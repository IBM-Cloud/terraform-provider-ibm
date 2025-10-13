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

func TestAccIbmBackupRecoveryProtectionGroupRunRequestBasic(t *testing.T) {
	objectId := 344
	runType := "kRegular"
	groupName := "tetst-terra-group-2" //"tf-group-5"

	resource.Test(t, resource.TestCase{
		PreCheck:                  func() { acc.TestAccPreCheck(t) },
		Providers:                 acc.TestAccProviders,
		CheckDestroy:              func(s *terraform.State) error { return nil },
		PreventPostDestroyRefresh: true,
		Steps: []resource.TestStep{
			{
				Destroy: false,
				Config:  testAccCreateIbmBaasProtectionGroupRunRequest(groupName, runType, objectId),
				Check: resource.ComposeTestCheckFunc(
					testRunExists("ibm_backup_recovery_protection_group_run_request.baas_protection_group_run_request_instance"),
				),
			},
			{
				Destroy: false,
				Config:  testAccCreateIbmBaasProtectionGroupRunCancelRequestConfigBasic(runType, groupName, objectId),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckProtectionRunCancelled("ibm_backup_recovery_perform_action_on_protection_group_run_request.baas_perform_action_on_protection_group_run_request_instance"),
				),
			},
		},
	})
}

func testAccCreateIbmBaasProtectionGroupRunRequest(groupName, runType string, objectID int) string {
	return fmt.Sprintf(`

		data "ibm_backup_recovery_protection_groups" "ibm_backup_recovery_protection_groups_instance" {
			backup_recovery_endpoint = "https://protectiondomain0103.us-east.backup-recovery-tests.cloud.ibm.com/v2"
			x_ibm_tenant_id = "%s"
			names = ["%s"]
		}

		resource "ibm_backup_recovery_protection_group_run_request" "baas_protection_group_run_request_instance" {
			x_ibm_tenant_id = "%s"
			run_type = "%s"
			backup_recovery_endpoint = "https://protectiondomain0103.us-east.backup-recovery-tests.cloud.ibm.com/v2"
			group_id = data.ibm_backup_recovery_protection_groups.ibm_backup_recovery_protection_groups_instance.protection_groups.0.id
			lifecycle {
				ignore_changes = ["x_ibm_tenant_id","run_type","group_id"]
			}
		}
	`, tenantId, groupName, tenantId, runType)
}

func testRunExists(n string) resource.TestCheckFunc {
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

func testAccCreateIbmBaasProtectionGroupRunCancelRequestConfigBasic(runType, groupName string, objectID int) string {
	return fmt.Sprintf(`
	data "ibm_backup_recovery_protection_groups" "baas_protection_group_existing_instance" {
		x_ibm_tenant_id = "%[1]s"
		backup_recovery_endpoint = "https://protectiondomain0103.us-east.backup-recovery-tests.cloud.ibm.com/v2"
		names = ["%[2]s"]
	}

	data "ibm_backup_recovery_protection_group_runs" "example_runs" {
		backup_recovery_endpoint = "https://protectiondomain0103.us-east.backup-recovery-tests.cloud.ibm.com/v2"
		x_ibm_tenant_id = "%[1]s"
		protection_group_id = data.ibm_backup_recovery_protection_groups.baas_protection_group_existing_instance.protection_groups.0.id
	}

	resource "ibm_backup_recovery_perform_action_on_protection_group_run_request" "baas_perform_action_on_protection_group_run_request_instance" {
		x_ibm_tenant_id = "%[1]s"
		backup_recovery_endpoint = "https://protectiondomain0103.us-east.backup-recovery-tests.cloud.ibm.com/v2"
		group_id = data.ibm_backup_recovery_protection_groups.baas_protection_group_existing_instance.protection_groups.0.id
		action = "Cancel"
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

func testAccCheckIbmBackupRecoveryProtectionGroupRunRequestExists(n string, obj backuprecoveryv1.ProtectionGroupRunsResponse) resource.TestCheckFunc {

	return func(s *terraform.State) error {
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
		protectionGroupRunResponse, _, err := backupRecoveryClient.GetProtectionGroupRuns(getProtectionGroupRunsOptions)
		if err != nil {
			return err
		}

		if len(protectionGroupRunResponse.Runs) > 0 {
			return nil
		}
		return nil
	}
}

func testAccCheckProtectionRunCancelled(n string) resource.TestCheckFunc {
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
			backupRecoveryClient.Service.Options.URL = "https://protectiondomain0103.us-east.backup-recovery-tests.cloud.ibm.com/v2"

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
