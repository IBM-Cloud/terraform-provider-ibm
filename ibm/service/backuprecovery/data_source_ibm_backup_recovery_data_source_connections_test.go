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

const (
	tenantId string = "wkk1yqrdce/"
)

func TestAccIbmBackupRecoveryDataSourceConnectionsDataSourceBasic(t *testing.T) {
	dataSourceConnectionConnectionName := fmt.Sprintf("tf_connection_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:   func() { acc.TestAccPreCheck(t) },
		Providers:  acc.TestAccProviders,
		IsUnitTest: true,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmBackupRecoveryDataSourceConnectionsDataSourceConfigBasic(dataSourceConnectionConnectionName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_data_source_connections.baas_data_source_connections_instance", "id"),
					resource.TestCheckResourceAttr("data.ibm_backup_recovery_data_source_connections.baas_data_source_connections_instance", "x_ibm_tenant_id", tenantId),
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_data_source_connections.baas_data_source_connections_instance", "connections.0.connection_id"),
					resource.TestCheckResourceAttr("data.ibm_backup_recovery_data_source_connections.baas_data_source_connections_instance", "connections.0.tenant_id", tenantId),
					resource.TestCheckResourceAttr("data.ibm_backup_recovery_data_source_connections.baas_data_source_connections_instance", "connections.0.connector_ids.#", "0"),
					resource.TestCheckResourceAttr("data.ibm_backup_recovery_data_source_connections.baas_data_source_connections_instance", "connections.0.connection_name", dataSourceConnectionConnectionName),
				),
			},
		},
	})
}

func testAccCheckIbmBackupRecoveryDataSourceConnectionsDataSourceConfigBasic(connectionName string) string {
	return fmt.Sprintf(`
	
	resource "ibm_backup_recovery_data_source_connection" "baas_data_source_connection_instance" {
		x_ibm_tenant_id = "%s"
		backup_recovery_endpoint = "https://protectiondomain0103.us-east.backup-recovery-tests.cloud.ibm.com/v2"
		connection_name = "%s"
	  }
	
	data "ibm_backup_recovery_data_source_connections" "baas_data_source_connections_instance" {
		x_ibm_tenant_id = ibm_backup_recovery_data_source_connection.baas_data_source_connection_instance.x_ibm_tenant_id
		backup_recovery_endpoint = "https://protectiondomain0103.us-east.backup-recovery-tests.cloud.ibm.com/v2"
		connection_ids = [ibm_backup_recovery_data_source_connection.baas_data_source_connection_instance.connection_id]
	  }
		`, tenantId, connectionName)
}
