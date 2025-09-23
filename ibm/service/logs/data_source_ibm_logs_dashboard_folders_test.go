// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package logs_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/logs-go-sdk/logsv0"
	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/service/logs"
	. "github.com/Mavrickk3/terraform-provider-ibm/ibm/unittest"
	"github.com/stretchr/testify/assert"
)

func TestAccIbmLogsDashboardFoldersDataSourceBasic(t *testing.T) {
	dashboardFolderName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckCloudLogs(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmLogsDashboardFoldersDataSourceConfigBasic(dashboardFolderName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_logs_dashboard_folders.logs_dashboard_folders_instance", "id"),
				),
			},
		},
	})
}

func testAccCheckIbmLogsDashboardFoldersDataSourceConfigBasic(dashboardFolderName string) string {
	return fmt.Sprintf(`
		resource "ibm_logs_dashboard_folder" "logs_dashboard_folder_instance" {
			instance_id = "%s"
			region      = "%s"
			name        = "%s"
		}

		data "ibm_logs_dashboard_folders" "logs_dashboard_folders_instance" {
			instance_id = ibm_logs_dashboard_folder.logs_dashboard_folder_instance.instance_id
			region      = ibm_logs_dashboard_folder.logs_dashboard_folder_instance.region
			depends_on = [
				ibm_logs_dashboard_folder.logs_dashboard_folder_instance
			]
		}
	`, acc.LogsInstanceId, acc.LogsInstanceRegion, dashboardFolderName)
}

func TestDataSourceIbmLogsDashboardFoldersDashboardFolderToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["id"] = "3dc02998-0b50-4ea8-b68a-4779d716fa1f"
		model["name"] = "My Folder"
		model["parent_id"] = "3dc02998-0b50-4ea8-b68a-4779d716fa1f"

		assert.Equal(t, result, model)
	}

	model := new(logsv0.DashboardFolder)
	model.ID = CreateMockUUID("3dc02998-0b50-4ea8-b68a-4779d716fa1f")
	model.Name = core.StringPtr("My Folder")
	model.ParentID = CreateMockUUID("3dc02998-0b50-4ea8-b68a-4779d716fa1f")

	result, err := logs.DataSourceIbmLogsDashboardFoldersDashboardFolderToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
