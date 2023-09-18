// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package scc_test

import (
	"fmt"
	"os"
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
					resource.TestCheckResourceAttrSet("data.ibm_scc_report_summary.scc_report_summary_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_report_summary.scc_report_summary_instance", "report_id"),
				),
			},
		},
	})
}

func testAccCheckIbmSccReportSummaryDataSourceConfigBasic() string {
	report_id := os.Getenv("IBMCLOUD_SCC_REPORT_ID")
	return fmt.Sprintf(`
		data "ibm_scc_report_summary" "scc_report_summary_instance" {
			report_id = "%s"
		}
	`, report_id)
}
