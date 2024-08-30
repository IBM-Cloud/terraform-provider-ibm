// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

package backuprecovery_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/conns"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/backuprecovery"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/stretchr/testify/assert"
	"github.ibm.com/BackupAndRecovery/ibm-backup-recovery-sdk-go/backuprecoveryv1"
)

func TestAccIbmRestorePointsBasic(t *testing.T) {
	var conf backuprecoveryv1.GetRestorePointsInTimeRangeResponse

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmRestorePointsDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmRestorePointsConfigBasic(),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmRestorePointsExists("ibm_restore_points.restore_points_instance", conf),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_restore_points.restore_points",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmRestorePointsConfigBasic() string {
	return fmt.Sprintf(`
		resource "ibm_restore_points" "restore_points_instance" {
		}
	`)
}

func testAccCheckIbmRestorePointsExists(n string, obj backuprecoveryv1.GetRestorePointsInTimeRangeResponse) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		_, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		backupRecoveryClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).BackupRecoveryV1()
		if err != nil {
			return err
		}

		getRestorePointsInTimeRangeOptions := &backuprecoveryv1.GetRestorePointsInTimeRangeOptions{}

		getRestorePointsInTimeRangeResponse, _, err := backupRecoveryClient.GetRestorePointsInTimeRange(getRestorePointsInTimeRangeOptions)
		if err != nil {
			return err
		}

		obj = *getRestorePointsInTimeRangeResponse
		return nil
	}
}

func testAccCheckIbmRestorePointsDestroy(s *terraform.State) error {
	backupRecoveryClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).BackupRecoveryV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_restore_points" {
			continue
		}

		getRestorePointsInTimeRangeOptions := &backuprecoveryv1.GetRestorePointsInTimeRangeOptions{}

		// Try to find the key
		_, response, err := backupRecoveryClient.GetRestorePointsInTimeRange(getRestorePointsInTimeRangeOptions)

		if err == nil {
			return fmt.Errorf("restore_points still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for restore_points (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

func TestResourceIbmRestorePointsFullSnapshotInfoToMap(t *testing.T) {
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

		archivalTargetSummaryInfoModel := make(map[string]interface{})
		archivalTargetSummaryInfoModel["target_id"] = int(26)
		archivalTargetSummaryInfoModel["archival_task_id"] = "testString"
		archivalTargetSummaryInfoModel["target_name"] = "testString"
		archivalTargetSummaryInfoModel["target_type"] = "Tape"
		archivalTargetSummaryInfoModel["usage_type"] = "Archival"
		archivalTargetSummaryInfoModel["ownership_context"] = "Local"
		archivalTargetSummaryInfoModel["tier_settings"] = []map[string]interface{}{archivalTargetTierInfoModel}

		customTagParamsModel := make(map[string]interface{})
		customTagParamsModel["key"] = "testString"
		customTagParamsModel["value"] = "testString"

		awsCloudSpinParamsModel := make(map[string]interface{})
		awsCloudSpinParamsModel["custom_tag_list"] = []map[string]interface{}{customTagParamsModel}
		awsCloudSpinParamsModel["region"] = int(26)
		awsCloudSpinParamsModel["subnet_id"] = int(26)
		awsCloudSpinParamsModel["vpc_id"] = int(26)

		azureCloudSpinParamsModel := make(map[string]interface{})
		azureCloudSpinParamsModel["availability_set_id"] = int(26)
		azureCloudSpinParamsModel["network_resource_group_id"] = int(26)
		azureCloudSpinParamsModel["resource_group_id"] = int(26)
		azureCloudSpinParamsModel["storage_account_id"] = int(26)
		azureCloudSpinParamsModel["storage_container_id"] = int(26)
		azureCloudSpinParamsModel["storage_resource_group_id"] = int(26)
		azureCloudSpinParamsModel["temp_vm_resource_group_id"] = int(26)
		azureCloudSpinParamsModel["temp_vm_storage_account_id"] = int(26)
		azureCloudSpinParamsModel["temp_vm_storage_container_id"] = int(26)
		azureCloudSpinParamsModel["temp_vm_subnet_id"] = int(26)
		azureCloudSpinParamsModel["temp_vm_virtual_network_id"] = int(26)

		cloudSpinTargetModel := make(map[string]interface{})
		cloudSpinTargetModel["aws_params"] = []map[string]interface{}{awsCloudSpinParamsModel}
		cloudSpinTargetModel["azure_params"] = []map[string]interface{}{azureCloudSpinParamsModel}
		cloudSpinTargetModel["id"] = int(26)

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

		objectProtectionStatsSummaryModel := make(map[string]interface{})
		objectProtectionStatsSummaryModel["environment"] = "kPhysical"
		objectProtectionStatsSummaryModel["protected_count"] = int(26)
		objectProtectionStatsSummaryModel["unprotected_count"] = int(26)
		objectProtectionStatsSummaryModel["deleted_protected_count"] = int(26)
		objectProtectionStatsSummaryModel["protected_size_bytes"] = int(26)
		objectProtectionStatsSummaryModel["unprotected_size_bytes"] = int(26)

		userModel := make(map[string]interface{})
		userModel["name"] = "testString"
		userModel["sid"] = "testString"
		userModel["domain"] = "testString"

		groupModel := make(map[string]interface{})
		groupModel["name"] = "testString"
		groupModel["sid"] = "testString"
		groupModel["domain"] = "testString"

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

		tenantModel := make(map[string]interface{})
		tenantModel["description"] = "testString"
		tenantModel["external_vendor_metadata"] = []map[string]interface{}{externalVendorTenantMetadataModel}
		tenantModel["id"] = "testString"
		tenantModel["is_managed_on_helios"] = true
		tenantModel["name"] = "testString"
		tenantModel["network"] = []map[string]interface{}{tenantNetworkModel}
		tenantModel["status"] = "Active"

		permissionInfoModel := make(map[string]interface{})
		permissionInfoModel["object_id"] = int(26)
		permissionInfoModel["users"] = []map[string]interface{}{userModel}
		permissionInfoModel["groups"] = []map[string]interface{}{groupModel}
		permissionInfoModel["tenant"] = []map[string]interface{}{tenantModel}

		aagInfoModel := make(map[string]interface{})
		aagInfoModel["name"] = "testString"
		aagInfoModel["object_id"] = int(26)

		hostInformationModel := make(map[string]interface{})
		hostInformationModel["id"] = "testString"
		hostInformationModel["name"] = "testString"
		hostInformationModel["environment"] = "kPhysical"

		objectMssqlParamsModel := make(map[string]interface{})
		objectMssqlParamsModel["aag_info"] = []map[string]interface{}{aagInfoModel}
		objectMssqlParamsModel["host_info"] = []map[string]interface{}{hostInformationModel}
		objectMssqlParamsModel["is_encrypted"] = true

		objectPhysicalParamsModel := make(map[string]interface{})
		objectPhysicalParamsModel["enable_system_backup"] = true

		objectModel := make(map[string]interface{})
		objectModel["id"] = int(26)
		objectModel["name"] = "testString"
		objectModel["source_id"] = int(26)
		objectModel["source_name"] = "testString"
		objectModel["environment"] = "kPhysical"
		objectModel["object_hash"] = "testString"
		objectModel["object_type"] = "kCluster"
		objectModel["logical_size_bytes"] = int(26)
		objectModel["uuid"] = "testString"
		objectModel["global_id"] = "testString"
		objectModel["protection_type"] = "kAgent"
		objectModel["sharepoint_site_summary"] = []map[string]interface{}{sharepointObjectParamsModel}
		objectModel["os_type"] = "kLinux"
		objectModel["child_objects"] = []map[string]interface{}{objectSummaryModel}
		objectModel["v_center_summary"] = []map[string]interface{}{objectTypeVCenterParamsModel}
		objectModel["windows_cluster_summary"] = []map[string]interface{}{objectTypeWindowsClusterParamsModel}
		objectModel["protection_stats"] = []map[string]interface{}{objectProtectionStatsSummaryModel}
		objectModel["permissions"] = []map[string]interface{}{permissionInfoModel}
		objectModel["mssql_params"] = []map[string]interface{}{objectMssqlParamsModel}
		objectModel["physical_params"] = []map[string]interface{}{objectPhysicalParamsModel}

		restoreInfoModel := make(map[string]interface{})
		restoreInfoModel["archival_target_info"] = []map[string]interface{}{archivalTargetSummaryInfoModel}
		restoreInfoModel["attempt_number"] = int(38)
		restoreInfoModel["cloud_deploy_target"] = []map[string]interface{}{cloudSpinTargetModel}
		restoreInfoModel["cloud_replication_target"] = []map[string]interface{}{cloudSpinTargetModel}
		restoreInfoModel["object_info"] = []map[string]interface{}{objectModel}
		restoreInfoModel["parent_object_info"] = []map[string]interface{}{objectModel}
		restoreInfoModel["protection_group_id"] = "testString"
		restoreInfoModel["run_start_time_usecs"] = int(26)
		restoreInfoModel["snapshot_relative_dir_path"] = "testString"
		restoreInfoModel["view_name"] = "testString"
		restoreInfoModel["vm_had_independent_disks"] = true

		targetScheduleModel := make(map[string]interface{})
		targetScheduleModel["unit"] = "Runs"
		targetScheduleModel["frequency"] = int(1)

		dataLockConfigModel := make(map[string]interface{})
		dataLockConfigModel["mode"] = "Compliance"
		dataLockConfigModel["unit"] = "Days"
		dataLockConfigModel["duration"] = int(1)
		dataLockConfigModel["enable_worm_on_external_target"] = true

		retentionModel := make(map[string]interface{})
		retentionModel["unit"] = "Days"
		retentionModel["duration"] = int(1)
		retentionModel["data_lock_config"] = []map[string]interface{}{dataLockConfigModel}

		cancellationTimeoutParamsModel := make(map[string]interface{})
		cancellationTimeoutParamsModel["timeout_mins"] = int(26)
		cancellationTimeoutParamsModel["backup_type"] = "kRegular"

		logRetentionModel := make(map[string]interface{})
		logRetentionModel["unit"] = "Days"
		logRetentionModel["duration"] = int(0)
		logRetentionModel["data_lock_config"] = []map[string]interface{}{dataLockConfigModel}

		awsTargetConfigModel := make(map[string]interface{})
		awsTargetConfigModel["region"] = int(26)
		awsTargetConfigModel["source_id"] = int(26)

		azureTargetConfigModel := make(map[string]interface{})
		azureTargetConfigModel["resource_group"] = int(26)
		azureTargetConfigModel["source_id"] = int(26)

		remoteTargetConfigModel := make(map[string]interface{})
		remoteTargetConfigModel["cluster_id"] = int(26)

		replicationTargetConfigurationModel := make(map[string]interface{})
		replicationTargetConfigurationModel["schedule"] = []map[string]interface{}{targetScheduleModel}
		replicationTargetConfigurationModel["retention"] = []map[string]interface{}{retentionModel}
		replicationTargetConfigurationModel["copy_on_run_success"] = true
		replicationTargetConfigurationModel["config_id"] = "testString"
		replicationTargetConfigurationModel["backup_run_type"] = "Regular"
		replicationTargetConfigurationModel["run_timeouts"] = []map[string]interface{}{cancellationTimeoutParamsModel}
		replicationTargetConfigurationModel["log_retention"] = []map[string]interface{}{logRetentionModel}
		replicationTargetConfigurationModel["aws_target_config"] = []map[string]interface{}{awsTargetConfigModel}
		replicationTargetConfigurationModel["azure_target_config"] = []map[string]interface{}{azureTargetConfigModel}
		replicationTargetConfigurationModel["target_type"] = "RemoteCluster"
		replicationTargetConfigurationModel["remote_target_config"] = []map[string]interface{}{remoteTargetConfigModel}

		tierLevelSettingsModel := make(map[string]interface{})
		tierLevelSettingsModel["aws_tiering"] = []map[string]interface{}{awsTiersModel}
		tierLevelSettingsModel["azure_tiering"] = []map[string]interface{}{azureTiersModel}
		tierLevelSettingsModel["cloud_platform"] = "AWS"
		tierLevelSettingsModel["google_tiering"] = []map[string]interface{}{googleTiersModel}
		tierLevelSettingsModel["oracle_tiering"] = []map[string]interface{}{oracleTiersModel}

		extendedRetentionScheduleModel := make(map[string]interface{})
		extendedRetentionScheduleModel["unit"] = "Runs"
		extendedRetentionScheduleModel["frequency"] = int(1)

		extendedRetentionPolicyModel := make(map[string]interface{})
		extendedRetentionPolicyModel["schedule"] = []map[string]interface{}{extendedRetentionScheduleModel}
		extendedRetentionPolicyModel["retention"] = []map[string]interface{}{retentionModel}
		extendedRetentionPolicyModel["run_type"] = "Regular"
		extendedRetentionPolicyModel["config_id"] = "testString"

		archivalTargetConfigurationModel := make(map[string]interface{})
		archivalTargetConfigurationModel["schedule"] = []map[string]interface{}{targetScheduleModel}
		archivalTargetConfigurationModel["retention"] = []map[string]interface{}{retentionModel}
		archivalTargetConfigurationModel["copy_on_run_success"] = true
		archivalTargetConfigurationModel["config_id"] = "testString"
		archivalTargetConfigurationModel["backup_run_type"] = "Regular"
		archivalTargetConfigurationModel["run_timeouts"] = []map[string]interface{}{cancellationTimeoutParamsModel}
		archivalTargetConfigurationModel["log_retention"] = []map[string]interface{}{logRetentionModel}
		archivalTargetConfigurationModel["target_id"] = int(26)
		archivalTargetConfigurationModel["tier_settings"] = []map[string]interface{}{tierLevelSettingsModel}
		archivalTargetConfigurationModel["extended_retention"] = []map[string]interface{}{extendedRetentionPolicyModel}

		cloudSpinTargetConfigurationModel := make(map[string]interface{})
		cloudSpinTargetConfigurationModel["schedule"] = []map[string]interface{}{targetScheduleModel}
		cloudSpinTargetConfigurationModel["retention"] = []map[string]interface{}{retentionModel}
		cloudSpinTargetConfigurationModel["copy_on_run_success"] = true
		cloudSpinTargetConfigurationModel["config_id"] = "testString"
		cloudSpinTargetConfigurationModel["backup_run_type"] = "Regular"
		cloudSpinTargetConfigurationModel["run_timeouts"] = []map[string]interface{}{cancellationTimeoutParamsModel}
		cloudSpinTargetConfigurationModel["log_retention"] = []map[string]interface{}{logRetentionModel}
		cloudSpinTargetConfigurationModel["target"] = []map[string]interface{}{cloudSpinTargetModel}

		onpremDeployParamsModel := make(map[string]interface{})
		onpremDeployParamsModel["id"] = int(26)

		onpremDeployTargetConfigurationModel := make(map[string]interface{})
		onpremDeployTargetConfigurationModel["schedule"] = []map[string]interface{}{targetScheduleModel}
		onpremDeployTargetConfigurationModel["retention"] = []map[string]interface{}{retentionModel}
		onpremDeployTargetConfigurationModel["copy_on_run_success"] = true
		onpremDeployTargetConfigurationModel["config_id"] = "testString"
		onpremDeployTargetConfigurationModel["backup_run_type"] = "Regular"
		onpremDeployTargetConfigurationModel["run_timeouts"] = []map[string]interface{}{cancellationTimeoutParamsModel}
		onpremDeployTargetConfigurationModel["log_retention"] = []map[string]interface{}{logRetentionModel}
		onpremDeployTargetConfigurationModel["params"] = []map[string]interface{}{onpremDeployParamsModel}

		rpaasTargetConfigurationModel := make(map[string]interface{})
		rpaasTargetConfigurationModel["schedule"] = []map[string]interface{}{targetScheduleModel}
		rpaasTargetConfigurationModel["retention"] = []map[string]interface{}{retentionModel}
		rpaasTargetConfigurationModel["copy_on_run_success"] = true
		rpaasTargetConfigurationModel["config_id"] = "testString"
		rpaasTargetConfigurationModel["backup_run_type"] = "Regular"
		rpaasTargetConfigurationModel["run_timeouts"] = []map[string]interface{}{cancellationTimeoutParamsModel}
		rpaasTargetConfigurationModel["log_retention"] = []map[string]interface{}{logRetentionModel}
		rpaasTargetConfigurationModel["target_id"] = int(26)
		rpaasTargetConfigurationModel["target_type"] = "Tape"

		targetsConfigurationModel := make(map[string]interface{})
		targetsConfigurationModel["replication_targets"] = []map[string]interface{}{replicationTargetConfigurationModel}
		targetsConfigurationModel["archival_targets"] = []map[string]interface{}{archivalTargetConfigurationModel}
		targetsConfigurationModel["cloud_spin_targets"] = []map[string]interface{}{cloudSpinTargetConfigurationModel}
		targetsConfigurationModel["onprem_deploy_targets"] = []map[string]interface{}{onpremDeployTargetConfigurationModel}
		targetsConfigurationModel["rpaas_targets"] = []map[string]interface{}{rpaasTargetConfigurationModel}

		model := make(map[string]interface{})
		model["restore_info"] = []map[string]interface{}{restoreInfoModel}
		model["targets_configuration"] = []map[string]interface{}{targetsConfigurationModel}

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

	archivalTargetSummaryInfoModel := new(backuprecoveryv1.ArchivalTargetSummaryInfo)
	archivalTargetSummaryInfoModel.TargetID = core.Int64Ptr(int64(26))
	archivalTargetSummaryInfoModel.ArchivalTaskID = core.StringPtr("testString")
	archivalTargetSummaryInfoModel.TargetName = core.StringPtr("testString")
	archivalTargetSummaryInfoModel.TargetType = core.StringPtr("Tape")
	archivalTargetSummaryInfoModel.UsageType = core.StringPtr("Archival")
	archivalTargetSummaryInfoModel.OwnershipContext = core.StringPtr("Local")
	archivalTargetSummaryInfoModel.TierSettings = archivalTargetTierInfoModel

	customTagParamsModel := new(backuprecoveryv1.CustomTagParams)
	customTagParamsModel.Key = core.StringPtr("testString")
	customTagParamsModel.Value = core.StringPtr("testString")

	awsCloudSpinParamsModel := new(backuprecoveryv1.AwsCloudSpinParams)
	awsCloudSpinParamsModel.CustomTagList = []backuprecoveryv1.CustomTagParams{*customTagParamsModel}
	awsCloudSpinParamsModel.Region = core.Int64Ptr(int64(26))
	awsCloudSpinParamsModel.SubnetID = core.Int64Ptr(int64(26))
	awsCloudSpinParamsModel.VpcID = core.Int64Ptr(int64(26))

	azureCloudSpinParamsModel := new(backuprecoveryv1.AzureCloudSpinParams)
	azureCloudSpinParamsModel.AvailabilitySetID = core.Int64Ptr(int64(26))
	azureCloudSpinParamsModel.NetworkResourceGroupID = core.Int64Ptr(int64(26))
	azureCloudSpinParamsModel.ResourceGroupID = core.Int64Ptr(int64(26))
	azureCloudSpinParamsModel.StorageAccountID = core.Int64Ptr(int64(26))
	azureCloudSpinParamsModel.StorageContainerID = core.Int64Ptr(int64(26))
	azureCloudSpinParamsModel.StorageResourceGroupID = core.Int64Ptr(int64(26))
	azureCloudSpinParamsModel.TempVmResourceGroupID = core.Int64Ptr(int64(26))
	azureCloudSpinParamsModel.TempVmStorageAccountID = core.Int64Ptr(int64(26))
	azureCloudSpinParamsModel.TempVmStorageContainerID = core.Int64Ptr(int64(26))
	azureCloudSpinParamsModel.TempVmSubnetID = core.Int64Ptr(int64(26))
	azureCloudSpinParamsModel.TempVmVirtualNetworkID = core.Int64Ptr(int64(26))

	cloudSpinTargetModel := new(backuprecoveryv1.CloudSpinTarget)
	cloudSpinTargetModel.AwsParams = awsCloudSpinParamsModel
	cloudSpinTargetModel.AzureParams = azureCloudSpinParamsModel
	cloudSpinTargetModel.ID = core.Int64Ptr(int64(26))

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

	objectProtectionStatsSummaryModel := new(backuprecoveryv1.ObjectProtectionStatsSummary)
	objectProtectionStatsSummaryModel.Environment = core.StringPtr("kPhysical")
	objectProtectionStatsSummaryModel.ProtectedCount = core.Int64Ptr(int64(26))
	objectProtectionStatsSummaryModel.UnprotectedCount = core.Int64Ptr(int64(26))
	objectProtectionStatsSummaryModel.DeletedProtectedCount = core.Int64Ptr(int64(26))
	objectProtectionStatsSummaryModel.ProtectedSizeBytes = core.Int64Ptr(int64(26))
	objectProtectionStatsSummaryModel.UnprotectedSizeBytes = core.Int64Ptr(int64(26))

	userModel := new(backuprecoveryv1.User)
	userModel.Name = core.StringPtr("testString")
	userModel.Sid = core.StringPtr("testString")
	userModel.Domain = core.StringPtr("testString")

	groupModel := new(backuprecoveryv1.Group)
	groupModel.Name = core.StringPtr("testString")
	groupModel.Sid = core.StringPtr("testString")
	groupModel.Domain = core.StringPtr("testString")

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

	tenantModel := new(backuprecoveryv1.Tenant)
	tenantModel.Description = core.StringPtr("testString")
	tenantModel.ExternalVendorMetadata = externalVendorTenantMetadataModel
	tenantModel.ID = core.StringPtr("testString")
	tenantModel.IsManagedOnHelios = core.BoolPtr(true)
	tenantModel.Name = core.StringPtr("testString")
	tenantModel.Network = tenantNetworkModel
	tenantModel.Status = core.StringPtr("Active")

	permissionInfoModel := new(backuprecoveryv1.PermissionInfo)
	permissionInfoModel.ObjectID = core.Int64Ptr(int64(26))
	permissionInfoModel.Users = []backuprecoveryv1.User{*userModel}
	permissionInfoModel.Groups = []backuprecoveryv1.Group{*groupModel}
	permissionInfoModel.Tenant = tenantModel

	aagInfoModel := new(backuprecoveryv1.AAGInfo)
	aagInfoModel.Name = core.StringPtr("testString")
	aagInfoModel.ObjectID = core.Int64Ptr(int64(26))

	hostInformationModel := new(backuprecoveryv1.HostInformation)
	hostInformationModel.ID = core.StringPtr("testString")
	hostInformationModel.Name = core.StringPtr("testString")
	hostInformationModel.Environment = core.StringPtr("kPhysical")

	objectMssqlParamsModel := new(backuprecoveryv1.ObjectMssqlParams)
	objectMssqlParamsModel.AagInfo = aagInfoModel
	objectMssqlParamsModel.HostInfo = hostInformationModel
	objectMssqlParamsModel.IsEncrypted = core.BoolPtr(true)

	objectPhysicalParamsModel := new(backuprecoveryv1.ObjectPhysicalParams)
	objectPhysicalParamsModel.EnableSystemBackup = core.BoolPtr(true)

	objectModel := new(backuprecoveryv1.Object)
	objectModel.ID = core.Int64Ptr(int64(26))
	objectModel.Name = core.StringPtr("testString")
	objectModel.SourceID = core.Int64Ptr(int64(26))
	objectModel.SourceName = core.StringPtr("testString")
	objectModel.Environment = core.StringPtr("kPhysical")
	objectModel.ObjectHash = core.StringPtr("testString")
	objectModel.ObjectType = core.StringPtr("kCluster")
	objectModel.LogicalSizeBytes = core.Int64Ptr(int64(26))
	objectModel.UUID = core.StringPtr("testString")
	objectModel.GlobalID = core.StringPtr("testString")
	objectModel.ProtectionType = core.StringPtr("kAgent")
	objectModel.SharepointSiteSummary = sharepointObjectParamsModel
	objectModel.OsType = core.StringPtr("kLinux")
	objectModel.ChildObjects = []backuprecoveryv1.ObjectSummary{*objectSummaryModel}
	objectModel.VCenterSummary = objectTypeVCenterParamsModel
	objectModel.WindowsClusterSummary = objectTypeWindowsClusterParamsModel
	objectModel.ProtectionStats = []backuprecoveryv1.ObjectProtectionStatsSummary{*objectProtectionStatsSummaryModel}
	objectModel.Permissions = permissionInfoModel
	objectModel.MssqlParams = objectMssqlParamsModel
	objectModel.PhysicalParams = objectPhysicalParamsModel

	restoreInfoModel := new(backuprecoveryv1.RestoreInfo)
	restoreInfoModel.ArchivalTargetInfo = archivalTargetSummaryInfoModel
	restoreInfoModel.AttemptNumber = core.Int64Ptr(int64(38))
	restoreInfoModel.CloudDeployTarget = cloudSpinTargetModel
	restoreInfoModel.CloudReplicationTarget = cloudSpinTargetModel
	restoreInfoModel.ObjectInfo = objectModel
	restoreInfoModel.ParentObjectInfo = objectModel
	restoreInfoModel.ProtectionGroupID = core.StringPtr("testString")
	restoreInfoModel.RunStartTimeUsecs = core.Int64Ptr(int64(26))
	restoreInfoModel.SnapshotRelativeDirPath = core.StringPtr("testString")
	restoreInfoModel.ViewName = core.StringPtr("testString")
	restoreInfoModel.VmHadIndependentDisks = core.BoolPtr(true)

	targetScheduleModel := new(backuprecoveryv1.TargetSchedule)
	targetScheduleModel.Unit = core.StringPtr("Runs")
	targetScheduleModel.Frequency = core.Int64Ptr(int64(1))

	dataLockConfigModel := new(backuprecoveryv1.DataLockConfig)
	dataLockConfigModel.Mode = core.StringPtr("Compliance")
	dataLockConfigModel.Unit = core.StringPtr("Days")
	dataLockConfigModel.Duration = core.Int64Ptr(int64(1))
	dataLockConfigModel.EnableWormOnExternalTarget = core.BoolPtr(true)

	retentionModel := new(backuprecoveryv1.Retention)
	retentionModel.Unit = core.StringPtr("Days")
	retentionModel.Duration = core.Int64Ptr(int64(1))
	retentionModel.DataLockConfig = dataLockConfigModel

	cancellationTimeoutParamsModel := new(backuprecoveryv1.CancellationTimeoutParams)
	cancellationTimeoutParamsModel.TimeoutMins = core.Int64Ptr(int64(26))
	cancellationTimeoutParamsModel.BackupType = core.StringPtr("kRegular")

	logRetentionModel := new(backuprecoveryv1.LogRetention)
	logRetentionModel.Unit = core.StringPtr("Days")
	logRetentionModel.Duration = core.Int64Ptr(int64(0))
	logRetentionModel.DataLockConfig = dataLockConfigModel

	awsTargetConfigModel := new(backuprecoveryv1.AWSTargetConfig)
	awsTargetConfigModel.Region = core.Int64Ptr(int64(26))
	awsTargetConfigModel.SourceID = core.Int64Ptr(int64(26))

	azureTargetConfigModel := new(backuprecoveryv1.AzureTargetConfig)
	azureTargetConfigModel.ResourceGroup = core.Int64Ptr(int64(26))
	azureTargetConfigModel.SourceID = core.Int64Ptr(int64(26))

	remoteTargetConfigModel := new(backuprecoveryv1.RemoteTargetConfig)
	remoteTargetConfigModel.ClusterID = core.Int64Ptr(int64(26))

	replicationTargetConfigurationModel := new(backuprecoveryv1.ReplicationTargetConfiguration)
	replicationTargetConfigurationModel.Schedule = targetScheduleModel
	replicationTargetConfigurationModel.Retention = retentionModel
	replicationTargetConfigurationModel.CopyOnRunSuccess = core.BoolPtr(true)
	replicationTargetConfigurationModel.ConfigID = core.StringPtr("testString")
	replicationTargetConfigurationModel.BackupRunType = core.StringPtr("Regular")
	replicationTargetConfigurationModel.RunTimeouts = []backuprecoveryv1.CancellationTimeoutParams{*cancellationTimeoutParamsModel}
	replicationTargetConfigurationModel.LogRetention = logRetentionModel
	replicationTargetConfigurationModel.AwsTargetConfig = awsTargetConfigModel
	replicationTargetConfigurationModel.AzureTargetConfig = azureTargetConfigModel
	replicationTargetConfigurationModel.TargetType = core.StringPtr("RemoteCluster")
	replicationTargetConfigurationModel.RemoteTargetConfig = remoteTargetConfigModel

	tierLevelSettingsModel := new(backuprecoveryv1.TierLevelSettings)
	tierLevelSettingsModel.AwsTiering = awsTiersModel
	tierLevelSettingsModel.AzureTiering = azureTiersModel
	tierLevelSettingsModel.CloudPlatform = core.StringPtr("AWS")
	tierLevelSettingsModel.GoogleTiering = googleTiersModel
	tierLevelSettingsModel.OracleTiering = oracleTiersModel

	extendedRetentionScheduleModel := new(backuprecoveryv1.ExtendedRetentionSchedule)
	extendedRetentionScheduleModel.Unit = core.StringPtr("Runs")
	extendedRetentionScheduleModel.Frequency = core.Int64Ptr(int64(1))

	extendedRetentionPolicyModel := new(backuprecoveryv1.ExtendedRetentionPolicy)
	extendedRetentionPolicyModel.Schedule = extendedRetentionScheduleModel
	extendedRetentionPolicyModel.Retention = retentionModel
	extendedRetentionPolicyModel.RunType = core.StringPtr("Regular")
	extendedRetentionPolicyModel.ConfigID = core.StringPtr("testString")

	archivalTargetConfigurationModel := new(backuprecoveryv1.ArchivalTargetConfiguration)
	archivalTargetConfigurationModel.Schedule = targetScheduleModel
	archivalTargetConfigurationModel.Retention = retentionModel
	archivalTargetConfigurationModel.CopyOnRunSuccess = core.BoolPtr(true)
	archivalTargetConfigurationModel.ConfigID = core.StringPtr("testString")
	archivalTargetConfigurationModel.BackupRunType = core.StringPtr("Regular")
	archivalTargetConfigurationModel.RunTimeouts = []backuprecoveryv1.CancellationTimeoutParams{*cancellationTimeoutParamsModel}
	archivalTargetConfigurationModel.LogRetention = logRetentionModel
	archivalTargetConfigurationModel.TargetID = core.Int64Ptr(int64(26))
	archivalTargetConfigurationModel.TierSettings = tierLevelSettingsModel
	archivalTargetConfigurationModel.ExtendedRetention = []backuprecoveryv1.ExtendedRetentionPolicy{*extendedRetentionPolicyModel}

	cloudSpinTargetConfigurationModel := new(backuprecoveryv1.CloudSpinTargetConfiguration)
	cloudSpinTargetConfigurationModel.Schedule = targetScheduleModel
	cloudSpinTargetConfigurationModel.Retention = retentionModel
	cloudSpinTargetConfigurationModel.CopyOnRunSuccess = core.BoolPtr(true)
	cloudSpinTargetConfigurationModel.ConfigID = core.StringPtr("testString")
	cloudSpinTargetConfigurationModel.BackupRunType = core.StringPtr("Regular")
	cloudSpinTargetConfigurationModel.RunTimeouts = []backuprecoveryv1.CancellationTimeoutParams{*cancellationTimeoutParamsModel}
	cloudSpinTargetConfigurationModel.LogRetention = logRetentionModel
	cloudSpinTargetConfigurationModel.Target = cloudSpinTargetModel

	onpremDeployParamsModel := new(backuprecoveryv1.OnpremDeployParams)
	onpremDeployParamsModel.ID = core.Int64Ptr(int64(26))

	onpremDeployTargetConfigurationModel := new(backuprecoveryv1.OnpremDeployTargetConfiguration)
	onpremDeployTargetConfigurationModel.Schedule = targetScheduleModel
	onpremDeployTargetConfigurationModel.Retention = retentionModel
	onpremDeployTargetConfigurationModel.CopyOnRunSuccess = core.BoolPtr(true)
	onpremDeployTargetConfigurationModel.ConfigID = core.StringPtr("testString")
	onpremDeployTargetConfigurationModel.BackupRunType = core.StringPtr("Regular")
	onpremDeployTargetConfigurationModel.RunTimeouts = []backuprecoveryv1.CancellationTimeoutParams{*cancellationTimeoutParamsModel}
	onpremDeployTargetConfigurationModel.LogRetention = logRetentionModel
	onpremDeployTargetConfigurationModel.Params = onpremDeployParamsModel

	rpaasTargetConfigurationModel := new(backuprecoveryv1.RpaasTargetConfiguration)
	rpaasTargetConfigurationModel.Schedule = targetScheduleModel
	rpaasTargetConfigurationModel.Retention = retentionModel
	rpaasTargetConfigurationModel.CopyOnRunSuccess = core.BoolPtr(true)
	rpaasTargetConfigurationModel.ConfigID = core.StringPtr("testString")
	rpaasTargetConfigurationModel.BackupRunType = core.StringPtr("Regular")
	rpaasTargetConfigurationModel.RunTimeouts = []backuprecoveryv1.CancellationTimeoutParams{*cancellationTimeoutParamsModel}
	rpaasTargetConfigurationModel.LogRetention = logRetentionModel
	rpaasTargetConfigurationModel.TargetID = core.Int64Ptr(int64(26))
	rpaasTargetConfigurationModel.TargetType = core.StringPtr("Tape")

	targetsConfigurationModel := new(backuprecoveryv1.TargetsConfiguration)
	targetsConfigurationModel.ReplicationTargets = []backuprecoveryv1.ReplicationTargetConfiguration{*replicationTargetConfigurationModel}
	targetsConfigurationModel.ArchivalTargets = []backuprecoveryv1.ArchivalTargetConfiguration{*archivalTargetConfigurationModel}
	targetsConfigurationModel.CloudSpinTargets = []backuprecoveryv1.CloudSpinTargetConfiguration{*cloudSpinTargetConfigurationModel}
	targetsConfigurationModel.OnpremDeployTargets = []backuprecoveryv1.OnpremDeployTargetConfiguration{*onpremDeployTargetConfigurationModel}
	targetsConfigurationModel.RpaasTargets = []backuprecoveryv1.RpaasTargetConfiguration{*rpaasTargetConfigurationModel}

	model := new(backuprecoveryv1.FullSnapshotInfo)
	model.RestoreInfo = restoreInfoModel
	model.TargetsConfiguration = []backuprecoveryv1.TargetsConfiguration{*targetsConfigurationModel}

	result, err := backuprecovery.ResourceIbmRestorePointsFullSnapshotInfoToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmRestorePointsRestoreInfoToMap(t *testing.T) {
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

		archivalTargetSummaryInfoModel := make(map[string]interface{})
		archivalTargetSummaryInfoModel["target_id"] = int(26)
		archivalTargetSummaryInfoModel["archival_task_id"] = "testString"
		archivalTargetSummaryInfoModel["target_name"] = "testString"
		archivalTargetSummaryInfoModel["target_type"] = "Tape"
		archivalTargetSummaryInfoModel["usage_type"] = "Archival"
		archivalTargetSummaryInfoModel["ownership_context"] = "Local"
		archivalTargetSummaryInfoModel["tier_settings"] = []map[string]interface{}{archivalTargetTierInfoModel}

		customTagParamsModel := make(map[string]interface{})
		customTagParamsModel["key"] = "testString"
		customTagParamsModel["value"] = "testString"

		awsCloudSpinParamsModel := make(map[string]interface{})
		awsCloudSpinParamsModel["custom_tag_list"] = []map[string]interface{}{customTagParamsModel}
		awsCloudSpinParamsModel["region"] = int(26)
		awsCloudSpinParamsModel["subnet_id"] = int(26)
		awsCloudSpinParamsModel["vpc_id"] = int(26)

		azureCloudSpinParamsModel := make(map[string]interface{})
		azureCloudSpinParamsModel["availability_set_id"] = int(26)
		azureCloudSpinParamsModel["network_resource_group_id"] = int(26)
		azureCloudSpinParamsModel["resource_group_id"] = int(26)
		azureCloudSpinParamsModel["storage_account_id"] = int(26)
		azureCloudSpinParamsModel["storage_container_id"] = int(26)
		azureCloudSpinParamsModel["storage_resource_group_id"] = int(26)
		azureCloudSpinParamsModel["temp_vm_resource_group_id"] = int(26)
		azureCloudSpinParamsModel["temp_vm_storage_account_id"] = int(26)
		azureCloudSpinParamsModel["temp_vm_storage_container_id"] = int(26)
		azureCloudSpinParamsModel["temp_vm_subnet_id"] = int(26)
		azureCloudSpinParamsModel["temp_vm_virtual_network_id"] = int(26)

		cloudSpinTargetModel := make(map[string]interface{})
		cloudSpinTargetModel["aws_params"] = []map[string]interface{}{awsCloudSpinParamsModel}
		cloudSpinTargetModel["azure_params"] = []map[string]interface{}{azureCloudSpinParamsModel}
		cloudSpinTargetModel["id"] = int(26)

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

		objectProtectionStatsSummaryModel := make(map[string]interface{})
		objectProtectionStatsSummaryModel["environment"] = "kPhysical"
		objectProtectionStatsSummaryModel["protected_count"] = int(26)
		objectProtectionStatsSummaryModel["unprotected_count"] = int(26)
		objectProtectionStatsSummaryModel["deleted_protected_count"] = int(26)
		objectProtectionStatsSummaryModel["protected_size_bytes"] = int(26)
		objectProtectionStatsSummaryModel["unprotected_size_bytes"] = int(26)

		userModel := make(map[string]interface{})
		userModel["name"] = "testString"
		userModel["sid"] = "testString"
		userModel["domain"] = "testString"

		groupModel := make(map[string]interface{})
		groupModel["name"] = "testString"
		groupModel["sid"] = "testString"
		groupModel["domain"] = "testString"

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

		tenantModel := make(map[string]interface{})
		tenantModel["description"] = "testString"
		tenantModel["external_vendor_metadata"] = []map[string]interface{}{externalVendorTenantMetadataModel}
		tenantModel["id"] = "testString"
		tenantModel["is_managed_on_helios"] = true
		tenantModel["name"] = "testString"
		tenantModel["network"] = []map[string]interface{}{tenantNetworkModel}
		tenantModel["status"] = "Active"

		permissionInfoModel := make(map[string]interface{})
		permissionInfoModel["object_id"] = int(26)
		permissionInfoModel["users"] = []map[string]interface{}{userModel}
		permissionInfoModel["groups"] = []map[string]interface{}{groupModel}
		permissionInfoModel["tenant"] = []map[string]interface{}{tenantModel}

		aagInfoModel := make(map[string]interface{})
		aagInfoModel["name"] = "testString"
		aagInfoModel["object_id"] = int(26)

		hostInformationModel := make(map[string]interface{})
		hostInformationModel["id"] = "testString"
		hostInformationModel["name"] = "testString"
		hostInformationModel["environment"] = "kPhysical"

		objectMssqlParamsModel := make(map[string]interface{})
		objectMssqlParamsModel["aag_info"] = []map[string]interface{}{aagInfoModel}
		objectMssqlParamsModel["host_info"] = []map[string]interface{}{hostInformationModel}
		objectMssqlParamsModel["is_encrypted"] = true

		objectPhysicalParamsModel := make(map[string]interface{})
		objectPhysicalParamsModel["enable_system_backup"] = true

		objectModel := make(map[string]interface{})
		objectModel["id"] = int(26)
		objectModel["name"] = "testString"
		objectModel["source_id"] = int(26)
		objectModel["source_name"] = "testString"
		objectModel["environment"] = "kPhysical"
		objectModel["object_hash"] = "testString"
		objectModel["object_type"] = "kCluster"
		objectModel["logical_size_bytes"] = int(26)
		objectModel["uuid"] = "testString"
		objectModel["global_id"] = "testString"
		objectModel["protection_type"] = "kAgent"
		objectModel["sharepoint_site_summary"] = []map[string]interface{}{sharepointObjectParamsModel}
		objectModel["os_type"] = "kLinux"
		objectModel["child_objects"] = []map[string]interface{}{objectSummaryModel}
		objectModel["v_center_summary"] = []map[string]interface{}{objectTypeVCenterParamsModel}
		objectModel["windows_cluster_summary"] = []map[string]interface{}{objectTypeWindowsClusterParamsModel}
		objectModel["protection_stats"] = []map[string]interface{}{objectProtectionStatsSummaryModel}
		objectModel["permissions"] = []map[string]interface{}{permissionInfoModel}
		objectModel["mssql_params"] = []map[string]interface{}{objectMssqlParamsModel}
		objectModel["physical_params"] = []map[string]interface{}{objectPhysicalParamsModel}

		model := make(map[string]interface{})
		model["archival_target_info"] = []map[string]interface{}{archivalTargetSummaryInfoModel}
		model["attempt_number"] = int(38)
		model["cloud_deploy_target"] = []map[string]interface{}{cloudSpinTargetModel}
		model["cloud_replication_target"] = []map[string]interface{}{cloudSpinTargetModel}
		model["object_info"] = []map[string]interface{}{objectModel}
		model["parent_object_info"] = []map[string]interface{}{objectModel}
		model["protection_group_id"] = "testString"
		model["run_start_time_usecs"] = int(26)
		model["snapshot_relative_dir_path"] = "testString"
		model["view_name"] = "testString"
		model["vm_had_independent_disks"] = true

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

	archivalTargetSummaryInfoModel := new(backuprecoveryv1.ArchivalTargetSummaryInfo)
	archivalTargetSummaryInfoModel.TargetID = core.Int64Ptr(int64(26))
	archivalTargetSummaryInfoModel.ArchivalTaskID = core.StringPtr("testString")
	archivalTargetSummaryInfoModel.TargetName = core.StringPtr("testString")
	archivalTargetSummaryInfoModel.TargetType = core.StringPtr("Tape")
	archivalTargetSummaryInfoModel.UsageType = core.StringPtr("Archival")
	archivalTargetSummaryInfoModel.OwnershipContext = core.StringPtr("Local")
	archivalTargetSummaryInfoModel.TierSettings = archivalTargetTierInfoModel

	customTagParamsModel := new(backuprecoveryv1.CustomTagParams)
	customTagParamsModel.Key = core.StringPtr("testString")
	customTagParamsModel.Value = core.StringPtr("testString")

	awsCloudSpinParamsModel := new(backuprecoveryv1.AwsCloudSpinParams)
	awsCloudSpinParamsModel.CustomTagList = []backuprecoveryv1.CustomTagParams{*customTagParamsModel}
	awsCloudSpinParamsModel.Region = core.Int64Ptr(int64(26))
	awsCloudSpinParamsModel.SubnetID = core.Int64Ptr(int64(26))
	awsCloudSpinParamsModel.VpcID = core.Int64Ptr(int64(26))

	azureCloudSpinParamsModel := new(backuprecoveryv1.AzureCloudSpinParams)
	azureCloudSpinParamsModel.AvailabilitySetID = core.Int64Ptr(int64(26))
	azureCloudSpinParamsModel.NetworkResourceGroupID = core.Int64Ptr(int64(26))
	azureCloudSpinParamsModel.ResourceGroupID = core.Int64Ptr(int64(26))
	azureCloudSpinParamsModel.StorageAccountID = core.Int64Ptr(int64(26))
	azureCloudSpinParamsModel.StorageContainerID = core.Int64Ptr(int64(26))
	azureCloudSpinParamsModel.StorageResourceGroupID = core.Int64Ptr(int64(26))
	azureCloudSpinParamsModel.TempVmResourceGroupID = core.Int64Ptr(int64(26))
	azureCloudSpinParamsModel.TempVmStorageAccountID = core.Int64Ptr(int64(26))
	azureCloudSpinParamsModel.TempVmStorageContainerID = core.Int64Ptr(int64(26))
	azureCloudSpinParamsModel.TempVmSubnetID = core.Int64Ptr(int64(26))
	azureCloudSpinParamsModel.TempVmVirtualNetworkID = core.Int64Ptr(int64(26))

	cloudSpinTargetModel := new(backuprecoveryv1.CloudSpinTarget)
	cloudSpinTargetModel.AwsParams = awsCloudSpinParamsModel
	cloudSpinTargetModel.AzureParams = azureCloudSpinParamsModel
	cloudSpinTargetModel.ID = core.Int64Ptr(int64(26))

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

	objectProtectionStatsSummaryModel := new(backuprecoveryv1.ObjectProtectionStatsSummary)
	objectProtectionStatsSummaryModel.Environment = core.StringPtr("kPhysical")
	objectProtectionStatsSummaryModel.ProtectedCount = core.Int64Ptr(int64(26))
	objectProtectionStatsSummaryModel.UnprotectedCount = core.Int64Ptr(int64(26))
	objectProtectionStatsSummaryModel.DeletedProtectedCount = core.Int64Ptr(int64(26))
	objectProtectionStatsSummaryModel.ProtectedSizeBytes = core.Int64Ptr(int64(26))
	objectProtectionStatsSummaryModel.UnprotectedSizeBytes = core.Int64Ptr(int64(26))

	userModel := new(backuprecoveryv1.User)
	userModel.Name = core.StringPtr("testString")
	userModel.Sid = core.StringPtr("testString")
	userModel.Domain = core.StringPtr("testString")

	groupModel := new(backuprecoveryv1.Group)
	groupModel.Name = core.StringPtr("testString")
	groupModel.Sid = core.StringPtr("testString")
	groupModel.Domain = core.StringPtr("testString")

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

	tenantModel := new(backuprecoveryv1.Tenant)
	tenantModel.Description = core.StringPtr("testString")
	tenantModel.ExternalVendorMetadata = externalVendorTenantMetadataModel
	tenantModel.ID = core.StringPtr("testString")
	tenantModel.IsManagedOnHelios = core.BoolPtr(true)
	tenantModel.Name = core.StringPtr("testString")
	tenantModel.Network = tenantNetworkModel
	tenantModel.Status = core.StringPtr("Active")

	permissionInfoModel := new(backuprecoveryv1.PermissionInfo)
	permissionInfoModel.ObjectID = core.Int64Ptr(int64(26))
	permissionInfoModel.Users = []backuprecoveryv1.User{*userModel}
	permissionInfoModel.Groups = []backuprecoveryv1.Group{*groupModel}
	permissionInfoModel.Tenant = tenantModel

	aagInfoModel := new(backuprecoveryv1.AAGInfo)
	aagInfoModel.Name = core.StringPtr("testString")
	aagInfoModel.ObjectID = core.Int64Ptr(int64(26))

	hostInformationModel := new(backuprecoveryv1.HostInformation)
	hostInformationModel.ID = core.StringPtr("testString")
	hostInformationModel.Name = core.StringPtr("testString")
	hostInformationModel.Environment = core.StringPtr("kPhysical")

	objectMssqlParamsModel := new(backuprecoveryv1.ObjectMssqlParams)
	objectMssqlParamsModel.AagInfo = aagInfoModel
	objectMssqlParamsModel.HostInfo = hostInformationModel
	objectMssqlParamsModel.IsEncrypted = core.BoolPtr(true)

	objectPhysicalParamsModel := new(backuprecoveryv1.ObjectPhysicalParams)
	objectPhysicalParamsModel.EnableSystemBackup = core.BoolPtr(true)

	objectModel := new(backuprecoveryv1.Object)
	objectModel.ID = core.Int64Ptr(int64(26))
	objectModel.Name = core.StringPtr("testString")
	objectModel.SourceID = core.Int64Ptr(int64(26))
	objectModel.SourceName = core.StringPtr("testString")
	objectModel.Environment = core.StringPtr("kPhysical")
	objectModel.ObjectHash = core.StringPtr("testString")
	objectModel.ObjectType = core.StringPtr("kCluster")
	objectModel.LogicalSizeBytes = core.Int64Ptr(int64(26))
	objectModel.UUID = core.StringPtr("testString")
	objectModel.GlobalID = core.StringPtr("testString")
	objectModel.ProtectionType = core.StringPtr("kAgent")
	objectModel.SharepointSiteSummary = sharepointObjectParamsModel
	objectModel.OsType = core.StringPtr("kLinux")
	objectModel.ChildObjects = []backuprecoveryv1.ObjectSummary{*objectSummaryModel}
	objectModel.VCenterSummary = objectTypeVCenterParamsModel
	objectModel.WindowsClusterSummary = objectTypeWindowsClusterParamsModel
	objectModel.ProtectionStats = []backuprecoveryv1.ObjectProtectionStatsSummary{*objectProtectionStatsSummaryModel}
	objectModel.Permissions = permissionInfoModel
	objectModel.MssqlParams = objectMssqlParamsModel
	objectModel.PhysicalParams = objectPhysicalParamsModel

	model := new(backuprecoveryv1.RestoreInfo)
	model.ArchivalTargetInfo = archivalTargetSummaryInfoModel
	model.AttemptNumber = core.Int64Ptr(int64(38))
	model.CloudDeployTarget = cloudSpinTargetModel
	model.CloudReplicationTarget = cloudSpinTargetModel
	model.ObjectInfo = objectModel
	model.ParentObjectInfo = objectModel
	model.ProtectionGroupID = core.StringPtr("testString")
	model.RunStartTimeUsecs = core.Int64Ptr(int64(26))
	model.SnapshotRelativeDirPath = core.StringPtr("testString")
	model.ViewName = core.StringPtr("testString")
	model.VmHadIndependentDisks = core.BoolPtr(true)

	result, err := backuprecovery.ResourceIbmRestorePointsRestoreInfoToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmRestorePointsArchivalTargetSummaryInfoToMap(t *testing.T) {
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

	model := new(backuprecoveryv1.ArchivalTargetSummaryInfo)
	model.TargetID = core.Int64Ptr(int64(26))
	model.ArchivalTaskID = core.StringPtr("testString")
	model.TargetName = core.StringPtr("testString")
	model.TargetType = core.StringPtr("Tape")
	model.UsageType = core.StringPtr("Archival")
	model.OwnershipContext = core.StringPtr("Local")
	model.TierSettings = archivalTargetTierInfoModel

	result, err := backuprecovery.ResourceIbmRestorePointsArchivalTargetSummaryInfoToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmRestorePointsArchivalTargetTierInfoToMap(t *testing.T) {
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

	result, err := backuprecovery.ResourceIbmRestorePointsArchivalTargetTierInfoToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmRestorePointsAWSTiersToMap(t *testing.T) {
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

	result, err := backuprecovery.ResourceIbmRestorePointsAWSTiersToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmRestorePointsAWSTierToMap(t *testing.T) {
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

	result, err := backuprecovery.ResourceIbmRestorePointsAWSTierToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmRestorePointsAzureTiersToMap(t *testing.T) {
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

	result, err := backuprecovery.ResourceIbmRestorePointsAzureTiersToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmRestorePointsAzureTierToMap(t *testing.T) {
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

	result, err := backuprecovery.ResourceIbmRestorePointsAzureTierToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmRestorePointsGoogleTiersToMap(t *testing.T) {
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

	result, err := backuprecovery.ResourceIbmRestorePointsGoogleTiersToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmRestorePointsGoogleTierToMap(t *testing.T) {
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

	result, err := backuprecovery.ResourceIbmRestorePointsGoogleTierToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmRestorePointsOracleTiersToMap(t *testing.T) {
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

	result, err := backuprecovery.ResourceIbmRestorePointsOracleTiersToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmRestorePointsOracleTierToMap(t *testing.T) {
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

	result, err := backuprecovery.ResourceIbmRestorePointsOracleTierToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmRestorePointsCloudSpinTargetToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		customTagParamsModel := make(map[string]interface{})
		customTagParamsModel["key"] = "testString"
		customTagParamsModel["value"] = "testString"

		awsCloudSpinParamsModel := make(map[string]interface{})
		awsCloudSpinParamsModel["custom_tag_list"] = []map[string]interface{}{customTagParamsModel}
		awsCloudSpinParamsModel["region"] = int(26)
		awsCloudSpinParamsModel["subnet_id"] = int(26)
		awsCloudSpinParamsModel["vpc_id"] = int(26)

		azureCloudSpinParamsModel := make(map[string]interface{})
		azureCloudSpinParamsModel["availability_set_id"] = int(26)
		azureCloudSpinParamsModel["network_resource_group_id"] = int(26)
		azureCloudSpinParamsModel["resource_group_id"] = int(26)
		azureCloudSpinParamsModel["storage_account_id"] = int(26)
		azureCloudSpinParamsModel["storage_container_id"] = int(26)
		azureCloudSpinParamsModel["storage_resource_group_id"] = int(26)
		azureCloudSpinParamsModel["temp_vm_resource_group_id"] = int(26)
		azureCloudSpinParamsModel["temp_vm_storage_account_id"] = int(26)
		azureCloudSpinParamsModel["temp_vm_storage_container_id"] = int(26)
		azureCloudSpinParamsModel["temp_vm_subnet_id"] = int(26)
		azureCloudSpinParamsModel["temp_vm_virtual_network_id"] = int(26)

		model := make(map[string]interface{})
		model["aws_params"] = []map[string]interface{}{awsCloudSpinParamsModel}
		model["azure_params"] = []map[string]interface{}{azureCloudSpinParamsModel}
		model["id"] = int(26)
		model["name"] = "testString"

		assert.Equal(t, result, model)
	}

	customTagParamsModel := new(backuprecoveryv1.CustomTagParams)
	customTagParamsModel.Key = core.StringPtr("testString")
	customTagParamsModel.Value = core.StringPtr("testString")

	awsCloudSpinParamsModel := new(backuprecoveryv1.AwsCloudSpinParams)
	awsCloudSpinParamsModel.CustomTagList = []backuprecoveryv1.CustomTagParams{*customTagParamsModel}
	awsCloudSpinParamsModel.Region = core.Int64Ptr(int64(26))
	awsCloudSpinParamsModel.SubnetID = core.Int64Ptr(int64(26))
	awsCloudSpinParamsModel.VpcID = core.Int64Ptr(int64(26))

	azureCloudSpinParamsModel := new(backuprecoveryv1.AzureCloudSpinParams)
	azureCloudSpinParamsModel.AvailabilitySetID = core.Int64Ptr(int64(26))
	azureCloudSpinParamsModel.NetworkResourceGroupID = core.Int64Ptr(int64(26))
	azureCloudSpinParamsModel.ResourceGroupID = core.Int64Ptr(int64(26))
	azureCloudSpinParamsModel.StorageAccountID = core.Int64Ptr(int64(26))
	azureCloudSpinParamsModel.StorageContainerID = core.Int64Ptr(int64(26))
	azureCloudSpinParamsModel.StorageResourceGroupID = core.Int64Ptr(int64(26))
	azureCloudSpinParamsModel.TempVmResourceGroupID = core.Int64Ptr(int64(26))
	azureCloudSpinParamsModel.TempVmStorageAccountID = core.Int64Ptr(int64(26))
	azureCloudSpinParamsModel.TempVmStorageContainerID = core.Int64Ptr(int64(26))
	azureCloudSpinParamsModel.TempVmSubnetID = core.Int64Ptr(int64(26))
	azureCloudSpinParamsModel.TempVmVirtualNetworkID = core.Int64Ptr(int64(26))

	model := new(backuprecoveryv1.CloudSpinTarget)
	model.AwsParams = awsCloudSpinParamsModel
	model.AzureParams = azureCloudSpinParamsModel
	model.ID = core.Int64Ptr(int64(26))
	model.Name = core.StringPtr("testString")

	result, err := backuprecovery.ResourceIbmRestorePointsCloudSpinTargetToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmRestorePointsAwsCloudSpinParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		customTagParamsModel := make(map[string]interface{})
		customTagParamsModel["key"] = "testString"
		customTagParamsModel["value"] = "testString"

		model := make(map[string]interface{})
		model["custom_tag_list"] = []map[string]interface{}{customTagParamsModel}
		model["region"] = int(26)
		model["subnet_id"] = int(26)
		model["vpc_id"] = int(26)

		assert.Equal(t, result, model)
	}

	customTagParamsModel := new(backuprecoveryv1.CustomTagParams)
	customTagParamsModel.Key = core.StringPtr("testString")
	customTagParamsModel.Value = core.StringPtr("testString")

	model := new(backuprecoveryv1.AwsCloudSpinParams)
	model.CustomTagList = []backuprecoveryv1.CustomTagParams{*customTagParamsModel}
	model.Region = core.Int64Ptr(int64(26))
	model.SubnetID = core.Int64Ptr(int64(26))
	model.VpcID = core.Int64Ptr(int64(26))

	result, err := backuprecovery.ResourceIbmRestorePointsAwsCloudSpinParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmRestorePointsCustomTagParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["key"] = "testString"
		model["value"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.CustomTagParams)
	model.Key = core.StringPtr("testString")
	model.Value = core.StringPtr("testString")

	result, err := backuprecovery.ResourceIbmRestorePointsCustomTagParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmRestorePointsAzureCloudSpinParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["availability_set_id"] = int(26)
		model["network_resource_group_id"] = int(26)
		model["resource_group_id"] = int(26)
		model["storage_account_id"] = int(26)
		model["storage_container_id"] = int(26)
		model["storage_resource_group_id"] = int(26)
		model["temp_vm_resource_group_id"] = int(26)
		model["temp_vm_storage_account_id"] = int(26)
		model["temp_vm_storage_container_id"] = int(26)
		model["temp_vm_subnet_id"] = int(26)
		model["temp_vm_virtual_network_id"] = int(26)

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.AzureCloudSpinParams)
	model.AvailabilitySetID = core.Int64Ptr(int64(26))
	model.NetworkResourceGroupID = core.Int64Ptr(int64(26))
	model.ResourceGroupID = core.Int64Ptr(int64(26))
	model.StorageAccountID = core.Int64Ptr(int64(26))
	model.StorageContainerID = core.Int64Ptr(int64(26))
	model.StorageResourceGroupID = core.Int64Ptr(int64(26))
	model.TempVmResourceGroupID = core.Int64Ptr(int64(26))
	model.TempVmStorageAccountID = core.Int64Ptr(int64(26))
	model.TempVmStorageContainerID = core.Int64Ptr(int64(26))
	model.TempVmSubnetID = core.Int64Ptr(int64(26))
	model.TempVmVirtualNetworkID = core.Int64Ptr(int64(26))

	result, err := backuprecovery.ResourceIbmRestorePointsAzureCloudSpinParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmRestorePointsObjectToMap(t *testing.T) {
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

		objectProtectionStatsSummaryModel := make(map[string]interface{})
		objectProtectionStatsSummaryModel["environment"] = "kPhysical"
		objectProtectionStatsSummaryModel["protected_count"] = int(26)
		objectProtectionStatsSummaryModel["unprotected_count"] = int(26)
		objectProtectionStatsSummaryModel["deleted_protected_count"] = int(26)
		objectProtectionStatsSummaryModel["protected_size_bytes"] = int(26)
		objectProtectionStatsSummaryModel["unprotected_size_bytes"] = int(26)

		userModel := make(map[string]interface{})
		userModel["name"] = "testString"
		userModel["sid"] = "testString"
		userModel["domain"] = "testString"

		groupModel := make(map[string]interface{})
		groupModel["name"] = "testString"
		groupModel["sid"] = "testString"
		groupModel["domain"] = "testString"

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

		tenantModel := make(map[string]interface{})
		tenantModel["description"] = "testString"
		tenantModel["external_vendor_metadata"] = []map[string]interface{}{externalVendorTenantMetadataModel}
		tenantModel["id"] = "testString"
		tenantModel["is_managed_on_helios"] = true
		tenantModel["name"] = "testString"
		tenantModel["network"] = []map[string]interface{}{tenantNetworkModel}
		tenantModel["status"] = "Active"

		permissionInfoModel := make(map[string]interface{})
		permissionInfoModel["object_id"] = int(26)
		permissionInfoModel["users"] = []map[string]interface{}{userModel}
		permissionInfoModel["groups"] = []map[string]interface{}{groupModel}
		permissionInfoModel["tenant"] = []map[string]interface{}{tenantModel}

		aagInfoModel := make(map[string]interface{})
		aagInfoModel["name"] = "testString"
		aagInfoModel["object_id"] = int(26)

		hostInformationModel := make(map[string]interface{})
		hostInformationModel["id"] = "testString"
		hostInformationModel["name"] = "testString"
		hostInformationModel["environment"] = "kPhysical"

		objectMssqlParamsModel := make(map[string]interface{})
		objectMssqlParamsModel["aag_info"] = []map[string]interface{}{aagInfoModel}
		objectMssqlParamsModel["host_info"] = []map[string]interface{}{hostInformationModel}
		objectMssqlParamsModel["is_encrypted"] = true

		objectPhysicalParamsModel := make(map[string]interface{})
		objectPhysicalParamsModel["enable_system_backup"] = true

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
		model["protection_stats"] = []map[string]interface{}{objectProtectionStatsSummaryModel}
		model["permissions"] = []map[string]interface{}{permissionInfoModel}
		model["mssql_params"] = []map[string]interface{}{objectMssqlParamsModel}
		model["physical_params"] = []map[string]interface{}{objectPhysicalParamsModel}

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

	objectProtectionStatsSummaryModel := new(backuprecoveryv1.ObjectProtectionStatsSummary)
	objectProtectionStatsSummaryModel.Environment = core.StringPtr("kPhysical")
	objectProtectionStatsSummaryModel.ProtectedCount = core.Int64Ptr(int64(26))
	objectProtectionStatsSummaryModel.UnprotectedCount = core.Int64Ptr(int64(26))
	objectProtectionStatsSummaryModel.DeletedProtectedCount = core.Int64Ptr(int64(26))
	objectProtectionStatsSummaryModel.ProtectedSizeBytes = core.Int64Ptr(int64(26))
	objectProtectionStatsSummaryModel.UnprotectedSizeBytes = core.Int64Ptr(int64(26))

	userModel := new(backuprecoveryv1.User)
	userModel.Name = core.StringPtr("testString")
	userModel.Sid = core.StringPtr("testString")
	userModel.Domain = core.StringPtr("testString")

	groupModel := new(backuprecoveryv1.Group)
	groupModel.Name = core.StringPtr("testString")
	groupModel.Sid = core.StringPtr("testString")
	groupModel.Domain = core.StringPtr("testString")

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

	tenantModel := new(backuprecoveryv1.Tenant)
	tenantModel.Description = core.StringPtr("testString")
	tenantModel.ExternalVendorMetadata = externalVendorTenantMetadataModel
	tenantModel.ID = core.StringPtr("testString")
	tenantModel.IsManagedOnHelios = core.BoolPtr(true)
	tenantModel.Name = core.StringPtr("testString")
	tenantModel.Network = tenantNetworkModel
	tenantModel.Status = core.StringPtr("Active")

	permissionInfoModel := new(backuprecoveryv1.PermissionInfo)
	permissionInfoModel.ObjectID = core.Int64Ptr(int64(26))
	permissionInfoModel.Users = []backuprecoveryv1.User{*userModel}
	permissionInfoModel.Groups = []backuprecoveryv1.Group{*groupModel}
	permissionInfoModel.Tenant = tenantModel

	aagInfoModel := new(backuprecoveryv1.AAGInfo)
	aagInfoModel.Name = core.StringPtr("testString")
	aagInfoModel.ObjectID = core.Int64Ptr(int64(26))

	hostInformationModel := new(backuprecoveryv1.HostInformation)
	hostInformationModel.ID = core.StringPtr("testString")
	hostInformationModel.Name = core.StringPtr("testString")
	hostInformationModel.Environment = core.StringPtr("kPhysical")

	objectMssqlParamsModel := new(backuprecoveryv1.ObjectMssqlParams)
	objectMssqlParamsModel.AagInfo = aagInfoModel
	objectMssqlParamsModel.HostInfo = hostInformationModel
	objectMssqlParamsModel.IsEncrypted = core.BoolPtr(true)

	objectPhysicalParamsModel := new(backuprecoveryv1.ObjectPhysicalParams)
	objectPhysicalParamsModel.EnableSystemBackup = core.BoolPtr(true)

	model := new(backuprecoveryv1.Object)
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
	model.ProtectionStats = []backuprecoveryv1.ObjectProtectionStatsSummary{*objectProtectionStatsSummaryModel}
	model.Permissions = permissionInfoModel
	model.MssqlParams = objectMssqlParamsModel
	model.PhysicalParams = objectPhysicalParamsModel

	result, err := backuprecovery.ResourceIbmRestorePointsObjectToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmRestorePointsSharepointObjectParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["site_web_url"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.SharepointObjectParams)
	model.SiteWebURL = core.StringPtr("testString")

	result, err := backuprecovery.ResourceIbmRestorePointsSharepointObjectParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmRestorePointsObjectSummaryToMap(t *testing.T) {
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

	result, err := backuprecovery.ResourceIbmRestorePointsObjectSummaryToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmRestorePointsObjectTypeVCenterParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["is_cloud_env"] = true

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.ObjectTypeVCenterParams)
	model.IsCloudEnv = core.BoolPtr(true)

	result, err := backuprecovery.ResourceIbmRestorePointsObjectTypeVCenterParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmRestorePointsObjectTypeWindowsClusterParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["cluster_source_type"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.ObjectTypeWindowsClusterParams)
	model.ClusterSourceType = core.StringPtr("testString")

	result, err := backuprecovery.ResourceIbmRestorePointsObjectTypeWindowsClusterParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmRestorePointsObjectProtectionStatsSummaryToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["environment"] = "kPhysical"
		model["protected_count"] = int(26)
		model["unprotected_count"] = int(26)
		model["deleted_protected_count"] = int(26)
		model["protected_size_bytes"] = int(26)
		model["unprotected_size_bytes"] = int(26)

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.ObjectProtectionStatsSummary)
	model.Environment = core.StringPtr("kPhysical")
	model.ProtectedCount = core.Int64Ptr(int64(26))
	model.UnprotectedCount = core.Int64Ptr(int64(26))
	model.DeletedProtectedCount = core.Int64Ptr(int64(26))
	model.ProtectedSizeBytes = core.Int64Ptr(int64(26))
	model.UnprotectedSizeBytes = core.Int64Ptr(int64(26))

	result, err := backuprecovery.ResourceIbmRestorePointsObjectProtectionStatsSummaryToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmRestorePointsPermissionInfoToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		userModel := make(map[string]interface{})
		userModel["name"] = "testString"
		userModel["sid"] = "testString"
		userModel["domain"] = "testString"

		groupModel := make(map[string]interface{})
		groupModel["name"] = "testString"
		groupModel["sid"] = "testString"
		groupModel["domain"] = "testString"

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

		tenantModel := make(map[string]interface{})
		tenantModel["description"] = "testString"
		tenantModel["external_vendor_metadata"] = []map[string]interface{}{externalVendorTenantMetadataModel}
		tenantModel["id"] = "testString"
		tenantModel["is_managed_on_helios"] = true
		tenantModel["name"] = "testString"
		tenantModel["network"] = []map[string]interface{}{tenantNetworkModel}
		tenantModel["status"] = "Active"

		model := make(map[string]interface{})
		model["object_id"] = int(26)
		model["users"] = []map[string]interface{}{userModel}
		model["groups"] = []map[string]interface{}{groupModel}
		model["tenant"] = []map[string]interface{}{tenantModel}

		assert.Equal(t, result, model)
	}

	userModel := new(backuprecoveryv1.User)
	userModel.Name = core.StringPtr("testString")
	userModel.Sid = core.StringPtr("testString")
	userModel.Domain = core.StringPtr("testString")

	groupModel := new(backuprecoveryv1.Group)
	groupModel.Name = core.StringPtr("testString")
	groupModel.Sid = core.StringPtr("testString")
	groupModel.Domain = core.StringPtr("testString")

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

	tenantModel := new(backuprecoveryv1.Tenant)
	tenantModel.Description = core.StringPtr("testString")
	tenantModel.ExternalVendorMetadata = externalVendorTenantMetadataModel
	tenantModel.ID = core.StringPtr("testString")
	tenantModel.IsManagedOnHelios = core.BoolPtr(true)
	tenantModel.Name = core.StringPtr("testString")
	tenantModel.Network = tenantNetworkModel
	tenantModel.Status = core.StringPtr("Active")

	model := new(backuprecoveryv1.PermissionInfo)
	model.ObjectID = core.Int64Ptr(int64(26))
	model.Users = []backuprecoveryv1.User{*userModel}
	model.Groups = []backuprecoveryv1.Group{*groupModel}
	model.Tenant = tenantModel

	result, err := backuprecovery.ResourceIbmRestorePointsPermissionInfoToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmRestorePointsUserToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["name"] = "testString"
		model["sid"] = "testString"
		model["domain"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.User)
	model.Name = core.StringPtr("testString")
	model.Sid = core.StringPtr("testString")
	model.Domain = core.StringPtr("testString")

	result, err := backuprecovery.ResourceIbmRestorePointsUserToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmRestorePointsGroupToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["name"] = "testString"
		model["sid"] = "testString"
		model["domain"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.Group)
	model.Name = core.StringPtr("testString")
	model.Sid = core.StringPtr("testString")
	model.Domain = core.StringPtr("testString")

	result, err := backuprecovery.ResourceIbmRestorePointsGroupToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmRestorePointsTenantToMap(t *testing.T) {
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

	result, err := backuprecovery.ResourceIbmRestorePointsTenantToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmRestorePointsExternalVendorTenantMetadataToMap(t *testing.T) {
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

	result, err := backuprecovery.ResourceIbmRestorePointsExternalVendorTenantMetadataToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmRestorePointsIbmTenantMetadataParamsToMap(t *testing.T) {
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

	result, err := backuprecovery.ResourceIbmRestorePointsIbmTenantMetadataParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmRestorePointsExternalVendorCustomPropertiesToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["key"] = "testString"
		model["value"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.ExternalVendorCustomProperties)
	model.Key = core.StringPtr("testString")
	model.Value = core.StringPtr("testString")

	result, err := backuprecovery.ResourceIbmRestorePointsExternalVendorCustomPropertiesToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmRestorePointsTenantNetworkToMap(t *testing.T) {
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

	result, err := backuprecovery.ResourceIbmRestorePointsTenantNetworkToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmRestorePointsObjectMssqlParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		aagInfoModel := make(map[string]interface{})
		aagInfoModel["name"] = "testString"
		aagInfoModel["object_id"] = int(26)

		hostInformationModel := make(map[string]interface{})
		hostInformationModel["id"] = "testString"
		hostInformationModel["name"] = "testString"
		hostInformationModel["environment"] = "kPhysical"

		model := make(map[string]interface{})
		model["aag_info"] = []map[string]interface{}{aagInfoModel}
		model["host_info"] = []map[string]interface{}{hostInformationModel}
		model["is_encrypted"] = true

		assert.Equal(t, result, model)
	}

	aagInfoModel := new(backuprecoveryv1.AAGInfo)
	aagInfoModel.Name = core.StringPtr("testString")
	aagInfoModel.ObjectID = core.Int64Ptr(int64(26))

	hostInformationModel := new(backuprecoveryv1.HostInformation)
	hostInformationModel.ID = core.StringPtr("testString")
	hostInformationModel.Name = core.StringPtr("testString")
	hostInformationModel.Environment = core.StringPtr("kPhysical")

	model := new(backuprecoveryv1.ObjectMssqlParams)
	model.AagInfo = aagInfoModel
	model.HostInfo = hostInformationModel
	model.IsEncrypted = core.BoolPtr(true)

	result, err := backuprecovery.ResourceIbmRestorePointsObjectMssqlParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmRestorePointsAAGInfoToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["name"] = "testString"
		model["object_id"] = int(26)

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.AAGInfo)
	model.Name = core.StringPtr("testString")
	model.ObjectID = core.Int64Ptr(int64(26))

	result, err := backuprecovery.ResourceIbmRestorePointsAAGInfoToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmRestorePointsHostInformationToMap(t *testing.T) {
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

	result, err := backuprecovery.ResourceIbmRestorePointsHostInformationToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmRestorePointsObjectPhysicalParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["enable_system_backup"] = true

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.ObjectPhysicalParams)
	model.EnableSystemBackup = core.BoolPtr(true)

	result, err := backuprecovery.ResourceIbmRestorePointsObjectPhysicalParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmRestorePointsTargetsConfigurationToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		targetScheduleModel := make(map[string]interface{})
		targetScheduleModel["unit"] = "Runs"
		targetScheduleModel["frequency"] = int(1)

		dataLockConfigModel := make(map[string]interface{})
		dataLockConfigModel["mode"] = "Compliance"
		dataLockConfigModel["unit"] = "Days"
		dataLockConfigModel["duration"] = int(1)
		dataLockConfigModel["enable_worm_on_external_target"] = true

		retentionModel := make(map[string]interface{})
		retentionModel["unit"] = "Days"
		retentionModel["duration"] = int(1)
		retentionModel["data_lock_config"] = []map[string]interface{}{dataLockConfigModel}

		cancellationTimeoutParamsModel := make(map[string]interface{})
		cancellationTimeoutParamsModel["timeout_mins"] = int(26)
		cancellationTimeoutParamsModel["backup_type"] = "kRegular"

		logRetentionModel := make(map[string]interface{})
		logRetentionModel["unit"] = "Days"
		logRetentionModel["duration"] = int(0)
		logRetentionModel["data_lock_config"] = []map[string]interface{}{dataLockConfigModel}

		awsTargetConfigModel := make(map[string]interface{})
		awsTargetConfigModel["region"] = int(26)
		awsTargetConfigModel["source_id"] = int(26)

		azureTargetConfigModel := make(map[string]interface{})
		azureTargetConfigModel["resource_group"] = int(26)
		azureTargetConfigModel["source_id"] = int(26)

		remoteTargetConfigModel := make(map[string]interface{})
		remoteTargetConfigModel["cluster_id"] = int(26)

		replicationTargetConfigurationModel := make(map[string]interface{})
		replicationTargetConfigurationModel["schedule"] = []map[string]interface{}{targetScheduleModel}
		replicationTargetConfigurationModel["retention"] = []map[string]interface{}{retentionModel}
		replicationTargetConfigurationModel["copy_on_run_success"] = true
		replicationTargetConfigurationModel["config_id"] = "testString"
		replicationTargetConfigurationModel["backup_run_type"] = "Regular"
		replicationTargetConfigurationModel["run_timeouts"] = []map[string]interface{}{cancellationTimeoutParamsModel}
		replicationTargetConfigurationModel["log_retention"] = []map[string]interface{}{logRetentionModel}
		replicationTargetConfigurationModel["aws_target_config"] = []map[string]interface{}{awsTargetConfigModel}
		replicationTargetConfigurationModel["azure_target_config"] = []map[string]interface{}{azureTargetConfigModel}
		replicationTargetConfigurationModel["target_type"] = "RemoteCluster"
		replicationTargetConfigurationModel["remote_target_config"] = []map[string]interface{}{remoteTargetConfigModel}

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

		tierLevelSettingsModel := make(map[string]interface{})
		tierLevelSettingsModel["aws_tiering"] = []map[string]interface{}{awsTiersModel}
		tierLevelSettingsModel["azure_tiering"] = []map[string]interface{}{azureTiersModel}
		tierLevelSettingsModel["cloud_platform"] = "AWS"
		tierLevelSettingsModel["google_tiering"] = []map[string]interface{}{googleTiersModel}
		tierLevelSettingsModel["oracle_tiering"] = []map[string]interface{}{oracleTiersModel}

		extendedRetentionScheduleModel := make(map[string]interface{})
		extendedRetentionScheduleModel["unit"] = "Runs"
		extendedRetentionScheduleModel["frequency"] = int(1)

		extendedRetentionPolicyModel := make(map[string]interface{})
		extendedRetentionPolicyModel["schedule"] = []map[string]interface{}{extendedRetentionScheduleModel}
		extendedRetentionPolicyModel["retention"] = []map[string]interface{}{retentionModel}
		extendedRetentionPolicyModel["run_type"] = "Regular"
		extendedRetentionPolicyModel["config_id"] = "testString"

		archivalTargetConfigurationModel := make(map[string]interface{})
		archivalTargetConfigurationModel["schedule"] = []map[string]interface{}{targetScheduleModel}
		archivalTargetConfigurationModel["retention"] = []map[string]interface{}{retentionModel}
		archivalTargetConfigurationModel["copy_on_run_success"] = true
		archivalTargetConfigurationModel["config_id"] = "testString"
		archivalTargetConfigurationModel["backup_run_type"] = "Regular"
		archivalTargetConfigurationModel["run_timeouts"] = []map[string]interface{}{cancellationTimeoutParamsModel}
		archivalTargetConfigurationModel["log_retention"] = []map[string]interface{}{logRetentionModel}
		archivalTargetConfigurationModel["target_id"] = int(26)
		archivalTargetConfigurationModel["tier_settings"] = []map[string]interface{}{tierLevelSettingsModel}
		archivalTargetConfigurationModel["extended_retention"] = []map[string]interface{}{extendedRetentionPolicyModel}

		customTagParamsModel := make(map[string]interface{})
		customTagParamsModel["key"] = "testString"
		customTagParamsModel["value"] = "testString"

		awsCloudSpinParamsModel := make(map[string]interface{})
		awsCloudSpinParamsModel["custom_tag_list"] = []map[string]interface{}{customTagParamsModel}
		awsCloudSpinParamsModel["region"] = int(26)
		awsCloudSpinParamsModel["subnet_id"] = int(26)
		awsCloudSpinParamsModel["vpc_id"] = int(26)

		azureCloudSpinParamsModel := make(map[string]interface{})
		azureCloudSpinParamsModel["availability_set_id"] = int(26)
		azureCloudSpinParamsModel["network_resource_group_id"] = int(26)
		azureCloudSpinParamsModel["resource_group_id"] = int(26)
		azureCloudSpinParamsModel["storage_account_id"] = int(26)
		azureCloudSpinParamsModel["storage_container_id"] = int(26)
		azureCloudSpinParamsModel["storage_resource_group_id"] = int(26)
		azureCloudSpinParamsModel["temp_vm_resource_group_id"] = int(26)
		azureCloudSpinParamsModel["temp_vm_storage_account_id"] = int(26)
		azureCloudSpinParamsModel["temp_vm_storage_container_id"] = int(26)
		azureCloudSpinParamsModel["temp_vm_subnet_id"] = int(26)
		azureCloudSpinParamsModel["temp_vm_virtual_network_id"] = int(26)

		cloudSpinTargetModel := make(map[string]interface{})
		cloudSpinTargetModel["aws_params"] = []map[string]interface{}{awsCloudSpinParamsModel}
		cloudSpinTargetModel["azure_params"] = []map[string]interface{}{azureCloudSpinParamsModel}
		cloudSpinTargetModel["id"] = int(26)

		cloudSpinTargetConfigurationModel := make(map[string]interface{})
		cloudSpinTargetConfigurationModel["schedule"] = []map[string]interface{}{targetScheduleModel}
		cloudSpinTargetConfigurationModel["retention"] = []map[string]interface{}{retentionModel}
		cloudSpinTargetConfigurationModel["copy_on_run_success"] = true
		cloudSpinTargetConfigurationModel["config_id"] = "testString"
		cloudSpinTargetConfigurationModel["backup_run_type"] = "Regular"
		cloudSpinTargetConfigurationModel["run_timeouts"] = []map[string]interface{}{cancellationTimeoutParamsModel}
		cloudSpinTargetConfigurationModel["log_retention"] = []map[string]interface{}{logRetentionModel}
		cloudSpinTargetConfigurationModel["target"] = []map[string]interface{}{cloudSpinTargetModel}

		onpremDeployParamsModel := make(map[string]interface{})
		onpremDeployParamsModel["id"] = int(26)

		onpremDeployTargetConfigurationModel := make(map[string]interface{})
		onpremDeployTargetConfigurationModel["schedule"] = []map[string]interface{}{targetScheduleModel}
		onpremDeployTargetConfigurationModel["retention"] = []map[string]interface{}{retentionModel}
		onpremDeployTargetConfigurationModel["copy_on_run_success"] = true
		onpremDeployTargetConfigurationModel["config_id"] = "testString"
		onpremDeployTargetConfigurationModel["backup_run_type"] = "Regular"
		onpremDeployTargetConfigurationModel["run_timeouts"] = []map[string]interface{}{cancellationTimeoutParamsModel}
		onpremDeployTargetConfigurationModel["log_retention"] = []map[string]interface{}{logRetentionModel}
		onpremDeployTargetConfigurationModel["params"] = []map[string]interface{}{onpremDeployParamsModel}

		rpaasTargetConfigurationModel := make(map[string]interface{})
		rpaasTargetConfigurationModel["schedule"] = []map[string]interface{}{targetScheduleModel}
		rpaasTargetConfigurationModel["retention"] = []map[string]interface{}{retentionModel}
		rpaasTargetConfigurationModel["copy_on_run_success"] = true
		rpaasTargetConfigurationModel["config_id"] = "testString"
		rpaasTargetConfigurationModel["backup_run_type"] = "Regular"
		rpaasTargetConfigurationModel["run_timeouts"] = []map[string]interface{}{cancellationTimeoutParamsModel}
		rpaasTargetConfigurationModel["log_retention"] = []map[string]interface{}{logRetentionModel}
		rpaasTargetConfigurationModel["target_id"] = int(26)
		rpaasTargetConfigurationModel["target_type"] = "Tape"

		model := make(map[string]interface{})
		model["replication_targets"] = []map[string]interface{}{replicationTargetConfigurationModel}
		model["archival_targets"] = []map[string]interface{}{archivalTargetConfigurationModel}
		model["cloud_spin_targets"] = []map[string]interface{}{cloudSpinTargetConfigurationModel}
		model["onprem_deploy_targets"] = []map[string]interface{}{onpremDeployTargetConfigurationModel}
		model["rpaas_targets"] = []map[string]interface{}{rpaasTargetConfigurationModel}

		assert.Equal(t, result, model)
	}

	targetScheduleModel := new(backuprecoveryv1.TargetSchedule)
	targetScheduleModel.Unit = core.StringPtr("Runs")
	targetScheduleModel.Frequency = core.Int64Ptr(int64(1))

	dataLockConfigModel := new(backuprecoveryv1.DataLockConfig)
	dataLockConfigModel.Mode = core.StringPtr("Compliance")
	dataLockConfigModel.Unit = core.StringPtr("Days")
	dataLockConfigModel.Duration = core.Int64Ptr(int64(1))
	dataLockConfigModel.EnableWormOnExternalTarget = core.BoolPtr(true)

	retentionModel := new(backuprecoveryv1.Retention)
	retentionModel.Unit = core.StringPtr("Days")
	retentionModel.Duration = core.Int64Ptr(int64(1))
	retentionModel.DataLockConfig = dataLockConfigModel

	cancellationTimeoutParamsModel := new(backuprecoveryv1.CancellationTimeoutParams)
	cancellationTimeoutParamsModel.TimeoutMins = core.Int64Ptr(int64(26))
	cancellationTimeoutParamsModel.BackupType = core.StringPtr("kRegular")

	logRetentionModel := new(backuprecoveryv1.LogRetention)
	logRetentionModel.Unit = core.StringPtr("Days")
	logRetentionModel.Duration = core.Int64Ptr(int64(0))
	logRetentionModel.DataLockConfig = dataLockConfigModel

	awsTargetConfigModel := new(backuprecoveryv1.AWSTargetConfig)
	awsTargetConfigModel.Region = core.Int64Ptr(int64(26))
	awsTargetConfigModel.SourceID = core.Int64Ptr(int64(26))

	azureTargetConfigModel := new(backuprecoveryv1.AzureTargetConfig)
	azureTargetConfigModel.ResourceGroup = core.Int64Ptr(int64(26))
	azureTargetConfigModel.SourceID = core.Int64Ptr(int64(26))

	remoteTargetConfigModel := new(backuprecoveryv1.RemoteTargetConfig)
	remoteTargetConfigModel.ClusterID = core.Int64Ptr(int64(26))

	replicationTargetConfigurationModel := new(backuprecoveryv1.ReplicationTargetConfiguration)
	replicationTargetConfigurationModel.Schedule = targetScheduleModel
	replicationTargetConfigurationModel.Retention = retentionModel
	replicationTargetConfigurationModel.CopyOnRunSuccess = core.BoolPtr(true)
	replicationTargetConfigurationModel.ConfigID = core.StringPtr("testString")
	replicationTargetConfigurationModel.BackupRunType = core.StringPtr("Regular")
	replicationTargetConfigurationModel.RunTimeouts = []backuprecoveryv1.CancellationTimeoutParams{*cancellationTimeoutParamsModel}
	replicationTargetConfigurationModel.LogRetention = logRetentionModel
	replicationTargetConfigurationModel.AwsTargetConfig = awsTargetConfigModel
	replicationTargetConfigurationModel.AzureTargetConfig = azureTargetConfigModel
	replicationTargetConfigurationModel.TargetType = core.StringPtr("RemoteCluster")
	replicationTargetConfigurationModel.RemoteTargetConfig = remoteTargetConfigModel

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

	tierLevelSettingsModel := new(backuprecoveryv1.TierLevelSettings)
	tierLevelSettingsModel.AwsTiering = awsTiersModel
	tierLevelSettingsModel.AzureTiering = azureTiersModel
	tierLevelSettingsModel.CloudPlatform = core.StringPtr("AWS")
	tierLevelSettingsModel.GoogleTiering = googleTiersModel
	tierLevelSettingsModel.OracleTiering = oracleTiersModel

	extendedRetentionScheduleModel := new(backuprecoveryv1.ExtendedRetentionSchedule)
	extendedRetentionScheduleModel.Unit = core.StringPtr("Runs")
	extendedRetentionScheduleModel.Frequency = core.Int64Ptr(int64(1))

	extendedRetentionPolicyModel := new(backuprecoveryv1.ExtendedRetentionPolicy)
	extendedRetentionPolicyModel.Schedule = extendedRetentionScheduleModel
	extendedRetentionPolicyModel.Retention = retentionModel
	extendedRetentionPolicyModel.RunType = core.StringPtr("Regular")
	extendedRetentionPolicyModel.ConfigID = core.StringPtr("testString")

	archivalTargetConfigurationModel := new(backuprecoveryv1.ArchivalTargetConfiguration)
	archivalTargetConfigurationModel.Schedule = targetScheduleModel
	archivalTargetConfigurationModel.Retention = retentionModel
	archivalTargetConfigurationModel.CopyOnRunSuccess = core.BoolPtr(true)
	archivalTargetConfigurationModel.ConfigID = core.StringPtr("testString")
	archivalTargetConfigurationModel.BackupRunType = core.StringPtr("Regular")
	archivalTargetConfigurationModel.RunTimeouts = []backuprecoveryv1.CancellationTimeoutParams{*cancellationTimeoutParamsModel}
	archivalTargetConfigurationModel.LogRetention = logRetentionModel
	archivalTargetConfigurationModel.TargetID = core.Int64Ptr(int64(26))
	archivalTargetConfigurationModel.TierSettings = tierLevelSettingsModel
	archivalTargetConfigurationModel.ExtendedRetention = []backuprecoveryv1.ExtendedRetentionPolicy{*extendedRetentionPolicyModel}

	customTagParamsModel := new(backuprecoveryv1.CustomTagParams)
	customTagParamsModel.Key = core.StringPtr("testString")
	customTagParamsModel.Value = core.StringPtr("testString")

	awsCloudSpinParamsModel := new(backuprecoveryv1.AwsCloudSpinParams)
	awsCloudSpinParamsModel.CustomTagList = []backuprecoveryv1.CustomTagParams{*customTagParamsModel}
	awsCloudSpinParamsModel.Region = core.Int64Ptr(int64(26))
	awsCloudSpinParamsModel.SubnetID = core.Int64Ptr(int64(26))
	awsCloudSpinParamsModel.VpcID = core.Int64Ptr(int64(26))

	azureCloudSpinParamsModel := new(backuprecoveryv1.AzureCloudSpinParams)
	azureCloudSpinParamsModel.AvailabilitySetID = core.Int64Ptr(int64(26))
	azureCloudSpinParamsModel.NetworkResourceGroupID = core.Int64Ptr(int64(26))
	azureCloudSpinParamsModel.ResourceGroupID = core.Int64Ptr(int64(26))
	azureCloudSpinParamsModel.StorageAccountID = core.Int64Ptr(int64(26))
	azureCloudSpinParamsModel.StorageContainerID = core.Int64Ptr(int64(26))
	azureCloudSpinParamsModel.StorageResourceGroupID = core.Int64Ptr(int64(26))
	azureCloudSpinParamsModel.TempVmResourceGroupID = core.Int64Ptr(int64(26))
	azureCloudSpinParamsModel.TempVmStorageAccountID = core.Int64Ptr(int64(26))
	azureCloudSpinParamsModel.TempVmStorageContainerID = core.Int64Ptr(int64(26))
	azureCloudSpinParamsModel.TempVmSubnetID = core.Int64Ptr(int64(26))
	azureCloudSpinParamsModel.TempVmVirtualNetworkID = core.Int64Ptr(int64(26))

	cloudSpinTargetModel := new(backuprecoveryv1.CloudSpinTarget)
	cloudSpinTargetModel.AwsParams = awsCloudSpinParamsModel
	cloudSpinTargetModel.AzureParams = azureCloudSpinParamsModel
	cloudSpinTargetModel.ID = core.Int64Ptr(int64(26))

	cloudSpinTargetConfigurationModel := new(backuprecoveryv1.CloudSpinTargetConfiguration)
	cloudSpinTargetConfigurationModel.Schedule = targetScheduleModel
	cloudSpinTargetConfigurationModel.Retention = retentionModel
	cloudSpinTargetConfigurationModel.CopyOnRunSuccess = core.BoolPtr(true)
	cloudSpinTargetConfigurationModel.ConfigID = core.StringPtr("testString")
	cloudSpinTargetConfigurationModel.BackupRunType = core.StringPtr("Regular")
	cloudSpinTargetConfigurationModel.RunTimeouts = []backuprecoveryv1.CancellationTimeoutParams{*cancellationTimeoutParamsModel}
	cloudSpinTargetConfigurationModel.LogRetention = logRetentionModel
	cloudSpinTargetConfigurationModel.Target = cloudSpinTargetModel

	onpremDeployParamsModel := new(backuprecoveryv1.OnpremDeployParams)
	onpremDeployParamsModel.ID = core.Int64Ptr(int64(26))

	onpremDeployTargetConfigurationModel := new(backuprecoveryv1.OnpremDeployTargetConfiguration)
	onpremDeployTargetConfigurationModel.Schedule = targetScheduleModel
	onpremDeployTargetConfigurationModel.Retention = retentionModel
	onpremDeployTargetConfigurationModel.CopyOnRunSuccess = core.BoolPtr(true)
	onpremDeployTargetConfigurationModel.ConfigID = core.StringPtr("testString")
	onpremDeployTargetConfigurationModel.BackupRunType = core.StringPtr("Regular")
	onpremDeployTargetConfigurationModel.RunTimeouts = []backuprecoveryv1.CancellationTimeoutParams{*cancellationTimeoutParamsModel}
	onpremDeployTargetConfigurationModel.LogRetention = logRetentionModel
	onpremDeployTargetConfigurationModel.Params = onpremDeployParamsModel

	rpaasTargetConfigurationModel := new(backuprecoveryv1.RpaasTargetConfiguration)
	rpaasTargetConfigurationModel.Schedule = targetScheduleModel
	rpaasTargetConfigurationModel.Retention = retentionModel
	rpaasTargetConfigurationModel.CopyOnRunSuccess = core.BoolPtr(true)
	rpaasTargetConfigurationModel.ConfigID = core.StringPtr("testString")
	rpaasTargetConfigurationModel.BackupRunType = core.StringPtr("Regular")
	rpaasTargetConfigurationModel.RunTimeouts = []backuprecoveryv1.CancellationTimeoutParams{*cancellationTimeoutParamsModel}
	rpaasTargetConfigurationModel.LogRetention = logRetentionModel
	rpaasTargetConfigurationModel.TargetID = core.Int64Ptr(int64(26))
	rpaasTargetConfigurationModel.TargetType = core.StringPtr("Tape")

	model := new(backuprecoveryv1.TargetsConfiguration)
	model.ReplicationTargets = []backuprecoveryv1.ReplicationTargetConfiguration{*replicationTargetConfigurationModel}
	model.ArchivalTargets = []backuprecoveryv1.ArchivalTargetConfiguration{*archivalTargetConfigurationModel}
	model.CloudSpinTargets = []backuprecoveryv1.CloudSpinTargetConfiguration{*cloudSpinTargetConfigurationModel}
	model.OnpremDeployTargets = []backuprecoveryv1.OnpremDeployTargetConfiguration{*onpremDeployTargetConfigurationModel}
	model.RpaasTargets = []backuprecoveryv1.RpaasTargetConfiguration{*rpaasTargetConfigurationModel}

	result, err := backuprecovery.ResourceIbmRestorePointsTargetsConfigurationToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmRestorePointsReplicationTargetConfigurationToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		targetScheduleModel := make(map[string]interface{})
		targetScheduleModel["unit"] = "Runs"
		targetScheduleModel["frequency"] = int(1)

		dataLockConfigModel := make(map[string]interface{})
		dataLockConfigModel["mode"] = "Compliance"
		dataLockConfigModel["unit"] = "Days"
		dataLockConfigModel["duration"] = int(1)
		dataLockConfigModel["enable_worm_on_external_target"] = true

		retentionModel := make(map[string]interface{})
		retentionModel["unit"] = "Days"
		retentionModel["duration"] = int(1)
		retentionModel["data_lock_config"] = []map[string]interface{}{dataLockConfigModel}

		cancellationTimeoutParamsModel := make(map[string]interface{})
		cancellationTimeoutParamsModel["timeout_mins"] = int(26)
		cancellationTimeoutParamsModel["backup_type"] = "kRegular"

		logRetentionModel := make(map[string]interface{})
		logRetentionModel["unit"] = "Days"
		logRetentionModel["duration"] = int(0)
		logRetentionModel["data_lock_config"] = []map[string]interface{}{dataLockConfigModel}

		awsTargetConfigModel := make(map[string]interface{})
		awsTargetConfigModel["region"] = int(26)
		awsTargetConfigModel["source_id"] = int(26)

		azureTargetConfigModel := make(map[string]interface{})
		azureTargetConfigModel["resource_group"] = int(26)
		azureTargetConfigModel["source_id"] = int(26)

		remoteTargetConfigModel := make(map[string]interface{})
		remoteTargetConfigModel["cluster_id"] = int(26)

		model := make(map[string]interface{})
		model["schedule"] = []map[string]interface{}{targetScheduleModel}
		model["retention"] = []map[string]interface{}{retentionModel}
		model["copy_on_run_success"] = true
		model["config_id"] = "testString"
		model["backup_run_type"] = "Regular"
		model["run_timeouts"] = []map[string]interface{}{cancellationTimeoutParamsModel}
		model["log_retention"] = []map[string]interface{}{logRetentionModel}
		model["aws_target_config"] = []map[string]interface{}{awsTargetConfigModel}
		model["azure_target_config"] = []map[string]interface{}{azureTargetConfigModel}
		model["target_type"] = "RemoteCluster"
		model["remote_target_config"] = []map[string]interface{}{remoteTargetConfigModel}

		assert.Equal(t, result, model)
	}

	targetScheduleModel := new(backuprecoveryv1.TargetSchedule)
	targetScheduleModel.Unit = core.StringPtr("Runs")
	targetScheduleModel.Frequency = core.Int64Ptr(int64(1))

	dataLockConfigModel := new(backuprecoveryv1.DataLockConfig)
	dataLockConfigModel.Mode = core.StringPtr("Compliance")
	dataLockConfigModel.Unit = core.StringPtr("Days")
	dataLockConfigModel.Duration = core.Int64Ptr(int64(1))
	dataLockConfigModel.EnableWormOnExternalTarget = core.BoolPtr(true)

	retentionModel := new(backuprecoveryv1.Retention)
	retentionModel.Unit = core.StringPtr("Days")
	retentionModel.Duration = core.Int64Ptr(int64(1))
	retentionModel.DataLockConfig = dataLockConfigModel

	cancellationTimeoutParamsModel := new(backuprecoveryv1.CancellationTimeoutParams)
	cancellationTimeoutParamsModel.TimeoutMins = core.Int64Ptr(int64(26))
	cancellationTimeoutParamsModel.BackupType = core.StringPtr("kRegular")

	logRetentionModel := new(backuprecoveryv1.LogRetention)
	logRetentionModel.Unit = core.StringPtr("Days")
	logRetentionModel.Duration = core.Int64Ptr(int64(0))
	logRetentionModel.DataLockConfig = dataLockConfigModel

	awsTargetConfigModel := new(backuprecoveryv1.AWSTargetConfig)
	awsTargetConfigModel.Region = core.Int64Ptr(int64(26))
	awsTargetConfigModel.SourceID = core.Int64Ptr(int64(26))

	azureTargetConfigModel := new(backuprecoveryv1.AzureTargetConfig)
	azureTargetConfigModel.ResourceGroup = core.Int64Ptr(int64(26))
	azureTargetConfigModel.SourceID = core.Int64Ptr(int64(26))

	remoteTargetConfigModel := new(backuprecoveryv1.RemoteTargetConfig)
	remoteTargetConfigModel.ClusterID = core.Int64Ptr(int64(26))

	model := new(backuprecoveryv1.ReplicationTargetConfiguration)
	model.Schedule = targetScheduleModel
	model.Retention = retentionModel
	model.CopyOnRunSuccess = core.BoolPtr(true)
	model.ConfigID = core.StringPtr("testString")
	model.BackupRunType = core.StringPtr("Regular")
	model.RunTimeouts = []backuprecoveryv1.CancellationTimeoutParams{*cancellationTimeoutParamsModel}
	model.LogRetention = logRetentionModel
	model.AwsTargetConfig = awsTargetConfigModel
	model.AzureTargetConfig = azureTargetConfigModel
	model.TargetType = core.StringPtr("RemoteCluster")
	model.RemoteTargetConfig = remoteTargetConfigModel

	result, err := backuprecovery.ResourceIbmRestorePointsReplicationTargetConfigurationToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmRestorePointsTargetScheduleToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["unit"] = "Runs"
		model["frequency"] = int(1)

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.TargetSchedule)
	model.Unit = core.StringPtr("Runs")
	model.Frequency = core.Int64Ptr(int64(1))

	result, err := backuprecovery.ResourceIbmRestorePointsTargetScheduleToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmRestorePointsRetentionToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		dataLockConfigModel := make(map[string]interface{})
		dataLockConfigModel["mode"] = "Compliance"
		dataLockConfigModel["unit"] = "Days"
		dataLockConfigModel["duration"] = int(1)
		dataLockConfigModel["enable_worm_on_external_target"] = true

		model := make(map[string]interface{})
		model["unit"] = "Days"
		model["duration"] = int(1)
		model["data_lock_config"] = []map[string]interface{}{dataLockConfigModel}

		assert.Equal(t, result, model)
	}

	dataLockConfigModel := new(backuprecoveryv1.DataLockConfig)
	dataLockConfigModel.Mode = core.StringPtr("Compliance")
	dataLockConfigModel.Unit = core.StringPtr("Days")
	dataLockConfigModel.Duration = core.Int64Ptr(int64(1))
	dataLockConfigModel.EnableWormOnExternalTarget = core.BoolPtr(true)

	model := new(backuprecoveryv1.Retention)
	model.Unit = core.StringPtr("Days")
	model.Duration = core.Int64Ptr(int64(1))
	model.DataLockConfig = dataLockConfigModel

	result, err := backuprecovery.ResourceIbmRestorePointsRetentionToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmRestorePointsDataLockConfigToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["mode"] = "Compliance"
		model["unit"] = "Days"
		model["duration"] = int(1)
		model["enable_worm_on_external_target"] = true

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.DataLockConfig)
	model.Mode = core.StringPtr("Compliance")
	model.Unit = core.StringPtr("Days")
	model.Duration = core.Int64Ptr(int64(1))
	model.EnableWormOnExternalTarget = core.BoolPtr(true)

	result, err := backuprecovery.ResourceIbmRestorePointsDataLockConfigToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmRestorePointsCancellationTimeoutParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["timeout_mins"] = int(26)
		model["backup_type"] = "kRegular"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.CancellationTimeoutParams)
	model.TimeoutMins = core.Int64Ptr(int64(26))
	model.BackupType = core.StringPtr("kRegular")

	result, err := backuprecovery.ResourceIbmRestorePointsCancellationTimeoutParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmRestorePointsLogRetentionToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		dataLockConfigModel := make(map[string]interface{})
		dataLockConfigModel["mode"] = "Compliance"
		dataLockConfigModel["unit"] = "Days"
		dataLockConfigModel["duration"] = int(1)
		dataLockConfigModel["enable_worm_on_external_target"] = true

		model := make(map[string]interface{})
		model["unit"] = "Days"
		model["duration"] = int(0)
		model["data_lock_config"] = []map[string]interface{}{dataLockConfigModel}

		assert.Equal(t, result, model)
	}

	dataLockConfigModel := new(backuprecoveryv1.DataLockConfig)
	dataLockConfigModel.Mode = core.StringPtr("Compliance")
	dataLockConfigModel.Unit = core.StringPtr("Days")
	dataLockConfigModel.Duration = core.Int64Ptr(int64(1))
	dataLockConfigModel.EnableWormOnExternalTarget = core.BoolPtr(true)

	model := new(backuprecoveryv1.LogRetention)
	model.Unit = core.StringPtr("Days")
	model.Duration = core.Int64Ptr(int64(0))
	model.DataLockConfig = dataLockConfigModel

	result, err := backuprecovery.ResourceIbmRestorePointsLogRetentionToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmRestorePointsAWSTargetConfigToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["name"] = "testString"
		model["region"] = int(26)
		model["region_name"] = "testString"
		model["source_id"] = int(26)

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.AWSTargetConfig)
	model.Name = core.StringPtr("testString")
	model.Region = core.Int64Ptr(int64(26))
	model.RegionName = core.StringPtr("testString")
	model.SourceID = core.Int64Ptr(int64(26))

	result, err := backuprecovery.ResourceIbmRestorePointsAWSTargetConfigToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmRestorePointsAzureTargetConfigToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["name"] = "testString"
		model["resource_group"] = int(26)
		model["resource_group_name"] = "testString"
		model["source_id"] = int(26)
		model["storage_account"] = int(38)
		model["storage_account_name"] = "testString"
		model["storage_container"] = int(38)
		model["storage_container_name"] = "testString"
		model["storage_resource_group"] = int(38)
		model["storage_resource_group_name"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.AzureTargetConfig)
	model.Name = core.StringPtr("testString")
	model.ResourceGroup = core.Int64Ptr(int64(26))
	model.ResourceGroupName = core.StringPtr("testString")
	model.SourceID = core.Int64Ptr(int64(26))
	model.StorageAccount = core.Int64Ptr(int64(38))
	model.StorageAccountName = core.StringPtr("testString")
	model.StorageContainer = core.Int64Ptr(int64(38))
	model.StorageContainerName = core.StringPtr("testString")
	model.StorageResourceGroup = core.Int64Ptr(int64(38))
	model.StorageResourceGroupName = core.StringPtr("testString")

	result, err := backuprecovery.ResourceIbmRestorePointsAzureTargetConfigToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmRestorePointsRemoteTargetConfigToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["cluster_id"] = int(26)
		model["cluster_name"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.RemoteTargetConfig)
	model.ClusterID = core.Int64Ptr(int64(26))
	model.ClusterName = core.StringPtr("testString")

	result, err := backuprecovery.ResourceIbmRestorePointsRemoteTargetConfigToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmRestorePointsArchivalTargetConfigurationToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		targetScheduleModel := make(map[string]interface{})
		targetScheduleModel["unit"] = "Runs"
		targetScheduleModel["frequency"] = int(1)

		dataLockConfigModel := make(map[string]interface{})
		dataLockConfigModel["mode"] = "Compliance"
		dataLockConfigModel["unit"] = "Days"
		dataLockConfigModel["duration"] = int(1)
		dataLockConfigModel["enable_worm_on_external_target"] = true

		retentionModel := make(map[string]interface{})
		retentionModel["unit"] = "Days"
		retentionModel["duration"] = int(1)
		retentionModel["data_lock_config"] = []map[string]interface{}{dataLockConfigModel}

		cancellationTimeoutParamsModel := make(map[string]interface{})
		cancellationTimeoutParamsModel["timeout_mins"] = int(26)
		cancellationTimeoutParamsModel["backup_type"] = "kRegular"

		logRetentionModel := make(map[string]interface{})
		logRetentionModel["unit"] = "Days"
		logRetentionModel["duration"] = int(0)
		logRetentionModel["data_lock_config"] = []map[string]interface{}{dataLockConfigModel}

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

		tierLevelSettingsModel := make(map[string]interface{})
		tierLevelSettingsModel["aws_tiering"] = []map[string]interface{}{awsTiersModel}
		tierLevelSettingsModel["azure_tiering"] = []map[string]interface{}{azureTiersModel}
		tierLevelSettingsModel["cloud_platform"] = "AWS"
		tierLevelSettingsModel["google_tiering"] = []map[string]interface{}{googleTiersModel}
		tierLevelSettingsModel["oracle_tiering"] = []map[string]interface{}{oracleTiersModel}

		extendedRetentionScheduleModel := make(map[string]interface{})
		extendedRetentionScheduleModel["unit"] = "Runs"
		extendedRetentionScheduleModel["frequency"] = int(1)

		extendedRetentionPolicyModel := make(map[string]interface{})
		extendedRetentionPolicyModel["schedule"] = []map[string]interface{}{extendedRetentionScheduleModel}
		extendedRetentionPolicyModel["retention"] = []map[string]interface{}{retentionModel}
		extendedRetentionPolicyModel["run_type"] = "Regular"
		extendedRetentionPolicyModel["config_id"] = "testString"

		model := make(map[string]interface{})
		model["schedule"] = []map[string]interface{}{targetScheduleModel}
		model["retention"] = []map[string]interface{}{retentionModel}
		model["copy_on_run_success"] = true
		model["config_id"] = "testString"
		model["backup_run_type"] = "Regular"
		model["run_timeouts"] = []map[string]interface{}{cancellationTimeoutParamsModel}
		model["log_retention"] = []map[string]interface{}{logRetentionModel}
		model["target_id"] = int(26)
		model["target_name"] = "testString"
		model["target_type"] = "Tape"
		model["tier_settings"] = []map[string]interface{}{tierLevelSettingsModel}
		model["extended_retention"] = []map[string]interface{}{extendedRetentionPolicyModel}

		assert.Equal(t, result, model)
	}

	targetScheduleModel := new(backuprecoveryv1.TargetSchedule)
	targetScheduleModel.Unit = core.StringPtr("Runs")
	targetScheduleModel.Frequency = core.Int64Ptr(int64(1))

	dataLockConfigModel := new(backuprecoveryv1.DataLockConfig)
	dataLockConfigModel.Mode = core.StringPtr("Compliance")
	dataLockConfigModel.Unit = core.StringPtr("Days")
	dataLockConfigModel.Duration = core.Int64Ptr(int64(1))
	dataLockConfigModel.EnableWormOnExternalTarget = core.BoolPtr(true)

	retentionModel := new(backuprecoveryv1.Retention)
	retentionModel.Unit = core.StringPtr("Days")
	retentionModel.Duration = core.Int64Ptr(int64(1))
	retentionModel.DataLockConfig = dataLockConfigModel

	cancellationTimeoutParamsModel := new(backuprecoveryv1.CancellationTimeoutParams)
	cancellationTimeoutParamsModel.TimeoutMins = core.Int64Ptr(int64(26))
	cancellationTimeoutParamsModel.BackupType = core.StringPtr("kRegular")

	logRetentionModel := new(backuprecoveryv1.LogRetention)
	logRetentionModel.Unit = core.StringPtr("Days")
	logRetentionModel.Duration = core.Int64Ptr(int64(0))
	logRetentionModel.DataLockConfig = dataLockConfigModel

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

	tierLevelSettingsModel := new(backuprecoveryv1.TierLevelSettings)
	tierLevelSettingsModel.AwsTiering = awsTiersModel
	tierLevelSettingsModel.AzureTiering = azureTiersModel
	tierLevelSettingsModel.CloudPlatform = core.StringPtr("AWS")
	tierLevelSettingsModel.GoogleTiering = googleTiersModel
	tierLevelSettingsModel.OracleTiering = oracleTiersModel

	extendedRetentionScheduleModel := new(backuprecoveryv1.ExtendedRetentionSchedule)
	extendedRetentionScheduleModel.Unit = core.StringPtr("Runs")
	extendedRetentionScheduleModel.Frequency = core.Int64Ptr(int64(1))

	extendedRetentionPolicyModel := new(backuprecoveryv1.ExtendedRetentionPolicy)
	extendedRetentionPolicyModel.Schedule = extendedRetentionScheduleModel
	extendedRetentionPolicyModel.Retention = retentionModel
	extendedRetentionPolicyModel.RunType = core.StringPtr("Regular")
	extendedRetentionPolicyModel.ConfigID = core.StringPtr("testString")

	model := new(backuprecoveryv1.ArchivalTargetConfiguration)
	model.Schedule = targetScheduleModel
	model.Retention = retentionModel
	model.CopyOnRunSuccess = core.BoolPtr(true)
	model.ConfigID = core.StringPtr("testString")
	model.BackupRunType = core.StringPtr("Regular")
	model.RunTimeouts = []backuprecoveryv1.CancellationTimeoutParams{*cancellationTimeoutParamsModel}
	model.LogRetention = logRetentionModel
	model.TargetID = core.Int64Ptr(int64(26))
	model.TargetName = core.StringPtr("testString")
	model.TargetType = core.StringPtr("Tape")
	model.TierSettings = tierLevelSettingsModel
	model.ExtendedRetention = []backuprecoveryv1.ExtendedRetentionPolicy{*extendedRetentionPolicyModel}

	result, err := backuprecovery.ResourceIbmRestorePointsArchivalTargetConfigurationToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmRestorePointsTierLevelSettingsToMap(t *testing.T) {
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

	model := new(backuprecoveryv1.TierLevelSettings)
	model.AwsTiering = awsTiersModel
	model.AzureTiering = azureTiersModel
	model.CloudPlatform = core.StringPtr("AWS")
	model.GoogleTiering = googleTiersModel
	model.OracleTiering = oracleTiersModel

	result, err := backuprecovery.ResourceIbmRestorePointsTierLevelSettingsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmRestorePointsExtendedRetentionPolicyToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		extendedRetentionScheduleModel := make(map[string]interface{})
		extendedRetentionScheduleModel["unit"] = "Runs"
		extendedRetentionScheduleModel["frequency"] = int(1)

		dataLockConfigModel := make(map[string]interface{})
		dataLockConfigModel["mode"] = "Compliance"
		dataLockConfigModel["unit"] = "Days"
		dataLockConfigModel["duration"] = int(1)
		dataLockConfigModel["enable_worm_on_external_target"] = true

		retentionModel := make(map[string]interface{})
		retentionModel["unit"] = "Days"
		retentionModel["duration"] = int(1)
		retentionModel["data_lock_config"] = []map[string]interface{}{dataLockConfigModel}

		model := make(map[string]interface{})
		model["schedule"] = []map[string]interface{}{extendedRetentionScheduleModel}
		model["retention"] = []map[string]interface{}{retentionModel}
		model["run_type"] = "Regular"
		model["config_id"] = "testString"

		assert.Equal(t, result, model)
	}

	extendedRetentionScheduleModel := new(backuprecoveryv1.ExtendedRetentionSchedule)
	extendedRetentionScheduleModel.Unit = core.StringPtr("Runs")
	extendedRetentionScheduleModel.Frequency = core.Int64Ptr(int64(1))

	dataLockConfigModel := new(backuprecoveryv1.DataLockConfig)
	dataLockConfigModel.Mode = core.StringPtr("Compliance")
	dataLockConfigModel.Unit = core.StringPtr("Days")
	dataLockConfigModel.Duration = core.Int64Ptr(int64(1))
	dataLockConfigModel.EnableWormOnExternalTarget = core.BoolPtr(true)

	retentionModel := new(backuprecoveryv1.Retention)
	retentionModel.Unit = core.StringPtr("Days")
	retentionModel.Duration = core.Int64Ptr(int64(1))
	retentionModel.DataLockConfig = dataLockConfigModel

	model := new(backuprecoveryv1.ExtendedRetentionPolicy)
	model.Schedule = extendedRetentionScheduleModel
	model.Retention = retentionModel
	model.RunType = core.StringPtr("Regular")
	model.ConfigID = core.StringPtr("testString")

	result, err := backuprecovery.ResourceIbmRestorePointsExtendedRetentionPolicyToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmRestorePointsExtendedRetentionScheduleToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["unit"] = "Runs"
		model["frequency"] = int(1)

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.ExtendedRetentionSchedule)
	model.Unit = core.StringPtr("Runs")
	model.Frequency = core.Int64Ptr(int64(1))

	result, err := backuprecovery.ResourceIbmRestorePointsExtendedRetentionScheduleToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmRestorePointsCloudSpinTargetConfigurationToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		targetScheduleModel := make(map[string]interface{})
		targetScheduleModel["unit"] = "Runs"
		targetScheduleModel["frequency"] = int(1)

		dataLockConfigModel := make(map[string]interface{})
		dataLockConfigModel["mode"] = "Compliance"
		dataLockConfigModel["unit"] = "Days"
		dataLockConfigModel["duration"] = int(1)
		dataLockConfigModel["enable_worm_on_external_target"] = true

		retentionModel := make(map[string]interface{})
		retentionModel["unit"] = "Days"
		retentionModel["duration"] = int(1)
		retentionModel["data_lock_config"] = []map[string]interface{}{dataLockConfigModel}

		cancellationTimeoutParamsModel := make(map[string]interface{})
		cancellationTimeoutParamsModel["timeout_mins"] = int(26)
		cancellationTimeoutParamsModel["backup_type"] = "kRegular"

		logRetentionModel := make(map[string]interface{})
		logRetentionModel["unit"] = "Days"
		logRetentionModel["duration"] = int(0)
		logRetentionModel["data_lock_config"] = []map[string]interface{}{dataLockConfigModel}

		customTagParamsModel := make(map[string]interface{})
		customTagParamsModel["key"] = "testString"
		customTagParamsModel["value"] = "testString"

		awsCloudSpinParamsModel := make(map[string]interface{})
		awsCloudSpinParamsModel["custom_tag_list"] = []map[string]interface{}{customTagParamsModel}
		awsCloudSpinParamsModel["region"] = int(26)
		awsCloudSpinParamsModel["subnet_id"] = int(26)
		awsCloudSpinParamsModel["vpc_id"] = int(26)

		azureCloudSpinParamsModel := make(map[string]interface{})
		azureCloudSpinParamsModel["availability_set_id"] = int(26)
		azureCloudSpinParamsModel["network_resource_group_id"] = int(26)
		azureCloudSpinParamsModel["resource_group_id"] = int(26)
		azureCloudSpinParamsModel["storage_account_id"] = int(26)
		azureCloudSpinParamsModel["storage_container_id"] = int(26)
		azureCloudSpinParamsModel["storage_resource_group_id"] = int(26)
		azureCloudSpinParamsModel["temp_vm_resource_group_id"] = int(26)
		azureCloudSpinParamsModel["temp_vm_storage_account_id"] = int(26)
		azureCloudSpinParamsModel["temp_vm_storage_container_id"] = int(26)
		azureCloudSpinParamsModel["temp_vm_subnet_id"] = int(26)
		azureCloudSpinParamsModel["temp_vm_virtual_network_id"] = int(26)

		cloudSpinTargetModel := make(map[string]interface{})
		cloudSpinTargetModel["aws_params"] = []map[string]interface{}{awsCloudSpinParamsModel}
		cloudSpinTargetModel["azure_params"] = []map[string]interface{}{azureCloudSpinParamsModel}
		cloudSpinTargetModel["id"] = int(26)

		model := make(map[string]interface{})
		model["schedule"] = []map[string]interface{}{targetScheduleModel}
		model["retention"] = []map[string]interface{}{retentionModel}
		model["copy_on_run_success"] = true
		model["config_id"] = "testString"
		model["backup_run_type"] = "Regular"
		model["run_timeouts"] = []map[string]interface{}{cancellationTimeoutParamsModel}
		model["log_retention"] = []map[string]interface{}{logRetentionModel}
		model["target"] = []map[string]interface{}{cloudSpinTargetModel}

		assert.Equal(t, result, model)
	}

	targetScheduleModel := new(backuprecoveryv1.TargetSchedule)
	targetScheduleModel.Unit = core.StringPtr("Runs")
	targetScheduleModel.Frequency = core.Int64Ptr(int64(1))

	dataLockConfigModel := new(backuprecoveryv1.DataLockConfig)
	dataLockConfigModel.Mode = core.StringPtr("Compliance")
	dataLockConfigModel.Unit = core.StringPtr("Days")
	dataLockConfigModel.Duration = core.Int64Ptr(int64(1))
	dataLockConfigModel.EnableWormOnExternalTarget = core.BoolPtr(true)

	retentionModel := new(backuprecoveryv1.Retention)
	retentionModel.Unit = core.StringPtr("Days")
	retentionModel.Duration = core.Int64Ptr(int64(1))
	retentionModel.DataLockConfig = dataLockConfigModel

	cancellationTimeoutParamsModel := new(backuprecoveryv1.CancellationTimeoutParams)
	cancellationTimeoutParamsModel.TimeoutMins = core.Int64Ptr(int64(26))
	cancellationTimeoutParamsModel.BackupType = core.StringPtr("kRegular")

	logRetentionModel := new(backuprecoveryv1.LogRetention)
	logRetentionModel.Unit = core.StringPtr("Days")
	logRetentionModel.Duration = core.Int64Ptr(int64(0))
	logRetentionModel.DataLockConfig = dataLockConfigModel

	customTagParamsModel := new(backuprecoveryv1.CustomTagParams)
	customTagParamsModel.Key = core.StringPtr("testString")
	customTagParamsModel.Value = core.StringPtr("testString")

	awsCloudSpinParamsModel := new(backuprecoveryv1.AwsCloudSpinParams)
	awsCloudSpinParamsModel.CustomTagList = []backuprecoveryv1.CustomTagParams{*customTagParamsModel}
	awsCloudSpinParamsModel.Region = core.Int64Ptr(int64(26))
	awsCloudSpinParamsModel.SubnetID = core.Int64Ptr(int64(26))
	awsCloudSpinParamsModel.VpcID = core.Int64Ptr(int64(26))

	azureCloudSpinParamsModel := new(backuprecoveryv1.AzureCloudSpinParams)
	azureCloudSpinParamsModel.AvailabilitySetID = core.Int64Ptr(int64(26))
	azureCloudSpinParamsModel.NetworkResourceGroupID = core.Int64Ptr(int64(26))
	azureCloudSpinParamsModel.ResourceGroupID = core.Int64Ptr(int64(26))
	azureCloudSpinParamsModel.StorageAccountID = core.Int64Ptr(int64(26))
	azureCloudSpinParamsModel.StorageContainerID = core.Int64Ptr(int64(26))
	azureCloudSpinParamsModel.StorageResourceGroupID = core.Int64Ptr(int64(26))
	azureCloudSpinParamsModel.TempVmResourceGroupID = core.Int64Ptr(int64(26))
	azureCloudSpinParamsModel.TempVmStorageAccountID = core.Int64Ptr(int64(26))
	azureCloudSpinParamsModel.TempVmStorageContainerID = core.Int64Ptr(int64(26))
	azureCloudSpinParamsModel.TempVmSubnetID = core.Int64Ptr(int64(26))
	azureCloudSpinParamsModel.TempVmVirtualNetworkID = core.Int64Ptr(int64(26))

	cloudSpinTargetModel := new(backuprecoveryv1.CloudSpinTarget)
	cloudSpinTargetModel.AwsParams = awsCloudSpinParamsModel
	cloudSpinTargetModel.AzureParams = azureCloudSpinParamsModel
	cloudSpinTargetModel.ID = core.Int64Ptr(int64(26))

	model := new(backuprecoveryv1.CloudSpinTargetConfiguration)
	model.Schedule = targetScheduleModel
	model.Retention = retentionModel
	model.CopyOnRunSuccess = core.BoolPtr(true)
	model.ConfigID = core.StringPtr("testString")
	model.BackupRunType = core.StringPtr("Regular")
	model.RunTimeouts = []backuprecoveryv1.CancellationTimeoutParams{*cancellationTimeoutParamsModel}
	model.LogRetention = logRetentionModel
	model.Target = cloudSpinTargetModel

	result, err := backuprecovery.ResourceIbmRestorePointsCloudSpinTargetConfigurationToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmRestorePointsOnpremDeployTargetConfigurationToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		targetScheduleModel := make(map[string]interface{})
		targetScheduleModel["unit"] = "Runs"
		targetScheduleModel["frequency"] = int(1)

		dataLockConfigModel := make(map[string]interface{})
		dataLockConfigModel["mode"] = "Compliance"
		dataLockConfigModel["unit"] = "Days"
		dataLockConfigModel["duration"] = int(1)
		dataLockConfigModel["enable_worm_on_external_target"] = true

		retentionModel := make(map[string]interface{})
		retentionModel["unit"] = "Days"
		retentionModel["duration"] = int(1)
		retentionModel["data_lock_config"] = []map[string]interface{}{dataLockConfigModel}

		cancellationTimeoutParamsModel := make(map[string]interface{})
		cancellationTimeoutParamsModel["timeout_mins"] = int(26)
		cancellationTimeoutParamsModel["backup_type"] = "kRegular"

		logRetentionModel := make(map[string]interface{})
		logRetentionModel["unit"] = "Days"
		logRetentionModel["duration"] = int(0)
		logRetentionModel["data_lock_config"] = []map[string]interface{}{dataLockConfigModel}

		onpremDeployParamsModel := make(map[string]interface{})
		onpremDeployParamsModel["id"] = int(26)

		model := make(map[string]interface{})
		model["schedule"] = []map[string]interface{}{targetScheduleModel}
		model["retention"] = []map[string]interface{}{retentionModel}
		model["copy_on_run_success"] = true
		model["config_id"] = "testString"
		model["backup_run_type"] = "Regular"
		model["run_timeouts"] = []map[string]interface{}{cancellationTimeoutParamsModel}
		model["log_retention"] = []map[string]interface{}{logRetentionModel}
		model["params"] = []map[string]interface{}{onpremDeployParamsModel}

		assert.Equal(t, result, model)
	}

	targetScheduleModel := new(backuprecoveryv1.TargetSchedule)
	targetScheduleModel.Unit = core.StringPtr("Runs")
	targetScheduleModel.Frequency = core.Int64Ptr(int64(1))

	dataLockConfigModel := new(backuprecoveryv1.DataLockConfig)
	dataLockConfigModel.Mode = core.StringPtr("Compliance")
	dataLockConfigModel.Unit = core.StringPtr("Days")
	dataLockConfigModel.Duration = core.Int64Ptr(int64(1))
	dataLockConfigModel.EnableWormOnExternalTarget = core.BoolPtr(true)

	retentionModel := new(backuprecoveryv1.Retention)
	retentionModel.Unit = core.StringPtr("Days")
	retentionModel.Duration = core.Int64Ptr(int64(1))
	retentionModel.DataLockConfig = dataLockConfigModel

	cancellationTimeoutParamsModel := new(backuprecoveryv1.CancellationTimeoutParams)
	cancellationTimeoutParamsModel.TimeoutMins = core.Int64Ptr(int64(26))
	cancellationTimeoutParamsModel.BackupType = core.StringPtr("kRegular")

	logRetentionModel := new(backuprecoveryv1.LogRetention)
	logRetentionModel.Unit = core.StringPtr("Days")
	logRetentionModel.Duration = core.Int64Ptr(int64(0))
	logRetentionModel.DataLockConfig = dataLockConfigModel

	onpremDeployParamsModel := new(backuprecoveryv1.OnpremDeployParams)
	onpremDeployParamsModel.ID = core.Int64Ptr(int64(26))

	model := new(backuprecoveryv1.OnpremDeployTargetConfiguration)
	model.Schedule = targetScheduleModel
	model.Retention = retentionModel
	model.CopyOnRunSuccess = core.BoolPtr(true)
	model.ConfigID = core.StringPtr("testString")
	model.BackupRunType = core.StringPtr("Regular")
	model.RunTimeouts = []backuprecoveryv1.CancellationTimeoutParams{*cancellationTimeoutParamsModel}
	model.LogRetention = logRetentionModel
	model.Params = onpremDeployParamsModel

	result, err := backuprecovery.ResourceIbmRestorePointsOnpremDeployTargetConfigurationToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmRestorePointsOnpremDeployParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["id"] = int(26)

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.OnpremDeployParams)
	model.ID = core.Int64Ptr(int64(26))

	result, err := backuprecovery.ResourceIbmRestorePointsOnpremDeployParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmRestorePointsRpaasTargetConfigurationToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		targetScheduleModel := make(map[string]interface{})
		targetScheduleModel["unit"] = "Runs"
		targetScheduleModel["frequency"] = int(1)

		dataLockConfigModel := make(map[string]interface{})
		dataLockConfigModel["mode"] = "Compliance"
		dataLockConfigModel["unit"] = "Days"
		dataLockConfigModel["duration"] = int(1)
		dataLockConfigModel["enable_worm_on_external_target"] = true

		retentionModel := make(map[string]interface{})
		retentionModel["unit"] = "Days"
		retentionModel["duration"] = int(1)
		retentionModel["data_lock_config"] = []map[string]interface{}{dataLockConfigModel}

		cancellationTimeoutParamsModel := make(map[string]interface{})
		cancellationTimeoutParamsModel["timeout_mins"] = int(26)
		cancellationTimeoutParamsModel["backup_type"] = "kRegular"

		logRetentionModel := make(map[string]interface{})
		logRetentionModel["unit"] = "Days"
		logRetentionModel["duration"] = int(0)
		logRetentionModel["data_lock_config"] = []map[string]interface{}{dataLockConfigModel}

		model := make(map[string]interface{})
		model["schedule"] = []map[string]interface{}{targetScheduleModel}
		model["retention"] = []map[string]interface{}{retentionModel}
		model["copy_on_run_success"] = true
		model["config_id"] = "testString"
		model["backup_run_type"] = "Regular"
		model["run_timeouts"] = []map[string]interface{}{cancellationTimeoutParamsModel}
		model["log_retention"] = []map[string]interface{}{logRetentionModel}
		model["target_id"] = int(26)
		model["target_name"] = "testString"
		model["target_type"] = "Tape"

		assert.Equal(t, result, model)
	}

	targetScheduleModel := new(backuprecoveryv1.TargetSchedule)
	targetScheduleModel.Unit = core.StringPtr("Runs")
	targetScheduleModel.Frequency = core.Int64Ptr(int64(1))

	dataLockConfigModel := new(backuprecoveryv1.DataLockConfig)
	dataLockConfigModel.Mode = core.StringPtr("Compliance")
	dataLockConfigModel.Unit = core.StringPtr("Days")
	dataLockConfigModel.Duration = core.Int64Ptr(int64(1))
	dataLockConfigModel.EnableWormOnExternalTarget = core.BoolPtr(true)

	retentionModel := new(backuprecoveryv1.Retention)
	retentionModel.Unit = core.StringPtr("Days")
	retentionModel.Duration = core.Int64Ptr(int64(1))
	retentionModel.DataLockConfig = dataLockConfigModel

	cancellationTimeoutParamsModel := new(backuprecoveryv1.CancellationTimeoutParams)
	cancellationTimeoutParamsModel.TimeoutMins = core.Int64Ptr(int64(26))
	cancellationTimeoutParamsModel.BackupType = core.StringPtr("kRegular")

	logRetentionModel := new(backuprecoveryv1.LogRetention)
	logRetentionModel.Unit = core.StringPtr("Days")
	logRetentionModel.Duration = core.Int64Ptr(int64(0))
	logRetentionModel.DataLockConfig = dataLockConfigModel

	model := new(backuprecoveryv1.RpaasTargetConfiguration)
	model.Schedule = targetScheduleModel
	model.Retention = retentionModel
	model.CopyOnRunSuccess = core.BoolPtr(true)
	model.ConfigID = core.StringPtr("testString")
	model.BackupRunType = core.StringPtr("Regular")
	model.RunTimeouts = []backuprecoveryv1.CancellationTimeoutParams{*cancellationTimeoutParamsModel}
	model.LogRetention = logRetentionModel
	model.TargetID = core.Int64Ptr(int64(26))
	model.TargetName = core.StringPtr("testString")
	model.TargetType = core.StringPtr("Tape")

	result, err := backuprecovery.ResourceIbmRestorePointsRpaasTargetConfigurationToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmRestorePointsTimeRangeInfoToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		recoveryTimeRangeInfoModel := make(map[string]interface{})
		recoveryTimeRangeInfoModel["end_time_usecs"] = int(26)
		recoveryTimeRangeInfoModel["protection_group_id"] = "testString"
		recoveryTimeRangeInfoModel["start_time_usecs"] = int(26)

		model := make(map[string]interface{})
		model["error_message"] = "testString"
		model["time_ranges"] = []map[string]interface{}{recoveryTimeRangeInfoModel}
		model["user_message"] = "testString"

		assert.Equal(t, result, model)
	}

	recoveryTimeRangeInfoModel := new(backuprecoveryv1.RecoveryTimeRangeInfo)
	recoveryTimeRangeInfoModel.EndTimeUsecs = core.Int64Ptr(int64(26))
	recoveryTimeRangeInfoModel.ProtectionGroupID = core.StringPtr("testString")
	recoveryTimeRangeInfoModel.StartTimeUsecs = core.Int64Ptr(int64(26))

	model := new(backuprecoveryv1.TimeRangeInfo)
	model.ErrorMessage = core.StringPtr("testString")
	model.TimeRanges = []backuprecoveryv1.RecoveryTimeRangeInfo{*recoveryTimeRangeInfoModel}
	model.UserMessage = core.StringPtr("testString")

	result, err := backuprecovery.ResourceIbmRestorePointsTimeRangeInfoToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmRestorePointsRecoveryTimeRangeInfoToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["end_time_usecs"] = int(26)
		model["protection_group_id"] = "testString"
		model["start_time_usecs"] = int(26)

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.RecoveryTimeRangeInfo)
	model.EndTimeUsecs = core.Int64Ptr(int64(26))
	model.ProtectionGroupID = core.StringPtr("testString")
	model.StartTimeUsecs = core.Int64Ptr(int64(26))

	result, err := backuprecovery.ResourceIbmRestorePointsRecoveryTimeRangeInfoToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
