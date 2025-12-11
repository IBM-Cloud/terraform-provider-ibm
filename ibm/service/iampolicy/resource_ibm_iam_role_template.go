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

func ResourceIBMIAMRoleTemplate() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMIAMRoleTemplateCreate,
		ReadContext:   resourceIBMIAMRoleTemplateVersionRead,
		UpdateContext: resourceIBMIAMRoleTemplateVersionUpdate,
		DeleteContext: resourceIBMIAMRoleTemplateVersionDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"role_template_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The role template ID.",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Required field when creating a new template. Otherwise, this field is optional. If the field is included, it changes the name value for all existing versions of the template.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Description of the role template. This is shown to users in the enterprise account. Use this to describe the purpose or context of the role for enterprise users managing IAM templates.",
			},
			"committed": {
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				Description: "Committed status of the template. If committed is set to true, then the template version can no longer be updated.",
			},
			"role": {
				Type:        schema.TypeList,
				MaxItems:    1,
				Optional:    true,
				ForceNew:    true,
				Description: "The role properties that are created in an action resource when the template is assigned.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The name of the role that is used in the CRN. This must be alphanumeric and capitalized.",
						},
						"display_name": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "The display the name of the role that is shown in the console.",
						},
						"service_name": {
							Type:        schema.TypeString,
							Required:    true,
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
			"account_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Enterprise account ID where this template is created.",
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
			"version": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The version number of the template used to identify different versions of same template.",
			},
			"state": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "State of role template.",
			},
		},
	}
}

func resourceIBMIAMRoleTemplateCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	iamPolicyManagementClient, err := meta.(conns.ClientSession).IAMPolicyManagementV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_role_template", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	userDetails, err := meta.(conns.ClientSession).BluemixUserDetails()
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to fetch BluemixUserDetails %s", err))
	}

	accountID := userDetails.UserAccount
	createRoleTemplateOptions := &iampolicymanagementv1.CreateRoleTemplateOptions{}

	createRoleTemplateOptions.SetAccountID(accountID)
	createRoleTemplateOptions.SetName(d.Get("name").(string))
	if _, ok := d.GetOk("description"); ok {
		createRoleTemplateOptions.SetDescription(d.Get("description").(string))
	}
	if _, ok := d.GetOk("committed"); ok {
		createRoleTemplateOptions.SetCommitted(d.Get("committed").(bool))
	}
	if _, ok := d.GetOk("role"); ok {
		roleModel, err := ResourceIBMIAMRoleTemplateMapToTemplateRole(d.Get("role.0").(map[string]interface{}))
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_role_template", "create", "parse-role").GetDiag()
		}
		createRoleTemplateOptions.SetRole(roleModel)
	}
	roleTemplate, _, err := iamPolicyManagementClient.CreateRoleTemplateWithContext(context, createRoleTemplateOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreateRoleTemplateWithContext failed: %s", err.Error()), "ibm_iam_role_template", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(fmt.Sprintf("%s/%s", *roleTemplate.ID, *roleTemplate.Version))

	return resourceIBMIAMRoleTemplateVersionRead(context, d, meta)
}

func ResourceIBMIAMRoleTemplateMapToTemplateRole(modelMap map[string]interface{}) (*iampolicymanagementv1.RoleTemplatePrototypeRole, error) {
	model := &iampolicymanagementv1.RoleTemplatePrototypeRole{}
	model.Name = core.StringPtr(modelMap["name"].(string))
	model.DisplayName = core.StringPtr(modelMap["display_name"].(string))
	model.ServiceName = core.StringPtr(modelMap["service_name"].(string))
	if modelMap["description"] != nil && modelMap["description"].(string) != "" {
		model.Description = core.StringPtr(modelMap["description"].(string))
	}
	actions := []string{}
	for _, actionsItem := range modelMap["actions"].([]interface{}) {
		actions = append(actions, actionsItem.(string))
	}
	model.Actions = actions
	return model, nil
}
