// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.116.0-df613dbc-20260803-154903
 */

package powerhaautomationservice

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/dra-go-sdk/powerhaautomationservicev1"
)

func DataSourceIBMPhaPowervsWorkspace() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMPhaPowervsWorkspaceRead,

		Schema: map[string]*schema.Schema{
			"pha_instance_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Unique identifier of the provisioned instance.",
			},
			"location_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Location ID value.",
			},
			"accept_language": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The language requested for the return document.",
			},
			"if_none_match": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "ETag for conditional requests (optional).",
			},
			"workspaces": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Array of workspace summaries within the region.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Unique identifier of the workspace.",
						},
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Name of the workspace.",
						},
					},
				},
			},
		},
	}
}

func dataSourceIBMPhaPowervsWorkspaceRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	powerhaAutomationServiceClient, err := meta.(conns.ClientSession).PowerhaAutomationServiceV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_pha_powervs_workspace", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getPhaPowervsWorkspaceOptions := &powerhaautomationservicev1.GetPhaPowervsWorkspaceOptions{}

	getPhaPowervsWorkspaceOptions.SetPhaInstanceID(d.Get("pha_instance_id").(string))
	getPhaPowervsWorkspaceOptions.SetLocationID(d.Get("location_id").(string))
	if _, ok := d.GetOk("accept_language"); ok {
		getPhaPowervsWorkspaceOptions.SetAcceptLanguage(d.Get("accept_language").(string))
	}
	if _, ok := d.GetOk("if_none_match"); ok {
		getPhaPowervsWorkspaceOptions.SetIfNoneMatch(d.Get("if_none_match").(string))
	}

	phaWorkspacesRegionResponse, response, err := powerhaAutomationServiceClient.GetPhaPowervsWorkspaceWithContext(context, getPhaPowervsWorkspaceOptions)
	if err != nil {
		detailedMsg := fmt.Sprintf("GetPowervsWorkspaceWithContext failed: %s", err.Error())
		// Include HTTP status & raw body if available
		if response != nil {
			detailedMsg = fmt.Sprintf(
				"GetPowervsWorkspaceWithContext failed: %s (status: %d, response: %s)",
				err.Error(), response.StatusCode, response.Result,
			)
		}
		tfErr := flex.TerraformErrorf(err, detailedMsg, "ibm_pha_powervs_workspaces", "create")
		log.Printf("[ERROR] %s", detailedMsg)
		return tfErr.GetDiag()
	}

	d.SetId(*phaWorkspacesRegionResponse.ID)

	workspaces := []map[string]interface{}{}
	for _, workspacesItem := range phaWorkspacesRegionResponse.Workspaces {
		workspacesItemMap, err := DataSourceIBMPhaPowervsWorkspacePhaWorkspaceSummaryToMap(&workspacesItem) // #nosec G601
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_pha_powervs_workspace", "read", "workspaces-to-map").GetDiag()
		}
		workspaces = append(workspaces, workspacesItemMap)
	}
	if err = d.Set("workspaces", workspaces); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting workspaces: %s", err), "(Data) ibm_pha_powervs_workspace", "read", "set-workspaces").GetDiag()
	}

	return nil
}

func DataSourceIBMPhaPowervsWorkspacePhaWorkspaceSummaryToMap(model *powerhaautomationservicev1.PhaWorkspaceSummary) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ID != nil {
		modelMap["id"] = *model.ID
	}
	if model.Name != nil {
		modelMap["name"] = *model.Name
	}
	return modelMap, nil
}
