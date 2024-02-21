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

func TestAccIBMEnEmailTemplateAllArgs(t *testing.T) {
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
				Config: testAccCheckIBMEnEmailTemplateConfig(instanceName, name, description),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMEnEmailTemplateExists("ibm_en_email_template.en_template_resource_1", params),
					resource.TestCheckResourceAttr("ibm_en_email_template.en_template_resource_1", "name", name),
					resource.TestCheckResourceAttr("ibm_en_email_template.en_template_resource_1", "type", "smtp_custom.notification"),
					resource.TestCheckResourceAttr("ibm_en_email_template.en_template_resource_1", "description", description),
				),
			},
			{
				Config: testAccCheckIBMEnEmailTemplateConfig(instanceName, newName, newDescription),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_en_email_template.en_template_resource_1", "name", newName),
					resource.TestCheckResourceAttr("ibm_en_email_template.en_template_resource_1", "type", "smtp_custom"),
					resource.TestCheckResourceAttr("ibm_en_email_template.en_template_resource_1", "description", newDescription),
				),
			},
			{
				ResourceName:      "ibm_en_email_template.en_template_resource_1",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMEnEmailTemplateConfig(instanceName, name, description string) string {
	return fmt.Sprintf(`
	resource "ibm_resource_instance" "en_template_resource" {
		name     = "%s"
		location = "us-south"
		plan     = "standard"
		service  = "event-notifications"
	}
	
	resource "ibm_en_email_template" "en_template_resource_1" {
		instance_guid = ibm_resource_instance.en_template_resource.guid
		name        = "%s"
		type        = "smtp_custom.notification"
		description = "%s"
		params {
			body  = "<!DOCTYPE html><html><head><title>Go To-Do list</title></head><body><p>To-Do list for user: {{ Data.issuer.p }}</p><table><tr><td>Task</td><td>Done</td></tr>{{#each Email}}<tr><td>{{ this }}</td></tr>{{/each}}</table></body></html>"
			subject = "HI This is the template test for the invitation"
		}
	}
	`, instanceName, name, description)
}

func testAccCheckIBMEnEmailTemplateExists(n string, obj en.Template) resource.TestCheckFunc {

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

func testAccCheckIBMEnEmailTemplateDestroy(s *terraform.State) error {
	enClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).EventNotificationsApiV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "en_destination_resource_1" {
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
			return fmt.Errorf("en_destination still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("[ERROR] Error checking for en_destination (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
