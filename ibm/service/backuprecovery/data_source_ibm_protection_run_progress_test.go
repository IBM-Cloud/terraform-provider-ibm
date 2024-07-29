// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package backuprecovery_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmProtectionRunProgressDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmProtectionRunProgressDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_protection_run_progress.protection_run_progress_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_protection_run_progress.protection_run_progress_instance", "run_id"),
				),
			},
		},
	})
}

func testAccCheckIbmProtectionRunProgressDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_protection_run_progress" "protection_run_progress_instance" {
			runId = "runId"
			objects = [ 1 ]
			tenantIds = [ "tenantIds" ]
			includeTenants = true
			includeFinishedTasks = true
			startTimeUsecs = 1
			endTimeUsecs = 1
			maxTasksNum = 1
			excludeObjectDetails = true
			includeEventLogs = true
			maxLogLevel = 1
			runTaskPath = "runTaskPath"
			objectTaskPaths = [ "objectTaskPaths" ]
		}
	`)
}
