package ibm

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccThemeTextDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: setupAppIDThemeTextDataSourceConfig(appIDTenantID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_appid_theme_text.text", "tenant_id", appIDTenantID),
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
