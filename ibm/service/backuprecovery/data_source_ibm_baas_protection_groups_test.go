// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.94.0-fa797aec-20240814-142622
 */

package backuprecovery_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmBaasProtectionGroupsDataSourceBasic(t *testing.T) {
	protectionGroupResponseTenantID := fmt.Sprintf("%d", acctest.RandIntRange(10, 100))
	protectionGroupResponseName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	protectionGroupResponsePolicyID := fmt.Sprintf("tf_policy_id_%d", acctest.RandIntRange(10, 100))
	protectionGroupResponseEnvironment := "kPhysical"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmBaasProtectionGroupsDataSourceConfigBasic(protectionGroupResponseTenantID, protectionGroupResponseName, protectionGroupResponsePolicyID, protectionGroupResponseEnvironment),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_baas_protection_groups.baas_protection_groups_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_protection_groups.baas_protection_groups_instance", "x_ibm_tenant_id"),
				),
			},
		},
	})
}

func TestAccIbmBaasProtectionGroupsDataSourceAllArgs(t *testing.T) {
	protectionGroupResponseTenantID := fmt.Sprintf("%d", acctest.RandIntRange(10, 100))
	protectionGroupResponseName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	protectionGroupResponsePolicyID := fmt.Sprintf("tf_policy_id_%d", acctest.RandIntRange(10, 100))
	protectionGroupResponsePriority := "kLow"
	protectionGroupResponseDescription := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	protectionGroupResponseEndTimeUsecs := fmt.Sprintf("%d", acctest.RandIntRange(10, 100))
	protectionGroupResponseLastModifiedTimestampUsecs := fmt.Sprintf("%d", acctest.RandIntRange(10, 100))
	protectionGroupResponseQosPolicy := "kBackupHDD"
	protectionGroupResponseAbortInBlackouts := "true"
	protectionGroupResponsePauseInBlackouts := "true"
	protectionGroupResponseIsPaused := "false"
	protectionGroupResponseEnvironment := "kPhysical"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmBaasProtectionGroupsDataSourceConfig(protectionGroupResponseTenantID, protectionGroupResponseName, protectionGroupResponsePolicyID, protectionGroupResponsePriority, protectionGroupResponseDescription, protectionGroupResponseEndTimeUsecs, protectionGroupResponseLastModifiedTimestampUsecs, protectionGroupResponseQosPolicy, protectionGroupResponseAbortInBlackouts, protectionGroupResponsePauseInBlackouts, protectionGroupResponseIsPaused, protectionGroupResponseEnvironment),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_baas_protection_groups.baas_protection_groups_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_protection_groups.baas_protection_groups_instance", "x_ibm_tenant_id"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_protection_groups.baas_protection_groups_instance", "request_initiator_type"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_protection_groups.baas_protection_groups_instance", "ids"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_protection_groups.baas_protection_groups_instance", "names"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_protection_groups.baas_protection_groups_instance", "policy_ids"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_protection_groups.baas_protection_groups_instance", "include_groups_with_datalock_only"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_protection_groups.baas_protection_groups_instance", "environments"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_protection_groups.baas_protection_groups_instance", "is_active"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_protection_groups.baas_protection_groups_instance", "is_deleted"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_protection_groups.baas_protection_groups_instance", "is_paused"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_protection_groups.baas_protection_groups_instance", "last_run_local_backup_status"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_protection_groups.baas_protection_groups_instance", "last_run_replication_status"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_protection_groups.baas_protection_groups_instance", "last_run_archival_status"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_protection_groups.baas_protection_groups_instance", "last_run_cloud_spin_status"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_protection_groups.baas_protection_groups_instance", "last_run_any_status"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_protection_groups.baas_protection_groups_instance", "is_last_run_sla_violated"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_protection_groups.baas_protection_groups_instance", "include_last_run_info"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_protection_groups.baas_protection_groups_instance", "prune_excluded_source_ids"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_protection_groups.baas_protection_groups_instance", "prune_source_ids"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_protection_groups.baas_protection_groups_instance", "use_cached_data"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_protection_groups.baas_protection_groups_instance", "source_ids"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_protection_groups.baas_protection_groups_instance", "protection_groups.#"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_protection_groups.baas_protection_groups_instance", "protection_groups.0.id"),
					resource.TestCheckResourceAttr("data.ibm_baas_protection_groups.baas_protection_groups_instance", "protection_groups.0.name", protectionGroupResponseName),
					resource.TestCheckResourceAttrSet("data.ibm_baas_protection_groups.baas_protection_groups_instance", "protection_groups.0.cluster_id"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_protection_groups.baas_protection_groups_instance", "protection_groups.0.region_id"),
					resource.TestCheckResourceAttr("data.ibm_baas_protection_groups.baas_protection_groups_instance", "protection_groups.0.policy_id", protectionGroupResponsePolicyID),
					resource.TestCheckResourceAttr("data.ibm_baas_protection_groups.baas_protection_groups_instance", "protection_groups.0.priority", protectionGroupResponsePriority),
					resource.TestCheckResourceAttr("data.ibm_baas_protection_groups.baas_protection_groups_instance", "protection_groups.0.description", protectionGroupResponseDescription),
					resource.TestCheckResourceAttr("data.ibm_baas_protection_groups.baas_protection_groups_instance", "protection_groups.0.end_time_usecs", protectionGroupResponseEndTimeUsecs),
					resource.TestCheckResourceAttr("data.ibm_baas_protection_groups.baas_protection_groups_instance", "protection_groups.0.last_modified_timestamp_usecs", protectionGroupResponseLastModifiedTimestampUsecs),
					resource.TestCheckResourceAttr("data.ibm_baas_protection_groups.baas_protection_groups_instance", "protection_groups.0.qos_policy", protectionGroupResponseQosPolicy),
					resource.TestCheckResourceAttr("data.ibm_baas_protection_groups.baas_protection_groups_instance", "protection_groups.0.abort_in_blackouts", protectionGroupResponseAbortInBlackouts),
					resource.TestCheckResourceAttr("data.ibm_baas_protection_groups.baas_protection_groups_instance", "protection_groups.0.pause_in_blackouts", protectionGroupResponsePauseInBlackouts),
					resource.TestCheckResourceAttrSet("data.ibm_baas_protection_groups.baas_protection_groups_instance", "protection_groups.0.is_active"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_protection_groups.baas_protection_groups_instance", "protection_groups.0.is_deleted"),
					resource.TestCheckResourceAttr("data.ibm_baas_protection_groups.baas_protection_groups_instance", "protection_groups.0.is_paused", protectionGroupResponseIsPaused),
					resource.TestCheckResourceAttr("data.ibm_baas_protection_groups.baas_protection_groups_instance", "protection_groups.0.environment", protectionGroupResponseEnvironment),
					resource.TestCheckResourceAttrSet("data.ibm_baas_protection_groups.baas_protection_groups_instance", "protection_groups.0.is_protect_once"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_protection_groups.baas_protection_groups_instance", "protection_groups.0.num_protected_objects"),
				),
			},
		},
	})
}

func testAccCheckIbmBaasProtectionGroupsDataSourceConfigBasic(protectionGroupResponseTenantID string, protectionGroupResponseName string, protectionGroupResponsePolicyID string, protectionGroupResponseEnvironment string) string {
	return fmt.Sprintf(`
		resource "ibm_baas_protection_group" "baas_protection_group_instance" {
			tenant_id = %s
			name = "%s"
			policy_id = "%s"
			environment = "%s"
		}

		data "ibm_baas_protection_groups" "baas_protection_groups_instance" {
			tenant_id = ibm_baas_protection_group.baas_protection_group_instance.tenant_id
			request_initiator_type = "UIUser"
			ids = [ "ids" ]
			names = [ "names" ]
			policy_ids = [ "policyIds" ]
			include_groups_with_datalock_only = true
			environments = [ "kPhysical" ]
			is_active = ibm_baas_protection_group.baas_protection_group_instance.is_active
			is_deleted = ibm_baas_protection_group.baas_protection_group_instance.is_deleted
			is_paused = ibm_baas_protection_group.baas_protection_group_instance.is_paused
			last_run_local_backup_status = [ "Accepted" ]
			last_run_replication_status = [ "Accepted" ]
			last_run_archival_status = [ "Accepted" ]
			last_run_cloud_spin_status = [ "Accepted" ]
			last_run_any_status = [ "Accepted" ]
			is_last_run_sla_violated = true
			include_last_run_info = true
			prune_excluded_source_ids = true
			prune_source_ids = true
			use_cached_data = true
			source_ids = [ 1 ]
		}
	`, protectionGroupResponseTenantID, protectionGroupResponseName, protectionGroupResponsePolicyID, protectionGroupResponseEnvironment)
}

func testAccCheckIbmBaasProtectionGroupsDataSourceConfig(protectionGroupResponseTenantID string, protectionGroupResponseName string, protectionGroupResponsePolicyID string, protectionGroupResponsePriority string, protectionGroupResponseDescription string, protectionGroupResponseEndTimeUsecs string, protectionGroupResponseLastModifiedTimestampUsecs string, protectionGroupResponseQosPolicy string, protectionGroupResponseAbortInBlackouts string, protectionGroupResponsePauseInBlackouts string, protectionGroupResponseIsPaused string, protectionGroupResponseEnvironment string) string {
	return fmt.Sprintf(`
		resource "ibm_baas_protection_group" "baas_protection_group_instance" {
			tenant_id = %s
			name = "%s"
			policy_id = "%s"
			priority = "%s"
			description = "%s"
			start_time {
				hour = 0
				minute = 0
				time_zone = "time_zone"
			}
			end_time_usecs = %s
			last_modified_timestamp_usecs = %s
			alert_policy {
				backup_run_status = [ "kSuccess" ]
				alert_targets {
					email_address = "email_address"
					language = "en-us"
					recipient_type = "kTo"
				}
				raise_object_level_failure_alert = true
				raise_object_level_failure_alert_after_last_attempt = true
				raise_object_level_failure_alert_after_each_attempt = true
			}
			sla {
				backup_run_type = "kIncremental"
				sla_minutes = 1
			}
			qos_policy = "%s"
			abort_in_blackouts = %s
			pause_in_blackouts = %s
			is_paused = %s
			environment = "%s"
			advanced_configs {
				key = "key"
				value = "value"
			}
			physical_params {
				protection_type = "kFile"
				volume_protection_type_params {
					objects {
						id = 1
						name = "name"
						volume_guids = [ "volumeGuids" ]
						enable_system_backup = true
						excluded_vss_writers = [ "excludedVssWriters" ]
					}
					indexing_policy {
						enable_indexing = true
						include_paths = [ "includePaths" ]
						exclude_paths = [ "excludePaths" ]
					}
					perform_source_side_deduplication = true
					quiesce = true
					continue_on_quiesce_failure = true
					incremental_backup_after_restart = true
					pre_post_script {
						pre_script {
							path = "path"
							params = "params"
							timeout_secs = 1
							is_active = true
							continue_on_error = true
						}
						post_script {
							path = "path"
							params = "params"
							timeout_secs = 1
							is_active = true
						}
					}
					dedup_exclusion_source_ids = [ 1 ]
					excluded_vss_writers = [ "excludedVssWriters" ]
					cobmr_backup = true
				}
				file_protection_type_params {
					excluded_vss_writers = [ "excludedVssWriters" ]
					objects {
						excluded_vss_writers = [ "excludedVssWriters" ]
						id = 1
						name = "name"
						file_paths {
							included_path = "included_path"
							excluded_paths = [ "excludedPaths" ]
							skip_nested_volumes = true
						}
						uses_path_level_skip_nested_volume_setting = true
						nested_volume_types_to_skip = [ "nestedVolumeTypesToSkip" ]
						follow_nas_symlink_target = true
						metadata_file_path = "metadata_file_path"
					}
					indexing_policy {
						enable_indexing = true
						include_paths = [ "includePaths" ]
						exclude_paths = [ "excludePaths" ]
					}
					perform_source_side_deduplication = true
					perform_brick_based_deduplication = true
					task_timeouts {
						timeout_mins = 1
						backup_type = "kRegular"
					}
					quiesce = true
					continue_on_quiesce_failure = true
					cobmr_backup = true
					pre_post_script {
						pre_script {
							path = "path"
							params = "params"
							timeout_secs = 1
							is_active = true
							continue_on_error = true
						}
						post_script {
							path = "path"
							params = "params"
							timeout_secs = 1
							is_active = true
						}
					}
					dedup_exclusion_source_ids = [ 1 ]
					global_exclude_paths = [ "globalExcludePaths" ]
					global_exclude_fs = [ "globalExcludeFS" ]
					ignorable_errors = [ "kEOF" ]
					allow_parallel_runs = true
				}
			}
			mssql_params {
				file_protection_type_params {
					aag_backup_preference_type = "kPrimaryReplicaOnly"
					advanced_settings {
						cloned_db_backup_status = "kError"
						db_backup_if_not_online_status = "kError"
						missing_db_backup_status = "kError"
						offline_restoring_db_backup_status = "kError"
						read_only_db_backup_status = "kError"
						report_all_non_autoprotect_db_errors = "kError"
					}
					backup_system_dbs = true
					exclude_filters {
						filter_string = "filter_string"
						is_regular_expression = true
					}
					full_backups_copy_only = true
					log_backup_num_streams = 1
					log_backup_with_clause = "log_backup_with_clause"
					pre_post_script {
						pre_script {
							path = "path"
							params = "params"
							timeout_secs = 1
							is_active = true
							continue_on_error = true
						}
						post_script {
							path = "path"
							params = "params"
							timeout_secs = 1
							is_active = true
						}
					}
					use_aag_preferences_from_server = true
					user_db_backup_preference_type = "kBackupAllDatabases"
					additional_host_params {
						disable_source_side_deduplication = true
						host_id = 1
						host_name = "host_name"
					}
					objects {
						id = 1
						name = "name"
						source_type = "source_type"
					}
					perform_source_side_deduplication = true
				}
				native_protection_type_params {
					aag_backup_preference_type = "kPrimaryReplicaOnly"
					advanced_settings {
						cloned_db_backup_status = "kError"
						db_backup_if_not_online_status = "kError"
						missing_db_backup_status = "kError"
						offline_restoring_db_backup_status = "kError"
						read_only_db_backup_status = "kError"
						report_all_non_autoprotect_db_errors = "kError"
					}
					backup_system_dbs = true
					exclude_filters {
						filter_string = "filter_string"
						is_regular_expression = true
					}
					full_backups_copy_only = true
					log_backup_num_streams = 1
					log_backup_with_clause = "log_backup_with_clause"
					pre_post_script {
						pre_script {
							path = "path"
							params = "params"
							timeout_secs = 1
							is_active = true
							continue_on_error = true
						}
						post_script {
							path = "path"
							params = "params"
							timeout_secs = 1
							is_active = true
						}
					}
					use_aag_preferences_from_server = true
					user_db_backup_preference_type = "kBackupAllDatabases"
					num_streams = 1
					objects {
						id = 1
						name = "name"
						source_type = "source_type"
					}
					with_clause = "with_clause"
				}
				protection_type = "kFile"
				volume_protection_type_params {
					aag_backup_preference_type = "kPrimaryReplicaOnly"
					advanced_settings {
						cloned_db_backup_status = "kError"
						db_backup_if_not_online_status = "kError"
						missing_db_backup_status = "kError"
						offline_restoring_db_backup_status = "kError"
						read_only_db_backup_status = "kError"
						report_all_non_autoprotect_db_errors = "kError"
					}
					backup_system_dbs = true
					exclude_filters {
						filter_string = "filter_string"
						is_regular_expression = true
					}
					full_backups_copy_only = true
					log_backup_num_streams = 1
					log_backup_with_clause = "log_backup_with_clause"
					pre_post_script {
						pre_script {
							path = "path"
							params = "params"
							timeout_secs = 1
							is_active = true
							continue_on_error = true
						}
						post_script {
							path = "path"
							params = "params"
							timeout_secs = 1
							is_active = true
						}
					}
					use_aag_preferences_from_server = true
					user_db_backup_preference_type = "kBackupAllDatabases"
					additional_host_params {
						enable_system_backup = true
						host_id = 1
						host_name = "host_name"
						volume_guids = [ "volumeGuids" ]
					}
					backup_db_volumes_only = true
					incremental_backup_after_restart = true
					indexing_policy {
						enable_indexing = true
						include_paths = [ "includePaths" ]
						exclude_paths = [ "excludePaths" ]
					}
					objects {
						id = 1
						name = "name"
						source_type = "source_type"
					}
				}
			}
		}

		data "ibm_baas_protection_groups" "baas_protection_groups_instance" {
			tenant_id = ibm_baas_protection_group.baas_protection_group_instance.tenant_id
			request_initiator_type = "UIUser"
			ids = [ "ids" ]
			names = [ "names" ]
			policy_ids = [ "policyIds" ]
			include_groups_with_datalock_only = true
			environments = [ "kPhysical" ]
			is_active = ibm_baas_protection_group.baas_protection_group_instance.is_active
			is_deleted = ibm_baas_protection_group.baas_protection_group_instance.is_deleted
			is_paused = ibm_baas_protection_group.baas_protection_group_instance.is_paused
			last_run_local_backup_status = [ "Accepted" ]
			last_run_replication_status = [ "Accepted" ]
			last_run_archival_status = [ "Accepted" ]
			last_run_cloud_spin_status = [ "Accepted" ]
			last_run_any_status = [ "Accepted" ]
			is_last_run_sla_violated = true
			include_last_run_info = true
			prune_excluded_source_ids = true
			prune_source_ids = true
			use_cached_data = true
			source_ids = [ 1 ]
		}
	`, protectionGroupResponseTenantID, protectionGroupResponseName, protectionGroupResponsePolicyID, protectionGroupResponsePriority, protectionGroupResponseDescription, protectionGroupResponseEndTimeUsecs, protectionGroupResponseLastModifiedTimestampUsecs, protectionGroupResponseQosPolicy, protectionGroupResponseAbortInBlackouts, protectionGroupResponsePauseInBlackouts, protectionGroupResponseIsPaused, protectionGroupResponseEnvironment)
}
