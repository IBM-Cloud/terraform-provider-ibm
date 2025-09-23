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

func TestAccIBMAppIDApplicationRoles_basic(t *testing.T) {
	appName := fmt.Sprintf("tf_testacc_app_roles_%d", acctest.RandIntRange(10, 100))
	roleName := fmt.Sprintf("tf_testacc_app_roles_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMAppIDApplicationRolesDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMAppIDApplicationRolesConfig(acc.AppIDTenantID, appName, roleName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_appid_application_roles.roles", "roles.#", "1"),
					resource.TestCheckResourceAttrPair("ibm_appid_role.role", "role_id", "ibm_appid_application_roles.roles", "roles.0"),
				),
			},
		},
	})
}

func testAccCheckIBMAppIDApplicationRolesConfig(tenantID string, appName string, roleName string) string {
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
	`, tenantID, appName, roleName)
}

func testAccCheckIBMAppIDApplicationRolesDestroy(s *terraform.State) error {
	appIDClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).AppIDAPI()

	if err != nil {
		return err
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_appid_application_roles" {
			continue
		}

		id := rs.Primary.ID
		idParts := strings.Split(id, "/")

		tenantID := idParts[0]
		clientID := idParts[1]

		_, _, err := appIDClient.GetApplicationRoles(&appid.GetApplicationRolesOptions{
			TenantID: &tenantID,
			ClientID: &clientID,
		})

		if err == nil {
			return fmt.Errorf("[ERROR] Error checking if AppID application roles resource (%s) has been destroyed", rs.Primary.ID)
		}
	}

	return nil
}
