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
			"project_config_id": &schema.Schema{
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
							Description: "The unique ID.",
						},
					},
				},
			},
			"last_save": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "A date and time value in the format YYYY-MM-DDTHH:mm:ssZ or YYYY-MM-DDTHH:mm:ss.sssZ, matching the date and time format as specified by RFC 3339.",
			},
			"last_validated": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The action job performed on the project configuration.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"href": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "A relative URL.",
						},
						"result": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The result of the last action.",
						},
						"job": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "A brief summary of an action.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The unique ID.",
									},
									"summary": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The summaries of jobs that were performed on the configuration.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"plan_summary": &schema.Schema{
													Type:        schema.TypeMap,
													Computed:    true,
													Description: "The summary of the plan jobs on the configuration.",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"apply_summary": &schema.Schema{
													Type:        schema.TypeMap,
													Computed:    true,
													Description: "The summary of the apply jobs on the configuration.",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"destroy_summary": &schema.Schema{
													Type:        schema.TypeMap,
													Computed:    true,
													Description: "The summary of the destroy jobs on the configuration.",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"message_summary": &schema.Schema{
													Type:        schema.TypeMap,
													Computed:    true,
													Description: "The message summaries of jobs on the configuration.",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"plan_messages": &schema.Schema{
													Type:        schema.TypeMap,
													Computed:    true,
													Description: "The messages of plan jobs on the configuration.",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"apply_messages": &schema.Schema{
													Type:        schema.TypeMap,
													Computed:    true,
													Description: "The messages of apply jobs on the configuration.",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"destroy_messages": &schema.Schema{
													Type:        schema.TypeMap,
													Computed:    true,
													Description: "The messages of destroy jobs on the configuration.",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
									},
								},
							},
						},
						"cost_estimate": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The cost estimate of the configuration.It only exists after the first configuration validation.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"version": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The version of the cost estimate of the configuration.",
									},
									"currency": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The currency of the cost estimate of the configuration.",
									},
									"total_hourly_cost": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The total hourly cost estimate of the configuration.",
									},
									"total_monthly_cost": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The total monthly cost estimate of the configuration.",
									},
									"past_total_hourly_cost": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The past total hourly cost estimate of the configuration.",
									},
									"past_total_monthly_cost": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The past total monthly cost estimate of the configuration.",
									},
									"diff_total_hourly_cost": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The difference between current and past total hourly cost estimates of the configuration.",
									},
									"diff_total_monthly_cost": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The difference between current and past total monthly cost estimates of the configuration.",
									},
									"time_generated": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "A date and time value in the format YYYY-MM-DDTHH:mm:ssZ or YYYY-MM-DDTHH:mm:ss.sssZ, matching the date and time format as specified by RFC 3339.",
									},
									"user_id": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The unique ID.",
									},
								},
							},
						},
						"cra_logs": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The Code Risk Analyzer logs of the configuration.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"cra_version": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The version of the Code Risk Analyzer logs of the configuration.",
									},
									"schema_version": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The schema version of Code Risk Analyzer logs of the configuration.",
									},
									"status": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The status of the Code Risk Analyzer logs of the configuration.",
									},
									"summary": &schema.Schema{
										Type:        schema.TypeMap,
										Computed:    true,
										Description: "The summary of the Code Risk Analyzer logs of the configuration.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
									"timestamp": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "A date and time value in the format YYYY-MM-DDTHH:mm:ssZ or YYYY-MM-DDTHH:mm:ss.sssZ, matching the date and time format as specified by RFC 3339.",
									},
								},
							},
						},
					},
				},
			},
			"last_deployed": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The action job performed on the project configuration.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"href": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "A relative URL.",
						},
						"result": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The result of the last action.",
						},
						"job": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "A brief summary of an action.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The unique ID.",
									},
									"summary": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The summaries of jobs that were performed on the configuration.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"plan_summary": &schema.Schema{
													Type:        schema.TypeMap,
													Computed:    true,
													Description: "The summary of the plan jobs on the configuration.",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"apply_summary": &schema.Schema{
													Type:        schema.TypeMap,
													Computed:    true,
													Description: "The summary of the apply jobs on the configuration.",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"destroy_summary": &schema.Schema{
													Type:        schema.TypeMap,
													Computed:    true,
													Description: "The summary of the destroy jobs on the configuration.",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"message_summary": &schema.Schema{
													Type:        schema.TypeMap,
													Computed:    true,
													Description: "The message summaries of jobs on the configuration.",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"plan_messages": &schema.Schema{
													Type:        schema.TypeMap,
													Computed:    true,
													Description: "The messages of plan jobs on the configuration.",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"apply_messages": &schema.Schema{
													Type:        schema.TypeMap,
													Computed:    true,
													Description: "The messages of apply jobs on the configuration.",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"destroy_messages": &schema.Schema{
													Type:        schema.TypeMap,
													Computed:    true,
													Description: "The messages of destroy jobs on the configuration.",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			"last_undeployed": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The action job performed on the project configuration.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"href": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "A relative URL.",
						},
						"result": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The result of the last action.",
						},
						"job": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "A brief summary of an action.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The unique ID.",
									},
									"summary": &schema.Schema{
										Type:        schema.TypeList,
										Computed:    true,
										Description: "The summaries of jobs that were performed on the configuration.",
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"plan_summary": &schema.Schema{
													Type:        schema.TypeMap,
													Computed:    true,
													Description: "The summary of the plan jobs on the configuration.",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"apply_summary": &schema.Schema{
													Type:        schema.TypeMap,
													Computed:    true,
													Description: "The summary of the apply jobs on the configuration.",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"destroy_summary": &schema.Schema{
													Type:        schema.TypeMap,
													Computed:    true,
													Description: "The summary of the destroy jobs on the configuration.",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"message_summary": &schema.Schema{
													Type:        schema.TypeMap,
													Computed:    true,
													Description: "The message summaries of jobs on the configuration.",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"plan_messages": &schema.Schema{
													Type:        schema.TypeMap,
													Computed:    true,
													Description: "The messages of plan jobs on the configuration.",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"apply_messages": &schema.Schema{
													Type:        schema.TypeMap,
													Computed:    true,
													Description: "The messages of apply jobs on the configuration.",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
												"destroy_messages": &schema.Schema{
													Type:        schema.TypeMap,
													Computed:    true,
													Description: "The messages of destroy jobs on the configuration.",
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
			"definition": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The type and output of a project configuration.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The configuration name.",
						},
						"description": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "A project configuration description.",
						},
						"labels": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The configuration labels.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
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
													Description: "The unique ID.",
												},
												"target_iam_id": &schema.Schema{
													Type:        schema.TypeString,
													Computed:    true,
													Description: "The unique ID.",
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
										Description: "The unique ID.",
									},
									"instance_id": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The unique ID.",
									},
									"instance_location": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The location of the compliance instance.",
									},
									"attachment_id": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The unique ID.",
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
							Description: "The input variables for configuration definition and environment.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{},
							},
						},
						"setting": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Schematics environment variables to use to deploy the configuration.Settings are only available if they were specified when the configuration was initially created.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{},
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

	projectConfig, response, err := projectClient.GetConfigWithContext(context, getConfigOptions)
	if err != nil {
		log.Printf("[DEBUG] GetConfigWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("GetConfigWithContext failed %s\n%s", err, response))
	}

	d.SetId(fmt.Sprintf("%s/%s", *getConfigOptions.ProjectID, *getConfigOptions.ID))

	if err = d.Set("project_config_id", projectConfig.ID); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting project_config_id: %s", err))
	}

	if err = d.Set("version", flex.IntValue(projectConfig.Version)); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting version: %s", err))
	}

	if err = d.Set("is_draft", projectConfig.IsDraft); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting is_draft: %s", err))
	}

	if err = d.Set("needs_attention_state", projectConfig.NeedsAttentionState); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting needs_attention_state: %s", err))
	}

	if err = d.Set("state", projectConfig.State); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting state: %s", err))
	}

	if err = d.Set("update_available", projectConfig.UpdateAvailable); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting update_available: %s", err))
	}

	if err = d.Set("created_at", flex.DateTimeToString(projectConfig.CreatedAt)); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting created_at: %s", err))
	}

	if err = d.Set("updated_at", flex.DateTimeToString(projectConfig.UpdatedAt)); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting updated_at: %s", err))
	}

	lastApproved := []map[string]interface{}{}
	if projectConfig.LastApproved != nil {
		modelMap, err := dataSourceIbmProjectConfigProjectConfigMetadataLastApprovedToMap(projectConfig.LastApproved)
		if err != nil {
			return diag.FromErr(err)
		}
		lastApproved = append(lastApproved, modelMap)
	}
	if err = d.Set("last_approved", lastApproved); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting last_approved %s", err))
	}

	if err = d.Set("last_save", flex.DateTimeToString(projectConfig.LastSave)); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting last_save: %s", err))
	}

	lastValidated := []map[string]interface{}{}
	if projectConfig.LastValidated != nil {
		modelMap, err := dataSourceIbmProjectConfigLastValidatedActionWithSummaryToMap(projectConfig.LastValidated)
		if err != nil {
			return diag.FromErr(err)
		}
		lastValidated = append(lastValidated, modelMap)
	}
	if err = d.Set("last_validated", lastValidated); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting last_validated %s", err))
	}

	lastDeployed := []map[string]interface{}{}
	if projectConfig.LastDeployed != nil {
		modelMap, err := dataSourceIbmProjectConfigLastActionWithSummaryToMap(projectConfig.LastDeployed)
		if err != nil {
			return diag.FromErr(err)
		}
		lastDeployed = append(lastDeployed, modelMap)
	}
	if err = d.Set("last_deployed", lastDeployed); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting last_deployed %s", err))
	}

	lastUndeployed := []map[string]interface{}{}
	if projectConfig.LastUndeployed != nil {
		modelMap, err := dataSourceIbmProjectConfigLastActionWithSummaryToMap(projectConfig.LastUndeployed)
		if err != nil {
			return diag.FromErr(err)
		}
		lastUndeployed = append(lastUndeployed, modelMap)
	}
	if err = d.Set("last_undeployed", lastUndeployed); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting last_undeployed %s", err))
	}

	definition := []map[string]interface{}{}
	if projectConfig.Definition != nil {
		modelMap, err := dataSourceIbmProjectConfigProjectConfigResponseDefinitionToMap(projectConfig.Definition)
		if err != nil {
			return diag.FromErr(err)
		}
		definition = append(definition, modelMap)
	}
	if err = d.Set("definition", definition); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting definition %s", err))
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

func dataSourceIbmProjectConfigLastValidatedActionWithSummaryToMap(model *projectv1.LastValidatedActionWithSummary) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Href != nil {
		modelMap["href"] = model.Href
	}
	if model.Result != nil {
		modelMap["result"] = model.Result
	}
	if model.Job != nil {
		jobMap, err := dataSourceIbmProjectConfigActionJobWithIdAndSummaryToMap(model.Job)
		if err != nil {
			return modelMap, err
		}
		modelMap["job"] = []map[string]interface{}{jobMap}
	}
	if model.CostEstimate != nil {
		costEstimateMap, err := dataSourceIbmProjectConfigProjectConfigMetadataCostEstimateToMap(model.CostEstimate)
		if err != nil {
			return modelMap, err
		}
		modelMap["cost_estimate"] = []map[string]interface{}{costEstimateMap}
	}
	if model.CraLogs != nil {
		craLogsMap, err := dataSourceIbmProjectConfigProjectConfigMetadataCraLogsToMap(model.CraLogs)
		if err != nil {
			return modelMap, err
		}
		modelMap["cra_logs"] = []map[string]interface{}{craLogsMap}
	}
	return modelMap, nil
}

func dataSourceIbmProjectConfigActionJobWithIdAndSummaryToMap(model *projectv1.ActionJobWithIdAndSummary) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ID != nil {
		modelMap["id"] = model.ID
	}
	if model.Summary != nil {
		summaryMap, err := dataSourceIbmProjectConfigActionJobSummaryToMap(model.Summary)
		if err != nil {
			return modelMap, err
		}
		modelMap["summary"] = []map[string]interface{}{summaryMap}
	}
	return modelMap, nil
}

func dataSourceIbmProjectConfigActionJobSummaryToMap(model *projectv1.ActionJobSummary) (map[string]interface{}, error) {
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

func dataSourceIbmProjectConfigProjectConfigMetadataCostEstimateToMap(model *projectv1.ProjectConfigMetadataCostEstimate) (map[string]interface{}, error) {
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

func dataSourceIbmProjectConfigProjectConfigMetadataCraLogsToMap(model *projectv1.ProjectConfigMetadataCraLogs) (map[string]interface{}, error) {
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

func dataSourceIbmProjectConfigLastActionWithSummaryToMap(model *projectv1.LastActionWithSummary) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Href != nil {
		modelMap["href"] = model.Href
	}
	if model.Result != nil {
		modelMap["result"] = model.Result
	}
	if model.Job != nil {
		jobMap, err := dataSourceIbmProjectConfigActionJobWithIdAndSummaryToMap(model.Job)
		if err != nil {
			return modelMap, err
		}
		modelMap["job"] = []map[string]interface{}{jobMap}
	}
	return modelMap, nil
}

func dataSourceIbmProjectConfigProjectConfigResponseDefinitionToMap(model *projectv1.ProjectConfigResponseDefinition) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["name"] = model.Name
	if model.Description != nil {
		modelMap["description"] = model.Description
	}
	if model.Labels != nil {
		modelMap["labels"] = model.Labels
	}
	if model.Authorizations != nil {
		authorizationsMap, err := dataSourceIbmProjectConfigProjectConfigAuthToMap(model.Authorizations)
		if err != nil {
			return modelMap, err
		}
		modelMap["authorizations"] = []map[string]interface{}{authorizationsMap}
	}
	if model.ComplianceProfile != nil {
		complianceProfileMap, err := dataSourceIbmProjectConfigProjectComplianceProfileToMap(model.ComplianceProfile)
		if err != nil {
			return modelMap, err
		}
		modelMap["compliance_profile"] = []map[string]interface{}{complianceProfileMap}
	}
	modelMap["locator_id"] = model.LocatorID
	if model.Input != nil {
		inputMap, err := dataSourceIbmProjectConfigInputVariableToMap(model.Input)
		if err != nil {
			return modelMap, err
		}
		modelMap["input"] = []map[string]interface{}{inputMap}
	}
	if model.Setting != nil {
		settingMap, err := dataSourceIbmProjectConfigProjectConfigSettingToMap(model.Setting)
		if err != nil {
			return modelMap, err
		}
		modelMap["setting"] = []map[string]interface{}{settingMap}
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

func dataSourceIbmProjectConfigProjectComplianceProfileToMap(model *projectv1.ProjectComplianceProfile) (map[string]interface{}, error) {
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
	return modelMap, nil
}

func dataSourceIbmProjectConfigProjectConfigSettingToMap(model *projectv1.ProjectConfigSetting) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
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
