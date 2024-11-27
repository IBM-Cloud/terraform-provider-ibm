package appid_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIBMAppIDApplicationDataSource_basic(t *testing.T) {
	appName := fmt.Sprintf("tf_testacc_app_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMAppIDApplicationDataSourceConfig(acc.AppIDTenantID, appName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_appid_application.test_app", "tenant_id", acc.AppIDTenantID),
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
