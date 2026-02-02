// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.107.1-41b0fbd0-20250825-080732
 */

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

func ResourceIBMIAMRoleTemplateVersion() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMIAMRoleTemplateVersionCreate,
		ReadContext:   resourceIBMIAMRoleTemplateVersionRead,
		UpdateContext: resourceIBMIAMRoleTemplateVersionUpdate,
		DeleteContext: resourceIBMIAMRoleTemplateVersionDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"role_template_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The role template ID.",
			},
			"name": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Required field when creating a new template. Otherwise, this field is optional. If the field is included, it changes the name value for all existing versions of the template.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Computed:    true,
				Description: "Description of the role template. This is shown to users in the enterprise account. Use this to describe the purpose or context of the role for enterprise users managing IAM templates.",
			},
			"committed": {
				Type:        schema.TypeBool,
				Optional:    true,
				Description: "Committed status of the template. If committed is set to true, then the template version can no longer be updated.",
			},
			"role": {
				Type:        schema.TypeList,
				MinItems:    1,
				MaxItems:    1,
				Optional:    true,
				Description: "The role properties that are created in an action resource when the template is assigned.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the role that is used in the CRN. This must be alphanumeric and capitalized.",
						},
						"display_name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The display the name of the role that is shown in the console.",
						},
						"service_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The service name that the role refers.",
						},
						"description": {
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Description of the role.",
						},
						"actions": {
							Type:        schema.TypeList,
							Required:    true,
							Description: "The actions of the role.",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
					},
				},
			},
			"account_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Enterprise account ID where this template is created.",
			},
			"id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The role template ID.",
			},
			"href": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The href URL that links to the role templates API by role template ID.",
			},
			"created_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The UTC timestamp when the role template was created.",
			},
			"created_by_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The IAM ID of the entity that created the role template.",
			},
			"last_modified_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The UTC timestamp when the role template was last modified.",
			},
			"last_modified_by_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The IAM ID of the entity that last modified the role template.",
			},
			"state": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "State of role template.",
			},
			"version": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The version number of the template used to identify different versions of same template.",
			},
		},
	}
}

func resourceIBMIAMRoleTemplateVersionCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	iamPolicyManagementClient, err := meta.(conns.ClientSession).IAMPolicyManagementV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_role_template_version", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	createRoleTemplateVersionOptions := &iampolicymanagementv1.CreateRoleTemplateVersionOptions{}

	createRoleTemplateVersionOptions.SetRoleTemplateID(d.Get("role_template_id").(string))

	if _, ok := d.GetOk("role"); ok {
		roleModel, err := ResourceIBMIAMRoleTemplateMapToTemplateVersionRole(d.Get("role.0").(map[string]interface{}))
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_role_template_version", "create", "parse-role").GetDiag()
		}
		createRoleTemplateVersionOptions.SetRole(roleModel)
	}

	if _, ok := d.GetOk("name"); ok {
		createRoleTemplateVersionOptions.SetName(d.Get("name").(string))
	}
	if _, ok := d.GetOk("description"); ok {
		createRoleTemplateVersionOptions.SetDescription(d.Get("description").(string))
	}
	if _, ok := d.GetOk("committed"); ok {
		createRoleTemplateVersionOptions.SetCommitted(d.Get("committed").(bool))
	}

	roleTemplateVersion, _, err := iamPolicyManagementClient.CreateRoleTemplateVersionWithContext(context, createRoleTemplateVersionOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreateRoleTemplateVersionWithContext failed: %s", err.Error()), "ibm_iam_role_template_version", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(fmt.Sprintf("%s/%s", *createRoleTemplateVersionOptions.RoleTemplateID, *roleTemplateVersion.Version))

	return resourceIBMIAMRoleTemplateVersionRead(context, d, meta)
}

func resourceIBMIAMRoleTemplateVersionRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	iamPolicyManagementClient, err := meta.(conns.ClientSession).IAMPolicyManagementV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_role_template_version", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getRoleTemplateVersionOptions := &iampolicymanagementv1.GetRoleTemplateVersionOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_role_template_version", "read", "sep-id-parts").GetDiag()
	}

	getRoleTemplateVersionOptions.SetRoleTemplateID(parts[0])
	getRoleTemplateVersionOptions.SetVersion(parts[1])

	roleTemplate, response, err := iamPolicyManagementClient.GetRoleTemplateVersionWithContext(context, getRoleTemplateVersionOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetRoleTemplateVersionWithContext failed: %s", err.Error()), "ibm_iam_role_template_version", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	if !core.IsNil(roleTemplate.ID) {
		if err = d.Set("role_template_id", roleTemplate.ID); err != nil {
			err = fmt.Errorf("Error setting role_template_id: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_role_template_version", "read", "set-role_template_id").GetDiag()
		}
	}

	if !core.IsNil(roleTemplate.Name) {
		if err = d.Set("name", roleTemplate.Name); err != nil {
			err = fmt.Errorf("Error setting name: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_role_template_version", "read", "set-name").GetDiag()
		}
	}
	if !core.IsNil(roleTemplate.Description) {
		if err = d.Set("description", roleTemplate.Description); err != nil {
			err = fmt.Errorf("Error setting description: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_role_template_version", "read", "set-description").GetDiag()
		}
	}
	if !core.IsNil(roleTemplate.Committed) {
		if err = d.Set("committed", roleTemplate.Committed); err != nil {
			err = fmt.Errorf("Error setting committed: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_role_template_version", "read", "set-committed").GetDiag()
		}
	}

	if _, ok := d.GetOk("role"); ok {
		roleMap, err := ResourceIBMIAMRoleTemplateVersionTemplateRoleToMap(roleTemplate.Role)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_role_template_version", "read", "role-to-map").GetDiag()
		}
		if err = d.Set("role", []map[string]interface{}{roleMap}); err != nil {
			err = fmt.Errorf("Error setting role: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_role_template_version", "read", "set-role").GetDiag()
		}
	}

	if err = d.Set("account_id", roleTemplate.AccountID); err != nil {
		err = fmt.Errorf("Error setting account_id: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_role_template_version", "read", "set-account_id").GetDiag()
	}
	if !core.IsNil(roleTemplate.Href) {
		if err = d.Set("href", roleTemplate.Href); err != nil {
			err = fmt.Errorf("Error setting href: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_role_template_version", "read", "set-href").GetDiag()
		}
	}
	if !core.IsNil(roleTemplate.CreatedAt) {
		if err = d.Set("created_at", flex.DateTimeToString(roleTemplate.CreatedAt)); err != nil {
			err = fmt.Errorf("Error setting created_at: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_role_template_version", "read", "set-created_at").GetDiag()
		}
	}
	if !core.IsNil(roleTemplate.CreatedByID) {
		if err = d.Set("created_by_id", roleTemplate.CreatedByID); err != nil {
			err = fmt.Errorf("Error setting created_by_id: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_role_template_version", "read", "set-created_by_id").GetDiag()
		}
	}
	if !core.IsNil(roleTemplate.LastModifiedAt) {
		if err = d.Set("last_modified_at", flex.DateTimeToString(roleTemplate.LastModifiedAt)); err != nil {
			err = fmt.Errorf("Error setting last_modified_at: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_role_template_version", "read", "set-last_modified_at").GetDiag()
		}
	}
	if !core.IsNil(roleTemplate.LastModifiedByID) {
		if err = d.Set("last_modified_by_id", roleTemplate.LastModifiedByID); err != nil {
			err = fmt.Errorf("Error setting last_modified_by_id: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_role_template_version", "read", "set-last_modified_by_id").GetDiag()
		}
	}
	if err = d.Set("state", roleTemplate.State); err != nil {
		err = fmt.Errorf("Error setting state: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_role_template_version", "read", "set-state").GetDiag()
	}

	if !core.IsNil(roleTemplate.Version) {
		if err = d.Set("version", roleTemplate.Version); err != nil {
			err = fmt.Errorf("Error setting version: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_role_template_version", "read", "set-version").GetDiag()
		}
	}

	return nil
}

func resourceIBMIAMRoleTemplateVersionUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	if d.HasChange("name") || d.HasChange("description") || d.HasChange("committed") || d.HasChange("role") {
		iamPolicyManagementClient, err := meta.(conns.ClientSession).IAMPolicyManagementV1API()
		if err != nil {
			tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_role_template_version", "update", "initialize-client")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}

		parts, err := flex.SepIdParts(d.Id(), "/")
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_role_template_version", "update", "sep-id-parts").GetDiag()
		}

		getRoleTemplateVersionOptions := &iampolicymanagementv1.GetRoleTemplateVersionOptions{
			RoleTemplateID: &parts[0],
			Version:        &parts[1],
		}

		roleTemplate, response, err := iamPolicyManagementClient.GetRoleTemplateVersionWithContext(context, getRoleTemplateVersionOptions)

		if err != nil || roleTemplate == nil {
			if response != nil && response.StatusCode == 404 {
				return nil
			}
			return diag.FromErr(fmt.Errorf("[ERROR] Error retrieving Policy Template: %s\n%s", err, response))
		}

		replaceRoleTemplateOptions := &iampolicymanagementv1.ReplaceRoleTemplateOptions{}

		replaceRoleTemplateOptions.SetRoleTemplateID(parts[0])
		replaceRoleTemplateOptions.SetVersion(parts[1])
		replaceRoleTemplateOptions.SetIfMatch(response.Headers.Get("ETag"))
		if _, ok := d.GetOk("name"); ok {
			replaceRoleTemplateOptions.SetName(d.Get("name").(string))
		}
		if _, ok := d.GetOk("description"); ok {
			replaceRoleTemplateOptions.SetDescription(d.Get("description").(string))
		}
		if _, ok := d.GetOk("role"); ok {
			role, err := ResourceIBMIAMRoleTemplateMapToTemplateVersionRole(d.Get("role.0").(map[string]interface{}))
			if err != nil {
				return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_role_template_version", "update", "parse-role").GetDiag()
			}
			replaceRoleTemplateOptions.SetRole(role)
		}
		if _, ok := d.GetOk("committed"); ok {
			replaceRoleTemplateOptions.SetCommitted(d.Get("committed").(bool))
		}
		_, _, err = iamPolicyManagementClient.ReplaceRoleTemplateWithContext(context, replaceRoleTemplateOptions)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, fmt.Sprintf("ReplaceRoleTemplateWithContext failed: %s", err.Error()), "ibm_role_template_version", "update")
			log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
			return tfErr.GetDiag()
		}
	}
	return resourceIBMIAMRoleTemplateVersionRead(context, d, meta)
}

func resourceIBMIAMRoleTemplateVersionDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	iamPolicyManagementClient, err := meta.(conns.ClientSession).IAMPolicyManagementV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_role_template_version", "delete", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	deleteRoleTemplateVersionOptions := &iampolicymanagementv1.DeleteRoleTemplateVersionOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_role_template_version", "delete", "sep-id-parts").GetDiag()
	}

	deleteRoleTemplateVersionOptions.SetRoleTemplateID(parts[0])
	deleteRoleTemplateVersionOptions.SetVersion(parts[1])

	_, err = iamPolicyManagementClient.DeleteRoleTemplateVersionWithContext(context, deleteRoleTemplateVersionOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("DeleteRoleTemplateVersionWithContext failed: %s", err.Error()), "ibm_iam_role_template_version", "delete")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId("")

	return nil
}

func ResourceIBMIAMRoleTemplateMapToTemplateVersionRole(modelMap map[string]interface{}) (*iampolicymanagementv1.TemplateRole, error) {
	model := &iampolicymanagementv1.TemplateRole{}
	if v, ok := modelMap["display_name"]; ok && v != nil {
		if str, ok := v.(string); ok && str != "" {
			model.DisplayName = core.StringPtr(str)
		}
	}
	if v, ok := modelMap["description"]; ok && v != nil {
		if str, ok := v.(string); ok && str != "" {
			model.Description = core.StringPtr(str)
		}
	}
	if v, ok := modelMap["actions"]; ok && v != nil {
		if list, ok := v.([]interface{}); ok {
			actions := make([]string, 0, len(list))
			for _, item := range list {
				if s, ok := item.(string); ok && s != "" {
					actions = append(actions, s)
				}
			}
			if len(actions) > 0 {
				model.Actions = actions
			}
		}
	}
	return model, nil
}

func ResourceIBMIAMRoleTemplateVersionTemplateRoleToMap(model *iampolicymanagementv1.RoleTemplatePrototypeRole) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Name != nil {
		modelMap["name"] = *model.Name
	}

	if model.DisplayName != nil {
		modelMap["display_name"] = *model.DisplayName
	}

	if model.ServiceName != nil {
		modelMap["service_name"] = *model.ServiceName
	}

	if model.Description != nil {
		modelMap["description"] = *model.Description
	}

	if model.Actions != nil {
		modelMap["actions"] = model.Actions
	}

	return modelMap, nil
}
