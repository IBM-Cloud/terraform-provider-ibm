// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.108.0-56772134-20251111-102802
 */

package metricsrouter_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/metricsrouter"
	. "github.com/IBM-Cloud/terraform-provider-ibm/ibm/unittest"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/metricsrouterv3"
	"github.com/stretchr/testify/assert"
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

func TestAccIBMMetricsRouterRoutesDataSourceAllArgs(t *testing.T) {
	routeName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	routeManagedBy := "enterprise"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMMetricsRouterRoutesDataSourceConfig(routeName, routeManagedBy),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_metrics_router_routes.metrics_router_routes_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_metrics_router_routes.metrics_router_routes_instance", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_metrics_router_routes.metrics_router_routes_instance", "routes.#"),
					resource.TestCheckResourceAttrSet("data.ibm_metrics_router_routes.metrics_router_routes_instance", "routes.0.id"),
					resource.TestCheckResourceAttr("data.ibm_metrics_router_routes.metrics_router_routes_instance", "routes.0.name", routeName),
					resource.TestCheckResourceAttrSet("data.ibm_metrics_router_routes.metrics_router_routes_instance", "routes.0.crn"),
					resource.TestCheckResourceAttrSet("data.ibm_metrics_router_routes.metrics_router_routes_instance", "routes.0.created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_metrics_router_routes.metrics_router_routes_instance", "routes.0.updated_at"),
					resource.TestCheckResourceAttr("data.ibm_metrics_router_routes.metrics_router_routes_instance", "routes.0.managed_by", routeManagedBy),
				),
			},
		},
	})
}

func testAccCheckIBMMetricsRouterRoutesDataSourceConfigBasic(routeName string) string {
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
					values = [ "us-south" ]
				}
			}
		}

		data "ibm_metrics_router_routes" "metrics_router_routes_instance" {
			name = ibm_metrics_router_route.metrics_router_route_instance.name
		}
	`, destinationCRN, routeName)
}

func testAccCheckIBMMetricsRouterRoutesDataSourceConfig(routeName string, routeManagedBy string) string {
	return fmt.Sprintf(`
		resource "ibm_metrics_router_target" "metrics_router_target_instance" {
			name = "my-mr-target"
			destination_crn = "%s"
			managed_by = "enterprise"
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
			managed_by = "%s"
		}

		data "ibm_metrics_router_routes" "metrics_router_routes_instance" {
			name = ibm_metrics_router_route.metrics_router_route_instance.name
		}
	`, destinationCRN, routeName, routeManagedBy)
}

func TestDataSourceIBMMetricsRouterRoutesRouteToMap(t *testing.T) {
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

		ruleModel := make(map[string]interface{})
		ruleModel["action"] = "send"
		ruleModel["targets"] = []map[string]interface{}{targetReferenceModel}
		ruleModel["inclusion_filters"] = []map[string]interface{}{inclusionFilterModel}

		model := make(map[string]interface{})
		model["id"] = "c3af557f-fb0e-4476-85c3-0889e7fe7bc4"
		model["name"] = "my-route"
		model["crn"] = "crn:v1:bluemix:public:metrics-router:global:a/0be5ad401ae913d8ff665d92680664ed:b6eec08b-5201-08ca-451b-cd71523e3626:route:c3af557f-fb0e-4476-85c3-0889e7fe7bc4"
		model["rules"] = []map[string]interface{}{ruleModel}
		model["created_at"] = "2021-05-18T20:15:12.353Z"
		model["updated_at"] = "2021-05-18T20:15:12.353Z"
		model["managed_by"] = "enterprise"

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

	ruleModel := new(metricsrouterv3.Rule)
	ruleModel.Action = core.StringPtr("send")
	ruleModel.Targets = []metricsrouterv3.TargetReference{*targetReferenceModel}
	ruleModel.InclusionFilters = []metricsrouterv3.InclusionFilter{*inclusionFilterModel}

	model := new(metricsrouterv3.Route)
	model.ID = core.StringPtr("c3af557f-fb0e-4476-85c3-0889e7fe7bc4")
	model.Name = core.StringPtr("my-route")
	model.CRN = core.StringPtr("crn:v1:bluemix:public:metrics-router:global:a/0be5ad401ae913d8ff665d92680664ed:b6eec08b-5201-08ca-451b-cd71523e3626:route:c3af557f-fb0e-4476-85c3-0889e7fe7bc4")
	model.Rules = []metricsrouterv3.Rule{*ruleModel}
	model.CreatedAt = CreateMockDateTime("2021-05-18T20:15:12.353Z")
	model.UpdatedAt = CreateMockDateTime("2021-05-18T20:15:12.353Z")
	model.ManagedBy = core.StringPtr("enterprise")

	result, err := metricsrouter.DataSourceIBMMetricsRouterRoutesRouteToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMMetricsRouterRoutesRuleToMap(t *testing.T) {
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

	result, err := metricsrouter.DataSourceIBMMetricsRouterRoutesRuleToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMMetricsRouterRoutesTargetReferenceToMap(t *testing.T) {
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

	result, err := metricsrouter.DataSourceIBMMetricsRouterRoutesTargetReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMMetricsRouterRoutesInclusionFilterToMap(t *testing.T) {
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

	result, err := metricsrouter.DataSourceIBMMetricsRouterRoutesInclusionFilterToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
