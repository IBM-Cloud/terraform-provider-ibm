// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package backuprecovery_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmProtectionGroupsDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmProtectionGroupsDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_protection_groups.protection_groups_instance", "id"),
				),
			},
		},
	})
}

func testAccCheckIbmProtectionGroupsDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_protection_groups" "protection_groups_instance" {
			requestInitiatorType = "UIUser"
			ids = [ "ids" ]
			names = [ "names" ]
			policyIds = [ "policyIds" ]
			storageDomainId = 1
			includeGroupsWithDatalockOnly = true
			environments = [ "kVMware" ]
			isActive = true
			isDeleted = true
			isPaused = true
			lastRunLocalBackupStatus = [ "Accepted" ]
			lastRunReplicationStatus = [ "Accepted" ]
			lastRunArchivalStatus = [ "Accepted" ]
			lastRunCloudSpinStatus = [ "Accepted" ]
			lastRunAnyStatus = [ "Accepted" ]
			isLastRunSlaViolated = true
			tenantIds = [ "tenantIds" ]
			includeTenants = true
			includeLastRunInfo = true
			pruneExcludedSourceIds = true
			pruneSourceIds = true
			useCachedData = true
		}
	`)
}
