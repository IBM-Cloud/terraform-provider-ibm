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

func TestAccIBMDatabaseRemotesDataSourceBasic(t *testing.T) {

	testName := fmt.Sprintf("tf-Pgress-%s", acctest.RandString(16))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMDatabaseRemotesDataSourceConfigBasic(testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_database_remotes.database_remotes", "deployment_id"),
					resource.TestCheckResourceAttrSet("data.ibm_database_remotes.database_remotes_replica", "deployment_id"),
					resource.TestCheckResourceAttr("data.ibm_database_remotes.database_remotes", "leader", ""),
					resource.TestCheckResourceAttrSet("data.ibm_database_remotes.database_remotes_replica", "leader"),
					resource.TestCheckResourceAttrSet("data.ibm_database_remotes.database_remotes", "replicas.#"),
				),
			},
		},
	},
	)
}

func testAccCheckIBMDatabaseDataSourceConfig4(name string) string {
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

	resource "ibm_database" "db_replica" {
		resource_group_id = data.ibm_resource_group.test_acc.id
    	remote_leader_id  = ibm_database.db.id
		name              = "%[1]s-replica"
		service           = "databases-for-postgresql"
		plan              = "standard"
		location          = "%[2]s"
		tags              = ["one:two"]

    depends_on = [
      ibm_database.db,
    ]
	}

				`, name, acc.IcdDbRegion)
}

func testAccCheckIBMDatabaseRemotesDataSourceConfigBasic(name string) string {
	return testAccCheckIBMDatabaseDataSourceConfig4(name) + `
		data "ibm_database_remotes" "database_remotes_replica" {
			deployment_id = ibm_database.db_replica.id

		depends_on = [
			ibm_database.db_replica,
		]
		}

		data "ibm_database_remotes" "database_remotes" {
			deployment_id = ibm_database.db.id
			
		depends_on = [
			ibm_database.db_replica,
		]
		}
	`
}
