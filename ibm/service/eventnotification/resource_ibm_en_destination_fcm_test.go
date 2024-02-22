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

func TestAccIBMEnFCMDestinationAllArgs(t *testing.T) {
	var config en.Destination
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	instanceName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	description := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	newName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	newDescription := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMEnFCMDestinationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMEnFCMDestinationConfig(instanceName, name, description),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMEnFCMDestinationExists("ibm_en_destination_android.en_destination_resource_fcm", config),
					resource.TestCheckResourceAttr("ibm_en_destination_android.en_destination_resource_fcm", "name", name),
					resource.TestCheckResourceAttr("ibm_en_destination_android.en_destination_resource_fcm", "type", "push_android"),
					resource.TestCheckResourceAttr("ibm_en_destination_android.en_destination_resource_fcm", "collect_failed_events", "false"),
					resource.TestCheckResourceAttr("ibm_en_destination_android.en_destination_resource_fcm", "description", description),
				),
			},
			{
				Config: testAccCheckIBMEnFCMDestinationConfig(instanceName, newName, newDescription),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_en_destination_android.en_destination_resource_fcm", "name", newName),
					resource.TestCheckResourceAttr("ibm_en_destination_android.en_destination_resource_fcm", "type", "push_android"),
					resource.TestCheckResourceAttr("ibm_en_destination_android.en_destination_resource_fcm", "collect_failed_events", "false"),
					resource.TestCheckResourceAttr("ibm_en_destination_android.en_destination_resource_fcm", "description", newDescription),
				),
			},
			{
				ResourceName:      "ibm_en_destination_android.en_destination_resource_1",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMEnFCMDestinationConfig(instanceName, name, description string) string {
	return fmt.Sprintf(`
	resource "ibm_resource_instance" "en_destination_resource" {
		name     = "%s"
		location = "us-south"
		plan     = "standard"
		service  = "event-notifications"
	}
	
	resource "ibm_en_destination_android" "en_destination_resource_fcm" {
		instance_guid = ibm_resource_instance.en_destination_resource.guid
		name        = "%s"
		type        = "push_android"
		description = "%s"
		config {
			params {
				project_id = "FCM_sender_id_value"
				private_key  = "FCM_Server_key_value"
				client_email = "testuser123@gmail.com"
				pre_prod = false
			}
		}
	}
	`, instanceName, name, description)
}

func testAccCheckIBMEnFCMDestinationExists(n string, obj en.Destination) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		enClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).EventNotificationsApiV1()
		if err != nil {
			return err
		}

		options := &en.GetDestinationOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		options.SetInstanceID(parts[0])
		options.SetID(parts[1])

		result, _, err := enClient.GetDestination(options)
		if err != nil {
			return err
		}

		obj = *result
		return nil
	}
}

func testAccCheckIBMEnFCMDestinationDestroy(s *terraform.State) error {
	enClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).EventNotificationsApiV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "en_destination_resource_fcm" {
			continue
		}

		options := &en.GetDestinationOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		options.SetInstanceID(parts[0])
		options.SetID(parts[1])

		// Try to find the key
		_, response, err := enClient.GetDestination(options)

		if err == nil {
			return fmt.Errorf("en_destination still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("[ERROR] Error checking for en_destination (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
