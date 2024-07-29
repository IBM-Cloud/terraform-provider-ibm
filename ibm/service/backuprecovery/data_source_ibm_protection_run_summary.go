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
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.ibm.com/BackupAndRecovery/ibm-backup-recovery-sdk-go/backuprecoveryv1"
)

func DataSourceIbmProtectionRunSummary() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmProtectionRunSummaryRead,

		Schema: map[string]*schema.Schema{
			"start_time_usecs": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Start time for time range filter. Specify the start time as a Unix epoch Timestamp (in microseconds), only runs executing after this time will be returned. By default it is endTimeUsecs minus an hour.",
			},
			"end_time_usecs": &schema.Schema{
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "End time for time range filter. Specify the end time as a Unix epoch Timestamp (in microseconds), only runs executing before this time will be returned. By default it is current time.",
			},
			"run_status": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Specifies a list of status, runs matching the status will be returned.<br> 'Running' indicates that the run is still running.<br> 'Canceled' indicates that the run has been canceled.<br> 'Canceling' indicates that the run is in the process of being canceled.<br> 'Failed' indicates that the run has failed.<br> 'Missed' indicates that the run was unable to take place at the scheduled time because the previous run was still happening.<br> 'Succeeded' indicates that the run has finished successfully.<br> 'SucceededWithWarning' indicates that the run finished successfully, but there were some warning messages.<br> 'Skipped' indicates that the run was skipped.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"protection_runs_summary": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Specifies a list of summaries of protection runs.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies the ID of the Protection Group run.",
						},
						"protection_group_id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "ProtectionGroupId to which this run belongs.",
						},
						"protection_group_name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Name of the Protection Group to which this run belongs.",
						},
						"environment": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Specifies the environment type of the Protection Group.",
						},
						"is_sla_violated": &schema.Schema{
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Indicated if SLA has been violated for this run.",
						},
						"start_time_usecs": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Specifies the start time of backup run in Unix epoch Timestamp(in microseconds).",
						},
						"end_time_usecs": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Specifies the end time of backup run in Unix epoch Timestamp(in microseconds).",
						},
						"status": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Status of the backup run. 'Running' indicates that the run is still running. 'Canceled' indicates that the run has been canceled. 'Canceling' indicates that the run is in the process of being canceled. 'Paused' indicates that the ongoing run has been paused. 'Failed' indicates that the run has failed. 'Missed' indicates that the run was unable to take place at the scheduled time because the previous run was still happening. 'Succeeded' indicates that the run has finished successfully. 'SucceededWithWarning' indicates that the run finished successfully, but there were some warning messages. 'Skipped' indicates that the run was skipped.",
						},
						"is_full_run": &schema.Schema{
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Specifies if the protection run is a full run.",
						},
						"total_objects_count": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Specifies the total number of objects protected in this run.",
						},
						"success_objects_count": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Specifies the number of objects which are successfully protected in this run.",
						},
						"logical_size_bytes": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Specifies total logical size of object(s) in bytes.",
						},
						"bytes_written": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Specifies total size of data in bytes written after taking backup.",
						},
					},
				},
			},
		},
	}
}

func dataSourceIbmProtectionRunSummaryRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	backupRecoveryClient, err := meta.(conns.ClientSession).BackupRecoveryV1()
	if err != nil {
		return diag.FromErr(err)
	}

	getProtectionRunsOptions := &backuprecoveryv1.GetProtectionRunsOptions{}

	if _, ok := d.GetOk("start_time_usecs"); ok {
		getProtectionRunsOptions.SetStartTimeUsecs(int64(d.Get("start_time_usecs").(int)))
	}
	if _, ok := d.GetOk("end_time_usecs"); ok {
		getProtectionRunsOptions.SetEndTimeUsecs(int64(d.Get("end_time_usecs").(int)))
	}
	if _, ok := d.GetOk("run_status"); ok {
		var runStatus []string
		for _, v := range d.Get("run_status").([]interface{}) {
			runStatusItem := v.(string)
			runStatus = append(runStatus, runStatusItem)
		}
		getProtectionRunsOptions.SetRunStatus(runStatus)
	}

	protectionRunsSummaryResponse, response, err := backupRecoveryClient.GetProtectionRunsWithContext(context, getProtectionRunsOptions)
	if err != nil {
		log.Printf("[DEBUG] GetProtectionRunsWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("GetProtectionRunsWithContext failed %s\n%s", err, response))
	}

	d.SetId(dataSourceIbmProtectionRunSummaryID(d))

	protectionRunsSummary := []map[string]interface{}{}
	if protectionRunsSummaryResponse.ProtectionRunsSummary != nil {
		for _, modelItem := range protectionRunsSummaryResponse.ProtectionRunsSummary {
			modelMap, err := dataSourceIbmProtectionRunSummaryProtectionRunSummaryToMap(&modelItem)
			if err != nil {
				return diag.FromErr(err)
			}
			protectionRunsSummary = append(protectionRunsSummary, modelMap)
		}
	}
	if err = d.Set("protection_runs_summary", protectionRunsSummary); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting protection_runs_summary %s", err))
	}

	return nil
}

// dataSourceIbmProtectionRunSummaryID returns a reasonable ID for the list.
func dataSourceIbmProtectionRunSummaryID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func dataSourceIbmProtectionRunSummaryProtectionRunSummaryToMap(model *backuprecoveryv1.ProtectionRunSummary) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ID != nil {
		modelMap["id"] = model.ID
	}
	if model.ProtectionGroupID != nil {
		modelMap["protection_group_id"] = model.ProtectionGroupID
	}
	if model.ProtectionGroupName != nil {
		modelMap["protection_group_name"] = model.ProtectionGroupName
	}
	if model.Environment != nil {
		modelMap["environment"] = model.Environment
	}
	if model.IsSlaViolated != nil {
		modelMap["is_sla_violated"] = model.IsSlaViolated
	}
	if model.StartTimeUsecs != nil {
		modelMap["start_time_usecs"] = flex.IntValue(model.StartTimeUsecs)
	}
	if model.EndTimeUsecs != nil {
		modelMap["end_time_usecs"] = flex.IntValue(model.EndTimeUsecs)
	}
	if model.Status != nil {
		modelMap["status"] = model.Status
	}
	if model.IsFullRun != nil {
		modelMap["is_full_run"] = model.IsFullRun
	}
	if model.TotalObjectsCount != nil {
		modelMap["total_objects_count"] = flex.IntValue(model.TotalObjectsCount)
	}
	if model.SuccessObjectsCount != nil {
		modelMap["success_objects_count"] = flex.IntValue(model.SuccessObjectsCount)
	}
	if model.LogicalSizeBytes != nil {
		modelMap["logical_size_bytes"] = flex.IntValue(model.LogicalSizeBytes)
	}
	if model.BytesWritten != nil {
		modelMap["bytes_written"] = flex.IntValue(model.BytesWritten)
	}
	return modelMap, nil
}
