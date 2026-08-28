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

func DataSourceIbmBrsMigrationDiscovers() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmBrsMigrationDiscoversRead,

		Schema: map[string]*schema.Schema{
			"migration_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The migration project ID (mgr-{uuid4} format).",
			},
			"env": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Filter discovery jobs by infrastructure environment.",
			},
			"state": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Filter discovery jobs by current state.",
			},
			"discover": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Discovery jobs on this page.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Unique discovery job ID (job-{uuid4} format).",
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
				},
			},
		},
	}
}

func dataSourceIbmBrsMigrationDiscoversRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	brsMigrationClient, err := meta.(conns.ClientSession).BrsMigrationV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_brs_migration_discovers", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	listDiscoverOptions := &brsmigrationv1.ListDiscoverOptions{}

	listDiscoverOptions.SetMigrationID(d.Get("migration_id").(string))
	if _, ok := d.GetOk("env"); ok {
		listDiscoverOptions.SetEnv(d.Get("env").(string))
	}
	if _, ok := d.GetOk("state"); ok {
		listDiscoverOptions.SetState(d.Get("state").(string))
	}

	var pager *brsmigrationv1.DiscoverPager
	pager, err = brsMigrationClient.NewDiscoverPager(listDiscoverOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, err.Error(), "(Data) ibm_brs_migration_discovers", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	allItems, err := pager.GetAll()
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DiscoverPager.GetAll() failed %s", err), "(Data) ibm_brs_migration_discovers", "read")
		log.Printf("[DEBUG] %s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(dataSourceIbmBrsMigrationDiscoversID(d))

	mapSlice := []map[string]interface{}{}
	for _, modelItem := range allItems {
		modelMap, err := DataSourceIbmBrsMigrationDiscoversDiscoverJobToMap(&modelItem) // #nosec G601
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_brs_migration_discovers", "read", "Discovery-to-map").GetDiag()
		}
		mapSlice = append(mapSlice, modelMap)
	}

	if err = d.Set("discover", mapSlice); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting discover %s", err), "(Data) ibm_brs_migration_discovers", "read", "discover-set").GetDiag()
	}

	return nil
}

// dataSourceIbmBrsMigrationDiscoversID returns a reasonable ID for the list.
func dataSourceIbmBrsMigrationDiscoversID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func DataSourceIbmBrsMigrationDiscoversDiscoverJobToMap(model *brsmigrationv1.DiscoverJob) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["id"] = *model.ID
	modelMap["env"] = *model.Env
	modelMap["state"] = *model.State
	if model.StartTime != nil {
		modelMap["start_time"] = model.StartTime.String()
	}
	if model.EndTime != nil {
		modelMap["end_time"] = model.EndTime.String()
	}
	if model.Message != nil {
		modelMap["message"] = *model.Message
	}
	if model.Summary != nil {
		summaryMap, err := DataSourceIbmBrsMigrationDiscoversDiscoverJobSummaryToMap(model.Summary)
		if err != nil {
			return modelMap, err
		}
		modelMap["summary"] = []map[string]interface{}{summaryMap}
	}
	return modelMap, nil
}

func DataSourceIbmBrsMigrationDiscoversDiscoverJobSummaryToMap(model *brsmigrationv1.DiscoverJobSummary) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Total != nil {
		modelMap["total"] = flex.IntValue(model.Total)
	}
	if model.Compute != nil {
		computeMap, err := DataSourceIbmBrsMigrationDiscoversDiscoverJobSummaryComputeToMap(model.Compute)
		if err != nil {
			return modelMap, err
		}
		modelMap["compute"] = []map[string]interface{}{computeMap}
	}
	if model.Storage != nil {
		storageMap, err := DataSourceIbmBrsMigrationDiscoversDiscoverJobSummaryStorageToMap(model.Storage)
		if err != nil {
			return modelMap, err
		}
		modelMap["storage"] = []map[string]interface{}{storageMap}
	}
	return modelMap, nil
}

func DataSourceIbmBrsMigrationDiscoversDiscoverJobSummaryComputeToMap(model *brsmigrationv1.DiscoverJobSummaryCompute) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.VirtualServer != nil {
		modelMap["virtual_server"] = flex.IntValue(model.VirtualServer)
	}
	if model.BareMetal != nil {
		modelMap["bare_metal"] = flex.IntValue(model.BareMetal)
	}
	return modelMap, nil
}

func DataSourceIbmBrsMigrationDiscoversDiscoverJobSummaryStorageToMap(model *brsmigrationv1.DiscoverJobSummaryStorage) (map[string]interface{}, error) {
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
