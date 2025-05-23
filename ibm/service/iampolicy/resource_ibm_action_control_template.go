// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.103.0-e8b84313-20250402-201816
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

func ResourceIBMIAMActionControlTemplate() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIBMIAMActionControlTemplateCreate,
		ReadContext:   resourceIBMIAMActionControlTemplateVersionRead,
		UpdateContext: resourceIBMIAMActionControlTemplateVersionUpdate,
		DeleteContext: resourceIBMIAMActionControlTemplateVersionDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"action_control_template_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The action control template ID.",
			},
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Required field when creating a new template. Otherwise, this field is optional. If the field is included, it changes the name value for all existing versions of the template.",
			},
			"description": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Description of the action control template. This is shown to users in the enterprise account. Use this to describe the purpose or context of the action control for enterprise users managing IAM templates.",
			},
			"committed": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Computed:    true,
				Description: "Committed status of the template. If committed is set to true, then the template version can no longer be updated.",
			},
			"action_control": &schema.Schema{
				Type:        schema.TypeList,
				MinItems:    1,
				Optional:    true,
				Description: "The action control properties that are created in an action resource when the template is assigned.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"service_name": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "The service name that the action control refers.",
						},
						"description": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Description of the action control.",
						},
						"actions": &schema.Schema{
							Type:        schema.TypeSet,
							MinItems:    1,
							Required:    true,
							Description: "List of actions to control access.",
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
			"version": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Template Version.",
			},
		},
	}
}

func resourceIBMIAMActionControlTemplateCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	iamPolicyManagementClient, err := meta.(conns.ClientSession).IAMPolicyManagementV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_action_control_template", "create", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	userDetails, err := meta.(conns.ClientSession).BluemixUserDetails()
	if err != nil {
		return diag.FromErr(fmt.Errorf("failed to fetch BluemixUserDetails %s", err))
	}

	accountID := userDetails.UserAccount

	createActionControlTemplateOptions := &iampolicymanagementv1.CreateActionControlTemplateOptions{}

	createActionControlTemplateOptions.SetAccountID(accountID)

	if _, ok := d.GetOk("name"); ok {
		createActionControlTemplateOptions.SetName(d.Get("name").(string))
	}
	if _, ok := d.GetOk("description"); ok {
		createActionControlTemplateOptions.SetDescription(d.Get("description").(string))
	}
	if _, ok := d.GetOk("action_control"); ok {
		actionControlModel, err := generateTemplateActionControl(d.Get("action_control.0").(map[string]interface{}))
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_action_control_template", "create", "parse-action_control").GetDiag()
		}
		createActionControlTemplateOptions.SetActionControl(actionControlModel)
	}
	if _, ok := d.GetOk("committed"); ok {
		createActionControlTemplateOptions.SetCommitted(d.Get("committed").(bool))
	}

	actionControlTemplate, _, err := iamPolicyManagementClient.CreateActionControlTemplateWithContext(context, createActionControlTemplateOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("CreateActionControlTemplateWithContext failed: %s", err.Error()), "ibm_iam_action_control_template", "create")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(fmt.Sprintf("%s/%s", *actionControlTemplate.ID, *actionControlTemplate.Version))

	return resourceIBMIAMActionControlTemplateVersionRead(context, d, meta)
}

func resourceIBMIAMActionControlTemplateRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	iamPolicyManagementClient, err := meta.(conns.ClientSession).IAMPolicyManagementV1API()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_action_control_template", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getActionControlTemplateVersionOptions := &iampolicymanagementv1.GetActionControlTemplateVersionOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_action_control_template", "read", "sep-id-parts").GetDiag()
	}

	getActionControlTemplateVersionOptions.SetActionControlTemplateID(parts[0])
	getActionControlTemplateVersionOptions.SetVersion(parts[1])

	actionControlTemplate, response, err := iamPolicyManagementClient.GetActionControlTemplateVersionWithContext(context, getActionControlTemplateVersionOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetActionControlTemplateVersionWithContext failed: %s", err.Error()), "ibm_iam_action_control_template", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	if !core.IsNil(actionControlTemplate.Name) {
		if err = d.Set("name", actionControlTemplate.Name); err != nil {
			err = fmt.Errorf("Error setting name: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_action_control_template", "read", "set-name").GetDiag()
		}
	}
	if !core.IsNil(actionControlTemplate.Description) {
		if err = d.Set("description", actionControlTemplate.Description); err != nil {
			err = fmt.Errorf("Error setting description: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_action_control_template", "read", "set-description").GetDiag()
		}
	}
	if !core.IsNil(actionControlTemplate.Committed) {
		if err = d.Set("committed", actionControlTemplate.Committed); err != nil {
			err = fmt.Errorf("Error setting committed: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_action_control_template", "read", "set-committed").GetDiag()
		}
	}
	if !core.IsNil(actionControlTemplate.ActionControl) {
		actionControlMap, err := flattenTemplateActionControl(actionControlTemplate.ActionControl)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_action_control_template", "read", "action_control-to-map").GetDiag()
		}
		if err = d.Set("action_control", []map[string]interface{}{actionControlMap}); err != nil {
			err = fmt.Errorf("Error setting action_control: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_action_control_template", "read", "set-action_control").GetDiag()
		}
	}
	if err = d.Set("account_id", actionControlTemplate.AccountID); err != nil {
		err = fmt.Errorf("Error setting account_id: %s", err)
		return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_action_control_template", "read", "set-account_id").GetDiag()
	}
	if !core.IsNil(actionControlTemplate.ID) {
		if err = d.Set("id", actionControlTemplate.ID); err != nil {
			err = fmt.Errorf("Error setting id: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_action_control_template", "read", "set-id").GetDiag()
		}
	}
	if !core.IsNil(actionControlTemplate.Version) {
		if err = d.Set("version", actionControlTemplate.Version); err != nil {
			err = fmt.Errorf("Error setting version: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_action_control_template", "read", "set-version").GetDiag()
		}
	}
	if !core.IsNil(actionControlTemplate.Href) {
		if err = d.Set("href", actionControlTemplate.Href); err != nil {
			err = fmt.Errorf("Error setting href: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_action_control_template", "read", "set-href").GetDiag()
		}
	}
	if !core.IsNil(actionControlTemplate.CreatedAt) {
		if err = d.Set("created_at", flex.DateTimeToString(actionControlTemplate.CreatedAt)); err != nil {
			err = fmt.Errorf("Error setting created_at: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_action_control_template", "read", "set-created_at").GetDiag()
		}
	}
	if !core.IsNil(actionControlTemplate.CreatedByID) {
		if err = d.Set("created_by_id", actionControlTemplate.CreatedByID); err != nil {
			err = fmt.Errorf("Error setting created_by_id: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_action_control_template", "read", "set-created_by_id").GetDiag()
		}
	}
	if !core.IsNil(actionControlTemplate.LastModifiedAt) {
		if err = d.Set("last_modified_at", flex.DateTimeToString(actionControlTemplate.LastModifiedAt)); err != nil {
			err = fmt.Errorf("Error setting last_modified_at: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_action_control_template", "read", "set-last_modified_at").GetDiag()
		}
	}
	if !core.IsNil(actionControlTemplate.LastModifiedByID) {
		if err = d.Set("last_modified_by_id", actionControlTemplate.LastModifiedByID); err != nil {
			err = fmt.Errorf("Error setting last_modified_by_id: %s", err)
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "ibm_iam_action_control_template", "read", "set-last_modified_by_id").GetDiag()
		}
	}

	return nil
}

func generateTemplateActionControl(modelMap map[string]interface{}) (*iampolicymanagementv1.TemplateActionControl, error) {
	model := &iampolicymanagementv1.TemplateActionControl{}
	model.ServiceName = core.StringPtr(modelMap["service_name"].(string))
	if modelMap["description"] != nil && modelMap["description"].(string) != "" {
		model.Description = core.StringPtr(modelMap["description"].(string))
	}
	actions := []string{}
	for _, actionsItem := range modelMap["actions"].(*schema.Set).List() {
		actions = append(actions, actionsItem.(string))
	}
	model.Actions = actions
	return model, nil
}

func flattenTemplateActionControl(model *iampolicymanagementv1.TemplateActionControl) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["service_name"] = *model.ServiceName
	if model.Description != nil {
		modelMap["description"] = *model.Description
	}
	modelMap["actions"] = model.Actions
	return modelMap, nil
}
