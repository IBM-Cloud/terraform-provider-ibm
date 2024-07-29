// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package backuprecovery_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmRecoveriesDataSourceBasic(t *testing.T) {
	recoveryName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	recoverySnapshotEnvironment := "kPhysical"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmRecoveriesDataSourceConfigBasic(recoveryName, recoverySnapshotEnvironment),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_recoveries.recoveries_instance", "id"),
				),
			},
		},
	})
}

func TestAccIbmRecoveriesDataSourceAllArgs(t *testing.T) {
	recoveryRequestInitiatorType := "UIUser"
	recoveryName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	recoverySnapshotEnvironment := "kPhysical"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmRecoveriesDataSourceConfig(recoveryRequestInitiatorType, recoveryName, recoverySnapshotEnvironment),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_recoveries.recoveries_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_recoveries.recoveries_instance", "ids"),
					resource.TestCheckResourceAttrSet("data.ibm_recoveries.recoveries_instance", "return_only_child_recoveries"),
					resource.TestCheckResourceAttrSet("data.ibm_recoveries.recoveries_instance", "tenant_ids"),
					resource.TestCheckResourceAttrSet("data.ibm_recoveries.recoveries_instance", "include_tenants"),
					resource.TestCheckResourceAttrSet("data.ibm_recoveries.recoveries_instance", "start_time_usecs"),
					resource.TestCheckResourceAttrSet("data.ibm_recoveries.recoveries_instance", "end_time_usecs"),
					resource.TestCheckResourceAttrSet("data.ibm_recoveries.recoveries_instance", "storage_domain_id"),
					resource.TestCheckResourceAttrSet("data.ibm_recoveries.recoveries_instance", "snapshot_target_type"),
					resource.TestCheckResourceAttrSet("data.ibm_recoveries.recoveries_instance", "archival_target_type"),
					resource.TestCheckResourceAttrSet("data.ibm_recoveries.recoveries_instance", "snapshot_environments"),
					resource.TestCheckResourceAttrSet("data.ibm_recoveries.recoveries_instance", "status"),
					resource.TestCheckResourceAttrSet("data.ibm_recoveries.recoveries_instance", "recovery_actions"),
					resource.TestCheckResourceAttrSet("data.ibm_recoveries.recoveries_instance", "recoveries.#"),
					resource.TestCheckResourceAttrSet("data.ibm_recoveries.recoveries_instance", "recoveries.0.id"),
					resource.TestCheckResourceAttr("data.ibm_recoveries.recoveries_instance", "recoveries.0.name", recoveryName),
					resource.TestCheckResourceAttrSet("data.ibm_recoveries.recoveries_instance", "recoveries.0.start_time_usecs"),
					resource.TestCheckResourceAttrSet("data.ibm_recoveries.recoveries_instance", "recoveries.0.end_time_usecs"),
					resource.TestCheckResourceAttrSet("data.ibm_recoveries.recoveries_instance", "recoveries.0.status"),
					resource.TestCheckResourceAttrSet("data.ibm_recoveries.recoveries_instance", "recoveries.0.progress_task_id"),
					resource.TestCheckResourceAttr("data.ibm_recoveries.recoveries_instance", "recoveries.0.snapshot_environment", recoverySnapshotEnvironment),
					resource.TestCheckResourceAttrSet("data.ibm_recoveries.recoveries_instance", "recoveries.0.recovery_action"),
					resource.TestCheckResourceAttrSet("data.ibm_recoveries.recoveries_instance", "recoveries.0.can_tear_down"),
					resource.TestCheckResourceAttrSet("data.ibm_recoveries.recoveries_instance", "recoveries.0.tear_down_status"),
					resource.TestCheckResourceAttrSet("data.ibm_recoveries.recoveries_instance", "recoveries.0.tear_down_message"),
					resource.TestCheckResourceAttrSet("data.ibm_recoveries.recoveries_instance", "recoveries.0.is_parent_recovery"),
					resource.TestCheckResourceAttrSet("data.ibm_recoveries.recoveries_instance", "recoveries.0.parent_recovery_id"),
					resource.TestCheckResourceAttrSet("data.ibm_recoveries.recoveries_instance", "recoveries.0.is_multi_stage_restore"),
				),
			},
		},
	})
}

func testAccCheckIbmRecoveriesDataSourceConfigBasic(recoveryName string, recoverySnapshotEnvironment string) string {
	return fmt.Sprintf(`
		resource "ibm_recovery" "recovery_instance" {
			name = "%s"
			snapshot_environment = "%s"
		}

		data "ibm_recoveries" "recoveries_instance" {
			ids = [ "ids" ]
			return_only_child_recoveries = true
			tenant_ids = [ "tenantIds" ]
			include_tenants = true
			start_time_usecs = ibm_recovery.recovery_instance.start_time_usecs
			end_time_usecs = ibm_recovery.recovery_instance.end_time_usecs
			storage_domain_id = 1
			snapshot_target_type = [ "Local" ]
			archival_target_type = [ "Tape" ]
			snapshot_environments = [ "kPhysical" ]
			status = ibm_recovery.recovery_instance.status
			recovery_actions = [ "RecoverVMs" ]
		}
	`, recoveryName, recoverySnapshotEnvironment)
}

func testAccCheckIbmRecoveriesDataSourceConfig(recoveryRequestInitiatorType string, recoveryName string, recoverySnapshotEnvironment string) string {
	return fmt.Sprintf(`
		resource "ibm_recovery" "recovery_instance" {
			request_initiator_type = "%s"
			name = "%s"
			snapshot_environment = "%s"
			physical_params {
				objects {
					snapshot_id = "snapshot_id"
					point_in_time_usecs = 1
					protection_group_id = "protection_group_id"
					protection_group_name = "protection_group_name"
					snapshot_creation_time_usecs = 1
					object_info {
						id = 1
						name = "name"
						source_id = 1
						source_name = "source_name"
						environment = "kPhysical"
						object_hash = "object_hash"
						object_type = "kCluster"
						logical_size_bytes = 1
						uuid = "uuid"
						global_id = "global_id"
						protection_type = "kAgent"
						os_type = "kLinux"
					}
					snapshot_target_type = "Local"
					storage_domain_id = 1
					archival_target_info {
						target_id = 1
						archival_task_id = "archival_task_id"
						target_name = "target_name"
						target_type = "Tape"
						usage_type = "Archival"
						ownership_context = "Local"
						tier_settings {
							cloud_platform = "Oracle"
							oracle_tiering {
								tiers {
									move_after_unit = "Days"
									move_after = 1
									tier_type = "kOracleTierStandard"
								}
							}
							current_tier_type = "kAmazonS3Standard"
						}
					}
					progress_task_id = "progress_task_id"
					recover_from_standby = true
					status = "Accepted"
					start_time_usecs = 1
					end_time_usecs = 1
					messages = [ "messages" ]
					bytes_restored = 1
				}
				recovery_action = "RecoverPhysicalVolumes"
				recover_volume_params {
					target_environment = "kPhysical"
					physical_target_params {
						mount_target {
							id = 1
							name = "name"
						}
						volume_mapping {
							source_volume_guid = "source_volume_guid"
							destination_volume_guid = "destination_volume_guid"
						}
						force_unmount_volume = true
						vlan_config {
							id = 1
							disable_vlan = true
							interface_name = "interface_name"
						}
					}
				}
				mount_volume_params {
					target_environment = "kPhysical"
					physical_target_params {
						mount_to_original_target = true
						original_target_config {
							server_credentials {
								username = "username"
								password = "password"
							}
						}
						new_target_config {
							mount_target {
								id = 1
								name = "name"
								parent_source_id = 1
								parent_source_name = "parent_source_name"
							}
							server_credentials {
								username = "username"
								password = "password"
							}
						}
						read_only_mount = true
						volume_names = [ "volumeNames" ]
						mounted_volume_mapping {
							original_volume = "original_volume"
							mounted_volume = "mounted_volume"
							file_system_type = "file_system_type"
						}
						vlan_config {
							id = 1
							disable_vlan = true
							interface_name = "interface_name"
						}
					}
				}
				recover_file_and_folder_params {
					files_and_folders {
						absolute_path = "absolute_path"
						destination_dir = "destination_dir"
						is_directory = true
						status = "NotStarted"
						messages = [ "messages" ]
						is_view_file_recovery = true
					}
					target_environment = "kPhysical"
					physical_target_params {
						recover_target {
							id = 1
							name = "name"
							parent_source_id = 1
							parent_source_name = "parent_source_name"
						}
						restore_to_original_paths = true
						overwrite_existing = true
						alternate_restore_directory = "alternate_restore_directory"
						preserve_attributes = true
						preserve_timestamps = true
						preserve_acls = true
						continue_on_error = true
						save_success_files = true
						vlan_config {
							id = 1
							disable_vlan = true
							interface_name = "interface_name"
						}
						restore_entity_type = "kRegular"
					}
				}
				download_file_and_folder_params {
					files_and_folders {
						absolute_path = "absolute_path"
						destination_dir = "destination_dir"
						is_directory = true
						status = "NotStarted"
						messages = [ "messages" ]
						is_view_file_recovery = true
					}
					download_file_path = "download_file_path"
				}
				system_recovery_params {
					full_nas_path = "full_nas_path"
				}
			}
			oracle_params {
				objects {
					snapshot_id = "snapshot_id"
					point_in_time_usecs = 1
					protection_group_id = "protection_group_id"
					protection_group_name = "protection_group_name"
					snapshot_creation_time_usecs = 1
					object_info {
						id = 1
						name = "name"
						source_id = 1
						source_name = "source_name"
						environment = "kPhysical"
						object_hash = "object_hash"
						object_type = "kCluster"
						logical_size_bytes = 1
						uuid = "uuid"
						global_id = "global_id"
						protection_type = "kAgent"
						os_type = "kLinux"
					}
					snapshot_target_type = "Local"
					storage_domain_id = 1
					archival_target_info {
						target_id = 1
						archival_task_id = "archival_task_id"
						target_name = "target_name"
						target_type = "Tape"
						usage_type = "Archival"
						ownership_context = "Local"
						tier_settings {
							cloud_platform = "Oracle"
							oracle_tiering {
								tiers {
									move_after_unit = "Days"
									move_after = 1
									tier_type = "kOracleTierStandard"
								}
							}
							current_tier_type = "kAmazonS3Standard"
						}
					}
					progress_task_id = "progress_task_id"
					recover_from_standby = true
					status = "Accepted"
					start_time_usecs = 1
					end_time_usecs = 1
					messages = [ "messages" ]
					bytes_restored = 1
					instant_recovery_info {
						progress_task_id = "progress_task_id"
						status = "Accepted"
						start_time_usecs = 1
						end_time_usecs = 1
					}
				}
				recovery_action = "RecoverApps"
				recover_app_params {
					target_environment = "kOracle"
					oracle_target_params {
						recover_to_new_source = true
						new_source_config {
							host {
								id = 1
								name = "name"
							}
							recovery_target = "RecoverDatabase"
							recover_database_params {
								restore_time_usecs = 1
								db_channels {
									archive_log_retention_days = 1
									archive_log_retention_hours = 1
									credentials {
										username = "username"
										password = "password"
									}
									database_unique_name = "database_unique_name"
									database_uuid = "database_uuid"
									default_channel_count = 1
									database_node_list {
										host_id = "host_id"
										channel_count = 1
										port = 1
										sbt_host_params {
											sbt_library_path = "sbt_library_path"
											view_fs_path = "view_fs_path"
											vip_list = [ "vipList" ]
											vlan_info_list {
												ip_list = [ "ipList" ]
												gateway = "gateway"
												id = 1
												subnet_ip = "subnet_ip"
											}
										}
									}
									max_host_count = 1
									enable_dg_primary_backup = true
									rman_backup_type = "kImageCopy"
								}
								recovery_mode = true
								shell_evironment_vars {
									key = "key"
									value = "value"
								}
								granular_restore_info {
									granularity_type = "kPDB"
									pdb_restore_params {
										drop_duplicate_pdb = true
										pdb_objects {
											db_id = "db_id"
											db_name = "db_name"
										}
										restore_to_existing_cdb = true
										rename_pdb_map {
											key = "key"
											value = "value"
										}
										include_in_restore = true
									}
								}
								oracle_archive_log_info {
									range_type = "Time"
									range_info_vec {
										start_of_range = 1
										end_of_range = 1
										protection_group_id = "protection_group_id"
										reset_log_id = 1
										incarnation_id = 1
										thread_id = 1
									}
									archive_log_restore_dest = "archive_log_restore_dest"
								}
								oracle_recovery_validation_info {
									create_dummy_instance = true
								}
								restore_spfile_or_pfile_info {
									should_restore_spfile_or_pfile = true
									file_location = "file_location"
								}
								use_scn_for_restore = true
								database_name = "database_name"
								oracle_base_folder = "oracle_base_folder"
								oracle_home_folder = "oracle_home_folder"
								db_files_destination = "db_files_destination"
								db_config_file_path = "db_config_file_path"
								enable_archive_log_mode = true
								pfile_parameter_map {
									key = "key"
									value = "value"
								}
								bct_file_path = "bct_file_path"
								num_tempfiles = 1
								redo_log_config {
									num_groups = 1
									member_prefix = "member_prefix"
									size_m_bytes = 1
									group_members = [ "groupMembers" ]
								}
								is_multi_stage_restore = true
								oracle_update_restore_options {
									delay_secs = 1
									target_path_vec = [ "targetPathVec" ]
								}
								skip_clone_nid = true
								no_filename_check = true
								new_name_clause = "new_name_clause"
							}
							recover_view_params {
								restore_time_usecs = 1
								db_channels {
									archive_log_retention_days = 1
									archive_log_retention_hours = 1
									credentials {
										username = "username"
										password = "password"
									}
									database_unique_name = "database_unique_name"
									database_uuid = "database_uuid"
									default_channel_count = 1
									database_node_list {
										host_id = "host_id"
										channel_count = 1
										port = 1
										sbt_host_params {
											sbt_library_path = "sbt_library_path"
											view_fs_path = "view_fs_path"
											vip_list = [ "vipList" ]
											vlan_info_list {
												ip_list = [ "ipList" ]
												gateway = "gateway"
												id = 1
												subnet_ip = "subnet_ip"
											}
										}
									}
									max_host_count = 1
									enable_dg_primary_backup = true
									rman_backup_type = "kImageCopy"
								}
								recovery_mode = true
								shell_evironment_vars {
									key = "key"
									value = "value"
								}
								granular_restore_info {
									granularity_type = "kPDB"
									pdb_restore_params {
										drop_duplicate_pdb = true
										pdb_objects {
											db_id = "db_id"
											db_name = "db_name"
										}
										restore_to_existing_cdb = true
										rename_pdb_map {
											key = "key"
											value = "value"
										}
										include_in_restore = true
									}
								}
								oracle_archive_log_info {
									range_type = "Time"
									range_info_vec {
										start_of_range = 1
										end_of_range = 1
										protection_group_id = "protection_group_id"
										reset_log_id = 1
										incarnation_id = 1
										thread_id = 1
									}
									archive_log_restore_dest = "archive_log_restore_dest"
								}
								oracle_recovery_validation_info {
									create_dummy_instance = true
								}
								restore_spfile_or_pfile_info {
									should_restore_spfile_or_pfile = true
									file_location = "file_location"
								}
								use_scn_for_restore = true
								view_mount_path = "view_mount_path"
							}
						}
						original_source_config {
							restore_time_usecs = 1
							db_channels {
								archive_log_retention_days = 1
								archive_log_retention_hours = 1
								credentials {
									username = "username"
									password = "password"
								}
								database_unique_name = "database_unique_name"
								database_uuid = "database_uuid"
								default_channel_count = 1
								database_node_list {
									host_id = "host_id"
									channel_count = 1
									port = 1
									sbt_host_params {
										sbt_library_path = "sbt_library_path"
										view_fs_path = "view_fs_path"
										vip_list = [ "vipList" ]
										vlan_info_list {
											ip_list = [ "ipList" ]
											gateway = "gateway"
											id = 1
											subnet_ip = "subnet_ip"
										}
									}
								}
								max_host_count = 1
								enable_dg_primary_backup = true
								rman_backup_type = "kImageCopy"
							}
							recovery_mode = true
							shell_evironment_vars {
								key = "key"
								value = "value"
							}
							granular_restore_info {
								granularity_type = "kPDB"
								pdb_restore_params {
									drop_duplicate_pdb = true
									pdb_objects {
										db_id = "db_id"
										db_name = "db_name"
									}
									restore_to_existing_cdb = true
									rename_pdb_map {
										key = "key"
										value = "value"
									}
									include_in_restore = true
								}
							}
							oracle_archive_log_info {
								range_type = "Time"
								range_info_vec {
									start_of_range = 1
									end_of_range = 1
									protection_group_id = "protection_group_id"
									reset_log_id = 1
									incarnation_id = 1
									thread_id = 1
								}
								archive_log_restore_dest = "archive_log_restore_dest"
							}
							oracle_recovery_validation_info {
								create_dummy_instance = true
							}
							restore_spfile_or_pfile_info {
								should_restore_spfile_or_pfile = true
								file_location = "file_location"
							}
							use_scn_for_restore = true
							roll_forward_log_path_vec = [ "rollForwardLogPathVec" ]
							attempt_complete_recovery = true
							roll_forward_time_msecs = 1
							stop_active_passive = true
						}
					}
					vlan_config {
						id = 1
						disable_vlan = true
						interface_name = "interface_name"
					}
				}
			}
		}

		data "ibm_recoveries" "recoveries_instance" {
			ids = [ "ids" ]
			return_only_child_recoveries = true
			tenant_ids = [ "tenantIds" ]
			include_tenants = true
			start_time_usecs = ibm_recovery.recovery_instance.start_time_usecs
			end_time_usecs = ibm_recovery.recovery_instance.end_time_usecs
			storage_domain_id = 1
			snapshot_target_type = [ "Local" ]
			archival_target_type = [ "Tape" ]
			snapshot_environments = [ "kPhysical" ]
			status = ibm_recovery.recovery_instance.status
			recovery_actions = [ "RecoverVMs" ]
		}
	`, recoveryRequestInitiatorType, recoveryName, recoverySnapshotEnvironment)
}
