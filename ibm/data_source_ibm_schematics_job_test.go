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

func TestAccIBMSchematicsJobDataSourceBasic(t *testing.T) {
	jobRefreshToken := fmt.Sprintf("refresh_token_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMSchematicsJobDataSourceConfigBasic(jobRefreshToken),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "job_id"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "command_object"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "command_object_id"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "command_name"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "command_parameter"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "command_options.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "job_inputs.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "job_env_settings.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "tags.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "description"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "location"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "resource_group"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "submitted_at"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "submitted_by"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "start_at"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "end_at"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "duration"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "status.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "data.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "targets_ini"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "bastion.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "job_log_summary.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "log_store_url"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "state_store_url"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "results_url"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "updated_at"),
				),
			},
		},
	})
}

func TestAccIBMSchematicsJobDataSourceAllArgs(t *testing.T) {
	jobRefreshToken := fmt.Sprintf("refresh_token_%d", acctest.RandIntRange(10, 100))
	jobCommandObject := "workspace"
	jobCommandObjectID := fmt.Sprintf("command_object_id_%d", acctest.RandIntRange(10, 100))
	jobCommandName := "workspace_init_flow"
	jobCommandParameter := fmt.Sprintf("command_parameter_%d", acctest.RandIntRange(10, 100))
	jobLocation := "us_south"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMSchematicsJobDataSourceConfig(jobRefreshToken, jobCommandObject, jobCommandObjectID, jobCommandName, jobCommandParameter, jobLocation),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "job_id"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "command_object"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "command_object_id"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "command_name"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "command_parameter"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "command_options.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "job_inputs.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "job_inputs.0.name"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "job_inputs.0.value"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "job_inputs.0.link"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "job_env_settings.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "job_env_settings.0.name"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "job_env_settings.0.value"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "job_env_settings.0.link"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "tags.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "description"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "location"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "resource_group"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "submitted_at"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "submitted_by"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "start_at"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "end_at"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "duration"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "status.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "data.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "targets_ini"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "bastion.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "job_log_summary.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "log_store_url"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "state_store_url"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "results_url"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_job.schematics_job", "updated_at"),
				),
			},
		},
	})
}

func testAccCheckIBMSchematicsJobDataSourceConfigBasic(jobRefreshToken string) string {
	return fmt.Sprintf(`
		resource "ibm_schematics_job" "schematics_job" {
			refresh_token = "%s"
		}

		data "ibm_schematics_job" "schematics_job" {
			job_id = "job_id"
		}
	`, jobRefreshToken)
}

func testAccCheckIBMSchematicsJobDataSourceConfig(jobRefreshToken string, jobCommandObject string, jobCommandObjectID string, jobCommandName string, jobCommandParameter string, jobLocation string) string {
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

		data "ibm_schematics_job" "schematics_job" {
			job_id = "job_id"
		}
	`, jobRefreshToken, jobCommandObject, jobCommandObjectID, jobCommandName, jobCommandParameter, jobLocation)
}
