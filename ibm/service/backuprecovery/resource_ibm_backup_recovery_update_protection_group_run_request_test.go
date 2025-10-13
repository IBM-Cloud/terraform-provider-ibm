// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package backuprecovery_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM/ibm-backup-recovery-sdk-go/backuprecoveryv1"
)

func TestAccIbmBackupRecoveryUpdateProtectionGroupRunRequestBasic(t *testing.T) {
	objectId := 344
	runType := "kRegular"
	groupName := "tetst-terra-group-2" // or can use "tf-group-5" //id: 5901263190628181:1725393921826:9414
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Destroy: false,
				Config:  testAccCreateIbmBaasProtectionGroupForUpdateRunRequest(groupName, runType, objectId),
				Check: resource.ComposeTestCheckFunc(
					testUpdateRunExists("ibm_backup_recovery_protection_group_run_request.baas_protection_group_run_request_instance"),
				),
			},
			{
				Destroy: false,
				Config:  testAccCreateIbmBaasProtectionGroupRunUpdateRequestConfigBasic(runType, groupName, objectId),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttrSet("ibm_backup_recovery_update_protection_group_run_request.baas_update_protection_group_run_request_instance", "group_id"),
					resource.TestCheckResourceAttr("ibm_backup_recovery_update_protection_group_run_request.baas_update_protection_group_run_request_instance", "update_protection_group_run_params.#", "1"),
				),
			},
			{
				Destroy: false,
				Config:  testAccCreateIbmBaasProtectionGroupRunCancelUpdateRequestConfigBasic(runType, groupName, objectId),
				Check:   resource.ComposeAggregateTestCheckFunc(),
			},
		},
	})
}

func testUpdateRunExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		timeout := time.Now().Add(4 * time.Minute) // Set a 3-Minute timeout for waiting
		for time.Now().Before(timeout) {

			rs, ok := s.RootModule().Resources[n]
			if !ok {
				return fmt.Errorf("Not found: %s", n)
			}
			backupRecoveryClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).BackupRecoveryV1()
			if err != nil {
				return err
			}
			backupRecoveryClient.Service.Options.URL = "https://protectiondomain0103.us-east.backup-recovery-tests.cloud.ibm.com/v2"

			getProtectionGroupRunsOptions := &backuprecoveryv1.GetProtectionGroupRunsOptions{}
			getProtectionGroupRunsOptions.SetID(rs.Primary.ID)
			getProtectionGroupRunsOptions.SetXIBMTenantID(tenantId)
			GetProtectionGroupRunsResponse, _, err := backupRecoveryClient.GetProtectionGroupRuns(getProtectionGroupRunsOptions)
			if err != nil {
				return err
			}

			if GetProtectionGroupRunsResponse != nil &&
				len(GetProtectionGroupRunsResponse.Runs) > 0 &&
				*(GetProtectionGroupRunsResponse.Runs[0].ProtectionGroupID) == rs.Primary.ID &&
				len(GetProtectionGroupRunsResponse.Runs) > 0 &&
				GetProtectionGroupRunsResponse.Runs[0].ArchivalInfo != nil &&
				len(GetProtectionGroupRunsResponse.Runs[0].ArchivalInfo.ArchivalTargetResults) > 0 &&
				*(GetProtectionGroupRunsResponse.Runs[0].ArchivalInfo.ArchivalTargetResults[0].ArchivalTaskID) != "" {

				return nil
			}
			time.Sleep(15 * time.Second)
		}
		return nil
	}
}

func testAccCreateIbmBaasProtectionGroupForUpdateRunRequest(groupName, runType string, objectID int) string {
	return fmt.Sprintf(`

		data "ibm_backup_recovery_protection_groups" "ibm_backup_recovery_protection_groups_instance" {
			x_ibm_tenant_id = "%s"
			backup_recovery_endpoint = "https://protectiondomain0103.us-east.backup-recovery-tests.cloud.ibm.com/v2"
			names = ["%s"]
		}

		resource "ibm_backup_recovery_protection_group_run_request" "baas_protection_group_run_request_instance" {
			backup_recovery_endpoint = "https://protectiondomain0103.us-east.backup-recovery-tests.cloud.ibm.com/v2"
			x_ibm_tenant_id = "%s"
			run_type = "%s"
			group_id = data.ibm_backup_recovery_protection_groups.ibm_backup_recovery_protection_groups_instance.protection_groups.0.id
			lifecycle {
				ignore_changes = ["x_ibm_tenant_id","run_type","group_id"]
			}
		}
	`, tenantId, groupName, tenantId, runType)
}

func testAccCreateIbmBaasProtectionGroupRunUpdateRequestConfigBasic(runType, groupName string, objectID int) string {
	return fmt.Sprintf(`
		data "ibm_backup_recovery_protection_groups" "baas_protection_group_existing_instance" {
			x_ibm_tenant_id = "%[1]s"
			backup_recovery_endpoint = "https://protectiondomain0103.us-east.backup-recovery-tests.cloud.ibm.com/v2"
			names = ["%[2]s"]
		}

		data "ibm_backup_recovery_protection_group_runs" "example_runs" {
			x_ibm_tenant_id = "%[1]s"
			backup_recovery_endpoint = "https://protectiondomain0103.us-east.backup-recovery-tests.cloud.ibm.com/v2"
			protection_group_id = data.ibm_backup_recovery_protection_groups.baas_protection_group_existing_instance.protection_groups.0.id
		}

		resource "ibm_backup_recovery_update_protection_group_run_request" "baas_update_protection_group_run_request_instance" {
			x_ibm_tenant_id = "%[1]s"
			backup_recovery_endpoint = "https://protectiondomain0103.us-east.backup-recovery-tests.cloud.ibm.com/v2"
			group_id = data.ibm_backup_recovery_protection_groups.baas_protection_group_existing_instance.protection_groups.0.id
			update_protection_group_run_params {
				run_id = data.ibm_backup_recovery_protection_group_runs.example_runs.runs.0.id
				local_snapshot_config {
					delete_snapshot = false
				}
			}
		}	
	`, tenantId, groupName)
}

func testAccCreateIbmBaasProtectionGroupRunCancelUpdateRequestConfigBasic(runType, groupName string, objectID int) string {
	return fmt.Sprintf(`
	data "ibm_backup_recovery_protection_groups" "baas_protection_group_existing_instance" {
		x_ibm_tenant_id = "%[1]s"
		backup_recovery_endpoint = "https://protectiondomain0103.us-east.backup-recovery-tests.cloud.ibm.com/v2"
		names = ["%[2]s"]
	}

	data "ibm_backup_recovery_protection_group_runs" "example_runs" {
		x_ibm_tenant_id = "%[1]s"
		backup_recovery_endpoint = "https://protectiondomain0103.us-east.backup-recovery-tests.cloud.ibm.com/v2"
		protection_group_id = data.ibm_backup_recovery_protection_groups.baas_protection_group_existing_instance.protection_groups.0.id
	}

	resource "ibm_backup_recovery_perform_action_on_protection_group_run_request" "baas_perform_action_on_updated_protection_group_run_request_instance" {
		x_ibm_tenant_id = "%[1]s"
		group_id = data.ibm_backup_recovery_protection_groups.baas_protection_group_existing_instance.protection_groups.0.id
		action = "Cancel"
		backup_recovery_endpoint = "https://protectiondomain0103.us-east.backup-recovery-tests.cloud.ibm.com/v2"
		cancel_params {
			run_id = data.ibm_backup_recovery_protection_group_runs.example_runs.runs.0.id
			local_task_id = data.ibm_backup_recovery_protection_group_runs.example_runs.runs.0.archival_info.0.archival_target_results.0.archival_task_id
		  }
		lifecycle {
			ignore_changes = ["x_ibm_tenant_id","group_id","action", "cancel_params"]
		}
	  }
	`, tenantId, groupName)
}
