package appid_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIBMAppIDApplicationScopesDataSource_basic(t *testing.T) {
	appName := fmt.Sprintf("tf_testacc_app_scopes_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMAppIDApplicationScopesDataSourceConfig(acc.AppIDTenantID, appName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_appid_application_scopes.scopes", "scopes.#", "3"),
					resource.TestCheckResourceAttr("data.ibm_appid_application_scopes.scopes", "scopes.0", "scope1"),
					resource.TestCheckResourceAttr("data.ibm_appid_application_scopes.scopes", "scopes.1", "scope2"),
					resource.TestCheckResourceAttr("data.ibm_appid_application_scopes.scopes", "scopes.2", "scope3"),
				),
			},
		},
	})
}

func testAccCheckIBMAppIDApplicationScopesDataSourceConfig(tenantID string, name string) string {
	return fmt.Sprintf(`
		resource "ibm_appid_application" "test_app" {
			tenant_id = "%s"
			name = "%s"  	
		}

		resource "ibm_appid_application_scopes" "scopes" {
		  tenant_id = ibm_appid_application.test_app.tenant_id
		  client_id = ibm_appid_application.test_app.client_id
		  scopes = ["scope1", "scope2", "scope3"]
		}

		data "ibm_appid_application_scopes" "scopes" {
			tenant_id = ibm_appid_application.test_app.tenant_id
			client_id = ibm_appid_application.test_app.client_id

			depends_on = [
				ibm_appid_application_scopes.scopes
			]
		}
	`, tenantID, name)
}
