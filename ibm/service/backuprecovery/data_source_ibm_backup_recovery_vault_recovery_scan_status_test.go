// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.113.1-d76630af-20260320-135953
*/

package backuprecovery_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/backuprecovery"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/ibm-backup-recovery-sdk-go/backuprecoveryv1"
	"github.com/stretchr/testify/assert"
	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmBackupRecoveryVaultRecoveryScanStatusDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmBackupRecoveryVaultRecoveryScanStatusDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_vault_recovery_scan_status.backup_recovery_vault_recovery_scan_status_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_vault_recovery_scan_status.backup_recovery_vault_recovery_scan_status_instance", "x_ibm_tenant_id"),
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_vault_recovery_scan_status.backup_recovery_vault_recovery_scan_status_instance", "cloud_type"),
				),
			},
		},
	})
}

func testAccCheckIbmBackupRecoveryVaultRecoveryScanStatusDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_backup_recovery_vault_recovery_scan_status" "backup_recovery_vault_recovery_scan_status_instance" {
			X-IBM-Tenant-Id = "tenantId"
			cloudType = "ibm"
			vaultIds = [ 1 ]
		}
	`)
}


func TestDataSourceIbmBackupRecoveryVaultRecoveryScanStatusBatchVaultRecoveryScanStatusToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		recoveryScanStatusModel := make(map[string]interface{})
		recoveryScanStatusModel["end_time_usecs"] = int(26)
		recoveryScanStatusModel["error_message"] = "testString"
		recoveryScanStatusModel["start_time_usecs"] = int(26)
		recoveryScanStatusModel["status"] = "Accepted"
		recoveryScanStatusModel["uid"] = "testString"

		model := make(map[string]interface{})
		model["status"] = []map[string]interface{}{recoveryScanStatusModel}
		model["vault_id"] = int(26)

		assert.Equal(t, result, model)
	}

	recoveryScanStatusModel := new(backuprecoveryv1.RecoveryScanStatus)
	recoveryScanStatusModel.EndTimeUsecs = core.Int64Ptr(int64(26))
	recoveryScanStatusModel.ErrorMessage = core.StringPtr("testString")
	recoveryScanStatusModel.StartTimeUsecs = core.Int64Ptr(int64(26))
	recoveryScanStatusModel.Status = core.StringPtr("Accepted")
	recoveryScanStatusModel.Uid = core.StringPtr("testString")

	model := new(backuprecoveryv1.BatchVaultRecoveryScanStatus)
	model.Status = recoveryScanStatusModel
	model.VaultID = core.Int64Ptr(int64(26))

	result, err := backuprecovery.DataSourceIbmBackupRecoveryVaultRecoveryScanStatusBatchVaultRecoveryScanStatusToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryVaultRecoveryScanStatusRecoveryScanStatusToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["end_time_usecs"] = int(26)
		model["error_message"] = "testString"
		model["start_time_usecs"] = int(26)
		model["status"] = "Accepted"
		model["uid"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.RecoveryScanStatus)
	model.EndTimeUsecs = core.Int64Ptr(int64(26))
	model.ErrorMessage = core.StringPtr("testString")
	model.StartTimeUsecs = core.Int64Ptr(int64(26))
	model.Status = core.StringPtr("Accepted")
	model.Uid = core.StringPtr("testString")

	result, err := backuprecovery.DataSourceIbmBackupRecoveryVaultRecoveryScanStatusRecoveryScanStatusToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
