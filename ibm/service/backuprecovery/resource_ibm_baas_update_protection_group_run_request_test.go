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

func TestAccIbmBaasUpdateProtectionGroupRunRequestBasic(t *testing.T) {
	var conf backuprecoveryv1.ProtectionGroupRunsResponse
	xIbmTenantID := fmt.Sprintf("tf_x_ibm_tenant_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmBaasUpdateProtectionGroupRunRequestDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmBaasUpdateProtectionGroupRunRequestConfigBasic(xIbmTenantID),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmBaasUpdateProtectionGroupRunRequestExists("ibm_baas_update_protection_group_run_request.baas_update_protection_group_run_request_instance", conf),
					resource.TestCheckResourceAttr("ibm_baas_update_protection_group_run_request.baas_update_protection_group_run_request_instance", "x_ibm_tenant_id", xIbmTenantID),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_baas_update_protection_group_run_request.baas_update_protection_group_run_request",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmBaasUpdateProtectionGroupRunRequestConfigBasic(xIbmTenantID string) string {
	return fmt.Sprintf(`
		resource "ibm_baas_update_protection_group_run_request" "baas_update_protection_group_run_request_instance" {
			x_ibm_tenant_id = "%s"
			update_protection_group_run_params {
				run_id = "run_id"
				local_snapshot_config {
					enable_legal_hold = true
					delete_snapshot = true
					data_lock = "Compliance"
					days_to_keep = 1
				}
				replication_snapshot_config {
					new_snapshot_config {
						id = 1
						retention {
							unit = "Days"
							duration = 1
							data_lock_config {
								mode = "Compliance"
								unit = "Days"
								duration = 1
								enable_worm_on_external_target = true
							}
						}
					}
					update_existing_snapshot_config {
						id = 1
						name = "name"
						enable_legal_hold = true
						delete_snapshot = true
						resync = true
						data_lock = "Compliance"
						days_to_keep = 1
					}
				}
				archival_snapshot_config {
					new_snapshot_config {
						id = 1
						archival_target_type = "Tape"
						retention {
							unit = "Days"
							duration = 1
							data_lock_config {
								mode = "Compliance"
								unit = "Days"
								duration = 1
								enable_worm_on_external_target = true
							}
						}
						copy_only_fully_successful = true
					}
					update_existing_snapshot_config {
						id = 1
						name = "name"
						archival_target_type = "Tape"
						enable_legal_hold = true
						delete_snapshot = true
						resync = true
						data_lock = "Compliance"
						days_to_keep = 1
					}
				}
			}
		}
	`, xIbmTenantID)
}

func testAccCheckIbmBaasUpdateProtectionGroupRunRequestExists(n string, obj backuprecoveryv1.ProtectionGroupRunsResponse) resource.TestCheckFunc {

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

		updateProtectionGroupRunRequest, _, err := backupRecoveryClient.GetProtectionGroupRuns(getProtectionGroupRunsOptions)
		if err != nil {
			return err
		}

		obj = *updateProtectionGroupRunRequest
		return nil
	}
}

func testAccCheckIbmBaasUpdateProtectionGroupRunRequestDestroy(s *terraform.State) error {
	backupRecoveryClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).BackupRecoveryV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_baas_update_protection_group_run_request" {
			continue
		}

		getProtectionGroupRunsOptions := &backuprecoveryv1.GetProtectionGroupRunsOptions{}

		getProtectionGroupRunsOptions.SetID(rs.Primary.ID)

		// Try to find the key
		_, response, err := backupRecoveryClient.GetProtectionGroupRuns(getProtectionGroupRunsOptions)

		if err == nil {
			return fmt.Errorf("Update Protection Group Run Request Body. still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for Update Protection Group Run Request Body. (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

func TestResourceIbmBaasUpdateProtectionGroupRunRequestUpdateProtectionGroupRunParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		updateLocalSnapshotConfigModel := make(map[string]interface{})
		updateLocalSnapshotConfigModel["enable_legal_hold"] = true
		updateLocalSnapshotConfigModel["delete_snapshot"] = true
		updateLocalSnapshotConfigModel["data_lock"] = "Compliance"
		updateLocalSnapshotConfigModel["days_to_keep"] = int(26)

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

		updateExistingReplicationSnapshotConfigModel := make(map[string]interface{})
		updateExistingReplicationSnapshotConfigModel["id"] = int(26)
		updateExistingReplicationSnapshotConfigModel["name"] = "testString"
		updateExistingReplicationSnapshotConfigModel["enable_legal_hold"] = true
		updateExistingReplicationSnapshotConfigModel["delete_snapshot"] = true
		updateExistingReplicationSnapshotConfigModel["resync"] = true
		updateExistingReplicationSnapshotConfigModel["data_lock"] = "Compliance"
		updateExistingReplicationSnapshotConfigModel["days_to_keep"] = int(26)

		updateReplicationSnapshotConfigModel := make(map[string]interface{})
		updateReplicationSnapshotConfigModel["new_snapshot_config"] = []map[string]interface{}{runReplicationConfigModel}
		updateReplicationSnapshotConfigModel["update_existing_snapshot_config"] = []map[string]interface{}{updateExistingReplicationSnapshotConfigModel}

		runArchivalConfigModel := make(map[string]interface{})
		runArchivalConfigModel["id"] = int(26)
		runArchivalConfigModel["archival_target_type"] = "Tape"
		runArchivalConfigModel["retention"] = []map[string]interface{}{retentionModel}
		runArchivalConfigModel["copy_only_fully_successful"] = true

		updateExistingArchivalSnapshotConfigModel := make(map[string]interface{})
		updateExistingArchivalSnapshotConfigModel["id"] = int(26)
		updateExistingArchivalSnapshotConfigModel["name"] = "testString"
		updateExistingArchivalSnapshotConfigModel["archival_target_type"] = "Tape"
		updateExistingArchivalSnapshotConfigModel["enable_legal_hold"] = true
		updateExistingArchivalSnapshotConfigModel["delete_snapshot"] = true
		updateExistingArchivalSnapshotConfigModel["resync"] = true
		updateExistingArchivalSnapshotConfigModel["data_lock"] = "Compliance"
		updateExistingArchivalSnapshotConfigModel["days_to_keep"] = int(26)

		updateArchivalSnapshotConfigModel := make(map[string]interface{})
		updateArchivalSnapshotConfigModel["new_snapshot_config"] = []map[string]interface{}{runArchivalConfigModel}
		updateArchivalSnapshotConfigModel["update_existing_snapshot_config"] = []map[string]interface{}{updateExistingArchivalSnapshotConfigModel}

		model := make(map[string]interface{})
		model["run_id"] = "testString"
		model["local_snapshot_config"] = []map[string]interface{}{updateLocalSnapshotConfigModel}
		model["replication_snapshot_config"] = []map[string]interface{}{updateReplicationSnapshotConfigModel}
		model["archival_snapshot_config"] = []map[string]interface{}{updateArchivalSnapshotConfigModel}

		assert.Equal(t, result, model)
	}

	updateLocalSnapshotConfigModel := new(backuprecoveryv1.UpdateLocalSnapshotConfig)
	updateLocalSnapshotConfigModel.EnableLegalHold = core.BoolPtr(true)
	updateLocalSnapshotConfigModel.DeleteSnapshot = core.BoolPtr(true)
	updateLocalSnapshotConfigModel.DataLock = core.StringPtr("Compliance")
	updateLocalSnapshotConfigModel.DaysToKeep = core.Int64Ptr(int64(26))

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

	updateExistingReplicationSnapshotConfigModel := new(backuprecoveryv1.UpdateExistingReplicationSnapshotConfig)
	updateExistingReplicationSnapshotConfigModel.ID = core.Int64Ptr(int64(26))
	updateExistingReplicationSnapshotConfigModel.Name = core.StringPtr("testString")
	updateExistingReplicationSnapshotConfigModel.EnableLegalHold = core.BoolPtr(true)
	updateExistingReplicationSnapshotConfigModel.DeleteSnapshot = core.BoolPtr(true)
	updateExistingReplicationSnapshotConfigModel.Resync = core.BoolPtr(true)
	updateExistingReplicationSnapshotConfigModel.DataLock = core.StringPtr("Compliance")
	updateExistingReplicationSnapshotConfigModel.DaysToKeep = core.Int64Ptr(int64(26))

	updateReplicationSnapshotConfigModel := new(backuprecoveryv1.UpdateReplicationSnapshotConfig)
	updateReplicationSnapshotConfigModel.NewSnapshotConfig = []backuprecoveryv1.RunReplicationConfig{*runReplicationConfigModel}
	updateReplicationSnapshotConfigModel.UpdateExistingSnapshotConfig = []backuprecoveryv1.UpdateExistingReplicationSnapshotConfig{*updateExistingReplicationSnapshotConfigModel}

	runArchivalConfigModel := new(backuprecoveryv1.RunArchivalConfig)
	runArchivalConfigModel.ID = core.Int64Ptr(int64(26))
	runArchivalConfigModel.ArchivalTargetType = core.StringPtr("Tape")
	runArchivalConfigModel.Retention = retentionModel
	runArchivalConfigModel.CopyOnlyFullySuccessful = core.BoolPtr(true)

	updateExistingArchivalSnapshotConfigModel := new(backuprecoveryv1.UpdateExistingArchivalSnapshotConfig)
	updateExistingArchivalSnapshotConfigModel.ID = core.Int64Ptr(int64(26))
	updateExistingArchivalSnapshotConfigModel.Name = core.StringPtr("testString")
	updateExistingArchivalSnapshotConfigModel.ArchivalTargetType = core.StringPtr("Tape")
	updateExistingArchivalSnapshotConfigModel.EnableLegalHold = core.BoolPtr(true)
	updateExistingArchivalSnapshotConfigModel.DeleteSnapshot = core.BoolPtr(true)
	updateExistingArchivalSnapshotConfigModel.Resync = core.BoolPtr(true)
	updateExistingArchivalSnapshotConfigModel.DataLock = core.StringPtr("Compliance")
	updateExistingArchivalSnapshotConfigModel.DaysToKeep = core.Int64Ptr(int64(26))

	updateArchivalSnapshotConfigModel := new(backuprecoveryv1.UpdateArchivalSnapshotConfig)
	updateArchivalSnapshotConfigModel.NewSnapshotConfig = []backuprecoveryv1.RunArchivalConfig{*runArchivalConfigModel}
	updateArchivalSnapshotConfigModel.UpdateExistingSnapshotConfig = []backuprecoveryv1.UpdateExistingArchivalSnapshotConfig{*updateExistingArchivalSnapshotConfigModel}

	model := new(backuprecoveryv1.UpdateProtectionGroupRunParams)
	model.RunID = core.StringPtr("testString")
	model.LocalSnapshotConfig = updateLocalSnapshotConfigModel
	model.ReplicationSnapshotConfig = updateReplicationSnapshotConfigModel
	model.ArchivalSnapshotConfig = updateArchivalSnapshotConfigModel

	result, err := backuprecovery.ResourceIbmBaasUpdateProtectionGroupRunRequestUpdateProtectionGroupRunParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasUpdateProtectionGroupRunRequestUpdateLocalSnapshotConfigToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["enable_legal_hold"] = true
		model["delete_snapshot"] = true
		model["data_lock"] = "Compliance"
		model["days_to_keep"] = int(26)

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.UpdateLocalSnapshotConfig)
	model.EnableLegalHold = core.BoolPtr(true)
	model.DeleteSnapshot = core.BoolPtr(true)
	model.DataLock = core.StringPtr("Compliance")
	model.DaysToKeep = core.Int64Ptr(int64(26))

	result, err := backuprecovery.ResourceIbmBaasUpdateProtectionGroupRunRequestUpdateLocalSnapshotConfigToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasUpdateProtectionGroupRunRequestUpdateReplicationSnapshotConfigToMap(t *testing.T) {
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

		updateExistingReplicationSnapshotConfigModel := make(map[string]interface{})
		updateExistingReplicationSnapshotConfigModel["id"] = int(26)
		updateExistingReplicationSnapshotConfigModel["name"] = "testString"
		updateExistingReplicationSnapshotConfigModel["enable_legal_hold"] = true
		updateExistingReplicationSnapshotConfigModel["delete_snapshot"] = true
		updateExistingReplicationSnapshotConfigModel["resync"] = true
		updateExistingReplicationSnapshotConfigModel["data_lock"] = "Compliance"
		updateExistingReplicationSnapshotConfigModel["days_to_keep"] = int(26)

		model := make(map[string]interface{})
		model["new_snapshot_config"] = []map[string]interface{}{runReplicationConfigModel}
		model["update_existing_snapshot_config"] = []map[string]interface{}{updateExistingReplicationSnapshotConfigModel}

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

	updateExistingReplicationSnapshotConfigModel := new(backuprecoveryv1.UpdateExistingReplicationSnapshotConfig)
	updateExistingReplicationSnapshotConfigModel.ID = core.Int64Ptr(int64(26))
	updateExistingReplicationSnapshotConfigModel.Name = core.StringPtr("testString")
	updateExistingReplicationSnapshotConfigModel.EnableLegalHold = core.BoolPtr(true)
	updateExistingReplicationSnapshotConfigModel.DeleteSnapshot = core.BoolPtr(true)
	updateExistingReplicationSnapshotConfigModel.Resync = core.BoolPtr(true)
	updateExistingReplicationSnapshotConfigModel.DataLock = core.StringPtr("Compliance")
	updateExistingReplicationSnapshotConfigModel.DaysToKeep = core.Int64Ptr(int64(26))

	model := new(backuprecoveryv1.UpdateReplicationSnapshotConfig)
	model.NewSnapshotConfig = []backuprecoveryv1.RunReplicationConfig{*runReplicationConfigModel}
	model.UpdateExistingSnapshotConfig = []backuprecoveryv1.UpdateExistingReplicationSnapshotConfig{*updateExistingReplicationSnapshotConfigModel}

	result, err := backuprecovery.ResourceIbmBaasUpdateProtectionGroupRunRequestUpdateReplicationSnapshotConfigToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasUpdateProtectionGroupRunRequestRunReplicationConfigToMap(t *testing.T) {
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

	result, err := backuprecovery.ResourceIbmBaasUpdateProtectionGroupRunRequestRunReplicationConfigToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasUpdateProtectionGroupRunRequestRetentionToMap(t *testing.T) {
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

	result, err := backuprecovery.ResourceIbmBaasUpdateProtectionGroupRunRequestRetentionToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasUpdateProtectionGroupRunRequestDataLockConfigToMap(t *testing.T) {
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

	result, err := backuprecovery.ResourceIbmBaasUpdateProtectionGroupRunRequestDataLockConfigToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasUpdateProtectionGroupRunRequestUpdateExistingReplicationSnapshotConfigToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["id"] = int(26)
		model["name"] = "testString"
		model["enable_legal_hold"] = true
		model["delete_snapshot"] = true
		model["resync"] = true
		model["data_lock"] = "Compliance"
		model["days_to_keep"] = int(26)

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.UpdateExistingReplicationSnapshotConfig)
	model.ID = core.Int64Ptr(int64(26))
	model.Name = core.StringPtr("testString")
	model.EnableLegalHold = core.BoolPtr(true)
	model.DeleteSnapshot = core.BoolPtr(true)
	model.Resync = core.BoolPtr(true)
	model.DataLock = core.StringPtr("Compliance")
	model.DaysToKeep = core.Int64Ptr(int64(26))

	result, err := backuprecovery.ResourceIbmBaasUpdateProtectionGroupRunRequestUpdateExistingReplicationSnapshotConfigToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasUpdateProtectionGroupRunRequestUpdateArchivalSnapshotConfigToMap(t *testing.T) {
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

		runArchivalConfigModel := make(map[string]interface{})
		runArchivalConfigModel["id"] = int(26)
		runArchivalConfigModel["archival_target_type"] = "Tape"
		runArchivalConfigModel["retention"] = []map[string]interface{}{retentionModel}
		runArchivalConfigModel["copy_only_fully_successful"] = true

		updateExistingArchivalSnapshotConfigModel := make(map[string]interface{})
		updateExistingArchivalSnapshotConfigModel["id"] = int(26)
		updateExistingArchivalSnapshotConfigModel["name"] = "testString"
		updateExistingArchivalSnapshotConfigModel["archival_target_type"] = "Tape"
		updateExistingArchivalSnapshotConfigModel["enable_legal_hold"] = true
		updateExistingArchivalSnapshotConfigModel["delete_snapshot"] = true
		updateExistingArchivalSnapshotConfigModel["resync"] = true
		updateExistingArchivalSnapshotConfigModel["data_lock"] = "Compliance"
		updateExistingArchivalSnapshotConfigModel["days_to_keep"] = int(26)

		model := make(map[string]interface{})
		model["new_snapshot_config"] = []map[string]interface{}{runArchivalConfigModel}
		model["update_existing_snapshot_config"] = []map[string]interface{}{updateExistingArchivalSnapshotConfigModel}

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

	runArchivalConfigModel := new(backuprecoveryv1.RunArchivalConfig)
	runArchivalConfigModel.ID = core.Int64Ptr(int64(26))
	runArchivalConfigModel.ArchivalTargetType = core.StringPtr("Tape")
	runArchivalConfigModel.Retention = retentionModel
	runArchivalConfigModel.CopyOnlyFullySuccessful = core.BoolPtr(true)

	updateExistingArchivalSnapshotConfigModel := new(backuprecoveryv1.UpdateExistingArchivalSnapshotConfig)
	updateExistingArchivalSnapshotConfigModel.ID = core.Int64Ptr(int64(26))
	updateExistingArchivalSnapshotConfigModel.Name = core.StringPtr("testString")
	updateExistingArchivalSnapshotConfigModel.ArchivalTargetType = core.StringPtr("Tape")
	updateExistingArchivalSnapshotConfigModel.EnableLegalHold = core.BoolPtr(true)
	updateExistingArchivalSnapshotConfigModel.DeleteSnapshot = core.BoolPtr(true)
	updateExistingArchivalSnapshotConfigModel.Resync = core.BoolPtr(true)
	updateExistingArchivalSnapshotConfigModel.DataLock = core.StringPtr("Compliance")
	updateExistingArchivalSnapshotConfigModel.DaysToKeep = core.Int64Ptr(int64(26))

	model := new(backuprecoveryv1.UpdateArchivalSnapshotConfig)
	model.NewSnapshotConfig = []backuprecoveryv1.RunArchivalConfig{*runArchivalConfigModel}
	model.UpdateExistingSnapshotConfig = []backuprecoveryv1.UpdateExistingArchivalSnapshotConfig{*updateExistingArchivalSnapshotConfigModel}

	result, err := backuprecovery.ResourceIbmBaasUpdateProtectionGroupRunRequestUpdateArchivalSnapshotConfigToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasUpdateProtectionGroupRunRequestRunArchivalConfigToMap(t *testing.T) {
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

	result, err := backuprecovery.ResourceIbmBaasUpdateProtectionGroupRunRequestRunArchivalConfigToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasUpdateProtectionGroupRunRequestUpdateExistingArchivalSnapshotConfigToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["id"] = int(26)
		model["name"] = "testString"
		model["archival_target_type"] = "Tape"
		model["enable_legal_hold"] = true
		model["delete_snapshot"] = true
		model["resync"] = true
		model["data_lock"] = "Compliance"
		model["days_to_keep"] = int(26)

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.UpdateExistingArchivalSnapshotConfig)
	model.ID = core.Int64Ptr(int64(26))
	model.Name = core.StringPtr("testString")
	model.ArchivalTargetType = core.StringPtr("Tape")
	model.EnableLegalHold = core.BoolPtr(true)
	model.DeleteSnapshot = core.BoolPtr(true)
	model.Resync = core.BoolPtr(true)
	model.DataLock = core.StringPtr("Compliance")
	model.DaysToKeep = core.Int64Ptr(int64(26))

	result, err := backuprecovery.ResourceIbmBaasUpdateProtectionGroupRunRequestUpdateExistingArchivalSnapshotConfigToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasUpdateProtectionGroupRunRequestMapToUpdateProtectionGroupRunParams(t *testing.T) {
	checkResult := func(result *backuprecoveryv1.UpdateProtectionGroupRunParams) {
		updateLocalSnapshotConfigModel := new(backuprecoveryv1.UpdateLocalSnapshotConfig)
		updateLocalSnapshotConfigModel.EnableLegalHold = core.BoolPtr(true)
		updateLocalSnapshotConfigModel.DeleteSnapshot = core.BoolPtr(true)
		updateLocalSnapshotConfigModel.DataLock = core.StringPtr("Compliance")
		updateLocalSnapshotConfigModel.DaysToKeep = core.Int64Ptr(int64(26))

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

		updateExistingReplicationSnapshotConfigModel := new(backuprecoveryv1.UpdateExistingReplicationSnapshotConfig)
		updateExistingReplicationSnapshotConfigModel.ID = core.Int64Ptr(int64(26))
		updateExistingReplicationSnapshotConfigModel.Name = core.StringPtr("testString")
		updateExistingReplicationSnapshotConfigModel.EnableLegalHold = core.BoolPtr(true)
		updateExistingReplicationSnapshotConfigModel.DeleteSnapshot = core.BoolPtr(true)
		updateExistingReplicationSnapshotConfigModel.Resync = core.BoolPtr(true)
		updateExistingReplicationSnapshotConfigModel.DataLock = core.StringPtr("Compliance")
		updateExistingReplicationSnapshotConfigModel.DaysToKeep = core.Int64Ptr(int64(26))

		updateReplicationSnapshotConfigModel := new(backuprecoveryv1.UpdateReplicationSnapshotConfig)
		updateReplicationSnapshotConfigModel.NewSnapshotConfig = []backuprecoveryv1.RunReplicationConfig{*runReplicationConfigModel}
		updateReplicationSnapshotConfigModel.UpdateExistingSnapshotConfig = []backuprecoveryv1.UpdateExistingReplicationSnapshotConfig{*updateExistingReplicationSnapshotConfigModel}

		runArchivalConfigModel := new(backuprecoveryv1.RunArchivalConfig)
		runArchivalConfigModel.ID = core.Int64Ptr(int64(26))
		runArchivalConfigModel.ArchivalTargetType = core.StringPtr("Tape")
		runArchivalConfigModel.Retention = retentionModel
		runArchivalConfigModel.CopyOnlyFullySuccessful = core.BoolPtr(true)

		updateExistingArchivalSnapshotConfigModel := new(backuprecoveryv1.UpdateExistingArchivalSnapshotConfig)
		updateExistingArchivalSnapshotConfigModel.ID = core.Int64Ptr(int64(26))
		updateExistingArchivalSnapshotConfigModel.Name = core.StringPtr("testString")
		updateExistingArchivalSnapshotConfigModel.ArchivalTargetType = core.StringPtr("Tape")
		updateExistingArchivalSnapshotConfigModel.EnableLegalHold = core.BoolPtr(true)
		updateExistingArchivalSnapshotConfigModel.DeleteSnapshot = core.BoolPtr(true)
		updateExistingArchivalSnapshotConfigModel.Resync = core.BoolPtr(true)
		updateExistingArchivalSnapshotConfigModel.DataLock = core.StringPtr("Compliance")
		updateExistingArchivalSnapshotConfigModel.DaysToKeep = core.Int64Ptr(int64(26))

		updateArchivalSnapshotConfigModel := new(backuprecoveryv1.UpdateArchivalSnapshotConfig)
		updateArchivalSnapshotConfigModel.NewSnapshotConfig = []backuprecoveryv1.RunArchivalConfig{*runArchivalConfigModel}
		updateArchivalSnapshotConfigModel.UpdateExistingSnapshotConfig = []backuprecoveryv1.UpdateExistingArchivalSnapshotConfig{*updateExistingArchivalSnapshotConfigModel}

		model := new(backuprecoveryv1.UpdateProtectionGroupRunParams)
		model.RunID = core.StringPtr("testString")
		model.LocalSnapshotConfig = updateLocalSnapshotConfigModel
		model.ReplicationSnapshotConfig = updateReplicationSnapshotConfigModel
		model.ArchivalSnapshotConfig = updateArchivalSnapshotConfigModel

		assert.Equal(t, result, model)
	}

	updateLocalSnapshotConfigModel := make(map[string]interface{})
	updateLocalSnapshotConfigModel["enable_legal_hold"] = true
	updateLocalSnapshotConfigModel["delete_snapshot"] = true
	updateLocalSnapshotConfigModel["data_lock"] = "Compliance"
	updateLocalSnapshotConfigModel["days_to_keep"] = int(26)

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

	updateExistingReplicationSnapshotConfigModel := make(map[string]interface{})
	updateExistingReplicationSnapshotConfigModel["id"] = int(26)
	updateExistingReplicationSnapshotConfigModel["name"] = "testString"
	updateExistingReplicationSnapshotConfigModel["enable_legal_hold"] = true
	updateExistingReplicationSnapshotConfigModel["delete_snapshot"] = true
	updateExistingReplicationSnapshotConfigModel["resync"] = true
	updateExistingReplicationSnapshotConfigModel["data_lock"] = "Compliance"
	updateExistingReplicationSnapshotConfigModel["days_to_keep"] = int(26)

	updateReplicationSnapshotConfigModel := make(map[string]interface{})
	updateReplicationSnapshotConfigModel["new_snapshot_config"] = []interface{}{runReplicationConfigModel}
	updateReplicationSnapshotConfigModel["update_existing_snapshot_config"] = []interface{}{updateExistingReplicationSnapshotConfigModel}

	runArchivalConfigModel := make(map[string]interface{})
	runArchivalConfigModel["id"] = int(26)
	runArchivalConfigModel["archival_target_type"] = "Tape"
	runArchivalConfigModel["retention"] = []interface{}{retentionModel}
	runArchivalConfigModel["copy_only_fully_successful"] = true

	updateExistingArchivalSnapshotConfigModel := make(map[string]interface{})
	updateExistingArchivalSnapshotConfigModel["id"] = int(26)
	updateExistingArchivalSnapshotConfigModel["name"] = "testString"
	updateExistingArchivalSnapshotConfigModel["archival_target_type"] = "Tape"
	updateExistingArchivalSnapshotConfigModel["enable_legal_hold"] = true
	updateExistingArchivalSnapshotConfigModel["delete_snapshot"] = true
	updateExistingArchivalSnapshotConfigModel["resync"] = true
	updateExistingArchivalSnapshotConfigModel["data_lock"] = "Compliance"
	updateExistingArchivalSnapshotConfigModel["days_to_keep"] = int(26)

	updateArchivalSnapshotConfigModel := make(map[string]interface{})
	updateArchivalSnapshotConfigModel["new_snapshot_config"] = []interface{}{runArchivalConfigModel}
	updateArchivalSnapshotConfigModel["update_existing_snapshot_config"] = []interface{}{updateExistingArchivalSnapshotConfigModel}

	model := make(map[string]interface{})
	model["run_id"] = "testString"
	model["local_snapshot_config"] = []interface{}{updateLocalSnapshotConfigModel}
	model["replication_snapshot_config"] = []interface{}{updateReplicationSnapshotConfigModel}
	model["archival_snapshot_config"] = []interface{}{updateArchivalSnapshotConfigModel}

	result, err := backuprecovery.ResourceIbmBaasUpdateProtectionGroupRunRequestMapToUpdateProtectionGroupRunParams(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasUpdateProtectionGroupRunRequestMapToUpdateLocalSnapshotConfig(t *testing.T) {
	checkResult := func(result *backuprecoveryv1.UpdateLocalSnapshotConfig) {
		model := new(backuprecoveryv1.UpdateLocalSnapshotConfig)
		model.EnableLegalHold = core.BoolPtr(true)
		model.DeleteSnapshot = core.BoolPtr(true)
		model.DataLock = core.StringPtr("Compliance")
		model.DaysToKeep = core.Int64Ptr(int64(26))

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["enable_legal_hold"] = true
	model["delete_snapshot"] = true
	model["data_lock"] = "Compliance"
	model["days_to_keep"] = int(26)

	result, err := backuprecovery.ResourceIbmBaasUpdateProtectionGroupRunRequestMapToUpdateLocalSnapshotConfig(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasUpdateProtectionGroupRunRequestMapToUpdateReplicationSnapshotConfig(t *testing.T) {
	checkResult := func(result *backuprecoveryv1.UpdateReplicationSnapshotConfig) {
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

		updateExistingReplicationSnapshotConfigModel := new(backuprecoveryv1.UpdateExistingReplicationSnapshotConfig)
		updateExistingReplicationSnapshotConfigModel.ID = core.Int64Ptr(int64(26))
		updateExistingReplicationSnapshotConfigModel.Name = core.StringPtr("testString")
		updateExistingReplicationSnapshotConfigModel.EnableLegalHold = core.BoolPtr(true)
		updateExistingReplicationSnapshotConfigModel.DeleteSnapshot = core.BoolPtr(true)
		updateExistingReplicationSnapshotConfigModel.Resync = core.BoolPtr(true)
		updateExistingReplicationSnapshotConfigModel.DataLock = core.StringPtr("Compliance")
		updateExistingReplicationSnapshotConfigModel.DaysToKeep = core.Int64Ptr(int64(26))

		model := new(backuprecoveryv1.UpdateReplicationSnapshotConfig)
		model.NewSnapshotConfig = []backuprecoveryv1.RunReplicationConfig{*runReplicationConfigModel}
		model.UpdateExistingSnapshotConfig = []backuprecoveryv1.UpdateExistingReplicationSnapshotConfig{*updateExistingReplicationSnapshotConfigModel}

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

	updateExistingReplicationSnapshotConfigModel := make(map[string]interface{})
	updateExistingReplicationSnapshotConfigModel["id"] = int(26)
	updateExistingReplicationSnapshotConfigModel["name"] = "testString"
	updateExistingReplicationSnapshotConfigModel["enable_legal_hold"] = true
	updateExistingReplicationSnapshotConfigModel["delete_snapshot"] = true
	updateExistingReplicationSnapshotConfigModel["resync"] = true
	updateExistingReplicationSnapshotConfigModel["data_lock"] = "Compliance"
	updateExistingReplicationSnapshotConfigModel["days_to_keep"] = int(26)

	model := make(map[string]interface{})
	model["new_snapshot_config"] = []interface{}{runReplicationConfigModel}
	model["update_existing_snapshot_config"] = []interface{}{updateExistingReplicationSnapshotConfigModel}

	result, err := backuprecovery.ResourceIbmBaasUpdateProtectionGroupRunRequestMapToUpdateReplicationSnapshotConfig(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasUpdateProtectionGroupRunRequestMapToRunReplicationConfig(t *testing.T) {
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

	result, err := backuprecovery.ResourceIbmBaasUpdateProtectionGroupRunRequestMapToRunReplicationConfig(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasUpdateProtectionGroupRunRequestMapToRetention(t *testing.T) {
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

	result, err := backuprecovery.ResourceIbmBaasUpdateProtectionGroupRunRequestMapToRetention(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasUpdateProtectionGroupRunRequestMapToDataLockConfig(t *testing.T) {
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

	result, err := backuprecovery.ResourceIbmBaasUpdateProtectionGroupRunRequestMapToDataLockConfig(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasUpdateProtectionGroupRunRequestMapToUpdateExistingReplicationSnapshotConfig(t *testing.T) {
	checkResult := func(result *backuprecoveryv1.UpdateExistingReplicationSnapshotConfig) {
		model := new(backuprecoveryv1.UpdateExistingReplicationSnapshotConfig)
		model.ID = core.Int64Ptr(int64(26))
		model.Name = core.StringPtr("testString")
		model.EnableLegalHold = core.BoolPtr(true)
		model.DeleteSnapshot = core.BoolPtr(true)
		model.Resync = core.BoolPtr(true)
		model.DataLock = core.StringPtr("Compliance")
		model.DaysToKeep = core.Int64Ptr(int64(26))

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["id"] = int(26)
	model["name"] = "testString"
	model["enable_legal_hold"] = true
	model["delete_snapshot"] = true
	model["resync"] = true
	model["data_lock"] = "Compliance"
	model["days_to_keep"] = int(26)

	result, err := backuprecovery.ResourceIbmBaasUpdateProtectionGroupRunRequestMapToUpdateExistingReplicationSnapshotConfig(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasUpdateProtectionGroupRunRequestMapToUpdateArchivalSnapshotConfig(t *testing.T) {
	checkResult := func(result *backuprecoveryv1.UpdateArchivalSnapshotConfig) {
		dataLockConfigModel := new(backuprecoveryv1.DataLockConfig)
		dataLockConfigModel.Mode = core.StringPtr("Compliance")
		dataLockConfigModel.Unit = core.StringPtr("Days")
		dataLockConfigModel.Duration = core.Int64Ptr(int64(1))
		dataLockConfigModel.EnableWormOnExternalTarget = core.BoolPtr(true)

		retentionModel := new(backuprecoveryv1.Retention)
		retentionModel.Unit = core.StringPtr("Days")
		retentionModel.Duration = core.Int64Ptr(int64(1))
		retentionModel.DataLockConfig = dataLockConfigModel

		runArchivalConfigModel := new(backuprecoveryv1.RunArchivalConfig)
		runArchivalConfigModel.ID = core.Int64Ptr(int64(26))
		runArchivalConfigModel.ArchivalTargetType = core.StringPtr("Tape")
		runArchivalConfigModel.Retention = retentionModel
		runArchivalConfigModel.CopyOnlyFullySuccessful = core.BoolPtr(true)

		updateExistingArchivalSnapshotConfigModel := new(backuprecoveryv1.UpdateExistingArchivalSnapshotConfig)
		updateExistingArchivalSnapshotConfigModel.ID = core.Int64Ptr(int64(26))
		updateExistingArchivalSnapshotConfigModel.Name = core.StringPtr("testString")
		updateExistingArchivalSnapshotConfigModel.ArchivalTargetType = core.StringPtr("Tape")
		updateExistingArchivalSnapshotConfigModel.EnableLegalHold = core.BoolPtr(true)
		updateExistingArchivalSnapshotConfigModel.DeleteSnapshot = core.BoolPtr(true)
		updateExistingArchivalSnapshotConfigModel.Resync = core.BoolPtr(true)
		updateExistingArchivalSnapshotConfigModel.DataLock = core.StringPtr("Compliance")
		updateExistingArchivalSnapshotConfigModel.DaysToKeep = core.Int64Ptr(int64(26))

		model := new(backuprecoveryv1.UpdateArchivalSnapshotConfig)
		model.NewSnapshotConfig = []backuprecoveryv1.RunArchivalConfig{*runArchivalConfigModel}
		model.UpdateExistingSnapshotConfig = []backuprecoveryv1.UpdateExistingArchivalSnapshotConfig{*updateExistingArchivalSnapshotConfigModel}

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

	runArchivalConfigModel := make(map[string]interface{})
	runArchivalConfigModel["id"] = int(26)
	runArchivalConfigModel["archival_target_type"] = "Tape"
	runArchivalConfigModel["retention"] = []interface{}{retentionModel}
	runArchivalConfigModel["copy_only_fully_successful"] = true

	updateExistingArchivalSnapshotConfigModel := make(map[string]interface{})
	updateExistingArchivalSnapshotConfigModel["id"] = int(26)
	updateExistingArchivalSnapshotConfigModel["name"] = "testString"
	updateExistingArchivalSnapshotConfigModel["archival_target_type"] = "Tape"
	updateExistingArchivalSnapshotConfigModel["enable_legal_hold"] = true
	updateExistingArchivalSnapshotConfigModel["delete_snapshot"] = true
	updateExistingArchivalSnapshotConfigModel["resync"] = true
	updateExistingArchivalSnapshotConfigModel["data_lock"] = "Compliance"
	updateExistingArchivalSnapshotConfigModel["days_to_keep"] = int(26)

	model := make(map[string]interface{})
	model["new_snapshot_config"] = []interface{}{runArchivalConfigModel}
	model["update_existing_snapshot_config"] = []interface{}{updateExistingArchivalSnapshotConfigModel}

	result, err := backuprecovery.ResourceIbmBaasUpdateProtectionGroupRunRequestMapToUpdateArchivalSnapshotConfig(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasUpdateProtectionGroupRunRequestMapToRunArchivalConfig(t *testing.T) {
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

	result, err := backuprecovery.ResourceIbmBaasUpdateProtectionGroupRunRequestMapToRunArchivalConfig(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasUpdateProtectionGroupRunRequestMapToUpdateExistingArchivalSnapshotConfig(t *testing.T) {
	checkResult := func(result *backuprecoveryv1.UpdateExistingArchivalSnapshotConfig) {
		model := new(backuprecoveryv1.UpdateExistingArchivalSnapshotConfig)
		model.ID = core.Int64Ptr(int64(26))
		model.Name = core.StringPtr("testString")
		model.ArchivalTargetType = core.StringPtr("Tape")
		model.EnableLegalHold = core.BoolPtr(true)
		model.DeleteSnapshot = core.BoolPtr(true)
		model.Resync = core.BoolPtr(true)
		model.DataLock = core.StringPtr("Compliance")
		model.DaysToKeep = core.Int64Ptr(int64(26))

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["id"] = int(26)
	model["name"] = "testString"
	model["archival_target_type"] = "Tape"
	model["enable_legal_hold"] = true
	model["delete_snapshot"] = true
	model["resync"] = true
	model["data_lock"] = "Compliance"
	model["days_to_keep"] = int(26)

	result, err := backuprecovery.ResourceIbmBaasUpdateProtectionGroupRunRequestMapToUpdateExistingArchivalSnapshotConfig(model)
	assert.Nil(t, err)
	checkResult(result)
}
