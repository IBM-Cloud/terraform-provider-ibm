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

func TestAccIBMDatabaseInstance_Valkey_Basic(t *testing.T) {
	t.Parallel()
	databaseResourceGroup := "default"
	var databaseInstanceOne string
	rnd := fmt.Sprintf("tf-valkey-%d", acctest.RandIntRange(10, 100))
	testName := rnd
	name := "ibm_database." + testName

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMDatabaseInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMDatabaseInstanceValkeyBasic(databaseResourceGroup, testName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(name, &databaseInstanceOne),
					resource.TestCheckResourceAttr(name, "name", testName),
					resource.TestCheckResourceAttr(name, "service", "databases-for-valkey"),
					resource.TestCheckResourceAttr(name, "plan", "standard-gen2"),
					resource.TestCheckResourceAttr(name, "location", acc.Region()),
					resource.TestCheckResourceAttr(name, "service_endpoints", "private"),
					resource.TestCheckResourceAttr(name, "groups.0.disk.0.allocation_mb", "20480"),
				),
			},
			{
				Config: testAccCheckIBMDatabaseInstanceValkeyScaled(databaseResourceGroup, testName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(name, "name", testName),
					resource.TestCheckResourceAttr(name, "service", "databases-for-valkey"),
					resource.TestCheckResourceAttr(name, "plan", "standard-gen2"),
					resource.TestCheckResourceAttr(name, "location", acc.Region()),
					resource.TestCheckResourceAttr(name, "groups.0.disk.0.allocation_mb", "30720"),
				),
			},
		},
	})
}

func TestAccIBMDatabaseInstanceValkeyImport(t *testing.T) {
	t.Parallel()
	databaseResourceGroup := "default"
	var databaseInstanceOne string
	serviceName := fmt.Sprintf("tf-valkey-%d", acctest.RandIntRange(10, 100))
	resourceName := "ibm_database." + serviceName

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMDatabaseInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMDatabaseInstanceValkeyBasic(databaseResourceGroup, serviceName),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(resourceName, &databaseInstanceOne),
					resource.TestCheckResourceAttr(resourceName, "name", serviceName),
					resource.TestCheckResourceAttr(resourceName, "service", "databases-for-valkey"),
					resource.TestCheckResourceAttr(resourceName, "plan", "standard-gen2"),
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

func TestAccIBMDatabaseInstanceValkeyKP_Encrypt(t *testing.T) {
	t.Parallel()
	databaseResourceGroup := "default"
	var databaseInstanceOne string
	rnd := fmt.Sprintf("tf-valkey-%d", acctest.RandIntRange(10, 100))
	testName := rnd
	kpInstanceName := fmt.Sprintf("tf_kp_instance_%d", acctest.RandIntRange(10, 100))
	kpKeyName := fmt.Sprintf("tf_kp_key_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMDatabaseInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMDatabaseInstanceValkeyKPEncrypt(databaseResourceGroup, kpInstanceName, kpKeyName, testName, acc.Region()),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists("ibm_database.database", &databaseInstanceOne),
					resource.TestCheckResourceAttr("ibm_database.database", "name", testName),
					resource.TestCheckResourceAttr("ibm_database.database", "service", "databases-for-valkey"),
					resource.TestCheckResourceAttrSet("ibm_database.database", "key_protect_key"),
				),
			},
		},
	})
}

func testAccCheckIBMDatabaseInstanceValkeyBasic(databaseResourceGroup string, name string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "test_acc" {
		is_default = true
	}

	resource "ibm_database" "%[2]s" {
		resource_group_id = data.ibm_resource_group.test_acc.id
		name              = "%[2]s"
		service           = "databases-for-valkey"
		plan              = "standard-gen2"
		location          = "%[3]s"
		service_endpoints = "private"

		group {
			group_id = "member"
			disk {
				allocation_mb = 20480
			}
			host_flavor {
				id = "bx3d.4x20"
			}
		}
	}
	`, databaseResourceGroup, name, acc.Region())
}

func testAccCheckIBMDatabaseInstanceValkeyScaled(databaseResourceGroup string, name string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "test_acc" {
		is_default = true
	}

	resource "ibm_database" "%[2]s" {
		resource_group_id = data.ibm_resource_group.test_acc.id
		name              = "%[2]s"
		service           = "databases-for-valkey"
		plan              = "standard-gen2"
		location          = "%[3]s"
		service_endpoints = "private"

		group {
			group_id = "member"
			disk {
				allocation_mb = 30720
			}
			host_flavor {
				id = "bx3d.4x20"
			}
		}
	}
	`, databaseResourceGroup, name, acc.Region())
}

func testAccCheckIBMDatabaseInstanceValkeyKPEncrypt(databaseResourceGroup string, kpInstanceName, kpKeyName, name, region string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "test_acc" {
		is_default = true
	}

	resource "ibm_resource_instance" "kp_instance" {
		name     = "%s"
		service  = "kms"
		plan     = "tiered-pricing"
		location = "%s"
	}

	resource "ibm_kp_key" "test" {
		key_protect_id = ibm_resource_instance.kp_instance.guid
		key_name       = "%s"
		force_delete   = true
	}

	resource "ibm_database" "database" {
		resource_group_id    = data.ibm_resource_group.test_acc.id
		name                 = "%s"
		service              = "databases-for-valkey"
		plan                 = "standard-gen2"
		location             = "%s"
		key_protect_key      = ibm_kp_key.test.id
		service_endpoints    = "private"

		group {
			group_id = "member"
			disk {
				allocation_mb = 20480
			}
			host_flavor {
				id = "bx3d.4x20"
			}
		}

		timeouts {
			create = "480m"
			update = "480m"
			delete = "15m"
		}
	}
	`, kpInstanceName, region, kpKeyName, name, region)
}
