package ibm

import (
	"errors"
	"fmt"
	"log"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.ibm.com/ibmcloud/networking-go-sdk/transitgatewayapisv1"
)

func TestAccIBMTransitGateway_basic(t *testing.T) {
	var instance string
	gatewayname := fmt.Sprintf("tg-gateway-name-%d", acctest.RandIntRange(10, 100))
	newgatewayname := fmt.Sprintf("newgateway-name-%d", acctest.RandIntRange(10, 100))
	location := fmt.Sprintf("us-south")

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
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
		client, err := transitgatewayClient(testAccProvider.Meta())
		if err != nil {
			return err
		}
		tgOptions := &transitgatewayapisv1.DetailTransitGatewayOptions{
			ID: &rs.Primary.ID,
		}
		instance1, response, err := client.DetailTransitGateway(tgOptions)
		if err != nil {
			return fmt.Errorf("Error Getting Transit Gateway: %s\n%s", err, response)
		}
		instance = *instance1.ID
		return nil
	}
}

func testAccCheckIBMTransitGatewayDestroy(s *terraform.State) error {
	client, err := transitgatewayClient(testAccProvider.Meta())
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_tg_gateway" {
			log.Printf("Destroy called ...%s", rs.Primary.ID)
			delOptions := &transitgatewayapisv1.DeleteTransitGatewayOptions{
				ID: &rs.Primary.ID,
			}
			response, err := client.DeleteTransitGateway(delOptions)
			if err != nil && response.StatusCode != 404 {
				log.Printf("Error deleting transit gateway :%s", response)
				return err
			}
			tgOptions := &transitgatewayapisv1.DetailTransitGatewayOptions{
				ID: &rs.Primary.ID,
			}
			_, response, err = client.DetailTransitGateway(tgOptions)

			if err == nil {
				return fmt.Errorf(" tarnsit gateway still exists: %s", rs.Primary.ID)
			}
		}
	}
	return nil
}

func TestAccIBMTransitGatewayImport(t *testing.T) {
	var instance string
	gatewayname := fmt.Sprintf("gateway-name-%d", acctest.RandIntRange(10, 100))
	location := fmt.Sprintf("us-south")
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMTransitGatewayDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMTransitGatewayConfig(gatewayname, location),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMTransitGatewayExists("ibm_tg_gateway.test_tg_gateway", instance),
					resource.TestCheckResourceAttr("ibm_tg_gateway.test_tg_gateway", "name", gatewayname),
					resource.TestCheckResourceAttr("ibm_tg_gateway.test_tg_gateway", "location", location),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_tg_gateway.test_tg_gateway",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
