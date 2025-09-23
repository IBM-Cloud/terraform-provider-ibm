// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package iampolicy

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM/platform-services-go-sdk/iampolicymanagementv1"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/conns"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/flex"
)

func DataSourceIBMIAMActionControlTemplate() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMActionControlTemplateRead,

		Schema: map[string]*schema.Schema{
			"action_control_templates": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of action control templates.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "name of template.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "description of template purpose.",
						},
						"account_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "account id where this template will be created.",
						},
						"version": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Template version.",
						},
						"committed": {
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Template version committed status.",
						},
						"action_control": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The action control properties that are created in an action resource when the template is assigned.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"service_name": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The service name that the action control refers.",
									},
									"description": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Description of the action control.",
									},
									"actions": &schema.Schema{
										Type:        schema.TypeSet,
										Computed:    true,
										Description: "List of actions to control access.",
										Elem:        &schema.Schema{Type: schema.TypeString},
									},
								},
							},
						},
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The action control template ID.",
						},
					},
				},
			},
		},
	}
}

func dataSourceIBMActionControlTemplateRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	iamPolicyManagementClient, err := meta.(conns.ClientSession).IAMPolicyManagementV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_iam_action_control_template", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	userDetails, err := meta.(conns.ClientSession).BluemixUserDetails()
	if err != nil {
		return diag.FromErr(fmt.Errorf("Failed to fetch BluemixUserDetails %s", err))
	}

	accountID := userDetails.UserAccount

	listActionControlTemplatesOptions := &iampolicymanagementv1.ListActionControlTemplatesOptions{}

	listActionControlTemplatesOptions.SetAccountID(accountID)

	var pager *iampolicymanagementv1.ActionControlTemplatesPager
	pager, err = iamPolicyManagementClient.NewActionControlTemplatesPager(listActionControlTemplatesOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, err.Error(), "(Data) ibm_iam_action_control_template", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	allItems, err := pager.GetAll()
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("ActionControlTemplatesPager.GetAll() failed %s", err), "(Data) ibm_iam_action_control_template", "read")
		log.Printf("[DEBUG] %s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(accountID)

	mapSlice := []map[string]interface{}{}
	for _, modelItem := range allItems {
		modelMap, err := DataSourceIBMListActionControlTemplatesActionControlTemplateToMap(&modelItem)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_iam_action_control_template", "read", "ActionControlTemplates-to-map").GetDiag()
		}
		mapSlice = append(mapSlice, modelMap)
	}

	if err = d.Set("action_control_templates", mapSlice); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting action_control_templates %s", err), "(Data) ibm_iam_action_control_template", "read", "action_control_templates-set").GetDiag()
	}

	return nil
}

func DataSourceIBMListActionControlTemplatesActionControlTemplateToMap(model *iampolicymanagementv1.ActionControlTemplate) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["name"] = *model.Name
	if model.Description != nil {
		modelMap["description"] = *model.Description
	}
	modelMap["account_id"] = *model.AccountID
	if model.Committed != nil {
		modelMap["committed"] = *model.Committed
	}
	if model.ActionControl != nil {
		actionControlMap, err := DataSourceIBMListActionControlTemplatesTemplateActionControlToMap(model.ActionControl)
		if err != nil {
			return modelMap, err
		}
		modelMap["action_control"] = []map[string]interface{}{actionControlMap}
	}
	if model.ID != nil {
		modelMap["id"] = *model.ID
	}
	modelMap["version"] = *model.Version
	return modelMap, nil
}

func DataSourceIBMListActionControlTemplatesTemplateActionControlToMap(model *iampolicymanagementv1.TemplateActionControl) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["service_name"] = *model.ServiceName
	if model.Description != nil {
		modelMap["description"] = *model.Description
	}
	modelMap["actions"] = model.Actions
	return modelMap, nil
}
