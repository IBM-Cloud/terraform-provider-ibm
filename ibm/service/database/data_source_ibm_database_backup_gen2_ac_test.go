// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package database_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

// TestAccIBMDatabaseBackupGen2DataSourceBasic validates the Gen2 datasource
// using the same single-test-step acceptance style as the legacy datasource test.
func TestAccIBMDatabaseBackupGen2DataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckEnterprise(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMDatabaseBackupGen2DataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_database_backup.database_backup", "backup_id"),
					resource.TestCheckResourceAttrSet("data.ibm_database_backup.database_backup", "deployment_id"),
					resource.TestCheckResourceAttrSet("data.ibm_database_backup.database_backup", "type"),
					resource.TestCheckResourceAttrSet("data.ibm_database_backup.database_backup", "status"),
					resource.TestCheckResourceAttrSet("data.ibm_database_backup.database_backup", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_database_backup.database_backup", "is_downloadable"),
					resource.TestCheckResourceAttrSet("data.ibm_database_backup.database_backup", "is_restorable"),
					resource.TestCheckResourceAttr("data.ibm_database_backup.database_backup", "download_link", ""),
				),
			},
		},
	})
}

func testAccCheckIBMDatabaseBackupGen2DataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_database_backup" "database_backup" {
			backup_id = "%[1]s"
		}
	`, acc.Gen2BackupId)
}

// TestAccIBMDatabaseBackupGen2DataSourceInvalidID verifies the error path when
// the backup_id does not correspond to a real Independent Backup instance.
func TestAccIBMDatabaseBackupGen2DataSourceInvalidID(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckEnterprise(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckIBMDatabaseBackupGen2DataSourceInvalidIDConfig(),
				ExpectError: regexp.MustCompile("Independent Backup not found|GetResourceInstance failed|not found"),
			},
		},
	})
}

func testAccCheckIBMDatabaseBackupGen2DataSourceInvalidIDConfig() string {
	return `
	data "ibm_database_backup" "invalid_backup" {
		backup_id = "crn:v1:bluemix:public:databases-independent-backups:us-south:a/00000000000000000000000000000000:00000000-0000-0000-0000-000000000000::"
	}
	`
}

// TestAccIBMDatabaseBackupGen2DataSourceS2SWarning verifies that reading a Gen2
// Independent Backup for an instance that has Independent Backups but no S2S
// authorizations configured emits a warning and still succeeds (non-blocking).
func TestAccIBMDatabaseBackupGen2DataSourceS2SWarning(t *testing.T) {
	t.Parallel()
	name := fmt.Sprintf("tf-gen2-s2s-backup-%s", acctest.RandString(8))
	dsName := "data.ibm_database_backup.gen2_s2s_backup"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckEnterprise(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMDatabaseBackupGen2S2SConfig(name),
				Check: resource.ComposeTestCheckFunc(
					// backup fields populated despite the S2S warning
					resource.TestCheckResourceAttrSet(dsName, "backup_id"),
					resource.TestCheckResourceAttrSet(dsName, "deployment_id"),
					resource.TestCheckResourceAttrSet(dsName, "type"),
					resource.TestCheckResourceAttrSet(dsName, "status"),
					resource.TestCheckResourceAttrSet(dsName, "created_at"),
					resource.TestCheckResourceAttr(dsName, "is_downloadable", "false"),
					resource.TestCheckResourceAttr(dsName, "download_link", ""),
				),
			},
		},
	})
}

// testAccCheckIBMDatabaseBackupGen2S2SConfig creates a Gen2 MySQL instance in
// eu-fr2 without S2S authorizations, waits for an Independent Backup to exist,
// then reads it via the ibm_database_backup datasource. The S2S warning fires
// but must not prevent the backup attributes from being populated.
func testAccCheckIBMDatabaseBackupGen2S2SConfig(name string) string {
	return fmt.Sprintf(`
data "ibm_resource_group" "test_acc" {
  is_default = true
}

resource "ibm_database" "gen2_s2s_backup" {
  resource_group_id = data.ibm_resource_group.test_acc.id
  name              = %[1]q
  service           = "databases-for-mysql"
  plan              = "standard-gen2"
  location          = "eu-fr2"
  tags              = ["terraform", "s2s-backup-test"]

  group {
    group_id = "member"
    members {
      allocation_count = 2
    }
  }

  timeouts {
    create = "120m"
    update = "60m"
    delete = "15m"
  }
}

data "ibm_database_backups" "gen2_s2s_list" {
  deployment_id = ibm_database.gen2_s2s_backup.id

  depends_on = [ibm_database.gen2_s2s_backup]
}

data "ibm_database_backup" "gen2_s2s_backup" {
  backup_id = data.ibm_database_backups.gen2_s2s_list.backups[0].backup_id

  depends_on = [data.ibm_database_backups.gen2_s2s_list]
}
`, name)
}
