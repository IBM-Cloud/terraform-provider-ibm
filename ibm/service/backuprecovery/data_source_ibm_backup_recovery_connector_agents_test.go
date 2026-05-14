// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.113.1-d76630af-20260320-135953
 */

package backuprecovery_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIbmBackupRecoveryConnectorAgentsDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmBackupRecoveryConnectorAgentsDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_connector_agents.backup_recovery_connector_agents_instance", "connector_agents.#"),
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_connector_agents.backup_recovery_connector_agents_instance", "connector_agents.0.connection_id"),
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_connector_agents.backup_recovery_connector_agents_instance", "connector_agents.0.connection_name"),
				),
			},
		},
	})
}

func testAccCheckIbmBackupRecoveryConnectorAgentsDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_backup_recovery_connector_agents" "backup_recovery_connector_agents_instance" {
			x_ibm_tenant_id = "4ugtn5idq2/"
			tenant_id = "4ugtn5idq2/"
		}
	`)
}
