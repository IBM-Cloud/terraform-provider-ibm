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

func TestAccIBMIsVPNGatewayDataSourceBasic(t *testing.T) {
	vpcname := fmt.Sprintf("tfvpnuat-vpc-%d", acctest.RandIntRange(100, 200))
	subnetname := fmt.Sprintf("tfvpnuat-subnet-%d", acctest.RandIntRange(100, 200))
	vpngwname := fmt.Sprintf("tfvpnuat-vpngw-%d", acctest.RandIntRange(100, 200))
	vpngwconname := fmt.Sprintf("tfvpnuat-vpngwconn-%d", acctest.RandIntRange(100, 200))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIsVPNGatewayDataSourceConfigBasic(vpcname, subnetname, vpngwname, vpngwconname),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateway.example", "connections.0.href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateway.example", "connections.0.id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateway.example", "connections.0.name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateway.example", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateway.example", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateway.example", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateway.example", "members.0.public_ip_address"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateway.example", "members.0.role"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateway.example", "members.0.status"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateway.example", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateway.example", "resource_group.0.id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateway.example", "resource_group.0.href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateway.example", "resource_group.0.name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateway.example", "resource_type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateway.example", "status"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateway.example", "subnet.0.href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateway.example", "subnet.0.id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateway.example", "subnet.0.name"),
				),
			},
			{
				Config: testAccCheckIBMIsVPNGatewayDataSourceConfigBasic(vpcname, subnetname, vpngwname, vpngwconname),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateway.example-name", "connections.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateway.example-name", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateway.example-name", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateway.example-name", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateway.example-name", "members.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateway.example-name", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateway.example-name", "resource_group.#"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateway.example-name", "resource_type"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateway.example-name", "status"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateway.example-name", "subnet.#"),
				),
			},
		},
	})
}

func testAccCheckIBMIsVPNGatewayDataSourceConfigBasic(vpc, subnet, vpngwname, vpngwconname string) string {
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
			name          = "%s"
			vpn_gateway   = ibm_is_vpn_gateway.example.id
			peer_address  = ibm_is_vpn_gateway.example.public_ip_address
			preshared_key = "VPNDemoPassword"
			local_cidrs   = [ibm_is_subnet.example.ipv4_cidr_block]
			
		}
		data "ibm_is_vpn_gateway" "example" {
			depends_on = [
				ibm_is_vpn_gateway_connection.example
			]
			vpn_gateway = ibm_is_vpn_gateway.example.id
		}
		data "ibm_is_vpn_gateway" "example-name" {
			depends_on = [
				ibm_is_vpn_gateway_connection.example
			]
			vpn_gateway_name = ibm_is_vpn_gateway.example.name
		}
	`, vpc, subnet, acc.ISZoneName, acc.ISCIDR, vpngwname, vpngwconname)
}
