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

func TestAccIBMEnCustomEmailDestinationAllArgs(t *testing.T) {
	var config en.Destination
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	instanceName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	description := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	newName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	newDescription := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMEnServiceNowDestinationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMEnCustomEmailDestinationConfig(instanceName, name, description),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMEnCustomEmailDestinationExists("ibm_en_destination_custom_email.en_destination_resource_1", config),
					resource.TestCheckResourceAttr("ibm_en_destination_custom_email.en_destination_resource_1", "name", name),
					resource.TestCheckResourceAttr("ibm_en_destination_custom_email.en_destination_resource_1", "type", "smtp_custom"),
					resource.TestCheckResourceAttr("ibm_en_destination_custom_email.en_destination_resource_1", "collect_failed_events", "false"),
					resource.TestCheckResourceAttr("ibm_en_destination_custom_email.en_destination_resource_1", "description", description),
					resource.TestCheckResourceAttr("ibm_en_destination_custom_email.en_destination_resource_1", "is_sandbox", "false"),
				),
			},
			{
				Config: testAccCheckIBMEnCustomEmailDestinationConfig(instanceName, newName, newDescription),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_en_destination_custom_email.en_destination_resource_1", "name", newName),
					resource.TestCheckResourceAttr("ibm_en_destination_custom_email.en_destination_resource_1", "type", "smtp_custom"),
					resource.TestCheckResourceAttr("ibm_en_destination_custom_email.en_destination_resource_1", "collect_failed_events", "false"),
					resource.TestCheckResourceAttr("ibm_en_destination_custom_email.en_destination_resource_1", "description", newDescription),
					resource.TestCheckResourceAttr("ibm_en_destination_custom_email.en_destination_resource_1", "is_sandbox", "false"),
				),
			},
			{
				ResourceName:      "ibm_en_destination_custom_email.en_destination_resource_1",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccIBMEnCustomEmailDestinationSandbox(t *testing.T) {
	var config en.Destination
	name := fmt.Sprintf("tf_sandbox_%d", acctest.RandIntRange(10, 100))
	instanceName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	description := fmt.Sprintf("tf_sandbox_description_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMEnCustomEmailDestinationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMEnCustomEmailDestinationSandboxConfig(instanceName, name, description),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMEnCustomEmailDestinationExists("ibm_en_destination_custom_email.en_destination_sandbox", config),
					resource.TestCheckResourceAttr("ibm_en_destination_custom_email.en_destination_sandbox", "name", name),
					resource.TestCheckResourceAttr("ibm_en_destination_custom_email.en_destination_sandbox", "type", "smtp_custom_sandbox"),
					resource.TestCheckResourceAttr("ibm_en_destination_custom_email.en_destination_sandbox", "description", description),
					resource.TestCheckResourceAttr("ibm_en_destination_custom_email.en_destination_sandbox", "is_sandbox", "true"),
				),
			},
			{
				ResourceName:      "ibm_en_destination_custom_email.en_destination_sandbox",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccIBMEnCustomEmailDestinationSandboxUpgrade(t *testing.T) {
	var config en.Destination
	name := fmt.Sprintf("tf_sandbox_upgrade_%d", acctest.RandIntRange(10, 100))
	instanceName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	description := fmt.Sprintf("tf_sandbox_upgrade_description_%d", acctest.RandIntRange(10, 100))
	domain := fmt.Sprintf("test%d.example.com", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMEnCustomEmailDestinationDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMEnCustomEmailDestinationSandboxUpgradeConfig(instanceName, name, description, domain, true),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMEnCustomEmailDestinationExists("ibm_en_destination_custom_email.en_destination_upgrade", config),
					resource.TestCheckResourceAttr("ibm_en_destination_custom_email.en_destination_upgrade", "name", name),
					resource.TestCheckResourceAttr("ibm_en_destination_custom_email.en_destination_upgrade", "type", "smtp_custom_sandbox"),
					resource.TestCheckResourceAttr("ibm_en_destination_custom_email.en_destination_upgrade", "description", description),
					resource.TestCheckResourceAttr("ibm_en_destination_custom_email.en_destination_upgrade", "is_sandbox", "true"),
				),
			},
			{
				Config: testAccCheckIBMEnCustomEmailDestinationSandboxUpgradeConfig(instanceName, name, description, domain, false),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMEnCustomEmailDestinationExists("ibm_en_destination_custom_email.en_destination_upgrade", config),
					resource.TestCheckResourceAttr("ibm_en_destination_custom_email.en_destination_upgrade", "name", name),
					resource.TestCheckResourceAttr("ibm_en_destination_custom_email.en_destination_upgrade", "type", "smtp_custom"),
					resource.TestCheckResourceAttr("ibm_en_destination_custom_email.en_destination_upgrade", "description", description),
					resource.TestCheckResourceAttr("ibm_en_destination_custom_email.en_destination_upgrade", "is_sandbox", "false"),
				),
			},
		},
	})
}

func testAccCheckIBMEnCustomEmailDestinationConfig(instanceName, name, description string) string {
	return fmt.Sprintf(`
	resource "ibm_resource_instance" "en_destination_resource" {
		name     = "%s"
		location = "us-south"
		plan     = "standard"
		service  = "event-notifications"
	}
	
	resource "ibm_en_destination_custom_email" "en_destination_resource_1" {
		instance_guid = ibm_resource_instance.en_destination_resource.guid
		name        = "%s"
		type        = "smtp_custom"
		description = "%s"
		is_sandbox  = false
		config {
			params {
				domain  = "mailx.com"
			}
		}
	}
	`, instanceName, name, description)
}

func testAccCheckIBMEnCustomEmailDestinationSandboxConfig(instanceName, name, description string) string {
	return fmt.Sprintf(`
	resource "ibm_resource_instance" "en_destination_resource" {
		name     = "%s"
		location = "us-south"
		plan     = "standard"
		service  = "event-notifications"
	}
	
	resource "ibm_en_destination_custom_email" "en_destination_sandbox" {
		instance_guid = ibm_resource_instance.en_destination_resource.guid
		name        = "%s"
		type        = "smtp_custom"
		description = "%s"
		is_sandbox  = true
	`, instanceName, name, description)
}

func testAccCheckIBMEnCustomEmailDestinationProductionConfig(instanceName, name, description, domain string) string {
	return fmt.Sprintf(`
	resource "ibm_resource_instance" "en_destination_resource" {
		name     = "%s"
		location = "us-south"
		plan     = "standard"
		service  = "event-notifications"
	}
	
	resource "ibm_en_destination_custom_email" "en_destination_production" {
		instance_guid = ibm_resource_instance.en_destination_resource.guid
		name        = "%s"
		type        = "smtp_custom"
		description = "%s"
		is_sandbox  = false
		config {
			params {
				domain  = "%s"
			}
		}
	}
	`, instanceName, name, description, domain)
}

func testAccCheckIBMEnCustomEmailDestinationSandboxUpgradeConfig(instanceName, name, description, domain string, isSandbox bool) string {
	return fmt.Sprintf(`
	resource "ibm_resource_instance" "en_destination_resource" {
		name     = "%s"
		location = "us-south"
		plan     = "standard"
		service  = "event-notifications"
	}
	
	resource "ibm_en_destination_custom_email" "en_destination_upgrade" {
		instance_guid = ibm_resource_instance.en_destination_resource.guid
		name        = "%s"
		type        = "smtp_custom"
		description = "%s"
		is_sandbox  = %t
		config {
			params {
				domain  = "%s"
			}
		}
	}
	`, instanceName, name, description, isSandbox, domain)
}

func testAccCheckIBMEnCustomEmailDestinationExists(n string, obj en.Destination) resource.TestCheckFunc {

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

func testAccCheckIBMEnCustomEmailDestinationDestroy(s *terraform.State) error {
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
