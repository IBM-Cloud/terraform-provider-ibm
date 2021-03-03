// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMISVpnGatewayConnectionsDataSource_basic(t *testing.T) {
	var vpnGatewayConnection string
	node := "data.ibm_is_vpn_gateway_connections.test1"
	vpcname := fmt.Sprintf("tfvpnuat-vpc-%d", acctest.RandIntRange(100, 200))
	subnetname := fmt.Sprintf("tfvpnuat-subnet-%d", acctest.RandIntRange(100, 200))
	vpngwname := fmt.Sprintf("tfvpnuat-vpngw-%d", acctest.RandIntRange(100, 200))
	name := fmt.Sprintf("tfvpnuat-createname-%d", acctest.RandIntRange(100, 200))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISVpnGatewayconnectionsDataSourceConfig(vpcname, subnetname, vpngwname, name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVPNGatewayConnectionExists("ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection", vpnGatewayConnection),
					resource.TestCheckResourceAttrSet(node, "connections.#"),
				),
			},
		},
	})
}

func testAccCheckIBMISVpnGatewayconnectionsDataSourceConfig(vpc, subnet, vpngwname, name string) string {
	// status filter defaults to empty
	return fmt.Sprintf(`
	
	data "ibm_resource_group" "rg" {
		name = "Proof of Concepts"
	}
	resource "ibm_is_vpc" "testacc_vpc" {
		name = "%s"
		resource_group = data.ibm_resource_group.rg.id
	}
	resource "ibm_is_subnet" "testacc_subnet" {
		name = "%s"
		vpc = "${ibm_is_vpc.testacc_vpc.id}"
		zone = "%s"
		ipv4_cidr_block = "%s"
		resource_group = data.ibm_resource_group.rg.id
	}
	resource "ibm_is_vpn_gateway" "testacc_vpnGateway" {
	name = "%s"
	subnet = "${ibm_is_subnet.testacc_subnet.id}"
	resource_group = data.ibm_resource_group.rg.id
	
	}
	resource "ibm_is_vpn_gateway_connection" "testacc_VPNGatewayConnection" {
		name = "%s"
		vpn_gateway = "${ibm_is_vpn_gateway.testacc_vpnGateway.id}"
		peer_address = "1.2.3.4"
		preshared_key = "VPNDemoPassword"
	}
	data "ibm_is_vpn_gateway_connections" "test1" {
		vpn_gateway = ibm_is_vpn_gateway.testacc_vpnGateway.id
	}`, vpc, subnet, ISZoneName, ISCIDR, vpngwname, name)

}
