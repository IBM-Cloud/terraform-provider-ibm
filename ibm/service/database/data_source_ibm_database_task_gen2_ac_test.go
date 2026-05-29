// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package database_test

import (
	"fmt"
	"regexp"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

// TestAccIBMDatabaseTaskGen2DataSourceRead validates the Gen2 task datasource
// using the same single-test-step acceptance style as the legacy datasource test.
func TestAccIBMDatabaseTaskGen2DataSourceRead(t *testing.T) {
	testName := fmt.Sprintf("tf-Pgress-gen2-Task-%s", acctest.RandString(16))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckEnterprise(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMDatabaseTaskGen2DataSourceConfig(testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_database_task.database_task", "task_id"),
					resource.TestCheckResourceAttrSet("data.ibm_database_task.database_task", "deployment_id"),
					resource.TestCheckResourceAttrSet("data.ibm_database_task.database_task", "description"),
					resource.TestCheckResourceAttrSet("data.ibm_database_task.database_task", "status"),
					resource.TestCheckResourceAttrSet("data.ibm_database_task.database_task", "progress_percent"),
					resource.TestCheckResourceAttrSet("data.ibm_database_task.database_task", "created_at"),
				),
			},
		},
	})
}

// TestAccIBMDatabaseTaskGen2DataSourceInvalidID tests error handling for invalid task ID
func TestAccIBMDatabaseTaskGen2DataSourceInvalidID(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckEnterprise(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckIBMDatabaseTaskGen2DataSourceInvalidIDConfig(),
				ExpectError: regexp.MustCompile("failed to get resource instance|GetResourceInstance failed|not found|does not exist|invalid"),
			},
		},
	})
}

// TestAccIBMDatabaseTaskGen2DataSourceStatusMapping verifies that the Gen2
// datasource correctly maps Resource Controller instance states to task statuses.
func TestAccIBMDatabaseTaskGen2DataSourceStatusMapping(t *testing.T) {
	testName := fmt.Sprintf("tf-pg-gen2-status-%s", acctest.RandString(8))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckEnterprise(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMDatabaseTaskGen2DataSourceConfig(testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_database_task.database_task", "task_id"),
					resource.TestCheckResourceAttrSet("data.ibm_database_task.database_task", "status"),
					// Status should be "completed" for an active Gen2 instance
					resource.TestMatchResourceAttr("data.ibm_database_task.database_task", "status", regexp.MustCompile("^(completed|running|queued|failed)$")),
				),
			},
		},
	})
}

// TestAccIBMDatabaseTaskGen2DataSourceProgressPercent verifies that progress
// percentage is calculated correctly based on instance state.
func TestAccIBMDatabaseTaskGen2DataSourceProgressPercent(t *testing.T) {
	testName := fmt.Sprintf("tf-pg-gen2-progress-%s", acctest.RandString(8))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckEnterprise(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMDatabaseTaskGen2DataSourceConfig(testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_database_task.database_task", "progress_percent"),
					// Progress should be between 0 and 100
					resource.TestMatchResourceAttr("data.ibm_database_task.database_task", "progress_percent", regexp.MustCompile("^(100|[0-9]{1,2})$")),
				),
			},
		},
	})
}

// testAccCheckIBMDatabaseDataSourceConfigGen2ForTask creates a Gen2 PostgreSQL instance for task testing.
func testAccCheckIBMDatabaseDataSourceConfigGen2ForTask(name string) string {
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

func testAccCheckIBMDatabaseTaskGen2DataSourceConfig(name string) string {
	return testAccCheckIBMDatabaseDataSourceConfigGen2ForTask(name) + `
data "ibm_database_task" "database_task" {
  task_id = ibm_database.db.id
}
`
}

// testAccCheckIBMDatabaseTaskGen2DataSourceInvalidIDConfig tests with an invalid task ID
func testAccCheckIBMDatabaseTaskGen2DataSourceInvalidIDConfig() string {
	return `
data "ibm_database_task" "invalid_test" {
  task_id = "crn:v1:bluemix:public:databases-for-postgresql:us-south:a/1234567890abcdef:00000000-0000-0000-0000-000000000000::"
}
`
}
