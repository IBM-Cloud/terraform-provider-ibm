package appid_test

import (
	"fmt"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMAppIDIDPCloudDirectoryDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: setupIBMAppIDCloudDirectoryIDPConfig(acc.AppIDTenantID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_appid_idp_cloud_directory.idp", "tenant_id", acc.AppIDTenantID),
					resource.TestCheckResourceAttr("data.ibm_appid_idp_cloud_directory.idp", "is_active", "true"),
					resource.TestCheckResourceAttr("data.ibm_appid_idp_cloud_directory.idp", "self_service_enabled", "false"),
					resource.TestCheckResourceAttr("data.ibm_appid_idp_cloud_directory.idp", "signup_enabled", "false"),
					resource.TestCheckResourceAttr("data.ibm_appid_idp_cloud_directory.idp", "welcome_enabled", "true"),
					resource.TestCheckResourceAttr("data.ibm_appid_idp_cloud_directory.idp", "reset_password_enabled", "false"),
					resource.TestCheckResourceAttr("data.ibm_appid_idp_cloud_directory.idp", "reset_password_notification_enabled", "false"),
					resource.TestCheckResourceAttr("data.ibm_appid_idp_cloud_directory.idp", "identity_confirm_access_mode", "FULL"),
					resource.TestCheckResourceAttr("data.ibm_appid_idp_cloud_directory.idp", "identity_confirm_methods.#", "1"),
					resource.TestCheckResourceAttr("data.ibm_appid_idp_cloud_directory.idp", "identity_confirm_methods.0", "email"),
					resource.TestCheckResourceAttr("data.ibm_appid_idp_cloud_directory.idp", "identity_field", "email"),
				),
			},
		},
	})
}

func setupIBMAppIDCloudDirectoryIDPConfig(tenantID string) string {
	return fmt.Sprintf(`
		resource "ibm_appid_idp_cloud_directory" "idp" {
			tenant_id = "%s"
			is_active = true
			identity_confirm_methods = [
				"email"
			]
			identity_field = "email"
			self_service_enabled = false
			signup_enabled = false
			welcome_enabled = true
			reset_password_enabled = false
			reset_password_notification_enabled = false	
		}
		data "ibm_appid_idp_cloud_directory" "idp" {
			tenant_id = ibm_appid_idp_cloud_directory.idp.tenant_id
			depends_on = [
				ibm_appid_idp_cloud_directory.idp
			]
		}
	`, tenantID)
}
