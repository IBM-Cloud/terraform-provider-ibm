// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package database_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

// TestAccIBMDatabaseTasksGen2DataSourceBasic validates the Gen2 tasks datasource
// Similar to classic TestAccIBMDatabaseTasksDataSourceBasic but for Gen2 databases
func TestAccIBMDatabaseTasksGen2DataSourceBasic(t *testing.T) {
	testName := fmt.Sprintf("tf-Pgress-gen2-%s", acctest.RandString(16))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckEnterprise(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMDatabaseTasksGen2DataSourceConfigBasic(testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_database_tasks.database_tasks", "deployment_id"),
					resource.TestCheckResourceAttrSet("data.ibm_database_tasks.database_tasks", "tasks.#"),
					// Gen2 returns at least 1 task (representing instance state)
					resource.TestCheckResourceAttr("data.ibm_database_tasks.database_tasks", "tasks.#", "1"),
					// Gen2 does not have task_id from last_operation, so it should be empty
					resource.TestCheckResourceAttr("data.ibm_database_tasks.database_tasks", "tasks.0.task_id", ""),
					resource.TestCheckResourceAttrSet("data.ibm_database_tasks.database_tasks", "tasks.0.deployment_id"),
					resource.TestCheckResourceAttrSet("data.ibm_database_tasks.database_tasks", "tasks.0.description"),
					resource.TestCheckResourceAttrSet("data.ibm_database_tasks.database_tasks", "tasks.0.status"),
					resource.TestCheckResourceAttrSet("data.ibm_database_tasks.database_tasks", "tasks.0.progress_percent"),
					resource.TestCheckResourceAttrSet("data.ibm_database_tasks.database_tasks", "tasks.0.created_at"),
				),
			},
		},
	})
}

// testAccCheckIBMDatabaseDataSourceConfigGen2ForTasks creates a Gen2 PostgreSQL instance for tasks testing.
func testAccCheckIBMDatabaseDataSourceConfigGen2ForTasks(name string) string {
	return fmt.Sprintf(`
data "ibm_resource_group" "test_acc" {
  is_default = true
}

resource "ibm_database" "db" {
  resource_group_id = data.ibm_resource_group.test_acc.id
  name              = "%[1]s"
  service           = "databases-for-postgresql"
  plan              = "standard-gen2"
  location          = "ca-mon"
  tags              = ["one:two"]
  service_endpoints = "private"

  timeouts {
    create = "60m"
  }

  group {
    group_id = "member"
    members {
      allocation_count = 2
    }
    host_flavor {
      id = "bx3d.4x20"
    }
    disk {
      allocation_mb = 10240
    }
  }
}
`, name)
}

func testAccCheckIBMDatabaseTasksGen2DataSourceConfigBasic(name string) string {
	return testAccCheckIBMDatabaseDataSourceConfigGen2ForTasks(name) + `
data "ibm_database_tasks" "database_tasks" {
  deployment_id = ibm_database.db.id
}
`
}

// Made with Bob
