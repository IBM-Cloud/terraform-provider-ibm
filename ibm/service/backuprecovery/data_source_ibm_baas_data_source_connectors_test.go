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
)

func TestAccIbmBaasDataSourceConnectorsDataSourceBasic(t *testing.T) {
	dataSourceConnectorConnectorID := fmt.Sprintf("tf_connector_id_%d", acctest.RandIntRange(10, 100))
	dataSourceConnectorTenantID := fmt.Sprintf("%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmBaasDataSourceConnectorsDataSourceConfigBasic(dataSourceConnectorConnectorID, dataSourceConnectorTenantID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_baas_data_source_connectors.baas_data_source_connectors_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_data_source_connectors.baas_data_source_connectors_instance", "tenant_id"),
				),
			},
		},
	})
}

func TestAccIbmBaasDataSourceConnectorsDataSourceAllArgs(t *testing.T) {
	dataSourceConnectorConnectorID := fmt.Sprintf("tf_connector_id_%d", acctest.RandIntRange(10, 100))
	dataSourceConnectorTenantID := fmt.Sprintf("%d", acctest.RandIntRange(10, 100))
	dataSourceConnectorConnectorName := fmt.Sprintf("tf_connector_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmBaasDataSourceConnectorsDataSourceConfig(dataSourceConnectorConnectorID, dataSourceConnectorTenantID, dataSourceConnectorConnectorName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_baas_data_source_connectors.baas_data_source_connectors_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_data_source_connectors.baas_data_source_connectors_instance", "tenant_id"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_data_source_connectors.baas_data_source_connectors_instance", "connector_ids"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_data_source_connectors.baas_data_source_connectors_instance", "connector_names"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_data_source_connectors.baas_data_source_connectors_instance", "connection_id"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_data_source_connectors.baas_data_source_connectors_instance", "connectors.#"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_data_source_connectors.baas_data_source_connectors_instance", "connectors.0.cluster_side_ip"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_data_source_connectors.baas_data_source_connectors_instance", "connectors.0.connection_id"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_data_source_connectors.baas_data_source_connectors_instance", "connectors.0.id"),
					resource.TestCheckResourceAttr("data.ibm_baas_data_source_connectors.baas_data_source_connectors_instance", "connectors.0.connector_name", dataSourceConnectorConnectorName),
					resource.TestCheckResourceAttrSet("data.ibm_baas_data_source_connectors.baas_data_source_connectors_instance", "connectors.0.software_version"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_data_source_connectors.baas_data_source_connectors_instance", "connectors.0.tenant_side_ip"),
				),
			},
		},
	})
}

func testAccCheckIbmBaasDataSourceConnectorsDataSourceConfigBasic(dataSourceConnectorConnectorID string, dataSourceConnectorTenantID string) string {
	return fmt.Sprintf(`
		resource "ibm_baas_data_source_connector_patch" "baas_data_source_connector_patch_instance" {
			connector_id = "%s"
			tenant_id = %s
		}

		data "ibm_baas_data_source_connectors" "baas_data_source_connectors_instance" {
			tenant_id = ibm_baas_data_source_connector_patch.baas_data_source_connector_patch_instance.tenant_id
			connector_ids = [ "connectorIds" ]
			connector_names = [ "connectorNames" ]
			connection_id = ibm_baas_data_source_connector_patch.baas_data_source_connector_patch_instance.connection_id
		}
	`, dataSourceConnectorConnectorID, dataSourceConnectorTenantID)
}

func testAccCheckIbmBaasDataSourceConnectorsDataSourceConfig(dataSourceConnectorConnectorID string, dataSourceConnectorTenantID string, dataSourceConnectorConnectorName string) string {
	return fmt.Sprintf(`
		resource "ibm_baas_data_source_connector_patch" "baas_data_source_connector_patch_instance" {
			connector_id = "%s"
			tenant_id = %s
			connector_name = "%s"
		}

		data "ibm_baas_data_source_connectors" "baas_data_source_connectors_instance" {
			tenant_id = ibm_baas_data_source_connector_patch.baas_data_source_connector_patch_instance.tenant_id
			connector_ids = [ "connectorIds" ]
			connector_names = [ "connectorNames" ]
			connection_id = ibm_baas_data_source_connector_patch.baas_data_source_connector_patch_instance.connection_id
		}
	`, dataSourceConnectorConnectorID, dataSourceConnectorTenantID, dataSourceConnectorConnectorName)
}
