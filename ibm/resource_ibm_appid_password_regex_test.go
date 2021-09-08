package ibm

import (
	"fmt"
	appid "github.com/IBM/appid-management-go-sdk/appidmanagementv4"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"testing"
)

func TestAccIBMAppIDPasswordRegex_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMAppIDPasswordRegexDestroy,
		Steps: []resource.TestStep{
			{
				Config: setupIBMAppIDPasswordRegexConfig(appIDTenantID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_appid_password_regex.rgx", "tenant_id", appIDTenantID),
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
	appIDClient, err := testAccProvider.Meta().(ClientSession).AppIDAPI()

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
			return fmt.Errorf("Error checking if AppID Password Regex was reset: %s", err)
		}

		// verify that configuration is reset to defaults
		if config == nil || (config.Base64EncodedRegex != nil && *config.Base64EncodedRegex != "") {
			return fmt.Errorf("Error checking if AppID Password Regex was reset")
		}
	}

	return nil
}
