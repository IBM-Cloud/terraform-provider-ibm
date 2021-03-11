package ibm

import (
	"fmt"

	"github.com/IBM/push-notifications-go-sdk/pushservicev1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceIBMPNApplicationChrome() *schema.Resource {
	return &schema.Resource{
		Read:   resourceApplicationChromeRead,
		Create: resourceApplicationChromeCreate,
		Update: resourceApplicationChromeUpdate,
		Delete: resourceApplicationChromeDelete,

		Schema: map[string]*schema.Schema{
			"application_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Unique ID of the application using the push service.",
			},
			"api_key": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "An API key that gives the push service an authorized access to Google services that is used for Chrome Web Push.",
			},
			"web_site_url": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "The URL of the WebSite / WebApp that should be permitted to subscribe to WebPush.",
			},
		},
	}
}

func resourceApplicationChromeCreate(d *schema.ResourceData, meta interface{}) error {
	pnClient, err := meta.(ClientSession).PushServiceV1()
	if err != nil {
		return err
	}

	serverKey := d.Get("api_key").(string)
	websiteURL := d.Get("web_site_url").(string)
	applicationGUID := d.Get("application_id").(string)

	_, response, err := pnClient.SaveChromeWebConf(&pushservicev1.SaveChromeWebConfOptions{
		ApplicationID: &applicationGUID,
		ApiKey:        &serverKey,
		WebSiteURL:    &websiteURL,
	})

	if err != nil {
		return fmt.Errorf("Error configuring chrome web platform: %s with responce code  %d", err, response.StatusCode)
	}

	return resourceApplicationChromeRead(d, meta)
}

func resourceApplicationChromeUpdate(d *schema.ResourceData, meta interface{}) error {

	if d.HasChanges("api_key", "web_site_url") {
		return resourceApplicationChromeCreate(d, meta)
	}
	return nil
}

func resourceApplicationChromeRead(d *schema.ResourceData, meta interface{}) error {
	pnClient, err := meta.(ClientSession).PushServiceV1()
	if err != nil {
		return err
	}

	applicationGUID := d.Get("application_id").(string)

	result, response, err := pnClient.GetChromeWebConf(&pushservicev1.GetChromeWebConfOptions{
		ApplicationID: &applicationGUID,
	})

	if err != nil {
		return fmt.Errorf("Error fetching chrome web platform configuration: %s with responce code  %d", err, response.StatusCode)
	}

	d.SetId(dataSourceIbmPnApplicationChromeID(d))

	if response.StatusCode == 200 {
		d.Set("api_key", *result.ApiKey)
		d.Set("web_site_url", *result.WebSiteURL)
	}
	return nil
}

func resourceApplicationChromeDelete(d *schema.ResourceData, meta interface{}) error {
	pnClient, err := meta.(ClientSession).PushServiceV1()
	if err != nil {
		return err
	}
	applicationGUID := d.Get("application_id").(string)

	response, e := pnClient.DeleteChromeWebConf(&pushservicev1.DeleteChromeWebConfOptions{
		ApplicationID: &applicationGUID,
	})

	if e != nil {
		return fmt.Errorf("Error deleting chrome web platform configuration: %s with responce code  %d", err, response.StatusCode)
	}

	d.SetId("")

	return nil

}
