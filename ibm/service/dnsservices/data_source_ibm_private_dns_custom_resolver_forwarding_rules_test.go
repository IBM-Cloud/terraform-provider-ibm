// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package dnsservices_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMPrivateDNSCustomResolverForwardingRulesDataSource_basic(t *testing.T) {
	forwardingRuleDescription := "test-forward-rule"
	forwardingRuleType := "zone"
	forwardingRuleMatch := "test.example.com"
	node := "data.ibm_dns_custom_resolver_forwarding_rules.test-fr"
	vpcname := fmt.Sprintf("d-fr-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("d-fr-subnet-name-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIbmDnsCrForwardingRulesDataSourceConfig(vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, forwardingRuleDescription, forwardingRuleType, forwardingRuleMatch),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(node, "rules.0.description"),
					resource.TestCheckResourceAttrSet(node, "rules.0.type"),
					resource.TestCheckResourceAttrSet(node, "rules.0.match"),
				),
			},
		},
	})
}

func testAccCheckIbmDnsCrForwardingRulesDataSourceConfig(vpcname, subnetname, zone, cidr, forwardingRuleDescription, forwardingRuleType string, forwardingRuleMatch string) string {
	return fmt.Sprintf(`

	data "ibm_resource_group" "rg" {
		is_default	= true
	}
	resource "ibm_is_vpc" "test-pdns-cr-vpc" {
		name			= "%s"
		resource_group	= data.ibm_resource_group.rg.id
	}
	resource "ibm_is_subnet" "test-pdns-cr-subnet1" {
		name			= "%s"
		vpc				= ibm_is_vpc.test-pdns-cr-vpc.id
		zone			= "%s"
		ipv4_cidr_block	= "%s"
		resource_group	= data.ibm_resource_group.rg.id
	}
	resource "ibm_resource_instance" "test-pdns-cr-instance" {
		name				= "test-pdns-cr-instance"
		resource_group_id	= data.ibm_resource_group.rg.id
		location			= "global"
		service				= "dns-svcs"
		plan				= "standard-dns"
	}
	resource "ibm_dns_custom_resolver" "test" {
		name		= "testpdnscustomresolver"
		instance_id = ibm_resource_instance.test-pdns-cr-instance.guid
		description = "new test CR - TF"
		high_availability = false
		enabled 	= true
		locations {
			subnet_crn	= ibm_is_subnet.test-pdns-cr-subnet1.crn
			enabled		= true
		}
	}
	resource "ibm_dns_custom_resolver_forwarding_rule" "dns_custom_resolver_forwarding_rule" {
		instance_id =  ibm_resource_instance.test-pdns-cr-instance.guid
		resolver_id = ibm_dns_custom_resolver.test.custom_resolver_id
		description = "Test Fw Rule"
		type = "%s"
		match = "%s"
		forward_to = ["168.20.22.122"]
	}		
	data "ibm_dns_custom_resolver_forwarding_rules" "test-fr" {
		depends_on  = [ibm_dns_custom_resolver.test]
		instance_id	= ibm_dns_custom_resolver.test.instance_id
		resolver_id = ibm_dns_custom_resolver.test.custom_resolver_id
	}
	`, vpcname, subnetname, zone, cidr, forwardingRuleType, forwardingRuleMatch)
}
