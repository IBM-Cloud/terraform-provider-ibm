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
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/backuprecovery"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/stretchr/testify/assert"
	"github.ibm.com/BackupAndRecovery/ibm-backup-recovery-sdk-go/backuprecoveryv1"
)

func TestAccIbmBaasProtectionGroupRunRequestBasic(t *testing.T) {
	var conf backuprecoveryv1.ProtectionGroupRunsResponse
	xIbmTenantID := fmt.Sprintf("tf_x_ibm_tenant_id_%d", acctest.RandIntRange(10, 100))
	groupID := fmt.Sprintf("group_id_%d", acctest.RandIntRange(10, 100))
	runType := "kRegular"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmBaasProtectionGroupRunRequestDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmBaasProtectionGroupRunRequestConfigBasic(xIbmTenantID, runType, groupID),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmBaasProtectionGroupRunRequestExists("ibm_baas_protection_group_run_request.baas_protection_group_run_request_instance", conf),
					resource.TestCheckResourceAttr("ibm_baas_protection_group_run_request.baas_protection_group_run_request_instance", "x_ibm_tenant_id", xIbmTenantID),
					resource.TestCheckResourceAttr("ibm_baas_protection_group_run_request.baas_protection_group_run_request_instance", "x_ibm_tenant_id", groupID),
					resource.TestCheckResourceAttr("ibm_baas_protection_group_run_request.baas_protection_group_run_request_instance", "group_id", runType),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_baas_protection_group_run_request.baas_protection_group_run_request",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmBaasProtectionGroupRunRequestConfigBasic(xIbmTenantID string, runType, groupID string) string {
	return fmt.Sprintf(`
		resource "ibm_baas_protection_group_run_request" "baas_protection_group_run_request_instance" {
			x_ibm_tenant_id = "%s"
			run_type = "%s"
			group_id = "%s"
		}
	`, xIbmTenantID, runType, groupID)
}

func testAccCheckIbmBaasProtectionGroupRunRequestExists(n string, obj backuprecoveryv1.ProtectionGroupRunsResponse) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		backupRecoveryClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).BackupRecoveryV1()
		if err != nil {
			return err
		}

		getProtectionGroupRunsOptions := &backuprecoveryv1.GetProtectionGroupRunsOptions{}

		getProtectionGroupRunsOptions.SetID(rs.Primary.ID)

		protectionGroupRunRequest, _, err := backupRecoveryClient.GetProtectionGroupRuns(getProtectionGroupRunsOptions)
		if err != nil {
			return err
		}

		obj = *protectionGroupRunRequest
		return nil
	}
}

func testAccCheckIbmBaasProtectionGroupRunRequestDestroy(s *terraform.State) error {
	backupRecoveryClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).BackupRecoveryV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_baas_protection_group_run_request" {
			continue
		}

		getProtectionGroupRunsOptions := &backuprecoveryv1.GetProtectionGroupRunsOptions{}

		getProtectionGroupRunsOptions.SetID(rs.Primary.ID)

		// Try to find the key
		_, response, err := backupRecoveryClient.GetProtectionGroupRuns(getProtectionGroupRunsOptions)

		if err == nil {
			return fmt.Errorf("baas_protection_group_run_request still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for baas_protection_group_run_request (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

func TestResourceIbmBaasProtectionGroupRunRequestRunObjectToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		runObjectPhysicalParamsModel := make(map[string]interface{})
		runObjectPhysicalParamsModel["metadata_file_path"] = "testString"

		model := make(map[string]interface{})
		model["id"] = int(26)
		model["app_ids"] = []int64{int64(26)}
		model["physical_params"] = []map[string]interface{}{runObjectPhysicalParamsModel}

		assert.Equal(t, result, model)
	}

	runObjectPhysicalParamsModel := new(backuprecoveryv1.RunObjectPhysicalParams)
	runObjectPhysicalParamsModel.MetadataFilePath = core.StringPtr("testString")

	model := new(backuprecoveryv1.RunObject)
	model.ID = core.Int64Ptr(int64(26))
	model.AppIds = []int64{int64(26)}
	model.PhysicalParams = runObjectPhysicalParamsModel

	result, err := backuprecovery.ResourceIbmBaasProtectionGroupRunRequestRunObjectToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasProtectionGroupRunRequestRunObjectPhysicalParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["metadata_file_path"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.RunObjectPhysicalParams)
	model.MetadataFilePath = core.StringPtr("testString")

	result, err := backuprecovery.ResourceIbmBaasProtectionGroupRunRequestRunObjectPhysicalParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasProtectionGroupRunRequestRunTargetsConfigurationToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		dataLockConfigModel := make(map[string]interface{})
		dataLockConfigModel["mode"] = "Compliance"
		dataLockConfigModel["unit"] = "Days"
		dataLockConfigModel["duration"] = int(1)
		dataLockConfigModel["enable_worm_on_external_target"] = true

		retentionModel := make(map[string]interface{})
		retentionModel["unit"] = "Days"
		retentionModel["duration"] = int(1)
		retentionModel["data_lock_config"] = []map[string]interface{}{dataLockConfigModel}

		runReplicationConfigModel := make(map[string]interface{})
		runReplicationConfigModel["id"] = int(26)
		runReplicationConfigModel["retention"] = []map[string]interface{}{retentionModel}

		runArchivalConfigModel := make(map[string]interface{})
		runArchivalConfigModel["id"] = int(26)
		runArchivalConfigModel["archival_target_type"] = "Tape"
		runArchivalConfigModel["retention"] = []map[string]interface{}{retentionModel}
		runArchivalConfigModel["copy_only_fully_successful"] = true

		awsTargetConfigModel := make(map[string]interface{})
		awsTargetConfigModel["region"] = int(26)
		awsTargetConfigModel["source_id"] = int(26)

		azureTargetConfigModel := make(map[string]interface{})
		azureTargetConfigModel["resource_group"] = int(26)
		azureTargetConfigModel["source_id"] = int(26)

		runCloudReplicationConfigModel := make(map[string]interface{})
		runCloudReplicationConfigModel["aws_target"] = []map[string]interface{}{awsTargetConfigModel}
		runCloudReplicationConfigModel["azure_target"] = []map[string]interface{}{azureTargetConfigModel}
		runCloudReplicationConfigModel["target_type"] = "AWS"
		runCloudReplicationConfigModel["retention"] = []map[string]interface{}{retentionModel}

		model := make(map[string]interface{})
		model["use_policy_defaults"] = false
		model["replications"] = []map[string]interface{}{runReplicationConfigModel}
		model["archivals"] = []map[string]interface{}{runArchivalConfigModel}
		model["cloud_replications"] = []map[string]interface{}{runCloudReplicationConfigModel}

		assert.Equal(t, result, model)
	}

	dataLockConfigModel := new(backuprecoveryv1.DataLockConfig)
	dataLockConfigModel.Mode = core.StringPtr("Compliance")
	dataLockConfigModel.Unit = core.StringPtr("Days")
	dataLockConfigModel.Duration = core.Int64Ptr(int64(1))
	dataLockConfigModel.EnableWormOnExternalTarget = core.BoolPtr(true)

	retentionModel := new(backuprecoveryv1.Retention)
	retentionModel.Unit = core.StringPtr("Days")
	retentionModel.Duration = core.Int64Ptr(int64(1))
	retentionModel.DataLockConfig = dataLockConfigModel

	runReplicationConfigModel := new(backuprecoveryv1.RunReplicationConfig)
	runReplicationConfigModel.ID = core.Int64Ptr(int64(26))
	runReplicationConfigModel.Retention = retentionModel

	runArchivalConfigModel := new(backuprecoveryv1.RunArchivalConfig)
	runArchivalConfigModel.ID = core.Int64Ptr(int64(26))
	runArchivalConfigModel.ArchivalTargetType = core.StringPtr("Tape")
	runArchivalConfigModel.Retention = retentionModel
	runArchivalConfigModel.CopyOnlyFullySuccessful = core.BoolPtr(true)

	awsTargetConfigModel := new(backuprecoveryv1.AWSTargetConfig)
	awsTargetConfigModel.Region = core.Int64Ptr(int64(26))
	awsTargetConfigModel.SourceID = core.Int64Ptr(int64(26))

	azureTargetConfigModel := new(backuprecoveryv1.AzureTargetConfig)
	azureTargetConfigModel.ResourceGroup = core.Int64Ptr(int64(26))
	azureTargetConfigModel.SourceID = core.Int64Ptr(int64(26))

	runCloudReplicationConfigModel := new(backuprecoveryv1.RunCloudReplicationConfig)
	runCloudReplicationConfigModel.AwsTarget = awsTargetConfigModel
	runCloudReplicationConfigModel.AzureTarget = azureTargetConfigModel
	runCloudReplicationConfigModel.TargetType = core.StringPtr("AWS")
	runCloudReplicationConfigModel.Retention = retentionModel

	model := new(backuprecoveryv1.RunTargetsConfiguration)
	model.UsePolicyDefaults = core.BoolPtr(false)
	model.Replications = []backuprecoveryv1.RunReplicationConfig{*runReplicationConfigModel}
	model.Archivals = []backuprecoveryv1.RunArchivalConfig{*runArchivalConfigModel}
	model.CloudReplications = []backuprecoveryv1.RunCloudReplicationConfig{*runCloudReplicationConfigModel}

	result, err := backuprecovery.ResourceIbmBaasProtectionGroupRunRequestRunTargetsConfigurationToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasProtectionGroupRunRequestRunReplicationConfigToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
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
		model["id"] = int(26)
		model["retention"] = []map[string]interface{}{retentionModel}

		assert.Equal(t, result, model)
	}

	dataLockConfigModel := new(backuprecoveryv1.DataLockConfig)
	dataLockConfigModel.Mode = core.StringPtr("Compliance")
	dataLockConfigModel.Unit = core.StringPtr("Days")
	dataLockConfigModel.Duration = core.Int64Ptr(int64(1))
	dataLockConfigModel.EnableWormOnExternalTarget = core.BoolPtr(true)

	retentionModel := new(backuprecoveryv1.Retention)
	retentionModel.Unit = core.StringPtr("Days")
	retentionModel.Duration = core.Int64Ptr(int64(1))
	retentionModel.DataLockConfig = dataLockConfigModel

	model := new(backuprecoveryv1.RunReplicationConfig)
	model.ID = core.Int64Ptr(int64(26))
	model.Retention = retentionModel

	result, err := backuprecovery.ResourceIbmBaasProtectionGroupRunRequestRunReplicationConfigToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasProtectionGroupRunRequestRetentionToMap(t *testing.T) {
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

	result, err := backuprecovery.ResourceIbmBaasProtectionGroupRunRequestRetentionToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasProtectionGroupRunRequestDataLockConfigToMap(t *testing.T) {
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

	result, err := backuprecovery.ResourceIbmBaasProtectionGroupRunRequestDataLockConfigToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasProtectionGroupRunRequestRunArchivalConfigToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
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
		model["id"] = int(26)
		model["archival_target_type"] = "Tape"
		model["retention"] = []map[string]interface{}{retentionModel}
		model["copy_only_fully_successful"] = true

		assert.Equal(t, result, model)
	}

	dataLockConfigModel := new(backuprecoveryv1.DataLockConfig)
	dataLockConfigModel.Mode = core.StringPtr("Compliance")
	dataLockConfigModel.Unit = core.StringPtr("Days")
	dataLockConfigModel.Duration = core.Int64Ptr(int64(1))
	dataLockConfigModel.EnableWormOnExternalTarget = core.BoolPtr(true)

	retentionModel := new(backuprecoveryv1.Retention)
	retentionModel.Unit = core.StringPtr("Days")
	retentionModel.Duration = core.Int64Ptr(int64(1))
	retentionModel.DataLockConfig = dataLockConfigModel

	model := new(backuprecoveryv1.RunArchivalConfig)
	model.ID = core.Int64Ptr(int64(26))
	model.ArchivalTargetType = core.StringPtr("Tape")
	model.Retention = retentionModel
	model.CopyOnlyFullySuccessful = core.BoolPtr(true)

	result, err := backuprecovery.ResourceIbmBaasProtectionGroupRunRequestRunArchivalConfigToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasProtectionGroupRunRequestRunCloudReplicationConfigToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		awsTargetConfigModel := make(map[string]interface{})
		awsTargetConfigModel["region"] = int(26)
		awsTargetConfigModel["source_id"] = int(26)

		azureTargetConfigModel := make(map[string]interface{})
		azureTargetConfigModel["resource_group"] = int(26)
		azureTargetConfigModel["source_id"] = int(26)

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
		model["aws_target"] = []map[string]interface{}{awsTargetConfigModel}
		model["azure_target"] = []map[string]interface{}{azureTargetConfigModel}
		model["target_type"] = "AWS"
		model["retention"] = []map[string]interface{}{retentionModel}

		assert.Equal(t, result, model)
	}

	awsTargetConfigModel := new(backuprecoveryv1.AWSTargetConfig)
	awsTargetConfigModel.Region = core.Int64Ptr(int64(26))
	awsTargetConfigModel.SourceID = core.Int64Ptr(int64(26))

	azureTargetConfigModel := new(backuprecoveryv1.AzureTargetConfig)
	azureTargetConfigModel.ResourceGroup = core.Int64Ptr(int64(26))
	azureTargetConfigModel.SourceID = core.Int64Ptr(int64(26))

	dataLockConfigModel := new(backuprecoveryv1.DataLockConfig)
	dataLockConfigModel.Mode = core.StringPtr("Compliance")
	dataLockConfigModel.Unit = core.StringPtr("Days")
	dataLockConfigModel.Duration = core.Int64Ptr(int64(1))
	dataLockConfigModel.EnableWormOnExternalTarget = core.BoolPtr(true)

	retentionModel := new(backuprecoveryv1.Retention)
	retentionModel.Unit = core.StringPtr("Days")
	retentionModel.Duration = core.Int64Ptr(int64(1))
	retentionModel.DataLockConfig = dataLockConfigModel

	model := new(backuprecoveryv1.RunCloudReplicationConfig)
	model.AwsTarget = awsTargetConfigModel
	model.AzureTarget = azureTargetConfigModel
	model.TargetType = core.StringPtr("AWS")
	model.Retention = retentionModel

	result, err := backuprecovery.ResourceIbmBaasProtectionGroupRunRequestRunCloudReplicationConfigToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasProtectionGroupRunRequestAWSTargetConfigToMap(t *testing.T) {
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

	result, err := backuprecovery.ResourceIbmBaasProtectionGroupRunRequestAWSTargetConfigToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasProtectionGroupRunRequestAzureTargetConfigToMap(t *testing.T) {
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

	result, err := backuprecovery.ResourceIbmBaasProtectionGroupRunRequestAzureTargetConfigToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasProtectionGroupRunRequestMapToRunObject(t *testing.T) {
	checkResult := func(result *backuprecoveryv1.RunObject) {
		runObjectPhysicalParamsModel := new(backuprecoveryv1.RunObjectPhysicalParams)
		runObjectPhysicalParamsModel.MetadataFilePath = core.StringPtr("testString")

		model := new(backuprecoveryv1.RunObject)
		model.ID = core.Int64Ptr(int64(26))
		model.AppIds = []int64{int64(26)}
		model.PhysicalParams = runObjectPhysicalParamsModel

		assert.Equal(t, result, model)
	}

	runObjectPhysicalParamsModel := make(map[string]interface{})
	runObjectPhysicalParamsModel["metadata_file_path"] = "testString"

	model := make(map[string]interface{})
	model["id"] = int(26)
	model["app_ids"] = []interface{}{int(26)}
	model["physical_params"] = []interface{}{runObjectPhysicalParamsModel}

	result, err := backuprecovery.ResourceIbmBaasProtectionGroupRunRequestMapToRunObject(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasProtectionGroupRunRequestMapToRunObjectPhysicalParams(t *testing.T) {
	checkResult := func(result *backuprecoveryv1.RunObjectPhysicalParams) {
		model := new(backuprecoveryv1.RunObjectPhysicalParams)
		model.MetadataFilePath = core.StringPtr("testString")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["metadata_file_path"] = "testString"

	result, err := backuprecovery.ResourceIbmBaasProtectionGroupRunRequestMapToRunObjectPhysicalParams(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasProtectionGroupRunRequestMapToRunTargetsConfiguration(t *testing.T) {
	checkResult := func(result *backuprecoveryv1.RunTargetsConfiguration) {
		dataLockConfigModel := new(backuprecoveryv1.DataLockConfig)
		dataLockConfigModel.Mode = core.StringPtr("Compliance")
		dataLockConfigModel.Unit = core.StringPtr("Days")
		dataLockConfigModel.Duration = core.Int64Ptr(int64(1))
		dataLockConfigModel.EnableWormOnExternalTarget = core.BoolPtr(true)

		retentionModel := new(backuprecoveryv1.Retention)
		retentionModel.Unit = core.StringPtr("Days")
		retentionModel.Duration = core.Int64Ptr(int64(1))
		retentionModel.DataLockConfig = dataLockConfigModel

		runReplicationConfigModel := new(backuprecoveryv1.RunReplicationConfig)
		runReplicationConfigModel.ID = core.Int64Ptr(int64(26))
		runReplicationConfigModel.Retention = retentionModel

		runArchivalConfigModel := new(backuprecoveryv1.RunArchivalConfig)
		runArchivalConfigModel.ID = core.Int64Ptr(int64(26))
		runArchivalConfigModel.ArchivalTargetType = core.StringPtr("Tape")
		runArchivalConfigModel.Retention = retentionModel
		runArchivalConfigModel.CopyOnlyFullySuccessful = core.BoolPtr(true)

		awsTargetConfigModel := new(backuprecoveryv1.AWSTargetConfig)
		awsTargetConfigModel.Region = core.Int64Ptr(int64(26))
		awsTargetConfigModel.SourceID = core.Int64Ptr(int64(26))

		azureTargetConfigModel := new(backuprecoveryv1.AzureTargetConfig)
		azureTargetConfigModel.ResourceGroup = core.Int64Ptr(int64(26))
		azureTargetConfigModel.SourceID = core.Int64Ptr(int64(26))

		runCloudReplicationConfigModel := new(backuprecoveryv1.RunCloudReplicationConfig)
		runCloudReplicationConfigModel.AwsTarget = awsTargetConfigModel
		runCloudReplicationConfigModel.AzureTarget = azureTargetConfigModel
		runCloudReplicationConfigModel.TargetType = core.StringPtr("AWS")
		runCloudReplicationConfigModel.Retention = retentionModel

		model := new(backuprecoveryv1.RunTargetsConfiguration)
		model.UsePolicyDefaults = core.BoolPtr(false)
		model.Replications = []backuprecoveryv1.RunReplicationConfig{*runReplicationConfigModel}
		model.Archivals = []backuprecoveryv1.RunArchivalConfig{*runArchivalConfigModel}
		model.CloudReplications = []backuprecoveryv1.RunCloudReplicationConfig{*runCloudReplicationConfigModel}

		assert.Equal(t, result, model)
	}

	dataLockConfigModel := make(map[string]interface{})
	dataLockConfigModel["mode"] = "Compliance"
	dataLockConfigModel["unit"] = "Days"
	dataLockConfigModel["duration"] = int(1)
	dataLockConfigModel["enable_worm_on_external_target"] = true

	retentionModel := make(map[string]interface{})
	retentionModel["unit"] = "Days"
	retentionModel["duration"] = int(1)
	retentionModel["data_lock_config"] = []interface{}{dataLockConfigModel}

	runReplicationConfigModel := make(map[string]interface{})
	runReplicationConfigModel["id"] = int(26)
	runReplicationConfigModel["retention"] = []interface{}{retentionModel}

	runArchivalConfigModel := make(map[string]interface{})
	runArchivalConfigModel["id"] = int(26)
	runArchivalConfigModel["archival_target_type"] = "Tape"
	runArchivalConfigModel["retention"] = []interface{}{retentionModel}
	runArchivalConfigModel["copy_only_fully_successful"] = true

	awsTargetConfigModel := make(map[string]interface{})
	awsTargetConfigModel["region"] = int(26)
	awsTargetConfigModel["source_id"] = int(26)

	azureTargetConfigModel := make(map[string]interface{})
	azureTargetConfigModel["resource_group"] = int(26)
	azureTargetConfigModel["source_id"] = int(26)

	runCloudReplicationConfigModel := make(map[string]interface{})
	runCloudReplicationConfigModel["aws_target"] = []interface{}{awsTargetConfigModel}
	runCloudReplicationConfigModel["azure_target"] = []interface{}{azureTargetConfigModel}
	runCloudReplicationConfigModel["target_type"] = "AWS"
	runCloudReplicationConfigModel["retention"] = []interface{}{retentionModel}

	model := make(map[string]interface{})
	model["use_policy_defaults"] = false
	model["replications"] = []interface{}{runReplicationConfigModel}
	model["archivals"] = []interface{}{runArchivalConfigModel}
	model["cloud_replications"] = []interface{}{runCloudReplicationConfigModel}

	result, err := backuprecovery.ResourceIbmBaasProtectionGroupRunRequestMapToRunTargetsConfiguration(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasProtectionGroupRunRequestMapToRunReplicationConfig(t *testing.T) {
	checkResult := func(result *backuprecoveryv1.RunReplicationConfig) {
		dataLockConfigModel := new(backuprecoveryv1.DataLockConfig)
		dataLockConfigModel.Mode = core.StringPtr("Compliance")
		dataLockConfigModel.Unit = core.StringPtr("Days")
		dataLockConfigModel.Duration = core.Int64Ptr(int64(1))
		dataLockConfigModel.EnableWormOnExternalTarget = core.BoolPtr(true)

		retentionModel := new(backuprecoveryv1.Retention)
		retentionModel.Unit = core.StringPtr("Days")
		retentionModel.Duration = core.Int64Ptr(int64(1))
		retentionModel.DataLockConfig = dataLockConfigModel

		model := new(backuprecoveryv1.RunReplicationConfig)
		model.ID = core.Int64Ptr(int64(26))
		model.Retention = retentionModel

		assert.Equal(t, result, model)
	}

	dataLockConfigModel := make(map[string]interface{})
	dataLockConfigModel["mode"] = "Compliance"
	dataLockConfigModel["unit"] = "Days"
	dataLockConfigModel["duration"] = int(1)
	dataLockConfigModel["enable_worm_on_external_target"] = true

	retentionModel := make(map[string]interface{})
	retentionModel["unit"] = "Days"
	retentionModel["duration"] = int(1)
	retentionModel["data_lock_config"] = []interface{}{dataLockConfigModel}

	model := make(map[string]interface{})
	model["id"] = int(26)
	model["retention"] = []interface{}{retentionModel}

	result, err := backuprecovery.ResourceIbmBaasProtectionGroupRunRequestMapToRunReplicationConfig(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasProtectionGroupRunRequestMapToRetention(t *testing.T) {
	checkResult := func(result *backuprecoveryv1.Retention) {
		dataLockConfigModel := new(backuprecoveryv1.DataLockConfig)
		dataLockConfigModel.Mode = core.StringPtr("Compliance")
		dataLockConfigModel.Unit = core.StringPtr("Days")
		dataLockConfigModel.Duration = core.Int64Ptr(int64(1))
		dataLockConfigModel.EnableWormOnExternalTarget = core.BoolPtr(true)

		model := new(backuprecoveryv1.Retention)
		model.Unit = core.StringPtr("Days")
		model.Duration = core.Int64Ptr(int64(1))
		model.DataLockConfig = dataLockConfigModel

		assert.Equal(t, result, model)
	}

	dataLockConfigModel := make(map[string]interface{})
	dataLockConfigModel["mode"] = "Compliance"
	dataLockConfigModel["unit"] = "Days"
	dataLockConfigModel["duration"] = int(1)
	dataLockConfigModel["enable_worm_on_external_target"] = true

	model := make(map[string]interface{})
	model["unit"] = "Days"
	model["duration"] = int(1)
	model["data_lock_config"] = []interface{}{dataLockConfigModel}

	result, err := backuprecovery.ResourceIbmBaasProtectionGroupRunRequestMapToRetention(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasProtectionGroupRunRequestMapToDataLockConfig(t *testing.T) {
	checkResult := func(result *backuprecoveryv1.DataLockConfig) {
		model := new(backuprecoveryv1.DataLockConfig)
		model.Mode = core.StringPtr("Compliance")
		model.Unit = core.StringPtr("Days")
		model.Duration = core.Int64Ptr(int64(1))
		model.EnableWormOnExternalTarget = core.BoolPtr(true)

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["mode"] = "Compliance"
	model["unit"] = "Days"
	model["duration"] = int(1)
	model["enable_worm_on_external_target"] = true

	result, err := backuprecovery.ResourceIbmBaasProtectionGroupRunRequestMapToDataLockConfig(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasProtectionGroupRunRequestMapToRunArchivalConfig(t *testing.T) {
	checkResult := func(result *backuprecoveryv1.RunArchivalConfig) {
		dataLockConfigModel := new(backuprecoveryv1.DataLockConfig)
		dataLockConfigModel.Mode = core.StringPtr("Compliance")
		dataLockConfigModel.Unit = core.StringPtr("Days")
		dataLockConfigModel.Duration = core.Int64Ptr(int64(1))
		dataLockConfigModel.EnableWormOnExternalTarget = core.BoolPtr(true)

		retentionModel := new(backuprecoveryv1.Retention)
		retentionModel.Unit = core.StringPtr("Days")
		retentionModel.Duration = core.Int64Ptr(int64(1))
		retentionModel.DataLockConfig = dataLockConfigModel

		model := new(backuprecoveryv1.RunArchivalConfig)
		model.ID = core.Int64Ptr(int64(26))
		model.ArchivalTargetType = core.StringPtr("Tape")
		model.Retention = retentionModel
		model.CopyOnlyFullySuccessful = core.BoolPtr(true)

		assert.Equal(t, result, model)
	}

	dataLockConfigModel := make(map[string]interface{})
	dataLockConfigModel["mode"] = "Compliance"
	dataLockConfigModel["unit"] = "Days"
	dataLockConfigModel["duration"] = int(1)
	dataLockConfigModel["enable_worm_on_external_target"] = true

	retentionModel := make(map[string]interface{})
	retentionModel["unit"] = "Days"
	retentionModel["duration"] = int(1)
	retentionModel["data_lock_config"] = []interface{}{dataLockConfigModel}

	model := make(map[string]interface{})
	model["id"] = int(26)
	model["archival_target_type"] = "Tape"
	model["retention"] = []interface{}{retentionModel}
	model["copy_only_fully_successful"] = true

	result, err := backuprecovery.ResourceIbmBaasProtectionGroupRunRequestMapToRunArchivalConfig(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasProtectionGroupRunRequestMapToRunCloudReplicationConfig(t *testing.T) {
	checkResult := func(result *backuprecoveryv1.RunCloudReplicationConfig) {
		awsTargetConfigModel := new(backuprecoveryv1.AWSTargetConfig)
		awsTargetConfigModel.Region = core.Int64Ptr(int64(26))
		awsTargetConfigModel.SourceID = core.Int64Ptr(int64(26))

		azureTargetConfigModel := new(backuprecoveryv1.AzureTargetConfig)
		azureTargetConfigModel.ResourceGroup = core.Int64Ptr(int64(26))
		azureTargetConfigModel.SourceID = core.Int64Ptr(int64(26))

		dataLockConfigModel := new(backuprecoveryv1.DataLockConfig)
		dataLockConfigModel.Mode = core.StringPtr("Compliance")
		dataLockConfigModel.Unit = core.StringPtr("Days")
		dataLockConfigModel.Duration = core.Int64Ptr(int64(1))
		dataLockConfigModel.EnableWormOnExternalTarget = core.BoolPtr(true)

		retentionModel := new(backuprecoveryv1.Retention)
		retentionModel.Unit = core.StringPtr("Days")
		retentionModel.Duration = core.Int64Ptr(int64(1))
		retentionModel.DataLockConfig = dataLockConfigModel

		model := new(backuprecoveryv1.RunCloudReplicationConfig)
		model.AwsTarget = awsTargetConfigModel
		model.AzureTarget = azureTargetConfigModel
		model.TargetType = core.StringPtr("AWS")
		model.Retention = retentionModel

		assert.Equal(t, result, model)
	}

	awsTargetConfigModel := make(map[string]interface{})
	awsTargetConfigModel["region"] = int(26)
	awsTargetConfigModel["source_id"] = int(26)

	azureTargetConfigModel := make(map[string]interface{})
	azureTargetConfigModel["resource_group"] = int(26)
	azureTargetConfigModel["source_id"] = int(26)

	dataLockConfigModel := make(map[string]interface{})
	dataLockConfigModel["mode"] = "Compliance"
	dataLockConfigModel["unit"] = "Days"
	dataLockConfigModel["duration"] = int(1)
	dataLockConfigModel["enable_worm_on_external_target"] = true

	retentionModel := make(map[string]interface{})
	retentionModel["unit"] = "Days"
	retentionModel["duration"] = int(1)
	retentionModel["data_lock_config"] = []interface{}{dataLockConfigModel}

	model := make(map[string]interface{})
	model["aws_target"] = []interface{}{awsTargetConfigModel}
	model["azure_target"] = []interface{}{azureTargetConfigModel}
	model["target_type"] = "AWS"
	model["retention"] = []interface{}{retentionModel}

	result, err := backuprecovery.ResourceIbmBaasProtectionGroupRunRequestMapToRunCloudReplicationConfig(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasProtectionGroupRunRequestMapToAWSTargetConfig(t *testing.T) {
	checkResult := func(result *backuprecoveryv1.AWSTargetConfig) {
		model := new(backuprecoveryv1.AWSTargetConfig)
		model.Name = core.StringPtr("testString")
		model.Region = core.Int64Ptr(int64(26))
		model.RegionName = core.StringPtr("testString")
		model.SourceID = core.Int64Ptr(int64(26))

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["name"] = "testString"
	model["region"] = int(26)
	model["region_name"] = "testString"
	model["source_id"] = int(26)

	result, err := backuprecovery.ResourceIbmBaasProtectionGroupRunRequestMapToAWSTargetConfig(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasProtectionGroupRunRequestMapToAzureTargetConfig(t *testing.T) {
	checkResult := func(result *backuprecoveryv1.AzureTargetConfig) {
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

		assert.Equal(t, result, model)
	}

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

	result, err := backuprecovery.ResourceIbmBaasProtectionGroupRunRequestMapToAzureTargetConfig(model)
	assert.Nil(t, err)
	checkResult(result)
}
