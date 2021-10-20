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
		instance_id		= "d515a480-a702-4837-9f40-6c0c285262fd"
		description		= "new test CR Locations - TF"
		high_availability =  true
		enabled		= true
		locations{
			subnet_crn	= "crn:v1:staging:public:is:us-south-1:a/01652b251c3ae2787110a995d8db0135::subnet:0716-a094c4e8-02cd-4b04-858d-7f31205b93b9"
			enabled		= true
		}
		locations {
			subnet_crn  = "crn:v1:staging:public:is:us-south-2:a/01652b251c3ae2787110a995d8db0135::subnet:0726-b6f3cb83-48f0-4c55-9023-202fe4570c83"
			enabled     = false
		}
	}
	resource "ibm_dns_custom_resolver_forwarding_rule" "dns_custom_resolver_forwarding_rule" {
		instance_id = "d515a480-a702-4837-9f40-6c0c285262fd"
		resolver_id = ibm_dns_custom_resolver.test.custom_resolver_id
		description = "Test Fw Rule"
		type = "%s"
		match = "%s"
		forward_to = ["168.20.22.122"]
	}		
	`, typeVar, match)
}
