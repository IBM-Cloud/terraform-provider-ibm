// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package project

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/project-go-sdk/projectv1"
)

func ResourceIbmProject() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIbmProjectCreate,
		ReadContext:   resourceIbmProjectRead,
		UpdateContext: resourceIbmProjectUpdate,
		DeleteContext: resourceIbmProjectDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"location": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validate.InvokeValidator("ibm_project", "location"),
				Description:  "The IBM Cloud location where a resource is deployed.",
			},
			"resource_group": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				// ValidateFunc: validate.InvokeValidator("ibm_project", "resource_group"),
				Description: "The resource group where the project's data and tools are created.",
			},
			"definition": &schema.Schema{
				Type:        schema.TypeList,
				MinItems:    1,
				MaxItems:    1,
				Required:    true,
				Description: "The definition of the project.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "The name of the project.",
						},
						"description": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "A brief explanation of the project's use in the configuration of a deployable architecture. It is possible to create a project without providing a description.",
						},
						"destroy_on_delete": &schema.Schema{
							Type:        schema.TypeBool,
							Required:    true,
							Description: "The policy that indicates whether the resources are destroyed or not when a project is deleted.",
						},
					},
				},
			},
			"crn": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "An IBM Cloud resource name, which uniquely identifies a resource.",
			},
			"created_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "A date and time value in the format YYYY-MM-DDTHH:mm:ssZ or YYYY-MM-DDTHH:mm:ss.sssZ, matching the date and time format as specified by RFC 3339.",
			},
			"cumulative_needs_attention_view": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The cumulative list of needs attention items for a project. If the view is successfully retrieved, an array which could be empty is returned.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"event": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The event name.",
						},
						"event_id": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "A unique ID for that individual event.",
						},
						"config_id": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "A unique ID for the configuration.",
						},
						"config_version": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "The version number of the configuration.",
						},
					},
				},
			},
			"cumulative_needs_attention_view_error": &schema.Schema{
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "True indicates that the fetch of the needs attention items failed. It only exists if there was an error while retrieving the cumulative needs attention view.",
			},
			"state": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The project status value.",
			},
			"event_notifications_crn": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The CRN of the event notifications instance if one is connected to this project.",
			},
			"configs": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The project configurations. These configurations are only included in the response of creating a project if a configs array is specified in the request payload.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "The ID of the configuration. If this parameter is empty, an ID is automatically created for the configuration.",
						},
						"project_id": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "The unique ID.",
						},
						"version": &schema.Schema{
							Type:        schema.TypeInt,
							Required:    true,
							Description: "The version of the configuration.",
						},
						"is_draft": &schema.Schema{
							Type:        schema.TypeBool,
							Required:    true,
							Description: "The flag that indicates whether the version of the configuration is draft, or active.",
						},
						"needs_attention_state": &schema.Schema{
							Type:        schema.TypeList,
							Optional:    true,
							Description: "The needs attention state of a configuration.",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"state": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "The state of the configuration.",
						},
						"update_available": &schema.Schema{
							Type:        schema.TypeBool,
							Required:    true,
							Description: "The flag that indicates whether a configuration update is available.",
						},
						"created_at": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "A date and time value in the format YYYY-MM-DDTHH:mm:ssZ or YYYY-MM-DDTHH:mm:ss.sssZ, matching the date and time format as specified by RFC 3339.",
						},
						"updated_at": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "A date and time value in the format YYYY-MM-DDTHH:mm:ssZ or YYYY-MM-DDTHH:mm:ss.sssZ, matching the date and time format as specified by RFC 3339.",
						},
						"last_save": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "A date and time value in the format YYYY-MM-DDTHH:mm:ssZ or YYYY-MM-DDTHH:mm:ss.sssZ, matching the date and time format as specified by RFC 3339.",
						},
						"href": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "A relative URL.",
						},
						"definition": &schema.Schema{
							Type:        schema.TypeList,
							MinItems:    1,
							MaxItems:    1,
							Required:    true,
							Description: "The name and description of a project configuration.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The configuration name.",
									},
									"description": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "A project configuration description.",
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func ResourceIbmProjectValidator() *validate.ResourceValidator {
	validateSchema := make([]validate.ValidateSchema, 0)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "location",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Required:                   true,
			Regexp:                     `^$|^(us-south|us-east|eu-gb|eu-de)$`,
			MinValueLength:             0,
			MaxValueLength:             12,
		},
		validate.ValidateSchema{
			Identifier:                 "resource_group",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Required:                   true,
			Regexp:                     `^(?!\s)(?!.*\s$)[^'"` + "`" + `<>{}\x00-\x1F]*$`,
			MinValueLength:             0,
			MaxValueLength:             64,
		},
	)

	resourceValidator := validate.ResourceValidator{ResourceName: "ibm_project", Schema: validateSchema}
	return &resourceValidator
}

func resourceIbmProjectCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	projectClient, err := meta.(conns.ClientSession).ProjectV1()
	if err != nil {
		return diag.FromErr(err)
	}

	createProjectOptions := &projectv1.CreateProjectOptions{}

	definitionModel, err := resourceIbmProjectMapToProjectPrototypeDefinition(d.Get("definition.0").(map[string]interface{}))
	if err != nil {
		return diag.FromErr(err)
	}
	createProjectOptions.SetDefinition(definitionModel)
	createProjectOptions.SetLocation(d.Get("location").(string))
	createProjectOptions.SetResourceGroup(d.Get("resource_group").(string))
	if _, ok := d.GetOk("configs"); ok {
		var configs []projectv1.ProjectConfigPrototype
		for _, v := range d.Get("configs").([]interface{}) {
			value := v.(map[string]interface{})
			configsItem, err := resourceIbmProjectMapToProjectConfigPrototype(value)
			if err != nil {
				return diag.FromErr(err)
			}
			configs = append(configs, *configsItem)
		}
		createProjectOptions.SetConfigs(configs)
	}

	project, response, err := projectClient.CreateProjectWithContext(context, createProjectOptions)
	if err != nil {
		log.Printf("[DEBUG] CreateProjectWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("CreateProjectWithContext failed %s\n%s", err, response))
	}

	d.SetId(*project.ID)

	return resourceIbmProjectRead(context, d, meta)
}

func resourceIbmProjectRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	projectClient, err := meta.(conns.ClientSession).ProjectV1()
	if err != nil {
		return diag.FromErr(err)
	}

	getProjectOptions := &projectv1.GetProjectOptions{}

	getProjectOptions.SetID(d.Id())

	project, response, err := projectClient.GetProjectWithContext(context, getProjectOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		log.Printf("[DEBUG] GetProjectWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("GetProjectWithContext failed %s\n%s", err, response))
	}

	if err = d.Set("location", project.Location); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting location: %s", err))
	}
	if err = d.Set("resource_group", project.ResourceGroup); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting resource_group: %s", err))
	}
	definitionMap, err := resourceIbmProjectProjectDefinitionPropertiesToMap(project.Definition)
	if err != nil {
		return diag.FromErr(err)
	}
	if err = d.Set("definition", []map[string]interface{}{definitionMap}); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting definition: %s", err))
	}
	if !core.IsNil(project.Crn) {
		if err = d.Set("crn", project.Crn); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting crn: %s", err))
		}
	}
	if !core.IsNil(project.CreatedAt) {
		if err = d.Set("created_at", flex.DateTimeToString(project.CreatedAt)); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting created_at: %s", err))
		}
	}
	if !core.IsNil(project.CumulativeNeedsAttentionView) {
		cumulativeNeedsAttentionView := []map[string]interface{}{}
		for _, cumulativeNeedsAttentionViewItem := range project.CumulativeNeedsAttentionView {
			cumulativeNeedsAttentionViewItemMap, err := resourceIbmProjectCumulativeNeedsAttentionToMap(&cumulativeNeedsAttentionViewItem)
			if err != nil {
				return diag.FromErr(err)
			}
			cumulativeNeedsAttentionView = append(cumulativeNeedsAttentionView, cumulativeNeedsAttentionViewItemMap)
		}
		if err = d.Set("cumulative_needs_attention_view", cumulativeNeedsAttentionView); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting cumulative_needs_attention_view: %s", err))
		}
	}
	if !core.IsNil(project.CumulativeNeedsAttentionViewError) {
		if err = d.Set("cumulative_needs_attention_view_error", project.CumulativeNeedsAttentionViewError); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting cumulative_needs_attention_view_error: %s", err))
		}
	}
	if !core.IsNil(project.State) {
		if err = d.Set("state", project.State); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting state: %s", err))
		}
	}
	if !core.IsNil(project.EventNotificationsCrn) {
		if err = d.Set("event_notifications_crn", project.EventNotificationsCrn); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting event_notifications_crn: %s", err))
		}
	}
	if !core.IsNil(project.Configs) {
		configs := []map[string]interface{}{}
		for _, configsItem := range project.Configs {
			configsItemMap, err := resourceIbmProjectProjectConfigCollectionMemberToMap(&configsItem)
			if err != nil {
				return diag.FromErr(err)
			}
			configs = append(configs, configsItemMap)
		}
		if err = d.Set("configs", configs); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting configs: %s", err))
		}
	}

	return nil
}

func resourceIbmProjectUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	projectClient, err := meta.(conns.ClientSession).ProjectV1()
	if err != nil {
		return diag.FromErr(err)
	}

	updateProjectOptions := &projectv1.UpdateProjectOptions{}

	updateProjectOptions.SetID(d.Id())

	hasChange := false

	if d.HasChange("definition") {
		definition, err := resourceIbmProjectMapToProjectPrototypePatchDefinitionBlock(d.Get("definition.0").(map[string]interface{}))
		if err != nil {
			return diag.FromErr(err)
		}
		updateProjectOptions.SetDefinition(definition)
		hasChange = true
	}

	if hasChange {
		_, response, err := projectClient.UpdateProjectWithContext(context, updateProjectOptions)
		if err != nil {
			log.Printf("[DEBUG] UpdateProjectWithContext failed %s\n%s", err, response)
			return diag.FromErr(fmt.Errorf("UpdateProjectWithContext failed %s\n%s", err, response))
		}
	}

	return resourceIbmProjectRead(context, d, meta)
}

func resourceIbmProjectDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	projectClient, err := meta.(conns.ClientSession).ProjectV1()
	if err != nil {
		return diag.FromErr(err)
	}

	deleteProjectOptions := &projectv1.DeleteProjectOptions{}

	deleteProjectOptions.SetID(d.Id())

	response, err := projectClient.DeleteProjectWithContext(context, deleteProjectOptions)
	if err != nil {
		log.Printf("[DEBUG] DeleteProjectWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("DeleteProjectWithContext failed %s\n%s", err, response))
	}

	d.SetId("")

	return nil
}

func resourceIbmProjectMapToProjectPrototypeDefinition(modelMap map[string]interface{}) (*projectv1.ProjectPrototypeDefinition, error) {
	model := &projectv1.ProjectPrototypeDefinition{}
	model.Name = core.StringPtr(modelMap["name"].(string))
	if modelMap["description"] != nil && modelMap["description"].(string) != "" {
		model.Description = core.StringPtr(modelMap["description"].(string))
	}
	if modelMap["destroy_on_delete"] != nil {
		model.DestroyOnDelete = core.BoolPtr(modelMap["destroy_on_delete"].(bool))
	}
	return model, nil
}

func resourceIbmProjectMapToProjectConfigPrototype(modelMap map[string]interface{}) (*projectv1.ProjectConfigPrototype, error) {
	model := &projectv1.ProjectConfigPrototype{}
	DefinitionModel, err := resourceIbmProjectMapToProjectConfigPrototypeDefinitionBlock(modelMap["definition"].([]interface{})[0].(map[string]interface{}))
	if err != nil {
		return model, err
	}
	model.Definition = DefinitionModel
	return model, nil
}

func resourceIbmProjectMapToProjectConfigPrototypeDefinitionBlock(modelMap map[string]interface{}) (*projectv1.ProjectConfigPrototypeDefinitionBlock, error) {
	model := &projectv1.ProjectConfigPrototypeDefinitionBlock{}
	model.Name = core.StringPtr(modelMap["name"].(string))
	if modelMap["description"] != nil && modelMap["description"].(string) != "" {
		model.Description = core.StringPtr(modelMap["description"].(string))
	}
	if modelMap["labels"] != nil {
		labels := []string{}
		for _, labelsItem := range modelMap["labels"].([]interface{}) {
			labels = append(labels, labelsItem.(string))
		}
		model.Labels = labels
	}
	if modelMap["authorizations"] != nil && len(modelMap["authorizations"].([]interface{})) > 0 {
		AuthorizationsModel, err := resourceIbmProjectMapToProjectConfigAuth(modelMap["authorizations"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.Authorizations = AuthorizationsModel
	}
	if modelMap["compliance_profile"] != nil && len(modelMap["compliance_profile"].([]interface{})) > 0 {
		ComplianceProfileModel, err := resourceIbmProjectMapToProjectComplianceProfile(modelMap["compliance_profile"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.ComplianceProfile = ComplianceProfileModel
	}
	model.LocatorID = core.StringPtr(modelMap["locator_id"].(string))
	if modelMap["input"] != nil && len(modelMap["input"].([]interface{})) > 0 {
		InputModel, err := resourceIbmProjectMapToInputVariable(modelMap["input"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.Input = InputModel
	}
	if modelMap["setting"] != nil && len(modelMap["setting"].([]interface{})) > 0 {
		SettingModel, err := resourceIbmProjectMapToProjectConfigSetting(modelMap["setting"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.Setting = SettingModel
	}
	return model, nil
}

func resourceIbmProjectMapToProjectConfigAuth(modelMap map[string]interface{}) (*projectv1.ProjectConfigAuth, error) {
	model := &projectv1.ProjectConfigAuth{}
	if modelMap["trusted_profile"] != nil && len(modelMap["trusted_profile"].([]interface{})) > 0 {
		TrustedProfileModel, err := resourceIbmProjectMapToProjectConfigAuthTrustedProfile(modelMap["trusted_profile"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.TrustedProfile = TrustedProfileModel
	}
	if modelMap["method"] != nil && modelMap["method"].(string) != "" {
		model.Method = core.StringPtr(modelMap["method"].(string))
	}
	if modelMap["api_key"] != nil && modelMap["api_key"].(string) != "" {
		model.ApiKey = core.StringPtr(modelMap["api_key"].(string))
	}
	return model, nil
}

func resourceIbmProjectMapToProjectConfigAuthTrustedProfile(modelMap map[string]interface{}) (*projectv1.ProjectConfigAuthTrustedProfile, error) {
	model := &projectv1.ProjectConfigAuthTrustedProfile{}
	if modelMap["id"] != nil && modelMap["id"].(string) != "" {
		model.ID = core.StringPtr(modelMap["id"].(string))
	}
	if modelMap["target_iam_id"] != nil && modelMap["target_iam_id"].(string) != "" {
		model.TargetIamID = core.StringPtr(modelMap["target_iam_id"].(string))
	}
	return model, nil
}

func resourceIbmProjectMapToProjectComplianceProfile(modelMap map[string]interface{}) (*projectv1.ProjectComplianceProfile, error) {
	model := &projectv1.ProjectComplianceProfile{}
	if modelMap["id"] != nil && modelMap["id"].(string) != "" {
		model.ID = core.StringPtr(modelMap["id"].(string))
	}
	if modelMap["instance_id"] != nil && modelMap["instance_id"].(string) != "" {
		model.InstanceID = core.StringPtr(modelMap["instance_id"].(string))
	}
	if modelMap["instance_location"] != nil && modelMap["instance_location"].(string) != "" {
		model.InstanceLocation = core.StringPtr(modelMap["instance_location"].(string))
	}
	if modelMap["attachment_id"] != nil && modelMap["attachment_id"].(string) != "" {
		model.AttachmentID = core.StringPtr(modelMap["attachment_id"].(string))
	}
	if modelMap["profile_name"] != nil && modelMap["profile_name"].(string) != "" {
		model.ProfileName = core.StringPtr(modelMap["profile_name"].(string))
	}
	return model, nil
}

func resourceIbmProjectMapToInputVariable(modelMap map[string]interface{}) (*projectv1.InputVariable, error) {
	model := &projectv1.InputVariable{}
	return model, nil
}

func resourceIbmProjectMapToProjectConfigSetting(modelMap map[string]interface{}) (*projectv1.ProjectConfigSetting, error) {
	model := &projectv1.ProjectConfigSetting{}
	return model, nil
}

func resourceIbmProjectMapToProjectPrototypePatchDefinitionBlock(modelMap map[string]interface{}) (*projectv1.ProjectPrototypePatchDefinitionBlock, error) {
	model := &projectv1.ProjectPrototypePatchDefinitionBlock{}
	if modelMap["name"] != nil && modelMap["name"].(string) != "" {
		model.Name = core.StringPtr(modelMap["name"].(string))
	}
	if modelMap["description"] != nil && modelMap["description"].(string) != "" {
		model.Description = core.StringPtr(modelMap["description"].(string))
	}
	if modelMap["destroy_on_delete"] != nil {
		model.DestroyOnDelete = core.BoolPtr(modelMap["destroy_on_delete"].(bool))
	}
	return model, nil
}

func resourceIbmProjectProjectDefinitionPropertiesToMap(model *projectv1.ProjectDefinitionProperties) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["name"] = model.Name
	if model.Description != nil {
		modelMap["description"] = model.Description
	}
	modelMap["destroy_on_delete"] = model.DestroyOnDelete
	return modelMap, nil
}

func resourceIbmProjectCumulativeNeedsAttentionToMap(model *projectv1.CumulativeNeedsAttention) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Event != nil {
		modelMap["event"] = model.Event
	}
	if model.EventID != nil {
		modelMap["event_id"] = model.EventID
	}
	if model.ConfigID != nil {
		modelMap["config_id"] = model.ConfigID
	}
	if model.ConfigVersion != nil {
		modelMap["config_version"] = flex.IntValue(model.ConfigVersion)
	}
	return modelMap, nil
}

func resourceIbmProjectProjectConfigCollectionMemberToMap(model *projectv1.ProjectConfigCollectionMember) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["id"] = model.ID
	modelMap["project_id"] = model.ProjectID
	modelMap["version"] = flex.IntValue(model.Version)
	modelMap["is_draft"] = model.IsDraft
	if model.NeedsAttentionState != nil {
		modelMap["needs_attention_state"] = model.NeedsAttentionState
	}
	modelMap["state"] = model.State
	modelMap["update_available"] = model.UpdateAvailable
	if model.CreatedAt != nil {
		modelMap["created_at"] = model.CreatedAt.String()
	}
	if model.UpdatedAt != nil {
		modelMap["updated_at"] = model.UpdatedAt.String()
	}
	if model.LastApproved != nil {
		lastApprovedMap, err := resourceIbmProjectProjectConfigMetadataLastApprovedToMap(model.LastApproved)
		if err != nil {
			return modelMap, err
		}
		modelMap["last_approved"] = []map[string]interface{}{lastApprovedMap}
	}
	if model.LastSave != nil {
		modelMap["last_save"] = model.LastSave.String()
	}
	if model.LastValidated != nil {
		lastValidatedMap, err := resourceIbmProjectLastValidatedActionWithSummaryToMap(model.LastValidated)
		if err != nil {
			return modelMap, err
		}
		modelMap["last_validated"] = []map[string]interface{}{lastValidatedMap}
	}
	if model.LastDeployed != nil {
		lastDeployedMap, err := resourceIbmProjectLastActionWithSummaryToMap(model.LastDeployed)
		if err != nil {
			return modelMap, err
		}
		modelMap["last_deployed"] = []map[string]interface{}{lastDeployedMap}
	}
	if model.LastUndeployed != nil {
		lastUndeployedMap, err := resourceIbmProjectLastActionWithSummaryToMap(model.LastUndeployed)
		if err != nil {
			return modelMap, err
		}
		modelMap["last_undeployed"] = []map[string]interface{}{lastUndeployedMap}
	}
	modelMap["href"] = model.Href
	definitionMap, err := resourceIbmProjectProjectConfigDefinitionNameDescriptionToMap(model.Definition)
	if err != nil {
		return modelMap, err
	}
	modelMap["definition"] = []map[string]interface{}{definitionMap}
	return modelMap, nil
}

func resourceIbmProjectProjectConfigMetadataLastApprovedToMap(model *projectv1.ProjectConfigMetadataLastApproved) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["is_forced"] = model.IsForced
	if model.Comment != nil {
		modelMap["comment"] = model.Comment
	}
	modelMap["timestamp"] = model.Timestamp.String()
	modelMap["user_id"] = model.UserID
	return modelMap, nil
}

func resourceIbmProjectLastValidatedActionWithSummaryToMap(model *projectv1.LastValidatedActionWithSummary) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Href != nil {
		modelMap["href"] = model.Href
	}
	if model.Result != nil {
		modelMap["result"] = model.Result
	}
	if model.Job != nil {
		jobMap, err := resourceIbmProjectActionJobWithIdAndSummaryToMap(model.Job)
		if err != nil {
			return modelMap, err
		}
		modelMap["job"] = []map[string]interface{}{jobMap}
	}
	if model.CostEstimate != nil {
		costEstimateMap, err := resourceIbmProjectProjectConfigMetadataCostEstimateToMap(model.CostEstimate)
		if err != nil {
			return modelMap, err
		}
		modelMap["cost_estimate"] = []map[string]interface{}{costEstimateMap}
	}
	if model.CraLogs != nil {
		craLogsMap, err := resourceIbmProjectProjectConfigMetadataCraLogsToMap(model.CraLogs)
		if err != nil {
			return modelMap, err
		}
		modelMap["cra_logs"] = []map[string]interface{}{craLogsMap}
	}
	return modelMap, nil
}

func resourceIbmProjectActionJobWithIdAndSummaryToMap(model *projectv1.ActionJobWithIdAndSummary) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ID != nil {
		modelMap["id"] = model.ID
	}
	if model.Summary != nil {
		summaryMap, err := resourceIbmProjectActionJobSummaryToMap(model.Summary)
		if err != nil {
			return modelMap, err
		}
		modelMap["summary"] = []map[string]interface{}{summaryMap}
	}
	return modelMap, nil
}

func resourceIbmProjectActionJobSummaryToMap(model *projectv1.ActionJobSummary) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.PlanSummary != nil {
		planSummary := make(map[string]interface{})
		for k, v := range model.PlanSummary {
			bytes, err := json.Marshal(v)
			if err != nil {
				return modelMap, err
			}
			planSummary[k] = string(bytes)
		}
		modelMap["plan_summary"] = planSummary
	}
	if model.ApplySummary != nil {
		applySummary := make(map[string]interface{})
		for k, v := range model.ApplySummary {
			bytes, err := json.Marshal(v)
			if err != nil {
				return modelMap, err
			}
			applySummary[k] = string(bytes)
		}
		modelMap["apply_summary"] = applySummary
	}
	if model.DestroySummary != nil {
		destroySummary := make(map[string]interface{})
		for k, v := range model.DestroySummary {
			bytes, err := json.Marshal(v)
			if err != nil {
				return modelMap, err
			}
			destroySummary[k] = string(bytes)
		}
		modelMap["destroy_summary"] = destroySummary
	}
	if model.MessageSummary != nil {
		messageSummary := make(map[string]interface{})
		for k, v := range model.MessageSummary {
			bytes, err := json.Marshal(v)
			if err != nil {
				return modelMap, err
			}
			messageSummary[k] = string(bytes)
		}
		modelMap["message_summary"] = messageSummary
	}
	if model.PlanMessages != nil {
		planMessages := make(map[string]interface{})
		for k, v := range model.PlanMessages {
			bytes, err := json.Marshal(v)
			if err != nil {
				return modelMap, err
			}
			planMessages[k] = string(bytes)
		}
		modelMap["plan_messages"] = planMessages
	}
	if model.ApplyMessages != nil {
		applyMessages := make(map[string]interface{})
		for k, v := range model.ApplyMessages {
			bytes, err := json.Marshal(v)
			if err != nil {
				return modelMap, err
			}
			applyMessages[k] = string(bytes)
		}
		modelMap["apply_messages"] = applyMessages
	}
	if model.DestroyMessages != nil {
		destroyMessages := make(map[string]interface{})
		for k, v := range model.DestroyMessages {
			bytes, err := json.Marshal(v)
			if err != nil {
				return modelMap, err
			}
			destroyMessages[k] = string(bytes)
		}
		modelMap["destroy_messages"] = destroyMessages
	}
	return modelMap, nil
}

func resourceIbmProjectProjectConfigMetadataCostEstimateToMap(model *projectv1.ProjectConfigMetadataCostEstimate) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Version != nil {
		modelMap["version"] = model.Version
	}
	if model.Currency != nil {
		modelMap["currency"] = model.Currency
	}
	if model.TotalHourlyCost != nil {
		modelMap["total_hourly_cost"] = model.TotalHourlyCost
	}
	if model.TotalMonthlyCost != nil {
		modelMap["total_monthly_cost"] = model.TotalMonthlyCost
	}
	if model.PastTotalHourlyCost != nil {
		modelMap["past_total_hourly_cost"] = model.PastTotalHourlyCost
	}
	if model.PastTotalMonthlyCost != nil {
		modelMap["past_total_monthly_cost"] = model.PastTotalMonthlyCost
	}
	if model.DiffTotalHourlyCost != nil {
		modelMap["diff_total_hourly_cost"] = model.DiffTotalHourlyCost
	}
	if model.DiffTotalMonthlyCost != nil {
		modelMap["diff_total_monthly_cost"] = model.DiffTotalMonthlyCost
	}
	if model.TimeGenerated != nil {
		modelMap["time_generated"] = model.TimeGenerated.String()
	}
	if model.UserID != nil {
		modelMap["user_id"] = model.UserID
	}
	return modelMap, nil
}

func resourceIbmProjectProjectConfigMetadataCraLogsToMap(model *projectv1.ProjectConfigMetadataCraLogs) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.CraVersion != nil {
		modelMap["cra_version"] = model.CraVersion
	}
	if model.SchemaVersion != nil {
		modelMap["schema_version"] = model.SchemaVersion
	}
	if model.Status != nil {
		modelMap["status"] = model.Status
	}
	if model.Summary != nil {
		summary := make(map[string]interface{})
		for k, v := range model.Summary {
			bytes, err := json.Marshal(v)
			if err != nil {
				return modelMap, err
			}
			summary[k] = string(bytes)
		}
		modelMap["summary"] = summary
	}
	if model.Timestamp != nil {
		modelMap["timestamp"] = model.Timestamp.String()
	}
	return modelMap, nil
}

func resourceIbmProjectLastActionWithSummaryToMap(model *projectv1.LastActionWithSummary) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Href != nil {
		modelMap["href"] = model.Href
	}
	if model.Result != nil {
		modelMap["result"] = model.Result
	}
	if model.Job != nil {
		jobMap, err := resourceIbmProjectActionJobWithIdAndSummaryToMap(model.Job)
		if err != nil {
			return modelMap, err
		}
		modelMap["job"] = []map[string]interface{}{jobMap}
	}
	return modelMap, nil
}

func resourceIbmProjectProjectConfigDefinitionNameDescriptionToMap(model *projectv1.ProjectConfigDefinitionNameDescription) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Name != nil {
		modelMap["name"] = model.Name
	}
	if model.Description != nil {
		modelMap["description"] = model.Description
	}
	return modelMap, nil
}
