// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package iamaccessgroup_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMIAMAccessGroupTemplateVersionsDataSourceBasic(t *testing.T) {
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMIAMAccessGroupTemplateVersionDataSourceConfigBasic(name, agName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_iam_access_group_template_versions.template", "group_template_versions.#"),
				),
			},
		},
	})
}

func testAccCheckIBMIAMAccessGroupTemplateVersionDataSourceConfigBasic(name string, agName string) string {
	return fmt.Sprintf(`

	resource "ibm_iam_access_group_template" "template" {
		name = "%s"
		description = "Testing4"
		group {
			name = "%s"
		}
	}
		data "ibm_iam_access_group_template_versions" "template" {
			template_id = ibm_iam_access_group_template.template.template_id
		}
	`, name, agName)
}
