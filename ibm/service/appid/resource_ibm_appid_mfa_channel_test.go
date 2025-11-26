package appid_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"

	appid "github.com/IBM/appid-management-go-sdk/appidmanagementv4"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccAppIDMFAChannel_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMAppIDMFAChannelDestroy,
		Steps: []resource.TestStep{
			{
				Config: setupIBMMFAChannelConfig(acc.AppIDTenantID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_appid_mfa_channel.mf", "tenant_id", acc.AppIDTenantID),
					resource.TestCheckResourceAttr("ibm_appid_mfa_channel.mf", "active", "sms"),
					resource.TestCheckResourceAttr("ibm_appid_mfa_channel.mf", "sms_config.#", "1"),
					resource.TestCheckResourceAttr("ibm_appid_mfa_channel.mf", "sms_config.0.key", "api_key"),
					resource.TestCheckResourceAttr("ibm_appid_mfa_channel.mf", "sms_config.0.secret", "api_secret"),
					resource.TestCheckResourceAttr("ibm_appid_mfa_channel.mf", "sms_config.0.from", "+12223334444"),
				),
			},
		},
	})
}

func setupIBMMFAChannelConfig(tenantID string) string {
	return fmt.Sprintf(`
		resource "ibm_appid_mfa_channel" "mf" {
			tenant_id = "%s"
			active = "sms"

			sms_config {
			  key = "api_key"
			  secret = "api_secret"
			  from = "+12223334444"
		    }
		}
	`, tenantID)
}

func testAccCheckIBMAppIDMFAChannelDestroy(s *terraform.State) error {
	appIDClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).AppIDAPI()

	if err != nil {
		return err
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_appid_mfa_channel" {
			continue
		}

		tenantID := rs.Primary.ID

		ch, _, err := appIDClient.ListChannels(&appid.ListChannelsOptions{
			TenantID: &tenantID,
		})

		if err != nil {
			return fmt.Errorf("[ERROR] Error checking if AppID MFA channel configuration was reset: %s", err)
		}

		for _, channel := range ch.Channels {
			if *channel.IsActive && *channel.Type != "email" {
				return fmt.Errorf("[ERROR] Error checking if AppID MFA channel configuration was reset")
			}
		}
	}

	return nil
}
