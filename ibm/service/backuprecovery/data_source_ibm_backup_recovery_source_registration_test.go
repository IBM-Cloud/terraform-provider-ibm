// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.94.0-fa797aec-20240814-142622
 */

package backuprecovery_test

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmBackupRecoverySourceRegistrationDataSourceBasic(t *testing.T) {
	// environment := "kPhysical"
	objectId := 344
	// endpoint := "172.26.1.24"
	// hostType := "kLinux"
	// physicalType := "kHost"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmBackupRecoverySourceRegistrationDataSourceConfigBasic(objectId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_backup_recovery_source_registration.baas_source_registration_instance", "x_ibm_tenant_id", tenantId),
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_source_registration.baas_source_registration_instance", "environment"),
					resource.TestCheckResourceAttr("data.ibm_backup_recovery_source_registration.baas_source_registration_instance", "id", fmt.Sprintf("%s::%d", tenantId, objectId)),
					resource.TestCheckResourceAttr("data.ibm_backup_recovery_source_registration.baas_source_registration_instance", "source_id", strconv.Itoa(objectId)),
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_source_registration.baas_source_registration_instance", "last_refreshed_time_msecs"),
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_source_registration.baas_source_registration_instance", "source_registration_id"),
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_source_registration.baas_source_registration_instance", "connections.#"),
					resource.TestCheckResourceAttr("data.ibm_backup_recovery_source_registration.baas_source_registration_instance", "source_info.#", "1"),
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_source_registration.baas_source_registration_instance", "source_info.0.name"),
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_source_registration.baas_source_registration_instance", "source_info.0.id"),
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_source_registration.baas_source_registration_instance", "source_info.0.os_type"),
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_source_registration.baas_source_registration_instance", "source_info.0.object_type"),
					resource.TestCheckResourceAttr("data.ibm_backup_recovery_source_registration.baas_source_registration_instance", "physical_params.#", "1"),
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_source_registration.baas_source_registration_instance", "physical_params.0.endpoint"),
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_source_registration.baas_source_registration_instance", "physical_params.0.host_type"),
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_source_registration.baas_source_registration_instance", "physical_params.0.physical_type"),
				),
			},
		},
	})
}

func testAccCheckIbmBackupRecoverySourceRegistrationDataSourceConfigBasic(objectId int) string {
	return fmt.Sprintf(`

			data "ibm_backup_recovery_source_registration" "baas_source_registration_instance" {
				backup_recovery_endpoint = "https://protectiondomain0103.us-east.backup-recovery-tests.cloud.ibm.com/v2"
				source_registration_id = %d
				x_ibm_tenant_id = "%s"
			}
	`, objectId, tenantId)
}
