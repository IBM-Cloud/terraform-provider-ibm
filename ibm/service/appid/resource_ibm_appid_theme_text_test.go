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

func TestAccThemeText_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMAppIDThemeTextDestroy,
		Steps: []resource.TestStep{
			{
				Config: setupAppIDThemeTextConfig(acc.AppIDTenantID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_appid_theme_text.text", "tenant_id", acc.AppIDTenantID),
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
	appIDClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).AppIDAPI()

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
			return fmt.Errorf("[ERROR] Error checking if AppID theme text configuration (%s) has been reset", rs.Primary.ID)
		}
	}

	return nil
}
