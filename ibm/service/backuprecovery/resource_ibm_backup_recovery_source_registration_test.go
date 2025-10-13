// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package backuprecovery_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM/ibm-backup-recovery-sdk-go/backuprecoveryv1"
)

var (
	tenantIdRegister = "wkk1yqrdce/"
)

func TestAccIbmBackupRecoverySourceRegistrationBasic(t *testing.T) {
	var conf backuprecoveryv1.SourceRegistrationResponseParams

	environment := "kPhysical"
	connectionId := "5128356219792164864"
	endpoint := "172.26.202.5"
	hostType := "kLinux"
	physicalType := "kHost"
	applications := ""

	applicationsUpdate := `applications = ["kSQL"]`

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmBackupRecoverySourceRegistrationDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmBackupRecoverySourceRegistrationConfigBasic(environment, applications, endpoint, hostType, physicalType, connectionId),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmBackupRecoverySourceRegistrationExists("ibm_backup_recovery_source_registration.baas_source_registration_instance", conf),
					resource.TestCheckResourceAttr("ibm_backup_recovery_source_registration.baas_source_registration_instance", "x_ibm_tenant_id", tenantIdRegister),
					resource.TestCheckResourceAttr("ibm_backup_recovery_source_registration.baas_source_registration_instance", "environment", environment),
					resource.TestCheckResourceAttrSet("ibm_backup_recovery_source_registration.baas_source_registration_instance", "id"),
					resource.TestCheckResourceAttrSet("ibm_backup_recovery_source_registration.baas_source_registration_instance", "connections.#"),
					resource.TestCheckResourceAttr("ibm_backup_recovery_source_registration.baas_source_registration_instance", "connection_id", connectionId),
					resource.TestCheckResourceAttr("ibm_backup_recovery_source_registration.baas_source_registration_instance", "source_info.#", "1"),
					resource.TestCheckResourceAttr("ibm_backup_recovery_source_registration.baas_source_registration_instance", "source_info.0.name", endpoint),
					resource.TestCheckResourceAttr("ibm_backup_recovery_source_registration.baas_source_registration_instance", "source_info.0.os_type", hostType),
					resource.TestCheckResourceAttr("ibm_backup_recovery_source_registration.baas_source_registration_instance", "source_info.0.object_type", physicalType),
					resource.TestCheckResourceAttr("ibm_backup_recovery_source_registration.baas_source_registration_instance", "physical_params.#", "1"),
					resource.TestCheckResourceAttr("ibm_backup_recovery_source_registration.baas_source_registration_instance", "physical_params.0.endpoint", endpoint),
					resource.TestCheckResourceAttr("ibm_backup_recovery_source_registration.baas_source_registration_instance", "physical_params.0.host_type", hostType),
					resource.TestCheckResourceAttr("ibm_backup_recovery_source_registration.baas_source_registration_instance", "physical_params.0.physical_type", physicalType),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmBackupRecoverySourceRegistrationConfigBasic(environment, applicationsUpdate, endpoint, hostType, physicalType, connectionId),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_backup_recovery_source_registration.baas_source_registration_instance", "x_ibm_tenant_id", tenantIdRegister),
					resource.TestCheckResourceAttr("ibm_backup_recovery_source_registration.baas_source_registration_instance", "environment", environment),
				),
			},
		},
	})
}

func testAccCheckIbmBackupRecoverySourceRegistrationConfigBasic(environment, applications, endpoint, hostType, physicalType string, connectionId string) string {
	return fmt.Sprintf(`
			resource "ibm_backup_recovery_source_registration" "baas_source_registration_instance" {
				backup_recovery_endpoint = "https://protectiondomain0103.us-east.backup-recovery-tests.cloud.ibm.com/v2"
				x_ibm_tenant_id = "%s"
				environment = "%s"
				connection_id = "%s"
				physical_params {
				endpoint = "%s"
				host_type = "%s"
				physical_type = "%s"
				%s
				}
			}
	`, tenantIdRegister, environment, connectionId, endpoint, hostType, physicalType, applications)
}

func testAccCheckIbmBackupRecoverySourceRegistrationExists(n string, obj backuprecoveryv1.SourceRegistrationResponseParams) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}
		backupRecoveryClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).BackupRecoveryV1()
		if err != nil {
			return err
		}

		getProtectionSourceRegistrationOptions := &backuprecoveryv1.GetProtectionSourceRegistrationOptions{}
		num, _ := strconv.Atoi(rs.Primary.Attributes["source_id"])
		getProtectionSourceRegistrationOptions.SetID(int64(num))
		getProtectionSourceRegistrationOptions.SetXIBMTenantID(tenantIdRegister)

		sourceRegistrationReponseParams, _, err := backupRecoveryClient.GetProtectionSourceRegistration(getProtectionSourceRegistrationOptions)
		if err != nil {
			return err
		}

		obj = *sourceRegistrationReponseParams
		return nil
	}
}

func testAccCheckIbmBackupRecoverySourceRegistrationDestroy(s *terraform.State) error {
	backupRecoveryClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).BackupRecoveryV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_backup_recovery_source_registration" {
			continue
		}

		getProtectionSourceRegistrationOptions := &backuprecoveryv1.GetProtectionSourceRegistrationOptions{}

		num, _ := strconv.Atoi(rs.Primary.Attributes["source_id"])

		getProtectionSourceRegistrationOptions.SetXIBMTenantID(tenantIdRegister)
		getProtectionSourceRegistrationOptions.SetID(int64(num))

		// Try to find the key
		_, response, err := backupRecoveryClient.GetProtectionSourceRegistration(getProtectionSourceRegistrationOptions)

		if err == nil {
			return fmt.Errorf("baas_source_registration still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for baas_source_registration (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
