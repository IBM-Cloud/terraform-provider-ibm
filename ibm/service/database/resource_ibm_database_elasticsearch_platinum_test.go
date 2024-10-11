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
					resource.TestCheckResourceAttr(name, "location", acc.Region()),
					resource.TestCheckResourceAttr(name, "adminuser", "admin"),
					resource.TestCheckResourceAttr(name, "allowlist.#", "1"),
					resource.TestCheckResourceAttr(name, "users.#", "1"),
				),
			},
			{
				Config: testAccCheckIBMDatabaseInstanceElasticsearchPlatinumFullyspecified(databaseResourceGroup, testName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(name, &databaseInstanceOne),
					resource.TestCheckResourceAttr(name, "name", testName),
					resource.TestCheckResourceAttr(name, "service", "databases-for-elasticsearch"),
					resource.TestCheckResourceAttr(name, "plan", "platinum"),
					resource.TestCheckResourceAttr(name, "location", acc.Region()),
					resource.TestCheckResourceAttr(name, "allowlist.#", "2"),
					resource.TestCheckResourceAttr(name, "users.#", "2"),
				),
			},
			{
				Config: testAccCheckIBMDatabaseInstanceElasticsearchPlatinumReduced(databaseResourceGroup, testName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(name, &databaseInstanceOne),
					resource.TestCheckResourceAttr(name, "name", testName),
					resource.TestCheckResourceAttr(name, "service", "databases-for-elasticsearch"),
					resource.TestCheckResourceAttr(name, "plan", "platinum"),
					resource.TestCheckResourceAttr(name, "location", acc.Region()),
					resource.TestCheckResourceAttr(name, "groups.0.memory.0.allocation_mb", "49152"),
					resource.TestCheckResourceAttr(name, "groups.0.disk.0.allocation_mb", "18432"),
					resource.TestCheckResourceAttr(name, "allowlist.#", "0"),
					resource.TestCheckResourceAttr(name, "users.#", "0"),
				),
			},
			{
				Config: testAccCheckIBMDatabaseInstanceElasticsearchPlatinumGroupMigration(databaseResourceGroup, testName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(name, &databaseInstanceOne),
					resource.TestCheckResourceAttr(name, "name", testName),
					resource.TestCheckResourceAttr(name, "service", "databases-for-elasticsearch"),
					resource.TestCheckResourceAttr(name, "plan", "platinum"),
					resource.TestCheckResourceAttr(name, "location", acc.Region()),
					resource.TestCheckResourceAttr(name, "groups.0.memory.0.allocation_mb", "49152"),
					resource.TestCheckResourceAttr(name, "groups.0.disk.0.allocation_mb", "18432"),
					resource.TestCheckResourceAttr(name, "whitelist.#", "0"),
					resource.TestCheckResourceAttr(name, "users.#", "0"),
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
					resource.TestCheckResourceAttr(name, "location", acc.Region()),
					resource.TestCheckResourceAttr(name, "adminuser", "admin"),

					resource.TestCheckResourceAttr(name, "allowlist.#", "1"),
					resource.TestCheckResourceAttr(name, "users.#", "1"),
				),
			},
			{
				Config: testAccCheckIBMDatabaseInstanceElasticsearchPlatinumNodeFullyspecified(databaseResourceGroup, testName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(name, &databaseInstanceOne),
					resource.TestCheckResourceAttr(name, "name", testName),
					resource.TestCheckResourceAttr(name, "service", "databases-for-elasticsearch"),
					resource.TestCheckResourceAttr(name, "plan", "platinum"),
					resource.TestCheckResourceAttr(name, "location", acc.Region()),
					resource.TestCheckResourceAttr(name, "allowlist.#", "2"),
					resource.TestCheckResourceAttr(name, "users.#", "2"),
				),
			},
			{
				Config: testAccCheckIBMDatabaseInstanceElasticsearchPlatinumNodeReduced(databaseResourceGroup, testName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(name, &databaseInstanceOne),
					resource.TestCheckResourceAttr(name, "name", testName),
					resource.TestCheckResourceAttr(name, "service", "databases-for-elasticsearch"),
					resource.TestCheckResourceAttr(name, "plan", "platinum"),
					resource.TestCheckResourceAttr(name, "location", acc.Region()),
					resource.TestCheckResourceAttr(name, "allowlist.#", "0"),
					resource.TestCheckResourceAttr(name, "users.#", "0"),
				),
			},
			{
				Config: testAccCheckIBMDatabaseInstanceElasticsearchPlatinumNodeScaleOut(databaseResourceGroup, testName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(name, &databaseInstanceOne),
					resource.TestCheckResourceAttr(name, "name", testName),
					resource.TestCheckResourceAttr(name, "service", "databases-for-elasticsearch"),
					resource.TestCheckResourceAttr(name, "plan", "platinum"),
					resource.TestCheckResourceAttr(name, "location", acc.Region()),
					resource.TestCheckResourceAttr(name, "allowlist.#", "0"),
					resource.TestCheckResourceAttr(name, "users.#", "0"),
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
					resource.TestCheckResourceAttr(name, "location", acc.Region()),
					resource.TestCheckResourceAttr(name, "adminuser", "admin"),

					resource.TestCheckResourceAttr(name, "allowlist.#", "1"),
					resource.TestCheckResourceAttr(name, "users.#", "1"),
				),
			},
			{
				Config: testAccCheckIBMDatabaseInstanceElasticsearchPlatinumGroupFullyspecified(databaseResourceGroup, testName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(name, &databaseInstanceOne),
					resource.TestCheckResourceAttr(name, "name", testName),
					resource.TestCheckResourceAttr(name, "service", "databases-for-elasticsearch"),
					resource.TestCheckResourceAttr(name, "plan", "platinum"),
					resource.TestCheckResourceAttr(name, "location", acc.Region()),
					resource.TestCheckResourceAttr(name, "groups.0.count", "3"),
					resource.TestCheckResourceAttr(name, "groups.0.memory.0.allocation_mb", "49152"),
					resource.TestCheckResourceAttr(name, "groups.0.disk.0.allocation_mb", "18432"),
					resource.TestCheckResourceAttr(name, "groups.0.cpu.0.allocation_count", "12"),
					resource.TestCheckResourceAttr(name, "allowlist.#", "2"),
					resource.TestCheckResourceAttr(name, "users.#", "2"),
				),
			},
			{
				Config: testAccCheckIBMDatabaseInstanceElasticsearchPlatinumGroupReduced(databaseResourceGroup, testName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(name, &databaseInstanceOne),
					resource.TestCheckResourceAttr(name, "name", testName),
					resource.TestCheckResourceAttr(name, "service", "databases-for-elasticsearch"),
					resource.TestCheckResourceAttr(name, "plan", "platinum"),
					resource.TestCheckResourceAttr(name, "location", acc.Region()),
					resource.TestCheckResourceAttr(name, "groups.0.count", "3"),
					resource.TestCheckResourceAttr(name, "groups.0.memory.0.allocation_mb", "49152"),
					resource.TestCheckResourceAttr(name, "groups.0.disk.0.allocation_mb", "18432"),
					resource.TestCheckResourceAttr(name, "groups.0.cpu.0.allocation_count", "12"),
					resource.TestCheckResourceAttr(name, "allowlist.#", "0"),
					resource.TestCheckResourceAttr(name, "users.#", "0"),
				),
			},
			{
				Config: testAccCheckIBMDatabaseInstanceElasticsearchPlatinumGroupScaleOut(databaseResourceGroup, testName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(name, &databaseInstanceOne),
					resource.TestCheckResourceAttr(name, "name", testName),
					resource.TestCheckResourceAttr(name, "service", "databases-for-elasticsearch"),
					resource.TestCheckResourceAttr(name, "plan", "platinum"),
					resource.TestCheckResourceAttr(name, "location", acc.Region()),
					resource.TestCheckResourceAttr(name, "groups.0.count", "4"),
					resource.TestCheckResourceAttr(name, "groups.0.memory.0.allocation_mb", "65536"),
					resource.TestCheckResourceAttr(name, "groups.0.disk.0.allocation_mb", "24576"),
					resource.TestCheckResourceAttr(name, "groups.0.cpu.0.allocation_count", "16"),
					resource.TestCheckResourceAttr(name, "allowlist.#", "0"),
					resource.TestCheckResourceAttr(name, "users.#", "0"),
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
					resource.TestCheckResourceAttr(resourceName, "location", acc.Region()),
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
		adminpassword                = "password12345678"
		service_endpoints            = "public-and-private"
		group {
			group_id = "member"

			host_flavor {
				id = "b3c.4x16.encrypted"
			}
			disk {
			  allocation_mb = 6144
			}
		}
		users {
		  name     = "user123"
		  password = "password12345678"
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
				`, databaseResourceGroup, name, acc.Region())
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
		adminpassword                = "password12345678"
		service_endpoints            = "public-and-private"
		group {
			group_id = "member"

			host_flavor {
				id = "b3c.4x16.encrypted"
			}
		}
		users {
		  name     = "user123"
		  password = "password12345678"
		}
		users {
		  name     = "user124"
		  password = "password12345678"
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

				`, databaseResourceGroup, name, acc.Region())
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
		adminpassword                = "password12345678"
		service_endpoints            = "public-and-private"
		group {
			group_id = "member"

			host_flavor {
				id = "b3c.4x16.encrypted"
			}
		}

		timeouts {
			create = "120m"
			update = "120m"
			delete = "15m"
		}
	}
				`, databaseResourceGroup, name, acc.Region())
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
		adminpassword                = "password12345678"
		service_endpoints            = "public-and-private"

		group {
		  group_id = "member"

		  host_flavor {
			id = "b3c.4x16.encrypted"
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
				`, databaseResourceGroup, name, acc.Region())
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
		adminpassword                = "password12345678"
		service_endpoints            = "public-and-private"
		group {
			group_id = "member"
			members {
			  allocation_count = 3
			}
			host_flavor {
				id = "b3c.4x16.encrypted"
			}
			disk {
			  allocation_mb = 5120
			}
		}

		users {
		  name     = "user123"
		  password = "password12345678"
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
				`, databaseResourceGroup, name, acc.Region())
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
		adminpassword                = "password12345678"
		service_endpoints            = "public-and-private"
		group {
			group_id = "member"
			members {
			  allocation_count = 3
			}
			host_flavor {
				id = "b3c.4x16.encrypted"
			}
			disk {
			  allocation_mb = 6144
			}
		}
		users {
		  name     = "user123"
		  password = "password12345678"
		}
		users {
		  name     = "user124"
		  password = "password12345678"
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
				`, databaseResourceGroup, name, acc.Region())
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
		adminpassword                = "password12345678"
		service_endpoints            = "public-and-private"
		group {
			group_id = "member"
			members {
			  allocation_count = 3
			}
			host_flavor {
				id = "b3c.4x16.encrypted"
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
				`, databaseResourceGroup, name, acc.Region())
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
		adminpassword                = "password12345678"
		service_endpoints            = "public-and-private"
		group {
			group_id = "member"
			members {
			  allocation_count = 4
			}
			host_flavor {
				id = "b3c.4x16.encrypted"
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
				`, databaseResourceGroup, name, acc.Region())
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
		adminpassword                = "password12345678"
		service_endpoints            = "public-and-private"

		group {
			group_id = "member"
			members {
				allocation_count = 3
			}
			host_flavor {
				id = "b3c.4x16.encrypted"
			}
			disk {
				allocation_mb = 5120
			}
		}

		users {
		  name     = "user123"
		  password = "password12345678"
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
				`, databaseResourceGroup, name, acc.Region())
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
		adminpassword                = "password12345678"
		service_endpoints            = "public-and-private"

		group {
		  group_id = "member"
		  members {
		    allocation_count = 3
		  }
		  host_flavor {
			id = "b3c.4x16.encrypted"
		  }
		  disk {
		    allocation_mb = 6144
		  }
		}
		users {
		  name     = "user123"
		  password = "password12345678"
		}
		users {
		  name     = "user124"
		  password = "password12345678"
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

				`, databaseResourceGroup, name, acc.Region())
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
		adminpassword                = "password12345678"
		service_endpoints            = "public-and-private"

		group {
		  group_id = "member"
		  members {
		    allocation_count = 3
		  }
		  host_flavor {
			id = "b3c.4x16.encrypted"
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
				`, databaseResourceGroup, name, acc.Region())
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
		adminpassword                = "password12345678"
		service_endpoints            = "public-and-private"

		group {
		  group_id = "member"
		  members {
		    allocation_count = 4
		  }
		  host_flavor {
			id = "b3c.4x16.encrypted"
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
				`, databaseResourceGroup, name, acc.Region())
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
		service_endpoints            = "public-and-private"

		timeouts {
			create = "120m"
			update = "120m"
			delete = "15m"
		}
	}

				`, databaseResourceGroup, name, acc.Region())
}
