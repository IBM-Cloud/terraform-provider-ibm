// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package backuprecovery_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmBackupRecoveryConnectorAccessTokenBasic(t *testing.T) {
	username := "admin"
	password := "newPassword7"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Destroy: false,
				Config:  testAccCheckIbmBackupRecoveryConnectorAccessTokenConfigBasic(username, password),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_backup_recovery_connector_access_token.backup_recovery_connector_access_token_instance", "username", username),
					resource.TestCheckResourceAttr("ibm_backup_recovery_connector_access_token.backup_recovery_connector_access_token_instance", "password", password),
				),
			},
		},
	})
}

func testAccCheckIbmBackupRecoveryConnectorAccessTokenConfigBasic(username, password string) string {
	return fmt.Sprintf(`
		resource "ibm_backup_recovery_connector_access_token" "backup_recovery_connector_access_token_instance" {
			username = "%s"
			password = "%s"
		}
	`, username, password)
}
