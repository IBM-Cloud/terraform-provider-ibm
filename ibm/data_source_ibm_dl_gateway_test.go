/**
 * © Copyright IBM Corporation 2020. All Rights Reserved.
 *
 * Licensed under the Mozilla Public License, version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at https://mozilla.org/MPL/2.0/
 */

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
	   data "ibm_dl_gateway" "test_dl_gateway_vc" {
			name = ibm_dl_gateway.test_dl_gateway.name
		 }
	  `, gatewayname, custname, carriername)
}
