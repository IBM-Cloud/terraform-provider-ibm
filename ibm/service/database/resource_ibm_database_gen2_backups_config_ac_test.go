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

// TestAccIBMDatabaseGen2BackupsConfig provisions a Gen2 PostgreSQL instance with
// a backups block and verifies the backups attributes are read back into state.
func TestAccIBMDatabaseGen2BackupsConfig(t *testing.T) {
	name := fmt.Sprintf("tf-gen2-backups-%s", acctest.RandString(8))
	rName := "ibm_database.gen2_backups"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckEnterprise(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccIBMDatabaseGen2BackupsConfigBasic(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(rName, "id"),
					resource.TestCheckResourceAttr(rName, "backups.#", "1"),
					resource.TestCheckResourceAttr(rName, "backups.0.preserve", "false"),
					resource.TestCheckResourceAttr(rName, "backups.0.automatic_backups.#", "1"),
					resource.TestCheckResourceAttr(rName, "backups.0.automatic_backups.0.copy.#", "1"),
					resource.TestCheckResourceAttr(rName, "backups.0.automatic_backups.0.copy.0.enabled", "true"),
				),
			},
		},
	})
}

// TestAccIBMDatabaseGen2BackupsPreserve verifies preserve=true is accepted and
// reflected in state.
func TestAccIBMDatabaseGen2BackupsPreserve(t *testing.T) {
	name := fmt.Sprintf("tf-gen2-backups-preserve-%s", acctest.RandString(8))
	rName := "ibm_database.gen2_backups_preserve"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckEnterprise(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccIBMDatabaseGen2BackupsConfigPreserve(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(rName, "id"),
					resource.TestCheckResourceAttr(rName, "backups.#", "1"),
					resource.TestCheckResourceAttr(rName, "backups.0.preserve", "true"),
				),
			},
		},
	})
}

// TestAccIBMDatabaseGen2BackupsCopyDisabled verifies copy.enabled=false is accepted.
func TestAccIBMDatabaseGen2BackupsCopyDisabled(t *testing.T) {
	name := fmt.Sprintf("tf-gen2-backups-nocopy-%s", acctest.RandString(8))
	rName := "ibm_database.gen2_backups_nocopy"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckEnterprise(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccIBMDatabaseGen2BackupsConfigCopyDisabled(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(rName, "id"),
					resource.TestCheckResourceAttr(rName, "backups.0.automatic_backups.0.copy.0.enabled", "false"),
				),
			},
		},
	})
}

// testAccIBMDatabaseGen2BackupsConfigBasic creates a Gen2 PostgreSQL instance with copy enabled.
// Uses in-che which supports automatic backup copy.
func testAccIBMDatabaseGen2BackupsConfigBasic(name string) string {
	return fmt.Sprintf(`
data "ibm_resource_group" "test_acc" {
  is_default = true
}

resource "ibm_database" "gen2_backups" {
  resource_group_id = data.ibm_resource_group.test_acc.id
  name              = %[1]q
  service           = "databases-for-postgresql"
  plan              = "standard-gen2"
  location          = "in-che"
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

  backups {
    preserve = false

    automatic_backups {
      copy {
        enabled = true
      }
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

// testAccIBMDatabaseGen2BackupsConfigPreserve creates an instance with preserve=true.
func testAccIBMDatabaseGen2BackupsConfigPreserve(name string) string {
	return fmt.Sprintf(`
data "ibm_resource_group" "test_acc" {
  is_default = true
}

resource "ibm_database" "gen2_backups_preserve" {
  resource_group_id = data.ibm_resource_group.test_acc.id
  name              = %[1]q
  service           = "databases-for-postgresql"
  plan              = "standard-gen2"
  location          = "in-che"
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

  backups {
    preserve = true
  }

  timeouts {
    create = "120m"
    update = "60m"
    delete = "15m"
  }
}
`, name)
}

// testAccIBMDatabaseGen2BackupsConfigCopyDisabled creates an instance with copy disabled.
func testAccIBMDatabaseGen2BackupsConfigCopyDisabled(name string) string {
	return fmt.Sprintf(`
data "ibm_resource_group" "test_acc" {
  is_default = true
}

resource "ibm_database" "gen2_backups_nocopy" {
  resource_group_id = data.ibm_resource_group.test_acc.id
  name              = %[1]q
  service           = "databases-for-postgresql"
  plan              = "standard-gen2"
  location          = "in-che"
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

  backups {
    automatic_backups {
      copy {
        enabled = false
      }
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
