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

func TestAccIbmSccReportRuleDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmSccReportRuleDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_scc_report_rule.scc_report_rule_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_report_rule.scc_report_rule_instance", "report_id"),
					resource.TestCheckResourceAttrSet("data.ibm_scc_report_rule.scc_report_rule_instance", "rule_id"),
				),
			},
		},
	})
}

func testAccCheckIbmSccReportRuleDataSourceConfigBasic() string {
	report_id := os.Getenv("IBMCLOUD_SCC_REPORT_ID")
	return fmt.Sprintf(`
		data "ibm_scc_report_rule" "scc_report_rule_instance" {
			report_id = "%s"
			rule_id = "rule-f8722625-1968-4d7a-93cb-4b0f8da726da"
		}
	`, report_id)
}
