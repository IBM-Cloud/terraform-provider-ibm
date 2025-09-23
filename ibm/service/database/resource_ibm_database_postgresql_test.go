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

func TestAccIBMDatabaseInstancePostgresBasic(t *testing.T) {
	t.Parallel()
	databaseResourceGroup := "default"
	var databaseInstanceOne string
	rnd := fmt.Sprintf("tf-Pgress-%d", acctest.RandIntRange(10, 100))
	testName := rnd
	name := "ibm_database." + testName

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMDatabaseInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMDatabaseInstancePostgresBasic(databaseResourceGroup, testName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(name, &databaseInstanceOne),
					resource.TestCheckResourceAttr(name, "name", testName),
					resource.TestCheckResourceAttr(name, "service", "databases-for-postgresql"),
					resource.TestCheckResourceAttr(name, "plan", "standard"),
					resource.TestCheckResourceAttr(name, "location", acc.Region()),
					resource.TestCheckResourceAttr(name, "adminuser", "admin"),
					resource.TestCheckResourceAttr(name, "groups.0.memory.0.allocation_mb", "8192"),
					resource.TestCheckResourceAttr(name, "groups.0.disk.0.allocation_mb", "20480"),
					resource.TestCheckResourceAttr(name, "groups.0.cpu.0.allocation_count", "0"),
					resource.TestCheckResourceAttr(name, "service_endpoints", "public"),
					resource.TestCheckResourceAttr(name, "allowlist.#", "1"),
					resource.TestCheckResourceAttr(name, "users.#", "1"),
					resource.TestCheckResourceAttr(name, "tags.#", "1"),
					resource.TestCheckResourceAttr(name, "logical_replication_slot.#", "1"),
				),
			},
			{
				Config: testAccCheckIBMDatabaseInstancePostgresFullyspecified(databaseResourceGroup, testName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(name, &databaseInstanceOne),
					resource.TestCheckResourceAttr(name, "name", testName),
					resource.TestCheckResourceAttr(name, "service", "databases-for-postgresql"),
					resource.TestCheckResourceAttr(name, "plan", "standard"),
					resource.TestCheckResourceAttr(name, "location", acc.Region()),
					resource.TestCheckResourceAttr(name, "groups.0.memory.0.allocation_mb", "16384"),
					resource.TestCheckResourceAttr(name, "groups.0.disk.0.allocation_mb", "28672"),
					resource.TestCheckResourceAttr(name, "service_endpoints", "public-and-private"),
					resource.TestCheckResourceAttr(name, "allowlist.#", "2"),
					resource.TestCheckResourceAttr(name, "users.#", "3"),
					resource.TestCheckResourceAttr(name, "tags.#", "1"),
					resource.TestCheckResourceAttr(name, "logical_replication_slot.#", "2"),
				),
			},
		},
	})
}

func TestAccIBMDatabaseInstancePostgresGroup(t *testing.T) {
	t.Parallel()
	databaseResourceGroup := "default"
	var databaseInstanceOne string
	rnd := fmt.Sprintf("tf-Pgress-%d", acctest.RandIntRange(10, 100))
	testName := rnd
	name := "ibm_database." + testName

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMDatabaseInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMDatabaseInstancePostgresGroupBasic(databaseResourceGroup, testName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(name, &databaseInstanceOne),
					resource.TestCheckResourceAttr(name, "name", testName),
					resource.TestCheckResourceAttr(name, "service", "databases-for-postgresql"),
					resource.TestCheckResourceAttr(name, "plan", "standard"),
					resource.TestCheckResourceAttr(name, "location", acc.Region()),
					resource.TestCheckResourceAttr(name, "adminuser", "admin"),
					resource.TestCheckResourceAttr(name, "groups.0.count", "2"),
					resource.TestCheckResourceAttr(name, "groups.0.memory.0.allocation_mb", "8192"),
					resource.TestCheckResourceAttr(name, "groups.0.disk.0.allocation_mb", "10240"),
					resource.TestCheckResourceAttr(name, "groups.0.cpu.0.allocation_count", "6"),
					resource.TestCheckResourceAttr(name, "service_endpoints", "public"),
					resource.TestCheckResourceAttr(name, "allowlist.#", "1"),
					resource.TestCheckResourceAttr(name, "users.#", "1"),
					resource.TestCheckResourceAttr(name, "tags.#", "1"),
				),
			},
			{
				Config: testAccCheckIBMDatabaseInstancePostgresGroupFullyspecified(databaseResourceGroup, testName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(name, &databaseInstanceOne),
					resource.TestCheckResourceAttr(name, "name", testName),
					resource.TestCheckResourceAttr(name, "service", "databases-for-postgresql"),
					resource.TestCheckResourceAttr(name, "plan", "standard"),
					resource.TestCheckResourceAttr(name, "location", acc.Region()),
					resource.TestCheckResourceAttr(name, "groups.0.count", "2"),
					resource.TestCheckResourceAttr(name, "groups.0.memory.0.allocation_mb", "16384"),
					resource.TestCheckResourceAttr(name, "groups.0.disk.0.allocation_mb", "14336"),
					resource.TestCheckResourceAttr(name, "groups.0.cpu.0.allocation_count", "6"),
					resource.TestCheckResourceAttr(name, "service_endpoints", "public-and-private"),
					resource.TestCheckResourceAttr(name, "allowlist.#", "2"),
					resource.TestCheckResourceAttr(name, "users.#", "2"),
					resource.TestCheckResourceAttr(name, "tags.#", "1"),
				),
			},
			{
				Config: testAccCheckIBMDatabaseInstancePostgresGroupReduced(databaseResourceGroup, testName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(name, &databaseInstanceOne),
					resource.TestCheckResourceAttr(name, "name", testName),
					resource.TestCheckResourceAttr(name, "service", "databases-for-postgresql"),
					resource.TestCheckResourceAttr(name, "plan", "standard"),
					resource.TestCheckResourceAttr(name, "location", acc.Region()),
					resource.TestCheckResourceAttr(name, "groups.0.count", "2"),
					resource.TestCheckResourceAttr(name, "groups.0.memory.0.allocation_mb", "8192"),
					resource.TestCheckResourceAttr(name, "groups.0.disk.0.allocation_mb", "14336"),
					resource.TestCheckResourceAttr(name, "groups.0.cpu.0.allocation_count", "6"),
					resource.TestCheckResourceAttr(name, "allowlist.#", "0"),
					resource.TestCheckResourceAttr(name, "users.#", "0"),
					resource.TestCheckResourceAttr(name, "tags.#", "1"),
				),
			},
			{
				Config: testAccCheckIBMDatabaseInstancePostgresGroupScaleOut(databaseResourceGroup, testName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(name, &databaseInstanceOne),
					resource.TestCheckResourceAttr(name, "name", testName),
					resource.TestCheckResourceAttr(name, "service", "databases-for-postgresql"),
					resource.TestCheckResourceAttr(name, "plan", "standard"),
					resource.TestCheckResourceAttr(name, "location", acc.Region()),
					resource.TestCheckResourceAttr(name, "groups.0.count", "3"),
					resource.TestCheckResourceAttr(name, "groups.0.memory.0.allocation_mb", "24576"),
					resource.TestCheckResourceAttr(name, "groups.0.disk.0.allocation_mb", "21504"),
					resource.TestCheckResourceAttr(name, "groups.0.cpu.0.allocation_count", "9"),
					resource.TestCheckResourceAttr(name, "allowlist.#", "0"),
					resource.TestCheckResourceAttr(name, "users.#", "0"),
					resource.TestCheckResourceAttr(name, "tags.#", "1"),
				),
			},
		},
	})
}

// TestAccIBMDatabaseInstance_CreateAfterManualDestroy not required as tested by resource_instance tests

func TestAccIBMDatabaseInstancePostgresImport(t *testing.T) {
	t.Parallel()

	databaseResourceGroup := "default"

	var databaseInstanceOne string

	serviceName := fmt.Sprintf("tf-Pgress-%d", acctest.RandIntRange(10, 100))
	resourceName := "ibm_database." + serviceName

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMDatabaseInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMDatabaseInstancePostgresImport(databaseResourceGroup, serviceName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(resourceName, &databaseInstanceOne),
					resource.TestCheckResourceAttr(resourceName, "name", serviceName),
					resource.TestCheckResourceAttr(resourceName, "service", "databases-for-postgresql"),
					resource.TestCheckResourceAttr(resourceName, "plan", "standard"),
					resource.TestCheckResourceAttr(resourceName, "location", acc.Region()),
				),
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"wait_time_minutes", "deletion_protection"},
			},
		},
	})
}

func TestAccIBMDatabaseInstancePostgresPITR(t *testing.T) {
	t.Parallel()

	databaseResourceGroup := "default"

	var databaseInstanceOne string
	var databaseInstanceTwo string

	serviceName := fmt.Sprintf("tf-Pgress-%d", acctest.RandIntRange(10, 100))
	pitrServiceName := serviceName + "-pitr"

	sourceResource := "ibm_database." + serviceName
	pitrResource := "ibm_database." + pitrServiceName

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMDatabaseInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMDatabaseInstancePostgresMinimal(databaseResourceGroup, serviceName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(sourceResource, &databaseInstanceOne),
					resource.TestCheckResourceAttr(sourceResource, "name", serviceName),
					resource.TestCheckResourceAttr(sourceResource, "service", "databases-for-postgresql"),
					resource.TestCheckResourceAttr(sourceResource, "plan", "standard"),
					resource.TestCheckResourceAttr(sourceResource, "location", acc.Region()),
				),
			},
			{
				Config: acc.ConfigCompose(
					testAccCheckIBMDatabaseInstancePostgresMinimal(databaseResourceGroup, serviceName),
					testAccCheckIBMDatabaseInstancePostgresMinimal_PITR(databaseResourceGroup, serviceName)),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(pitrResource, &databaseInstanceTwo),
					resource.TestCheckResourceAttr(pitrResource, "name", pitrServiceName),
					resource.TestCheckResourceAttr(pitrResource, "service", "databases-for-postgresql"),
					resource.TestCheckResourceAttr(pitrResource, "plan", "standard"),
					resource.TestCheckResourceAttr(pitrResource, "location", acc.Region()),
				),
			},
		},
	})
}

func TestAccIBMDatabaseInstancePostgresReadReplicaPromotion(t *testing.T) {
	t.Parallel()

	databaseResourceGroup := "default"

	var sourceInstanceCRN string
	var replicaInstanceCRN string

	serviceName := fmt.Sprintf("tf-Pgress-%d", acctest.RandIntRange(10, 100))
	readReplicaName := serviceName + "-replica"

	sourceResource := "ibm_database." + serviceName
	replicaReplicaResource := "ibm_database." + readReplicaName

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMDatabaseInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMDatabaseInstancePostgresMinimal(databaseResourceGroup, serviceName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(sourceResource, &sourceInstanceCRN),
					resource.TestCheckResourceAttr(sourceResource, "name", serviceName),
					resource.TestCheckResourceAttr(sourceResource, "service", "databases-for-postgresql"),
					resource.TestCheckResourceAttr(sourceResource, "plan", "standard"),
					resource.TestCheckResourceAttr(sourceResource, "location", acc.Region()),
				),
			},
			{
				Config: acc.ConfigCompose(
					testAccCheckIBMDatabaseInstancePostgresMinimal(databaseResourceGroup, serviceName),
					testAccCheckIBMDatabaseInstancePostgresMinimal_ReadReplica(databaseResourceGroup, serviceName)),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(replicaReplicaResource, &replicaInstanceCRN),
					resource.TestCheckResourceAttr(replicaReplicaResource, "name", readReplicaName),
					resource.TestCheckResourceAttr(replicaReplicaResource, "service", "databases-for-postgresql"),
					resource.TestCheckResourceAttr(replicaReplicaResource, "plan", "standard"),
					resource.TestCheckResourceAttr(replicaReplicaResource, "location", acc.Region()),
				),
			},
			{
				Config: acc.ConfigCompose(
					testAccCheckIBMDatabaseInstancePostgresMinimal(databaseResourceGroup, serviceName),
					testAccCheckIBMDatabaseInstancePostgresReadReplicaPromote(databaseResourceGroup, readReplicaName)),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(replicaReplicaResource, &replicaInstanceCRN),
					resource.TestCheckResourceAttr(replicaReplicaResource, "name", readReplicaName),
					resource.TestCheckResourceAttr(replicaReplicaResource, "service", "databases-for-postgresql"),
					resource.TestCheckResourceAttr(replicaReplicaResource, "plan", "standard"),
					resource.TestCheckResourceAttr(replicaReplicaResource, "location", acc.Region()),
					resource.TestCheckResourceAttr(replicaReplicaResource, "remote_leader_id", ""),
					resource.TestCheckResourceAttr(replicaReplicaResource, "skip_initial_backup", "true"),
				),
			},
		},
	})
}

func testAccCheckIBMDatabaseInstancePostgresBasic(databaseResourceGroup string, name string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "test_acc" {
		name = "%[1]s"
	}

	resource "ibm_database" "%[2]s" {
		resource_group_id            = data.ibm_resource_group.test_acc.id
		name                         = "%[2]s"
		service                      = "databases-for-postgresql"
		plan                         = "standard"
		location                     = "%[3]s"
		adminpassword                = "secure-Password12345"
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
		tags                         = ["one:two"]
		users {
			name     = "user123"
			password = "secure-Password12345"
		}
		allowlist {
			address     = "172.168.1.2/32"
			description = "desc1"
		}
		configuration                = <<CONFIGURATION
		{
		  "wal_level": "logical",
		  "max_replication_slots": 21,
		  "max_wal_senders": 21
		}
		CONFIGURATION
		logical_replication_slot {
			name = "wj123"
			database_name = "ibmclouddb"
			plugin_type = "wal2json"
		}
	}
	`, databaseResourceGroup, name, acc.Region())
}

func testAccCheckIBMDatabaseInstancePostgresFullyspecified(databaseResourceGroup string, name string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "test_acc" {
		name = "%[1]s"
	}

	resource "ibm_database" "%[2]s" {
		resource_group_id            = data.ibm_resource_group.test_acc.id
		name                         = "%[2]s"
		service                      = "databases-for-postgresql"
		plan                         = "standard"
		location                     = "%[3]s"
		adminpassword                = "secure-Password12345"
		group {
			group_id = "member"
			memory {
			  allocation_mb = 8192
			}
			disk {
			  allocation_mb = 14336
			}
			cpu {
			  allocation_count = 6
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
		users {
			name     = "repl"
			password = "repl123456Password"
		}
		configuration                   = <<CONFIGURATION
		{
		  "wal_level": "logical",
		  "max_replication_slots": 21,
		  "max_wal_senders": 21
		}
		CONFIGURATION
		allowlist {
			address     = "172.168.1.2/32"
			description = "desc1"
		}
		allowlist {
			address     = "172.168.1.1/32"
			description = "desc"
		}
		logical_replication_slot {
			name = "wj123"
			database_name = "ibmclouddb"
			plugin_type = "wal2json"
		}
		logical_replication_slot {
			name = "wj321"
			database_name = "ibmclouddb"
			plugin_type = "wal2json"
		}
	}
	`, databaseResourceGroup, name, acc.Region())
}

func testAccCheckIBMDatabaseInstancePostgresGroupBasic(databaseResourceGroup string, name string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "test_acc" {
		name = "%[1]s"
	}

	resource "ibm_database" "%[2]s" {
		resource_group_id            = data.ibm_resource_group.test_acc.id
		name                         = "%[2]s"
		service                      = "databases-for-postgresql"
		plan                         = "standard"
		location                     = "%[3]s"
		adminpassword                = "secure-Password12345"
		tags                         = ["one:two"]
		group {
			group_id = "member"
			members {
				allocation_count = 2
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
		service_endpoints            = "public"
		users {
			name     = "user123"
			password = "secure-Password12345"
		}
		allowlist {
			address     = "172.168.1.2/32"
			description = "desc1"
		}
	}
	`, databaseResourceGroup, name, acc.Region())
}

func testAccCheckIBMDatabaseInstancePostgresGroupFullyspecified(databaseResourceGroup string, name string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "test_acc" {
		name = "%[1]s"
	}

	resource "ibm_database" "%[2]s" {
		resource_group_id            = data.ibm_resource_group.test_acc.id
		name                         = "%[2]s"
		service                      = "databases-for-postgresql"
		plan                         = "standard"
		location                     = "%[3]s"
		adminpassword                = "secure-Password12345"
		service_endpoints            = "public-and-private"
		tags                         = ["one:two"]
		group {
			group_id = "member"
			members {
				allocation_count = 2
			}
			memory {
				allocation_mb = 8192
			}
			disk {
				allocation_mb = 7168
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
	}
	`, databaseResourceGroup, name, acc.Region())
}

func testAccCheckIBMDatabaseInstancePostgresGroupReduced(databaseResourceGroup string, name string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "test_acc" {
		name = "%[1]s"
	}

	resource "ibm_database" "%[2]s" {
		resource_group_id            = data.ibm_resource_group.test_acc.id
		name                         = "%[2]s"
		service                      = "databases-for-postgresql"
		plan                         = "standard"
		location                     = "%[3]s"
		adminpassword                = "secure-Password12345"
		service_endpoints            = "public"
		tags                         = ["one:two"]
		group {
			group_id = "member"
			members {
				allocation_count = 2
			}
			memory {
				allocation_mb = 4096
			}
			disk {
				allocation_mb = 7168
			}
			cpu {
				allocation_count = 3
			}
			host_flavor {
				id = "multitenant"
			}
		}
	}
	`, databaseResourceGroup, name, acc.Region())
}

func testAccCheckIBMDatabaseInstancePostgresGroupScaleOut(databaseResourceGroup string, name string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "test_acc" {
		name = "%[1]s"
	}

	resource "ibm_database" "%[2]s" {
		resource_group_id            = data.ibm_resource_group.test_acc.id
		name                         = "%[2]s"
		service                      = "databases-for-postgresql"
		plan                         = "standard"
		location                     = "%[3]s"
		adminpassword                = "secure-Password12345"
		group {
			group_id = "member"
			members {
				allocation_count = 3
			}
			memory {
				allocation_mb = 8192
			}
			disk {
				allocation_mb = 7168
			}
			cpu {
				allocation_count = 3
			}
			host_flavor {
				id = "multitenant"
			}
		}
		service_endpoints            = "public"
		tags                         = ["one:two"]
	}
	`, databaseResourceGroup, name, acc.Region())
}

func testAccCheckIBMDatabaseInstancePostgresImport(databaseResourceGroup string, name string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "test_acc" {
		is_default = true
		# name = "%[1]s"
	}

	resource "ibm_database" "%[2]s" {
		resource_group_id = data.ibm_resource_group.test_acc.id
		name              = "%[2]s"
		service           = "databases-for-postgresql"
		plan              = "standard"
		location          = "%[3]s"
		service_endpoints = "public-and-private"
	  }
	`, databaseResourceGroup, name, acc.Region())
}

func testAccCheckIBMDatabaseInstancePostgresMinimal(databaseResourceGroup string, name string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "test_acc" {
		is_default = true
		# name = "%[1]s"
	}

	resource "ibm_database" "%[2]s" {
		resource_group_id = data.ibm_resource_group.test_acc.id
		name              = "%[2]s"
		service           = "databases-for-postgresql"
		plan              = "standard"
		location          = "%[3]s"
		service_endpoints            = "public-and-private"
	}
	`, databaseResourceGroup, name, acc.Region())
}

func testAccCheckIBMDatabaseInstancePostgresMinimal_PITR(databaseResourceGroup string, name string) string {
	return fmt.Sprintf(`
	resource "ibm_database" "%[2]s-pitr" {
		resource_group_id                     = data.ibm_resource_group.test_acc.id
		name                                  = "%[2]s-pitr"
		service                               = "databases-for-postgresql"
		plan                                  = "standard"
		location                              = "%[3]s"
		point_in_time_recovery_deployment_id  = ibm_database.%[2]s.id
		point_in_time_recovery_time           = ""
		service_endpoints                     = "public-and-private"
	}
	`, databaseResourceGroup, name, acc.Region())
}

func testAccCheckIBMDatabaseInstancePostgresMinimal_ReadReplica(databaseResourceGroup string, name string) string {
	return fmt.Sprintf(`
	resource "ibm_database" "%[2]s-replica" {
		resource_group_id = data.ibm_resource_group.test_acc.id
		name                = "%[2]s-replica"
		service             = "databases-for-postgresql"
		plan                = "standard"
		location            = "%[3]s"
		service_endpoints   = "public-and-private"
		remote_leader_id    = ibm_database.%[2]s.id
	}
	`, databaseResourceGroup, name, acc.Region())
}

func testAccCheckIBMDatabaseInstancePostgresReadReplicaPromote(databaseResourceGroup string, readReplicaName string) string {
	return fmt.Sprintf(`
	resource "ibm_database" "%[2]s" {
		resource_group_id = data.ibm_resource_group.test_acc.id
		name                = "%[2]s"
		service             = "databases-for-postgresql"
		plan                = "standard"
		location            = "%[3]s"
		service_endpoints   = "public-and-private"
		remote_leader_id    = ""
		skip_initial_backup = true
	}
	`, databaseResourceGroup, readReplicaName, acc.Region())
}
