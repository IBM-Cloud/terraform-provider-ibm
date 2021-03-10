package ibm

import (
	"github.com/IBM/push-notifications-go-sdk/pushservicev1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceIBMPNApplicationChrome() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceApplicationChromeRead,

		Schema: map[string]*schema.Schema{
			"service_instance_guid": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Instance guid of the push notifications instance",
			},
			"server_key": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "Server key that provides Push Notification service authorized access to Google services that is used for Chrome Web Push.",
			},
			"website_url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL of the website/web application that should be permitted to subscribe to Web Push.",
			},
		},
	}
}

func dataSourceApplicationChromeRead(d *schema.ResourceData, meta interface{}) error {
	pnClient, err := meta.(ClientSession).PushNotificationsV1API()
	if err != nil {
		return err
	}

	serviceInstanceGUID := d.Get("service_instance_guid").(string)

	result, response, err := pnClient.GetChromeWebConf(&pushservicev1.GetChromeWebConfOptions{
		ApplicationID: &serviceInstanceGUID,
	})

	if err != nil {
		return err
	}

	if response.StatusCode == 200 {
		d.SetId(serviceInstanceGUID)
		d.Set("server_key", *result.ApiKey)
		d.Set("website_url", *result.WebSiteURL)
	}

	return nil
}
