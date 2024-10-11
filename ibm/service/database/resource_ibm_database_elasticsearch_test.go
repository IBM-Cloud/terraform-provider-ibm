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

func TestAccIBMDatabaseInstance_Elasticsearch_Basic(t *testing.T) {
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
				Config: testAccCheckIBMDatabaseInstanceElasticsearchBasic(databaseResourceGroup, testName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(name, &databaseInstanceOne),
					resource.TestCheckResourceAttr(name, "name", testName),
					resource.TestCheckResourceAttr(name, "service", "databases-for-elasticsearch"),
					resource.TestCheckResourceAttr(name, "plan", "standard"),
					resource.TestCheckResourceAttr(name, "location", acc.Region()),
					resource.TestCheckResourceAttr(name, "adminuser", "admin"),
					resource.TestCheckResourceAttr(name, "allowlist.#", "1"),
					resource.TestCheckResourceAttr(name, "users.#", "1"),
				),
			},
			{
				Config: testAccCheckIBMDatabaseInstanceElasticsearchFullyspecified(databaseResourceGroup, testName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(name, &databaseInstanceOne),
					resource.TestCheckResourceAttr(name, "name", testName),
					resource.TestCheckResourceAttr(name, "service", "databases-for-elasticsearch"),
					resource.TestCheckResourceAttr(name, "plan", "standard"),
					resource.TestCheckResourceAttr(name, "location", acc.Region()),
					resource.TestCheckResourceAttr(name, "allowlist.#", "2"),
					resource.TestCheckResourceAttr(name, "users.#", "2"),
				),
			},
			{
				Config: testAccCheckIBMDatabaseInstanceElasticsearchReduced(databaseResourceGroup, testName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(name, &databaseInstanceOne),
					resource.TestCheckResourceAttr(name, "name", testName),
					resource.TestCheckResourceAttr(name, "service", "databases-for-elasticsearch"),
					resource.TestCheckResourceAttr(name, "plan", "standard"),
					resource.TestCheckResourceAttr(name, "location", acc.Region()),
					resource.TestCheckResourceAttr(name, "allowlist.#", "0"),
					resource.TestCheckResourceAttr(name, "users.#", "0"),
				),
			},
			{
				Config: testAccCheckIBMDatabaseInstanceElasticsearchGroupMigration(databaseResourceGroup, testName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(name, &databaseInstanceOne),
					resource.TestCheckResourceAttr(name, "name", testName),
					resource.TestCheckResourceAttr(name, "service", "databases-for-elasticsearch"),
					resource.TestCheckResourceAttr(name, "plan", "standard"),
					resource.TestCheckResourceAttr(name, "location", acc.Region()),
					resource.TestCheckResourceAttr(name, "groups.0.memory.0.allocation_mb", "12288"),
					resource.TestCheckResourceAttr(name, "groups.0.disk.0.allocation_mb", "18432"),
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

func TestAccIBMDatabaseInstance_Elasticsearch_Node(t *testing.T) {
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
				Config: testAccCheckIBMDatabaseInstanceElasticsearchNodeBasic(databaseResourceGroup, testName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(name, &databaseInstanceOne),
					resource.TestCheckResourceAttr(name, "name", testName),
					resource.TestCheckResourceAttr(name, "service", "databases-for-elasticsearch"),
					resource.TestCheckResourceAttr(name, "plan", "standard"),
					resource.TestCheckResourceAttr(name, "location", acc.Region()),
					resource.TestCheckResourceAttr(name, "adminuser", "admin"),

					resource.TestCheckResourceAttr(name, "allowlist.#", "1"),
					resource.TestCheckResourceAttr(name, "users.#", "1"),
				),
			},
			{
				Config: testAccCheckIBMDatabaseInstanceElasticsearchNodeFullyspecified(databaseResourceGroup, testName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(name, &databaseInstanceOne),
					resource.TestCheckResourceAttr(name, "name", testName),
					resource.TestCheckResourceAttr(name, "service", "databases-for-elasticsearch"),
					resource.TestCheckResourceAttr(name, "plan", "standard"),
					resource.TestCheckResourceAttr(name, "location", acc.Region()),
					resource.TestCheckResourceAttr(name, "allowlist.#", "2"),
					resource.TestCheckResourceAttr(name, "users.#", "2"),
				),
			},
			{
				Config: testAccCheckIBMDatabaseInstanceElasticsearchNodeReduced(databaseResourceGroup, testName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(name, &databaseInstanceOne),
					resource.TestCheckResourceAttr(name, "name", testName),
					resource.TestCheckResourceAttr(name, "service", "databases-for-elasticsearch"),
					resource.TestCheckResourceAttr(name, "plan", "standard"),
					resource.TestCheckResourceAttr(name, "location", acc.Region()),
					resource.TestCheckResourceAttr(name, "allowlist.#", "0"),
					resource.TestCheckResourceAttr(name, "users.#", "0"),
				),
			},
			{
				Config: testAccCheckIBMDatabaseInstanceElasticsearchNodeScaleOut(databaseResourceGroup, testName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(name, &databaseInstanceOne),
					resource.TestCheckResourceAttr(name, "name", testName),
					resource.TestCheckResourceAttr(name, "service", "databases-for-elasticsearch"),
					resource.TestCheckResourceAttr(name, "plan", "standard"),
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

func TestAccIBMDatabaseInstance_Elasticsearch_Group(t *testing.T) {
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
				Config: testAccCheckIBMDatabaseInstanceElasticsearchGroupBasic(databaseResourceGroup, testName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(name, &databaseInstanceOne),
					resource.TestCheckResourceAttr(name, "name", testName),
					resource.TestCheckResourceAttr(name, "service", "databases-for-elasticsearch"),
					resource.TestCheckResourceAttr(name, "plan", "standard"),
					resource.TestCheckResourceAttr(name, "location", acc.Region()),
					resource.TestCheckResourceAttr(name, "adminuser", "admin"),

					resource.TestCheckResourceAttr(name, "allowlist.#", "1"),
					resource.TestCheckResourceAttr(name, "users.#", "1"),
				),
			},
			{
				Config: testAccCheckIBMDatabaseInstanceElasticsearchGroupFullyspecified(databaseResourceGroup, testName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(name, &databaseInstanceOne),
					resource.TestCheckResourceAttr(name, "name", testName),
					resource.TestCheckResourceAttr(name, "service", "databases-for-elasticsearch"),
					resource.TestCheckResourceAttr(name, "plan", "standard"),
					resource.TestCheckResourceAttr(name, "location", acc.Region()),
					resource.TestCheckResourceAttr(name, "groups.0.count", "3"),
					resource.TestCheckResourceAttr(name, "groups.0.memory.0.allocation_mb", "12288"),
					resource.TestCheckResourceAttr(name, "groups.0.disk.0.allocation_mb", "18432"),
					resource.TestCheckResourceAttr(name, "groups.0.cpu.0.allocation_count", "9"),
					resource.TestCheckResourceAttr(name, "allowlist.#", "2"),
					resource.TestCheckResourceAttr(name, "users.#", "2"),
				),
			},
			{
				Config: testAccCheckIBMDatabaseInstanceElasticsearchGroupReduced(databaseResourceGroup, testName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(name, &databaseInstanceOne),
					resource.TestCheckResourceAttr(name, "name", testName),
					resource.TestCheckResourceAttr(name, "service", "databases-for-elasticsearch"),
					resource.TestCheckResourceAttr(name, "plan", "standard"),
					resource.TestCheckResourceAttr(name, "location", acc.Region()),
					resource.TestCheckResourceAttr(name, "groups.0.count", "3"),
					resource.TestCheckResourceAttr(name, "groups.0.memory.0.allocation_mb", "12288"),
					resource.TestCheckResourceAttr(name, "groups.0.disk.0.allocation_mb", "18432"),
					resource.TestCheckResourceAttr(name, "groups.0.cpu.0.allocation_count", "9"),
					resource.TestCheckResourceAttr(name, "allowlist.#", "0"),
					resource.TestCheckResourceAttr(name, "users.#", "0"),
				),
			},
			{
				Config: testAccCheckIBMDatabaseInstanceElasticsearchGroupScaleOut(databaseResourceGroup, testName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(name, &databaseInstanceOne),
					resource.TestCheckResourceAttr(name, "name", testName),
					resource.TestCheckResourceAttr(name, "service", "databases-for-elasticsearch"),
					resource.TestCheckResourceAttr(name, "plan", "standard"),
					resource.TestCheckResourceAttr(name, "location", acc.Region()),
					resource.TestCheckResourceAttr(name, "groups.0.count", "4"),
					resource.TestCheckResourceAttr(name, "groups.0.memory.0.allocation_mb", "15360"),
					resource.TestCheckResourceAttr(name, "groups.0.disk.0.allocation_mb", "24576"),
					resource.TestCheckResourceAttr(name, "groups.0.cpu.0.allocation_count", "12"),
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

func TestAccIBMDatabaseInstanceElasticsearchImport(t *testing.T) {
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
				Config: testAccCheckIBMDatabaseInstanceElasticsearchImport(databaseResourceGroup, serviceName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(resourceName, &databaseInstanceOne),
					resource.TestCheckResourceAttr(resourceName, "name", serviceName),
					resource.TestCheckResourceAttr(resourceName, "service", "databases-for-elasticsearch"),
					resource.TestCheckResourceAttr(resourceName, "plan", "standard"),
					resource.TestCheckResourceAttr(resourceName, "location", acc.Region()),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"wait_time_minutes"},
			},
		},
	})
}

// func testAccCheckIBMDatabaseInstanceDestroy(s *terraform.State) etc in resource_ibm_database_postgresql_test.go

func testAccCheckIBMDatabaseInstanceElasticsearchBasic(databaseResourceGroup string, name string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "test_acc" {
		is_default = true
		# name = "%[1]s"
	}

	resource "ibm_database" "%[2]s" {
		resource_group_id            = data.ibm_resource_group.test_acc.id
		name                         = "%[2]s"
		service                      = "databases-for-elasticsearch"
		plan                         = "standard"
		location                     = "%[3]s"
		adminpassword                = "password12345678"
		service_endpoints            = "public"
		users {
		  name     = "user123"
		  password = "password12345678"
		}
		allowlist {
		  address     = "172.168.1.2/32"
		  description = "desc1"
		}

		group {
			group_id = "member"
			memory {
				allocation_mb = 4096
			}
			host_flavor {
				id = "multitenant"
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

func testAccCheckIBMDatabaseInstanceElasticsearchFullyspecified(databaseResourceGroup string, name string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "test_acc" {
		is_default = true
		# name = "%[1]s"
	}

	resource "ibm_database" "%[2]s" {
		resource_group_id            = data.ibm_resource_group.test_acc.id
		name                         = "%[2]s"
		service                      = "databases-for-elasticsearch"
		plan                         = "standard"
		location                     = "%[3]s"
		adminpassword                = "password12345678"
		service_endpoints            = "public"
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

		group {
			group_id = "member"
			memory {
				allocation_mb = 4096
			}
			host_flavor {
				id = "multitenant"
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

func testAccCheckIBMDatabaseInstanceElasticsearchReduced(databaseResourceGroup string, name string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "test_acc" {
		is_default = true
		# name = "%[1]s"
	}

	resource "ibm_database" "%[2]s" {
		resource_group_id            = data.ibm_resource_group.test_acc.id
		name                         = "%[2]s"
		service                      = "databases-for-elasticsearch"
		plan                         = "standard"
		location                     = "%[3]s"
		adminpassword                = "password12345678"
		service_endpoints            = "public"

		group {
			group_id = "member"
			memory {
				allocation_mb = 4096
			}
			host_flavor {
				id = "multitenant"
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

func testAccCheckIBMDatabaseInstanceElasticsearchGroupMigration(databaseResourceGroup string, name string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "test_acc" {
		is_default = true
		# name = "%[1]s"
	}

	resource "ibm_database" "%[2]s" {
		resource_group_id            = data.ibm_resource_group.test_acc.id
		name                         = "%[2]s"
		service                      = "databases-for-elasticsearch"
		plan                         = "standard"
		location                     = "%[3]s"
		adminpassword                = "password12345678"
		service_endpoints            = "public"

		group {
		  group_id = "member"

		  memory {
		    allocation_mb = 4096
		  }
		  host_flavor {
			  id = "multitenant"
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

func testAccCheckIBMDatabaseInstanceElasticsearchNodeBasic(databaseResourceGroup string, name string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "test_acc" {
		is_default = true
		# name = "%[1]s"
	}

	resource "ibm_database" "%[2]s" {
		resource_group_id            = data.ibm_resource_group.test_acc.id
		name                         = "%[2]s"
		service                      = "databases-for-elasticsearch"
		plan                         = "standard"
		location                     = "%[3]s"
		adminpassword                = "password12345678"
		service_endpoints            = "public"

		group {
			group_id = "member"
			members {
			  allocation_count = 3
			}
			memory {
			  allocation_mb = 4096
			}
			disk {
			  allocation_mb = 5120
			}
			cpu {
			  allocation_count = 3
			}
			host_flavor {
				id = "multitenant"
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

func testAccCheckIBMDatabaseInstanceElasticsearchNodeFullyspecified(databaseResourceGroup string, name string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "test_acc" {
		is_default = true
		# name = "%[1]s"
	}

	resource "ibm_database" "%[2]s" {
		resource_group_id            = data.ibm_resource_group.test_acc.id
		name                         = "%[2]s"
		service                      = "databases-for-elasticsearch"
		plan                         = "standard"
		location                     = "%[3]s"
		adminpassword                = "password12345678"
		service_endpoints            = "public"
		group {
			group_id = "member"
			members {
			  allocation_count = 3
			}
			memory {
			  allocation_mb = 5120
			}
			disk {
			  allocation_mb = 6144
			}
			cpu {
			  allocation_count = 3
			}
			host_flavor {
				id = "multitenant"
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

func testAccCheckIBMDatabaseInstanceElasticsearchNodeReduced(databaseResourceGroup string, name string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "test_acc" {
		is_default = true
		# name = "%[1]s"
	}

	resource "ibm_database" "%[2]s" {
		resource_group_id            = data.ibm_resource_group.test_acc.id
		name                         = "%[2]s"
		service                      = "databases-for-elasticsearch"
		plan                         = "standard"
		location                     = "%[3]s"
		adminpassword                = "password12345678"
		service_endpoints            = "public"
		group {
			group_id = "member"
			members {
			  allocation_count = 3
			}
			memory {
			  allocation_mb = 4096
			}
			disk {
			  allocation_mb = 6144
			}
			cpu {
			  allocation_count = 3
			}
			host_flavor {
				id = "multitenant"
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

func testAccCheckIBMDatabaseInstanceElasticsearchNodeScaleOut(databaseResourceGroup string, name string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "test_acc" {
		is_default = true
		# name = "%[1]s"
	}

	resource "ibm_database" "%[2]s" {
		resource_group_id            = data.ibm_resource_group.test_acc.id
		name                         = "%[2]s"
		service                      = "databases-for-elasticsearch"
		plan                         = "standard"
		location                     = "%[3]s"
		adminpassword                = "password12345678"
		service_endpoints            = "public"
		group {
			group_id = "member"
			members {
			  allocation_count = 4
			}
			memory {
			  allocation_mb = 4096
			}
			disk {
			  allocation_mb = 6144
			}
			cpu {
			  allocation_count = 3
			}
			host_flavor {
				id = "multitenant"
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

func testAccCheckIBMDatabaseInstanceElasticsearchGroupBasic(databaseResourceGroup string, name string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "test_acc" {
		is_default = true
		# name = "%[1]s"
	}

	resource "ibm_database" "%[2]s" {
		resource_group_id            = data.ibm_resource_group.test_acc.id
		name                         = "%[2]s"
		service                      = "databases-for-elasticsearch"
		plan                         = "standard"
		location                     = "%[3]s"
		adminpassword                = "password12345678"
		service_endpoints            = "public"

		group {
			group_id = "member"
			members {
				allocation_count = 3
			}
			memory {
				allocation_mb = 4096
			}
			disk {
				allocation_mb = 5120
			}
			cpu {
				allocation_count = 3
			}
			host_flavor {
				id = "multitenant"
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

func testAccCheckIBMDatabaseInstanceElasticsearchGroupFullyspecified(databaseResourceGroup string, name string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "test_acc" {
		is_default = true
		# name = "%[1]s"
	}

	resource "ibm_database" "%[2]s" {
		resource_group_id            = data.ibm_resource_group.test_acc.id
		name                         = "%[2]s"
		service                      = "databases-for-elasticsearch"
		plan                         = "standard"
		location                     = "%[3]s"
		adminpassword                = "password12345678"
		service_endpoints            = "public"

		group {
		  group_id = "member"
		  members {
		    allocation_count = 3
		  }
		  memory {
		    allocation_mb = 4096
		  }
		  disk {
		    allocation_mb = 6144
		  }
		  cpu {
		    allocation_count = 3
		  }
		  host_flavor {
			  id = "multitenant"
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

func testAccCheckIBMDatabaseInstanceElasticsearchGroupReduced(databaseResourceGroup string, name string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "test_acc" {
		is_default = true
		# name = "%[1]s"
	}

	resource "ibm_database" "%[2]s" {
		resource_group_id            = data.ibm_resource_group.test_acc.id
		name                         = "%[2]s"
		service                      = "databases-for-elasticsearch"
		plan                         = "standard"
		location                     = "%[3]s"
		adminpassword                = "password12345678"
		service_endpoints            = "public"

		group {
		  group_id = "member"
		  members {
		    allocation_count = 3
		  }
		  memory {
		    allocation_mb = 4096
		  }
		  disk {
		    allocation_mb = 6144
		  }
		  cpu {
		    allocation_count = 3
		  }
		  host_flavor {
			  id = "multitenant"
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

func testAccCheckIBMDatabaseInstanceElasticsearchGroupScaleOut(databaseResourceGroup string, name string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "test_acc" {
		is_default = true
		# name = "%[1]s"
	}

	resource "ibm_database" "%[2]s" {
		resource_group_id            = data.ibm_resource_group.test_acc.id
		name                         = "%[2]s"
		service                      = "databases-for-elasticsearch"
		plan                         = "standard"
		location                     = "%[3]s"
		adminpassword                = "password12345678"
		service_endpoints            = "public"

		group {
		  group_id = "member"
		  members {
		    allocation_count = 4
		  }
		  memory {
		    allocation_mb = 4096
		  }
		  disk {
		    allocation_mb = 6144
		  }
		  cpu {
		    allocation_count = 3
		  }
		  host_flavor {
			  id = "multitenant"
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

func testAccCheckIBMDatabaseInstanceElasticsearchImport(databaseResourceGroup string, name string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "test_acc" {
		is_default = true
		# name = "%[1]s"
	}

	resource "ibm_database" "%[2]s" {
		resource_group_id = data.ibm_resource_group.test_acc.id
		name              = "%[2]s"
		service           = "databases-for-elasticsearch"
		plan              = "standard"
		location          = "%[3]s"
		service_endpoints = "public"

		timeouts {
			create = "120m"
			update = "120m"
			delete = "15m"
		}
	}

				`, databaseResourceGroup, name, acc.Region())
}
