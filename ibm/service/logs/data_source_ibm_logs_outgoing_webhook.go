// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package logs

import (
	"context"
	"fmt"
	"log"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/observability-c/dragonlog-logs-go-sdk/logsv0"
)

func DataSourceIbmLogsOutgoingWebhook() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmLogsOutgoingWebhookRead,

		Schema: map[string]*schema.Schema{
			"logs_outgoing_webhook_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Outbound webhook ID.",
			},
			"type": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Outbound webhook type.",
			},
			"name": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The name of the outbound webhook.",
			},
			"url": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL of the outbound webhook.",
			},
			"created_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The creation time of the outbound webhook.",
			},
			"updated_at": &schema.Schema{
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The update time of the outbound webhook.",
			},
			"external_id": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "The external ID of the outbound webhook.",
			},
			"ibm_event_notifications": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "The configuration of an IBM Event Notifications outbound webhook.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"event_notifications_instance_id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The instance ID of the IBM Event Notifications configuration.",
						},
						"region_id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The region ID of the IBM Event Notifications configuration.",
						},
					},
				},
			},
		},
	}
}

func dataSourceIbmLogsOutgoingWebhookRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logsClient, err := meta.(conns.ClientSession).LogsV0()
	if err != nil {
		tfErr := flex.TerraformErrorf(err, err.Error(), "(Data) ibm_logs_outgoing_webhook", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	region := getLogsInstanceRegion(logsClient, d)
	instanceId := d.Get("instance_id").(string)
	logsClient = getClientWithLogsInstanceEndpoint(logsClient, instanceId, region, getLogsInstanceEndpointType(logsClient, d))

	getOutgoingWebhookOptions := &logsv0.GetOutgoingWebhookOptions{}

	getOutgoingWebhookOptions.SetID(d.Get("logs_outgoing_webhook_id").(string))

	outgoingWebhookIntf, _, err := logsClient.GetOutgoingWebhookWithContext(context, getOutgoingWebhookOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetOutgoingWebhookWithContext failed: %s", err.Error()), "(Data) ibm_logs_outgoing_webhook", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}
	outgoingWebhook := outgoingWebhookIntf.(*logsv0.OutgoingWebhook)

	d.SetId(fmt.Sprintf("%s", *getOutgoingWebhookOptions.ID))

	if err = d.Set("type", outgoingWebhook.Type); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting type: %s", err), "(Data) ibm_logs_outgoing_webhook", "read")
		return tfErr.GetDiag()
	}

	if err = d.Set("name", outgoingWebhook.Name); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting name: %s", err), "(Data) ibm_logs_outgoing_webhook", "read")
		return tfErr.GetDiag()
	}

	if err = d.Set("url", outgoingWebhook.URL); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting url: %s", err), "(Data) ibm_logs_outgoing_webhook", "read")
		return tfErr.GetDiag()
	}

	if err = d.Set("created_at", flex.DateTimeToString(outgoingWebhook.CreatedAt)); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting created_at: %s", err), "(Data) ibm_logs_outgoing_webhook", "read")
		return tfErr.GetDiag()
	}

	if err = d.Set("updated_at", flex.DateTimeToString(outgoingWebhook.UpdatedAt)); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting updated_at: %s", err), "(Data) ibm_logs_outgoing_webhook", "read")
		return tfErr.GetDiag()
	}

	if err = d.Set("external_id", flex.IntValue(outgoingWebhook.ExternalID)); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting external_id: %s", err), "(Data) ibm_logs_outgoing_webhook", "read")
		return tfErr.GetDiag()
	}

	ibmEventNotifications := []map[string]interface{}{}
	if outgoingWebhook.IbmEventNotifications != nil {
		modelMap, err := DataSourceIbmLogsOutgoingWebhookOutgoingWebhooksV1IbmEventNotificationsConfigToMap(outgoingWebhook.IbmEventNotifications)
		if err != nil {
			tfErr := flex.TerraformErrorf(err, err.Error(), "(Data) ibm_logs_outgoing_webhook", "read")
			return tfErr.GetDiag()
		}
		ibmEventNotifications = append(ibmEventNotifications, modelMap)
	}
	if err = d.Set("ibm_event_notifications", ibmEventNotifications); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting ibm_event_notifications: %s", err), "(Data) ibm_logs_outgoing_webhook", "read")
		return tfErr.GetDiag()
	}

	return nil
}

func DataSourceIbmLogsOutgoingWebhookOutgoingWebhooksV1IbmEventNotificationsConfigToMap(model *logsv0.OutgoingWebhooksV1IbmEventNotificationsConfig) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["event_notifications_instance_id"] = model.EventNotificationsInstanceID.String()
	modelMap["region_id"] = *model.RegionID
	return modelMap, nil
}
