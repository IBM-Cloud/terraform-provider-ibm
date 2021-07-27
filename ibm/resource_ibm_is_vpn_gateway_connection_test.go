// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"errors"
	"fmt"
	"testing"

	"github.com/IBM/vpc-go-sdk/vpcv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIBMISVPNGatewayConnection_basic(t *testing.T) {
	var VPNGatewayConnection string
	vpcname1 := fmt.Sprintf("tfvpngc-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname1 := fmt.Sprintf("tfvpngc-subnet-%d", acctest.RandIntRange(10, 100))
	vpnname1 := fmt.Sprintf("tfvpngc-vpn-%d", acctest.RandIntRange(10, 100))
	name1 := fmt.Sprintf("tfvpngc-createname-%d", acctest.RandIntRange(10, 100))

	vpcname2 := fmt.Sprintf("tfvpngc-vpc-%d", acctest.RandIntRange(10, 100))
	subnetname2 := fmt.Sprintf("tfvpngc-subnet-%d", acctest.RandIntRange(10, 100))
	vpnname2 := fmt.Sprintf("tfvpngc-vpn-%d", acctest.RandIntRange(10, 100))
	name2 := fmt.Sprintf("tfvpngc-createname-%d", acctest.RandIntRange(10, 100))
	updname2 := fmt.Sprintf("tfvpngc-updatename-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMISVPNGatewayConnectionDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMISVPNGatewayConnectionConfig(vpcname1, subnetname1, vpnname1, name1, vpcname2, subnetname2, vpnname2, name2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVPNGatewayConnectionExists("ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", VPNGatewayConnection),
					resource.TestCheckResourceAttr(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "name", name1),
					resource.TestCheckResourceAttrSet(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "gateway_connection"),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMISVPNGatewayConnectionUpdate(vpcname1, subnetname1, vpnname1, name1, vpcname2, subnetname2, vpnname2, updname2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVPNGatewayConnectionExists("ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection2", VPNGatewayConnection),
					resource.TestCheckResourceAttr(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection2", "name", updname2),
				),
			},
		},
	})
}

func TestAccIBMISVPNGatewayConnection_route(t *testing.T) {
	var VPNGatewayConnection string
	vpcname1 := fmt.Sprintf("tfvpngc-vpc-%d", acctest.RandIntRange(100, 200))
	subnetname1 := fmt.Sprintf("tfvpngc-subnet-%d", acctest.RandIntRange(100, 200))
	vpnname1 := fmt.Sprintf("tfvpngc-vpn-%d", acctest.RandIntRange(100, 200))
	name1 := fmt.Sprintf("tfvpngc-createname-%d", acctest.RandIntRange(100, 200))

	vpcname2 := fmt.Sprintf("tfvpngc-vpc-%d", acctest.RandIntRange(100, 200))
	subnetname2 := fmt.Sprintf("tfvpngc-subnet-%d", acctest.RandIntRange(100, 200))
	vpnname2 := fmt.Sprintf("tfvpngc-vpn-%d", acctest.RandIntRange(100, 200))
	name2 := fmt.Sprintf("tfvpngc-createname-%d", acctest.RandIntRange(100, 200))
	updname2 := fmt.Sprintf("tfvpngc-updatename-%d", acctest.RandIntRange(100, 200))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMISVPNGatewayConnectionDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMISVPNGatewayConnectionRouteConfig(vpcname1, subnetname1, vpnname1, name1, vpcname2, subnetname2, vpnname2, name2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVPNGatewayConnectionExists("ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", VPNGatewayConnection),
					resource.TestCheckResourceAttr(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "name", name1),
					resource.TestCheckResourceAttr(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection1", "mode", "route"),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMISVPNGatewayConnectionRouteUpdate(vpcname1, subnetname1, vpnname1, name1, vpcname2, subnetname2, vpnname2, updname2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVPNGatewayConnectionExists("ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection2", VPNGatewayConnection),
					resource.TestCheckResourceAttr(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection2", "name", updname2),
					resource.TestCheckResourceAttr(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection2", "mode", "route"),
				),
			},
		},
	})
}

func testAccCheckIBMISVPNGatewayConnectionDestroy(s *terraform.State) error {

	sess, _ := testAccProvider.Meta().(ClientSession).VpcV1API()
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_is_vpn_gateway_connection" {
			continue
		}
		parts, err := idParts(rs.Primary.ID)
		if err != nil {
			return err
		}

		gID := parts[0]
		gConnID := parts[1]

		getvpngcoptions := &vpcv1.GetVPNGatewayConnectionOptions{
			VPNGatewayID: &gID,
			ID:           &gConnID,
		}
		_, _, err1 := sess.GetVPNGatewayConnection(getvpngcoptions)

		if err1 == nil {
			return fmt.Errorf("VPNGatewayConnection still exists: %s", rs.Primary.ID)
		}
	}

	return nil
}

func testAccCheckIBMISVPNGatewayConnectionExists(n, vpngcID string) resource.TestCheckFunc {
	return func(s *terraform.State) error {

		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return errors.New("No Record ID is set")
		}
		parts, err := idParts(rs.Primary.ID)
		if err != nil {
			return err
		}

		gID := parts[0]
		gConnID := parts[1]

		sess, _ := testAccProvider.Meta().(ClientSession).VpcV1API()
		getvpngcoptions := &vpcv1.GetVPNGatewayConnectionOptions{
			VPNGatewayID: &gID,
			ID:           &gConnID,
		}
		foundvpngcIntf, res, err := sess.GetVPNGatewayConnection(getvpngcoptions)
		if err != nil {
			return fmt.Errorf("Error Getting VPN Gateway connection: %s\n%s", err, res)
		}
		foundvpngc := foundvpngcIntf.(*vpcv1.VPNGatewayConnection)
		vpngcID = *foundvpngc.ID
		return nil
	}
}

func testAccCheckIBMISVPNGatewayConnectionConfig(vpc1, subnet1, vpnname1, name1, vpc2, subnet2, vpnname2, name2 string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc1" {
		name = "%s"
	}

	resource "ibm_is_subnet" "testacc_subnet1" {
		name = "%s"
		vpc = "${ibm_is_vpc.testacc_vpc1.id}"
		zone = "%s"
		ipv4_cidr_block = "%s"
	}
	resource "ibm_is_vpn_gateway" "testacc_VPNGateway1" {
		name = "%s"
		subnet = "${ibm_is_subnet.testacc_subnet1.id}"
	}
	resource "ibm_is_vpn_gateway_connection" "testacc_VPNGatewayConnection1" {
		name = "%s"
		vpn_gateway = "${ibm_is_vpn_gateway.testacc_VPNGateway1.id}"
		peer_address = "${ibm_is_vpn_gateway.testacc_VPNGateway1.public_ip_address}"
		preshared_key = "VPNDemoPassword"
		local_cidrs = ["${ibm_is_subnet.testacc_subnet1.ipv4_cidr_block}"]
		peer_cidrs = ["${ibm_is_subnet.testacc_subnet2.ipv4_cidr_block}"]

	}

	resource "ibm_is_vpc" "testacc_vpc2" {
		name = "%s"
	}

	resource "ibm_is_subnet" "testacc_subnet2" {
		name = "%s"
		vpc = "${ibm_is_vpc.testacc_vpc2.id}"
		zone = "%s"
		ipv4_cidr_block = "%s"
	}

	resource "ibm_is_vpn_gateway" "testacc_VPNGateway2" {
		name = "%s"
		subnet = "${ibm_is_subnet.testacc_subnet2.id}"
		mode = "policy"
	}
	resource "ibm_is_vpn_gateway_connection" "testacc_VPNGatewayConnection2" {
		name = "%s"
		vpn_gateway = "${ibm_is_vpn_gateway.testacc_VPNGateway2.id}"
		peer_address = "${ibm_is_vpn_gateway.testacc_VPNGateway2.public_ip_address}"
		preshared_key = "VPNDemoPassword"
		local_cidrs = ["${ibm_is_subnet.testacc_subnet2.ipv4_cidr_block}"]
		peer_cidrs = ["${ibm_is_subnet.testacc_subnet1.ipv4_cidr_block}"]

	}

	`, vpc1, subnet1, ISZoneName, ISCIDR, vpnname1, name1, vpc2, subnet2, ISZoneName, ISCIDR, vpnname2, name2)

}

func testAccCheckIBMISVPNGatewayConnectionUpdate(vpc1, subnet1, vpnname1, name1, vpc2, subnet2, vpnname2, name2 string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc1" {
		name = "%s"
	}

	resource "ibm_is_subnet" "testacc_subnet1" {
		name = "%s"
		vpc = "${ibm_is_vpc.testacc_vpc1.id}"
		zone = "%s"
		ipv4_cidr_block = "%s"
	}
	resource "ibm_is_vpn_gateway" "testacc_VPNGateway1" {
		name = "%s"
		subnet = "${ibm_is_subnet.testacc_subnet1.id}"
		mode = "policy"
	}
	resource "ibm_is_vpn_gateway_connection" "testacc_VPNGatewayConnection1" {
		name = "%s"
		vpn_gateway = "${ibm_is_vpn_gateway.testacc_VPNGateway1.id}"
		peer_address = "${ibm_is_vpn_gateway.testacc_VPNGateway1.public_ip_address}"
		preshared_key = "VPNDemoPassword"
		local_cidrs = ["${ibm_is_subnet.testacc_subnet1.ipv4_cidr_block}"]
		peer_cidrs = ["${ibm_is_subnet.testacc_subnet2.ipv4_cidr_block}"]

	}

	resource "ibm_is_vpc" "testacc_vpc2" {
		name = "%s"
	}

	resource "ibm_is_subnet" "testacc_subnet2" {
		name = "%s"
		vpc = "${ibm_is_vpc.testacc_vpc2.id}"
		zone = "%s"
		ipv4_cidr_block = "%s"
	}

	resource "ibm_is_vpn_gateway" "testacc_VPNGateway2" {
		name = "%s"
		subnet = "${ibm_is_subnet.testacc_subnet2.id}"
		mode = "policy"
	}
	resource "ibm_is_vpn_gateway_connection" "testacc_VPNGatewayConnection2" {
		name = "%s"
		vpn_gateway = "${ibm_is_vpn_gateway.testacc_VPNGateway2.id}"
		peer_address = "${ibm_is_vpn_gateway.testacc_VPNGateway2.public_ip_address}"
		preshared_key = "VPNDemoPassword"
		local_cidrs = ["${ibm_is_subnet.testacc_subnet2.ipv4_cidr_block}"]
		peer_cidrs = ["${ibm_is_subnet.testacc_subnet1.ipv4_cidr_block}"]

	}

	`, vpc1, subnet1, ISZoneName, ISCIDR, vpnname1, name1, vpc2, subnet2, ISZoneName, ISCIDR, vpnname2, name2)

}

func testAccCheckIBMISVPNGatewayConnectionRouteConfig(vpc1, subnet1, vpnname1, name1, vpc2, subnet2, vpnname2, name2 string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc1" {
		name = "%s"
	}
	resource "ibm_is_subnet" "testacc_subnet1" {
		name = "%s"
		vpc = "${ibm_is_vpc.testacc_vpc1.id}"
		zone = "%s"
		ipv4_cidr_block = "%s"
	}
	resource "ibm_is_vpn_gateway" "testacc_VPNGateway1" {
		name = "%s"
		subnet = "${ibm_is_subnet.testacc_subnet1.id}"
		mode = "route"
	}
	resource "ibm_is_vpn_gateway_connection" "testacc_VPNGatewayConnection1" {
		name = "%s"
		vpn_gateway = "${ibm_is_vpn_gateway.testacc_VPNGateway1.id}"
		peer_address = "${ibm_is_vpn_gateway.testacc_VPNGateway1.public_ip_address}"
		preshared_key = "VPNDemoPassword"
	}
	resource "ibm_is_vpc" "testacc_vpc2" {
		name = "%s"
	}
	resource "ibm_is_subnet" "testacc_subnet2" {
		name = "%s"
		vpc = "${ibm_is_vpc.testacc_vpc2.id}"
		zone = "%s"
		ipv4_cidr_block = "%s"
	}
	resource "ibm_is_vpn_gateway" "testacc_VPNGateway2" {
		name = "%s"
		subnet = "${ibm_is_subnet.testacc_subnet2.id}"
		mode = "route"
	}
	resource "ibm_is_vpn_gateway_connection" "testacc_VPNGatewayConnection2" {
		name = "%s"
		vpn_gateway = "${ibm_is_vpn_gateway.testacc_VPNGateway2.id}"
		peer_address = "${ibm_is_vpn_gateway.testacc_VPNGateway2.public_ip_address}"
		preshared_key = "VPNDemoPassword"
	}
	`, vpc1, subnet1, ISZoneName, ISCIDR, vpnname1, name1, vpc2, subnet2, ISZoneName, ISCIDR, vpnname2, name2)

}

func testAccCheckIBMISVPNGatewayConnectionRouteUpdate(vpc1, subnet1, vpnname1, name1, vpc2, subnet2, vpnname2, name2 string) string {
	return fmt.Sprintf(`
	resource "ibm_is_vpc" "testacc_vpc1" {
		name = "%s"
	}
	resource "ibm_is_subnet" "testacc_subnet1" {
		name = "%s"
		vpc = "${ibm_is_vpc.testacc_vpc1.id}"
		zone = "%s"
		ipv4_cidr_block = "%s"
	}
	resource "ibm_is_vpn_gateway" "testacc_VPNGateway1" {
		name = "%s"
		subnet = "${ibm_is_subnet.testacc_subnet1.id}"
		mode = "route"
	}
	resource "ibm_is_vpn_gateway_connection" "testacc_VPNGatewayConnection1" {
		name = "%s"
		vpn_gateway = "${ibm_is_vpn_gateway.testacc_VPNGateway1.id}"
		peer_address = "${ibm_is_vpn_gateway.testacc_VPNGateway1.public_ip_address}"
		preshared_key = "VPNDemoPassword"
	}
	resource "ibm_is_vpc" "testacc_vpc2" {
		name = "%s"
	}
	resource "ibm_is_subnet" "testacc_subnet2" {
		name = "%s"
		vpc = "${ibm_is_vpc.testacc_vpc2.id}"
		zone = "%s"
		ipv4_cidr_block = "%s"
	}
	resource "ibm_is_vpn_gateway" "testacc_VPNGateway2" {
		name = "%s"
		subnet = "${ibm_is_subnet.testacc_subnet2.id}"
		mode = "route"
	}
	resource "ibm_is_vpn_gateway_connection" "testacc_VPNGatewayConnection2" {
		name = "%s"
		vpn_gateway = "${ibm_is_vpn_gateway.testacc_VPNGateway2.id}"
		peer_address = "${ibm_is_vpn_gateway.testacc_VPNGateway2.public_ip_address}"
		preshared_key = "VPNDemoPassword"
	}
	`, vpc1, subnet1, ISZoneName, ISCIDR, vpnname1, name1, vpc2, subnet2, ISZoneName, ISCIDR, vpnname2, name2)

}
