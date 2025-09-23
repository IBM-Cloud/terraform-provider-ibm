// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package logs_test

import (
	"fmt"
	"testing"

	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/logs-go-sdk/logsv0"
	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/conns"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/flex"
)

func TestAccIbmLogsDashboardFolderBasic(t *testing.T) {
	var conf logsv0.DashboardFolder
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheckCloudLogs(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmLogsDashboardFolderDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmLogsDashboardFolderConfigBasic(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmLogsDashboardFolderExists("ibm_logs_dashboard_folder.logs_dashboard_folder_instance", conf),
					resource.TestCheckResourceAttr("ibm_logs_dashboard_folder.logs_dashboard_folder_instance", "name", name),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmLogsDashboardFolderConfigBasic(nameUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_logs_dashboard_folder.logs_dashboard_folder_instance", "name", nameUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_logs_dashboard_folder.logs_dashboard_folder_instance",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmLogsDashboardFolderConfigBasic(name string) string {
	return fmt.Sprintf(`
		resource "ibm_logs_dashboard_folder" "logs_dashboard_folder_instance" {
			instance_id = "%s"
			region      = "%s"
			name = "%s"
		}
	`, acc.LogsInstanceId, acc.LogsInstanceRegion, name)
}

func testAccCheckIbmLogsDashboardFolderExists(n string, obj logsv0.DashboardFolder) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		logsClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).LogsV0()
		if err != nil {
			return err
		}
		logsClient = getTestClientWithLogsInstanceEndpoint(logsClient)
		resourceID, err := flex.IdParts(rs.Primary.ID)
		if err != nil {
			return err
		}

		listDashboardFoldersOptions := &logsv0.ListDashboardFoldersOptions{}

		dashboardFolder, _, err := logsClient.ListDashboardFolders(listDashboardFoldersOptions)
		if err != nil {
			return err
		}
		for _, folder := range dashboardFolder.Folders {
			if folder.ID == core.UUIDPtr(strfmt.UUID(resourceID[2])) {
				obj = folder
				return nil
			}
		}
		return nil
	}
}

func testAccCheckIbmLogsDashboardFolderDestroy(s *terraform.State) error {
	logsClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).LogsV0()
	if err != nil {
		return err
	}
	logsClient = getTestClientWithLogsInstanceEndpoint(logsClient)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_logs_dashboard_folder" {
			continue
		}
		resourceID, err := flex.IdParts(rs.Primary.ID)
		if err != nil {
			return err
		}
		listDashboardFoldersOptions := &logsv0.ListDashboardFoldersOptions{}

		// Try to find the key
		dashboardFolder, _, _ := logsClient.ListDashboardFolders(listDashboardFoldersOptions)

		for _, folder := range dashboardFolder.Folders {
			if folder.ID == core.UUIDPtr(strfmt.UUID(resourceID[2])) {
				return fmt.Errorf("logs_dashboard_folder still exists: %s", rs.Primary.ID)
			}
		}
	}

	return nil
}
