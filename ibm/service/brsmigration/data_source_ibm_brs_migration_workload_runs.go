// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.114.3-943fbc81-20260603-173645
*/

package brsmigration

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/ibm-brs-migration-sdk-go/brsmigrationv1"
)

func DataSourceIbmBrsMigrationWorkloadRuns() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmBrsMigrationWorkloadRunsRead,

		Schema: map[string]*schema.Schema{
			"migration_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The migration project ID (mgr-{uuid4} format).",
			},
			"workload_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The migration service workload ID (wl-{uuid4} format).",
			},
			"status": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "Filter by run status.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"run_type": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Filter by run type (scheduled or on-demand).",
			},
			"runs": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of runs ordered by `startedAt` descending.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Unique run ID (run-{uuid4} format).",
						},
						"workload_id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "ID of the parent workload.",
						},
						"operation_type": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Whether this run is a backup or a restore operation.",
						},
						"run_type": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Whether this run was triggered on-demand or by the schedule.",
						},
						"status": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Current execution status of the run.",
						},
						"started_at": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Time the run started (ISO 8601 UTC).",
						},
						"completed_at": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Time the run completed. Null if still in progress.",
						},
						"duration_seconds": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Wall-clock duration of the run in seconds.",
						},
						"message": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Human-readable status message or error detail.",
						},
						"stats": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Data-transfer statistics for a workload run or payload.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"logical_size_bytes": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Total logical size of all data processed, in bytes.",
									},
									"bytes_transferred": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Number of bytes successfully transferred.",
									},
									"bytes_read": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Number of bytes read from the source.",
									},
									"total_file_count": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Total number of files or objects processed.",
									},
									"transferred_file_count": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Number of files or objects successfully transferred.",
									},
								},
							},
						},
						"payload_results": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Per-payload breakdown of the run.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"payload_id": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "ID of the workload payload this result belongs to.",
									},
									"status": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Status of this individual payload transfer.",
									},
									"message": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Error or warning detail specific to this payload.",
									},
									"stats": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "Data-transfer statistics for a workload run or payload.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"logical_size_bytes": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Total logical size of all data processed, in bytes.",
												},
												"bytes_transferred": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Number of bytes successfully transferred.",
												},
												"bytes_read": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Number of bytes read from the source.",
												},
												"total_file_count": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Total number of files or objects processed.",
												},
												"transferred_file_count": &schema.Schema{
													Type:        schema.TypeInt,
													Computed:    true,
													Description: "Number of files or objects successfully transferred.",
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
	}
}

func dataSourceIbmBrsMigrationWorkloadRunsRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	brsMigrationClient, err := meta.(conns.ClientSession).BrsMigrationV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_brs_migration_workload_runs", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	listWorkloadRunsOptions := &brsmigrationv1.ListWorkloadRunsOptions{}

	listWorkloadRunsOptions.SetMigrationID(d.Get("migration_id").(string))
	listWorkloadRunsOptions.SetWorkloadID(d.Get("workload_id").(string))
	if _, ok := d.GetOk("status"); ok {
		var status []string
		for _, v := range d.Get("status").([]interface{}) {
			statusItem := v.(string)
			status = append(status, statusItem)
		}
		listWorkloadRunsOptions.SetStatus(status)
	}
	if _, ok := d.GetOk("run_type"); ok {
		listWorkloadRunsOptions.SetRunType(d.Get("run_type").(string))
	}

	var pager *brsmigrationv1.WorkloadRunsPager
	pager, err = brsMigrationClient.NewWorkloadRunsPager(listWorkloadRunsOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, err.Error(), "(Data) ibm_brs_migration_workload_runs", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	allItems, err := pager.GetAll()
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("WorkloadRunsPager.GetAll() failed %s", err), "(Data) ibm_brs_migration_workload_runs", "read")
		log.Printf("[DEBUG] %s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(dataSourceIbmBrsMigrationWorkloadRunsID(d))

	mapSlice := []map[string]interface{}{}
	for _, modelItem := range allItems {
		modelMap, err := DataSourceIbmBrsMigrationWorkloadRunsWorkloadRunToMap(&modelItem) // #nosec G601
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_brs_migration_workload_runs", "read", "WorkloadRuns-to-map").GetDiag()
		}
		mapSlice = append(mapSlice, modelMap)
	}

	if err = d.Set("runs", mapSlice); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting runs %s", err), "(Data) ibm_brs_migration_workload_runs", "read", "runs-set").GetDiag()
	}

	return nil
}

// dataSourceIbmBrsMigrationWorkloadRunsID returns a reasonable ID for the list.
func dataSourceIbmBrsMigrationWorkloadRunsID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func DataSourceIbmBrsMigrationWorkloadRunsWorkloadRunToMap(model *brsmigrationv1.WorkloadRun) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["id"] = *model.ID
	modelMap["workload_id"] = *model.WorkloadID
	modelMap["operation_type"] = *model.OperationType
	modelMap["run_type"] = *model.RunType
	modelMap["status"] = *model.Status
	modelMap["started_at"] = model.StartedAt.String()
	if model.CompletedAt != nil {
		modelMap["completed_at"] = model.CompletedAt.String()
	}
	if model.DurationSeconds != nil {
		modelMap["duration_seconds"] = flex.IntValue(model.DurationSeconds)
	}
	if model.Message != nil {
		modelMap["message"] = *model.Message
	}
	if model.Stats != nil {
		statsMap, err := DataSourceIbmBrsMigrationWorkloadRunsWorkloadRunStatsToMap(model.Stats)
		if err != nil {
			return modelMap, err
		}
		modelMap["stats"] = []map[string]interface{}{statsMap}
	}
	payloadResults := []map[string]interface{}{}
	for _, payloadResultsItem := range model.PayloadResults {
		payloadResultsItemMap, err := DataSourceIbmBrsMigrationWorkloadRunsPayloadResultToMap(&payloadResultsItem) // #nosec G601
		if err != nil {
			return modelMap, err
		}
		payloadResults = append(payloadResults, payloadResultsItemMap)
	}
	modelMap["payload_results"] = payloadResults
	return modelMap, nil
}

func DataSourceIbmBrsMigrationWorkloadRunsWorkloadRunStatsToMap(model *brsmigrationv1.WorkloadRunStats) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.LogicalSizeBytes != nil {
		modelMap["logical_size_bytes"] = flex.IntValue(model.LogicalSizeBytes)
	}
	if model.BytesTransferred != nil {
		modelMap["bytes_transferred"] = flex.IntValue(model.BytesTransferred)
	}
	if model.BytesRead != nil {
		modelMap["bytes_read"] = flex.IntValue(model.BytesRead)
	}
	if model.TotalFileCount != nil {
		modelMap["total_file_count"] = flex.IntValue(model.TotalFileCount)
	}
	if model.TransferredFileCount != nil {
		modelMap["transferred_file_count"] = flex.IntValue(model.TransferredFileCount)
	}
	return modelMap, nil
}

func DataSourceIbmBrsMigrationWorkloadRunsPayloadResultToMap(model *brsmigrationv1.PayloadResult) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["payload_id"] = *model.PayloadID
	modelMap["status"] = *model.Status
	if model.Message != nil {
		modelMap["message"] = *model.Message
	}
	if model.Stats != nil {
		statsMap, err := DataSourceIbmBrsMigrationWorkloadRunsPayloadResultStatsToMap(model.Stats)
		if err != nil {
			return modelMap, err
		}
		modelMap["stats"] = []map[string]interface{}{statsMap}
	}
	return modelMap, nil
}

func DataSourceIbmBrsMigrationWorkloadRunsPayloadResultStatsToMap(model *brsmigrationv1.PayloadResultStats) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.LogicalSizeBytes != nil {
		modelMap["logical_size_bytes"] = flex.IntValue(model.LogicalSizeBytes)
	}
	if model.BytesTransferred != nil {
		modelMap["bytes_transferred"] = flex.IntValue(model.BytesTransferred)
	}
	if model.BytesRead != nil {
		modelMap["bytes_read"] = flex.IntValue(model.BytesRead)
	}
	if model.TotalFileCount != nil {
		modelMap["total_file_count"] = flex.IntValue(model.TotalFileCount)
	}
	if model.TransferredFileCount != nil {
		modelMap["transferred_file_count"] = flex.IntValue(model.TransferredFileCount)
	}
	return modelMap, nil
}
