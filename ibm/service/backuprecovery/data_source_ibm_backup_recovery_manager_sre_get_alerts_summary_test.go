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

func TestAccIbmBackupRecoveryManagerSreGetAlertsSummaryDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmBackupRecoveryManagerSreGetAlertsSummaryDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_manager_sre_get_alerts_summary.backup_recovery_manager_sre_get_alerts_summary_instance", "id"),
				),
			},
		},
	})
}

func testAccCheckIbmBackupRecoveryManagerSreGetAlertsSummaryDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_backup_recovery_manager_sre_get_alerts_summary" "backup_recovery_manager_sre_get_alerts_summary_instance" {
			startTimeUsecs = 1
			endTimeUsecs = 1
			includeTenants = true
			tenantIds = [ "tenantIds" ]
			statesList = [ "kResolved" ]
			X-Scope-Identifier = "X-Scope-Identifier"
		}
	`)
}
