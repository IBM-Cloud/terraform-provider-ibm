package ibm

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceIBMNotificationChannels() *schema.Resource {

	return &schema.Resource{
		Read: dataSourceIBMNotificationChannelsRead,
		Schema: map[string]*schema.Schema{
			"notification_channels": {
				Type:        schema.TypeList,
				Description: "Collection of SA Notification Channels",
				Computed:    true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"description": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"type": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"endpoint": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"enabled": {
							Type:     schema.TypeBool,
							Computed: true,
						},
						"frequency": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"severity": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"critical": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"high": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"medium": {
										Type:     schema.TypeBool,
										Computed: true,
									},
									"low": {
										Type:     schema.TypeBool,
										Computed: true,
									},
								},
							},
						},
						"alert_source": {
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"provider_name": {
										Type:     schema.TypeString,
										Computed: true,
									},
									"finding_types": {
										Type:     schema.TypeList,
										Computed: true,
										Elem:     &schema.Schema{Type: schema.TypeString},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

func dataSourceIBMNotificationChannelsRead(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).NotificationsV1API()
	if err != nil {
		return err
	}

	userDetails, err := meta.(ClientSession).BluemixUserDetails()
	if err != nil {
		return err
	}
	accountID := userDetails.userAccount

	listChannelsOptions := sess.NewListAllChannelsOptions(accountID)
	channels, _, err := sess.ListAllChannels(listChannelsOptions)

	if err != nil {
		return fmt.Errorf("error occurred while listing notification channels: %v", err)
	}

	notificationChannelsList := make([]map[string]interface{}, 0)
	for _, channel := range channels.Channels {

		notificationChannelAlertSourceList := make([]map[string]interface{}, 0)
		for _, alertSource := range channel.AlertSource {
			notificationChannelAlertSourceObject := map[string]interface{}{}
			notificationChannelAlertSourceObject["provider_name"] = alertSource.ProviderName
			notificationChannelAlertSourceObject["finding_types"] = alertSource.FindingTypes
			notificationChannelAlertSourceList = append(notificationChannelAlertSourceList, notificationChannelAlertSourceObject)
		}

		notificationChannelSeverityObjectList := make([]map[string]interface{}, 1)
		notificationChannelSeverityObject := map[string]interface{}{}
		notificationChannelSeverityObject["critical"] = false
		notificationChannelSeverityObject["high"] = false
		notificationChannelSeverityObject["medium"] = false
		notificationChannelSeverityObject["low"] = false
		if channel.Severity.Critical != nil && *channel.Severity.Critical {
			notificationChannelSeverityObject["critical"] = true
		}
		if channel.Severity.High != nil && *channel.Severity.High {
			notificationChannelSeverityObject["high"] = true
		}
		if channel.Severity.Medium != nil && *channel.Severity.Medium {
			notificationChannelSeverityObject["medium"] = true
		}
		if channel.Severity.Low != nil && *channel.Severity.Low {
			notificationChannelSeverityObject["low"] = true
		}
		notificationChannelSeverityObjectList[0] = notificationChannelSeverityObject

		notificationChannelObject := map[string]interface{}{}
		notificationChannelObject["id"] = channel.ChannelID
		notificationChannelObject["name"] = channel.Name
		notificationChannelObject["description"] = channel.Description
		notificationChannelObject["type"] = channel.Type
		notificationChannelObject["endpoint"] = channel.Endpoint
		notificationChannelObject["enabled"] = channel.Enabled
		notificationChannelObject["frequency"] = channel.Frequency
		notificationChannelObject["severity"] = notificationChannelSeverityObjectList
		notificationChannelObject["alert_source"] = notificationChannelAlertSourceList

		notificationChannelsList = append(notificationChannelsList, notificationChannelObject)
	}

	d.Set("notification_channels", notificationChannelsList)
	d.SetId(accountID)

	return nil
}
