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
	projectResourceGroup := fmt.Sprintf("tf_resource_group_%d", acctest.RandIntRange(10, 100))
	projectLocation := fmt.Sprintf("tf_location_%d", acctest.RandIntRange(10, 100))
	projectName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmProjectDataSourceConfigBasic(projectResourceGroup, projectLocation, projectName),
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
	projectResourceGroup := fmt.Sprintf("tf_resource_group_%d", acctest.RandIntRange(10, 100))
	projectLocation := fmt.Sprintf("tf_location_%d", acctest.RandIntRange(10, 100))
	projectName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	projectDescription := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	projectDestroyOnDelete := "false"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmProjectDataSourceConfig(projectResourceGroup, projectLocation, projectName, projectDescription, projectDestroyOnDelete),
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
					resource.TestCheckResourceAttrSet("data.ibm_project.project", "configs.#"),
					resource.TestCheckResourceAttr("data.ibm_project.project", "configs.0.name", projectName),
					resource.TestCheckResourceAttr("data.ibm_project.project", "configs.0.description", projectDescription),
					resource.TestCheckResourceAttrSet("data.ibm_project.project", "configs.0.locator_id"),
					resource.TestCheckResourceAttrSet("data.ibm_project.project", "configs.0.type"),
					resource.TestCheckResourceAttrSet("data.ibm_project.project", "configs.0.id"),
				),
			},
		},
	})
}

func testAccCheckIbmProjectDataSourceConfigBasic(projectResourceGroup string, projectLocation string, projectName string) string {
	return fmt.Sprintf(`
		resource "ibm_project" "project_instance" {
			resource_group = "%s"
			location = "%s"
			name = "%s"
		}

		data "ibm_project" "project_instance" {
			id = ibm_project.project_instance.project_id
		}
	`, projectResourceGroup, projectLocation, projectName)
}

func testAccCheckIbmProjectDataSourceConfig(projectResourceGroup string, projectLocation string, projectName string, projectDescription string, projectDestroyOnDelete string) string {
	return fmt.Sprintf(`
		resource "ibm_project" "project_instance" {
			resource_group = "%s"
			location = "%s"
			name = "%s"
			description = "%s"
			destroy_on_delete = %s
			configs {
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
				type = "terraform_template"
				input {
					name = "name"
					type = "array"
					value = "anything as a string"
					required = true
				}
				output {
					name = "name"
					description = "description"
					value = "anything as a string"
				}
				setting {
					name = "name"
					value = "value"
				}
				id = "id"
			}
		}

		data "ibm_project" "project_instance" {
			id = ibm_project.project_instance.project_id
		}
	`, projectResourceGroup, projectLocation, projectName, projectDescription, projectDestroyOnDelete)
}
