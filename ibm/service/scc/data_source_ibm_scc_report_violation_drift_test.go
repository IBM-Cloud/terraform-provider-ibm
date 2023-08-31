// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package scc_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmSccReportViolationDriftDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSccReportViolationDriftDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_scc_report_violation_drift.scc_report_violation_drift", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_report_violation_drift.scc_report_violation_drift", "report_id"),
				),
			},
		},
	})
}

func testAccCheckIbmSccReportViolationDriftDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_scc_report_violation_drift" "scc_report_violation_drift_instance" {
			report_id = "report_id"
			X-Correlation-Id = "X-Correlation-Id"
			X-Request-Id = "X-Request-Id"
			scan_time_duration = 0
		}
	`)
}
