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

func TestAccIBMDatabaseInstanceMongodbBasic(t *testing.T) {
	t.Parallel()
	databaseResourceGroup := "default"
	var databaseInstanceOne string
	rnd := fmt.Sprintf("tf-Mongo-%d", acctest.RandIntRange(10, 100))
	testName := rnd
	name := "ibm_database." + testName

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMDatabaseInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMDatabaseInstanceMongodbBasic(databaseResourceGroup, testName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(name, &databaseInstanceOne),
					resource.TestCheckResourceAttr(name, "name", testName),
					resource.TestCheckResourceAttr(name, "service", "databases-for-mongodb"),
					resource.TestCheckResourceAttr(name, "plan", "standard"),
					resource.TestCheckResourceAttr(name, "location", acc.Region()),
					resource.TestCheckResourceAttr(name, "adminuser", "admin"),
					resource.TestCheckResourceAttr(name, "allowlist.#", "1"),
					resource.TestCheckResourceAttr(name, "users.#", "1"),
					resource.TestCheckResourceAttr(name, "groups.0.memory.0.allocation_mb", "12288"),
					resource.TestCheckResourceAttr(name, "groups.0.disk.0.allocation_mb", "30720"),
				),
			},
			{
				Config: testAccCheckIBMDatabaseInstanceMongodbFullyspecified(databaseResourceGroup, testName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(name, &databaseInstanceOne),
					resource.TestCheckResourceAttr(name, "name", testName),
					resource.TestCheckResourceAttr(name, "service", "databases-for-mongodb"),
					resource.TestCheckResourceAttr(name, "plan", "standard"),
					resource.TestCheckResourceAttr(name, "location", acc.Region()),
					resource.TestCheckResourceAttr(name, "allowlist.#", "2"),
					resource.TestCheckResourceAttr(name, "users.#", "2"),
					resource.TestCheckResourceAttr(name, "groups.0.memory.0.allocation_mb", "15360"),
					resource.TestCheckResourceAttr(name, "groups.0.disk.0.allocation_mb", "30720"),
				),
			},
			{
				Config: testAccCheckIBMDatabaseInstanceMongodbReduced(databaseResourceGroup, testName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(name, &databaseInstanceOne),
					resource.TestCheckResourceAttr(name, "name", testName),
					resource.TestCheckResourceAttr(name, "service", "databases-for-mongodb"),
					resource.TestCheckResourceAttr(name, "plan", "standard"),
					resource.TestCheckResourceAttr(name, "location", acc.Region()),
					resource.TestCheckResourceAttr(name, "groups.0.memory.0.allocation_mb", "12288"),
					resource.TestCheckResourceAttr(name, "groups.0.disk.0.allocation_mb", "30720"),
					resource.TestCheckResourceAttr(name, "allowlist.#", "0"),
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

// TestAccIBMDatabaseInstance_CreateAfterManualDestroy not required as tested by resource_instance tests

func TestAccIBMDatabaseInstanceMongodbImport(t *testing.T) {
	t.Parallel()
	databaseResourceGroup := "default"
	var databaseInstanceOne string
	serviceName := fmt.Sprintf("tf-Mongo-%d", acctest.RandIntRange(10, 100))
	//serviceName := "test_acc"
	resourceName := "ibm_database." + serviceName

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMDatabaseInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMDatabaseInstanceMongodbImport(databaseResourceGroup, serviceName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(resourceName, &databaseInstanceOne),
					resource.TestCheckResourceAttr(resourceName, "name", serviceName),
					resource.TestCheckResourceAttr(resourceName, "service", "databases-for-mongodb"),
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

func testAccCheckIBMDatabaseInstanceMongodbBasic(databaseResourceGroup string, name string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "test_acc" {
		name = "%[1]s"
	}

	resource "ibm_database" "%[2]s" {
		resource_group_id            = data.ibm_resource_group.test_acc.id
		name                         = "%[2]s"
		service                      = "databases-for-mongodb"
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
				allocation_mb = 10240
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
	}
				`, databaseResourceGroup, name, acc.Region())
}

func testAccCheckIBMDatabaseInstanceMongodbFullyspecified(databaseResourceGroup string, name string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "test_acc" {
		name = "%[1]s"
	}

	resource "ibm_database" "%[2]s" {
		resource_group_id            = data.ibm_resource_group.test_acc.id
		name                         = "%[2]s"
		service                      = "databases-for-mongodb"
		plan                         = "standard"
		location                     = "%[3]s"
		adminpassword                = "password12345678"
		service_endpoints            = "public"
		group {
			group_id = "member"
			memory {
				allocation_mb = 5120
			}
			host_flavor {
				id = "multitenant"
			}
			disk {
				allocation_mb = 10240
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
	}
				`, databaseResourceGroup, name, acc.Region())
}

func testAccCheckIBMDatabaseInstanceMongodbReduced(databaseResourceGroup string, name string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "test_acc" {
		name = "%[1]s"
	}

	resource "ibm_database" "%[2]s" {
		resource_group_id            = data.ibm_resource_group.test_acc.id
		name                         = "%[2]s"
		service                      = "databases-for-mongodb"
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
				allocation_mb = 10240
			}
		}
	}

				`, databaseResourceGroup, name, acc.Region())
}

func testAccCheckIBMDatabaseInstanceMongodbImport(databaseResourceGroup string, name string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "test_acc" {
		is_default = true
		# name = "%[1]s"
	}

	resource "ibm_database" "%[2]s" {
		resource_group_id = data.ibm_resource_group.test_acc.id
		name              = "%[2]s"
		service           = "databases-for-mongodb"
		plan              = "standard"
		location          = "%[3]s"
		service_endpoints            = "public"
	}

				`, databaseResourceGroup, name, acc.Region())
}
