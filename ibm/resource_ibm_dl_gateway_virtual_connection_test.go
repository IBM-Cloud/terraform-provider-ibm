package ibm

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"testing"
)

func TestAccIBMDLGatewayVC_basic(t *testing.T) {
	//var instance string
	vcName := fmt.Sprintf("vc-name-%d", acctest.RandIntRange(10, 100))
	gatewayname := fmt.Sprintf("gateway-name-%d", acctest.RandIntRange(10, 100))
	custname := fmt.Sprintf("customer-name-%d", acctest.RandIntRange(10, 100))
	carriername := fmt.Sprintf("carrier-name-%d", acctest.RandIntRange(10, 100))
	vctype := "vpc"
	vpcname := fmt.Sprintf("tf-vpcname-%d", acctest.RandIntRange(100, 200))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				//Create test case
				Config: testAccCheckIBMDLGatewayVCConfig(vctype, vcName, gatewayname, custname, carriername, vpcname),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_dl_gateway_vc.test_dl_gateway_vc", "name", vcName),
				),
			},
		},
	})
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
	
	resource "ibm_dl_gateway_vc" "test_dl_gateway_vc"{
		depends_on = [ibm_is_vpc.test_dl_vc_vpc,ibm_dl_gateway.test_dl_gateway]
		gateway = ibm_dl_gateway.test_dl_gateway.id
		name = "%s"
		type = "%s"
		network_id = ibm_is_vpc.test_dl_vc_vpc.resource_crn
	   }

	   data "ibm_dl_gateway_virtualconnections" "test_dl_gateway_virtualconnections" {
			gateway = ibm_dl_gateway.test_dl_gateway.id
		 }
	  `, vpcname, gatewayname, custname, carriername, vcName, vctype)

}
