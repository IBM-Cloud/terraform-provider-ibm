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
			instance_id =  "d515a480-a702-4837-9f40-6c0c285262fd"
			description = "FW rules"
			high_availability = false
			enabled = true
			locations {
				subnet_crn = "crn:v1:staging:public:is:us-south-1:a/01652b251c3ae2787110a995d8db0135::subnet:0716-a094c4e8-02cd-4b04-858d-7f31205b93b9"
				enabled    = true
			}
		}
		resource "ibm_dns_custom_resolver_forwarding_rule" "dns_custom_resolver_forwarding_rule" {
			instance_id =  "d515a480-a702-4837-9f40-6c0c285262fd"
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
