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

func TestAccIBMSchematicsActionDataSourceBasic(t *testing.T) {

	actionName := fmt.Sprintf("acc-test-schematics-actions_%d", acctest.RandIntRange(10, 65))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMSchematicsActionDataSourceConfigBasic(actionName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "description"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "location"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "resource_group"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "user_state.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "created_by"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "sys_lock.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "state.#"),
				),
			},
		},
	})
}

func testAccCheckIBMSchematicsActionDataSourceConfigBasic(actionName string) string {
	return fmt.Sprintf(`
		resource "ibm_schematics_action" "schematics_action" {
			name = "%s"
			description = "acc-test-schematics-actions"
			location = "us-east"
			resource_group = "default"
		}

		data "ibm_schematics_action" "schematics_action" {
			action_id = ibm_schematics_action.schematics_action.id
		}
	`, actionName)
}
