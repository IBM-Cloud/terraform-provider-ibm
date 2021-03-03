// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"errors"
	"fmt"
	"testing"

	"github.com/IBM/vpc-go-sdk/vpcclassicv1"
	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccIBMISVPNGateway_basic(t *testing.T) {
	var vpnGateway string
	vpcname := fmt.Sprintf("tfvpnuat-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tfvpnuat-subnet-%d", acctest.RandIntRange(10, 100))
	name1 := fmt.Sprintf("tfvpnuat-createname-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMISVPNGatewayDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISVPNGatewayConfig(vpcname, subnetname, name1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVPNGatewayExists("ibm_is_vpn_gateway.testacc_vpnGateway", vpnGateway),
					resource.TestCheckResourceAttr(
						"ibm_is_vpn_gateway.testacc_vpnGateway", "name", name1),
				),
			},
		},
	})
}

func TestAccIBMISVPNGateway_route(t *testing.T) {
	var vpnGateway string
	vpcname := fmt.Sprintf("tfvpnuat-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tfvpnuat-subnet-%d", acctest.RandIntRange(10, 100))
	name1 := fmt.Sprintf("tfvpnuat-createname-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMISVPNGatewayDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISVPNGatewayRouteConfig(vpcname, subnetname, name1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVPNGatewayExists("ibm_is_vpn_gateway.testacc_vpnGateway", vpnGateway),
					resource.TestCheckResourceAttr(
						"ibm_is_vpn_gateway.testacc_vpnGateway", "name", name1),
					resource.TestCheckResourceAttr(
						"ibm_is_vpn_gateway.testacc_vpnGateway", "mode", "route"),
				),
			},
		},
	})
}

func testAccCheckIBMISVPNGatewayDestroy(s *terraform.State) error {
	userDetails, _ := testAccProvider.Meta().(ClientSession).BluemixUserDetails()

	if userDetails.generation == 1 {
		sess, _ := testAccProvider.Meta().(ClientSession).VpcClassicV1API()
		for _, rs := range s.RootModule().Resources {
			if rs.Type != "ibm_is_vpn_gateway" {
				continue
			}

			getvpngcptions := &vpcclassicv1.GetVPNGatewayConnectionOptions{
				ID: &rs.Primary.ID,
			}
			_, _, err := sess.GetVPNGatewayConnection(getvpngcptions)

			if err == nil {
				return fmt.Errorf("vpnGateway still exists: %s", rs.Primary.ID)
			}
		}
	} else {
		sess, _ := testAccProvider.Meta().(ClientSession).VpcV1API()
		for _, rs := range s.RootModule().Resources {
			if rs.Type != "ibm_is_vpn_gateway" {
				continue
			}

			getvpngcptions := &vpcv1.GetVPNGatewayConnectionOptions{
				ID: &rs.Primary.ID,
			}
			_, _, err := sess.GetVPNGatewayConnection(getvpngcptions)

			if err == nil {
				return fmt.Errorf("vpnGateway still exists: %s", rs.Primary.ID)
			}
		}
	}

	return nil
}

func testAccCheckIBMISVPNGatewayExists(n, vpnGatewayID string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return errors.New("No Record ID is set")
		}

		userDetails, _ := testAccProvider.Meta().(ClientSession).BluemixUserDetails()

		if userDetails.generation == 1 {
			sess, _ := testAccProvider.Meta().(ClientSession).VpcClassicV1API()
			getvpngcptions := &vpcclassicv1.GetVPNGatewayOptions{
				ID: &rs.Primary.ID,
			}
			foundvpnGatewayIntf, _, err := sess.GetVPNGateway(getvpngcptions)
			if err != nil {
				return err
			}
			foundvpnGateway := foundvpnGatewayIntf.(*vpcclassicv1.VPNGateway)
			vpnGatewayID = *foundvpnGateway.ID
		} else {
			sess, _ := testAccProvider.Meta().(ClientSession).VpcV1API()
			getvpngcptions := &vpcv1.GetVPNGatewayOptions{
				ID: &rs.Primary.ID,
			}
			foundvpnGatewayIntf, _, err := sess.GetVPNGateway(getvpngcptions)
			if err != nil {
				return err
			}
			foundvpnGateway := foundvpnGatewayIntf.(*vpcv1.VPNGateway)
			vpnGatewayID = *foundvpnGateway.ID
		}
		return nil
	}
}

func testAccCheckIBMISVPNGatewayConfig(vpc, subnet, name string) string {
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
	mode = "policy"
	}`, vpc, subnet, ISZoneName, ISCIDR, name)

}

func testAccCheckIBMISVPNGatewayRouteConfig(vpc, subnet, name string) string {
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
	mode = "route"
	}`, vpc, subnet, ISZoneName, ISCIDR, name)

}
