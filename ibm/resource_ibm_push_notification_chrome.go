package ibm

import (
	"log"

	"github.com/IBM/push-notifications-go-sdk/pushservicev1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIBMPNApplicationChrome() *schema.Resource {
	return &schema.Resource{
		Create: resourceApplicationChromeCreate,
		Update: resourceApplicationChromeUpdate,
		Delete: resourceApplicationChromeDelete,
		Read:   resourceApplicationChromeRead,

		Schema: map[string]*schema.Schema{
			"service_instance_guid": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Instance guid of the push notifications instance",
			},
			"server_key": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Server key that provides Push Notification service authorized access to Google services that is used for Chrome Web Push.",
			},
			"website_url": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The URL of the website/web application that should be permitted to subscribe to Web Push.",
			},
		},
	}
}

func resourceApplicationChromeCreate(d *schema.ResourceData, meta interface{}) error {
	pnClient, err := meta.(ClientSession).PushNotificationsV1API()
	if err != nil {
		return err
	}

	serverKey := d.Get("server_key").(string)
	websiteURL := d.Get("website_url").(string)
	serviceInstanceGUID := d.Get("service_instance_guid").(string)

	_, _, e := pnClient.SaveChromeWebConf(&pushservicev1.SaveChromeWebConfOptions{
		ApplicationID: &serviceInstanceGUID,
		ApiKey:        &serverKey,
		WebSiteURL:    &websiteURL,
	})

	if e != nil {
		log.Fatal(e)
		return e
	}

	return resourceApplicationChromeRead(d, meta)
}

func resourceApplicationChromeUpdate(d *schema.ResourceData, meta interface{}) error {

	if d.HasChanges("server_key", "website_url") {
		return resourceApplicationChromeCreate(d, meta)
	}
	return nil
}

func resourceApplicationChromeRead(d *schema.ResourceData, meta interface{}) error {
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

func resourceApplicationChromeDelete(d *schema.ResourceData, meta interface{}) error {
	pnClient, err := meta.(ClientSession).PushNotificationsV1API()
	if err != nil {
		return err
	}
	serviceInstanceGUID := d.Get("service_instance_guid").(string)

	_, e := pnClient.DeleteChromeWebConf(&pushservicev1.DeleteChromeWebConfOptions{
		ApplicationID: &serviceInstanceGUID,
	})

	if e != nil {
		return e
	}

	d.SetId("")

	return nil

}
