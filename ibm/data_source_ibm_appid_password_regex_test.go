package ibm

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccIBMAppIDPasswordRegexDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: setupIBMAppIDPasswordRegexDataSourceConfig(appIDTenantID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_appid_password_regex.rgx", "tenant_id", appIDTenantID),
					resource.TestCheckResourceAttr("data.ibm_appid_password_regex.rgx", "regex", "^(?:(?=.*\\d)(?=.*[a-z])(?=.*[A-Z]).*)$"),
					resource.TestCheckResourceAttr("data.ibm_appid_password_regex.rgx", "error_message", "test error"),
				),
			},
		},
	})
}

func setupIBMAppIDPasswordRegexDataSourceConfig(tenantID string) string {
	return fmt.Sprintf(`
		resource "ibm_appid_password_regex" "rgx" {
			tenant_id = "%s"
			regex = "^(?:(?=.*\\d)(?=.*[a-z])(?=.*[A-Z]).*)$"
			error_message = "test error"
		}

		data "ibm_appid_password_regex" "rgx" {
			tenant_id = ibm_appid_password_regex.rgx.tenant_id

			depends_on = [
				ibm_appid_password_regex.rgx
			]
		}
	`, tenantID)
}
