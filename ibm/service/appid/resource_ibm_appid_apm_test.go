package appid_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"

	appid "github.com/IBM/appid-management-go-sdk/appidmanagementv4"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/Mavrickk3/bluemix-go/helpers"
	"github.com/google/go-cmp/cmp"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIBMAppIDAPM_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMAppIDAPMDestroy,
		Steps: []resource.TestStep{
			{
				Config: setupIBMAppIDAPMResourceConfig(acc.AppIDTenantID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_appid_apm.apm", "tenant_id", acc.AppIDTenantID),
					resource.TestCheckResourceAttr("ibm_appid_apm.apm", "enabled", "true"),
					resource.TestCheckResourceAttr("ibm_appid_apm.apm", "prevent_password_with_username", "true"),
					resource.TestCheckResourceAttr("ibm_appid_apm.apm", "password_reuse.0.enabled", "true"),
					resource.TestCheckResourceAttr("ibm_appid_apm.apm", "password_reuse.0.max_password_reuse", "4"),
					resource.TestCheckResourceAttr("ibm_appid_apm.apm", "password_expiration.0.enabled", "true"),
					resource.TestCheckResourceAttr("ibm_appid_apm.apm", "password_expiration.0.days_to_expire", "25"),
					resource.TestCheckResourceAttr("ibm_appid_apm.apm", "lockout_policy.0.enabled", "true"),
					resource.TestCheckResourceAttr("ibm_appid_apm.apm", "lockout_policy.0.lockout_time_sec", "2600"),
					resource.TestCheckResourceAttr("ibm_appid_apm.apm", "lockout_policy.0.num_of_attempts", "4"),
					resource.TestCheckResourceAttr("ibm_appid_apm.apm", "min_password_change_interval.0.enabled", "true"),
					resource.TestCheckResourceAttr("ibm_appid_apm.apm", "min_password_change_interval.0.min_hours_to_change_password", "1"),
				),
			},
		},
	})
}

func setupIBMAppIDAPMResourceConfig(tenantID string) string {
	return fmt.Sprintf(`
		resource "ibm_appid_apm" "apm" {
			tenant_id = "%s"
			enabled = true
			prevent_password_with_username = true
		
			password_reuse {
				enabled = true
				max_password_reuse = 4
			}
		
			password_expiration {
				enabled = true
				days_to_expire = 25
			}
		
			lockout_policy {
				enabled = true
				lockout_time_sec = 2600
				num_of_attempts = 4
			}
		
			min_password_change_interval {
				enabled = true
				min_hours_to_change_password = 1
			}
		}
	`, tenantID)
}

func testAccCheckIBMAppIDAPMDestroy(s *terraform.State) error {
	appIDClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).AppIDAPI()

	if err != nil {
		return err
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_appid_apm" {
			continue
		}

		tenantID := rs.Primary.ID

		config, _, err := appIDClient.GetCloudDirectoryAdvancedPasswordManagement(&appid.GetCloudDirectoryAdvancedPasswordManagementOptions{
			TenantID: &tenantID,
		})

		if err != nil {
			return fmt.Errorf("[ERROR] Error checking if AppID APM configuration was reset: %s", err)
		}

		// verify that configuration is reset to defaults
		defaults := &appid.ApmSchemaAdvancedPasswordManagement{
			Enabled: helpers.Bool(false),
			PasswordReuse: &appid.ApmSchemaAdvancedPasswordManagementPasswordReuse{
				Enabled: helpers.Bool(false),
				Config: &appid.ApmSchemaAdvancedPasswordManagementPasswordReuseConfig{
					MaxPasswordReuse: core.Int64Ptr(8),
				},
			},
			PasswordExpiration: &appid.ApmSchemaAdvancedPasswordManagementPasswordExpiration{
				Enabled: helpers.Bool(false),
				Config: &appid.ApmSchemaAdvancedPasswordManagementPasswordExpirationConfig{
					DaysToExpire: core.Int64Ptr(30),
				},
			},
			MinPasswordChangeInterval: &appid.ApmSchemaAdvancedPasswordManagementMinPasswordChangeInterval{
				Enabled: helpers.Bool(false),
				Config: &appid.ApmSchemaAdvancedPasswordManagementMinPasswordChangeIntervalConfig{
					MinHoursToChangePassword: core.Int64Ptr(0),
				},
			},
			LockOutPolicy: &appid.ApmSchemaAdvancedPasswordManagementLockOutPolicy{
				Enabled: helpers.Bool(false),
				Config: &appid.ApmSchemaAdvancedPasswordManagementLockOutPolicyConfig{
					LockOutTimeSec: core.Int64Ptr(1800),
					NumOfAttempts:  core.Int64Ptr(3),
				},
			},
			PreventPasswordWithUsername: &appid.ApmSchemaAdvancedPasswordManagementPreventPasswordWithUsername{
				Enabled: helpers.Bool(false),
			},
		}

		diff := cmp.Diff(config.AdvancedPasswordManagement, defaults)

		if config == nil || diff != "" {
			return fmt.Errorf("[ERROR] Error checking if AppID APM configuration was reset: %s", diff)
		}
	}

	return nil
}
