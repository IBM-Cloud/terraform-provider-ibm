// Copyright IBM Corp. 2017, 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package database_test

import (
	"fmt"
	"regexp"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

// TestAccIBMDatabaseMongoDBShardingGen2Basic provisions an enterprise-sharding-gen2
// instance with shards=2 and verifies basic resource attributes. A data source
// lookup is added in a second step to confirm shards, plan, and Gen2-specific behaviour
func TestAccIBMDatabaseMongoDBShardingGen2Basic(t *testing.T) {
	t.Parallel()
	var databaseInstanceOne string
	testName := fmt.Sprintf("tf-mongo-sharding-gen2-%d", acctest.RandIntRange(10, 100))
	name := "ibm_database." + testName
	dataName := "data.ibm_database.lookup"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheckEnterprise(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMDatabaseInstanceDestroy,
		Steps: []resource.TestStep{
			// Step 1 — provision and verify resource attributes
			{
				Config: testAccCheckIBMDatabaseInstanceMongoDBShardingGen2WithShards(testName, 2),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(name, &databaseInstanceOne),
					resource.TestCheckResourceAttr(name, "name", testName),
					resource.TestCheckResourceAttr(name, "service", "databases-for-mongodb"),
					resource.TestCheckResourceAttr(name, "plan", "enterprise-sharding-gen2"),
					resource.TestCheckResourceAttr(name, "location", "ca-mon"),
					resource.TestCheckResourceAttr(name, "service_endpoints", "private"),
					resource.TestCheckResourceAttr(name, "shards", "2"),
					// Gen2 has no admin user
					resource.TestCheckResourceAttr(name, "adminuser", ""),
					resource.TestCheckResourceAttrSet(name, "groups.0.disk.0.allocation_mb"),
					resource.TestCheckResourceAttrSet(name, "groups.0.host_flavor.0.id"),
				),
			},
			// Step 2 — add data source and verify shards + Gen2 specifics read back correctly
			{
				Config: testAccCheckIBMDatabaseInstanceMongoDBShardingGen2WithDataSource(testName, 2),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(dataName, "shards", "2"),
					resource.TestCheckResourceAttr(dataName, "plan", "enterprise-sharding-gen2"),
					resource.TestCheckResourceAttr(dataName, "service", "databases-for-mongodb"),
					resource.TestCheckResourceAttrSet(dataName, "guid"),
					resource.TestCheckResourceAttr(dataName, "adminuser", ""),
					resource.TestCheckResourceAttr(dataName, "auto_scaling.#", "0"),
					resource.TestCheckResourceAttr(dataName, "allowlist.#", "0"),
				),
			},
		},
	})
}

// TestAccIBMDatabaseMongoDBShardingGen2DowngradeRejected verifies that attempting
// to decrease shard count is rejected during terraform plan (CustomizeDiff),
// before any API call is made.
func TestAccIBMDatabaseMongoDBShardingGen2DowngradeRejected(t *testing.T) {
	t.Parallel()
	var databaseInstanceOne string
	testName := fmt.Sprintf("tf-mongo-sharding-gen2-%d", acctest.RandIntRange(10, 100))
	name := "ibm_database." + testName

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheckEnterprise(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMDatabaseInstanceDestroy,
		Steps: []resource.TestStep{
			// Step 1 — create with 2 shards
			{
				Config: testAccCheckIBMDatabaseInstanceMongoDBShardingGen2WithShards(testName, 2),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(name, &databaseInstanceOne),
					resource.TestCheckResourceAttr(name, "shards", "2"),
				),
			},
			// Step 2 — attempt to decrease to 1 shard; must fail at plan time
			{
				Config:      testAccCheckIBMDatabaseInstanceMongoDBShardingGen2WithShards(testName, 1),
				ExpectError: regexp.MustCompile(`Shard count cannot be decreased`),
			},
		},
	})
}

// TestAccIBMDatabaseMongoDBShardingGen2InvalidShardCount verifies that shards
// values outside 1–3 are rejected by the schema ValidateFunc before any API
// call is made. No real resource is created.
func TestAccIBMDatabaseMongoDBShardingGen2InvalidShardCount(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckEnterprise(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckIBMDatabaseInstanceMongoDBShardingGen2WithShards("tf-invalid-shards", 4),
				ExpectError: regexp.MustCompile(`shard count must be between 1 and 3`),
			},
			{
				Config:      testAccCheckIBMDatabaseInstanceMongoDBShardingGen2WithShards("tf-invalid-shards", 0),
				ExpectError: regexp.MustCompile(`shard count must be between 1 and 3`),
			},
		},
	})
}

// TestAccIBMDatabaseMongoDBShardingGen2ShardsOnWrongPlan verifies that setting
// shards on a non-enterprise-sharding-gen2 plan is rejected at plan time.
func TestAccIBMDatabaseMongoDBShardingGen2ShardsOnWrongPlan(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckEnterprise(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckIBMDatabaseInstanceMongoDBStandardGen2WithShards("tf-wrong-plan-shards"),
				ExpectError: regexp.MustCompile(`(?s)shards.*only supported.*enterprise-sharding-gen2`),
			},
		},
	})
}

func testAccCheckIBMDatabaseInstanceMongoDBShardingGen2WithShards(name string, shards int) string {
	return fmt.Sprintf(`
data "ibm_resource_group" "test_acc" {
  is_default = true
}

resource "ibm_database" "%[1]s" {
  resource_group_id = data.ibm_resource_group.test_acc.id
  name              = "%[1]s"
  service           = "databases-for-mongodb"
  plan              = "enterprise-sharding-gen2"
  location          = "ca-mon"
  service_endpoints = "private"
  shards            = %[2]d

  group {
    group_id = "member"

    host_flavor {
      id = "bxf.16x64"
    }

    disk {
      allocation_mb = 20480
    }
  }

  timeouts {
    create = "240m"
    update = "120m"
    delete = "15m"
  }
}
`, name, shards)
}

// testAccCheckIBMDatabaseInstanceMongoDBStandardGen2WithShards produces a config
// that sets shards on a standard-gen2 plan — used to verify plan-time rejection.
func testAccCheckIBMDatabaseInstanceMongoDBStandardGen2WithShards(name string) string {
	return fmt.Sprintf(`
data "ibm_resource_group" "test_acc" {
  is_default = true
}

resource "ibm_database" "%[1]s" {
  resource_group_id = data.ibm_resource_group.test_acc.id
  name              = "%[1]s"
  service           = "databases-for-mongodb"
  plan              = "standard-gen2"
  location          = "ca-mon"
  service_endpoints = "private"
  shards            = 2

  group {
    group_id = "member"

    host_flavor {
      id = "bxf.16x64"
    }

    disk {
      allocation_mb = 20480
    }
  }

  timeouts {
    create = "240m"
    update = "120m"
    delete = "15m"
  }
}
`, name)
}

// testAccCheckIBMDatabaseInstanceMongoDBShardingGen2WithDataSource extends the
// WithShards config by attaching a data source lookup — used to verify that the
// data source reads back shards and Gen2-specific attributes correctly.
func testAccCheckIBMDatabaseInstanceMongoDBShardingGen2WithDataSource(name string, shards int) string {
	return testAccCheckIBMDatabaseInstanceMongoDBShardingGen2WithShards(name, shards) + fmt.Sprintf(`
data "ibm_database" "lookup" {
  resource_group_id = data.ibm_resource_group.test_acc.id
  name              = ibm_database.%[1]s.name
  service           = "databases-for-mongodb"
  location          = "ca-mon"
}
`, name)
}
