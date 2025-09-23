// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package dnsservices_test

import (
	"fmt"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMPrivateDNSNetworkDataSource_basic(t *testing.T) {
	node := "data.ibm_dns_permitted_networks.test"
	riname := fmt.Sprintf("tf-instnace-%d", acctest.RandIntRange(100, 200))
	vpcname := fmt.Sprintf("tfc-vpc-name-%d", acctest.RandIntRange(10, 100))
	zonename := fmt.Sprintf("tf-dnszone-%d.com", acctest.RandIntRange(100, 200))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMpDNSPermittedNetworksDataSourceConfig(riname, vpcname, zonename),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(node, "dns_permitted_networks.0.permitted_network_id"),
					resource.TestCheckResourceAttrSet(node, "dns_permitted_networks.0.state"),
					resource.TestCheckResourceAttrSet(node, "dns_permitted_networks.0.type"),
				),
			},
		},
	})
}

func testAccCheckIBMpDNSPermittedNetworksDataSourceConfig(riname, vpcname, zonename string) string {
	// status filter defaults to empty
	return fmt.Sprintf(`
	data "ibm_resource_group" "rg" {
		is_default=true
	}

	resource "ibm_resource_instance" "test-pdns-instance" {
		name = "%s"
		resource_group_id = data.ibm_resource_group.rg.id
		location = "global"
		service = "dns-svcs"
		plan = "standard-dns"
	}

	resource "ibm_is_vpc" "test_pdns_vpc" {
		name = "%s"
		resource_group = data.ibm_resource_group.rg.id
	}

	resource "ibm_dns_zone" "test-pdns-zone" {
		name        = "%s"
		instance_id = ibm_resource_instance.test-pdns-instance.guid
		description = "testdescription6"
		label       = "testlabel-updated6"
	}

	resource "ibm_dns_permitted_network" "test-pdns-permitted-network-nw" {
		instance_id = ibm_dns_zone.test-pdns-zone.instance_id
		zone_id     = ibm_dns_zone.test-pdns-zone.zone_id
		vpc_crn     = ibm_is_vpc.test_pdns_vpc.crn
	}

	data "ibm_dns_permitted_networks" "test" {
		instance_id = ibm_dns_permitted_network.test-pdns-permitted-network-nw.instance_id
		zone_id = ibm_dns_permitted_network.test-pdns-permitted-network-nw.zone_id
	}`, riname, vpcname, zonename)
}
