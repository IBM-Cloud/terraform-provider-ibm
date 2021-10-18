// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	en "github.ibm.com/Notification-Hub/event-notifications-go-admin-sdk/eventnotificationsapiv1"
)

func dataSourceIBMEnSubscriptions() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIBMEnSubscriptionsRead,

		Schema: map[string]*schema.Schema{
			"instance_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Unique identifier for IBM Cloud Event Notifications instance.",
			},
			"offset": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Current offset.",
			},
			"limit": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "limit to show subscriptions.",
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

func dataSourceIBMEnSubscriptionsRead(d *schema.ResourceData, meta interface{}) error {
	enClient, err := meta.(ClientSession).EventNotificationsApiV1()
	if err != nil {
		return err
	}

	options := &en.ListSubscriptionsOptions{}

	options.SetInstanceID(d.Get("instance_id").(string))

	if _, ok := d.GetOk("limit"); ok {
		options.SetLimit(d.Get("limit").(int64))
	}
	if _, ok := d.GetOk("offset"); ok {
		options.SetOffset(d.Get("offset").(int64))
	}
	if _, ok := d.GetOk("search_key"); ok {
		options.SetSearch(d.Get("search_key").(string))
	}

	result, response, err := enClient.ListSubscriptions(options)
	if err != nil {
		return fmt.Errorf("ListSubscriptions failed %s\n%s", err, response)
	}

	d.SetId(fmt.Sprintf("subscriptions_%s", d.Get("instance_id").(string)))

	if err = d.Set("total_count", intValue(result.TotalCount)); err != nil {
		return fmt.Errorf("error setting total_count: %s", err)
	}
	if err = d.Set("offset", intValue(result.Offset)); err != nil {
		return fmt.Errorf("error setting offset: %s", err)
	}
	if err = d.Set("limit", intValue(result.Limit)); err != nil {
		return fmt.Errorf("error setting limit: %s", err)
	}

	if result.Subscriptions != nil {
		err = d.Set("subscriptions", enFlattenSubscriptionList(result.Subscriptions))
		if err != nil {
			return fmt.Errorf("error setting subscriptions %s", err)
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

	fmt.Printf("Pajabjkd %+v", subscriptionsItem)
	return subscriptionsMap
}
