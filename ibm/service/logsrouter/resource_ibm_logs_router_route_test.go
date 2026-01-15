// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package logsrouter_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/logsrouter"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/logsrouterv3"
	"github.com/stretchr/testify/assert"
)

func TestAccIBMLogsRouterRouteBasic(t *testing.T) {
	var conf logsrouterv3.Route
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMLogsRouterRouteDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMLogsRouterRouteConfigBasic(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMLogsRouterRouteExists("ibm_logs_router_route.logs_router_route_instance", conf),
					resource.TestCheckResourceAttr("ibm_logs_router_route.logs_router_route_instance", "name", name),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMLogsRouterRouteConfigBasic(nameUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_logs_router_route.logs_router_route_instance", "name", nameUpdate),
				),
			},
		},
	})
}

func TestAccIBMLogsRouterRouteAllArgs(t *testing.T) {
	var conf logsrouterv3.Route
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	managedBy := "enterprise"
	nameUpdate := fmt.Sprintf("%s_update", name)
	managedByUpdate := "enterprise" // managedBy is immutable

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMLogsRouterRouteDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMLogsRouterRouteConfig(name, managedBy),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMLogsRouterRouteExists("ibm_logs_router_route.logs_router_route_instance", conf),
					resource.TestCheckResourceAttr("ibm_logs_router_route.logs_router_route_instance", "name", name),
					resource.TestCheckResourceAttr("ibm_logs_router_route.logs_router_route_instance", "managed_by", managedBy),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMLogsRouterRouteConfig(nameUpdate, managedByUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_logs_router_route.logs_router_route_instance", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_logs_router_route.logs_router_route_instance", "managed_by", managedByUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_logs_router_route.logs_router_route_instance",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMLogsRouterRouteConfigBasic(name string) string {
	return fmt.Sprintf(`
		resource "ibm_logs_router_target" "logs_router_target_instance" {
			name = "my-lr-target1"
			destination_crn = "crn:v1:bluemix:public:logs:us-south:a/0be5ad401ae913d8ff665d92680664ed:22222222-2222-2222-2222-222222222222::"
			region = "us-south"
			managed_by = "account"
		}
		resource "ibm_logs_router_route" "logs_router_route_instance" {
			name = "%s"
			rules {
				action = "send"
				targets {
					id = ibm_logs_router_target.logs_router_target_instance.id
				}
				inclusion_filters {
					operand = "location"
					operator = "is"
					values = [ "us-south" ]
				}
			}
			managed_by = "account"
		}
	`, name)
}

func testAccCheckIBMLogsRouterRouteConfig(name string, managedBy string) string {
	return fmt.Sprintf(`
		resource "ibm_logs_router_target" "logs_router_target_instance" {
			name = "my-lr-target2"
			destination_crn = "crn:v1:bluemix:public:logs:us-south:a/0be5ad401ae913d8ff665d92680664ed:22222222-2222-2222-2222-222222222222::"
            managed_by = "%s"
			region = "us-south"
		}

		resource "ibm_logs_router_route" "logs_router_route_instance" {
			name = "%s"
			rules {
				action = "send"
				targets {
					id = ibm_logs_router_target.logs_router_target_instance.id
				}
				inclusion_filters {
					operand = "location"
					operator = "is"
					values = [ "us-south" ]
				}
			}
			managed_by = "%s"
		}
	`, managedBy, name, managedBy)
}

func testAccCheckIBMLogsRouterRouteExists(n string, obj logsrouterv3.Route) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		logsRouterClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).LogsRouterV3()
		if err != nil {
			return err
		}

		getRouteOptions := &logsrouterv3.GetRouteOptions{}

		getRouteOptions.SetID(rs.Primary.ID)

		route, _, err := logsRouterClient.GetRoute(getRouteOptions)
		if err != nil {
			return err
		}

		obj = *route
		return nil
	}
}

func testAccCheckIBMLogsRouterRouteDestroy(s *terraform.State) error {
	logsRouterClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).LogsRouterV3()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_logs_router_route" {
			continue
		}

		getRouteOptions := &logsrouterv3.GetRouteOptions{}

		getRouteOptions.SetID(rs.Primary.ID)

		// Try to find the key
		_, response, err := logsRouterClient.GetRoute(getRouteOptions)

		if err == nil {
			return fmt.Errorf("logs_router_route still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for logs_router_route (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

func TestResourceIBMLogsRouterRouteRuleToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		targetReferenceModel := make(map[string]interface{})
		targetReferenceModel["id"] = "c3af557f-fb0e-4476-85c3-0889e7fe7bc4"
		targetReferenceModel["crn"] = "crn:v1:bluemix:public:logs:us-south:a/0be5ad401ae913d8ff665d92680664ed:22222222-2222-2222-2222-222222222222::"
		targetReferenceModel["name"] = "a-lr-target-us-south"
		targetReferenceModel["target_type"] = "cloud_logs"

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

	targetReferenceModel := new(logsrouterv3.TargetReference)
	targetReferenceModel.ID = core.StringPtr("c3af557f-fb0e-4476-85c3-0889e7fe7bc4")
	targetReferenceModel.CRN = core.StringPtr("crn:v1:bluemix:public:logs:us-south:a/0be5ad401ae913d8ff665d92680664ed:22222222-2222-2222-2222-222222222222::")
	targetReferenceModel.Name = core.StringPtr("a-lr-target-us-south")
	targetReferenceModel.TargetType = core.StringPtr("cloud_logs")

	inclusionFilterModel := new(logsrouterv3.InclusionFilter)
	inclusionFilterModel.Operand = core.StringPtr("location")
	inclusionFilterModel.Operator = core.StringPtr("is")
	inclusionFilterModel.Values = []string{"us-south"}

	model := new(logsrouterv3.Rule)
	model.Action = core.StringPtr("send")
	model.Targets = []logsrouterv3.TargetReference{*targetReferenceModel}
	model.InclusionFilters = []logsrouterv3.InclusionFilter{*inclusionFilterModel}

	result, err := logsrouter.ResourceIBMLogsRouterRouteRuleToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMLogsRouterRouteTargetReferenceToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["id"] = "c3af557f-fb0e-4476-85c3-0889e7fe7bc4"
		model["crn"] = "crn:v1:bluemix:public:logs:us-south:a/0be5ad401ae913d8ff665d92680664ed:22222222-2222-2222-2222-222222222222::"
		model["name"] = "a-lr-target-us-south"
		model["target_type"] = "cloud_logs"

		assert.Equal(t, result, model)
	}

	model := new(logsrouterv3.TargetReference)
	model.ID = core.StringPtr("c3af557f-fb0e-4476-85c3-0889e7fe7bc4")
	model.CRN = core.StringPtr("crn:v1:bluemix:public:logs:us-south:a/0be5ad401ae913d8ff665d92680664ed:22222222-2222-2222-2222-222222222222::")
	model.Name = core.StringPtr("a-lr-target-us-south")
	model.TargetType = core.StringPtr("cloud_logs")

	result, err := logsrouter.ResourceIBMLogsRouterRouteTargetReferenceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMLogsRouterRouteInclusionFilterToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["operand"] = "location"
		model["operator"] = "is"
		model["values"] = []string{"us-south"}

		assert.Equal(t, result, model)
	}

	model := new(logsrouterv3.InclusionFilter)
	model.Operand = core.StringPtr("location")
	model.Operator = core.StringPtr("is")
	model.Values = []string{"us-south"}

	result, err := logsrouter.ResourceIBMLogsRouterRouteInclusionFilterToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMLogsRouterRouteMapToRulePrototype(t *testing.T) {
	checkResult := func(result *logsrouterv3.RulePrototype) {
		targetIdentityModel := new(logsrouterv3.TargetIdentity)
		targetIdentityModel.ID = core.StringPtr("c3af557f-fb0e-4476-85c3-0889e7fe7bc4")

		inclusionFilterPrototypeModel := new(logsrouterv3.InclusionFilterPrototype)
		inclusionFilterPrototypeModel.Operand = core.StringPtr("location")
		inclusionFilterPrototypeModel.Operator = core.StringPtr("is")
		inclusionFilterPrototypeModel.Values = []string{"us-south"}

		model := new(logsrouterv3.RulePrototype)
		model.Action = core.StringPtr("send")
		model.Targets = []logsrouterv3.TargetIdentity{*targetIdentityModel}
		model.InclusionFilters = []logsrouterv3.InclusionFilterPrototype{*inclusionFilterPrototypeModel}

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

	result, err := logsrouter.ResourceIBMLogsRouterRouteMapToRulePrototype(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMLogsRouterRouteMapToTargetIdentity(t *testing.T) {
	checkResult := func(result *logsrouterv3.TargetIdentity) {
		model := new(logsrouterv3.TargetIdentity)
		model.ID = core.StringPtr("c3af557f-fb0e-4476-85c3-0889e7fe7bc4")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["id"] = "c3af557f-fb0e-4476-85c3-0889e7fe7bc4"

	result, err := logsrouter.ResourceIBMLogsRouterRouteMapToTargetIdentity(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMLogsRouterRouteMapToInclusionFilterPrototype(t *testing.T) {
	checkResult := func(result *logsrouterv3.InclusionFilterPrototype) {
		model := new(logsrouterv3.InclusionFilterPrototype)
		model.Operand = core.StringPtr("location")
		model.Operator = core.StringPtr("is")
		model.Values = []string{"us-south"}

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["operand"] = "location"
	model["operator"] = "is"
	model["values"] = []interface{}{"us-south"}

	result, err := logsrouter.ResourceIBMLogsRouterRouteMapToInclusionFilterPrototype(model)
	assert.Nil(t, err)
	checkResult(result)
}
