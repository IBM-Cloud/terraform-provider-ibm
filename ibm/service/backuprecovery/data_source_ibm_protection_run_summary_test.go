// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package backuprecovery_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmProtectionRunSummaryDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmProtectionRunSummaryDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_protection_run_summary.protection_run_summary_instance", "id"),
				),
			},
		},
	})
}

func testAccCheckIbmProtectionRunSummaryDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_protection_run_summary" "protection_run_summary_instance" {
			startTimeUsecs = 1
			endTimeUsecs = 1
			runStatus = [ "Accepted" ]
		}
	`)
}
