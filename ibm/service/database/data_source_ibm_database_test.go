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

func TestAccIBMDatabaseDataSource_basic(t *testing.T) {
	t.Parallel()
	databaseResourceGroup := "default"
	var databaseInstanceOne string
	testName := fmt.Sprintf("tf-Pgress-%s", acctest.RandString(16))
	dataName := "data.ibm_database." + testName
	resourceName := "ibm_database.db"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config:  testAccCheckIBMDatabaseDataSourceConfig(databaseResourceGroup, testName),
				Destroy: false,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(resourceName, &databaseInstanceOne),
					testAccCheckIBMDatabaseInstanceExists(dataName, &databaseInstanceOne),
					resource.TestCheckResourceAttr(dataName, "name", testName),
					resource.TestCheckResourceAttr(dataName, "service", "databases-for-postgresql"),
					resource.TestCheckResourceAttr(dataName, "plan", "standard"),
					resource.TestCheckResourceAttr(dataName, "location", acc.Region()),
					resource.TestCheckResourceAttr(dataName, "adminuser", "admin"),
					resource.TestCheckResourceAttr(dataName, "groups.0.memory.0.allocation_mb", "2048"),
					resource.TestCheckResourceAttr(dataName, "groups.0.disk.0.allocation_mb", "10240"),
					resource.TestCheckResourceAttr(dataName, "allowlist.#", "0"),
					resource.TestCheckResourceAttr(dataName, "tags.#", "1"),
				),
			},
		},
	})
}

func testAccCheckIBMDatabaseDataSourceConfig(databaseResourceGroup string, name string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "test_acc" {
		is_default = true
	}

	data "ibm_database" "%[2]s" {
		resource_group_id = data.ibm_resource_group.test_acc.id
		name              = ibm_database.db.name
	}

	resource "ibm_database" "db" {
		resource_group_id = data.ibm_resource_group.test_acc.id
		name              = "%[2]s"
		service           = "databases-for-postgresql"
		plan              = "standard"
		location          = "%[3]s"
		tags              = ["one:two"]
		service_endpoints = "public"
	}

				`, databaseResourceGroup, name, acc.Region())
}
