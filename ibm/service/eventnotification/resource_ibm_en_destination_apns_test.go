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

func TestAccIBMEnAPNSDestinationAllArgs(t *testing.T) {
	var config en.Destination
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	instanceName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	description := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	newName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	newDescription := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMEnAPNSDestinationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMEnAPNSDestinationConfig(instanceName, name, description),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMEnAPNSDestinationExists("ibm_en_apns_destination.en_destination_resource_apns", config),
					resource.TestCheckResourceAttr("ibm_en_destination_ios.en_destination_resource_apns", "name", name),
					resource.TestCheckResourceAttr("ibm_en_destination_ios.en_destination_resource_apns", "type", "push_ios"),
					resource.TestCheckResourceAttr("ibm_en_destination_ios.en_destination_resource_apns", "description", description),
				),
			},
			{
				Config: testAccCheckIBMEnAPNSDestinationConfig(instanceName, newName, newDescription),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_en_destination_ios.en_destination_resource_apns", "name", newName),
					resource.TestCheckResourceAttr("ibm_en_destination_ios.en_destination_resource_apns", "type", "push_ios"),
					resource.TestCheckResourceAttr("ibm_en_destination_ios.en_destination_resource_apns", "description", newDescription),
				),
			},
			{
				ResourceName:      "ibm_en_destination_ios.en_destination_resource_apns",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMEnAPNSDestinationConfig(instanceName, name, description string) string {
	return fmt.Sprintf(`
	resource "ibm_resource_instance" "en_destination_resource" {
		name     = "%s"
		location = "us-south"
		plan     = "standard"
		service  = "event-notifications"
	}
	
	resource "ibm_en_destination_ios" "en_destination_resource_apns" {
		instance_guid = ibm_resource_instance.en_destination_resource.guid
		name        = "%s"
		type        = "push_ios"
		certificate_content_type = "p12"
        certificate = "${path.module}/cert.p12"
		description = "%s"
		config {
			params {
				cert_type = "p12"
                is_sandbox = true
                password = "certpassword"
			}
		}
	}
	`, instanceName, name, description)
}

func testAccCheckIBMEnAPNSDestinationExists(n string, obj en.Destination) resource.TestCheckFunc {

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

func testAccCheckIBMEnAPNSDestinationDestroy(s *terraform.State) error {
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
			return fmt.Errorf("en_destination still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("[ERROR] Error checking for en_destination (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
