// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package backuprecovery_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmProtectionGroupRunsDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmProtectionGroupRunsDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_protection_group_runs.protection_group_runs_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_protection_group_runs.protection_group_runs_instance", "protection_group_runs_id"),
				),
			},
		},
	})
}

func testAccCheckIbmProtectionGroupRunsDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_protection_group_runs" "protection_group_runs_instance" {
			id = "id"
			requestInitiatorType = "UIUser"
			runId = "runId"
			startTimeUsecs = 1
			endTimeUsecs = 1
			tenantIds = [ "tenantIds" ]
			includeTenants = true
			runTypes = [ "kAll" ]
			includeObjectDetails = true
			localBackupRunStatus = [ "Accepted" ]
			replicationRunStatus = [ "Accepted" ]
			archivalRunStatus = [ "Accepted" ]
			cloudSpinRunStatus = [ "Accepted" ]
			numRuns = 1
			excludeNonRestorableRuns = true
			runTags = [ "runTags" ]
			useCachedData = true
			filterByEndTime = true
			snapshotTargetTypes = [ "Local" ]
			onlyReturnSuccessfulCopyRun = true
			filterByCopyTaskEndTime = true
		}
	`)
}
