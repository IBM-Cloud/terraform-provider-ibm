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
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The project name.",
			},
			"description": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "A project descriptive text.",
			},
			"destroy_on_delete": &schema.Schema{
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "The policy that indicates whether the resources are destroyed or not when a project is deleted.",
			},
			"metadata": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The metadata of the project.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
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
							Description: "The cumulative list of needs attention items for a project.",
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
										Description: "The unique ID of a project.",
									},
									"config_id": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "The unique ID of a project.",
									},
									"config_version": &schema.Schema{
										Type:        schema.TypeInt,
										Computed:    true,
										Description: "The version number of the configuration.",
									},
								},
							},
						},
						"cumulative_needs_attention_view_err": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "True indicates that the fetch of the needs attention items failed.",
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

	project, response, err := projectClient.GetProjectWithContext(context, getProjectOptions)
	if err != nil {
		log.Printf("[DEBUG] GetProjectWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("GetProjectWithContext failed %s\n%s", err, response))
	}

	d.SetId(fmt.Sprintf("%s", *getProjectOptions.ID))

	if err = d.Set("name", project.Name); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting name: %s", err))
	}

	if err = d.Set("description", project.Description); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting description: %s", err))
	}

	if err = d.Set("destroy_on_delete", project.DestroyOnDelete); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting destroy_on_delete: %s", err))
	}

	metadata := []map[string]interface{}{}
	if project.Metadata != nil {
		modelMap, err := dataSourceIbmProjectProjectMetadataToMap(project.Metadata)
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

func dataSourceIbmProjectProjectMetadataToMap(model *projectv1.ProjectMetadata) (map[string]interface{}, error) {
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
			cumulativeNeedsAttentionViewItemMap, err := dataSourceIbmProjectCumulativeNeedsAttentionToMap(&cumulativeNeedsAttentionViewItem)
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
