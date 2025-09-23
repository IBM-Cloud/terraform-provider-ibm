// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package transitgateway_test

import (
	"fmt"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMTransitGatewaysDataSource_basic(t *testing.T) {
	resName := "data.ibm_tg_gateways.test1"
	gatewayname := fmt.Sprintf("gateway-name-%d", acctest.RandIntRange(10, 100))
	location := "us-south"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{

			{
				Config: testAccCheckIBMTransitGatewayDataSourceConfig(gatewayname, location),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(
						"data.ibm_tg_gateway.test_tg_gateway", "name", gatewayname),
					resource.TestCheckResourceAttr(
						"data.ibm_tg_gateway.test_tg_gateway", "location", location),
				),
			},
			{
				Config: testAccCheckIBMTransitGatewaysDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(resName, "transit_gateways.0.name"),
					resource.TestCheckResourceAttrSet(resName, "transit_gateways.0.location"),
					resource.TestCheckResourceAttrSet(resName, "transit_gateways.0.global"),
					resource.TestCheckResourceAttrSet(resName, "transit_gateways.0.gre_enhanced_route_propagation"),
				),
			},
		},
	})
}

func testAccCheckIBMTransitGatewaysDataSourceConfig() string {
	// status filter defaults to empty
	return `
      data "ibm_tg_gateways" "test1" {
      }`
}
