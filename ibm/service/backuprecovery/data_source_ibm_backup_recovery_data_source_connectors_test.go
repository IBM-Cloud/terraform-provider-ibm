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
)

func TestAccIbmBackupRecoveryDataSourceConnectorsDataSourceBasic(t *testing.T) {
	dataSourceConnectorConnectionId := "5128356219792164864"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmBackupRecoveryDataSourceConnectorsDataSourceConfigBasic(dataSourceConnectorConnectionId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_data_source_connectors.baas_data_source_connectors_instance", "id"),
					resource.TestCheckResourceAttr("data.ibm_backup_recovery_data_source_connectors.baas_data_source_connectors_instance", "x_ibm_tenant_id", tenantId),
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_data_source_connectors.baas_data_source_connectors_instance", "connectors.#"),
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_data_source_connectors.baas_data_source_connectors_instance", "connectors.0.connection_id"),
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_data_source_connectors.baas_data_source_connectors_instance", "connectors.0.connector_name"),
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_data_source_connectors.baas_data_source_connectors_instance", "connectors.0.tenant_side_ip"),
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_data_source_connectors.baas_data_source_connectors_instance", "connectors.0.cluster_side_ip"),
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_data_source_connectors.baas_data_source_connectors_instance", "connectors.0.software_version"),
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_data_source_connectors.baas_data_source_connectors_instance", "connectors.0.connectivity_status.#"),
				),
			},
		},
	})
}

func testAccCheckIbmBackupRecoveryDataSourceConnectorsDataSourceConfigBasic(dataSourceConnectorConnectioID string) string {
	return fmt.Sprintf(`
		data "ibm_backup_recovery_data_source_connectors" "baas_data_source_connectors_instance" {
			x_ibm_tenant_id = "%s"
			
			connection_id = "%s"
		  }
	`, tenantId, dataSourceConnectorConnectioID)
}
