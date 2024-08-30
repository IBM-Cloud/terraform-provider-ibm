// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.94.0-fa797aec-20240814-142622
 */

package backuprecovery_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/backuprecovery"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/stretchr/testify/assert"
	"github.ibm.com/BackupAndRecovery/ibm-backup-recovery-sdk-go/backuprecoveryv1"
)

func TestAccIbmDataSourceConnectionsDataSourceBasic(t *testing.T) {
	dataSourceConnectionConnectionName := fmt.Sprintf("tf_connection_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmDataSourceConnectionsDataSourceConfigBasic(dataSourceConnectionConnectionName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_data_source_connections.data_source_connections_instance", "id"),
				),
			},
		},
	})
}

func TestAccIbmDataSourceConnectionsDataSourceAllArgs(t *testing.T) {
	dataSourceConnectionTenantID := fmt.Sprintf("%d", acctest.RandIntRange(10, 100))
	dataSourceConnectionConnectionName := fmt.Sprintf("tf_connection_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmDataSourceConnectionsDataSourceConfig(dataSourceConnectionTenantID, dataSourceConnectionConnectionName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_data_source_connections.data_source_connections_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_data_source_connections.data_source_connections_instance", "connection_ids"),
					resource.TestCheckResourceAttrSet("data.ibm_data_source_connections.data_source_connections_instance", "tenant_id"),
					resource.TestCheckResourceAttrSet("data.ibm_data_source_connections.data_source_connections_instance", "connection_names"),
					resource.TestCheckResourceAttrSet("data.ibm_data_source_connections.data_source_connections_instance", "connections.#"),
					resource.TestCheckResourceAttrSet("data.ibm_data_source_connections.data_source_connections_instance", "connections.0.id"),
					resource.TestCheckResourceAttr("data.ibm_data_source_connections.data_source_connections_instance", "connections.0.connection_name", dataSourceConnectionConnectionName),
					resource.TestCheckResourceAttrSet("data.ibm_data_source_connections.data_source_connections_instance", "connections.0.registration_token"),
				),
			},
		},
	})
}

func testAccCheckIbmDataSourceConnectionsDataSourceConfigBasic(dataSourceConnectionConnectionName string) string {
	return fmt.Sprintf(`
		resource "ibm_data_source_connection" "data_source_connection_instance" {
			connection_name = "%s"
		}

		data "ibm_data_source_connections" "data_source_connections_instance" {
			connection_ids = [ "connectionIds" ]
			tenant_id = ibm_data_source_connection.data_source_connection_instance.tenant_id
			connection_names = [ "connectionNames" ]
		}
	`, dataSourceConnectionConnectionName)
}

func testAccCheckIbmDataSourceConnectionsDataSourceConfig(dataSourceConnectionTenantID string, dataSourceConnectionConnectionName string) string {
	return fmt.Sprintf(`
		resource "ibm_data_source_connection" "data_source_connection_instance" {
			tenant_id = %s
			connection_name = "%s"
		}

		data "ibm_data_source_connections" "data_source_connections_instance" {
			connection_ids = [ "connectionIds" ]
			tenant_id = ibm_data_source_connection.data_source_connection_instance.tenant_id
			connection_names = [ "connectionNames" ]
		}
	`, dataSourceConnectionTenantID, dataSourceConnectionConnectionName)
}

func TestDataSourceIbmDataSourceConnectionsDataSourceConnectionToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		networkSettingsModel := make(map[string]interface{})
		networkSettingsModel["cluster_fqdn"] = "testString"
		networkSettingsModel["dns"] = []string{"testString"}
		networkSettingsModel["network_gateway"] = "testString"
		networkSettingsModel["ntp"] = "testString"

		model := make(map[string]interface{})
		model["id"] = "testString"
		model["connection_name"] = "testString"
		model["connector_ids"] = []string{"testString"}
		model["network_settings"] = []map[string]interface{}{networkSettingsModel}
		model["registration_token"] = "testString"

		assert.Equal(t, result, model)
	}

	networkSettingsModel := new(backuprecoveryv1.NetworkSettings)
	networkSettingsModel.ClusterFqdn = core.StringPtr("testString")
	networkSettingsModel.Dns = []string{"testString"}
	networkSettingsModel.NetworkGateway = core.StringPtr("testString")
	networkSettingsModel.Ntp = core.StringPtr("testString")

	model := new(backuprecoveryv1.DataSourceConnection)
	model.ConnectionID = core.StringPtr("testString")
	model.ConnectionName = core.StringPtr("testString")
	model.ConnectorIds = []string{"testString"}
	model.NetworkSettings = networkSettingsModel
	model.RegistrationToken = core.StringPtr("testString")

	result, err := backuprecovery.DataSourceIbmDataSourceConnectionsDataSourceConnectionToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmDataSourceConnectionsNetworkSettingsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["cluster_fqdn"] = "testString"
		model["dns"] = []string{"testString"}
		model["network_gateway"] = "testString"
		model["ntp"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.NetworkSettings)
	model.ClusterFqdn = core.StringPtr("testString")
	model.Dns = []string{"testString"}
	model.NetworkGateway = core.StringPtr("testString")
	model.Ntp = core.StringPtr("testString")

	result, err := backuprecovery.DataSourceIbmDataSourceConnectionsNetworkSettingsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
