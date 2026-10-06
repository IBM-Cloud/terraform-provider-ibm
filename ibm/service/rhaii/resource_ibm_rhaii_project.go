// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package rhaii

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	rc "github.com/IBM/platform-services-go-sdk/resourcecontrollerv2"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const rhaiiProjectResourceName = "ibm_rhaii_project"

// ResourceIBMRhaiiProject manages a Red Hat AI Inference (RHAII) project.
// A project is a resource controller instance of the "instructlab" service.
// The resource talks to the Global Catalog, Resource Controller and Global
// Tagging APIs directly, so it does not depend on ibm_resource_instance.
func ResourceIBMRhaiiProject() *schema.Resource {
	s := map[string]*schema.Schema{
		"name": {
			Type:        schema.TypeString,
			Required:    true,
			Description: "The name of the project.",
		},
		"location": {
			Type:        schema.TypeString,
			Optional:    true,
			ForceNew:    true,
			Default:     rhaiiDefaultLocation,
			Description: "The region where the project is created.",
		},
		"plan": {
			Type:        schema.TypeString,
			Optional:    true,
			Default:     rhaiiDefaultPlan,
			Description: "The pricing plan of the project.",
		},
		"resource_group_id": {
			Type:        schema.TypeString,
			Optional:    true,
			Computed:    true,
			ForceNew:    true,
			Description: "The ID of the resource group of the project. The default resource group of the account is used when not set.",
		},
		"tags": {
			Type:        schema.TypeSet,
			Optional:    true,
			Computed:    true,
			Elem:        &schema.Schema{Type: schema.TypeString, ValidateFunc: validate.InvokeValidator(rhaiiProjectResourceName, "tags")},
			Set:         flex.ResourceIBMVPCHash,
			Description: "The user tags of the project.",
		},
		"access_tags": {
			Type:        schema.TypeSet,
			Optional:    true,
			Computed:    true,
			Elem:        &schema.Schema{Type: schema.TypeString, ValidateFunc: validate.InvokeValidator(rhaiiProjectResourceName, "access_tags")},
			Set:         flex.ResourceIBMVPCHash,
			Description: "The access management tags of the project. The tags must already exist in the account.",
		},
	}
	for k, v := range rhaiiProjectComputedSchema() {
		s[k] = v
	}

	return &schema.Resource{
		CreateContext: resourceIBMRhaiiProjectCreate,
		ReadContext:   resourceIBMRhaiiProjectRead,
		UpdateContext: resourceIBMRhaiiProjectUpdate,
		DeleteContext: resourceIBMRhaiiProjectDelete,
		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		CustomizeDiff: customdiff.Sequence(
			func(_ context.Context, diff *schema.ResourceDiff, _ interface{}) error {
				return flex.ResourceTagsCustomizeDiff(diff)
			},
		),

		Schema: s,
	}
}

// ResourceIBMRhaiiProjectValidator validates the arguments of ibm_rhaii_project.
func ResourceIBMRhaiiProjectValidator() *validate.ResourceValidator {
	validateSchema := []validate.ValidateSchema{
		{
			Identifier:                 "tags",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^[A-Za-z0-9:_ .-]+$`,
			MinValueLength:             1,
			MaxValueLength:             128,
		},
		{
			Identifier:                 "access_tags",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^([A-Za-z0-9_.-]|[A-Za-z0-9_.-][A-Za-z0-9_ .-]*[A-Za-z0-9_.-]):([A-Za-z0-9_.-]|[A-Za-z0-9_.-][A-Za-z0-9_ .-]*[A-Za-z0-9_.-])$`,
			MinValueLength:             1,
			MaxValueLength:             128,
		},
	}
	return &validate.ResourceValidator{ResourceName: rhaiiProjectResourceName, Schema: validateSchema}
}

func resourceIBMRhaiiProjectCreate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	rcClient, err := meta.(conns.ClientSession).ResourceControllerV2API()
	if err != nil {
		return rhaiiDiag(err, "create", "initialize-client")
	}

	gcClient, err := meta.(conns.ClientSession).GlobalCatalogV1API()
	if err != nil {
		return rhaiiDiag(err, "create", "initialize-client")
	}

	location := d.Get("location").(string)
	planID, err := rhaiiResolvePlanID(ctx, gcClient, d.Get("plan").(string))
	if err != nil {
		return rhaiiDiag(err, "create", "resolve-plan")
	}
	targetCRN, err := rhaiiResolveTargetCRN(ctx, gcClient, planID, location)
	if err != nil {
		return rhaiiDiag(err, "create", "resolve-deployment")
	}

	resourceGroupID := d.Get("resource_group_id").(string)
	if resourceGroupID == "" {
		resourceGroupID, err = flex.DefaultResourceGroup(meta)
		if err != nil {
			return rhaiiDiag(err, "create", "default-resource-group")
		}
	}

	createOptions := &rc.CreateResourceInstanceOptions{
		Name:           ptr(d.Get("name").(string)),
		Target:         &targetCRN,
		ResourceGroup:  &resourceGroupID,
		ResourcePlanID: &planID,
	}
	instance, _, err := rcClient.CreateResourceInstanceWithContext(ctx, createOptions)
	if err != nil {
		return rhaiiDiag(fmt.Errorf("CreateResourceInstanceWithContext failed: %s", err), "create", "create-instance")
	}
	d.SetId(*instance.ID)

	if err := waitForRhaiiProject(ctx, rcClient, d.Id(), rhaiiWaitCreating, d.Timeout(schema.TimeoutCreate), rhaiiCreateRefresh); err != nil {
		return rhaiiDiag(fmt.Errorf("error waiting for project (%s) to be active: %s", d.Id(), err), "create", "wait-for-state")
	}

	// Tag errors are logged and not returned, so that a created project is
	// never lost from the state. The next plan shows the difference.
	if _, ok := d.GetOk("tags"); ok || os.Getenv("IC_ENV_TAGS") != "" {
		oldList, newList := d.GetChange("tags")
		if err := flex.UpdateTagsUsingCRN(oldList, newList, meta, *instance.CRN); err != nil {
			log.Printf("[ERROR] Error on create of %s (%s) tags: %s", rhaiiProjectResourceName, d.Id(), err)
		}
	}
	if _, ok := d.GetOk("access_tags"); ok {
		oldList, newList := d.GetChange("access_tags")
		if err := flex.UpdateGlobalTagsUsingCRN(oldList, newList, meta, *instance.CRN, "", rhaiiAccessTagType); err != nil {
			log.Printf("[ERROR] Error on create of %s (%s) access tags: %s", rhaiiProjectResourceName, d.Id(), err)
		}
	}

	return resourceIBMRhaiiProjectRead(ctx, d, meta)
}

func resourceIBMRhaiiProjectRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	rcClient, err := meta.(conns.ClientSession).ResourceControllerV2API()
	if err != nil {
		return rhaiiDiag(err, "read", "initialize-client")
	}

	instance, resp, err := rcClient.GetResourceInstanceWithContext(ctx, &rc.GetResourceInstanceOptions{ID: ptr(d.Id())})
	if err != nil {
		if resp != nil && resp.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		return rhaiiDiag(fmt.Errorf("GetResourceInstanceWithContext failed: %s", err), "read", "get-instance")
	}
	if isRhaiiInstanceGone(instance) {
		log.Printf("[WARN] Removing %s (%s) from state because it is in state %q", rhaiiProjectResourceName, d.Id(), *instance.State)
		d.SetId("")
		return nil
	}
	if !isRhaiiInstance(instance) {
		return rhaiiDiag(fmt.Errorf("resource instance %s is not a Red Hat AI Inference project (service is %q, expected %q)", d.Id(), strDeref(instance.ResourceID), rhaiiServiceName), "read", "check-service")
	}

	// Normalize the ID to the CRN, so that an import by GUID gives the same state as a create.
	d.SetId(*instance.ID)

	if err := setRhaiiProjectAttributes(ctx, d, meta, instance); err != nil {
		return rhaiiDiag(err, "read", "set-attributes")
	}

	tags, err := flex.GetTagsUsingCRN(meta, *instance.CRN)
	if err != nil {
		log.Printf("[ERROR] Error on get of %s (%s) tags: %s", rhaiiProjectResourceName, d.Id(), err)
	}
	d.Set("tags", tags)

	accessTags, err := flex.GetGlobalTagsUsingCRN(meta, *instance.CRN, "", rhaiiAccessTagType)
	if err != nil {
		log.Printf("[ERROR] Error on get of %s (%s) access tags: %s", rhaiiProjectResourceName, d.Id(), err)
	}
	d.Set("access_tags", accessTags)

	return nil
}

func resourceIBMRhaiiProjectUpdate(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	rcClient, err := meta.(conns.ClientSession).ResourceControllerV2API()
	if err != nil {
		return rhaiiDiag(err, "update", "initialize-client")
	}

	if d.HasChange("name") || d.HasChange("plan") {
		updateOptions := &rc.UpdateResourceInstanceOptions{ID: ptr(d.Id())}
		if d.HasChange("name") {
			updateOptions.Name = ptr(d.Get("name").(string))
		}
		if d.HasChange("plan") {
			gcClient, err := meta.(conns.ClientSession).GlobalCatalogV1API()
			if err != nil {
				return rhaiiDiag(err, "update", "initialize-client")
			}
			planID, err := rhaiiResolvePlanID(ctx, gcClient, d.Get("plan").(string))
			if err != nil {
				return rhaiiDiag(err, "update", "resolve-plan")
			}
			updateOptions.ResourcePlanID = &planID
		}

		if _, _, err := rcClient.UpdateResourceInstanceWithContext(ctx, updateOptions); err != nil {
			return rhaiiDiag(fmt.Errorf("UpdateResourceInstanceWithContext failed: %s", err), "update", "update-instance")
		}
		if err := waitForRhaiiProject(ctx, rcClient, d.Id(), rhaiiWaitUpdating, d.Timeout(schema.TimeoutUpdate), rhaiiUpdateRefresh); err != nil {
			return rhaiiDiag(fmt.Errorf("error waiting for project (%s) to be updated: %s", d.Id(), err), "update", "wait-for-state")
		}
	}

	crn := d.Get("crn").(string)
	if d.HasChange("tags") {
		oldList, newList := d.GetChange("tags")
		if err := flex.UpdateTagsUsingCRN(oldList, newList, meta, crn); err != nil {
			return rhaiiDiag(fmt.Errorf("error updating tags: %s", err), "update", "update-tags")
		}
	}
	if d.HasChange("access_tags") {
		oldList, newList := d.GetChange("access_tags")
		if err := flex.UpdateGlobalTagsUsingCRN(oldList, newList, meta, crn, "", rhaiiAccessTagType); err != nil {
			return rhaiiDiag(fmt.Errorf("error updating access tags: %s", err), "update", "update-access-tags")
		}
	}

	return resourceIBMRhaiiProjectRead(ctx, d, meta)
}

func resourceIBMRhaiiProjectDelete(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	rcClient, err := meta.(conns.ClientSession).ResourceControllerV2API()
	if err != nil {
		return rhaiiDiag(err, "delete", "initialize-client")
	}

	resp, err := rcClient.DeleteResourceInstanceWithContext(ctx, &rc.DeleteResourceInstanceOptions{
		ID:        ptr(d.Id()),
		Recursive: flex.PtrToBool(true),
	})
	if err != nil {
		if resp != nil && (resp.StatusCode == 404 || resp.StatusCode == 410) {
			d.SetId("")
			return nil
		}
		return rhaiiDiag(fmt.Errorf("DeleteResourceInstanceWithContext failed: %s", err), "delete", "delete-instance")
	}

	if err := waitForRhaiiProject(ctx, rcClient, d.Id(), rhaiiWaitDeleting, d.Timeout(schema.TimeoutDelete), rhaiiDeleteRefresh); err != nil {
		return rhaiiDiag(fmt.Errorf("error waiting for project (%s) to be deleted: %s", d.Id(), err), "delete", "wait-for-state")
	}

	d.SetId("")
	return nil
}

// rhaiiDiag converts an error to diagnostics in the format used across the provider.
func rhaiiDiag(err error, operation, discriminator string) diag.Diagnostics {
	return rhaiiDiagFor(err, rhaiiProjectResourceName, operation, discriminator)
}

func rhaiiDiagFor(err error, resource, operation, discriminator string) diag.Diagnostics {
	tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), resource, operation, discriminator)
	log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
	return tfErr.GetDiag()
}
