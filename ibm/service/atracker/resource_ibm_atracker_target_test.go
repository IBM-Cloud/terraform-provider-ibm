// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package atracker_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/platform-services-go-sdk/atrackerv2"
	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/conns"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/service/atracker"
	. "github.com/Mavrickk3/terraform-provider-ibm/ibm/unittest"
	"github.com/stretchr/testify/assert"
)

func TestAccIBMAtrackerTargetBasic(t *testing.T) {
	var conf atrackerv2.Target
	name := fmt.Sprintf("tf_basic_name_1")
	targetType := "cloud_object_storage"
	nameUpdate := fmt.Sprintf("tf_basic_name_2")

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMAtrackerTargetDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMAtrackerTargetConfigBasic(name, targetType),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMAtrackerTargetExists("ibm_atracker_target.atracker_target_instance", conf),
					resource.TestCheckResourceAttr("ibm_atracker_target.atracker_target_instance", "name", name),
					resource.TestCheckResourceAttr("ibm_atracker_target.atracker_target_instance", "target_type", targetType),
				),
			},
			{
				Config: testAccCheckIBMAtrackerTargetConfigBasic(nameUpdate, targetType),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_atracker_target.atracker_target_instance", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_atracker_target.atracker_target_instance", "target_type", targetType),
				),
			},
		},
	})
}

func TestAccIBMAtrackerTargetAllArgs(t *testing.T) {
	var conf atrackerv2.Target
	name := fmt.Sprintf("tf_all_name_1")
	targetType := "cloud_object_storage"
	region := fmt.Sprintf("us-south")
	// targetType and region cannot be changed
	nameUpdate := fmt.Sprintf("tf_all_name_2")

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMAtrackerTargetDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMAtrackerTargetConfig(name, targetType, region),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMAtrackerTargetExists("ibm_atracker_target.atracker_target_instance", conf),
					resource.TestCheckResourceAttr("ibm_atracker_target.atracker_target_instance", "name", name),
					resource.TestCheckResourceAttr("ibm_atracker_target.atracker_target_instance", "target_type", targetType),
					resource.TestCheckResourceAttr("ibm_atracker_target.atracker_target_instance", "region", region),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMAtrackerTargetConfig(nameUpdate, targetType, region),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_atracker_target.atracker_target_instance", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_atracker_target.atracker_target_instance", "target_type", targetType),
					resource.TestCheckResourceAttr("ibm_atracker_target.atracker_target_instance", "region", region),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_atracker_target.atracker_target_instance",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMAtrackerTargetConfigBasic(name string, targetType string) string {
	return fmt.Sprintf(`
		resource "ibm_atracker_target" "atracker_target_instance" {
			name = "%s"
			target_type = "%s"
			cos_endpoint {
					endpoint = "s3.private.us-east.cloud-object-storage.appdomain.cloud"
					target_crn = "crn:v1:bluemix:public:cloud-object-storage:global:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::"
					bucket = "my-atracker-bucket"
					api_key = "xxxxxxxxxxxxxx"
					service_to_service_enabled = true
			}
		}
	`, name, targetType)
}

func testAccCheckIBMAtrackerTargetConfig(name string, targetType string, region string) string {
	return fmt.Sprintf(`
		resource "ibm_atracker_target" "atracker_target_instance" {
			name = "%s"
			target_type = "%s"
			region = "%s"
			cos_endpoint {
				endpoint = "s3.private.us-east.cloud-object-storage.appdomain.cloud"
				target_crn = "crn:v1:bluemix:public:cloud-object-storage:global:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::"
				bucket = "my-atracker-bucket"
				api_key = "xxxxxxxxxxxxxx"
				service_to_service_enabled = true
			}
		}
	`, name, targetType, region)
}

func testAccCheckIBMAtrackerTargetExists(n string, obj atrackerv2.Target) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		atrackerClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).AtrackerV2()
		if err != nil {
			return err
		}

		getTargetOptions := &atrackerv2.GetTargetOptions{}

		getTargetOptions.SetID(rs.Primary.ID)

		target, _, err := atrackerClient.GetTarget(getTargetOptions)
		if err != nil {
			return err
		}

		obj = *target
		return nil
	}
}

func testAccCheckIBMAtrackerTargetDestroy(s *terraform.State) error {
	atrackerClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).AtrackerV2()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_atracker_target" {
			continue
		}

		getTargetOptions := &atrackerv2.GetTargetOptions{}

		getTargetOptions.SetID(rs.Primary.ID)

		// Try to find the key
		_, response, err := atrackerClient.GetTarget(getTargetOptions)

		if err == nil {
			return fmt.Errorf("Activity Tracker Target still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("[ERROR] Error checking for Activity Tracker Target (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

func TestResourceIBMAtrackerTargetCosEndpointToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["endpoint"] = core.StringPtr("s3.private.us-east.cloud-object-storage.appdomain.cloud")
		model["target_crn"] = core.StringPtr("crn:v1:bluemix:public:cloud-object-storage:global:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::")
		model["bucket"] = core.StringPtr("my-atracker-bucket")
		model["service_to_service_enabled"] = core.BoolPtr(true)
		model["api_key"] = "REDACTED"

		assert.Equal(t, result, model)
	}

	model := new(atrackerv2.CosEndpoint)
	model.Endpoint = core.StringPtr("s3.private.us-east.cloud-object-storage.appdomain.cloud")
	model.TargetCRN = core.StringPtr("crn:v1:bluemix:public:cloud-object-storage:global:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::")
	model.Bucket = core.StringPtr("my-atracker-bucket")
	model.ServiceToServiceEnabled = core.BoolPtr(true)

	result, err := atracker.ResourceIBMAtrackerTargetCosEndpointToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMAtrackerTargetEventstreamsEndpointToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["target_crn"] = core.StringPtr("crn:v1:bluemix:public:messagehub:us-south:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::")
		model["brokers"] = []string{"kafka-x:9094"}
		model["topic"] = core.StringPtr("my-topic")
		model["api_key"] = "REDACTED"
		model["service_to_service_enabled"] = core.BoolPtr(false)

		assert.Equal(t, result, model)
	}

	model := new(atrackerv2.EventstreamsEndpoint)
	model.TargetCRN = core.StringPtr("crn:v1:bluemix:public:messagehub:us-south:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::")
	model.Brokers = []string{"kafka-x:9094"}
	model.Topic = core.StringPtr("my-topic")
	model.APIKey = core.StringPtr("xxxxxxxxxxxxxx")
	model.ServiceToServiceEnabled = core.BoolPtr(false)

	result, err := atracker.ResourceIBMAtrackerTargetEventstreamsEndpointToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMAtrackerTargetCloudLogsEndpointToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["target_crn"] = core.StringPtr("crn:v1:bluemix:public:logs:eu-es:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::")

		assert.Equal(t, result, model)
	}

	model := new(atrackerv2.CloudLogsEndpoint)
	model.TargetCRN = core.StringPtr("crn:v1:bluemix:public:logs:eu-es:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::")

	result, err := atracker.ResourceIBMAtrackerTargetCloudLogsEndpointToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMAtrackerTargetWriteStatusToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["status"] = core.StringPtr("success")
		model["last_failure"] = "2021-05-18T20:15:12.353Z"
		model["reason_for_last_failure"] = core.StringPtr("Provided API key could not be found")

		assert.Equal(t, result, model)
	}

	model := new(atrackerv2.WriteStatus)
	model.Status = core.StringPtr("success")
	model.LastFailure = CreateMockDateTime("2021-05-18T20:15:12.353Z")
	model.ReasonForLastFailure = core.StringPtr("Provided API key could not be found")

	result, err := atracker.ResourceIBMAtrackerTargetWriteStatusToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMAtrackerTargetMapToCosEndpointPrototype(t *testing.T) {
	checkResult := func(result *atrackerv2.CosEndpointPrototype) {
		model := new(atrackerv2.CosEndpointPrototype)
		model.Endpoint = core.StringPtr("s3.private.us-east.cloud-object-storage.appdomain.cloud")
		model.TargetCRN = core.StringPtr("crn:v1:bluemix:public:cloud-object-storage:global:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::")
		model.Bucket = core.StringPtr("my-atracker-bucket")
		model.APIKey = core.StringPtr("xxxxxxxxxxxxxx")
		model.ServiceToServiceEnabled = core.BoolPtr(true)

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["endpoint"] = "s3.private.us-east.cloud-object-storage.appdomain.cloud"
	model["target_crn"] = "crn:v1:bluemix:public:cloud-object-storage:global:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::"
	model["bucket"] = "my-atracker-bucket"
	model["api_key"] = "xxxxxxxxxxxxxx"
	model["service_to_service_enabled"] = true

	result, err := atracker.ResourceIBMAtrackerTargetMapToCosEndpointPrototype(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMAtrackerTargetMapToEventstreamsEndpointPrototype(t *testing.T) {
	checkResult := func(result *atrackerv2.EventstreamsEndpointPrototype) {
		model := new(atrackerv2.EventstreamsEndpointPrototype)
		model.TargetCRN = core.StringPtr("crn:v1:bluemix:public:messagehub:us-south:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::")
		model.Brokers = []string{"kafka-x:9094"}
		model.Topic = core.StringPtr("my-topic")
		model.APIKey = core.StringPtr("xxxxxxxxxxxxxx")
		model.ServiceToServiceEnabled = core.BoolPtr(false)

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["target_crn"] = "crn:v1:bluemix:public:messagehub:us-south:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::"
	model["brokers"] = []interface{}{"kafka-x:9094"}
	model["topic"] = "my-topic"
	model["api_key"] = "xxxxxxxxxxxxxx"
	model["service_to_service_enabled"] = false

	result, err := atracker.ResourceIBMAtrackerTargetMapToEventstreamsEndpointPrototype(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIBMAtrackerTargetMapToCloudLogsEndpointPrototype(t *testing.T) {
	checkResult := func(result *atrackerv2.CloudLogsEndpointPrototype) {
		model := new(atrackerv2.CloudLogsEndpointPrototype)
		model.TargetCRN = core.StringPtr("crn:v1:bluemix:public:logs:eu-es:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["target_crn"] = "crn:v1:bluemix:public:logs:eu-es:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::"

	result, err := atracker.ResourceIBMAtrackerTargetMapToCloudLogsEndpointPrototype(model)
	assert.Nil(t, err)
	checkResult(result)
}
