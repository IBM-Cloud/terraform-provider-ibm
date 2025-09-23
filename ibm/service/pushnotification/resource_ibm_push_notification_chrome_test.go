package pushnotification_test

import (
	"fmt"
	"strings"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/conns"

	"github.com/IBM/push-notifications-go-sdk/pushservicev1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIBMPNApplicationChrome_Basic(t *testing.T) {
	var conf pushservicev1.ChromeWebPushCredendialsModel
	name := fmt.Sprintf("terraform_PN_%d", acctest.RandIntRange(10, 100))
	serverKey := fmt.Sprint(acctest.RandString(45))    // dummy value
	newServerKey := fmt.Sprint(acctest.RandString(45)) // dummy value
	websiteURL := "http://xyz.mybluemix.net"           // dummy url
	newWebsiteURL := "http://abc.mybluemix.net"        // dummy url
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMPNApplicationChromeDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPNApplicationChrome(name, serverKey, websiteURL),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPNApplicationChromeExists("ibm_pn_application_chrome.application_chrome", conf),
					resource.TestCheckResourceAttrSet("ibm_pn_application_chrome.application_chrome", "id"),
				),
			},
			{
				Config: testAccCheckIBMPNApplicationChromeUpdate(name, newServerKey, newWebsiteURL),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMPNApplicationChromeExists("ibm_pn_application_chrome.application_chrome", conf),
					resource.TestCheckResourceAttr("ibm_pn_application_chrome.application_chrome", "server_key", newServerKey),
					resource.TestCheckResourceAttr("ibm_pn_application_chrome.application_chrome", "web_site_url", newWebsiteURL),
				),
			},
		},
	})
}

func testAccCheckIBMPNApplicationChrome(name, serverKey, websiteURL string) string {
	return fmt.Sprintf(`
	resource "ibm_resource_instance" "push_notification"{
		name     = "%s"
		location = "us-south"
		service  = "imfpush"
		plan     = "lite"
	}
	resource "ibm_pn_application_chrome" "application_chrome" {
		server_key   = "%s"
		web_site_url = "%s"
		guid         = ibm_resource_instance.push_notification.guid
	}`, name, serverKey, websiteURL)
}

func testAccCheckIBMPNApplicationChromeUpdate(name, newServerKey, newWebsiteURL string) string {
	return fmt.Sprintf(`
	resource "ibm_resource_instance" "push_notification"{
		name     = "%s"
		location = "us-south"
		service  = "imfpush"
		plan     = "lite"
	}
	resource "ibm_pn_application_chrome" "application_chrome" {
		server_key    = "%s"
		web_site_url  = "%s"
		guid 					= ibm_resource_instance.push_notification.guid
	}`, name, newServerKey, newWebsiteURL)
}

func testAccCheckIBMPNApplicationChromeExists(n string, obj pushservicev1.ChromeWebPushCredendialsModel) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		pushServiceClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).PushServiceV1()
		if err != nil {
			return err
		}

		getChromeWebConfOptions := &pushservicev1.GetChromeWebConfOptions{}

		guid := rs.Primary.ID

		getChromeWebConfOptions.SetApplicationID(guid)

		chromeWebConf, _, err := pushServiceClient.GetChromeWebConf(getChromeWebConfOptions)
		if err != nil {
			return err
		}

		obj = *chromeWebConf
		return nil
	}
}

func testAccCheckIBMPNApplicationChromeDestroy(s *terraform.State) error {
	pushServiceClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).PushServiceV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_pn_application_chrome" {
			continue
		}

		getChromeWebConfOptions := &pushservicev1.GetChromeWebConfOptions{}

		guid := rs.Primary.ID

		getChromeWebConfOptions.SetApplicationID(guid)

		// Try to find the config
		_, _, err := pushServiceClient.GetChromeWebConf(getChromeWebConfOptions)

		if err != nil && !strings.Contains(err.Error(), "not found") {
			return fmt.Errorf("[ERROR] Error checking for chrome web config (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
