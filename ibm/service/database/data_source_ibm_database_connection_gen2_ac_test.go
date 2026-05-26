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

// TestAccIBMDatabaseConnectionGen2DataSourceRead validates the Gen2 datasource
// when a matching resource key exists for the requested user_id.
// Note: this depends on IBM Cloud successfully creating a resource key for a Gen2
// PostgreSQL deployment, which may still be affected by broker-side issues.
func TestAccIBMDatabaseConnectionGen2DataSourceRead(t *testing.T) {
	t.Parallel()

	testName := fmt.Sprintf("tf-pg-gen2-read-%s", acctest.RandString(8))
	dataName := "data.ibm_database_connection." + testName
	resourceName := "ibm_database.db"
	expectedKeyName := testName + "-key"
	var databaseInstanceOne string

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckEnterprise(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config:  testAccCheckIBMDatabaseConnectionGen2DataSourceConfig(testName),
				Destroy: false,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(resourceName, &databaseInstanceOne),
					resource.TestCheckResourceAttrPair(dataName, "deployment_id", resourceName, "id"),
					resource.TestCheckResourceAttr(dataName, "user_type", "database"),
					resource.TestCheckResourceAttr(dataName, "user_id", expectedKeyName),
					resource.TestCheckResourceAttr(dataName, "endpoint_type", "private"),
					resource.TestCheckResourceAttrSet(dataName, "postgres.#"),
					resource.TestCheckResourceAttrSet(dataName, "postgres.0.composed.#"),
					resource.TestCheckResourceAttrSet(dataName, "postgres.0.hosts.#"),
					resource.TestCheckResourceAttrSet(dataName, "postgres.0.authentication.#"),
					resource.TestCheckResourceAttrSet(dataName, "postgres.0.database"),
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

// TestAccIBMDatabaseConnectionGen2DataSourceMissingResourceKey verifies the
// datasource error path when no keys exist for the Gen2 deployment.
func TestAccIBMDatabaseConnectionGen2DataSourceMissingResourceKey(t *testing.T) {
	t.Parallel()

	testName := fmt.Sprintf("tf-pg-gen2-nokey-%s", acctest.RandString(8))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckEnterprise(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckIBMDatabaseConnectionGen2DataSourceNoKeyConfig(testName),
				ExpectError: regexp.MustCompile("No resource keys found for Gen2 database"),
			},
		},
	})
}

// TestAccIBMDatabaseConnectionGen2DataSourceFallsBackToFirstKey verifies the
// implementation fallback path: if the requested key name is not found, the
// datasource uses the first available resource key returned by the API.
func TestAccIBMDatabaseConnectionGen2DataSourceFallsBackToFirstKey(t *testing.T) {
	t.Parallel()

	testName := fmt.Sprintf("tf-pg-gen2-fallback-%s", acctest.RandString(8))
	dataName := "data.ibm_database_connection." + testName
	resourceName := "ibm_database.db"
	expectedFallbackKeyName := testName + "-key-a"
	var databaseInstanceOne string

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckEnterprise(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config:  testAccCheckIBMDatabaseConnectionGen2DataSourceFallbackKeyConfig(testName),
				Destroy: false,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(resourceName, &databaseInstanceOne),
					resource.TestCheckResourceAttrPair(dataName, "deployment_id", resourceName, "id"),
					resource.TestCheckResourceAttr(dataName, "user_type", "database"),
					resource.TestCheckResourceAttr(dataName, "user_id", expectedFallbackKeyName),
					resource.TestCheckResourceAttr(dataName, "endpoint_type", "private"),
					resource.TestCheckResourceAttrSet(dataName, "postgres.#"),
				),
			},
		},
	})
}

// testAccCheckIBMDatabaseDataSourceConfigGen2 creates a Gen2 PostgreSQL instance.
func testAccCheckIBMDatabaseDataSourceConfigGen2(name string) string {
	return fmt.Sprintf(`
data "ibm_resource_group" "test_acc" {
  is_default = true
}

resource "ibm_database" "db" {
  resource_group_id = data.ibm_resource_group.test_acc.id
  name              = "%[1]s"
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
`, name)
}

func testAccCheckIBMDatabaseConnectionGen2DataSourceConfig(name string) string {
	return testAccCheckIBMDatabaseDataSourceConfigGen2(name) + fmt.Sprintf(`
resource "ibm_resource_key" "db_key" {
  name                 = "%[1]s-key"
  resource_instance_id = ibm_database.db.id
  role                 = "Operator"
}

data "ibm_database_connection" "%[1]s" {
  deployment_id = ibm_database.db.id
  user_type     = "database"
  user_id       = ibm_resource_key.db_key.name
  endpoint_type = "private"

  depends_on = [ibm_resource_key.db_key]
}
`, name)
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

// testAccCheckIBMDatabaseConnectionGen2DataSourceNoKeyConfig creates a Gen2
// database without any resource keys so the datasource returns the expected error.
func testAccCheckIBMDatabaseConnectionGen2DataSourceNoKeyConfig(name string) string {
	return testAccCheckIBMDatabaseDataSourceConfigGen2(name) + fmt.Sprintf(`
data "ibm_database_connection" "%[1]s" {
  deployment_id = ibm_database.db.id
  user_type     = "database"
  user_id       = "nonexistent-key"
  endpoint_type = "private"
}
`, name)
}

// testAccCheckIBMDatabaseConnectionGen2DataSourceFallbackKeyConfig creates
// multiple keys but asks the datasource for a non-existent key name, which
// exercises the implementation fallback to the first available key.
func testAccCheckIBMDatabaseConnectionGen2DataSourceFallbackKeyConfig(name string) string {
	return testAccCheckIBMDatabaseDataSourceConfigGen2(name) + fmt.Sprintf(`
resource "ibm_resource_key" "db_key_a" {
  name                 = "%[1]s-key-a"
  resource_instance_id = ibm_database.db.id
  role                 = "Operator"
}

resource "ibm_resource_key" "db_key_b" {
  name                 = "%[1]s-key-b"
  resource_instance_id = ibm_database.db.id
  role                 = "Operator"
}

data "ibm_database_connection" "%[1]s" {
  deployment_id = ibm_database.db.id
  user_type     = "database"
  user_id       = "%[1]s-key-does-not-exist"
  endpoint_type = "private"

  depends_on = [
    ibm_resource_key.db_key_a,
    ibm_resource_key.db_key_b,
  ]
}
`, name)
}
