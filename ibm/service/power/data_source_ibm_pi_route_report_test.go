// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIBMPIRouteReportDataSource_basic(t *testing.T) {
	routeReportData := "data.ibm_pi_route_report.route_report_data"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPIRouteReportDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(routeReportData, "id"),
				),
			},
		},
	})
}

func testAccCheckIBMPIRouteReportDataSourceConfig() string {
	return fmt.Sprintf(`
		data "ibm_pi_route_report" "route_report_data" {
			pi_cloud_instance_id = "%s"
		}`, acc.Pi_cloud_instance_id)
}
