// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.105.0-3c13b041-20250605-193116
 */

package drautomationservice

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.ibm.com/DRAutomation/dra-go-sdk/drautomationservicev1"
)

func DataSourceIBMPdrGetEvents() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMPdrGetEventsRead,

		Schema: map[string]*schema.Schema{
			"instance_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "instance id of the service.",
			},
			"time": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "(deprecated - use from_time) A time in either ISO 8601 or unix epoch format.",
			},
			"from_time": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "A from query time in either ISO 8601 or unix epoch format.",
			},
			"to_time": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "A to query time in either ISO 8601 or unix epoch format.",
			},
			"accept_language": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "The language requested for the return document.",
			},
			"event": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Events.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"action": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Type of action for this event.",
						},
						"api_source": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Source of API when it being executed.",
						},
						"event_id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "ID of the Activity.",
						},
						"level": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Level of the event (notice, info, warning, error).",
						},
						"message": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The (translated) message of the event.",
						},
						"message_data": &schema.Schema{
							Type:        schema.TypeMap,
							Computed:    true,
							Description: "A flexible schema placeholder to allow any JSON value (aligns with interface{} in Go).",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"metadata": &schema.Schema{
							Type:        schema.TypeMap,
							Computed:    true,
							Description: "A flexible schema placeholder to allow any JSON value (aligns with interface{} in Go).",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"resource": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Type of resource for this event.",
						},
						"time": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Time of activity in ISO 8601 - RFC3339.",
						},
						"timestamp": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Time of activity in unix epoch.",
						},
						"user": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Information about a user associated with an event.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"email": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Email of the User.",
									},
									"name": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "Name of the User.",
									},
									"user_id": &schema.Schema{
										Type:        schema.TypeString,
										Computed:    true,
										Description: "ID of user who created/caused the event.",
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

func dataSourceIBMPdrGetEventsRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	drAutomationServiceClient, err := meta.(conns.ClientSession).DrAutomationServiceV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_pdr_get_events", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	listEventsOptions := &drautomationservicev1.ListEventsOptions{}

	listEventsOptions.SetInstanceID(d.Get("instance_id").(string))
	if _, ok := d.GetOk("time"); ok {
		listEventsOptions.SetTime(d.Get("time").(string))
	}
	if _, ok := d.GetOk("from_time"); ok {
		listEventsOptions.SetFromTime(d.Get("from_time").(string))
	}
	if _, ok := d.GetOk("to_time"); ok {
		listEventsOptions.SetToTime(d.Get("to_time").(string))
	}
	if _, ok := d.GetOk("accept_language"); ok {
		listEventsOptions.SetAcceptLanguage(d.Get("accept_language").(string))
	}
	eventCollection, response, err := drAutomationServiceClient.ListEventsWithContext(context, listEventsOptions)
	if err != nil {
		detailedMsg := fmt.Sprintf("ListEventsWithContext failed: %s", err.Error())
		// Include HTTP status & raw body if available
		if response != nil {
			detailedMsg = fmt.Sprintf(
				"ListEventsWithContext failed: %s (status: %d, response: %s)",
				err.Error(), response.StatusCode, response.Result,
			)
		}
		tfErr := flex.TerraformErrorf(err, detailedMsg, "(Data) ibm_pdr_get_events", "read")
		log.Printf("[ERROR] %s", detailedMsg)
		return tfErr.GetDiag()
	}

	d.SetId(dataSourceIBMPdrGetEventsID(d))

	event := []map[string]interface{}{}
	for _, eventItem := range eventCollection.Event {
		eventItemMap, err := DataSourceIBMPdrGetEventsEventToMap(&eventItem) // #nosec G601
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_pdr_get_events", "read", "event-to-map").GetDiag()
		}
		event = append(event, eventItemMap)
	}
	if err = d.Set("event", event); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting event: %s", err), "(Data) ibm_pdr_get_events", "read", "set-event").GetDiag()
	}

	return nil
}

// dataSourceIBMPdrGetEventsID returns a reasonable ID for the list.
func dataSourceIBMPdrGetEventsID(d *schema.ResourceData) string {
	parts := strings.Split(d.Get("instance_id").(string), ":")
	if len(parts) > 7 {
		return parts[7]
	}
	return d.Get("instance_id").(string)
}

func DataSourceIBMPdrGetEventsEventToMap(model *drautomationservicev1.Event) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["action"] = *model.Action
	if model.APISource != nil {
		modelMap["api_source"] = *model.APISource
	}
	modelMap["event_id"] = *model.EventID
	modelMap["level"] = *model.Level
	modelMap["message"] = *model.Message
	if model.MessageData != nil {
		messageData := make(map[string]interface{})
		for k, v := range model.MessageData {
			messageData[k] = flex.Stringify(v)
		}
		modelMap["message_data"] = messageData
	}
	if model.Metadata != nil {
		metadata := make(map[string]interface{})
		for k, v := range model.Metadata {
			metadata[k] = flex.Stringify(v)
		}
		modelMap["metadata"] = metadata
	}
	modelMap["resource"] = *model.Resource
	modelMap["time"] = model.Time.String()
	modelMap["timestamp"] = *model.Timestamp
	if model.User != nil {
		userMap, err := DataSourceIBMPdrGetEventsEventUserToMap(model.User)
		if err != nil {
			return modelMap, err
		}
		modelMap["user"] = []map[string]interface{}{userMap}
	}
	return modelMap, nil
}

func DataSourceIBMPdrGetEventsEventUserToMap(model *drautomationservicev1.EventUser) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Email != nil {
		modelMap["email"] = *model.Email
	}
	if model.Name != nil {
		modelMap["name"] = *model.Name
	}
	modelMap["user_id"] = *model.UserID
	return modelMap, nil
}
