// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package metricsrouter_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMMetricsRouterRoutesDataSourceBasic(t *testing.T) {
	routeName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMMetricsRouterRoutesDataSourceConfigBasic(routeName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_metrics_router_routes.metrics_router_routes_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_metrics_router_routes.metrics_router_routes_instance", "routes.#"),
					resource.TestCheckResourceAttr("data.ibm_metrics_router_routes.metrics_router_routes_instance", "routes.0.name", routeName),
					resource.TestCheckResourceAttr("data.ibm_metrics_router_routes.metrics_router_routes_instance", "routes.0.rules.0.inclusion_filters.0.operand", "location"),
					resource.TestCheckResourceAttr("data.ibm_metrics_router_routes.metrics_router_routes_instance", "routes.0.rules.0.inclusion_filters.0.operator", "is"),
					resource.TestCheckResourceAttr("data.ibm_metrics_router_routes.metrics_router_routes_instance", "routes.0.rules.0.inclusion_filters.0.value.0", "value"),
				),
			},
		},
	})
}

func testAccCheckIBMMetricsRouterRoutesDataSourceConfigBasic(routeName string) string {
	return fmt.Sprintf(`
		resource "ibm_metrics_router_target" "metrics_router_target_instance" {
			name = "my-mr-target"
			destination_crn = "crn:v1:bluemix:public:sysdig-monitor:us-south:a/0be5ad401ae913d8ff665d92680664ed:22222222-2222-2222-2222-222222222222::"
		}

		resource "ibm_metrics_router_route" "metrics_router_route_instance" {
			name = "%s"
			rules {
				target_ids = [ ibm_metrics_router_target.metrics_router_target_instance.id ]
				inclusion_filters {
					operand = "location"
					operator = "is"
					value = [ "value" ]
				}
			}
		}

		data "ibm_metrics_router_routes" "metrics_router_routes_instance" {
			name = ibm_metrics_router_route.metrics_router_route_instance.name
		}
	`, routeName)
}
