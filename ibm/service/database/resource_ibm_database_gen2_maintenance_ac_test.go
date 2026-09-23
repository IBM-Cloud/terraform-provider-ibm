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

// TestAccIBMDatabaseGen2MaintenanceWindowCustom provisions a Gen2 PostgreSQL
// instance with a custom maintenance window, verifies the window is stored and
// read back, then updates it to a different window in a second step.
func TestAccIBMDatabaseGen2MaintenanceWindowCustom(t *testing.T) {
	t.Parallel()
	name := fmt.Sprintf("tf-gen2-maint-%s", acctest.RandString(10))
	resourceName := "ibm_database." + name

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMDatabaseInstanceDestroy,
		Steps: []resource.TestStep{
			{
				// Step 1: create with a custom maintenance window
				Config: testAccIBMDatabaseGen2MaintenanceWindowCustomConfig(name, "05:00Z", "Wednesday", "Thursday"),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(resourceName, new(string)),
					resource.TestCheckResourceAttr(resourceName, "name", name),
					resource.TestCheckResourceAttr(resourceName, "plan", "standard-gen2"),
					resource.TestCheckResourceAttr(resourceName, "maintenance.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "maintenance.0.window.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "maintenance.0.window.0.start_time", "05:00Z"),
					resource.TestCheckTypeSetElemAttr(resourceName, "maintenance.0.window.0.days.*", "Wednesday"),
					resource.TestCheckTypeSetElemAttr(resourceName, "maintenance.0.window.0.days.*", "Thursday"),
					resource.TestCheckResourceAttr(resourceName, "maintenance.0.window.0.system_assigned", "false"),
				),
			},
			{
				// Step 2: update to a different window — single call to RC
				Config: testAccIBMDatabaseGen2MaintenanceWindowCustomConfig(name, "03:00Z", "Saturday", "Sunday"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "maintenance.0.window.0.start_time", "03:00Z"),
					resource.TestCheckTypeSetElemAttr(resourceName, "maintenance.0.window.0.days.*", "Saturday"),
					resource.TestCheckTypeSetElemAttr(resourceName, "maintenance.0.window.0.days.*", "Sunday"),
				),
			},
			{
				// Step 3: import and verify round-trip
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"wait_time_minutes", "deletion_protection",
				},
			},
		},
	})
}

// TestAccIBMDatabaseGen2MaintenanceWindowSystemAssigned provisions a Gen2
// instance with system_assigned=true (IBM-controlled schedule), then updates
// it to a custom window and back again.
func TestAccIBMDatabaseGen2MaintenanceWindowSystemAssigned(t *testing.T) {
	t.Parallel()
	name := fmt.Sprintf("tf-gen2-maint-sys-%s", acctest.RandString(10))
	resourceName := "ibm_database." + name

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMDatabaseInstanceDestroy,
		Steps: []resource.TestStep{
			{
				// Step 1: create with system_assigned
				Config: testAccIBMDatabaseGen2MaintenanceWindowSystemAssignedConfig(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(resourceName, new(string)),
					resource.TestCheckResourceAttr(resourceName, "maintenance.0.window.0.system_assigned", "true"),
				),
			},
			{
				// Step 2: switch to a custom window
				Config: testAccIBMDatabaseGen2MaintenanceWindowCustomConfig(name, "07:00Z", "Monday", "Tuesday"),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "maintenance.0.window.0.start_time", "07:00Z"),
					resource.TestCheckTypeSetElemAttr(resourceName, "maintenance.0.window.0.days.*", "Monday"),
					resource.TestCheckTypeSetElemAttr(resourceName, "maintenance.0.window.0.days.*", "Tuesday"),
					resource.TestCheckResourceAttr(resourceName, "maintenance.0.window.0.system_assigned", "false"),
				),
			},
			{
				// Step 3: reset back to system_assigned
				Config: testAccIBMDatabaseGen2MaintenanceWindowSystemAssignedConfig(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "maintenance.0.window.0.system_assigned", "true"),
				),
			},
		},
	})
}

// ---------------------------------------------------------------------------
// Terraform config helpers
// ---------------------------------------------------------------------------

// testAccIBMDatabaseGen2MaintenanceWindowCustomConfig creates a Gen2 PostgreSQL
// instance with a custom maintenance window (start_time + two days).
func testAccIBMDatabaseGen2MaintenanceWindowCustomConfig(name, startTime, day1, day2 string) string {
	return fmt.Sprintf(`
data "ibm_resource_group" "test_acc" {
  is_default = true
}

resource "ibm_database" %[1]q {
  resource_group_id = data.ibm_resource_group.test_acc.id
  name              = %[1]q
  service           = "databases-for-postgresql"
  plan              = "standard-gen2"
  location          = "ca-mon"
  service_endpoints = "private"

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

  maintenance {
    window {
      start_time = %[2]q
      days       = [%[3]q, %[4]q]
    }
  }

  timeouts {
    create = "120m"
    update = "60m"
    delete = "15m"
  }
}
`, name, startTime, day1, day2)
}

// testAccIBMDatabaseGen2MaintenanceWindowSystemAssignedConfig creates a Gen2
// PostgreSQL instance with system_assigned=true (IBM-controlled schedule).
func testAccIBMDatabaseGen2MaintenanceWindowSystemAssignedConfig(name string) string {
	return fmt.Sprintf(`
data "ibm_resource_group" "test_acc" {
  is_default = true
}

resource "ibm_database" %[1]q {
  resource_group_id = data.ibm_resource_group.test_acc.id
  name              = %[1]q
  service           = "databases-for-postgresql"
  plan              = "standard-gen2"
  location          = "ca-mon"
  service_endpoints = "private"

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

  maintenance {
    window {
      system_assigned = true
    }
  }

  timeouts {
    create = "120m"
    update = "60m"
    delete = "15m"
  }
}
`, name)
}

