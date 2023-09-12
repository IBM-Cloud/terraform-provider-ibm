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

func TestAccIBMDatabaseInstance_ElasticsearchPlatinum_Basic(t *testing.T) {
	t.Parallel()
	databaseResourceGroup := "default"
	var databaseInstanceOne string
	rnd := fmt.Sprintf("tf-Es-%d", acctest.RandIntRange(10, 100))
	testName := rnd
	name := "ibm_database." + testName

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMDatabaseInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMDatabaseInstanceElasticsearchPlatinumBasic(databaseResourceGroup, testName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(name, &databaseInstanceOne),
					resource.TestCheckResourceAttr(name, "name", testName),
					resource.TestCheckResourceAttr(name, "service", "databases-for-elasticsearch"),
					resource.TestCheckResourceAttr(name, "plan", "platinum"),
					resource.TestCheckResourceAttr(name, "location", acc.IcdDbRegion),
					resource.TestCheckResourceAttr(name, "adminuser", "admin"),
					resource.TestCheckResourceAttr(name, "members_memory_allocation_mb", "3072"),
					resource.TestCheckResourceAttr(name, "members_disk_allocation_mb", "15360"),
					resource.TestCheckResourceAttr(name, "allowlist.#", "1"),
					resource.TestCheckResourceAttr(name, "users.#", "1"),
					resource.TestCheckResourceAttr(name, "connectionstrings.#", "2"),
					resource.TestCheckResourceAttr(name, "connectionstrings.1.name", "admin"),
					resource.TestCheckResourceAttr(name, "connectionstrings.0.hosts.#", "1"),
					resource.TestCheckResourceAttr(name, "connectionstrings.0.database", ""),
				),
			},
			{
				Config: testAccCheckIBMDatabaseInstanceElasticsearchPlatinumFullyspecified(databaseResourceGroup, testName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(name, &databaseInstanceOne),
					resource.TestCheckResourceAttr(name, "name", testName),
					resource.TestCheckResourceAttr(name, "service", "databases-for-elasticsearch"),
					resource.TestCheckResourceAttr(name, "plan", "platinum"),
					resource.TestCheckResourceAttr(name, "location", acc.IcdDbRegion),
					resource.TestCheckResourceAttr(name, "members_memory_allocation_mb", "6144"),
					resource.TestCheckResourceAttr(name, "members_disk_allocation_mb", "18432"),
					resource.TestCheckResourceAttr(name, "allowlist.#", "2"),
					resource.TestCheckResourceAttr(name, "users.#", "2"),
					resource.TestCheckResourceAttr(name, "connectionstrings.#", "3"),
					resource.TestCheckResourceAttr(name, "connectionstrings.2.name", "admin"),
				),
			},
			{
				Config: testAccCheckIBMDatabaseInstanceElasticsearchPlatinumReduced(databaseResourceGroup, testName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(name, &databaseInstanceOne),
					resource.TestCheckResourceAttr(name, "name", testName),
					resource.TestCheckResourceAttr(name, "service", "databases-for-elasticsearch"),
					resource.TestCheckResourceAttr(name, "plan", "platinum"),
					resource.TestCheckResourceAttr(name, "location", acc.IcdDbRegion),
					resource.TestCheckResourceAttr(name, "members_memory_allocation_mb", "3072"),
					resource.TestCheckResourceAttr(name, "members_disk_allocation_mb", "18432"),
					resource.TestCheckResourceAttr(name, "allowlist.#", "0"),
					resource.TestCheckResourceAttr(name, "users.#", "0"),
					resource.TestCheckResourceAttr(name, "connectionstrings.#", "1"),
				),
			},
			{
				Config: testAccCheckIBMDatabaseInstanceElasticsearchPlatinumGroupMigration(databaseResourceGroup, testName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(name, &databaseInstanceOne),
					resource.TestCheckResourceAttr(name, "name", testName),
					resource.TestCheckResourceAttr(name, "service", "databases-for-elasticsearch"),
					resource.TestCheckResourceAttr(name, "plan", "platinum"),
					resource.TestCheckResourceAttr(name, "location", acc.IcdDbRegion),
					resource.TestCheckResourceAttr(name, "groups.0.memory.0.allocation_mb", "3072"),
					resource.TestCheckResourceAttr(name, "groups.0.disk.0.allocation_mb", "18432"),
					resource.TestCheckResourceAttr(name, "whitelist.#", "0"),
					resource.TestCheckResourceAttr(name, "users.#", "0"),
					resource.TestCheckResourceAttr(name, "connectionstrings.#", "1"),
				),
			},
			// {
			// 	ResourceName:      name,
			// 	ImportState:       true,
			// 	ImportStateVerify: true,
			// },
		},
	})
}

func TestAccIBMDatabaseInstance_ElasticsearchPlatinum_Node(t *testing.T) {
	t.Parallel()
	databaseResourceGroup := "default"
	var databaseInstanceOne string
	rnd := fmt.Sprintf("tf-Es-%d", acctest.RandIntRange(10, 100))
	testName := rnd
	name := "ibm_database." + testName

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMDatabaseInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMDatabaseInstanceElasticsearchPlatinumNodeBasic(databaseResourceGroup, testName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(name, &databaseInstanceOne),
					resource.TestCheckResourceAttr(name, "name", testName),
					resource.TestCheckResourceAttr(name, "service", "databases-for-elasticsearch"),
					resource.TestCheckResourceAttr(name, "plan", "platinum"),
					resource.TestCheckResourceAttr(name, "location", acc.IcdDbRegion),
					resource.TestCheckResourceAttr(name, "adminuser", "admin"),
					resource.TestCheckResourceAttr(name, "node_count", "3"),
					resource.TestCheckResourceAttr(name, "node_memory_allocation_mb", "1024"),
					resource.TestCheckResourceAttr(name, "node_disk_allocation_mb", "5120"),
					resource.TestCheckResourceAttr(name, "node_cpu_allocation_count", "3"),

					resource.TestCheckResourceAttr(name, "allowlist.#", "1"),
					resource.TestCheckResourceAttr(name, "users.#", "1"),
					resource.TestCheckResourceAttr(name, "connectionstrings.#", "2"),
					resource.TestCheckResourceAttr(name, "connectionstrings.1.name", "admin"),
					resource.TestCheckResourceAttr(name, "connectionstrings.0.hosts.#", "1"),
					resource.TestCheckResourceAttr(name, "connectionstrings.0.database", ""),
				),
			},
			{
				Config: testAccCheckIBMDatabaseInstanceElasticsearchPlatinumNodeFullyspecified(databaseResourceGroup, testName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(name, &databaseInstanceOne),
					resource.TestCheckResourceAttr(name, "name", testName),
					resource.TestCheckResourceAttr(name, "service", "databases-for-elasticsearch"),
					resource.TestCheckResourceAttr(name, "plan", "platinum"),
					resource.TestCheckResourceAttr(name, "location", acc.IcdDbRegion),
					resource.TestCheckResourceAttr(name, "node_count", "3"),
					resource.TestCheckResourceAttr(name, "node_memory_allocation_mb", "1024"),
					resource.TestCheckResourceAttr(name, "node_disk_allocation_mb", "6144"),
					resource.TestCheckResourceAttr(name, "node_cpu_allocation_count", "3"),
					resource.TestCheckResourceAttr(name, "allowlist.#", "2"),
					resource.TestCheckResourceAttr(name, "users.#", "2"),
					resource.TestCheckResourceAttr(name, "connectionstrings.#", "3"),
					resource.TestCheckResourceAttr(name, "connectionstrings.2.name", "admin"),
				),
			},
			{
				Config: testAccCheckIBMDatabaseInstanceElasticsearchPlatinumNodeReduced(databaseResourceGroup, testName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(name, &databaseInstanceOne),
					resource.TestCheckResourceAttr(name, "name", testName),
					resource.TestCheckResourceAttr(name, "service", "databases-for-elasticsearch"),
					resource.TestCheckResourceAttr(name, "plan", "platinum"),
					resource.TestCheckResourceAttr(name, "location", acc.IcdDbRegion),
					resource.TestCheckResourceAttr(name, "node_count", "3"),
					resource.TestCheckResourceAttr(name, "node_memory_allocation_mb", "1024"),
					resource.TestCheckResourceAttr(name, "node_disk_allocation_mb", "6144"),
					resource.TestCheckResourceAttr(name, "node_cpu_allocation_count", "3"),
					resource.TestCheckResourceAttr(name, "allowlist.#", "0"),
					resource.TestCheckResourceAttr(name, "users.#", "0"),
					resource.TestCheckResourceAttr(name, "connectionstrings.#", "1"),
				),
			},
			{
				Config: testAccCheckIBMDatabaseInstanceElasticsearchPlatinumNodeScaleOut(databaseResourceGroup, testName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(name, &databaseInstanceOne),
					resource.TestCheckResourceAttr(name, "name", testName),
					resource.TestCheckResourceAttr(name, "service", "databases-for-elasticsearch"),
					resource.TestCheckResourceAttr(name, "plan", "platinum"),
					resource.TestCheckResourceAttr(name, "location", acc.IcdDbRegion),
					resource.TestCheckResourceAttr(name, "node_count", "4"),
					resource.TestCheckResourceAttr(name, "node_memory_allocation_mb", "1024"),
					resource.TestCheckResourceAttr(name, "node_disk_allocation_mb", "6144"),
					resource.TestCheckResourceAttr(name, "node_cpu_allocation_count", "3"),
					resource.TestCheckResourceAttr(name, "allowlist.#", "0"),
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

func TestAccIBMDatabaseInstance_ElasticsearchPlatinum_Group(t *testing.T) {
	t.Parallel()
	databaseResourceGroup := "default"
	var databaseInstanceOne string
	rnd := fmt.Sprintf("tf-Es-%d", acctest.RandIntRange(10, 100))
	testName := rnd
	name := "ibm_database." + testName

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMDatabaseInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMDatabaseInstanceElasticsearchPlatinumGroupBasic(databaseResourceGroup, testName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(name, &databaseInstanceOne),
					resource.TestCheckResourceAttr(name, "name", testName),
					resource.TestCheckResourceAttr(name, "service", "databases-for-elasticsearch"),
					resource.TestCheckResourceAttr(name, "plan", "platinum"),
					resource.TestCheckResourceAttr(name, "location", acc.IcdDbRegion),
					resource.TestCheckResourceAttr(name, "adminuser", "admin"),
					resource.TestCheckResourceAttr(name, "node_count", "3"),
					resource.TestCheckResourceAttr(name, "node_memory_allocation_mb", "1024"),
					resource.TestCheckResourceAttr(name, "node_disk_allocation_mb", "5120"),
					resource.TestCheckResourceAttr(name, "node_cpu_allocation_count", "3"),

					resource.TestCheckResourceAttr(name, "allowlist.#", "1"),
					resource.TestCheckResourceAttr(name, "users.#", "1"),
					resource.TestCheckResourceAttr(name, "connectionstrings.#", "2"),
					resource.TestCheckResourceAttr(name, "connectionstrings.1.name", "admin"),
					resource.TestCheckResourceAttr(name, "connectionstrings.0.hosts.#", "1"),
					resource.TestCheckResourceAttr(name, "connectionstrings.0.database", ""),
				),
			},
			{
				Config: testAccCheckIBMDatabaseInstanceElasticsearchPlatinumGroupFullyspecified(databaseResourceGroup, testName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(name, &databaseInstanceOne),
					resource.TestCheckResourceAttr(name, "name", testName),
					resource.TestCheckResourceAttr(name, "service", "databases-for-elasticsearch"),
					resource.TestCheckResourceAttr(name, "plan", "platinum"),
					resource.TestCheckResourceAttr(name, "location", acc.IcdDbRegion),
					resource.TestCheckResourceAttr(name, "node_count", "3"),
					resource.TestCheckResourceAttr(name, "node_memory_allocation_mb", "1024"),
					resource.TestCheckResourceAttr(name, "node_disk_allocation_mb", "6144"),
					resource.TestCheckResourceAttr(name, "node_cpu_allocation_count", "3"),
					resource.TestCheckResourceAttr(name, "groups.0.count", "3"),
					resource.TestCheckResourceAttr(name, "groups.0.memory.0.allocation_mb", "3072"),
					resource.TestCheckResourceAttr(name, "groups.0.disk.0.allocation_mb", "18432"),
					resource.TestCheckResourceAttr(name, "groups.0.cpu.0.allocation_count", "9"),
					resource.TestCheckResourceAttr(name, "allowlist.#", "2"),
					resource.TestCheckResourceAttr(name, "users.#", "2"),
					resource.TestCheckResourceAttr(name, "connectionstrings.#", "3"),
					resource.TestCheckResourceAttr(name, "connectionstrings.2.name", "admin"),
				),
			},
			{
				Config: testAccCheckIBMDatabaseInstanceElasticsearchPlatinumGroupReduced(databaseResourceGroup, testName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(name, &databaseInstanceOne),
					resource.TestCheckResourceAttr(name, "name", testName),
					resource.TestCheckResourceAttr(name, "service", "databases-for-elasticsearch"),
					resource.TestCheckResourceAttr(name, "plan", "platinum"),
					resource.TestCheckResourceAttr(name, "location", acc.IcdDbRegion),
					resource.TestCheckResourceAttr(name, "node_count", "3"),
					resource.TestCheckResourceAttr(name, "node_memory_allocation_mb", "1024"),
					resource.TestCheckResourceAttr(name, "node_disk_allocation_mb", "6144"),
					resource.TestCheckResourceAttr(name, "node_cpu_allocation_count", "3"),
					resource.TestCheckResourceAttr(name, "groups.0.count", "3"),
					resource.TestCheckResourceAttr(name, "groups.0.memory.0.allocation_mb", "3072"),
					resource.TestCheckResourceAttr(name, "groups.0.disk.0.allocation_mb", "18432"),
					resource.TestCheckResourceAttr(name, "groups.0.cpu.0.allocation_count", "9"),
					resource.TestCheckResourceAttr(name, "allowlist.#", "0"),
					resource.TestCheckResourceAttr(name, "users.#", "0"),
					resource.TestCheckResourceAttr(name, "connectionstrings.#", "1"),
				),
			},
			{
				Config: testAccCheckIBMDatabaseInstanceElasticsearchPlatinumGroupScaleOut(databaseResourceGroup, testName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(name, &databaseInstanceOne),
					resource.TestCheckResourceAttr(name, "name", testName),
					resource.TestCheckResourceAttr(name, "service", "databases-for-elasticsearch"),
					resource.TestCheckResourceAttr(name, "plan", "platinum"),
					resource.TestCheckResourceAttr(name, "location", acc.IcdDbRegion),
					resource.TestCheckResourceAttr(name, "node_count", "4"),
					resource.TestCheckResourceAttr(name, "node_memory_allocation_mb", "1024"),
					resource.TestCheckResourceAttr(name, "node_disk_allocation_mb", "6144"),
					resource.TestCheckResourceAttr(name, "node_cpu_allocation_count", "3"),
					resource.TestCheckResourceAttr(name, "groups.0.count", "4"),
					resource.TestCheckResourceAttr(name, "groups.0.memory.0.allocation_mb", "4096"),
					resource.TestCheckResourceAttr(name, "groups.0.disk.0.allocation_mb", "24576"),
					resource.TestCheckResourceAttr(name, "groups.0.cpu.0.allocation_count", "12"),
					resource.TestCheckResourceAttr(name, "allowlist.#", "0"),
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

// TestAccIBMDatabaseInstance_CreateAfterManualDestroy not required as tested by resource_instance tests

func TestAccIBMDatabaseInstanceElasticsearchPlatinumImport(t *testing.T) {
	t.Parallel()
	databaseResourceGroup := "default"
	var databaseInstanceOne string
	serviceName := fmt.Sprintf("tf-Es-%d", acctest.RandIntRange(10, 100))
	//serviceName := "test_acc"
	resourceName := "ibm_database." + serviceName

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMDatabaseInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMDatabaseInstanceElasticsearchPlatinumImport(databaseResourceGroup, serviceName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(resourceName, &databaseInstanceOne),
					resource.TestCheckResourceAttr(resourceName, "name", serviceName),
					resource.TestCheckResourceAttr(resourceName, "service", "databases-for-elasticsearch"),
					resource.TestCheckResourceAttr(resourceName, "plan", "platinum"),
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

func testAccCheckIBMDatabaseInstanceElasticsearchPlatinumBasic(databaseResourceGroup string, name string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "test_acc" {
		is_default = true
		# name = "%[1]s"
	}

	resource "ibm_database" "%[2]s" {
		resource_group_id            = data.ibm_resource_group.test_acc.id
		name                         = "%[2]s"
		service                      = "databases-for-elasticsearch"
		plan                         = "platinum"
		location                     = "%[3]s"
		adminpassword                = "password12"
		members_memory_allocation_mb = 3072
		members_disk_allocation_mb   = 15360
		users {
		  name     = "user123"
		  password = "password12"
		}
		allowlist {
		  address     = "172.168.1.2/32"
		  description = "desc1"
		}

		timeouts {
			create = "120m"
			update = "120m"
			delete = "15m"
		}
	}
				`, databaseResourceGroup, name, acc.IcdDbRegion)
}

func testAccCheckIBMDatabaseInstanceElasticsearchPlatinumFullyspecified(databaseResourceGroup string, name string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "test_acc" {
		is_default = true
		# name = "%[1]s"
	}

	resource "ibm_database" "%[2]s" {
		resource_group_id            = data.ibm_resource_group.test_acc.id
		name                         = "%[2]s"
		service                      = "databases-for-elasticsearch"
		plan                         = "platinum"
		location                     = "%[3]s"
		adminpassword                = "password12"
		members_memory_allocation_mb = 6144
		members_disk_allocation_mb   = 18432
		users {
		  name     = "user123"
		  password = "password12"
		}
		users {
		  name     = "user124"
		  password = "password12"
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
			create = "120m"
			update = "120m"
			delete = "15m"
		}
	}

				`, databaseResourceGroup, name, acc.IcdDbRegion)
}

func testAccCheckIBMDatabaseInstanceElasticsearchPlatinumReduced(databaseResourceGroup string, name string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "test_acc" {
		is_default = true
		# name = "%[1]s"
	}

	resource "ibm_database" "%[2]s" {
		resource_group_id            = data.ibm_resource_group.test_acc.id
		name                         = "%[2]s"
		service                      = "databases-for-elasticsearch"
		plan                         = "platinum"
		location                     = "%[3]s"
		adminpassword                = "password12"
		members_memory_allocation_mb = 3072
		members_disk_allocation_mb   = 18432

		timeouts {
			create = "120m"
			update = "120m"
			delete = "15m"
		}
	}
				`, databaseResourceGroup, name, acc.IcdDbRegion)
}

func testAccCheckIBMDatabaseInstanceElasticsearchPlatinumGroupMigration(databaseResourceGroup string, name string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "test_acc" {
		is_default = true
		# name = "%[1]s"
	}

	resource "ibm_database" "%[2]s" {
		resource_group_id            = data.ibm_resource_group.test_acc.id
		name                         = "%[2]s"
		service                      = "databases-for-elasticsearch"
		plan                         = "platinum"
		location                     = "%[3]s"
		adminpassword                = "password12"

		group {
		  group_id = "member"

		  memory {
		    allocation_mb = 1024
		  }
		  disk {
		    allocation_mb = 6144
		  }
		}

		timeouts {
			create = "120m"
			update = "120m"
			delete = "15m"
		}
	}
				`, databaseResourceGroup, name, acc.IcdDbRegion)
}

func testAccCheckIBMDatabaseInstanceElasticsearchPlatinumNodeBasic(databaseResourceGroup string, name string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "test_acc" {
		is_default = true
		# name = "%[1]s"
	}

	resource "ibm_database" "%[2]s" {
		resource_group_id            = data.ibm_resource_group.test_acc.id
		name                         = "%[2]s"
		service                      = "databases-for-elasticsearch"
		plan                         = "platinum"
		location                     = "%[3]s"
		adminpassword                = "password12"
		node_count					 = 3
		node_memory_allocation_mb    = 1024
		node_disk_allocation_mb      = 5120
        node_cpu_allocation_count    = 3

		users {
		  name     = "user123"
		  password = "password12"
		}
		allowlist {
		  address     = "172.168.1.2/32"
		  description = "desc1"
		}

		timeouts {
			create = "120m"
			update = "120m"
			delete = "15m"
		}
	}
				`, databaseResourceGroup, name, acc.IcdDbRegion)
}

func testAccCheckIBMDatabaseInstanceElasticsearchPlatinumNodeFullyspecified(databaseResourceGroup string, name string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "test_acc" {
		is_default = true
		# name = "%[1]s"
	}

	resource "ibm_database" "%[2]s" {
		resource_group_id            = data.ibm_resource_group.test_acc.id
		name                         = "%[2]s"
		service                      = "databases-for-elasticsearch"
		plan                         = "platinum"
		location                     = "%[3]s"
		adminpassword                = "password12"
		node_count					 = 3
		node_memory_allocation_mb    = 1024
		node_disk_allocation_mb      = 6144
        node_cpu_allocation_count    = 3
		users {
		  name     = "user123"
		  password = "password12"
		}
		users {
		  name     = "user124"
		  password = "password12"
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
			create = "120m"
			update = "120m"
			delete = "15m"
		}
	}
				`, databaseResourceGroup, name, acc.IcdDbRegion)
}

func testAccCheckIBMDatabaseInstanceElasticsearchPlatinumNodeReduced(databaseResourceGroup string, name string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "test_acc" {
		is_default = true
		# name = "%[1]s"
	}

	resource "ibm_database" "%[2]s" {
		resource_group_id            = data.ibm_resource_group.test_acc.id
		name                         = "%[2]s"
		service                      = "databases-for-elasticsearch"
		plan                         = "platinum"
		location                     = "%[3]s"
		adminpassword                = "password12"
		node_count					 = 3
		node_memory_allocation_mb    = 1024
		node_disk_allocation_mb      = 6144
        node_cpu_allocation_count    = 3

		timeouts {
			create = "120m"
			update = "120m"
			delete = "15m"
		}
	}
				`, databaseResourceGroup, name, acc.IcdDbRegion)
}

func testAccCheckIBMDatabaseInstanceElasticsearchPlatinumNodeScaleOut(databaseResourceGroup string, name string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "test_acc" {
		is_default = true
		# name = "%[1]s"
	}

	resource "ibm_database" "%[2]s" {
		resource_group_id            = data.ibm_resource_group.test_acc.id
		name                         = "%[2]s"
		service                      = "databases-for-elasticsearch"
		plan                         = "platinum"
		location                     = "%[3]s"
		adminpassword                = "password12"
        node_count                   = 4
		node_memory_allocation_mb    = 1024
		node_disk_allocation_mb      = 6144
        node_cpu_allocation_count    = 3

		timeouts {
			create = "120m"
			update = "120m"
			delete = "15m"
		}
	}
				`, databaseResourceGroup, name, acc.IcdDbRegion)
}

func testAccCheckIBMDatabaseInstanceElasticsearchPlatinumGroupBasic(databaseResourceGroup string, name string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "test_acc" {
		is_default = true
		# name = "%[1]s"
	}

	resource "ibm_database" "%[2]s" {
		resource_group_id            = data.ibm_resource_group.test_acc.id
		name                         = "%[2]s"
		service                      = "databases-for-elasticsearch"
		plan                         = "platinum"
		location                     = "%[3]s"
		adminpassword                = "password12"

		group {
			group_id = "member"
			members {
				allocation_count = 3
			}
			memory {
				allocation_mb = 1024
			}
			disk {
				allocation_mb = 5120
			}
			cpu {
				allocation_count = 3
			}
		}

		users {
		  name     = "user123"
		  password = "password12"
		}
		allowlist {
		  address     = "172.168.1.2/32"
		  description = "desc1"
		}

		timeouts {
			create = "120m"
			update = "120m"
			delete = "15m"
		}
	}
				`, databaseResourceGroup, name, acc.IcdDbRegion)
}

func testAccCheckIBMDatabaseInstanceElasticsearchPlatinumGroupFullyspecified(databaseResourceGroup string, name string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "test_acc" {
		is_default = true
		# name = "%[1]s"
	}

	resource "ibm_database" "%[2]s" {
		resource_group_id            = data.ibm_resource_group.test_acc.id
		name                         = "%[2]s"
		service                      = "databases-for-elasticsearch"
		plan                         = "platinum"
		location                     = "%[3]s"
		adminpassword                = "password12"

		group {
		  group_id = "member"
		  members {
		    allocation_count = 3
		  }
		  memory {
		    allocation_mb = 1024
		  }
		  disk {
		    allocation_mb = 6144
		  }
		  cpu {
		    allocation_count = 3
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
		allowlist {
		  address     = "172.168.1.2/32"
		  description = "desc1"
		}
		allowlist {
		  address     = "172.168.1.1/32"
		  description = "desc"
		}

		timeouts {
			create = "120m"
			update = "120m"
			delete = "15m"
		}
	}

				`, databaseResourceGroup, name, acc.IcdDbRegion)
}

func testAccCheckIBMDatabaseInstanceElasticsearchPlatinumGroupReduced(databaseResourceGroup string, name string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "test_acc" {
		is_default = true
		# name = "%[1]s"
	}

	resource "ibm_database" "%[2]s" {
		resource_group_id            = data.ibm_resource_group.test_acc.id
		name                         = "%[2]s"
		service                      = "databases-for-elasticsearch"
		plan                         = "platinum"
		location                     = "%[3]s"
		adminpassword                = "password12"

		group {
		  group_id = "member"
		  members {
		    allocation_count = 3
		  }
		  memory {
		    allocation_mb = 1024
		  }
		  disk {
		    allocation_mb = 6144
		  }
		  cpu {
		    allocation_count = 3
		  }
		}

		timeouts {
			create = "120m"
			update = "120m"
			delete = "15m"
		}
	}
				`, databaseResourceGroup, name, acc.IcdDbRegion)
}

func testAccCheckIBMDatabaseInstanceElasticsearchPlatinumGroupScaleOut(databaseResourceGroup string, name string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "test_acc" {
		is_default = true
		# name = "%[1]s"
	}

	resource "ibm_database" "%[2]s" {
		resource_group_id            = data.ibm_resource_group.test_acc.id
		name                         = "%[2]s"
		service                      = "databases-for-elasticsearch"
		plan                         = "platinum"
		location                     = "%[3]s"
		adminpassword                = "password12"

		group {
		  group_id = "member"
		  members {
		    allocation_count = 4
		  }
		  memory {
		    allocation_mb = 1024
		  }
		  disk {
		    allocation_mb = 6144
		  }
		  cpu {
		    allocation_count = 3
		  }
		}
		timeouts {
			create = "120m"
			update = "120m"
			delete = "15m"
		}
	}
				`, databaseResourceGroup, name, acc.IcdDbRegion)
}

func testAccCheckIBMDatabaseInstanceElasticsearchPlatinumImport(databaseResourceGroup string, name string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "test_acc" {
		is_default = true
		# name = "%[1]s"
	}

	resource "ibm_database" "%[2]s" {
		resource_group_id = data.ibm_resource_group.test_acc.id
		name              = "%[2]s"
		service           = "databases-for-elasticsearch"
		plan              = "platinum"
		location          = "%[3]s"

		timeouts {
			create = "120m"
			update = "120m"
			delete = "15m"
		}
	}

				`, databaseResourceGroup, name, acc.IcdDbRegion)
}
