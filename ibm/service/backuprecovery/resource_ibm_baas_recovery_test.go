// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package backuprecovery_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.ibm.com/BackupAndRecovery/ibm-backup-recovery-sdk-go/backuprecoveryv1"
)

func TestAccIbmBaasRecoveryBasic(t *testing.T) {
	var conf backuprecoveryv1.Recovery
	xIbmTenantID := fmt.Sprintf("tf_x_ibm_tenant_id_%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	snapshotEnvironment := "kPhysical"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmBaasRecoveryDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmBaasRecoveryConfigBasic(xIbmTenantID, name, snapshotEnvironment),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmBaasRecoveryExists("ibm_baas_recovery.baas_recovery_instance", conf),
					resource.TestCheckResourceAttr("ibm_baas_recovery.baas_recovery_instance", "x_ibm_tenant_id", xIbmTenantID),
					resource.TestCheckResourceAttr("ibm_baas_recovery.baas_recovery_instance", "name", name),
					resource.TestCheckResourceAttr("ibm_baas_recovery.baas_recovery_instance", "snapshot_environment", snapshotEnvironment),
				),
			},
		},
	})
}

func TestAccIbmBaasRecoveryAllArgs(t *testing.T) {
	var conf backuprecoveryv1.Recovery
	xIbmTenantID := fmt.Sprintf("tf_x_ibm_tenant_id_%d", acctest.RandIntRange(10, 100))
	requestInitiatorType := "UIUser"
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	snapshotEnvironment := "kPhysical"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmBaasRecoveryDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmBaasRecoveryConfig(xIbmTenantID, requestInitiatorType, name, snapshotEnvironment),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmBaasRecoveryExists("ibm_baas_recovery.baas_recovery_instance", conf),
					resource.TestCheckResourceAttr("ibm_baas_recovery.baas_recovery_instance", "x_ibm_tenant_id", xIbmTenantID),
					resource.TestCheckResourceAttr("ibm_baas_recovery.baas_recovery_instance", "request_initiator_type", requestInitiatorType),
					resource.TestCheckResourceAttr("ibm_baas_recovery.baas_recovery_instance", "name", name),
					resource.TestCheckResourceAttr("ibm_baas_recovery.baas_recovery_instance", "snapshot_environment", snapshotEnvironment),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_baas_recovery.baas_recovery",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmBaasRecoveryConfigBasic(xIbmTenantID string, name string, snapshotEnvironment string) string {
	return fmt.Sprintf(`
		resource "ibm_baas_recovery" "baas_recovery_instance" {
			x_ibm_tenant_id = "%s"
			name = "%s"
			snapshot_environment = "%s"
		}
	`, xIbmTenantID, name, snapshotEnvironment)
}

func testAccCheckIbmBaasRecoveryConfig(xIbmTenantID string, requestInitiatorType string, name string, snapshotEnvironment string) string {
	return fmt.Sprintf(`

		resource "ibm_baas_recovery" "baas_recovery_instance" {
			x_ibm_tenant_id = "%s"
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
					storage_domain_id = 1
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
					storage_domain_id = 1
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
		}
	`, xIbmTenantID, requestInitiatorType, name, snapshotEnvironment)
}

func testAccCheckIbmBaasRecoveryExists(n string, obj backuprecoveryv1.Recovery) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		backupRecoveryClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).BackupRecoveryV1()
		if err != nil {
			return err
		}

		getRecoveryByIdOptions := &backuprecoveryv1.GetRecoveryByIdOptions{}

		getRecoveryByIdOptions.SetID(rs.Primary.ID)

		recovery, _, err := backupRecoveryClient.GetRecoveryByID(getRecoveryByIdOptions)
		if err != nil {
			return err
		}

		obj = *recovery
		return nil
	}
}

func testAccCheckIbmBaasRecoveryDestroy(s *terraform.State) error {
	backupRecoveryClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).BackupRecoveryV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_baas_recovery" {
			continue
		}

		getRecoveryByIdOptions := &backuprecoveryv1.GetRecoveryByIdOptions{}

		getRecoveryByIdOptions.SetID(rs.Primary.ID)

		// Try to find the key
		_, response, err := backupRecoveryClient.GetRecoveryByID(getRecoveryByIdOptions)

		if err == nil {
			return fmt.Errorf("Common Recovery Response Params. still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for Common Recovery Response Params. (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
