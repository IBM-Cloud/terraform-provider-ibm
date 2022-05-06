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

func TestAccIBMSchematicsActionBasic(t *testing.T) {
	var conf schematicsv1.Action

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMSchematicsActionDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMSchematicsActionConfigBasic(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMSchematicsActionExists("ibm_schematics_action.schematics_action", conf),
				),
			},
		},
	})
}

func TestAccIBMSchematicsActionAllArgs(t *testing.T) {
	var conf schematicsv1.Action
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	description := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	location := "us-south"
	resourceGroup := fmt.Sprintf("tf_resource_group_%d", acctest.RandIntRange(10, 100))
	bastionConnectionType := "ssh"
	inventoryConnectionType := "ssh"
	sourceReadmeURL := fmt.Sprintf("tf_source_readme_url_%d", acctest.RandIntRange(10, 100))
	sourceType := "local"
	commandParameter := fmt.Sprintf("tf_command_parameter_%d", acctest.RandIntRange(10, 100))
	inventory := fmt.Sprintf("tf_inventory_%d", acctest.RandIntRange(10, 100))
	targetsIni := fmt.Sprintf("tf_targets_ini_%d", acctest.RandIntRange(10, 100))
	xGithubToken := fmt.Sprintf("tf_x_github_token_%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	descriptionUpdate := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	locationUpdate := "eu-de"
	resourceGroupUpdate := fmt.Sprintf("tf_resource_group_%d", acctest.RandIntRange(10, 100))
	bastionConnectionTypeUpdate := "ssh"
	inventoryConnectionTypeUpdate := "winrm"
	sourceReadmeURLUpdate := fmt.Sprintf("tf_source_readme_url_%d", acctest.RandIntRange(10, 100))
	sourceTypeUpdate := "external_scm"
	commandParameterUpdate := fmt.Sprintf("tf_command_parameter_%d", acctest.RandIntRange(10, 100))
	inventoryUpdate := fmt.Sprintf("tf_inventory_%d", acctest.RandIntRange(10, 100))
	targetsIniUpdate := fmt.Sprintf("tf_targets_ini_%d", acctest.RandIntRange(10, 100))
	xGithubTokenUpdate := fmt.Sprintf("tf_x_github_token_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMSchematicsActionDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMSchematicsActionConfig(name, description, location, resourceGroup, bastionConnectionType, inventoryConnectionType, sourceReadmeURL, sourceType, commandParameter, inventory, targetsIni, xGithubToken),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMSchematicsActionExists("ibm_schematics_action.schematics_action", conf),
					resource.TestCheckResourceAttr("ibm_schematics_action.schematics_action", "name", name),
					resource.TestCheckResourceAttr("ibm_schematics_action.schematics_action", "description", description),
					resource.TestCheckResourceAttr("ibm_schematics_action.schematics_action", "location", location),
					resource.TestCheckResourceAttr("ibm_schematics_action.schematics_action", "resource_group", resourceGroup),
					resource.TestCheckResourceAttr("ibm_schematics_action.schematics_action", "bastion_connection_type", bastionConnectionType),
					resource.TestCheckResourceAttr("ibm_schematics_action.schematics_action", "inventory_connection_type", inventoryConnectionType),
					resource.TestCheckResourceAttr("ibm_schematics_action.schematics_action", "source_readme_url", sourceReadmeURL),
					resource.TestCheckResourceAttr("ibm_schematics_action.schematics_action", "source_type", sourceType),
					resource.TestCheckResourceAttr("ibm_schematics_action.schematics_action", "command_parameter", commandParameter),
					resource.TestCheckResourceAttr("ibm_schematics_action.schematics_action", "inventory", inventory),
					resource.TestCheckResourceAttr("ibm_schematics_action.schematics_action", "targets_ini", targetsIni),
					resource.TestCheckResourceAttr("ibm_schematics_action.schematics_action", "x_github_token", xGithubToken),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMSchematicsActionConfig(nameUpdate, descriptionUpdate, locationUpdate, resourceGroupUpdate, bastionConnectionTypeUpdate, inventoryConnectionTypeUpdate, sourceReadmeURLUpdate, sourceTypeUpdate, commandParameterUpdate, inventoryUpdate, targetsIniUpdate, xGithubTokenUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_schematics_action.schematics_action", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_schematics_action.schematics_action", "description", descriptionUpdate),
					resource.TestCheckResourceAttr("ibm_schematics_action.schematics_action", "location", locationUpdate),
					resource.TestCheckResourceAttr("ibm_schematics_action.schematics_action", "resource_group", resourceGroupUpdate),
					resource.TestCheckResourceAttr("ibm_schematics_action.schematics_action", "bastion_connection_type", bastionConnectionTypeUpdate),
					resource.TestCheckResourceAttr("ibm_schematics_action.schematics_action", "inventory_connection_type", inventoryConnectionTypeUpdate),
					resource.TestCheckResourceAttr("ibm_schematics_action.schematics_action", "source_readme_url", sourceReadmeURLUpdate),
					resource.TestCheckResourceAttr("ibm_schematics_action.schematics_action", "source_type", sourceTypeUpdate),
					resource.TestCheckResourceAttr("ibm_schematics_action.schematics_action", "command_parameter", commandParameterUpdate),
					resource.TestCheckResourceAttr("ibm_schematics_action.schematics_action", "inventory", inventoryUpdate),
					resource.TestCheckResourceAttr("ibm_schematics_action.schematics_action", "targets_ini", targetsIniUpdate),
					resource.TestCheckResourceAttr("ibm_schematics_action.schematics_action", "x_github_token", xGithubTokenUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_schematics_action.schematics_action",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMSchematicsActionConfigBasic() string {
	return fmt.Sprintf(`

		resource "ibm_schematics_action" "schematics_action" {
		}
	`)
}

func testAccCheckIBMSchematicsActionConfig(name string, description string, location string, resourceGroup string, bastionConnectionType string, inventoryConnectionType string, sourceReadmeURL string, sourceType string, commandParameter string, inventory string, targetsIni string, xGithubToken string) string {
	return fmt.Sprintf(`

		resource "ibm_schematics_action" "schematics_action" {
			name = "%s"
			description = "%s"
			location = "%s"
			resource_group = "%s"
			bastion_connection_type = "%s"
			inventory_connection_type = "%s"
			tags = ["FIXME"]
			user_state {
				state = "draft"
				set_by = "set_by"
				set_at = "2021-01-31T09:44:12Z"
			}
			source_readme_url = "%s"
			source {
				source_type = "local"
				git {
					computed_git_repo_url = "computed_git_repo_url"
					git_repo_url = "git_repo_url"
					git_token = "git_token"
					git_repo_folder = "git_repo_folder"
					git_release = "git_release"
					git_branch = "git_branch"
				}
				catalog {
					catalog_name = "catalog_name"
					offering_name = "offering_name"
					offering_version = "offering_version"
					offering_kind = "offering_kind"
					offering_id = "offering_id"
					offering_version_id = "offering_version_id"
					offering_repo_url = "offering_repo_url"
				}
			}
			source_type = "%s"
			command_parameter = "%s"
			inventory = "%s"
			credentials {
				name = "name"
				value = "value"
				use_default = true
				metadata {
					type = "string"
					aliases = [ "aliases" ]
					description = "description"
					cloud_data_type = "cloud_data_type"
					default_value = "default_value"
					link_status = "normal"
					immutable = true
					hidden = true
					required = true
					position = 1
					group_by = "group_by"
					source = "source"
				}
				link = "link"
			}
			bastion {
				name = "name"
				host = "host"
			}
			bastion_credential {
				name = "name"
				value = "value"
				use_default = true
				metadata {
					type = "string"
					aliases = [ "aliases" ]
					description = "description"
					cloud_data_type = "cloud_data_type"
					default_value = "default_value"
					link_status = "normal"
					immutable = true
					hidden = true
					required = true
					position = 1
					group_by = "group_by"
					source = "source"
				}
				link = "link"
			}
			targets_ini = "%s"
			action_inputs {
				name = "name"
				value = "value"
				use_default = true
				metadata {
					type = "boolean"
					aliases = [ "aliases" ]
					description = "description"
					cloud_data_type = "cloud_data_type"
					default_value = "default_value"
					link_status = "normal"
					secure = true
					immutable = true
					hidden = true
					required = true
					options = [ "options" ]
					min_value = 1
					max_value = 1
					min_length = 1
					max_length = 1
					matches = "matches"
					position = 1
					group_by = "group_by"
					source = "source"
				}
				link = "link"
			}
			action_outputs {
				name = "name"
				value = "value"
				use_default = true
				metadata {
					type = "boolean"
					aliases = [ "aliases" ]
					description = "description"
					cloud_data_type = "cloud_data_type"
					default_value = "default_value"
					link_status = "normal"
					secure = true
					immutable = true
					hidden = true
					required = true
					options = [ "options" ]
					min_value = 1
					max_value = 1
					min_length = 1
					max_length = 1
					matches = "matches"
					position = 1
					group_by = "group_by"
					source = "source"
				}
				link = "link"
			}
			settings {
				name = "name"
				value = "value"
				use_default = true
				metadata {
					type = "boolean"
					aliases = [ "aliases" ]
					description = "description"
					cloud_data_type = "cloud_data_type"
					default_value = "default_value"
					link_status = "normal"
					secure = true
					immutable = true
					hidden = true
					required = true
					options = [ "options" ]
					min_value = 1
					max_value = 1
					min_length = 1
					max_length = 1
					matches = "matches"
					position = 1
					group_by = "group_by"
					source = "source"
				}
				link = "link"
			}
			state {
				status_code = "normal"
				status_job_id = "status_job_id"
				status_message = "status_message"
			}
			sys_lock {
				sys_locked = true
				sys_locked_by = "sys_locked_by"
				sys_locked_at = "2021-01-31T09:44:12Z"
			}
			x_github_token = "%s"
		}
	`, name, description, location, resourceGroup, bastionConnectionType, inventoryConnectionType, sourceReadmeURL, sourceType, commandParameter, inventory, targetsIni, xGithubToken)
}

func testAccCheckIBMSchematicsActionExists(n string, obj schematicsv1.Action) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		schematicsClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).SchematicsV1()
		if err != nil {
			return err
		}

		getActionOptions := &schematicsv1.GetActionOptions{}

		getActionOptions.SetActionID(rs.Primary.ID)

		action, _, err := schematicsClient.GetAction(getActionOptions)
		if err != nil {
			return err
		}

		obj = *action
		return nil
	}
}

func testAccCheckIBMSchematicsActionDestroy(s *terraform.State) error {
	schematicsClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).SchematicsV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_schematics_action" {
			continue
		}

		getActionOptions := &schematicsv1.GetActionOptions{}

		getActionOptions.SetActionID(rs.Primary.ID)

		// Try to find the key
		_, response, err := schematicsClient.GetAction(getActionOptions)

		if err == nil {
			return fmt.Errorf("schematics_action still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for schematics_action (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
