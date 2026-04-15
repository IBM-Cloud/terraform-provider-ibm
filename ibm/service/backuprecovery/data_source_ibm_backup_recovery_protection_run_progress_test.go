// Copyright IBM Corp. 2026 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.113.1-d76630af-20260320-135953
*/

package backuprecovery_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/backuprecovery"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/ibm-backup-recovery-sdk-go/backuprecoveryv1"
	"github.com/stretchr/testify/assert"
	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
)

func TestAccIbmBackupRecoveryProtectionRunProgressDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmBackupRecoveryProtectionRunProgressDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_protection_run_progress.backup_recovery_protection_run_progress_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_protection_run_progress.backup_recovery_protection_run_progress_instance", "x_ibm_tenant_id"),
					resource.TestCheckResourceAttrSet("data.ibm_backup_recovery_protection_run_progress.backup_recovery_protection_run_progress_instance", "run_id"),
				),
			},
		},
	})
}

func testAccCheckIbmBackupRecoveryProtectionRunProgressDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_backup_recovery_protection_run_progress" "backup_recovery_protection_run_progress_instance" {
			X-IBM-Tenant-Id = "tenantId"
			runId = "runId"
			objects = [ 1 ]
			tenantIds = [ "tenantIds" ]
			includeTenants = true
			includeFinishedTasks = true
			startTimeUsecs = 1
			endTimeUsecs = 1
			maxTasksNum = 1
			excludeObjectDetails = true
			includeEventLogs = true
			maxLogLevel = 1
			runTaskPath = "runTaskPath"
			objectTaskPaths = [ "objectTaskPaths" ]
		}
	`)
}


func TestDataSourceIbmBackupRecoveryProtectionRunProgressArchivalTargetProgressInfoToMap(t *testing.T) {
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

		progressTaskEventModel := make(map[string]interface{})
		progressTaskEventModel["message"] = "testString"
		progressTaskEventModel["occured_at_usecs"] = int(26)

		progressStatsModel := make(map[string]interface{})
		progressStatsModel["backup_file_count"] = int(26)
		progressStatsModel["file_walk_done"] = true
		progressStatsModel["total_file_count"] = int(26)

		progressTaskInfoModel := make(map[string]interface{})
		progressTaskInfoModel["end_time_usecs"] = int(26)
		progressTaskInfoModel["events"] = []map[string]interface{}{progressTaskEventModel}
		progressTaskInfoModel["expected_remaining_time_usecs"] = int(26)
		progressTaskInfoModel["percentage_completed"] = float64(36.0)
		progressTaskInfoModel["start_time_usecs"] = int(26)
		progressTaskInfoModel["stats"] = []map[string]interface{}{progressStatsModel}
		progressTaskInfoModel["status"] = "Active"

		objectProgressInfoModel := make(map[string]interface{})
		objectProgressInfoModel["id"] = int(26)
		objectProgressInfoModel["name"] = "testString"
		objectProgressInfoModel["source_id"] = int(26)
		objectProgressInfoModel["source_name"] = "testString"
		objectProgressInfoModel["environment"] = "kPhysical"
		objectProgressInfoModel["end_time_usecs"] = int(26)
		objectProgressInfoModel["events"] = []map[string]interface{}{progressTaskEventModel}
		objectProgressInfoModel["expected_remaining_time_usecs"] = int(26)
		objectProgressInfoModel["percentage_completed"] = float64(36.0)
		objectProgressInfoModel["start_time_usecs"] = int(26)
		objectProgressInfoModel["stats"] = []map[string]interface{}{progressStatsModel}
		objectProgressInfoModel["status"] = "Active"
		objectProgressInfoModel["failed_attempts"] = []map[string]interface{}{progressTaskInfoModel}

		model := make(map[string]interface{})
		model["target_id"] = int(26)
		model["archival_task_id"] = "testString"
		model["target_name"] = "testString"
		model["target_type"] = "Tape"
		model["usage_type"] = "Archival"
		model["ownership_context"] = "Local"
		model["tier_settings"] = []map[string]interface{}{archivalTargetTierInfoModel}
		model["end_time_usecs"] = int(26)
		model["events"] = []map[string]interface{}{progressTaskEventModel}
		model["expected_remaining_time_usecs"] = int(26)
		model["percentage_completed"] = float64(36.0)
		model["start_time_usecs"] = int(26)
		model["stats"] = []map[string]interface{}{progressStatsModel}
		model["status"] = "Active"
		model["objects"] = []map[string]interface{}{objectProgressInfoModel}

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

	progressTaskEventModel := new(backuprecoveryv1.ProgressTaskEvent)
	progressTaskEventModel.Message = core.StringPtr("testString")
	progressTaskEventModel.OccuredAtUsecs = core.Int64Ptr(int64(26))

	progressStatsModel := new(backuprecoveryv1.ProgressStats)
	progressStatsModel.BackupFileCount = core.Int64Ptr(int64(26))
	progressStatsModel.FileWalkDone = core.BoolPtr(true)
	progressStatsModel.TotalFileCount = core.Int64Ptr(int64(26))

	progressTaskInfoModel := new(backuprecoveryv1.ProgressTaskInfo)
	progressTaskInfoModel.EndTimeUsecs = core.Int64Ptr(int64(26))
	progressTaskInfoModel.Events = []backuprecoveryv1.ProgressTaskEvent{*progressTaskEventModel}
	progressTaskInfoModel.ExpectedRemainingTimeUsecs = core.Int64Ptr(int64(26))
	progressTaskInfoModel.PercentageCompleted = core.Float32Ptr(float32(36.0))
	progressTaskInfoModel.StartTimeUsecs = core.Int64Ptr(int64(26))
	progressTaskInfoModel.Stats = progressStatsModel
	progressTaskInfoModel.Status = core.StringPtr("Active")

	objectProgressInfoModel := new(backuprecoveryv1.ObjectProgressInfo)
	objectProgressInfoModel.ID = core.Int64Ptr(int64(26))
	objectProgressInfoModel.Name = core.StringPtr("testString")
	objectProgressInfoModel.SourceID = core.Int64Ptr(int64(26))
	objectProgressInfoModel.SourceName = core.StringPtr("testString")
	objectProgressInfoModel.Environment = core.StringPtr("kPhysical")
	objectProgressInfoModel.EndTimeUsecs = core.Int64Ptr(int64(26))
	objectProgressInfoModel.Events = []backuprecoveryv1.ProgressTaskEvent{*progressTaskEventModel}
	objectProgressInfoModel.ExpectedRemainingTimeUsecs = core.Int64Ptr(int64(26))
	objectProgressInfoModel.PercentageCompleted = core.Float32Ptr(float32(36.0))
	objectProgressInfoModel.StartTimeUsecs = core.Int64Ptr(int64(26))
	objectProgressInfoModel.Stats = progressStatsModel
	objectProgressInfoModel.Status = core.StringPtr("Active")
	objectProgressInfoModel.FailedAttempts = []backuprecoveryv1.ProgressTaskInfo{*progressTaskInfoModel}

	model := new(backuprecoveryv1.ArchivalTargetProgressInfo)
	model.TargetID = core.Int64Ptr(int64(26))
	model.ArchivalTaskID = core.StringPtr("testString")
	model.TargetName = core.StringPtr("testString")
	model.TargetType = core.StringPtr("Tape")
	model.UsageType = core.StringPtr("Archival")
	model.OwnershipContext = core.StringPtr("Local")
	model.TierSettings = archivalTargetTierInfoModel
	model.EndTimeUsecs = core.Int64Ptr(int64(26))
	model.Events = []backuprecoveryv1.ProgressTaskEvent{*progressTaskEventModel}
	model.ExpectedRemainingTimeUsecs = core.Int64Ptr(int64(26))
	model.PercentageCompleted = core.Float32Ptr(float32(36.0))
	model.StartTimeUsecs = core.Int64Ptr(int64(26))
	model.Stats = progressStatsModel
	model.Status = core.StringPtr("Active")
	model.Objects = []backuprecoveryv1.ObjectProgressInfo{*objectProgressInfoModel}

	result, err := backuprecovery.DataSourceIbmBackupRecoveryProtectionRunProgressArchivalTargetProgressInfoToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryProtectionRunProgressArchivalTargetTierInfoToMap(t *testing.T) {
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

	result, err := backuprecovery.DataSourceIbmBackupRecoveryProtectionRunProgressArchivalTargetTierInfoToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryProtectionRunProgressAWSTiersToMap(t *testing.T) {
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

	result, err := backuprecovery.DataSourceIbmBackupRecoveryProtectionRunProgressAWSTiersToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryProtectionRunProgressAWSTierToMap(t *testing.T) {
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

	result, err := backuprecovery.DataSourceIbmBackupRecoveryProtectionRunProgressAWSTierToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryProtectionRunProgressAzureTiersToMap(t *testing.T) {
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

	result, err := backuprecovery.DataSourceIbmBackupRecoveryProtectionRunProgressAzureTiersToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryProtectionRunProgressAzureTierToMap(t *testing.T) {
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

	result, err := backuprecovery.DataSourceIbmBackupRecoveryProtectionRunProgressAzureTierToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryProtectionRunProgressGoogleTiersToMap(t *testing.T) {
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

	result, err := backuprecovery.DataSourceIbmBackupRecoveryProtectionRunProgressGoogleTiersToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryProtectionRunProgressGoogleTierToMap(t *testing.T) {
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

	result, err := backuprecovery.DataSourceIbmBackupRecoveryProtectionRunProgressGoogleTierToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryProtectionRunProgressOracleTiersToMap(t *testing.T) {
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

	result, err := backuprecovery.DataSourceIbmBackupRecoveryProtectionRunProgressOracleTiersToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryProtectionRunProgressOracleTierToMap(t *testing.T) {
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

	result, err := backuprecovery.DataSourceIbmBackupRecoveryProtectionRunProgressOracleTierToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryProtectionRunProgressProgressTaskEventToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["message"] = "testString"
		model["occured_at_usecs"] = int(26)

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.ProgressTaskEvent)
	model.Message = core.StringPtr("testString")
	model.OccuredAtUsecs = core.Int64Ptr(int64(26))

	result, err := backuprecovery.DataSourceIbmBackupRecoveryProtectionRunProgressProgressTaskEventToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryProtectionRunProgressProgressStatsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["backup_file_count"] = int(26)
		model["file_walk_done"] = true
		model["total_file_count"] = int(26)

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.ProgressStats)
	model.BackupFileCount = core.Int64Ptr(int64(26))
	model.FileWalkDone = core.BoolPtr(true)
	model.TotalFileCount = core.Int64Ptr(int64(26))

	result, err := backuprecovery.DataSourceIbmBackupRecoveryProtectionRunProgressProgressStatsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryProtectionRunProgressObjectProgressInfoToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		progressTaskEventModel := make(map[string]interface{})
		progressTaskEventModel["message"] = "testString"
		progressTaskEventModel["occured_at_usecs"] = int(26)

		progressStatsModel := make(map[string]interface{})
		progressStatsModel["backup_file_count"] = int(26)
		progressStatsModel["file_walk_done"] = true
		progressStatsModel["total_file_count"] = int(26)

		progressTaskInfoModel := make(map[string]interface{})
		progressTaskInfoModel["end_time_usecs"] = int(26)
		progressTaskInfoModel["events"] = []map[string]interface{}{progressTaskEventModel}
		progressTaskInfoModel["expected_remaining_time_usecs"] = int(26)
		progressTaskInfoModel["percentage_completed"] = float64(36.0)
		progressTaskInfoModel["start_time_usecs"] = int(26)
		progressTaskInfoModel["stats"] = []map[string]interface{}{progressStatsModel}
		progressTaskInfoModel["status"] = "Active"

		model := make(map[string]interface{})
		model["id"] = int(26)
		model["name"] = "testString"
		model["source_id"] = int(26)
		model["source_name"] = "testString"
		model["environment"] = "kPhysical"
		model["end_time_usecs"] = int(26)
		model["events"] = []map[string]interface{}{progressTaskEventModel}
		model["expected_remaining_time_usecs"] = int(26)
		model["percentage_completed"] = float64(36.0)
		model["start_time_usecs"] = int(26)
		model["stats"] = []map[string]interface{}{progressStatsModel}
		model["status"] = "Active"
		model["failed_attempts"] = []map[string]interface{}{progressTaskInfoModel}

		assert.Equal(t, result, model)
	}

	progressTaskEventModel := new(backuprecoveryv1.ProgressTaskEvent)
	progressTaskEventModel.Message = core.StringPtr("testString")
	progressTaskEventModel.OccuredAtUsecs = core.Int64Ptr(int64(26))

	progressStatsModel := new(backuprecoveryv1.ProgressStats)
	progressStatsModel.BackupFileCount = core.Int64Ptr(int64(26))
	progressStatsModel.FileWalkDone = core.BoolPtr(true)
	progressStatsModel.TotalFileCount = core.Int64Ptr(int64(26))

	progressTaskInfoModel := new(backuprecoveryv1.ProgressTaskInfo)
	progressTaskInfoModel.EndTimeUsecs = core.Int64Ptr(int64(26))
	progressTaskInfoModel.Events = []backuprecoveryv1.ProgressTaskEvent{*progressTaskEventModel}
	progressTaskInfoModel.ExpectedRemainingTimeUsecs = core.Int64Ptr(int64(26))
	progressTaskInfoModel.PercentageCompleted = core.Float32Ptr(float32(36.0))
	progressTaskInfoModel.StartTimeUsecs = core.Int64Ptr(int64(26))
	progressTaskInfoModel.Stats = progressStatsModel
	progressTaskInfoModel.Status = core.StringPtr("Active")

	model := new(backuprecoveryv1.ObjectProgressInfo)
	model.ID = core.Int64Ptr(int64(26))
	model.Name = core.StringPtr("testString")
	model.SourceID = core.Int64Ptr(int64(26))
	model.SourceName = core.StringPtr("testString")
	model.Environment = core.StringPtr("kPhysical")
	model.EndTimeUsecs = core.Int64Ptr(int64(26))
	model.Events = []backuprecoveryv1.ProgressTaskEvent{*progressTaskEventModel}
	model.ExpectedRemainingTimeUsecs = core.Int64Ptr(int64(26))
	model.PercentageCompleted = core.Float32Ptr(float32(36.0))
	model.StartTimeUsecs = core.Int64Ptr(int64(26))
	model.Stats = progressStatsModel
	model.Status = core.StringPtr("Active")
	model.FailedAttempts = []backuprecoveryv1.ProgressTaskInfo{*progressTaskInfoModel}

	result, err := backuprecovery.DataSourceIbmBackupRecoveryProtectionRunProgressObjectProgressInfoToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryProtectionRunProgressProgressTaskInfoToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		progressTaskEventModel := make(map[string]interface{})
		progressTaskEventModel["message"] = "testString"
		progressTaskEventModel["occured_at_usecs"] = int(26)

		progressStatsModel := make(map[string]interface{})
		progressStatsModel["backup_file_count"] = int(26)
		progressStatsModel["file_walk_done"] = true
		progressStatsModel["total_file_count"] = int(26)

		model := make(map[string]interface{})
		model["end_time_usecs"] = int(26)
		model["events"] = []map[string]interface{}{progressTaskEventModel}
		model["expected_remaining_time_usecs"] = int(26)
		model["percentage_completed"] = float64(36.0)
		model["start_time_usecs"] = int(26)
		model["stats"] = []map[string]interface{}{progressStatsModel}
		model["status"] = "Active"

		assert.Equal(t, result, model)
	}

	progressTaskEventModel := new(backuprecoveryv1.ProgressTaskEvent)
	progressTaskEventModel.Message = core.StringPtr("testString")
	progressTaskEventModel.OccuredAtUsecs = core.Int64Ptr(int64(26))

	progressStatsModel := new(backuprecoveryv1.ProgressStats)
	progressStatsModel.BackupFileCount = core.Int64Ptr(int64(26))
	progressStatsModel.FileWalkDone = core.BoolPtr(true)
	progressStatsModel.TotalFileCount = core.Int64Ptr(int64(26))

	model := new(backuprecoveryv1.ProgressTaskInfo)
	model.EndTimeUsecs = core.Int64Ptr(int64(26))
	model.Events = []backuprecoveryv1.ProgressTaskEvent{*progressTaskEventModel}
	model.ExpectedRemainingTimeUsecs = core.Int64Ptr(int64(26))
	model.PercentageCompleted = core.Float32Ptr(float32(36.0))
	model.StartTimeUsecs = core.Int64Ptr(int64(26))
	model.Stats = progressStatsModel
	model.Status = core.StringPtr("Active")

	result, err := backuprecovery.DataSourceIbmBackupRecoveryProtectionRunProgressProgressTaskInfoToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryProtectionRunProgressBackupRunProgressInfoToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		progressTaskEventModel := make(map[string]interface{})
		progressTaskEventModel["message"] = "testString"
		progressTaskEventModel["occured_at_usecs"] = int(26)

		progressStatsModel := make(map[string]interface{})
		progressStatsModel["backup_file_count"] = int(26)
		progressStatsModel["file_walk_done"] = true
		progressStatsModel["total_file_count"] = int(26)

		progressTaskInfoModel := make(map[string]interface{})
		progressTaskInfoModel["end_time_usecs"] = int(26)
		progressTaskInfoModel["events"] = []map[string]interface{}{progressTaskEventModel}
		progressTaskInfoModel["expected_remaining_time_usecs"] = int(26)
		progressTaskInfoModel["percentage_completed"] = float64(36.0)
		progressTaskInfoModel["start_time_usecs"] = int(26)
		progressTaskInfoModel["stats"] = []map[string]interface{}{progressStatsModel}
		progressTaskInfoModel["status"] = "Active"

		objectProgressInfoModel := make(map[string]interface{})
		objectProgressInfoModel["id"] = int(26)
		objectProgressInfoModel["name"] = "testString"
		objectProgressInfoModel["source_id"] = int(26)
		objectProgressInfoModel["source_name"] = "testString"
		objectProgressInfoModel["environment"] = "kPhysical"
		objectProgressInfoModel["end_time_usecs"] = int(26)
		objectProgressInfoModel["events"] = []map[string]interface{}{progressTaskEventModel}
		objectProgressInfoModel["expected_remaining_time_usecs"] = int(26)
		objectProgressInfoModel["percentage_completed"] = float64(36.0)
		objectProgressInfoModel["start_time_usecs"] = int(26)
		objectProgressInfoModel["stats"] = []map[string]interface{}{progressStatsModel}
		objectProgressInfoModel["status"] = "Active"
		objectProgressInfoModel["failed_attempts"] = []map[string]interface{}{progressTaskInfoModel}

		model := make(map[string]interface{})
		model["end_time_usecs"] = int(26)
		model["events"] = []map[string]interface{}{progressTaskEventModel}
		model["expected_remaining_time_usecs"] = int(26)
		model["percentage_completed"] = float64(36.0)
		model["start_time_usecs"] = int(26)
		model["stats"] = []map[string]interface{}{progressStatsModel}
		model["status"] = "Active"
		model["objects"] = []map[string]interface{}{objectProgressInfoModel}

		assert.Equal(t, result, model)
	}

	progressTaskEventModel := new(backuprecoveryv1.ProgressTaskEvent)
	progressTaskEventModel.Message = core.StringPtr("testString")
	progressTaskEventModel.OccuredAtUsecs = core.Int64Ptr(int64(26))

	progressStatsModel := new(backuprecoveryv1.ProgressStats)
	progressStatsModel.BackupFileCount = core.Int64Ptr(int64(26))
	progressStatsModel.FileWalkDone = core.BoolPtr(true)
	progressStatsModel.TotalFileCount = core.Int64Ptr(int64(26))

	progressTaskInfoModel := new(backuprecoveryv1.ProgressTaskInfo)
	progressTaskInfoModel.EndTimeUsecs = core.Int64Ptr(int64(26))
	progressTaskInfoModel.Events = []backuprecoveryv1.ProgressTaskEvent{*progressTaskEventModel}
	progressTaskInfoModel.ExpectedRemainingTimeUsecs = core.Int64Ptr(int64(26))
	progressTaskInfoModel.PercentageCompleted = core.Float32Ptr(float32(36.0))
	progressTaskInfoModel.StartTimeUsecs = core.Int64Ptr(int64(26))
	progressTaskInfoModel.Stats = progressStatsModel
	progressTaskInfoModel.Status = core.StringPtr("Active")

	objectProgressInfoModel := new(backuprecoveryv1.ObjectProgressInfo)
	objectProgressInfoModel.ID = core.Int64Ptr(int64(26))
	objectProgressInfoModel.Name = core.StringPtr("testString")
	objectProgressInfoModel.SourceID = core.Int64Ptr(int64(26))
	objectProgressInfoModel.SourceName = core.StringPtr("testString")
	objectProgressInfoModel.Environment = core.StringPtr("kPhysical")
	objectProgressInfoModel.EndTimeUsecs = core.Int64Ptr(int64(26))
	objectProgressInfoModel.Events = []backuprecoveryv1.ProgressTaskEvent{*progressTaskEventModel}
	objectProgressInfoModel.ExpectedRemainingTimeUsecs = core.Int64Ptr(int64(26))
	objectProgressInfoModel.PercentageCompleted = core.Float32Ptr(float32(36.0))
	objectProgressInfoModel.StartTimeUsecs = core.Int64Ptr(int64(26))
	objectProgressInfoModel.Stats = progressStatsModel
	objectProgressInfoModel.Status = core.StringPtr("Active")
	objectProgressInfoModel.FailedAttempts = []backuprecoveryv1.ProgressTaskInfo{*progressTaskInfoModel}

	model := new(backuprecoveryv1.BackupRunProgressInfo)
	model.EndTimeUsecs = core.Int64Ptr(int64(26))
	model.Events = []backuprecoveryv1.ProgressTaskEvent{*progressTaskEventModel}
	model.ExpectedRemainingTimeUsecs = core.Int64Ptr(int64(26))
	model.PercentageCompleted = core.Float32Ptr(float32(36.0))
	model.StartTimeUsecs = core.Int64Ptr(int64(26))
	model.Stats = progressStatsModel
	model.Status = core.StringPtr("Active")
	model.Objects = []backuprecoveryv1.ObjectProgressInfo{*objectProgressInfoModel}

	result, err := backuprecovery.DataSourceIbmBackupRecoveryProtectionRunProgressBackupRunProgressInfoToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryProtectionRunProgressReplicationTargetProgressInfoToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		awsTargetConfigModel := make(map[string]interface{})
		awsTargetConfigModel["region"] = int(26)
		awsTargetConfigModel["source_id"] = int(26)

		azureTargetConfigModel := make(map[string]interface{})
		azureTargetConfigModel["resource_group"] = int(26)
		azureTargetConfigModel["source_id"] = int(26)

		progressTaskEventModel := make(map[string]interface{})
		progressTaskEventModel["message"] = "testString"
		progressTaskEventModel["occured_at_usecs"] = int(26)

		progressStatsModel := make(map[string]interface{})
		progressStatsModel["backup_file_count"] = int(26)
		progressStatsModel["file_walk_done"] = true
		progressStatsModel["total_file_count"] = int(26)

		progressTaskInfoModel := make(map[string]interface{})
		progressTaskInfoModel["end_time_usecs"] = int(26)
		progressTaskInfoModel["events"] = []map[string]interface{}{progressTaskEventModel}
		progressTaskInfoModel["expected_remaining_time_usecs"] = int(26)
		progressTaskInfoModel["percentage_completed"] = float64(36.0)
		progressTaskInfoModel["start_time_usecs"] = int(26)
		progressTaskInfoModel["stats"] = []map[string]interface{}{progressStatsModel}
		progressTaskInfoModel["status"] = "Active"

		objectProgressInfoModel := make(map[string]interface{})
		objectProgressInfoModel["id"] = int(26)
		objectProgressInfoModel["name"] = "testString"
		objectProgressInfoModel["source_id"] = int(26)
		objectProgressInfoModel["source_name"] = "testString"
		objectProgressInfoModel["environment"] = "kPhysical"
		objectProgressInfoModel["end_time_usecs"] = int(26)
		objectProgressInfoModel["events"] = []map[string]interface{}{progressTaskEventModel}
		objectProgressInfoModel["expected_remaining_time_usecs"] = int(26)
		objectProgressInfoModel["percentage_completed"] = float64(36.0)
		objectProgressInfoModel["start_time_usecs"] = int(26)
		objectProgressInfoModel["stats"] = []map[string]interface{}{progressStatsModel}
		objectProgressInfoModel["status"] = "Active"
		objectProgressInfoModel["failed_attempts"] = []map[string]interface{}{progressTaskInfoModel}

		model := make(map[string]interface{})
		model["cluster_id"] = int(26)
		model["cluster_incarnation_id"] = int(26)
		model["cluster_name"] = "testString"
		model["aws_target_config"] = []map[string]interface{}{awsTargetConfigModel}
		model["azure_target_config"] = []map[string]interface{}{azureTargetConfigModel}
		model["end_time_usecs"] = int(26)
		model["events"] = []map[string]interface{}{progressTaskEventModel}
		model["expected_remaining_time_usecs"] = int(26)
		model["percentage_completed"] = float64(36.0)
		model["start_time_usecs"] = int(26)
		model["stats"] = []map[string]interface{}{progressStatsModel}
		model["status"] = "Active"
		model["objects"] = []map[string]interface{}{objectProgressInfoModel}

		assert.Equal(t, result, model)
	}

	awsTargetConfigModel := new(backuprecoveryv1.AWSTargetConfig)
	awsTargetConfigModel.Region = core.Int64Ptr(int64(26))
	awsTargetConfigModel.SourceID = core.Int64Ptr(int64(26))

	azureTargetConfigModel := new(backuprecoveryv1.AzureTargetConfig)
	azureTargetConfigModel.ResourceGroup = core.Int64Ptr(int64(26))
	azureTargetConfigModel.SourceID = core.Int64Ptr(int64(26))

	progressTaskEventModel := new(backuprecoveryv1.ProgressTaskEvent)
	progressTaskEventModel.Message = core.StringPtr("testString")
	progressTaskEventModel.OccuredAtUsecs = core.Int64Ptr(int64(26))

	progressStatsModel := new(backuprecoveryv1.ProgressStats)
	progressStatsModel.BackupFileCount = core.Int64Ptr(int64(26))
	progressStatsModel.FileWalkDone = core.BoolPtr(true)
	progressStatsModel.TotalFileCount = core.Int64Ptr(int64(26))

	progressTaskInfoModel := new(backuprecoveryv1.ProgressTaskInfo)
	progressTaskInfoModel.EndTimeUsecs = core.Int64Ptr(int64(26))
	progressTaskInfoModel.Events = []backuprecoveryv1.ProgressTaskEvent{*progressTaskEventModel}
	progressTaskInfoModel.ExpectedRemainingTimeUsecs = core.Int64Ptr(int64(26))
	progressTaskInfoModel.PercentageCompleted = core.Float32Ptr(float32(36.0))
	progressTaskInfoModel.StartTimeUsecs = core.Int64Ptr(int64(26))
	progressTaskInfoModel.Stats = progressStatsModel
	progressTaskInfoModel.Status = core.StringPtr("Active")

	objectProgressInfoModel := new(backuprecoveryv1.ObjectProgressInfo)
	objectProgressInfoModel.ID = core.Int64Ptr(int64(26))
	objectProgressInfoModel.Name = core.StringPtr("testString")
	objectProgressInfoModel.SourceID = core.Int64Ptr(int64(26))
	objectProgressInfoModel.SourceName = core.StringPtr("testString")
	objectProgressInfoModel.Environment = core.StringPtr("kPhysical")
	objectProgressInfoModel.EndTimeUsecs = core.Int64Ptr(int64(26))
	objectProgressInfoModel.Events = []backuprecoveryv1.ProgressTaskEvent{*progressTaskEventModel}
	objectProgressInfoModel.ExpectedRemainingTimeUsecs = core.Int64Ptr(int64(26))
	objectProgressInfoModel.PercentageCompleted = core.Float32Ptr(float32(36.0))
	objectProgressInfoModel.StartTimeUsecs = core.Int64Ptr(int64(26))
	objectProgressInfoModel.Stats = progressStatsModel
	objectProgressInfoModel.Status = core.StringPtr("Active")
	objectProgressInfoModel.FailedAttempts = []backuprecoveryv1.ProgressTaskInfo{*progressTaskInfoModel}

	model := new(backuprecoveryv1.ReplicationTargetProgressInfo)
	model.ClusterID = core.Int64Ptr(int64(26))
	model.ClusterIncarnationID = core.Int64Ptr(int64(26))
	model.ClusterName = core.StringPtr("testString")
	model.AwsTargetConfig = awsTargetConfigModel
	model.AzureTargetConfig = azureTargetConfigModel
	model.EndTimeUsecs = core.Int64Ptr(int64(26))
	model.Events = []backuprecoveryv1.ProgressTaskEvent{*progressTaskEventModel}
	model.ExpectedRemainingTimeUsecs = core.Int64Ptr(int64(26))
	model.PercentageCompleted = core.Float32Ptr(float32(36.0))
	model.StartTimeUsecs = core.Int64Ptr(int64(26))
	model.Stats = progressStatsModel
	model.Status = core.StringPtr("Active")
	model.Objects = []backuprecoveryv1.ObjectProgressInfo{*objectProgressInfoModel}

	result, err := backuprecovery.DataSourceIbmBackupRecoveryProtectionRunProgressReplicationTargetProgressInfoToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryProtectionRunProgressAWSTargetConfigToMap(t *testing.T) {
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

	result, err := backuprecovery.DataSourceIbmBackupRecoveryProtectionRunProgressAWSTargetConfigToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBackupRecoveryProtectionRunProgressAzureTargetConfigToMap(t *testing.T) {
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

	result, err := backuprecovery.DataSourceIbmBackupRecoveryProtectionRunProgressAzureTargetConfigToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
