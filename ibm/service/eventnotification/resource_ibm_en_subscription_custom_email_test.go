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

func TestAccIBMEnCustomDomainSubscriptionProductionDestination(t *testing.T) {
	var conf en.Subscription
	instanceName := fmt.Sprintf("tf_instance_%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	description := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	newName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	newDescription := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMEnCustomDomainSubscriptionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMEnCustomDomainSubscriptionProductionConfig(instanceName, name, description),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMEnCustomDomainSubscriptionExists("ibm_en_subscription_custom_email.en_subscription_production", conf),
					resource.TestCheckResourceAttrSet("ibm_en_subscription_custom_email.en_subscription_production", "name"),
					resource.TestCheckResourceAttrSet("ibm_en_subscription_custom_email.en_subscription_production", "description"),
					resource.TestCheckResourceAttrSet("ibm_en_subscription_custom_email.en_subscription_production", "topic_id"),
					resource.TestCheckResourceAttrSet("ibm_en_subscription_custom_email.en_subscription_production", "updated_at"),
					resource.TestCheckResourceAttrSet("ibm_en_subscription_custom_email.en_subscription_production", "instance_guid"),
					resource.TestCheckResourceAttrSet("ibm_en_subscription_custom_email.en_subscription_production", "destination_id"),
					resource.TestCheckResourceAttrSet("ibm_en_subscription_custom_email.en_subscription_production", "destination_type"),
					resource.TestCheckResourceAttrSet("ibm_en_subscription_custom_email.en_subscription_production", "subscription_id"),
					resource.TestCheckResourceAttrSet("ibm_en_subscription_custom_email.en_subscription_production", "attributes.#"),
					resource.TestCheckResourceAttrSet("ibm_en_subscription_custom_email.en_subscription_production", "attributes.0.reply_to_mail"),
					resource.TestCheckResourceAttrSet("ibm_en_subscription_custom_email.en_subscription_production", "attributes.0.reply_to_name"),
					resource.TestCheckResourceAttrSet("ibm_en_subscription_custom_email.en_subscription_production", "attributes.0.from_name"),
					resource.TestCheckResourceAttrSet("ibm_en_subscription_custom_email.en_subscription_production", "attributes.0.from_email"),
					resource.TestCheckResourceAttrSet("ibm_en_subscription_custom_email.en_subscription_production", "attributes.0.add_notification_payload"),
					resource.TestCheckResourceAttrSet("ibm_en_subscription_custom_email.en_subscription_production", "attributes.0.invited"),
				),
			},
			{
				Config: testAccCheckIBMEnCustomDomainSubscriptionProductionConfig(instanceName, newName, newDescription),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_en_subscription_custom_email.en_subscription_production", "name", newName),
					resource.TestCheckResourceAttr("ibm_en_subscription_custom_email.en_subscription_production", "description", newDescription),
				),
			},
			{
				ResourceName:      "ibm_en_subscription_custom_email.en_subscription_production",
				ImportState:       true,
				ImportStateVerify: false,
			},
		},
	})
}

func TestAccIBMEnCustomDomainSubscriptionSandboxDestination(t *testing.T) {
	var conf en.Subscription
	instanceName := fmt.Sprintf("tf_instance_%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf_sandbox_sub_%d", acctest.RandIntRange(10, 100))
	description := fmt.Sprintf("tf_sandbox_description_%d", acctest.RandIntRange(10, 100))
	newName := fmt.Sprintf("tf_sandbox_sub_%d", acctest.RandIntRange(10, 100))
	newDescription := fmt.Sprintf("tf_sandbox_updated_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMEnCustomDomainSubscriptionDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMEnCustomDomainSubscriptionSandboxConfig(instanceName, name, description),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMEnCustomDomainSubscriptionExists("ibm_en_subscription_custom_email.en_subscription_sandbox", conf),
					resource.TestCheckResourceAttrSet("ibm_en_subscription_custom_email.en_subscription_sandbox", "name"),
					resource.TestCheckResourceAttrSet("ibm_en_subscription_custom_email.en_subscription_sandbox", "description"),
					resource.TestCheckResourceAttrSet("ibm_en_subscription_custom_email.en_subscription_sandbox", "topic_id"),
					resource.TestCheckResourceAttrSet("ibm_en_subscription_custom_email.en_subscription_sandbox", "updated_at"),
					resource.TestCheckResourceAttrSet("ibm_en_subscription_custom_email.en_subscription_sandbox", "instance_guid"),
					resource.TestCheckResourceAttrSet("ibm_en_subscription_custom_email.en_subscription_sandbox", "destination_id"),
					resource.TestCheckResourceAttrSet("ibm_en_subscription_custom_email.en_subscription_sandbox", "destination_type"),
					resource.TestCheckResourceAttrSet("ibm_en_subscription_custom_email.en_subscription_sandbox", "subscription_id"),
					resource.TestCheckResourceAttrSet("ibm_en_subscription_custom_email.en_subscription_sandbox", "attributes.#"),
					resource.TestCheckResourceAttrSet("ibm_en_subscription_custom_email.en_subscription_sandbox", "attributes.0.reply_to_mail"),
					resource.TestCheckResourceAttrSet("ibm_en_subscription_custom_email.en_subscription_sandbox", "attributes.0.reply_to_name"),
					resource.TestCheckResourceAttrSet("ibm_en_subscription_custom_email.en_subscription_sandbox", "attributes.0.add_notification_payload"),
					resource.TestCheckResourceAttrSet("ibm_en_subscription_custom_email.en_subscription_sandbox", "attributes.0.invited"),
					// Note: from_name and from_email should NOT be set for sandbox destinations
				),
			},
			{
				Config: testAccCheckIBMEnCustomDomainSubscriptionSandboxConfig(instanceName, newName, newDescription),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_en_subscription_custom_email.en_subscription_sandbox", "name", newName),
					resource.TestCheckResourceAttr("ibm_en_subscription_custom_email.en_subscription_sandbox", "description", newDescription),
				),
			},
			{
				ResourceName:      "ibm_en_subscription_custom_email.en_subscription_sandbox",
				ImportState:       true,
				ImportStateVerify: false,
			},
		},
	})
}

func testAccCheckIBMEnCustomDomainSubscriptionProductionConfig(instanceName, name, description string) string {
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
	
	resource "ibm_en_destination_custom_email" "en_destination_production" {
		instance_guid = ibm_resource_instance.en_subscription_resource.guid
		name        = "Production Email Destination"
		type        = "smtp_custom"
		description = "Production custom email destination"
		is_sandbox  = false
		config {
			params {
				domain  = "production.example.com"
			}
		}
	}
	
	resource "ibm_en_subscription_custom_email" "en_subscription_production" {
		name             = "%s"
		description 	 = "%s"
		instance_guid    = ibm_resource_instance.en_subscription_resource.guid
		topic_id         = ibm_en_topic.en_topic_resource_2.topic_id
		destination_id   = ibm_en_destination_custom_email.en_destination_production.destination_id
		attributes {
			add_notification_payload = true
			reply_to_mail = "en@ibm.com"
			reply_to_name = "EYS ORG"
			from_name = "ABC ORG"
			from_email = "testuser@production.example.com"
			invited = ["testmail@gmail.com"]
		}
	}
	`, instanceName, name, description)
}

func testAccCheckIBMEnCustomDomainSubscriptionSandboxConfig(instanceName, name, description string) string {
	return fmt.Sprintf(`
	resource "ibm_resource_instance" "en_subscription_resource" {
		name     = "%s"
		location = "us-south"
		plan     = "standard"
		service  = "event-notifications"
	}
	
	resource "ibm_en_topic" "en_topic_resource_2" {
		instance_guid = ibm_resource_instance.en_subscription_resource.guid
		name        = "tf_topic_name_sandbox"
		description = "tf_topic_description_sandbox"
	}
	
	resource "ibm_en_destination_custom_email" "en_destination_sandbox" {
		instance_guid = ibm_resource_instance.en_subscription_resource.guid
		name        = "Sandbox Email Destination"
		type        = "smtp_custom"
		description = "Sandbox custom email destination"
		is_sandbox  = true
		config {
			params {
				domain  = "sandbox.example.com"
			}
		}
	}
	
	resource "ibm_en_subscription_custom_email" "en_subscription_sandbox" {
		name             = "%s"
		description 	 = "%s"
		instance_guid    = ibm_resource_instance.en_subscription_resource.guid
		topic_id         = ibm_en_topic.en_topic_resource_2.topic_id
		destination_id   = ibm_en_destination_custom_email.en_destination_sandbox.destination_id
		attributes {
			add_notification_payload = true
			reply_to_mail = "sandbox@ibm.com"
			reply_to_name = "Sandbox Team"
			invited = ["sandboxtest@gmail.com"]
			# Note: from_name and from_email are NOT included for sandbox destinations
		}
	}
	`, instanceName, name, description)
}

func testAccCheckIBMEnCustomDomainSubscriptionExists(n string, obj en.Subscription) resource.TestCheckFunc {

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

func testAccCheckIBMEnCustomDomainSubscriptionDestroy(s *terraform.State) error {
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
