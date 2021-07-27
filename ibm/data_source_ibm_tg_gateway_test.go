// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMTransitGatewayDataSource_basic(t *testing.T) {
	gatewayname := fmt.Sprintf("gateway-name-%d", acctest.RandIntRange(10, 100))
	location := fmt.Sprintf("us-south")

	classicConnName := fmt.Sprintf("classic-connection-name-%d", acctest.RandIntRange(10, 100))
	greConnName := fmt.Sprintf("gre-connection-name-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMTransitGatewayDataSourceConfig(gatewayname, location),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.ibm_tg_gateway.test_tg_gateway", "name", gatewayname),
					resource.TestCheckResourceAttr(
						"data.ibm_tg_gateway.test_tg_gateway", "location", location),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMTransitGatewayGreConnectionDataSourceConfig(gatewayname, location, classicConnName, greConnName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_tg_gateway.test_tg_gateway", "connections.#"),
				),
			},
		},
	})
}

func testAccCheckIBMTransitGatewayDataSourceConfig(gatewayname, location string) string {
	return fmt.Sprintf(`
	
	resource "ibm_tg_gateway" "test_tg_gateway" {
		name="%s"
		location="%s"
		global=true
	  }
	
	   data "ibm_tg_gateway" "test_tg_gateway" {
			name = ibm_tg_gateway.test_tg_gateway.name
		 }
	  `, gatewayname, location)
}

func testAccCheckIBMTransitGatewayGreConnectionDataSourceConfig(gatewayName, location, classicConnName, greConnName string) string {
	return fmt.Sprintf(`
    
    resource "ibm_tg_gateway" "test_tg_gateway" {
        name="%s"
        location="%s"
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
    
       data "ibm_tg_gateway" "test_tg_gateway" {
            name = ibm_tg_gateway.test_tg_gateway.name
         }
      `, gatewayName, location, classicConnName, greConnName)
}
