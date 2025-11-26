// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.94.0-fa797aec-20240814-142622
 */

package backuprecovery_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmBackupRecoveryDownloadIndexedFilesDataSourceBasic(t *testing.T) {
	objectId := 344
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmBackupRecoveryDownloadIndexedFilesDataSourceConfigBasic(objectId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_download_indexed_files.baas_download_indexed_files_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_download_indexed_files.baas_download_indexed_files_instance", "x_ibm_tenant_id"),
				),
			},
		},
	})
}

func testAccCheckIbmBackupRecoveryDownloadIndexedFilesDataSourceConfigBasic(objectId int) string {

	return fmt.Sprintf(`
		data "ibm_backup_recovery_object_snapshots" "baas_object_snapshots_instance" {
			x_ibm_tenant_id = "%s"
			
			object_id = %d
		}

		data "ibm_backup_recovery_download_indexed_files" "baas_download_indexed_files_instance" {
			x_ibm_tenant_id = "%s"
			snapshots_id = data.ibm_backup_recovery_object_snapshots.baas_object_snapshots_instance.snapshots.0.id
			
			file_path = "/mnt"
		}
	`, tenantId, objectId, tenantId)
}
