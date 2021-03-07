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
		},
	})
}

func testAccCheckIBMSchematicsWorkspaceConfigBasic() string {
	return fmt.Sprintf(`

		resource "ibm_schematics_workspace" "schematics_workspace" {
			description = "tf-acc-test-schematics"
			name = "tf-acc-test-schematics"
			location = "us-east"
			resource_group = "default"
			template_type = "terraform_v0.12.20"
		}
	`)
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
