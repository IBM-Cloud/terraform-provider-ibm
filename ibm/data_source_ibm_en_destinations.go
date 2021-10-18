// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	en "github.ibm.com/Notification-Hub/event-notifications-go-admin-sdk/eventnotificationsapiv1"
)

func dataSourceIBMEnDestinations() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceIBMEnDestinationsRead,

		Schema: map[string]*schema.Schema{
			"instance_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Unique identifier for IBM Cloud Event Notifications instance.",
			},
			"search_key": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: "Filter the destinations by name or type.",
			},
			"offset": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "Current offset.",
			},
			"limit": {
				Type:        schema.TypeInt,
				Optional:    true,
				Description: "limit to show destinations.",
			},
			"total_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Total number of destinations.",
			},
			"destinations": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: "List of destinations.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Destination ID.",
						},
						"name": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Destination name.",
						},
						"description": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Destination description.",
						},
						"type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Destination type Email/SMS/Webhook.",
						},
						"subscription_count": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: "Subscription count.",
						},
						"subscription_names": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: "Names of subscriptions.",
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
						"updated_at": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "Updated at.",
						},
					},
				},
			},
		},
	}
}

func dataSourceIBMEnDestinationsRead(d *schema.ResourceData, meta interface{}) error {
	enClient, err := meta.(ClientSession).EventNotificationsApiV1()
	if err != nil {
		return err
	}

	options := &en.ListDestinationsOptions{}

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

	result, response, err := enClient.ListDestinations(options)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("ListDestinations failed %s\n%s", err, response)
	}

	d.SetId(fmt.Sprintf("destinations/%s", *options.InstanceID))

	if err = d.Set("total_count", intValue(result.TotalCount)); err != nil {
		return fmt.Errorf("error setting total_count: %s", err)
	}

	if err = d.Set("offset", intValue(result.Offset)); err != nil {
		return fmt.Errorf("error setting offset: %s", err)
	}

	if err = d.Set("limit", intValue(result.Limit)); err != nil {
		return fmt.Errorf("error setting limit: %s", err)
	}

	if result.Destinations != nil {
		if err = d.Set("destinations", enFlattenDestinationsList(result.Destinations)); err != nil {
			return fmt.Errorf("error setting destinations %s", err)
		}
	}

	return nil
}

func enFlattenDestinationsList(result []en.DestinationLisItem) (destinations []map[string]interface{}) {
	for _, destinationsItem := range result {
		destinations = append(destinations, enDestinationListToMap(destinationsItem))
	}

	return destinations
}

func enDestinationListToMap(destinationItem en.DestinationLisItem) (destination map[string]interface{}) {
	destination = map[string]interface{}{}

	if destinationItem.ID != nil {
		destination["id"] = destinationItem.ID
	}
	if destinationItem.Name != nil {
		destination["name"] = destinationItem.Name
	}
	if destinationItem.Description != nil {
		destination["description"] = destinationItem.Description
	}
	if destinationItem.Type != nil {
		destination["type"] = destinationItem.Type
	}
	if destinationItem.SubscriptionCount != nil {
		destination["subscription_count"] = destinationItem.SubscriptionCount
	}
	if destinationItem.SubscriptionNames != nil {
		destination["subscription_names"] = destinationItem.SubscriptionNames
	}
	if destinationItem.UpdatedAt != nil {
		destination["updated_at"] = dateTimeToString(destinationItem.UpdatedAt)
	}

	return destination
}
