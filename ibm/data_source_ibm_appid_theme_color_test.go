package ibm

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccThemeColorDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: setupAppIDThemeColorDataSourceConfig(appIDTenantID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_appid_theme_color.color", "tenant_id", appIDTenantID),
					resource.TestCheckResourceAttr("data.ibm_appid_theme_color.color", "header_color", "#000000"),
				),
			},
		},
	})
}

func setupAppIDThemeColorDataSourceConfig(tenantID string) string {
	return fmt.Sprintf(`
		resource "ibm_appid_theme_color" "color" {
			tenant_id = "%s"
			header_color = "#000000"
		}

		data "ibm_appid_theme_color" "color" {
			tenant_id = ibm_appid_theme_color.color.tenant_id

			depends_on = [
				ibm_appid_theme_color.color
			]
		}
	`, tenantID)
}
