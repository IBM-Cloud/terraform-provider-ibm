package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMResourcePNApplicationChrome_Basic(t *testing.T) {
	name := fmt.Sprintf("terraform_PN_%d", acctest.RandIntRange(10, 100))
	apiKey := fmt.Sprint(acctest.RandString(45))                  // dummy value
	websiteURL := "http://webpushnotificatons.mybluemix.net"      // dummy url
	newWebsiteURL := "http://chromepushnotificaton.mybluemix.net" // dummy url
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMResourcePNApplicationChromeConfig(name, apiKey, websiteURL),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("ibm_pn_application_chrome.application_chrome", "id"),
				),
			},
			{
				Config: testAccCheckIBMResourcePNApplicationChromeUpdate(name, apiKey, newWebsiteURL),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_pn_application_chrome.application_chrome", "web_site_url", newWebsiteURL),
				),
			},
		},
	})
}

func testAccCheckIBMResourcePNApplicationChromeConfig(name, apiKey, websiteURL string) string {
	return fmt.Sprintf(`
	resource "ibm_resource_instance" "push_notification"{
		name     = "%s"
		location = "us-south"
		service  = "imfpush"
		plan     = "lite"
	}
	resource "ibm_pn_application_chrome" "application_chrome" {
		api_key            		= "%s"
		web_site_url           = "%s"
		application_id = ibm_resource_instance.push_notification.guid
	}`, name, apiKey, websiteURL)
}

func testAccCheckIBMResourcePNApplicationChromeUpdate(name, apiKey, newWebsiteURL string) string {
	return fmt.Sprintf(`
	resource "ibm_resource_instance" "push_notification"{
		name     = "%s"
		location = "us-south"
		service  = "imfpush"
		plan     = "lite"
	}
		resource "ibm_pn_application_chrome" "application_chrome" {
			api_key            		= "%s"
			web_site_url           = "%s"
			application_id = ibm_resource_instance.push_notification.guid
		}`, name, apiKey, newWebsiteURL)
}
