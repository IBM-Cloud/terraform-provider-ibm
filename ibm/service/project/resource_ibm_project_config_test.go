// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package project_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	"github.com/IBM/project-go-sdk/projectv1"
	acc "github.com/Mavrickk3/terraform-provider-ibm/ibm/acctest"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/conns"
	"github.com/Mavrickk3/terraform-provider-ibm/ibm/flex"
)

func TestAccIbmProjectConfigBasic(t *testing.T) {
	var conf projectv1.ProjectConfig

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmProjectConfigDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmProjectConfigConfigBasic(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmProjectConfigExists("ibm_project_config.project_config_instance", conf),
				),
			},
			resource.TestStep{
				ResourceName:            "ibm_project_config.project_config_instance",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"project_id"},
			},
		},
	})
}

func testAccCheckIbmProjectConfigConfigBasic() string {
	return fmt.Sprintf(`
		resource "ibm_project" "project_instance" {
			location = "ca-tor"
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
	`, acc.ProjectsConfigApiKey)
}

func testAccCheckIbmProjectConfigExists(n string, obj projectv1.ProjectConfig) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		projectClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).ProjectV1()
		if err != nil {
			return err
		}

		getConfigOptions := &projectv1.GetConfigOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getConfigOptions.SetProjectID(parts[0])
		getConfigOptions.SetID(parts[1])

		projectConfig, _, err := projectClient.GetConfig(getConfigOptions)
		if err != nil {
			return err
		}

		obj = *projectConfig
		return nil
	}
}

func testAccCheckIbmProjectConfigDestroy(s *terraform.State) error {
	projectClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).ProjectV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_project_config" {
			continue
		}

		getConfigOptions := &projectv1.GetConfigOptions{}

		parts, err := flex.SepIdParts(rs.Primary.ID, "/")
		if err != nil {
			return err
		}

		getConfigOptions.SetProjectID(parts[0])
		getConfigOptions.SetID(parts[1])

		// Try to find the key
		_, response, err := projectClient.GetConfig(getConfigOptions)

		if err == nil {
			return fmt.Errorf("project_config still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for project_config (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
