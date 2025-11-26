// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMIsVPNGatewayAdvertisedCidrsDataSourceBasic(t *testing.T) {
	vpcname := fmt.Sprintf("tfvpnuat-vpc-%d", acctest.RandIntRange(100, 200))
	subnetname := fmt.Sprintf("tfvpnuat-subnet-%d", acctest.RandIntRange(100, 200))
	vpngwname := fmt.Sprintf("tfvpnuat-vpngw-%d", acctest.RandIntRange(100, 200))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIsVPNGatewayAdvertisedCidrsDataSourceConfigBasic(vpcname, subnetname, vpngwname),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateway_advertised_cidrs.is_vpn_gateway_advertised_cidrs", "advertised_cidrs.#"),
				),
			},
		},
	})
}

func TestAccIBMIsVPNGatewayAdvertisedCidrsVPNGatewayNameDataSourceBasic(t *testing.T) {
	vpcname := fmt.Sprintf("tfvpnuat-vpc-%d", acctest.RandIntRange(100, 200))
	subnetname := fmt.Sprintf("tfvpnuat-subnet-%d", acctest.RandIntRange(100, 200))
	vpngwname := fmt.Sprintf("tfvpnuat-vpngw-%d", acctest.RandIntRange(100, 200))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIsVPNGatewayAdvertisedCidrsWithVPNGatewayNameDataSourceConfigBasic(vpcname, subnetname, vpngwname),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_is_vpn_gateway_advertised_cidrs.is_vpn_gateway_advertised_cidrs", "advertised_cidrs.#"),
				),
			},
		},
	})
}

func testAccCheckIBMIsVPNGatewayAdvertisedCidrsDataSourceConfigBasic(vpc, subnet, vpngwname string) string {
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
		mode   = "route"
		local_asn = 64520
		lifecycle {
			ignore_changes = [
				advertised_cidrs
			]
  		}
	  }

	  resource "ibm_is_vpn_gateway_advertised_cidr" "example" {
		vpn_gateway = ibm_is_vpn_gateway.example.id
		cidr        = "10.45.0.0/25"
	  }
	  data "ibm_is_vpn_gateway_advertised_cidrs" "is_vpn_gateway_advertised_cidrs" {
	  	depends_on = [resource.ibm_is_vpn_gateway_advertised_cidr.example]
		vpn_gateway = ibm_is_vpn_gateway.example.id
	  }
	`, vpc, subnet, acc.ISZoneName, acc.ISCIDR, vpngwname)
}

func testAccCheckIBMIsVPNGatewayAdvertisedCidrsWithVPNGatewayNameDataSourceConfigBasic(vpc, subnet, vpngwname string) string {
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
		mode   = "route"
		lifecycle {
			ignore_changes = [
				advertised_cidrs
			]
  		}
	  }
	  resource "ibm_is_vpn_gateway_advertised_cidr" "example" {
		vpn_gateway = ibm_is_vpn_gateway.example.id
		cidr        = "10.45.0.0/25"
	  }
	  data "ibm_is_vpn_gateway_advertised_cidrs" "is_vpn_gateway_advertised_cidrs" {
	  	depends_on = [resource.ibm_is_vpn_gateway_advertised_cidr.example]
		vpn_gateway_name = ibm_is_vpn_gateway.example.name
	  }
	`, vpc, subnet, acc.ISZoneName, acc.ISCIDR, vpngwname)
}
