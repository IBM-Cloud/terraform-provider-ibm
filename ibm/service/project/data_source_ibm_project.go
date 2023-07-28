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

func DataSourceIbmProject() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmProjectRead,

		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The unique project ID.",
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
							Computed:    true,
							Description: "The event name.",
						},
						"event_id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "A unique ID for that individual event.",
						},
						"config_id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "A unique ID for the configuration.",
						},
						"config_version": &schema.Schema{
							Type:        schema.TypeInt,
							Computed:    true,
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
			"location": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The IBM Cloud location where a resource is deployed.",
			},
			"resource_group": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The resource group where the project's data and tools are created.",
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
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name of the project.",
			},
			"description": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "A brief explanation of the project's use in the configuration of a deployable architecture. It is possible to create a project without providing a description.",
			},
			"destroy_on_delete": &schema.Schema{
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "The policy that indicates whether the resources are destroyed or not when a project is deleted.",
			},
			"definition": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The definition of the project.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the project.",
						},
						"description": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "A brief explanation of the project's use in the configuration of a deployable architecture. It is possible to create a project without providing a description.",
						},
						"destroy_on_delete": &schema.Schema{
							Type:        schema.TypeBool,
							Computed:    true,
							Description: "The policy that indicates whether the resources are destroyed or not when a project is deleted.",
						},
					},
				},
			},
			"configs": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The project configurations. These configurations are only included in the response of creating a project if a configs array is specified in the request payload.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the configuration. If this parameter is empty, an ID is automatically created for the configuration.",
						},
						"project_id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique ID of a project.",
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
				},
			},
		},
	}
}

func dataSourceIbmProjectRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	projectClient, err := meta.(conns.ClientSession).ProjectV1()
	if err != nil {
		return diag.FromErr(err)
	}

	getProjectOptions := &projectv1.GetProjectOptions{}

	getProjectOptions.SetID(d.Get("id").(string))

	projectCanonical, response, err := projectClient.GetProjectWithContext(context, getProjectOptions)
	if err != nil {
		log.Printf("[DEBUG] GetProjectWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("GetProjectWithContext failed %s\n%s", err, response))
	}

	d.SetId(fmt.Sprintf("%s", *getProjectOptions.ID))

	if err = d.Set("crn", projectCanonical.Crn); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting crn: %s", err))
	}

	if err = d.Set("created_at", flex.DateTimeToString(projectCanonical.CreatedAt)); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting created_at: %s", err))
	}

	cumulativeNeedsAttentionView := []map[string]interface{}{}
	if projectCanonical.CumulativeNeedsAttentionView != nil {
		for _, modelItem := range projectCanonical.CumulativeNeedsAttentionView {
			modelMap, err := dataSourceIbmProjectCumulativeNeedsAttentionToMap(&modelItem)
			if err != nil {
				return diag.FromErr(err)
			}
			cumulativeNeedsAttentionView = append(cumulativeNeedsAttentionView, modelMap)
		}
	}
	if err = d.Set("cumulative_needs_attention_view", cumulativeNeedsAttentionView); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting cumulative_needs_attention_view %s", err))
	}

	if err = d.Set("cumulative_needs_attention_view_error", projectCanonical.CumulativeNeedsAttentionViewError); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting cumulative_needs_attention_view_error: %s", err))
	}

	if err = d.Set("location", projectCanonical.Location); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting location: %s", err))
	}

	if err = d.Set("resource_group", projectCanonical.ResourceGroup); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting resource_group: %s", err))
	}

	if err = d.Set("state", projectCanonical.State); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting state: %s", err))
	}

	if err = d.Set("event_notifications_crn", projectCanonical.EventNotificationsCrn); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting event_notifications_crn: %s", err))
	}

	if err = d.Set("name", projectCanonical.Name); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting name: %s", err))
	}

	if err = d.Set("description", projectCanonical.Description); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting description: %s", err))
	}

	if err = d.Set("destroy_on_delete", projectCanonical.DestroyOnDelete); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting destroy_on_delete: %s", err))
	}

	definition := []map[string]interface{}{}
	if projectCanonical.Definition != nil {
		modelMap, err := dataSourceIbmProjectProjectDefinitionResponseToMap(projectCanonical.Definition)
		if err != nil {
			return diag.FromErr(err)
		}
		definition = append(definition, modelMap)
	}
	if err = d.Set("definition", definition); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting definition %s", err))
	}

	configs := []map[string]interface{}{}
	if projectCanonical.Configs != nil {
		for _, modelItem := range projectCanonical.Configs {
			modelMap, err := dataSourceIbmProjectProjectConfigCanonicalToMap(&modelItem)
			if err != nil {
				return diag.FromErr(err)
			}
			configs = append(configs, modelMap)
		}
	}
	if err = d.Set("configs", configs); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting configs %s", err))
	}

	return nil
}

func dataSourceIbmProjectCumulativeNeedsAttentionToMap(model *projectv1.CumulativeNeedsAttention) (map[string]interface{}, error) {
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

func dataSourceIbmProjectProjectDefinitionResponseToMap(model *projectv1.ProjectDefinitionResponse) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["name"] = model.Name
	modelMap["description"] = model.Description
	modelMap["destroy_on_delete"] = model.DestroyOnDelete
	return modelMap, nil
}

func dataSourceIbmProjectProjectConfigCanonicalToMap(model *projectv1.ProjectConfigCanonical) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ID != nil {
		modelMap["id"] = model.ID
	}
	if model.ProjectID != nil {
		modelMap["project_id"] = model.ProjectID
	}
	if model.Version != nil {
		modelMap["version"] = flex.IntValue(model.Version)
	}
	if model.IsDraft != nil {
		modelMap["is_draft"] = model.IsDraft
	}
	if model.NeedsAttentionState != nil {
		modelMap["needs_attention_state"] = model.NeedsAttentionState
	}
	if model.State != nil {
		modelMap["state"] = model.State
	}
	if model.PipelineState != nil {
		modelMap["pipeline_state"] = model.PipelineState
	}
	if model.UpdateAvailable != nil {
		modelMap["update_available"] = model.UpdateAvailable
	}
	if model.CreatedAt != nil {
		modelMap["created_at"] = model.CreatedAt.String()
	}
	if model.UpdatedAt != nil {
		modelMap["updated_at"] = model.UpdatedAt.String()
	}
	if model.LastApproved != nil {
		lastApprovedMap, err := dataSourceIbmProjectProjectConfigMetadataLastApprovedToMap(model.LastApproved)
		if err != nil {
			return modelMap, err
		}
		modelMap["last_approved"] = []map[string]interface{}{lastApprovedMap}
	}
	if model.LastSave != nil {
		modelMap["last_save"] = model.LastSave.String()
	}
	modelMap["name"] = model.Name
	if model.Labels != nil {
		modelMap["labels"] = model.Labels
	}
	if model.Description != nil {
		modelMap["description"] = model.Description
	}
	if model.Authorizations != nil {
		authorizationsMap, err := dataSourceIbmProjectProjectConfigAuthToMap(model.Authorizations)
		if err != nil {
			return modelMap, err
		}
		modelMap["authorizations"] = []map[string]interface{}{authorizationsMap}
	}
	if model.ComplianceProfile != nil {
		complianceProfileMap, err := dataSourceIbmProjectProjectConfigComplianceProfileToMap(model.ComplianceProfile)
		if err != nil {
			return modelMap, err
		}
		modelMap["compliance_profile"] = []map[string]interface{}{complianceProfileMap}
	}
	modelMap["locator_id"] = model.LocatorID
	if model.Input != nil {
		input := []map[string]interface{}{}
		for _, inputItem := range model.Input {
			inputItemMap, err := dataSourceIbmProjectProjectConfigInputVariableToMap(&inputItem)
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
			settingItemMap, err := dataSourceIbmProjectProjectConfigSettingCollectionToMap(&settingItem)
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
			outputItemMap, err := dataSourceIbmProjectOutputValueToMap(&outputItem)
			if err != nil {
				return modelMap, err
			}
			output = append(output, outputItemMap)
		}
		modelMap["output"] = output
	}
	if model.ActiveDraft != nil {
		activeDraftMap, err := dataSourceIbmProjectProjectConfigVersionSummaryToMap(model.ActiveDraft)
		if err != nil {
			return modelMap, err
		}
		modelMap["active_draft"] = []map[string]interface{}{activeDraftMap}
	}
	if model.Definition != nil {
		definitionMap, err := dataSourceIbmProjectProjectConfigDefinitionToMap(model.Definition)
		if err != nil {
			return modelMap, err
		}
		modelMap["definition"] = []map[string]interface{}{definitionMap}
	}
	if model.Href != nil {
		modelMap["href"] = model.Href
	}
	return modelMap, nil
}

func dataSourceIbmProjectProjectConfigMetadataLastApprovedToMap(model *projectv1.ProjectConfigMetadataLastApproved) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["is_forced"] = model.IsForced
	if model.Comment != nil {
		modelMap["comment"] = model.Comment
	}
	modelMap["timestamp"] = model.Timestamp.String()
	modelMap["user_id"] = model.UserID
	return modelMap, nil
}

func dataSourceIbmProjectProjectConfigAuthToMap(model *projectv1.ProjectConfigAuth) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.TrustedProfile != nil {
		trustedProfileMap, err := dataSourceIbmProjectProjectConfigAuthTrustedProfileToMap(model.TrustedProfile)
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

func dataSourceIbmProjectProjectConfigAuthTrustedProfileToMap(model *projectv1.ProjectConfigAuthTrustedProfile) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ID != nil {
		modelMap["id"] = model.ID
	}
	if model.TargetIamID != nil {
		modelMap["target_iam_id"] = model.TargetIamID
	}
	return modelMap, nil
}

func dataSourceIbmProjectProjectConfigComplianceProfileToMap(model *projectv1.ProjectConfigComplianceProfile) (map[string]interface{}, error) {
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

func dataSourceIbmProjectProjectConfigInputVariableToMap(model *projectv1.ProjectConfigInputVariable) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["name"] = model.Name
	if model.Value != nil {
		modelMap["value"] = model.Value
	}
	return modelMap, nil
}

func dataSourceIbmProjectProjectConfigSettingCollectionToMap(model *projectv1.ProjectConfigSettingCollection) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["name"] = model.Name
	modelMap["value"] = model.Value
	return modelMap, nil
}

func dataSourceIbmProjectOutputValueToMap(model *projectv1.OutputValue) (map[string]interface{}, error) {
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

func dataSourceIbmProjectProjectConfigVersionSummaryToMap(model *projectv1.ProjectConfigVersionSummary) (map[string]interface{}, error) {
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

func dataSourceIbmProjectProjectConfigDefinitionToMap(model *projectv1.ProjectConfigDefinition) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["name"] = model.Name
	if model.Labels != nil {
		modelMap["labels"] = model.Labels
	}
	if model.Description != nil {
		modelMap["description"] = model.Description
	}
	if model.Authorizations != nil {
		authorizationsMap, err := dataSourceIbmProjectProjectConfigAuthToMap(model.Authorizations)
		if err != nil {
			return modelMap, err
		}
		modelMap["authorizations"] = []map[string]interface{}{authorizationsMap}
	}
	if model.ComplianceProfile != nil {
		complianceProfileMap, err := dataSourceIbmProjectProjectConfigComplianceProfileToMap(model.ComplianceProfile)
		if err != nil {
			return modelMap, err
		}
		modelMap["compliance_profile"] = []map[string]interface{}{complianceProfileMap}
	}
	modelMap["locator_id"] = model.LocatorID
	if model.Input != nil {
		input := []map[string]interface{}{}
		for _, inputItem := range model.Input {
			inputItemMap, err := dataSourceIbmProjectProjectConfigInputVariableToMap(&inputItem)
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
			settingItemMap, err := dataSourceIbmProjectProjectConfigSettingCollectionToMap(&settingItem)
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
			outputItemMap, err := dataSourceIbmProjectOutputValueToMap(&outputItem)
			if err != nil {
				return modelMap, err
			}
			output = append(output, outputItemMap)
		}
		modelMap["output"] = output
	}
	return modelMap, nil
}
