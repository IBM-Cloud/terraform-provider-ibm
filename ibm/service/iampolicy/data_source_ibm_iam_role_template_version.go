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

func DataSourceIBMIAMRoleTemplateVersion() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMListRoleTemplatesRead,

		Schema: map[string]*schema.Schema{
			"role_template_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "role template ID.",
			},
			"version": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "role template version.",
			},
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Required field when creating a new template. Otherwise, this field is optional. If the field is included, it changes the name value for all existing versions of the template.",
			},
			"description": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Description of the role  template. This is shown to users in the enterprise account. Use this to describe the purpose or context of the role  for enterprise users managing IAM templates.",
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
			"role": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The role  properties that are created in an role resource when the template is assigned.",
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
							Description: "The service name that the role  refers.",
						},
						"description": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Description of the role .",
						},
						"actions": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "List of actions to  access.",
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
				Description: "The href URL that links to the role templates API by role template ID.",
			},
			"created_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The UTC timestamp when the role template was created.",
			},
			"created_by_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The IAM ID of the entity that created the role template.",
			},
			"last_modified_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The UTC timestamp when the role template was last modified.",
			},
			"last_modified_by_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The IAM ID of the entity that last modified the role template.",
			},
			"state": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "State of role template.",
			},
		},
	}
}

func dataSourceIBMListRoleTemplatesRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	iamPolicyManagementClient, err := meta.(conns.ClientSession).IAMPolicyManagementV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_list_role_templates", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getRoleTemplateVersionOptions := &iampolicymanagementv1.GetRoleTemplateVersionOptions{}

	getRoleTemplateVersionOptions.SetRoleTemplateID(d.Get("role_template_id").(string))
	getRoleTemplateVersionOptions.SetVersion(d.Get("version").(string))

	roleTemplate, _, err := iamPolicyManagementClient.GetRoleTemplateVersionWithContext(context, getRoleTemplateVersionOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetRoleTemplateVersionWithContext failed: %s", err.Error()), "(Data) ibm_list_role_templates", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(*roleTemplate.ID)

	if err = d.Set("name", roleTemplate.Name); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting name: %s", err), "(Data) ibm_list_role_templates", "read", "set-name").GetDiag()
	}

	if err = d.Set("description", roleTemplate.Description); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting description: %s", err), "(Data) ibm_list_role_templates", "read", "set-description").GetDiag()
	}

	if err = d.Set("account_id", roleTemplate.AccountID); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting account_id: %s", err), "(Data) ibm_list_role_templates", "read", "set-account_id").GetDiag()
	}

	if !core.IsNil(roleTemplate.Committed) {
		if err = d.Set("committed", roleTemplate.Committed); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting committed: %s", err), "(Data) ibm_list_role_templates", "read", "set-committed").GetDiag()
		}
	}

	if !core.IsNil(roleTemplate.Role) {
		role := []map[string]interface{}{}
		roleMap, err := ResourceIBMIAMRoleTemplateVersionTemplateRoleToMap(roleTemplate.Role)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_list_role_templates", "read", "role-to-map").GetDiag()
		}
		role = append(role, roleMap)
		if err = d.Set("role", role); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting role: %s", err), "(Data) ibm_list_role_templates", "read", "set-role").GetDiag()
		}
	}

	return nil
}
