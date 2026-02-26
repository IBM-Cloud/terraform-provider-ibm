// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.108.0-56772134-20251111-102802
 */

package powerhaautomationservice

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.ibm.com/DRAutomation/dra-go-sdk/powerhaautomationservicev1"
)

func DataSourceIBMPhaListServiceInstanceEvents() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMPhaListServiceInstanceEventsRead,

		Schema: map[string]*schema.Schema{
			"pha_instance_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "instance id of instance to provision.",
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
			"if_none_match": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "ETag for conditional requests (optional).",
			},
			"events": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "Pha automation Events.",
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
							Description: "Dynamic key-value data related to the event message.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"meta_data": &schema.Schema{
							Type:        schema.TypeMap,
							Computed:    true,
							Description: "Metadata providing additional context for the event.",
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
						"time_stamp": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Timestamp indicating when the event occurred.",
						},
						"user": &schema.Schema{
							Type:        schema.TypeList,
							Computed:    true,
							Description: "events for pha user.",
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

func dataSourceIBMPhaListServiceInstanceEventsRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	powerhaAutomationServiceClient, err := meta.(conns.ClientSession).PowerhaAutomationServiceV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_pha_list_service_instance_events", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	listServiceInstanceEventsOptions := &powerhaautomationservicev1.ListServiceInstanceEventsOptions{}

	listServiceInstanceEventsOptions.SetPhaInstanceID(d.Get("pha_instance_id").(string))
	if _, ok := d.GetOk("time"); ok {
		listServiceInstanceEventsOptions.SetTime(d.Get("time").(string))
	}
	if _, ok := d.GetOk("from_time"); ok {
		listServiceInstanceEventsOptions.SetFromTime(d.Get("from_time").(string))
	}
	if _, ok := d.GetOk("to_time"); ok {
		listServiceInstanceEventsOptions.SetToTime(d.Get("to_time").(string))
	}
	if _, ok := d.GetOk("accept_language"); ok {
		listServiceInstanceEventsOptions.SetAcceptLanguage(d.Get("accept_language").(string))
	}
	if _, ok := d.GetOk("if_none_match"); ok {
		listServiceInstanceEventsOptions.SetIfNoneMatch(d.Get("if_none_match").(string))
	}

	phaEventCollection, response, err := powerhaAutomationServiceClient.ListServiceInstanceEventsWithContext(context, listServiceInstanceEventsOptions)
	if err != nil {
		detailedMsg := fmt.Sprintf("ListServiceInstanceEventsWithContext failed: %s", err.Error())
		// Include HTTP status & raw body if available
		if response != nil {
			detailedMsg = fmt.Sprintf(
				"ListServiceInstanceEventsWithContext failed: %s (status: %d, response: %s)",
				err.Error(), response.StatusCode, response.Result,
			)
		}
		tfErr := flex.TerraformErrorf(err, detailedMsg, "ibm_pha_list_service_instance_events", "create")
		log.Printf("[ERROR] %s", detailedMsg)
		return tfErr.GetDiag()
		// tfErr := flex.TerraformErrorf(err, fmt.Sprintf("ListServiceInstanceEventsWithContext failed: %s", err.Error()), "(Data) ibm_pha_list_service_instance_events", "read")
		// log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		// return tfErr.GetDiag()
	}

	d.SetId(dataSourceIBMPhaListServiceInstanceEventsID(d))

	events := []map[string]interface{}{}
	for _, eventsItem := range phaEventCollection.Events {
		eventsItemMap, err := DataSourceIBMPhaListServiceInstanceEventsPhaEventToMap(&eventsItem) // #nosec G601
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_pha_list_service_instance_events", "read", "events-to-map").GetDiag()
		}
		events = append(events, eventsItemMap)
	}
	if err = d.Set("events", events); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting events: %s", err), "(Data) ibm_pha_list_service_instance_events", "read", "set-events").GetDiag()
	}

	return nil
}

// dataSourceIBMPhaListServiceInstanceEventsID returns a reasonable ID for the list.
func dataSourceIBMPhaListServiceInstanceEventsID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

// func DataSourceIBMPhaListServiceInstanceEventsPhaEventToMap(model *powerhaautomationservicev1.PhaEvent) (map[string]interface{}, error) {
// 	modelMap := make(map[string]interface{})
// 	modelMap["action"] = *model.Action
// 	if model.APISource != nil {
// 		modelMap["api_source"] = *model.APISource
// 	}
// 	modelMap["event_id"] = *model.EventID
// 	modelMap["level"] = *model.Level
// 	modelMap["message"] = *model.Message
// 	if model.MessageData != nil {
// 		messageData := make(map[string]interface{})
// 		for k, v := range model.MessageData {
// 			messageData[k] = flex.Stringify(v)
// 		}
// 		modelMap["message_data"] = messageData
// 	}
// 	if model.MetaData != nil {
// 		metaData := make(map[string]interface{})
// 		for k, v := range model.MetaData {
// 			metaData[k] = flex.Stringify(v)
// 		}
// 		modelMap["meta_data"] = metaData
// 	}
// 	modelMap["resource"] = *model.Resource
// 	modelMap["time"] = *model.Time
// 	modelMap["time_stamp"] = *model.TimeStamp
// 	if model.User != nil {
// 		userMap, err := DataSourceIBMPhaListServiceInstanceEventsPhaEventUserToMap(model.User)
// 		if err != nil {
// 			return modelMap, err
// 		}
// 		modelMap["user"] = []map[string]interface{}{userMap}
// 	}
// 	return modelMap, nil
// }

func DataSourceIBMPhaListServiceInstanceEventsPhaEventToMap(model *powerhaautomationservicev1.PhaEvent) (map[string]interface{}, error) {
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
	if model.Level != nil {
		modelMap["level"] = *model.Level
	}
	if model.Message != nil {
		modelMap["message"] = *model.Message
	}

	if model.MessageData != nil {
		messageData := make(map[string]interface{})
		for k, v := range model.MessageData {
			messageData[k] = flex.Stringify(v)
		}
		modelMap["message_data"] = messageData
	}

	if model.MetaData != nil {
		metaData := make(map[string]interface{})
		for k, v := range model.MetaData {
			metaData[k] = flex.Stringify(v)
		}
		modelMap["meta_data"] = metaData
	}

	if model.Resource != nil {
		modelMap["resource"] = *model.Resource
	}
	if model.Time != nil {
		modelMap["time"] = *model.Time
	}
	if model.TimeStamp != nil {
		modelMap["time_stamp"] = *model.TimeStamp
	} else {
		modelMap["time_stamp"] = ""
	}

	if model.User != nil {
		userMap, err := DataSourceIBMPhaListServiceInstanceEventsPhaEventUserToMap(model.User)
		if err != nil {
			return modelMap, err
		}
		modelMap["user"] = []map[string]interface{}{userMap}
	}

	return modelMap, nil
}

func DataSourceIBMPhaListServiceInstanceEventsPhaEventUserToMap(model *powerhaautomationservicev1.PhaEventUser) (map[string]interface{}, error) {
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
