// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package iampolicy

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/platform-services-go-sdk/iampolicymanagementv1"
)

func DataSourceIBMIAMRoleTemplate() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMRoleTemplateRead,

		Schema: map[string]*schema.Schema{
			"role_templates": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of role templates.",
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
						"role": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The role properties that are created in an action resource when the template is assigned.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The name of the role that is used in the CRN. This must be alphanumeric and capitalized.",
									},
									"display_name": &schema.Schema{
										Type:        schema.TypeString,
										Required:    true,
										Description: "The display the name of the role that is shown in the console.",
									},
									"service_name": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The service name that the role refers.",
									},
									"description": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Description of the role.",
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
							Description: "The role template ID.",
						},
					},
				},
			},
		},
	}
}

func dataSourceIBMRoleTemplateRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	iamPolicyManagementClient, err := meta.(conns.ClientSession).IAMPolicyManagementV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_iam_role_template", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	userDetails, err := meta.(conns.ClientSession).BluemixUserDetails()
	if err != nil {
		return diag.FromErr(fmt.Errorf("Failed to fetch BluemixUserDetails %s", err))
	}

	accountID := userDetails.UserAccount

	listRoleTemplatesOptions := &iampolicymanagementv1.ListRoleTemplatesOptions{}

	listRoleTemplatesOptions.SetAccountID(accountID)

	var pager *iampolicymanagementv1.RoleTemplatesPager
	pager, err = iamPolicyManagementClient.NewRoleTemplatesPager(listRoleTemplatesOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, err.Error(), "(Data) ibm_iam_role_template", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	allItems, err := pager.GetAll()
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("RoleTemplatesPager.GetAll() failed %s", err), "(Data) ibm_iam_role_template", "read")
		log.Printf("[DEBUG] %s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(accountID)

	mapSlice := []map[string]interface{}{}
	for _, modelItem := range allItems {
		modelMap, err := DataSourceIBMListRoleTemplatesRoleTemplateToMap(&modelItem)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_iam_role_template", "read", "RoleTemplates-to-map").GetDiag()
		}
		mapSlice = append(mapSlice, modelMap)
	}

	if err = d.Set("role_templates", mapSlice); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting role_templates %s", err), "(Data) ibm_iam_role_template", "read", "role_templates-set").GetDiag()
	}

	return nil
}

func DataSourceIBMListRoleTemplatesRoleTemplateToMap(model *iampolicymanagementv1.RoleTemplate) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Name != nil {
		modelMap["name"] = *model.Name
	}
	if model.Description != nil {
		modelMap["description"] = *model.Description
	}
	modelMap["account_id"] = *model.AccountID
	if model.Committed != nil {
		modelMap["committed"] = *model.Committed
	}
	if model.Role != nil {
		roleMap, err := ResourceIBMIAMRoleTemplateVersionTemplateRoleToMap(model.Role)
		if err != nil {
			return modelMap, err
		}
		modelMap["role"] = []map[string]interface{}{roleMap}
	}
	if model.ID != nil {
		modelMap["id"] = *model.ID
	}
	modelMap["version"] = *model.Version
	return modelMap, nil
}
