// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package database_test

import (
	"fmt"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMDatabaseTasksDataSourceBasic(t *testing.T) {
	testName := fmt.Sprintf("tf-Pgress-%s", acctest.RandString(16))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMDatabaseTasksDataSourceConfigBasic(testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_database_tasks.database_tasks", "deployment_id"),
					resource.TestCheckResourceAttrSet("data.ibm_database_tasks.database_tasks", "tasks.#"),
				),
			},
		},
	})
}

func testAccCheckIBMDatabaseTasksDataSourceConfigBasic(name string) string {
	return fmt.Sprintf(`
		data "ibm_database_tasks" "database_tasks" {
			deployment_id = "%[1]s"
		}
	`, acc.IcdDbDeploymentId)
}
