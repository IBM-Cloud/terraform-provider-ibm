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

func TestAccIBMIAMActionControlAssignmentDataSourceBasic(t *testing.T) {
	name := fmt.Sprintf("TerraformTemplateTest%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMActionControlAssignmentDataSourceConfigBasic(name, acc.TargetAccountId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_iam_action_control_assignment.action_control_assignment", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_action_control_assignment.action_control_assignment", "account_id"),
				),
			},
		},
	})
}

func testAccCheckIBMActionControlAssignmentDataSourceConfigBasic(name string, target string) string {
	return fmt.Sprintf(`
	
	resource "ibm_iam_action_control_template" "action_control_template" {
		name = "%s"
		action_control {
			actions = ["am-test-service.test.delete" ]
			service_name="am-test-service"
		}
		committed=true
}

resource "ibm_iam_action_control_assignment" "action_control_assignment" {
	target  ={
		type = "Account"
		id = "%s"
	}

	templates{
		id = ibm_iam_action_control_template.action_control_template.action_control_template_id 
		version = ibm_iam_action_control_template.action_control_template.version
	}
}

data "ibm_iam_action_control_assignment" "action_control_assignment" {
	assignment_id = ibm_iam_action_control_assignment.action_control_assignment.id
}`, name, target)
}
