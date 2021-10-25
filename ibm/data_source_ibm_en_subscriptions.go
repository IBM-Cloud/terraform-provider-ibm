// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	en "github.com/IBM/event-notifications-go-admin-sdk/eventnotificationsv1"
)

func dataSourceIBMEnSubscriptions() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceIBMEnSubscriptionsRead,

		Schema: map[string]*schema.Schema{
			"instance_guid": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Unique identifier for IBM Cloud Event Notifications instance.",
			},
			"search_key": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Filter the subscriptions by name",
			},
			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Number of subscriptions.",
			},
			"subscriptions": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of subscriptions.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "ID of the subscription.",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Name of the subscription.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Description of the subscription.",
						},
						"destination_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "ID of the destination.",
						},
						"destination_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The destination name.",
						},
						"destination_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "The type of destination.",
						},
						"topic_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "ID of the topic.",
						},
						"topic_name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Topic name.",
						},
						"updated_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Last updated time of the subscription.",
						},
					},
				},
			},
		},
	}
}

func dataSourceIBMEnSubscriptionsRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	enClient, err := meta.(ClientSession).EventNotificationsApiV1()
	if err != nil {
		return diag.FromErr(err)
	}

	options := &en.ListSubscriptionsOptions{}

	options.SetInstanceID(d.Get("instance_guid").(string))

	if _, ok := d.GetOk("search_key"); ok {
		options.SetSearch(d.Get("search_key").(string))
	}
	var subscriptionList *en.SubscriptionList

	finalList := []en.SubscriptionListItem{}

	var offset int64 = 0
	var limit int64 = 100

	options.SetLimit(limit)

	for {
		options.SetOffset(offset)

		result, response, err := enClient.ListSubscriptionsWithContext(context, options)

		subscriptionList = result

		if err != nil {
			return diag.FromErr(fmt.Errorf("ListSubscriptionsWithContext failed %s\n%s", err, response))
		}

		offset = offset + limit

		finalList = append(finalList, result.Subscriptions...)

		if offset > *result.TotalCount {
			break
		}
	}

	subscriptionList.Subscriptions = finalList

	d.SetId(fmt.Sprintf("subscriptions_%s", d.Get("instance_guid").(string)))

	if err = d.Set("total_count", intValue(subscriptionList.TotalCount)); err != nil {
		return diag.FromErr(fmt.Errorf("error setting total_count: %s", err))
	}

	if subscriptionList.Subscriptions != nil {
		err = d.Set("subscriptions", enFlattenSubscriptionList(subscriptionList.Subscriptions))
		if err != nil {
			return diag.FromErr(fmt.Errorf("error setting subscriptions %s", err))
		}
	}

	return nil
}

func enFlattenSubscriptionList(result []en.SubscriptionListItem) (subscriptions []map[string]interface{}) {
	subscriptions = []map[string]interface{}{}
	for _, subscriptionsItem := range result {
		subscriptions = append(subscriptions, enSubscriptionListToMap(subscriptionsItem))
	}

	return subscriptions
}

func enSubscriptionListToMap(subscriptionsItem en.SubscriptionListItem) (subscriptionsMap map[string]interface{}) {
	subscriptionsMap = map[string]interface{}{}

	if subscriptionsItem.ID != nil {
		subscriptionsMap["id"] = subscriptionsItem.ID
	}
	if subscriptionsItem.Name != nil {
		subscriptionsMap["name"] = subscriptionsItem.Name
	}
	if subscriptionsItem.Description != nil {
		subscriptionsMap["description"] = subscriptionsItem.Description
	}
	if subscriptionsItem.DestinationID != nil {
		subscriptionsMap["destination_id"] = subscriptionsItem.DestinationID
	}
	if subscriptionsItem.DestinationName != nil {
		subscriptionsMap["destination_name"] = subscriptionsItem.DestinationName
	}
	if subscriptionsItem.DestinationType != nil {
		subscriptionsMap["destination_type"] = subscriptionsItem.DestinationType
	}
	if subscriptionsItem.TopicID != nil {
		subscriptionsMap["topic_id"] = subscriptionsItem.TopicID
	}
	if subscriptionsItem.TopicName != nil {
		subscriptionsMap["topic_name"] = subscriptionsItem.TopicName
	}
	if subscriptionsItem.UpdatedAt != nil {
		subscriptionsMap["updated_at"] = subscriptionsItem.UpdatedAt.String()
	}

	return subscriptionsMap
}
