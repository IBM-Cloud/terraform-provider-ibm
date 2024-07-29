// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package backuprecovery

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.ibm.com/BackupAndRecovery/ibm-backup-recovery-sdk-go/backuprecoveryv1"
)

func ResourceIbmPerformActionOnProtectionGroupRunRequest() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIbmPerformActionOnProtectionGroupRunRequestCreate,
		ReadContext:   resourceIbmPerformActionOnProtectionGroupRunRequestRead,
		DeleteContext: resourceIbmPerformActionOnProtectionGroupRunRequestDelete,
		UpdateContext: resourceIbmPerformActionOnProtectionGroupRunRequestUpdate,
		CustomizeDiff: checkResourceIbmPerformActionOnProtectionGroupRunDiff,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"action": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.InvokeValidator("ibm_perform_action_on_protection_group_run_request", "action"),
				Description:  "Specifies the type of the action which will be performed on protection runs.",
			},
			"pause_params": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: "Specifies the pause action params for a protection run.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"run_id": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "Specifies a unique run id of the Protection Group run.",
						},
					},
				},
			},
			"resume_params": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: "Specifies the resume action params for a protection run.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"run_id": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "Specifies a unique run id of the Protection Group run.",
						},
					},
				},
			},
			"cancel_params": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: "Specifies the cancel action params for a protection run.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"run_id": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "Specifies a unique run id of the Protection Group run.",
						},
						"local_task_id": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Specifies the task id of the local run.",
						},
						"object_ids": &schema.Schema{
							Type:        schema.TypeList,
							Optional:    true,
							Description: "List of entity ids for which we need to cancel the backup tasks. If this is provided it will not cancel the complete run but will cancel only subset of backup tasks (if backup tasks are cancelled correspoding copy task will also get cancelled). If the backup tasks are completed successfully it will not cancel those backup tasks.",
							Elem:        &schema.Schema{Type: schema.TypeInt},
						},
						"replication_task_id": &schema.Schema{
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Specifies the task id of the replication run.",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"archival_task_id": &schema.Schema{
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Specifies the task id of the archival run.",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"cloud_spin_task_id": &schema.Schema{
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Specifies the task id of the cloudSpin run.",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
					},
				},
			},
			"run_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The unique ID.",
			},
			"group_id": {
				Type:     schema.TypeString,
				Required: true,
				// ValidateFunc: validate.InvokeValidator("ibm_create_protection_group_run_request", "run_type"),
				Description: "Protection group id",
			},
		},
	}
}

func checkResourceIbmPerformActionOnProtectionGroupRunDiff(context context.Context, d *schema.ResourceDiff, meta interface{}) error {
	// skip if it's a new resource
	if d.Id() == "" {
		return nil
	}

	resourceSchema := ResourceIbmPerformActionOnProtectionGroupRunRequest().Schema

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

func ResourceIbmPerformActionOnProtectionGroupRunRequestValidator() *validate.ResourceValidator {
	validateSchema := make([]validate.ValidateSchema, 0)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "action",
			ValidateFunctionIdentifier: validate.ValidateAllowedStringValue,
			Type:                       validate.TypeString,
			Required:                   true,
			AllowedValues:              "Cancel, Pause, Resume",
		},
	)

	resourceValidator := validate.ResourceValidator{ResourceName: "ibm_perform_action_on_protection_group_run_request", Schema: validateSchema}
	return &resourceValidator
}

func resourceIbmPerformActionOnProtectionGroupRunRequestCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	backupRecoveryClient, err := meta.(conns.ClientSession).BackupRecoveryV1()
	if err != nil {
		return diag.FromErr(err)
	}

	performActionOnProtectionGroupRunOptions := &backuprecoveryv1.PerformActionOnProtectionGroupRunOptions{}
	performActionOnProtectionGroupRunOptions.SetID(d.Get("group_id").(string))
	performActionOnProtectionGroupRunOptions.SetAction(d.Get("action").(string))
	if _, ok := d.GetOk("pause_params"); ok {
		var newPauseParams []backuprecoveryv1.PauseProtectionRunActionParams
		for _, v := range d.Get("pause_params").([]interface{}) {
			value := v.(map[string]interface{})
			newPauseParamsItem, err := resourceIbmPerformActionOnProtectionGroupRunRequestMapToPauseProtectionRunActionParams(value)
			if err != nil {
				return diag.FromErr(err)
			}
			newPauseParams = append(newPauseParams, *newPauseParamsItem)
		}
		performActionOnProtectionGroupRunOptions.SetPauseParams(newPauseParams)
	}
	if _, ok := d.GetOk("resume_params"); ok {
		var newResumeParams []backuprecoveryv1.ResumeProtectionRunActionParams
		for _, v := range d.Get("resume_params").([]interface{}) {
			value := v.(map[string]interface{})
			newResumeParamsItem, err := resourceIbmPerformActionOnProtectionGroupRunRequestMapToResumeProtectionRunActionParams(value)
			if err != nil {
				return diag.FromErr(err)
			}
			newResumeParams = append(newResumeParams, *newResumeParamsItem)
		}
		performActionOnProtectionGroupRunOptions.SetResumeParams(newResumeParams)
	}
	if _, ok := d.GetOk("cancel_params"); ok {
		var newCancelParams []backuprecoveryv1.CancelProtectionGroupRunRequest
		for _, v := range d.Get("cancel_params").([]interface{}) {
			value := v.(map[string]interface{})
			newCancelParamsItem, err := resourceIbmPerformActionOnProtectionGroupRunRequestMapToCancelProtectionGroupRunRequest(value)
			if err != nil {
				return diag.FromErr(err)
			}
			newCancelParams = append(newCancelParams, *newCancelParamsItem)
		}
		performActionOnProtectionGroupRunOptions.SetCancelParams(newCancelParams)
	}

	performRunActionResponse, response, err := backupRecoveryClient.PerformActionOnProtectionGroupRunWithContext(context, performActionOnProtectionGroupRunOptions)
	if err != nil {
		log.Printf("[DEBUG] PerformActionOnProtectionGroupRunWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("PerformActionOnProtectionGroupRunWithContext failed %s\n%s", err, response))
	}

	d.SetId(resourceIbmProtectionRunActionID(d))

	d.Set("action", performRunActionResponse.Action)

	if !core.IsNil(performRunActionResponse.PauseParams) {
		pauseParams := []map[string]interface{}{}
		for _, pauseParamsItem := range performRunActionResponse.PauseParams {
			pauseParamsItemMap, err := resourceIbmPerformActionOnProtectionGroupRunRequestPauseProtectionRunActionParamsToMap(&pauseParamsItem)
			if err != nil {
				return diag.FromErr(err)
			}
			pauseParams = append(pauseParams, pauseParamsItemMap)
		}
		_ = d.Set("pause_params", []map[string]interface{}{})
		if err = d.Set("pause_params", pauseParams); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting pause_params: %s", err))
		}
	}
	if !core.IsNil(performRunActionResponse.ResumeParams) {
		resumeParams := []map[string]interface{}{}
		for _, resumeParamsItem := range performRunActionResponse.ResumeParams {
			resumeParamsItemMap, err := resourceIbmPerformActionOnProtectionGroupRunRequestResumeProtectionRunActionParamsToMap(&resumeParamsItem)
			if err != nil {
				return diag.FromErr(err)
			}
			resumeParams = append(resumeParams, resumeParamsItemMap)
		}
		_ = d.Set("resume_params", []map[string]interface{}{})
		if err = d.Set("resume_params", resumeParams); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting resume_params: %s", err))
		}
	}
	if !core.IsNil(performRunActionResponse.CancelParams) {
		cancelParams := []map[string]interface{}{}
		for _, cancelParamsItem := range performRunActionResponse.CancelParams {
			cancelParamsItemMap, err := resourceIbmPerformActionOnProtectionGroupRunRequestCancelProtectionGroupRunRequestToMap(&cancelParamsItem)
			if err != nil {
				return diag.FromErr(err)
			}
			cancelParams = append(cancelParams, cancelParamsItemMap)
		}

		if err = d.Set("cancel_params", cancelParams); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting cancel_params: %s", err))
		}
	}
	return resourceIbmPerformActionOnProtectionGroupRunRequestRead(context, d, meta)
}

func resourceIbmProtectionRunActionID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func resourceIbmPerformActionOnProtectionGroupRunRequestRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return nil
}

func resourceIbmPerformActionOnProtectionGroupRunRequestDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

func resourceIbmPerformActionOnProtectionGroupRunRequestUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

func resourceIbmPerformActionOnProtectionGroupRunRequestMapToPauseProtectionRunActionParams(modelMap map[string]interface{}) (*backuprecoveryv1.PauseProtectionRunActionParams, error) {
	model := &backuprecoveryv1.PauseProtectionRunActionParams{}
	model.RunID = core.StringPtr(modelMap["run_id"].(string))
	return model, nil
}

func resourceIbmPerformActionOnProtectionGroupRunRequestMapToResumeProtectionRunActionParams(modelMap map[string]interface{}) (*backuprecoveryv1.ResumeProtectionRunActionParams, error) {
	model := &backuprecoveryv1.ResumeProtectionRunActionParams{}
	model.RunID = core.StringPtr(modelMap["run_id"].(string))
	return model, nil
}

func resourceIbmPerformActionOnProtectionGroupRunRequestMapToCancelProtectionGroupRunRequest(modelMap map[string]interface{}) (*backuprecoveryv1.CancelProtectionGroupRunRequest, error) {
	model := &backuprecoveryv1.CancelProtectionGroupRunRequest{}
	model.RunID = core.StringPtr(modelMap["run_id"].(string))
	if modelMap["local_task_id"] != nil && modelMap["local_task_id"].(string) != "" {
		model.LocalTaskID = core.StringPtr(modelMap["local_task_id"].(string))
	}
	if modelMap["object_ids"] != nil {
		objectIds := []int64{}
		for _, objectIdsItem := range modelMap["object_ids"].([]interface{}) {
			objectIds = append(objectIds, int64(objectIdsItem.(int)))
		}
		model.ObjectIds = objectIds
	}
	if modelMap["replication_task_id"] != nil {
		replicationTaskID := []string{}
		for _, replicationTaskIDItem := range modelMap["replication_task_id"].([]interface{}) {
			replicationTaskID = append(replicationTaskID, replicationTaskIDItem.(string))
		}
		model.ReplicationTaskID = replicationTaskID
	}
	if modelMap["archival_task_id"] != nil {
		archivalTaskID := []string{}
		for _, archivalTaskIDItem := range modelMap["archival_task_id"].([]interface{}) {
			archivalTaskID = append(archivalTaskID, archivalTaskIDItem.(string))
		}
		model.ArchivalTaskID = archivalTaskID
	}
	if modelMap["cloud_spin_task_id"] != nil {
		cloudSpinTaskID := []string{}
		for _, cloudSpinTaskIDItem := range modelMap["cloud_spin_task_id"].([]interface{}) {
			cloudSpinTaskID = append(cloudSpinTaskID, cloudSpinTaskIDItem.(string))
		}
		model.CloudSpinTaskID = cloudSpinTaskID
	}
	return model, nil
}

func resourceIbmPerformActionOnProtectionGroupRunRequestPauseProtectionRunActionParamsToMap(model *backuprecoveryv1.PauseProtectionRunActionResponseParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["run_id"] = model.RunID
	return modelMap, nil
}

func resourceIbmPerformActionOnProtectionGroupRunRequestResumeProtectionRunActionParamsToMap(model *backuprecoveryv1.ResumeProtectionRunActionResponseParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["run_id"] = model.RunID
	return modelMap, nil
}

func resourceIbmPerformActionOnProtectionGroupRunRequestCancelProtectionGroupRunRequestToMap(model *backuprecoveryv1.CancelProtectionGroupRunResponseParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["run_id"] = model.RunID
	return modelMap, nil
}
