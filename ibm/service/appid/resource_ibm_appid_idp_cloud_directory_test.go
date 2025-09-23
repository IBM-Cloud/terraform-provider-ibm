package appid_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"

	appid "github.com/IBM/appid-management-go-sdk/appidmanagementv4"
	"github.com/Mavrickk3/bluemix-go/helpers"
	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIBMAppIDIDPCloudDirectory_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMAppIDIDPCloudDirectoryDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMAppIDIDPCloudDirectoryConfig(acc.AppIDTenantID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_appid_idp_cloud_directory.idp", "tenant_id", acc.AppIDTenantID),
					resource.TestCheckResourceAttr("ibm_appid_idp_cloud_directory.idp", "is_active", "true"),
					resource.TestCheckResourceAttr("ibm_appid_idp_cloud_directory.idp", "self_service_enabled", "false"),
					resource.TestCheckResourceAttr("ibm_appid_idp_cloud_directory.idp", "signup_enabled", "false"),
					resource.TestCheckResourceAttr("ibm_appid_idp_cloud_directory.idp", "welcome_enabled", "true"),
					resource.TestCheckResourceAttr("ibm_appid_idp_cloud_directory.idp", "reset_password_enabled", "false"),
					resource.TestCheckResourceAttr("ibm_appid_idp_cloud_directory.idp", "reset_password_notification_enabled", "false"),
					resource.TestCheckResourceAttr("ibm_appid_idp_cloud_directory.idp", "identity_confirm_access_mode", "FULL"),
					resource.TestCheckResourceAttr("ibm_appid_idp_cloud_directory.idp", "identity_confirm_methods.#", "1"),
					resource.TestCheckResourceAttr("ibm_appid_idp_cloud_directory.idp", "identity_confirm_methods.0", "email"),
					resource.TestCheckResourceAttr("ibm_appid_idp_cloud_directory.idp", "identity_field", "email"),
				),
			},
		},
	})
}

func testAccCheckIBMAppIDIDPCloudDirectoryConfig(tenantID string) string {
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
	`, tenantID)
}

func testAccCheckIBMAppIDIDPCloudDirectoryDestroy(s *terraform.State) error {
	appIDClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).AppIDAPI()

	if err != nil {
		return err
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_appid_idp_cloud_directory" {
			continue
		}

		tenantID := rs.Primary.ID

		config, _, err := appIDClient.GetCloudDirectoryIDP(&appid.GetCloudDirectoryIDPOptions{
			TenantID: &tenantID,
		})

		if err != nil {
			return fmt.Errorf("[ERROR] Error checking if AppID IDP Cloud Directory configuration was reset: %s", err)
		}

		// verify that configuration is reset to defaults
		defaults := &appid.SetCloudDirectoryIDPOptions{
			TenantID: &tenantID,
			IsActive: helpers.Bool(false),
			Config: &appid.CloudDirectoryConfigParams{
				SelfServiceEnabled: helpers.Bool(true),
				SignupEnabled:      helpers.Bool(true),
				Interactions: &appid.CloudDirectoryConfigParamsInteractions{
					IdentityConfirmation: &appid.CloudDirectoryConfigParamsInteractionsIdentityConfirmation{
						AccessMode: helpers.String("FULL"),
						Methods:    []string{"email"},
					},
					WelcomeEnabled:                  helpers.Bool(true),
					ResetPasswordEnabled:            helpers.Bool(true),
					ResetPasswordNotificationEnable: helpers.Bool(true),
				},
			},
		}

		diff := cmp.Diff(&appid.SetCloudDirectoryIDPOptions{
			TenantID: &tenantID,
			IsActive: config.IsActive,
			Config:   config.Config,
		}, defaults)

		if config == nil || diff != "" {
			return fmt.Errorf("[ERROR] Error checking if AppID IDP Cloud Directory configuration was reset: %s", diff)
		}
	}

	return nil
}
