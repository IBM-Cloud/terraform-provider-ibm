package appid_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIBMAppIDRedirectURLsDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMAppIDRedirectURLsDataSourceConfig(acc.AppIDTenantID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_appid_redirect_urls.urls", "urls.#", "3"),
					resource.TestCheckResourceAttr("data.ibm_appid_redirect_urls.urls", "urls.0", "https://test-url-1.com"),
					resource.TestCheckResourceAttr("data.ibm_appid_redirect_urls.urls", "urls.1", "https://test-url-2.com"),
					resource.TestCheckResourceAttr("data.ibm_appid_redirect_urls.urls", "urls.2", "https://test-url-3.com"),
				),
			},
		},
	})
}

func testAccCheckIBMAppIDRedirectURLsDataSourceConfig(tenantID string) string {
	return fmt.Sprintf(`
		resource "ibm_appid_redirect_urls" "urls" {
			tenant_id = "%s"
			urls = [
				"https://test-url-1.com",
				"https://test-url-2.com",
				"https://test-url-3.com",
			]
		}

		data "ibm_appid_redirect_urls" "urls" {
			tenant_id = ibm_appid_redirect_urls.urls.tenant_id

			depends_on = [
				ibm_appid_redirect_urls.urls
			]
		}
		
	`, tenantID)
}
