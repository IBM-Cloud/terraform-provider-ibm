package appid_test

import (
	"fmt"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMAppIDRolesDataSource_basic(t *testing.T) {
	roleName1 := fmt.Sprintf("tf_testacc_role_1_%d", acctest.RandIntRange(10, 100))
	roleName2 := fmt.Sprintf("tf_testacc_role_2_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: setupAppIDRolesConfig(acc.AppIDTenantID, roleName1, roleName2),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_appid_roles.roles", "roles.#", "2"),
					resource.TestCheckResourceAttrSet("data.ibm_appid_roles.roles", "roles.0.role_id"),
					resource.TestCheckResourceAttr("data.ibm_appid_roles.roles", "roles.0.name", roleName1),
					resource.TestCheckResourceAttr("data.ibm_appid_roles.roles", "roles.0.description", "test role 1"),
					resource.TestCheckResourceAttrSet("data.ibm_appid_roles.roles", "roles.1.role_id"),
					resource.TestCheckResourceAttr("data.ibm_appid_roles.roles", "roles.1.name", roleName2),
					resource.TestCheckResourceAttr("data.ibm_appid_roles.roles", "roles.1.description", "test role 2"),
				),
			},
		},
	})
}

// Test assumes there are no pre-existing roles
func setupAppIDRolesConfig(tenantID string, roleName1 string, roleName2 string) string {
	return fmt.Sprintf(`	
		resource "ibm_appid_role" "role1" {
			tenant_id = "%s"
			name = "%s"
			description = "test role 1"	
		}

		resource "ibm_appid_role" "role2" {
			tenant_id = ibm_appid_role.role1.tenant_id
			name = "%s"
			description = "test role 2"	
		}

		data "ibm_appid_roles" "roles" {
			tenant_id = ibm_appid_role.role1.tenant_id
			
			depends_on = [
				ibm_appid_role.role1,
				ibm_appid_role.role2
			]
		}
	`, tenantID, roleName1, roleName2)
}
