package ibm

import (
	"fmt"
	appid "github.com/IBM/appid-management-go-sdk/appidmanagementv4"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"testing"
)

func TestAccAppIDMFAChannel_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMAppIDMFAChannelDestroy,
		Steps: []resource.TestStep{
			{
				Config: setupIBMMFAChannelConfig(appIDTenantID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_appid_mfa_channel.mf", "tenant_id", appIDTenantID),
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
	appIDClient, err := testAccProvider.Meta().(ClientSession).AppIDAPI()

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
			return fmt.Errorf("Error checking if AppID MFA channel configuration was reset: %s", err)
		}

		for _, channel := range ch.Channels {
			if *channel.IsActive && *channel.Type != "email" {
				return fmt.Errorf("Error checking if AppID MFA channel configuration was reset")
			}
		}
	}

	return nil
}
