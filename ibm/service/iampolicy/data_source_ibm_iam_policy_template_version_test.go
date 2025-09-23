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

func TestAccIBMPolicyTemplateVersionDataSourceBasic(t *testing.T) {
	name := fmt.Sprintf("TerraformTemplateTest%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPolicyTemplateVersionDataSourceConfigBasic(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_iam_policy_template_version.policy_template", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_policy_template_version.policy_template", "account_id"),
				),
			},
		},
	})
}

func testAccCheckIBMPolicyTemplateVersionDataSourceConfigBasic(name string) string {
	return fmt.Sprintf(`
	
	resource "ibm_iam_policy_template" "policy_template" {
		name = "%s"
		policy {
			type = "access"
			description = "description"
			resource {
				attributes {
					key = "serviceName"
					operator = "stringEquals"
					value = "is"
				}
			}
			roles = ["Operator"]
		}
	}

	data "ibm_iam_policy_template_version" "policy_template" {
		policy_template_id = ibm_iam_policy_template.policy_template.template_id
		version = ibm_iam_policy_template.policy_template.version
	}
	`, name)
}
