package appid_test

import (
	"fmt"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccThemeColorDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: setupAppIDThemeColorDataSourceConfig(acc.AppIDTenantID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_appid_theme_color.color", "tenant_id", acc.AppIDTenantID),
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
