package pushnotification_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIBMPNApplicationChromeDataSource_Basic(t *testing.T) {
	name := fmt.Sprintf("terraform_PN_%d", acctest.RandIntRange(10, 100))
	serverKey := fmt.Sprint(acctest.RandString(45))         // dummy value                       //dummy value
	websiteURL := "http://webpushnotificaton.mybluemix.net" // dummy url
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPNApplicationChromeDataSourceConfig(name, serverKey, websiteURL),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_pn_application_chrome.chrome", "server_key"),
					resource.TestCheckResourceAttrSet("data.ibm_pn_application_chrome.chrome", "web_site_url"),
				),
			},
		},
	})
}

func testAccCheckIBMPNApplicationChromeDataSourceConfig(name, serverKey, websiteURL string) string {
	return fmt.Sprintf(`
		resource "ibm_resource_instance" "push_notification"{
			name     = "%s"
			location = "us-south"
			service  = "imfpush"
			plan     = "lite"
		}
		resource "ibm_pn_application_chrome" "application_chrome" {
			server_key          = "%s"
			web_site_url       = "%s"
			guid = ibm_resource_instance.push_notification.guid
		}
		data "ibm_pn_application_chrome" "chrome" {
			guid = ibm_pn_application_chrome.application_chrome.guid
		}`, name, serverKey, websiteURL)
}
