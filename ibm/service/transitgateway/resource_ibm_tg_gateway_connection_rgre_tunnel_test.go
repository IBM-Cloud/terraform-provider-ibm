// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package transitgateway_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMTransitGatewayConnectionTunnel_basic(t *testing.T) {
	var tgConnection string
	var randNum = acctest.RandIntRange(10, 100)
	connectionName := fmt.Sprintf("tg-connection-tunnel_name-%d", randNum)
	gatewayName := fmt.Sprintf("tg-gateway-name-%d", randNum)
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			// tg RGGRe tunnel creation  test
			{
				//Create test case
				Config: testAccCheckIBMTransitCreateTunnelConfig(gatewayName, connectionName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMTransitGatewayConnectionExists("ibm_tg_connection.test_tg_rgre_connection", tgConnection),
					resource.TestCheckResourceAttr("ibm_tg_connection.test_tg_rgre_connection", "name", connectionName),
				),
			},
		},
	})
}

func testAccCheckIBMTransitCreateTunnelConfig(gatewayName, connectionName string) string {
	return fmt.Sprintf(`	  	
	resource "ibm_tg_gateway" "test_tg_gateway"{
		name="%s"
		location="us-south"
		global=true
	}	
		
	resource "ibm_tg_connection" "test_tg_rgre_connection"{
		gateway = "${ibm_tg_gateway.test_tg_gateway.id}"
		network_type = "redundant_gre"
		name= "%s"
		base_network_type = "classic"
  tunnels {
           local_gateway_ip = "192.189.100.1"
           local_tunnel_ip = "192.118.239.2"
           name =  "tunne1_test1"
           remote_gateway_ip = "10.186.203.4"
           remote_tunnel_ip = "192.118.239.1"
           zone =  "us-south-1"
        }    
 tunnels {
             local_gateway_ip = "192.189.120.1"
             local_tunnel_ip = "192.138.239.2"
             name =  "tunne2_test2"
             remote_gateway_ip = "10.186.203.4"
             remote_tunnel_ip = "192.138.239.1"
             zone =  "us-south-1"
         } 
}
	  `, gatewayName, connectionName)
}
