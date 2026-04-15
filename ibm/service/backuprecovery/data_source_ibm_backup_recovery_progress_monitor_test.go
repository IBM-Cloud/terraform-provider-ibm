// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.113.1-d76630af-20260320-135953
 */

package backuprecovery_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmBackupRecoveryProgressMonitorDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmBackupRecoveryProgressMonitorDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_progress_monitor.backup_recovery_progress_monitor_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_progress_monitor.backup_recovery_progress_monitor_instance", "x_ibm_tenant_id"),
				),
			},
		},
	})
}

func testAccCheckIbmBackupRecoveryProgressMonitorDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_backup_recovery_progress_monitor" "backup_recovery_progress_monitor_instance" {
			X-IBM-Tenant-Id = "tenantId"
			attributeVec = [ { key="key", value={ data={ oneof_data={  } }, type=1 } } ]
			endTimeSecs = 1
			excludeSubTasks = true
			fetchLogsMaxLevel = 1
			includeEventLogs = true
			includeFinishedTasks = true
			maxTasks = 1
			startTimeSecs = 1
			taskPathVec = [ "new_taskPathVec" ]
			taskPathVec = [ "taskPathVec" ]
			includeFinishedTasks = true
			startTimeSecs = 1
			endTimeSecs = 1
			maxTasks = 1
			excludeSubTasks = true
		}
	`)
}
