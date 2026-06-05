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

// TestAccIBMDatabaseDataSourceGen2Basic tests the Gen2 database data source.
// Note: This test creates a real database instance which can take 30-60 minutes.
// Run with: go test -timeout 120m -run TestAccIBMDatabaseDataSourceGen2Basic ./ibm/service/database/...
func TestAccIBMDatabaseDataSourceGen2Basic(t *testing.T) {
	t.Parallel()
	testName := fmt.Sprintf("tf-gen2-db-%s", acctest.RandString(10))
	dataName := "data.ibm_database.test"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckEnterprise(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMDatabaseDataSourceGen2Config(testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(dataName, "id"),
					resource.TestCheckResourceAttrSet(dataName, "name"),
					resource.TestCheckResourceAttrSet(dataName, "service"),
					resource.TestCheckResourceAttrSet(dataName, "plan"),
					resource.TestCheckResourceAttrSet(dataName, "location"),
					resource.TestCheckResourceAttr(dataName, "name", testName),
					resource.TestCheckResourceAttr(dataName, "service", "databases-for-postgresql"),
					resource.TestCheckResourceAttr(dataName, "plan", "standard-gen2"),
					// Verify Gen2-specific behavior: adminuser, auto_scaling, and allowlist are empty/nil
					resource.TestCheckResourceAttr(dataName, "adminuser", ""),
					resource.TestCheckResourceAttr(dataName, "auto_scaling.#", "0"),
					resource.TestCheckResourceAttr(dataName, "allowlist.#", "0"),
					// Verify groups are set
					resource.TestCheckResourceAttrSet(dataName, "groups.#"),
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
	data "ibm_resource_group" "test_acc" {
		is_default = true
	}

	resource "ibm_database" "test" {
		resource_group_id = data.ibm_resource_group.test_acc.id
		name              = "%[1]s"
		service           = "databases-for-postgresql"
		plan              = "standard-gen2"
		location          = "ca-mon"
		service_endpoints = "private"

		group {
			group_id = "member"

			host_flavor {
				id = "bx3d.4x20"
			}

			disk {
				allocation_mb = 20480
			}
		}

		timeouts {
			create = "120m"
			update = "60m"
			delete = "60m"
		}
	}

	data "ibm_database" "test" {
		resource_group_id = data.ibm_resource_group.test_acc.id
		name              = ibm_database.test.name
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
