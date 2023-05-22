// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package project

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/project-go-sdk/projectv1"
)

func ResourceIbmProject() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIbmProjectCreate,
		ReadContext:   resourceIbmProjectRead,
		DeleteContext: resourceIbmProjectDelete,
		Importer:      &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"resource_group": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.InvokeValidator("ibm_project", "resource_group"),
				Description:  "The resource group where the project's data and tools are created.",
			},
			"location": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.InvokeValidator("ibm_project", "location"),
				Description:  "The location where the project's data and tools are created.",
			},
			"name": &schema.Schema{
				Type:         schema.TypeString,
				Required:     true,
				ForceNew:     true,
				ValidateFunc: validate.InvokeValidator("ibm_project", "name"),
				Description:  "The project name.",
			},
			"description": &schema.Schema{
				Type:         schema.TypeString,
				Optional:     true,
				ForceNew:     true,
				ValidateFunc: validate.InvokeValidator("ibm_project", "description"),
				Description:  "A project's descriptive text.",
			},
			"destroy_on_delete": &schema.Schema{
				Type:        schema.TypeBool,
				Optional:    true,
				Default:     true,
				ForceNew:    true,
				Description: "The policy that indicates whether the resources are destroyed or not when a project is deleted.",
			},
			"configs": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				ForceNew:    true,
				Description: "The project configurations.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The ID of the configuration. If this parameter is empty, an ID is automatically created for the configuration.",
						},
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "The configuration name.",
						},
						"labels": &schema.Schema{
							Type:        schema.TypeList,
							Optional:    true,
							Description: "A collection of configuration labels.",
							Elem:        &schema.Schema{Type: schema.TypeString},
						},
						"description": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The project configuration description.",
						},
						"authorizations": &schema.Schema{
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "The authorization for a configuration. You can authorize by using a trusted profile or an API key in Secrets Manager.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"trusted_profile": &schema.Schema{
										Type:        schema.TypeList,
										MaxItems:    1,
										Optional:    true,
										Description: "The trusted profile for authorizations.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"id": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "The unique ID of a project.",
												},
												"target_iam_id": &schema.Schema{
													Type:        schema.TypeString,
													Optional:    true,
													Description: "The unique ID of a project.",
												},
											},
										},
									},
									"method": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The authorization for a configuration. You can authorize by using a trusted profile or an API key in Secrets Manager.",
									},
									"api_key": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The IBM Cloud API Key.",
									},
								},
							},
						},
						"compliance_profile": &schema.Schema{
							Type:        schema.TypeList,
							MaxItems:    1,
							Optional:    true,
							Description: "The profile required for compliance.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The unique ID of a project.",
									},
									"instance_id": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The unique ID of a project.",
									},
									"instance_location": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The location of the compliance instance.",
									},
									"attachment_id": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The unique ID of a project.",
									},
									"profile_name": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "The name of the compliance profile.",
									},
								},
							},
						},
						"locator_id": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "A dotted value of catalogID.versionID.",
						},
						"input": &schema.Schema{
							Type:        schema.TypeList,
							Optional:    true,
							Description: "The input values to use to deploy the configuration.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:        schema.TypeString,
										Required:    true,
										Description: "The variable name.",
									},
									"value": &schema.Schema{
										Type:        schema.TypeString,
										Optional:    true,
										Description: "Can be any value - a string, number, boolean, array, or object.",
									},
								},
							},
						},
						"setting": &schema.Schema{
							Type:        schema.TypeList,
							Optional:    true,
							Description: "Schematics environment variables to use to deploy the configuration.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:        schema.TypeString,
										Required:    true,
										Description: "The name of the configuration setting.",
									},
									"value": &schema.Schema{
										Type:        schema.TypeString,
										Required:    true,
										Description: "The value of the configuration setting.",
									},
								},
							},
						},
					},
				},
			},
			"crn": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "An IBM Cloud resource name, which uniquely identifies a resource.",
			},
			"metadata": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The metadata of the project.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"crn": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "An IBM Cloud resource name, which uniquely identifies a resource.",
						},
						"created_at": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "A date and time value in the format YYYY-MM-DDTHH:mm:ssZ or YYYY-MM-DDTHH:mm:ss.sssZ, matching the date and time format as specified by RFC 3339.",
						},
						"cumulative_needs_attention_view": &schema.Schema{
							Type:        schema.TypeList,
							Optional:    true,
							Description: "The cumulative list of needs attention items for a project.",
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
						"cumulative_needs_attention_view_err": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "True indicates that the fetch of the needs attention items failed.",
						},
						"location": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The IBM Cloud location where a resource is deployed.",
						},
						"resource_group": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The resource group where the project's data and tools are created.",
						},
						"state": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The project status value.",
						},
						"event_notifications_crn": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The CRN of the event notifications instance if one is connected to this project.",
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
			Identifier:                 "resource_group",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Required:                   true,
			Regexp:                     `^$|^(?!\s)(?!.*\s$)[^'"<>{}\x00-\x1F]*$`,
			MinValueLength:             0,
			MaxValueLength:             40,
		},
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
			Identifier:                 "name",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Required:                   true,
			Regexp:                     `^(?!\s)(?!.*\s$)[^'"<>{}\x00-\x1F]+$`,
			MinValueLength:             1,
			MaxValueLength:             64,
		},
		validate.ValidateSchema{
			Identifier:                 "description",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Optional:                   true,
			Regexp:                     `^$|^(?!\s).*\S$`,
			MinValueLength:             0,
			MaxValueLength:             1024,
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

	createProjectOptions.SetResourceGroup(d.Get("resource_group").(string))
	createProjectOptions.SetLocation(d.Get("location").(string))
	createProjectOptions.SetName(d.Get("name").(string))
	if _, ok := d.GetOk("description"); ok {
		createProjectOptions.SetDescription(d.Get("description").(string))
	}
	if _, ok := d.GetOk("destroy_on_delete"); ok {
		createProjectOptions.SetDestroyOnDelete(d.Get("destroy_on_delete").(bool))
	}
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

	if err = d.Set("name", project.Name); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting name: %s", err))
	}
	if !core.IsNil(project.Description) {
		if err = d.Set("description", project.Description); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting description: %s", err))
		}
	}
	if !core.IsNil(project.DestroyOnDelete) {
		if err = d.Set("destroy_on_delete", project.DestroyOnDelete); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting destroy_on_delete: %s", err))
		}
	}
	if !core.IsNil(project.Configs) {
		configs := []map[string]interface{}{}
		for _, configsItem := range project.Configs {
			configsItemMap, err := resourceIbmProjectProjectConfigPrototypeToMap(&configsItem)
			if err != nil {
				return diag.FromErr(err)
			}
			configs = append(configs, configsItemMap)
		}
		if err = d.Set("configs", configs); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting configs: %s", err))
		}
	}
	if !core.IsNil(project.Crn) {
		if err = d.Set("crn", project.Crn); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting crn: %s", err))
		}
	}
	if !core.IsNil(project.Metadata) {
		metadataMap, err := resourceIbmProjectProjectMetadataToMap(project.Metadata)
		if err != nil {
			return diag.FromErr(err)
		}
		if err = d.Set("metadata", []map[string]interface{}{metadataMap}); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting metadata: %s", err))
		}
	}

	return nil
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

func resourceIbmProjectMapToProjectConfigPrototype(modelMap map[string]interface{}) (*projectv1.ProjectConfigPrototype, error) {
	model := &projectv1.ProjectConfigPrototype{}
	if modelMap["id"] != nil && modelMap["id"].(string) != "" {
		model.ID = core.StringPtr(modelMap["id"].(string))
	}
	model.Name = core.StringPtr(modelMap["name"].(string))
	if modelMap["labels"] != nil {
		labels := []string{}
		for _, labelsItem := range modelMap["labels"].([]interface{}) {
			labels = append(labels, labelsItem.(string))
		}
		model.Labels = labels
	}
	if modelMap["description"] != nil && modelMap["description"].(string) != "" {
		model.Description = core.StringPtr(modelMap["description"].(string))
	}
	if modelMap["authorizations"] != nil && len(modelMap["authorizations"].([]interface{})) > 0 {
		AuthorizationsModel, err := resourceIbmProjectMapToProjectConfigAuth(modelMap["authorizations"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.Authorizations = AuthorizationsModel
	}
	if modelMap["compliance_profile"] != nil && len(modelMap["compliance_profile"].([]interface{})) > 0 {
		ComplianceProfileModel, err := resourceIbmProjectMapToProjectConfigComplianceProfile(modelMap["compliance_profile"].([]interface{})[0].(map[string]interface{}))
		if err != nil {
			return model, err
		}
		model.ComplianceProfile = ComplianceProfileModel
	}
	model.LocatorID = core.StringPtr(modelMap["locator_id"].(string))
	if modelMap["input"] != nil {
		input := []projectv1.ProjectConfigInputVariable{}
		for _, inputItem := range modelMap["input"].([]interface{}) {
			inputItemModel, err := resourceIbmProjectMapToProjectConfigInputVariable(inputItem.(map[string]interface{}))
			if err != nil {
				return model, err
			}
			input = append(input, *inputItemModel)
		}
		model.Input = input
	}
	if modelMap["setting"] != nil {
		setting := []projectv1.ProjectConfigSettingCollection{}
		for _, settingItem := range modelMap["setting"].([]interface{}) {
			settingItemModel, err := resourceIbmProjectMapToProjectConfigSettingCollection(settingItem.(map[string]interface{}))
			if err != nil {
				return model, err
			}
			setting = append(setting, *settingItemModel)
		}
		model.Setting = setting
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

func resourceIbmProjectMapToProjectConfigComplianceProfile(modelMap map[string]interface{}) (*projectv1.ProjectConfigComplianceProfile, error) {
	model := &projectv1.ProjectConfigComplianceProfile{}
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

func resourceIbmProjectMapToProjectConfigInputVariable(modelMap map[string]interface{}) (*projectv1.ProjectConfigInputVariable, error) {
	model := &projectv1.ProjectConfigInputVariable{}
	model.Name = core.StringPtr(modelMap["name"].(string))
	if modelMap["value"] != nil {
		model.Value = modelMap["value"].(string)
	}
	return model, nil
}

func resourceIbmProjectMapToProjectConfigSettingCollection(modelMap map[string]interface{}) (*projectv1.ProjectConfigSettingCollection, error) {
	model := &projectv1.ProjectConfigSettingCollection{}
	model.Name = core.StringPtr(modelMap["name"].(string))
	model.Value = core.StringPtr(modelMap["value"].(string))
	return model, nil
}

func resourceIbmProjectProjectConfigPrototypeToMap(model *projectv1.ProjectConfigPrototype) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ID != nil {
		modelMap["id"] = model.ID
	}
	modelMap["name"] = model.Name
	if model.Labels != nil {
		modelMap["labels"] = model.Labels
	}
	if model.Description != nil {
		modelMap["description"] = model.Description
	}
	if model.Authorizations != nil {
		authorizationsMap, err := resourceIbmProjectProjectConfigAuthToMap(model.Authorizations)
		if err != nil {
			return modelMap, err
		}
		modelMap["authorizations"] = []map[string]interface{}{authorizationsMap}
	}
	if model.ComplianceProfile != nil {
		complianceProfileMap, err := resourceIbmProjectProjectConfigComplianceProfileToMap(model.ComplianceProfile)
		if err != nil {
			return modelMap, err
		}
		modelMap["compliance_profile"] = []map[string]interface{}{complianceProfileMap}
	}
	modelMap["locator_id"] = model.LocatorID
	if model.Input != nil {
		input := []map[string]interface{}{}
		for _, inputItem := range model.Input {
			inputItemMap, err := resourceIbmProjectProjectConfigInputVariableToMap(&inputItem)
			if err != nil {
				return modelMap, err
			}
			input = append(input, inputItemMap)
		}
		modelMap["input"] = input
	}
	if model.Setting != nil {
		setting := []map[string]interface{}{}
		for _, settingItem := range model.Setting {
			settingItemMap, err := resourceIbmProjectProjectConfigSettingCollectionToMap(&settingItem)
			if err != nil {
				return modelMap, err
			}
			setting = append(setting, settingItemMap)
		}
		modelMap["setting"] = setting
	}
	return modelMap, nil
}

func resourceIbmProjectProjectConfigAuthToMap(model *projectv1.ProjectConfigAuth) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.TrustedProfile != nil {
		trustedProfileMap, err := resourceIbmProjectProjectConfigAuthTrustedProfileToMap(model.TrustedProfile)
		if err != nil {
			return modelMap, err
		}
		modelMap["trusted_profile"] = []map[string]interface{}{trustedProfileMap}
	}
	if model.Method != nil {
		modelMap["method"] = model.Method
	}
	if model.ApiKey != nil {
		modelMap["api_key"] = model.ApiKey
	}
	return modelMap, nil
}

func resourceIbmProjectProjectConfigAuthTrustedProfileToMap(model *projectv1.ProjectConfigAuthTrustedProfile) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ID != nil {
		modelMap["id"] = model.ID
	}
	if model.TargetIamID != nil {
		modelMap["target_iam_id"] = model.TargetIamID
	}
	return modelMap, nil
}

func resourceIbmProjectProjectConfigComplianceProfileToMap(model *projectv1.ProjectConfigComplianceProfile) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ID != nil {
		modelMap["id"] = model.ID
	}
	if model.InstanceID != nil {
		modelMap["instance_id"] = model.InstanceID
	}
	if model.InstanceLocation != nil {
		modelMap["instance_location"] = model.InstanceLocation
	}
	if model.AttachmentID != nil {
		modelMap["attachment_id"] = model.AttachmentID
	}
	if model.ProfileName != nil {
		modelMap["profile_name"] = model.ProfileName
	}
	return modelMap, nil
}

func resourceIbmProjectProjectConfigInputVariableToMap(model *projectv1.ProjectConfigInputVariable) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["name"] = model.Name
	if model.Value != nil {
		modelMap["value"] = model.Value
	}
	return modelMap, nil
}

func resourceIbmProjectProjectConfigSettingCollectionToMap(model *projectv1.ProjectConfigSettingCollection) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["name"] = model.Name
	modelMap["value"] = model.Value
	return modelMap, nil
}

func resourceIbmProjectProjectMetadataToMap(model *projectv1.ProjectMetadata) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Crn != nil {
		modelMap["crn"] = model.Crn
	}
	if model.CreatedAt != nil {
		modelMap["created_at"] = model.CreatedAt.String()
	}
	if model.CumulativeNeedsAttentionView != nil {
		cumulativeNeedsAttentionView := []map[string]interface{}{}
		for _, cumulativeNeedsAttentionViewItem := range model.CumulativeNeedsAttentionView {
			cumulativeNeedsAttentionViewItemMap, err := resourceIbmProjectCumulativeNeedsAttentionToMap(&cumulativeNeedsAttentionViewItem)
			if err != nil {
				return modelMap, err
			}
			cumulativeNeedsAttentionView = append(cumulativeNeedsAttentionView, cumulativeNeedsAttentionViewItemMap)
		}
		modelMap["cumulative_needs_attention_view"] = cumulativeNeedsAttentionView
	}
	if model.CumulativeNeedsAttentionViewErr != nil {
		modelMap["cumulative_needs_attention_view_err"] = model.CumulativeNeedsAttentionViewErr
	}
	if model.Location != nil {
		modelMap["location"] = model.Location
	}
	if model.ResourceGroup != nil {
		modelMap["resource_group"] = model.ResourceGroup
	}
	if model.State != nil {
		modelMap["state"] = model.State
	}
	if model.EventNotificationsCrn != nil {
		modelMap["event_notifications_crn"] = model.EventNotificationsCrn
	}
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
		modelMap["config_version"] = model.ConfigVersion
	}
	return modelMap, nil
}
