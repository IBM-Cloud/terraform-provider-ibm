// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"

	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccIBMISVPNGatewayAdvertisedCidr_basic(t *testing.T) {
	var advertisedCidr string

	vpcname := fmt.Sprintf("tf-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tf-subnet-%d", acctest.RandIntRange(10, 100))
	vpnname := fmt.Sprintf("tfvpngc-vpn-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISVPNGatewayAdvertisedCidrDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISVPNGatewayAdvertisedCidrConfig(vpcname, subnetname, vpnname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVPNGatewayAdvertisedCidrExists("ibm_is_vpn_gateway_advertised_cidr.example", &advertisedCidr),
				),
			},
		},
	})
}

func testAccCheckIBMISVPNGatewayAdvertisedCidrDestroy(s *terraform.State) error {

	sess, err := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_is_vpn_gateway_advertised_cidr" {
			continue
		}
		parts, err := flex.IdParts(rs.Primary.ID)
		if err != nil {
			return err
		}
		gID := parts[0]
		gCidr := parts[1]

		removeVPNGatewayAdvertisedCIDROptions := &vpcv1.RemoveVPNGatewayAdvertisedCIDROptions{
			VPNGatewayID: &gID,
			CIDR:         &gCidr,
		}
		response, err := sess.RemoveVPNGatewayAdvertisedCIDR(removeVPNGatewayAdvertisedCIDROptions)
		if err == nil {
			return fmt.Errorf("Advertised Cidr still exists: %v", response)
		}
	}
	return nil
}

func testAccCheckIBMISVPNGatewayAdvertisedCidrExists(n string, advertisedCidr *string) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}
		if rs.Primary.ID == "" {
			return fmt.Errorf("[ERROR] No Advertised Cidr is set")
		}

		sess, err := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
		if err != nil {
			return err
		}

		parts, err := flex.IdParts(rs.Primary.ID)
		if err != nil {
			return err
		}
		gID := parts[0]
		gAdvertisedCidr := parts[1]

		checkVPNGatewayAdvertisedCIDROptions := &vpcv1.CheckVPNGatewayAdvertisedCIDROptions{
			VPNGatewayID: &gID,
			CIDR:         &gAdvertisedCidr,
		}

		response, err := sess.CheckVPNGatewayAdvertisedCIDR(checkVPNGatewayAdvertisedCIDROptions)
		if err != nil {
			if response != nil && response.StatusCode == 404 {
				*advertisedCidr = ""
				return nil
			}
			return fmt.Errorf("[ERROR] Error getting Advertised Cidr : %s\n%s", err, response)
		}
		*advertisedCidr = fmt.Sprintf("%s/%s", gID, gAdvertisedCidr)
		return nil
	}
}

func testAccCheckIBMISVPNGatewayAdvertisedCidrConfig(vpcname, subnetname, vpnname string) string {
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
	  }`, vpcname, subnetname, acc.ISZoneName, acc.ISCIDR, vpnname)

}
