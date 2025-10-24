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

func DataSourceIbmPdrSchematicWorkspaces() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmPdrSchematicWorkspacesRead,

		Schema: map[string]*schema.Schema{
			"instance_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "instance id of instance to provision.",
			},
			"if_none_match": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "ETag for conditional requests (optional).",
			},
			"workspaces": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of Schematics workspaces associated with the DR automation service instance.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"catalog_ref": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Reference to a catalog item associated with the DR automation workspace.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"item_name": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Name of the catalog item that defines the resource or configuration.",
									},
								},
							},
						},
						"created_at": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Timestamp when the Schematics workspace was created, in ISO 8601 format (UTC).",
						},
						"created_by": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "CRN of the user or service that created the Schematics workspace.",
						},
						"crn": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Cloud Resource Name (CRN) of the Schematics workspace.",
						},
						"description": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Detailed description of the Schematics workspace.",
						},
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Unique identifier of the Schematics workspace.",
						},
						"location": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Region where the Schematics workspace is hosted.",
						},
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Human-readable name of the Schematics workspace.",
						},
						"status": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Current lifecycle status of the Schematics workspace.",
						},
					},
				},
			},
		},
	}
}

func dataSourceIbmPdrSchematicWorkspacesRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	drAutomationServiceClient, err := meta.(conns.ClientSession).DrAutomationServiceV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_pdr_schematic_workspaces", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getSchematicWorkspaceOptions := &drautomationservicev1.GetSchematicWorkspaceOptions{}

	getSchematicWorkspaceOptions.SetInstanceID(d.Get("instance_id").(string))
	if _, ok := d.GetOk("if_none_match"); ok {
		getSchematicWorkspaceOptions.SetIfNoneMatch(d.Get("if_none_match").(string))
	}

	schematicWorkspacesResponse, _, err := drAutomationServiceClient.GetSchematicWorkspaceWithContext(context, getSchematicWorkspaceOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetSchematicWorkspaceWithContext failed: %s", err.Error()), "(Data) ibm_pdr_schematic_workspaces", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(dataSourceIbmPdrSchematicWorkspacesID(d))

	workspaces := []map[string]interface{}{}
	for _, workspacesItem := range schematicWorkspacesResponse.Workspaces {
		workspacesItemMap, err := DataSourceIbmPdrSchematicWorkspacesDrAutomationSchematicsWorkspaceToMap(&workspacesItem) // #nosec G601
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_pdr_schematic_workspaces", "read", "workspaces-to-map").GetDiag()
		}
		workspaces = append(workspaces, workspacesItemMap)
	}
	if err = d.Set("workspaces", workspaces); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting workspaces: %s", err), "(Data) ibm_pdr_schematic_workspaces", "read", "set-workspaces").GetDiag()
	}

	return nil
}

// dataSourceIbmPdrSchematicWorkspacesID returns a reasonable ID for the list.
func dataSourceIbmPdrSchematicWorkspacesID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func DataSourceIbmPdrSchematicWorkspacesDrAutomationSchematicsWorkspaceToMap(model *drautomationservicev1.DrAutomationSchematicsWorkspace) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.CatalogRef != nil {
		catalogRefMap, err := DataSourceIbmPdrSchematicWorkspacesDrAutomationCatalogRefToMap(model.CatalogRef)
		if err != nil {
			return modelMap, err
		}
		modelMap["catalog_ref"] = []map[string]interface{}{catalogRefMap}
	}
	if model.CreatedAt != nil {
		modelMap["created_at"] = model.CreatedAt.String()
	}
	if model.CreatedBy != nil {
		modelMap["created_by"] = *model.CreatedBy
	}
	if model.CRN != nil {
		modelMap["crn"] = *model.CRN
	}
	if model.Description != nil {
		modelMap["description"] = *model.Description
	}
	if model.ID != nil {
		modelMap["id"] = *model.ID
	}
	if model.Location != nil {
		modelMap["location"] = *model.Location
	}
	if model.Name != nil {
		modelMap["name"] = *model.Name
	}
	if model.Status != nil {
		modelMap["status"] = *model.Status
	}
	return modelMap, nil
}

func DataSourceIbmPdrSchematicWorkspacesDrAutomationCatalogRefToMap(model *drautomationservicev1.DrAutomationCatalogRef) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ItemName != nil {
		modelMap["item_name"] = *model.ItemName
	}
	return modelMap, nil
}
