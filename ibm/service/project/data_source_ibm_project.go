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
						"approved_version": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The project configuration version.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
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
									"version": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The version number of the configuration.",
									},
									"href": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "A relative URL.",
									},
								},
							},
						},
						"installed_version": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "The project configuration version.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
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
									"version": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The version number of the configuration.",
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
							Description: "The project configuration definition summary.",
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
		modelMap, err := dataSourceIbmProjectProjectDefinitionTerraformToMap(projectCanonical.Definition)
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
			modelMap, err := dataSourceIbmProjectProjectConfigCollectionMemberTerraformToMap(&modelItem)
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

func dataSourceIbmProjectProjectDefinitionTerraformToMap(model *projectv1.ProjectDefinitionTerraform) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Name != nil {
		modelMap["name"] = model.Name
	}
	if model.Description != nil {
		modelMap["description"] = model.Description
	}
	if model.DestroyOnDelete != nil {
		modelMap["destroy_on_delete"] = model.DestroyOnDelete
	}
	return modelMap, nil
}

func dataSourceIbmProjectProjectConfigCollectionMemberTerraformToMap(model *projectv1.ProjectConfigCollectionMemberTerraform) (map[string]interface{}, error) {
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
	if model.ApprovedVersion != nil {
		approvedVersionMap, err := dataSourceIbmProjectProjectConfigVersionSummaryToMap(model.ApprovedVersion)
		if err != nil {
			return modelMap, err
		}
		modelMap["approved_version"] = []map[string]interface{}{approvedVersionMap}
	}
	if model.InstalledVersion != nil {
		installedVersionMap, err := dataSourceIbmProjectProjectConfigVersionSummaryToMap(model.InstalledVersion)
		if err != nil {
			return modelMap, err
		}
		modelMap["installed_version"] = []map[string]interface{}{installedVersionMap}
	}
	if model.Definition != nil {
		definitionMap, err := dataSourceIbmProjectProjectConfigDefinitionSummaryToMap(model.Definition)
		if err != nil {
			return modelMap, err
		}
		modelMap["definition"] = []map[string]interface{}{definitionMap}
	}
	if model.CheckJob != nil {
		checkJobMap, err := dataSourceIbmProjectActionJobWithIdAndHrefToMap(model.CheckJob)
		if err != nil {
			return modelMap, err
		}
		modelMap["check_job"] = []map[string]interface{}{checkJobMap}
	}
	if model.InstallJob != nil {
		installJobMap, err := dataSourceIbmProjectActionJobWithIdAndHrefToMap(model.InstallJob)
		if err != nil {
			return modelMap, err
		}
		modelMap["install_job"] = []map[string]interface{}{installJobMap}
	}
	if model.UninstallJob != nil {
		uninstallJobMap, err := dataSourceIbmProjectActionJobWithIdAndHrefToMap(model.UninstallJob)
		if err != nil {
			return modelMap, err
		}
		modelMap["uninstall_job"] = []map[string]interface{}{uninstallJobMap}
	}
	if model.Href != nil {
		modelMap["href"] = model.Href
	}
	return modelMap, nil
}

func dataSourceIbmProjectProjectConfigVersionSummaryToMap(model *projectv1.ProjectConfigVersionSummary) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.NeedsAttentionState != nil {
		modelMap["needs_attention_state"] = model.NeedsAttentionState
	}
	modelMap["state"] = model.State
	modelMap["version"] = flex.IntValue(model.Version)
	if model.Href != nil {
		modelMap["href"] = model.Href
	}
	return modelMap, nil
}

func dataSourceIbmProjectProjectConfigDefinitionSummaryToMap(model *projectv1.ProjectConfigDefinitionSummary) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["name"] = model.Name
	if model.Description != nil {
		modelMap["description"] = model.Description
	}
	return modelMap, nil
}

func dataSourceIbmProjectActionJobWithIdAndHrefToMap(model *projectv1.ActionJobWithIdAndHref) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.ID != nil {
		modelMap["id"] = model.ID
	}
	if model.Href != nil {
		modelMap["href"] = model.Href
	}
	return modelMap, nil
}
