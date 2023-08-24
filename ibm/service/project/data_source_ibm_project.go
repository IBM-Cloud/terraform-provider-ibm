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
			"configs": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The project configurations. These configurations are only included in the response of creating a project if a configs array is specified in the request payload.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The name of the configuration.",
						},
						"description": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The description of the project configuration.",
						},
						"labels": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "A collection of configuration labels.",
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
							Description: "The input variables for the configuration definition.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
								},
							},
						},
						"setting": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Schematics environment variables to use to deploy the configuration.Settings are only available if they were specified when the configuration was initially created.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
								},
							},
						},
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the configuration. If this parameter is empty, an ID is automatically created for the configuration.",
						},
						"project_id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The unique ID.",
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
						"check_job": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The action job performed on the project configuration.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The unique ID.",
									},
									"href": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "A relative URL.",
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
						"install_job": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The action job performed on the project configuration.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The unique ID.",
									},
									"href": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "A relative URL.",
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
						"uninstall_job": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The action job performed on the project configuration.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The unique ID.",
									},
									"href": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "A relative URL.",
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
						"type": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of a project configuration manual property.",
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

func dataSourceIbmProjectProjectConfigCanonicalToMap(model *projectv1.ProjectConfigCanonical) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Name != nil {
		modelMap["name"] = model.Name
	}
	if model.Description != nil {
		modelMap["description"] = model.Description
	}
	if model.Labels != nil {
		modelMap["labels"] = model.Labels
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
	if model.LocatorID != nil {
		modelMap["locator_id"] = model.LocatorID
	}
	if model.Input != nil {
		inputMap, err := dataSourceIbmProjectInputVariableToMap(model.Input)
		if err != nil {
			return modelMap, err
		}
		modelMap["input"] = []map[string]interface{}{inputMap}
	}
	if model.Setting != nil {
		settingMap, err := dataSourceIbmProjectProjectConfigSettingToMap(model.Setting)
		if err != nil {
			return modelMap, err
		}
		modelMap["setting"] = []map[string]interface{}{settingMap}
	}
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
	if model.CraLogs != nil {
		craLogsMap, err := dataSourceIbmProjectProjectConfigMetadataCraLogsToMap(model.CraLogs)
		if err != nil {
			return modelMap, err
		}
		modelMap["cra_logs"] = []map[string]interface{}{craLogsMap}
	}
	if model.CostEstimate != nil {
		costEstimateMap, err := dataSourceIbmProjectProjectConfigMetadataCostEstimateToMap(model.CostEstimate)
		if err != nil {
			return modelMap, err
		}
		modelMap["cost_estimate"] = []map[string]interface{}{costEstimateMap}
	}
	if model.CheckJob != nil {
		checkJobMap, err := dataSourceIbmProjectActionJobWithSummaryAndHrefToMap(model.CheckJob)
		if err != nil {
			return modelMap, err
		}
		modelMap["check_job"] = []map[string]interface{}{checkJobMap}
	}
	if model.InstallJob != nil {
		installJobMap, err := dataSourceIbmProjectActionJobWithSummaryAndHrefToMap(model.InstallJob)
		if err != nil {
			return modelMap, err
		}
		modelMap["install_job"] = []map[string]interface{}{installJobMap}
	}
	if model.UninstallJob != nil {
		uninstallJobMap, err := dataSourceIbmProjectActionJobWithSummaryAndHrefToMap(model.UninstallJob)
		if err != nil {
			return modelMap, err
		}
		modelMap["uninstall_job"] = []map[string]interface{}{uninstallJobMap}
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
	if model.Type != nil {
		modelMap["type"] = model.Type
	}
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

func dataSourceIbmProjectInputVariableToMap(model *projectv1.InputVariable) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	return modelMap, nil
}

func dataSourceIbmProjectProjectConfigSettingToMap(model *projectv1.ProjectConfigSetting) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
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

func dataSourceIbmProjectProjectConfigMetadataCraLogsToMap(model *projectv1.ProjectConfigMetadataCraLogs) (map[string]interface{}, error) {
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

func dataSourceIbmProjectProjectConfigMetadataCostEstimateToMap(model *projectv1.ProjectConfigMetadataCostEstimate) (map[string]interface{}, error) {
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

func dataSourceIbmProjectActionJobWithSummaryAndHrefToMap(model *projectv1.ActionJobWithSummaryAndHref) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ID != nil {
		modelMap["id"] = model.ID
	}
	if model.Href != nil {
		modelMap["href"] = model.Href
	}
	if model.Summary != nil {
		summaryMap, err := dataSourceIbmProjectActionJobSummaryToMap(model.Summary)
		if err != nil {
			return modelMap, err
		}
		modelMap["summary"] = []map[string]interface{}{summaryMap}
	}
	return modelMap, nil
}

func dataSourceIbmProjectActionJobSummaryToMap(model *projectv1.ActionJobSummary) (map[string]interface{}, error) {
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
