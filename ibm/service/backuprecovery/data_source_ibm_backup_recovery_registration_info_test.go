// Copyright IBM Corp. 2025 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.105.1-067d600b-20250616-154447
 */

package backuprecovery_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/backuprecovery"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/ibm-backup-recovery-sdk-go/backuprecoveryv1"
	"github.com/stretchr/testify/assert"
)

func TestAccIbmBackupRecoveryRegistrationInfoDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmBackupRecoveryRegistrationInfoDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_registration_info.backup_recovery_registration_info_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_registration_info.backup_recovery_registration_info_instance", "x_ibm_tenant_id"),
				),
			},
		},
	})
}

func testAccCheckIbmBackupRecoveryRegistrationInfoDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_backup_recovery_registration_info" "backup_recovery_registration_info_instance" {
			X-IBM-Tenant-Id = "tenantId"
			environments = [ "kVMware" ]
			ids = [ 1 ]
			includeEntityPermissionInfo = true
			sids = [ "sids" ]
			includeSourceCredentials = true
			encryptionKey = "encryptionKey"
			includeApplicationsTreeInfo = true
			pruneNonCriticalInfo = true
			requestInitiatorType = "requestInitiatorType"
			useCachedData = true
			includeExternalMetadata = true
			maintenanceStatus = "UnderMaintenance"
			tenantIds = [ "tenantIds" ]
			allUnderHierarchy = true
		}
	`)
}

func TestDataSourceIbmBackupRecoveryRegistrationInfoProtectionSourceTreeInfoToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		cbtFileVersionModel := make(map[string]interface{})
		cbtFileVersionModel["build_ver"] = float64(72.5)
		cbtFileVersionModel["major_ver"] = float64(72.5)
		cbtFileVersionModel["minor_ver"] = float64(72.5)
		cbtFileVersionModel["revision_num"] = float64(72.5)

		cbtServiceStateModel := make(map[string]interface{})
		cbtServiceStateModel["name"] = "testString"
		cbtServiceStateModel["state"] = "testString"

		cbtInfoModel := make(map[string]interface{})
		cbtInfoModel["file_version"] = []map[string]interface{}{cbtFileVersionModel}
		cbtInfoModel["is_installed"] = true
		cbtInfoModel["reboot_status"] = "kRebooted"
		cbtInfoModel["service_state"] = []map[string]interface{}{cbtServiceStateModel}

		agentAccessInfoModel := make(map[string]interface{})
		agentAccessInfoModel["connection_id"] = int(26)
		agentAccessInfoModel["connector_group_id"] = int(26)
		agentAccessInfoModel["endpoint"] = "testString"
		agentAccessInfoModel["environment"] = "kPhysical"
		agentAccessInfoModel["id"] = int(26)
		agentAccessInfoModel["version"] = int(26)

		timeModel := make(map[string]interface{})
		timeModel["hour"] = int(38)
		timeModel["minute"] = int(38)

		dayTimeParamsModel := make(map[string]interface{})
		dayTimeParamsModel["day"] = "kSunday"
		dayTimeParamsModel["time"] = []map[string]interface{}{timeModel}

		dayTimeWindowModel := make(map[string]interface{})
		dayTimeWindowModel["end_time"] = []map[string]interface{}{dayTimeParamsModel}
		dayTimeWindowModel["start_time"] = []map[string]interface{}{dayTimeParamsModel}

		throttlingWindowModel := make(map[string]interface{})
		throttlingWindowModel["day_time_window"] = []map[string]interface{}{dayTimeWindowModel}
		throttlingWindowModel["threshold"] = int(26)

		throttlingConfigurationParamsModel := make(map[string]interface{})
		throttlingConfigurationParamsModel["fixed_threshold"] = int(26)
		throttlingConfigurationParamsModel["pattern_type"] = "kNoThrottling"
		throttlingConfigurationParamsModel["throttling_windows"] = []map[string]interface{}{throttlingWindowModel}

		throttlingConfigModel := make(map[string]interface{})
		throttlingConfigModel["cpu_throttling_config"] = []map[string]interface{}{throttlingConfigurationParamsModel}
		throttlingConfigModel["network_throttling_config"] = []map[string]interface{}{throttlingConfigurationParamsModel}

		agentPhysicalParamsModel := make(map[string]interface{})
		agentPhysicalParamsModel["applications"] = []string{"kSQL"}
		agentPhysicalParamsModel["password"] = "testString"
		agentPhysicalParamsModel["throttling_config"] = []map[string]interface{}{throttlingConfigModel}
		agentPhysicalParamsModel["username"] = "testString"

		hostSettingsCheckResultModel := make(map[string]interface{})
		hostSettingsCheckResultModel["check_type"] = "kIsAgentPortAccessible"
		hostSettingsCheckResultModel["result_type"] = "kPass"
		hostSettingsCheckResultModel["user_message"] = "testString"

		registeredAppInfoModel := make(map[string]interface{})
		registeredAppInfoModel["authentication_error_message"] = "testString"
		registeredAppInfoModel["authentication_status"] = "kPending"
		registeredAppInfoModel["environment"] = "kPhysical"
		registeredAppInfoModel["host_settings_check_results"] = []map[string]interface{}{hostSettingsCheckResultModel}
		registeredAppInfoModel["refresh_error_message"] = "testString"

		subnetModel := make(map[string]interface{})
		subnetModel["component"] = "testString"
		subnetModel["description"] = "testString"
		subnetModel["id"] = float64(72.5)
		subnetModel["ip"] = "testString"
		subnetModel["netmask_bits"] = float64(72.5)
		subnetModel["netmask_ip4"] = "testString"
		subnetModel["nfs_access"] = "kDisabled"
		subnetModel["nfs_all_squash"] = true
		subnetModel["nfs_root_squash"] = true
		subnetModel["s3_access"] = "kDisabled"
		subnetModel["smb_access"] = "kDisabled"
		subnetModel["tenant_id"] = "testString"

		latencyThresholdsModel := make(map[string]interface{})
		latencyThresholdsModel["active_task_msecs"] = int(26)
		latencyThresholdsModel["new_task_msecs"] = int(26)

		nasSourceParamsModel := make(map[string]interface{})
		nasSourceParamsModel["max_parallel_metadata_fetch_full_percentage"] = float64(72.5)
		nasSourceParamsModel["max_parallel_metadata_fetch_incremental_percentage"] = float64(72.5)
		nasSourceParamsModel["max_parallel_read_write_full_percentage"] = float64(72.5)
		nasSourceParamsModel["max_parallel_read_write_incremental_percentage"] = float64(72.5)

		storageArraySnapshotMaxSpaceConfigModel := make(map[string]interface{})
		storageArraySnapshotMaxSpaceConfigModel["max_snapshot_space_percentage"] = float64(72.5)

		maxSnapshotConfigModel := make(map[string]interface{})
		maxSnapshotConfigModel["max_snapshots"] = float64(72.5)

		maxSpaceConfigModel := make(map[string]interface{})
		maxSpaceConfigModel["max_snapshot_space_percentage"] = float64(72.5)

		storageArraySnapshotThrottlingPoliciesModel := make(map[string]interface{})
		storageArraySnapshotThrottlingPoliciesModel["id"] = int(26)
		storageArraySnapshotThrottlingPoliciesModel["is_max_snapshots_config_enabled"] = true
		storageArraySnapshotThrottlingPoliciesModel["is_max_space_config_enabled"] = true
		storageArraySnapshotThrottlingPoliciesModel["max_snapshot_config"] = []map[string]interface{}{maxSnapshotConfigModel}
		storageArraySnapshotThrottlingPoliciesModel["max_space_config"] = []map[string]interface{}{maxSpaceConfigModel}

		storageArraySnapshotConfigModel := make(map[string]interface{})
		storageArraySnapshotConfigModel["is_max_snapshots_config_enabled"] = true
		storageArraySnapshotConfigModel["is_max_space_config_enabled"] = true
		storageArraySnapshotConfigModel["storage_array_snapshot_max_space_config"] = []map[string]interface{}{storageArraySnapshotMaxSpaceConfigModel}
		storageArraySnapshotConfigModel["storage_array_snapshot_throttling_policies"] = []map[string]interface{}{storageArraySnapshotThrottlingPoliciesModel}

		throttlingPolicyModel := make(map[string]interface{})
		throttlingPolicyModel["enforce_max_streams"] = true
		throttlingPolicyModel["enforce_registered_source_max_backups"] = true
		throttlingPolicyModel["is_enabled"] = true
		throttlingPolicyModel["latency_thresholds"] = []map[string]interface{}{latencyThresholdsModel}
		throttlingPolicyModel["max_concurrent_streams"] = float64(72.5)
		throttlingPolicyModel["nas_source_params"] = []map[string]interface{}{nasSourceParamsModel}
		throttlingPolicyModel["registered_source_max_concurrent_backups"] = float64(72.5)
		throttlingPolicyModel["storage_array_snapshot_config"] = []map[string]interface{}{storageArraySnapshotConfigModel}

		throttlingPolicyOverridesModel := make(map[string]interface{})
		throttlingPolicyOverridesModel["datastore_id"] = int(26)
		throttlingPolicyOverridesModel["datastore_name"] = "testString"
		throttlingPolicyOverridesModel["throttling_policy"] = []map[string]interface{}{throttlingPolicyModel}

		registeredSourceVlanConfigModel := make(map[string]interface{})
		registeredSourceVlanConfigModel["vlan"] = float64(72.5)
		registeredSourceVlanConfigModel["disable_vlan"] = true
		registeredSourceVlanConfigModel["interface_name"] = "testString"

		agentRegistrationInfoModel := make(map[string]interface{})
		agentRegistrationInfoModel["access_info"] = []map[string]interface{}{agentAccessInfoModel}
		agentRegistrationInfoModel["allowed_ip_addresses"] = []string{"testString"}
		agentRegistrationInfoModel["authentication_error_message"] = "testString"
		agentRegistrationInfoModel["authentication_status"] = "kPending"
		agentRegistrationInfoModel["blacklisted_ip_addresses"] = []string{"testString"}
		agentRegistrationInfoModel["denied_ip_addresses"] = []string{"testString"}
		agentRegistrationInfoModel["environments"] = []string{"kPhysical"}
		agentRegistrationInfoModel["is_db_authenticated"] = true
		agentRegistrationInfoModel["is_storage_array_snapshot_enabled"] = true
		agentRegistrationInfoModel["link_vms_across_vcenter"] = true
		agentRegistrationInfoModel["minimum_free_space_gb"] = int(26)
		agentRegistrationInfoModel["minimum_free_space_percent"] = int(26)
		agentRegistrationInfoModel["password"] = "testString"
		agentRegistrationInfoModel["physical_params"] = []map[string]interface{}{agentPhysicalParamsModel}
		agentRegistrationInfoModel["progress_monitor_path"] = "testString"
		agentRegistrationInfoModel["refresh_error_message"] = "testString"
		agentRegistrationInfoModel["refresh_time_usecs"] = int(26)
		agentRegistrationInfoModel["registered_apps_info"] = []map[string]interface{}{registeredAppInfoModel}
		agentRegistrationInfoModel["registration_time_usecs"] = int(26)
		agentRegistrationInfoModel["subnets"] = []map[string]interface{}{subnetModel}
		agentRegistrationInfoModel["throttling_policy"] = []map[string]interface{}{throttlingPolicyModel}
		agentRegistrationInfoModel["throttling_policy_overrides"] = []map[string]interface{}{throttlingPolicyOverridesModel}
		agentRegistrationInfoModel["use_o_auth_for_exchange_online"] = true
		agentRegistrationInfoModel["use_vm_bios_uuid"] = true
		agentRegistrationInfoModel["user_messages"] = []string{"testString"}
		agentRegistrationInfoModel["username"] = "testString"
		agentRegistrationInfoModel["vlan_params"] = []map[string]interface{}{registeredSourceVlanConfigModel}
		agentRegistrationInfoModel["warning_messages"] = []string{"testString"}

		agentInformationModel := make(map[string]interface{})
		agentInformationModel["cbmr_version"] = "testString"
		agentInformationModel["file_cbt_info"] = []map[string]interface{}{cbtInfoModel}
		agentInformationModel["host_type"] = "kLinux"
		agentInformationModel["id"] = int(26)
		agentInformationModel["name"] = "testString"
		agentInformationModel["oracle_multi_node_channel_supported"] = true
		agentInformationModel["registration_info"] = []map[string]interface{}{agentRegistrationInfoModel}
		agentInformationModel["source_side_dedup_enabled"] = true
		agentInformationModel["status"] = "kUnknown"
		agentInformationModel["status_message"] = "testString"
		agentInformationModel["upgradability"] = "kUpgradable"
		agentInformationModel["upgrade_status"] = "kIdle"
		agentInformationModel["upgrade_status_message"] = "testString"
		agentInformationModel["version"] = "testString"
		agentInformationModel["vol_cbt_info"] = []map[string]interface{}{cbtInfoModel}

		uniqueGlobalIdModel := make(map[string]interface{})
		uniqueGlobalIdModel["cluster_id"] = int(26)
		uniqueGlobalIdModel["cluster_incarnation_id"] = int(26)
		uniqueGlobalIdModel["id"] = int(26)

		clusterNetworkingEndpointModel := make(map[string]interface{})
		clusterNetworkingEndpointModel["fqdn"] = "testString"
		clusterNetworkingEndpointModel["ipv4_addr"] = "testString"
		clusterNetworkingEndpointModel["ipv6_addr"] = "testString"

		clusterNetworkResourceInformationModel := make(map[string]interface{})
		clusterNetworkResourceInformationModel["endpoints"] = []map[string]interface{}{clusterNetworkingEndpointModel}
		clusterNetworkResourceInformationModel["type"] = "testString"

		networkingInformationModel := make(map[string]interface{})
		networkingInformationModel["resource_vec"] = []map[string]interface{}{clusterNetworkResourceInformationModel}

		physicalVolumeModel := make(map[string]interface{})
		physicalVolumeModel["device_path"] = "testString"
		physicalVolumeModel["guid"] = "testString"
		physicalVolumeModel["is_boot_volume"] = true
		physicalVolumeModel["is_extended_attributes_supported"] = true
		physicalVolumeModel["is_protected"] = true
		physicalVolumeModel["is_shared_volume"] = true
		physicalVolumeModel["label"] = "testString"
		physicalVolumeModel["logical_size_bytes"] = float64(72.5)
		physicalVolumeModel["mount_points"] = []string{"testString"}
		physicalVolumeModel["mount_type"] = "testString"
		physicalVolumeModel["network_path"] = "testString"
		physicalVolumeModel["used_size_bytes"] = float64(72.5)

		vssWritersModel := make(map[string]interface{})
		vssWritersModel["is_writer_excluded"] = true
		vssWritersModel["writer_name"] = true

		physicalProtectionSourceModel := make(map[string]interface{})
		physicalProtectionSourceModel["agents"] = []map[string]interface{}{agentInformationModel}
		physicalProtectionSourceModel["cluster_source_type"] = "testString"
		physicalProtectionSourceModel["host_name"] = "testString"
		physicalProtectionSourceModel["host_type"] = "kLinux"
		physicalProtectionSourceModel["id"] = []map[string]interface{}{uniqueGlobalIdModel}
		physicalProtectionSourceModel["is_proxy_host"] = true
		physicalProtectionSourceModel["memory_size_bytes"] = int(26)
		physicalProtectionSourceModel["name"] = "testString"
		physicalProtectionSourceModel["networking_info"] = []map[string]interface{}{networkingInformationModel}
		physicalProtectionSourceModel["num_processors"] = int(26)
		physicalProtectionSourceModel["os_name"] = "testString"
		physicalProtectionSourceModel["type"] = "kGroup"
		physicalProtectionSourceModel["vcs_version"] = "testString"
		physicalProtectionSourceModel["volumes"] = []map[string]interface{}{physicalVolumeModel}
		physicalProtectionSourceModel["vsswriters"] = []map[string]interface{}{vssWritersModel}

		vlanParametersModel := make(map[string]interface{})
		vlanParametersModel["disable_vlan"] = true
		vlanParametersModel["interface_name"] = "testString"
		vlanParametersModel["vlan"] = int(38)

		kubernetesLabelAttributeModel := make(map[string]interface{})
		kubernetesLabelAttributeModel["id"] = int(26)
		kubernetesLabelAttributeModel["name"] = "testString"
		kubernetesLabelAttributeModel["uuid"] = "testString"

		k8sLabelModel := make(map[string]interface{})
		k8sLabelModel["key"] = "testString"
		k8sLabelModel["value"] = "testString"

		serviceAnnotationsEntryModel := make(map[string]interface{})
		serviceAnnotationsEntryModel["key"] = "testString"
		serviceAnnotationsEntryModel["value"] = "testString"

		kubernetesStorageClassInfoModel := make(map[string]interface{})
		kubernetesStorageClassInfoModel["name"] = "testString"
		kubernetesStorageClassInfoModel["provisioner"] = "testString"

		kubernetesServiceAnnotationObjectModel := make(map[string]interface{})
		kubernetesServiceAnnotationObjectModel["key"] = "testString"
		kubernetesServiceAnnotationObjectModel["value"] = "testString"

		vlanParamsModel := make(map[string]interface{})
		vlanParamsModel["disable_vlan"] = true
		vlanParamsModel["interface_name"] = "testString"
		vlanParamsModel["vlan_id"] = int(38)

		kubernetesVlanInfoModel := make(map[string]interface{})
		kubernetesVlanInfoModel["service_annotations"] = []map[string]interface{}{kubernetesServiceAnnotationObjectModel}
		kubernetesVlanInfoModel["vlan_params"] = []map[string]interface{}{vlanParamsModel}

		kubernetesProtectionSourceModel := make(map[string]interface{})
		kubernetesProtectionSourceModel["datamover_image_location"] = "testString"
		kubernetesProtectionSourceModel["datamover_service_type"] = int(38)
		kubernetesProtectionSourceModel["datamover_upgradability"] = "testString"
		kubernetesProtectionSourceModel["default_vlan_params"] = []map[string]interface{}{vlanParametersModel}
		kubernetesProtectionSourceModel["description"] = "testString"
		kubernetesProtectionSourceModel["distribution"] = "kMainline"
		kubernetesProtectionSourceModel["init_container_image_location"] = "testString"
		kubernetesProtectionSourceModel["label_attributes"] = []map[string]interface{}{kubernetesLabelAttributeModel}
		kubernetesProtectionSourceModel["name"] = "testString"
		kubernetesProtectionSourceModel["priority_class_name"] = "testString"
		kubernetesProtectionSourceModel["resource_annotation_list"] = []map[string]interface{}{k8sLabelModel}
		kubernetesProtectionSourceModel["resource_label_list"] = []map[string]interface{}{k8sLabelModel}
		kubernetesProtectionSourceModel["san_field"] = []string{"testString"}
		kubernetesProtectionSourceModel["service_annotations"] = []map[string]interface{}{serviceAnnotationsEntryModel}
		kubernetesProtectionSourceModel["storage_class"] = []map[string]interface{}{kubernetesStorageClassInfoModel}
		kubernetesProtectionSourceModel["type"] = "kCluster"
		kubernetesProtectionSourceModel["uuid"] = "testString"
		kubernetesProtectionSourceModel["velero_aws_plugin_image_location"] = "testString"
		kubernetesProtectionSourceModel["velero_image_location"] = "testString"
		kubernetesProtectionSourceModel["velero_openshift_plugin_image_location"] = "testString"
		kubernetesProtectionSourceModel["velero_upgradability"] = "testString"
		kubernetesProtectionSourceModel["vlan_info_vec"] = []map[string]interface{}{kubernetesVlanInfoModel}

		databaseFileInformationModel := make(map[string]interface{})
		databaseFileInformationModel["file_type"] = "kRows"
		databaseFileInformationModel["full_path"] = "testString"
		databaseFileInformationModel["size_bytes"] = int(26)

		sqlSourceIdModel := make(map[string]interface{})
		sqlSourceIdModel["created_date_msecs"] = int(26)
		sqlSourceIdModel["database_id"] = int(26)
		sqlSourceIdModel["instance_id"] = "testString"

		sqlServerInstanceVersionModel := make(map[string]interface{})
		sqlServerInstanceVersionModel["build"] = float64(72.5)
		sqlServerInstanceVersionModel["major_version"] = float64(72.5)
		sqlServerInstanceVersionModel["minor_version"] = float64(72.5)
		sqlServerInstanceVersionModel["revision"] = float64(72.5)
		sqlServerInstanceVersionModel["version_string"] = float64(72.5)

		sqlProtectionSourceModel := make(map[string]interface{})
		sqlProtectionSourceModel["is_available_for_vss_backup"] = true
		sqlProtectionSourceModel["created_timestamp"] = "testString"
		sqlProtectionSourceModel["database_name"] = "testString"
		sqlProtectionSourceModel["db_aag_entity_id"] = int(26)
		sqlProtectionSourceModel["db_aag_name"] = "testString"
		sqlProtectionSourceModel["db_compatibility_level"] = int(26)
		sqlProtectionSourceModel["db_file_groups"] = []string{"testString"}
		sqlProtectionSourceModel["db_files"] = []map[string]interface{}{databaseFileInformationModel}
		sqlProtectionSourceModel["db_owner_username"] = "testString"
		sqlProtectionSourceModel["default_database_location"] = "testString"
		sqlProtectionSourceModel["default_log_location"] = "testString"
		sqlProtectionSourceModel["id"] = []map[string]interface{}{sqlSourceIdModel}
		sqlProtectionSourceModel["is_encrypted"] = true
		sqlProtectionSourceModel["name"] = "testString"
		sqlProtectionSourceModel["owner_id"] = int(26)
		sqlProtectionSourceModel["recovery_model"] = "kSimpleRecoveryModel"
		sqlProtectionSourceModel["sql_server_db_state"] = "kOnline"
		sqlProtectionSourceModel["sql_server_instance_version"] = []map[string]interface{}{sqlServerInstanceVersionModel}
		sqlProtectionSourceModel["type"] = "kInstance"

		protectionSourceNodeModel := make(map[string]interface{})
		protectionSourceNodeModel["connection_id"] = int(26)
		protectionSourceNodeModel["connector_group_id"] = int(26)
		protectionSourceNodeModel["custom_name"] = "testString"
		protectionSourceNodeModel["environment"] = "kPhysical"
		protectionSourceNodeModel["id"] = int(26)
		protectionSourceNodeModel["name"] = "testString"
		protectionSourceNodeModel["parent_id"] = int(26)
		protectionSourceNodeModel["physical_protection_source"] = []map[string]interface{}{physicalProtectionSourceModel}
		protectionSourceNodeModel["kubernetes_protection_source"] = []map[string]interface{}{kubernetesProtectionSourceModel}
		protectionSourceNodeModel["sql_protection_source"] = []map[string]interface{}{sqlProtectionSourceModel}

		applicationInfoModel := make(map[string]interface{})
		applicationInfoModel["application_tree_info"] = []map[string]interface{}{protectionSourceNodeModel}
		applicationInfoModel["environment"] = "kVMware"

		groupInfoModel := make(map[string]interface{})
		groupInfoModel["domain"] = "testString"
		groupInfoModel["group_name"] = "testString"
		groupInfoModel["sid"] = "testString"
		groupInfoModel["tenant_ids"] = []string{"testString"}

		tenantInfoModel := make(map[string]interface{})
		tenantInfoModel["bifrost_enabled"] = true
		tenantInfoModel["is_managed_on_helios"] = true
		tenantInfoModel["name"] = "testString"
		tenantInfoModel["tenant_id"] = "testString"

		userInfoModel := make(map[string]interface{})
		userInfoModel["domain"] = "testString"
		userInfoModel["sid"] = "testString"
		userInfoModel["tenant_id"] = "testString"
		userInfoModel["user_name"] = "testString"

		entityPermissionInformationModel := make(map[string]interface{})
		entityPermissionInformationModel["entity_id"] = int(26)
		entityPermissionInformationModel["groups"] = []map[string]interface{}{groupInfoModel}
		entityPermissionInformationModel["is_inferred"] = true
		entityPermissionInformationModel["is_registered_by_sp"] = true
		entityPermissionInformationModel["registering_tenant_id"] = "testString"
		entityPermissionInformationModel["tenant"] = []map[string]interface{}{tenantInfoModel}
		entityPermissionInformationModel["users"] = []map[string]interface{}{userInfoModel}

		timeRangeUsecsModel := make(map[string]interface{})
		timeRangeUsecsModel["end_time_usecs"] = int(26)
		timeRangeUsecsModel["start_time_usecs"] = int(26)

		timeWindowModel := make(map[string]interface{})
		timeWindowModel["day_of_the_week"] = "Sunday"
		timeWindowModel["end_time"] = []map[string]interface{}{timeModel}
		timeWindowModel["start_time"] = []map[string]interface{}{timeModel}

		scheduleModel := make(map[string]interface{})
		scheduleModel["periodic_time_windows"] = []map[string]interface{}{timeWindowModel}
		scheduleModel["schedule_type"] = "PeriodicTimeWindows"
		scheduleModel["time_ranges"] = []map[string]interface{}{timeRangeUsecsModel}
		scheduleModel["timezone"] = "testString"

		workflowInterventionSpecModel := make(map[string]interface{})
		workflowInterventionSpecModel["intervention"] = "NoIntervention"
		workflowInterventionSpecModel["workflow_type"] = "BackupRun"

		maintenanceModeConfigModel := make(map[string]interface{})
		maintenanceModeConfigModel["activation_time_intervals"] = []map[string]interface{}{timeRangeUsecsModel}
		maintenanceModeConfigModel["maintenance_schedule"] = []map[string]interface{}{scheduleModel}
		maintenanceModeConfigModel["user_message"] = "testString"
		maintenanceModeConfigModel["workflow_intervention_spec_list"] = []map[string]interface{}{workflowInterventionSpecModel}

		connectorParametersModel := make(map[string]interface{})
		connectorParametersModel["connection_id"] = int(26)
		connectorParametersModel["connector_group_id"] = int(26)
		connectorParametersModel["endpoint"] = "testString"
		connectorParametersModel["environment"] = "kVMware"
		connectorParametersModel["id"] = int(26)
		connectorParametersModel["version"] = int(26)

		cassandraPortsInfoModel := make(map[string]interface{})
		cassandraPortsInfoModel["jmx_port"] = int(38)
		cassandraPortsInfoModel["native_transport_port"] = int(38)
		cassandraPortsInfoModel["rpc_port"] = int(38)
		cassandraPortsInfoModel["ssl_storage_port"] = int(38)
		cassandraPortsInfoModel["storage_port"] = int(38)

		cassandraSecurityInfoModel := make(map[string]interface{})
		cassandraSecurityInfoModel["cassandra_auth_required"] = true
		cassandraSecurityInfoModel["cassandra_auth_type"] = "PASSWORD"
		cassandraSecurityInfoModel["cassandra_authorizer"] = "testString"
		cassandraSecurityInfoModel["client_encryption"] = true
		cassandraSecurityInfoModel["dse_authorization"] = true
		cassandraSecurityInfoModel["server_encryption_req_client_auth"] = true
		cassandraSecurityInfoModel["server_internode_encryption_type"] = "testString"

		cassandraConnectParamsModel := make(map[string]interface{})
		cassandraConnectParamsModel["cassandra_ports_info"] = []map[string]interface{}{cassandraPortsInfoModel}
		cassandraConnectParamsModel["cassandra_security_info"] = []map[string]interface{}{cassandraSecurityInfoModel}
		cassandraConnectParamsModel["cassandra_version"] = "testString"
		cassandraConnectParamsModel["commit_log_backup_location"] = "testString"
		cassandraConnectParamsModel["config_directory"] = "testString"
		cassandraConnectParamsModel["data_centers"] = []string{"testString"}
		cassandraConnectParamsModel["dse_config_directory"] = "testString"
		cassandraConnectParamsModel["dse_version"] = "testString"
		cassandraConnectParamsModel["is_dse_authenticator"] = true
		cassandraConnectParamsModel["is_dse_tiered_storage"] = true
		cassandraConnectParamsModel["is_jmx_auth_enable"] = true
		cassandraConnectParamsModel["kerberos_principal"] = "testString"
		cassandraConnectParamsModel["primary_host"] = "testString"
		cassandraConnectParamsModel["seeds"] = []string{"testString"}
		cassandraConnectParamsModel["solr_nodes"] = []string{"testString"}
		cassandraConnectParamsModel["solr_port"] = int(38)

		couchbaseConnectParamsModel := make(map[string]interface{})
		couchbaseConnectParamsModel["carrier_direct_port"] = int(38)
		couchbaseConnectParamsModel["http_direct_port"] = int(38)
		couchbaseConnectParamsModel["requires_ssl"] = true
		couchbaseConnectParamsModel["seeds"] = []string{"testString"}

		hadoopDiscoveryParamsModel := make(map[string]interface{})
		hadoopDiscoveryParamsModel["config_directory"] = "testString"
		hadoopDiscoveryParamsModel["host"] = "testString"

		hBaseConnectParamsModel := make(map[string]interface{})
		hBaseConnectParamsModel["hbase_discovery_params"] = []map[string]interface{}{hadoopDiscoveryParamsModel}
		hBaseConnectParamsModel["hdfs_entity_id"] = int(26)
		hBaseConnectParamsModel["kerberos_principal"] = "testString"
		hBaseConnectParamsModel["root_data_directory"] = "testString"
		hBaseConnectParamsModel["zookeeper_quorum"] = []string{"testString"}

		hdfsConnectParamsModel := make(map[string]interface{})
		hdfsConnectParamsModel["hadoop_distribution"] = "CDH"
		hdfsConnectParamsModel["hadoop_version"] = "testString"
		hdfsConnectParamsModel["hdfs_connection_type"] = "DFS"
		hdfsConnectParamsModel["hdfs_discovery_params"] = []map[string]interface{}{hadoopDiscoveryParamsModel}
		hdfsConnectParamsModel["kerberos_principal"] = "testString"
		hdfsConnectParamsModel["namenode"] = "testString"
		hdfsConnectParamsModel["port"] = int(38)

		hiveConnectParamsModel := make(map[string]interface{})
		hiveConnectParamsModel["entity_threshold_exceeded"] = true
		hiveConnectParamsModel["hdfs_entity_id"] = int(26)
		hiveConnectParamsModel["hive_discovery_params"] = []map[string]interface{}{hadoopDiscoveryParamsModel}
		hiveConnectParamsModel["kerberos_principal"] = "testString"
		hiveConnectParamsModel["metastore"] = "testString"
		hiveConnectParamsModel["thrift_port"] = int(38)

		networkPoolConfigModel := make(map[string]interface{})
		networkPoolConfigModel["pool_name"] = "testString"
		networkPoolConfigModel["subnet"] = "testString"
		networkPoolConfigModel["use_smart_connect"] = true

		zoneConfigModel := make(map[string]interface{})
		zoneConfigModel["dynamic_network_pool_config"] = []map[string]interface{}{networkPoolConfigModel}

		registeredProtectionSourceIsilonParamsModel := make(map[string]interface{})
		registeredProtectionSourceIsilonParamsModel["zone_config_list"] = []map[string]interface{}{zoneConfigModel}

		mongoDbConnectParamsModel := make(map[string]interface{})
		mongoDbConnectParamsModel["auth_type"] = "SCRAM"
		mongoDbConnectParamsModel["authenticating_database_name"] = "testString"
		mongoDbConnectParamsModel["requires_ssl"] = true
		mongoDbConnectParamsModel["secondary_node_tag"] = "testString"
		mongoDbConnectParamsModel["seeds"] = []string{"testString"}
		mongoDbConnectParamsModel["use_fixed_node_for_backup"] = true
		mongoDbConnectParamsModel["use_secondary_for_backup"] = true

		nasServerCredentialsModel := make(map[string]interface{})
		nasServerCredentialsModel["domain"] = "testString"
		nasServerCredentialsModel["nas_protocol"] = "kNoProtocol"

		sitesDiscoveryParamsModel := make(map[string]interface{})
		sitesDiscoveryParamsModel["enable_site_tagging"] = true

		teamsAdditionalParamsModel := make(map[string]interface{})
		teamsAdditionalParamsModel["allow_posts_backup"] = true

		usersDiscoveryParamsModel := make(map[string]interface{})
		usersDiscoveryParamsModel["allow_chats_backup"] = true
		usersDiscoveryParamsModel["discover_users_with_mailbox"] = true
		usersDiscoveryParamsModel["discover_users_with_onedrive"] = true
		usersDiscoveryParamsModel["fetch_mailbox_info"] = true
		usersDiscoveryParamsModel["fetch_one_drive_info"] = true
		usersDiscoveryParamsModel["skip_users_without_my_site"] = true

		objectsDiscoveryParamsModel := make(map[string]interface{})
		objectsDiscoveryParamsModel["discoverable_object_type_list"] = []string{"testString"}
		objectsDiscoveryParamsModel["sites_discovery_params"] = []map[string]interface{}{sitesDiscoveryParamsModel}
		objectsDiscoveryParamsModel["teams_additional_params"] = []map[string]interface{}{teamsAdditionalParamsModel}
		objectsDiscoveryParamsModel["users_discovery_params"] = []map[string]interface{}{usersDiscoveryParamsModel}

		m365CsmParamsModel := make(map[string]interface{})
		m365CsmParamsModel["backup_allowed"] = true

		o365ConnectParamsModel := make(map[string]interface{})
		o365ConnectParamsModel["objects_discovery_params"] = []map[string]interface{}{objectsDiscoveryParamsModel}
		o365ConnectParamsModel["csm_params"] = []map[string]interface{}{m365CsmParamsModel}

		office365CredentialsModel := make(map[string]interface{})
		office365CredentialsModel["client_id"] = "testString"
		office365CredentialsModel["client_secret"] = "testString"
		office365CredentialsModel["grant_type"] = "testString"
		office365CredentialsModel["scope"] = "testString"
		office365CredentialsModel["use_o_auth_for_exchange_online"] = true

		credentialsModel := make(map[string]interface{})
		credentialsModel["username"] = "testString"
		credentialsModel["password"] = "testString"

		throttlingConfigurationModel := make(map[string]interface{})
		throttlingConfigurationModel["fixed_threshold"] = int(26)
		throttlingConfigurationModel["pattern_type"] = "kNoThrottling"
		throttlingConfigurationModel["throttling_windows"] = []map[string]interface{}{throttlingWindowModel}

		sourceThrottlingConfigurationModel := make(map[string]interface{})
		sourceThrottlingConfigurationModel["cpu_throttling_config"] = []map[string]interface{}{throttlingConfigurationModel}
		sourceThrottlingConfigurationModel["network_throttling_config"] = []map[string]interface{}{throttlingConfigurationModel}

		physicalParamsModel := make(map[string]interface{})
		physicalParamsModel["applications"] = []string{"kVMware"}
		physicalParamsModel["password"] = "testString"
		physicalParamsModel["throttling_config"] = []map[string]interface{}{sourceThrottlingConfigurationModel}
		physicalParamsModel["username"] = "testString"

		sfdcParamsModel := make(map[string]interface{})
		sfdcParamsModel["access_token"] = "testString"
		sfdcParamsModel["concurrent_api_requests_limit"] = int(26)
		sfdcParamsModel["consumer_key"] = "testString"
		sfdcParamsModel["consumer_secret"] = "testString"
		sfdcParamsModel["daily_api_limit"] = int(26)
		sfdcParamsModel["endpoint"] = "testString"
		sfdcParamsModel["endpoint_type"] = "PROD"
		sfdcParamsModel["metadata_endpoint_url"] = "testString"
		sfdcParamsModel["refresh_token"] = "testString"
		sfdcParamsModel["soap_endpoint_url"] = "testString"
		sfdcParamsModel["use_bulk_api"] = true

		nasSourceThrottlingParamsModel := make(map[string]interface{})
		nasSourceThrottlingParamsModel["max_parallel_metadata_fetch_full_percentage"] = int(38)
		nasSourceThrottlingParamsModel["max_parallel_metadata_fetch_incremental_percentage"] = int(38)
		nasSourceThrottlingParamsModel["max_parallel_read_write_full_percentage"] = int(38)
		nasSourceThrottlingParamsModel["max_parallel_read_write_incremental_percentage"] = int(38)

		storageArraySnapshotMaxSpaceConfigParamsModel := make(map[string]interface{})
		storageArraySnapshotMaxSpaceConfigParamsModel["max_snapshot_space_percentage"] = int(38)

		storageArraySnapshotMaxSnapshotConfigParamsModel := make(map[string]interface{})
		storageArraySnapshotMaxSnapshotConfigParamsModel["max_snapshots"] = int(38)

		storageArraySnapshotThrottlingPolicyModel := make(map[string]interface{})
		storageArraySnapshotThrottlingPolicyModel["id"] = int(26)
		storageArraySnapshotThrottlingPolicyModel["is_max_snapshots_config_enabled"] = true
		storageArraySnapshotThrottlingPolicyModel["is_max_space_config_enabled"] = true
		storageArraySnapshotThrottlingPolicyModel["max_snapshot_config"] = []map[string]interface{}{storageArraySnapshotMaxSnapshotConfigParamsModel}
		storageArraySnapshotThrottlingPolicyModel["max_space_config"] = []map[string]interface{}{storageArraySnapshotMaxSpaceConfigParamsModel}

		storageArraySnapshotConfigParamsModel := make(map[string]interface{})
		storageArraySnapshotConfigParamsModel["is_max_snapshots_config_enabled"] = true
		storageArraySnapshotConfigParamsModel["is_max_space_config_enabled"] = true
		storageArraySnapshotConfigParamsModel["storage_array_snapshot_max_space_config"] = []map[string]interface{}{storageArraySnapshotMaxSpaceConfigParamsModel}
		storageArraySnapshotConfigParamsModel["storage_array_snapshot_throttling_policies"] = []map[string]interface{}{storageArraySnapshotThrottlingPolicyModel}

		throttlingPolicyParametersModel := make(map[string]interface{})
		throttlingPolicyParametersModel["enforce_max_streams"] = true
		throttlingPolicyParametersModel["enforce_registered_source_max_backups"] = true
		throttlingPolicyParametersModel["is_enabled"] = true
		throttlingPolicyParametersModel["latency_thresholds"] = []map[string]interface{}{latencyThresholdsModel}
		throttlingPolicyParametersModel["max_concurrent_streams"] = int(38)
		throttlingPolicyParametersModel["nas_source_params"] = []map[string]interface{}{nasSourceThrottlingParamsModel}
		throttlingPolicyParametersModel["registered_source_max_concurrent_backups"] = int(38)
		throttlingPolicyParametersModel["storage_array_snapshot_config"] = []map[string]interface{}{storageArraySnapshotConfigParamsModel}

		throttlingPolicyOverrideModel := make(map[string]interface{})
		throttlingPolicyOverrideModel["datastore_id"] = int(26)
		throttlingPolicyOverrideModel["datastore_name"] = "testString"
		throttlingPolicyOverrideModel["throttling_policy"] = []map[string]interface{}{throttlingPolicyParametersModel}

		udaSourceCapabilitiesModel := make(map[string]interface{})
		udaSourceCapabilitiesModel["auto_log_backup"] = true
		udaSourceCapabilitiesModel["dynamic_config"] = true
		udaSourceCapabilitiesModel["entity_support"] = true
		udaSourceCapabilitiesModel["et_log_backup"] = true
		udaSourceCapabilitiesModel["external_disks"] = true
		udaSourceCapabilitiesModel["full_backup"] = true
		udaSourceCapabilitiesModel["incr_backup"] = true
		udaSourceCapabilitiesModel["log_backup"] = true
		udaSourceCapabilitiesModel["multi_object_restore"] = true
		udaSourceCapabilitiesModel["pause_resume_backup"] = true
		udaSourceCapabilitiesModel["post_backup_job_script"] = true
		udaSourceCapabilitiesModel["post_restore_job_script"] = true
		udaSourceCapabilitiesModel["pre_backup_job_script"] = true
		udaSourceCapabilitiesModel["pre_restore_job_script"] = true
		udaSourceCapabilitiesModel["resource_throttling"] = true
		udaSourceCapabilitiesModel["snapfs_cert"] = true

		keyValueStrPairModel := make(map[string]interface{})
		keyValueStrPairModel["key"] = "testString"
		keyValueStrPairModel["value"] = "testString"

		udaConnectParamsModel := make(map[string]interface{})
		udaConnectParamsModel["capabilities"] = []map[string]interface{}{udaSourceCapabilitiesModel}
		udaConnectParamsModel["credentials"] = []map[string]interface{}{credentialsModel}
		udaConnectParamsModel["et_enable_log_backup_policy"] = true
		udaConnectParamsModel["et_enable_run_now"] = true
		udaConnectParamsModel["host_type"] = "kLinux"
		udaConnectParamsModel["hosts"] = []string{"testString"}
		udaConnectParamsModel["live_data_view"] = true
		udaConnectParamsModel["live_log_view"] = true
		udaConnectParamsModel["mount_dir"] = "testString"
		udaConnectParamsModel["mount_view"] = true
		udaConnectParamsModel["script_dir"] = "testString"
		udaConnectParamsModel["source_args"] = "testString"
		udaConnectParamsModel["source_registration_arguments"] = []map[string]interface{}{keyValueStrPairModel}
		udaConnectParamsModel["source_type"] = "testString"

		registeredSourceInfoModel := make(map[string]interface{})
		registeredSourceInfoModel["access_info"] = []map[string]interface{}{connectorParametersModel}
		registeredSourceInfoModel["allowed_ip_addresses"] = []string{"testString"}
		registeredSourceInfoModel["authentication_error_message"] = "testString"
		registeredSourceInfoModel["authentication_status"] = "kPending"
		registeredSourceInfoModel["blacklisted_ip_addresses"] = []string{"testString"}
		registeredSourceInfoModel["cassandra_params"] = []map[string]interface{}{cassandraConnectParamsModel}
		registeredSourceInfoModel["couchbase_params"] = []map[string]interface{}{couchbaseConnectParamsModel}
		registeredSourceInfoModel["denied_ip_addresses"] = []string{"testString"}
		registeredSourceInfoModel["environments"] = []string{"kVMware"}
		registeredSourceInfoModel["hbase_params"] = []map[string]interface{}{hBaseConnectParamsModel}
		registeredSourceInfoModel["hdfs_params"] = []map[string]interface{}{hdfsConnectParamsModel}
		registeredSourceInfoModel["hive_params"] = []map[string]interface{}{hiveConnectParamsModel}
		registeredSourceInfoModel["is_db_authenticated"] = true
		registeredSourceInfoModel["is_storage_array_snapshot_enabled"] = true
		registeredSourceInfoModel["isilon_params"] = []map[string]interface{}{registeredProtectionSourceIsilonParamsModel}
		registeredSourceInfoModel["link_vms_across_vcenter"] = true
		registeredSourceInfoModel["minimum_free_space_gb"] = int(26)
		registeredSourceInfoModel["minimum_free_space_percent"] = int(26)
		registeredSourceInfoModel["mongodb_params"] = []map[string]interface{}{mongoDbConnectParamsModel}
		registeredSourceInfoModel["nas_mount_credentials"] = []map[string]interface{}{nasServerCredentialsModel}
		registeredSourceInfoModel["o365_params"] = []map[string]interface{}{o365ConnectParamsModel}
		registeredSourceInfoModel["office365_credentials_list"] = []map[string]interface{}{office365CredentialsModel}
		registeredSourceInfoModel["office365_region"] = "testString"
		registeredSourceInfoModel["office365_service_account_credentials_list"] = []map[string]interface{}{credentialsModel}
		registeredSourceInfoModel["password"] = "testString"
		registeredSourceInfoModel["physical_params"] = []map[string]interface{}{physicalParamsModel}
		registeredSourceInfoModel["progress_monitor_path"] = "testString"
		registeredSourceInfoModel["refresh_error_message"] = "testString"
		registeredSourceInfoModel["refresh_time_usecs"] = int(26)
		registeredSourceInfoModel["registered_apps_info"] = []map[string]interface{}{registeredAppInfoModel}
		registeredSourceInfoModel["registration_time_usecs"] = int(26)
		registeredSourceInfoModel["sfdc_params"] = []map[string]interface{}{sfdcParamsModel}
		registeredSourceInfoModel["subnets"] = []map[string]interface{}{subnetModel}
		registeredSourceInfoModel["throttling_policy"] = []map[string]interface{}{throttlingPolicyParametersModel}
		registeredSourceInfoModel["throttling_policy_overrides"] = []map[string]interface{}{throttlingPolicyOverrideModel}
		registeredSourceInfoModel["uda_params"] = []map[string]interface{}{udaConnectParamsModel}
		registeredSourceInfoModel["update_last_backup_details"] = true
		registeredSourceInfoModel["use_o_auth_for_exchange_online"] = true
		registeredSourceInfoModel["use_vm_bios_uuid"] = true
		registeredSourceInfoModel["user_messages"] = []string{"testString"}
		registeredSourceInfoModel["username"] = "testString"
		registeredSourceInfoModel["vlan_params"] = []map[string]interface{}{vlanParametersModel}
		registeredSourceInfoModel["warning_messages"] = []string{"testString"}

		protectionSourceTreeInfoStatsModel := make(map[string]interface{})
		protectionSourceTreeInfoStatsModel["protected_count"] = int(26)
		protectionSourceTreeInfoStatsModel["protected_size"] = int(26)
		protectionSourceTreeInfoStatsModel["unprotected_count"] = int(26)
		protectionSourceTreeInfoStatsModel["unprotected_size"] = int(26)

		protectionSummaryForK8sDistributionsModel := make(map[string]interface{})
		protectionSummaryForK8sDistributionsModel["distribution"] = "kMainline"
		protectionSummaryForK8sDistributionsModel["protected_count"] = int(26)
		protectionSummaryForK8sDistributionsModel["protected_size"] = int(26)
		protectionSummaryForK8sDistributionsModel["total_registered_clusters"] = int(26)
		protectionSummaryForK8sDistributionsModel["unprotected_count"] = int(26)
		protectionSummaryForK8sDistributionsModel["unprotected_size"] = int(26)

		protectionSummaryByEnvModel := make(map[string]interface{})
		protectionSummaryByEnvModel["environment"] = "kVMware"
		protectionSummaryByEnvModel["kubernetes_distribution_stats"] = []map[string]interface{}{protectionSummaryForK8sDistributionsModel}
		protectionSummaryByEnvModel["protected_count"] = int(26)
		protectionSummaryByEnvModel["protected_size"] = int(26)
		protectionSummaryByEnvModel["unprotected_count"] = int(26)
		protectionSummaryByEnvModel["unprotected_size"] = int(26)

		model := make(map[string]interface{})
		model["applications"] = []map[string]interface{}{applicationInfoModel}
		model["entity_permission_info"] = []map[string]interface{}{entityPermissionInformationModel}
		model["logical_size_bytes"] = int(26)
		model["maintenance_mode_config"] = []map[string]interface{}{maintenanceModeConfigModel}
		model["registration_info"] = []map[string]interface{}{registeredSourceInfoModel}
		model["root_node"] = []map[string]interface{}{protectionSourceNodeModel}
		model["stats"] = []map[string]interface{}{protectionSourceTreeInfoStatsModel}
		model["stats_by_env"] = []map[string]interface{}{protectionSummaryByEnvModel}
		model["total_downtiered_size_in_bytes"] = int(26)
		model["total_uptiered_size_in_bytes"] = int(26)

		assert.Equal(t, result, model)
	}

	cbtFileVersionModel := new(backuprecoveryv1.CbtFileVersion)
	cbtFileVersionModel.BuildVer = core.Float64Ptr(float64(72.5))
	cbtFileVersionModel.MajorVer = core.Float64Ptr(float64(72.5))
	cbtFileVersionModel.MinorVer = core.Float64Ptr(float64(72.5))
	cbtFileVersionModel.RevisionNum = core.Float64Ptr(float64(72.5))

	cbtServiceStateModel := new(backuprecoveryv1.CbtServiceState)
	cbtServiceStateModel.Name = core.StringPtr("testString")
	cbtServiceStateModel.State = core.StringPtr("testString")

	cbtInfoModel := new(backuprecoveryv1.CbtInfo)
	cbtInfoModel.FileVersion = cbtFileVersionModel
	cbtInfoModel.IsInstalled = core.BoolPtr(true)
	cbtInfoModel.RebootStatus = core.StringPtr("kRebooted")
	cbtInfoModel.ServiceState = cbtServiceStateModel

	agentAccessInfoModel := new(backuprecoveryv1.AgentAccessInfo)
	agentAccessInfoModel.ConnectionID = core.Int64Ptr(int64(26))
	agentAccessInfoModel.ConnectorGroupID = core.Int64Ptr(int64(26))
	agentAccessInfoModel.Endpoint = core.StringPtr("testString")
	agentAccessInfoModel.Environment = core.StringPtr("kPhysical")
	agentAccessInfoModel.ID = core.Int64Ptr(int64(26))
	agentAccessInfoModel.Version = core.Int64Ptr(int64(26))

	timeModel := new(backuprecoveryv1.Time)
	timeModel.Hour = core.Int64Ptr(int64(38))
	timeModel.Minute = core.Int64Ptr(int64(38))

	dayTimeParamsModel := new(backuprecoveryv1.DayTimeParams)
	dayTimeParamsModel.Day = core.StringPtr("kSunday")
	dayTimeParamsModel.Time = timeModel

	dayTimeWindowModel := new(backuprecoveryv1.DayTimeWindow)
	dayTimeWindowModel.EndTime = dayTimeParamsModel
	dayTimeWindowModel.StartTime = dayTimeParamsModel

	throttlingWindowModel := new(backuprecoveryv1.ThrottlingWindow)
	throttlingWindowModel.DayTimeWindow = dayTimeWindowModel
	throttlingWindowModel.Threshold = core.Int64Ptr(int64(26))

	throttlingConfigurationParamsModel := new(backuprecoveryv1.ThrottlingConfigurationParams)
	throttlingConfigurationParamsModel.FixedThreshold = core.Int64Ptr(int64(26))
	throttlingConfigurationParamsModel.PatternType = core.StringPtr("kNoThrottling")
	throttlingConfigurationParamsModel.ThrottlingWindows = []backuprecoveryv1.ThrottlingWindow{*throttlingWindowModel}

	throttlingConfigModel := new(backuprecoveryv1.ThrottlingConfig)
	throttlingConfigModel.CpuThrottlingConfig = throttlingConfigurationParamsModel
	throttlingConfigModel.NetworkThrottlingConfig = throttlingConfigurationParamsModel

	agentPhysicalParamsModel := new(backuprecoveryv1.AgentPhysicalParams)
	agentPhysicalParamsModel.Applications = []string{"kSQL"}
	agentPhysicalParamsModel.Password = core.StringPtr("testString")
	agentPhysicalParamsModel.ThrottlingConfig = throttlingConfigModel
	agentPhysicalParamsModel.Username = core.StringPtr("testString")

	hostSettingsCheckResultModel := new(backuprecoveryv1.HostSettingsCheckResult)
	hostSettingsCheckResultModel.CheckType = core.StringPtr("kIsAgentPortAccessible")
	hostSettingsCheckResultModel.ResultType = core.StringPtr("kPass")
	hostSettingsCheckResultModel.UserMessage = core.StringPtr("testString")

	registeredAppInfoModel := new(backuprecoveryv1.RegisteredAppInfo)
	registeredAppInfoModel.AuthenticationErrorMessage = core.StringPtr("testString")
	registeredAppInfoModel.AuthenticationStatus = core.StringPtr("kPending")
	registeredAppInfoModel.Environment = core.StringPtr("kPhysical")
	registeredAppInfoModel.HostSettingsCheckResults = []backuprecoveryv1.HostSettingsCheckResult{*hostSettingsCheckResultModel}
	registeredAppInfoModel.RefreshErrorMessage = core.StringPtr("testString")

	subnetModel := new(backuprecoveryv1.Subnet)
	subnetModel.Component = core.StringPtr("testString")
	subnetModel.Description = core.StringPtr("testString")
	subnetModel.ID = core.Float64Ptr(float64(72.5))
	subnetModel.Ip = core.StringPtr("testString")
	subnetModel.NetmaskBits = core.Float64Ptr(float64(72.5))
	subnetModel.NetmaskIp4 = core.StringPtr("testString")
	subnetModel.NfsAccess = core.StringPtr("kDisabled")
	subnetModel.NfsAllSquash = core.BoolPtr(true)
	subnetModel.NfsRootSquash = core.BoolPtr(true)
	subnetModel.S3Access = core.StringPtr("kDisabled")
	subnetModel.SmbAccess = core.StringPtr("kDisabled")
	subnetModel.TenantID = core.StringPtr("testString")

	latencyThresholdsModel := new(backuprecoveryv1.LatencyThresholds)
	latencyThresholdsModel.ActiveTaskMsecs = core.Int64Ptr(int64(26))
	latencyThresholdsModel.NewTaskMsecs = core.Int64Ptr(int64(26))

	nasSourceParamsModel := new(backuprecoveryv1.NasSourceParams)
	nasSourceParamsModel.MaxParallelMetadataFetchFullPercentage = core.Float64Ptr(float64(72.5))
	nasSourceParamsModel.MaxParallelMetadataFetchIncrementalPercentage = core.Float64Ptr(float64(72.5))
	nasSourceParamsModel.MaxParallelReadWriteFullPercentage = core.Float64Ptr(float64(72.5))
	nasSourceParamsModel.MaxParallelReadWriteIncrementalPercentage = core.Float64Ptr(float64(72.5))

	storageArraySnapshotMaxSpaceConfigModel := new(backuprecoveryv1.StorageArraySnapshotMaxSpaceConfig)
	storageArraySnapshotMaxSpaceConfigModel.MaxSnapshotSpacePercentage = core.Float64Ptr(float64(72.5))

	maxSnapshotConfigModel := new(backuprecoveryv1.MaxSnapshotConfig)
	maxSnapshotConfigModel.MaxSnapshots = core.Float64Ptr(float64(72.5))

	maxSpaceConfigModel := new(backuprecoveryv1.MaxSpaceConfig)
	maxSpaceConfigModel.MaxSnapshotSpacePercentage = core.Float64Ptr(float64(72.5))

	storageArraySnapshotThrottlingPoliciesModel := new(backuprecoveryv1.StorageArraySnapshotThrottlingPolicies)
	storageArraySnapshotThrottlingPoliciesModel.ID = core.Int64Ptr(int64(26))
	storageArraySnapshotThrottlingPoliciesModel.IsMaxSnapshotsConfigEnabled = core.BoolPtr(true)
	storageArraySnapshotThrottlingPoliciesModel.IsMaxSpaceConfigEnabled = core.BoolPtr(true)
	storageArraySnapshotThrottlingPoliciesModel.MaxSnapshotConfig = maxSnapshotConfigModel
	storageArraySnapshotThrottlingPoliciesModel.MaxSpaceConfig = maxSpaceConfigModel

	storageArraySnapshotConfigModel := new(backuprecoveryv1.StorageArraySnapshotConfig)
	storageArraySnapshotConfigModel.IsMaxSnapshotsConfigEnabled = core.BoolPtr(true)
	storageArraySnapshotConfigModel.IsMaxSpaceConfigEnabled = core.BoolPtr(true)
	storageArraySnapshotConfigModel.StorageArraySnapshotMaxSpaceConfig = storageArraySnapshotMaxSpaceConfigModel
	storageArraySnapshotConfigModel.StorageArraySnapshotThrottlingPolicies = []backuprecoveryv1.StorageArraySnapshotThrottlingPolicies{*storageArraySnapshotThrottlingPoliciesModel}

	throttlingPolicyModel := new(backuprecoveryv1.ThrottlingPolicy)
	throttlingPolicyModel.EnforceMaxStreams = core.BoolPtr(true)
	throttlingPolicyModel.EnforceRegisteredSourceMaxBackups = core.BoolPtr(true)
	throttlingPolicyModel.IsEnabled = core.BoolPtr(true)
	throttlingPolicyModel.LatencyThresholds = latencyThresholdsModel
	throttlingPolicyModel.MaxConcurrentStreams = core.Float64Ptr(float64(72.5))
	throttlingPolicyModel.NasSourceParams = nasSourceParamsModel
	throttlingPolicyModel.RegisteredSourceMaxConcurrentBackups = core.Float64Ptr(float64(72.5))
	throttlingPolicyModel.StorageArraySnapshotConfig = storageArraySnapshotConfigModel

	throttlingPolicyOverridesModel := new(backuprecoveryv1.ThrottlingPolicyOverrides)
	throttlingPolicyOverridesModel.DatastoreID = core.Int64Ptr(int64(26))
	throttlingPolicyOverridesModel.DatastoreName = core.StringPtr("testString")
	throttlingPolicyOverridesModel.ThrottlingPolicy = throttlingPolicyModel

	registeredSourceVlanConfigModel := new(backuprecoveryv1.RegisteredSourceVlanConfig)
	registeredSourceVlanConfigModel.Vlan = core.Float64Ptr(float64(72.5))
	registeredSourceVlanConfigModel.DisableVlan = core.BoolPtr(true)
	registeredSourceVlanConfigModel.InterfaceName = core.StringPtr("testString")

	agentRegistrationInfoModel := new(backuprecoveryv1.AgentRegistrationInfo)
	agentRegistrationInfoModel.AccessInfo = agentAccessInfoModel
	agentRegistrationInfoModel.AllowedIpAddresses = []string{"testString"}
	agentRegistrationInfoModel.AuthenticationErrorMessage = core.StringPtr("testString")
	agentRegistrationInfoModel.AuthenticationStatus = core.StringPtr("kPending")
	agentRegistrationInfoModel.BlacklistedIpAddresses = []string{"testString"}
	agentRegistrationInfoModel.DeniedIpAddresses = []string{"testString"}
	agentRegistrationInfoModel.Environments = []string{"kPhysical"}
	agentRegistrationInfoModel.IsDbAuthenticated = core.BoolPtr(true)
	agentRegistrationInfoModel.IsStorageArraySnapshotEnabled = core.BoolPtr(true)
	agentRegistrationInfoModel.LinkVmsAcrossVcenter = core.BoolPtr(true)
	agentRegistrationInfoModel.MinimumFreeSpaceGB = core.Int64Ptr(int64(26))
	agentRegistrationInfoModel.MinimumFreeSpacePercent = core.Int64Ptr(int64(26))
	agentRegistrationInfoModel.Password = core.StringPtr("testString")
	agentRegistrationInfoModel.PhysicalParams = agentPhysicalParamsModel
	agentRegistrationInfoModel.ProgressMonitorPath = core.StringPtr("testString")
	agentRegistrationInfoModel.RefreshErrorMessage = core.StringPtr("testString")
	agentRegistrationInfoModel.RefreshTimeUsecs = core.Int64Ptr(int64(26))
	agentRegistrationInfoModel.RegisteredAppsInfo = []backuprecoveryv1.RegisteredAppInfo{*registeredAppInfoModel}
	agentRegistrationInfoModel.RegistrationTimeUsecs = core.Int64Ptr(int64(26))
	agentRegistrationInfoModel.Subnets = []backuprecoveryv1.Subnet{*subnetModel}
	agentRegistrationInfoModel.ThrottlingPolicy = throttlingPolicyModel
	agentRegistrationInfoModel.ThrottlingPolicyOverrides = []backuprecoveryv1.ThrottlingPolicyOverrides{*throttlingPolicyOverridesModel}
	agentRegistrationInfoModel.UseOAuthForExchangeOnline = core.BoolPtr(true)
	agentRegistrationInfoModel.UseVmBiosUUID = core.BoolPtr(true)
	agentRegistrationInfoModel.UserMessages = []string{"testString"}
	agentRegistrationInfoModel.Username = core.StringPtr("testString")
	agentRegistrationInfoModel.VlanParams = registeredSourceVlanConfigModel
	agentRegistrationInfoModel.WarningMessages = []string{"testString"}

	agentInformationModel := new(backuprecoveryv1.AgentInformation)
	agentInformationModel.CbmrVersion = core.StringPtr("testString")
	agentInformationModel.FileCbtInfo = cbtInfoModel
	agentInformationModel.HostType = core.StringPtr("kLinux")
	agentInformationModel.ID = core.Int64Ptr(int64(26))
	agentInformationModel.Name = core.StringPtr("testString")
	agentInformationModel.OracleMultiNodeChannelSupported = core.BoolPtr(true)
	agentInformationModel.RegistrationInfo = agentRegistrationInfoModel
	agentInformationModel.SourceSideDedupEnabled = core.BoolPtr(true)
	agentInformationModel.Status = core.StringPtr("kUnknown")
	agentInformationModel.StatusMessage = core.StringPtr("testString")
	agentInformationModel.Upgradability = core.StringPtr("kUpgradable")
	agentInformationModel.UpgradeStatus = core.StringPtr("kIdle")
	agentInformationModel.UpgradeStatusMessage = core.StringPtr("testString")
	agentInformationModel.Version = core.StringPtr("testString")
	agentInformationModel.VolCbtInfo = cbtInfoModel

	uniqueGlobalIdModel := new(backuprecoveryv1.UniqueGlobalID)
	uniqueGlobalIdModel.ClusterID = core.Int64Ptr(int64(26))
	uniqueGlobalIdModel.ClusterIncarnationID = core.Int64Ptr(int64(26))
	uniqueGlobalIdModel.ID = core.Int64Ptr(int64(26))

	clusterNetworkingEndpointModel := new(backuprecoveryv1.ClusterNetworkingEndpoint)
	clusterNetworkingEndpointModel.Fqdn = core.StringPtr("testString")
	clusterNetworkingEndpointModel.Ipv4Addr = core.StringPtr("testString")
	clusterNetworkingEndpointModel.Ipv6Addr = core.StringPtr("testString")

	clusterNetworkResourceInformationModel := new(backuprecoveryv1.ClusterNetworkResourceInformation)
	clusterNetworkResourceInformationModel.Endpoints = []backuprecoveryv1.ClusterNetworkingEndpoint{*clusterNetworkingEndpointModel}
	clusterNetworkResourceInformationModel.Type = core.StringPtr("testString")

	networkingInformationModel := new(backuprecoveryv1.NetworkingInformation)
	networkingInformationModel.ResourceVec = []backuprecoveryv1.ClusterNetworkResourceInformation{*clusterNetworkResourceInformationModel}

	physicalVolumeModel := new(backuprecoveryv1.PhysicalVolume)
	physicalVolumeModel.DevicePath = core.StringPtr("testString")
	physicalVolumeModel.Guid = core.StringPtr("testString")
	physicalVolumeModel.IsBootVolume = core.BoolPtr(true)
	physicalVolumeModel.IsExtendedAttributesSupported = core.BoolPtr(true)
	physicalVolumeModel.IsProtected = core.BoolPtr(true)
	physicalVolumeModel.IsSharedVolume = core.BoolPtr(true)
	physicalVolumeModel.Label = core.StringPtr("testString")
	physicalVolumeModel.LogicalSizeBytes = core.Float64Ptr(float64(72.5))
	physicalVolumeModel.MountPoints = []string{"testString"}
	physicalVolumeModel.MountType = core.StringPtr("testString")
	physicalVolumeModel.NetworkPath = core.StringPtr("testString")
	physicalVolumeModel.UsedSizeBytes = core.Float64Ptr(float64(72.5))

	vssWritersModel := new(backuprecoveryv1.VssWriters)
	vssWritersModel.IsWriterExcluded = core.BoolPtr(true)
	vssWritersModel.WriterName = core.BoolPtr(true)

	physicalProtectionSourceModel := new(backuprecoveryv1.PhysicalProtectionSource)
	physicalProtectionSourceModel.Agents = []backuprecoveryv1.AgentInformation{*agentInformationModel}
	physicalProtectionSourceModel.ClusterSourceType = core.StringPtr("testString")
	physicalProtectionSourceModel.HostName = core.StringPtr("testString")
	physicalProtectionSourceModel.HostType = core.StringPtr("kLinux")
	physicalProtectionSourceModel.ID = uniqueGlobalIdModel
	physicalProtectionSourceModel.IsProxyHost = core.BoolPtr(true)
	physicalProtectionSourceModel.MemorySizeBytes = core.Int64Ptr(int64(26))
	physicalProtectionSourceModel.Name = core.StringPtr("testString")
	physicalProtectionSourceModel.NetworkingInfo = networkingInformationModel
	physicalProtectionSourceModel.NumProcessors = core.Int64Ptr(int64(26))
	physicalProtectionSourceModel.OsName = core.StringPtr("testString")
	physicalProtectionSourceModel.Type = core.StringPtr("kGroup")
	physicalProtectionSourceModel.VcsVersion = core.StringPtr("testString")
	physicalProtectionSourceModel.Volumes = []backuprecoveryv1.PhysicalVolume{*physicalVolumeModel}
	physicalProtectionSourceModel.Vsswriters = []backuprecoveryv1.VssWriters{*vssWritersModel}

	vlanParametersModel := new(backuprecoveryv1.VlanParameters)
	vlanParametersModel.DisableVlan = core.BoolPtr(true)
	vlanParametersModel.InterfaceName = core.StringPtr("testString")
	vlanParametersModel.Vlan = core.Int64Ptr(int64(38))

	kubernetesLabelAttributeModel := new(backuprecoveryv1.KubernetesLabelAttribute)
	kubernetesLabelAttributeModel.ID = core.Int64Ptr(int64(26))
	kubernetesLabelAttributeModel.Name = core.StringPtr("testString")
	kubernetesLabelAttributeModel.UUID = core.StringPtr("testString")

	k8sLabelModel := new(backuprecoveryv1.K8sLabel)
	k8sLabelModel.Key = core.StringPtr("testString")
	k8sLabelModel.Value = core.StringPtr("testString")

	serviceAnnotationsEntryModel := new(backuprecoveryv1.ServiceAnnotationsEntry)
	serviceAnnotationsEntryModel.Key = core.StringPtr("testString")
	serviceAnnotationsEntryModel.Value = core.StringPtr("testString")

	kubernetesStorageClassInfoModel := new(backuprecoveryv1.KubernetesStorageClassInfo)
	kubernetesStorageClassInfoModel.Name = core.StringPtr("testString")
	kubernetesStorageClassInfoModel.Provisioner = core.StringPtr("testString")

	kubernetesServiceAnnotationObjectModel := new(backuprecoveryv1.KubernetesServiceAnnotationObject)
	kubernetesServiceAnnotationObjectModel.Key = core.StringPtr("testString")
	kubernetesServiceAnnotationObjectModel.Value = core.StringPtr("testString")

	vlanParamsModel := new(backuprecoveryv1.VlanParams)
	vlanParamsModel.DisableVlan = core.BoolPtr(true)
	vlanParamsModel.InterfaceName = core.StringPtr("testString")
	vlanParamsModel.VlanID = core.Int64Ptr(int64(38))

	kubernetesVlanInfoModel := new(backuprecoveryv1.KubernetesVlanInfo)
	kubernetesVlanInfoModel.ServiceAnnotations = []backuprecoveryv1.KubernetesServiceAnnotationObject{*kubernetesServiceAnnotationObjectModel}
	kubernetesVlanInfoModel.VlanParams = vlanParamsModel

	kubernetesProtectionSourceModel := new(backuprecoveryv1.KubernetesProtectionSource)
	kubernetesProtectionSourceModel.DatamoverImageLocation = core.StringPtr("testString")
	kubernetesProtectionSourceModel.DatamoverServiceType = core.Int64Ptr(int64(38))
	kubernetesProtectionSourceModel.DatamoverUpgradability = core.StringPtr("kCurrent")
	kubernetesProtectionSourceModel.DefaultVlanParams = vlanParametersModel
	kubernetesProtectionSourceModel.Description = core.StringPtr("testString")
	kubernetesProtectionSourceModel.Distribution = core.StringPtr("kMainline")
	kubernetesProtectionSourceModel.InitContainerImageLocation = core.StringPtr("testString")
	kubernetesProtectionSourceModel.LabelAttributes = []backuprecoveryv1.KubernetesLabelAttribute{*kubernetesLabelAttributeModel}
	kubernetesProtectionSourceModel.Name = core.StringPtr("testString")
	kubernetesProtectionSourceModel.PriorityClassName = core.StringPtr("testString")
	kubernetesProtectionSourceModel.ResourceAnnotationList = []backuprecoveryv1.K8sLabel{*k8sLabelModel}
	kubernetesProtectionSourceModel.ResourceLabelList = []backuprecoveryv1.K8sLabel{*k8sLabelModel}
	kubernetesProtectionSourceModel.SanField = []string{"testString"}
	kubernetesProtectionSourceModel.ServiceAnnotations = []backuprecoveryv1.ServiceAnnotationsEntry{*serviceAnnotationsEntryModel}
	kubernetesProtectionSourceModel.StorageClass = []backuprecoveryv1.KubernetesStorageClassInfo{*kubernetesStorageClassInfoModel}
	kubernetesProtectionSourceModel.Type = core.StringPtr("kCluster")
	kubernetesProtectionSourceModel.UUID = core.StringPtr("testString")
	kubernetesProtectionSourceModel.VeleroAwsPluginImageLocation = core.StringPtr("testString")
	kubernetesProtectionSourceModel.VeleroImageLocation = core.StringPtr("testString")
	kubernetesProtectionSourceModel.VeleroOpenshiftPluginImageLocation = core.StringPtr("testString")
	kubernetesProtectionSourceModel.VeleroUpgradability = core.StringPtr("testString")
	kubernetesProtectionSourceModel.VlanInfoVec = []backuprecoveryv1.KubernetesVlanInfo{*kubernetesVlanInfoModel}

	databaseFileInformationModel := new(backuprecoveryv1.DatabaseFileInformation)
	databaseFileInformationModel.FileType = core.StringPtr("kRows")
	databaseFileInformationModel.FullPath = core.StringPtr("testString")
	databaseFileInformationModel.SizeBytes = core.Int64Ptr(int64(26))

	sqlSourceIdModel := new(backuprecoveryv1.SQLSourceID)
	sqlSourceIdModel.CreatedDateMsecs = core.Int64Ptr(int64(26))
	sqlSourceIdModel.DatabaseID = core.Int64Ptr(int64(26))
	sqlSourceIdModel.InstanceID = core.StringPtr("testString")

	sqlServerInstanceVersionModel := new(backuprecoveryv1.SQLServerInstanceVersion)
	sqlServerInstanceVersionModel.Build = core.Float64Ptr(float64(72.5))
	sqlServerInstanceVersionModel.MajorVersion = core.Float64Ptr(float64(72.5))
	sqlServerInstanceVersionModel.MinorVersion = core.Float64Ptr(float64(72.5))
	sqlServerInstanceVersionModel.Revision = core.Float64Ptr(float64(72.5))
	sqlServerInstanceVersionModel.VersionString = core.Float64Ptr(float64(72.5))

	sqlProtectionSourceModel := new(backuprecoveryv1.SqlProtectionSource)
	sqlProtectionSourceModel.IsAvailableForVssBackup = core.BoolPtr(true)
	sqlProtectionSourceModel.CreatedTimestamp = core.StringPtr("testString")
	sqlProtectionSourceModel.DatabaseName = core.StringPtr("testString")
	sqlProtectionSourceModel.DbAagEntityID = core.Int64Ptr(int64(26))
	sqlProtectionSourceModel.DbAagName = core.StringPtr("testString")
	sqlProtectionSourceModel.DbCompatibilityLevel = core.Int64Ptr(int64(26))
	sqlProtectionSourceModel.DbFileGroups = []string{"testString"}
	sqlProtectionSourceModel.DbFiles = []backuprecoveryv1.DatabaseFileInformation{*databaseFileInformationModel}
	sqlProtectionSourceModel.DbOwnerUsername = core.StringPtr("testString")
	sqlProtectionSourceModel.DefaultDatabaseLocation = core.StringPtr("testString")
	sqlProtectionSourceModel.DefaultLogLocation = core.StringPtr("testString")
	sqlProtectionSourceModel.ID = sqlSourceIdModel
	sqlProtectionSourceModel.IsEncrypted = core.BoolPtr(true)
	sqlProtectionSourceModel.Name = core.StringPtr("testString")
	sqlProtectionSourceModel.OwnerID = core.Int64Ptr(int64(26))
	sqlProtectionSourceModel.RecoveryModel = core.StringPtr("kSimpleRecoveryModel")
	sqlProtectionSourceModel.SqlServerDbState = core.StringPtr("kOnline")
	sqlProtectionSourceModel.SqlServerInstanceVersion = sqlServerInstanceVersionModel
	sqlProtectionSourceModel.Type = core.StringPtr("kInstance")

	protectionSourceNodeModel := new(backuprecoveryv1.ProtectionSourceNode)
	protectionSourceNodeModel.ConnectionID = core.Int64Ptr(int64(26))
	protectionSourceNodeModel.ConnectorGroupID = core.Int64Ptr(int64(26))
	protectionSourceNodeModel.CustomName = core.StringPtr("testString")
	protectionSourceNodeModel.Environment = core.StringPtr("kPhysical")
	protectionSourceNodeModel.ID = core.Int64Ptr(int64(26))
	protectionSourceNodeModel.Name = core.StringPtr("testString")
	protectionSourceNodeModel.ParentID = core.Int64Ptr(int64(26))
	protectionSourceNodeModel.PhysicalProtectionSource = physicalProtectionSourceModel
	protectionSourceNodeModel.KubernetesProtectionSource = kubernetesProtectionSourceModel
	protectionSourceNodeModel.SqlProtectionSource = sqlProtectionSourceModel

	applicationInfoModel := new(backuprecoveryv1.ApplicationInfo)
	applicationInfoModel.ApplicationTreeInfo = []backuprecoveryv1.ProtectionSourceNode{*protectionSourceNodeModel}
	applicationInfoModel.Environment = core.StringPtr("kVMware")

	groupInfoModel := new(backuprecoveryv1.GroupInfo)
	groupInfoModel.Domain = core.StringPtr("testString")
	groupInfoModel.GroupName = core.StringPtr("testString")
	groupInfoModel.Sid = core.StringPtr("testString")
	groupInfoModel.TenantIds = []string{"testString"}

	tenantInfoModel := new(backuprecoveryv1.TenantInfo)
	tenantInfoModel.BifrostEnabled = core.BoolPtr(true)
	tenantInfoModel.IsManagedOnHelios = core.BoolPtr(true)
	tenantInfoModel.Name = core.StringPtr("testString")
	tenantInfoModel.TenantID = core.StringPtr("testString")

	userInfoModel := new(backuprecoveryv1.UserInfo)
	userInfoModel.Domain = core.StringPtr("testString")
	userInfoModel.Sid = core.StringPtr("testString")
	userInfoModel.TenantID = core.StringPtr("testString")
	userInfoModel.UserName = core.StringPtr("testString")

	entityPermissionInformationModel := new(backuprecoveryv1.EntityPermissionInformation)
	entityPermissionInformationModel.EntityID = core.Int64Ptr(int64(26))
	entityPermissionInformationModel.Groups = []backuprecoveryv1.GroupInfo{*groupInfoModel}
	entityPermissionInformationModel.IsInferred = core.BoolPtr(true)
	entityPermissionInformationModel.IsRegisteredBySp = core.BoolPtr(true)
	entityPermissionInformationModel.RegisteringTenantID = core.StringPtr("testString")
	entityPermissionInformationModel.Tenant = tenantInfoModel
	entityPermissionInformationModel.Users = []backuprecoveryv1.UserInfo{*userInfoModel}

	timeRangeUsecsModel := new(backuprecoveryv1.TimeRangeUsecs)
	timeRangeUsecsModel.EndTimeUsecs = core.Int64Ptr(int64(26))
	timeRangeUsecsModel.StartTimeUsecs = core.Int64Ptr(int64(26))

	timeWindowModel := new(backuprecoveryv1.TimeWindow)
	timeWindowModel.DayOfTheWeek = core.StringPtr("Sunday")
	timeWindowModel.EndTime = timeModel
	timeWindowModel.StartTime = timeModel

	scheduleModel := new(backuprecoveryv1.Schedule)
	scheduleModel.PeriodicTimeWindows = []backuprecoveryv1.TimeWindow{*timeWindowModel}
	scheduleModel.ScheduleType = core.StringPtr("PeriodicTimeWindows")
	scheduleModel.TimeRanges = []backuprecoveryv1.TimeRangeUsecs{*timeRangeUsecsModel}
	scheduleModel.Timezone = core.StringPtr("testString")

	workflowInterventionSpecModel := new(backuprecoveryv1.WorkflowInterventionSpec)
	workflowInterventionSpecModel.Intervention = core.StringPtr("NoIntervention")
	workflowInterventionSpecModel.WorkflowType = core.StringPtr("BackupRun")

	maintenanceModeConfigModel := new(backuprecoveryv1.MaintenanceModeConfig)
	maintenanceModeConfigModel.ActivationTimeIntervals = []backuprecoveryv1.TimeRangeUsecs{*timeRangeUsecsModel}
	maintenanceModeConfigModel.MaintenanceSchedule = scheduleModel
	maintenanceModeConfigModel.UserMessage = core.StringPtr("testString")
	maintenanceModeConfigModel.WorkflowInterventionSpecList = []backuprecoveryv1.WorkflowInterventionSpec{*workflowInterventionSpecModel}

	connectorParametersModel := new(backuprecoveryv1.ConnectorParameters)
	connectorParametersModel.ConnectionID = core.Int64Ptr(int64(26))
	connectorParametersModel.ConnectorGroupID = core.Int64Ptr(int64(26))
	connectorParametersModel.Endpoint = core.StringPtr("testString")
	connectorParametersModel.Environment = core.StringPtr("kVMware")
	connectorParametersModel.ID = core.Int64Ptr(int64(26))
	connectorParametersModel.Version = core.Int64Ptr(int64(26))

	cassandraPortsInfoModel := new(backuprecoveryv1.CassandraPortsInfo)
	cassandraPortsInfoModel.JmxPort = core.Int64Ptr(int64(38))
	cassandraPortsInfoModel.NativeTransportPort = core.Int64Ptr(int64(38))
	cassandraPortsInfoModel.RpcPort = core.Int64Ptr(int64(38))
	cassandraPortsInfoModel.SslStoragePort = core.Int64Ptr(int64(38))
	cassandraPortsInfoModel.StoragePort = core.Int64Ptr(int64(38))

	cassandraSecurityInfoModel := new(backuprecoveryv1.CassandraSecurityInfo)
	cassandraSecurityInfoModel.CassandraAuthRequired = core.BoolPtr(true)
	cassandraSecurityInfoModel.CassandraAuthType = core.StringPtr("PASSWORD")
	cassandraSecurityInfoModel.CassandraAuthorizer = core.StringPtr("testString")
	cassandraSecurityInfoModel.ClientEncryption = core.BoolPtr(true)
	cassandraSecurityInfoModel.DseAuthorization = core.BoolPtr(true)
	cassandraSecurityInfoModel.ServerEncryptionReqClientAuth = core.BoolPtr(true)
	cassandraSecurityInfoModel.ServerInternodeEncryptionType = core.StringPtr("testString")

	cassandraConnectParamsModel := new(backuprecoveryv1.CassandraConnectParams)
	cassandraConnectParamsModel.CassandraPortsInfo = cassandraPortsInfoModel
	cassandraConnectParamsModel.CassandraSecurityInfo = cassandraSecurityInfoModel
	cassandraConnectParamsModel.CassandraVersion = core.StringPtr("testString")
	cassandraConnectParamsModel.CommitLogBackupLocation = core.StringPtr("testString")
	cassandraConnectParamsModel.ConfigDirectory = core.StringPtr("testString")
	cassandraConnectParamsModel.DataCenters = []string{"testString"}
	cassandraConnectParamsModel.DseConfigDirectory = core.StringPtr("testString")
	cassandraConnectParamsModel.DseVersion = core.StringPtr("testString")
	cassandraConnectParamsModel.IsDseAuthenticator = core.BoolPtr(true)
	cassandraConnectParamsModel.IsDseTieredStorage = core.BoolPtr(true)
	cassandraConnectParamsModel.IsJmxAuthEnable = core.BoolPtr(true)
	cassandraConnectParamsModel.KerberosPrincipal = core.StringPtr("testString")
	cassandraConnectParamsModel.PrimaryHost = core.StringPtr("testString")
	cassandraConnectParamsModel.Seeds = []string{"testString"}
	cassandraConnectParamsModel.SolrNodes = []string{"testString"}
	cassandraConnectParamsModel.SolrPort = core.Int64Ptr(int64(38))

	couchbaseConnectParamsModel := new(backuprecoveryv1.CouchbaseConnectParams)
	couchbaseConnectParamsModel.CarrierDirectPort = core.Int64Ptr(int64(38))
	couchbaseConnectParamsModel.HttpDirectPort = core.Int64Ptr(int64(38))
	couchbaseConnectParamsModel.RequiresSsl = core.BoolPtr(true)
	couchbaseConnectParamsModel.Seeds = []string{"testString"}

	hadoopDiscoveryParamsModel := new(backuprecoveryv1.HadoopDiscoveryParams)
	hadoopDiscoveryParamsModel.ConfigDirectory = core.StringPtr("testString")
	hadoopDiscoveryParamsModel.Host = core.StringPtr("testString")

	hBaseConnectParamsModel := new(backuprecoveryv1.HBaseConnectParams)
	hBaseConnectParamsModel.HbaseDiscoveryParams = hadoopDiscoveryParamsModel
	hBaseConnectParamsModel.HdfsEntityID = core.Int64Ptr(int64(26))
	hBaseConnectParamsModel.KerberosPrincipal = core.StringPtr("testString")
	hBaseConnectParamsModel.RootDataDirectory = core.StringPtr("testString")
	hBaseConnectParamsModel.ZookeeperQuorum = []string{"testString"}

	hdfsConnectParamsModel := new(backuprecoveryv1.HdfsConnectParams)
	hdfsConnectParamsModel.HadoopDistribution = core.StringPtr("CDH")
	hdfsConnectParamsModel.HadoopVersion = core.StringPtr("testString")
	hdfsConnectParamsModel.HdfsConnectionType = core.StringPtr("DFS")
	hdfsConnectParamsModel.HdfsDiscoveryParams = hadoopDiscoveryParamsModel
	hdfsConnectParamsModel.KerberosPrincipal = core.StringPtr("testString")
	hdfsConnectParamsModel.Namenode = core.StringPtr("testString")
	hdfsConnectParamsModel.Port = core.Int64Ptr(int64(38))

	hiveConnectParamsModel := new(backuprecoveryv1.HiveConnectParams)
	hiveConnectParamsModel.EntityThresholdExceeded = core.BoolPtr(true)
	hiveConnectParamsModel.HdfsEntityID = core.Int64Ptr(int64(26))
	hiveConnectParamsModel.HiveDiscoveryParams = hadoopDiscoveryParamsModel
	hiveConnectParamsModel.KerberosPrincipal = core.StringPtr("testString")
	hiveConnectParamsModel.Metastore = core.StringPtr("testString")
	hiveConnectParamsModel.ThriftPort = core.Int64Ptr(int64(38))

	networkPoolConfigModel := new(backuprecoveryv1.NetworkPoolConfig)
	networkPoolConfigModel.PoolName = core.StringPtr("testString")
	networkPoolConfigModel.Subnet = core.StringPtr("testString")
	networkPoolConfigModel.UseSmartConnect = core.BoolPtr(true)

	zoneConfigModel := new(backuprecoveryv1.ZoneConfig)
	zoneConfigModel.DynamicNetworkPoolConfig = networkPoolConfigModel

	registeredProtectionSourceIsilonParamsModel := new(backuprecoveryv1.RegisteredProtectionSourceIsilonParams)
	registeredProtectionSourceIsilonParamsModel.ZoneConfigList = []backuprecoveryv1.ZoneConfig{*zoneConfigModel}

	mongoDbConnectParamsModel := new(backuprecoveryv1.MongoDBConnectParams)
	mongoDbConnectParamsModel.AuthType = core.StringPtr("SCRAM")
	mongoDbConnectParamsModel.AuthenticatingDatabaseName = core.StringPtr("testString")
	mongoDbConnectParamsModel.RequiresSsl = core.BoolPtr(true)
	mongoDbConnectParamsModel.SecondaryNodeTag = core.StringPtr("testString")
	mongoDbConnectParamsModel.Seeds = []string{"testString"}
	mongoDbConnectParamsModel.UseFixedNodeForBackup = core.BoolPtr(true)
	mongoDbConnectParamsModel.UseSecondaryForBackup = core.BoolPtr(true)

	nasServerCredentialsModel := new(backuprecoveryv1.NASServerCredentials)
	nasServerCredentialsModel.Domain = core.StringPtr("testString")
	nasServerCredentialsModel.NasProtocol = core.StringPtr("kNoProtocol")

	sitesDiscoveryParamsModel := new(backuprecoveryv1.SitesDiscoveryParams)
	sitesDiscoveryParamsModel.EnableSiteTagging = core.BoolPtr(true)

	teamsAdditionalParamsModel := new(backuprecoveryv1.TeamsAdditionalParams)
	teamsAdditionalParamsModel.AllowPostsBackup = core.BoolPtr(true)

	usersDiscoveryParamsModel := new(backuprecoveryv1.UsersDiscoveryParams)
	usersDiscoveryParamsModel.AllowChatsBackup = core.BoolPtr(true)
	usersDiscoveryParamsModel.DiscoverUsersWithMailbox = core.BoolPtr(true)
	usersDiscoveryParamsModel.DiscoverUsersWithOnedrive = core.BoolPtr(true)
	usersDiscoveryParamsModel.FetchMailboxInfo = core.BoolPtr(true)
	usersDiscoveryParamsModel.FetchOneDriveInfo = core.BoolPtr(true)
	usersDiscoveryParamsModel.SkipUsersWithoutMySite = core.BoolPtr(true)

	objectsDiscoveryParamsModel := new(backuprecoveryv1.ObjectsDiscoveryParams)
	objectsDiscoveryParamsModel.DiscoverableObjectTypeList = []string{"testString"}
	objectsDiscoveryParamsModel.SitesDiscoveryParams = sitesDiscoveryParamsModel
	objectsDiscoveryParamsModel.TeamsAdditionalParams = teamsAdditionalParamsModel
	objectsDiscoveryParamsModel.UsersDiscoveryParams = usersDiscoveryParamsModel

	m365CsmParamsModel := new(backuprecoveryv1.M365CsmParams)
	m365CsmParamsModel.BackupAllowed = core.BoolPtr(true)

	o365ConnectParamsModel := new(backuprecoveryv1.O365ConnectParams)
	o365ConnectParamsModel.ObjectsDiscoveryParams = objectsDiscoveryParamsModel
	o365ConnectParamsModel.CsmParams = m365CsmParamsModel

	office365CredentialsModel := new(backuprecoveryv1.Office365Credentials)
	office365CredentialsModel.ClientID = core.StringPtr("testString")
	office365CredentialsModel.ClientSecret = core.StringPtr("testString")
	office365CredentialsModel.GrantType = core.StringPtr("testString")
	office365CredentialsModel.Scope = core.StringPtr("testString")
	office365CredentialsModel.UseOAuthForExchangeOnline = core.BoolPtr(true)

	credentialsModel := new(backuprecoveryv1.Credentials)
	credentialsModel.Username = core.StringPtr("testString")
	credentialsModel.Password = core.StringPtr("testString")

	throttlingConfigurationModel := new(backuprecoveryv1.ThrottlingConfiguration)
	throttlingConfigurationModel.FixedThreshold = core.Int64Ptr(int64(26))
	throttlingConfigurationModel.PatternType = core.StringPtr("kNoThrottling")
	throttlingConfigurationModel.ThrottlingWindows = []backuprecoveryv1.ThrottlingWindow{*throttlingWindowModel}

	sourceThrottlingConfigurationModel := new(backuprecoveryv1.SourceThrottlingConfiguration)
	sourceThrottlingConfigurationModel.CpuThrottlingConfig = throttlingConfigurationModel
	sourceThrottlingConfigurationModel.NetworkThrottlingConfig = throttlingConfigurationModel

	physicalParamsModel := new(backuprecoveryv1.PhysicalParams)
	physicalParamsModel.Applications = []string{"kVMware"}
	physicalParamsModel.Password = core.StringPtr("testString")
	physicalParamsModel.ThrottlingConfig = sourceThrottlingConfigurationModel
	physicalParamsModel.Username = core.StringPtr("testString")

	sfdcParamsModel := new(backuprecoveryv1.SfdcParams)
	sfdcParamsModel.AccessToken = core.StringPtr("testString")
	sfdcParamsModel.ConcurrentApiRequestsLimit = core.Int64Ptr(int64(26))
	sfdcParamsModel.ConsumerKey = core.StringPtr("testString")
	sfdcParamsModel.ConsumerSecret = core.StringPtr("testString")
	sfdcParamsModel.DailyApiLimit = core.Int64Ptr(int64(26))
	sfdcParamsModel.Endpoint = core.StringPtr("testString")
	sfdcParamsModel.EndpointType = core.StringPtr("PROD")
	sfdcParamsModel.MetadataEndpointURL = core.StringPtr("testString")
	sfdcParamsModel.RefreshToken = core.StringPtr("testString")
	sfdcParamsModel.SoapEndpointURL = core.StringPtr("testString")
	sfdcParamsModel.UseBulkApi = core.BoolPtr(true)

	nasSourceThrottlingParamsModel := new(backuprecoveryv1.NasSourceThrottlingParams)
	nasSourceThrottlingParamsModel.MaxParallelMetadataFetchFullPercentage = core.Int64Ptr(int64(38))
	nasSourceThrottlingParamsModel.MaxParallelMetadataFetchIncrementalPercentage = core.Int64Ptr(int64(38))
	nasSourceThrottlingParamsModel.MaxParallelReadWriteFullPercentage = core.Int64Ptr(int64(38))
	nasSourceThrottlingParamsModel.MaxParallelReadWriteIncrementalPercentage = core.Int64Ptr(int64(38))

	storageArraySnapshotMaxSpaceConfigParamsModel := new(backuprecoveryv1.StorageArraySnapshotMaxSpaceConfigParams)
	storageArraySnapshotMaxSpaceConfigParamsModel.MaxSnapshotSpacePercentage = core.Int64Ptr(int64(38))

	storageArraySnapshotMaxSnapshotConfigParamsModel := new(backuprecoveryv1.StorageArraySnapshotMaxSnapshotConfigParams)
	storageArraySnapshotMaxSnapshotConfigParamsModel.MaxSnapshots = core.Int64Ptr(int64(38))

	storageArraySnapshotThrottlingPolicyModel := new(backuprecoveryv1.StorageArraySnapshotThrottlingPolicy)
	storageArraySnapshotThrottlingPolicyModel.ID = core.Int64Ptr(int64(26))
	storageArraySnapshotThrottlingPolicyModel.IsMaxSnapshotsConfigEnabled = core.BoolPtr(true)
	storageArraySnapshotThrottlingPolicyModel.IsMaxSpaceConfigEnabled = core.BoolPtr(true)
	storageArraySnapshotThrottlingPolicyModel.MaxSnapshotConfig = storageArraySnapshotMaxSnapshotConfigParamsModel
	storageArraySnapshotThrottlingPolicyModel.MaxSpaceConfig = storageArraySnapshotMaxSpaceConfigParamsModel

	storageArraySnapshotConfigParamsModel := new(backuprecoveryv1.StorageArraySnapshotConfigParams)
	storageArraySnapshotConfigParamsModel.IsMaxSnapshotsConfigEnabled = core.BoolPtr(true)
	storageArraySnapshotConfigParamsModel.IsMaxSpaceConfigEnabled = core.BoolPtr(true)
	storageArraySnapshotConfigParamsModel.StorageArraySnapshotMaxSpaceConfig = storageArraySnapshotMaxSpaceConfigParamsModel
	storageArraySnapshotConfigParamsModel.StorageArraySnapshotThrottlingPolicies = []backuprecoveryv1.StorageArraySnapshotThrottlingPolicy{*storageArraySnapshotThrottlingPolicyModel}

	throttlingPolicyParametersModel := new(backuprecoveryv1.ThrottlingPolicyParameters)
	throttlingPolicyParametersModel.EnforceMaxStreams = core.BoolPtr(true)
	throttlingPolicyParametersModel.EnforceRegisteredSourceMaxBackups = core.BoolPtr(true)
	throttlingPolicyParametersModel.IsEnabled = core.BoolPtr(true)
	throttlingPolicyParametersModel.LatencyThresholds = latencyThresholdsModel
	throttlingPolicyParametersModel.MaxConcurrentStreams = core.Int64Ptr(int64(38))
	throttlingPolicyParametersModel.NasSourceParams = nasSourceThrottlingParamsModel
	throttlingPolicyParametersModel.RegisteredSourceMaxConcurrentBackups = core.Int64Ptr(int64(38))
	throttlingPolicyParametersModel.StorageArraySnapshotConfig = storageArraySnapshotConfigParamsModel

	throttlingPolicyOverrideModel := new(backuprecoveryv1.ThrottlingPolicyOverride)
	throttlingPolicyOverrideModel.DatastoreID = core.Int64Ptr(int64(26))
	throttlingPolicyOverrideModel.DatastoreName = core.StringPtr("testString")
	throttlingPolicyOverrideModel.ThrottlingPolicy = throttlingPolicyParametersModel

	udaSourceCapabilitiesModel := new(backuprecoveryv1.UdaSourceCapabilities)
	udaSourceCapabilitiesModel.AutoLogBackup = core.BoolPtr(true)
	udaSourceCapabilitiesModel.DynamicConfig = core.BoolPtr(true)
	udaSourceCapabilitiesModel.EntitySupport = core.BoolPtr(true)
	udaSourceCapabilitiesModel.EtLogBackup = core.BoolPtr(true)
	udaSourceCapabilitiesModel.ExternalDisks = core.BoolPtr(true)
	udaSourceCapabilitiesModel.FullBackup = core.BoolPtr(true)
	udaSourceCapabilitiesModel.IncrBackup = core.BoolPtr(true)
	udaSourceCapabilitiesModel.LogBackup = core.BoolPtr(true)
	udaSourceCapabilitiesModel.MultiObjectRestore = core.BoolPtr(true)
	udaSourceCapabilitiesModel.PauseResumeBackup = core.BoolPtr(true)
	udaSourceCapabilitiesModel.PostBackupJobScript = core.BoolPtr(true)
	udaSourceCapabilitiesModel.PostRestoreJobScript = core.BoolPtr(true)
	udaSourceCapabilitiesModel.PreBackupJobScript = core.BoolPtr(true)
	udaSourceCapabilitiesModel.PreRestoreJobScript = core.BoolPtr(true)
	udaSourceCapabilitiesModel.ResourceThrottling = core.BoolPtr(true)
	udaSourceCapabilitiesModel.SnapfsCert = core.BoolPtr(true)

	keyValueStrPairModel := new(backuprecoveryv1.KeyValueStrPair)
	keyValueStrPairModel.Key = core.StringPtr("testString")
	keyValueStrPairModel.Value = core.StringPtr("testString")

	udaConnectParamsModel := new(backuprecoveryv1.UdaConnectParams)
	udaConnectParamsModel.Capabilities = udaSourceCapabilitiesModel
	udaConnectParamsModel.Credentials = credentialsModel
	udaConnectParamsModel.EtEnableLogBackupPolicy = core.BoolPtr(true)
	udaConnectParamsModel.EtEnableRunNow = core.BoolPtr(true)
	udaConnectParamsModel.HostType = core.StringPtr("kLinux")
	udaConnectParamsModel.Hosts = []string{"testString"}
	udaConnectParamsModel.LiveDataView = core.BoolPtr(true)
	udaConnectParamsModel.LiveLogView = core.BoolPtr(true)
	udaConnectParamsModel.MountDir = core.StringPtr("testString")
	udaConnectParamsModel.MountView = core.BoolPtr(true)
	udaConnectParamsModel.ScriptDir = core.StringPtr("testString")
	udaConnectParamsModel.SourceArgs = core.StringPtr("testString")
	udaConnectParamsModel.SourceRegistrationArguments = []backuprecoveryv1.KeyValueStrPair{*keyValueStrPairModel}
	udaConnectParamsModel.SourceType = core.StringPtr("testString")

	registeredSourceInfoModel := new(backuprecoveryv1.RegisteredSourceInfo)
	registeredSourceInfoModel.AccessInfo = connectorParametersModel
	registeredSourceInfoModel.AllowedIpAddresses = []string{"testString"}
	registeredSourceInfoModel.AuthenticationErrorMessage = core.StringPtr("testString")
	registeredSourceInfoModel.AuthenticationStatus = core.StringPtr("kPending")
	registeredSourceInfoModel.BlacklistedIpAddresses = []string{"testString"}
	registeredSourceInfoModel.CassandraParams = cassandraConnectParamsModel
	registeredSourceInfoModel.CouchbaseParams = couchbaseConnectParamsModel
	registeredSourceInfoModel.DeniedIpAddresses = []string{"testString"}
	registeredSourceInfoModel.Environments = []string{"kVMware"}
	registeredSourceInfoModel.HbaseParams = hBaseConnectParamsModel
	registeredSourceInfoModel.HdfsParams = hdfsConnectParamsModel
	registeredSourceInfoModel.HiveParams = hiveConnectParamsModel
	registeredSourceInfoModel.IsDbAuthenticated = core.BoolPtr(true)
	registeredSourceInfoModel.IsStorageArraySnapshotEnabled = core.BoolPtr(true)
	registeredSourceInfoModel.IsilonParams = registeredProtectionSourceIsilonParamsModel
	registeredSourceInfoModel.LinkVmsAcrossVcenter = core.BoolPtr(true)
	registeredSourceInfoModel.MinimumFreeSpaceGB = core.Int64Ptr(int64(26))
	registeredSourceInfoModel.MinimumFreeSpacePercent = core.Int64Ptr(int64(26))
	registeredSourceInfoModel.MongodbParams = mongoDbConnectParamsModel
	registeredSourceInfoModel.NasMountCredentials = nasServerCredentialsModel
	registeredSourceInfoModel.O365Params = o365ConnectParamsModel
	registeredSourceInfoModel.Office365CredentialsList = []backuprecoveryv1.Office365Credentials{*office365CredentialsModel}
	registeredSourceInfoModel.Office365Region = core.StringPtr("testString")
	registeredSourceInfoModel.Office365ServiceAccountCredentialsList = []backuprecoveryv1.Credentials{*credentialsModel}
	registeredSourceInfoModel.Password = core.StringPtr("testString")
	registeredSourceInfoModel.PhysicalParams = physicalParamsModel
	registeredSourceInfoModel.ProgressMonitorPath = core.StringPtr("testString")
	registeredSourceInfoModel.RefreshErrorMessage = core.StringPtr("testString")
	registeredSourceInfoModel.RefreshTimeUsecs = core.Int64Ptr(int64(26))
	registeredSourceInfoModel.RegisteredAppsInfo = []backuprecoveryv1.RegisteredAppInfo{*registeredAppInfoModel}
	registeredSourceInfoModel.RegistrationTimeUsecs = core.Int64Ptr(int64(26))
	registeredSourceInfoModel.SfdcParams = sfdcParamsModel
	registeredSourceInfoModel.Subnets = []backuprecoveryv1.Subnet{*subnetModel}
	registeredSourceInfoModel.ThrottlingPolicy = throttlingPolicyParametersModel
	registeredSourceInfoModel.ThrottlingPolicyOverrides = []backuprecoveryv1.ThrottlingPolicyOverride{*throttlingPolicyOverrideModel}
	registeredSourceInfoModel.UdaParams = udaConnectParamsModel
	registeredSourceInfoModel.UpdateLastBackupDetails = core.BoolPtr(true)
	registeredSourceInfoModel.UseOAuthForExchangeOnline = core.BoolPtr(true)
	registeredSourceInfoModel.UseVmBiosUUID = core.BoolPtr(true)
	registeredSourceInfoModel.UserMessages = []string{"testString"}
	registeredSourceInfoModel.Username = core.StringPtr("testString")
	registeredSourceInfoModel.VlanParams = vlanParametersModel
	registeredSourceInfoModel.WarningMessages = []string{"testString"}

	protectionSourceTreeInfoStatsModel := new(backuprecoveryv1.ProtectionSourceTreeInfoStats)
	protectionSourceTreeInfoStatsModel.ProtectedCount = core.Int64Ptr(int64(26))
	protectionSourceTreeInfoStatsModel.ProtectedSize = core.Int64Ptr(int64(26))
	protectionSourceTreeInfoStatsModel.UnprotectedCount = core.Int64Ptr(int64(26))
	protectionSourceTreeInfoStatsModel.UnprotectedSize = core.Int64Ptr(int64(26))

	protectionSummaryForK8sDistributionsModel := new(backuprecoveryv1.ProtectionSummaryForK8sDistributions)
	protectionSummaryForK8sDistributionsModel.Distribution = core.StringPtr("kMainline")
	protectionSummaryForK8sDistributionsModel.ProtectedCount = core.Int64Ptr(int64(26))
	protectionSummaryForK8sDistributionsModel.ProtectedSize = core.Int64Ptr(int64(26))
	protectionSummaryForK8sDistributionsModel.TotalRegisteredClusters = core.Int64Ptr(int64(26))
	protectionSummaryForK8sDistributionsModel.UnprotectedCount = core.Int64Ptr(int64(26))
	protectionSummaryForK8sDistributionsModel.UnprotectedSize = core.Int64Ptr(int64(26))

	protectionSummaryByEnvModel := new(backuprecoveryv1.ProtectionSummaryByEnv)
	protectionSummaryByEnvModel.Environment = core.StringPtr("kVMware")
	protectionSummaryByEnvModel.KubernetesDistributionStats = []backuprecoveryv1.ProtectionSummaryForK8sDistributions{*protectionSummaryForK8sDistributionsModel}
	protectionSummaryByEnvModel.ProtectedCount = core.Int64Ptr(int64(26))
	protectionSummaryByEnvModel.ProtectedSize = core.Int64Ptr(int64(26))
	protectionSummaryByEnvModel.UnprotectedCount = core.Int64Ptr(int64(26))
	protectionSummaryByEnvModel.UnprotectedSize = core.Int64Ptr(int64(26))

	model := new(backuprecoveryv1.ProtectionSourceTreeInfo)
	model.Applications = []backuprecoveryv1.ApplicationInfo{*applicationInfoModel}
	model.EntityPermissionInfo = entityPermissionInformationModel
	model.LogicalSizeBytes = core.Int64Ptr(int64(26))
	model.MaintenanceModeConfig = maintenanceModeConfigModel
	model.RegistrationInfo = registeredSourceInfoModel
	model.RootNode = protectionSourceNodeModel
	model.Stats = protectionSourceTreeInfoStatsModel
	model.StatsByEnv = []backuprecoveryv1.ProtectionSummaryByEnv{*protectionSummaryByEnvModel}
	model.TotalDowntieredSizeInBytes = core.Int64Ptr(int64(26))
	model.TotalUptieredSizeInBytes = core.Int64Ptr(int64(26))

	result, err := backuprecovery.DataSourceIbmBackupRecoveryRegistrationInfoProtectionSourceTreeInfoToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryRegistrationInfoApplicationInfoToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		cbtFileVersionModel := make(map[string]interface{})
		cbtFileVersionModel["build_ver"] = float64(72.5)
		cbtFileVersionModel["major_ver"] = float64(72.5)
		cbtFileVersionModel["minor_ver"] = float64(72.5)
		cbtFileVersionModel["revision_num"] = float64(72.5)

		cbtServiceStateModel := make(map[string]interface{})
		cbtServiceStateModel["name"] = "testString"
		cbtServiceStateModel["state"] = "testString"

		cbtInfoModel := make(map[string]interface{})
		cbtInfoModel["file_version"] = []map[string]interface{}{cbtFileVersionModel}
		cbtInfoModel["is_installed"] = true
		cbtInfoModel["reboot_status"] = "kRebooted"
		cbtInfoModel["service_state"] = []map[string]interface{}{cbtServiceStateModel}

		agentAccessInfoModel := make(map[string]interface{})
		agentAccessInfoModel["connection_id"] = int(26)
		agentAccessInfoModel["connector_group_id"] = int(26)
		agentAccessInfoModel["endpoint"] = "testString"
		agentAccessInfoModel["environment"] = "kPhysical"
		agentAccessInfoModel["id"] = int(26)
		agentAccessInfoModel["version"] = int(26)

		timeModel := make(map[string]interface{})
		timeModel["hour"] = int(38)
		timeModel["minute"] = int(38)

		dayTimeParamsModel := make(map[string]interface{})
		dayTimeParamsModel["day"] = "kSunday"
		dayTimeParamsModel["time"] = []map[string]interface{}{timeModel}

		dayTimeWindowModel := make(map[string]interface{})
		dayTimeWindowModel["end_time"] = []map[string]interface{}{dayTimeParamsModel}
		dayTimeWindowModel["start_time"] = []map[string]interface{}{dayTimeParamsModel}

		throttlingWindowModel := make(map[string]interface{})
		throttlingWindowModel["day_time_window"] = []map[string]interface{}{dayTimeWindowModel}
		throttlingWindowModel["threshold"] = int(26)

		throttlingConfigurationParamsModel := make(map[string]interface{})
		throttlingConfigurationParamsModel["fixed_threshold"] = int(26)
		throttlingConfigurationParamsModel["pattern_type"] = "kNoThrottling"
		throttlingConfigurationParamsModel["throttling_windows"] = []map[string]interface{}{throttlingWindowModel}

		throttlingConfigModel := make(map[string]interface{})
		throttlingConfigModel["cpu_throttling_config"] = []map[string]interface{}{throttlingConfigurationParamsModel}
		throttlingConfigModel["network_throttling_config"] = []map[string]interface{}{throttlingConfigurationParamsModel}

		agentPhysicalParamsModel := make(map[string]interface{})
		agentPhysicalParamsModel["applications"] = []string{"kSQL"}
		agentPhysicalParamsModel["password"] = "testString"
		agentPhysicalParamsModel["throttling_config"] = []map[string]interface{}{throttlingConfigModel}
		agentPhysicalParamsModel["username"] = "testString"

		hostSettingsCheckResultModel := make(map[string]interface{})
		hostSettingsCheckResultModel["check_type"] = "kIsAgentPortAccessible"
		hostSettingsCheckResultModel["result_type"] = "kPass"
		hostSettingsCheckResultModel["user_message"] = "testString"

		registeredAppInfoModel := make(map[string]interface{})
		registeredAppInfoModel["authentication_error_message"] = "testString"
		registeredAppInfoModel["authentication_status"] = "kPending"
		registeredAppInfoModel["environment"] = "kPhysical"
		registeredAppInfoModel["host_settings_check_results"] = []map[string]interface{}{hostSettingsCheckResultModel}
		registeredAppInfoModel["refresh_error_message"] = "testString"

		subnetModel := make(map[string]interface{})
		subnetModel["component"] = "testString"
		subnetModel["description"] = "testString"
		subnetModel["id"] = float64(72.5)
		subnetModel["ip"] = "testString"
		subnetModel["netmask_bits"] = float64(72.5)
		subnetModel["netmask_ip4"] = "testString"
		subnetModel["nfs_access"] = "kDisabled"
		subnetModel["nfs_all_squash"] = true
		subnetModel["nfs_root_squash"] = true
		subnetModel["s3_access"] = "kDisabled"
		subnetModel["smb_access"] = "kDisabled"
		subnetModel["tenant_id"] = "testString"

		latencyThresholdsModel := make(map[string]interface{})
		latencyThresholdsModel["active_task_msecs"] = int(26)
		latencyThresholdsModel["new_task_msecs"] = int(26)

		nasSourceParamsModel := make(map[string]interface{})
		nasSourceParamsModel["max_parallel_metadata_fetch_full_percentage"] = float64(72.5)
		nasSourceParamsModel["max_parallel_metadata_fetch_incremental_percentage"] = float64(72.5)
		nasSourceParamsModel["max_parallel_read_write_full_percentage"] = float64(72.5)
		nasSourceParamsModel["max_parallel_read_write_incremental_percentage"] = float64(72.5)

		storageArraySnapshotMaxSpaceConfigModel := make(map[string]interface{})
		storageArraySnapshotMaxSpaceConfigModel["max_snapshot_space_percentage"] = float64(72.5)

		maxSnapshotConfigModel := make(map[string]interface{})
		maxSnapshotConfigModel["max_snapshots"] = float64(72.5)

		maxSpaceConfigModel := make(map[string]interface{})
		maxSpaceConfigModel["max_snapshot_space_percentage"] = float64(72.5)

		storageArraySnapshotThrottlingPoliciesModel := make(map[string]interface{})
		storageArraySnapshotThrottlingPoliciesModel["id"] = int(26)
		storageArraySnapshotThrottlingPoliciesModel["is_max_snapshots_config_enabled"] = true
		storageArraySnapshotThrottlingPoliciesModel["is_max_space_config_enabled"] = true
		storageArraySnapshotThrottlingPoliciesModel["max_snapshot_config"] = []map[string]interface{}{maxSnapshotConfigModel}
		storageArraySnapshotThrottlingPoliciesModel["max_space_config"] = []map[string]interface{}{maxSpaceConfigModel}

		storageArraySnapshotConfigModel := make(map[string]interface{})
		storageArraySnapshotConfigModel["is_max_snapshots_config_enabled"] = true
		storageArraySnapshotConfigModel["is_max_space_config_enabled"] = true
		storageArraySnapshotConfigModel["storage_array_snapshot_max_space_config"] = []map[string]interface{}{storageArraySnapshotMaxSpaceConfigModel}
		storageArraySnapshotConfigModel["storage_array_snapshot_throttling_policies"] = []map[string]interface{}{storageArraySnapshotThrottlingPoliciesModel}

		throttlingPolicyModel := make(map[string]interface{})
		throttlingPolicyModel["enforce_max_streams"] = true
		throttlingPolicyModel["enforce_registered_source_max_backups"] = true
		throttlingPolicyModel["is_enabled"] = true
		throttlingPolicyModel["latency_thresholds"] = []map[string]interface{}{latencyThresholdsModel}
		throttlingPolicyModel["max_concurrent_streams"] = float64(72.5)
		throttlingPolicyModel["nas_source_params"] = []map[string]interface{}{nasSourceParamsModel}
		throttlingPolicyModel["registered_source_max_concurrent_backups"] = float64(72.5)
		throttlingPolicyModel["storage_array_snapshot_config"] = []map[string]interface{}{storageArraySnapshotConfigModel}

		throttlingPolicyOverridesModel := make(map[string]interface{})
		throttlingPolicyOverridesModel["datastore_id"] = int(26)
		throttlingPolicyOverridesModel["datastore_name"] = "testString"
		throttlingPolicyOverridesModel["throttling_policy"] = []map[string]interface{}{throttlingPolicyModel}

		registeredSourceVlanConfigModel := make(map[string]interface{})
		registeredSourceVlanConfigModel["vlan"] = float64(72.5)
		registeredSourceVlanConfigModel["disable_vlan"] = true
		registeredSourceVlanConfigModel["interface_name"] = "testString"

		agentRegistrationInfoModel := make(map[string]interface{})
		agentRegistrationInfoModel["access_info"] = []map[string]interface{}{agentAccessInfoModel}
		agentRegistrationInfoModel["allowed_ip_addresses"] = []string{"testString"}
		agentRegistrationInfoModel["authentication_error_message"] = "testString"
		agentRegistrationInfoModel["authentication_status"] = "kPending"
		agentRegistrationInfoModel["blacklisted_ip_addresses"] = []string{"testString"}
		agentRegistrationInfoModel["denied_ip_addresses"] = []string{"testString"}
		agentRegistrationInfoModel["environments"] = []string{"kPhysical"}
		agentRegistrationInfoModel["is_db_authenticated"] = true
		agentRegistrationInfoModel["is_storage_array_snapshot_enabled"] = true
		agentRegistrationInfoModel["link_vms_across_vcenter"] = true
		agentRegistrationInfoModel["minimum_free_space_gb"] = int(26)
		agentRegistrationInfoModel["minimum_free_space_percent"] = int(26)
		agentRegistrationInfoModel["password"] = "testString"
		agentRegistrationInfoModel["physical_params"] = []map[string]interface{}{agentPhysicalParamsModel}
		agentRegistrationInfoModel["progress_monitor_path"] = "testString"
		agentRegistrationInfoModel["refresh_error_message"] = "testString"
		agentRegistrationInfoModel["refresh_time_usecs"] = int(26)
		agentRegistrationInfoModel["registered_apps_info"] = []map[string]interface{}{registeredAppInfoModel}
		agentRegistrationInfoModel["registration_time_usecs"] = int(26)
		agentRegistrationInfoModel["subnets"] = []map[string]interface{}{subnetModel}
		agentRegistrationInfoModel["throttling_policy"] = []map[string]interface{}{throttlingPolicyModel}
		agentRegistrationInfoModel["throttling_policy_overrides"] = []map[string]interface{}{throttlingPolicyOverridesModel}
		agentRegistrationInfoModel["use_o_auth_for_exchange_online"] = true
		agentRegistrationInfoModel["use_vm_bios_uuid"] = true
		agentRegistrationInfoModel["user_messages"] = []string{"testString"}
		agentRegistrationInfoModel["username"] = "testString"
		agentRegistrationInfoModel["vlan_params"] = []map[string]interface{}{registeredSourceVlanConfigModel}
		agentRegistrationInfoModel["warning_messages"] = []string{"testString"}

		agentInformationModel := make(map[string]interface{})
		agentInformationModel["cbmr_version"] = "testString"
		agentInformationModel["file_cbt_info"] = []map[string]interface{}{cbtInfoModel}
		agentInformationModel["host_type"] = "kLinux"
		agentInformationModel["id"] = int(26)
		agentInformationModel["name"] = "testString"
		agentInformationModel["oracle_multi_node_channel_supported"] = true
		agentInformationModel["registration_info"] = []map[string]interface{}{agentRegistrationInfoModel}
		agentInformationModel["source_side_dedup_enabled"] = true
		agentInformationModel["status"] = "kUnknown"
		agentInformationModel["status_message"] = "testString"
		agentInformationModel["upgradability"] = "kUpgradable"
		agentInformationModel["upgrade_status"] = "kIdle"
		agentInformationModel["upgrade_status_message"] = "testString"
		agentInformationModel["version"] = "testString"
		agentInformationModel["vol_cbt_info"] = []map[string]interface{}{cbtInfoModel}

		uniqueGlobalIdModel := make(map[string]interface{})
		uniqueGlobalIdModel["cluster_id"] = int(26)
		uniqueGlobalIdModel["cluster_incarnation_id"] = int(26)
		uniqueGlobalIdModel["id"] = int(26)

		clusterNetworkingEndpointModel := make(map[string]interface{})
		clusterNetworkingEndpointModel["fqdn"] = "testString"
		clusterNetworkingEndpointModel["ipv4_addr"] = "testString"
		clusterNetworkingEndpointModel["ipv6_addr"] = "testString"

		clusterNetworkResourceInformationModel := make(map[string]interface{})
		clusterNetworkResourceInformationModel["endpoints"] = []map[string]interface{}{clusterNetworkingEndpointModel}
		clusterNetworkResourceInformationModel["type"] = "testString"

		networkingInformationModel := make(map[string]interface{})
		networkingInformationModel["resource_vec"] = []map[string]interface{}{clusterNetworkResourceInformationModel}

		physicalVolumeModel := make(map[string]interface{})
		physicalVolumeModel["device_path"] = "testString"
		physicalVolumeModel["guid"] = "testString"
		physicalVolumeModel["is_boot_volume"] = true
		physicalVolumeModel["is_extended_attributes_supported"] = true
		physicalVolumeModel["is_protected"] = true
		physicalVolumeModel["is_shared_volume"] = true
		physicalVolumeModel["label"] = "testString"
		physicalVolumeModel["logical_size_bytes"] = float64(72.5)
		physicalVolumeModel["mount_points"] = []string{"testString"}
		physicalVolumeModel["mount_type"] = "testString"
		physicalVolumeModel["network_path"] = "testString"
		physicalVolumeModel["used_size_bytes"] = float64(72.5)

		vssWritersModel := make(map[string]interface{})
		vssWritersModel["is_writer_excluded"] = true
		vssWritersModel["writer_name"] = true

		physicalProtectionSourceModel := make(map[string]interface{})
		physicalProtectionSourceModel["agents"] = []map[string]interface{}{agentInformationModel}
		physicalProtectionSourceModel["cluster_source_type"] = "testString"
		physicalProtectionSourceModel["host_name"] = "testString"
		physicalProtectionSourceModel["host_type"] = "kLinux"
		physicalProtectionSourceModel["id"] = []map[string]interface{}{uniqueGlobalIdModel}
		physicalProtectionSourceModel["is_proxy_host"] = true
		physicalProtectionSourceModel["memory_size_bytes"] = int(26)
		physicalProtectionSourceModel["name"] = "testString"
		physicalProtectionSourceModel["networking_info"] = []map[string]interface{}{networkingInformationModel}
		physicalProtectionSourceModel["num_processors"] = int(26)
		physicalProtectionSourceModel["os_name"] = "testString"
		physicalProtectionSourceModel["type"] = "kGroup"
		physicalProtectionSourceModel["vcs_version"] = "testString"
		physicalProtectionSourceModel["volumes"] = []map[string]interface{}{physicalVolumeModel}
		physicalProtectionSourceModel["vsswriters"] = []map[string]interface{}{vssWritersModel}

		vlanParametersModel := make(map[string]interface{})
		vlanParametersModel["disable_vlan"] = true
		vlanParametersModel["interface_name"] = "testString"
		vlanParametersModel["vlan"] = int(38)

		kubernetesLabelAttributeModel := make(map[string]interface{})
		kubernetesLabelAttributeModel["id"] = int(26)
		kubernetesLabelAttributeModel["name"] = "testString"
		kubernetesLabelAttributeModel["uuid"] = "testString"

		k8sLabelModel := make(map[string]interface{})
		k8sLabelModel["key"] = "testString"
		k8sLabelModel["value"] = "testString"

		serviceAnnotationsEntryModel := make(map[string]interface{})
		serviceAnnotationsEntryModel["key"] = "testString"
		serviceAnnotationsEntryModel["value"] = "testString"

		kubernetesStorageClassInfoModel := make(map[string]interface{})
		kubernetesStorageClassInfoModel["name"] = "testString"
		kubernetesStorageClassInfoModel["provisioner"] = "testString"

		kubernetesServiceAnnotationObjectModel := make(map[string]interface{})
		kubernetesServiceAnnotationObjectModel["key"] = "testString"
		kubernetesServiceAnnotationObjectModel["value"] = "testString"

		vlanParamsModel := make(map[string]interface{})
		vlanParamsModel["disable_vlan"] = true
		vlanParamsModel["interface_name"] = "testString"
		vlanParamsModel["vlan_id"] = int(38)

		kubernetesVlanInfoModel := make(map[string]interface{})
		kubernetesVlanInfoModel["service_annotations"] = []map[string]interface{}{kubernetesServiceAnnotationObjectModel}
		kubernetesVlanInfoModel["vlan_params"] = []map[string]interface{}{vlanParamsModel}

		kubernetesProtectionSourceModel := make(map[string]interface{})
		kubernetesProtectionSourceModel["datamover_image_location"] = "testString"
		kubernetesProtectionSourceModel["datamover_service_type"] = int(38)
		kubernetesProtectionSourceModel["datamover_upgradability"] = "testString"
		kubernetesProtectionSourceModel["default_vlan_params"] = []map[string]interface{}{vlanParametersModel}
		kubernetesProtectionSourceModel["description"] = "testString"
		kubernetesProtectionSourceModel["distribution"] = "kMainline"
		kubernetesProtectionSourceModel["init_container_image_location"] = "testString"
		kubernetesProtectionSourceModel["label_attributes"] = []map[string]interface{}{kubernetesLabelAttributeModel}
		kubernetesProtectionSourceModel["name"] = "testString"
		kubernetesProtectionSourceModel["priority_class_name"] = "testString"
		kubernetesProtectionSourceModel["resource_annotation_list"] = []map[string]interface{}{k8sLabelModel}
		kubernetesProtectionSourceModel["resource_label_list"] = []map[string]interface{}{k8sLabelModel}
		kubernetesProtectionSourceModel["san_field"] = []string{"testString"}
		kubernetesProtectionSourceModel["service_annotations"] = []map[string]interface{}{serviceAnnotationsEntryModel}
		kubernetesProtectionSourceModel["storage_class"] = []map[string]interface{}{kubernetesStorageClassInfoModel}
		kubernetesProtectionSourceModel["type"] = "kCluster"
		kubernetesProtectionSourceModel["uuid"] = "testString"
		kubernetesProtectionSourceModel["velero_aws_plugin_image_location"] = "testString"
		kubernetesProtectionSourceModel["velero_image_location"] = "testString"
		kubernetesProtectionSourceModel["velero_openshift_plugin_image_location"] = "testString"
		kubernetesProtectionSourceModel["velero_upgradability"] = "testString"
		kubernetesProtectionSourceModel["vlan_info_vec"] = []map[string]interface{}{kubernetesVlanInfoModel}

		databaseFileInformationModel := make(map[string]interface{})
		databaseFileInformationModel["file_type"] = "kRows"
		databaseFileInformationModel["full_path"] = "testString"
		databaseFileInformationModel["size_bytes"] = int(26)

		sqlSourceIdModel := make(map[string]interface{})
		sqlSourceIdModel["created_date_msecs"] = int(26)
		sqlSourceIdModel["database_id"] = int(26)
		sqlSourceIdModel["instance_id"] = "testString"

		sqlServerInstanceVersionModel := make(map[string]interface{})
		sqlServerInstanceVersionModel["build"] = float64(72.5)
		sqlServerInstanceVersionModel["major_version"] = float64(72.5)
		sqlServerInstanceVersionModel["minor_version"] = float64(72.5)
		sqlServerInstanceVersionModel["revision"] = float64(72.5)
		sqlServerInstanceVersionModel["version_string"] = float64(72.5)

		sqlProtectionSourceModel := make(map[string]interface{})
		sqlProtectionSourceModel["is_available_for_vss_backup"] = true
		sqlProtectionSourceModel["created_timestamp"] = "testString"
		sqlProtectionSourceModel["database_name"] = "testString"
		sqlProtectionSourceModel["db_aag_entity_id"] = int(26)
		sqlProtectionSourceModel["db_aag_name"] = "testString"
		sqlProtectionSourceModel["db_compatibility_level"] = int(26)
		sqlProtectionSourceModel["db_file_groups"] = []string{"testString"}
		sqlProtectionSourceModel["db_files"] = []map[string]interface{}{databaseFileInformationModel}
		sqlProtectionSourceModel["db_owner_username"] = "testString"
		sqlProtectionSourceModel["default_database_location"] = "testString"
		sqlProtectionSourceModel["default_log_location"] = "testString"
		sqlProtectionSourceModel["id"] = []map[string]interface{}{sqlSourceIdModel}
		sqlProtectionSourceModel["is_encrypted"] = true
		sqlProtectionSourceModel["name"] = "testString"
		sqlProtectionSourceModel["owner_id"] = int(26)
		sqlProtectionSourceModel["recovery_model"] = "kSimpleRecoveryModel"
		sqlProtectionSourceModel["sql_server_db_state"] = "kOnline"
		sqlProtectionSourceModel["sql_server_instance_version"] = []map[string]interface{}{sqlServerInstanceVersionModel}
		sqlProtectionSourceModel["type"] = "kInstance"

		protectionSourceNodeModel := make(map[string]interface{})
		protectionSourceNodeModel["connection_id"] = int(26)
		protectionSourceNodeModel["connector_group_id"] = int(26)
		protectionSourceNodeModel["custom_name"] = "testString"
		protectionSourceNodeModel["environment"] = "kPhysical"
		protectionSourceNodeModel["id"] = int(26)
		protectionSourceNodeModel["name"] = "testString"
		protectionSourceNodeModel["parent_id"] = int(26)
		protectionSourceNodeModel["physical_protection_source"] = []map[string]interface{}{physicalProtectionSourceModel}
		protectionSourceNodeModel["kubernetes_protection_source"] = []map[string]interface{}{kubernetesProtectionSourceModel}
		protectionSourceNodeModel["sql_protection_source"] = []map[string]interface{}{sqlProtectionSourceModel}

		model := make(map[string]interface{})
		model["application_tree_info"] = []map[string]interface{}{protectionSourceNodeModel}
		model["environment"] = "kVMware"

		assert.Equal(t, result, model)
	}

	cbtFileVersionModel := new(backuprecoveryv1.CbtFileVersion)
	cbtFileVersionModel.BuildVer = core.Float64Ptr(float64(72.5))
	cbtFileVersionModel.MajorVer = core.Float64Ptr(float64(72.5))
	cbtFileVersionModel.MinorVer = core.Float64Ptr(float64(72.5))
	cbtFileVersionModel.RevisionNum = core.Float64Ptr(float64(72.5))

	cbtServiceStateModel := new(backuprecoveryv1.CbtServiceState)
	cbtServiceStateModel.Name = core.StringPtr("testString")
	cbtServiceStateModel.State = core.StringPtr("testString")

	cbtInfoModel := new(backuprecoveryv1.CbtInfo)
	cbtInfoModel.FileVersion = cbtFileVersionModel
	cbtInfoModel.IsInstalled = core.BoolPtr(true)
	cbtInfoModel.RebootStatus = core.StringPtr("kRebooted")
	cbtInfoModel.ServiceState = cbtServiceStateModel

	agentAccessInfoModel := new(backuprecoveryv1.AgentAccessInfo)
	agentAccessInfoModel.ConnectionID = core.Int64Ptr(int64(26))
	agentAccessInfoModel.ConnectorGroupID = core.Int64Ptr(int64(26))
	agentAccessInfoModel.Endpoint = core.StringPtr("testString")
	agentAccessInfoModel.Environment = core.StringPtr("kPhysical")
	agentAccessInfoModel.ID = core.Int64Ptr(int64(26))
	agentAccessInfoModel.Version = core.Int64Ptr(int64(26))

	timeModel := new(backuprecoveryv1.Time)
	timeModel.Hour = core.Int64Ptr(int64(38))
	timeModel.Minute = core.Int64Ptr(int64(38))

	dayTimeParamsModel := new(backuprecoveryv1.DayTimeParams)
	dayTimeParamsModel.Day = core.StringPtr("kSunday")
	dayTimeParamsModel.Time = timeModel

	dayTimeWindowModel := new(backuprecoveryv1.DayTimeWindow)
	dayTimeWindowModel.EndTime = dayTimeParamsModel
	dayTimeWindowModel.StartTime = dayTimeParamsModel

	throttlingWindowModel := new(backuprecoveryv1.ThrottlingWindow)
	throttlingWindowModel.DayTimeWindow = dayTimeWindowModel
	throttlingWindowModel.Threshold = core.Int64Ptr(int64(26))

	throttlingConfigurationParamsModel := new(backuprecoveryv1.ThrottlingConfigurationParams)
	throttlingConfigurationParamsModel.FixedThreshold = core.Int64Ptr(int64(26))
	throttlingConfigurationParamsModel.PatternType = core.StringPtr("kNoThrottling")
	throttlingConfigurationParamsModel.ThrottlingWindows = []backuprecoveryv1.ThrottlingWindow{*throttlingWindowModel}

	throttlingConfigModel := new(backuprecoveryv1.ThrottlingConfig)
	throttlingConfigModel.CpuThrottlingConfig = throttlingConfigurationParamsModel
	throttlingConfigModel.NetworkThrottlingConfig = throttlingConfigurationParamsModel

	agentPhysicalParamsModel := new(backuprecoveryv1.AgentPhysicalParams)
	agentPhysicalParamsModel.Applications = []string{"kSQL"}
	agentPhysicalParamsModel.Password = core.StringPtr("testString")
	agentPhysicalParamsModel.ThrottlingConfig = throttlingConfigModel
	agentPhysicalParamsModel.Username = core.StringPtr("testString")

	hostSettingsCheckResultModel := new(backuprecoveryv1.HostSettingsCheckResult)
	hostSettingsCheckResultModel.CheckType = core.StringPtr("kIsAgentPortAccessible")
	hostSettingsCheckResultModel.ResultType = core.StringPtr("kPass")
	hostSettingsCheckResultModel.UserMessage = core.StringPtr("testString")

	registeredAppInfoModel := new(backuprecoveryv1.RegisteredAppInfo)
	registeredAppInfoModel.AuthenticationErrorMessage = core.StringPtr("testString")
	registeredAppInfoModel.AuthenticationStatus = core.StringPtr("kPending")
	registeredAppInfoModel.Environment = core.StringPtr("kPhysical")
	registeredAppInfoModel.HostSettingsCheckResults = []backuprecoveryv1.HostSettingsCheckResult{*hostSettingsCheckResultModel}
	registeredAppInfoModel.RefreshErrorMessage = core.StringPtr("testString")

	subnetModel := new(backuprecoveryv1.Subnet)
	subnetModel.Component = core.StringPtr("testString")
	subnetModel.Description = core.StringPtr("testString")
	subnetModel.ID = core.Float64Ptr(float64(72.5))
	subnetModel.Ip = core.StringPtr("testString")
	subnetModel.NetmaskBits = core.Float64Ptr(float64(72.5))
	subnetModel.NetmaskIp4 = core.StringPtr("testString")
	subnetModel.NfsAccess = core.StringPtr("kDisabled")
	subnetModel.NfsAllSquash = core.BoolPtr(true)
	subnetModel.NfsRootSquash = core.BoolPtr(true)
	subnetModel.S3Access = core.StringPtr("kDisabled")
	subnetModel.SmbAccess = core.StringPtr("kDisabled")
	subnetModel.TenantID = core.StringPtr("testString")

	latencyThresholdsModel := new(backuprecoveryv1.LatencyThresholds)
	latencyThresholdsModel.ActiveTaskMsecs = core.Int64Ptr(int64(26))
	latencyThresholdsModel.NewTaskMsecs = core.Int64Ptr(int64(26))

	nasSourceParamsModel := new(backuprecoveryv1.NasSourceParams)
	nasSourceParamsModel.MaxParallelMetadataFetchFullPercentage = core.Float64Ptr(float64(72.5))
	nasSourceParamsModel.MaxParallelMetadataFetchIncrementalPercentage = core.Float64Ptr(float64(72.5))
	nasSourceParamsModel.MaxParallelReadWriteFullPercentage = core.Float64Ptr(float64(72.5))
	nasSourceParamsModel.MaxParallelReadWriteIncrementalPercentage = core.Float64Ptr(float64(72.5))

	storageArraySnapshotMaxSpaceConfigModel := new(backuprecoveryv1.StorageArraySnapshotMaxSpaceConfig)
	storageArraySnapshotMaxSpaceConfigModel.MaxSnapshotSpacePercentage = core.Float64Ptr(float64(72.5))

	maxSnapshotConfigModel := new(backuprecoveryv1.MaxSnapshotConfig)
	maxSnapshotConfigModel.MaxSnapshots = core.Float64Ptr(float64(72.5))

	maxSpaceConfigModel := new(backuprecoveryv1.MaxSpaceConfig)
	maxSpaceConfigModel.MaxSnapshotSpacePercentage = core.Float64Ptr(float64(72.5))

	storageArraySnapshotThrottlingPoliciesModel := new(backuprecoveryv1.StorageArraySnapshotThrottlingPolicies)
	storageArraySnapshotThrottlingPoliciesModel.ID = core.Int64Ptr(int64(26))
	storageArraySnapshotThrottlingPoliciesModel.IsMaxSnapshotsConfigEnabled = core.BoolPtr(true)
	storageArraySnapshotThrottlingPoliciesModel.IsMaxSpaceConfigEnabled = core.BoolPtr(true)
	storageArraySnapshotThrottlingPoliciesModel.MaxSnapshotConfig = maxSnapshotConfigModel
	storageArraySnapshotThrottlingPoliciesModel.MaxSpaceConfig = maxSpaceConfigModel

	storageArraySnapshotConfigModel := new(backuprecoveryv1.StorageArraySnapshotConfig)
	storageArraySnapshotConfigModel.IsMaxSnapshotsConfigEnabled = core.BoolPtr(true)
	storageArraySnapshotConfigModel.IsMaxSpaceConfigEnabled = core.BoolPtr(true)
	storageArraySnapshotConfigModel.StorageArraySnapshotMaxSpaceConfig = storageArraySnapshotMaxSpaceConfigModel
	storageArraySnapshotConfigModel.StorageArraySnapshotThrottlingPolicies = []backuprecoveryv1.StorageArraySnapshotThrottlingPolicies{*storageArraySnapshotThrottlingPoliciesModel}

	throttlingPolicyModel := new(backuprecoveryv1.ThrottlingPolicy)
	throttlingPolicyModel.EnforceMaxStreams = core.BoolPtr(true)
	throttlingPolicyModel.EnforceRegisteredSourceMaxBackups = core.BoolPtr(true)
	throttlingPolicyModel.IsEnabled = core.BoolPtr(true)
	throttlingPolicyModel.LatencyThresholds = latencyThresholdsModel
	throttlingPolicyModel.MaxConcurrentStreams = core.Float64Ptr(float64(72.5))
	throttlingPolicyModel.NasSourceParams = nasSourceParamsModel
	throttlingPolicyModel.RegisteredSourceMaxConcurrentBackups = core.Float64Ptr(float64(72.5))
	throttlingPolicyModel.StorageArraySnapshotConfig = storageArraySnapshotConfigModel

	throttlingPolicyOverridesModel := new(backuprecoveryv1.ThrottlingPolicyOverrides)
	throttlingPolicyOverridesModel.DatastoreID = core.Int64Ptr(int64(26))
	throttlingPolicyOverridesModel.DatastoreName = core.StringPtr("testString")
	throttlingPolicyOverridesModel.ThrottlingPolicy = throttlingPolicyModel

	registeredSourceVlanConfigModel := new(backuprecoveryv1.RegisteredSourceVlanConfig)
	registeredSourceVlanConfigModel.Vlan = core.Float64Ptr(float64(72.5))
	registeredSourceVlanConfigModel.DisableVlan = core.BoolPtr(true)
	registeredSourceVlanConfigModel.InterfaceName = core.StringPtr("testString")

	agentRegistrationInfoModel := new(backuprecoveryv1.AgentRegistrationInfo)
	agentRegistrationInfoModel.AccessInfo = agentAccessInfoModel
	agentRegistrationInfoModel.AllowedIpAddresses = []string{"testString"}
	agentRegistrationInfoModel.AuthenticationErrorMessage = core.StringPtr("testString")
	agentRegistrationInfoModel.AuthenticationStatus = core.StringPtr("kPending")
	agentRegistrationInfoModel.BlacklistedIpAddresses = []string{"testString"}
	agentRegistrationInfoModel.DeniedIpAddresses = []string{"testString"}
	agentRegistrationInfoModel.Environments = []string{"kPhysical"}
	agentRegistrationInfoModel.IsDbAuthenticated = core.BoolPtr(true)
	agentRegistrationInfoModel.IsStorageArraySnapshotEnabled = core.BoolPtr(true)
	agentRegistrationInfoModel.LinkVmsAcrossVcenter = core.BoolPtr(true)
	agentRegistrationInfoModel.MinimumFreeSpaceGB = core.Int64Ptr(int64(26))
	agentRegistrationInfoModel.MinimumFreeSpacePercent = core.Int64Ptr(int64(26))
	agentRegistrationInfoModel.Password = core.StringPtr("testString")
	agentRegistrationInfoModel.PhysicalParams = agentPhysicalParamsModel
	agentRegistrationInfoModel.ProgressMonitorPath = core.StringPtr("testString")
	agentRegistrationInfoModel.RefreshErrorMessage = core.StringPtr("testString")
	agentRegistrationInfoModel.RefreshTimeUsecs = core.Int64Ptr(int64(26))
	agentRegistrationInfoModel.RegisteredAppsInfo = []backuprecoveryv1.RegisteredAppInfo{*registeredAppInfoModel}
	agentRegistrationInfoModel.RegistrationTimeUsecs = core.Int64Ptr(int64(26))
	agentRegistrationInfoModel.Subnets = []backuprecoveryv1.Subnet{*subnetModel}
	agentRegistrationInfoModel.ThrottlingPolicy = throttlingPolicyModel
	agentRegistrationInfoModel.ThrottlingPolicyOverrides = []backuprecoveryv1.ThrottlingPolicyOverrides{*throttlingPolicyOverridesModel}
	agentRegistrationInfoModel.UseOAuthForExchangeOnline = core.BoolPtr(true)
	agentRegistrationInfoModel.UseVmBiosUUID = core.BoolPtr(true)
	agentRegistrationInfoModel.UserMessages = []string{"testString"}
	agentRegistrationInfoModel.Username = core.StringPtr("testString")
	agentRegistrationInfoModel.VlanParams = registeredSourceVlanConfigModel
	agentRegistrationInfoModel.WarningMessages = []string{"testString"}

	agentInformationModel := new(backuprecoveryv1.AgentInformation)
	agentInformationModel.CbmrVersion = core.StringPtr("testString")
	agentInformationModel.FileCbtInfo = cbtInfoModel
	agentInformationModel.HostType = core.StringPtr("kLinux")
	agentInformationModel.ID = core.Int64Ptr(int64(26))
	agentInformationModel.Name = core.StringPtr("testString")
	agentInformationModel.OracleMultiNodeChannelSupported = core.BoolPtr(true)
	agentInformationModel.RegistrationInfo = agentRegistrationInfoModel
	agentInformationModel.SourceSideDedupEnabled = core.BoolPtr(true)
	agentInformationModel.Status = core.StringPtr("kUnknown")
	agentInformationModel.StatusMessage = core.StringPtr("testString")
	agentInformationModel.Upgradability = core.StringPtr("kUpgradable")
	agentInformationModel.UpgradeStatus = core.StringPtr("kIdle")
	agentInformationModel.UpgradeStatusMessage = core.StringPtr("testString")
	agentInformationModel.Version = core.StringPtr("testString")
	agentInformationModel.VolCbtInfo = cbtInfoModel

	uniqueGlobalIdModel := new(backuprecoveryv1.UniqueGlobalID)
	uniqueGlobalIdModel.ClusterID = core.Int64Ptr(int64(26))
	uniqueGlobalIdModel.ClusterIncarnationID = core.Int64Ptr(int64(26))
	uniqueGlobalIdModel.ID = core.Int64Ptr(int64(26))

	clusterNetworkingEndpointModel := new(backuprecoveryv1.ClusterNetworkingEndpoint)
	clusterNetworkingEndpointModel.Fqdn = core.StringPtr("testString")
	clusterNetworkingEndpointModel.Ipv4Addr = core.StringPtr("testString")
	clusterNetworkingEndpointModel.Ipv6Addr = core.StringPtr("testString")

	clusterNetworkResourceInformationModel := new(backuprecoveryv1.ClusterNetworkResourceInformation)
	clusterNetworkResourceInformationModel.Endpoints = []backuprecoveryv1.ClusterNetworkingEndpoint{*clusterNetworkingEndpointModel}
	clusterNetworkResourceInformationModel.Type = core.StringPtr("testString")

	networkingInformationModel := new(backuprecoveryv1.NetworkingInformation)
	networkingInformationModel.ResourceVec = []backuprecoveryv1.ClusterNetworkResourceInformation{*clusterNetworkResourceInformationModel}

	physicalVolumeModel := new(backuprecoveryv1.PhysicalVolume)
	physicalVolumeModel.DevicePath = core.StringPtr("testString")
	physicalVolumeModel.Guid = core.StringPtr("testString")
	physicalVolumeModel.IsBootVolume = core.BoolPtr(true)
	physicalVolumeModel.IsExtendedAttributesSupported = core.BoolPtr(true)
	physicalVolumeModel.IsProtected = core.BoolPtr(true)
	physicalVolumeModel.IsSharedVolume = core.BoolPtr(true)
	physicalVolumeModel.Label = core.StringPtr("testString")
	physicalVolumeModel.LogicalSizeBytes = core.Float64Ptr(float64(72.5))
	physicalVolumeModel.MountPoints = []string{"testString"}
	physicalVolumeModel.MountType = core.StringPtr("testString")
	physicalVolumeModel.NetworkPath = core.StringPtr("testString")
	physicalVolumeModel.UsedSizeBytes = core.Float64Ptr(float64(72.5))

	vssWritersModel := new(backuprecoveryv1.VssWriters)
	vssWritersModel.IsWriterExcluded = core.BoolPtr(true)
	vssWritersModel.WriterName = core.BoolPtr(true)

	physicalProtectionSourceModel := new(backuprecoveryv1.PhysicalProtectionSource)
	physicalProtectionSourceModel.Agents = []backuprecoveryv1.AgentInformation{*agentInformationModel}
	physicalProtectionSourceModel.ClusterSourceType = core.StringPtr("testString")
	physicalProtectionSourceModel.HostName = core.StringPtr("testString")
	physicalProtectionSourceModel.HostType = core.StringPtr("kLinux")
	physicalProtectionSourceModel.ID = uniqueGlobalIdModel
	physicalProtectionSourceModel.IsProxyHost = core.BoolPtr(true)
	physicalProtectionSourceModel.MemorySizeBytes = core.Int64Ptr(int64(26))
	physicalProtectionSourceModel.Name = core.StringPtr("testString")
	physicalProtectionSourceModel.NetworkingInfo = networkingInformationModel
	physicalProtectionSourceModel.NumProcessors = core.Int64Ptr(int64(26))
	physicalProtectionSourceModel.OsName = core.StringPtr("testString")
	physicalProtectionSourceModel.Type = core.StringPtr("kGroup")
	physicalProtectionSourceModel.VcsVersion = core.StringPtr("testString")
	physicalProtectionSourceModel.Volumes = []backuprecoveryv1.PhysicalVolume{*physicalVolumeModel}
	physicalProtectionSourceModel.Vsswriters = []backuprecoveryv1.VssWriters{*vssWritersModel}

	vlanParametersModel := new(backuprecoveryv1.VlanParameters)
	vlanParametersModel.DisableVlan = core.BoolPtr(true)
	vlanParametersModel.InterfaceName = core.StringPtr("testString")
	vlanParametersModel.Vlan = core.Int64Ptr(int64(38))

	kubernetesLabelAttributeModel := new(backuprecoveryv1.KubernetesLabelAttribute)
	kubernetesLabelAttributeModel.ID = core.Int64Ptr(int64(26))
	kubernetesLabelAttributeModel.Name = core.StringPtr("testString")
	kubernetesLabelAttributeModel.UUID = core.StringPtr("testString")

	k8sLabelModel := new(backuprecoveryv1.K8sLabel)
	k8sLabelModel.Key = core.StringPtr("testString")
	k8sLabelModel.Value = core.StringPtr("testString")

	serviceAnnotationsEntryModel := new(backuprecoveryv1.ServiceAnnotationsEntry)
	serviceAnnotationsEntryModel.Key = core.StringPtr("testString")
	serviceAnnotationsEntryModel.Value = core.StringPtr("testString")

	kubernetesStorageClassInfoModel := new(backuprecoveryv1.KubernetesStorageClassInfo)
	kubernetesStorageClassInfoModel.Name = core.StringPtr("testString")
	kubernetesStorageClassInfoModel.Provisioner = core.StringPtr("testString")

	kubernetesServiceAnnotationObjectModel := new(backuprecoveryv1.KubernetesServiceAnnotationObject)
	kubernetesServiceAnnotationObjectModel.Key = core.StringPtr("testString")
	kubernetesServiceAnnotationObjectModel.Value = core.StringPtr("testString")

	vlanParamsModel := new(backuprecoveryv1.VlanParams)
	vlanParamsModel.DisableVlan = core.BoolPtr(true)
	vlanParamsModel.InterfaceName = core.StringPtr("testString")
	vlanParamsModel.VlanID = core.Int64Ptr(int64(38))

	kubernetesVlanInfoModel := new(backuprecoveryv1.KubernetesVlanInfo)
	kubernetesVlanInfoModel.ServiceAnnotations = []backuprecoveryv1.KubernetesServiceAnnotationObject{*kubernetesServiceAnnotationObjectModel}
	kubernetesVlanInfoModel.VlanParams = vlanParamsModel

	kubernetesProtectionSourceModel := new(backuprecoveryv1.KubernetesProtectionSource)
	kubernetesProtectionSourceModel.DatamoverImageLocation = core.StringPtr("testString")
	kubernetesProtectionSourceModel.DatamoverServiceType = core.Int64Ptr(int64(38))
	kubernetesProtectionSourceModel.DatamoverUpgradability = core.StringPtr("kCurrent")
	kubernetesProtectionSourceModel.DefaultVlanParams = vlanParametersModel
	kubernetesProtectionSourceModel.Description = core.StringPtr("testString")
	kubernetesProtectionSourceModel.Distribution = core.StringPtr("kMainline")
	kubernetesProtectionSourceModel.InitContainerImageLocation = core.StringPtr("testString")
	kubernetesProtectionSourceModel.LabelAttributes = []backuprecoveryv1.KubernetesLabelAttribute{*kubernetesLabelAttributeModel}
	kubernetesProtectionSourceModel.Name = core.StringPtr("testString")
	kubernetesProtectionSourceModel.PriorityClassName = core.StringPtr("testString")
	kubernetesProtectionSourceModel.ResourceAnnotationList = []backuprecoveryv1.K8sLabel{*k8sLabelModel}
	kubernetesProtectionSourceModel.ResourceLabelList = []backuprecoveryv1.K8sLabel{*k8sLabelModel}
	kubernetesProtectionSourceModel.SanField = []string{"testString"}
	kubernetesProtectionSourceModel.ServiceAnnotations = []backuprecoveryv1.ServiceAnnotationsEntry{*serviceAnnotationsEntryModel}
	kubernetesProtectionSourceModel.StorageClass = []backuprecoveryv1.KubernetesStorageClassInfo{*kubernetesStorageClassInfoModel}
	kubernetesProtectionSourceModel.Type = core.StringPtr("kCluster")
	kubernetesProtectionSourceModel.UUID = core.StringPtr("testString")
	kubernetesProtectionSourceModel.VeleroAwsPluginImageLocation = core.StringPtr("testString")
	kubernetesProtectionSourceModel.VeleroImageLocation = core.StringPtr("testString")
	kubernetesProtectionSourceModel.VeleroOpenshiftPluginImageLocation = core.StringPtr("testString")
	kubernetesProtectionSourceModel.VeleroUpgradability = core.StringPtr("testString")
	kubernetesProtectionSourceModel.VlanInfoVec = []backuprecoveryv1.KubernetesVlanInfo{*kubernetesVlanInfoModel}

	databaseFileInformationModel := new(backuprecoveryv1.DatabaseFileInformation)
	databaseFileInformationModel.FileType = core.StringPtr("kRows")
	databaseFileInformationModel.FullPath = core.StringPtr("testString")
	databaseFileInformationModel.SizeBytes = core.Int64Ptr(int64(26))

	sqlSourceIdModel := new(backuprecoveryv1.SQLSourceID)
	sqlSourceIdModel.CreatedDateMsecs = core.Int64Ptr(int64(26))
	sqlSourceIdModel.DatabaseID = core.Int64Ptr(int64(26))
	sqlSourceIdModel.InstanceID = core.StringPtr("testString")

	sqlServerInstanceVersionModel := new(backuprecoveryv1.SQLServerInstanceVersion)
	sqlServerInstanceVersionModel.Build = core.Float64Ptr(float64(72.5))
	sqlServerInstanceVersionModel.MajorVersion = core.Float64Ptr(float64(72.5))
	sqlServerInstanceVersionModel.MinorVersion = core.Float64Ptr(float64(72.5))
	sqlServerInstanceVersionModel.Revision = core.Float64Ptr(float64(72.5))
	sqlServerInstanceVersionModel.VersionString = core.Float64Ptr(float64(72.5))

	sqlProtectionSourceModel := new(backuprecoveryv1.SqlProtectionSource)
	sqlProtectionSourceModel.IsAvailableForVssBackup = core.BoolPtr(true)
	sqlProtectionSourceModel.CreatedTimestamp = core.StringPtr("testString")
	sqlProtectionSourceModel.DatabaseName = core.StringPtr("testString")
	sqlProtectionSourceModel.DbAagEntityID = core.Int64Ptr(int64(26))
	sqlProtectionSourceModel.DbAagName = core.StringPtr("testString")
	sqlProtectionSourceModel.DbCompatibilityLevel = core.Int64Ptr(int64(26))
	sqlProtectionSourceModel.DbFileGroups = []string{"testString"}
	sqlProtectionSourceModel.DbFiles = []backuprecoveryv1.DatabaseFileInformation{*databaseFileInformationModel}
	sqlProtectionSourceModel.DbOwnerUsername = core.StringPtr("testString")
	sqlProtectionSourceModel.DefaultDatabaseLocation = core.StringPtr("testString")
	sqlProtectionSourceModel.DefaultLogLocation = core.StringPtr("testString")
	sqlProtectionSourceModel.ID = sqlSourceIdModel
	sqlProtectionSourceModel.IsEncrypted = core.BoolPtr(true)
	sqlProtectionSourceModel.Name = core.StringPtr("testString")
	sqlProtectionSourceModel.OwnerID = core.Int64Ptr(int64(26))
	sqlProtectionSourceModel.RecoveryModel = core.StringPtr("kSimpleRecoveryModel")
	sqlProtectionSourceModel.SqlServerDbState = core.StringPtr("kOnline")
	sqlProtectionSourceModel.SqlServerInstanceVersion = sqlServerInstanceVersionModel
	sqlProtectionSourceModel.Type = core.StringPtr("kInstance")

	protectionSourceNodeModel := new(backuprecoveryv1.ProtectionSourceNode)
	protectionSourceNodeModel.ConnectionID = core.Int64Ptr(int64(26))
	protectionSourceNodeModel.ConnectorGroupID = core.Int64Ptr(int64(26))
	protectionSourceNodeModel.CustomName = core.StringPtr("testString")
	protectionSourceNodeModel.Environment = core.StringPtr("kPhysical")
	protectionSourceNodeModel.ID = core.Int64Ptr(int64(26))
	protectionSourceNodeModel.Name = core.StringPtr("testString")
	protectionSourceNodeModel.ParentID = core.Int64Ptr(int64(26))
	protectionSourceNodeModel.PhysicalProtectionSource = physicalProtectionSourceModel
	protectionSourceNodeModel.KubernetesProtectionSource = kubernetesProtectionSourceModel
	protectionSourceNodeModel.SqlProtectionSource = sqlProtectionSourceModel

	model := new(backuprecoveryv1.ApplicationInfo)
	model.ApplicationTreeInfo = []backuprecoveryv1.ProtectionSourceNode{*protectionSourceNodeModel}
	model.Environment = core.StringPtr("kVMware")

	result, err := backuprecovery.DataSourceIbmBackupRecoveryRegistrationInfoApplicationInfoToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryRegistrationInfoProtectionSourceNodeToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		cbtFileVersionModel := make(map[string]interface{})
		cbtFileVersionModel["build_ver"] = float64(72.5)
		cbtFileVersionModel["major_ver"] = float64(72.5)
		cbtFileVersionModel["minor_ver"] = float64(72.5)
		cbtFileVersionModel["revision_num"] = float64(72.5)

		cbtServiceStateModel := make(map[string]interface{})
		cbtServiceStateModel["name"] = "testString"
		cbtServiceStateModel["state"] = "testString"

		cbtInfoModel := make(map[string]interface{})
		cbtInfoModel["file_version"] = []map[string]interface{}{cbtFileVersionModel}
		cbtInfoModel["is_installed"] = true
		cbtInfoModel["reboot_status"] = "kRebooted"
		cbtInfoModel["service_state"] = []map[string]interface{}{cbtServiceStateModel}

		agentAccessInfoModel := make(map[string]interface{})
		agentAccessInfoModel["connection_id"] = int(26)
		agentAccessInfoModel["connector_group_id"] = int(26)
		agentAccessInfoModel["endpoint"] = "testString"
		agentAccessInfoModel["environment"] = "kPhysical"
		agentAccessInfoModel["id"] = int(26)
		agentAccessInfoModel["version"] = int(26)

		timeModel := make(map[string]interface{})
		timeModel["hour"] = int(38)
		timeModel["minute"] = int(38)

		dayTimeParamsModel := make(map[string]interface{})
		dayTimeParamsModel["day"] = "kSunday"
		dayTimeParamsModel["time"] = []map[string]interface{}{timeModel}

		dayTimeWindowModel := make(map[string]interface{})
		dayTimeWindowModel["end_time"] = []map[string]interface{}{dayTimeParamsModel}
		dayTimeWindowModel["start_time"] = []map[string]interface{}{dayTimeParamsModel}

		throttlingWindowModel := make(map[string]interface{})
		throttlingWindowModel["day_time_window"] = []map[string]interface{}{dayTimeWindowModel}
		throttlingWindowModel["threshold"] = int(26)

		throttlingConfigurationParamsModel := make(map[string]interface{})
		throttlingConfigurationParamsModel["fixed_threshold"] = int(26)
		throttlingConfigurationParamsModel["pattern_type"] = "kNoThrottling"
		throttlingConfigurationParamsModel["throttling_windows"] = []map[string]interface{}{throttlingWindowModel}

		throttlingConfigModel := make(map[string]interface{})
		throttlingConfigModel["cpu_throttling_config"] = []map[string]interface{}{throttlingConfigurationParamsModel}
		throttlingConfigModel["network_throttling_config"] = []map[string]interface{}{throttlingConfigurationParamsModel}

		agentPhysicalParamsModel := make(map[string]interface{})
		agentPhysicalParamsModel["applications"] = []string{"kSQL"}
		agentPhysicalParamsModel["password"] = "testString"
		agentPhysicalParamsModel["throttling_config"] = []map[string]interface{}{throttlingConfigModel}
		agentPhysicalParamsModel["username"] = "testString"

		hostSettingsCheckResultModel := make(map[string]interface{})
		hostSettingsCheckResultModel["check_type"] = "kIsAgentPortAccessible"
		hostSettingsCheckResultModel["result_type"] = "kPass"
		hostSettingsCheckResultModel["user_message"] = "testString"

		registeredAppInfoModel := make(map[string]interface{})
		registeredAppInfoModel["authentication_error_message"] = "testString"
		registeredAppInfoModel["authentication_status"] = "kPending"
		registeredAppInfoModel["environment"] = "kPhysical"
		registeredAppInfoModel["host_settings_check_results"] = []map[string]interface{}{hostSettingsCheckResultModel}
		registeredAppInfoModel["refresh_error_message"] = "testString"

		subnetModel := make(map[string]interface{})
		subnetModel["component"] = "testString"
		subnetModel["description"] = "testString"
		subnetModel["id"] = float64(72.5)
		subnetModel["ip"] = "testString"
		subnetModel["netmask_bits"] = float64(72.5)
		subnetModel["netmask_ip4"] = "testString"
		subnetModel["nfs_access"] = "kDisabled"
		subnetModel["nfs_all_squash"] = true
		subnetModel["nfs_root_squash"] = true
		subnetModel["s3_access"] = "kDisabled"
		subnetModel["smb_access"] = "kDisabled"
		subnetModel["tenant_id"] = "testString"

		latencyThresholdsModel := make(map[string]interface{})
		latencyThresholdsModel["active_task_msecs"] = int(26)
		latencyThresholdsModel["new_task_msecs"] = int(26)

		nasSourceParamsModel := make(map[string]interface{})
		nasSourceParamsModel["max_parallel_metadata_fetch_full_percentage"] = float64(72.5)
		nasSourceParamsModel["max_parallel_metadata_fetch_incremental_percentage"] = float64(72.5)
		nasSourceParamsModel["max_parallel_read_write_full_percentage"] = float64(72.5)
		nasSourceParamsModel["max_parallel_read_write_incremental_percentage"] = float64(72.5)

		storageArraySnapshotMaxSpaceConfigModel := make(map[string]interface{})
		storageArraySnapshotMaxSpaceConfigModel["max_snapshot_space_percentage"] = float64(72.5)

		maxSnapshotConfigModel := make(map[string]interface{})
		maxSnapshotConfigModel["max_snapshots"] = float64(72.5)

		maxSpaceConfigModel := make(map[string]interface{})
		maxSpaceConfigModel["max_snapshot_space_percentage"] = float64(72.5)

		storageArraySnapshotThrottlingPoliciesModel := make(map[string]interface{})
		storageArraySnapshotThrottlingPoliciesModel["id"] = int(26)
		storageArraySnapshotThrottlingPoliciesModel["is_max_snapshots_config_enabled"] = true
		storageArraySnapshotThrottlingPoliciesModel["is_max_space_config_enabled"] = true
		storageArraySnapshotThrottlingPoliciesModel["max_snapshot_config"] = []map[string]interface{}{maxSnapshotConfigModel}
		storageArraySnapshotThrottlingPoliciesModel["max_space_config"] = []map[string]interface{}{maxSpaceConfigModel}

		storageArraySnapshotConfigModel := make(map[string]interface{})
		storageArraySnapshotConfigModel["is_max_snapshots_config_enabled"] = true
		storageArraySnapshotConfigModel["is_max_space_config_enabled"] = true
		storageArraySnapshotConfigModel["storage_array_snapshot_max_space_config"] = []map[string]interface{}{storageArraySnapshotMaxSpaceConfigModel}
		storageArraySnapshotConfigModel["storage_array_snapshot_throttling_policies"] = []map[string]interface{}{storageArraySnapshotThrottlingPoliciesModel}

		throttlingPolicyModel := make(map[string]interface{})
		throttlingPolicyModel["enforce_max_streams"] = true
		throttlingPolicyModel["enforce_registered_source_max_backups"] = true
		throttlingPolicyModel["is_enabled"] = true
		throttlingPolicyModel["latency_thresholds"] = []map[string]interface{}{latencyThresholdsModel}
		throttlingPolicyModel["max_concurrent_streams"] = float64(72.5)
		throttlingPolicyModel["nas_source_params"] = []map[string]interface{}{nasSourceParamsModel}
		throttlingPolicyModel["registered_source_max_concurrent_backups"] = float64(72.5)
		throttlingPolicyModel["storage_array_snapshot_config"] = []map[string]interface{}{storageArraySnapshotConfigModel}

		throttlingPolicyOverridesModel := make(map[string]interface{})
		throttlingPolicyOverridesModel["datastore_id"] = int(26)
		throttlingPolicyOverridesModel["datastore_name"] = "testString"
		throttlingPolicyOverridesModel["throttling_policy"] = []map[string]interface{}{throttlingPolicyModel}

		registeredSourceVlanConfigModel := make(map[string]interface{})
		registeredSourceVlanConfigModel["vlan"] = float64(72.5)
		registeredSourceVlanConfigModel["disable_vlan"] = true
		registeredSourceVlanConfigModel["interface_name"] = "testString"

		agentRegistrationInfoModel := make(map[string]interface{})
		agentRegistrationInfoModel["access_info"] = []map[string]interface{}{agentAccessInfoModel}
		agentRegistrationInfoModel["allowed_ip_addresses"] = []string{"testString"}
		agentRegistrationInfoModel["authentication_error_message"] = "testString"
		agentRegistrationInfoModel["authentication_status"] = "kPending"
		agentRegistrationInfoModel["blacklisted_ip_addresses"] = []string{"testString"}
		agentRegistrationInfoModel["denied_ip_addresses"] = []string{"testString"}
		agentRegistrationInfoModel["environments"] = []string{"kPhysical"}
		agentRegistrationInfoModel["is_db_authenticated"] = true
		agentRegistrationInfoModel["is_storage_array_snapshot_enabled"] = true
		agentRegistrationInfoModel["link_vms_across_vcenter"] = true
		agentRegistrationInfoModel["minimum_free_space_gb"] = int(26)
		agentRegistrationInfoModel["minimum_free_space_percent"] = int(26)
		agentRegistrationInfoModel["password"] = "testString"
		agentRegistrationInfoModel["physical_params"] = []map[string]interface{}{agentPhysicalParamsModel}
		agentRegistrationInfoModel["progress_monitor_path"] = "testString"
		agentRegistrationInfoModel["refresh_error_message"] = "testString"
		agentRegistrationInfoModel["refresh_time_usecs"] = int(26)
		agentRegistrationInfoModel["registered_apps_info"] = []map[string]interface{}{registeredAppInfoModel}
		agentRegistrationInfoModel["registration_time_usecs"] = int(26)
		agentRegistrationInfoModel["subnets"] = []map[string]interface{}{subnetModel}
		agentRegistrationInfoModel["throttling_policy"] = []map[string]interface{}{throttlingPolicyModel}
		agentRegistrationInfoModel["throttling_policy_overrides"] = []map[string]interface{}{throttlingPolicyOverridesModel}
		agentRegistrationInfoModel["use_o_auth_for_exchange_online"] = true
		agentRegistrationInfoModel["use_vm_bios_uuid"] = true
		agentRegistrationInfoModel["user_messages"] = []string{"testString"}
		agentRegistrationInfoModel["username"] = "testString"
		agentRegistrationInfoModel["vlan_params"] = []map[string]interface{}{registeredSourceVlanConfigModel}
		agentRegistrationInfoModel["warning_messages"] = []string{"testString"}

		agentInformationModel := make(map[string]interface{})
		agentInformationModel["cbmr_version"] = "testString"
		agentInformationModel["file_cbt_info"] = []map[string]interface{}{cbtInfoModel}
		agentInformationModel["host_type"] = "kLinux"
		agentInformationModel["id"] = int(26)
		agentInformationModel["name"] = "testString"
		agentInformationModel["oracle_multi_node_channel_supported"] = true
		agentInformationModel["registration_info"] = []map[string]interface{}{agentRegistrationInfoModel}
		agentInformationModel["source_side_dedup_enabled"] = true
		agentInformationModel["status"] = "kUnknown"
		agentInformationModel["status_message"] = "testString"
		agentInformationModel["upgradability"] = "kUpgradable"
		agentInformationModel["upgrade_status"] = "kIdle"
		agentInformationModel["upgrade_status_message"] = "testString"
		agentInformationModel["version"] = "testString"
		agentInformationModel["vol_cbt_info"] = []map[string]interface{}{cbtInfoModel}

		uniqueGlobalIdModel := make(map[string]interface{})
		uniqueGlobalIdModel["cluster_id"] = int(26)
		uniqueGlobalIdModel["cluster_incarnation_id"] = int(26)
		uniqueGlobalIdModel["id"] = int(26)

		clusterNetworkingEndpointModel := make(map[string]interface{})
		clusterNetworkingEndpointModel["fqdn"] = "testString"
		clusterNetworkingEndpointModel["ipv4_addr"] = "testString"
		clusterNetworkingEndpointModel["ipv6_addr"] = "testString"

		clusterNetworkResourceInformationModel := make(map[string]interface{})
		clusterNetworkResourceInformationModel["endpoints"] = []map[string]interface{}{clusterNetworkingEndpointModel}
		clusterNetworkResourceInformationModel["type"] = "testString"

		networkingInformationModel := make(map[string]interface{})
		networkingInformationModel["resource_vec"] = []map[string]interface{}{clusterNetworkResourceInformationModel}

		physicalVolumeModel := make(map[string]interface{})
		physicalVolumeModel["device_path"] = "testString"
		physicalVolumeModel["guid"] = "testString"
		physicalVolumeModel["is_boot_volume"] = true
		physicalVolumeModel["is_extended_attributes_supported"] = true
		physicalVolumeModel["is_protected"] = true
		physicalVolumeModel["is_shared_volume"] = true
		physicalVolumeModel["label"] = "testString"
		physicalVolumeModel["logical_size_bytes"] = float64(72.5)
		physicalVolumeModel["mount_points"] = []string{"testString"}
		physicalVolumeModel["mount_type"] = "testString"
		physicalVolumeModel["network_path"] = "testString"
		physicalVolumeModel["used_size_bytes"] = float64(72.5)

		vssWritersModel := make(map[string]interface{})
		vssWritersModel["is_writer_excluded"] = true
		vssWritersModel["writer_name"] = true

		physicalProtectionSourceModel := make(map[string]interface{})
		physicalProtectionSourceModel["agents"] = []map[string]interface{}{agentInformationModel}
		physicalProtectionSourceModel["cluster_source_type"] = "testString"
		physicalProtectionSourceModel["host_name"] = "testString"
		physicalProtectionSourceModel["host_type"] = "kLinux"
		physicalProtectionSourceModel["id"] = []map[string]interface{}{uniqueGlobalIdModel}
		physicalProtectionSourceModel["is_proxy_host"] = true
		physicalProtectionSourceModel["memory_size_bytes"] = int(26)
		physicalProtectionSourceModel["name"] = "testString"
		physicalProtectionSourceModel["networking_info"] = []map[string]interface{}{networkingInformationModel}
		physicalProtectionSourceModel["num_processors"] = int(26)
		physicalProtectionSourceModel["os_name"] = "testString"
		physicalProtectionSourceModel["type"] = "kGroup"
		physicalProtectionSourceModel["vcs_version"] = "testString"
		physicalProtectionSourceModel["volumes"] = []map[string]interface{}{physicalVolumeModel}
		physicalProtectionSourceModel["vsswriters"] = []map[string]interface{}{vssWritersModel}

		vlanParametersModel := make(map[string]interface{})
		vlanParametersModel["disable_vlan"] = true
		vlanParametersModel["interface_name"] = "testString"
		vlanParametersModel["vlan"] = int(38)

		kubernetesLabelAttributeModel := make(map[string]interface{})
		kubernetesLabelAttributeModel["id"] = int(26)
		kubernetesLabelAttributeModel["name"] = "testString"
		kubernetesLabelAttributeModel["uuid"] = "testString"

		k8sLabelModel := make(map[string]interface{})
		k8sLabelModel["key"] = "testString"
		k8sLabelModel["value"] = "testString"

		serviceAnnotationsEntryModel := make(map[string]interface{})
		serviceAnnotationsEntryModel["key"] = "testString"
		serviceAnnotationsEntryModel["value"] = "testString"

		kubernetesStorageClassInfoModel := make(map[string]interface{})
		kubernetesStorageClassInfoModel["name"] = "testString"
		kubernetesStorageClassInfoModel["provisioner"] = "testString"

		kubernetesServiceAnnotationObjectModel := make(map[string]interface{})
		kubernetesServiceAnnotationObjectModel["key"] = "testString"
		kubernetesServiceAnnotationObjectModel["value"] = "testString"

		vlanParamsModel := make(map[string]interface{})
		vlanParamsModel["disable_vlan"] = true
		vlanParamsModel["interface_name"] = "testString"
		vlanParamsModel["vlan_id"] = int(38)

		kubernetesVlanInfoModel := make(map[string]interface{})
		kubernetesVlanInfoModel["service_annotations"] = []map[string]interface{}{kubernetesServiceAnnotationObjectModel}
		kubernetesVlanInfoModel["vlan_params"] = []map[string]interface{}{vlanParamsModel}

		kubernetesProtectionSourceModel := make(map[string]interface{})
		kubernetesProtectionSourceModel["datamover_image_location"] = "testString"
		kubernetesProtectionSourceModel["datamover_service_type"] = int(38)
		kubernetesProtectionSourceModel["datamover_upgradability"] = "testString"
		kubernetesProtectionSourceModel["default_vlan_params"] = []map[string]interface{}{vlanParametersModel}
		kubernetesProtectionSourceModel["description"] = "testString"
		kubernetesProtectionSourceModel["distribution"] = "kMainline"
		kubernetesProtectionSourceModel["init_container_image_location"] = "testString"
		kubernetesProtectionSourceModel["label_attributes"] = []map[string]interface{}{kubernetesLabelAttributeModel}
		kubernetesProtectionSourceModel["name"] = "testString"
		kubernetesProtectionSourceModel["priority_class_name"] = "testString"
		kubernetesProtectionSourceModel["resource_annotation_list"] = []map[string]interface{}{k8sLabelModel}
		kubernetesProtectionSourceModel["resource_label_list"] = []map[string]interface{}{k8sLabelModel}
		kubernetesProtectionSourceModel["san_field"] = []string{"testString"}
		kubernetesProtectionSourceModel["service_annotations"] = []map[string]interface{}{serviceAnnotationsEntryModel}
		kubernetesProtectionSourceModel["storage_class"] = []map[string]interface{}{kubernetesStorageClassInfoModel}
		kubernetesProtectionSourceModel["type"] = "kCluster"
		kubernetesProtectionSourceModel["uuid"] = "testString"
		kubernetesProtectionSourceModel["velero_aws_plugin_image_location"] = "testString"
		kubernetesProtectionSourceModel["velero_image_location"] = "testString"
		kubernetesProtectionSourceModel["velero_openshift_plugin_image_location"] = "testString"
		kubernetesProtectionSourceModel["velero_upgradability"] = "testString"
		kubernetesProtectionSourceModel["vlan_info_vec"] = []map[string]interface{}{kubernetesVlanInfoModel}

		databaseFileInformationModel := make(map[string]interface{})
		databaseFileInformationModel["file_type"] = "kRows"
		databaseFileInformationModel["full_path"] = "testString"
		databaseFileInformationModel["size_bytes"] = int(26)

		sqlSourceIdModel := make(map[string]interface{})
		sqlSourceIdModel["created_date_msecs"] = int(26)
		sqlSourceIdModel["database_id"] = int(26)
		sqlSourceIdModel["instance_id"] = "testString"

		sqlServerInstanceVersionModel := make(map[string]interface{})
		sqlServerInstanceVersionModel["build"] = float64(72.5)
		sqlServerInstanceVersionModel["major_version"] = float64(72.5)
		sqlServerInstanceVersionModel["minor_version"] = float64(72.5)
		sqlServerInstanceVersionModel["revision"] = float64(72.5)
		sqlServerInstanceVersionModel["version_string"] = float64(72.5)

		sqlProtectionSourceModel := make(map[string]interface{})
		sqlProtectionSourceModel["is_available_for_vss_backup"] = true
		sqlProtectionSourceModel["created_timestamp"] = "testString"
		sqlProtectionSourceModel["database_name"] = "testString"
		sqlProtectionSourceModel["db_aag_entity_id"] = int(26)
		sqlProtectionSourceModel["db_aag_name"] = "testString"
		sqlProtectionSourceModel["db_compatibility_level"] = int(26)
		sqlProtectionSourceModel["db_file_groups"] = []string{"testString"}
		sqlProtectionSourceModel["db_files"] = []map[string]interface{}{databaseFileInformationModel}
		sqlProtectionSourceModel["db_owner_username"] = "testString"
		sqlProtectionSourceModel["default_database_location"] = "testString"
		sqlProtectionSourceModel["default_log_location"] = "testString"
		sqlProtectionSourceModel["id"] = []map[string]interface{}{sqlSourceIdModel}
		sqlProtectionSourceModel["is_encrypted"] = true
		sqlProtectionSourceModel["name"] = "testString"
		sqlProtectionSourceModel["owner_id"] = int(26)
		sqlProtectionSourceModel["recovery_model"] = "kSimpleRecoveryModel"
		sqlProtectionSourceModel["sql_server_db_state"] = "kOnline"
		sqlProtectionSourceModel["sql_server_instance_version"] = []map[string]interface{}{sqlServerInstanceVersionModel}
		sqlProtectionSourceModel["type"] = "kInstance"

		model := make(map[string]interface{})
		model["connection_id"] = int(26)
		model["connector_group_id"] = int(26)
		model["custom_name"] = "testString"
		model["environment"] = "kPhysical"
		model["id"] = int(26)
		model["name"] = "testString"
		model["parent_id"] = int(26)
		model["physical_protection_source"] = []map[string]interface{}{physicalProtectionSourceModel}
		model["kubernetes_protection_source"] = []map[string]interface{}{kubernetesProtectionSourceModel}
		model["sql_protection_source"] = []map[string]interface{}{sqlProtectionSourceModel}

		assert.Equal(t, result, model)
	}

	cbtFileVersionModel := new(backuprecoveryv1.CbtFileVersion)
	cbtFileVersionModel.BuildVer = core.Float64Ptr(float64(72.5))
	cbtFileVersionModel.MajorVer = core.Float64Ptr(float64(72.5))
	cbtFileVersionModel.MinorVer = core.Float64Ptr(float64(72.5))
	cbtFileVersionModel.RevisionNum = core.Float64Ptr(float64(72.5))

	cbtServiceStateModel := new(backuprecoveryv1.CbtServiceState)
	cbtServiceStateModel.Name = core.StringPtr("testString")
	cbtServiceStateModel.State = core.StringPtr("testString")

	cbtInfoModel := new(backuprecoveryv1.CbtInfo)
	cbtInfoModel.FileVersion = cbtFileVersionModel
	cbtInfoModel.IsInstalled = core.BoolPtr(true)
	cbtInfoModel.RebootStatus = core.StringPtr("kRebooted")
	cbtInfoModel.ServiceState = cbtServiceStateModel

	agentAccessInfoModel := new(backuprecoveryv1.AgentAccessInfo)
	agentAccessInfoModel.ConnectionID = core.Int64Ptr(int64(26))
	agentAccessInfoModel.ConnectorGroupID = core.Int64Ptr(int64(26))
	agentAccessInfoModel.Endpoint = core.StringPtr("testString")
	agentAccessInfoModel.Environment = core.StringPtr("kPhysical")
	agentAccessInfoModel.ID = core.Int64Ptr(int64(26))
	agentAccessInfoModel.Version = core.Int64Ptr(int64(26))

	timeModel := new(backuprecoveryv1.Time)
	timeModel.Hour = core.Int64Ptr(int64(38))
	timeModel.Minute = core.Int64Ptr(int64(38))

	dayTimeParamsModel := new(backuprecoveryv1.DayTimeParams)
	dayTimeParamsModel.Day = core.StringPtr("kSunday")
	dayTimeParamsModel.Time = timeModel

	dayTimeWindowModel := new(backuprecoveryv1.DayTimeWindow)
	dayTimeWindowModel.EndTime = dayTimeParamsModel
	dayTimeWindowModel.StartTime = dayTimeParamsModel

	throttlingWindowModel := new(backuprecoveryv1.ThrottlingWindow)
	throttlingWindowModel.DayTimeWindow = dayTimeWindowModel
	throttlingWindowModel.Threshold = core.Int64Ptr(int64(26))

	throttlingConfigurationParamsModel := new(backuprecoveryv1.ThrottlingConfigurationParams)
	throttlingConfigurationParamsModel.FixedThreshold = core.Int64Ptr(int64(26))
	throttlingConfigurationParamsModel.PatternType = core.StringPtr("kNoThrottling")
	throttlingConfigurationParamsModel.ThrottlingWindows = []backuprecoveryv1.ThrottlingWindow{*throttlingWindowModel}

	throttlingConfigModel := new(backuprecoveryv1.ThrottlingConfig)
	throttlingConfigModel.CpuThrottlingConfig = throttlingConfigurationParamsModel
	throttlingConfigModel.NetworkThrottlingConfig = throttlingConfigurationParamsModel

	agentPhysicalParamsModel := new(backuprecoveryv1.AgentPhysicalParams)
	agentPhysicalParamsModel.Applications = []string{"kSQL"}
	agentPhysicalParamsModel.Password = core.StringPtr("testString")
	agentPhysicalParamsModel.ThrottlingConfig = throttlingConfigModel
	agentPhysicalParamsModel.Username = core.StringPtr("testString")

	hostSettingsCheckResultModel := new(backuprecoveryv1.HostSettingsCheckResult)
	hostSettingsCheckResultModel.CheckType = core.StringPtr("kIsAgentPortAccessible")
	hostSettingsCheckResultModel.ResultType = core.StringPtr("kPass")
	hostSettingsCheckResultModel.UserMessage = core.StringPtr("testString")

	registeredAppInfoModel := new(backuprecoveryv1.RegisteredAppInfo)
	registeredAppInfoModel.AuthenticationErrorMessage = core.StringPtr("testString")
	registeredAppInfoModel.AuthenticationStatus = core.StringPtr("kPending")
	registeredAppInfoModel.Environment = core.StringPtr("kPhysical")
	registeredAppInfoModel.HostSettingsCheckResults = []backuprecoveryv1.HostSettingsCheckResult{*hostSettingsCheckResultModel}
	registeredAppInfoModel.RefreshErrorMessage = core.StringPtr("testString")

	subnetModel := new(backuprecoveryv1.Subnet)
	subnetModel.Component = core.StringPtr("testString")
	subnetModel.Description = core.StringPtr("testString")
	subnetModel.ID = core.Float64Ptr(float64(72.5))
	subnetModel.Ip = core.StringPtr("testString")
	subnetModel.NetmaskBits = core.Float64Ptr(float64(72.5))
	subnetModel.NetmaskIp4 = core.StringPtr("testString")
	subnetModel.NfsAccess = core.StringPtr("kDisabled")
	subnetModel.NfsAllSquash = core.BoolPtr(true)
	subnetModel.NfsRootSquash = core.BoolPtr(true)
	subnetModel.S3Access = core.StringPtr("kDisabled")
	subnetModel.SmbAccess = core.StringPtr("kDisabled")
	subnetModel.TenantID = core.StringPtr("testString")

	latencyThresholdsModel := new(backuprecoveryv1.LatencyThresholds)
	latencyThresholdsModel.ActiveTaskMsecs = core.Int64Ptr(int64(26))
	latencyThresholdsModel.NewTaskMsecs = core.Int64Ptr(int64(26))

	nasSourceParamsModel := new(backuprecoveryv1.NasSourceParams)
	nasSourceParamsModel.MaxParallelMetadataFetchFullPercentage = core.Float64Ptr(float64(72.5))
	nasSourceParamsModel.MaxParallelMetadataFetchIncrementalPercentage = core.Float64Ptr(float64(72.5))
	nasSourceParamsModel.MaxParallelReadWriteFullPercentage = core.Float64Ptr(float64(72.5))
	nasSourceParamsModel.MaxParallelReadWriteIncrementalPercentage = core.Float64Ptr(float64(72.5))

	storageArraySnapshotMaxSpaceConfigModel := new(backuprecoveryv1.StorageArraySnapshotMaxSpaceConfig)
	storageArraySnapshotMaxSpaceConfigModel.MaxSnapshotSpacePercentage = core.Float64Ptr(float64(72.5))

	maxSnapshotConfigModel := new(backuprecoveryv1.MaxSnapshotConfig)
	maxSnapshotConfigModel.MaxSnapshots = core.Float64Ptr(float64(72.5))

	maxSpaceConfigModel := new(backuprecoveryv1.MaxSpaceConfig)
	maxSpaceConfigModel.MaxSnapshotSpacePercentage = core.Float64Ptr(float64(72.5))

	storageArraySnapshotThrottlingPoliciesModel := new(backuprecoveryv1.StorageArraySnapshotThrottlingPolicies)
	storageArraySnapshotThrottlingPoliciesModel.ID = core.Int64Ptr(int64(26))
	storageArraySnapshotThrottlingPoliciesModel.IsMaxSnapshotsConfigEnabled = core.BoolPtr(true)
	storageArraySnapshotThrottlingPoliciesModel.IsMaxSpaceConfigEnabled = core.BoolPtr(true)
	storageArraySnapshotThrottlingPoliciesModel.MaxSnapshotConfig = maxSnapshotConfigModel
	storageArraySnapshotThrottlingPoliciesModel.MaxSpaceConfig = maxSpaceConfigModel

	storageArraySnapshotConfigModel := new(backuprecoveryv1.StorageArraySnapshotConfig)
	storageArraySnapshotConfigModel.IsMaxSnapshotsConfigEnabled = core.BoolPtr(true)
	storageArraySnapshotConfigModel.IsMaxSpaceConfigEnabled = core.BoolPtr(true)
	storageArraySnapshotConfigModel.StorageArraySnapshotMaxSpaceConfig = storageArraySnapshotMaxSpaceConfigModel
	storageArraySnapshotConfigModel.StorageArraySnapshotThrottlingPolicies = []backuprecoveryv1.StorageArraySnapshotThrottlingPolicies{*storageArraySnapshotThrottlingPoliciesModel}

	throttlingPolicyModel := new(backuprecoveryv1.ThrottlingPolicy)
	throttlingPolicyModel.EnforceMaxStreams = core.BoolPtr(true)
	throttlingPolicyModel.EnforceRegisteredSourceMaxBackups = core.BoolPtr(true)
	throttlingPolicyModel.IsEnabled = core.BoolPtr(true)
	throttlingPolicyModel.LatencyThresholds = latencyThresholdsModel
	throttlingPolicyModel.MaxConcurrentStreams = core.Float64Ptr(float64(72.5))
	throttlingPolicyModel.NasSourceParams = nasSourceParamsModel
	throttlingPolicyModel.RegisteredSourceMaxConcurrentBackups = core.Float64Ptr(float64(72.5))
	throttlingPolicyModel.StorageArraySnapshotConfig = storageArraySnapshotConfigModel

	throttlingPolicyOverridesModel := new(backuprecoveryv1.ThrottlingPolicyOverrides)
	throttlingPolicyOverridesModel.DatastoreID = core.Int64Ptr(int64(26))
	throttlingPolicyOverridesModel.DatastoreName = core.StringPtr("testString")
	throttlingPolicyOverridesModel.ThrottlingPolicy = throttlingPolicyModel

	registeredSourceVlanConfigModel := new(backuprecoveryv1.RegisteredSourceVlanConfig)
	registeredSourceVlanConfigModel.Vlan = core.Float64Ptr(float64(72.5))
	registeredSourceVlanConfigModel.DisableVlan = core.BoolPtr(true)
	registeredSourceVlanConfigModel.InterfaceName = core.StringPtr("testString")

	agentRegistrationInfoModel := new(backuprecoveryv1.AgentRegistrationInfo)
	agentRegistrationInfoModel.AccessInfo = agentAccessInfoModel
	agentRegistrationInfoModel.AllowedIpAddresses = []string{"testString"}
	agentRegistrationInfoModel.AuthenticationErrorMessage = core.StringPtr("testString")
	agentRegistrationInfoModel.AuthenticationStatus = core.StringPtr("kPending")
	agentRegistrationInfoModel.BlacklistedIpAddresses = []string{"testString"}
	agentRegistrationInfoModel.DeniedIpAddresses = []string{"testString"}
	agentRegistrationInfoModel.Environments = []string{"kPhysical"}
	agentRegistrationInfoModel.IsDbAuthenticated = core.BoolPtr(true)
	agentRegistrationInfoModel.IsStorageArraySnapshotEnabled = core.BoolPtr(true)
	agentRegistrationInfoModel.LinkVmsAcrossVcenter = core.BoolPtr(true)
	agentRegistrationInfoModel.MinimumFreeSpaceGB = core.Int64Ptr(int64(26))
	agentRegistrationInfoModel.MinimumFreeSpacePercent = core.Int64Ptr(int64(26))
	agentRegistrationInfoModel.Password = core.StringPtr("testString")
	agentRegistrationInfoModel.PhysicalParams = agentPhysicalParamsModel
	agentRegistrationInfoModel.ProgressMonitorPath = core.StringPtr("testString")
	agentRegistrationInfoModel.RefreshErrorMessage = core.StringPtr("testString")
	agentRegistrationInfoModel.RefreshTimeUsecs = core.Int64Ptr(int64(26))
	agentRegistrationInfoModel.RegisteredAppsInfo = []backuprecoveryv1.RegisteredAppInfo{*registeredAppInfoModel}
	agentRegistrationInfoModel.RegistrationTimeUsecs = core.Int64Ptr(int64(26))
	agentRegistrationInfoModel.Subnets = []backuprecoveryv1.Subnet{*subnetModel}
	agentRegistrationInfoModel.ThrottlingPolicy = throttlingPolicyModel
	agentRegistrationInfoModel.ThrottlingPolicyOverrides = []backuprecoveryv1.ThrottlingPolicyOverrides{*throttlingPolicyOverridesModel}
	agentRegistrationInfoModel.UseOAuthForExchangeOnline = core.BoolPtr(true)
	agentRegistrationInfoModel.UseVmBiosUUID = core.BoolPtr(true)
	agentRegistrationInfoModel.UserMessages = []string{"testString"}
	agentRegistrationInfoModel.Username = core.StringPtr("testString")
	agentRegistrationInfoModel.VlanParams = registeredSourceVlanConfigModel
	agentRegistrationInfoModel.WarningMessages = []string{"testString"}

	agentInformationModel := new(backuprecoveryv1.AgentInformation)
	agentInformationModel.CbmrVersion = core.StringPtr("testString")
	agentInformationModel.FileCbtInfo = cbtInfoModel
	agentInformationModel.HostType = core.StringPtr("kLinux")
	agentInformationModel.ID = core.Int64Ptr(int64(26))
	agentInformationModel.Name = core.StringPtr("testString")
	agentInformationModel.OracleMultiNodeChannelSupported = core.BoolPtr(true)
	agentInformationModel.RegistrationInfo = agentRegistrationInfoModel
	agentInformationModel.SourceSideDedupEnabled = core.BoolPtr(true)
	agentInformationModel.Status = core.StringPtr("kUnknown")
	agentInformationModel.StatusMessage = core.StringPtr("testString")
	agentInformationModel.Upgradability = core.StringPtr("kUpgradable")
	agentInformationModel.UpgradeStatus = core.StringPtr("kIdle")
	agentInformationModel.UpgradeStatusMessage = core.StringPtr("testString")
	agentInformationModel.Version = core.StringPtr("testString")
	agentInformationModel.VolCbtInfo = cbtInfoModel

	uniqueGlobalIdModel := new(backuprecoveryv1.UniqueGlobalID)
	uniqueGlobalIdModel.ClusterID = core.Int64Ptr(int64(26))
	uniqueGlobalIdModel.ClusterIncarnationID = core.Int64Ptr(int64(26))
	uniqueGlobalIdModel.ID = core.Int64Ptr(int64(26))

	clusterNetworkingEndpointModel := new(backuprecoveryv1.ClusterNetworkingEndpoint)
	clusterNetworkingEndpointModel.Fqdn = core.StringPtr("testString")
	clusterNetworkingEndpointModel.Ipv4Addr = core.StringPtr("testString")
	clusterNetworkingEndpointModel.Ipv6Addr = core.StringPtr("testString")

	clusterNetworkResourceInformationModel := new(backuprecoveryv1.ClusterNetworkResourceInformation)
	clusterNetworkResourceInformationModel.Endpoints = []backuprecoveryv1.ClusterNetworkingEndpoint{*clusterNetworkingEndpointModel}
	clusterNetworkResourceInformationModel.Type = core.StringPtr("testString")

	networkingInformationModel := new(backuprecoveryv1.NetworkingInformation)
	networkingInformationModel.ResourceVec = []backuprecoveryv1.ClusterNetworkResourceInformation{*clusterNetworkResourceInformationModel}

	physicalVolumeModel := new(backuprecoveryv1.PhysicalVolume)
	physicalVolumeModel.DevicePath = core.StringPtr("testString")
	physicalVolumeModel.Guid = core.StringPtr("testString")
	physicalVolumeModel.IsBootVolume = core.BoolPtr(true)
	physicalVolumeModel.IsExtendedAttributesSupported = core.BoolPtr(true)
	physicalVolumeModel.IsProtected = core.BoolPtr(true)
	physicalVolumeModel.IsSharedVolume = core.BoolPtr(true)
	physicalVolumeModel.Label = core.StringPtr("testString")
	physicalVolumeModel.LogicalSizeBytes = core.Float64Ptr(float64(72.5))
	physicalVolumeModel.MountPoints = []string{"testString"}
	physicalVolumeModel.MountType = core.StringPtr("testString")
	physicalVolumeModel.NetworkPath = core.StringPtr("testString")
	physicalVolumeModel.UsedSizeBytes = core.Float64Ptr(float64(72.5))

	vssWritersModel := new(backuprecoveryv1.VssWriters)
	vssWritersModel.IsWriterExcluded = core.BoolPtr(true)
	vssWritersModel.WriterName = core.BoolPtr(true)

	physicalProtectionSourceModel := new(backuprecoveryv1.PhysicalProtectionSource)
	physicalProtectionSourceModel.Agents = []backuprecoveryv1.AgentInformation{*agentInformationModel}
	physicalProtectionSourceModel.ClusterSourceType = core.StringPtr("testString")
	physicalProtectionSourceModel.HostName = core.StringPtr("testString")
	physicalProtectionSourceModel.HostType = core.StringPtr("kLinux")
	physicalProtectionSourceModel.ID = uniqueGlobalIdModel
	physicalProtectionSourceModel.IsProxyHost = core.BoolPtr(true)
	physicalProtectionSourceModel.MemorySizeBytes = core.Int64Ptr(int64(26))
	physicalProtectionSourceModel.Name = core.StringPtr("testString")
	physicalProtectionSourceModel.NetworkingInfo = networkingInformationModel
	physicalProtectionSourceModel.NumProcessors = core.Int64Ptr(int64(26))
	physicalProtectionSourceModel.OsName = core.StringPtr("testString")
	physicalProtectionSourceModel.Type = core.StringPtr("kGroup")
	physicalProtectionSourceModel.VcsVersion = core.StringPtr("testString")
	physicalProtectionSourceModel.Volumes = []backuprecoveryv1.PhysicalVolume{*physicalVolumeModel}
	physicalProtectionSourceModel.Vsswriters = []backuprecoveryv1.VssWriters{*vssWritersModel}

	vlanParametersModel := new(backuprecoveryv1.VlanParameters)
	vlanParametersModel.DisableVlan = core.BoolPtr(true)
	vlanParametersModel.InterfaceName = core.StringPtr("testString")
	vlanParametersModel.Vlan = core.Int64Ptr(int64(38))

	kubernetesLabelAttributeModel := new(backuprecoveryv1.KubernetesLabelAttribute)
	kubernetesLabelAttributeModel.ID = core.Int64Ptr(int64(26))
	kubernetesLabelAttributeModel.Name = core.StringPtr("testString")
	kubernetesLabelAttributeModel.UUID = core.StringPtr("testString")

	k8sLabelModel := new(backuprecoveryv1.K8sLabel)
	k8sLabelModel.Key = core.StringPtr("testString")
	k8sLabelModel.Value = core.StringPtr("testString")

	serviceAnnotationsEntryModel := new(backuprecoveryv1.ServiceAnnotationsEntry)
	serviceAnnotationsEntryModel.Key = core.StringPtr("testString")
	serviceAnnotationsEntryModel.Value = core.StringPtr("testString")

	kubernetesStorageClassInfoModel := new(backuprecoveryv1.KubernetesStorageClassInfo)
	kubernetesStorageClassInfoModel.Name = core.StringPtr("testString")
	kubernetesStorageClassInfoModel.Provisioner = core.StringPtr("testString")

	kubernetesServiceAnnotationObjectModel := new(backuprecoveryv1.KubernetesServiceAnnotationObject)
	kubernetesServiceAnnotationObjectModel.Key = core.StringPtr("testString")
	kubernetesServiceAnnotationObjectModel.Value = core.StringPtr("testString")

	vlanParamsModel := new(backuprecoveryv1.VlanParams)
	vlanParamsModel.DisableVlan = core.BoolPtr(true)
	vlanParamsModel.InterfaceName = core.StringPtr("testString")
	vlanParamsModel.VlanID = core.Int64Ptr(int64(38))

	kubernetesVlanInfoModel := new(backuprecoveryv1.KubernetesVlanInfo)
	kubernetesVlanInfoModel.ServiceAnnotations = []backuprecoveryv1.KubernetesServiceAnnotationObject{*kubernetesServiceAnnotationObjectModel}
	kubernetesVlanInfoModel.VlanParams = vlanParamsModel

	kubernetesProtectionSourceModel := new(backuprecoveryv1.KubernetesProtectionSource)
	kubernetesProtectionSourceModel.DatamoverImageLocation = core.StringPtr("testString")
	kubernetesProtectionSourceModel.DatamoverServiceType = core.Int64Ptr(int64(38))
	kubernetesProtectionSourceModel.DatamoverUpgradability = core.StringPtr("kCurrent")
	kubernetesProtectionSourceModel.DefaultVlanParams = vlanParametersModel
	kubernetesProtectionSourceModel.Description = core.StringPtr("testString")
	kubernetesProtectionSourceModel.Distribution = core.StringPtr("kMainline")
	kubernetesProtectionSourceModel.InitContainerImageLocation = core.StringPtr("testString")
	kubernetesProtectionSourceModel.LabelAttributes = []backuprecoveryv1.KubernetesLabelAttribute{*kubernetesLabelAttributeModel}
	kubernetesProtectionSourceModel.Name = core.StringPtr("testString")
	kubernetesProtectionSourceModel.PriorityClassName = core.StringPtr("testString")
	kubernetesProtectionSourceModel.ResourceAnnotationList = []backuprecoveryv1.K8sLabel{*k8sLabelModel}
	kubernetesProtectionSourceModel.ResourceLabelList = []backuprecoveryv1.K8sLabel{*k8sLabelModel}
	kubernetesProtectionSourceModel.SanField = []string{"testString"}
	kubernetesProtectionSourceModel.ServiceAnnotations = []backuprecoveryv1.ServiceAnnotationsEntry{*serviceAnnotationsEntryModel}
	kubernetesProtectionSourceModel.StorageClass = []backuprecoveryv1.KubernetesStorageClassInfo{*kubernetesStorageClassInfoModel}
	kubernetesProtectionSourceModel.Type = core.StringPtr("kCluster")
	kubernetesProtectionSourceModel.UUID = core.StringPtr("testString")
	kubernetesProtectionSourceModel.VeleroAwsPluginImageLocation = core.StringPtr("testString")
	kubernetesProtectionSourceModel.VeleroImageLocation = core.StringPtr("testString")
	kubernetesProtectionSourceModel.VeleroOpenshiftPluginImageLocation = core.StringPtr("testString")
	kubernetesProtectionSourceModel.VeleroUpgradability = core.StringPtr("testString")
	kubernetesProtectionSourceModel.VlanInfoVec = []backuprecoveryv1.KubernetesVlanInfo{*kubernetesVlanInfoModel}

	databaseFileInformationModel := new(backuprecoveryv1.DatabaseFileInformation)
	databaseFileInformationModel.FileType = core.StringPtr("kRows")
	databaseFileInformationModel.FullPath = core.StringPtr("testString")
	databaseFileInformationModel.SizeBytes = core.Int64Ptr(int64(26))

	sqlSourceIdModel := new(backuprecoveryv1.SQLSourceID)
	sqlSourceIdModel.CreatedDateMsecs = core.Int64Ptr(int64(26))
	sqlSourceIdModel.DatabaseID = core.Int64Ptr(int64(26))
	sqlSourceIdModel.InstanceID = core.StringPtr("testString")

	sqlServerInstanceVersionModel := new(backuprecoveryv1.SQLServerInstanceVersion)
	sqlServerInstanceVersionModel.Build = core.Float64Ptr(float64(72.5))
	sqlServerInstanceVersionModel.MajorVersion = core.Float64Ptr(float64(72.5))
	sqlServerInstanceVersionModel.MinorVersion = core.Float64Ptr(float64(72.5))
	sqlServerInstanceVersionModel.Revision = core.Float64Ptr(float64(72.5))
	sqlServerInstanceVersionModel.VersionString = core.Float64Ptr(float64(72.5))

	sqlProtectionSourceModel := new(backuprecoveryv1.SqlProtectionSource)
	sqlProtectionSourceModel.IsAvailableForVssBackup = core.BoolPtr(true)
	sqlProtectionSourceModel.CreatedTimestamp = core.StringPtr("testString")
	sqlProtectionSourceModel.DatabaseName = core.StringPtr("testString")
	sqlProtectionSourceModel.DbAagEntityID = core.Int64Ptr(int64(26))
	sqlProtectionSourceModel.DbAagName = core.StringPtr("testString")
	sqlProtectionSourceModel.DbCompatibilityLevel = core.Int64Ptr(int64(26))
	sqlProtectionSourceModel.DbFileGroups = []string{"testString"}
	sqlProtectionSourceModel.DbFiles = []backuprecoveryv1.DatabaseFileInformation{*databaseFileInformationModel}
	sqlProtectionSourceModel.DbOwnerUsername = core.StringPtr("testString")
	sqlProtectionSourceModel.DefaultDatabaseLocation = core.StringPtr("testString")
	sqlProtectionSourceModel.DefaultLogLocation = core.StringPtr("testString")
	sqlProtectionSourceModel.ID = sqlSourceIdModel
	sqlProtectionSourceModel.IsEncrypted = core.BoolPtr(true)
	sqlProtectionSourceModel.Name = core.StringPtr("testString")
	sqlProtectionSourceModel.OwnerID = core.Int64Ptr(int64(26))
	sqlProtectionSourceModel.RecoveryModel = core.StringPtr("kSimpleRecoveryModel")
	sqlProtectionSourceModel.SqlServerDbState = core.StringPtr("kOnline")
	sqlProtectionSourceModel.SqlServerInstanceVersion = sqlServerInstanceVersionModel
	sqlProtectionSourceModel.Type = core.StringPtr("kInstance")

	model := new(backuprecoveryv1.ProtectionSourceNode)
	model.ConnectionID = core.Int64Ptr(int64(26))
	model.ConnectorGroupID = core.Int64Ptr(int64(26))
	model.CustomName = core.StringPtr("testString")
	model.Environment = core.StringPtr("kPhysical")
	model.ID = core.Int64Ptr(int64(26))
	model.Name = core.StringPtr("testString")
	model.ParentID = core.Int64Ptr(int64(26))
	model.PhysicalProtectionSource = physicalProtectionSourceModel
	model.KubernetesProtectionSource = kubernetesProtectionSourceModel
	model.SqlProtectionSource = sqlProtectionSourceModel

	result, err := backuprecovery.DataSourceIbmBackupRecoveryRegistrationInfoProtectionSourceNodeToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryRegistrationInfoPhysicalProtectionSourceToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		cbtFileVersionModel := make(map[string]interface{})
		cbtFileVersionModel["build_ver"] = float64(72.5)
		cbtFileVersionModel["major_ver"] = float64(72.5)
		cbtFileVersionModel["minor_ver"] = float64(72.5)
		cbtFileVersionModel["revision_num"] = float64(72.5)

		cbtServiceStateModel := make(map[string]interface{})
		cbtServiceStateModel["name"] = "testString"
		cbtServiceStateModel["state"] = "testString"

		cbtInfoModel := make(map[string]interface{})
		cbtInfoModel["file_version"] = []map[string]interface{}{cbtFileVersionModel}
		cbtInfoModel["is_installed"] = true
		cbtInfoModel["reboot_status"] = "kRebooted"
		cbtInfoModel["service_state"] = []map[string]interface{}{cbtServiceStateModel}

		agentAccessInfoModel := make(map[string]interface{})
		agentAccessInfoModel["connection_id"] = int(26)
		agentAccessInfoModel["connector_group_id"] = int(26)
		agentAccessInfoModel["endpoint"] = "testString"
		agentAccessInfoModel["environment"] = "kPhysical"
		agentAccessInfoModel["id"] = int(26)
		agentAccessInfoModel["version"] = int(26)

		timeModel := make(map[string]interface{})
		timeModel["hour"] = int(38)
		timeModel["minute"] = int(38)

		dayTimeParamsModel := make(map[string]interface{})
		dayTimeParamsModel["day"] = "kSunday"
		dayTimeParamsModel["time"] = []map[string]interface{}{timeModel}

		dayTimeWindowModel := make(map[string]interface{})
		dayTimeWindowModel["end_time"] = []map[string]interface{}{dayTimeParamsModel}
		dayTimeWindowModel["start_time"] = []map[string]interface{}{dayTimeParamsModel}

		throttlingWindowModel := make(map[string]interface{})
		throttlingWindowModel["day_time_window"] = []map[string]interface{}{dayTimeWindowModel}
		throttlingWindowModel["threshold"] = int(26)

		throttlingConfigurationParamsModel := make(map[string]interface{})
		throttlingConfigurationParamsModel["fixed_threshold"] = int(26)
		throttlingConfigurationParamsModel["pattern_type"] = "kNoThrottling"
		throttlingConfigurationParamsModel["throttling_windows"] = []map[string]interface{}{throttlingWindowModel}

		throttlingConfigModel := make(map[string]interface{})
		throttlingConfigModel["cpu_throttling_config"] = []map[string]interface{}{throttlingConfigurationParamsModel}
		throttlingConfigModel["network_throttling_config"] = []map[string]interface{}{throttlingConfigurationParamsModel}

		agentPhysicalParamsModel := make(map[string]interface{})
		agentPhysicalParamsModel["applications"] = []string{"kSQL"}
		agentPhysicalParamsModel["password"] = "testString"
		agentPhysicalParamsModel["throttling_config"] = []map[string]interface{}{throttlingConfigModel}
		agentPhysicalParamsModel["username"] = "testString"

		hostSettingsCheckResultModel := make(map[string]interface{})
		hostSettingsCheckResultModel["check_type"] = "kIsAgentPortAccessible"
		hostSettingsCheckResultModel["result_type"] = "kPass"
		hostSettingsCheckResultModel["user_message"] = "testString"

		registeredAppInfoModel := make(map[string]interface{})
		registeredAppInfoModel["authentication_error_message"] = "testString"
		registeredAppInfoModel["authentication_status"] = "kPending"
		registeredAppInfoModel["environment"] = "kPhysical"
		registeredAppInfoModel["host_settings_check_results"] = []map[string]interface{}{hostSettingsCheckResultModel}
		registeredAppInfoModel["refresh_error_message"] = "testString"

		subnetModel := make(map[string]interface{})
		subnetModel["component"] = "testString"
		subnetModel["description"] = "testString"
		subnetModel["id"] = float64(72.5)
		subnetModel["ip"] = "testString"
		subnetModel["netmask_bits"] = float64(72.5)
		subnetModel["netmask_ip4"] = "testString"
		subnetModel["nfs_access"] = "kDisabled"
		subnetModel["nfs_all_squash"] = true
		subnetModel["nfs_root_squash"] = true
		subnetModel["s3_access"] = "kDisabled"
		subnetModel["smb_access"] = "kDisabled"
		subnetModel["tenant_id"] = "testString"

		latencyThresholdsModel := make(map[string]interface{})
		latencyThresholdsModel["active_task_msecs"] = int(26)
		latencyThresholdsModel["new_task_msecs"] = int(26)

		nasSourceParamsModel := make(map[string]interface{})
		nasSourceParamsModel["max_parallel_metadata_fetch_full_percentage"] = float64(72.5)
		nasSourceParamsModel["max_parallel_metadata_fetch_incremental_percentage"] = float64(72.5)
		nasSourceParamsModel["max_parallel_read_write_full_percentage"] = float64(72.5)
		nasSourceParamsModel["max_parallel_read_write_incremental_percentage"] = float64(72.5)

		storageArraySnapshotMaxSpaceConfigModel := make(map[string]interface{})
		storageArraySnapshotMaxSpaceConfigModel["max_snapshot_space_percentage"] = float64(72.5)

		maxSnapshotConfigModel := make(map[string]interface{})
		maxSnapshotConfigModel["max_snapshots"] = float64(72.5)

		maxSpaceConfigModel := make(map[string]interface{})
		maxSpaceConfigModel["max_snapshot_space_percentage"] = float64(72.5)

		storageArraySnapshotThrottlingPoliciesModel := make(map[string]interface{})
		storageArraySnapshotThrottlingPoliciesModel["id"] = int(26)
		storageArraySnapshotThrottlingPoliciesModel["is_max_snapshots_config_enabled"] = true
		storageArraySnapshotThrottlingPoliciesModel["is_max_space_config_enabled"] = true
		storageArraySnapshotThrottlingPoliciesModel["max_snapshot_config"] = []map[string]interface{}{maxSnapshotConfigModel}
		storageArraySnapshotThrottlingPoliciesModel["max_space_config"] = []map[string]interface{}{maxSpaceConfigModel}

		storageArraySnapshotConfigModel := make(map[string]interface{})
		storageArraySnapshotConfigModel["is_max_snapshots_config_enabled"] = true
		storageArraySnapshotConfigModel["is_max_space_config_enabled"] = true
		storageArraySnapshotConfigModel["storage_array_snapshot_max_space_config"] = []map[string]interface{}{storageArraySnapshotMaxSpaceConfigModel}
		storageArraySnapshotConfigModel["storage_array_snapshot_throttling_policies"] = []map[string]interface{}{storageArraySnapshotThrottlingPoliciesModel}

		throttlingPolicyModel := make(map[string]interface{})
		throttlingPolicyModel["enforce_max_streams"] = true
		throttlingPolicyModel["enforce_registered_source_max_backups"] = true
		throttlingPolicyModel["is_enabled"] = true
		throttlingPolicyModel["latency_thresholds"] = []map[string]interface{}{latencyThresholdsModel}
		throttlingPolicyModel["max_concurrent_streams"] = float64(72.5)
		throttlingPolicyModel["nas_source_params"] = []map[string]interface{}{nasSourceParamsModel}
		throttlingPolicyModel["registered_source_max_concurrent_backups"] = float64(72.5)
		throttlingPolicyModel["storage_array_snapshot_config"] = []map[string]interface{}{storageArraySnapshotConfigModel}

		throttlingPolicyOverridesModel := make(map[string]interface{})
		throttlingPolicyOverridesModel["datastore_id"] = int(26)
		throttlingPolicyOverridesModel["datastore_name"] = "testString"
		throttlingPolicyOverridesModel["throttling_policy"] = []map[string]interface{}{throttlingPolicyModel}

		registeredSourceVlanConfigModel := make(map[string]interface{})
		registeredSourceVlanConfigModel["vlan"] = float64(72.5)
		registeredSourceVlanConfigModel["disable_vlan"] = true
		registeredSourceVlanConfigModel["interface_name"] = "testString"

		agentRegistrationInfoModel := make(map[string]interface{})
		agentRegistrationInfoModel["access_info"] = []map[string]interface{}{agentAccessInfoModel}
		agentRegistrationInfoModel["allowed_ip_addresses"] = []string{"testString"}
		agentRegistrationInfoModel["authentication_error_message"] = "testString"
		agentRegistrationInfoModel["authentication_status"] = "kPending"
		agentRegistrationInfoModel["blacklisted_ip_addresses"] = []string{"testString"}
		agentRegistrationInfoModel["denied_ip_addresses"] = []string{"testString"}
		agentRegistrationInfoModel["environments"] = []string{"kPhysical"}
		agentRegistrationInfoModel["is_db_authenticated"] = true
		agentRegistrationInfoModel["is_storage_array_snapshot_enabled"] = true
		agentRegistrationInfoModel["link_vms_across_vcenter"] = true
		agentRegistrationInfoModel["minimum_free_space_gb"] = int(26)
		agentRegistrationInfoModel["minimum_free_space_percent"] = int(26)
		agentRegistrationInfoModel["password"] = "testString"
		agentRegistrationInfoModel["physical_params"] = []map[string]interface{}{agentPhysicalParamsModel}
		agentRegistrationInfoModel["progress_monitor_path"] = "testString"
		agentRegistrationInfoModel["refresh_error_message"] = "testString"
		agentRegistrationInfoModel["refresh_time_usecs"] = int(26)
		agentRegistrationInfoModel["registered_apps_info"] = []map[string]interface{}{registeredAppInfoModel}
		agentRegistrationInfoModel["registration_time_usecs"] = int(26)
		agentRegistrationInfoModel["subnets"] = []map[string]interface{}{subnetModel}
		agentRegistrationInfoModel["throttling_policy"] = []map[string]interface{}{throttlingPolicyModel}
		agentRegistrationInfoModel["throttling_policy_overrides"] = []map[string]interface{}{throttlingPolicyOverridesModel}
		agentRegistrationInfoModel["use_o_auth_for_exchange_online"] = true
		agentRegistrationInfoModel["use_vm_bios_uuid"] = true
		agentRegistrationInfoModel["user_messages"] = []string{"testString"}
		agentRegistrationInfoModel["username"] = "testString"
		agentRegistrationInfoModel["vlan_params"] = []map[string]interface{}{registeredSourceVlanConfigModel}
		agentRegistrationInfoModel["warning_messages"] = []string{"testString"}

		agentInformationModel := make(map[string]interface{})
		agentInformationModel["cbmr_version"] = "testString"
		agentInformationModel["file_cbt_info"] = []map[string]interface{}{cbtInfoModel}
		agentInformationModel["host_type"] = "kLinux"
		agentInformationModel["id"] = int(26)
		agentInformationModel["name"] = "testString"
		agentInformationModel["oracle_multi_node_channel_supported"] = true
		agentInformationModel["registration_info"] = []map[string]interface{}{agentRegistrationInfoModel}
		agentInformationModel["source_side_dedup_enabled"] = true
		agentInformationModel["status"] = "kUnknown"
		agentInformationModel["status_message"] = "testString"
		agentInformationModel["upgradability"] = "kUpgradable"
		agentInformationModel["upgrade_status"] = "kIdle"
		agentInformationModel["upgrade_status_message"] = "testString"
		agentInformationModel["version"] = "testString"
		agentInformationModel["vol_cbt_info"] = []map[string]interface{}{cbtInfoModel}

		uniqueGlobalIdModel := make(map[string]interface{})
		uniqueGlobalIdModel["cluster_id"] = int(26)
		uniqueGlobalIdModel["cluster_incarnation_id"] = int(26)
		uniqueGlobalIdModel["id"] = int(26)

		clusterNetworkingEndpointModel := make(map[string]interface{})
		clusterNetworkingEndpointModel["fqdn"] = "testString"
		clusterNetworkingEndpointModel["ipv4_addr"] = "testString"
		clusterNetworkingEndpointModel["ipv6_addr"] = "testString"

		clusterNetworkResourceInformationModel := make(map[string]interface{})
		clusterNetworkResourceInformationModel["endpoints"] = []map[string]interface{}{clusterNetworkingEndpointModel}
		clusterNetworkResourceInformationModel["type"] = "testString"

		networkingInformationModel := make(map[string]interface{})
		networkingInformationModel["resource_vec"] = []map[string]interface{}{clusterNetworkResourceInformationModel}

		physicalVolumeModel := make(map[string]interface{})
		physicalVolumeModel["device_path"] = "testString"
		physicalVolumeModel["guid"] = "testString"
		physicalVolumeModel["is_boot_volume"] = true
		physicalVolumeModel["is_extended_attributes_supported"] = true
		physicalVolumeModel["is_protected"] = true
		physicalVolumeModel["is_shared_volume"] = true
		physicalVolumeModel["label"] = "testString"
		physicalVolumeModel["logical_size_bytes"] = float64(72.5)
		physicalVolumeModel["mount_points"] = []string{"testString"}
		physicalVolumeModel["mount_type"] = "testString"
		physicalVolumeModel["network_path"] = "testString"
		physicalVolumeModel["used_size_bytes"] = float64(72.5)

		vssWritersModel := make(map[string]interface{})
		vssWritersModel["is_writer_excluded"] = true
		vssWritersModel["writer_name"] = true

		model := make(map[string]interface{})
		model["agents"] = []map[string]interface{}{agentInformationModel}
		model["cluster_source_type"] = "testString"
		model["host_name"] = "testString"
		model["host_type"] = "kLinux"
		model["id"] = []map[string]interface{}{uniqueGlobalIdModel}
		model["is_proxy_host"] = true
		model["memory_size_bytes"] = int(26)
		model["name"] = "testString"
		model["networking_info"] = []map[string]interface{}{networkingInformationModel}
		model["num_processors"] = int(26)
		model["os_name"] = "testString"
		model["type"] = "kGroup"
		model["vcs_version"] = "testString"
		model["volumes"] = []map[string]interface{}{physicalVolumeModel}
		model["vsswriters"] = []map[string]interface{}{vssWritersModel}

		assert.Equal(t, result, model)
	}

	cbtFileVersionModel := new(backuprecoveryv1.CbtFileVersion)
	cbtFileVersionModel.BuildVer = core.Float64Ptr(float64(72.5))
	cbtFileVersionModel.MajorVer = core.Float64Ptr(float64(72.5))
	cbtFileVersionModel.MinorVer = core.Float64Ptr(float64(72.5))
	cbtFileVersionModel.RevisionNum = core.Float64Ptr(float64(72.5))

	cbtServiceStateModel := new(backuprecoveryv1.CbtServiceState)
	cbtServiceStateModel.Name = core.StringPtr("testString")
	cbtServiceStateModel.State = core.StringPtr("testString")

	cbtInfoModel := new(backuprecoveryv1.CbtInfo)
	cbtInfoModel.FileVersion = cbtFileVersionModel
	cbtInfoModel.IsInstalled = core.BoolPtr(true)
	cbtInfoModel.RebootStatus = core.StringPtr("kRebooted")
	cbtInfoModel.ServiceState = cbtServiceStateModel

	agentAccessInfoModel := new(backuprecoveryv1.AgentAccessInfo)
	agentAccessInfoModel.ConnectionID = core.Int64Ptr(int64(26))
	agentAccessInfoModel.ConnectorGroupID = core.Int64Ptr(int64(26))
	agentAccessInfoModel.Endpoint = core.StringPtr("testString")
	agentAccessInfoModel.Environment = core.StringPtr("kPhysical")
	agentAccessInfoModel.ID = core.Int64Ptr(int64(26))
	agentAccessInfoModel.Version = core.Int64Ptr(int64(26))

	timeModel := new(backuprecoveryv1.Time)
	timeModel.Hour = core.Int64Ptr(int64(38))
	timeModel.Minute = core.Int64Ptr(int64(38))

	dayTimeParamsModel := new(backuprecoveryv1.DayTimeParams)
	dayTimeParamsModel.Day = core.StringPtr("kSunday")
	dayTimeParamsModel.Time = timeModel

	dayTimeWindowModel := new(backuprecoveryv1.DayTimeWindow)
	dayTimeWindowModel.EndTime = dayTimeParamsModel
	dayTimeWindowModel.StartTime = dayTimeParamsModel

	throttlingWindowModel := new(backuprecoveryv1.ThrottlingWindow)
	throttlingWindowModel.DayTimeWindow = dayTimeWindowModel
	throttlingWindowModel.Threshold = core.Int64Ptr(int64(26))

	throttlingConfigurationParamsModel := new(backuprecoveryv1.ThrottlingConfigurationParams)
	throttlingConfigurationParamsModel.FixedThreshold = core.Int64Ptr(int64(26))
	throttlingConfigurationParamsModel.PatternType = core.StringPtr("kNoThrottling")
	throttlingConfigurationParamsModel.ThrottlingWindows = []backuprecoveryv1.ThrottlingWindow{*throttlingWindowModel}

	throttlingConfigModel := new(backuprecoveryv1.ThrottlingConfig)
	throttlingConfigModel.CpuThrottlingConfig = throttlingConfigurationParamsModel
	throttlingConfigModel.NetworkThrottlingConfig = throttlingConfigurationParamsModel

	agentPhysicalParamsModel := new(backuprecoveryv1.AgentPhysicalParams)
	agentPhysicalParamsModel.Applications = []string{"kSQL"}
	agentPhysicalParamsModel.Password = core.StringPtr("testString")
	agentPhysicalParamsModel.ThrottlingConfig = throttlingConfigModel
	agentPhysicalParamsModel.Username = core.StringPtr("testString")

	hostSettingsCheckResultModel := new(backuprecoveryv1.HostSettingsCheckResult)
	hostSettingsCheckResultModel.CheckType = core.StringPtr("kIsAgentPortAccessible")
	hostSettingsCheckResultModel.ResultType = core.StringPtr("kPass")
	hostSettingsCheckResultModel.UserMessage = core.StringPtr("testString")

	registeredAppInfoModel := new(backuprecoveryv1.RegisteredAppInfo)
	registeredAppInfoModel.AuthenticationErrorMessage = core.StringPtr("testString")
	registeredAppInfoModel.AuthenticationStatus = core.StringPtr("kPending")
	registeredAppInfoModel.Environment = core.StringPtr("kPhysical")
	registeredAppInfoModel.HostSettingsCheckResults = []backuprecoveryv1.HostSettingsCheckResult{*hostSettingsCheckResultModel}
	registeredAppInfoModel.RefreshErrorMessage = core.StringPtr("testString")

	subnetModel := new(backuprecoveryv1.Subnet)
	subnetModel.Component = core.StringPtr("testString")
	subnetModel.Description = core.StringPtr("testString")
	subnetModel.ID = core.Float64Ptr(float64(72.5))
	subnetModel.Ip = core.StringPtr("testString")
	subnetModel.NetmaskBits = core.Float64Ptr(float64(72.5))
	subnetModel.NetmaskIp4 = core.StringPtr("testString")
	subnetModel.NfsAccess = core.StringPtr("kDisabled")
	subnetModel.NfsAllSquash = core.BoolPtr(true)
	subnetModel.NfsRootSquash = core.BoolPtr(true)
	subnetModel.S3Access = core.StringPtr("kDisabled")
	subnetModel.SmbAccess = core.StringPtr("kDisabled")
	subnetModel.TenantID = core.StringPtr("testString")

	latencyThresholdsModel := new(backuprecoveryv1.LatencyThresholds)
	latencyThresholdsModel.ActiveTaskMsecs = core.Int64Ptr(int64(26))
	latencyThresholdsModel.NewTaskMsecs = core.Int64Ptr(int64(26))

	nasSourceParamsModel := new(backuprecoveryv1.NasSourceParams)
	nasSourceParamsModel.MaxParallelMetadataFetchFullPercentage = core.Float64Ptr(float64(72.5))
	nasSourceParamsModel.MaxParallelMetadataFetchIncrementalPercentage = core.Float64Ptr(float64(72.5))
	nasSourceParamsModel.MaxParallelReadWriteFullPercentage = core.Float64Ptr(float64(72.5))
	nasSourceParamsModel.MaxParallelReadWriteIncrementalPercentage = core.Float64Ptr(float64(72.5))

	storageArraySnapshotMaxSpaceConfigModel := new(backuprecoveryv1.StorageArraySnapshotMaxSpaceConfig)
	storageArraySnapshotMaxSpaceConfigModel.MaxSnapshotSpacePercentage = core.Float64Ptr(float64(72.5))

	maxSnapshotConfigModel := new(backuprecoveryv1.MaxSnapshotConfig)
	maxSnapshotConfigModel.MaxSnapshots = core.Float64Ptr(float64(72.5))

	maxSpaceConfigModel := new(backuprecoveryv1.MaxSpaceConfig)
	maxSpaceConfigModel.MaxSnapshotSpacePercentage = core.Float64Ptr(float64(72.5))

	storageArraySnapshotThrottlingPoliciesModel := new(backuprecoveryv1.StorageArraySnapshotThrottlingPolicies)
	storageArraySnapshotThrottlingPoliciesModel.ID = core.Int64Ptr(int64(26))
	storageArraySnapshotThrottlingPoliciesModel.IsMaxSnapshotsConfigEnabled = core.BoolPtr(true)
	storageArraySnapshotThrottlingPoliciesModel.IsMaxSpaceConfigEnabled = core.BoolPtr(true)
	storageArraySnapshotThrottlingPoliciesModel.MaxSnapshotConfig = maxSnapshotConfigModel
	storageArraySnapshotThrottlingPoliciesModel.MaxSpaceConfig = maxSpaceConfigModel

	storageArraySnapshotConfigModel := new(backuprecoveryv1.StorageArraySnapshotConfig)
	storageArraySnapshotConfigModel.IsMaxSnapshotsConfigEnabled = core.BoolPtr(true)
	storageArraySnapshotConfigModel.IsMaxSpaceConfigEnabled = core.BoolPtr(true)
	storageArraySnapshotConfigModel.StorageArraySnapshotMaxSpaceConfig = storageArraySnapshotMaxSpaceConfigModel
	storageArraySnapshotConfigModel.StorageArraySnapshotThrottlingPolicies = []backuprecoveryv1.StorageArraySnapshotThrottlingPolicies{*storageArraySnapshotThrottlingPoliciesModel}

	throttlingPolicyModel := new(backuprecoveryv1.ThrottlingPolicy)
	throttlingPolicyModel.EnforceMaxStreams = core.BoolPtr(true)
	throttlingPolicyModel.EnforceRegisteredSourceMaxBackups = core.BoolPtr(true)
	throttlingPolicyModel.IsEnabled = core.BoolPtr(true)
	throttlingPolicyModel.LatencyThresholds = latencyThresholdsModel
	throttlingPolicyModel.MaxConcurrentStreams = core.Float64Ptr(float64(72.5))
	throttlingPolicyModel.NasSourceParams = nasSourceParamsModel
	throttlingPolicyModel.RegisteredSourceMaxConcurrentBackups = core.Float64Ptr(float64(72.5))
	throttlingPolicyModel.StorageArraySnapshotConfig = storageArraySnapshotConfigModel

	throttlingPolicyOverridesModel := new(backuprecoveryv1.ThrottlingPolicyOverrides)
	throttlingPolicyOverridesModel.DatastoreID = core.Int64Ptr(int64(26))
	throttlingPolicyOverridesModel.DatastoreName = core.StringPtr("testString")
	throttlingPolicyOverridesModel.ThrottlingPolicy = throttlingPolicyModel

	registeredSourceVlanConfigModel := new(backuprecoveryv1.RegisteredSourceVlanConfig)
	registeredSourceVlanConfigModel.Vlan = core.Float64Ptr(float64(72.5))
	registeredSourceVlanConfigModel.DisableVlan = core.BoolPtr(true)
	registeredSourceVlanConfigModel.InterfaceName = core.StringPtr("testString")

	agentRegistrationInfoModel := new(backuprecoveryv1.AgentRegistrationInfo)
	agentRegistrationInfoModel.AccessInfo = agentAccessInfoModel
	agentRegistrationInfoModel.AllowedIpAddresses = []string{"testString"}
	agentRegistrationInfoModel.AuthenticationErrorMessage = core.StringPtr("testString")
	agentRegistrationInfoModel.AuthenticationStatus = core.StringPtr("kPending")
	agentRegistrationInfoModel.BlacklistedIpAddresses = []string{"testString"}
	agentRegistrationInfoModel.DeniedIpAddresses = []string{"testString"}
	agentRegistrationInfoModel.Environments = []string{"kPhysical"}
	agentRegistrationInfoModel.IsDbAuthenticated = core.BoolPtr(true)
	agentRegistrationInfoModel.IsStorageArraySnapshotEnabled = core.BoolPtr(true)
	agentRegistrationInfoModel.LinkVmsAcrossVcenter = core.BoolPtr(true)
	agentRegistrationInfoModel.MinimumFreeSpaceGB = core.Int64Ptr(int64(26))
	agentRegistrationInfoModel.MinimumFreeSpacePercent = core.Int64Ptr(int64(26))
	agentRegistrationInfoModel.Password = core.StringPtr("testString")
	agentRegistrationInfoModel.PhysicalParams = agentPhysicalParamsModel
	agentRegistrationInfoModel.ProgressMonitorPath = core.StringPtr("testString")
	agentRegistrationInfoModel.RefreshErrorMessage = core.StringPtr("testString")
	agentRegistrationInfoModel.RefreshTimeUsecs = core.Int64Ptr(int64(26))
	agentRegistrationInfoModel.RegisteredAppsInfo = []backuprecoveryv1.RegisteredAppInfo{*registeredAppInfoModel}
	agentRegistrationInfoModel.RegistrationTimeUsecs = core.Int64Ptr(int64(26))
	agentRegistrationInfoModel.Subnets = []backuprecoveryv1.Subnet{*subnetModel}
	agentRegistrationInfoModel.ThrottlingPolicy = throttlingPolicyModel
	agentRegistrationInfoModel.ThrottlingPolicyOverrides = []backuprecoveryv1.ThrottlingPolicyOverrides{*throttlingPolicyOverridesModel}
	agentRegistrationInfoModel.UseOAuthForExchangeOnline = core.BoolPtr(true)
	agentRegistrationInfoModel.UseVmBiosUUID = core.BoolPtr(true)
	agentRegistrationInfoModel.UserMessages = []string{"testString"}
	agentRegistrationInfoModel.Username = core.StringPtr("testString")
	agentRegistrationInfoModel.VlanParams = registeredSourceVlanConfigModel
	agentRegistrationInfoModel.WarningMessages = []string{"testString"}

	agentInformationModel := new(backuprecoveryv1.AgentInformation)
	agentInformationModel.CbmrVersion = core.StringPtr("testString")
	agentInformationModel.FileCbtInfo = cbtInfoModel
	agentInformationModel.HostType = core.StringPtr("kLinux")
	agentInformationModel.ID = core.Int64Ptr(int64(26))
	agentInformationModel.Name = core.StringPtr("testString")
	agentInformationModel.OracleMultiNodeChannelSupported = core.BoolPtr(true)
	agentInformationModel.RegistrationInfo = agentRegistrationInfoModel
	agentInformationModel.SourceSideDedupEnabled = core.BoolPtr(true)
	agentInformationModel.Status = core.StringPtr("kUnknown")
	agentInformationModel.StatusMessage = core.StringPtr("testString")
	agentInformationModel.Upgradability = core.StringPtr("kUpgradable")
	agentInformationModel.UpgradeStatus = core.StringPtr("kIdle")
	agentInformationModel.UpgradeStatusMessage = core.StringPtr("testString")
	agentInformationModel.Version = core.StringPtr("testString")
	agentInformationModel.VolCbtInfo = cbtInfoModel

	uniqueGlobalIdModel := new(backuprecoveryv1.UniqueGlobalID)
	uniqueGlobalIdModel.ClusterID = core.Int64Ptr(int64(26))
	uniqueGlobalIdModel.ClusterIncarnationID = core.Int64Ptr(int64(26))
	uniqueGlobalIdModel.ID = core.Int64Ptr(int64(26))

	clusterNetworkingEndpointModel := new(backuprecoveryv1.ClusterNetworkingEndpoint)
	clusterNetworkingEndpointModel.Fqdn = core.StringPtr("testString")
	clusterNetworkingEndpointModel.Ipv4Addr = core.StringPtr("testString")
	clusterNetworkingEndpointModel.Ipv6Addr = core.StringPtr("testString")

	clusterNetworkResourceInformationModel := new(backuprecoveryv1.ClusterNetworkResourceInformation)
	clusterNetworkResourceInformationModel.Endpoints = []backuprecoveryv1.ClusterNetworkingEndpoint{*clusterNetworkingEndpointModel}
	clusterNetworkResourceInformationModel.Type = core.StringPtr("testString")

	networkingInformationModel := new(backuprecoveryv1.NetworkingInformation)
	networkingInformationModel.ResourceVec = []backuprecoveryv1.ClusterNetworkResourceInformation{*clusterNetworkResourceInformationModel}

	physicalVolumeModel := new(backuprecoveryv1.PhysicalVolume)
	physicalVolumeModel.DevicePath = core.StringPtr("testString")
	physicalVolumeModel.Guid = core.StringPtr("testString")
	physicalVolumeModel.IsBootVolume = core.BoolPtr(true)
	physicalVolumeModel.IsExtendedAttributesSupported = core.BoolPtr(true)
	physicalVolumeModel.IsProtected = core.BoolPtr(true)
	physicalVolumeModel.IsSharedVolume = core.BoolPtr(true)
	physicalVolumeModel.Label = core.StringPtr("testString")
	physicalVolumeModel.LogicalSizeBytes = core.Float64Ptr(float64(72.5))
	physicalVolumeModel.MountPoints = []string{"testString"}
	physicalVolumeModel.MountType = core.StringPtr("testString")
	physicalVolumeModel.NetworkPath = core.StringPtr("testString")
	physicalVolumeModel.UsedSizeBytes = core.Float64Ptr(float64(72.5))

	vssWritersModel := new(backuprecoveryv1.VssWriters)
	vssWritersModel.IsWriterExcluded = core.BoolPtr(true)
	vssWritersModel.WriterName = core.BoolPtr(true)

	model := new(backuprecoveryv1.PhysicalProtectionSource)
	model.Agents = []backuprecoveryv1.AgentInformation{*agentInformationModel}
	model.ClusterSourceType = core.StringPtr("testString")
	model.HostName = core.StringPtr("testString")
	model.HostType = core.StringPtr("kLinux")
	model.ID = uniqueGlobalIdModel
	model.IsProxyHost = core.BoolPtr(true)
	model.MemorySizeBytes = core.Int64Ptr(int64(26))
	model.Name = core.StringPtr("testString")
	model.NetworkingInfo = networkingInformationModel
	model.NumProcessors = core.Int64Ptr(int64(26))
	model.OsName = core.StringPtr("testString")
	model.Type = core.StringPtr("kGroup")
	model.VcsVersion = core.StringPtr("testString")
	model.Volumes = []backuprecoveryv1.PhysicalVolume{*physicalVolumeModel}
	model.Vsswriters = []backuprecoveryv1.VssWriters{*vssWritersModel}

	result, err := backuprecovery.DataSourceIbmBackupRecoveryRegistrationInfoPhysicalProtectionSourceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryRegistrationInfoAgentInformationToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		cbtFileVersionModel := make(map[string]interface{})
		cbtFileVersionModel["build_ver"] = float64(72.5)
		cbtFileVersionModel["major_ver"] = float64(72.5)
		cbtFileVersionModel["minor_ver"] = float64(72.5)
		cbtFileVersionModel["revision_num"] = float64(72.5)

		cbtServiceStateModel := make(map[string]interface{})
		cbtServiceStateModel["name"] = "testString"
		cbtServiceStateModel["state"] = "testString"

		cbtInfoModel := make(map[string]interface{})
		cbtInfoModel["file_version"] = []map[string]interface{}{cbtFileVersionModel}
		cbtInfoModel["is_installed"] = true
		cbtInfoModel["reboot_status"] = "kRebooted"
		cbtInfoModel["service_state"] = []map[string]interface{}{cbtServiceStateModel}

		agentAccessInfoModel := make(map[string]interface{})
		agentAccessInfoModel["connection_id"] = int(26)
		agentAccessInfoModel["connector_group_id"] = int(26)
		agentAccessInfoModel["endpoint"] = "testString"
		agentAccessInfoModel["environment"] = "kPhysical"
		agentAccessInfoModel["id"] = int(26)
		agentAccessInfoModel["version"] = int(26)

		timeModel := make(map[string]interface{})
		timeModel["hour"] = int(38)
		timeModel["minute"] = int(38)

		dayTimeParamsModel := make(map[string]interface{})
		dayTimeParamsModel["day"] = "kSunday"
		dayTimeParamsModel["time"] = []map[string]interface{}{timeModel}

		dayTimeWindowModel := make(map[string]interface{})
		dayTimeWindowModel["end_time"] = []map[string]interface{}{dayTimeParamsModel}
		dayTimeWindowModel["start_time"] = []map[string]interface{}{dayTimeParamsModel}

		throttlingWindowModel := make(map[string]interface{})
		throttlingWindowModel["day_time_window"] = []map[string]interface{}{dayTimeWindowModel}
		throttlingWindowModel["threshold"] = int(26)

		throttlingConfigurationParamsModel := make(map[string]interface{})
		throttlingConfigurationParamsModel["fixed_threshold"] = int(26)
		throttlingConfigurationParamsModel["pattern_type"] = "kNoThrottling"
		throttlingConfigurationParamsModel["throttling_windows"] = []map[string]interface{}{throttlingWindowModel}

		throttlingConfigModel := make(map[string]interface{})
		throttlingConfigModel["cpu_throttling_config"] = []map[string]interface{}{throttlingConfigurationParamsModel}
		throttlingConfigModel["network_throttling_config"] = []map[string]interface{}{throttlingConfigurationParamsModel}

		agentPhysicalParamsModel := make(map[string]interface{})
		agentPhysicalParamsModel["applications"] = []string{"kSQL"}
		agentPhysicalParamsModel["password"] = "testString"
		agentPhysicalParamsModel["throttling_config"] = []map[string]interface{}{throttlingConfigModel}
		agentPhysicalParamsModel["username"] = "testString"

		hostSettingsCheckResultModel := make(map[string]interface{})
		hostSettingsCheckResultModel["check_type"] = "kIsAgentPortAccessible"
		hostSettingsCheckResultModel["result_type"] = "kPass"
		hostSettingsCheckResultModel["user_message"] = "testString"

		registeredAppInfoModel := make(map[string]interface{})
		registeredAppInfoModel["authentication_error_message"] = "testString"
		registeredAppInfoModel["authentication_status"] = "kPending"
		registeredAppInfoModel["environment"] = "kPhysical"
		registeredAppInfoModel["host_settings_check_results"] = []map[string]interface{}{hostSettingsCheckResultModel}
		registeredAppInfoModel["refresh_error_message"] = "testString"

		subnetModel := make(map[string]interface{})
		subnetModel["component"] = "testString"
		subnetModel["description"] = "testString"
		subnetModel["id"] = float64(72.5)
		subnetModel["ip"] = "testString"
		subnetModel["netmask_bits"] = float64(72.5)
		subnetModel["netmask_ip4"] = "testString"
		subnetModel["nfs_access"] = "kDisabled"
		subnetModel["nfs_all_squash"] = true
		subnetModel["nfs_root_squash"] = true
		subnetModel["s3_access"] = "kDisabled"
		subnetModel["smb_access"] = "kDisabled"
		subnetModel["tenant_id"] = "testString"

		latencyThresholdsModel := make(map[string]interface{})
		latencyThresholdsModel["active_task_msecs"] = int(26)
		latencyThresholdsModel["new_task_msecs"] = int(26)

		nasSourceParamsModel := make(map[string]interface{})
		nasSourceParamsModel["max_parallel_metadata_fetch_full_percentage"] = float64(72.5)
		nasSourceParamsModel["max_parallel_metadata_fetch_incremental_percentage"] = float64(72.5)
		nasSourceParamsModel["max_parallel_read_write_full_percentage"] = float64(72.5)
		nasSourceParamsModel["max_parallel_read_write_incremental_percentage"] = float64(72.5)

		storageArraySnapshotMaxSpaceConfigModel := make(map[string]interface{})
		storageArraySnapshotMaxSpaceConfigModel["max_snapshot_space_percentage"] = float64(72.5)

		maxSnapshotConfigModel := make(map[string]interface{})
		maxSnapshotConfigModel["max_snapshots"] = float64(72.5)

		maxSpaceConfigModel := make(map[string]interface{})
		maxSpaceConfigModel["max_snapshot_space_percentage"] = float64(72.5)

		storageArraySnapshotThrottlingPoliciesModel := make(map[string]interface{})
		storageArraySnapshotThrottlingPoliciesModel["id"] = int(26)
		storageArraySnapshotThrottlingPoliciesModel["is_max_snapshots_config_enabled"] = true
		storageArraySnapshotThrottlingPoliciesModel["is_max_space_config_enabled"] = true
		storageArraySnapshotThrottlingPoliciesModel["max_snapshot_config"] = []map[string]interface{}{maxSnapshotConfigModel}
		storageArraySnapshotThrottlingPoliciesModel["max_space_config"] = []map[string]interface{}{maxSpaceConfigModel}

		storageArraySnapshotConfigModel := make(map[string]interface{})
		storageArraySnapshotConfigModel["is_max_snapshots_config_enabled"] = true
		storageArraySnapshotConfigModel["is_max_space_config_enabled"] = true
		storageArraySnapshotConfigModel["storage_array_snapshot_max_space_config"] = []map[string]interface{}{storageArraySnapshotMaxSpaceConfigModel}
		storageArraySnapshotConfigModel["storage_array_snapshot_throttling_policies"] = []map[string]interface{}{storageArraySnapshotThrottlingPoliciesModel}

		throttlingPolicyModel := make(map[string]interface{})
		throttlingPolicyModel["enforce_max_streams"] = true
		throttlingPolicyModel["enforce_registered_source_max_backups"] = true
		throttlingPolicyModel["is_enabled"] = true
		throttlingPolicyModel["latency_thresholds"] = []map[string]interface{}{latencyThresholdsModel}
		throttlingPolicyModel["max_concurrent_streams"] = float64(72.5)
		throttlingPolicyModel["nas_source_params"] = []map[string]interface{}{nasSourceParamsModel}
		throttlingPolicyModel["registered_source_max_concurrent_backups"] = float64(72.5)
		throttlingPolicyModel["storage_array_snapshot_config"] = []map[string]interface{}{storageArraySnapshotConfigModel}

		throttlingPolicyOverridesModel := make(map[string]interface{})
		throttlingPolicyOverridesModel["datastore_id"] = int(26)
		throttlingPolicyOverridesModel["datastore_name"] = "testString"
		throttlingPolicyOverridesModel["throttling_policy"] = []map[string]interface{}{throttlingPolicyModel}

		registeredSourceVlanConfigModel := make(map[string]interface{})
		registeredSourceVlanConfigModel["vlan"] = float64(72.5)
		registeredSourceVlanConfigModel["disable_vlan"] = true
		registeredSourceVlanConfigModel["interface_name"] = "testString"

		agentRegistrationInfoModel := make(map[string]interface{})
		agentRegistrationInfoModel["access_info"] = []map[string]interface{}{agentAccessInfoModel}
		agentRegistrationInfoModel["allowed_ip_addresses"] = []string{"testString"}
		agentRegistrationInfoModel["authentication_error_message"] = "testString"
		agentRegistrationInfoModel["authentication_status"] = "kPending"
		agentRegistrationInfoModel["blacklisted_ip_addresses"] = []string{"testString"}
		agentRegistrationInfoModel["denied_ip_addresses"] = []string{"testString"}
		agentRegistrationInfoModel["environments"] = []string{"kPhysical"}
		agentRegistrationInfoModel["is_db_authenticated"] = true
		agentRegistrationInfoModel["is_storage_array_snapshot_enabled"] = true
		agentRegistrationInfoModel["link_vms_across_vcenter"] = true
		agentRegistrationInfoModel["minimum_free_space_gb"] = int(26)
		agentRegistrationInfoModel["minimum_free_space_percent"] = int(26)
		agentRegistrationInfoModel["password"] = "testString"
		agentRegistrationInfoModel["physical_params"] = []map[string]interface{}{agentPhysicalParamsModel}
		agentRegistrationInfoModel["progress_monitor_path"] = "testString"
		agentRegistrationInfoModel["refresh_error_message"] = "testString"
		agentRegistrationInfoModel["refresh_time_usecs"] = int(26)
		agentRegistrationInfoModel["registered_apps_info"] = []map[string]interface{}{registeredAppInfoModel}
		agentRegistrationInfoModel["registration_time_usecs"] = int(26)
		agentRegistrationInfoModel["subnets"] = []map[string]interface{}{subnetModel}
		agentRegistrationInfoModel["throttling_policy"] = []map[string]interface{}{throttlingPolicyModel}
		agentRegistrationInfoModel["throttling_policy_overrides"] = []map[string]interface{}{throttlingPolicyOverridesModel}
		agentRegistrationInfoModel["use_o_auth_for_exchange_online"] = true
		agentRegistrationInfoModel["use_vm_bios_uuid"] = true
		agentRegistrationInfoModel["user_messages"] = []string{"testString"}
		agentRegistrationInfoModel["username"] = "testString"
		agentRegistrationInfoModel["vlan_params"] = []map[string]interface{}{registeredSourceVlanConfigModel}
		agentRegistrationInfoModel["warning_messages"] = []string{"testString"}

		model := make(map[string]interface{})
		model["cbmr_version"] = "testString"
		model["file_cbt_info"] = []map[string]interface{}{cbtInfoModel}
		model["host_type"] = "kLinux"
		model["id"] = int(26)
		model["name"] = "testString"
		model["oracle_multi_node_channel_supported"] = true
		model["registration_info"] = []map[string]interface{}{agentRegistrationInfoModel}
		model["source_side_dedup_enabled"] = true
		model["status"] = "kUnknown"
		model["status_message"] = "testString"
		model["upgradability"] = "kUpgradable"
		model["upgrade_status"] = "kIdle"
		model["upgrade_status_message"] = "testString"
		model["version"] = "testString"
		model["vol_cbt_info"] = []map[string]interface{}{cbtInfoModel}

		assert.Equal(t, result, model)
	}

	cbtFileVersionModel := new(backuprecoveryv1.CbtFileVersion)
	cbtFileVersionModel.BuildVer = core.Float64Ptr(float64(72.5))
	cbtFileVersionModel.MajorVer = core.Float64Ptr(float64(72.5))
	cbtFileVersionModel.MinorVer = core.Float64Ptr(float64(72.5))
	cbtFileVersionModel.RevisionNum = core.Float64Ptr(float64(72.5))

	cbtServiceStateModel := new(backuprecoveryv1.CbtServiceState)
	cbtServiceStateModel.Name = core.StringPtr("testString")
	cbtServiceStateModel.State = core.StringPtr("testString")

	cbtInfoModel := new(backuprecoveryv1.CbtInfo)
	cbtInfoModel.FileVersion = cbtFileVersionModel
	cbtInfoModel.IsInstalled = core.BoolPtr(true)
	cbtInfoModel.RebootStatus = core.StringPtr("kRebooted")
	cbtInfoModel.ServiceState = cbtServiceStateModel

	agentAccessInfoModel := new(backuprecoveryv1.AgentAccessInfo)
	agentAccessInfoModel.ConnectionID = core.Int64Ptr(int64(26))
	agentAccessInfoModel.ConnectorGroupID = core.Int64Ptr(int64(26))
	agentAccessInfoModel.Endpoint = core.StringPtr("testString")
	agentAccessInfoModel.Environment = core.StringPtr("kPhysical")
	agentAccessInfoModel.ID = core.Int64Ptr(int64(26))
	agentAccessInfoModel.Version = core.Int64Ptr(int64(26))

	timeModel := new(backuprecoveryv1.Time)
	timeModel.Hour = core.Int64Ptr(int64(38))
	timeModel.Minute = core.Int64Ptr(int64(38))

	dayTimeParamsModel := new(backuprecoveryv1.DayTimeParams)
	dayTimeParamsModel.Day = core.StringPtr("kSunday")
	dayTimeParamsModel.Time = timeModel

	dayTimeWindowModel := new(backuprecoveryv1.DayTimeWindow)
	dayTimeWindowModel.EndTime = dayTimeParamsModel
	dayTimeWindowModel.StartTime = dayTimeParamsModel

	throttlingWindowModel := new(backuprecoveryv1.ThrottlingWindow)
	throttlingWindowModel.DayTimeWindow = dayTimeWindowModel
	throttlingWindowModel.Threshold = core.Int64Ptr(int64(26))

	throttlingConfigurationParamsModel := new(backuprecoveryv1.ThrottlingConfigurationParams)
	throttlingConfigurationParamsModel.FixedThreshold = core.Int64Ptr(int64(26))
	throttlingConfigurationParamsModel.PatternType = core.StringPtr("kNoThrottling")
	throttlingConfigurationParamsModel.ThrottlingWindows = []backuprecoveryv1.ThrottlingWindow{*throttlingWindowModel}

	throttlingConfigModel := new(backuprecoveryv1.ThrottlingConfig)
	throttlingConfigModel.CpuThrottlingConfig = throttlingConfigurationParamsModel
	throttlingConfigModel.NetworkThrottlingConfig = throttlingConfigurationParamsModel

	agentPhysicalParamsModel := new(backuprecoveryv1.AgentPhysicalParams)
	agentPhysicalParamsModel.Applications = []string{"kSQL"}
	agentPhysicalParamsModel.Password = core.StringPtr("testString")
	agentPhysicalParamsModel.ThrottlingConfig = throttlingConfigModel
	agentPhysicalParamsModel.Username = core.StringPtr("testString")

	hostSettingsCheckResultModel := new(backuprecoveryv1.HostSettingsCheckResult)
	hostSettingsCheckResultModel.CheckType = core.StringPtr("kIsAgentPortAccessible")
	hostSettingsCheckResultModel.ResultType = core.StringPtr("kPass")
	hostSettingsCheckResultModel.UserMessage = core.StringPtr("testString")

	registeredAppInfoModel := new(backuprecoveryv1.RegisteredAppInfo)
	registeredAppInfoModel.AuthenticationErrorMessage = core.StringPtr("testString")
	registeredAppInfoModel.AuthenticationStatus = core.StringPtr("kPending")
	registeredAppInfoModel.Environment = core.StringPtr("kPhysical")
	registeredAppInfoModel.HostSettingsCheckResults = []backuprecoveryv1.HostSettingsCheckResult{*hostSettingsCheckResultModel}
	registeredAppInfoModel.RefreshErrorMessage = core.StringPtr("testString")

	subnetModel := new(backuprecoveryv1.Subnet)
	subnetModel.Component = core.StringPtr("testString")
	subnetModel.Description = core.StringPtr("testString")
	subnetModel.ID = core.Float64Ptr(float64(72.5))
	subnetModel.Ip = core.StringPtr("testString")
	subnetModel.NetmaskBits = core.Float64Ptr(float64(72.5))
	subnetModel.NetmaskIp4 = core.StringPtr("testString")
	subnetModel.NfsAccess = core.StringPtr("kDisabled")
	subnetModel.NfsAllSquash = core.BoolPtr(true)
	subnetModel.NfsRootSquash = core.BoolPtr(true)
	subnetModel.S3Access = core.StringPtr("kDisabled")
	subnetModel.SmbAccess = core.StringPtr("kDisabled")
	subnetModel.TenantID = core.StringPtr("testString")

	latencyThresholdsModel := new(backuprecoveryv1.LatencyThresholds)
	latencyThresholdsModel.ActiveTaskMsecs = core.Int64Ptr(int64(26))
	latencyThresholdsModel.NewTaskMsecs = core.Int64Ptr(int64(26))

	nasSourceParamsModel := new(backuprecoveryv1.NasSourceParams)
	nasSourceParamsModel.MaxParallelMetadataFetchFullPercentage = core.Float64Ptr(float64(72.5))
	nasSourceParamsModel.MaxParallelMetadataFetchIncrementalPercentage = core.Float64Ptr(float64(72.5))
	nasSourceParamsModel.MaxParallelReadWriteFullPercentage = core.Float64Ptr(float64(72.5))
	nasSourceParamsModel.MaxParallelReadWriteIncrementalPercentage = core.Float64Ptr(float64(72.5))

	storageArraySnapshotMaxSpaceConfigModel := new(backuprecoveryv1.StorageArraySnapshotMaxSpaceConfig)
	storageArraySnapshotMaxSpaceConfigModel.MaxSnapshotSpacePercentage = core.Float64Ptr(float64(72.5))

	maxSnapshotConfigModel := new(backuprecoveryv1.MaxSnapshotConfig)
	maxSnapshotConfigModel.MaxSnapshots = core.Float64Ptr(float64(72.5))

	maxSpaceConfigModel := new(backuprecoveryv1.MaxSpaceConfig)
	maxSpaceConfigModel.MaxSnapshotSpacePercentage = core.Float64Ptr(float64(72.5))

	storageArraySnapshotThrottlingPoliciesModel := new(backuprecoveryv1.StorageArraySnapshotThrottlingPolicies)
	storageArraySnapshotThrottlingPoliciesModel.ID = core.Int64Ptr(int64(26))
	storageArraySnapshotThrottlingPoliciesModel.IsMaxSnapshotsConfigEnabled = core.BoolPtr(true)
	storageArraySnapshotThrottlingPoliciesModel.IsMaxSpaceConfigEnabled = core.BoolPtr(true)
	storageArraySnapshotThrottlingPoliciesModel.MaxSnapshotConfig = maxSnapshotConfigModel
	storageArraySnapshotThrottlingPoliciesModel.MaxSpaceConfig = maxSpaceConfigModel

	storageArraySnapshotConfigModel := new(backuprecoveryv1.StorageArraySnapshotConfig)
	storageArraySnapshotConfigModel.IsMaxSnapshotsConfigEnabled = core.BoolPtr(true)
	storageArraySnapshotConfigModel.IsMaxSpaceConfigEnabled = core.BoolPtr(true)
	storageArraySnapshotConfigModel.StorageArraySnapshotMaxSpaceConfig = storageArraySnapshotMaxSpaceConfigModel
	storageArraySnapshotConfigModel.StorageArraySnapshotThrottlingPolicies = []backuprecoveryv1.StorageArraySnapshotThrottlingPolicies{*storageArraySnapshotThrottlingPoliciesModel}

	throttlingPolicyModel := new(backuprecoveryv1.ThrottlingPolicy)
	throttlingPolicyModel.EnforceMaxStreams = core.BoolPtr(true)
	throttlingPolicyModel.EnforceRegisteredSourceMaxBackups = core.BoolPtr(true)
	throttlingPolicyModel.IsEnabled = core.BoolPtr(true)
	throttlingPolicyModel.LatencyThresholds = latencyThresholdsModel
	throttlingPolicyModel.MaxConcurrentStreams = core.Float64Ptr(float64(72.5))
	throttlingPolicyModel.NasSourceParams = nasSourceParamsModel
	throttlingPolicyModel.RegisteredSourceMaxConcurrentBackups = core.Float64Ptr(float64(72.5))
	throttlingPolicyModel.StorageArraySnapshotConfig = storageArraySnapshotConfigModel

	throttlingPolicyOverridesModel := new(backuprecoveryv1.ThrottlingPolicyOverrides)
	throttlingPolicyOverridesModel.DatastoreID = core.Int64Ptr(int64(26))
	throttlingPolicyOverridesModel.DatastoreName = core.StringPtr("testString")
	throttlingPolicyOverridesModel.ThrottlingPolicy = throttlingPolicyModel

	registeredSourceVlanConfigModel := new(backuprecoveryv1.RegisteredSourceVlanConfig)
	registeredSourceVlanConfigModel.Vlan = core.Float64Ptr(float64(72.5))
	registeredSourceVlanConfigModel.DisableVlan = core.BoolPtr(true)
	registeredSourceVlanConfigModel.InterfaceName = core.StringPtr("testString")

	agentRegistrationInfoModel := new(backuprecoveryv1.AgentRegistrationInfo)
	agentRegistrationInfoModel.AccessInfo = agentAccessInfoModel
	agentRegistrationInfoModel.AllowedIpAddresses = []string{"testString"}
	agentRegistrationInfoModel.AuthenticationErrorMessage = core.StringPtr("testString")
	agentRegistrationInfoModel.AuthenticationStatus = core.StringPtr("kPending")
	agentRegistrationInfoModel.BlacklistedIpAddresses = []string{"testString"}
	agentRegistrationInfoModel.DeniedIpAddresses = []string{"testString"}
	agentRegistrationInfoModel.Environments = []string{"kPhysical"}
	agentRegistrationInfoModel.IsDbAuthenticated = core.BoolPtr(true)
	agentRegistrationInfoModel.IsStorageArraySnapshotEnabled = core.BoolPtr(true)
	agentRegistrationInfoModel.LinkVmsAcrossVcenter = core.BoolPtr(true)
	agentRegistrationInfoModel.MinimumFreeSpaceGB = core.Int64Ptr(int64(26))
	agentRegistrationInfoModel.MinimumFreeSpacePercent = core.Int64Ptr(int64(26))
	agentRegistrationInfoModel.Password = core.StringPtr("testString")
	agentRegistrationInfoModel.PhysicalParams = agentPhysicalParamsModel
	agentRegistrationInfoModel.ProgressMonitorPath = core.StringPtr("testString")
	agentRegistrationInfoModel.RefreshErrorMessage = core.StringPtr("testString")
	agentRegistrationInfoModel.RefreshTimeUsecs = core.Int64Ptr(int64(26))
	agentRegistrationInfoModel.RegisteredAppsInfo = []backuprecoveryv1.RegisteredAppInfo{*registeredAppInfoModel}
	agentRegistrationInfoModel.RegistrationTimeUsecs = core.Int64Ptr(int64(26))
	agentRegistrationInfoModel.Subnets = []backuprecoveryv1.Subnet{*subnetModel}
	agentRegistrationInfoModel.ThrottlingPolicy = throttlingPolicyModel
	agentRegistrationInfoModel.ThrottlingPolicyOverrides = []backuprecoveryv1.ThrottlingPolicyOverrides{*throttlingPolicyOverridesModel}
	agentRegistrationInfoModel.UseOAuthForExchangeOnline = core.BoolPtr(true)
	agentRegistrationInfoModel.UseVmBiosUUID = core.BoolPtr(true)
	agentRegistrationInfoModel.UserMessages = []string{"testString"}
	agentRegistrationInfoModel.Username = core.StringPtr("testString")
	agentRegistrationInfoModel.VlanParams = registeredSourceVlanConfigModel
	agentRegistrationInfoModel.WarningMessages = []string{"testString"}

	model := new(backuprecoveryv1.AgentInformation)
	model.CbmrVersion = core.StringPtr("testString")
	model.FileCbtInfo = cbtInfoModel
	model.HostType = core.StringPtr("kLinux")
	model.ID = core.Int64Ptr(int64(26))
	model.Name = core.StringPtr("testString")
	model.OracleMultiNodeChannelSupported = core.BoolPtr(true)
	model.RegistrationInfo = agentRegistrationInfoModel
	model.SourceSideDedupEnabled = core.BoolPtr(true)
	model.Status = core.StringPtr("kUnknown")
	model.StatusMessage = core.StringPtr("testString")
	model.Upgradability = core.StringPtr("kUpgradable")
	model.UpgradeStatus = core.StringPtr("kIdle")
	model.UpgradeStatusMessage = core.StringPtr("testString")
	model.Version = core.StringPtr("testString")
	model.VolCbtInfo = cbtInfoModel

	result, err := backuprecovery.DataSourceIbmBackupRecoveryRegistrationInfoAgentInformationToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryRegistrationInfoCbtInfoToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		cbtFileVersionModel := make(map[string]interface{})
		cbtFileVersionModel["build_ver"] = float64(72.5)
		cbtFileVersionModel["major_ver"] = float64(72.5)
		cbtFileVersionModel["minor_ver"] = float64(72.5)
		cbtFileVersionModel["revision_num"] = float64(72.5)

		cbtServiceStateModel := make(map[string]interface{})
		cbtServiceStateModel["name"] = "testString"
		cbtServiceStateModel["state"] = "testString"

		model := make(map[string]interface{})
		model["file_version"] = []map[string]interface{}{cbtFileVersionModel}
		model["is_installed"] = true
		model["reboot_status"] = "kRebooted"
		model["service_state"] = []map[string]interface{}{cbtServiceStateModel}

		assert.Equal(t, result, model)
	}

	cbtFileVersionModel := new(backuprecoveryv1.CbtFileVersion)
	cbtFileVersionModel.BuildVer = core.Float64Ptr(float64(72.5))
	cbtFileVersionModel.MajorVer = core.Float64Ptr(float64(72.5))
	cbtFileVersionModel.MinorVer = core.Float64Ptr(float64(72.5))
	cbtFileVersionModel.RevisionNum = core.Float64Ptr(float64(72.5))

	cbtServiceStateModel := new(backuprecoveryv1.CbtServiceState)
	cbtServiceStateModel.Name = core.StringPtr("testString")
	cbtServiceStateModel.State = core.StringPtr("testString")

	model := new(backuprecoveryv1.CbtInfo)
	model.FileVersion = cbtFileVersionModel
	model.IsInstalled = core.BoolPtr(true)
	model.RebootStatus = core.StringPtr("kRebooted")
	model.ServiceState = cbtServiceStateModel

	result, err := backuprecovery.DataSourceIbmBackupRecoveryRegistrationInfoCbtInfoToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryRegistrationInfoCbtFileVersionToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["build_ver"] = float64(72.5)
		model["major_ver"] = float64(72.5)
		model["minor_ver"] = float64(72.5)
		model["revision_num"] = float64(72.5)

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.CbtFileVersion)
	model.BuildVer = core.Float64Ptr(float64(72.5))
	model.MajorVer = core.Float64Ptr(float64(72.5))
	model.MinorVer = core.Float64Ptr(float64(72.5))
	model.RevisionNum = core.Float64Ptr(float64(72.5))

	result, err := backuprecovery.DataSourceIbmBackupRecoveryRegistrationInfoCbtFileVersionToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryRegistrationInfoCbtServiceStateToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["name"] = "testString"
		model["state"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.CbtServiceState)
	model.Name = core.StringPtr("testString")
	model.State = core.StringPtr("testString")

	result, err := backuprecovery.DataSourceIbmBackupRecoveryRegistrationInfoCbtServiceStateToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryRegistrationInfoAgentRegistrationInfoToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		agentAccessInfoModel := make(map[string]interface{})
		agentAccessInfoModel["connection_id"] = int(26)
		agentAccessInfoModel["connector_group_id"] = int(26)
		agentAccessInfoModel["endpoint"] = "testString"
		agentAccessInfoModel["environment"] = "kPhysical"
		agentAccessInfoModel["id"] = int(26)
		agentAccessInfoModel["version"] = int(26)

		timeModel := make(map[string]interface{})
		timeModel["hour"] = int(38)
		timeModel["minute"] = int(38)

		dayTimeParamsModel := make(map[string]interface{})
		dayTimeParamsModel["day"] = "kSunday"
		dayTimeParamsModel["time"] = []map[string]interface{}{timeModel}

		dayTimeWindowModel := make(map[string]interface{})
		dayTimeWindowModel["end_time"] = []map[string]interface{}{dayTimeParamsModel}
		dayTimeWindowModel["start_time"] = []map[string]interface{}{dayTimeParamsModel}

		throttlingWindowModel := make(map[string]interface{})
		throttlingWindowModel["day_time_window"] = []map[string]interface{}{dayTimeWindowModel}
		throttlingWindowModel["threshold"] = int(26)

		throttlingConfigurationParamsModel := make(map[string]interface{})
		throttlingConfigurationParamsModel["fixed_threshold"] = int(26)
		throttlingConfigurationParamsModel["pattern_type"] = "kNoThrottling"
		throttlingConfigurationParamsModel["throttling_windows"] = []map[string]interface{}{throttlingWindowModel}

		throttlingConfigModel := make(map[string]interface{})
		throttlingConfigModel["cpu_throttling_config"] = []map[string]interface{}{throttlingConfigurationParamsModel}
		throttlingConfigModel["network_throttling_config"] = []map[string]interface{}{throttlingConfigurationParamsModel}

		agentPhysicalParamsModel := make(map[string]interface{})
		agentPhysicalParamsModel["applications"] = []string{"kSQL"}
		agentPhysicalParamsModel["password"] = "testString"
		agentPhysicalParamsModel["throttling_config"] = []map[string]interface{}{throttlingConfigModel}
		agentPhysicalParamsModel["username"] = "testString"

		hostSettingsCheckResultModel := make(map[string]interface{})
		hostSettingsCheckResultModel["check_type"] = "kIsAgentPortAccessible"
		hostSettingsCheckResultModel["result_type"] = "kPass"
		hostSettingsCheckResultModel["user_message"] = "testString"

		registeredAppInfoModel := make(map[string]interface{})
		registeredAppInfoModel["authentication_error_message"] = "testString"
		registeredAppInfoModel["authentication_status"] = "kPending"
		registeredAppInfoModel["environment"] = "kPhysical"
		registeredAppInfoModel["host_settings_check_results"] = []map[string]interface{}{hostSettingsCheckResultModel}
		registeredAppInfoModel["refresh_error_message"] = "testString"

		subnetModel := make(map[string]interface{})
		subnetModel["component"] = "testString"
		subnetModel["description"] = "testString"
		subnetModel["id"] = float64(72.5)
		subnetModel["ip"] = "testString"
		subnetModel["netmask_bits"] = float64(72.5)
		subnetModel["netmask_ip4"] = "testString"
		subnetModel["nfs_access"] = "kDisabled"
		subnetModel["nfs_all_squash"] = true
		subnetModel["nfs_root_squash"] = true
		subnetModel["s3_access"] = "kDisabled"
		subnetModel["smb_access"] = "kDisabled"
		subnetModel["tenant_id"] = "testString"

		latencyThresholdsModel := make(map[string]interface{})
		latencyThresholdsModel["active_task_msecs"] = int(26)
		latencyThresholdsModel["new_task_msecs"] = int(26)

		nasSourceParamsModel := make(map[string]interface{})
		nasSourceParamsModel["max_parallel_metadata_fetch_full_percentage"] = float64(72.5)
		nasSourceParamsModel["max_parallel_metadata_fetch_incremental_percentage"] = float64(72.5)
		nasSourceParamsModel["max_parallel_read_write_full_percentage"] = float64(72.5)
		nasSourceParamsModel["max_parallel_read_write_incremental_percentage"] = float64(72.5)

		storageArraySnapshotMaxSpaceConfigModel := make(map[string]interface{})
		storageArraySnapshotMaxSpaceConfigModel["max_snapshot_space_percentage"] = float64(72.5)

		maxSnapshotConfigModel := make(map[string]interface{})
		maxSnapshotConfigModel["max_snapshots"] = float64(72.5)

		maxSpaceConfigModel := make(map[string]interface{})
		maxSpaceConfigModel["max_snapshot_space_percentage"] = float64(72.5)

		storageArraySnapshotThrottlingPoliciesModel := make(map[string]interface{})
		storageArraySnapshotThrottlingPoliciesModel["id"] = int(26)
		storageArraySnapshotThrottlingPoliciesModel["is_max_snapshots_config_enabled"] = true
		storageArraySnapshotThrottlingPoliciesModel["is_max_space_config_enabled"] = true
		storageArraySnapshotThrottlingPoliciesModel["max_snapshot_config"] = []map[string]interface{}{maxSnapshotConfigModel}
		storageArraySnapshotThrottlingPoliciesModel["max_space_config"] = []map[string]interface{}{maxSpaceConfigModel}

		storageArraySnapshotConfigModel := make(map[string]interface{})
		storageArraySnapshotConfigModel["is_max_snapshots_config_enabled"] = true
		storageArraySnapshotConfigModel["is_max_space_config_enabled"] = true
		storageArraySnapshotConfigModel["storage_array_snapshot_max_space_config"] = []map[string]interface{}{storageArraySnapshotMaxSpaceConfigModel}
		storageArraySnapshotConfigModel["storage_array_snapshot_throttling_policies"] = []map[string]interface{}{storageArraySnapshotThrottlingPoliciesModel}

		throttlingPolicyModel := make(map[string]interface{})
		throttlingPolicyModel["enforce_max_streams"] = true
		throttlingPolicyModel["enforce_registered_source_max_backups"] = true
		throttlingPolicyModel["is_enabled"] = true
		throttlingPolicyModel["latency_thresholds"] = []map[string]interface{}{latencyThresholdsModel}
		throttlingPolicyModel["max_concurrent_streams"] = float64(72.5)
		throttlingPolicyModel["nas_source_params"] = []map[string]interface{}{nasSourceParamsModel}
		throttlingPolicyModel["registered_source_max_concurrent_backups"] = float64(72.5)
		throttlingPolicyModel["storage_array_snapshot_config"] = []map[string]interface{}{storageArraySnapshotConfigModel}

		throttlingPolicyOverridesModel := make(map[string]interface{})
		throttlingPolicyOverridesModel["datastore_id"] = int(26)
		throttlingPolicyOverridesModel["datastore_name"] = "testString"
		throttlingPolicyOverridesModel["throttling_policy"] = []map[string]interface{}{throttlingPolicyModel}

		registeredSourceVlanConfigModel := make(map[string]interface{})
		registeredSourceVlanConfigModel["vlan"] = float64(72.5)
		registeredSourceVlanConfigModel["disable_vlan"] = true
		registeredSourceVlanConfigModel["interface_name"] = "testString"

		model := make(map[string]interface{})
		model["access_info"] = []map[string]interface{}{agentAccessInfoModel}
		model["allowed_ip_addresses"] = []string{"testString"}
		model["authentication_error_message"] = "testString"
		model["authentication_status"] = "kPending"
		model["blacklisted_ip_addresses"] = []string{"testString"}
		model["denied_ip_addresses"] = []string{"testString"}
		model["environments"] = []string{"kPhysical"}
		model["is_db_authenticated"] = true
		model["is_storage_array_snapshot_enabled"] = true
		model["link_vms_across_vcenter"] = true
		model["minimum_free_space_gb"] = int(26)
		model["minimum_free_space_percent"] = int(26)
		model["password"] = "testString"
		model["physical_params"] = []map[string]interface{}{agentPhysicalParamsModel}
		model["progress_monitor_path"] = "testString"
		model["refresh_error_message"] = "testString"
		model["refresh_time_usecs"] = int(26)
		model["registered_apps_info"] = []map[string]interface{}{registeredAppInfoModel}
		model["registration_time_usecs"] = int(26)
		model["subnets"] = []map[string]interface{}{subnetModel}
		model["throttling_policy"] = []map[string]interface{}{throttlingPolicyModel}
		model["throttling_policy_overrides"] = []map[string]interface{}{throttlingPolicyOverridesModel}
		model["use_o_auth_for_exchange_online"] = true
		model["use_vm_bios_uuid"] = true
		model["user_messages"] = []string{"testString"}
		model["username"] = "testString"
		model["vlan_params"] = []map[string]interface{}{registeredSourceVlanConfigModel}
		model["warning_messages"] = []string{"testString"}

		assert.Equal(t, result, model)
	}

	agentAccessInfoModel := new(backuprecoveryv1.AgentAccessInfo)
	agentAccessInfoModel.ConnectionID = core.Int64Ptr(int64(26))
	agentAccessInfoModel.ConnectorGroupID = core.Int64Ptr(int64(26))
	agentAccessInfoModel.Endpoint = core.StringPtr("testString")
	agentAccessInfoModel.Environment = core.StringPtr("kPhysical")
	agentAccessInfoModel.ID = core.Int64Ptr(int64(26))
	agentAccessInfoModel.Version = core.Int64Ptr(int64(26))

	timeModel := new(backuprecoveryv1.Time)
	timeModel.Hour = core.Int64Ptr(int64(38))
	timeModel.Minute = core.Int64Ptr(int64(38))

	dayTimeParamsModel := new(backuprecoveryv1.DayTimeParams)
	dayTimeParamsModel.Day = core.StringPtr("kSunday")
	dayTimeParamsModel.Time = timeModel

	dayTimeWindowModel := new(backuprecoveryv1.DayTimeWindow)
	dayTimeWindowModel.EndTime = dayTimeParamsModel
	dayTimeWindowModel.StartTime = dayTimeParamsModel

	throttlingWindowModel := new(backuprecoveryv1.ThrottlingWindow)
	throttlingWindowModel.DayTimeWindow = dayTimeWindowModel
	throttlingWindowModel.Threshold = core.Int64Ptr(int64(26))

	throttlingConfigurationParamsModel := new(backuprecoveryv1.ThrottlingConfigurationParams)
	throttlingConfigurationParamsModel.FixedThreshold = core.Int64Ptr(int64(26))
	throttlingConfigurationParamsModel.PatternType = core.StringPtr("kNoThrottling")
	throttlingConfigurationParamsModel.ThrottlingWindows = []backuprecoveryv1.ThrottlingWindow{*throttlingWindowModel}

	throttlingConfigModel := new(backuprecoveryv1.ThrottlingConfig)
	throttlingConfigModel.CpuThrottlingConfig = throttlingConfigurationParamsModel
	throttlingConfigModel.NetworkThrottlingConfig = throttlingConfigurationParamsModel

	agentPhysicalParamsModel := new(backuprecoveryv1.AgentPhysicalParams)
	agentPhysicalParamsModel.Applications = []string{"kSQL"}
	agentPhysicalParamsModel.Password = core.StringPtr("testString")
	agentPhysicalParamsModel.ThrottlingConfig = throttlingConfigModel
	agentPhysicalParamsModel.Username = core.StringPtr("testString")

	hostSettingsCheckResultModel := new(backuprecoveryv1.HostSettingsCheckResult)
	hostSettingsCheckResultModel.CheckType = core.StringPtr("kIsAgentPortAccessible")
	hostSettingsCheckResultModel.ResultType = core.StringPtr("kPass")
	hostSettingsCheckResultModel.UserMessage = core.StringPtr("testString")

	registeredAppInfoModel := new(backuprecoveryv1.RegisteredAppInfo)
	registeredAppInfoModel.AuthenticationErrorMessage = core.StringPtr("testString")
	registeredAppInfoModel.AuthenticationStatus = core.StringPtr("kPending")
	registeredAppInfoModel.Environment = core.StringPtr("kPhysical")
	registeredAppInfoModel.HostSettingsCheckResults = []backuprecoveryv1.HostSettingsCheckResult{*hostSettingsCheckResultModel}
	registeredAppInfoModel.RefreshErrorMessage = core.StringPtr("testString")

	subnetModel := new(backuprecoveryv1.Subnet)
	subnetModel.Component = core.StringPtr("testString")
	subnetModel.Description = core.StringPtr("testString")
	subnetModel.ID = core.Float64Ptr(float64(72.5))
	subnetModel.Ip = core.StringPtr("testString")
	subnetModel.NetmaskBits = core.Float64Ptr(float64(72.5))
	subnetModel.NetmaskIp4 = core.StringPtr("testString")
	subnetModel.NfsAccess = core.StringPtr("kDisabled")
	subnetModel.NfsAllSquash = core.BoolPtr(true)
	subnetModel.NfsRootSquash = core.BoolPtr(true)
	subnetModel.S3Access = core.StringPtr("kDisabled")
	subnetModel.SmbAccess = core.StringPtr("kDisabled")
	subnetModel.TenantID = core.StringPtr("testString")

	latencyThresholdsModel := new(backuprecoveryv1.LatencyThresholds)
	latencyThresholdsModel.ActiveTaskMsecs = core.Int64Ptr(int64(26))
	latencyThresholdsModel.NewTaskMsecs = core.Int64Ptr(int64(26))

	nasSourceParamsModel := new(backuprecoveryv1.NasSourceParams)
	nasSourceParamsModel.MaxParallelMetadataFetchFullPercentage = core.Float64Ptr(float64(72.5))
	nasSourceParamsModel.MaxParallelMetadataFetchIncrementalPercentage = core.Float64Ptr(float64(72.5))
	nasSourceParamsModel.MaxParallelReadWriteFullPercentage = core.Float64Ptr(float64(72.5))
	nasSourceParamsModel.MaxParallelReadWriteIncrementalPercentage = core.Float64Ptr(float64(72.5))

	storageArraySnapshotMaxSpaceConfigModel := new(backuprecoveryv1.StorageArraySnapshotMaxSpaceConfig)
	storageArraySnapshotMaxSpaceConfigModel.MaxSnapshotSpacePercentage = core.Float64Ptr(float64(72.5))

	maxSnapshotConfigModel := new(backuprecoveryv1.MaxSnapshotConfig)
	maxSnapshotConfigModel.MaxSnapshots = core.Float64Ptr(float64(72.5))

	maxSpaceConfigModel := new(backuprecoveryv1.MaxSpaceConfig)
	maxSpaceConfigModel.MaxSnapshotSpacePercentage = core.Float64Ptr(float64(72.5))

	storageArraySnapshotThrottlingPoliciesModel := new(backuprecoveryv1.StorageArraySnapshotThrottlingPolicies)
	storageArraySnapshotThrottlingPoliciesModel.ID = core.Int64Ptr(int64(26))
	storageArraySnapshotThrottlingPoliciesModel.IsMaxSnapshotsConfigEnabled = core.BoolPtr(true)
	storageArraySnapshotThrottlingPoliciesModel.IsMaxSpaceConfigEnabled = core.BoolPtr(true)
	storageArraySnapshotThrottlingPoliciesModel.MaxSnapshotConfig = maxSnapshotConfigModel
	storageArraySnapshotThrottlingPoliciesModel.MaxSpaceConfig = maxSpaceConfigModel

	storageArraySnapshotConfigModel := new(backuprecoveryv1.StorageArraySnapshotConfig)
	storageArraySnapshotConfigModel.IsMaxSnapshotsConfigEnabled = core.BoolPtr(true)
	storageArraySnapshotConfigModel.IsMaxSpaceConfigEnabled = core.BoolPtr(true)
	storageArraySnapshotConfigModel.StorageArraySnapshotMaxSpaceConfig = storageArraySnapshotMaxSpaceConfigModel
	storageArraySnapshotConfigModel.StorageArraySnapshotThrottlingPolicies = []backuprecoveryv1.StorageArraySnapshotThrottlingPolicies{*storageArraySnapshotThrottlingPoliciesModel}

	throttlingPolicyModel := new(backuprecoveryv1.ThrottlingPolicy)
	throttlingPolicyModel.EnforceMaxStreams = core.BoolPtr(true)
	throttlingPolicyModel.EnforceRegisteredSourceMaxBackups = core.BoolPtr(true)
	throttlingPolicyModel.IsEnabled = core.BoolPtr(true)
	throttlingPolicyModel.LatencyThresholds = latencyThresholdsModel
	throttlingPolicyModel.MaxConcurrentStreams = core.Float64Ptr(float64(72.5))
	throttlingPolicyModel.NasSourceParams = nasSourceParamsModel
	throttlingPolicyModel.RegisteredSourceMaxConcurrentBackups = core.Float64Ptr(float64(72.5))
	throttlingPolicyModel.StorageArraySnapshotConfig = storageArraySnapshotConfigModel

	throttlingPolicyOverridesModel := new(backuprecoveryv1.ThrottlingPolicyOverrides)
	throttlingPolicyOverridesModel.DatastoreID = core.Int64Ptr(int64(26))
	throttlingPolicyOverridesModel.DatastoreName = core.StringPtr("testString")
	throttlingPolicyOverridesModel.ThrottlingPolicy = throttlingPolicyModel

	registeredSourceVlanConfigModel := new(backuprecoveryv1.RegisteredSourceVlanConfig)
	registeredSourceVlanConfigModel.Vlan = core.Float64Ptr(float64(72.5))
	registeredSourceVlanConfigModel.DisableVlan = core.BoolPtr(true)
	registeredSourceVlanConfigModel.InterfaceName = core.StringPtr("testString")

	model := new(backuprecoveryv1.AgentRegistrationInfo)
	model.AccessInfo = agentAccessInfoModel
	model.AllowedIpAddresses = []string{"testString"}
	model.AuthenticationErrorMessage = core.StringPtr("testString")
	model.AuthenticationStatus = core.StringPtr("kPending")
	model.BlacklistedIpAddresses = []string{"testString"}
	model.DeniedIpAddresses = []string{"testString"}
	model.Environments = []string{"kPhysical"}
	model.IsDbAuthenticated = core.BoolPtr(true)
	model.IsStorageArraySnapshotEnabled = core.BoolPtr(true)
	model.LinkVmsAcrossVcenter = core.BoolPtr(true)
	model.MinimumFreeSpaceGB = core.Int64Ptr(int64(26))
	model.MinimumFreeSpacePercent = core.Int64Ptr(int64(26))
	model.Password = core.StringPtr("testString")
	model.PhysicalParams = agentPhysicalParamsModel
	model.ProgressMonitorPath = core.StringPtr("testString")
	model.RefreshErrorMessage = core.StringPtr("testString")
	model.RefreshTimeUsecs = core.Int64Ptr(int64(26))
	model.RegisteredAppsInfo = []backuprecoveryv1.RegisteredAppInfo{*registeredAppInfoModel}
	model.RegistrationTimeUsecs = core.Int64Ptr(int64(26))
	model.Subnets = []backuprecoveryv1.Subnet{*subnetModel}
	model.ThrottlingPolicy = throttlingPolicyModel
	model.ThrottlingPolicyOverrides = []backuprecoveryv1.ThrottlingPolicyOverrides{*throttlingPolicyOverridesModel}
	model.UseOAuthForExchangeOnline = core.BoolPtr(true)
	model.UseVmBiosUUID = core.BoolPtr(true)
	model.UserMessages = []string{"testString"}
	model.Username = core.StringPtr("testString")
	model.VlanParams = registeredSourceVlanConfigModel
	model.WarningMessages = []string{"testString"}

	result, err := backuprecovery.DataSourceIbmBackupRecoveryRegistrationInfoAgentRegistrationInfoToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryRegistrationInfoAgentAccessInfoToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["connection_id"] = int(26)
		model["connector_group_id"] = int(26)
		model["endpoint"] = "testString"
		model["environment"] = "kPhysical"
		model["id"] = int(26)
		model["version"] = int(26)

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.AgentAccessInfo)
	model.ConnectionID = core.Int64Ptr(int64(26))
	model.ConnectorGroupID = core.Int64Ptr(int64(26))
	model.Endpoint = core.StringPtr("testString")
	model.Environment = core.StringPtr("kPhysical")
	model.ID = core.Int64Ptr(int64(26))
	model.Version = core.Int64Ptr(int64(26))

	result, err := backuprecovery.DataSourceIbmBackupRecoveryRegistrationInfoAgentAccessInfoToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryRegistrationInfoAgentPhysicalParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		timeModel := make(map[string]interface{})
		timeModel["hour"] = int(38)
		timeModel["minute"] = int(38)

		dayTimeParamsModel := make(map[string]interface{})
		dayTimeParamsModel["day"] = "kSunday"
		dayTimeParamsModel["time"] = []map[string]interface{}{timeModel}

		dayTimeWindowModel := make(map[string]interface{})
		dayTimeWindowModel["end_time"] = []map[string]interface{}{dayTimeParamsModel}
		dayTimeWindowModel["start_time"] = []map[string]interface{}{dayTimeParamsModel}

		throttlingWindowModel := make(map[string]interface{})
		throttlingWindowModel["day_time_window"] = []map[string]interface{}{dayTimeWindowModel}
		throttlingWindowModel["threshold"] = int(26)

		throttlingConfigurationParamsModel := make(map[string]interface{})
		throttlingConfigurationParamsModel["fixed_threshold"] = int(26)
		throttlingConfigurationParamsModel["pattern_type"] = "kNoThrottling"
		throttlingConfigurationParamsModel["throttling_windows"] = []map[string]interface{}{throttlingWindowModel}

		throttlingConfigModel := make(map[string]interface{})
		throttlingConfigModel["cpu_throttling_config"] = []map[string]interface{}{throttlingConfigurationParamsModel}
		throttlingConfigModel["network_throttling_config"] = []map[string]interface{}{throttlingConfigurationParamsModel}

		model := make(map[string]interface{})
		model["applications"] = []string{"kSQL"}
		model["password"] = "testString"
		model["throttling_config"] = []map[string]interface{}{throttlingConfigModel}
		model["username"] = "testString"

		assert.Equal(t, result, model)
	}

	timeModel := new(backuprecoveryv1.Time)
	timeModel.Hour = core.Int64Ptr(int64(38))
	timeModel.Minute = core.Int64Ptr(int64(38))

	dayTimeParamsModel := new(backuprecoveryv1.DayTimeParams)
	dayTimeParamsModel.Day = core.StringPtr("kSunday")
	dayTimeParamsModel.Time = timeModel

	dayTimeWindowModel := new(backuprecoveryv1.DayTimeWindow)
	dayTimeWindowModel.EndTime = dayTimeParamsModel
	dayTimeWindowModel.StartTime = dayTimeParamsModel

	throttlingWindowModel := new(backuprecoveryv1.ThrottlingWindow)
	throttlingWindowModel.DayTimeWindow = dayTimeWindowModel
	throttlingWindowModel.Threshold = core.Int64Ptr(int64(26))

	throttlingConfigurationParamsModel := new(backuprecoveryv1.ThrottlingConfigurationParams)
	throttlingConfigurationParamsModel.FixedThreshold = core.Int64Ptr(int64(26))
	throttlingConfigurationParamsModel.PatternType = core.StringPtr("kNoThrottling")
	throttlingConfigurationParamsModel.ThrottlingWindows = []backuprecoveryv1.ThrottlingWindow{*throttlingWindowModel}

	throttlingConfigModel := new(backuprecoveryv1.ThrottlingConfig)
	throttlingConfigModel.CpuThrottlingConfig = throttlingConfigurationParamsModel
	throttlingConfigModel.NetworkThrottlingConfig = throttlingConfigurationParamsModel

	model := new(backuprecoveryv1.AgentPhysicalParams)
	model.Applications = []string{"kSQL"}
	model.Password = core.StringPtr("testString")
	model.ThrottlingConfig = throttlingConfigModel
	model.Username = core.StringPtr("testString")

	result, err := backuprecovery.DataSourceIbmBackupRecoveryRegistrationInfoAgentPhysicalParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryRegistrationInfoThrottlingConfigToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		timeModel := make(map[string]interface{})
		timeModel["hour"] = int(38)
		timeModel["minute"] = int(38)

		dayTimeParamsModel := make(map[string]interface{})
		dayTimeParamsModel["day"] = "kSunday"
		dayTimeParamsModel["time"] = []map[string]interface{}{timeModel}

		dayTimeWindowModel := make(map[string]interface{})
		dayTimeWindowModel["end_time"] = []map[string]interface{}{dayTimeParamsModel}
		dayTimeWindowModel["start_time"] = []map[string]interface{}{dayTimeParamsModel}

		throttlingWindowModel := make(map[string]interface{})
		throttlingWindowModel["day_time_window"] = []map[string]interface{}{dayTimeWindowModel}
		throttlingWindowModel["threshold"] = int(26)

		throttlingConfigurationParamsModel := make(map[string]interface{})
		throttlingConfigurationParamsModel["fixed_threshold"] = int(26)
		throttlingConfigurationParamsModel["pattern_type"] = "kNoThrottling"
		throttlingConfigurationParamsModel["throttling_windows"] = []map[string]interface{}{throttlingWindowModel}

		model := make(map[string]interface{})
		model["cpu_throttling_config"] = []map[string]interface{}{throttlingConfigurationParamsModel}
		model["network_throttling_config"] = []map[string]interface{}{throttlingConfigurationParamsModel}

		assert.Equal(t, result, model)
	}

	timeModel := new(backuprecoveryv1.Time)
	timeModel.Hour = core.Int64Ptr(int64(38))
	timeModel.Minute = core.Int64Ptr(int64(38))

	dayTimeParamsModel := new(backuprecoveryv1.DayTimeParams)
	dayTimeParamsModel.Day = core.StringPtr("kSunday")
	dayTimeParamsModel.Time = timeModel

	dayTimeWindowModel := new(backuprecoveryv1.DayTimeWindow)
	dayTimeWindowModel.EndTime = dayTimeParamsModel
	dayTimeWindowModel.StartTime = dayTimeParamsModel

	throttlingWindowModel := new(backuprecoveryv1.ThrottlingWindow)
	throttlingWindowModel.DayTimeWindow = dayTimeWindowModel
	throttlingWindowModel.Threshold = core.Int64Ptr(int64(26))

	throttlingConfigurationParamsModel := new(backuprecoveryv1.ThrottlingConfigurationParams)
	throttlingConfigurationParamsModel.FixedThreshold = core.Int64Ptr(int64(26))
	throttlingConfigurationParamsModel.PatternType = core.StringPtr("kNoThrottling")
	throttlingConfigurationParamsModel.ThrottlingWindows = []backuprecoveryv1.ThrottlingWindow{*throttlingWindowModel}

	model := new(backuprecoveryv1.ThrottlingConfig)
	model.CpuThrottlingConfig = throttlingConfigurationParamsModel
	model.NetworkThrottlingConfig = throttlingConfigurationParamsModel

	result, err := backuprecovery.DataSourceIbmBackupRecoveryRegistrationInfoThrottlingConfigToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryRegistrationInfoThrottlingConfigurationParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		timeModel := make(map[string]interface{})
		timeModel["hour"] = int(38)
		timeModel["minute"] = int(38)

		dayTimeParamsModel := make(map[string]interface{})
		dayTimeParamsModel["day"] = "kSunday"
		dayTimeParamsModel["time"] = []map[string]interface{}{timeModel}

		dayTimeWindowModel := make(map[string]interface{})
		dayTimeWindowModel["end_time"] = []map[string]interface{}{dayTimeParamsModel}
		dayTimeWindowModel["start_time"] = []map[string]interface{}{dayTimeParamsModel}

		throttlingWindowModel := make(map[string]interface{})
		throttlingWindowModel["day_time_window"] = []map[string]interface{}{dayTimeWindowModel}
		throttlingWindowModel["threshold"] = int(26)

		model := make(map[string]interface{})
		model["fixed_threshold"] = int(26)
		model["pattern_type"] = "kNoThrottling"
		model["throttling_windows"] = []map[string]interface{}{throttlingWindowModel}

		assert.Equal(t, result, model)
	}

	timeModel := new(backuprecoveryv1.Time)
	timeModel.Hour = core.Int64Ptr(int64(38))
	timeModel.Minute = core.Int64Ptr(int64(38))

	dayTimeParamsModel := new(backuprecoveryv1.DayTimeParams)
	dayTimeParamsModel.Day = core.StringPtr("kSunday")
	dayTimeParamsModel.Time = timeModel

	dayTimeWindowModel := new(backuprecoveryv1.DayTimeWindow)
	dayTimeWindowModel.EndTime = dayTimeParamsModel
	dayTimeWindowModel.StartTime = dayTimeParamsModel

	throttlingWindowModel := new(backuprecoveryv1.ThrottlingWindow)
	throttlingWindowModel.DayTimeWindow = dayTimeWindowModel
	throttlingWindowModel.Threshold = core.Int64Ptr(int64(26))

	model := new(backuprecoveryv1.ThrottlingConfigurationParams)
	model.FixedThreshold = core.Int64Ptr(int64(26))
	model.PatternType = core.StringPtr("kNoThrottling")
	model.ThrottlingWindows = []backuprecoveryv1.ThrottlingWindow{*throttlingWindowModel}

	result, err := backuprecovery.DataSourceIbmBackupRecoveryRegistrationInfoThrottlingConfigurationParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryRegistrationInfoThrottlingWindowToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		timeModel := make(map[string]interface{})
		timeModel["hour"] = int(38)
		timeModel["minute"] = int(38)

		dayTimeParamsModel := make(map[string]interface{})
		dayTimeParamsModel["day"] = "kSunday"
		dayTimeParamsModel["time"] = []map[string]interface{}{timeModel}

		dayTimeWindowModel := make(map[string]interface{})
		dayTimeWindowModel["end_time"] = []map[string]interface{}{dayTimeParamsModel}
		dayTimeWindowModel["start_time"] = []map[string]interface{}{dayTimeParamsModel}

		model := make(map[string]interface{})
		model["day_time_window"] = []map[string]interface{}{dayTimeWindowModel}
		model["threshold"] = int(26)

		assert.Equal(t, result, model)
	}

	timeModel := new(backuprecoveryv1.Time)
	timeModel.Hour = core.Int64Ptr(int64(38))
	timeModel.Minute = core.Int64Ptr(int64(38))

	dayTimeParamsModel := new(backuprecoveryv1.DayTimeParams)
	dayTimeParamsModel.Day = core.StringPtr("kSunday")
	dayTimeParamsModel.Time = timeModel

	dayTimeWindowModel := new(backuprecoveryv1.DayTimeWindow)
	dayTimeWindowModel.EndTime = dayTimeParamsModel
	dayTimeWindowModel.StartTime = dayTimeParamsModel

	model := new(backuprecoveryv1.ThrottlingWindow)
	model.DayTimeWindow = dayTimeWindowModel
	model.Threshold = core.Int64Ptr(int64(26))

	result, err := backuprecovery.DataSourceIbmBackupRecoveryRegistrationInfoThrottlingWindowToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryRegistrationInfoDayTimeWindowToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		timeModel := make(map[string]interface{})
		timeModel["hour"] = int(38)
		timeModel["minute"] = int(38)

		dayTimeParamsModel := make(map[string]interface{})
		dayTimeParamsModel["day"] = "kSunday"
		dayTimeParamsModel["time"] = []map[string]interface{}{timeModel}

		model := make(map[string]interface{})
		model["end_time"] = []map[string]interface{}{dayTimeParamsModel}
		model["start_time"] = []map[string]interface{}{dayTimeParamsModel}

		assert.Equal(t, result, model)
	}

	timeModel := new(backuprecoveryv1.Time)
	timeModel.Hour = core.Int64Ptr(int64(38))
	timeModel.Minute = core.Int64Ptr(int64(38))

	dayTimeParamsModel := new(backuprecoveryv1.DayTimeParams)
	dayTimeParamsModel.Day = core.StringPtr("kSunday")
	dayTimeParamsModel.Time = timeModel

	model := new(backuprecoveryv1.DayTimeWindow)
	model.EndTime = dayTimeParamsModel
	model.StartTime = dayTimeParamsModel

	result, err := backuprecovery.DataSourceIbmBackupRecoveryRegistrationInfoDayTimeWindowToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryRegistrationInfoDayTimeParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		timeModel := make(map[string]interface{})
		timeModel["hour"] = int(38)
		timeModel["minute"] = int(38)

		model := make(map[string]interface{})
		model["day"] = "kSunday"
		model["time"] = []map[string]interface{}{timeModel}

		assert.Equal(t, result, model)
	}

	timeModel := new(backuprecoveryv1.Time)
	timeModel.Hour = core.Int64Ptr(int64(38))
	timeModel.Minute = core.Int64Ptr(int64(38))

	model := new(backuprecoveryv1.DayTimeParams)
	model.Day = core.StringPtr("kSunday")
	model.Time = timeModel

	result, err := backuprecovery.DataSourceIbmBackupRecoveryRegistrationInfoDayTimeParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryRegistrationInfoTimeToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["hour"] = int(38)
		model["minute"] = int(38)

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.Time)
	model.Hour = core.Int64Ptr(int64(38))
	model.Minute = core.Int64Ptr(int64(38))

	result, err := backuprecovery.DataSourceIbmBackupRecoveryRegistrationInfoTimeToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryRegistrationInfoRegisteredAppInfoToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		hostSettingsCheckResultModel := make(map[string]interface{})
		hostSettingsCheckResultModel["check_type"] = "kIsAgentPortAccessible"
		hostSettingsCheckResultModel["result_type"] = "kPass"
		hostSettingsCheckResultModel["user_message"] = "testString"

		model := make(map[string]interface{})
		model["authentication_error_message"] = "testString"
		model["authentication_status"] = "kPending"
		model["environment"] = "kPhysical"
		model["host_settings_check_results"] = []map[string]interface{}{hostSettingsCheckResultModel}
		model["refresh_error_message"] = "testString"

		assert.Equal(t, result, model)
	}

	hostSettingsCheckResultModel := new(backuprecoveryv1.HostSettingsCheckResult)
	hostSettingsCheckResultModel.CheckType = core.StringPtr("kIsAgentPortAccessible")
	hostSettingsCheckResultModel.ResultType = core.StringPtr("kPass")
	hostSettingsCheckResultModel.UserMessage = core.StringPtr("testString")

	model := new(backuprecoveryv1.RegisteredAppInfo)
	model.AuthenticationErrorMessage = core.StringPtr("testString")
	model.AuthenticationStatus = core.StringPtr("kPending")
	model.Environment = core.StringPtr("kPhysical")
	model.HostSettingsCheckResults = []backuprecoveryv1.HostSettingsCheckResult{*hostSettingsCheckResultModel}
	model.RefreshErrorMessage = core.StringPtr("testString")

	result, err := backuprecovery.DataSourceIbmBackupRecoveryRegistrationInfoRegisteredAppInfoToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryRegistrationInfoHostSettingsCheckResultToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["check_type"] = "kIsAgentPortAccessible"
		model["result_type"] = "kPass"
		model["user_message"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.HostSettingsCheckResult)
	model.CheckType = core.StringPtr("kIsAgentPortAccessible")
	model.ResultType = core.StringPtr("kPass")
	model.UserMessage = core.StringPtr("testString")

	result, err := backuprecovery.DataSourceIbmBackupRecoveryRegistrationInfoHostSettingsCheckResultToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryRegistrationInfoSubnetToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["component"] = "testString"
		model["description"] = "testString"
		model["id"] = float64(72.5)
		model["ip"] = "testString"
		model["netmask_bits"] = float64(72.5)
		model["netmask_ip4"] = "testString"
		model["nfs_access"] = "kDisabled"
		model["nfs_all_squash"] = true
		model["nfs_root_squash"] = true
		model["s3_access"] = "kDisabled"
		model["smb_access"] = "kDisabled"
		model["tenant_id"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.Subnet)
	model.Component = core.StringPtr("testString")
	model.Description = core.StringPtr("testString")
	model.ID = core.Float64Ptr(float64(72.5))
	model.Ip = core.StringPtr("testString")
	model.NetmaskBits = core.Float64Ptr(float64(72.5))
	model.NetmaskIp4 = core.StringPtr("testString")
	model.NfsAccess = core.StringPtr("kDisabled")
	model.NfsAllSquash = core.BoolPtr(true)
	model.NfsRootSquash = core.BoolPtr(true)
	model.S3Access = core.StringPtr("kDisabled")
	model.SmbAccess = core.StringPtr("kDisabled")
	model.TenantID = core.StringPtr("testString")

	result, err := backuprecovery.DataSourceIbmBackupRecoveryRegistrationInfoSubnetToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryRegistrationInfoThrottlingPolicyToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		latencyThresholdsModel := make(map[string]interface{})
		latencyThresholdsModel["active_task_msecs"] = int(26)
		latencyThresholdsModel["new_task_msecs"] = int(26)

		nasSourceParamsModel := make(map[string]interface{})
		nasSourceParamsModel["max_parallel_metadata_fetch_full_percentage"] = float64(72.5)
		nasSourceParamsModel["max_parallel_metadata_fetch_incremental_percentage"] = float64(72.5)
		nasSourceParamsModel["max_parallel_read_write_full_percentage"] = float64(72.5)
		nasSourceParamsModel["max_parallel_read_write_incremental_percentage"] = float64(72.5)

		storageArraySnapshotMaxSpaceConfigModel := make(map[string]interface{})
		storageArraySnapshotMaxSpaceConfigModel["max_snapshot_space_percentage"] = float64(72.5)

		maxSnapshotConfigModel := make(map[string]interface{})
		maxSnapshotConfigModel["max_snapshots"] = float64(72.5)

		maxSpaceConfigModel := make(map[string]interface{})
		maxSpaceConfigModel["max_snapshot_space_percentage"] = float64(72.5)

		storageArraySnapshotThrottlingPoliciesModel := make(map[string]interface{})
		storageArraySnapshotThrottlingPoliciesModel["id"] = int(26)
		storageArraySnapshotThrottlingPoliciesModel["is_max_snapshots_config_enabled"] = true
		storageArraySnapshotThrottlingPoliciesModel["is_max_space_config_enabled"] = true
		storageArraySnapshotThrottlingPoliciesModel["max_snapshot_config"] = []map[string]interface{}{maxSnapshotConfigModel}
		storageArraySnapshotThrottlingPoliciesModel["max_space_config"] = []map[string]interface{}{maxSpaceConfigModel}

		storageArraySnapshotConfigModel := make(map[string]interface{})
		storageArraySnapshotConfigModel["is_max_snapshots_config_enabled"] = true
		storageArraySnapshotConfigModel["is_max_space_config_enabled"] = true
		storageArraySnapshotConfigModel["storage_array_snapshot_max_space_config"] = []map[string]interface{}{storageArraySnapshotMaxSpaceConfigModel}
		storageArraySnapshotConfigModel["storage_array_snapshot_throttling_policies"] = []map[string]interface{}{storageArraySnapshotThrottlingPoliciesModel}

		model := make(map[string]interface{})
		model["enforce_max_streams"] = true
		model["enforce_registered_source_max_backups"] = true
		model["is_enabled"] = true
		model["latency_thresholds"] = []map[string]interface{}{latencyThresholdsModel}
		model["max_concurrent_streams"] = float64(72.5)
		model["nas_source_params"] = []map[string]interface{}{nasSourceParamsModel}
		model["registered_source_max_concurrent_backups"] = float64(72.5)
		model["storage_array_snapshot_config"] = []map[string]interface{}{storageArraySnapshotConfigModel}

		assert.Equal(t, result, model)
	}

	latencyThresholdsModel := new(backuprecoveryv1.LatencyThresholds)
	latencyThresholdsModel.ActiveTaskMsecs = core.Int64Ptr(int64(26))
	latencyThresholdsModel.NewTaskMsecs = core.Int64Ptr(int64(26))

	nasSourceParamsModel := new(backuprecoveryv1.NasSourceParams)
	nasSourceParamsModel.MaxParallelMetadataFetchFullPercentage = core.Float64Ptr(float64(72.5))
	nasSourceParamsModel.MaxParallelMetadataFetchIncrementalPercentage = core.Float64Ptr(float64(72.5))
	nasSourceParamsModel.MaxParallelReadWriteFullPercentage = core.Float64Ptr(float64(72.5))
	nasSourceParamsModel.MaxParallelReadWriteIncrementalPercentage = core.Float64Ptr(float64(72.5))

	storageArraySnapshotMaxSpaceConfigModel := new(backuprecoveryv1.StorageArraySnapshotMaxSpaceConfig)
	storageArraySnapshotMaxSpaceConfigModel.MaxSnapshotSpacePercentage = core.Float64Ptr(float64(72.5))

	maxSnapshotConfigModel := new(backuprecoveryv1.MaxSnapshotConfig)
	maxSnapshotConfigModel.MaxSnapshots = core.Float64Ptr(float64(72.5))

	maxSpaceConfigModel := new(backuprecoveryv1.MaxSpaceConfig)
	maxSpaceConfigModel.MaxSnapshotSpacePercentage = core.Float64Ptr(float64(72.5))

	storageArraySnapshotThrottlingPoliciesModel := new(backuprecoveryv1.StorageArraySnapshotThrottlingPolicies)
	storageArraySnapshotThrottlingPoliciesModel.ID = core.Int64Ptr(int64(26))
	storageArraySnapshotThrottlingPoliciesModel.IsMaxSnapshotsConfigEnabled = core.BoolPtr(true)
	storageArraySnapshotThrottlingPoliciesModel.IsMaxSpaceConfigEnabled = core.BoolPtr(true)
	storageArraySnapshotThrottlingPoliciesModel.MaxSnapshotConfig = maxSnapshotConfigModel
	storageArraySnapshotThrottlingPoliciesModel.MaxSpaceConfig = maxSpaceConfigModel

	storageArraySnapshotConfigModel := new(backuprecoveryv1.StorageArraySnapshotConfig)
	storageArraySnapshotConfigModel.IsMaxSnapshotsConfigEnabled = core.BoolPtr(true)
	storageArraySnapshotConfigModel.IsMaxSpaceConfigEnabled = core.BoolPtr(true)
	storageArraySnapshotConfigModel.StorageArraySnapshotMaxSpaceConfig = storageArraySnapshotMaxSpaceConfigModel
	storageArraySnapshotConfigModel.StorageArraySnapshotThrottlingPolicies = []backuprecoveryv1.StorageArraySnapshotThrottlingPolicies{*storageArraySnapshotThrottlingPoliciesModel}

	model := new(backuprecoveryv1.ThrottlingPolicy)
	model.EnforceMaxStreams = core.BoolPtr(true)
	model.EnforceRegisteredSourceMaxBackups = core.BoolPtr(true)
	model.IsEnabled = core.BoolPtr(true)
	model.LatencyThresholds = latencyThresholdsModel
	model.MaxConcurrentStreams = core.Float64Ptr(float64(72.5))
	model.NasSourceParams = nasSourceParamsModel
	model.RegisteredSourceMaxConcurrentBackups = core.Float64Ptr(float64(72.5))
	model.StorageArraySnapshotConfig = storageArraySnapshotConfigModel

	result, err := backuprecovery.DataSourceIbmBackupRecoveryRegistrationInfoThrottlingPolicyToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryRegistrationInfoLatencyThresholdsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["active_task_msecs"] = int(26)
		model["new_task_msecs"] = int(26)

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.LatencyThresholds)
	model.ActiveTaskMsecs = core.Int64Ptr(int64(26))
	model.NewTaskMsecs = core.Int64Ptr(int64(26))

	result, err := backuprecovery.DataSourceIbmBackupRecoveryRegistrationInfoLatencyThresholdsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryRegistrationInfoNasSourceParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["max_parallel_metadata_fetch_full_percentage"] = float64(72.5)
		model["max_parallel_metadata_fetch_incremental_percentage"] = float64(72.5)
		model["max_parallel_read_write_full_percentage"] = float64(72.5)
		model["max_parallel_read_write_incremental_percentage"] = float64(72.5)

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.NasSourceParams)
	model.MaxParallelMetadataFetchFullPercentage = core.Float64Ptr(float64(72.5))
	model.MaxParallelMetadataFetchIncrementalPercentage = core.Float64Ptr(float64(72.5))
	model.MaxParallelReadWriteFullPercentage = core.Float64Ptr(float64(72.5))
	model.MaxParallelReadWriteIncrementalPercentage = core.Float64Ptr(float64(72.5))

	result, err := backuprecovery.DataSourceIbmBackupRecoveryRegistrationInfoNasSourceParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryRegistrationInfoStorageArraySnapshotConfigToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		storageArraySnapshotMaxSpaceConfigModel := make(map[string]interface{})
		storageArraySnapshotMaxSpaceConfigModel["max_snapshot_space_percentage"] = float64(72.5)

		maxSnapshotConfigModel := make(map[string]interface{})
		maxSnapshotConfigModel["max_snapshots"] = float64(72.5)

		maxSpaceConfigModel := make(map[string]interface{})
		maxSpaceConfigModel["max_snapshot_space_percentage"] = float64(72.5)

		storageArraySnapshotThrottlingPoliciesModel := make(map[string]interface{})
		storageArraySnapshotThrottlingPoliciesModel["id"] = int(26)
		storageArraySnapshotThrottlingPoliciesModel["is_max_snapshots_config_enabled"] = true
		storageArraySnapshotThrottlingPoliciesModel["is_max_space_config_enabled"] = true
		storageArraySnapshotThrottlingPoliciesModel["max_snapshot_config"] = []map[string]interface{}{maxSnapshotConfigModel}
		storageArraySnapshotThrottlingPoliciesModel["max_space_config"] = []map[string]interface{}{maxSpaceConfigModel}

		model := make(map[string]interface{})
		model["is_max_snapshots_config_enabled"] = true
		model["is_max_space_config_enabled"] = true
		model["storage_array_snapshot_max_space_config"] = []map[string]interface{}{storageArraySnapshotMaxSpaceConfigModel}
		model["storage_array_snapshot_throttling_policies"] = []map[string]interface{}{storageArraySnapshotThrottlingPoliciesModel}

		assert.Equal(t, result, model)
	}

	storageArraySnapshotMaxSpaceConfigModel := new(backuprecoveryv1.StorageArraySnapshotMaxSpaceConfig)
	storageArraySnapshotMaxSpaceConfigModel.MaxSnapshotSpacePercentage = core.Float64Ptr(float64(72.5))

	maxSnapshotConfigModel := new(backuprecoveryv1.MaxSnapshotConfig)
	maxSnapshotConfigModel.MaxSnapshots = core.Float64Ptr(float64(72.5))

	maxSpaceConfigModel := new(backuprecoveryv1.MaxSpaceConfig)
	maxSpaceConfigModel.MaxSnapshotSpacePercentage = core.Float64Ptr(float64(72.5))

	storageArraySnapshotThrottlingPoliciesModel := new(backuprecoveryv1.StorageArraySnapshotThrottlingPolicies)
	storageArraySnapshotThrottlingPoliciesModel.ID = core.Int64Ptr(int64(26))
	storageArraySnapshotThrottlingPoliciesModel.IsMaxSnapshotsConfigEnabled = core.BoolPtr(true)
	storageArraySnapshotThrottlingPoliciesModel.IsMaxSpaceConfigEnabled = core.BoolPtr(true)
	storageArraySnapshotThrottlingPoliciesModel.MaxSnapshotConfig = maxSnapshotConfigModel
	storageArraySnapshotThrottlingPoliciesModel.MaxSpaceConfig = maxSpaceConfigModel

	model := new(backuprecoveryv1.StorageArraySnapshotConfig)
	model.IsMaxSnapshotsConfigEnabled = core.BoolPtr(true)
	model.IsMaxSpaceConfigEnabled = core.BoolPtr(true)
	model.StorageArraySnapshotMaxSpaceConfig = storageArraySnapshotMaxSpaceConfigModel
	model.StorageArraySnapshotThrottlingPolicies = []backuprecoveryv1.StorageArraySnapshotThrottlingPolicies{*storageArraySnapshotThrottlingPoliciesModel}

	result, err := backuprecovery.DataSourceIbmBackupRecoveryRegistrationInfoStorageArraySnapshotConfigToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryRegistrationInfoStorageArraySnapshotMaxSpaceConfigToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["max_snapshot_space_percentage"] = float64(72.5)

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.StorageArraySnapshotMaxSpaceConfig)
	model.MaxSnapshotSpacePercentage = core.Float64Ptr(float64(72.5))

	result, err := backuprecovery.DataSourceIbmBackupRecoveryRegistrationInfoStorageArraySnapshotMaxSpaceConfigToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryRegistrationInfoStorageArraySnapshotThrottlingPoliciesToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		maxSnapshotConfigModel := make(map[string]interface{})
		maxSnapshotConfigModel["max_snapshots"] = float64(72.5)

		maxSpaceConfigModel := make(map[string]interface{})
		maxSpaceConfigModel["max_snapshot_space_percentage"] = float64(72.5)

		model := make(map[string]interface{})
		model["id"] = int(26)
		model["is_max_snapshots_config_enabled"] = true
		model["is_max_space_config_enabled"] = true
		model["max_snapshot_config"] = []map[string]interface{}{maxSnapshotConfigModel}
		model["max_space_config"] = []map[string]interface{}{maxSpaceConfigModel}

		assert.Equal(t, result, model)
	}

	maxSnapshotConfigModel := new(backuprecoveryv1.MaxSnapshotConfig)
	maxSnapshotConfigModel.MaxSnapshots = core.Float64Ptr(float64(72.5))

	maxSpaceConfigModel := new(backuprecoveryv1.MaxSpaceConfig)
	maxSpaceConfigModel.MaxSnapshotSpacePercentage = core.Float64Ptr(float64(72.5))

	model := new(backuprecoveryv1.StorageArraySnapshotThrottlingPolicies)
	model.ID = core.Int64Ptr(int64(26))
	model.IsMaxSnapshotsConfigEnabled = core.BoolPtr(true)
	model.IsMaxSpaceConfigEnabled = core.BoolPtr(true)
	model.MaxSnapshotConfig = maxSnapshotConfigModel
	model.MaxSpaceConfig = maxSpaceConfigModel

	result, err := backuprecovery.DataSourceIbmBackupRecoveryRegistrationInfoStorageArraySnapshotThrottlingPoliciesToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryRegistrationInfoMaxSnapshotConfigToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["max_snapshots"] = float64(72.5)

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.MaxSnapshotConfig)
	model.MaxSnapshots = core.Float64Ptr(float64(72.5))

	result, err := backuprecovery.DataSourceIbmBackupRecoveryRegistrationInfoMaxSnapshotConfigToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryRegistrationInfoMaxSpaceConfigToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["max_snapshot_space_percentage"] = float64(72.5)

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.MaxSpaceConfig)
	model.MaxSnapshotSpacePercentage = core.Float64Ptr(float64(72.5))

	result, err := backuprecovery.DataSourceIbmBackupRecoveryRegistrationInfoMaxSpaceConfigToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryRegistrationInfoThrottlingPolicyOverridesToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		latencyThresholdsModel := make(map[string]interface{})
		latencyThresholdsModel["active_task_msecs"] = int(26)
		latencyThresholdsModel["new_task_msecs"] = int(26)

		nasSourceParamsModel := make(map[string]interface{})
		nasSourceParamsModel["max_parallel_metadata_fetch_full_percentage"] = float64(72.5)
		nasSourceParamsModel["max_parallel_metadata_fetch_incremental_percentage"] = float64(72.5)
		nasSourceParamsModel["max_parallel_read_write_full_percentage"] = float64(72.5)
		nasSourceParamsModel["max_parallel_read_write_incremental_percentage"] = float64(72.5)

		storageArraySnapshotMaxSpaceConfigModel := make(map[string]interface{})
		storageArraySnapshotMaxSpaceConfigModel["max_snapshot_space_percentage"] = float64(72.5)

		maxSnapshotConfigModel := make(map[string]interface{})
		maxSnapshotConfigModel["max_snapshots"] = float64(72.5)

		maxSpaceConfigModel := make(map[string]interface{})
		maxSpaceConfigModel["max_snapshot_space_percentage"] = float64(72.5)

		storageArraySnapshotThrottlingPoliciesModel := make(map[string]interface{})
		storageArraySnapshotThrottlingPoliciesModel["id"] = int(26)
		storageArraySnapshotThrottlingPoliciesModel["is_max_snapshots_config_enabled"] = true
		storageArraySnapshotThrottlingPoliciesModel["is_max_space_config_enabled"] = true
		storageArraySnapshotThrottlingPoliciesModel["max_snapshot_config"] = []map[string]interface{}{maxSnapshotConfigModel}
		storageArraySnapshotThrottlingPoliciesModel["max_space_config"] = []map[string]interface{}{maxSpaceConfigModel}

		storageArraySnapshotConfigModel := make(map[string]interface{})
		storageArraySnapshotConfigModel["is_max_snapshots_config_enabled"] = true
		storageArraySnapshotConfigModel["is_max_space_config_enabled"] = true
		storageArraySnapshotConfigModel["storage_array_snapshot_max_space_config"] = []map[string]interface{}{storageArraySnapshotMaxSpaceConfigModel}
		storageArraySnapshotConfigModel["storage_array_snapshot_throttling_policies"] = []map[string]interface{}{storageArraySnapshotThrottlingPoliciesModel}

		throttlingPolicyModel := make(map[string]interface{})
		throttlingPolicyModel["enforce_max_streams"] = true
		throttlingPolicyModel["enforce_registered_source_max_backups"] = true
		throttlingPolicyModel["is_enabled"] = true
		throttlingPolicyModel["latency_thresholds"] = []map[string]interface{}{latencyThresholdsModel}
		throttlingPolicyModel["max_concurrent_streams"] = float64(72.5)
		throttlingPolicyModel["nas_source_params"] = []map[string]interface{}{nasSourceParamsModel}
		throttlingPolicyModel["registered_source_max_concurrent_backups"] = float64(72.5)
		throttlingPolicyModel["storage_array_snapshot_config"] = []map[string]interface{}{storageArraySnapshotConfigModel}

		model := make(map[string]interface{})
		model["datastore_id"] = int(26)
		model["datastore_name"] = "testString"
		model["throttling_policy"] = []map[string]interface{}{throttlingPolicyModel}

		assert.Equal(t, result, model)
	}

	latencyThresholdsModel := new(backuprecoveryv1.LatencyThresholds)
	latencyThresholdsModel.ActiveTaskMsecs = core.Int64Ptr(int64(26))
	latencyThresholdsModel.NewTaskMsecs = core.Int64Ptr(int64(26))

	nasSourceParamsModel := new(backuprecoveryv1.NasSourceParams)
	nasSourceParamsModel.MaxParallelMetadataFetchFullPercentage = core.Float64Ptr(float64(72.5))
	nasSourceParamsModel.MaxParallelMetadataFetchIncrementalPercentage = core.Float64Ptr(float64(72.5))
	nasSourceParamsModel.MaxParallelReadWriteFullPercentage = core.Float64Ptr(float64(72.5))
	nasSourceParamsModel.MaxParallelReadWriteIncrementalPercentage = core.Float64Ptr(float64(72.5))

	storageArraySnapshotMaxSpaceConfigModel := new(backuprecoveryv1.StorageArraySnapshotMaxSpaceConfig)
	storageArraySnapshotMaxSpaceConfigModel.MaxSnapshotSpacePercentage = core.Float64Ptr(float64(72.5))

	maxSnapshotConfigModel := new(backuprecoveryv1.MaxSnapshotConfig)
	maxSnapshotConfigModel.MaxSnapshots = core.Float64Ptr(float64(72.5))

	maxSpaceConfigModel := new(backuprecoveryv1.MaxSpaceConfig)
	maxSpaceConfigModel.MaxSnapshotSpacePercentage = core.Float64Ptr(float64(72.5))

	storageArraySnapshotThrottlingPoliciesModel := new(backuprecoveryv1.StorageArraySnapshotThrottlingPolicies)
	storageArraySnapshotThrottlingPoliciesModel.ID = core.Int64Ptr(int64(26))
	storageArraySnapshotThrottlingPoliciesModel.IsMaxSnapshotsConfigEnabled = core.BoolPtr(true)
	storageArraySnapshotThrottlingPoliciesModel.IsMaxSpaceConfigEnabled = core.BoolPtr(true)
	storageArraySnapshotThrottlingPoliciesModel.MaxSnapshotConfig = maxSnapshotConfigModel
	storageArraySnapshotThrottlingPoliciesModel.MaxSpaceConfig = maxSpaceConfigModel

	storageArraySnapshotConfigModel := new(backuprecoveryv1.StorageArraySnapshotConfig)
	storageArraySnapshotConfigModel.IsMaxSnapshotsConfigEnabled = core.BoolPtr(true)
	storageArraySnapshotConfigModel.IsMaxSpaceConfigEnabled = core.BoolPtr(true)
	storageArraySnapshotConfigModel.StorageArraySnapshotMaxSpaceConfig = storageArraySnapshotMaxSpaceConfigModel
	storageArraySnapshotConfigModel.StorageArraySnapshotThrottlingPolicies = []backuprecoveryv1.StorageArraySnapshotThrottlingPolicies{*storageArraySnapshotThrottlingPoliciesModel}

	throttlingPolicyModel := new(backuprecoveryv1.ThrottlingPolicy)
	throttlingPolicyModel.EnforceMaxStreams = core.BoolPtr(true)
	throttlingPolicyModel.EnforceRegisteredSourceMaxBackups = core.BoolPtr(true)
	throttlingPolicyModel.IsEnabled = core.BoolPtr(true)
	throttlingPolicyModel.LatencyThresholds = latencyThresholdsModel
	throttlingPolicyModel.MaxConcurrentStreams = core.Float64Ptr(float64(72.5))
	throttlingPolicyModel.NasSourceParams = nasSourceParamsModel
	throttlingPolicyModel.RegisteredSourceMaxConcurrentBackups = core.Float64Ptr(float64(72.5))
	throttlingPolicyModel.StorageArraySnapshotConfig = storageArraySnapshotConfigModel

	model := new(backuprecoveryv1.ThrottlingPolicyOverrides)
	model.DatastoreID = core.Int64Ptr(int64(26))
	model.DatastoreName = core.StringPtr("testString")
	model.ThrottlingPolicy = throttlingPolicyModel

	result, err := backuprecovery.DataSourceIbmBackupRecoveryRegistrationInfoThrottlingPolicyOverridesToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryRegistrationInfoRegisteredSourceVlanConfigToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["vlan"] = float64(72.5)
		model["disable_vlan"] = true
		model["interface_name"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.RegisteredSourceVlanConfig)
	model.Vlan = core.Float64Ptr(float64(72.5))
	model.DisableVlan = core.BoolPtr(true)
	model.InterfaceName = core.StringPtr("testString")

	result, err := backuprecovery.DataSourceIbmBackupRecoveryRegistrationInfoRegisteredSourceVlanConfigToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryRegistrationInfoUniqueGlobalIDToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["cluster_id"] = int(26)
		model["cluster_incarnation_id"] = int(26)
		model["id"] = int(26)

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.UniqueGlobalID)
	model.ClusterID = core.Int64Ptr(int64(26))
	model.ClusterIncarnationID = core.Int64Ptr(int64(26))
	model.ID = core.Int64Ptr(int64(26))

	result, err := backuprecovery.DataSourceIbmBackupRecoveryRegistrationInfoUniqueGlobalIDToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryRegistrationInfoNetworkingInformationToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		clusterNetworkingEndpointModel := make(map[string]interface{})
		clusterNetworkingEndpointModel["fqdn"] = "testString"
		clusterNetworkingEndpointModel["ipv4_addr"] = "testString"
		clusterNetworkingEndpointModel["ipv6_addr"] = "testString"

		clusterNetworkResourceInformationModel := make(map[string]interface{})
		clusterNetworkResourceInformationModel["endpoints"] = []map[string]interface{}{clusterNetworkingEndpointModel}
		clusterNetworkResourceInformationModel["type"] = "testString"

		model := make(map[string]interface{})
		model["resource_vec"] = []map[string]interface{}{clusterNetworkResourceInformationModel}

		assert.Equal(t, result, model)
	}

	clusterNetworkingEndpointModel := new(backuprecoveryv1.ClusterNetworkingEndpoint)
	clusterNetworkingEndpointModel.Fqdn = core.StringPtr("testString")
	clusterNetworkingEndpointModel.Ipv4Addr = core.StringPtr("testString")
	clusterNetworkingEndpointModel.Ipv6Addr = core.StringPtr("testString")

	clusterNetworkResourceInformationModel := new(backuprecoveryv1.ClusterNetworkResourceInformation)
	clusterNetworkResourceInformationModel.Endpoints = []backuprecoveryv1.ClusterNetworkingEndpoint{*clusterNetworkingEndpointModel}
	clusterNetworkResourceInformationModel.Type = core.StringPtr("testString")

	model := new(backuprecoveryv1.NetworkingInformation)
	model.ResourceVec = []backuprecoveryv1.ClusterNetworkResourceInformation{*clusterNetworkResourceInformationModel}

	result, err := backuprecovery.DataSourceIbmBackupRecoveryRegistrationInfoNetworkingInformationToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryRegistrationInfoClusterNetworkResourceInformationToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		clusterNetworkingEndpointModel := make(map[string]interface{})
		clusterNetworkingEndpointModel["fqdn"] = "testString"
		clusterNetworkingEndpointModel["ipv4_addr"] = "testString"
		clusterNetworkingEndpointModel["ipv6_addr"] = "testString"

		model := make(map[string]interface{})
		model["endpoints"] = []map[string]interface{}{clusterNetworkingEndpointModel}
		model["type"] = "testString"

		assert.Equal(t, result, model)
	}

	clusterNetworkingEndpointModel := new(backuprecoveryv1.ClusterNetworkingEndpoint)
	clusterNetworkingEndpointModel.Fqdn = core.StringPtr("testString")
	clusterNetworkingEndpointModel.Ipv4Addr = core.StringPtr("testString")
	clusterNetworkingEndpointModel.Ipv6Addr = core.StringPtr("testString")

	model := new(backuprecoveryv1.ClusterNetworkResourceInformation)
	model.Endpoints = []backuprecoveryv1.ClusterNetworkingEndpoint{*clusterNetworkingEndpointModel}
	model.Type = core.StringPtr("testString")

	result, err := backuprecovery.DataSourceIbmBackupRecoveryRegistrationInfoClusterNetworkResourceInformationToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryRegistrationInfoClusterNetworkingEndpointToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["fqdn"] = "testString"
		model["ipv4_addr"] = "testString"
		model["ipv6_addr"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.ClusterNetworkingEndpoint)
	model.Fqdn = core.StringPtr("testString")
	model.Ipv4Addr = core.StringPtr("testString")
	model.Ipv6Addr = core.StringPtr("testString")

	result, err := backuprecovery.DataSourceIbmBackupRecoveryRegistrationInfoClusterNetworkingEndpointToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryRegistrationInfoPhysicalVolumeToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["device_path"] = "testString"
		model["guid"] = "testString"
		model["is_boot_volume"] = true
		model["is_extended_attributes_supported"] = true
		model["is_protected"] = true
		model["is_shared_volume"] = true
		model["label"] = "testString"
		model["logical_size_bytes"] = float64(72.5)
		model["mount_points"] = []string{"testString"}
		model["mount_type"] = "testString"
		model["network_path"] = "testString"
		model["used_size_bytes"] = float64(72.5)

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.PhysicalVolume)
	model.DevicePath = core.StringPtr("testString")
	model.Guid = core.StringPtr("testString")
	model.IsBootVolume = core.BoolPtr(true)
	model.IsExtendedAttributesSupported = core.BoolPtr(true)
	model.IsProtected = core.BoolPtr(true)
	model.IsSharedVolume = core.BoolPtr(true)
	model.Label = core.StringPtr("testString")
	model.LogicalSizeBytes = core.Float64Ptr(float64(72.5))
	model.MountPoints = []string{"testString"}
	model.MountType = core.StringPtr("testString")
	model.NetworkPath = core.StringPtr("testString")
	model.UsedSizeBytes = core.Float64Ptr(float64(72.5))

	result, err := backuprecovery.DataSourceIbmBackupRecoveryRegistrationInfoPhysicalVolumeToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryRegistrationInfoVssWritersToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["is_writer_excluded"] = true
		model["writer_name"] = true

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.VssWriters)
	model.IsWriterExcluded = core.BoolPtr(true)
	model.WriterName = core.BoolPtr(true)

	result, err := backuprecovery.DataSourceIbmBackupRecoveryRegistrationInfoVssWritersToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryRegistrationInfoKubernetesProtectionSourceToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		vlanParametersModel := make(map[string]interface{})
		vlanParametersModel["disable_vlan"] = true
		vlanParametersModel["interface_name"] = "testString"
		vlanParametersModel["vlan"] = int(38)

		kubernetesLabelAttributeModel := make(map[string]interface{})
		kubernetesLabelAttributeModel["id"] = int(26)
		kubernetesLabelAttributeModel["name"] = "testString"
		kubernetesLabelAttributeModel["uuid"] = "testString"

		k8sLabelModel := make(map[string]interface{})
		k8sLabelModel["key"] = "testString"
		k8sLabelModel["value"] = "testString"

		serviceAnnotationsEntryModel := make(map[string]interface{})
		serviceAnnotationsEntryModel["key"] = "testString"
		serviceAnnotationsEntryModel["value"] = "testString"

		kubernetesStorageClassInfoModel := make(map[string]interface{})
		kubernetesStorageClassInfoModel["name"] = "testString"
		kubernetesStorageClassInfoModel["provisioner"] = "testString"

		kubernetesServiceAnnotationObjectModel := make(map[string]interface{})
		kubernetesServiceAnnotationObjectModel["key"] = "testString"
		kubernetesServiceAnnotationObjectModel["value"] = "testString"

		vlanParamsModel := make(map[string]interface{})
		vlanParamsModel["disable_vlan"] = true
		vlanParamsModel["interface_name"] = "testString"
		vlanParamsModel["vlan_id"] = int(38)

		kubernetesVlanInfoModel := make(map[string]interface{})
		kubernetesVlanInfoModel["service_annotations"] = []map[string]interface{}{kubernetesServiceAnnotationObjectModel}
		kubernetesVlanInfoModel["vlan_params"] = []map[string]interface{}{vlanParamsModel}

		model := make(map[string]interface{})
		model["datamover_image_location"] = "testString"
		model["datamover_service_type"] = int(38)
		model["datamover_upgradability"] = "testString"
		model["default_vlan_params"] = []map[string]interface{}{vlanParametersModel}
		model["description"] = "testString"
		model["distribution"] = "kMainline"
		model["init_container_image_location"] = "testString"
		model["label_attributes"] = []map[string]interface{}{kubernetesLabelAttributeModel}
		model["name"] = "testString"
		model["priority_class_name"] = "testString"
		model["resource_annotation_list"] = []map[string]interface{}{k8sLabelModel}
		model["resource_label_list"] = []map[string]interface{}{k8sLabelModel}
		model["san_field"] = []string{"testString"}
		model["service_annotations"] = []map[string]interface{}{serviceAnnotationsEntryModel}
		model["storage_class"] = []map[string]interface{}{kubernetesStorageClassInfoModel}
		model["type"] = "kCluster"
		model["uuid"] = "testString"
		model["velero_aws_plugin_image_location"] = "testString"
		model["velero_image_location"] = "testString"
		model["velero_openshift_plugin_image_location"] = "testString"
		model["velero_upgradability"] = "testString"
		model["vlan_info_vec"] = []map[string]interface{}{kubernetesVlanInfoModel}

		assert.Equal(t, result, model)
	}

	vlanParametersModel := new(backuprecoveryv1.VlanParameters)
	vlanParametersModel.DisableVlan = core.BoolPtr(true)
	vlanParametersModel.InterfaceName = core.StringPtr("testString")
	vlanParametersModel.Vlan = core.Int64Ptr(int64(38))

	kubernetesLabelAttributeModel := new(backuprecoveryv1.KubernetesLabelAttribute)
	kubernetesLabelAttributeModel.ID = core.Int64Ptr(int64(26))
	kubernetesLabelAttributeModel.Name = core.StringPtr("testString")
	kubernetesLabelAttributeModel.UUID = core.StringPtr("testString")

	k8sLabelModel := new(backuprecoveryv1.K8sLabel)
	k8sLabelModel.Key = core.StringPtr("testString")
	k8sLabelModel.Value = core.StringPtr("testString")

	serviceAnnotationsEntryModel := new(backuprecoveryv1.ServiceAnnotationsEntry)
	serviceAnnotationsEntryModel.Key = core.StringPtr("testString")
	serviceAnnotationsEntryModel.Value = core.StringPtr("testString")

	kubernetesStorageClassInfoModel := new(backuprecoveryv1.KubernetesStorageClassInfo)
	kubernetesStorageClassInfoModel.Name = core.StringPtr("testString")
	kubernetesStorageClassInfoModel.Provisioner = core.StringPtr("testString")

	kubernetesServiceAnnotationObjectModel := new(backuprecoveryv1.KubernetesServiceAnnotationObject)
	kubernetesServiceAnnotationObjectModel.Key = core.StringPtr("testString")
	kubernetesServiceAnnotationObjectModel.Value = core.StringPtr("testString")

	vlanParamsModel := new(backuprecoveryv1.VlanParams)
	vlanParamsModel.DisableVlan = core.BoolPtr(true)
	vlanParamsModel.InterfaceName = core.StringPtr("testString")
	vlanParamsModel.VlanID = core.Int64Ptr(int64(38))

	kubernetesVlanInfoModel := new(backuprecoveryv1.KubernetesVlanInfo)
	kubernetesVlanInfoModel.ServiceAnnotations = []backuprecoveryv1.KubernetesServiceAnnotationObject{*kubernetesServiceAnnotationObjectModel}
	kubernetesVlanInfoModel.VlanParams = vlanParamsModel

	model := new(backuprecoveryv1.KubernetesProtectionSource)
	model.DatamoverImageLocation = core.StringPtr("testString")
	model.DatamoverServiceType = core.Int64Ptr(int64(38))
	model.DatamoverUpgradability = core.StringPtr("kCurrent")
	model.DefaultVlanParams = vlanParametersModel
	model.Description = core.StringPtr("testString")
	model.Distribution = core.StringPtr("kMainline")
	model.InitContainerImageLocation = core.StringPtr("testString")
	model.LabelAttributes = []backuprecoveryv1.KubernetesLabelAttribute{*kubernetesLabelAttributeModel}
	model.Name = core.StringPtr("testString")
	model.PriorityClassName = core.StringPtr("testString")
	model.ResourceAnnotationList = []backuprecoveryv1.K8sLabel{*k8sLabelModel}
	model.ResourceLabelList = []backuprecoveryv1.K8sLabel{*k8sLabelModel}
	model.SanField = []string{"testString"}
	model.ServiceAnnotations = []backuprecoveryv1.ServiceAnnotationsEntry{*serviceAnnotationsEntryModel}
	model.StorageClass = []backuprecoveryv1.KubernetesStorageClassInfo{*kubernetesStorageClassInfoModel}
	model.Type = core.StringPtr("kCluster")
	model.UUID = core.StringPtr("testString")
	model.VeleroAwsPluginImageLocation = core.StringPtr("testString")
	model.VeleroImageLocation = core.StringPtr("testString")
	model.VeleroOpenshiftPluginImageLocation = core.StringPtr("testString")
	model.VeleroUpgradability = core.StringPtr("testString")
	model.VlanInfoVec = []backuprecoveryv1.KubernetesVlanInfo{*kubernetesVlanInfoModel}

	result, err := backuprecovery.DataSourceIbmBackupRecoveryRegistrationInfoKubernetesProtectionSourceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryRegistrationInfoVlanParametersToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["disable_vlan"] = true
		model["interface_name"] = "testString"
		model["vlan"] = int(38)

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.VlanParameters)
	model.DisableVlan = core.BoolPtr(true)
	model.InterfaceName = core.StringPtr("testString")
	model.Vlan = core.Int64Ptr(int64(38))

	result, err := backuprecovery.DataSourceIbmBackupRecoveryRegistrationInfoVlanParametersToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryRegistrationInfoKubernetesLabelAttributeToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["id"] = int(26)
		model["name"] = "testString"
		model["uuid"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.KubernetesLabelAttribute)
	model.ID = core.Int64Ptr(int64(26))
	model.Name = core.StringPtr("testString")
	model.UUID = core.StringPtr("testString")

	result, err := backuprecovery.DataSourceIbmBackupRecoveryRegistrationInfoKubernetesLabelAttributeToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryRegistrationInfoK8sLabelToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["key"] = "testString"
		model["value"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.K8sLabel)
	model.Key = core.StringPtr("testString")
	model.Value = core.StringPtr("testString")

	result, err := backuprecovery.DataSourceIbmBackupRecoveryRegistrationInfoK8sLabelToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryRegistrationInfoServiceAnnotationsEntryToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["key"] = "testString"
		model["value"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.ServiceAnnotationsEntry)
	model.Key = core.StringPtr("testString")
	model.Value = core.StringPtr("testString")

	result, err := backuprecovery.DataSourceIbmBackupRecoveryRegistrationInfoServiceAnnotationsEntryToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryRegistrationInfoKubernetesStorageClassInfoToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["name"] = "testString"
		model["provisioner"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.KubernetesStorageClassInfo)
	model.Name = core.StringPtr("testString")
	model.Provisioner = core.StringPtr("testString")

	result, err := backuprecovery.DataSourceIbmBackupRecoveryRegistrationInfoKubernetesStorageClassInfoToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryRegistrationInfoKubernetesVlanInfoToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		kubernetesServiceAnnotationObjectModel := make(map[string]interface{})
		kubernetesServiceAnnotationObjectModel["key"] = "testString"
		kubernetesServiceAnnotationObjectModel["value"] = "testString"

		vlanParamsModel := make(map[string]interface{})
		vlanParamsModel["disable_vlan"] = true
		vlanParamsModel["interface_name"] = "testString"
		vlanParamsModel["vlan_id"] = int(38)

		model := make(map[string]interface{})
		model["service_annotations"] = []map[string]interface{}{kubernetesServiceAnnotationObjectModel}
		model["vlan_params"] = []map[string]interface{}{vlanParamsModel}

		assert.Equal(t, result, model)
	}

	kubernetesServiceAnnotationObjectModel := new(backuprecoveryv1.KubernetesServiceAnnotationObject)
	kubernetesServiceAnnotationObjectModel.Key = core.StringPtr("testString")
	kubernetesServiceAnnotationObjectModel.Value = core.StringPtr("testString")

	vlanParamsModel := new(backuprecoveryv1.VlanParams)
	vlanParamsModel.DisableVlan = core.BoolPtr(true)
	vlanParamsModel.InterfaceName = core.StringPtr("testString")
	vlanParamsModel.VlanID = core.Int64Ptr(int64(38))

	model := new(backuprecoveryv1.KubernetesVlanInfo)
	model.ServiceAnnotations = []backuprecoveryv1.KubernetesServiceAnnotationObject{*kubernetesServiceAnnotationObjectModel}
	model.VlanParams = vlanParamsModel

	result, err := backuprecovery.DataSourceIbmBackupRecoveryRegistrationInfoKubernetesVlanInfoToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryRegistrationInfoKubernetesServiceAnnotationObjectToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["key"] = "testString"
		model["value"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.KubernetesServiceAnnotationObject)
	model.Key = core.StringPtr("testString")
	model.Value = core.StringPtr("testString")

	result, err := backuprecovery.DataSourceIbmBackupRecoveryRegistrationInfoKubernetesServiceAnnotationObjectToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryRegistrationInfoVlanParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["disable_vlan"] = true
		model["interface_name"] = "testString"
		model["vlan_id"] = int(38)

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.VlanParams)
	model.DisableVlan = core.BoolPtr(true)
	model.InterfaceName = core.StringPtr("testString")
	model.VlanID = core.Int64Ptr(int64(38))

	result, err := backuprecovery.DataSourceIbmBackupRecoveryRegistrationInfoVlanParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryRegistrationInfoSqlProtectionSourceToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		databaseFileInformationModel := make(map[string]interface{})
		databaseFileInformationModel["file_type"] = "kRows"
		databaseFileInformationModel["full_path"] = "testString"
		databaseFileInformationModel["size_bytes"] = int(26)

		sqlSourceIdModel := make(map[string]interface{})
		sqlSourceIdModel["created_date_msecs"] = int(26)
		sqlSourceIdModel["database_id"] = int(26)
		sqlSourceIdModel["instance_id"] = "testString"

		sqlServerInstanceVersionModel := make(map[string]interface{})
		sqlServerInstanceVersionModel["build"] = float64(72.5)
		sqlServerInstanceVersionModel["major_version"] = float64(72.5)
		sqlServerInstanceVersionModel["minor_version"] = float64(72.5)
		sqlServerInstanceVersionModel["revision"] = float64(72.5)
		sqlServerInstanceVersionModel["version_string"] = float64(72.5)

		model := make(map[string]interface{})
		model["is_available_for_vss_backup"] = true
		model["created_timestamp"] = "testString"
		model["database_name"] = "testString"
		model["db_aag_entity_id"] = int(26)
		model["db_aag_name"] = "testString"
		model["db_compatibility_level"] = int(26)
		model["db_file_groups"] = []string{"testString"}
		model["db_files"] = []map[string]interface{}{databaseFileInformationModel}
		model["db_owner_username"] = "testString"
		model["default_database_location"] = "testString"
		model["default_log_location"] = "testString"
		model["id"] = []map[string]interface{}{sqlSourceIdModel}
		model["is_encrypted"] = true
		model["name"] = "testString"
		model["owner_id"] = int(26)
		model["recovery_model"] = "kSimpleRecoveryModel"
		model["sql_server_db_state"] = "kOnline"
		model["sql_server_instance_version"] = []map[string]interface{}{sqlServerInstanceVersionModel}
		model["type"] = "kInstance"

		assert.Equal(t, result, model)
	}

	databaseFileInformationModel := new(backuprecoveryv1.DatabaseFileInformation)
	databaseFileInformationModel.FileType = core.StringPtr("kRows")
	databaseFileInformationModel.FullPath = core.StringPtr("testString")
	databaseFileInformationModel.SizeBytes = core.Int64Ptr(int64(26))

	sqlSourceIdModel := new(backuprecoveryv1.SQLSourceID)
	sqlSourceIdModel.CreatedDateMsecs = core.Int64Ptr(int64(26))
	sqlSourceIdModel.DatabaseID = core.Int64Ptr(int64(26))
	sqlSourceIdModel.InstanceID = core.StringPtr("testString")

	sqlServerInstanceVersionModel := new(backuprecoveryv1.SQLServerInstanceVersion)
	sqlServerInstanceVersionModel.Build = core.Float64Ptr(float64(72.5))
	sqlServerInstanceVersionModel.MajorVersion = core.Float64Ptr(float64(72.5))
	sqlServerInstanceVersionModel.MinorVersion = core.Float64Ptr(float64(72.5))
	sqlServerInstanceVersionModel.Revision = core.Float64Ptr(float64(72.5))
	sqlServerInstanceVersionModel.VersionString = core.Float64Ptr(float64(72.5))

	model := new(backuprecoveryv1.SqlProtectionSource)
	model.IsAvailableForVssBackup = core.BoolPtr(true)
	model.CreatedTimestamp = core.StringPtr("testString")
	model.DatabaseName = core.StringPtr("testString")
	model.DbAagEntityID = core.Int64Ptr(int64(26))
	model.DbAagName = core.StringPtr("testString")
	model.DbCompatibilityLevel = core.Int64Ptr(int64(26))
	model.DbFileGroups = []string{"testString"}
	model.DbFiles = []backuprecoveryv1.DatabaseFileInformation{*databaseFileInformationModel}
	model.DbOwnerUsername = core.StringPtr("testString")
	model.DefaultDatabaseLocation = core.StringPtr("testString")
	model.DefaultLogLocation = core.StringPtr("testString")
	model.ID = sqlSourceIdModel
	model.IsEncrypted = core.BoolPtr(true)
	model.Name = core.StringPtr("testString")
	model.OwnerID = core.Int64Ptr(int64(26))
	model.RecoveryModel = core.StringPtr("kSimpleRecoveryModel")
	model.SqlServerDbState = core.StringPtr("kOnline")
	model.SqlServerInstanceVersion = sqlServerInstanceVersionModel
	model.Type = core.StringPtr("kInstance")

	result, err := backuprecovery.DataSourceIbmBackupRecoveryRegistrationInfoSqlProtectionSourceToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryRegistrationInfoDatabaseFileInformationToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["file_type"] = "kRows"
		model["full_path"] = "testString"
		model["size_bytes"] = int(26)

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.DatabaseFileInformation)
	model.FileType = core.StringPtr("kRows")
	model.FullPath = core.StringPtr("testString")
	model.SizeBytes = core.Int64Ptr(int64(26))

	result, err := backuprecovery.DataSourceIbmBackupRecoveryRegistrationInfoDatabaseFileInformationToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryRegistrationInfoSQLSourceIDToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["created_date_msecs"] = int(26)
		model["database_id"] = int(26)
		model["instance_id"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.SQLSourceID)
	model.CreatedDateMsecs = core.Int64Ptr(int64(26))
	model.DatabaseID = core.Int64Ptr(int64(26))
	model.InstanceID = core.StringPtr("testString")

	result, err := backuprecovery.DataSourceIbmBackupRecoveryRegistrationInfoSQLSourceIDToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryRegistrationInfoSQLServerInstanceVersionToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["build"] = float64(72.5)
		model["major_version"] = float64(72.5)
		model["minor_version"] = float64(72.5)
		model["revision"] = float64(72.5)
		model["version_string"] = float64(72.5)

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.SQLServerInstanceVersion)
	model.Build = core.Float64Ptr(float64(72.5))
	model.MajorVersion = core.Float64Ptr(float64(72.5))
	model.MinorVersion = core.Float64Ptr(float64(72.5))
	model.Revision = core.Float64Ptr(float64(72.5))
	model.VersionString = core.Float64Ptr(float64(72.5))

	result, err := backuprecovery.DataSourceIbmBackupRecoveryRegistrationInfoSQLServerInstanceVersionToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryRegistrationInfoEntityPermissionInformationToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		groupInfoModel := make(map[string]interface{})
		groupInfoModel["domain"] = "testString"
		groupInfoModel["group_name"] = "testString"
		groupInfoModel["sid"] = "testString"
		groupInfoModel["tenant_ids"] = []string{"testString"}

		tenantInfoModel := make(map[string]interface{})
		tenantInfoModel["bifrost_enabled"] = true
		tenantInfoModel["is_managed_on_helios"] = true
		tenantInfoModel["name"] = "testString"
		tenantInfoModel["tenant_id"] = "testString"

		userInfoModel := make(map[string]interface{})
		userInfoModel["domain"] = "testString"
		userInfoModel["sid"] = "testString"
		userInfoModel["tenant_id"] = "testString"
		userInfoModel["user_name"] = "testString"

		model := make(map[string]interface{})
		model["entity_id"] = int(26)
		model["groups"] = []map[string]interface{}{groupInfoModel}
		model["is_inferred"] = true
		model["is_registered_by_sp"] = true
		model["registering_tenant_id"] = "testString"
		model["tenant"] = []map[string]interface{}{tenantInfoModel}
		model["users"] = []map[string]interface{}{userInfoModel}

		assert.Equal(t, result, model)
	}

	groupInfoModel := new(backuprecoveryv1.GroupInfo)
	groupInfoModel.Domain = core.StringPtr("testString")
	groupInfoModel.GroupName = core.StringPtr("testString")
	groupInfoModel.Sid = core.StringPtr("testString")
	groupInfoModel.TenantIds = []string{"testString"}

	tenantInfoModel := new(backuprecoveryv1.TenantInfo)
	tenantInfoModel.BifrostEnabled = core.BoolPtr(true)
	tenantInfoModel.IsManagedOnHelios = core.BoolPtr(true)
	tenantInfoModel.Name = core.StringPtr("testString")
	tenantInfoModel.TenantID = core.StringPtr("testString")

	userInfoModel := new(backuprecoveryv1.UserInfo)
	userInfoModel.Domain = core.StringPtr("testString")
	userInfoModel.Sid = core.StringPtr("testString")
	userInfoModel.TenantID = core.StringPtr("testString")
	userInfoModel.UserName = core.StringPtr("testString")

	model := new(backuprecoveryv1.EntityPermissionInformation)
	model.EntityID = core.Int64Ptr(int64(26))
	model.Groups = []backuprecoveryv1.GroupInfo{*groupInfoModel}
	model.IsInferred = core.BoolPtr(true)
	model.IsRegisteredBySp = core.BoolPtr(true)
	model.RegisteringTenantID = core.StringPtr("testString")
	model.Tenant = tenantInfoModel
	model.Users = []backuprecoveryv1.UserInfo{*userInfoModel}

	result, err := backuprecovery.DataSourceIbmBackupRecoveryRegistrationInfoEntityPermissionInformationToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryRegistrationInfoGroupInfoToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["domain"] = "testString"
		model["group_name"] = "testString"
		model["sid"] = "testString"
		model["tenant_ids"] = []string{"testString"}

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.GroupInfo)
	model.Domain = core.StringPtr("testString")
	model.GroupName = core.StringPtr("testString")
	model.Sid = core.StringPtr("testString")
	model.TenantIds = []string{"testString"}

	result, err := backuprecovery.DataSourceIbmBackupRecoveryRegistrationInfoGroupInfoToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryRegistrationInfoTenantInfoToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["bifrost_enabled"] = true
		model["is_managed_on_helios"] = true
		model["name"] = "testString"
		model["tenant_id"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.TenantInfo)
	model.BifrostEnabled = core.BoolPtr(true)
	model.IsManagedOnHelios = core.BoolPtr(true)
	model.Name = core.StringPtr("testString")
	model.TenantID = core.StringPtr("testString")

	result, err := backuprecovery.DataSourceIbmBackupRecoveryRegistrationInfoTenantInfoToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryRegistrationInfoUserInfoToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["domain"] = "testString"
		model["sid"] = "testString"
		model["tenant_id"] = "testString"
		model["user_name"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.UserInfo)
	model.Domain = core.StringPtr("testString")
	model.Sid = core.StringPtr("testString")
	model.TenantID = core.StringPtr("testString")
	model.UserName = core.StringPtr("testString")

	result, err := backuprecovery.DataSourceIbmBackupRecoveryRegistrationInfoUserInfoToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryRegistrationInfoMaintenanceModeConfigToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		timeRangeUsecsModel := make(map[string]interface{})
		timeRangeUsecsModel["end_time_usecs"] = int(26)
		timeRangeUsecsModel["start_time_usecs"] = int(26)

		timeModel := make(map[string]interface{})
		timeModel["hour"] = int(38)
		timeModel["minute"] = int(38)

		timeWindowModel := make(map[string]interface{})
		timeWindowModel["day_of_the_week"] = "Sunday"
		timeWindowModel["end_time"] = []map[string]interface{}{timeModel}
		timeWindowModel["start_time"] = []map[string]interface{}{timeModel}

		scheduleModel := make(map[string]interface{})
		scheduleModel["periodic_time_windows"] = []map[string]interface{}{timeWindowModel}
		scheduleModel["schedule_type"] = "PeriodicTimeWindows"
		scheduleModel["time_ranges"] = []map[string]interface{}{timeRangeUsecsModel}
		scheduleModel["timezone"] = "testString"

		workflowInterventionSpecModel := make(map[string]interface{})
		workflowInterventionSpecModel["intervention"] = "NoIntervention"
		workflowInterventionSpecModel["workflow_type"] = "BackupRun"

		model := make(map[string]interface{})
		model["activation_time_intervals"] = []map[string]interface{}{timeRangeUsecsModel}
		model["maintenance_schedule"] = []map[string]interface{}{scheduleModel}
		model["user_message"] = "testString"
		model["workflow_intervention_spec_list"] = []map[string]interface{}{workflowInterventionSpecModel}

		assert.Equal(t, result, model)
	}

	timeRangeUsecsModel := new(backuprecoveryv1.TimeRangeUsecs)
	timeRangeUsecsModel.EndTimeUsecs = core.Int64Ptr(int64(26))
	timeRangeUsecsModel.StartTimeUsecs = core.Int64Ptr(int64(26))

	timeModel := new(backuprecoveryv1.Time)
	timeModel.Hour = core.Int64Ptr(int64(38))
	timeModel.Minute = core.Int64Ptr(int64(38))

	timeWindowModel := new(backuprecoveryv1.TimeWindow)
	timeWindowModel.DayOfTheWeek = core.StringPtr("Sunday")
	timeWindowModel.EndTime = timeModel
	timeWindowModel.StartTime = timeModel

	scheduleModel := new(backuprecoveryv1.Schedule)
	scheduleModel.PeriodicTimeWindows = []backuprecoveryv1.TimeWindow{*timeWindowModel}
	scheduleModel.ScheduleType = core.StringPtr("PeriodicTimeWindows")
	scheduleModel.TimeRanges = []backuprecoveryv1.TimeRangeUsecs{*timeRangeUsecsModel}
	scheduleModel.Timezone = core.StringPtr("testString")

	workflowInterventionSpecModel := new(backuprecoveryv1.WorkflowInterventionSpec)
	workflowInterventionSpecModel.Intervention = core.StringPtr("NoIntervention")
	workflowInterventionSpecModel.WorkflowType = core.StringPtr("BackupRun")

	model := new(backuprecoveryv1.MaintenanceModeConfig)
	model.ActivationTimeIntervals = []backuprecoveryv1.TimeRangeUsecs{*timeRangeUsecsModel}
	model.MaintenanceSchedule = scheduleModel
	model.UserMessage = core.StringPtr("testString")
	model.WorkflowInterventionSpecList = []backuprecoveryv1.WorkflowInterventionSpec{*workflowInterventionSpecModel}

	result, err := backuprecovery.DataSourceIbmBackupRecoveryRegistrationInfoMaintenanceModeConfigToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryRegistrationInfoTimeRangeUsecsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["end_time_usecs"] = int(26)
		model["start_time_usecs"] = int(26)

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.TimeRangeUsecs)
	model.EndTimeUsecs = core.Int64Ptr(int64(26))
	model.StartTimeUsecs = core.Int64Ptr(int64(26))

	result, err := backuprecovery.DataSourceIbmBackupRecoveryRegistrationInfoTimeRangeUsecsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryRegistrationInfoScheduleToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		timeModel := make(map[string]interface{})
		timeModel["hour"] = int(38)
		timeModel["minute"] = int(38)

		timeWindowModel := make(map[string]interface{})
		timeWindowModel["day_of_the_week"] = "Sunday"
		timeWindowModel["end_time"] = []map[string]interface{}{timeModel}
		timeWindowModel["start_time"] = []map[string]interface{}{timeModel}

		timeRangeUsecsModel := make(map[string]interface{})
		timeRangeUsecsModel["end_time_usecs"] = int(26)
		timeRangeUsecsModel["start_time_usecs"] = int(26)

		model := make(map[string]interface{})
		model["periodic_time_windows"] = []map[string]interface{}{timeWindowModel}
		model["schedule_type"] = "PeriodicTimeWindows"
		model["time_ranges"] = []map[string]interface{}{timeRangeUsecsModel}
		model["timezone"] = "testString"

		assert.Equal(t, result, model)
	}

	timeModel := new(backuprecoveryv1.Time)
	timeModel.Hour = core.Int64Ptr(int64(38))
	timeModel.Minute = core.Int64Ptr(int64(38))

	timeWindowModel := new(backuprecoveryv1.TimeWindow)
	timeWindowModel.DayOfTheWeek = core.StringPtr("Sunday")
	timeWindowModel.EndTime = timeModel
	timeWindowModel.StartTime = timeModel

	timeRangeUsecsModel := new(backuprecoveryv1.TimeRangeUsecs)
	timeRangeUsecsModel.EndTimeUsecs = core.Int64Ptr(int64(26))
	timeRangeUsecsModel.StartTimeUsecs = core.Int64Ptr(int64(26))

	model := new(backuprecoveryv1.Schedule)
	model.PeriodicTimeWindows = []backuprecoveryv1.TimeWindow{*timeWindowModel}
	model.ScheduleType = core.StringPtr("PeriodicTimeWindows")
	model.TimeRanges = []backuprecoveryv1.TimeRangeUsecs{*timeRangeUsecsModel}
	model.Timezone = core.StringPtr("testString")

	result, err := backuprecovery.DataSourceIbmBackupRecoveryRegistrationInfoScheduleToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryRegistrationInfoTimeWindowToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		timeModel := make(map[string]interface{})
		timeModel["hour"] = int(38)
		timeModel["minute"] = int(38)

		model := make(map[string]interface{})
		model["day_of_the_week"] = "Sunday"
		model["end_time"] = []map[string]interface{}{timeModel}
		model["start_time"] = []map[string]interface{}{timeModel}

		assert.Equal(t, result, model)
	}

	timeModel := new(backuprecoveryv1.Time)
	timeModel.Hour = core.Int64Ptr(int64(38))
	timeModel.Minute = core.Int64Ptr(int64(38))

	model := new(backuprecoveryv1.TimeWindow)
	model.DayOfTheWeek = core.StringPtr("Sunday")
	model.EndTime = timeModel
	model.StartTime = timeModel

	result, err := backuprecovery.DataSourceIbmBackupRecoveryRegistrationInfoTimeWindowToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryRegistrationInfoWorkflowInterventionSpecToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["intervention"] = "NoIntervention"
		model["workflow_type"] = "BackupRun"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.WorkflowInterventionSpec)
	model.Intervention = core.StringPtr("NoIntervention")
	model.WorkflowType = core.StringPtr("BackupRun")

	result, err := backuprecovery.DataSourceIbmBackupRecoveryRegistrationInfoWorkflowInterventionSpecToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryRegistrationInfoRegisteredSourceInfoToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		connectorParametersModel := make(map[string]interface{})
		connectorParametersModel["connection_id"] = int(26)
		connectorParametersModel["connector_group_id"] = int(26)
		connectorParametersModel["endpoint"] = "testString"
		connectorParametersModel["environment"] = "kVMware"
		connectorParametersModel["id"] = int(26)
		connectorParametersModel["version"] = int(26)

		cassandraPortsInfoModel := make(map[string]interface{})
		cassandraPortsInfoModel["jmx_port"] = int(38)
		cassandraPortsInfoModel["native_transport_port"] = int(38)
		cassandraPortsInfoModel["rpc_port"] = int(38)
		cassandraPortsInfoModel["ssl_storage_port"] = int(38)
		cassandraPortsInfoModel["storage_port"] = int(38)

		cassandraSecurityInfoModel := make(map[string]interface{})
		cassandraSecurityInfoModel["cassandra_auth_required"] = true
		cassandraSecurityInfoModel["cassandra_auth_type"] = "PASSWORD"
		cassandraSecurityInfoModel["cassandra_authorizer"] = "testString"
		cassandraSecurityInfoModel["client_encryption"] = true
		cassandraSecurityInfoModel["dse_authorization"] = true
		cassandraSecurityInfoModel["server_encryption_req_client_auth"] = true
		cassandraSecurityInfoModel["server_internode_encryption_type"] = "testString"

		cassandraConnectParamsModel := make(map[string]interface{})
		cassandraConnectParamsModel["cassandra_ports_info"] = []map[string]interface{}{cassandraPortsInfoModel}
		cassandraConnectParamsModel["cassandra_security_info"] = []map[string]interface{}{cassandraSecurityInfoModel}
		cassandraConnectParamsModel["cassandra_version"] = "testString"
		cassandraConnectParamsModel["commit_log_backup_location"] = "testString"
		cassandraConnectParamsModel["config_directory"] = "testString"
		cassandraConnectParamsModel["data_centers"] = []string{"testString"}
		cassandraConnectParamsModel["dse_config_directory"] = "testString"
		cassandraConnectParamsModel["dse_version"] = "testString"
		cassandraConnectParamsModel["is_dse_authenticator"] = true
		cassandraConnectParamsModel["is_dse_tiered_storage"] = true
		cassandraConnectParamsModel["is_jmx_auth_enable"] = true
		cassandraConnectParamsModel["kerberos_principal"] = "testString"
		cassandraConnectParamsModel["primary_host"] = "testString"
		cassandraConnectParamsModel["seeds"] = []string{"testString"}
		cassandraConnectParamsModel["solr_nodes"] = []string{"testString"}
		cassandraConnectParamsModel["solr_port"] = int(38)

		couchbaseConnectParamsModel := make(map[string]interface{})
		couchbaseConnectParamsModel["carrier_direct_port"] = int(38)
		couchbaseConnectParamsModel["http_direct_port"] = int(38)
		couchbaseConnectParamsModel["requires_ssl"] = true
		couchbaseConnectParamsModel["seeds"] = []string{"testString"}

		hadoopDiscoveryParamsModel := make(map[string]interface{})
		hadoopDiscoveryParamsModel["config_directory"] = "testString"
		hadoopDiscoveryParamsModel["host"] = "testString"

		hBaseConnectParamsModel := make(map[string]interface{})
		hBaseConnectParamsModel["hbase_discovery_params"] = []map[string]interface{}{hadoopDiscoveryParamsModel}
		hBaseConnectParamsModel["hdfs_entity_id"] = int(26)
		hBaseConnectParamsModel["kerberos_principal"] = "testString"
		hBaseConnectParamsModel["root_data_directory"] = "testString"
		hBaseConnectParamsModel["zookeeper_quorum"] = []string{"testString"}

		hdfsConnectParamsModel := make(map[string]interface{})
		hdfsConnectParamsModel["hadoop_distribution"] = "CDH"
		hdfsConnectParamsModel["hadoop_version"] = "testString"
		hdfsConnectParamsModel["hdfs_connection_type"] = "DFS"
		hdfsConnectParamsModel["hdfs_discovery_params"] = []map[string]interface{}{hadoopDiscoveryParamsModel}
		hdfsConnectParamsModel["kerberos_principal"] = "testString"
		hdfsConnectParamsModel["namenode"] = "testString"
		hdfsConnectParamsModel["port"] = int(38)

		hiveConnectParamsModel := make(map[string]interface{})
		hiveConnectParamsModel["entity_threshold_exceeded"] = true
		hiveConnectParamsModel["hdfs_entity_id"] = int(26)
		hiveConnectParamsModel["hive_discovery_params"] = []map[string]interface{}{hadoopDiscoveryParamsModel}
		hiveConnectParamsModel["kerberos_principal"] = "testString"
		hiveConnectParamsModel["metastore"] = "testString"
		hiveConnectParamsModel["thrift_port"] = int(38)

		networkPoolConfigModel := make(map[string]interface{})
		networkPoolConfigModel["pool_name"] = "testString"
		networkPoolConfigModel["subnet"] = "testString"
		networkPoolConfigModel["use_smart_connect"] = true

		zoneConfigModel := make(map[string]interface{})
		zoneConfigModel["dynamic_network_pool_config"] = []map[string]interface{}{networkPoolConfigModel}

		registeredProtectionSourceIsilonParamsModel := make(map[string]interface{})
		registeredProtectionSourceIsilonParamsModel["zone_config_list"] = []map[string]interface{}{zoneConfigModel}

		mongoDbConnectParamsModel := make(map[string]interface{})
		mongoDbConnectParamsModel["auth_type"] = "SCRAM"
		mongoDbConnectParamsModel["authenticating_database_name"] = "testString"
		mongoDbConnectParamsModel["requires_ssl"] = true
		mongoDbConnectParamsModel["secondary_node_tag"] = "testString"
		mongoDbConnectParamsModel["seeds"] = []string{"testString"}
		mongoDbConnectParamsModel["use_fixed_node_for_backup"] = true
		mongoDbConnectParamsModel["use_secondary_for_backup"] = true

		nasServerCredentialsModel := make(map[string]interface{})
		nasServerCredentialsModel["domain"] = "testString"
		nasServerCredentialsModel["nas_protocol"] = "kNoProtocol"

		sitesDiscoveryParamsModel := make(map[string]interface{})
		sitesDiscoveryParamsModel["enable_site_tagging"] = true

		teamsAdditionalParamsModel := make(map[string]interface{})
		teamsAdditionalParamsModel["allow_posts_backup"] = true

		usersDiscoveryParamsModel := make(map[string]interface{})
		usersDiscoveryParamsModel["allow_chats_backup"] = true
		usersDiscoveryParamsModel["discover_users_with_mailbox"] = true
		usersDiscoveryParamsModel["discover_users_with_onedrive"] = true
		usersDiscoveryParamsModel["fetch_mailbox_info"] = true
		usersDiscoveryParamsModel["fetch_one_drive_info"] = true
		usersDiscoveryParamsModel["skip_users_without_my_site"] = true

		objectsDiscoveryParamsModel := make(map[string]interface{})
		objectsDiscoveryParamsModel["discoverable_object_type_list"] = []string{"testString"}
		objectsDiscoveryParamsModel["sites_discovery_params"] = []map[string]interface{}{sitesDiscoveryParamsModel}
		objectsDiscoveryParamsModel["teams_additional_params"] = []map[string]interface{}{teamsAdditionalParamsModel}
		objectsDiscoveryParamsModel["users_discovery_params"] = []map[string]interface{}{usersDiscoveryParamsModel}

		m365CsmParamsModel := make(map[string]interface{})
		m365CsmParamsModel["backup_allowed"] = true

		o365ConnectParamsModel := make(map[string]interface{})
		o365ConnectParamsModel["objects_discovery_params"] = []map[string]interface{}{objectsDiscoveryParamsModel}
		o365ConnectParamsModel["csm_params"] = []map[string]interface{}{m365CsmParamsModel}

		office365CredentialsModel := make(map[string]interface{})
		office365CredentialsModel["client_id"] = "testString"
		office365CredentialsModel["client_secret"] = "testString"
		office365CredentialsModel["grant_type"] = "testString"
		office365CredentialsModel["scope"] = "testString"
		office365CredentialsModel["use_o_auth_for_exchange_online"] = true

		credentialsModel := make(map[string]interface{})
		credentialsModel["username"] = "testString"
		credentialsModel["password"] = "testString"

		timeModel := make(map[string]interface{})
		timeModel["hour"] = int(38)
		timeModel["minute"] = int(38)

		dayTimeParamsModel := make(map[string]interface{})
		dayTimeParamsModel["day"] = "kSunday"
		dayTimeParamsModel["time"] = []map[string]interface{}{timeModel}

		dayTimeWindowModel := make(map[string]interface{})
		dayTimeWindowModel["end_time"] = []map[string]interface{}{dayTimeParamsModel}
		dayTimeWindowModel["start_time"] = []map[string]interface{}{dayTimeParamsModel}

		throttlingWindowModel := make(map[string]interface{})
		throttlingWindowModel["day_time_window"] = []map[string]interface{}{dayTimeWindowModel}
		throttlingWindowModel["threshold"] = int(26)

		throttlingConfigurationModel := make(map[string]interface{})
		throttlingConfigurationModel["fixed_threshold"] = int(26)
		throttlingConfigurationModel["pattern_type"] = "kNoThrottling"
		throttlingConfigurationModel["throttling_windows"] = []map[string]interface{}{throttlingWindowModel}

		sourceThrottlingConfigurationModel := make(map[string]interface{})
		sourceThrottlingConfigurationModel["cpu_throttling_config"] = []map[string]interface{}{throttlingConfigurationModel}
		sourceThrottlingConfigurationModel["network_throttling_config"] = []map[string]interface{}{throttlingConfigurationModel}

		physicalParamsModel := make(map[string]interface{})
		physicalParamsModel["applications"] = []string{"kVMware"}
		physicalParamsModel["password"] = "testString"
		physicalParamsModel["throttling_config"] = []map[string]interface{}{sourceThrottlingConfigurationModel}
		physicalParamsModel["username"] = "testString"

		hostSettingsCheckResultModel := make(map[string]interface{})
		hostSettingsCheckResultModel["check_type"] = "kIsAgentPortAccessible"
		hostSettingsCheckResultModel["result_type"] = "kPass"
		hostSettingsCheckResultModel["user_message"] = "testString"

		registeredAppInfoModel := make(map[string]interface{})
		registeredAppInfoModel["authentication_error_message"] = "testString"
		registeredAppInfoModel["authentication_status"] = "kPending"
		registeredAppInfoModel["environment"] = "kPhysical"
		registeredAppInfoModel["host_settings_check_results"] = []map[string]interface{}{hostSettingsCheckResultModel}
		registeredAppInfoModel["refresh_error_message"] = "testString"

		sfdcParamsModel := make(map[string]interface{})
		sfdcParamsModel["access_token"] = "testString"
		sfdcParamsModel["concurrent_api_requests_limit"] = int(26)
		sfdcParamsModel["consumer_key"] = "testString"
		sfdcParamsModel["consumer_secret"] = "testString"
		sfdcParamsModel["daily_api_limit"] = int(26)
		sfdcParamsModel["endpoint"] = "testString"
		sfdcParamsModel["endpoint_type"] = "PROD"
		sfdcParamsModel["metadata_endpoint_url"] = "testString"
		sfdcParamsModel["refresh_token"] = "testString"
		sfdcParamsModel["soap_endpoint_url"] = "testString"
		sfdcParamsModel["use_bulk_api"] = true

		subnetModel := make(map[string]interface{})
		subnetModel["component"] = "testString"
		subnetModel["description"] = "testString"
		subnetModel["id"] = float64(72.5)
		subnetModel["ip"] = "testString"
		subnetModel["netmask_bits"] = float64(72.5)
		subnetModel["netmask_ip4"] = "testString"
		subnetModel["nfs_access"] = "kDisabled"
		subnetModel["nfs_all_squash"] = true
		subnetModel["nfs_root_squash"] = true
		subnetModel["s3_access"] = "kDisabled"
		subnetModel["smb_access"] = "kDisabled"
		subnetModel["tenant_id"] = "testString"

		latencyThresholdsModel := make(map[string]interface{})
		latencyThresholdsModel["active_task_msecs"] = int(26)
		latencyThresholdsModel["new_task_msecs"] = int(26)

		nasSourceThrottlingParamsModel := make(map[string]interface{})
		nasSourceThrottlingParamsModel["max_parallel_metadata_fetch_full_percentage"] = int(38)
		nasSourceThrottlingParamsModel["max_parallel_metadata_fetch_incremental_percentage"] = int(38)
		nasSourceThrottlingParamsModel["max_parallel_read_write_full_percentage"] = int(38)
		nasSourceThrottlingParamsModel["max_parallel_read_write_incremental_percentage"] = int(38)

		storageArraySnapshotMaxSpaceConfigParamsModel := make(map[string]interface{})
		storageArraySnapshotMaxSpaceConfigParamsModel["max_snapshot_space_percentage"] = int(38)

		storageArraySnapshotMaxSnapshotConfigParamsModel := make(map[string]interface{})
		storageArraySnapshotMaxSnapshotConfigParamsModel["max_snapshots"] = int(38)

		storageArraySnapshotThrottlingPolicyModel := make(map[string]interface{})
		storageArraySnapshotThrottlingPolicyModel["id"] = int(26)
		storageArraySnapshotThrottlingPolicyModel["is_max_snapshots_config_enabled"] = true
		storageArraySnapshotThrottlingPolicyModel["is_max_space_config_enabled"] = true
		storageArraySnapshotThrottlingPolicyModel["max_snapshot_config"] = []map[string]interface{}{storageArraySnapshotMaxSnapshotConfigParamsModel}
		storageArraySnapshotThrottlingPolicyModel["max_space_config"] = []map[string]interface{}{storageArraySnapshotMaxSpaceConfigParamsModel}

		storageArraySnapshotConfigParamsModel := make(map[string]interface{})
		storageArraySnapshotConfigParamsModel["is_max_snapshots_config_enabled"] = true
		storageArraySnapshotConfigParamsModel["is_max_space_config_enabled"] = true
		storageArraySnapshotConfigParamsModel["storage_array_snapshot_max_space_config"] = []map[string]interface{}{storageArraySnapshotMaxSpaceConfigParamsModel}
		storageArraySnapshotConfigParamsModel["storage_array_snapshot_throttling_policies"] = []map[string]interface{}{storageArraySnapshotThrottlingPolicyModel}

		throttlingPolicyParametersModel := make(map[string]interface{})
		throttlingPolicyParametersModel["enforce_max_streams"] = true
		throttlingPolicyParametersModel["enforce_registered_source_max_backups"] = true
		throttlingPolicyParametersModel["is_enabled"] = true
		throttlingPolicyParametersModel["latency_thresholds"] = []map[string]interface{}{latencyThresholdsModel}
		throttlingPolicyParametersModel["max_concurrent_streams"] = int(38)
		throttlingPolicyParametersModel["nas_source_params"] = []map[string]interface{}{nasSourceThrottlingParamsModel}
		throttlingPolicyParametersModel["registered_source_max_concurrent_backups"] = int(38)
		throttlingPolicyParametersModel["storage_array_snapshot_config"] = []map[string]interface{}{storageArraySnapshotConfigParamsModel}

		throttlingPolicyOverrideModel := make(map[string]interface{})
		throttlingPolicyOverrideModel["datastore_id"] = int(26)
		throttlingPolicyOverrideModel["datastore_name"] = "testString"
		throttlingPolicyOverrideModel["throttling_policy"] = []map[string]interface{}{throttlingPolicyParametersModel}

		udaSourceCapabilitiesModel := make(map[string]interface{})
		udaSourceCapabilitiesModel["auto_log_backup"] = true
		udaSourceCapabilitiesModel["dynamic_config"] = true
		udaSourceCapabilitiesModel["entity_support"] = true
		udaSourceCapabilitiesModel["et_log_backup"] = true
		udaSourceCapabilitiesModel["external_disks"] = true
		udaSourceCapabilitiesModel["full_backup"] = true
		udaSourceCapabilitiesModel["incr_backup"] = true
		udaSourceCapabilitiesModel["log_backup"] = true
		udaSourceCapabilitiesModel["multi_object_restore"] = true
		udaSourceCapabilitiesModel["pause_resume_backup"] = true
		udaSourceCapabilitiesModel["post_backup_job_script"] = true
		udaSourceCapabilitiesModel["post_restore_job_script"] = true
		udaSourceCapabilitiesModel["pre_backup_job_script"] = true
		udaSourceCapabilitiesModel["pre_restore_job_script"] = true
		udaSourceCapabilitiesModel["resource_throttling"] = true
		udaSourceCapabilitiesModel["snapfs_cert"] = true

		keyValueStrPairModel := make(map[string]interface{})
		keyValueStrPairModel["key"] = "testString"
		keyValueStrPairModel["value"] = "testString"

		udaConnectParamsModel := make(map[string]interface{})
		udaConnectParamsModel["capabilities"] = []map[string]interface{}{udaSourceCapabilitiesModel}
		udaConnectParamsModel["credentials"] = []map[string]interface{}{credentialsModel}
		udaConnectParamsModel["et_enable_log_backup_policy"] = true
		udaConnectParamsModel["et_enable_run_now"] = true
		udaConnectParamsModel["host_type"] = "kLinux"
		udaConnectParamsModel["hosts"] = []string{"testString"}
		udaConnectParamsModel["live_data_view"] = true
		udaConnectParamsModel["live_log_view"] = true
		udaConnectParamsModel["mount_dir"] = "testString"
		udaConnectParamsModel["mount_view"] = true
		udaConnectParamsModel["script_dir"] = "testString"
		udaConnectParamsModel["source_args"] = "testString"
		udaConnectParamsModel["source_registration_arguments"] = []map[string]interface{}{keyValueStrPairModel}
		udaConnectParamsModel["source_type"] = "testString"

		vlanParametersModel := make(map[string]interface{})
		vlanParametersModel["disable_vlan"] = true
		vlanParametersModel["interface_name"] = "testString"
		vlanParametersModel["vlan"] = int(38)

		model := make(map[string]interface{})
		model["access_info"] = []map[string]interface{}{connectorParametersModel}
		model["allowed_ip_addresses"] = []string{"testString"}
		model["authentication_error_message"] = "testString"
		model["authentication_status"] = "kPending"
		model["blacklisted_ip_addresses"] = []string{"testString"}
		model["cassandra_params"] = []map[string]interface{}{cassandraConnectParamsModel}
		model["couchbase_params"] = []map[string]interface{}{couchbaseConnectParamsModel}
		model["denied_ip_addresses"] = []string{"testString"}
		model["environments"] = []string{"kVMware"}
		model["hbase_params"] = []map[string]interface{}{hBaseConnectParamsModel}
		model["hdfs_params"] = []map[string]interface{}{hdfsConnectParamsModel}
		model["hive_params"] = []map[string]interface{}{hiveConnectParamsModel}
		model["is_db_authenticated"] = true
		model["is_storage_array_snapshot_enabled"] = true
		model["isilon_params"] = []map[string]interface{}{registeredProtectionSourceIsilonParamsModel}
		model["link_vms_across_vcenter"] = true
		model["minimum_free_space_gb"] = int(26)
		model["minimum_free_space_percent"] = int(26)
		model["mongodb_params"] = []map[string]interface{}{mongoDbConnectParamsModel}
		model["nas_mount_credentials"] = []map[string]interface{}{nasServerCredentialsModel}
		model["o365_params"] = []map[string]interface{}{o365ConnectParamsModel}
		model["office365_credentials_list"] = []map[string]interface{}{office365CredentialsModel}
		model["office365_region"] = "testString"
		model["office365_service_account_credentials_list"] = []map[string]interface{}{credentialsModel}
		model["password"] = "testString"
		model["physical_params"] = []map[string]interface{}{physicalParamsModel}
		model["progress_monitor_path"] = "testString"
		model["refresh_error_message"] = "testString"
		model["refresh_time_usecs"] = int(26)
		model["registered_apps_info"] = []map[string]interface{}{registeredAppInfoModel}
		model["registration_time_usecs"] = int(26)
		model["sfdc_params"] = []map[string]interface{}{sfdcParamsModel}
		model["subnets"] = []map[string]interface{}{subnetModel}
		model["throttling_policy"] = []map[string]interface{}{throttlingPolicyParametersModel}
		model["throttling_policy_overrides"] = []map[string]interface{}{throttlingPolicyOverrideModel}
		model["uda_params"] = []map[string]interface{}{udaConnectParamsModel}
		model["update_last_backup_details"] = true
		model["use_o_auth_for_exchange_online"] = true
		model["use_vm_bios_uuid"] = true
		model["user_messages"] = []string{"testString"}
		model["username"] = "testString"
		model["vlan_params"] = []map[string]interface{}{vlanParametersModel}
		model["warning_messages"] = []string{"testString"}

		assert.Equal(t, result, model)
	}

	connectorParametersModel := new(backuprecoveryv1.ConnectorParameters)
	connectorParametersModel.ConnectionID = core.Int64Ptr(int64(26))
	connectorParametersModel.ConnectorGroupID = core.Int64Ptr(int64(26))
	connectorParametersModel.Endpoint = core.StringPtr("testString")
	connectorParametersModel.Environment = core.StringPtr("kVMware")
	connectorParametersModel.ID = core.Int64Ptr(int64(26))
	connectorParametersModel.Version = core.Int64Ptr(int64(26))

	cassandraPortsInfoModel := new(backuprecoveryv1.CassandraPortsInfo)
	cassandraPortsInfoModel.JmxPort = core.Int64Ptr(int64(38))
	cassandraPortsInfoModel.NativeTransportPort = core.Int64Ptr(int64(38))
	cassandraPortsInfoModel.RpcPort = core.Int64Ptr(int64(38))
	cassandraPortsInfoModel.SslStoragePort = core.Int64Ptr(int64(38))
	cassandraPortsInfoModel.StoragePort = core.Int64Ptr(int64(38))

	cassandraSecurityInfoModel := new(backuprecoveryv1.CassandraSecurityInfo)
	cassandraSecurityInfoModel.CassandraAuthRequired = core.BoolPtr(true)
	cassandraSecurityInfoModel.CassandraAuthType = core.StringPtr("PASSWORD")
	cassandraSecurityInfoModel.CassandraAuthorizer = core.StringPtr("testString")
	cassandraSecurityInfoModel.ClientEncryption = core.BoolPtr(true)
	cassandraSecurityInfoModel.DseAuthorization = core.BoolPtr(true)
	cassandraSecurityInfoModel.ServerEncryptionReqClientAuth = core.BoolPtr(true)
	cassandraSecurityInfoModel.ServerInternodeEncryptionType = core.StringPtr("testString")

	cassandraConnectParamsModel := new(backuprecoveryv1.CassandraConnectParams)
	cassandraConnectParamsModel.CassandraPortsInfo = cassandraPortsInfoModel
	cassandraConnectParamsModel.CassandraSecurityInfo = cassandraSecurityInfoModel
	cassandraConnectParamsModel.CassandraVersion = core.StringPtr("testString")
	cassandraConnectParamsModel.CommitLogBackupLocation = core.StringPtr("testString")
	cassandraConnectParamsModel.ConfigDirectory = core.StringPtr("testString")
	cassandraConnectParamsModel.DataCenters = []string{"testString"}
	cassandraConnectParamsModel.DseConfigDirectory = core.StringPtr("testString")
	cassandraConnectParamsModel.DseVersion = core.StringPtr("testString")
	cassandraConnectParamsModel.IsDseAuthenticator = core.BoolPtr(true)
	cassandraConnectParamsModel.IsDseTieredStorage = core.BoolPtr(true)
	cassandraConnectParamsModel.IsJmxAuthEnable = core.BoolPtr(true)
	cassandraConnectParamsModel.KerberosPrincipal = core.StringPtr("testString")
	cassandraConnectParamsModel.PrimaryHost = core.StringPtr("testString")
	cassandraConnectParamsModel.Seeds = []string{"testString"}
	cassandraConnectParamsModel.SolrNodes = []string{"testString"}
	cassandraConnectParamsModel.SolrPort = core.Int64Ptr(int64(38))

	couchbaseConnectParamsModel := new(backuprecoveryv1.CouchbaseConnectParams)
	couchbaseConnectParamsModel.CarrierDirectPort = core.Int64Ptr(int64(38))
	couchbaseConnectParamsModel.HttpDirectPort = core.Int64Ptr(int64(38))
	couchbaseConnectParamsModel.RequiresSsl = core.BoolPtr(true)
	couchbaseConnectParamsModel.Seeds = []string{"testString"}

	hadoopDiscoveryParamsModel := new(backuprecoveryv1.HadoopDiscoveryParams)
	hadoopDiscoveryParamsModel.ConfigDirectory = core.StringPtr("testString")
	hadoopDiscoveryParamsModel.Host = core.StringPtr("testString")

	hBaseConnectParamsModel := new(backuprecoveryv1.HBaseConnectParams)
	hBaseConnectParamsModel.HbaseDiscoveryParams = hadoopDiscoveryParamsModel
	hBaseConnectParamsModel.HdfsEntityID = core.Int64Ptr(int64(26))
	hBaseConnectParamsModel.KerberosPrincipal = core.StringPtr("testString")
	hBaseConnectParamsModel.RootDataDirectory = core.StringPtr("testString")
	hBaseConnectParamsModel.ZookeeperQuorum = []string{"testString"}

	hdfsConnectParamsModel := new(backuprecoveryv1.HdfsConnectParams)
	hdfsConnectParamsModel.HadoopDistribution = core.StringPtr("CDH")
	hdfsConnectParamsModel.HadoopVersion = core.StringPtr("testString")
	hdfsConnectParamsModel.HdfsConnectionType = core.StringPtr("DFS")
	hdfsConnectParamsModel.HdfsDiscoveryParams = hadoopDiscoveryParamsModel
	hdfsConnectParamsModel.KerberosPrincipal = core.StringPtr("testString")
	hdfsConnectParamsModel.Namenode = core.StringPtr("testString")
	hdfsConnectParamsModel.Port = core.Int64Ptr(int64(38))

	hiveConnectParamsModel := new(backuprecoveryv1.HiveConnectParams)
	hiveConnectParamsModel.EntityThresholdExceeded = core.BoolPtr(true)
	hiveConnectParamsModel.HdfsEntityID = core.Int64Ptr(int64(26))
	hiveConnectParamsModel.HiveDiscoveryParams = hadoopDiscoveryParamsModel
	hiveConnectParamsModel.KerberosPrincipal = core.StringPtr("testString")
	hiveConnectParamsModel.Metastore = core.StringPtr("testString")
	hiveConnectParamsModel.ThriftPort = core.Int64Ptr(int64(38))

	networkPoolConfigModel := new(backuprecoveryv1.NetworkPoolConfig)
	networkPoolConfigModel.PoolName = core.StringPtr("testString")
	networkPoolConfigModel.Subnet = core.StringPtr("testString")
	networkPoolConfigModel.UseSmartConnect = core.BoolPtr(true)

	zoneConfigModel := new(backuprecoveryv1.ZoneConfig)
	zoneConfigModel.DynamicNetworkPoolConfig = networkPoolConfigModel

	registeredProtectionSourceIsilonParamsModel := new(backuprecoveryv1.RegisteredProtectionSourceIsilonParams)
	registeredProtectionSourceIsilonParamsModel.ZoneConfigList = []backuprecoveryv1.ZoneConfig{*zoneConfigModel}

	mongoDbConnectParamsModel := new(backuprecoveryv1.MongoDBConnectParams)
	mongoDbConnectParamsModel.AuthType = core.StringPtr("SCRAM")
	mongoDbConnectParamsModel.AuthenticatingDatabaseName = core.StringPtr("testString")
	mongoDbConnectParamsModel.RequiresSsl = core.BoolPtr(true)
	mongoDbConnectParamsModel.SecondaryNodeTag = core.StringPtr("testString")
	mongoDbConnectParamsModel.Seeds = []string{"testString"}
	mongoDbConnectParamsModel.UseFixedNodeForBackup = core.BoolPtr(true)
	mongoDbConnectParamsModel.UseSecondaryForBackup = core.BoolPtr(true)

	nasServerCredentialsModel := new(backuprecoveryv1.NASServerCredentials)
	nasServerCredentialsModel.Domain = core.StringPtr("testString")
	nasServerCredentialsModel.NasProtocol = core.StringPtr("kNoProtocol")

	sitesDiscoveryParamsModel := new(backuprecoveryv1.SitesDiscoveryParams)
	sitesDiscoveryParamsModel.EnableSiteTagging = core.BoolPtr(true)

	teamsAdditionalParamsModel := new(backuprecoveryv1.TeamsAdditionalParams)
	teamsAdditionalParamsModel.AllowPostsBackup = core.BoolPtr(true)

	usersDiscoveryParamsModel := new(backuprecoveryv1.UsersDiscoveryParams)
	usersDiscoveryParamsModel.AllowChatsBackup = core.BoolPtr(true)
	usersDiscoveryParamsModel.DiscoverUsersWithMailbox = core.BoolPtr(true)
	usersDiscoveryParamsModel.DiscoverUsersWithOnedrive = core.BoolPtr(true)
	usersDiscoveryParamsModel.FetchMailboxInfo = core.BoolPtr(true)
	usersDiscoveryParamsModel.FetchOneDriveInfo = core.BoolPtr(true)
	usersDiscoveryParamsModel.SkipUsersWithoutMySite = core.BoolPtr(true)

	objectsDiscoveryParamsModel := new(backuprecoveryv1.ObjectsDiscoveryParams)
	objectsDiscoveryParamsModel.DiscoverableObjectTypeList = []string{"testString"}
	objectsDiscoveryParamsModel.SitesDiscoveryParams = sitesDiscoveryParamsModel
	objectsDiscoveryParamsModel.TeamsAdditionalParams = teamsAdditionalParamsModel
	objectsDiscoveryParamsModel.UsersDiscoveryParams = usersDiscoveryParamsModel

	m365CsmParamsModel := new(backuprecoveryv1.M365CsmParams)
	m365CsmParamsModel.BackupAllowed = core.BoolPtr(true)

	o365ConnectParamsModel := new(backuprecoveryv1.O365ConnectParams)
	o365ConnectParamsModel.ObjectsDiscoveryParams = objectsDiscoveryParamsModel
	o365ConnectParamsModel.CsmParams = m365CsmParamsModel

	office365CredentialsModel := new(backuprecoveryv1.Office365Credentials)
	office365CredentialsModel.ClientID = core.StringPtr("testString")
	office365CredentialsModel.ClientSecret = core.StringPtr("testString")
	office365CredentialsModel.GrantType = core.StringPtr("testString")
	office365CredentialsModel.Scope = core.StringPtr("testString")
	office365CredentialsModel.UseOAuthForExchangeOnline = core.BoolPtr(true)

	credentialsModel := new(backuprecoveryv1.Credentials)
	credentialsModel.Username = core.StringPtr("testString")
	credentialsModel.Password = core.StringPtr("testString")

	timeModel := new(backuprecoveryv1.Time)
	timeModel.Hour = core.Int64Ptr(int64(38))
	timeModel.Minute = core.Int64Ptr(int64(38))

	dayTimeParamsModel := new(backuprecoveryv1.DayTimeParams)
	dayTimeParamsModel.Day = core.StringPtr("kSunday")
	dayTimeParamsModel.Time = timeModel

	dayTimeWindowModel := new(backuprecoveryv1.DayTimeWindow)
	dayTimeWindowModel.EndTime = dayTimeParamsModel
	dayTimeWindowModel.StartTime = dayTimeParamsModel

	throttlingWindowModel := new(backuprecoveryv1.ThrottlingWindow)
	throttlingWindowModel.DayTimeWindow = dayTimeWindowModel
	throttlingWindowModel.Threshold = core.Int64Ptr(int64(26))

	throttlingConfigurationModel := new(backuprecoveryv1.ThrottlingConfiguration)
	throttlingConfigurationModel.FixedThreshold = core.Int64Ptr(int64(26))
	throttlingConfigurationModel.PatternType = core.StringPtr("kNoThrottling")
	throttlingConfigurationModel.ThrottlingWindows = []backuprecoveryv1.ThrottlingWindow{*throttlingWindowModel}

	sourceThrottlingConfigurationModel := new(backuprecoveryv1.SourceThrottlingConfiguration)
	sourceThrottlingConfigurationModel.CpuThrottlingConfig = throttlingConfigurationModel
	sourceThrottlingConfigurationModel.NetworkThrottlingConfig = throttlingConfigurationModel

	physicalParamsModel := new(backuprecoveryv1.PhysicalParams)
	physicalParamsModel.Applications = []string{"kVMware"}
	physicalParamsModel.Password = core.StringPtr("testString")
	physicalParamsModel.ThrottlingConfig = sourceThrottlingConfigurationModel
	physicalParamsModel.Username = core.StringPtr("testString")

	hostSettingsCheckResultModel := new(backuprecoveryv1.HostSettingsCheckResult)
	hostSettingsCheckResultModel.CheckType = core.StringPtr("kIsAgentPortAccessible")
	hostSettingsCheckResultModel.ResultType = core.StringPtr("kPass")
	hostSettingsCheckResultModel.UserMessage = core.StringPtr("testString")

	registeredAppInfoModel := new(backuprecoveryv1.RegisteredAppInfo)
	registeredAppInfoModel.AuthenticationErrorMessage = core.StringPtr("testString")
	registeredAppInfoModel.AuthenticationStatus = core.StringPtr("kPending")
	registeredAppInfoModel.Environment = core.StringPtr("kPhysical")
	registeredAppInfoModel.HostSettingsCheckResults = []backuprecoveryv1.HostSettingsCheckResult{*hostSettingsCheckResultModel}
	registeredAppInfoModel.RefreshErrorMessage = core.StringPtr("testString")

	sfdcParamsModel := new(backuprecoveryv1.SfdcParams)
	sfdcParamsModel.AccessToken = core.StringPtr("testString")
	sfdcParamsModel.ConcurrentApiRequestsLimit = core.Int64Ptr(int64(26))
	sfdcParamsModel.ConsumerKey = core.StringPtr("testString")
	sfdcParamsModel.ConsumerSecret = core.StringPtr("testString")
	sfdcParamsModel.DailyApiLimit = core.Int64Ptr(int64(26))
	sfdcParamsModel.Endpoint = core.StringPtr("testString")
	sfdcParamsModel.EndpointType = core.StringPtr("PROD")
	sfdcParamsModel.MetadataEndpointURL = core.StringPtr("testString")
	sfdcParamsModel.RefreshToken = core.StringPtr("testString")
	sfdcParamsModel.SoapEndpointURL = core.StringPtr("testString")
	sfdcParamsModel.UseBulkApi = core.BoolPtr(true)

	subnetModel := new(backuprecoveryv1.Subnet)
	subnetModel.Component = core.StringPtr("testString")
	subnetModel.Description = core.StringPtr("testString")
	subnetModel.ID = core.Float64Ptr(float64(72.5))
	subnetModel.Ip = core.StringPtr("testString")
	subnetModel.NetmaskBits = core.Float64Ptr(float64(72.5))
	subnetModel.NetmaskIp4 = core.StringPtr("testString")
	subnetModel.NfsAccess = core.StringPtr("kDisabled")
	subnetModel.NfsAllSquash = core.BoolPtr(true)
	subnetModel.NfsRootSquash = core.BoolPtr(true)
	subnetModel.S3Access = core.StringPtr("kDisabled")
	subnetModel.SmbAccess = core.StringPtr("kDisabled")
	subnetModel.TenantID = core.StringPtr("testString")

	latencyThresholdsModel := new(backuprecoveryv1.LatencyThresholds)
	latencyThresholdsModel.ActiveTaskMsecs = core.Int64Ptr(int64(26))
	latencyThresholdsModel.NewTaskMsecs = core.Int64Ptr(int64(26))

	nasSourceThrottlingParamsModel := new(backuprecoveryv1.NasSourceThrottlingParams)
	nasSourceThrottlingParamsModel.MaxParallelMetadataFetchFullPercentage = core.Int64Ptr(int64(38))
	nasSourceThrottlingParamsModel.MaxParallelMetadataFetchIncrementalPercentage = core.Int64Ptr(int64(38))
	nasSourceThrottlingParamsModel.MaxParallelReadWriteFullPercentage = core.Int64Ptr(int64(38))
	nasSourceThrottlingParamsModel.MaxParallelReadWriteIncrementalPercentage = core.Int64Ptr(int64(38))

	storageArraySnapshotMaxSpaceConfigParamsModel := new(backuprecoveryv1.StorageArraySnapshotMaxSpaceConfigParams)
	storageArraySnapshotMaxSpaceConfigParamsModel.MaxSnapshotSpacePercentage = core.Int64Ptr(int64(38))

	storageArraySnapshotMaxSnapshotConfigParamsModel := new(backuprecoveryv1.StorageArraySnapshotMaxSnapshotConfigParams)
	storageArraySnapshotMaxSnapshotConfigParamsModel.MaxSnapshots = core.Int64Ptr(int64(38))

	storageArraySnapshotThrottlingPolicyModel := new(backuprecoveryv1.StorageArraySnapshotThrottlingPolicy)
	storageArraySnapshotThrottlingPolicyModel.ID = core.Int64Ptr(int64(26))
	storageArraySnapshotThrottlingPolicyModel.IsMaxSnapshotsConfigEnabled = core.BoolPtr(true)
	storageArraySnapshotThrottlingPolicyModel.IsMaxSpaceConfigEnabled = core.BoolPtr(true)
	storageArraySnapshotThrottlingPolicyModel.MaxSnapshotConfig = storageArraySnapshotMaxSnapshotConfigParamsModel
	storageArraySnapshotThrottlingPolicyModel.MaxSpaceConfig = storageArraySnapshotMaxSpaceConfigParamsModel

	storageArraySnapshotConfigParamsModel := new(backuprecoveryv1.StorageArraySnapshotConfigParams)
	storageArraySnapshotConfigParamsModel.IsMaxSnapshotsConfigEnabled = core.BoolPtr(true)
	storageArraySnapshotConfigParamsModel.IsMaxSpaceConfigEnabled = core.BoolPtr(true)
	storageArraySnapshotConfigParamsModel.StorageArraySnapshotMaxSpaceConfig = storageArraySnapshotMaxSpaceConfigParamsModel
	storageArraySnapshotConfigParamsModel.StorageArraySnapshotThrottlingPolicies = []backuprecoveryv1.StorageArraySnapshotThrottlingPolicy{*storageArraySnapshotThrottlingPolicyModel}

	throttlingPolicyParametersModel := new(backuprecoveryv1.ThrottlingPolicyParameters)
	throttlingPolicyParametersModel.EnforceMaxStreams = core.BoolPtr(true)
	throttlingPolicyParametersModel.EnforceRegisteredSourceMaxBackups = core.BoolPtr(true)
	throttlingPolicyParametersModel.IsEnabled = core.BoolPtr(true)
	throttlingPolicyParametersModel.LatencyThresholds = latencyThresholdsModel
	throttlingPolicyParametersModel.MaxConcurrentStreams = core.Int64Ptr(int64(38))
	throttlingPolicyParametersModel.NasSourceParams = nasSourceThrottlingParamsModel
	throttlingPolicyParametersModel.RegisteredSourceMaxConcurrentBackups = core.Int64Ptr(int64(38))
	throttlingPolicyParametersModel.StorageArraySnapshotConfig = storageArraySnapshotConfigParamsModel

	throttlingPolicyOverrideModel := new(backuprecoveryv1.ThrottlingPolicyOverride)
	throttlingPolicyOverrideModel.DatastoreID = core.Int64Ptr(int64(26))
	throttlingPolicyOverrideModel.DatastoreName = core.StringPtr("testString")
	throttlingPolicyOverrideModel.ThrottlingPolicy = throttlingPolicyParametersModel

	udaSourceCapabilitiesModel := new(backuprecoveryv1.UdaSourceCapabilities)
	udaSourceCapabilitiesModel.AutoLogBackup = core.BoolPtr(true)
	udaSourceCapabilitiesModel.DynamicConfig = core.BoolPtr(true)
	udaSourceCapabilitiesModel.EntitySupport = core.BoolPtr(true)
	udaSourceCapabilitiesModel.EtLogBackup = core.BoolPtr(true)
	udaSourceCapabilitiesModel.ExternalDisks = core.BoolPtr(true)
	udaSourceCapabilitiesModel.FullBackup = core.BoolPtr(true)
	udaSourceCapabilitiesModel.IncrBackup = core.BoolPtr(true)
	udaSourceCapabilitiesModel.LogBackup = core.BoolPtr(true)
	udaSourceCapabilitiesModel.MultiObjectRestore = core.BoolPtr(true)
	udaSourceCapabilitiesModel.PauseResumeBackup = core.BoolPtr(true)
	udaSourceCapabilitiesModel.PostBackupJobScript = core.BoolPtr(true)
	udaSourceCapabilitiesModel.PostRestoreJobScript = core.BoolPtr(true)
	udaSourceCapabilitiesModel.PreBackupJobScript = core.BoolPtr(true)
	udaSourceCapabilitiesModel.PreRestoreJobScript = core.BoolPtr(true)
	udaSourceCapabilitiesModel.ResourceThrottling = core.BoolPtr(true)
	udaSourceCapabilitiesModel.SnapfsCert = core.BoolPtr(true)

	keyValueStrPairModel := new(backuprecoveryv1.KeyValueStrPair)
	keyValueStrPairModel.Key = core.StringPtr("testString")
	keyValueStrPairModel.Value = core.StringPtr("testString")

	udaConnectParamsModel := new(backuprecoveryv1.UdaConnectParams)
	udaConnectParamsModel.Capabilities = udaSourceCapabilitiesModel
	udaConnectParamsModel.Credentials = credentialsModel
	udaConnectParamsModel.EtEnableLogBackupPolicy = core.BoolPtr(true)
	udaConnectParamsModel.EtEnableRunNow = core.BoolPtr(true)
	udaConnectParamsModel.HostType = core.StringPtr("kLinux")
	udaConnectParamsModel.Hosts = []string{"testString"}
	udaConnectParamsModel.LiveDataView = core.BoolPtr(true)
	udaConnectParamsModel.LiveLogView = core.BoolPtr(true)
	udaConnectParamsModel.MountDir = core.StringPtr("testString")
	udaConnectParamsModel.MountView = core.BoolPtr(true)
	udaConnectParamsModel.ScriptDir = core.StringPtr("testString")
	udaConnectParamsModel.SourceArgs = core.StringPtr("testString")
	udaConnectParamsModel.SourceRegistrationArguments = []backuprecoveryv1.KeyValueStrPair{*keyValueStrPairModel}
	udaConnectParamsModel.SourceType = core.StringPtr("testString")

	vlanParametersModel := new(backuprecoveryv1.VlanParameters)
	vlanParametersModel.DisableVlan = core.BoolPtr(true)
	vlanParametersModel.InterfaceName = core.StringPtr("testString")
	vlanParametersModel.Vlan = core.Int64Ptr(int64(38))

	model := new(backuprecoveryv1.RegisteredSourceInfo)
	model.AccessInfo = connectorParametersModel
	model.AllowedIpAddresses = []string{"testString"}
	model.AuthenticationErrorMessage = core.StringPtr("testString")
	model.AuthenticationStatus = core.StringPtr("kPending")
	model.BlacklistedIpAddresses = []string{"testString"}
	model.CassandraParams = cassandraConnectParamsModel
	model.CouchbaseParams = couchbaseConnectParamsModel
	model.DeniedIpAddresses = []string{"testString"}
	model.Environments = []string{"kVMware"}
	model.HbaseParams = hBaseConnectParamsModel
	model.HdfsParams = hdfsConnectParamsModel
	model.HiveParams = hiveConnectParamsModel
	model.IsDbAuthenticated = core.BoolPtr(true)
	model.IsStorageArraySnapshotEnabled = core.BoolPtr(true)
	model.IsilonParams = registeredProtectionSourceIsilonParamsModel
	model.LinkVmsAcrossVcenter = core.BoolPtr(true)
	model.MinimumFreeSpaceGB = core.Int64Ptr(int64(26))
	model.MinimumFreeSpacePercent = core.Int64Ptr(int64(26))
	model.MongodbParams = mongoDbConnectParamsModel
	model.NasMountCredentials = nasServerCredentialsModel
	model.O365Params = o365ConnectParamsModel
	model.Office365CredentialsList = []backuprecoveryv1.Office365Credentials{*office365CredentialsModel}
	model.Office365Region = core.StringPtr("testString")
	model.Office365ServiceAccountCredentialsList = []backuprecoveryv1.Credentials{*credentialsModel}
	model.Password = core.StringPtr("testString")
	model.PhysicalParams = physicalParamsModel
	model.ProgressMonitorPath = core.StringPtr("testString")
	model.RefreshErrorMessage = core.StringPtr("testString")
	model.RefreshTimeUsecs = core.Int64Ptr(int64(26))
	model.RegisteredAppsInfo = []backuprecoveryv1.RegisteredAppInfo{*registeredAppInfoModel}
	model.RegistrationTimeUsecs = core.Int64Ptr(int64(26))
	model.SfdcParams = sfdcParamsModel
	model.Subnets = []backuprecoveryv1.Subnet{*subnetModel}
	model.ThrottlingPolicy = throttlingPolicyParametersModel
	model.ThrottlingPolicyOverrides = []backuprecoveryv1.ThrottlingPolicyOverride{*throttlingPolicyOverrideModel}
	model.UdaParams = udaConnectParamsModel
	model.UpdateLastBackupDetails = core.BoolPtr(true)
	model.UseOAuthForExchangeOnline = core.BoolPtr(true)
	model.UseVmBiosUUID = core.BoolPtr(true)
	model.UserMessages = []string{"testString"}
	model.Username = core.StringPtr("testString")
	model.VlanParams = vlanParametersModel
	model.WarningMessages = []string{"testString"}

	result, err := backuprecovery.DataSourceIbmBackupRecoveryRegistrationInfoRegisteredSourceInfoToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryRegistrationInfoConnectorParametersToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["connection_id"] = int(26)
		model["connector_group_id"] = int(26)
		model["endpoint"] = "testString"
		model["environment"] = "kVMware"
		model["id"] = int(26)
		model["version"] = int(26)

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.ConnectorParameters)
	model.ConnectionID = core.Int64Ptr(int64(26))
	model.ConnectorGroupID = core.Int64Ptr(int64(26))
	model.Endpoint = core.StringPtr("testString")
	model.Environment = core.StringPtr("kVMware")
	model.ID = core.Int64Ptr(int64(26))
	model.Version = core.Int64Ptr(int64(26))

	result, err := backuprecovery.DataSourceIbmBackupRecoveryRegistrationInfoConnectorParametersToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryRegistrationInfoCassandraConnectParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		cassandraPortsInfoModel := make(map[string]interface{})
		cassandraPortsInfoModel["jmx_port"] = int(38)
		cassandraPortsInfoModel["native_transport_port"] = int(38)
		cassandraPortsInfoModel["rpc_port"] = int(38)
		cassandraPortsInfoModel["ssl_storage_port"] = int(38)
		cassandraPortsInfoModel["storage_port"] = int(38)

		cassandraSecurityInfoModel := make(map[string]interface{})
		cassandraSecurityInfoModel["cassandra_auth_required"] = true
		cassandraSecurityInfoModel["cassandra_auth_type"] = "PASSWORD"
		cassandraSecurityInfoModel["cassandra_authorizer"] = "testString"
		cassandraSecurityInfoModel["client_encryption"] = true
		cassandraSecurityInfoModel["dse_authorization"] = true
		cassandraSecurityInfoModel["server_encryption_req_client_auth"] = true
		cassandraSecurityInfoModel["server_internode_encryption_type"] = "testString"

		model := make(map[string]interface{})
		model["cassandra_ports_info"] = []map[string]interface{}{cassandraPortsInfoModel}
		model["cassandra_security_info"] = []map[string]interface{}{cassandraSecurityInfoModel}
		model["cassandra_version"] = "testString"
		model["commit_log_backup_location"] = "testString"
		model["config_directory"] = "testString"
		model["data_centers"] = []string{"testString"}
		model["dse_config_directory"] = "testString"
		model["dse_version"] = "testString"
		model["is_dse_authenticator"] = true
		model["is_dse_tiered_storage"] = true
		model["is_jmx_auth_enable"] = true
		model["kerberos_principal"] = "testString"
		model["primary_host"] = "testString"
		model["seeds"] = []string{"testString"}
		model["solr_nodes"] = []string{"testString"}
		model["solr_port"] = int(38)

		assert.Equal(t, result, model)
	}

	cassandraPortsInfoModel := new(backuprecoveryv1.CassandraPortsInfo)
	cassandraPortsInfoModel.JmxPort = core.Int64Ptr(int64(38))
	cassandraPortsInfoModel.NativeTransportPort = core.Int64Ptr(int64(38))
	cassandraPortsInfoModel.RpcPort = core.Int64Ptr(int64(38))
	cassandraPortsInfoModel.SslStoragePort = core.Int64Ptr(int64(38))
	cassandraPortsInfoModel.StoragePort = core.Int64Ptr(int64(38))

	cassandraSecurityInfoModel := new(backuprecoveryv1.CassandraSecurityInfo)
	cassandraSecurityInfoModel.CassandraAuthRequired = core.BoolPtr(true)
	cassandraSecurityInfoModel.CassandraAuthType = core.StringPtr("PASSWORD")
	cassandraSecurityInfoModel.CassandraAuthorizer = core.StringPtr("testString")
	cassandraSecurityInfoModel.ClientEncryption = core.BoolPtr(true)
	cassandraSecurityInfoModel.DseAuthorization = core.BoolPtr(true)
	cassandraSecurityInfoModel.ServerEncryptionReqClientAuth = core.BoolPtr(true)
	cassandraSecurityInfoModel.ServerInternodeEncryptionType = core.StringPtr("testString")

	model := new(backuprecoveryv1.CassandraConnectParams)
	model.CassandraPortsInfo = cassandraPortsInfoModel
	model.CassandraSecurityInfo = cassandraSecurityInfoModel
	model.CassandraVersion = core.StringPtr("testString")
	model.CommitLogBackupLocation = core.StringPtr("testString")
	model.ConfigDirectory = core.StringPtr("testString")
	model.DataCenters = []string{"testString"}
	model.DseConfigDirectory = core.StringPtr("testString")
	model.DseVersion = core.StringPtr("testString")
	model.IsDseAuthenticator = core.BoolPtr(true)
	model.IsDseTieredStorage = core.BoolPtr(true)
	model.IsJmxAuthEnable = core.BoolPtr(true)
	model.KerberosPrincipal = core.StringPtr("testString")
	model.PrimaryHost = core.StringPtr("testString")
	model.Seeds = []string{"testString"}
	model.SolrNodes = []string{"testString"}
	model.SolrPort = core.Int64Ptr(int64(38))

	result, err := backuprecovery.DataSourceIbmBackupRecoveryRegistrationInfoCassandraConnectParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryRegistrationInfoCassandraPortsInfoToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["jmx_port"] = int(38)
		model["native_transport_port"] = int(38)
		model["rpc_port"] = int(38)
		model["ssl_storage_port"] = int(38)
		model["storage_port"] = int(38)

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.CassandraPortsInfo)
	model.JmxPort = core.Int64Ptr(int64(38))
	model.NativeTransportPort = core.Int64Ptr(int64(38))
	model.RpcPort = core.Int64Ptr(int64(38))
	model.SslStoragePort = core.Int64Ptr(int64(38))
	model.StoragePort = core.Int64Ptr(int64(38))

	result, err := backuprecovery.DataSourceIbmBackupRecoveryRegistrationInfoCassandraPortsInfoToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryRegistrationInfoCassandraSecurityInfoToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["cassandra_auth_required"] = true
		model["cassandra_auth_type"] = "PASSWORD"
		model["cassandra_authorizer"] = "testString"
		model["client_encryption"] = true
		model["dse_authorization"] = true
		model["server_encryption_req_client_auth"] = true
		model["server_internode_encryption_type"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.CassandraSecurityInfo)
	model.CassandraAuthRequired = core.BoolPtr(true)
	model.CassandraAuthType = core.StringPtr("PASSWORD")
	model.CassandraAuthorizer = core.StringPtr("testString")
	model.ClientEncryption = core.BoolPtr(true)
	model.DseAuthorization = core.BoolPtr(true)
	model.ServerEncryptionReqClientAuth = core.BoolPtr(true)
	model.ServerInternodeEncryptionType = core.StringPtr("testString")

	result, err := backuprecovery.DataSourceIbmBackupRecoveryRegistrationInfoCassandraSecurityInfoToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryRegistrationInfoCouchbaseConnectParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["carrier_direct_port"] = int(38)
		model["http_direct_port"] = int(38)
		model["requires_ssl"] = true
		model["seeds"] = []string{"testString"}

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.CouchbaseConnectParams)
	model.CarrierDirectPort = core.Int64Ptr(int64(38))
	model.HttpDirectPort = core.Int64Ptr(int64(38))
	model.RequiresSsl = core.BoolPtr(true)
	model.Seeds = []string{"testString"}

	result, err := backuprecovery.DataSourceIbmBackupRecoveryRegistrationInfoCouchbaseConnectParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryRegistrationInfoHBaseConnectParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		hadoopDiscoveryParamsModel := make(map[string]interface{})
		hadoopDiscoveryParamsModel["config_directory"] = "testString"
		hadoopDiscoveryParamsModel["host"] = "testString"

		model := make(map[string]interface{})
		model["hbase_discovery_params"] = []map[string]interface{}{hadoopDiscoveryParamsModel}
		model["hdfs_entity_id"] = int(26)
		model["kerberos_principal"] = "testString"
		model["root_data_directory"] = "testString"
		model["zookeeper_quorum"] = []string{"testString"}

		assert.Equal(t, result, model)
	}

	hadoopDiscoveryParamsModel := new(backuprecoveryv1.HadoopDiscoveryParams)
	hadoopDiscoveryParamsModel.ConfigDirectory = core.StringPtr("testString")
	hadoopDiscoveryParamsModel.Host = core.StringPtr("testString")

	model := new(backuprecoveryv1.HBaseConnectParams)
	model.HbaseDiscoveryParams = hadoopDiscoveryParamsModel
	model.HdfsEntityID = core.Int64Ptr(int64(26))
	model.KerberosPrincipal = core.StringPtr("testString")
	model.RootDataDirectory = core.StringPtr("testString")
	model.ZookeeperQuorum = []string{"testString"}

	result, err := backuprecovery.DataSourceIbmBackupRecoveryRegistrationInfoHBaseConnectParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryRegistrationInfoHadoopDiscoveryParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["config_directory"] = "testString"
		model["host"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.HadoopDiscoveryParams)
	model.ConfigDirectory = core.StringPtr("testString")
	model.Host = core.StringPtr("testString")

	result, err := backuprecovery.DataSourceIbmBackupRecoveryRegistrationInfoHadoopDiscoveryParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryRegistrationInfoHdfsConnectParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		hadoopDiscoveryParamsModel := make(map[string]interface{})
		hadoopDiscoveryParamsModel["config_directory"] = "testString"
		hadoopDiscoveryParamsModel["host"] = "testString"

		model := make(map[string]interface{})
		model["hadoop_distribution"] = "CDH"
		model["hadoop_version"] = "testString"
		model["hdfs_connection_type"] = "DFS"
		model["hdfs_discovery_params"] = []map[string]interface{}{hadoopDiscoveryParamsModel}
		model["kerberos_principal"] = "testString"
		model["namenode"] = "testString"
		model["port"] = int(38)

		assert.Equal(t, result, model)
	}

	hadoopDiscoveryParamsModel := new(backuprecoveryv1.HadoopDiscoveryParams)
	hadoopDiscoveryParamsModel.ConfigDirectory = core.StringPtr("testString")
	hadoopDiscoveryParamsModel.Host = core.StringPtr("testString")

	model := new(backuprecoveryv1.HdfsConnectParams)
	model.HadoopDistribution = core.StringPtr("CDH")
	model.HadoopVersion = core.StringPtr("testString")
	model.HdfsConnectionType = core.StringPtr("DFS")
	model.HdfsDiscoveryParams = hadoopDiscoveryParamsModel
	model.KerberosPrincipal = core.StringPtr("testString")
	model.Namenode = core.StringPtr("testString")
	model.Port = core.Int64Ptr(int64(38))

	result, err := backuprecovery.DataSourceIbmBackupRecoveryRegistrationInfoHdfsConnectParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryRegistrationInfoHiveConnectParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		hadoopDiscoveryParamsModel := make(map[string]interface{})
		hadoopDiscoveryParamsModel["config_directory"] = "testString"
		hadoopDiscoveryParamsModel["host"] = "testString"

		model := make(map[string]interface{})
		model["entity_threshold_exceeded"] = true
		model["hdfs_entity_id"] = int(26)
		model["hive_discovery_params"] = []map[string]interface{}{hadoopDiscoveryParamsModel}
		model["kerberos_principal"] = "testString"
		model["metastore"] = "testString"
		model["thrift_port"] = int(38)

		assert.Equal(t, result, model)
	}

	hadoopDiscoveryParamsModel := new(backuprecoveryv1.HadoopDiscoveryParams)
	hadoopDiscoveryParamsModel.ConfigDirectory = core.StringPtr("testString")
	hadoopDiscoveryParamsModel.Host = core.StringPtr("testString")

	model := new(backuprecoveryv1.HiveConnectParams)
	model.EntityThresholdExceeded = core.BoolPtr(true)
	model.HdfsEntityID = core.Int64Ptr(int64(26))
	model.HiveDiscoveryParams = hadoopDiscoveryParamsModel
	model.KerberosPrincipal = core.StringPtr("testString")
	model.Metastore = core.StringPtr("testString")
	model.ThriftPort = core.Int64Ptr(int64(38))

	result, err := backuprecovery.DataSourceIbmBackupRecoveryRegistrationInfoHiveConnectParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryRegistrationInfoRegisteredProtectionSourceIsilonParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		networkPoolConfigModel := make(map[string]interface{})
		networkPoolConfigModel["pool_name"] = "testString"
		networkPoolConfigModel["subnet"] = "testString"
		networkPoolConfigModel["use_smart_connect"] = true

		zoneConfigModel := make(map[string]interface{})
		zoneConfigModel["dynamic_network_pool_config"] = []map[string]interface{}{networkPoolConfigModel}

		model := make(map[string]interface{})
		model["zone_config_list"] = []map[string]interface{}{zoneConfigModel}

		assert.Equal(t, result, model)
	}

	networkPoolConfigModel := new(backuprecoveryv1.NetworkPoolConfig)
	networkPoolConfigModel.PoolName = core.StringPtr("testString")
	networkPoolConfigModel.Subnet = core.StringPtr("testString")
	networkPoolConfigModel.UseSmartConnect = core.BoolPtr(true)

	zoneConfigModel := new(backuprecoveryv1.ZoneConfig)
	zoneConfigModel.DynamicNetworkPoolConfig = networkPoolConfigModel

	model := new(backuprecoveryv1.RegisteredProtectionSourceIsilonParams)
	model.ZoneConfigList = []backuprecoveryv1.ZoneConfig{*zoneConfigModel}

	result, err := backuprecovery.DataSourceIbmBackupRecoveryRegistrationInfoRegisteredProtectionSourceIsilonParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryRegistrationInfoZoneConfigToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		networkPoolConfigModel := make(map[string]interface{})
		networkPoolConfigModel["pool_name"] = "testString"
		networkPoolConfigModel["subnet"] = "testString"
		networkPoolConfigModel["use_smart_connect"] = true

		model := make(map[string]interface{})
		model["dynamic_network_pool_config"] = []map[string]interface{}{networkPoolConfigModel}

		assert.Equal(t, result, model)
	}

	networkPoolConfigModel := new(backuprecoveryv1.NetworkPoolConfig)
	networkPoolConfigModel.PoolName = core.StringPtr("testString")
	networkPoolConfigModel.Subnet = core.StringPtr("testString")
	networkPoolConfigModel.UseSmartConnect = core.BoolPtr(true)

	model := new(backuprecoveryv1.ZoneConfig)
	model.DynamicNetworkPoolConfig = networkPoolConfigModel

	result, err := backuprecovery.DataSourceIbmBackupRecoveryRegistrationInfoZoneConfigToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryRegistrationInfoNetworkPoolConfigToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["pool_name"] = "testString"
		model["subnet"] = "testString"
		model["use_smart_connect"] = true

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.NetworkPoolConfig)
	model.PoolName = core.StringPtr("testString")
	model.Subnet = core.StringPtr("testString")
	model.UseSmartConnect = core.BoolPtr(true)

	result, err := backuprecovery.DataSourceIbmBackupRecoveryRegistrationInfoNetworkPoolConfigToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryRegistrationInfoMongoDBConnectParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["auth_type"] = "SCRAM"
		model["authenticating_database_name"] = "testString"
		model["requires_ssl"] = true
		model["secondary_node_tag"] = "testString"
		model["seeds"] = []string{"testString"}
		model["use_fixed_node_for_backup"] = true
		model["use_secondary_for_backup"] = true

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.MongoDBConnectParams)
	model.AuthType = core.StringPtr("SCRAM")
	model.AuthenticatingDatabaseName = core.StringPtr("testString")
	model.RequiresSsl = core.BoolPtr(true)
	model.SecondaryNodeTag = core.StringPtr("testString")
	model.Seeds = []string{"testString"}
	model.UseFixedNodeForBackup = core.BoolPtr(true)
	model.UseSecondaryForBackup = core.BoolPtr(true)

	result, err := backuprecovery.DataSourceIbmBackupRecoveryRegistrationInfoMongoDBConnectParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryRegistrationInfoNASServerCredentialsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["domain"] = "testString"
		model["nas_protocol"] = "kNoProtocol"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.NASServerCredentials)
	model.Domain = core.StringPtr("testString")
	model.NasProtocol = core.StringPtr("kNoProtocol")

	result, err := backuprecovery.DataSourceIbmBackupRecoveryRegistrationInfoNASServerCredentialsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryRegistrationInfoO365ConnectParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		sitesDiscoveryParamsModel := make(map[string]interface{})
		sitesDiscoveryParamsModel["enable_site_tagging"] = true

		teamsAdditionalParamsModel := make(map[string]interface{})
		teamsAdditionalParamsModel["allow_posts_backup"] = true

		usersDiscoveryParamsModel := make(map[string]interface{})
		usersDiscoveryParamsModel["allow_chats_backup"] = true
		usersDiscoveryParamsModel["discover_users_with_mailbox"] = true
		usersDiscoveryParamsModel["discover_users_with_onedrive"] = true
		usersDiscoveryParamsModel["fetch_mailbox_info"] = true
		usersDiscoveryParamsModel["fetch_one_drive_info"] = true
		usersDiscoveryParamsModel["skip_users_without_my_site"] = true

		objectsDiscoveryParamsModel := make(map[string]interface{})
		objectsDiscoveryParamsModel["discoverable_object_type_list"] = []string{"testString"}
		objectsDiscoveryParamsModel["sites_discovery_params"] = []map[string]interface{}{sitesDiscoveryParamsModel}
		objectsDiscoveryParamsModel["teams_additional_params"] = []map[string]interface{}{teamsAdditionalParamsModel}
		objectsDiscoveryParamsModel["users_discovery_params"] = []map[string]interface{}{usersDiscoveryParamsModel}

		m365CsmParamsModel := make(map[string]interface{})
		m365CsmParamsModel["backup_allowed"] = true

		model := make(map[string]interface{})
		model["objects_discovery_params"] = []map[string]interface{}{objectsDiscoveryParamsModel}
		model["csm_params"] = []map[string]interface{}{m365CsmParamsModel}

		assert.Equal(t, result, model)
	}

	sitesDiscoveryParamsModel := new(backuprecoveryv1.SitesDiscoveryParams)
	sitesDiscoveryParamsModel.EnableSiteTagging = core.BoolPtr(true)

	teamsAdditionalParamsModel := new(backuprecoveryv1.TeamsAdditionalParams)
	teamsAdditionalParamsModel.AllowPostsBackup = core.BoolPtr(true)

	usersDiscoveryParamsModel := new(backuprecoveryv1.UsersDiscoveryParams)
	usersDiscoveryParamsModel.AllowChatsBackup = core.BoolPtr(true)
	usersDiscoveryParamsModel.DiscoverUsersWithMailbox = core.BoolPtr(true)
	usersDiscoveryParamsModel.DiscoverUsersWithOnedrive = core.BoolPtr(true)
	usersDiscoveryParamsModel.FetchMailboxInfo = core.BoolPtr(true)
	usersDiscoveryParamsModel.FetchOneDriveInfo = core.BoolPtr(true)
	usersDiscoveryParamsModel.SkipUsersWithoutMySite = core.BoolPtr(true)

	objectsDiscoveryParamsModel := new(backuprecoveryv1.ObjectsDiscoveryParams)
	objectsDiscoveryParamsModel.DiscoverableObjectTypeList = []string{"testString"}
	objectsDiscoveryParamsModel.SitesDiscoveryParams = sitesDiscoveryParamsModel
	objectsDiscoveryParamsModel.TeamsAdditionalParams = teamsAdditionalParamsModel
	objectsDiscoveryParamsModel.UsersDiscoveryParams = usersDiscoveryParamsModel

	m365CsmParamsModel := new(backuprecoveryv1.M365CsmParams)
	m365CsmParamsModel.BackupAllowed = core.BoolPtr(true)

	model := new(backuprecoveryv1.O365ConnectParams)
	model.ObjectsDiscoveryParams = objectsDiscoveryParamsModel
	model.CsmParams = m365CsmParamsModel

	result, err := backuprecovery.DataSourceIbmBackupRecoveryRegistrationInfoO365ConnectParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryRegistrationInfoObjectsDiscoveryParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		sitesDiscoveryParamsModel := make(map[string]interface{})
		sitesDiscoveryParamsModel["enable_site_tagging"] = true

		teamsAdditionalParamsModel := make(map[string]interface{})
		teamsAdditionalParamsModel["allow_posts_backup"] = true

		usersDiscoveryParamsModel := make(map[string]interface{})
		usersDiscoveryParamsModel["allow_chats_backup"] = true
		usersDiscoveryParamsModel["discover_users_with_mailbox"] = true
		usersDiscoveryParamsModel["discover_users_with_onedrive"] = true
		usersDiscoveryParamsModel["fetch_mailbox_info"] = true
		usersDiscoveryParamsModel["fetch_one_drive_info"] = true
		usersDiscoveryParamsModel["skip_users_without_my_site"] = true

		model := make(map[string]interface{})
		model["discoverable_object_type_list"] = []string{"testString"}
		model["sites_discovery_params"] = []map[string]interface{}{sitesDiscoveryParamsModel}
		model["teams_additional_params"] = []map[string]interface{}{teamsAdditionalParamsModel}
		model["users_discovery_params"] = []map[string]interface{}{usersDiscoveryParamsModel}

		assert.Equal(t, result, model)
	}

	sitesDiscoveryParamsModel := new(backuprecoveryv1.SitesDiscoveryParams)
	sitesDiscoveryParamsModel.EnableSiteTagging = core.BoolPtr(true)

	teamsAdditionalParamsModel := new(backuprecoveryv1.TeamsAdditionalParams)
	teamsAdditionalParamsModel.AllowPostsBackup = core.BoolPtr(true)

	usersDiscoveryParamsModel := new(backuprecoveryv1.UsersDiscoveryParams)
	usersDiscoveryParamsModel.AllowChatsBackup = core.BoolPtr(true)
	usersDiscoveryParamsModel.DiscoverUsersWithMailbox = core.BoolPtr(true)
	usersDiscoveryParamsModel.DiscoverUsersWithOnedrive = core.BoolPtr(true)
	usersDiscoveryParamsModel.FetchMailboxInfo = core.BoolPtr(true)
	usersDiscoveryParamsModel.FetchOneDriveInfo = core.BoolPtr(true)
	usersDiscoveryParamsModel.SkipUsersWithoutMySite = core.BoolPtr(true)

	model := new(backuprecoveryv1.ObjectsDiscoveryParams)
	model.DiscoverableObjectTypeList = []string{"testString"}
	model.SitesDiscoveryParams = sitesDiscoveryParamsModel
	model.TeamsAdditionalParams = teamsAdditionalParamsModel
	model.UsersDiscoveryParams = usersDiscoveryParamsModel

	result, err := backuprecovery.DataSourceIbmBackupRecoveryRegistrationInfoObjectsDiscoveryParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryRegistrationInfoSitesDiscoveryParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["enable_site_tagging"] = true

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.SitesDiscoveryParams)
	model.EnableSiteTagging = core.BoolPtr(true)

	result, err := backuprecovery.DataSourceIbmBackupRecoveryRegistrationInfoSitesDiscoveryParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryRegistrationInfoTeamsAdditionalParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["allow_posts_backup"] = true

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.TeamsAdditionalParams)
	model.AllowPostsBackup = core.BoolPtr(true)

	result, err := backuprecovery.DataSourceIbmBackupRecoveryRegistrationInfoTeamsAdditionalParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryRegistrationInfoUsersDiscoveryParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["allow_chats_backup"] = true
		model["discover_users_with_mailbox"] = true
		model["discover_users_with_onedrive"] = true
		model["fetch_mailbox_info"] = true
		model["fetch_one_drive_info"] = true
		model["skip_users_without_my_site"] = true

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.UsersDiscoveryParams)
	model.AllowChatsBackup = core.BoolPtr(true)
	model.DiscoverUsersWithMailbox = core.BoolPtr(true)
	model.DiscoverUsersWithOnedrive = core.BoolPtr(true)
	model.FetchMailboxInfo = core.BoolPtr(true)
	model.FetchOneDriveInfo = core.BoolPtr(true)
	model.SkipUsersWithoutMySite = core.BoolPtr(true)

	result, err := backuprecovery.DataSourceIbmBackupRecoveryRegistrationInfoUsersDiscoveryParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryRegistrationInfoM365CsmParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["backup_allowed"] = true

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.M365CsmParams)
	model.BackupAllowed = core.BoolPtr(true)

	result, err := backuprecovery.DataSourceIbmBackupRecoveryRegistrationInfoM365CsmParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryRegistrationInfoOffice365CredentialsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["client_id"] = "testString"
		model["client_secret"] = "testString"
		model["grant_type"] = "testString"
		model["scope"] = "testString"
		model["use_o_auth_for_exchange_online"] = true

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.Office365Credentials)
	model.ClientID = core.StringPtr("testString")
	model.ClientSecret = core.StringPtr("testString")
	model.GrantType = core.StringPtr("testString")
	model.Scope = core.StringPtr("testString")
	model.UseOAuthForExchangeOnline = core.BoolPtr(true)

	result, err := backuprecovery.DataSourceIbmBackupRecoveryRegistrationInfoOffice365CredentialsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryRegistrationInfoCredentialsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["username"] = "testString"
		model["password"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.Credentials)
	model.Username = core.StringPtr("testString")
	model.Password = core.StringPtr("testString")

	result, err := backuprecovery.DataSourceIbmBackupRecoveryRegistrationInfoCredentialsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryRegistrationInfoPhysicalParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		timeModel := make(map[string]interface{})
		timeModel["hour"] = int(38)
		timeModel["minute"] = int(38)

		dayTimeParamsModel := make(map[string]interface{})
		dayTimeParamsModel["day"] = "kSunday"
		dayTimeParamsModel["time"] = []map[string]interface{}{timeModel}

		dayTimeWindowModel := make(map[string]interface{})
		dayTimeWindowModel["end_time"] = []map[string]interface{}{dayTimeParamsModel}
		dayTimeWindowModel["start_time"] = []map[string]interface{}{dayTimeParamsModel}

		throttlingWindowModel := make(map[string]interface{})
		throttlingWindowModel["day_time_window"] = []map[string]interface{}{dayTimeWindowModel}
		throttlingWindowModel["threshold"] = int(26)

		throttlingConfigurationModel := make(map[string]interface{})
		throttlingConfigurationModel["fixed_threshold"] = int(26)
		throttlingConfigurationModel["pattern_type"] = "kNoThrottling"
		throttlingConfigurationModel["throttling_windows"] = []map[string]interface{}{throttlingWindowModel}

		sourceThrottlingConfigurationModel := make(map[string]interface{})
		sourceThrottlingConfigurationModel["cpu_throttling_config"] = []map[string]interface{}{throttlingConfigurationModel}
		sourceThrottlingConfigurationModel["network_throttling_config"] = []map[string]interface{}{throttlingConfigurationModel}

		model := make(map[string]interface{})
		model["applications"] = []string{"kVMware"}
		model["password"] = "testString"
		model["throttling_config"] = []map[string]interface{}{sourceThrottlingConfigurationModel}
		model["username"] = "testString"

		assert.Equal(t, result, model)
	}

	timeModel := new(backuprecoveryv1.Time)
	timeModel.Hour = core.Int64Ptr(int64(38))
	timeModel.Minute = core.Int64Ptr(int64(38))

	dayTimeParamsModel := new(backuprecoveryv1.DayTimeParams)
	dayTimeParamsModel.Day = core.StringPtr("kSunday")
	dayTimeParamsModel.Time = timeModel

	dayTimeWindowModel := new(backuprecoveryv1.DayTimeWindow)
	dayTimeWindowModel.EndTime = dayTimeParamsModel
	dayTimeWindowModel.StartTime = dayTimeParamsModel

	throttlingWindowModel := new(backuprecoveryv1.ThrottlingWindow)
	throttlingWindowModel.DayTimeWindow = dayTimeWindowModel
	throttlingWindowModel.Threshold = core.Int64Ptr(int64(26))

	throttlingConfigurationModel := new(backuprecoveryv1.ThrottlingConfiguration)
	throttlingConfigurationModel.FixedThreshold = core.Int64Ptr(int64(26))
	throttlingConfigurationModel.PatternType = core.StringPtr("kNoThrottling")
	throttlingConfigurationModel.ThrottlingWindows = []backuprecoveryv1.ThrottlingWindow{*throttlingWindowModel}

	sourceThrottlingConfigurationModel := new(backuprecoveryv1.SourceThrottlingConfiguration)
	sourceThrottlingConfigurationModel.CpuThrottlingConfig = throttlingConfigurationModel
	sourceThrottlingConfigurationModel.NetworkThrottlingConfig = throttlingConfigurationModel

	model := new(backuprecoveryv1.PhysicalParams)
	model.Applications = []string{"kVMware"}
	model.Password = core.StringPtr("testString")
	model.ThrottlingConfig = sourceThrottlingConfigurationModel
	model.Username = core.StringPtr("testString")

	result, err := backuprecovery.DataSourceIbmBackupRecoveryRegistrationInfoPhysicalParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryRegistrationInfoSourceThrottlingConfigurationToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		timeModel := make(map[string]interface{})
		timeModel["hour"] = int(38)
		timeModel["minute"] = int(38)

		dayTimeParamsModel := make(map[string]interface{})
		dayTimeParamsModel["day"] = "kSunday"
		dayTimeParamsModel["time"] = []map[string]interface{}{timeModel}

		dayTimeWindowModel := make(map[string]interface{})
		dayTimeWindowModel["end_time"] = []map[string]interface{}{dayTimeParamsModel}
		dayTimeWindowModel["start_time"] = []map[string]interface{}{dayTimeParamsModel}

		throttlingWindowModel := make(map[string]interface{})
		throttlingWindowModel["day_time_window"] = []map[string]interface{}{dayTimeWindowModel}
		throttlingWindowModel["threshold"] = int(26)

		throttlingConfigurationModel := make(map[string]interface{})
		throttlingConfigurationModel["fixed_threshold"] = int(26)
		throttlingConfigurationModel["pattern_type"] = "kNoThrottling"
		throttlingConfigurationModel["throttling_windows"] = []map[string]interface{}{throttlingWindowModel}

		model := make(map[string]interface{})
		model["cpu_throttling_config"] = []map[string]interface{}{throttlingConfigurationModel}
		model["network_throttling_config"] = []map[string]interface{}{throttlingConfigurationModel}

		assert.Equal(t, result, model)
	}

	timeModel := new(backuprecoveryv1.Time)
	timeModel.Hour = core.Int64Ptr(int64(38))
	timeModel.Minute = core.Int64Ptr(int64(38))

	dayTimeParamsModel := new(backuprecoveryv1.DayTimeParams)
	dayTimeParamsModel.Day = core.StringPtr("kSunday")
	dayTimeParamsModel.Time = timeModel

	dayTimeWindowModel := new(backuprecoveryv1.DayTimeWindow)
	dayTimeWindowModel.EndTime = dayTimeParamsModel
	dayTimeWindowModel.StartTime = dayTimeParamsModel

	throttlingWindowModel := new(backuprecoveryv1.ThrottlingWindow)
	throttlingWindowModel.DayTimeWindow = dayTimeWindowModel
	throttlingWindowModel.Threshold = core.Int64Ptr(int64(26))

	throttlingConfigurationModel := new(backuprecoveryv1.ThrottlingConfiguration)
	throttlingConfigurationModel.FixedThreshold = core.Int64Ptr(int64(26))
	throttlingConfigurationModel.PatternType = core.StringPtr("kNoThrottling")
	throttlingConfigurationModel.ThrottlingWindows = []backuprecoveryv1.ThrottlingWindow{*throttlingWindowModel}

	model := new(backuprecoveryv1.SourceThrottlingConfiguration)
	model.CpuThrottlingConfig = throttlingConfigurationModel
	model.NetworkThrottlingConfig = throttlingConfigurationModel

	result, err := backuprecovery.DataSourceIbmBackupRecoveryRegistrationInfoSourceThrottlingConfigurationToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryRegistrationInfoThrottlingConfigurationToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		timeModel := make(map[string]interface{})
		timeModel["hour"] = int(38)
		timeModel["minute"] = int(38)

		dayTimeParamsModel := make(map[string]interface{})
		dayTimeParamsModel["day"] = "kSunday"
		dayTimeParamsModel["time"] = []map[string]interface{}{timeModel}

		dayTimeWindowModel := make(map[string]interface{})
		dayTimeWindowModel["end_time"] = []map[string]interface{}{dayTimeParamsModel}
		dayTimeWindowModel["start_time"] = []map[string]interface{}{dayTimeParamsModel}

		throttlingWindowModel := make(map[string]interface{})
		throttlingWindowModel["day_time_window"] = []map[string]interface{}{dayTimeWindowModel}
		throttlingWindowModel["threshold"] = int(26)

		model := make(map[string]interface{})
		model["fixed_threshold"] = int(26)
		model["pattern_type"] = "kNoThrottling"
		model["throttling_windows"] = []map[string]interface{}{throttlingWindowModel}

		assert.Equal(t, result, model)
	}

	timeModel := new(backuprecoveryv1.Time)
	timeModel.Hour = core.Int64Ptr(int64(38))
	timeModel.Minute = core.Int64Ptr(int64(38))

	dayTimeParamsModel := new(backuprecoveryv1.DayTimeParams)
	dayTimeParamsModel.Day = core.StringPtr("kSunday")
	dayTimeParamsModel.Time = timeModel

	dayTimeWindowModel := new(backuprecoveryv1.DayTimeWindow)
	dayTimeWindowModel.EndTime = dayTimeParamsModel
	dayTimeWindowModel.StartTime = dayTimeParamsModel

	throttlingWindowModel := new(backuprecoveryv1.ThrottlingWindow)
	throttlingWindowModel.DayTimeWindow = dayTimeWindowModel
	throttlingWindowModel.Threshold = core.Int64Ptr(int64(26))

	model := new(backuprecoveryv1.ThrottlingConfiguration)
	model.FixedThreshold = core.Int64Ptr(int64(26))
	model.PatternType = core.StringPtr("kNoThrottling")
	model.ThrottlingWindows = []backuprecoveryv1.ThrottlingWindow{*throttlingWindowModel}

	result, err := backuprecovery.DataSourceIbmBackupRecoveryRegistrationInfoThrottlingConfigurationToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryRegistrationInfoSfdcParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["access_token"] = "testString"
		model["concurrent_api_requests_limit"] = int(26)
		model["consumer_key"] = "testString"
		model["consumer_secret"] = "testString"
		model["daily_api_limit"] = int(26)
		model["endpoint"] = "testString"
		model["endpoint_type"] = "PROD"
		model["metadata_endpoint_url"] = "testString"
		model["refresh_token"] = "testString"
		model["soap_endpoint_url"] = "testString"
		model["use_bulk_api"] = true

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.SfdcParams)
	model.AccessToken = core.StringPtr("testString")
	model.ConcurrentApiRequestsLimit = core.Int64Ptr(int64(26))
	model.ConsumerKey = core.StringPtr("testString")
	model.ConsumerSecret = core.StringPtr("testString")
	model.DailyApiLimit = core.Int64Ptr(int64(26))
	model.Endpoint = core.StringPtr("testString")
	model.EndpointType = core.StringPtr("PROD")
	model.MetadataEndpointURL = core.StringPtr("testString")
	model.RefreshToken = core.StringPtr("testString")
	model.SoapEndpointURL = core.StringPtr("testString")
	model.UseBulkApi = core.BoolPtr(true)

	result, err := backuprecovery.DataSourceIbmBackupRecoveryRegistrationInfoSfdcParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryRegistrationInfoThrottlingPolicyParametersToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		latencyThresholdsModel := make(map[string]interface{})
		latencyThresholdsModel["active_task_msecs"] = int(26)
		latencyThresholdsModel["new_task_msecs"] = int(26)

		nasSourceThrottlingParamsModel := make(map[string]interface{})
		nasSourceThrottlingParamsModel["max_parallel_metadata_fetch_full_percentage"] = int(38)
		nasSourceThrottlingParamsModel["max_parallel_metadata_fetch_incremental_percentage"] = int(38)
		nasSourceThrottlingParamsModel["max_parallel_read_write_full_percentage"] = int(38)
		nasSourceThrottlingParamsModel["max_parallel_read_write_incremental_percentage"] = int(38)

		storageArraySnapshotMaxSpaceConfigParamsModel := make(map[string]interface{})
		storageArraySnapshotMaxSpaceConfigParamsModel["max_snapshot_space_percentage"] = int(38)

		storageArraySnapshotMaxSnapshotConfigParamsModel := make(map[string]interface{})
		storageArraySnapshotMaxSnapshotConfigParamsModel["max_snapshots"] = int(38)

		storageArraySnapshotThrottlingPolicyModel := make(map[string]interface{})
		storageArraySnapshotThrottlingPolicyModel["id"] = int(26)
		storageArraySnapshotThrottlingPolicyModel["is_max_snapshots_config_enabled"] = true
		storageArraySnapshotThrottlingPolicyModel["is_max_space_config_enabled"] = true
		storageArraySnapshotThrottlingPolicyModel["max_snapshot_config"] = []map[string]interface{}{storageArraySnapshotMaxSnapshotConfigParamsModel}
		storageArraySnapshotThrottlingPolicyModel["max_space_config"] = []map[string]interface{}{storageArraySnapshotMaxSpaceConfigParamsModel}

		storageArraySnapshotConfigParamsModel := make(map[string]interface{})
		storageArraySnapshotConfigParamsModel["is_max_snapshots_config_enabled"] = true
		storageArraySnapshotConfigParamsModel["is_max_space_config_enabled"] = true
		storageArraySnapshotConfigParamsModel["storage_array_snapshot_max_space_config"] = []map[string]interface{}{storageArraySnapshotMaxSpaceConfigParamsModel}
		storageArraySnapshotConfigParamsModel["storage_array_snapshot_throttling_policies"] = []map[string]interface{}{storageArraySnapshotThrottlingPolicyModel}

		model := make(map[string]interface{})
		model["enforce_max_streams"] = true
		model["enforce_registered_source_max_backups"] = true
		model["is_enabled"] = true
		model["latency_thresholds"] = []map[string]interface{}{latencyThresholdsModel}
		model["max_concurrent_streams"] = int(38)
		model["nas_source_params"] = []map[string]interface{}{nasSourceThrottlingParamsModel}
		model["registered_source_max_concurrent_backups"] = int(38)
		model["storage_array_snapshot_config"] = []map[string]interface{}{storageArraySnapshotConfigParamsModel}

		assert.Equal(t, result, model)
	}

	latencyThresholdsModel := new(backuprecoveryv1.LatencyThresholds)
	latencyThresholdsModel.ActiveTaskMsecs = core.Int64Ptr(int64(26))
	latencyThresholdsModel.NewTaskMsecs = core.Int64Ptr(int64(26))

	nasSourceThrottlingParamsModel := new(backuprecoveryv1.NasSourceThrottlingParams)
	nasSourceThrottlingParamsModel.MaxParallelMetadataFetchFullPercentage = core.Int64Ptr(int64(38))
	nasSourceThrottlingParamsModel.MaxParallelMetadataFetchIncrementalPercentage = core.Int64Ptr(int64(38))
	nasSourceThrottlingParamsModel.MaxParallelReadWriteFullPercentage = core.Int64Ptr(int64(38))
	nasSourceThrottlingParamsModel.MaxParallelReadWriteIncrementalPercentage = core.Int64Ptr(int64(38))

	storageArraySnapshotMaxSpaceConfigParamsModel := new(backuprecoveryv1.StorageArraySnapshotMaxSpaceConfigParams)
	storageArraySnapshotMaxSpaceConfigParamsModel.MaxSnapshotSpacePercentage = core.Int64Ptr(int64(38))

	storageArraySnapshotMaxSnapshotConfigParamsModel := new(backuprecoveryv1.StorageArraySnapshotMaxSnapshotConfigParams)
	storageArraySnapshotMaxSnapshotConfigParamsModel.MaxSnapshots = core.Int64Ptr(int64(38))

	storageArraySnapshotThrottlingPolicyModel := new(backuprecoveryv1.StorageArraySnapshotThrottlingPolicy)
	storageArraySnapshotThrottlingPolicyModel.ID = core.Int64Ptr(int64(26))
	storageArraySnapshotThrottlingPolicyModel.IsMaxSnapshotsConfigEnabled = core.BoolPtr(true)
	storageArraySnapshotThrottlingPolicyModel.IsMaxSpaceConfigEnabled = core.BoolPtr(true)
	storageArraySnapshotThrottlingPolicyModel.MaxSnapshotConfig = storageArraySnapshotMaxSnapshotConfigParamsModel
	storageArraySnapshotThrottlingPolicyModel.MaxSpaceConfig = storageArraySnapshotMaxSpaceConfigParamsModel

	storageArraySnapshotConfigParamsModel := new(backuprecoveryv1.StorageArraySnapshotConfigParams)
	storageArraySnapshotConfigParamsModel.IsMaxSnapshotsConfigEnabled = core.BoolPtr(true)
	storageArraySnapshotConfigParamsModel.IsMaxSpaceConfigEnabled = core.BoolPtr(true)
	storageArraySnapshotConfigParamsModel.StorageArraySnapshotMaxSpaceConfig = storageArraySnapshotMaxSpaceConfigParamsModel
	storageArraySnapshotConfigParamsModel.StorageArraySnapshotThrottlingPolicies = []backuprecoveryv1.StorageArraySnapshotThrottlingPolicy{*storageArraySnapshotThrottlingPolicyModel}

	model := new(backuprecoveryv1.ThrottlingPolicyParameters)
	model.EnforceMaxStreams = core.BoolPtr(true)
	model.EnforceRegisteredSourceMaxBackups = core.BoolPtr(true)
	model.IsEnabled = core.BoolPtr(true)
	model.LatencyThresholds = latencyThresholdsModel
	model.MaxConcurrentStreams = core.Int64Ptr(int64(38))
	model.NasSourceParams = nasSourceThrottlingParamsModel
	model.RegisteredSourceMaxConcurrentBackups = core.Int64Ptr(int64(38))
	model.StorageArraySnapshotConfig = storageArraySnapshotConfigParamsModel

	result, err := backuprecovery.DataSourceIbmBackupRecoveryRegistrationInfoThrottlingPolicyParametersToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryRegistrationInfoNasSourceThrottlingParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["max_parallel_metadata_fetch_full_percentage"] = int(38)
		model["max_parallel_metadata_fetch_incremental_percentage"] = int(38)
		model["max_parallel_read_write_full_percentage"] = int(38)
		model["max_parallel_read_write_incremental_percentage"] = int(38)

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.NasSourceThrottlingParams)
	model.MaxParallelMetadataFetchFullPercentage = core.Int64Ptr(int64(38))
	model.MaxParallelMetadataFetchIncrementalPercentage = core.Int64Ptr(int64(38))
	model.MaxParallelReadWriteFullPercentage = core.Int64Ptr(int64(38))
	model.MaxParallelReadWriteIncrementalPercentage = core.Int64Ptr(int64(38))

	result, err := backuprecovery.DataSourceIbmBackupRecoveryRegistrationInfoNasSourceThrottlingParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryRegistrationInfoStorageArraySnapshotConfigParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		storageArraySnapshotMaxSpaceConfigParamsModel := make(map[string]interface{})
		storageArraySnapshotMaxSpaceConfigParamsModel["max_snapshot_space_percentage"] = int(38)

		storageArraySnapshotMaxSnapshotConfigParamsModel := make(map[string]interface{})
		storageArraySnapshotMaxSnapshotConfigParamsModel["max_snapshots"] = int(38)

		storageArraySnapshotThrottlingPolicyModel := make(map[string]interface{})
		storageArraySnapshotThrottlingPolicyModel["id"] = int(26)
		storageArraySnapshotThrottlingPolicyModel["is_max_snapshots_config_enabled"] = true
		storageArraySnapshotThrottlingPolicyModel["is_max_space_config_enabled"] = true
		storageArraySnapshotThrottlingPolicyModel["max_snapshot_config"] = []map[string]interface{}{storageArraySnapshotMaxSnapshotConfigParamsModel}
		storageArraySnapshotThrottlingPolicyModel["max_space_config"] = []map[string]interface{}{storageArraySnapshotMaxSpaceConfigParamsModel}

		model := make(map[string]interface{})
		model["is_max_snapshots_config_enabled"] = true
		model["is_max_space_config_enabled"] = true
		model["storage_array_snapshot_max_space_config"] = []map[string]interface{}{storageArraySnapshotMaxSpaceConfigParamsModel}
		model["storage_array_snapshot_throttling_policies"] = []map[string]interface{}{storageArraySnapshotThrottlingPolicyModel}

		assert.Equal(t, result, model)
	}

	storageArraySnapshotMaxSpaceConfigParamsModel := new(backuprecoveryv1.StorageArraySnapshotMaxSpaceConfigParams)
	storageArraySnapshotMaxSpaceConfigParamsModel.MaxSnapshotSpacePercentage = core.Int64Ptr(int64(38))

	storageArraySnapshotMaxSnapshotConfigParamsModel := new(backuprecoveryv1.StorageArraySnapshotMaxSnapshotConfigParams)
	storageArraySnapshotMaxSnapshotConfigParamsModel.MaxSnapshots = core.Int64Ptr(int64(38))

	storageArraySnapshotThrottlingPolicyModel := new(backuprecoveryv1.StorageArraySnapshotThrottlingPolicy)
	storageArraySnapshotThrottlingPolicyModel.ID = core.Int64Ptr(int64(26))
	storageArraySnapshotThrottlingPolicyModel.IsMaxSnapshotsConfigEnabled = core.BoolPtr(true)
	storageArraySnapshotThrottlingPolicyModel.IsMaxSpaceConfigEnabled = core.BoolPtr(true)
	storageArraySnapshotThrottlingPolicyModel.MaxSnapshotConfig = storageArraySnapshotMaxSnapshotConfigParamsModel
	storageArraySnapshotThrottlingPolicyModel.MaxSpaceConfig = storageArraySnapshotMaxSpaceConfigParamsModel

	model := new(backuprecoveryv1.StorageArraySnapshotConfigParams)
	model.IsMaxSnapshotsConfigEnabled = core.BoolPtr(true)
	model.IsMaxSpaceConfigEnabled = core.BoolPtr(true)
	model.StorageArraySnapshotMaxSpaceConfig = storageArraySnapshotMaxSpaceConfigParamsModel
	model.StorageArraySnapshotThrottlingPolicies = []backuprecoveryv1.StorageArraySnapshotThrottlingPolicy{*storageArraySnapshotThrottlingPolicyModel}

	result, err := backuprecovery.DataSourceIbmBackupRecoveryRegistrationInfoStorageArraySnapshotConfigParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryRegistrationInfoStorageArraySnapshotMaxSpaceConfigParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["max_snapshot_space_percentage"] = int(38)

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.StorageArraySnapshotMaxSpaceConfigParams)
	model.MaxSnapshotSpacePercentage = core.Int64Ptr(int64(38))

	result, err := backuprecovery.DataSourceIbmBackupRecoveryRegistrationInfoStorageArraySnapshotMaxSpaceConfigParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryRegistrationInfoStorageArraySnapshotThrottlingPolicyToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		storageArraySnapshotMaxSnapshotConfigParamsModel := make(map[string]interface{})
		storageArraySnapshotMaxSnapshotConfigParamsModel["max_snapshots"] = int(38)

		storageArraySnapshotMaxSpaceConfigParamsModel := make(map[string]interface{})
		storageArraySnapshotMaxSpaceConfigParamsModel["max_snapshot_space_percentage"] = int(38)

		model := make(map[string]interface{})
		model["id"] = int(26)
		model["is_max_snapshots_config_enabled"] = true
		model["is_max_space_config_enabled"] = true
		model["max_snapshot_config"] = []map[string]interface{}{storageArraySnapshotMaxSnapshotConfigParamsModel}
		model["max_space_config"] = []map[string]interface{}{storageArraySnapshotMaxSpaceConfigParamsModel}

		assert.Equal(t, result, model)
	}

	storageArraySnapshotMaxSnapshotConfigParamsModel := new(backuprecoveryv1.StorageArraySnapshotMaxSnapshotConfigParams)
	storageArraySnapshotMaxSnapshotConfigParamsModel.MaxSnapshots = core.Int64Ptr(int64(38))

	storageArraySnapshotMaxSpaceConfigParamsModel := new(backuprecoveryv1.StorageArraySnapshotMaxSpaceConfigParams)
	storageArraySnapshotMaxSpaceConfigParamsModel.MaxSnapshotSpacePercentage = core.Int64Ptr(int64(38))

	model := new(backuprecoveryv1.StorageArraySnapshotThrottlingPolicy)
	model.ID = core.Int64Ptr(int64(26))
	model.IsMaxSnapshotsConfigEnabled = core.BoolPtr(true)
	model.IsMaxSpaceConfigEnabled = core.BoolPtr(true)
	model.MaxSnapshotConfig = storageArraySnapshotMaxSnapshotConfigParamsModel
	model.MaxSpaceConfig = storageArraySnapshotMaxSpaceConfigParamsModel

	result, err := backuprecovery.DataSourceIbmBackupRecoveryRegistrationInfoStorageArraySnapshotThrottlingPolicyToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryRegistrationInfoStorageArraySnapshotMaxSnapshotConfigParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["max_snapshots"] = int(38)

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.StorageArraySnapshotMaxSnapshotConfigParams)
	model.MaxSnapshots = core.Int64Ptr(int64(38))

	result, err := backuprecovery.DataSourceIbmBackupRecoveryRegistrationInfoStorageArraySnapshotMaxSnapshotConfigParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryRegistrationInfoThrottlingPolicyOverrideToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		latencyThresholdsModel := make(map[string]interface{})
		latencyThresholdsModel["active_task_msecs"] = int(26)
		latencyThresholdsModel["new_task_msecs"] = int(26)

		nasSourceThrottlingParamsModel := make(map[string]interface{})
		nasSourceThrottlingParamsModel["max_parallel_metadata_fetch_full_percentage"] = int(38)
		nasSourceThrottlingParamsModel["max_parallel_metadata_fetch_incremental_percentage"] = int(38)
		nasSourceThrottlingParamsModel["max_parallel_read_write_full_percentage"] = int(38)
		nasSourceThrottlingParamsModel["max_parallel_read_write_incremental_percentage"] = int(38)

		storageArraySnapshotMaxSpaceConfigParamsModel := make(map[string]interface{})
		storageArraySnapshotMaxSpaceConfigParamsModel["max_snapshot_space_percentage"] = int(38)

		storageArraySnapshotMaxSnapshotConfigParamsModel := make(map[string]interface{})
		storageArraySnapshotMaxSnapshotConfigParamsModel["max_snapshots"] = int(38)

		storageArraySnapshotThrottlingPolicyModel := make(map[string]interface{})
		storageArraySnapshotThrottlingPolicyModel["id"] = int(26)
		storageArraySnapshotThrottlingPolicyModel["is_max_snapshots_config_enabled"] = true
		storageArraySnapshotThrottlingPolicyModel["is_max_space_config_enabled"] = true
		storageArraySnapshotThrottlingPolicyModel["max_snapshot_config"] = []map[string]interface{}{storageArraySnapshotMaxSnapshotConfigParamsModel}
		storageArraySnapshotThrottlingPolicyModel["max_space_config"] = []map[string]interface{}{storageArraySnapshotMaxSpaceConfigParamsModel}

		storageArraySnapshotConfigParamsModel := make(map[string]interface{})
		storageArraySnapshotConfigParamsModel["is_max_snapshots_config_enabled"] = true
		storageArraySnapshotConfigParamsModel["is_max_space_config_enabled"] = true
		storageArraySnapshotConfigParamsModel["storage_array_snapshot_max_space_config"] = []map[string]interface{}{storageArraySnapshotMaxSpaceConfigParamsModel}
		storageArraySnapshotConfigParamsModel["storage_array_snapshot_throttling_policies"] = []map[string]interface{}{storageArraySnapshotThrottlingPolicyModel}

		throttlingPolicyParametersModel := make(map[string]interface{})
		throttlingPolicyParametersModel["enforce_max_streams"] = true
		throttlingPolicyParametersModel["enforce_registered_source_max_backups"] = true
		throttlingPolicyParametersModel["is_enabled"] = true
		throttlingPolicyParametersModel["latency_thresholds"] = []map[string]interface{}{latencyThresholdsModel}
		throttlingPolicyParametersModel["max_concurrent_streams"] = int(38)
		throttlingPolicyParametersModel["nas_source_params"] = []map[string]interface{}{nasSourceThrottlingParamsModel}
		throttlingPolicyParametersModel["registered_source_max_concurrent_backups"] = int(38)
		throttlingPolicyParametersModel["storage_array_snapshot_config"] = []map[string]interface{}{storageArraySnapshotConfigParamsModel}

		model := make(map[string]interface{})
		model["datastore_id"] = int(26)
		model["datastore_name"] = "testString"
		model["throttling_policy"] = []map[string]interface{}{throttlingPolicyParametersModel}

		assert.Equal(t, result, model)
	}

	latencyThresholdsModel := new(backuprecoveryv1.LatencyThresholds)
	latencyThresholdsModel.ActiveTaskMsecs = core.Int64Ptr(int64(26))
	latencyThresholdsModel.NewTaskMsecs = core.Int64Ptr(int64(26))

	nasSourceThrottlingParamsModel := new(backuprecoveryv1.NasSourceThrottlingParams)
	nasSourceThrottlingParamsModel.MaxParallelMetadataFetchFullPercentage = core.Int64Ptr(int64(38))
	nasSourceThrottlingParamsModel.MaxParallelMetadataFetchIncrementalPercentage = core.Int64Ptr(int64(38))
	nasSourceThrottlingParamsModel.MaxParallelReadWriteFullPercentage = core.Int64Ptr(int64(38))
	nasSourceThrottlingParamsModel.MaxParallelReadWriteIncrementalPercentage = core.Int64Ptr(int64(38))

	storageArraySnapshotMaxSpaceConfigParamsModel := new(backuprecoveryv1.StorageArraySnapshotMaxSpaceConfigParams)
	storageArraySnapshotMaxSpaceConfigParamsModel.MaxSnapshotSpacePercentage = core.Int64Ptr(int64(38))

	storageArraySnapshotMaxSnapshotConfigParamsModel := new(backuprecoveryv1.StorageArraySnapshotMaxSnapshotConfigParams)
	storageArraySnapshotMaxSnapshotConfigParamsModel.MaxSnapshots = core.Int64Ptr(int64(38))

	storageArraySnapshotThrottlingPolicyModel := new(backuprecoveryv1.StorageArraySnapshotThrottlingPolicy)
	storageArraySnapshotThrottlingPolicyModel.ID = core.Int64Ptr(int64(26))
	storageArraySnapshotThrottlingPolicyModel.IsMaxSnapshotsConfigEnabled = core.BoolPtr(true)
	storageArraySnapshotThrottlingPolicyModel.IsMaxSpaceConfigEnabled = core.BoolPtr(true)
	storageArraySnapshotThrottlingPolicyModel.MaxSnapshotConfig = storageArraySnapshotMaxSnapshotConfigParamsModel
	storageArraySnapshotThrottlingPolicyModel.MaxSpaceConfig = storageArraySnapshotMaxSpaceConfigParamsModel

	storageArraySnapshotConfigParamsModel := new(backuprecoveryv1.StorageArraySnapshotConfigParams)
	storageArraySnapshotConfigParamsModel.IsMaxSnapshotsConfigEnabled = core.BoolPtr(true)
	storageArraySnapshotConfigParamsModel.IsMaxSpaceConfigEnabled = core.BoolPtr(true)
	storageArraySnapshotConfigParamsModel.StorageArraySnapshotMaxSpaceConfig = storageArraySnapshotMaxSpaceConfigParamsModel
	storageArraySnapshotConfigParamsModel.StorageArraySnapshotThrottlingPolicies = []backuprecoveryv1.StorageArraySnapshotThrottlingPolicy{*storageArraySnapshotThrottlingPolicyModel}

	throttlingPolicyParametersModel := new(backuprecoveryv1.ThrottlingPolicyParameters)
	throttlingPolicyParametersModel.EnforceMaxStreams = core.BoolPtr(true)
	throttlingPolicyParametersModel.EnforceRegisteredSourceMaxBackups = core.BoolPtr(true)
	throttlingPolicyParametersModel.IsEnabled = core.BoolPtr(true)
	throttlingPolicyParametersModel.LatencyThresholds = latencyThresholdsModel
	throttlingPolicyParametersModel.MaxConcurrentStreams = core.Int64Ptr(int64(38))
	throttlingPolicyParametersModel.NasSourceParams = nasSourceThrottlingParamsModel
	throttlingPolicyParametersModel.RegisteredSourceMaxConcurrentBackups = core.Int64Ptr(int64(38))
	throttlingPolicyParametersModel.StorageArraySnapshotConfig = storageArraySnapshotConfigParamsModel

	model := new(backuprecoveryv1.ThrottlingPolicyOverride)
	model.DatastoreID = core.Int64Ptr(int64(26))
	model.DatastoreName = core.StringPtr("testString")
	model.ThrottlingPolicy = throttlingPolicyParametersModel

	result, err := backuprecovery.DataSourceIbmBackupRecoveryRegistrationInfoThrottlingPolicyOverrideToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryRegistrationInfoUdaConnectParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		udaSourceCapabilitiesModel := make(map[string]interface{})
		udaSourceCapabilitiesModel["auto_log_backup"] = true
		udaSourceCapabilitiesModel["dynamic_config"] = true
		udaSourceCapabilitiesModel["entity_support"] = true
		udaSourceCapabilitiesModel["et_log_backup"] = true
		udaSourceCapabilitiesModel["external_disks"] = true
		udaSourceCapabilitiesModel["full_backup"] = true
		udaSourceCapabilitiesModel["incr_backup"] = true
		udaSourceCapabilitiesModel["log_backup"] = true
		udaSourceCapabilitiesModel["multi_object_restore"] = true
		udaSourceCapabilitiesModel["pause_resume_backup"] = true
		udaSourceCapabilitiesModel["post_backup_job_script"] = true
		udaSourceCapabilitiesModel["post_restore_job_script"] = true
		udaSourceCapabilitiesModel["pre_backup_job_script"] = true
		udaSourceCapabilitiesModel["pre_restore_job_script"] = true
		udaSourceCapabilitiesModel["resource_throttling"] = true
		udaSourceCapabilitiesModel["snapfs_cert"] = true

		credentialsModel := make(map[string]interface{})
		credentialsModel["username"] = "testString"
		credentialsModel["password"] = "testString"

		keyValueStrPairModel := make(map[string]interface{})
		keyValueStrPairModel["key"] = "testString"
		keyValueStrPairModel["value"] = "testString"

		model := make(map[string]interface{})
		model["capabilities"] = []map[string]interface{}{udaSourceCapabilitiesModel}
		model["credentials"] = []map[string]interface{}{credentialsModel}
		model["et_enable_log_backup_policy"] = true
		model["et_enable_run_now"] = true
		model["host_type"] = "kLinux"
		model["hosts"] = []string{"testString"}
		model["live_data_view"] = true
		model["live_log_view"] = true
		model["mount_dir"] = "testString"
		model["mount_view"] = true
		model["script_dir"] = "testString"
		model["source_args"] = "testString"
		model["source_registration_arguments"] = []map[string]interface{}{keyValueStrPairModel}
		model["source_type"] = "testString"

		assert.Equal(t, result, model)
	}

	udaSourceCapabilitiesModel := new(backuprecoveryv1.UdaSourceCapabilities)
	udaSourceCapabilitiesModel.AutoLogBackup = core.BoolPtr(true)
	udaSourceCapabilitiesModel.DynamicConfig = core.BoolPtr(true)
	udaSourceCapabilitiesModel.EntitySupport = core.BoolPtr(true)
	udaSourceCapabilitiesModel.EtLogBackup = core.BoolPtr(true)
	udaSourceCapabilitiesModel.ExternalDisks = core.BoolPtr(true)
	udaSourceCapabilitiesModel.FullBackup = core.BoolPtr(true)
	udaSourceCapabilitiesModel.IncrBackup = core.BoolPtr(true)
	udaSourceCapabilitiesModel.LogBackup = core.BoolPtr(true)
	udaSourceCapabilitiesModel.MultiObjectRestore = core.BoolPtr(true)
	udaSourceCapabilitiesModel.PauseResumeBackup = core.BoolPtr(true)
	udaSourceCapabilitiesModel.PostBackupJobScript = core.BoolPtr(true)
	udaSourceCapabilitiesModel.PostRestoreJobScript = core.BoolPtr(true)
	udaSourceCapabilitiesModel.PreBackupJobScript = core.BoolPtr(true)
	udaSourceCapabilitiesModel.PreRestoreJobScript = core.BoolPtr(true)
	udaSourceCapabilitiesModel.ResourceThrottling = core.BoolPtr(true)
	udaSourceCapabilitiesModel.SnapfsCert = core.BoolPtr(true)

	credentialsModel := new(backuprecoveryv1.Credentials)
	credentialsModel.Username = core.StringPtr("testString")
	credentialsModel.Password = core.StringPtr("testString")

	keyValueStrPairModel := new(backuprecoveryv1.KeyValueStrPair)
	keyValueStrPairModel.Key = core.StringPtr("testString")
	keyValueStrPairModel.Value = core.StringPtr("testString")

	model := new(backuprecoveryv1.UdaConnectParams)
	model.Capabilities = udaSourceCapabilitiesModel
	model.Credentials = credentialsModel
	model.EtEnableLogBackupPolicy = core.BoolPtr(true)
	model.EtEnableRunNow = core.BoolPtr(true)
	model.HostType = core.StringPtr("kLinux")
	model.Hosts = []string{"testString"}
	model.LiveDataView = core.BoolPtr(true)
	model.LiveLogView = core.BoolPtr(true)
	model.MountDir = core.StringPtr("testString")
	model.MountView = core.BoolPtr(true)
	model.ScriptDir = core.StringPtr("testString")
	model.SourceArgs = core.StringPtr("testString")
	model.SourceRegistrationArguments = []backuprecoveryv1.KeyValueStrPair{*keyValueStrPairModel}
	model.SourceType = core.StringPtr("testString")

	result, err := backuprecovery.DataSourceIbmBackupRecoveryRegistrationInfoUdaConnectParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryRegistrationInfoUdaSourceCapabilitiesToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["auto_log_backup"] = true
		model["dynamic_config"] = true
		model["entity_support"] = true
		model["et_log_backup"] = true
		model["external_disks"] = true
		model["full_backup"] = true
		model["incr_backup"] = true
		model["log_backup"] = true
		model["multi_object_restore"] = true
		model["pause_resume_backup"] = true
		model["post_backup_job_script"] = true
		model["post_restore_job_script"] = true
		model["pre_backup_job_script"] = true
		model["pre_restore_job_script"] = true
		model["resource_throttling"] = true
		model["snapfs_cert"] = true

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.UdaSourceCapabilities)
	model.AutoLogBackup = core.BoolPtr(true)
	model.DynamicConfig = core.BoolPtr(true)
	model.EntitySupport = core.BoolPtr(true)
	model.EtLogBackup = core.BoolPtr(true)
	model.ExternalDisks = core.BoolPtr(true)
	model.FullBackup = core.BoolPtr(true)
	model.IncrBackup = core.BoolPtr(true)
	model.LogBackup = core.BoolPtr(true)
	model.MultiObjectRestore = core.BoolPtr(true)
	model.PauseResumeBackup = core.BoolPtr(true)
	model.PostBackupJobScript = core.BoolPtr(true)
	model.PostRestoreJobScript = core.BoolPtr(true)
	model.PreBackupJobScript = core.BoolPtr(true)
	model.PreRestoreJobScript = core.BoolPtr(true)
	model.ResourceThrottling = core.BoolPtr(true)
	model.SnapfsCert = core.BoolPtr(true)

	result, err := backuprecovery.DataSourceIbmBackupRecoveryRegistrationInfoUdaSourceCapabilitiesToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryRegistrationInfoKeyValueStrPairToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["key"] = "testString"
		model["value"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.KeyValueStrPair)
	model.Key = core.StringPtr("testString")
	model.Value = core.StringPtr("testString")

	result, err := backuprecovery.DataSourceIbmBackupRecoveryRegistrationInfoKeyValueStrPairToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryRegistrationInfoProtectionSourceTreeInfoStatsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["protected_count"] = int(26)
		model["protected_size"] = int(26)
		model["unprotected_count"] = int(26)
		model["unprotected_size"] = int(26)

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.ProtectionSourceTreeInfoStats)
	model.ProtectedCount = core.Int64Ptr(int64(26))
	model.ProtectedSize = core.Int64Ptr(int64(26))
	model.UnprotectedCount = core.Int64Ptr(int64(26))
	model.UnprotectedSize = core.Int64Ptr(int64(26))

	result, err := backuprecovery.DataSourceIbmBackupRecoveryRegistrationInfoProtectionSourceTreeInfoStatsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryRegistrationInfoProtectionSummaryByEnvToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		protectionSummaryForK8sDistributionsModel := make(map[string]interface{})
		protectionSummaryForK8sDistributionsModel["distribution"] = "kMainline"
		protectionSummaryForK8sDistributionsModel["protected_count"] = int(26)
		protectionSummaryForK8sDistributionsModel["protected_size"] = int(26)
		protectionSummaryForK8sDistributionsModel["total_registered_clusters"] = int(26)
		protectionSummaryForK8sDistributionsModel["unprotected_count"] = int(26)
		protectionSummaryForK8sDistributionsModel["unprotected_size"] = int(26)

		model := make(map[string]interface{})
		model["environment"] = "kVMware"
		model["kubernetes_distribution_stats"] = []map[string]interface{}{protectionSummaryForK8sDistributionsModel}
		model["protected_count"] = int(26)
		model["protected_size"] = int(26)
		model["unprotected_count"] = int(26)
		model["unprotected_size"] = int(26)

		assert.Equal(t, result, model)
	}

	protectionSummaryForK8sDistributionsModel := new(backuprecoveryv1.ProtectionSummaryForK8sDistributions)
	protectionSummaryForK8sDistributionsModel.Distribution = core.StringPtr("kMainline")
	protectionSummaryForK8sDistributionsModel.ProtectedCount = core.Int64Ptr(int64(26))
	protectionSummaryForK8sDistributionsModel.ProtectedSize = core.Int64Ptr(int64(26))
	protectionSummaryForK8sDistributionsModel.TotalRegisteredClusters = core.Int64Ptr(int64(26))
	protectionSummaryForK8sDistributionsModel.UnprotectedCount = core.Int64Ptr(int64(26))
	protectionSummaryForK8sDistributionsModel.UnprotectedSize = core.Int64Ptr(int64(26))

	model := new(backuprecoveryv1.ProtectionSummaryByEnv)
	model.Environment = core.StringPtr("kVMware")
	model.KubernetesDistributionStats = []backuprecoveryv1.ProtectionSummaryForK8sDistributions{*protectionSummaryForK8sDistributionsModel}
	model.ProtectedCount = core.Int64Ptr(int64(26))
	model.ProtectedSize = core.Int64Ptr(int64(26))
	model.UnprotectedCount = core.Int64Ptr(int64(26))
	model.UnprotectedSize = core.Int64Ptr(int64(26))

	result, err := backuprecovery.DataSourceIbmBackupRecoveryRegistrationInfoProtectionSummaryByEnvToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryRegistrationInfoProtectionSummaryForK8sDistributionsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["distribution"] = "kMainline"
		model["protected_count"] = int(26)
		model["protected_size"] = int(26)
		model["total_registered_clusters"] = int(26)
		model["unprotected_count"] = int(26)
		model["unprotected_size"] = int(26)

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.ProtectionSummaryForK8sDistributions)
	model.Distribution = core.StringPtr("kMainline")
	model.ProtectedCount = core.Int64Ptr(int64(26))
	model.ProtectedSize = core.Int64Ptr(int64(26))
	model.TotalRegisteredClusters = core.Int64Ptr(int64(26))
	model.UnprotectedCount = core.Int64Ptr(int64(26))
	model.UnprotectedSize = core.Int64Ptr(int64(26))

	result, err := backuprecovery.DataSourceIbmBackupRecoveryRegistrationInfoProtectionSummaryForK8sDistributionsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryRegistrationInfoListProtectionSourcesRegistrationInfoResponseStatsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["protected_count"] = int(26)
		model["protected_size"] = int(26)
		model["unprotected_count"] = int(26)
		model["unprotected_size"] = int(26)

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.GetRegistrationInfoResponseStats)
	model.ProtectedCount = core.Int64Ptr(int64(26))
	model.ProtectedSize = core.Int64Ptr(int64(26))
	model.UnprotectedCount = core.Int64Ptr(int64(26))
	model.UnprotectedSize = core.Int64Ptr(int64(26))

	result, err := backuprecovery.DataSourceIbmBackupRecoveryRegistrationInfoListProtectionSourcesRegistrationInfoResponseStatsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
