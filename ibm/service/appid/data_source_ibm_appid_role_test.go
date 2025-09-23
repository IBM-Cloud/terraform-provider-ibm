package appid_test

import (
	"fmt"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMAppIDRoleDataSource_basic(t *testing.T) {
	roleName := fmt.Sprintf("tf_testacc_role_%d", acctest.RandIntRange(10, 100))
	appName := fmt.Sprintf("tf_testacc_role_%d", acctest.RandIntRange(10, 100))
	description := "test role"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: setupRoleConfig(acc.AppIDTenantID, appName, roleName, description),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_appid_role.role", "tenant_id", acc.AppIDTenantID),
					resource.TestCheckResourceAttr("data.ibm_appid_role.role", "name", roleName),
					resource.TestCheckResourceAttr("data.ibm_appid_role.role", "description", description),
					resource.TestCheckResourceAttr("data.ibm_appid_role.role", "access.#", "1"),
					resource.TestCheckResourceAttrSet("data.ibm_appid_role.role", "access.0.application_id"),
					resource.TestCheckResourceAttr("data.ibm_appid_role.role", "access.0.scopes.#", "1"),
					resource.TestCheckResourceAttr("data.ibm_appid_role.role", "access.0.scopes.0", "pancakes"),
				),
			},
		},
	})
}

func setupRoleConfig(tenantID string, appName string, roleName string, description string) string {
	return fmt.Sprintf(`
		resource "ibm_appid_application" "app" {
			tenant_id = "%s"
			name = "%s"  
			type = "singlepageapp"
	  	}

		resource "ibm_appid_application_scopes" "scopes" {
			tenant_id = ibm_appid_application.app.tenant_id
			client_id = ibm_appid_application.app.client_id
			
			scopes = ["pancakes", "cartoons"]
		}

		resource "ibm_appid_role" "role" {
			tenant_id = ibm_appid_application.app.tenant_id
			name = "%s"
			description = "%s"
			access {
				application_id = ibm_appid_application.app.client_id
				scopes = [
					"pancakes",
				]
			}

			depends_on = [
				ibm_appid_application_scopes.scopes
			]
		}

		data "ibm_appid_role" "role" {
			tenant_id = ibm_appid_role.role.tenant_id
			role_id = ibm_appid_role.role.role_id
		}
	`, tenantID, appName, roleName, description)
}
