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

func TestAccIBMAppIDPasswordRegex_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMAppIDPasswordRegexDestroy,
		Steps: []resource.TestStep{
			{
				Config: setupIBMAppIDPasswordRegexConfig(acc.AppIDTenantID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_appid_password_regex.rgx", "tenant_id", acc.AppIDTenantID),
					resource.TestCheckResourceAttr("ibm_appid_password_regex.rgx", "regex", "^(?:(?=.*\\d)(?=.*[a-z])(?=.*[A-Z]).*)$"),
					resource.TestCheckResourceAttr("ibm_appid_password_regex.rgx", "error_message", "test error"),
				),
			},
		},
	})
}

func setupIBMAppIDPasswordRegexConfig(tenantID string) string {
	return fmt.Sprintf(`
		resource "ibm_appid_password_regex" "rgx" {
			tenant_id = "%s"
			regex = "^(?:(?=.*\\d)(?=.*[a-z])(?=.*[A-Z]).*)$"
			error_message = "test error"
		}
	`, tenantID)
}

func testAccCheckIBMAppIDPasswordRegexDestroy(s *terraform.State) error {
	appIDClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).AppIDAPI()

	if err != nil {
		return err
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_appid_password_regex" {
			continue
		}

		tenantID := rs.Primary.ID

		config, _, err := appIDClient.GetCloudDirectoryPasswordRegex(&appid.GetCloudDirectoryPasswordRegexOptions{
			TenantID: &tenantID,
		})

		if err != nil {
			return fmt.Errorf("[ERROR] Error checking if AppID Password Regex was reset: %s", err)
		}

		// verify that configuration is reset to defaults
		if config == nil || (config.Base64EncodedRegex != nil && *config.Base64EncodedRegex != "") {
			return fmt.Errorf("[ERROR] Error checking if AppID Password Regex was reset")
		}
	}

	return nil
}
