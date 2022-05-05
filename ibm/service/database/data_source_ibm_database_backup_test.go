// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package clouddatabases_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMDatabaseBackupDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMDatabaseBackupDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_database_backup.database_backup", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_database_backup.database_backup", "backup_id"),
				),
			},
		},
	})
}

func testAccCheckIBMDatabaseBackupDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_database_backup" "database_backup" {
			backup_id = "backup_id"
		}
	`)
}

