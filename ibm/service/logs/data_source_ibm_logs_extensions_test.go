// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package logs_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmLogsExtensionsDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckCloudLogs(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmLogsExtensionsDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_logs_extensions.logs_extensions_instance", "id"),
				),
			},
		},
	})
}

func testAccCheckIbmLogsExtensionsDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_logs_extensions" "logs_extensions_instance" {
			instance_id = "%s"
			region = "%s"
			deployed = true
		}
	`, acc.LogsInstanceId, acc.LogsInstanceRegion)
}
