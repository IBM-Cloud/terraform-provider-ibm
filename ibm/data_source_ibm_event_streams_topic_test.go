/*
* IBM Confidential
* Object Code Only Source Materials
* 5747-SM3
* (c) Copyright IBM Corp. 2017,2021
*
* The source code for this program is not published or otherwise divested
* of its trade secrets, irrespective of what has been deposited with the
* U.S. Copyright Office.
 */

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

var (
	MZREnterpriseInstanceName = "mh-preprod-customer-us-south-wp"
	SZREnterpriseInstanceName = "mh-preprod-customer-us-south-szr"
	standardInstanceName      = "hyperion-preprod-spp-a-service"
	topicName                 = "kafka-java-console-sample-topic"
)

func TestAccIBMEventStreamsTopicDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
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
			resource.TestStep{
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
			resource.TestStep{
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
		name = "Default"
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
