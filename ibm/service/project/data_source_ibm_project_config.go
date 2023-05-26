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
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/project-go-sdk/projectv1"
)

func DataSourceIbmProjectConfig() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmProjectConfigRead,

		Schema: map[string]*schema.Schema{
			"project_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The unique project ID.",
			},
			"id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The unique config ID.",
			},
			"version": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "active",
				Description: "The version of the configuration to return.",
			},
			"project_config_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The ID of the configuration. If this parameter is empty, an ID is automatically created for the configuration.",
			},
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The configuration name.",
			},
			"labels": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "A collection of configuration labels.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"description": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The project configuration description.",
			},
			"authorizations": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The authorization for a configuration. You can authorize by using a trusted profile or an API key in Secrets Manager.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"trusted_profile": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The trusted profile for authorizations.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The unique ID of a project.",
									},
									"target_iam_id": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The unique ID of a project.",
									},
								},
							},
						},
						"method": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The authorization for a configuration. You can authorize by using a trusted profile or an API key in Secrets Manager.",
						},
						"api_key": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The IBM Cloud API Key.",
						},
					},
				},
			},
			"compliance_profile": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The profile required for compliance.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique ID of a project.",
						},
						"instance_id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique ID of a project.",
						},
						"instance_location": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The location of the compliance instance.",
						},
						"attachment_id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique ID of a project.",
						},
						"profile_name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the compliance profile.",
						},
					},
				},
			},
			"locator_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "A dotted value of catalogID.versionID.",
			},
			"type": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The type of a project configuration manual property.",
			},
			"input": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The outputs of a Schematics template property.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The variable name.",
						},
						"type": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The variable type.",
						},
						"value": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Can be any value - a string, number, boolean, array, or object.",
						},
						"required": &schema.Schema{
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "Whether the variable is required or not.",
						},
					},
				},
			},
			"output": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The outputs of a Schematics template property.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The variable name.",
						},
						"description": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "A short explanation of the output value.",
						},
						"value": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Can be any value - a string, number, boolean, array, or object.",
						},
					},
				},
			},
			"setting": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Schematics environment variables to use to deploy the configuration.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the configuration setting.",
						},
						"value": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The value of the configuration setting.",
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
							Computed:    true,
							Description: "The unique ID of a project.",
						},
						"version": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "The version number of the configuration.",
						},
						"state": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The state of the configuration draft.",
						},
						"pipeline_state": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The pipeline state of the configuration.",
						},
					},
				},
			},
		},
	}
}

func dataSourceIbmProjectConfigRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	projectClient, err := meta.(conns.ClientSession).ProjectV1()
	if err != nil {
		return diag.FromErr(err)
	}

	getConfigOptions := &projectv1.GetConfigOptions{}

	getConfigOptions.SetProjectID(d.Get("project_id").(string))
	getConfigOptions.SetID(d.Get("id").(string))
	if _, ok := d.GetOk("version"); ok {
		getConfigOptions.SetVersion(d.Get("version").(string))
	}

	projectConfig, response, err := projectClient.GetConfigWithContext(context, getConfigOptions)
	if err != nil {
		log.Printf("[DEBUG] GetConfigWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("GetConfigWithContext failed %s\n%s", err, response))
	}

	d.SetId(fmt.Sprintf("%s/%s", *getConfigOptions.ProjectID, *getConfigOptions.ID))

	if err = d.Set("project_config_id", projectConfig.ID); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting project_config_id: %s", err))
	}

	if err = d.Set("name", projectConfig.Name); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting name: %s", err))
	}

	if err = d.Set("description", projectConfig.Description); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting description: %s", err))
	}

	authorizations := []map[string]interface{}{}
	if projectConfig.Authorizations != nil {
		modelMap, err := dataSourceIbmProjectConfigProjectConfigAuthToMap(projectConfig.Authorizations)
		if err != nil {
			return diag.FromErr(err)
		}
		authorizations = append(authorizations, modelMap)
	}
	if err = d.Set("authorizations", authorizations); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting authorizations %s", err))
	}

	complianceProfile := []map[string]interface{}{}
	if projectConfig.ComplianceProfile != nil {
		modelMap, err := dataSourceIbmProjectConfigProjectConfigComplianceProfileToMap(projectConfig.ComplianceProfile)
		if err != nil {
			return diag.FromErr(err)
		}
		complianceProfile = append(complianceProfile, modelMap)
	}
	if err = d.Set("compliance_profile", complianceProfile); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting compliance_profile %s", err))
	}

	if err = d.Set("locator_id", projectConfig.LocatorID); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting locator_id: %s", err))
	}

	if err = d.Set("type", projectConfig.Type); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting type: %s", err))
	}

	input := []map[string]interface{}{}
	if projectConfig.Input != nil {
		for _, modelItem := range projectConfig.Input {
			modelMap, err := dataSourceIbmProjectConfigInputVariableToMap(&modelItem)
			if err != nil {
				return diag.FromErr(err)
			}
			input = append(input, modelMap)
		}
	}
	if err = d.Set("input", input); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting input %s", err))
	}

	output := []map[string]interface{}{}
	if projectConfig.Output != nil {
		for _, modelItem := range projectConfig.Output {
			modelMap, err := dataSourceIbmProjectConfigOutputValueToMap(&modelItem)
			if err != nil {
				return diag.FromErr(err)
			}
			output = append(output, modelMap)
		}
	}
	if err = d.Set("output", output); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting output %s", err))
	}

	setting := []map[string]interface{}{}
	if projectConfig.Setting != nil {
		for _, modelItem := range projectConfig.Setting {
			modelMap, err := dataSourceIbmProjectConfigProjectConfigSettingCollectionToMap(&modelItem)
			if err != nil {
				return diag.FromErr(err)
			}
			setting = append(setting, modelMap)
		}
	}
	if err = d.Set("setting", setting); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting setting %s", err))
	}

	metadata := []map[string]interface{}{}
	if projectConfig.Metadata != nil {
		modelMap, err := dataSourceIbmProjectConfigProjectConfigDraftMetadataToMap(projectConfig.Metadata)
		if err != nil {
			return diag.FromErr(err)
		}
		metadata = append(metadata, modelMap)
	}
	if err = d.Set("metadata", metadata); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting metadata %s", err))
	}

	return nil
}

func dataSourceIbmProjectConfigProjectConfigAuthToMap(model *projectv1.ProjectConfigAuth) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.TrustedProfile != nil {
		trustedProfileMap, err := dataSourceIbmProjectConfigProjectConfigAuthTrustedProfileToMap(model.TrustedProfile)
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

func dataSourceIbmProjectConfigProjectConfigAuthTrustedProfileToMap(model *projectv1.ProjectConfigAuthTrustedProfile) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ID != nil {
		modelMap["id"] = model.ID
	}
	if model.TargetIamID != nil {
		modelMap["target_iam_id"] = model.TargetIamID
	}
	return modelMap, nil
}

func dataSourceIbmProjectConfigProjectConfigComplianceProfileToMap(model *projectv1.ProjectConfigComplianceProfile) (map[string]interface{}, error) {
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

func dataSourceIbmProjectConfigInputVariableToMap(model *projectv1.InputVariable) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["name"] = model.Name
	modelMap["type"] = model.Type
	if model.Value != nil {
		modelMap["value"] = model.Value
	}
	if model.Required != nil {
		modelMap["required"] = model.Required
	}
	return modelMap, nil
}

func dataSourceIbmProjectConfigOutputValueToMap(model *projectv1.OutputValue) (map[string]interface{}, error) {
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

func dataSourceIbmProjectConfigProjectConfigSettingCollectionToMap(model *projectv1.ProjectConfigSettingCollection) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["name"] = model.Name
	modelMap["value"] = model.Value
	return modelMap, nil
}

func dataSourceIbmProjectConfigProjectConfigDraftMetadataToMap(model *projectv1.ProjectConfigDraftMetadata) (map[string]interface{}, error) {
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
