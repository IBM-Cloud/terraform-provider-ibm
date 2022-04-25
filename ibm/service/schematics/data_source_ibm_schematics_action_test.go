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

func TestAccIBMSchematicsActionDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMSchematicsActionDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "action_id"),
				),
			},
		},
	})
}

func TestAccIBMSchematicsActionDataSourceAllArgs(t *testing.T) {
	actionName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	actionDescription := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	actionLocation := "us-south"
	actionResourceGroup := fmt.Sprintf("tf_resource_group_%d", acctest.RandIntRange(10, 100))
	actionBastionConnectionType := "ssh"
	actionInventoryConnectionType := "ssh"
	actionSourceReadmeURL := fmt.Sprintf("tf_source_readme_url_%d", acctest.RandIntRange(10, 100))
	actionSourceType := "local"
	actionCommandParameter := fmt.Sprintf("tf_command_parameter_%d", acctest.RandIntRange(10, 100))
	actionInventory := fmt.Sprintf("tf_inventory_%d", acctest.RandIntRange(10, 100))
	actionTargetsIni := fmt.Sprintf("tf_targets_ini_%d", acctest.RandIntRange(10, 100))
	actionXGithubToken := fmt.Sprintf("tf_x_github_token_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMSchematicsActionDataSourceConfig(actionName, actionDescription, actionLocation, actionResourceGroup, actionBastionConnectionType, actionInventoryConnectionType, actionSourceReadmeURL, actionSourceType, actionCommandParameter, actionInventory, actionTargetsIni, actionXGithubToken),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "action_id"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "description"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "location"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "resource_group"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "bastion_connection_type"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "inventory_connection_type"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "tags.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "user_state.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "source_readme_url"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "source.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "source_type"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "command_parameter"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "inventory"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "credentials.#"),
					resource.TestCheckResourceAttr("data.ibm_schematics_action.schematics_action", "credentials.0.name", actionName),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "credentials.0.value"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "credentials.0.use_default"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "credentials.0.link"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "bastion.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "bastion_credential.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "targets_ini"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "action_inputs.#"),
					resource.TestCheckResourceAttr("data.ibm_schematics_action.schematics_action", "action_inputs.0.name", actionName),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "action_inputs.0.value"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "action_inputs.0.use_default"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "action_inputs.0.link"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "action_outputs.#"),
					resource.TestCheckResourceAttr("data.ibm_schematics_action.schematics_action", "action_outputs.0.name", actionName),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "action_outputs.0.value"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "action_outputs.0.use_default"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "action_outputs.0.link"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "settings.#"),
					resource.TestCheckResourceAttr("data.ibm_schematics_action.schematics_action", "settings.0.name", actionName),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "settings.0.value"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "settings.0.use_default"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "settings.0.link"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "crn"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "account"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "source_created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "source_created_by"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "source_updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "source_updated_by"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "created_at"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "created_by"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "updated_at"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "updated_by"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "state.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "playbook_names.#"),
					resource.TestCheckResourceAttrSet("data.ibm_schematics_action.schematics_action", "sys_lock.#"),
				),
			},
		},
	})
}

func testAccCheckIBMSchematicsActionDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		resource "ibm_schematics_action" "schematics_action" {
		}

		data "ibm_schematics_action" "schematics_action" {
			action_id = "action_id"
		}
	`)
}

func testAccCheckIBMSchematicsActionDataSourceConfig(actionName string, actionDescription string, actionLocation string, actionResourceGroup string, actionBastionConnectionType string, actionInventoryConnectionType string, actionSourceReadmeURL string, actionSourceType string, actionCommandParameter string, actionInventory string, actionTargetsIni string, actionXGithubToken string) string {
	return fmt.Sprintf(`
		resource "ibm_schematics_action" "schematics_action" {
			name = "%s"
			description = "%s"
			location = "%s"
			resource_group = "%s"
			bastion_connection_type = "%s"
			inventory_connection_type = "%s"
			tags = "FIXME"
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

		data "ibm_schematics_action" "schematics_action" {
			action_id = "action_id"
		}
	`, actionName, actionDescription, actionLocation, actionResourceGroup, actionBastionConnectionType, actionInventoryConnectionType, actionSourceReadmeURL, actionSourceType, actionCommandParameter, actionInventory, actionTargetsIni, actionXGithubToken)
}
