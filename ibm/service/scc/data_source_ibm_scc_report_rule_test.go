// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package scc_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmSccReportRuleDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSccReportRuleDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_scc_report_rule.scc_report_rule", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_report_rule.scc_report_rule", "report_id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_report_rule.scc_report_rule", "rule_id"),
				),
			},
		},
	})
}

func testAccCheckIbmSccReportRuleDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_scc_report_rule" "scc_report_rule_instance" {
			report_id = "report_id"
			rule_id = "rule-8d444f8c-fd1d-48de-bcaa-f43732568761"
		}
	`)
}
