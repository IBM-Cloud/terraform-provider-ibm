// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package directlink_test

import (
	"fmt"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMDLRouteReportDataSource_basic(t *testing.T) {
	node := "data.ibm_dl_route_report.dl_route_report"
	gatewayname := fmt.Sprintf("gateway-name-%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMDLRouteReportDataSourceConfig(gatewayname),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(node, "status", "complete"),
				),
			},
		},
	})
}

func testAccCheckIBMDLRouteReportDataSourceConfig(gatewayname string) string {
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

	data "ibm_dl_route_report" "dl_route_report" {
		gateway = ibm_dl_gateway.test_dl_gateway.id
		route_report= ibm_dl_route_report.dl_route_report.route_report_id
	}

	  `, gatewayname)
}
