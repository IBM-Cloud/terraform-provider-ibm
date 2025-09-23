// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cis_test

import (
	"fmt"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMCisFirewallRulesDataSource_Basic(t *testing.T) {
	name := "data.ibm_cis_firewall_rules.test"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmFirewallRulesDataSourceConfigBasic("test", acc.CisDomainStatic),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(name, "id"),
				),
			},
		},
	})
}

func testAccCheckIbmFirewallRulesDataSourceConfigBasic(id, CisDomainStatic string) string {
	return testAccCheckIBMCisDomainDataSourceConfigBasic1() + fmt.Sprintf(`
	data "ibm_cis_firewall_rules" "%[1]s" {
		cis_id          = data.ibm_cis.cis.id
		domain_id       = data.ibm_cis_domain.cis_domain.domain_id
	  }
`, id, acc.CisDomainStatic)
}
