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
	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmBrsMigrationWorkloadRunDataSourceBasic(t *testing.T) {
	workloadRunMigrationID := fmt.Sprintf("tf_migration_id_%d", acctest.RandIntRange(10, 100))
	workloadRunWorkloadID := fmt.Sprintf("tf_workload_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmBrsMigrationWorkloadRunDataSourceConfigBasic(workloadRunMigrationID, workloadRunWorkloadID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_brs_migration_workload_run.brs_migration_workload_run_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_brs_migration_workload_run.brs_migration_workload_run_instance", "migration_id"),
					resource.TestCheckResourceAttrSet("data.ibm_brs_migration_workload_run.brs_migration_workload_run_instance", "workload_id"),
					resource.TestCheckResourceAttrSet("data.ibm_brs_migration_workload_run.brs_migration_workload_run_instance", "run_id"),
					resource.TestCheckResourceAttrSet("data.ibm_brs_migration_workload_run.brs_migration_workload_run_instance", "operation_type"),
					resource.TestCheckResourceAttrSet("data.ibm_brs_migration_workload_run.brs_migration_workload_run_instance", "run_type"),
					resource.TestCheckResourceAttrSet("data.ibm_brs_migration_workload_run.brs_migration_workload_run_instance", "status"),
					resource.TestCheckResourceAttrSet("data.ibm_brs_migration_workload_run.brs_migration_workload_run_instance", "started_at"),
					resource.TestCheckResourceAttrSet("data.ibm_brs_migration_workload_run.brs_migration_workload_run_instance", "payload_results.#"),
				),
			},
		},
	})
}

func testAccCheckIbmBrsMigrationWorkloadRunDataSourceConfigBasic(workloadRunMigrationID string, workloadRunWorkloadID string) string {
	return fmt.Sprintf(`
		resource "ibm_brs_migration_workload_run" "brs_migration_workload_run_instance" {
			migration_id = "%s"
			workload_id = "%s"
		}

		data "ibm_brs_migration_workload_run" "brs_migration_workload_run_instance" {
			migration_id = ibm_brs_migration_workload_run.brs_migration_workload_run_instance.migration_id
			workload_id = ibm_brs_migration_workload_run.brs_migration_workload_run_instance.workload_id
			run_id = ibm_brs_migration_workload_run.brs_migration_workload_run_instance.run_id
		}
	`, workloadRunMigrationID, workloadRunWorkloadID)
}


func TestDataSourceIbmBrsMigrationWorkloadRunWorkloadRunStatsToMap(t *testing.T) {
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

	result, err := brsmigration.DataSourceIbmBrsMigrationWorkloadRunWorkloadRunStatsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBrsMigrationWorkloadRunPayloadResultToMap(t *testing.T) {
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

	result, err := brsmigration.DataSourceIbmBrsMigrationWorkloadRunPayloadResultToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBrsMigrationWorkloadRunPayloadResultStatsToMap(t *testing.T) {
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

	result, err := brsmigration.DataSourceIbmBrsMigrationWorkloadRunPayloadResultStatsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
