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

func TestAccIbmBaasDataSourceConnectorRegistrationBasic(t *testing.T) {
	var conf backuprecoveryv1.DataSourceConnectorList
	xIbmTenantID := fmt.Sprintf("tf_x_ibm_tenant_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmBaasDataSourceConnectorRegistrationDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmBaasDataSourceConnectorRegistrationConfigBasic(xIbmTenantID),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmBaasDataSourceConnectorRegistrationExists("ibm_baas_data_source_connector_registration.baas_data_source_connector_registration_instance", conf),
					resource.TestCheckResourceAttr("ibm_baas_data_source_connector_registration.baas_data_source_connector_registration_instance", "x_ibm_tenant_id", xIbmTenantID),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_baas_data_source_connector_registration.baas_data_source_connector_registration",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmBaasDataSourceConnectorRegistrationConfigBasic(xIbmTenantID string) string {
	return fmt.Sprintf(`
		resource "ibm_baas_data_source_connector_registration" "baas_data_source_connector_registration_instance" {
			x_ibm_tenant_id = "%s"
		}
	`, xIbmTenantID)
}

func testAccCheckIbmBaasDataSourceConnectorRegistrationExists(n string, obj backuprecoveryv1.DataSourceConnectorList) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		_, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		backupRecoveryClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).BackupRecoveryV1()
		if err != nil {
			return err
		}

		getDataSourceConnectorsOptions := &backuprecoveryv1.GetDataSourceConnectorsOptions{}

		dataSourceConnector, _, err := backupRecoveryClient.GetDataSourceConnectors(getDataSourceConnectorsOptions)
		if err != nil {
			return err
		}

		obj = *dataSourceConnector
		return nil
	}
}

func testAccCheckIbmBaasDataSourceConnectorRegistrationDestroy(s *terraform.State) error {
	backupRecoveryClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).BackupRecoveryV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_baas_data_source_connector_registration" {
			continue
		}

		getDataSourceConnectorsOptions := &backuprecoveryv1.GetDataSourceConnectorsOptions{}

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

func TestResourceIbmBaasDataSourceConnectorRegistrationDataSourceConnectorStatusToMap(t *testing.T) {
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

	result, err := backuprecovery.ResourceIbmBaasDataSourceConnectorRegistrationDataSourceConnectorStatusToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
