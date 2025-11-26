// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package backuprecovery_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM/ibm-backup-recovery-sdk-go/backuprecoveryv1"
)

func TestAccIbmBackupRecoveryDataSourceConnectorPatchBasic(t *testing.T) {
	connectorID := fmt.Sprintf("tf_connector_id_%d", acctest.RandIntRange(10, 100))
	xIbmTenantID := fmt.Sprintf("tf_x_ibm_tenant_id_%d", acctest.RandIntRange(10, 100))
	connectorName := fmt.Sprintf("tf_connector_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmBackupRecoveryDataSourceConnectorPatchConfigBasic(connectorID, xIbmTenantID, connectorName),
				Check:  resource.ComposeAggregateTestCheckFunc(),
			},
		},
	})
}

func testAccCheckIbmBackupRecoveryDataSourceConnectorPatchConfigBasic(connectorID string, xIbmTenantID, connectorName string) string {
	return fmt.Sprintf(`
	`)
}

func testAccCheckIbmBackupRecoveryDataSourceConnectorPatchExists(n string, obj backuprecoveryv1.DataSourceConnectorList) resource.TestCheckFunc {

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
