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
	"github.com/IBM/go-sdk-core/v5/core"
	"github.ibm.com/DRAutomation/dra-go-sdk/powerhaautomationservicev1"
)

func DataSourceIBMPhaGetServiceInstanceEvent() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMPhaGetServiceInstanceEventRead,

		Schema: map[string]*schema.Schema{
			"pha_instance_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "instance id of instance to provision.",
			},
			"event_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Event ID.",
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
	}
}

func dataSourceIBMPhaGetServiceInstanceEventRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	powerhaAutomationServiceClient, err := meta.(conns.ClientSession).PowerhaAutomationServiceV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_pha_get_service_instance_event", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getServiceInstanceEventOptions := &powerhaautomationservicev1.GetServiceInstanceEventOptions{}

	getServiceInstanceEventOptions.SetPhaInstanceID(d.Get("pha_instance_id").(string))
	getServiceInstanceEventOptions.SetEventID(d.Get("event_id").(string))
	if _, ok := d.GetOk("accept_language"); ok {
		getServiceInstanceEventOptions.SetAcceptLanguage(d.Get("accept_language").(string))
	}
	if _, ok := d.GetOk("if_none_match"); ok {
		getServiceInstanceEventOptions.SetIfNoneMatch(d.Get("if_none_match").(string))
	}

	phaEvent, response, err := powerhaAutomationServiceClient.GetServiceInstanceEventWithContext(context, getServiceInstanceEventOptions)
	if err != nil {
		detailedMsg := fmt.Sprintf("GetServiceInstanceEventWithContext failed: %s", err.Error())
		// Include HTTP status & raw body if available
		if response != nil {
			detailedMsg = fmt.Sprintf(
				"GetServiceInstanceEventWithContext failed: %s (status: %d, response: %s)",
				err.Error(), response.StatusCode, response.Result,
			)
		}
		tfErr := flex.TerraformErrorf(err, detailedMsg, "ibm_pha_get_service_instance_event", "create")
		log.Printf("[ERROR] %s", detailedMsg)
		return tfErr.GetDiag()

		// tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetServiceInstanceEventWithContext failed: %s", err.Error()), "(Data) ibm_pha_get_service_instance_event", "read")
		// log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		// return tfErr.GetDiag()
	}

	d.SetId(dataSourceIBMPhaGetServiceInstanceEventID(d))

	if err = d.Set("action", phaEvent.Action); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting action: %s", err), "(Data) ibm_pha_get_service_instance_event", "read", "set-action").GetDiag()
	}

	if !core.IsNil(phaEvent.APISource) {
		if err = d.Set("api_source", phaEvent.APISource); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting api_source: %s", err), "(Data) ibm_pha_get_service_instance_event", "read", "set-api_source").GetDiag()
		}
	}

	if err = d.Set("level", phaEvent.Level); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting level: %s", err), "(Data) ibm_pha_get_service_instance_event", "read", "set-level").GetDiag()
	}

	if err = d.Set("message", phaEvent.Message); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting message: %s", err), "(Data) ibm_pha_get_service_instance_event", "read", "set-message").GetDiag()
	}

	if !core.IsNil(phaEvent.MessageData) {
		convertedMap := make(map[string]interface{}, len(phaEvent.MessageData))
		for k, v := range phaEvent.MessageData {
			convertedMap[k] = v
		}
		if err = d.Set("message_data", flex.Flatten(convertedMap)); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting message_data: %s", err), "(Data) ibm_pha_get_service_instance_event", "read", "set-message_data").GetDiag()
		}
	}

	if !core.IsNil(phaEvent.MetaData) {
		if err = d.Set("meta_data", phaEvent.MetaData); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting meta_data: %s", err), "(Data) ibm_pha_get_service_instance_event", "read", "set-meta_data").GetDiag()
		}
	}

	if err = d.Set("resource", phaEvent.Resource); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting resource: %s", err), "(Data) ibm_pha_get_service_instance_event", "read", "set-resource").GetDiag()
	}

	if err = d.Set("time", phaEvent.Time); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting time: %s", err), "(Data) ibm_pha_get_service_instance_event", "read", "set-time").GetDiag()
	}

	if err = d.Set("time_stamp", phaEvent.TimeStamp); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting time_stamp: %s", err), "(Data) ibm_pha_get_service_instance_event", "read", "set-time_stamp").GetDiag()
	}

	if !core.IsNil(phaEvent.User) {
		user := []map[string]interface{}{}
		userMap, err := DataSourceIBMPhaGetServiceInstanceEventPhaEventUserToMap(phaEvent.User)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_pha_get_service_instance_event", "read", "user-to-map").GetDiag()
		}
		user = append(user, userMap)
		if err = d.Set("user", user); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting user: %s", err), "(Data) ibm_pha_get_service_instance_event", "read", "set-user").GetDiag()
		}
	}

	return nil
}

// dataSourceIBMPhaGetServiceInstanceEventID returns a reasonable ID for the list.
func dataSourceIBMPhaGetServiceInstanceEventID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func DataSourceIBMPhaGetServiceInstanceEventPhaEventUserToMap(model *powerhaautomationservicev1.PhaEventUser) (map[string]interface{}, error) {
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
