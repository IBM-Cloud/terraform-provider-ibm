// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package project_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmProjectEnvironmentDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmProjectEnvironmentDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_project_environment.project_environment_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_project_environment.project_environment_instance", "project_id"),
					resource.TestCheckResourceAttrSet("data.ibm_project_environment.project_environment_instance", "project_environment_id"),
					resource.TestCheckResourceAttrSet("data.ibm_project_environment.project_environment_instance", "project.#"),
					resource.TestCheckResourceAttrSet("data.ibm_project_environment.project_environment_instance", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_project_environment.project_environment_instance", "modified_at"),
					resource.TestCheckResourceAttrSet("data.ibm_project_environment.project_environment_instance", "href"),
					resource.TestCheckResourceAttrSet("data.ibm_project_environment.project_environment_instance", "definition.#"),
				),
			},
		},
	})
}

func testAccCheckIbmProjectEnvironmentDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		resource "ibm_project" "project_instance" {
            location = "us-south"
            resource_group = "Default"
            definition {
                name = "acme-microservice"
                description = "acme-microservice description"
                destroy_on_delete = true
                monitoring_enabled = true
            }
        }

        resource "ibm_project_environment" "project_environment_instance" {
            project_id = ibm_project.project_instance.id
            definition {
                name = "environment-stage"
                description = "environment for stage project"
                authorizations {
                    method = "api_key"
                    api_key = "%s"
               }
            }
            lifecycle {
                ignore_changes = [
                    definition[0].authorizations[0].api_key,
                ]
            }
        }

		data "ibm_project_environment" "project_environment_instance" {
			project_id = ibm_project_environment.project_environment_instance.project_id
			project_environment_id = ibm_project_environment.project_environment_instance.project_environment_id
		}
	`, acc.ProjectsConfigApiKey)
}
