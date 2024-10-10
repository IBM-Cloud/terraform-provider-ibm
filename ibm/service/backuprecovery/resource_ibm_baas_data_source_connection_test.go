// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package backuprecovery_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
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
					testAccCheckIbmBaasDataSourceConnectionExists("ibm_backup_recovery_data_source_connection.baas_data_source_connection_instance", conf, connectionName),
					resource.TestCheckResourceAttr("ibm_backup_recovery_data_source_connection.baas_data_source_connection_instance", "connection_name", connectionName),
					resource.TestCheckResourceAttrSet("ibm_backup_recovery_data_source_connection.baas_data_source_connection_instance", "registration_token"),
					resource.TestCheckResourceAttrSet("ibm_backup_recovery_data_source_connection.baas_data_source_connection_instance", "connection_id"),
					resource.TestCheckResourceAttr("ibm_backup_recovery_data_source_connection.baas_data_source_connection_instance", "x_ibm_tenant_id", tenantId),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmBaasDataSourceConnectionConfigBasic(connectionNameUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_backup_recovery_data_source_connection.baas_data_source_connection_instance", "connection_name", connectionNameUpdate),
					resource.TestCheckResourceAttrSet("ibm_backup_recovery_data_source_connection.baas_data_source_connection_instance", "registration_token"),
					resource.TestCheckResourceAttrSet("ibm_backup_recovery_data_source_connection.baas_data_source_connection_instance", "id"),
					resource.TestCheckResourceAttr("ibm_backup_recovery_data_source_connection.baas_data_source_connection_instance", "x_ibm_tenant_id", tenantId),
				),
			},
		},
	})
}

func testAccCheckIbmBaasDataSourceConnectionConfigBasic(connectionName string) string {
	return fmt.Sprintf(`
	resource "ibm_backup_recovery_data_source_connection" "baas_data_source_connection_instance" {
		x_ibm_tenant_id = "%s"
		connection_name = "%s"
	  }
	`, tenantId, connectionName)
}

func testAccCheckIbmBaasDataSourceConnectionExists(resource string, obj backuprecoveryv1.DataSourceConnection, connName string) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[resource]
		if !ok {
			return fmt.Errorf("Not found: %s", resource)
		}

		backupRecoveryClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).BackupRecoveryV1()
		if err != nil {
			return err
		}

		getDataSourceConnectionsOptions := &backuprecoveryv1.GetDataSourceConnectionsOptions{}

		getDataSourceConnectionsOptions.SetXIBMTenantID(tenantId)
		getDataSourceConnectionsOptions.SetConnectionIds([]string{rs.Primary.ID})

		dataSourceConnection, _, err := backupRecoveryClient.GetDataSourceConnections(getDataSourceConnectionsOptions)
		if err != nil {
			return err
		}
		if (dataSourceConnection.Connections != nil) && (len(dataSourceConnection.Connections) > 0) && (*(dataSourceConnection.Connections[0].ConnectionName) == connName) {
			return nil
		} else {
			return fmt.Errorf("Not found: %s", resource)
		}

	}
}

func testAccCheckIbmBaasDataSourceConnectionDestroy(s *terraform.State) error {
	backupRecoveryClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).BackupRecoveryV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_backup_recovery_data_source_connection" {
			continue
		}

		getDataSourceConnectionsOptions := &backuprecoveryv1.GetDataSourceConnectionsOptions{}

		getDataSourceConnectionsOptions.SetXIBMTenantID(tenantId)
		getDataSourceConnectionsOptions.SetConnectionIds([]string{rs.Primary.ID})

		// Try to find the key
		_, response, err := backupRecoveryClient.GetDataSourceConnections(getDataSourceConnectionsOptions)
		if err == nil {
			return fmt.Errorf("Data-Source Connection still exists: %s", rs.Primary.ID)
		}
		if strings.Contains(response.String(), "does not exist in organization") {
			return nil
		} else {
			return fmt.Errorf("Error checking for Data-Source Connection (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
