package appid_test

import (
	"fmt"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMAppIDActionURLDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: setupAppIDActionURLDataSourceConfig(acc.AppIDTenantID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_appid_action_url.url", "tenant_id", acc.AppIDTenantID),
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
