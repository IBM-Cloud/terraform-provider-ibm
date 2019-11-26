package ibm

import (
	"fmt"
	"testing"

	"log"

	bluemix "github.com/IBM-Cloud/bluemix-go"
	sch "github.com/IBM-Cloud/bluemix-go/api/schematics"
	"github.com/IBM-Cloud/bluemix-go/session"
	"github.com/hashicorp/terraform/helper/acctest"
	"github.com/hashicorp/terraform/helper/resource"
)

func TestAccIBMSchematicsWorkspaceDataSource_basic(t *testing.T) {
	workspaceName := fmt.Sprintf("terraform_%d", acctest.RandInt())

	var payload = sch.Payload{
		Name:        workspaceName,
		Type:        []string{"terraform-v1.0"},
		Description: "terraform workspace",
		Tags:        []string{"department:HR", "application:compensation", "environment:staging"},
		WorkspaceStatus: sch.WorkspaceStatus{
			Frozen: true,
		},
		TemplateRepo: sch.TemplateRepo{
			URL: "https://github.com/ptaube/tf_cloudless_sleepy",
		},
		TemplateRef: "ibm-open-liberty-2ae855ce3ca4",
		TemplateData: []sch.TemplateData{
			{
				Folder: ".",
				Type:   "terraform-v1.0",

				Variablestore: []sch.Variablestore{
					{
						Name:        "sample_var",
						Secure:      false,
						Value:       "THIS IS IBM CLOUD TERRAFORM CLI DEMO",
						Description: "Description of sample_var",
					},
					{
						Name:  "sleepy_time",
						Value: "15",
					},
				},
			},
		},
	}

	c := new(bluemix.Config)
	sess, err := session.New(c)
	if err != nil {
		log.Fatal(err)
	}
	schClient, err := sch.New(sess)
	if err != nil {
		log.Fatal(err)
	}
	wrkAPI := schClient.Workspaces()

	WorkspaceInfo, err := wrkAPI.CreateWorkspace(payload)

	WorkspaceID := WorkspaceInfo.ID

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMSchematicsWorkspaceDataSourceConfig(WorkspaceID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.ibm_schematics_workspace.test", "workspace_id", WorkspaceID),
				),
			},
		},
	})
}

func testAccCheckIBMSchematicsWorkspaceDataSourceConfig(WorkspaceID string) string {
	return fmt.Sprintf(`
	data "ibm_schematics_workspace" "test" {
		workspace_id = "%s"
	  }
	  
	  output "WorkSpace Values" {
		value = "${data.ibm_schematics_workspace.test.template_id.0}"
	  }
`, WorkspaceID)
}
