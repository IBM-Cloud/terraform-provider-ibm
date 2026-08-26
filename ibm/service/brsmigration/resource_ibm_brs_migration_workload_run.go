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

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.ibm.com/BackupAndRecovery/brs-migration-orchestrator/brsmigrationv2"
)

func ResourceIbmBrsMigrationWorkloadRun() *schema.Resource {
	return &schema.Resource{
		CreateContext:   resourceIbmBrsMigrationWorkloadRunCreate,
		ReadContext:     resourceIbmBrsMigrationWorkloadRunRead,
		DeleteContext:   resourceIbmBrsMigrationWorkloadRunDelete,
		Importer: &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"migration_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				ValidateFunc: validate.InvokeValidator("ibm_brs_migration_workload_run", "migration_id"),
				Description: "The migration project ID (mgr-{uuid4} format).",
			},
			"workload_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				ValidateFunc: validate.InvokeValidator("ibm_brs_migration_workload_run", "workload_id"),
				Description: "The migration service workload ID (wl-{uuid4} format).",
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
							Optional:    true,
							Computed:    true,
							Description: "Total logical size of all data processed, in bytes.",
						},
						"bytes_transferred": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Computed:    true,
							Description: "Number of bytes successfully transferred.",
						},
						"bytes_read": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Computed:    true,
							Description: "Number of bytes read from the source.",
						},
						"total_file_count": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Computed:    true,
							Description: "Total number of files or objects processed.",
						},
						"transferred_file_count": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
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
							Optional:    true,
							Computed:    true,
							Description: "Error or warning detail specific to this payload.",
						},
						"stats": &schema.Schema{
							Type:        schema.TypeList,
							Optional:    true,
							Computed:    true,
							Description: "Data-transfer statistics for a workload run or payload.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"logical_size_bytes": &schema.Schema{
										Type:        schema.TypeInt,
										Optional:    true,
										Computed:    true,
										Description: "Total logical size of all data processed, in bytes.",
									},
									"bytes_transferred": &schema.Schema{
										Type:        schema.TypeInt,
										Optional:    true,
										Computed:    true,
										Description: "Number of bytes successfully transferred.",
									},
									"bytes_read": &schema.Schema{
										Type:        schema.TypeInt,
										Optional:    true,
										Computed:    true,
										Description: "Number of bytes read from the source.",
									},
									"total_file_count": &schema.Schema{
										Type:        schema.TypeInt,
										Optional:    true,
										Computed:    true,
										Description: "Total number of files or objects processed.",
									},
									"transferred_file_count": &schema.Schema{
										Type:        schema.TypeInt,
										Optional:    true,
										Computed:    true,
										Description: "Number of files or objects successfully transferred.",
									},
								},
							},
						},
					},
				},
			},
			"run_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Unique run ID (run-{uuid4} format).",
			},
		},
	}
}

func ResourceIbmBrsMigrationWorkloadRunValidator() *validate.ResourceValidator {
	validateSchema := make([]validate.ValidateSchema, 0)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "migration_id",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Required:                   true,
			Regexp:                     `^mgr-[0-9a-f-]{36}$`,
			MinValueLength:             40,
			MaxValueLength:             40,
		},
		validate.ValidateSchema{
			Identifier:                 "workload_id",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Required:                   true,
			Regexp:                     `^wl-[0-9a-f-]{36}$`,
			MinValueLength:             39,
			MaxValueLength:             39,
		},
	)

	resourceValidator := validate.ResourceValidator{ResourceName: "ibm_brs_migration_workload_run", Schema: validateSchema}
	return &resourceValidator
}

func resourceIbmBrsMigrationWorkloadRunCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	brsMigrationClient, err := meta.(conns.ClientSession).BrsMigrationV2()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_brs_migration_workload_run", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	createWorkloadRunOptions := &brsmigrationv2.CreateWorkloadRunOptions{}

	createWorkloadRunOptions.SetMigrationID(d.Get("migration_id").(string))
	createWorkloadRunOptions.SetWorkloadID(d.Get("workload_id").(string))
	createWorkloadRunOptions.SetBackupType(d.Get("backup_type").(string))

	workloadRunAccepted, _, err := brsMigrationClient.CreateWorkloadRunWithContext(context, createWorkloadRunOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreateWorkloadRunWithContext failed: %s", err.Error()), "ibm_brs_migration_workload_run", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(fmt.Sprintf("%s/%s/%s", *createWorkloadRunOptions.MigrationID, *createWorkloadRunOptions.WorkloadID, *workloadRunAccepted.ID))

	return resourceIbmBrsMigrationWorkloadRunRead(context, d, meta)
}

func resourceIbmBrsMigrationWorkloadRunRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	brsMigrationClient, err := meta.(conns.ClientSession).BrsMigrationV2()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_brs_migration_workload_run", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getWorkloadRunOptions := &brsmigrationv2.GetWorkloadRunOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_brs_migration_workload_run", "read", "sep-id-parts").GetDiag()
	}

	getWorkloadRunOptions.SetMigrationID(parts[0])
	getWorkloadRunOptions.SetWorkloadID(parts[1])
	getWorkloadRunOptions.SetRunID(parts[2])

	workloadRun, response, err := brsMigrationClient.GetWorkloadRunWithContext(context, getWorkloadRunOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetWorkloadRunWithContext failed: %s", err.Error()), "ibm_brs_migration_workload_run", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	if err = d.Set("workload_id", workloadRun.WorkloadID); err != nil {
		err = fmt.Errorf("Error setting workload_id: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_brs_migration_workload_run", "read", "set-workload_id").GetDiag()
	}
	if err = d.Set("operation_type", workloadRun.OperationType); err != nil {
		err = fmt.Errorf("Error setting operation_type: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_brs_migration_workload_run", "read", "set-operation_type").GetDiag()
	}
	if err = d.Set("run_type", workloadRun.RunType); err != nil {
		err = fmt.Errorf("Error setting run_type: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_brs_migration_workload_run", "read", "set-run_type").GetDiag()
	}
	if err = d.Set("status", workloadRun.Status); err != nil {
		err = fmt.Errorf("Error setting status: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_brs_migration_workload_run", "read", "set-status").GetDiag()
	}
	if err = d.Set("started_at", flex.DateTimeToString(workloadRun.StartedAt)); err != nil {
		err = fmt.Errorf("Error setting started_at: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_brs_migration_workload_run", "read", "set-started_at").GetDiag()
	}
	if !core.IsNil(workloadRun.CompletedAt) {
		if err = d.Set("completed_at", flex.DateTimeToString(workloadRun.CompletedAt)); err != nil {
			err = fmt.Errorf("Error setting completed_at: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_brs_migration_workload_run", "read", "set-completed_at").GetDiag()
		}
	}
	if !core.IsNil(workloadRun.DurationSeconds) {
		if err = d.Set("duration_seconds", flex.IntValue(workloadRun.DurationSeconds)); err != nil {
			err = fmt.Errorf("Error setting duration_seconds: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_brs_migration_workload_run", "read", "set-duration_seconds").GetDiag()
		}
	}
	if !core.IsNil(workloadRun.Message) {
		if err = d.Set("message", workloadRun.Message); err != nil {
			err = fmt.Errorf("Error setting message: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_brs_migration_workload_run", "read", "set-message").GetDiag()
		}
	}
	if !core.IsNil(workloadRun.Stats) {
		statsMap, err := ResourceIbmBrsMigrationWorkloadRunWorkloadRunStatsToMap(workloadRun.Stats)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_brs_migration_workload_run", "read", "stats-to-map").GetDiag()
		}
		if err = d.Set("stats", []map[string]interface{}{statsMap}); err != nil {
			err = fmt.Errorf("Error setting stats: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_brs_migration_workload_run", "read", "set-stats").GetDiag()
		}
	}
	payloadResults := []map[string]interface{}{}
	for _, payloadResultsItem := range workloadRun.PayloadResults {
		payloadResultsItemMap, err := ResourceIbmBrsMigrationWorkloadRunPayloadResultToMap(&payloadResultsItem) // #nosec G601
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_brs_migration_workload_run", "read", "payload_results-to-map").GetDiag()
		}
		payloadResults = append(payloadResults, payloadResultsItemMap)
	}
	if err = d.Set("payload_results", payloadResults); err != nil {
		err = fmt.Errorf("Error setting payload_results: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_brs_migration_workload_run", "read", "set-payload_results").GetDiag()
	}
	if err = d.Set("run_id", workloadRun.ID); err != nil {
		err = fmt.Errorf("Error setting run_id: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_brs_migration_workload_run", "read", "set-run_id").GetDiag()
	}

	return nil
}

func resourceIbmBrsMigrationWorkloadRunDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	// This resource does not support a "delete" operation.
	d.SetId("")
	return nil
}

func ResourceIbmBrsMigrationWorkloadRunWorkloadRunStatsToMap(model *brsmigrationv2.WorkloadRunStats) (map[string]interface{}, error) {
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

func ResourceIbmBrsMigrationWorkloadRunPayloadResultToMap(model *brsmigrationv2.PayloadResult) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["payload_id"] = *model.PayloadID
	modelMap["status"] = *model.Status
	if model.Message != nil {
		modelMap["message"] = *model.Message
	}
	if model.Stats != nil {
		statsMap, err := ResourceIbmBrsMigrationWorkloadRunPayloadResultStatsToMap(model.Stats)
		if err != nil {
			return modelMap, err
		}
		modelMap["stats"] = []map[string]interface{}{statsMap}
	}
	return modelMap, nil
}

func ResourceIbmBrsMigrationWorkloadRunPayloadResultStatsToMap(model *brsmigrationv2.PayloadResultStats) (map[string]interface{}, error) {
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
