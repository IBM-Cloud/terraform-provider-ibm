package appid_test

import (
	"fmt"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMAppIDUserRolesRolesDataSource_basic(t *testing.T) {
	roleName := fmt.Sprintf("tf_testacc_role_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: setupAppIDUserRolesDataSourceConfig(acc.AppIDTenantID, roleName, acc.AppIDTestUserEmail),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_appid_user_roles.roles", "roles.#", "1"),
					resource.TestCheckResourceAttrPair("data.ibm_appid_user_roles.roles", "roles.0.id", "ibm_appid_role.role", "role_id"),
					resource.TestCheckResourceAttrPair("data.ibm_appid_user_roles.roles", "roles.0.name", "ibm_appid_role.role", "name"),
				),
			},
		},
	})
}

// Test assumes there are no pre-existing roles
func setupAppIDUserRolesDataSourceConfig(tenantID string, roleName string, email string) string {
	return fmt.Sprintf(`	
		resource "ibm_appid_role" "role" {
			tenant_id = "%s"
			name = "%s"
			description = "test role"
		}

		resource "ibm_appid_cloud_directory_user" "test_user" {
			tenant_id = ibm_appid_role.role.tenant_id
			email {
				value = "%s"
				primary = true
			}
			password = "P@ssw0rd"
			status = "PENDING"
		}

		resource "ibm_appid_user_roles" "roles" {
			tenant_id = ibm_appid_role.role.tenant_id
			subject = ibm_appid_cloud_directory_user.test_user.subject
			role_ids = [ibm_appid_role.role.role_id]
		}

		data "ibm_appid_user_roles" "roles" {
			tenant_id = ibm_appid_role.role.tenant_id
			subject = ibm_appid_cloud_directory_user.test_user.subject

			depends_on = [
				ibm_appid_user_roles.roles
			]
		}
	`, tenantID, roleName, email)
}
