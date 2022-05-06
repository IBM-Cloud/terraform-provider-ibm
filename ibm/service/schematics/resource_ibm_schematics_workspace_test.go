// Copyright IBM Corp. 2022 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package schematics_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM/schematics-go-sdk/schematicsv1"
)

func TestAccIBMSchematicsWorkspaceBasic(t *testing.T) {
	var conf schematicsv1.WorkspaceResponse

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMSchematicsWorkspaceDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMSchematicsWorkspaceConfigBasic(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMSchematicsWorkspaceExists("ibm_schematics_workspace.schematics_workspace", conf),
				),
			},
		},
	})
}

func TestAccIBMSchematicsWorkspaceAllArgs(t *testing.T) {
	var conf schematicsv1.WorkspaceResponse
	description := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	location := fmt.Sprintf("tf_location_%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	resourceGroup := fmt.Sprintf("tf_resource_group_%d", acctest.RandIntRange(10, 100))
	templateRef := fmt.Sprintf("tf_template_ref_%d", acctest.RandIntRange(10, 100))
	xGithubToken := fmt.Sprintf("tf_x_github_token_%d", acctest.RandIntRange(10, 100))
	descriptionUpdate := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	locationUpdate := fmt.Sprintf("tf_location_%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	resourceGroupUpdate := fmt.Sprintf("tf_resource_group_%d", acctest.RandIntRange(10, 100))
	templateRefUpdate := fmt.Sprintf("tf_template_ref_%d", acctest.RandIntRange(10, 100))
	xGithubTokenUpdate := fmt.Sprintf("tf_x_github_token_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMSchematicsWorkspaceDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMSchematicsWorkspaceConfig(description, location, name, resourceGroup, templateRef, xGithubToken),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMSchematicsWorkspaceExists("ibm_schematics_workspace.schematics_workspace", conf),
					resource.TestCheckResourceAttr("ibm_schematics_workspace.schematics_workspace", "description", description),
					resource.TestCheckResourceAttr("ibm_schematics_workspace.schematics_workspace", "location", location),
					resource.TestCheckResourceAttr("ibm_schematics_workspace.schematics_workspace", "name", name),
					resource.TestCheckResourceAttr("ibm_schematics_workspace.schematics_workspace", "resource_group", resourceGroup),
					resource.TestCheckResourceAttr("ibm_schematics_workspace.schematics_workspace", "template_ref", templateRef),
					resource.TestCheckResourceAttr("ibm_schematics_workspace.schematics_workspace", "x_github_token", xGithubToken),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMSchematicsWorkspaceConfig(descriptionUpdate, locationUpdate, nameUpdate, resourceGroupUpdate, templateRefUpdate, xGithubTokenUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_schematics_workspace.schematics_workspace", "description", descriptionUpdate),
					resource.TestCheckResourceAttr("ibm_schematics_workspace.schematics_workspace", "location", locationUpdate),
					resource.TestCheckResourceAttr("ibm_schematics_workspace.schematics_workspace", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_schematics_workspace.schematics_workspace", "resource_group", resourceGroupUpdate),
					resource.TestCheckResourceAttr("ibm_schematics_workspace.schematics_workspace", "template_ref", templateRefUpdate),
					resource.TestCheckResourceAttr("ibm_schematics_workspace.schematics_workspace", "x_github_token", xGithubTokenUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_schematics_workspace.schematics_workspace",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMSchematicsWorkspaceConfigBasic() string {
	return fmt.Sprintf(`

		resource "ibm_schematics_workspace" "schematics_workspace" {
		}
	`)
}

func testAccCheckIBMSchematicsWorkspaceConfig(description string, location string, name string, resourceGroup string, templateRef string, xGithubToken string) string {
	return fmt.Sprintf(`

		resource "ibm_schematics_workspace" "schematics_workspace" {
			// applied_shareddata_ids = "FIXME"
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
			tags = ["FIXME"]
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
			type = "terraform_v1.1"
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
	`, description, location, name, resourceGroup, templateRef, xGithubToken)
}

func testAccCheckIBMSchematicsWorkspaceExists(n string, obj schematicsv1.WorkspaceResponse) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		schematicsClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).SchematicsV1()
		if err != nil {
			return err
		}

		getWorkspaceOptions := &schematicsv1.GetWorkspaceOptions{}

		getWorkspaceOptions.SetWID(rs.Primary.ID)

		workspaceResponse, _, err := schematicsClient.GetWorkspace(getWorkspaceOptions)
		if err != nil {
			return err
		}

		obj = *workspaceResponse
		return nil
	}
}

func testAccCheckIBMSchematicsWorkspaceDestroy(s *terraform.State) error {
	schematicsClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).SchematicsV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_schematics_workspace" {
			continue
		}

		getWorkspaceOptions := &schematicsv1.GetWorkspaceOptions{}

		getWorkspaceOptions.SetWID(rs.Primary.ID)

		// Try to find the key
		_, response, err := schematicsClient.GetWorkspace(getWorkspaceOptions)

		if err == nil {
			return fmt.Errorf("schematics_workspace still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for schematics_workspace (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
