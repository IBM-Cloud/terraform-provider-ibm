package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMAppIDApplicationDataSource_basic(t *testing.T) {
	appName := fmt.Sprintf("tf_testacc_app_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMAppIDApplicationDataSourceConfig(appIDTenantID, appName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_appid_application.test_app", "tenant_id", appIDTenantID),
					resource.TestCheckResourceAttr("data.ibm_appid_application.test_app", "name", appName),
					resource.TestCheckResourceAttr("data.ibm_appid_application.test_app", "type", "singlepageapp"),
					resource.TestCheckResourceAttrSet("data.ibm_appid_application.test_app", "client_id"),
				),
			},
		},
	})
}

func testAccCheckIBMAppIDApplicationDataSourceConfig(tenantID string, name string) string {
	return fmt.Sprintf(`
		resource "ibm_appid_application" "test_app" {
			tenant_id = "%s"
			name = "%s"  
			type = "singlepageapp"
		}

		data "ibm_appid_application" "test_app" {
			tenant_id = ibm_appid_application.test_app.tenant_id
			client_id = ibm_appid_application.test_app.client_id

			depends_on = [
				ibm_appid_application.test_app
			]
		}
	`, tenantID, name)
}
