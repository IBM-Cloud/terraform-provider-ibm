// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.96.1-5136e54a-20241108-203028
 */

package backuprecovery_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmBackupRecoveryManagerSreGetAlertsDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmBackupRecoveryManagerSreGetAlertsDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_manager_sre_get_alerts.backup_recovery_manager_sre_get_alerts_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_manager_sre_get_alerts.backup_recovery_manager_sre_get_alerts_instance", "alerts.#"),
				),
			},
		},
	})
}

func testAccCheckIbmBackupRecoveryManagerSreGetAlertsDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_backup_recovery_manager_sre_get_alerts" "backup_recovery_manager_sre_get_alerts_instance" {
			alertIds = [ "alertIds" ]
			alertTypes = [ 1 ]
			alertCategories = [ "kDisk" ]
			alertStates = [ "kResolved" ]
			alertSeverities = [ "kCritical" ]
			alertTypeBuckets = [ "kHardware" ]
			startTimeUsecs = 1
			endTimeUsecs = 1
			maxAlerts = 1
			propertyKey = "propertyKey"
			propertyValue = "propertyValue"
			alertName = "alertName"
			resolutionIds = [ 1 ]
			tenantIds = [ "tenantIds" ]
			allUnderHierarchy = true
			X-Scope-Identifier = "X-Scope-Identifier"
		}
	`)
}
