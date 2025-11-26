// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package backuprecovery_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM/ibm-backup-recovery-sdk-go/backuprecoveryv1"
)

func TestAccIbmBackupRecoveryPerformActionOnProtectionGroupRunRequestBasic(t *testing.T) {
	objectId := 344
	runType := "kRegular"
	groupName := "tetst-terra-group-2" //"tf-group-5"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,

		Steps: []resource.TestStep{
			{
				Destroy: false,
				Config:  testAccCheckIbmBackupRecoveryProtectionGroupRunRequest(groupName, runType, objectId),
				Check: resource.ComposeTestCheckFunc(
					testPerformRunExists("ibm_backup_recovery_protection_group_run_request.baas_protection_group_run_request_instance"),
				),
			},
			{
				Destroy: false,
				Config:  testAccCheckIbmBackupRecoveryPerformActionOnProtectionGroupRunRequestConfigBasic(objectId, groupName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckProtectionRunPerformActionCancelled("ibm_backup_recovery_perform_action_on_protection_group_run_request.baas_perform_action_on_protection_group_run_request_instance"),
				),
			},
		},
	})
}

func testAccCheckIbmBackupRecoveryProtectionGroupRunRequest(groupName, runType string, objectID int) string {
	return fmt.Sprintf(`

		data "ibm_backup_recovery_protection_groups" "ibm_backup_recovery_protection_groups_instance" {
			x_ibm_tenant_id = "%s"
			names = ["%s"]
			
		}

		resource "ibm_backup_recovery_protection_group_run_request" "baas_protection_group_run_request_instance" {
			x_ibm_tenant_id = "%s"
			run_type = "%s"
			
			group_id = data.ibm_backup_recovery_protection_groups.ibm_backup_recovery_protection_groups_instance.protection_groups.0.id
			lifecycle {
				ignore_changes = ["x_ibm_tenant_id","run_type","group_id"]
			}
		}
	`, tenantId, groupName, tenantId, runType)
}

func testPerformRunExists(n string) resource.TestCheckFunc {
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

			getProtectionGroupRunsOptions := &backuprecoveryv1.GetProtectionGroupRunsOptions{}
			getProtectionGroupRunsOptions.SetID(rs.Primary.ID)
			getProtectionGroupRunsOptions.SetXIBMTenantID(tenantId)
			performActionOnProtectionGroupRunResponse, _, err := backupRecoveryClient.GetProtectionGroupRuns(getProtectionGroupRunsOptions)
			if err != nil {
				return err
			}

			if performActionOnProtectionGroupRunResponse != nil &&
				len(performActionOnProtectionGroupRunResponse.Runs) > 0 &&
				*(performActionOnProtectionGroupRunResponse.Runs[0].ProtectionGroupID) == rs.Primary.ID &&
				len(performActionOnProtectionGroupRunResponse.Runs) > 0 &&
				performActionOnProtectionGroupRunResponse.Runs[0].ArchivalInfo != nil &&
				len(performActionOnProtectionGroupRunResponse.Runs[0].ArchivalInfo.ArchivalTargetResults) > 0 &&
				*(performActionOnProtectionGroupRunResponse.Runs[0].ArchivalInfo.ArchivalTargetResults[0].ArchivalTaskID) != "" {

				return nil
			}
			time.Sleep(15 * time.Second)
		}
		return nil
	}
}

func testAccCheckIbmBackupRecoveryPerformActionOnProtectionGroupRunRequestConfigBasic(objectId int, groupName string) string {
	return fmt.Sprintf(`

	data "ibm_backup_recovery_protection_groups" "ibm_backup_recovery_protection_groups_instance" {
		x_ibm_tenant_id = "%s"
		
		names = ["%s"]
	  }

	data "ibm_backup_recovery_protection_group_runs" "baas_protection_group_runs_instance" {
		x_ibm_tenant_id = "%s"
		
		protection_group_id = data.ibm_backup_recovery_protection_groups.ibm_backup_recovery_protection_groups_instance.protection_groups.0.id
	}

	resource "ibm_backup_recovery_perform_action_on_protection_group_run_request" "baas_perform_action_on_protection_group_run_request_instance" {
		x_ibm_tenant_id = "%s"
		
		group_id = data.ibm_backup_recovery_protection_group_runs.baas_protection_group_runs_instance.protection_group_id
		action = "Cancel"
		cancel_params {
		  run_id = data.ibm_backup_recovery_protection_group_runs.baas_protection_group_runs_instance.runs.0.id
		  local_task_id = data.ibm_backup_recovery_protection_group_runs.baas_protection_group_runs_instance.runs.0.archival_info.0.archival_target_results.0.archival_task_id
		}
		lifecycle {
			ignore_changes = ["x_ibm_tenant_id","group_id","action", "cancel_params"]
		}
	}
	`, tenantId, groupName, tenantId, tenantId)
}

func testAccCheckProtectionRunPerformActionCancelled(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		timeout := time.Now().Add(4 * time.Minute) // Set a 3-Minute timeout for waiting
		for time.Now().Before(timeout) {

			rs, ok := s.RootModule().Resources[n]
			if !ok {
				return fmt.Errorf("Not found: %s", n)
			}
			runId := rs.Primary.Attributes["cancel_params.0.run_id"]
			groupId := rs.Primary.Attributes["group_id"]

			backupRecoveryClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).BackupRecoveryV1()
			if err != nil {
				return err
			}

			getProtectionGroupRunsOptions := &backuprecoveryv1.GetProtectionGroupRunsOptions{}

			getProtectionGroupRunsOptions.SetID(groupId)
			getProtectionGroupRunsOptions.SetXIBMTenantID(tenantId)
			getProtectionGroupRunsOptions.SetRunID(runId)

			performActionOnProtectionGroupRunResponse, _, err := backupRecoveryClient.GetProtectionGroupRuns(getProtectionGroupRunsOptions)
			if err != nil {
				return err
			}

			if performActionOnProtectionGroupRunResponse != nil &&
				len(performActionOnProtectionGroupRunResponse.Runs) > 0 &&
				*(performActionOnProtectionGroupRunResponse.Runs[0].ProtectionGroupID) == groupId &&
				len(performActionOnProtectionGroupRunResponse.Runs) > 0 &&
				performActionOnProtectionGroupRunResponse.Runs[0].ArchivalInfo != nil &&
				len(performActionOnProtectionGroupRunResponse.Runs[0].ArchivalInfo.ArchivalTargetResults) > 0 &&
				*(performActionOnProtectionGroupRunResponse.Runs[0].ArchivalInfo.ArchivalTargetResults[0].Status) == "Canceled" {

				return nil
			}
			time.Sleep(15 * time.Second)
		}
		return nil
	}
}
