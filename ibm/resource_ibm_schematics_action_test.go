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

func TestAccIBMSchematicsActionBasic(t *testing.T) {
	var conf schematicsv1.Action

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMSchematicsActionDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMSchematicsActionConfigBasic(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMSchematicsActionExists("ibm_schematics_action.schematics_action", conf),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMSchematicsActionConfigBasic(),
				Check:  resource.ComposeAggregateTestCheckFunc(),
			},
		},
	})
}

func TestAccIBMSchematicsActionAllArgs(t *testing.T) {
	var conf schematicsv1.Action
	name := fmt.Sprintf("name_%d", acctest.RandIntRange(10, 100))
	description := fmt.Sprintf("description_%d", acctest.RandIntRange(10, 100))
	location := "us_south"
	resourceGroup := fmt.Sprintf("resource_group_%d", acctest.RandIntRange(10, 100))
	sourceReadmeURL := fmt.Sprintf("source_readme_url_%d", acctest.RandIntRange(10, 100))
	sourceType := "local"
	commandParameter := fmt.Sprintf("command_parameter_%d", acctest.RandIntRange(10, 100))
	targetsIni := fmt.Sprintf("targets_ini_%d", acctest.RandIntRange(10, 100))
	triggerRecordID := fmt.Sprintf("trigger_record_id_%d", acctest.RandIntRange(10, 100))
	xGithubToken := fmt.Sprintf("X-Github-token_%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("name_%d", acctest.RandIntRange(10, 100))
	descriptionUpdate := fmt.Sprintf("description_%d", acctest.RandIntRange(10, 100))
	locationUpdate := "eu_de"
	resourceGroupUpdate := fmt.Sprintf("resource_group_%d", acctest.RandIntRange(10, 100))
	sourceReadmeURLUpdate := fmt.Sprintf("source_readme_url_%d", acctest.RandIntRange(10, 100))
	sourceTypeUpdate := "external_scm"
	commandParameterUpdate := fmt.Sprintf("command_parameter_%d", acctest.RandIntRange(10, 100))
	targetsIniUpdate := fmt.Sprintf("targets_ini_%d", acctest.RandIntRange(10, 100))
	triggerRecordIDUpdate := fmt.Sprintf("trigger_record_id_%d", acctest.RandIntRange(10, 100))
	xGithubTokenUpdate := fmt.Sprintf("X-Github-token_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMSchematicsActionDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMSchematicsActionConfig(name, description, location, resourceGroup, sourceReadmeURL, sourceType, commandParameter, targetsIni, triggerRecordID, xGithubToken),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMSchematicsActionExists("ibm_schematics_action.schematics_action", conf),
					resource.TestCheckResourceAttr("ibm_schematics_action.schematics_action", "name", name),
					resource.TestCheckResourceAttr("ibm_schematics_action.schematics_action", "description", description),
					resource.TestCheckResourceAttr("ibm_schematics_action.schematics_action", "location", location),
					resource.TestCheckResourceAttr("ibm_schematics_action.schematics_action", "resource_group", resourceGroup),
					resource.TestCheckResourceAttr("ibm_schematics_action.schematics_action", "source_readme_url", sourceReadmeURL),
					resource.TestCheckResourceAttr("ibm_schematics_action.schematics_action", "source_type", sourceType),
					resource.TestCheckResourceAttr("ibm_schematics_action.schematics_action", "command_parameter", commandParameter),
					resource.TestCheckResourceAttr("ibm_schematics_action.schematics_action", "targets_ini", targetsIni),
					resource.TestCheckResourceAttr("ibm_schematics_action.schematics_action", "trigger_record_id", triggerRecordID),
					resource.TestCheckResourceAttr("ibm_schematics_action.schematics_action", "X-Github-token", xGithubToken),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMSchematicsActionConfig(nameUpdate, descriptionUpdate, locationUpdate, resourceGroupUpdate, sourceReadmeURLUpdate, sourceTypeUpdate, commandParameterUpdate, targetsIniUpdate, triggerRecordIDUpdate, xGithubTokenUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_schematics_action.schematics_action", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_schematics_action.schematics_action", "description", descriptionUpdate),
					resource.TestCheckResourceAttr("ibm_schematics_action.schematics_action", "location", locationUpdate),
					resource.TestCheckResourceAttr("ibm_schematics_action.schematics_action", "resource_group", resourceGroupUpdate),
					resource.TestCheckResourceAttr("ibm_schematics_action.schematics_action", "source_readme_url", sourceReadmeURLUpdate),
					resource.TestCheckResourceAttr("ibm_schematics_action.schematics_action", "source_type", sourceTypeUpdate),
					resource.TestCheckResourceAttr("ibm_schematics_action.schematics_action", "command_parameter", commandParameterUpdate),
					resource.TestCheckResourceAttr("ibm_schematics_action.schematics_action", "targets_ini", targetsIniUpdate),
					resource.TestCheckResourceAttr("ibm_schematics_action.schematics_action", "trigger_record_id", triggerRecordIDUpdate),
					resource.TestCheckResourceAttr("ibm_schematics_action.schematics_action", "X-Github-token", xGithubTokenUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_schematics_action.schematics_action",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMSchematicsActionConfigBasic() string {
	return fmt.Sprintf(`

		resource "ibm_schematics_action" "schematics_action" {
		}
	`)
}

func testAccCheckIBMSchematicsActionConfig(name string, description string, location string, resourceGroup string, sourceReadmeURL string, sourceType string, commandParameter string, targetsIni string, triggerRecordID string, xGithubToken string) string {
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
	`, name, description, location, resourceGroup, sourceReadmeURL, sourceType, commandParameter, targetsIni, triggerRecordID, xGithubToken)
}

func testAccCheckIBMSchematicsActionExists(n string, obj schematicsv1.Action) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		schematicsClient, err := testAccProvider.Meta().(ClientSession).SchematicsV1()
		if err != nil {
			return err
		}

		getActionOptions := &schematicsv1.GetActionOptions{}

		getActionOptions.SetActionID(rs.Primary.ID)

		action, _, err := schematicsClient.GetAction(getActionOptions)
		if err != nil {
			return err
		}

		obj = *action
		return nil
	}
}

func testAccCheckIBMSchematicsActionDestroy(s *terraform.State) error {
	schematicsClient, err := testAccProvider.Meta().(ClientSession).SchematicsV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_schematics_action" {
			continue
		}

		getActionOptions := &schematicsv1.GetActionOptions{}

		getActionOptions.SetActionID(rs.Primary.ID)

		// Try to find the key
		_, response, err := schematicsClient.GetAction(getActionOptions)

		if err == nil {
			return fmt.Errorf("schematics_action still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for schematics_action (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
