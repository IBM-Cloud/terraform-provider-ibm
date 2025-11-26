// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package backuprecovery_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/IBM/ibm-backup-recovery-sdk-go/backuprecoveryv1"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
)

func TestAccIbmBackupRecoveryDownloadFilesFoldersBasic(t *testing.T) {
	name := fmt.Sprintf("tf_recovery_download_files_folders_name_%d", acctest.RandIntRange(10, 100))
	objectId := 344
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmBackupRecoveryDownloadFilesFoldersConfigBasic(name, objectId),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmBackupRecoveryDownloadFilesFoldersExists("ibm_backup_recovery_download_files_folders.baas_recovery_download_files_folders_instance"),
					resource.TestCheckResourceAttr("ibm_backup_recovery_download_files_folders.baas_recovery_download_files_folders_instance", "x_ibm_tenant_id", tenantId),
					resource.TestCheckResourceAttr("ibm_backup_recovery_download_files_folders.baas_recovery_download_files_folders_instance", "name", name),
					resource.TestCheckResourceAttrSet("ibm_backup_recovery_download_files_folders.baas_recovery_download_files_folders_instance", "id"),
					resource.TestCheckResourceAttr("ibm_backup_recovery_download_files_folders.baas_recovery_download_files_folders_instance", "recovery_physical_params.#", "1"),
					resource.TestCheckResourceAttr("ibm_backup_recovery_download_files_folders.baas_recovery_download_files_folders_instance", "recovery_physical_params.0.objects.#", "1"),
					resource.TestCheckResourceAttr("ibm_backup_recovery_download_files_folders.baas_recovery_download_files_folders_instance", "recovery_physical_params.0.objects.0.object_info.#", "1"),
					resource.TestCheckResourceAttr("ibm_backup_recovery_download_files_folders.baas_recovery_download_files_folders_instance", "recovery_physical_params.0.objects.0.object_info.0.id", strconv.Itoa(objectId)),
				),
				Destroy: false,
			},
		},
	})
}

func testAccCheckIbmBackupRecoveryDownloadFilesFoldersConfigBasic(name string, objectId int) string {
	return fmt.Sprintf(`
	
	data "ibm_backup_recovery_object_snapshots" "baas_object_snapshots_instance" {
		x_ibm_tenant_id = "%s"
		
		object_id = %d
	  }

	resource "ibm_backup_recovery_download_files_folders" "baas_recovery_download_files_folders_instance" {
		x_ibm_tenant_id = "%s"
		name = "%s"
		
		object {
		  snapshot_id = data.ibm_backup_recovery_object_snapshots.baas_object_snapshots_instance.snapshots[0].id
		}
		files_and_folders {
			absolute_path = "/mnt"
		}
	  }
	`, tenantId, objectId, tenantId, name)
}

func testAccCheckIbmBackupRecoveryDownloadFilesFoldersExists(n string) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		backupRecoveryClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).BackupRecoveryV1()
		if err != nil {
			return err
		}

		getDownloadFilesFromRecoveryOptionsOptions := &backuprecoveryv1.DownloadFilesFromRecoveryOptions{}

		getDownloadFilesFromRecoveryOptionsOptions.SetID(rs.Primary.ID)
		getDownloadFilesFromRecoveryOptionsOptions.SetXIBMTenantID(tenantId)

		_, err = backupRecoveryClient.DownloadFilesFromRecovery(getDownloadFilesFromRecoveryOptionsOptions)
		if err != nil {
			return err
		}
		return nil
	}
}
