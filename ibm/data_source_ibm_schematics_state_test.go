// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMSchematicsStateDataSourceBasic(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMSchematicsStateDataSourceConfigBasic(workspaceID, templateID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_schematics_state.schematics_state", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_state.schematics_state", "state_store"),
				),
			},
		},
	})
}

func testAccCheckIBMSchematicsStateDataSourceConfigBasic(workspaceID string, templateID string) string {
	return fmt.Sprintf(`
		 data "ibm_schematics_state" "schematics_state" {
			workspace_id = "%s"
			 template_id = "%s"
		 }
	 `, workspaceID, templateID)
}
