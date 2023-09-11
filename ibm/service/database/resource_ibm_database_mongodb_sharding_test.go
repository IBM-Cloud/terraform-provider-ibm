// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
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
					resource.TestCheckResourceAttr(name, "location", acc.IcdDbRegion),
					resource.TestCheckResourceAttr(name, "adminuser", "admin"),
					resource.TestCheckResourceAttr(name, "members_memory_allocation_mb", "86016"),
					resource.TestCheckResourceAttr(name, "members_disk_allocation_mb", "122880"),
					resource.TestCheckResourceAttr(name, "allowlist.#", "1"),
					resource.TestCheckResourceAttr(name, "users.#", "1"),
					resource.TestCheckResourceAttr(name, "connectionstrings.#", "2"),
					resource.TestCheckResourceAttr(name, "connectionstrings.1.name", "admin"),
					resource.TestMatchResourceAttr(name, "connectionstrings.1.certname", regexp.MustCompile("[-a-z0-9]*")),
					resource.TestMatchResourceAttr(name, "connectionstrings.1.certbase64", regexp.MustCompile("^(?:[A-Za-z0-9+/]{4})*(?:[A-Za-z0-9+/]{2}==|[A-Za-z0-9+/]{3}=)?$")),
				),
			},
			{
				Config: testAccCheckIBMDatabaseInstanceMongoDBShardingFullyspecified(databaseResourceGroup, testName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(name, &databaseInstanceOne),
					resource.TestCheckResourceAttr(name, "name", testName),
					resource.TestCheckResourceAttr(name, "service", "databases-for-mongodb"),
					resource.TestCheckResourceAttr(name, "plan", "enterprise-sharding"),
					resource.TestCheckResourceAttr(name, "location", acc.IcdDbRegion),
					resource.TestCheckResourceAttr(name, "members_memory_allocation_mb", "86016"),
					resource.TestCheckResourceAttr(name, "members_disk_allocation_mb", "122880"),
					resource.TestCheckResourceAttr(name, "allowlist.#", "2"),
					resource.TestCheckResourceAttr(name, "users.#", "2"),
					resource.TestCheckResourceAttr(name, "connectionstrings.#", "3"),
					resource.TestCheckResourceAttr(name, "connectionstrings.2.name", "admin"),
					resource.TestCheckResourceAttr(name, "connectionstrings.0.scheme", "mongodb"),
				),
			},
			{
				Config: testAccCheckIBMDatabaseInstanceMongoDBShardingReduced(databaseResourceGroup, testName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(name, &databaseInstanceOne),
					resource.TestCheckResourceAttr(name, "name", testName),
					resource.TestCheckResourceAttr(name, "service", "databases-for-mongodb"),
					resource.TestCheckResourceAttr(name, "plan", "enterprise-sharding"),
					resource.TestCheckResourceAttr(name, "location", acc.IcdDbRegion),
					resource.TestCheckResourceAttr(name, "members_memory_allocation_mb", "86016"),
					resource.TestCheckResourceAttr(name, "members_disk_allocation_mb", "122880"),
					resource.TestCheckResourceAttr(name, "allowlist.#", "0"),
					resource.TestCheckResourceAttr(name, "users.#", "0"),
					resource.TestCheckResourceAttr(name, "connectionstrings.#", "1"),
				),
			},
		},
	})
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
		adminpassword                = "password12"
		members_disk_allocation_mb   = 122880
    	members_memory_allocation_mb = 86016
		users {
		  name     = "user123"
		  password = "password12"
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
				`, databaseResourceGroup, name, acc.IcdDbRegion)
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
		adminpassword                = "password12"
		members_memory_allocation_mb = 86016
		members_disk_allocation_mb   = 122880
		members_cpu_allocation_count = 36
		users {
		  name     = "user123"
		  password = "password12"
		  type     = "database"
		}
		users {
		  name     = "user124"
		  password = "password12password"
		  type     = "database"
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
				`, databaseResourceGroup, name, acc.IcdDbRegion)
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
		adminpassword                = "password12"
		members_disk_allocation_mb   = 122880
    	members_memory_allocation_mb = 86016
		service_endpoints            = "public"
		timeouts {
			create = "480m"
			update = "480m"
			delete = "15m"
		}
	  }
				`, databaseResourceGroup, name, acc.IcdDbRegion)
}
