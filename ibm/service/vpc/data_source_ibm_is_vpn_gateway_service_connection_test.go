// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIBMIsVPNGatewayServiceConnectionDataSourceBasic(t *testing.T) {
	vpcname := fmt.Sprintf("tfvpnuat-vpc-%d", acctest.RandIntRange(100, 200))
	subnetname := fmt.Sprintf("tfvpnuat-subnet-%d", acctest.RandIntRange(100, 200))
	vpngwname := fmt.Sprintf("tfvpnuat-vpngw-%d", acctest.RandIntRange(100, 200))
	name := fmt.Sprintf("tfvpnuat-createname-%d", acctest.RandIntRange(100, 200))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIsVPNGatewayServiceConnectionDataSourceConfigBasic(vpcname, subnetname, vpngwname, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateway_service_connection.example", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateway_service_connection.example", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateway_service_connection.example", "lifecycle_state"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateway_service_connection.example", "status"),
				),
			},
			{
				Config: testAccCheckIBMIsVPNGatewayServiceConnectionDataSourceConfigBasic(vpcname, subnetname, vpngwname, name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateway_service_connection.example1", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateway_service_connection.example1", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateway_service_connection.example1", "lifecycle_state"),
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateway_service_connection.example1", "status"),
				),
			},
		},
	})
}

func testAccCheckIBMIsVPNGatewayServiceConnectionDataSourceConfigBasic(vpc, subnet, vpngwname, name string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "example" {
		name = "%s"
	  
	  }
	  resource "ibm_is_subnet" "example" {
		name            = "%s"
		vpc             = ibm_is_vpc.example.id
		zone            = "%s"
		ipv4_cidr_block = "%s"
	  
	  }
	  resource "ibm_is_vpn_gateway" "example" {
		name   = "%s"
		subnet = ibm_is_subnet.example.id
		mode   = "policy"
	  
	  }
	  resource "ibm_is_vpn_gateway_connection" "example" {
		name          = "%s"
		vpn_gateway   = ibm_is_vpn_gateway.example.id
		peer_address  = "1.2.3.4"
		peer_cidrs    = [ibm_is_subnet.example.ipv4_cidr_block]
		local_cidrs   = [ibm_is_subnet.example.ipv4_cidr_block]
		preshared_key = "VPNDemoPassword"
	  }
	  data "ibm_is_vpn_gateway_service_connection" "example" {
		vpn_gateway            = ibm_is_vpn_gateway.example.id
		vpn_gateway_service_connection = ibm_is_vpn_gateway_connection.example.gateway_connection
	  }
	  data "ibm_is_vpn_gateway_service_connection" "example1" {
		vpn_gateway_name       = ibm_is_vpn_gateway.example.name
		vpn_gateway_service_connection = ibm_is_vpn_gateway_connection.example.gateway_connection
	  }
	`, vpc, subnet, acc.ISZoneName, acc.ISCIDR, vpngwname, name)
}
