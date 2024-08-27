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

func TestAccIBMDatabaseInstance_Etcd_Basic(t *testing.T) {
	t.Parallel()
	databaseResourceGroup := "default"
	var databaseInstanceOne string
	rnd := fmt.Sprintf("tf-Etcd-%d", acctest.RandIntRange(10, 100))
	testName := rnd
	name := "ibm_database." + testName

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMDatabaseInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMDatabaseInstanceEtcdBasic(databaseResourceGroup, testName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(name, &databaseInstanceOne),
					resource.TestCheckResourceAttr(name, "name", testName),
					resource.TestCheckResourceAttr(name, "service", "databases-for-etcd"),
					resource.TestCheckResourceAttr(name, "plan", "standard"),
					resource.TestCheckResourceAttr(name, "location", acc.Region()),
					resource.TestCheckResourceAttr(name, "adminuser", "root"),
					resource.TestCheckResourceAttr(name, "groups.0.memory.0.allocation_mb", "12288"),
					resource.TestCheckResourceAttr(name, "groups.0.disk.0.allocation_mb", "184320"),
					resource.TestCheckResourceAttr(name, "allowlist.#", "1"),
					resource.TestCheckResourceAttr(name, "users.#", "1"),
				),
			},
			{
				Config: testAccCheckIBMDatabaseInstanceEtcdFullyspecified(databaseResourceGroup, testName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", testName),
					resource.TestCheckResourceAttr(name, "service", "databases-for-etcd"),
					resource.TestCheckResourceAttr(name, "plan", "standard"),
					resource.TestCheckResourceAttr(name, "location", acc.Region()),
					resource.TestCheckResourceAttr(name, "groups.0.memory.0.allocation_mb", "49152"),
					resource.TestCheckResourceAttr(name, "groups.0.disk.0.allocation_mb", "193536"),
					resource.TestCheckResourceAttr(name, "allowlist.#", "2"),
					resource.TestCheckResourceAttr(name, "users.#", "2"),
				),
			},
			{
				Config: testAccCheckIBMDatabaseInstanceEtcdReduced(databaseResourceGroup, testName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", testName),
					resource.TestCheckResourceAttr(name, "service", "databases-for-etcd"),
					resource.TestCheckResourceAttr(name, "plan", "standard"),
					resource.TestCheckResourceAttr(name, "location", acc.Region()),
					resource.TestCheckResourceAttr(name, "groups.0.memory.0.allocation_mb", "12288"),
					resource.TestCheckResourceAttr(name, "groups.0.disk.0.allocation_mb", "193536"),
					resource.TestCheckResourceAttr(name, "allowlist.#", "0"),
					resource.TestCheckResourceAttr(name, "users.#", "0"),
				),
			},
		},
	})
}

// TestAccIBMDatabaseInstance_CreateAfterManualDestroy not required as tested by resource_instance tests

func TestAccIBMDatabaseInstanceEtcdImport(t *testing.T) {
	t.Parallel()
	databaseResourceGroup := "default"
	var databaseInstanceOne string
	serviceName := fmt.Sprintf("tf-Etcd-%d", acctest.RandIntRange(10, 100))
	//serviceName := "test_acc"
	resourceName := "ibm_database." + serviceName

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMDatabaseInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMDatabaseInstanceEtcdImport(databaseResourceGroup, serviceName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(resourceName, &databaseInstanceOne),
					resource.TestCheckResourceAttr(resourceName, "name", serviceName),
					resource.TestCheckResourceAttr(resourceName, "service", "databases-for-etcd"),
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

func testAccCheckIBMDatabaseInstanceEtcdBasic(databaseResourceGroup string, name string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "test_acc" {
		is_default = true
		# name = "%[1]s"
	  }

	resource "ibm_database" "%[2]s" {
		resource_group_id            = data.ibm_resource_group.test_acc.id
		name                         = "%[2]s"
		service                      = "databases-for-etcd"
		plan                         = "standard"
		location                     = "%[3]s"
		adminpassword                = "password12345678"
		service_endpoints            = "public-and-private"
		group {
			group_id = "member"
			memory {
			  allocation_mb = 4096
			}
			host_flavor {
				id = "multitenant"
			}
			disk {
			  allocation_mb = 61440
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

func testAccCheckIBMDatabaseInstanceEtcdFullyspecified(databaseResourceGroup string, name string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "test_acc" {
		is_default = true
		# name = "%[1]s"
	}

	resource "ibm_database" "%[2]s" {
		resource_group_id            = data.ibm_resource_group.test_acc.id
		name                         = "%[2]s"
		service                      = "databases-for-etcd"
		plan                         = "standard"
		location                     = "%[3]s"
		adminpassword                = "password12345678"
		service_endpoints            = "public-and-private"
		group {
			group_id = "member"
			host_flavor {
				id = "b3c.4x16.encrypted"
			}
			disk {
			  allocation_mb = 64512
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

func testAccCheckIBMDatabaseInstanceEtcdReduced(databaseResourceGroup string, name string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "test_acc" {
		is_default = true
		# name = "%[1]s"
	}

	resource "ibm_database" "%[2]s" {
		resource_group_id            = data.ibm_resource_group.test_acc.id
		name                         = "%[2]s"
		service                      = "databases-for-etcd"
		plan                         = "standard"
		location                     = "%[3]s"
		adminpassword                = "password12345678"
		service_endpoints            = "public-and-private"
		group {
			group_id = "member"
			memory {
			  allocation_mb = 4096
			}
			host_flavor {
				id = "multitenant"
			}
			disk {
			  allocation_mb = 64512
			}
		}
	}
				`, databaseResourceGroup, name, acc.Region())
}

func testAccCheckIBMDatabaseInstanceEtcdImport(databaseResourceGroup string, name string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "test_acc" {
		is_default = true
		# name = "%[1]s"
	}

	resource "ibm_database" "%[2]s" {
		resource_group_id = data.ibm_resource_group.test_acc.id
		name              = "%[2]s"
		service           = "databases-for-etcd"
		plan              = "standard"
		location          = "%[3]s"
		service_endpoints            = "public-and-private"
	}
				`, databaseResourceGroup, name, acc.Region())
}
