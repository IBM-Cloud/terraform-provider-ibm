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

func TestAccIBMDatabaseTaskDataSourceBasic(t *testing.T) {
	testName := fmt.Sprintf("tf-Pgress-%s", acctest.RandString(16))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMDatabaseTaskDataSourceConfigBasic(testName),
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

func testAccCheckIBMDatabaseTaskDataSourceConfigBasic(name string) string {
	return fmt.Sprintf(`
		data "ibm_database_task" "database_task" {
			task_id = "crn:v1:bluemix:public:databases-for-redis:au-syd:a/40ddc34a953a8c02f10987b59085b60e:5042afe1-72c2-4231-89cc-c949e5d56251:task:72cda75b-bc2f-4e84-abb0-96f63dbb02b0"
		}
	`)
}
