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
	"github.ibm.com/BackupAndRecovery/ibm-backup-recovery-sdk-go/backuprecoveryv1"
)

func TestAccIbmBaasRecoveryBasic(t *testing.T) {
	name := fmt.Sprintf("tf_recovery_name_%d", acctest.RandIntRange(10, 100))
	snapshotEnvironment := "kPhysical"
	objectId := 72
	targetenvironment := "kPhysical"
	absolutePath := "/"
	restoreEntityType := "kRegular"
	recoveryAction := "RecoverFiles"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Destroy: false,
				Config:  testAccCheckIbmBaasRecoveryConfigBasic(objectId, name, snapshotEnvironment, targetenvironment, absolutePath, restoreEntityType, recoveryAction),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmBaasRecoveryExists("ibm_baas_recovery.baas_recovery_instance"),
					resource.TestCheckResourceAttr("ibm_baas_recovery.baas_recovery_instance", "x_ibm_tenant_id", tenantId),
					resource.TestCheckResourceAttr("ibm_baas_recovery.baas_recovery_instance", "name", name),
				),
			},
		},
	})
}

func testAccCheckIbmBaasRecoveryConfigBasic(objectId int, name, snapshotEnvironment, targetenvironment, absolutePath, restoreEntityType, recoveryAction string) string {
	return fmt.Sprintf(`

	data "ibm_baas_object_snapshots" "object_snapshot" {
		x_ibm_tenant_id = "%s"
		baas_object_id = %d
	  }

	resource "ibm_baas_recovery" "baas_recovery_instance" {
		x_ibm_tenant_id = "%s"
		snapshot_environment = "%s"
		name = "%s"
		physical_params {
		  recovery_action = "%s"
		  objects {
			snapshot_id = data.ibm_baas_object_snapshots.object_snapshot.snapshots.0.id
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

func testAccCheckIbmBaasRecoveryExists(n string) resource.TestCheckFunc {

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

		getRecoveryByIdOptions.SetID(rs.Primary.ID)
		getRecoveryByIdOptions.SetXIBMTenantID(tenantId)

		_, _, err = backupRecoveryClient.GetRecoveryByID(getRecoveryByIdOptions)
		if err != nil {
			return err
		}
		return nil
	}
}
