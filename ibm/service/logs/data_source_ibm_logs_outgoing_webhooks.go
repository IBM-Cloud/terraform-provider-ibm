// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package logs

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/observability-c/dragonlog-logs-go-sdk/logsv0"
)

func DataSourceIbmLogsOutgoingWebhooks() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIbmLogsOutgoingWebhooksRead,

		Schema: map[string]*schema.Schema{
			"type": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "ibm_event_notifications",
				Description: "Outbound webhook type.",
			},
			"outgoing_webhooks": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of deployed outbound webhooks.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The ID of the outbound webhook.",
						},
						"name": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of the outbound webhook.",
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
					},
				},
			},
		},
	}
}

func dataSourceIbmLogsOutgoingWebhooksRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	logsClient, err := meta.(conns.ClientSession).LogsV0()
	if err != nil {
		tfErr := flex.TerraformErrorf(err, err.Error(), "(Data) ibm_logs_outgoing_webhooks", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	region := getLogsInstanceRegion(logsClient, d)
	instanceId := d.Get("instance_id").(string)
	logsClient = getClientWithLogsInstanceEndpoint(logsClient, instanceId, region, getLogsInstanceEndpointType(logsClient, d))

	listOutgoingWebhooksOptions := &logsv0.ListOutgoingWebhooksOptions{}

	if _, ok := d.GetOk("type"); ok {
		listOutgoingWebhooksOptions.SetType(d.Get("type").(string))
	}

	outgoingWebhookCollection, _, err := logsClient.ListOutgoingWebhooksWithContext(context, listOutgoingWebhooksOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("ListOutgoingWebhooksWithContext failed: %s", err.Error()), "(Data) ibm_logs_outgoing_webhooks", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(dataSourceIbmLogsOutgoingWebhooksID(d))

	outgoingWebhooks := []map[string]interface{}{}
	if outgoingWebhookCollection.OutgoingWebhooks != nil {
		for _, modelItem := range outgoingWebhookCollection.OutgoingWebhooks {
			modelMap, err := DataSourceIbmLogsOutgoingWebhooksOutgoingWebhookSummaryToMap(&modelItem)
			if err != nil {
				tfErr := flex.TerraformErrorf(err, err.Error(), "(Data) ibm_logs_outgoing_webhooks", "read")
				return tfErr.GetDiag()
			}
			outgoingWebhooks = append(outgoingWebhooks, modelMap)
		}
	}
	if err = d.Set("outgoing_webhooks", outgoingWebhooks); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting outgoing_webhooks: %s", err), "(Data) ibm_logs_outgoing_webhooks", "read")
		return tfErr.GetDiag()
	}

	return nil
}

// dataSourceIbmLogsOutgoingWebhooksID returns a reasonable ID for the list.
func dataSourceIbmLogsOutgoingWebhooksID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func DataSourceIbmLogsOutgoingWebhooksOutgoingWebhookSummaryToMap(model *logsv0.OutgoingWebhookSummary) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["id"] = model.ID.String()
	modelMap["name"] = *model.Name
	modelMap["url"] = *model.URL
	modelMap["created_at"] = model.CreatedAt.String()
	modelMap["updated_at"] = model.UpdatedAt.String()
	modelMap["external_id"] = flex.IntValue(model.ExternalID)
	return modelMap, nil
}
