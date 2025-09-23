// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package eventnotification_test

import (
	"fmt"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMEnTopicDataSourceBasic(t *testing.T) {
	instanceName := fmt.Sprintf("tf_instance_%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	description := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMEnTopicDataSourceConfigBasic(instanceName, name, description),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_en_topic.en_topic_datasource_2", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_en_topic.en_topic_datasource_2", "instance_guid"),
					resource.TestCheckResourceAttrSet("data.ibm_en_topic.en_topic_datasource_2", "topic_id"),
					resource.TestCheckResourceAttrSet("data.ibm_en_topic.en_topic_datasource_2", "description"),
					resource.TestCheckResourceAttrSet("data.ibm_en_topic.en_topic_datasource_2", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_en_topic.en_topic_datasource_2", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_en_topic.en_topic_datasource_2", "source_count"),
					resource.TestCheckResourceAttrSet("data.ibm_en_topic.en_topic_datasource_2", "subscription_count"),
				),
			},
		},
	})
}

func testAccCheckIBMEnTopicDataSourceConfigBasic(instanceName, name, description string) string {
	return fmt.Sprintf(`
	resource "ibm_resource_instance" "en_topic_datasource" {
		name     = "%s"
		location = "us-south"
		plan     = "standard"
		service  = "event-notifications"
	}
	
	resource "ibm_en_topic" "en_topic_datasource_1" {
		name        = "%s"
		description = "%s"
		instance_guid = ibm_resource_instance.en_topic_datasource.guid
	}
	
	data "ibm_en_topic" "en_topic_datasource_2" {
		instance_guid = ibm_resource_instance.en_topic_datasource.guid
		topic_id    = ibm_en_topic.en_topic_datasource_1.topic_id
	}
	`, instanceName, name, description)
}
