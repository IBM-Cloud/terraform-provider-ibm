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

func TestAccIbmProtectionPolicyBasic(t *testing.T) {
	var conf backuprecoveryv1.ProtectionPolicyResponse
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmProtectionPolicyDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmProtectionPolicyConfigBasic(name),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmProtectionPolicyExists("ibm_protection_policy.protection_policy_instance", conf),
					resource.TestCheckResourceAttr("ibm_protection_policy.protection_policy_instance", "name", name),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmProtectionPolicyConfigBasic(nameUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_protection_policy.protection_policy_instance", "name", nameUpdate),
				),
			},
		},
	})
}

func TestAccIbmProtectionPolicyAllArgs(t *testing.T) {
	var conf backuprecoveryv1.ProtectionPolicyResponse
	name := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	description := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	dataLock := "Compliance"
	version := fmt.Sprintf("%d", acctest.RandIntRange(10, 100))
	isCbsEnabled := "true"
	lastModificationTimeUsecs := fmt.Sprintf("%d", acctest.RandIntRange(10, 100))
	templateID := fmt.Sprintf("tf_template_id_%d", acctest.RandIntRange(10, 100))
	nameUpdate := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	descriptionUpdate := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	dataLockUpdate := "Administrative"
	versionUpdate := fmt.Sprintf("%d", acctest.RandIntRange(10, 100))
	isCbsEnabledUpdate := "false"
	lastModificationTimeUsecsUpdate := fmt.Sprintf("%d", acctest.RandIntRange(10, 100))
	templateIDUpdate := fmt.Sprintf("tf_template_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { acc.TestAccPreCheck(t) },
		Providers:    acc.TestAccProviders,
		CheckDestroy: testAccCheckIbmProtectionPolicyDestroy,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmProtectionPolicyConfig(name, description, dataLock, version, isCbsEnabled, lastModificationTimeUsecs, templateID),
				Check: resource.ComposeAggregateTestCheckFunc(
					testAccCheckIbmProtectionPolicyExists("ibm_protection_policy.protection_policy_instance", conf),
					resource.TestCheckResourceAttr("ibm_protection_policy.protection_policy_instance", "name", name),
					resource.TestCheckResourceAttr("ibm_protection_policy.protection_policy_instance", "description", description),
					resource.TestCheckResourceAttr("ibm_protection_policy.protection_policy_instance", "data_lock", dataLock),
					resource.TestCheckResourceAttr("ibm_protection_policy.protection_policy_instance", "version", version),
					resource.TestCheckResourceAttr("ibm_protection_policy.protection_policy_instance", "is_cbs_enabled", isCbsEnabled),
					resource.TestCheckResourceAttr("ibm_protection_policy.protection_policy_instance", "last_modification_time_usecs", lastModificationTimeUsecs),
					resource.TestCheckResourceAttr("ibm_protection_policy.protection_policy_instance", "template_id", templateID),
				),
			},
			resource.TestStep{
				Config: testAccCheckIbmProtectionPolicyConfig(nameUpdate, descriptionUpdate, dataLockUpdate, versionUpdate, isCbsEnabledUpdate, lastModificationTimeUsecsUpdate, templateIDUpdate),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("ibm_protection_policy.protection_policy_instance", "name", nameUpdate),
					resource.TestCheckResourceAttr("ibm_protection_policy.protection_policy_instance", "description", descriptionUpdate),
					resource.TestCheckResourceAttr("ibm_protection_policy.protection_policy_instance", "data_lock", dataLockUpdate),
					resource.TestCheckResourceAttr("ibm_protection_policy.protection_policy_instance", "version", versionUpdate),
					resource.TestCheckResourceAttr("ibm_protection_policy.protection_policy_instance", "is_cbs_enabled", isCbsEnabledUpdate),
					resource.TestCheckResourceAttr("ibm_protection_policy.protection_policy_instance", "last_modification_time_usecs", lastModificationTimeUsecsUpdate),
					resource.TestCheckResourceAttr("ibm_protection_policy.protection_policy_instance", "template_id", templateIDUpdate),
				),
			},
			resource.TestStep{
				ResourceName:      "ibm_protection_policy.protection_policy",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckIbmProtectionPolicyConfigBasic(name string) string {
	return fmt.Sprintf(`
		resource "ibm_protection_policy" "protection_policy_instance" {
			name = "%s"
			backup_policy {
				regular {
					incremental {
						schedule {
							unit = "Minutes"
							minute_schedule {
								frequency = 1
							}
							hour_schedule {
								frequency = 1
							}
							day_schedule {
								frequency = 1
							}
							week_schedule {
								day_of_week = [ "Sunday" ]
							}
							month_schedule {
								day_of_week = [ "Sunday" ]
								week_of_month = "First"
								day_of_month = 1
							}
							year_schedule {
								day_of_year = "First"
							}
						}
					}
					full {
						schedule {
							unit = "Days"
							day_schedule {
								frequency = 1
							}
							week_schedule {
								day_of_week = [ "Sunday" ]
							}
							month_schedule {
								day_of_week = [ "Sunday" ]
								week_of_month = "First"
								day_of_month = 1
							}
							year_schedule {
								day_of_year = "First"
							}
						}
					}
					full_backups {
						schedule {
							unit = "Days"
							day_schedule {
								frequency = 1
							}
							week_schedule {
								day_of_week = [ "Sunday" ]
							}
							month_schedule {
								day_of_week = [ "Sunday" ]
								week_of_month = "First"
								day_of_month = 1
							}
							year_schedule {
								day_of_year = "First"
							}
						}
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
					primary_backup_target {
						target_type = "Local"
						archival_target_settings {
							target_id = 1
							target_name = "target_name"
							tier_settings {
								cloud_platform = "Oracle"
								oracle_tiering {
									tiers {
										move_after_unit = "Days"
										move_after = 1
										tier_type = "kOracleTierStandard"
									}
								}
							}
						}
					}
				}
				log {
					schedule {
						unit = "Minutes"
						minute_schedule {
							frequency = 1
						}
						hour_schedule {
							frequency = 1
						}
					}
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
				bmr {
					schedule {
						unit = "Days"
						day_schedule {
							frequency = 1
						}
						week_schedule {
							day_of_week = [ "Sunday" ]
						}
						month_schedule {
							day_of_week = [ "Sunday" ]
							week_of_month = "First"
							day_of_month = 1
						}
						year_schedule {
							day_of_year = "First"
						}
					}
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
				cdp {
					retention {
						unit = "Minutes"
						duration = 1
						data_lock_config {
							mode = "Compliance"
							unit = "Days"
							duration = 1
							enable_worm_on_external_target = true
						}
					}
				}
				storage_array_snapshot {
					schedule {
						unit = "Minutes"
						minute_schedule {
							frequency = 1
						}
						hour_schedule {
							frequency = 1
						}
						day_schedule {
							frequency = 1
						}
						week_schedule {
							day_of_week = [ "Sunday" ]
						}
						month_schedule {
							day_of_week = [ "Sunday" ]
							week_of_month = "First"
							day_of_month = 1
						}
						year_schedule {
							day_of_year = "First"
						}
					}
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
				run_timeouts {
					timeout_mins = 1
					backup_type = "kRegular"
				}
			}
		}
	`, name)
}

func testAccCheckIbmProtectionPolicyConfig(name string, description string, dataLock string, version string, isCbsEnabled string, lastModificationTimeUsecs string, templateID string) string {
	return fmt.Sprintf(`

		resource "ibm_protection_policy" "protection_policy_instance" {
			name = "%s"
			backup_policy {
				regular {
					incremental {
						schedule {
							unit = "Minutes"
							minute_schedule {
								frequency = 1
							}
							hour_schedule {
								frequency = 1
							}
							day_schedule {
								frequency = 1
							}
							week_schedule {
								day_of_week = [ "Sunday" ]
							}
							month_schedule {
								day_of_week = [ "Sunday" ]
								week_of_month = "First"
								day_of_month = 1
							}
							year_schedule {
								day_of_year = "First"
							}
						}
					}
					full {
						schedule {
							unit = "Days"
							day_schedule {
								frequency = 1
							}
							week_schedule {
								day_of_week = [ "Sunday" ]
							}
							month_schedule {
								day_of_week = [ "Sunday" ]
								week_of_month = "First"
								day_of_month = 1
							}
							year_schedule {
								day_of_year = "First"
							}
						}
					}
					full_backups {
						schedule {
							unit = "Days"
							day_schedule {
								frequency = 1
							}
							week_schedule {
								day_of_week = [ "Sunday" ]
							}
							month_schedule {
								day_of_week = [ "Sunday" ]
								week_of_month = "First"
								day_of_month = 1
							}
							year_schedule {
								day_of_year = "First"
							}
						}
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
					primary_backup_target {
						target_type = "Local"
						archival_target_settings {
							target_id = 1
							target_name = "target_name"
							tier_settings {
								cloud_platform = "Oracle"
								oracle_tiering {
									tiers {
										move_after_unit = "Days"
										move_after = 1
										tier_type = "kOracleTierStandard"
									}
								}
							}
						}
					}
				}
				log {
					schedule {
						unit = "Minutes"
						minute_schedule {
							frequency = 1
						}
						hour_schedule {
							frequency = 1
						}
					}
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
				bmr {
					schedule {
						unit = "Days"
						day_schedule {
							frequency = 1
						}
						week_schedule {
							day_of_week = [ "Sunday" ]
						}
						month_schedule {
							day_of_week = [ "Sunday" ]
							week_of_month = "First"
							day_of_month = 1
						}
						year_schedule {
							day_of_year = "First"
						}
					}
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
				cdp {
					retention {
						unit = "Minutes"
						duration = 1
						data_lock_config {
							mode = "Compliance"
							unit = "Days"
							duration = 1
							enable_worm_on_external_target = true
						}
					}
				}
				storage_array_snapshot {
					schedule {
						unit = "Minutes"
						minute_schedule {
							frequency = 1
						}
						hour_schedule {
							frequency = 1
						}
						day_schedule {
							frequency = 1
						}
						week_schedule {
							day_of_week = [ "Sunday" ]
						}
						month_schedule {
							day_of_week = [ "Sunday" ]
							week_of_month = "First"
							day_of_month = 1
						}
						year_schedule {
							day_of_year = "First"
						}
					}
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
				run_timeouts {
					timeout_mins = 1
					backup_type = "kRegular"
				}
			}
			description = "%s"
			blackout_window {
				day = "Sunday"
				start_time {
					hour = 0
					minute = 0
					time_zone = "time_zone"
				}
				end_time {
					hour = 0
					minute = 0
					time_zone = "time_zone"
				}
				config_id = "config_id"
			}
			extended_retention {
				schedule {
					unit = "Runs"
					frequency = 1
				}
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
				run_type = "Regular"
				config_id = "config_id"
			}
			remote_target_policy {
				replication_targets {
					schedule {
						unit = "Runs"
						frequency = 1
					}
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
					copy_on_run_success = true
					config_id = "config_id"
					backup_run_type = "Regular"
					run_timeouts {
						timeout_mins = 1
						backup_type = "kRegular"
					}
					log_retention {
						unit = "Days"
						duration = 1
						data_lock_config {
							mode = "Compliance"
							unit = "Days"
							duration = 1
							enable_worm_on_external_target = true
						}
					}
					target_type = "RemoteCluster"
					remote_target_config {
						cluster_id = 1
						cluster_name = "cluster_name"
					}
				}
				archival_targets {
					schedule {
						unit = "Runs"
						frequency = 1
					}
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
					copy_on_run_success = true
					config_id = "config_id"
					backup_run_type = "Regular"
					run_timeouts {
						timeout_mins = 1
						backup_type = "kRegular"
					}
					log_retention {
						unit = "Days"
						duration = 1
						data_lock_config {
							mode = "Compliance"
							unit = "Days"
							duration = 1
							enable_worm_on_external_target = true
						}
					}
					target_id = 1
					target_name = "target_name"
					target_type = "Tape"
					tier_settings {
						cloud_platform = "Oracle"
						oracle_tiering {
							tiers {
								move_after_unit = "Days"
								move_after = 1
								tier_type = "kOracleTierStandard"
							}
						}
					}
					extended_retention {
						schedule {
							unit = "Runs"
							frequency = 1
						}
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
						run_type = "Regular"
						config_id = "config_id"
					}
				}
				cloud_spin_targets {
					schedule {
						unit = "Runs"
						frequency = 1
					}
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
					copy_on_run_success = true
					config_id = "config_id"
					backup_run_type = "Regular"
					run_timeouts {
						timeout_mins = 1
						backup_type = "kRegular"
					}
					log_retention {
						unit = "Days"
						duration = 1
						data_lock_config {
							mode = "Compliance"
							unit = "Days"
							duration = 1
							enable_worm_on_external_target = true
						}
					}
					target {
						id = 1
						name = "name"
					}
				}
				onprem_deploy_targets {
					schedule {
						unit = "Runs"
						frequency = 1
					}
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
					copy_on_run_success = true
					config_id = "config_id"
					backup_run_type = "Regular"
					run_timeouts {
						timeout_mins = 1
						backup_type = "kRegular"
					}
					log_retention {
						unit = "Days"
						duration = 1
						data_lock_config {
							mode = "Compliance"
							unit = "Days"
							duration = 1
							enable_worm_on_external_target = true
						}
					}
					params {
						id = 1
						restore_v_mware_params {
							target_vm_folder_id = 1
							target_data_store_id = 1
							enable_copy_recovery = true
							resource_pool_id = 1
							datastore_ids = [ 1 ]
							overwrite_existing_vm = true
							power_off_and_rename_existing_vm = true
							attempt_differential_restore = true
							is_on_prem_deploy = true
						}
					}
				}
				rpaas_targets {
					schedule {
						unit = "Runs"
						frequency = 1
					}
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
					copy_on_run_success = true
					config_id = "config_id"
					backup_run_type = "Regular"
					run_timeouts {
						timeout_mins = 1
						backup_type = "kRegular"
					}
					log_retention {
						unit = "Days"
						duration = 1
						data_lock_config {
							mode = "Compliance"
							unit = "Days"
							duration = 1
							enable_worm_on_external_target = true
						}
					}
					target_id = 1
					target_name = "target_name"
					target_type = "Tape"
				}
			}
			cascaded_targets_config {
				source_cluster_id = 1
				remote_targets {
					replication_targets {
						schedule {
							unit = "Runs"
							frequency = 1
						}
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
						copy_on_run_success = true
						config_id = "config_id"
						backup_run_type = "Regular"
						run_timeouts {
							timeout_mins = 1
							backup_type = "kRegular"
						}
						log_retention {
							unit = "Days"
							duration = 1
							data_lock_config {
								mode = "Compliance"
								unit = "Days"
								duration = 1
								enable_worm_on_external_target = true
							}
						}
						target_type = "RemoteCluster"
						remote_target_config {
							cluster_id = 1
							cluster_name = "cluster_name"
						}
					}
					archival_targets {
						schedule {
							unit = "Runs"
							frequency = 1
						}
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
						copy_on_run_success = true
						config_id = "config_id"
						backup_run_type = "Regular"
						run_timeouts {
							timeout_mins = 1
							backup_type = "kRegular"
						}
						log_retention {
							unit = "Days"
							duration = 1
							data_lock_config {
								mode = "Compliance"
								unit = "Days"
								duration = 1
								enable_worm_on_external_target = true
							}
						}
						target_id = 1
						target_name = "target_name"
						target_type = "Tape"
						tier_settings {
							cloud_platform = "Oracle"
							oracle_tiering {
								tiers {
									move_after_unit = "Days"
									move_after = 1
									tier_type = "kOracleTierStandard"
								}
							}
						}
						extended_retention {
							schedule {
								unit = "Runs"
								frequency = 1
							}
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
							run_type = "Regular"
							config_id = "config_id"
						}
					}
					cloud_spin_targets {
						schedule {
							unit = "Runs"
							frequency = 1
						}
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
						copy_on_run_success = true
						config_id = "config_id"
						backup_run_type = "Regular"
						run_timeouts {
							timeout_mins = 1
							backup_type = "kRegular"
						}
						log_retention {
							unit = "Days"
							duration = 1
							data_lock_config {
								mode = "Compliance"
								unit = "Days"
								duration = 1
								enable_worm_on_external_target = true
							}
						}
						target {
							id = 1
							name = "name"
						}
					}
					onprem_deploy_targets {
						schedule {
							unit = "Runs"
							frequency = 1
						}
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
						copy_on_run_success = true
						config_id = "config_id"
						backup_run_type = "Regular"
						run_timeouts {
							timeout_mins = 1
							backup_type = "kRegular"
						}
						log_retention {
							unit = "Days"
							duration = 1
							data_lock_config {
								mode = "Compliance"
								unit = "Days"
								duration = 1
								enable_worm_on_external_target = true
							}
						}
						params {
							id = 1
							restore_v_mware_params {
								target_vm_folder_id = 1
								target_data_store_id = 1
								enable_copy_recovery = true
								resource_pool_id = 1
								datastore_ids = [ 1 ]
								overwrite_existing_vm = true
								power_off_and_rename_existing_vm = true
								attempt_differential_restore = true
								is_on_prem_deploy = true
							}
						}
					}
					rpaas_targets {
						schedule {
							unit = "Runs"
							frequency = 1
						}
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
						copy_on_run_success = true
						config_id = "config_id"
						backup_run_type = "Regular"
						run_timeouts {
							timeout_mins = 1
							backup_type = "kRegular"
						}
						log_retention {
							unit = "Days"
							duration = 1
							data_lock_config {
								mode = "Compliance"
								unit = "Days"
								duration = 1
								enable_worm_on_external_target = true
							}
						}
						target_id = 1
						target_name = "target_name"
						target_type = "Tape"
					}
				}
			}
			retry_options {
				retries = 0
				retry_interval_mins = 1
			}
			data_lock = "%s"
			version = %s
			is_cbs_enabled = %s
			last_modification_time_usecs = %s
			template_id = "%s"
		}
	`, name, description, dataLock, version, isCbsEnabled, lastModificationTimeUsecs, templateID)
}

func testAccCheckIbmProtectionPolicyExists(n string, obj backuprecoveryv1.ProtectionPolicyResponse) resource.TestCheckFunc {

	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]
		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		backupRecoveryClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).BackupRecoveryV1()
		if err != nil {
			return err
		}

		getProtectionPolicyByIdOptions := &backuprecoveryv1.GetProtectionPolicyByIdOptions{}

		getProtectionPolicyByIdOptions.SetID(rs.Primary.ID)

		protectionPolicyResponse, _, err := backupRecoveryClient.GetProtectionPolicyByID(getProtectionPolicyByIdOptions)
		if err != nil {
			return err
		}

		obj = *protectionPolicyResponse
		return nil
	}
}

func testAccCheckIbmProtectionPolicyDestroy(s *terraform.State) error {
	backupRecoveryClient, err := acc.TestAccProvider.Meta().(conns.ClientSession).BackupRecoveryV1()
	if err != nil {
		return err
	}
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "ibm_protection_policy" {
			continue
		}

		getProtectionPolicyByIdOptions := &backuprecoveryv1.GetProtectionPolicyByIdOptions{}

		getProtectionPolicyByIdOptions.SetID(rs.Primary.ID)

		// Try to find the key
		_, response, err := backupRecoveryClient.GetProtectionPolicyByID(getProtectionPolicyByIdOptions)

		if err == nil {
			return fmt.Errorf("protection_policy still exists: %s", rs.Primary.ID)
		} else if response.StatusCode != 404 {
			return fmt.Errorf("Error checking for protection_policy (%s) has been destroyed: %s", rs.Primary.ID, err)
		}
	}

	return nil
}
