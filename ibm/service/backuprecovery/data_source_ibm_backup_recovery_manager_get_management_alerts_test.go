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

func TestAccIbmBackupRecoveryManagerGetManagementAlertsDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmBackupRecoveryManagerGetManagementAlertsDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_manager_get_management_alerts.backup_recovery_manager_get_management_alerts_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_manager_get_management_alerts.backup_recovery_manager_get_management_alerts_instance", "alerts_list.#"),
				),
			},
		},
	})
}

func testAccCheckIbmBackupRecoveryManagerGetManagementAlertsDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_backup_recovery_manager_get_management_alerts" "backup_recovery_manager_get_management_alerts_instance" {

		}
	`)
}
