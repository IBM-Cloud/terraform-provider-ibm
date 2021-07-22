// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMAtrackerRoutesDataSourceBasic(t *testing.T) {
	routeName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	routeReceiveGlobalEvents := "false"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMAtrackerRoutesDataSourceConfigBasic(routeName, routeReceiveGlobalEvents),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_atracker_routes.atracker_routes", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_atracker_routes.atracker_routes", "routes.#"),
					resource.TestCheckResourceAttr("data.ibm_atracker_routes.atracker_routes", "routes.0.name", routeName),
					resource.TestCheckResourceAttr("data.ibm_atracker_routes.atracker_routes", "routes.0.receive_global_events", routeReceiveGlobalEvents),
				),
			},
		},
	})
}

func testAccCheckIBMAtrackerRoutesDataSourceConfigBasic(routeName string, routeReceiveGlobalEvents string) string {
	return fmt.Sprintf(`
		resource "ibm_atracker_route" "atracker_route" {
			name = "%s"
			receive_global_events = %s
			rules {
				target_ids = ["target_ids"]
			}
		}

		data "ibm_atracker_routes" "atracker_routes" {
			name = ibm_atracker_route.atracker_route.name
		}
	`, routeName, routeReceiveGlobalEvents)
}
