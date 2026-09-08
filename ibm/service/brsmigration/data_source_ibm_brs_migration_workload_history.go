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

func DataSourceIbmBrsMigrationWorkloadHistory() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmBrsMigrationWorkloadHistoryRead,

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
			"history": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Workload execution history entries on this page.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Unique identifier for this history entry.",
						},
						"state": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Final execution state of this history entry.",
						},
						"started_at": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Timestamp when this run started.",
						},
						"completed_at": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Timestamp when this run completed.",
						},
						"message": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Human-readable status or error message.",
						},
					},
				},
			},
		},
	}
}

func dataSourceIbmBrsMigrationWorkloadHistoryRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	brsMigrationClient, err := meta.(conns.ClientSession).BrsMigrationV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_brs_migration_workload_history", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	listWorkloadHistoryOptions := &brsmigrationv1.ListWorkloadHistoryOptions{}

	listWorkloadHistoryOptions.SetMigrationID(d.Get("migration_id").(string))
	listWorkloadHistoryOptions.SetWorkloadID(d.Get("workload_id").(string))

	var pager *brsmigrationv1.WorkloadHistoryPager
	pager, err = brsMigrationClient.NewWorkloadHistoryPager(listWorkloadHistoryOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, err.Error(), "(Data) ibm_brs_migration_workload_history", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	allItems, err := pager.GetAll()
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("WorkloadHistoryPager.GetAll() failed %s", err), "(Data) ibm_brs_migration_workload_history", "read")
		log.Printf("[DEBUG] %s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(dataSourceIbmBrsMigrationWorkloadHistoryID(d))

	mapSlice := []map[string]interface{}{}
	for _, modelItem := range allItems {
		modelMap, err := DataSourceIbmBrsMigrationWorkloadHistoryWorkloadHistoryEntryToMap(&modelItem) // #nosec G601
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_brs_migration_workload_history", "read", "WorkloadHistory-to-map").GetDiag()
		}
		mapSlice = append(mapSlice, modelMap)
	}

	if err = d.Set("history", mapSlice); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting history %s", err), "(Data) ibm_brs_migration_workload_history", "read", "history-set").GetDiag()
	}

	return nil
}

// dataSourceIbmBrsMigrationWorkloadHistoryID returns a reasonable ID for the list.
func dataSourceIbmBrsMigrationWorkloadHistoryID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func DataSourceIbmBrsMigrationWorkloadHistoryWorkloadHistoryEntryToMap(model *brsmigrationv1.WorkloadHistoryEntry) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["id"] = *model.ID
	modelMap["state"] = *model.State
	modelMap["started_at"] = model.StartedAt.String()
	if model.CompletedAt != nil {
		modelMap["completed_at"] = model.CompletedAt.String()
	}
	if model.Message != nil {
		modelMap["message"] = *model.Message
	}
	return modelMap, nil
}
