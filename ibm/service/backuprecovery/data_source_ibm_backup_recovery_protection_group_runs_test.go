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

func TestAccIbmBackupRecoveryProtectionGroupRunsDataSourceBasic(t *testing.T) {
	groupName := "tetst-terra-group-2"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmBackupRecoveryProtectionGroupRunsDataSourceConfigBasic(groupName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_protection_group_runs.baas_protection_group_runs_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_protection_group_runs.baas_protection_group_runs_instance", "x_ibm_tenant_id"),
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_protection_group_runs.baas_protection_group_runs_instance", "runs.#"),
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_protection_group_runs.baas_protection_group_runs_instance", "total_runs"),
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_protection_group_runs.baas_protection_group_runs_instance", "runs.#"),
					resource.TestCheckResourceAttr("data.ibm_backup_recovery_protection_group_runs.baas_protection_group_runs_instance", "runs.0.protection_group_name", groupName),
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_protection_group_runs.baas_protection_group_runs_instance", "runs.0.id"),
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_protection_group_runs.baas_protection_group_runs_instance", "runs.0.protection_group_id"),
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_protection_group_runs.baas_protection_group_runs_instance", "runs.0.protection_group_instance_id"),
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_protection_group_runs.baas_protection_group_runs_instance", "runs.0.permissions.#"),
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_protection_group_runs.baas_protection_group_runs_instance", "runs.0.archival_info.#"),
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_protection_group_runs.baas_protection_group_runs_instance", "runs.0.is_cloud_archival_direct"),
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_protection_group_runs.baas_protection_group_runs_instance", "runs.0.is_local_snapshots_deleted"),
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_protection_group_runs.baas_protection_group_runs_instance", "runs.0.environment"),
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_protection_group_runs.baas_protection_group_runs_instance", "runs.0.is_replication_run"),
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_protection_group_runs.baas_protection_group_runs_instance", "runs.0.on_legal_hold"),
				),
			},
		},
	})
}

func testAccCheckIbmBackupRecoveryProtectionGroupRunsDataSourceConfigBasic(groupName string) string {
	return fmt.Sprintf(`
	data "ibm_backup_recovery_protection_groups" "ibm_backup_recovery_protection_groups_instance" {
		x_ibm_tenant_id = "%s"
		
		names = ["%s"]
	}

	data "ibm_backup_recovery_protection_group_runs" "baas_protection_group_runs_instance" {
		x_ibm_tenant_id = "%[1]s"
		
		protection_group_id = data.ibm_backup_recovery_protection_groups.ibm_backup_recovery_protection_groups_instance.protection_groups.0.id
	}
	`, tenantId, groupName)
}
