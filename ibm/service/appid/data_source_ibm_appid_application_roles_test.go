package appid_test

import (
	"fmt"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMAppIDApplicationRolesDataSource_basic(t *testing.T) {
	appName := fmt.Sprintf("tf_testacc_app_roles_%d", acctest.RandIntRange(10, 100))
	roleName := fmt.Sprintf("tf_testacc_app_roles_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMAppIDApplicationRolesDataSourceConfig(acc.AppIDTenantID, appName, roleName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_appid_application_roles.roles", "roles.#", "1"),
					resource.TestCheckResourceAttrPair("ibm_appid_role.role", "role_id", "data.ibm_appid_application_roles.roles", "roles.0.id"),
					resource.TestCheckResourceAttr("data.ibm_appid_application_roles.roles", "roles.0.name", roleName),
				),
			},
		},
	})
}

func testAccCheckIBMAppIDApplicationRolesDataSourceConfig(tenantID string, appName string, roleName string) string {
	return fmt.Sprintf(`
		resource "ibm_appid_application" "test_app" {
			tenant_id = "%s"
			name = "%s"  	
		}

		resource "ibm_appid_role" "role" {
			tenant_id = ibm_appid_application.test_app.tenant_id
			name = "%s"
		}

		resource "ibm_appid_application_roles" "roles" {
			tenant_id = ibm_appid_application.test_app.tenant_id
			client_id = ibm_appid_application.test_app.client_id
			roles = [ibm_appid_role.role.role_id]        
		}

		data "ibm_appid_application_roles" "roles" {
			tenant_id = ibm_appid_application.test_app.tenant_id
			client_id = ibm_appid_application.test_app.client_id
	
			depends_on = [
				ibm_appid_application_roles.roles
			]
		}
	`, tenantID, appName, roleName)
}
