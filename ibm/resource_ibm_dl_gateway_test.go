// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"errors"
	"fmt"
	"log"
	"testing"

	"github.com/IBM/networking-go-sdk/directlinkv1"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
)

func TestAccIBMDLGateway_basic(t *testing.T) {
	var instance string
	gatewayname := fmt.Sprintf("gateway-name-%d", acctest.RandIntRange(10, 100))
	newgatewayname := fmt.Sprintf("newgateway-name-%d", acctest.RandIntRange(10, 100))
	custname := fmt.Sprintf("customer-name-%d", acctest.RandIntRange(10, 100))
	carriername := fmt.Sprintf("carrier-name-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMDLGatewayDestroy, // Delete test case
		Steps: []resource.TestStep{
			{
				//Create test case
				Config: testAccCheckIBMDLGatewayConfig(gatewayname, custname, carriername),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMDLGatewayExists("ibm_dl_gateway.test_dl_gateway", instance),
					resource.TestCheckResourceAttr("ibm_dl_gateway.test_dl_gateway", "name", gatewayname),
					resource.TestCheckResourceAttr("ibm_dl_gateway.test_dl_gateway", "customer_name", custname),
					resource.TestCheckResourceAttr("ibm_dl_gateway.test_dl_gateway", "carrier_name", carriername),
				),
			},
			{
				//Update test case
				Config: testAccCheckIBMDLGatewayConfig(newgatewayname, custname, carriername),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMDLGatewayExists("ibm_dl_gateway.test_dl_gateway", instance),
					resource.TestCheckResourceAttr("ibm_dl_gateway.test_dl_gateway", "name", newgatewayname),
				),
			},
		},
	})
}
func TestAccIBMDLGatewayConnect_basic(t *testing.T) {
	var instance string
	connectgatewayname := fmt.Sprintf("gateway-connect-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMDLGatewayDestroy, // Delete test case
		Steps: []resource.TestStep{

			{
				//dl connect  test case
				Config: testAccCheckIBMDLConnectGatewayConfig(connectgatewayname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMDLGatewayExists("ibm_dl_gateway.test_dl_connect", instance),
					resource.TestCheckResourceAttr("ibm_dl_gateway.test_dl_connect", "name", connectgatewayname),
				),
			},
		},
	})
}

func testAccCheckIBMDLGatewayConfig(gatewayname, custname, carriername string) string {
	return fmt.Sprintf(`
	data "ibm_dl_routers" "test1" {
		offering_type = "dedicated"
		location_name = "dal10"
	}
	  resource "ibm_dl_gateway" "test_dl_gateway" {
		bgp_asn =  64999
        global = true
        metered = false
        name = "%s"
        speed_mbps = 1000
        type =  "dedicated"
		cross_connect_router = data.ibm_dl_routers.test1.cross_connect_routers[0].router_name
        location_name = data.ibm_dl_routers.test1.location_name
		customer_name = "%s"
        carrier_name = "%s"
	  }
	  
	  `, gatewayname, custname, carriername)
}

func testAccCheckIBMDLConnectGatewayConfig(gatewayname string) string {
	return fmt.Sprintf(`
	data "ibm_dl_ports" "test_ds_dl_ports" {
	}
	  resource "ibm_dl_gateway" "test_dl_connect" {
		bgp_asn =  64999
        global = true
        metered = false
        name = "%s"
        speed_mbps = 1000
		type =  "connect"
		port =  data.ibm_dl_ports.test_ds_dl_ports.ports[0].port_id
	}
	  `, gatewayname)
}

func testAccCheckIBMDLGatewayExists(n string, instance string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}
		if rs.Primary.ID == "" {
			return errors.New("No Record ID is set")
		}
		directLink, err := directlinkClient(testAccProvider.Meta())
		if err != nil {
			return err
		}
		getOptions := &directlinkv1.GetGatewayOptions{
			ID: &rs.Primary.ID,
		}
		instance1, response, err := directLink.GetGateway(getOptions)
		if err != nil {
			return fmt.Errorf("Error Getting Direct Link Gateway (Dedicated Template): %s\n%s", err, response)
		}
		instance = *instance1.ID
		return nil
	}
}

func testAccCheckIBMDLGatewayDestroy(s *terraform.State) error {
	directLink, err := directlinkClient(testAccProvider.Meta())
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_dl_gateway" {
			log.Printf("Destroy called ...%s", rs.Primary.ID)
			getOptions := &directlinkv1.GetGatewayOptions{
				ID: &rs.Primary.ID,
			}
			_, _, err = directLink.GetGateway(getOptions)

			if err == nil {
				return fmt.Errorf("gateway still exists: %s", rs.Primary.ID)
			}
		}
	}
	return nil
}
