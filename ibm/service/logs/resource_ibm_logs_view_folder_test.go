// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package logs_test

import (
	"fmt"
	"testing"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/IBM/logs-go-sdk/logsv0"
	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/conns"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/flex"
)

func TestAccIbmLogsViewFolderBasic(t *testing.T) {
	var conf logsv0.ViewFolder
	name := fmt.Sprintf("tf_name_応答時間モニター_%d", acctest.RandIntRange(10, 100))
	// nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheckCloudLogs(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmLogsViewFolderDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmLogsViewFolderConfigBasic(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmLogsViewFolderExists("ibm_logs_view_folder.logs_view_folder_instance", conf),
					resource.TestCheckResourceAttr("ibm_logs_view_folder.logs_view_folder_instance", "name", name),
				),
			},
			// resource.TestStep{ #Todo @kavya498 enable update test once it is fixed in backend.
			// 	Config: testAccCheckIbmLogsViewFolderConfigBasic(nameUpdate),
			// 	Check: resource.ComposeAggregateTestCheckFunc(
			// 		resource.TestCheckResourceAttr("ibm_logs_view_folder.logs_view_folder_instance", "name", nameUpdate),
			// 	),
			// },
			resource.TestStep{
				ResourceName:      "ibm_logs_view_folder.logs_view_folder_instance",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmLogsViewFolderConfigBasic(name string) string {
	return fmt.Sprintf(`
	resource "ibm_logs_view_folder" "logs_view_folder_instance" {
		instance_id = "%s"
		region      = "%s"
		name        = "%s"
	  }
	`, acc.LogsInstanceId, acc.LogsInstanceRegion, name)
}

func testAccCheckIbmLogsViewFolderExists(n string, obj logsv0.ViewFolder) resource.TestCheckFunc {

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
		getViewFolderOptions := &logsv0.GetViewFolderOptions{}

		getViewFolderOptions.SetID(core.UUIDPtr(strfmt.UUID(resourceID[2])))

		viewFolder, _, err := logsClient.GetViewFolder(getViewFolderOptions)
		if err != nil {
			return err
		}

		obj = *viewFolder
		return nil
	}
}

func testAccCheckIbmLogsViewFolderDestroy(s *terraform.State) error {
	logsClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).LogsV0()
	if err != nil {
		return err
	}
	logsClient = getTestClientWithLogsInstanceEndpoint(logsClient)
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_logs_view_folder" {
			continue
		}
		resourceID, err := flex.IdParts(rs.Primary.ID)
		if err != nil {
			return err
		}

		getViewFolderOptions := &logsv0.GetViewFolderOptions{}

		getViewFolderOptions.SetID(core.UUIDPtr(strfmt.UUID(resourceID[2])))

		// Try to find the key
		_, response, err := logsClient.GetViewFolder(getViewFolderOptions)

		if err == nil {
			return fmt.Errorf("logs_view_folder still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for logs_view_folder (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
