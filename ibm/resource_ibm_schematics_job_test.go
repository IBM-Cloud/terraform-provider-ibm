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

func TestAccIBMSchematicsJobBasic(t *testing.T) {
	var conf schematicsv1.Job
	refreshToken := fmt.Sprintf("refresh_token_%d", acctest.RandIntRange(10, 100))
	refreshTokenUpdate := fmt.Sprintf("refresh_token_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMSchematicsJobDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMSchematicsJobConfigBasic(refreshToken),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMSchematicsJobExists("ibm_schematics_job.schematics_job", conf),
					resource.TestCheckResourceAttr("ibm_schematics_job.schematics_job", "refresh_token", refreshToken),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMSchematicsJobConfigBasic(refreshTokenUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_schematics_job.schematics_job", "refresh_token", refreshTokenUpdate),
				),
			},
		},
	})
}

func TestAccIBMSchematicsJobAllArgs(t *testing.T) {
	var conf schematicsv1.Job
	refreshToken := fmt.Sprintf("refresh_token_%d", acctest.RandIntRange(10, 100))
	commandObject := "workspace"
	commandObjectID := fmt.Sprintf("command_object_id_%d", acctest.RandIntRange(10, 100))
	commandName := "workspace_init_flow"
	commandParameter := fmt.Sprintf("command_parameter_%d", acctest.RandIntRange(10, 100))
	location := "us_south"
	refreshTokenUpdate := fmt.Sprintf("refresh_token_%d", acctest.RandIntRange(10, 100))
	commandObjectUpdate := "action"
	commandObjectIDUpdate := fmt.Sprintf("command_object_id_%d", acctest.RandIntRange(10, 100))
	commandNameUpdate := "opa_evaluate"
	commandParameterUpdate := fmt.Sprintf("command_parameter_%d", acctest.RandIntRange(10, 100))
	locationUpdate := "eu_de"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckIBMSchematicsJobDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMSchematicsJobConfig(refreshToken, commandObject, commandObjectID, commandName, commandParameter, location),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMSchematicsJobExists("ibm_schematics_job.schematics_job", conf),
					resource.TestCheckResourceAttr("ibm_schematics_job.schematics_job", "refresh_token", refreshToken),
					resource.TestCheckResourceAttr("ibm_schematics_job.schematics_job", "command_object", commandObject),
					resource.TestCheckResourceAttr("ibm_schematics_job.schematics_job", "command_object_id", commandObjectID),
					resource.TestCheckResourceAttr("ibm_schematics_job.schematics_job", "command_name", commandName),
					resource.TestCheckResourceAttr("ibm_schematics_job.schematics_job", "command_parameter", commandParameter),
					resource.TestCheckResourceAttr("ibm_schematics_job.schematics_job", "location", location),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMSchematicsJobConfig(refreshTokenUpdate, commandObjectUpdate, commandObjectIDUpdate, commandNameUpdate, commandParameterUpdate, locationUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_schematics_job.schematics_job", "refresh_token", refreshTokenUpdate),
					resource.TestCheckResourceAttr("ibm_schematics_job.schematics_job", "command_object", commandObjectUpdate),
					resource.TestCheckResourceAttr("ibm_schematics_job.schematics_job", "command_object_id", commandObjectIDUpdate),
					resource.TestCheckResourceAttr("ibm_schematics_job.schematics_job", "command_name", commandNameUpdate),
					resource.TestCheckResourceAttr("ibm_schematics_job.schematics_job", "command_parameter", commandParameterUpdate),
					resource.TestCheckResourceAttr("ibm_schematics_job.schematics_job", "location", locationUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_schematics_job.schematics_job",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMSchematicsJobConfigBasic(refreshToken string) string {
	return fmt.Sprintf(`

		resource "ibm_schematics_job" "schematics_job" {
			refresh_token = "%s"
		}
	`, refreshToken)
}

func testAccCheckIBMSchematicsJobConfig(refreshToken string, commandObject string, commandObjectID string, commandName string, commandParameter string, location string) string {
	return fmt.Sprintf(`

		resource "ibm_schematics_job" "schematics_job" {
			refresh_token = "%s"
			command_object = "%s"
			command_object_id = "%s"
			command_name = "%s"
			command_parameter = "%s"
			command_options = "FIXME"
			inputs = { example: "object" }
			settings = { example: "object" }
			tags = "FIXME"
			location = "%s"
			status = { example: "object" }
			data {
				job_type = "repo_download_job"
			}
			bastion = { example: "object" }
			log_summary = { example: "object" }
		}
	`, refreshToken, commandObject, commandObjectID, commandName, commandParameter, location)
}

func testAccCheckIBMSchematicsJobExists(n string, obj schematicsv1.Job) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		schematicsClient, err := testAccProvider.Meta().(ClientSession).SchematicsV1()
		if err != nil {
			return err
		}

		getJobOptions := &schematicsv1.GetJobOptions{}

		getJobOptions.SetJobID(rs.Primary.ID)

		job, _, err := schematicsClient.GetJob(getJobOptions)
		if err != nil {
			return err
		}

		obj = *job
		return nil
	}
}

func testAccCheckIBMSchematicsJobDestroy(s *terraform.State) error {
	schematicsClient, err := testAccProvider.Meta().(ClientSession).SchematicsV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_schematics_job" {
			continue
		}

		getJobOptions := &schematicsv1.GetJobOptions{}

		getJobOptions.SetJobID(rs.Primary.ID)

		// Try to find the key
		_, response, err := schematicsClient.GetJob(getJobOptions)

		if err == nil {
			return fmt.Errorf("schematics_job still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for schematics_job (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
