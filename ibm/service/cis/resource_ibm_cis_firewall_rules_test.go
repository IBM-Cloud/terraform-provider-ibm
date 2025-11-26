// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package cis_test

import (
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIBMCisFirewallrules_Basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheckCis(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckCisFirewallrules_basic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_cis_firewall_rule.firewall_rules_instance", "action", "allow"),
					resource.TestCheckResourceAttr("ibm_cis_firewall_rule.firewall_rules_instance", "priority", "5"),
					resource.TestCheckResourceAttr("ibm_cis_firewall_rule.firewall_rules_instance", "paused", "true"),
				),
			},
		},
	})
}
func testAccCheckCisFirewallrules_basic() string {
	return testAccCheckIBMCisDomainDataSourceConfigBasic1() + `
	resource "ibm_cis_filter" "test" {
		cis_id =  data.ibm_cis.cis.id
		domain_id = data.ibm_cis_domain.cis_domain.domain_id
		expression = "(ip.src eq 156.25.53.188 and http.request.uri.path eq \"^.*/wp-login[0-9].php$\")"
		paused =  true
		description = "Filter-creation"
	}
	resource "ibm_cis_firewall_rule" "firewall_rules_instance" {
		depends_on  = [ibm_cis_filter.test]
		cis_id =  data.ibm_cis.cis.id
		domain_id = data.ibm_cis_domain.cis_domain.domain_id
  		filter_id = ibm_cis_filter.test.filter_id
  		action = "allow"
  		priority = 5
		paused = true
	  }
`
}
