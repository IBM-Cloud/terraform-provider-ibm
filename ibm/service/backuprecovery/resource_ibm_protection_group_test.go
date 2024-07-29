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

func TestAccIbmProtectionGroupBasic(t *testing.T) {
	var conf backuprecoveryv1.ProtectionGroupResponse
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	policyID := fmt.Sprintf("tf_policy_id_%d", acctest.RandIntRange(10, 100))
	environment := "kPhysical"
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	policyIDUpdate := fmt.Sprintf("tf_policy_id_%d", acctest.RandIntRange(10, 100))
	environmentUpdate := "kOracle"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmProtectionGroupDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmProtectionGroupConfigBasic(name, policyID, environment),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmProtectionGroupExists("ibm_protection_group.protection_group_instance", conf),
					resource.TestCheckResourceAttr("ibm_protection_group.protection_group_instance", "name", name),
					resource.TestCheckResourceAttr("ibm_protection_group.protection_group_instance", "policy_id", policyID),
					resource.TestCheckResourceAttr("ibm_protection_group.protection_group_instance", "environment", environment),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmProtectionGroupConfigBasic(nameUpdate, policyIDUpdate, environmentUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_protection_group.protection_group_instance", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_protection_group.protection_group_instance", "policy_id", policyIDUpdate),
					resource.TestCheckResourceAttr("ibm_protection_group.protection_group_instance", "environment", environmentUpdate),
				),
			},
		},
	})
}

func TestAccIbmProtectionGroupAllArgs(t *testing.T) {
	var conf backuprecoveryv1.ProtectionGroupResponse
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	policyID := fmt.Sprintf("tf_policy_id_%d", acctest.RandIntRange(10, 100))
	priority := "kLow"
	storageDomainID := fmt.Sprintf("%d", acctest.RandIntRange(10, 100))
	description := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	endTimeUsecs := fmt.Sprintf("%d", acctest.RandIntRange(10, 100))
	lastModifiedTimestampUsecs := fmt.Sprintf("%d", acctest.RandIntRange(10, 100))
	qosPolicy := "kBackupHDD"
	abortInBlackouts := "true"
	pauseInBlackouts := "true"
	isPaused := "false"
	environment := "kPhysical"
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	policyIDUpdate := fmt.Sprintf("tf_policy_id_%d", acctest.RandIntRange(10, 100))
	priorityUpdate := "kHigh"
	storageDomainIDUpdate := fmt.Sprintf("%d", acctest.RandIntRange(10, 100))
	descriptionUpdate := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	endTimeUsecsUpdate := fmt.Sprintf("%d", acctest.RandIntRange(10, 100))
	lastModifiedTimestampUsecsUpdate := fmt.Sprintf("%d", acctest.RandIntRange(10, 100))
	qosPolicyUpdate := "kBackupAll"
	abortInBlackoutsUpdate := "false"
	pauseInBlackoutsUpdate := "false"
	isPausedUpdate := "true"
	environmentUpdate := "kOracle"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmProtectionGroupDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmProtectionGroupConfig(name, policyID, priority, storageDomainID, description, endTimeUsecs, lastModifiedTimestampUsecs, qosPolicy, abortInBlackouts, pauseInBlackouts, isPaused, environment),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmProtectionGroupExists("ibm_protection_group.protection_group_instance", conf),
					resource.TestCheckResourceAttr("ibm_protection_group.protection_group_instance", "name", name),
					resource.TestCheckResourceAttr("ibm_protection_group.protection_group_instance", "policy_id", policyID),
					resource.TestCheckResourceAttr("ibm_protection_group.protection_group_instance", "priority", priority),
					resource.TestCheckResourceAttr("ibm_protection_group.protection_group_instance", "storage_domain_id", storageDomainID),
					resource.TestCheckResourceAttr("ibm_protection_group.protection_group_instance", "description", description),
					resource.TestCheckResourceAttr("ibm_protection_group.protection_group_instance", "end_time_usecs", endTimeUsecs),
					resource.TestCheckResourceAttr("ibm_protection_group.protection_group_instance", "last_modified_timestamp_usecs", lastModifiedTimestampUsecs),
					resource.TestCheckResourceAttr("ibm_protection_group.protection_group_instance", "qos_policy", qosPolicy),
					resource.TestCheckResourceAttr("ibm_protection_group.protection_group_instance", "abort_in_blackouts", abortInBlackouts),
					resource.TestCheckResourceAttr("ibm_protection_group.protection_group_instance", "pause_in_blackouts", pauseInBlackouts),
					resource.TestCheckResourceAttr("ibm_protection_group.protection_group_instance", "is_paused", isPaused),
					resource.TestCheckResourceAttr("ibm_protection_group.protection_group_instance", "environment", environment),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmProtectionGroupConfig(nameUpdate, policyIDUpdate, priorityUpdate, storageDomainIDUpdate, descriptionUpdate, endTimeUsecsUpdate, lastModifiedTimestampUsecsUpdate, qosPolicyUpdate, abortInBlackoutsUpdate, pauseInBlackoutsUpdate, isPausedUpdate, environmentUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_protection_group.protection_group_instance", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_protection_group.protection_group_instance", "policy_id", policyIDUpdate),
					resource.TestCheckResourceAttr("ibm_protection_group.protection_group_instance", "priority", priorityUpdate),
					resource.TestCheckResourceAttr("ibm_protection_group.protection_group_instance", "storage_domain_id", storageDomainIDUpdate),
					resource.TestCheckResourceAttr("ibm_protection_group.protection_group_instance", "description", descriptionUpdate),
					resource.TestCheckResourceAttr("ibm_protection_group.protection_group_instance", "end_time_usecs", endTimeUsecsUpdate),
					resource.TestCheckResourceAttr("ibm_protection_group.protection_group_instance", "last_modified_timestamp_usecs", lastModifiedTimestampUsecsUpdate),
					resource.TestCheckResourceAttr("ibm_protection_group.protection_group_instance", "qos_policy", qosPolicyUpdate),
					resource.TestCheckResourceAttr("ibm_protection_group.protection_group_instance", "abort_in_blackouts", abortInBlackoutsUpdate),
					resource.TestCheckResourceAttr("ibm_protection_group.protection_group_instance", "pause_in_blackouts", pauseInBlackoutsUpdate),
					resource.TestCheckResourceAttr("ibm_protection_group.protection_group_instance", "is_paused", isPausedUpdate),
					resource.TestCheckResourceAttr("ibm_protection_group.protection_group_instance", "environment", environmentUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_protection_group.protection_group",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmProtectionGroupConfigBasic(name string, policyID string, environment string) string {
	return fmt.Sprintf(`
		resource "ibm_protection_group" "protection_group_instance" {
			name = "%s"
			policy_id = "%s"
			environment = "%s"
		}
	`, name, policyID, environment)
}

func testAccCheckIbmProtectionGroupConfig(name string, policyID string, priority string, storageDomainID string, description string, endTimeUsecs string, lastModifiedTimestampUsecs string, qosPolicy string, abortInBlackouts string, pauseInBlackouts string, isPaused string, environment string) string {
	return fmt.Sprintf(`

		resource "ibm_protection_group" "protection_group_instance" {
			name = "%s"
			policy_id = "%s"
			priority = "%s"
			storage_domain_id = %s
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
			environment = "%s"
			is_paused = %s
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
					objects {
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
			oracle_params {
				objects {
					source_id = 1
					source_name = "source_name"
					db_params {
						database_id = 1
						database_name = "database_name"
						db_channels {
							archive_log_retention_days = 1
							archive_log_retention_hours = 1
							credentials {
								username = "username"
								password = "password"
							}
							database_unique_name = "database_unique_name"
							database_uuid = "database_uuid"
							default_channel_count = 1
							database_node_list {
								host_id = "host_id"
								channel_count = 1
								port = 1
								sbt_host_params {
									sbt_library_path = "sbt_library_path"
									view_fs_path = "view_fs_path"
									vip_list = [ "vipList" ]
									vlan_info_list {
										ip_list = [ "ipList" ]
										gateway = "gateway"
										id = 1
										subnet_ip = "subnet_ip"
									}
								}
							}
							max_host_count = 1
							enable_dg_primary_backup = true
							rman_backup_type = "kImageCopy"
						}
					}
				}
				persist_mountpoints = true
				vlan_params {
					vlan_id = 1
					disable_vlan = true
					interface_name = "interface_name"
				}
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
				log_auto_kill_timeout_secs = 1
				incr_auto_kill_timeout_secs = 1
				full_auto_kill_timeout_secs = 1
			}
		}
	`, name, policyID, priority, storageDomainID, description, endTimeUsecs, lastModifiedTimestampUsecs, qosPolicy, abortInBlackouts, pauseInBlackouts, isPaused, environment)
}

func testAccCheckIbmProtectionGroupExists(n string, obj backuprecoveryv1.ProtectionGroupResponse) resource.TestCheckFunc {

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

func testAccCheckIbmProtectionGroupDestroy(s *terraform.State) error {
	backupRecoveryClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).BackupRecoveryV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_protection_group" {
			continue
		}

		getProtectionGroupByIdOptions := &backuprecoveryv1.GetProtectionGroupByIdOptions{}

		getProtectionGroupByIdOptions.SetID(rs.Primary.ID)

		// Try to find the key
		_, response, err := backupRecoveryClient.GetProtectionGroupByID(getProtectionGroupByIdOptions)

		if err == nil {
			return fmt.Errorf("protection_group still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for protection_group (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
