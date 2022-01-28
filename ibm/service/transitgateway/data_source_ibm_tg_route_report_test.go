// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package transitgateway_test

import (
	"fmt"
	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccIBMTransitGatewayRouteReportDataSource_basic(t *testing.T) {
	gatewayname := fmt.Sprintf("gateway-name-%d", acctest.RandIntRange(10, 100))
	location := fmt.Sprintf("us-south")

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMTransitGatewayDataRouteReportSourceConfig(gatewayname, location),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_tg_route_report.test_tg_route_get", "connections.#"),
				),
			},
		},
	})
}

func testAccCheckIBMTransitGatewayDataRouteReportSourceConfig(gatewayname, location string) string {
	return fmt.Sprintf(`
	
	resource "ibm_tg_gateway" "test_tg_gateway" {
		name="%s"
		location="%s"
		global=true
	}

	resource "ibm_tg_route_report" "test_tg_route" {
		gateway = ibm_tg_gateway.test_tg_gateway.id
	}

	data "ibm_tg_route_report" "test_tg_route_get" {
		gateway = ibm_tg_gateway.test_tg_gateway.id
		route_report = ibm_tg_route_report.test_tg_route.route_report_id
	}
	`, gatewayname, location)
}
