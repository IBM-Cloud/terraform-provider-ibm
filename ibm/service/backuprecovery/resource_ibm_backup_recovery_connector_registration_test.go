// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package backuprecovery_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.ibm.com/BackupAndRecovery/ibm-backup-recovery-sdk-go/backuprecoveryv1"
)

func TestAccIbmBackupRecoveryConnectorRegistrationBasic(t *testing.T) {
	// var conf backuprecoveryv1.DataSourceConnectorRegistrationRequest
	xIbmTenantID := fmt.Sprintf("tf_x_ibm_tenant_id_%d", acctest.RandIntRange(10, 100))
	registrationToken := fmt.Sprintf("tf_registration_token_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmBackupRecoveryConnectorRegistrationDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmBackupRecoveryConnectorRegistrationConfigBasic(xIbmTenantID, registrationToken),
				Check: resource.ComposeAggregateTestCheckFunc(
					// testAccCheckIbmBackupRecoveryConnectorRegistrationExists("ibm_backup_recovery_connector_registration.backup_recovery_connector_registration_instance", conf),
					resource.TestCheckResourceAttr("ibm_backup_recovery_connector_registration.backup_recovery_connector_registration_instance", "x_ibm_tenant_id", xIbmTenantID),
					resource.TestCheckResourceAttr("ibm_backup_recovery_connector_registration.backup_recovery_connector_registration_instance", "registration_token", registrationToken),
				),
			},
		},
	})
}

func TestAccIbmBackupRecoveryConnectorRegistrationAllArgs(t *testing.T) {
	// var conf backuprecoveryv1.DataSourceConnectorRegistrationRequest
	// xIbmTenantID := fmt.Sprintf("tf_x_ibm_tenant_id_%d", acctest.RandIntRange(10, 100))
	// connectorID := fmt.Sprintf("%d", acctest.RandIntRange(10, 100))
	// registrationToken := fmt.Sprintf("tf_registration_token_%d", acctest.RandIntRange(10, 100))

	// resource.Test(t, resource.TestCase{
	// 	PreCheck:     func() { acc.TestAccPreCheck(t) },
	// 	Providers:    acc.TestAccProviders,
	// 	CheckDestroy: testAccCheckIbmBackupRecoveryConnectorRegistrationDestroy,
	// 	Steps: []resource.TestStep{
	// 		resource.TestStep{
	// 			Config: testAccCheckIbmBackupRecoveryConnectorRegistrationConfig(xIbmTenantID, connectorID, registrationToken),
	// 			Check: resource.ComposeAggregateTestCheckFunc(
	// 				testAccCheckIbmBackupRecoveryConnectorRegistrationExists("ibm_backup_recovery_connector_registration.backup_recovery_connector_registration_instance", conf),
	// 				resource.TestCheckResourceAttr("ibm_backup_recovery_connector_registration.backup_recovery_connector_registration_instance", "x_ibm_tenant_id", xIbmTenantID),
	// 				resource.TestCheckResourceAttr("ibm_backup_recovery_connector_registration.backup_recovery_connector_registration_instance", "connector_id", connectorID),
	// 				resource.TestCheckResourceAttr("ibm_backup_recovery_connector_registration.backup_recovery_connector_registration_instance", "registration_token", registrationToken),
	// 			),
	// 		},
	// 		resource.TestStep{
	// 			ResourceName:      "ibm_backup_recovery_connector_registration.backup_recovery_connector_registration",
	// 			ImportState:       true,
	// 			ImportStateVerify: true,
	// 		},
	// 	},
	// })
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

func testAccCheckIbmBackupRecoveryConnectorRegistrationExists(n string, obj backuprecoveryv1.DataSourceConnectorRegistrationStatus) resource.TestCheckFunc {

	// return func(s *terraform.State) error {
	// 	_, ok := s.RootModule().Resources[n]
	// 	if !ok {
	// 		return fmt.Errorf("Not found: %s", n)
	// 	}

	// 	backupRecoveryClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).BackupRecoveryV1()
	// 	if err != nil {
	// 		return err
	// 	}

	// 	registerDataSourceConnectorOptions := &backuprecoveryv1.RegisterDataSourceConnectorOptions{}

	// 	dataSourceConnectorRegistrationRequest, _, err := backupRecoveryClient.RegisterDataSourceConnector(registerDataSourceConnectorOptions)
	// 	if err != nil {
	// 		return err
	// 	}

	// 	obj = *dataSourceConnectorRegistrationRequest
	// 	return nil
	// }

	return nil
}

func testAccCheckIbmBackupRecoveryConnectorRegistrationDestroy(s *terraform.State) error {
	// backupRecoveryClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).BackupRecoveryV1()
	// if err != nil {
	// 	return err
	// }
	// for _, rs := range s.RootModule().Resources {
	// 	if rs.Type != "ibm_backup_recovery_connector_registration" {
	// 		continue
	// 	}

	// 	registerDataSourceConnectorOptions := &backuprecoveryv1.RegisterDataSourceConnectorOptions{}

	// 	// Try to find the key
	// 	_, response, err := backupRecoveryClient.RegisterDataSourceConnector(registerDataSourceConnectorOptions)

	// 	if err == nil {
	// 		return fmt.Errorf("Data-Source Connector Registration Request still exists: %s", rs.Primary.ID)
	// 	} else if response.StatusCode != 404 {
	// 		return fmt.Errorf("Error checking for Data-Source Connector Registration Request (%s) has been destroyed: %s", rs.Primary.ID, err)
	// 	}
	// }

	return nil
}
