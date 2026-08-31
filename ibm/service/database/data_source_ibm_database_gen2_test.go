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

// ---------------------------------------------------------------------------
// S2S authorization warning — ibm_database data source
// ---------------------------------------------------------------------------

// TestAccIBMDatabaseDataSourceGen2S2SWarning verifies that the ibm_database
// data source for a Gen2 MySQL instance without S2S authorizations configured:
//   - Fully populates all attributes (warning is non-blocking).
//   - Gen2-unsupported fields (adminuser, auto_scaling, allowlist) are empty.
//
// The S2S warning "Database backup authorization required" is emitted during the
// Read refresh that runs on every plan/apply.  resource.TestCase does not expose
// a way to assert on warnings directly; the important assertion is that state is
// fully set despite the warning.
//
// Run with:
//
//	IC_API_KEY=<key> go test -v -timeout 240m \
//	  -run TestAccIBMDatabaseDataSourceGen2S2SWarning \
//	  ./ibm/service/database/...
func TestAccIBMDatabaseDataSourceGen2S2SWarning(t *testing.T) {
	t.Parallel()
	name := fmt.Sprintf("tf-gen2-s2s-ds-%s", acctest.RandString(8))
	dsName := "data.ibm_database.gen2_s2s_lookup"
	resName := "ibm_database.gen2_s2s"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckEnterprise(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				// State must be fully populated even though the S2S warning fires.
				Config: testAccCheckIBMDatabaseDataSourceGen2S2SConfig(name),
				Check: resource.ComposeTestCheckFunc(
					// resource fully created
					resource.TestCheckResourceAttrSet(resName, "id"),
					resource.TestCheckResourceAttr(resName, "name", name),
					resource.TestCheckResourceAttr(resName, "service", "databases-for-mysql"),
					resource.TestCheckResourceAttr(resName, "plan", "standard-gen2"),
					resource.TestCheckResourceAttr(resName, "location", "eu-fr2"),
					// data source fully populated
					resource.TestCheckResourceAttrSet(dsName, "id"),
					resource.TestCheckResourceAttr(dsName, "name", name),
					resource.TestCheckResourceAttr(dsName, "service", "databases-for-mysql"),
					resource.TestCheckResourceAttr(dsName, "plan", "standard-gen2"),
					resource.TestCheckResourceAttr(dsName, "location", "eu-fr2"),
					resource.TestCheckResourceAttrSet(dsName, "groups.#"),
					// Gen2: unsupported attributes must be empty
					resource.TestCheckResourceAttr(dsName, "adminuser", ""),
					resource.TestCheckResourceAttr(dsName, "auto_scaling.#", "0"),
					resource.TestCheckResourceAttr(dsName, "allowlist.#", "0"),
				),
			},
		},
	})
}

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

// testAccCheckIBMDatabaseDataSourceGen2S2SConfig creates a Gen2 MySQL instance
// in eu-fr2 (exactly 2 members required) and a data source that looks it up by
// name.  No S2S authorizations are configured, so the S2S warning fires on read.
func testAccCheckIBMDatabaseDataSourceGen2S2SConfig(name string) string {
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

data "ibm_database" "gen2_s2s_lookup" {
  resource_group_id = data.ibm_resource_group.test_acc.id
  name              = ibm_database.gen2_s2s.name
}
`, name)
}
