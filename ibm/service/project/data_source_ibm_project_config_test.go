// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.92.1-44330004-20240620-143510
 */

package project_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmProjectConfigDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmProjectConfigDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config_instance", "project_id"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config_instance", "project_config_id"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config_instance", "version"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config_instance", "is_draft"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config_instance", "needs_attention_state.#"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config_instance", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config_instance", "modified_at"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config_instance", "outputs.#"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config_instance", "project.#"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config_instance", "state"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config_instance", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config_instance", "definition.#"),
				),
			},
		},
	})
}

func testAccCheckIbmProjectConfigDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		resource "ibm_project" "project_instance" {
			location = "us-south"
			resource_group = "Default"
			definition {
                name = "acme-microservice"
                description = "acme-microservice description"
                destroy_on_delete = true
                monitoring_enabled = true
                auto_deploy = false
            }
		}

		resource "ibm_project_config" "project_config_instance" {
			project_id = ibm_project.project_instance.id
            definition {
                name = "stage-environment"
                authorizations {
                    method = "api_key"
                    api_key = "%s"
                }
                locator_id = "1082e7d2-5e2f-0a11-a3bc-f88a8e1931fc.cd596f95-95a2-4f21-9b84-477f21fd1e95-global"
                inputs = {
                    app_repo_name = "grit-repo-name"
                }
            }
            lifecycle {
                ignore_changes = [
                    definition[0].authorizations[0].api_key,
                ]
            }
		}

		data "ibm_project_config" "project_config_instance" {
			project_id = ibm_project_config.project_config_instance.project_id
			project_config_id = ibm_project_config.project_config_instance.project_config_id
		}
	`, acc.ProjectsConfigApiKey)
}
