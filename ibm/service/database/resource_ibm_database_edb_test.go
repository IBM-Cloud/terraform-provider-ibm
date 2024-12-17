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

func TestAccIBMEDBDatabaseInstanceBasic(t *testing.T) {
	t.Parallel()
	databaseResourceGroup := "default"
	var databaseInstanceOne string
	rnd := fmt.Sprintf("tf-edb-%d", acctest.RandIntRange(10, 100))
	testName := rnd
	name := "ibm_database." + testName

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMDatabaseInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMDatabaseInstanceEDBBasic(databaseResourceGroup, testName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(name, &databaseInstanceOne),
					resource.TestCheckResourceAttr(name, "name", testName),
					resource.TestCheckResourceAttr(name, "service", "databases-for-enterprisedb"),
					resource.TestCheckResourceAttr(name, "plan", "standard"),
					resource.TestCheckResourceAttr(name, "location", acc.Region()),
					resource.TestCheckResourceAttr(name, "adminuser", "admin"),
					resource.TestCheckResourceAttr(name, "groups.0.memory.0.allocation_mb", "49152"),
					resource.TestCheckResourceAttr(name, "groups.0.disk.0.allocation_mb", "61440"),
					resource.TestCheckResourceAttr(name, "service_endpoints", "public-and-private"),
					resource.TestCheckResourceAttr(name, "allowlist.#", "1"),
					resource.TestCheckResourceAttr(name, "users.#", "1"),
					resource.TestCheckResourceAttr(name, "tags.#", "1"),
				),
			},
			{
				Config: testAccCheckIBMDatabaseInstanceEDBFullyspecified(databaseResourceGroup, testName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(name, &databaseInstanceOne),
					resource.TestCheckResourceAttr(name, "name", testName),
					resource.TestCheckResourceAttr(name, "service", "databases-for-enterprisedb"),
					resource.TestCheckResourceAttr(name, "plan", "standard"),
					resource.TestCheckResourceAttr(name, "location", acc.Region()),
					resource.TestCheckResourceAttr(name, "groups.0.memory.0.allocation_mb", "98304"),
					resource.TestCheckResourceAttr(name, "groups.0.disk.0.allocation_mb", "92160"),
					resource.TestCheckResourceAttr(name, "service_endpoints", "public-and-private"),
					resource.TestCheckResourceAttr(name, "allowlist.#", "2"),
					resource.TestCheckResourceAttr(name, "users.#", "2"),
					resource.TestCheckResourceAttr(name, "tags.#", "1"),
				),
			},
			{
				Config: testAccCheckIBMDatabaseInstanceEDBReduced(databaseResourceGroup, testName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(name, &databaseInstanceOne),
					resource.TestCheckResourceAttr(name, "name", testName),
					resource.TestCheckResourceAttr(name, "service", "databases-for-enterprisedb"),
					resource.TestCheckResourceAttr(name, "plan", "standard"),
					resource.TestCheckResourceAttr(name, "location", acc.Region()),
					resource.TestCheckResourceAttr(name, "groups.0.memory.0.allocation_mb", "49152"),
					resource.TestCheckResourceAttr(name, "groups.0.disk.0.allocation_mb", "92160"),
					resource.TestCheckResourceAttr(name, "allowlist.#", "0"),
					resource.TestCheckResourceAttr(name, "users.#", "0"),
					resource.TestCheckResourceAttr(name, "tags.#", "1"),
				),
			},
			{
				ResourceName:      name,
				ImportState:       true,
				ImportStateVerify: false,
				ImportStateVerifyIgnore: []string{
					"wait_time_minutes", "adminpassword"},
			},
		},
	})
}

func TestAccIBMDatabaseInstanceEDBReadReplicaPromotion(t *testing.T) {
	t.Parallel()

	databaseResourceGroup := "default"

	var sourceInstanceCRN string
	var replicaInstanceCRN string

	serviceName := fmt.Sprintf("tf-edb-%d", acctest.RandIntRange(10, 100))
	readReplicaName := serviceName + "-replica"

	sourceResource := "ibm_database." + serviceName
	replicaReplicaResource := "ibm_database." + readReplicaName

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMDatabaseInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMDatabaseInstanceEDBMinimal(databaseResourceGroup, serviceName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(sourceResource, &sourceInstanceCRN),
					resource.TestCheckResourceAttr(sourceResource, "name", serviceName),
					resource.TestCheckResourceAttr(sourceResource, "service", "databases-for-enterprisedb"),
					resource.TestCheckResourceAttr(sourceResource, "plan", "standard"),
					resource.TestCheckResourceAttr(sourceResource, "location", acc.Region()),
				),
			},
			{
				Config: acc.ConfigCompose(
					testAccCheckIBMDatabaseInstanceEDBMinimal(databaseResourceGroup, serviceName),
					testAccCheckIBMDatabaseInstanceEDBMinimal_ReadReplica(databaseResourceGroup, serviceName)),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(replicaReplicaResource, &replicaInstanceCRN),
					resource.TestCheckResourceAttr(replicaReplicaResource, "name", readReplicaName),
					resource.TestCheckResourceAttr(replicaReplicaResource, "service", "databases-for-enterprisedb"),
					resource.TestCheckResourceAttr(replicaReplicaResource, "plan", "standard"),
					resource.TestCheckResourceAttr(replicaReplicaResource, "location", acc.Region()),
				),
			},
			{
				Config: acc.ConfigCompose(
					testAccCheckIBMDatabaseInstanceEDBMinimal(databaseResourceGroup, serviceName),
					testAccCheckIBMDatabaseInstanceEDBReadReplicaPromote(databaseResourceGroup, readReplicaName)),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(sourceResource, &sourceInstanceCRN),
					testAccCheckIBMDatabaseInstanceExists(replicaReplicaResource, &replicaInstanceCRN),
					resource.TestCheckResourceAttr(replicaReplicaResource, "name", readReplicaName),
					resource.TestCheckResourceAttr(replicaReplicaResource, "service", "databases-for-enterprisedb"),
					resource.TestCheckResourceAttr(replicaReplicaResource, "plan", "standard"),
					resource.TestCheckResourceAttr(replicaReplicaResource, "location", acc.Region()),
					resource.TestCheckResourceAttr(replicaReplicaResource, "remote_leader_id", ""),
					resource.TestCheckResourceAttr(replicaReplicaResource, "skip_initial_backup", "true"),
				),
			},
		},
	})
}

func testAccCheckIBMDatabaseInstanceEDBMinimal(databaseResourceGroup string, name string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "test_acc" {
		is_default = true
		# name = "%[1]s"
	}

	resource "ibm_database" "%[2]s" {
		resource_group_id = data.ibm_resource_group.test_acc.id
		name              = "%[2]s"
		service           = "databases-for-enterprisedb"
		plan              = "standard"
		location          = "%[3]s"
		service_endpoints = "public-and-private"

		group {
			group_id = "member"
			host_flavor {
				id = "b3c.4x16.encrypted"
			}
			disk {
			  allocation_mb = 20480
			}
		}
	}
	`, databaseResourceGroup, name, acc.Region())
}

func testAccCheckIBMDatabaseInstanceEDBBasic(databaseResourceGroup string, name string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "test_acc" {
		name = "%[1]s"
	}

	resource "ibm_database" "%[2]s" {
		resource_group_id            = data.ibm_resource_group.test_acc.id
		name                         = "%[2]s"
		service                      = "databases-for-enterprisedb"
		plan                         = "standard"
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
		service_endpoints            = "public-and-private"
		tags                         = ["one:two"]
		users {
		  name     = "user123"
		  password = "secure-Password12345"
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

func testAccCheckIBMDatabaseInstanceEDBFullyspecified(databaseResourceGroup string, name string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "test_acc" {
		name = "%[1]s"
	}

	resource "ibm_database" "%[2]s" {
		resource_group_id            = data.ibm_resource_group.test_acc.id
		name                         = "%[2]s"
		service                      = "databases-for-enterprisedb"
		plan                         = "standard"
		location                     = "%[3]s"
		adminpassword                = "secure-Password12345"
		group {
			group_id = "member"
			host_flavor {
				id = "b3c.8x32.encrypted"
			}
			disk {
			  allocation_mb = 30720
			}
		}
		service_endpoints            = "public-and-private"
		tags                         = ["one:two"]
		users {
		  name     = "user123"
		  password = "secure-Password12345"
		}
		users {
		  name     = "user124"
		  password = "secure-Password12345"
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

func testAccCheckIBMDatabaseInstanceEDBReduced(databaseResourceGroup string, name string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "test_acc" {
		name = "%[1]s"
	  }

	  resource "ibm_database" "%[2]s" {
		resource_group_id            = data.ibm_resource_group.test_acc.id
		name                         = "%[2]s"
		service                      = "databases-for-enterprisedb"
		plan                         = "standard"
		location                     = "%[3]s"
		adminpassword                = "secure-Password12345"
		group {
			group_id = "member"
			host_flavor {
				id = "b3c.4x16.encrypted"
			}
			disk {
			  allocation_mb = 30720
			}
		}
		service_endpoints            = "public"
		tags                         = ["one:two"]
		timeouts {
			create = "120m"
			update = "120m"
			delete = "15m"
		}
	  }
				`, databaseResourceGroup, name, acc.Region())
}

func testAccCheckIBMDatabaseInstanceEDBMinimal_ReadReplica(databaseResourceGroup string, name string) string {
	return fmt.Sprintf(`
	resource "ibm_database" "%[2]s-replica" {
		resource_group_id = data.ibm_resource_group.test_acc.id
		name                = "%[2]s-replica"
		service             = "databases-for-enterprisedb"
		plan                = "standard"
		location            = "%[3]s"
		service_endpoints   = "public-and-private"
		remote_leader_id    = ibm_database.%[2]s.id

		group {
			group_id = "member"
			host_flavor {
				id = "b3c.4x16.encrypted"
			}
			disk {
			  allocation_mb = 20480
			}
		}
	}
	`, databaseResourceGroup, name, acc.Region())
}

func testAccCheckIBMDatabaseInstanceEDBReadReplicaPromote(databaseResourceGroup string, readReplicaName string) string {
	return fmt.Sprintf(`
	resource "ibm_database" "%[2]s" {
		resource_group_id = data.ibm_resource_group.test_acc.id
		name                = "%[2]s"
		service             = "databases-for-enterprisedb"
		plan                = "standard"
		location            = "%[3]s"
		service_endpoints   = "public-and-private"
		remote_leader_id    = ""
		skip_initial_backup = true

		group {
			group_id = "member"
			host_flavor {
				id = "b3c.4x16.encrypted"
			}
			disk {
			  allocation_mb = 20480
			}
		}
	}
	`, databaseResourceGroup, readReplicaName, acc.Region())
}
