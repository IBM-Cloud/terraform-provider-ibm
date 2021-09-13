package ibm

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccIBMAppIDActionURLDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: setupAppIDActionURLDataSourceConfig(appIDTenantID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_appid_action_url.url", "tenant_id", appIDTenantID),
					resource.TestCheckResourceAttr("data.ibm_appid_action_url.url", "action", "on_user_verified"),
					resource.TestCheckResourceAttr("data.ibm_appid_action_url.url", "url", "https://www.example.com/?user=verified"),
				),
			},
		},
	})
}

func setupAppIDActionURLDataSourceConfig(tenantID string) string {
	return fmt.Sprintf(`
		resource "ibm_appid_action_url" "url" {
			tenant_id = "%s"
			action = "on_user_verified"
			url = "https://www.example.com/?user=verified"
		}

		data "ibm_appid_action_url" "url" {
			tenant_id = ibm_appid_action_url.url.tenant_id
			action = "on_user_verified"

			depends_on = [
				ibm_appid_action_url.url
			]
		}
	`, tenantID)
}
