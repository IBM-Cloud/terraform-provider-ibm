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

func TestAccIBMMongoDBEnterpriseDatabaseInstanceBasic(t *testing.T) {
	t.Parallel()
	databaseResourceGroup := "default"
	var databaseInstanceOne string
	rnd := fmt.Sprintf("tf-mongoEnterprise-%d", acctest.RandIntRange(10, 100))
	testName := rnd
	name := "ibm_database." + testName

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMDatabaseInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMDatabaseInstanceMongoDBEnterpriseBasic(databaseResourceGroup, testName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(name, &databaseInstanceOne),
					resource.TestCheckResourceAttr(name, "name", testName),
					resource.TestCheckResourceAttr(name, "service", "databases-for-mongodb"),
					resource.TestCheckResourceAttr(name, "plan", "enterprise"),
					resource.TestCheckResourceAttr(name, "location", acc.Region()),
					resource.TestCheckResourceAttr(name, "adminuser", "admin"),
					resource.TestCheckResourceAttr(name, "service_endpoints", "public"),
					resource.TestCheckResourceAttr(name, "groups.0.memory.0.allocation_mb", "49152"),
					resource.TestCheckResourceAttr(name, "groups.0.disk.0.allocation_mb", "61440"),
					resource.TestCheckResourceAttr(name, "allowlist.#", "1"),
					resource.TestCheckResourceAttr(name, "users.#", "1"),
					resource.TestCheckResourceAttr(name, "tags.#", "1"),
				),
			},
			{
				Config: testAccCheckIBMDatabaseInstanceMongoDBEnterpriseFullyspecified(databaseResourceGroup, testName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(name, &databaseInstanceOne),
					resource.TestCheckResourceAttr(name, "name", testName),
					resource.TestCheckResourceAttr(name, "service", "databases-for-mongodb"),
					resource.TestCheckResourceAttr(name, "plan", "enterprise"),
					resource.TestCheckResourceAttr(name, "location", acc.Region()),
					resource.TestCheckResourceAttr(name, "service_endpoints", "public"),
					resource.TestCheckResourceAttr(name, "groups.0.memory.0.allocation_mb", "98304"),
					resource.TestCheckResourceAttr(name, "groups.0.disk.0.allocation_mb", "122880"),
					resource.TestCheckResourceAttr(name, "allowlist.#", "2"),
					resource.TestCheckResourceAttr(name, "users.#", "2"),
					resource.TestCheckResourceAttr(name, "users.1.type", "ops_manager"),
					resource.TestCheckResourceAttr(name, "tags.#", "1"),
				),
			},
			{
				Config: testAccCheckIBMDatabaseInstanceMongoDBEnterpriseReduced(databaseResourceGroup, testName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(name, &databaseInstanceOne),
					resource.TestCheckResourceAttr(name, "name", testName),
					resource.TestCheckResourceAttr(name, "service", "databases-for-mongodb"),
					resource.TestCheckResourceAttr(name, "plan", "enterprise"),
					resource.TestCheckResourceAttr(name, "location", acc.Region()),
					resource.TestCheckResourceAttr(name, "allowlist.#", "0"),
					resource.TestCheckResourceAttr(name, "groups.0.memory.0.allocation_mb", "49152"),
					resource.TestCheckResourceAttr(name, "groups.0.disk.0.allocation_mb", "122880"),
					resource.TestCheckResourceAttr(name, "users.#", "0"),
					resource.TestCheckResourceAttr(name, "tags.#", "1"),
				),
			},
			{
				ResourceName:      name,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"wait_time_minutes", "adminpassword", "group"},
			},
		},
	})
}

func TestAccIBMMongoDBEnterpriseDatabaseInstanceGroupBasic(t *testing.T) {
	t.Parallel()
	databaseResourceGroup := "default"
	var databaseInstanceOne string
	rnd := fmt.Sprintf("tf-mongoEnterprise-%d", acctest.RandIntRange(10, 100))
	testName := rnd
	name := "ibm_database." + testName

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMDatabaseInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMDatabaseInstanceMongoDBEnterpriseGroupBasic(databaseResourceGroup, testName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(name, &databaseInstanceOne),
					resource.TestCheckResourceAttr(name, "name", testName),
					resource.TestCheckResourceAttr(name, "service", "databases-for-mongodb"),
					resource.TestCheckResourceAttr(name, "plan", "enterprise"),
					resource.TestCheckResourceAttr(name, "location", acc.Region()),
					resource.TestCheckResourceAttr(name, "adminuser", "admin"),
					resource.TestCheckResourceAttr(name, "service_endpoints", "public"),
					resource.TestCheckResourceAttr(name, "groups.0.memory.0.allocation_mb", "49152"),
					resource.TestCheckResourceAttr(name, "groups.0.disk.0.allocation_mb", "61440"),
					resource.TestCheckResourceAttr(name, "groups.0.count", "3"),
					resource.TestCheckResourceAttr(name, "groups.1.count", "1"),
					resource.TestCheckResourceAttr(name, "groups.2.count", "1"),
					resource.TestCheckResourceAttr(name, "tags.#", "1"),
				),
			},
		},
	})
}

func TestAccIBMMongoDBEnterpriseDatabaseInstancePITR(t *testing.T) {
	t.Parallel()
	databaseResourceGroup := "Default"
	var databaseInstanceOne string
	var databaseInstanceTwo string
	serviceName := fmt.Sprintf("tf-mongodbee-%d", acctest.RandIntRange(10, 100))
	pitrServiceName := serviceName + "-pitr"
	resourceName := "ibm_database." + serviceName
	pitrResourceName := "ibm_database." + pitrServiceName

	resource.Test(t, resource.TestCase{
		PreCheck: func() { acc.TestAccPreCheck(t) },
		ExternalProviders: map[string]resource.ExternalProvider{
			"time": {
				Source:            "hashicorp/time",
				VersionConstraint: ">=0.9.1",
			},
		},
		ProviderFactories: acc.TestAccProviderFactories(),
		CheckDestroy:      testAccCheckIBMDatabaseInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMDatabaseInstanceMongoDBEnterpriseMinimal(databaseResourceGroup, serviceName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(resourceName, &databaseInstanceOne),
					resource.TestCheckResourceAttr(resourceName, "name", serviceName),
					resource.TestCheckResourceAttr(resourceName, "service", "databases-for-mongodb"),
					resource.TestCheckResourceAttr(resourceName, "plan", "enterprise"),
					resource.TestCheckResourceAttr(resourceName, "location", acc.Region()),

					resource.TestCheckResourceAttr(resourceName, "groups.0.count", "3"),
					resource.TestCheckResourceAttr(resourceName, "groups.1.count", "0"),
					resource.TestCheckResourceAttr(resourceName, "groups.2.count", "0"),
				),
			},
			{
				Config: acc.ConfigCompose(acc.ConfigAlternateRegionProvider(),
					testAccCheckIBMDatabaseInstanceMongoDBEnterpriseMinimal(databaseResourceGroup, serviceName),
					testAccCheckIBMDatabaseInstanceMongoDBEnterpriseMinimal_PITR(databaseResourceGroup, serviceName)),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(resourceName, &databaseInstanceOne),
					testAccCheckIBMDatabaseInstanceExists(pitrResourceName, &databaseInstanceTwo),
					resource.TestCheckResourceAttr(pitrResourceName, "name", pitrServiceName),
					resource.TestCheckResourceAttr(pitrResourceName, "service", "databases-for-mongodb"),
					resource.TestCheckResourceAttr(pitrResourceName, "plan", "enterprise"),
					resource.TestCheckResourceAttr(pitrResourceName, "location", acc.RegionAlternate()),
					resource.TestCheckResourceAttr(pitrResourceName, "adminuser", "admin"),
					resource.TestCheckResourceAttr(pitrResourceName, "groups.0.count", "3"),
					resource.TestCheckResourceAttr(pitrResourceName, "groups.1.count", "0"),
					resource.TestCheckResourceAttr(pitrResourceName, "groups.2.count", "0"),
				),
			},
		},
	})
}

func testAccCheckIBMDatabaseInstanceMongoDBEnterpriseBasic(databaseResourceGroup string, name string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "test_acc" {
		name = "%[1]s"
	}

	resource "ibm_database" "%[2]s" {
		resource_group_id            = data.ibm_resource_group.test_acc.id
		name                         = "%[2]s"
		service                      = "databases-for-mongodb"
		plan                         = "enterprise"
		location                     = "%[3]s"
		adminpassword                = "password12345678"
		tags                         = ["one:two"]
		service_endpoints            = "public"
		group {
			group_id = "member"
			host_flavor {
				id = "b3c.4x16.encrypted"
			}
			disk {
				allocation_mb = 20480
			}
		}
		users {
		  name     = "user123"
		  password = "password12345678"
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

func testAccCheckIBMDatabaseInstanceMongoDBEnterpriseFullyspecified(databaseResourceGroup string, name string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "test_acc" {
		name = "%[1]s"
	}

	resource "ibm_database" "%[2]s" {
		resource_group_id            = data.ibm_resource_group.test_acc.id
		name                         = "%[2]s"
		service                      = "databases-for-mongodb"
		plan                         = "enterprise"
		location                     = "%[3]s"
		adminpassword                = "password12345678"
		tags                         = ["one:two"]
		service_endpoints            = "public"
		group {
			group_id = "member"
			host_flavor {
				id = "b3c.8x32.encrypted"
			}
			disk {
				allocation_mb = 40960
			}
		}
		users {
		  name     = "user123"
		  password = "password12345678"
		  type     = "database"
		}
		users {
		  name     = "user124"
		  password = "password12345678$password"
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

func testAccCheckIBMDatabaseInstanceMongoDBEnterpriseReduced(databaseResourceGroup string, name string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "test_acc" {
		name = "%[1]s"
	  }

	  resource "ibm_database" "%[2]s" {
		resource_group_id            = data.ibm_resource_group.test_acc.id
		name                         = "%[2]s"
		service                      = "databases-for-mongodb"
		plan                         = "enterprise"
		location                     = "%[3]s"
		adminpassword                = "password12345678"
		service_endpoints            = "public"
		tags                         = ["one:two"]
		group {
			group_id = "member"
			host_flavor {
				id = "b3c.4x16.encrypted"
			}
			disk {
				allocation_mb = 40960
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

func testAccCheckIBMDatabaseInstanceMongoDBEnterpriseGroupBasic(databaseResourceGroup string, name string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "test_acc" {
		name = "%[1]s"
	}

	resource "ibm_database" "%[2]s" {
		resource_group_id            = data.ibm_resource_group.test_acc.id
		name                         = "%[2]s"
		service                      = "databases-for-mongodb"
		plan                         = "enterprise"
		location                     = "%[3]s"
		adminpassword                = "password12345678"
		tags                         = ["one:two"]
		service_endpoints            = "public"

		group {
			group_id = "member"

			host_flavor {
				id = "b3c.4x16.encrypted"
			}
			disk {
				allocation_mb = 20480
			}
		}

		group {
			group_id = "bi_connector"

			members {
				allocation_count = 1
			}
		}

		group {
			group_id = "analytics"

			members {
				allocation_count = 1
			}
		}

		timeouts {
			create = "4h"
			update = "4h"
			delete = "15m"
		}
	}
				`, databaseResourceGroup, name, acc.Region())
}

func testAccCheckIBMDatabaseInstanceMongoDBEnterpriseMinimal(databaseResourceGroup string, name string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "test_acc" {
		is_default = true
	}

	resource "ibm_database" "%[2]s" {
		resource_group_id            = data.ibm_resource_group.test_acc.id
		name                         = "%[2]s"
		service                      = "databases-for-mongodb"
		plan                         = "enterprise"
		location                     = "%[3]s"
		service_endpoints            = "public"

		group {
			group_id = "member"

			host_flavor {
				id = "b3c.4x16.encrypted"
			}
		}
		timeouts {
			create = "4h"
			update = "4h"
			delete = "15m"
		}
	}
				`, databaseResourceGroup, name, acc.Region())
}

func testAccCheckIBMDatabaseInstanceMongoDBEnterpriseMinimal_PITR(databaseResourceGroup string, name string) string {
	return fmt.Sprintf(`
	resource "time_sleep" "wait_time" {
		create_duration = "1h"
		depends_on      = [ibm_database.%[2]s]
	}

	resource "ibm_database" "%[2]s-pitr" {
		provider                              = "%[1]s"
		depends_on                            = [time_sleep.wait_time, ibm_database.%[2]s]

		resource_group_id                     = data.ibm_resource_group.test_acc.id
		name                                  = "%[2]s-pitr"
		service                               = "databases-for-mongodb"
		plan                                  = "enterprise"
		location                              = "%[3]s"
		point_in_time_recovery_deployment_id  = ibm_database.%[2]s.id
		point_in_time_recovery_time           = ""
    	offline_restore                       = true
		service_endpoints            = "public"

		group {
			group_id = "member"

			host_flavor {
				id = "b3c.4x16.encrypted"
			}
		}

		timeouts {
			create = "4h"
			update = "4h"
			delete = "15m"
		}
	}

				`, acc.ProviderNameAlternate, name, acc.RegionAlternate())
}
