package ibm

import (
	"fmt"
	"os"
	"testing"

	"github.com/hashicorp/terraform/helper/resource"
)

func init() {
	workspaceID = os.Getenv("WORKSPACE_ID")
	if workspaceID == "" {
		workspaceID = "outwork-2737f163-b966-44"
		fmt.Println("[INFO] Set the environment variable WORKSPACE_ID for testing data_source_ibm_schematics_state_test else it is set to default value null")
	}
	templateID = os.Getenv("TEMPLATE_ID")
	if templateID == "" {
		templateID = "653f60a4-f64f-41"
		fmt.Println("[INFO] Set the environment variable TEMPLATE_ID for testing data_source_ibm_schematics_state_test else it is set to default value null")
	}

}

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
	  
	  output "WorkSpace Values" {
		value = "${data.ibm_schematics_output.test.output_values}"
	  }
`, WorkspaceID, templateID)
}
