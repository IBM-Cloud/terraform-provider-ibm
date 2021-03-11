package ibm

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/IBM/push-notifications-go-sdk/pushservicev1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func dataSourceIBMPNApplicationChrome() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceApplicationChromeRead,

		Schema: map[string]*schema.Schema{
			"application_id": {
				Type:        schema.TypeString,
				Required:    true,
				Description: "Unique ID of the application using the push service.",
			},
			"api_key": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "An API key that gives the push service an authorized access to Google services that is used for Chrome Web Push.",
			},
			"web_site_url": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: "The URL of the WebSite / WebApp that should be permitted to subscribe to WebPush.",
			},
		},
	}
}

func dataSourceApplicationChromeRead(context context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	pushServiceClient, err := meta.(ClientSession).PushServiceV1()
	if err != nil {
		return diag.FromErr(err)
	}

	getChromeWebConfOptions := &pushservicev1.GetChromeWebConfOptions{}

	getChromeWebConfOptions.SetApplicationID(d.Get("application_id").(string))

	chromeWebPushCredendialsModel, response, err := pushServiceClient.GetChromeWebConfWithContext(context, getChromeWebConfOptions)
	if err != nil {
		log.Printf("[DEBUG] GetChromeWebConfWithContext failed %s\n%d", err, response.StatusCode)
		return diag.FromErr(err)
	}

	d.SetId(dataSourceIbmPnApplicationChromeID(d))
	if err = d.Set("api_key", chromeWebPushCredendialsModel.ApiKey); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting api_key: %s", err))
	}
	if err = d.Set("web_site_url", chromeWebPushCredendialsModel.WebSiteURL); err != nil {
		return diag.FromErr(fmt.Errorf("Error setting web_site_url: %s", err))
	}

	return nil
}

// dataSourceIbmPnApplicationChromeID returns a reasonable ID for the list.
func dataSourceIbmPnApplicationChromeID(d *schema.ResourceData) string {
	return time.Now().UTC().String()
}
