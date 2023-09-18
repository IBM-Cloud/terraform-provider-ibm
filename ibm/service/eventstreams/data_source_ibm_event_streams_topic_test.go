// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package eventstreams_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

var (
	MZREnterpriseInstanceName = "ES Preprod Pipeline MZR"
	SZREnterpriseInstanceName = "mh-preprod-customer-us-south-szr"
	standardInstanceName      = "hyperion-preprod-spp09-us-south-a-service"
	topicName                 = "kafka-java-console-sample-topic"
)

func TestAccIBMEventStreamsTopicDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMEventStreamsTopicDataSourceConfigBasic(MZREnterpriseInstanceName, topicName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_resource_instance.es_instance", "extensions.kafka_brokers_sasl.0"),
					resource.TestCheckResourceAttrSet("data.ibm_resource_instance.es_instance", "extensions.kafka_http_url"),
					resource.TestCheckResourceAttrSet("data.ibm_event_streams_topic.es_topic", "id"),
					resource.TestCheckResourceAttr("data.ibm_event_streams_topic.es_topic", "name", topicName),
					resource.TestCheckResourceAttrSet("data.ibm_event_streams_topic.es_topic", "kafka_brokers_sasl.0"),
					resource.TestCheckResourceAttrSet("data.ibm_event_streams_topic.es_topic", "kafka_http_url"),
				),
			},
			{
				Config: testAccCheckIBMEventStreamsTopicDataSourceConfigBasic(SZREnterpriseInstanceName, topicName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_resource_instance.es_instance", "extensions.kafka_brokers_sasl.0"),
					resource.TestCheckResourceAttrSet("data.ibm_resource_instance.es_instance", "extensions.kafka_http_url"),
					resource.TestCheckResourceAttrSet("data.ibm_event_streams_topic.es_topic", "id"),
					resource.TestCheckResourceAttr("data.ibm_event_streams_topic.es_topic", "name", topicName),
					resource.TestCheckResourceAttrSet("data.ibm_event_streams_topic.es_topic", "kafka_brokers_sasl.0"),
					resource.TestCheckResourceAttrSet("data.ibm_event_streams_topic.es_topic", "kafka_http_url"),
				),
			},
			{
				Config: testAccCheckIBMEventStreamsTopicDataSourceConfigBasic(standardInstanceName, topicName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_resource_instance.es_instance", "extensions.kafka_brokers_sasl.0"),
					resource.TestCheckResourceAttrSet("data.ibm_resource_instance.es_instance", "extensions.kafka_http_url"),
					resource.TestCheckResourceAttrSet("data.ibm_event_streams_topic.es_topic", "id"),
					resource.TestCheckResourceAttr("data.ibm_event_streams_topic.es_topic", "name", topicName),
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
