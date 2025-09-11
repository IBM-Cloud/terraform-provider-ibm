// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package eventnotification_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMEnPreDefinedTemplateDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMEnPreDefinedTemplateDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_en_pre_defined_template.en_pre_defined_template_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_en_pre_defined_template.en_pre_defined_template_instance", "instance_id"),
					resource.TestCheckResourceAttrSet("data.ibm_en_pre_defined_template.en_pre_defined_template_instance", "template_id"),
					resource.TestCheckResourceAttrSet("data.ibm_en_pre_defined_template.en_pre_defined_template_instance", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_en_pre_defined_template.en_pre_defined_template_instance", "description"),
					resource.TestCheckResourceAttrSet("data.ibm_en_pre_defined_template.en_pre_defined_template_instance", "source"),
					resource.TestCheckResourceAttrSet("data.ibm_en_pre_defined_template.en_pre_defined_template_instance", "type"),
					resource.TestCheckResourceAttrSet("data.ibm_en_pre_defined_template.en_pre_defined_template_instance", "updated_at"),
				),
			},
		},
	})
}

func testAccCheckIBMEnPreDefinedTemplateDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_en_pre_defined_template" "en_pre_defined_template_instance" {
			instance_guid = "instance_id"
			template_id = "id"
		}
	`)
}
