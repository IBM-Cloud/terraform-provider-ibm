/* IBM Confidential
*  Object Code Only Source Materials
*  5747-SM3
*  (c) Copyright IBM Corp. 2017,2021
*
*  The source code for this program is not published or otherwise divested
*  of its trade secrets, irrespective of what has been deposited with the
*  U.S. Copyright Office. */

package ibm

import (
	"fmt"
	"github.com/IBM/networking-go-sdk/transitgatewayapisv1"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"log"
	"testing"
)

func TestAccIBMTransitGatewayConnection_basic(t *testing.T) {
	var tgConnection string
	tgConnectionName := fmt.Sprintf("tg-connection-name-%d", acctest.RandIntRange(10, 100))
	gatewayName := fmt.Sprintf("tg-gateway-name-%d", acctest.RandIntRange(10, 100))
	updateVcName := fmt.Sprintf("newtg-connection-name-%d", acctest.RandIntRange(10, 100))
	vpcName := fmt.Sprintf("vpc-name-%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMTransitGatewayConnectionDestroy,
		Steps: []resource.TestStep{

			resource.TestStep{
				//Create test case
				Config: testAccCheckIBMTransitGatewayConnectionConfig(tgConnectionName, gatewayName, vpcName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMTransitGatewayConnectionExists("ibm_tg_connection.test_ibm_tg_connection", tgConnection),
					resource.TestCheckResourceAttr("ibm_tg_connection.test_ibm_tg_connection", "name", tgConnectionName),
				),
			},
			//update
			resource.TestStep{
				Config: testAccCheckIBMTransitGatewayConnectionConfig(updateVcName, gatewayName, vpcName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMTransitGatewayConnectionExists("ibm_tg_connection.test_ibm_tg_connection", tgConnection),
					resource.TestCheckResourceAttr("ibm_tg_connection.test_ibm_tg_connection", "name", updateVcName),
				),
			},
			// tg cross account test
			resource.TestStep{
				//Create test case
				Config: testAccCheckIBMTransitGatewayCrossAccConnectionConfig(tgConnectionName, gatewayName, vpcName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMTransitGatewayConnectionExists("ibm_tg_connection.test_ibm_tg_connection", tgConnection),
					resource.TestCheckResourceAttr("ibm_tg_connection.test_ibm_tg_connection", "name", tgConnectionName),
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
	  `, vpcName, gatewayName, vcName, tg_cross_network_id, tg_cross_network_account_id)

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
func testAccCheckIBMTransitGatewayConnectionDestroy(s *terraform.State) error {
	client, err := transitgatewayClient(testAccProvider.Meta())
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_tg_connection" {
			continue
		}

		parts, err := idParts(rs.Primary.ID)
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
		client, err := transitgatewayClient(testAccProvider.Meta())
		if err != nil {
			return err
		}
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}
		parts, err := idParts(rs.Primary.ID)
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
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMTransitGatewayConnectionDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMTransitGatewayConnectionConfig(tgConnectionName, gatewayname, vpcName),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMTransitGatewayConnectionExists("ibm_tg_connection.test_ibm_tg_connection", virtualConnection),
					resource.TestCheckResourceAttr("ibm_tg_connection.test_ibm_tg_connection", "name", tgConnectionName),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_tg_connection.test_ibm_tg_connection",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"updated_at"},
			},
		},
	})
}
