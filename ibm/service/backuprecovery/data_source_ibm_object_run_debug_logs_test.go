// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package backuprecovery_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmObjectRunDebugLogsDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmObjectRunDebugLogsDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_object_run_debug_logs.object_run_debug_logs_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_object_run_debug_logs.object_run_debug_logs_instance", "object_run_debug_logs_id"),
					resource.TestCheckResourceAttrSet("data.ibm_object_run_debug_logs.object_run_debug_logs_instance", "run_id"),
					resource.TestCheckResourceAttrSet("data.ibm_object_run_debug_logs.object_run_debug_logs_instance", "object_id"),
				),
			},
		},
	})
}

func testAccCheckIbmObjectRunDebugLogsDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_object_run_debug_logs" "object_run_debug_logs_instance" {
			id = "id"
			runId = "runId"
			objectId = "objectId"
		}
	`)
}
