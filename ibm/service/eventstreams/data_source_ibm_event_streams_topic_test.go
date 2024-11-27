// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package eventstreams_test

import (
	"fmt"
	"os"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

const (
	mzrKey                    = "MZR_INSTANCE_NAME"
	szrKey                    = "SZR_INSTANCE_NAME"
	stdKey                    = "STD_INSTANCE_NAME"
	MZREnterpriseInstanceName = "ES Preprod Pipeline MZR"
	SZREnterpriseInstanceName = "ES Integration Pipeline SZR"
	standardInstanceName      = "hyperion-preprod-spp09-us-south-a-service"
	testTopicName             = "kafka-java-console-sample-topic"
)

func getTestInstanceName(envKey string) string {
	instance := os.Getenv(envKey)
	if len(instance) == 0 {
		switch envKey {
		case mzrKey:
			instance = MZREnterpriseInstanceName
		case szrKey:
			instance = SZREnterpriseInstanceName
		case stdKey:
			instance = standardInstanceName
		}
	}
	return instance
}

func getTestTopicName() string {
	topic := os.Getenv("TEST_TOPIC")
	if len(topic) == 0 {
		topic = testTopicName
	}
	return topic
}

func TestAccIBMEventStreamsTopicDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMEventStreamsTopicDataSourceConfigBasic(getTestInstanceName(mzrKey), getTestTopicName()),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_resource_instance.es_instance", "extensions.kafka_brokers_sasl.0"),
					resource.TestCheckResourceAttrSet("data.ibm_resource_instance.es_instance", "extensions.kafka_http_url"),
					resource.TestCheckResourceAttrSet("data.ibm_event_streams_topic.es_topic", "id"),
					resource.TestCheckResourceAttr("data.ibm_event_streams_topic.es_topic", "name", getTestTopicName()),
					resource.TestCheckResourceAttrSet("data.ibm_event_streams_topic.es_topic", "kafka_brokers_sasl.0"),
					resource.TestCheckResourceAttrSet("data.ibm_event_streams_topic.es_topic", "kafka_http_url"),
				),
			},
			{
				Config: testAccCheckIBMEventStreamsTopicDataSourceConfigBasic(getTestInstanceName(szrKey), getTestTopicName()),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_resource_instance.es_instance", "extensions.kafka_brokers_sasl.0"),
					resource.TestCheckResourceAttrSet("data.ibm_resource_instance.es_instance", "extensions.kafka_http_url"),
					resource.TestCheckResourceAttrSet("data.ibm_event_streams_topic.es_topic", "id"),
					resource.TestCheckResourceAttr("data.ibm_event_streams_topic.es_topic", "name", getTestTopicName()),
					resource.TestCheckResourceAttrSet("data.ibm_event_streams_topic.es_topic", "kafka_brokers_sasl.0"),
					resource.TestCheckResourceAttrSet("data.ibm_event_streams_topic.es_topic", "kafka_http_url"),
				),
			},
			{
				Config: testAccCheckIBMEventStreamsTopicDataSourceConfigBasic(getTestInstanceName(stdKey), getTestTopicName()),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_resource_instance.es_instance", "extensions.kafka_brokers_sasl.0"),
					resource.TestCheckResourceAttrSet("data.ibm_resource_instance.es_instance", "extensions.kafka_http_url"),
					resource.TestCheckResourceAttrSet("data.ibm_event_streams_topic.es_topic", "id"),
					resource.TestCheckResourceAttr("data.ibm_event_streams_topic.es_topic", "name", getTestTopicName()),
					resource.TestCheckResourceAttrSet("data.ibm_event_streams_topic.es_topic", "kafka_brokers_sasl.0"),
					resource.TestCheckResourceAttrSet("data.ibm_event_streams_topic.es_topic", "kafka_http_url"),
				),
			},
		},
	})
}

func testAccCheckIBMEventStreamsTopicDataSourceConfigBasic(instancecName, topicName string) string {
	return fmt.Sprintf(`
	data "ibm_resource_group" "my_group" {
		is_default=true
	  }
	data "ibm_resource_instance" "es_instance" {
		resource_group_id = data.ibm_resource_group.my_group.id
		name              = "%s"
	}
	data "ibm_event_streams_topic" "es_topic" {
		resource_instance_id = data.ibm_resource_instance.es_instance.id
		name                 = "%s"
	}`, instancecName, topicName)
}
