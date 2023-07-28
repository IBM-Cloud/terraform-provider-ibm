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

func TestAccIbmProjectDataSourceBasic(t *testing.T) {
	projectCanonicalResourceGroup := fmt.Sprintf("tf_resource_group_%d", acctest.RandIntRange(10, 100))
	projectCanonicalLocation := fmt.Sprintf("tf_location_%d", acctest.RandIntRange(10, 100))
	projectCanonicalName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmProjectDataSourceConfigBasic(projectCanonicalResourceGroup, projectCanonicalLocation, projectCanonicalName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_project.project", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_project.project", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_project.project", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_project.project", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_project.project", "location"),
					resource.TestCheckResourceAttrSet("data.ibm_project.project", "resource_group"),
					resource.TestCheckResourceAttrSet("data.ibm_project.project", "state"),
					resource.TestCheckResourceAttrSet("data.ibm_project.project", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_project.project", "description"),
					resource.TestCheckResourceAttrSet("data.ibm_project.project", "destroy_on_delete"),
				),
			},
		},
	})
}

func TestAccIbmProjectDataSourceAllArgs(t *testing.T) {
	projectCanonicalResourceGroup := fmt.Sprintf("tf_resource_group_%d", acctest.RandIntRange(10, 100))
	projectCanonicalLocation := fmt.Sprintf("tf_location_%d", acctest.RandIntRange(10, 100))
	projectCanonicalName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	projectCanonicalDescription := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	projectCanonicalDestroyOnDelete := "false"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmProjectDataSourceConfig(projectCanonicalResourceGroup, projectCanonicalLocation, projectCanonicalName, projectCanonicalDescription, projectCanonicalDestroyOnDelete),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_project.project", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_project.project", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_project.project", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_project.project", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_project.project", "cumulative_needs_attention_view.#"),
					resource.TestCheckResourceAttrSet("data.ibm_project.project", "cumulative_needs_attention_view.0.event"),
					resource.TestCheckResourceAttrSet("data.ibm_project.project", "cumulative_needs_attention_view.0.event_id"),
					resource.TestCheckResourceAttrSet("data.ibm_project.project", "cumulative_needs_attention_view.0.config_id"),
					resource.TestCheckResourceAttrSet("data.ibm_project.project", "cumulative_needs_attention_view.0.config_version"),
					resource.TestCheckResourceAttrSet("data.ibm_project.project", "cumulative_needs_attention_view_error"),
					resource.TestCheckResourceAttrSet("data.ibm_project.project", "location"),
					resource.TestCheckResourceAttrSet("data.ibm_project.project", "resource_group"),
					resource.TestCheckResourceAttrSet("data.ibm_project.project", "state"),
					resource.TestCheckResourceAttrSet("data.ibm_project.project", "event_notifications_crn"),
					resource.TestCheckResourceAttrSet("data.ibm_project.project", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_project.project", "description"),
					resource.TestCheckResourceAttrSet("data.ibm_project.project", "destroy_on_delete"),
					resource.TestCheckResourceAttrSet("data.ibm_project.project", "definition.#"),
					resource.TestCheckResourceAttrSet("data.ibm_project.project", "configs.#"),
					resource.TestCheckResourceAttrSet("data.ibm_project.project", "configs.0.id"),
					resource.TestCheckResourceAttrSet("data.ibm_project.project", "configs.0.project_id"),
					resource.TestCheckResourceAttrSet("data.ibm_project.project", "configs.0.version"),
					resource.TestCheckResourceAttrSet("data.ibm_project.project", "configs.0.is_draft"),
					resource.TestCheckResourceAttrSet("data.ibm_project.project", "configs.0.needs_attention_state"),
					resource.TestCheckResourceAttrSet("data.ibm_project.project", "configs.0.state"),
					resource.TestCheckResourceAttrSet("data.ibm_project.project", "configs.0.pipeline_state"),
					resource.TestCheckResourceAttrSet("data.ibm_project.project", "configs.0.update_available"),
					resource.TestCheckResourceAttrSet("data.ibm_project.project", "configs.0.created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_project.project", "configs.0.updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_project.project", "configs.0.last_save"),
					resource.TestCheckResourceAttr("data.ibm_project.project", "configs.0.name", projectCanonicalName),
					resource.TestCheckResourceAttr("data.ibm_project.project", "configs.0.description", projectCanonicalDescription),
					resource.TestCheckResourceAttrSet("data.ibm_project.project", "configs.0.locator_id"),
					resource.TestCheckResourceAttrSet("data.ibm_project.project", "configs.0.type"),
					resource.TestCheckResourceAttrSet("data.ibm_project.project", "configs.0.href"),
				),
			},
		},
	})
}

func testAccCheckIbmProjectDataSourceConfigBasic(projectCanonicalResourceGroup string, projectCanonicalLocation string, projectCanonicalName string) string {
	return fmt.Sprintf(`
		resource "ibm_project" "project_instance" {
			resource_group = "%s"
			location = "%s"
			name = "%s"
		}

		data "ibm_project" "project_instance" {
			id = ibm_project.project_instance.projectCanonical_id
		}
	`, projectCanonicalResourceGroup, projectCanonicalLocation, projectCanonicalName)
}

func testAccCheckIbmProjectDataSourceConfig(projectCanonicalResourceGroup string, projectCanonicalLocation string, projectCanonicalName string, projectCanonicalDescription string, projectCanonicalDestroyOnDelete string) string {
	return fmt.Sprintf(`
		resource "ibm_project" "project_instance" {
			resource_group = "%s"
			location = "%s"
			name = "%s"
			description = "%s"
			destroy_on_delete = %s
			configs {
				id = "id"
				project_id = "project_id"
				version = 1
				is_draft = true
				needs_attention_state = [ "anything as a string" ]
				state = "approved"
				pipeline_state = "pipeline_failed"
				update_available = true
				created_at = "2021-01-31T09:44:12Z"
				updated_at = "2021-01-31T09:44:12Z"
				last_approved {
					is_forced = true
					comment = "comment"
					timestamp = "2021-01-31T09:44:12Z"
					user_id = "user_id"
				}
				last_save = "2021-01-31T09:44:12Z"
				name = "name"
				labels = [ "labels" ]
				description = "description"
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
				locator_id = "locator_id"
				input {
					name = "name"
					value = "anything as a string"
				}
				setting {
					name = "name"
					value = "value"
				}
				type = "terraform_template"
				output {
					name = "name"
					description = "description"
					value = "anything as a string"
				}
				active_draft {
					version = 1
					state = "discarded"
					pipeline_state = "pipeline_failed"
					href = "href"
				}
				definition {
					name = "name"
					labels = [ "labels" ]
					description = "description"
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
					locator_id = "locator_id"
					input {
						name = "name"
						value = "anything as a string"
					}
					setting {
						name = "name"
						value = "value"
					}
					type = "terraform_template"
					output {
						name = "name"
						description = "description"
						value = "anything as a string"
					}
				}
				href = "href"
			}
		}

		data "ibm_project" "project_instance" {
			id = ibm_project.project_instance.projectCanonical_id
		}
	`, projectCanonicalResourceGroup, projectCanonicalLocation, projectCanonicalName, projectCanonicalDescription, projectCanonicalDestroyOnDelete)
}
