package ibm

import (
	"fmt"
	appid "github.com/IBM/appid-management-go-sdk/appidmanagementv4"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"testing"
)

func TestAccIBMAppIDRedirectURLs_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMAppIDRedirectURLsDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMAppIDRedirectURLsConfig(appIDTenantID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_appid_redirect_urls.urls", "urls.#", "3"),
					resource.TestCheckResourceAttr("ibm_appid_redirect_urls.urls", "urls.0", "https://test-url-1.com"),
					resource.TestCheckResourceAttr("ibm_appid_redirect_urls.urls", "urls.1", "https://test-url-2.com"),
					resource.TestCheckResourceAttr("ibm_appid_redirect_urls.urls", "urls.2", "https://test-url-3.com"),
				),
			},
		},
	})
}

func testAccCheckIBMAppIDRedirectURLsConfig(tenantID string) string {
	return fmt.Sprintf(`
		resource "ibm_appid_redirect_urls" "urls" {
			tenant_id = "%s"
			urls = [
				"https://test-url-1.com",
				"https://test-url-2.com",
				"https://test-url-3.com",
			]
		}
	`, tenantID)
}

func testAccCheckIBMAppIDRedirectURLsDestroy(s *terraform.State) error {
	appIDClient, err := testAccProvider.Meta().(ClientSession).AppIDAPI()

	if err != nil {
		return err
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_appid_redirect_urls" {
			continue
		}

		tenantID := rs.Primary.ID

		urls, _, err := appIDClient.GetRedirectUris(&appid.GetRedirectUrisOptions{
			TenantID: &tenantID,
		})

		if err != nil || len(urls.RedirectUris) != 0 {
			return fmt.Errorf("error checking if AppID redirect URLs resource (%s) has been destroyed", rs.Primary.ID)
		}
	}

	return nil
}
