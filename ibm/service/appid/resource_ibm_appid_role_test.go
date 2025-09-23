package appid_test

import (
	"fmt"
	"strings"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/conns"

	appid "github.com/IBM/appid-management-go-sdk/appidmanagementv4"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIBMAppIDRole_basic(t *testing.T) {
	appName := fmt.Sprintf("tf_testacc_app_%d", acctest.RandIntRange(10, 100))
	roleName := fmt.Sprintf("tf_testacc_role_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMAppIDRoleDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMAppIDRoleConfig(acc.AppIDTenantID, roleName, appName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_appid_role.test_role", "tenant_id", acc.AppIDTenantID),
					resource.TestCheckResourceAttr("ibm_appid_role.test_role", "name", roleName),
					resource.TestCheckResourceAttr("ibm_appid_role.test_role", "access.#", "1"),
					resource.TestCheckResourceAttrSet("ibm_appid_role.test_role", "access.0.application_id"),
					resource.TestCheckResourceAttr("ibm_appid_role.test_role", "access.0.scopes.#", "2"),
					resource.TestCheckResourceAttr("ibm_appid_role.test_role", "access.0.scopes.0", "scope1"),
					resource.TestCheckResourceAttr("ibm_appid_role.test_role", "access.0.scopes.1", "scope3"),
				),
			},
		},
	})
}

func testAccCheckIBMAppIDRoleConfig(tenantID string, roleName string, appName string) string {
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
		
		resource "ibm_appid_role" "test_role" {
			tenant_id = ibm_appid_application.test_app.tenant_id
			name = "%s"
			
			access {
				application_id = ibm_appid_application.test_app.client_id
				scopes = ["scope1", "scope3"]
			}

			depends_on = [
				ibm_appid_application_scopes.scopes
			]
		}	
	`, tenantID, appName, roleName)
}

func testAccCheckIBMAppIDRoleDestroy(s *terraform.State) error {
	appIDClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).AppIDAPI()

	if err != nil {
		return err
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_appid_role" {
			continue
		}

		id := rs.Primary.ID
		idParts := strings.Split(id, "/")

		tenantID := idParts[0]
		roleID := idParts[1]

		_, _, err := appIDClient.GetRole(&appid.GetRoleOptions{
			TenantID: &tenantID,
			RoleID:   &roleID,
		})

		if err == nil {
			return fmt.Errorf("[ERROR] Error checking if AppID role (%s) has been destroyed", rs.Primary.ID)
		}
	}

	return nil
}
