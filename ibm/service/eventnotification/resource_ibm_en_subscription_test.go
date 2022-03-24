// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package eventnotification_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	en "github.com/IBM/event-notifications-go-admin-sdk/eventnotificationsv1"
)

func TestAccIBMEnSubscriptionAllArgs(t *testing.T) {
	var conf en.Subscription
	instanceName := fmt.Sprintf("tf_instance_%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	description := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	newName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	newDescription := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMEnSubscriptionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMEnSubscriptionConfig(instanceName, name, description),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMEnSubscriptionExists("ibm_en_subscription.en_subscription_resource_1", conf),
					resource.TestCheckResourceAttrSet("ibm_en_subscription.en_subscription_resource_1", "name"),
					resource.TestCheckResourceAttrSet("ibm_en_subscription.en_subscription_resource_1", "description"),
					resource.TestCheckResourceAttrSet("ibm_en_subscription.en_subscription_resource_1", "topic_id"),
					resource.TestCheckResourceAttrSet("ibm_en_subscription.en_subscription_resource_1", "updated_at"),
					resource.TestCheckResourceAttrSet("ibm_en_subscription.en_subscription_resource_1", "instance_guid"),
					resource.TestCheckResourceAttrSet("ibm_en_subscription.en_subscription_resource_1", "destination_id"),
					resource.TestCheckResourceAttrSet("ibm_en_subscription.en_subscription_resource_1", "destination_type"),
					resource.TestCheckResourceAttrSet("ibm_en_subscription.en_subscription_resource_1", "subscription_id"),
					resource.TestCheckResourceAttrSet("ibm_en_subscription.en_subscription_resource_1", "attributes.#"),
					resource.TestCheckResourceAttrSet("ibm_en_subscription.en_subscription_resource_1", "attributes.0.signing_enabled"),
					resource.TestCheckResourceAttrSet("ibm_en_subscription.en_subscription_resource_1", "attributes.0.add_notification_payload"),
				),
			},
			{
				Config: testAccCheckIBMEnSubscriptionConfig(instanceName, newName, newDescription),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_en_subscription.en_subscription_resource_1", "name", newName),
					resource.TestCheckResourceAttr("ibm_en_subscription.en_subscription_resource_1", "description", newDescription),
				),
			},
			{
				ResourceName:      "ibm_en_subscription.en_subscription_resource_1",
				ImportState:       true,
				ImportStateVerify: false,
			},
		},
	})
}

func testAccCheckIBMEnSubscriptionConfig(instanceName, name, description string) string {
	return fmt.Sprintf(`
	resource "ibm_resource_instance" "en_subscription_resource" {
		name     = "%s"
		location = "us-south"
		plan     = "standard"
		service  = "event-notifications"
	}
	
	resource "ibm_en_topic" "en_topic_resource_2" {
		instance_guid = ibm_resource_instance.en_subscription_resource.guid
		name        = "tf_topic_name_0234"
		description = "tf_topic_description_0235"
	}
	
	resource "ibm_en_destination" "en_destination_resource_2" {
		instance_guid = ibm_resource_instance.en_subscription_resource.guid
		name        = "tf_destination_name_02983"
		type        = "webhook"
		description = "tf_destinatios_description_0364"
		config {
			params {
				verb = "POST"
				url  = "https://demo.webhook.com"
			}
		}
	}
	
	resource "ibm_en_subscription" "en_subscription_resource_1" {
		name           = "%s"
		description 	 = "%s"
		instance_guid    = ibm_resource_instance.en_subscription_resource.guid
		topic_id       = ibm_en_topic.en_topic_resource_2.topic_id
		destination_id = ibm_en_destination.en_destination_resource_2.destination_id
		attributes {
			signing_enabled          = true
		}
	}
	`, instanceName, name, description)
}

func testAccCheckIBMEnSubscriptionExists(n string, obj en.Subscription) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		enClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).EventNotificationsApiV1()
		if err != nil {
			return err
		}

		options := &en.GetSubscriptionOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		options.SetInstanceID(parts[0])
		options.SetID(parts[1])

		subscription, _, err := enClient.GetSubscription(options)
		if err != nil {
			return err
		}

		obj = *subscription
		return nil
	}
}

func testAccCheckIBMEnSubscriptionDestroy(s *terraform.State) error {
	enClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).EventNotificationsApiV1()
	if err != nil {
		return err
	}

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "en_subscription_resource_1" {
			continue
		}

		options := &en.GetSubscriptionOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		options.SetInstanceID(parts[0])
		options.SetID(parts[1])

		// Try to find the key
		_, response, err := enClient.GetSubscription(options)

		if err == nil {
			return fmt.Errorf("en_subscription still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("[ERROR] Error checking for en_subscription (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
