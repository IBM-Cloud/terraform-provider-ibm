// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package schematics_test

import (
	"fmt"
	"testing"

	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMSchematicsStateDataSourceBasic(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMSchematicsStateDataSourceConfigBasic(acc.WorkspaceID, acc.TemplateID),
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
	 `, acc.WorkspaceID, templateID)
}
