// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package database_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMDatabasePitrDataSourceBasic(t *testing.T) {
	testName := fmt.Sprintf("tf-Pgress-%s", acctest.RandString(16))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMDatabasePitrDataSourceConfigBasic(testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_database_pitr.database_pitr", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_database_pitr.database_pitr", "id"),
				),
			},
		},
	})
}

func testAccCheckIBMDatabaseDataSourceConfig3(name string) string {
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
		location          = "au-syd"
		tags              = ["one:two"]
	}

				`, name)
}

func testAccCheckIBMDatabasePitrDataSourceConfigBasic(name string) string {
	return testAccCheckIBMDatabaseDataSourceConfig3(name) + `
		data "ibm_database_pitr" "database_pitr" {
			id = ibm_database.db.id
		}
	`
}
