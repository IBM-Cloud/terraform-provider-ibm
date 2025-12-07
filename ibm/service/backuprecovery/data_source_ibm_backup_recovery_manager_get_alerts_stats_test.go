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

func TestAccIbmBackupRecoveryManagerGetAlertsStatsDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmBackupRecoveryManagerGetAlertsStatsDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_manager_get_alerts_stats.backup_recovery_manager_get_alerts_stats_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_manager_get_alerts_stats.backup_recovery_manager_get_alerts_stats_instance", "start_time_usecs"),
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_manager_get_alerts_stats.backup_recovery_manager_get_alerts_stats_instance", "end_time_usecs"),
				),
			},
		},
	})
}

func testAccCheckIbmBackupRecoveryManagerGetAlertsStatsDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_backup_recovery_manager_get_alerts_stats" "backup_recovery_manager_get_alerts_stats_instance" {
			start_time_usecs = 1748739600000
			end_time_usecs = 1848739600000
		}
	`)
}
