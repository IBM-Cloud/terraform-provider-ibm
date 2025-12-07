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

func TestAccIbmBackupRecoveryManagerGetUpgradesInfoDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmBackupRecoveryManagerGetUpgradesInfoDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_manager_get_upgrades_info.backup_recovery_manager_get_upgrades_info_instance", "id"),
				),
			},
		},
	})
}

func testAccCheckIbmBackupRecoveryManagerGetUpgradesInfoDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_backup_recovery_manager_get_upgrades_info" "backup_recovery_manager_get_upgrades_info_instance" {
			cluster_identifiers = [ "3524800407225868","8305184241232842","7463743295903869","1589079364046703","90858563991288","4532338433036076" ]
		}
	`)
}
