// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package transitgateway_test

import (
	"fmt"
	"log"
	"testing"

	"github.com/IBM/networking-go-sdk/transitgatewayapisv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
)

func TestAccIBMTransitGatewayConnection_basic(t *testing.T) {
	var tgConnection string
	tgConnectionName := fmt.Sprintf("tg-connection-name-%d", acctest.RandIntRange(10, 100))
	tgSecondConnectionName := fmt.Sprintf("tg-connection-name-%d", acctest.RandIntRange(10, 100))
	gatewayName := fmt.Sprintf("tg-gateway-name-%d", acctest.RandIntRange(10, 100))
	updateVcName := fmt.Sprintf("newtg-connection-name-%d", acctest.RandIntRange(10, 100))
	vpcName := fmt.Sprintf("vpc-name-%d", acctest.RandIntRange(10, 100))
	vpnConnectionName := fmt.Sprintf("vpn-connection-name-%d", acctest.RandIntRange(10, 100))
	dlGatewayName := fmt.Sprintf("dl-gateway-name-%d", acctest.RandIntRange(10, 100))
	vpnNetworkId := "test-vpn-id"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMTransitGatewayConnectionDestroy,
		Steps: []resource.TestStep{

			{
				//Create test case
				Config: testAccCheckIBMTransitGatewayConnectionConfig(tgConnectionName, gatewayName, vpcName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMTransitGatewayConnectionExists("ibm_tg_connection.test_ibm_tg_connection", tgConnection),
					resource.TestCheckResourceAttr("ibm_tg_connection.test_ibm_tg_connection", "name", tgConnectionName),
				),
			},
			//update
			{
				Config: testAccCheckIBMTransitGatewayConnectionConfig(updateVcName, gatewayName, vpcName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMTransitGatewayConnectionExists("ibm_tg_connection.test_ibm_tg_connection", tgConnection),
					resource.TestCheckResourceAttr("ibm_tg_connection.test_ibm_tg_connection", "name", updateVcName),
				),
			},
			// tg cross account test
			{
				//Create test case
				Config: testAccCheckIBMTransitGatewayCrossAccConnectionConfig(tgConnectionName, gatewayName, vpcName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMTransitGatewayConnectionExists("ibm_tg_connection.test_ibm_tg_connection", tgConnection),
					resource.TestCheckResourceAttr("ibm_tg_connection.test_ibm_tg_connection", "name", tgConnectionName),
				),
			},
			// tg gre tunnel test
			{
				//Create test case
				Config: testAccCheckIBMTransitGatewayGreConnectionConfig(tgConnectionName, gatewayName, tgSecondConnectionName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMTransitGatewayConnectionExists("ibm_tg_connection.test_ibm_tg_gre_connection", tgConnection),
					resource.TestCheckResourceAttr("ibm_tg_connection.test_ibm_tg_gre_connection", "name", tgSecondConnectionName),
				),
			},
			// tg unbound gre test
			{
				//Create test case
				Config: testAccCheckIBMTransitGatewayUnboundGreConnectionConfig(gatewayName, tgConnectionName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMTransitGatewayConnectionExists("ibm_tg_connection.test_ibm_tg_unbound_gre_connection", tgConnection),
					resource.TestCheckResourceAttr("ibm_tg_connection.test_ibm_tg_unbound_gre_connection", "base_network_type", "classic"),
				),
			},
			// tg directlink test
			{
				//Create test case
				Config: testAccCheckIBMTransitGatewayDirectlinkConnectionConfig(dlGatewayName, gatewayName, tgConnectionName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMTransitGatewayConnectionExists("ibm_tg_connection.test_ibm_tg_dl_connection", tgConnection),
					resource.TestCheckResourceAttr("ibm_tg_connection.test_ibm_tg_dl_connection", "name", tgConnectionName),
				),
			},
			// tg power vs test
			{
				//Create test case
				Config: testAccCheckIBMTransitGatewayPowerVSConnectionConfig(gatewayName, tgConnectionName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMTransitGatewayConnectionExists("ibm_tg_connection.test_tg_powervs_connection", tgConnection),
					resource.TestCheckResourceAttr("ibm_tg_connection.test_tg_powervs_connection", "name", tgConnectionName),
				),
			},
			// tg vpn gateway test
			{
				//Create test case
				Config: testAccCheckIBMTransitGatewayVPNGatewayConnectionConfig(gatewayName, vpnNetworkId, vpnConnectionName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMTransitGatewayConnectionExists("ibm_tg_connection.test_tg_vpn_connection", tgConnection),
					resource.TestCheckResourceAttr("ibm_tg_connection.test_tg_vpn_connection", "name", tgConnectionName),
				),
			},
		},
	},
	)
}

func testAccCheckIBMTransitGatewayCrossAccConnectionConfig(vcName, gatewayName, vpcName string) string {
	return fmt.Sprintf(`	
	resource "ibm_is_vpc" "test_tg_vpc" {
		name = "%s"
		}    	
resource "ibm_tg_gateway" "test_tg_gateway"{
		name="%s"
		location="us-south"
		global=true
		}
	 
	
resource "ibm_tg_connection" "test_ibm_tg_connection"{
		gateway = "${ibm_tg_gateway.test_tg_gateway.id}"
		network_type = "vpc"
		name = "%s"
		network_id = "%s"
		network_account_id = "%s"
}	   
	  `, vpcName, gatewayName, vcName, acc.Tg_cross_network_id, acc.Tg_cross_network_account_id)

}

func testAccCheckIBMTransitGatewayConnectionConfig(vcName, gatewayName, vpcName string) string {
	return fmt.Sprintf(`	
	resource "ibm_is_vpc" "test_tg_vpc" {
		name = "%s"
		}    	
resource "ibm_tg_gateway" "test_tg_gateway"{
		name="%s"
		location="us-south"
		global=true
		}
	 
	
resource "ibm_tg_connection" "test_ibm_tg_connection"{
		gateway = "${ibm_tg_gateway.test_tg_gateway.id}"
		network_type = "vpc"
		name= "%s"
		network_id = ibm_is_vpc.test_tg_vpc.resource_crn
}
	   
	  `, vpcName, gatewayName, vcName)

}

func testAccCheckIBMTransitGatewayGreConnectionConfig(gatewayName, classicConnName, greConnName string) string {
	return fmt.Sprintf(`	 	
resource "ibm_tg_gateway" "test_tg_gateway"{
		name="%s"
		location="us-south"
		global=true
		}
	 
	
resource "ibm_tg_connection" "test_ibm_tg_classic_connection"{
		gateway = "${ibm_tg_gateway.test_tg_gateway.id}"
		network_type = "classic"
		name = "%s"
}

resource "ibm_tg_connection" "test_ibm_tg_gre_connection"{
		gateway = "${ibm_tg_gateway.test_tg_gateway.id}"
		network_type = "gre_tunnel"
		name = "%s"
        base_connection_id = "${ibm_tg_connection.test_ibm_tg_classic_connection.connection_id}"
        local_gateway_ip = "192.168.100.1"
        local_tunnel_ip = "192.168.101.1"
        remote_gateway_ip = "10.242.63.12"
        remote_tunnel_ip = "192.168.101.2"
        zone = "us-south-1"
}
	   
	  `, gatewayName, classicConnName, greConnName)
}

func testAccCheckIBMTransitGatewayUnboundGreConnectionConfig(gatewayName, greConnName string) string {
	return fmt.Sprintf(`	 	
resource "ibm_tg_gateway" "test_tg_gateway"{
	name="%s"
	location="us-south"
	global=true
}

resource "ibm_tg_connection" "test_ibm_tg_unbound_gre_connection"{
	gateway = "${ibm_tg_gateway.test_tg_gateway.id}"
	network_type = "unbound_gre_tunnel"
	name = "%s"
	base_network_type = "classic"
	local_gateway_ip = "192.168.100.1"
	local_tunnel_ip = "192.168.101.1"
	network_account_id = "%s"
	remote_gateway_ip = "10.242.63.12"
	remote_tunnel_ip = "192.168.101.2"
	zone = "us-south-1"
}
	  `, gatewayName, greConnName, acc.Tg_cross_network_account_id)
}

func testAccCheckIBMTransitGatewayDirectlinkConnectionConfig(dlGatewayName, gatewayName, dlConnectionName string) string {
	return fmt.Sprintf(`	
data "ibm_dl_ports" "test_dl_ports" {
}
resource "ibm_dl_gateway" "test_dl_gateway"{
		bgp_asn = 64999
		global = true
		name ="%s"
		speed_mbps = 1000
		metered = false
		connection_mode = "transit"
		type = "connect"
		port = data.ibm_dl_ports.test_dl_ports.ports[0].port_id
}

resource "ibm_tg_gateway" "test_tg_gateway"{
		name="%s"
		location="us-south"
		global=true
}
	 
resource "ibm_tg_connection" "test_ibm_tg_dl_connection"{
		gateway = "${ibm_tg_gateway.test_tg_gateway.id}"
		network_type = "directlink"
		name= "%s"
		network_id = "${ibm_dl_gateway.test_dl_gateway.crn}"
}
	  `, dlGatewayName, gatewayName, dlConnectionName)

}
func testAccCheckIBMTransitGatewayPowerVSConnectionConfig(gatewayName, powerVSConnName string) string {
	return fmt.Sprintf(`	   	
resource "ibm_tg_gateway" "test_tg_gateway"{
		name="%s"
		location="us-south"
		global=true
}	 
	
resource "ibm_tg_connection" "test_tg_powervs_connection"{
		gateway = "${ibm_tg_gateway.test_tg_gateway.id}"
		network_type = "power_virtual_server"
		name = "%s"
		network_id = "%s"
}	   
	  `, gatewayName, powerVSConnName, acc.Tg_power_vs_network_id)
}

func testAccCheckIBMTransitGatewayVPNGatewayConnectionConfig(transitGatewayName, vpnNetworkId, vpnGatewayConnectionName string) string {
	return fmt.Sprintf(`	   	
resource "ibm_tg_gateway" "test_tg_gateway"{
		name="%s"
		location="us-south"
		global=true
}	 
	
resource "ibm_tg_connection" "test_tg_vpn_gateway_connection"{
		gateway = "${ibm_tg_gateway.test_tg_gateway.id}"
		network_type = "vpn_gateway"
		name = "%s"
		network_id = "%s"
}	   
	  `, transitGatewayName, vpnGatewayConnectionName, vpnNetworkId)
}

func transitgatewayClient(meta interface{}) (*transitgatewayapisv1.TransitGatewayApisV1, error) {
	sess, err := meta.(conns.ClientSession).TransitGatewayV1API()
	return sess, err
}

func testAccCheckIBMTransitGatewayConnectionDestroy(s *terraform.State) error {
	client, err := transitgatewayClient(acc.TestAccProvider.Meta())
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_tg_connection" {
			continue
		}

		parts, err := flex.IdParts(rs.Primary.ID)
		if err != nil {
			return err
		}
		gatewayId := parts[0]
		ID := parts[1]

		detailTransitGatewayConnectionOptions := &transitgatewayapisv1.GetTransitGatewayConnectionOptions{}
		detailTransitGatewayConnectionOptions.SetTransitGatewayID(gatewayId)
		detailTransitGatewayConnectionOptions.SetID(ID)
		_, _, err = client.GetTransitGatewayConnection(detailTransitGatewayConnectionOptions)
		if err == nil {
			return fmt.Errorf(" transit gateway connection still exists: %s", rs.Primary.ID)
		}
	}
	return nil
}

func testAccCheckIBMTransitGatewayConnectionExists(n string, vc string) resource.TestCheckFunc {
	log.Printf("Inside testAccCheckIBMTransitGatewayConnectionExists :  %s", vc)
	return func(s *terraform.State) error {
		client, err := transitgatewayClient(acc.TestAccProvider.Meta())
		if err != nil {
			return err
		}
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}
		parts, err := flex.IdParts(rs.Primary.ID)
		if err != nil {
			return err
		}
		gatewayId := parts[0]
		ID := parts[1]

		getVCOptions := &transitgatewayapisv1.GetTransitGatewayConnectionOptions{
			ID: &ID,
		}
		getVCOptions.SetTransitGatewayID(gatewayId)
		r, response, err := client.GetTransitGatewayConnection(getVCOptions)
		if err != nil {
			return fmt.Errorf("testAccCheckIBMTransitGatewayConnectionExists: Error Getting Transit Gateway  Connection: %s\n%s", err, response)
		}

		vc = *r.ID
		return nil
	}
}

func testAccCheckIBMTransitGatewayConnectionUpdate(vcName, gatewayName, vpcName string) string {
	return fmt.Sprintf(`

	resource "ibm_is_vpc" "test_dl_vc_vpc" {
		name = "%s"
		}    	
resource "ibm_tg_gateway" "test_tg_gateway"{
		name="%s"
		location="us-south"
		global=true
		}
	 
	
resource "ibm_tg_connection" "test_ibm_tg_connection"{
	depends_on = [ibm_is_vpc.test_dl_vc_vpc,ibm_tg_gateway.test_tg_gateway]
		gateway = "${ibm_tg_gateway.test_tg_gateway.id}"
		network_type = "vpc"
		name= "%s"
		network_id = ibm_is_vpc.test_dl_vc_vpc.resource_crn
}
	`, vpcName, gatewayName, vcName)

}

func TestAccIBMTransitGatewayConnectionImport(t *testing.T) {
	var virtualConnection string
	tgConnectionName := fmt.Sprintf("tg-connection-name-%d", acctest.RandIntRange(10, 100))
	gatewayname := fmt.Sprintf("tg-gateway-name-%d", acctest.RandIntRange(10, 100))
	vpcName := fmt.Sprintf("vpc-name-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMTransitGatewayConnectionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMTransitGatewayConnectionConfig(tgConnectionName, gatewayname, vpcName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMTransitGatewayConnectionExists("ibm_tg_connection.test_ibm_tg_connection", virtualConnection),
					resource.TestCheckResourceAttr("ibm_tg_connection.test_ibm_tg_connection", "name", tgConnectionName),
				),
			},
			{
				ResourceName:      "ibm_tg_connection.test_ibm_tg_connection",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"updated_at"},
			},
		},
	})
}
