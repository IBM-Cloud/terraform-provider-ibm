// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package eventnotification_test

import (
	"fmt"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/conns"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/flex"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	en "github.com/IBM/event-notifications-go-admin-sdk/eventnotificationsv1"
)

func TestAccIBMEnSourceAllArgs(t *testing.T) {
	var config en.Source
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	instanceName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	description := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	newName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	newDescription := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	enabled := true
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMEnSourceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMEnSourceConfig(instanceName, name, description, enabled),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMEnSourceExists("ibm_en_source.en_source_resource_1", config),
					resource.TestCheckResourceAttr("ibm_en_source.en_source_resource_1", "name", name),
					resource.TestCheckResourceAttr("ibm_en_source.en_source_resource_1", "enabled", "enabled"),
					resource.TestCheckResourceAttr("ibm_en_source.en_source_resource_1", "description", description),
				),
			},
			{
				Config: testAccCheckIBMEnSourceConfig(instanceName, newName, newDescription, enabled),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_en_source.en_source_resource_1", "name", newName),
					resource.TestCheckResourceAttr("ibm_en_source.en_source_resource_1", "enabled", "enabled"),
					resource.TestCheckResourceAttr("ibm_en_source.en_source_resource_1", "description", newDescription),
				),
			},
			{
				ResourceName:      "ibm_en_source.en_source_resource_1",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMEnSourceConfig(instanceName, name, description string, enabled bool) string {
	return fmt.Sprintf(`
	resource "ibm_resource_instance" "en_source_resource" {
		name     = "%s"
		location = "us-south"
		plan     = "standard"
		service  = "event-notifications"
	}
	
	resource "ibm_en_source" "en_source_resource_1" {
		instance_guid = ibm_resource_instance.en_destination_resource.guid
		name        = "%s"
		description = "%s"
		enabled = %t
	}
	`, instanceName, name, description, enabled)
}

func testAccCheckIBMEnSourceExists(n string, obj en.Source) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		enClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).EventNotificationsApiV1()
		if err != nil {
			return err
		}

		options := &en.GetSourceOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		options.SetInstanceID(parts[0])
		options.SetID(parts[1])

		result, _, err := enClient.GetSource(options)
		if err != nil {
			return err
		}

		obj = *result
		return nil
	}
}

func testAccCheckIBMEnSourceDestroy(s *terraform.State) error {
	enClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).EventNotificationsApiV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "en_source_resource_1" {
			continue
		}

		options := &en.GetSourceOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		options.SetInstanceID(parts[0])
		options.SetID(parts[1])

		// Try to find the key
		_, response, err := enClient.GetSource(options)

		if err == nil {
			return fmt.Errorf("en_source still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("[ERROR] Error checking for en_source (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
