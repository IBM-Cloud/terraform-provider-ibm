package ibm

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
)

func dataSourceIBMNotificationChannel() *schema.Resource {

	return &schema.Resource{
		Read: dataSourceIBMNotificationChannelRead,
		Schema: map[string]*schema.Schema{
			"channel_id": {
				Type:     schema.TypeString,
				Required: true,
			},
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
	}
}

func dataSourceIBMNotificationChannelRead(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).NotificationsV1API()
	if err != nil {
		return err
	}

	userDetails, err := meta.(ClientSession).BluemixUserDetails()
	if err != nil {
		return err
	}
	accountID := userDetails.userAccount
	channelID := d.Get("channel_id").(string)

	getChannelOptions := sess.NewGetNotificationChannelOptions(accountID, channelID)
	channel, _, err := sess.GetNotificationChannel(getChannelOptions)

	if err != nil {
		return fmt.Errorf("error occurred while reading notification channel: %v", err)
	}

	if channel.Channel == nil {
		return fmt.Errorf("no such notification channel found: %v", channelID)
	}

	notificationChannelAlertSourceList := make([]map[string]interface{}, 0)
	for _, alertSource := range channel.Channel.AlertSource {
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

	if channel.Channel.Severity.Critical != nil && *channel.Channel.Severity.Critical {
		notificationChannelSeverityObject["critical"] = true
	}
	if channel.Channel.Severity.High != nil && *channel.Channel.Severity.High {
		notificationChannelSeverityObject["high"] = true
	}
	if channel.Channel.Severity.Medium != nil && *channel.Channel.Severity.Medium {
		notificationChannelSeverityObject["medium"] = true
	}
	if channel.Channel.Severity.Low != nil && *channel.Channel.Severity.Low {
		notificationChannelSeverityObject["low"] = true
	}
	notificationChannelSeverityObjectList[0] = notificationChannelSeverityObject

	d.Set("name", channel.Channel.Name)
	d.Set("description", channel.Channel.Description)
	d.Set("type", channel.Channel.Type)
	d.Set("endpoint", channel.Channel.Endpoint)
	d.Set("enabled", channel.Channel.Enabled)
	d.Set("frequency", channel.Channel.Frequency)
	d.Set("severity", notificationChannelSeverityObjectList)
	d.Set("alert_source", notificationChannelAlertSourceList)
	d.SetId(*channel.Channel.ChannelID)

	return nil
}
