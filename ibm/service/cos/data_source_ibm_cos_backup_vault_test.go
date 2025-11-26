// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cos_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIBMCOSBackupVaultDataSource_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckCOS(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccIBMCOSBackupVaultDataSourceConfig_basic_read(acc.BackupVaultName, acc.CosCRN, acc.RegionName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_cos_backup_vault.vault", "backup_vault_name"),
					resource.TestCheckResourceAttrSet("data.ibm_cos_backup_vault.vault", "service_instance_id"),
					resource.TestCheckResourceAttrSet("data.ibm_cos_backup_vault.vault", "region"),
				),
			},
		},
	})
}

func testAccIBMCOSBackupVaultDataSourceConfig_basic_read(name string, crn string, region string) string {
	return fmt.Sprintf(`

		data "ibm_cos_backup_vault" "vault" {
			backup_vault_name          = "%s"
			service_instance_id = "%s"
			region = "%s"
		}`, name, crn, region)
}
