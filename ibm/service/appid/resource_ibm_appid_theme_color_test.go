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

func TestAccThemeColor_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMAppIDThemeColorDestroy,
		Steps: []resource.TestStep{
			{
				Config: setupAppIDThemeColorConfig(acc.AppIDTenantID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_appid_theme_color.color", "tenant_id", acc.AppIDTenantID),
					resource.TestCheckResourceAttr("ibm_appid_theme_color.color", "header_color", "#000000"),
				),
			},
		},
	})
}

func setupAppIDThemeColorConfig(tenantID string) string {
	return fmt.Sprintf(`
		resource "ibm_appid_theme_color" "color" {
			tenant_id = "%s"
			header_color = "#000000"
		}
	`, tenantID)
}

func testAccCheckIBMAppIDThemeColorDestroy(s *terraform.State) error {
	appIDClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).AppIDAPI()

	if err != nil {
		return err
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_appid_theme_color" {
			continue
		}

		tenantID := rs.Primary.ID

		cfg, _, err := appIDClient.GetThemeColor(&appid.GetThemeColorOptions{
			TenantID: &tenantID,
		})

		// AppID default default #EEF2F5
		if err != nil || (cfg.HeaderColor != nil && *cfg.HeaderColor != "#EEF2F5") {
			return fmt.Errorf("[ERROR] Error checking if AppID theme color configuration (%s) has been reset", rs.Primary.ID)
		}
	}

	return nil
}
