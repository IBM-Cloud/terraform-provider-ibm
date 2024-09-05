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

func TestAccIbmBaasDataSourceConnectionsDataSourceBasic(t *testing.T) {
	dataSourceConnectionConnectionName := fmt.Sprintf("tf_connection_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmBaasDataSourceConnectionsDataSourceConfigBasic(dataSourceConnectionConnectionName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_baas_data_source_connections.baas_data_source_connections_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_data_source_connections.baas_data_source_connections_instance", "tenant_id"),
				),
			},
		},
	})
}

func TestAccIbmBaasDataSourceConnectionsDataSourceAllArgs(t *testing.T) {
	dataSourceConnectionTenantID := fmt.Sprintf("%d", acctest.RandIntRange(10, 100))
	dataSourceConnectionConnectionName := fmt.Sprintf("tf_connection_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmBaasDataSourceConnectionsDataSourceConfig(dataSourceConnectionTenantID, dataSourceConnectionConnectionName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_baas_data_source_connections.baas_data_source_connections_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_data_source_connections.baas_data_source_connections_instance", "tenant_id"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_data_source_connections.baas_data_source_connections_instance", "connection_ids"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_data_source_connections.baas_data_source_connections_instance", "connection_names"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_data_source_connections.baas_data_source_connections_instance", "connections.#"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_data_source_connections.baas_data_source_connections_instance", "connections.0.id"),
					resource.TestCheckResourceAttr("data.ibm_baas_data_source_connections.baas_data_source_connections_instance", "connections.0.connection_name", dataSourceConnectionConnectionName),
					resource.TestCheckResourceAttrSet("data.ibm_baas_data_source_connections.baas_data_source_connections_instance", "connections.0.registration_token"),
				),
			},
		},
	})
}

func testAccCheckIbmBaasDataSourceConnectionsDataSourceConfigBasic(dataSourceConnectionConnectionName string) string {
	return fmt.Sprintf(`
		resource "ibm_baas_data_source_connection" "baas_data_source_connection_instance" {
			connection_name = "%s"
		}

		data "ibm_baas_data_source_connections" "baas_data_source_connections_instance" {
			tenant_id = ibm_baas_data_source_connection.baas_data_source_connection_instance.tenant_id
			connection_ids = [ "connectionIds" ]
			connection_names = [ "connectionNames" ]
		}
	`, dataSourceConnectionConnectionName)
}

func testAccCheckIbmBaasDataSourceConnectionsDataSourceConfig(dataSourceConnectionTenantID string, dataSourceConnectionConnectionName string) string {
	return fmt.Sprintf(`
		resource "ibm_baas_data_source_connection" "baas_data_source_connection_instance" {
			tenant_id = %s
			connection_name = "%s"
		}

		data "ibm_baas_data_source_connections" "baas_data_source_connections_instance" {
			tenant_id = ibm_baas_data_source_connection.baas_data_source_connection_instance.tenant_id
			connection_ids = [ "connectionIds" ]
			connection_names = [ "connectionNames" ]
		}
	`, dataSourceConnectionTenantID, dataSourceConnectionConnectionName)
}

func TestDataSourceIbmBaasDataSourceConnectionsDataSourceConnectionToMap(t *testing.T) {
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

	result, err := backuprecovery.DataSourceIbmBaasDataSourceConnectionsDataSourceConnectionToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasDataSourceConnectionsNetworkSettingsToMap(t *testing.T) {
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

	result, err := backuprecovery.DataSourceIbmBaasDataSourceConnectionsNetworkSettingsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
