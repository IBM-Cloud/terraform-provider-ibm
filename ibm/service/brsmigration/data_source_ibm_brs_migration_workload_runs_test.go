// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.114.3-943fbc81-20260603-173645
*/

package brsmigration_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/brsmigration"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/stretchr/testify/assert"
	"github.ibm.com/BackupAndRecovery/brs-migration-orchestrator/brsmigrationv2"
	. "github.com/IBM-Cloud/terraform-provider-ibm/ibm/unittest"
	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmBrsMigrationWorkloadRunsDataSourceBasic(t *testing.T) {
	workloadRunMigrationID := fmt.Sprintf("tf_migration_id_%d", acctest.RandIntRange(10, 100))
	workloadRunWorkloadID := fmt.Sprintf("tf_workload_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmBrsMigrationWorkloadRunsDataSourceConfigBasic(workloadRunMigrationID, workloadRunWorkloadID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_brs_migration_workload_runs.brs_migration_workload_runs_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_brs_migration_workload_runs.brs_migration_workload_runs_instance", "migration_id"),
					resource.TestCheckResourceAttrSet("data.ibm_brs_migration_workload_runs.brs_migration_workload_runs_instance", "workload_id"),
					resource.TestCheckResourceAttrSet("data.ibm_brs_migration_workload_runs.brs_migration_workload_runs_instance", "runs.#"),
					resource.TestCheckResourceAttr("data.ibm_brs_migration_workload_runs.brs_migration_workload_runs_instance", "runs.0.workload_id", workloadRunWorkloadID),
				),
			},
		},
	})
}

func testAccCheckIbmBrsMigrationWorkloadRunsDataSourceConfigBasic(workloadRunMigrationID string, workloadRunWorkloadID string) string {
	return fmt.Sprintf(`
		resource "ibm_brs_migration_workload_run" "brs_migration_workload_run_instance" {
			migration_id = "%s"
			workload_id = "%s"
		}

		data "ibm_brs_migration_workload_runs" "brs_migration_workload_runs_instance" {
			migration_id = ibm_brs_migration_workload_run.brs_migration_workload_run_instance.migration_id
			workload_id = ibm_brs_migration_workload_run.brs_migration_workload_run_instance.workload_id
			status = ibm_brs_migration_workload_run.brs_migration_workload_run_instance.status
			run_type = ibm_brs_migration_workload_run.brs_migration_workload_run_instance.run_type
		}
	`, workloadRunMigrationID, workloadRunWorkloadID)
}


func TestDataSourceIbmBrsMigrationWorkloadRunsWorkloadRunToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		workloadRunStatsModel := make(map[string]interface{})
		workloadRunStatsModel["logical_size_bytes"] = int(0)
		workloadRunStatsModel["bytes_transferred"] = int(0)
		workloadRunStatsModel["bytes_read"] = int(0)
		workloadRunStatsModel["total_file_count"] = int(0)
		workloadRunStatsModel["transferred_file_count"] = int(0)

		payloadResultStatsModel := make(map[string]interface{})
		payloadResultStatsModel["logical_size_bytes"] = int(0)
		payloadResultStatsModel["bytes_transferred"] = int(0)
		payloadResultStatsModel["bytes_read"] = int(0)
		payloadResultStatsModel["total_file_count"] = int(0)
		payloadResultStatsModel["transferred_file_count"] = int(0)

		payloadResultModel := make(map[string]interface{})
		payloadResultModel["payload_id"] = "pl-c3d4e5f6-a7b8-9012-cdef-012345678901"
		payloadResultModel["status"] = "accepted"
		payloadResultModel["message"] = "Source volume unreachable during transfer."
		payloadResultModel["stats"] = []map[string]interface{}{payloadResultStatsModel}

		model := make(map[string]interface{})
		model["id"] = "run-e5f6a7b8-c9d0-1234-ef01-234567890123"
		model["workload_id"] = "wl-a1b2c3d4-e5f6-7890-abcd-ef1234567890"
		model["operation_type"] = "backup"
		model["run_type"] = "scheduled"
		model["status"] = "accepted"
		model["started_at"] = "2019-01-01T12:00:00.000Z"
		model["completed_at"] = "2019-01-01T12:00:00.000Z"
		model["duration_seconds"] = int(0)
		model["message"] = "Run completed successfully with 2 payloads transferred."
		model["stats"] = []map[string]interface{}{workloadRunStatsModel}
		model["payload_results"] = []map[string]interface{}{payloadResultModel}

		assert.Equal(t, result, model)
	}

	workloadRunStatsModel := new(brsmigrationv2.WorkloadRunStats)
	workloadRunStatsModel.LogicalSizeBytes = core.Int64Ptr(int64(0))
	workloadRunStatsModel.BytesTransferred = core.Int64Ptr(int64(0))
	workloadRunStatsModel.BytesRead = core.Int64Ptr(int64(0))
	workloadRunStatsModel.TotalFileCount = core.Int64Ptr(int64(0))
	workloadRunStatsModel.TransferredFileCount = core.Int64Ptr(int64(0))

	payloadResultStatsModel := new(brsmigrationv2.PayloadResultStats)
	payloadResultStatsModel.LogicalSizeBytes = core.Int64Ptr(int64(0))
	payloadResultStatsModel.BytesTransferred = core.Int64Ptr(int64(0))
	payloadResultStatsModel.BytesRead = core.Int64Ptr(int64(0))
	payloadResultStatsModel.TotalFileCount = core.Int64Ptr(int64(0))
	payloadResultStatsModel.TransferredFileCount = core.Int64Ptr(int64(0))

	payloadResultModel := new(brsmigrationv2.PayloadResult)
	payloadResultModel.PayloadID = core.StringPtr("pl-c3d4e5f6-a7b8-9012-cdef-012345678901")
	payloadResultModel.Status = core.StringPtr("accepted")
	payloadResultModel.Message = core.StringPtr("Source volume unreachable during transfer.")
	payloadResultModel.Stats = payloadResultStatsModel

	model := new(brsmigrationv2.WorkloadRun)
	model.ID = core.StringPtr("run-e5f6a7b8-c9d0-1234-ef01-234567890123")
	model.WorkloadID = core.StringPtr("wl-a1b2c3d4-e5f6-7890-abcd-ef1234567890")
	model.OperationType = core.StringPtr("backup")
	model.RunType = core.StringPtr("scheduled")
	model.Status = core.StringPtr("accepted")
	model.StartedAt = CreateMockDateTime("2019-01-01T12:00:00.000Z")
	model.CompletedAt = CreateMockDateTime("2019-01-01T12:00:00.000Z")
	model.DurationSeconds = core.Int64Ptr(int64(0))
	model.Message = core.StringPtr("Run completed successfully with 2 payloads transferred.")
	model.Stats = workloadRunStatsModel
	model.PayloadResults = []brsmigrationv2.PayloadResult{*payloadResultModel}

	result, err := brsmigration.DataSourceIbmBrsMigrationWorkloadRunsWorkloadRunToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBrsMigrationWorkloadRunsWorkloadRunStatsToMap(t *testing.T) {
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

	result, err := brsmigration.DataSourceIbmBrsMigrationWorkloadRunsWorkloadRunStatsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBrsMigrationWorkloadRunsPayloadResultToMap(t *testing.T) {
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

	result, err := brsmigration.DataSourceIbmBrsMigrationWorkloadRunsPayloadResultToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBrsMigrationWorkloadRunsPayloadResultStatsToMap(t *testing.T) {
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

	result, err := brsmigration.DataSourceIbmBrsMigrationWorkloadRunsPayloadResultStatsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
