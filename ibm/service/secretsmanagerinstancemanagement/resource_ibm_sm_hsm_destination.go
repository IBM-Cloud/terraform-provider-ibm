// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.117.3-3e91ab11-20261028-182643
 */

package secretsmanagerinstancemanagement

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/secrets-manager-management-go-sdk/v2/secretsmanagerinstancemanagementv2"
)

func ResourceIbmSmHsmDestination() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIbmSmHsmDestinationCreate,
		ReadContext:   resourceIbmSmHsmDestinationRead,
		UpdateContext: resourceIbmSmHsmDestinationUpdate,
		DeleteContext: resourceIbmSmHsmDestinationDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"region": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Computed:    true,
				Description: "The region of the Secrets Manager instance.",
			},
			"instance_id": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.InvokeValidator("ibm_sm_hsm_destination", "instance_id"),
				Description:  "Secrets Manager instance ID.",
			},
			"name": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validate.InvokeValidator("ibm_sm_hsm_destination", "name"),
				Description:  "Human-readable display name.",
			},
			"description": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validate.InvokeValidator("ibm_sm_hsm_destination", "description"),
				Description:  "Optional description. Omitted if not set.",
			},
			"vendor": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.InvokeValidator("ibm_sm_hsm_destination", "vendor"),
				Description:  "HSM vendor. Immutable after creation.",
			},
			"href": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "URL of the destination resource.",
			},
			"account_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "IBM Cloud account ID. System-managed. Required for account policy setup.",
			},
			"provisioning_state": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Lifecycle state of the destination:- `completed`: Destination created and available for configuration. Initial state on POST.- `deleting`: Hard delete initiated; asynchronous cleanup in progress.- `deletion_failed`: Asynchronous delete failed; the destination still exists. Retry DELETE to re-initiate.",
			},
			"health_state": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "AS OF (point-in-time) health state. Written exclusively by POST …/health and POST …/apply:- `unknown`: Health check not yet performed, or destination is deleting/deletion_failed.- `evaluating`: A health check (POST …/health) is in progress. Set on acceptance; replaced when the check completes.- `in_progress`: An apply operation (POST …/apply) is running. Set on acceptance; replaced by the health check result once apply completes (or set to `unknown` on failure).- `not_configured`: Prerequisites not met (needs ≥1 HA Group, ≥1 path with active non-expired Server CA, and ≥1 active non-expired client certificate).- `healthy`: All non-`not_configured` HSM HA Groups are `healthy` or `degraded`.- `degraded`: Some non-`not_configured` HSM HA Groups are `unhealthy`/`unknown` AND at least one is `healthy`.- `unhealthy`: All non-`not_configured` HSM HA Groups are `unhealthy` or `unknown`.",
			},
			"health_as_of": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Timestamp of the last health evaluation. On creation equals `created_at`. Updated only by POST …/health.",
			},
			"created_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Computed field. When the resource was created.",
			},
			"updated_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Computed field. Omitted until first update.",
			},
			"created_by": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "IAM identity of creator. Computed field.",
			},
			"updated_by": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "IAM identity of last updater. Computed field. Omitted until first update.",
			},
			"destination_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Destination ID. Immutable. System-managed.",
			},
		},
	}
}

func ResourceIbmSmHsmDestinationValidator() *validate.ResourceValidator {
	validateSchema := make([]validate.ValidateSchema, 0)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "instance_id",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Required:                   true,
			Regexp:                     `^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`,
			MinValueLength:             36,
			MaxValueLength:             36,
		},
		validate.ValidateSchema{
			Identifier:                 "name",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Required:                   true,
			Regexp:                     `^[A-Za-z0-9_][A-Za-z0-9_]*(?:_*-*\.*[A-Za-z0-9]*)*[A-Za-z0-9]+$`,
			MinValueLength:             2,
			MaxValueLength:             128,
		},
		validate.ValidateSchema{
			Identifier:                 "description",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^(.*?)$`,
			MinValueLength:             0,
			MaxValueLength:             1024,
		},
		validate.ValidateSchema{
			Identifier:                 "vendor",
			ValidateFunctionIdentifier: validate.ValidateAllowedStringValue,
			Type:                       validate.TypeString,
			Required:                   true,
			AllowedValues:              "thales_luna",
		},
	)

	resourceValidator := validate.ResourceValidator{ResourceName: "ibm_sm_hsm_destination", Schema: validateSchema}
	return &resourceValidator
}

func resourceIbmSmHsmDestinationCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	secretsManagerInstanceManagementClient, err := getInstanceManagementClient(d, meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_destination", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	region := getRegion(secretsManagerInstanceManagementClient, d)
	if err = d.Set("region", region); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting region"), "ibm_sm_hsm_destination", "create")
		return tfErr.GetDiag()
	}

	createInstanceHsmDestinationOptions := &secretsmanagerinstancemanagementv2.CreateInstanceHsmDestinationOptions{}

	createInstanceHsmDestinationOptions.SetInstanceID(d.Get("instance_id").(string))
	createInstanceHsmDestinationOptions.SetName(d.Get("name").(string))
	createInstanceHsmDestinationOptions.SetVendor(d.Get("vendor").(string))
	if _, ok := d.GetOk("description"); ok {
		createInstanceHsmDestinationOptions.SetDescription(d.Get("description").(string))
	}

	hsmDestination, _, err := secretsManagerInstanceManagementClient.CreateInstanceHsmDestinationWithContext(context, createInstanceHsmDestinationOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreateInstanceHsmDestinationWithContext failed: %s", err.Error()), "ibm_sm_hsm_destination", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(fmt.Sprintf("%s/%s", *createInstanceHsmDestinationOptions.InstanceID, *hsmDestination.ID))

	return resourceIbmSmHsmDestinationRead(context, d, meta)
}

func resourceIbmSmHsmDestinationRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	secretsManagerInstanceManagementClient, err := getInstanceManagementClient(d, meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_destination", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getInstanceHsmDestinationOptions := &secretsmanagerinstancemanagementv2.GetInstanceHsmDestinationOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_destination", "read", "sep-id-parts").GetDiag()
	}

	getInstanceHsmDestinationOptions.SetInstanceID(parts[0])
	getInstanceHsmDestinationOptions.SetDestinationID(parts[1])

	hsmDestination, response, err := secretsManagerInstanceManagementClient.GetInstanceHsmDestinationWithContext(context, getInstanceHsmDestinationOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetInstanceHsmDestinationWithContext failed: %s", err.Error()), "ibm_sm_hsm_destination", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	if err = d.Set("instance_id", parts[0]); err != nil {
		err = fmt.Errorf("error setting instance_id: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_destination", "read", "set-instance_id").GetDiag()
	}
	if err = d.Set("name", hsmDestination.Name); err != nil {
		err = fmt.Errorf("error setting name: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_destination", "read", "set-name").GetDiag()
	}
	if !core.IsNil(hsmDestination.Description) {
		if err = d.Set("description", hsmDestination.Description); err != nil {
			err = fmt.Errorf("error setting description: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_destination", "read", "set-description").GetDiag()
		}
	}
	if err = d.Set("vendor", hsmDestination.Vendor); err != nil {
		err = fmt.Errorf("error setting vendor: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_destination", "read", "set-vendor").GetDiag()
	}
	if !core.IsNil(hsmDestination.Href) {
		if err = d.Set("href", hsmDestination.Href); err != nil {
			err = fmt.Errorf("error setting href: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_destination", "read", "set-href").GetDiag()
		}
	}
	if !core.IsNil(hsmDestination.AccountID) {
		if err = d.Set("account_id", hsmDestination.AccountID); err != nil {
			err = fmt.Errorf("error setting account_id: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_destination", "read", "set-account_id").GetDiag()
		}
	}
	if err = d.Set("provisioning_state", hsmDestination.ProvisioningState); err != nil {
		err = fmt.Errorf("error setting provisioning_state: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_destination", "read", "set-provisioning_state").GetDiag()
	}
	if err = d.Set("health_state", hsmDestination.HealthState); err != nil {
		err = fmt.Errorf("error setting health_state: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_destination", "read", "set-health_state").GetDiag()
	}
	if err = d.Set("health_as_of", flex.DateTimeToString(hsmDestination.HealthAsOf)); err != nil {
		err = fmt.Errorf("error setting health_as_of: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_destination", "read", "set-health_as_of").GetDiag()
	}
	if err = d.Set("created_at", flex.DateTimeToString(hsmDestination.CreatedAt)); err != nil {
		err = fmt.Errorf("error setting created_at: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_destination", "read", "set-created_at").GetDiag()
	}
	if !core.IsNil(hsmDestination.UpdatedAt) {
		if err = d.Set("updated_at", flex.DateTimeToString(hsmDestination.UpdatedAt)); err != nil {
			err = fmt.Errorf("error setting updated_at: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_destination", "read", "set-updated_at").GetDiag()
		}
	}
	if err = d.Set("created_by", hsmDestination.CreatedBy); err != nil {
		err = fmt.Errorf("error setting created_by: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_destination", "read", "set-created_by").GetDiag()
	}
	if !core.IsNil(hsmDestination.UpdatedBy) {
		if err = d.Set("updated_by", hsmDestination.UpdatedBy); err != nil {
			err = fmt.Errorf("error setting updated_by: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_destination", "read", "set-updated_by").GetDiag()
		}
	}
	if !core.IsNil(hsmDestination.ID) {
		if err = d.Set("destination_id", hsmDestination.ID); err != nil {
			err = fmt.Errorf("error setting destination_id: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_destination", "read", "set-destination_id").GetDiag()
		}
	}

	return nil
}

func resourceIbmSmHsmDestinationUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	secretsManagerInstanceManagementClient, err := getInstanceManagementClient(d, meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_destination", "update", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	updateInstanceHsmDestinationOptions := &secretsmanagerinstancemanagementv2.UpdateInstanceHsmDestinationOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_destination", "update", "sep-id-parts").GetDiag()
	}

	updateInstanceHsmDestinationOptions.SetInstanceID(parts[0])
	updateInstanceHsmDestinationOptions.SetDestinationID(parts[1])

	hasChange := false

	patchVals := &secretsmanagerinstancemanagementv2.HSMDestinationPatch{}
	if d.HasChange("instance_id") {
		errMsg := fmt.Sprintf("cannot update resource property \"%s\" with the ForceNew annotation."+
			" The resource must be re-created to update this property.", "instance_id")
		return flex.DiscriminatedTerraformErrorf(nil, errMsg, "ibm_sm_hsm_destination", "update", "instance_id-forces-new").GetDiag()
	}
	if d.HasChange("name") {
		newName := d.Get("name").(string)
		patchVals.Name = &newName
		hasChange = true
	}
	if d.HasChange("description") {
		newDescription := d.Get("description").(string)
		patchVals.Description = &newDescription
		hasChange = true
	}

	if hasChange {
		// Fields with `nil` values are omitted from the generic map,
		// so we need to re-add them to support removing arguments
		// in merge-patch operations sent to the service.
		updateInstanceHsmDestinationOptions.HsmDestinationPatch = ResourceIbmSmHsmDestinationHSMDestinationPatchAsPatch(patchVals, d)

		_, _, err = secretsManagerInstanceManagementClient.UpdateInstanceHsmDestinationWithContext(context, updateInstanceHsmDestinationOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("UpdateInstanceHsmDestinationWithContext failed: %s", err.Error()), "ibm_sm_hsm_destination", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
	}

	return resourceIbmSmHsmDestinationRead(context, d, meta)
}

func resourceIbmSmHsmDestinationDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	secretsManagerInstanceManagementClient, err := getInstanceManagementClient(d, meta)
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_destination", "delete", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	deleteInstanceHsmDestinationOptions := &secretsmanagerinstancemanagementv2.DeleteInstanceHsmDestinationOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_sm_hsm_destination", "delete", "sep-id-parts").GetDiag()
	}

	deleteInstanceHsmDestinationOptions.SetInstanceID(parts[0])
	deleteInstanceHsmDestinationOptions.SetDestinationID(parts[1])

	_, _, err = secretsManagerInstanceManagementClient.DeleteInstanceHsmDestinationWithContext(context, deleteInstanceHsmDestinationOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DeleteInstanceHsmDestinationWithContext failed: %s", err.Error()), "ibm_sm_hsm_destination", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	_, err = waitForIbmSmHsmDestinationDelete(d, meta)
	if err != nil {
		errMsg := fmt.Sprintf("error waiting for resource IbmSmHsmDestination (%s) to be deleted: %s", d.Id(), err)
		return flex.DiscriminatedTerraformErrorf(err, errMsg, "ibm_sm_hsm_destination", "delete", "wait-for-state").GetDiag()
	}

	d.SetId("")

	return nil
}

func waitForIbmSmHsmDestinationDelete(d *schema.ResourceData, meta interface{}) (interface{}, error) {
	secretsManagerInstanceManagementClient, err := getInstanceManagementClient(d, meta)
	if err != nil {
		return false, err
	}
	getInstanceHsmDestinationOptions := &secretsmanagerinstancemanagementv2.GetInstanceHsmDestinationOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return false, err
	}

	getInstanceHsmDestinationOptions.SetInstanceID(parts[0])
	getInstanceHsmDestinationOptions.SetDestinationID(parts[1])

	stateConf := &resource.StateChangeConf{
		Pending: []string{"deleting"},
		Target:  []string{"success", "succeeded", "valid", "active", "done", "normal", "ready", "provisioned", "available", "deployed", "healthy", "pending_reclamation", "deleted", "removed", "stable", "suspended"},
		Refresh: func() (interface{}, string, error) {
			stateObj, response, err := secretsManagerInstanceManagementClient.GetInstanceHsmDestination(getInstanceHsmDestinationOptions)
			if err != nil {
				if _, ok := err.(*core.SDKProblem); ok && response != nil && response.GetStatusCode() == 404 {
					return response, "success", nil
				}
				return nil, "", err
			}
			failStates := map[string]bool{"deletion_failed": true}
			if failStates[*stateObj.ProvisioningState] {
				return stateObj, *stateObj.ProvisioningState, fmt.Errorf("the instance %s failed: %s", "getInstanceHsmDestinationOptions", err)
			}
			return stateObj, *stateObj.ProvisioningState, nil
		},
		Timeout:    d.Timeout(schema.TimeoutCreate),
		Delay:      60 * time.Second,
		MinTimeout: 60 * time.Second,
	}

	return stateConf.WaitForStateContext(context.Background())
}

func ResourceIbmSmHsmDestinationHSMDestinationPatchAsPatch(patchVals *secretsmanagerinstancemanagementv2.HSMDestinationPatch, d *schema.ResourceData) map[string]interface{} {
	patch, _ := patchVals.AsPatch()
	var path string

	path = "name"
	if _, exists := d.GetOk(path); d.HasChange(path) && !exists {
		patch["name"] = nil
	} else if !exists {
		delete(patch, "name")
	}
	path = "description"
	if _, exists := d.GetOk(path); d.HasChange(path) && !exists {
		patch["description"] = nil
	} else if !exists {
		delete(patch, "description")
	}

	return patch
}
