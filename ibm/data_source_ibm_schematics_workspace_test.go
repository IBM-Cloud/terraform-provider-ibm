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
)

func TestAccIBMSchematicsWorkspaceDataSourceBasic(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMSchematicsWorkspaceDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "workspace_id"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "applied_shareddata_ids.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "catalog_ref.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "created_by"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "description"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "last_health_check_at"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "location"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "resource_group"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "runtime_data.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "shared_data.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "status"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "tags.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "template_data.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "template_ref"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "template_repo.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "type.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "updated_by"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "workspace_status.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "workspace_status_msg.#"),
				),
			},
		},
	})
}

func TestAccIBMSchematicsWorkspaceDataSourceAllArgs(t *testing.T) {
	workspaceResponseDescription := fmt.Sprintf("description_%d", acctest.RandIntRange(10, 100))
	workspaceResponseLocation := fmt.Sprintf("location_%d", acctest.RandIntRange(10, 100))
	workspaceResponseName := fmt.Sprintf("name_%d", acctest.RandIntRange(10, 100))
	workspaceResponseResourceGroup := fmt.Sprintf("resource_group_%d", acctest.RandIntRange(10, 100))
	workspaceResponseTemplateRef := fmt.Sprintf("template_ref_%d", acctest.RandIntRange(10, 100))
	workspaceResponseXGithubToken := fmt.Sprintf("x_github_token_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMSchematicsWorkspaceDataSourceConfig(workspaceResponseDescription, workspaceResponseLocation, workspaceResponseName, workspaceResponseResourceGroup, workspaceResponseTemplateRef, workspaceResponseXGithubToken),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "workspace_id"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "applied_shareddata_ids.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "catalog_ref.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "created_by"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "description"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "last_health_check_at"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "location"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "resource_group"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "runtime_data.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "runtime_data.0.engine_cmd"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "runtime_data.0.engine_name"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "runtime_data.0.engine_version"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "runtime_data.0.id"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "runtime_data.0.log_store_url"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "runtime_data.0.state_store_url"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "shared_data.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "status"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "tags.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "template_data.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "template_data.0.folder"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "template_data.0.has_githubtoken"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "template_data.0.id"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "template_data.0.type"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "template_data.0.uninstall_script_name"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "template_data.0.values"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "template_data.0.values_url"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "template_ref"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "template_repo.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "type.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "updated_by"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "workspace_status.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "workspace_status_msg.#"),
				),
			},
		},
	})
}

func testAccCheckIBMSchematicsWorkspaceDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		 resource "ibm_schematics_workspace" "schematics_workspace" {
		 }
 
		 data "ibm_schematics_workspace" "schematics_workspace" {
			 workspace_id = "workspace_id"
		 }
	 `)
}

func testAccCheckIBMSchematicsWorkspaceDataSourceConfig(workspaceResponseDescription string, workspaceResponseLocation string, workspaceResponseName string, workspaceResponseResourceGroup string, workspaceResponseTemplateRef string, workspaceResponseXGithubToken string) string {
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
 
		 data "ibm_schematics_workspace" "schematics_workspace" {
			 workspace_id = "workspace_id"
		 }
	 `, workspaceResponseDescription, workspaceResponseLocation, workspaceResponseName, workspaceResponseResourceGroup, workspaceResponseTemplateRef, workspaceResponseXGithubToken)
}
