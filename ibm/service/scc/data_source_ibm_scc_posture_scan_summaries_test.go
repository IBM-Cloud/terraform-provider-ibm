// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package scc_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMSccPostureScanSummariesDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMSccPostureScanSummariesDataSourceConfigBasic(acc.Scc_posture_report_setting_id),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_scc_posture_scan_summaries.scan_summaries", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_posture_scan_summaries.scan_summaries", "report_setting_id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_posture_scan_summaries.scan_summaries", "summaries.#"),
				),
			},
		},
	})
}

func testAccCheckIBMSccPostureScanSummariesDataSourceConfigBasic(report_setting_id string) string {
	return fmt.Sprintf(`
		data "ibm_scc_posture_scan_summaries" "scan_summaries" {
			report_setting_id = "%s"
		}
	`, report_setting_id)
}
