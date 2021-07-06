// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIbmDnsCrForwardingRulesDataSourceBasic(t *testing.T) {
	forwardingRuleDescription := "test forward rule"
	forwardingRuleType := "zone"
	forwardingRuleMatch := "test.example.com"
	node := "data.ibm_dns_cr_forwarding_rules.dns_cr_forwarding_rules"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmDnsCrForwardingRulesDataSourceConfig(forwardingRuleDescription, forwardingRuleType, forwardingRuleMatch),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(node, "ibm_dns_cr_forwarding_rule.0.description"),
					resource.TestCheckResourceAttrSet(node, "ibm_dns_cr_forwarding_rule.0.type"),
					resource.TestCheckResourceAttrSet(node, "ibm_dns_cr_forwarding_rule.0.match"),
				),
			},
		},
	})
}

func testAccCheckIbmDnsCrForwardingRulesDataSourceConfig(forwardingRuleDescription string, forwardingRuleType string, forwardingRuleMatch string) string {
	return fmt.Sprintf(`
		resource "ibm_dns_cr_forwarding_rule" "dns_cr_forwarding_rule" {
			instance_id = "345ca2c4-83bf-4c04-bb09-5d8ec4d425a8"
			resolver_id = "db857778-2aa9-4d81-b954-08bcaba0dadf"
			description = "%s"
			type = "%s"
			match = "%s"
			forward_to = ["168.20.22.122"]
		}
		data "ibm_dns_cr_forwarding_rules" "dns_cr_forwarding_rules" {
			instance_id = ibm_dns_cr_forwarding_rule.dns_cr_forwarding_rule.instance_id
			resolver_id = ibm_dns_cr_forwarding_rule.dns_cr_forwarding_rule.resolver_id
		  }
	`, forwardingRuleDescription, forwardingRuleType, forwardingRuleMatch)
}
