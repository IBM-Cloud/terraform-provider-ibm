// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package database_test

import (
	"fmt"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMMysqlDatabaseInstanceBasic(t *testing.T) {
	t.Parallel()
	databaseResourceGroup := "default"
	var databaseInstanceOne string
	rnd := fmt.Sprintf("tf-mysql-%s", acctest.RandString(6))
	testName := rnd
	name := "ibm_database." + testName

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMDatabaseInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMDatabaseInstanceMysqlBasic(databaseResourceGroup, testName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(name, &databaseInstanceOne),
					resource.TestCheckResourceAttr(name, "name", testName),
					resource.TestCheckResourceAttr(name, "service", "databases-for-mysql"),
					resource.TestCheckResourceAttr(name, "plan", "standard"),
					resource.TestCheckResourceAttr(name, "location", acc.Region()),
					resource.TestCheckResourceAttr(name, "adminuser", "admin"),
					resource.TestCheckResourceAttr(name, "groups.0.memory.0.allocation_mb", "12288"),
					resource.TestCheckResourceAttr(name, "groups.0.disk.0.allocation_mb", "61440"),
					resource.TestCheckResourceAttr(name, "service_endpoints", "public"),
					resource.TestCheckResourceAttr(name, "allowlist.#", "1"),
					resource.TestCheckResourceAttr(name, "users.#", "1"),
					resource.TestCheckResourceAttr(name, "tags.#", "1"),
				),
			},
			{
				Config: testAccCheckIBMDatabaseInstanceMysqlFullyspecified(databaseResourceGroup, testName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(name, &databaseInstanceOne),
					resource.TestCheckResourceAttr(name, "name", testName),
					resource.TestCheckResourceAttr(name, "service", "databases-for-mysql"),
					resource.TestCheckResourceAttr(name, "plan", "standard"),
					resource.TestCheckResourceAttr(name, "location", acc.Region()),
					resource.TestCheckResourceAttr(name, "groups.0.memory.0.allocation_mb", "15360"),
					resource.TestCheckResourceAttr(name, "groups.0.disk.0.allocation_mb", "92160"),
					resource.TestCheckResourceAttr(name, "service_endpoints", "public-and-private"),
					resource.TestCheckResourceAttr(name, "allowlist.#", "2"),
					resource.TestCheckResourceAttr(name, "users.#", "2"),
					resource.TestCheckResourceAttr(name, "tags.#", "1"),
				),
			},
		},
	})
}

func TestAccIBMDatabaseInstanceMySQLReadReplicaPromotion(t *testing.T) {
	t.Parallel()

	databaseResourceGroup := "default"

	var sourceInstanceCRN string
	var replicaInstanceCRN string

	serviceName := fmt.Sprintf("tf-mysql-%d", acctest.RandIntRange(10, 100))
	readReplicaName := serviceName + "-replica"

	sourceResource := "ibm_database." + serviceName
	replicaReplicaResource := "ibm_database." + readReplicaName

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMDatabaseInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMDatabaseInstanceMySQLMinimal(databaseResourceGroup, serviceName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(sourceResource, &sourceInstanceCRN),
					resource.TestCheckResourceAttr(sourceResource, "name", serviceName),
					resource.TestCheckResourceAttr(sourceResource, "service", "databases-for-mysql"),
					resource.TestCheckResourceAttr(sourceResource, "plan", "standard"),
					resource.TestCheckResourceAttr(sourceResource, "location", acc.Region()),
				),
			},
			{
				Config: acc.ConfigCompose(
					testAccCheckIBMDatabaseInstanceMySQLMinimal(databaseResourceGroup, serviceName),
					testAccCheckIBMDatabaseInstanceMySQLMinimal_ReadReplica(databaseResourceGroup, serviceName)),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(replicaReplicaResource, &replicaInstanceCRN),
					resource.TestCheckResourceAttr(replicaReplicaResource, "name", readReplicaName),
					resource.TestCheckResourceAttr(replicaReplicaResource, "service", "databases-for-mysql"),
					resource.TestCheckResourceAttr(replicaReplicaResource, "plan", "standard"),
					resource.TestCheckResourceAttr(replicaReplicaResource, "location", acc.Region()),
				),
			},
			{
				Config: acc.ConfigCompose(
					testAccCheckIBMDatabaseInstanceMySQLMinimal(databaseResourceGroup, serviceName),
					testAccCheckIBMDatabaseInstanceMySQLReadReplicaPromote(databaseResourceGroup, readReplicaName)),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(replicaReplicaResource, &replicaInstanceCRN),
					resource.TestCheckResourceAttr(replicaReplicaResource, "name", readReplicaName),
					resource.TestCheckResourceAttr(replicaReplicaResource, "service", "databases-for-mysql"),
					resource.TestCheckResourceAttr(replicaReplicaResource, "plan", "standard"),
					resource.TestCheckResourceAttr(replicaReplicaResource, "location", acc.Region()),
					resource.TestCheckResourceAttr(replicaReplicaResource, "remote_leader_id", ""),
					resource.TestCheckResourceAttr(replicaReplicaResource, "skip_initial_backup", "true"),
				),
			},
		},
	})
}

func testAccCheckIBMDatabaseInstanceMySQLMinimal(databaseResourceGroup string, name string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "test_acc" {
		is_default = true
		# name = "%[1]s"
	}

	resource "ibm_database" "%[2]s" {
		resource_group_id = data.ibm_resource_group.test_acc.id
		name              = "%[2]s"
		service           = "databases-for-mysql"
		plan              = "standard"
		location          = "%[3]s"
		service_endpoints = "public-and-private"
	}
	`, databaseResourceGroup, name, acc.Region())
}

func testAccCheckIBMDatabaseInstanceMySQLMinimal_ReadReplica(databaseResourceGroup string, name string) string {
	return fmt.Sprintf(`
	resource "ibm_database" "%[2]s-replica" {
		resource_group_id = data.ibm_resource_group.test_acc.id
		name                = "%[2]s-replica"
		service             = "databases-for-mysql"
		plan                = "standard"
		location            = "%[3]s"
		service_endpoints   = "public-and-private"
		remote_leader_id    = ibm_database.%[2]s.id
	}
	`, databaseResourceGroup, name, acc.Region())
}

func testAccCheckIBMDatabaseInstanceMySQLReadReplicaPromote(databaseResourceGroup string, readReplicaName string) string {
	return fmt.Sprintf(`
	resource "ibm_database" "%[2]s" {
		resource_group_id = data.ibm_resource_group.test_acc.id
		name                = "%[2]s"
		service             = "databases-for-mysql"
		plan                = "standard"
		location            = "%[3]s"
		service_endpoints   = "public-and-private"
		remote_leader_id    = ""
		skip_initial_backup = true
	}
	`, databaseResourceGroup, readReplicaName, acc.Region())
}

func testAccCheckIBMDatabaseInstanceMysqlBasic(databaseResourceGroup string, name string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "test_acc" {
		name = "%[1]s"
	}

	resource "ibm_database" "%[2]s" {
		resource_group_id            = data.ibm_resource_group.test_acc.id
		name                         = "%[2]s"
		service                      = "databases-for-mysql"
		plan                         = "standard"
		location                     = "%[3]s"
		adminpassword                = "secure-Password12345"
		group {
			group_id = "member"
			memory {
				allocation_mb = 4096
			}
			host_flavor {
				id = "multitenant"
			}
			disk {
				allocation_mb = 20480
			}
		}
		service_endpoints            = "public"
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

func testAccCheckIBMDatabaseInstanceMysqlFullyspecified(databaseResourceGroup string, name string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "test_acc" {
		name = "%[1]s"
	}

	resource "ibm_database" "%[2]s" {
		resource_group_id            = data.ibm_resource_group.test_acc.id
		name                         = "%[2]s"
		service                      = "databases-for-mysql"
		plan                         = "standard"
		location                     = "%[3]s"
		adminpassword                = "secure-Password12345"
		group {
			group_id = "member"
			memory {
				allocation_mb = 5120
			}
			disk {
				allocation_mb = 30720
			}
			cpu {
				allocation_count = 4
			}
			host_flavor {
				id = "multitenant"
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
		configuration = <<CONFIGURATION
		{
			"mysql_max_binlog_age_sec": 2000,
			"innodb_buffer_pool_size_percentage": 60
		}
		CONFIGURATION
		timeouts {
			create = "120m"
			update = "120m"
			delete = "15m"
		}
	}
				`, databaseResourceGroup, name, acc.Region())
}
