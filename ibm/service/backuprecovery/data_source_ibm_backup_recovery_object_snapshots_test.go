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

func TestAccIbmBackupRecoveryObjectSnapshotsDataSourceBasic(t *testing.T) {
	objectId := 18
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmBackupRecoveryObjectSnapshotsDataSourceConfigBasic(objectId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_object_snapshots.baas_object_snapshots_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_object_snapshots.baas_object_snapshots_instance", "snapshots.#"),
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_object_snapshots.baas_object_snapshots_instance", "x_ibm_tenant_id"),
					resource.TestCheckResourceAttr("data.ibm_backup_recovery_object_snapshots.baas_object_snapshots_instance", "snapshots.0.object_id", strconv.Itoa(objectId)),
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_object_snapshots.baas_object_snapshots_instance", "snapshots.0.id"),
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_object_snapshots.baas_object_snapshots_instance", "snapshots.0.storage_domain_id"),
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_object_snapshots.baas_object_snapshots_instance", "snapshots.0.source_id"),
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_object_snapshots.baas_object_snapshots_instance", "snapshots.0.snapshot_target_type"),
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_object_snapshots.baas_object_snapshots_instance", "snapshots.0.run_type"),
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_object_snapshots.baas_object_snapshots_instance", "snapshots.0.protection_group_run_id"),
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_object_snapshots.baas_object_snapshots_instance", "snapshots.0.protection_group_id"),
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_object_snapshots.baas_object_snapshots_instance", "snapshots.0.run_instance_id"),
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_object_snapshots.baas_object_snapshots_instance", "snapshots.0.protection_group_name"),
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_object_snapshots.baas_object_snapshots_instance", "snapshots.0.ownership_context"),
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_object_snapshots.baas_object_snapshots_instance", "snapshots.0.object_name"),
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_object_snapshots.baas_object_snapshots_instance", "snapshots.0.has_data_lock"),
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_object_snapshots.baas_object_snapshots_instance", "snapshots.0.environment"),
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_object_snapshots.baas_object_snapshots_instance", "snapshots.0.cluster_id"),
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_object_snapshots.baas_object_snapshots_instance", "snapshots.0.expiry_time_usecs"),
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_object_snapshots.baas_object_snapshots_instance", "snapshots.0.run_start_time_usecs"),
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_object_snapshots.baas_object_snapshots_instance", "snapshots.0.snapshot_timestamp_usecs"),
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_object_snapshots.baas_object_snapshots_instance", "snapshots.0.indexing_status"),
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_object_snapshots.baas_object_snapshots_instance", "snapshots.0.external_target_info.#"),
				),
			},
		},
	})
}

func testAccCheckIbmBackupRecoveryObjectSnapshotsDataSourceConfigBasic(objectId int) string {
	return fmt.Sprintf(`
	data "ibm_backup_recovery_object_snapshots" "baas_object_snapshots_instance" {
		x_ibm_tenant_id = "%s"
		object_id = %d
	  }
	`, tenantId, objectId)
}
