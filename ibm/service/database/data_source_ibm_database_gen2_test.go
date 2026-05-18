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

// TestAccIBMDatabaseDataSourceGen2Basic tests basic Gen2 data source read
func TestAccIBMDatabaseDataSourceGen2Basic(t *testing.T) {
	testName := fmt.Sprintf("tf-gen2-db-%s", acctest.RandString(16))
	dataName := "data.ibm_database." + testName

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckEnterprise(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMDatabaseDataSourceGen2Config(testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(dataName, "name", testName),
					resource.TestCheckResourceAttr(dataName, "service", "databases-for-mongodb"),
					resource.TestCheckResourceAttr(dataName, "plan", "standard-gen2"),
					resource.TestCheckResourceAttr(dataName, "location", "ca-mon"),
					resource.TestCheckResourceAttrSet(dataName, "status"),
					resource.TestCheckResourceAttrSet(dataName, "version"),
					resource.TestCheckResourceAttrSet(dataName, "guid"),
					resource.TestCheckResourceAttr(dataName, "groups.#", "1"),
					resource.TestCheckResourceAttrSet(dataName, "groups.0.memory.0.allocation_mb"),
					resource.TestCheckResourceAttrSet(dataName, "groups.0.disk.0.allocation_mb"),
					resource.TestCheckResourceAttrSet(dataName, "groups.0.cpu.0.allocation_count"),
					// Verify Gen2-unsupported attributes are NOT set
					resource.TestCheckResourceAttr(dataName, "adminuser", ""),
					resource.TestCheckResourceAttr(dataName, "adminpassword", ""),
					resource.TestCheckResourceAttr(dataName, "auto_scaling.#", "0"),
					resource.TestCheckResourceAttr(dataName, "allowlist.#", "0"),
					resource.TestCheckResourceAttr(dataName, "users.#", "0"),
					resource.TestCheckResourceAttr(dataName, "configuration_schema", ""),
				),
			},
		},
	})
}

// TestAccIBMDatabaseDataSourceGen2UnsupportedAttributes validates that
// Gen2-unsupported attributes remain empty even when explicitly checked
func TestAccIBMDatabaseDataSourceGen2UnsupportedAttributes(t *testing.T) {
	testName := fmt.Sprintf("tf-gen2-unsup-%s", acctest.RandString(16))
	dataName := "data.ibm_database." + testName

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckEnterprise(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMDatabaseDataSourceGen2Config(testName),
				Check: resource.ComposeTestCheckFunc(
					// Explicitly verify each unsupported attribute is empty
					resource.TestCheckResourceAttr(dataName, "adminuser", ""),
					resource.TestCheckResourceAttr(dataName, "adminpassword", ""),
					resource.TestCheckResourceAttr(dataName, "auto_scaling.#", "0"),
					resource.TestCheckResourceAttr(dataName, "allowlist.#", "0"),
					resource.TestCheckResourceAttr(dataName, "users.#", "0"),
					resource.TestCheckResourceAttr(dataName, "configuration_schema", ""),
				),
			},
		},
	})
}

// TestAccIBMDatabaseDataSourceGen2InvalidInput tests error handling
// for invalid database name
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

func testAccCheckIBMDatabaseDataSourceGen2Config(name string) string {
	return fmt.Sprintf(`
data "ibm_resource_group" "test_acc" {
  is_default = true
}

resource "ibm_database" "db" {
  name              = "%[1]s"
  plan              = "standard-gen2"
  location          = "ca-mon"
  service           = "databases-for-mongodb"
  resource_group_id = data.ibm_resource_group.test_acc.id
  
  group {
    group_id = "member"
    host_flavor {
      id = "bx3d.4x20"
    }
    disk {
      allocation_mb = 20480
    }
  }

  tags = ["gen2-test"]
}

data "ibm_database" "%[1]s" {
  name              = ibm_database.db.name
  resource_group_id = data.ibm_resource_group.test_acc.id
  location          = ibm_database.db.location
  service           = ibm_database.db.service
}
`, name)
}

func testAccCheckIBMDatabaseDataSourceGen2InvalidConfig() string {
	return fmt.Sprintf(`
data "ibm_resource_group" "test_acc" {
  is_default = true
}

data "ibm_database" "nonexistent" {
  name              = "this-database-does-not-exist-gen2-%s"
  resource_group_id = data.ibm_resource_group.test_acc.id
  location          = "ca-mon"
  service           = "databases-for-mongodb"
}
`, acctest.RandString(8))
}

func testAccCheckIBMDatabaseGen2CheckPlansConfig(service string) string {
	return fmt.Sprintf(`
data "ibm_resource_group" "test_acc" {
  is_default = true
}

resource "ibm_database" "gen2_check" {
  name              = "tf-gen2-check-%s"
  plan              = "standard-gen2"
  location          = "ca-mon"
  service           = "%s"
  resource_group_id = data.ibm_resource_group.test_acc.id
  
  group {
    group_id = "member"
    host_flavor {
      id = "bx3d.4x20"
    }
    disk {
      allocation_mb = 20480
    }
  }

  tags = ["gen2-availability-check"]
}
`, acctest.RandString(8), service)
}
