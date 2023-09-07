// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package scc_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmSccReportEvaluationsDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSccReportEvaluationsDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_scc_report_evaluations.scc_report_evaluations", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_report_evaluations.scc_report_evaluations", "report_id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_report_evaluations.scc_report_evaluations", "first.#"),
				),
			},
		},
	})
}

func testAccCheckIbmSccReportEvaluationsDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_scc_report_evaluations" "scc_report_evaluations_instance" {
			report_id = "report_id"
			assessment_id = "assessment_id"
			component_id = "component_id"
			target_id = "target_id"
			target_name = "target_name"
			status = "failure"
		}
	`)
}
