// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.94.0-fa797aec-20240814-142622
 */

package backuprecovery_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmBackupRecoveryConnectorLogsDataSourceBasic(t *testing.T) {
	username := "admin"
	password := "newPassword7"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmBackupRecoveryConnectorLogsDataSourceConfigBasic(username, password),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_connector_logs.backup_recovery_connector_logs_instance", "connector_logs.#"),
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_connector_logs.backup_recovery_connector_logs_instance", "id"),
				),
			},
		},
	})
}

func testAccCheckIbmBackupRecoveryConnectorLogsDataSourceConfigBasic(username, password string) string {
	return fmt.Sprintf(`

	resource "ibm_backup_recovery_connector_access_token" "backup_recovery_connector_access_token_instance" {
		username = "%s"
		password = "%s"
	}
		data "ibm_backup_recovery_connector_logs" "backup_recovery_connector_logs_instance" {
			access_token = resource.ibm_backup_recovery_connector_access_token.backup_recovery_connector_access_token_instance.access_token
		}
	`, username, password)
}
