// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package backuprecovery_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmBackupRecoveryConnectionRegistrationTokenBasic(t *testing.T) {
	connectionName := fmt.Sprintf("tf_connection_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmBackupRecoveryConnectionRegistrationTokenConfigBasic(connectionName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("ibm_backup_recovery_connection_registration_token.baas_connection_registration_token_instance", "connection_id"),
					resource.TestCheckResourceAttr("ibm_backup_recovery_connection_registration_token.baas_connection_registration_token_instance", "x_ibm_tenant_id", tenantId),
					resource.TestCheckResourceAttrSet("ibm_backup_recovery_connection_registration_token.baas_connection_registration_token_instance", "registration_token"),
				),
				Destroy: false,
			},
		},
	})
}

func testAccCheckIbmBackupRecoveryConnectionRegistrationTokenConfigBasic(connectionName string) string {
	return fmt.Sprintf(`

		resource "ibm_backup_recovery_data_source_connection" "baas_data_source_connection_instance_1" {
			x_ibm_tenant_id = "%s"
			connection_name = "%s"
		}
		resource "ibm_backup_recovery_connection_registration_token" "baas_connection_registration_token_instance" {
			connection_id = ibm_backup_recovery_data_source_connection.baas_data_source_connection_instance_1.connection_id
			x_ibm_tenant_id = "%s"
		}
	`, tenantId, connectionName, tenantId)
}
