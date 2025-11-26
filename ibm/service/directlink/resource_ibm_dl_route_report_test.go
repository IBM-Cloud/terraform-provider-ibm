// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package directlink_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/networking-go-sdk/directlinkv1"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccIBMDLRouteReportResource_basic(t *testing.T) {
	var dlRouteReport string
	node := "ibm_dl_route_report.dl_route_report"
	gatewayname := fmt.Sprintf("gateway-name-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMDLRouteReportDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMDLRouteReportResourceConfig(gatewayname),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIBMDLGatewayRouteReportExists(node, dlRouteReport),
					resource.TestCheckResourceAttrSet(node, "route_report_id"),
				),
			},
		},
	})
}

func testAccCheckIBMDLRouteReportResourceConfig(gatewayname string) string {
	return fmt.Sprintf(`
	data "ibm_dl_ports" "ds_dlports" {
	}
	
	resource ibm_dl_gateway test_dl_gateway {
		bgp_asn =  64999
		global = true 
		metered = false
		name = "%s"
		speed_mbps = 1000 
		type =  "connect" 
		port = data.ibm_dl_ports.ds_dlports.ports[0].port_id
	} 

	resource ibm_dl_route_report dl_route_report {
		gateway = ibm_dl_gateway.test_dl_gateway.id
	 }

	  `, gatewayname)
}

func testAccCheckIBMDLRouteReportDestroy(s *terraform.State) error {
	client, err := directlinkClient(acc.TestAccProvider.Meta())
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_dl_route_report" {
			continue
		}

		parts, err := flex.IdParts(rs.Primary.ID)
		if err != nil {
			return err
		}
		gatewayId := parts[0]
		ID := parts[1]

		getGatewayRouteReportOptions := &directlinkv1.GetGatewayRouteReportOptions{}
		getGatewayRouteReportOptions.SetGatewayID(gatewayId)
		getGatewayRouteReportOptions.SetID(ID)
		_, _, err = client.GetGatewayRouteReport(getGatewayRouteReportOptions)
		if err == nil {
			return fmt.Errorf(" DL gateway route report still exists: %s", rs.Primary.ID)
		}
	}
	return nil
}

func testAccCheckIBMDLGatewayRouteReportExists(n string, vc string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		client, err := directlinkClient(acc.TestAccProvider.Meta())
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

		getGatewayRouteReportOptions := &directlinkv1.GetGatewayRouteReportOptions{
			ID:        &ID,
			GatewayID: &gatewayId,
		}
		r, response, err := client.GetGatewayRouteReport(getGatewayRouteReportOptions)
		if err != nil {
			return fmt.Errorf("testAccCheckIBMDLGatewayRouteReportExists: Error Getting DL Gateway Route Report: %s\n%s", err, response)
		}

		vc = *r.ID
		return nil
	}
}
