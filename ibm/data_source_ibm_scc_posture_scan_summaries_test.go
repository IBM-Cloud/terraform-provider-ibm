// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMSccPostureScanSummariesDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMSccPostureScanSummariesDataSourceConfigBasic(scc_posture_profile_id, scc_posture_scope_id),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_scc_posture_scan_summaries.scan_summaries", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_posture_scan_summaries.scan_summaries", "profile_id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_posture_scan_summaries.scan_summaries", "scope_id"),
				),
			},
		},
	})
}

func testAccCheckIBMSccPostureScanSummariesDataSourceConfigBasic(profileId string, scopeId string) string {
	return fmt.Sprintf(`
		data "ibm_scc_posture_scan_summaries" "scan_summaries" {
			profile_id = "%s"
			scope_id = "%s"
		}
	`, profileId, scopeId)
}
