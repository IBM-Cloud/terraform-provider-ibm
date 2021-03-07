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
)

func TestAccIBMSchematicsStateDataSourceBasic(t *testing.T) {

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMSchematicsStateDataSourceConfigBasic(workspaceID, templateID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_schematics_state.schematics_state", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_state.schematics_state", "state_store"),
				),
			},
		},
	})
}

func testAccCheckIBMSchematicsStateDataSourceConfigBasic(workspaceID string, templateID string) string {
	return fmt.Sprintf(`
		 data "ibm_schematics_state" "schematics_state" {
			workspace_id = "%s"
			 template_id = "%s"
		 }
	 `, workspaceID, templateID)
}
