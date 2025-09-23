package appid_test

import (
	"fmt"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccThemeTextDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: setupAppIDThemeTextDataSourceConfig(acc.AppIDTenantID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_appid_theme_text.text", "tenant_id", acc.AppIDTenantID),
					resource.TestCheckResourceAttr("data.ibm_appid_theme_text.text", "tab_title", "test title"),
					resource.TestCheckResourceAttr("data.ibm_appid_theme_text.text", "footnote", "test footnote"),
				),
			},
		},
	})
}

func setupAppIDThemeTextDataSourceConfig(tenantID string) string {
	return fmt.Sprintf(`
		resource "ibm_appid_theme_text" "text" {
			tenant_id = "%s"
			tab_title = "test title"
			footnote = "test footnote"
		}
		data "ibm_appid_theme_text" "text" {
			tenant_id = ibm_appid_theme_text.text.tenant_id
			depends_on = [
				ibm_appid_theme_text.text
			]
		}
	`, tenantID)
}
