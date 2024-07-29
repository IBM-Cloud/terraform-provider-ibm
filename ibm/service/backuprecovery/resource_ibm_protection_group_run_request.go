// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package backuprecovery

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.ibm.com/BackupAndRecovery/ibm-backup-recovery-sdk-go/backuprecoveryv1"
)

func ResourceIbmProtectionGroupRunRequest() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIbmProtectionGroupRunRequestCreate,
		ReadContext:   resourceIbmProtectionGroupRunRequestRead,
		DeleteContext: resourceIbmProtectionGroupRunRequestDelete,
		UpdateContext: resourceIbmProtectionGroupRunRequestUpdate,
		Importer:      &schema.ResourceImporter{},
		CustomizeDiff: checkResourceIbmProtectionGroupRunDiff,
		Schema: map[string]*schema.Schema{
			"group_id": {
				Type:     schema.TypeString,
				Required: true,
				// ValidateFunc: validate.InvokeValidator("ibm_create_protection_group_run_request", "run_type"),
				Description: "Protection group id",
			},
			"run_type": {
				Type:     schema.TypeString,
				Required: true,
				//ForceNew: true,
				//ValidateFunc: validate.InvokeValidator("ibm_protection_group_run_request", "run_type"),
				Description: "Type of protection run. 'kRegular' indicates an incremental (CBT) backup. Incremental backups utilizing CBT (if supported) are captured of the target protection objects. The first run of a kRegular schedule captures all the blocks. 'kFull' indicates a full (no CBT) backup. A complete backup (all blocks) of the target protection objects are always captured and Change Block Tracking (CBT) is not utilized. 'kLog' indicates a Database Log backup. Capture the database transaction logs to allow rolling back to a specific point in time. 'kSystem' indicates system volume backup. It produces an image for bare metal recovery.",
			},
			"objects": {
				Type:     schema.TypeList,
				Optional: true,
				//ForceNew:    true,
				Description: "Specifies the list of objects to be protected by this Protection Group run. These can be leaf objects or non-leaf objects in the protection hierarchy. This must be specified only if a subset of objects from the Protection Groups needs to be protected.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeInt,
							Required:    true,
							Description: "Specifies the id of object.",
						},
						"app_ids": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Specifies a list of ids of applications.",
							Elem:        &schema.Schema{Type: schema.TypeInt},
						},
						"physical_params": {
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "Specifies physical parameters for this run.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"metadata_file_path": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Specifies metadata file path during run-now requests for physical file based backups for some specific source. If specified, it will override any default metadata/directive file path set at the object level for the source. Also note that if the job default does not specify a metadata/directive file path for the source, then specifying this field for that source during run-now request will be rejected.",
									},
								},
							},
						},
					},
				},
			},
			"targets_config": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				//ForceNew:    true,
				Description: "Specifies the replication and archival targets.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"use_policy_defaults": {
							Type:        schema.TypeBool,
							Optional:    true,
							Default:     false,
							Description: "Specifies whether to use default policy settings or not. If specified as true then 'replications' and 'arcihvals' should not be specified. In case of true value, replicatioan targets congfigured in the policy will be added internally.",
						},
						"replications": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Specifies a list of replication targets configurations.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:        schema.TypeInt,
										Required:    true,
										Description: "Specifies id of Remote Cluster to copy the Snapshots to.",
									},
									"retention": {
										Type:        schema.TypeList,
										MaxItems:    1,
										Optional:    true,
										Description: "Specifies the retention of a backup.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"unit": {
													Type:        schema.TypeString,
													Required:    true,
													Description: "Specificies the Retention Unit of a backup measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Years' then number of retention days will be 365 * 2 = 730 days.",
												},
												"duration": {
													Type:        schema.TypeInt,
													Required:    true,
													Description: "Specifies the duration for a backup retention. <br> Example. If duration is 7 and unit is Months, the retention of a backup is 7 * 30 = 210 days.",
												},
												"data_lock_config": {
													Type:        schema.TypeList,
													MaxItems:    1,
													Optional:    true,
													Description: "Specifies WORM retention type for the snapshots. When a WORM retention type is specified, the snapshots of the Protection Groups using this policy will be kept for the last N days as specified in the duration of the datalock. During that time, the snapshots cannot be deleted.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"mode": {
																Type:        schema.TypeString,
																Required:    true,
																Description: "Specifies the type of WORM retention type. 'Compliance' implies WORM retention is set for compliance reason. 'Administrative' implies WORM retention is set for administrative purposes.",
															},
															"unit": {
																Type:        schema.TypeString,
																Required:    true,
																Description: "Specificies the Retention Unit of a dataLock measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Months' then number of retention days will be 365 * 2 = 730 days.",
															},
															"duration": {
																Type:        schema.TypeInt,
																Required:    true,
																Description: "Specifies the duration for a dataLock. <br> Example. If duration is 7 and unit is Months, the dataLock is enabled for last 7 * 30 = 210 days of the backup.",
															},
															"enable_worm_on_external_target": {
																Type:        schema.TypeBool,
																Optional:    true,
																Description: "Specifies whether objects in the external target associated with this policy need to be made immutable.",
															},
														},
													},
												},
											},
										},
									},
								},
							},
						},
						"archivals": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Specifies a list of archival targets configurations.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": {
										Type:        schema.TypeInt,
										Required:    true,
										Description: "Specifies the Archival target to copy the Snapshots to.",
									},
									"archival_target_type": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "Specifies the snapshot's archival target type from which recovery has been performed.",
									},
									"retention": {
										Type:        schema.TypeList,
										MaxItems:    1,
										Optional:    true,
										Description: "Specifies the retention of a backup.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"unit": {
													Type:        schema.TypeString,
													Required:    true,
													Description: "Specificies the Retention Unit of a backup measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Years' then number of retention days will be 365 * 2 = 730 days.",
												},
												"duration": {
													Type:        schema.TypeInt,
													Required:    true,
													Description: "Specifies the duration for a backup retention. <br> Example. If duration is 7 and unit is Months, the retention of a backup is 7 * 30 = 210 days.",
												},
												"data_lock_config": {
													Type:        schema.TypeList,
													MaxItems:    1,
													Optional:    true,
													Description: "Specifies WORM retention type for the snapshots. When a WORM retention type is specified, the snapshots of the Protection Groups using this policy will be kept for the last N days as specified in the duration of the datalock. During that time, the snapshots cannot be deleted.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"mode": {
																Type:        schema.TypeString,
																Required:    true,
																Description: "Specifies the type of WORM retention type. 'Compliance' implies WORM retention is set for compliance reason. 'Administrative' implies WORM retention is set for administrative purposes.",
															},
															"unit": {
																Type:        schema.TypeString,
																Required:    true,
																Description: "Specificies the Retention Unit of a dataLock measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Months' then number of retention days will be 365 * 2 = 730 days.",
															},
															"duration": {
																Type:        schema.TypeInt,
																Required:    true,
																Description: "Specifies the duration for a dataLock. <br> Example. If duration is 7 and unit is Months, the dataLock is enabled for last 7 * 30 = 210 days of the backup.",
															},
															"enable_worm_on_external_target": {
																Type:        schema.TypeBool,
																Optional:    true,
																Description: "Specifies whether objects in the external target associated with this policy need to be made immutable.",
															},
														},
													},
												},
											},
										},
									},
									"copy_only_fully_successful": {
										Type:        schema.TypeBool,
										Optional:    true,
										Description: "Specifies if Snapshots are copied from a fully successful Protection Group Run or a partially successful Protection Group Run. If false, Snapshots are copied the Protection Group Run, even if the Run was not fully successful i.e. Snapshots were not captured for all Objects in the Protection Group. If true, Snapshots are copied only when the run is fully successful.",
									},
								},
							},
						},
						"cloud_replications": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Specifies a list of cloud replication targets configurations.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"target_type": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "Specifies the type of target to which replication need to be performed.",
									},
									"retention": {
										Type:        schema.TypeList,
										MaxItems:    1,
										Optional:    true,
										Description: "Specifies the retention of a backup.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"unit": {
													Type:        schema.TypeString,
													Required:    true,
													Description: "Specificies the Retention Unit of a backup measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Years' then number of retention days will be 365 * 2 = 730 days.",
												},
												"duration": {
													Type:        schema.TypeInt,
													Required:    true,
													Description: "Specifies the duration for a backup retention. <br> Example. If duration is 7 and unit is Months, the retention of a backup is 7 * 30 = 210 days.",
												},
												"data_lock_config": {
													Type:        schema.TypeList,
													MaxItems:    1,
													Optional:    true,
													Description: "Specifies WORM retention type for the snapshots. When a WORM retention type is specified, the snapshots of the Protection Groups using this policy will be kept for the last N days as specified in the duration of the datalock. During that time, the snapshots cannot be deleted.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"mode": {
																Type:        schema.TypeString,
																Required:    true,
																Description: "Specifies the type of WORM retention type. 'Compliance' implies WORM retention is set for compliance reason. 'Administrative' implies WORM retention is set for administrative purposes.",
															},
															"unit": {
																Type:        schema.TypeString,
																Required:    true,
																Description: "Specificies the Retention Unit of a dataLock measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Months' then number of retention days will be 365 * 2 = 730 days.",
															},
															"duration": {
																Type:        schema.TypeInt,
																Required:    true,
																Description: "Specifies the duration for a dataLock. <br> Example. If duration is 7 and unit is Months, the dataLock is enabled for last 7 * 30 = 210 days of the backup.",
															},
															"enable_worm_on_external_target": {
																Type:        schema.TypeBool,
																Optional:    true,
																Description: "Specifies whether objects in the external target associated with this policy need to be made immutable.",
															},
														},
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			"uda_params": {
				Type:     schema.TypeList,
				MaxItems: 1,
				Optional: true,
				//ForceNew:    true,
				Description: "Specifies the parameters for Universal Data Adapter protection run.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"externally_triggered_run_params": {
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "Specifies the parameters for an externally triggered run.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"control_node": {
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Specifies the IP or FQDN of the source host where this backup will run.",
									},
									"backup_args": {
										Type:        schema.TypeList,
										Optional:    true,
										Description: "Specifies a map of custom arguments to be supplied to the plugin.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"key": {
													Type:        schema.TypeString,
													Required:    true,
													Description: "key.",
												},
												"value": {
													Type:        schema.TypeString,
													Required:    true,
													Description: "value.",
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			"run_id": {
				Type: schema.TypeString,
				// Computed:    true,
				Optional:    true,
				Description: "The unique ID.",
			},
		},
	}
}

func checkResourceIbmProtectionGroupRunDiff(context context.Context, d *schema.ResourceDiff, meta interface{}) error {
	// skip if it's a new resource
	if d.Id() == "" {
		return nil
	}

	resourceSchema := ResourceIbmProtectionGroupRunRequest().Schema

	// handle update resource
LOOP:
	for key := range resourceSchema {
		if d.HasChange(key) {
			log.Println("[WARNING] Update operation is not supported for this resource. No changes will be applied. Please use ibm_update_protection_group_run_request resource for updates.")
			break LOOP
		}
	}

	return nil
}

func ResourceIbmProtectionGroupRunRequestValidator() *validate.ResourceValidator {
	validateSchema := make([]validate.ValidateSchema, 0)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "run_type",
			ValidateFunctionIdentifier: validate.ValidateAllowedStringValue,
			Type:                       validate.TypeString,
			Required:                   true,
			AllowedValues:              "kFull, kHydrateCDP, kLog, kRegular, kStorageArraySnapshot, kSystem",
		},
		validate.ValidateSchema{
			Identifier:                 "group_id",
			ValidateFunctionIdentifier: validate.ValidateAllowedStringValue,
			Type:                       validate.TypeString,
			Required:                   true,
		},
	)

	resourceValidator := validate.ResourceValidator{ResourceName: "ibm_protection_group_run_request", Schema: validateSchema}
	return &resourceValidator
}

func resourceIbmProtectionGroupRunRequestCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	backupRecoveryClient, err := meta.(conns.ClientSession).BackupRecoveryV1()
	if err != nil {
		return diag.FromErr(err)
	}

	createProtectionGroupRunOptions := &backuprecoveryv1.CreateProtectionGroupRunOptions{}

	createProtectionGroupRunOptions.SetID(d.Get("group_id").(string))

	createProtectionGroupRunOptions.SetRunType(d.Get("run_type").(string))
	if _, ok := d.GetOk("objects"); ok {
		var newObjects []backuprecoveryv1.RunObject
		for _, v := range d.Get("objects").([]interface{}) {
			value := v.(map[string]interface{})
			newObjectsItem, err := resourceIbmProtectionGroupRunRequestMapToRunObject(value)
			if err != nil {
				return diag.FromErr(err)
			}
			newObjects = append(newObjects, *newObjectsItem)
		}
		createProtectionGroupRunOptions.SetObjects(newObjects)
	}
	if _, ok := d.GetOk("targets_config"); ok {
		newTargetsConfigModel, err := resourceIbmProtectionGroupRunRequestMapToRunTargetsConfiguration(d.Get("targets_config.0").(map[string]interface{}))
		if err != nil {
			return diag.FromErr(err)
		}
		createProtectionGroupRunOptions.SetTargetsConfig(newTargetsConfigModel)
	}
	if _, ok := d.GetOk("uda_params"); ok {
		newUdaParamsModel, err := resourceIbmProtectionGroupRunRequestMapToUdaProtectionRunParams(d.Get("uda_params.0").(map[string]interface{}))
		if err != nil {
			return diag.FromErr(err)
		}
		createProtectionGroupRunOptions.SetUdaParams(newUdaParamsModel)
	}

	createProtectionGroupRunResponseBody, response, err := backupRecoveryClient.CreateProtectionGroupRunWithContext(context, createProtectionGroupRunOptions)
	if err != nil {
		log.Printf("[DEBUG] CreateProtectionGroupRunWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("CreateProtectionGroupRunWithContext failed %s\n%s", err, response))
	}

	d.SetId(*createProtectionGroupRunResponseBody.ProtectionGroupID)

	if !core.IsNil(createProtectionGroupRunResponseBody.UdaParams) {
		udaParamsMap, err := resourceIbmProtectionGroupRunRequestUdaProtectionRunParamsToMap(createProtectionGroupRunResponseBody.UdaParams)
		if err != nil {
			return diag.FromErr(err)
		}
		if err = d.Set("uda_params", []map[string]interface{}{udaParamsMap}); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting uda_params: %s", err))
		}
	}

	return nil
	// return resourceIbmCreateProtectionGroupRunsRequestRead(context, d, meta)
}

func resourceIbmProtectionGroupRunRequestUdaProtectionRunParamsToMap(model *backuprecoveryv1.UdaCreateRunResponseParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ExternallyTriggeredRunID != nil {
		modelMap["externally_triggered_run_id"] = model.ExternallyTriggeredRunID
	}
	return modelMap, nil
}

func resourceIbmProtectionGroupRunRequestUdaExternallyTriggeredRunParamsToMap(model *backuprecoveryv1.UdaExternallyTriggeredRunParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ControlNode != nil {
		modelMap["control_node"] = model.ControlNode
	}
	if model.BackupArgs != nil {
		backupArgs := []map[string]interface{}{}
		for _, backupArgsItem := range model.BackupArgs {
			backupArgsItemMap, err := resourceIbmProtectionGroupRunRequestKeyValuePairToMap(&backupArgsItem)
			if err != nil {
				return modelMap, err
			}
			backupArgs = append(backupArgs, backupArgsItemMap)
		}
		modelMap["backup_args"] = backupArgs
	}
	return modelMap, nil
}

func resourceIbmProtectionGroupRunRequestRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {

	return nil
}

func resourceIbmProtectionGroupRunRequestDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
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

func resourceIbmProtectionGroupRunRequestUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// This resource does not support a "delete" operation.
	var diags diag.Diagnostics
	warning := diag.Diagnostic{
		Severity: diag.Warning,
		Summary:  "Update Not Supported",
		Detail:   "Update operation is not supported for this resource. No changes will be applied. Please use ibm_update_protection_group_run_request resource for updates.",
	}
	diags = append(diags, warning)
	d.SetId("")
	return diags
}

func resourceIbmProtectionGroupRunRequestMapToRunObject(modelMap map[string]interface{}) (*backuprecoveryv1.RunObject, error) {
	model := &backuprecoveryv1.RunObject{}
	model.ID = core.Int64Ptr(int64(modelMap["id"].(int)))
	if modelMap["app_ids"] != nil {
		appIds := []int64{}
		for _, appIdsItem := range modelMap["app_ids"].([]interface{}) {
			appIds = append(appIds, int64(appIdsItem.(int)))
		}
		model.AppIds = appIds
	}
	if modelMap["physical_params"] != nil && len(modelMap["physical_params"].([]interface{})) > 0 {
		PhysicalParamsModel, err := resourceIbmProtectionGroupRunRequestMapToRunObjectPhysicalParams(modelMap["physical_params"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.PhysicalParams = PhysicalParamsModel
	}
	return model, nil
}

func resourceIbmProtectionGroupRunRequestMapToRunObjectPhysicalParams(modelMap map[string]interface{}) (*backuprecoveryv1.RunObjectPhysicalParams, error) {
	model := &backuprecoveryv1.RunObjectPhysicalParams{}
	if modelMap["metadata_file_path"] != nil && modelMap["metadata_file_path"].(string) != "" {
		model.MetadataFilePath = core.StringPtr(modelMap["metadata_file_path"].(string))
	}
	return model, nil
}

func resourceIbmProtectionGroupRunRequestMapToRunTargetsConfiguration(modelMap map[string]interface{}) (*backuprecoveryv1.RunTargetsConfiguration, error) {
	model := &backuprecoveryv1.RunTargetsConfiguration{}
	if modelMap["use_policy_defaults"] != nil {
		model.UsePolicyDefaults = core.BoolPtr(modelMap["use_policy_defaults"].(bool))
	}
	if modelMap["replications"] != nil {
		replications := []backuprecoveryv1.RunReplicationConfig{}
		for _, replicationsItem := range modelMap["replications"].([]interface{}) {
			replicationsItemModel, err := resourceIbmProtectionGroupRunRequestMapToRunReplicationConfig(replicationsItem.(map[string]interface{}))
			if err != nil {
				return model, err
			}
			replications = append(replications, *replicationsItemModel)
		}
		model.Replications = replications
	}
	if modelMap["archivals"] != nil {
		archivals := []backuprecoveryv1.RunArchivalConfig{}
		for _, archivalsItem := range modelMap["archivals"].([]interface{}) {
			archivalsItemModel, err := resourceIbmProtectionGroupRunRequestMapToRunArchivalConfig(archivalsItem.(map[string]interface{}))
			if err != nil {
				return model, err
			}
			archivals = append(archivals, *archivalsItemModel)
		}
		model.Archivals = archivals
	}
	if modelMap["cloud_replications"] != nil {
		cloudReplications := []backuprecoveryv1.RunCloudReplicationConfig{}
		for _, cloudReplicationsItem := range modelMap["cloud_replications"].([]interface{}) {
			cloudReplicationsItemModel, err := resourceIbmProtectionGroupRunRequestMapToRunCloudReplicationConfig(cloudReplicationsItem.(map[string]interface{}))
			if err != nil {
				return model, err
			}
			cloudReplications = append(cloudReplications, *cloudReplicationsItemModel)
		}
		model.CloudReplications = cloudReplications
	}
	return model, nil
}

func resourceIbmProtectionGroupRunRequestMapToRunReplicationConfig(modelMap map[string]interface{}) (*backuprecoveryv1.RunReplicationConfig, error) {
	model := &backuprecoveryv1.RunReplicationConfig{}
	model.ID = core.Int64Ptr(int64(modelMap["id"].(int)))
	if modelMap["retention"] != nil && len(modelMap["retention"].([]interface{})) > 0 {
		RetentionModel, err := resourceIbmProtectionGroupRunRequestMapToRetention(modelMap["retention"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.Retention = RetentionModel
	}
	return model, nil
}

func resourceIbmProtectionGroupRunRequestMapToRetention(modelMap map[string]interface{}) (*backuprecoveryv1.Retention, error) {
	model := &backuprecoveryv1.Retention{}
	model.Unit = core.StringPtr(modelMap["unit"].(string))
	model.Duration = core.Int64Ptr(int64(modelMap["duration"].(int)))
	if modelMap["data_lock_config"] != nil && len(modelMap["data_lock_config"].([]interface{})) > 0 {
		DataLockConfigModel, err := resourceIbmProtectionGroupRunRequestMapToDataLockConfig(modelMap["data_lock_config"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.DataLockConfig = DataLockConfigModel
	}
	return model, nil
}

func resourceIbmProtectionGroupRunRequestMapToDataLockConfig(modelMap map[string]interface{}) (*backuprecoveryv1.DataLockConfig, error) {
	model := &backuprecoveryv1.DataLockConfig{}
	model.Mode = core.StringPtr(modelMap["mode"].(string))
	model.Unit = core.StringPtr(modelMap["unit"].(string))
	model.Duration = core.Int64Ptr(int64(modelMap["duration"].(int)))
	if modelMap["enable_worm_on_external_target"] != nil {
		model.EnableWormOnExternalTarget = core.BoolPtr(modelMap["enable_worm_on_external_target"].(bool))
	}
	return model, nil
}

func resourceIbmProtectionGroupRunRequestMapToRunArchivalConfig(modelMap map[string]interface{}) (*backuprecoveryv1.RunArchivalConfig, error) {
	model := &backuprecoveryv1.RunArchivalConfig{}
	model.ID = core.Int64Ptr(int64(modelMap["id"].(int)))
	model.ArchivalTargetType = core.StringPtr(modelMap["archival_target_type"].(string))
	if modelMap["retention"] != nil && len(modelMap["retention"].([]interface{})) > 0 {
		RetentionModel, err := resourceIbmProtectionGroupRunRequestMapToRetention(modelMap["retention"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.Retention = RetentionModel
	}
	if modelMap["copy_only_fully_successful"] != nil {
		model.CopyOnlyFullySuccessful = core.BoolPtr(modelMap["copy_only_fully_successful"].(bool))
	}
	return model, nil
}

func resourceIbmProtectionGroupRunRequestMapToRunCloudReplicationConfig(modelMap map[string]interface{}) (*backuprecoveryv1.RunCloudReplicationConfig, error) {
	model := &backuprecoveryv1.RunCloudReplicationConfig{}
	model.TargetType = core.StringPtr(modelMap["target_type"].(string))
	if modelMap["retention"] != nil && len(modelMap["retention"].([]interface{})) > 0 {
		RetentionModel, err := resourceIbmProtectionGroupRunRequestMapToRetention(modelMap["retention"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.Retention = RetentionModel
	}
	return model, nil
}

func resourceIbmProtectionGroupRunRequestMapToUdaProtectionRunParams(modelMap map[string]interface{}) (*backuprecoveryv1.UdaProtectionRunParams, error) {
	model := &backuprecoveryv1.UdaProtectionRunParams{}
	if modelMap["externally_triggered_run_params"] != nil && len(modelMap["externally_triggered_run_params"].([]interface{})) > 0 {
		ExternallyTriggeredRunParamsModel, err := resourceIbmProtectionGroupRunRequestMapToUdaExternallyTriggeredRunParams(modelMap["externally_triggered_run_params"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.ExternallyTriggeredRunParams = ExternallyTriggeredRunParamsModel
	}
	return model, nil
}

func resourceIbmProtectionGroupRunRequestMapToUdaExternallyTriggeredRunParams(modelMap map[string]interface{}) (*backuprecoveryv1.UdaExternallyTriggeredRunParams, error) {
	model := &backuprecoveryv1.UdaExternallyTriggeredRunParams{}
	if modelMap["control_node"] != nil && modelMap["control_node"].(string) != "" {
		model.ControlNode = core.StringPtr(modelMap["control_node"].(string))
	}
	if modelMap["backup_args"] != nil {
		backupArgs := []backuprecoveryv1.KeyValuePair{}
		for _, backupArgsItem := range modelMap["backup_args"].([]interface{}) {
			backupArgsItemModel, err := resourceIbmProtectionGroupRunRequestMapToKeyValuePair(backupArgsItem.(map[string]interface{}))
			if err != nil {
				return model, err
			}
			backupArgs = append(backupArgs, *backupArgsItemModel)
		}
		model.BackupArgs = backupArgs
	}
	return model, nil
}

func resourceIbmProtectionGroupRunRequestMapToKeyValuePair(modelMap map[string]interface{}) (*backuprecoveryv1.KeyValuePair, error) {
	model := &backuprecoveryv1.KeyValuePair{}
	model.Key = core.StringPtr(modelMap["key"].(string))
	model.Value = core.StringPtr(modelMap["value"].(string))
	return model, nil
}

//--

func resourceIbmProtectionGroupRunRequestRunCloudReplicationConfigToMap(model *backuprecoveryv1.RunCloudReplicationConfig) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["target_type"] = model.TargetType
	if model.Retention != nil {
		retentionMap, err := resourceIbmProtectionGroupRunRequestRetentionToMap(model.Retention)
		if err != nil {
			return modelMap, err
		}
		modelMap["retention"] = []map[string]interface{}{retentionMap}
	}
	return modelMap, nil
}

func resourceIbmProtectionGroupRunRequestRunObjectToMap(model *backuprecoveryv1.RunObject) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["id"] = flex.IntValue(model.ID)
	if model.AppIds != nil {
		modelMap["app_ids"] = model.AppIds
	}
	if model.PhysicalParams != nil {
		physicalParamsMap, err := resourceIbmProtectionGroupRunRequestRunObjectPhysicalParamsToMap(model.PhysicalParams)
		if err != nil {
			return modelMap, err
		}
		modelMap["physical_params"] = []map[string]interface{}{physicalParamsMap}
	}
	return modelMap, nil
}

func resourceIbmProtectionGroupRunRequestRunObjectPhysicalParamsToMap(model *backuprecoveryv1.RunObjectPhysicalParams) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.MetadataFilePath != nil {
		modelMap["metadata_file_path"] = model.MetadataFilePath
	}
	return modelMap, nil
}

func resourceIbmProtectionGroupRunRequestRunTargetsConfigurationToMap(model *backuprecoveryv1.RunTargetsConfiguration) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.UsePolicyDefaults != nil {
		modelMap["use_policy_defaults"] = model.UsePolicyDefaults
	}
	if model.Replications != nil {
		replications := []map[string]interface{}{}
		for _, replicationsItem := range model.Replications {
			replicationsItemMap, err := resourceIbmProtectionGroupRunRequestRunReplicationConfigToMap(&replicationsItem)
			if err != nil {
				return modelMap, err
			}
			replications = append(replications, replicationsItemMap)
		}
		modelMap["replications"] = replications
	}
	if model.Archivals != nil {
		archivals := []map[string]interface{}{}
		for _, archivalsItem := range model.Archivals {
			archivalsItemMap, err := resourceIbmProtectionGroupRunRequestRunArchivalConfigToMap(&archivalsItem)
			if err != nil {
				return modelMap, err
			}
			archivals = append(archivals, archivalsItemMap)
		}
		modelMap["archivals"] = archivals
	}
	if model.CloudReplications != nil {
		cloudReplications := []map[string]interface{}{}
		for _, cloudReplicationsItem := range model.CloudReplications {
			cloudReplicationsItemMap, err := resourceIbmProtectionGroupRunRequestRunCloudReplicationConfigToMap(&cloudReplicationsItem)
			if err != nil {
				return modelMap, err
			}
			cloudReplications = append(cloudReplications, cloudReplicationsItemMap)
		}
		modelMap["cloud_replications"] = cloudReplications
	}
	return modelMap, nil
}

func resourceIbmProtectionGroupRunRequestRunReplicationConfigToMap(model *backuprecoveryv1.RunReplicationConfig) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["id"] = flex.IntValue(model.ID)
	if model.Retention != nil {
		retentionMap, err := resourceIbmProtectionGroupRunRequestRetentionToMap(model.Retention)
		if err != nil {
			return modelMap, err
		}
		modelMap["retention"] = []map[string]interface{}{retentionMap}
	}
	return modelMap, nil
}

func resourceIbmProtectionGroupRunRequestRetentionToMap(model *backuprecoveryv1.Retention) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["unit"] = model.Unit
	modelMap["duration"] = flex.IntValue(model.Duration)
	if model.DataLockConfig != nil {
		dataLockConfigMap, err := resourceIbmProtectionGroupRunRequestDataLockConfigToMap(model.DataLockConfig)
		if err != nil {
			return modelMap, err
		}
		modelMap["data_lock_config"] = []map[string]interface{}{dataLockConfigMap}
	}
	return modelMap, nil
}

func resourceIbmProtectionGroupRunRequestDataLockConfigToMap(model *backuprecoveryv1.DataLockConfig) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["mode"] = model.Mode
	modelMap["unit"] = model.Unit
	modelMap["duration"] = flex.IntValue(model.Duration)
	if model.EnableWormOnExternalTarget != nil {
		modelMap["enable_worm_on_external_target"] = model.EnableWormOnExternalTarget
	}
	return modelMap, nil
}

func resourceIbmProtectionGroupRunRequestRunArchivalConfigToMap(model *backuprecoveryv1.RunArchivalConfig) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["id"] = flex.IntValue(model.ID)
	modelMap["archival_target_type"] = model.ArchivalTargetType
	if model.Retention != nil {
		retentionMap, err := resourceIbmProtectionGroupRunRequestRetentionToMap(model.Retention)
		if err != nil {
			return modelMap, err
		}
		modelMap["retention"] = []map[string]interface{}{retentionMap}
	}
	if model.CopyOnlyFullySuccessful != nil {
		modelMap["copy_only_fully_successful"] = model.CopyOnlyFullySuccessful
	}
	return modelMap, nil
}

func resourceIbmProtectionGroupRunRequestKeyValuePairToMap(model *backuprecoveryv1.KeyValuePair) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["key"] = model.Key
	modelMap["value"] = model.Value
	return modelMap, nil
}
