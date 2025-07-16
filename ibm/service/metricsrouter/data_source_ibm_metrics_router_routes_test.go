// Copyright IBM Corp. 2023 All Rights Reserved.
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
			resource.TestStep{
				Config: testAccCheckIBMMetricsRouterRoutesDataSourceConfigBasic(routeName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_metrics_router_routes.metrics_router_routes_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_metrics_router_routes.metrics_router_routes_instance", "routes.#"),
					resource.TestCheckResourceAttr("data.ibm_metrics_router_routes.metrics_router_routes_instance", "routes.0.name", routeName),
					resource.TestCheckResourceAttr("data.ibm_metrics_router_routes.metrics_router_routes_instance", "routes.0.rules.0.inclusion_filters.0.operand", "location"),
					resource.TestCheckResourceAttr("data.ibm_metrics_router_routes.metrics_router_routes_instance", "routes.0.rules.0.inclusion_filters.0.operator", "is"),
					resource.TestCheckResourceAttr("data.ibm_metrics_router_routes.metrics_router_routes_instance", "routes.0.rules.0.inclusion_filters.0.values.0", "us-south"),
				),
			},
		},
	})
}

func testAccCheckIBMMetricsRouterRoutesDataSourceConfigBasic(routeName string) string {
	return fmt.Sprintf(`
        resource "ibm_metrics_router_settings" "mr_settings_instance" {
            permitted_target_regions = ["us-east", "us-south"]
            primary_metadata_region = "us-south"
            backup_metadata_region = "us-east"
        }

		resource "ibm_metrics_router_target" "metrics_router_target_instance" {
			name = "my-mr-target"
			destination_crn = "%s"
		}

		resource "ibm_metrics_router_route" "metrics_router_route_instance" {
			name = "%s"
			rules {
				action = "send"
				targets {
					id = ibm_metrics_router_target.metrics_router_target_instance.id
				}
				inclusion_filters {
					operand = "location"
					operator = "is"
					values = [ "us-south" ]
				}
			}
		}

		data "ibm_metrics_router_routes" "metrics_router_routes_instance" {
			name = ibm_metrics_router_route.metrics_router_route_instance.name
		}
	`, destinationCRN, routeName)
}
