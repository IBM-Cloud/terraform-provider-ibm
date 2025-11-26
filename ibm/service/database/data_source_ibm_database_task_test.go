// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package database_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIBMDatabaseTaskDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMDatabaseTaskDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_database_task.database_task", "task_id"),
					resource.TestCheckResourceAttrSet("data.ibm_database_task.database_task", "description"),
					resource.TestCheckResourceAttrSet("data.ibm_database_task.database_task", "status"),
					resource.TestCheckResourceAttrSet("data.ibm_database_task.database_task", "deployment_id"),
					resource.TestCheckResourceAttrSet("data.ibm_database_task.database_task", "progress_percent"),
					resource.TestCheckResourceAttrSet("data.ibm_database_task.database_task", "created_at"),
				),
			},
		},
	})
}

func testAccCheckIBMDatabaseTaskDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_database_task" "database_task" {
			task_id = "%[1]s"
		}
	`, acc.IcdDbTaskId)
}
