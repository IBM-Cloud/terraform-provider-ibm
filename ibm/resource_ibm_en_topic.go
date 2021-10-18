// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/IBM/go-sdk-core/v5/core"
	en "github.ibm.com/Notification-Hub/event-notifications-go-admin-sdk/eventnotificationsapiv1"
)

func resourceIBMEnTopic() *schema.Resource {
	return &schema.Resource{
		Create:   resourceIBMEnTopicCreate,
		Read:     resourceIBMEnTopicRead,
		Update:   resourceIBMEnTopicUpdate,
		Delete:   resourceIBMEnTopicDelete,
		Importer: &schema.ResourceImporter{},

		Schema: map[string]*schema.Schema{
			"instance_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: "Unique identifier for IBM Cloud Event Notifications instance.",
			},
			"name": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Name of the topic.",
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Default:     "",
				Description: "Description of the topic.",
			},
			"sources": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: "List of sources.",
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:        schema.TypeString,
							Required:    true,
							Description: "ID of the source.",
						},
						"rules": {
							Type:        schema.TypeList,
							Optional:    true,
							Description: "List of rules.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"enabled": {
										Type:        schema.TypeBool,
										Optional:    true,
										Default:     true,
										Description: "Whether the rule is enabled or not.",
									},
									"event_type_filter": {
										Type:        schema.TypeString,
										Required:    true,
										Description: "Event type filter.",
									},
									"notification_filter": {
										Type:        schema.TypeString,
										Optional:    true,
										Default:     "",
										Description: "Notification filter.",
									},
								},
							},
						},
					},
				},
			},
			"topic_id": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Topic ID.",
			},
			"updated_at": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Last time the topic was updated.",
			},
			"source_count": {
				Type:        schema.TypeInt,
				Computed:    true,
				Description: "Number of sources.",
			},
			"subscription_count": {
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
							Optional:    true,
							Description: "Description of the subscription.",
						},
						"destination_id": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: "ID of the destination.",
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

func resourceIBMEnTopicCreate(d *schema.ResourceData, meta interface{}) error {
	enClient, err := meta.(ClientSession).EventNotificationsApiV1()
	if err != nil {
		return err
	}

	options := &en.CreateTopicOptions{}

	options.SetInstanceID(d.Get("instance_id").(string))
	options.SetName(d.Get("name").(string))

	if _, ok := d.GetOk("description"); ok {
		options.SetDescription(d.Get("description").(string))
	}

	if _, ok := d.GetOk("sources"); ok {
		var sources []en.TopicUpdateSourcesItem
		for _, e := range d.Get("sources").([]interface{}) {
			value := e.(map[string]interface{})
			sourcesItem := enTopicUpdateSourcesItem(value)
			sources = append(sources, sourcesItem)
		}
		options.SetSources(sources)
	}

	result, response, err := enClient.CreateTopic(options)

	if err != nil {
		return fmt.Errorf("CreateTopic failed %s\n%s", err, response)
	}

	d.SetId(fmt.Sprintf("%s/%s", *options.InstanceID, *result.ID))

	return resourceIBMEnTopicRead(d, meta)
}

func resourceIBMEnTopicRead(d *schema.ResourceData, meta interface{}) error {
	enClient, err := meta.(ClientSession).EventNotificationsApiV1()
	if err != nil {
		return err
	}

	options := &en.GetTopicOptions{}

	parts, err := sepIdParts(d.Id(), "/")
	if err != nil {
		return err
	}

	options.SetInstanceID(parts[0])
	options.SetID(parts[1])
	options.SetHeaders(map[string]string{"include": "subscriptions"})

	result, response, err := enClient.GetTopic(options)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}

		return fmt.Errorf("GetTopic failed %s\n%s", err, response)
	}

	d.Set("instance_id", options.InstanceID)
	d.Set("topic_id", result.ID)

	if err = d.Set("name", result.Name); err != nil {
		return fmt.Errorf("error setting name: %s", err)
	}

	if err = d.Set("description", result.Description); err != nil {
		return fmt.Errorf("error setting description: %s", err)
	}

	if result.Sources != nil {
		sources := []map[string]interface{}{}
		for _, sourcesItem := range result.Sources {
			sourcesItemMap := enTopicUpdateSourcesItemToMap(sourcesItem)
			sources = append(sources, sourcesItemMap)
		}
		if err = d.Set("sources", sources); err != nil {
			return fmt.Errorf("error setting sources: %s", err)
		}
	}

	if err = d.Set("updated_at", result.UpdatedAt); err != nil {
		return fmt.Errorf("error setting updated_at: %s", err)
	}

	if err = d.Set("source_count", intValue(result.SourceCount)); err != nil {
		return fmt.Errorf("error setting source_count: %s", err)
	}

	if err = d.Set("subscription_count", intValue(result.SubscriptionCount)); err != nil {
		return fmt.Errorf("error setting subscription_count: %s", err)
	}

	subscriptions := []map[string]interface{}{}
	for _, subscriptionsItem := range result.Subscriptions {
		subscriptionsItemMap := enTopicSubscriptionToMap(subscriptionsItem)
		subscriptions = append(subscriptions, subscriptionsItemMap)
	}

	if err = d.Set("subscriptions", subscriptions); err != nil {
		return fmt.Errorf("error setting subscriptions: %s", err)
	}

	return nil
}

func resourceIBMEnTopicUpdate(d *schema.ResourceData, meta interface{}) error {
	enClient, err := meta.(ClientSession).EventNotificationsApiV1()
	if err != nil {
		return err
	}

	options := &en.ReplaceTopicOptions{}

	parts, err := sepIdParts(d.Id(), "/")
	if err != nil {
		return err
	}

	options.SetInstanceID(parts[0])
	options.SetID(parts[1])
	options.SetInstanceID(d.Get("instance_id").(string))

	options.SetName(d.Get("name").(string))

	if _, ok := d.GetOk("description"); ok {
		options.SetDescription(d.Get("description").(string))
	}

	if _, ok := d.GetOk("sources"); ok {
		var sources []en.TopicUpdateSourcesItem
		for _, e := range d.Get("sources").([]interface{}) {
			value := e.(map[string]interface{})
			sourcesItem := enTopicUpdateSourcesItem(value)
			sources = append(sources, sourcesItem)
		}
		options.SetSources(sources)
	}

	_, response, err := enClient.ReplaceTopic(options)
	if err != nil {
		return fmt.Errorf("ReplaceTopic failed %s\n%s", err, response)
	}

	return resourceIBMEnTopicRead(d, meta)
}

func resourceIBMEnTopicDelete(d *schema.ResourceData, meta interface{}) error {
	enClient, err := meta.(ClientSession).EventNotificationsApiV1()
	if err != nil {
		return err
	}

	options := &en.DeleteTopicOptions{}

	parts, err := sepIdParts(d.Id(), "/")
	if err != nil {
		return err
	}

	options.SetInstanceID(parts[0])
	options.SetID(parts[1])

	response, err := enClient.DeleteTopic(options)
	if err != nil {
		if response != nil && response.StatusCode == 404 {
			d.SetId("")
			return nil
		}
		return fmt.Errorf("DeleteTopic failed %s\n%s", err, response)
	}

	d.SetId("")

	return nil
}

func enTopicUpdateSourcesItem(topicUpdateSourcesItemMap map[string]interface{}) en.TopicUpdateSourcesItem {
	topicUpdateSourcesItem := en.TopicUpdateSourcesItem{}

	if topicUpdateSourcesItemMap["id"] != nil {
		topicUpdateSourcesItem.ID = core.StringPtr(topicUpdateSourcesItemMap["id"].(string))
	}
	if topicUpdateSourcesItemMap["rules"] != nil {
		rules := []en.Rules{}
		for _, rulesItem := range topicUpdateSourcesItemMap["rules"].([]interface{}) {
			rulesItemModel := resourceIBMEnTopicMapToRules(rulesItem.(map[string]interface{}))
			rules = append(rules, rulesItemModel)
		}
		topicUpdateSourcesItem.Rules = rules
	}

	return topicUpdateSourcesItem
}

func resourceIBMEnTopicMapToRules(rulesMap map[string]interface{}) en.Rules {
	rules := en.Rules{}

	if rulesMap["enabled"] != nil {
		rules.Enabled = core.BoolPtr(rulesMap["enabled"].(bool))
	}

	rules.EventTypeFilter = core.StringPtr(rulesMap["event_type_filter"].(string))

	if rulesMap["notification_filter"] != nil {
		rules.NotificationFilter = core.StringPtr(rulesMap["notification_filter"].(string))
	}

	return rules
}

func enTopicUpdateSourcesItemToMap(source en.TopicSourcesItem) map[string]interface{} {
	sourceMap := map[string]interface{}{}

	if source.ID != nil {
		sourceMap["id"] = source.ID
	}

	if source.Rules != nil {
		rules := []map[string]interface{}{}
		for _, rulesItem := range source.Rules {
			rulesItemMap := resourceIBMEnTopicRulesToMap(rulesItem)
			rules = append(rules, rulesItemMap)
		}
		sourceMap["rules"] = rules
	}

	return sourceMap
}

func resourceIBMEnTopicRulesToMap(rules en.RulesGet) map[string]interface{} {
	rulesMap := map[string]interface{}{}

	if rules.Enabled != nil {
		rulesMap["enabled"] = rules.Enabled
	}

	rulesMap["event_type_filter"] = rules.EventTypeFilter

	if rules.NotificationFilter != nil {
		rulesMap["notification_filter"] = rules.NotificationFilter
	}

	return rulesMap
}

func enTopicSubscriptionToMap(subscriptionListItem en.SubscriptionListItem) map[string]interface{} {
	subscriptionMap := map[string]interface{}{}

	subscriptionMap["id"] = subscriptionListItem.ID

	subscriptionMap["name"] = subscriptionListItem.Name

	subscriptionMap["destination_id"] = subscriptionListItem.DestinationID

	if subscriptionListItem.Description != nil {
		subscriptionMap["description"] = subscriptionListItem.Description
	}

	subscriptionMap["destination_type"] = subscriptionListItem.DestinationType

	subscriptionMap["topic_id"] = subscriptionListItem.TopicID

	subscriptionMap["updated_at"] = subscriptionListItem.UpdatedAt.String()

	return subscriptionMap
}
