// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package logs_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/logs"
	. "github.com/IBM-Cloud/terraform-provider-ibm/ibm/unittest"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/logs-go-sdk/logsv0"
	"github.com/stretchr/testify/assert"
)

func TestAccIbmLogsViewFoldersDataSourceBasic(t *testing.T) {
	viewFolderName := fmt.Sprintf("TF_LOG_アクセスログ_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckCloudLogs(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmLogsViewFoldersDataSourceConfigBasic(viewFolderName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_logs_view_folders.logs_view_folders_instance", "id"),
				),
			},
		},
	})
}

func testAccCheckIbmLogsViewFoldersDataSourceConfigBasic(viewFolderName string) string {
	return fmt.Sprintf(`
	resource "ibm_logs_view_folder" "logs_view_folder_instance" {
		instance_id = "%s"
		region      = "%s"
		name        = "%s"
	  }

	data "ibm_logs_view_folders" "logs_view_folders_instance" {
		depends_on = [
			ibm_logs_view_folder.logs_view_folder_instance
		]
		instance_id = ibm_logs_view_folder.logs_view_folder_instance.instance_id
		region      = ibm_logs_view_folder.logs_view_folder_instance.region
	}
	`, acc.LogsInstanceId, acc.LogsInstanceRegion, viewFolderName)
}

func TestDataSourceIbmLogsViewFoldersViewFolderToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["id"] = "3dc02998-0b50-4ea8-b68a-4779d716fa1f"
		model["name"] = "My Folder"

		assert.Equal(t, result, model)
	}

	model := new(logsv0.ViewFolder)
	model.ID = CreateMockUUID("3dc02998-0b50-4ea8-b68a-4779d716fa1f")
	model.Name = core.StringPtr("My Folder")

	result, err := logs.DataSourceIbmLogsViewFoldersViewFolderToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
