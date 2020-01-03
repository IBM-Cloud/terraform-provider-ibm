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
