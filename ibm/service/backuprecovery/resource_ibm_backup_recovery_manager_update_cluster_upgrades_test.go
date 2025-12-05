// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package backuprecovery_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmBackupRecoveryManagerUpdateClusterUpgradesBasic(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmBackupRecoveryManagerUpdateClusterUpgradesConfigBasic(),
				Check:  resource.ComposeAggregateTestCheckFunc(),
			},
			resource.TestStep{
				ResourceName:      "ibm_backup_recovery_manager_update_cluster_upgrades.backup_recovery_manager_update_cluster_upgrades",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmBackupRecoveryManagerUpdateClusterUpgradesConfigBasic() string {
	return fmt.Sprintf(`
		resource "ibm_backup_recovery_manager_update_cluster_upgrades" "backup_recovery_manager_update_cluster_upgrades_instance" {
		}
	`)
}
