package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
)

func TestAccIBMSchematicsOutputDataSource_basic(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMSchematicsOutputDataSourceConfig(workspaceID, templateID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_schematics_output.test", "workspace_id", workspaceID),
				),
			},
		},
	})
}

func testAccCheckIBMSchematicsOutputDataSourceConfig(WorkspaceID, templateID string) string {
	return fmt.Sprintf(`
	data "ibm_schematics_output" "test" {
		workspace_id = "%s"
		template_id = "%s"
	  }
	  
	  output "output_values" {
		value = data.ibm_schematics_output.test.output_values
	  }
`, WorkspaceID, templateID)
}
