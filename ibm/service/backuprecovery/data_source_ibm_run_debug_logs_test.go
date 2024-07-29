// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package backuprecovery_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmRunDebugLogsDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmRunDebugLogsDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_run_debug_logs.run_debug_logs_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_run_debug_logs.run_debug_logs_instance", "run_debug_logs_id"),
					resource.TestCheckResourceAttrSet("data.ibm_run_debug_logs.run_debug_logs_instance", "run_id"),
				),
			},
		},
	})
}

func testAccCheckIbmRunDebugLogsDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_run_debug_logs" "run_debug_logs_instance" {
			id = "id"
			runId = "runId"
			objectId = "objectId"
		}
	`)
}
