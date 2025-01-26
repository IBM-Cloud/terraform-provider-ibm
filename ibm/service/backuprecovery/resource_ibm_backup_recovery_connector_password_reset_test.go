// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package backuprecovery_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.ibm.com/BackupAndRecovery/ibm-backup-recovery-sdk-go/backuprecoveryv1"
)

func TestAccIbmBackupRecoveryConnectorPasswordResetBasic(t *testing.T) {
	var conf backuprecoveryv1.TokenResponse
	xIbmTenantID := fmt.Sprintf("tf_x_ibm_tenant_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmBackupRecoveryConnectorAccessTokenDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmBackupRecoveryConnectorAccessTokenConfigBasic(xIbmTenantID),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmBackupRecoveryConnectorAccessTokenExists("ibm_backup_recovery_connector_access_token.backup_recovery_connector_access_token_instance", conf),
					resource.TestCheckResourceAttr("ibm_backup_recovery_connector_access_token.backup_recovery_connector_access_token_instance", "x_ibm_tenant_id", xIbmTenantID),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_backup_recovery_connector_access_token.backup_recovery_connector_access_token",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccIbmBackupRecoveryConnectorPasswordResetBasic(xIbmTenantID string) string {
	return fmt.Sprintf(`
		resource "ibm_backup_recovery_connector_access_token" "backup_recovery_connector_access_token_instance" {
			x_ibm_tenant_id = "%s"
		}
	`, xIbmTenantID)
}

func testAccIbmBackupRecoveryConnectorPasswordResetExists(n string, obj backuprecoveryv1.TokenResponse) resource.TestCheckFunc {

	return nil
	// return func(s *terraform.State) error {
	// 	_, ok := s.RootModule().Resources[n]
	// 	if !ok {
	// 		return fmt.Errorf("Not found: %s", n)
	// 	}

	// 	backupRecoveryClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).BackupRecoveryV1()
	// 	if err != nil {
	// 		return err
	// 	}

	// 	createAccessTokenOptions := &backuprecoveryv1.CreateAccessTokenOptions{}

	// 	tokenResponse, _, err := backupRecoveryClient.CreateAccessToken(createAccessTokenOptions)
	// 	if err != nil {
	// 		return err
	// 	}

	// 	obj = *tokenResponse
	// 	return nil
	// }
}

func testAccIbmBackupRecoveryConnectorPasswordResetDestroty(s *terraform.State) error {
	// backupRecoveryClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).BackupRecoveryV1()
	// if err != nil {
	// 	return err
	// }
	// for _, rs := range s.RootModule().Resources {
	// 	if rs.Type != "ibm_backup_recovery_connector_access_token" {
	// 		continue
	// 	}

	// 	createAccessTokenOptions := &backuprecoveryv1.CreateAccessTokenOptions{}

	// 	// Try to find the key
	// 	_, response, err := backupRecoveryClient.CreateAccessToken(createAccessTokenOptions)

	// 	if err == nil {
	// 		return fmt.Errorf("backup_recovery_connector_access_token still exists: %s", rs.Primary.ID)
	// 	} else if response.StatusCode != 404 {
	// 		return fmt.Errorf("Error checking for backup_recovery_connector_access_token (%s) has been destroyed: %s", rs.Primary.ID, err)
	// 	}
	// }

	return nil
}
