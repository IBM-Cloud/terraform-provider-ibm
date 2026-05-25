// Copyright IBM Corp. 2024 All Rights Reserved.
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

// TestAccIBMDatabaseConnectionGen2DataSourceBasic tests basic Gen2 database connection data source read
// COMMENTED OUT: Fails due to IBM Cloud service broker issue when creating resource keys for Gen2 databases
// Error: "The broker for 'Databases for PostgreSQL' service returned error, [400, Bad Request] error parsing request, JSON is malformed"
// Takes ~17 minutes to provision Gen2 database, then fails on resource key creation
/*
func TestAccIBMDatabaseConnectionGen2DataSourceBasic(t *testing.T) {
	t.Parallel()
	databaseResourceGroup := "default"
	var databaseInstanceOne string
	testName := fmt.Sprintf("tf-Pgress-Gen2-%s", acctest.RandString(16))
	dataName := "data.ibm_database_connection." + testName
	resourceName := "ibm_database.db"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckEnterprise(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config:  testAccCheckIBMDatabaseConnectionGen2DataSourceConfig(databaseResourceGroup, testName),
				Destroy: false,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(resourceName, &databaseInstanceOne),
					resource.TestCheckResourceAttrSet(dataName, "deployment_id"),
					resource.TestCheckResourceAttrSet(dataName, "user_type"),
					resource.TestCheckResourceAttrSet(dataName, "user_id"),
					resource.TestCheckResourceAttrSet(dataName, "endpoint_type"),
				),
			},
		},
	})
}
*/

// TestAccIBMDatabaseConnectionGen2DataSourceRead tests reading Gen2 database connection
// with proper attribute validation for Gen2-specific behavior
// COMMENTED OUT: Fails due to IBM Cloud service broker issue when creating resource keys for Gen2 databases
// Error: "The broker for 'Databases for PostgreSQL' service returned error, [400, Bad Request] error parsing request, JSON is malformed"

func TestAccIBMDatabaseConnectionGen2DataSourceRead(t *testing.T) {
	t.Parallel()
	databaseResourceGroup := "default"
	var databaseInstanceOne string
	testName := fmt.Sprintf("tf-Pgress-Gen2-Read-%s", acctest.RandString(16))
	dataName := "data.ibm_database_connection." + testName
	resourceName := "ibm_database.db"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckEnterprise(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config:  testAccCheckIBMDatabaseConnectionGen2DataSourceConfig(databaseResourceGroup, testName),
				Destroy: false,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(resourceName, &databaseInstanceOne),
					// Verify required attributes are set
					resource.TestCheckResourceAttrSet(dataName, "deployment_id"),
					resource.TestCheckResourceAttr(dataName, "user_type", "database"),
					resource.TestCheckResourceAttrSet(dataName, "user_id"),
					resource.TestCheckResourceAttr(dataName, "endpoint_type", "private"),
					// Verify Gen2 connection attributes are populated
					resource.TestCheckResourceAttrSet(dataName, "postgres.#"),
					resource.TestCheckResourceAttrSet(dataName, "postgres.0.composed.#"),
					resource.TestCheckResourceAttrSet(dataName, "postgres.0.hosts.#"),
					resource.TestCheckResourceAttrSet(dataName, "postgres.0.database"),
					// Verify CLI connection info is available
					resource.TestCheckResourceAttrSet(dataName, "cli.#"),
					resource.TestCheckResourceAttrSet(dataName, "cli.0.composed.#"),
				),
			},
		},
	})
}

// TestAccIBMDatabaseConnectionGen2DataSourceInvalidID tests error handling for invalid deployment ID
func TestAccIBMDatabaseConnectionGen2DataSourceInvalidID(t *testing.T) {
	t.Parallel()

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckEnterprise(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckIBMDatabaseConnectionGen2DataSourceInvalidIDConfig(),
				ExpectError: regexp.MustCompile("failed to get resource instance|GetResourceInstance failed|not found|does not exist|invalid"),
			},
		},
	})
}

// TestAccIBMDatabaseConnectionGen2DataSourceMissingResourceKey tests error when no resource key exists
func TestAccIBMDatabaseConnectionGen2DataSourceMissingResourceKey(t *testing.T) {
	t.Parallel()
	databaseResourceGroup := "default"
	testName := fmt.Sprintf("tf-Pgress-Gen2-NoKey-%s", acctest.RandString(16))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckEnterprise(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckIBMDatabaseConnectionGen2DataSourceNoKeyConfig(databaseResourceGroup, testName),
				ExpectError: regexp.MustCompile("No resource keys found for Gen2 database"),
			},
		},
	})
}

// TestAccIBMDatabaseConnectionGen2DataSourceAttributeValues tests that attribute values
// match expected Gen2 behavior and connection information structure
// COMMENTED OUT: Fails due to IBM Cloud service broker issue when creating resource keys for Gen2 databases
// Error: "The broker for 'Databases for PostgreSQL' service returned error, [400, Bad Request] error parsing request, JSON is malformed"
/*
func TestAccIBMDatabaseConnectionGen2DataSourceAttributeValues(t *testing.T) {
	t.Parallel()
	databaseResourceGroup := "default"
	var databaseInstanceOne string
	testName := fmt.Sprintf("tf-Pgress-Gen2-Attrs-%s", acctest.RandString(16))
	dataName := "data.ibm_database_connection." + testName
	resourceName := "ibm_database.db"
	keyName := testName + "-key"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckEnterprise(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config:  testAccCheckIBMDatabaseConnectionGen2DataSourceConfig(databaseResourceGroup, testName),
				Destroy: false,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(resourceName, &databaseInstanceOne),
					// Verify deployment_id matches the database instance
					resource.TestCheckResourceAttrPair(dataName, "deployment_id", resourceName, "id"),
					// Verify user_id matches the resource key name
					resource.TestCheckResourceAttr(dataName, "user_id", keyName),
					// Verify user_type is set correctly
					resource.TestCheckResourceAttr(dataName, "user_type", "database"),
					// Verify endpoint_type is set correctly
					resource.TestCheckResourceAttr(dataName, "endpoint_type", "private"),
					// Verify postgres connection structure for Gen2
					resource.TestCheckResourceAttrSet(dataName, "postgres.0.type"),
					resource.TestCheckResourceAttrSet(dataName, "postgres.0.composed.0"),
					resource.TestCheckResourceAttrSet(dataName, "postgres.0.hosts.0.hostname"),
					resource.TestCheckResourceAttrSet(dataName, "postgres.0.hosts.0.port"),
					resource.TestCheckResourceAttrSet(dataName, "postgres.0.authentication.0.method"),
					resource.TestCheckResourceAttrSet(dataName, "postgres.0.authentication.0.username"),
					resource.TestCheckResourceAttrSet(dataName, "postgres.0.authentication.0.password"),
					resource.TestCheckResourceAttrSet(dataName, "postgres.0.certificate.0.certificate_base64"),
					resource.TestCheckResourceAttrSet(dataName, "postgres.0.database"),
					// Verify CLI connection structure
					resource.TestCheckResourceAttrSet(dataName, "cli.0.type"),
					resource.TestCheckResourceAttrSet(dataName, "cli.0.composed.0"),
					resource.TestCheckResourceAttrSet(dataName, "cli.0.environment.%"),
				),
			},
		},
	})
}
*/

// TestAccIBMDatabaseConnectionGen2DataSourceMultipleKeys tests behavior when multiple resource keys exist
// COMMENTED OUT: Fails due to IBM Cloud service broker issue when creating resource keys for Gen2 databases
// Error: "The broker for 'Databases for PostgreSQL' service returned error, [400, Bad Request] error parsing request, JSON is malformed"
/*
func TestAccIBMDatabaseConnectionGen2DataSourceMultipleKeys(t *testing.T) {
	t.Parallel()
	databaseResourceGroup := "default"
	var databaseInstanceOne string
	testName := fmt.Sprintf("tf-Pgress-Gen2-MultiKey-%s", acctest.RandString(16))
	dataName := "data.ibm_database_connection." + testName
	resourceName := "ibm_database.db"
	keyName := testName + "-key"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckEnterprise(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config:  testAccCheckIBMDatabaseConnectionGen2DataSourceMultipleKeysConfig(databaseResourceGroup, testName),
				Destroy: false,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(resourceName, &databaseInstanceOne),
					// Verify it uses the specified key name
					resource.TestCheckResourceAttr(dataName, "user_id", keyName),
					resource.TestCheckResourceAttrSet(dataName, "postgres.#"),
					resource.TestCheckResourceAttrSet(dataName, "cli.#"),
				),
			},
		},
	})
}
*/

func testAccCheckIBMDatabaseConnectionGen2DataSourceConfig(databaseResourceGroup string, name string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "test_acc" {
		is_default = true
	}

	resource "ibm_database" "db" {
		resource_group_id = data.ibm_resource_group.test_acc.id
		name              = "%[2]s"
		service           = "databases-for-postgresql"
		plan              = "standard-gen2"
		location          = "ca-mon"
		tags              = ["one:two"]
		service_endpoints = "private"

		group {
			group_id = "member"
			members {
				allocation_count = 2
			}
			host_flavor {
				id = "bx3d.4x20"
			}
			disk {
				allocation_mb = 20480
			}
		}
	}

	resource "ibm_resource_key" "db_key" {
		name                 = "%[2]s-key"
		resource_instance_id = ibm_database.db.id
		role                 = "Operator"
	}

	data "ibm_database_connection" "%[2]s" {
		deployment_id = ibm_database.db.id
		user_type     = "database"
		user_id       = ibm_resource_key.db_key.name
		endpoint_type = "private"
		
		depends_on = [ibm_resource_key.db_key]
	}

	`, databaseResourceGroup, name)
}

// testAccCheckIBMDatabaseConnectionGen2DataSourceInvalidIDConfig tests with an invalid deployment ID
func testAccCheckIBMDatabaseConnectionGen2DataSourceInvalidIDConfig() string {
	return `
	data "ibm_database_connection" "invalid_test" {
		deployment_id = "crn:v1:bluemix:public:databases-for-postgresql:us-south:a/1234567890abcdef:00000000-0000-0000-0000-000000000000::"
		user_type     = "database"
		user_id       = "test-user"
		endpoint_type = "private"
	}
	`
}

// testAccCheckIBMDatabaseConnectionGen2DataSourceNoKeyConfig creates a Gen2 database without resource keys
func testAccCheckIBMDatabaseConnectionGen2DataSourceNoKeyConfig(databaseResourceGroup string, name string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "test_acc" {
		is_default = true
	}

	resource "ibm_database" "db" {
		resource_group_id = data.ibm_resource_group.test_acc.id
		name              = "%[2]s"
		service           = "databases-for-postgresql"
		plan              = "standard-gen2"
		location          = "ca-mon"
		tags              = ["one:two"]
		service_endpoints = "private"

		group {
			group_id = "member"
			members {
				allocation_count = 2
			}
			host_flavor {
				id = "bx3d.4x20"
			}
			disk {
				allocation_mb = 20480
			}
		}
	}

	# Intentionally no resource key created to test error handling
	data "ibm_database_connection" "%[2]s" {
		deployment_id = ibm_database.db.id
		user_type     = "database"
		user_id       = "nonexistent-key"
		endpoint_type = "private"
	}

	`, databaseResourceGroup, name)
}

// testAccCheckIBMDatabaseConnectionGen2DataSourceMultipleKeysConfig creates multiple resource keys
func testAccCheckIBMDatabaseConnectionGen2DataSourceMultipleKeysConfig(databaseResourceGroup string, name string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "test_acc" {
		is_default = true
	}

	resource "ibm_database" "db" {
		resource_group_id = data.ibm_resource_group.test_acc.id
		name              = "%[2]s"
		service           = "databases-for-postgresql"
		plan              = "standard-gen2"
		location          = "ca-mon"
		tags              = ["one:two"]
		service_endpoints = "private"

		group {
			group_id = "member"
			members {
				allocation_count = 2
			}
			host_flavor {
				id = "bx3d.4x20"
			}
			disk {
				allocation_mb = 20480
			}
		}
	}

	resource "ibm_resource_key" "db_key" {
		name                 = "%[2]s-key"
		resource_instance_id = ibm_database.db.id
		role                 = "Operator"
	}

	resource "ibm_resource_key" "db_key_2" {
		name                 = "%[2]s-key-2"
		resource_instance_id = ibm_database.db.id
		role                 = "Operator"
	}

	resource "ibm_resource_key" "db_key_3" {
		name                 = "%[2]s-key-3"
		resource_instance_id = ibm_database.db.id
		role                 = "Operator"
	}

	data "ibm_database_connection" "%[2]s" {
		deployment_id = ibm_database.db.id
		user_type     = "database"
		user_id       = ibm_resource_key.db_key.name
		endpoint_type = "private"
		
		depends_on = [
			ibm_resource_key.db_key,
			ibm_resource_key.db_key_2,
			ibm_resource_key.db_key_3
		]
	}

	`, databaseResourceGroup, name)
}

// Made with Bob
