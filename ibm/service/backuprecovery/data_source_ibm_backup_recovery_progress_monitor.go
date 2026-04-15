// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.113.1-d76630af-20260320-135953
 */

package backuprecovery

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/ibm-backup-recovery-sdk-go/backuprecoveryv1"
)

func DataSourceIbmBackupRecoveryProgressMonitor() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmBackupRecoveryProgressMonitorRead,

		Schema: map[string]*schema.Schema{
			"x_ibm_tenant_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Specifies the unique id of the tenant.",
			},
			"attribute_vec": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "If specified, tasks matching the current query are futher filtered by these KeyValuePairs. This gives client an ability to search by custom attributes that they specified during the task creation. Only the Tasks having 'all' of the specified key=value pairs will be returned.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"key": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies the name of the key.",
						},
						"value": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Specifies a value for the key.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"data": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Specifies the fields to store data of a given type.Specify data in the appropriate field for the current data type.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"oneof_data": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "Types that are valid to be assigned to OneofData:Value_Data_Int64ValueValue_Data_DoubleValueValue_Data_StringValueValue_Data_BytesValue.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{},
													},
												},
											},
										},
									},
									"type": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Specifies the type of value. 0 specifies a data point of type Int64. 1 specifies a data point of type Double. 2 specifies a data point of type String. 3 specifies a data point of type Bytes.",
									},
								},
							},
						},
					},
				},
			},
			"end_time_secs": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Tasks that ended before this time.",
			},
			"exclude_sub_tasks": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Skip information about the sub tasks of the matching root and sub tasks. By default, the entire task tree will be returned for matching tasks.",
			},
			"fetch_logs_max_level": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Number of levels till which we need to fetch the event logs for a pulse tree. Note that it is applicable only when include_event_logs is true.",
			},
			"include_event_logs": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "If set, the event logs will be included in the response message. Otherwise they will be cleared out.",
			},
			"include_finished_tasks": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Returns finished tasks as well. By default, Pulse only returns active tasks.",
			},
			"max_tasks": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Only return at most these many matching tasks. This constraint is applied with each query's result group.",
			},
			"start_time_secs": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Tasks that started after this time.",
			},
			"task_path_vec": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "The hierarchical paths to the names of the tasks being queried. The task path-name specified here can be a prefix. Clients can specify multiple paths/prefixes. Pulse will return one ResultGroup for each path query.Each path is treated separately by Pulse, so if there are duplicate paths, Pulse will return duplicate results.Both root tasks and sub tasks can be specified in @task_path_vec.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"error": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Proto to describe the error returned by pulse.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"error_msg": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "A string describing the errors encountered.",
						},
						"type": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The type of error encountered.",
						},
					},
				},
			},
			"result_group_vec": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"task_vec": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "All tasks that match the corresponding query.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"progress": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The progress on this task.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"approx_percent_unknown_work": &schema.Schema{
													Type:        schema.TypeFloat,
													Computed:    true,
													Description: "If set this indicate the percentage of work which is not know at this time. This will be useful if client does not know total amount of work that has to done. But client know how much work it has completed and approximate how much more work need to be done. This is usually reported by the clients for leaf tasks. For non-leaf tasks, the progress may be dynamically inferred.(see ReportTaskProgressArg).",
												},
												"attribute_vec": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "The latest attributes (if any) reported for this task.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"key": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "key.",
															},
															"value": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "value.",
															},
														},
													},
												},
												"end_time_secs": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "The time when the task finished.",
												},
												"event_vec": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "The events (if any) reported for this task.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"event_msg": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "Message associated with the event.",
															},
															"owner_percent_finished": &schema.Schema{
																Type:        schema.TypeFloat,
																Computed:    true,
																Description: "How much the owning task completed when this event occurred.",
															},
															"owner_remaining_work_count": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "How much work was remaining for the owning task when this event occurred.",
															},
															"timestamp_secs": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "The timestamp at which the event occurred.",
															},
														},
													},
												},
												"expected_end_time_secs": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "The expected end time of this task (if it hasn't ended). This is extrapolated using the current progress, and any historic data about this task if it occurs periodically. TODO(gaurav): Deprecate this field once Iris has stopped using it.",
												},
												"expected_time_remaining_secs": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Expected time remaining for this task (if it hasn't ended).",
												},
												"expected_total_work_count": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "The expected raw count of the total work remaining. This is the highest work count value reported by the client. This field can be set to let pulse compute percent_finished by looking at the currently reported remaining_work_count and the expected_total_work_count.",
												},
												"last_update_time_secs": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "The timestamp at which task progress was last reported.",
												},
												"percent_finished": &schema.Schema{
													Type:        schema.TypeFloat,
													Computed:    true,
													Description: "The reported progress on this task. This is usually reported by clients for leaf tasks. For non-leaf tasks, the progress may be dynamically inferred.(see ReportTaskProgressArg).",
												},
												"start_time_secs": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "The time when the task was started.",
												},
												"status": &schema.Schema{
													Type:        schema.TypeList,
													Computed:    true,
													Description: "The status of the task.",
													Elem: &schema.Resource{
														Schema: map[string]*schema.Schema{
															"error_msg": &schema.Schema{
																Type:        schema.TypeString,
																Computed:    true,
																Description: "The error message (if any).",
															},
															"type": &schema.Schema{
																Type:        schema.TypeInt,
																Computed:    true,
																Description: "The return type.",
															},
														},
													},
												},
											},
										},
									},
									"sub_task_vec": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Information about all the sub tasks for this task.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"task_path": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The hierarchical name of the task.",
									},
									"weight": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The weight of this task.",
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceIbmBackupRecoveryProgressMonitorRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	backupRecoveryClient, err := meta.(conns.ClientSession).BackupRecoveryV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_backup_recovery_progress_monitor", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getProgressMonitorsOptions := &backuprecoveryv1.GetProgressMonitorsOptions{}

	getProgressMonitorsOptions.XIBMTenantID = (d.Get("x_ibm_tenant_id").(*string))
	if _, ok := d.GetOk("attribute_vec"); ok {
		var newAttributeVec []backuprecoveryv1.KeyValuePair
		for _, v := range d.Get("attribute_vec").([]interface{}) {
			value := v.(map[string]interface{})
			newAttributeVecItem, err := ResourceMapToKeyValueProgressPair(value)
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_backup_recovery_progress_monitor", "read", "parse-attribute_vec").GetDiag()
			}
			newAttributeVec = append(newAttributeVec, *newAttributeVecItem)
		}
		getProgressMonitorsOptions.SetNewAttributeVec(newAttributeVec)
	}
	if _, ok := d.GetOk("end_time_secs"); ok {
		getProgressMonitorsOptions.SetNewEndTimeSecs(int64(d.Get("end_time_secs").(int)))
	}
	if _, ok := d.GetOk("exclude_sub_tasks"); ok {
		getProgressMonitorsOptions.SetNewExcludeSubTasks(d.Get("exclude_sub_tasks").(bool))
	}
	if _, ok := d.GetOk("fetch_logs_max_level"); ok {
		getProgressMonitorsOptions.SetNewFetchLogsMaxLevel(int64(d.Get("fetch_logs_max_level").(int)))
	}
	if _, ok := d.GetOk("include_event_logs"); ok {
		getProgressMonitorsOptions.SetNewIncludeEventLogs(d.Get("include_event_logs").(bool))
	}
	if _, ok := d.GetOk("include_finished_tasks"); ok {
		getProgressMonitorsOptions.SetNewIncludeFinishedTasks(d.Get("include_finished_tasks").(bool))
	}
	if _, ok := d.GetOk("max_tasks"); ok {
		getProgressMonitorsOptions.SetNewMaxTasks(int64(d.Get("max_tasks").(int)))
	}
	if _, ok := d.GetOk("start_time_secs"); ok {
		getProgressMonitorsOptions.SetNewStartTimeSecs(int64(d.Get("start_time_secs").(int)))
	}
	if _, ok := d.GetOk("task_path_vec"); ok {
		var newTaskPathVec []string
		for _, v := range d.Get("task_path_vec").([]interface{}) {
			newTaskPathVecItem := v.(string)
			newTaskPathVec = append(newTaskPathVec, newTaskPathVecItem)
		}
		getProgressMonitorsOptions.SetNewTaskPathVec(newTaskPathVec)
	}
	if _, ok := d.GetOk("task_path_vec"); ok {
		var taskPathVec []string
		for _, v := range d.Get("task_path_vec").([]interface{}) {
			taskPathVecItem := v.(string)
			taskPathVec = append(taskPathVec, taskPathVecItem)
		}
		getProgressMonitorsOptions.SetTaskPathVec(taskPathVec)
	}
	if _, ok := d.GetOk("include_finished_tasks"); ok {
		getProgressMonitorsOptions.SetIncludeFinishedTasks(d.Get("include_finished_tasks").(bool))
	}
	if _, ok := d.GetOk("start_time_secs"); ok {
		getProgressMonitorsOptions.SetStartTimeSecs(int64(d.Get("start_time_secs").(int)))
	}
	if _, ok := d.GetOk("end_time_secs"); ok {
		getProgressMonitorsOptions.SetEndTimeSecs(int64(d.Get("end_time_secs").(int)))
	}
	if _, ok := d.GetOk("max_tasks"); ok {
		getProgressMonitorsOptions.SetMaxTasks(int64(d.Get("max_tasks").(int)))
	}
	if _, ok := d.GetOk("exclude_sub_tasks"); ok {
		getProgressMonitorsOptions.SetExcludeSubTasks(d.Get("exclude_sub_tasks").(bool))
	}

	getTasksResult, _, err := backupRecoveryClient.GetProgressMonitorsWithContext(context, getProgressMonitorsOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetProgressMonitorsWithContext failed: %s", err.Error()), "(Data) ibm_backup_recovery_progress_monitor", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(dataSourceIbmBackupRecoveryProgressMonitorID(d))

	if !core.IsNil(getTasksResult.Error) {
		error := []map[string]interface{}{}
		errorMap, err := DataSourceIbmBackupRecoveryProgressMonitorPrivateErrorProtoToMap(getTasksResult.Error)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_backup_recovery_progress_monitor", "read", "error-to-map").GetDiag()
		}
		error = append(error, errorMap)
		if err = d.Set("error", error); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting error: %s", err), "(Data) ibm_backup_recovery_progress_monitor", "read", "set-error").GetDiag()
		}
	}

	if !core.IsNil(getTasksResult.ResultGroupVec) {
		resultGroupVec := []map[string]interface{}{}
		for _, resultGroupVecItem := range getTasksResult.ResultGroupVec {
			resultGroupVecItemMap, err := DataSourceIbmBackupRecoveryProgressMonitorGetTasksResultResultGroupToMap(&resultGroupVecItem) // #nosec G601
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_backup_recovery_progress_monitor", "read", "result_group_vec-to-map").GetDiag()
			}
			resultGroupVec = append(resultGroupVec, resultGroupVecItemMap)
		}
		if err = d.Set("result_group_vec", resultGroupVec); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting result_group_vec: %s", err), "(Data) ibm_backup_recovery_progress_monitor", "read", "set-result_group_vec").GetDiag()
		}
	}

	return nil
}

func ResourceMapToKeyValueProgressPair(modelMap map[string]interface{}) (*backuprecoveryv1.KeyValuePair, error) {
	model := &backuprecoveryv1.KeyValuePair{}
	model.Key = core.StringPtr(modelMap["key"].(string))
	model.Value = core.StringPtr(modelMap["value"].(string))
	return model, nil
}

// dataSourceIbmBackupRecoveryProgressMonitorID returns a reasonable ID for the list.
func dataSourceIbmBackupRecoveryProgressMonitorID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func DataSourceIbmBackupRecoveryProgressMonitorPrivateErrorProtoToMap(model *backuprecoveryv1.PrivateErrorProto) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ErrorMsg != nil {
		modelMap["error_msg"] = *model.ErrorMsg
	}
	if model.Type != nil {
		modelMap["type"] = flex.IntValue(model.Type)
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryProgressMonitorGetTasksResultResultGroupToMap(model *backuprecoveryv1.GetTasksResultResultGroup) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.TaskVec != nil {
		taskVec := []map[string]interface{}{}
		for _, taskVecItem := range model.TaskVec {
			taskVecItemMap, err := DataSourceIbmBackupRecoveryProgressMonitorGetTasksResultResultGroupTaskToMap(&taskVecItem) // #nosec G601
			if err != nil {
				return modelMap, err
			}
			taskVec = append(taskVec, taskVecItemMap)
		}
		modelMap["task_vec"] = taskVec
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryProgressMonitorGetTasksResultResultGroupTaskToMap(model *backuprecoveryv1.GetTasksResultResultGroupTask) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Progress != nil {
		progressMap, err := DataSourceIbmBackupRecoveryProgressMonitorTaskProgressToMap(model.Progress)
		if err != nil {
			return modelMap, err
		}
		modelMap["progress"] = []map[string]interface{}{progressMap}
	}
	if model.SubTaskVec != nil {
		modelMap["sub_task_vec"] = model.SubTaskVec
	}
	if model.TaskPath != nil {
		modelMap["task_path"] = *model.TaskPath
	}
	if model.Weight != nil {
		modelMap["weight"] = flex.IntValue(model.Weight)
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryProgressMonitorTaskProgressToMap(model *backuprecoveryv1.TaskProgress) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ApproxPercentUnknownWork != nil {
		modelMap["approx_percent_unknown_work"] = flex.Float64Value(model.ApproxPercentUnknownWork)
	}
	if model.AttributeVec != nil {
		attributeVec := []map[string]interface{}{}
		for _, attributeVecItem := range model.AttributeVec {
			attributeVecItemMap, err := DataSourceIbmBackupRecoveryProgressMonitorKeyValuePairToMap(&attributeVecItem) // #nosec G601
			if err != nil {
				return modelMap, err
			}
			attributeVec = append(attributeVec, attributeVecItemMap)
		}
		modelMap["attribute_vec"] = attributeVec
	}
	if model.EndTimeSecs != nil {
		modelMap["end_time_secs"] = flex.IntValue(model.EndTimeSecs)
	}
	if model.EventVec != nil {
		eventVec := []map[string]interface{}{}
		for _, eventVecItem := range model.EventVec {
			eventVecItemMap, err := DataSourceIbmBackupRecoveryProgressMonitorPrivateTaskEventToMap(&eventVecItem) // #nosec G601
			if err != nil {
				return modelMap, err
			}
			eventVec = append(eventVec, eventVecItemMap)
		}
		modelMap["event_vec"] = eventVec
	}
	if model.ExpectedEndTimeSecs != nil {
		modelMap["expected_end_time_secs"] = flex.IntValue(model.ExpectedEndTimeSecs)
	}
	if model.ExpectedTimeRemainingSecs != nil {
		modelMap["expected_time_remaining_secs"] = flex.IntValue(model.ExpectedTimeRemainingSecs)
	}
	if model.ExpectedTotalWorkCount != nil {
		modelMap["expected_total_work_count"] = flex.IntValue(model.ExpectedTotalWorkCount)
	}
	if model.LastUpdateTimeSecs != nil {
		modelMap["last_update_time_secs"] = flex.IntValue(model.LastUpdateTimeSecs)
	}
	if model.PercentFinished != nil {
		modelMap["percent_finished"] = flex.Float64Value(model.PercentFinished)
	}
	if model.StartTimeSecs != nil {
		modelMap["start_time_secs"] = flex.IntValue(model.StartTimeSecs)
	}
	if model.Status != nil {
		statusMap, err := DataSourceIbmBackupRecoveryProgressMonitorTaskStatusToMap(model.Status)
		if err != nil {
			return modelMap, err
		}
		modelMap["status"] = []map[string]interface{}{statusMap}
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryProgressMonitorKeyValuePairToMap(model *backuprecoveryv1.KeyValueProgressPair) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["key"] = *model.Key
	modelMap["value"] = *model.Value
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryProgressMonitorPrivateTaskEventToMap(model *backuprecoveryv1.PrivateTaskEvent) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.EventMsg != nil {
		modelMap["event_msg"] = *model.EventMsg
	}
	if model.OwnerPercentFinished != nil {
		modelMap["owner_percent_finished"] = flex.Float64Value(model.OwnerPercentFinished)
	}
	if model.OwnerRemainingWorkCount != nil {
		modelMap["owner_remaining_work_count"] = flex.IntValue(model.OwnerRemainingWorkCount)
	}
	if model.TimestampSecs != nil {
		modelMap["timestamp_secs"] = flex.IntValue(model.TimestampSecs)
	}
	return modelMap, nil
}

func DataSourceIbmBackupRecoveryProgressMonitorTaskStatusToMap(model *backuprecoveryv1.TaskStatus) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ErrorMsg != nil {
		modelMap["error_msg"] = *model.ErrorMsg
	}
	if model.Type != nil {
		modelMap["type"] = flex.IntValue(model.Type)
	}
	return modelMap, nil
}
