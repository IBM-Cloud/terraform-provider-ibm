// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package transitgateway_test

import (
	"fmt"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/networking-go-sdk/transitgatewayapisv1"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"log"
	"testing"
)

func TestAccIBMTransitGatewayRouteReport_basic(t *testing.T) {
	var tgRouteReport string
	gatewayName := fmt.Sprintf("tg-gateway-name-%d", acctest.RandIntRange(10, 100))
	location := fmt.Sprintf("us-south")
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMTransitGatewayRouteReportDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMTransitGatewayRouteReportConfig(gatewayName, location),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMTransitGatewayRouteReportExists("ibm_tg_route_report.test_tg_route", tgRouteReport),
					resource.TestCheckResourceAttrSet("ibm_tg_route_report.test_tg_route", "route_report_id"),
				),
			},
		},
	},
	)
}

func testAccCheckIBMTransitGatewayRouteReportConfig(gatewayname, location string) string {
	return fmt.Sprintf(`
	
	resource "ibm_tg_gateway" "test_tg_gateway" {
		name="%s"
		location="%s"
		global=true
	}

	resource "ibm_tg_route_report" "test_tg_route" {
		gateway = ibm_tg_gateway.test_tg_gateway.id
	}
	`, gatewayname, location)
}

func testAccCheckIBMTransitGatewayRouteReportDestroy(s *terraform.State) error {
	client, err := transitgatewayClient(acc.TestAccProvider.Meta())
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_tg_route_report" {
			continue
		}

		parts, err := flex.IdParts(rs.Primary.ID)
		if err != nil {
			return err
		}
		gatewayId := parts[0]
		ID := parts[1]

		detailTransitGatewayRouteReportOptions := &transitgatewayapisv1.GetTransitGatewayRouteReportOptions{}
		detailTransitGatewayRouteReportOptions.SetTransitGatewayID(gatewayId)
		detailTransitGatewayRouteReportOptions.SetID(ID)
		_, _, err = client.GetTransitGatewayRouteReport(detailTransitGatewayRouteReportOptions)
		if err == nil {
			return fmt.Errorf(" transit gateway route report still exists: %s", rs.Primary.ID)
		}
	}
	return nil
}

func testAccCheckIBMTransitGatewayRouteReportExists(n string, vc string) resource.TestCheckFunc {
	log.Printf("Inside testAccCheckIBMTransitGatewayRouteReportExists :  %s", vc)
	return func(s *terraform.State) error {
		client, err := transitgatewayClient(acc.TestAccProvider.Meta())
		if err != nil {
			return err
		}
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}
		parts, err := flex.IdParts(rs.Primary.ID)
		if err != nil {
			return err
		}
		gatewayId := parts[0]
		ID := parts[1]

		detailTransitGatewayRouteReportOptions := &transitgatewayapisv1.GetTransitGatewayRouteReportOptions{
			ID: &ID,
		}
		detailTransitGatewayRouteReportOptions.SetTransitGatewayID(gatewayId)
		r, response, err := client.GetTransitGatewayRouteReport(detailTransitGatewayRouteReportOptions)
		if err != nil {
			return fmt.Errorf("testAccCheckIBMTransitGatewayRouteReportExists: Error Getting Transit Gateway Route Report: %s\n%s", err, response)
		}

		vc = *r.ID
		return nil
	}
}
