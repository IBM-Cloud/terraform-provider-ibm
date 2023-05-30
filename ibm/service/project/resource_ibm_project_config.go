// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package project

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/validate"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/project-go-sdk/projectv1"
)

func ResourceIbmProjectConfig() *schema.Resource {
	return &schema.Resource{
		CreateContext: resourceIbmProjectConfigCreate,
		ReadContext:   resourceIbmProjectConfigRead,
		UpdateContext: resourceIbmProjectConfigUpdate,
		DeleteContext: resourceIbmProjectConfigDelete,
		Importer:      &schema.ResourceImporter{},
		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(1 * time.Minute),
			Delete: schema.DefaultTimeout(1 * time.Minute),
		},

		Schema: map[string]*schema.Schema{
			"project_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "The unique project ID.",
			},
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The configuration name.",
			},
			"locator_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "A dotted value of catalogID.versionID.",
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
			"input": &schema.Schema{
				Type:        schema.TypeList,
				Optional:    true,
				Description: "The inputs of a Schematics template property.",
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
			"type": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The type of a project configuration manual property.",
			},
			"output": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The outputs of a Schematics template property.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Required:    true,
							Description: "The variable name.",
						},
						"description": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "A short explanation of the output value.",
						},
						"value": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "Can be any value - a string, number, boolean, array, or object.",
						},
					},
				},
			},
			"metadata": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The project configuration draft.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"project_id": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The unique ID of a project.",
						},
						"version": &schema.Schema{
							Type:        schema.TypeInt,
							Optional:    true,
							Description: "The version number of the configuration.",
						},
						"state": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The state of the configuration draft.",
						},
						"pipeline_state": &schema.Schema{
							Type:        schema.TypeString,
							Optional:    true,
							Description: "The pipeline state of the configuration.",
						},
					},
				},
			},
			"project_config_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The ID of the configuration. If this parameter is empty, an ID is automatically created for the configuration.",
			},
		},
	}
}

func ResourceIbmProjectConfigValidator() *validate.ResourceValidator {
	validateSchema := make([]validate.ValidateSchema, 0)
	validateSchema = append(validateSchema,
		validate.ValidateSchema{
			Identifier:                 "project_id",
			ValidateFunctionIdentifier: validate.ValidateRegexp,
			Type:                       validate.TypeString,
			Required:                   true,
			Regexp:                     `^[\.\-0-9a-zA-Z]+$`,
			MaxValueLength:             128,
		},
		validate.ValidateSchema{
			Identifier:                 "name",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Required:                   true,
			Regexp:                     `^[a-zA-Z0-9][a-zA-Z0-9-_ ]*$`,
			MinValueLength:             1,
			MaxValueLength:             64,
		},
		validate.ValidateSchema{
			Identifier:                 "locator_id",
			ValidateFunctionIdentifier: validate.ValidateRegexpLen,
			Type:                       validate.TypeString,
			Required:                   true,
			Regexp:                     `^(?!\s)(?!.*\s$)[\.0-9a-z-A-Z_-]+$`,
			MinValueLength:             1,
			MaxValueLength:             512,
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

	resourceValidator := validate.ResourceValidator{ResourceName: "ibm_project_config", Schema: validateSchema}
	return &resourceValidator
}

func resourceIbmProjectConfigCreate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	projectClient, err := meta.(conns.ClientSession).ProjectV1()
	if err != nil {
		return diag.FromErr(err)
	}

	createConfigOptions := &projectv1.CreateConfigOptions{}

	createConfigOptions.SetProjectID(d.Get("project_id").(string))
	createConfigOptions.SetName(d.Get("name").(string))
	createConfigOptions.SetLocatorID(d.Get("locator_id").(string))
	if _, ok := d.GetOk("labels"); ok {
		var labels []string
		for _, v := range d.Get("labels").([]interface{}) {
			labelsItem := v.(string)
			labels = append(labels, labelsItem)
		}
		createConfigOptions.SetLabels(labels)
	}
	if _, ok := d.GetOk("description"); ok {
		createConfigOptions.SetDescription(d.Get("description").(string))
	}
	if _, ok := d.GetOk("authorizations"); ok {
		authorizationsModel, err := resourceIbmProjectConfigMapToProjectConfigAuth(d.Get("authorizations.0").(map[string]interface{}))
		if err != nil {
			return diag.FromErr(err)
		}
		createConfigOptions.SetAuthorizations(authorizationsModel)
	}
	if _, ok := d.GetOk("compliance_profile"); ok {
		complianceProfileModel, err := resourceIbmProjectConfigMapToProjectConfigComplianceProfile(d.Get("compliance_profile.0").(map[string]interface{}))
		if err != nil {
			return diag.FromErr(err)
		}
		createConfigOptions.SetComplianceProfile(complianceProfileModel)
	}
	if _, ok := d.GetOk("input"); ok {
		var input []projectv1.ProjectConfigInputVariable
		for _, v := range d.Get("input").([]interface{}) {
			value := v.(map[string]interface{})
			inputItem, err := resourceIbmProjectConfigMapToProjectConfigInputVariable(value)
			if err != nil {
				return diag.FromErr(err)
			}
			input = append(input, *inputItem)
		}
		createConfigOptions.SetInput(input)
	}
	if _, ok := d.GetOk("setting"); ok {
		var setting []projectv1.ProjectConfigSettingCollection
		for _, v := range d.Get("setting").([]interface{}) {
			value := v.(map[string]interface{})
			settingItem, err := resourceIbmProjectConfigMapToProjectConfigSettingCollection(value)
			if err != nil {
				return diag.FromErr(err)
			}
			setting = append(setting, *settingItem)
		}
		createConfigOptions.SetSetting(setting)
	}

	projectConfigGetResponse, response, err := projectClient.CreateConfigWithContext(context, createConfigOptions)
	if err != nil {
		log.Printf("[DEBUG] CreateConfigWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("CreateConfigWithContext failed %s\n%s", err, response))
	}

	d.SetId(fmt.Sprintf("%s/%s", *createConfigOptions.ProjectID, *projectConfigGetResponse.ID))

	_, err = waitForProjectConfigCreate(context, d, meta)
	if err != nil {
		return diag.FromErr(fmt.Errorf("[ERROR] Error waiting for create config instance (%s) to be succeeded: %s", d.Id(), err))
	}

	return resourceIbmProjectConfigRead(context, d, meta)
}

func waitForProjectConfigCreate(context context.Context, d *schema.ResourceData, meta interface{}) (interface{}, error) {
	projectClient, err := meta.(conns.ClientSession).ProjectV1()
	if err != nil {
		return false, err
	}

	parts, err := flex.SepIdParts(d.Id(), "/")
	getConfigOptions := &projectv1.GetConfigOptions{}
	getConfigOptions.SetProjectID(parts[0])
	getConfigOptions.SetID(parts[1])
	getConfigOptions.SetVersion("draft") // newly created config is in "draft" state, need to pass the version state

	stateConf := &resource.StateChangeConf{
		Pending: []string{"not_exists"},
		Target:  []string{"exists"},
		Refresh: func() (interface{}, string, error) {
			_, resp, err := projectClient.GetConfigWithContext(context, getConfigOptions)
			if err == nil {
				if resp != nil && resp.StatusCode == 200 {
					return resp, "exists", nil
				} else {
					return resp, "not_exists", nil
				}
			} else {
				return nil, "", fmt.Errorf("[ERROR] Get the config instance %s failed with resp code: %d, err: %v", d.Id(), resp.StatusCode, err)
			}
		},
		Timeout:    d.Timeout(schema.TimeoutCreate),
		Delay:      2 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func resourceIbmProjectConfigRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	projectClient, err := meta.(conns.ClientSession).ProjectV1()
	if err != nil {
		return diag.FromErr(err)
	}

	getConfigOptions := &projectv1.GetConfigOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return diag.FromErr(err)
	}

	getConfigOptions.SetProjectID(parts[0])
	getConfigOptions.SetID(parts[1])

	projectConfigGetResponse, response, err := projectClient.GetConfigWithContext(context, getConfigOptions)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		log.Printf("[DEBUG] GetConfigWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("GetConfigWithContext failed %s\n%s", err, response))
	}

	if err = d.Set("project_id", getConfigOptions.ProjectID); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting project_id: %s", err))
	}
	if err = d.Set("name", projectConfigGetResponse.Name); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting name: %s", err))
	}
	if err = d.Set("locator_id", projectConfigGetResponse.LocatorID); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting locator_id: %s", err))
	}
	if !core.IsNil(projectConfigGetResponse.Labels) {
		if err = d.Set("labels", projectConfigGetResponse.Labels); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting labels: %s", err))
		}
	}
	if !core.IsNil(projectConfigGetResponse.Description) {
		if err = d.Set("description", projectConfigGetResponse.Description); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting description: %s", err))
		}
	}
	if !core.IsNil(projectConfigGetResponse.Authorizations) {
		authorizationsMap, err := resourceIbmProjectConfigProjectConfigAuthToMap(projectConfigGetResponse.Authorizations)
		if err != nil {
			return diag.FromErr(err)
		}
		if err = d.Set("authorizations", []map[string]interface{}{authorizationsMap}); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting authorizations: %s", err))
		}
	}
	if !core.IsNil(projectConfigGetResponse.ComplianceProfile) {
		complianceProfileMap, err := resourceIbmProjectConfigProjectConfigComplianceProfileToMap(projectConfigGetResponse.ComplianceProfile)
		if err != nil {
			return diag.FromErr(err)
		}
		if err = d.Set("compliance_profile", []map[string]interface{}{complianceProfileMap}); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting compliance_profile: %s", err))
		}
	}

	if !core.IsNil(projectConfigGetResponse.Input) {
		input := []map[string]interface{}{}
		for _, inputItem := range projectConfigGetResponse.Input {
			inputItemMap, err := resourceIbmProjectConfigProjectConfigInputVariableToMap(&inputItem)
			if err != nil {
				return diag.FromErr(err)
			}
			input = append(input, inputItemMap)
		}
		if err = d.Set("input", input); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting input: %s", err))
		}
	}

	if !core.IsNil(projectConfigGetResponse.Setting) {
		setting := []map[string]interface{}{}
		for _, settingItem := range projectConfigGetResponse.Setting {
			settingItemMap, err := resourceIbmProjectConfigProjectConfigSettingCollectionToMap(&settingItem)
			if err != nil {
				return diag.FromErr(err)
			}
			setting = append(setting, settingItemMap)
		}
		if err = d.Set("setting", setting); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting setting: %s", err))
		}
	}
	if err = d.Set("type", projectConfigGetResponse.Type); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting type: %s", err))
	}
	if !core.IsNil(projectConfigGetResponse.Output) {
		output := []map[string]interface{}{}
		for _, outputItem := range projectConfigGetResponse.Output {
			outputItemMap, err := resourceIbmProjectConfigOutputValueToMap(&outputItem)
			if err != nil {
				return diag.FromErr(err)
			}
			output = append(output, outputItemMap)
		}
		if err = d.Set("output", output); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting output: %s", err))
		}
	}

	if !core.IsNil(projectConfigGetResponse.Metadata) {
		metadataMap, err := resourceIbmProjectConfigProjectConfigDraftMetadataToMap(projectConfigGetResponse.Metadata)
		if err != nil {
			return diag.FromErr(err)
		}
		if err = d.Set("metadata", []map[string]interface{}{metadataMap}); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting metadata: %s", err))
		}
	}

	if !core.IsNil(projectConfigGetResponse.ID) {
		if err = d.Set("project_config_id", projectConfigGetResponse.ID); err != nil {
			return diag.FromErr(fmt.Errorf("Error setting project_config_id: %s", err))
		}
	}

	return nil
}

func resourceIbmProjectConfigUpdate(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	projectClient, err := meta.(conns.ClientSession).ProjectV1()
	if err != nil {
		return diag.FromErr(err)
	}

	updateConfigOptions := &projectv1.UpdateConfigOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return diag.FromErr(err)
	}

	updateConfigOptions.SetProjectID(parts[0])
	updateConfigOptions.SetID(parts[1])

	hasChange := false

	if d.HasChange("project_id") {
		return diag.FromErr(fmt.Errorf("Cannot update resource property \"%s\" with the ForceNew annotation."+
			" The resource must be re-created to update this property.", "project_id"))
	}
	if d.HasChange("name") || d.HasChange("locator_id") {
		updateConfigOptions.SetName(d.Get("name").(string))
		updateConfigOptions.SetLocatorID(d.Get("locator_id").(string))
		hasChange = true
	}
	if d.HasChange("labels") {
		var labels []string
		for _, v := range d.Get("labels").([]interface{}) {
			labelsItem := v.(string)
			labels = append(labels, labelsItem)
		}
		updateConfigOptions.SetLabels(labels)
		hasChange = true
	}
	if d.HasChange("description") {
		updateConfigOptions.SetDescription(d.Get("description").(string))
		hasChange = true
	}
	if d.HasChange("authorizations") {
		authorizations, err := resourceIbmProjectConfigMapToProjectConfigAuth(d.Get("authorizations.0").(map[string]interface{}))
		if err != nil {
			return diag.FromErr(err)
		}
		updateConfigOptions.SetAuthorizations(authorizations)
		hasChange = true
	}
	if d.HasChange("compliance_profile") {
		complianceProfile, err := resourceIbmProjectConfigMapToProjectConfigComplianceProfile(d.Get("compliance_profile.0").(map[string]interface{}))
		if err != nil {
			return diag.FromErr(err)
		}
		updateConfigOptions.SetComplianceProfile(complianceProfile)
		hasChange = true
	}
	if d.HasChange("input") {
		var input []projectv1.ProjectConfigInputVariable
		for _, v := range d.Get("input").([]interface{}) {
			value := v.(map[string]interface{})
			inputItem, err := resourceIbmProjectConfigMapToProjectConfigInputVariable(value)
			if err != nil {
				return diag.FromErr(err)
			}
			input = append(input, *inputItem)
		}
		updateConfigOptions.SetInput(input)
		hasChange = true
	}
	if d.HasChange("setting") {
		var setting []projectv1.ProjectConfigSettingCollection
		for _, v := range d.Get("setting").([]interface{}) {
			value := v.(map[string]interface{})
			settingItem, err := resourceIbmProjectConfigMapToProjectConfigSettingCollection(value)
			if err != nil {
				return diag.FromErr(err)
			}
			setting = append(setting, *settingItem)
		}
		updateConfigOptions.SetSetting(setting)
		hasChange = true
	}

	if hasChange {
		_, response, err := projectClient.UpdateConfigWithContext(context, updateConfigOptions)
		if err != nil {
			log.Printf("[DEBUG] UpdateConfigWithContext failed %s\n%s", err, response)
			return diag.FromErr(fmt.Errorf("UpdateConfigWithContext failed %s\n%s", err, response))
		}
	}

	return resourceIbmProjectConfigRead(context, d, meta)
}

func resourceIbmProjectConfigDelete(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	projectClient, err := meta.(conns.ClientSession).ProjectV1()
	if err != nil {
		return diag.FromErr(err)
	}

	deleteConfigOptions := &projectv1.DeleteConfigOptions{}

	parts, err := flex.SepIdParts(d.Id(), "/")
	if err != nil {
		return diag.FromErr(err)
	}

	deleteConfigOptions.SetProjectID(parts[0])
	deleteConfigOptions.SetID(parts[1])

	_, response, err := projectClient.DeleteConfigWithContext(context, deleteConfigOptions)
	if err != nil {
		log.Printf("[DEBUG] DeleteConfigWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("DeleteConfigWithContext failed %s\n%s", err, response))
	}

	_, err = waitForProjectConfigDelete(context, d, meta)
	if err != nil {
		return diag.FromErr(fmt.Errorf("[ERROR] Error waiting for delete config instance (%s) to be succeeded: %s", d.Id(), err))
	}

	d.SetId("")

	return nil
}

func waitForProjectConfigDelete(context context.Context, d *schema.ResourceData, meta interface{}) (interface{}, error) {
	projectClient, err := meta.(conns.ClientSession).ProjectV1()
	if err != nil {
		return false, err
	}

	parts, err := flex.SepIdParts(d.Id(), "/")
	getConfigOptions := &projectv1.GetConfigOptions{}
	getConfigOptions.SetProjectID(parts[0])
	getConfigOptions.SetID(parts[1])

	stateConf := &resource.StateChangeConf{
		Pending: []string{"exists"},
		Target:  []string{"not_exists"},
		Refresh: func() (interface{}, string, error) {
			_, resp, err := projectClient.GetConfigWithContext(context, getConfigOptions)
			if err != nil {
				if resp != nil && resp.StatusCode == 404 {
					return resp, "not_exists", nil
				} else {
					return resp, "exists", nil
				}
			} else {
				return resp, "exists", nil
			}
		},
		Timeout:    d.Timeout(schema.TimeoutDelete),
		Delay:      2 * time.Second,
		MinTimeout: 10 * time.Second,
	}

	return stateConf.WaitForState()
}

func resourceIbmProjectConfigMapToProjectConfigAuth(modelMap map[string]interface{}) (*projectv1.ProjectConfigAuth, error) {
	model := &projectv1.ProjectConfigAuth{}
	if modelMap["trusted_profile"] != nil && len(modelMap["trusted_profile"].([]interface{})) > 0 {
		TrustedProfileModel, err := resourceIbmProjectConfigMapToProjectConfigAuthTrustedProfile(modelMap["trusted_profile"].([]interface{})[0].(map[string]interface{}))
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

func resourceIbmProjectConfigMapToProjectConfigAuthTrustedProfile(modelMap map[string]interface{}) (*projectv1.ProjectConfigAuthTrustedProfile, error) {
	model := &projectv1.ProjectConfigAuthTrustedProfile{}
	if modelMap["id"] != nil && modelMap["id"].(string) != "" {
		model.ID = core.StringPtr(modelMap["id"].(string))
	}
	if modelMap["target_iam_id"] != nil && modelMap["target_iam_id"].(string) != "" {
		model.TargetIamID = core.StringPtr(modelMap["target_iam_id"].(string))
	}
	return model, nil
}

func resourceIbmProjectConfigMapToProjectConfigComplianceProfile(modelMap map[string]interface{}) (*projectv1.ProjectConfigComplianceProfile, error) {
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

func resourceIbmProjectConfigMapToProjectConfigInputVariable(modelMap map[string]interface{}) (*projectv1.ProjectConfigInputVariable, error) {
	model := &projectv1.ProjectConfigInputVariable{}
	model.Name = core.StringPtr(modelMap["name"].(string))
	if modelMap["value"] != nil {
		model.Value = modelMap["value"].(string)
	}
	return model, nil
}

func resourceIbmProjectConfigMapToProjectConfigSettingCollection(modelMap map[string]interface{}) (*projectv1.ProjectConfigSettingCollection, error) {
	model := &projectv1.ProjectConfigSettingCollection{}
	model.Name = core.StringPtr(modelMap["name"].(string))
	model.Value = core.StringPtr(modelMap["value"].(string))
	return model, nil
}

func resourceIbmProjectConfigProjectConfigAuthToMap(model *projectv1.ProjectConfigAuth) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.TrustedProfile != nil {
		trustedProfileMap, err := resourceIbmProjectConfigProjectConfigAuthTrustedProfileToMap(model.TrustedProfile)
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

func resourceIbmProjectConfigProjectConfigAuthTrustedProfileToMap(model *projectv1.ProjectConfigAuthTrustedProfile) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ID != nil {
		modelMap["id"] = model.ID
	}
	if model.TargetIamID != nil {
		modelMap["target_iam_id"] = model.TargetIamID
	}
	return modelMap, nil
}

func resourceIbmProjectConfigProjectConfigComplianceProfileToMap(model *projectv1.ProjectConfigComplianceProfile) (map[string]interface{}, error) {
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

func resourceIbmProjectConfigProjectConfigInputVariableToMap(model *projectv1.ProjectConfigInputVariable) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["name"] = model.Name
	if model.Value != nil {
		modelMap["value"] = model.Value
	}
	return modelMap, nil
}

func resourceIbmProjectConfigProjectConfigSettingCollectionToMap(model *projectv1.ProjectConfigSettingCollection) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["name"] = model.Name
	modelMap["value"] = model.Value
	return modelMap, nil
}

func resourceIbmProjectConfigOutputValueToMap(model *projectv1.OutputValue) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["name"] = model.Name
	if model.Description != nil {
		modelMap["description"] = model.Description
	}
	if model.Value != nil {
		modelMap["value"] = model.Value
	}
	return modelMap, nil
}

func resourceIbmProjectConfigProjectConfigDraftMetadataToMap(model *projectv1.ProjectConfigDraftMetadata) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ProjectID != nil {
		modelMap["project_id"] = model.ProjectID
	}
	if model.Version != nil {
		modelMap["version"] = flex.IntValue(model.Version)
	}
	if model.State != nil {
		modelMap["state"] = model.State
	}
	if model.PipelineState != nil {
		modelMap["pipeline_state"] = model.PipelineState
	}
	return modelMap, nil
}
