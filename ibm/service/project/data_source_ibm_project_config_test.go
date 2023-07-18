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
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config_instance", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config_instance", "locator_id"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config_instance", "type"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config_instance", "project_config_id"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config_instance", "version"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config_instance", "is_draft"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config_instance", "state"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config_instance", "update_available"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config_instance", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config_instance", "updated_at"),
				),
			},
		},
	})
}

func TestAccIbmProjectConfigDataSourceAllArgs(t *testing.T) {
	projectConfigName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	projectConfigDescription := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	projectConfigLocatorID := fmt.Sprintf("1082e7d2-5e2f-0a11-a3bc-f88a8e1931fc.cd596f95-95a2-4f21-9b84-477f21fd1e95-global")

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmProjectConfigDataSourceConfig(projectConfigName, projectConfigDescription, projectConfigLocatorID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config_instance", "project_id"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config_instance", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config_instance", "labels.#"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config_instance", "description"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config_instance", "authorizations.#"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config_instance", "compliance_profile.#"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config_instance", "locator_id"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config_instance", "type"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config_instance", "input.#"),
					resource.TestCheckResourceAttr("data.ibm_project_config.project_config_instance", "input.0.name", projectConfigName),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config_instance", "input.0.type"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config_instance", "input.0.value"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config_instance", "input.0.required"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config_instance", "output.#"),
					resource.TestCheckResourceAttr("data.ibm_project_config.project_config_instance", "output.0.name", projectConfigName),
					resource.TestCheckResourceAttr("data.ibm_project_config.project_config_instance", "output.0.description", projectConfigDescription),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config_instance", "output.0.value"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config_instance", "setting.#"),
					resource.TestCheckResourceAttr("data.ibm_project_config.project_config_instance", "setting.0.name", projectConfigName),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config_instance", "setting.0.value"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config_instance", "project_config_id"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config_instance", "version"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config_instance", "is_draft"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config_instance", "needs_attention_state"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config_instance", "needs_attention_state.#"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config_instance", "state"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config_instance", "pipeline_state"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config_instance", "update_available"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config_instance", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config_instance", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config_instance", "last_approved.#"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config_instance", "last_save"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config_instance", "job_summary.#"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config_instance", "cra_logs.#"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config_instance", "cost_estimate.#"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config_instance", "last_deployment_job_summary.#"),
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
			name = "acme-microservice"
		}

		resource "ibm_project_config" "project_config_instance" {
			project_id = ibm_project.project_instance.id
			name = "%s"
			locator_id = "%s"
		}

		data "ibm_project_config" "project_config_instance" {
			project_id = ibm_project_config.project_config_instance.project_id
			id = ibm_project_config.project_config_instance.project_config_id
		}
	`, projectConfigName, projectConfigLocatorID)
}

func testAccCheckIbmProjectConfigDataSourceConfig(projectConfigName string, projectConfigDescription string, projectConfigLocatorID string) string {
	return fmt.Sprintf(`
		resource "ibm_project" "project_instance" {
			resource_group = "Default"
			location = "us-south"
			name = "acme-microservice"
		}

		resource "ibm_project_config" "project_config_instance" {
			project_id = ibm_project.project_instance.id
			name = "%s"
			labels = [ "labels" ]
			description = "%s"
			authorizations {
                method = "API_KEY"
                api_key = "xxx"
            }
			locator_id = "%s"
			input {
				name = "name"
				type = "array"
				value = "anything as a string"
				required = true
			}
			setting {
				name = "name"
				value = "value"
			}
		}

		data "ibm_project_config" "project_config_instance" {
			project_id = ibm_project_config.project_config_instance.project_id
			id = ibm_project_config.project_config_instance.project_config_id
		}
	`, projectConfigName, projectConfigDescription, projectConfigLocatorID)
}
