// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cis_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMCISRulesetRule_Basic(t *testing.T) {
	name := "data.ibm_cis_ruleset_rule.test"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckCis(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCisAlertsDataSource_basic("test", acc.CisDomainStatic),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "ruleset_id", "dcdec3fe0cbe41edac08619503da8de5"),
				),
			},
		},
	})
}

func testAccCheckCisRulesetsRule_basic(id, CisDomainStatic string) string {
	return testAccCheckIBMCisDomainDataSourceConfigBasic1() + fmt.Sprintf(`
	resource "ibm_cis_ruleset_rule" "%[1]s" {
		cis_id = data.ibm_cis.cis.id
		domain_id = data.ibm_cis_domain.cis_domain.domain_id
		ruleset_id = "dcdec3fe0cbe41edac08619503da8de5"
		rules {
			{
			  action =  "block"
			  description = "Manually created rule"
			  enabled = true
			}
		}
	  }
`, id, acc.CisDomainStatic)
}
