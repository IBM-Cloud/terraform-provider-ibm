// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package directlink_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/networking-go-sdk/directlinkv1"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIBMDLGatewayVC_basic(t *testing.T) {
	var virtualConnection string
	vcName := fmt.Sprintf("vc-name-%d", acctest.RandIntRange(10, 100))
	gatewayname := fmt.Sprintf("gateway-name-%d", acctest.RandIntRange(10, 100))
	custname := fmt.Sprintf("customer-name-%d", acctest.RandIntRange(10, 100))
	carriername := fmt.Sprintf("carrier-name-%d", acctest.RandIntRange(10, 100))
	vctype := "vpc"
	vpcname := fmt.Sprintf("tf-vpcname-%d", acctest.RandIntRange(100, 200))
	updvcName := fmt.Sprintf("vc-name-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMDLGatewayVCDestroy,
		Steps: []resource.TestStep{

			{

				//Create test case
				Config: testAccCheckIBMDLGatewayVCConfig(vctype, vcName, gatewayname, custname, carriername, vpcname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMDLGatewayVCExists("ibm_dl_virtual_connection.test_dl_gateway_vc", virtualConnection),
					resource.TestCheckResourceAttr("ibm_dl_virtual_connection.test_dl_gateway_vc", "name", vcName),
				),
			},
			//update
			{
				Config: testAccCheckIBMDLGatewayVCUpdate(vctype, updvcName, gatewayname, custname, carriername, vpcname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMDLGatewayVCExists("ibm_dl_virtual_connection.test_dl_gateway_vc", virtualConnection),
					resource.TestCheckResourceAttr("ibm_dl_virtual_connection.test_dl_gateway_vc", "name", updvcName),
				),
			},
		},
	},
	)
}

func testAccCheckIBMDLGatewayVCConfig(vctype, vcName, gatewayname, custname, carriername, vpcname string) string {
	return fmt.Sprintf(`	  	
	data "ibm_dl_routers" "test1" {
		offering_type = "dedicated"
		location_name = "dal10"
	}
	resource "ibm_is_vpc" "test_dl_vc_vpc" {
		name = "%s"
		}  
	resource "ibm_dl_gateway" "test_dl_gateway" {
		bgp_asn =  64999
        global = true
        metered = false
        name = "%s"
        speed_mbps = 1000
        type = "dedicated"
		cross_connect_router = data.ibm_dl_routers.test1.cross_connect_routers[0].router_name
        location_name = data.ibm_dl_routers.test1.location_name
	    customer_name = "%s"
        carrier_name = "%s"
	  }
	
	resource "ibm_dl_virtual_connection" "test_dl_gateway_vc"{
		depends_on = [ibm_is_vpc.test_dl_vc_vpc,ibm_dl_gateway.test_dl_gateway]
		gateway = ibm_dl_gateway.test_dl_gateway.id
		name = "%s"
		type = "%s"
		network_id = ibm_is_vpc.test_dl_vc_vpc.resource_crn
	   }
	   
	  `, vpcname, gatewayname, custname, carriername, vcName, vctype)

}

func testAccCheckIBMDLGatewayVCDestroy(s *terraform.State) error {
	directLink, err := directlinkClient(acc.TestAccProvider.Meta())
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_dl_virtual_connection" {
			continue
		}

		parts, err := flex.IdParts(rs.Primary.ID)
		if err != nil {
			return err
		}
		gatewayId := parts[0]
		ID := parts[1]

		getGatewayVirtualConnectionOptions := &directlinkv1.GetGatewayVirtualConnectionOptions{}
		getGatewayVirtualConnectionOptions.SetGatewayID(gatewayId)
		getGatewayVirtualConnectionOptions.SetID(ID)
		_, _, err = directLink.GetGatewayVirtualConnection(getGatewayVirtualConnectionOptions)

		if err == nil {
			return fmt.Errorf("dl connection still exists: %s", rs.Primary.ID)
		}
	}
	return nil
}

func testAccCheckIBMDLGatewayVCExists(n string, vc string) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		directLink, err := directlinkClient(acc.TestAccProvider.Meta())
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

		getVCOptions := &directlinkv1.GetGatewayVirtualConnectionOptions{
			ID: &ID,
		}
		getVCOptions.SetGatewayID(gatewayId)
		r, response, err := directLink.GetGatewayVirtualConnection(getVCOptions)
		if err != nil {
			return fmt.Errorf("testAccCheckIBMDLGatewayVCExists: Error Getting Direct Link Gateway (Dedicated Template) Virtual Connection: %s\n%s", err, response)
		}

		vc = *r.ID
		return nil
	}
}

func testAccCheckIBMDLGatewayVCUpdate(vctype, vcName, gatewayname, custname, carriername, vpcname string) string {
	return fmt.Sprintf(`
	data "ibm_dl_routers" "test1" {
		offering_type = "dedicated"
		location_name = "dal10"
	}
	resource "ibm_is_vpc" "test_dl_vc_vpc" {
		name = "%s"
		}  
	resource "ibm_dl_gateway" "test_dl_gateway" {
		bgp_asn =  64999
        global = true
        metered = false
        name = "%s"
        speed_mbps = 1000
        type = "dedicated"
		cross_connect_router = data.ibm_dl_routers.test1.cross_connect_routers[0].router_name
        location_name = data.ibm_dl_routers.test1.location_name
		customer_name = "%s"
        carrier_name = "%s"
	  }
	
	resource "ibm_dl_virtual_connection" "test_dl_gateway_vc"{
		depends_on = [ibm_is_vpc.test_dl_vc_vpc,ibm_dl_gateway.test_dl_gateway]
		gateway = ibm_dl_gateway.test_dl_gateway.id
		name = "%s"
		type = "%s"
		network_id = ibm_is_vpc.test_dl_vc_vpc.resource_crn
	   }

	`, vpcname, gatewayname, custname, carriername, vcName, vctype)

}

func TestAccIBMDLGatewayVCImport(t *testing.T) {
	var virtualConnection string
	vcName := fmt.Sprintf("vc-name-%d", acctest.RandIntRange(10, 100))
	gatewayname := fmt.Sprintf("gateway-name-%d", acctest.RandIntRange(10, 100))
	custname := fmt.Sprintf("customer-name-%d", acctest.RandIntRange(10, 100))
	carriername := fmt.Sprintf("carrier-name-%d", acctest.RandIntRange(10, 100))
	vctype := "vpc"
	vpcname := fmt.Sprintf("tf-vpcname-%d", acctest.RandIntRange(100, 200))
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMDLGatewayVCDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMDLGatewayVCConfig(vctype, vcName, gatewayname, custname, carriername, vpcname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMDLGatewayVCExists("ibm_dl_virtual_connection.test_dl_gateway_vc", virtualConnection),
					resource.TestCheckResourceAttr("ibm_dl_virtual_connection.test_dl_gateway_vc", "name", vcName),
				),
			},
			{
				ResourceName:      "ibm_dl_virtual_connection.test_dl_gateway_vc",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
