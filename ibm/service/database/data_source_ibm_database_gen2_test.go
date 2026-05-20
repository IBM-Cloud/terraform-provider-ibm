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
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckEnterprise(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMDatabaseDataSourceGen2ConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_database.database_gen2", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_database.database_gen2", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_database.database_gen2", "status"),
					resource.TestCheckResourceAttrSet("data.ibm_database.database_gen2", "version"),
					resource.TestCheckResourceAttrSet("data.ibm_database.database_gen2", "guid"),
					resource.TestCheckResourceAttrSet("data.ibm_database.database_gen2", "location"),
					// Test Gen2-specific behavior: unsupported attributes should be empty
					resource.TestCheckResourceAttr("data.ibm_database.database_gen2", "adminuser", ""),
					resource.TestCheckResourceAttr("data.ibm_database.database_gen2", "adminpassword", ""),
					resource.TestCheckResourceAttr("data.ibm_database.database_gen2", "users.#", "0"),
					resource.TestCheckResourceAttr("data.ibm_database.database_gen2", "allowlist.#", "0"),
					resource.TestCheckResourceAttr("data.ibm_database.database_gen2", "auto_scaling.#", "0"),
					resource.TestCheckResourceAttr("data.ibm_database.database_gen2", "configuration_schema", ""),
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
			resource.TestStep{
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
			resource.TestStep{
				Config:      testAccCheckIBMDatabaseDataSourceGen2InvalidIDConfig(),
				ExpectError: regexp.MustCompile("No resource instance found|invalid"),
			},
		},
	})
}

func testAccCheckIBMDatabaseDataSourceGen2ConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_database" "database_gen2" {
			name = "%[1]s"
		}
	`, acc.IcdDbGen2DeploymentId)
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
