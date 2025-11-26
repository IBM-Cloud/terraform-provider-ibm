// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package eventstreams_test

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
	"strings"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
	"gotest.tools/assert"

	"github.com/IBM-Cloud/bluemix-go/models"
)

func TestAccIBMEventStreamsTopicResourceBasic(t *testing.T) {
	instanceName := fmt.Sprintf("terraform_support_%d", acctest.RandInt())
	planID := "standard"
	serviceName := "messagehub"
	location := "us-south"
	topicName := fmt.Sprintf("es_topic_%d", acctest.RandInt())
	partitions := 1
	cleanupPolicy := "compact,delete"
	retentionBytes := 10485760
	retentionMs := 3600000
	segmentBytes := 10485760
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMEventStreamsInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMEventStreamsTopicWithoutConfig(instanceName, serviceName, planID, location, topicName, partitions),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMEventStreamsTopicExists("ibm_event_streams_topic.es_topic", topicName),
					resource.TestCheckResourceAttr("ibm_resource_instance.es_instance", "name", instanceName),
					resource.TestCheckResourceAttr("ibm_resource_instance.es_instance", "service", serviceName),
					resource.TestCheckResourceAttr("ibm_resource_instance.es_instance", "plan", planID),
					resource.TestCheckResourceAttr("ibm_resource_instance.es_instance", "location", location),
					resource.TestCheckResourceAttrSet("ibm_event_streams_topic.es_topic", "id"),
					resource.TestCheckResourceAttr("ibm_event_streams_topic.es_topic", "name", topicName),
					resource.TestCheckResourceAttr("ibm_event_streams_topic.es_topic", "partitions", strconv.Itoa(partitions)),
				),
			},
			{
				Config: testAccCheckIBMEventStreamsTopicWithConfig(instanceName, serviceName, planID, location, topicName, partitions, cleanupPolicy, retentionBytes, retentionMs, segmentBytes),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMEventStreamsTopicExists("ibm_event_streams_topic.es_topic", topicName),
					resource.TestCheckResourceAttr("ibm_resource_instance.es_instance", "name", instanceName),
					resource.TestCheckResourceAttr("ibm_resource_instance.es_instance", "service", serviceName),
					resource.TestCheckResourceAttr("ibm_resource_instance.es_instance", "plan", planID),
					resource.TestCheckResourceAttr("ibm_resource_instance.es_instance", "location", location),
					resource.TestCheckResourceAttrSet("ibm_event_streams_topic.es_topic", "id"),
					resource.TestCheckResourceAttr("ibm_event_streams_topic.es_topic", "name", topicName),
					resource.TestCheckResourceAttr("ibm_event_streams_topic.es_topic", "partitions", strconv.Itoa(partitions)),
					resource.TestCheckResourceAttr("ibm_event_streams_topic.es_topic", "config.cleanup.policy", cleanupPolicy),
					resource.TestCheckResourceAttr("ibm_event_streams_topic.es_topic", "config.retention.bytes", strconv.Itoa(retentionBytes)),
					resource.TestCheckResourceAttr("ibm_event_streams_topic.es_topic", "config.retention.ms", strconv.Itoa(retentionMs)),
					resource.TestCheckResourceAttr("ibm_event_streams_topic.es_topic", "config.segment.bytes", strconv.Itoa(segmentBytes)),
				),
			},
		},
	})
}

func TestAccIBMEventStreamsTopicResourceWithExistingInstance(t *testing.T) {
	topicName := fmt.Sprintf("es_topic_%d", acctest.RandInt())
	partitions := 1
	cleanupPolicy := "compact,delete"
	retentionBytes := 10485760
	retentionMs := 3600000
	segmentBytes := 10485760
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMEventStreamsTopicWithExistingInstanceWithoutConfig(getTestInstanceName(stdKey), topicName, partitions),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMEventStreamsTopicExists("ibm_event_streams_topic.es_topic", topicName),
					resource.TestCheckResourceAttrSet("data.ibm_resource_instance.es_instance", "extensions.kafka_brokers_sasl.0"),
					resource.TestCheckResourceAttrSet("data.ibm_resource_instance.es_instance", "extensions.kafka_http_url"),
					resource.TestCheckResourceAttrSet("ibm_event_streams_topic.es_topic", "id"),
					resource.TestCheckResourceAttrSet("ibm_event_streams_topic.es_topic", "kafka_brokers_sasl.0"),
					resource.TestCheckResourceAttrSet("ibm_event_streams_topic.es_topic", "kafka_http_url"),
					resource.TestCheckResourceAttrSet("ibm_event_streams_topic.es_topic", "id"),
					resource.TestCheckResourceAttr("ibm_event_streams_topic.es_topic", "name", topicName),
					resource.TestCheckResourceAttr("ibm_event_streams_topic.es_topic", "partitions", strconv.Itoa(partitions)),
				),
			},
			{
				Config: testAccCheckIBMEventStreamsTopicWithExistingInstanceWithConfig(getTestInstanceName(stdKey), topicName, partitions, cleanupPolicy, retentionBytes, retentionMs, segmentBytes),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMEventStreamsTopicExists("ibm_event_streams_topic.es_topic", topicName),
					resource.TestCheckResourceAttrSet("data.ibm_resource_instance.es_instance", "extensions.kafka_brokers_sasl.0"),
					resource.TestCheckResourceAttrSet("data.ibm_resource_instance.es_instance", "extensions.kafka_http_url"),
					resource.TestCheckResourceAttrSet("ibm_event_streams_topic.es_topic", "id"),
					resource.TestCheckResourceAttrSet("ibm_event_streams_topic.es_topic", "kafka_brokers_sasl.0"),
					resource.TestCheckResourceAttrSet("ibm_event_streams_topic.es_topic", "kafka_http_url"),
					resource.TestCheckResourceAttr("ibm_event_streams_topic.es_topic", "name", topicName),
					resource.TestCheckResourceAttr("ibm_event_streams_topic.es_topic", "partitions", strconv.Itoa(partitions)),
					resource.TestCheckResourceAttr("ibm_event_streams_topic.es_topic", "config.cleanup.policy", cleanupPolicy),
					resource.TestCheckResourceAttr("ibm_event_streams_topic.es_topic", "config.retention.bytes", strconv.Itoa(retentionBytes)),
					resource.TestCheckResourceAttr("ibm_event_streams_topic.es_topic", "config.retention.ms", strconv.Itoa(retentionMs)),
					resource.TestCheckResourceAttr("ibm_event_streams_topic.es_topic", "config.segment.bytes", strconv.Itoa(segmentBytes)),
				),
			},
		},
	})
}

func TestAccIBMEventStreamsTopicImport(t *testing.T) {
	instanceName := fmt.Sprintf("terraform_support_%d", acctest.RandInt())
	planID := "standard"
	serviceName := "messagehub"
	location := "us-south"
	topicName := fmt.Sprintf("es_topic_%d", acctest.RandInt())
	partitions := 1
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMEventStreamsInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMEventStreamsTopicWithoutConfig(instanceName, serviceName, planID, location, topicName, partitions),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMEventStreamsTopicExists("ibm_event_streams_topic.es_topic", topicName),
					resource.TestCheckResourceAttr("ibm_event_streams_topic.es_topic", "name", topicName),
					resource.TestCheckResourceAttr("ibm_event_streams_topic.es_topic", "partitions", strconv.Itoa(partitions)),
				),
			},
			{
				ResourceName:      "ibm_event_streams_topic.es_topic",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccIBMEventStreamsEnterprise(t *testing.T) {
	instanceName := fmt.Sprintf("terraform_support_%d", acctest.RandInt())
	planID := "enterprise-3nodes-2tb"
	serviceName := "messagehub"
	location := "eu-gb"
	topicName := fmt.Sprintf("es_topic_%d", acctest.RandInt())
	partitions := 1
	parameters := map[string]string{
		"service-endpoints": "public",
		"throughput":        "150",
		"storage_size":      "2048",
		"kms_key_crn":       "crn:v1:staging:public:kms:us-south:a/6db1b0d0b5c54ee5c201552547febcd8:0aa69b09-941b-41b2-bbf9-9f9f0f6a6f79:key:dd37a0b6-eff4-4708-8459-e29ae0a8f256", //preprod-byok-customer-key from KMS instance keyprotect-preprod-customer-keys
		//ES Preprod Pipeline Mirror E2E
		"source_crn":   "crn:v1:staging:public:messagehub:eu-gb:a/6db1b0d0b5c54ee5c201552547febcd8:4414aa6e-237a-495d-b0b1-1b69d64d874b::",
		"target_alias": "target-cluster",
		"source_alias": "source-cluster",
	}
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMEventStreamsInstanceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMEventStreamsEnterpriseWithParameters(instanceName, serviceName, planID, location, topicName, partitions, parameters),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMEventStreamsTopicExists("ibm_event_streams_topic.es_topic", topicName),
					resource.TestCheckResourceAttr("ibm_resource_instance.es_instance", "name", instanceName),
					resource.TestCheckResourceAttr("ibm_resource_instance.es_instance", "service", serviceName),
					resource.TestCheckResourceAttr("ibm_resource_instance.es_instance", "plan", planID),
					resource.TestCheckResourceAttr("ibm_resource_instance.es_instance", "location", location),
					resource.TestCheckResourceAttrSet("ibm_event_streams_topic.es_topic", "id"),
					resource.TestCheckResourceAttr("ibm_event_streams_topic.es_topic", "name", topicName),
					resource.TestCheckResourceAttr("ibm_event_streams_topic.es_topic", "partitions", strconv.Itoa(partitions)),
					testAccCheckIBMEventStreamsMirroringConfigProperties("ibm_event_streams_mirroring_config.es_mirroring_config"),
					resource.TestCheckResourceAttr("ibm_event_streams_mirroring_config.es_mirroring_config", "mirroring_topic_patterns.#", "1"),
					resource.TestCheckResourceAttr("ibm_event_streams_mirroring_config.es_mirroring_config", "mirroring_topic_patterns.0", ".*"),
				),
			},
		},
	})
}

func testAccCheckIBMEventStreamsTopicWithoutConfig(instanceName, serviceName, planID, location,
	topicName string, partitions int) string {
	return createPlatformResources(instanceName, serviceName, planID, location, nil) + "\n" +
		createEventStreamsTopicResourceWithoutConfig(true, topicName, partitions)
}

func testAccCheckIBMEventStreamsTopicWithConfig(instanceName, serviceName, planID, location,
	topicName string, partitions int, cleanupPolicy string, retentionBytes int, retentionMs int, segmentBytes int) string {
	return createPlatformResources(instanceName, serviceName, planID, location, nil) + "\n" +
		createEventStreamsTopicResourceWithConfig(true, topicName, partitions, cleanupPolicy, retentionBytes, retentionMs, segmentBytes)
}

func testAccCheckIBMEventStreamsEnterpriseWithParameters(instanceName, serviceName, planID, location, topicName string, partitions int, params map[string]string) string {
	return createPlatformResources(instanceName, serviceName, planID, location, params) + "\n" +
		createEventStreamsTopicResourceWithoutConfig(true, topicName, partitions) + "\n" +
		createEventStreamsMirroringConfig(serviceName, params["source_crn"])
}

func testAccCheckIBMEventStreamsTopicWithExistingInstanceWithoutConfig(instanceName,
	topicName string, partitions int) string {
	return getPlatformResource(instanceName) + "\n" +
		createEventStreamsTopicResourceWithoutConfig(false, topicName, partitions)
}

func testAccCheckIBMEventStreamsTopicWithExistingInstanceWithConfig(instanceName,
	topicName string, partitions int, cleanupPolicy string, retentionBytes int, retentionMs int, segmentBytes int) string {
	return getPlatformResource(instanceName) + "\n" +
		createEventStreamsTopicResourceWithConfig(false, topicName, partitions, cleanupPolicy, retentionBytes, retentionMs, segmentBytes)
}

func getPlatformResource(instanceName string) string {
	return fmt.Sprintf(`
	  data "ibm_resource_group" "group" {
		is_default=true
	  }
	  data "ibm_resource_instance" "es_instance" {
		resource_group_id = data.ibm_resource_group.group.id
		name              = "%s"
	  }`, instanceName)
}

func createPlatformResources(instanceName, serviceName, planID, location string, parameters map[string]string) string {
	if planID == "standard" || planID == "lite" {
		return fmt.Sprintf(`
		data "ibm_resource_group" "group" {
		  is_default=true
		}
		resource "ibm_resource_instance" "es_instance" {
		  name              = "%s"
		  service           = "%s"
		  plan              = "%s"
		  location          = "%s"
		  resource_group_id = data.ibm_resource_group.group.id
		}`, instanceName, serviceName, planID, location)
	}
	// create enterprise instance
	return fmt.Sprintf(`
data "ibm_resource_group" "group" {
  is_default = true
}
resource "ibm_resource_instance" "es_instance" {
  name              = "%s"
  service           = "%s"
  plan              = "%s"
  location          = "%s"
  resource_group_id = data.ibm_resource_group.group.id
  parameters_json = jsonencode(
    {
      service-endpoints = "%s"
      throughput        = "%s"
      storage_size      = "%s"
      kms_key_crn       = "%s"
      mirroring = {
        source_crn   = "%s"
        source_alias = "%s"
        target_alias = "%s"
      }
    }
  )
  timeouts {
    create = "3h"
    update = "1h"
    delete = "15m"
  }
}`, instanceName, serviceName, planID, location,
		parameters["service-endpoints"], parameters["throughput"], parameters["storage_size"], parameters["kms_key_crn"], parameters["source_crn"], parameters["source_alias"], parameters["target_alias"])
}

func createEventStreamsMirroringConfig(serviceName, targetCrn string) string {
	//get target guid from target crn
	targetGuid := strings.Split(targetCrn, ":")[7]
	return fmt.Sprintf(`
resource "ibm_iam_authorization_policy" "instance_policy" {
  source_service_name         = "%s"
  source_resource_instance_id = ibm_resource_instance.es_instance.guid
  target_service_name         = "%s"
  target_resource_instance_id = "%s"
  roles                       = ["Reader"]
  description                 = "test mirroring setup via terraform"
}

resource "ibm_event_streams_mirroring_config" "es_mirroring_config" {
  resource_instance_id     = ibm_resource_instance.es_instance.id
  mirroring_topic_patterns = [".*"]
}`, serviceName, serviceName, targetGuid)
}

func createEventStreamsTopicResourceWithoutConfig(createInstance bool, topicName string, partitions int) string {
	var resourceInstanceID string
	if createInstance {
		resourceInstanceID = "ibm_resource_instance.es_instance.id"
	} else {
		resourceInstanceID = "data.ibm_resource_instance.es_instance.id"
	}
	return fmt.Sprintf(`
		resource "ibm_event_streams_topic" "es_topic" {
		  resource_instance_id 	= %s
		  name            		= "%s"
		  partitions      		= %d
		}`, resourceInstanceID, topicName, partitions)

}

func createEventStreamsTopicResourceWithConfig(createInstance bool, topicName string, partitions int, cleanupPolicy string, retentionBytes int, retentionMs int, segmentBytes int) string {
	var resourceInstanceID string
	if createInstance {
		resourceInstanceID = "ibm_resource_instance.es_instance.id"
	} else {
		resourceInstanceID = "data.ibm_resource_instance.es_instance.id"
	}
	return fmt.Sprintf(`
		resource "ibm_event_streams_topic" "es_topic" {
			resource_instance_id = %s
			name                 = "%s"
			partitions           = %d
			config = {
			  "cleanup.policy"  = "%s"
			  "retention.bytes" = %d
			  "retention.ms"    = %d
			  "segment.bytes"   = %d
			}
		}`, resourceInstanceID, topicName, partitions, cleanupPolicy, retentionBytes, retentionMs, segmentBytes)
}

func testAccCheckIBMEventStreamsInstanceDestroy(s *terraform.State) error {
	rsContClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).ResourceControllerAPIV2()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_resource_instance" && rs.Type != "ibm_iam_authorization_policy" {
			continue
		}
		// adding authotization policy check for mirroring instance
		if rs.Type == "ibm_iam_authorization_policy" {
			iamPolicyManagementClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).IAMPolicyManagementV1API()
			if err != nil {
				return err
			}
			authPolicyID := rs.Primary.ID

			getPolicyOptions := iamPolicyManagementClient.NewGetPolicyOptions(
				authPolicyID,
			)
			destroyedPolicy, response, err := iamPolicyManagementClient.GetPolicy(getPolicyOptions)

			if err == nil && *destroyedPolicy.State != "deleted" {
				return fmt.Errorf("Authorization policy still exists: %s\n", rs.Primary.ID)
			} else if response.StatusCode != 404 && destroyedPolicy.State != nil && *destroyedPolicy.State != "deleted" {
				return fmt.Errorf("[ERROR] Error waiting for authorization policy (%s) to be destroyed: %s", rs.Primary.ID, err)
			}
		}
		if rs.Type == "ibm_resource_instance" {
			instanceID := rs.Primary.ID
			instance, err := rsContClient.ResourceServiceInstanceV2().GetInstance(instanceID)
			if err != nil {
				if strings.Contains(err.Error(), "404") {
					return nil
				}
				return fmt.Errorf("[ERROR] Error checking if instance (%s) has been destroyed: %s", rs.Primary.ID, err)
			}

			if !reflect.DeepEqual(instance, models.ServiceInstanceV2{}) &&
				instance.State != "removed" && instance.State != "pending_reclamation" {
				return fmt.Errorf("[ERROR] Instance (%s) is not removed", rs.Primary.ID)
			}
		}
	}
	return nil
}

func testAccCheckIBMEventStreamsTopicExists(n, topicName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		log.Printf("testAccCheckIBMEventStreamsTopicExists")
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}
		topicID := rs.Primary.ID
		if topicID == "" {
			return fmt.Errorf("[ERROR] No topic ID is set")
		}
		if strings.HasSuffix(topicID, topicName) {
			return nil
		}
		return fmt.Errorf("topic %s not found", topicName)
	}
}

var (
	instanceCRN = "crn:v1:staging:public:messagehub:us-south:a/6db1b0d0b5c54ee5c201552547febcd8:c822a30e-bfff-4867-85ec-b805eeab1835::"
	topicID     = "crn:v1:staging:public:messagehub:us-south:a/6db1b0d0b5c54ee5c201552547febcd8:c822a30e-bfff-4867-85ec-b805eeab1835:topic:" + getTestTopicName()
)

func TestGetTopicID(t *testing.T) {
	gotTopicID := getTopicID(instanceCRN, getTestTopicName())
	assert.Equal(t, topicID, gotTopicID)
}

func TestGetInstanceCRN(t *testing.T) {
	gotInstanceCRN := getInstanceCRN(topicID)
	assert.Equal(t, instanceCRN, gotInstanceCRN)
}

func TestGetTopicName(t *testing.T) {
	gotTopicName := getTopicName(topicID)
	assert.Equal(t, getTestTopicName(), gotTopicName)
}

func getTopicID(instanceCRN string, topicName string) string {
	crnSegments := strings.Split(instanceCRN, ":")
	crnSegments[8] = "topic"
	crnSegments[9] = topicName
	return strings.Join(crnSegments, ":")
}

func getTopicName(topicID string) string {
	return strings.Split(topicID, ":")[9]
}

func getInstanceCRN(topicID string) string {
	crnSegments := strings.Split(topicID, ":")
	crnSegments[8] = ""
	crnSegments[9] = ""
	return strings.Join(crnSegments, ":")
}
