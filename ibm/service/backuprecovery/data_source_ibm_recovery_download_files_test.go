// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package backuprecovery_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmRecoveryDownloadFilesDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmRecoveryDownloadFilesDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_recovery_download_files.recovery_download_files_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_recovery_download_files.recovery_download_files_instance", "recovery_download_files_id"),
				),
			},
		},
	})
}

func testAccCheckIbmRecoveryDownloadFilesDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_recovery_download_files" "recovery_download_files_instance" {
			id = "id"
			startOffset = 1
			length = 1
			fileType = "fileType"
			sourceName = "sourceName"
			startTime = "startTime"
			includeTenants = true
		}
	`)
}
