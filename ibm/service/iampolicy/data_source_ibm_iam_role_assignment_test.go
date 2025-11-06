// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package iampolicy_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMIAMRoleAssignmentDataSourceBasic(t *testing.T) {
	name := fmt.Sprintf("TerraformTemplateTest%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMRoleAssignmentDataSourceConfigBasic(name, acc.TargetAccountId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_iam_role_assignment.role_assignment", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_role_assignment.role_assignment", "account_id"),
				),
			},
		},
	})
}

func testAccCheckIBMRoleAssignmentDataSourceConfigBasic(name string, target string) string {
	return fmt.Sprintf(`
	
	resource "ibm_iam_role_template" "role_template" {
		name = "%s"
		role {
			name = "DataSourceTerraformAssignments"
			display_name = "TerraformAssignments"
			actions = ["am-test-service.test.delete", "iam.policy.delete" ]
			service_name="am-test-service"
		}
		committed=true
}

resource "ibm_iam_role_assignment" "role_assignment" {
	target  ={
		type = "Account"
		id = "%s"
	}

	templates{
		id = ibm_iam_role_template.role_template.role_template_id 
		version = ibm_iam_role_template.role_template.version
	}
}

data "ibm_iam_role_assignment" "role_assignment" {
	assignment_id = ibm_iam_role_assignment.role_assignment.id
}`, name, target)
}
