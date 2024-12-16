// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cis_test

import (
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIBMCisPageRuleDataSource_basic(t *testing.T) {
	node := "data.ibm_cis_page_rules.test"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMCisPageRuleDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(node, "cis_page_rules.0.id"),
					resource.TestCheckResourceAttrSet(node, "cis_page_rules.0.rule_id"),
				),
			},
		},
	})
}

func testAccCheckIBMCisPageRuleDataSourceConfig() string {
	// status filter defaults to empty
	return testAccCheckIBMCisPageRuleConfigBasic() + `
	data "ibm_cis_page_rules" "test" {
		cis_id     = ibm_cis_page_rule.page_rule.cis_id
		domain_id  = ibm_cis_page_rule.page_rule.domain_id
	  }`
}
