// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package eventnotification_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMEnTemplatesDataSourceBasic(t *testing.T) {
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	instanceName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	description := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMEnTemplatesDatasourceConfig(instanceName, name, description),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_en_email_templates.data_email_template_1", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_en_email_templates.data_email_template_1", "instance_guid"),
					resource.TestCheckResourceAttrSet("data.ibm_en_email_templates.data_email_template_1", "total_count"),
					resource.TestCheckResourceAttrSet("data.ibm_en_email_templates.data_email_template_1", "templates.#"),
					resource.TestCheckResourceAttrSet("data.ibm_en_email_templates.data_email_template_1", "templates.0.name"),
					resource.TestCheckResourceAttrSet("data.ibm_en_email_templates.data_email_template_1", "templates.0.updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_en_email_templates.data_email_template_1", "templates.0.type"),
					resource.TestCheckResourceAttrSet("data.ibm_en_email_templates.data_email_template_1", "templates.0.id"),
					resource.TestCheckResourceAttrSet("data.ibm_en_email_templates.data_email_template_1", "templates.0.description"),
				),
			},
		},
	})
}

func testAccCheckIBMEnTemplatesDatasourceConfig(instanceName, name, description string) string {
	return fmt.Sprintf(`
	resource "ibm_resource_instance" "en_template_datasource" {
		name     = "%s"
		location = "us-south"
		plan     = "standard"
		service  = "event-notifications"
	}
	
	resource "ibm_en_email_template" "en_template_datasource_1" {
		instance_guid = ibm_resource_instance.en_template_datasource.guid
		name        = "%s"
		type        = "smtp_custom.notification"
		description = "%s"
		params {
			body  = "<!DOCTYPE html><html><head><title>Go To-Do list</title></head><body><p>To-Do list for user: {{ Data.issuer.p }}</p><table><tr><td>Task</td><td>Done</td></tr>{{#each Email}}<tr><td>{{ this }}</td></tr>{{/each}}</table></body></html>"
			subject = "HI This is the template test for the invitation"
		}
	}

	data "ibm_en_destinations" "data_email_template_1" {
		instance_guid = ibm_resource_instance.en_template_datasource_1.guid
	}
	`, instanceName, name, description)
}
