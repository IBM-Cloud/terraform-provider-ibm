// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package directlink_test

import (
	"fmt"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMDLGatewayDataSource_basic(t *testing.T) {
	node := "data.ibm_dl_gateway.test_dl_gateway"
	gatewayname := fmt.Sprintf("gateway-name-ds-%d", acctest.RandIntRange(10, 100))
	custname := fmt.Sprintf("customer-name-%d", acctest.RandIntRange(10, 100))
	carriername := fmt.Sprintf("carrier-name-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMDLGatewayVCsDataSourceConfig(gatewayname, custname, carriername),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(node, "name", gatewayname),
					// resource.TestCheckResourceAttrSet(node, "as_prepends.#"),
				),
			},
		},
	})
}

func testAccCheckIBMDLGatewayVCsDataSourceConfig(gatewayname, custname, carriername string) string {
	return fmt.Sprintf(`
	data "ibm_dl_routers" "test1" {
		offering_type = "dedicated"
		location_name = "dal10"
	}
	/*
	data "ibm_resource_group" "rg" {
		is_default	= true
	}
	*/
	
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
		//resource_group=data.ibm_resource_group.rg.id
	  }
	   data "ibm_dl_gateway" "test_dl_gateway" {
			name = ibm_dl_gateway.test_dl_gateway.name
		 }
	  `, gatewayname, custname, carriername)
}
