// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package atracker_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/atracker"
	"github.com/IBM/platform-services-go-sdk/atrackerv2"
	"github.com/stretchr/testify/assert"
)

func TestAccIBMAtrackerRouteBasic(t *testing.T) {
	var conf atrackerv2.Route
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMAtrackerRouteDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMAtrackerRouteConfigBasic(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMAtrackerRouteExists("ibm_atracker_route.atracker_route_instance", conf),
					resource.TestCheckResourceAttr("ibm_atracker_route.atracker_route_instance", "name", name),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMAtrackerRouteConfigBasic(nameUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_atracker_route.atracker_route_instance", "name", nameUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_atracker_route.atracker_route_instance",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccIBMAtrackerRouteBasicMultipleRules(t *testing.T) {
	var conf atrackerv2.Route
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMAtrackerRouteDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMAtrackerRouteConfigBasicMultipleRules(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMAtrackerRouteExists("ibm_atracker_route.atracker_route_instance", conf),
					resource.TestCheckResourceAttr("ibm_atracker_route.atracker_route_instance", "name", name),
				),
			},
			{
				Config: testAccCheckIBMAtrackerRouteConfigBasicMultipleRules(nameUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_atracker_route.atracker_route_instance", "name", nameUpdate),
				),
			},
			{
				ResourceName:      "ibm_atracker_route.atracker_route_instance",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMAtrackerRouteConfigBasic(name string) string {
	return fmt.Sprintf(`
		resource "ibm_atracker_target" "atracker_target_instance" {
			name = "my-cos-target"
			target_type = "cloud_object_storage"
			region = "us-south"
			cos_endpoint {
				endpoint = "s3.private.us-east.cloud-object-storage.appdomain.cloud"
				target_crn = "crn:v1:bluemix:public:cloud-object-storage:global:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::"
				bucket = "my-atracker-bucket"
				api_key = "xxxxxxxxxxxxxx"
				service_to_service_enabled = false
			}
		}

		resource "ibm_atracker_route" "atracker_route_instance" {
			name = "%s"
			rules {
				target_ids = [ ibm_atracker_target.atracker_target_instance.id ]
				locations = [ "us-south" ]
			}
		}
	`, name)
}

func testAccCheckIBMAtrackerRouteConfigBasicMultipleRules(name string) string {
	return fmt.Sprintf(`
		resource "ibm_atracker_target" "atracker_target_instance" {
			name = "my-cos-target"
			target_type = "cloud_object_storage"
			region = "us-south"
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
			rules {
				target_ids = [ ibm_atracker_target.atracker_target_instance.id ]
				locations = [ "us-east" ]
			}
		}
	`, name)
}

func testAccCheckIBMAtrackerRouteExists(n string, obj atrackerv2.Route) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		atrackerClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).AtrackerV2()
		if err != nil {
			return err
		}

		getRouteOptions := &atrackerv2.GetRouteOptions{}

		getRouteOptions.SetID(rs.Primary.ID)

		route, _, err := atrackerClient.GetRoute(getRouteOptions)
		if err != nil {
			return err
		}

		obj = *route
		return nil
	}
}

func testAccCheckIBMAtrackerRouteDestroy(s *terraform.State) error {
	atrackerClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).AtrackerV2()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_atracker_route" {
			continue
		}

		getRouteOptions := &atrackerv2.GetRouteOptions{}

		getRouteOptions.SetID(rs.Primary.ID)

		// Try to find the key
		_, response, err := atrackerClient.GetRoute(getRouteOptions)

		if err == nil {
			return fmt.Errorf("atracker_route_instance still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for atracker_route_instance (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

func TestResourceIBMAtrackerRouteRuleToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["target_ids"] = []string{"c3af557f-fb0e-4476-85c3-0889e7fe7bc4"}
		model["locations"] = []string{"us-south"}

		assert.Equal(t, result, model)
	}

	model := new(atrackerv2.Rule)
	model.TargetIds = []string{"c3af557f-fb0e-4476-85c3-0889e7fe7bc4"}
	model.Locations = []string{"us-south"}

	result, err := atracker.ResourceIBMAtrackerRouteRuleToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMAtrackerRouteMapToRulePrototype(t *testing.T) {
	checkResult := func(result *atrackerv2.RulePrototype) {
		model := new(atrackerv2.RulePrototype)
		model.TargetIds = []string{"c3af557f-fb0e-4476-85c3-0889e7fe7bc4"}
		model.Locations = []string{"us-south"}

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["target_ids"] = []interface{}{"c3af557f-fb0e-4476-85c3-0889e7fe7bc4"}
	model["locations"] = []interface{}{"us-south"}

	result, err := atracker.ResourceIBMAtrackerRouteMapToRulePrototype(model)
	assert.Nil(t, err)
	checkResult(result)
}
