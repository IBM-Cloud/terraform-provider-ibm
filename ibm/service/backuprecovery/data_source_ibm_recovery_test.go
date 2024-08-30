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
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/backuprecovery"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/stretchr/testify/assert"
	"github.ibm.com/BackupAndRecovery/ibm-backup-recovery-sdk-go/backuprecoveryv1"
)

func TestAccIbmRecoveryDataSourceBasic(t *testing.T) {
	recoveryName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	recoverySnapshotEnvironment := "kPhysical"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmRecoveryDataSourceConfigBasic(recoveryName, recoverySnapshotEnvironment),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_recovery.recovery_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_recovery.recovery_instance", "recovery_id"),
				),
			},
		},
	})
}

func TestAccIbmRecoveryDataSourceAllArgs(t *testing.T) {
	recoveryRequestInitiatorType := "UIUser"
	recoveryName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	recoverySnapshotEnvironment := "kPhysical"

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmRecoveryDataSourceConfig(recoveryRequestInitiatorType, recoveryName, recoverySnapshotEnvironment),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_recovery.recovery_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_recovery.recovery_instance", "recovery_id"),
					resource.TestCheckResourceAttrSet("data.ibm_recovery.recovery_instance", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_recovery.recovery_instance", "start_time_usecs"),
					resource.TestCheckResourceAttrSet("data.ibm_recovery.recovery_instance", "end_time_usecs"),
					resource.TestCheckResourceAttrSet("data.ibm_recovery.recovery_instance", "status"),
					resource.TestCheckResourceAttrSet("data.ibm_recovery.recovery_instance", "progress_task_id"),
					resource.TestCheckResourceAttrSet("data.ibm_recovery.recovery_instance", "snapshot_environment"),
					resource.TestCheckResourceAttrSet("data.ibm_recovery.recovery_instance", "recovery_action"),
					resource.TestCheckResourceAttrSet("data.ibm_recovery.recovery_instance", "permissions.#"),
					resource.TestCheckResourceAttrSet("data.ibm_recovery.recovery_instance", "permissions.0.created_at_time_msecs"),
					resource.TestCheckResourceAttrSet("data.ibm_recovery.recovery_instance", "permissions.0.deleted_at_time_msecs"),
					resource.TestCheckResourceAttrSet("data.ibm_recovery.recovery_instance", "permissions.0.description"),
					resource.TestCheckResourceAttrSet("data.ibm_recovery.recovery_instance", "permissions.0.id"),
					resource.TestCheckResourceAttrSet("data.ibm_recovery.recovery_instance", "permissions.0.is_managed_on_helios"),
					resource.TestCheckResourceAttrSet("data.ibm_recovery.recovery_instance", "permissions.0.last_updated_at_time_msecs"),
					resource.TestCheckResourceAttr("data.ibm_recovery.recovery_instance", "permissions.0.name", recoveryName),
					resource.TestCheckResourceAttrSet("data.ibm_recovery.recovery_instance", "permissions.0.status"),
					resource.TestCheckResourceAttrSet("data.ibm_recovery.recovery_instance", "creation_info.#"),
					resource.TestCheckResourceAttrSet("data.ibm_recovery.recovery_instance", "can_tear_down"),
					resource.TestCheckResourceAttrSet("data.ibm_recovery.recovery_instance", "tear_down_status"),
					resource.TestCheckResourceAttrSet("data.ibm_recovery.recovery_instance", "tear_down_message"),
					resource.TestCheckResourceAttrSet("data.ibm_recovery.recovery_instance", "messages.#"),
					resource.TestCheckResourceAttrSet("data.ibm_recovery.recovery_instance", "is_parent_recovery"),
					resource.TestCheckResourceAttrSet("data.ibm_recovery.recovery_instance", "parent_recovery_id"),
					resource.TestCheckResourceAttrSet("data.ibm_recovery.recovery_instance", "retrieve_archive_tasks.#"),
					resource.TestCheckResourceAttrSet("data.ibm_recovery.recovery_instance", "retrieve_archive_tasks.0.task_uid"),
					resource.TestCheckResourceAttrSet("data.ibm_recovery.recovery_instance", "is_multi_stage_restore"),
					resource.TestCheckResourceAttrSet("data.ibm_recovery.recovery_instance", "physical_params.#"),
					resource.TestCheckResourceAttrSet("data.ibm_recovery.recovery_instance", "mssql_params.#"),
				),
			},
		},
	})
}

func testAccCheckIbmRecoveryDataSourceConfigBasic(recoveryName string, recoverySnapshotEnvironment string) string {
	return fmt.Sprintf(`
		resource "ibm_recovery" "recovery_instance" {
			name = "%s"
			snapshot_environment = "%s"
		}

		data "ibm_recovery" "recovery_instance" {
			recovery_id = "recovery_id"
		}
	`, recoveryName, recoverySnapshotEnvironment)
}

func testAccCheckIbmRecoveryDataSourceConfig(recoveryRequestInitiatorType string, recoveryName string, recoverySnapshotEnvironment string) string {
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

		data "ibm_recovery" "recovery_instance" {
			recovery_id = "recovery_id"
		}
	`, recoveryRequestInitiatorType, recoveryName, recoverySnapshotEnvironment)
}

func TestDataSourceIbmRecoveryTenantToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		externalVendorCustomPropertiesModel := make(map[string]interface{})
		externalVendorCustomPropertiesModel["key"] = "testString"
		externalVendorCustomPropertiesModel["value"] = "testString"

		ibmTenantMetadataParamsModel := make(map[string]interface{})
		ibmTenantMetadataParamsModel["account_id"] = "testString"
		ibmTenantMetadataParamsModel["crn"] = "testString"
		ibmTenantMetadataParamsModel["custom_properties"] = []map[string]interface{}{externalVendorCustomPropertiesModel}
		ibmTenantMetadataParamsModel["liveness_mode"] = "Active"
		ibmTenantMetadataParamsModel["ownership_mode"] = "Primary"
		ibmTenantMetadataParamsModel["resource_group_id"] = "testString"

		externalVendorTenantMetadataModel := make(map[string]interface{})
		externalVendorTenantMetadataModel["ibm_tenant_metadata_params"] = []map[string]interface{}{ibmTenantMetadataParamsModel}
		externalVendorTenantMetadataModel["type"] = "IBM"

		tenantNetworkModel := make(map[string]interface{})
		tenantNetworkModel["connector_enabled"] = true
		tenantNetworkModel["cluster_hostname"] = "testString"
		tenantNetworkModel["cluster_ips"] = []string{"testString"}

		model := make(map[string]interface{})
		model["created_at_time_msecs"] = int(26)
		model["deleted_at_time_msecs"] = int(26)
		model["description"] = "testString"
		model["external_vendor_metadata"] = []map[string]interface{}{externalVendorTenantMetadataModel}
		model["id"] = "testString"
		model["is_managed_on_helios"] = true
		model["last_updated_at_time_msecs"] = int(26)
		model["name"] = "testString"
		model["network"] = []map[string]interface{}{tenantNetworkModel}
		model["status"] = "Active"

		assert.Equal(t, result, model)
	}

	externalVendorCustomPropertiesModel := new(backuprecoveryv1.ExternalVendorCustomProperties)
	externalVendorCustomPropertiesModel.Key = core.StringPtr("testString")
	externalVendorCustomPropertiesModel.Value = core.StringPtr("testString")

	ibmTenantMetadataParamsModel := new(backuprecoveryv1.IbmTenantMetadataParams)
	ibmTenantMetadataParamsModel.AccountID = core.StringPtr("testString")
	ibmTenantMetadataParamsModel.Crn = core.StringPtr("testString")
	ibmTenantMetadataParamsModel.CustomProperties = []backuprecoveryv1.ExternalVendorCustomProperties{*externalVendorCustomPropertiesModel}
	ibmTenantMetadataParamsModel.LivenessMode = core.StringPtr("Active")
	ibmTenantMetadataParamsModel.OwnershipMode = core.StringPtr("Primary")
	ibmTenantMetadataParamsModel.ResourceGroupID = core.StringPtr("testString")

	externalVendorTenantMetadataModel := new(backuprecoveryv1.ExternalVendorTenantMetadata)
	externalVendorTenantMetadataModel.IbmTenantMetadataParams = ibmTenantMetadataParamsModel
	externalVendorTenantMetadataModel.Type = core.StringPtr("IBM")

	tenantNetworkModel := new(backuprecoveryv1.TenantNetwork)
	tenantNetworkModel.ConnectorEnabled = core.BoolPtr(true)
	tenantNetworkModel.ClusterHostname = core.StringPtr("testString")
	tenantNetworkModel.ClusterIps = []string{"testString"}

	model := new(backuprecoveryv1.Tenant)
	model.CreatedAtTimeMsecs = core.Int64Ptr(int64(26))
	model.DeletedAtTimeMsecs = core.Int64Ptr(int64(26))
	model.Description = core.StringPtr("testString")
	model.ExternalVendorMetadata = externalVendorTenantMetadataModel
	model.ID = core.StringPtr("testString")
	model.IsManagedOnHelios = core.BoolPtr(true)
	model.LastUpdatedAtTimeMsecs = core.Int64Ptr(int64(26))
	model.Name = core.StringPtr("testString")
	model.Network = tenantNetworkModel
	model.Status = core.StringPtr("Active")

	result, err := backuprecovery.DataSourceIbmRecoveryTenantToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmRecoveryExternalVendorTenantMetadataToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		externalVendorCustomPropertiesModel := make(map[string]interface{})
		externalVendorCustomPropertiesModel["key"] = "testString"
		externalVendorCustomPropertiesModel["value"] = "testString"

		ibmTenantMetadataParamsModel := make(map[string]interface{})
		ibmTenantMetadataParamsModel["account_id"] = "testString"
		ibmTenantMetadataParamsModel["crn"] = "testString"
		ibmTenantMetadataParamsModel["custom_properties"] = []map[string]interface{}{externalVendorCustomPropertiesModel}
		ibmTenantMetadataParamsModel["liveness_mode"] = "Active"
		ibmTenantMetadataParamsModel["ownership_mode"] = "Primary"
		ibmTenantMetadataParamsModel["resource_group_id"] = "testString"

		model := make(map[string]interface{})
		model["ibm_tenant_metadata_params"] = []map[string]interface{}{ibmTenantMetadataParamsModel}
		model["type"] = "IBM"

		assert.Equal(t, result, model)
	}

	externalVendorCustomPropertiesModel := new(backuprecoveryv1.ExternalVendorCustomProperties)
	externalVendorCustomPropertiesModel.Key = core.StringPtr("testString")
	externalVendorCustomPropertiesModel.Value = core.StringPtr("testString")

	ibmTenantMetadataParamsModel := new(backuprecoveryv1.IbmTenantMetadataParams)
	ibmTenantMetadataParamsModel.AccountID = core.StringPtr("testString")
	ibmTenantMetadataParamsModel.Crn = core.StringPtr("testString")
	ibmTenantMetadataParamsModel.CustomProperties = []backuprecoveryv1.ExternalVendorCustomProperties{*externalVendorCustomPropertiesModel}
	ibmTenantMetadataParamsModel.LivenessMode = core.StringPtr("Active")
	ibmTenantMetadataParamsModel.OwnershipMode = core.StringPtr("Primary")
	ibmTenantMetadataParamsModel.ResourceGroupID = core.StringPtr("testString")

	model := new(backuprecoveryv1.ExternalVendorTenantMetadata)
	model.IbmTenantMetadataParams = ibmTenantMetadataParamsModel
	model.Type = core.StringPtr("IBM")

	result, err := backuprecovery.DataSourceIbmRecoveryExternalVendorTenantMetadataToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmRecoveryIbmTenantMetadataParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		externalVendorCustomPropertiesModel := make(map[string]interface{})
		externalVendorCustomPropertiesModel["key"] = "testString"
		externalVendorCustomPropertiesModel["value"] = "testString"

		model := make(map[string]interface{})
		model["account_id"] = "testString"
		model["crn"] = "testString"
		model["custom_properties"] = []map[string]interface{}{externalVendorCustomPropertiesModel}
		model["liveness_mode"] = "Active"
		model["ownership_mode"] = "Primary"
		model["resource_group_id"] = "testString"

		assert.Equal(t, result, model)
	}

	externalVendorCustomPropertiesModel := new(backuprecoveryv1.ExternalVendorCustomProperties)
	externalVendorCustomPropertiesModel.Key = core.StringPtr("testString")
	externalVendorCustomPropertiesModel.Value = core.StringPtr("testString")

	model := new(backuprecoveryv1.IbmTenantMetadataParams)
	model.AccountID = core.StringPtr("testString")
	model.Crn = core.StringPtr("testString")
	model.CustomProperties = []backuprecoveryv1.ExternalVendorCustomProperties{*externalVendorCustomPropertiesModel}
	model.LivenessMode = core.StringPtr("Active")
	model.OwnershipMode = core.StringPtr("Primary")
	model.ResourceGroupID = core.StringPtr("testString")

	result, err := backuprecovery.DataSourceIbmRecoveryIbmTenantMetadataParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmRecoveryExternalVendorCustomPropertiesToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["key"] = "testString"
		model["value"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.ExternalVendorCustomProperties)
	model.Key = core.StringPtr("testString")
	model.Value = core.StringPtr("testString")

	result, err := backuprecovery.DataSourceIbmRecoveryExternalVendorCustomPropertiesToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmRecoveryTenantNetworkToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["connector_enabled"] = true
		model["cluster_hostname"] = "testString"
		model["cluster_ips"] = []string{"testString"}

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.TenantNetwork)
	model.ConnectorEnabled = core.BoolPtr(true)
	model.ClusterHostname = core.StringPtr("testString")
	model.ClusterIps = []string{"testString"}

	result, err := backuprecovery.DataSourceIbmRecoveryTenantNetworkToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmRecoveryCreationInfoToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["user_name"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.CreationInfo)
	model.UserName = core.StringPtr("testString")

	result, err := backuprecovery.DataSourceIbmRecoveryCreationInfoToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmRecoveryRetrieveArchiveTaskToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["task_uid"] = "testString"
		model["uptier_expiry_times"] = []int64{int64(26)}

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.RetrieveArchiveTask)
	model.TaskUid = core.StringPtr("testString")
	model.UptierExpiryTimes = []int64{int64(26)}

	result, err := backuprecovery.DataSourceIbmRecoveryRetrieveArchiveTaskToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmRecoveryRecoverPhysicalParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		sharepointObjectParamsModel := make(map[string]interface{})
		sharepointObjectParamsModel["site_web_url"] = "testString"

		objectTypeVCenterParamsModel := make(map[string]interface{})
		objectTypeVCenterParamsModel["is_cloud_env"] = true

		objectTypeWindowsClusterParamsModel := make(map[string]interface{})
		objectTypeWindowsClusterParamsModel["cluster_source_type"] = "testString"

		objectSummaryModel := make(map[string]interface{})
		objectSummaryModel["id"] = int(26)
		objectSummaryModel["name"] = "testString"
		objectSummaryModel["source_id"] = int(26)
		objectSummaryModel["source_name"] = "testString"
		objectSummaryModel["environment"] = "kPhysical"
		objectSummaryModel["object_hash"] = "testString"
		objectSummaryModel["object_type"] = "kCluster"
		objectSummaryModel["logical_size_bytes"] = int(26)
		objectSummaryModel["uuid"] = "testString"
		objectSummaryModel["global_id"] = "testString"
		objectSummaryModel["protection_type"] = "kAgent"
		objectSummaryModel["sharepoint_site_summary"] = []map[string]interface{}{sharepointObjectParamsModel}
		objectSummaryModel["os_type"] = "kLinux"
		objectSummaryModel["v_center_summary"] = []map[string]interface{}{objectTypeVCenterParamsModel}
		objectSummaryModel["windows_cluster_summary"] = []map[string]interface{}{objectTypeWindowsClusterParamsModel}

		commonRecoverObjectSnapshotParamsObjectInfoModel := make(map[string]interface{})
		commonRecoverObjectSnapshotParamsObjectInfoModel["id"] = int(26)
		commonRecoverObjectSnapshotParamsObjectInfoModel["name"] = "testString"
		commonRecoverObjectSnapshotParamsObjectInfoModel["source_id"] = int(26)
		commonRecoverObjectSnapshotParamsObjectInfoModel["source_name"] = "testString"
		commonRecoverObjectSnapshotParamsObjectInfoModel["environment"] = "kPhysical"
		commonRecoverObjectSnapshotParamsObjectInfoModel["object_hash"] = "testString"
		commonRecoverObjectSnapshotParamsObjectInfoModel["object_type"] = "kCluster"
		commonRecoverObjectSnapshotParamsObjectInfoModel["logical_size_bytes"] = int(26)
		commonRecoverObjectSnapshotParamsObjectInfoModel["uuid"] = "testString"
		commonRecoverObjectSnapshotParamsObjectInfoModel["global_id"] = "testString"
		commonRecoverObjectSnapshotParamsObjectInfoModel["protection_type"] = "kAgent"
		commonRecoverObjectSnapshotParamsObjectInfoModel["sharepoint_site_summary"] = []map[string]interface{}{sharepointObjectParamsModel}
		commonRecoverObjectSnapshotParamsObjectInfoModel["os_type"] = "kLinux"
		commonRecoverObjectSnapshotParamsObjectInfoModel["child_objects"] = []map[string]interface{}{objectSummaryModel}
		commonRecoverObjectSnapshotParamsObjectInfoModel["v_center_summary"] = []map[string]interface{}{objectTypeVCenterParamsModel}
		commonRecoverObjectSnapshotParamsObjectInfoModel["windows_cluster_summary"] = []map[string]interface{}{objectTypeWindowsClusterParamsModel}

		awsTierModel := make(map[string]interface{})
		awsTierModel["move_after_unit"] = "Days"
		awsTierModel["move_after"] = int(26)
		awsTierModel["tier_type"] = "kAmazonS3Standard"

		awsTiersModel := make(map[string]interface{})
		awsTiersModel["tiers"] = []map[string]interface{}{awsTierModel}

		azureTierModel := make(map[string]interface{})
		azureTierModel["move_after_unit"] = "Days"
		azureTierModel["move_after"] = int(26)
		azureTierModel["tier_type"] = "kAzureTierHot"

		azureTiersModel := make(map[string]interface{})
		azureTiersModel["tiers"] = []map[string]interface{}{azureTierModel}

		googleTierModel := make(map[string]interface{})
		googleTierModel["move_after_unit"] = "Days"
		googleTierModel["move_after"] = int(26)
		googleTierModel["tier_type"] = "kGoogleStandard"

		googleTiersModel := make(map[string]interface{})
		googleTiersModel["tiers"] = []map[string]interface{}{googleTierModel}

		oracleTierModel := make(map[string]interface{})
		oracleTierModel["move_after_unit"] = "Days"
		oracleTierModel["move_after"] = int(26)
		oracleTierModel["tier_type"] = "kOracleTierStandard"

		oracleTiersModel := make(map[string]interface{})
		oracleTiersModel["tiers"] = []map[string]interface{}{oracleTierModel}

		archivalTargetTierInfoModel := make(map[string]interface{})
		archivalTargetTierInfoModel["aws_tiering"] = []map[string]interface{}{awsTiersModel}
		archivalTargetTierInfoModel["azure_tiering"] = []map[string]interface{}{azureTiersModel}
		archivalTargetTierInfoModel["cloud_platform"] = "AWS"
		archivalTargetTierInfoModel["google_tiering"] = []map[string]interface{}{googleTiersModel}
		archivalTargetTierInfoModel["oracle_tiering"] = []map[string]interface{}{oracleTiersModel}
		archivalTargetTierInfoModel["current_tier_type"] = "kAmazonS3Standard"

		commonRecoverObjectSnapshotParamsArchivalTargetInfoModel := make(map[string]interface{})
		commonRecoverObjectSnapshotParamsArchivalTargetInfoModel["target_id"] = int(26)
		commonRecoverObjectSnapshotParamsArchivalTargetInfoModel["archival_task_id"] = "testString"
		commonRecoverObjectSnapshotParamsArchivalTargetInfoModel["target_name"] = "testString"
		commonRecoverObjectSnapshotParamsArchivalTargetInfoModel["target_type"] = "Tape"
		commonRecoverObjectSnapshotParamsArchivalTargetInfoModel["usage_type"] = "Archival"
		commonRecoverObjectSnapshotParamsArchivalTargetInfoModel["ownership_context"] = "Local"
		commonRecoverObjectSnapshotParamsArchivalTargetInfoModel["tier_settings"] = []map[string]interface{}{archivalTargetTierInfoModel}

		commonRecoverObjectSnapshotParamsModel := make(map[string]interface{})
		commonRecoverObjectSnapshotParamsModel["snapshot_id"] = "testString"
		commonRecoverObjectSnapshotParamsModel["point_in_time_usecs"] = int(26)
		commonRecoverObjectSnapshotParamsModel["protection_group_id"] = "testString"
		commonRecoverObjectSnapshotParamsModel["protection_group_name"] = "testString"
		commonRecoverObjectSnapshotParamsModel["object_info"] = []map[string]interface{}{commonRecoverObjectSnapshotParamsObjectInfoModel}
		commonRecoverObjectSnapshotParamsModel["archival_target_info"] = []map[string]interface{}{commonRecoverObjectSnapshotParamsArchivalTargetInfoModel}
		commonRecoverObjectSnapshotParamsModel["recover_from_standby"] = true

		physicalTargetParamsForRecoverVolumeMountTargetModel := make(map[string]interface{})
		physicalTargetParamsForRecoverVolumeMountTargetModel["id"] = int(26)

		recoverVolumeMappingModel := make(map[string]interface{})
		recoverVolumeMappingModel["source_volume_guid"] = "testString"
		recoverVolumeMappingModel["destination_volume_guid"] = "testString"

		physicalTargetParamsForRecoverVolumeVlanConfigModel := make(map[string]interface{})
		physicalTargetParamsForRecoverVolumeVlanConfigModel["id"] = int(38)
		physicalTargetParamsForRecoverVolumeVlanConfigModel["disable_vlan"] = true

		recoverPhysicalVolumeParamsPhysicalTargetParamsModel := make(map[string]interface{})
		recoverPhysicalVolumeParamsPhysicalTargetParamsModel["mount_target"] = []map[string]interface{}{physicalTargetParamsForRecoverVolumeMountTargetModel}
		recoverPhysicalVolumeParamsPhysicalTargetParamsModel["volume_mapping"] = []map[string]interface{}{recoverVolumeMappingModel}
		recoverPhysicalVolumeParamsPhysicalTargetParamsModel["force_unmount_volume"] = true
		recoverPhysicalVolumeParamsPhysicalTargetParamsModel["vlan_config"] = []map[string]interface{}{physicalTargetParamsForRecoverVolumeVlanConfigModel}

		recoverPhysicalParamsRecoverVolumeParamsModel := make(map[string]interface{})
		recoverPhysicalParamsRecoverVolumeParamsModel["target_environment"] = "kPhysical"
		recoverPhysicalParamsRecoverVolumeParamsModel["physical_target_params"] = []map[string]interface{}{recoverPhysicalVolumeParamsPhysicalTargetParamsModel}

		physicalMountVolumesOriginalTargetConfigServerCredentialsModel := make(map[string]interface{})
		physicalMountVolumesOriginalTargetConfigServerCredentialsModel["username"] = "testString"
		physicalMountVolumesOriginalTargetConfigServerCredentialsModel["password"] = "testString"

		physicalTargetParamsForMountVolumeOriginalTargetConfigModel := make(map[string]interface{})
		physicalTargetParamsForMountVolumeOriginalTargetConfigModel["server_credentials"] = []map[string]interface{}{physicalMountVolumesOriginalTargetConfigServerCredentialsModel}

		recoverTargetModel := make(map[string]interface{})
		recoverTargetModel["id"] = int(26)

		physicalMountVolumesNewTargetConfigServerCredentialsModel := make(map[string]interface{})
		physicalMountVolumesNewTargetConfigServerCredentialsModel["username"] = "testString"
		physicalMountVolumesNewTargetConfigServerCredentialsModel["password"] = "testString"

		physicalTargetParamsForMountVolumeNewTargetConfigModel := make(map[string]interface{})
		physicalTargetParamsForMountVolumeNewTargetConfigModel["mount_target"] = []map[string]interface{}{recoverTargetModel}
		physicalTargetParamsForMountVolumeNewTargetConfigModel["server_credentials"] = []map[string]interface{}{physicalMountVolumesNewTargetConfigServerCredentialsModel}

		physicalTargetParamsForMountVolumeVlanConfigModel := make(map[string]interface{})
		physicalTargetParamsForMountVolumeVlanConfigModel["id"] = int(38)
		physicalTargetParamsForMountVolumeVlanConfigModel["disable_vlan"] = true

		mountPhysicalVolumeParamsPhysicalTargetParamsModel := make(map[string]interface{})
		mountPhysicalVolumeParamsPhysicalTargetParamsModel["mount_to_original_target"] = true
		mountPhysicalVolumeParamsPhysicalTargetParamsModel["original_target_config"] = []map[string]interface{}{physicalTargetParamsForMountVolumeOriginalTargetConfigModel}
		mountPhysicalVolumeParamsPhysicalTargetParamsModel["new_target_config"] = []map[string]interface{}{physicalTargetParamsForMountVolumeNewTargetConfigModel}
		mountPhysicalVolumeParamsPhysicalTargetParamsModel["read_only_mount"] = true
		mountPhysicalVolumeParamsPhysicalTargetParamsModel["volume_names"] = []string{"testString"}
		mountPhysicalVolumeParamsPhysicalTargetParamsModel["vlan_config"] = []map[string]interface{}{physicalTargetParamsForMountVolumeVlanConfigModel}

		recoverPhysicalParamsMountVolumeParamsModel := make(map[string]interface{})
		recoverPhysicalParamsMountVolumeParamsModel["target_environment"] = "kPhysical"
		recoverPhysicalParamsMountVolumeParamsModel["physical_target_params"] = []map[string]interface{}{mountPhysicalVolumeParamsPhysicalTargetParamsModel}

		commonRecoverFileAndFolderInfoModel := make(map[string]interface{})
		commonRecoverFileAndFolderInfoModel["absolute_path"] = "testString"
		commonRecoverFileAndFolderInfoModel["is_directory"] = true
		commonRecoverFileAndFolderInfoModel["is_view_file_recovery"] = true

		physicalTargetParamsForRecoverFileAndFolderRecoverTargetModel := make(map[string]interface{})
		physicalTargetParamsForRecoverFileAndFolderRecoverTargetModel["id"] = int(26)

		physicalTargetParamsForRecoverFileAndFolderVlanConfigModel := make(map[string]interface{})
		physicalTargetParamsForRecoverFileAndFolderVlanConfigModel["id"] = int(38)
		physicalTargetParamsForRecoverFileAndFolderVlanConfigModel["disable_vlan"] = true

		recoverPhysicalFileAndFolderParamsPhysicalTargetParamsModel := make(map[string]interface{})
		recoverPhysicalFileAndFolderParamsPhysicalTargetParamsModel["recover_target"] = []map[string]interface{}{physicalTargetParamsForRecoverFileAndFolderRecoverTargetModel}
		recoverPhysicalFileAndFolderParamsPhysicalTargetParamsModel["restore_to_original_paths"] = true
		recoverPhysicalFileAndFolderParamsPhysicalTargetParamsModel["overwrite_existing"] = true
		recoverPhysicalFileAndFolderParamsPhysicalTargetParamsModel["alternate_restore_directory"] = "testString"
		recoverPhysicalFileAndFolderParamsPhysicalTargetParamsModel["preserve_attributes"] = true
		recoverPhysicalFileAndFolderParamsPhysicalTargetParamsModel["preserve_timestamps"] = true
		recoverPhysicalFileAndFolderParamsPhysicalTargetParamsModel["preserve_acls"] = true
		recoverPhysicalFileAndFolderParamsPhysicalTargetParamsModel["continue_on_error"] = true
		recoverPhysicalFileAndFolderParamsPhysicalTargetParamsModel["save_success_files"] = true
		recoverPhysicalFileAndFolderParamsPhysicalTargetParamsModel["vlan_config"] = []map[string]interface{}{physicalTargetParamsForRecoverFileAndFolderVlanConfigModel}
		recoverPhysicalFileAndFolderParamsPhysicalTargetParamsModel["restore_entity_type"] = "kRegular"

		recoverPhysicalParamsRecoverFileAndFolderParamsModel := make(map[string]interface{})
		recoverPhysicalParamsRecoverFileAndFolderParamsModel["files_and_folders"] = []map[string]interface{}{commonRecoverFileAndFolderInfoModel}
		recoverPhysicalParamsRecoverFileAndFolderParamsModel["target_environment"] = "kPhysical"
		recoverPhysicalParamsRecoverFileAndFolderParamsModel["physical_target_params"] = []map[string]interface{}{recoverPhysicalFileAndFolderParamsPhysicalTargetParamsModel}

		recoverPhysicalParamsDownloadFileAndFolderParamsModel := make(map[string]interface{})
		recoverPhysicalParamsDownloadFileAndFolderParamsModel["expiry_time_usecs"] = int(26)
		recoverPhysicalParamsDownloadFileAndFolderParamsModel["files_and_folders"] = []map[string]interface{}{commonRecoverFileAndFolderInfoModel}
		recoverPhysicalParamsDownloadFileAndFolderParamsModel["download_file_path"] = "testString"

		recoverPhysicalParamsSystemRecoveryParamsModel := make(map[string]interface{})
		recoverPhysicalParamsSystemRecoveryParamsModel["full_nas_path"] = "testString"

		model := make(map[string]interface{})
		model["objects"] = []map[string]interface{}{commonRecoverObjectSnapshotParamsModel}
		model["recovery_action"] = "RecoverPhysicalVolumes"
		model["recover_volume_params"] = []map[string]interface{}{recoverPhysicalParamsRecoverVolumeParamsModel}
		model["mount_volume_params"] = []map[string]interface{}{recoverPhysicalParamsMountVolumeParamsModel}
		model["recover_file_and_folder_params"] = []map[string]interface{}{recoverPhysicalParamsRecoverFileAndFolderParamsModel}
		model["download_file_and_folder_params"] = []map[string]interface{}{recoverPhysicalParamsDownloadFileAndFolderParamsModel}
		model["system_recovery_params"] = []map[string]interface{}{recoverPhysicalParamsSystemRecoveryParamsModel}

		assert.Equal(t, result, model)
	}

	sharepointObjectParamsModel := new(backuprecoveryv1.SharepointObjectParams)
	sharepointObjectParamsModel.SiteWebURL = core.StringPtr("testString")

	objectTypeVCenterParamsModel := new(backuprecoveryv1.ObjectTypeVCenterParams)
	objectTypeVCenterParamsModel.IsCloudEnv = core.BoolPtr(true)

	objectTypeWindowsClusterParamsModel := new(backuprecoveryv1.ObjectTypeWindowsClusterParams)
	objectTypeWindowsClusterParamsModel.ClusterSourceType = core.StringPtr("testString")

	objectSummaryModel := new(backuprecoveryv1.ObjectSummary)
	objectSummaryModel.ID = core.Int64Ptr(int64(26))
	objectSummaryModel.Name = core.StringPtr("testString")
	objectSummaryModel.SourceID = core.Int64Ptr(int64(26))
	objectSummaryModel.SourceName = core.StringPtr("testString")
	objectSummaryModel.Environment = core.StringPtr("kPhysical")
	objectSummaryModel.ObjectHash = core.StringPtr("testString")
	objectSummaryModel.ObjectType = core.StringPtr("kCluster")
	objectSummaryModel.LogicalSizeBytes = core.Int64Ptr(int64(26))
	objectSummaryModel.UUID = core.StringPtr("testString")
	objectSummaryModel.GlobalID = core.StringPtr("testString")
	objectSummaryModel.ProtectionType = core.StringPtr("kAgent")
	objectSummaryModel.SharepointSiteSummary = sharepointObjectParamsModel
	objectSummaryModel.OsType = core.StringPtr("kLinux")
	objectSummaryModel.VCenterSummary = objectTypeVCenterParamsModel
	objectSummaryModel.WindowsClusterSummary = objectTypeWindowsClusterParamsModel

	commonRecoverObjectSnapshotParamsObjectInfoModel := new(backuprecoveryv1.CommonRecoverObjectSnapshotParamsObjectInfo)
	commonRecoverObjectSnapshotParamsObjectInfoModel.ID = core.Int64Ptr(int64(26))
	commonRecoverObjectSnapshotParamsObjectInfoModel.Name = core.StringPtr("testString")
	commonRecoverObjectSnapshotParamsObjectInfoModel.SourceID = core.Int64Ptr(int64(26))
	commonRecoverObjectSnapshotParamsObjectInfoModel.SourceName = core.StringPtr("testString")
	commonRecoverObjectSnapshotParamsObjectInfoModel.Environment = core.StringPtr("kPhysical")
	commonRecoverObjectSnapshotParamsObjectInfoModel.ObjectHash = core.StringPtr("testString")
	commonRecoverObjectSnapshotParamsObjectInfoModel.ObjectType = core.StringPtr("kCluster")
	commonRecoverObjectSnapshotParamsObjectInfoModel.LogicalSizeBytes = core.Int64Ptr(int64(26))
	commonRecoverObjectSnapshotParamsObjectInfoModel.UUID = core.StringPtr("testString")
	commonRecoverObjectSnapshotParamsObjectInfoModel.GlobalID = core.StringPtr("testString")
	commonRecoverObjectSnapshotParamsObjectInfoModel.ProtectionType = core.StringPtr("kAgent")
	commonRecoverObjectSnapshotParamsObjectInfoModel.SharepointSiteSummary = sharepointObjectParamsModel
	commonRecoverObjectSnapshotParamsObjectInfoModel.OsType = core.StringPtr("kLinux")
	commonRecoverObjectSnapshotParamsObjectInfoModel.ChildObjects = []backuprecoveryv1.ObjectSummary{*objectSummaryModel}
	commonRecoverObjectSnapshotParamsObjectInfoModel.VCenterSummary = objectTypeVCenterParamsModel
	commonRecoverObjectSnapshotParamsObjectInfoModel.WindowsClusterSummary = objectTypeWindowsClusterParamsModel

	awsTierModel := new(backuprecoveryv1.AWSTier)
	awsTierModel.MoveAfterUnit = core.StringPtr("Days")
	awsTierModel.MoveAfter = core.Int64Ptr(int64(26))
	awsTierModel.TierType = core.StringPtr("kAmazonS3Standard")

	awsTiersModel := new(backuprecoveryv1.AWSTiers)
	awsTiersModel.Tiers = []backuprecoveryv1.AWSTier{*awsTierModel}

	azureTierModel := new(backuprecoveryv1.AzureTier)
	azureTierModel.MoveAfterUnit = core.StringPtr("Days")
	azureTierModel.MoveAfter = core.Int64Ptr(int64(26))
	azureTierModel.TierType = core.StringPtr("kAzureTierHot")

	azureTiersModel := new(backuprecoveryv1.AzureTiers)
	azureTiersModel.Tiers = []backuprecoveryv1.AzureTier{*azureTierModel}

	googleTierModel := new(backuprecoveryv1.GoogleTier)
	googleTierModel.MoveAfterUnit = core.StringPtr("Days")
	googleTierModel.MoveAfter = core.Int64Ptr(int64(26))
	googleTierModel.TierType = core.StringPtr("kGoogleStandard")

	googleTiersModel := new(backuprecoveryv1.GoogleTiers)
	googleTiersModel.Tiers = []backuprecoveryv1.GoogleTier{*googleTierModel}

	oracleTierModel := new(backuprecoveryv1.OracleTier)
	oracleTierModel.MoveAfterUnit = core.StringPtr("Days")
	oracleTierModel.MoveAfter = core.Int64Ptr(int64(26))
	oracleTierModel.TierType = core.StringPtr("kOracleTierStandard")

	oracleTiersModel := new(backuprecoveryv1.OracleTiers)
	oracleTiersModel.Tiers = []backuprecoveryv1.OracleTier{*oracleTierModel}

	archivalTargetTierInfoModel := new(backuprecoveryv1.ArchivalTargetTierInfo)
	archivalTargetTierInfoModel.AwsTiering = awsTiersModel
	archivalTargetTierInfoModel.AzureTiering = azureTiersModel
	archivalTargetTierInfoModel.CloudPlatform = core.StringPtr("AWS")
	archivalTargetTierInfoModel.GoogleTiering = googleTiersModel
	archivalTargetTierInfoModel.OracleTiering = oracleTiersModel
	archivalTargetTierInfoModel.CurrentTierType = core.StringPtr("kAmazonS3Standard")

	commonRecoverObjectSnapshotParamsArchivalTargetInfoModel := new(backuprecoveryv1.CommonRecoverObjectSnapshotParamsArchivalTargetInfo)
	commonRecoverObjectSnapshotParamsArchivalTargetInfoModel.TargetID = core.Int64Ptr(int64(26))
	commonRecoverObjectSnapshotParamsArchivalTargetInfoModel.ArchivalTaskID = core.StringPtr("testString")
	commonRecoverObjectSnapshotParamsArchivalTargetInfoModel.TargetName = core.StringPtr("testString")
	commonRecoverObjectSnapshotParamsArchivalTargetInfoModel.TargetType = core.StringPtr("Tape")
	commonRecoverObjectSnapshotParamsArchivalTargetInfoModel.UsageType = core.StringPtr("Archival")
	commonRecoverObjectSnapshotParamsArchivalTargetInfoModel.OwnershipContext = core.StringPtr("Local")
	commonRecoverObjectSnapshotParamsArchivalTargetInfoModel.TierSettings = archivalTargetTierInfoModel

	commonRecoverObjectSnapshotParamsModel := new(backuprecoveryv1.CommonRecoverObjectSnapshotParams)
	commonRecoverObjectSnapshotParamsModel.SnapshotID = core.StringPtr("testString")
	commonRecoverObjectSnapshotParamsModel.PointInTimeUsecs = core.Int64Ptr(int64(26))
	commonRecoverObjectSnapshotParamsModel.ProtectionGroupID = core.StringPtr("testString")
	commonRecoverObjectSnapshotParamsModel.ProtectionGroupName = core.StringPtr("testString")
	commonRecoverObjectSnapshotParamsModel.ObjectInfo = commonRecoverObjectSnapshotParamsObjectInfoModel
	commonRecoverObjectSnapshotParamsModel.ArchivalTargetInfo = commonRecoverObjectSnapshotParamsArchivalTargetInfoModel
	commonRecoverObjectSnapshotParamsModel.RecoverFromStandby = core.BoolPtr(true)

	physicalTargetParamsForRecoverVolumeMountTargetModel := new(backuprecoveryv1.PhysicalTargetParamsForRecoverVolumeMountTarget)
	physicalTargetParamsForRecoverVolumeMountTargetModel.ID = core.Int64Ptr(int64(26))

	recoverVolumeMappingModel := new(backuprecoveryv1.RecoverVolumeMapping)
	recoverVolumeMappingModel.SourceVolumeGuid = core.StringPtr("testString")
	recoverVolumeMappingModel.DestinationVolumeGuid = core.StringPtr("testString")

	physicalTargetParamsForRecoverVolumeVlanConfigModel := new(backuprecoveryv1.PhysicalTargetParamsForRecoverVolumeVlanConfig)
	physicalTargetParamsForRecoverVolumeVlanConfigModel.ID = core.Int64Ptr(int64(38))
	physicalTargetParamsForRecoverVolumeVlanConfigModel.DisableVlan = core.BoolPtr(true)

	recoverPhysicalVolumeParamsPhysicalTargetParamsModel := new(backuprecoveryv1.RecoverPhysicalVolumeParamsPhysicalTargetParams)
	recoverPhysicalVolumeParamsPhysicalTargetParamsModel.MountTarget = physicalTargetParamsForRecoverVolumeMountTargetModel
	recoverPhysicalVolumeParamsPhysicalTargetParamsModel.VolumeMapping = []backuprecoveryv1.RecoverVolumeMapping{*recoverVolumeMappingModel}
	recoverPhysicalVolumeParamsPhysicalTargetParamsModel.ForceUnmountVolume = core.BoolPtr(true)
	recoverPhysicalVolumeParamsPhysicalTargetParamsModel.VlanConfig = physicalTargetParamsForRecoverVolumeVlanConfigModel

	recoverPhysicalParamsRecoverVolumeParamsModel := new(backuprecoveryv1.RecoverPhysicalParamsRecoverVolumeParams)
	recoverPhysicalParamsRecoverVolumeParamsModel.TargetEnvironment = core.StringPtr("kPhysical")
	recoverPhysicalParamsRecoverVolumeParamsModel.PhysicalTargetParams = recoverPhysicalVolumeParamsPhysicalTargetParamsModel

	physicalMountVolumesOriginalTargetConfigServerCredentialsModel := new(backuprecoveryv1.PhysicalMountVolumesOriginalTargetConfigServerCredentials)
	physicalMountVolumesOriginalTargetConfigServerCredentialsModel.Username = core.StringPtr("testString")
	physicalMountVolumesOriginalTargetConfigServerCredentialsModel.Password = core.StringPtr("testString")

	physicalTargetParamsForMountVolumeOriginalTargetConfigModel := new(backuprecoveryv1.PhysicalTargetParamsForMountVolumeOriginalTargetConfig)
	physicalTargetParamsForMountVolumeOriginalTargetConfigModel.ServerCredentials = physicalMountVolumesOriginalTargetConfigServerCredentialsModel

	recoverTargetModel := new(backuprecoveryv1.RecoverTarget)
	recoverTargetModel.ID = core.Int64Ptr(int64(26))

	physicalMountVolumesNewTargetConfigServerCredentialsModel := new(backuprecoveryv1.PhysicalMountVolumesNewTargetConfigServerCredentials)
	physicalMountVolumesNewTargetConfigServerCredentialsModel.Username = core.StringPtr("testString")
	physicalMountVolumesNewTargetConfigServerCredentialsModel.Password = core.StringPtr("testString")

	physicalTargetParamsForMountVolumeNewTargetConfigModel := new(backuprecoveryv1.PhysicalTargetParamsForMountVolumeNewTargetConfig)
	physicalTargetParamsForMountVolumeNewTargetConfigModel.MountTarget = recoverTargetModel
	physicalTargetParamsForMountVolumeNewTargetConfigModel.ServerCredentials = physicalMountVolumesNewTargetConfigServerCredentialsModel

	physicalTargetParamsForMountVolumeVlanConfigModel := new(backuprecoveryv1.PhysicalTargetParamsForMountVolumeVlanConfig)
	physicalTargetParamsForMountVolumeVlanConfigModel.ID = core.Int64Ptr(int64(38))
	physicalTargetParamsForMountVolumeVlanConfigModel.DisableVlan = core.BoolPtr(true)

	mountPhysicalVolumeParamsPhysicalTargetParamsModel := new(backuprecoveryv1.MountPhysicalVolumeParamsPhysicalTargetParams)
	mountPhysicalVolumeParamsPhysicalTargetParamsModel.MountToOriginalTarget = core.BoolPtr(true)
	mountPhysicalVolumeParamsPhysicalTargetParamsModel.OriginalTargetConfig = physicalTargetParamsForMountVolumeOriginalTargetConfigModel
	mountPhysicalVolumeParamsPhysicalTargetParamsModel.NewTargetConfig = physicalTargetParamsForMountVolumeNewTargetConfigModel
	mountPhysicalVolumeParamsPhysicalTargetParamsModel.ReadOnlyMount = core.BoolPtr(true)
	mountPhysicalVolumeParamsPhysicalTargetParamsModel.VolumeNames = []string{"testString"}
	mountPhysicalVolumeParamsPhysicalTargetParamsModel.VlanConfig = physicalTargetParamsForMountVolumeVlanConfigModel

	recoverPhysicalParamsMountVolumeParamsModel := new(backuprecoveryv1.RecoverPhysicalParamsMountVolumeParams)
	recoverPhysicalParamsMountVolumeParamsModel.TargetEnvironment = core.StringPtr("kPhysical")
	recoverPhysicalParamsMountVolumeParamsModel.PhysicalTargetParams = mountPhysicalVolumeParamsPhysicalTargetParamsModel

	commonRecoverFileAndFolderInfoModel := new(backuprecoveryv1.CommonRecoverFileAndFolderInfo)
	commonRecoverFileAndFolderInfoModel.AbsolutePath = core.StringPtr("testString")
	commonRecoverFileAndFolderInfoModel.IsDirectory = core.BoolPtr(true)
	commonRecoverFileAndFolderInfoModel.IsViewFileRecovery = core.BoolPtr(true)

	physicalTargetParamsForRecoverFileAndFolderRecoverTargetModel := new(backuprecoveryv1.PhysicalTargetParamsForRecoverFileAndFolderRecoverTarget)
	physicalTargetParamsForRecoverFileAndFolderRecoverTargetModel.ID = core.Int64Ptr(int64(26))

	physicalTargetParamsForRecoverFileAndFolderVlanConfigModel := new(backuprecoveryv1.PhysicalTargetParamsForRecoverFileAndFolderVlanConfig)
	physicalTargetParamsForRecoverFileAndFolderVlanConfigModel.ID = core.Int64Ptr(int64(38))
	physicalTargetParamsForRecoverFileAndFolderVlanConfigModel.DisableVlan = core.BoolPtr(true)

	recoverPhysicalFileAndFolderParamsPhysicalTargetParamsModel := new(backuprecoveryv1.RecoverPhysicalFileAndFolderParamsPhysicalTargetParams)
	recoverPhysicalFileAndFolderParamsPhysicalTargetParamsModel.RecoverTarget = physicalTargetParamsForRecoverFileAndFolderRecoverTargetModel
	recoverPhysicalFileAndFolderParamsPhysicalTargetParamsModel.RestoreToOriginalPaths = core.BoolPtr(true)
	recoverPhysicalFileAndFolderParamsPhysicalTargetParamsModel.OverwriteExisting = core.BoolPtr(true)
	recoverPhysicalFileAndFolderParamsPhysicalTargetParamsModel.AlternateRestoreDirectory = core.StringPtr("testString")
	recoverPhysicalFileAndFolderParamsPhysicalTargetParamsModel.PreserveAttributes = core.BoolPtr(true)
	recoverPhysicalFileAndFolderParamsPhysicalTargetParamsModel.PreserveTimestamps = core.BoolPtr(true)
	recoverPhysicalFileAndFolderParamsPhysicalTargetParamsModel.PreserveAcls = core.BoolPtr(true)
	recoverPhysicalFileAndFolderParamsPhysicalTargetParamsModel.ContinueOnError = core.BoolPtr(true)
	recoverPhysicalFileAndFolderParamsPhysicalTargetParamsModel.SaveSuccessFiles = core.BoolPtr(true)
	recoverPhysicalFileAndFolderParamsPhysicalTargetParamsModel.VlanConfig = physicalTargetParamsForRecoverFileAndFolderVlanConfigModel
	recoverPhysicalFileAndFolderParamsPhysicalTargetParamsModel.RestoreEntityType = core.StringPtr("kRegular")

	recoverPhysicalParamsRecoverFileAndFolderParamsModel := new(backuprecoveryv1.RecoverPhysicalParamsRecoverFileAndFolderParams)
	recoverPhysicalParamsRecoverFileAndFolderParamsModel.FilesAndFolders = []backuprecoveryv1.CommonRecoverFileAndFolderInfo{*commonRecoverFileAndFolderInfoModel}
	recoverPhysicalParamsRecoverFileAndFolderParamsModel.TargetEnvironment = core.StringPtr("kPhysical")
	recoverPhysicalParamsRecoverFileAndFolderParamsModel.PhysicalTargetParams = recoverPhysicalFileAndFolderParamsPhysicalTargetParamsModel

	recoverPhysicalParamsDownloadFileAndFolderParamsModel := new(backuprecoveryv1.RecoverPhysicalParamsDownloadFileAndFolderParams)
	recoverPhysicalParamsDownloadFileAndFolderParamsModel.ExpiryTimeUsecs = core.Int64Ptr(int64(26))
	recoverPhysicalParamsDownloadFileAndFolderParamsModel.FilesAndFolders = []backuprecoveryv1.CommonRecoverFileAndFolderInfo{*commonRecoverFileAndFolderInfoModel}
	recoverPhysicalParamsDownloadFileAndFolderParamsModel.DownloadFilePath = core.StringPtr("testString")

	recoverPhysicalParamsSystemRecoveryParamsModel := new(backuprecoveryv1.RecoverPhysicalParamsSystemRecoveryParams)
	recoverPhysicalParamsSystemRecoveryParamsModel.FullNasPath = core.StringPtr("testString")

	model := new(backuprecoveryv1.RecoverPhysicalParams)
	model.Objects = []backuprecoveryv1.CommonRecoverObjectSnapshotParams{*commonRecoverObjectSnapshotParamsModel}
	model.RecoveryAction = core.StringPtr("RecoverPhysicalVolumes")
	model.RecoverVolumeParams = recoverPhysicalParamsRecoverVolumeParamsModel
	model.MountVolumeParams = recoverPhysicalParamsMountVolumeParamsModel
	model.RecoverFileAndFolderParams = recoverPhysicalParamsRecoverFileAndFolderParamsModel
	model.DownloadFileAndFolderParams = recoverPhysicalParamsDownloadFileAndFolderParamsModel
	model.SystemRecoveryParams = recoverPhysicalParamsSystemRecoveryParamsModel

	result, err := backuprecovery.DataSourceIbmRecoveryRecoverPhysicalParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmRecoveryCommonRecoverObjectSnapshotParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		sharepointObjectParamsModel := make(map[string]interface{})
		sharepointObjectParamsModel["site_web_url"] = "testString"

		objectTypeVCenterParamsModel := make(map[string]interface{})
		objectTypeVCenterParamsModel["is_cloud_env"] = true

		objectTypeWindowsClusterParamsModel := make(map[string]interface{})
		objectTypeWindowsClusterParamsModel["cluster_source_type"] = "testString"

		objectSummaryModel := make(map[string]interface{})
		objectSummaryModel["id"] = int(26)
		objectSummaryModel["name"] = "testString"
		objectSummaryModel["source_id"] = int(26)
		objectSummaryModel["source_name"] = "testString"
		objectSummaryModel["environment"] = "kPhysical"
		objectSummaryModel["object_hash"] = "testString"
		objectSummaryModel["object_type"] = "kCluster"
		objectSummaryModel["logical_size_bytes"] = int(26)
		objectSummaryModel["uuid"] = "testString"
		objectSummaryModel["global_id"] = "testString"
		objectSummaryModel["protection_type"] = "kAgent"
		objectSummaryModel["sharepoint_site_summary"] = []map[string]interface{}{sharepointObjectParamsModel}
		objectSummaryModel["os_type"] = "kLinux"
		objectSummaryModel["v_center_summary"] = []map[string]interface{}{objectTypeVCenterParamsModel}
		objectSummaryModel["windows_cluster_summary"] = []map[string]interface{}{objectTypeWindowsClusterParamsModel}

		commonRecoverObjectSnapshotParamsObjectInfoModel := make(map[string]interface{})
		commonRecoverObjectSnapshotParamsObjectInfoModel["id"] = int(26)
		commonRecoverObjectSnapshotParamsObjectInfoModel["name"] = "testString"
		commonRecoverObjectSnapshotParamsObjectInfoModel["source_id"] = int(26)
		commonRecoverObjectSnapshotParamsObjectInfoModel["source_name"] = "testString"
		commonRecoverObjectSnapshotParamsObjectInfoModel["environment"] = "kPhysical"
		commonRecoverObjectSnapshotParamsObjectInfoModel["object_hash"] = "testString"
		commonRecoverObjectSnapshotParamsObjectInfoModel["object_type"] = "kCluster"
		commonRecoverObjectSnapshotParamsObjectInfoModel["logical_size_bytes"] = int(26)
		commonRecoverObjectSnapshotParamsObjectInfoModel["uuid"] = "testString"
		commonRecoverObjectSnapshotParamsObjectInfoModel["global_id"] = "testString"
		commonRecoverObjectSnapshotParamsObjectInfoModel["protection_type"] = "kAgent"
		commonRecoverObjectSnapshotParamsObjectInfoModel["sharepoint_site_summary"] = []map[string]interface{}{sharepointObjectParamsModel}
		commonRecoverObjectSnapshotParamsObjectInfoModel["os_type"] = "kLinux"
		commonRecoverObjectSnapshotParamsObjectInfoModel["child_objects"] = []map[string]interface{}{objectSummaryModel}
		commonRecoverObjectSnapshotParamsObjectInfoModel["v_center_summary"] = []map[string]interface{}{objectTypeVCenterParamsModel}
		commonRecoverObjectSnapshotParamsObjectInfoModel["windows_cluster_summary"] = []map[string]interface{}{objectTypeWindowsClusterParamsModel}

		awsTierModel := make(map[string]interface{})
		awsTierModel["move_after_unit"] = "Days"
		awsTierModel["move_after"] = int(26)
		awsTierModel["tier_type"] = "kAmazonS3Standard"

		awsTiersModel := make(map[string]interface{})
		awsTiersModel["tiers"] = []map[string]interface{}{awsTierModel}

		azureTierModel := make(map[string]interface{})
		azureTierModel["move_after_unit"] = "Days"
		azureTierModel["move_after"] = int(26)
		azureTierModel["tier_type"] = "kAzureTierHot"

		azureTiersModel := make(map[string]interface{})
		azureTiersModel["tiers"] = []map[string]interface{}{azureTierModel}

		googleTierModel := make(map[string]interface{})
		googleTierModel["move_after_unit"] = "Days"
		googleTierModel["move_after"] = int(26)
		googleTierModel["tier_type"] = "kGoogleStandard"

		googleTiersModel := make(map[string]interface{})
		googleTiersModel["tiers"] = []map[string]interface{}{googleTierModel}

		oracleTierModel := make(map[string]interface{})
		oracleTierModel["move_after_unit"] = "Days"
		oracleTierModel["move_after"] = int(26)
		oracleTierModel["tier_type"] = "kOracleTierStandard"

		oracleTiersModel := make(map[string]interface{})
		oracleTiersModel["tiers"] = []map[string]interface{}{oracleTierModel}

		archivalTargetTierInfoModel := make(map[string]interface{})
		archivalTargetTierInfoModel["aws_tiering"] = []map[string]interface{}{awsTiersModel}
		archivalTargetTierInfoModel["azure_tiering"] = []map[string]interface{}{azureTiersModel}
		archivalTargetTierInfoModel["cloud_platform"] = "AWS"
		archivalTargetTierInfoModel["google_tiering"] = []map[string]interface{}{googleTiersModel}
		archivalTargetTierInfoModel["oracle_tiering"] = []map[string]interface{}{oracleTiersModel}
		archivalTargetTierInfoModel["current_tier_type"] = "kAmazonS3Standard"

		commonRecoverObjectSnapshotParamsArchivalTargetInfoModel := make(map[string]interface{})
		commonRecoverObjectSnapshotParamsArchivalTargetInfoModel["target_id"] = int(26)
		commonRecoverObjectSnapshotParamsArchivalTargetInfoModel["archival_task_id"] = "testString"
		commonRecoverObjectSnapshotParamsArchivalTargetInfoModel["target_name"] = "testString"
		commonRecoverObjectSnapshotParamsArchivalTargetInfoModel["target_type"] = "Tape"
		commonRecoverObjectSnapshotParamsArchivalTargetInfoModel["usage_type"] = "Archival"
		commonRecoverObjectSnapshotParamsArchivalTargetInfoModel["ownership_context"] = "Local"
		commonRecoverObjectSnapshotParamsArchivalTargetInfoModel["tier_settings"] = []map[string]interface{}{archivalTargetTierInfoModel}

		model := make(map[string]interface{})
		model["snapshot_id"] = "testString"
		model["point_in_time_usecs"] = int(26)
		model["protection_group_id"] = "testString"
		model["protection_group_name"] = "testString"
		model["snapshot_creation_time_usecs"] = int(26)
		model["object_info"] = []map[string]interface{}{commonRecoverObjectSnapshotParamsObjectInfoModel}
		model["snapshot_target_type"] = "Local"
		model["storage_domain_id"] = int(26)
		model["archival_target_info"] = []map[string]interface{}{commonRecoverObjectSnapshotParamsArchivalTargetInfoModel}
		model["progress_task_id"] = "testString"
		model["recover_from_standby"] = true
		model["status"] = "Accepted"
		model["start_time_usecs"] = int(26)
		model["end_time_usecs"] = int(26)
		model["messages"] = []string{"testString"}
		model["bytes_restored"] = int(26)

		assert.Equal(t, result, model)
	}

	sharepointObjectParamsModel := new(backuprecoveryv1.SharepointObjectParams)
	sharepointObjectParamsModel.SiteWebURL = core.StringPtr("testString")

	objectTypeVCenterParamsModel := new(backuprecoveryv1.ObjectTypeVCenterParams)
	objectTypeVCenterParamsModel.IsCloudEnv = core.BoolPtr(true)

	objectTypeWindowsClusterParamsModel := new(backuprecoveryv1.ObjectTypeWindowsClusterParams)
	objectTypeWindowsClusterParamsModel.ClusterSourceType = core.StringPtr("testString")

	objectSummaryModel := new(backuprecoveryv1.ObjectSummary)
	objectSummaryModel.ID = core.Int64Ptr(int64(26))
	objectSummaryModel.Name = core.StringPtr("testString")
	objectSummaryModel.SourceID = core.Int64Ptr(int64(26))
	objectSummaryModel.SourceName = core.StringPtr("testString")
	objectSummaryModel.Environment = core.StringPtr("kPhysical")
	objectSummaryModel.ObjectHash = core.StringPtr("testString")
	objectSummaryModel.ObjectType = core.StringPtr("kCluster")
	objectSummaryModel.LogicalSizeBytes = core.Int64Ptr(int64(26))
	objectSummaryModel.UUID = core.StringPtr("testString")
	objectSummaryModel.GlobalID = core.StringPtr("testString")
	objectSummaryModel.ProtectionType = core.StringPtr("kAgent")
	objectSummaryModel.SharepointSiteSummary = sharepointObjectParamsModel
	objectSummaryModel.OsType = core.StringPtr("kLinux")
	objectSummaryModel.VCenterSummary = objectTypeVCenterParamsModel
	objectSummaryModel.WindowsClusterSummary = objectTypeWindowsClusterParamsModel

	commonRecoverObjectSnapshotParamsObjectInfoModel := new(backuprecoveryv1.CommonRecoverObjectSnapshotParamsObjectInfo)
	commonRecoverObjectSnapshotParamsObjectInfoModel.ID = core.Int64Ptr(int64(26))
	commonRecoverObjectSnapshotParamsObjectInfoModel.Name = core.StringPtr("testString")
	commonRecoverObjectSnapshotParamsObjectInfoModel.SourceID = core.Int64Ptr(int64(26))
	commonRecoverObjectSnapshotParamsObjectInfoModel.SourceName = core.StringPtr("testString")
	commonRecoverObjectSnapshotParamsObjectInfoModel.Environment = core.StringPtr("kPhysical")
	commonRecoverObjectSnapshotParamsObjectInfoModel.ObjectHash = core.StringPtr("testString")
	commonRecoverObjectSnapshotParamsObjectInfoModel.ObjectType = core.StringPtr("kCluster")
	commonRecoverObjectSnapshotParamsObjectInfoModel.LogicalSizeBytes = core.Int64Ptr(int64(26))
	commonRecoverObjectSnapshotParamsObjectInfoModel.UUID = core.StringPtr("testString")
	commonRecoverObjectSnapshotParamsObjectInfoModel.GlobalID = core.StringPtr("testString")
	commonRecoverObjectSnapshotParamsObjectInfoModel.ProtectionType = core.StringPtr("kAgent")
	commonRecoverObjectSnapshotParamsObjectInfoModel.SharepointSiteSummary = sharepointObjectParamsModel
	commonRecoverObjectSnapshotParamsObjectInfoModel.OsType = core.StringPtr("kLinux")
	commonRecoverObjectSnapshotParamsObjectInfoModel.ChildObjects = []backuprecoveryv1.ObjectSummary{*objectSummaryModel}
	commonRecoverObjectSnapshotParamsObjectInfoModel.VCenterSummary = objectTypeVCenterParamsModel
	commonRecoverObjectSnapshotParamsObjectInfoModel.WindowsClusterSummary = objectTypeWindowsClusterParamsModel

	awsTierModel := new(backuprecoveryv1.AWSTier)
	awsTierModel.MoveAfterUnit = core.StringPtr("Days")
	awsTierModel.MoveAfter = core.Int64Ptr(int64(26))
	awsTierModel.TierType = core.StringPtr("kAmazonS3Standard")

	awsTiersModel := new(backuprecoveryv1.AWSTiers)
	awsTiersModel.Tiers = []backuprecoveryv1.AWSTier{*awsTierModel}

	azureTierModel := new(backuprecoveryv1.AzureTier)
	azureTierModel.MoveAfterUnit = core.StringPtr("Days")
	azureTierModel.MoveAfter = core.Int64Ptr(int64(26))
	azureTierModel.TierType = core.StringPtr("kAzureTierHot")

	azureTiersModel := new(backuprecoveryv1.AzureTiers)
	azureTiersModel.Tiers = []backuprecoveryv1.AzureTier{*azureTierModel}

	googleTierModel := new(backuprecoveryv1.GoogleTier)
	googleTierModel.MoveAfterUnit = core.StringPtr("Days")
	googleTierModel.MoveAfter = core.Int64Ptr(int64(26))
	googleTierModel.TierType = core.StringPtr("kGoogleStandard")

	googleTiersModel := new(backuprecoveryv1.GoogleTiers)
	googleTiersModel.Tiers = []backuprecoveryv1.GoogleTier{*googleTierModel}

	oracleTierModel := new(backuprecoveryv1.OracleTier)
	oracleTierModel.MoveAfterUnit = core.StringPtr("Days")
	oracleTierModel.MoveAfter = core.Int64Ptr(int64(26))
	oracleTierModel.TierType = core.StringPtr("kOracleTierStandard")

	oracleTiersModel := new(backuprecoveryv1.OracleTiers)
	oracleTiersModel.Tiers = []backuprecoveryv1.OracleTier{*oracleTierModel}

	archivalTargetTierInfoModel := new(backuprecoveryv1.ArchivalTargetTierInfo)
	archivalTargetTierInfoModel.AwsTiering = awsTiersModel
	archivalTargetTierInfoModel.AzureTiering = azureTiersModel
	archivalTargetTierInfoModel.CloudPlatform = core.StringPtr("AWS")
	archivalTargetTierInfoModel.GoogleTiering = googleTiersModel
	archivalTargetTierInfoModel.OracleTiering = oracleTiersModel
	archivalTargetTierInfoModel.CurrentTierType = core.StringPtr("kAmazonS3Standard")

	commonRecoverObjectSnapshotParamsArchivalTargetInfoModel := new(backuprecoveryv1.CommonRecoverObjectSnapshotParamsArchivalTargetInfo)
	commonRecoverObjectSnapshotParamsArchivalTargetInfoModel.TargetID = core.Int64Ptr(int64(26))
	commonRecoverObjectSnapshotParamsArchivalTargetInfoModel.ArchivalTaskID = core.StringPtr("testString")
	commonRecoverObjectSnapshotParamsArchivalTargetInfoModel.TargetName = core.StringPtr("testString")
	commonRecoverObjectSnapshotParamsArchivalTargetInfoModel.TargetType = core.StringPtr("Tape")
	commonRecoverObjectSnapshotParamsArchivalTargetInfoModel.UsageType = core.StringPtr("Archival")
	commonRecoverObjectSnapshotParamsArchivalTargetInfoModel.OwnershipContext = core.StringPtr("Local")
	commonRecoverObjectSnapshotParamsArchivalTargetInfoModel.TierSettings = archivalTargetTierInfoModel

	model := new(backuprecoveryv1.CommonRecoverObjectSnapshotParams)
	model.SnapshotID = core.StringPtr("testString")
	model.PointInTimeUsecs = core.Int64Ptr(int64(26))
	model.ProtectionGroupID = core.StringPtr("testString")
	model.ProtectionGroupName = core.StringPtr("testString")
	model.SnapshotCreationTimeUsecs = core.Int64Ptr(int64(26))
	model.ObjectInfo = commonRecoverObjectSnapshotParamsObjectInfoModel
	model.SnapshotTargetType = core.StringPtr("Local")
	model.StorageDomainID = core.Int64Ptr(int64(26))
	model.ArchivalTargetInfo = commonRecoverObjectSnapshotParamsArchivalTargetInfoModel
	model.ProgressTaskID = core.StringPtr("testString")
	model.RecoverFromStandby = core.BoolPtr(true)
	model.Status = core.StringPtr("Accepted")
	model.StartTimeUsecs = core.Int64Ptr(int64(26))
	model.EndTimeUsecs = core.Int64Ptr(int64(26))
	model.Messages = []string{"testString"}
	model.BytesRestored = core.Int64Ptr(int64(26))

	result, err := backuprecovery.DataSourceIbmRecoveryCommonRecoverObjectSnapshotParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmRecoveryCommonRecoverObjectSnapshotParamsObjectInfoToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		sharepointObjectParamsModel := make(map[string]interface{})
		sharepointObjectParamsModel["site_web_url"] = "testString"

		objectTypeVCenterParamsModel := make(map[string]interface{})
		objectTypeVCenterParamsModel["is_cloud_env"] = true

		objectTypeWindowsClusterParamsModel := make(map[string]interface{})
		objectTypeWindowsClusterParamsModel["cluster_source_type"] = "testString"

		objectSummaryModel := make(map[string]interface{})
		objectSummaryModel["id"] = int(26)
		objectSummaryModel["name"] = "testString"
		objectSummaryModel["source_id"] = int(26)
		objectSummaryModel["source_name"] = "testString"
		objectSummaryModel["environment"] = "kPhysical"
		objectSummaryModel["object_hash"] = "testString"
		objectSummaryModel["object_type"] = "kCluster"
		objectSummaryModel["logical_size_bytes"] = int(26)
		objectSummaryModel["uuid"] = "testString"
		objectSummaryModel["global_id"] = "testString"
		objectSummaryModel["protection_type"] = "kAgent"
		objectSummaryModel["sharepoint_site_summary"] = []map[string]interface{}{sharepointObjectParamsModel}
		objectSummaryModel["os_type"] = "kLinux"
		objectSummaryModel["v_center_summary"] = []map[string]interface{}{objectTypeVCenterParamsModel}
		objectSummaryModel["windows_cluster_summary"] = []map[string]interface{}{objectTypeWindowsClusterParamsModel}

		model := make(map[string]interface{})
		model["id"] = int(26)
		model["name"] = "testString"
		model["source_id"] = int(26)
		model["source_name"] = "testString"
		model["environment"] = "kPhysical"
		model["object_hash"] = "testString"
		model["object_type"] = "kCluster"
		model["logical_size_bytes"] = int(26)
		model["uuid"] = "testString"
		model["global_id"] = "testString"
		model["protection_type"] = "kAgent"
		model["sharepoint_site_summary"] = []map[string]interface{}{sharepointObjectParamsModel}
		model["os_type"] = "kLinux"
		model["child_objects"] = []map[string]interface{}{objectSummaryModel}
		model["v_center_summary"] = []map[string]interface{}{objectTypeVCenterParamsModel}
		model["windows_cluster_summary"] = []map[string]interface{}{objectTypeWindowsClusterParamsModel}

		assert.Equal(t, result, model)
	}

	sharepointObjectParamsModel := new(backuprecoveryv1.SharepointObjectParams)
	sharepointObjectParamsModel.SiteWebURL = core.StringPtr("testString")

	objectTypeVCenterParamsModel := new(backuprecoveryv1.ObjectTypeVCenterParams)
	objectTypeVCenterParamsModel.IsCloudEnv = core.BoolPtr(true)

	objectTypeWindowsClusterParamsModel := new(backuprecoveryv1.ObjectTypeWindowsClusterParams)
	objectTypeWindowsClusterParamsModel.ClusterSourceType = core.StringPtr("testString")

	objectSummaryModel := new(backuprecoveryv1.ObjectSummary)
	objectSummaryModel.ID = core.Int64Ptr(int64(26))
	objectSummaryModel.Name = core.StringPtr("testString")
	objectSummaryModel.SourceID = core.Int64Ptr(int64(26))
	objectSummaryModel.SourceName = core.StringPtr("testString")
	objectSummaryModel.Environment = core.StringPtr("kPhysical")
	objectSummaryModel.ObjectHash = core.StringPtr("testString")
	objectSummaryModel.ObjectType = core.StringPtr("kCluster")
	objectSummaryModel.LogicalSizeBytes = core.Int64Ptr(int64(26))
	objectSummaryModel.UUID = core.StringPtr("testString")
	objectSummaryModel.GlobalID = core.StringPtr("testString")
	objectSummaryModel.ProtectionType = core.StringPtr("kAgent")
	objectSummaryModel.SharepointSiteSummary = sharepointObjectParamsModel
	objectSummaryModel.OsType = core.StringPtr("kLinux")
	objectSummaryModel.VCenterSummary = objectTypeVCenterParamsModel
	objectSummaryModel.WindowsClusterSummary = objectTypeWindowsClusterParamsModel

	model := new(backuprecoveryv1.CommonRecoverObjectSnapshotParamsObjectInfo)
	model.ID = core.Int64Ptr(int64(26))
	model.Name = core.StringPtr("testString")
	model.SourceID = core.Int64Ptr(int64(26))
	model.SourceName = core.StringPtr("testString")
	model.Environment = core.StringPtr("kPhysical")
	model.ObjectHash = core.StringPtr("testString")
	model.ObjectType = core.StringPtr("kCluster")
	model.LogicalSizeBytes = core.Int64Ptr(int64(26))
	model.UUID = core.StringPtr("testString")
	model.GlobalID = core.StringPtr("testString")
	model.ProtectionType = core.StringPtr("kAgent")
	model.SharepointSiteSummary = sharepointObjectParamsModel
	model.OsType = core.StringPtr("kLinux")
	model.ChildObjects = []backuprecoveryv1.ObjectSummary{*objectSummaryModel}
	model.VCenterSummary = objectTypeVCenterParamsModel
	model.WindowsClusterSummary = objectTypeWindowsClusterParamsModel

	result, err := backuprecovery.DataSourceIbmRecoveryCommonRecoverObjectSnapshotParamsObjectInfoToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmRecoverySharepointObjectParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["site_web_url"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.SharepointObjectParams)
	model.SiteWebURL = core.StringPtr("testString")

	result, err := backuprecovery.DataSourceIbmRecoverySharepointObjectParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmRecoveryObjectSummaryToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		sharepointObjectParamsModel := make(map[string]interface{})
		sharepointObjectParamsModel["site_web_url"] = "testString"

		objectTypeVCenterParamsModel := make(map[string]interface{})
		objectTypeVCenterParamsModel["is_cloud_env"] = true

		objectTypeWindowsClusterParamsModel := make(map[string]interface{})
		objectTypeWindowsClusterParamsModel["cluster_source_type"] = "testString"

		model := make(map[string]interface{})
		model["id"] = int(26)
		model["name"] = "testString"
		model["source_id"] = int(26)
		model["source_name"] = "testString"
		model["environment"] = "kPhysical"
		model["object_hash"] = "testString"
		model["object_type"] = "kCluster"
		model["logical_size_bytes"] = int(26)
		model["uuid"] = "testString"
		model["global_id"] = "testString"
		model["protection_type"] = "kAgent"
		model["sharepoint_site_summary"] = []map[string]interface{}{sharepointObjectParamsModel}
		model["os_type"] = "kLinux"
		model["v_center_summary"] = []map[string]interface{}{objectTypeVCenterParamsModel}
		model["windows_cluster_summary"] = []map[string]interface{}{objectTypeWindowsClusterParamsModel}

		assert.Equal(t, result, model)
	}

	sharepointObjectParamsModel := new(backuprecoveryv1.SharepointObjectParams)
	sharepointObjectParamsModel.SiteWebURL = core.StringPtr("testString")

	objectTypeVCenterParamsModel := new(backuprecoveryv1.ObjectTypeVCenterParams)
	objectTypeVCenterParamsModel.IsCloudEnv = core.BoolPtr(true)

	objectTypeWindowsClusterParamsModel := new(backuprecoveryv1.ObjectTypeWindowsClusterParams)
	objectTypeWindowsClusterParamsModel.ClusterSourceType = core.StringPtr("testString")

	model := new(backuprecoveryv1.ObjectSummary)
	model.ID = core.Int64Ptr(int64(26))
	model.Name = core.StringPtr("testString")
	model.SourceID = core.Int64Ptr(int64(26))
	model.SourceName = core.StringPtr("testString")
	model.Environment = core.StringPtr("kPhysical")
	model.ObjectHash = core.StringPtr("testString")
	model.ObjectType = core.StringPtr("kCluster")
	model.LogicalSizeBytes = core.Int64Ptr(int64(26))
	model.UUID = core.StringPtr("testString")
	model.GlobalID = core.StringPtr("testString")
	model.ProtectionType = core.StringPtr("kAgent")
	model.SharepointSiteSummary = sharepointObjectParamsModel
	model.OsType = core.StringPtr("kLinux")
	model.VCenterSummary = objectTypeVCenterParamsModel
	model.WindowsClusterSummary = objectTypeWindowsClusterParamsModel

	result, err := backuprecovery.DataSourceIbmRecoveryObjectSummaryToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmRecoveryObjectTypeVCenterParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["is_cloud_env"] = true

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.ObjectTypeVCenterParams)
	model.IsCloudEnv = core.BoolPtr(true)

	result, err := backuprecovery.DataSourceIbmRecoveryObjectTypeVCenterParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmRecoveryObjectTypeWindowsClusterParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["cluster_source_type"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.ObjectTypeWindowsClusterParams)
	model.ClusterSourceType = core.StringPtr("testString")

	result, err := backuprecovery.DataSourceIbmRecoveryObjectTypeWindowsClusterParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmRecoveryCommonRecoverObjectSnapshotParamsArchivalTargetInfoToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		awsTierModel := make(map[string]interface{})
		awsTierModel["move_after_unit"] = "Days"
		awsTierModel["move_after"] = int(26)
		awsTierModel["tier_type"] = "kAmazonS3Standard"

		awsTiersModel := make(map[string]interface{})
		awsTiersModel["tiers"] = []map[string]interface{}{awsTierModel}

		azureTierModel := make(map[string]interface{})
		azureTierModel["move_after_unit"] = "Days"
		azureTierModel["move_after"] = int(26)
		azureTierModel["tier_type"] = "kAzureTierHot"

		azureTiersModel := make(map[string]interface{})
		azureTiersModel["tiers"] = []map[string]interface{}{azureTierModel}

		googleTierModel := make(map[string]interface{})
		googleTierModel["move_after_unit"] = "Days"
		googleTierModel["move_after"] = int(26)
		googleTierModel["tier_type"] = "kGoogleStandard"

		googleTiersModel := make(map[string]interface{})
		googleTiersModel["tiers"] = []map[string]interface{}{googleTierModel}

		oracleTierModel := make(map[string]interface{})
		oracleTierModel["move_after_unit"] = "Days"
		oracleTierModel["move_after"] = int(26)
		oracleTierModel["tier_type"] = "kOracleTierStandard"

		oracleTiersModel := make(map[string]interface{})
		oracleTiersModel["tiers"] = []map[string]interface{}{oracleTierModel}

		archivalTargetTierInfoModel := make(map[string]interface{})
		archivalTargetTierInfoModel["aws_tiering"] = []map[string]interface{}{awsTiersModel}
		archivalTargetTierInfoModel["azure_tiering"] = []map[string]interface{}{azureTiersModel}
		archivalTargetTierInfoModel["cloud_platform"] = "AWS"
		archivalTargetTierInfoModel["google_tiering"] = []map[string]interface{}{googleTiersModel}
		archivalTargetTierInfoModel["oracle_tiering"] = []map[string]interface{}{oracleTiersModel}
		archivalTargetTierInfoModel["current_tier_type"] = "kAmazonS3Standard"

		model := make(map[string]interface{})
		model["target_id"] = int(26)
		model["archival_task_id"] = "testString"
		model["target_name"] = "testString"
		model["target_type"] = "Tape"
		model["usage_type"] = "Archival"
		model["ownership_context"] = "Local"
		model["tier_settings"] = []map[string]interface{}{archivalTargetTierInfoModel}

		assert.Equal(t, result, model)
	}

	awsTierModel := new(backuprecoveryv1.AWSTier)
	awsTierModel.MoveAfterUnit = core.StringPtr("Days")
	awsTierModel.MoveAfter = core.Int64Ptr(int64(26))
	awsTierModel.TierType = core.StringPtr("kAmazonS3Standard")

	awsTiersModel := new(backuprecoveryv1.AWSTiers)
	awsTiersModel.Tiers = []backuprecoveryv1.AWSTier{*awsTierModel}

	azureTierModel := new(backuprecoveryv1.AzureTier)
	azureTierModel.MoveAfterUnit = core.StringPtr("Days")
	azureTierModel.MoveAfter = core.Int64Ptr(int64(26))
	azureTierModel.TierType = core.StringPtr("kAzureTierHot")

	azureTiersModel := new(backuprecoveryv1.AzureTiers)
	azureTiersModel.Tiers = []backuprecoveryv1.AzureTier{*azureTierModel}

	googleTierModel := new(backuprecoveryv1.GoogleTier)
	googleTierModel.MoveAfterUnit = core.StringPtr("Days")
	googleTierModel.MoveAfter = core.Int64Ptr(int64(26))
	googleTierModel.TierType = core.StringPtr("kGoogleStandard")

	googleTiersModel := new(backuprecoveryv1.GoogleTiers)
	googleTiersModel.Tiers = []backuprecoveryv1.GoogleTier{*googleTierModel}

	oracleTierModel := new(backuprecoveryv1.OracleTier)
	oracleTierModel.MoveAfterUnit = core.StringPtr("Days")
	oracleTierModel.MoveAfter = core.Int64Ptr(int64(26))
	oracleTierModel.TierType = core.StringPtr("kOracleTierStandard")

	oracleTiersModel := new(backuprecoveryv1.OracleTiers)
	oracleTiersModel.Tiers = []backuprecoveryv1.OracleTier{*oracleTierModel}

	archivalTargetTierInfoModel := new(backuprecoveryv1.ArchivalTargetTierInfo)
	archivalTargetTierInfoModel.AwsTiering = awsTiersModel
	archivalTargetTierInfoModel.AzureTiering = azureTiersModel
	archivalTargetTierInfoModel.CloudPlatform = core.StringPtr("AWS")
	archivalTargetTierInfoModel.GoogleTiering = googleTiersModel
	archivalTargetTierInfoModel.OracleTiering = oracleTiersModel
	archivalTargetTierInfoModel.CurrentTierType = core.StringPtr("kAmazonS3Standard")

	model := new(backuprecoveryv1.CommonRecoverObjectSnapshotParamsArchivalTargetInfo)
	model.TargetID = core.Int64Ptr(int64(26))
	model.ArchivalTaskID = core.StringPtr("testString")
	model.TargetName = core.StringPtr("testString")
	model.TargetType = core.StringPtr("Tape")
	model.UsageType = core.StringPtr("Archival")
	model.OwnershipContext = core.StringPtr("Local")
	model.TierSettings = archivalTargetTierInfoModel

	result, err := backuprecovery.DataSourceIbmRecoveryCommonRecoverObjectSnapshotParamsArchivalTargetInfoToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmRecoveryArchivalTargetTierInfoToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		awsTierModel := make(map[string]interface{})
		awsTierModel["move_after_unit"] = "Days"
		awsTierModel["move_after"] = int(26)
		awsTierModel["tier_type"] = "kAmazonS3Standard"

		awsTiersModel := make(map[string]interface{})
		awsTiersModel["tiers"] = []map[string]interface{}{awsTierModel}

		azureTierModel := make(map[string]interface{})
		azureTierModel["move_after_unit"] = "Days"
		azureTierModel["move_after"] = int(26)
		azureTierModel["tier_type"] = "kAzureTierHot"

		azureTiersModel := make(map[string]interface{})
		azureTiersModel["tiers"] = []map[string]interface{}{azureTierModel}

		googleTierModel := make(map[string]interface{})
		googleTierModel["move_after_unit"] = "Days"
		googleTierModel["move_after"] = int(26)
		googleTierModel["tier_type"] = "kGoogleStandard"

		googleTiersModel := make(map[string]interface{})
		googleTiersModel["tiers"] = []map[string]interface{}{googleTierModel}

		oracleTierModel := make(map[string]interface{})
		oracleTierModel["move_after_unit"] = "Days"
		oracleTierModel["move_after"] = int(26)
		oracleTierModel["tier_type"] = "kOracleTierStandard"

		oracleTiersModel := make(map[string]interface{})
		oracleTiersModel["tiers"] = []map[string]interface{}{oracleTierModel}

		model := make(map[string]interface{})
		model["aws_tiering"] = []map[string]interface{}{awsTiersModel}
		model["azure_tiering"] = []map[string]interface{}{azureTiersModel}
		model["cloud_platform"] = "AWS"
		model["google_tiering"] = []map[string]interface{}{googleTiersModel}
		model["oracle_tiering"] = []map[string]interface{}{oracleTiersModel}
		model["current_tier_type"] = "kAmazonS3Standard"

		assert.Equal(t, result, model)
	}

	awsTierModel := new(backuprecoveryv1.AWSTier)
	awsTierModel.MoveAfterUnit = core.StringPtr("Days")
	awsTierModel.MoveAfter = core.Int64Ptr(int64(26))
	awsTierModel.TierType = core.StringPtr("kAmazonS3Standard")

	awsTiersModel := new(backuprecoveryv1.AWSTiers)
	awsTiersModel.Tiers = []backuprecoveryv1.AWSTier{*awsTierModel}

	azureTierModel := new(backuprecoveryv1.AzureTier)
	azureTierModel.MoveAfterUnit = core.StringPtr("Days")
	azureTierModel.MoveAfter = core.Int64Ptr(int64(26))
	azureTierModel.TierType = core.StringPtr("kAzureTierHot")

	azureTiersModel := new(backuprecoveryv1.AzureTiers)
	azureTiersModel.Tiers = []backuprecoveryv1.AzureTier{*azureTierModel}

	googleTierModel := new(backuprecoveryv1.GoogleTier)
	googleTierModel.MoveAfterUnit = core.StringPtr("Days")
	googleTierModel.MoveAfter = core.Int64Ptr(int64(26))
	googleTierModel.TierType = core.StringPtr("kGoogleStandard")

	googleTiersModel := new(backuprecoveryv1.GoogleTiers)
	googleTiersModel.Tiers = []backuprecoveryv1.GoogleTier{*googleTierModel}

	oracleTierModel := new(backuprecoveryv1.OracleTier)
	oracleTierModel.MoveAfterUnit = core.StringPtr("Days")
	oracleTierModel.MoveAfter = core.Int64Ptr(int64(26))
	oracleTierModel.TierType = core.StringPtr("kOracleTierStandard")

	oracleTiersModel := new(backuprecoveryv1.OracleTiers)
	oracleTiersModel.Tiers = []backuprecoveryv1.OracleTier{*oracleTierModel}

	model := new(backuprecoveryv1.ArchivalTargetTierInfo)
	model.AwsTiering = awsTiersModel
	model.AzureTiering = azureTiersModel
	model.CloudPlatform = core.StringPtr("AWS")
	model.GoogleTiering = googleTiersModel
	model.OracleTiering = oracleTiersModel
	model.CurrentTierType = core.StringPtr("kAmazonS3Standard")

	result, err := backuprecovery.DataSourceIbmRecoveryArchivalTargetTierInfoToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmRecoveryAWSTiersToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		awsTierModel := make(map[string]interface{})
		awsTierModel["move_after_unit"] = "Days"
		awsTierModel["move_after"] = int(26)
		awsTierModel["tier_type"] = "kAmazonS3Standard"

		model := make(map[string]interface{})
		model["tiers"] = []map[string]interface{}{awsTierModel}

		assert.Equal(t, result, model)
	}

	awsTierModel := new(backuprecoveryv1.AWSTier)
	awsTierModel.MoveAfterUnit = core.StringPtr("Days")
	awsTierModel.MoveAfter = core.Int64Ptr(int64(26))
	awsTierModel.TierType = core.StringPtr("kAmazonS3Standard")

	model := new(backuprecoveryv1.AWSTiers)
	model.Tiers = []backuprecoveryv1.AWSTier{*awsTierModel}

	result, err := backuprecovery.DataSourceIbmRecoveryAWSTiersToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmRecoveryAWSTierToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["move_after_unit"] = "Days"
		model["move_after"] = int(26)
		model["tier_type"] = "kAmazonS3Standard"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.AWSTier)
	model.MoveAfterUnit = core.StringPtr("Days")
	model.MoveAfter = core.Int64Ptr(int64(26))
	model.TierType = core.StringPtr("kAmazonS3Standard")

	result, err := backuprecovery.DataSourceIbmRecoveryAWSTierToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmRecoveryAzureTiersToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		azureTierModel := make(map[string]interface{})
		azureTierModel["move_after_unit"] = "Days"
		azureTierModel["move_after"] = int(26)
		azureTierModel["tier_type"] = "kAzureTierHot"

		model := make(map[string]interface{})
		model["tiers"] = []map[string]interface{}{azureTierModel}

		assert.Equal(t, result, model)
	}

	azureTierModel := new(backuprecoveryv1.AzureTier)
	azureTierModel.MoveAfterUnit = core.StringPtr("Days")
	azureTierModel.MoveAfter = core.Int64Ptr(int64(26))
	azureTierModel.TierType = core.StringPtr("kAzureTierHot")

	model := new(backuprecoveryv1.AzureTiers)
	model.Tiers = []backuprecoveryv1.AzureTier{*azureTierModel}

	result, err := backuprecovery.DataSourceIbmRecoveryAzureTiersToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmRecoveryAzureTierToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["move_after_unit"] = "Days"
		model["move_after"] = int(26)
		model["tier_type"] = "kAzureTierHot"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.AzureTier)
	model.MoveAfterUnit = core.StringPtr("Days")
	model.MoveAfter = core.Int64Ptr(int64(26))
	model.TierType = core.StringPtr("kAzureTierHot")

	result, err := backuprecovery.DataSourceIbmRecoveryAzureTierToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmRecoveryGoogleTiersToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		googleTierModel := make(map[string]interface{})
		googleTierModel["move_after_unit"] = "Days"
		googleTierModel["move_after"] = int(26)
		googleTierModel["tier_type"] = "kGoogleStandard"

		model := make(map[string]interface{})
		model["tiers"] = []map[string]interface{}{googleTierModel}

		assert.Equal(t, result, model)
	}

	googleTierModel := new(backuprecoveryv1.GoogleTier)
	googleTierModel.MoveAfterUnit = core.StringPtr("Days")
	googleTierModel.MoveAfter = core.Int64Ptr(int64(26))
	googleTierModel.TierType = core.StringPtr("kGoogleStandard")

	model := new(backuprecoveryv1.GoogleTiers)
	model.Tiers = []backuprecoveryv1.GoogleTier{*googleTierModel}

	result, err := backuprecovery.DataSourceIbmRecoveryGoogleTiersToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmRecoveryGoogleTierToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["move_after_unit"] = "Days"
		model["move_after"] = int(26)
		model["tier_type"] = "kGoogleStandard"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.GoogleTier)
	model.MoveAfterUnit = core.StringPtr("Days")
	model.MoveAfter = core.Int64Ptr(int64(26))
	model.TierType = core.StringPtr("kGoogleStandard")

	result, err := backuprecovery.DataSourceIbmRecoveryGoogleTierToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmRecoveryOracleTiersToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		oracleTierModel := make(map[string]interface{})
		oracleTierModel["move_after_unit"] = "Days"
		oracleTierModel["move_after"] = int(26)
		oracleTierModel["tier_type"] = "kOracleTierStandard"

		model := make(map[string]interface{})
		model["tiers"] = []map[string]interface{}{oracleTierModel}

		assert.Equal(t, result, model)
	}

	oracleTierModel := new(backuprecoveryv1.OracleTier)
	oracleTierModel.MoveAfterUnit = core.StringPtr("Days")
	oracleTierModel.MoveAfter = core.Int64Ptr(int64(26))
	oracleTierModel.TierType = core.StringPtr("kOracleTierStandard")

	model := new(backuprecoveryv1.OracleTiers)
	model.Tiers = []backuprecoveryv1.OracleTier{*oracleTierModel}

	result, err := backuprecovery.DataSourceIbmRecoveryOracleTiersToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmRecoveryOracleTierToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["move_after_unit"] = "Days"
		model["move_after"] = int(26)
		model["tier_type"] = "kOracleTierStandard"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.OracleTier)
	model.MoveAfterUnit = core.StringPtr("Days")
	model.MoveAfter = core.Int64Ptr(int64(26))
	model.TierType = core.StringPtr("kOracleTierStandard")

	result, err := backuprecovery.DataSourceIbmRecoveryOracleTierToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmRecoveryRecoverPhysicalParamsRecoverVolumeParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		physicalTargetParamsForRecoverVolumeMountTargetModel := make(map[string]interface{})
		physicalTargetParamsForRecoverVolumeMountTargetModel["id"] = int(26)

		recoverVolumeMappingModel := make(map[string]interface{})
		recoverVolumeMappingModel["source_volume_guid"] = "testString"
		recoverVolumeMappingModel["destination_volume_guid"] = "testString"

		physicalTargetParamsForRecoverVolumeVlanConfigModel := make(map[string]interface{})
		physicalTargetParamsForRecoverVolumeVlanConfigModel["id"] = int(38)
		physicalTargetParamsForRecoverVolumeVlanConfigModel["disable_vlan"] = true

		recoverPhysicalVolumeParamsPhysicalTargetParamsModel := make(map[string]interface{})
		recoverPhysicalVolumeParamsPhysicalTargetParamsModel["mount_target"] = []map[string]interface{}{physicalTargetParamsForRecoverVolumeMountTargetModel}
		recoverPhysicalVolumeParamsPhysicalTargetParamsModel["volume_mapping"] = []map[string]interface{}{recoverVolumeMappingModel}
		recoverPhysicalVolumeParamsPhysicalTargetParamsModel["force_unmount_volume"] = true
		recoverPhysicalVolumeParamsPhysicalTargetParamsModel["vlan_config"] = []map[string]interface{}{physicalTargetParamsForRecoverVolumeVlanConfigModel}

		model := make(map[string]interface{})
		model["target_environment"] = "kPhysical"
		model["physical_target_params"] = []map[string]interface{}{recoverPhysicalVolumeParamsPhysicalTargetParamsModel}

		assert.Equal(t, result, model)
	}

	physicalTargetParamsForRecoverVolumeMountTargetModel := new(backuprecoveryv1.PhysicalTargetParamsForRecoverVolumeMountTarget)
	physicalTargetParamsForRecoverVolumeMountTargetModel.ID = core.Int64Ptr(int64(26))

	recoverVolumeMappingModel := new(backuprecoveryv1.RecoverVolumeMapping)
	recoverVolumeMappingModel.SourceVolumeGuid = core.StringPtr("testString")
	recoverVolumeMappingModel.DestinationVolumeGuid = core.StringPtr("testString")

	physicalTargetParamsForRecoverVolumeVlanConfigModel := new(backuprecoveryv1.PhysicalTargetParamsForRecoverVolumeVlanConfig)
	physicalTargetParamsForRecoverVolumeVlanConfigModel.ID = core.Int64Ptr(int64(38))
	physicalTargetParamsForRecoverVolumeVlanConfigModel.DisableVlan = core.BoolPtr(true)

	recoverPhysicalVolumeParamsPhysicalTargetParamsModel := new(backuprecoveryv1.RecoverPhysicalVolumeParamsPhysicalTargetParams)
	recoverPhysicalVolumeParamsPhysicalTargetParamsModel.MountTarget = physicalTargetParamsForRecoverVolumeMountTargetModel
	recoverPhysicalVolumeParamsPhysicalTargetParamsModel.VolumeMapping = []backuprecoveryv1.RecoverVolumeMapping{*recoverVolumeMappingModel}
	recoverPhysicalVolumeParamsPhysicalTargetParamsModel.ForceUnmountVolume = core.BoolPtr(true)
	recoverPhysicalVolumeParamsPhysicalTargetParamsModel.VlanConfig = physicalTargetParamsForRecoverVolumeVlanConfigModel

	model := new(backuprecoveryv1.RecoverPhysicalParamsRecoverVolumeParams)
	model.TargetEnvironment = core.StringPtr("kPhysical")
	model.PhysicalTargetParams = recoverPhysicalVolumeParamsPhysicalTargetParamsModel

	result, err := backuprecovery.DataSourceIbmRecoveryRecoverPhysicalParamsRecoverVolumeParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmRecoveryRecoverPhysicalVolumeParamsPhysicalTargetParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		physicalTargetParamsForRecoverVolumeMountTargetModel := make(map[string]interface{})
		physicalTargetParamsForRecoverVolumeMountTargetModel["id"] = int(26)

		recoverVolumeMappingModel := make(map[string]interface{})
		recoverVolumeMappingModel["source_volume_guid"] = "testString"
		recoverVolumeMappingModel["destination_volume_guid"] = "testString"

		physicalTargetParamsForRecoverVolumeVlanConfigModel := make(map[string]interface{})
		physicalTargetParamsForRecoverVolumeVlanConfigModel["id"] = int(38)
		physicalTargetParamsForRecoverVolumeVlanConfigModel["disable_vlan"] = true

		model := make(map[string]interface{})
		model["mount_target"] = []map[string]interface{}{physicalTargetParamsForRecoverVolumeMountTargetModel}
		model["volume_mapping"] = []map[string]interface{}{recoverVolumeMappingModel}
		model["force_unmount_volume"] = true
		model["vlan_config"] = []map[string]interface{}{physicalTargetParamsForRecoverVolumeVlanConfigModel}

		assert.Equal(t, result, model)
	}

	physicalTargetParamsForRecoverVolumeMountTargetModel := new(backuprecoveryv1.PhysicalTargetParamsForRecoverVolumeMountTarget)
	physicalTargetParamsForRecoverVolumeMountTargetModel.ID = core.Int64Ptr(int64(26))

	recoverVolumeMappingModel := new(backuprecoveryv1.RecoverVolumeMapping)
	recoverVolumeMappingModel.SourceVolumeGuid = core.StringPtr("testString")
	recoverVolumeMappingModel.DestinationVolumeGuid = core.StringPtr("testString")

	physicalTargetParamsForRecoverVolumeVlanConfigModel := new(backuprecoveryv1.PhysicalTargetParamsForRecoverVolumeVlanConfig)
	physicalTargetParamsForRecoverVolumeVlanConfigModel.ID = core.Int64Ptr(int64(38))
	physicalTargetParamsForRecoverVolumeVlanConfigModel.DisableVlan = core.BoolPtr(true)

	model := new(backuprecoveryv1.RecoverPhysicalVolumeParamsPhysicalTargetParams)
	model.MountTarget = physicalTargetParamsForRecoverVolumeMountTargetModel
	model.VolumeMapping = []backuprecoveryv1.RecoverVolumeMapping{*recoverVolumeMappingModel}
	model.ForceUnmountVolume = core.BoolPtr(true)
	model.VlanConfig = physicalTargetParamsForRecoverVolumeVlanConfigModel

	result, err := backuprecovery.DataSourceIbmRecoveryRecoverPhysicalVolumeParamsPhysicalTargetParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmRecoveryPhysicalTargetParamsForRecoverVolumeMountTargetToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["id"] = int(26)
		model["name"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.PhysicalTargetParamsForRecoverVolumeMountTarget)
	model.ID = core.Int64Ptr(int64(26))
	model.Name = core.StringPtr("testString")

	result, err := backuprecovery.DataSourceIbmRecoveryPhysicalTargetParamsForRecoverVolumeMountTargetToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmRecoveryRecoverVolumeMappingToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["source_volume_guid"] = "testString"
		model["destination_volume_guid"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.RecoverVolumeMapping)
	model.SourceVolumeGuid = core.StringPtr("testString")
	model.DestinationVolumeGuid = core.StringPtr("testString")

	result, err := backuprecovery.DataSourceIbmRecoveryRecoverVolumeMappingToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmRecoveryPhysicalTargetParamsForRecoverVolumeVlanConfigToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["id"] = int(38)
		model["disable_vlan"] = true
		model["interface_name"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.PhysicalTargetParamsForRecoverVolumeVlanConfig)
	model.ID = core.Int64Ptr(int64(38))
	model.DisableVlan = core.BoolPtr(true)
	model.InterfaceName = core.StringPtr("testString")

	result, err := backuprecovery.DataSourceIbmRecoveryPhysicalTargetParamsForRecoverVolumeVlanConfigToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmRecoveryRecoverPhysicalParamsMountVolumeParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		physicalMountVolumesOriginalTargetConfigServerCredentialsModel := make(map[string]interface{})
		physicalMountVolumesOriginalTargetConfigServerCredentialsModel["username"] = "testString"
		physicalMountVolumesOriginalTargetConfigServerCredentialsModel["password"] = "testString"

		physicalTargetParamsForMountVolumeOriginalTargetConfigModel := make(map[string]interface{})
		physicalTargetParamsForMountVolumeOriginalTargetConfigModel["server_credentials"] = []map[string]interface{}{physicalMountVolumesOriginalTargetConfigServerCredentialsModel}

		recoverTargetModel := make(map[string]interface{})
		recoverTargetModel["id"] = int(26)

		physicalMountVolumesNewTargetConfigServerCredentialsModel := make(map[string]interface{})
		physicalMountVolumesNewTargetConfigServerCredentialsModel["username"] = "testString"
		physicalMountVolumesNewTargetConfigServerCredentialsModel["password"] = "testString"

		physicalTargetParamsForMountVolumeNewTargetConfigModel := make(map[string]interface{})
		physicalTargetParamsForMountVolumeNewTargetConfigModel["mount_target"] = []map[string]interface{}{recoverTargetModel}
		physicalTargetParamsForMountVolumeNewTargetConfigModel["server_credentials"] = []map[string]interface{}{physicalMountVolumesNewTargetConfigServerCredentialsModel}

		physicalTargetParamsForMountVolumeVlanConfigModel := make(map[string]interface{})
		physicalTargetParamsForMountVolumeVlanConfigModel["id"] = int(38)
		physicalTargetParamsForMountVolumeVlanConfigModel["disable_vlan"] = true

		mountPhysicalVolumeParamsPhysicalTargetParamsModel := make(map[string]interface{})
		mountPhysicalVolumeParamsPhysicalTargetParamsModel["mount_to_original_target"] = true
		mountPhysicalVolumeParamsPhysicalTargetParamsModel["original_target_config"] = []map[string]interface{}{physicalTargetParamsForMountVolumeOriginalTargetConfigModel}
		mountPhysicalVolumeParamsPhysicalTargetParamsModel["new_target_config"] = []map[string]interface{}{physicalTargetParamsForMountVolumeNewTargetConfigModel}
		mountPhysicalVolumeParamsPhysicalTargetParamsModel["read_only_mount"] = true
		mountPhysicalVolumeParamsPhysicalTargetParamsModel["volume_names"] = []string{"testString"}
		mountPhysicalVolumeParamsPhysicalTargetParamsModel["vlan_config"] = []map[string]interface{}{physicalTargetParamsForMountVolumeVlanConfigModel}

		model := make(map[string]interface{})
		model["target_environment"] = "kPhysical"
		model["physical_target_params"] = []map[string]interface{}{mountPhysicalVolumeParamsPhysicalTargetParamsModel}

		assert.Equal(t, result, model)
	}

	physicalMountVolumesOriginalTargetConfigServerCredentialsModel := new(backuprecoveryv1.PhysicalMountVolumesOriginalTargetConfigServerCredentials)
	physicalMountVolumesOriginalTargetConfigServerCredentialsModel.Username = core.StringPtr("testString")
	physicalMountVolumesOriginalTargetConfigServerCredentialsModel.Password = core.StringPtr("testString")

	physicalTargetParamsForMountVolumeOriginalTargetConfigModel := new(backuprecoveryv1.PhysicalTargetParamsForMountVolumeOriginalTargetConfig)
	physicalTargetParamsForMountVolumeOriginalTargetConfigModel.ServerCredentials = physicalMountVolumesOriginalTargetConfigServerCredentialsModel

	recoverTargetModel := new(backuprecoveryv1.RecoverTarget)
	recoverTargetModel.ID = core.Int64Ptr(int64(26))

	physicalMountVolumesNewTargetConfigServerCredentialsModel := new(backuprecoveryv1.PhysicalMountVolumesNewTargetConfigServerCredentials)
	physicalMountVolumesNewTargetConfigServerCredentialsModel.Username = core.StringPtr("testString")
	physicalMountVolumesNewTargetConfigServerCredentialsModel.Password = core.StringPtr("testString")

	physicalTargetParamsForMountVolumeNewTargetConfigModel := new(backuprecoveryv1.PhysicalTargetParamsForMountVolumeNewTargetConfig)
	physicalTargetParamsForMountVolumeNewTargetConfigModel.MountTarget = recoverTargetModel
	physicalTargetParamsForMountVolumeNewTargetConfigModel.ServerCredentials = physicalMountVolumesNewTargetConfigServerCredentialsModel

	physicalTargetParamsForMountVolumeVlanConfigModel := new(backuprecoveryv1.PhysicalTargetParamsForMountVolumeVlanConfig)
	physicalTargetParamsForMountVolumeVlanConfigModel.ID = core.Int64Ptr(int64(38))
	physicalTargetParamsForMountVolumeVlanConfigModel.DisableVlan = core.BoolPtr(true)

	mountPhysicalVolumeParamsPhysicalTargetParamsModel := new(backuprecoveryv1.MountPhysicalVolumeParamsPhysicalTargetParams)
	mountPhysicalVolumeParamsPhysicalTargetParamsModel.MountToOriginalTarget = core.BoolPtr(true)
	mountPhysicalVolumeParamsPhysicalTargetParamsModel.OriginalTargetConfig = physicalTargetParamsForMountVolumeOriginalTargetConfigModel
	mountPhysicalVolumeParamsPhysicalTargetParamsModel.NewTargetConfig = physicalTargetParamsForMountVolumeNewTargetConfigModel
	mountPhysicalVolumeParamsPhysicalTargetParamsModel.ReadOnlyMount = core.BoolPtr(true)
	mountPhysicalVolumeParamsPhysicalTargetParamsModel.VolumeNames = []string{"testString"}
	mountPhysicalVolumeParamsPhysicalTargetParamsModel.VlanConfig = physicalTargetParamsForMountVolumeVlanConfigModel

	model := new(backuprecoveryv1.RecoverPhysicalParamsMountVolumeParams)
	model.TargetEnvironment = core.StringPtr("kPhysical")
	model.PhysicalTargetParams = mountPhysicalVolumeParamsPhysicalTargetParamsModel

	result, err := backuprecovery.DataSourceIbmRecoveryRecoverPhysicalParamsMountVolumeParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmRecoveryMountPhysicalVolumeParamsPhysicalTargetParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		physicalMountVolumesOriginalTargetConfigServerCredentialsModel := make(map[string]interface{})
		physicalMountVolumesOriginalTargetConfigServerCredentialsModel["username"] = "testString"
		physicalMountVolumesOriginalTargetConfigServerCredentialsModel["password"] = "testString"

		physicalTargetParamsForMountVolumeOriginalTargetConfigModel := make(map[string]interface{})
		physicalTargetParamsForMountVolumeOriginalTargetConfigModel["server_credentials"] = []map[string]interface{}{physicalMountVolumesOriginalTargetConfigServerCredentialsModel}

		recoverTargetModel := make(map[string]interface{})
		recoverTargetModel["id"] = int(26)

		physicalMountVolumesNewTargetConfigServerCredentialsModel := make(map[string]interface{})
		physicalMountVolumesNewTargetConfigServerCredentialsModel["username"] = "testString"
		physicalMountVolumesNewTargetConfigServerCredentialsModel["password"] = "testString"

		physicalTargetParamsForMountVolumeNewTargetConfigModel := make(map[string]interface{})
		physicalTargetParamsForMountVolumeNewTargetConfigModel["mount_target"] = []map[string]interface{}{recoverTargetModel}
		physicalTargetParamsForMountVolumeNewTargetConfigModel["server_credentials"] = []map[string]interface{}{physicalMountVolumesNewTargetConfigServerCredentialsModel}

		physicalTargetParamsForMountVolumeVlanConfigModel := make(map[string]interface{})
		physicalTargetParamsForMountVolumeVlanConfigModel["id"] = int(38)
		physicalTargetParamsForMountVolumeVlanConfigModel["disable_vlan"] = true

		model := make(map[string]interface{})
		model["mount_to_original_target"] = true
		model["original_target_config"] = []map[string]interface{}{physicalTargetParamsForMountVolumeOriginalTargetConfigModel}
		model["new_target_config"] = []map[string]interface{}{physicalTargetParamsForMountVolumeNewTargetConfigModel}
		model["read_only_mount"] = true
		model["volume_names"] = []string{"testString"}
		model["vlan_config"] = []map[string]interface{}{physicalTargetParamsForMountVolumeVlanConfigModel}

		assert.Equal(t, result, model)
	}

	physicalMountVolumesOriginalTargetConfigServerCredentialsModel := new(backuprecoveryv1.PhysicalMountVolumesOriginalTargetConfigServerCredentials)
	physicalMountVolumesOriginalTargetConfigServerCredentialsModel.Username = core.StringPtr("testString")
	physicalMountVolumesOriginalTargetConfigServerCredentialsModel.Password = core.StringPtr("testString")

	physicalTargetParamsForMountVolumeOriginalTargetConfigModel := new(backuprecoveryv1.PhysicalTargetParamsForMountVolumeOriginalTargetConfig)
	physicalTargetParamsForMountVolumeOriginalTargetConfigModel.ServerCredentials = physicalMountVolumesOriginalTargetConfigServerCredentialsModel

	recoverTargetModel := new(backuprecoveryv1.RecoverTarget)
	recoverTargetModel.ID = core.Int64Ptr(int64(26))

	physicalMountVolumesNewTargetConfigServerCredentialsModel := new(backuprecoveryv1.PhysicalMountVolumesNewTargetConfigServerCredentials)
	physicalMountVolumesNewTargetConfigServerCredentialsModel.Username = core.StringPtr("testString")
	physicalMountVolumesNewTargetConfigServerCredentialsModel.Password = core.StringPtr("testString")

	physicalTargetParamsForMountVolumeNewTargetConfigModel := new(backuprecoveryv1.PhysicalTargetParamsForMountVolumeNewTargetConfig)
	physicalTargetParamsForMountVolumeNewTargetConfigModel.MountTarget = recoverTargetModel
	physicalTargetParamsForMountVolumeNewTargetConfigModel.ServerCredentials = physicalMountVolumesNewTargetConfigServerCredentialsModel

	physicalTargetParamsForMountVolumeVlanConfigModel := new(backuprecoveryv1.PhysicalTargetParamsForMountVolumeVlanConfig)
	physicalTargetParamsForMountVolumeVlanConfigModel.ID = core.Int64Ptr(int64(38))
	physicalTargetParamsForMountVolumeVlanConfigModel.DisableVlan = core.BoolPtr(true)

	model := new(backuprecoveryv1.MountPhysicalVolumeParamsPhysicalTargetParams)
	model.MountToOriginalTarget = core.BoolPtr(true)
	model.OriginalTargetConfig = physicalTargetParamsForMountVolumeOriginalTargetConfigModel
	model.NewTargetConfig = physicalTargetParamsForMountVolumeNewTargetConfigModel
	model.ReadOnlyMount = core.BoolPtr(true)
	model.VolumeNames = []string{"testString"}
	model.MountedVolumeMapping = []backuprecoveryv1.MountedVolumeMapping{}
	model.VlanConfig = physicalTargetParamsForMountVolumeVlanConfigModel

	result, err := backuprecovery.DataSourceIbmRecoveryMountPhysicalVolumeParamsPhysicalTargetParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmRecoveryPhysicalTargetParamsForMountVolumeOriginalTargetConfigToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		physicalMountVolumesOriginalTargetConfigServerCredentialsModel := make(map[string]interface{})
		physicalMountVolumesOriginalTargetConfigServerCredentialsModel["username"] = "testString"
		physicalMountVolumesOriginalTargetConfigServerCredentialsModel["password"] = "testString"

		model := make(map[string]interface{})
		model["server_credentials"] = []map[string]interface{}{physicalMountVolumesOriginalTargetConfigServerCredentialsModel}

		assert.Equal(t, result, model)
	}

	physicalMountVolumesOriginalTargetConfigServerCredentialsModel := new(backuprecoveryv1.PhysicalMountVolumesOriginalTargetConfigServerCredentials)
	physicalMountVolumesOriginalTargetConfigServerCredentialsModel.Username = core.StringPtr("testString")
	physicalMountVolumesOriginalTargetConfigServerCredentialsModel.Password = core.StringPtr("testString")

	model := new(backuprecoveryv1.PhysicalTargetParamsForMountVolumeOriginalTargetConfig)
	model.ServerCredentials = physicalMountVolumesOriginalTargetConfigServerCredentialsModel

	result, err := backuprecovery.DataSourceIbmRecoveryPhysicalTargetParamsForMountVolumeOriginalTargetConfigToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmRecoveryPhysicalMountVolumesOriginalTargetConfigServerCredentialsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["username"] = "testString"
		model["password"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.PhysicalMountVolumesOriginalTargetConfigServerCredentials)
	model.Username = core.StringPtr("testString")
	model.Password = core.StringPtr("testString")

	result, err := backuprecovery.DataSourceIbmRecoveryPhysicalMountVolumesOriginalTargetConfigServerCredentialsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmRecoveryPhysicalTargetParamsForMountVolumeNewTargetConfigToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		recoverTargetModel := make(map[string]interface{})
		recoverTargetModel["id"] = int(26)

		physicalMountVolumesNewTargetConfigServerCredentialsModel := make(map[string]interface{})
		physicalMountVolumesNewTargetConfigServerCredentialsModel["username"] = "testString"
		physicalMountVolumesNewTargetConfigServerCredentialsModel["password"] = "testString"

		model := make(map[string]interface{})
		model["mount_target"] = []map[string]interface{}{recoverTargetModel}
		model["server_credentials"] = []map[string]interface{}{physicalMountVolumesNewTargetConfigServerCredentialsModel}

		assert.Equal(t, result, model)
	}

	recoverTargetModel := new(backuprecoveryv1.RecoverTarget)
	recoverTargetModel.ID = core.Int64Ptr(int64(26))

	physicalMountVolumesNewTargetConfigServerCredentialsModel := new(backuprecoveryv1.PhysicalMountVolumesNewTargetConfigServerCredentials)
	physicalMountVolumesNewTargetConfigServerCredentialsModel.Username = core.StringPtr("testString")
	physicalMountVolumesNewTargetConfigServerCredentialsModel.Password = core.StringPtr("testString")

	model := new(backuprecoveryv1.PhysicalTargetParamsForMountVolumeNewTargetConfig)
	model.MountTarget = recoverTargetModel
	model.ServerCredentials = physicalMountVolumesNewTargetConfigServerCredentialsModel

	result, err := backuprecovery.DataSourceIbmRecoveryPhysicalTargetParamsForMountVolumeNewTargetConfigToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmRecoveryRecoverTargetToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["id"] = int(26)
		model["name"] = "testString"
		model["parent_source_id"] = int(26)
		model["parent_source_name"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.RecoverTarget)
	model.ID = core.Int64Ptr(int64(26))
	model.Name = core.StringPtr("testString")
	model.ParentSourceID = core.Int64Ptr(int64(26))
	model.ParentSourceName = core.StringPtr("testString")

	result, err := backuprecovery.DataSourceIbmRecoveryRecoverTargetToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmRecoveryPhysicalMountVolumesNewTargetConfigServerCredentialsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["username"] = "testString"
		model["password"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.PhysicalMountVolumesNewTargetConfigServerCredentials)
	model.Username = core.StringPtr("testString")
	model.Password = core.StringPtr("testString")

	result, err := backuprecovery.DataSourceIbmRecoveryPhysicalMountVolumesNewTargetConfigServerCredentialsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmRecoveryMountedVolumeMappingToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["original_volume"] = "testString"
		model["mounted_volume"] = "testString"
		model["file_system_type"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.MountedVolumeMapping)
	model.OriginalVolume = core.StringPtr("testString")
	model.MountedVolume = core.StringPtr("testString")
	model.FileSystemType = core.StringPtr("testString")

	result, err := backuprecovery.DataSourceIbmRecoveryMountedVolumeMappingToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmRecoveryPhysicalTargetParamsForMountVolumeVlanConfigToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["id"] = int(38)
		model["disable_vlan"] = true
		model["interface_name"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.PhysicalTargetParamsForMountVolumeVlanConfig)
	model.ID = core.Int64Ptr(int64(38))
	model.DisableVlan = core.BoolPtr(true)
	model.InterfaceName = core.StringPtr("testString")

	result, err := backuprecovery.DataSourceIbmRecoveryPhysicalTargetParamsForMountVolumeVlanConfigToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmRecoveryRecoverPhysicalParamsRecoverFileAndFolderParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		commonRecoverFileAndFolderInfoModel := make(map[string]interface{})
		commonRecoverFileAndFolderInfoModel["absolute_path"] = "testString"
		commonRecoverFileAndFolderInfoModel["is_directory"] = true
		commonRecoverFileAndFolderInfoModel["is_view_file_recovery"] = true

		physicalTargetParamsForRecoverFileAndFolderRecoverTargetModel := make(map[string]interface{})
		physicalTargetParamsForRecoverFileAndFolderRecoverTargetModel["id"] = int(26)

		physicalTargetParamsForRecoverFileAndFolderVlanConfigModel := make(map[string]interface{})
		physicalTargetParamsForRecoverFileAndFolderVlanConfigModel["id"] = int(38)
		physicalTargetParamsForRecoverFileAndFolderVlanConfigModel["disable_vlan"] = true

		recoverPhysicalFileAndFolderParamsPhysicalTargetParamsModel := make(map[string]interface{})
		recoverPhysicalFileAndFolderParamsPhysicalTargetParamsModel["recover_target"] = []map[string]interface{}{physicalTargetParamsForRecoverFileAndFolderRecoverTargetModel}
		recoverPhysicalFileAndFolderParamsPhysicalTargetParamsModel["restore_to_original_paths"] = true
		recoverPhysicalFileAndFolderParamsPhysicalTargetParamsModel["overwrite_existing"] = true
		recoverPhysicalFileAndFolderParamsPhysicalTargetParamsModel["alternate_restore_directory"] = "testString"
		recoverPhysicalFileAndFolderParamsPhysicalTargetParamsModel["preserve_attributes"] = true
		recoverPhysicalFileAndFolderParamsPhysicalTargetParamsModel["preserve_timestamps"] = true
		recoverPhysicalFileAndFolderParamsPhysicalTargetParamsModel["preserve_acls"] = true
		recoverPhysicalFileAndFolderParamsPhysicalTargetParamsModel["continue_on_error"] = true
		recoverPhysicalFileAndFolderParamsPhysicalTargetParamsModel["save_success_files"] = true
		recoverPhysicalFileAndFolderParamsPhysicalTargetParamsModel["vlan_config"] = []map[string]interface{}{physicalTargetParamsForRecoverFileAndFolderVlanConfigModel}
		recoverPhysicalFileAndFolderParamsPhysicalTargetParamsModel["restore_entity_type"] = "kRegular"

		model := make(map[string]interface{})
		model["files_and_folders"] = []map[string]interface{}{commonRecoverFileAndFolderInfoModel}
		model["target_environment"] = "kPhysical"
		model["physical_target_params"] = []map[string]interface{}{recoverPhysicalFileAndFolderParamsPhysicalTargetParamsModel}

		assert.Equal(t, result, model)
	}

	commonRecoverFileAndFolderInfoModel := new(backuprecoveryv1.CommonRecoverFileAndFolderInfo)
	commonRecoverFileAndFolderInfoModel.AbsolutePath = core.StringPtr("testString")
	commonRecoverFileAndFolderInfoModel.IsDirectory = core.BoolPtr(true)
	commonRecoverFileAndFolderInfoModel.IsViewFileRecovery = core.BoolPtr(true)

	physicalTargetParamsForRecoverFileAndFolderRecoverTargetModel := new(backuprecoveryv1.PhysicalTargetParamsForRecoverFileAndFolderRecoverTarget)
	physicalTargetParamsForRecoverFileAndFolderRecoverTargetModel.ID = core.Int64Ptr(int64(26))

	physicalTargetParamsForRecoverFileAndFolderVlanConfigModel := new(backuprecoveryv1.PhysicalTargetParamsForRecoverFileAndFolderVlanConfig)
	physicalTargetParamsForRecoverFileAndFolderVlanConfigModel.ID = core.Int64Ptr(int64(38))
	physicalTargetParamsForRecoverFileAndFolderVlanConfigModel.DisableVlan = core.BoolPtr(true)

	recoverPhysicalFileAndFolderParamsPhysicalTargetParamsModel := new(backuprecoveryv1.RecoverPhysicalFileAndFolderParamsPhysicalTargetParams)
	recoverPhysicalFileAndFolderParamsPhysicalTargetParamsModel.RecoverTarget = physicalTargetParamsForRecoverFileAndFolderRecoverTargetModel
	recoverPhysicalFileAndFolderParamsPhysicalTargetParamsModel.RestoreToOriginalPaths = core.BoolPtr(true)
	recoverPhysicalFileAndFolderParamsPhysicalTargetParamsModel.OverwriteExisting = core.BoolPtr(true)
	recoverPhysicalFileAndFolderParamsPhysicalTargetParamsModel.AlternateRestoreDirectory = core.StringPtr("testString")
	recoverPhysicalFileAndFolderParamsPhysicalTargetParamsModel.PreserveAttributes = core.BoolPtr(true)
	recoverPhysicalFileAndFolderParamsPhysicalTargetParamsModel.PreserveTimestamps = core.BoolPtr(true)
	recoverPhysicalFileAndFolderParamsPhysicalTargetParamsModel.PreserveAcls = core.BoolPtr(true)
	recoverPhysicalFileAndFolderParamsPhysicalTargetParamsModel.ContinueOnError = core.BoolPtr(true)
	recoverPhysicalFileAndFolderParamsPhysicalTargetParamsModel.SaveSuccessFiles = core.BoolPtr(true)
	recoverPhysicalFileAndFolderParamsPhysicalTargetParamsModel.VlanConfig = physicalTargetParamsForRecoverFileAndFolderVlanConfigModel
	recoverPhysicalFileAndFolderParamsPhysicalTargetParamsModel.RestoreEntityType = core.StringPtr("kRegular")

	model := new(backuprecoveryv1.RecoverPhysicalParamsRecoverFileAndFolderParams)
	model.FilesAndFolders = []backuprecoveryv1.CommonRecoverFileAndFolderInfo{*commonRecoverFileAndFolderInfoModel}
	model.TargetEnvironment = core.StringPtr("kPhysical")
	model.PhysicalTargetParams = recoverPhysicalFileAndFolderParamsPhysicalTargetParamsModel

	result, err := backuprecovery.DataSourceIbmRecoveryRecoverPhysicalParamsRecoverFileAndFolderParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmRecoveryCommonRecoverFileAndFolderInfoToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["absolute_path"] = "testString"
		model["destination_dir"] = "testString"
		model["is_directory"] = true
		model["status"] = "NotStarted"
		model["messages"] = []string{"testString"}
		model["is_view_file_recovery"] = true

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.CommonRecoverFileAndFolderInfo)
	model.AbsolutePath = core.StringPtr("testString")
	model.DestinationDir = core.StringPtr("testString")
	model.IsDirectory = core.BoolPtr(true)
	model.Status = core.StringPtr("NotStarted")
	model.Messages = []string{"testString"}
	model.IsViewFileRecovery = core.BoolPtr(true)

	result, err := backuprecovery.DataSourceIbmRecoveryCommonRecoverFileAndFolderInfoToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmRecoveryRecoverPhysicalFileAndFolderParamsPhysicalTargetParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		physicalTargetParamsForRecoverFileAndFolderRecoverTargetModel := make(map[string]interface{})
		physicalTargetParamsForRecoverFileAndFolderRecoverTargetModel["id"] = int(26)

		physicalTargetParamsForRecoverFileAndFolderVlanConfigModel := make(map[string]interface{})
		physicalTargetParamsForRecoverFileAndFolderVlanConfigModel["id"] = int(38)
		physicalTargetParamsForRecoverFileAndFolderVlanConfigModel["disable_vlan"] = true

		model := make(map[string]interface{})
		model["recover_target"] = []map[string]interface{}{physicalTargetParamsForRecoverFileAndFolderRecoverTargetModel}
		model["restore_to_original_paths"] = true
		model["overwrite_existing"] = true
		model["alternate_restore_directory"] = "testString"
		model["preserve_attributes"] = true
		model["preserve_timestamps"] = true
		model["preserve_acls"] = true
		model["continue_on_error"] = true
		model["save_success_files"] = true
		model["vlan_config"] = []map[string]interface{}{physicalTargetParamsForRecoverFileAndFolderVlanConfigModel}
		model["restore_entity_type"] = "kRegular"

		assert.Equal(t, result, model)
	}

	physicalTargetParamsForRecoverFileAndFolderRecoverTargetModel := new(backuprecoveryv1.PhysicalTargetParamsForRecoverFileAndFolderRecoverTarget)
	physicalTargetParamsForRecoverFileAndFolderRecoverTargetModel.ID = core.Int64Ptr(int64(26))

	physicalTargetParamsForRecoverFileAndFolderVlanConfigModel := new(backuprecoveryv1.PhysicalTargetParamsForRecoverFileAndFolderVlanConfig)
	physicalTargetParamsForRecoverFileAndFolderVlanConfigModel.ID = core.Int64Ptr(int64(38))
	physicalTargetParamsForRecoverFileAndFolderVlanConfigModel.DisableVlan = core.BoolPtr(true)

	model := new(backuprecoveryv1.RecoverPhysicalFileAndFolderParamsPhysicalTargetParams)
	model.RecoverTarget = physicalTargetParamsForRecoverFileAndFolderRecoverTargetModel
	model.RestoreToOriginalPaths = core.BoolPtr(true)
	model.OverwriteExisting = core.BoolPtr(true)
	model.AlternateRestoreDirectory = core.StringPtr("testString")
	model.PreserveAttributes = core.BoolPtr(true)
	model.PreserveTimestamps = core.BoolPtr(true)
	model.PreserveAcls = core.BoolPtr(true)
	model.ContinueOnError = core.BoolPtr(true)
	model.SaveSuccessFiles = core.BoolPtr(true)
	model.VlanConfig = physicalTargetParamsForRecoverFileAndFolderVlanConfigModel
	model.RestoreEntityType = core.StringPtr("kRegular")

	result, err := backuprecovery.DataSourceIbmRecoveryRecoverPhysicalFileAndFolderParamsPhysicalTargetParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmRecoveryPhysicalTargetParamsForRecoverFileAndFolderRecoverTargetToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["id"] = int(26)
		model["name"] = "testString"
		model["parent_source_id"] = int(26)
		model["parent_source_name"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.PhysicalTargetParamsForRecoverFileAndFolderRecoverTarget)
	model.ID = core.Int64Ptr(int64(26))
	model.Name = core.StringPtr("testString")
	model.ParentSourceID = core.Int64Ptr(int64(26))
	model.ParentSourceName = core.StringPtr("testString")

	result, err := backuprecovery.DataSourceIbmRecoveryPhysicalTargetParamsForRecoverFileAndFolderRecoverTargetToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmRecoveryPhysicalTargetParamsForRecoverFileAndFolderVlanConfigToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["id"] = int(38)
		model["disable_vlan"] = true
		model["interface_name"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.PhysicalTargetParamsForRecoverFileAndFolderVlanConfig)
	model.ID = core.Int64Ptr(int64(38))
	model.DisableVlan = core.BoolPtr(true)
	model.InterfaceName = core.StringPtr("testString")

	result, err := backuprecovery.DataSourceIbmRecoveryPhysicalTargetParamsForRecoverFileAndFolderVlanConfigToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmRecoveryRecoverPhysicalParamsDownloadFileAndFolderParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		commonRecoverFileAndFolderInfoModel := make(map[string]interface{})
		commonRecoverFileAndFolderInfoModel["absolute_path"] = "testString"
		commonRecoverFileAndFolderInfoModel["is_directory"] = true
		commonRecoverFileAndFolderInfoModel["is_view_file_recovery"] = true

		model := make(map[string]interface{})
		model["expiry_time_usecs"] = int(26)
		model["files_and_folders"] = []map[string]interface{}{commonRecoverFileAndFolderInfoModel}
		model["download_file_path"] = "testString"

		assert.Equal(t, result, model)
	}

	commonRecoverFileAndFolderInfoModel := new(backuprecoveryv1.CommonRecoverFileAndFolderInfo)
	commonRecoverFileAndFolderInfoModel.AbsolutePath = core.StringPtr("testString")
	commonRecoverFileAndFolderInfoModel.IsDirectory = core.BoolPtr(true)
	commonRecoverFileAndFolderInfoModel.IsViewFileRecovery = core.BoolPtr(true)

	model := new(backuprecoveryv1.RecoverPhysicalParamsDownloadFileAndFolderParams)
	model.ExpiryTimeUsecs = core.Int64Ptr(int64(26))
	model.FilesAndFolders = []backuprecoveryv1.CommonRecoverFileAndFolderInfo{*commonRecoverFileAndFolderInfoModel}
	model.DownloadFilePath = core.StringPtr("testString")

	result, err := backuprecovery.DataSourceIbmRecoveryRecoverPhysicalParamsDownloadFileAndFolderParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmRecoveryRecoverPhysicalParamsSystemRecoveryParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["full_nas_path"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.RecoverPhysicalParamsSystemRecoveryParams)
	model.FullNasPath = core.StringPtr("testString")

	result, err := backuprecovery.DataSourceIbmRecoveryRecoverPhysicalParamsSystemRecoveryParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmRecoveryRecoverSqlParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		sharepointObjectParamsModel := make(map[string]interface{})
		sharepointObjectParamsModel["site_web_url"] = "testString"

		objectTypeVCenterParamsModel := make(map[string]interface{})
		objectTypeVCenterParamsModel["is_cloud_env"] = true

		objectTypeWindowsClusterParamsModel := make(map[string]interface{})
		objectTypeWindowsClusterParamsModel["cluster_source_type"] = "testString"

		objectSummaryModel := make(map[string]interface{})
		objectSummaryModel["id"] = int(26)
		objectSummaryModel["name"] = "testString"
		objectSummaryModel["source_id"] = int(26)
		objectSummaryModel["source_name"] = "testString"
		objectSummaryModel["environment"] = "kPhysical"
		objectSummaryModel["object_hash"] = "testString"
		objectSummaryModel["object_type"] = "kCluster"
		objectSummaryModel["logical_size_bytes"] = int(26)
		objectSummaryModel["uuid"] = "testString"
		objectSummaryModel["global_id"] = "testString"
		objectSummaryModel["protection_type"] = "kAgent"
		objectSummaryModel["sharepoint_site_summary"] = []map[string]interface{}{sharepointObjectParamsModel}
		objectSummaryModel["os_type"] = "kLinux"
		objectSummaryModel["v_center_summary"] = []map[string]interface{}{objectTypeVCenterParamsModel}
		objectSummaryModel["windows_cluster_summary"] = []map[string]interface{}{objectTypeWindowsClusterParamsModel}

		commonRecoverObjectSnapshotParamsObjectInfoModel := make(map[string]interface{})
		commonRecoverObjectSnapshotParamsObjectInfoModel["id"] = int(26)
		commonRecoverObjectSnapshotParamsObjectInfoModel["name"] = "testString"
		commonRecoverObjectSnapshotParamsObjectInfoModel["source_id"] = int(26)
		commonRecoverObjectSnapshotParamsObjectInfoModel["source_name"] = "testString"
		commonRecoverObjectSnapshotParamsObjectInfoModel["environment"] = "kPhysical"
		commonRecoverObjectSnapshotParamsObjectInfoModel["object_hash"] = "testString"
		commonRecoverObjectSnapshotParamsObjectInfoModel["object_type"] = "kCluster"
		commonRecoverObjectSnapshotParamsObjectInfoModel["logical_size_bytes"] = int(26)
		commonRecoverObjectSnapshotParamsObjectInfoModel["uuid"] = "testString"
		commonRecoverObjectSnapshotParamsObjectInfoModel["global_id"] = "testString"
		commonRecoverObjectSnapshotParamsObjectInfoModel["protection_type"] = "kAgent"
		commonRecoverObjectSnapshotParamsObjectInfoModel["sharepoint_site_summary"] = []map[string]interface{}{sharepointObjectParamsModel}
		commonRecoverObjectSnapshotParamsObjectInfoModel["os_type"] = "kLinux"
		commonRecoverObjectSnapshotParamsObjectInfoModel["child_objects"] = []map[string]interface{}{objectSummaryModel}
		commonRecoverObjectSnapshotParamsObjectInfoModel["v_center_summary"] = []map[string]interface{}{objectTypeVCenterParamsModel}
		commonRecoverObjectSnapshotParamsObjectInfoModel["windows_cluster_summary"] = []map[string]interface{}{objectTypeWindowsClusterParamsModel}

		awsTierModel := make(map[string]interface{})
		awsTierModel["move_after_unit"] = "Days"
		awsTierModel["move_after"] = int(26)
		awsTierModel["tier_type"] = "kAmazonS3Standard"

		awsTiersModel := make(map[string]interface{})
		awsTiersModel["tiers"] = []map[string]interface{}{awsTierModel}

		azureTierModel := make(map[string]interface{})
		azureTierModel["move_after_unit"] = "Days"
		azureTierModel["move_after"] = int(26)
		azureTierModel["tier_type"] = "kAzureTierHot"

		azureTiersModel := make(map[string]interface{})
		azureTiersModel["tiers"] = []map[string]interface{}{azureTierModel}

		googleTierModel := make(map[string]interface{})
		googleTierModel["move_after_unit"] = "Days"
		googleTierModel["move_after"] = int(26)
		googleTierModel["tier_type"] = "kGoogleStandard"

		googleTiersModel := make(map[string]interface{})
		googleTiersModel["tiers"] = []map[string]interface{}{googleTierModel}

		oracleTierModel := make(map[string]interface{})
		oracleTierModel["move_after_unit"] = "Days"
		oracleTierModel["move_after"] = int(26)
		oracleTierModel["tier_type"] = "kOracleTierStandard"

		oracleTiersModel := make(map[string]interface{})
		oracleTiersModel["tiers"] = []map[string]interface{}{oracleTierModel}

		archivalTargetTierInfoModel := make(map[string]interface{})
		archivalTargetTierInfoModel["aws_tiering"] = []map[string]interface{}{awsTiersModel}
		archivalTargetTierInfoModel["azure_tiering"] = []map[string]interface{}{azureTiersModel}
		archivalTargetTierInfoModel["cloud_platform"] = "AWS"
		archivalTargetTierInfoModel["google_tiering"] = []map[string]interface{}{googleTiersModel}
		archivalTargetTierInfoModel["oracle_tiering"] = []map[string]interface{}{oracleTiersModel}
		archivalTargetTierInfoModel["current_tier_type"] = "kAmazonS3Standard"

		commonRecoverObjectSnapshotParamsArchivalTargetInfoModel := make(map[string]interface{})
		commonRecoverObjectSnapshotParamsArchivalTargetInfoModel["target_id"] = int(26)
		commonRecoverObjectSnapshotParamsArchivalTargetInfoModel["archival_task_id"] = "testString"
		commonRecoverObjectSnapshotParamsArchivalTargetInfoModel["target_name"] = "testString"
		commonRecoverObjectSnapshotParamsArchivalTargetInfoModel["target_type"] = "Tape"
		commonRecoverObjectSnapshotParamsArchivalTargetInfoModel["usage_type"] = "Archival"
		commonRecoverObjectSnapshotParamsArchivalTargetInfoModel["ownership_context"] = "Local"
		commonRecoverObjectSnapshotParamsArchivalTargetInfoModel["tier_settings"] = []map[string]interface{}{archivalTargetTierInfoModel}

		aagInfoModel := make(map[string]interface{})
		aagInfoModel["name"] = "testString"
		aagInfoModel["object_id"] = int(26)

		hostInformationModel := make(map[string]interface{})
		hostInformationModel["id"] = "testString"
		hostInformationModel["name"] = "testString"
		hostInformationModel["environment"] = "kPhysical"

		multiStageRestoreOptionsModel := make(map[string]interface{})
		multiStageRestoreOptionsModel["enable_auto_sync"] = true
		multiStageRestoreOptionsModel["enable_multi_stage_restore"] = true

		filenamePatternToDirectoryModel := make(map[string]interface{})
		filenamePatternToDirectoryModel["directory"] = "testString"
		filenamePatternToDirectoryModel["filename_pattern"] = "testString"

		recoveryObjectIdentifierModel := make(map[string]interface{})
		recoveryObjectIdentifierModel["id"] = int(26)

		recoverSqlAppNewSourceConfigModel := make(map[string]interface{})
		recoverSqlAppNewSourceConfigModel["keep_cdc"] = true
		recoverSqlAppNewSourceConfigModel["multi_stage_restore_options"] = []map[string]interface{}{multiStageRestoreOptionsModel}
		recoverSqlAppNewSourceConfigModel["native_log_recovery_with_clause"] = "testString"
		recoverSqlAppNewSourceConfigModel["native_recovery_with_clause"] = "testString"
		recoverSqlAppNewSourceConfigModel["overwriting_policy"] = "FailIfExists"
		recoverSqlAppNewSourceConfigModel["replay_entire_last_log"] = true
		recoverSqlAppNewSourceConfigModel["restore_time_usecs"] = int(26)
		recoverSqlAppNewSourceConfigModel["secondary_data_files_dir_list"] = []map[string]interface{}{filenamePatternToDirectoryModel}
		recoverSqlAppNewSourceConfigModel["with_no_recovery"] = true
		recoverSqlAppNewSourceConfigModel["data_file_directory_location"] = "testString"
		recoverSqlAppNewSourceConfigModel["database_name"] = "testString"
		recoverSqlAppNewSourceConfigModel["host"] = []map[string]interface{}{recoveryObjectIdentifierModel}
		recoverSqlAppNewSourceConfigModel["instance_name"] = "testString"
		recoverSqlAppNewSourceConfigModel["log_file_directory_location"] = "testString"

		recoverSqlAppOriginalSourceConfigModel := make(map[string]interface{})
		recoverSqlAppOriginalSourceConfigModel["keep_cdc"] = true
		recoverSqlAppOriginalSourceConfigModel["multi_stage_restore_options"] = []map[string]interface{}{multiStageRestoreOptionsModel}
		recoverSqlAppOriginalSourceConfigModel["native_log_recovery_with_clause"] = "testString"
		recoverSqlAppOriginalSourceConfigModel["native_recovery_with_clause"] = "testString"
		recoverSqlAppOriginalSourceConfigModel["overwriting_policy"] = "FailIfExists"
		recoverSqlAppOriginalSourceConfigModel["replay_entire_last_log"] = true
		recoverSqlAppOriginalSourceConfigModel["restore_time_usecs"] = int(26)
		recoverSqlAppOriginalSourceConfigModel["secondary_data_files_dir_list"] = []map[string]interface{}{filenamePatternToDirectoryModel}
		recoverSqlAppOriginalSourceConfigModel["with_no_recovery"] = true
		recoverSqlAppOriginalSourceConfigModel["capture_tail_logs"] = true
		recoverSqlAppOriginalSourceConfigModel["data_file_directory_location"] = "testString"
		recoverSqlAppOriginalSourceConfigModel["log_file_directory_location"] = "testString"
		recoverSqlAppOriginalSourceConfigModel["new_database_name"] = "testString"

		sqlTargetParamsForRecoverSqlAppModel := make(map[string]interface{})
		sqlTargetParamsForRecoverSqlAppModel["new_source_config"] = []map[string]interface{}{recoverSqlAppNewSourceConfigModel}
		sqlTargetParamsForRecoverSqlAppModel["original_source_config"] = []map[string]interface{}{recoverSqlAppOriginalSourceConfigModel}
		sqlTargetParamsForRecoverSqlAppModel["recover_to_new_source"] = true

		recoverSqlAppParamsModel := make(map[string]interface{})
		recoverSqlAppParamsModel["snapshot_id"] = "testString"
		recoverSqlAppParamsModel["point_in_time_usecs"] = int(26)
		recoverSqlAppParamsModel["protection_group_id"] = "testString"
		recoverSqlAppParamsModel["protection_group_name"] = "testString"
		recoverSqlAppParamsModel["object_info"] = []map[string]interface{}{commonRecoverObjectSnapshotParamsObjectInfoModel}
		recoverSqlAppParamsModel["archival_target_info"] = []map[string]interface{}{commonRecoverObjectSnapshotParamsArchivalTargetInfoModel}
		recoverSqlAppParamsModel["recover_from_standby"] = true
		recoverSqlAppParamsModel["aag_info"] = []map[string]interface{}{aagInfoModel}
		recoverSqlAppParamsModel["host_info"] = []map[string]interface{}{hostInformationModel}
		recoverSqlAppParamsModel["is_encrypted"] = true
		recoverSqlAppParamsModel["sql_target_params"] = []map[string]interface{}{sqlTargetParamsForRecoverSqlAppModel}
		recoverSqlAppParamsModel["target_environment"] = "kSQL"

		recoveryVlanConfigModel := make(map[string]interface{})
		recoveryVlanConfigModel["id"] = int(38)
		recoveryVlanConfigModel["disable_vlan"] = true

		model := make(map[string]interface{})
		model["recover_app_params"] = []map[string]interface{}{recoverSqlAppParamsModel}
		model["recovery_action"] = "RecoverApps"
		model["vlan_config"] = []map[string]interface{}{recoveryVlanConfigModel}

		assert.Equal(t, result, model)
	}

	sharepointObjectParamsModel := new(backuprecoveryv1.SharepointObjectParams)
	sharepointObjectParamsModel.SiteWebURL = core.StringPtr("testString")

	objectTypeVCenterParamsModel := new(backuprecoveryv1.ObjectTypeVCenterParams)
	objectTypeVCenterParamsModel.IsCloudEnv = core.BoolPtr(true)

	objectTypeWindowsClusterParamsModel := new(backuprecoveryv1.ObjectTypeWindowsClusterParams)
	objectTypeWindowsClusterParamsModel.ClusterSourceType = core.StringPtr("testString")

	objectSummaryModel := new(backuprecoveryv1.ObjectSummary)
	objectSummaryModel.ID = core.Int64Ptr(int64(26))
	objectSummaryModel.Name = core.StringPtr("testString")
	objectSummaryModel.SourceID = core.Int64Ptr(int64(26))
	objectSummaryModel.SourceName = core.StringPtr("testString")
	objectSummaryModel.Environment = core.StringPtr("kPhysical")
	objectSummaryModel.ObjectHash = core.StringPtr("testString")
	objectSummaryModel.ObjectType = core.StringPtr("kCluster")
	objectSummaryModel.LogicalSizeBytes = core.Int64Ptr(int64(26))
	objectSummaryModel.UUID = core.StringPtr("testString")
	objectSummaryModel.GlobalID = core.StringPtr("testString")
	objectSummaryModel.ProtectionType = core.StringPtr("kAgent")
	objectSummaryModel.SharepointSiteSummary = sharepointObjectParamsModel
	objectSummaryModel.OsType = core.StringPtr("kLinux")
	objectSummaryModel.VCenterSummary = objectTypeVCenterParamsModel
	objectSummaryModel.WindowsClusterSummary = objectTypeWindowsClusterParamsModel

	commonRecoverObjectSnapshotParamsObjectInfoModel := new(backuprecoveryv1.CommonRecoverObjectSnapshotParamsObjectInfo)
	commonRecoverObjectSnapshotParamsObjectInfoModel.ID = core.Int64Ptr(int64(26))
	commonRecoverObjectSnapshotParamsObjectInfoModel.Name = core.StringPtr("testString")
	commonRecoverObjectSnapshotParamsObjectInfoModel.SourceID = core.Int64Ptr(int64(26))
	commonRecoverObjectSnapshotParamsObjectInfoModel.SourceName = core.StringPtr("testString")
	commonRecoverObjectSnapshotParamsObjectInfoModel.Environment = core.StringPtr("kPhysical")
	commonRecoverObjectSnapshotParamsObjectInfoModel.ObjectHash = core.StringPtr("testString")
	commonRecoverObjectSnapshotParamsObjectInfoModel.ObjectType = core.StringPtr("kCluster")
	commonRecoverObjectSnapshotParamsObjectInfoModel.LogicalSizeBytes = core.Int64Ptr(int64(26))
	commonRecoverObjectSnapshotParamsObjectInfoModel.UUID = core.StringPtr("testString")
	commonRecoverObjectSnapshotParamsObjectInfoModel.GlobalID = core.StringPtr("testString")
	commonRecoverObjectSnapshotParamsObjectInfoModel.ProtectionType = core.StringPtr("kAgent")
	commonRecoverObjectSnapshotParamsObjectInfoModel.SharepointSiteSummary = sharepointObjectParamsModel
	commonRecoverObjectSnapshotParamsObjectInfoModel.OsType = core.StringPtr("kLinux")
	commonRecoverObjectSnapshotParamsObjectInfoModel.ChildObjects = []backuprecoveryv1.ObjectSummary{*objectSummaryModel}
	commonRecoverObjectSnapshotParamsObjectInfoModel.VCenterSummary = objectTypeVCenterParamsModel
	commonRecoverObjectSnapshotParamsObjectInfoModel.WindowsClusterSummary = objectTypeWindowsClusterParamsModel

	awsTierModel := new(backuprecoveryv1.AWSTier)
	awsTierModel.MoveAfterUnit = core.StringPtr("Days")
	awsTierModel.MoveAfter = core.Int64Ptr(int64(26))
	awsTierModel.TierType = core.StringPtr("kAmazonS3Standard")

	awsTiersModel := new(backuprecoveryv1.AWSTiers)
	awsTiersModel.Tiers = []backuprecoveryv1.AWSTier{*awsTierModel}

	azureTierModel := new(backuprecoveryv1.AzureTier)
	azureTierModel.MoveAfterUnit = core.StringPtr("Days")
	azureTierModel.MoveAfter = core.Int64Ptr(int64(26))
	azureTierModel.TierType = core.StringPtr("kAzureTierHot")

	azureTiersModel := new(backuprecoveryv1.AzureTiers)
	azureTiersModel.Tiers = []backuprecoveryv1.AzureTier{*azureTierModel}

	googleTierModel := new(backuprecoveryv1.GoogleTier)
	googleTierModel.MoveAfterUnit = core.StringPtr("Days")
	googleTierModel.MoveAfter = core.Int64Ptr(int64(26))
	googleTierModel.TierType = core.StringPtr("kGoogleStandard")

	googleTiersModel := new(backuprecoveryv1.GoogleTiers)
	googleTiersModel.Tiers = []backuprecoveryv1.GoogleTier{*googleTierModel}

	oracleTierModel := new(backuprecoveryv1.OracleTier)
	oracleTierModel.MoveAfterUnit = core.StringPtr("Days")
	oracleTierModel.MoveAfter = core.Int64Ptr(int64(26))
	oracleTierModel.TierType = core.StringPtr("kOracleTierStandard")

	oracleTiersModel := new(backuprecoveryv1.OracleTiers)
	oracleTiersModel.Tiers = []backuprecoveryv1.OracleTier{*oracleTierModel}

	archivalTargetTierInfoModel := new(backuprecoveryv1.ArchivalTargetTierInfo)
	archivalTargetTierInfoModel.AwsTiering = awsTiersModel
	archivalTargetTierInfoModel.AzureTiering = azureTiersModel
	archivalTargetTierInfoModel.CloudPlatform = core.StringPtr("AWS")
	archivalTargetTierInfoModel.GoogleTiering = googleTiersModel
	archivalTargetTierInfoModel.OracleTiering = oracleTiersModel
	archivalTargetTierInfoModel.CurrentTierType = core.StringPtr("kAmazonS3Standard")

	commonRecoverObjectSnapshotParamsArchivalTargetInfoModel := new(backuprecoveryv1.CommonRecoverObjectSnapshotParamsArchivalTargetInfo)
	commonRecoverObjectSnapshotParamsArchivalTargetInfoModel.TargetID = core.Int64Ptr(int64(26))
	commonRecoverObjectSnapshotParamsArchivalTargetInfoModel.ArchivalTaskID = core.StringPtr("testString")
	commonRecoverObjectSnapshotParamsArchivalTargetInfoModel.TargetName = core.StringPtr("testString")
	commonRecoverObjectSnapshotParamsArchivalTargetInfoModel.TargetType = core.StringPtr("Tape")
	commonRecoverObjectSnapshotParamsArchivalTargetInfoModel.UsageType = core.StringPtr("Archival")
	commonRecoverObjectSnapshotParamsArchivalTargetInfoModel.OwnershipContext = core.StringPtr("Local")
	commonRecoverObjectSnapshotParamsArchivalTargetInfoModel.TierSettings = archivalTargetTierInfoModel

	aagInfoModel := new(backuprecoveryv1.AAGInfo)
	aagInfoModel.Name = core.StringPtr("testString")
	aagInfoModel.ObjectID = core.Int64Ptr(int64(26))

	hostInformationModel := new(backuprecoveryv1.HostInformation)
	hostInformationModel.ID = core.StringPtr("testString")
	hostInformationModel.Name = core.StringPtr("testString")
	hostInformationModel.Environment = core.StringPtr("kPhysical")

	multiStageRestoreOptionsModel := new(backuprecoveryv1.MultiStageRestoreOptions)
	multiStageRestoreOptionsModel.EnableAutoSync = core.BoolPtr(true)
	multiStageRestoreOptionsModel.EnableMultiStageRestore = core.BoolPtr(true)

	filenamePatternToDirectoryModel := new(backuprecoveryv1.FilenamePatternToDirectory)
	filenamePatternToDirectoryModel.Directory = core.StringPtr("testString")
	filenamePatternToDirectoryModel.FilenamePattern = core.StringPtr("testString")

	recoveryObjectIdentifierModel := new(backuprecoveryv1.RecoveryObjectIdentifier)
	recoveryObjectIdentifierModel.ID = core.Int64Ptr(int64(26))

	recoverSqlAppNewSourceConfigModel := new(backuprecoveryv1.RecoverSqlAppNewSourceConfig)
	recoverSqlAppNewSourceConfigModel.KeepCdc = core.BoolPtr(true)
	recoverSqlAppNewSourceConfigModel.MultiStageRestoreOptions = multiStageRestoreOptionsModel
	recoverSqlAppNewSourceConfigModel.NativeLogRecoveryWithClause = core.StringPtr("testString")
	recoverSqlAppNewSourceConfigModel.NativeRecoveryWithClause = core.StringPtr("testString")
	recoverSqlAppNewSourceConfigModel.OverwritingPolicy = core.StringPtr("FailIfExists")
	recoverSqlAppNewSourceConfigModel.ReplayEntireLastLog = core.BoolPtr(true)
	recoverSqlAppNewSourceConfigModel.RestoreTimeUsecs = core.Int64Ptr(int64(26))
	recoverSqlAppNewSourceConfigModel.SecondaryDataFilesDirList = []backuprecoveryv1.FilenamePatternToDirectory{*filenamePatternToDirectoryModel}
	recoverSqlAppNewSourceConfigModel.WithNoRecovery = core.BoolPtr(true)
	recoverSqlAppNewSourceConfigModel.DataFileDirectoryLocation = core.StringPtr("testString")
	recoverSqlAppNewSourceConfigModel.DatabaseName = core.StringPtr("testString")
	recoverSqlAppNewSourceConfigModel.Host = recoveryObjectIdentifierModel
	recoverSqlAppNewSourceConfigModel.InstanceName = core.StringPtr("testString")
	recoverSqlAppNewSourceConfigModel.LogFileDirectoryLocation = core.StringPtr("testString")

	recoverSqlAppOriginalSourceConfigModel := new(backuprecoveryv1.RecoverSqlAppOriginalSourceConfig)
	recoverSqlAppOriginalSourceConfigModel.KeepCdc = core.BoolPtr(true)
	recoverSqlAppOriginalSourceConfigModel.MultiStageRestoreOptions = multiStageRestoreOptionsModel
	recoverSqlAppOriginalSourceConfigModel.NativeLogRecoveryWithClause = core.StringPtr("testString")
	recoverSqlAppOriginalSourceConfigModel.NativeRecoveryWithClause = core.StringPtr("testString")
	recoverSqlAppOriginalSourceConfigModel.OverwritingPolicy = core.StringPtr("FailIfExists")
	recoverSqlAppOriginalSourceConfigModel.ReplayEntireLastLog = core.BoolPtr(true)
	recoverSqlAppOriginalSourceConfigModel.RestoreTimeUsecs = core.Int64Ptr(int64(26))
	recoverSqlAppOriginalSourceConfigModel.SecondaryDataFilesDirList = []backuprecoveryv1.FilenamePatternToDirectory{*filenamePatternToDirectoryModel}
	recoverSqlAppOriginalSourceConfigModel.WithNoRecovery = core.BoolPtr(true)
	recoverSqlAppOriginalSourceConfigModel.CaptureTailLogs = core.BoolPtr(true)
	recoverSqlAppOriginalSourceConfigModel.DataFileDirectoryLocation = core.StringPtr("testString")
	recoverSqlAppOriginalSourceConfigModel.LogFileDirectoryLocation = core.StringPtr("testString")
	recoverSqlAppOriginalSourceConfigModel.NewDatabaseName = core.StringPtr("testString")

	sqlTargetParamsForRecoverSqlAppModel := new(backuprecoveryv1.SqlTargetParamsForRecoverSqlApp)
	sqlTargetParamsForRecoverSqlAppModel.NewSourceConfig = recoverSqlAppNewSourceConfigModel
	sqlTargetParamsForRecoverSqlAppModel.OriginalSourceConfig = recoverSqlAppOriginalSourceConfigModel
	sqlTargetParamsForRecoverSqlAppModel.RecoverToNewSource = core.BoolPtr(true)

	recoverSqlAppParamsModel := new(backuprecoveryv1.RecoverSqlAppParams)
	recoverSqlAppParamsModel.SnapshotID = core.StringPtr("testString")
	recoverSqlAppParamsModel.PointInTimeUsecs = core.Int64Ptr(int64(26))
	recoverSqlAppParamsModel.ProtectionGroupID = core.StringPtr("testString")
	recoverSqlAppParamsModel.ProtectionGroupName = core.StringPtr("testString")
	recoverSqlAppParamsModel.ObjectInfo = commonRecoverObjectSnapshotParamsObjectInfoModel
	recoverSqlAppParamsModel.ArchivalTargetInfo = commonRecoverObjectSnapshotParamsArchivalTargetInfoModel
	recoverSqlAppParamsModel.RecoverFromStandby = core.BoolPtr(true)
	recoverSqlAppParamsModel.AagInfo = aagInfoModel
	recoverSqlAppParamsModel.HostInfo = hostInformationModel
	recoverSqlAppParamsModel.IsEncrypted = core.BoolPtr(true)
	recoverSqlAppParamsModel.SqlTargetParams = sqlTargetParamsForRecoverSqlAppModel
	recoverSqlAppParamsModel.TargetEnvironment = core.StringPtr("kSQL")

	recoveryVlanConfigModel := new(backuprecoveryv1.RecoveryVlanConfig)
	recoveryVlanConfigModel.ID = core.Int64Ptr(int64(38))
	recoveryVlanConfigModel.DisableVlan = core.BoolPtr(true)

	model := new(backuprecoveryv1.RecoverSqlParams)
	model.RecoverAppParams = []backuprecoveryv1.RecoverSqlAppParams{*recoverSqlAppParamsModel}
	model.RecoveryAction = core.StringPtr("RecoverApps")
	model.VlanConfig = recoveryVlanConfigModel

	result, err := backuprecovery.DataSourceIbmRecoveryRecoverSqlParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmRecoveryRecoverSqlAppParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		sharepointObjectParamsModel := make(map[string]interface{})
		sharepointObjectParamsModel["site_web_url"] = "testString"

		objectTypeVCenterParamsModel := make(map[string]interface{})
		objectTypeVCenterParamsModel["is_cloud_env"] = true

		objectTypeWindowsClusterParamsModel := make(map[string]interface{})
		objectTypeWindowsClusterParamsModel["cluster_source_type"] = "testString"

		objectSummaryModel := make(map[string]interface{})
		objectSummaryModel["id"] = int(26)
		objectSummaryModel["name"] = "testString"
		objectSummaryModel["source_id"] = int(26)
		objectSummaryModel["source_name"] = "testString"
		objectSummaryModel["environment"] = "kPhysical"
		objectSummaryModel["object_hash"] = "testString"
		objectSummaryModel["object_type"] = "kCluster"
		objectSummaryModel["logical_size_bytes"] = int(26)
		objectSummaryModel["uuid"] = "testString"
		objectSummaryModel["global_id"] = "testString"
		objectSummaryModel["protection_type"] = "kAgent"
		objectSummaryModel["sharepoint_site_summary"] = []map[string]interface{}{sharepointObjectParamsModel}
		objectSummaryModel["os_type"] = "kLinux"
		objectSummaryModel["v_center_summary"] = []map[string]interface{}{objectTypeVCenterParamsModel}
		objectSummaryModel["windows_cluster_summary"] = []map[string]interface{}{objectTypeWindowsClusterParamsModel}

		commonRecoverObjectSnapshotParamsObjectInfoModel := make(map[string]interface{})
		commonRecoverObjectSnapshotParamsObjectInfoModel["id"] = int(26)
		commonRecoverObjectSnapshotParamsObjectInfoModel["name"] = "testString"
		commonRecoverObjectSnapshotParamsObjectInfoModel["source_id"] = int(26)
		commonRecoverObjectSnapshotParamsObjectInfoModel["source_name"] = "testString"
		commonRecoverObjectSnapshotParamsObjectInfoModel["environment"] = "kPhysical"
		commonRecoverObjectSnapshotParamsObjectInfoModel["object_hash"] = "testString"
		commonRecoverObjectSnapshotParamsObjectInfoModel["object_type"] = "kCluster"
		commonRecoverObjectSnapshotParamsObjectInfoModel["logical_size_bytes"] = int(26)
		commonRecoverObjectSnapshotParamsObjectInfoModel["uuid"] = "testString"
		commonRecoverObjectSnapshotParamsObjectInfoModel["global_id"] = "testString"
		commonRecoverObjectSnapshotParamsObjectInfoModel["protection_type"] = "kAgent"
		commonRecoverObjectSnapshotParamsObjectInfoModel["sharepoint_site_summary"] = []map[string]interface{}{sharepointObjectParamsModel}
		commonRecoverObjectSnapshotParamsObjectInfoModel["os_type"] = "kLinux"
		commonRecoverObjectSnapshotParamsObjectInfoModel["child_objects"] = []map[string]interface{}{objectSummaryModel}
		commonRecoverObjectSnapshotParamsObjectInfoModel["v_center_summary"] = []map[string]interface{}{objectTypeVCenterParamsModel}
		commonRecoverObjectSnapshotParamsObjectInfoModel["windows_cluster_summary"] = []map[string]interface{}{objectTypeWindowsClusterParamsModel}

		awsTierModel := make(map[string]interface{})
		awsTierModel["move_after_unit"] = "Days"
		awsTierModel["move_after"] = int(26)
		awsTierModel["tier_type"] = "kAmazonS3Standard"

		awsTiersModel := make(map[string]interface{})
		awsTiersModel["tiers"] = []map[string]interface{}{awsTierModel}

		azureTierModel := make(map[string]interface{})
		azureTierModel["move_after_unit"] = "Days"
		azureTierModel["move_after"] = int(26)
		azureTierModel["tier_type"] = "kAzureTierHot"

		azureTiersModel := make(map[string]interface{})
		azureTiersModel["tiers"] = []map[string]interface{}{azureTierModel}

		googleTierModel := make(map[string]interface{})
		googleTierModel["move_after_unit"] = "Days"
		googleTierModel["move_after"] = int(26)
		googleTierModel["tier_type"] = "kGoogleStandard"

		googleTiersModel := make(map[string]interface{})
		googleTiersModel["tiers"] = []map[string]interface{}{googleTierModel}

		oracleTierModel := make(map[string]interface{})
		oracleTierModel["move_after_unit"] = "Days"
		oracleTierModel["move_after"] = int(26)
		oracleTierModel["tier_type"] = "kOracleTierStandard"

		oracleTiersModel := make(map[string]interface{})
		oracleTiersModel["tiers"] = []map[string]interface{}{oracleTierModel}

		archivalTargetTierInfoModel := make(map[string]interface{})
		archivalTargetTierInfoModel["aws_tiering"] = []map[string]interface{}{awsTiersModel}
		archivalTargetTierInfoModel["azure_tiering"] = []map[string]interface{}{azureTiersModel}
		archivalTargetTierInfoModel["cloud_platform"] = "AWS"
		archivalTargetTierInfoModel["google_tiering"] = []map[string]interface{}{googleTiersModel}
		archivalTargetTierInfoModel["oracle_tiering"] = []map[string]interface{}{oracleTiersModel}
		archivalTargetTierInfoModel["current_tier_type"] = "kAmazonS3Standard"

		commonRecoverObjectSnapshotParamsArchivalTargetInfoModel := make(map[string]interface{})
		commonRecoverObjectSnapshotParamsArchivalTargetInfoModel["target_id"] = int(26)
		commonRecoverObjectSnapshotParamsArchivalTargetInfoModel["archival_task_id"] = "testString"
		commonRecoverObjectSnapshotParamsArchivalTargetInfoModel["target_name"] = "testString"
		commonRecoverObjectSnapshotParamsArchivalTargetInfoModel["target_type"] = "Tape"
		commonRecoverObjectSnapshotParamsArchivalTargetInfoModel["usage_type"] = "Archival"
		commonRecoverObjectSnapshotParamsArchivalTargetInfoModel["ownership_context"] = "Local"
		commonRecoverObjectSnapshotParamsArchivalTargetInfoModel["tier_settings"] = []map[string]interface{}{archivalTargetTierInfoModel}

		aagInfoModel := make(map[string]interface{})
		aagInfoModel["name"] = "testString"
		aagInfoModel["object_id"] = int(26)

		hostInformationModel := make(map[string]interface{})
		hostInformationModel["id"] = "testString"
		hostInformationModel["name"] = "testString"
		hostInformationModel["environment"] = "kPhysical"

		multiStageRestoreOptionsModel := make(map[string]interface{})
		multiStageRestoreOptionsModel["enable_auto_sync"] = true
		multiStageRestoreOptionsModel["enable_multi_stage_restore"] = true

		filenamePatternToDirectoryModel := make(map[string]interface{})
		filenamePatternToDirectoryModel["directory"] = "testString"
		filenamePatternToDirectoryModel["filename_pattern"] = "testString"

		recoveryObjectIdentifierModel := make(map[string]interface{})
		recoveryObjectIdentifierModel["id"] = int(26)

		recoverSqlAppNewSourceConfigModel := make(map[string]interface{})
		recoverSqlAppNewSourceConfigModel["keep_cdc"] = true
		recoverSqlAppNewSourceConfigModel["multi_stage_restore_options"] = []map[string]interface{}{multiStageRestoreOptionsModel}
		recoverSqlAppNewSourceConfigModel["native_log_recovery_with_clause"] = "testString"
		recoverSqlAppNewSourceConfigModel["native_recovery_with_clause"] = "testString"
		recoverSqlAppNewSourceConfigModel["overwriting_policy"] = "FailIfExists"
		recoverSqlAppNewSourceConfigModel["replay_entire_last_log"] = true
		recoverSqlAppNewSourceConfigModel["restore_time_usecs"] = int(26)
		recoverSqlAppNewSourceConfigModel["secondary_data_files_dir_list"] = []map[string]interface{}{filenamePatternToDirectoryModel}
		recoverSqlAppNewSourceConfigModel["with_no_recovery"] = true
		recoverSqlAppNewSourceConfigModel["data_file_directory_location"] = "testString"
		recoverSqlAppNewSourceConfigModel["database_name"] = "testString"
		recoverSqlAppNewSourceConfigModel["host"] = []map[string]interface{}{recoveryObjectIdentifierModel}
		recoverSqlAppNewSourceConfigModel["instance_name"] = "testString"
		recoverSqlAppNewSourceConfigModel["log_file_directory_location"] = "testString"

		recoverSqlAppOriginalSourceConfigModel := make(map[string]interface{})
		recoverSqlAppOriginalSourceConfigModel["keep_cdc"] = true
		recoverSqlAppOriginalSourceConfigModel["multi_stage_restore_options"] = []map[string]interface{}{multiStageRestoreOptionsModel}
		recoverSqlAppOriginalSourceConfigModel["native_log_recovery_with_clause"] = "testString"
		recoverSqlAppOriginalSourceConfigModel["native_recovery_with_clause"] = "testString"
		recoverSqlAppOriginalSourceConfigModel["overwriting_policy"] = "FailIfExists"
		recoverSqlAppOriginalSourceConfigModel["replay_entire_last_log"] = true
		recoverSqlAppOriginalSourceConfigModel["restore_time_usecs"] = int(26)
		recoverSqlAppOriginalSourceConfigModel["secondary_data_files_dir_list"] = []map[string]interface{}{filenamePatternToDirectoryModel}
		recoverSqlAppOriginalSourceConfigModel["with_no_recovery"] = true
		recoverSqlAppOriginalSourceConfigModel["capture_tail_logs"] = true
		recoverSqlAppOriginalSourceConfigModel["data_file_directory_location"] = "testString"
		recoverSqlAppOriginalSourceConfigModel["log_file_directory_location"] = "testString"
		recoverSqlAppOriginalSourceConfigModel["new_database_name"] = "testString"

		sqlTargetParamsForRecoverSqlAppModel := make(map[string]interface{})
		sqlTargetParamsForRecoverSqlAppModel["new_source_config"] = []map[string]interface{}{recoverSqlAppNewSourceConfigModel}
		sqlTargetParamsForRecoverSqlAppModel["original_source_config"] = []map[string]interface{}{recoverSqlAppOriginalSourceConfigModel}
		sqlTargetParamsForRecoverSqlAppModel["recover_to_new_source"] = true

		model := make(map[string]interface{})
		model["snapshot_id"] = "testString"
		model["point_in_time_usecs"] = int(26)
		model["protection_group_id"] = "testString"
		model["protection_group_name"] = "testString"
		model["snapshot_creation_time_usecs"] = int(26)
		model["object_info"] = []map[string]interface{}{commonRecoverObjectSnapshotParamsObjectInfoModel}
		model["snapshot_target_type"] = "Local"
		model["storage_domain_id"] = int(26)
		model["archival_target_info"] = []map[string]interface{}{commonRecoverObjectSnapshotParamsArchivalTargetInfoModel}
		model["progress_task_id"] = "testString"
		model["recover_from_standby"] = true
		model["status"] = "Accepted"
		model["start_time_usecs"] = int(26)
		model["end_time_usecs"] = int(26)
		model["messages"] = []string{"testString"}
		model["bytes_restored"] = int(26)
		model["aag_info"] = []map[string]interface{}{aagInfoModel}
		model["host_info"] = []map[string]interface{}{hostInformationModel}
		model["is_encrypted"] = true
		model["sql_target_params"] = []map[string]interface{}{sqlTargetParamsForRecoverSqlAppModel}
		model["target_environment"] = "kSQL"

		assert.Equal(t, result, model)
	}

	sharepointObjectParamsModel := new(backuprecoveryv1.SharepointObjectParams)
	sharepointObjectParamsModel.SiteWebURL = core.StringPtr("testString")

	objectTypeVCenterParamsModel := new(backuprecoveryv1.ObjectTypeVCenterParams)
	objectTypeVCenterParamsModel.IsCloudEnv = core.BoolPtr(true)

	objectTypeWindowsClusterParamsModel := new(backuprecoveryv1.ObjectTypeWindowsClusterParams)
	objectTypeWindowsClusterParamsModel.ClusterSourceType = core.StringPtr("testString")

	objectSummaryModel := new(backuprecoveryv1.ObjectSummary)
	objectSummaryModel.ID = core.Int64Ptr(int64(26))
	objectSummaryModel.Name = core.StringPtr("testString")
	objectSummaryModel.SourceID = core.Int64Ptr(int64(26))
	objectSummaryModel.SourceName = core.StringPtr("testString")
	objectSummaryModel.Environment = core.StringPtr("kPhysical")
	objectSummaryModel.ObjectHash = core.StringPtr("testString")
	objectSummaryModel.ObjectType = core.StringPtr("kCluster")
	objectSummaryModel.LogicalSizeBytes = core.Int64Ptr(int64(26))
	objectSummaryModel.UUID = core.StringPtr("testString")
	objectSummaryModel.GlobalID = core.StringPtr("testString")
	objectSummaryModel.ProtectionType = core.StringPtr("kAgent")
	objectSummaryModel.SharepointSiteSummary = sharepointObjectParamsModel
	objectSummaryModel.OsType = core.StringPtr("kLinux")
	objectSummaryModel.VCenterSummary = objectTypeVCenterParamsModel
	objectSummaryModel.WindowsClusterSummary = objectTypeWindowsClusterParamsModel

	commonRecoverObjectSnapshotParamsObjectInfoModel := new(backuprecoveryv1.CommonRecoverObjectSnapshotParamsObjectInfo)
	commonRecoverObjectSnapshotParamsObjectInfoModel.ID = core.Int64Ptr(int64(26))
	commonRecoverObjectSnapshotParamsObjectInfoModel.Name = core.StringPtr("testString")
	commonRecoverObjectSnapshotParamsObjectInfoModel.SourceID = core.Int64Ptr(int64(26))
	commonRecoverObjectSnapshotParamsObjectInfoModel.SourceName = core.StringPtr("testString")
	commonRecoverObjectSnapshotParamsObjectInfoModel.Environment = core.StringPtr("kPhysical")
	commonRecoverObjectSnapshotParamsObjectInfoModel.ObjectHash = core.StringPtr("testString")
	commonRecoverObjectSnapshotParamsObjectInfoModel.ObjectType = core.StringPtr("kCluster")
	commonRecoverObjectSnapshotParamsObjectInfoModel.LogicalSizeBytes = core.Int64Ptr(int64(26))
	commonRecoverObjectSnapshotParamsObjectInfoModel.UUID = core.StringPtr("testString")
	commonRecoverObjectSnapshotParamsObjectInfoModel.GlobalID = core.StringPtr("testString")
	commonRecoverObjectSnapshotParamsObjectInfoModel.ProtectionType = core.StringPtr("kAgent")
	commonRecoverObjectSnapshotParamsObjectInfoModel.SharepointSiteSummary = sharepointObjectParamsModel
	commonRecoverObjectSnapshotParamsObjectInfoModel.OsType = core.StringPtr("kLinux")
	commonRecoverObjectSnapshotParamsObjectInfoModel.ChildObjects = []backuprecoveryv1.ObjectSummary{*objectSummaryModel}
	commonRecoverObjectSnapshotParamsObjectInfoModel.VCenterSummary = objectTypeVCenterParamsModel
	commonRecoverObjectSnapshotParamsObjectInfoModel.WindowsClusterSummary = objectTypeWindowsClusterParamsModel

	awsTierModel := new(backuprecoveryv1.AWSTier)
	awsTierModel.MoveAfterUnit = core.StringPtr("Days")
	awsTierModel.MoveAfter = core.Int64Ptr(int64(26))
	awsTierModel.TierType = core.StringPtr("kAmazonS3Standard")

	awsTiersModel := new(backuprecoveryv1.AWSTiers)
	awsTiersModel.Tiers = []backuprecoveryv1.AWSTier{*awsTierModel}

	azureTierModel := new(backuprecoveryv1.AzureTier)
	azureTierModel.MoveAfterUnit = core.StringPtr("Days")
	azureTierModel.MoveAfter = core.Int64Ptr(int64(26))
	azureTierModel.TierType = core.StringPtr("kAzureTierHot")

	azureTiersModel := new(backuprecoveryv1.AzureTiers)
	azureTiersModel.Tiers = []backuprecoveryv1.AzureTier{*azureTierModel}

	googleTierModel := new(backuprecoveryv1.GoogleTier)
	googleTierModel.MoveAfterUnit = core.StringPtr("Days")
	googleTierModel.MoveAfter = core.Int64Ptr(int64(26))
	googleTierModel.TierType = core.StringPtr("kGoogleStandard")

	googleTiersModel := new(backuprecoveryv1.GoogleTiers)
	googleTiersModel.Tiers = []backuprecoveryv1.GoogleTier{*googleTierModel}

	oracleTierModel := new(backuprecoveryv1.OracleTier)
	oracleTierModel.MoveAfterUnit = core.StringPtr("Days")
	oracleTierModel.MoveAfter = core.Int64Ptr(int64(26))
	oracleTierModel.TierType = core.StringPtr("kOracleTierStandard")

	oracleTiersModel := new(backuprecoveryv1.OracleTiers)
	oracleTiersModel.Tiers = []backuprecoveryv1.OracleTier{*oracleTierModel}

	archivalTargetTierInfoModel := new(backuprecoveryv1.ArchivalTargetTierInfo)
	archivalTargetTierInfoModel.AwsTiering = awsTiersModel
	archivalTargetTierInfoModel.AzureTiering = azureTiersModel
	archivalTargetTierInfoModel.CloudPlatform = core.StringPtr("AWS")
	archivalTargetTierInfoModel.GoogleTiering = googleTiersModel
	archivalTargetTierInfoModel.OracleTiering = oracleTiersModel
	archivalTargetTierInfoModel.CurrentTierType = core.StringPtr("kAmazonS3Standard")

	commonRecoverObjectSnapshotParamsArchivalTargetInfoModel := new(backuprecoveryv1.CommonRecoverObjectSnapshotParamsArchivalTargetInfo)
	commonRecoverObjectSnapshotParamsArchivalTargetInfoModel.TargetID = core.Int64Ptr(int64(26))
	commonRecoverObjectSnapshotParamsArchivalTargetInfoModel.ArchivalTaskID = core.StringPtr("testString")
	commonRecoverObjectSnapshotParamsArchivalTargetInfoModel.TargetName = core.StringPtr("testString")
	commonRecoverObjectSnapshotParamsArchivalTargetInfoModel.TargetType = core.StringPtr("Tape")
	commonRecoverObjectSnapshotParamsArchivalTargetInfoModel.UsageType = core.StringPtr("Archival")
	commonRecoverObjectSnapshotParamsArchivalTargetInfoModel.OwnershipContext = core.StringPtr("Local")
	commonRecoverObjectSnapshotParamsArchivalTargetInfoModel.TierSettings = archivalTargetTierInfoModel

	aagInfoModel := new(backuprecoveryv1.AAGInfo)
	aagInfoModel.Name = core.StringPtr("testString")
	aagInfoModel.ObjectID = core.Int64Ptr(int64(26))

	hostInformationModel := new(backuprecoveryv1.HostInformation)
	hostInformationModel.ID = core.StringPtr("testString")
	hostInformationModel.Name = core.StringPtr("testString")
	hostInformationModel.Environment = core.StringPtr("kPhysical")

	multiStageRestoreOptionsModel := new(backuprecoveryv1.MultiStageRestoreOptions)
	multiStageRestoreOptionsModel.EnableAutoSync = core.BoolPtr(true)
	multiStageRestoreOptionsModel.EnableMultiStageRestore = core.BoolPtr(true)

	filenamePatternToDirectoryModel := new(backuprecoveryv1.FilenamePatternToDirectory)
	filenamePatternToDirectoryModel.Directory = core.StringPtr("testString")
	filenamePatternToDirectoryModel.FilenamePattern = core.StringPtr("testString")

	recoveryObjectIdentifierModel := new(backuprecoveryv1.RecoveryObjectIdentifier)
	recoveryObjectIdentifierModel.ID = core.Int64Ptr(int64(26))

	recoverSqlAppNewSourceConfigModel := new(backuprecoveryv1.RecoverSqlAppNewSourceConfig)
	recoverSqlAppNewSourceConfigModel.KeepCdc = core.BoolPtr(true)
	recoverSqlAppNewSourceConfigModel.MultiStageRestoreOptions = multiStageRestoreOptionsModel
	recoverSqlAppNewSourceConfigModel.NativeLogRecoveryWithClause = core.StringPtr("testString")
	recoverSqlAppNewSourceConfigModel.NativeRecoveryWithClause = core.StringPtr("testString")
	recoverSqlAppNewSourceConfigModel.OverwritingPolicy = core.StringPtr("FailIfExists")
	recoverSqlAppNewSourceConfigModel.ReplayEntireLastLog = core.BoolPtr(true)
	recoverSqlAppNewSourceConfigModel.RestoreTimeUsecs = core.Int64Ptr(int64(26))
	recoverSqlAppNewSourceConfigModel.SecondaryDataFilesDirList = []backuprecoveryv1.FilenamePatternToDirectory{*filenamePatternToDirectoryModel}
	recoverSqlAppNewSourceConfigModel.WithNoRecovery = core.BoolPtr(true)
	recoverSqlAppNewSourceConfigModel.DataFileDirectoryLocation = core.StringPtr("testString")
	recoverSqlAppNewSourceConfigModel.DatabaseName = core.StringPtr("testString")
	recoverSqlAppNewSourceConfigModel.Host = recoveryObjectIdentifierModel
	recoverSqlAppNewSourceConfigModel.InstanceName = core.StringPtr("testString")
	recoverSqlAppNewSourceConfigModel.LogFileDirectoryLocation = core.StringPtr("testString")

	recoverSqlAppOriginalSourceConfigModel := new(backuprecoveryv1.RecoverSqlAppOriginalSourceConfig)
	recoverSqlAppOriginalSourceConfigModel.KeepCdc = core.BoolPtr(true)
	recoverSqlAppOriginalSourceConfigModel.MultiStageRestoreOptions = multiStageRestoreOptionsModel
	recoverSqlAppOriginalSourceConfigModel.NativeLogRecoveryWithClause = core.StringPtr("testString")
	recoverSqlAppOriginalSourceConfigModel.NativeRecoveryWithClause = core.StringPtr("testString")
	recoverSqlAppOriginalSourceConfigModel.OverwritingPolicy = core.StringPtr("FailIfExists")
	recoverSqlAppOriginalSourceConfigModel.ReplayEntireLastLog = core.BoolPtr(true)
	recoverSqlAppOriginalSourceConfigModel.RestoreTimeUsecs = core.Int64Ptr(int64(26))
	recoverSqlAppOriginalSourceConfigModel.SecondaryDataFilesDirList = []backuprecoveryv1.FilenamePatternToDirectory{*filenamePatternToDirectoryModel}
	recoverSqlAppOriginalSourceConfigModel.WithNoRecovery = core.BoolPtr(true)
	recoverSqlAppOriginalSourceConfigModel.CaptureTailLogs = core.BoolPtr(true)
	recoverSqlAppOriginalSourceConfigModel.DataFileDirectoryLocation = core.StringPtr("testString")
	recoverSqlAppOriginalSourceConfigModel.LogFileDirectoryLocation = core.StringPtr("testString")
	recoverSqlAppOriginalSourceConfigModel.NewDatabaseName = core.StringPtr("testString")

	sqlTargetParamsForRecoverSqlAppModel := new(backuprecoveryv1.SqlTargetParamsForRecoverSqlApp)
	sqlTargetParamsForRecoverSqlAppModel.NewSourceConfig = recoverSqlAppNewSourceConfigModel
	sqlTargetParamsForRecoverSqlAppModel.OriginalSourceConfig = recoverSqlAppOriginalSourceConfigModel
	sqlTargetParamsForRecoverSqlAppModel.RecoverToNewSource = core.BoolPtr(true)

	model := new(backuprecoveryv1.RecoverSqlAppParams)
	model.SnapshotID = core.StringPtr("testString")
	model.PointInTimeUsecs = core.Int64Ptr(int64(26))
	model.ProtectionGroupID = core.StringPtr("testString")
	model.ProtectionGroupName = core.StringPtr("testString")
	model.SnapshotCreationTimeUsecs = core.Int64Ptr(int64(26))
	model.ObjectInfo = commonRecoverObjectSnapshotParamsObjectInfoModel
	model.SnapshotTargetType = core.StringPtr("Local")
	model.StorageDomainID = core.Int64Ptr(int64(26))
	model.ArchivalTargetInfo = commonRecoverObjectSnapshotParamsArchivalTargetInfoModel
	model.ProgressTaskID = core.StringPtr("testString")
	model.RecoverFromStandby = core.BoolPtr(true)
	model.Status = core.StringPtr("Accepted")
	model.StartTimeUsecs = core.Int64Ptr(int64(26))
	model.EndTimeUsecs = core.Int64Ptr(int64(26))
	model.Messages = []string{"testString"}
	model.BytesRestored = core.Int64Ptr(int64(26))
	model.AagInfo = aagInfoModel
	model.HostInfo = hostInformationModel
	model.IsEncrypted = core.BoolPtr(true)
	model.SqlTargetParams = sqlTargetParamsForRecoverSqlAppModel
	model.TargetEnvironment = core.StringPtr("kSQL")

	result, err := backuprecovery.DataSourceIbmRecoveryRecoverSqlAppParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmRecoveryAAGInfoToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["name"] = "testString"
		model["object_id"] = int(26)

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.AAGInfo)
	model.Name = core.StringPtr("testString")
	model.ObjectID = core.Int64Ptr(int64(26))

	result, err := backuprecovery.DataSourceIbmRecoveryAAGInfoToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmRecoveryHostInformationToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["id"] = "testString"
		model["name"] = "testString"
		model["environment"] = "kPhysical"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.HostInformation)
	model.ID = core.StringPtr("testString")
	model.Name = core.StringPtr("testString")
	model.Environment = core.StringPtr("kPhysical")

	result, err := backuprecovery.DataSourceIbmRecoveryHostInformationToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmRecoverySqlTargetParamsForRecoverSqlAppToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		multiStageRestoreOptionsModel := make(map[string]interface{})
		multiStageRestoreOptionsModel["enable_auto_sync"] = true
		multiStageRestoreOptionsModel["enable_multi_stage_restore"] = true

		filenamePatternToDirectoryModel := make(map[string]interface{})
		filenamePatternToDirectoryModel["directory"] = "testString"
		filenamePatternToDirectoryModel["filename_pattern"] = "testString"

		recoveryObjectIdentifierModel := make(map[string]interface{})
		recoveryObjectIdentifierModel["id"] = int(26)

		recoverSqlAppNewSourceConfigModel := make(map[string]interface{})
		recoverSqlAppNewSourceConfigModel["keep_cdc"] = true
		recoverSqlAppNewSourceConfigModel["multi_stage_restore_options"] = []map[string]interface{}{multiStageRestoreOptionsModel}
		recoverSqlAppNewSourceConfigModel["native_log_recovery_with_clause"] = "testString"
		recoverSqlAppNewSourceConfigModel["native_recovery_with_clause"] = "testString"
		recoverSqlAppNewSourceConfigModel["overwriting_policy"] = "FailIfExists"
		recoverSqlAppNewSourceConfigModel["replay_entire_last_log"] = true
		recoverSqlAppNewSourceConfigModel["restore_time_usecs"] = int(26)
		recoverSqlAppNewSourceConfigModel["secondary_data_files_dir_list"] = []map[string]interface{}{filenamePatternToDirectoryModel}
		recoverSqlAppNewSourceConfigModel["with_no_recovery"] = true
		recoverSqlAppNewSourceConfigModel["data_file_directory_location"] = "testString"
		recoverSqlAppNewSourceConfigModel["database_name"] = "testString"
		recoverSqlAppNewSourceConfigModel["host"] = []map[string]interface{}{recoveryObjectIdentifierModel}
		recoverSqlAppNewSourceConfigModel["instance_name"] = "testString"
		recoverSqlAppNewSourceConfigModel["log_file_directory_location"] = "testString"

		recoverSqlAppOriginalSourceConfigModel := make(map[string]interface{})
		recoverSqlAppOriginalSourceConfigModel["keep_cdc"] = true
		recoverSqlAppOriginalSourceConfigModel["multi_stage_restore_options"] = []map[string]interface{}{multiStageRestoreOptionsModel}
		recoverSqlAppOriginalSourceConfigModel["native_log_recovery_with_clause"] = "testString"
		recoverSqlAppOriginalSourceConfigModel["native_recovery_with_clause"] = "testString"
		recoverSqlAppOriginalSourceConfigModel["overwriting_policy"] = "FailIfExists"
		recoverSqlAppOriginalSourceConfigModel["replay_entire_last_log"] = true
		recoverSqlAppOriginalSourceConfigModel["restore_time_usecs"] = int(26)
		recoverSqlAppOriginalSourceConfigModel["secondary_data_files_dir_list"] = []map[string]interface{}{filenamePatternToDirectoryModel}
		recoverSqlAppOriginalSourceConfigModel["with_no_recovery"] = true
		recoverSqlAppOriginalSourceConfigModel["capture_tail_logs"] = true
		recoverSqlAppOriginalSourceConfigModel["data_file_directory_location"] = "testString"
		recoverSqlAppOriginalSourceConfigModel["log_file_directory_location"] = "testString"
		recoverSqlAppOriginalSourceConfigModel["new_database_name"] = "testString"

		model := make(map[string]interface{})
		model["new_source_config"] = []map[string]interface{}{recoverSqlAppNewSourceConfigModel}
		model["original_source_config"] = []map[string]interface{}{recoverSqlAppOriginalSourceConfigModel}
		model["recover_to_new_source"] = true

		assert.Equal(t, result, model)
	}

	multiStageRestoreOptionsModel := new(backuprecoveryv1.MultiStageRestoreOptions)
	multiStageRestoreOptionsModel.EnableAutoSync = core.BoolPtr(true)
	multiStageRestoreOptionsModel.EnableMultiStageRestore = core.BoolPtr(true)

	filenamePatternToDirectoryModel := new(backuprecoveryv1.FilenamePatternToDirectory)
	filenamePatternToDirectoryModel.Directory = core.StringPtr("testString")
	filenamePatternToDirectoryModel.FilenamePattern = core.StringPtr("testString")

	recoveryObjectIdentifierModel := new(backuprecoveryv1.RecoveryObjectIdentifier)
	recoveryObjectIdentifierModel.ID = core.Int64Ptr(int64(26))

	recoverSqlAppNewSourceConfigModel := new(backuprecoveryv1.RecoverSqlAppNewSourceConfig)
	recoverSqlAppNewSourceConfigModel.KeepCdc = core.BoolPtr(true)
	recoverSqlAppNewSourceConfigModel.MultiStageRestoreOptions = multiStageRestoreOptionsModel
	recoverSqlAppNewSourceConfigModel.NativeLogRecoveryWithClause = core.StringPtr("testString")
	recoverSqlAppNewSourceConfigModel.NativeRecoveryWithClause = core.StringPtr("testString")
	recoverSqlAppNewSourceConfigModel.OverwritingPolicy = core.StringPtr("FailIfExists")
	recoverSqlAppNewSourceConfigModel.ReplayEntireLastLog = core.BoolPtr(true)
	recoverSqlAppNewSourceConfigModel.RestoreTimeUsecs = core.Int64Ptr(int64(26))
	recoverSqlAppNewSourceConfigModel.SecondaryDataFilesDirList = []backuprecoveryv1.FilenamePatternToDirectory{*filenamePatternToDirectoryModel}
	recoverSqlAppNewSourceConfigModel.WithNoRecovery = core.BoolPtr(true)
	recoverSqlAppNewSourceConfigModel.DataFileDirectoryLocation = core.StringPtr("testString")
	recoverSqlAppNewSourceConfigModel.DatabaseName = core.StringPtr("testString")
	recoverSqlAppNewSourceConfigModel.Host = recoveryObjectIdentifierModel
	recoverSqlAppNewSourceConfigModel.InstanceName = core.StringPtr("testString")
	recoverSqlAppNewSourceConfigModel.LogFileDirectoryLocation = core.StringPtr("testString")

	recoverSqlAppOriginalSourceConfigModel := new(backuprecoveryv1.RecoverSqlAppOriginalSourceConfig)
	recoverSqlAppOriginalSourceConfigModel.KeepCdc = core.BoolPtr(true)
	recoverSqlAppOriginalSourceConfigModel.MultiStageRestoreOptions = multiStageRestoreOptionsModel
	recoverSqlAppOriginalSourceConfigModel.NativeLogRecoveryWithClause = core.StringPtr("testString")
	recoverSqlAppOriginalSourceConfigModel.NativeRecoveryWithClause = core.StringPtr("testString")
	recoverSqlAppOriginalSourceConfigModel.OverwritingPolicy = core.StringPtr("FailIfExists")
	recoverSqlAppOriginalSourceConfigModel.ReplayEntireLastLog = core.BoolPtr(true)
	recoverSqlAppOriginalSourceConfigModel.RestoreTimeUsecs = core.Int64Ptr(int64(26))
	recoverSqlAppOriginalSourceConfigModel.SecondaryDataFilesDirList = []backuprecoveryv1.FilenamePatternToDirectory{*filenamePatternToDirectoryModel}
	recoverSqlAppOriginalSourceConfigModel.WithNoRecovery = core.BoolPtr(true)
	recoverSqlAppOriginalSourceConfigModel.CaptureTailLogs = core.BoolPtr(true)
	recoverSqlAppOriginalSourceConfigModel.DataFileDirectoryLocation = core.StringPtr("testString")
	recoverSqlAppOriginalSourceConfigModel.LogFileDirectoryLocation = core.StringPtr("testString")
	recoverSqlAppOriginalSourceConfigModel.NewDatabaseName = core.StringPtr("testString")

	model := new(backuprecoveryv1.SqlTargetParamsForRecoverSqlApp)
	model.NewSourceConfig = recoverSqlAppNewSourceConfigModel
	model.OriginalSourceConfig = recoverSqlAppOriginalSourceConfigModel
	model.RecoverToNewSource = core.BoolPtr(true)

	result, err := backuprecovery.DataSourceIbmRecoverySqlTargetParamsForRecoverSqlAppToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmRecoveryRecoverSqlAppNewSourceConfigToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		multiStageRestoreOptionsModel := make(map[string]interface{})
		multiStageRestoreOptionsModel["enable_auto_sync"] = true
		multiStageRestoreOptionsModel["enable_multi_stage_restore"] = true

		filenamePatternToDirectoryModel := make(map[string]interface{})
		filenamePatternToDirectoryModel["directory"] = "testString"
		filenamePatternToDirectoryModel["filename_pattern"] = "testString"

		recoveryObjectIdentifierModel := make(map[string]interface{})
		recoveryObjectIdentifierModel["id"] = int(26)

		model := make(map[string]interface{})
		model["keep_cdc"] = true
		model["multi_stage_restore_options"] = []map[string]interface{}{multiStageRestoreOptionsModel}
		model["native_log_recovery_with_clause"] = "testString"
		model["native_recovery_with_clause"] = "testString"
		model["overwriting_policy"] = "FailIfExists"
		model["replay_entire_last_log"] = true
		model["restore_time_usecs"] = int(26)
		model["secondary_data_files_dir_list"] = []map[string]interface{}{filenamePatternToDirectoryModel}
		model["with_no_recovery"] = true
		model["data_file_directory_location"] = "testString"
		model["database_name"] = "testString"
		model["host"] = []map[string]interface{}{recoveryObjectIdentifierModel}
		model["instance_name"] = "testString"
		model["log_file_directory_location"] = "testString"

		assert.Equal(t, result, model)
	}

	multiStageRestoreOptionsModel := new(backuprecoveryv1.MultiStageRestoreOptions)
	multiStageRestoreOptionsModel.EnableAutoSync = core.BoolPtr(true)
	multiStageRestoreOptionsModel.EnableMultiStageRestore = core.BoolPtr(true)

	filenamePatternToDirectoryModel := new(backuprecoveryv1.FilenamePatternToDirectory)
	filenamePatternToDirectoryModel.Directory = core.StringPtr("testString")
	filenamePatternToDirectoryModel.FilenamePattern = core.StringPtr("testString")

	recoveryObjectIdentifierModel := new(backuprecoveryv1.RecoveryObjectIdentifier)
	recoveryObjectIdentifierModel.ID = core.Int64Ptr(int64(26))

	model := new(backuprecoveryv1.RecoverSqlAppNewSourceConfig)
	model.KeepCdc = core.BoolPtr(true)
	model.MultiStageRestoreOptions = multiStageRestoreOptionsModel
	model.NativeLogRecoveryWithClause = core.StringPtr("testString")
	model.NativeRecoveryWithClause = core.StringPtr("testString")
	model.OverwritingPolicy = core.StringPtr("FailIfExists")
	model.ReplayEntireLastLog = core.BoolPtr(true)
	model.RestoreTimeUsecs = core.Int64Ptr(int64(26))
	model.SecondaryDataFilesDirList = []backuprecoveryv1.FilenamePatternToDirectory{*filenamePatternToDirectoryModel}
	model.WithNoRecovery = core.BoolPtr(true)
	model.DataFileDirectoryLocation = core.StringPtr("testString")
	model.DatabaseName = core.StringPtr("testString")
	model.Host = recoveryObjectIdentifierModel
	model.InstanceName = core.StringPtr("testString")
	model.LogFileDirectoryLocation = core.StringPtr("testString")

	result, err := backuprecovery.DataSourceIbmRecoveryRecoverSqlAppNewSourceConfigToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmRecoveryMultiStageRestoreOptionsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["enable_auto_sync"] = true
		model["enable_multi_stage_restore"] = true

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.MultiStageRestoreOptions)
	model.EnableAutoSync = core.BoolPtr(true)
	model.EnableMultiStageRestore = core.BoolPtr(true)

	result, err := backuprecovery.DataSourceIbmRecoveryMultiStageRestoreOptionsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmRecoveryFilenamePatternToDirectoryToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["directory"] = "testString"
		model["filename_pattern"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.FilenamePatternToDirectory)
	model.Directory = core.StringPtr("testString")
	model.FilenamePattern = core.StringPtr("testString")

	result, err := backuprecovery.DataSourceIbmRecoveryFilenamePatternToDirectoryToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmRecoveryRecoveryObjectIdentifierToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["id"] = int(26)
		model["name"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.RecoveryObjectIdentifier)
	model.ID = core.Int64Ptr(int64(26))
	model.Name = core.StringPtr("testString")

	result, err := backuprecovery.DataSourceIbmRecoveryRecoveryObjectIdentifierToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmRecoveryRecoverSqlAppOriginalSourceConfigToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		multiStageRestoreOptionsModel := make(map[string]interface{})
		multiStageRestoreOptionsModel["enable_auto_sync"] = true
		multiStageRestoreOptionsModel["enable_multi_stage_restore"] = true

		filenamePatternToDirectoryModel := make(map[string]interface{})
		filenamePatternToDirectoryModel["directory"] = "testString"
		filenamePatternToDirectoryModel["filename_pattern"] = "testString"

		model := make(map[string]interface{})
		model["keep_cdc"] = true
		model["multi_stage_restore_options"] = []map[string]interface{}{multiStageRestoreOptionsModel}
		model["native_log_recovery_with_clause"] = "testString"
		model["native_recovery_with_clause"] = "testString"
		model["overwriting_policy"] = "FailIfExists"
		model["replay_entire_last_log"] = true
		model["restore_time_usecs"] = int(26)
		model["secondary_data_files_dir_list"] = []map[string]interface{}{filenamePatternToDirectoryModel}
		model["with_no_recovery"] = true
		model["capture_tail_logs"] = true
		model["data_file_directory_location"] = "testString"
		model["log_file_directory_location"] = "testString"
		model["new_database_name"] = "testString"

		assert.Equal(t, result, model)
	}

	multiStageRestoreOptionsModel := new(backuprecoveryv1.MultiStageRestoreOptions)
	multiStageRestoreOptionsModel.EnableAutoSync = core.BoolPtr(true)
	multiStageRestoreOptionsModel.EnableMultiStageRestore = core.BoolPtr(true)

	filenamePatternToDirectoryModel := new(backuprecoveryv1.FilenamePatternToDirectory)
	filenamePatternToDirectoryModel.Directory = core.StringPtr("testString")
	filenamePatternToDirectoryModel.FilenamePattern = core.StringPtr("testString")

	model := new(backuprecoveryv1.RecoverSqlAppOriginalSourceConfig)
	model.KeepCdc = core.BoolPtr(true)
	model.MultiStageRestoreOptions = multiStageRestoreOptionsModel
	model.NativeLogRecoveryWithClause = core.StringPtr("testString")
	model.NativeRecoveryWithClause = core.StringPtr("testString")
	model.OverwritingPolicy = core.StringPtr("FailIfExists")
	model.ReplayEntireLastLog = core.BoolPtr(true)
	model.RestoreTimeUsecs = core.Int64Ptr(int64(26))
	model.SecondaryDataFilesDirList = []backuprecoveryv1.FilenamePatternToDirectory{*filenamePatternToDirectoryModel}
	model.WithNoRecovery = core.BoolPtr(true)
	model.CaptureTailLogs = core.BoolPtr(true)
	model.DataFileDirectoryLocation = core.StringPtr("testString")
	model.LogFileDirectoryLocation = core.StringPtr("testString")
	model.NewDatabaseName = core.StringPtr("testString")

	result, err := backuprecovery.DataSourceIbmRecoveryRecoverSqlAppOriginalSourceConfigToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmRecoveryRecoveryVlanConfigToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["id"] = int(38)
		model["disable_vlan"] = true
		model["interface_name"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.RecoveryVlanConfig)
	model.ID = core.Int64Ptr(int64(38))
	model.DisableVlan = core.BoolPtr(true)
	model.InterfaceName = core.StringPtr("testString")

	result, err := backuprecovery.DataSourceIbmRecoveryRecoveryVlanConfigToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
