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

func TestAccIBMEnSubscriptionsDataSourceBasic(t *testing.T) {
	instanceName := fmt.Sprintf("tf_instance_%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	description := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMEnSubscriptionsDataSourceConfigBasic(instanceName, name, description),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_en_subscriptions.data_subscription_4", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_en_subscriptions.data_subscription_4", "instance_guid"),
					resource.TestCheckResourceAttrSet("data.ibm_en_subscriptions.data_subscription_4", "total_count"),
					resource.TestCheckResourceAttrSet("data.ibm_en_subscriptions.data_subscription_4", "subscriptions.#"),
					resource.TestCheckResourceAttrSet("data.ibm_en_subscriptions.data_subscription_4", "subscriptions.0.id"),
					resource.TestCheckResourceAttrSet("data.ibm_en_subscriptions.data_subscription_4", "subscriptions.0.name"),
					resource.TestCheckResourceAttrSet("data.ibm_en_subscriptions.data_subscription_4", "subscriptions.0.description"),
					resource.TestCheckResourceAttrSet("data.ibm_en_subscriptions.data_subscription_4", "subscriptions.0.destination_id"),
					resource.TestCheckResourceAttrSet("data.ibm_en_subscriptions.data_subscription_4", "subscriptions.0.topic_id"),
					resource.TestCheckResourceAttrSet("data.ibm_en_subscriptions.data_subscription_4", "subscriptions.0.topic_name"),
					resource.TestCheckResourceAttrSet("data.ibm_en_subscriptions.data_subscription_4", "subscriptions.0.updated_at"),
				),
			},
		},
	})
}

func testAccCheckIBMEnSubscriptionsDataSourceConfigBasic(instanceName, name, description string) string {
	return fmt.Sprintf(`
	resource "ibm_resource_instance" "en_subscription_datasource_1" {
		name     = "%s"
		location = "us-south"
		plan     = "standard"
		service  = "event-notifications"
	}
	
	resource "ibm_en_topic" "en_topic_resource_6" {
		instance_guid = ibm_resource_instance.en_subscription_datasource_1.guid
		name        = "tf_topic_name_0664"
		description = "tf_topic_description_0455"
	}
	
	resource "ibm_en_destination" "en_destination_resource_6" {
		instance_guid = ibm_resource_instance.en_subscription_datasource_1.guid
		name        = "tf_destination_name_02944"
		type        = "webhook"
		description = "tf_destinatios_description_0364"
		config {
			params {
				verb = "POST"
				url  = "https://demo.webhook.com"
			}
		}
	}
	
	resource "ibm_en_subscription" "en_subscription_resource_6" {
		name           = "%s"
		description 	 = "%s"
		instance_guid    = ibm_resource_instance.en_subscription_datasource_1.guid
		topic_id       = ibm_en_topic.en_topic_resource_6.topic_id
		destination_id = ibm_en_destination.en_destination_resource_6.destination_id
		attributes {
			add_notification_payload = true
			signing_enabled          = true
		}
	}

	data "ibm_en_subscriptions" "data_subscription_4" {
		instance_guid     = ibm_en_subscription.en_subscription_resource_6.instance_guid
	}

	`, instanceName, name, description)
}
