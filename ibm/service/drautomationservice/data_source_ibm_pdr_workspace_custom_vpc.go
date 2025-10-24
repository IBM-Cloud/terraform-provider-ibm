// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.105.0-3c13b041-20250605-193116
 */

package drautomationservice

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.ibm.com/DRAutomation/dra-go-sdk/drautomationservicev1"
)

func DataSourceIbmPdrWorkspaceCustomVpc() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmPdrWorkspaceCustomVpcRead,

		Schema: map[string]*schema.Schema{
			"instance_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "instance id of instance to provision.",
			},
			"location_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Location ID value.",
			},
			"vpc_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "vpc id value.",
			},
			"tg_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "transit gateway id value.",
			},
			"if_none_match": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "ETag for conditional requests (optional).",
			},
			"dr_standby_workspaces": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of standby disaster recovery workspaces.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"details": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Detailed information of the standby DR workspace.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"crn": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Cloud Resource Name (CRN) of the DR workspace.",
									},
								},
							},
						},
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Unique identifier of the standby workspace.",
						},
						"location": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Location information of the standby workspace.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"region": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The region identifier of the DR location.",
									},
									"type": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The type of location (e.g., data-center, cloud-region).",
									},
									"url": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The URL endpoint to access the DR location.",
									},
								},
							},
						},
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Name of the standby workspace.",
						},
						"status": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Current status of the standby workspace.",
						},
					},
				},
			},
			"dr_workspaces": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of primary disaster recovery workspaces.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"default": &schema.Schema{
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Indicates if this is the default DR workspace.",
						},
						"details": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Detailed information about the DR workspace.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"crn": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Cloud Resource Name (CRN) of the DR workspace.",
									},
								},
							},
						},
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Unique identifier of the DR workspace.",
						},
						"location": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Location information of the DR workspace.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"region": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The region identifier of the DR location.",
									},
									"type": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The type of location (e.g., data-center, cloud-region).",
									},
									"url": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The URL endpoint to access the DR location.",
									},
								},
							},
						},
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Name of the DR workspace.",
						},
						"status": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Current status of the DR workspace.",
						},
					},
				},
			},
		},
	}
}

func dataSourceIbmPdrWorkspaceCustomVpcRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	drAutomationServiceClient, err := meta.(conns.ClientSession).DrAutomationServiceV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_pdr_workspace_custom_vpc", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getPowervsWorkspacesForCustomVpcOptions := &drautomationservicev1.GetPowervsWorkspacesForCustomVPCOptions{}

	getPowervsWorkspacesForCustomVpcOptions.SetInstanceID(d.Get("instance_id").(string))
	getPowervsWorkspacesForCustomVpcOptions.SetLocationID(d.Get("location_id").(string))
	getPowervsWorkspacesForCustomVpcOptions.SetVPCID(d.Get("vpc_id").(string))
	getPowervsWorkspacesForCustomVpcOptions.SetTgID(d.Get("tg_id").(string))
	if _, ok := d.GetOk("if_none_match"); ok {
		getPowervsWorkspacesForCustomVpcOptions.SetIfNoneMatch(d.Get("if_none_match").(string))
	}

	drDataCustomVPC, _, err := drAutomationServiceClient.GetPowervsWorkspacesForCustomVPCWithContext(context, getPowervsWorkspacesForCustomVpcOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetPowervsWorkspacesForCustomVpcWithContext failed: %s", err.Error()), "(Data) ibm_pdr_workspace_custom_vpc", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(dataSourceIbmPdrWorkspaceCustomVpcID(d))

	drStandbyWorkspaces := []map[string]interface{}{}
	for _, drStandbyWorkspacesItem := range drDataCustomVPC.DrStandbyWorkspaces {
		drStandbyWorkspacesItemMap, err := DataSourceIbmPdrWorkspaceCustomVpcDRStandbyWorkspaceToMap(&drStandbyWorkspacesItem) // #nosec G601
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_pdr_workspace_custom_vpc", "read", "dr_standby_workspaces-to-map").GetDiag()
		}
		drStandbyWorkspaces = append(drStandbyWorkspaces, drStandbyWorkspacesItemMap)
	}
	if err = d.Set("dr_standby_workspaces", drStandbyWorkspaces); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting dr_standby_workspaces: %s", err), "(Data) ibm_pdr_workspace_custom_vpc", "read", "set-dr_standby_workspaces").GetDiag()
	}

	drWorkspaces := []map[string]interface{}{}
	for _, drWorkspacesItem := range drDataCustomVPC.DrWorkspaces {
		drWorkspacesItemMap, err := DataSourceIbmPdrWorkspaceCustomVpcDRWorkspaceToMap(&drWorkspacesItem) // #nosec G601
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_pdr_workspace_custom_vpc", "read", "dr_workspaces-to-map").GetDiag()
		}
		drWorkspaces = append(drWorkspaces, drWorkspacesItemMap)
	}
	if err = d.Set("dr_workspaces", drWorkspaces); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting dr_workspaces: %s", err), "(Data) ibm_pdr_workspace_custom_vpc", "read", "set-dr_workspaces").GetDiag()
	}

	return nil
}

// dataSourceIbmPdrWorkspaceCustomVpcID returns a reasonable ID for the list.
func dataSourceIbmPdrWorkspaceCustomVpcID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func DataSourceIbmPdrWorkspaceCustomVpcDRStandbyWorkspaceToMap(model *drautomationservicev1.DrStandbyWorkspace) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Details != nil {
		detailsMap, err := DataSourceIbmPdrWorkspaceCustomVpcDetailsDrToMap(model.Details)
		if err != nil {
			return modelMap, err
		}
		modelMap["details"] = []map[string]interface{}{detailsMap}
	}
	if model.ID != nil {
		modelMap["id"] = *model.ID
	}
	if model.Location != nil {
		locationMap, err := DataSourceIbmPdrWorkspaceCustomVpcLocationDrToMap(model.Location)
		if err != nil {
			return modelMap, err
		}
		modelMap["location"] = []map[string]interface{}{locationMap}
	}
	if model.Name != nil {
		modelMap["name"] = *model.Name
	}
	if model.Status != nil {
		modelMap["status"] = *model.Status
	}
	return modelMap, nil
}

func DataSourceIbmPdrWorkspaceCustomVpcDetailsDrToMap(model *drautomationservicev1.DetailsDr) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.CRN != nil {
		modelMap["crn"] = *model.CRN
	}
	return modelMap, nil
}

func DataSourceIbmPdrWorkspaceCustomVpcLocationDrToMap(model *drautomationservicev1.LocationDr) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Region != nil {
		modelMap["region"] = *model.Region
	}
	if model.Type != nil {
		modelMap["type"] = *model.Type
	}
	if model.URL != nil {
		modelMap["url"] = *model.URL
	}
	return modelMap, nil
}

func DataSourceIbmPdrWorkspaceCustomVpcDRWorkspaceToMap(model *drautomationservicev1.DrWorkspace) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Default != nil {
		modelMap["default"] = *model.Default
	}
	if model.Details != nil {
		detailsMap, err := DataSourceIbmPdrWorkspaceCustomVpcDetailsDrToMap(model.Details)
		if err != nil {
			return modelMap, err
		}
		modelMap["details"] = []map[string]interface{}{detailsMap}
	}
	if model.ID != nil {
		modelMap["id"] = *model.ID
	}
	if model.Location != nil {
		locationMap, err := DataSourceIbmPdrWorkspaceCustomVpcLocationDrToMap(model.Location)
		if err != nil {
			return modelMap, err
		}
		modelMap["location"] = []map[string]interface{}{locationMap}
	}
	if model.Name != nil {
		modelMap["name"] = *model.Name
	}
	if model.Status != nil {
		modelMap["status"] = *model.Status
	}
	return modelMap, nil
}
