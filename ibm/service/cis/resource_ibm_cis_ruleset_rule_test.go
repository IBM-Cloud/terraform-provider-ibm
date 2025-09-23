// Copyright IBM Corp. 2017, 2025. All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cis_test

import (
	"fmt"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMCISRulesetRule_Basic(t *testing.T) {
	name := "ibm_cis_ruleset_rule.test"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckCis(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCisRulesetsRule_basic("test", acc.CisDomainStatic),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(name, "ruleset_id", "eb5efc50d5ec49d8b0b0f44b357b8d7b"),
					resource.TestCheckResourceAttr(name, "rule.#", "1"),
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
		ruleset_id = "eb5efc50d5ec49d8b0b0f44b357b8d7b"
			rule {
			  action =  "block"
			  description = "Testing rule creation"
			  enabled = true
			  expression = "true"
			}
	  }
`, id, acc.CisDomainStatic)
}
