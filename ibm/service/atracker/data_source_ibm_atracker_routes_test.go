// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.101.0-62624c1e-20250225-192301
 */

package atracker_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/atracker"
	. "github.com/IBM-Cloud/terraform-provider-ibm/ibm/unittest"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/atrackerv2"
	"github.com/stretchr/testify/assert"
)

func TestAccIBMAtrackerRoutesDataSourceBasic(t *testing.T) {
	routeName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMAtrackerRoutesDataSourceConfigBasic(routeName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_atracker_routes.atracker_routes_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_atracker_routes.atracker_routes_instance", "routes.#"),
					resource.TestCheckResourceAttr("data.ibm_atracker_routes.atracker_routes_instance", "routes.0.name", routeName),
				),
			},
		},
	})
}

func testAccCheckIBMAtrackerRoutesDataSourceConfigBasic(routeName string) string {
	return fmt.Sprintf(`
		resource "ibm_atracker_target" "atracker_target_instance" {
			name = "my-cos-target"
			target_type = "cloud_object_storage"
			cos_endpoint {
				endpoint = "s3.private.us-east.cloud-object-storage.appdomain.cloud"
				target_crn = "crn:v1:bluemix:public:cloud-object-storage:global:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::"
				bucket = "my-atracker-bucket"
				api_key = "xxxxxxxxxxxxxx"
			}
		}

		resource "ibm_atracker_route" "atracker_route_instance" {
			name = "%s"
			rules {
				target_ids = [ ibm_atracker_target.atracker_target_instance.id ]
				locations = [ "us-south" ]
			}
		}

		data "ibm_atracker_routes" "atracker_routes_instance" {
			name = ibm_atracker_route.atracker_route_instance.name
		}
	`, routeName)
}

func TestDataSourceIBMAtrackerRoutesRouteToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		ruleModel := make(map[string]interface{})
		ruleModel["target_ids"] = []string{"c3af557f-fb0e-4476-85c3-0889e7fe7bc4"}
		ruleModel["locations"] = []string{"us-south"}

		model := make(map[string]interface{})
		model["id"] = "c3af557f-fb0e-4476-85c3-0889e7fe7bc4"
		model["name"] = "my-route"
		model["crn"] = "crn:v1:bluemix:public:atracker:global:a/11111111111111111111111111111111:b6eec08b-5201-08ca-451b-cd71523e3626:route:c3af557f-fb0e-4476-85c3-0889e7fe7bc4"
		model["version"] = int(0)
		model["rules"] = []map[string]interface{}{ruleModel}
		model["created_at"] = "2021-05-18T20:15:12.353Z"
		model["updated_at"] = "2021-05-18T20:15:12.353Z"
		model["api_version"] = int(2)
		model["message"] = "Route was created successfully."

		assert.Equal(t, result, model)
	}

	ruleModel := new(atrackerv2.Rule)
	ruleModel.TargetIds = []string{"c3af557f-fb0e-4476-85c3-0889e7fe7bc4"}
	ruleModel.Locations = []string{"us-south"}

	model := new(atrackerv2.Route)
	model.ID = core.StringPtr("c3af557f-fb0e-4476-85c3-0889e7fe7bc4")
	model.Name = core.StringPtr("my-route")
	model.CRN = core.StringPtr("crn:v1:bluemix:public:atracker:global:a/11111111111111111111111111111111:b6eec08b-5201-08ca-451b-cd71523e3626:route:c3af557f-fb0e-4476-85c3-0889e7fe7bc4")
	model.Version = core.Int64Ptr(int64(0))
	model.Rules = []atrackerv2.Rule{*ruleModel}
	model.CreatedAt = CreateMockDateTime("2021-05-18T20:15:12.353Z")
	model.UpdatedAt = CreateMockDateTime("2021-05-18T20:15:12.353Z")
	model.APIVersion = core.Int64Ptr(int64(2))
	model.Message = core.StringPtr("Route was created successfully.")

	result, err := atracker.DataSourceIBMAtrackerRoutesRouteToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMAtrackerRoutesRuleToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["target_ids"] = []string{"c3af557f-fb0e-4476-85c3-0889e7fe7bc4"}
		model["locations"] = []string{"us-south"}

		assert.Equal(t, result, model)
	}

	model := new(atrackerv2.Rule)
	model.TargetIds = []string{"c3af557f-fb0e-4476-85c3-0889e7fe7bc4"}
	model.Locations = []string{"us-south"}

	result, err := atracker.DataSourceIBMAtrackerRoutesRuleToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
