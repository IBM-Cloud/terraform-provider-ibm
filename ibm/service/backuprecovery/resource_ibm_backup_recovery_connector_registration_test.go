// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package backuprecovery_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmBackupRecoveryConnectorRegistrationBasic(t *testing.T) {
	// var conf backuprecoveryv1.DataSourceConnectorRegistrationRequest
	xIbmTenantID := fmt.Sprintf("tf_x_ibm_tenant_id_%d", acctest.RandIntRange(10, 100))
	registrationToken := fmt.Sprintf("tf_registration_token_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Destroy: false,
				Config:  testAccCheckIbmBackupRecoveryConnectorRegistrationConfigBasic(xIbmTenantID, registrationToken),
				Check: resource.ComposeAggregateTestCheckFunc(
					// testAccCheckIbmBackupRecoveryConnectorRegistrationExists("ibm_backup_recovery_connector_registration.backup_recovery_connector_registration_instance", conf),
					resource.TestCheckResourceAttr("ibm_backup_recovery_connector_registration.backup_recovery_connector_registration_instance", "x_ibm_tenant_id", xIbmTenantID),
					resource.TestCheckResourceAttr("ibm_backup_recovery_connector_registration.backup_recovery_connector_registration_instance", "registration_token", registrationToken),
				),
			},
		},
	})
}

func testAccCheckIbmBackupRecoveryConnectorRegistrationConfigBasic(xIbmTenantID string, registrationToken string) string {
	return fmt.Sprintf(`
		resource "ibm_backup_recovery_connector_registration" "backup_recovery_connector_registration_instance" {
			x_ibm_tenant_id = "%s"
			registration_token = "%s"
		}
	`, xIbmTenantID, registrationToken)
}

func testAccCheckIbmBackupRecoveryConnectorRegistrationConfig(xIbmTenantID string, connectorID string, registrationToken string) string {
	return fmt.Sprintf(`

		resource "ibm_backup_recovery_connector_registration" "backup_recovery_connector_registration_instance" {
			x_ibm_tenant_id = "%s"
			connector_id = %s
			registration_token = "%s"
		}
	`, xIbmTenantID, connectorID, registrationToken)
}
