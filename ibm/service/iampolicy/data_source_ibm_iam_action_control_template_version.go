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
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/iampolicymanagementv1"
)

func DataSourceIBMIAMActionControlTemplateVersion() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMListActionControlTemplatesRead,

		Schema: map[string]*schema.Schema{
			"action_control_template_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Action control template ID.",
			},
			"version": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Action control template version.",
			},
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Required field when creating a new template. Otherwise, this field is optional. If the field is included, it changes the name value for all existing versions of the template.",
			},
			"description": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Description of the action control template. This is shown to users in the enterprise account. Use this to describe the purpose or context of the action control for enterprise users managing IAM templates.",
			},
			"account_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Enterprise account ID where this template is created.",
			},
			"committed": &schema.Schema{
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "Committed status of the template. If committed is set to true, then the template version can no longer be updated.",
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
							Computed:    true,
							Description: "Description of the action control.",
						},
						"actions": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "List of actions to control access.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"href": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The href URL that links to the action control templates API by action control template ID.",
			},
			"created_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The UTC timestamp when the action control template was created.",
			},
			"created_by_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The IAM ID of the entity that created the action control template.",
			},
			"last_modified_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The UTC timestamp when the action control template was last modified.",
			},
			"last_modified_by_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The IAM ID of the entity that last modified the action control template.",
			},
			"state": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "State of action control template.",
			},
		},
	}
}

func dataSourceIBMListActionControlTemplatesRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	iamPolicyManagementClient, err := meta.(conns.ClientSession).IAMPolicyManagementV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_list_action_control_templates", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getActionControlTemplateVersionOptions := &iampolicymanagementv1.GetActionControlTemplateVersionOptions{}

	getActionControlTemplateVersionOptions.SetActionControlTemplateID(d.Get("action_control_template_id").(string))
	getActionControlTemplateVersionOptions.SetVersion(d.Get("version").(string))

	actionControlTemplate, _, err := iamPolicyManagementClient.GetActionControlTemplateVersionWithContext(context, getActionControlTemplateVersionOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetActionControlTemplateVersionWithContext failed: %s", err.Error()), "(Data) ibm_list_action_control_templates", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(*actionControlTemplate.ID)

	if err = d.Set("name", actionControlTemplate.Name); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting name: %s", err), "(Data) ibm_list_action_control_templates", "read", "set-name").GetDiag()
	}

	if err = d.Set("description", actionControlTemplate.Description); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting description: %s", err), "(Data) ibm_list_action_control_templates", "read", "set-description").GetDiag()
	}

	if err = d.Set("account_id", actionControlTemplate.AccountID); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting account_id: %s", err), "(Data) ibm_list_action_control_templates", "read", "set-account_id").GetDiag()
	}

	if !core.IsNil(actionControlTemplate.Committed) {
		if err = d.Set("committed", actionControlTemplate.Committed); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting committed: %s", err), "(Data) ibm_list_action_control_templates", "read", "set-committed").GetDiag()
		}
	}

	if !core.IsNil(actionControlTemplate.ActionControl) {
		actionControl := []map[string]interface{}{}
		actionControlMap, err := DataSourceIBMListActionControlTemplatesTemplateActionControlToMap(actionControlTemplate.ActionControl)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_list_action_control_templates", "read", "action_control-to-map").GetDiag()
		}
		actionControl = append(actionControl, actionControlMap)
		if err = d.Set("action_control", actionControl); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting action_control: %s", err), "(Data) ibm_list_action_control_templates", "read", "set-action_control").GetDiag()
		}
	}

	return nil
}
