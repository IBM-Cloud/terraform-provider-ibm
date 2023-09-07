// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package scc_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmSccReportSummaryDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSccReportSummaryDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_scc_report_summary.scc_report_summary", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_report_summary.scc_report_summary", "report_id"),
				),
			},
		},
	})
}

func testAccCheckIbmSccReportSummaryDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_scc_report_summary" "scc_report_summary_instance" {
			report_id = "report_id"
		}
	`)
}
