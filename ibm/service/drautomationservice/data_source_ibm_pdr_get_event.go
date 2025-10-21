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
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.ibm.com/DRAutomation/dra-go-sdk/drautomationservicev1"
)

func DataSourceIbmPdrGetEvent() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmPdrGetEventRead,

		Schema: map[string]*schema.Schema{
			"provision_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "provision id.",
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
	}
}

func dataSourceIbmPdrGetEventRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	drAutomationServiceClient, err := meta.(conns.ClientSession).DrAutomationServiceV1()
	if err != nil {
		tfErr := flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_pdr_get_event", "read", "initialize-client")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getEventOptions := &drautomationservicev1.GetEventOptions{}

	getEventOptions.SetProvisionID(d.Get("provision_id").(string))
	getEventOptions.SetEventID(d.Get("event_id").(string))
	if _, ok := d.GetOk("accept_language"); ok {
		getEventOptions.SetAcceptLanguage(d.Get("accept_language").(string))
	}
	if _, ok := d.GetOk("if_none_match"); ok {
		getEventOptions.SetIfNoneMatch(d.Get("if_none_match").(string))
	}

	event, _, err := drAutomationServiceClient.GetEventWithContext(context, getEventOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetEventWithContext failed: %s", err.Error()), "(Data) ibm_pdr_get_event", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(dataSourceIbmPdrGetEventID(d))

	if err = d.Set("action", event.Action); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting action: %s", err), "(Data) ibm_pdr_get_event", "read", "set-action").GetDiag()
	}

	if !core.IsNil(event.APISource) {
		if err = d.Set("api_source", event.APISource); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting api_source: %s", err), "(Data) ibm_pdr_get_event", "read", "set-api_source").GetDiag()
		}
	}

	if err = d.Set("level", event.Level); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting level: %s", err), "(Data) ibm_pdr_get_event", "read", "set-level").GetDiag()
	}

	if err = d.Set("message", event.Message); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting message: %s", err), "(Data) ibm_pdr_get_event", "read", "set-message").GetDiag()
	}

	if !core.IsNil(event.MessageData) {
		convertedMap := make(map[string]interface{}, len(event.MessageData))
		for k, v := range event.MessageData {
			convertedMap[k] = v
		}
		if err = d.Set("message_data", flex.Flatten(convertedMap)); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting message_data: %s", err), "(Data) ibm_pdr_get_event", "read", "set-message_data").GetDiag()
		}
	}

	if !core.IsNil(event.Metadata) {
		convertedMap := make(map[string]interface{}, len(event.Metadata))
		for k, v := range event.Metadata {
			convertedMap[k] = v
		}
		if err = d.Set("metadata", flex.Flatten(convertedMap)); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting metadata: %s", err), "(Data) ibm_pdr_get_event", "read", "set-metadata").GetDiag()
		}
	}

	if err = d.Set("resource", event.Resource); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting resource: %s", err), "(Data) ibm_pdr_get_event", "read", "set-resource").GetDiag()
	}

	if err = d.Set("time", flex.DateTimeToString(event.Time)); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting time: %s", err), "(Data) ibm_pdr_get_event", "read", "set-time").GetDiag()
	}

	if err = d.Set("timestamp", event.Timestamp); err != nil {
		return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting timestamp: %s", err), "(Data) ibm_pdr_get_event", "read", "set-timestamp").GetDiag()
	}

	if !core.IsNil(event.User) {
		user := []map[string]interface{}{}
		userMap, err := DataSourceIbmPdrGetEventEventUserToMap(event.User)
		if err != nil {
			return flex.DiscriminatedTerraformErrorf(err, err.Error(), "(Data) ibm_pdr_get_event", "read", "user-to-map").GetDiag()
		}
		user = append(user, userMap)
		if err = d.Set("user", user); err != nil {
			return flex.DiscriminatedTerraformErrorf(err, fmt.Sprintf("Error setting user: %s", err), "(Data) ibm_pdr_get_event", "read", "set-user").GetDiag()
		}
	}

	return nil
}

// dataSourceIbmPdrGetEventID returns a reasonable ID for the list.
func dataSourceIbmPdrGetEventID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func DataSourceIbmPdrGetEventEventUserToMap(model *drautomationservicev1.EventUser) (map[string]interface{}, error) {
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
