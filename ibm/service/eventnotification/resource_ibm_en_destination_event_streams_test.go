// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package eventnotification_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"

	en "github.com/IBM/event-notifications-go-admin-sdk/eventnotificationsv1"
)

func TestAccIBMEnEventStreamsDestinationAllArgs(t *testing.T) {
	var config en.Destination
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	instanceName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	description := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	newName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	newDescription := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMEnPagerDutyDestinationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMEnEventStreamsDestinationConfig(instanceName, name, description),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMEnEventStreamsDestinationExists("ibm_en_destination_event_streams.en_destination_resource_1", config),
					resource.TestCheckResourceAttr("ibm_en_destination_event_streams.en_destination_resource_1", "name", name),
					resource.TestCheckResourceAttr("ibm_en_destination_event_streams.en_destination_resource_1", "type", "event_streams"),
					resource.TestCheckResourceAttr("ibm_en_destination_event_streams.en_destination_resource_1", "collect_failed_events", "false"),
					resource.TestCheckResourceAttr("ibm_en_destination_event_streams.en_destination_resource_1", "description", description),
				),
			},
			{
				Config: testAccCheckIBMEnEventStreamsDestinationConfig(instanceName, newName, newDescription),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_en_destination_event_streams.en_destination_resource_1", "name", newName),
					resource.TestCheckResourceAttr("ibm_en_destination_event_streams.en_destination_resource_1", "type", "event_streams"),
					resource.TestCheckResourceAttr("ibm_en_destination_event_streams.en_destination_resource_1", "collect_failed_events", "false"),
					resource.TestCheckResourceAttr("ibm_en_destination_event_streams.en_destination_resource_1", "description", newDescription),
				),
			},
			{
				ResourceName:      "ibm_en_destination_event_streams.en_destination_resource_1",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMEnEventStreamsDestinationConfig(instanceName, name, description string) string {
	return fmt.Sprintf(`
	resource "ibm_resource_instance" "en_destination_resource" {
		name     = "%s"
		location = "us-south"
		plan     = "standard"
		service  = "event-notifications"
	}
	
	resource "ibm_en_destination_event_streams" "en_destination_resource_1" {
		instance_guid = ibm_resource_instance.en_destination_resource.guid
		name        = "%s"
		type        = "event_streams"
		description = "%s"
		config {
			params {
				crn = "crn:v1:staging:public:messagehub:us-south:a/9f007405a9fe4a5d9345fa8c131610c8:a292db6e-af78-4c0b-b3db-7d6794b40aeb::"
				endpoint = "https://n6627w6t7y62cfgd.svc09.us-south.eventstreams.test.cloud.ibm.com"
				topic = "test_demo"
			}
		}
	}
	`, instanceName, name, description)
}

func testAccCheckIBMEnEventStreamsDestinationExists(n string, obj en.Destination) resource.TestCheckFunc {

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

func testAccCheckIBMEnEventStreamsDestinationDestroy(s *terraform.State) error {
	enClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).EventNotificationsApiV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "en_destination_resource_1" {
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
			return fmt.Errorf("en_destination_resource_1 still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("[ERROR] Error checking for en_destination_resource_1 (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
