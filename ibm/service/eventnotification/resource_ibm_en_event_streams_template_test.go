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

func TestAccIBMEnEventStreamsTemplateAllArgs(t *testing.T) {
	var params en.Template
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	instanceName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	description := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	newName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	newDescription := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMEnEmailTemplateDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMEnEventStreamsTemplateConfig(instanceName, name, description),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMEnSlackTemplateExists("ibm_en_event_streams_template.en_template_resource_1", params),
					resource.TestCheckResourceAttr("ibm_en_event_streams_template.en_template_resource_1", "name", name),
					resource.TestCheckResourceAttr("ibm_en_event_streams_template.en_template_resource_1", "type", "event_streams.notification"),
					resource.TestCheckResourceAttr("ibm_en_event_streams_template.en_template_resource_1", "description", description),
				),
			},
			{
				Config: testAccCheckIBMEnEventStreamsTemplateConfig(instanceName, newName, newDescription),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_en_event_streams_template.en_template_resource_1", "name", newName),
					resource.TestCheckResourceAttr("ibm_en_event_streams_template.en_template_resource_1", "type", "event_streams.notification"),
					resource.TestCheckResourceAttr("ibm_en_event_streams_template.en_template_resource_1", "description", newDescription),
				),
			},
			{
				ResourceName:      "ibm_en_event_streams_template.en_template_resource_1",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMEnEventStreamsTemplateConfig(instanceName, name, description string) string {
	return fmt.Sprintf(`
	resource "ibm_resource_instance" "en_template_resource" {
		name     = "%s"
		location = "us-south"
		plan     = "standard"
		service  = "event-notifications"
	}
	
	resource "ibm_en_event_streams_template" "en_template_resource_1" {
		instance_guid = ibm_resource_instance.en_template_resource.guid
		name        = "%s"
		type        = "event_streams.notification"
		description = "%s"
		params {
			body  = "eyJuYW1lIjoie3tkYXRhLm5hbWV9fSIifQ=="
		}
	}
	`, instanceName, name, description)
}

func testAccCheckIBMEnEventStreamsTemplateExists(n string, obj en.Template) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		enClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).EventNotificationsApiV1()
		if err != nil {
			return err
		}

		options := &en.GetTemplateOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		options.SetInstanceID(parts[0])
		options.SetID(parts[1])

		result, _, err := enClient.GetTemplate(options)
		if err != nil {
			return err
		}

		obj = *result
		return nil
	}
}

func testAccCheckIBMEnEventStreamsTemplateDestroy(s *terraform.State) error {
	enClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).EventNotificationsApiV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "en_template_resource_1" {
			continue
		}

		options := &en.GetTemplateOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		options.SetInstanceID(parts[0])
		options.SetID(parts[1])

		// Try to find the key
		_, response, err := enClient.GetTemplate(options)

		if err == nil {
			return fmt.Errorf("en_template_resource still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("[ERROR] Error checking for en_template_resource_1 (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
