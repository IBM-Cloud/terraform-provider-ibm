// Copyright IBM Corp. 2017, 2026 All Rights Reserved.
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

// TestAccIBMDatabaseGen2BackupConfiguration provisions a Gen2 instance with backup_configuration,
// updates it in place, and verifies the values read back through the resource and the data source.
//
// preserve is intentionally left at its default (false) here: the ADR makes preserve=true
// irreversible and it would leave independent backups behind after the instance is destroyed.
// The preserve rules are covered by the plan-time tests below and by unit tests.
//
// target_regions uses ca-tor as a valid cross-region copy target for a ca-mon source; adjust if the
// available copy regions change.
func TestAccIBMDatabaseGen2BackupConfiguration(t *testing.T) {
	t.Parallel()
	var databaseInstanceOne string
	testName := fmt.Sprintf("tf-gen2-backupcfg-%d", acctest.RandIntRange(10, 100))
	name := "ibm_database." + testName
	dataName := "data.ibm_database.lookup"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheckEnterprise(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMDatabaseInstanceDestroy,
		Steps: []resource.TestStep{
			// Step 1 - create with retention + a cross-region copy target
			{
				Config: testAccCheckIBMDatabaseGen2BackupConfiguration(testName, 45, true, `["ca-tor"]`),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMDatabaseInstanceExists(name, &databaseInstanceOne),
					resource.TestCheckResourceAttr(name, "plan", "standard-gen2"),
					resource.TestCheckResourceAttr(name, "backup_configuration.#", "1"),
					resource.TestCheckResourceAttr(name, "backup_configuration.0.retention", "45"),
					resource.TestCheckResourceAttr(name, "backup_configuration.0.preserve", "false"),
					resource.TestCheckResourceAttr(name, "backup_configuration.0.copy_enabled", "true"),
					resource.TestCheckResourceAttr(name, "backup_configuration.0.target_regions.#", "1"),
					resource.TestCheckTypeSetElemAttr(name, "backup_configuration.0.target_regions.*", "ca-tor"),
				),
			},
			// Step 2 - update in place: change retention and turn automatic cross-region copies off
			{
				Config: testAccCheckIBMDatabaseGen2BackupConfiguration(testName, 60, false, `["ca-tor"]`),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(name, "backup_configuration.0.retention", "60"),
					resource.TestCheckResourceAttr(name, "backup_configuration.0.copy_enabled", "false"),
					resource.TestCheckResourceAttr(name, "backup_configuration.0.preserve", "false"),
				),
			},
			// Step 3 - the data source reads the same configuration back
			{
				Config: testAccCheckIBMDatabaseGen2BackupConfiguration(testName, 60, false, `["ca-tor"]`) +
					fmt.Sprintf(`
data "ibm_database" "lookup" {
  resource_group_id = data.ibm_resource_group.test_acc.id
  name              = ibm_database.%[1]s.name
  service           = "databases-for-postgresql"
  location          = "ca-mon"
}
`, testName),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr(dataName, "backup_configuration.#", "1"),
					resource.TestCheckResourceAttr(dataName, "backup_configuration.0.retention", "60"),
					resource.TestCheckResourceAttr(dataName, "backup_configuration.0.copy_enabled", "false"),
					resource.TestCheckResourceAttr(dataName, "backup_configuration.0.preserve", "false"),
				),
			},
			// Step 4 - re-planning the same config yields no diff (no perpetual diff from the set/defaults)
			{
				Config:   testAccCheckIBMDatabaseGen2BackupConfiguration(testName, 60, false, `["ca-tor"]`),
				PlanOnly: true,
			},
		},
	})
}

// TestAccIBMDatabaseGen2BackupConfigurationRetentionOutOfRange verifies the 30-60 day range is
// enforced at plan time, before any API call. No resource is created.
func TestAccIBMDatabaseGen2BackupConfigurationRetentionOutOfRange(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckEnterprise(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckIBMDatabaseGen2BackupConfiguration("tf-backupcfg-low", 29, true, `[]`),
				ExpectError: regexp.MustCompile(`retention in backup_configuration must be between 30 and 60 days`),
			},
			{
				Config:      testAccCheckIBMDatabaseGen2BackupConfiguration("tf-backupcfg-high", 61, true, `[]`),
				ExpectError: regexp.MustCompile(`retention in backup_configuration must be between 30 and 60 days`),
			},
		},
	})
}

// TestAccIBMDatabaseGen2BackupConfigurationMySQLPreserveRejected verifies preserve=true is rejected
// for Databases for MySQL at plan time. No resource is created.
func TestAccIBMDatabaseGen2BackupConfigurationMySQLPreserveRejected(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckEnterprise(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckIBMDatabaseBackupConfigurationMySQLPreserve("tf-backupcfg-mysql"),
				ExpectError: regexp.MustCompile(`preserve backups is not supported for Databases for MySQL`),
			},
		},
	})
}

// TestAccIBMDatabaseClassicBackupConfigurationRejected verifies a Classic plan rejects
// backup_configuration at plan time with the unsupported-attribute error. No resource is created.
func TestAccIBMDatabaseClassicBackupConfigurationRejected(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckEnterprise(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckIBMDatabaseClassicBackupConfiguration("tf-classic-backupcfg"),
				ExpectError: regexp.MustCompile(`(?s)not supported for Classic databases.*backup_configuration`),
			},
		},
	})
}

func testAccCheckIBMDatabaseGen2BackupConfiguration(name string, retention int, copyEnabled bool, targetRegions string) string {
	return fmt.Sprintf(`
data "ibm_resource_group" "test_acc" {
  is_default = true
}

resource "ibm_database" "%[1]s" {
  resource_group_id = data.ibm_resource_group.test_acc.id
  name              = "%[1]s"
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

  backup_configuration {
    retention      = %[2]d
    copy_enabled   = %[3]t
    target_regions = %[4]s
  }

  timeouts {
    create = "120m"
    update = "60m"
    delete = "15m"
  }
}
`, name, retention, copyEnabled, targetRegions)
}

func testAccCheckIBMDatabaseBackupConfigurationMySQLPreserve(name string) string {
	return fmt.Sprintf(`
data "ibm_resource_group" "test_acc" {
  is_default = true
}

resource "ibm_database" "%[1]s" {
  resource_group_id = data.ibm_resource_group.test_acc.id
  name              = "%[1]s"
  service           = "databases-for-mysql"
  plan              = "standard-gen2"
  location          = "eu-fr2"

  group {
    group_id = "member"
    members {
      allocation_count = 2
    }
  }

  backup_configuration {
    preserve = true
  }
}
`, name)
}

func testAccCheckIBMDatabaseClassicBackupConfiguration(name string) string {
	return fmt.Sprintf(`
data "ibm_resource_group" "test_acc" {
  is_default = true
}

resource "ibm_database" "%[1]s" {
  resource_group_id = data.ibm_resource_group.test_acc.id
  name              = "%[1]s"
  service           = "databases-for-postgresql"
  plan              = "standard"
  location          = "us-south"
  service_endpoints = "public"

  backup_configuration {
    retention = 30
  }
}
`, name)
}
