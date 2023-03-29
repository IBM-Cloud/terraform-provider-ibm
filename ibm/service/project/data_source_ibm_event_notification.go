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
	"github.com/damianovesperini/platform-services-go-sdk/projectv1"
)

func DataSourceIbmEventNotification() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmEventNotificationRead,

		Schema: map[string]*schema.Schema{
			"id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "The ID of the project, which uniquely identifies it.",
			},
			"description": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "A description of the instance of the event.",
			},
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name of the instance of the event.",
			},
			"enabled": &schema.Schema{
				Type:        schema.TypeBool,
				Computed:    true,
				Description: "The status of instance of the event.",
			},
			"type": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The type of the instance of event.",
			},
			"updated_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "A date/time value in the format YYYY-MM-DDTHH:mm:ssZ or YYYY-MM-DDTHH:mm:ss.sssZ, matching the date-time format as specified by RFC 3339.",
			},
			"topic_count": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The topic count of the instance of the event.",
			},
			"topic_names": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The topic names of the instance of the event.",
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
		},
	}
}

func dataSourceIbmEventNotificationRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	projectClient, err := meta.(conns.ClientSession).ProjectV1()
	if err != nil {
		return diag.FromErr(err)
	}

	getEventNotificationsIntegrationOptions := &projectv1.GetEventNotificationsIntegrationOptions{}

	getEventNotificationsIntegrationOptions.SetID(d.Get("id").(string))

	getEventNotificationsIntegrationResponse, response, err := projectClient.GetEventNotificationsIntegrationWithContext(context, getEventNotificationsIntegrationOptions)
	if err != nil {
		log.Printf("[DEBUG] GetEventNotificationsIntegrationWithContext failed %s\n%s", err, response)
		return diag.FromErr(fmt.Errorf("GetEventNotificationsIntegrationWithContext failed %s\n%s", err, response))
	}

	d.SetId(*getEventNotificationsIntegrationResponse.ID)

	if err = d.Set("description", getEventNotificationsIntegrationResponse.Description); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting description: %s", err))
	}

	if err = d.Set("name", getEventNotificationsIntegrationResponse.Name); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting name: %s", err))
	}

	if err = d.Set("enabled", getEventNotificationsIntegrationResponse.Enabled); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting enabled: %s", err))
	}

	if err = d.Set("type", getEventNotificationsIntegrationResponse.Type); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting type: %s", err))
	}

	if err = d.Set("updated_at", flex.DateTimeToString(getEventNotificationsIntegrationResponse.UpdatedAt)); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting updated_at: %s", err))
	}

	if err = d.Set("topic_count", flex.IntValue(getEventNotificationsIntegrationResponse.TopicCount)); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting topic_count: %s", err))
	}


	return nil
}
