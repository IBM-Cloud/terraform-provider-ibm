// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.94.0-fa797aec-20240814-142622
 */

package backuprecovery_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmBaasProtectionGroupRunDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmBaasProtectionGroupRunDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_baas_protection_group_run.baas_protection_group_run_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_protection_group_run.baas_protection_group_run_instance", "baas_protection_group_run_id"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_protection_group_run.baas_protection_group_run_instance", "tenant_id"),
				),
			},
		},
	})
}

func testAccCheckIbmBaasProtectionGroupRunDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_baas_protection_group_run" "baas_protection_group_run_instance" {
			id = "id"
			tenantId = 1
			requestInitiatorType = "UIUser"
			runId = "runId"
			startTimeUsecs = 1
			endTimeUsecs = 1
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
