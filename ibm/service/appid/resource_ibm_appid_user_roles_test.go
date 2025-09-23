package appid_test

import (
	"fmt"
	"strings"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/conns"

	appid "github.com/IBM/appid-management-go-sdk/appidmanagementv4"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMAppIDUserRolesRoles_basic(t *testing.T) {
	roleName := fmt.Sprintf("tf_testacc_role_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMAppIDUserRolesDestroy,
		Steps: []resource.TestStep{
			{
				Config: setupAppIDUserRolesConfig(acc.AppIDTenantID, roleName, acc.AppIDTestUserEmail),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_appid_user_roles.roles", "role_ids.#", "1"),
					resource.TestCheckResourceAttrPair("ibm_appid_user_roles.roles", "role_ids.0", "ibm_appid_role.role", "role_id"),
				),
			},
		},
	})
}

// Test assumes there are no pre-existing roles
func setupAppIDUserRolesConfig(tenantID string, roleName string, email string) string {
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
	`, tenantID, roleName, email)
}

func testAccCheckIBMAppIDUserRolesDestroy(s *terraform.State) error {
	appIDClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).AppIDAPI()

	if err != nil {
		return err
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_appid_user_roles" {
			continue
		}

		id := rs.Primary.ID
		idParts := strings.Split(id, "/")

		tenantID := idParts[0]
		subject := idParts[1]

		roles, _, err := appIDClient.GetUserRoles(&appid.GetUserRolesOptions{
			TenantID: &tenantID,
			ID:       &subject,
		})

		if err != nil {
			return fmt.Errorf("[ERROR] Error checking if AppID user roles have been destroyed: %s", err)
		}

		if roles.Roles != nil && len(roles.Roles) > 0 {
			return fmt.Errorf("[ERROR] Error checking if AppID user roles have been destroyed")
		}
	}

	return nil
}
