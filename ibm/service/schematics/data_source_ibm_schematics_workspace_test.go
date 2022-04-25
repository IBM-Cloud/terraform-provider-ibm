// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package schematics_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIBMSchematicsWorkspaceDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMSchematicsWorkspaceDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "workspace_id"),
				),
			},
		},
	})
}

func TestAccIBMSchematicsWorkspaceDataSourceAllArgs(t *testing.T) {
	workspaceResponseDescription := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	workspaceResponseLocation := fmt.Sprintf("tf_location_%d", acctest.RandIntRange(10, 100))
	workspaceResponseName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	workspaceResponseResourceGroup := fmt.Sprintf("tf_resource_group_%d", acctest.RandIntRange(10, 100))
	workspaceResponseTemplateRef := fmt.Sprintf("tf_template_ref_%d", acctest.RandIntRange(10, 100))
	workspaceResponseXGithubToken := fmt.Sprintf("tf_x_github_token_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMSchematicsWorkspaceDataSourceConfig(workspaceResponseDescription, workspaceResponseLocation, workspaceResponseName, workspaceResponseResourceGroup, workspaceResponseTemplateRef, workspaceResponseXGithubToken),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "workspace_id"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "catalog_ref.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "created_by"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "dependencies.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "description"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "last_health_check_at"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "location"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "resource_group"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "runtime_data.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "runtime_data.0.engine_cmd"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "runtime_data.0.engine_name"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "runtime_data.0.engine_version"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "runtime_data.0.id"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "runtime_data.0.log_store_url"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "runtime_data.0.state_store_url"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "shared_data.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "status"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "tags.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "template_data.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "template_data.0.folder"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "template_data.0.compact"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "template_data.0.has_githubtoken"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "template_data.0.id"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "template_data.0.type"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "template_data.0.uninstall_script_name"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "template_data.0.values"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "template_data.0.values_url"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "template_ref"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "template_repo.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "type.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "updated_by"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "cart_id"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "last_action_name"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "last_activity_id"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "last_job.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "workspace_status.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_workspace.schematics_workspace", "workspace_status_msg.#"),
				),
			},
		},
	})
}

func testAccCheckIBMSchematicsWorkspaceDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		resource "ibm_schematics_workspace" "schematics_workspace" {
		}

		data "ibm_schematics_workspace" "schematics_workspace" {
			w_id = "w_id"
		}
	`)
}

func testAccCheckIBMSchematicsWorkspaceDataSourceConfig(workspaceResponseDescription string, workspaceResponseLocation string, workspaceResponseName string, workspaceResponseResourceGroup string, workspaceResponseTemplateRef string, workspaceResponseXGithubToken string) string {
	return fmt.Sprintf(`
		resource "ibm_schematics_workspace" "schematics_workspace" {
			applied_shareddata_ids = "FIXME"
			catalog_ref {
				dry_run = true
				owning_account = "owning_account"
				item_icon_url = "item_icon_url"
				item_id = "item_id"
				item_name = "item_name"
				item_readme_url = "item_readme_url"
				item_url = "item_url"
				launch_url = "launch_url"
				offering_version = "offering_version"
			}
			dependencies {
				parents = [ "parents" ]
				children = [ "children" ]
			}
			description = "%s"
			location = "%s"
			name = "%s"
			resource_group = "%s"
			shared_data {
				cluster_created_on = "cluster_created_on"
				cluster_id = "cluster_id"
				cluster_name = "cluster_name"
				cluster_type = "cluster_type"
				entitlement_keys = [ null ]
				namespace = "namespace"
				region = "region"
				resource_group_id = "resource_group_id"
				worker_count = 1
				worker_machine_type = "worker_machine_type"
			}
			tags = "FIXME"
			template_data {
				env_values = [ null ]
				env_values_metadata {
					hidden = true
					name = "name"
					secure = true
				}
				folder = "folder"
				compact = true
				init_state_file = "init_state_file"
				injectors {
					tft_git_url = "tft_git_url"
					tft_git_token = "tft_git_token"
					tft_prefix = "tft_prefix"
					injection_type = "injection_type"
					tft_name = "tft_name"
					tft_parameters {
						name = "name"
						value = "value"
					}
				}
				type = "type"
				uninstall_script_name = "uninstall_script_name"
				values = "values"
				values_metadata = [ null ]
				variablestore {
					description = "description"
					name = "name"
					secure = true
					type = "type"
					use_default = true
					value = "value"
				}
			}
			template_ref = "%s"
			template_repo {
				branch = "branch"
				release = "release"
				repo_sha_value = "repo_sha_value"
				repo_url = "repo_url"
				url = "url"
			}
			type = "FIXME"
			workspace_status {
				frozen = true
				frozen_at = "2021-01-31T09:44:12Z"
				frozen_by = "frozen_by"
				locked = true
				locked_by = "locked_by"
				locked_time = "2021-01-31T09:44:12Z"
			}
			x_github_token = "%s"
		}

		data "ibm_schematics_workspace" "schematics_workspace" {
			w_id = "w_id"
		}
	`, workspaceResponseDescription, workspaceResponseLocation, workspaceResponseName, workspaceResponseResourceGroup, workspaceResponseTemplateRef, workspaceResponseXGithubToken)
}
