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

func TestAccIBMPolicyAssignmentDataSourceBasic(t *testing.T) {
	name := fmt.Sprintf("TerraformTemplateTest%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPolicyAssignmentDataSourceConfigBasic(name, acc.TargetAccountId),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_iam_policy_assignment.policy_assignment", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_policy_assignment.policy_assignment", "account_id"),
				),
			},
		},
	})
}

func testAccCheckIBMPolicyAssignmentDataSourceConfigBasic(name string, target string) string {
	return fmt.Sprintf(`
	
	resource "ibm_iam_policy_template" "policy_s2s_template" {
		name = "%s"
		policy {
			type = "authorization"
			description = "description"
			resource {
				attributes {
					key = "serviceName"
					operator = "stringEquals"
					value = "kms"
				}
			}
			subject {
				attributes {
					key = "serviceName"
					operator = "stringEquals"
					value = "compliance"
				}
			}
			roles = ["Reader"]
		}
		committed=true
}

resource "ibm_iam_policy_assignment" "policy_assignment" {
	version ="1.0"
	target  ={
		type = "Account"
		id = "%s"
	}

	templates{
		id = ibm_iam_policy_template.policy_s2s_template.template_id 
		version = ibm_iam_policy_template.policy_s2s_template.version
	}
}

data "ibm_iam_policy_assignment" "policy_assignment" {
	assignment_id = ibm_iam_policy_assignment.policy_assignment.id
}`, name, target)
}
