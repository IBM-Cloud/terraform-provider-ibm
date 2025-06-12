// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package scc_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmSccReportDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckScc(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmSccReportDataSourceConfigBasic(acc.SccInstanceID, acc.SccReportID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_scc_report.scc_report_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_report.scc_report_instance", "report_id"),
				),
			},
		},
	})
}

func testAccCheckIbmSccReportDataSourceConfigBasic(instanceID, reportID string) string {
	return fmt.Sprintf(`
		data "ibm_scc_report" "scc_report_instance" {
			instance_id = "%s"
			report_id = "%s"
		}
	`, instanceID, reportID)
}
