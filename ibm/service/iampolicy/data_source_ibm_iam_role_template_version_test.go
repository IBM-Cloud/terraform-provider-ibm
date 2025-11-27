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

func TestAccIBMIAMRoleTemplateVersionDataSourceBasic(t *testing.T) {
	name := fmt.Sprintf("TerraformTemplateTest%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMRoleTemplateVersionDataSourceConfigBasic(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_iam_role_template_version.role_template", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_role_template_version.role_template", "account_id"),
				),
			},
		},
	})
}

func testAccCheckIBMRoleTemplateVersionDataSourceConfigBasic(name string) string {
	return fmt.Sprintf(`
	
	resource "ibm_iam_role_template" "role_template" {
		name = "%s"
		description = "Create Action Control templates through Terraform datasource"
        role {
			name = "TerraformDataSourceRoleTest"
			display_name = "TerraformDataSourceRoleDisplayNameTest"
			actions = ["am-test-service.test.create" ]
			service_name="am-test-service"
		}
        committed =  true
	}

	data "ibm_iam_role_template_version" "role_template" {
		role_template_id = ibm_iam_role_template.role_template.role_template_id
		version = ibm_iam_role_template.role_template.version
	}
	`, name)
}
