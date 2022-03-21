// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc

import (
	"context"
	"fmt"
	"log"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/vpc-go-sdk/vpcv1"
)

func ResourceIBMIsBackupPolicy() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMIsBackupPolicyCreate,
		ReadContext:   resourceIBMIsBackupPolicyRead,
		UpdateContext: resourceIBMIsBackupPolicyUpdate,
		DeleteContext: resourceIBMIsBackupPolicyDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"match_resource_types": &schema.Schema{
				Type:        schema.TypeSet,
				Optional:    true,
				Computed:    true,
				Set:         schema.HashString,
				Description: "A resource type this backup policy applies to. Resources that have both a matching type and a matching user tag will be subject to the backup policy.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"match_user_tags": &schema.Schema{
				Type:        schema.TypeSet,
				Required:    true,
				Description: "The user tags this backup policy applies to. Resources that have both a matching user tag and a matching type will be subject to the backup policy.",
				Elem:        &schema.Schema{Type: schema.TypeString},
				Set:         schema.HashString,
			},
			"name": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validate.InvokeValidator("ibm_is_backup_policy", "name"),
				Description:  "The user-defined name for this backup policy. Names must be unique within the region this backup policy resides in. If unspecified, the name will be a hyphenated list of randomly-selected words.",
			},
			"resource_group": &schema.Schema{
				Type:        schema.TypeList,
				MaxItems:    1,
				ForceNew:    true,
				Optional:    true,
				Computed:    true,
				Description: "The resource group to use. If unspecified, the account's [default resourcegroup](https://cloud.ibm.com/apidocs/resource-manager#introduction) is used.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The unique identifier for this resource group.",
						},
					},
				},
			},
			"created_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time that the backup policy was created.",
			},
			"crn": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The CRN for this backup policy.",
			},
			"href": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL for this backup policy.",
			},
			"last_job_completed_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The date and time that the most recent job for this backup policy completed.",
			},
			"lifecycle_state": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The lifecycle state of the backup policy.",
			},
			"resource_type": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The resource type.",
			},
			"version": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

func ResourceIBMIsBackupPolicyValidator() *validate.ResourceValidator {
	validateSchema := make([]validate.ValidateSchema, 1)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "name",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^([a-z]|[a-z][-a-z0-9]*[a-z0-9]|[0-9][-a-z0-9]*([a-z]|[-a-z][-a-z0-9]*[a-z0-9]))$`,
			MinValueLength:             1,
			MaxValueLength:             63,
		},
	)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "match_user_tags",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^[A-Za-z0-9:_ .-]+$`,
			MinValueLength:             1,
			MaxValueLength:             128,
		},
	)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "match_resource_types",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^[a-z][a-z0-9]*(_[a-z0-9]+)*$`,
			MinValueLength:             1,
			MaxValueLength:             128,
		},
	)
	resourceValidator := validate.ResourceValidator{ResourceName: "ibm_is_backup_policy", Schema: validateSchema}
	return &resourceValidator
}

func resourceIBMIsBackupPolicyCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := vpcClient(meta)
	if err != nil {
		return diag.FromErr(err)
	}

	createBackupPolicyOptions := &vpcv1.CreateBackupPolicyOptions{}

	if _, ok := d.GetOk("match_resource_types"); ok {
		createBackupPolicyOptions.SetMatchResourceTypes(flex.ExpandStringList((d.Get("match_resource_types").(*schema.Set)).List()))
	}
	if _, ok := d.GetOk("match_user_tags"); ok {
		createBackupPolicyOptions.SetMatchUserTags((flex.ExpandStringList((d.Get("match_user_tags").(*schema.Set)).List())))
	}
	if _, ok := d.GetOk("name"); ok {
		createBackupPolicyOptions.SetName(d.Get("name").(string))
	}
	if _, ok := d.GetOk("resource_group"); ok {
		resourceGroup := resourceIBMIsBackupPolicyMapToResourceGroupIdentity(d.Get("resource_group.0").(map[string]interface{}))
		createBackupPolicyOptions.ResourceGroup = resourceGroup
	}

	backupPolicy, response, err := vpcClient.CreateBackupPolicyWithContext(context, createBackupPolicyOptions)
	if err != nil {
		log.Printf("[DEBUG] CreateBackupPolicyWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("CreateBackupPolicyWithContext failed %s\n%s", err, response))
	}

	d.SetId(*backupPolicy.ID)

	return resourceIBMIsBackupPolicyRead(context, d, meta)
}

func resourceIBMIsBackupPolicyMapToResourceGroupIdentity(resourceGroupIdentityMap map[string]interface{}) vpcv1.ResourceGroupIdentityIntf {
	resourceGroupIdentity := vpcv1.ResourceGroupIdentity{}

	if resourceGroupIdentityMap["id"] != nil {
		resourceGroupIdentity.ID = core.StringPtr(resourceGroupIdentityMap["id"].(string))
	}

	return &resourceGroupIdentity
}

func resourceIBMIsBackupPolicyRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := vpcClient(meta)
	if err != nil {
		return diag.FromErr(err)
	}

	getBackupPolicyOptions := &vpcv1.GetBackupPolicyOptions{}

	getBackupPolicyOptions.SetID(d.Id())

	backupPolicy, response, err := vpcClient.GetBackupPolicyWithContext(context, getBackupPolicyOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		log.Printf("[DEBUG] GetBackupPolicyWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("GetBackupPolicyWithContext failed %s\n%s", err, response))
	}

	if backupPolicy.MatchResourceTypes != nil {
		if err = d.Set("match_resource_types", backupPolicy.MatchResourceTypes); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting match_resource_types: %s", err))
		}
	}
	if backupPolicy.MatchUserTags != nil {
		if err = d.Set("match_user_tags", backupPolicy.MatchUserTags); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting match_user_tags: %s", err))
		}
	}
	if backupPolicy.Name != nil {
		if err = d.Set("name", backupPolicy.Name); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting name: %s", err))
		}
	}
	if backupPolicy.ResourceGroup != nil {
		resourceGroupMap := resourceIBMIsBackupPolicyResourceGroupIdentityToMap(*backupPolicy.ResourceGroup)
		if err = d.Set("resource_group", []map[string]interface{}{resourceGroupMap}); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting resource_group: %s", err))
		}
	}
	if backupPolicy.CreatedAt != nil {
		if err = d.Set("created_at", flex.DateTimeToString(backupPolicy.CreatedAt)); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting created_at: %s", err))
		}
	}

	if backupPolicy.CRN != nil {
		if err = d.Set("crn", backupPolicy.CRN); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting crn: %s", err))
		}
	}

	if backupPolicy.Href != nil {
		if err = d.Set("href", backupPolicy.Href); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting href: %s", err))
		}
	}

	if backupPolicy.LastJobCompletedAt != nil {
		if err = d.Set("last_job_completed_at", flex.DateTimeToString(backupPolicy.LastJobCompletedAt)); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting last_job_completed_at: %s", err))
		}
	}

	if backupPolicy.LifecycleState != nil {
		if err = d.Set("lifecycle_state", backupPolicy.LifecycleState); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting lifecycle_state: %s", err))
		}
	}

	if backupPolicy.ResourceType != nil {
		if err = d.Set("resource_type", backupPolicy.ResourceType); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting resource_type: %s", err))
		}
	}

	if err = d.Set("version", response.Headers.Get("Etag")); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting version: %s", err))
	}

	return nil
}
func resourceIBMIsBackupPolicyResourceGroupIdentityToMap(resourceGroupIdentity vpcv1.ResourceGroupReference) map[string]interface{} {
	resourceGroupIdentityMap := map[string]interface{}{}
	resourceGroupIdentityMap["id"] = resourceGroupIdentity.ID
	return resourceGroupIdentityMap
}

func resourceIBMIsBackupPolicyUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := vpcClient(meta)
	if err != nil {
		return diag.FromErr(err)
	}
	updateBackupPolicyOptions := &vpcv1.UpdateBackupPolicyOptions{}
	updateBackupPolicyOptions.SetID(d.Id())
	hasChange := false
	patchVals := &vpcv1.BackupPolicyPatch{}
	if d.HasChange("match_user_tags") {
		patchVals.MatchUserTags = (flex.ExpandStringList((d.Get("match_user_tags").(*schema.Set)).List()))
		hasChange = true
	}
	if d.HasChange("name") {
		patchVals.Name = core.StringPtr(d.Get("name").(string))
		hasChange = true
	}
	updateBackupPolicyOptions.SetIfMatch(d.Get("version").(string))
	if hasChange {
		updateBackupPolicyOptions.BackupPolicyPatch, _ = patchVals.AsPatch()
		_, response, err := vpcClient.UpdateBackupPolicyWithContext(context, updateBackupPolicyOptions)
		if err != nil {
			log.Printf("[DEBUG] UpdateBackupPolicyWithContext failed %s\n%s", err, response)
			return diag.FromErr(fmt.Errorf("UpdateBackupPolicyWithContext failed %s\n%s", err, response))
		}
	}

	if d.HasChange("plans") && !d.IsNewResource() {
		getBackupPolicyOptions := &vpcv1.GetBackupPolicyOptions{}
		getBackupPolicyOptions.SetID(d.Id())
		backupPolicy, response, err := vpcClient.GetBackupPolicyWithContext(context, getBackupPolicyOptions)
		if err != nil {
			if response != nil && response.StatusCode == 404 {
				d.SetId("")
				return nil
			}
			log.Printf("[DEBUG] GetBackupPolicyWithContext failed %s\n%s", err, response)
			return diag.FromErr(fmt.Errorf("GetBackupPolicyWithContext failed %s\n%s", err, response))
		}
		backupPolicyPlanID := backupPolicy.Plans[0].ID
		getBackupPolicyPlanOptions := &vpcv1.GetBackupPolicyPlanOptions{}
		getBackupPolicyPlanOptions.SetBackupPolicyID(d.Id())
		getBackupPolicyPlanOptions.SetID(*backupPolicyPlanID)
		_, response, err = vpcClient.GetBackupPolicyPlanWithContext(context, getBackupPolicyPlanOptions)
		if err != nil {
			if response != nil && response.StatusCode == 404 {
				d.SetId("")
				return nil
			}
			log.Printf("[DEBUG] GetBackupPolicyPlanWithContext failed %s\n%s", err, response)
			return diag.FromErr(fmt.Errorf("GetBackupPolicyPlanWithContext failed %s\n%s", err, response))
		}
		etag := response.Headers.Get("E-tag")
		plans := d.Get("plans").([]interface{})
		for i := range plans {
			updateBackupPolicyPlanOptions := &vpcv1.UpdateBackupPolicyPlanOptions{}
			updateBackupPolicyPlanOptions.SetBackupPolicyID(d.Id())
			updateBackupPolicyPlanOptions.SetID(*backupPolicyPlanID)
			hasChange := false
			patchVals := &vpcv1.BackupPolicyPlanPatch{}

			cronSpec := fmt.Sprintf("plans.%d.cron_spec", i)
			if d.HasChange(cronSpec) {
				patchVals.CronSpec = core.StringPtr(d.Get(cronSpec).(string))
				hasChange = true
			}

			active := fmt.Sprintf("plans.%d.active", i)
			if d.HasChange(active) {
				patchVals.Active = core.BoolPtr(d.Get(active).(bool))
				hasChange = true
			}

			attachUserTags := fmt.Sprintf("plans.%d.attach_user_tags", i)
			if d.HasChange(attachUserTags) {
				patchVals.AttachUserTags = (flex.ExpandStringList((d.Get(attachUserTags).(*schema.Set)).List()))
				hasChange = true
			}

			copyUserTags := fmt.Sprintf("plans.%d.copy_user_tags", i)
			if d.HasChange(copyUserTags) {
				patchVals.CopyUserTags = core.BoolPtr(d.Get(copyUserTags).(bool))
				hasChange = true
			}

			deletionTrigger := fmt.Sprintf("plans.%d.deletion_trigger", i)
			if d.HasChange(deletionTrigger) {
				deletionTriggerItem := fmt.Sprintf("plans.%d.deletion_trigger.0", i)
				deletionTrigger := resourceIBMIsBackupPolicyPlanMapToBackupPolicyPlanDeletionTriggerPatch(d.Get(deletionTriggerItem).(map[string]interface{}))
				patchVals.DeletionTrigger = &deletionTrigger
				hasChange = true
			}

			name := fmt.Sprintf("plans.%d.name", i)
			if d.HasChange(name) {
				patchVals.Name = core.StringPtr(d.Get(name).(string))
				hasChange = true
			}
			updateBackupPolicyPlanOptions.SetIfMatch(etag)

			if hasChange {
				updateBackupPolicyPlanOptions.BackupPolicyPlanPatch, _ = patchVals.AsPatch()
				_, response, err := vpcClient.UpdateBackupPolicyPlanWithContext(context, updateBackupPolicyPlanOptions)
				if err != nil {
					log.Printf("[DEBUG] UpdateBackupPolicyPlanWithContext failed %s\n%s", err, response)
					return diag.FromErr(fmt.Errorf("UpdateBackupPolicyPlanWithContext failed %s\n%s", err, response))
				}
			}

		}

	}

	return resourceIBMIsBackupPolicyRead(context, d, meta)
}

func resourceIBMIsBackupPolicyDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	vpcClient, err := vpcClient(meta)
	if err != nil {
		return diag.FromErr(err)
	}
	deleteBackupPolicyOptions := &vpcv1.DeleteBackupPolicyOptions{}
	deleteBackupPolicyOptions.SetID(d.Id())
	deleteBackupPolicyOptions.SetIfMatch(d.Get("version").(string))
	_, response, err := vpcClient.DeleteBackupPolicyWithContext(context, deleteBackupPolicyOptions)
	if err != nil {
		log.Printf("[DEBUG] DeleteBackupPolicyWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("DeleteBackupPolicyWithContext failed %s\n%s", err, response))
	}
	d.SetId("")
	return nil
}
