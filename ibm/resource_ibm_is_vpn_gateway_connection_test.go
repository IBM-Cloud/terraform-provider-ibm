package ibm

import (
	"errors"
	"fmt"
	"testing"

	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"github.ibm.com/Bluemix/riaas-go-client/clients/vpn"
	"github.ibm.com/riaas/rias-api/riaas/models"
)

func TestAccIBMISVPNGatewayConnection_basic(t *testing.T) {
	var VPNGatewayConnection *models.VPNGatewayConnection
	vpcname1 := fmt.Sprintf("terraformvpnuatConn_vpc_%d", acctest.RandInt())
	subnetname1 := fmt.Sprintf("terraformvpnuatConn_subnet_%d", acctest.RandInt())
	vpnname1 := fmt.Sprintf("terraformvpnuatConn_vpn_%d", acctest.RandInt())
	name1 := fmt.Sprintf("terraformvpnconnuat_create_step_name_%d", acctest.RandInt())

	vpcname2 := fmt.Sprintf("terraformvpnuatConn_vpc_%d", acctest.RandInt())
	subnetname2 := fmt.Sprintf("terraformvpnuatConn_subnet_%d", acctest.RandInt())
	vpnname2 := fmt.Sprintf("terraformvpnuatConn_vpn_%d", acctest.RandInt())
	name2 := fmt.Sprintf("terraformvpnconnuat_create_step_name_%d", acctest.RandInt())

	updname2 := fmt.Sprintf("terraformvpnconnuat_update_step_name_%d", acctest.RandInt())

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMISVPNGatewayConnectionDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMISVPNGatewayConnectionConfig(vpcname1, subnetname1, vpnname1, name1, vpcname2, subnetname2, vpnname2, name2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVPNGatewayConnectionExists("ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection", &VPNGatewayConnection),
					resource.TestCheckResourceAttr(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection", "name", name1),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMISVPNGatewayConnectionUpdate(vpcname1, subnetname1, vpnname1, name1, vpcname2, subnetname2, vpnname2, updname2),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMISVPNGatewayConnectionExists("ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection", &VPNGatewayConnection),
					resource.TestCheckResourceAttr(
						"ibm_is_vpn_gateway_connection.testacc_VPNGatewayConnection", "name", updname2),
				),
			},
		},
	})
}

func testAccCheckIBMISVPNGatewayConnectionDestroy(s *terraform.State) error {
	sess, _ := testAccProvider.Meta().(ClientSession).ISSession()

	VPNGatewayConnectionC := vpn.NewVpnClient(sess)

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

		_, err = VPNGatewayConnectionC.GetConnection(gID, gConnID)

		if err == nil {
			return fmt.Errorf("VPNGatewayConnection still exists: %s", rs.Primary.ID)
		}
	}

	return nil
}

func testAccCheckIBMISVPNGatewayConnectionExists(n string, VPNGatewayConnection **models.VPNGatewayConnection) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return errors.New("No Record ID is set")
		}

		sess, _ := testAccProvider.Meta().(ClientSession).ISSession()
		VPNGatewayConnectionC := vpn.NewVpnClient(sess)

		parts, err := idParts(rs.Primary.ID)
		if err != nil {
			return err
		}
		gID := parts[0]
		gConnID := parts[1]

		foundVPNGatewayConnection, err := VPNGatewayConnectionC.GetConnection(gID, gConnID)
		if err != nil {
			return err
		}

		*VPNGatewayConnection = foundVPNGatewayConnection
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
