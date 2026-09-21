// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.116.0-df613dbc-20260803-154903
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
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/dra-go-sdk/drautomationservicev1"
)

func DataSourceIBMPdrPowervsWorkspace() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMPdrPowervsWorkspaceRead,

		Schema: map[string]*schema.Schema{
			"instance_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Service Instance ID.",
			},
			"location_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Location ID value.",
			},
			"dr_workspaces": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The list of primary disaster recovery workspaces.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"details": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Details of the DR workspace.",
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
							Description: "The unique identifier of the DR workspace.",
						},
						"location": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Represents a disaster recovery location.",
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
							Description: "The name of the DR workspace.",
						},
						"status": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of the DR workspace.",
						},
						"default": &schema.Schema{
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Indicates if this is the default DR workspace.",
						},
					},
				},
			},
			"dr_standby_workspaces": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The list of standby disaster recovery workspaces.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"details": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Details of the DR workspace.",
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
							Description: "The unique identifier of the standby workspace.",
						},
						"location": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Represents a disaster recovery location.",
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
							Description: "The name of the standby workspace.",
						},
						"status": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The status of the standby workspace.",
						},
					},
				},
			},
			"dr_workspace_description": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Description of Workspace.",
			},
			"dr_standby_workspace_description": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Description of Standby Workspace.",
			},
		},
	}
}

func dataSourceIBMPdrPowervsWorkspaceRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	drAutomationServiceClient, err := meta.(conns.ClientSession).DrAutomationServiceV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_pdr_powervs_workspace", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getPowervsWorkspaceOptions := &drautomationservicev1.GetPowervsWorkspaceOptions{}

	getPowervsWorkspaceOptions.SetInstanceID(d.Get("instance_id").(string))
	getPowervsWorkspaceOptions.SetLocationID(d.Get("location_id").(string))

	drData, response, err := drAutomationServiceClient.GetPowervsWorkspaceWithContext(context, getPowervsWorkspaceOptions)
	if err != nil {
		detailedMsg := fmt.Sprintf("GetPowervsWorkspacesWithContext failed: %s", err.Error())
		// Include HTTP status & raw body if available
		if response != nil {
			detailedMsg = fmt.Sprintf(
				"GetPowervsWorkspacesWithContext failed: %s (status: %d, response: %s)",
				err.Error(), response.StatusCode, response.Result,
			)
		}
		tfErr := flex.TerraformErrorf(err, detailedMsg, "(Data) ibm_pdr_powervs_workspaces", "read")
		log.Printf("[ERROR] %s", detailedMsg)
		return tfErr.GetDiag()
	}

	d.SetId(dataSourceIBMPdrPowervsWorkspaceID(d))

	drWorkspaces := []map[string]interface{}{}
	for _, drWorkspacesItem := range drData.DrWorkspaces {
		drWorkspacesItemMap, err := DataSourceIBMPdrPowervsWorkspaceDrWorkspaceToMap(&drWorkspacesItem) // #nosec G601
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_pdr_powervs_workspace", "read", "dr_workspaces-to-map").GetDiag()
		}
		drWorkspaces = append(drWorkspaces, drWorkspacesItemMap)
	}
	if err = d.Set("dr_workspaces", drWorkspaces); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting dr_workspaces: %s", err), "(Data) ibm_pdr_powervs_workspace", "read", "set-dr_workspaces").GetDiag()
	}

	drStandbyWorkspaces := []map[string]interface{}{}
	for _, drStandbyWorkspacesItem := range drData.DrStandbyWorkspaces {
		drStandbyWorkspacesItemMap, err := DataSourceIBMPdrPowervsWorkspaceDrStandbyWorkspaceToMap(&drStandbyWorkspacesItem) // #nosec G601
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_pdr_powervs_workspace", "read", "dr_standby_workspaces-to-map").GetDiag()
		}
		drStandbyWorkspaces = append(drStandbyWorkspaces, drStandbyWorkspacesItemMap)
	}
	if err = d.Set("dr_standby_workspaces", drStandbyWorkspaces); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting dr_standby_workspaces: %s", err), "(Data) ibm_pdr_powervs_workspace", "read", "set-dr_standby_workspaces").GetDiag()
	}

	if !core.IsNil(drData.DrWorkspaceDescription) {
		if err = d.Set("dr_workspace_description", drData.DrWorkspaceDescription); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting dr_workspace_description: %s", err), "(Data) ibm_pdr_powervs_workspace", "read", "set-dr_workspace_description").GetDiag()
		}
	}

	if !core.IsNil(drData.DrStandbyWorkspaceDescription) {
		if err = d.Set("dr_standby_workspace_description", drData.DrStandbyWorkspaceDescription); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting dr_standby_workspace_description: %s", err), "(Data) ibm_pdr_powervs_workspace", "read", "set-dr_standby_workspace_description").GetDiag()
		}
	}

	return nil
}

// dataSourceIBMPdrPowervsWorkspaceID returns a reasonable ID for the list.
func dataSourceIBMPdrPowervsWorkspaceID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func DataSourceIBMPdrPowervsWorkspaceDrWorkspaceToMap(model *drautomationservicev1.DrWorkspace) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Details != nil {
		detailsMap, err := DataSourceIBMPdrPowervsWorkspaceDetailsDrToMap(model.Details)
		if err != nil {
			return modelMap, err
		}
		modelMap["details"] = []map[string]interface{}{detailsMap}
	}
	if model.ID != nil {
		modelMap["id"] = *model.ID
	}
	if model.Location != nil {
		locationMap, err := DataSourceIBMPdrPowervsWorkspaceLocationDrToMap(model.Location)
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
	if model.Default != nil {
		modelMap["default"] = *model.Default
	}
	return modelMap, nil
}

func DataSourceIBMPdrPowervsWorkspaceDetailsDrToMap(model *drautomationservicev1.DetailsDr) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.CRN != nil {
		modelMap["crn"] = *model.CRN
	}
	return modelMap, nil
}

func DataSourceIBMPdrPowervsWorkspaceLocationDrToMap(model *drautomationservicev1.LocationDr) (map[string]interface{}, error) {
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

func DataSourceIBMPdrPowervsWorkspaceDrStandbyWorkspaceToMap(model *drautomationservicev1.DrStandbyWorkspace) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Details != nil {
		detailsMap, err := DataSourceIBMPdrPowervsWorkspaceDetailsDrToMap(model.Details)
		if err != nil {
			return modelMap, err
		}
		modelMap["details"] = []map[string]interface{}{detailsMap}
	}
	if model.ID != nil {
		modelMap["id"] = *model.ID
	}
	if model.Location != nil {
		locationMap, err := DataSourceIBMPdrPowervsWorkspaceLocationDrToMap(model.Location)
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
