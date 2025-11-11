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

func TestAccIBMISVpnGatewayServoceConnectionsDataSource_basic(t *testing.T) {
	conn := "data.ibm_is_vpn_gateway_service_connections.test1"
	vpcname := fmt.Sprintf("tfvpnuat-vpc-%d", acctest.RandIntRange(100, 200))
	subnetname := fmt.Sprintf("tfvpnuat-subnet-%d", acctest.RandIntRange(100, 200))
	vpngwname := fmt.Sprintf("tfvpnuat-vpngw-%d", acctest.RandIntRange(100, 200))
	name := fmt.Sprintf("tfvpnuat-createname-%d", acctest.RandIntRange(100, 200))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISVpnGatewayServiceconnectionsDataSourceConfig(vpcname, subnetname, vpngwname, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(conn, "service_connections.#"),
				),
			},
		},
	})
}

func testAccCheckIBMISVpnGatewayServiceconnectionsDataSourceConfig(vpc, subnet, vpngwname, name string) string {
	return fmt.Sprintf(`
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
	data "ibm_is_vpn_gateway_service_connections" "test1" {
		vpn_gateway = ibm_is_vpn_gateway.testacc_vpnGateway.id
	}`, vpc, subnet, acc.ISZoneName, acc.ISCIDR, vpngwname, name)

}
