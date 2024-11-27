// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package directlink_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIBMDLGatewaysDataSource_basic(t *testing.T) {
	var instance string
	node := "data.ibm_dl_gateways.test1"
	gatewayname := fmt.Sprintf("gateway-name-dgws-%d", acctest.RandIntRange(10, 100))
	custname := fmt.Sprintf("customer-name-%d", acctest.RandIntRange(10, 100))
	carriername := fmt.Sprintf("carrier-name-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMDLGatewaysDataSourceConfig(gatewayname, custname, carriername),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMDLGatewayExists("ibm_dl_gateway.test_dl_gateway", instance),
					resource.TestCheckResourceAttrSet(node, "gateways.#"),
				),
			},
		},
	})
}

func testAccCheckIBMDLGatewaysDataSourceConfig(gatewayname, custname, carriername string) string {
	return fmt.Sprintf(`
	data "ibm_dl_routers" "test1" {
		offering_type = "dedicated"
		location_name = "dal10"
	}

	data "ibm_resource_group" "rg" {
		is_default	= true
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
		resource_group=data.ibm_resource_group.rg.id
	  }
	 
	  data "ibm_dl_gateways" "test1" {
	}
	  
	  `, gatewayname, custname, carriername)
}
