// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.94.0-fa797aec-20240814-142622
 */

package backuprecovery_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmBackupRecoveryProtectionPolicyDataSourceBasic(t *testing.T) {

	name := fmt.Sprintf("tf_policyname_%d", acctest.RandIntRange(10, 100))
	duration := 1

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmBackupRecoveryProtectionPolicyDataSourceConfigBasic(name, duration),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_protection_policy.baas_protection_policy_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_protection_policy.baas_protection_policy_instance", "policy_id"),
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_protection_policy.baas_protection_policy_instance", "x_ibm_tenant_id"),
					resource.TestCheckResourceAttr("data.ibm_backup_recovery_protection_policy.baas_protection_policy_instance", "name", name),
					resource.TestCheckResourceAttr("data.ibm_backup_recovery_protection_policy.baas_protection_policy_instance", "backup_policy.0.regular.0.retention.0.duration", strconv.Itoa(duration)),
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_protection_policy.baas_protection_policy_instance", "backup_policy.0.regular.0.retention.0.unit"),
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_protection_policy.baas_protection_policy_instance", "backup_policy.0.regular.0.incremental.0.schedule.#"),
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_protection_policy.baas_protection_policy_instance", "backup_policy.0.regular.0.incremental.0.schedule.0.unit"),
				),
			},
		},
	})
}

func testAccCheckIbmBackupRecoveryProtectionPolicyDataSourceConfigBasic(name string, duration int) string {
	return fmt.Sprintf(`
		resource "ibm_backup_recovery_protection_policy" "baas_protection_policy_instance" {
			x_ibm_tenant_id = "%s"
			name = "%s"
			backup_policy {
					regular {
						incremental{
							schedule{
									day_schedule {
										frequency = 1
									}
									unit = "Days"
								}
						}
						retention {
							duration = %d
							unit = "Weeks"
						}
						primary_backup_target {
							use_default_backup_target = true
						}
					}
			}
			retry_options {
			retries = 3
			retry_interval_mins = 5
			}
		}

		data "ibm_backup_recovery_protection_policy" "baas_protection_policy_instance" {
			protection_policy_id = ibm_backup_recovery_protection_policy.baas_protection_policy_instance.policy_id
			x_ibm_tenant_id = "%[1]s"
		}

	`, tenantId, name, duration)

}
