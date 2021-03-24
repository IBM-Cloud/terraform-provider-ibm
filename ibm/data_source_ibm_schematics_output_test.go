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

func TestAccIBMSchematicsOutputDataSourceBasic(t *testing.T) {
	wID := workspaceID

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMSchematicsOutputDataSourceConfigBasic(wID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_schematics_output.schematics_output", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_output.schematics_output", "output_values.#"),
				),
			},
		},
	})
}

func testAccCheckIBMSchematicsOutputDataSourceConfigBasic(wID string) string {
	return fmt.Sprintf(`
		  data "ibm_schematics_output" "schematics_output" {
			workspace_id = "%s"
		  }
	  `, wID)
}
