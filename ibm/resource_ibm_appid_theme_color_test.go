package ibm

import (
	"fmt"
	appid "github.com/IBM/appid-management-go-sdk/appidmanagementv4"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"testing"
)

func TestAccThemeColor_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMAppIDThemeColorDestroy,
		Steps: []resource.TestStep{
			{
				Config: setupAppIDThemeColorConfig(appIDTenantID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_appid_theme_color.color", "tenant_id", appIDTenantID),
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
	appIDClient, err := testAccProvider.Meta().(ClientSession).AppIDAPI()

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
			return fmt.Errorf("error checking if AppID theme color configuration (%s) has been reset", rs.Primary.ID)
		}
	}

	return nil
}
