// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package database_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMDatabaseConnectionDataSourceBasic(t *testing.T) {
	testName := fmt.Sprintf("tf-Pgress-%s", acctest.RandString(16))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMDatabaseInstancePostgresql(testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_database_connection.database_connection", "deployment_id"),
					resource.TestCheckResourceAttrSet("data.ibm_database_connection.database_connection", "user_type"),
					resource.TestCheckResourceAttrSet("data.ibm_database_connection.database_connection", "user_id"),
					resource.TestCheckResourceAttrSet("data.ibm_database_connection.database_connection", "endpoint_type"),
					resource.TestCheckResourceAttrSet("data.ibm_database_connection.database_connection", "certificate_root"),
				),
			},
		},
	})
}

func testAccCheckIBMDatabaseDataSourceConfig2(name string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "test_acc" {
		is_default = true
	}

	data "ibm_database" "%[1]s" {
		resource_group_id = data.ibm_resource_group.test_acc.id
		name              = ibm_database.db.name
	}

	resource "ibm_database" "db" {
		resource_group_id = data.ibm_resource_group.test_acc.id
		name              = "%[1]s"
		service           = "databases-for-postgresql"
		plan              = "standard"
		location          = "%[2]s"
		tags              = ["one:two"]
		service_endpoints = "public"
	}

				`, name, acc.Region())
}

func testAccCheckIBMDatabaseInstancePostgresql(name string) string {
	return testAccCheckIBMDatabaseDataSourceConfig2(name) + `
		data "ibm_database_connection" "database_connection" {
			deployment_id = ibm_database.db.id
			user_type = "database"
			user_id = "user_id"
			endpoint_type = "public"
			certificate_root = "./test/path"
		}
	  `
}
