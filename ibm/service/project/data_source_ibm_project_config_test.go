// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package project_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmProjectConfigDataSourceBasic(t *testing.T) {
	projectConfigName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	projectConfigLocatorID := fmt.Sprintf("1082e7d2-5e2f-0a11-a3bc-f88a8e1931fc.cd596f95-95a2-4f21-9b84-477f21fd1e95-global")

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmProjectConfigDataSourceConfigBasic(projectConfigName, projectConfigLocatorID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config_instance", "project_id"),
					resource.TestCheckResourceAttr("data.ibm_project_config.project_config_instance", "name", projectConfigName),
					resource.TestCheckResourceAttr("data.ibm_project_config.project_config_instance", "locator_id", projectConfigLocatorID),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config_instance", "type"),
				),
			},
		},
	})
}

func TestAccIbmProjectConfigDataSourceAllArgs(t *testing.T) {
	projectConfigName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	projectConfigLocatorID := fmt.Sprintf("1082e7d2-5e2f-0a11-a3bc-f88a8e1931fc.cd596f95-95a2-4f21-9b84-477f21fd1e95-global")
	projectConfigDescription := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmProjectConfigDataSourceConfig(projectConfigName, projectConfigLocatorID, projectConfigDescription),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config_instance", "project_id"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config_instance", "version"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config_instance", "project_config_id"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config_instance", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config_instance", "description"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config_instance", "authorizations.#"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config_instance", "compliance_profile.#"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config_instance", "locator_id"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config_instance", "type"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config_instance", "input.#"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config_instance", "output.#"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config_instance", "output.0.name"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config_instance", "setting.#"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config_instance", "setting.0.name"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config_instance", "setting.0.value"),
				),
			},
		},
	})
}

func testAccCheckIbmProjectConfigDataSourceConfigBasic(projectConfigName string, projectConfigLocatorID string) string {
	return fmt.Sprintf(`
		resource "ibm_project" "project_instance" {
			resource_group = "Default"
			location = "us-south"
			name = "acme-microservice-1"
		}

		resource "ibm_project_config" "project_config_instance" {
			project_id = ibm_project.project_instance.id
			name = "%s"
			locator_id = "%s"
		}

		data "ibm_project_config" "project_config_instance" {
			project_id = ibm_project.project_instance.id
			id = ibm_project_config.project_config_instance.project_config_id
			version = "draft"
		}
	`, projectConfigName, projectConfigLocatorID)
}

func testAccCheckIbmProjectConfigDataSourceConfig(projectConfigName string, projectConfigLocatorID string, projectConfigDescription string) string {
	return fmt.Sprintf(`
		resource "ibm_project" "project_instance" {
			resource_group = "Default"
			location = "us-south"
			name = "acme-microservice-2"
		}

		resource "ibm_project_config" "project_config_instance" {
			project_id = ibm_project.project_instance.id
			name = "%s"
			locator_id = "%s"
			labels = [ "labels" ]
			description = "%s"
			authorizations {
				method = "API_KEY"
				api_key = "xxx"
			}
			input {
				name = "name"
				value = "anything as a string"
			}
			setting {
				name = "name"
				value = "value"
			}
		}

		data "ibm_project_config" "project_config_instance" {
			project_id = ibm_project.project_instance.id
			id = ibm_project_config.project_config_instance.project_config_id
			version = "draft"
		}
	`, projectConfigName, projectConfigLocatorID, projectConfigDescription)
}
