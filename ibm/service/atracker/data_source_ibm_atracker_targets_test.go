// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.108.0-56772134-20251111-102802
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

func TestAccIBMAtrackerTargetsDataSourceBasic(t *testing.T) {
	targetName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	targetTargetType := "cloud_object_storage"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMAtrackerTargetsDataSourceConfigBasic(targetName, targetTargetType),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_atracker_targets.atracker_targets_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_atracker_targets.atracker_targets_instance", "targets.#"),
					resource.TestCheckResourceAttr("data.ibm_atracker_targets.atracker_targets_instance", "targets.0.name", targetName),
					resource.TestCheckResourceAttr("data.ibm_atracker_targets.atracker_targets_instance", "targets.0.target_type", targetTargetType),
				),
			},
		},
	})
}

func TestAccIBMAtrackerTargetsDataSourceAllArgs(t *testing.T) {
	targetName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	targetTargetType := "cloud_object_storage"
	targetRegion := "us-south"
	targetManagedBy := "enterprise"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMAtrackerTargetsDataSourceConfig(targetName, targetTargetType, targetRegion, targetManagedBy),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_atracker_targets.atracker_targets_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_atracker_targets.atracker_targets_instance", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_atracker_targets.atracker_targets_instance", "targets.#"),
					resource.TestCheckResourceAttrSet("data.ibm_atracker_targets.atracker_targets_instance", "targets.0.id"),
					resource.TestCheckResourceAttr("data.ibm_atracker_targets.atracker_targets_instance", "targets.0.name", targetName),
					resource.TestCheckResourceAttrSet("data.ibm_atracker_targets.atracker_targets_instance", "targets.0.crn"),
					resource.TestCheckResourceAttr("data.ibm_atracker_targets.atracker_targets_instance", "targets.0.target_type", targetTargetType),
					resource.TestCheckResourceAttr("data.ibm_atracker_targets.atracker_targets_instance", "targets.0.region", targetRegion),
					resource.TestCheckResourceAttrSet("data.ibm_atracker_targets.atracker_targets_instance", "targets.0.created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_atracker_targets.atracker_targets_instance", "targets.0.updated_at"),
					resource.TestCheckResourceAttr("data.ibm_atracker_targets.atracker_targets_instance", "targets.0.api_version", "2"),
					resource.TestCheckResourceAttr("data.ibm_atracker_targets.atracker_targets_instance", "targets.0.managed_by", targetManagedBy),
				),
			},
		},
	})
}

func testAccCheckIBMAtrackerTargetsDataSourceConfigBasic(targetName string, targetTargetType string) string {
	return fmt.Sprintf(`
		resource "ibm_atracker_target" "atracker_target_instance" {
			name = "%s"
			target_type = "%s"
			region = "us-south"
			cos_endpoint {
				endpoint = "s3.private.us-east.cloud-object-storage.appdomain.cloud"
				target_crn = "crn:v1:bluemix:public:cloud-object-storage:global:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::"
				bucket = "my-atracker-bucket"
				api_key = "%s" // pragma: allowlist secret
				service_to_service_enabled = true
			}
			managed_by = "enterprise"
		}

		data "ibm_atracker_targets" "atracker_targets_instance" {
			name = ibm_atracker_target.atracker_target_instance.name
		}
	`, targetName, targetTargetType, acc.COSApiKey)
}

func testAccCheckIBMAtrackerTargetsDataSourceConfig(targetName string, targetTargetType string, targetRegion string, targetManagedBy string) string {
	return fmt.Sprintf(`
		resource "ibm_atracker_target" "atracker_target_instance" {
			name = "%s"
			target_type = "%s"
			region = "%s"
			cos_endpoint {
				endpoint = "s3.private.us-east.cloud-object-storage.appdomain.cloud"
				target_crn = "crn:v1:bluemix:public:cloud-object-storage:global:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::"
				bucket = "my-atracker-bucket"
				api_key = "%s" // pragma: allowlist secret
				service_to_service_enabled = true
			}
			eventstreams_endpoint {
				target_crn = "crn:v1:bluemix:public:messagehub:us-south:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::"
				brokers = [ "kafka-x:9094" ]
				topic = "my-topic"
				api_key = "%s" // pragma: allowlist secret
				service_to_service_enabled = false
			}
			cloudlogs_endpoint {
				target_crn = "crn:v1:bluemix:public:logs:eu-es:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::"
			}
			managed_by = "%s"
		}

		data "ibm_atracker_targets" "atracker_targets_instance" {
			name = ibm_atracker_target.atracker_target_instance.name
		}
	`, targetName, targetTargetType, targetRegion, acc.COSApiKey, acc.IesApiKey, targetManagedBy)
}

func TestDataSourceIBMAtrackerTargetsTargetToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		cosEndpointModel := make(map[string]interface{})
		cosEndpointModel["endpoint"] = "s3.private.us-east.cloud-object-storage.appdomain.cloud"
		cosEndpointModel["target_crn"] = "crn:v1:bluemix:public:cloud-object-storage:global:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::"
		cosEndpointModel["bucket"] = "my-atracker-bucket"
		cosEndpointModel["service_to_service_enabled"] = true

		eventstreamsEndpointModel := make(map[string]interface{})
		eventstreamsEndpointModel["target_crn"] = "crn:v1:bluemix:public:messagehub:us-south:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::"
		eventstreamsEndpointModel["brokers"] = []string{"kafka-x:9094"}
		eventstreamsEndpointModel["topic"] = "my-topic"
		eventstreamsEndpointModel["api_key"] = "xxxxxxxxxxxxxx"
		eventstreamsEndpointModel["service_to_service_enabled"] = false

		cloudLogsEndpointModel := make(map[string]interface{})
		cloudLogsEndpointModel["target_crn"] = "crn:v1:bluemix:public:eu-es:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::"

		writeStatusModel := make(map[string]interface{})
		writeStatusModel["status"] = "success"
		writeStatusModel["last_failure"] = "2021-05-18T20:15:12.353Z"
		writeStatusModel["reason_for_last_failure"] = "Provided API key could not be found"

		model := make(map[string]interface{})
		model["id"] = "f7dcfae6-e7c5-08ca-451b-fdfa696c9bb6"
		model["name"] = "a-cos-target-us-south"
		model["crn"] = "crn:v1:bluemix:public:atracker:us-south:a/11111111111111111111111111111111:b6eec08b-5201-08ca-451b-cd71523e3626:target:f7dcfae6-e7c5-08ca-451b-fdfa696c9bb6"
		model["target_type"] = "cloud_object_storage"
		model["region"] = "us-south"
		model["cos_endpoint"] = []map[string]interface{}{cosEndpointModel}
		model["eventstreams_endpoint"] = []map[string]interface{}{eventstreamsEndpointModel}
		model["cloudlogs_endpoint"] = []map[string]interface{}{cloudLogsEndpointModel}
		model["write_status"] = []map[string]interface{}{writeStatusModel}
		model["created_at"] = "2021-05-18T20:15:12.353Z"
		model["updated_at"] = "2021-05-18T20:15:12.353Z"
		model["message"] = "This is a valid target. However, there is another target already defined with the same target endpoint."
		model["api_version"] = int(2)
		model["managed_by"] = "enterprise"

		assert.Equal(t, result, model)
	}

	cosEndpointModel := new(atrackerv2.CosEndpoint)
	cosEndpointModel.Endpoint = core.StringPtr("s3.private.us-east.cloud-object-storage.appdomain.cloud")
	cosEndpointModel.TargetCRN = core.StringPtr("crn:v1:bluemix:public:cloud-object-storage:global:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::")
	cosEndpointModel.Bucket = core.StringPtr("my-atracker-bucket")
	cosEndpointModel.ServiceToServiceEnabled = core.BoolPtr(true)

	eventstreamsEndpointModel := new(atrackerv2.EventstreamsEndpoint)
	eventstreamsEndpointModel.TargetCRN = core.StringPtr("crn:v1:bluemix:public:messagehub:us-south:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::")
	eventstreamsEndpointModel.Brokers = []string{"kafka-x:9094"}
	eventstreamsEndpointModel.Topic = core.StringPtr("my-topic")
	eventstreamsEndpointModel.APIKey = core.StringPtr("xxxxxxxxxxxxxx")
	eventstreamsEndpointModel.ServiceToServiceEnabled = core.BoolPtr(false)

	cloudLogsEndpointModel := new(atrackerv2.CloudLogsEndpoint)
	cloudLogsEndpointModel.TargetCRN = core.StringPtr("crn:v1:bluemix:public:eu-es:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::")

	writeStatusModel := new(atrackerv2.WriteStatus)
	writeStatusModel.Status = core.StringPtr("success")
	writeStatusModel.LastFailure = CreateMockDateTime("2021-05-18T20:15:12.353Z")
	writeStatusModel.ReasonForLastFailure = core.StringPtr("Provided API key could not be found")

	model := new(atrackerv2.Target)
	model.ID = core.StringPtr("f7dcfae6-e7c5-08ca-451b-fdfa696c9bb6")
	model.Name = core.StringPtr("a-cos-target-us-south")
	model.CRN = core.StringPtr("crn:v1:bluemix:public:atracker:us-south:a/11111111111111111111111111111111:b6eec08b-5201-08ca-451b-cd71523e3626:target:f7dcfae6-e7c5-08ca-451b-fdfa696c9bb6")
	model.TargetType = core.StringPtr("cloud_object_storage")
	model.Region = core.StringPtr("us-south")
	model.CosEndpoint = cosEndpointModel
	model.EventstreamsEndpoint = eventstreamsEndpointModel
	model.CloudlogsEndpoint = cloudLogsEndpointModel
	model.WriteStatus = writeStatusModel
	model.CreatedAt = CreateMockDateTime("2021-05-18T20:15:12.353Z")
	model.UpdatedAt = CreateMockDateTime("2021-05-18T20:15:12.353Z")
	model.Message = core.StringPtr("This is a valid target. However, there is another target already defined with the same target endpoint.")
	model.APIVersion = core.Int64Ptr(int64(2))
	model.ManagedBy = core.StringPtr("enterprise")

	result, err := atracker.DataSourceIBMAtrackerTargetsTargetToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMAtrackerTargetsCosEndpointToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["endpoint"] = "s3.private.us-east.cloud-object-storage.appdomain.cloud"
		model["target_crn"] = "crn:v1:bluemix:public:cloud-object-storage:global:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::"
		model["bucket"] = "my-atracker-bucket"
		model["service_to_service_enabled"] = true

		assert.Equal(t, result, model)
	}

	model := new(atrackerv2.CosEndpoint)
	model.Endpoint = core.StringPtr("s3.private.us-east.cloud-object-storage.appdomain.cloud")
	model.TargetCRN = core.StringPtr("crn:v1:bluemix:public:cloud-object-storage:global:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::")
	model.Bucket = core.StringPtr("my-atracker-bucket")
	model.ServiceToServiceEnabled = core.BoolPtr(true)

	result, err := atracker.DataSourceIBMAtrackerTargetsCosEndpointToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMAtrackerTargetsEventstreamsEndpointToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["target_crn"] = "crn:v1:bluemix:public:messagehub:us-south:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::"
		model["brokers"] = []string{"kafka-x:9094"}
		model["topic"] = "my-topic"
		model["api_key"] = "xxxxxxxxxxxxxx"
		model["service_to_service_enabled"] = false

		assert.Equal(t, result, model)
	}

	model := new(atrackerv2.EventstreamsEndpoint)
	model.TargetCRN = core.StringPtr("crn:v1:bluemix:public:messagehub:us-south:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::")
	model.Brokers = []string{"kafka-x:9094"}
	model.Topic = core.StringPtr("my-topic")
	model.APIKey = core.StringPtr("xxxxxxxxxxxxxx")
	model.ServiceToServiceEnabled = core.BoolPtr(false)

	result, err := atracker.DataSourceIBMAtrackerTargetsEventstreamsEndpointToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMAtrackerTargetsCloudLogsEndpointToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["target_crn"] = "crn:v1:bluemix:public:eu-es:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::"

		assert.Equal(t, result, model)
	}

	model := new(atrackerv2.CloudLogsEndpoint)
	model.TargetCRN = core.StringPtr("crn:v1:bluemix:public:eu-es:a/11111111111111111111111111111111:22222222-2222-2222-2222-222222222222::")

	result, err := atracker.DataSourceIBMAtrackerTargetsCloudLogsEndpointToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIBMAtrackerTargetsWriteStatusToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["status"] = "success"
		model["last_failure"] = "2021-05-18T20:15:12.353Z"
		model["reason_for_last_failure"] = "Provided API key could not be found"

		assert.Equal(t, result, model)
	}

	model := new(atrackerv2.WriteStatus)
	model.Status = core.StringPtr("success")
	model.LastFailure = CreateMockDateTime("2021-05-18T20:15:12.353Z")
	model.ReasonForLastFailure = core.StringPtr("Provided API key could not be found")

	result, err := atracker.DataSourceIBMAtrackerTargetsWriteStatusToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
