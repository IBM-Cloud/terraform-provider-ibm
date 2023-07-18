// Copyright IBM Corp. 2023 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package project_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM/project-go-sdk/projectv1"
)

func TestAccIbmProjectBasic(t *testing.T) {
	var conf projectv1.Project
	resourceGroup := fmt.Sprintf("tf_resource_group_%d", acctest.RandIntRange(10, 100))
	location := fmt.Sprintf("tf_location_%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	resourceGroupUpdate := fmt.Sprintf("tf_resource_group_%d", acctest.RandIntRange(10, 100))
	locationUpdate := fmt.Sprintf("tf_location_%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmProjectDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmProjectConfigBasic(resourceGroup, location, name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmProjectExists("ibm_project.project", conf),
					resource.TestCheckResourceAttr("ibm_project.project", "resource_group", resourceGroup),
					resource.TestCheckResourceAttr("ibm_project.project", "location", location),
					resource.TestCheckResourceAttr("ibm_project.project", "name", name),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmProjectConfigBasic(resourceGroupUpdate, locationUpdate, nameUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_project.project", "resource_group", resourceGroupUpdate),
					resource.TestCheckResourceAttr("ibm_project.project", "location", locationUpdate),
					resource.TestCheckResourceAttr("ibm_project.project", "name", nameUpdate),
				),
			},
		},
	})
}

func TestAccIbmProjectAllArgs(t *testing.T) {
	var conf projectv1.Project
	resourceGroup := fmt.Sprintf("tf_resource_group_%d", acctest.RandIntRange(10, 100))
	location := fmt.Sprintf("tf_location_%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	description := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	destroyOnDelete := "false"
	resourceGroupUpdate := fmt.Sprintf("tf_resource_group_%d", acctest.RandIntRange(10, 100))
	locationUpdate := fmt.Sprintf("tf_location_%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	descriptionUpdate := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	destroyOnDeleteUpdate := "true"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmProjectDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmProjectConfig(resourceGroup, location, name, description, destroyOnDelete),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmProjectExists("ibm_project.project", conf),
					resource.TestCheckResourceAttr("ibm_project.project", "resource_group", resourceGroup),
					resource.TestCheckResourceAttr("ibm_project.project", "location", location),
					resource.TestCheckResourceAttr("ibm_project.project", "name", name),
					resource.TestCheckResourceAttr("ibm_project.project", "description", description),
					resource.TestCheckResourceAttr("ibm_project.project", "destroy_on_delete", destroyOnDelete),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmProjectConfig(resourceGroupUpdate, locationUpdate, nameUpdate, descriptionUpdate, destroyOnDeleteUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_project.project", "resource_group", resourceGroupUpdate),
					resource.TestCheckResourceAttr("ibm_project.project", "location", locationUpdate),
					resource.TestCheckResourceAttr("ibm_project.project", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_project.project", "description", descriptionUpdate),
					resource.TestCheckResourceAttr("ibm_project.project", "destroy_on_delete", destroyOnDeleteUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_project.project",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmProjectConfigBasic(resourceGroup string, location string, name string) string {
	return fmt.Sprintf(`
		resource "ibm_project" "project_instance" {
			resource_group = "%s"
			location = "%s"
			name = "%s"
		}
	`, resourceGroup, location, name)
}

func testAccCheckIbmProjectConfig(resourceGroup string, location string, name string, description string, destroyOnDelete string) string {
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
				project_id = "project_id"
				version = 1
				is_draft = true
				needs_attention_state = [ "anything as a string" ]
				state = "deleted"
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
				job_summary {
					plan_summary = { "key" = "anything as a string" }
					apply_summary = { "key" = "anything as a string" }
					destroy_summary = { "key" = "anything as a string" }
					message_summary = { "key" = "anything as a string" }
					plan_messages = { "key" = "anything as a string" }
					apply_messages = { "key" = "anything as a string" }
					destroy_messages = { "key" = "anything as a string" }
				}
				cra_logs {
					cra_version = "cra_version"
					schema_version = "schema_version"
					status = "status"
					summary = { "key" = "anything as a string" }
					timestamp = "2021-01-31T09:44:12Z"
				}
				cost_estimate {
					version = "version"
					currency = "currency"
					total_hourly_cost = "total_hourly_cost"
					total_monthly_cost = "total_monthly_cost"
					past_total_hourly_cost = "past_total_hourly_cost"
					past_total_monthly_cost = "past_total_monthly_cost"
					diff_total_hourly_cost = "diff_total_hourly_cost"
					diff_total_monthly_cost = "diff_total_monthly_cost"
					time_generated = "2021-01-31T09:44:12Z"
					user_id = "user_id"
				}
				last_deployment_job_summary {
					plan_summary = { "key" = "anything as a string" }
					apply_summary = { "key" = "anything as a string" }
					destroy_summary = { "key" = "anything as a string" }
					message_summary = { "key" = "anything as a string" }
					plan_messages = { "key" = "anything as a string" }
					apply_messages = { "key" = "anything as a string" }
					destroy_messages = { "key" = "anything as a string" }
				}
			}
		}
	`, resourceGroup, location, name, description, destroyOnDelete)
}

func testAccCheckIbmProjectExists(n string, obj projectv1.Project) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		projectClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).ProjectV1()
		if err != nil {
			return err
		}

		getProjectOptions := &projectv1.GetProjectOptions{}

		getProjectOptions.SetID(rs.Primary.ID)

		project, _, err := projectClient.GetProject(getProjectOptions)
		if err != nil {
			return err
		}

		obj = *project
		return nil
	}
}

func testAccCheckIbmProjectDestroy(s *terraform.State) error {
	projectClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).ProjectV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_project" {
			continue
		}

		getProjectOptions := &projectv1.GetProjectOptions{}

		getProjectOptions.SetID(rs.Primary.ID)

		// Try to find the key
		_, response, err := projectClient.GetProject(getProjectOptions)

		if err == nil {
			return fmt.Errorf("project still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for project (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
