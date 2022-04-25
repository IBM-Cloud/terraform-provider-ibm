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

func TestAccIBMSchematicsJobBasic(t *testing.T) {
	var conf schematicsv1.Job
	refreshToken := fmt.Sprintf("tf_refresh_token_%d", acctest.RandIntRange(10, 100))
	refreshTokenUpdate := fmt.Sprintf("tf_refresh_token_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMSchematicsJobDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMSchematicsJobConfigBasic(refreshToken),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMSchematicsJobExists("ibm_schematics_job.schematics_job", conf),
					resource.TestCheckResourceAttr("ibm_schematics_job.schematics_job", "refresh_token", refreshToken),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMSchematicsJobConfigBasic(refreshTokenUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_schematics_job.schematics_job", "refresh_token", refreshTokenUpdate),
				),
			},
		},
	})
}

func TestAccIBMSchematicsJobAllArgs(t *testing.T) {
	var conf schematicsv1.Job
	refreshToken := fmt.Sprintf("tf_refresh_token_%d", acctest.RandIntRange(10, 100))
	commandObject := "workspace"
	commandObjectID := fmt.Sprintf("tf_command_object_id_%d", acctest.RandIntRange(10, 100))
	commandName := "workspace_plan"
	commandParameter := fmt.Sprintf("tf_command_parameter_%d", acctest.RandIntRange(10, 100))
	location := "us-south"
	refreshTokenUpdate := fmt.Sprintf("tf_refresh_token_%d", acctest.RandIntRange(10, 100))
	commandObjectUpdate := "environment"
	commandObjectIDUpdate := fmt.Sprintf("tf_command_object_id_%d", acctest.RandIntRange(10, 100))
	commandNameUpdate := "terraform_commands"
	commandParameterUpdate := fmt.Sprintf("tf_command_parameter_%d", acctest.RandIntRange(10, 100))
	locationUpdate := "eu-de"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIBMSchematicsJobDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIBMSchematicsJobConfig(refreshToken, commandObject, commandObjectID, commandName, commandParameter, location),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIBMSchematicsJobExists("ibm_schematics_job.schematics_job", conf),
					resource.TestCheckResourceAttr("ibm_schematics_job.schematics_job", "refresh_token", refreshToken),
					resource.TestCheckResourceAttr("ibm_schematics_job.schematics_job", "command_object", commandObject),
					resource.TestCheckResourceAttr("ibm_schematics_job.schematics_job", "command_object_id", commandObjectID),
					resource.TestCheckResourceAttr("ibm_schematics_job.schematics_job", "command_name", commandName),
					resource.TestCheckResourceAttr("ibm_schematics_job.schematics_job", "command_parameter", commandParameter),
					resource.TestCheckResourceAttr("ibm_schematics_job.schematics_job", "location", location),
				),
			},
			resource.TestStep{
				Config: testAccCheckIBMSchematicsJobConfig(refreshTokenUpdate, commandObjectUpdate, commandObjectIDUpdate, commandNameUpdate, commandParameterUpdate, locationUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_schematics_job.schematics_job", "refresh_token", refreshTokenUpdate),
					resource.TestCheckResourceAttr("ibm_schematics_job.schematics_job", "command_object", commandObjectUpdate),
					resource.TestCheckResourceAttr("ibm_schematics_job.schematics_job", "command_object_id", commandObjectIDUpdate),
					resource.TestCheckResourceAttr("ibm_schematics_job.schematics_job", "command_name", commandNameUpdate),
					resource.TestCheckResourceAttr("ibm_schematics_job.schematics_job", "command_parameter", commandParameterUpdate),
					resource.TestCheckResourceAttr("ibm_schematics_job.schematics_job", "location", locationUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_schematics_job.schematics_job",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIBMSchematicsJobConfigBasic(refreshToken string) string {
	return fmt.Sprintf(`

		resource "ibm_schematics_job" "schematics_job" {
			refresh_token = "%s"
		}
	`, refreshToken)
}

func testAccCheckIBMSchematicsJobConfig(refreshToken string, commandObject string, commandObjectID string, commandName string, commandParameter string, location string) string {
	return fmt.Sprintf(`

		resource "ibm_schematics_job" "schematics_job" {
			refresh_token = "%s"
			command_object = "%s"
			command_object_id = "%s"
			command_name = "%s"
			command_parameter = "%s"
			// command_options = ["FIXME"]
			job_inputs {
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
			job_env_settings {
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
			tags = ["FIXME"]
			location = "%s"
			status {
				position_in_queue = 1.0
				total_in_queue = 1.0
				workspace_job_status {
					workspace_name = "workspace_name"
					status_code = "job_pending"
					status_message = "status_message"
					flow_status {
						flow_id = "flow_id"
						flow_name = "flow_name"
						status_code = "job_pending"
						status_message = "status_message"
						workitems {
							workspace_id = "workspace_id"
							workspace_name = "workspace_name"
							job_id = "job_id"
							status_code = "job_pending"
							status_message = "status_message"
							updated_at = "2021-01-31T09:44:12Z"
						}
						updated_at = "2021-01-31T09:44:12Z"
					}
					template_status {
						template_id = "template_id"
						template_name = "template_name"
						flow_index = 1
						status_code = "job_pending"
						status_message = "status_message"
						updated_at = "2021-01-31T09:44:12Z"
					}
					updated_at = "2021-01-31T09:44:12Z"
					commands {
						name = "name"
						outcome = "outcome"
					}
				}
				action_job_status {
					action_name = "action_name"
					status_code = "job_pending"
					status_message = "status_message"
					bastion_status_code = "none"
					bastion_status_message = "bastion_status_message"
					targets_status_code = "none"
					targets_status_message = "targets_status_message"
					updated_at = "2021-01-31T09:44:12Z"
				}
				system_job_status {
					system_status_message = "system_status_message"
					system_status_code = "job_pending"
					schematics_resource_status {
						status_code = "job_pending"
						status_message = "status_message"
						schematics_resource_id = "schematics_resource_id"
						updated_at = "2021-01-31T09:44:12Z"
					}
					updated_at = "2021-01-31T09:44:12Z"
				}
				flow_job_status {
					flow_id = "flow_id"
					flow_name = "flow_name"
					status_code = "job_pending"
					status_message = "status_message"
					workitems {
						workspace_id = "workspace_id"
						workspace_name = "workspace_name"
						job_id = "job_id"
						status_code = "job_pending"
						status_message = "status_message"
						updated_at = "2021-01-31T09:44:12Z"
					}
					updated_at = "2021-01-31T09:44:12Z"
				}
			}
			data {
				job_type = "repo_download_job"
				workspace_job_data {
					workspace_name = "workspace_name"
					flow_id = "flow_id"
					flow_name = "flow_name"
					inputs {
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
					outputs {
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
					template_data {
						template_id = "template_id"
						template_name = "template_name"
						flow_index = 1
						inputs {
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
						outputs {
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
						updated_at = "2021-01-31T09:44:12Z"
					}
					updated_at = "2021-01-31T09:44:12Z"
				}
				action_job_data {
					action_name = "action_name"
					inputs {
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
					outputs {
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
					updated_at = "2021-01-31T09:44:12Z"
					inventory_record {
						name = "name"
						id = "id"
						description = "description"
						location = "us-south"
						resource_group = "resource_group"
						created_at = "2021-01-31T09:44:12Z"
						created_by = "created_by"
						updated_at = "2021-01-31T09:44:12Z"
						updated_by = "updated_by"
						inventories_ini = "inventories_ini"
						resource_queries = [ "resource_queries" ]
					}
					materialized_inventory = "materialized_inventory"
				}
				system_job_data {
					key_id = "key_id"
					schematics_resource_id = [ "schematics_resource_id" ]
					updated_at = "2021-01-31T09:44:12Z"
				}
				flow_job_data {
					flow_id = "flow_id"
					flow_name = "flow_name"
					workitems {
						command_object_id = "command_object_id"
						command_object_name = "command_object_name"
						layers = "layers"
						source_type = "local"
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
						inputs {
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
						outputs {
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
						last_job {
							command_object = "workspace"
							command_object_name = "command_object_name"
							command_object_id = "command_object_id"
							command_name = "workspace_plan"
							job_id = "job_id"
							job_status = "job_pending"
						}
						updated_at = "2021-01-31T09:44:12Z"
					}
					updated_at = "2021-01-31T09:44:12Z"
				}
			}
			bastion {
				name = "name"
				host = "host"
			}
			log_summary {
				job_id = "job_id"
				job_type = "repo_download_job"
				log_start_at = "2021-01-31T09:44:12Z"
				log_analyzed_till = "2021-01-31T09:44:12Z"
				elapsed_time = 1.0
				log_errors {
					error_code = "error_code"
					error_msg = "error_msg"
					error_count = 1.0
				}
				repo_download_job {
					scanned_file_count = 1.0
					quarantined_file_count = 1.0
					detected_filetype = "detected_filetype"
					inputs_count = "inputs_count"
					outputs_count = "outputs_count"
				}
				workspace_job {
					resources_add = 1.0
					resources_modify = 1.0
					resources_destroy = 1.0
				}
				flow_job {
					workitems_completed = 1.0
					workitems_pending = 1.0
					workitems_failed = 1.0
					workitems {
						workspace_id = "workspace_id"
						job_id = "job_id"
						resources_add = 1.0
						resources_modify = 1.0
						resources_destroy = 1.0
						log_url = "log_url"
					}
				}
				action_job {
					target_count = 1.0
					task_count = 1.0
					play_count = 1.0
					recap {
						target = [ "target" ]
						ok = 1.0
						changed = 1.0
						failed = 1.0
						skipped = 1.0
						unreachable = 1.0
					}
				}
				system_job {
					target_count = 1.0
					success = 1.0
					failed = 1.0
				}
			}
		}
	`, refreshToken, commandObject, commandObjectID, commandName, commandParameter, location)
}

func testAccCheckIBMSchematicsJobExists(n string, obj schematicsv1.Job) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		schematicsClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).SchematicsV1()
		if err != nil {
			return err
		}

		getJobOptions := &schematicsv1.GetJobOptions{}

		getJobOptions.SetJobID(rs.Primary.ID)

		job, _, err := schematicsClient.GetJob(getJobOptions)
		if err != nil {
			return err
		}

		obj = *job
		return nil
	}
}

func testAccCheckIBMSchematicsJobDestroy(s *terraform.State) error {
	schematicsClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).SchematicsV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_schematics_job" {
			continue
		}

		getJobOptions := &schematicsv1.GetJobOptions{}

		getJobOptions.SetJobID(rs.Primary.ID)

		// Try to find the key
		_, response, err := schematicsClient.GetJob(getJobOptions)

		if err == nil {
			return fmt.Errorf("schematics_job still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for schematics_job (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
