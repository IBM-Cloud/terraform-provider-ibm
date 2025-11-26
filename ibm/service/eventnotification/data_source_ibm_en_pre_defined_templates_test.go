// Copyright IBM Corp. 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package eventnotification_test

import (
	"fmt"
	"testing"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccIBMEnPreDefinedTemplatesDataSourceBasic(t *testing.T) {
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMEnPreDefinedTemplatesDatasourceConfig(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_en_pre_defined_templates.data_pre_defined_template_1", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_en_pre_defined_templates.data_pre_defined_template_1", "instance_guid"),
					resource.TestCheckResourceAttrSet("data.ibm_en_pre_defined_templates.data_pre_defined_template_1", "total_count"),
					resource.TestCheckResourceAttrSet("data.ibm_en_pre_defined_templates.data_pre_defined_template_1", "source"),
					resource.TestCheckResourceAttrSet("data.ibm_en_pre_defined_templates.data_pre_defined_template_1", "type"),
					resource.TestCheckResourceAttrSet("data.ibm_en_pre_defined_templates.data_pre_defined_template_1", "templates.#"),
					resource.TestCheckResourceAttrSet("data.ibm_en_pre_defined_templates.data_pre_defined_template_1", "templates.0.name"),
					resource.TestCheckResourceAttrSet("data.ibm_en_pre_defined_templates.data_pre_defined_template_1", "templates.0.source"),
					resource.TestCheckResourceAttrSet("data.ibm_en_pre_defined_templates.data_pre_defined_template_1", "templates.0.updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_en_pre_defined_templates.data_pre_defined_template_1", "templates.0.type"),
					resource.TestCheckResourceAttrSet("data.ibm_en_pre_defined_templates.data_pre_defined_template_1", "templates.0.id"),
					resource.TestCheckResourceAttrSet("data.ibm_en_pre_defined_templates.data_pre_defined_template_1", "templates.0.description"),
				),
			},
		},
	})
}

func testAccCheckIBMEnPreDefinedTemplatesDatasourceConfig(name string) string {
	return fmt.Sprintf(`
	resource "ibm_resource_instance" "en_template_datasource" {
		name     = "%s"
		location = "us-south"
		plan     = "standard"
		service  = "event-notifications"
	}

	data "ibm_en_destinations" "data_pre_defined_template_1" {
		instance_guid = ibm_resource_instance.en_template_datasource_1.guid
		type        = "slack.notification"
		source = "logs"
	}
	`, name)
}
