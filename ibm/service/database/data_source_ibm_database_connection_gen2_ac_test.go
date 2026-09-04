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
// using the same single-test-step acceptance style as the legacy datasource test.
func TestAccIBMDatabaseConnectionGen2DataSourceRead(t *testing.T) {
	testName := fmt.Sprintf("tf-Pgress-gen2-Test-%s", acctest.RandString(16))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckEnterprise(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMDatabaseConnectionGen2DataSourceConfig(testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_database_connection.database_connection", "deployment_id"),
					resource.TestCheckResourceAttr("data.ibm_database_connection.database_connection", "user_type", "database"),
					resource.TestCheckResourceAttr("data.ibm_database_connection.database_connection", "user_id", testName+"-key"),
					resource.TestCheckResourceAttr("data.ibm_database_connection.database_connection", "endpoint_type", "private"),
					resource.TestCheckResourceAttrSet("data.ibm_database_connection.database_connection", "postgres.#"),
					resource.TestCheckResourceAttrSet("data.ibm_database_connection.database_connection", "postgres.0.composed.#"),
					resource.TestCheckResourceAttrSet("data.ibm_database_connection.database_connection", "postgres.0.hosts.#"),
					resource.TestCheckResourceAttrSet("data.ibm_database_connection.database_connection", "postgres.0.authentication.#"),
					resource.TestCheckResourceAttrSet("data.ibm_database_connection.database_connection", "postgres.0.database"),
					resource.TestCheckResourceAttrSet("data.ibm_database_connection.database_connection", "cli.#"),
					resource.TestCheckResourceAttrSet("data.ibm_database_connection.database_connection", "cli.0.composed.#"),
				),
			},
		},
	})
}

// TestAccIBMDatabaseConnectionGen2DataSourceInvalidID tests error handling for invalid deployment ID
func TestAccIBMDatabaseConnectionGen2DataSourceInvalidID(t *testing.T) {
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
// TestAccIBMDatabaseConnectionGen2DataSourceFallsBackToFirstKey verifies fallback
// behavior while keeping the acceptance test structure consistent with the legacy file.
func TestAccIBMDatabaseConnectionGen2DataSourceFallsBackToFirstKey(t *testing.T) {
	testName := fmt.Sprintf("tf-pg-gen2-fallback-%s", acctest.RandString(8))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckEnterprise(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMDatabaseConnectionGen2DataSourceFallbackKeyConfig(testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_database_connection.database_connection", "deployment_id"),
					resource.TestCheckResourceAttr("data.ibm_database_connection.database_connection", "user_type", "database"),
					resource.TestCheckResourceAttr("data.ibm_database_connection.database_connection", "user_id", testName+"-key-a"),
					resource.TestCheckResourceAttr("data.ibm_database_connection.database_connection", "endpoint_type", "private"),
					resource.TestCheckResourceAttrSet("data.ibm_database_connection.database_connection", "postgres.#"),
				),
			},
		},
	})
}

// TestAccIBMDatabaseConnectionGen2DataSourceS2SWarning verifies the S2S warning
// fires on a Gen2 MySQL connection data source read and is non-blocking.
func TestAccIBMDatabaseConnectionGen2DataSourceS2SWarning(t *testing.T) {
	t.Parallel()
	name := fmt.Sprintf("tf-gen2-s2s-conn-%s", acctest.RandString(8))
	dsName := "data.ibm_database_connection.gen2_s2s_conn"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckEnterprise(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMDatabaseConnectionGen2S2SConfig(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(dsName, "deployment_id"),
					resource.TestCheckResourceAttr(dsName, "user_type", "database"),
					resource.TestCheckResourceAttr(dsName, "endpoint_type", "private"),
					// mysql connection populated despite S2S warning
					resource.TestCheckResourceAttrSet(dsName, "mysql.#"),
					resource.TestCheckResourceAttrSet(dsName, "mysql.0.hosts.#"),
					resource.TestCheckResourceAttrSet(dsName, "mysql.0.composed.#"),
					// cli populated
					resource.TestCheckResourceAttrSet(dsName, "cli.#"),
					resource.TestCheckResourceAttrSet(dsName, "cli.0.composed.#"),
				),
			},
		},
	})
}

// testAccCheckIBMDatabaseConnectionGen2S2SConfig creates a Gen2 MySQL instance
// in eu-fr2 with a resource key, then reads connection info via the datasource.
// No S2S authorizations are configured so the S2S warning fires during the read.
func testAccCheckIBMDatabaseConnectionGen2S2SConfig(name string) string {
	return fmt.Sprintf(`
data "ibm_resource_group" "test_acc" {
  is_default = true
}

resource "ibm_database" "gen2_s2s" {
  resource_group_id = data.ibm_resource_group.test_acc.id
  name              = %[1]q
  service           = "databases-for-mysql"
  plan              = "standard-gen2"
  location          = "eu-fr2"
  tags              = ["terraform", "s2s-test"]

  # eu-fr2 Gen2 MySQL only allows exactly 2 members
  group {
    group_id = "member"
    members {
      allocation_count = 2
    }
  }

  timeouts {
    create = "120m"
    update = "60m"
    delete = "15m"
  }
}

resource "ibm_resource_key" "gen2_s2s_key" {
  name                 = %[2]q
  resource_instance_id = ibm_database.gen2_s2s.id
}

data "ibm_database_connection" "gen2_s2s_conn" {
  deployment_id = ibm_database.gen2_s2s.id
  user_type     = "database"
  user_id       = ibm_resource_key.gen2_s2s_key.name
  endpoint_type = "private"

  depends_on = [ibm_resource_key.gen2_s2s_key]
}
`, name, name+"-key")
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

  timeouts {
    create = "60m"
  }

  group {
    group_id = "member"
    members {
      allocation_count = 2
    }
    host_flavor {
      id = "bx3d.4x20"
    }
    disk {
      allocation_mb = 10240
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
}

data "ibm_database_connection" "database_connection" {
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
	return testAccCheckIBMDatabaseDataSourceConfigGen2(name) + `
data "ibm_database_connection" "database_connection" {
  deployment_id = ibm_database.db.id
  user_type     = "database"
  user_id       = "nonexistent-key"
  endpoint_type = "private"
}
`
}

// testAccCheckIBMDatabaseConnectionGen2DataSourceFallbackKeyConfig creates
// multiple keys but asks the datasource for a non-existent key name, which
// exercises the implementation fallback to the first available key.
func testAccCheckIBMDatabaseConnectionGen2DataSourceFallbackKeyConfig(name string) string {
	return testAccCheckIBMDatabaseDataSourceConfigGen2(name) + fmt.Sprintf(`
resource "ibm_resource_key" "db_key_a" {
  name                 = "%[1]s-key-a"
  resource_instance_id = ibm_database.db.id
}

resource "ibm_resource_key" "db_key_b" {
  name                 = "%[1]s-key-b"
  resource_instance_id = ibm_database.db.id
}

data "ibm_database_connection" "database_connection" {
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
