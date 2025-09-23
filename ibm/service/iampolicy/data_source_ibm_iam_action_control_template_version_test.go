// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package iampolicy_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMIAMActionControlTemplateVersionDataSourceBasic(t *testing.T) {
	name := fmt.Sprintf("TerraformTemplateTest%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMActionControlTemplateVersionDataSourceConfigBasic(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_iam_action_control_template_version.action_control_template", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_action_control_template_version.action_control_template", "account_id"),
				),
			},
		},
	})
}

func testAccCheckIBMActionControlTemplateVersionDataSourceConfigBasic(name string) string {
	return fmt.Sprintf(`
	
	resource "ibm_iam_action_control_template" "action_control_template" {
		name = "%s"
		description = "Create Action Control templates through Terraform datasource"
        action_control {
		actions = ["am-test-service.test.create" ]
		service_name="am-test-service"
	} 
        committed =  true
	}

	data "ibm_iam_action_control_template_version" "action_control_template" {
		action_control_template_id = ibm_iam_action_control_template.action_control_template.action_control_template_id
		version = ibm_iam_action_control_template.action_control_template.version
	}
	`, name)
}
