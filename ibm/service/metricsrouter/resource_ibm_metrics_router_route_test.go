// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package metricsrouter_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM/platform-services-go-sdk/metricsrouterv3"
)

func TestAccIBMMetricsRouterRouteBasic(t *testing.T) {
	var conf metricsrouterv3.Route
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	filterValue := "value"
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	updatedFilterValue := "value1"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMMetricsRouterRouteDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMMetricsRouterRouteConfigBasic(name, filterValue),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMMetricsRouterRouteExists("ibm_metrics_router_route.metrics_router_route_instance", conf),
					resource.TestCheckResourceAttr("ibm_metrics_router_route.metrics_router_route_instance", "name", name),
					resource.TestCheckResourceAttr("ibm_metrics_router_route.metrics_router_route_instance", "rules.0.inclusion_filters.0.operand", "location"),
					resource.TestCheckResourceAttr("ibm_metrics_router_route.metrics_router_route_instance", "rules.0.inclusion_filters.0.operator", "is"),
					resource.TestCheckResourceAttr("ibm_metrics_router_route.metrics_router_route_instance", "rules.0.inclusion_filters.0.values.0", filterValue),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMMetricsRouterRouteConfigBasic(nameUpdate, updatedFilterValue),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_metrics_router_route.metrics_router_route_instance", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_metrics_router_route.metrics_router_route_instance", "rules.0.inclusion_filters.0.operand", "location"),
					resource.TestCheckResourceAttr("ibm_metrics_router_route.metrics_router_route_instance", "rules.0.inclusion_filters.0.operator", "is"),
					resource.TestCheckResourceAttr("ibm_metrics_router_route.metrics_router_route_instance", "rules.0.inclusion_filters.0.values.0", updatedFilterValue),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_metrics_router_route.metrics_router_route_instance",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccIBMMetricsRouterRouteConfigBasicWithoutAction(t *testing.T) {
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMMetricsRouterRouteDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMMetricsRouterRouteConfigBasicWithoutAction(name),
			},
		},
	})
}

func TestAccIBMMetricsRouterRouteSendEmptyTarget(t *testing.T) {
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	action := "send"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMMetricsRouterRouteDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckIBMMetricsRouterRouteEmptyTarget(name, action),
				ExpectError: regexp.MustCompile("should match regexp"),
			},
		},
	})
}

func TestAccIBMMetricsRouterRouteDropEmptyTarget(t *testing.T) {
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	action := "drop"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMMetricsRouterRouteDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckIBMMetricsRouterRouteEmptyTarget(name, action),
				ExpectError: regexp.MustCompile("should match regexp"),
			},
		},
	})
}

func TestAccIBMMetricsRouterRouteDropNoTarget(t *testing.T) {
	var conf metricsrouterv3.Route
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	action := "drop"
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMMetricsRouterRouteDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMMetricsRouterRouteNoTarget(name, action),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMMetricsRouterRouteExists("ibm_metrics_router_route.metrics_router_route_instance", conf),
					resource.TestCheckResourceAttr("ibm_metrics_router_route.metrics_router_route_instance", "rules.0.inclusion_filters.0.operand", "resource_type"),
					resource.TestCheckResourceAttr("ibm_metrics_router_route.metrics_router_route_instance", "rules.0.inclusion_filters.0.operator", "is"),
					resource.TestCheckResourceAttr("ibm_metrics_router_route.metrics_router_route_instance", "rules.0.inclusion_filters.0.values.0", "worker"),
				),
			},
		},
	})
}

func TestAccIBMMetricsRouterRouteSendNoTarget(t *testing.T) {
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	action := "send"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMMetricsRouterRouteDestroy,
		Steps: []resource.TestStep{
			{
				Config:      testAccCheckIBMMetricsRouterRouteNoTarget(name, action),
				ExpectError: regexp.MustCompile("You have a rule with empty targets."),
			},
		},
	})
}

func testAccCheckIBMMetricsRouterRouteConfigBasic(name, filter_value string) string {
	return fmt.Sprintf(`
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
					values = [ "%s" ]
				}
			}
		}
	`, destinationCRN, name, filter_value)
}

func testAccCheckIBMMetricsRouterRouteConfigBasicWithoutAction(name string) string {
	return fmt.Sprintf(`
		resource "ibm_metrics_router_target" "metrics_router_target_instance" {
			name = "my-mr-target"
			destination_crn = "%s"
		}

		resource "ibm_metrics_router_route" "metrics_router_route_instance" {
			name = "%s"
			rules {
				targets {
					id = ibm_metrics_router_target.metrics_router_target_instance.id
				}
				inclusion_filters {
					operand = "resource_type"
					operator = "is"
					values = ["worker"]
				}
			}
		}`, destinationCRN, name)
}

func testAccCheckIBMMetricsRouterRouteEmptyTarget(name string, action string) string {
	return fmt.Sprintf(`
	resource "ibm_metrics_router_route" "metrics_router_route_instance" {
		name = "%s"
		rules {
			action = "%s"
			targets {
				id = ""
			}
			inclusion_filters {
				operand = "resource_type"
				operator = "is"
				values = ["worker"]
			}
		}
	}`, name, action)
}

func testAccCheckIBMMetricsRouterRouteNoTarget(name string, action string) string {
	return fmt.Sprintf(`
	resource "ibm_metrics_router_route" "metrics_router_route_instance" {
		name = "%s"
		rules {
			action = "%s"
			inclusion_filters {
				operand = "resource_type"
				operator = "is"
				values = ["worker"]
			}
		}
	}`, name, action)
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
