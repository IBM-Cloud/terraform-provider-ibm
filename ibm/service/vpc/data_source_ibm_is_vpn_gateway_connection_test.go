// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMIsVPNGatewayConnectionDataSourceBasic(t *testing.T) {
	vpcname := fmt.Sprintf("tfvpnuat-vpc-%d", acctest.RandIntRange(100, 200))
	subnetname := fmt.Sprintf("tfvpnuat-subnet-%d", acctest.RandIntRange(100, 200))
	vpngwname := fmt.Sprintf("tfvpnuat-vpngw-%d", acctest.RandIntRange(100, 200))
	name := fmt.Sprintf("tfvpnuat-createname-%d", acctest.RandIntRange(100, 200))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIsVPNGatewayConnectionDataSourceConfigBasic(vpcname, subnetname, vpngwname, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateway_connection.example", "admin_state_up"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateway_connection.example", "authentication_mode"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateway_connection.example", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateway_connection.example", "dead_peer_detection.0.action"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateway_connection.example", "dead_peer_detection.0.interval"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateway_connection.example", "dead_peer_detection.0.timeout"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateway_connection.example", "tunnels.0.public_ip_address"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateway_connection.example", "tunnels.0.status"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateway_connection.example", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateway_connection.example", "mode"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateway_connection.example", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateway_connection.example", "peer_address"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateway_connection.example", "psk"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateway_connection.example", "resource_type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateway_connection.example", "status"),
				),
			},
			{
				Config: testAccCheckIBMIsVPNGatewayConnectionDataSourceConfigBasic(vpcname, subnetname, vpngwname, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateway_connection.example1", "admin_state_up"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateway_connection.example1", "authentication_mode"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateway_connection.example1", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateway_connection.example", "dead_peer_detection.0.action"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateway_connection.example", "dead_peer_detection.0.interval"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateway_connection.example", "dead_peer_detection.0.timeout"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateway_connection.example", "tunnels.0.public_ip_address"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateway_connection.example", "tunnels.0.status"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateway_connection.example1", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateway_connection.example1", "mode"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateway_connection.example1", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateway_connection.example1", "peer_address"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateway_connection.example1", "psk"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateway_connection.example1", "resource_type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateway_connection.example1", "status"),
				),
			},
			{
				Config: testAccCheckIBMIsVPNGatewayConnectionDataSourceConfigBasic(vpcname, subnetname, vpngwname, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateway_connection.example2", "admin_state_up"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateway_connection.example2", "authentication_mode"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateway_connection.example2", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateway_connection.example", "dead_peer_detection.0.action"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateway_connection.example", "dead_peer_detection.0.interval"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateway_connection.example", "dead_peer_detection.0.timeout"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateway_connection.example", "tunnels.0.public_ip_address"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateway_connection.example", "tunnels.0.status"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateway_connection.example2", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateway_connection.example2", "mode"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateway_connection.example2", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateway_connection.example2", "peer_address"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateway_connection.example2", "psk"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateway_connection.example2", "resource_type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateway_connection.example2", "status"),
				),
			},
			{
				Config: testAccCheckIBMIsVPNGatewayConnectionDataSourceConfigBasic(vpcname, subnetname, vpngwname, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateway_connection.example3", "admin_state_up"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateway_connection.example3", "authentication_mode"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateway_connection.example3", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateway_connection.example", "dead_peer_detection.0.action"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateway_connection.example", "dead_peer_detection.0.interval"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateway_connection.example", "dead_peer_detection.0.timeout"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateway_connection.example", "tunnels.0.public_ip_address"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateway_connection.example", "tunnels.0.status"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateway_connection.example3", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateway_connection.example3", "mode"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateway_connection.example3", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateway_connection.example3", "peer_address"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateway_connection.example3", "psk"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateway_connection.example3", "resource_type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateway_connection.example3", "status"),
				),
			},
		},
	})
}

func testAccCheckIBMIsVPNGatewayConnectionDataSourceConfigBasic(vpc, subnet, vpngwname, name string) string {
	return fmt.Sprintf(`
    	resource "ibm_is_vpc" "example" {
    		name = "%s"
    		
    	}
    	resource "ibm_is_subnet" "example" {
    		name = "%s"
    		vpc = "${ibm_is_vpc.example.id}"
    		zone = "%s"
    		ipv4_cidr_block = "%s"
    		
    	}
    	resource "ibm_is_vpn_gateway" "example" {
			name = "%s"
			subnet = "${ibm_is_subnet.example.id}"
			
    	
    	}
    	resource "ibm_is_vpn_gateway_connection" "example" {
    		name = "%s"
    		vpn_gateway = "${ibm_is_vpn_gateway.example.id}"
    		peer_address = "1.2.3.4"
			local_cidrs   = [ibm_is_subnet.example.ipv4_cidr_block]
    		preshared_key = "VPNDemoPassword"
    	}
		data "ibm_is_vpn_gateway_connection" "example" {
    		vpn_gateway = ibm_is_vpn_gateway.example.id
    		vpn_gateway_connection = ibm_is_vpn_gateway_connection.example.gateway_connection
    	}
    	data "ibm_is_vpn_gateway_connection" "example1" {
    		vpn_gateway = ibm_is_vpn_gateway.example.id
    		vpn_gateway_connection_name = ibm_is_vpn_gateway_connection.example.name
    	}
    	data "ibm_is_vpn_gateway_connection" "example2" {
    		vpn_gateway_name = ibm_is_vpn_gateway.example.name
    		vpn_gateway_connection = ibm_is_vpn_gateway_connection.example.gateway_connection
    	}
    	data "ibm_is_vpn_gateway_connection" "example3" {
    		vpn_gateway_name = ibm_is_vpn_gateway.example.name
    		vpn_gateway_connection_name = ibm_is_vpn_gateway_connection.example.name
    	}
	`, vpc, subnet, acc.ISZoneName, acc.ISCIDR, vpngwname, name)
}
