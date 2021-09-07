// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMScanSummariesDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMScanSummariesDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_scan_summaries.scan_summaries", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_scan_summaries.scan_summaries", "profile_id"),
					resource.TestCheckResourceAttrSet("data.ibm_scan_summaries.scan_summaries", "scope_id"),
				),
			},
		},
	})
}

func testAccCheckIBMScanSummariesDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_scan_summaries" "scan_summaries" {
			profile_id = "profile_id"
			scope_id = "scope_id"
			scan_id = "262"
		}
	`)
}

