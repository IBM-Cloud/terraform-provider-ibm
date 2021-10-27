// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMPrivateDNSCustomResolverForwardingRule_basic(t *testing.T) {
	typeVar := "zone"
	match := "test.example.com"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmDnsCrForwardingRuleConfig(typeVar, match),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_dns_custom_resolver_forwarding_rule.dns_custom_resolver_forwarding_rule", "type", typeVar),
					resource.TestCheckResourceAttr("ibm_dns_custom_resolver_forwarding_rule.dns_custom_resolver_forwarding_rule", "match", match),
				),
			},
		},
	})
}

func testAccCheckIbmDnsCrForwardingRuleConfig(typeVar, match string) string {
	return fmt.Sprintf(`
		
	resource "ibm_dns_custom_resolver" "test" {
		name			= "testpdnscustomresolver"
		instance_id		= "c9e23743-b039-4f33-ba8a-c3bf35e9b450"
		description		= "new test CR Locations - TF"
		high_availability =  true
		enabled		= true
		locations{
			subnet_crn	= "crn:v1:bluemix:public:is:us-south-3:a/bcf1865e99742d38d2d5fc3fb80a5496::subnet:0737-0d198509-3221-4162-b2d8-4a9326d3d7ad"
			enabled		= false
		}
		locations {
			subnet_crn  = "crn:v1:bluemix:public:is:us-south-2:a/bcf1865e99742d38d2d5fc3fb80a5496::subnet:0727-f17967f2-2bbe-427c-bcf6-22f8c2395285"
			enabled     = true
		}
	}
	resource "ibm_dns_custom_resolver_forwarding_rule" "dns_custom_resolver_forwarding_rule" {
		instance_id = "c9e23743-b039-4f33-ba8a-c3bf35e9b450"
		resolver_id = ibm_dns_custom_resolver.test.custom_resolver_id
		description = "Test Fw Rule"
		type = "%s"
		match = "%s"
		forward_to = ["168.20.22.122"]
	}		
	`, typeVar, match)
}
