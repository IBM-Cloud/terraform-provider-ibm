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
			resource.TestStep{
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
	data "ibm_schematics_state" "test" {
		workspace_id = "%s"
		template_id = "%s"
	  }
	  
	  output "state_store_values" {
		value = "${data.ibm_schematics_state.test.state_store}"
	  }
`, WorkspaceID, templateID)
}
