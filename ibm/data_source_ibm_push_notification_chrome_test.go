package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMDataSourcePNApplicationChrome_Basic(t *testing.T) {
	name := fmt.Sprintf("terraform_PN_%d", acctest.RandIntRange(10, 100))
	apiKey := fmt.Sprint(acctest.RandString(45))            // dummy value                       //dummy value
	websiteURL := "http://webpushnotificaton.mybluemix.net" // dummy url
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMDataSourcePNApplicationChrome(name, apiKey, websiteURL),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_pn_application_chrome.chrome", "api_key"),
					resource.TestCheckResourceAttrSet("data.ibm_pn_application_chrome.chrome", "web_site_url"),
				),
			},
		},
	})
}

func testAccCheckIBMDataSourcePNApplicationChrome(name, apiKey, websiteURL string) string {
	return fmt.Sprintf(`
		resource "ibm_resource_instance" "push_notification"{
			name     = "%s"
			location = "us-south"
			service  = "imfpush"
			plan     = "lite"
		}
		resource "ibm_pn_application_chrome" "application_chrome" {
			api_key            = "%s"
			web_site_url       = "%s"
			application_id = ibm_resource_instance.push_notification.guid
		}
		data "ibm_pn_application_chrome" "chrome" {
			application_id = ibm_pn_application_chrome.application_chrome.application_id
		}`, name, apiKey, websiteURL)
}
