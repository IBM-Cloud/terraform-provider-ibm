// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package database_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

// TestAccIBMDatabaseBackupGen2DataSourceBasic validates the Gen2 datasource
// using the same single-test-step acceptance style as the legacy datasource test.
func TestAccIBMDatabaseBackupGen2DataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMDatabaseBackupGen2DataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_database_backup.database_backup", "backup_id"),
					resource.TestCheckResourceAttrSet("data.ibm_database_backup.database_backup", "created_at"),
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
	`, acc.IcdDbBackupId)
}
