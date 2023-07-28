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
			"project_config_canonical_id": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The ID of the configuration. If this parameter is empty, an ID is automatically created for the configuration.",
			},
			"version": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The version of the configuration.",
			},
			"is_draft": &schema.Schema{
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "The flag that indicates whether the version of the configuration is draft, or active.",
			},
			"needs_attention_state": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The needs attention state of a configuration.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"state": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The state of the configuration.",
			},
			"pipeline_state": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The pipeline state of the configuration. It only exists after the first configuration validation.",
			},
			"update_available": &schema.Schema{
				Type:        schema.TypeBool,
				Computed:    true,
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
			"last_approved": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The last approved metadata of the configuration.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"is_forced": &schema.Schema{
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "The flag that indicates whether the approval was forced approved.",
						},
						"comment": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The comment left by the user who approved the configuration.",
						},
						"timestamp": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "A date and time value in the format YYYY-MM-DDTHH:mm:ssZ or YYYY-MM-DDTHH:mm:ss.sssZ, matching the date and time format as specified by RFC 3339.",
						},
						"user_id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique ID of a project.",
						},
					},
				},
			},
			"last_save": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "A date and time value in the format YYYY-MM-DDTHH:mm:ssZ or YYYY-MM-DDTHH:mm:ss.sssZ, matching the date and time format as specified by RFC 3339.",
			},
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name of the configuration.",
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
				Description: "The description of the project configuration.",
			},
			"authorizations": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The authorization for a configuration.You can authorize by using a trusted profile or an API key in Secrets Manager.",
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
			"input": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The inputs of a Schematics template property.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The variable name.",
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
				Description: "Schematics environment variables to use to deploy the configuration. Settings are only available if they were specified when the configuration was initially created.",
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
			"active_draft": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The project configuration version.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
							Description: "The pipeline state of the configuration. It only exists after the first configuration validation.",
						},
						"href": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "A relative URL.",
						},
					},
				},
			},
			"definition": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The project configuration definition.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the configuration.",
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
							Description: "The description of the project configuration.",
						},
						"authorizations": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The authorization for a configuration.You can authorize by using a trusted profile or an API key in Secrets Manager.",
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
						"input": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The inputs of a Schematics template property.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The variable name.",
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
							Description: "Schematics environment variables to use to deploy the configuration. Settings are only available if they were specified when the configuration was initially created.",
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
					},
				},
			},
			"href": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "A relative URL.",
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

	projectConfigCanonical, response, err := projectClient.GetConfigWithContext(context, getConfigOptions)
	if err != nil {
		log.Printf("[DEBUG] GetConfigWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("GetConfigWithContext failed %s\n%s", err, response))
	}

	d.SetId(fmt.Sprintf("%s/%s", *getConfigOptions.ProjectID, *getConfigOptions.ID))

	if err = d.Set("project_config_canonical_id", projectConfigCanonical.ID); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting project_config_canonical_id: %s", err))
	}

	if err = d.Set("version", flex.IntValue(projectConfigCanonical.Version)); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting version: %s", err))
	}

	if err = d.Set("is_draft", projectConfigCanonical.IsDraft); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting is_draft: %s", err))
	}

	if err = d.Set("needs_attention_state", projectConfigCanonical.NeedsAttentionState); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting needs_attention_state: %s", err))
	}

	if err = d.Set("state", projectConfigCanonical.State); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting state: %s", err))
	}

	if err = d.Set("pipeline_state", projectConfigCanonical.PipelineState); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting pipeline_state: %s", err))
	}

	if err = d.Set("update_available", projectConfigCanonical.UpdateAvailable); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting update_available: %s", err))
	}

	if err = d.Set("created_at", flex.DateTimeToString(projectConfigCanonical.CreatedAt)); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting created_at: %s", err))
	}

	if err = d.Set("updated_at", flex.DateTimeToString(projectConfigCanonical.UpdatedAt)); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting updated_at: %s", err))
	}

	lastApproved := []map[string]interface{}{}
	if projectConfigCanonical.LastApproved != nil {
		modelMap, err := dataSourceIbmProjectConfigProjectConfigMetadataLastApprovedToMap(projectConfigCanonical.LastApproved)
		if err != nil {
			return diag.FromErr(err)
		}
		lastApproved = append(lastApproved, modelMap)
	}
	if err = d.Set("last_approved", lastApproved); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting last_approved %s", err))
	}

	if err = d.Set("last_save", flex.DateTimeToString(projectConfigCanonical.LastSave)); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting last_save: %s", err))
	}

	if err = d.Set("name", projectConfigCanonical.Name); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting name: %s", err))
	}

	if err = d.Set("description", projectConfigCanonical.Description); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting description: %s", err))
	}

	authorizations := []map[string]interface{}{}
	if projectConfigCanonical.Authorizations != nil {
		modelMap, err := dataSourceIbmProjectConfigProjectConfigAuthToMap(projectConfigCanonical.Authorizations)
		if err != nil {
			return diag.FromErr(err)
		}
		authorizations = append(authorizations, modelMap)
	}
	if err = d.Set("authorizations", authorizations); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting authorizations %s", err))
	}

	complianceProfile := []map[string]interface{}{}
	if projectConfigCanonical.ComplianceProfile != nil {
		modelMap, err := dataSourceIbmProjectConfigProjectConfigComplianceProfileToMap(projectConfigCanonical.ComplianceProfile)
		if err != nil {
			return diag.FromErr(err)
		}
		complianceProfile = append(complianceProfile, modelMap)
	}
	if err = d.Set("compliance_profile", complianceProfile); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting compliance_profile %s", err))
	}

	if err = d.Set("locator_id", projectConfigCanonical.LocatorID); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting locator_id: %s", err))
	}

	input := []map[string]interface{}{}
	if projectConfigCanonical.Input != nil {
		for _, modelItem := range projectConfigCanonical.Input {
			modelMap, err := dataSourceIbmProjectConfigProjectConfigInputVariableToMap(&modelItem)
			if err != nil {
				return diag.FromErr(err)
			}
			input = append(input, modelMap)
		}
	}
	if err = d.Set("input", input); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting input %s", err))
	}

	setting := []map[string]interface{}{}
	if projectConfigCanonical.Setting != nil {
		for _, modelItem := range projectConfigCanonical.Setting {
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

	if err = d.Set("type", projectConfigCanonical.Type); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting type: %s", err))
	}

	output := []map[string]interface{}{}
	if projectConfigCanonical.Output != nil {
		for _, modelItem := range projectConfigCanonical.Output {
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

	activeDraft := []map[string]interface{}{}
	if projectConfigCanonical.ActiveDraft != nil {
		modelMap, err := dataSourceIbmProjectConfigProjectConfigVersionSummaryToMap(projectConfigCanonical.ActiveDraft)
		if err != nil {
			return diag.FromErr(err)
		}
		activeDraft = append(activeDraft, modelMap)
	}
	if err = d.Set("active_draft", activeDraft); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting active_draft %s", err))
	}

	definition := []map[string]interface{}{}
	if projectConfigCanonical.Definition != nil {
		modelMap, err := dataSourceIbmProjectConfigProjectConfigDefinitionToMap(projectConfigCanonical.Definition)
		if err != nil {
			return diag.FromErr(err)
		}
		definition = append(definition, modelMap)
	}
	if err = d.Set("definition", definition); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting definition %s", err))
	}

	if err = d.Set("href", projectConfigCanonical.Href); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting href: %s", err))
	}

	return nil
}

func dataSourceIbmProjectConfigProjectConfigMetadataLastApprovedToMap(model *projectv1.ProjectConfigMetadataLastApproved) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["is_forced"] = model.IsForced
	if model.Comment != nil {
		modelMap["comment"] = model.Comment
	}
	modelMap["timestamp"] = model.Timestamp.String()
	modelMap["user_id"] = model.UserID
	return modelMap, nil
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

func dataSourceIbmProjectConfigProjectConfigInputVariableToMap(model *projectv1.ProjectConfigInputVariable) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["name"] = model.Name
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

func dataSourceIbmProjectConfigProjectConfigVersionSummaryToMap(model *projectv1.ProjectConfigVersionSummary) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["version"] = flex.IntValue(model.Version)
	modelMap["state"] = model.State
	if model.PipelineState != nil {
		modelMap["pipeline_state"] = model.PipelineState
	}
	if model.Href != nil {
		modelMap["href"] = model.Href
	}
	return modelMap, nil
}

func dataSourceIbmProjectConfigProjectConfigDefinitionToMap(model *projectv1.ProjectConfigDefinition) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["name"] = model.Name
	if model.Labels != nil {
		modelMap["labels"] = model.Labels
	}
	if model.Description != nil {
		modelMap["description"] = model.Description
	}
	if model.Authorizations != nil {
		authorizationsMap, err := dataSourceIbmProjectConfigProjectConfigAuthToMap(model.Authorizations)
		if err != nil {
			return modelMap, err
		}
		modelMap["authorizations"] = []map[string]interface{}{authorizationsMap}
	}
	if model.ComplianceProfile != nil {
		complianceProfileMap, err := dataSourceIbmProjectConfigProjectConfigComplianceProfileToMap(model.ComplianceProfile)
		if err != nil {
			return modelMap, err
		}
		modelMap["compliance_profile"] = []map[string]interface{}{complianceProfileMap}
	}
	modelMap["locator_id"] = model.LocatorID
	if model.Input != nil {
		input := []map[string]interface{}{}
		for _, inputItem := range model.Input {
			inputItemMap, err := dataSourceIbmProjectConfigProjectConfigInputVariableToMap(&inputItem)
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
			settingItemMap, err := dataSourceIbmProjectConfigProjectConfigSettingCollectionToMap(&settingItem)
			if err != nil {
				return modelMap, err
			}
			setting = append(setting, settingItemMap)
		}
		modelMap["setting"] = setting
	}
	if model.Type != nil {
		modelMap["type"] = model.Type
	}
	if model.Output != nil {
		output := []map[string]interface{}{}
		for _, outputItem := range model.Output {
			outputItemMap, err := dataSourceIbmProjectConfigOutputValueToMap(&outputItem)
			if err != nil {
				return modelMap, err
			}
			output = append(output, outputItemMap)
		}
		modelMap["output"] = output
	}
	return modelMap, nil
}
