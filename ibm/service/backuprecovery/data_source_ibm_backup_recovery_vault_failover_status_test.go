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

func TestAccIbmBackupRecoveryVaultFailoverStatusDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmBackupRecoveryVaultFailoverStatusDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_vault_failover_status.backup_recovery_vault_failover_status_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_vault_failover_status.backup_recovery_vault_failover_status_instance", "x_ibm_tenant_id"),
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_vault_failover_status.backup_recovery_vault_failover_status_instance", "cloud_type"),
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_vault_failover_status.backup_recovery_vault_failover_status_instance", "vault_ids"),
				),
			},
		},
	})
}

func testAccCheckIbmBackupRecoveryVaultFailoverStatusDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_backup_recovery_vault_failover_status" "backup_recovery_vault_failover_status_instance" {
			X-IBM-Tenant-Id = "tenantId"
			cloudType = "ibm"
			vaultIds = [ 1 ]
		}
	`)
}


func TestDataSourceIbmBackupRecoveryVaultFailoverStatusBatchVaultFailoverStatusToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		vaultFailoverStatusModel := make(map[string]interface{})
		vaultFailoverStatusModel["end_time_usecs"] = int(26)
		vaultFailoverStatusModel["error_message"] = "testString"
		vaultFailoverStatusModel["start_time_usecs"] = int(26)
		vaultFailoverStatusModel["status"] = "Accepted"
		vaultFailoverStatusModel["uid"] = "testString"

		model := make(map[string]interface{})
		model["status"] = []map[string]interface{}{vaultFailoverStatusModel}
		model["vault_id"] = int(26)

		assert.Equal(t, result, model)
	}

	vaultFailoverStatusModel := new(backuprecoveryv1.VaultFailoverStatus)
	vaultFailoverStatusModel.EndTimeUsecs = core.Int64Ptr(int64(26))
	vaultFailoverStatusModel.ErrorMessage = core.StringPtr("testString")
	vaultFailoverStatusModel.StartTimeUsecs = core.Int64Ptr(int64(26))
	vaultFailoverStatusModel.Status = core.StringPtr("Accepted")
	vaultFailoverStatusModel.Uid = core.StringPtr("testString")

	model := new(backuprecoveryv1.BatchVaultFailoverStatus)
	model.Status = vaultFailoverStatusModel
	model.VaultID = core.Int64Ptr(int64(26))

	result, err := backuprecovery.DataSourceIbmBackupRecoveryVaultFailoverStatusBatchVaultFailoverStatusToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryVaultFailoverStatusVaultFailoverStatusToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["end_time_usecs"] = int(26)
		model["error_message"] = "testString"
		model["start_time_usecs"] = int(26)
		model["status"] = "Accepted"
		model["uid"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.VaultFailoverStatus)
	model.EndTimeUsecs = core.Int64Ptr(int64(26))
	model.ErrorMessage = core.StringPtr("testString")
	model.StartTimeUsecs = core.Int64Ptr(int64(26))
	model.Status = core.StringPtr("Accepted")
	model.Uid = core.StringPtr("testString")

	result, err := backuprecovery.DataSourceIbmBackupRecoveryVaultFailoverStatusVaultFailoverStatusToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
