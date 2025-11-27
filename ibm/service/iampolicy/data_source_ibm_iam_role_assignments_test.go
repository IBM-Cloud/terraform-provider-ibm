// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package iampolicy_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMIAMRoleAssignmentsDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMRoleAssignmentsDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_iam_role_assignments.role_assignment", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_iam_role_assignments.role_assignment", "assignments.#"),
				),
			},
		},
	})
}

func testAccCheckIBMRoleAssignmentsDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_iam_role_assignments" "role_assignment" {
		}`)
}
