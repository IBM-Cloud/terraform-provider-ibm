// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"

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
	name = "%s"
	subnet = "${ibm_is_subnet.testacc_subnet.id}"
	}
	data "ibm_is_vpn_gateways" "test1" {
		
	}`, vpc, subnet, acc.ISZoneName, acc.ISCIDR, name)

}
