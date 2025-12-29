// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package backuprecovery_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmBackupRecoveryManagerCancelClusterUpgradesBasic(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config:  testAccCheckIbmBackupRecoveryManagerCancelClusterUpgradesConfigBasic(),
				Destroy: false,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("resource.ibm_backup_recovery_manager_cancel_cluster_upgrades.backup_recovery_manager_cancel_cluster_upgrades_instance", "id"),
				),
			},
		},
	})
}

func testAccCheckIbmBackupRecoveryManagerCancelClusterUpgradesConfigBasic() string {
	return fmt.Sprintf(`
		resource "ibm_backup_recovery_manager_cancel_cluster_upgrades" "backup_recovery_manager_cancel_cluster_upgrades_instance" {
			cluster_identifiers = ["3524800407225868"]
		}
	`)
}
