// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package backuprecovery_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/IBM/ibm-backup-recovery-sdk-go/backuprecoveryv1"
	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/conns"
)

func TestAccIbmBackupRecoveryProtectionPolicyBasic(t *testing.T) {
	var conf backuprecoveryv1.ProtectionPolicyResponse
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	duration := 1
	durationUpdate := 2

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmBackupRecoveryProtectionPolicyDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmBackupRecoveryProtectionPolicyConfigBasic(name, duration),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmBackupRecoveryProtectionPolicyExists("ibm_backup_recovery_protection_policy.baas_protection_policy_instance", conf),
					resource.TestCheckResourceAttr("ibm_backup_recovery_protection_policy.baas_protection_policy_instance", "x_ibm_tenant_id", tenantId),
					resource.TestCheckResourceAttr("ibm_backup_recovery_protection_policy.baas_protection_policy_instance", "name", name),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmBackupRecoveryProtectionPolicyConfigBasic(name, durationUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_backup_recovery_protection_policy.baas_protection_policy_instance", "x_ibm_tenant_id", tenantId),
					resource.TestCheckResourceAttr("ibm_backup_recovery_protection_policy.baas_protection_policy_instance", "name", name),
				),
			},
		},
	})
}

func testAccCheckIbmBackupRecoveryProtectionPolicyConfigBasic(name string, duration int) string {
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
	`, tenantId, name, duration)
}

func testAccCheckIbmBackupRecoveryProtectionPolicyExists(n string, obj backuprecoveryv1.ProtectionPolicyResponse) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		backupRecoveryClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).BackupRecoveryV1()
		if err != nil {
			return err
		}

		getProtectionPolicyByIdOptions := &backuprecoveryv1.GetProtectionPolicyByIdOptions{}

		getProtectionPolicyByIdOptions.SetXIBMTenantID(tenantId)
		getProtectionPolicyByIdOptions.SetID(rs.Primary.Attributes["policy_id"])

		protectionPolicyResponse, _, err := backupRecoveryClient.GetProtectionPolicyByID(getProtectionPolicyByIdOptions)
		if err != nil {
			return err
		}

		obj = *protectionPolicyResponse
		return nil
	}
}

func testAccCheckIbmBackupRecoveryProtectionPolicyDestroy(s *terraform.State) error {
	backupRecoveryClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).BackupRecoveryV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_backup_recovery_protection_policy" {
			continue
		}

		getProtectionPolicyByIdOptions := &backuprecoveryv1.GetProtectionPolicyByIdOptions{}

		getProtectionPolicyByIdOptions.SetXIBMTenantID(tenantId)
		getProtectionPolicyByIdOptions.SetID(rs.Primary.Attributes["policy_id"])

		// Try to find the key
		policyResponse, response, err := backupRecoveryClient.GetProtectionPolicyByID(getProtectionPolicyByIdOptions)

		if err == nil {
			if strings.Contains(*policyResponse.Name, fmt.Sprintf("%s_DELETED", rs.Primary.Attributes["name"])) {
				return nil
			}
			return fmt.Errorf("baas_protection_policy still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for baas_protection_policy (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
