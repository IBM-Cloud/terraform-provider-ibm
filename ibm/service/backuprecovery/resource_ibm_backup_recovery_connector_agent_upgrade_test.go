// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package backuprecovery_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmBackupRecoveryConnectorAgentUpgradeBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmBackupRecoveryConnectorAgentUpgradeConfigBasic(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("ibm_backup_recovery_connector_agent_upgrade.backup_recovery_connector_agent_upgrade_instance", "upgrade_id"),
					resource.TestCheckResourceAttrSet("ibm_backup_recovery_connector_agent_upgrade.backup_recovery_connector_agent_upgrade_instance", "status"),
				),
			},
		},
	})
}

func testAccCheckIbmBackupRecoveryConnectorAgentUpgradeConfigBasic() string {
	return fmt.Sprintf(`
		resource "ibm_backup_recovery_connector_agent_upgrade" "backup_recovery_connector_agent_upgrade_instance" {
			connection_ids = [ 1 ]
			retry_forever  = false
		}
	`)
}
