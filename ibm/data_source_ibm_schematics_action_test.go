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

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMSchematicsActionDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "action_id"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "description"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "location"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "resource_group"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "tags.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "user_state.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "source_readme_url"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "source.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "source_type"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "command_parameter"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "bastion.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "targets_ini"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "credentials.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "action_inputs.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "action_outputs.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "settings.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "trigger_record_id"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "account"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "source_created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "source_created_by"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "source_updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "source_updated_by"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "created_by"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "updated_by"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "namespace"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "state.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "playbook_names.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "sys_lock.#"),
				),
			},
		},
	})
}

func TestAccIBMSchematicsActionDataSourceAllArgs(t *testing.T) {
	actionName := fmt.Sprintf("name_%d", acctest.RandIntRange(10, 100))
	actionDescription := fmt.Sprintf("description_%d", acctest.RandIntRange(10, 100))
	actionLocation := "us_south"
	actionResourceGroup := fmt.Sprintf("resource_group_%d", acctest.RandIntRange(10, 100))
	actionSourceReadmeURL := fmt.Sprintf("source_readme_url_%d", acctest.RandIntRange(10, 100))
	actionSourceType := "local"
	actionCommandParameter := fmt.Sprintf("command_parameter_%d", acctest.RandIntRange(10, 100))
	actionTargetsIni := fmt.Sprintf("targets_ini_%d", acctest.RandIntRange(10, 100))
	actionTriggerRecordID := fmt.Sprintf("trigger_record_id_%d", acctest.RandIntRange(10, 100))
	actionXGithubToken := fmt.Sprintf("x_github_token_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMSchematicsActionDataSourceConfig(actionName, actionDescription, actionLocation, actionResourceGroup, actionSourceReadmeURL, actionSourceType, actionCommandParameter, actionTargetsIni, actionTriggerRecordID, actionXGithubToken),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "action_id"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "description"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "location"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "resource_group"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "tags.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "user_state.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "source_readme_url"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "source.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "source_type"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "command_parameter"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "bastion.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "targets_ini"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "credentials.#"),
					resource.TestCheckResourceAttr("data.ibm_schematics_action.schematics_action", "credentials.0.name", actionName),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "credentials.0.value"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "credentials.0.link"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "action_inputs.#"),
					resource.TestCheckResourceAttr("data.ibm_schematics_action.schematics_action", "action_inputs.0.name", actionName),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "action_inputs.0.value"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "action_inputs.0.link"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "action_outputs.#"),
					resource.TestCheckResourceAttr("data.ibm_schematics_action.schematics_action", "action_outputs.0.name", actionName),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "action_outputs.0.value"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "action_outputs.0.link"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "settings.#"),
					resource.TestCheckResourceAttr("data.ibm_schematics_action.schematics_action", "settings.0.name", actionName),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "settings.0.value"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "settings.0.link"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "trigger_record_id"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "account"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "source_created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "source_created_by"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "source_updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "source_updated_by"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "created_by"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "updated_by"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "namespace"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "state.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "playbook_names.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "sys_lock.#"),
				),
			},
		},
	})
}

func testAccCheckIBMSchematicsActionDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		resource "ibm_schematics_action" "schematics_action" {
		}

		data "ibm_schematics_action" "schematics_action" {
			action_id = "action_id"
		}
	`)
}

func testAccCheckIBMSchematicsActionDataSourceConfig(actionName string, actionDescription string, actionLocation string, actionResourceGroup string, actionSourceReadmeURL string, actionSourceType string, actionCommandParameter string, actionTargetsIni string, actionTriggerRecordID string, actionXGithubToken string) string {
	return fmt.Sprintf(`
		resource "ibm_schematics_action" "schematics_action" {
			name = "%s"
			description = "%s"
			location = "%s"
			resource_group = "%s"
			tags = "FIXME"
			user_state = { example: "object" }
			source_readme_url = "%s"
			source {
				source_type = "local"
			}
			source_type = "%s"
			command_parameter = "%s"
			bastion = { example: "object" }
			targets_ini = "%s"
			credentials = { example: "object" }
			inputs = { example: "object" }
			outputs = { example: "object" }
			settings = { example: "object" }
			trigger_record_id = "%s"
			state = { example: "object" }
			sys_lock = { example: "object" }
			X-Github-token = "%s"
		}

		data "ibm_schematics_action" "schematics_action" {
			action_id = "action_id"
		}
	`, actionName, actionDescription, actionLocation, actionResourceGroup, actionSourceReadmeURL, actionSourceType, actionCommandParameter, actionTargetsIni, actionTriggerRecordID, actionXGithubToken)
}
