// Copyright IBM Corp. 2025 All Rights Reserved.
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
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/metricsrouter"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/metricsrouterv3"
	"github.com/stretchr/testify/assert"
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

func TestAccIBMMetricsRouterRouteAllArgs(t *testing.T) {
	var conf metricsrouterv3.Route
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	managedBy := "enterprise"
	managedByTarget := managedBy

	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	managedByUpdate := "account"
	managedByTargetUpdate := managedByUpdate

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMMetricsRouterRouteDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMMetricsRouterRouteConfig(name, managedBy, managedByTarget),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMMetricsRouterRouteExists("ibm_metrics_router_route.metrics_router_route_instance", conf),
					resource.TestCheckResourceAttr("ibm_metrics_router_route.metrics_router_route_instance", "name", name),
					resource.TestCheckResourceAttr("ibm_metrics_router_route.metrics_router_route_instance", "managed_by", managedBy),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMMetricsRouterRouteConfig(nameUpdate, managedByUpdate, managedByTargetUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_metrics_router_route.metrics_router_route_instance", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_metrics_router_route.metrics_router_route_instance", "managed_by", managedByUpdate),
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

func testAccCheckIBMMetricsRouterRouteConfig(name string, managedBy string, managedTargetBy string) string {
	return fmt.Sprintf(`
		resource "ibm_metrics_router_target" "metrics_router_target_instance" {
			name = "my-mr-target"
			destination_crn = "%s"
			managed_by = "%s"
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
			managed_by = "%s"
		}
	`, destinationCRN, managedTargetBy, name, managedBy)

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

func TestResourceIBMMetricsRouterRouteRuleToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		targetReferenceModel := make(map[string]interface{})
		targetReferenceModel["id"] = "c3af557f-fb0e-4476-85c3-0889e7fe7bc4"
		targetReferenceModel["crn"] = "crn:v1:bluemix:public:sysdig-monitor:us-south:a/0be5ad401ae913d8ff665d92680664ed:22222222-2222-2222-2222-222222222222::"
		targetReferenceModel["name"] = "a-mr-target-us-south"
		targetReferenceModel["target_type"] = "sysdig_monitor"

		inclusionFilterModel := make(map[string]interface{})
		inclusionFilterModel["operand"] = "location"
		inclusionFilterModel["operator"] = "is"
		inclusionFilterModel["values"] = []string{"us-south"}

		model := make(map[string]interface{})
		model["action"] = "send"
		model["targets"] = []map[string]interface{}{targetReferenceModel}
		model["inclusion_filters"] = []map[string]interface{}{inclusionFilterModel}

		assert.Equal(t, result, model)
	}

	targetReferenceModel := new(metricsrouterv3.TargetReference)
	targetReferenceModel.ID = core.StringPtr("c3af557f-fb0e-4476-85c3-0889e7fe7bc4")
	targetReferenceModel.CRN = core.StringPtr("crn:v1:bluemix:public:sysdig-monitor:us-south:a/0be5ad401ae913d8ff665d92680664ed:22222222-2222-2222-2222-222222222222::")
	targetReferenceModel.Name = core.StringPtr("a-mr-target-us-south")
	targetReferenceModel.TargetType = core.StringPtr("sysdig_monitor")

	inclusionFilterModel := new(metricsrouterv3.InclusionFilter)
	inclusionFilterModel.Operand = core.StringPtr("location")
	inclusionFilterModel.Operator = core.StringPtr("is")
	inclusionFilterModel.Values = []string{"us-south"}

	model := new(metricsrouterv3.Rule)
	model.Action = core.StringPtr("send")
	model.Targets = []metricsrouterv3.TargetReference{*targetReferenceModel}
	model.InclusionFilters = []metricsrouterv3.InclusionFilter{*inclusionFilterModel}

	result, err := metricsrouter.ResourceIBMMetricsRouterRouteRuleToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMMetricsRouterRouteTargetReferenceToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["id"] = "c3af557f-fb0e-4476-85c3-0889e7fe7bc4"
		model["crn"] = "crn:v1:bluemix:public:sysdig-monitor:us-south:a/0be5ad401ae913d8ff665d92680664ed:22222222-2222-2222-2222-222222222222::"
		model["name"] = "a-mr-target-us-south"
		model["target_type"] = "sysdig_monitor"

		assert.Equal(t, result, model)
	}

	model := new(metricsrouterv3.TargetReference)
	model.ID = core.StringPtr("c3af557f-fb0e-4476-85c3-0889e7fe7bc4")
	model.CRN = core.StringPtr("crn:v1:bluemix:public:sysdig-monitor:us-south:a/0be5ad401ae913d8ff665d92680664ed:22222222-2222-2222-2222-222222222222::")
	model.Name = core.StringPtr("a-mr-target-us-south")
	model.TargetType = core.StringPtr("sysdig_monitor")

	result, err := metricsrouter.ResourceIBMMetricsRouterRouteTargetReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMMetricsRouterRouteInclusionFilterToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["operand"] = "location"
		model["operator"] = "is"
		model["values"] = []string{"us-south"}

		assert.Equal(t, result, model)
	}

	model := new(metricsrouterv3.InclusionFilter)
	model.Operand = core.StringPtr("location")
	model.Operator = core.StringPtr("is")
	model.Values = []string{"us-south"}

	result, err := metricsrouter.ResourceIBMMetricsRouterRouteInclusionFilterToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMMetricsRouterRouteMapToRulePrototype(t *testing.T) {
	checkResult := func(result *metricsrouterv3.RulePrototype) {
		targetIdentityModel := new(metricsrouterv3.TargetIdentity)
		targetIdentityModel.ID = core.StringPtr("c3af557f-fb0e-4476-85c3-0889e7fe7bc4")

		inclusionFilterPrototypeModel := new(metricsrouterv3.InclusionFilterPrototype)
		inclusionFilterPrototypeModel.Operand = core.StringPtr("location")
		inclusionFilterPrototypeModel.Operator = core.StringPtr("is")
		inclusionFilterPrototypeModel.Values = []string{"us-south"}

		model := new(metricsrouterv3.RulePrototype)
		model.Action = core.StringPtr("send")
		model.Targets = []metricsrouterv3.TargetIdentity{*targetIdentityModel}
		model.InclusionFilters = []metricsrouterv3.InclusionFilterPrototype{*inclusionFilterPrototypeModel}

		assert.Equal(t, result, model)
	}

	targetIdentityModel := make(map[string]interface{})
	targetIdentityModel["id"] = "c3af557f-fb0e-4476-85c3-0889e7fe7bc4"

	inclusionFilterPrototypeModel := make(map[string]interface{})
	inclusionFilterPrototypeModel["operand"] = "location"
	inclusionFilterPrototypeModel["operator"] = "is"
	inclusionFilterPrototypeModel["values"] = []interface{}{"us-south"}

	model := make(map[string]interface{})
	model["action"] = "send"
	model["targets"] = []interface{}{targetIdentityModel}
	model["inclusion_filters"] = []interface{}{inclusionFilterPrototypeModel}

	result, err := metricsrouter.ResourceIBMMetricsRouterRouteMapToRulePrototype(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMMetricsRouterRouteMapToTargetIdentity(t *testing.T) {
	checkResult := func(result *metricsrouterv3.TargetIdentity) {
		model := new(metricsrouterv3.TargetIdentity)
		model.ID = core.StringPtr("c3af557f-fb0e-4476-85c3-0889e7fe7bc4")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["id"] = "c3af557f-fb0e-4476-85c3-0889e7fe7bc4"

	result, err := metricsrouter.ResourceIBMMetricsRouterRouteMapToTargetIdentity(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMMetricsRouterRouteMapToInclusionFilterPrototype(t *testing.T) {
	checkResult := func(result *metricsrouterv3.InclusionFilterPrototype) {
		model := new(metricsrouterv3.InclusionFilterPrototype)
		model.Operand = core.StringPtr("location")
		model.Operator = core.StringPtr("is")
		model.Values = []string{"us-south"}

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["operand"] = "location"
	model["operator"] = "is"
	model["values"] = []interface{}{"us-south"}

	result, err := metricsrouter.ResourceIBMMetricsRouterRouteMapToInclusionFilterPrototype(model)
	assert.Nil(t, err)
	checkResult(result)
}
