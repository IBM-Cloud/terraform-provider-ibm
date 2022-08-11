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

func TestAccIBMCassandraDatabaseInstanceBasic(t *testing.T) {
	t.Parallel()
	databaseResourceGroup := "default"
	var databaseInstanceOne string
	rnd := fmt.Sprintf("tf-Datastax-%d", acctest.RandIntRange(10, 100))
	testName := rnd
	name := "ibm_database." + testName

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMDatabaseInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMDatabaseInstanceCassandraBasic(databaseResourceGroup, testName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(name, &databaseInstanceOne),
					resource.TestCheckResourceAttr(name, "name", testName),
					resource.TestCheckResourceAttr(name, "service", "databases-for-cassandra"),
					resource.TestCheckResourceAttr(name, "plan", "enterprise"),
					resource.TestCheckResourceAttr(name, "location", acc.IcdDbRegion),
					resource.TestCheckResourceAttr(name, "adminuser", "admin"),
					resource.TestCheckResourceAttr(name, "members_memory_allocation_mb", "36864"),
					resource.TestCheckResourceAttr(name, "members_disk_allocation_mb", "61440"),
					resource.TestCheckResourceAttr(name, "whitelist.#", "1"),
					resource.TestCheckResourceAttr(name, "users.#", "1"),
					resource.TestCheckResourceAttr(name, "connectionstrings.#", "2"),
					resource.TestCheckResourceAttr(name, "connectionstrings.1.name", "admin"),
					resource.TestCheckResourceAttr(name, "connectionstrings.0.hosts.#", "1"),
				),
			},
			{
				Config: testAccCheckIBMDatabaseInstanceCassandraFullyspecified(databaseResourceGroup, testName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(name, &databaseInstanceOne),
					resource.TestCheckResourceAttr(name, "name", testName),
					resource.TestCheckResourceAttr(name, "service", "databases-for-cassandra"),
					resource.TestCheckResourceAttr(name, "plan", "enterprise"),
					resource.TestCheckResourceAttr(name, "location", acc.IcdDbRegion),
					resource.TestCheckResourceAttr(name, "members_memory_allocation_mb", "38400"),
					resource.TestCheckResourceAttr(name, "members_disk_allocation_mb", "61440"),
					resource.TestCheckResourceAttr(name, "whitelist.#", "2"),
					resource.TestCheckResourceAttr(name, "users.#", "2"),
					resource.TestCheckResourceAttr(name, "connectionstrings.#", "3"),
					resource.TestCheckResourceAttr(name, "connectionstrings.2.name", "admin"),
				),
			},
			{
				Config: testAccCheckIBMDatabaseInstanceCassandraReduced(databaseResourceGroup, testName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(name, &databaseInstanceOne),
					resource.TestCheckResourceAttr(name, "name", testName),
					resource.TestCheckResourceAttr(name, "service", "databases-for-cassandra"),
					resource.TestCheckResourceAttr(name, "plan", "enterprise"),
					resource.TestCheckResourceAttr(name, "location", acc.IcdDbRegion),
					resource.TestCheckResourceAttr(name, "members_memory_allocation_mb", "36864"),
					resource.TestCheckResourceAttr(name, "members_disk_allocation_mb", "61440"),
					resource.TestCheckResourceAttr(name, "whitelist.#", "0"),
					resource.TestCheckResourceAttr(name, "users.#", "0"),
					resource.TestCheckResourceAttr(name, "connectionstrings.#", "1"),
				),
			},
		},
	})
}

func TestAccIBMDatabaseInstance_Cassandra_Node(t *testing.T) {
	t.Parallel()
	databaseResourceGroup := "default"
	var databaseInstanceOne string
	rnd := fmt.Sprintf("tf-Datastax-%d", acctest.RandIntRange(10, 100))
	testName := rnd
	name := "ibm_database." + testName

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMDatabaseInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMDatabaseInstanceCassandraNodeBasic(databaseResourceGroup, testName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(name, &databaseInstanceOne),
					resource.TestCheckResourceAttr(name, "name", testName),
					resource.TestCheckResourceAttr(name, "service", "databases-for-cassandra"),
					resource.TestCheckResourceAttr(name, "plan", "enterprise"),
					resource.TestCheckResourceAttr(name, "location", acc.IcdDbRegion),
					resource.TestCheckResourceAttr(name, "adminuser", "admin"),
					resource.TestCheckResourceAttr(name, "node_count", "3"),
					resource.TestCheckResourceAttr(name, "node_memory_allocation_mb", "12288"),
					resource.TestCheckResourceAttr(name, "node_disk_allocation_mb", "20480"),
					resource.TestCheckResourceAttr(name, "node_cpu_allocation_count", "6"),

					resource.TestCheckResourceAttr(name, "whitelist.#", "1"),
					resource.TestCheckResourceAttr(name, "users.#", "1"),
					resource.TestCheckResourceAttr(name, "connectionstrings.#", "2"),
					resource.TestCheckResourceAttr(name, "connectionstrings.1.name", "admin"),
					resource.TestCheckResourceAttr(name, "connectionstrings.0.hosts.#", "1"),
					resource.TestCheckResourceAttr(name, "connectionstrings.0.database", ""),
				),
			},
			{
				Config: testAccCheckIBMDatabaseInstanceCassandraNodeFullyspecified(databaseResourceGroup, testName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(name, &databaseInstanceOne),
					resource.TestCheckResourceAttr(name, "name", testName),
					resource.TestCheckResourceAttr(name, "service", "databases-for-cassandra"),
					resource.TestCheckResourceAttr(name, "plan", "enterprise"),
					resource.TestCheckResourceAttr(name, "location", acc.IcdDbRegion),
					resource.TestCheckResourceAttr(name, "node_count", "3"),
					resource.TestCheckResourceAttr(name, "node_memory_allocation_mb", "12416"),
					resource.TestCheckResourceAttr(name, "node_disk_allocation_mb", "20480"),
					resource.TestCheckResourceAttr(name, "node_cpu_allocation_count", "6"),
					resource.TestCheckResourceAttr(name, "whitelist.#", "2"),
					resource.TestCheckResourceAttr(name, "users.#", "2"),
					resource.TestCheckResourceAttr(name, "connectionstrings.#", "3"),
					resource.TestCheckResourceAttr(name, "connectionstrings.2.name", "admin"),
				),
			},
			{
				Config: testAccCheckIBMDatabaseInstanceCassandraNodeReduced(databaseResourceGroup, testName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(name, &databaseInstanceOne),
					resource.TestCheckResourceAttr(name, "name", testName),
					resource.TestCheckResourceAttr(name, "service", "databases-for-cassandra"),
					resource.TestCheckResourceAttr(name, "plan", "enterprise"),
					resource.TestCheckResourceAttr(name, "location", acc.IcdDbRegion),
					resource.TestCheckResourceAttr(name, "node_count", "3"),
					resource.TestCheckResourceAttr(name, "node_memory_allocation_mb", "12288"),
					resource.TestCheckResourceAttr(name, "node_disk_allocation_mb", "20480"),
					resource.TestCheckResourceAttr(name, "node_cpu_allocation_count", "6"),
					resource.TestCheckResourceAttr(name, "whitelist.#", "0"),
					resource.TestCheckResourceAttr(name, "users.#", "0"),
					resource.TestCheckResourceAttr(name, "connectionstrings.#", "1"),
				),
			},
			{
				Config: testAccCheckIBMDatabaseInstanceCassandraNodeScaleOut(databaseResourceGroup, testName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(name, &databaseInstanceOne),
					resource.TestCheckResourceAttr(name, "name", testName),
					resource.TestCheckResourceAttr(name, "service", "databases-for-cassandra"),
					resource.TestCheckResourceAttr(name, "plan", "enterprise"),
					resource.TestCheckResourceAttr(name, "location", acc.IcdDbRegion),
					resource.TestCheckResourceAttr(name, "node_count", "4"),
					resource.TestCheckResourceAttr(name, "node_memory_allocation_mb", "12288"),
					resource.TestCheckResourceAttr(name, "node_disk_allocation_mb", "20480"),
					resource.TestCheckResourceAttr(name, "node_cpu_allocation_count", "6"),
					resource.TestCheckResourceAttr(name, "whitelist.#", "0"),
					resource.TestCheckResourceAttr(name, "users.#", "0"),
					resource.TestCheckResourceAttr(name, "connectionstrings.#", "1"),
				),
			},
			//{
			//	ResourceName:      name,
			//	ImportState:       true,
			//	ImportStateVerify: true,
			//},
		},
	})
}

func TestAccIBMDatabaseInstance_Cassandra_Group(t *testing.T) {
	t.Parallel()
	databaseResourceGroup := "default"
	var databaseInstanceOne string
	rnd := fmt.Sprintf("tf-Datastax-%d", acctest.RandIntRange(10, 100))
	testName := rnd
	name := "ibm_database." + testName

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMDatabaseInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMDatabaseInstanceCassandraGroupBasic(databaseResourceGroup, testName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(name, &databaseInstanceOne),
					resource.TestCheckResourceAttr(name, "name", testName),
					resource.TestCheckResourceAttr(name, "service", "databases-for-cassandra"),
					resource.TestCheckResourceAttr(name, "plan", "enterprise"),
					resource.TestCheckResourceAttr(name, "location", acc.IcdDbRegion),
					resource.TestCheckResourceAttr(name, "adminuser", "admin"),
					resource.TestCheckResourceAttr(name, "node_count", "3"),
					resource.TestCheckResourceAttr(name, "node_memory_allocation_mb", "12288"),
					resource.TestCheckResourceAttr(name, "node_disk_allocation_mb", "20480"),
					resource.TestCheckResourceAttr(name, "node_cpu_allocation_count", "6"),
					resource.TestCheckResourceAttr(name, "groups.0.count", "3"),
					resource.TestCheckResourceAttr(name, "groups.0.memory.0.allocation_mb", "36864"),
					resource.TestCheckResourceAttr(name, "groups.0.disk.0.allocation_mb", "61440"),
					resource.TestCheckResourceAttr(name, "groups.0.cpu.0.allocation_count", "18"),
					resource.TestCheckResourceAttr(name, "whitelist.#", "1"),
					resource.TestCheckResourceAttr(name, "users.#", "1"),
					resource.TestCheckResourceAttr(name, "connectionstrings.#", "2"),
					resource.TestCheckResourceAttr(name, "connectionstrings.1.name", "admin"),
					resource.TestCheckResourceAttr(name, "connectionstrings.0.hosts.#", "1"),
					resource.TestCheckResourceAttr(name, "connectionstrings.0.database", ""),
				),
			},
			{
				Config: testAccCheckIBMDatabaseInstanceCassandraGroupFullyspecified(databaseResourceGroup, testName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(name, &databaseInstanceOne),
					resource.TestCheckResourceAttr(name, "name", testName),
					resource.TestCheckResourceAttr(name, "service", "databases-for-cassandra"),
					resource.TestCheckResourceAttr(name, "plan", "enterprise"),
					resource.TestCheckResourceAttr(name, "location", acc.IcdDbRegion),
					resource.TestCheckResourceAttr(name, "node_count", "3"),
					resource.TestCheckResourceAttr(name, "node_memory_allocation_mb", "12416"),
					resource.TestCheckResourceAttr(name, "node_disk_allocation_mb", "20480"),
					resource.TestCheckResourceAttr(name, "node_cpu_allocation_count", "6"),
					resource.TestCheckResourceAttr(name, "groups.0.count", "3"),
					resource.TestCheckResourceAttr(name, "groups.0.memory.0.allocation_mb", "37248"),
					resource.TestCheckResourceAttr(name, "groups.0.disk.0.allocation_mb", "61440"),
					resource.TestCheckResourceAttr(name, "groups.0.cpu.0.allocation_count", "18"),
					resource.TestCheckResourceAttr(name, "whitelist.#", "2"),
					resource.TestCheckResourceAttr(name, "users.#", "2"),
					resource.TestCheckResourceAttr(name, "connectionstrings.#", "3"),
					resource.TestCheckResourceAttr(name, "connectionstrings.2.name", "admin"),
				),
			},
			{
				Config: testAccCheckIBMDatabaseInstanceCassandraGroupReduced(databaseResourceGroup, testName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(name, &databaseInstanceOne),
					resource.TestCheckResourceAttr(name, "name", testName),
					resource.TestCheckResourceAttr(name, "service", "databases-for-cassandra"),
					resource.TestCheckResourceAttr(name, "plan", "enterprise"),
					resource.TestCheckResourceAttr(name, "location", acc.IcdDbRegion),
					resource.TestCheckResourceAttr(name, "node_count", "3"),
					resource.TestCheckResourceAttr(name, "node_memory_allocation_mb", "12288"),
					resource.TestCheckResourceAttr(name, "node_disk_allocation_mb", "20480"),
					resource.TestCheckResourceAttr(name, "node_cpu_allocation_count", "6"),
					resource.TestCheckResourceAttr(name, "groups.0.count", "3"),
					resource.TestCheckResourceAttr(name, "groups.0.memory.0.allocation_mb", "36864"),
					resource.TestCheckResourceAttr(name, "groups.0.disk.0.allocation_mb", "61440"),
					resource.TestCheckResourceAttr(name, "groups.0.cpu.0.allocation_count", "18"),
					resource.TestCheckResourceAttr(name, "whitelist.#", "0"),
					resource.TestCheckResourceAttr(name, "users.#", "0"),
					resource.TestCheckResourceAttr(name, "connectionstrings.#", "1"),
				),
			},
			{
				Config: testAccCheckIBMDatabaseInstanceCassandraGroupScaleOut(databaseResourceGroup, testName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(name, &databaseInstanceOne),
					resource.TestCheckResourceAttr(name, "name", testName),
					resource.TestCheckResourceAttr(name, "service", "databases-for-cassandra"),
					resource.TestCheckResourceAttr(name, "plan", "enterprise"),
					resource.TestCheckResourceAttr(name, "location", acc.IcdDbRegion),
					resource.TestCheckResourceAttr(name, "node_count", "4"),
					resource.TestCheckResourceAttr(name, "node_memory_allocation_mb", "12288"),
					resource.TestCheckResourceAttr(name, "node_disk_allocation_mb", "20480"),
					resource.TestCheckResourceAttr(name, "node_cpu_allocation_count", "6"),
					resource.TestCheckResourceAttr(name, "groups.0.count", "4"),
					resource.TestCheckResourceAttr(name, "groups.0.memory.0.allocation_mb", "49152"),
					resource.TestCheckResourceAttr(name, "groups.0.disk.0.allocation_mb", "81920"),
					resource.TestCheckResourceAttr(name, "groups.0.cpu.0.allocation_count", "24"),
					resource.TestCheckResourceAttr(name, "groups.1.count", "3"),
					resource.TestCheckResourceAttr(name, "whitelist.#", "0"),
					resource.TestCheckResourceAttr(name, "users.#", "0"),
					resource.TestCheckResourceAttr(name, "connectionstrings.#", "1"),
				),
			},
		},
	})
}

// TestAccIBMDatabaseInstance_CreateAfterManualDestroy not required as tested by resource_instance tests

func TestAccIBMDatabaseInstanceCassandraImport(t *testing.T) {
	t.Parallel()
	databaseResourceGroup := "default"
	var databaseInstanceOne string
	serviceName := fmt.Sprintf("tf-Datastax-%d", acctest.RandIntRange(10, 100))
	//serviceName := "test_acc"
	resourceName := "ibm_database." + serviceName

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMDatabaseInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMDatabaseInstanceCassandraImport(databaseResourceGroup, serviceName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(resourceName, &databaseInstanceOne),
					resource.TestCheckResourceAttr(resourceName, "name", serviceName),
					resource.TestCheckResourceAttr(resourceName, "service", "databases-for-cassandra"),
					resource.TestCheckResourceAttr(resourceName, "plan", "enterprise"),
					resource.TestCheckResourceAttr(resourceName, "location", acc.IcdDbRegion),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"wait_time_minutes", "plan_validation"},
			},
		},
	})
}

// func testAccCheckIBMDatabaseInstanceDestroy(s *terraform.State) etc in resource_ibm_database_postgresql_test.go

func testAccCheckIBMDatabaseInstanceCassandraBasic(databaseResourceGroup string, name string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "test_acc" {
		is_default = true
		# name = "%[1]s"
	}

	resource "ibm_database" "%[2]s" {
		resource_group_id            = data.ibm_resource_group.test_acc.id
		name                         = "%[2]s"
		service                      = "databases-for-cassandra"
		plan                         = "enterprise"
		location                     = "%[3]s"
		adminpassword                = "password12"
		members_memory_allocation_mb = 36864
		members_disk_allocation_mb   = 61440
		users {
		  name     = "user123"
		  password = "password12"
		}
		whitelist {
		  address     = "172.168.1.2/32"
		  description = "desc1"
		}

		timeouts {
			create = "480m"
			update = "480m"
			delete = "15m"
		}
	}
				`, databaseResourceGroup, name, acc.IcdDbRegion)
}

func testAccCheckIBMDatabaseInstanceCassandraFullyspecified(databaseResourceGroup string, name string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "test_acc" {
		is_default = true
		# name = "%[1]s"
	}

	resource "ibm_database" "%[2]s" {
		resource_group_id            = data.ibm_resource_group.test_acc.id
		name                         = "%[2]s"
		service                      = "databases-for-cassandra"
		plan                         = "enterprise"
		location                     = "%[3]s"
		adminpassword                = "password12"
		members_memory_allocation_mb = 38400
		members_disk_allocation_mb   = 61440
		users {
		  name     = "user123"
		  password = "password12"
		}
		users {
		  name     = "user124"
		  password = "password12"
		}
		whitelist {
		  address     = "172.168.1.2/32"
		  description = "desc1"
		}
		whitelist {
		  address     = "172.168.1.1/32"
		  description = "desc"
		}

		timeouts {
			create = "480m"
			update = "480m"
			delete = "15m"
		}
	}
				`, databaseResourceGroup, name, acc.IcdDbRegion)
}

func testAccCheckIBMDatabaseInstanceCassandraReduced(databaseResourceGroup string, name string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "test_acc" {
		is_default = true
		# name = "%[1]s"
	}

	resource "ibm_database" "%[2]s" {
		resource_group_id            = data.ibm_resource_group.test_acc.id
		name                         = "%[2]s"
		service                      = "databases-for-cassandra"
		plan                         = "enterprise"
		location                     = "%[3]s"
		adminpassword                = "password12"
		members_memory_allocation_mb = 36864
		members_disk_allocation_mb   = 61440

		timeouts {
			create = "480m"
			update = "480m"
			delete = "15m"
		}
	}
				`, databaseResourceGroup, name, acc.IcdDbRegion)
}

func testAccCheckIBMDatabaseInstanceCassandraNodeBasic(databaseResourceGroup string, name string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "test_acc" {
		is_default = true
		# name = "%[1]s"
	}

	resource "ibm_database" "%[2]s" {
		resource_group_id            = data.ibm_resource_group.test_acc.id
		name                         = "%[2]s"
		service                      = "databases-for-cassandra"
		plan                         = "enterprise"
		location                     = "%[3]s"
		adminpassword                = "password12"
		node_count                   = 3
		node_memory_allocation_mb    = 12288
		node_disk_allocation_mb      = 20480
		node_cpu_allocation_count    = 6

		users {
		  name     = "user123"
		  password = "password12"
		}
		whitelist {
		  address     = "172.168.1.2/32"
		  description = "desc1"
		}

		timeouts {
			create = "480m"
			update = "480m"
			delete = "15m"
		}
	}
				`, databaseResourceGroup, name, acc.IcdDbRegion)
}

func testAccCheckIBMDatabaseInstanceCassandraNodeFullyspecified(databaseResourceGroup string, name string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "test_acc" {
		is_default = true
		# name = "%[1]s"
	}

	resource "ibm_database" "%[2]s" {
		resource_group_id            = data.ibm_resource_group.test_acc.id
		name                         = "%[2]s"
		service                      = "databases-for-cassandra"
		plan                         = "enterprise"
		location                     = "%[3]s"
		version                      = "5.1"
		adminpassword                = "password12"
		node_count                   = 3
		node_memory_allocation_mb    = 12416
		node_disk_allocation_mb      = 20480
		node_cpu_allocation_count    = 6

		users {
			name     = "user123"
			password = "password12"
		}
		users {
			name     = "user124"
			password = "password12"
		}
		whitelist {
			address     = "172.168.1.2/32"
			description = "desc1"
		}
		whitelist {
			address     = "172.168.1.1/32"
			description = "desc"
		}

		timeouts {
			create = "480m"
			update = "480m"
			delete = "15m"
		}
	}
				`, databaseResourceGroup, name, acc.IcdDbRegion)
}

func testAccCheckIBMDatabaseInstanceCassandraNodeReduced(databaseResourceGroup string, name string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "test_acc" {
		is_default = true
		# name = "%[1]s"
	}

	resource "ibm_database" "%[2]s" {
		resource_group_id            = data.ibm_resource_group.test_acc.id
		name                         = "%[2]s"
		service                      = "databases-for-cassandra"
		plan                         = "enterprise"
		location                     = "%[3]s"
		version                      = "5.1"
		adminpassword                = "password12"
		node_count                   = 3
		node_memory_allocation_mb    = 12288
		node_disk_allocation_mb      = 20480
		node_cpu_allocation_count    = 6

		timeouts {
			create = "480m"
			update = "480m"
			delete = "15m"
		}
	}
				`, databaseResourceGroup, name, acc.IcdDbRegion)
}

func testAccCheckIBMDatabaseInstanceCassandraNodeScaleOut(databaseResourceGroup string, name string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "test_acc" {
		is_default = true
		# name = "%[1]s"
	}

	resource "ibm_database" "%[2]s" {
		resource_group_id            = data.ibm_resource_group.test_acc.id
		name                         = "%[2]s"
		service                      = "databases-for-cassandra"
		plan                         = "enterprise"
		location                     = "%[3]s"
		adminpassword                = "password12"
		node_count                   = 4
		node_memory_allocation_mb    = 12288
		node_disk_allocation_mb      = 20480
		node_cpu_allocation_count    = 6

		timeouts {
			create = "480m"
			update = "480m"
			delete = "15m"
		}
	}
				`, databaseResourceGroup, name, acc.IcdDbRegion)
}

func testAccCheckIBMDatabaseInstanceCassandraGroupBasic(databaseResourceGroup string, name string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "test_acc" {
		is_default = true
		# name = "%[1]s"
	}

	resource "ibm_database" "%[2]s" {
		resource_group_id            = data.ibm_resource_group.test_acc.id
		name                         = "%[2]s"
		service                      = "databases-for-cassandra"
		plan                         = "enterprise"
		version                      = "5.1"
		location                     = "%[3]s"
		adminpassword                = "password12"

		group {
			group_id = "member"
			members {
				allocation_count = 3
			}
			memory {
				allocation_mb = 12288
			}
			 disk {
				allocation_mb = 20480
			}
			cpu {
				allocation_count = 6
			}
		}
		users {
			name     = "user123"
			password = "password12"
		}
		whitelist {
			address     = "172.168.1.2/32"
			description = "desc1"
		}

		timeouts {
			create = "480m"
			update = "480m"
			delete = "15m"
		}
	}
				`, databaseResourceGroup, name, acc.IcdDbRegion)
}

func testAccCheckIBMDatabaseInstanceCassandraGroupFullyspecified(databaseResourceGroup string, name string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "test_acc" {
		is_default = true
		# name = "%[1]s"
	}

	resource "ibm_database" "%[2]s" {
		resource_group_id            = data.ibm_resource_group.test_acc.id
		name                         = "%[2]s"
		service                      = "databases-for-cassandra"
		plan                         = "enterprise"
		version                      = "5.1"
		location                     = "%[3]s"
		adminpassword                = "password12"

		group {
			group_id = "member"
			members {
				allocation_count = 3
			}
			memory {
				allocation_mb = 12416
			}
			disk {
				allocation_mb = 20480
			}
			cpu {
				allocation_count = 6
			}
		}
		users {
			name     = "user123"
			password = "password12"
		}
		users {
			name     = "user124"
			password = "password12"
		}
		whitelist {
			address     = "172.168.1.2/32"
			description = "desc1"
		}
		whitelist {
			address     = "172.168.1.1/32"
			description = "desc"
		}

		timeouts {
			create = "480m"
			update = "480m"
			delete = "15m"
		}
	}

				`, databaseResourceGroup, name, acc.IcdDbRegion)
}

func testAccCheckIBMDatabaseInstanceCassandraGroupReduced(databaseResourceGroup string, name string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "test_acc" {
		is_default = true
		# name = "%[1]s"
	}

	resource "ibm_database" "%[2]s" {
		resource_group_id            = data.ibm_resource_group.test_acc.id
		name                         = "%[2]s"
		service                      = "databases-for-cassandra"
		plan                         = "enterprise"
		version                      = "5.1"
		location                     = "%[3]s"
		adminpassword                = "password12"
		group {
		group_id = "member"
			members {
				allocation_count = 3
			}
			memory {
				allocation_mb = 12288
			}
			disk {
				allocation_mb = 20480
			}
			cpu {
				allocation_count = 6
			}
		}

		timeouts {
			create = "480m"
			update = "480m"
			delete = "15m"
		}
	}
				`, databaseResourceGroup, name, acc.IcdDbRegion)
}

func testAccCheckIBMDatabaseInstanceCassandraGroupScaleOut(databaseResourceGroup string, name string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "test_acc" {
		is_default = true
		# name = "%[1]s"
	}

	resource "ibm_database" "%[2]s" {
		resource_group_id            = data.ibm_resource_group.test_acc.id
		name                         = "%[2]s"
		service                      = "databases-for-cassandra"
		plan                         = "enterprise"
		version                      = "5.1"
		location                     = "%[3]s"
		adminpassword                = "password12"

		group {
		group_id = "member"
			members {
				allocation_count = 4
			}
			memory {
				allocation_mb = 12288
			}
			disk {
				allocation_mb = 20480
			}
			cpu {
				allocation_count = 6
			}
		}

		group {
			group_id = "search"
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
				`, databaseResourceGroup, name, acc.IcdDbRegion)
}

func testAccCheckIBMDatabaseInstanceCassandraImport(databaseResourceGroup string, name string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "test_acc" {
		is_default = true
		# name = "%[1]s"
	}

	resource "ibm_database" "%[2]s" {
		resource_group_id = data.ibm_resource_group.test_acc.id
		name              = "%[2]s"
		service           = "databases-for-cassandra"
		plan              = "enterprise"
		location          = "%[3]s"

		timeouts {
			create = "480m"
			update = "480m"
			delete = "15m"
		}

	}
				`, databaseResourceGroup, name, acc.IcdDbRegion)
}
