// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package brsmigration_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/brsmigration"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/stretchr/testify/assert"
	"github.ibm.com/BackupAndRecovery/brs-migration-orchestrator/brsmigrationv2"
	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmBrsMigrationWorkloadRunBasic(t *testing.T) {
	var conf brsmigrationv2.WorkloadRun
	migrationID := fmt.Sprintf("tf_migration_id_%d", acctest.RandIntRange(10, 100))
	workloadID := fmt.Sprintf("tf_workload_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmBrsMigrationWorkloadRunDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmBrsMigrationWorkloadRunConfigBasic(migrationID, workloadID),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmBrsMigrationWorkloadRunExists("ibm_brs_migration_workload_run.brs_migration_workload_run_instance", conf),
					resource.TestCheckResourceAttr("ibm_brs_migration_workload_run.brs_migration_workload_run_instance", "migration_id", migrationID),
					resource.TestCheckResourceAttr("ibm_brs_migration_workload_run.brs_migration_workload_run_instance", "workload_id", workloadID),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_brs_migration_workload_run.brs_migration_workload_run_instance",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmBrsMigrationWorkloadRunConfigBasic(migrationID string, workloadID string) string {
	return fmt.Sprintf(`
		resource "ibm_brs_migration_workload_run" "brs_migration_workload_run_instance" {
			migration_id = "%s"
			workload_id = "%s"
		}
	`, migrationID, workloadID)
}

func testAccCheckIbmBrsMigrationWorkloadRunExists(n string, obj brsmigrationv2.WorkloadRun) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		brsMigrationClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).BrsMigrationV2()
		if err != nil {
			return err
		}

		getWorkloadRunOptions := &brsmigrationv2.GetWorkloadRunOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getWorkloadRunOptions.SetMigrationID(parts[0])
		getWorkloadRunOptions.SetWorkloadID(parts[1])
		getWorkloadRunOptions.SetRunID(parts[2])

		workloadRun, _, err := brsMigrationClient.GetWorkloadRun(getWorkloadRunOptions)
		if err != nil {
			return err
		}

		obj = *workloadRun
		return nil
	}
}

func testAccCheckIbmBrsMigrationWorkloadRunDestroy(s *terraform.State) error {
	brsMigrationClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).BrsMigrationV2()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_brs_migration_workload_run" {
			continue
		}

		getWorkloadRunOptions := &brsmigrationv2.GetWorkloadRunOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getWorkloadRunOptions.SetMigrationID(parts[0])
		getWorkloadRunOptions.SetWorkloadID(parts[1])
		getWorkloadRunOptions.SetRunID(parts[2])

		// Try to find the key
		_, response, err := brsMigrationClient.GetWorkloadRun(getWorkloadRunOptions)

		if err == nil {
			return fmt.Errorf("brs_migration_workload_run still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for brs_migration_workload_run (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

func TestResourceIbmBrsMigrationWorkloadRunWorkloadRunStatsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["logical_size_bytes"] = int(0)
		model["bytes_transferred"] = int(0)
		model["bytes_read"] = int(0)
		model["total_file_count"] = int(0)
		model["transferred_file_count"] = int(0)

		assert.Equal(t, result, model)
	}

	model := new(brsmigrationv2.WorkloadRunStats)
	model.LogicalSizeBytes = core.Int64Ptr(int64(0))
	model.BytesTransferred = core.Int64Ptr(int64(0))
	model.BytesRead = core.Int64Ptr(int64(0))
	model.TotalFileCount = core.Int64Ptr(int64(0))
	model.TransferredFileCount = core.Int64Ptr(int64(0))

	result, err := brsmigration.ResourceIbmBrsMigrationWorkloadRunWorkloadRunStatsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBrsMigrationWorkloadRunPayloadResultToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		payloadResultStatsModel := make(map[string]interface{})
		payloadResultStatsModel["logical_size_bytes"] = int(0)
		payloadResultStatsModel["bytes_transferred"] = int(0)
		payloadResultStatsModel["bytes_read"] = int(0)
		payloadResultStatsModel["total_file_count"] = int(0)
		payloadResultStatsModel["transferred_file_count"] = int(0)

		model := make(map[string]interface{})
		model["payload_id"] = "pl-c3d4e5f6-a7b8-9012-cdef-012345678901"
		model["status"] = "accepted"
		model["message"] = "Source volume unreachable during transfer."
		model["stats"] = []map[string]interface{}{payloadResultStatsModel}

		assert.Equal(t, result, model)
	}

	payloadResultStatsModel := new(brsmigrationv2.PayloadResultStats)
	payloadResultStatsModel.LogicalSizeBytes = core.Int64Ptr(int64(0))
	payloadResultStatsModel.BytesTransferred = core.Int64Ptr(int64(0))
	payloadResultStatsModel.BytesRead = core.Int64Ptr(int64(0))
	payloadResultStatsModel.TotalFileCount = core.Int64Ptr(int64(0))
	payloadResultStatsModel.TransferredFileCount = core.Int64Ptr(int64(0))

	model := new(brsmigrationv2.PayloadResult)
	model.PayloadID = core.StringPtr("pl-c3d4e5f6-a7b8-9012-cdef-012345678901")
	model.Status = core.StringPtr("accepted")
	model.Message = core.StringPtr("Source volume unreachable during transfer.")
	model.Stats = payloadResultStatsModel

	result, err := brsmigration.ResourceIbmBrsMigrationWorkloadRunPayloadResultToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBrsMigrationWorkloadRunPayloadResultStatsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["logical_size_bytes"] = int(0)
		model["bytes_transferred"] = int(0)
		model["bytes_read"] = int(0)
		model["total_file_count"] = int(0)
		model["transferred_file_count"] = int(0)

		assert.Equal(t, result, model)
	}

	model := new(brsmigrationv2.PayloadResultStats)
	model.LogicalSizeBytes = core.Int64Ptr(int64(0))
	model.BytesTransferred = core.Int64Ptr(int64(0))
	model.BytesRead = core.Int64Ptr(int64(0))
	model.TotalFileCount = core.Int64Ptr(int64(0))
	model.TransferredFileCount = core.Int64Ptr(int64(0))

	result, err := brsmigration.ResourceIbmBrsMigrationWorkloadRunPayloadResultStatsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
