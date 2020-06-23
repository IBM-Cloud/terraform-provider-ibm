package ibm

import (
	"errors"
	"fmt"
	"log"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.ibm.com/ibmcloud/networking-go-sdk/directlinkapisv1"
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

func testAccCheckIBMDLGatewayConfig(gatewayname, custname, carriername string) string {
	return fmt.Sprintf(`
	  
	  resource "ibm_dl_gateway" "test_dl_gateway" {
		bgp_asn =  64999
        bgp_base_cidr =  "169.254.0.0/16"
        global = true
        metered = false
        name = "%s"
        speed_mbps = 1000
        type =  "dedicated"
        cross_connect_router = "LAB-xcr01.dal09"
        location_name = "dal09"
        customer_name = "%s"
        carrier_name = "%s"
	  }
	  
	  `, gatewayname, custname, carriername)
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
		getOptions := &directlinkapisv1.GetGatewayOptions{
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
			delOptions := &directlinkapisv1.DeleteGatewayOptions{
				ID: &rs.Primary.ID,
			}
			response, err := directLink.DeleteGateway(delOptions)
			if err != nil && response.StatusCode != 404 {
				log.Printf("Error deleting direct link gateway dedicated:%s", response)
				return err
			}
			getOptions := &directlinkapisv1.GetGatewayOptions{
				ID: &rs.Primary.ID,
			}
			_, response, err = directLink.GetGateway(getOptions)

			if err == nil {
				return fmt.Errorf("gateway still exists: %s", rs.Primary.ID)
			}
		}
	}
	return nil
}
