// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package backuprecovery_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/backuprecovery"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/stretchr/testify/assert"
	"github.ibm.com/BackupAndRecovery/ibm-backup-recovery-sdk-go/backuprecoveryv1"
)

func TestAccIbmBaasDataSourceConnectionBasic(t *testing.T) {
	var conf backuprecoveryv1.DataSourceConnection
	connectionName := fmt.Sprintf("tf_connection_name_%d", acctest.RandIntRange(10, 100))
	connectionNameUpdate := fmt.Sprintf("tf_connection_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmBaasDataSourceConnectionDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmBaasDataSourceConnectionConfigBasic(connectionName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmBaasDataSourceConnectionExists("ibm_baas_data_source_connection.baas_data_source_connection_instance", conf),
					resource.TestCheckResourceAttr("ibm_baas_data_source_connection.baas_data_source_connection_instance", "connection_name", connectionName),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmBaasDataSourceConnectionConfigBasic(connectionNameUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_baas_data_source_connection.baas_data_source_connection_instance", "connection_name", connectionNameUpdate),
				),
			},
		},
	})
}

func TestAccIbmBaasDataSourceConnectionAllArgs(t *testing.T) {
	var conf backuprecoveryv1.DataSourceConnection
	xIbmTenantID := fmt.Sprintf("tf_x_ibm_tenant_id_%d", acctest.RandIntRange(10, 100))
	connectionName := fmt.Sprintf("tf_connection_name_%d", acctest.RandIntRange(10, 100))
	xIbmTenantIDUpdate := fmt.Sprintf("tf_x_ibm_tenant_id_%d", acctest.RandIntRange(10, 100))
	connectionNameUpdate := fmt.Sprintf("tf_connection_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmBaasDataSourceConnectionDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmBaasDataSourceConnectionConfig(xIbmTenantID, connectionName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmBaasDataSourceConnectionExists("ibm_baas_data_source_connection.baas_data_source_connection_instance", conf),
					resource.TestCheckResourceAttr("ibm_baas_data_source_connection.baas_data_source_connection_instance", "x_ibm_tenant_id", xIbmTenantID),
					resource.TestCheckResourceAttr("ibm_baas_data_source_connection.baas_data_source_connection_instance", "connection_name", connectionName),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmBaasDataSourceConnectionConfig(xIbmTenantIDUpdate, connectionNameUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_baas_data_source_connection.baas_data_source_connection_instance", "x_ibm_tenant_id", xIbmTenantIDUpdate),
					resource.TestCheckResourceAttr("ibm_baas_data_source_connection.baas_data_source_connection_instance", "connection_name", connectionNameUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_baas_data_source_connection.baas_data_source_connection",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmBaasDataSourceConnectionConfigBasic(connectionName string) string {
	return fmt.Sprintf(`
		resource "ibm_baas_data_source_connection" "baas_data_source_connection_instance" {
			connection_name = "%s"
		}
	`, connectionName)
}

func testAccCheckIbmBaasDataSourceConnectionConfig(xIbmTenantID string, connectionName string) string {
	return fmt.Sprintf(`

		resource "ibm_baas_data_source_connection" "baas_data_source_connection_instance" {
			x_ibm_tenant_id = "%s"
			connection_name = "%s"
		}
	`, xIbmTenantID, connectionName)
}

func testAccCheckIbmBaasDataSourceConnectionExists(n string, obj backuprecoveryv1.DataSourceConnection) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		backupRecoveryClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).BackupRecoveryV1()
		if err != nil {
			return err
		}

		getDataSourceConnectionsOptions := &backuprecoveryv1.GetDataSourceConnectionsOptions{}

		getDataSourceConnectionsOptions.SetConnectionIds([]string{rs.Primary.ID})

		dataSourceConnection, _, err := backupRecoveryClient.GetDataSourceConnections(getDataSourceConnectionsOptions)
		if err != nil {
			return err
		}

		obj = *&dataSourceConnection.Connections[0]
		return nil
	}
}

func testAccCheckIbmBaasDataSourceConnectionDestroy(s *terraform.State) error {
	backupRecoveryClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).BackupRecoveryV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_baas_data_source_connection" {
			continue
		}

		getDataSourceConnectionsOptions := &backuprecoveryv1.GetDataSourceConnectionsOptions{}

		getDataSourceConnectionsOptions.SetConnectionIds([]string{rs.Primary.ID})

		// Try to find the key
		_, response, err := backupRecoveryClient.GetDataSourceConnections(getDataSourceConnectionsOptions)

		if err == nil {
			return fmt.Errorf("Data-Source Connection still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for Data-Source Connection (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

func TestResourceIbmBaasDataSourceConnectionNetworkSettingsToMap(t *testing.T) {
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

	result, err := backuprecovery.ResourceIbmBaasDataSourceConnectionNetworkSettingsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
