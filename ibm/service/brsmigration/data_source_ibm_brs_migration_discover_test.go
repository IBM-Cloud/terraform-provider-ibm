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
	"github.com/IBM/ibm-brs-migration-sdk-go/brsmigrationv1"
	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmBrsMigrationDiscoverDataSourceBasic(t *testing.T) {
	discoverJobMigrationID := fmt.Sprintf("tf_migration_id_%d", acctest.RandIntRange(10, 100))
	discoverJobEnv := "classic"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmBrsMigrationDiscoverDataSourceConfigBasic(discoverJobMigrationID, discoverJobEnv),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_brs_migration_discover.brs_migration_discover_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_brs_migration_discover.brs_migration_discover_instance", "migration_id"),
					resource.TestCheckResourceAttrSet("data.ibm_brs_migration_discover.brs_migration_discover_instance", "job_id"),
					resource.TestCheckResourceAttrSet("data.ibm_brs_migration_discover.brs_migration_discover_instance", "env"),
					resource.TestCheckResourceAttrSet("data.ibm_brs_migration_discover.brs_migration_discover_instance", "state"),
				),
			},
		},
	})
}

func testAccCheckIbmBrsMigrationDiscoverDataSourceConfigBasic(discoverJobMigrationID string, discoverJobEnv string) string {
	return fmt.Sprintf(`
		resource "ibm_brs_migration_discover" "brs_migration_discover_instance" {
			migration_id = "%s"
			env = "%s"
		}

		data "ibm_brs_migration_discover" "brs_migration_discover_instance" {
			migration_id = ibm_brs_migration_discover.brs_migration_discover_instance.migration_id
			job_id = ibm_brs_migration_discover.brs_migration_discover_instance.job_id
		}
	`, discoverJobMigrationID, discoverJobEnv)
}


func TestDataSourceIbmBrsMigrationDiscoverDiscoverJobSummaryToMap(t *testing.T) {
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

	discoverJobSummaryComputeModel := new(brsmigrationv1.DiscoverJobSummaryCompute)
	discoverJobSummaryComputeModel.VirtualServer = core.Int64Ptr(int64(0))
	discoverJobSummaryComputeModel.BareMetal = core.Int64Ptr(int64(0))

	discoverJobSummaryStorageModel := new(brsmigrationv1.DiscoverJobSummaryStorage)
	discoverJobSummaryStorageModel.Block = core.Int64Ptr(int64(0))
	discoverJobSummaryStorageModel.File = core.Int64Ptr(int64(0))
	discoverJobSummaryStorageModel.San = core.Int64Ptr(int64(0))
	discoverJobSummaryStorageModel.Local = core.Int64Ptr(int64(0))

	model := new(brsmigrationv1.DiscoverJobSummary)
	model.Total = core.Int64Ptr(int64(0))
	model.Compute = discoverJobSummaryComputeModel
	model.Storage = discoverJobSummaryStorageModel

	result, err := brsmigration.DataSourceIbmBrsMigrationDiscoverDiscoverJobSummaryToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBrsMigrationDiscoverDiscoverJobSummaryComputeToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["virtual_server"] = int(0)
		model["bare_metal"] = int(0)

		assert.Equal(t, result, model)
	}

	model := new(brsmigrationv1.DiscoverJobSummaryCompute)
	model.VirtualServer = core.Int64Ptr(int64(0))
	model.BareMetal = core.Int64Ptr(int64(0))

	result, err := brsmigration.DataSourceIbmBrsMigrationDiscoverDiscoverJobSummaryComputeToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBrsMigrationDiscoverDiscoverJobSummaryStorageToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["block"] = int(0)
		model["file"] = int(0)
		model["san"] = int(0)
		model["local"] = int(0)

		assert.Equal(t, result, model)
	}

	model := new(brsmigrationv1.DiscoverJobSummaryStorage)
	model.Block = core.Int64Ptr(int64(0))
	model.File = core.Int64Ptr(int64(0))
	model.San = core.Int64Ptr(int64(0))
	model.Local = core.Int64Ptr(int64(0))

	result, err := brsmigration.DataSourceIbmBrsMigrationDiscoverDiscoverJobSummaryStorageToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
