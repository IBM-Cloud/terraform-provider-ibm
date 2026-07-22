// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package backuprecovery_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmBackupRecoveryConnectorAgentRegistrationBasic(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmBackupRecoveryConnectorAgentRegistrationConfigBasic(),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("ibm_backup_recovery_connector_agent_registration.ibm_backup_recovery_connector_agent_registration", "registration_status"),
				),
				Destroy: false,
			},
		},
	})
}

func testAccCheckIbmBackupRecoveryConnectorAgentRegistrationConfigBasic() string {
	return fmt.Sprintf(`
	resource "ibm_backup_recovery_connector_agent_registration" "ibm_backup_recovery_connector_agent_registration"{
		registration_token = ""
		connection_name = "terra-conn-register-Connector-2"
		join_existing_connection = false
	}
	`)
}
