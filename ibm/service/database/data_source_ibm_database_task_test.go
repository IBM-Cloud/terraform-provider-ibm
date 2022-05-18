// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package database_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMDatabaseTaskDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMDatabaseTaskDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_database_task.database_task", "id"),
				),
			},
		},
	})
}

func testAccCheckIBMDatabaseDataSourceConfig5(name string) string {
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
	}

				`, name, acc.IcdDbRegion)
}

func testAccCheckIBMDatabaseTaskDataSourceConfigBasic() string {
	return testAccCheckIBMDatabaseDataSourceConfig5(name) + `
		data "ibm_database_task" "database_task" {
			id = taskIDLink
		}
	`
}
