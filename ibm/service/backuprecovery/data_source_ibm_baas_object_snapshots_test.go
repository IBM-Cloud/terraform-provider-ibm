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

func TestAccIbmBaasObjectSnapshotsDataSourceBasic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmBaasObjectSnapshotsDataSourceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_baas_object_snapshots.baas_object_snapshots_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_object_snapshots.baas_object_snapshots_instance", "baas_object_snapshots_id"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_object_snapshots.baas_object_snapshots_instance", "x_ibm_tenant_id"),
				),
			},
		},
	})
}

func testAccCheckIbmBaasObjectSnapshotsDataSourceConfigBasic() string {
	return fmt.Sprintf(`
		data "ibm_baas_object_snapshots" "baas_object_snapshots_instance" {
			id = 1
			X-IBM-Tenant-Id = "X-IBM-Tenant-Id"
			fromTimeUsecs = 1
			toTimeUsecs = 1
			runStartFromTimeUsecs = 1
			runStartToTimeUsecs = 1
			snapshotActions = [ "RecoverVMs" ]
			runTypes = [ "kRegular" ]
			protectionGroupIds = [ "protectionGroupIds" ]
			runInstanceIds = [ 1 ]
			regionIds = [ "regionIds" ]
			objectActionKeys = [ "kVMware" ]
		}
	`)
}

func TestDataSourceIbmBaasObjectSnapshotsObjectSnapshotToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		awsSnapshotParamsModel := make(map[string]interface{})
		awsSnapshotParamsModel["protection_type"] = "kAgent"

		azureSnapshotParamsModel := make(map[string]interface{})
		azureSnapshotParamsModel["protection_type"] = "kAgent"

		commonNasObjectParamsModel := make(map[string]interface{})
		commonNasObjectParamsModel["supported_nas_mount_protocols"] = []string{"kNoProtocol"}

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

		flashbladeObjectParamsModel := make(map[string]interface{})
		flashbladeObjectParamsModel["supported_nas_mount_protocols"] = []string{"kNfs"}

		hypervSnapshotParamsModel := make(map[string]interface{})
		hypervSnapshotParamsModel["protection_type"] = "kAuto"

		isilonObjectParamsModel := make(map[string]interface{})
		isilonObjectParamsModel["supported_nas_mount_protocols"] = []string{"kNfs"}

		netappObjectParamsModel := make(map[string]interface{})
		netappObjectParamsModel["supported_nas_mount_protocols"] = []string{"kNfs"}
		netappObjectParamsModel["volume_extended_style"] = "kFlexVol"
		netappObjectParamsModel["volume_type"] = "ReadWrite"

		physicalSnapshotParamsModel := make(map[string]interface{})
		physicalSnapshotParamsModel["enable_system_backup"] = true
		physicalSnapshotParamsModel["protection_type"] = "kFile"

		sfdcObjectParamsModel := make(map[string]interface{})
		sfdcObjectParamsModel["records_added"] = int(26)
		sfdcObjectParamsModel["records_modified"] = int(26)
		sfdcObjectParamsModel["records_removed"] = int(26)

		model := make(map[string]interface{})
		model["aws_params"] = []map[string]interface{}{awsSnapshotParamsModel}
		model["azure_params"] = []map[string]interface{}{azureSnapshotParamsModel}
		model["cluster_id"] = int(26)
		model["cluster_incarnation_id"] = int(26)
		model["elastifile_params"] = []map[string]interface{}{commonNasObjectParamsModel}
		model["environment"] = "kVMware"
		model["expiry_time_usecs"] = int(26)
		model["external_target_info"] = []map[string]interface{}{archivalTargetSummaryInfoModel}
		model["flashblade_params"] = []map[string]interface{}{flashbladeObjectParamsModel}
		model["generic_nas_params"] = []map[string]interface{}{commonNasObjectParamsModel}
		model["gpfs_params"] = []map[string]interface{}{commonNasObjectParamsModel}
		model["has_data_lock"] = true
		model["hyperv_params"] = []map[string]interface{}{hypervSnapshotParamsModel}
		model["id"] = "testString"
		model["indexing_status"] = "InProgress"
		model["isilon_params"] = []map[string]interface{}{isilonObjectParamsModel}
		model["netapp_params"] = []map[string]interface{}{netappObjectParamsModel}
		model["object_id"] = int(26)
		model["object_name"] = "testString"
		model["on_legal_hold"] = true
		model["ownership_context"] = "Local"
		model["physical_params"] = []map[string]interface{}{physicalSnapshotParamsModel}
		model["protection_group_id"] = "testString"
		model["protection_group_name"] = "testString"
		model["protection_group_run_id"] = "testString"
		model["region_id"] = "testString"
		model["run_instance_id"] = int(26)
		model["run_start_time_usecs"] = int(26)
		model["run_type"] = "kRegular"
		model["sfdc_params"] = []map[string]interface{}{sfdcObjectParamsModel}
		model["snapshot_target_type"] = "Local"
		model["snapshot_timestamp_usecs"] = int(26)
		model["source_group_id"] = "testString"
		model["source_id"] = int(26)
		model["storage_domain_id"] = int(26)

		assert.Equal(t, result, model)
	}

	awsSnapshotParamsModel := new(backuprecoveryv1.AwsSnapshotParams)
	awsSnapshotParamsModel.ProtectionType = core.StringPtr("kAgent")

	azureSnapshotParamsModel := new(backuprecoveryv1.AzureSnapshotParams)
	azureSnapshotParamsModel.ProtectionType = core.StringPtr("kAgent")

	commonNasObjectParamsModel := new(backuprecoveryv1.CommonNasObjectParams)
	commonNasObjectParamsModel.SupportedNasMountProtocols = []string{"kNoProtocol"}

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

	flashbladeObjectParamsModel := new(backuprecoveryv1.FlashbladeObjectParams)
	flashbladeObjectParamsModel.SupportedNasMountProtocols = []string{"kNfs"}

	hypervSnapshotParamsModel := new(backuprecoveryv1.HypervSnapshotParams)
	hypervSnapshotParamsModel.ProtectionType = core.StringPtr("kAuto")

	isilonObjectParamsModel := new(backuprecoveryv1.IsilonObjectParams)
	isilonObjectParamsModel.SupportedNasMountProtocols = []string{"kNfs"}

	netappObjectParamsModel := new(backuprecoveryv1.NetappObjectParams)
	netappObjectParamsModel.SupportedNasMountProtocols = []string{"kNfs"}
	netappObjectParamsModel.VolumeExtendedStyle = core.StringPtr("kFlexVol")
	netappObjectParamsModel.VolumeType = core.StringPtr("ReadWrite")

	physicalSnapshotParamsModel := new(backuprecoveryv1.PhysicalSnapshotParams)
	physicalSnapshotParamsModel.EnableSystemBackup = core.BoolPtr(true)
	physicalSnapshotParamsModel.ProtectionType = core.StringPtr("kFile")

	sfdcObjectParamsModel := new(backuprecoveryv1.SfdcObjectParams)
	sfdcObjectParamsModel.RecordsAdded = core.Int64Ptr(int64(26))
	sfdcObjectParamsModel.RecordsModified = core.Int64Ptr(int64(26))
	sfdcObjectParamsModel.RecordsRemoved = core.Int64Ptr(int64(26))

	model := new(backuprecoveryv1.ObjectSnapshot)
	model.AwsParams = awsSnapshotParamsModel
	model.AzureParams = azureSnapshotParamsModel
	model.ClusterID = core.Int64Ptr(int64(26))
	model.ClusterIncarnationID = core.Int64Ptr(int64(26))
	model.ElastifileParams = commonNasObjectParamsModel
	model.Environment = core.StringPtr("kVMware")
	model.ExpiryTimeUsecs = core.Int64Ptr(int64(26))
	model.ExternalTargetInfo = archivalTargetSummaryInfoModel
	model.FlashbladeParams = flashbladeObjectParamsModel
	model.GenericNasParams = commonNasObjectParamsModel
	model.GpfsParams = commonNasObjectParamsModel
	model.HasDataLock = core.BoolPtr(true)
	model.HypervParams = hypervSnapshotParamsModel
	model.ID = core.StringPtr("testString")
	model.IndexingStatus = core.StringPtr("InProgress")
	model.IsilonParams = isilonObjectParamsModel
	model.NetappParams = netappObjectParamsModel
	model.ObjectID = core.Int64Ptr(int64(26))
	model.ObjectName = core.StringPtr("testString")
	model.OnLegalHold = core.BoolPtr(true)
	model.OwnershipContext = core.StringPtr("Local")
	model.PhysicalParams = physicalSnapshotParamsModel
	model.ProtectionGroupID = core.StringPtr("testString")
	model.ProtectionGroupName = core.StringPtr("testString")
	model.ProtectionGroupRunID = core.StringPtr("testString")
	model.RegionID = core.StringPtr("testString")
	model.RunInstanceID = core.Int64Ptr(int64(26))
	model.RunStartTimeUsecs = core.Int64Ptr(int64(26))
	model.RunType = core.StringPtr("kRegular")
	model.SfdcParams = sfdcObjectParamsModel
	model.SnapshotTargetType = core.StringPtr("Local")
	model.SnapshotTimestampUsecs = core.Int64Ptr(int64(26))
	model.SourceGroupID = core.StringPtr("testString")
	model.SourceID = core.Int64Ptr(int64(26))
	model.StorageDomainID = core.Int64Ptr(int64(26))

	result, err := backuprecovery.DataSourceIbmBaasObjectSnapshotsObjectSnapshotToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasObjectSnapshotsAwsSnapshotParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["protection_type"] = "kAgent"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.AwsSnapshotParams)
	model.ProtectionType = core.StringPtr("kAgent")

	result, err := backuprecovery.DataSourceIbmBaasObjectSnapshotsAwsSnapshotParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasObjectSnapshotsAzureSnapshotParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["protection_type"] = "kAgent"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.AzureSnapshotParams)
	model.ProtectionType = core.StringPtr("kAgent")

	result, err := backuprecovery.DataSourceIbmBaasObjectSnapshotsAzureSnapshotParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasObjectSnapshotsCommonNasObjectParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["supported_nas_mount_protocols"] = []string{"kNoProtocol"}

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.CommonNasObjectParams)
	model.SupportedNasMountProtocols = []string{"kNoProtocol"}

	result, err := backuprecovery.DataSourceIbmBaasObjectSnapshotsCommonNasObjectParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasObjectSnapshotsArchivalTargetSummaryInfoToMap(t *testing.T) {
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

	result, err := backuprecovery.DataSourceIbmBaasObjectSnapshotsArchivalTargetSummaryInfoToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasObjectSnapshotsArchivalTargetTierInfoToMap(t *testing.T) {
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

	result, err := backuprecovery.DataSourceIbmBaasObjectSnapshotsArchivalTargetTierInfoToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasObjectSnapshotsAWSTiersToMap(t *testing.T) {
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

	result, err := backuprecovery.DataSourceIbmBaasObjectSnapshotsAWSTiersToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasObjectSnapshotsAWSTierToMap(t *testing.T) {
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

	result, err := backuprecovery.DataSourceIbmBaasObjectSnapshotsAWSTierToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasObjectSnapshotsAzureTiersToMap(t *testing.T) {
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

	result, err := backuprecovery.DataSourceIbmBaasObjectSnapshotsAzureTiersToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasObjectSnapshotsAzureTierToMap(t *testing.T) {
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

	result, err := backuprecovery.DataSourceIbmBaasObjectSnapshotsAzureTierToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasObjectSnapshotsGoogleTiersToMap(t *testing.T) {
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

	result, err := backuprecovery.DataSourceIbmBaasObjectSnapshotsGoogleTiersToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasObjectSnapshotsGoogleTierToMap(t *testing.T) {
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

	result, err := backuprecovery.DataSourceIbmBaasObjectSnapshotsGoogleTierToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasObjectSnapshotsOracleTiersToMap(t *testing.T) {
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

	result, err := backuprecovery.DataSourceIbmBaasObjectSnapshotsOracleTiersToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasObjectSnapshotsOracleTierToMap(t *testing.T) {
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

	result, err := backuprecovery.DataSourceIbmBaasObjectSnapshotsOracleTierToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasObjectSnapshotsFlashbladeObjectParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["supported_nas_mount_protocols"] = []string{"kNfs"}

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.FlashbladeObjectParams)
	model.SupportedNasMountProtocols = []string{"kNfs"}

	result, err := backuprecovery.DataSourceIbmBaasObjectSnapshotsFlashbladeObjectParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasObjectSnapshotsHypervSnapshotParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["protection_type"] = "kAuto"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.HypervSnapshotParams)
	model.ProtectionType = core.StringPtr("kAuto")

	result, err := backuprecovery.DataSourceIbmBaasObjectSnapshotsHypervSnapshotParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasObjectSnapshotsIsilonObjectParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["supported_nas_mount_protocols"] = []string{"kNfs"}

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.IsilonObjectParams)
	model.SupportedNasMountProtocols = []string{"kNfs"}

	result, err := backuprecovery.DataSourceIbmBaasObjectSnapshotsIsilonObjectParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasObjectSnapshotsNetappObjectParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["supported_nas_mount_protocols"] = []string{"kNfs"}
		model["volume_extended_style"] = "kFlexVol"
		model["volume_type"] = "ReadWrite"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.NetappObjectParams)
	model.SupportedNasMountProtocols = []string{"kNfs"}
	model.VolumeExtendedStyle = core.StringPtr("kFlexVol")
	model.VolumeType = core.StringPtr("ReadWrite")

	result, err := backuprecovery.DataSourceIbmBaasObjectSnapshotsNetappObjectParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasObjectSnapshotsPhysicalSnapshotParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["enable_system_backup"] = true
		model["protection_type"] = "kFile"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.PhysicalSnapshotParams)
	model.EnableSystemBackup = core.BoolPtr(true)
	model.ProtectionType = core.StringPtr("kFile")

	result, err := backuprecovery.DataSourceIbmBaasObjectSnapshotsPhysicalSnapshotParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasObjectSnapshotsSfdcObjectParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["records_added"] = int(26)
		model["records_modified"] = int(26)
		model["records_removed"] = int(26)

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.SfdcObjectParams)
	model.RecordsAdded = core.Int64Ptr(int64(26))
	model.RecordsModified = core.Int64Ptr(int64(26))
	model.RecordsRemoved = core.Int64Ptr(int64(26))

	result, err := backuprecovery.DataSourceIbmBaasObjectSnapshotsSfdcObjectParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
