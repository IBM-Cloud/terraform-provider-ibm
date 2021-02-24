package ibm

import (
	"fmt"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/ibm-cloud-security/security-advisor-sdk-go/notificationsapiv1"
)

func resourceIBMNotificationChannel() *schema.Resource {
	return &schema.Resource{
		Create: resourceIBMNotificationChannelCreate,
		Read:   resourceIBMNotificationChannelRead,
		Update: resourceIBMNotificationChannelUpdate,
		Delete: resourceIBMNotificationChannelDelete,
		Exists: resourceIBMNotificationChannelExists,

		Schema: map[string]*schema.Schema{
			"channel_id": {
				Type:     schema.TypeString,
				Computed: true,
				ForceNew: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"type": {
				Type:     schema.TypeString,
				Required: true,
			},
			"endpoint": {
				Type:     schema.TypeString,
				Required: true,
			},
			"enabled": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"frequency": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"severity": {
				Type:     schema.TypeList,
				Optional: true,
				Elem:     &schema.Schema{Type: schema.TypeString},
			},
			"alert_source": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"provider_name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"finding_types": {
							Type:     schema.TypeList,
							Optional: true,
							Elem:     &schema.Schema{Type: schema.TypeString},
						},
					},
				},
			},
		},
	}
}

func resourceIBMNotificationChannelCreate(d *schema.ResourceData, meta interface{}) error {
	channelName := d.Get("name").(string)
	channelDescription := d.Get("description").(string)
	channelType := d.Get("type").(string)
	channelEndpoint := d.Get("endpoint").(string)
	channelEnabled := d.Get("enabled").(bool)
	channelSeverity := d.Get("severity").([]interface{})
	channelAlertSource := d.Get("alert_source").([]interface{})

	sess, err := meta.(ClientSession).NotificationsV1API()
	if err != nil {
		return err
	}

	userDetails, err := meta.(ClientSession).BluemixUserDetails()
	if err != nil {
		return err
	}
	accountID := userDetails.userAccount

	channelSeveritySet := make([]string, 0)
	for _, channelSeverity := range channelSeverity {
		channelSeveritySet = append(channelSeveritySet, channelSeverity.(string))
	}

	channelAlertSourceObjectList := make([]notificationsapiv1.NotificationChannelAlertSourceItem, 0)
	for _, channelAlertSource := range channelAlertSource {
		channelAlertSourceObject := notificationsapiv1.NotificationChannelAlertSourceItem{}
		providerName := channelAlertSource.(map[string]interface{})["provider_name"].(string)
		findingTypes := make([]string, 0)
		findingTypesInterfaceList := channelAlertSource.(map[string]interface{})["finding_types"].([]interface{})
		for _, findingType := range findingTypesInterfaceList {
			findingTypes = append(findingTypes, findingType.(string))
		}
		channelAlertSourceObject.ProviderName = &providerName
		channelAlertSourceObject.FindingTypes = findingTypes
		channelAlertSourceObjectList = append(channelAlertSourceObjectList, channelAlertSourceObject)
	}

	createChannelOptions := sess.NewCreateNotificationChannelOptions(accountID, channelName, channelType, channelEndpoint)
	if channelDescription != "" {
		createChannelOptions.SetDescription(channelDescription)
	}
	if len(channelSeveritySet) != 0 {
		createChannelOptions.SetSeverity(channelSeveritySet)
	}
	if len(channelAlertSourceObjectList) != 0 {
		createChannelOptions.SetAlertSource(channelAlertSourceObjectList)
	}
	if channelEnabled {
		createChannelOptions.SetEnabled(channelEnabled)
	}
	channel, _, err := sess.CreateNotificationChannel(createChannelOptions)
	if err != nil {
		return fmt.Errorf("error while creating notification channel: %v", err)
	}
	if channel.ChannelID != nil {
		d.SetId(*channel.ChannelID)
		return resourceIBMNotificationChannelRead(d, meta)
	}

	return nil
}

func resourceIBMNotificationChannelRead(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).NotificationsV1API()
	if err != nil {
		return err
	}

	userDetails, err := meta.(ClientSession).BluemixUserDetails()
	if err != nil {
		return err
	}
	accountID := userDetails.userAccount
	channelID := d.Id()

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

func resourceIBMNotificationChannelExists(d *schema.ResourceData, meta interface{}) (bool, error) {
	sess, err := meta.(ClientSession).NotificationsV1API()
	if err != nil {
		return false, err
	}

	userDetails, err := meta.(ClientSession).BluemixUserDetails()
	if err != nil {
		return false, err
	}
	accountID := userDetails.userAccount
	channelID := d.Id()

	getChannelOptions := sess.NewGetNotificationChannelOptions(accountID, channelID)
	_, resp, err := sess.GetNotificationChannel(getChannelOptions)

	if err != nil {
		return false, fmt.Errorf("error while reading notification channel: %v", err)
	}

	if resp.StatusCode == 404 {
		return false, nil
	}

	return true, nil
}

func resourceIBMNotificationChannelUpdate(d *schema.ResourceData, meta interface{}) error {
	channelName := d.Get("name").(string)
	channelDescription := d.Get("description").(string)
	channelType := d.Get("type").(string)
	channelEndpoint := d.Get("endpoint").(string)
	channelEnabled := d.Get("enabled").(bool)
	channelSeverity := d.Get("severity").([]interface{})
	channelAlertSource := d.Get("alert_source").([]interface{})

	sess, err := meta.(ClientSession).NotificationsV1API()
	if err != nil {
		return err
	}

	userDetails, err := meta.(ClientSession).BluemixUserDetails()
	if err != nil {
		return err
	}
	accountID := userDetails.userAccount
	channelID := d.Id()

	channelSeveritySet := make([]string, 0)
	for _, channelSeverity := range channelSeverity {
		channelSeveritySet = append(channelSeveritySet, channelSeverity.(string))
	}

	channelAlertSourceObjectList := make([]notificationsapiv1.NotificationChannelAlertSourceItem, 0)
	for _, channelAlertSource := range channelAlertSource {
		channelAlertSourceObject := notificationsapiv1.NotificationChannelAlertSourceItem{}
		providerName := channelAlertSource.(map[string]interface{})["provider_name"].(string)
		findingTypes := make([]string, 0)
		findingTypesInterfaceList := channelAlertSource.(map[string]interface{})["finding_types"].([]interface{})
		for _, findingType := range findingTypesInterfaceList {
			findingTypes = append(findingTypes, findingType.(string))
		}
		channelAlertSourceObject.ProviderName = &providerName
		channelAlertSourceObject.FindingTypes = findingTypes
		channelAlertSourceObjectList = append(channelAlertSourceObjectList, channelAlertSourceObject)
	}

	updateChannelOptions := sess.NewUpdateNotificationChannelOptions(accountID, channelID, channelName, channelType, channelEndpoint)
	if channelDescription != "" {
		updateChannelOptions.SetDescription(channelDescription)
	}
	if len(channelSeveritySet) != 0 {
		updateChannelOptions.SetSeverity(channelSeveritySet)
	}
	if len(channelAlertSourceObjectList) != 0 {
		updateChannelOptions.SetAlertSource(channelAlertSourceObjectList)
	}
	if channelEnabled {
		updateChannelOptions.SetEnabled(channelEnabled)
	}
	channel, _, err := sess.UpdateNotificationChannel(updateChannelOptions)
	if err != nil {
		return fmt.Errorf("error while updating notification channel: %v", err)
	}
	if channel.ChannelID != nil {
		d.Set("channel_id", *channel.ChannelID)
		return resourceIBMNotificationChannelRead(d, meta)
	}

	return nil
}

func resourceIBMNotificationChannelDelete(d *schema.ResourceData, meta interface{}) error {
	sess, err := meta.(ClientSession).NotificationsV1API()
	if err != nil {
		return err
	}

	userDetails, err := meta.(ClientSession).BluemixUserDetails()
	if err != nil {
		return err
	}
	accountID := userDetails.userAccount
	channelID := d.Id()

	deleteChannelOptions := sess.NewDeleteNotificationChannelOptions(accountID, channelID)
	_, _, err = sess.DeleteNotificationChannel(deleteChannelOptions)

	if err != nil {
		return fmt.Errorf("error occurred while deleting notification channel: %v", err)
	}

	d.SetId("")

	return nil
}
