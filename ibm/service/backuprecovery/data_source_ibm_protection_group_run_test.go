// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package backuprecovery_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmProtectionGroupRunDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmProtectionGroupRunDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_protection_group_run.protection_group_run_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_protection_group_run.protection_group_run_instance", "protection_group_run_id"),
					resource.TestCheckResourceAttrSet("data.ibm_protection_group_run.protection_group_run_instance", "run_id"),
				),
			},
		},
	})
}

func testAccCheckIbmProtectionGroupRunDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_protection_group_run" "protection_group_run_instance" {
			id = "id"
			runId = "runId"
			requestInitiatorType = "UIUser"
			tenantIds = [ "tenantIds" ]
			includeTenants = true
			includeObjectDetails = true
			useCachedData = true
		}
	`)
}
