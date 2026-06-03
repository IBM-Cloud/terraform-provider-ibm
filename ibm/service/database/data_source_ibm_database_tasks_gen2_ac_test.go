// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package database_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

// TestAccIBMDatabaseTasksGen2DataSourceBasic validates the Gen2 tasks datasource
// using environment variable for existing Gen2 deployment (same pattern as classic test)
func TestAccIBMDatabaseTasksGen2DataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckEnterprise(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMDatabaseTasksGen2DataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_database_tasks.database_tasks", "deployment_id"),
					resource.TestCheckResourceAttrSet("data.ibm_database_tasks.database_tasks", "tasks.#"),
				),
			},
		},
	})
}

func testAccCheckIBMDatabaseTasksGen2DataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_database_tasks" "database_tasks" {
			deployment_id = "%[1]s"
		}
	`, acc.IcdDbGen2DeploymentId)
}
