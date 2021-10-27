// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMPrivateDNSCustomResolverForwardingRulesDataSource_basic(t *testing.T) {
	forwardingRuleDescription := "test-forward-rule"
	forwardingRuleType := "zone"
	forwardingRuleMatch := "test.example.com"
	node := "data.ibm_dns_custom_resolver_forwarding_rules.test-fr"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmDnsCrForwardingRulesDataSourceConfig(forwardingRuleDescription, forwardingRuleType, forwardingRuleMatch),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(node, "rules.0.description"),
					resource.TestCheckResourceAttrSet(node, "rules.0.type"),
					resource.TestCheckResourceAttrSet(node, "rules.0.match"),
				),
			},
		},
	})
}

func testAccCheckIbmDnsCrForwardingRulesDataSourceConfig(forwardingRuleDescription string, forwardingRuleType string, forwardingRuleMatch string) string {
	return fmt.Sprintf(`
		resource "ibm_dns_custom_resolver" "test" {
			name        = "CustomResolverFW"
			instance_id = "c9e23743-b039-4f33-ba8a-c3bf35e9b450"
			description = "FW rules"
			high_availability = false
			enabled = true
			locations {
				subnet_crn = "crn:v1:bluemix:public:is:us-south-1:a/bcf1865e99742d38d2d5fc3fb80a5496::subnet:0717-4f53a236-cd7a-4688-9347-066bb5058a5c"
				enabled    = true
			}
		}
		resource "ibm_dns_custom_resolver_forwarding_rule" "dns_custom_resolver_forwarding_rule" {
			instance_id =  "c9e23743-b039-4f33-ba8a-c3bf35e9b450"
			resolver_id = ibm_dns_custom_resolver.test.custom_resolver_id
			description = "%s"
			type = "%s"
			match = "%s"
			forward_to = ["168.20.22.122"]
		}
	
		data "ibm_dns_custom_resolver_forwarding_rules" "test-fr" {
			depends_on  = [ibm_dns_custom_resolver.test]
			instance_id	= ibm_dns_custom_resolver.test.instance_id
			resolver_id = ibm_dns_custom_resolver.test.custom_resolver_id
		}
	`, forwardingRuleDescription, forwardingRuleType, forwardingRuleMatch)
}
