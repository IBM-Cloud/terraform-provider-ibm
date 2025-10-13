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

func TestAccIbmRecoveryDownloadFilesDataSourceBasic(t *testing.T) {
	name := fmt.Sprintf("tf_recovery_download_files_folders_name_%d", acctest.RandIntRange(10, 100))
	objectId := 344
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmRecoveryDownloadFilesDataSourceConfigBasic(name, objectId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_download_files.recovery_download_files_instance", "id"),
					resource.TestCheckResourceAttr("data.ibm_backup_recovery_download_files.recovery_download_files_instance", "x_ibm_tenant_id", tenantId),
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_download_files.recovery_download_files_instance", "recovery_download_files_id"),
				),
			},
		},
	})
}

func testAccCheckIbmRecoveryDownloadFilesDataSourceConfigBasic(name string, objectId int) string {
	return fmt.Sprintf(`
	data "ibm_backup_recovery_object_snapshots" "baas_object_snapshots_instance" {
		x_ibm_tenant_id = "%s"
		backup_recovery_endpoint = "https://protectiondomain0103.us-east.backup-recovery-tests.cloud.ibm.com/v2"
		object_id = %d
	  }

	resource "ibm_backup_recovery_download_files_folders" "baas_recovery_download_files_folders_instance" {
		x_ibm_tenant_id = "%s"
		name = "%s"
		backup_recovery_endpoint = "https://protectiondomain0103.us-east.backup-recovery-tests.cloud.ibm.com/v2"
		object {
		  snapshot_id = data.ibm_backup_recovery_object_snapshots.baas_object_snapshots_instance.snapshots[0].id
		}
		files_and_folders {
			absolute_path = "/mnt"
		}
	  }
	  data "ibm_backup_recovery_download_files" "recovery_download_files_instance" {
		x_ibm_tenant_id = "%[1]s"
		backup_recovery_endpoint = "https://protectiondomain0103.us-east.backup-recovery-tests.cloud.ibm.com/v2"
		recovery_download_files_id = ibm_backup_recovery_download_files_folders.baas_recovery_download_files_folders_instance.id
	}
	`, tenantId, objectId, tenantId, name)

}
