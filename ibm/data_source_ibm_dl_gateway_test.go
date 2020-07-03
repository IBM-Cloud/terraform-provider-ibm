package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMDLGatewayDataSource_basic(t *testing.T) {
	node := "data.ibm_dl_gateway.test_dl_gateway_vc"
	gatewayname := fmt.Sprintf("gateway-name-%d", acctest.RandIntRange(10, 100))
	custname := fmt.Sprintf("customer-name-%d", acctest.RandIntRange(10, 100))
	carriername := fmt.Sprintf("carrier-name-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMDLGatewayVCsDataSourceConfig(gatewayname, custname, carriername),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(node, "name", gatewayname),
				),
			},
		},
	})
}

func testAccCheckIBMDLGatewayVCsDataSourceConfig(gatewayname, custname, carriername string) string {
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
	   data "ibm_dl_gateway" "test_dl_gateway_vc" {
			name = ibm_dl_gateway.test_dl_gateway.name
		 }
	  `, gatewayname, custname, carriername)
}
