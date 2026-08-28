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
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/ibm-brs-migration-sdk-go/brsmigrationv1"
)

func DataSourceIbmBrsMigrationDiscover() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmBrsMigrationDiscoverRead,

		Schema: map[string]*schema.Schema{
			"migration_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The migration project ID (mgr-{uuid4} format).",
			},
			"job_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The unique ID of the discovery job (job-{uuid4} format).",
			},
			"env": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Infrastructure environment being discovered.",
			},
			"state": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Current lifecycle state of the discovery job.",
			},
			"start_time": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Start of the time window used for this discovery run.",
			},
			"end_time": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "End of the time window used for this discovery run.",
			},
			"message": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Human-readable status or error message.",
			},
			"summary": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Counts of discovered resources by compute and storage type.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"total": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Total number of compute resources discovered.",
						},
						"compute": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Compute resource counts by type.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"virtual_server": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Number of Virtual Server Instances discovered.",
									},
									"bare_metal": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Number of bare metal servers discovered.",
									},
								},
							},
						},
						"storage": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Storage volume counts by type.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"block": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Number of block volumes discovered.",
									},
									"file": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Number of file shares discovered.",
									},
									"san": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Number of SAN volumes discovered (Classic only).",
									},
									"local": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "Number of local disks discovered.",
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

func dataSourceIbmBrsMigrationDiscoverRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	brsMigrationClient, err := meta.(conns.ClientSession).BrsMigrationV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_brs_migration_discover", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getDiscoverOptions := &brsmigrationv1.GetDiscoverOptions{}

	getDiscoverOptions.SetMigrationID(d.Get("migration_id").(string))
	getDiscoverOptions.SetJobID(d.Get("job_id").(string))

	discoverJob, _, err := brsMigrationClient.GetDiscoverWithContext(context, getDiscoverOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetDiscoverWithContext failed: %s", err.Error()), "(Data) ibm_brs_migration_discover", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(fmt.Sprintf("%s/%s", *getDiscoverOptions.MigrationID, *getDiscoverOptions.JobID))

	if err = d.Set("env", discoverJob.Env); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting env: %s", err), "(Data) ibm_brs_migration_discover", "read", "set-env").GetDiag()
	}

	if err = d.Set("state", discoverJob.State); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting state: %s", err), "(Data) ibm_brs_migration_discover", "read", "set-state").GetDiag()
	}

	if !core.IsNil(discoverJob.StartTime) {
		if err = d.Set("start_time", flex.DateTimeToString(discoverJob.StartTime)); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting start_time: %s", err), "(Data) ibm_brs_migration_discover", "read", "set-start_time").GetDiag()
		}
	}

	if !core.IsNil(discoverJob.EndTime) {
		if err = d.Set("end_time", flex.DateTimeToString(discoverJob.EndTime)); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting end_time: %s", err), "(Data) ibm_brs_migration_discover", "read", "set-end_time").GetDiag()
		}
	}

	if !core.IsNil(discoverJob.Message) {
		if err = d.Set("message", discoverJob.Message); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting message: %s", err), "(Data) ibm_brs_migration_discover", "read", "set-message").GetDiag()
		}
	}

	if !core.IsNil(discoverJob.Summary) {
		summary := []map[string]interface{}{}
		summaryMap, err := DataSourceIbmBrsMigrationDiscoverDiscoverJobSummaryToMap(discoverJob.Summary)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_brs_migration_discover", "read", "summary-to-map").GetDiag()
		}
		summary = append(summary, summaryMap)
		if err = d.Set("summary", summary); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting summary: %s", err), "(Data) ibm_brs_migration_discover", "read", "set-summary").GetDiag()
		}
	}

	return nil
}

func DataSourceIbmBrsMigrationDiscoverDiscoverJobSummaryToMap(model *brsmigrationv1.DiscoverJobSummary) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Total != nil {
		modelMap["total"] = flex.IntValue(model.Total)
	}
	if model.Compute != nil {
		computeMap, err := DataSourceIbmBrsMigrationDiscoverDiscoverJobSummaryComputeToMap(model.Compute)
		if err != nil {
			return modelMap, err
		}
		modelMap["compute"] = []map[string]interface{}{computeMap}
	}
	if model.Storage != nil {
		storageMap, err := DataSourceIbmBrsMigrationDiscoverDiscoverJobSummaryStorageToMap(model.Storage)
		if err != nil {
			return modelMap, err
		}
		modelMap["storage"] = []map[string]interface{}{storageMap}
	}
	return modelMap, nil
}

func DataSourceIbmBrsMigrationDiscoverDiscoverJobSummaryComputeToMap(model *brsmigrationv1.DiscoverJobSummaryCompute) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.VirtualServer != nil {
		modelMap["virtual_server"] = flex.IntValue(model.VirtualServer)
	}
	if model.BareMetal != nil {
		modelMap["bare_metal"] = flex.IntValue(model.BareMetal)
	}
	return modelMap, nil
}

func DataSourceIbmBrsMigrationDiscoverDiscoverJobSummaryStorageToMap(model *brsmigrationv1.DiscoverJobSummaryStorage) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Block != nil {
		modelMap["block"] = flex.IntValue(model.Block)
	}
	if model.File != nil {
		modelMap["file"] = flex.IntValue(model.File)
	}
	if model.San != nil {
		modelMap["san"] = flex.IntValue(model.San)
	}
	if model.Local != nil {
		modelMap["local"] = flex.IntValue(model.Local)
	}
	return modelMap, nil
}
