// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package eventnotification_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMEnSlackSubscriptionDataSourceAllArgs(t *testing.T) {
	instanceName := fmt.Sprintf("tf_instance_%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	description := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMEnSlackSubscriptionDataSourceConfig(instanceName, name, description),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_en_subscription_slack.data_subscription_1", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_en_subscription_slack.data_subscription_1", "instance_guid"),
					resource.TestCheckResourceAttrSet("data.ibm_en_subscription_slack.data_subscription_1", "subscription_id"),
					resource.TestCheckResourceAttrSet("data.ibm_en_subscription_slack.data_subscription_1", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_en_subscription_slack.data_subscription_1", "description"),
					resource.TestCheckResourceAttrSet("data.ibm_en_subscription_slack.data_subscription_1", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_en_subscription_slack.data_subscription_1", "destination_type"),
					resource.TestCheckResourceAttrSet("data.ibm_en_subscription_slack.data_subscription_1", "destination_id"),
					resource.TestCheckResourceAttrSet("data.ibm_en_subscription_slack.data_subscription_1", "destination_name"),
					resource.TestCheckResourceAttrSet("data.ibm_en_subscription_slack.data_subscription_1", "topic_id"),
					resource.TestCheckResourceAttrSet("data.ibm_en_subscription_slack.data_subscription_1", "topic_name"),
				),
			},
		},
	})
}

func testAccCheckIBMEnSlackSubscriptionDataSourceConfig(instanceName, name, description string) string {
	return fmt.Sprintf(`
	resource "ibm_resource_instance" "en_subscription_datasource" {
		name     = "%s"
		location = "us-south"
		plan     = "standard"
		service  = "event-notifications"
	}
	
	resource "ibm_en_topic" "en_topic_resource_4" {
		instance_guid = ibm_resource_instance.en_subscription_datasource.guid
		name        = "tf_topic_name_0664"
		description = "tf_topic_description_0455"
	}
	
	resource "ibm_en_destination_slack" "en_destination_resource_4" {
		instance_guid = ibm_resource_instance.en_subscription_datasource.guid
		name        = "tf_destination_name_02944"
		type        = "slack"
		description = "tf_destinatios_description_0364"
		config {
			params {
			    type = "direct_message"
				token  = "dsgdewgfpwfgpewweeffewsgcvsaksdksff"
			}
		}
	}
	
	resource "ibm_en_subscription_slack" "en_subscription_resource_4" {
		name           = "%s"
		description 	 = "%s"
		instance_guid    = ibm_resource_instance.en_subscription_datasource.guid
		topic_id       = ibm_en_topic.en_topic_resource_4.topic_id
		destination_id = ibm_en_destination_slack.en_destination_resource_4.destination_id
		attributes {
		    channels {
			id = "GFDRTUIIOBVD"
			}
			
		}
	}

	data "ibm_en_subscription_slack" "data_subscription_1" {
		instance_guid     = ibm_resource_instance.en_subscription_datasource.guid
		subscription_id = ibm_en_subscription_slack.en_subscription_resource_4.subscription_id
	}

	`, instanceName, name, description)
}
