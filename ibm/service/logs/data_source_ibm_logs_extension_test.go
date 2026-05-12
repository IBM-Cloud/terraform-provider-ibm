// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package logs_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmLogsExtensionDataSourceBasic(t *testing.T) {
	// NOTE: This test requires a valid extension ID from your Cloud Logs instance
	// Extensions are IBM-provided resources. To get a valid extension ID:
	// 1. Run: make testacc TEST=./ibm/service/logs TESTARGS='-run=TestAccIbmLogsExtensionsDataSourceBasic'
	// 2. Check the terraform state or API response for available extension IDs
	// 3. Replace the hardcoded extension ID below with a real one

	// TODO: Replace with actual extension ID from your Cloud Logs instance
	extensionID := "IBMCloudant" // Replace with real UUID

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckCloudLogs(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmLogsExtensionDataSourceConfigBasic(extensionID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_logs_extension.logs_extension_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_extension.logs_extension_instance", "logs_extension_id"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_extension.logs_extension_instance", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_extension.logs_extension_instance", "revisions.#"),
					resource.TestCheckResourceAttrSet("data.ibm_logs_extension.logs_extension_instance", "revisions.0.items.#"),
				),
			},
		},
	})
}

func testAccCheckIbmLogsExtensionDataSourceConfigBasic(extensionID string) string {
	return fmt.Sprintf(`
		data "ibm_logs_extension" "logs_extension_instance" {
			instance_id       = "%s"
			region            = "%s"
			logs_extension_id = "%s"
		}
	`, acc.LogsInstanceId, acc.LogsInstanceRegion, extensionID)
}
