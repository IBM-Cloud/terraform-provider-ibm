// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMPrivateDNSCustomResolverForwardingRulesDataSource_basic(t *testing.T) {
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
	data "ibm_resource_group" "rg" {
		is_default=true
	}
	resource "ibm_is_vpc" "test-pdns-cr-vpc" {
		name = "test-pdns-custom-resolver-vpc"
		resource_group = data.ibm_resource_group.rg.id
	}
	resource "ibm_is_subnet" "test-pdns-cr-subnet1" {
		name                    = "test-pdns-cr-subnet1"
		vpc                     = ibm_is_vpc.test-pdns-cr-vpc.id
		zone            		= "us-south-1"
		ipv4_cidr_block 		= "10.240.0.0/24"
		resource_group 			= data.ibm_resource_group.rg.id
	}
	resource "ibm_is_subnet" "test-pdns-cr-subnet2" {
		name                    = "test-pdns-cr-subnet2"
		vpc                     = ibm_is_vpc.test-pdns-cr-vpc.id
		zone            		= "us-south-1"
		ipv4_cidr_block 		= "10.240.64.0/24"
		resource_group 			= data.ibm_resource_group.rg.id
	}
	resource "ibm_resource_instance" "test-pdns-cr-instance" {
		name = "test-pdns-cr-instance"
		resource_group_id = data.ibm_resource_group.rg.id
		location = "global"
		service = "dns-svcs"
		plan = "standard-dns"
	}
	resource "ibm_dns_custom_resolver" "test" {
		name        = "test-pdns-customresolver"
		instance_id = ibm_resource_instance.test-pdns-cr-instance.guid
		description = "test-pdns-cr-instance"
		locations {
			subnet_crn = ibm_is_subnet.test-pdns-cr-subnet1.crn
			enabled     = true
		}
		locations {
			subnet_crn = ibm_is_subnet.test-pdns-cr-subnet2.crn
			enabled     = true
		
	}
	resource "ibm_dns_custom_resolver_forwarding_rule" "dns_custom_resolver_forwarding_rule" {
		instance_id = ibm_resource_instance.test-pdns-cr-instance.guid
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
