// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIbmDnsCrForwardingRuleBasic(t *testing.T) {

	instanceID := "345ca2c4-83bf-4c04-bb09-5d8ec4d425a8"
	resolverID := "db857778-2aa9-4d81-b954-08bcaba0dadf"
	description := "test forward rule"
	typeVar := "zone"
	match := "test.example.com"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmDnsCrForwardingRuleConfig(instanceID, resolverID, description, typeVar, match),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_dns_cr_forwarding_rule.dns_cr_forwarding_rule", "instance_id", instanceID),
					resource.TestCheckResourceAttr("ibm_dns_cr_forwarding_rule.dns_cr_forwarding_rule", "resolver_id", resolverID),
					resource.TestCheckResourceAttr("ibm_dns_cr_forwarding_rule.dns_cr_forwarding_rule", "description", description),
					resource.TestCheckResourceAttr("ibm_dns_cr_forwarding_rule.dns_cr_forwarding_rule", "type", typeVar),
					resource.TestCheckResourceAttr("ibm_dns_cr_forwarding_rule.dns_cr_forwarding_rule", "match", match),
				),
			},
			{
				ResourceName:      "ibm_dns_cr_forwarding_rule.dns_cr_forwarding_rule",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"type"},
			},
		},
	})
}

func testAccCheckIbmDnsCrForwardingRuleConfig(instanceID string, resolverID string, description string, typeVar string, match string) string {
	return fmt.Sprintf(`

		resource "ibm_dns_cr_forwarding_rule" "dns_cr_forwarding_rule" {
			instance_id = "%s"
			resolver_id = "%s"
			description = "%s"
			type = "%s"
			match = "%s"
			forward_to = ["168.20.22.122"]
		}
	`, instanceID, resolverID, description, typeVar, match)
}
