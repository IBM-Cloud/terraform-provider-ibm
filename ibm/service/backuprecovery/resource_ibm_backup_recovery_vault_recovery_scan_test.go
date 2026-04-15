// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package backuprecovery_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/backuprecovery"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/ibm-backup-recovery-sdk-go/backuprecoveryv1"
	"github.com/stretchr/testify/assert"
	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmBackupRecoveryVaultRecoveryScanBasic(t *testing.T) {
	var conf backuprecoveryv1.RecoveryScan
	xIbmTenantID := fmt.Sprintf("tf_x_ibm_tenant_id_%d", acctest.RandIntRange(10, 100))
	cloudType := "ibm"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmBackupRecoveryVaultRecoveryScanDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmBackupRecoveryVaultRecoveryScanConfigBasic(xIbmTenantID, cloudType),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmBackupRecoveryVaultRecoveryScanExists("ibm_backup_recovery_vault_recovery_scan.backup_recovery_vault_recovery_scan_instance", conf),
					resource.TestCheckResourceAttr("ibm_backup_recovery_vault_recovery_scan.backup_recovery_vault_recovery_scan_instance", "x_ibm_tenant_id", xIbmTenantID),
					resource.TestCheckResourceAttr("ibm_backup_recovery_vault_recovery_scan.backup_recovery_vault_recovery_scan_instance", "cloud_type", cloudType),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_backup_recovery_vault_recovery_scan.backup_recovery_vault_recovery_scan_instance",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmBackupRecoveryVaultRecoveryScanConfigBasic(xIbmTenantID string, cloudType string) string {
	return fmt.Sprintf(`
		resource "ibm_backup_recovery_vault_recovery_scan" "backup_recovery_vault_recovery_scan_instance" {
			x_ibm_tenant_id = "%s"
			cloud_type = "%s"
		}
	`, xIbmTenantID, cloudType)
}

func testAccCheckIbmBackupRecoveryVaultRecoveryScanExists(n string, obj backuprecoveryv1.RecoveryScan) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		_, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		backupRecoveryClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).BackupRecoveryV1()
		if err != nil {
			return err
		}

		getBatchVaultRecoveryScanStatusOptions := &backuprecoveryv1.GetBatchVaultRecoveryScanStatusOptions{}


		recoveryScan, _, err := backupRecoveryClient.GetBatchVaultRecoveryScanStatus(getBatchVaultRecoveryScanStatusOptions)
		if err != nil {
			return err
		}

		obj = *recoveryScan
		return nil
	}
}

func testAccCheckIbmBackupRecoveryVaultRecoveryScanDestroy(s *terraform.State) error {
	backupRecoveryClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).BackupRecoveryV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_backup_recovery_vault_recovery_scan" {
			continue
		}

		getBatchVaultRecoveryScanStatusOptions := &backuprecoveryv1.GetBatchVaultRecoveryScanStatusOptions{}


		// Try to find the key
		_, response, err := backupRecoveryClient.GetBatchVaultRecoveryScanStatus(getBatchVaultRecoveryScanStatusOptions)

		if err == nil {
			return fmt.Errorf("backup_recovery_vault_recovery_scan still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for backup_recovery_vault_recovery_scan (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

func TestResourceIbmBackupRecoveryVaultRecoveryScanRecoveryScanRequestParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["vault_id"] = int(26)

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.RecoveryScanRequestParams)
	model.VaultID = core.Int64Ptr(int64(26))

	result, err := backuprecovery.ResourceIbmBackupRecoveryVaultRecoveryScanRecoveryScanRequestParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBackupRecoveryVaultRecoveryScanMapToRecoveryScanRequestParams(t *testing.T) {
	checkResult := func(result *backuprecoveryv1.RecoveryScanRequestParams) {
		model := new(backuprecoveryv1.RecoveryScanRequestParams)
		model.VaultID = core.Int64Ptr(int64(26))

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["vault_id"] = int(26)

	result, err := backuprecovery.ResourceIbmBackupRecoveryVaultRecoveryScanMapToRecoveryScanRequestParams(model)
	assert.Nil(t, err)
	checkResult(result)
}
