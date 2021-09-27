package ibm

import (
	"fmt"
	appid "github.com/IBM/appid-management-go-sdk/appidmanagementv4"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"strings"
	"testing"
)

func TestAccIBMAppIDLanguages_basic(t *testing.T) {
	languages := []string{"en", "es-ES", "fr-FR"}

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMAppIDLanguagesDestroy,
		Steps: []resource.TestStep{
			{
				Config: setupIBMAppIDLanguagesConfig(appIDTenantID, languages),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_appid_languages.lang", "tenant_id", appIDTenantID),
					resource.TestCheckResourceAttr("ibm_appid_languages.lang", "languages.#", "3"),
					resource.TestCheckResourceAttr("ibm_appid_languages.lang", "languages.0", "en"),
					resource.TestCheckResourceAttr("ibm_appid_languages.lang", "languages.1", "es-ES"),
					resource.TestCheckResourceAttr("ibm_appid_languages.lang", "languages.2", "fr-FR"),
				),
			},
		},
	})
}

func setupIBMAppIDLanguagesConfig(tenantID string, languages []string) string {
	langs := strings.Replace(fmt.Sprintf("%q", languages), " ", ", ", -1)

	return fmt.Sprintf(`
		resource "ibm_appid_languages" "lang" {
			tenant_id = "%s"
			languages = %s
		}
	`, tenantID, langs)
}

func testAccCheckIBMAppIDLanguagesDestroy(s *terraform.State) error {
	appIDClient, err := testAccProvider.Meta().(ClientSession).AppIDAPI()

	if err != nil {
		return err
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_appid_languages" {
			continue
		}

		tenantID := rs.Primary.ID

		config, _, err := appIDClient.GetLocalization(&appid.GetLocalizationOptions{
			TenantID: &tenantID,
		})

		if err != nil {
			return fmt.Errorf("Error checking if AppID languages were reset: %s", err)
		}

		// verify that configuration is reset to defaults
		if config == nil || len(config.Languages) != 1 || (len(config.Languages) == 1 && config.Languages[0] != "en") {
			return fmt.Errorf("Error checking if AppID languages were reset")
		}
	}

	return nil
}
