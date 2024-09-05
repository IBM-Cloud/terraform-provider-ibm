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

func TestAccIbmBaasDataSourceConnectorPatchBasic(t *testing.T) {
	var conf backuprecoveryv1.DataSourceConnectorList
	connectorID := fmt.Sprintf("tf_connector_id_%d", acctest.RandIntRange(10, 100))
	xIbmTenantID := fmt.Sprintf("tf_x_ibm_tenant_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmBaasDataSourceConnectorPatchDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmBaasDataSourceConnectorPatchConfigBasic(connectorID, xIbmTenantID),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmBaasDataSourceConnectorPatchExists("ibm_baas_data_source_connector_patch.baas_data_source_connector_patch_instance", conf),
					resource.TestCheckResourceAttr("ibm_baas_data_source_connector_patch.baas_data_source_connector_patch_instance", "connector_id", connectorID),
					resource.TestCheckResourceAttr("ibm_baas_data_source_connector_patch.baas_data_source_connector_patch_instance", "x_ibm_tenant_id", xIbmTenantID),
				),
			},
		},
	})
}

func TestAccIbmBaasDataSourceConnectorPatchAllArgs(t *testing.T) {
	var conf backuprecoveryv1.DataSourceConnectorList
	connectorID := fmt.Sprintf("tf_connector_id_%d", acctest.RandIntRange(10, 100))
	xIbmTenantID := fmt.Sprintf("tf_x_ibm_tenant_id_%d", acctest.RandIntRange(10, 100))
	connectorName := fmt.Sprintf("tf_connector_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmBaasDataSourceConnectorPatchDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmBaasDataSourceConnectorPatchConfig(connectorID, xIbmTenantID, connectorName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmBaasDataSourceConnectorPatchExists("ibm_baas_data_source_connector_patch.baas_data_source_connector_patch_instance", conf),
					resource.TestCheckResourceAttr("ibm_baas_data_source_connector_patch.baas_data_source_connector_patch_instance", "connector_id", connectorID),
					resource.TestCheckResourceAttr("ibm_baas_data_source_connector_patch.baas_data_source_connector_patch_instance", "x_ibm_tenant_id", xIbmTenantID),
					resource.TestCheckResourceAttr("ibm_baas_data_source_connector_patch.baas_data_source_connector_patch_instance", "connector_name", connectorName),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_baas_data_source_connector_patch.baas_data_source_connector_patch",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmBaasDataSourceConnectorPatchConfigBasic(connectorID string, xIbmTenantID string) string {
	return fmt.Sprintf(`
		resource "ibm_baas_data_source_connector_patch" "baas_data_source_connector_patch_instance" {
			connector_id = "%s"
			x_ibm_tenant_id = "%s"
		}
	`, connectorID, xIbmTenantID)
}

func testAccCheckIbmBaasDataSourceConnectorPatchConfig(connectorID string, xIbmTenantID string, connectorName string) string {
	return fmt.Sprintf(`

		resource "ibm_baas_data_source_connector_patch" "baas_data_source_connector_patch_instance" {
			connector_id = "%s"
			x_ibm_tenant_id = "%s"
			connector_name = "%s"
		}
	`, connectorID, xIbmTenantID, connectorName)
}

func testAccCheckIbmBaasDataSourceConnectorPatchExists(n string, obj backuprecoveryv1.DataSourceConnectorList) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		backupRecoveryClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).BackupRecoveryV1()
		if err != nil {
			return err
		}

		getDataSourceConnectorsOptions := &backuprecoveryv1.GetDataSourceConnectorsOptions{}

		getDataSourceConnectorsOptions.SetConnectorIds([]string{rs.Primary.ID})

		dataSourceConnector, _, err := backupRecoveryClient.GetDataSourceConnectors(getDataSourceConnectorsOptions)
		if err != nil {
			return err
		}

		obj = *dataSourceConnector
		return nil
	}
}

func testAccCheckIbmBaasDataSourceConnectorPatchDestroy(s *terraform.State) error {
	backupRecoveryClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).BackupRecoveryV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_baas_data_source_connector_patch" {
			continue
		}

		getDataSourceConnectorsOptions := &backuprecoveryv1.GetDataSourceConnectorsOptions{}

		getDataSourceConnectorsOptions.SetConnectorIds([]string{rs.Primary.ID})

		// Try to find the key
		_, response, err := backupRecoveryClient.GetDataSourceConnectors(getDataSourceConnectorsOptions)

		if err == nil {
			return fmt.Errorf("Data-Source Connector still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for Data-Source Connector (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

func TestResourceIbmBaasDataSourceConnectorPatchDataSourceConnectorStatusToMap(t *testing.T) {
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

	result, err := backuprecovery.ResourceIbmBaasDataSourceConnectorPatchDataSourceConnectorStatusToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
