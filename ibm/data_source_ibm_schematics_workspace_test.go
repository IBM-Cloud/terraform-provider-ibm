/* IBM Confidential
*  Object Code Only Source Materials
*  5747-SM3
*  (c) Copyright IBM Corp. 2017,2021
*
*  The source code for this program is not published or otherwise divested
*  of its trade secrets, irrespective of what has been deposited with the
*  U.S. Copyright Office. */

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMSchematicsWorkspaceDataSource_basic(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
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
