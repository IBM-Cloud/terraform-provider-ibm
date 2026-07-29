// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package database_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMMongoDBShardingDatabaseInstanceBasic(t *testing.T) {
	t.Parallel()
	databaseResourceGroup := "default"
	var databaseInstanceOne string
	rnd := fmt.Sprintf("tf-mongoSharding-%d", acctest.RandIntRange(10, 100))
	testName := rnd
	name := "ibm_database." + testName

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMDatabaseInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMDatabaseInstanceMongoDBShardingBasic(databaseResourceGroup, testName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(name, &databaseInstanceOne),
					resource.TestCheckResourceAttr(name, "name", testName),
					resource.TestCheckResourceAttr(name, "service", "databases-for-mongodb"),
					resource.TestCheckResourceAttr(name, "plan", "enterprise-sharding"),
					resource.TestCheckResourceAttr(name, "location", acc.Region()),
					resource.TestCheckResourceAttr(name, "adminuser", "admin"),
					resource.TestCheckResourceAttr(name, "groups.0.memory.0.allocation_mb", "98304"),
					resource.TestCheckResourceAttr(name, "groups.0.disk.0.allocation_mb", "122880"),
					resource.TestCheckResourceAttr(name, "allowlist.#", "1"),
					resource.TestCheckResourceAttr(name, "users.#", "1"),
				),
			},
			{
				Config: testAccCheckIBMDatabaseInstanceMongoDBShardingFullyspecified(databaseResourceGroup, testName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(name, &databaseInstanceOne),
					resource.TestCheckResourceAttr(name, "name", testName),
					resource.TestCheckResourceAttr(name, "service", "databases-for-mongodb"),
					resource.TestCheckResourceAttr(name, "plan", "enterprise-sharding"),
					resource.TestCheckResourceAttr(name, "location", acc.Region()),
					resource.TestCheckResourceAttr(name, "groups.0.memory.0.allocation_mb", "196608"),
					resource.TestCheckResourceAttr(name, "groups.0.disk.0.allocation_mb", "245760"),
					resource.TestCheckResourceAttr(name, "allowlist.#", "2"),
					resource.TestCheckResourceAttr(name, "users.#", "2"),
				),
			},
			{
				Config: testAccCheckIBMDatabaseInstanceMongoDBShardingReduced(databaseResourceGroup, testName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(name, &databaseInstanceOne),
					resource.TestCheckResourceAttr(name, "name", testName),
					resource.TestCheckResourceAttr(name, "service", "databases-for-mongodb"),
					resource.TestCheckResourceAttr(name, "plan", "enterprise-sharding"),
					resource.TestCheckResourceAttr(name, "location", acc.Region()),
					resource.TestCheckResourceAttr(name, "groups.0.memory.0.allocation_mb", "98304"),
					resource.TestCheckResourceAttr(name, "groups.0.disk.0.allocation_mb", "245760"),
					resource.TestCheckResourceAttr(name, "allowlist.#", "0"),
					resource.TestCheckResourceAttr(name, "users.#", "0"),
				),
			},
		},
	})
}

func TestAccIBMDatabaseInstanceMongoShardingHscale(t *testing.T) {
	t.Parallel()

	databaseResourceGroup := "default"
	var databaseInstanceOne string
	rnd := fmt.Sprintf("tf-mongoSharding-%d", acctest.RandIntRange(10, 100))
	testName := rnd
	name := "ibm_database." + testName

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMDatabaseInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMDatabaseInstanceMongoDBShardingMinimal(databaseResourceGroup, testName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(name, &databaseInstanceOne),
					resource.TestCheckResourceAttr(name, "name", testName),
					resource.TestCheckResourceAttr(name, "service", "databases-for-mongodb"),
					resource.TestCheckResourceAttr(name, "plan", "enterprise-sharding"),
					resource.TestCheckResourceAttr(name, "location", acc.Region()),
					resource.TestCheckResourceAttr(name, "groups.0.count", "3"),
				),
			},
			{
				Config: testAccCheckIBMDatabaseInstanceMongoDBShardingHScale(databaseResourceGroup, testName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(name, &databaseInstanceOne),
					resource.TestCheckResourceAttr(name, "name", testName),
					resource.TestCheckResourceAttr(name, "service", "databases-for-mongodb"),
					resource.TestCheckResourceAttr(name, "plan", "enterprise-sharding"),
					resource.TestCheckResourceAttr(name, "location", acc.Region()),
					resource.TestCheckResourceAttr(name, "groups.0.count", "6"),
				),
			},
		},
	})
}

// TestAccIBMMongoDBShardingGen2DatabaseInstanceBasic provisions a Gen2 MongoDB
// enterprise-sharding instance using the enterprise-sharding-gen2 plan. This plan
// is multitenant (no dedicated host flavors) and only supports private service endpoints.
// The shard member count is fixed at 3 by the plan — members and host_flavor must not
// be specified in the group block. Minimum disk is 30 GB (30720 MB) per member.
//
// Run against a real account with:
// go test -timeout 240m -run TestAccIBMMongoDBShardingGen2DatabaseInstanceBasic ./ibm/service/database/...
func TestAccIBMMongoDBShardingGen2DatabaseInstanceBasic(t *testing.T) {
	t.Parallel()
	var databaseInstanceOne string
	rnd := fmt.Sprintf("tf-mongoShardingGen2-%d", acctest.RandIntRange(10, 100))
	testName := rnd
	name := "ibm_database." + testName

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheckEnterprise(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMDatabaseInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMDatabaseInstanceMongoDBShardingGen2Basic(testName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(name, &databaseInstanceOne),
					resource.TestCheckResourceAttr(name, "name", testName),
					resource.TestCheckResourceAttr(name, "service", "databases-for-mongodb"),
					resource.TestCheckResourceAttr(name, "plan", "enterprise-sharding-gen2"),
					resource.TestCheckResourceAttr(name, "location", "ca-mon"),
					resource.TestCheckResourceAttr(name, "service_endpoints", "private"),
					resource.TestCheckResourceAttr(name, "groups.0.count", "3"),
				),
			},
		},
	})
}

func testAccCheckIBMDatabaseInstanceMongoDBShardingGen2Basic(name string) string {
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

		# host_flavor must NOT be set: enterprise-sharding-gen2 is multitenant
		# members  must NOT be set: shard count is fixed at 3 by the plan; broker rejects the field
		group {
			group_id = "member"

			disk {
				allocation_mb = 30720 # minimum 30 GB; step size 3 GB
			}
		}

		timeouts {
			create = "480m"
			update = "480m"
			delete = "15m"
		}
	}
			`, name, acc.Region())
}

func testAccCheckIBMDatabaseInstanceMongoDBShardingBasic(databaseResourceGroup string, name string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "test_acc" {
		name = "%[1]s"
	}

	resource "ibm_database" "%[2]s" {
		resource_group_id            = data.ibm_resource_group.test_acc.id
		name                         = "%[2]s"
		service                      = "databases-for-mongodb"
		plan                         = "enterprise-sharding"
		location                     = "%[3]s"
		adminpassword                = "secure-Password12345"
		group {
			group_id = "member"
			host_flavor {
				id = "b3c.4x16.encrypted"
			}
			disk {
				allocation_mb = 20480
			}
		}
		service_endpoints            = "public"
		users {
		  name     = "user123"
		  password = "secure-Password12345"
		  type     = "database"
		}
		allowlist {
		  address     = "172.168.1.2/32"
		  description = "desc1"
		}
		timeouts {
			create = "480m"
			update = "480m"
			delete = "15m"
		}
	}
				`, databaseResourceGroup, name, acc.Region())
}

func testAccCheckIBMDatabaseInstanceMongoDBShardingFullyspecified(databaseResourceGroup string, name string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "test_acc" {
		name = "%[1]s"
	}

	resource "ibm_database" "%[2]s" {
		resource_group_id            = data.ibm_resource_group.test_acc.id
		name                         = "%[2]s"
		service                      = "databases-for-mongodb"
		plan                         = "enterprise-sharding"
		location                     = "%[3]s"
		adminpassword                = "secure-Password12345"
		group {
			group_id = "member"
			host_flavor {
				id = "b3c.8x32.encrypted"
			}
			disk {
				allocation_mb = 40960
			}
		}
		service_endpoints            = "public"
		users {
		  name     = "user123"
		  password = "secure-Password12345"
		  type     = "database"
		}
		users {
		  name     = "user124"
		  password = "password$12Password"
		  type     = "ops_manager"
		}
		allowlist {
		  address     = "172.168.1.2/32"
		  description = "desc1"
		}
		allowlist {
		  address     = "172.168.1.1/32"
		  description = "desc"
		}
		timeouts {
			create = "480m"
			update = "480m"
			delete = "15m"
		}
	}
				`, databaseResourceGroup, name, acc.Region())
}

func testAccCheckIBMDatabaseInstanceMongoDBShardingReduced(databaseResourceGroup string, name string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "test_acc" {
		name = "%[1]s"
	  }

	  resource "ibm_database" "%[2]s" {
		resource_group_id            = data.ibm_resource_group.test_acc.id
		name                         = "%[2]s"
		service                      = "databases-for-mongodb"
		plan                         = "enterprise-sharding"
		location                     = "%[3]s"
		adminpassword                = "secure-Password12345"
		group {
			group_id = "member"
			host_flavor {
				id = "b3c.4x16.encrypted"
			}
			disk {
				allocation_mb = 40960
			}
		}
		service_endpoints            = "public"
		timeouts {
			create = "480m"
			update = "480m"
			delete = "15m"
		}
	  }
				`, databaseResourceGroup, name, acc.Region())
}

func testAccCheckIBMDatabaseInstanceMongoDBShardingMinimal(databaseResourceGroup string, name string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "test_acc" {
		name = "%[1]s"
	}

	resource "ibm_database" "%[2]s" {
		resource_group_id            = data.ibm_resource_group.test_acc.id
		name                         = "%[2]s"
		service                      = "databases-for-mongodb"
		plan                         = "enterprise-sharding"
		location                     = "%[3]s"
		service_endpoints            = "private"

		group {
			group_id = "member"

			host_flavor {
				id = "b3c.4x16.encrypted"
			}

			members {
				allocation_count = 3
			}
		}

		timeouts {
			create = "480m"
			update = "480m"
			delete = "15m"
		}
	}
				`, databaseResourceGroup, name, acc.Region())
}

func testAccCheckIBMDatabaseInstanceMongoDBShardingHScale(databaseResourceGroup string, name string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "test_acc" {
		name = "%[1]s"
	}

	resource "ibm_database" "%[2]s" {
		resource_group_id            = data.ibm_resource_group.test_acc.id
		name                         = "%[2]s"
		service                      = "databases-for-mongodb"
		plan                         = "enterprise-sharding"
		location                     = "%[3]s"
		service_endpoints            = "private"

		group {
			group_id = "member"

			host_flavor {
				id = "b3c.4x16.encrypted"
			}

			members {
				allocation_count = 6
			}
		}
		timeouts {
			create = "480m"
			update = "480m"
			delete = "15m"
		}
	}
				`, databaseResourceGroup, name, acc.Region())
}
