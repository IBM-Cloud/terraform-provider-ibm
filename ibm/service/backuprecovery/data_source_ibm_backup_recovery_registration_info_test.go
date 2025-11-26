// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.105.1-067d600b-20250616-154447
 */

package backuprecovery_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIbmBackupRecoveryRegistrationInfoDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmBackupRecoveryRegistrationInfoDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_registration_info.baas_recovery_registration_info", "x_ibm_tenant_id"),
				),
			},
		},
	})
}

func testAccCheckIbmBackupRecoveryRegistrationInfoDataSourceConfigBasic() string {
	return fmt.Sprintf(`
	 data "ibm_backup_recovery_registration_info" "baas_recovery_registration_info" {
		     x_ibm_tenant_id = "wkk1yqrdce/"
		 }
		
	`)
}
