// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package backuprecovery

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.ibm.com/BackupAndRecovery/ibm-backup-recovery-sdk-go/backuprecoveryv1"
)

func ResourceIbmProtectionGroupState() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIbmProtectionGroupStateCreate,
		ReadContext:   resourceIbmProtectionGroupStateRead,
		DeleteContext: resourceIbmProtectionGroupStateDelete,
		UpdateContext: resourceIbmProtectionGroupStateUpdate,
		CustomizeDiff: checkResourceIbmProtectionGroupStateDiff,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"action": &schema.Schema{
				Type: schema.TypeString,
				//ForceNew:: true,
				Required: true,
				// ValidateFunc: validate.InvokeValidator("ibm_protection_group_state", "action"),
				Description: "Specifies the action to be performed on all the specfied Protection Groups. 'kActivate' specifies that Protection Group should be activated. 'kDeactivate' sepcifies that Protection Group should be deactivated. 'kPause' specifies that Protection Group should be paused. 'kResume' specifies that Protection Group should be resumed.",
			},
			"ids": &schema.Schema{
				Type: schema.TypeList,
				//ForceNew::    true,
				Required:    true,
				Description: "Specifies a list of Protection Group ids for which the state should change.",
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"successful_group_ids": &schema.Schema{
				Type: schema.TypeString,
				//ForceNew::    true,
				Computed:    true,
				Description: "Specifies a list of Protection Group ids for which the state should change.",
			},
			"failed_groups": &schema.Schema{
				Type: schema.TypeList,
				//ForceNew::    true,
				Computed:    true,
				Description: "Specfies the list of connections for the source.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"run_id": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Specifies the id of the connection.",
						},
						"error_message": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Specifies the entity id of the source. The source can a non-root entity.",
						},
					},
				},
			},
		},
	}
}

func checkResourceIbmProtectionGroupStateDiff(context context.Context, d *schema.ResourceDiff, meta interface{}) error {
	// skip if it's a new resource
	if d.Id() == "" {
		return nil
	}

	resourceSchema := ResourceIbmProtectionGroupState().Schema

	// handle update resource
LOOP:
	for key := range resourceSchema {
		if d.HasChange(key) {
			log.Println("[WARNING] Update operation is not supported for this resource. No changes will be applied.")
			break LOOP
		}
	}

	return nil
}

func ResourceIbmProtectionGroupStateValidator() *validate.ResourceValidator {
	validateSchema := make([]validate.ValidateSchema, 0)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "action",
			ValidateFunctionIdentifier: validate.ValidateAllowedStringValue,
			Type:                       validate.TypeString,
			Required:                   true,
			AllowedValues:              "kPause, kResume",
		},
	)

	resourceValidator := validate.ResourceValidator{ResourceName: "ibm_protection_group_state", Schema: validateSchema}
	return &resourceValidator
}

func resourceIbmProtectionGroupStateCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	backupRecoveryClient, err := meta.(conns.ClientSession).BackupRecoveryV1()
	if err != nil {
		return diag.FromErr(err)
	}

	updateProtectionGroupsStateOptions := &backuprecoveryv1.UpdateProtectionGroupsStateOptions{}

	updateProtectionGroupsStateOptions.SetAction(d.Get("action").(string))
	var ids []string
	for _, v := range d.Get("ids").([]interface{}) {
		idsItem := v.(string)
		ids = append(ids, idsItem)
	}
	updateProtectionGroupsStateOptions.SetIds(ids)

	updateProtectionGroupsState, response, err := backupRecoveryClient.UpdateProtectionGroupsStateWithContext(context, updateProtectionGroupsStateOptions)
	if err != nil {
		log.Printf("[DEBUG] UpdateProtectionGroupsStateWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("UpdateProtectionGroupsStateWithContext failed %s\n%s", err, response))
	}

	if err = d.Set("successful_group_ids", strings.Join(updateProtectionGroupsState.SuccessfulProtectionGroupIds[:], ",")); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting successful_group_ids: %s", err))
	}

	if len(updateProtectionGroupsState.FailedProtectionGroups) > 0 {
		failedGroups := make([]string, 0)
		for _, failedRun := range updateProtectionGroupsState.FailedProtectionGroups {
			failedGroups = append(failedGroups, fmt.Sprintf("run_id: %s, error_message: %s", *failedRun.ProtectionGroupID, *failedRun.ErrorMessage))
		}
		if err = d.Set("failed_groups", strings.Join(failedGroups[:], ",")); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting failed_groups: %s", err))
		}
	}

	if !core.IsNil(updateProtectionGroupsState.FailedProtectionGroups) {
		failedGroups := []map[string]interface{}{}
		for _, failedGroup := range updateProtectionGroupsState.FailedProtectionGroups {
			failedGroupsMap, err := resourceIbmProtectionGroupStateRunRequestMapToProtectionGroupStateFailedGroups(&failedGroup)
			if err != nil {
				return diag.FromErr(err)
			}
			failedGroups = append(failedGroups, failedGroupsMap)
		}
		if err = d.Set("failed_groups", failedGroups); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting failedRuns: %s", err))
		}
	}

	if err = d.Set("action", updateProtectionGroupsStateOptions.Action); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting action: %s", err))
	}
	if err = d.Set("ids", updateProtectionGroupsStateOptions.Ids); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting ids: %s", err))
	}

	d.SetId(resourceIbmProtectionGroupStateID(d))

	return resourceIbmProtectionGroupStateRead(context, d, meta)
}

func resourceIbmProtectionGroupStateRunRequestMapToProtectionGroupStateFailedGroups(model *backuprecoveryv1.FailedProtectionGroup) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ProtectionGroupID != nil {
		modelMap["protection_group_id"] = model.ProtectionGroupID
	}
	if model.ErrorMessage != nil {
		modelMap["error_message"] = model.ErrorMessage
	}
	return modelMap, nil
}

func resourceIbmProtectionGroupStateID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func resourceIbmProtectionGroupStateRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return nil
}

func resourceIbmProtectionGroupStateDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// This resource does not support a "delete" operation.

	var diags diag.Diagnostics
	warning := diag.Diagnostic{
		Severity: diag.Warning,
		Summary:  "Delete Not Supported",
		Detail:   "Delete operation is not supported for this resource. The resource will be removed from the terraform file but will continue to exist in the backend.",
	}
	diags = append(diags, warning)
	d.SetId("")
	return diags
}

func resourceIbmProtectionGroupStateUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// This resource does not support a "delete" operation.
	var diags diag.Diagnostics
	warning := diag.Diagnostic{
		Severity: diag.Warning,
		Summary:  "Update Not Supported",
		Detail:   "Update operation is not supported for this resource. No changes will be applied.",
	}
	diags = append(diags, warning)
	d.SetId("")
	return diags
}
