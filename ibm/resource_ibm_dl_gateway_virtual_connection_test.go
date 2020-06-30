package ibm

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"github.ibm.com/ibmcloud/networking-go-sdk/directlinkapisv1"
	"log"
	"testing"
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
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMDLGatewayVCDestroy,
		Steps: []resource.TestStep{

			resource.TestStep{

				//Create test case
				Config: testAccCheckIBMDLGatewayVCConfig(vctype, vcName, gatewayname, custname, carriername, vpcname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMDLGatewayVCExists("ibm_dl_virtual_connection.test_dl_gateway_vc", virtualConnection),
					resource.TestCheckResourceAttr("ibm_dl_virtual_connection.test_dl_gateway_vc", "name", vcName),
				),
			},
			//update
			resource.TestStep{
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
	
	resource "ibm_is_vpc" "test_dl_vc_vpc" {
		name = "%s"
		}  
	resource "ibm_dl_gateway" "test_dl_gateway" {
		bgp_asn =  64999
        bgp_base_cidr =  "169.254.0.0/16"
        global = true
        metered = false
        name = "%s"
        speed_mbps = 1000
        type = "dedicated"
        cross_connect_router = "LAB-xcr01.dal09"
        location_name = "dal09"
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
	directLink, err := directlinkClient(testAccProvider.Meta())
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_dl_virtual_connection" {
			continue
		}

		parts, err := idParts(rs.Primary.ID)
		if err != nil {
			return err
		}
		gatewayId := parts[0]
		ID := parts[1]

		delVCOptions := &directlinkapisv1.DeleteGatewayVirtualConnectionOptions{
			ID: &ID,
		}
		delVCOptions.SetGatewayID(gatewayId)
		response, err := directLink.DeleteGatewayVirtualConnection(delVCOptions)

		if err != nil && response.StatusCode != 404 {
			log.Printf("testAccCheckIBMDLGatewayVCDestroy:Error deleting Direct Link Gateway (Dedicated Template) Virtual Connection: %s", response)
			return err
		}
	}
	return nil
}

func testAccCheckIBMDLGatewayVCExists(n string, vc string) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		directLink, err := directlinkClient(testAccProvider.Meta())
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

		getVCOptions := &directlinkapisv1.GetGatewayVirtualConnectionOptions{
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
	resource "ibm_is_vpc" "test_dl_vc_vpc" {
		name = "%s"
		}  
	resource "ibm_dl_gateway" "test_dl_gateway" {
		bgp_asn =  64999
        bgp_base_cidr =  "169.254.0.0/16"
        global = true
        metered = false
        name = "%s"
        speed_mbps = 1000
        type = "dedicated"
        cross_connect_router = "LAB-xcr01.dal09"
        location_name = "dal09"
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
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMDLGatewayVCDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMDLGatewayVCConfig(vctype, vcName, gatewayname, custname, carriername, vpcname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMDLGatewayVCExists("ibm_dl_virtual_connection.test_dl_gateway_vc", virtualConnection),
					resource.TestCheckResourceAttr("ibm_dl_virtual_connection.test_dl_gateway_vc", "name", vcName),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_dl_virtual_connection.test_dl_gateway_vc",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
