// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.116.0-df613dbc-20260803-154903
 */

package drautomationservice

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/dra-go-sdk/drautomationservicev1"
)

func DataSourceIBMPdrEvents() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMPdrEventsRead,

		Schema: map[string]*schema.Schema{
			"instance_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Service Instance ID.",
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
			"events": &schema.Schema{
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
							Optional:    true,
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
							Description: "Any message data associated with the event.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"metadata": &schema.Schema{
							Type:        schema.TypeMap,
							Computed:    true,
							Description: "Any metadata associated with the event.",
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
							Description: "Time of activity in Unix epoch.",
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

func dataSourceIBMPdrEventsRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	drAutomationServiceClient, err := meta.(conns.ClientSession).DrAutomationServiceV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_pdr_events", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	listServiceInstanceEventsOptions := &drautomationservicev1.ListServiceInstanceEventsOptions{}

	listServiceInstanceEventsOptions.SetInstanceID(d.Get("instance_id").(string))
	if _, ok := d.GetOk("time"); ok {
		fmtDateTimeTime, err := core.ParseDateTime(d.Get("time").(string))
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_pdr_events", "read", "parse-time").GetDiag()
		}
		listServiceInstanceEventsOptions.SetTime(&fmtDateTimeTime)
	}
	if _, ok := d.GetOk("from_time"); ok {
		fmtDateTimeFromTime, err := core.ParseDateTime(d.Get("from_time").(string))
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_pdr_events", "read", "parse-from_time").GetDiag()
		}
		listServiceInstanceEventsOptions.SetFromTime(&fmtDateTimeFromTime)
	}
	if _, ok := d.GetOk("to_time"); ok {
		fmtDateTimeToTime, err := core.ParseDateTime(d.Get("to_time").(string))
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_pdr_events", "read", "parse-to_time").GetDiag()
		}
		listServiceInstanceEventsOptions.SetToTime(&fmtDateTimeToTime)
	}
	if _, ok := d.GetOk("accept_language"); ok {
		listServiceInstanceEventsOptions.SetAcceptLanguage(d.Get("accept_language").(string))
	}

	eventCollection, response, err := drAutomationServiceClient.ListServiceInstanceEventsWithContext(context, listServiceInstanceEventsOptions)
	if err != nil {
		detailedMsg := fmt.Sprintf("ListEventsWithContext failed: %s", err.Error())
		// Include HTTP status & raw body if available
		if response != nil {
			detailedMsg = fmt.Sprintf(
				"ListEventsWithContext failed: %s (status: %d, response: %s)",
				err.Error(), response.StatusCode, response.Result,
			)
		}
		tfErr := flex.TerraformErrorf(err, detailedMsg, "(Data) ibm_pdr_events", "read")
		log.Printf("[ERROR] %s", detailedMsg)
		return tfErr.GetDiag()
	}

	d.SetId(dataSourceIBMPdrEventsID(d))

	events := []map[string]interface{}{}
	for _, eventsItem := range eventCollection.Events {
		eventsItemMap, err := DataSourceIBMPdrEventsEventToMap(&eventsItem) // #nosec G601
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_pdr_events", "read", "events-to-map").GetDiag()
		}
		events = append(events, eventsItemMap)
	}
	if err = d.Set("events", events); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting events: %s", err), "(Data) ibm_pdr_events", "read", "set-events").GetDiag()
	}

	return nil
}

// dataSourceIBMPdrEventsID returns a reasonable ID for the list.
func dataSourceIBMPdrEventsID(d *schema.ResourceData) string {
	parts := strings.Split(d.Get("instance_id").(string), ":")
	if len(parts) > 7 {
		return parts[7]
	}
	return d.Get("instance_id").(string)
}

func DataSourceIBMPdrEventsEventToMap(model *drautomationservicev1.Event) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Action != nil {
		modelMap["action"] = *model.Action
	}
	if model.APISource != nil {
		modelMap["api_source"] = *model.APISource
	}
	if model.EventID != nil {
		modelMap["event_id"] = *model.EventID
	}
	modelMap["level"] = *model.Level
	if model.Message != nil {
		modelMap["message"] = *model.Message
	}
	messageData := make(map[string]interface{})
	for k, v := range model.MessageData {
		messageData[k] = flex.Stringify(v)
	}
	modelMap["message_data"] = messageData
	metadata := make(map[string]interface{})
	for k, v := range model.Metadata {
		metadata[k] = flex.Stringify(v)
	}
	modelMap["metadata"] = metadata
	if model.Resource != nil {
		modelMap["resource"] = *model.Resource
	}
	if model.Time != nil {
		modelMap["time"] = model.Time.String()
	}
	if model.Timestamp != nil {
		modelMap["timestamp"] = *model.Timestamp
	}
	if model.User != nil {
		userMap, err := DataSourceIBMPdrEventsEventUserToMap(model.User)
		if err != nil {
			return modelMap, err
		}
		modelMap["user"] = []map[string]interface{}{userMap}
	}
	return modelMap, nil
}

func DataSourceIBMPdrEventsEventUserToMap(model *drautomationservicev1.EventUser) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	if model.Email != nil {
		modelMap["email"] = *model.Email
	}
	if model.Name != nil {
		modelMap["name"] = *model.Name
	}
	if model.UserID != nil {
		modelMap["user_id"] = *model.UserID
	}
	return modelMap, nil
}
