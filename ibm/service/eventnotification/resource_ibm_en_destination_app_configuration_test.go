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

func TestAccIBMEnAppConfigurationDestinationAllArgs(t *testing.T) {
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
				Config: testAccCheckIBMEnAppConfigurationDestinationConfig(instanceName, name, description),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMEnAppConfigurationDestinationExists("ibm_en_destination_app_configuration.en_destination_resource_1", config),
					resource.TestCheckResourceAttr("ibm_en_destination_app_configuration.en_destination_resource_1", "name", name),
					resource.TestCheckResourceAttr("ibm_en_destination_app_configuration.en_destination_resource_1", "type", "app_configuration"),
					resource.TestCheckResourceAttr("ibm_en_destination_app_configuration.en_destination_resource_1", "collect_failed_events", "false"),
					resource.TestCheckResourceAttr("ibm_en_destination_app_configuration.en_destination_resource_1", "description", description),
				),
			},
			{
				Config: testAccCheckIBMEnAppConfigurationDestinationConfig(instanceName, newName, newDescription),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_en_destination_app_configuration.en_destination_resource_1", "name", newName),
					resource.TestCheckResourceAttr("ibm_en_destination_app_configuration.en_destination_resource_1", "type", "app_configuration"),
					resource.TestCheckResourceAttr("ibm_en_destination_app_configuration.en_destination_resource_1", "collect_failed_events", "false"),
					resource.TestCheckResourceAttr("ibm_en_destination_app_configuration.en_destination_resource_1", "description", newDescription),
				),
			},
			{
				ResourceName:      "ibm_en_destination_app_configuration.en_destination_resource_1",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMEnAppConfigurationDestinationConfig(instanceName, name, description string) string {
	return fmt.Sprintf(`
	resource "ibm_resource_instance" "en_destination_resource" {
		name     = "%s"
		location = "us-south"
		plan     = "standard"
		service  = "event-notifications"
	}
	
	resource "ibm_en_destination_app_configuration" "en_destination_resource_1" {
		instance_guid = ibm_resource_instance.en_destination_resource.guid
		name        = "%s"
		type        = "app_configuration"
		description = "%s"
		config {
			params {
			    type = "features"
				crn = "crn:v1:staging:public:apprapp:us-south:a/9f007405a9fe4a5d9345fa8c131610c8:3a86a8e4-fe8b-4e43-9727-2f2cf987f1c8::"
				feature_id = "cross"
				environment_id = "dev"
			}
		}
	}
	`, instanceName, name, description)
}

func testAccCheckIBMEnAppConfigurationDestinationExists(n string, obj en.Destination) resource.TestCheckFunc {

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

func testAccCheckIBMEnAppConfigurationDestinationDestroy(s *terraform.State) error {
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
