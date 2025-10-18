// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package backuprecovery_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM/ibm-backup-recovery-sdk-go/backuprecoveryv1"
)

func TestAccIbmBackupRecoveryBasic(t *testing.T) {
	name := fmt.Sprintf("tf_recovery_name_%d", acctest.RandIntRange(10, 100))
	snapshotEnvironment := "kPhysical"
	objectId := 18
	targetenvironment := "kPhysical"
	absolutePath := "/data/"
	restoreEntityType := "kRegular"
	recoveryAction := "RecoverFiles"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Destroy: false,
				Config:  testAccCheckIbmBackupRecoveryConfigBasic(objectId, name, snapshotEnvironment, targetenvironment, absolutePath, restoreEntityType, recoveryAction),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmBackupRecoveryExists("ibm_backup_recovery.baas_recovery_instance"),
					resource.TestCheckResourceAttr("ibm_backup_recovery.baas_recovery_instance", "x_ibm_tenant_id", tenantId),
					resource.TestCheckResourceAttr("ibm_backup_recovery.baas_recovery_instance", "name", name),
				),
			},
		},
	})
}

func testAccCheckIbmBackupRecoveryConfigBasic(objectId int, name, snapshotEnvironment, targetenvironment, absolutePath, restoreEntityType, recoveryAction string) string {
	return fmt.Sprintf(`

	data "ibm_backup_recovery_object_snapshots" "object_snapshot" {
		x_ibm_tenant_id = "%s"
		object_id = %d
	  }

	resource "ibm_backup_recovery" "baas_recovery_instance" {
		x_ibm_tenant_id = "%s"
		snapshot_environment = "%s"
		name = "%s"
		physical_params {
		  recovery_action = "%s"
		  objects {
			snapshot_id = data.ibm_backup_recovery_object_snapshots.object_snapshot.snapshots.0.id
		  }
		  recover_file_and_folder_params {
			 target_environment = "%s"
			 files_and_folders {
			   absolute_path = "%s"
			 }
			 physical_target_params {
			   recover_target {
				 id = %d
			   }
			   restore_entity_type = "%s"
			   alternate_restore_directory = "%s"
			 }
		  }
		}
	  }
	`, tenantId, objectId, tenantId, snapshotEnvironment, name, recoveryAction, targetenvironment, absolutePath, objectId, restoreEntityType, absolutePath)
}

func TestAccIbmBackupRecoveryKubernetesBasic(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Destroy: false,
				Config:  testAccCheckIbmBackupRecoveryKubernetesConfigBasic(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_backup_recovery.baas_recovery_instance", "name", "terra-k8s-recovery-1"),
				),
			},
		},
	})
}

func testAccCheckIbmBackupRecoveryKubernetesConfigBasic() string {
	return fmt.Sprintf(`

	 resource "ibm_backup_recovery" "baas_recovery_instance" {
		 	x_ibm_tenant_id = "wkk1yqrdce/"
		 	snapshot_environment = "kKubernetes"
		 	name = "terra-k8s-recovery-1"
		 	kubernetes_params {
		 		recovery_action = "RecoverNamespaces"
		 		recover_namespace_params {
		 			target_environment = "kKubernetes"
		 			kubernetes_target_params {
		 				objects {
		 					snapshot_id = "eyJhX2NsdXN0ZXJJZCI6ODMwNTE4NDI0MTIzMjg0MiwiYl9jbHVzdGVySW5jYXJuYXRpb25JZCI6MTc1NzMzMTc4MTI1NCwiY19qb2JJZCI6MjQ4MDM2LCJlX2pvYkluc3RhbmNlSWQiOjI0ODA1MywiZl9ydW5TdGFydFRpbWVVc2VjcyI6MTc2MDI3MDk2NjI2ODg4MywiZ19vYmplY3RJZCI6MzEyMCwiaF92YXVsdElkIjoxMjQzNDU3fQ=="
		 				}
		 				rename_recovered_namespaces_params {
		 				  prefix = "0-copy-sdk-workload"
		 				}
		 				recovery_target_config {
		 					recover_to_new_source = false
		 				}
		 			}
		 		}
		 	}
		 }


	`)
}

func testAccCheckIbmBackupRecoveryExists(n string) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found ...: %s", n)
		}

		backupRecoveryClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).BackupRecoveryV1()
		if err != nil {
			return err
		}

		getRecoveryByIdOptions := &backuprecoveryv1.GetRecoveryByIdOptions{}

		getRecoveryByIdOptions.SetID(rs.Primary.Attributes["recovery_id"])
		getRecoveryByIdOptions.SetXIBMTenantID(tenantId)

		_, _, err = backupRecoveryClient.GetRecoveryByID(getRecoveryByIdOptions)
		if err != nil {
			return err
		}
		return nil
	}
}
