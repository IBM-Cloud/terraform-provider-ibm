// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package project_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/flex"
	"github.com/IBM/project-go-sdk/projectv1"
)

func TestAccIbmProjectEnvironmentBasic(t *testing.T) {
	var conf projectv1.Environment

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmProjectEnvironmentDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmProjectEnvironmentConfigBasic(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmProjectEnvironmentExists("ibm_project_environment.project_environment_instance", conf),
				),
			},
			resource.TestStep{
				ResourceName:            "ibm_project_environment.project_environment_instance",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"project_id"},
			},
		},
	})
}

func testAccCheckIbmProjectEnvironmentConfigBasic() string {
	return fmt.Sprintf(`
		resource "ibm_project" "project_instance" {
            location = "us-south"
            resource_group = "Default"
            definition {
                name = "acme-microservice"
                description = "acme-microservice description"
                destroy_on_delete = true
                monitoring_enabled = true
                auto_deploy = true
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
	`, acc.ProjectsConfigApiKey)
}

func testAccCheckIbmProjectEnvironmentExists(n string, obj projectv1.Environment) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		projectClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).ProjectV1()
		if err != nil {
			return err
		}

		getProjectEnvironmentOptions := &projectv1.GetProjectEnvironmentOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getProjectEnvironmentOptions.SetProjectID(parts[0])
		getProjectEnvironmentOptions.SetID(parts[1])

		environment, _, err := projectClient.GetProjectEnvironment(getProjectEnvironmentOptions)
		if err != nil {
			return err
		}

		obj = *environment
		return nil
	}
}

func testAccCheckIbmProjectEnvironmentDestroy(s *terraform.State) error {
	projectClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).ProjectV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_project_environment" {
			continue
		}

		getProjectEnvironmentOptions := &projectv1.GetProjectEnvironmentOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getProjectEnvironmentOptions.SetProjectID(parts[0])
		getProjectEnvironmentOptions.SetID(parts[1])

		// Try to find the key
		_, response, err := projectClient.GetProjectEnvironment(getProjectEnvironmentOptions)

		if err == nil {
			return fmt.Errorf("project_environment still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for project_environment (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
