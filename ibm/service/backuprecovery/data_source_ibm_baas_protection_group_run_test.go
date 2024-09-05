// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.94.0-fa797aec-20240814-142622
 */

package backuprecovery_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/backuprecovery"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/stretchr/testify/assert"
	"github.ibm.com/BackupAndRecovery/ibm-backup-recovery-sdk-go/backuprecoveryv1"
)

func TestAccIbmBaasProtectionGroupRunDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmBaasProtectionGroupRunDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_baas_protection_group_run.baas_protection_group_run_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_protection_group_run.baas_protection_group_run_instance", "baas_protection_group_run_id"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_protection_group_run.baas_protection_group_run_instance", "tenant_id"),
				),
			},
		},
	})
}

func testAccCheckIbmBaasProtectionGroupRunDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_baas_protection_group_run" "baas_protection_group_run_instance" {
			id = "id"
			tenantId = 1
			requestInitiatorType = "UIUser"
			runId = "runId"
			startTimeUsecs = 1
			endTimeUsecs = 1
			runTypes = [ "kAll" ]
			includeObjectDetails = true
			localBackupRunStatus = [ "Accepted" ]
			replicationRunStatus = [ "Accepted" ]
			archivalRunStatus = [ "Accepted" ]
			cloudSpinRunStatus = [ "Accepted" ]
			numRuns = 1
			excludeNonRestorableRuns = true
			runTags = [ "runTags" ]
			useCachedData = true
			filterByEndTime = true
			snapshotTargetTypes = [ "Local" ]
			onlyReturnSuccessfulCopyRun = true
			filterByCopyTaskEndTime = true
		}
	`)
}

func TestDataSourceIbmBaasProtectionGroupRunClusterIdentifierToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["cluster_id"] = int(26)
		model["cluster_incarnation_id"] = int(26)
		model["cluster_name"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.ClusterIdentifier)
	model.ClusterID = core.Int64Ptr(int64(26))
	model.ClusterIncarnationID = core.Int64Ptr(int64(26))
	model.ClusterName = core.StringPtr("testString")

	result, err := backuprecovery.DataSourceIbmBaasProtectionGroupRunClusterIdentifierToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasProtectionGroupRunObjectRunResultToMap(t *testing.T) {
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

		backupDataStatsModel := make(map[string]interface{})
		backupDataStatsModel["logical_size_bytes"] = int(26)
		backupDataStatsModel["bytes_written"] = int(26)
		backupDataStatsModel["bytes_read"] = int(26)

		dataLockConstraintsModel := make(map[string]interface{})
		dataLockConstraintsModel["mode"] = "Compliance"
		dataLockConstraintsModel["expiry_time_usecs"] = int(26)

		snapshotInfoModel := make(map[string]interface{})
		snapshotInfoModel["snapshot_id"] = "testString"
		snapshotInfoModel["status"] = "kInProgress"
		snapshotInfoModel["status_message"] = "testString"
		snapshotInfoModel["start_time_usecs"] = int(26)
		snapshotInfoModel["end_time_usecs"] = int(26)
		snapshotInfoModel["admitted_time_usecs"] = int(26)
		snapshotInfoModel["permit_grant_time_usecs"] = int(26)
		snapshotInfoModel["queue_duration_usecs"] = int(26)
		snapshotInfoModel["snapshot_creation_time_usecs"] = int(26)
		snapshotInfoModel["stats"] = []map[string]interface{}{backupDataStatsModel}
		snapshotInfoModel["progress_task_id"] = "testString"
		snapshotInfoModel["indexing_task_id"] = "testString"
		snapshotInfoModel["stats_task_id"] = "testString"
		snapshotInfoModel["warnings"] = []string{"testString"}
		snapshotInfoModel["is_manually_deleted"] = true
		snapshotInfoModel["expiry_time_usecs"] = int(26)
		snapshotInfoModel["total_file_count"] = int(26)
		snapshotInfoModel["backup_file_count"] = int(26)
		snapshotInfoModel["data_lock_constraints"] = []map[string]interface{}{dataLockConstraintsModel}

		backupAttemptModel := make(map[string]interface{})
		backupAttemptModel["start_time_usecs"] = int(26)
		backupAttemptModel["end_time_usecs"] = int(26)
		backupAttemptModel["admitted_time_usecs"] = int(26)
		backupAttemptModel["permit_grant_time_usecs"] = int(26)
		backupAttemptModel["queue_duration_usecs"] = int(26)
		backupAttemptModel["snapshot_creation_time_usecs"] = int(26)
		backupAttemptModel["status"] = "Accepted"
		backupAttemptModel["stats"] = []map[string]interface{}{backupDataStatsModel}
		backupAttemptModel["progress_task_id"] = "testString"
		backupAttemptModel["message"] = "testString"

		backupRunModel := make(map[string]interface{})
		backupRunModel["snapshot_info"] = []map[string]interface{}{snapshotInfoModel}
		backupRunModel["failed_attempts"] = []map[string]interface{}{backupAttemptModel}

		awsTargetConfigModel := make(map[string]interface{})
		awsTargetConfigModel["region"] = int(26)
		awsTargetConfigModel["source_id"] = int(26)

		azureTargetConfigModel := make(map[string]interface{})
		azureTargetConfigModel["resource_group"] = int(26)
		azureTargetConfigModel["source_id"] = int(26)

		replicationDataStatsModel := make(map[string]interface{})
		replicationDataStatsModel["logical_size_bytes"] = int(26)
		replicationDataStatsModel["logical_bytes_transferred"] = int(26)
		replicationDataStatsModel["physical_bytes_transferred"] = int(26)

		replicationTargetResultModel := make(map[string]interface{})
		replicationTargetResultModel["cluster_id"] = int(26)
		replicationTargetResultModel["cluster_incarnation_id"] = int(26)
		replicationTargetResultModel["aws_target_config"] = []map[string]interface{}{awsTargetConfigModel}
		replicationTargetResultModel["azure_target_config"] = []map[string]interface{}{azureTargetConfigModel}
		replicationTargetResultModel["start_time_usecs"] = int(26)
		replicationTargetResultModel["end_time_usecs"] = int(26)
		replicationTargetResultModel["queued_time_usecs"] = int(26)
		replicationTargetResultModel["status"] = "Accepted"
		replicationTargetResultModel["message"] = "testString"
		replicationTargetResultModel["percentage_completed"] = int(38)
		replicationTargetResultModel["stats"] = []map[string]interface{}{replicationDataStatsModel}
		replicationTargetResultModel["is_manually_deleted"] = true
		replicationTargetResultModel["expiry_time_usecs"] = int(26)
		replicationTargetResultModel["replication_task_id"] = "testString"
		replicationTargetResultModel["entries_changed"] = int(26)
		replicationTargetResultModel["is_in_bound"] = true
		replicationTargetResultModel["data_lock_constraints"] = []map[string]interface{}{dataLockConstraintsModel}
		replicationTargetResultModel["on_legal_hold"] = true
		replicationTargetResultModel["multi_object_replication"] = true

		replicationRunModel := make(map[string]interface{})
		replicationRunModel["replication_target_results"] = []map[string]interface{}{replicationTargetResultModel}

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

		archivalDataStatsModel := make(map[string]interface{})
		archivalDataStatsModel["logical_size_bytes"] = int(26)
		archivalDataStatsModel["bytes_read"] = int(26)
		archivalDataStatsModel["logical_bytes_transferred"] = int(26)
		archivalDataStatsModel["physical_bytes_transferred"] = int(26)
		archivalDataStatsModel["avg_logical_transfer_rate_bps"] = int(26)
		archivalDataStatsModel["file_walk_done"] = true
		archivalDataStatsModel["total_file_count"] = int(26)
		archivalDataStatsModel["backup_file_count"] = int(26)

		wormPropertiesModel := make(map[string]interface{})
		wormPropertiesModel["is_archive_worm_compliant"] = true
		wormPropertiesModel["worm_non_compliance_reason"] = "testString"
		wormPropertiesModel["worm_expiry_time_usecs"] = int(26)

		archivalTargetResultModel := make(map[string]interface{})
		archivalTargetResultModel["target_id"] = int(26)
		archivalTargetResultModel["archival_task_id"] = "testString"
		archivalTargetResultModel["target_name"] = "testString"
		archivalTargetResultModel["target_type"] = "Tape"
		archivalTargetResultModel["usage_type"] = "Archival"
		archivalTargetResultModel["ownership_context"] = "Local"
		archivalTargetResultModel["tier_settings"] = []map[string]interface{}{archivalTargetTierInfoModel}
		archivalTargetResultModel["run_type"] = "kRegular"
		archivalTargetResultModel["is_sla_violated"] = true
		archivalTargetResultModel["snapshot_id"] = "testString"
		archivalTargetResultModel["start_time_usecs"] = int(26)
		archivalTargetResultModel["end_time_usecs"] = int(26)
		archivalTargetResultModel["queued_time_usecs"] = int(26)
		archivalTargetResultModel["is_incremental"] = true
		archivalTargetResultModel["is_forever_incremental"] = true
		archivalTargetResultModel["is_cad_archive"] = true
		archivalTargetResultModel["status"] = "Accepted"
		archivalTargetResultModel["message"] = "testString"
		archivalTargetResultModel["progress_task_id"] = "testString"
		archivalTargetResultModel["stats_task_id"] = "testString"
		archivalTargetResultModel["indexing_task_id"] = "testString"
		archivalTargetResultModel["successful_objects_count"] = int(26)
		archivalTargetResultModel["failed_objects_count"] = int(26)
		archivalTargetResultModel["cancelled_objects_count"] = int(26)
		archivalTargetResultModel["successful_app_objects_count"] = int(38)
		archivalTargetResultModel["failed_app_objects_count"] = int(38)
		archivalTargetResultModel["cancelled_app_objects_count"] = int(38)
		archivalTargetResultModel["stats"] = []map[string]interface{}{archivalDataStatsModel}
		archivalTargetResultModel["is_manually_deleted"] = true
		archivalTargetResultModel["expiry_time_usecs"] = int(26)
		archivalTargetResultModel["data_lock_constraints"] = []map[string]interface{}{dataLockConstraintsModel}
		archivalTargetResultModel["on_legal_hold"] = true
		archivalTargetResultModel["worm_properties"] = []map[string]interface{}{wormPropertiesModel}

		archivalRunModel := make(map[string]interface{})
		archivalRunModel["archival_target_results"] = []map[string]interface{}{archivalTargetResultModel}

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

		cloudSpinDataStatsModel := make(map[string]interface{})
		cloudSpinDataStatsModel["physical_bytes_transferred"] = int(26)

		cloudSpinTargetResultModel := make(map[string]interface{})
		cloudSpinTargetResultModel["aws_params"] = []map[string]interface{}{awsCloudSpinParamsModel}
		cloudSpinTargetResultModel["azure_params"] = []map[string]interface{}{azureCloudSpinParamsModel}
		cloudSpinTargetResultModel["id"] = int(26)
		cloudSpinTargetResultModel["start_time_usecs"] = int(26)
		cloudSpinTargetResultModel["end_time_usecs"] = int(26)
		cloudSpinTargetResultModel["status"] = "Accepted"
		cloudSpinTargetResultModel["message"] = "testString"
		cloudSpinTargetResultModel["stats"] = []map[string]interface{}{cloudSpinDataStatsModel}
		cloudSpinTargetResultModel["is_manually_deleted"] = true
		cloudSpinTargetResultModel["expiry_time_usecs"] = int(26)
		cloudSpinTargetResultModel["cloudspin_task_id"] = "testString"
		cloudSpinTargetResultModel["progress_task_id"] = "testString"
		cloudSpinTargetResultModel["data_lock_constraints"] = []map[string]interface{}{dataLockConstraintsModel}
		cloudSpinTargetResultModel["on_legal_hold"] = true

		cloudSpinRunModel := make(map[string]interface{})
		cloudSpinRunModel["cloud_spin_target_results"] = []map[string]interface{}{cloudSpinTargetResultModel}

		model := make(map[string]interface{})
		model["object"] = []map[string]interface{}{objectSummaryModel}
		model["local_snapshot_info"] = []map[string]interface{}{backupRunModel}
		model["original_backup_info"] = []map[string]interface{}{backupRunModel}
		model["replication_info"] = []map[string]interface{}{replicationRunModel}
		model["archival_info"] = []map[string]interface{}{archivalRunModel}
		model["cloud_spin_info"] = []map[string]interface{}{cloudSpinRunModel}
		model["on_legal_hold"] = true

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

	backupDataStatsModel := new(backuprecoveryv1.BackupDataStats)
	backupDataStatsModel.LogicalSizeBytes = core.Int64Ptr(int64(26))
	backupDataStatsModel.BytesWritten = core.Int64Ptr(int64(26))
	backupDataStatsModel.BytesRead = core.Int64Ptr(int64(26))

	dataLockConstraintsModel := new(backuprecoveryv1.DataLockConstraints)
	dataLockConstraintsModel.Mode = core.StringPtr("Compliance")
	dataLockConstraintsModel.ExpiryTimeUsecs = core.Int64Ptr(int64(26))

	snapshotInfoModel := new(backuprecoveryv1.SnapshotInfo)
	snapshotInfoModel.SnapshotID = core.StringPtr("testString")
	snapshotInfoModel.Status = core.StringPtr("kInProgress")
	snapshotInfoModel.StatusMessage = core.StringPtr("testString")
	snapshotInfoModel.StartTimeUsecs = core.Int64Ptr(int64(26))
	snapshotInfoModel.EndTimeUsecs = core.Int64Ptr(int64(26))
	snapshotInfoModel.AdmittedTimeUsecs = core.Int64Ptr(int64(26))
	snapshotInfoModel.PermitGrantTimeUsecs = core.Int64Ptr(int64(26))
	snapshotInfoModel.QueueDurationUsecs = core.Int64Ptr(int64(26))
	snapshotInfoModel.SnapshotCreationTimeUsecs = core.Int64Ptr(int64(26))
	snapshotInfoModel.Stats = backupDataStatsModel
	snapshotInfoModel.ProgressTaskID = core.StringPtr("testString")
	snapshotInfoModel.IndexingTaskID = core.StringPtr("testString")
	snapshotInfoModel.StatsTaskID = core.StringPtr("testString")
	snapshotInfoModel.Warnings = []string{"testString"}
	snapshotInfoModel.IsManuallyDeleted = core.BoolPtr(true)
	snapshotInfoModel.ExpiryTimeUsecs = core.Int64Ptr(int64(26))
	snapshotInfoModel.TotalFileCount = core.Int64Ptr(int64(26))
	snapshotInfoModel.BackupFileCount = core.Int64Ptr(int64(26))
	snapshotInfoModel.DataLockConstraints = dataLockConstraintsModel

	backupAttemptModel := new(backuprecoveryv1.BackupAttempt)
	backupAttemptModel.StartTimeUsecs = core.Int64Ptr(int64(26))
	backupAttemptModel.EndTimeUsecs = core.Int64Ptr(int64(26))
	backupAttemptModel.AdmittedTimeUsecs = core.Int64Ptr(int64(26))
	backupAttemptModel.PermitGrantTimeUsecs = core.Int64Ptr(int64(26))
	backupAttemptModel.QueueDurationUsecs = core.Int64Ptr(int64(26))
	backupAttemptModel.SnapshotCreationTimeUsecs = core.Int64Ptr(int64(26))
	backupAttemptModel.Status = core.StringPtr("Accepted")
	backupAttemptModel.Stats = backupDataStatsModel
	backupAttemptModel.ProgressTaskID = core.StringPtr("testString")
	backupAttemptModel.Message = core.StringPtr("testString")

	backupRunModel := new(backuprecoveryv1.BackupRun)
	backupRunModel.SnapshotInfo = snapshotInfoModel
	backupRunModel.FailedAttempts = []backuprecoveryv1.BackupAttempt{*backupAttemptModel}

	awsTargetConfigModel := new(backuprecoveryv1.AWSTargetConfig)
	awsTargetConfigModel.Region = core.Int64Ptr(int64(26))
	awsTargetConfigModel.SourceID = core.Int64Ptr(int64(26))

	azureTargetConfigModel := new(backuprecoveryv1.AzureTargetConfig)
	azureTargetConfigModel.ResourceGroup = core.Int64Ptr(int64(26))
	azureTargetConfigModel.SourceID = core.Int64Ptr(int64(26))

	replicationDataStatsModel := new(backuprecoveryv1.ReplicationDataStats)
	replicationDataStatsModel.LogicalSizeBytes = core.Int64Ptr(int64(26))
	replicationDataStatsModel.LogicalBytesTransferred = core.Int64Ptr(int64(26))
	replicationDataStatsModel.PhysicalBytesTransferred = core.Int64Ptr(int64(26))

	replicationTargetResultModel := new(backuprecoveryv1.ReplicationTargetResult)
	replicationTargetResultModel.ClusterID = core.Int64Ptr(int64(26))
	replicationTargetResultModel.ClusterIncarnationID = core.Int64Ptr(int64(26))
	replicationTargetResultModel.AwsTargetConfig = awsTargetConfigModel
	replicationTargetResultModel.AzureTargetConfig = azureTargetConfigModel
	replicationTargetResultModel.StartTimeUsecs = core.Int64Ptr(int64(26))
	replicationTargetResultModel.EndTimeUsecs = core.Int64Ptr(int64(26))
	replicationTargetResultModel.QueuedTimeUsecs = core.Int64Ptr(int64(26))
	replicationTargetResultModel.Status = core.StringPtr("Accepted")
	replicationTargetResultModel.Message = core.StringPtr("testString")
	replicationTargetResultModel.PercentageCompleted = core.Int64Ptr(int64(38))
	replicationTargetResultModel.Stats = replicationDataStatsModel
	replicationTargetResultModel.IsManuallyDeleted = core.BoolPtr(true)
	replicationTargetResultModel.ExpiryTimeUsecs = core.Int64Ptr(int64(26))
	replicationTargetResultModel.ReplicationTaskID = core.StringPtr("testString")
	replicationTargetResultModel.EntriesChanged = core.Int64Ptr(int64(26))
	replicationTargetResultModel.IsInBound = core.BoolPtr(true)
	replicationTargetResultModel.DataLockConstraints = dataLockConstraintsModel
	replicationTargetResultModel.OnLegalHold = core.BoolPtr(true)
	replicationTargetResultModel.MultiObjectReplication = core.BoolPtr(true)

	replicationRunModel := new(backuprecoveryv1.ReplicationRun)
	replicationRunModel.ReplicationTargetResults = []backuprecoveryv1.ReplicationTargetResult{*replicationTargetResultModel}

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

	archivalDataStatsModel := new(backuprecoveryv1.ArchivalDataStats)
	archivalDataStatsModel.LogicalSizeBytes = core.Int64Ptr(int64(26))
	archivalDataStatsModel.BytesRead = core.Int64Ptr(int64(26))
	archivalDataStatsModel.LogicalBytesTransferred = core.Int64Ptr(int64(26))
	archivalDataStatsModel.PhysicalBytesTransferred = core.Int64Ptr(int64(26))
	archivalDataStatsModel.AvgLogicalTransferRateBps = core.Int64Ptr(int64(26))
	archivalDataStatsModel.FileWalkDone = core.BoolPtr(true)
	archivalDataStatsModel.TotalFileCount = core.Int64Ptr(int64(26))
	archivalDataStatsModel.BackupFileCount = core.Int64Ptr(int64(26))

	wormPropertiesModel := new(backuprecoveryv1.WormProperties)
	wormPropertiesModel.IsArchiveWormCompliant = core.BoolPtr(true)
	wormPropertiesModel.WormNonComplianceReason = core.StringPtr("testString")
	wormPropertiesModel.WormExpiryTimeUsecs = core.Int64Ptr(int64(26))

	archivalTargetResultModel := new(backuprecoveryv1.ArchivalTargetResult)
	archivalTargetResultModel.TargetID = core.Int64Ptr(int64(26))
	archivalTargetResultModel.ArchivalTaskID = core.StringPtr("testString")
	archivalTargetResultModel.TargetName = core.StringPtr("testString")
	archivalTargetResultModel.TargetType = core.StringPtr("Tape")
	archivalTargetResultModel.UsageType = core.StringPtr("Archival")
	archivalTargetResultModel.OwnershipContext = core.StringPtr("Local")
	archivalTargetResultModel.TierSettings = archivalTargetTierInfoModel
	archivalTargetResultModel.RunType = core.StringPtr("kRegular")
	archivalTargetResultModel.IsSlaViolated = core.BoolPtr(true)
	archivalTargetResultModel.SnapshotID = core.StringPtr("testString")
	archivalTargetResultModel.StartTimeUsecs = core.Int64Ptr(int64(26))
	archivalTargetResultModel.EndTimeUsecs = core.Int64Ptr(int64(26))
	archivalTargetResultModel.QueuedTimeUsecs = core.Int64Ptr(int64(26))
	archivalTargetResultModel.IsIncremental = core.BoolPtr(true)
	archivalTargetResultModel.IsForeverIncremental = core.BoolPtr(true)
	archivalTargetResultModel.IsCadArchive = core.BoolPtr(true)
	archivalTargetResultModel.Status = core.StringPtr("Accepted")
	archivalTargetResultModel.Message = core.StringPtr("testString")
	archivalTargetResultModel.ProgressTaskID = core.StringPtr("testString")
	archivalTargetResultModel.StatsTaskID = core.StringPtr("testString")
	archivalTargetResultModel.IndexingTaskID = core.StringPtr("testString")
	archivalTargetResultModel.SuccessfulObjectsCount = core.Int64Ptr(int64(26))
	archivalTargetResultModel.FailedObjectsCount = core.Int64Ptr(int64(26))
	archivalTargetResultModel.CancelledObjectsCount = core.Int64Ptr(int64(26))
	archivalTargetResultModel.SuccessfulAppObjectsCount = core.Int64Ptr(int64(38))
	archivalTargetResultModel.FailedAppObjectsCount = core.Int64Ptr(int64(38))
	archivalTargetResultModel.CancelledAppObjectsCount = core.Int64Ptr(int64(38))
	archivalTargetResultModel.Stats = archivalDataStatsModel
	archivalTargetResultModel.IsManuallyDeleted = core.BoolPtr(true)
	archivalTargetResultModel.ExpiryTimeUsecs = core.Int64Ptr(int64(26))
	archivalTargetResultModel.DataLockConstraints = dataLockConstraintsModel
	archivalTargetResultModel.OnLegalHold = core.BoolPtr(true)
	archivalTargetResultModel.WormProperties = wormPropertiesModel

	archivalRunModel := new(backuprecoveryv1.ArchivalRun)
	archivalRunModel.ArchivalTargetResults = []backuprecoveryv1.ArchivalTargetResult{*archivalTargetResultModel}

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

	cloudSpinDataStatsModel := new(backuprecoveryv1.CloudSpinDataStats)
	cloudSpinDataStatsModel.PhysicalBytesTransferred = core.Int64Ptr(int64(26))

	cloudSpinTargetResultModel := new(backuprecoveryv1.CloudSpinTargetResult)
	cloudSpinTargetResultModel.AwsParams = awsCloudSpinParamsModel
	cloudSpinTargetResultModel.AzureParams = azureCloudSpinParamsModel
	cloudSpinTargetResultModel.ID = core.Int64Ptr(int64(26))
	cloudSpinTargetResultModel.StartTimeUsecs = core.Int64Ptr(int64(26))
	cloudSpinTargetResultModel.EndTimeUsecs = core.Int64Ptr(int64(26))
	cloudSpinTargetResultModel.Status = core.StringPtr("Accepted")
	cloudSpinTargetResultModel.Message = core.StringPtr("testString")
	cloudSpinTargetResultModel.Stats = cloudSpinDataStatsModel
	cloudSpinTargetResultModel.IsManuallyDeleted = core.BoolPtr(true)
	cloudSpinTargetResultModel.ExpiryTimeUsecs = core.Int64Ptr(int64(26))
	cloudSpinTargetResultModel.CloudspinTaskID = core.StringPtr("testString")
	cloudSpinTargetResultModel.ProgressTaskID = core.StringPtr("testString")
	cloudSpinTargetResultModel.DataLockConstraints = dataLockConstraintsModel
	cloudSpinTargetResultModel.OnLegalHold = core.BoolPtr(true)

	cloudSpinRunModel := new(backuprecoveryv1.CloudSpinRun)
	cloudSpinRunModel.CloudSpinTargetResults = []backuprecoveryv1.CloudSpinTargetResult{*cloudSpinTargetResultModel}

	model := new(backuprecoveryv1.ObjectRunResult)
	model.Object = objectSummaryModel
	model.LocalSnapshotInfo = backupRunModel
	model.OriginalBackupInfo = backupRunModel
	model.ReplicationInfo = replicationRunModel
	model.ArchivalInfo = archivalRunModel
	model.CloudSpinInfo = cloudSpinRunModel
	model.OnLegalHold = core.BoolPtr(true)

	result, err := backuprecovery.DataSourceIbmBaasProtectionGroupRunObjectRunResultToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasProtectionGroupRunObjectSummaryToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		sharepointObjectParamsModel := make(map[string]interface{})
		sharepointObjectParamsModel["site_web_url"] = "testString"

		objectTypeVCenterParamsModel := make(map[string]interface{})
		objectTypeVCenterParamsModel["is_cloud_env"] = true

		objectTypeWindowsClusterParamsModel := make(map[string]interface{})
		objectTypeWindowsClusterParamsModel["cluster_source_type"] = "testString"
		objectSummaryParamsModel := make(map[string]interface{})
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
		model["child_objects"] = []map[string]interface{}{objectSummaryParamsModel}
		model["v_center_summary"] = []map[string]interface{}{objectTypeVCenterParamsModel}
		model["windows_cluster_summary"] = []map[string]interface{}{objectTypeWindowsClusterParamsModel}

		assert.Equal(t, result, model)
	}

	sharepointObjectParamsModel := new(backuprecoveryv1.SharepointObjectParams)

	sharepointObjectParamsModel.SiteWebURL = core.StringPtr("testString")

	objectTypeVCenterParamsModel := new(backuprecoveryv1.ObjectTypeVCenterParams)
	objectSummaryParamsModel := new([]backuprecoveryv1.ObjectSummary)
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
	model.ChildObjects = *objectSummaryParamsModel
	model.VCenterSummary = objectTypeVCenterParamsModel
	model.WindowsClusterSummary = objectTypeWindowsClusterParamsModel

	result, err := backuprecovery.DataSourceIbmBaasProtectionGroupRunObjectSummaryToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasProtectionGroupRunSharepointObjectParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["site_web_url"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.SharepointObjectParams)
	model.SiteWebURL = core.StringPtr("testString")

	result, err := backuprecovery.DataSourceIbmBaasProtectionGroupRunSharepointObjectParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasProtectionGroupRunObjectTypeVCenterParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["is_cloud_env"] = true

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.ObjectTypeVCenterParams)
	model.IsCloudEnv = core.BoolPtr(true)

	result, err := backuprecovery.DataSourceIbmBaasProtectionGroupRunObjectTypeVCenterParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasProtectionGroupRunObjectTypeWindowsClusterParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["cluster_source_type"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.ObjectTypeWindowsClusterParams)
	model.ClusterSourceType = core.StringPtr("testString")

	result, err := backuprecovery.DataSourceIbmBaasProtectionGroupRunObjectTypeWindowsClusterParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasProtectionGroupRunBackupRunToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		backupDataStatsModel := make(map[string]interface{})
		backupDataStatsModel["logical_size_bytes"] = int(26)
		backupDataStatsModel["bytes_written"] = int(26)
		backupDataStatsModel["bytes_read"] = int(26)

		dataLockConstraintsModel := make(map[string]interface{})
		dataLockConstraintsModel["mode"] = "Compliance"
		dataLockConstraintsModel["expiry_time_usecs"] = int(26)

		snapshotInfoModel := make(map[string]interface{})
		snapshotInfoModel["snapshot_id"] = "testString"
		snapshotInfoModel["status"] = "kInProgress"
		snapshotInfoModel["status_message"] = "testString"
		snapshotInfoModel["start_time_usecs"] = int(26)
		snapshotInfoModel["end_time_usecs"] = int(26)
		snapshotInfoModel["admitted_time_usecs"] = int(26)
		snapshotInfoModel["permit_grant_time_usecs"] = int(26)
		snapshotInfoModel["queue_duration_usecs"] = int(26)
		snapshotInfoModel["snapshot_creation_time_usecs"] = int(26)
		snapshotInfoModel["stats"] = []map[string]interface{}{backupDataStatsModel}
		snapshotInfoModel["progress_task_id"] = "testString"
		snapshotInfoModel["indexing_task_id"] = "testString"
		snapshotInfoModel["stats_task_id"] = "testString"
		snapshotInfoModel["warnings"] = []string{"testString"}
		snapshotInfoModel["is_manually_deleted"] = true
		snapshotInfoModel["expiry_time_usecs"] = int(26)
		snapshotInfoModel["total_file_count"] = int(26)
		snapshotInfoModel["backup_file_count"] = int(26)
		snapshotInfoModel["data_lock_constraints"] = []map[string]interface{}{dataLockConstraintsModel}

		backupAttemptModel := make(map[string]interface{})
		backupAttemptModel["start_time_usecs"] = int(26)
		backupAttemptModel["end_time_usecs"] = int(26)
		backupAttemptModel["admitted_time_usecs"] = int(26)
		backupAttemptModel["permit_grant_time_usecs"] = int(26)
		backupAttemptModel["queue_duration_usecs"] = int(26)
		backupAttemptModel["snapshot_creation_time_usecs"] = int(26)
		backupAttemptModel["status"] = "Accepted"
		backupAttemptModel["stats"] = []map[string]interface{}{backupDataStatsModel}
		backupAttemptModel["progress_task_id"] = "testString"
		backupAttemptModel["message"] = "testString"

		model := make(map[string]interface{})
		model["snapshot_info"] = []map[string]interface{}{snapshotInfoModel}
		model["failed_attempts"] = []map[string]interface{}{backupAttemptModel}

		assert.Equal(t, result, model)
	}

	backupDataStatsModel := new(backuprecoveryv1.BackupDataStats)
	backupDataStatsModel.LogicalSizeBytes = core.Int64Ptr(int64(26))
	backupDataStatsModel.BytesWritten = core.Int64Ptr(int64(26))
	backupDataStatsModel.BytesRead = core.Int64Ptr(int64(26))

	dataLockConstraintsModel := new(backuprecoveryv1.DataLockConstraints)
	dataLockConstraintsModel.Mode = core.StringPtr("Compliance")
	dataLockConstraintsModel.ExpiryTimeUsecs = core.Int64Ptr(int64(26))

	snapshotInfoModel := new(backuprecoveryv1.SnapshotInfo)
	snapshotInfoModel.SnapshotID = core.StringPtr("testString")
	snapshotInfoModel.Status = core.StringPtr("kInProgress")
	snapshotInfoModel.StatusMessage = core.StringPtr("testString")
	snapshotInfoModel.StartTimeUsecs = core.Int64Ptr(int64(26))
	snapshotInfoModel.EndTimeUsecs = core.Int64Ptr(int64(26))
	snapshotInfoModel.AdmittedTimeUsecs = core.Int64Ptr(int64(26))
	snapshotInfoModel.PermitGrantTimeUsecs = core.Int64Ptr(int64(26))
	snapshotInfoModel.QueueDurationUsecs = core.Int64Ptr(int64(26))
	snapshotInfoModel.SnapshotCreationTimeUsecs = core.Int64Ptr(int64(26))
	snapshotInfoModel.Stats = backupDataStatsModel
	snapshotInfoModel.ProgressTaskID = core.StringPtr("testString")
	snapshotInfoModel.IndexingTaskID = core.StringPtr("testString")
	snapshotInfoModel.StatsTaskID = core.StringPtr("testString")
	snapshotInfoModel.Warnings = []string{"testString"}
	snapshotInfoModel.IsManuallyDeleted = core.BoolPtr(true)
	snapshotInfoModel.ExpiryTimeUsecs = core.Int64Ptr(int64(26))
	snapshotInfoModel.TotalFileCount = core.Int64Ptr(int64(26))
	snapshotInfoModel.BackupFileCount = core.Int64Ptr(int64(26))
	snapshotInfoModel.DataLockConstraints = dataLockConstraintsModel

	backupAttemptModel := new(backuprecoveryv1.BackupAttempt)
	backupAttemptModel.StartTimeUsecs = core.Int64Ptr(int64(26))
	backupAttemptModel.EndTimeUsecs = core.Int64Ptr(int64(26))
	backupAttemptModel.AdmittedTimeUsecs = core.Int64Ptr(int64(26))
	backupAttemptModel.PermitGrantTimeUsecs = core.Int64Ptr(int64(26))
	backupAttemptModel.QueueDurationUsecs = core.Int64Ptr(int64(26))
	backupAttemptModel.SnapshotCreationTimeUsecs = core.Int64Ptr(int64(26))
	backupAttemptModel.Status = core.StringPtr("Accepted")
	backupAttemptModel.Stats = backupDataStatsModel
	backupAttemptModel.ProgressTaskID = core.StringPtr("testString")
	backupAttemptModel.Message = core.StringPtr("testString")

	model := new(backuprecoveryv1.BackupRun)
	model.SnapshotInfo = snapshotInfoModel
	model.FailedAttempts = []backuprecoveryv1.BackupAttempt{*backupAttemptModel}

	result, err := backuprecovery.DataSourceIbmBaasProtectionGroupRunBackupRunToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasProtectionGroupRunSnapshotInfoToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		backupDataStatsModel := make(map[string]interface{})
		backupDataStatsModel["logical_size_bytes"] = int(26)
		backupDataStatsModel["bytes_written"] = int(26)
		backupDataStatsModel["bytes_read"] = int(26)

		dataLockConstraintsModel := make(map[string]interface{})
		dataLockConstraintsModel["mode"] = "Compliance"
		dataLockConstraintsModel["expiry_time_usecs"] = int(26)

		model := make(map[string]interface{})
		model["snapshot_id"] = "testString"
		model["status"] = "kInProgress"
		model["status_message"] = "testString"
		model["start_time_usecs"] = int(26)
		model["end_time_usecs"] = int(26)
		model["admitted_time_usecs"] = int(26)
		model["permit_grant_time_usecs"] = int(26)
		model["queue_duration_usecs"] = int(26)
		model["snapshot_creation_time_usecs"] = int(26)
		model["stats"] = []map[string]interface{}{backupDataStatsModel}
		model["progress_task_id"] = "testString"
		model["indexing_task_id"] = "testString"
		model["stats_task_id"] = "testString"
		model["warnings"] = []string{"testString"}
		model["is_manually_deleted"] = true
		model["expiry_time_usecs"] = int(26)
		model["total_file_count"] = int(26)
		model["backup_file_count"] = int(26)
		model["data_lock_constraints"] = []map[string]interface{}{dataLockConstraintsModel}

		assert.Equal(t, result, model)
	}

	backupDataStatsModel := new(backuprecoveryv1.BackupDataStats)
	backupDataStatsModel.LogicalSizeBytes = core.Int64Ptr(int64(26))
	backupDataStatsModel.BytesWritten = core.Int64Ptr(int64(26))
	backupDataStatsModel.BytesRead = core.Int64Ptr(int64(26))

	dataLockConstraintsModel := new(backuprecoveryv1.DataLockConstraints)
	dataLockConstraintsModel.Mode = core.StringPtr("Compliance")
	dataLockConstraintsModel.ExpiryTimeUsecs = core.Int64Ptr(int64(26))

	model := new(backuprecoveryv1.SnapshotInfo)
	model.SnapshotID = core.StringPtr("testString")
	model.Status = core.StringPtr("kInProgress")
	model.StatusMessage = core.StringPtr("testString")
	model.StartTimeUsecs = core.Int64Ptr(int64(26))
	model.EndTimeUsecs = core.Int64Ptr(int64(26))
	model.AdmittedTimeUsecs = core.Int64Ptr(int64(26))
	model.PermitGrantTimeUsecs = core.Int64Ptr(int64(26))
	model.QueueDurationUsecs = core.Int64Ptr(int64(26))
	model.SnapshotCreationTimeUsecs = core.Int64Ptr(int64(26))
	model.Stats = backupDataStatsModel
	model.ProgressTaskID = core.StringPtr("testString")
	model.IndexingTaskID = core.StringPtr("testString")
	model.StatsTaskID = core.StringPtr("testString")
	model.Warnings = []string{"testString"}
	model.IsManuallyDeleted = core.BoolPtr(true)
	model.ExpiryTimeUsecs = core.Int64Ptr(int64(26))
	model.TotalFileCount = core.Int64Ptr(int64(26))
	model.BackupFileCount = core.Int64Ptr(int64(26))
	model.DataLockConstraints = dataLockConstraintsModel

	result, err := backuprecovery.DataSourceIbmBaasProtectionGroupRunSnapshotInfoToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasProtectionGroupRunBackupDataStatsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["logical_size_bytes"] = int(26)
		model["bytes_written"] = int(26)
		model["bytes_read"] = int(26)

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.BackupDataStats)
	model.LogicalSizeBytes = core.Int64Ptr(int64(26))
	model.BytesWritten = core.Int64Ptr(int64(26))
	model.BytesRead = core.Int64Ptr(int64(26))

	result, err := backuprecovery.DataSourceIbmBaasProtectionGroupRunBackupDataStatsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasProtectionGroupRunDataLockConstraintsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["mode"] = "Compliance"
		model["expiry_time_usecs"] = int(26)

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.DataLockConstraints)
	model.Mode = core.StringPtr("Compliance")
	model.ExpiryTimeUsecs = core.Int64Ptr(int64(26))

	result, err := backuprecovery.DataSourceIbmBaasProtectionGroupRunDataLockConstraintsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasProtectionGroupRunBackupAttemptToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		backupDataStatsModel := make(map[string]interface{})
		backupDataStatsModel["logical_size_bytes"] = int(26)
		backupDataStatsModel["bytes_written"] = int(26)
		backupDataStatsModel["bytes_read"] = int(26)

		model := make(map[string]interface{})
		model["start_time_usecs"] = int(26)
		model["end_time_usecs"] = int(26)
		model["admitted_time_usecs"] = int(26)
		model["permit_grant_time_usecs"] = int(26)
		model["queue_duration_usecs"] = int(26)
		model["snapshot_creation_time_usecs"] = int(26)
		model["status"] = "Accepted"
		model["stats"] = []map[string]interface{}{backupDataStatsModel}
		model["progress_task_id"] = "testString"
		model["message"] = "testString"

		assert.Equal(t, result, model)
	}

	backupDataStatsModel := new(backuprecoveryv1.BackupDataStats)
	backupDataStatsModel.LogicalSizeBytes = core.Int64Ptr(int64(26))
	backupDataStatsModel.BytesWritten = core.Int64Ptr(int64(26))
	backupDataStatsModel.BytesRead = core.Int64Ptr(int64(26))

	model := new(backuprecoveryv1.BackupAttempt)
	model.StartTimeUsecs = core.Int64Ptr(int64(26))
	model.EndTimeUsecs = core.Int64Ptr(int64(26))
	model.AdmittedTimeUsecs = core.Int64Ptr(int64(26))
	model.PermitGrantTimeUsecs = core.Int64Ptr(int64(26))
	model.QueueDurationUsecs = core.Int64Ptr(int64(26))
	model.SnapshotCreationTimeUsecs = core.Int64Ptr(int64(26))
	model.Status = core.StringPtr("Accepted")
	model.Stats = backupDataStatsModel
	model.ProgressTaskID = core.StringPtr("testString")
	model.Message = core.StringPtr("testString")

	result, err := backuprecovery.DataSourceIbmBaasProtectionGroupRunBackupAttemptToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasProtectionGroupRunReplicationRunToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		awsTargetConfigModel := make(map[string]interface{})
		awsTargetConfigModel["region"] = int(26)
		awsTargetConfigModel["source_id"] = int(26)

		azureTargetConfigModel := make(map[string]interface{})
		azureTargetConfigModel["resource_group"] = int(26)
		azureTargetConfigModel["source_id"] = int(26)

		replicationDataStatsModel := make(map[string]interface{})
		replicationDataStatsModel["logical_size_bytes"] = int(26)
		replicationDataStatsModel["logical_bytes_transferred"] = int(26)
		replicationDataStatsModel["physical_bytes_transferred"] = int(26)

		dataLockConstraintsModel := make(map[string]interface{})
		dataLockConstraintsModel["mode"] = "Compliance"
		dataLockConstraintsModel["expiry_time_usecs"] = int(26)

		replicationTargetResultModel := make(map[string]interface{})
		replicationTargetResultModel["cluster_id"] = int(26)
		replicationTargetResultModel["cluster_incarnation_id"] = int(26)
		replicationTargetResultModel["aws_target_config"] = []map[string]interface{}{awsTargetConfigModel}
		replicationTargetResultModel["azure_target_config"] = []map[string]interface{}{azureTargetConfigModel}
		replicationTargetResultModel["start_time_usecs"] = int(26)
		replicationTargetResultModel["end_time_usecs"] = int(26)
		replicationTargetResultModel["queued_time_usecs"] = int(26)
		replicationTargetResultModel["status"] = "Accepted"
		replicationTargetResultModel["message"] = "testString"
		replicationTargetResultModel["percentage_completed"] = int(38)
		replicationTargetResultModel["stats"] = []map[string]interface{}{replicationDataStatsModel}
		replicationTargetResultModel["is_manually_deleted"] = true
		replicationTargetResultModel["expiry_time_usecs"] = int(26)
		replicationTargetResultModel["replication_task_id"] = "testString"
		replicationTargetResultModel["entries_changed"] = int(26)
		replicationTargetResultModel["is_in_bound"] = true
		replicationTargetResultModel["data_lock_constraints"] = []map[string]interface{}{dataLockConstraintsModel}
		replicationTargetResultModel["on_legal_hold"] = true
		replicationTargetResultModel["multi_object_replication"] = true

		model := make(map[string]interface{})
		model["replication_target_results"] = []map[string]interface{}{replicationTargetResultModel}

		assert.Equal(t, result, model)
	}

	awsTargetConfigModel := new(backuprecoveryv1.AWSTargetConfig)
	awsTargetConfigModel.Region = core.Int64Ptr(int64(26))
	awsTargetConfigModel.SourceID = core.Int64Ptr(int64(26))

	azureTargetConfigModel := new(backuprecoveryv1.AzureTargetConfig)
	azureTargetConfigModel.ResourceGroup = core.Int64Ptr(int64(26))
	azureTargetConfigModel.SourceID = core.Int64Ptr(int64(26))

	replicationDataStatsModel := new(backuprecoveryv1.ReplicationDataStats)
	replicationDataStatsModel.LogicalSizeBytes = core.Int64Ptr(int64(26))
	replicationDataStatsModel.LogicalBytesTransferred = core.Int64Ptr(int64(26))
	replicationDataStatsModel.PhysicalBytesTransferred = core.Int64Ptr(int64(26))

	dataLockConstraintsModel := new(backuprecoveryv1.DataLockConstraints)
	dataLockConstraintsModel.Mode = core.StringPtr("Compliance")
	dataLockConstraintsModel.ExpiryTimeUsecs = core.Int64Ptr(int64(26))

	replicationTargetResultModel := new(backuprecoveryv1.ReplicationTargetResult)
	replicationTargetResultModel.ClusterID = core.Int64Ptr(int64(26))
	replicationTargetResultModel.ClusterIncarnationID = core.Int64Ptr(int64(26))
	replicationTargetResultModel.AwsTargetConfig = awsTargetConfigModel
	replicationTargetResultModel.AzureTargetConfig = azureTargetConfigModel
	replicationTargetResultModel.StartTimeUsecs = core.Int64Ptr(int64(26))
	replicationTargetResultModel.EndTimeUsecs = core.Int64Ptr(int64(26))
	replicationTargetResultModel.QueuedTimeUsecs = core.Int64Ptr(int64(26))
	replicationTargetResultModel.Status = core.StringPtr("Accepted")
	replicationTargetResultModel.Message = core.StringPtr("testString")
	replicationTargetResultModel.PercentageCompleted = core.Int64Ptr(int64(38))
	replicationTargetResultModel.Stats = replicationDataStatsModel
	replicationTargetResultModel.IsManuallyDeleted = core.BoolPtr(true)
	replicationTargetResultModel.ExpiryTimeUsecs = core.Int64Ptr(int64(26))
	replicationTargetResultModel.ReplicationTaskID = core.StringPtr("testString")
	replicationTargetResultModel.EntriesChanged = core.Int64Ptr(int64(26))
	replicationTargetResultModel.IsInBound = core.BoolPtr(true)
	replicationTargetResultModel.DataLockConstraints = dataLockConstraintsModel
	replicationTargetResultModel.OnLegalHold = core.BoolPtr(true)
	replicationTargetResultModel.MultiObjectReplication = core.BoolPtr(true)

	model := new(backuprecoveryv1.ReplicationRun)
	model.ReplicationTargetResults = []backuprecoveryv1.ReplicationTargetResult{*replicationTargetResultModel}

	result, err := backuprecovery.DataSourceIbmBaasProtectionGroupRunReplicationRunToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasProtectionGroupRunReplicationTargetResultToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		awsTargetConfigModel := make(map[string]interface{})
		awsTargetConfigModel["region"] = int(26)
		awsTargetConfigModel["source_id"] = int(26)

		azureTargetConfigModel := make(map[string]interface{})
		azureTargetConfigModel["resource_group"] = int(26)
		azureTargetConfigModel["source_id"] = int(26)

		replicationDataStatsModel := make(map[string]interface{})
		replicationDataStatsModel["logical_size_bytes"] = int(26)
		replicationDataStatsModel["logical_bytes_transferred"] = int(26)
		replicationDataStatsModel["physical_bytes_transferred"] = int(26)

		dataLockConstraintsModel := make(map[string]interface{})
		dataLockConstraintsModel["mode"] = "Compliance"
		dataLockConstraintsModel["expiry_time_usecs"] = int(26)

		model := make(map[string]interface{})
		model["cluster_id"] = int(26)
		model["cluster_incarnation_id"] = int(26)
		model["cluster_name"] = "testString"
		model["aws_target_config"] = []map[string]interface{}{awsTargetConfigModel}
		model["azure_target_config"] = []map[string]interface{}{azureTargetConfigModel}
		model["start_time_usecs"] = int(26)
		model["end_time_usecs"] = int(26)
		model["queued_time_usecs"] = int(26)
		model["status"] = "Accepted"
		model["message"] = "testString"
		model["percentage_completed"] = int(38)
		model["stats"] = []map[string]interface{}{replicationDataStatsModel}
		model["is_manually_deleted"] = true
		model["expiry_time_usecs"] = int(26)
		model["replication_task_id"] = "testString"
		model["entries_changed"] = int(26)
		model["is_in_bound"] = true
		model["data_lock_constraints"] = []map[string]interface{}{dataLockConstraintsModel}
		model["on_legal_hold"] = true
		model["multi_object_replication"] = true

		assert.Equal(t, result, model)
	}

	awsTargetConfigModel := new(backuprecoveryv1.AWSTargetConfig)
	awsTargetConfigModel.Region = core.Int64Ptr(int64(26))
	awsTargetConfigModel.SourceID = core.Int64Ptr(int64(26))

	azureTargetConfigModel := new(backuprecoveryv1.AzureTargetConfig)
	azureTargetConfigModel.ResourceGroup = core.Int64Ptr(int64(26))
	azureTargetConfigModel.SourceID = core.Int64Ptr(int64(26))

	replicationDataStatsModel := new(backuprecoveryv1.ReplicationDataStats)
	replicationDataStatsModel.LogicalSizeBytes = core.Int64Ptr(int64(26))
	replicationDataStatsModel.LogicalBytesTransferred = core.Int64Ptr(int64(26))
	replicationDataStatsModel.PhysicalBytesTransferred = core.Int64Ptr(int64(26))

	dataLockConstraintsModel := new(backuprecoveryv1.DataLockConstraints)
	dataLockConstraintsModel.Mode = core.StringPtr("Compliance")
	dataLockConstraintsModel.ExpiryTimeUsecs = core.Int64Ptr(int64(26))

	model := new(backuprecoveryv1.ReplicationTargetResult)
	model.ClusterID = core.Int64Ptr(int64(26))
	model.ClusterIncarnationID = core.Int64Ptr(int64(26))
	model.ClusterName = core.StringPtr("testString")
	model.AwsTargetConfig = awsTargetConfigModel
	model.AzureTargetConfig = azureTargetConfigModel
	model.StartTimeUsecs = core.Int64Ptr(int64(26))
	model.EndTimeUsecs = core.Int64Ptr(int64(26))
	model.QueuedTimeUsecs = core.Int64Ptr(int64(26))
	model.Status = core.StringPtr("Accepted")
	model.Message = core.StringPtr("testString")
	model.PercentageCompleted = core.Int64Ptr(int64(38))
	model.Stats = replicationDataStatsModel
	model.IsManuallyDeleted = core.BoolPtr(true)
	model.ExpiryTimeUsecs = core.Int64Ptr(int64(26))
	model.ReplicationTaskID = core.StringPtr("testString")
	model.EntriesChanged = core.Int64Ptr(int64(26))
	model.IsInBound = core.BoolPtr(true)
	model.DataLockConstraints = dataLockConstraintsModel
	model.OnLegalHold = core.BoolPtr(true)
	model.MultiObjectReplication = core.BoolPtr(true)

	result, err := backuprecovery.DataSourceIbmBaasProtectionGroupRunReplicationTargetResultToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasProtectionGroupRunAWSTargetConfigToMap(t *testing.T) {
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

	result, err := backuprecovery.DataSourceIbmBaasProtectionGroupRunAWSTargetConfigToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasProtectionGroupRunAzureTargetConfigToMap(t *testing.T) {
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

	result, err := backuprecovery.DataSourceIbmBaasProtectionGroupRunAzureTargetConfigToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasProtectionGroupRunReplicationDataStatsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["logical_size_bytes"] = int(26)
		model["logical_bytes_transferred"] = int(26)
		model["physical_bytes_transferred"] = int(26)

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.ReplicationDataStats)
	model.LogicalSizeBytes = core.Int64Ptr(int64(26))
	model.LogicalBytesTransferred = core.Int64Ptr(int64(26))
	model.PhysicalBytesTransferred = core.Int64Ptr(int64(26))

	result, err := backuprecovery.DataSourceIbmBaasProtectionGroupRunReplicationDataStatsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasProtectionGroupRunArchivalRunToMap(t *testing.T) {
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

		archivalDataStatsModel := make(map[string]interface{})
		archivalDataStatsModel["logical_size_bytes"] = int(26)
		archivalDataStatsModel["bytes_read"] = int(26)
		archivalDataStatsModel["logical_bytes_transferred"] = int(26)
		archivalDataStatsModel["physical_bytes_transferred"] = int(26)
		archivalDataStatsModel["avg_logical_transfer_rate_bps"] = int(26)
		archivalDataStatsModel["file_walk_done"] = true
		archivalDataStatsModel["total_file_count"] = int(26)
		archivalDataStatsModel["backup_file_count"] = int(26)

		dataLockConstraintsModel := make(map[string]interface{})
		dataLockConstraintsModel["mode"] = "Compliance"
		dataLockConstraintsModel["expiry_time_usecs"] = int(26)

		wormPropertiesModel := make(map[string]interface{})
		wormPropertiesModel["is_archive_worm_compliant"] = true
		wormPropertiesModel["worm_non_compliance_reason"] = "testString"
		wormPropertiesModel["worm_expiry_time_usecs"] = int(26)

		archivalTargetResultModel := make(map[string]interface{})
		archivalTargetResultModel["target_id"] = int(26)
		archivalTargetResultModel["archival_task_id"] = "testString"
		archivalTargetResultModel["target_name"] = "testString"
		archivalTargetResultModel["target_type"] = "Tape"
		archivalTargetResultModel["usage_type"] = "Archival"
		archivalTargetResultModel["ownership_context"] = "Local"
		archivalTargetResultModel["tier_settings"] = []map[string]interface{}{archivalTargetTierInfoModel}
		archivalTargetResultModel["run_type"] = "kRegular"
		archivalTargetResultModel["is_sla_violated"] = true
		archivalTargetResultModel["snapshot_id"] = "testString"
		archivalTargetResultModel["start_time_usecs"] = int(26)
		archivalTargetResultModel["end_time_usecs"] = int(26)
		archivalTargetResultModel["queued_time_usecs"] = int(26)
		archivalTargetResultModel["is_incremental"] = true
		archivalTargetResultModel["is_forever_incremental"] = true
		archivalTargetResultModel["is_cad_archive"] = true
		archivalTargetResultModel["status"] = "Accepted"
		archivalTargetResultModel["message"] = "testString"
		archivalTargetResultModel["progress_task_id"] = "testString"
		archivalTargetResultModel["stats_task_id"] = "testString"
		archivalTargetResultModel["indexing_task_id"] = "testString"
		archivalTargetResultModel["successful_objects_count"] = int(26)
		archivalTargetResultModel["failed_objects_count"] = int(26)
		archivalTargetResultModel["cancelled_objects_count"] = int(26)
		archivalTargetResultModel["successful_app_objects_count"] = int(38)
		archivalTargetResultModel["failed_app_objects_count"] = int(38)
		archivalTargetResultModel["cancelled_app_objects_count"] = int(38)
		archivalTargetResultModel["stats"] = []map[string]interface{}{archivalDataStatsModel}
		archivalTargetResultModel["is_manually_deleted"] = true
		archivalTargetResultModel["expiry_time_usecs"] = int(26)
		archivalTargetResultModel["data_lock_constraints"] = []map[string]interface{}{dataLockConstraintsModel}
		archivalTargetResultModel["on_legal_hold"] = true
		archivalTargetResultModel["worm_properties"] = []map[string]interface{}{wormPropertiesModel}

		model := make(map[string]interface{})
		model["archival_target_results"] = []map[string]interface{}{archivalTargetResultModel}

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

	archivalDataStatsModel := new(backuprecoveryv1.ArchivalDataStats)
	archivalDataStatsModel.LogicalSizeBytes = core.Int64Ptr(int64(26))
	archivalDataStatsModel.BytesRead = core.Int64Ptr(int64(26))
	archivalDataStatsModel.LogicalBytesTransferred = core.Int64Ptr(int64(26))
	archivalDataStatsModel.PhysicalBytesTransferred = core.Int64Ptr(int64(26))
	archivalDataStatsModel.AvgLogicalTransferRateBps = core.Int64Ptr(int64(26))
	archivalDataStatsModel.FileWalkDone = core.BoolPtr(true)
	archivalDataStatsModel.TotalFileCount = core.Int64Ptr(int64(26))
	archivalDataStatsModel.BackupFileCount = core.Int64Ptr(int64(26))

	dataLockConstraintsModel := new(backuprecoveryv1.DataLockConstraints)
	dataLockConstraintsModel.Mode = core.StringPtr("Compliance")
	dataLockConstraintsModel.ExpiryTimeUsecs = core.Int64Ptr(int64(26))

	wormPropertiesModel := new(backuprecoveryv1.WormProperties)
	wormPropertiesModel.IsArchiveWormCompliant = core.BoolPtr(true)
	wormPropertiesModel.WormNonComplianceReason = core.StringPtr("testString")
	wormPropertiesModel.WormExpiryTimeUsecs = core.Int64Ptr(int64(26))

	archivalTargetResultModel := new(backuprecoveryv1.ArchivalTargetResult)
	archivalTargetResultModel.TargetID = core.Int64Ptr(int64(26))
	archivalTargetResultModel.ArchivalTaskID = core.StringPtr("testString")
	archivalTargetResultModel.TargetName = core.StringPtr("testString")
	archivalTargetResultModel.TargetType = core.StringPtr("Tape")
	archivalTargetResultModel.UsageType = core.StringPtr("Archival")
	archivalTargetResultModel.OwnershipContext = core.StringPtr("Local")
	archivalTargetResultModel.TierSettings = archivalTargetTierInfoModel
	archivalTargetResultModel.RunType = core.StringPtr("kRegular")
	archivalTargetResultModel.IsSlaViolated = core.BoolPtr(true)
	archivalTargetResultModel.SnapshotID = core.StringPtr("testString")
	archivalTargetResultModel.StartTimeUsecs = core.Int64Ptr(int64(26))
	archivalTargetResultModel.EndTimeUsecs = core.Int64Ptr(int64(26))
	archivalTargetResultModel.QueuedTimeUsecs = core.Int64Ptr(int64(26))
	archivalTargetResultModel.IsIncremental = core.BoolPtr(true)
	archivalTargetResultModel.IsForeverIncremental = core.BoolPtr(true)
	archivalTargetResultModel.IsCadArchive = core.BoolPtr(true)
	archivalTargetResultModel.Status = core.StringPtr("Accepted")
	archivalTargetResultModel.Message = core.StringPtr("testString")
	archivalTargetResultModel.ProgressTaskID = core.StringPtr("testString")
	archivalTargetResultModel.StatsTaskID = core.StringPtr("testString")
	archivalTargetResultModel.IndexingTaskID = core.StringPtr("testString")
	archivalTargetResultModel.SuccessfulObjectsCount = core.Int64Ptr(int64(26))
	archivalTargetResultModel.FailedObjectsCount = core.Int64Ptr(int64(26))
	archivalTargetResultModel.CancelledObjectsCount = core.Int64Ptr(int64(26))
	archivalTargetResultModel.SuccessfulAppObjectsCount = core.Int64Ptr(int64(38))
	archivalTargetResultModel.FailedAppObjectsCount = core.Int64Ptr(int64(38))
	archivalTargetResultModel.CancelledAppObjectsCount = core.Int64Ptr(int64(38))
	archivalTargetResultModel.Stats = archivalDataStatsModel
	archivalTargetResultModel.IsManuallyDeleted = core.BoolPtr(true)
	archivalTargetResultModel.ExpiryTimeUsecs = core.Int64Ptr(int64(26))
	archivalTargetResultModel.DataLockConstraints = dataLockConstraintsModel
	archivalTargetResultModel.OnLegalHold = core.BoolPtr(true)
	archivalTargetResultModel.WormProperties = wormPropertiesModel

	model := new(backuprecoveryv1.ArchivalRun)
	model.ArchivalTargetResults = []backuprecoveryv1.ArchivalTargetResult{*archivalTargetResultModel}

	result, err := backuprecovery.DataSourceIbmBaasProtectionGroupRunArchivalRunToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasProtectionGroupRunArchivalTargetResultToMap(t *testing.T) {
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

		archivalDataStatsModel := make(map[string]interface{})
		archivalDataStatsModel["logical_size_bytes"] = int(26)
		archivalDataStatsModel["bytes_read"] = int(26)
		archivalDataStatsModel["logical_bytes_transferred"] = int(26)
		archivalDataStatsModel["physical_bytes_transferred"] = int(26)
		archivalDataStatsModel["avg_logical_transfer_rate_bps"] = int(26)
		archivalDataStatsModel["file_walk_done"] = true
		archivalDataStatsModel["total_file_count"] = int(26)
		archivalDataStatsModel["backup_file_count"] = int(26)

		dataLockConstraintsModel := make(map[string]interface{})
		dataLockConstraintsModel["mode"] = "Compliance"
		dataLockConstraintsModel["expiry_time_usecs"] = int(26)

		wormPropertiesModel := make(map[string]interface{})
		wormPropertiesModel["is_archive_worm_compliant"] = true
		wormPropertiesModel["worm_non_compliance_reason"] = "testString"
		wormPropertiesModel["worm_expiry_time_usecs"] = int(26)

		model := make(map[string]interface{})
		model["target_id"] = int(26)
		model["archival_task_id"] = "testString"
		model["target_name"] = "testString"
		model["target_type"] = "Tape"
		model["usage_type"] = "Archival"
		model["ownership_context"] = "Local"
		model["tier_settings"] = []map[string]interface{}{archivalTargetTierInfoModel}
		model["run_type"] = "kRegular"
		model["is_sla_violated"] = true
		model["snapshot_id"] = "testString"
		model["start_time_usecs"] = int(26)
		model["end_time_usecs"] = int(26)
		model["queued_time_usecs"] = int(26)
		model["is_incremental"] = true
		model["is_forever_incremental"] = true
		model["is_cad_archive"] = true
		model["status"] = "Accepted"
		model["message"] = "testString"
		model["progress_task_id"] = "testString"
		model["stats_task_id"] = "testString"
		model["indexing_task_id"] = "testString"
		model["successful_objects_count"] = int(26)
		model["failed_objects_count"] = int(26)
		model["cancelled_objects_count"] = int(26)
		model["successful_app_objects_count"] = int(38)
		model["failed_app_objects_count"] = int(38)
		model["cancelled_app_objects_count"] = int(38)
		model["stats"] = []map[string]interface{}{archivalDataStatsModel}
		model["is_manually_deleted"] = true
		model["expiry_time_usecs"] = int(26)
		model["data_lock_constraints"] = []map[string]interface{}{dataLockConstraintsModel}
		model["on_legal_hold"] = true
		model["worm_properties"] = []map[string]interface{}{wormPropertiesModel}

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

	archivalDataStatsModel := new(backuprecoveryv1.ArchivalDataStats)
	archivalDataStatsModel.LogicalSizeBytes = core.Int64Ptr(int64(26))
	archivalDataStatsModel.BytesRead = core.Int64Ptr(int64(26))
	archivalDataStatsModel.LogicalBytesTransferred = core.Int64Ptr(int64(26))
	archivalDataStatsModel.PhysicalBytesTransferred = core.Int64Ptr(int64(26))
	archivalDataStatsModel.AvgLogicalTransferRateBps = core.Int64Ptr(int64(26))
	archivalDataStatsModel.FileWalkDone = core.BoolPtr(true)
	archivalDataStatsModel.TotalFileCount = core.Int64Ptr(int64(26))
	archivalDataStatsModel.BackupFileCount = core.Int64Ptr(int64(26))

	dataLockConstraintsModel := new(backuprecoveryv1.DataLockConstraints)
	dataLockConstraintsModel.Mode = core.StringPtr("Compliance")
	dataLockConstraintsModel.ExpiryTimeUsecs = core.Int64Ptr(int64(26))

	wormPropertiesModel := new(backuprecoveryv1.WormProperties)
	wormPropertiesModel.IsArchiveWormCompliant = core.BoolPtr(true)
	wormPropertiesModel.WormNonComplianceReason = core.StringPtr("testString")
	wormPropertiesModel.WormExpiryTimeUsecs = core.Int64Ptr(int64(26))

	model := new(backuprecoveryv1.ArchivalTargetResult)
	model.TargetID = core.Int64Ptr(int64(26))
	model.ArchivalTaskID = core.StringPtr("testString")
	model.TargetName = core.StringPtr("testString")
	model.TargetType = core.StringPtr("Tape")
	model.UsageType = core.StringPtr("Archival")
	model.OwnershipContext = core.StringPtr("Local")
	model.TierSettings = archivalTargetTierInfoModel
	model.RunType = core.StringPtr("kRegular")
	model.IsSlaViolated = core.BoolPtr(true)
	model.SnapshotID = core.StringPtr("testString")
	model.StartTimeUsecs = core.Int64Ptr(int64(26))
	model.EndTimeUsecs = core.Int64Ptr(int64(26))
	model.QueuedTimeUsecs = core.Int64Ptr(int64(26))
	model.IsIncremental = core.BoolPtr(true)
	model.IsForeverIncremental = core.BoolPtr(true)
	model.IsCadArchive = core.BoolPtr(true)
	model.Status = core.StringPtr("Accepted")
	model.Message = core.StringPtr("testString")
	model.ProgressTaskID = core.StringPtr("testString")
	model.StatsTaskID = core.StringPtr("testString")
	model.IndexingTaskID = core.StringPtr("testString")
	model.SuccessfulObjectsCount = core.Int64Ptr(int64(26))
	model.FailedObjectsCount = core.Int64Ptr(int64(26))
	model.CancelledObjectsCount = core.Int64Ptr(int64(26))
	model.SuccessfulAppObjectsCount = core.Int64Ptr(int64(38))
	model.FailedAppObjectsCount = core.Int64Ptr(int64(38))
	model.CancelledAppObjectsCount = core.Int64Ptr(int64(38))
	model.Stats = archivalDataStatsModel
	model.IsManuallyDeleted = core.BoolPtr(true)
	model.ExpiryTimeUsecs = core.Int64Ptr(int64(26))
	model.DataLockConstraints = dataLockConstraintsModel
	model.OnLegalHold = core.BoolPtr(true)
	model.WormProperties = wormPropertiesModel

	result, err := backuprecovery.DataSourceIbmBaasProtectionGroupRunArchivalTargetResultToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasProtectionGroupRunArchivalTargetTierInfoToMap(t *testing.T) {
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

	result, err := backuprecovery.DataSourceIbmBaasProtectionGroupRunArchivalTargetTierInfoToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasProtectionGroupRunAWSTiersToMap(t *testing.T) {
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

	result, err := backuprecovery.DataSourceIbmBaasProtectionGroupRunAWSTiersToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasProtectionGroupRunAWSTierToMap(t *testing.T) {
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

	result, err := backuprecovery.DataSourceIbmBaasProtectionGroupRunAWSTierToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasProtectionGroupRunAzureTiersToMap(t *testing.T) {
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

	result, err := backuprecovery.DataSourceIbmBaasProtectionGroupRunAzureTiersToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasProtectionGroupRunAzureTierToMap(t *testing.T) {
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

	result, err := backuprecovery.DataSourceIbmBaasProtectionGroupRunAzureTierToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasProtectionGroupRunGoogleTiersToMap(t *testing.T) {
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

	result, err := backuprecovery.DataSourceIbmBaasProtectionGroupRunGoogleTiersToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasProtectionGroupRunGoogleTierToMap(t *testing.T) {
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

	result, err := backuprecovery.DataSourceIbmBaasProtectionGroupRunGoogleTierToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasProtectionGroupRunOracleTiersToMap(t *testing.T) {
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

	result, err := backuprecovery.DataSourceIbmBaasProtectionGroupRunOracleTiersToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasProtectionGroupRunOracleTierToMap(t *testing.T) {
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

	result, err := backuprecovery.DataSourceIbmBaasProtectionGroupRunOracleTierToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasProtectionGroupRunArchivalDataStatsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["logical_size_bytes"] = int(26)
		model["bytes_read"] = int(26)
		model["logical_bytes_transferred"] = int(26)
		model["physical_bytes_transferred"] = int(26)
		model["avg_logical_transfer_rate_bps"] = int(26)
		model["file_walk_done"] = true
		model["total_file_count"] = int(26)
		model["backup_file_count"] = int(26)

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.ArchivalDataStats)
	model.LogicalSizeBytes = core.Int64Ptr(int64(26))
	model.BytesRead = core.Int64Ptr(int64(26))
	model.LogicalBytesTransferred = core.Int64Ptr(int64(26))
	model.PhysicalBytesTransferred = core.Int64Ptr(int64(26))
	model.AvgLogicalTransferRateBps = core.Int64Ptr(int64(26))
	model.FileWalkDone = core.BoolPtr(true)
	model.TotalFileCount = core.Int64Ptr(int64(26))
	model.BackupFileCount = core.Int64Ptr(int64(26))

	result, err := backuprecovery.DataSourceIbmBaasProtectionGroupRunArchivalDataStatsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasProtectionGroupRunWormPropertiesToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["is_archive_worm_compliant"] = true
		model["worm_non_compliance_reason"] = "testString"
		model["worm_expiry_time_usecs"] = int(26)

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.WormProperties)
	model.IsArchiveWormCompliant = core.BoolPtr(true)
	model.WormNonComplianceReason = core.StringPtr("testString")
	model.WormExpiryTimeUsecs = core.Int64Ptr(int64(26))

	result, err := backuprecovery.DataSourceIbmBaasProtectionGroupRunWormPropertiesToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasProtectionGroupRunCloudSpinRunToMap(t *testing.T) {
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

		cloudSpinDataStatsModel := make(map[string]interface{})
		cloudSpinDataStatsModel["physical_bytes_transferred"] = int(26)

		dataLockConstraintsModel := make(map[string]interface{})
		dataLockConstraintsModel["mode"] = "Compliance"
		dataLockConstraintsModel["expiry_time_usecs"] = int(26)

		cloudSpinTargetResultModel := make(map[string]interface{})
		cloudSpinTargetResultModel["aws_params"] = []map[string]interface{}{awsCloudSpinParamsModel}
		cloudSpinTargetResultModel["azure_params"] = []map[string]interface{}{azureCloudSpinParamsModel}
		cloudSpinTargetResultModel["id"] = int(26)
		cloudSpinTargetResultModel["start_time_usecs"] = int(26)
		cloudSpinTargetResultModel["end_time_usecs"] = int(26)
		cloudSpinTargetResultModel["status"] = "Accepted"
		cloudSpinTargetResultModel["message"] = "testString"
		cloudSpinTargetResultModel["stats"] = []map[string]interface{}{cloudSpinDataStatsModel}
		cloudSpinTargetResultModel["is_manually_deleted"] = true
		cloudSpinTargetResultModel["expiry_time_usecs"] = int(26)
		cloudSpinTargetResultModel["cloudspin_task_id"] = "testString"
		cloudSpinTargetResultModel["progress_task_id"] = "testString"
		cloudSpinTargetResultModel["data_lock_constraints"] = []map[string]interface{}{dataLockConstraintsModel}
		cloudSpinTargetResultModel["on_legal_hold"] = true

		model := make(map[string]interface{})
		model["cloud_spin_target_results"] = []map[string]interface{}{cloudSpinTargetResultModel}

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

	cloudSpinDataStatsModel := new(backuprecoveryv1.CloudSpinDataStats)
	cloudSpinDataStatsModel.PhysicalBytesTransferred = core.Int64Ptr(int64(26))

	dataLockConstraintsModel := new(backuprecoveryv1.DataLockConstraints)
	dataLockConstraintsModel.Mode = core.StringPtr("Compliance")
	dataLockConstraintsModel.ExpiryTimeUsecs = core.Int64Ptr(int64(26))

	cloudSpinTargetResultModel := new(backuprecoveryv1.CloudSpinTargetResult)
	cloudSpinTargetResultModel.AwsParams = awsCloudSpinParamsModel
	cloudSpinTargetResultModel.AzureParams = azureCloudSpinParamsModel
	cloudSpinTargetResultModel.ID = core.Int64Ptr(int64(26))
	cloudSpinTargetResultModel.StartTimeUsecs = core.Int64Ptr(int64(26))
	cloudSpinTargetResultModel.EndTimeUsecs = core.Int64Ptr(int64(26))
	cloudSpinTargetResultModel.Status = core.StringPtr("Accepted")
	cloudSpinTargetResultModel.Message = core.StringPtr("testString")
	cloudSpinTargetResultModel.Stats = cloudSpinDataStatsModel
	cloudSpinTargetResultModel.IsManuallyDeleted = core.BoolPtr(true)
	cloudSpinTargetResultModel.ExpiryTimeUsecs = core.Int64Ptr(int64(26))
	cloudSpinTargetResultModel.CloudspinTaskID = core.StringPtr("testString")
	cloudSpinTargetResultModel.ProgressTaskID = core.StringPtr("testString")
	cloudSpinTargetResultModel.DataLockConstraints = dataLockConstraintsModel
	cloudSpinTargetResultModel.OnLegalHold = core.BoolPtr(true)

	model := new(backuprecoveryv1.CloudSpinRun)
	model.CloudSpinTargetResults = []backuprecoveryv1.CloudSpinTargetResult{*cloudSpinTargetResultModel}

	result, err := backuprecovery.DataSourceIbmBaasProtectionGroupRunCloudSpinRunToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasProtectionGroupRunCloudSpinTargetResultToMap(t *testing.T) {
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

		cloudSpinDataStatsModel := make(map[string]interface{})
		cloudSpinDataStatsModel["physical_bytes_transferred"] = int(26)

		dataLockConstraintsModel := make(map[string]interface{})
		dataLockConstraintsModel["mode"] = "Compliance"
		dataLockConstraintsModel["expiry_time_usecs"] = int(26)

		model := make(map[string]interface{})
		model["aws_params"] = []map[string]interface{}{awsCloudSpinParamsModel}
		model["azure_params"] = []map[string]interface{}{azureCloudSpinParamsModel}
		model["id"] = int(26)
		model["name"] = "testString"
		model["start_time_usecs"] = int(26)
		model["end_time_usecs"] = int(26)
		model["status"] = "Accepted"
		model["message"] = "testString"
		model["stats"] = []map[string]interface{}{cloudSpinDataStatsModel}
		model["is_manually_deleted"] = true
		model["expiry_time_usecs"] = int(26)
		model["cloudspin_task_id"] = "testString"
		model["progress_task_id"] = "testString"
		model["data_lock_constraints"] = []map[string]interface{}{dataLockConstraintsModel}
		model["on_legal_hold"] = true

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

	cloudSpinDataStatsModel := new(backuprecoveryv1.CloudSpinDataStats)
	cloudSpinDataStatsModel.PhysicalBytesTransferred = core.Int64Ptr(int64(26))

	dataLockConstraintsModel := new(backuprecoveryv1.DataLockConstraints)
	dataLockConstraintsModel.Mode = core.StringPtr("Compliance")
	dataLockConstraintsModel.ExpiryTimeUsecs = core.Int64Ptr(int64(26))

	model := new(backuprecoveryv1.CloudSpinTargetResult)
	model.AwsParams = awsCloudSpinParamsModel
	model.AzureParams = azureCloudSpinParamsModel
	model.ID = core.Int64Ptr(int64(26))
	model.Name = core.StringPtr("testString")
	model.StartTimeUsecs = core.Int64Ptr(int64(26))
	model.EndTimeUsecs = core.Int64Ptr(int64(26))
	model.Status = core.StringPtr("Accepted")
	model.Message = core.StringPtr("testString")
	model.Stats = cloudSpinDataStatsModel
	model.IsManuallyDeleted = core.BoolPtr(true)
	model.ExpiryTimeUsecs = core.Int64Ptr(int64(26))
	model.CloudspinTaskID = core.StringPtr("testString")
	model.ProgressTaskID = core.StringPtr("testString")
	model.DataLockConstraints = dataLockConstraintsModel
	model.OnLegalHold = core.BoolPtr(true)

	result, err := backuprecovery.DataSourceIbmBaasProtectionGroupRunCloudSpinTargetResultToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasProtectionGroupRunAwsCloudSpinParamsToMap(t *testing.T) {
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

	result, err := backuprecovery.DataSourceIbmBaasProtectionGroupRunAwsCloudSpinParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasProtectionGroupRunCustomTagParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["key"] = "testString"
		model["value"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.CustomTagParams)
	model.Key = core.StringPtr("testString")
	model.Value = core.StringPtr("testString")

	result, err := backuprecovery.DataSourceIbmBaasProtectionGroupRunCustomTagParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasProtectionGroupRunAzureCloudSpinParamsToMap(t *testing.T) {
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

	result, err := backuprecovery.DataSourceIbmBaasProtectionGroupRunAzureCloudSpinParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasProtectionGroupRunCloudSpinDataStatsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["physical_bytes_transferred"] = int(26)

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.CloudSpinDataStats)
	model.PhysicalBytesTransferred = core.Int64Ptr(int64(26))

	result, err := backuprecovery.DataSourceIbmBaasProtectionGroupRunCloudSpinDataStatsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasProtectionGroupRunBackupRunSummaryToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		backupDataStatsModel := make(map[string]interface{})
		backupDataStatsModel["logical_size_bytes"] = int(26)
		backupDataStatsModel["bytes_written"] = int(26)
		backupDataStatsModel["bytes_read"] = int(26)

		dataLockConstraintsModel := make(map[string]interface{})
		dataLockConstraintsModel["mode"] = "Compliance"
		dataLockConstraintsModel["expiry_time_usecs"] = int(26)

		model := make(map[string]interface{})
		model["run_type"] = "kRegular"
		model["is_sla_violated"] = true
		model["start_time_usecs"] = int(26)
		model["end_time_usecs"] = int(26)
		model["status"] = "Accepted"
		model["messages"] = []string{"testString"}
		model["successful_objects_count"] = int(26)
		model["skipped_objects_count"] = int(26)
		model["failed_objects_count"] = int(26)
		model["cancelled_objects_count"] = int(26)
		model["successful_app_objects_count"] = int(38)
		model["failed_app_objects_count"] = int(38)
		model["cancelled_app_objects_count"] = int(38)
		model["local_snapshot_stats"] = []map[string]interface{}{backupDataStatsModel}
		model["indexing_task_id"] = "testString"
		model["progress_task_id"] = "testString"
		model["stats_task_id"] = "testString"
		model["data_lock"] = "Compliance"
		model["local_task_id"] = "testString"
		model["data_lock_constraints"] = []map[string]interface{}{dataLockConstraintsModel}

		assert.Equal(t, result, model)
	}

	backupDataStatsModel := new(backuprecoveryv1.BackupDataStats)
	backupDataStatsModel.LogicalSizeBytes = core.Int64Ptr(int64(26))
	backupDataStatsModel.BytesWritten = core.Int64Ptr(int64(26))
	backupDataStatsModel.BytesRead = core.Int64Ptr(int64(26))

	dataLockConstraintsModel := new(backuprecoveryv1.DataLockConstraints)
	dataLockConstraintsModel.Mode = core.StringPtr("Compliance")
	dataLockConstraintsModel.ExpiryTimeUsecs = core.Int64Ptr(int64(26))

	model := new(backuprecoveryv1.BackupRunSummary)
	model.RunType = core.StringPtr("kRegular")
	model.IsSlaViolated = core.BoolPtr(true)
	model.StartTimeUsecs = core.Int64Ptr(int64(26))
	model.EndTimeUsecs = core.Int64Ptr(int64(26))
	model.Status = core.StringPtr("Accepted")
	model.Messages = []string{"testString"}
	model.SuccessfulObjectsCount = core.Int64Ptr(int64(26))
	model.SkippedObjectsCount = core.Int64Ptr(int64(26))
	model.FailedObjectsCount = core.Int64Ptr(int64(26))
	model.CancelledObjectsCount = core.Int64Ptr(int64(26))
	model.SuccessfulAppObjectsCount = core.Int64Ptr(int64(38))
	model.FailedAppObjectsCount = core.Int64Ptr(int64(38))
	model.CancelledAppObjectsCount = core.Int64Ptr(int64(38))
	model.LocalSnapshotStats = backupDataStatsModel
	model.IndexingTaskID = core.StringPtr("testString")
	model.ProgressTaskID = core.StringPtr("testString")
	model.StatsTaskID = core.StringPtr("testString")
	model.DataLock = core.StringPtr("Compliance")
	model.LocalTaskID = core.StringPtr("testString")
	model.DataLockConstraints = dataLockConstraintsModel

	result, err := backuprecovery.DataSourceIbmBaasProtectionGroupRunBackupRunSummaryToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasProtectionGroupRunReplicationRunSummaryToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		awsTargetConfigModel := make(map[string]interface{})
		awsTargetConfigModel["region"] = int(26)
		awsTargetConfigModel["source_id"] = int(26)

		azureTargetConfigModel := make(map[string]interface{})
		azureTargetConfigModel["resource_group"] = int(26)
		azureTargetConfigModel["source_id"] = int(26)

		replicationDataStatsModel := make(map[string]interface{})
		replicationDataStatsModel["logical_size_bytes"] = int(26)
		replicationDataStatsModel["logical_bytes_transferred"] = int(26)
		replicationDataStatsModel["physical_bytes_transferred"] = int(26)

		dataLockConstraintsModel := make(map[string]interface{})
		dataLockConstraintsModel["mode"] = "Compliance"
		dataLockConstraintsModel["expiry_time_usecs"] = int(26)

		replicationTargetResultModel := make(map[string]interface{})
		replicationTargetResultModel["cluster_id"] = int(26)
		replicationTargetResultModel["cluster_incarnation_id"] = int(26)
		replicationTargetResultModel["aws_target_config"] = []map[string]interface{}{awsTargetConfigModel}
		replicationTargetResultModel["azure_target_config"] = []map[string]interface{}{azureTargetConfigModel}
		replicationTargetResultModel["start_time_usecs"] = int(26)
		replicationTargetResultModel["end_time_usecs"] = int(26)
		replicationTargetResultModel["queued_time_usecs"] = int(26)
		replicationTargetResultModel["status"] = "Accepted"
		replicationTargetResultModel["message"] = "testString"
		replicationTargetResultModel["percentage_completed"] = int(38)
		replicationTargetResultModel["stats"] = []map[string]interface{}{replicationDataStatsModel}
		replicationTargetResultModel["is_manually_deleted"] = true
		replicationTargetResultModel["expiry_time_usecs"] = int(26)
		replicationTargetResultModel["replication_task_id"] = "testString"
		replicationTargetResultModel["entries_changed"] = int(26)
		replicationTargetResultModel["is_in_bound"] = true
		replicationTargetResultModel["data_lock_constraints"] = []map[string]interface{}{dataLockConstraintsModel}
		replicationTargetResultModel["on_legal_hold"] = true
		replicationTargetResultModel["multi_object_replication"] = true

		model := make(map[string]interface{})
		model["replication_target_results"] = []map[string]interface{}{replicationTargetResultModel}

		assert.Equal(t, result, model)
	}

	awsTargetConfigModel := new(backuprecoveryv1.AWSTargetConfig)
	awsTargetConfigModel.Region = core.Int64Ptr(int64(26))
	awsTargetConfigModel.SourceID = core.Int64Ptr(int64(26))

	azureTargetConfigModel := new(backuprecoveryv1.AzureTargetConfig)
	azureTargetConfigModel.ResourceGroup = core.Int64Ptr(int64(26))
	azureTargetConfigModel.SourceID = core.Int64Ptr(int64(26))

	replicationDataStatsModel := new(backuprecoveryv1.ReplicationDataStats)
	replicationDataStatsModel.LogicalSizeBytes = core.Int64Ptr(int64(26))
	replicationDataStatsModel.LogicalBytesTransferred = core.Int64Ptr(int64(26))
	replicationDataStatsModel.PhysicalBytesTransferred = core.Int64Ptr(int64(26))

	dataLockConstraintsModel := new(backuprecoveryv1.DataLockConstraints)
	dataLockConstraintsModel.Mode = core.StringPtr("Compliance")
	dataLockConstraintsModel.ExpiryTimeUsecs = core.Int64Ptr(int64(26))

	replicationTargetResultModel := new(backuprecoveryv1.ReplicationTargetResult)
	replicationTargetResultModel.ClusterID = core.Int64Ptr(int64(26))
	replicationTargetResultModel.ClusterIncarnationID = core.Int64Ptr(int64(26))
	replicationTargetResultModel.AwsTargetConfig = awsTargetConfigModel
	replicationTargetResultModel.AzureTargetConfig = azureTargetConfigModel
	replicationTargetResultModel.StartTimeUsecs = core.Int64Ptr(int64(26))
	replicationTargetResultModel.EndTimeUsecs = core.Int64Ptr(int64(26))
	replicationTargetResultModel.QueuedTimeUsecs = core.Int64Ptr(int64(26))
	replicationTargetResultModel.Status = core.StringPtr("Accepted")
	replicationTargetResultModel.Message = core.StringPtr("testString")
	replicationTargetResultModel.PercentageCompleted = core.Int64Ptr(int64(38))
	replicationTargetResultModel.Stats = replicationDataStatsModel
	replicationTargetResultModel.IsManuallyDeleted = core.BoolPtr(true)
	replicationTargetResultModel.ExpiryTimeUsecs = core.Int64Ptr(int64(26))
	replicationTargetResultModel.ReplicationTaskID = core.StringPtr("testString")
	replicationTargetResultModel.EntriesChanged = core.Int64Ptr(int64(26))
	replicationTargetResultModel.IsInBound = core.BoolPtr(true)
	replicationTargetResultModel.DataLockConstraints = dataLockConstraintsModel
	replicationTargetResultModel.OnLegalHold = core.BoolPtr(true)
	replicationTargetResultModel.MultiObjectReplication = core.BoolPtr(true)

	model := new(backuprecoveryv1.ReplicationRunSummary)
	model.ReplicationTargetResults = []backuprecoveryv1.ReplicationTargetResult{*replicationTargetResultModel}

	result, err := backuprecovery.DataSourceIbmBaasProtectionGroupRunReplicationRunSummaryToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasProtectionGroupRunArchivalRunSummaryToMap(t *testing.T) {
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

		archivalDataStatsModel := make(map[string]interface{})
		archivalDataStatsModel["logical_size_bytes"] = int(26)
		archivalDataStatsModel["bytes_read"] = int(26)
		archivalDataStatsModel["logical_bytes_transferred"] = int(26)
		archivalDataStatsModel["physical_bytes_transferred"] = int(26)
		archivalDataStatsModel["avg_logical_transfer_rate_bps"] = int(26)
		archivalDataStatsModel["file_walk_done"] = true
		archivalDataStatsModel["total_file_count"] = int(26)
		archivalDataStatsModel["backup_file_count"] = int(26)

		dataLockConstraintsModel := make(map[string]interface{})
		dataLockConstraintsModel["mode"] = "Compliance"
		dataLockConstraintsModel["expiry_time_usecs"] = int(26)

		wormPropertiesModel := make(map[string]interface{})
		wormPropertiesModel["is_archive_worm_compliant"] = true
		wormPropertiesModel["worm_non_compliance_reason"] = "testString"
		wormPropertiesModel["worm_expiry_time_usecs"] = int(26)

		archivalTargetResultModel := make(map[string]interface{})
		archivalTargetResultModel["target_id"] = int(26)
		archivalTargetResultModel["archival_task_id"] = "testString"
		archivalTargetResultModel["target_name"] = "testString"
		archivalTargetResultModel["target_type"] = "Tape"
		archivalTargetResultModel["usage_type"] = "Archival"
		archivalTargetResultModel["ownership_context"] = "Local"
		archivalTargetResultModel["tier_settings"] = []map[string]interface{}{archivalTargetTierInfoModel}
		archivalTargetResultModel["run_type"] = "kRegular"
		archivalTargetResultModel["is_sla_violated"] = true
		archivalTargetResultModel["snapshot_id"] = "testString"
		archivalTargetResultModel["start_time_usecs"] = int(26)
		archivalTargetResultModel["end_time_usecs"] = int(26)
		archivalTargetResultModel["queued_time_usecs"] = int(26)
		archivalTargetResultModel["is_incremental"] = true
		archivalTargetResultModel["is_forever_incremental"] = true
		archivalTargetResultModel["is_cad_archive"] = true
		archivalTargetResultModel["status"] = "Accepted"
		archivalTargetResultModel["message"] = "testString"
		archivalTargetResultModel["progress_task_id"] = "testString"
		archivalTargetResultModel["stats_task_id"] = "testString"
		archivalTargetResultModel["indexing_task_id"] = "testString"
		archivalTargetResultModel["successful_objects_count"] = int(26)
		archivalTargetResultModel["failed_objects_count"] = int(26)
		archivalTargetResultModel["cancelled_objects_count"] = int(26)
		archivalTargetResultModel["successful_app_objects_count"] = int(38)
		archivalTargetResultModel["failed_app_objects_count"] = int(38)
		archivalTargetResultModel["cancelled_app_objects_count"] = int(38)
		archivalTargetResultModel["stats"] = []map[string]interface{}{archivalDataStatsModel}
		archivalTargetResultModel["is_manually_deleted"] = true
		archivalTargetResultModel["expiry_time_usecs"] = int(26)
		archivalTargetResultModel["data_lock_constraints"] = []map[string]interface{}{dataLockConstraintsModel}
		archivalTargetResultModel["on_legal_hold"] = true
		archivalTargetResultModel["worm_properties"] = []map[string]interface{}{wormPropertiesModel}

		model := make(map[string]interface{})
		model["archival_target_results"] = []map[string]interface{}{archivalTargetResultModel}

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

	archivalDataStatsModel := new(backuprecoveryv1.ArchivalDataStats)
	archivalDataStatsModel.LogicalSizeBytes = core.Int64Ptr(int64(26))
	archivalDataStatsModel.BytesRead = core.Int64Ptr(int64(26))
	archivalDataStatsModel.LogicalBytesTransferred = core.Int64Ptr(int64(26))
	archivalDataStatsModel.PhysicalBytesTransferred = core.Int64Ptr(int64(26))
	archivalDataStatsModel.AvgLogicalTransferRateBps = core.Int64Ptr(int64(26))
	archivalDataStatsModel.FileWalkDone = core.BoolPtr(true)
	archivalDataStatsModel.TotalFileCount = core.Int64Ptr(int64(26))
	archivalDataStatsModel.BackupFileCount = core.Int64Ptr(int64(26))

	dataLockConstraintsModel := new(backuprecoveryv1.DataLockConstraints)
	dataLockConstraintsModel.Mode = core.StringPtr("Compliance")
	dataLockConstraintsModel.ExpiryTimeUsecs = core.Int64Ptr(int64(26))

	wormPropertiesModel := new(backuprecoveryv1.WormProperties)
	wormPropertiesModel.IsArchiveWormCompliant = core.BoolPtr(true)
	wormPropertiesModel.WormNonComplianceReason = core.StringPtr("testString")
	wormPropertiesModel.WormExpiryTimeUsecs = core.Int64Ptr(int64(26))

	archivalTargetResultModel := new(backuprecoveryv1.ArchivalTargetResult)
	archivalTargetResultModel.TargetID = core.Int64Ptr(int64(26))
	archivalTargetResultModel.ArchivalTaskID = core.StringPtr("testString")
	archivalTargetResultModel.TargetName = core.StringPtr("testString")
	archivalTargetResultModel.TargetType = core.StringPtr("Tape")
	archivalTargetResultModel.UsageType = core.StringPtr("Archival")
	archivalTargetResultModel.OwnershipContext = core.StringPtr("Local")
	archivalTargetResultModel.TierSettings = archivalTargetTierInfoModel
	archivalTargetResultModel.RunType = core.StringPtr("kRegular")
	archivalTargetResultModel.IsSlaViolated = core.BoolPtr(true)
	archivalTargetResultModel.SnapshotID = core.StringPtr("testString")
	archivalTargetResultModel.StartTimeUsecs = core.Int64Ptr(int64(26))
	archivalTargetResultModel.EndTimeUsecs = core.Int64Ptr(int64(26))
	archivalTargetResultModel.QueuedTimeUsecs = core.Int64Ptr(int64(26))
	archivalTargetResultModel.IsIncremental = core.BoolPtr(true)
	archivalTargetResultModel.IsForeverIncremental = core.BoolPtr(true)
	archivalTargetResultModel.IsCadArchive = core.BoolPtr(true)
	archivalTargetResultModel.Status = core.StringPtr("Accepted")
	archivalTargetResultModel.Message = core.StringPtr("testString")
	archivalTargetResultModel.ProgressTaskID = core.StringPtr("testString")
	archivalTargetResultModel.StatsTaskID = core.StringPtr("testString")
	archivalTargetResultModel.IndexingTaskID = core.StringPtr("testString")
	archivalTargetResultModel.SuccessfulObjectsCount = core.Int64Ptr(int64(26))
	archivalTargetResultModel.FailedObjectsCount = core.Int64Ptr(int64(26))
	archivalTargetResultModel.CancelledObjectsCount = core.Int64Ptr(int64(26))
	archivalTargetResultModel.SuccessfulAppObjectsCount = core.Int64Ptr(int64(38))
	archivalTargetResultModel.FailedAppObjectsCount = core.Int64Ptr(int64(38))
	archivalTargetResultModel.CancelledAppObjectsCount = core.Int64Ptr(int64(38))
	archivalTargetResultModel.Stats = archivalDataStatsModel
	archivalTargetResultModel.IsManuallyDeleted = core.BoolPtr(true)
	archivalTargetResultModel.ExpiryTimeUsecs = core.Int64Ptr(int64(26))
	archivalTargetResultModel.DataLockConstraints = dataLockConstraintsModel
	archivalTargetResultModel.OnLegalHold = core.BoolPtr(true)
	archivalTargetResultModel.WormProperties = wormPropertiesModel

	model := new(backuprecoveryv1.ArchivalRunSummary)
	model.ArchivalTargetResults = []backuprecoveryv1.ArchivalTargetResult{*archivalTargetResultModel}

	result, err := backuprecovery.DataSourceIbmBaasProtectionGroupRunArchivalRunSummaryToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasProtectionGroupRunCloudSpinRunSummaryToMap(t *testing.T) {
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

		cloudSpinDataStatsModel := make(map[string]interface{})
		cloudSpinDataStatsModel["physical_bytes_transferred"] = int(26)

		dataLockConstraintsModel := make(map[string]interface{})
		dataLockConstraintsModel["mode"] = "Compliance"
		dataLockConstraintsModel["expiry_time_usecs"] = int(26)

		cloudSpinTargetResultModel := make(map[string]interface{})
		cloudSpinTargetResultModel["aws_params"] = []map[string]interface{}{awsCloudSpinParamsModel}
		cloudSpinTargetResultModel["azure_params"] = []map[string]interface{}{azureCloudSpinParamsModel}
		cloudSpinTargetResultModel["id"] = int(26)
		cloudSpinTargetResultModel["start_time_usecs"] = int(26)
		cloudSpinTargetResultModel["end_time_usecs"] = int(26)
		cloudSpinTargetResultModel["status"] = "Accepted"
		cloudSpinTargetResultModel["message"] = "testString"
		cloudSpinTargetResultModel["stats"] = []map[string]interface{}{cloudSpinDataStatsModel}
		cloudSpinTargetResultModel["is_manually_deleted"] = true
		cloudSpinTargetResultModel["expiry_time_usecs"] = int(26)
		cloudSpinTargetResultModel["cloudspin_task_id"] = "testString"
		cloudSpinTargetResultModel["progress_task_id"] = "testString"
		cloudSpinTargetResultModel["data_lock_constraints"] = []map[string]interface{}{dataLockConstraintsModel}
		cloudSpinTargetResultModel["on_legal_hold"] = true

		model := make(map[string]interface{})
		model["cloud_spin_target_results"] = []map[string]interface{}{cloudSpinTargetResultModel}

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

	cloudSpinDataStatsModel := new(backuprecoveryv1.CloudSpinDataStats)
	cloudSpinDataStatsModel.PhysicalBytesTransferred = core.Int64Ptr(int64(26))

	dataLockConstraintsModel := new(backuprecoveryv1.DataLockConstraints)
	dataLockConstraintsModel.Mode = core.StringPtr("Compliance")
	dataLockConstraintsModel.ExpiryTimeUsecs = core.Int64Ptr(int64(26))

	cloudSpinTargetResultModel := new(backuprecoveryv1.CloudSpinTargetResult)
	cloudSpinTargetResultModel.AwsParams = awsCloudSpinParamsModel
	cloudSpinTargetResultModel.AzureParams = azureCloudSpinParamsModel
	cloudSpinTargetResultModel.ID = core.Int64Ptr(int64(26))
	cloudSpinTargetResultModel.StartTimeUsecs = core.Int64Ptr(int64(26))
	cloudSpinTargetResultModel.EndTimeUsecs = core.Int64Ptr(int64(26))
	cloudSpinTargetResultModel.Status = core.StringPtr("Accepted")
	cloudSpinTargetResultModel.Message = core.StringPtr("testString")
	cloudSpinTargetResultModel.Stats = cloudSpinDataStatsModel
	cloudSpinTargetResultModel.IsManuallyDeleted = core.BoolPtr(true)
	cloudSpinTargetResultModel.ExpiryTimeUsecs = core.Int64Ptr(int64(26))
	cloudSpinTargetResultModel.CloudspinTaskID = core.StringPtr("testString")
	cloudSpinTargetResultModel.ProgressTaskID = core.StringPtr("testString")
	cloudSpinTargetResultModel.DataLockConstraints = dataLockConstraintsModel
	cloudSpinTargetResultModel.OnLegalHold = core.BoolPtr(true)

	model := new(backuprecoveryv1.CloudSpinRunSummary)
	model.CloudSpinTargetResults = []backuprecoveryv1.CloudSpinTargetResult{*cloudSpinTargetResultModel}

	result, err := backuprecovery.DataSourceIbmBaasProtectionGroupRunCloudSpinRunSummaryToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasProtectionGroupRunTenantToMap(t *testing.T) {
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

	result, err := backuprecovery.DataSourceIbmBaasProtectionGroupRunTenantToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasProtectionGroupRunExternalVendorTenantMetadataToMap(t *testing.T) {
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

	result, err := backuprecovery.DataSourceIbmBaasProtectionGroupRunExternalVendorTenantMetadataToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasProtectionGroupRunIbmTenantMetadataParamsToMap(t *testing.T) {
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

	result, err := backuprecovery.DataSourceIbmBaasProtectionGroupRunIbmTenantMetadataParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasProtectionGroupRunExternalVendorCustomPropertiesToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["key"] = "testString"
		model["value"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.ExternalVendorCustomProperties)
	model.Key = core.StringPtr("testString")
	model.Value = core.StringPtr("testString")

	result, err := backuprecovery.DataSourceIbmBaasProtectionGroupRunExternalVendorCustomPropertiesToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasProtectionGroupRunTenantNetworkToMap(t *testing.T) {
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

	result, err := backuprecovery.DataSourceIbmBaasProtectionGroupRunTenantNetworkToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
