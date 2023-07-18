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
	projectConfigLocatorID := fmt.Sprintf("tf_locator_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmProjectConfigDataSourceConfigBasic(projectConfigName, projectConfigLocatorID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config", "project_id"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config", "locator_id"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config", "type"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config", "project_config_id"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config", "version"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config", "is_draft"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config", "state"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config", "update_available"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config", "updated_at"),
				),
			},
		},
	})
}

func TestAccIbmProjectConfigDataSourceAllArgs(t *testing.T) {
	projectConfigName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	projectConfigDescription := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	projectConfigLocatorID := fmt.Sprintf("tf_locator_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmProjectConfigDataSourceConfig(projectConfigName, projectConfigDescription, projectConfigLocatorID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config", "project_id"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config", "labels.#"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config", "description"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config", "authorizations.#"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config", "compliance_profile.#"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config", "locator_id"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config", "type"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config", "input.#"),
					resource.TestCheckResourceAttr("data.ibm_project_config.project_config", "input.0.name", projectConfigName),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config", "input.0.type"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config", "input.0.value"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config", "input.0.required"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config", "output.#"),
					resource.TestCheckResourceAttr("data.ibm_project_config.project_config", "output.0.name", projectConfigName),
					resource.TestCheckResourceAttr("data.ibm_project_config.project_config", "output.0.description", projectConfigDescription),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config", "output.0.value"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config", "setting.#"),
					resource.TestCheckResourceAttr("data.ibm_project_config.project_config", "setting.0.name", projectConfigName),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config", "setting.0.value"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config", "project_config_id"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config", "version"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config", "is_draft"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config", "needs_attention_state"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config", "needs_attention_state.#"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config", "state"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config", "pipeline_state"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config", "update_available"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config", "last_approved.#"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config", "last_save"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config", "job_summary.#"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config", "cra_logs.#"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config", "cost_estimate.#"),
					resource.TestCheckResourceAttrSet("data.ibm_project_config.project_config", "last_deployment_job_summary.#"),
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
			project_id = ibm_project_config.project_config.project_id
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
			labels = "FIXME"
			description = "%s"
			authorizations {
				trusted_profile {
					id = "id"
					target_iam_id = "target_iam_id"
				}
				method = "method"
				api_key = "api_key"
			}
			compliance_profile {
				id = "id"
				instance_id = "instance_id"
				instance_location = "instance_location"
				attachment_id = "attachment_id"
				profile_name = "profile_name"
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
			project_id = ibm_project_config.project_config.project_id
			id = ibm_project_config.project_config_instance.project_config_id
		}
	`, projectConfigName, projectConfigDescription, projectConfigLocatorID)
}
