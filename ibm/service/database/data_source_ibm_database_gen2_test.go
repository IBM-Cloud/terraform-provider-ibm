// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package database_test

import (
	"fmt"
	"regexp"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMDatabaseDataSourceGen2Basic(t *testing.T) {
	t.Parallel()
	var databaseInstanceOne string
	testName := acc.IcdDbGen2DeploymentId
	dataName := "data.ibm_database.test"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckEnterprise(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMDatabaseDataSourceGen2Config(testName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(dataName, &databaseInstanceOne),
					resource.TestCheckResourceAttr(dataName, "name", testName),
					resource.TestCheckResourceAttr(dataName, "service", "databases-for-postgresql"),
					resource.TestCheckResourceAttr(dataName, "plan", "standard-gen2"),
					resource.TestCheckResourceAttrSet(dataName, "location"),
					resource.TestCheckResourceAttrSet(dataName, "groups.0.memory.0.allocation_mb"),
					resource.TestCheckResourceAttrSet(dataName, "groups.0.disk.0.allocation_mb"),
					// Test Gen2-specific behavior: unsupported attributes should be empty
					resource.TestCheckResourceAttr(dataName, "adminuser", ""),
					resource.TestCheckResourceAttr(dataName, "adminpassword", ""),
					resource.TestCheckResourceAttr(dataName, "users.#", "0"),
					resource.TestCheckResourceAttr(dataName, "allowlist.#", "0"),
					resource.TestCheckResourceAttr(dataName, "auto_scaling.#", "0"),
					resource.TestCheckResourceAttr(dataName, "configuration_schema", ""),
				),
			},
		},
	})
}

func TestAccIBMDatabaseDataSourceGen2WithResourceGroupID(t *testing.T) {
	t.Parallel()
	var databaseInstanceOne string
	testName := acc.IcdDbGen2DeploymentId
	dataName := "data.ibm_database.test"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckEnterprise(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMDatabaseDataSourceGen2ConfigWithResourceGroup(testName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(dataName, &databaseInstanceOne),
					resource.TestCheckResourceAttr(dataName, "name", testName),
					resource.TestCheckResourceAttr(dataName, "service", "databases-for-postgresql"),
					resource.TestCheckResourceAttr(dataName, "plan", "standard-gen2"),
					resource.TestCheckResourceAttrSet(dataName, "location"),
					resource.TestCheckResourceAttrSet(dataName, "resource_group_id"),
					resource.TestCheckResourceAttrSet(dataName, "groups.0.memory.0.allocation_mb"),
					resource.TestCheckResourceAttrSet(dataName, "groups.0.disk.0.allocation_mb"),
				),
			},
		},
	})
}

func TestAccIBMDatabaseDataSourceGen2InvalidInput(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckEnterprise(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckIBMDatabaseDataSourceGen2InvalidConfig(),
				ExpectError: regexp.MustCompile("No resource instance found"),
			},
		},
	})
}

func TestAccIBMDatabaseDataSourceGen2InvalidID(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckEnterprise(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckIBMDatabaseDataSourceGen2InvalidIDConfig(),
				ExpectError: regexp.MustCompile("No resource instance found|invalid"),
			},
		},
	})
}

func testAccCheckIBMDatabaseDataSourceGen2Config(name string) string {
	return fmt.Sprintf(`
	data "ibm_database" "test" {
		name = "%s"
	}
	`, name)
}

func testAccCheckIBMDatabaseDataSourceGen2ConfigWithResourceGroup(name string) string {
	return fmt.Sprintf(`
	data "ibm_database" "test" {
		name = "%s"
	}
	`, name)
}

func testAccCheckIBMDatabaseDataSourceGen2InvalidConfig() string {
	return `
		data "ibm_database" "nonexistent" {
			name = "this-database-does-not-exist-gen2-test"
		}
	`
}

func testAccCheckIBMDatabaseDataSourceGen2InvalidIDConfig() string {
	return `
		data "ibm_database" "invalid_id" {
			name = "invalid@#$%^&*()id"
		}
	`
}
