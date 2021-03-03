// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMSchematicsStateDataSource_basic(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMSchematicsStateDataSourceConfig(workspaceID, templateID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_schematics_state.test", "workspace_id", workspaceID),
				),
			},
		},
	})
}

func testAccCheckIBMSchematicsStateDataSourceConfig(WorkspaceID, templateID string) string {
	return fmt.Sprintf(`
	data "ibm_schematics_workspace" "test" {
		workspace_id = "%s"
	}
	data "ibm_schematics_state" "test" {
		workspace_id = data.ibm_schematics_workspace.test.workspace_id
		template_id = data.ibm_schematics_workspace.test.template_id.0
	}
	  
	output "state_store_values" {
		value = data.ibm_schematics_state.test.state_store
	}
`, WorkspaceID)
}
