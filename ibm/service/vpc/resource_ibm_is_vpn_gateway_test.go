// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package vpc_test

import (
	"errors"
	"fmt"
	"regexp"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"

	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIBMISVPNGateway_basic(t *testing.T) {
	var vpnGateway string
	vpcname := fmt.Sprintf("tfvpnuat-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tfvpnuat-subnet-%d", acctest.RandIntRange(10, 100))
	name1 := fmt.Sprintf("tfvpnuat-createname-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISVPNGatewayDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMISVPNGatewayConfig(vpcname, subnetname, name1),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVPNGatewayExists("ibm_is_vpn_gateway.testacc_vpnGateway", vpnGateway),
					resource.TestCheckResourceAttr(
						"ibm_is_vpn_gateway.testacc_vpnGateway", "name", name1),
					resource.TestCheckResourceAttrSet("ibm_is_vpn_gateway.testacc_vpnGateway", "lifecycle_state"),
					resource.TestCheckResourceAttrSet("ibm_is_vpn_gateway.testacc_vpnGateway", "health_state"),
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
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
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
			{
				Config: testAccCheckIBMISVPNGatewayRouteConfig(vpcname, subnetname, name1),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("ibm_is_vpn_gateway.testacc_vpnGateway", "vpc.#"),
					resource.TestCheckResourceAttrSet("ibm_is_vpn_gateway.testacc_vpnGateway", "vpc.0.name"),
					resource.TestCheckResourceAttrSet("ibm_is_vpn_gateway.testacc_vpnGateway", "vpc.0.crn"),
					resource.TestCheckResourceAttrSet("ibm_is_vpn_gateway.testacc_vpnGateway", "vpc.0.href"),
					resource.TestCheckResourceAttrSet("ibm_is_vpn_gateway.testacc_vpnGateway", "vpc.0.id"),
				),
			},
		},
	})
}

func testAccCheckIBMISVPNGatewayDestroy(s *terraform.State) error {

	sess, _ := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
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
		sess, _ := acc.TestAccProvider.Meta().(conns.ClientSession).VpcV1API()
		getvpngcptions := &vpcv1.GetVPNGatewayOptions{
			ID: &rs.Primary.ID,
		}
		foundvpnGatewayIntf, _, err := sess.GetVPNGateway(getvpngcptions)
		if err != nil {
			return err
		}
		foundvpnGateway := foundvpnGatewayIntf.(*vpcv1.VPNGateway)
		vpnGatewayID = *foundvpnGateway.ID
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
	}`, vpc, subnet, acc.ISZoneName, acc.ISCIDR, name)

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
	}`, vpc, subnet, acc.ISZoneName, acc.ISCIDR, name)

}

func testAccCheckIBMISVPNGatewayTaintConfig(vpc, subnet, name string) string {
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
		name 	= "%s"
		subnet 	= "${ibm_is_subnet.testacc_subnet.id}"
		mode 	= "policy"
		timeouts{
			create = "2m"
		}
		lifecycle {
			create_before_destroy = true
		}
	}`, vpc, subnet, acc.ISZoneName, acc.ISCIDR, name)

}
func testAccCheckIBMISVPNGatewayTaintConfig2(vpc, subnet, name string) string {
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
		name 	= "%s"
		subnet 	= "${ibm_is_subnet.testacc_subnet.id}"
		mode 	= "policy"
		timeouts{
			create = "12m"
		}
		lifecycle {
			create_before_destroy = true
		}
	}`, vpc, subnet, acc.ISZoneName, acc.ISCIDR, name)

}

func TestAccIBMISVPNGateway_taint(t *testing.T) {
	var vpnGateway string
	vpcname := fmt.Sprintf("tfvpnuat-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname := fmt.Sprintf("tfvpnuat-subnet-%d", acctest.RandIntRange(10, 100))
	name1 := fmt.Sprintf("tfvpnuat-taintname-%d", acctest.RandIntRange(10, 100))
	name2 := fmt.Sprintf("tfvpnuat-createname-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMISVPNGatewayDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckIBMISVPNGatewayTaintConfig(vpcname, subnetname, name1),
				ExpectError: regexp.MustCompile(fmt.Sprintf("timeout while waiting for state to become 'done,")),
			},
			{
				Config: testAccCheckIBMISVPNGatewayTaintConfig2(vpcname, subnetname, name2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVPNGatewayExists("ibm_is_vpn_gateway.testacc_vpnGateway", vpnGateway),
					resource.TestCheckResourceAttr(
						"ibm_is_vpn_gateway.testacc_vpnGateway", "name", name2),
				),
			},
		},
	})
}
