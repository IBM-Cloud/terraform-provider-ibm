package ibm

import (
	"fmt"
	appid "github.com/IBM/appid-management-go-sdk/appidmanagementv4"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"testing"
)

func TestAccThemeText_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMAppIDThemeTextDestroy,
		Steps: []resource.TestStep{
			{
				Config: setupAppIDThemeTextConfig(appIDTenantID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_appid_theme_text.text", "tenant_id", appIDTenantID),
					resource.TestCheckResourceAttr("ibm_appid_theme_text.text", "tab_title", "resource test title"),
					resource.TestCheckResourceAttr("ibm_appid_theme_text.text", "footnote", "resource test footnote"),
				),
			},
		},
	})
}

func setupAppIDThemeTextConfig(tenantID string) string {
	return fmt.Sprintf(`
		resource "ibm_appid_theme_text" "text" {
			tenant_id = "%s"
			tab_title = "resource test title"
			footnote = "resource test footnote"
		}
	`, tenantID)
}

func testAccCheckIBMAppIDThemeTextDestroy(s *terraform.State) error {
	appIDClient, err := testAccProvider.Meta().(ClientSession).AppIDAPI()

	if err != nil {
		return err
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_appid_theme_text" {
			continue
		}

		tenantID := rs.Primary.ID

		cfg, _, err := appIDClient.GetThemeText(&appid.GetThemeTextOptions{
			TenantID: &tenantID,
		})

		if err != nil || (cfg.TabTitle != nil && *cfg.TabTitle != "Login") || (cfg.Footnote != nil && *cfg.Footnote != "Powered by App ID") {
			return fmt.Errorf("error checking if AppID theme text configuration (%s) has been reset", rs.Primary.ID)
		}
	}

	return nil
}
