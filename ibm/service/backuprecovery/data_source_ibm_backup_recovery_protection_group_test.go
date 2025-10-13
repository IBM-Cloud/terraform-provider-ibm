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

func TestAccIbmBackupRecoveryProtectionGroupDataSourceBasic(t *testing.T) {
	groupName := fmt.Sprintf("tf_groupname_%d", acctest.RandIntRange(10, 100))
	policyName := fmt.Sprintf("tf_policyname_%d", acctest.RandIntRange(10, 100))
	objectId := 344
	environment := "kPhysical"
	includedPath := "/data2/data/"
	protectionType := "kFile"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmBackupRecoveryProtectionGroupDataSourceConfigBasic(groupName, environment, includedPath, protectionType, policyName, objectId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_protection_group.baas_protection_group_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_protection_group.baas_protection_group_instance", "group_id"),
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_protection_group.baas_protection_group_instance", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_protection_group.baas_protection_group_instance", "physical_params.#"),
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_protection_group.baas_protection_group_instance", "physical_params.0.file_protection_type_params.#"),
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_protection_group.baas_protection_group_instance", "physical_params.0.file_protection_type_params.0.objects.#"),
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_protection_group.baas_protection_group_instance", "physical_params.0.file_protection_type_params.0.objects.0.name"),
					resource.TestCheckResourceAttr("data.ibm_backup_recovery_protection_group.baas_protection_group_instance", "physical_params.0.file_protection_type_params.0.objects.0.id", strconv.Itoa(objectId)),
					resource.TestCheckResourceAttr("data.ibm_backup_recovery_protection_group.baas_protection_group_instance", "x_ibm_tenant_id", tenantId),
				),
			},
		},
	})
}

func testAccCheckIbmBackupRecoveryProtectionGroupDataSourceConfigBasic(name, environment, includedPath, protectionType, policyName string, objectId int) string {
	return fmt.Sprintf(`
		resource "ibm_backup_recovery_protection_policy" "baas_protection_policy_instance" {
			x_ibm_tenant_id = "%s"
			name = "%s"
			backup_recovery_endpoint = "https://protectiondomain0103.us-east.backup-recovery-tests.cloud.ibm.com/v2"
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
			backup_recovery_endpoint = "https://protectiondomain0103.us-east.backup-recovery-tests.cloud.ibm.com/v2"
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

		data "ibm_backup_recovery_protection_group" "baas_protection_group_instance" {
			protection_group_id = ibm_backup_recovery_protection_group.baas_protection_group_instance.group_id
			backup_recovery_endpoint = "https://protectiondomain0103.us-east.backup-recovery-tests.cloud.ibm.com/v2"
			x_ibm_tenant_id = "%[1]s"
		}
	`, tenantId, policyName, tenantId, name, environment, protectionType, objectId, includedPath)
}
