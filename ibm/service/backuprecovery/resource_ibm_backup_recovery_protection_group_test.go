// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package backuprecovery_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM/ibm-backup-recovery-sdk-go/backuprecoveryv1"
)

func TestAccIbmBackupRecoveryProtectionGroupBasic(t *testing.T) {
	var conf backuprecoveryv1.ProtectionGroupResponse
	groupName := fmt.Sprintf("tf_groupname_%d", acctest.RandIntRange(10, 100))
	policyName := fmt.Sprintf("tf_name_policy_%d", acctest.RandIntRange(10, 100))
	environment := "kPhysical"
	includedPath := "/data2/data/"
	includedPathUpdate := "/data1/"
	protectionType := "kFile"
	objectId := 344

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmBackupRecoveryProtectionGroupDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmBackupRecoveryProtectionGroupConfigBasic(groupName, environment, includedPath, protectionType, policyName, objectId),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmBackupRecoveryProtectionGroupExists("ibm_backup_recovery_protection_group.baas_protection_group_instance", conf),
					resource.TestCheckResourceAttr("ibm_backup_recovery_protection_group.baas_protection_group_instance", "x_ibm_tenant_id", tenantId),
					resource.TestCheckResourceAttr("ibm_backup_recovery_protection_group.baas_protection_group_instance", "name", groupName),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmBackupRecoveryProtectionGroupConfigBasic(groupName, environment, includedPathUpdate, protectionType, policyName, objectId),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_backup_recovery_protection_group.baas_protection_group_instance", "x_ibm_tenant_id", tenantId),
					resource.TestCheckResourceAttr("ibm_backup_recovery_protection_group.baas_protection_group_instance", "name", groupName),
				),
			},
		},
	})
}

func testAccCheckIbmBackupRecoveryProtectionGroupConfigBasic(name, environment, includedPath, protectionType, policyName string, objectId int) string {
	return fmt.Sprintf(`
			resource "ibm_backup_recovery_protection_policy" "baas_protection_policy_instance" {
				x_ibm_tenant_id = "%s"
				name = "%s"
				backup_policy {
						regular {
							incremental{
								schedule{
										day_schedule {
											frequency = 2
										}
										unit = "Days"
									}
							}
							retention {
								duration = 2
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
	`, tenantId, policyName, tenantId, name, environment, protectionType, objectId, includedPath)
}

func TestAccIbmBackupRecoveryProtectionGroupKubernetesBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmBackupRecoveryProtectionKubernetesGroupConfigBasic(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_backup_recovery_protection_group.baas_protection_group_instance", "x_ibm_tenant_id", "wkk1yqrdce/"),
					resource.TestCheckResourceAttr("ibm_backup_recovery_protection_group.baas_protection_group_instance", "name", "terra-test-group-100"),
				),
			},
		},
	})
}

func testAccCheckIbmBackupRecoveryProtectionKubernetesGroupConfigBasic() string {
	return fmt.Sprintf(`
	resource "ibm_backup_recovery_protection_group" "baas_protection_group_instance" {
		x_ibm_tenant_id = "wkk1yqrdce/"
		policy_id = "8305184241232842:1757331781254:65366"
		name = "terra-test-group-100"
		environment = "kKubernetes"
		priority = "kMedium"
		qos_policy = "kBackupHDD"
		kubernetes_params {
			objects {
				id = 3120
			}
		}
}
	`)
}

func testAccCheckIbmBackupRecoveryProtectionGroupExists(n string, obj backuprecoveryv1.ProtectionGroupResponse) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		backupRecoveryClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).BackupRecoveryV1()
		if err != nil {
			return err
		}

		getProtectionGroupByIdOptions := &backuprecoveryv1.GetProtectionGroupByIdOptions{}

		getProtectionGroupByIdOptions.SetID(rs.Primary.Attributes["group_id"])
		getProtectionGroupByIdOptions.SetXIBMTenantID(tenantId)

		protectionGroupResponse, _, err := backupRecoveryClient.GetProtectionGroupByID(getProtectionGroupByIdOptions)
		if err != nil {
			return err
		}

		obj = *protectionGroupResponse
		return nil
	}
}

func testAccCheckIbmBackupRecoveryProtectionGroupDestroy(s *terraform.State) error {
	backupRecoveryClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).BackupRecoveryV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_backup_recovery_protection_group" {
			continue
		}

		getProtectionGroupByIdOptions := &backuprecoveryv1.GetProtectionGroupByIdOptions{}

		getProtectionGroupByIdOptions.SetID(rs.Primary.Attributes["group_id"])
		getProtectionGroupByIdOptions.SetXIBMTenantID(tenantId)

		// Try to find the key
		groupResponse, response, err := backupRecoveryClient.GetProtectionGroupByID(getProtectionGroupByIdOptions)

		if err == nil {
			if strings.Contains(*groupResponse.Name, fmt.Sprintf("_DELETED_%s", rs.Primary.Attributes["name"])) {
				return nil
			}
			return fmt.Errorf("baas_protection_group still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for baas_protection_group (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
