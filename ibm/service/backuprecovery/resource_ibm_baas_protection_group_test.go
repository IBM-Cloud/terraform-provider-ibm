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

func TestAccIbmBaasProtectionGroupBasic(t *testing.T) {
	var conf backuprecoveryv1.ProtectionGroupResponse
	xIbmTenantID := fmt.Sprintf("tf_x_ibm_tenant_id_%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	policyID := fmt.Sprintf("tf_policy_id_%d", acctest.RandIntRange(10, 100))
	environment := "kPhysical"
	xIbmTenantIDUpdate := fmt.Sprintf("tf_x_ibm_tenant_id_%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	policyIDUpdate := fmt.Sprintf("tf_policy_id_%d", acctest.RandIntRange(10, 100))
	environmentUpdate := "kSQL"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmBaasProtectionGroupDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmBaasProtectionGroupConfigBasic(xIbmTenantID, name, policyID, environment),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmBaasProtectionGroupExists("ibm_baas_protection_group.baas_protection_group_instance", conf),
					resource.TestCheckResourceAttr("ibm_baas_protection_group.baas_protection_group_instance", "x_ibm_tenant_id", xIbmTenantID),
					resource.TestCheckResourceAttr("ibm_baas_protection_group.baas_protection_group_instance", "name", name),
					resource.TestCheckResourceAttr("ibm_baas_protection_group.baas_protection_group_instance", "policy_id", policyID),
					resource.TestCheckResourceAttr("ibm_baas_protection_group.baas_protection_group_instance", "environment", environment),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmBaasProtectionGroupConfigBasic(xIbmTenantIDUpdate, nameUpdate, policyIDUpdate, environmentUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_baas_protection_group.baas_protection_group_instance", "x_ibm_tenant_id", xIbmTenantIDUpdate),
					resource.TestCheckResourceAttr("ibm_baas_protection_group.baas_protection_group_instance", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_baas_protection_group.baas_protection_group_instance", "policy_id", policyIDUpdate),
					resource.TestCheckResourceAttr("ibm_baas_protection_group.baas_protection_group_instance", "environment", environmentUpdate),
				),
			},
		},
	})
}

func TestAccIbmBaasProtectionGroupAllArgs(t *testing.T) {
	var conf backuprecoveryv1.ProtectionGroupResponse
	xIbmTenantID := fmt.Sprintf("tf_x_ibm_tenant_id_%d", acctest.RandIntRange(10, 100))
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	policyID := fmt.Sprintf("tf_policy_id_%d", acctest.RandIntRange(10, 100))
	priority := "kLow"
	description := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	endTimeUsecs := fmt.Sprintf("%d", acctest.RandIntRange(10, 100))
	lastModifiedTimestampUsecs := fmt.Sprintf("%d", acctest.RandIntRange(10, 100))
	qosPolicy := "kBackupHDD"
	abortInBlackouts := "true"
	pauseInBlackouts := "true"
	isPaused := "false"
	environment := "kPhysical"
	xIbmTenantIDUpdate := fmt.Sprintf("tf_x_ibm_tenant_id_%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	policyIDUpdate := fmt.Sprintf("tf_policy_id_%d", acctest.RandIntRange(10, 100))
	priorityUpdate := "kHigh"
	descriptionUpdate := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	endTimeUsecsUpdate := fmt.Sprintf("%d", acctest.RandIntRange(10, 100))
	lastModifiedTimestampUsecsUpdate := fmt.Sprintf("%d", acctest.RandIntRange(10, 100))
	qosPolicyUpdate := "kBackupAll"
	abortInBlackoutsUpdate := "false"
	pauseInBlackoutsUpdate := "false"
	isPausedUpdate := "true"
	environmentUpdate := "kSQL"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmBaasProtectionGroupDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmBaasProtectionGroupConfig(xIbmTenantID, name, policyID, priority, description, endTimeUsecs, lastModifiedTimestampUsecs, qosPolicy, abortInBlackouts, pauseInBlackouts, isPaused, environment),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmBaasProtectionGroupExists("ibm_baas_protection_group.baas_protection_group_instance", conf),
					resource.TestCheckResourceAttr("ibm_baas_protection_group.baas_protection_group_instance", "x_ibm_tenant_id", xIbmTenantID),
					resource.TestCheckResourceAttr("ibm_baas_protection_group.baas_protection_group_instance", "name", name),
					resource.TestCheckResourceAttr("ibm_baas_protection_group.baas_protection_group_instance", "policy_id", policyID),
					resource.TestCheckResourceAttr("ibm_baas_protection_group.baas_protection_group_instance", "priority", priority),
					resource.TestCheckResourceAttr("ibm_baas_protection_group.baas_protection_group_instance", "description", description),
					resource.TestCheckResourceAttr("ibm_baas_protection_group.baas_protection_group_instance", "end_time_usecs", endTimeUsecs),
					resource.TestCheckResourceAttr("ibm_baas_protection_group.baas_protection_group_instance", "last_modified_timestamp_usecs", lastModifiedTimestampUsecs),
					resource.TestCheckResourceAttr("ibm_baas_protection_group.baas_protection_group_instance", "qos_policy", qosPolicy),
					resource.TestCheckResourceAttr("ibm_baas_protection_group.baas_protection_group_instance", "abort_in_blackouts", abortInBlackouts),
					resource.TestCheckResourceAttr("ibm_baas_protection_group.baas_protection_group_instance", "pause_in_blackouts", pauseInBlackouts),
					resource.TestCheckResourceAttr("ibm_baas_protection_group.baas_protection_group_instance", "is_paused", isPaused),
					resource.TestCheckResourceAttr("ibm_baas_protection_group.baas_protection_group_instance", "environment", environment),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmBaasProtectionGroupConfig(xIbmTenantIDUpdate, nameUpdate, policyIDUpdate, priorityUpdate, descriptionUpdate, endTimeUsecsUpdate, lastModifiedTimestampUsecsUpdate, qosPolicyUpdate, abortInBlackoutsUpdate, pauseInBlackoutsUpdate, isPausedUpdate, environmentUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_baas_protection_group.baas_protection_group_instance", "x_ibm_tenant_id", xIbmTenantIDUpdate),
					resource.TestCheckResourceAttr("ibm_baas_protection_group.baas_protection_group_instance", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_baas_protection_group.baas_protection_group_instance", "policy_id", policyIDUpdate),
					resource.TestCheckResourceAttr("ibm_baas_protection_group.baas_protection_group_instance", "priority", priorityUpdate),
					resource.TestCheckResourceAttr("ibm_baas_protection_group.baas_protection_group_instance", "description", descriptionUpdate),
					resource.TestCheckResourceAttr("ibm_baas_protection_group.baas_protection_group_instance", "end_time_usecs", endTimeUsecsUpdate),
					resource.TestCheckResourceAttr("ibm_baas_protection_group.baas_protection_group_instance", "last_modified_timestamp_usecs", lastModifiedTimestampUsecsUpdate),
					resource.TestCheckResourceAttr("ibm_baas_protection_group.baas_protection_group_instance", "qos_policy", qosPolicyUpdate),
					resource.TestCheckResourceAttr("ibm_baas_protection_group.baas_protection_group_instance", "abort_in_blackouts", abortInBlackoutsUpdate),
					resource.TestCheckResourceAttr("ibm_baas_protection_group.baas_protection_group_instance", "pause_in_blackouts", pauseInBlackoutsUpdate),
					resource.TestCheckResourceAttr("ibm_baas_protection_group.baas_protection_group_instance", "is_paused", isPausedUpdate),
					resource.TestCheckResourceAttr("ibm_baas_protection_group.baas_protection_group_instance", "environment", environmentUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_baas_protection_group.baas_protection_group",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmBaasProtectionGroupConfigBasic(xIbmTenantID string, name string, policyID string, environment string) string {
	return fmt.Sprintf(`
		resource "ibm_baas_protection_group" "baas_protection_group_instance" {
			x_ibm_tenant_id = "%s"
			name = "%s"
			policy_id = "%s"
			environment = "%s"
		}
	`, xIbmTenantID, name, policyID, environment)
}

func testAccCheckIbmBaasProtectionGroupConfig(xIbmTenantID string, name string, policyID string, priority string, description string, endTimeUsecs string, lastModifiedTimestampUsecs string, qosPolicy string, abortInBlackouts string, pauseInBlackouts string, isPaused string, environment string) string {
	return fmt.Sprintf(`

		resource "ibm_baas_protection_group" "baas_protection_group_instance" {
			x_ibm_tenant_id = "%s"
			name = "%s"
			policy_id = "%s"
			priority = "%s"
			description = "%s"
			start_time {
				hour = 0
				minute = 0
				time_zone = "time_zone"
			}
			end_time_usecs = %s
			last_modified_timestamp_usecs = %s
			alert_policy {
				backup_run_status = [ "kSuccess" ]
				alert_targets {
					email_address = "email_address"
					language = "en-us"
					recipient_type = "kTo"
				}
				raise_object_level_failure_alert = true
				raise_object_level_failure_alert_after_last_attempt = true
				raise_object_level_failure_alert_after_each_attempt = true
			}
			sla {
				backup_run_type = "kIncremental"
				sla_minutes = 1
			}
			qos_policy = "%s"
			abort_in_blackouts = %s
			pause_in_blackouts = %s
			is_paused = %s
			environment = "%s"
			advanced_configs {
				key = "key"
				value = "value"
			}
			physical_params {
				protection_type = "kFile"
				volume_protection_type_params {
					objects {
						id = 1
						name = "name"
						volume_guids = [ "volumeGuids" ]
						enable_system_backup = true
						excluded_vss_writers = [ "excludedVssWriters" ]
					}
					indexing_policy {
						enable_indexing = true
						include_paths = [ "includePaths" ]
						exclude_paths = [ "excludePaths" ]
					}
					perform_source_side_deduplication = true
					quiesce = true
					continue_on_quiesce_failure = true
					incremental_backup_after_restart = true
					pre_post_script {
						pre_script {
							path = "path"
							params = "params"
							timeout_secs = 1
							is_active = true
							continue_on_error = true
						}
						post_script {
							path = "path"
							params = "params"
							timeout_secs = 1
							is_active = true
						}
					}
					dedup_exclusion_source_ids = [ 1 ]
					excluded_vss_writers = [ "excludedVssWriters" ]
					cobmr_backup = true
				}
				file_protection_type_params {
					excluded_vss_writers = [ "excludedVssWriters" ]
					objects {
						excluded_vss_writers = [ "excludedVssWriters" ]
						id = 1
						name = "name"
						file_paths {
							included_path = "included_path"
							excluded_paths = [ "excludedPaths" ]
							skip_nested_volumes = true
						}
						uses_path_level_skip_nested_volume_setting = true
						nested_volume_types_to_skip = [ "nestedVolumeTypesToSkip" ]
						follow_nas_symlink_target = true
						metadata_file_path = "metadata_file_path"
					}
					indexing_policy {
						enable_indexing = true
						include_paths = [ "includePaths" ]
						exclude_paths = [ "excludePaths" ]
					}
					perform_source_side_deduplication = true
					perform_brick_based_deduplication = true
					task_timeouts {
						timeout_mins = 1
						backup_type = "kRegular"
					}
					quiesce = true
					continue_on_quiesce_failure = true
					cobmr_backup = true
					pre_post_script {
						pre_script {
							path = "path"
							params = "params"
							timeout_secs = 1
							is_active = true
							continue_on_error = true
						}
						post_script {
							path = "path"
							params = "params"
							timeout_secs = 1
							is_active = true
						}
					}
					dedup_exclusion_source_ids = [ 1 ]
					global_exclude_paths = [ "globalExcludePaths" ]
					global_exclude_fs = [ "globalExcludeFS" ]
					ignorable_errors = [ "kEOF" ]
					allow_parallel_runs = true
				}
			}
			mssql_params {
				file_protection_type_params {
					aag_backup_preference_type = "kPrimaryReplicaOnly"
					advanced_settings {
						cloned_db_backup_status = "kError"
						db_backup_if_not_online_status = "kError"
						missing_db_backup_status = "kError"
						offline_restoring_db_backup_status = "kError"
						read_only_db_backup_status = "kError"
						report_all_non_autoprotect_db_errors = "kError"
					}
					backup_system_dbs = true
					exclude_filters {
						filter_string = "filter_string"
						is_regular_expression = true
					}
					full_backups_copy_only = true
					log_backup_num_streams = 1
					log_backup_with_clause = "log_backup_with_clause"
					pre_post_script {
						pre_script {
							path = "path"
							params = "params"
							timeout_secs = 1
							is_active = true
							continue_on_error = true
						}
						post_script {
							path = "path"
							params = "params"
							timeout_secs = 1
							is_active = true
						}
					}
					use_aag_preferences_from_server = true
					user_db_backup_preference_type = "kBackupAllDatabases"
					additional_host_params {
						disable_source_side_deduplication = true
						host_id = 1
						host_name = "host_name"
					}
					objects {
						id = 1
						name = "name"
						source_type = "source_type"
					}
					perform_source_side_deduplication = true
				}
				native_protection_type_params {
					aag_backup_preference_type = "kPrimaryReplicaOnly"
					advanced_settings {
						cloned_db_backup_status = "kError"
						db_backup_if_not_online_status = "kError"
						missing_db_backup_status = "kError"
						offline_restoring_db_backup_status = "kError"
						read_only_db_backup_status = "kError"
						report_all_non_autoprotect_db_errors = "kError"
					}
					backup_system_dbs = true
					exclude_filters {
						filter_string = "filter_string"
						is_regular_expression = true
					}
					full_backups_copy_only = true
					log_backup_num_streams = 1
					log_backup_with_clause = "log_backup_with_clause"
					pre_post_script {
						pre_script {
							path = "path"
							params = "params"
							timeout_secs = 1
							is_active = true
							continue_on_error = true
						}
						post_script {
							path = "path"
							params = "params"
							timeout_secs = 1
							is_active = true
						}
					}
					use_aag_preferences_from_server = true
					user_db_backup_preference_type = "kBackupAllDatabases"
					num_streams = 1
					objects {
						id = 1
						name = "name"
						source_type = "source_type"
					}
					with_clause = "with_clause"
				}
				protection_type = "kFile"
				volume_protection_type_params {
					aag_backup_preference_type = "kPrimaryReplicaOnly"
					advanced_settings {
						cloned_db_backup_status = "kError"
						db_backup_if_not_online_status = "kError"
						missing_db_backup_status = "kError"
						offline_restoring_db_backup_status = "kError"
						read_only_db_backup_status = "kError"
						report_all_non_autoprotect_db_errors = "kError"
					}
					backup_system_dbs = true
					exclude_filters {
						filter_string = "filter_string"
						is_regular_expression = true
					}
					full_backups_copy_only = true
					log_backup_num_streams = 1
					log_backup_with_clause = "log_backup_with_clause"
					pre_post_script {
						pre_script {
							path = "path"
							params = "params"
							timeout_secs = 1
							is_active = true
							continue_on_error = true
						}
						post_script {
							path = "path"
							params = "params"
							timeout_secs = 1
							is_active = true
						}
					}
					use_aag_preferences_from_server = true
					user_db_backup_preference_type = "kBackupAllDatabases"
					additional_host_params {
						enable_system_backup = true
						host_id = 1
						host_name = "host_name"
						volume_guids = [ "volumeGuids" ]
					}
					backup_db_volumes_only = true
					incremental_backup_after_restart = true
					indexing_policy {
						enable_indexing = true
						include_paths = [ "includePaths" ]
						exclude_paths = [ "excludePaths" ]
					}
					objects {
						id = 1
						name = "name"
						source_type = "source_type"
					}
				}
			}
		}
	`, xIbmTenantID, name, policyID, priority, description, endTimeUsecs, lastModifiedTimestampUsecs, qosPolicy, abortInBlackouts, pauseInBlackouts, isPaused, environment)
}

func testAccCheckIbmBaasProtectionGroupExists(n string, obj backuprecoveryv1.ProtectionGroupResponse) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		backupRecoveryClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).BackupRecoveryV1()
		if err != nil {
			return err
		}

		getProtectionGroupByIdOptions := &backuprecoveryv1.GetProtectionGroupByIdOptions{}

		getProtectionGroupByIdOptions.SetID(rs.Primary.ID)

		protectionGroupResponse, _, err := backupRecoveryClient.GetProtectionGroupByID(getProtectionGroupByIdOptions)
		if err != nil {
			return err
		}

		obj = *protectionGroupResponse
		return nil
	}
}

func testAccCheckIbmBaasProtectionGroupDestroy(s *terraform.State) error {
	backupRecoveryClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).BackupRecoveryV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_baas_protection_group" {
			continue
		}

		getProtectionGroupByIdOptions := &backuprecoveryv1.GetProtectionGroupByIdOptions{}

		getProtectionGroupByIdOptions.SetID(rs.Primary.ID)

		// Try to find the key
		_, response, err := backupRecoveryClient.GetProtectionGroupByID(getProtectionGroupByIdOptions)

		if err == nil {
			return fmt.Errorf("baas_protection_group still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for baas_protection_group (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}

func TestResourceIbmBaasProtectionGroupTimeOfDayToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["hour"] = int(0)
		model["minute"] = int(0)
		model["time_zone"] = "America/Los_Angeles"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.TimeOfDay)
	model.Hour = core.Int64Ptr(int64(0))
	model.Minute = core.Int64Ptr(int64(0))
	model.TimeZone = core.StringPtr("America/Los_Angeles")

	result, err := backuprecovery.ResourceIbmBaasProtectionGroupTimeOfDayToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasProtectionGroupProtectionGroupAlertingPolicyToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		alertTargetModel := make(map[string]interface{})
		alertTargetModel["email_address"] = "testString"
		alertTargetModel["language"] = "en-us"
		alertTargetModel["recipient_type"] = "kTo"

		model := make(map[string]interface{})
		model["backup_run_status"] = []string{"kSuccess"}
		model["alert_targets"] = []map[string]interface{}{alertTargetModel}
		model["raise_object_level_failure_alert"] = true
		model["raise_object_level_failure_alert_after_last_attempt"] = true
		model["raise_object_level_failure_alert_after_each_attempt"] = true

		assert.Equal(t, result, model)
	}

	alertTargetModel := new(backuprecoveryv1.AlertTarget)
	alertTargetModel.EmailAddress = core.StringPtr("testString")
	alertTargetModel.Language = core.StringPtr("en-us")
	alertTargetModel.RecipientType = core.StringPtr("kTo")

	model := new(backuprecoveryv1.ProtectionGroupAlertingPolicy)
	model.BackupRunStatus = []string{"kSuccess"}
	model.AlertTargets = []backuprecoveryv1.AlertTarget{*alertTargetModel}
	model.RaiseObjectLevelFailureAlert = core.BoolPtr(true)
	model.RaiseObjectLevelFailureAlertAfterLastAttempt = core.BoolPtr(true)
	model.RaiseObjectLevelFailureAlertAfterEachAttempt = core.BoolPtr(true)

	result, err := backuprecovery.ResourceIbmBaasProtectionGroupProtectionGroupAlertingPolicyToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasProtectionGroupAlertTargetToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["email_address"] = "testString"
		model["language"] = "en-us"
		model["recipient_type"] = "kTo"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.AlertTarget)
	model.EmailAddress = core.StringPtr("testString")
	model.Language = core.StringPtr("en-us")
	model.RecipientType = core.StringPtr("kTo")

	result, err := backuprecovery.ResourceIbmBaasProtectionGroupAlertTargetToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasProtectionGroupSlaRuleToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["backup_run_type"] = "kIncremental"
		model["sla_minutes"] = int(1)

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.SlaRule)
	model.BackupRunType = core.StringPtr("kIncremental")
	model.SlaMinutes = core.Int64Ptr(int64(1))

	result, err := backuprecovery.ResourceIbmBaasProtectionGroupSlaRuleToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasProtectionGroupKeyValuePairToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["key"] = "testString"
		model["value"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.KeyValuePair)
	model.Key = core.StringPtr("testString")
	model.Value = core.StringPtr("testString")

	result, err := backuprecovery.ResourceIbmBaasProtectionGroupKeyValuePairToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasProtectionGroupPhysicalProtectionGroupParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		physicalVolumeProtectionGroupObjectParamsModel := make(map[string]interface{})
		physicalVolumeProtectionGroupObjectParamsModel["id"] = int(26)
		physicalVolumeProtectionGroupObjectParamsModel["name"] = "testString"
		physicalVolumeProtectionGroupObjectParamsModel["volume_guids"] = []string{"testString"}
		physicalVolumeProtectionGroupObjectParamsModel["enable_system_backup"] = true
		physicalVolumeProtectionGroupObjectParamsModel["excluded_vss_writers"] = []string{"testString"}

		indexingPolicyModel := make(map[string]interface{})
		indexingPolicyModel["enable_indexing"] = true
		indexingPolicyModel["include_paths"] = []string{"testString"}
		indexingPolicyModel["exclude_paths"] = []string{"testString"}

		commonPreBackupScriptParamsModel := make(map[string]interface{})
		commonPreBackupScriptParamsModel["path"] = "testString"
		commonPreBackupScriptParamsModel["params"] = "testString"
		commonPreBackupScriptParamsModel["timeout_secs"] = int(1)
		commonPreBackupScriptParamsModel["is_active"] = true
		commonPreBackupScriptParamsModel["continue_on_error"] = true

		commonPostBackupScriptParamsModel := make(map[string]interface{})
		commonPostBackupScriptParamsModel["path"] = "testString"
		commonPostBackupScriptParamsModel["params"] = "testString"
		commonPostBackupScriptParamsModel["timeout_secs"] = int(1)
		commonPostBackupScriptParamsModel["is_active"] = true

		prePostScriptParamsModel := make(map[string]interface{})
		prePostScriptParamsModel["pre_script"] = []map[string]interface{}{commonPreBackupScriptParamsModel}
		prePostScriptParamsModel["post_script"] = []map[string]interface{}{commonPostBackupScriptParamsModel}

		physicalVolumeProtectionGroupParamsModel := make(map[string]interface{})
		physicalVolumeProtectionGroupParamsModel["objects"] = []map[string]interface{}{physicalVolumeProtectionGroupObjectParamsModel}
		physicalVolumeProtectionGroupParamsModel["indexing_policy"] = []map[string]interface{}{indexingPolicyModel}
		physicalVolumeProtectionGroupParamsModel["perform_source_side_deduplication"] = true
		physicalVolumeProtectionGroupParamsModel["quiesce"] = true
		physicalVolumeProtectionGroupParamsModel["continue_on_quiesce_failure"] = true
		physicalVolumeProtectionGroupParamsModel["incremental_backup_after_restart"] = true
		physicalVolumeProtectionGroupParamsModel["pre_post_script"] = []map[string]interface{}{prePostScriptParamsModel}
		physicalVolumeProtectionGroupParamsModel["dedup_exclusion_source_ids"] = []int64{int64(26)}
		physicalVolumeProtectionGroupParamsModel["excluded_vss_writers"] = []string{"testString"}
		physicalVolumeProtectionGroupParamsModel["cobmr_backup"] = true

		physicalFileBackupPathParamsModel := make(map[string]interface{})
		physicalFileBackupPathParamsModel["included_path"] = "testString"
		physicalFileBackupPathParamsModel["excluded_paths"] = []string{"testString"}
		physicalFileBackupPathParamsModel["skip_nested_volumes"] = true

		physicalFileProtectionGroupObjectParamsModel := make(map[string]interface{})
		physicalFileProtectionGroupObjectParamsModel["excluded_vss_writers"] = []string{"testString"}
		physicalFileProtectionGroupObjectParamsModel["id"] = int(26)
		physicalFileProtectionGroupObjectParamsModel["file_paths"] = []map[string]interface{}{physicalFileBackupPathParamsModel}
		physicalFileProtectionGroupObjectParamsModel["uses_path_level_skip_nested_volume_setting"] = true
		physicalFileProtectionGroupObjectParamsModel["nested_volume_types_to_skip"] = []string{"testString"}
		physicalFileProtectionGroupObjectParamsModel["follow_nas_symlink_target"] = true
		physicalFileProtectionGroupObjectParamsModel["metadata_file_path"] = "testString"

		cancellationTimeoutParamsModel := make(map[string]interface{})
		cancellationTimeoutParamsModel["timeout_mins"] = int(26)
		cancellationTimeoutParamsModel["backup_type"] = "kRegular"

		physicalFileProtectionGroupParamsModel := make(map[string]interface{})
		physicalFileProtectionGroupParamsModel["excluded_vss_writers"] = []string{"testString"}
		physicalFileProtectionGroupParamsModel["objects"] = []map[string]interface{}{physicalFileProtectionGroupObjectParamsModel}
		physicalFileProtectionGroupParamsModel["indexing_policy"] = []map[string]interface{}{indexingPolicyModel}
		physicalFileProtectionGroupParamsModel["perform_source_side_deduplication"] = true
		physicalFileProtectionGroupParamsModel["perform_brick_based_deduplication"] = true
		physicalFileProtectionGroupParamsModel["task_timeouts"] = []map[string]interface{}{cancellationTimeoutParamsModel}
		physicalFileProtectionGroupParamsModel["quiesce"] = true
		physicalFileProtectionGroupParamsModel["continue_on_quiesce_failure"] = true
		physicalFileProtectionGroupParamsModel["cobmr_backup"] = true
		physicalFileProtectionGroupParamsModel["pre_post_script"] = []map[string]interface{}{prePostScriptParamsModel}
		physicalFileProtectionGroupParamsModel["dedup_exclusion_source_ids"] = []int64{int64(26)}
		physicalFileProtectionGroupParamsModel["global_exclude_paths"] = []string{"testString"}
		physicalFileProtectionGroupParamsModel["global_exclude_fs"] = []string{"testString"}
		physicalFileProtectionGroupParamsModel["ignorable_errors"] = []string{"kEOF"}
		physicalFileProtectionGroupParamsModel["allow_parallel_runs"] = true

		model := make(map[string]interface{})
		model["protection_type"] = "kFile"
		model["volume_protection_type_params"] = []map[string]interface{}{physicalVolumeProtectionGroupParamsModel}
		model["file_protection_type_params"] = []map[string]interface{}{physicalFileProtectionGroupParamsModel}

		assert.Equal(t, result, model)
	}

	physicalVolumeProtectionGroupObjectParamsModel := new(backuprecoveryv1.PhysicalVolumeProtectionGroupObjectParams)
	physicalVolumeProtectionGroupObjectParamsModel.ID = core.Int64Ptr(int64(26))
	physicalVolumeProtectionGroupObjectParamsModel.Name = core.StringPtr("testString")
	physicalVolumeProtectionGroupObjectParamsModel.VolumeGuids = []string{"testString"}
	physicalVolumeProtectionGroupObjectParamsModel.EnableSystemBackup = core.BoolPtr(true)
	physicalVolumeProtectionGroupObjectParamsModel.ExcludedVssWriters = []string{"testString"}

	indexingPolicyModel := new(backuprecoveryv1.IndexingPolicy)
	indexingPolicyModel.EnableIndexing = core.BoolPtr(true)
	indexingPolicyModel.IncludePaths = []string{"testString"}
	indexingPolicyModel.ExcludePaths = []string{"testString"}

	commonPreBackupScriptParamsModel := new(backuprecoveryv1.CommonPreBackupScriptParams)
	commonPreBackupScriptParamsModel.Path = core.StringPtr("testString")
	commonPreBackupScriptParamsModel.Params = core.StringPtr("testString")
	commonPreBackupScriptParamsModel.TimeoutSecs = core.Int64Ptr(int64(1))
	commonPreBackupScriptParamsModel.IsActive = core.BoolPtr(true)
	commonPreBackupScriptParamsModel.ContinueOnError = core.BoolPtr(true)

	commonPostBackupScriptParamsModel := new(backuprecoveryv1.CommonPostBackupScriptParams)
	commonPostBackupScriptParamsModel.Path = core.StringPtr("testString")
	commonPostBackupScriptParamsModel.Params = core.StringPtr("testString")
	commonPostBackupScriptParamsModel.TimeoutSecs = core.Int64Ptr(int64(1))
	commonPostBackupScriptParamsModel.IsActive = core.BoolPtr(true)

	prePostScriptParamsModel := new(backuprecoveryv1.PrePostScriptParams)
	prePostScriptParamsModel.PreScript = commonPreBackupScriptParamsModel
	prePostScriptParamsModel.PostScript = commonPostBackupScriptParamsModel

	physicalVolumeProtectionGroupParamsModel := new(backuprecoveryv1.PhysicalVolumeProtectionGroupParams)
	physicalVolumeProtectionGroupParamsModel.Objects = []backuprecoveryv1.PhysicalVolumeProtectionGroupObjectParams{*physicalVolumeProtectionGroupObjectParamsModel}
	physicalVolumeProtectionGroupParamsModel.IndexingPolicy = indexingPolicyModel
	physicalVolumeProtectionGroupParamsModel.PerformSourceSideDeduplication = core.BoolPtr(true)
	physicalVolumeProtectionGroupParamsModel.Quiesce = core.BoolPtr(true)
	physicalVolumeProtectionGroupParamsModel.ContinueOnQuiesceFailure = core.BoolPtr(true)
	physicalVolumeProtectionGroupParamsModel.IncrementalBackupAfterRestart = core.BoolPtr(true)
	physicalVolumeProtectionGroupParamsModel.PrePostScript = prePostScriptParamsModel
	physicalVolumeProtectionGroupParamsModel.DedupExclusionSourceIds = []int64{int64(26)}
	physicalVolumeProtectionGroupParamsModel.ExcludedVssWriters = []string{"testString"}
	physicalVolumeProtectionGroupParamsModel.CobmrBackup = core.BoolPtr(true)

	physicalFileBackupPathParamsModel := new(backuprecoveryv1.PhysicalFileBackupPathParams)
	physicalFileBackupPathParamsModel.IncludedPath = core.StringPtr("testString")
	physicalFileBackupPathParamsModel.ExcludedPaths = []string{"testString"}
	physicalFileBackupPathParamsModel.SkipNestedVolumes = core.BoolPtr(true)

	physicalFileProtectionGroupObjectParamsModel := new(backuprecoveryv1.PhysicalFileProtectionGroupObjectParams)
	physicalFileProtectionGroupObjectParamsModel.ExcludedVssWriters = []string{"testString"}
	physicalFileProtectionGroupObjectParamsModel.ID = core.Int64Ptr(int64(26))
	physicalFileProtectionGroupObjectParamsModel.FilePaths = []backuprecoveryv1.PhysicalFileBackupPathParams{*physicalFileBackupPathParamsModel}
	physicalFileProtectionGroupObjectParamsModel.UsesPathLevelSkipNestedVolumeSetting = core.BoolPtr(true)
	physicalFileProtectionGroupObjectParamsModel.NestedVolumeTypesToSkip = []string{"testString"}
	physicalFileProtectionGroupObjectParamsModel.FollowNasSymlinkTarget = core.BoolPtr(true)
	physicalFileProtectionGroupObjectParamsModel.MetadataFilePath = core.StringPtr("testString")

	cancellationTimeoutParamsModel := new(backuprecoveryv1.CancellationTimeoutParams)
	cancellationTimeoutParamsModel.TimeoutMins = core.Int64Ptr(int64(26))
	cancellationTimeoutParamsModel.BackupType = core.StringPtr("kRegular")

	physicalFileProtectionGroupParamsModel := new(backuprecoveryv1.PhysicalFileProtectionGroupParams)
	physicalFileProtectionGroupParamsModel.ExcludedVssWriters = []string{"testString"}
	physicalFileProtectionGroupParamsModel.Objects = []backuprecoveryv1.PhysicalFileProtectionGroupObjectParams{*physicalFileProtectionGroupObjectParamsModel}
	physicalFileProtectionGroupParamsModel.IndexingPolicy = indexingPolicyModel
	physicalFileProtectionGroupParamsModel.PerformSourceSideDeduplication = core.BoolPtr(true)
	physicalFileProtectionGroupParamsModel.PerformBrickBasedDeduplication = core.BoolPtr(true)
	physicalFileProtectionGroupParamsModel.TaskTimeouts = []backuprecoveryv1.CancellationTimeoutParams{*cancellationTimeoutParamsModel}
	physicalFileProtectionGroupParamsModel.Quiesce = core.BoolPtr(true)
	physicalFileProtectionGroupParamsModel.ContinueOnQuiesceFailure = core.BoolPtr(true)
	physicalFileProtectionGroupParamsModel.CobmrBackup = core.BoolPtr(true)
	physicalFileProtectionGroupParamsModel.PrePostScript = prePostScriptParamsModel
	physicalFileProtectionGroupParamsModel.DedupExclusionSourceIds = []int64{int64(26)}
	physicalFileProtectionGroupParamsModel.GlobalExcludePaths = []string{"testString"}
	physicalFileProtectionGroupParamsModel.GlobalExcludeFS = []string{"testString"}
	physicalFileProtectionGroupParamsModel.IgnorableErrors = []string{"kEOF"}
	physicalFileProtectionGroupParamsModel.AllowParallelRuns = core.BoolPtr(true)

	model := new(backuprecoveryv1.PhysicalProtectionGroupParams)
	model.ProtectionType = core.StringPtr("kFile")
	model.VolumeProtectionTypeParams = physicalVolumeProtectionGroupParamsModel
	model.FileProtectionTypeParams = physicalFileProtectionGroupParamsModel

	result, err := backuprecovery.ResourceIbmBaasProtectionGroupPhysicalProtectionGroupParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasProtectionGroupPhysicalVolumeProtectionGroupParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		physicalVolumeProtectionGroupObjectParamsModel := make(map[string]interface{})
		physicalVolumeProtectionGroupObjectParamsModel["id"] = int(26)
		physicalVolumeProtectionGroupObjectParamsModel["name"] = "testString"
		physicalVolumeProtectionGroupObjectParamsModel["volume_guids"] = []string{"testString"}
		physicalVolumeProtectionGroupObjectParamsModel["enable_system_backup"] = true
		physicalVolumeProtectionGroupObjectParamsModel["excluded_vss_writers"] = []string{"testString"}

		indexingPolicyModel := make(map[string]interface{})
		indexingPolicyModel["enable_indexing"] = true
		indexingPolicyModel["include_paths"] = []string{"testString"}
		indexingPolicyModel["exclude_paths"] = []string{"testString"}

		commonPreBackupScriptParamsModel := make(map[string]interface{})
		commonPreBackupScriptParamsModel["path"] = "testString"
		commonPreBackupScriptParamsModel["params"] = "testString"
		commonPreBackupScriptParamsModel["timeout_secs"] = int(1)
		commonPreBackupScriptParamsModel["is_active"] = true
		commonPreBackupScriptParamsModel["continue_on_error"] = true

		commonPostBackupScriptParamsModel := make(map[string]interface{})
		commonPostBackupScriptParamsModel["path"] = "testString"
		commonPostBackupScriptParamsModel["params"] = "testString"
		commonPostBackupScriptParamsModel["timeout_secs"] = int(1)
		commonPostBackupScriptParamsModel["is_active"] = true

		prePostScriptParamsModel := make(map[string]interface{})
		prePostScriptParamsModel["pre_script"] = []map[string]interface{}{commonPreBackupScriptParamsModel}
		prePostScriptParamsModel["post_script"] = []map[string]interface{}{commonPostBackupScriptParamsModel}

		model := make(map[string]interface{})
		model["objects"] = []map[string]interface{}{physicalVolumeProtectionGroupObjectParamsModel}
		model["indexing_policy"] = []map[string]interface{}{indexingPolicyModel}
		model["perform_source_side_deduplication"] = true
		model["quiesce"] = true
		model["continue_on_quiesce_failure"] = true
		model["incremental_backup_after_restart"] = true
		model["pre_post_script"] = []map[string]interface{}{prePostScriptParamsModel}
		model["dedup_exclusion_source_ids"] = []int64{int64(26)}
		model["excluded_vss_writers"] = []string{"testString"}
		model["cobmr_backup"] = true

		assert.Equal(t, result, model)
	}

	physicalVolumeProtectionGroupObjectParamsModel := new(backuprecoveryv1.PhysicalVolumeProtectionGroupObjectParams)
	physicalVolumeProtectionGroupObjectParamsModel.ID = core.Int64Ptr(int64(26))
	physicalVolumeProtectionGroupObjectParamsModel.Name = core.StringPtr("testString")
	physicalVolumeProtectionGroupObjectParamsModel.VolumeGuids = []string{"testString"}
	physicalVolumeProtectionGroupObjectParamsModel.EnableSystemBackup = core.BoolPtr(true)
	physicalVolumeProtectionGroupObjectParamsModel.ExcludedVssWriters = []string{"testString"}

	indexingPolicyModel := new(backuprecoveryv1.IndexingPolicy)
	indexingPolicyModel.EnableIndexing = core.BoolPtr(true)
	indexingPolicyModel.IncludePaths = []string{"testString"}
	indexingPolicyModel.ExcludePaths = []string{"testString"}

	commonPreBackupScriptParamsModel := new(backuprecoveryv1.CommonPreBackupScriptParams)
	commonPreBackupScriptParamsModel.Path = core.StringPtr("testString")
	commonPreBackupScriptParamsModel.Params = core.StringPtr("testString")
	commonPreBackupScriptParamsModel.TimeoutSecs = core.Int64Ptr(int64(1))
	commonPreBackupScriptParamsModel.IsActive = core.BoolPtr(true)
	commonPreBackupScriptParamsModel.ContinueOnError = core.BoolPtr(true)

	commonPostBackupScriptParamsModel := new(backuprecoveryv1.CommonPostBackupScriptParams)
	commonPostBackupScriptParamsModel.Path = core.StringPtr("testString")
	commonPostBackupScriptParamsModel.Params = core.StringPtr("testString")
	commonPostBackupScriptParamsModel.TimeoutSecs = core.Int64Ptr(int64(1))
	commonPostBackupScriptParamsModel.IsActive = core.BoolPtr(true)

	prePostScriptParamsModel := new(backuprecoveryv1.PrePostScriptParams)
	prePostScriptParamsModel.PreScript = commonPreBackupScriptParamsModel
	prePostScriptParamsModel.PostScript = commonPostBackupScriptParamsModel

	model := new(backuprecoveryv1.PhysicalVolumeProtectionGroupParams)
	model.Objects = []backuprecoveryv1.PhysicalVolumeProtectionGroupObjectParams{*physicalVolumeProtectionGroupObjectParamsModel}
	model.IndexingPolicy = indexingPolicyModel
	model.PerformSourceSideDeduplication = core.BoolPtr(true)
	model.Quiesce = core.BoolPtr(true)
	model.ContinueOnQuiesceFailure = core.BoolPtr(true)
	model.IncrementalBackupAfterRestart = core.BoolPtr(true)
	model.PrePostScript = prePostScriptParamsModel
	model.DedupExclusionSourceIds = []int64{int64(26)}
	model.ExcludedVssWriters = []string{"testString"}
	model.CobmrBackup = core.BoolPtr(true)

	result, err := backuprecovery.ResourceIbmBaasProtectionGroupPhysicalVolumeProtectionGroupParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasProtectionGroupPhysicalVolumeProtectionGroupObjectParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["id"] = int(26)
		model["name"] = "testString"
		model["volume_guids"] = []string{"testString"}
		model["enable_system_backup"] = true
		model["excluded_vss_writers"] = []string{"testString"}

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.PhysicalVolumeProtectionGroupObjectParams)
	model.ID = core.Int64Ptr(int64(26))
	model.Name = core.StringPtr("testString")
	model.VolumeGuids = []string{"testString"}
	model.EnableSystemBackup = core.BoolPtr(true)
	model.ExcludedVssWriters = []string{"testString"}

	result, err := backuprecovery.ResourceIbmBaasProtectionGroupPhysicalVolumeProtectionGroupObjectParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasProtectionGroupIndexingPolicyToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["enable_indexing"] = true
		model["include_paths"] = []string{"testString"}
		model["exclude_paths"] = []string{"testString"}

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.IndexingPolicy)
	model.EnableIndexing = core.BoolPtr(true)
	model.IncludePaths = []string{"testString"}
	model.ExcludePaths = []string{"testString"}

	result, err := backuprecovery.ResourceIbmBaasProtectionGroupIndexingPolicyToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasProtectionGroupPrePostScriptParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		commonPreBackupScriptParamsModel := make(map[string]interface{})
		commonPreBackupScriptParamsModel["path"] = "testString"
		commonPreBackupScriptParamsModel["params"] = "testString"
		commonPreBackupScriptParamsModel["timeout_secs"] = int(1)
		commonPreBackupScriptParamsModel["is_active"] = true
		commonPreBackupScriptParamsModel["continue_on_error"] = true

		commonPostBackupScriptParamsModel := make(map[string]interface{})
		commonPostBackupScriptParamsModel["path"] = "testString"
		commonPostBackupScriptParamsModel["params"] = "testString"
		commonPostBackupScriptParamsModel["timeout_secs"] = int(1)
		commonPostBackupScriptParamsModel["is_active"] = true

		model := make(map[string]interface{})
		model["pre_script"] = []map[string]interface{}{commonPreBackupScriptParamsModel}
		model["post_script"] = []map[string]interface{}{commonPostBackupScriptParamsModel}

		assert.Equal(t, result, model)
	}

	commonPreBackupScriptParamsModel := new(backuprecoveryv1.CommonPreBackupScriptParams)
	commonPreBackupScriptParamsModel.Path = core.StringPtr("testString")
	commonPreBackupScriptParamsModel.Params = core.StringPtr("testString")
	commonPreBackupScriptParamsModel.TimeoutSecs = core.Int64Ptr(int64(1))
	commonPreBackupScriptParamsModel.IsActive = core.BoolPtr(true)
	commonPreBackupScriptParamsModel.ContinueOnError = core.BoolPtr(true)

	commonPostBackupScriptParamsModel := new(backuprecoveryv1.CommonPostBackupScriptParams)
	commonPostBackupScriptParamsModel.Path = core.StringPtr("testString")
	commonPostBackupScriptParamsModel.Params = core.StringPtr("testString")
	commonPostBackupScriptParamsModel.TimeoutSecs = core.Int64Ptr(int64(1))
	commonPostBackupScriptParamsModel.IsActive = core.BoolPtr(true)

	model := new(backuprecoveryv1.PrePostScriptParams)
	model.PreScript = commonPreBackupScriptParamsModel
	model.PostScript = commonPostBackupScriptParamsModel

	result, err := backuprecovery.ResourceIbmBaasProtectionGroupPrePostScriptParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasProtectionGroupCommonPreBackupScriptParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["path"] = "testString"
		model["params"] = "testString"
		model["timeout_secs"] = int(1)
		model["is_active"] = true
		model["continue_on_error"] = true

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.CommonPreBackupScriptParams)
	model.Path = core.StringPtr("testString")
	model.Params = core.StringPtr("testString")
	model.TimeoutSecs = core.Int64Ptr(int64(1))
	model.IsActive = core.BoolPtr(true)
	model.ContinueOnError = core.BoolPtr(true)

	result, err := backuprecovery.ResourceIbmBaasProtectionGroupCommonPreBackupScriptParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasProtectionGroupCommonPostBackupScriptParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["path"] = "testString"
		model["params"] = "testString"
		model["timeout_secs"] = int(1)
		model["is_active"] = true

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.CommonPostBackupScriptParams)
	model.Path = core.StringPtr("testString")
	model.Params = core.StringPtr("testString")
	model.TimeoutSecs = core.Int64Ptr(int64(1))
	model.IsActive = core.BoolPtr(true)

	result, err := backuprecovery.ResourceIbmBaasProtectionGroupCommonPostBackupScriptParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasProtectionGroupPhysicalFileProtectionGroupParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		physicalFileBackupPathParamsModel := make(map[string]interface{})
		physicalFileBackupPathParamsModel["included_path"] = "testString"
		physicalFileBackupPathParamsModel["excluded_paths"] = []string{"testString"}
		physicalFileBackupPathParamsModel["skip_nested_volumes"] = true

		physicalFileProtectionGroupObjectParamsModel := make(map[string]interface{})
		physicalFileProtectionGroupObjectParamsModel["excluded_vss_writers"] = []string{"testString"}
		physicalFileProtectionGroupObjectParamsModel["id"] = int(26)
		physicalFileProtectionGroupObjectParamsModel["file_paths"] = []map[string]interface{}{physicalFileBackupPathParamsModel}
		physicalFileProtectionGroupObjectParamsModel["uses_path_level_skip_nested_volume_setting"] = true
		physicalFileProtectionGroupObjectParamsModel["nested_volume_types_to_skip"] = []string{"testString"}
		physicalFileProtectionGroupObjectParamsModel["follow_nas_symlink_target"] = true
		physicalFileProtectionGroupObjectParamsModel["metadata_file_path"] = "testString"

		indexingPolicyModel := make(map[string]interface{})
		indexingPolicyModel["enable_indexing"] = true
		indexingPolicyModel["include_paths"] = []string{"testString"}
		indexingPolicyModel["exclude_paths"] = []string{"testString"}

		cancellationTimeoutParamsModel := make(map[string]interface{})
		cancellationTimeoutParamsModel["timeout_mins"] = int(26)
		cancellationTimeoutParamsModel["backup_type"] = "kRegular"

		commonPreBackupScriptParamsModel := make(map[string]interface{})
		commonPreBackupScriptParamsModel["path"] = "testString"
		commonPreBackupScriptParamsModel["params"] = "testString"
		commonPreBackupScriptParamsModel["timeout_secs"] = int(1)
		commonPreBackupScriptParamsModel["is_active"] = true
		commonPreBackupScriptParamsModel["continue_on_error"] = true

		commonPostBackupScriptParamsModel := make(map[string]interface{})
		commonPostBackupScriptParamsModel["path"] = "testString"
		commonPostBackupScriptParamsModel["params"] = "testString"
		commonPostBackupScriptParamsModel["timeout_secs"] = int(1)
		commonPostBackupScriptParamsModel["is_active"] = true

		prePostScriptParamsModel := make(map[string]interface{})
		prePostScriptParamsModel["pre_script"] = []map[string]interface{}{commonPreBackupScriptParamsModel}
		prePostScriptParamsModel["post_script"] = []map[string]interface{}{commonPostBackupScriptParamsModel}

		model := make(map[string]interface{})
		model["excluded_vss_writers"] = []string{"testString"}
		model["objects"] = []map[string]interface{}{physicalFileProtectionGroupObjectParamsModel}
		model["indexing_policy"] = []map[string]interface{}{indexingPolicyModel}
		model["perform_source_side_deduplication"] = true
		model["perform_brick_based_deduplication"] = true
		model["task_timeouts"] = []map[string]interface{}{cancellationTimeoutParamsModel}
		model["quiesce"] = true
		model["continue_on_quiesce_failure"] = true
		model["cobmr_backup"] = true
		model["pre_post_script"] = []map[string]interface{}{prePostScriptParamsModel}
		model["dedup_exclusion_source_ids"] = []int64{int64(26)}
		model["global_exclude_paths"] = []string{"testString"}
		model["global_exclude_fs"] = []string{"testString"}
		model["ignorable_errors"] = []string{"kEOF"}
		model["allow_parallel_runs"] = true

		assert.Equal(t, result, model)
	}

	physicalFileBackupPathParamsModel := new(backuprecoveryv1.PhysicalFileBackupPathParams)
	physicalFileBackupPathParamsModel.IncludedPath = core.StringPtr("testString")
	physicalFileBackupPathParamsModel.ExcludedPaths = []string{"testString"}
	physicalFileBackupPathParamsModel.SkipNestedVolumes = core.BoolPtr(true)

	physicalFileProtectionGroupObjectParamsModel := new(backuprecoveryv1.PhysicalFileProtectionGroupObjectParams)
	physicalFileProtectionGroupObjectParamsModel.ExcludedVssWriters = []string{"testString"}
	physicalFileProtectionGroupObjectParamsModel.ID = core.Int64Ptr(int64(26))
	physicalFileProtectionGroupObjectParamsModel.FilePaths = []backuprecoveryv1.PhysicalFileBackupPathParams{*physicalFileBackupPathParamsModel}
	physicalFileProtectionGroupObjectParamsModel.UsesPathLevelSkipNestedVolumeSetting = core.BoolPtr(true)
	physicalFileProtectionGroupObjectParamsModel.NestedVolumeTypesToSkip = []string{"testString"}
	physicalFileProtectionGroupObjectParamsModel.FollowNasSymlinkTarget = core.BoolPtr(true)
	physicalFileProtectionGroupObjectParamsModel.MetadataFilePath = core.StringPtr("testString")

	indexingPolicyModel := new(backuprecoveryv1.IndexingPolicy)
	indexingPolicyModel.EnableIndexing = core.BoolPtr(true)
	indexingPolicyModel.IncludePaths = []string{"testString"}
	indexingPolicyModel.ExcludePaths = []string{"testString"}

	cancellationTimeoutParamsModel := new(backuprecoveryv1.CancellationTimeoutParams)
	cancellationTimeoutParamsModel.TimeoutMins = core.Int64Ptr(int64(26))
	cancellationTimeoutParamsModel.BackupType = core.StringPtr("kRegular")

	commonPreBackupScriptParamsModel := new(backuprecoveryv1.CommonPreBackupScriptParams)
	commonPreBackupScriptParamsModel.Path = core.StringPtr("testString")
	commonPreBackupScriptParamsModel.Params = core.StringPtr("testString")
	commonPreBackupScriptParamsModel.TimeoutSecs = core.Int64Ptr(int64(1))
	commonPreBackupScriptParamsModel.IsActive = core.BoolPtr(true)
	commonPreBackupScriptParamsModel.ContinueOnError = core.BoolPtr(true)

	commonPostBackupScriptParamsModel := new(backuprecoveryv1.CommonPostBackupScriptParams)
	commonPostBackupScriptParamsModel.Path = core.StringPtr("testString")
	commonPostBackupScriptParamsModel.Params = core.StringPtr("testString")
	commonPostBackupScriptParamsModel.TimeoutSecs = core.Int64Ptr(int64(1))
	commonPostBackupScriptParamsModel.IsActive = core.BoolPtr(true)

	prePostScriptParamsModel := new(backuprecoveryv1.PrePostScriptParams)
	prePostScriptParamsModel.PreScript = commonPreBackupScriptParamsModel
	prePostScriptParamsModel.PostScript = commonPostBackupScriptParamsModel

	model := new(backuprecoveryv1.PhysicalFileProtectionGroupParams)
	model.ExcludedVssWriters = []string{"testString"}
	model.Objects = []backuprecoveryv1.PhysicalFileProtectionGroupObjectParams{*physicalFileProtectionGroupObjectParamsModel}
	model.IndexingPolicy = indexingPolicyModel
	model.PerformSourceSideDeduplication = core.BoolPtr(true)
	model.PerformBrickBasedDeduplication = core.BoolPtr(true)
	model.TaskTimeouts = []backuprecoveryv1.CancellationTimeoutParams{*cancellationTimeoutParamsModel}
	model.Quiesce = core.BoolPtr(true)
	model.ContinueOnQuiesceFailure = core.BoolPtr(true)
	model.CobmrBackup = core.BoolPtr(true)
	model.PrePostScript = prePostScriptParamsModel
	model.DedupExclusionSourceIds = []int64{int64(26)}
	model.GlobalExcludePaths = []string{"testString"}
	model.GlobalExcludeFS = []string{"testString"}
	model.IgnorableErrors = []string{"kEOF"}
	model.AllowParallelRuns = core.BoolPtr(true)

	result, err := backuprecovery.ResourceIbmBaasProtectionGroupPhysicalFileProtectionGroupParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasProtectionGroupPhysicalFileProtectionGroupObjectParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		physicalFileBackupPathParamsModel := make(map[string]interface{})
		physicalFileBackupPathParamsModel["included_path"] = "testString"
		physicalFileBackupPathParamsModel["excluded_paths"] = []string{"testString"}
		physicalFileBackupPathParamsModel["skip_nested_volumes"] = true

		model := make(map[string]interface{})
		model["excluded_vss_writers"] = []string{"testString"}
		model["id"] = int(26)
		model["name"] = "testString"
		model["file_paths"] = []map[string]interface{}{physicalFileBackupPathParamsModel}
		model["uses_path_level_skip_nested_volume_setting"] = true
		model["nested_volume_types_to_skip"] = []string{"testString"}
		model["follow_nas_symlink_target"] = true
		model["metadata_file_path"] = "testString"

		assert.Equal(t, result, model)
	}

	physicalFileBackupPathParamsModel := new(backuprecoveryv1.PhysicalFileBackupPathParams)
	physicalFileBackupPathParamsModel.IncludedPath = core.StringPtr("testString")
	physicalFileBackupPathParamsModel.ExcludedPaths = []string{"testString"}
	physicalFileBackupPathParamsModel.SkipNestedVolumes = core.BoolPtr(true)

	model := new(backuprecoveryv1.PhysicalFileProtectionGroupObjectParams)
	model.ExcludedVssWriters = []string{"testString"}
	model.ID = core.Int64Ptr(int64(26))
	model.Name = core.StringPtr("testString")
	model.FilePaths = []backuprecoveryv1.PhysicalFileBackupPathParams{*physicalFileBackupPathParamsModel}
	model.UsesPathLevelSkipNestedVolumeSetting = core.BoolPtr(true)
	model.NestedVolumeTypesToSkip = []string{"testString"}
	model.FollowNasSymlinkTarget = core.BoolPtr(true)
	model.MetadataFilePath = core.StringPtr("testString")

	result, err := backuprecovery.ResourceIbmBaasProtectionGroupPhysicalFileProtectionGroupObjectParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasProtectionGroupPhysicalFileBackupPathParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["included_path"] = "testString"
		model["excluded_paths"] = []string{"testString"}
		model["skip_nested_volumes"] = true

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.PhysicalFileBackupPathParams)
	model.IncludedPath = core.StringPtr("testString")
	model.ExcludedPaths = []string{"testString"}
	model.SkipNestedVolumes = core.BoolPtr(true)

	result, err := backuprecovery.ResourceIbmBaasProtectionGroupPhysicalFileBackupPathParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasProtectionGroupCancellationTimeoutParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["timeout_mins"] = int(26)
		model["backup_type"] = "kRegular"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.CancellationTimeoutParams)
	model.TimeoutMins = core.Int64Ptr(int64(26))
	model.BackupType = core.StringPtr("kRegular")

	result, err := backuprecovery.ResourceIbmBaasProtectionGroupCancellationTimeoutParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasProtectionGroupMSSQLProtectionGroupParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		advancedSettingsModel := make(map[string]interface{})
		advancedSettingsModel["cloned_db_backup_status"] = "kError"
		advancedSettingsModel["db_backup_if_not_online_status"] = "kError"
		advancedSettingsModel["missing_db_backup_status"] = "kError"
		advancedSettingsModel["offline_restoring_db_backup_status"] = "kError"
		advancedSettingsModel["read_only_db_backup_status"] = "kError"
		advancedSettingsModel["report_all_non_autoprotect_db_errors"] = "kError"

		filterModel := make(map[string]interface{})
		filterModel["filter_string"] = "testString"
		filterModel["is_regular_expression"] = false

		commonPreBackupScriptParamsModel := make(map[string]interface{})
		commonPreBackupScriptParamsModel["path"] = "testString"
		commonPreBackupScriptParamsModel["params"] = "testString"
		commonPreBackupScriptParamsModel["timeout_secs"] = int(1)
		commonPreBackupScriptParamsModel["is_active"] = true
		commonPreBackupScriptParamsModel["continue_on_error"] = true

		commonPostBackupScriptParamsModel := make(map[string]interface{})
		commonPostBackupScriptParamsModel["path"] = "testString"
		commonPostBackupScriptParamsModel["params"] = "testString"
		commonPostBackupScriptParamsModel["timeout_secs"] = int(1)
		commonPostBackupScriptParamsModel["is_active"] = true

		prePostScriptParamsModel := make(map[string]interface{})
		prePostScriptParamsModel["pre_script"] = []map[string]interface{}{commonPreBackupScriptParamsModel}
		prePostScriptParamsModel["post_script"] = []map[string]interface{}{commonPostBackupScriptParamsModel}

		mssqlFileProtectionGroupHostParamsModel := make(map[string]interface{})
		mssqlFileProtectionGroupHostParamsModel["disable_source_side_deduplication"] = true
		mssqlFileProtectionGroupHostParamsModel["host_id"] = int(26)

		mssqlFileProtectionGroupObjectParamsModel := make(map[string]interface{})
		mssqlFileProtectionGroupObjectParamsModel["id"] = int(26)

		mssqlFileProtectionGroupParamsModel := make(map[string]interface{})
		mssqlFileProtectionGroupParamsModel["aag_backup_preference_type"] = "kPrimaryReplicaOnly"
		mssqlFileProtectionGroupParamsModel["advanced_settings"] = []map[string]interface{}{advancedSettingsModel}
		mssqlFileProtectionGroupParamsModel["backup_system_dbs"] = true
		mssqlFileProtectionGroupParamsModel["exclude_filters"] = []map[string]interface{}{filterModel}
		mssqlFileProtectionGroupParamsModel["full_backups_copy_only"] = true
		mssqlFileProtectionGroupParamsModel["log_backup_num_streams"] = int(38)
		mssqlFileProtectionGroupParamsModel["log_backup_with_clause"] = "testString"
		mssqlFileProtectionGroupParamsModel["pre_post_script"] = []map[string]interface{}{prePostScriptParamsModel}
		mssqlFileProtectionGroupParamsModel["use_aag_preferences_from_server"] = true
		mssqlFileProtectionGroupParamsModel["user_db_backup_preference_type"] = "kBackupAllDatabases"
		mssqlFileProtectionGroupParamsModel["additional_host_params"] = []map[string]interface{}{mssqlFileProtectionGroupHostParamsModel}
		mssqlFileProtectionGroupParamsModel["objects"] = []map[string]interface{}{mssqlFileProtectionGroupObjectParamsModel}
		mssqlFileProtectionGroupParamsModel["perform_source_side_deduplication"] = true

		mssqlNativeProtectionGroupObjectParamsModel := make(map[string]interface{})
		mssqlNativeProtectionGroupObjectParamsModel["id"] = int(26)

		mssqlNativeProtectionGroupParamsModel := make(map[string]interface{})
		mssqlNativeProtectionGroupParamsModel["aag_backup_preference_type"] = "kPrimaryReplicaOnly"
		mssqlNativeProtectionGroupParamsModel["advanced_settings"] = []map[string]interface{}{advancedSettingsModel}
		mssqlNativeProtectionGroupParamsModel["backup_system_dbs"] = true
		mssqlNativeProtectionGroupParamsModel["exclude_filters"] = []map[string]interface{}{filterModel}
		mssqlNativeProtectionGroupParamsModel["full_backups_copy_only"] = true
		mssqlNativeProtectionGroupParamsModel["log_backup_num_streams"] = int(38)
		mssqlNativeProtectionGroupParamsModel["log_backup_with_clause"] = "testString"
		mssqlNativeProtectionGroupParamsModel["pre_post_script"] = []map[string]interface{}{prePostScriptParamsModel}
		mssqlNativeProtectionGroupParamsModel["use_aag_preferences_from_server"] = true
		mssqlNativeProtectionGroupParamsModel["user_db_backup_preference_type"] = "kBackupAllDatabases"
		mssqlNativeProtectionGroupParamsModel["num_streams"] = int(38)
		mssqlNativeProtectionGroupParamsModel["objects"] = []map[string]interface{}{mssqlNativeProtectionGroupObjectParamsModel}
		mssqlNativeProtectionGroupParamsModel["with_clause"] = "testString"

		mssqlVolumeProtectionGroupHostParamsModel := make(map[string]interface{})
		mssqlVolumeProtectionGroupHostParamsModel["enable_system_backup"] = true
		mssqlVolumeProtectionGroupHostParamsModel["host_id"] = int(26)
		mssqlVolumeProtectionGroupHostParamsModel["volume_guids"] = []string{"testString"}

		indexingPolicyModel := make(map[string]interface{})
		indexingPolicyModel["enable_indexing"] = true
		indexingPolicyModel["include_paths"] = []string{"testString"}
		indexingPolicyModel["exclude_paths"] = []string{"testString"}

		mssqlVolumeProtectionGroupObjectParamsModel := make(map[string]interface{})
		mssqlVolumeProtectionGroupObjectParamsModel["id"] = int(26)

		mssqlVolumeProtectionGroupParamsModel := make(map[string]interface{})
		mssqlVolumeProtectionGroupParamsModel["aag_backup_preference_type"] = "kPrimaryReplicaOnly"
		mssqlVolumeProtectionGroupParamsModel["advanced_settings"] = []map[string]interface{}{advancedSettingsModel}
		mssqlVolumeProtectionGroupParamsModel["backup_system_dbs"] = true
		mssqlVolumeProtectionGroupParamsModel["exclude_filters"] = []map[string]interface{}{filterModel}
		mssqlVolumeProtectionGroupParamsModel["full_backups_copy_only"] = true
		mssqlVolumeProtectionGroupParamsModel["log_backup_num_streams"] = int(38)
		mssqlVolumeProtectionGroupParamsModel["log_backup_with_clause"] = "testString"
		mssqlVolumeProtectionGroupParamsModel["pre_post_script"] = []map[string]interface{}{prePostScriptParamsModel}
		mssqlVolumeProtectionGroupParamsModel["use_aag_preferences_from_server"] = true
		mssqlVolumeProtectionGroupParamsModel["user_db_backup_preference_type"] = "kBackupAllDatabases"
		mssqlVolumeProtectionGroupParamsModel["additional_host_params"] = []map[string]interface{}{mssqlVolumeProtectionGroupHostParamsModel}
		mssqlVolumeProtectionGroupParamsModel["backup_db_volumes_only"] = true
		mssqlVolumeProtectionGroupParamsModel["incremental_backup_after_restart"] = true
		mssqlVolumeProtectionGroupParamsModel["indexing_policy"] = []map[string]interface{}{indexingPolicyModel}
		mssqlVolumeProtectionGroupParamsModel["objects"] = []map[string]interface{}{mssqlVolumeProtectionGroupObjectParamsModel}

		model := make(map[string]interface{})
		model["file_protection_type_params"] = []map[string]interface{}{mssqlFileProtectionGroupParamsModel}
		model["native_protection_type_params"] = []map[string]interface{}{mssqlNativeProtectionGroupParamsModel}
		model["protection_type"] = "kFile"
		model["volume_protection_type_params"] = []map[string]interface{}{mssqlVolumeProtectionGroupParamsModel}

		assert.Equal(t, result, model)
	}

	advancedSettingsModel := new(backuprecoveryv1.AdvancedSettings)
	advancedSettingsModel.ClonedDbBackupStatus = core.StringPtr("kError")
	advancedSettingsModel.DbBackupIfNotOnlineStatus = core.StringPtr("kError")
	advancedSettingsModel.MissingDbBackupStatus = core.StringPtr("kError")
	advancedSettingsModel.OfflineRestoringDbBackupStatus = core.StringPtr("kError")
	advancedSettingsModel.ReadOnlyDbBackupStatus = core.StringPtr("kError")
	advancedSettingsModel.ReportAllNonAutoprotectDbErrors = core.StringPtr("kError")

	filterModel := new(backuprecoveryv1.Filter)
	filterModel.FilterString = core.StringPtr("testString")
	filterModel.IsRegularExpression = core.BoolPtr(false)

	commonPreBackupScriptParamsModel := new(backuprecoveryv1.CommonPreBackupScriptParams)
	commonPreBackupScriptParamsModel.Path = core.StringPtr("testString")
	commonPreBackupScriptParamsModel.Params = core.StringPtr("testString")
	commonPreBackupScriptParamsModel.TimeoutSecs = core.Int64Ptr(int64(1))
	commonPreBackupScriptParamsModel.IsActive = core.BoolPtr(true)
	commonPreBackupScriptParamsModel.ContinueOnError = core.BoolPtr(true)

	commonPostBackupScriptParamsModel := new(backuprecoveryv1.CommonPostBackupScriptParams)
	commonPostBackupScriptParamsModel.Path = core.StringPtr("testString")
	commonPostBackupScriptParamsModel.Params = core.StringPtr("testString")
	commonPostBackupScriptParamsModel.TimeoutSecs = core.Int64Ptr(int64(1))
	commonPostBackupScriptParamsModel.IsActive = core.BoolPtr(true)

	prePostScriptParamsModel := new(backuprecoveryv1.PrePostScriptParams)
	prePostScriptParamsModel.PreScript = commonPreBackupScriptParamsModel
	prePostScriptParamsModel.PostScript = commonPostBackupScriptParamsModel

	mssqlFileProtectionGroupHostParamsModel := new(backuprecoveryv1.MSSQLFileProtectionGroupHostParams)
	mssqlFileProtectionGroupHostParamsModel.DisableSourceSideDeduplication = core.BoolPtr(true)
	mssqlFileProtectionGroupHostParamsModel.HostID = core.Int64Ptr(int64(26))

	mssqlFileProtectionGroupObjectParamsModel := new(backuprecoveryv1.MSSQLFileProtectionGroupObjectParams)
	mssqlFileProtectionGroupObjectParamsModel.ID = core.Int64Ptr(int64(26))

	mssqlFileProtectionGroupParamsModel := new(backuprecoveryv1.MSSQLFileProtectionGroupParams)
	mssqlFileProtectionGroupParamsModel.AagBackupPreferenceType = core.StringPtr("kPrimaryReplicaOnly")
	mssqlFileProtectionGroupParamsModel.AdvancedSettings = advancedSettingsModel
	mssqlFileProtectionGroupParamsModel.BackupSystemDbs = core.BoolPtr(true)
	mssqlFileProtectionGroupParamsModel.ExcludeFilters = []backuprecoveryv1.Filter{*filterModel}
	mssqlFileProtectionGroupParamsModel.FullBackupsCopyOnly = core.BoolPtr(true)
	mssqlFileProtectionGroupParamsModel.LogBackupNumStreams = core.Int64Ptr(int64(38))
	mssqlFileProtectionGroupParamsModel.LogBackupWithClause = core.StringPtr("testString")
	mssqlFileProtectionGroupParamsModel.PrePostScript = prePostScriptParamsModel
	mssqlFileProtectionGroupParamsModel.UseAagPreferencesFromServer = core.BoolPtr(true)
	mssqlFileProtectionGroupParamsModel.UserDbBackupPreferenceType = core.StringPtr("kBackupAllDatabases")
	mssqlFileProtectionGroupParamsModel.AdditionalHostParams = []backuprecoveryv1.MSSQLFileProtectionGroupHostParams{*mssqlFileProtectionGroupHostParamsModel}
	mssqlFileProtectionGroupParamsModel.Objects = []backuprecoveryv1.MSSQLFileProtectionGroupObjectParams{*mssqlFileProtectionGroupObjectParamsModel}
	mssqlFileProtectionGroupParamsModel.PerformSourceSideDeduplication = core.BoolPtr(true)

	mssqlNativeProtectionGroupObjectParamsModel := new(backuprecoveryv1.MSSQLNativeProtectionGroupObjectParams)
	mssqlNativeProtectionGroupObjectParamsModel.ID = core.Int64Ptr(int64(26))

	mssqlNativeProtectionGroupParamsModel := new(backuprecoveryv1.MSSQLNativeProtectionGroupParams)
	mssqlNativeProtectionGroupParamsModel.AagBackupPreferenceType = core.StringPtr("kPrimaryReplicaOnly")
	mssqlNativeProtectionGroupParamsModel.AdvancedSettings = advancedSettingsModel
	mssqlNativeProtectionGroupParamsModel.BackupSystemDbs = core.BoolPtr(true)
	mssqlNativeProtectionGroupParamsModel.ExcludeFilters = []backuprecoveryv1.Filter{*filterModel}
	mssqlNativeProtectionGroupParamsModel.FullBackupsCopyOnly = core.BoolPtr(true)
	mssqlNativeProtectionGroupParamsModel.LogBackupNumStreams = core.Int64Ptr(int64(38))
	mssqlNativeProtectionGroupParamsModel.LogBackupWithClause = core.StringPtr("testString")
	mssqlNativeProtectionGroupParamsModel.PrePostScript = prePostScriptParamsModel
	mssqlNativeProtectionGroupParamsModel.UseAagPreferencesFromServer = core.BoolPtr(true)
	mssqlNativeProtectionGroupParamsModel.UserDbBackupPreferenceType = core.StringPtr("kBackupAllDatabases")
	mssqlNativeProtectionGroupParamsModel.NumStreams = core.Int64Ptr(int64(38))
	mssqlNativeProtectionGroupParamsModel.Objects = []backuprecoveryv1.MSSQLNativeProtectionGroupObjectParams{*mssqlNativeProtectionGroupObjectParamsModel}
	mssqlNativeProtectionGroupParamsModel.WithClause = core.StringPtr("testString")

	mssqlVolumeProtectionGroupHostParamsModel := new(backuprecoveryv1.MSSQLVolumeProtectionGroupHostParams)
	mssqlVolumeProtectionGroupHostParamsModel.EnableSystemBackup = core.BoolPtr(true)
	mssqlVolumeProtectionGroupHostParamsModel.HostID = core.Int64Ptr(int64(26))
	mssqlVolumeProtectionGroupHostParamsModel.VolumeGuids = []string{"testString"}

	indexingPolicyModel := new(backuprecoveryv1.IndexingPolicy)
	indexingPolicyModel.EnableIndexing = core.BoolPtr(true)
	indexingPolicyModel.IncludePaths = []string{"testString"}
	indexingPolicyModel.ExcludePaths = []string{"testString"}

	mssqlVolumeProtectionGroupObjectParamsModel := new(backuprecoveryv1.MSSQLVolumeProtectionGroupObjectParams)
	mssqlVolumeProtectionGroupObjectParamsModel.ID = core.Int64Ptr(int64(26))

	mssqlVolumeProtectionGroupParamsModel := new(backuprecoveryv1.MSSQLVolumeProtectionGroupParams)
	mssqlVolumeProtectionGroupParamsModel.AagBackupPreferenceType = core.StringPtr("kPrimaryReplicaOnly")
	mssqlVolumeProtectionGroupParamsModel.AdvancedSettings = advancedSettingsModel
	mssqlVolumeProtectionGroupParamsModel.BackupSystemDbs = core.BoolPtr(true)
	mssqlVolumeProtectionGroupParamsModel.ExcludeFilters = []backuprecoveryv1.Filter{*filterModel}
	mssqlVolumeProtectionGroupParamsModel.FullBackupsCopyOnly = core.BoolPtr(true)
	mssqlVolumeProtectionGroupParamsModel.LogBackupNumStreams = core.Int64Ptr(int64(38))
	mssqlVolumeProtectionGroupParamsModel.LogBackupWithClause = core.StringPtr("testString")
	mssqlVolumeProtectionGroupParamsModel.PrePostScript = prePostScriptParamsModel
	mssqlVolumeProtectionGroupParamsModel.UseAagPreferencesFromServer = core.BoolPtr(true)
	mssqlVolumeProtectionGroupParamsModel.UserDbBackupPreferenceType = core.StringPtr("kBackupAllDatabases")
	mssqlVolumeProtectionGroupParamsModel.AdditionalHostParams = []backuprecoveryv1.MSSQLVolumeProtectionGroupHostParams{*mssqlVolumeProtectionGroupHostParamsModel}
	mssqlVolumeProtectionGroupParamsModel.BackupDbVolumesOnly = core.BoolPtr(true)
	mssqlVolumeProtectionGroupParamsModel.IncrementalBackupAfterRestart = core.BoolPtr(true)
	mssqlVolumeProtectionGroupParamsModel.IndexingPolicy = indexingPolicyModel
	mssqlVolumeProtectionGroupParamsModel.Objects = []backuprecoveryv1.MSSQLVolumeProtectionGroupObjectParams{*mssqlVolumeProtectionGroupObjectParamsModel}

	model := new(backuprecoveryv1.MSSQLProtectionGroupParams)
	model.FileProtectionTypeParams = mssqlFileProtectionGroupParamsModel
	model.NativeProtectionTypeParams = mssqlNativeProtectionGroupParamsModel
	model.ProtectionType = core.StringPtr("kFile")
	model.VolumeProtectionTypeParams = mssqlVolumeProtectionGroupParamsModel

	result, err := backuprecovery.ResourceIbmBaasProtectionGroupMSSQLProtectionGroupParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasProtectionGroupMSSQLFileProtectionGroupParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		advancedSettingsModel := make(map[string]interface{})
		advancedSettingsModel["cloned_db_backup_status"] = "kError"
		advancedSettingsModel["db_backup_if_not_online_status"] = "kError"
		advancedSettingsModel["missing_db_backup_status"] = "kError"
		advancedSettingsModel["offline_restoring_db_backup_status"] = "kError"
		advancedSettingsModel["read_only_db_backup_status"] = "kError"
		advancedSettingsModel["report_all_non_autoprotect_db_errors"] = "kError"

		filterModel := make(map[string]interface{})
		filterModel["filter_string"] = "testString"
		filterModel["is_regular_expression"] = false

		commonPreBackupScriptParamsModel := make(map[string]interface{})
		commonPreBackupScriptParamsModel["path"] = "testString"
		commonPreBackupScriptParamsModel["params"] = "testString"
		commonPreBackupScriptParamsModel["timeout_secs"] = int(1)
		commonPreBackupScriptParamsModel["is_active"] = true
		commonPreBackupScriptParamsModel["continue_on_error"] = true

		commonPostBackupScriptParamsModel := make(map[string]interface{})
		commonPostBackupScriptParamsModel["path"] = "testString"
		commonPostBackupScriptParamsModel["params"] = "testString"
		commonPostBackupScriptParamsModel["timeout_secs"] = int(1)
		commonPostBackupScriptParamsModel["is_active"] = true

		prePostScriptParamsModel := make(map[string]interface{})
		prePostScriptParamsModel["pre_script"] = []map[string]interface{}{commonPreBackupScriptParamsModel}
		prePostScriptParamsModel["post_script"] = []map[string]interface{}{commonPostBackupScriptParamsModel}

		mssqlFileProtectionGroupHostParamsModel := make(map[string]interface{})
		mssqlFileProtectionGroupHostParamsModel["disable_source_side_deduplication"] = true
		mssqlFileProtectionGroupHostParamsModel["host_id"] = int(26)

		mssqlFileProtectionGroupObjectParamsModel := make(map[string]interface{})
		mssqlFileProtectionGroupObjectParamsModel["id"] = int(26)

		model := make(map[string]interface{})
		model["aag_backup_preference_type"] = "kPrimaryReplicaOnly"
		model["advanced_settings"] = []map[string]interface{}{advancedSettingsModel}
		model["backup_system_dbs"] = true
		model["exclude_filters"] = []map[string]interface{}{filterModel}
		model["full_backups_copy_only"] = true
		model["log_backup_num_streams"] = int(38)
		model["log_backup_with_clause"] = "testString"
		model["pre_post_script"] = []map[string]interface{}{prePostScriptParamsModel}
		model["use_aag_preferences_from_server"] = true
		model["user_db_backup_preference_type"] = "kBackupAllDatabases"
		model["additional_host_params"] = []map[string]interface{}{mssqlFileProtectionGroupHostParamsModel}
		model["objects"] = []map[string]interface{}{mssqlFileProtectionGroupObjectParamsModel}
		model["perform_source_side_deduplication"] = true

		assert.Equal(t, result, model)
	}

	advancedSettingsModel := new(backuprecoveryv1.AdvancedSettings)
	advancedSettingsModel.ClonedDbBackupStatus = core.StringPtr("kError")
	advancedSettingsModel.DbBackupIfNotOnlineStatus = core.StringPtr("kError")
	advancedSettingsModel.MissingDbBackupStatus = core.StringPtr("kError")
	advancedSettingsModel.OfflineRestoringDbBackupStatus = core.StringPtr("kError")
	advancedSettingsModel.ReadOnlyDbBackupStatus = core.StringPtr("kError")
	advancedSettingsModel.ReportAllNonAutoprotectDbErrors = core.StringPtr("kError")

	filterModel := new(backuprecoveryv1.Filter)
	filterModel.FilterString = core.StringPtr("testString")
	filterModel.IsRegularExpression = core.BoolPtr(false)

	commonPreBackupScriptParamsModel := new(backuprecoveryv1.CommonPreBackupScriptParams)
	commonPreBackupScriptParamsModel.Path = core.StringPtr("testString")
	commonPreBackupScriptParamsModel.Params = core.StringPtr("testString")
	commonPreBackupScriptParamsModel.TimeoutSecs = core.Int64Ptr(int64(1))
	commonPreBackupScriptParamsModel.IsActive = core.BoolPtr(true)
	commonPreBackupScriptParamsModel.ContinueOnError = core.BoolPtr(true)

	commonPostBackupScriptParamsModel := new(backuprecoveryv1.CommonPostBackupScriptParams)
	commonPostBackupScriptParamsModel.Path = core.StringPtr("testString")
	commonPostBackupScriptParamsModel.Params = core.StringPtr("testString")
	commonPostBackupScriptParamsModel.TimeoutSecs = core.Int64Ptr(int64(1))
	commonPostBackupScriptParamsModel.IsActive = core.BoolPtr(true)

	prePostScriptParamsModel := new(backuprecoveryv1.PrePostScriptParams)
	prePostScriptParamsModel.PreScript = commonPreBackupScriptParamsModel
	prePostScriptParamsModel.PostScript = commonPostBackupScriptParamsModel

	mssqlFileProtectionGroupHostParamsModel := new(backuprecoveryv1.MSSQLFileProtectionGroupHostParams)
	mssqlFileProtectionGroupHostParamsModel.DisableSourceSideDeduplication = core.BoolPtr(true)
	mssqlFileProtectionGroupHostParamsModel.HostID = core.Int64Ptr(int64(26))

	mssqlFileProtectionGroupObjectParamsModel := new(backuprecoveryv1.MSSQLFileProtectionGroupObjectParams)
	mssqlFileProtectionGroupObjectParamsModel.ID = core.Int64Ptr(int64(26))

	model := new(backuprecoveryv1.MSSQLFileProtectionGroupParams)
	model.AagBackupPreferenceType = core.StringPtr("kPrimaryReplicaOnly")
	model.AdvancedSettings = advancedSettingsModel
	model.BackupSystemDbs = core.BoolPtr(true)
	model.ExcludeFilters = []backuprecoveryv1.Filter{*filterModel}
	model.FullBackupsCopyOnly = core.BoolPtr(true)
	model.LogBackupNumStreams = core.Int64Ptr(int64(38))
	model.LogBackupWithClause = core.StringPtr("testString")
	model.PrePostScript = prePostScriptParamsModel
	model.UseAagPreferencesFromServer = core.BoolPtr(true)
	model.UserDbBackupPreferenceType = core.StringPtr("kBackupAllDatabases")
	model.AdditionalHostParams = []backuprecoveryv1.MSSQLFileProtectionGroupHostParams{*mssqlFileProtectionGroupHostParamsModel}
	model.Objects = []backuprecoveryv1.MSSQLFileProtectionGroupObjectParams{*mssqlFileProtectionGroupObjectParamsModel}
	model.PerformSourceSideDeduplication = core.BoolPtr(true)

	result, err := backuprecovery.ResourceIbmBaasProtectionGroupMSSQLFileProtectionGroupParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasProtectionGroupAdvancedSettingsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["cloned_db_backup_status"] = "kError"
		model["db_backup_if_not_online_status"] = "kError"
		model["missing_db_backup_status"] = "kError"
		model["offline_restoring_db_backup_status"] = "kError"
		model["read_only_db_backup_status"] = "kError"
		model["report_all_non_autoprotect_db_errors"] = "kError"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.AdvancedSettings)
	model.ClonedDbBackupStatus = core.StringPtr("kError")
	model.DbBackupIfNotOnlineStatus = core.StringPtr("kError")
	model.MissingDbBackupStatus = core.StringPtr("kError")
	model.OfflineRestoringDbBackupStatus = core.StringPtr("kError")
	model.ReadOnlyDbBackupStatus = core.StringPtr("kError")
	model.ReportAllNonAutoprotectDbErrors = core.StringPtr("kError")

	result, err := backuprecovery.ResourceIbmBaasProtectionGroupAdvancedSettingsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasProtectionGroupFilterToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["filter_string"] = "testString"
		model["is_regular_expression"] = false

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.Filter)
	model.FilterString = core.StringPtr("testString")
	model.IsRegularExpression = core.BoolPtr(false)

	result, err := backuprecovery.ResourceIbmBaasProtectionGroupFilterToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasProtectionGroupMSSQLFileProtectionGroupHostParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["disable_source_side_deduplication"] = true
		model["host_id"] = int(26)
		model["host_name"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.MSSQLFileProtectionGroupHostParams)
	model.DisableSourceSideDeduplication = core.BoolPtr(true)
	model.HostID = core.Int64Ptr(int64(26))
	model.HostName = core.StringPtr("testString")

	result, err := backuprecovery.ResourceIbmBaasProtectionGroupMSSQLFileProtectionGroupHostParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasProtectionGroupMSSQLFileProtectionGroupObjectParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["id"] = int(26)
		model["name"] = "testString"
		model["source_type"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.MSSQLFileProtectionGroupObjectParams)
	model.ID = core.Int64Ptr(int64(26))
	model.Name = core.StringPtr("testString")
	model.SourceType = core.StringPtr("testString")

	result, err := backuprecovery.ResourceIbmBaasProtectionGroupMSSQLFileProtectionGroupObjectParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasProtectionGroupMSSQLNativeProtectionGroupParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		advancedSettingsModel := make(map[string]interface{})
		advancedSettingsModel["cloned_db_backup_status"] = "kError"
		advancedSettingsModel["db_backup_if_not_online_status"] = "kError"
		advancedSettingsModel["missing_db_backup_status"] = "kError"
		advancedSettingsModel["offline_restoring_db_backup_status"] = "kError"
		advancedSettingsModel["read_only_db_backup_status"] = "kError"
		advancedSettingsModel["report_all_non_autoprotect_db_errors"] = "kError"

		filterModel := make(map[string]interface{})
		filterModel["filter_string"] = "testString"
		filterModel["is_regular_expression"] = false

		commonPreBackupScriptParamsModel := make(map[string]interface{})
		commonPreBackupScriptParamsModel["path"] = "testString"
		commonPreBackupScriptParamsModel["params"] = "testString"
		commonPreBackupScriptParamsModel["timeout_secs"] = int(1)
		commonPreBackupScriptParamsModel["is_active"] = true
		commonPreBackupScriptParamsModel["continue_on_error"] = true

		commonPostBackupScriptParamsModel := make(map[string]interface{})
		commonPostBackupScriptParamsModel["path"] = "testString"
		commonPostBackupScriptParamsModel["params"] = "testString"
		commonPostBackupScriptParamsModel["timeout_secs"] = int(1)
		commonPostBackupScriptParamsModel["is_active"] = true

		prePostScriptParamsModel := make(map[string]interface{})
		prePostScriptParamsModel["pre_script"] = []map[string]interface{}{commonPreBackupScriptParamsModel}
		prePostScriptParamsModel["post_script"] = []map[string]interface{}{commonPostBackupScriptParamsModel}

		mssqlNativeProtectionGroupObjectParamsModel := make(map[string]interface{})
		mssqlNativeProtectionGroupObjectParamsModel["id"] = int(26)

		model := make(map[string]interface{})
		model["aag_backup_preference_type"] = "kPrimaryReplicaOnly"
		model["advanced_settings"] = []map[string]interface{}{advancedSettingsModel}
		model["backup_system_dbs"] = true
		model["exclude_filters"] = []map[string]interface{}{filterModel}
		model["full_backups_copy_only"] = true
		model["log_backup_num_streams"] = int(38)
		model["log_backup_with_clause"] = "testString"
		model["pre_post_script"] = []map[string]interface{}{prePostScriptParamsModel}
		model["use_aag_preferences_from_server"] = true
		model["user_db_backup_preference_type"] = "kBackupAllDatabases"
		model["num_streams"] = int(38)
		model["objects"] = []map[string]interface{}{mssqlNativeProtectionGroupObjectParamsModel}
		model["with_clause"] = "testString"

		assert.Equal(t, result, model)
	}

	advancedSettingsModel := new(backuprecoveryv1.AdvancedSettings)
	advancedSettingsModel.ClonedDbBackupStatus = core.StringPtr("kError")
	advancedSettingsModel.DbBackupIfNotOnlineStatus = core.StringPtr("kError")
	advancedSettingsModel.MissingDbBackupStatus = core.StringPtr("kError")
	advancedSettingsModel.OfflineRestoringDbBackupStatus = core.StringPtr("kError")
	advancedSettingsModel.ReadOnlyDbBackupStatus = core.StringPtr("kError")
	advancedSettingsModel.ReportAllNonAutoprotectDbErrors = core.StringPtr("kError")

	filterModel := new(backuprecoveryv1.Filter)
	filterModel.FilterString = core.StringPtr("testString")
	filterModel.IsRegularExpression = core.BoolPtr(false)

	commonPreBackupScriptParamsModel := new(backuprecoveryv1.CommonPreBackupScriptParams)
	commonPreBackupScriptParamsModel.Path = core.StringPtr("testString")
	commonPreBackupScriptParamsModel.Params = core.StringPtr("testString")
	commonPreBackupScriptParamsModel.TimeoutSecs = core.Int64Ptr(int64(1))
	commonPreBackupScriptParamsModel.IsActive = core.BoolPtr(true)
	commonPreBackupScriptParamsModel.ContinueOnError = core.BoolPtr(true)

	commonPostBackupScriptParamsModel := new(backuprecoveryv1.CommonPostBackupScriptParams)
	commonPostBackupScriptParamsModel.Path = core.StringPtr("testString")
	commonPostBackupScriptParamsModel.Params = core.StringPtr("testString")
	commonPostBackupScriptParamsModel.TimeoutSecs = core.Int64Ptr(int64(1))
	commonPostBackupScriptParamsModel.IsActive = core.BoolPtr(true)

	prePostScriptParamsModel := new(backuprecoveryv1.PrePostScriptParams)
	prePostScriptParamsModel.PreScript = commonPreBackupScriptParamsModel
	prePostScriptParamsModel.PostScript = commonPostBackupScriptParamsModel

	mssqlNativeProtectionGroupObjectParamsModel := new(backuprecoveryv1.MSSQLNativeProtectionGroupObjectParams)
	mssqlNativeProtectionGroupObjectParamsModel.ID = core.Int64Ptr(int64(26))

	model := new(backuprecoveryv1.MSSQLNativeProtectionGroupParams)
	model.AagBackupPreferenceType = core.StringPtr("kPrimaryReplicaOnly")
	model.AdvancedSettings = advancedSettingsModel
	model.BackupSystemDbs = core.BoolPtr(true)
	model.ExcludeFilters = []backuprecoveryv1.Filter{*filterModel}
	model.FullBackupsCopyOnly = core.BoolPtr(true)
	model.LogBackupNumStreams = core.Int64Ptr(int64(38))
	model.LogBackupWithClause = core.StringPtr("testString")
	model.PrePostScript = prePostScriptParamsModel
	model.UseAagPreferencesFromServer = core.BoolPtr(true)
	model.UserDbBackupPreferenceType = core.StringPtr("kBackupAllDatabases")
	model.NumStreams = core.Int64Ptr(int64(38))
	model.Objects = []backuprecoveryv1.MSSQLNativeProtectionGroupObjectParams{*mssqlNativeProtectionGroupObjectParamsModel}
	model.WithClause = core.StringPtr("testString")

	result, err := backuprecovery.ResourceIbmBaasProtectionGroupMSSQLNativeProtectionGroupParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasProtectionGroupMSSQLNativeProtectionGroupObjectParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["id"] = int(26)
		model["name"] = "testString"
		model["source_type"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.MSSQLNativeProtectionGroupObjectParams)
	model.ID = core.Int64Ptr(int64(26))
	model.Name = core.StringPtr("testString")
	model.SourceType = core.StringPtr("testString")

	result, err := backuprecovery.ResourceIbmBaasProtectionGroupMSSQLNativeProtectionGroupObjectParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasProtectionGroupMSSQLVolumeProtectionGroupParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		advancedSettingsModel := make(map[string]interface{})
		advancedSettingsModel["cloned_db_backup_status"] = "kError"
		advancedSettingsModel["db_backup_if_not_online_status"] = "kError"
		advancedSettingsModel["missing_db_backup_status"] = "kError"
		advancedSettingsModel["offline_restoring_db_backup_status"] = "kError"
		advancedSettingsModel["read_only_db_backup_status"] = "kError"
		advancedSettingsModel["report_all_non_autoprotect_db_errors"] = "kError"

		filterModel := make(map[string]interface{})
		filterModel["filter_string"] = "testString"
		filterModel["is_regular_expression"] = false

		commonPreBackupScriptParamsModel := make(map[string]interface{})
		commonPreBackupScriptParamsModel["path"] = "testString"
		commonPreBackupScriptParamsModel["params"] = "testString"
		commonPreBackupScriptParamsModel["timeout_secs"] = int(1)
		commonPreBackupScriptParamsModel["is_active"] = true
		commonPreBackupScriptParamsModel["continue_on_error"] = true

		commonPostBackupScriptParamsModel := make(map[string]interface{})
		commonPostBackupScriptParamsModel["path"] = "testString"
		commonPostBackupScriptParamsModel["params"] = "testString"
		commonPostBackupScriptParamsModel["timeout_secs"] = int(1)
		commonPostBackupScriptParamsModel["is_active"] = true

		prePostScriptParamsModel := make(map[string]interface{})
		prePostScriptParamsModel["pre_script"] = []map[string]interface{}{commonPreBackupScriptParamsModel}
		prePostScriptParamsModel["post_script"] = []map[string]interface{}{commonPostBackupScriptParamsModel}

		mssqlVolumeProtectionGroupHostParamsModel := make(map[string]interface{})
		mssqlVolumeProtectionGroupHostParamsModel["enable_system_backup"] = true
		mssqlVolumeProtectionGroupHostParamsModel["host_id"] = int(26)
		mssqlVolumeProtectionGroupHostParamsModel["volume_guids"] = []string{"testString"}

		indexingPolicyModel := make(map[string]interface{})
		indexingPolicyModel["enable_indexing"] = true
		indexingPolicyModel["include_paths"] = []string{"testString"}
		indexingPolicyModel["exclude_paths"] = []string{"testString"}

		mssqlVolumeProtectionGroupObjectParamsModel := make(map[string]interface{})
		mssqlVolumeProtectionGroupObjectParamsModel["id"] = int(26)

		model := make(map[string]interface{})
		model["aag_backup_preference_type"] = "kPrimaryReplicaOnly"
		model["advanced_settings"] = []map[string]interface{}{advancedSettingsModel}
		model["backup_system_dbs"] = true
		model["exclude_filters"] = []map[string]interface{}{filterModel}
		model["full_backups_copy_only"] = true
		model["log_backup_num_streams"] = int(38)
		model["log_backup_with_clause"] = "testString"
		model["pre_post_script"] = []map[string]interface{}{prePostScriptParamsModel}
		model["use_aag_preferences_from_server"] = true
		model["user_db_backup_preference_type"] = "kBackupAllDatabases"
		model["additional_host_params"] = []map[string]interface{}{mssqlVolumeProtectionGroupHostParamsModel}
		model["backup_db_volumes_only"] = true
		model["incremental_backup_after_restart"] = true
		model["indexing_policy"] = []map[string]interface{}{indexingPolicyModel}
		model["objects"] = []map[string]interface{}{mssqlVolumeProtectionGroupObjectParamsModel}

		assert.Equal(t, result, model)
	}

	advancedSettingsModel := new(backuprecoveryv1.AdvancedSettings)
	advancedSettingsModel.ClonedDbBackupStatus = core.StringPtr("kError")
	advancedSettingsModel.DbBackupIfNotOnlineStatus = core.StringPtr("kError")
	advancedSettingsModel.MissingDbBackupStatus = core.StringPtr("kError")
	advancedSettingsModel.OfflineRestoringDbBackupStatus = core.StringPtr("kError")
	advancedSettingsModel.ReadOnlyDbBackupStatus = core.StringPtr("kError")
	advancedSettingsModel.ReportAllNonAutoprotectDbErrors = core.StringPtr("kError")

	filterModel := new(backuprecoveryv1.Filter)
	filterModel.FilterString = core.StringPtr("testString")
	filterModel.IsRegularExpression = core.BoolPtr(false)

	commonPreBackupScriptParamsModel := new(backuprecoveryv1.CommonPreBackupScriptParams)
	commonPreBackupScriptParamsModel.Path = core.StringPtr("testString")
	commonPreBackupScriptParamsModel.Params = core.StringPtr("testString")
	commonPreBackupScriptParamsModel.TimeoutSecs = core.Int64Ptr(int64(1))
	commonPreBackupScriptParamsModel.IsActive = core.BoolPtr(true)
	commonPreBackupScriptParamsModel.ContinueOnError = core.BoolPtr(true)

	commonPostBackupScriptParamsModel := new(backuprecoveryv1.CommonPostBackupScriptParams)
	commonPostBackupScriptParamsModel.Path = core.StringPtr("testString")
	commonPostBackupScriptParamsModel.Params = core.StringPtr("testString")
	commonPostBackupScriptParamsModel.TimeoutSecs = core.Int64Ptr(int64(1))
	commonPostBackupScriptParamsModel.IsActive = core.BoolPtr(true)

	prePostScriptParamsModel := new(backuprecoveryv1.PrePostScriptParams)
	prePostScriptParamsModel.PreScript = commonPreBackupScriptParamsModel
	prePostScriptParamsModel.PostScript = commonPostBackupScriptParamsModel

	mssqlVolumeProtectionGroupHostParamsModel := new(backuprecoveryv1.MSSQLVolumeProtectionGroupHostParams)
	mssqlVolumeProtectionGroupHostParamsModel.EnableSystemBackup = core.BoolPtr(true)
	mssqlVolumeProtectionGroupHostParamsModel.HostID = core.Int64Ptr(int64(26))
	mssqlVolumeProtectionGroupHostParamsModel.VolumeGuids = []string{"testString"}

	indexingPolicyModel := new(backuprecoveryv1.IndexingPolicy)
	indexingPolicyModel.EnableIndexing = core.BoolPtr(true)
	indexingPolicyModel.IncludePaths = []string{"testString"}
	indexingPolicyModel.ExcludePaths = []string{"testString"}

	mssqlVolumeProtectionGroupObjectParamsModel := new(backuprecoveryv1.MSSQLVolumeProtectionGroupObjectParams)
	mssqlVolumeProtectionGroupObjectParamsModel.ID = core.Int64Ptr(int64(26))

	model := new(backuprecoveryv1.MSSQLVolumeProtectionGroupParams)
	model.AagBackupPreferenceType = core.StringPtr("kPrimaryReplicaOnly")
	model.AdvancedSettings = advancedSettingsModel
	model.BackupSystemDbs = core.BoolPtr(true)
	model.ExcludeFilters = []backuprecoveryv1.Filter{*filterModel}
	model.FullBackupsCopyOnly = core.BoolPtr(true)
	model.LogBackupNumStreams = core.Int64Ptr(int64(38))
	model.LogBackupWithClause = core.StringPtr("testString")
	model.PrePostScript = prePostScriptParamsModel
	model.UseAagPreferencesFromServer = core.BoolPtr(true)
	model.UserDbBackupPreferenceType = core.StringPtr("kBackupAllDatabases")
	model.AdditionalHostParams = []backuprecoveryv1.MSSQLVolumeProtectionGroupHostParams{*mssqlVolumeProtectionGroupHostParamsModel}
	model.BackupDbVolumesOnly = core.BoolPtr(true)
	model.IncrementalBackupAfterRestart = core.BoolPtr(true)
	model.IndexingPolicy = indexingPolicyModel
	model.Objects = []backuprecoveryv1.MSSQLVolumeProtectionGroupObjectParams{*mssqlVolumeProtectionGroupObjectParamsModel}

	result, err := backuprecovery.ResourceIbmBaasProtectionGroupMSSQLVolumeProtectionGroupParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasProtectionGroupMSSQLVolumeProtectionGroupHostParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["enable_system_backup"] = true
		model["host_id"] = int(26)
		model["host_name"] = "testString"
		model["volume_guids"] = []string{"testString"}

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.MSSQLVolumeProtectionGroupHostParams)
	model.EnableSystemBackup = core.BoolPtr(true)
	model.HostID = core.Int64Ptr(int64(26))
	model.HostName = core.StringPtr("testString")
	model.VolumeGuids = []string{"testString"}

	result, err := backuprecovery.ResourceIbmBaasProtectionGroupMSSQLVolumeProtectionGroupHostParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasProtectionGroupMSSQLVolumeProtectionGroupObjectParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["id"] = int(26)
		model["name"] = "testString"
		model["source_type"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.MSSQLVolumeProtectionGroupObjectParams)
	model.ID = core.Int64Ptr(int64(26))
	model.Name = core.StringPtr("testString")
	model.SourceType = core.StringPtr("testString")

	result, err := backuprecovery.ResourceIbmBaasProtectionGroupMSSQLVolumeProtectionGroupObjectParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasProtectionGroupProtectionGroupRunToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		clusterIdentifierModel := make(map[string]interface{})
		clusterIdentifierModel["cluster_id"] = int(26)
		clusterIdentifierModel["cluster_incarnation_id"] = int(26)

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

		objectRunResultModel := make(map[string]interface{})
		objectRunResultModel["object"] = []map[string]interface{}{objectSummaryModel}
		objectRunResultModel["local_snapshot_info"] = []map[string]interface{}{backupRunModel}
		objectRunResultModel["original_backup_info"] = []map[string]interface{}{backupRunModel}
		objectRunResultModel["replication_info"] = []map[string]interface{}{replicationRunModel}
		objectRunResultModel["archival_info"] = []map[string]interface{}{archivalRunModel}
		objectRunResultModel["cloud_spin_info"] = []map[string]interface{}{cloudSpinRunModel}
		objectRunResultModel["on_legal_hold"] = true

		backupRunSummaryModel := make(map[string]interface{})
		backupRunSummaryModel["run_type"] = "kRegular"
		backupRunSummaryModel["is_sla_violated"] = true
		backupRunSummaryModel["start_time_usecs"] = int(26)
		backupRunSummaryModel["end_time_usecs"] = int(26)
		backupRunSummaryModel["status"] = "Accepted"
		backupRunSummaryModel["messages"] = []string{"testString"}
		backupRunSummaryModel["successful_objects_count"] = int(26)
		backupRunSummaryModel["skipped_objects_count"] = int(26)
		backupRunSummaryModel["failed_objects_count"] = int(26)
		backupRunSummaryModel["cancelled_objects_count"] = int(26)
		backupRunSummaryModel["successful_app_objects_count"] = int(38)
		backupRunSummaryModel["failed_app_objects_count"] = int(38)
		backupRunSummaryModel["cancelled_app_objects_count"] = int(38)
		backupRunSummaryModel["local_snapshot_stats"] = []map[string]interface{}{backupDataStatsModel}
		backupRunSummaryModel["indexing_task_id"] = "testString"
		backupRunSummaryModel["progress_task_id"] = "testString"
		backupRunSummaryModel["stats_task_id"] = "testString"
		backupRunSummaryModel["data_lock"] = "Compliance"
		backupRunSummaryModel["local_task_id"] = "testString"
		backupRunSummaryModel["data_lock_constraints"] = []map[string]interface{}{dataLockConstraintsModel}

		replicationRunSummaryModel := make(map[string]interface{})
		replicationRunSummaryModel["replication_target_results"] = []map[string]interface{}{replicationTargetResultModel}

		archivalRunSummaryModel := make(map[string]interface{})
		archivalRunSummaryModel["archival_target_results"] = []map[string]interface{}{archivalTargetResultModel}

		cloudSpinRunSummaryModel := make(map[string]interface{})
		cloudSpinRunSummaryModel["cloud_spin_target_results"] = []map[string]interface{}{cloudSpinTargetResultModel}

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
		model["id"] = "testString"
		model["protection_group_instance_id"] = int(26)
		model["protection_group_id"] = "testString"
		model["is_replication_run"] = true
		model["origin_cluster_identifier"] = []map[string]interface{}{clusterIdentifierModel}
		model["origin_protection_group_id"] = "testString"
		model["protection_group_name"] = "testString"
		model["is_local_snapshots_deleted"] = true
		model["objects"] = []map[string]interface{}{objectRunResultModel}
		model["local_backup_info"] = []map[string]interface{}{backupRunSummaryModel}
		model["original_backup_info"] = []map[string]interface{}{backupRunSummaryModel}
		model["replication_info"] = []map[string]interface{}{replicationRunSummaryModel}
		model["archival_info"] = []map[string]interface{}{archivalRunSummaryModel}
		model["cloud_spin_info"] = []map[string]interface{}{cloudSpinRunSummaryModel}
		model["on_legal_hold"] = true
		model["permissions"] = []map[string]interface{}{tenantModel}
		model["is_cloud_archival_direct"] = true
		model["has_local_snapshot"] = true
		model["environment"] = "testString"
		model["externally_triggered_backup_tag"] = "testString"

		assert.Equal(t, result, model)
	}

	clusterIdentifierModel := new(backuprecoveryv1.ClusterIdentifier)
	clusterIdentifierModel.ClusterID = core.Int64Ptr(int64(26))
	clusterIdentifierModel.ClusterIncarnationID = core.Int64Ptr(int64(26))

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

	objectRunResultModel := new(backuprecoveryv1.ObjectRunResult)
	objectRunResultModel.Object = objectSummaryModel
	objectRunResultModel.LocalSnapshotInfo = backupRunModel
	objectRunResultModel.OriginalBackupInfo = backupRunModel
	objectRunResultModel.ReplicationInfo = replicationRunModel
	objectRunResultModel.ArchivalInfo = archivalRunModel
	objectRunResultModel.CloudSpinInfo = cloudSpinRunModel
	objectRunResultModel.OnLegalHold = core.BoolPtr(true)

	backupRunSummaryModel := new(backuprecoveryv1.BackupRunSummary)
	backupRunSummaryModel.RunType = core.StringPtr("kRegular")
	backupRunSummaryModel.IsSlaViolated = core.BoolPtr(true)
	backupRunSummaryModel.StartTimeUsecs = core.Int64Ptr(int64(26))
	backupRunSummaryModel.EndTimeUsecs = core.Int64Ptr(int64(26))
	backupRunSummaryModel.Status = core.StringPtr("Accepted")
	backupRunSummaryModel.Messages = []string{"testString"}
	backupRunSummaryModel.SuccessfulObjectsCount = core.Int64Ptr(int64(26))
	backupRunSummaryModel.SkippedObjectsCount = core.Int64Ptr(int64(26))
	backupRunSummaryModel.FailedObjectsCount = core.Int64Ptr(int64(26))
	backupRunSummaryModel.CancelledObjectsCount = core.Int64Ptr(int64(26))
	backupRunSummaryModel.SuccessfulAppObjectsCount = core.Int64Ptr(int64(38))
	backupRunSummaryModel.FailedAppObjectsCount = core.Int64Ptr(int64(38))
	backupRunSummaryModel.CancelledAppObjectsCount = core.Int64Ptr(int64(38))
	backupRunSummaryModel.LocalSnapshotStats = backupDataStatsModel
	backupRunSummaryModel.IndexingTaskID = core.StringPtr("testString")
	backupRunSummaryModel.ProgressTaskID = core.StringPtr("testString")
	backupRunSummaryModel.StatsTaskID = core.StringPtr("testString")
	backupRunSummaryModel.DataLock = core.StringPtr("Compliance")
	backupRunSummaryModel.LocalTaskID = core.StringPtr("testString")
	backupRunSummaryModel.DataLockConstraints = dataLockConstraintsModel

	replicationRunSummaryModel := new(backuprecoveryv1.ReplicationRunSummary)
	replicationRunSummaryModel.ReplicationTargetResults = []backuprecoveryv1.ReplicationTargetResult{*replicationTargetResultModel}

	archivalRunSummaryModel := new(backuprecoveryv1.ArchivalRunSummary)
	archivalRunSummaryModel.ArchivalTargetResults = []backuprecoveryv1.ArchivalTargetResult{*archivalTargetResultModel}

	cloudSpinRunSummaryModel := new(backuprecoveryv1.CloudSpinRunSummary)
	cloudSpinRunSummaryModel.CloudSpinTargetResults = []backuprecoveryv1.CloudSpinTargetResult{*cloudSpinTargetResultModel}

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

	model := new(backuprecoveryv1.ProtectionGroupRun)
	model.ID = core.StringPtr("testString")
	model.ProtectionGroupInstanceID = core.Int64Ptr(int64(26))
	model.ProtectionGroupID = core.StringPtr("testString")
	model.IsReplicationRun = core.BoolPtr(true)
	model.OriginClusterIdentifier = clusterIdentifierModel
	model.OriginProtectionGroupID = core.StringPtr("testString")
	model.ProtectionGroupName = core.StringPtr("testString")
	model.IsLocalSnapshotsDeleted = core.BoolPtr(true)
	model.Objects = []backuprecoveryv1.ObjectRunResult{*objectRunResultModel}
	model.LocalBackupInfo = backupRunSummaryModel
	model.OriginalBackupInfo = backupRunSummaryModel
	model.ReplicationInfo = replicationRunSummaryModel
	model.ArchivalInfo = archivalRunSummaryModel
	model.CloudSpinInfo = cloudSpinRunSummaryModel
	model.OnLegalHold = core.BoolPtr(true)
	model.Permissions = []backuprecoveryv1.Tenant{*tenantModel}
	model.IsCloudArchivalDirect = core.BoolPtr(true)
	model.HasLocalSnapshot = core.BoolPtr(true)
	model.Environment = core.StringPtr("testString")
	model.ExternallyTriggeredBackupTag = core.StringPtr("testString")

	result, err := backuprecovery.ResourceIbmBaasProtectionGroupProtectionGroupRunToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasProtectionGroupClusterIdentifierToMap(t *testing.T) {
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

	result, err := backuprecovery.ResourceIbmBaasProtectionGroupClusterIdentifierToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasProtectionGroupObjectRunResultToMap(t *testing.T) {
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

	result, err := backuprecovery.ResourceIbmBaasProtectionGroupObjectRunResultToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasProtectionGroupObjectSummaryToMap(t *testing.T) {
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

	result, err := backuprecovery.ResourceIbmBaasProtectionGroupObjectSummaryToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasProtectionGroupSharepointObjectParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["site_web_url"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.SharepointObjectParams)
	model.SiteWebURL = core.StringPtr("testString")

	result, err := backuprecovery.ResourceIbmBaasProtectionGroupSharepointObjectParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasProtectionGroupObjectTypeVCenterParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["is_cloud_env"] = true

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.ObjectTypeVCenterParams)
	model.IsCloudEnv = core.BoolPtr(true)

	result, err := backuprecovery.ResourceIbmBaasProtectionGroupObjectTypeVCenterParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasProtectionGroupObjectTypeWindowsClusterParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["cluster_source_type"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.ObjectTypeWindowsClusterParams)
	model.ClusterSourceType = core.StringPtr("testString")

	result, err := backuprecovery.ResourceIbmBaasProtectionGroupObjectTypeWindowsClusterParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasProtectionGroupBackupRunToMap(t *testing.T) {
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

	result, err := backuprecovery.ResourceIbmBaasProtectionGroupBackupRunToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasProtectionGroupSnapshotInfoToMap(t *testing.T) {
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

	result, err := backuprecovery.ResourceIbmBaasProtectionGroupSnapshotInfoToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasProtectionGroupBackupDataStatsToMap(t *testing.T) {
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

	result, err := backuprecovery.ResourceIbmBaasProtectionGroupBackupDataStatsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasProtectionGroupDataLockConstraintsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["mode"] = "Compliance"
		model["expiry_time_usecs"] = int(26)

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.DataLockConstraints)
	model.Mode = core.StringPtr("Compliance")
	model.ExpiryTimeUsecs = core.Int64Ptr(int64(26))

	result, err := backuprecovery.ResourceIbmBaasProtectionGroupDataLockConstraintsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasProtectionGroupBackupAttemptToMap(t *testing.T) {
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

	result, err := backuprecovery.ResourceIbmBaasProtectionGroupBackupAttemptToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasProtectionGroupReplicationRunToMap(t *testing.T) {
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

	result, err := backuprecovery.ResourceIbmBaasProtectionGroupReplicationRunToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasProtectionGroupReplicationTargetResultToMap(t *testing.T) {
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

	result, err := backuprecovery.ResourceIbmBaasProtectionGroupReplicationTargetResultToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasProtectionGroupAWSTargetConfigToMap(t *testing.T) {
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

	result, err := backuprecovery.ResourceIbmBaasProtectionGroupAWSTargetConfigToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasProtectionGroupAzureTargetConfigToMap(t *testing.T) {
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

	result, err := backuprecovery.ResourceIbmBaasProtectionGroupAzureTargetConfigToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasProtectionGroupReplicationDataStatsToMap(t *testing.T) {
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

	result, err := backuprecovery.ResourceIbmBaasProtectionGroupReplicationDataStatsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasProtectionGroupArchivalRunToMap(t *testing.T) {
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

	result, err := backuprecovery.ResourceIbmBaasProtectionGroupArchivalRunToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasProtectionGroupArchivalTargetResultToMap(t *testing.T) {
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

	result, err := backuprecovery.ResourceIbmBaasProtectionGroupArchivalTargetResultToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasProtectionGroupArchivalTargetTierInfoToMap(t *testing.T) {
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

	result, err := backuprecovery.ResourceIbmBaasProtectionGroupArchivalTargetTierInfoToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasProtectionGroupAWSTiersToMap(t *testing.T) {
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

	result, err := backuprecovery.ResourceIbmBaasProtectionGroupAWSTiersToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasProtectionGroupAWSTierToMap(t *testing.T) {
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

	result, err := backuprecovery.ResourceIbmBaasProtectionGroupAWSTierToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasProtectionGroupAzureTiersToMap(t *testing.T) {
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

	result, err := backuprecovery.ResourceIbmBaasProtectionGroupAzureTiersToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasProtectionGroupAzureTierToMap(t *testing.T) {
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

	result, err := backuprecovery.ResourceIbmBaasProtectionGroupAzureTierToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasProtectionGroupGoogleTiersToMap(t *testing.T) {
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

	result, err := backuprecovery.ResourceIbmBaasProtectionGroupGoogleTiersToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasProtectionGroupGoogleTierToMap(t *testing.T) {
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

	result, err := backuprecovery.ResourceIbmBaasProtectionGroupGoogleTierToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasProtectionGroupOracleTiersToMap(t *testing.T) {
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

	result, err := backuprecovery.ResourceIbmBaasProtectionGroupOracleTiersToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasProtectionGroupOracleTierToMap(t *testing.T) {
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

	result, err := backuprecovery.ResourceIbmBaasProtectionGroupOracleTierToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasProtectionGroupArchivalDataStatsToMap(t *testing.T) {
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

	result, err := backuprecovery.ResourceIbmBaasProtectionGroupArchivalDataStatsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasProtectionGroupWormPropertiesToMap(t *testing.T) {
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

	result, err := backuprecovery.ResourceIbmBaasProtectionGroupWormPropertiesToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasProtectionGroupCloudSpinRunToMap(t *testing.T) {
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

	result, err := backuprecovery.ResourceIbmBaasProtectionGroupCloudSpinRunToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasProtectionGroupCloudSpinTargetResultToMap(t *testing.T) {
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

	result, err := backuprecovery.ResourceIbmBaasProtectionGroupCloudSpinTargetResultToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasProtectionGroupAwsCloudSpinParamsToMap(t *testing.T) {
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

	result, err := backuprecovery.ResourceIbmBaasProtectionGroupAwsCloudSpinParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasProtectionGroupCustomTagParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["key"] = "testString"
		model["value"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.CustomTagParams)
	model.Key = core.StringPtr("testString")
	model.Value = core.StringPtr("testString")

	result, err := backuprecovery.ResourceIbmBaasProtectionGroupCustomTagParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasProtectionGroupAzureCloudSpinParamsToMap(t *testing.T) {
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

	result, err := backuprecovery.ResourceIbmBaasProtectionGroupAzureCloudSpinParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasProtectionGroupCloudSpinDataStatsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["physical_bytes_transferred"] = int(26)

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.CloudSpinDataStats)
	model.PhysicalBytesTransferred = core.Int64Ptr(int64(26))

	result, err := backuprecovery.ResourceIbmBaasProtectionGroupCloudSpinDataStatsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasProtectionGroupBackupRunSummaryToMap(t *testing.T) {
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

	result, err := backuprecovery.ResourceIbmBaasProtectionGroupBackupRunSummaryToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasProtectionGroupReplicationRunSummaryToMap(t *testing.T) {
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

	result, err := backuprecovery.ResourceIbmBaasProtectionGroupReplicationRunSummaryToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasProtectionGroupArchivalRunSummaryToMap(t *testing.T) {
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

	result, err := backuprecovery.ResourceIbmBaasProtectionGroupArchivalRunSummaryToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasProtectionGroupCloudSpinRunSummaryToMap(t *testing.T) {
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

	result, err := backuprecovery.ResourceIbmBaasProtectionGroupCloudSpinRunSummaryToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasProtectionGroupTenantToMap(t *testing.T) {
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

	result, err := backuprecovery.ResourceIbmBaasProtectionGroupTenantToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasProtectionGroupExternalVendorTenantMetadataToMap(t *testing.T) {
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

	result, err := backuprecovery.ResourceIbmBaasProtectionGroupExternalVendorTenantMetadataToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasProtectionGroupIbmTenantMetadataParamsToMap(t *testing.T) {
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

	result, err := backuprecovery.ResourceIbmBaasProtectionGroupIbmTenantMetadataParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasProtectionGroupExternalVendorCustomPropertiesToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["key"] = "testString"
		model["value"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.ExternalVendorCustomProperties)
	model.Key = core.StringPtr("testString")
	model.Value = core.StringPtr("testString")

	result, err := backuprecovery.ResourceIbmBaasProtectionGroupExternalVendorCustomPropertiesToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasProtectionGroupTenantNetworkToMap(t *testing.T) {
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

	result, err := backuprecovery.ResourceIbmBaasProtectionGroupTenantNetworkToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasProtectionGroupMissingEntityParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["id"] = int(26)
		model["name"] = "testString"
		model["parent_source_id"] = int(26)
		model["parent_source_name"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.MissingEntityParams)
	model.ID = core.Int64Ptr(int64(26))
	model.Name = core.StringPtr("testString")
	model.ParentSourceID = core.Int64Ptr(int64(26))
	model.ParentSourceName = core.StringPtr("testString")

	result, err := backuprecovery.ResourceIbmBaasProtectionGroupMissingEntityParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasProtectionGroupMapToTimeOfDay(t *testing.T) {
	checkResult := func(result *backuprecoveryv1.TimeOfDay) {
		model := new(backuprecoveryv1.TimeOfDay)
		model.Hour = core.Int64Ptr(int64(0))
		model.Minute = core.Int64Ptr(int64(0))
		model.TimeZone = core.StringPtr("America/Los_Angeles")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["hour"] = int(0)
	model["minute"] = int(0)
	model["time_zone"] = "America/Los_Angeles"

	result, err := backuprecovery.ResourceIbmBaasProtectionGroupMapToTimeOfDay(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasProtectionGroupMapToProtectionGroupAlertingPolicy(t *testing.T) {
	checkResult := func(result *backuprecoveryv1.ProtectionGroupAlertingPolicy) {
		alertTargetModel := new(backuprecoveryv1.AlertTarget)
		alertTargetModel.EmailAddress = core.StringPtr("testString")
		alertTargetModel.Language = core.StringPtr("en-us")
		alertTargetModel.RecipientType = core.StringPtr("kTo")

		model := new(backuprecoveryv1.ProtectionGroupAlertingPolicy)
		model.BackupRunStatus = []string{"kSuccess"}
		model.AlertTargets = []backuprecoveryv1.AlertTarget{*alertTargetModel}
		model.RaiseObjectLevelFailureAlert = core.BoolPtr(true)
		model.RaiseObjectLevelFailureAlertAfterLastAttempt = core.BoolPtr(true)
		model.RaiseObjectLevelFailureAlertAfterEachAttempt = core.BoolPtr(true)

		assert.Equal(t, result, model)
	}

	alertTargetModel := make(map[string]interface{})
	alertTargetModel["email_address"] = "testString"
	alertTargetModel["language"] = "en-us"
	alertTargetModel["recipient_type"] = "kTo"

	model := make(map[string]interface{})
	model["backup_run_status"] = []interface{}{"kSuccess"}
	model["alert_targets"] = []interface{}{alertTargetModel}
	model["raise_object_level_failure_alert"] = true
	model["raise_object_level_failure_alert_after_last_attempt"] = true
	model["raise_object_level_failure_alert_after_each_attempt"] = true

	result, err := backuprecovery.ResourceIbmBaasProtectionGroupMapToProtectionGroupAlertingPolicy(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasProtectionGroupMapToAlertTarget(t *testing.T) {
	checkResult := func(result *backuprecoveryv1.AlertTarget) {
		model := new(backuprecoveryv1.AlertTarget)
		model.EmailAddress = core.StringPtr("testString")
		model.Language = core.StringPtr("en-us")
		model.RecipientType = core.StringPtr("kTo")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["email_address"] = "testString"
	model["language"] = "en-us"
	model["recipient_type"] = "kTo"

	result, err := backuprecovery.ResourceIbmBaasProtectionGroupMapToAlertTarget(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasProtectionGroupMapToSlaRule(t *testing.T) {
	checkResult := func(result *backuprecoveryv1.SlaRule) {
		model := new(backuprecoveryv1.SlaRule)
		model.BackupRunType = core.StringPtr("kIncremental")
		model.SlaMinutes = core.Int64Ptr(int64(1))

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["backup_run_type"] = "kIncremental"
	model["sla_minutes"] = int(1)

	result, err := backuprecovery.ResourceIbmBaasProtectionGroupMapToSlaRule(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasProtectionGroupMapToKeyValuePair(t *testing.T) {
	checkResult := func(result *backuprecoveryv1.KeyValuePair) {
		model := new(backuprecoveryv1.KeyValuePair)
		model.Key = core.StringPtr("testString")
		model.Value = core.StringPtr("testString")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["key"] = "testString"
	model["value"] = "testString"

	result, err := backuprecovery.ResourceIbmBaasProtectionGroupMapToKeyValuePair(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasProtectionGroupMapToPhysicalProtectionGroupParams(t *testing.T) {
	checkResult := func(result *backuprecoveryv1.PhysicalProtectionGroupParams) {
		physicalVolumeProtectionGroupObjectParamsModel := new(backuprecoveryv1.PhysicalVolumeProtectionGroupObjectParams)
		physicalVolumeProtectionGroupObjectParamsModel.ID = core.Int64Ptr(int64(26))
		physicalVolumeProtectionGroupObjectParamsModel.Name = core.StringPtr("testString")
		physicalVolumeProtectionGroupObjectParamsModel.VolumeGuids = []string{"testString"}
		physicalVolumeProtectionGroupObjectParamsModel.EnableSystemBackup = core.BoolPtr(true)
		physicalVolumeProtectionGroupObjectParamsModel.ExcludedVssWriters = []string{"testString"}

		indexingPolicyModel := new(backuprecoveryv1.IndexingPolicy)
		indexingPolicyModel.EnableIndexing = core.BoolPtr(true)
		indexingPolicyModel.IncludePaths = []string{"testString"}
		indexingPolicyModel.ExcludePaths = []string{"testString"}

		commonPreBackupScriptParamsModel := new(backuprecoveryv1.CommonPreBackupScriptParams)
		commonPreBackupScriptParamsModel.Path = core.StringPtr("testString")
		commonPreBackupScriptParamsModel.Params = core.StringPtr("testString")
		commonPreBackupScriptParamsModel.TimeoutSecs = core.Int64Ptr(int64(1))
		commonPreBackupScriptParamsModel.IsActive = core.BoolPtr(true)
		commonPreBackupScriptParamsModel.ContinueOnError = core.BoolPtr(true)

		commonPostBackupScriptParamsModel := new(backuprecoveryv1.CommonPostBackupScriptParams)
		commonPostBackupScriptParamsModel.Path = core.StringPtr("testString")
		commonPostBackupScriptParamsModel.Params = core.StringPtr("testString")
		commonPostBackupScriptParamsModel.TimeoutSecs = core.Int64Ptr(int64(1))
		commonPostBackupScriptParamsModel.IsActive = core.BoolPtr(true)

		prePostScriptParamsModel := new(backuprecoveryv1.PrePostScriptParams)
		prePostScriptParamsModel.PreScript = commonPreBackupScriptParamsModel
		prePostScriptParamsModel.PostScript = commonPostBackupScriptParamsModel

		physicalVolumeProtectionGroupParamsModel := new(backuprecoveryv1.PhysicalVolumeProtectionGroupParams)
		physicalVolumeProtectionGroupParamsModel.Objects = []backuprecoveryv1.PhysicalVolumeProtectionGroupObjectParams{*physicalVolumeProtectionGroupObjectParamsModel}
		physicalVolumeProtectionGroupParamsModel.IndexingPolicy = indexingPolicyModel
		physicalVolumeProtectionGroupParamsModel.PerformSourceSideDeduplication = core.BoolPtr(true)
		physicalVolumeProtectionGroupParamsModel.Quiesce = core.BoolPtr(true)
		physicalVolumeProtectionGroupParamsModel.ContinueOnQuiesceFailure = core.BoolPtr(true)
		physicalVolumeProtectionGroupParamsModel.IncrementalBackupAfterRestart = core.BoolPtr(true)
		physicalVolumeProtectionGroupParamsModel.PrePostScript = prePostScriptParamsModel
		physicalVolumeProtectionGroupParamsModel.DedupExclusionSourceIds = []int64{int64(26)}
		physicalVolumeProtectionGroupParamsModel.ExcludedVssWriters = []string{"testString"}
		physicalVolumeProtectionGroupParamsModel.CobmrBackup = core.BoolPtr(true)

		physicalFileBackupPathParamsModel := new(backuprecoveryv1.PhysicalFileBackupPathParams)
		physicalFileBackupPathParamsModel.IncludedPath = core.StringPtr("testString")
		physicalFileBackupPathParamsModel.ExcludedPaths = []string{"testString"}
		physicalFileBackupPathParamsModel.SkipNestedVolumes = core.BoolPtr(true)

		physicalFileProtectionGroupObjectParamsModel := new(backuprecoveryv1.PhysicalFileProtectionGroupObjectParams)
		physicalFileProtectionGroupObjectParamsModel.ExcludedVssWriters = []string{"testString"}
		physicalFileProtectionGroupObjectParamsModel.ID = core.Int64Ptr(int64(26))
		physicalFileProtectionGroupObjectParamsModel.FilePaths = []backuprecoveryv1.PhysicalFileBackupPathParams{*physicalFileBackupPathParamsModel}
		physicalFileProtectionGroupObjectParamsModel.UsesPathLevelSkipNestedVolumeSetting = core.BoolPtr(true)
		physicalFileProtectionGroupObjectParamsModel.NestedVolumeTypesToSkip = []string{"testString"}
		physicalFileProtectionGroupObjectParamsModel.FollowNasSymlinkTarget = core.BoolPtr(true)
		physicalFileProtectionGroupObjectParamsModel.MetadataFilePath = core.StringPtr("testString")

		cancellationTimeoutParamsModel := new(backuprecoveryv1.CancellationTimeoutParams)
		cancellationTimeoutParamsModel.TimeoutMins = core.Int64Ptr(int64(26))
		cancellationTimeoutParamsModel.BackupType = core.StringPtr("kRegular")

		physicalFileProtectionGroupParamsModel := new(backuprecoveryv1.PhysicalFileProtectionGroupParams)
		physicalFileProtectionGroupParamsModel.ExcludedVssWriters = []string{"testString"}
		physicalFileProtectionGroupParamsModel.Objects = []backuprecoveryv1.PhysicalFileProtectionGroupObjectParams{*physicalFileProtectionGroupObjectParamsModel}
		physicalFileProtectionGroupParamsModel.IndexingPolicy = indexingPolicyModel
		physicalFileProtectionGroupParamsModel.PerformSourceSideDeduplication = core.BoolPtr(true)
		physicalFileProtectionGroupParamsModel.PerformBrickBasedDeduplication = core.BoolPtr(true)
		physicalFileProtectionGroupParamsModel.TaskTimeouts = []backuprecoveryv1.CancellationTimeoutParams{*cancellationTimeoutParamsModel}
		physicalFileProtectionGroupParamsModel.Quiesce = core.BoolPtr(true)
		physicalFileProtectionGroupParamsModel.ContinueOnQuiesceFailure = core.BoolPtr(true)
		physicalFileProtectionGroupParamsModel.CobmrBackup = core.BoolPtr(true)
		physicalFileProtectionGroupParamsModel.PrePostScript = prePostScriptParamsModel
		physicalFileProtectionGroupParamsModel.DedupExclusionSourceIds = []int64{int64(26)}
		physicalFileProtectionGroupParamsModel.GlobalExcludePaths = []string{"testString"}
		physicalFileProtectionGroupParamsModel.GlobalExcludeFS = []string{"testString"}
		physicalFileProtectionGroupParamsModel.IgnorableErrors = []string{"kEOF"}
		physicalFileProtectionGroupParamsModel.AllowParallelRuns = core.BoolPtr(true)

		model := new(backuprecoveryv1.PhysicalProtectionGroupParams)
		model.ProtectionType = core.StringPtr("kFile")
		model.VolumeProtectionTypeParams = physicalVolumeProtectionGroupParamsModel
		model.FileProtectionTypeParams = physicalFileProtectionGroupParamsModel

		assert.Equal(t, result, model)
	}

	physicalVolumeProtectionGroupObjectParamsModel := make(map[string]interface{})
	physicalVolumeProtectionGroupObjectParamsModel["id"] = int(26)
	physicalVolumeProtectionGroupObjectParamsModel["name"] = "testString"
	physicalVolumeProtectionGroupObjectParamsModel["volume_guids"] = []interface{}{"testString"}
	physicalVolumeProtectionGroupObjectParamsModel["enable_system_backup"] = true
	physicalVolumeProtectionGroupObjectParamsModel["excluded_vss_writers"] = []interface{}{"testString"}

	indexingPolicyModel := make(map[string]interface{})
	indexingPolicyModel["enable_indexing"] = true
	indexingPolicyModel["include_paths"] = []interface{}{"testString"}
	indexingPolicyModel["exclude_paths"] = []interface{}{"testString"}

	commonPreBackupScriptParamsModel := make(map[string]interface{})
	commonPreBackupScriptParamsModel["path"] = "testString"
	commonPreBackupScriptParamsModel["params"] = "testString"
	commonPreBackupScriptParamsModel["timeout_secs"] = int(1)
	commonPreBackupScriptParamsModel["is_active"] = true
	commonPreBackupScriptParamsModel["continue_on_error"] = true

	commonPostBackupScriptParamsModel := make(map[string]interface{})
	commonPostBackupScriptParamsModel["path"] = "testString"
	commonPostBackupScriptParamsModel["params"] = "testString"
	commonPostBackupScriptParamsModel["timeout_secs"] = int(1)
	commonPostBackupScriptParamsModel["is_active"] = true

	prePostScriptParamsModel := make(map[string]interface{})
	prePostScriptParamsModel["pre_script"] = []interface{}{commonPreBackupScriptParamsModel}
	prePostScriptParamsModel["post_script"] = []interface{}{commonPostBackupScriptParamsModel}

	physicalVolumeProtectionGroupParamsModel := make(map[string]interface{})
	physicalVolumeProtectionGroupParamsModel["objects"] = []interface{}{physicalVolumeProtectionGroupObjectParamsModel}
	physicalVolumeProtectionGroupParamsModel["indexing_policy"] = []interface{}{indexingPolicyModel}
	physicalVolumeProtectionGroupParamsModel["perform_source_side_deduplication"] = true
	physicalVolumeProtectionGroupParamsModel["quiesce"] = true
	physicalVolumeProtectionGroupParamsModel["continue_on_quiesce_failure"] = true
	physicalVolumeProtectionGroupParamsModel["incremental_backup_after_restart"] = true
	physicalVolumeProtectionGroupParamsModel["pre_post_script"] = []interface{}{prePostScriptParamsModel}
	physicalVolumeProtectionGroupParamsModel["dedup_exclusion_source_ids"] = []interface{}{int(26)}
	physicalVolumeProtectionGroupParamsModel["excluded_vss_writers"] = []interface{}{"testString"}
	physicalVolumeProtectionGroupParamsModel["cobmr_backup"] = true

	physicalFileBackupPathParamsModel := make(map[string]interface{})
	physicalFileBackupPathParamsModel["included_path"] = "testString"
	physicalFileBackupPathParamsModel["excluded_paths"] = []interface{}{"testString"}
	physicalFileBackupPathParamsModel["skip_nested_volumes"] = true

	physicalFileProtectionGroupObjectParamsModel := make(map[string]interface{})
	physicalFileProtectionGroupObjectParamsModel["excluded_vss_writers"] = []interface{}{"testString"}
	physicalFileProtectionGroupObjectParamsModel["id"] = int(26)
	physicalFileProtectionGroupObjectParamsModel["file_paths"] = []interface{}{physicalFileBackupPathParamsModel}
	physicalFileProtectionGroupObjectParamsModel["uses_path_level_skip_nested_volume_setting"] = true
	physicalFileProtectionGroupObjectParamsModel["nested_volume_types_to_skip"] = []interface{}{"testString"}
	physicalFileProtectionGroupObjectParamsModel["follow_nas_symlink_target"] = true
	physicalFileProtectionGroupObjectParamsModel["metadata_file_path"] = "testString"

	cancellationTimeoutParamsModel := make(map[string]interface{})
	cancellationTimeoutParamsModel["timeout_mins"] = int(26)
	cancellationTimeoutParamsModel["backup_type"] = "kRegular"

	physicalFileProtectionGroupParamsModel := make(map[string]interface{})
	physicalFileProtectionGroupParamsModel["excluded_vss_writers"] = []interface{}{"testString"}
	physicalFileProtectionGroupParamsModel["objects"] = []interface{}{physicalFileProtectionGroupObjectParamsModel}
	physicalFileProtectionGroupParamsModel["indexing_policy"] = []interface{}{indexingPolicyModel}
	physicalFileProtectionGroupParamsModel["perform_source_side_deduplication"] = true
	physicalFileProtectionGroupParamsModel["perform_brick_based_deduplication"] = true
	physicalFileProtectionGroupParamsModel["task_timeouts"] = []interface{}{cancellationTimeoutParamsModel}
	physicalFileProtectionGroupParamsModel["quiesce"] = true
	physicalFileProtectionGroupParamsModel["continue_on_quiesce_failure"] = true
	physicalFileProtectionGroupParamsModel["cobmr_backup"] = true
	physicalFileProtectionGroupParamsModel["pre_post_script"] = []interface{}{prePostScriptParamsModel}
	physicalFileProtectionGroupParamsModel["dedup_exclusion_source_ids"] = []interface{}{int(26)}
	physicalFileProtectionGroupParamsModel["global_exclude_paths"] = []interface{}{"testString"}
	physicalFileProtectionGroupParamsModel["global_exclude_fs"] = []interface{}{"testString"}
	physicalFileProtectionGroupParamsModel["ignorable_errors"] = []interface{}{"kEOF"}
	physicalFileProtectionGroupParamsModel["allow_parallel_runs"] = true

	model := make(map[string]interface{})
	model["protection_type"] = "kFile"
	model["volume_protection_type_params"] = []interface{}{physicalVolumeProtectionGroupParamsModel}
	model["file_protection_type_params"] = []interface{}{physicalFileProtectionGroupParamsModel}

	result, err := backuprecovery.ResourceIbmBaasProtectionGroupMapToPhysicalProtectionGroupParams(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasProtectionGroupMapToPhysicalVolumeProtectionGroupParams(t *testing.T) {
	checkResult := func(result *backuprecoveryv1.PhysicalVolumeProtectionGroupParams) {
		physicalVolumeProtectionGroupObjectParamsModel := new(backuprecoveryv1.PhysicalVolumeProtectionGroupObjectParams)
		physicalVolumeProtectionGroupObjectParamsModel.ID = core.Int64Ptr(int64(26))
		physicalVolumeProtectionGroupObjectParamsModel.Name = core.StringPtr("testString")
		physicalVolumeProtectionGroupObjectParamsModel.VolumeGuids = []string{"testString"}
		physicalVolumeProtectionGroupObjectParamsModel.EnableSystemBackup = core.BoolPtr(true)
		physicalVolumeProtectionGroupObjectParamsModel.ExcludedVssWriters = []string{"testString"}

		indexingPolicyModel := new(backuprecoveryv1.IndexingPolicy)
		indexingPolicyModel.EnableIndexing = core.BoolPtr(true)
		indexingPolicyModel.IncludePaths = []string{"testString"}
		indexingPolicyModel.ExcludePaths = []string{"testString"}

		commonPreBackupScriptParamsModel := new(backuprecoveryv1.CommonPreBackupScriptParams)
		commonPreBackupScriptParamsModel.Path = core.StringPtr("testString")
		commonPreBackupScriptParamsModel.Params = core.StringPtr("testString")
		commonPreBackupScriptParamsModel.TimeoutSecs = core.Int64Ptr(int64(1))
		commonPreBackupScriptParamsModel.IsActive = core.BoolPtr(true)
		commonPreBackupScriptParamsModel.ContinueOnError = core.BoolPtr(true)

		commonPostBackupScriptParamsModel := new(backuprecoveryv1.CommonPostBackupScriptParams)
		commonPostBackupScriptParamsModel.Path = core.StringPtr("testString")
		commonPostBackupScriptParamsModel.Params = core.StringPtr("testString")
		commonPostBackupScriptParamsModel.TimeoutSecs = core.Int64Ptr(int64(1))
		commonPostBackupScriptParamsModel.IsActive = core.BoolPtr(true)

		prePostScriptParamsModel := new(backuprecoveryv1.PrePostScriptParams)
		prePostScriptParamsModel.PreScript = commonPreBackupScriptParamsModel
		prePostScriptParamsModel.PostScript = commonPostBackupScriptParamsModel

		model := new(backuprecoveryv1.PhysicalVolumeProtectionGroupParams)
		model.Objects = []backuprecoveryv1.PhysicalVolumeProtectionGroupObjectParams{*physicalVolumeProtectionGroupObjectParamsModel}
		model.IndexingPolicy = indexingPolicyModel
		model.PerformSourceSideDeduplication = core.BoolPtr(true)
		model.Quiesce = core.BoolPtr(true)
		model.ContinueOnQuiesceFailure = core.BoolPtr(true)
		model.IncrementalBackupAfterRestart = core.BoolPtr(true)
		model.PrePostScript = prePostScriptParamsModel
		model.DedupExclusionSourceIds = []int64{int64(26)}
		model.ExcludedVssWriters = []string{"testString"}
		model.CobmrBackup = core.BoolPtr(true)

		assert.Equal(t, result, model)
	}

	physicalVolumeProtectionGroupObjectParamsModel := make(map[string]interface{})
	physicalVolumeProtectionGroupObjectParamsModel["id"] = int(26)
	physicalVolumeProtectionGroupObjectParamsModel["name"] = "testString"
	physicalVolumeProtectionGroupObjectParamsModel["volume_guids"] = []interface{}{"testString"}
	physicalVolumeProtectionGroupObjectParamsModel["enable_system_backup"] = true
	physicalVolumeProtectionGroupObjectParamsModel["excluded_vss_writers"] = []interface{}{"testString"}

	indexingPolicyModel := make(map[string]interface{})
	indexingPolicyModel["enable_indexing"] = true
	indexingPolicyModel["include_paths"] = []interface{}{"testString"}
	indexingPolicyModel["exclude_paths"] = []interface{}{"testString"}

	commonPreBackupScriptParamsModel := make(map[string]interface{})
	commonPreBackupScriptParamsModel["path"] = "testString"
	commonPreBackupScriptParamsModel["params"] = "testString"
	commonPreBackupScriptParamsModel["timeout_secs"] = int(1)
	commonPreBackupScriptParamsModel["is_active"] = true
	commonPreBackupScriptParamsModel["continue_on_error"] = true

	commonPostBackupScriptParamsModel := make(map[string]interface{})
	commonPostBackupScriptParamsModel["path"] = "testString"
	commonPostBackupScriptParamsModel["params"] = "testString"
	commonPostBackupScriptParamsModel["timeout_secs"] = int(1)
	commonPostBackupScriptParamsModel["is_active"] = true

	prePostScriptParamsModel := make(map[string]interface{})
	prePostScriptParamsModel["pre_script"] = []interface{}{commonPreBackupScriptParamsModel}
	prePostScriptParamsModel["post_script"] = []interface{}{commonPostBackupScriptParamsModel}

	model := make(map[string]interface{})
	model["objects"] = []interface{}{physicalVolumeProtectionGroupObjectParamsModel}
	model["indexing_policy"] = []interface{}{indexingPolicyModel}
	model["perform_source_side_deduplication"] = true
	model["quiesce"] = true
	model["continue_on_quiesce_failure"] = true
	model["incremental_backup_after_restart"] = true
	model["pre_post_script"] = []interface{}{prePostScriptParamsModel}
	model["dedup_exclusion_source_ids"] = []interface{}{int(26)}
	model["excluded_vss_writers"] = []interface{}{"testString"}
	model["cobmr_backup"] = true

	result, err := backuprecovery.ResourceIbmBaasProtectionGroupMapToPhysicalVolumeProtectionGroupParams(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasProtectionGroupMapToPhysicalVolumeProtectionGroupObjectParams(t *testing.T) {
	checkResult := func(result *backuprecoveryv1.PhysicalVolumeProtectionGroupObjectParams) {
		model := new(backuprecoveryv1.PhysicalVolumeProtectionGroupObjectParams)
		model.ID = core.Int64Ptr(int64(26))
		model.Name = core.StringPtr("testString")
		model.VolumeGuids = []string{"testString"}
		model.EnableSystemBackup = core.BoolPtr(true)
		model.ExcludedVssWriters = []string{"testString"}

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["id"] = int(26)
	model["name"] = "testString"
	model["volume_guids"] = []interface{}{"testString"}
	model["enable_system_backup"] = true
	model["excluded_vss_writers"] = []interface{}{"testString"}

	result, err := backuprecovery.ResourceIbmBaasProtectionGroupMapToPhysicalVolumeProtectionGroupObjectParams(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasProtectionGroupMapToIndexingPolicy(t *testing.T) {
	checkResult := func(result *backuprecoveryv1.IndexingPolicy) {
		model := new(backuprecoveryv1.IndexingPolicy)
		model.EnableIndexing = core.BoolPtr(true)
		model.IncludePaths = []string{"testString"}
		model.ExcludePaths = []string{"testString"}

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["enable_indexing"] = true
	model["include_paths"] = []interface{}{"testString"}
	model["exclude_paths"] = []interface{}{"testString"}

	result, err := backuprecovery.ResourceIbmBaasProtectionGroupMapToIndexingPolicy(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasProtectionGroupMapToPrePostScriptParams(t *testing.T) {
	checkResult := func(result *backuprecoveryv1.PrePostScriptParams) {
		commonPreBackupScriptParamsModel := new(backuprecoveryv1.CommonPreBackupScriptParams)
		commonPreBackupScriptParamsModel.Path = core.StringPtr("testString")
		commonPreBackupScriptParamsModel.Params = core.StringPtr("testString")
		commonPreBackupScriptParamsModel.TimeoutSecs = core.Int64Ptr(int64(1))
		commonPreBackupScriptParamsModel.IsActive = core.BoolPtr(true)
		commonPreBackupScriptParamsModel.ContinueOnError = core.BoolPtr(true)

		commonPostBackupScriptParamsModel := new(backuprecoveryv1.CommonPostBackupScriptParams)
		commonPostBackupScriptParamsModel.Path = core.StringPtr("testString")
		commonPostBackupScriptParamsModel.Params = core.StringPtr("testString")
		commonPostBackupScriptParamsModel.TimeoutSecs = core.Int64Ptr(int64(1))
		commonPostBackupScriptParamsModel.IsActive = core.BoolPtr(true)

		model := new(backuprecoveryv1.PrePostScriptParams)
		model.PreScript = commonPreBackupScriptParamsModel
		model.PostScript = commonPostBackupScriptParamsModel

		assert.Equal(t, result, model)
	}

	commonPreBackupScriptParamsModel := make(map[string]interface{})
	commonPreBackupScriptParamsModel["path"] = "testString"
	commonPreBackupScriptParamsModel["params"] = "testString"
	commonPreBackupScriptParamsModel["timeout_secs"] = int(1)
	commonPreBackupScriptParamsModel["is_active"] = true
	commonPreBackupScriptParamsModel["continue_on_error"] = true

	commonPostBackupScriptParamsModel := make(map[string]interface{})
	commonPostBackupScriptParamsModel["path"] = "testString"
	commonPostBackupScriptParamsModel["params"] = "testString"
	commonPostBackupScriptParamsModel["timeout_secs"] = int(1)
	commonPostBackupScriptParamsModel["is_active"] = true

	model := make(map[string]interface{})
	model["pre_script"] = []interface{}{commonPreBackupScriptParamsModel}
	model["post_script"] = []interface{}{commonPostBackupScriptParamsModel}

	result, err := backuprecovery.ResourceIbmBaasProtectionGroupMapToPrePostScriptParams(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasProtectionGroupMapToCommonPreBackupScriptParams(t *testing.T) {
	checkResult := func(result *backuprecoveryv1.CommonPreBackupScriptParams) {
		model := new(backuprecoveryv1.CommonPreBackupScriptParams)
		model.Path = core.StringPtr("testString")
		model.Params = core.StringPtr("testString")
		model.TimeoutSecs = core.Int64Ptr(int64(1))
		model.IsActive = core.BoolPtr(true)
		model.ContinueOnError = core.BoolPtr(true)

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["path"] = "testString"
	model["params"] = "testString"
	model["timeout_secs"] = int(1)
	model["is_active"] = true
	model["continue_on_error"] = true

	result, err := backuprecovery.ResourceIbmBaasProtectionGroupMapToCommonPreBackupScriptParams(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasProtectionGroupMapToCommonPostBackupScriptParams(t *testing.T) {
	checkResult := func(result *backuprecoveryv1.CommonPostBackupScriptParams) {
		model := new(backuprecoveryv1.CommonPostBackupScriptParams)
		model.Path = core.StringPtr("testString")
		model.Params = core.StringPtr("testString")
		model.TimeoutSecs = core.Int64Ptr(int64(1))
		model.IsActive = core.BoolPtr(true)

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["path"] = "testString"
	model["params"] = "testString"
	model["timeout_secs"] = int(1)
	model["is_active"] = true

	result, err := backuprecovery.ResourceIbmBaasProtectionGroupMapToCommonPostBackupScriptParams(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasProtectionGroupMapToPhysicalFileProtectionGroupParams(t *testing.T) {
	checkResult := func(result *backuprecoveryv1.PhysicalFileProtectionGroupParams) {
		physicalFileBackupPathParamsModel := new(backuprecoveryv1.PhysicalFileBackupPathParams)
		physicalFileBackupPathParamsModel.IncludedPath = core.StringPtr("testString")
		physicalFileBackupPathParamsModel.ExcludedPaths = []string{"testString"}
		physicalFileBackupPathParamsModel.SkipNestedVolumes = core.BoolPtr(true)

		physicalFileProtectionGroupObjectParamsModel := new(backuprecoveryv1.PhysicalFileProtectionGroupObjectParams)
		physicalFileProtectionGroupObjectParamsModel.ExcludedVssWriters = []string{"testString"}
		physicalFileProtectionGroupObjectParamsModel.ID = core.Int64Ptr(int64(26))
		physicalFileProtectionGroupObjectParamsModel.FilePaths = []backuprecoveryv1.PhysicalFileBackupPathParams{*physicalFileBackupPathParamsModel}
		physicalFileProtectionGroupObjectParamsModel.UsesPathLevelSkipNestedVolumeSetting = core.BoolPtr(true)
		physicalFileProtectionGroupObjectParamsModel.NestedVolumeTypesToSkip = []string{"testString"}
		physicalFileProtectionGroupObjectParamsModel.FollowNasSymlinkTarget = core.BoolPtr(true)
		physicalFileProtectionGroupObjectParamsModel.MetadataFilePath = core.StringPtr("testString")

		indexingPolicyModel := new(backuprecoveryv1.IndexingPolicy)
		indexingPolicyModel.EnableIndexing = core.BoolPtr(true)
		indexingPolicyModel.IncludePaths = []string{"testString"}
		indexingPolicyModel.ExcludePaths = []string{"testString"}

		cancellationTimeoutParamsModel := new(backuprecoveryv1.CancellationTimeoutParams)
		cancellationTimeoutParamsModel.TimeoutMins = core.Int64Ptr(int64(26))
		cancellationTimeoutParamsModel.BackupType = core.StringPtr("kRegular")

		commonPreBackupScriptParamsModel := new(backuprecoveryv1.CommonPreBackupScriptParams)
		commonPreBackupScriptParamsModel.Path = core.StringPtr("testString")
		commonPreBackupScriptParamsModel.Params = core.StringPtr("testString")
		commonPreBackupScriptParamsModel.TimeoutSecs = core.Int64Ptr(int64(1))
		commonPreBackupScriptParamsModel.IsActive = core.BoolPtr(true)
		commonPreBackupScriptParamsModel.ContinueOnError = core.BoolPtr(true)

		commonPostBackupScriptParamsModel := new(backuprecoveryv1.CommonPostBackupScriptParams)
		commonPostBackupScriptParamsModel.Path = core.StringPtr("testString")
		commonPostBackupScriptParamsModel.Params = core.StringPtr("testString")
		commonPostBackupScriptParamsModel.TimeoutSecs = core.Int64Ptr(int64(1))
		commonPostBackupScriptParamsModel.IsActive = core.BoolPtr(true)

		prePostScriptParamsModel := new(backuprecoveryv1.PrePostScriptParams)
		prePostScriptParamsModel.PreScript = commonPreBackupScriptParamsModel
		prePostScriptParamsModel.PostScript = commonPostBackupScriptParamsModel

		model := new(backuprecoveryv1.PhysicalFileProtectionGroupParams)
		model.ExcludedVssWriters = []string{"testString"}
		model.Objects = []backuprecoveryv1.PhysicalFileProtectionGroupObjectParams{*physicalFileProtectionGroupObjectParamsModel}
		model.IndexingPolicy = indexingPolicyModel
		model.PerformSourceSideDeduplication = core.BoolPtr(true)
		model.PerformBrickBasedDeduplication = core.BoolPtr(true)
		model.TaskTimeouts = []backuprecoveryv1.CancellationTimeoutParams{*cancellationTimeoutParamsModel}
		model.Quiesce = core.BoolPtr(true)
		model.ContinueOnQuiesceFailure = core.BoolPtr(true)
		model.CobmrBackup = core.BoolPtr(true)
		model.PrePostScript = prePostScriptParamsModel
		model.DedupExclusionSourceIds = []int64{int64(26)}
		model.GlobalExcludePaths = []string{"testString"}
		model.GlobalExcludeFS = []string{"testString"}
		model.IgnorableErrors = []string{"kEOF"}
		model.AllowParallelRuns = core.BoolPtr(true)

		assert.Equal(t, result, model)
	}

	physicalFileBackupPathParamsModel := make(map[string]interface{})
	physicalFileBackupPathParamsModel["included_path"] = "testString"
	physicalFileBackupPathParamsModel["excluded_paths"] = []interface{}{"testString"}
	physicalFileBackupPathParamsModel["skip_nested_volumes"] = true

	physicalFileProtectionGroupObjectParamsModel := make(map[string]interface{})
	physicalFileProtectionGroupObjectParamsModel["excluded_vss_writers"] = []interface{}{"testString"}
	physicalFileProtectionGroupObjectParamsModel["id"] = int(26)
	physicalFileProtectionGroupObjectParamsModel["file_paths"] = []interface{}{physicalFileBackupPathParamsModel}
	physicalFileProtectionGroupObjectParamsModel["uses_path_level_skip_nested_volume_setting"] = true
	physicalFileProtectionGroupObjectParamsModel["nested_volume_types_to_skip"] = []interface{}{"testString"}
	physicalFileProtectionGroupObjectParamsModel["follow_nas_symlink_target"] = true
	physicalFileProtectionGroupObjectParamsModel["metadata_file_path"] = "testString"

	indexingPolicyModel := make(map[string]interface{})
	indexingPolicyModel["enable_indexing"] = true
	indexingPolicyModel["include_paths"] = []interface{}{"testString"}
	indexingPolicyModel["exclude_paths"] = []interface{}{"testString"}

	cancellationTimeoutParamsModel := make(map[string]interface{})
	cancellationTimeoutParamsModel["timeout_mins"] = int(26)
	cancellationTimeoutParamsModel["backup_type"] = "kRegular"

	commonPreBackupScriptParamsModel := make(map[string]interface{})
	commonPreBackupScriptParamsModel["path"] = "testString"
	commonPreBackupScriptParamsModel["params"] = "testString"
	commonPreBackupScriptParamsModel["timeout_secs"] = int(1)
	commonPreBackupScriptParamsModel["is_active"] = true
	commonPreBackupScriptParamsModel["continue_on_error"] = true

	commonPostBackupScriptParamsModel := make(map[string]interface{})
	commonPostBackupScriptParamsModel["path"] = "testString"
	commonPostBackupScriptParamsModel["params"] = "testString"
	commonPostBackupScriptParamsModel["timeout_secs"] = int(1)
	commonPostBackupScriptParamsModel["is_active"] = true

	prePostScriptParamsModel := make(map[string]interface{})
	prePostScriptParamsModel["pre_script"] = []interface{}{commonPreBackupScriptParamsModel}
	prePostScriptParamsModel["post_script"] = []interface{}{commonPostBackupScriptParamsModel}

	model := make(map[string]interface{})
	model["excluded_vss_writers"] = []interface{}{"testString"}
	model["objects"] = []interface{}{physicalFileProtectionGroupObjectParamsModel}
	model["indexing_policy"] = []interface{}{indexingPolicyModel}
	model["perform_source_side_deduplication"] = true
	model["perform_brick_based_deduplication"] = true
	model["task_timeouts"] = []interface{}{cancellationTimeoutParamsModel}
	model["quiesce"] = true
	model["continue_on_quiesce_failure"] = true
	model["cobmr_backup"] = true
	model["pre_post_script"] = []interface{}{prePostScriptParamsModel}
	model["dedup_exclusion_source_ids"] = []interface{}{int(26)}
	model["global_exclude_paths"] = []interface{}{"testString"}
	model["global_exclude_fs"] = []interface{}{"testString"}
	model["ignorable_errors"] = []interface{}{"kEOF"}
	model["allow_parallel_runs"] = true

	result, err := backuprecovery.ResourceIbmBaasProtectionGroupMapToPhysicalFileProtectionGroupParams(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasProtectionGroupMapToPhysicalFileProtectionGroupObjectParams(t *testing.T) {
	checkResult := func(result *backuprecoveryv1.PhysicalFileProtectionGroupObjectParams) {
		physicalFileBackupPathParamsModel := new(backuprecoveryv1.PhysicalFileBackupPathParams)
		physicalFileBackupPathParamsModel.IncludedPath = core.StringPtr("testString")
		physicalFileBackupPathParamsModel.ExcludedPaths = []string{"testString"}
		physicalFileBackupPathParamsModel.SkipNestedVolumes = core.BoolPtr(true)

		model := new(backuprecoveryv1.PhysicalFileProtectionGroupObjectParams)
		model.ExcludedVssWriters = []string{"testString"}
		model.ID = core.Int64Ptr(int64(26))
		model.Name = core.StringPtr("testString")
		model.FilePaths = []backuprecoveryv1.PhysicalFileBackupPathParams{*physicalFileBackupPathParamsModel}
		model.UsesPathLevelSkipNestedVolumeSetting = core.BoolPtr(true)
		model.NestedVolumeTypesToSkip = []string{"testString"}
		model.FollowNasSymlinkTarget = core.BoolPtr(true)
		model.MetadataFilePath = core.StringPtr("testString")

		assert.Equal(t, result, model)
	}

	physicalFileBackupPathParamsModel := make(map[string]interface{})
	physicalFileBackupPathParamsModel["included_path"] = "testString"
	physicalFileBackupPathParamsModel["excluded_paths"] = []interface{}{"testString"}
	physicalFileBackupPathParamsModel["skip_nested_volumes"] = true

	model := make(map[string]interface{})
	model["excluded_vss_writers"] = []interface{}{"testString"}
	model["id"] = int(26)
	model["name"] = "testString"
	model["file_paths"] = []interface{}{physicalFileBackupPathParamsModel}
	model["uses_path_level_skip_nested_volume_setting"] = true
	model["nested_volume_types_to_skip"] = []interface{}{"testString"}
	model["follow_nas_symlink_target"] = true
	model["metadata_file_path"] = "testString"

	result, err := backuprecovery.ResourceIbmBaasProtectionGroupMapToPhysicalFileProtectionGroupObjectParams(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasProtectionGroupMapToPhysicalFileBackupPathParams(t *testing.T) {
	checkResult := func(result *backuprecoveryv1.PhysicalFileBackupPathParams) {
		model := new(backuprecoveryv1.PhysicalFileBackupPathParams)
		model.IncludedPath = core.StringPtr("testString")
		model.ExcludedPaths = []string{"testString"}
		model.SkipNestedVolumes = core.BoolPtr(true)

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["included_path"] = "testString"
	model["excluded_paths"] = []interface{}{"testString"}
	model["skip_nested_volumes"] = true

	result, err := backuprecovery.ResourceIbmBaasProtectionGroupMapToPhysicalFileBackupPathParams(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasProtectionGroupMapToCancellationTimeoutParams(t *testing.T) {
	checkResult := func(result *backuprecoveryv1.CancellationTimeoutParams) {
		model := new(backuprecoveryv1.CancellationTimeoutParams)
		model.TimeoutMins = core.Int64Ptr(int64(26))
		model.BackupType = core.StringPtr("kRegular")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["timeout_mins"] = int(26)
	model["backup_type"] = "kRegular"

	result, err := backuprecovery.ResourceIbmBaasProtectionGroupMapToCancellationTimeoutParams(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasProtectionGroupMapToMSSQLProtectionGroupParams(t *testing.T) {
	checkResult := func(result *backuprecoveryv1.MSSQLProtectionGroupParams) {
		advancedSettingsModel := new(backuprecoveryv1.AdvancedSettings)
		advancedSettingsModel.ClonedDbBackupStatus = core.StringPtr("kError")
		advancedSettingsModel.DbBackupIfNotOnlineStatus = core.StringPtr("kError")
		advancedSettingsModel.MissingDbBackupStatus = core.StringPtr("kError")
		advancedSettingsModel.OfflineRestoringDbBackupStatus = core.StringPtr("kError")
		advancedSettingsModel.ReadOnlyDbBackupStatus = core.StringPtr("kError")
		advancedSettingsModel.ReportAllNonAutoprotectDbErrors = core.StringPtr("kError")

		filterModel := new(backuprecoveryv1.Filter)
		filterModel.FilterString = core.StringPtr("testString")
		filterModel.IsRegularExpression = core.BoolPtr(false)

		commonPreBackupScriptParamsModel := new(backuprecoveryv1.CommonPreBackupScriptParams)
		commonPreBackupScriptParamsModel.Path = core.StringPtr("testString")
		commonPreBackupScriptParamsModel.Params = core.StringPtr("testString")
		commonPreBackupScriptParamsModel.TimeoutSecs = core.Int64Ptr(int64(1))
		commonPreBackupScriptParamsModel.IsActive = core.BoolPtr(true)
		commonPreBackupScriptParamsModel.ContinueOnError = core.BoolPtr(true)

		commonPostBackupScriptParamsModel := new(backuprecoveryv1.CommonPostBackupScriptParams)
		commonPostBackupScriptParamsModel.Path = core.StringPtr("testString")
		commonPostBackupScriptParamsModel.Params = core.StringPtr("testString")
		commonPostBackupScriptParamsModel.TimeoutSecs = core.Int64Ptr(int64(1))
		commonPostBackupScriptParamsModel.IsActive = core.BoolPtr(true)

		prePostScriptParamsModel := new(backuprecoveryv1.PrePostScriptParams)
		prePostScriptParamsModel.PreScript = commonPreBackupScriptParamsModel
		prePostScriptParamsModel.PostScript = commonPostBackupScriptParamsModel

		mssqlFileProtectionGroupHostParamsModel := new(backuprecoveryv1.MSSQLFileProtectionGroupHostParams)
		mssqlFileProtectionGroupHostParamsModel.DisableSourceSideDeduplication = core.BoolPtr(true)
		mssqlFileProtectionGroupHostParamsModel.HostID = core.Int64Ptr(int64(26))

		mssqlFileProtectionGroupObjectParamsModel := new(backuprecoveryv1.MSSQLFileProtectionGroupObjectParams)
		mssqlFileProtectionGroupObjectParamsModel.ID = core.Int64Ptr(int64(26))

		mssqlFileProtectionGroupParamsModel := new(backuprecoveryv1.MSSQLFileProtectionGroupParams)
		mssqlFileProtectionGroupParamsModel.AagBackupPreferenceType = core.StringPtr("kPrimaryReplicaOnly")
		mssqlFileProtectionGroupParamsModel.AdvancedSettings = advancedSettingsModel
		mssqlFileProtectionGroupParamsModel.BackupSystemDbs = core.BoolPtr(true)
		mssqlFileProtectionGroupParamsModel.ExcludeFilters = []backuprecoveryv1.Filter{*filterModel}
		mssqlFileProtectionGroupParamsModel.FullBackupsCopyOnly = core.BoolPtr(true)
		mssqlFileProtectionGroupParamsModel.LogBackupNumStreams = core.Int64Ptr(int64(38))
		mssqlFileProtectionGroupParamsModel.LogBackupWithClause = core.StringPtr("testString")
		mssqlFileProtectionGroupParamsModel.PrePostScript = prePostScriptParamsModel
		mssqlFileProtectionGroupParamsModel.UseAagPreferencesFromServer = core.BoolPtr(true)
		mssqlFileProtectionGroupParamsModel.UserDbBackupPreferenceType = core.StringPtr("kBackupAllDatabases")
		mssqlFileProtectionGroupParamsModel.AdditionalHostParams = []backuprecoveryv1.MSSQLFileProtectionGroupHostParams{*mssqlFileProtectionGroupHostParamsModel}
		mssqlFileProtectionGroupParamsModel.Objects = []backuprecoveryv1.MSSQLFileProtectionGroupObjectParams{*mssqlFileProtectionGroupObjectParamsModel}
		mssqlFileProtectionGroupParamsModel.PerformSourceSideDeduplication = core.BoolPtr(true)

		mssqlNativeProtectionGroupObjectParamsModel := new(backuprecoveryv1.MSSQLNativeProtectionGroupObjectParams)
		mssqlNativeProtectionGroupObjectParamsModel.ID = core.Int64Ptr(int64(26))

		mssqlNativeProtectionGroupParamsModel := new(backuprecoveryv1.MSSQLNativeProtectionGroupParams)
		mssqlNativeProtectionGroupParamsModel.AagBackupPreferenceType = core.StringPtr("kPrimaryReplicaOnly")
		mssqlNativeProtectionGroupParamsModel.AdvancedSettings = advancedSettingsModel
		mssqlNativeProtectionGroupParamsModel.BackupSystemDbs = core.BoolPtr(true)
		mssqlNativeProtectionGroupParamsModel.ExcludeFilters = []backuprecoveryv1.Filter{*filterModel}
		mssqlNativeProtectionGroupParamsModel.FullBackupsCopyOnly = core.BoolPtr(true)
		mssqlNativeProtectionGroupParamsModel.LogBackupNumStreams = core.Int64Ptr(int64(38))
		mssqlNativeProtectionGroupParamsModel.LogBackupWithClause = core.StringPtr("testString")
		mssqlNativeProtectionGroupParamsModel.PrePostScript = prePostScriptParamsModel
		mssqlNativeProtectionGroupParamsModel.UseAagPreferencesFromServer = core.BoolPtr(true)
		mssqlNativeProtectionGroupParamsModel.UserDbBackupPreferenceType = core.StringPtr("kBackupAllDatabases")
		mssqlNativeProtectionGroupParamsModel.NumStreams = core.Int64Ptr(int64(38))
		mssqlNativeProtectionGroupParamsModel.Objects = []backuprecoveryv1.MSSQLNativeProtectionGroupObjectParams{*mssqlNativeProtectionGroupObjectParamsModel}
		mssqlNativeProtectionGroupParamsModel.WithClause = core.StringPtr("testString")

		mssqlVolumeProtectionGroupHostParamsModel := new(backuprecoveryv1.MSSQLVolumeProtectionGroupHostParams)
		mssqlVolumeProtectionGroupHostParamsModel.EnableSystemBackup = core.BoolPtr(true)
		mssqlVolumeProtectionGroupHostParamsModel.HostID = core.Int64Ptr(int64(26))
		mssqlVolumeProtectionGroupHostParamsModel.VolumeGuids = []string{"testString"}

		indexingPolicyModel := new(backuprecoveryv1.IndexingPolicy)
		indexingPolicyModel.EnableIndexing = core.BoolPtr(true)
		indexingPolicyModel.IncludePaths = []string{"testString"}
		indexingPolicyModel.ExcludePaths = []string{"testString"}

		mssqlVolumeProtectionGroupObjectParamsModel := new(backuprecoveryv1.MSSQLVolumeProtectionGroupObjectParams)
		mssqlVolumeProtectionGroupObjectParamsModel.ID = core.Int64Ptr(int64(26))

		mssqlVolumeProtectionGroupParamsModel := new(backuprecoveryv1.MSSQLVolumeProtectionGroupParams)
		mssqlVolumeProtectionGroupParamsModel.AagBackupPreferenceType = core.StringPtr("kPrimaryReplicaOnly")
		mssqlVolumeProtectionGroupParamsModel.AdvancedSettings = advancedSettingsModel
		mssqlVolumeProtectionGroupParamsModel.BackupSystemDbs = core.BoolPtr(true)
		mssqlVolumeProtectionGroupParamsModel.ExcludeFilters = []backuprecoveryv1.Filter{*filterModel}
		mssqlVolumeProtectionGroupParamsModel.FullBackupsCopyOnly = core.BoolPtr(true)
		mssqlVolumeProtectionGroupParamsModel.LogBackupNumStreams = core.Int64Ptr(int64(38))
		mssqlVolumeProtectionGroupParamsModel.LogBackupWithClause = core.StringPtr("testString")
		mssqlVolumeProtectionGroupParamsModel.PrePostScript = prePostScriptParamsModel
		mssqlVolumeProtectionGroupParamsModel.UseAagPreferencesFromServer = core.BoolPtr(true)
		mssqlVolumeProtectionGroupParamsModel.UserDbBackupPreferenceType = core.StringPtr("kBackupAllDatabases")
		mssqlVolumeProtectionGroupParamsModel.AdditionalHostParams = []backuprecoveryv1.MSSQLVolumeProtectionGroupHostParams{*mssqlVolumeProtectionGroupHostParamsModel}
		mssqlVolumeProtectionGroupParamsModel.BackupDbVolumesOnly = core.BoolPtr(true)
		mssqlVolumeProtectionGroupParamsModel.IncrementalBackupAfterRestart = core.BoolPtr(true)
		mssqlVolumeProtectionGroupParamsModel.IndexingPolicy = indexingPolicyModel
		mssqlVolumeProtectionGroupParamsModel.Objects = []backuprecoveryv1.MSSQLVolumeProtectionGroupObjectParams{*mssqlVolumeProtectionGroupObjectParamsModel}

		model := new(backuprecoveryv1.MSSQLProtectionGroupParams)
		model.FileProtectionTypeParams = mssqlFileProtectionGroupParamsModel
		model.NativeProtectionTypeParams = mssqlNativeProtectionGroupParamsModel
		model.ProtectionType = core.StringPtr("kFile")
		model.VolumeProtectionTypeParams = mssqlVolumeProtectionGroupParamsModel

		assert.Equal(t, result, model)
	}

	advancedSettingsModel := make(map[string]interface{})
	advancedSettingsModel["cloned_db_backup_status"] = "kError"
	advancedSettingsModel["db_backup_if_not_online_status"] = "kError"
	advancedSettingsModel["missing_db_backup_status"] = "kError"
	advancedSettingsModel["offline_restoring_db_backup_status"] = "kError"
	advancedSettingsModel["read_only_db_backup_status"] = "kError"
	advancedSettingsModel["report_all_non_autoprotect_db_errors"] = "kError"

	filterModel := make(map[string]interface{})
	filterModel["filter_string"] = "testString"
	filterModel["is_regular_expression"] = false

	commonPreBackupScriptParamsModel := make(map[string]interface{})
	commonPreBackupScriptParamsModel["path"] = "testString"
	commonPreBackupScriptParamsModel["params"] = "testString"
	commonPreBackupScriptParamsModel["timeout_secs"] = int(1)
	commonPreBackupScriptParamsModel["is_active"] = true
	commonPreBackupScriptParamsModel["continue_on_error"] = true

	commonPostBackupScriptParamsModel := make(map[string]interface{})
	commonPostBackupScriptParamsModel["path"] = "testString"
	commonPostBackupScriptParamsModel["params"] = "testString"
	commonPostBackupScriptParamsModel["timeout_secs"] = int(1)
	commonPostBackupScriptParamsModel["is_active"] = true

	prePostScriptParamsModel := make(map[string]interface{})
	prePostScriptParamsModel["pre_script"] = []interface{}{commonPreBackupScriptParamsModel}
	prePostScriptParamsModel["post_script"] = []interface{}{commonPostBackupScriptParamsModel}

	mssqlFileProtectionGroupHostParamsModel := make(map[string]interface{})
	mssqlFileProtectionGroupHostParamsModel["disable_source_side_deduplication"] = true
	mssqlFileProtectionGroupHostParamsModel["host_id"] = int(26)

	mssqlFileProtectionGroupObjectParamsModel := make(map[string]interface{})
	mssqlFileProtectionGroupObjectParamsModel["id"] = int(26)

	mssqlFileProtectionGroupParamsModel := make(map[string]interface{})
	mssqlFileProtectionGroupParamsModel["aag_backup_preference_type"] = "kPrimaryReplicaOnly"
	mssqlFileProtectionGroupParamsModel["advanced_settings"] = []interface{}{advancedSettingsModel}
	mssqlFileProtectionGroupParamsModel["backup_system_dbs"] = true
	mssqlFileProtectionGroupParamsModel["exclude_filters"] = []interface{}{filterModel}
	mssqlFileProtectionGroupParamsModel["full_backups_copy_only"] = true
	mssqlFileProtectionGroupParamsModel["log_backup_num_streams"] = int(38)
	mssqlFileProtectionGroupParamsModel["log_backup_with_clause"] = "testString"
	mssqlFileProtectionGroupParamsModel["pre_post_script"] = []interface{}{prePostScriptParamsModel}
	mssqlFileProtectionGroupParamsModel["use_aag_preferences_from_server"] = true
	mssqlFileProtectionGroupParamsModel["user_db_backup_preference_type"] = "kBackupAllDatabases"
	mssqlFileProtectionGroupParamsModel["additional_host_params"] = []interface{}{mssqlFileProtectionGroupHostParamsModel}
	mssqlFileProtectionGroupParamsModel["objects"] = []interface{}{mssqlFileProtectionGroupObjectParamsModel}
	mssqlFileProtectionGroupParamsModel["perform_source_side_deduplication"] = true

	mssqlNativeProtectionGroupObjectParamsModel := make(map[string]interface{})
	mssqlNativeProtectionGroupObjectParamsModel["id"] = int(26)

	mssqlNativeProtectionGroupParamsModel := make(map[string]interface{})
	mssqlNativeProtectionGroupParamsModel["aag_backup_preference_type"] = "kPrimaryReplicaOnly"
	mssqlNativeProtectionGroupParamsModel["advanced_settings"] = []interface{}{advancedSettingsModel}
	mssqlNativeProtectionGroupParamsModel["backup_system_dbs"] = true
	mssqlNativeProtectionGroupParamsModel["exclude_filters"] = []interface{}{filterModel}
	mssqlNativeProtectionGroupParamsModel["full_backups_copy_only"] = true
	mssqlNativeProtectionGroupParamsModel["log_backup_num_streams"] = int(38)
	mssqlNativeProtectionGroupParamsModel["log_backup_with_clause"] = "testString"
	mssqlNativeProtectionGroupParamsModel["pre_post_script"] = []interface{}{prePostScriptParamsModel}
	mssqlNativeProtectionGroupParamsModel["use_aag_preferences_from_server"] = true
	mssqlNativeProtectionGroupParamsModel["user_db_backup_preference_type"] = "kBackupAllDatabases"
	mssqlNativeProtectionGroupParamsModel["num_streams"] = int(38)
	mssqlNativeProtectionGroupParamsModel["objects"] = []interface{}{mssqlNativeProtectionGroupObjectParamsModel}
	mssqlNativeProtectionGroupParamsModel["with_clause"] = "testString"

	mssqlVolumeProtectionGroupHostParamsModel := make(map[string]interface{})
	mssqlVolumeProtectionGroupHostParamsModel["enable_system_backup"] = true
	mssqlVolumeProtectionGroupHostParamsModel["host_id"] = int(26)
	mssqlVolumeProtectionGroupHostParamsModel["volume_guids"] = []interface{}{"testString"}

	indexingPolicyModel := make(map[string]interface{})
	indexingPolicyModel["enable_indexing"] = true
	indexingPolicyModel["include_paths"] = []interface{}{"testString"}
	indexingPolicyModel["exclude_paths"] = []interface{}{"testString"}

	mssqlVolumeProtectionGroupObjectParamsModel := make(map[string]interface{})
	mssqlVolumeProtectionGroupObjectParamsModel["id"] = int(26)

	mssqlVolumeProtectionGroupParamsModel := make(map[string]interface{})
	mssqlVolumeProtectionGroupParamsModel["aag_backup_preference_type"] = "kPrimaryReplicaOnly"
	mssqlVolumeProtectionGroupParamsModel["advanced_settings"] = []interface{}{advancedSettingsModel}
	mssqlVolumeProtectionGroupParamsModel["backup_system_dbs"] = true
	mssqlVolumeProtectionGroupParamsModel["exclude_filters"] = []interface{}{filterModel}
	mssqlVolumeProtectionGroupParamsModel["full_backups_copy_only"] = true
	mssqlVolumeProtectionGroupParamsModel["log_backup_num_streams"] = int(38)
	mssqlVolumeProtectionGroupParamsModel["log_backup_with_clause"] = "testString"
	mssqlVolumeProtectionGroupParamsModel["pre_post_script"] = []interface{}{prePostScriptParamsModel}
	mssqlVolumeProtectionGroupParamsModel["use_aag_preferences_from_server"] = true
	mssqlVolumeProtectionGroupParamsModel["user_db_backup_preference_type"] = "kBackupAllDatabases"
	mssqlVolumeProtectionGroupParamsModel["additional_host_params"] = []interface{}{mssqlVolumeProtectionGroupHostParamsModel}
	mssqlVolumeProtectionGroupParamsModel["backup_db_volumes_only"] = true
	mssqlVolumeProtectionGroupParamsModel["incremental_backup_after_restart"] = true
	mssqlVolumeProtectionGroupParamsModel["indexing_policy"] = []interface{}{indexingPolicyModel}
	mssqlVolumeProtectionGroupParamsModel["objects"] = []interface{}{mssqlVolumeProtectionGroupObjectParamsModel}

	model := make(map[string]interface{})
	model["file_protection_type_params"] = []interface{}{mssqlFileProtectionGroupParamsModel}
	model["native_protection_type_params"] = []interface{}{mssqlNativeProtectionGroupParamsModel}
	model["protection_type"] = "kFile"
	model["volume_protection_type_params"] = []interface{}{mssqlVolumeProtectionGroupParamsModel}

	result, err := backuprecovery.ResourceIbmBaasProtectionGroupMapToMSSQLProtectionGroupParams(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasProtectionGroupMapToMSSQLFileProtectionGroupParams(t *testing.T) {
	checkResult := func(result *backuprecoveryv1.MSSQLFileProtectionGroupParams) {
		advancedSettingsModel := new(backuprecoveryv1.AdvancedSettings)
		advancedSettingsModel.ClonedDbBackupStatus = core.StringPtr("kError")
		advancedSettingsModel.DbBackupIfNotOnlineStatus = core.StringPtr("kError")
		advancedSettingsModel.MissingDbBackupStatus = core.StringPtr("kError")
		advancedSettingsModel.OfflineRestoringDbBackupStatus = core.StringPtr("kError")
		advancedSettingsModel.ReadOnlyDbBackupStatus = core.StringPtr("kError")
		advancedSettingsModel.ReportAllNonAutoprotectDbErrors = core.StringPtr("kError")

		filterModel := new(backuprecoveryv1.Filter)
		filterModel.FilterString = core.StringPtr("testString")
		filterModel.IsRegularExpression = core.BoolPtr(false)

		commonPreBackupScriptParamsModel := new(backuprecoveryv1.CommonPreBackupScriptParams)
		commonPreBackupScriptParamsModel.Path = core.StringPtr("testString")
		commonPreBackupScriptParamsModel.Params = core.StringPtr("testString")
		commonPreBackupScriptParamsModel.TimeoutSecs = core.Int64Ptr(int64(1))
		commonPreBackupScriptParamsModel.IsActive = core.BoolPtr(true)
		commonPreBackupScriptParamsModel.ContinueOnError = core.BoolPtr(true)

		commonPostBackupScriptParamsModel := new(backuprecoveryv1.CommonPostBackupScriptParams)
		commonPostBackupScriptParamsModel.Path = core.StringPtr("testString")
		commonPostBackupScriptParamsModel.Params = core.StringPtr("testString")
		commonPostBackupScriptParamsModel.TimeoutSecs = core.Int64Ptr(int64(1))
		commonPostBackupScriptParamsModel.IsActive = core.BoolPtr(true)

		prePostScriptParamsModel := new(backuprecoveryv1.PrePostScriptParams)
		prePostScriptParamsModel.PreScript = commonPreBackupScriptParamsModel
		prePostScriptParamsModel.PostScript = commonPostBackupScriptParamsModel

		mssqlFileProtectionGroupHostParamsModel := new(backuprecoveryv1.MSSQLFileProtectionGroupHostParams)
		mssqlFileProtectionGroupHostParamsModel.DisableSourceSideDeduplication = core.BoolPtr(true)
		mssqlFileProtectionGroupHostParamsModel.HostID = core.Int64Ptr(int64(26))

		mssqlFileProtectionGroupObjectParamsModel := new(backuprecoveryv1.MSSQLFileProtectionGroupObjectParams)
		mssqlFileProtectionGroupObjectParamsModel.ID = core.Int64Ptr(int64(26))

		model := new(backuprecoveryv1.MSSQLFileProtectionGroupParams)
		model.AagBackupPreferenceType = core.StringPtr("kPrimaryReplicaOnly")
		model.AdvancedSettings = advancedSettingsModel
		model.BackupSystemDbs = core.BoolPtr(true)
		model.ExcludeFilters = []backuprecoveryv1.Filter{*filterModel}
		model.FullBackupsCopyOnly = core.BoolPtr(true)
		model.LogBackupNumStreams = core.Int64Ptr(int64(38))
		model.LogBackupWithClause = core.StringPtr("testString")
		model.PrePostScript = prePostScriptParamsModel
		model.UseAagPreferencesFromServer = core.BoolPtr(true)
		model.UserDbBackupPreferenceType = core.StringPtr("kBackupAllDatabases")
		model.AdditionalHostParams = []backuprecoveryv1.MSSQLFileProtectionGroupHostParams{*mssqlFileProtectionGroupHostParamsModel}
		model.Objects = []backuprecoveryv1.MSSQLFileProtectionGroupObjectParams{*mssqlFileProtectionGroupObjectParamsModel}
		model.PerformSourceSideDeduplication = core.BoolPtr(true)

		assert.Equal(t, result, model)
	}

	advancedSettingsModel := make(map[string]interface{})
	advancedSettingsModel["cloned_db_backup_status"] = "kError"
	advancedSettingsModel["db_backup_if_not_online_status"] = "kError"
	advancedSettingsModel["missing_db_backup_status"] = "kError"
	advancedSettingsModel["offline_restoring_db_backup_status"] = "kError"
	advancedSettingsModel["read_only_db_backup_status"] = "kError"
	advancedSettingsModel["report_all_non_autoprotect_db_errors"] = "kError"

	filterModel := make(map[string]interface{})
	filterModel["filter_string"] = "testString"
	filterModel["is_regular_expression"] = false

	commonPreBackupScriptParamsModel := make(map[string]interface{})
	commonPreBackupScriptParamsModel["path"] = "testString"
	commonPreBackupScriptParamsModel["params"] = "testString"
	commonPreBackupScriptParamsModel["timeout_secs"] = int(1)
	commonPreBackupScriptParamsModel["is_active"] = true
	commonPreBackupScriptParamsModel["continue_on_error"] = true

	commonPostBackupScriptParamsModel := make(map[string]interface{})
	commonPostBackupScriptParamsModel["path"] = "testString"
	commonPostBackupScriptParamsModel["params"] = "testString"
	commonPostBackupScriptParamsModel["timeout_secs"] = int(1)
	commonPostBackupScriptParamsModel["is_active"] = true

	prePostScriptParamsModel := make(map[string]interface{})
	prePostScriptParamsModel["pre_script"] = []interface{}{commonPreBackupScriptParamsModel}
	prePostScriptParamsModel["post_script"] = []interface{}{commonPostBackupScriptParamsModel}

	mssqlFileProtectionGroupHostParamsModel := make(map[string]interface{})
	mssqlFileProtectionGroupHostParamsModel["disable_source_side_deduplication"] = true
	mssqlFileProtectionGroupHostParamsModel["host_id"] = int(26)

	mssqlFileProtectionGroupObjectParamsModel := make(map[string]interface{})
	mssqlFileProtectionGroupObjectParamsModel["id"] = int(26)

	model := make(map[string]interface{})
	model["aag_backup_preference_type"] = "kPrimaryReplicaOnly"
	model["advanced_settings"] = []interface{}{advancedSettingsModel}
	model["backup_system_dbs"] = true
	model["exclude_filters"] = []interface{}{filterModel}
	model["full_backups_copy_only"] = true
	model["log_backup_num_streams"] = int(38)
	model["log_backup_with_clause"] = "testString"
	model["pre_post_script"] = []interface{}{prePostScriptParamsModel}
	model["use_aag_preferences_from_server"] = true
	model["user_db_backup_preference_type"] = "kBackupAllDatabases"
	model["additional_host_params"] = []interface{}{mssqlFileProtectionGroupHostParamsModel}
	model["objects"] = []interface{}{mssqlFileProtectionGroupObjectParamsModel}
	model["perform_source_side_deduplication"] = true

	result, err := backuprecovery.ResourceIbmBaasProtectionGroupMapToMSSQLFileProtectionGroupParams(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasProtectionGroupMapToAdvancedSettings(t *testing.T) {
	checkResult := func(result *backuprecoveryv1.AdvancedSettings) {
		model := new(backuprecoveryv1.AdvancedSettings)
		model.ClonedDbBackupStatus = core.StringPtr("kError")
		model.DbBackupIfNotOnlineStatus = core.StringPtr("kError")
		model.MissingDbBackupStatus = core.StringPtr("kError")
		model.OfflineRestoringDbBackupStatus = core.StringPtr("kError")
		model.ReadOnlyDbBackupStatus = core.StringPtr("kError")
		model.ReportAllNonAutoprotectDbErrors = core.StringPtr("kError")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["cloned_db_backup_status"] = "kError"
	model["db_backup_if_not_online_status"] = "kError"
	model["missing_db_backup_status"] = "kError"
	model["offline_restoring_db_backup_status"] = "kError"
	model["read_only_db_backup_status"] = "kError"
	model["report_all_non_autoprotect_db_errors"] = "kError"

	result, err := backuprecovery.ResourceIbmBaasProtectionGroupMapToAdvancedSettings(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasProtectionGroupMapToFilter(t *testing.T) {
	checkResult := func(result *backuprecoveryv1.Filter) {
		model := new(backuprecoveryv1.Filter)
		model.FilterString = core.StringPtr("testString")
		model.IsRegularExpression = core.BoolPtr(false)

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["filter_string"] = "testString"
	model["is_regular_expression"] = false

	result, err := backuprecovery.ResourceIbmBaasProtectionGroupMapToFilter(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasProtectionGroupMapToMSSQLFileProtectionGroupHostParams(t *testing.T) {
	checkResult := func(result *backuprecoveryv1.MSSQLFileProtectionGroupHostParams) {
		model := new(backuprecoveryv1.MSSQLFileProtectionGroupHostParams)
		model.DisableSourceSideDeduplication = core.BoolPtr(true)
		model.HostID = core.Int64Ptr(int64(26))
		model.HostName = core.StringPtr("testString")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["disable_source_side_deduplication"] = true
	model["host_id"] = int(26)
	model["host_name"] = "testString"

	result, err := backuprecovery.ResourceIbmBaasProtectionGroupMapToMSSQLFileProtectionGroupHostParams(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasProtectionGroupMapToMSSQLFileProtectionGroupObjectParams(t *testing.T) {
	checkResult := func(result *backuprecoveryv1.MSSQLFileProtectionGroupObjectParams) {
		model := new(backuprecoveryv1.MSSQLFileProtectionGroupObjectParams)
		model.ID = core.Int64Ptr(int64(26))
		model.Name = core.StringPtr("testString")
		model.SourceType = core.StringPtr("testString")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["id"] = int(26)
	model["name"] = "testString"
	model["source_type"] = "testString"

	result, err := backuprecovery.ResourceIbmBaasProtectionGroupMapToMSSQLFileProtectionGroupObjectParams(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasProtectionGroupMapToMSSQLNativeProtectionGroupParams(t *testing.T) {
	checkResult := func(result *backuprecoveryv1.MSSQLNativeProtectionGroupParams) {
		advancedSettingsModel := new(backuprecoveryv1.AdvancedSettings)
		advancedSettingsModel.ClonedDbBackupStatus = core.StringPtr("kError")
		advancedSettingsModel.DbBackupIfNotOnlineStatus = core.StringPtr("kError")
		advancedSettingsModel.MissingDbBackupStatus = core.StringPtr("kError")
		advancedSettingsModel.OfflineRestoringDbBackupStatus = core.StringPtr("kError")
		advancedSettingsModel.ReadOnlyDbBackupStatus = core.StringPtr("kError")
		advancedSettingsModel.ReportAllNonAutoprotectDbErrors = core.StringPtr("kError")

		filterModel := new(backuprecoveryv1.Filter)
		filterModel.FilterString = core.StringPtr("testString")
		filterModel.IsRegularExpression = core.BoolPtr(false)

		commonPreBackupScriptParamsModel := new(backuprecoveryv1.CommonPreBackupScriptParams)
		commonPreBackupScriptParamsModel.Path = core.StringPtr("testString")
		commonPreBackupScriptParamsModel.Params = core.StringPtr("testString")
		commonPreBackupScriptParamsModel.TimeoutSecs = core.Int64Ptr(int64(1))
		commonPreBackupScriptParamsModel.IsActive = core.BoolPtr(true)
		commonPreBackupScriptParamsModel.ContinueOnError = core.BoolPtr(true)

		commonPostBackupScriptParamsModel := new(backuprecoveryv1.CommonPostBackupScriptParams)
		commonPostBackupScriptParamsModel.Path = core.StringPtr("testString")
		commonPostBackupScriptParamsModel.Params = core.StringPtr("testString")
		commonPostBackupScriptParamsModel.TimeoutSecs = core.Int64Ptr(int64(1))
		commonPostBackupScriptParamsModel.IsActive = core.BoolPtr(true)

		prePostScriptParamsModel := new(backuprecoveryv1.PrePostScriptParams)
		prePostScriptParamsModel.PreScript = commonPreBackupScriptParamsModel
		prePostScriptParamsModel.PostScript = commonPostBackupScriptParamsModel

		mssqlNativeProtectionGroupObjectParamsModel := new(backuprecoveryv1.MSSQLNativeProtectionGroupObjectParams)
		mssqlNativeProtectionGroupObjectParamsModel.ID = core.Int64Ptr(int64(26))

		model := new(backuprecoveryv1.MSSQLNativeProtectionGroupParams)
		model.AagBackupPreferenceType = core.StringPtr("kPrimaryReplicaOnly")
		model.AdvancedSettings = advancedSettingsModel
		model.BackupSystemDbs = core.BoolPtr(true)
		model.ExcludeFilters = []backuprecoveryv1.Filter{*filterModel}
		model.FullBackupsCopyOnly = core.BoolPtr(true)
		model.LogBackupNumStreams = core.Int64Ptr(int64(38))
		model.LogBackupWithClause = core.StringPtr("testString")
		model.PrePostScript = prePostScriptParamsModel
		model.UseAagPreferencesFromServer = core.BoolPtr(true)
		model.UserDbBackupPreferenceType = core.StringPtr("kBackupAllDatabases")
		model.NumStreams = core.Int64Ptr(int64(38))
		model.Objects = []backuprecoveryv1.MSSQLNativeProtectionGroupObjectParams{*mssqlNativeProtectionGroupObjectParamsModel}
		model.WithClause = core.StringPtr("testString")

		assert.Equal(t, result, model)
	}

	advancedSettingsModel := make(map[string]interface{})
	advancedSettingsModel["cloned_db_backup_status"] = "kError"
	advancedSettingsModel["db_backup_if_not_online_status"] = "kError"
	advancedSettingsModel["missing_db_backup_status"] = "kError"
	advancedSettingsModel["offline_restoring_db_backup_status"] = "kError"
	advancedSettingsModel["read_only_db_backup_status"] = "kError"
	advancedSettingsModel["report_all_non_autoprotect_db_errors"] = "kError"

	filterModel := make(map[string]interface{})
	filterModel["filter_string"] = "testString"
	filterModel["is_regular_expression"] = false

	commonPreBackupScriptParamsModel := make(map[string]interface{})
	commonPreBackupScriptParamsModel["path"] = "testString"
	commonPreBackupScriptParamsModel["params"] = "testString"
	commonPreBackupScriptParamsModel["timeout_secs"] = int(1)
	commonPreBackupScriptParamsModel["is_active"] = true
	commonPreBackupScriptParamsModel["continue_on_error"] = true

	commonPostBackupScriptParamsModel := make(map[string]interface{})
	commonPostBackupScriptParamsModel["path"] = "testString"
	commonPostBackupScriptParamsModel["params"] = "testString"
	commonPostBackupScriptParamsModel["timeout_secs"] = int(1)
	commonPostBackupScriptParamsModel["is_active"] = true

	prePostScriptParamsModel := make(map[string]interface{})
	prePostScriptParamsModel["pre_script"] = []interface{}{commonPreBackupScriptParamsModel}
	prePostScriptParamsModel["post_script"] = []interface{}{commonPostBackupScriptParamsModel}

	mssqlNativeProtectionGroupObjectParamsModel := make(map[string]interface{})
	mssqlNativeProtectionGroupObjectParamsModel["id"] = int(26)

	model := make(map[string]interface{})
	model["aag_backup_preference_type"] = "kPrimaryReplicaOnly"
	model["advanced_settings"] = []interface{}{advancedSettingsModel}
	model["backup_system_dbs"] = true
	model["exclude_filters"] = []interface{}{filterModel}
	model["full_backups_copy_only"] = true
	model["log_backup_num_streams"] = int(38)
	model["log_backup_with_clause"] = "testString"
	model["pre_post_script"] = []interface{}{prePostScriptParamsModel}
	model["use_aag_preferences_from_server"] = true
	model["user_db_backup_preference_type"] = "kBackupAllDatabases"
	model["num_streams"] = int(38)
	model["objects"] = []interface{}{mssqlNativeProtectionGroupObjectParamsModel}
	model["with_clause"] = "testString"

	result, err := backuprecovery.ResourceIbmBaasProtectionGroupMapToMSSQLNativeProtectionGroupParams(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasProtectionGroupMapToMSSQLNativeProtectionGroupObjectParams(t *testing.T) {
	checkResult := func(result *backuprecoveryv1.MSSQLNativeProtectionGroupObjectParams) {
		model := new(backuprecoveryv1.MSSQLNativeProtectionGroupObjectParams)
		model.ID = core.Int64Ptr(int64(26))
		model.Name = core.StringPtr("testString")
		model.SourceType = core.StringPtr("testString")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["id"] = int(26)
	model["name"] = "testString"
	model["source_type"] = "testString"

	result, err := backuprecovery.ResourceIbmBaasProtectionGroupMapToMSSQLNativeProtectionGroupObjectParams(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasProtectionGroupMapToMSSQLVolumeProtectionGroupParams(t *testing.T) {
	checkResult := func(result *backuprecoveryv1.MSSQLVolumeProtectionGroupParams) {
		advancedSettingsModel := new(backuprecoveryv1.AdvancedSettings)
		advancedSettingsModel.ClonedDbBackupStatus = core.StringPtr("kError")
		advancedSettingsModel.DbBackupIfNotOnlineStatus = core.StringPtr("kError")
		advancedSettingsModel.MissingDbBackupStatus = core.StringPtr("kError")
		advancedSettingsModel.OfflineRestoringDbBackupStatus = core.StringPtr("kError")
		advancedSettingsModel.ReadOnlyDbBackupStatus = core.StringPtr("kError")
		advancedSettingsModel.ReportAllNonAutoprotectDbErrors = core.StringPtr("kError")

		filterModel := new(backuprecoveryv1.Filter)
		filterModel.FilterString = core.StringPtr("testString")
		filterModel.IsRegularExpression = core.BoolPtr(false)

		commonPreBackupScriptParamsModel := new(backuprecoveryv1.CommonPreBackupScriptParams)
		commonPreBackupScriptParamsModel.Path = core.StringPtr("testString")
		commonPreBackupScriptParamsModel.Params = core.StringPtr("testString")
		commonPreBackupScriptParamsModel.TimeoutSecs = core.Int64Ptr(int64(1))
		commonPreBackupScriptParamsModel.IsActive = core.BoolPtr(true)
		commonPreBackupScriptParamsModel.ContinueOnError = core.BoolPtr(true)

		commonPostBackupScriptParamsModel := new(backuprecoveryv1.CommonPostBackupScriptParams)
		commonPostBackupScriptParamsModel.Path = core.StringPtr("testString")
		commonPostBackupScriptParamsModel.Params = core.StringPtr("testString")
		commonPostBackupScriptParamsModel.TimeoutSecs = core.Int64Ptr(int64(1))
		commonPostBackupScriptParamsModel.IsActive = core.BoolPtr(true)

		prePostScriptParamsModel := new(backuprecoveryv1.PrePostScriptParams)
		prePostScriptParamsModel.PreScript = commonPreBackupScriptParamsModel
		prePostScriptParamsModel.PostScript = commonPostBackupScriptParamsModel

		mssqlVolumeProtectionGroupHostParamsModel := new(backuprecoveryv1.MSSQLVolumeProtectionGroupHostParams)
		mssqlVolumeProtectionGroupHostParamsModel.EnableSystemBackup = core.BoolPtr(true)
		mssqlVolumeProtectionGroupHostParamsModel.HostID = core.Int64Ptr(int64(26))
		mssqlVolumeProtectionGroupHostParamsModel.VolumeGuids = []string{"testString"}

		indexingPolicyModel := new(backuprecoveryv1.IndexingPolicy)
		indexingPolicyModel.EnableIndexing = core.BoolPtr(true)
		indexingPolicyModel.IncludePaths = []string{"testString"}
		indexingPolicyModel.ExcludePaths = []string{"testString"}

		mssqlVolumeProtectionGroupObjectParamsModel := new(backuprecoveryv1.MSSQLVolumeProtectionGroupObjectParams)
		mssqlVolumeProtectionGroupObjectParamsModel.ID = core.Int64Ptr(int64(26))

		model := new(backuprecoveryv1.MSSQLVolumeProtectionGroupParams)
		model.AagBackupPreferenceType = core.StringPtr("kPrimaryReplicaOnly")
		model.AdvancedSettings = advancedSettingsModel
		model.BackupSystemDbs = core.BoolPtr(true)
		model.ExcludeFilters = []backuprecoveryv1.Filter{*filterModel}
		model.FullBackupsCopyOnly = core.BoolPtr(true)
		model.LogBackupNumStreams = core.Int64Ptr(int64(38))
		model.LogBackupWithClause = core.StringPtr("testString")
		model.PrePostScript = prePostScriptParamsModel
		model.UseAagPreferencesFromServer = core.BoolPtr(true)
		model.UserDbBackupPreferenceType = core.StringPtr("kBackupAllDatabases")
		model.AdditionalHostParams = []backuprecoveryv1.MSSQLVolumeProtectionGroupHostParams{*mssqlVolumeProtectionGroupHostParamsModel}
		model.BackupDbVolumesOnly = core.BoolPtr(true)
		model.IncrementalBackupAfterRestart = core.BoolPtr(true)
		model.IndexingPolicy = indexingPolicyModel
		model.Objects = []backuprecoveryv1.MSSQLVolumeProtectionGroupObjectParams{*mssqlVolumeProtectionGroupObjectParamsModel}

		assert.Equal(t, result, model)
	}

	advancedSettingsModel := make(map[string]interface{})
	advancedSettingsModel["cloned_db_backup_status"] = "kError"
	advancedSettingsModel["db_backup_if_not_online_status"] = "kError"
	advancedSettingsModel["missing_db_backup_status"] = "kError"
	advancedSettingsModel["offline_restoring_db_backup_status"] = "kError"
	advancedSettingsModel["read_only_db_backup_status"] = "kError"
	advancedSettingsModel["report_all_non_autoprotect_db_errors"] = "kError"

	filterModel := make(map[string]interface{})
	filterModel["filter_string"] = "testString"
	filterModel["is_regular_expression"] = false

	commonPreBackupScriptParamsModel := make(map[string]interface{})
	commonPreBackupScriptParamsModel["path"] = "testString"
	commonPreBackupScriptParamsModel["params"] = "testString"
	commonPreBackupScriptParamsModel["timeout_secs"] = int(1)
	commonPreBackupScriptParamsModel["is_active"] = true
	commonPreBackupScriptParamsModel["continue_on_error"] = true

	commonPostBackupScriptParamsModel := make(map[string]interface{})
	commonPostBackupScriptParamsModel["path"] = "testString"
	commonPostBackupScriptParamsModel["params"] = "testString"
	commonPostBackupScriptParamsModel["timeout_secs"] = int(1)
	commonPostBackupScriptParamsModel["is_active"] = true

	prePostScriptParamsModel := make(map[string]interface{})
	prePostScriptParamsModel["pre_script"] = []interface{}{commonPreBackupScriptParamsModel}
	prePostScriptParamsModel["post_script"] = []interface{}{commonPostBackupScriptParamsModel}

	mssqlVolumeProtectionGroupHostParamsModel := make(map[string]interface{})
	mssqlVolumeProtectionGroupHostParamsModel["enable_system_backup"] = true
	mssqlVolumeProtectionGroupHostParamsModel["host_id"] = int(26)
	mssqlVolumeProtectionGroupHostParamsModel["volume_guids"] = []interface{}{"testString"}

	indexingPolicyModel := make(map[string]interface{})
	indexingPolicyModel["enable_indexing"] = true
	indexingPolicyModel["include_paths"] = []interface{}{"testString"}
	indexingPolicyModel["exclude_paths"] = []interface{}{"testString"}

	mssqlVolumeProtectionGroupObjectParamsModel := make(map[string]interface{})
	mssqlVolumeProtectionGroupObjectParamsModel["id"] = int(26)

	model := make(map[string]interface{})
	model["aag_backup_preference_type"] = "kPrimaryReplicaOnly"
	model["advanced_settings"] = []interface{}{advancedSettingsModel}
	model["backup_system_dbs"] = true
	model["exclude_filters"] = []interface{}{filterModel}
	model["full_backups_copy_only"] = true
	model["log_backup_num_streams"] = int(38)
	model["log_backup_with_clause"] = "testString"
	model["pre_post_script"] = []interface{}{prePostScriptParamsModel}
	model["use_aag_preferences_from_server"] = true
	model["user_db_backup_preference_type"] = "kBackupAllDatabases"
	model["additional_host_params"] = []interface{}{mssqlVolumeProtectionGroupHostParamsModel}
	model["backup_db_volumes_only"] = true
	model["incremental_backup_after_restart"] = true
	model["indexing_policy"] = []interface{}{indexingPolicyModel}
	model["objects"] = []interface{}{mssqlVolumeProtectionGroupObjectParamsModel}

	result, err := backuprecovery.ResourceIbmBaasProtectionGroupMapToMSSQLVolumeProtectionGroupParams(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasProtectionGroupMapToMSSQLVolumeProtectionGroupHostParams(t *testing.T) {
	checkResult := func(result *backuprecoveryv1.MSSQLVolumeProtectionGroupHostParams) {
		model := new(backuprecoveryv1.MSSQLVolumeProtectionGroupHostParams)
		model.EnableSystemBackup = core.BoolPtr(true)
		model.HostID = core.Int64Ptr(int64(26))
		model.HostName = core.StringPtr("testString")
		model.VolumeGuids = []string{"testString"}

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["enable_system_backup"] = true
	model["host_id"] = int(26)
	model["host_name"] = "testString"
	model["volume_guids"] = []interface{}{"testString"}

	result, err := backuprecovery.ResourceIbmBaasProtectionGroupMapToMSSQLVolumeProtectionGroupHostParams(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestResourceIbmBaasProtectionGroupMapToMSSQLVolumeProtectionGroupObjectParams(t *testing.T) {
	checkResult := func(result *backuprecoveryv1.MSSQLVolumeProtectionGroupObjectParams) {
		model := new(backuprecoveryv1.MSSQLVolumeProtectionGroupObjectParams)
		model.ID = core.Int64Ptr(int64(26))
		model.Name = core.StringPtr("testString")
		model.SourceType = core.StringPtr("testString")

		assert.Equal(t, result, model)
	}

	model := make(map[string]interface{})
	model["id"] = int(26)
	model["name"] = "testString"
	model["source_type"] = "testString"

	result, err := backuprecovery.ResourceIbmBaasProtectionGroupMapToMSSQLVolumeProtectionGroupObjectParams(model)
	assert.Nil(t, err)
	checkResult(result)
}
