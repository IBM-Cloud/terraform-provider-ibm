// Copyright IBM Corp. 2024 All Rights Reserved.
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
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/backuprecovery"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/stretchr/testify/assert"
	"github.ibm.com/BackupAndRecovery/ibm-backup-recovery-sdk-go/backuprecoveryv1"
)

func TestAccIbmDataSourceConnectorsDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmDataSourceConnectorsDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_data_source_connectors.data_source_connectors_instance", "id"),
				),
			},
		},
	})
}

func testAccCheckIbmDataSourceConnectorsDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_data_source_connectors" "data_source_connectors_instance" {
			connectorIds = [ "connectorIds" ]
			connectorNames = [ "connectorNames" ]
			tenantId = "tenantId"
			connectionId = "connectionId"
		}
	`)
}

func TestDataSourceIbmDataSourceConnectorsDataSourceConnectorToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		dataSourceConnectorStatusModel := make(map[string]interface{})
		dataSourceConnectorStatusModel["is_connected"] = true
		dataSourceConnectorStatusModel["last_connected_timestamp_secs"] = int(26)
		dataSourceConnectorStatusModel["message"] = "testString"

		model := make(map[string]interface{})
		model["cluster_side_ip"] = "testString"
		model["connection_id"] = "testString"
		model["connector_id"] = "testString"
		model["connector_name"] = "testString"
		model["connector_status"] = []map[string]interface{}{dataSourceConnectorStatusModel}
		model["software_version"] = "testString"
		model["tenant_side_ip"] = "testString"

		assert.Equal(t, result, model)
	}

	dataSourceConnectorStatusModel := new(backuprecoveryv1.DataSourceConnectorStatus)
	dataSourceConnectorStatusModel.IsConnected = core.BoolPtr(true)
	dataSourceConnectorStatusModel.LastConnectedTimestampSecs = core.Int64Ptr(int64(26))
	dataSourceConnectorStatusModel.Message = core.StringPtr("testString")

	model := new(backuprecoveryv1.DataSourceConnector)
	model.ClusterSideIp = core.StringPtr("testString")
	model.ConnectionID = core.StringPtr("testString")
	model.ConnectorID = core.StringPtr("testString")
	model.ConnectorName = core.StringPtr("testString")
	model.ConnectorStatus = dataSourceConnectorStatusModel
	model.SoftwareVersion = core.StringPtr("testString")
	model.TenantSideIp = core.StringPtr("testString")

	result, err := backuprecovery.DataSourceIbmDataSourceConnectorsDataSourceConnectorToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmDataSourceConnectorsDataSourceConnectorStatusToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["is_connected"] = true
		model["last_connected_timestamp_secs"] = int(26)
		model["message"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.DataSourceConnectorStatus)
	model.IsConnected = core.BoolPtr(true)
	model.LastConnectedTimestampSecs = core.Int64Ptr(int64(26))
	model.Message = core.StringPtr("testString")

	result, err := backuprecovery.DataSourceIbmDataSourceConnectorsDataSourceConnectorStatusToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
