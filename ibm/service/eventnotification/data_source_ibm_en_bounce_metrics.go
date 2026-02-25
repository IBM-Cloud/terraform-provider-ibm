// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package eventnotification

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/event-notifications-go-admin-sdk/eventnotificationsv1"
)

func DataSourceIBMEnBounceMetrics() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMEnBounceMetricsRead,

		Schema: map[string]*schema.Schema{
			"instance_id": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Unique identifier for IBM Cloud Event Notifications instance.",
			},
			"destination_type": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "Destination type. Allowed values are [smtp_custom].",
			},
			"gte": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "GTE (greater than equal), start timestamp in UTC.",
			},
			"lte": &schema.Schema{
				Type:        schema.TypeString,
				Required:    true,
				Description: "LTE (less than equal), end timestamp in UTC.",
			},
			"destination_id": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Unique identifier for Destination.",
			},
			"subscription_id": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Unique identifier for Subscription.",
			},
			"source_id": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Unique identifier for Source.",
			},
			"email_to": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Receiver email id.",
			},
			"notification_id": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Notification Id.",
			},
			"subject": &schema.Schema{
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Email subject.",
			},
			"metrics": &schema.Schema{
				Type:        schema.TypeList,
				Computed:    true,
				Description: "array of bounce metrics.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"email_address": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Email address.",
						},
						"subject": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Subject.",
						},
						"error_message": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Error message.",
						},
						"ip_address": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "IP address.",
						},
						"subscription_id": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Subscription ID.",
						},
						"timestamp": &schema.Schema{
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Bounced at.",
						},
					},
				},
			},
			"total_count": &schema.Schema{
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "total number of bounce metrics.",
			},
		},
	}
}

func dataSourceIBMEnBounceMetricsRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	eventNotificationsClient, err := meta.(conns.ClientSession).EventNotificationsApiV1()
	if err != nil {
		tfErr := flex.TerraformErrorf(err, err.Error(), "(Data) ibm_en_bounce_metrics", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	getBounceMetricsOptions := &eventnotificationsv1.GetBounceMetricsOptions{}

	getBounceMetricsOptions.SetInstanceID(d.Get("instance_id").(string))
	getBounceMetricsOptions.SetDestinationType(d.Get("destination_type").(string))
	getBounceMetricsOptions.SetGte(d.Get("gte").(string))
	getBounceMetricsOptions.SetLte(d.Get("lte").(string))
	if _, ok := d.GetOk("destination_id"); ok {
		getBounceMetricsOptions.SetDestinationID(d.Get("destination_id").(string))
	}
	if _, ok := d.GetOk("subscription_id"); ok {
		getBounceMetricsOptions.SetSubscriptionID(d.Get("subscription_id").(string))
	}
	if _, ok := d.GetOk("source_id"); ok {
		getBounceMetricsOptions.SetSourceID(d.Get("source_id").(string))
	}
	if _, ok := d.GetOk("email_to"); ok {
		getBounceMetricsOptions.SetEmailTo(d.Get("email_to").(string))
	}
	if _, ok := d.GetOk("notification_id"); ok {
		getBounceMetricsOptions.SetNotificationID(d.Get("notification_id").(string))
	}
	if _, ok := d.GetOk("subject"); ok {
		getBounceMetricsOptions.SetSubject(d.Get("subject").(string))
	}

	bounceMetrics, _, err := eventNotificationsClient.GetBounceMetricsWithContext(context, getBounceMetricsOptions)
	if err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("GetBounceMetricsWithContext failed: %s", err.Error()), "(Data) ibm_en_bounce_metrics", "read")
		log.Printf("[DEBUG]\n%s", tfErr.GetDebugMessage())
		return tfErr.GetDiag()
	}

	d.SetId(dataSourceIBMEnBounceMetricsID(d))

	metrics := []map[string]interface{}{}
	if bounceMetrics.Metrics != nil {
		for _, modelItem := range bounceMetrics.Metrics {
			modelMap, err := dataSourceIBMEnBounceMetricsBounceMetricItemToMap(&modelItem)
			if err != nil {
				tfErr := flex.TerraformErrorf(err, err.Error(), "(Data) ibm_en_bounce_metrics", "read")
				return tfErr.GetDiag()
			}
			metrics = append(metrics, modelMap)
		}
	}
	if err = d.Set("metrics", metrics); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting metrics: %s", err), "(Data) ibm_en_bounce_metrics", "read")
		return tfErr.GetDiag()
	}

	if err = d.Set("total_count", flex.IntValue(bounceMetrics.TotalCount)); err != nil {
		tfErr := flex.TerraformErrorf(err, fmt.Sprintf("Error setting total_count: %s", err), "(Data) ibm_en_bounce_metrics", "read")
		return tfErr.GetDiag()
	}

	return nil
}

// dataSourceIBMEnBounceMetricsID returns a reasonable ID for the list.
func dataSourceIBMEnBounceMetricsID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}

func dataSourceIBMEnBounceMetricsBounceMetricItemToMap(model *eventnotificationsv1.BounceMetricItem) (map[string]interface{}, error) {
	modelMap := make(map[string]interface{})
	modelMap["email_address"] = model.EmailAddress
	modelMap["subject"] = model.Subject
	modelMap["error_message"] = model.ErrorMessage
	if model.IPAddress != nil {
		modelMap["ip_address"] = model.IPAddress
	}
	if model.SubscriptionID != nil {
		modelMap["subscription_id"] = model.SubscriptionID
	}
	modelMap["timestamp"] = model.Timestamp.String()
	return modelMap, nil
}
