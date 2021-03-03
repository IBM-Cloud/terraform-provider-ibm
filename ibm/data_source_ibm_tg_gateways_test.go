// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMTransitGatewaysDataSource_basic(t *testing.T) {
	resName := "data.ibm_tg_gateways.test1"
	gatewayname := fmt.Sprintf("gateway-name-%d", acctest.RandIntRange(10, 100))
	location := fmt.Sprintf("us-south")

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
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
				),
			},
		},
	})
}

func testAccCheckIBMTransitGatewaysDataSourceConfig() string {
	// status filter defaults to empty
	return fmt.Sprintf(`
      data "ibm_tg_gateways" "test1" {
      }`)
}
