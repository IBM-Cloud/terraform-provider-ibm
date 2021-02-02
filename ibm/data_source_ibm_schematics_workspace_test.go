// Copyright IBM Corp. 2017, 2021 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccIBMSchematicsWorkspaceDataSource_basic(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckIBMSchematicsWorkspaceDataSourceConfig(workspaceID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_schematics_workspace.test", "workspace_id", workspaceID),
				),
			},
		},
	})
}

func testAccCheckIBMSchematicsWorkspaceDataSourceConfig(workspaceID string) string {
	return fmt.Sprintf(`
	data "ibm_schematics_workspace" "test" {
		workspace_id = "%s"
	}
	  
	output "WorkSpaceValues" {
		value = data.ibm_schematics_workspace.test.template_id.0
	}
`, workspaceID)
}
