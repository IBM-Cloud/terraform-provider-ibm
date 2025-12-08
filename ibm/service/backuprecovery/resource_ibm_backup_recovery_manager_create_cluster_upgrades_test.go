// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package backuprecovery_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmBackupRecoveryManagerCreateClusterUpgradesBasic(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmBackupRecoveryManagerCreateClusterUpgradesConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("resource.ibm_backup_recovery_manager_create_cluster_upgrades.backup_recovery_manager_create_cluster_upgrades_instance", "id"),
				),
				Destroy: false,
			},
		},
	})
}

func testAccCheckIbmBackupRecoveryManagerCreateClusterUpgradesConfigBasic() string {
	return fmt.Sprintf(`
		resource "ibm_backup_recovery_manager_create_cluster_upgrades" "backup_recovery_manager_create_cluster_upgrades_instance" {
			clusters {
				cluster_id = "3524800407225868"
				cluster_incarnation_id = "1758305184241232842"
			}
			package_url = "https://s3.us-east.cloud-object-storage.appdomain.cloud/7.2.15/cluster_artifacts/cohesity-7.2.15_release-20250721_6aa24701.tar.gz"
		}
	`)
}
