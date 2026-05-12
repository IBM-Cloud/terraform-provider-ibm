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

func TestAccIbmBackupRecoveryConnectorAgentConfigDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmBackupRecoveryConnectorAgentConfigDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_connector_agent_config.backup_recovery_connector_agent_config_instance", "id"),
				),
			},
		},
	})
}

func testAccCheckIbmBackupRecoveryConnectorAgentConfigDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_backup_recovery_connector_agent_config" "backup_recovery_connector_agent_config_instance" {
		}
	`)
}
