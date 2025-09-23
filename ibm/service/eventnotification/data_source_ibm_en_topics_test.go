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

func TestAccIBMEnTopicsDataSourceBasic(t *testing.T) {
	instanceName := fmt.Sprintf("tf_instance_%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	description := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMEnTopicsDataSourceConfigBasic(instanceName, name, description),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_en_topics.en_topic_datasource_6", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_en_topics.en_topic_datasource_6", "instance_guid"),
					resource.TestCheckResourceAttrSet("data.ibm_en_topics.en_topic_datasource_6", "total_count"),
					resource.TestCheckResourceAttrSet("data.ibm_en_topics.en_topic_datasource_6", "topics.#"),
					resource.TestCheckResourceAttrSet("data.ibm_en_topics.en_topic_datasource_6", "topics.0.id"),
					resource.TestCheckResourceAttrSet("data.ibm_en_topics.en_topic_datasource_6", "topics.0.name"),
					resource.TestCheckResourceAttrSet("data.ibm_en_topics.en_topic_datasource_6", "topics.0.description"),
					resource.TestCheckResourceAttrSet("data.ibm_en_topics.en_topic_datasource_6", "topics.0.source_count"),
					resource.TestCheckResourceAttrSet("data.ibm_en_topics.en_topic_datasource_6", "topics.0.subscription_count"),
				),
			},
		},
	})
}

func testAccCheckIBMEnTopicsDataSourceConfigBasic(instanceName, name, description string) string {
	return fmt.Sprintf(`
	resource "ibm_resource_instance" "en_topic_datasource_1" {
		name     = "%s"
		location = "us-south"
		plan     = "standard"
		service  = "event-notifications"
	}
	
	resource "ibm_en_topic" "en_topic_datasource_4" {
		name        = "%s"
		description = "%s"
		instance_guid = ibm_resource_instance.en_topic_datasource_1.guid
	}
	
	data "ibm_en_topics" "en_topic_datasource_6" {
		instance_guid = ibm_resource_instance.en_topic_datasource_1.guid
	}
	`, instanceName, name, description)
}
