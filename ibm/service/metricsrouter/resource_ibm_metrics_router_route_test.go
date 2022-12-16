// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package metricsrouter_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM/platform-services-go-sdk/metricsrouterv3"
)

func TestAccIBMMetricsRouterRouteBasic(t *testing.T) {
	var conf metricsrouterv3.Route
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	filterValue := "value"
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	updatedFilterValue := "value"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMMetricsRouterRouteDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMMetricsRouterRouteConfigBasic(name, filterValue),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMMetricsRouterRouteExists("ibm_metrics_router_route.metrics_router_route_instance", conf),
					resource.TestCheckResourceAttr("ibm_metrics_router_route.metrics_router_route_instance", "name", name),
					resource.TestCheckResourceAttr("ibm_metrics_router_route.metrics_router_route_instance", "rules.0.inclusion_filters.0.operand", "location"),
					resource.TestCheckResourceAttr("ibm_metrics_router_route.metrics_router_route_instance", "rules.0.inclusion_filters.0.operator", "is"),
					resource.TestCheckResourceAttr("ibm_metrics_router_route.metrics_router_route_instance", "rules.0.inclusion_filters.0.value.0", filterValue),
				),
			},
			{
				Config: testAccCheckIBMMetricsRouterRouteConfigBasic(nameUpdate, updatedFilterValue),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_metrics_router_route.metrics_router_route_instance", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_metrics_router_route.metrics_router_route_instance", "rules.0.inclusion_filters.0.operand", "location"),
					resource.TestCheckResourceAttr("ibm_metrics_router_route.metrics_router_route_instance", "rules.0.inclusion_filters.0.operator", "is"),
					resource.TestCheckResourceAttr("ibm_metrics_router_route.metrics_router_route_instance", "rules.0.inclusion_filters.0.value.0", updatedFilterValue),
				),
			},
			{
				ResourceName:      "ibm_metrics_router_route.metrics_router_route_instance",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMMetricsRouterRouteConfigBasic(name, filter_value string) string {
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
					value = [ "%s" ]
				}
			}
		}
	`, name, filter_value)
}

func testAccCheckIBMMetricsRouterRouteExists(n string, obj metricsrouterv3.Route) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		metricsRouterClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).MetricsRouterV3()
		if err != nil {
			return err
		}

		getRouteOptions := &metricsrouterv3.GetRouteOptions{}

		getRouteOptions.SetID(rs.Primary.ID)

		route, _, err := metricsRouterClient.GetRoute(getRouteOptions)
		if err != nil {
			return err
		}

		obj = *route
		return nil
	}
}

func testAccCheckIBMMetricsRouterRouteDestroy(s *terraform.State) error {
	metricsRouterClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).MetricsRouterV3()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_metrics_router_route" {
			continue
		}

		getRouteOptions := &metricsrouterv3.GetRouteOptions{}

		getRouteOptions.SetID(rs.Primary.ID)

		// Try to find the key
		_, response, err := metricsRouterClient.GetRoute(getRouteOptions)

		if err == nil {
			return fmt.Errorf("metrics_router_route still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for metrics_router_route (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
