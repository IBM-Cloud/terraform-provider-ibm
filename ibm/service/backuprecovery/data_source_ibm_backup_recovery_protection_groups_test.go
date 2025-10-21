// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.94.0-fa797aec-20240814-142622
 */

package backuprecovery_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmBackupRecoveryProtectionGroupsDataSourceBasic(t *testing.T) {
	groupName := fmt.Sprintf("tf_groupname_%d", acctest.RandIntRange(10, 100))
	policyName := fmt.Sprintf("tf_policyname_%d", acctest.RandIntRange(10, 100))
	objectId := 344
	environment := "kPhysical"
	includedPath := "/data1/data/dat2/"
	protectionType := "kFile"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmBackupRecoveryProtectionGroupsDataSourceConfigBasic(groupName, environment, includedPath, protectionType, policyName, objectId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_protection_groups.baas_protection_groups_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_protection_groups.baas_protection_groups_instance", "x_ibm_tenant_id"),
					resource.TestCheckResourceAttr("data.ibm_backup_recovery_protection_groups.baas_protection_groups_instance", "protection_groups.#", "1"),
					resource.TestCheckResourceAttr("data.ibm_backup_recovery_protection_groups.baas_protection_groups_instance", "protection_groups.0.name", groupName),
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_protection_groups.baas_protection_groups_instance", "protection_groups.0.id"),
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_protection_groups.baas_protection_groups_instance", "protection_groups.0.permissions.#"),
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_protection_groups.baas_protection_groups_instance", "protection_groups.0.physical_params.#"),
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_protection_groups.baas_protection_groups_instance", "protection_groups.0.sla.#"),
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_protection_groups.baas_protection_groups_instance", "protection_groups.0.start_time.#"),
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_protection_groups.baas_protection_groups_instance", "protection_groups.0.policy_id"),
				),
			},
		},
	})
}

func TestAccIbmBackupRecoveryProtectionGroupsDataSourceKubernetesBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmBackupRecoveryProtectionGroupsDataSourceKubernetesConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_protection_groups.baas_protection_groups_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_protection_groups.baas_protection_groups_instance", "x_ibm_tenant_id"),
				),
			},
		},
	})
}

func testAccCheckIbmBackupRecoveryProtectionGroupsDataSourceConfigBasic(name, environment, includedPath, protectionType, policyName string, objectId int) string {
	return fmt.Sprintf(`
	resource "ibm_backup_recovery_protection_policy" "baas_protection_policy_instance" {
		x_ibm_tenant_id = "%s"
		name = "%s"
		

		backup_policy {
				regular {
					incremental{
						schedule{
								day_schedule {
									frequency = 10
								}
								unit = "Days"
							}
					}
					retention {
						duration = 1
						unit = "Weeks"
					}
					primary_backup_target {
						use_default_backup_target = true
					}
				}
		}
		retry_options {
			retries = 0
			retry_interval_mins = 10
		}
	}

	resource "ibm_backup_recovery_protection_group" "baas_protection_group_instance" {
		x_ibm_tenant_id = "%s"
		policy_id = ibm_backup_recovery_protection_policy.baas_protection_policy_instance.policy_id
		
		name = "%s"
		environment = "%s"
		physical_params {
			protection_type = "%s"
			file_protection_type_params {
			objects {
				id = %d
				file_paths{
					included_path = "%s"
				}
			}
			}
		}
	}
		data "ibm_backup_recovery_protection_groups" "baas_protection_groups_instance" {
			x_ibm_tenant_id = "%[1]s"
			
			ids = [ ibm_backup_recovery_protection_group.baas_protection_group_instance.group_id ]
		}
	`, tenantId, policyName, tenantId, name, environment, protectionType, objectId, includedPath)
}

func testAccCheckIbmBackupRecoveryProtectionGroupsDataSourceKubernetesConfigBasic() string {
	return `
		data "ibm_backup_recovery_protection_groups" "baas_protection_groups_instance" {
			x_ibm_tenant_id = "wkk1yqrdce/"
		}
	`
}
