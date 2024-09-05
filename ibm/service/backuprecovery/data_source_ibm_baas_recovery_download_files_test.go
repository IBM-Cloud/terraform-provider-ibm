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
	recoveryDownloadFilesXIBMTenantID := fmt.Sprintf("tf_x_ibm_tenant_id_%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmRecoveryDownloadFilesDataSourceConfigBasic(recoveryDownloadFilesXIBMTenantID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_recovery_download_files.recovery_download_files_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_recovery_download_files.recovery_download_files_instance", "recovery_download_files_id"),
				),
			},
		},
	})
}

func testAccCheckIbmRecoveryDownloadFilesDataSourceConfigBasic(recoveryDownloadFilesXIBMTenantID string) string {
	return fmt.Sprintf(`
		data "ibm_recovery_download_files" "recovery_download_files_instance" {
			x_ibm_tenant_id = "%s"
			id = "id"
			startOffset = 1
			length = 1
			fileType = "fileType"
			sourceName = "sourceName"
			startTime = "startTime"
			includeTenants = true
		}
	`, recoveryDownloadFilesXIBMTenantID)
}
