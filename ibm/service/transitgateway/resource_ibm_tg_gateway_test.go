// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package transitgateway_test

import (
	"errors"
	"fmt"
	"log"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/IBM/networking-go-sdk/transitgatewayapisv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccIBMTransitGateway_basic(t *testing.T) {
	var instance string
	gatewayname := fmt.Sprintf("tg-gateway-name-%d", acctest.RandIntRange(10, 100))
	newgatewayname := fmt.Sprintf("newgateway-name-%d", acctest.RandIntRange(10, 100))
	location := "us-south"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMTransitGatewayDestroy, // Delete test case
		Steps: []resource.TestStep{
			{
				//Create test case
				Config: testAccCheckIBMTransitGatewayConfig(gatewayname, location),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMTransitGatewayExists("ibm_tg_gateway.test_tg_gateway", instance),
					resource.TestCheckResourceAttr("ibm_tg_gateway.test_tg_gateway", "name", gatewayname),
					resource.TestCheckResourceAttr("ibm_tg_gateway.test_tg_gateway", "location", location),
				),
			},
			{
				//Update test case
				Config: testAccCheckIBMTransitGatewayConfig(newgatewayname, location),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMTransitGatewayExists("ibm_tg_gateway.test_tg_gateway", instance),
					resource.TestCheckResourceAttr("ibm_tg_gateway.test_tg_gateway", "name", newgatewayname),
				),
			},
		},
	})
}

func testAccCheckIBMTransitGatewayConfig(gatewayname, location string) string {
	return fmt.Sprintf(`
	  
	resource "ibm_tg_gateway" "test_tg_gateway"{
		name="%s"
		location="%s"
		global=true
		}
	  `, gatewayname, location)
}

func testAccCheckIBMTransitGatewayExists(n string, instance string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}
		if rs.Primary.ID == "" {
			return errors.New("No Record ID is set")
		}
		client, err := transitgatewayClient(acc.TestAccProvider.Meta())
		if err != nil {
			return err
		}
		tgOptions := &transitgatewayapisv1.GetTransitGatewayOptions{
			ID: &rs.Primary.ID,
		}
		instance1, response, err := client.GetTransitGateway(tgOptions)
		if err != nil {
			return fmt.Errorf("[ERROR] Error Getting Transit Gateway: %s\n%s", err, response)
		}
		instance = *instance1.ID
		return nil
	}
}

func testAccCheckIBMTransitGatewayDestroy(s *terraform.State) error {
	client, err := transitgatewayClient(acc.TestAccProvider.Meta())
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_tg_gateway" {
			log.Printf("Destroy called ...%s", rs.Primary.ID)

			tgOptions := &transitgatewayapisv1.GetTransitGatewayOptions{
				ID: &rs.Primary.ID,
			}
			_, _, err = client.GetTransitGateway(tgOptions)

			if err == nil {
				return fmt.Errorf(" transit gateway still exists: %s", rs.Primary.ID)
			}
		}
	}
	return nil
}

func TestAccIBMTransitGatewayImport(t *testing.T) {
	var instance string
	gatewayname := fmt.Sprintf("gateway-name-%d", acctest.RandIntRange(10, 100))
	location := "us-south"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMTransitGatewayDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMTransitGatewayConfig(gatewayname, location),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMTransitGatewayExists("ibm_tg_gateway.test_tg_gateway", instance),
					resource.TestCheckResourceAttr("ibm_tg_gateway.test_tg_gateway", "name", gatewayname),
					resource.TestCheckResourceAttr("ibm_tg_gateway.test_tg_gateway", "location", location),
				),
			},
			{
				ResourceName:      "ibm_tg_gateway.test_tg_gateway",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
