package appid_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIBMAppIDAPMDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: setupIBMAppIDAPMDataSourceConfig(acc.AppIDTenantID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_appid_apm.apm", "tenant_id", acc.AppIDTenantID),
					resource.TestCheckResourceAttr("data.ibm_appid_apm.apm", "enabled", "true"),
					resource.TestCheckResourceAttr("data.ibm_appid_apm.apm", "prevent_password_with_username", "true"),
					resource.TestCheckResourceAttr("data.ibm_appid_apm.apm", "password_reuse.0.enabled", "true"),
					resource.TestCheckResourceAttr("data.ibm_appid_apm.apm", "password_reuse.0.max_password_reuse", "4"),
					resource.TestCheckResourceAttr("data.ibm_appid_apm.apm", "password_expiration.0.enabled", "true"),
					resource.TestCheckResourceAttr("data.ibm_appid_apm.apm", "password_expiration.0.days_to_expire", "25"),
					resource.TestCheckResourceAttr("data.ibm_appid_apm.apm", "lockout_policy.0.enabled", "true"),
					resource.TestCheckResourceAttr("data.ibm_appid_apm.apm", "lockout_policy.0.lockout_time_sec", "2600"),
					resource.TestCheckResourceAttr("data.ibm_appid_apm.apm", "lockout_policy.0.num_of_attempts", "4"),
					resource.TestCheckResourceAttr("data.ibm_appid_apm.apm", "min_password_change_interval.0.enabled", "true"),
					resource.TestCheckResourceAttr("data.ibm_appid_apm.apm", "min_password_change_interval.0.min_hours_to_change_password", "1"),
				),
			},
		},
	})
}

func setupIBMAppIDAPMDataSourceConfig(tenantID string) string {
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

		data "ibm_appid_apm" "apm" {
			tenant_id = ibm_appid_apm.apm.tenant_id

			depends_on = [
				ibm_appid_apm.apm
			]
		}
	`, tenantID)
}
