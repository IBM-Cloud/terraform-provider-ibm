// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.114.3-943fbc81-20260603-173645
 */

package backuprecovery_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/backuprecovery"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/ibm-backup-recovery-sdk-go/backuprecoveryv1"
	"github.com/stretchr/testify/assert"
)

func TestAccIbmBackupRecoveryConnectorAgentsUpgradeStatusDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmBackupRecoveryConnectorAgentsUpgradeStatusDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_connector_agents_upgrade_status.backup_recovery_connector_agents_upgrade_status_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_connector_agents_upgrade_status.backup_recovery_connector_agents_upgrade_status_instance", "upgrade_id"),
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_connector_agents_upgrade_status.backup_recovery_connector_agents_upgrade_status_instance", "status"),
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_connector_agents_upgrade_status.backup_recovery_connector_agents_upgrade_status_instance", "tenant_statuses.#"),
				),
			},
		},
	})
}

func testAccCheckIbmBackupRecoveryConnectorAgentsUpgradeStatusDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_backup_recovery_connector_agents_upgrade_status" "backup_recovery_connector_agents_upgrade_status_instance" {
			upgrade_id = "9fab83da-98cb-4f18-a7ba-b6f0435c9673"
			tenant_ids = [ "tenantIds" ]
			connection_ids = [ 1 ]
		}
	`)
}

func TestDataSourceIbmBackupRecoveryConnectorAgentsUpgradeStatusConnectorAgentTenantUpgradeStatusToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		connectorAgentUpgradeConnectionStatusModel := make(map[string]interface{})
		connectorAgentUpgradeConnectionStatusModel["connection_id"] = int(26)
		connectorAgentUpgradeConnectionStatusModel["status"] = "NotStarted"
		connectorAgentUpgradeConnectionStatusModel["attempt_count"] = int(38)
		connectorAgentUpgradeConnectionStatusModel["last_attempt_timestamp_msecs"] = int(26)
		connectorAgentUpgradeConnectionStatusModel["last_status_fetched_timestamp_msecs"] = int(26)
		connectorAgentUpgradeConnectionStatusModel["error_message"] = "testString"

		connectorAgentTenantUpgradeInfoModel := make(map[string]interface{})
		connectorAgentTenantUpgradeInfoModel["state"] = "NotStarted"
		connectorAgentTenantUpgradeInfoModel["target_version"] = "testString"
		connectorAgentTenantUpgradeInfoModel["retry_forever"] = true
		connectorAgentTenantUpgradeInfoModel["total_connections"] = int(38)
		connectorAgentTenantUpgradeInfoModel["succeeded_connections"] = int(38)
		connectorAgentTenantUpgradeInfoModel["failed_connections"] = int(38)
		connectorAgentTenantUpgradeInfoModel["created_timestamp_usecs"] = int(26)
		connectorAgentTenantUpgradeInfoModel["finished_timestamp_usecs"] = int(26)
		connectorAgentTenantUpgradeInfoModel["connection_statuses"] = []map[string]interface{}{connectorAgentUpgradeConnectionStatusModel}

		model := make(map[string]interface{})
		model["tenant_id"] = "testString"
		model["info"] = []map[string]interface{}{connectorAgentTenantUpgradeInfoModel}

		assert.Equal(t, result, model)
	}

	connectorAgentUpgradeConnectionStatusModel := new(backuprecoveryv1.ConnectorAgentUpgradeConnectionStatus)
	connectorAgentUpgradeConnectionStatusModel.ConnectionID = core.Int64Ptr(int64(26))
	connectorAgentUpgradeConnectionStatusModel.Status = core.StringPtr("NotStarted")
	connectorAgentUpgradeConnectionStatusModel.AttemptCount = core.Int64Ptr(int64(38))
	connectorAgentUpgradeConnectionStatusModel.LastAttemptTimestampMsecs = core.Int64Ptr(int64(26))
	connectorAgentUpgradeConnectionStatusModel.LastStatusFetchedTimestampMsecs = core.Int64Ptr(int64(26))
	connectorAgentUpgradeConnectionStatusModel.ErrorMessage = core.StringPtr("testString")

	connectorAgentTenantUpgradeInfoModel := new(backuprecoveryv1.ConnectorAgentTenantUpgradeInfo)
	connectorAgentTenantUpgradeInfoModel.State = core.StringPtr("NotStarted")
	connectorAgentTenantUpgradeInfoModel.TargetVersion = core.StringPtr("testString")
	connectorAgentTenantUpgradeInfoModel.RetryForever = core.BoolPtr(true)
	connectorAgentTenantUpgradeInfoModel.TotalConnections = core.Int64Ptr(int64(38))
	connectorAgentTenantUpgradeInfoModel.SucceededConnections = core.Int64Ptr(int64(38))
	connectorAgentTenantUpgradeInfoModel.FailedConnections = core.Int64Ptr(int64(38))
	connectorAgentTenantUpgradeInfoModel.CreatedTimestampUsecs = core.Int64Ptr(int64(26))
	connectorAgentTenantUpgradeInfoModel.FinishedTimestampUsecs = core.Int64Ptr(int64(26))
	connectorAgentTenantUpgradeInfoModel.ConnectionStatuses = []backuprecoveryv1.ConnectorAgentUpgradeConnectionStatus{*connectorAgentUpgradeConnectionStatusModel}

	model := new(backuprecoveryv1.ConnectorAgentTenantUpgradeStatus)
	model.TenantID = core.StringPtr("testString")
	model.Info = connectorAgentTenantUpgradeInfoModel

	result, err := backuprecovery.DataSourceIbmBackupRecoveryConnectorAgentsUpgradeStatusConnectorAgentTenantUpgradeStatusToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryConnectorAgentsUpgradeStatusConnectorAgentTenantUpgradeInfoToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		connectorAgentUpgradeConnectionStatusModel := make(map[string]interface{})
		connectorAgentUpgradeConnectionStatusModel["connection_id"] = int(26)
		connectorAgentUpgradeConnectionStatusModel["status"] = "NotStarted"
		connectorAgentUpgradeConnectionStatusModel["attempt_count"] = int(38)
		connectorAgentUpgradeConnectionStatusModel["last_attempt_timestamp_msecs"] = int(26)
		connectorAgentUpgradeConnectionStatusModel["last_status_fetched_timestamp_msecs"] = int(26)
		connectorAgentUpgradeConnectionStatusModel["error_message"] = "testString"

		model := make(map[string]interface{})
		model["state"] = "NotStarted"
		model["target_version"] = "testString"
		model["retry_forever"] = true
		model["total_connections"] = int(38)
		model["succeeded_connections"] = int(38)
		model["failed_connections"] = int(38)
		model["created_timestamp_usecs"] = int(26)
		model["finished_timestamp_usecs"] = int(26)
		model["connection_statuses"] = []map[string]interface{}{connectorAgentUpgradeConnectionStatusModel}

		assert.Equal(t, result, model)
	}

	connectorAgentUpgradeConnectionStatusModel := new(backuprecoveryv1.ConnectorAgentUpgradeConnectionStatus)
	connectorAgentUpgradeConnectionStatusModel.ConnectionID = core.Int64Ptr(int64(26))
	connectorAgentUpgradeConnectionStatusModel.Status = core.StringPtr("NotStarted")
	connectorAgentUpgradeConnectionStatusModel.AttemptCount = core.Int64Ptr(int64(38))
	connectorAgentUpgradeConnectionStatusModel.LastAttemptTimestampMsecs = core.Int64Ptr(int64(26))
	connectorAgentUpgradeConnectionStatusModel.LastStatusFetchedTimestampMsecs = core.Int64Ptr(int64(26))
	connectorAgentUpgradeConnectionStatusModel.ErrorMessage = core.StringPtr("testString")

	model := new(backuprecoveryv1.ConnectorAgentTenantUpgradeInfo)
	model.State = core.StringPtr("NotStarted")
	model.TargetVersion = core.StringPtr("testString")
	model.RetryForever = core.BoolPtr(true)
	model.TotalConnections = core.Int64Ptr(int64(38))
	model.SucceededConnections = core.Int64Ptr(int64(38))
	model.FailedConnections = core.Int64Ptr(int64(38))
	model.CreatedTimestampUsecs = core.Int64Ptr(int64(26))
	model.FinishedTimestampUsecs = core.Int64Ptr(int64(26))
	model.ConnectionStatuses = []backuprecoveryv1.ConnectorAgentUpgradeConnectionStatus{*connectorAgentUpgradeConnectionStatusModel}

	result, err := backuprecovery.DataSourceIbmBackupRecoveryConnectorAgentsUpgradeStatusConnectorAgentTenantUpgradeInfoToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryConnectorAgentsUpgradeStatusConnectorAgentUpgradeConnectionStatusToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["connection_id"] = int(26)
		model["status"] = "NotStarted"
		model["attempt_count"] = int(38)
		model["last_attempt_timestamp_msecs"] = int(26)
		model["last_status_fetched_timestamp_msecs"] = int(26)
		model["error_message"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.ConnectorAgentUpgradeConnectionStatus)
	model.ConnectionID = core.Int64Ptr(int64(26))
	model.Status = core.StringPtr("NotStarted")
	model.AttemptCount = core.Int64Ptr(int64(38))
	model.LastAttemptTimestampMsecs = core.Int64Ptr(int64(26))
	model.LastStatusFetchedTimestampMsecs = core.Int64Ptr(int64(26))
	model.ErrorMessage = core.StringPtr("testString")

	result, err := backuprecovery.DataSourceIbmBackupRecoveryConnectorAgentsUpgradeStatusConnectorAgentUpgradeConnectionStatusToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
