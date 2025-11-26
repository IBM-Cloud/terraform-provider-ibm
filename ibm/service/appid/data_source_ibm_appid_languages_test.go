package appid_test

import (
	"fmt"
	"strings"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIBMAppIDLanguagesDataSource_basic(t *testing.T) {
	languages := []string{"en", "es-ES", "fr-FR"}

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: setupIBMAppIDLanguagesDataSourceConfig(acc.AppIDTenantID, languages),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_appid_languages.lang", "tenant_id", acc.AppIDTenantID),
					resource.TestCheckResourceAttr("data.ibm_appid_languages.lang", "languages.#", "3"),
					resource.TestCheckResourceAttr("data.ibm_appid_languages.lang", "languages.0", "en"),
					resource.TestCheckResourceAttr("data.ibm_appid_languages.lang", "languages.1", "es-ES"),
					resource.TestCheckResourceAttr("data.ibm_appid_languages.lang", "languages.2", "fr-FR"),
				),
			},
		},
	})
}

func setupIBMAppIDLanguagesDataSourceConfig(tenantID string, languages []string) string {
	langs := strings.Replace(fmt.Sprintf("%q", languages), " ", ", ", -1)

	return fmt.Sprintf(`
		resource "ibm_appid_languages" "lang" {
			tenant_id = "%s"
			languages = %s
		}
		data "ibm_appid_languages" "lang" {
			tenant_id = ibm_appid_languages.lang.tenant_id

			depends_on = [
				ibm_appid_languages.lang
			]
		}
	`, tenantID, langs)
}
