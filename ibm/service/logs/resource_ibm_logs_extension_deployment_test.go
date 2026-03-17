// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package logs_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/logs-go-sdk/logsv0"
)

func TestAccIbmLogsExtensionDeploymentBasic(t *testing.T) {
	// This test uses the IBMCloudant extension which is available in all Cloud Logs instances
	// The test fetches the extension first to get available versions and item IDs dynamically
	extensionID := "IBMCloudant"
	applications := "test-app"
	subsystems := "test-subsystem"
	updateApplications := "update-app"
	updateSubsystems := "update-subsystem"

	var conf logsv0.ExtensionDeployment

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheckCloudLogs(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmLogsExtensionDeploymentDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmLogsExtensionDeploymentConfigBasic(extensionID),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmLogsExtensionDeploymentExists("ibm_logs_extension_deployment.logs_extension_deployment_instance", conf),
					resource.TestCheckResourceAttr("ibm_logs_extension_deployment.logs_extension_deployment_instance", "logs_extension_id", extensionID),
					resource.TestCheckResourceAttrSet("ibm_logs_extension_deployment.logs_extension_deployment_instance", "version"),
					resource.TestCheckResourceAttrSet("ibm_logs_extension_deployment.logs_extension_deployment_instance", "item_ids.#"),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmLogsExtensionDeploymentConfigUpdate(extensionID, applications, subsystems),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_logs_extension_deployment.logs_extension_deployment_instance", "logs_extension_id", extensionID),
					resource.TestCheckResourceAttrSet("ibm_logs_extension_deployment.logs_extension_deployment_instance", "version"),
					resource.TestCheckResourceAttrSet("ibm_logs_extension_deployment.logs_extension_deployment_instance", "item_ids.#"),
					resource.TestCheckResourceAttr("ibm_logs_extension_deployment.logs_extension_deployment_instance", "applications.#", "1"),
					resource.TestCheckResourceAttr("ibm_logs_extension_deployment.logs_extension_deployment_instance", "applications.0", applications),
					resource.TestCheckResourceAttr("ibm_logs_extension_deployment.logs_extension_deployment_instance", "subsystems.#", "1"),
					resource.TestCheckResourceAttr("ibm_logs_extension_deployment.logs_extension_deployment_instance", "subsystems.0", subsystems),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmLogsExtensionDeploymentConfigUpdate(extensionID, updateApplications, updateSubsystems),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_logs_extension_deployment.logs_extension_deployment_instance", "logs_extension_id", extensionID),
					resource.TestCheckResourceAttrSet("ibm_logs_extension_deployment.logs_extension_deployment_instance", "version"),
					resource.TestCheckResourceAttr("ibm_logs_extension_deployment.logs_extension_deployment_instance", "item_ids.#", "2"),
					resource.TestCheckResourceAttr("ibm_logs_extension_deployment.logs_extension_deployment_instance", "applications.#", "1"),
					resource.TestCheckResourceAttr("ibm_logs_extension_deployment.logs_extension_deployment_instance", "applications.0", updateApplications),
					resource.TestCheckResourceAttr("ibm_logs_extension_deployment.logs_extension_deployment_instance", "subsystems.#", "1"),
					resource.TestCheckResourceAttr("ibm_logs_extension_deployment.logs_extension_deployment_instance", "subsystems.0", updateSubsystems),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_logs_extension_deployment.logs_extension_deployment_instance",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmLogsExtensionDeploymentConfigBasic(extensionID string) string {
	return fmt.Sprintf(`
		data "ibm_logs_extension" "extension" {
			instance_id       = "%[1]s"
			region            = "%[2]s"
			logs_extension_id = "%[3]s"
		}

		resource "ibm_logs_extension_deployment" "logs_extension_deployment_instance" {
			instance_id       = "%[1]s"
			region            = "%[2]s"
			logs_extension_id = "%[3]s"
			version           = data.ibm_logs_extension.extension.revisions[0].version
			item_ids          = [for item in data.ibm_logs_extension.extension.revisions[0].items : item.id]
		}
	`, acc.LogsInstanceId, acc.LogsInstanceRegion, extensionID)
}

func testAccCheckIbmLogsExtensionDeploymentConfigUpdate(extensionID, applications, subsystems string) string {
	return fmt.Sprintf(`
		data "ibm_logs_extension" "extension" {
			instance_id       = "%[1]s"
			region            = "%[2]s"
			logs_extension_id = "%[3]s"
		}

		resource "ibm_logs_extension_deployment" "logs_extension_deployment_instance" {
			instance_id       = "%[1]s"
			region            = "%[2]s"
			logs_extension_id = "%[3]s"
			version           = data.ibm_logs_extension.extension.revisions[0].version
			item_ids          = slice([for item in data.ibm_logs_extension.extension.revisions[0].items : item.id], 0, 2)
			applications      = ["%[4]s"]
			subsystems        = ["%[5]s"]
		}
	`, acc.LogsInstanceId, acc.LogsInstanceRegion, extensionID, applications, subsystems)
}

func testAccCheckIbmLogsExtensionDeploymentExists(n string, obj logsv0.ExtensionDeployment) resource.TestCheckFunc {

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

		extensionID := resourceID[2]

		getExtensionDeploymentOptions := &logsv0.GetExtensionDeploymentOptions{}
		getExtensionDeploymentOptions.SetID(extensionID)

		extensionDeployment, _, err := logsClient.GetExtensionDeployment(getExtensionDeploymentOptions)
		if err != nil {
			return err
		}

		obj = *extensionDeployment
		return nil
	}
}

func testAccCheckIbmLogsExtensionDeploymentDestroy(s *terraform.State) error {
	logsClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).LogsV0()
	if err != nil {
		return err
	}
	logsClient = getTestClientWithLogsInstanceEndpoint(logsClient)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_logs_extension_deployment" {
			continue
		}

		resourceID, err := flex.IdParts(rs.Primary.ID)
		if err != nil {
			return err
		}

		extensionID := resourceID[2]

		getExtensionDeploymentOptions := &logsv0.GetExtensionDeploymentOptions{}
		getExtensionDeploymentOptions.SetID(extensionID)

		// Try to find the key
		_, response, err := logsClient.GetExtensionDeployment(getExtensionDeploymentOptions)

		if err == nil {
			return fmt.Errorf("Extension deployment still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for Extension deployment (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
