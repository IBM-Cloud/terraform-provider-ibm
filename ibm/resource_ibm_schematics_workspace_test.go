/**
 * (C) Copyright IBM Corp. 2021.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package ibm

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/IBM/schematics-go-sdk/schematicsv1"
)

func TestAccIBMSchematicsWorkspaceBasic(t *testing.T) {
	var conf schematicsv1.WorkspaceResponse

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMSchematicsWorkspaceDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMSchematicsWorkspaceConfigBasic(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMSchematicsWorkspaceExists("ibm_schematics_workspace.schematics_workspace", conf),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMSchematicsWorkspaceConfigBasic(),
				Check: resource.ComposeAggregateTestCheckFunc(
				),
			},
		},
	})
}

func TestAccIBMSchematicsWorkspaceAllArgs(t *testing.T) {
	var conf schematicsv1.WorkspaceResponse
	description := fmt.Sprintf("description_%d", acctest.RandIntRange(10, 100))
	location := fmt.Sprintf("location_%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("name_%d", acctest.RandIntRange(10, 100))
	resourceGroup := fmt.Sprintf("resource_group_%d", acctest.RandIntRange(10, 100))
	templateRef := fmt.Sprintf("template_ref_%d", acctest.RandIntRange(10, 100))
	xGithubToken := fmt.Sprintf("X-Github-token_%d", acctest.RandIntRange(10, 100))
	descriptionUpdate := fmt.Sprintf("description_%d", acctest.RandIntRange(10, 100))
	locationUpdate := fmt.Sprintf("location_%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("name_%d", acctest.RandIntRange(10, 100))
	resourceGroupUpdate := fmt.Sprintf("resource_group_%d", acctest.RandIntRange(10, 100))
	templateRefUpdate := fmt.Sprintf("template_ref_%d", acctest.RandIntRange(10, 100))
	xGithubTokenUpdate := fmt.Sprintf("X-Github-token_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMSchematicsWorkspaceDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMSchematicsWorkspaceConfig(description, location, name, resourceGroup, templateRef, xGithubToken),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMSchematicsWorkspaceExists("ibm_schematics_workspace.schematics_workspace", conf),
					resource.TestCheckResourceAttr("ibm_schematics_workspace.schematics_workspace", "description", description),
					resource.TestCheckResourceAttr("ibm_schematics_workspace.schematics_workspace", "location", location),
					resource.TestCheckResourceAttr("ibm_schematics_workspace.schematics_workspace", "name", name),
					resource.TestCheckResourceAttr("ibm_schematics_workspace.schematics_workspace", "resource_group", resourceGroup),
					resource.TestCheckResourceAttr("ibm_schematics_workspace.schematics_workspace", "template_ref", templateRef),
					resource.TestCheckResourceAttr("ibm_schematics_workspace.schematics_workspace", "X-Github-token", xGithubToken),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMSchematicsWorkspaceConfig(descriptionUpdate, locationUpdate, nameUpdate, resourceGroupUpdate, templateRefUpdate, xGithubTokenUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_schematics_workspace.schematics_workspace", "description", descriptionUpdate),
					resource.TestCheckResourceAttr("ibm_schematics_workspace.schematics_workspace", "location", locationUpdate),
					resource.TestCheckResourceAttr("ibm_schematics_workspace.schematics_workspace", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_schematics_workspace.schematics_workspace", "resource_group", resourceGroupUpdate),
					resource.TestCheckResourceAttr("ibm_schematics_workspace.schematics_workspace", "template_ref", templateRefUpdate),
					resource.TestCheckResourceAttr("ibm_schematics_workspace.schematics_workspace", "X-Github-token", xGithubTokenUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_schematics_workspace.schematics_workspace",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMSchematicsWorkspaceConfigBasic() string {
	return fmt.Sprintf(`

		resource "ibm_schematics_workspace" "schematics_workspace" {
		}
	`, )
}

func testAccCheckIBMSchematicsWorkspaceConfig(description string, location string, name string, resourceGroup string, templateRef string, xGithubToken string) string {
	return fmt.Sprintf(`

		resource "ibm_schematics_workspace" "schematics_workspace" {
			applied_shareddata_ids = "FIXME"
			catalog_ref = { example: "object" }
			description = "%s"
			location = "%s"
			name = "%s"
			resource_group = "%s"
			shared_data = { example: "object" }
			tags = "FIXME"
			template_data = { example: "object" }
			template_ref = "%s"
			template_repo = { example: "object" }
			type = "FIXME"
			workspace_status = { example: "object" }
			X-Github-token = "%s"
		}
	`, description, location, name, resourceGroup, templateRef, xGithubToken)
}

func testAccCheckIBMSchematicsWorkspaceExists(n string, obj schematicsv1.WorkspaceResponse) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		schematicsClient, err := testAccProvider.Meta().(ClientSession).SchematicsV1()
		if err != nil {
			return err
		}

		getWorkspaceOptions := &schematicsv1.GetWorkspaceOptions{}

		getWorkspaceOptions.SetWID(rs.Primary.ID)

		workspaceResponse, _, err := schematicsClient.GetWorkspace(getWorkspaceOptions)
		if err != nil {
			return err
		}

		obj = *workspaceResponse
		return nil
	}
}

func testAccCheckIBMSchematicsWorkspaceDestroy(s *terraform.State) error {
	schematicsClient, err := testAccProvider.Meta().(ClientSession).SchematicsV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_schematics_workspace" {
			continue
		}

		getWorkspaceOptions := &schematicsv1.GetWorkspaceOptions{}

		getWorkspaceOptions.SetWID(rs.Primary.ID)

		// Try to find the key
		_, response, err := schematicsClient.GetWorkspace(getWorkspaceOptions)

		if err == nil {
			return fmt.Errorf("schematics_workspace still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for schematics_workspace (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
