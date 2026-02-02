---
layout: "ibm"
page_title: "IBM : ibm_backup_recovery"
description: |-
  Manages Common Recovery Response Params..
subcategory: "IBM Backup Recovery"
---

# ibm_backup_recovery

Create Recovery with this resource.

**Note**
ibm_backup_recovery resource does not support update or delete operations due to the absence of corresponding API endpoints. As a result, Terraform cannot manage these operations for those resources. Users should be aware that removing these resources from the configuration (main.tf) will only remove them from the Terraform state and will not affect the actual resources in the backend. Similarly updating these resources will throw an error in the plan phase stating that the resource cannot be updated.

**Important:** When managing resources that lack complete CRUD operations, users should exercise caution and consider the limitations described above. Manual intervention may be required to manage these resources in the backend if updates or deletions are necessary.**

## Example Usage

```hcl
resource "ibm_backup_recovery" "backup_recovery_instance" {
  kubernetes_params {
		download_file_and_folder_params {
			expiry_time_usecs = 1
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
				sharepoint_site_summary {
					site_web_url = "site_web_url"
				}
				os_type = "kLinux"
				child_objects {
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
					sharepoint_site_summary {
						site_web_url = "site_web_url"
					}
					os_type = "kLinux"
					v_center_summary {
						is_cloud_env = true
					}
					windows_cluster_summary {
						cluster_source_type = "cluster_source_type"
					}
				}
				v_center_summary {
					is_cloud_env = true
				}
				windows_cluster_summary {
					cluster_source_type = "cluster_source_type"
				}
			}
			snapshot_target_type = "Local"
			archival_target_info {
				target_id = 1
				archival_task_id = "archival_task_id"
				target_name = "target_name"
				target_type = "Tape"
				usage_type = "Archival"
				ownership_context = "Local"
				tier_settings {
					aws_tiering {
						tiers {
							move_after_unit = "Days"
							move_after = 1
							tier_type = "kAmazonS3Standard"
						}
					}
					azure_tiering {
						tiers {
							move_after_unit = "Days"
							move_after = 1
							tier_type = "kAzureTierHot"
						}
					}
					cloud_platform = "AWS"
					google_tiering {
						tiers {
							move_after_unit = "Days"
							move_after = 1
							tier_type = "kGoogleStandard"
						}
					}
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
		recover_file_and_folder_params {
			files_and_folders {
				absolute_path = "absolute_path"
				destination_dir = "destination_dir"
				is_directory = true
				status = "NotStarted"
				messages = [ "messages" ]
				is_view_file_recovery = true
			}
			kubernetes_target_params {
				continue_on_error = true
				new_target_config {
					absolute_path = "absolute_path"
					target_namespace {
						id = 1
						name = "name"
						parent_source_id = 1
						parent_source_name = "parent_source_name"
					}
					target_pvc {
						id = 1
						name = "name"
						parent_source_id = 1
						parent_source_name = "parent_source_name"
					}
					target_source {
						id = 1
						name = "name"
						parent_source_id = 1
						parent_source_name = "parent_source_name"
					}
				}
				original_target_config {
					alternate_path = "alternate_path"
					recover_to_original_path = true
				}
				overwrite_existing = true
				preserve_attributes = true
				recover_to_original_target = true
				vlan_config {
					id = 1
					disable_vlan = true
					interface_name = "interface_name"
				}
			}
			target_environment = "kKubernetes"
		}
		recover_namespace_params {
			kubernetes_target_params {
				exclude_params {
					label_combination_method = "AND"
					label_vector {
						key = "key"
						value = "value"
					}
					objects = [ 1 ]
					selected_resources {
						api_group = "api_group"
						is_cluster_scoped = true
						kind = "kind"
						name = "name"
						resource_list {
							entity_id = 1
							name = "name"
						}
						version = "version"
					}
				}
				excluded_pvcs {
					id = 1
					name = "name"
				}
				include_params {
					label_combination_method = "AND"
					label_vector {
						key = "key"
						value = "value"
					}
					objects = [ 1 ]
					selected_resources {
						api_group = "api_group"
						is_cluster_scoped = true
						kind = "kind"
						name = "name"
						resource_list {
							entity_id = 1
							name = "name"
						}
						version = "version"
					}
				}
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
						sharepoint_site_summary {
							site_web_url = "site_web_url"
						}
						os_type = "kLinux"
						child_objects {
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
							sharepoint_site_summary {
								site_web_url = "site_web_url"
							}
							os_type = "kLinux"
							v_center_summary {
								is_cloud_env = true
							}
							windows_cluster_summary {
								cluster_source_type = "cluster_source_type"
							}
						}
						v_center_summary {
							is_cloud_env = true
						}
						windows_cluster_summary {
							cluster_source_type = "cluster_source_type"
						}
					}
					snapshot_target_type = "Local"
					archival_target_info {
						target_id = 1
						archival_task_id = "archival_task_id"
						target_name = "target_name"
						target_type = "Tape"
						usage_type = "Archival"
						ownership_context = "Local"
						tier_settings {
							aws_tiering {
								tiers {
									move_after_unit = "Days"
									move_after = 1
									tier_type = "kAmazonS3Standard"
								}
							}
							azure_tiering {
								tiers {
									move_after_unit = "Days"
									move_after = 1
									tier_type = "kAzureTierHot"
								}
							}
							cloud_platform = "AWS"
							google_tiering {
								tiers {
									move_after_unit = "Days"
									move_after = 1
									tier_type = "kGoogleStandard"
								}
							}
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
					exclude_params {
						label_combination_method = "AND"
						label_vector {
							key = "key"
							value = "value"
						}
						objects = [ 1 ]
						selected_resources {
							api_group = "api_group"
							is_cluster_scoped = true
							kind = "kind"
							name = "name"
							resource_list {
								entity_id = 1
								name = "name"
							}
							version = "version"
						}
					}
					include_params {
						label_combination_method = "AND"
						label_vector {
							key = "key"
							value = "value"
						}
						objects = [ 1 ]
						selected_resources {
							api_group = "api_group"
							is_cluster_scoped = true
							kind = "kind"
							name = "name"
							resource_list {
								entity_id = 1
								name = "name"
							}
							version = "version"
						}
					}
					recover_pvcs_only = true
					storage_class {
						storage_class_mapping {
							key = "key"
							value = "value"
						}
						use_storage_class_mapping = true
					}
					unbind_pvcs = true
				}
				recover_cluster_scoped_resources {
					snapshot_id = "snapshot_id"
				}
				recover_protection_group_runs_params {
					archival_target_id = 1
					protection_group_id = "protection_group_id"
					protection_group_instance_id = 1
					protection_group_run_id = "protection_group_run_id"
				}
				recover_pvcs_only = true
				recovery_region_migration_params {
					current_value = "current_value"
					new_value = "new_value"
				}
				recovery_target_config {
					new_source_config {
						source {
							id = 1
							name = "name"
						}
					}
					recover_to_new_source = true
				}
				recovery_zone_migration_params {
					current_value = "current_value"
					new_value = "new_value"
				}
				rename_recovered_namespaces_params {
					prefix = "prefix"
					suffix = "suffix"
				}
				skip_cluster_compatibility_check = true
				storage_class {
					storage_class_mapping {
						key = "key"
						value = "value"
					}
					use_storage_class_mapping = true
				}
			}
			target_environment = "kKubernetes"
			vlan_config {
				id = 1
				disable_vlan = true
				interface_name = "interface_name"
			}
		}
		recovery_action = "RecoverNamespaces"
  }
  mssql_params {
		recover_app_params {
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
				sharepoint_site_summary {
					site_web_url = "site_web_url"
				}
				os_type = "kLinux"
				child_objects {
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
					sharepoint_site_summary {
						site_web_url = "site_web_url"
					}
					os_type = "kLinux"
					v_center_summary {
						is_cloud_env = true
					}
					windows_cluster_summary {
						cluster_source_type = "cluster_source_type"
					}
				}
				v_center_summary {
					is_cloud_env = true
				}
				windows_cluster_summary {
					cluster_source_type = "cluster_source_type"
				}
			}
			snapshot_target_type = "Local"
			archival_target_info {
				target_id = 1
				archival_task_id = "archival_task_id"
				target_name = "target_name"
				target_type = "Tape"
				usage_type = "Archival"
				ownership_context = "Local"
				tier_settings {
					aws_tiering {
						tiers {
							move_after_unit = "Days"
							move_after = 1
							tier_type = "kAmazonS3Standard"
						}
					}
					azure_tiering {
						tiers {
							move_after_unit = "Days"
							move_after = 1
							tier_type = "kAzureTierHot"
						}
					}
					cloud_platform = "AWS"
					google_tiering {
						tiers {
							move_after_unit = "Days"
							move_after = 1
							tier_type = "kGoogleStandard"
						}
					}
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
			aag_info {
				name = "name"
				object_id = 1
			}
			host_info {
				id = "id"
				name = "name"
				environment = "kPhysical"
			}
			is_encrypted = true
			sql_target_params {
				new_source_config {
					keep_cdc = true
					multi_stage_restore_options {
						enable_auto_sync = true
						enable_multi_stage_restore = true
					}
					native_log_recovery_with_clause = "native_log_recovery_with_clause"
					native_recovery_with_clause = "native_recovery_with_clause"
					overwriting_policy = "FailIfExists"
					replay_entire_last_log = true
					restore_time_usecs = 1
					secondary_data_files_dir_list {
						directory = "directory"
						filename_pattern = "filename_pattern"
					}
					with_no_recovery = true
					data_file_directory_location = "data_file_directory_location"
					database_name = "database_name"
					host {
						id = 1
						name = "name"
					}
					instance_name = "instance_name"
					log_file_directory_location = "log_file_directory_location"
				}
				original_source_config {
					keep_cdc = true
					multi_stage_restore_options {
						enable_auto_sync = true
						enable_multi_stage_restore = true
					}
					native_log_recovery_with_clause = "native_log_recovery_with_clause"
					native_recovery_with_clause = "native_recovery_with_clause"
					overwriting_policy = "FailIfExists"
					replay_entire_last_log = true
					restore_time_usecs = 1
					secondary_data_files_dir_list {
						directory = "directory"
						filename_pattern = "filename_pattern"
					}
					with_no_recovery = true
					capture_tail_logs = true
					data_file_directory_location = "data_file_directory_location"
					log_file_directory_location = "log_file_directory_location"
					new_database_name = "new_database_name"
				}
				recover_to_new_source = true
			}
			target_environment = "kSQL"
		}
		recovery_action = "RecoverApps"
		vlan_config {
			id = 1
			disable_vlan = true
			interface_name = "interface_name"
		}
  }
  name = "name"
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
				sharepoint_site_summary {
					site_web_url = "site_web_url"
				}
				os_type = "kLinux"
				child_objects {
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
					sharepoint_site_summary {
						site_web_url = "site_web_url"
					}
					os_type = "kLinux"
					v_center_summary {
						is_cloud_env = true
					}
					windows_cluster_summary {
						cluster_source_type = "cluster_source_type"
					}
				}
				v_center_summary {
					is_cloud_env = true
				}
				windows_cluster_summary {
					cluster_source_type = "cluster_source_type"
				}
			}
			snapshot_target_type = "Local"
			archival_target_info {
				target_id = 1
				archival_task_id = "archival_task_id"
				target_name = "target_name"
				target_type = "Tape"
				usage_type = "Archival"
				ownership_context = "Local"
				tier_settings {
					aws_tiering {
						tiers {
							move_after_unit = "Days"
							move_after = 1
							tier_type = "kAmazonS3Standard"
						}
					}
					azure_tiering {
						tiers {
							move_after_unit = "Days"
							move_after = 1
							tier_type = "kAzureTierHot"
						}
					}
					cloud_platform = "AWS"
					google_tiering {
						tiers {
							move_after_unit = "Days"
							move_after = 1
							tier_type = "kGoogleStandard"
						}
					}
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
			expiry_time_usecs = 1
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
  snapshot_environment = "kPhysical"
  x_ibm_tenant_id = "x_ibm_tenant_id"
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `kubernetes_params` - (Optional, Forces new resource, List) Specifies the recovery options specific to Kubernetes environment.
Nested schema for **kubernetes_params**:
	* `download_file_and_folder_params` - (Optional, List) Specifies the parameters to download files and folders.
	Nested schema for **download_file_and_folder_params**:
		* `download_file_path` - (Optional, String) Specifies the path location to download the files and folders.
		* `expiry_time_usecs` - (Optional, Integer) Specifies the time upto which the download link is available.
		* `files_and_folders` - (Optional, List) Specifies the info about the files and folders to be recovered.
		Nested schema for **files_and_folders**:
			* `absolute_path` - (Required, String) Specifies the absolute path to the file or folder.
			* `destination_dir` - (Computed, String) Specifies the destination directory where the file/directory was copied.
			* `is_directory` - (Optional, Boolean) Specifies whether this is a directory or not.
			* `is_view_file_recovery` - (Optional, Boolean) Specify if the recovery is of type view file/folder.
			* `messages` - (Computed, List) Specify error messages about the file during recovery.
			* `status` - (Computed, String) Specifies the recovery status for this file or folder.
			  * Constraints: Allowable values are: `NotStarted`, `EstimationInProgress`, `EstimationDone`, `CopyInProgress`, `Finished`.
	* `objects` - (Optional, List) Specifies the list of objects which need to be recovered.
	Nested schema for **objects**:
		* `archival_target_info` - (Optional, List) Specifies the archival target information if the snapshot is an archival snapshot.
		Nested schema for **archival_target_info**:
			* `archival_task_id` - (Computed, String) Specifies the archival task id. This is a protection group UID which only applies when archival type is 'Tape'.
			* `ownership_context` - (Computed, String) Specifies the ownership context for the target.
			  * Constraints: Allowable values are: `Local`, `FortKnox`.
			* `target_id` - (Computed, Integer) Specifies the archival target ID.
			* `target_name` - (Computed, String) Specifies the archival target name.
			* `target_type` - (Computed, String) Specifies the archival target type.
			  * Constraints: Allowable values are: `Tape`, `Cloud`, `Nas`.
			* `tier_settings` - (Optional, List) Specifies the tier info for archival.
			Nested schema for **tier_settings**:
				* `aws_tiering` - (Optional, List) Specifies aws tiers.
				Nested schema for **aws_tiering**:
					* `tiers` - (Required, List) Specifies the tiers that are used to move the archived backup from current tier to next tier. The order of the tiers determines which tier will be used next for moving the archived backup. The first tier input should always be default tier where backup will be acrhived. Each tier specifies how much time after the backup will be moved to next tier from the current tier.
					Nested schema for **tiers**:
						* `move_after` - (Computed, Integer) Specifies the time period after which the backup will be moved from current tier to next tier.
						* `move_after_unit` - (Computed, String) Specifies the unit for moving the data from current tier to next tier. This unit will be a base unit for the 'moveAfter' field specified below.
						  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
						* `tier_type` - (Computed, String) Specifies the AWS tier types.
						  * Constraints: Allowable values are: `kAmazonS3Standard`, `kAmazonS3StandardIA`, `kAmazonS3OneZoneIA`, `kAmazonS3IntelligentTiering`, `kAmazonS3Glacier`, `kAmazonS3GlacierDeepArchive`.
				* `azure_tiering` - (Optional, List) Specifies Azure tiers.
				Nested schema for **azure_tiering**:
					* `tiers` - (Optional, List) Specifies the tiers that are used to move the archived backup from current tier to next tier. The order of the tiers determines which tier will be used next for moving the archived backup. The first tier input should always be default tier where backup will be acrhived. Each tier specifies how much time after the backup will be moved to next tier from the current tier.
					Nested schema for **tiers**:
						* `move_after` - (Computed, Integer) Specifies the time period after which the backup will be moved from current tier to next tier.
						* `move_after_unit` - (Computed, String) Specifies the unit for moving the data from current tier to next tier. This unit will be a base unit for the 'moveAfter' field specified below.
						  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
						* `tier_type` - (Computed, String) Specifies the Azure tier types.
						  * Constraints: Allowable values are: `kAzureTierHot`, `kAzureTierCool`, `kAzureTierArchive`.
				* `cloud_platform` - (Computed, String) Specifies the cloud platform to enable tiering.
				  * Constraints: Allowable values are: `AWS`, `Azure`, `Oracle`, `Google`.
				* `current_tier_type` - (Computed, String) Specifies the type of the current tier where the snapshot resides. This will be specified if the run is a CAD run.
				  * Constraints: Allowable values are: `kAmazonS3Standard`, `kAmazonS3StandardIA`, `kAmazonS3OneZoneIA`, `kAmazonS3IntelligentTiering`, `kAmazonS3Glacier`, `kAmazonS3GlacierDeepArchive`, `kAzureTierHot`, `kAzureTierCool`, `kAzureTierArchive`, `kGoogleStandard`, `kGoogleRegional`, `kGoogleMultiRegional`, `kGoogleNearline`, `kGoogleColdline`, `kOracleTierStandard`, `kOracleTierArchive`.
				* `google_tiering` - (Optional, List) Specifies Google tiers.
				Nested schema for **google_tiering**:
					* `tiers` - (Required, List) Specifies the tiers that are used to move the archived backup from current tier to next tier. The order of the tiers determines which tier will be used next for moving the archived backup. The first tier input should always be default tier where backup will be acrhived. Each tier specifies how much time after the backup will be moved to next tier from the current tier.
					Nested schema for **tiers**:
						* `move_after` - (Computed, Integer) Specifies the time period after which the backup will be moved from current tier to next tier.
						* `move_after_unit` - (Computed, String) Specifies the unit for moving the data from current tier to next tier. This unit will be a base unit for the 'moveAfter' field specified below.
						  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
						* `tier_type` - (Computed, String) Specifies the Google tier types.
						  * Constraints: Allowable values are: `kGoogleStandard`, `kGoogleRegional`, `kGoogleMultiRegional`, `kGoogleNearline`, `kGoogleColdline`.
				* `oracle_tiering` - (Optional, List) Specifies Oracle tiers.
				Nested schema for **oracle_tiering**:
					* `tiers` - (Required, List) Specifies the tiers that are used to move the archived backup from current tier to next tier. The order of the tiers determines which tier will be used next for moving the archived backup. The first tier input should always be default tier where backup will be acrhived. Each tier specifies how much time after the backup will be moved to next tier from the current tier.
					Nested schema for **tiers**:
						* `move_after` - (Computed, Integer) Specifies the time period after which the backup will be moved from current tier to next tier.
						* `move_after_unit` - (Computed, String) Specifies the unit for moving the data from current tier to next tier. This unit will be a base unit for the 'moveAfter' field specified below.
						  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
						* `tier_type` - (Computed, String) Specifies the Oracle tier types.
						  * Constraints: Allowable values are: `kOracleTierStandard`, `kOracleTierArchive`.
			* `usage_type` - (Computed, String) Specifies the usage type for the target.
			  * Constraints: Allowable values are: `Archival`, `Tiering`, `Rpaas`.
		* `bytes_restored` - (Computed, Integer) Specify the total bytes restored.
		* `end_time_usecs` - (Computed, Integer) Specifies the end time of the Recovery in Unix timestamp epoch in microseconds. This field will be populated only after Recovery is finished.
		* `messages` - (Computed, List) Specify error messages about the object.
		* `object_info` - (Optional, List) Specifies the information about the object for which the snapshot is taken.
		Nested schema for **object_info**:
			* `child_objects` - (Optional, List) Specifies child object details.
			Nested schema for **child_objects**:
				* `child_objects` - (Optional, List) Specifies child object details.
				Nested schema for **child_objects**:
				* `environment` - (Computed, String) Specifies the environment of the object.
				  * Constraints: Allowable values are: `kPhysical`, `kSQL`.
				* `global_id` - (Computed, String) Specifies the global id which is a unique identifier of the object.
				* `id` - (Computed, Integer) Specifies object id.
				* `logical_size_bytes` - (Computed, Integer) Specifies the logical size of object in bytes.
				* `name` - (Computed, String) Specifies the name of the object.
				* `object_hash` - (Computed, String) Specifies the hash identifier of the object.
				* `object_type` - (Computed, String) Specifies the type of the object.
				  * Constraints: Allowable values are: `kCluster`, `kVserver`, `kVolume`, `kVCenter`, `kStandaloneHost`, `kvCloudDirector`, `kFolder`, `kDatacenter`, `kComputeResource`, `kClusterComputeResource`, `kResourcePool`, `kDatastore`, `kHostSystem`, `kVirtualMachine`, `kVirtualApp`, `kStoragePod`, `kNetwork`, `kDistributedVirtualPortgroup`, `kTagCategory`, `kTag`, `kOpaqueNetwork`, `kOrganization`, `kVirtualDatacenter`, `kCatalog`, `kOrgMetadata`, `kStoragePolicy`, `kVirtualAppTemplate`, `kDomain`, `kOutlook`, `kMailbox`, `kUsers`, `kGroups`, `kSites`, `kUser`, `kGroup`, `kSite`, `kApplication`, `kGraphUser`, `kPublicFolders`, `kPublicFolder`, `kTeams`, `kTeam`, `kRootPublicFolder`, `kO365Exchange`, `kO365OneDrive`, `kO365Sharepoint`, `kKeyspace`, `kTable`, `kDatabase`, `kCollection`, `kBucket`, `kNamespace`, `kSCVMMServer`, `kStandaloneCluster`, `kHostGroup`, `kHypervHost`, `kHostCluster`, `kCustomProperty`, `kTenant`, `kSubscription`, `kResourceGroup`, `kStorageAccount`, `kStorageKey`, `kStorageContainer`, `kStorageBlob`, `kNetworkSecurityGroup`, `kVirtualNetwork`, `kSubnet`, `kComputeOptions`, `kSnapshotManagerPermit`, `kAvailabilitySet`, `kOVirtManager`, `kHost`, `kStorageDomain`, `kVNicProfile`, `kIAMUser`, `kRegion`, `kAvailabilityZone`, `kEC2Instance`, `kVPC`, `kInstanceType`, `kKeyPair`, `kRDSOptionGroup`, `kRDSParameterGroup`, `kRDSInstance`, `kRDSSubnet`, `kRDSTag`, `kAuroraTag`, `kAuroraCluster`, `kAccount`, `kSubTaskPermit`, `kS3Bucket`, `kS3Tag`, `kKmsKey`, `kProject`, `kLabel`, `kMetadata`, `kVPCConnector`, `kPrismCentral`, `kOtherHypervisorCluster`, `kZone`, `kMountPoint`, `kStorageArray`, `kFileSystem`, `kContainer`, `kFilesystem`, `kFileset`, `kPureProtectionGroup`, `kVolumeGroup`, `kStoragePool`, `kViewBox`, `kView`, `kWindowsCluster`, `kOracleRACCluster`, `kOracleAPCluster`, `kService`, `kPVC`, `kPersistentVolumeClaim`, `kPersistentVolume`, `kRootContainer`, `kDAGRootContainer`, `kExchangeNode`, `kExchangeDAGDatabaseCopy`, `kExchangeStandaloneDatabase`, `kExchangeDAG`, `kExchangeDAGDatabase`, `kDomainController`, `kInstance`, `kAAG`, `kAAGRootContainer`, `kAAGDatabase`, `kRACRootContainer`, `kTableSpace`, `kPDB`, `kObject`, `kOrg`, `kAppInstance`.
				* `os_type` - (Computed, String) Specifies the operating system type of the object.
				  * Constraints: Allowable values are: `kLinux`, `kWindows`.
				* `protection_type` - (Computed, String) Specifies the protection type of the object if any.
				  * Constraints: Allowable values are: `kAgent`, `kNative`, `kSnapshotManager`, `kRDSSnapshotManager`, `kAuroraSnapshotManager`, `kAwsS3`, `kAwsRDSPostgresBackup`, `kAwsAuroraPostgres`, `kAwsRDSPostgres`, `kAzureSQL`, `kFile`, `kVolume`.
				* `sharepoint_site_summary` - (Optional, List) Specifies the common parameters for Sharepoint site objects.
				Nested schema for **sharepoint_site_summary**:
					* `site_web_url` - (Computed, String) Specifies the web url for the Sharepoint site.
				* `source_id` - (Computed, Integer) Specifies registered source id to which object belongs.
				* `source_name` - (Computed, String) Specifies registered source name to which object belongs.
				* `uuid` - (Computed, String) Specifies the uuid which is a unique identifier of the object.
				* `v_center_summary` - (Optional, List)
				Nested schema for **v_center_summary**:
					* `is_cloud_env` - (Computed, Boolean) Specifies that registered vCenter source is a VMC (VMware Cloud) environment or not.
				* `windows_cluster_summary` - (Optional, List)
				Nested schema for **windows_cluster_summary**:
					* `cluster_source_type` - (Computed, String) Specifies the type of cluster resource this source represents.
			* `environment` - (Computed, String) Specifies the environment of the object.
			  * Constraints: Allowable values are: `kPhysical`, `kSQL`.
			* `global_id` - (Computed, String) Specifies the global id which is a unique identifier of the object.
			* `id` - (Computed, Integer) Specifies object id.
			* `logical_size_bytes` - (Computed, Integer) Specifies the logical size of object in bytes.
			* `name` - (Computed, String) Specifies the name of the object.
			* `object_hash` - (Computed, String) Specifies the hash identifier of the object.
			* `object_type` - (Computed, String) Specifies the type of the object.
			  * Constraints: Allowable values are: `kCluster`, `kVserver`, `kVolume`, `kVCenter`, `kStandaloneHost`, `kvCloudDirector`, `kFolder`, `kDatacenter`, `kComputeResource`, `kClusterComputeResource`, `kResourcePool`, `kDatastore`, `kHostSystem`, `kVirtualMachine`, `kVirtualApp`, `kStoragePod`, `kNetwork`, `kDistributedVirtualPortgroup`, `kTagCategory`, `kTag`, `kOpaqueNetwork`, `kOrganization`, `kVirtualDatacenter`, `kCatalog`, `kOrgMetadata`, `kStoragePolicy`, `kVirtualAppTemplate`, `kDomain`, `kOutlook`, `kMailbox`, `kUsers`, `kGroups`, `kSites`, `kUser`, `kGroup`, `kSite`, `kApplication`, `kGraphUser`, `kPublicFolders`, `kPublicFolder`, `kTeams`, `kTeam`, `kRootPublicFolder`, `kO365Exchange`, `kO365OneDrive`, `kO365Sharepoint`, `kKeyspace`, `kTable`, `kDatabase`, `kCollection`, `kBucket`, `kNamespace`, `kSCVMMServer`, `kStandaloneCluster`, `kHostGroup`, `kHypervHost`, `kHostCluster`, `kCustomProperty`, `kTenant`, `kSubscription`, `kResourceGroup`, `kStorageAccount`, `kStorageKey`, `kStorageContainer`, `kStorageBlob`, `kNetworkSecurityGroup`, `kVirtualNetwork`, `kSubnet`, `kComputeOptions`, `kSnapshotManagerPermit`, `kAvailabilitySet`, `kOVirtManager`, `kHost`, `kStorageDomain`, `kVNicProfile`, `kIAMUser`, `kRegion`, `kAvailabilityZone`, `kEC2Instance`, `kVPC`, `kInstanceType`, `kKeyPair`, `kRDSOptionGroup`, `kRDSParameterGroup`, `kRDSInstance`, `kRDSSubnet`, `kRDSTag`, `kAuroraTag`, `kAuroraCluster`, `kAccount`, `kSubTaskPermit`, `kS3Bucket`, `kS3Tag`, `kKmsKey`, `kProject`, `kLabel`, `kMetadata`, `kVPCConnector`, `kPrismCentral`, `kOtherHypervisorCluster`, `kZone`, `kMountPoint`, `kStorageArray`, `kFileSystem`, `kContainer`, `kFilesystem`, `kFileset`, `kPureProtectionGroup`, `kVolumeGroup`, `kStoragePool`, `kViewBox`, `kView`, `kWindowsCluster`, `kOracleRACCluster`, `kOracleAPCluster`, `kService`, `kPVC`, `kPersistentVolumeClaim`, `kPersistentVolume`, `kRootContainer`, `kDAGRootContainer`, `kExchangeNode`, `kExchangeDAGDatabaseCopy`, `kExchangeStandaloneDatabase`, `kExchangeDAG`, `kExchangeDAGDatabase`, `kDomainController`, `kInstance`, `kAAG`, `kAAGRootContainer`, `kAAGDatabase`, `kRACRootContainer`, `kTableSpace`, `kPDB`, `kObject`, `kOrg`, `kAppInstance`.
			* `os_type` - (Computed, String) Specifies the operating system type of the object.
			  * Constraints: Allowable values are: `kLinux`, `kWindows`.
			* `protection_type` - (Computed, String) Specifies the protection type of the object if any.
			  * Constraints: Allowable values are: `kAgent`, `kNative`, `kSnapshotManager`, `kRDSSnapshotManager`, `kAuroraSnapshotManager`, `kAwsS3`, `kAwsRDSPostgresBackup`, `kAwsAuroraPostgres`, `kAwsRDSPostgres`, `kAzureSQL`, `kFile`, `kVolume`.
			* `sharepoint_site_summary` - (Optional, List) Specifies the common parameters for Sharepoint site objects.
			Nested schema for **sharepoint_site_summary**:
				* `site_web_url` - (Computed, String) Specifies the web url for the Sharepoint site.
			* `source_id` - (Computed, Integer) Specifies registered source id to which object belongs.
			* `source_name` - (Computed, String) Specifies registered source name to which object belongs.
			* `uuid` - (Computed, String) Specifies the uuid which is a unique identifier of the object.
			* `v_center_summary` - (Optional, List)
			Nested schema for **v_center_summary**:
				* `is_cloud_env` - (Computed, Boolean) Specifies that registered vCenter source is a VMC (VMware Cloud) environment or not.
			* `windows_cluster_summary` - (Optional, List)
			Nested schema for **windows_cluster_summary**:
				* `cluster_source_type` - (Computed, String) Specifies the type of cluster resource this source represents.
		* `point_in_time_usecs` - (Optional, Integer) Specifies the timestamp (in microseconds. from epoch) for recovering to a point-in-time in the past.
		* `progress_task_id` - (Computed, String) Progress monitor task id for Recovery of VM.
		* `protection_group_id` - (Optional, String) Specifies the protection group id of the object snapshot.
		* `protection_group_name` - (Optional, String) Specifies the protection group name of the object snapshot.
		* `recover_from_standby` - (Optional, Boolean) Specifies that user wants to perform standby restore if it is enabled for this object.
		* `snapshot_creation_time_usecs` - (Computed, Integer) Specifies the time when the snapshot is created in Unix timestamp epoch in microseconds.
		* `snapshot_id` - (Required, String) Specifies the snapshot id.
		* `snapshot_target_type` - (Computed, String) Specifies the snapshot target type.
		  * Constraints: Allowable values are: `Local`, `Archival`, `RpaasArchival`, `StorageArraySnapshot`, `Remote`.
		* `start_time_usecs` - (Computed, Integer) Specifies the start time of the Recovery in Unix timestamp epoch in microseconds.
		* `status` - (Computed, String) Status of the Recovery. 'Running' indicates that the Recovery is still running. 'Canceled' indicates that the Recovery has been cancelled. 'Canceling' indicates that the Recovery is in the process of being cancelled. 'Failed' indicates that the Recovery has failed. 'Succeeded' indicates that the Recovery has finished successfully. 'SucceededWithWarning' indicates that the Recovery finished successfully, but there were some warning messages. 'Skipped' indicates that the Recovery task was skipped.
		  * Constraints: Allowable values are: `Accepted`, `Running`, `Canceled`, `Canceling`, `Failed`, `Missed`, `Succeeded`, `SucceededWithWarning`, `OnHold`, `Finalizing`, `Skipped`, `LegalHold`.
	* `recover_file_and_folder_params` - (Optional, List) Specifies the parameters to perform a file and folder recovery.
	Nested schema for **recover_file_and_folder_params**:
		* `files_and_folders` - (Required, List) Specifies the information about the files and folders to be recovered.
		Nested schema for **files_and_folders**:
			* `absolute_path` - (Required, String) Specifies the absolute path to the file or folder.
			* `destination_dir` - (Computed, String) Specifies the destination directory where the file/directory was copied.
			* `is_directory` - (Optional, Boolean) Specifies whether this is a directory or not.
			* `is_view_file_recovery` - (Optional, Boolean) Specify if the recovery is of type view file/folder.
			* `messages` - (Computed, List) Specify error messages about the file during recovery.
			* `status` - (Computed, String) Specifies the recovery status for this file or folder.
			  * Constraints: Allowable values are: `NotStarted`, `EstimationInProgress`, `EstimationDone`, `CopyInProgress`, `Finished`.
		* `kubernetes_target_params` - (Optional, List) Specifies the parameters to recover to a Kubernetes target.
		Nested schema for **kubernetes_target_params**:
			* `continue_on_error` - (Optional, Boolean) Specifies whether to continue recovering other files if one of files or folders failed to recover. Default value is false.
			* `new_target_config` - (Optional, List) Specifies the configuration for recovering to a new target.
			Nested schema for **new_target_config**:
				* `absolute_path` - (Required, String) Specifies the absolute path of the file.
				* `target_namespace` - (Optional, List) Specifies the target namespace to recover files and folders to.
				Nested schema for **target_namespace**:
					* `id` - (Required, Integer) Specifies the id of the object.
					* `name` - (Computed, String) Specifies the name of the object.
					* `parent_source_id` - (Computed, Integer) Specifies the id of the parent source of the target.
					* `parent_source_name` - (Computed, String) Specifies the name of the parent source of the target.
				* `target_pvc` - (Required, List) Specifies the target PVC(Persistent Volume Claim) to recover files and folders to.
				Nested schema for **target_pvc**:
					* `id` - (Required, Integer) Specifies the id of the object.
					* `name` - (Computed, String) Specifies the name of the object.
					* `parent_source_id` - (Computed, Integer) Specifies the id of the parent source of the target.
					* `parent_source_name` - (Computed, String) Specifies the name of the parent source of the target.
				* `target_source` - (Optional, List) Specifies the target kubernetes to recover files and folders to.
				Nested schema for **target_source**:
					* `id` - (Required, Integer) Specifies the id of the object.
					* `name` - (Computed, String) Specifies the name of the object.
					* `parent_source_id` - (Computed, Integer) Specifies the id of the parent source of the target.
					* `parent_source_name` - (Computed, String) Specifies the name of the parent source of the target.
			* `original_target_config` - (Optional, List) Specifies the configuration for recovering to the original target.
			Nested schema for **original_target_config**:
				* `alternate_path` - (Optional, String) Specifies the alternate path location to recover files to.
				* `recover_to_original_path` - (Required, Boolean) Specifies whether to recover files and folders to the original path location. If false, alternatePath must be specified.
			* `overwrite_existing` - (Optional, Boolean) Specifies whether to overwrite the existing files. Default is true.
			* `preserve_attributes` - (Optional, Boolean) Specifies whether to preserve original attributes. Default is true.
			* `recover_to_original_target` - (Required, Boolean) Specifies whether to recover to the original target. If true, originalTargetConfig must be specified. If false, newTargetConfig must be specified.
			* `vlan_config` - (Optional, List) Specifies VLAN Params associated with the recovered files and folders. If this is not specified, then the VLAN settings will be automatically selected from one of the below options: a. If VLANs are configured on Cohesity, then the VLAN host/VIP will be automatically based on the client's (e.g. ESXI host) IP address. b. If VLANs are not configured on Cohesity, then the partition hostname or VIPs will be used for Recovery.
			Nested schema for **vlan_config**:
				* `disable_vlan` - (Optional, Boolean) If this is set to true, then even if VLANs are configured on the system, the partition VIPs will be used for the Recovery.
				* `id` - (Optional, Integer) If this is set, then the Cohesity host name or the IP address associated with this vlan is used for mounting Cohesity's view on the remote host.
				* `interface_name` - (Computed, String) Interface group to use for Recovery.
		* `target_environment` - (Required, String) Specifies the environment of the recovery target. The corresponding params below must be filled out.
		  * Constraints: Allowable values are: `kKubernetes`.
	* `recover_namespace_params` - (Optional, List) Specifies the parameters to recover Kubernetes Namespaces.
	Nested schema for **recover_namespace_params**:
		* `kubernetes_target_params` - (Optional, List) Specifies the params for recovering to a Kubernetes host.
		Nested schema for **kubernetes_target_params**:
			* `exclude_params` - (Optional, List) Specifies the parameters to in/exclude objects (e.g.: volumes). An object satisfying any of these criteria will be included by this filter.
			Nested schema for **exclude_params**:
				* `label_combination_method` - (Optional, String) Whether to include all the labels or any of them while performing inclusion/exclusion of objects.
				  * Constraints: Allowable values are: `AND`, `OR`.
				* `label_vector` - (Optional, List) Array of Object to represent Label that Specify Objects (e.g.: Persistent Volumes and Persistent Volume Claims) to Include or Exclude.It will be a two-dimensional array, where each inner array will consist of a key and value representing labels. Using this two dimensional array of Labels, the Cluster generates a list of items to include in the filter, which are derived from intersections or the union of these labels, as decided by operation parameter.
				Nested schema for **label_vector**:
					* `key` - (Computed, String) The key of the label, used to identify the label.
					* `value` - (Computed, String) The value associated with the label key.
				* `objects` - (Optional, List) Array of objects that are to be included.
				* `selected_resources` - (Optional, List) Array of Object which has group, version, kind, etc. as its fields to identify a resource type and a resource list which is essentially the list of instances of that resource type.
				Nested schema for **selected_resources**:
					* `api_group` - (Optional, String) API group name of the resource (excluding the version). (Eg. apps, kubevirt.io).
					* `is_cluster_scoped` - (Optional, Boolean) Boolean indicating whether the resource is cluster scoped or not. This field is ignored for resource selection during recovery.
					* `kind` - (Optional, String) The kind of the resource type. (Eg. VirtualMachine).
					* `name` - (Optional, String) The name of the resource. This field is ignored for resource selection during recovery.
					* `resource_list` - (Optional, List) Array of the instances of the resource with group, version and kind mentioned above.
					Nested schema for **resource_list**:
						* `entity_id` - (Optional, Integer) The id of the specific entity to be backed up or restored.
						* `name` - (Optional, String) The name of the specific entity/resource to be backed up or restored.
					* `version` - (Optional, String) The version under the API group for the resource. This field is ignored for resource selection during recovery.
			* `excluded_pvcs` - (Optional, List) Specifies the list of pvc to be excluded from recovery. This will be deprecated in the future. This is overridden by the object level param.
			Nested schema for **excluded_pvcs**:
				* `id` - (Computed, Integer) Specifies the id of the pvc.
				* `name` - (Computed, String) Name of the pvc.
			* `include_params` - (Optional, List) Specifies the parameters to in/exclude objects (e.g.: volumes). An object satisfying any of these criteria will be included by this filter.
			Nested schema for **include_params**:
				* `label_combination_method` - (Optional, String) Whether to include all the labels or any of them while performing inclusion/exclusion of objects.
				  * Constraints: Allowable values are: `AND`, `OR`.
				* `label_vector` - (Optional, List) Array of Object to represent Label that Specify Objects (e.g.: Persistent Volumes and Persistent Volume Claims) to Include or Exclude.It will be a two-dimensional array, where each inner array will consist of a key and value representing labels. Using this two dimensional array of Labels, the Cluster generates a list of items to include in the filter, which are derived from intersections or the union of these labels, as decided by operation parameter.
				Nested schema for **label_vector**:
					* `key` - (Computed, String) The key of the label, used to identify the label.
					* `value` - (Computed, String) The value associated with the label key.
				* `objects` - (Optional, List) Array of objects that are to be included.
				* `selected_resources` - (Optional, List) Array of Object which has group, version, kind, etc. as its fields to identify a resource type and a resource list which is essentially the list of instances of that resource type.
				Nested schema for **selected_resources**:
					* `api_group` - (Optional, String) API group name of the resource (excluding the version). (Eg. apps, kubevirt.io).
					* `is_cluster_scoped` - (Optional, Boolean) Boolean indicating whether the resource is cluster scoped or not. This field is ignored for resource selection during recovery.
					* `kind` - (Optional, String) The kind of the resource type. (Eg. VirtualMachine).
					* `name` - (Optional, String) The name of the resource. This field is ignored for resource selection during recovery.
					* `resource_list` - (Optional, List) Array of the instances of the resource with group, version and kind mentioned above.
					Nested schema for **resource_list**:
						* `entity_id` - (Optional, Integer) The id of the specific entity to be backed up or restored.
						* `name` - (Optional, String) The name of the specific entity/resource to be backed up or restored.
					* `version` - (Optional, String) The version under the API group for the resource. This field is ignored for resource selection during recovery.
			* `objects` - (Optional, List) Specifies the objects to be recovered.
			Nested schema for **objects**:
				* `archival_target_info` - (Optional, List) Specifies the archival target information if the snapshot is an archival snapshot.
				Nested schema for **archival_target_info**:
					* `archival_task_id` - (Computed, String) Specifies the archival task id. This is a protection group UID which only applies when archival type is 'Tape'.
					* `ownership_context` - (Computed, String) Specifies the ownership context for the target.
					  * Constraints: Allowable values are: `Local`, `FortKnox`.
					* `target_id` - (Computed, Integer) Specifies the archival target ID.
					* `target_name` - (Computed, String) Specifies the archival target name.
					* `target_type` - (Computed, String) Specifies the archival target type.
					  * Constraints: Allowable values are: `Tape`, `Cloud`, `Nas`.
					* `tier_settings` - (Optional, List) Specifies the tier info for archival.
					Nested schema for **tier_settings**:
						* `aws_tiering` - (Optional, List) Specifies aws tiers.
						Nested schema for **aws_tiering**:
							* `tiers` - (Required, List) Specifies the tiers that are used to move the archived backup from current tier to next tier. The order of the tiers determines which tier will be used next for moving the archived backup. The first tier input should always be default tier where backup will be acrhived. Each tier specifies how much time after the backup will be moved to next tier from the current tier.
							Nested schema for **tiers**:
								* `move_after` - (Computed, Integer) Specifies the time period after which the backup will be moved from current tier to next tier.
								* `move_after_unit` - (Computed, String) Specifies the unit for moving the data from current tier to next tier. This unit will be a base unit for the 'moveAfter' field specified below.
								  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
								* `tier_type` - (Computed, String) Specifies the AWS tier types.
								  * Constraints: Allowable values are: `kAmazonS3Standard`, `kAmazonS3StandardIA`, `kAmazonS3OneZoneIA`, `kAmazonS3IntelligentTiering`, `kAmazonS3Glacier`, `kAmazonS3GlacierDeepArchive`.
						* `azure_tiering` - (Optional, List) Specifies Azure tiers.
						Nested schema for **azure_tiering**:
							* `tiers` - (Optional, List) Specifies the tiers that are used to move the archived backup from current tier to next tier. The order of the tiers determines which tier will be used next for moving the archived backup. The first tier input should always be default tier where backup will be acrhived. Each tier specifies how much time after the backup will be moved to next tier from the current tier.
							Nested schema for **tiers**:
								* `move_after` - (Computed, Integer) Specifies the time period after which the backup will be moved from current tier to next tier.
								* `move_after_unit` - (Computed, String) Specifies the unit for moving the data from current tier to next tier. This unit will be a base unit for the 'moveAfter' field specified below.
								  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
								* `tier_type` - (Computed, String) Specifies the Azure tier types.
								  * Constraints: Allowable values are: `kAzureTierHot`, `kAzureTierCool`, `kAzureTierArchive`.
						* `cloud_platform` - (Computed, String) Specifies the cloud platform to enable tiering.
						  * Constraints: Allowable values are: `AWS`, `Azure`, `Oracle`, `Google`.
						* `current_tier_type` - (Computed, String) Specifies the type of the current tier where the snapshot resides. This will be specified if the run is a CAD run.
						  * Constraints: Allowable values are: `kAmazonS3Standard`, `kAmazonS3StandardIA`, `kAmazonS3OneZoneIA`, `kAmazonS3IntelligentTiering`, `kAmazonS3Glacier`, `kAmazonS3GlacierDeepArchive`, `kAzureTierHot`, `kAzureTierCool`, `kAzureTierArchive`, `kGoogleStandard`, `kGoogleRegional`, `kGoogleMultiRegional`, `kGoogleNearline`, `kGoogleColdline`, `kOracleTierStandard`, `kOracleTierArchive`.
						* `google_tiering` - (Optional, List) Specifies Google tiers.
						Nested schema for **google_tiering**:
							* `tiers` - (Required, List) Specifies the tiers that are used to move the archived backup from current tier to next tier. The order of the tiers determines which tier will be used next for moving the archived backup. The first tier input should always be default tier where backup will be acrhived. Each tier specifies how much time after the backup will be moved to next tier from the current tier.
							Nested schema for **tiers**:
								* `move_after` - (Computed, Integer) Specifies the time period after which the backup will be moved from current tier to next tier.
								* `move_after_unit` - (Computed, String) Specifies the unit for moving the data from current tier to next tier. This unit will be a base unit for the 'moveAfter' field specified below.
								  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
								* `tier_type` - (Computed, String) Specifies the Google tier types.
								  * Constraints: Allowable values are: `kGoogleStandard`, `kGoogleRegional`, `kGoogleMultiRegional`, `kGoogleNearline`, `kGoogleColdline`.
						* `oracle_tiering` - (Optional, List) Specifies Oracle tiers.
						Nested schema for **oracle_tiering**:
							* `tiers` - (Required, List) Specifies the tiers that are used to move the archived backup from current tier to next tier. The order of the tiers determines which tier will be used next for moving the archived backup. The first tier input should always be default tier where backup will be acrhived. Each tier specifies how much time after the backup will be moved to next tier from the current tier.
							Nested schema for **tiers**:
								* `move_after` - (Computed, Integer) Specifies the time period after which the backup will be moved from current tier to next tier.
								* `move_after_unit` - (Computed, String) Specifies the unit for moving the data from current tier to next tier. This unit will be a base unit for the 'moveAfter' field specified below.
								  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
								* `tier_type` - (Computed, String) Specifies the Oracle tier types.
								  * Constraints: Allowable values are: `kOracleTierStandard`, `kOracleTierArchive`.
					* `usage_type` - (Computed, String) Specifies the usage type for the target.
					  * Constraints: Allowable values are: `Archival`, `Tiering`, `Rpaas`.
				* `bytes_restored` - (Computed, Integer) Specify the total bytes restored.
				* `end_time_usecs` - (Computed, Integer) Specifies the end time of the Recovery in Unix timestamp epoch in microseconds. This field will be populated only after Recovery is finished.
				* `exclude_params` - (Optional, List) Specifies the parameters to in/exclude objects (e.g.: volumes). An object satisfying any of these criteria will be included by this filter.
				Nested schema for **exclude_params**:
					* `label_combination_method` - (Optional, String) Whether to include all the labels or any of them while performing inclusion/exclusion of objects.
					  * Constraints: Allowable values are: `AND`, `OR`.
					* `label_vector` - (Optional, List) Array of Object to represent Label that Specify Objects (e.g.: Persistent Volumes and Persistent Volume Claims) to Include or Exclude.It will be a two-dimensional array, where each inner array will consist of a key and value representing labels. Using this two dimensional array of Labels, the Cluster generates a list of items to include in the filter, which are derived from intersections or the union of these labels, as decided by operation parameter.
					Nested schema for **label_vector**:
						* `key` - (Computed, String) The key of the label, used to identify the label.
						* `value` - (Computed, String) The value associated with the label key.
					* `objects` - (Optional, List) Array of objects that are to be included.
					* `selected_resources` - (Optional, List) Array of Object which has group, version, kind, etc. as its fields to identify a resource type and a resource list which is essentially the list of instances of that resource type.
					Nested schema for **selected_resources**:
						* `api_group` - (Optional, String) API group name of the resource (excluding the version). (Eg. apps, kubevirt.io).
						* `is_cluster_scoped` - (Optional, Boolean) Boolean indicating whether the resource is cluster scoped or not. This field is ignored for resource selection during recovery.
						* `kind` - (Optional, String) The kind of the resource type. (Eg. VirtualMachine).
						* `name` - (Optional, String) The name of the resource. This field is ignored for resource selection during recovery.
						* `resource_list` - (Optional, List) Array of the instances of the resource with group, version and kind mentioned above.
						Nested schema for **resource_list**:
							* `entity_id` - (Optional, Integer) The id of the specific entity to be backed up or restored.
							* `name` - (Optional, String) The name of the specific entity/resource to be backed up or restored.
						* `version` - (Optional, String) The version under the API group for the resource. This field is ignored for resource selection during recovery.
				* `include_params` - (Optional, List) Specifies the parameters to in/exclude objects (e.g.: volumes). An object satisfying any of these criteria will be included by this filter.
				Nested schema for **include_params**:
					* `label_combination_method` - (Optional, String) Whether to include all the labels or any of them while performing inclusion/exclusion of objects.
					  * Constraints: Allowable values are: `AND`, `OR`.
					* `label_vector` - (Optional, List) Array of Object to represent Label that Specify Objects (e.g.: Persistent Volumes and Persistent Volume Claims) to Include or Exclude.It will be a two-dimensional array, where each inner array will consist of a key and value representing labels. Using this two dimensional array of Labels, the Cluster generates a list of items to include in the filter, which are derived from intersections or the union of these labels, as decided by operation parameter.
					Nested schema for **label_vector**:
						* `key` - (Computed, String) The key of the label, used to identify the label.
						* `value` - (Computed, String) The value associated with the label key.
					* `objects` - (Optional, List) Array of objects that are to be included.
					* `selected_resources` - (Optional, List) Array of Object which has group, version, kind, etc. as its fields to identify a resource type and a resource list which is essentially the list of instances of that resource type.
					Nested schema for **selected_resources**:
						* `api_group` - (Optional, String) API group name of the resource (excluding the version). (Eg. apps, kubevirt.io).
						* `is_cluster_scoped` - (Optional, Boolean) Boolean indicating whether the resource is cluster scoped or not. This field is ignored for resource selection during recovery.
						* `kind` - (Optional, String) The kind of the resource type. (Eg. VirtualMachine).
						* `name` - (Optional, String) The name of the resource. This field is ignored for resource selection during recovery.
						* `resource_list` - (Optional, List) Array of the instances of the resource with group, version and kind mentioned above.
						Nested schema for **resource_list**:
							* `entity_id` - (Optional, Integer) The id of the specific entity to be backed up or restored.
							* `name` - (Optional, String) The name of the specific entity/resource to be backed up or restored.
						* `version` - (Optional, String) The version under the API group for the resource. This field is ignored for resource selection during recovery.
				* `messages` - (Computed, List) Specify error messages about the object.
				* `object_info` - (Optional, List) Specifies the information about the object for which the snapshot is taken.
				Nested schema for **object_info**:
					* `child_objects` - (Optional, List) Specifies child object details.
					Nested schema for **child_objects**:
						* `child_objects` - (Optional, List) Specifies child object details.
						Nested schema for **child_objects**:
						* `environment` - (Computed, String) Specifies the environment of the object.
						  * Constraints: Allowable values are: `kPhysical`, `kSQL`.
						* `global_id` - (Computed, String) Specifies the global id which is a unique identifier of the object.
						* `id` - (Computed, Integer) Specifies object id.
						* `logical_size_bytes` - (Computed, Integer) Specifies the logical size of object in bytes.
						* `name` - (Computed, String) Specifies the name of the object.
						* `object_hash` - (Computed, String) Specifies the hash identifier of the object.
						* `object_type` - (Computed, String) Specifies the type of the object.
						  * Constraints: Allowable values are: `kCluster`, `kVserver`, `kVolume`, `kVCenter`, `kStandaloneHost`, `kvCloudDirector`, `kFolder`, `kDatacenter`, `kComputeResource`, `kClusterComputeResource`, `kResourcePool`, `kDatastore`, `kHostSystem`, `kVirtualMachine`, `kVirtualApp`, `kStoragePod`, `kNetwork`, `kDistributedVirtualPortgroup`, `kTagCategory`, `kTag`, `kOpaqueNetwork`, `kOrganization`, `kVirtualDatacenter`, `kCatalog`, `kOrgMetadata`, `kStoragePolicy`, `kVirtualAppTemplate`, `kDomain`, `kOutlook`, `kMailbox`, `kUsers`, `kGroups`, `kSites`, `kUser`, `kGroup`, `kSite`, `kApplication`, `kGraphUser`, `kPublicFolders`, `kPublicFolder`, `kTeams`, `kTeam`, `kRootPublicFolder`, `kO365Exchange`, `kO365OneDrive`, `kO365Sharepoint`, `kKeyspace`, `kTable`, `kDatabase`, `kCollection`, `kBucket`, `kNamespace`, `kSCVMMServer`, `kStandaloneCluster`, `kHostGroup`, `kHypervHost`, `kHostCluster`, `kCustomProperty`, `kTenant`, `kSubscription`, `kResourceGroup`, `kStorageAccount`, `kStorageKey`, `kStorageContainer`, `kStorageBlob`, `kNetworkSecurityGroup`, `kVirtualNetwork`, `kSubnet`, `kComputeOptions`, `kSnapshotManagerPermit`, `kAvailabilitySet`, `kOVirtManager`, `kHost`, `kStorageDomain`, `kVNicProfile`, `kIAMUser`, `kRegion`, `kAvailabilityZone`, `kEC2Instance`, `kVPC`, `kInstanceType`, `kKeyPair`, `kRDSOptionGroup`, `kRDSParameterGroup`, `kRDSInstance`, `kRDSSubnet`, `kRDSTag`, `kAuroraTag`, `kAuroraCluster`, `kAccount`, `kSubTaskPermit`, `kS3Bucket`, `kS3Tag`, `kKmsKey`, `kProject`, `kLabel`, `kMetadata`, `kVPCConnector`, `kPrismCentral`, `kOtherHypervisorCluster`, `kZone`, `kMountPoint`, `kStorageArray`, `kFileSystem`, `kContainer`, `kFilesystem`, `kFileset`, `kPureProtectionGroup`, `kVolumeGroup`, `kStoragePool`, `kViewBox`, `kView`, `kWindowsCluster`, `kOracleRACCluster`, `kOracleAPCluster`, `kService`, `kPVC`, `kPersistentVolumeClaim`, `kPersistentVolume`, `kRootContainer`, `kDAGRootContainer`, `kExchangeNode`, `kExchangeDAGDatabaseCopy`, `kExchangeStandaloneDatabase`, `kExchangeDAG`, `kExchangeDAGDatabase`, `kDomainController`, `kInstance`, `kAAG`, `kAAGRootContainer`, `kAAGDatabase`, `kRACRootContainer`, `kTableSpace`, `kPDB`, `kObject`, `kOrg`, `kAppInstance`.
						* `os_type` - (Computed, String) Specifies the operating system type of the object.
						  * Constraints: Allowable values are: `kLinux`, `kWindows`.
						* `protection_type` - (Computed, String) Specifies the protection type of the object if any.
						  * Constraints: Allowable values are: `kAgent`, `kNative`, `kSnapshotManager`, `kRDSSnapshotManager`, `kAuroraSnapshotManager`, `kAwsS3`, `kAwsRDSPostgresBackup`, `kAwsAuroraPostgres`, `kAwsRDSPostgres`, `kAzureSQL`, `kFile`, `kVolume`.
						* `sharepoint_site_summary` - (Optional, List) Specifies the common parameters for Sharepoint site objects.
						Nested schema for **sharepoint_site_summary**:
							* `site_web_url` - (Computed, String) Specifies the web url for the Sharepoint site.
						* `source_id` - (Computed, Integer) Specifies registered source id to which object belongs.
						* `source_name` - (Computed, String) Specifies registered source name to which object belongs.
						* `uuid` - (Computed, String) Specifies the uuid which is a unique identifier of the object.
						* `v_center_summary` - (Optional, List)
						Nested schema for **v_center_summary**:
							* `is_cloud_env` - (Computed, Boolean) Specifies that registered vCenter source is a VMC (VMware Cloud) environment or not.
						* `windows_cluster_summary` - (Optional, List)
						Nested schema for **windows_cluster_summary**:
							* `cluster_source_type` - (Computed, String) Specifies the type of cluster resource this source represents.
					* `environment` - (Computed, String) Specifies the environment of the object.
					  * Constraints: Allowable values are: `kPhysical`, `kSQL`.
					* `global_id` - (Computed, String) Specifies the global id which is a unique identifier of the object.
					* `id` - (Computed, Integer) Specifies object id.
					* `logical_size_bytes` - (Computed, Integer) Specifies the logical size of object in bytes.
					* `name` - (Computed, String) Specifies the name of the object.
					* `object_hash` - (Computed, String) Specifies the hash identifier of the object.
					* `object_type` - (Computed, String) Specifies the type of the object.
					  * Constraints: Allowable values are: `kCluster`, `kVserver`, `kVolume`, `kVCenter`, `kStandaloneHost`, `kvCloudDirector`, `kFolder`, `kDatacenter`, `kComputeResource`, `kClusterComputeResource`, `kResourcePool`, `kDatastore`, `kHostSystem`, `kVirtualMachine`, `kVirtualApp`, `kStoragePod`, `kNetwork`, `kDistributedVirtualPortgroup`, `kTagCategory`, `kTag`, `kOpaqueNetwork`, `kOrganization`, `kVirtualDatacenter`, `kCatalog`, `kOrgMetadata`, `kStoragePolicy`, `kVirtualAppTemplate`, `kDomain`, `kOutlook`, `kMailbox`, `kUsers`, `kGroups`, `kSites`, `kUser`, `kGroup`, `kSite`, `kApplication`, `kGraphUser`, `kPublicFolders`, `kPublicFolder`, `kTeams`, `kTeam`, `kRootPublicFolder`, `kO365Exchange`, `kO365OneDrive`, `kO365Sharepoint`, `kKeyspace`, `kTable`, `kDatabase`, `kCollection`, `kBucket`, `kNamespace`, `kSCVMMServer`, `kStandaloneCluster`, `kHostGroup`, `kHypervHost`, `kHostCluster`, `kCustomProperty`, `kTenant`, `kSubscription`, `kResourceGroup`, `kStorageAccount`, `kStorageKey`, `kStorageContainer`, `kStorageBlob`, `kNetworkSecurityGroup`, `kVirtualNetwork`, `kSubnet`, `kComputeOptions`, `kSnapshotManagerPermit`, `kAvailabilitySet`, `kOVirtManager`, `kHost`, `kStorageDomain`, `kVNicProfile`, `kIAMUser`, `kRegion`, `kAvailabilityZone`, `kEC2Instance`, `kVPC`, `kInstanceType`, `kKeyPair`, `kRDSOptionGroup`, `kRDSParameterGroup`, `kRDSInstance`, `kRDSSubnet`, `kRDSTag`, `kAuroraTag`, `kAuroraCluster`, `kAccount`, `kSubTaskPermit`, `kS3Bucket`, `kS3Tag`, `kKmsKey`, `kProject`, `kLabel`, `kMetadata`, `kVPCConnector`, `kPrismCentral`, `kOtherHypervisorCluster`, `kZone`, `kMountPoint`, `kStorageArray`, `kFileSystem`, `kContainer`, `kFilesystem`, `kFileset`, `kPureProtectionGroup`, `kVolumeGroup`, `kStoragePool`, `kViewBox`, `kView`, `kWindowsCluster`, `kOracleRACCluster`, `kOracleAPCluster`, `kService`, `kPVC`, `kPersistentVolumeClaim`, `kPersistentVolume`, `kRootContainer`, `kDAGRootContainer`, `kExchangeNode`, `kExchangeDAGDatabaseCopy`, `kExchangeStandaloneDatabase`, `kExchangeDAG`, `kExchangeDAGDatabase`, `kDomainController`, `kInstance`, `kAAG`, `kAAGRootContainer`, `kAAGDatabase`, `kRACRootContainer`, `kTableSpace`, `kPDB`, `kObject`, `kOrg`, `kAppInstance`.
					* `os_type` - (Computed, String) Specifies the operating system type of the object.
					  * Constraints: Allowable values are: `kLinux`, `kWindows`.
					* `protection_type` - (Computed, String) Specifies the protection type of the object if any.
					  * Constraints: Allowable values are: `kAgent`, `kNative`, `kSnapshotManager`, `kRDSSnapshotManager`, `kAuroraSnapshotManager`, `kAwsS3`, `kAwsRDSPostgresBackup`, `kAwsAuroraPostgres`, `kAwsRDSPostgres`, `kAzureSQL`, `kFile`, `kVolume`.
					* `sharepoint_site_summary` - (Optional, List) Specifies the common parameters for Sharepoint site objects.
					Nested schema for **sharepoint_site_summary**:
						* `site_web_url` - (Computed, String) Specifies the web url for the Sharepoint site.
					* `source_id` - (Computed, Integer) Specifies registered source id to which object belongs.
					* `source_name` - (Computed, String) Specifies registered source name to which object belongs.
					* `uuid` - (Computed, String) Specifies the uuid which is a unique identifier of the object.
					* `v_center_summary` - (Optional, List)
					Nested schema for **v_center_summary**:
						* `is_cloud_env` - (Computed, Boolean) Specifies that registered vCenter source is a VMC (VMware Cloud) environment or not.
					* `windows_cluster_summary` - (Optional, List)
					Nested schema for **windows_cluster_summary**:
						* `cluster_source_type` - (Computed, String) Specifies the type of cluster resource this source represents.
				* `point_in_time_usecs` - (Optional, Integer) Specifies the timestamp (in microseconds. from epoch) for recovering to a point-in-time in the past.
				* `progress_task_id` - (Computed, String) Progress monitor task id for Recovery of VM.
				* `protection_group_id` - (Optional, String) Specifies the protection group id of the object snapshot.
				* `protection_group_name` - (Optional, String) Specifies the protection group name of the object snapshot.
				* `recover_from_standby` - (Optional, Boolean) Specifies that user wants to perform standby restore if it is enabled for this object.
				* `recover_pvcs_only` - (Optional, Boolean) Specifies whether to recover PVCs only during recovery. Default: false.
				* `snapshot_creation_time_usecs` - (Computed, Integer) Specifies the time when the snapshot is created in Unix timestamp epoch in microseconds.
				* `snapshot_id` - (Required, String) Specifies the snapshot id.
				* `snapshot_target_type` - (Computed, String) Specifies the snapshot target type.
				  * Constraints: Allowable values are: `Local`, `Archival`, `RpaasArchival`, `StorageArraySnapshot`, `Remote`.
				* `start_time_usecs` - (Computed, Integer) Specifies the start time of the Recovery in Unix timestamp epoch in microseconds.
				* `status` - (Computed, String) Status of the Recovery. 'Running' indicates that the Recovery is still running. 'Canceled' indicates that the Recovery has been cancelled. 'Canceling' indicates that the Recovery is in the process of being cancelled. 'Failed' indicates that the Recovery has failed. 'Succeeded' indicates that the Recovery has finished successfully. 'SucceededWithWarning' indicates that the Recovery finished successfully, but there were some warning messages. 'Skipped' indicates that the Recovery task was skipped.
				  * Constraints: Allowable values are: `Accepted`, `Running`, `Canceled`, `Canceling`, `Failed`, `Missed`, `Succeeded`, `SucceededWithWarning`, `OnHold`, `Finalizing`, `Skipped`, `LegalHold`.
			* `storage_class` - (Optional, List) Specifies the storage class parameters for recovery of namespace.
				Nested schema for **storage_class**:
					* `storage_class_mapping` - (Optional, List) Specifies mapping of storage classes.
					Nested schema for **storage_class_mapping**:
						* `key` - (Computed, String) The key of the label, used to identify the label.
						* `value` - (Computed, String) The value associated with the label key.
					* `use_storage_class_mapping` - (Optional, Boolean) Specifies whether or not to use storage class mapping.
				* `unbind_pvcs` - (Optional, Boolean) Specifies whether the volume bindings will be removed from all restored PVCs. This will effectively unbind the PVCs from their original PVs. Default: false.
			* `recover_cluster_scoped_resources` - (Optional, List) Specifies the parameters from where the cluster scoped resources would be recovered.
			Nested schema for **recover_cluster_scoped_resources**:
				* `snapshot_id` - (Optional, String) Specifies the snapshot id of the namespace from where the cluster scoped resources are to be recovered.
			* `recover_protection_group_runs_params` - (Optional, List) Specifies the Protection Group Runs params to recover. All the VM's that are successfully backed up by specified Runs will be recovered. This can be specified along with individual snapshots of VMs. User has to make sure that specified Object snapshots and Protection Group Runs should not have any intersection. For example, user cannot specify multiple Runs which has same Object or an Object snapshot and a Run which has same Object's snapshot.
			Nested schema for **recover_protection_group_runs_params**:
				* `archival_target_id` - (Optional, Integer) Specifies the archival target id. If specified and Protection Group run has an archival snapshot then VMs are recovered from the specified archival snapshot. If not specified (default), VMs are recovered from local snapshot.
				* `protection_group_id` - (Optional, String) Specifies the local Protection Group id. In case of recovering a replication Run, this field should be provided with local Protection Group id.
				* `protection_group_instance_id` - (Required, Integer) Specifies the Protection Group Instance id.
				* `protection_group_run_id` - (Required, String) Specifies the Protection Group Run id from which to recover VMs. All the VM's that are successfully protected by this Run will be recovered.
				  * Constraints: The value must match regular expression `/^\\d+:\\d+$/`.
			* `recover_pvcs_only` - (Optional, Boolean) Specifies whether to recover PVCs only during recovery.. This is overridden with the object level settings and will be deprecated in the future.
			* `recovery_region_migration_params` - (Optional, List) Specifies an individual migration rule for mapping a region/zone to another for cross region recovery.
			Nested schema for **recovery_region_migration_params**:
				* `current_value` - (Required, String) Specifies the current value for the mapping that needs to be mutated.
				* `new_value` - (Required, String) Specifies the new value for the mapping with which the fields need to be updated with.
			* `recovery_target_config` - (Required, List) Specifies the recovery target configuration of the Namespace recovery.
			Nested schema for **recovery_target_config**:
				* `new_source_config` - (Optional, List) Specifies the new source configuration if a Kubernetes Namespace is being restored to a different source than the one from which it was protected.
				Nested schema for **new_source_config**:
					* `source` - (Required, List) Specifies the id of the parent source to recover the Namespaces.
					Nested schema for **source**:
						* `id` - (Required, Integer) Specifies the id of the object.
						* `name` - (Computed, String) Specifies the name of the object.
				* `recover_to_new_source` - (Required, Boolean) Specifies whether or not to recover the Namespaces to a different source than they were backed up from.
			* `recovery_zone_migration_params` - (Optional, List) Specifies rules for performing zone migrations during recovery. Used in case of recovery to new location and the namespace being recovered is in a different zone.
			Nested schema for **recovery_zone_migration_params**:
				* `current_value` - (Required, String) Specifies the current value for the mapping that needs to be mutated.
				* `new_value` - (Required, String) Specifies the new value for the mapping with which the fields need to be updated with.
			* `rename_recovered_namespaces_params` - (Optional, List) Specifies params to rename the Namespaces that are recovered. If not specified, the original names of the Namespaces are preserved. If a name collision occurs then the Namespace being recovered will overwrite the Namespace already present on the source.
			Nested schema for **rename_recovered_namespaces_params**:
				* `prefix` - (Optional, String) Specifies the prefix string to be added to recovered or cloned object names.
				* `suffix` - (Optional, String) Specifies the suffix string to be added to recovered or cloned object names.
			* `skip_cluster_compatibility_check` - (Optional, Boolean) Specifies whether to skip checking if the target cluster, to restore to, is compatible or not. By default restore allowed to compatible cluster only.
			* `storage_class` - (Optional, List) Specifies the storage class parameters for recovery of namespace.
			Nested schema for **storage_class**:
				* `storage_class_mapping` - (Optional, List) Specifies mapping of storage classes.
				Nested schema for **storage_class_mapping**:
					* `key` - (Computed, String) The key of the label, used to identify the label.
					* `value` - (Computed, String) The value associated with the label key.
				* `use_storage_class_mapping` - (Optional, Boolean) Specifies whether or not to use storage class mapping.
		* `target_environment` - (Required, String) Specifies the environment of the recovery target. The corresponding params below must be filled out. As of now only kubernetes target environment is supported.
		  * Constraints: Allowable values are: `kKubernetes`.
		* `vlan_config` - (Optional, List) Specifies VLAN Params associated with the recovered. If this is not specified, then the VLAN settings will be automatically selected from one of the below options: a. If VLANs are configured on Cohesity, then the VLAN host/VIP will be automatically based on the client's (e.g. ESXI host) IP address. b. If VLANs are not configured on Cohesity, then the partition hostname or VIPs will be used for Recovery.
		Nested schema for **vlan_config**:
			* `disable_vlan` - (Optional, Boolean) If this is set to true, then even if VLANs are configured on the system, the partition VIPs will be used for the Recovery.
			* `id` - (Optional, Integer) If this is set, then the Cohesity host name or the IP address associated with this vlan is used for mounting Cohesity's view on the remote host.
			* `interface_name` - (Computed, String) Interface group to use for Recovery.
	* `recovery_action` - (Required, String) Specifies the type of recover action to be performed.
	  * Constraints: Allowable values are: `RecoverNamespaces`, `RecoverFiles`, `DownloadFilesAndFolders`.
* `mssql_params` - (Optional, Forces new resource, List) Specifies the recovery options specific to Sql environment.
Nested schema for **mssql_params**:
	* `recover_app_params` - (Optional, List) Specifies the parameters to recover Sql databases.
	  * Constraints: The minimum length is `1` item.
	Nested schema for **recover_app_params**:
		* `aag_info` - (Optional, List) Object details for Mssql.
		Nested schema for **aag_info**:
			* `name` - (Optional, String) Specifies the AAG name.
			* `object_id` - (Optional, Integer) Specifies the AAG object Id.
		* `archival_target_info` - (Optional, List) Specifies the archival target information if the snapshot is an archival snapshot.
		Nested schema for **archival_target_info**:
			* `archival_task_id` - (Computed, String) Specifies the archival task id. This is a protection group UID which only applies when archival type is 'Tape'.
			* `ownership_context` - (Computed, String) Specifies the ownership context for the target.
			  * Constraints: Allowable values are: `Local`, `FortKnox`.
			* `target_id` - (Computed, Integer) Specifies the archival target ID.
			* `target_name` - (Computed, String) Specifies the archival target name.
			* `target_type` - (Computed, String) Specifies the archival target type.
			  * Constraints: Allowable values are: `Tape`, `Cloud`, `Nas`.
			* `tier_settings` - (Optional, List) Specifies the tier info for archival.
			Nested schema for **tier_settings**:
				* `aws_tiering` - (Optional, List) Specifies aws tiers.
				Nested schema for **aws_tiering**:
					* `tiers` - (Required, List) Specifies the tiers that are used to move the archived backup from current tier to next tier. The order of the tiers determines which tier will be used next for moving the archived backup. The first tier input should always be default tier where backup will be acrhived. Each tier specifies how much time after the backup will be moved to next tier from the current tier.
					Nested schema for **tiers**:
						* `move_after` - (Computed, Integer) Specifies the time period after which the backup will be moved from current tier to next tier.
						* `move_after_unit` - (Computed, String) Specifies the unit for moving the data from current tier to next tier. This unit will be a base unit for the 'moveAfter' field specified below.
						  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
						* `tier_type` - (Computed, String) Specifies the AWS tier types.
						  * Constraints: Allowable values are: `kAmazonS3Standard`, `kAmazonS3StandardIA`, `kAmazonS3OneZoneIA`, `kAmazonS3IntelligentTiering`, `kAmazonS3Glacier`, `kAmazonS3GlacierDeepArchive`.
				* `azure_tiering` - (Optional, List) Specifies Azure tiers.
				Nested schema for **azure_tiering**:
					* `tiers` - (Optional, List) Specifies the tiers that are used to move the archived backup from current tier to next tier. The order of the tiers determines which tier will be used next for moving the archived backup. The first tier input should always be default tier where backup will be acrhived. Each tier specifies how much time after the backup will be moved to next tier from the current tier.
					Nested schema for **tiers**:
						* `move_after` - (Computed, Integer) Specifies the time period after which the backup will be moved from current tier to next tier.
						* `move_after_unit` - (Computed, String) Specifies the unit for moving the data from current tier to next tier. This unit will be a base unit for the 'moveAfter' field specified below.
						  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
						* `tier_type` - (Computed, String) Specifies the Azure tier types.
						  * Constraints: Allowable values are: `kAzureTierHot`, `kAzureTierCool`, `kAzureTierArchive`.
				* `cloud_platform` - (Computed, String) Specifies the cloud platform to enable tiering.
				  * Constraints: Allowable values are: `AWS`, `Azure`, `Oracle`, `Google`.
				* `current_tier_type` - (Computed, String) Specifies the type of the current tier where the snapshot resides. This will be specified if the run is a CAD run.
				  * Constraints: Allowable values are: `kAmazonS3Standard`, `kAmazonS3StandardIA`, `kAmazonS3OneZoneIA`, `kAmazonS3IntelligentTiering`, `kAmazonS3Glacier`, `kAmazonS3GlacierDeepArchive`, `kAzureTierHot`, `kAzureTierCool`, `kAzureTierArchive`, `kGoogleStandard`, `kGoogleRegional`, `kGoogleMultiRegional`, `kGoogleNearline`, `kGoogleColdline`, `kOracleTierStandard`, `kOracleTierArchive`.
				* `google_tiering` - (Optional, List) Specifies Google tiers.
				Nested schema for **google_tiering**:
					* `tiers` - (Required, List) Specifies the tiers that are used to move the archived backup from current tier to next tier. The order of the tiers determines which tier will be used next for moving the archived backup. The first tier input should always be default tier where backup will be acrhived. Each tier specifies how much time after the backup will be moved to next tier from the current tier.
					Nested schema for **tiers**:
						* `move_after` - (Computed, Integer) Specifies the time period after which the backup will be moved from current tier to next tier.
						* `move_after_unit` - (Computed, String) Specifies the unit for moving the data from current tier to next tier. This unit will be a base unit for the 'moveAfter' field specified below.
						  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
						* `tier_type` - (Computed, String) Specifies the Google tier types.
						  * Constraints: Allowable values are: `kGoogleStandard`, `kGoogleRegional`, `kGoogleMultiRegional`, `kGoogleNearline`, `kGoogleColdline`.
				* `oracle_tiering` - (Optional, List) Specifies Oracle tiers.
				Nested schema for **oracle_tiering**:
					* `tiers` - (Required, List) Specifies the tiers that are used to move the archived backup from current tier to next tier. The order of the tiers determines which tier will be used next for moving the archived backup. The first tier input should always be default tier where backup will be acrhived. Each tier specifies how much time after the backup will be moved to next tier from the current tier.
					Nested schema for **tiers**:
						* `move_after` - (Computed, Integer) Specifies the time period after which the backup will be moved from current tier to next tier.
						* `move_after_unit` - (Computed, String) Specifies the unit for moving the data from current tier to next tier. This unit will be a base unit for the 'moveAfter' field specified below.
						  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
						* `tier_type` - (Computed, String) Specifies the Oracle tier types.
						  * Constraints: Allowable values are: `kOracleTierStandard`, `kOracleTierArchive`.
			* `usage_type` - (Computed, String) Specifies the usage type for the target.
			  * Constraints: Allowable values are: `Archival`, `Tiering`, `Rpaas`.
		* `bytes_restored` - (Computed, Integer) Specify the total bytes restored.
		* `end_time_usecs` - (Computed, Integer) Specifies the end time of the Recovery in Unix timestamp epoch in microseconds. This field will be populated only after Recovery is finished.
		* `host_info` - (Optional, List) Specifies the host information for a objects. This is mainly populated in case of App objects where app object is hosted by another object such as VM or physical server.
		Nested schema for **host_info**:
			* `environment` - (Optional, String) Specifies the environment of the object.
			  * Constraints: Allowable values are: `kPhysical`, `kSQL`.
			* `id` - (Optional, String) Specifies the id of the host object.
			* `name` - (Optional, String) Specifies the name of the host object.
		* `is_encrypted` - (Optional, Boolean) Specifies whether the database is TDE enabled.
		* `messages` - (Computed, List) Specify error messages about the object.
		* `object_info` - (Optional, List) Specifies the information about the object for which the snapshot is taken.
		Nested schema for **object_info**:
			* `child_objects` - (Optional, List) Specifies child object details.
			Nested schema for **child_objects**:
				* `child_objects` - (Optional, List) Specifies child object details.
				Nested schema for **child_objects**:
				* `environment` - (Computed, String) Specifies the environment of the object.
				  * Constraints: Allowable values are: `kPhysical`, `kSQL`.
				* `global_id` - (Computed, String) Specifies the global id which is a unique identifier of the object.
				* `id` - (Computed, Integer) Specifies object id.
				* `logical_size_bytes` - (Computed, Integer) Specifies the logical size of object in bytes.
				* `name` - (Computed, String) Specifies the name of the object.
				* `object_hash` - (Computed, String) Specifies the hash identifier of the object.
				* `object_type` - (Computed, String) Specifies the type of the object.
				  * Constraints: Allowable values are: `kCluster`, `kVserver`, `kVolume`, `kVCenter`, `kStandaloneHost`, `kvCloudDirector`, `kFolder`, `kDatacenter`, `kComputeResource`, `kClusterComputeResource`, `kResourcePool`, `kDatastore`, `kHostSystem`, `kVirtualMachine`, `kVirtualApp`, `kStoragePod`, `kNetwork`, `kDistributedVirtualPortgroup`, `kTagCategory`, `kTag`, `kOpaqueNetwork`, `kOrganization`, `kVirtualDatacenter`, `kCatalog`, `kOrgMetadata`, `kStoragePolicy`, `kVirtualAppTemplate`, `kDomain`, `kOutlook`, `kMailbox`, `kUsers`, `kGroups`, `kSites`, `kUser`, `kGroup`, `kSite`, `kApplication`, `kGraphUser`, `kPublicFolders`, `kPublicFolder`, `kTeams`, `kTeam`, `kRootPublicFolder`, `kO365Exchange`, `kO365OneDrive`, `kO365Sharepoint`, `kKeyspace`, `kTable`, `kDatabase`, `kCollection`, `kBucket`, `kNamespace`, `kSCVMMServer`, `kStandaloneCluster`, `kHostGroup`, `kHypervHost`, `kHostCluster`, `kCustomProperty`, `kTenant`, `kSubscription`, `kResourceGroup`, `kStorageAccount`, `kStorageKey`, `kStorageContainer`, `kStorageBlob`, `kNetworkSecurityGroup`, `kVirtualNetwork`, `kSubnet`, `kComputeOptions`, `kSnapshotManagerPermit`, `kAvailabilitySet`, `kOVirtManager`, `kHost`, `kStorageDomain`, `kVNicProfile`, `kIAMUser`, `kRegion`, `kAvailabilityZone`, `kEC2Instance`, `kVPC`, `kInstanceType`, `kKeyPair`, `kRDSOptionGroup`, `kRDSParameterGroup`, `kRDSInstance`, `kRDSSubnet`, `kRDSTag`, `kAuroraTag`, `kAuroraCluster`, `kAccount`, `kSubTaskPermit`, `kS3Bucket`, `kS3Tag`, `kKmsKey`, `kProject`, `kLabel`, `kMetadata`, `kVPCConnector`, `kPrismCentral`, `kOtherHypervisorCluster`, `kZone`, `kMountPoint`, `kStorageArray`, `kFileSystem`, `kContainer`, `kFilesystem`, `kFileset`, `kPureProtectionGroup`, `kVolumeGroup`, `kStoragePool`, `kViewBox`, `kView`, `kWindowsCluster`, `kOracleRACCluster`, `kOracleAPCluster`, `kService`, `kPVC`, `kPersistentVolumeClaim`, `kPersistentVolume`, `kRootContainer`, `kDAGRootContainer`, `kExchangeNode`, `kExchangeDAGDatabaseCopy`, `kExchangeStandaloneDatabase`, `kExchangeDAG`, `kExchangeDAGDatabase`, `kDomainController`, `kInstance`, `kAAG`, `kAAGRootContainer`, `kAAGDatabase`, `kRACRootContainer`, `kTableSpace`, `kPDB`, `kObject`, `kOrg`, `kAppInstance`.
				* `os_type` - (Computed, String) Specifies the operating system type of the object.
				  * Constraints: Allowable values are: `kLinux`, `kWindows`.
				* `protection_type` - (Computed, String) Specifies the protection type of the object if any.
				  * Constraints: Allowable values are: `kAgent`, `kNative`, `kSnapshotManager`, `kRDSSnapshotManager`, `kAuroraSnapshotManager`, `kAwsS3`, `kAwsRDSPostgresBackup`, `kAwsAuroraPostgres`, `kAwsRDSPostgres`, `kAzureSQL`, `kFile`, `kVolume`.
				* `sharepoint_site_summary` - (Optional, List) Specifies the common parameters for Sharepoint site objects.
				Nested schema for **sharepoint_site_summary**:
					* `site_web_url` - (Computed, String) Specifies the web url for the Sharepoint site.
				* `source_id` - (Computed, Integer) Specifies registered source id to which object belongs.
				* `source_name` - (Computed, String) Specifies registered source name to which object belongs.
				* `uuid` - (Computed, String) Specifies the uuid which is a unique identifier of the object.
				* `v_center_summary` - (Optional, List)
				Nested schema for **v_center_summary**:
					* `is_cloud_env` - (Computed, Boolean) Specifies that registered vCenter source is a VMC (VMware Cloud) environment or not.
				* `windows_cluster_summary` - (Optional, List)
				Nested schema for **windows_cluster_summary**:
					* `cluster_source_type` - (Computed, String) Specifies the type of cluster resource this source represents.
			* `environment` - (Computed, String) Specifies the environment of the object.
			  * Constraints: Allowable values are: `kPhysical`, `kSQL`.
			* `global_id` - (Computed, String) Specifies the global id which is a unique identifier of the object.
			* `id` - (Computed, Integer) Specifies object id.
			* `logical_size_bytes` - (Computed, Integer) Specifies the logical size of object in bytes.
			* `name` - (Computed, String) Specifies the name of the object.
			* `object_hash` - (Computed, String) Specifies the hash identifier of the object.
			* `object_type` - (Computed, String) Specifies the type of the object.
			  * Constraints: Allowable values are: `kCluster`, `kVserver`, `kVolume`, `kVCenter`, `kStandaloneHost`, `kvCloudDirector`, `kFolder`, `kDatacenter`, `kComputeResource`, `kClusterComputeResource`, `kResourcePool`, `kDatastore`, `kHostSystem`, `kVirtualMachine`, `kVirtualApp`, `kStoragePod`, `kNetwork`, `kDistributedVirtualPortgroup`, `kTagCategory`, `kTag`, `kOpaqueNetwork`, `kOrganization`, `kVirtualDatacenter`, `kCatalog`, `kOrgMetadata`, `kStoragePolicy`, `kVirtualAppTemplate`, `kDomain`, `kOutlook`, `kMailbox`, `kUsers`, `kGroups`, `kSites`, `kUser`, `kGroup`, `kSite`, `kApplication`, `kGraphUser`, `kPublicFolders`, `kPublicFolder`, `kTeams`, `kTeam`, `kRootPublicFolder`, `kO365Exchange`, `kO365OneDrive`, `kO365Sharepoint`, `kKeyspace`, `kTable`, `kDatabase`, `kCollection`, `kBucket`, `kNamespace`, `kSCVMMServer`, `kStandaloneCluster`, `kHostGroup`, `kHypervHost`, `kHostCluster`, `kCustomProperty`, `kTenant`, `kSubscription`, `kResourceGroup`, `kStorageAccount`, `kStorageKey`, `kStorageContainer`, `kStorageBlob`, `kNetworkSecurityGroup`, `kVirtualNetwork`, `kSubnet`, `kComputeOptions`, `kSnapshotManagerPermit`, `kAvailabilitySet`, `kOVirtManager`, `kHost`, `kStorageDomain`, `kVNicProfile`, `kIAMUser`, `kRegion`, `kAvailabilityZone`, `kEC2Instance`, `kVPC`, `kInstanceType`, `kKeyPair`, `kRDSOptionGroup`, `kRDSParameterGroup`, `kRDSInstance`, `kRDSSubnet`, `kRDSTag`, `kAuroraTag`, `kAuroraCluster`, `kAccount`, `kSubTaskPermit`, `kS3Bucket`, `kS3Tag`, `kKmsKey`, `kProject`, `kLabel`, `kMetadata`, `kVPCConnector`, `kPrismCentral`, `kOtherHypervisorCluster`, `kZone`, `kMountPoint`, `kStorageArray`, `kFileSystem`, `kContainer`, `kFilesystem`, `kFileset`, `kPureProtectionGroup`, `kVolumeGroup`, `kStoragePool`, `kViewBox`, `kView`, `kWindowsCluster`, `kOracleRACCluster`, `kOracleAPCluster`, `kService`, `kPVC`, `kPersistentVolumeClaim`, `kPersistentVolume`, `kRootContainer`, `kDAGRootContainer`, `kExchangeNode`, `kExchangeDAGDatabaseCopy`, `kExchangeStandaloneDatabase`, `kExchangeDAG`, `kExchangeDAGDatabase`, `kDomainController`, `kInstance`, `kAAG`, `kAAGRootContainer`, `kAAGDatabase`, `kRACRootContainer`, `kTableSpace`, `kPDB`, `kObject`, `kOrg`, `kAppInstance`.
			* `os_type` - (Computed, String) Specifies the operating system type of the object.
			  * Constraints: Allowable values are: `kLinux`, `kWindows`.
			* `protection_type` - (Computed, String) Specifies the protection type of the object if any.
			  * Constraints: Allowable values are: `kAgent`, `kNative`, `kSnapshotManager`, `kRDSSnapshotManager`, `kAuroraSnapshotManager`, `kAwsS3`, `kAwsRDSPostgresBackup`, `kAwsAuroraPostgres`, `kAwsRDSPostgres`, `kAzureSQL`, `kFile`, `kVolume`.
			* `sharepoint_site_summary` - (Optional, List) Specifies the common parameters for Sharepoint site objects.
			Nested schema for **sharepoint_site_summary**:
				* `site_web_url` - (Computed, String) Specifies the web url for the Sharepoint site.
			* `source_id` - (Computed, Integer) Specifies registered source id to which object belongs.
			* `source_name` - (Computed, String) Specifies registered source name to which object belongs.
			* `uuid` - (Computed, String) Specifies the uuid which is a unique identifier of the object.
			* `v_center_summary` - (Optional, List)
			Nested schema for **v_center_summary**:
				* `is_cloud_env` - (Computed, Boolean) Specifies that registered vCenter source is a VMC (VMware Cloud) environment or not.
			* `windows_cluster_summary` - (Optional, List)
			Nested schema for **windows_cluster_summary**:
				* `cluster_source_type` - (Computed, String) Specifies the type of cluster resource this source represents.
		* `point_in_time_usecs` - (Optional, Integer) Specifies the timestamp (in microseconds. from epoch) for recovering to a point-in-time in the past.
		* `progress_task_id` - (Computed, String) Progress monitor task id for Recovery of VM.
		* `protection_group_id` - (Optional, String) Specifies the protection group id of the object snapshot.
		* `protection_group_name` - (Optional, String) Specifies the protection group name of the object snapshot.
		* `recover_from_standby` - (Optional, Boolean) Specifies that user wants to perform standby restore if it is enabled for this object.
		* `snapshot_creation_time_usecs` - (Computed, Integer) Specifies the time when the snapshot is created in Unix timestamp epoch in microseconds.
		* `snapshot_id` - (Required, String) Specifies the snapshot id.
		* `snapshot_target_type` - (Computed, String) Specifies the snapshot target type.
		  * Constraints: Allowable values are: `Local`, `Archival`, `RpaasArchival`, `StorageArraySnapshot`, `Remote`.
		* `sql_target_params` - (Optional, List) Specifies the params for recovering to a sql host. Specifiy seperate settings for each db object that need to be recovered. Provided sql backup should be recovered to same type of target host. For Example: If you have sql backup taken from a physical host then that should be recovered to physical host only.
		Nested schema for **sql_target_params**:
			* `new_source_config` - (Optional, List) Specifies the destination Source configuration parameters where the databases will be recovered. This is mandatory if recoverToNewSource is set to true.
			Nested schema for **new_source_config**:
				* `data_file_directory_location` - (Required, String) Specifies the directory where to put the database data files. Missing directory will be automatically created.
				* `database_name` - (Optional, String) Specifies a new name for the restored database. If this field is not specified, then the original database will be overwritten after recovery.
				* `host` - (Required, List) Specifies the source id of target host where databases will be recovered. This source id can be a physical host or virtual machine.
				Nested schema for **host**:
					* `id` - (Required, Integer) Specifies the id of the object.
					* `name` - (Computed, String) Specifies the name of the object.
				* `instance_name` - (Required, String) Specifies an instance name of the Sql Server that should be used for restoring databases to.
				* `keep_cdc` - (Optional, Boolean) Specifies whether to keep CDC (Change Data Capture) on recovered databases or not. If not passed, this is assumed to be true. If withNoRecovery is passed as true, then this field must not be set to true. Passing this field as true in this scenario will be a invalid request.
				* `log_file_directory_location` - (Required, String) Specifies the directory where to put the database log files. Missing directory will be automatically created.
				* `multi_stage_restore_options` - (Optional, List) Specifies the parameters related to multi stage Sql restore.
				Nested schema for **multi_stage_restore_options**:
					* `enable_auto_sync` - (Optional, Boolean) Set this to true if you want to enable auto sync for multi stage restore.
					* `enable_multi_stage_restore` - (Optional, Boolean) Set this to true if you are creating a multi-stage Sql restore task needed for features such as Hot-Standby.
				* `native_log_recovery_with_clause` - (Optional, String) Specifies the WITH clause to be used in native sql log restore command. This is only applicable for native log restore.
				* `native_recovery_with_clause` - (Optional, String) 'with_clause' contains 'with clause' to be used in native sql restore command. This is only applicable for database restore of native sql backup. Here user can specify multiple restore options. Example: 'WITH BUFFERCOUNT = 575, MAXTRANSFERSIZE = 2097152'.
				* `overwriting_policy` - (Optional, String) Specifies a policy to be used while recovering existing databases.
				  * Constraints: Allowable values are: `FailIfExists`, `Overwrite`.
				* `replay_entire_last_log` - (Optional, Boolean) Specifies the option to set replay last log bit while creating the sql restore task and doing restore to latest point-in-time. If this is set to true, we will replay the entire last log without STOPAT.
				* `restore_time_usecs` - (Optional, Integer) Specifies the time in the past to which the Sql database needs to be restored. This allows for granular recovery of Sql databases. If this is not set, the Sql database will be restored from the full/incremental snapshot.
				* `secondary_data_files_dir_list` - (Optional, List) Specifies the secondary data filename pattern and corresponding direcories of the DB. Secondary data files are optional and are user defined. The recommended file extention for secondary files is ".ndf". If this option is specified and the destination folders do not exist they will be automatically created.
				Nested schema for **secondary_data_files_dir_list**:
					* `directory` - (Optional, String) Specifies the directory where to keep the files matching the pattern.
					* `filename_pattern` - (Optional, String) Specifies a pattern to be matched with filenames. This can be a regex expression.
				* `with_no_recovery` - (Optional, Boolean) Specifies the flag to bring DBs online or not after successful recovery. If this is passed as true, then it means DBs won't be brought online.
			* `original_source_config` - (Optional, List) Specifies the Source configuration if databases are being recovered to Original Source. If not specified, all the configuration parameters will be retained.
			Nested schema for **original_source_config**:
				* `capture_tail_logs` - (Optional, Boolean) Set this to true if tail logs are to be captured before the recovery operation. This is only applicable if database is not being renamed.
				* `data_file_directory_location` - (Optional, String) Specifies the directory where to put the database data files. Missing directory will be automatically created. If you are overwriting the existing database then this field will be ignored.
				* `keep_cdc` - (Optional, Boolean) Specifies whether to keep CDC (Change Data Capture) on recovered databases or not. If not passed, this is assumed to be true. If withNoRecovery is passed as true, then this field must not be set to true. Passing this field as true in this scenario will be a invalid request.
				* `log_file_directory_location` - (Optional, String) Specifies the directory where to put the database log files. Missing directory will be automatically created. If you are overwriting the existing database then this field will be ignored.
				* `multi_stage_restore_options` - (Optional, List) Specifies the parameters related to multi stage Sql restore.
				Nested schema for **multi_stage_restore_options**:
					* `enable_auto_sync` - (Optional, Boolean) Set this to true if you want to enable auto sync for multi stage restore.
					* `enable_multi_stage_restore` - (Optional, Boolean) Set this to true if you are creating a multi-stage Sql restore task needed for features such as Hot-Standby.
				* `native_log_recovery_with_clause` - (Optional, String) Specifies the WITH clause to be used in native sql log restore command. This is only applicable for native log restore.
				* `native_recovery_with_clause` - (Optional, String) 'with_clause' contains 'with clause' to be used in native sql restore command. This is only applicable for database restore of native sql backup. Here user can specify multiple restore options. Example: 'WITH BUFFERCOUNT = 575, MAXTRANSFERSIZE = 2097152'.
				* `new_database_name` - (Optional, String) Specifies a new name for the restored database. If this field is not specified, then the original database will be overwritten after recovery.
				* `overwriting_policy` - (Optional, String) Specifies a policy to be used while recovering existing databases.
				  * Constraints: Allowable values are: `FailIfExists`, `Overwrite`.
				* `replay_entire_last_log` - (Optional, Boolean) Specifies the option to set replay last log bit while creating the sql restore task and doing restore to latest point-in-time. If this is set to true, we will replay the entire last log without STOPAT.
				* `restore_time_usecs` - (Optional, Integer) Specifies the time in the past to which the Sql database needs to be restored. This allows for granular recovery of Sql databases. If this is not set, the Sql database will be restored from the full/incremental snapshot.
				* `secondary_data_files_dir_list` - (Optional, List) Specifies the secondary data filename pattern and corresponding direcories of the DB. Secondary data files are optional and are user defined. The recommended file extention for secondary files is ".ndf". If this option is specified and the destination folders do not exist they will be automatically created.
				Nested schema for **secondary_data_files_dir_list**:
					* `directory` - (Optional, String) Specifies the directory where to keep the files matching the pattern.
					* `filename_pattern` - (Optional, String) Specifies a pattern to be matched with filenames. This can be a regex expression.
				* `with_no_recovery` - (Optional, Boolean) Specifies the flag to bring DBs online or not after successful recovery. If this is passed as true, then it means DBs won't be brought online.
			* `recover_to_new_source` - (Required, Boolean) Specifies the parameter whether the recovery should be performed to a new sources or an original Source Target.
		* `start_time_usecs` - (Computed, Integer) Specifies the start time of the Recovery in Unix timestamp epoch in microseconds.
		* `status` - (Computed, String) Status of the Recovery. 'Running' indicates that the Recovery is still running. 'Canceled' indicates that the Recovery has been cancelled. 'Canceling' indicates that the Recovery is in the process of being cancelled. 'Failed' indicates that the Recovery has failed. 'Succeeded' indicates that the Recovery has finished successfully. 'SucceededWithWarning' indicates that the Recovery finished successfully, but there were some warning messages. 'Skipped' indicates that the Recovery task was skipped.
		  * Constraints: Allowable values are: `Accepted`, `Running`, `Canceled`, `Canceling`, `Failed`, `Missed`, `Succeeded`, `SucceededWithWarning`, `OnHold`, `Finalizing`, `Skipped`, `LegalHold`.
		* `target_environment` - (Required, String) Specifies the environment of the recovery target. The corresponding params below must be filled out.
		  * Constraints: Allowable values are: `kSQL`.
	* `recovery_action` - (Required, String) Specifies the type of recover action to be performed.
	  * Constraints: Allowable values are: `RecoverApps`, `CloneApps`.
	* `vlan_config` - (Optional, List) Specifies VLAN Params associated with the recovered. If this is not specified, then the VLAN settings will be automatically selected from one of the below options: a. If VLANs are configured on IBM, then the VLAN host/VIP will be automatically based on the client's (e.g. ESXI host) IP address. b. If VLANs are not configured on IBM, then the partition hostname or VIPs will be used for Recovery.
	Nested schema for **vlan_config**:
		* `disable_vlan` - (Optional, Boolean) If this is set to true, then even if VLANs are configured on the system, the partition VIPs will be used for the Recovery.
		* `id` - (Optional, Integer) If this is set, then the Cohesity host name or the IP address associated with this vlan is used for mounting Cohesity's view on the remote host.
		* `interface_name` - (Computed, String) Interface group to use for Recovery.
* `name` - (Required, Forces new resource, String) Specifies the name of the Recovery.
* `endpoint_type` - (Optional, String) Backup Recovery Endpoint type. By default set to "public".
* `instance_id` - (Optional, String) Backup Recovery instance ID. If provided here along with region, the provider constructs the endpoint URL using them, which overrides any value set through environment variables or the `endpoints.json` file.
* `region` - (Optional, String) Backup Recovery region. If provided here along with instance_id, the provider constructs the endpoint URL using them, which overrides any value set through environment variables or the `endpoints.json` file.  
* `physical_params` - (Optional, Forces new resource, List) Specifies the recovery options specific to Physical environment.
Nested schema for **physical_params**:
	* `download_file_and_folder_params` - (Optional, List) Specifies the parameters to download files and folders.
	Nested schema for **download_file_and_folder_params**:
		* `download_file_path` - (Optional, String) Specifies the path location to download the files and folders.
		* `expiry_time_usecs` - (Optional, Integer) Specifies the time upto which the download link is available.
		* `files_and_folders` - (Optional, List) Specifies the info about the files and folders to be recovered.
		Nested schema for **files_and_folders**:
			* `absolute_path` - (Required, String) Specifies the absolute path to the file or folder.
			* `destination_dir` - (Computed, String) Specifies the destination directory where the file/directory was copied.
			* `is_directory` - (Optional, Boolean) Specifies whether this is a directory or not.
			* `is_view_file_recovery` - (Optional, Boolean) Specify if the recovery is of type view file/folder.
			* `messages` - (Computed, List) Specify error messages about the file during recovery.
			* `status` - (Computed, String) Specifies the recovery status for this file or folder.
			  * Constraints: Allowable values are: `NotStarted`, `EstimationInProgress`, `EstimationDone`, `CopyInProgress`, `Finished`.
	* `mount_volume_params` - (Optional, List) Specifies the parameters to mount Physical Volumes.
	Nested schema for **mount_volume_params**:
		* `physical_target_params` - (Optional, List) Specifies the params for recovering to a physical target.
		Nested schema for **physical_target_params**:
			* `mount_to_original_target` - (Required, Boolean) Specifies whether to mount to the original target. If true, originalTargetConfig must be specified. If false, newTargetConfig must be specified.
			* `mounted_volume_mapping` - (Optional, List) Specifies the mapping of original volumes and mounted volumes.
			Nested schema for **mounted_volume_mapping**:
				* `file_system_type` - (Computed, String) Specifies the type of the file system of the volume.
				* `mounted_volume` - (Computed, String) Specifies the name of the point where the volume is mounted.
				* `original_volume` - (Computed, String) Specifies the name of the original volume.
			* `new_target_config` - (Optional, List) Specifies the configuration for mounting to a new target.
			Nested schema for **new_target_config**:
				* `mount_target` - (Required, List) Specifies the target entity to recover to.
				Nested schema for **mount_target**:
					* `id` - (Required, Integer) Specifies the id of the object.
					* `name` - (Computed, String) Specifies the name of the object.
					* `parent_source_id` - (Computed, Integer) Specifies the id of the parent source of the target.
					* `parent_source_name` - (Computed, String) Specifies the name of the parent source of the target.
				* `server_credentials` - (Optional, List) Specifies credentials to access the target server. This is required if the server is of Linux OS.
				Nested schema for **server_credentials**:
					* `password` - (Required, String) Specifies the password to access target entity.
					* `username` - (Required, String) Specifies the username to access target entity.
			* `original_target_config` - (Optional, List) Specifies the configuration for mounting to the original target.
			Nested schema for **original_target_config**:
				* `server_credentials` - (Optional, List) Specifies credentials to access the target server. This is required if the server is of Linux OS.
				Nested schema for **server_credentials**:
					* `password` - (Required, String) Specifies the password to access target entity.
					* `username` - (Required, String) Specifies the username to access target entity.
			* `read_only_mount` - (Optional, Boolean) Specifies whether to perform a read-only mount. Default is false.
			* `vlan_config` - (Optional, List) Specifies VLAN Params associated with the recovered. If this is not specified, then the VLAN settings will be automatically selected from one of the below options: a. If VLANs are configured on Cohesity, then the VLAN host/VIP will be automatically based on the client's (e.g. ESXI host) IP address. b. If VLANs are not configured on Cohesity, then the partition hostname or VIPs will be used for Recovery.
			Nested schema for **vlan_config**:
				* `disable_vlan` - (Optional, Boolean) If this is set to true, then even if VLANs are configured on the system, the partition VIPs will be used for the Recovery.
				* `id` - (Optional, Integer) If this is set, then the Cohesity host name or the IP address associated with this vlan is used for mounting Cohesity's view on the remote host.
				* `interface_name` - (Computed, String) Interface group to use for Recovery.
			* `volume_names` - (Optional, List) Specifies the names of volumes that need to be mounted. If this is not specified then all volumes that are part of the source VM will be mounted on the target VM.
		* `target_environment` - (Required, String) Specifies the environment of the recovery target. The corresponding params below must be filled out.
		  * Constraints: Allowable values are: `kPhysical`.
	* `objects` - (Required, List) Specifies the list of Recover Object parameters. For recovering files, specifies the object contains the file to recover.
	Nested schema for **objects**:
		* `archival_target_info` - (Optional, List) Specifies the archival target information if the snapshot is an archival snapshot.
		Nested schema for **archival_target_info**:
			* `archival_task_id` - (Computed, String) Specifies the archival task id. This is a protection group UID which only applies when archival type is 'Tape'.
			* `ownership_context` - (Computed, String) Specifies the ownership context for the target.
			  * Constraints: Allowable values are: `Local`, `FortKnox`.
			* `target_id` - (Computed, Integer) Specifies the archival target ID.
			* `target_name` - (Computed, String) Specifies the archival target name.
			* `target_type` - (Computed, String) Specifies the archival target type.
			  * Constraints: Allowable values are: `Tape`, `Cloud`, `Nas`.
			* `tier_settings` - (Optional, List) Specifies the tier info for archival.
			Nested schema for **tier_settings**:
				* `aws_tiering` - (Optional, List) Specifies aws tiers.
				Nested schema for **aws_tiering**:
					* `tiers` - (Required, List) Specifies the tiers that are used to move the archived backup from current tier to next tier. The order of the tiers determines which tier will be used next for moving the archived backup. The first tier input should always be default tier where backup will be acrhived. Each tier specifies how much time after the backup will be moved to next tier from the current tier.
					Nested schema for **tiers**:
						* `move_after` - (Computed, Integer) Specifies the time period after which the backup will be moved from current tier to next tier.
						* `move_after_unit` - (Computed, String) Specifies the unit for moving the data from current tier to next tier. This unit will be a base unit for the 'moveAfter' field specified below.
						  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
						* `tier_type` - (Computed, String) Specifies the AWS tier types.
						  * Constraints: Allowable values are: `kAmazonS3Standard`, `kAmazonS3StandardIA`, `kAmazonS3OneZoneIA`, `kAmazonS3IntelligentTiering`, `kAmazonS3Glacier`, `kAmazonS3GlacierDeepArchive`.
				* `azure_tiering` - (Optional, List) Specifies Azure tiers.
				Nested schema for **azure_tiering**:
					* `tiers` - (Optional, List) Specifies the tiers that are used to move the archived backup from current tier to next tier. The order of the tiers determines which tier will be used next for moving the archived backup. The first tier input should always be default tier where backup will be acrhived. Each tier specifies how much time after the backup will be moved to next tier from the current tier.
					Nested schema for **tiers**:
						* `move_after` - (Computed, Integer) Specifies the time period after which the backup will be moved from current tier to next tier.
						* `move_after_unit` - (Computed, String) Specifies the unit for moving the data from current tier to next tier. This unit will be a base unit for the 'moveAfter' field specified below.
						  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
						* `tier_type` - (Computed, String) Specifies the Azure tier types.
						  * Constraints: Allowable values are: `kAzureTierHot`, `kAzureTierCool`, `kAzureTierArchive`.
				* `cloud_platform` - (Computed, String) Specifies the cloud platform to enable tiering.
				  * Constraints: Allowable values are: `AWS`, `Azure`, `Oracle`, `Google`.
				* `current_tier_type` - (Computed, String) Specifies the type of the current tier where the snapshot resides. This will be specified if the run is a CAD run.
				  * Constraints: Allowable values are: `kAmazonS3Standard`, `kAmazonS3StandardIA`, `kAmazonS3OneZoneIA`, `kAmazonS3IntelligentTiering`, `kAmazonS3Glacier`, `kAmazonS3GlacierDeepArchive`, `kAzureTierHot`, `kAzureTierCool`, `kAzureTierArchive`, `kGoogleStandard`, `kGoogleRegional`, `kGoogleMultiRegional`, `kGoogleNearline`, `kGoogleColdline`, `kOracleTierStandard`, `kOracleTierArchive`.
				* `google_tiering` - (Optional, List) Specifies Google tiers.
				Nested schema for **google_tiering**:
					* `tiers` - (Required, List) Specifies the tiers that are used to move the archived backup from current tier to next tier. The order of the tiers determines which tier will be used next for moving the archived backup. The first tier input should always be default tier where backup will be acrhived. Each tier specifies how much time after the backup will be moved to next tier from the current tier.
					Nested schema for **tiers**:
						* `move_after` - (Computed, Integer) Specifies the time period after which the backup will be moved from current tier to next tier.
						* `move_after_unit` - (Computed, String) Specifies the unit for moving the data from current tier to next tier. This unit will be a base unit for the 'moveAfter' field specified below.
						  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
						* `tier_type` - (Computed, String) Specifies the Google tier types.
						  * Constraints: Allowable values are: `kGoogleStandard`, `kGoogleRegional`, `kGoogleMultiRegional`, `kGoogleNearline`, `kGoogleColdline`.
				* `oracle_tiering` - (Optional, List) Specifies Oracle tiers.
				Nested schema for **oracle_tiering**:
					* `tiers` - (Required, List) Specifies the tiers that are used to move the archived backup from current tier to next tier. The order of the tiers determines which tier will be used next for moving the archived backup. The first tier input should always be default tier where backup will be acrhived. Each tier specifies how much time after the backup will be moved to next tier from the current tier.
					Nested schema for **tiers**:
						* `move_after` - (Computed, Integer) Specifies the time period after which the backup will be moved from current tier to next tier.
						* `move_after_unit` - (Computed, String) Specifies the unit for moving the data from current tier to next tier. This unit will be a base unit for the 'moveAfter' field specified below.
						  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
						* `tier_type` - (Computed, String) Specifies the Oracle tier types.
						  * Constraints: Allowable values are: `kOracleTierStandard`, `kOracleTierArchive`.
			* `usage_type` - (Computed, String) Specifies the usage type for the target.
			  * Constraints: Allowable values are: `Archival`, `Tiering`, `Rpaas`.
		* `bytes_restored` - (Computed, Integer) Specify the total bytes restored.
		* `end_time_usecs` - (Computed, Integer) Specifies the end time of the Recovery in Unix timestamp epoch in microseconds. This field will be populated only after Recovery is finished.
		* `messages` - (Computed, List) Specify error messages about the object.
		* `object_info` - (Optional, List) Specifies the information about the object for which the snapshot is taken.
		Nested schema for **object_info**:
			* `child_objects` - (Optional, List) Specifies child object details.
			Nested schema for **child_objects**:
				* `child_objects` - (Optional, List) Specifies child object details.
				Nested schema for **child_objects**:
				* `environment` - (Computed, String) Specifies the environment of the object.
				  * Constraints: Allowable values are: `kPhysical`, `kSQL`.
				* `global_id` - (Computed, String) Specifies the global id which is a unique identifier of the object.
				* `id` - (Computed, Integer) Specifies object id.
				* `logical_size_bytes` - (Computed, Integer) Specifies the logical size of object in bytes.
				* `name` - (Computed, String) Specifies the name of the object.
				* `object_hash` - (Computed, String) Specifies the hash identifier of the object.
				* `object_type` - (Computed, String) Specifies the type of the object.
				  * Constraints: Allowable values are: `kCluster`, `kVserver`, `kVolume`, `kVCenter`, `kStandaloneHost`, `kvCloudDirector`, `kFolder`, `kDatacenter`, `kComputeResource`, `kClusterComputeResource`, `kResourcePool`, `kDatastore`, `kHostSystem`, `kVirtualMachine`, `kVirtualApp`, `kStoragePod`, `kNetwork`, `kDistributedVirtualPortgroup`, `kTagCategory`, `kTag`, `kOpaqueNetwork`, `kOrganization`, `kVirtualDatacenter`, `kCatalog`, `kOrgMetadata`, `kStoragePolicy`, `kVirtualAppTemplate`, `kDomain`, `kOutlook`, `kMailbox`, `kUsers`, `kGroups`, `kSites`, `kUser`, `kGroup`, `kSite`, `kApplication`, `kGraphUser`, `kPublicFolders`, `kPublicFolder`, `kTeams`, `kTeam`, `kRootPublicFolder`, `kO365Exchange`, `kO365OneDrive`, `kO365Sharepoint`, `kKeyspace`, `kTable`, `kDatabase`, `kCollection`, `kBucket`, `kNamespace`, `kSCVMMServer`, `kStandaloneCluster`, `kHostGroup`, `kHypervHost`, `kHostCluster`, `kCustomProperty`, `kTenant`, `kSubscription`, `kResourceGroup`, `kStorageAccount`, `kStorageKey`, `kStorageContainer`, `kStorageBlob`, `kNetworkSecurityGroup`, `kVirtualNetwork`, `kSubnet`, `kComputeOptions`, `kSnapshotManagerPermit`, `kAvailabilitySet`, `kOVirtManager`, `kHost`, `kStorageDomain`, `kVNicProfile`, `kIAMUser`, `kRegion`, `kAvailabilityZone`, `kEC2Instance`, `kVPC`, `kInstanceType`, `kKeyPair`, `kRDSOptionGroup`, `kRDSParameterGroup`, `kRDSInstance`, `kRDSSubnet`, `kRDSTag`, `kAuroraTag`, `kAuroraCluster`, `kAccount`, `kSubTaskPermit`, `kS3Bucket`, `kS3Tag`, `kKmsKey`, `kProject`, `kLabel`, `kMetadata`, `kVPCConnector`, `kPrismCentral`, `kOtherHypervisorCluster`, `kZone`, `kMountPoint`, `kStorageArray`, `kFileSystem`, `kContainer`, `kFilesystem`, `kFileset`, `kPureProtectionGroup`, `kVolumeGroup`, `kStoragePool`, `kViewBox`, `kView`, `kWindowsCluster`, `kOracleRACCluster`, `kOracleAPCluster`, `kService`, `kPVC`, `kPersistentVolumeClaim`, `kPersistentVolume`, `kRootContainer`, `kDAGRootContainer`, `kExchangeNode`, `kExchangeDAGDatabaseCopy`, `kExchangeStandaloneDatabase`, `kExchangeDAG`, `kExchangeDAGDatabase`, `kDomainController`, `kInstance`, `kAAG`, `kAAGRootContainer`, `kAAGDatabase`, `kRACRootContainer`, `kTableSpace`, `kPDB`, `kObject`, `kOrg`, `kAppInstance`.
				* `os_type` - (Computed, String) Specifies the operating system type of the object.
				  * Constraints: Allowable values are: `kLinux`, `kWindows`.
				* `protection_type` - (Computed, String) Specifies the protection type of the object if any.
				  * Constraints: Allowable values are: `kAgent`, `kNative`, `kSnapshotManager`, `kRDSSnapshotManager`, `kAuroraSnapshotManager`, `kAwsS3`, `kAwsRDSPostgresBackup`, `kAwsAuroraPostgres`, `kAwsRDSPostgres`, `kAzureSQL`, `kFile`, `kVolume`.
				* `sharepoint_site_summary` - (Optional, List) Specifies the common parameters for Sharepoint site objects.
				Nested schema for **sharepoint_site_summary**:
					* `site_web_url` - (Computed, String) Specifies the web url for the Sharepoint site.
				* `source_id` - (Computed, Integer) Specifies registered source id to which object belongs.
				* `source_name` - (Computed, String) Specifies registered source name to which object belongs.
				* `uuid` - (Computed, String) Specifies the uuid which is a unique identifier of the object.
				* `v_center_summary` - (Optional, List)
				Nested schema for **v_center_summary**:
					* `is_cloud_env` - (Computed, Boolean) Specifies that registered vCenter source is a VMC (VMware Cloud) environment or not.
				* `windows_cluster_summary` - (Optional, List)
				Nested schema for **windows_cluster_summary**:
					* `cluster_source_type` - (Computed, String) Specifies the type of cluster resource this source represents.
			* `environment` - (Computed, String) Specifies the environment of the object.
			  * Constraints: Allowable values are: `kPhysical`, `kSQL`.
			* `global_id` - (Computed, String) Specifies the global id which is a unique identifier of the object.
			* `id` - (Computed, Integer) Specifies object id.
			* `logical_size_bytes` - (Computed, Integer) Specifies the logical size of object in bytes.
			* `name` - (Computed, String) Specifies the name of the object.
			* `object_hash` - (Computed, String) Specifies the hash identifier of the object.
			* `object_type` - (Computed, String) Specifies the type of the object.
			  * Constraints: Allowable values are: `kCluster`, `kVserver`, `kVolume`, `kVCenter`, `kStandaloneHost`, `kvCloudDirector`, `kFolder`, `kDatacenter`, `kComputeResource`, `kClusterComputeResource`, `kResourcePool`, `kDatastore`, `kHostSystem`, `kVirtualMachine`, `kVirtualApp`, `kStoragePod`, `kNetwork`, `kDistributedVirtualPortgroup`, `kTagCategory`, `kTag`, `kOpaqueNetwork`, `kOrganization`, `kVirtualDatacenter`, `kCatalog`, `kOrgMetadata`, `kStoragePolicy`, `kVirtualAppTemplate`, `kDomain`, `kOutlook`, `kMailbox`, `kUsers`, `kGroups`, `kSites`, `kUser`, `kGroup`, `kSite`, `kApplication`, `kGraphUser`, `kPublicFolders`, `kPublicFolder`, `kTeams`, `kTeam`, `kRootPublicFolder`, `kO365Exchange`, `kO365OneDrive`, `kO365Sharepoint`, `kKeyspace`, `kTable`, `kDatabase`, `kCollection`, `kBucket`, `kNamespace`, `kSCVMMServer`, `kStandaloneCluster`, `kHostGroup`, `kHypervHost`, `kHostCluster`, `kCustomProperty`, `kTenant`, `kSubscription`, `kResourceGroup`, `kStorageAccount`, `kStorageKey`, `kStorageContainer`, `kStorageBlob`, `kNetworkSecurityGroup`, `kVirtualNetwork`, `kSubnet`, `kComputeOptions`, `kSnapshotManagerPermit`, `kAvailabilitySet`, `kOVirtManager`, `kHost`, `kStorageDomain`, `kVNicProfile`, `kIAMUser`, `kRegion`, `kAvailabilityZone`, `kEC2Instance`, `kVPC`, `kInstanceType`, `kKeyPair`, `kRDSOptionGroup`, `kRDSParameterGroup`, `kRDSInstance`, `kRDSSubnet`, `kRDSTag`, `kAuroraTag`, `kAuroraCluster`, `kAccount`, `kSubTaskPermit`, `kS3Bucket`, `kS3Tag`, `kKmsKey`, `kProject`, `kLabel`, `kMetadata`, `kVPCConnector`, `kPrismCentral`, `kOtherHypervisorCluster`, `kZone`, `kMountPoint`, `kStorageArray`, `kFileSystem`, `kContainer`, `kFilesystem`, `kFileset`, `kPureProtectionGroup`, `kVolumeGroup`, `kStoragePool`, `kViewBox`, `kView`, `kWindowsCluster`, `kOracleRACCluster`, `kOracleAPCluster`, `kService`, `kPVC`, `kPersistentVolumeClaim`, `kPersistentVolume`, `kRootContainer`, `kDAGRootContainer`, `kExchangeNode`, `kExchangeDAGDatabaseCopy`, `kExchangeStandaloneDatabase`, `kExchangeDAG`, `kExchangeDAGDatabase`, `kDomainController`, `kInstance`, `kAAG`, `kAAGRootContainer`, `kAAGDatabase`, `kRACRootContainer`, `kTableSpace`, `kPDB`, `kObject`, `kOrg`, `kAppInstance`.
			* `os_type` - (Computed, String) Specifies the operating system type of the object.
			  * Constraints: Allowable values are: `kLinux`, `kWindows`.
			* `protection_type` - (Computed, String) Specifies the protection type of the object if any.
			  * Constraints: Allowable values are: `kAgent`, `kNative`, `kSnapshotManager`, `kRDSSnapshotManager`, `kAuroraSnapshotManager`, `kAwsS3`, `kAwsRDSPostgresBackup`, `kAwsAuroraPostgres`, `kAwsRDSPostgres`, `kAzureSQL`, `kFile`, `kVolume`.
			* `sharepoint_site_summary` - (Optional, List) Specifies the common parameters for Sharepoint site objects.
			Nested schema for **sharepoint_site_summary**:
				* `site_web_url` - (Computed, String) Specifies the web url for the Sharepoint site.
			* `source_id` - (Computed, Integer) Specifies registered source id to which object belongs.
			* `source_name` - (Computed, String) Specifies registered source name to which object belongs.
			* `uuid` - (Computed, String) Specifies the uuid which is a unique identifier of the object.
			* `v_center_summary` - (Optional, List)
			Nested schema for **v_center_summary**:
				* `is_cloud_env` - (Computed, Boolean) Specifies that registered vCenter source is a VMC (VMware Cloud) environment or not.
			* `windows_cluster_summary` - (Optional, List)
			Nested schema for **windows_cluster_summary**:
				* `cluster_source_type` - (Computed, String) Specifies the type of cluster resource this source represents.
		* `point_in_time_usecs` - (Optional, Integer) Specifies the timestamp (in microseconds. from epoch) for recovering to a point-in-time in the past.
		* `progress_task_id` - (Computed, String) Progress monitor task id for Recovery of VM.
		* `protection_group_id` - (Optional, String) Specifies the protection group id of the object snapshot.
		* `protection_group_name` - (Optional, String) Specifies the protection group name of the object snapshot.
		* `recover_from_standby` - (Optional, Boolean) Specifies that user wants to perform standby restore if it is enabled for this object.
		* `snapshot_creation_time_usecs` - (Computed, Integer) Specifies the time when the snapshot is created in Unix timestamp epoch in microseconds.
		* `snapshot_id` - (Required, String) Specifies the snapshot id.
		* `snapshot_target_type` - (Computed, String) Specifies the snapshot target type.
		  * Constraints: Allowable values are: `Local`, `Archival`, `RpaasArchival`, `StorageArraySnapshot`, `Remote`.
		* `start_time_usecs` - (Computed, Integer) Specifies the start time of the Recovery in Unix timestamp epoch in microseconds.
		* `status` - (Computed, String) Status of the Recovery. 'Running' indicates that the Recovery is still running. 'Canceled' indicates that the Recovery has been cancelled. 'Canceling' indicates that the Recovery is in the process of being cancelled. 'Failed' indicates that the Recovery has failed. 'Succeeded' indicates that the Recovery has finished successfully. 'SucceededWithWarning' indicates that the Recovery finished successfully, but there were some warning messages. 'Skipped' indicates that the Recovery task was skipped.
		  * Constraints: Allowable values are: `Accepted`, `Running`, `Canceled`, `Canceling`, `Failed`, `Missed`, `Succeeded`, `SucceededWithWarning`, `OnHold`, `Finalizing`, `Skipped`, `LegalHold`.
	* `recover_file_and_folder_params` - (Optional, List) Specifies the parameters to perform a file and folder recovery.
	Nested schema for **recover_file_and_folder_params**:
		* `files_and_folders` - (Required, List) Specifies the information about the files and folders to be recovered.
		Nested schema for **files_and_folders**:
			* `absolute_path` - (Required, String) Specifies the absolute path to the file or folder.
			* `destination_dir` - (Computed, String) Specifies the destination directory where the file/directory was copied.
			* `is_directory` - (Optional, Boolean) Specifies whether this is a directory or not.
			* `is_view_file_recovery` - (Optional, Boolean) Specify if the recovery is of type view file/folder.
			* `messages` - (Computed, List) Specify error messages about the file during recovery.
			* `status` - (Computed, String) Specifies the recovery status for this file or folder.
			  * Constraints: Allowable values are: `NotStarted`, `EstimationInProgress`, `EstimationDone`, `CopyInProgress`, `Finished`.
		* `physical_target_params` - (Optional, List) Specifies the parameters to recover to a Physical target.
		Nested schema for **physical_target_params**:
			* `alternate_restore_directory` - (Optional, String) Specifies the directory path where restore should happen if restore_to_original_paths is set to false.
			* `continue_on_error` - (Optional, Boolean) Specifies whether to continue recovering other volumes if one of the volumes fails to recover. Default value is false.
			* `overwrite_existing` - (Optional, Boolean) Specifies whether to overwrite existing file/folder during recovery.
			* `preserve_acls` - (Optional, Boolean) Whether to preserve the ACLs of the original file.
			* `preserve_attributes` - (Optional, Boolean) Specifies whether to preserve file/folder attributes during recovery.
			* `preserve_timestamps` - (Optional, Boolean) Whether to preserve the original time stamps.
			* `recover_target` - (Required, List) Specifies the target entity where the volumes are being mounted.
			Nested schema for **recover_target**:
				* `id` - (Required, Integer) Specifies the id of the object.
				* `name` - (Computed, String) Specifies the name of the object.
				* `parent_source_id` - (Computed, Integer) Specifies the id of the parent source of the target.
				* `parent_source_name` - (Computed, String) Specifies the name of the parent source of the target.
			* `restore_entity_type` - (Optional, String) Specifies the restore type (restore everything or ACLs only) when restoring or downloading files or folders from a Physical file based or block based backup snapshot.
			  * Constraints: Allowable values are: `kRegular`, `kACLOnly`.
			* `restore_to_original_paths` - (Optional, Boolean) If this is true, then files will be restored to original paths.
			* `save_success_files` - (Optional, Boolean) Specifies whether to save success files or not. Default value is false.
			* `vlan_config` - (Optional, List) Specifies VLAN Params associated with the recovered. If this is not specified, then the VLAN settings will be automatically selected from one of the below options: a. If VLANs are configured on Cohesity, then the VLAN host/VIP will be automatically based on the client's (e.g. ESXI host) IP address. b. If VLANs are not configured on Cohesity, then the partition hostname or VIPs will be used for Recovery.
			Nested schema for **vlan_config**:
				* `disable_vlan` - (Optional, Boolean) If this is set to true, then even if VLANs are configured on the system, the partition VIPs will be used for the Recovery.
				* `id` - (Optional, Integer) If this is set, then the Cohesity host name or the IP address associated with this vlan is used for mounting Cohesity's view on the remote host.
				* `interface_name` - (Computed, String) Interface group to use for Recovery.
		* `target_environment` - (Required, String) Specifies the environment of the recovery target. The corresponding params below must be filled out.
		  * Constraints: Allowable values are: `kPhysical`.
	* `recover_volume_params` - (Optional, List) Specifies the parameters to recover Physical Volumes.
	Nested schema for **recover_volume_params**:
		* `physical_target_params` - (Optional, List) Specifies the params for recovering to a physical target.
		Nested schema for **physical_target_params**:
			* `force_unmount_volume` - (Optional, Boolean) Specifies whether volume would be dismounted first during LockVolume failure. If not specified, default is false.
			* `mount_target` - (Required, List) Specifies the target entity where the volumes are being mounted.
			Nested schema for **mount_target**:
				* `id` - (Required, Integer) Specifies the id of the object.
				* `name` - (Computed, String) Specifies the name of the object.
			* `vlan_config` - (Optional, List) Specifies VLAN Params associated with the recovered. If this is not specified, then the VLAN settings will be automatically selected from one of the below options: a. If VLANs are configured on Cohesity, then the VLAN host/VIP will be automatically based on the client's (e.g. ESXI host) IP address. b. If VLANs are not configured on Cohesity, then the partition hostname or VIPs will be used for Recovery.
			Nested schema for **vlan_config**:
				* `disable_vlan` - (Optional, Boolean) If this is set to true, then even if VLANs are configured on the system, the partition VIPs will be used for the Recovery.
				* `id` - (Optional, Integer) If this is set, then the Cohesity host name or the IP address associated with this vlan is used for mounting Cohesity's view on the remote host.
				* `interface_name` - (Computed, String) Interface group to use for Recovery.
			* `volume_mapping` - (Required, List) Specifies the mapping from source volumes to destination volumes.
			Nested schema for **volume_mapping**:
				* `destination_volume_guid` - (Required, String) Specifies the guid of the destination volume.
				* `source_volume_guid` - (Required, String) Specifies the guid of the source volume.
		* `target_environment` - (Required, String) Specifies the environment of the recovery target. The corresponding params below must be filled out.
		  * Constraints: Allowable values are: `kPhysical`.
	* `recovery_action` - (Required, String) Specifies the type of recover action to be performed.
	  * Constraints: Allowable values are: `RecoverPhysicalVolumes`, `InstantVolumeMount`, `RecoverFiles`, `RecoverSystem`.
	* `system_recovery_params` - (Optional, List) Specifies the parameters to perform a system recovery.
	Nested schema for **system_recovery_params**:
		* `full_nas_path` - (Optional, String) Specifies the path to the recovery view.
* `request_initiator_type` - (Optional, Forces new resource, String) Specifies the type of request from UI, which is used for services like magneto to determine the priority of requests.
  * Constraints: Allowable values are: `UIUser`, `UIAuto`, `Helios`.
* `snapshot_environment` - (Required, Forces new resource, String) Specifies the type of snapshot environment for which the Recovery was performed.
  * Constraints: Allowable values are: `kPhysical`, `kSQL`.
* `x_ibm_tenant_id` - (Required, Forces new resource, String) Specifies the key to be used to encrypt the source credential. If includeSourceCredentials is set to true this key must be specified.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the Common Recovery Response Params..
* `can_tear_down` - (Boolean) Specifies whether it's possible to tear down the objects created by the recovery.
* `recovery_id` - The unique identifier of Recovery
* `creation_info` - (List) Specifies the information about the creation of the protection group or recovery.
Nested schema for **creation_info**:
	* `user_name` - (String) Specifies the name of the user who created the protection group or recovery.
* `end_time_usecs` - (Integer) Specifies the end time of the Recovery in Unix timestamp epoch in microseconds. This field will be populated only after Recovery is finished.
* `is_multi_stage_restore` - (Boolean) Specifies whether the current recovery operation is a multi-stage restore operation. This is currently used by VMware recoveres for the migration/hot-standby use case.
* `is_parent_recovery` - (Boolean) Specifies whether the current recovery operation has created child recoveries. This is currently used in SQL recovery where multiple child recoveries can be tracked under a common/parent recovery.
* `messages` - (List) Specifies messages about the recovery.
* `parent_recovery_id` - (String) If current recovery is child recovery triggered by another parent recovery operation, then this field willt specify the id of the parent recovery.
  * Constraints: The value must match regular expression `/^\\d+:\\d+:\\d+$/`.
* `permissions` - (List) Specifies the list of tenants that have permissions for this recovery.
Nested schema for **permissions**:
	* `created_at_time_msecs` - (Integer) Epoch time when tenant was created.
	* `deleted_at_time_msecs` - (Integer) Epoch time when tenant was last updated.
	* `description` - (String) Description about the tenant.
	* `external_vendor_metadata` - (List) Specifies the additional metadata for the tenant that is specifically set by the external vendors who are responsible for managing tenants. This field will only applicable if tenant creation is happening for a specially provisioned clusters for external vendors.
	Nested schema for **external_vendor_metadata**:
		* `ibm_tenant_metadata_params` - (List) Specifies the additional metadata for the tenant that is specifically set by the external vendor of type 'IBM'.
		Nested schema for **ibm_tenant_metadata_params**:
			* `account_id` - (String) Specifies the unique identifier of the IBM's account ID.
			* `crn` - (String) Specifies the unique CRN associated with the tenant.
			* `custom_properties` - (List) Specifies the list of custom properties associated with the tenant. External vendors can choose to set any properties inside following list. Note that the fields set inside the following will not be available for direct filtering. API callers should make sure that no sensitive information such as passwords is sent in these fields.
			Nested schema for **custom_properties**:
				* `key` - (String) Specifies the unique key for custom property.
				* `value` - (String) Specifies the value for the above custom key.
			* `liveness_mode` - (String) Specifies the current liveness mode of the tenant. This mode may change based on AZ failures when vendor chooses to failover or failback the tenants to other AZs.
			  * Constraints: Allowable values are: `Active`, `Standby`.
			* `metrics_config` - (List) Specifies the metadata for metrics configuration. The metadata defined here will be used by cluster to send the usgae metrics to IBM cloud metering service for calculating the tenant billing.
			Nested schema for **metrics_config**:
				* `cos_resource_config` - (List) Specifies the details of COS resource configuration required for posting metrics and trackinb billing information for IBM tenants.
				Nested schema for **cos_resource_config**:
					* `resource_url` - (String) Specifies the resource COS resource configuration endpoint that will be used for fetching bucket usage for a given tenant.
				* `iam_metrics_config` - (List) Specifies the IAM configuration that will be used for accessing the billing service in IBM cloud.
				Nested schema for **iam_metrics_config**:
					* `billing_api_key_secret_id` - (String) Specifies Id of the secret that contains the API key.
					* `iam_url` - (String) Specifies the IAM URL needed to fetch the operator token from IBM. The operator token is needed to make service API calls to IBM billing service.
				* `metering_config` - (List) Specifies the metering configuration that will be used for IBM cluster to send the billing details to IBM billing service.
				Nested schema for **metering_config**:
					* `part_ids` - (List) Specifies the list of part identifiers used for metrics identification.
					  * Constraints: Allowable list items are: `USAGETERABYTE`. The minimum length is `1` item.
					* `submission_interval_in_secs` - (Integer) Specifies the frequency in seconds at which the metrics will be pushed to IBM billing service from cluster.
					* `url` - (String) Specifies the base metering URL that will be used by cluster to send the billing information.
			* `ownership_mode` - (String) Specifies the current ownership mode for the tenant. The ownership of the tenant represents the active role for functioning of the tenant.
			  * Constraints: Allowable values are: `Primary`, `Secondary`.
			* `plan_id` - (String) Specifies the Plan Id associated with the tenant. This field is introduced for tracking purposes inside IBM enviournment.
			* `resource_group_id` - (String) Specifies the Resource Group ID associated with the tenant.
			* `resource_instance_id` - (String) Specifies the Resource Instance ID associated with the tenant. This field is introduced for tracking purposes inside IBM enviournment.
		* `type` - (String) Specifies the type of the external vendor. The type specific parameters must be specified the provided type.
		  * Constraints: Allowable values are: `IBM`.
	* `id` - (String) The tenant id.
	* `is_managed_on_helios` - (Boolean) Flag to indicate if tenant is managed on helios.
	* `last_updated_at_time_msecs` - (Integer) Epoch time when tenant was last updated.
	* `name` - (String) Name of the Tenant.
	* `network` - (List) Networking information about a Tenant on a Cluster.
	Nested schema for **network**:
		* `cluster_hostname` - (String) The hostname for Cohesity cluster as seen by tenants and as is routable from the tenant's network. Tenant's VLAN's hostname, if available can be used instead but it is mandatory to provide this value if there's no VLAN hostname to use. Also, when set, this field would take precedence over VLAN hostname.
		* `cluster_ips` - (List) Set of IPs as seen from the tenant's network for the Cohesity cluster. Only one from 'clusterHostname' and 'clusterIps' is needed.
		* `connector_enabled` - (Boolean) Whether connector (hybrid extender) is enabled.
	* `status` - (String) Current Status of the Tenant.
	  * Constraints: Allowable values are: `Active`, `Inactive`, `MarkedForDeletion`, `Deleted`.
* `progress_task_id` - (String) Progress monitor task id for Recovery.
* `recovery_action` - (String) Specifies the type of recover action.
  * Constraints: Allowable values are: `RecoverVMs`, `RecoverFiles`, `InstantVolumeMount`, `RecoverVmDisks`, `RecoverVApps`, `RecoverVAppTemplates`, `UptierSnapshot`, `RecoverRDS`, `RecoverAurora`, `RecoverS3Buckets`, `RecoverRDSPostgres`, `RecoverAzureSQL`, `RecoverApps`, `CloneApps`, `RecoverNasVolume`, `RecoverPhysicalVolumes`, `RecoverSystem`, `RecoverExchangeDbs`, `CloneAppView`, `RecoverSanVolumes`, `RecoverSanGroup`, `RecoverMailbox`, `RecoverOneDrive`, `RecoverSharePoint`, `RecoverPublicFolders`, `RecoverMsGroup`, `RecoverMsTeam`, `ConvertToPst`, `DownloadChats`, `RecoverMailboxCSM`, `RecoverOneDriveCSM`, `RecoverSharePointCSM`, `RecoverNamespaces`, `RecoverObjects`, `RecoverSfdcObjects`, `RecoverSfdcOrg`, `RecoverSfdcRecords`, `DownloadFilesAndFolders`, `CloneVMs`, `CloneView`, `CloneRefreshApp`, `CloneVMsToView`, `ConvertAndDeployVMs`, `DeployVMs`.
* `retrieve_archive_tasks` - (List) Specifies the list of persistent state of a retrieve of an archive task.
Nested schema for **retrieve_archive_tasks**:
	* `task_uid` - (String) Specifies the globally unique id for this retrieval of an archive task.
	  * Constraints: The value must match regular expression `/^\\d+:\\d+:\\d+$/`.
	* `uptier_expiry_times` - (List) Specifies how much time the retrieved entity is present in the hot-tiers.
* `start_time_usecs` - (Integer) Specifies the start time of the Recovery in Unix timestamp epoch in microseconds.
* `status` - (String) Status of the Recovery. 'Running' indicates that the Recovery is still running. 'Canceled' indicates that the Recovery has been cancelled. 'Canceling' indicates that the Recovery is in the process of being cancelled. 'Failed' indicates that the Recovery has failed. 'Succeeded' indicates that the Recovery has finished successfully. 'SucceededWithWarning' indicates that the Recovery finished successfully, but there were some warning messages. 'Skipped' indicates that the Recovery task was skipped.
  * Constraints: Allowable values are: `Accepted`, `Running`, `Canceled`, `Canceling`, `Failed`, `Missed`, `Succeeded`, `SucceededWithWarning`, `OnHold`, `Finalizing`, `Skipped`, `LegalHold`.
* `tear_down_message` - (String) Specifies the error message about the tear down operation if it fails.
* `tear_down_status` - (String) Specifies the status of the tear down operation. This is only set when the canTearDown is set to true. 'DestroyScheduled' indicates that the tear down is ready to schedule. 'Destroying' indicates that the tear down is still running. 'Destroyed' indicates that the tear down succeeded. 'DestroyError' indicates that the tear down failed.
  * Constraints: Allowable values are: `DestroyScheduled`, `Destroying`, `Destroyed`, `DestroyError`.


## Import

You can import the `ibm_backup_recovery` resource by using `id`. Specifies the id of the Recovery.The ID is formed using tenantID and resourceId.
`id = <tenantId>::<recovery_id>`. 


#### Syntax
```
import {
	to = <ibm_backup_recovery_resource>
	id = "<tenantId>::<recovery_id>"
}
```

#### Example
```
resource "ibm_backup_recovery" "backup_recovery_instance" {
		x_ibm_tenant_id = "jhxqx715r9/"
		snapshot_environment = "kPhysical"
		name = "terra-recovery-1"
		physical_params {
		  recovery_action = "RecoverFiles"
		  objects {
			snapshot_id = data.ibm_backup_recovery_object_snapshots.object_snapshot.snapshots.0.id
		  }
		  recover_file_and_folder_params {
			 target_environment = "kPhysical"
			 files_and_folders {
			   absolute_path = "/data/"
			 }
			 physical_target_params {
			   recover_target {
				 id = 3
			   }
			   restore_entity_type = "kRegular"
			   alternate_restore_directory = "/data/"
			 }
		  }
		}
	  }

import {
	to = ibm_backup_recovery.backup_recovery_instance
	id = "jhxqx715r9/::5170815044477768:1732541085048:484"
}
```
