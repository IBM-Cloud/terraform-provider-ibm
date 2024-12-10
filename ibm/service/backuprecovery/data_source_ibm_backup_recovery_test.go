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

func TestAccIbmBackupRecoveryDataSourceBasic(t *testing.T) {
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
				Config: testAccCheckIbmBackupRecoveryDataSourceConfigBasic(objectId, name, snapshotEnvironment, targetenvironment, absolutePath, restoreEntityType, recoveryAction),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery.baas_recovery_instance", "id"),
					resource.TestCheckResourceAttr("data.ibm_backup_recovery.baas_recovery_instance", "x_ibm_tenant_id", tenantId),
					resource.TestCheckResourceAttr("data.ibm_backup_recovery.baas_recovery_instance", "name", name),
					resource.TestCheckResourceAttr("data.ibm_backup_recovery.baas_recovery_instance", "physical_params.0.objects.0.object_info.0.id", strconv.Itoa(objectId)),
				),
			},
		},
	})
}

func testAccCheckIbmBackupRecoveryDataSourceConfigBasic(objectId int, name, snapshotEnvironment, targetenvironment, absolutePath, restoreEntityType, recoveryAction string) string {

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

	  data "ibm_backup_recovery" "baas_recovery_instance" {
		recovery_id = ibm_backup_recovery.baas_recovery_instance.recovery_id
		x_ibm_tenant_id = "%[1]s"
	}

	`, tenantId, objectId, tenantId, snapshotEnvironment, name, recoveryAction, targetenvironment, absolutePath, objectId, restoreEntityType, absolutePath)

}
