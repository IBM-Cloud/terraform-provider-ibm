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
	"github.com/IBM/ibm-brs-migration-sdk-go/brsmigrationv1"
	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmBrsMigrationDiscoverBasic(t *testing.T) {
	var conf brsmigrationv1.DiscoverJob
	migrationID := fmt.Sprintf("tf_migration_id_%d", acctest.RandIntRange(10, 100))
	env := "classic"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmBrsMigrationDiscoverDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmBrsMigrationDiscoverConfigBasic(migrationID, env),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmBrsMigrationDiscoverExists("ibm_brs_migration_discover.brs_migration_discover_instance", conf),
					resource.TestCheckResourceAttr("ibm_brs_migration_discover.brs_migration_discover_instance", "migration_id", migrationID),
					resource.TestCheckResourceAttr("ibm_brs_migration_discover.brs_migration_discover_instance", "env", env),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_brs_migration_discover.brs_migration_discover_instance",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmBrsMigrationDiscoverConfigBasic(migrationID string, env string) string {
	return fmt.Sprintf(`
		resource "ibm_brs_migration_discover" "brs_migration_discover_instance" {
			migration_id = "%s"
			env = "%s"
		}
	`, migrationID, env)
}

func testAccCheckIbmBrsMigrationDiscoverExists(n string, obj brsmigrationv1.DiscoverJob) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		brsMigrationClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).BrsMigrationV1()
		if err != nil {
			return err
		}

		getDiscoverOptions := &brsmigrationv1.GetDiscoverOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getDiscoverOptions.SetMigrationID(parts[0])
		getDiscoverOptions.SetJobID(parts[1])

		discoverJob, _, err := brsMigrationClient.GetDiscover(getDiscoverOptions)
		if err != nil {
			return err
		}

		obj = *discoverJob
		return nil
	}
}

func testAccCheckIbmBrsMigrationDiscoverDestroy(s *terraform.State) error {
	brsMigrationClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).BrsMigrationV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_brs_migration_discover" {
			continue
		}

		getDiscoverOptions := &brsmigrationv1.GetDiscoverOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getDiscoverOptions.SetMigrationID(parts[0])
		getDiscoverOptions.SetJobID(parts[1])

		// Try to find the key
		_, response, err := brsMigrationClient.GetDiscover(getDiscoverOptions)

		if err == nil {
			return fmt.Errorf("brs_migration_discover still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for brs_migration_discover (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

func TestResourceIbmBrsMigrationDiscoverDiscoverJobSummaryToMap(t *testing.T) {
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

	result, err := brsmigration.ResourceIbmBrsMigrationDiscoverDiscoverJobSummaryToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBrsMigrationDiscoverDiscoverJobSummaryComputeToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["virtual_server"] = int(0)
		model["bare_metal"] = int(0)

		assert.Equal(t, result, model)
	}

	model := new(brsmigrationv1.DiscoverJobSummaryCompute)
	model.VirtualServer = core.Int64Ptr(int64(0))
	model.BareMetal = core.Int64Ptr(int64(0))

	result, err := brsmigration.ResourceIbmBrsMigrationDiscoverDiscoverJobSummaryComputeToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBrsMigrationDiscoverDiscoverJobSummaryStorageToMap(t *testing.T) {
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

	result, err := brsmigration.ResourceIbmBrsMigrationDiscoverDiscoverJobSummaryStorageToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBrsMigrationDiscoverMapToDiscoverJobPrototypeLocation(t *testing.T) {
	checkResult := func(result *brsmigrationv1.DiscoverJobPrototypeLocation) {
		model := new(brsmigrationv1.DiscoverJobPrototypeLocation)
		model.Datacenters = []string{"testString"}
		model.Regions = []string{"testString"}
		model.Zones = []string{"testString"}

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["datacenters"] = []interface{}{"testString"}
	model["regions"] = []interface{}{"testString"}
	model["zones"] = []interface{}{"testString"}

	result, err := brsmigration.ResourceIbmBrsMigrationDiscoverMapToDiscoverJobPrototypeLocation(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBrsMigrationDiscoverMapToDiscoverJobPrototypeCompute(t *testing.T) {
	checkResult := func(result *brsmigrationv1.DiscoverJobPrototypeCompute) {
		model := new(brsmigrationv1.DiscoverJobPrototypeCompute)
		model.Types = []string{"virtual_server"}

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["types"] = []interface{}{"virtual_server"}

	result, err := brsmigration.ResourceIbmBrsMigrationDiscoverMapToDiscoverJobPrototypeCompute(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBrsMigrationDiscoverMapToDiscoverJobPrototypeStorage(t *testing.T) {
	checkResult := func(result *brsmigrationv1.DiscoverJobPrototypeStorage) {
		model := new(brsmigrationv1.DiscoverJobPrototypeStorage)
		model.Types = []string{"block"}

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["types"] = []interface{}{"block"}

	result, err := brsmigration.ResourceIbmBrsMigrationDiscoverMapToDiscoverJobPrototypeStorage(model)
	assert.Nil(t, err)
	checkResult(result)
}
