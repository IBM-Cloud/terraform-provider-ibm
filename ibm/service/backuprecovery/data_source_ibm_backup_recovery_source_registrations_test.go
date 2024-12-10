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

func TestAccIbmBackupRecoverySourceRegistrationsDataSourceBasic(t *testing.T) {
	// environment := "kPhysical"
	// connectionId := "4980716806983529472"
	// endpoint := "172.26.1.24"
	// hostType := "kLinux"
	// physicalType := "kHost"

	objectId := 18
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmBackupRecoverySourceRegistrationsDataSourceConfigBasic(objectId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_source_registrations.baas_source_registrations_instance", "id"),
					resource.TestCheckResourceAttr("data.ibm_backup_recovery_source_registrations.baas_source_registrations_instance", "x_ibm_tenant_id", tenantId),
					resource.TestCheckResourceAttr("data.ibm_backup_recovery_source_registrations.baas_source_registrations_instance", "ids.#", "1"),
					resource.TestCheckResourceAttr("data.ibm_backup_recovery_source_registrations.baas_source_registrations_instance", "registrations.#", "1"),
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_source_registrations.baas_source_registrations_instance", "registrations.0.id"),
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_source_registrations.baas_source_registrations_instance", "registrations.0.source_id"),
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_source_registrations.baas_source_registrations_instance", "registrations.0.environment"),
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_source_registrations.baas_source_registrations_instance", "registrations.0.connection_id"),
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_source_registrations.baas_source_registrations_instance", "registrations.0.connector_group_id"),
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_source_registrations.baas_source_registrations_instance", "registrations.0.data_source_connection_id"),
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_source_registrations.baas_source_registrations_instance", "registrations.0.authentication_status"),
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_source_registrations.baas_source_registrations_instance", "registrations.0.registration_time_msecs"),
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_source_registrations.baas_source_registrations_instance", "registrations.0.last_refreshed_time_msecs"),
					resource.TestCheckResourceAttr("data.ibm_backup_recovery_source_registrations.baas_source_registrations_instance", "registrations.0.source_info.#", "1"),
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_source_registrations.baas_source_registrations_instance", "registrations.0.source_info.0.name"),
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_source_registrations.baas_source_registrations_instance", "registrations.0.source_info.0.id"),
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_source_registrations.baas_source_registrations_instance", "registrations.0.source_info.0.os_type"),
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_source_registrations.baas_source_registrations_instance", "registrations.0.source_info.0.object_type"),
					resource.TestCheckResourceAttr("data.ibm_backup_recovery_source_registrations.baas_source_registrations_instance", "registrations.0.physical_params.#", "1"),
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_source_registrations.baas_source_registrations_instance", "registrations.0.physical_params.0.endpoint"),
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_source_registrations.baas_source_registrations_instance", "registrations.0.physical_params.0.host_type"),
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_source_registrations.baas_source_registrations_instance", "registrations.0.physical_params.0.physical_type"),
				),
			},
		},
	})
}

func testAccCheckIbmBackupRecoverySourceRegistrationsDataSourceConfigBasic(objectId int) string {
	return fmt.Sprintf(`

			data "ibm_backup_recovery_source_registrations" "baas_source_registrations_instance" {
				ids = [%d]
				x_ibm_tenant_id = "%s"
			}
	`, objectId, tenantId)
}
