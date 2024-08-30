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

func TestAccIbmDownloadIndexedFilesDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmDownloadIndexedFilesDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_download_indexed_files.download_indexed_files_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_download_indexed_files.download_indexed_files_instance", "snapshots_id"),
				),
			},
		},
	})
}

func testAccCheckIbmDownloadIndexedFilesDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_download_indexed_files" "download_indexed_files_instance" {
			snapshotsId = "snapshotsId"
			filePath = "filePath"
			nvramFile = true
			retryAttempt = 1
			startOffset = 1
			length = 1
		}
	`)
}
