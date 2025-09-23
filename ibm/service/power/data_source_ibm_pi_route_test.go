// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package power_test

import (
	"fmt"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMPIRouteDataSource_basic(t *testing.T) {
	routeData := "data.ibm_pi_route.route_data"
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPIRouteDataSourceConfig(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet(routeData, "id"),
				),
			},
		},
	})
}

func testAccCheckIBMPIRouteDataSourceConfig() string {
	return fmt.Sprintf(`
		data "ibm_pi_route" "route_data" {
			pi_cloud_instance_id = "%[2]s"
			pi_route_id          = "%[1]s"
		}`, acc.Pi_route_id, acc.Pi_cloud_instance_id)
}
