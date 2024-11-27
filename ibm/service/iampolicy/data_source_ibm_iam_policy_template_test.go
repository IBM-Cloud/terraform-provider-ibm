// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package iampolicy_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMPolicyTemplateDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMPolicyTemplateDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_iam_policy_template.policy_template", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_policy_template.policy_template", "policy_templates.#"),
				),
			},
		},
	})
}

func testAccCheckIBMPolicyTemplateDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_iam_policy_template" "policy_template" {
			
		}`)
}
