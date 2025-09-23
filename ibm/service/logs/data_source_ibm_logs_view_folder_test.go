// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package logs_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmLogsViewFolderDataSourceBasic(t *testing.T) {
	viewFolderName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckCloudLogs(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmLogsViewFolderDataSourceConfigBasic(viewFolderName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_logs_view_folder.logs_view_folder_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_view_folder.logs_view_folder_instance", "logs_view_folder_id"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_view_folder.logs_view_folder_instance", "name"),
				),
			},
		},
	})
}

func testAccCheckIbmLogsViewFolderDataSourceConfigBasic(viewFolderName string) string {
	return fmt.Sprintf(`
	resource "ibm_logs_view_folder" "logs_view_folder_instance" {
		instance_id = "%s"
		region      = "%s"
		name        = "%s"
	  }

	data "ibm_logs_view_folder" "logs_view_folder_instance" {
		instance_id = ibm_logs_view_folder.logs_view_folder_instance.instance_id
		region      = ibm_logs_view_folder.logs_view_folder_instance.region
		logs_view_folder_id = ibm_logs_view_folder.logs_view_folder_instance.view_folder_id
	}
	`, acc.LogsInstanceId, acc.LogsInstanceRegion, viewFolderName)
}
