// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMISVpnGatewaysDataSource_basic(t *testing.T) {
	var vpnGateway string
	node := "data.ibm_is_vpn_gateways.test1"
	vpcname := fmt.Sprintf("tfvpnuat-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tfvpnuat-subnet-%d", acctest.RandIntRange(10, 100))
	name1 := fmt.Sprintf("tfvpnuat-createname-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISVpnGatewaysDataSourceConfig(vpcname, subnetname, name1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVPNGatewayExists("ibm_is_vpn_gateway.testacc_vpnGateway", vpnGateway),
					resource.TestCheckResourceAttrSet(node, "vpn_gateways.#"),
				),
			},
			{
				Config: testAccCheckIBMISVpnGatewaysDataSourceConfig(vpcname, subnetname, name1),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateways.test1", "vpn_gateways.0.vpc.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateways.test1", "vpn_gateways.0.vpc.0.name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateways.test1", "vpn_gateways.0.vpc.0.crn"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateways.test1", "vpn_gateways.0.vpc.0.href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateways.test1", "vpn_gateways.0.vpc.0.id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateways.test1", "vpn_gateways.0.local_asn"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateways.test1", "vpn_gateways.0.advertised_cidrs.#"),
				),
			},
		},
	})
}

func testAccCheckIBMISVpnGatewaysDataSourceConfig(vpc, subnet, name string) string {
	// status filter defaults to empty
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
		
	}
	resource "ibm_is_subnet" "testacc_subnet" {
		name = "%s"
		vpc = "${ibm_is_vpc.testacc_vpc.id}"
		zone = "%s"
		ipv4_cidr_block = "%s"
		
	}
	resource "ibm_is_vpn_gateway" "testacc_vpnGateway" {
		name   = "%s"
		subnet = ibm_is_subnet.testacc_subnet.id
		mode   = "route"
		local_asn = 64520
		lifecycle {
			ignore_changes = [
				advertised_cidrs
			]
  		}
	}
	resource "ibm_is_vpn_gateway_advertised_cidr" "example" {
		vpn_gateway = ibm_is_vpn_gateway.testacc_vpnGateway.id
		cidr        = "10.45.0.0/25"
	}	
	data "ibm_is_vpn_gateways" "test1" {
		
	}`, vpc, subnet, acc.ISZoneName, acc.ISCIDR, name)

}

func TestAccIBMISVpnGatewaysDataSource_regional(t *testing.T) {
	var vpnGateway string
	node := "data.ibm_is_vpn_gateways.test1"
	vpcname := fmt.Sprintf("tfvpnuat-vpc-%d", acctest.RandIntRange(10, 100))
	subnet1name := fmt.Sprintf("tfvpnuat-subnet1-%d", acctest.RandIntRange(10, 100))
	subnet2name := fmt.Sprintf("tfvpnuat-subnet2-%d", acctest.RandIntRange(10, 100))
	name1 := fmt.Sprintf("tfvpnuat-regional-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISVpnGatewaysDataSourceRegionalConfig(vpcname, subnet1name, subnet2name, name1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVPNGatewayExists("ibm_is_vpn_gateway.testacc_vpnGateway", vpnGateway),
					resource.TestCheckResourceAttrSet(node, "vpn_gateways.#"),
					resource.TestCheckResourceAttrSet(node, "vpn_gateways.0.availability_mode"),
					resource.TestCheckResourceAttrSet(node, "vpn_gateways.0.members.#"),
				),
			},
			{
				Config: testAccCheckIBMISVpnGatewaysDataSourceRegionalConfig(vpcname, subnet1name, subnet2name, name1),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateways.test1", "vpn_gateways.0.vpc.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateways.test1", "vpn_gateways.0.vpc.0.name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateways.test1", "vpn_gateways.0.vpc.0.crn"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateways.test1", "vpn_gateways.0.vpc.0.href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateways.test1", "vpn_gateways.0.vpc.0.id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateways.test1", "vpn_gateways.0.local_asn"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateways.test1", "vpn_gateways.0.advertised_cidrs.#"),
					resource.TestCheckResourceAttr("data.ibm_is_vpn_gateways.test1", "vpn_gateways.0.members.#", "2"),
				),
			},
		},
	})
}

func testAccCheckIBMISVpnGatewaysDataSourceRegionalConfig(vpc, subnet1, subnet2, name string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
		
	}
	resource "ibm_is_subnet" "testacc_subnet1" {
		name = "%s"
		vpc = ibm_is_vpc.testacc_vpc.id
		zone = "%s"
		ipv4_cidr_block = "10.240.20.0/24"
		
	}
	resource "ibm_is_subnet" "testacc_subnet2" {
		name = "%s"
		vpc = ibm_is_vpc.testacc_vpc.id
		zone = "%s"
		ipv4_cidr_block = "10.240.21.0/24"
		
	}
	resource "ibm_is_vpn_gateway" "testacc_vpnGateway" {
		name   = "%s"
		availability_mode = "regional"
		mode   = "route"
		local_asn = 64520
		members {
			private_ip {
				subnet {
					id = ibm_is_subnet.testacc_subnet1.id
				}
			}
		}
		members {
			private_ip {
				subnet {
					id = ibm_is_subnet.testacc_subnet2.id
				}
			}
		}
		lifecycle {
			ignore_changes = [
				advertised_cidrs
			]
  		}
	}
	resource "ibm_is_vpn_gateway_advertised_cidr" "example" {
		vpn_gateway = ibm_is_vpn_gateway.testacc_vpnGateway.id
		cidr        = "10.45.0.0/25"
	}
	data "ibm_is_vpn_gateways" "test1" {
		
	}`, vpc, subnet1, acc.ISZoneName, subnet2, acc.ISZoneName2, name)

}
