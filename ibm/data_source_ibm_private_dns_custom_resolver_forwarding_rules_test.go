// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMDNSCustomResolverForwardingRulesDataSource_basic(t *testing.T) {
	forwardingRuleDescription := "test forward rule"
	forwardingRuleType := "zone"
	forwardingRuleMatch := "test.example.com"
	node := "data.ibm_dns_custom_resolver_forwarding_rules.dns_custom_resolver_forwarding_rules"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmDnsCrForwardingRulesDataSourceConfig(forwardingRuleDescription, forwardingRuleType, forwardingRuleMatch),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(node, "ibm_dns_custom_resolver_forwarding_rule.0.description"),
					resource.TestCheckResourceAttrSet(node, "ibm_dns_custom_resolver_forwarding_rule.0.type"),
					resource.TestCheckResourceAttrSet(node, "ibm_dns_custom_resolver_forwarding_rule.0.match"),
				),
			},
		},
	})
}

func testAccCheckIbmDnsCrForwardingRulesDataSourceConfig(forwardingRuleDescription string, forwardingRuleType string, forwardingRuleMatch string) string {
	return fmt.Sprintf(`
		resource "ibm_dns_custom_resolver" "test" {
			name        = "CustomResolverFW"
			instance_id = "345ca2c4-83bf-4c04-bb09-5d8ec4d425a8"
			description = "FW rules"
			enabled = true
			locations {
				subnet_crn = "crn:v1:staging:public:is:us-south-1:a/01652b251c3ae2787110a995d8db0135::subnet:0716-03d54d71-b438-4d20-b943-76d3d2a1a590"
				enabled    = true
			}
		}
		resource "ibm_dns_custom_resolver_forwarding_rule" "dns_custom_resolver_forwarding_rule" {
			instance_id = "345ca2c4-83bf-4c04-bb09-5d8ec4d425a8"
			resolver_id = ibm_dns_custom_resolver.test.custom_resolver_id
			description = "%s"
			type = "%s"
			match = "%s"
			forward_to = ["168.20.22.122"]
		}
		data "ibm_dns_custom_resolver_forwarding_rules" "dns_custom_resolver_forwarding_rules" {
			instance_id = ibm_dns_custom_resolver_forwarding_rule.dns_custom_resolver_forwarding_rule.instance_id
			resolver_id = ibm_dns_custom_resolver_forwarding_rule.dns_custom_resolver_forwarding_rule.resolver_id
		  }
	`, forwardingRuleDescription, forwardingRuleType, forwardingRuleMatch)
}
