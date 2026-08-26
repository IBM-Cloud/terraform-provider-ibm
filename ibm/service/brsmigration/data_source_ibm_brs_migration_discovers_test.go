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

func TestAccIbmBrsMigrationDiscoversDataSourceBasic(t *testing.T) {
	discoverJobMigrationID := fmt.Sprintf("tf_migration_id_%d", acctest.RandIntRange(10, 100))
	discoverJobEnv := "classic"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmBrsMigrationDiscoversDataSourceConfigBasic(discoverJobMigrationID, discoverJobEnv),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_brs_migration_discovers.brs_migration_discovers_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_brs_migration_discovers.brs_migration_discovers_instance", "migration_id"),
					resource.TestCheckResourceAttrSet("data.ibm_brs_migration_discovers.brs_migration_discovers_instance", "discover.#"),
					resource.TestCheckResourceAttr("data.ibm_brs_migration_discovers.brs_migration_discovers_instance", "discover.0.env", discoverJobEnv),
				),
			},
		},
	})
}

func testAccCheckIbmBrsMigrationDiscoversDataSourceConfigBasic(discoverJobMigrationID string, discoverJobEnv string) string {
	return fmt.Sprintf(`
		resource "ibm_brs_migration_discover" "brs_migration_discover_instance" {
			migration_id = "%s"
			env = "%s"
		}

		data "ibm_brs_migration_discovers" "brs_migration_discovers_instance" {
			migration_id = ibm_brs_migration_discover.brs_migration_discover_instance.migration_id
			env = ibm_brs_migration_discover.brs_migration_discover_instance.env
			state = ibm_brs_migration_discover.brs_migration_discover_instance.state
		}
	`, discoverJobMigrationID, discoverJobEnv)
}


func TestDataSourceIbmBrsMigrationDiscoversDiscoverJobToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		discoverJobSummaryComputeModel := make(map[string]interface{})
		discoverJobSummaryComputeModel["virtual_server"] = int(0)
		discoverJobSummaryComputeModel["bare_metal"] = int(0)

		discoverJobSummaryStorageModel := make(map[string]interface{})
		discoverJobSummaryStorageModel["block"] = int(0)
		discoverJobSummaryStorageModel["file"] = int(0)
		discoverJobSummaryStorageModel["san"] = int(0)
		discoverJobSummaryStorageModel["local"] = int(0)

		discoverJobSummaryModel := make(map[string]interface{})
		discoverJobSummaryModel["total"] = int(0)
		discoverJobSummaryModel["compute"] = []map[string]interface{}{discoverJobSummaryComputeModel}
		discoverJobSummaryModel["storage"] = []map[string]interface{}{discoverJobSummaryStorageModel}

		model := make(map[string]interface{})
		model["id"] = "job-12345678-abcd-ef01-2345-678901234567"
		model["env"] = "classic"
		model["state"] = "pending"
		model["start_time"] = "2019-01-01T12:00:00.000Z"
		model["end_time"] = "2019-01-01T12:00:00.000Z"
		model["message"] = "Discovery completed: 12 hosts and 34 volumes found."
		model["summary"] = []map[string]interface{}{discoverJobSummaryModel}

		assert.Equal(t, result, model)
	}

	discoverJobSummaryComputeModel := new(brsmigrationv2.DiscoverJobSummaryCompute)
	discoverJobSummaryComputeModel.VirtualServer = core.Int64Ptr(int64(0))
	discoverJobSummaryComputeModel.BareMetal = core.Int64Ptr(int64(0))

	discoverJobSummaryStorageModel := new(brsmigrationv2.DiscoverJobSummaryStorage)
	discoverJobSummaryStorageModel.Block = core.Int64Ptr(int64(0))
	discoverJobSummaryStorageModel.File = core.Int64Ptr(int64(0))
	discoverJobSummaryStorageModel.San = core.Int64Ptr(int64(0))
	discoverJobSummaryStorageModel.Local = core.Int64Ptr(int64(0))

	discoverJobSummaryModel := new(brsmigrationv2.DiscoverJobSummary)
	discoverJobSummaryModel.Total = core.Int64Ptr(int64(0))
	discoverJobSummaryModel.Compute = discoverJobSummaryComputeModel
	discoverJobSummaryModel.Storage = discoverJobSummaryStorageModel

	model := new(brsmigrationv2.DiscoverJob)
	model.ID = core.StringPtr("job-12345678-abcd-ef01-2345-678901234567")
	model.Env = core.StringPtr("classic")
	model.State = core.StringPtr("pending")
	model.StartTime = CreateMockDateTime("2019-01-01T12:00:00.000Z")
	model.EndTime = CreateMockDateTime("2019-01-01T12:00:00.000Z")
	model.Message = core.StringPtr("Discovery completed: 12 hosts and 34 volumes found.")
	model.Summary = discoverJobSummaryModel

	result, err := brsmigration.DataSourceIbmBrsMigrationDiscoversDiscoverJobToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBrsMigrationDiscoversDiscoverJobSummaryToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		discoverJobSummaryComputeModel := make(map[string]interface{})
		discoverJobSummaryComputeModel["virtual_server"] = int(0)
		discoverJobSummaryComputeModel["bare_metal"] = int(0)

		discoverJobSummaryStorageModel := make(map[string]interface{})
		discoverJobSummaryStorageModel["block"] = int(0)
		discoverJobSummaryStorageModel["file"] = int(0)
		discoverJobSummaryStorageModel["san"] = int(0)
		discoverJobSummaryStorageModel["local"] = int(0)

		model := make(map[string]interface{})
		model["total"] = int(0)
		model["compute"] = []map[string]interface{}{discoverJobSummaryComputeModel}
		model["storage"] = []map[string]interface{}{discoverJobSummaryStorageModel}

		assert.Equal(t, result, model)
	}

	discoverJobSummaryComputeModel := new(brsmigrationv2.DiscoverJobSummaryCompute)
	discoverJobSummaryComputeModel.VirtualServer = core.Int64Ptr(int64(0))
	discoverJobSummaryComputeModel.BareMetal = core.Int64Ptr(int64(0))

	discoverJobSummaryStorageModel := new(brsmigrationv2.DiscoverJobSummaryStorage)
	discoverJobSummaryStorageModel.Block = core.Int64Ptr(int64(0))
	discoverJobSummaryStorageModel.File = core.Int64Ptr(int64(0))
	discoverJobSummaryStorageModel.San = core.Int64Ptr(int64(0))
	discoverJobSummaryStorageModel.Local = core.Int64Ptr(int64(0))

	model := new(brsmigrationv2.DiscoverJobSummary)
	model.Total = core.Int64Ptr(int64(0))
	model.Compute = discoverJobSummaryComputeModel
	model.Storage = discoverJobSummaryStorageModel

	result, err := brsmigration.DataSourceIbmBrsMigrationDiscoversDiscoverJobSummaryToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBrsMigrationDiscoversDiscoverJobSummaryComputeToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["virtual_server"] = int(0)
		model["bare_metal"] = int(0)

		assert.Equal(t, result, model)
	}

	model := new(brsmigrationv2.DiscoverJobSummaryCompute)
	model.VirtualServer = core.Int64Ptr(int64(0))
	model.BareMetal = core.Int64Ptr(int64(0))

	result, err := brsmigration.DataSourceIbmBrsMigrationDiscoversDiscoverJobSummaryComputeToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBrsMigrationDiscoversDiscoverJobSummaryStorageToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["block"] = int(0)
		model["file"] = int(0)
		model["san"] = int(0)
		model["local"] = int(0)

		assert.Equal(t, result, model)
	}

	model := new(brsmigrationv2.DiscoverJobSummaryStorage)
	model.Block = core.Int64Ptr(int64(0))
	model.File = core.Int64Ptr(int64(0))
	model.San = core.Int64Ptr(int64(0))
	model.Local = core.Int64Ptr(int64(0))

	result, err := brsmigration.DataSourceIbmBrsMigrationDiscoversDiscoverJobSummaryStorageToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
