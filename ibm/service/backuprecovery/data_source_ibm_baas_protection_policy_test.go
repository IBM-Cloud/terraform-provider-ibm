// Copyright IBM Corp. 2024 All Rights Reserved.
// Licensed under the Mozilla Public License v2.0

/*
 * IBM OpenAPI Terraform Generator Version: 3.94.0-fa797aec-20240814-142622
 */

package backuprecovery_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"

	acc "github.com/IBM-Cloud/terraform-provider-ibm/ibm/acctest"
	"github.com/IBM-Cloud/terraform-provider-ibm/ibm/service/backuprecovery"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/stretchr/testify/assert"
	"github.ibm.com/BackupAndRecovery/ibm-backup-recovery-sdk-go/backuprecoveryv1"
)

func TestAccIbmBaasProtectionPolicyDataSourceBasic(t *testing.T) {
	protectionPolicyResponseTenantID := fmt.Sprintf("%d", acctest.RandIntRange(10, 100))
	protectionPolicyResponseName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmBaasProtectionPolicyDataSourceConfigBasic(protectionPolicyResponseTenantID, protectionPolicyResponseName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_baas_protection_policy.baas_protection_policy_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_protection_policy.baas_protection_policy_instance", "baas_protection_policy_id"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_protection_policy.baas_protection_policy_instance", "tenant_id"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_protection_policy.baas_protection_policy_instance", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_protection_policy.baas_protection_policy_instance", "backup_policy.#"),
				),
			},
		},
	})
}

func TestAccIbmBaasProtectionPolicyDataSourceAllArgs(t *testing.T) {
	protectionPolicyResponseTenantID := fmt.Sprintf("%d", acctest.RandIntRange(10, 100))
	protectionPolicyResponseName := fmt.Sprintf("tf_name_%d", acctest.RandIntRange(10, 100))
	protectionPolicyResponseDescription := fmt.Sprintf("tf_description_%d", acctest.RandIntRange(10, 100))
	protectionPolicyResponseDataLock := "Compliance"
	protectionPolicyResponseVersion := fmt.Sprintf("%d", acctest.RandIntRange(10, 100))
	protectionPolicyResponseIsCBSEnabled := "true"
	protectionPolicyResponseLastModificationTimeUsecs := fmt.Sprintf("%d", acctest.RandIntRange(10, 100))
	protectionPolicyResponseTemplateID := fmt.Sprintf("tf_template_id_%d", acctest.RandIntRange(10, 100))

	resource.Test(t, resource.TestCase{
		PreCheck:  func() { acc.TestAccPreCheck(t) },
		Providers: acc.TestAccProviders,
		Steps: []resource.TestStep{
			resource.TestStep{
				Config: testAccCheckIbmBaasProtectionPolicyDataSourceConfig(protectionPolicyResponseTenantID, protectionPolicyResponseName, protectionPolicyResponseDescription, protectionPolicyResponseDataLock, protectionPolicyResponseVersion, protectionPolicyResponseIsCBSEnabled, protectionPolicyResponseLastModificationTimeUsecs, protectionPolicyResponseTemplateID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("data.ibm_baas_protection_policy.baas_protection_policy_instance", "id"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_protection_policy.baas_protection_policy_instance", "protection_policy_id"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_protection_policy.baas_protection_policy_instance", "tenant_id"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_protection_policy.baas_protection_policy_instance", "request_initiator_type"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_protection_policy.baas_protection_policy_instance", "name"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_protection_policy.baas_protection_policy_instance", "backup_policy.#"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_protection_policy.baas_protection_policy_instance", "description"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_protection_policy.baas_protection_policy_instance", "blackout_window.#"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_protection_policy.baas_protection_policy_instance", "blackout_window.0.day"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_protection_policy.baas_protection_policy_instance", "blackout_window.0.config_id"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_protection_policy.baas_protection_policy_instance", "extended_retention.#"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_protection_policy.baas_protection_policy_instance", "extended_retention.0.run_type"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_protection_policy.baas_protection_policy_instance", "extended_retention.0.config_id"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_protection_policy.baas_protection_policy_instance", "remote_target_policy.#"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_protection_policy.baas_protection_policy_instance", "cascaded_targets_config.#"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_protection_policy.baas_protection_policy_instance", "cascaded_targets_config.0.source_cluster_id"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_protection_policy.baas_protection_policy_instance", "retry_options.#"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_protection_policy.baas_protection_policy_instance", "data_lock"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_protection_policy.baas_protection_policy_instance", "version"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_protection_policy.baas_protection_policy_instance", "is_cbs_enabled"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_protection_policy.baas_protection_policy_instance", "last_modification_time_usecs"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_protection_policy.baas_protection_policy_instance", "template_id"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_protection_policy.baas_protection_policy_instance", "is_usable"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_protection_policy.baas_protection_policy_instance", "is_replicated"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_protection_policy.baas_protection_policy_instance", "num_protection_groups"),
					resource.TestCheckResourceAttrSet("data.ibm_baas_protection_policy.baas_protection_policy_instance", "num_protected_objects"),
				),
			},
		},
	})
}

func testAccCheckIbmBaasProtectionPolicyDataSourceConfigBasic(protectionPolicyResponseTenantID string, protectionPolicyResponseName string) string {
	return fmt.Sprintf(`
		resource "ibm_baas_protection_policy" "baas_protection_policy_instance" {
			tenant_id = %s
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
								aws_tiering {
									tiers {
										move_after_unit = "Days"
										move_after = 1
										tier_type = "kAmazonS3Standard"
									}
								}
								azure_tiering {
									tiers {
										move_after_unit = "Days"
										move_after = 1
										tier_type = "kAzureTierHot"
									}
								}
								cloud_platform = "AWS"
								google_tiering {
									tiers {
										move_after_unit = "Days"
										move_after = 1
										tier_type = "kGoogleStandard"
									}
								}
								oracle_tiering {
									tiers {
										move_after_unit = "Days"
										move_after = 1
										tier_type = "kOracleTierStandard"
									}
								}
							}
						}
						use_default_backup_target = true
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

		data "ibm_baas_protection_policy" "baas_protection_policy_instance" {
			protection_policy_id = "protection_policy_id"
			tenant_id = ibm_baas_protection_policy.baas_protection_policy_instance.tenant_id
			request_initiator_type = "UIUser"
		}
	`, protectionPolicyResponseTenantID, protectionPolicyResponseName)
}

func testAccCheckIbmBaasProtectionPolicyDataSourceConfig(protectionPolicyResponseTenantID string, protectionPolicyResponseName string, protectionPolicyResponseDescription string, protectionPolicyResponseDataLock string, protectionPolicyResponseVersion string, protectionPolicyResponseIsCBSEnabled string, protectionPolicyResponseLastModificationTimeUsecs string, protectionPolicyResponseTemplateID string) string {
	return fmt.Sprintf(`
		resource "ibm_baas_protection_policy" "baas_protection_policy_instance" {
			tenant_id = %s
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
								aws_tiering {
									tiers {
										move_after_unit = "Days"
										move_after = 1
										tier_type = "kAmazonS3Standard"
									}
								}
								azure_tiering {
									tiers {
										move_after_unit = "Days"
										move_after = 1
										tier_type = "kAzureTierHot"
									}
								}
								cloud_platform = "AWS"
								google_tiering {
									tiers {
										move_after_unit = "Days"
										move_after = 1
										tier_type = "kGoogleStandard"
									}
								}
								oracle_tiering {
									tiers {
										move_after_unit = "Days"
										move_after = 1
										tier_type = "kOracleTierStandard"
									}
								}
							}
						}
						use_default_backup_target = true
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
						duration = 0
						data_lock_config {
							mode = "Compliance"
							unit = "Days"
							duration = 1
							enable_worm_on_external_target = true
						}
					}
					aws_target_config {
						name = "name"
						region = 1
						region_name = "region_name"
						source_id = 1
					}
					azure_target_config {
						name = "name"
						resource_group = 1
						resource_group_name = "resource_group_name"
						source_id = 1
						storage_account = 1
						storage_account_name = "storage_account_name"
						storage_container = 1
						storage_container_name = "storage_container_name"
						storage_resource_group = 1
						storage_resource_group_name = "storage_resource_group_name"
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
						duration = 0
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
						aws_tiering {
							tiers {
								move_after_unit = "Days"
								move_after = 1
								tier_type = "kAmazonS3Standard"
							}
						}
						azure_tiering {
							tiers {
								move_after_unit = "Days"
								move_after = 1
								tier_type = "kAzureTierHot"
							}
						}
						cloud_platform = "AWS"
						google_tiering {
							tiers {
								move_after_unit = "Days"
								move_after = 1
								tier_type = "kGoogleStandard"
							}
						}
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
						duration = 0
						data_lock_config {
							mode = "Compliance"
							unit = "Days"
							duration = 1
							enable_worm_on_external_target = true
						}
					}
					target {
						aws_params {
							custom_tag_list {
								key = "key"
								value = "value"
							}
							region = 1
							subnet_id = 1
							vpc_id = 1
						}
						azure_params {
							availability_set_id = 1
							network_resource_group_id = 1
							resource_group_id = 1
							storage_account_id = 1
							storage_container_id = 1
							storage_resource_group_id = 1
							temp_vm_resource_group_id = 1
							temp_vm_storage_account_id = 1
							temp_vm_storage_container_id = 1
							temp_vm_subnet_id = 1
							temp_vm_virtual_network_id = 1
						}
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
						duration = 0
						data_lock_config {
							mode = "Compliance"
							unit = "Days"
							duration = 1
							enable_worm_on_external_target = true
						}
					}
					params {
						id = 1
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
						duration = 0
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
							duration = 0
							data_lock_config {
								mode = "Compliance"
								unit = "Days"
								duration = 1
								enable_worm_on_external_target = true
							}
						}
						aws_target_config {
							name = "name"
							region = 1
							region_name = "region_name"
							source_id = 1
						}
						azure_target_config {
							name = "name"
							resource_group = 1
							resource_group_name = "resource_group_name"
							source_id = 1
							storage_account = 1
							storage_account_name = "storage_account_name"
							storage_container = 1
							storage_container_name = "storage_container_name"
							storage_resource_group = 1
							storage_resource_group_name = "storage_resource_group_name"
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
							duration = 0
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
							aws_tiering {
								tiers {
									move_after_unit = "Days"
									move_after = 1
									tier_type = "kAmazonS3Standard"
								}
							}
							azure_tiering {
								tiers {
									move_after_unit = "Days"
									move_after = 1
									tier_type = "kAzureTierHot"
								}
							}
							cloud_platform = "AWS"
							google_tiering {
								tiers {
									move_after_unit = "Days"
									move_after = 1
									tier_type = "kGoogleStandard"
								}
							}
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
							duration = 0
							data_lock_config {
								mode = "Compliance"
								unit = "Days"
								duration = 1
								enable_worm_on_external_target = true
							}
						}
						target {
							aws_params {
								custom_tag_list {
									key = "key"
									value = "value"
								}
								region = 1
								subnet_id = 1
								vpc_id = 1
							}
							azure_params {
								availability_set_id = 1
								network_resource_group_id = 1
								resource_group_id = 1
								storage_account_id = 1
								storage_container_id = 1
								storage_resource_group_id = 1
								temp_vm_resource_group_id = 1
								temp_vm_storage_account_id = 1
								temp_vm_storage_container_id = 1
								temp_vm_subnet_id = 1
								temp_vm_virtual_network_id = 1
							}
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
							duration = 0
							data_lock_config {
								mode = "Compliance"
								unit = "Days"
								duration = 1
								enable_worm_on_external_target = true
							}
						}
						params {
							id = 1
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
							duration = 0
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

		data "ibm_baas_protection_policy" "baas_protection_policy_instance" {
			protection_policy_id = "protection_policy_id"
			tenant_id = ibm_baas_protection_policy.baas_protection_policy_instance.tenant_id
			request_initiator_type = "UIUser"
		}
	`, protectionPolicyResponseTenantID, protectionPolicyResponseName, protectionPolicyResponseDescription, protectionPolicyResponseDataLock, protectionPolicyResponseVersion, protectionPolicyResponseIsCBSEnabled, protectionPolicyResponseLastModificationTimeUsecs, protectionPolicyResponseTemplateID)
}

func TestDataSourceIbmBaasProtectionPolicyBackupPolicyToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		minuteScheduleModel := make(map[string]interface{})
		minuteScheduleModel["frequency"] = int(1)

		hourScheduleModel := make(map[string]interface{})
		hourScheduleModel["frequency"] = int(1)

		dayScheduleModel := make(map[string]interface{})
		dayScheduleModel["frequency"] = int(1)

		weekScheduleModel := make(map[string]interface{})
		weekScheduleModel["day_of_week"] = []string{"Sunday"}

		monthScheduleModel := make(map[string]interface{})
		monthScheduleModel["day_of_week"] = []string{"Sunday"}
		monthScheduleModel["week_of_month"] = "First"
		monthScheduleModel["day_of_month"] = int(38)

		yearScheduleModel := make(map[string]interface{})
		yearScheduleModel["day_of_year"] = "First"

		incrementalScheduleModel := make(map[string]interface{})
		incrementalScheduleModel["unit"] = "Minutes"
		incrementalScheduleModel["minute_schedule"] = []map[string]interface{}{minuteScheduleModel}
		incrementalScheduleModel["hour_schedule"] = []map[string]interface{}{hourScheduleModel}
		incrementalScheduleModel["day_schedule"] = []map[string]interface{}{dayScheduleModel}
		incrementalScheduleModel["week_schedule"] = []map[string]interface{}{weekScheduleModel}
		incrementalScheduleModel["month_schedule"] = []map[string]interface{}{monthScheduleModel}
		incrementalScheduleModel["year_schedule"] = []map[string]interface{}{yearScheduleModel}

		incrementalBackupPolicyModel := make(map[string]interface{})
		incrementalBackupPolicyModel["schedule"] = []map[string]interface{}{incrementalScheduleModel}

		fullScheduleModel := make(map[string]interface{})
		fullScheduleModel["unit"] = "Days"
		fullScheduleModel["day_schedule"] = []map[string]interface{}{dayScheduleModel}
		fullScheduleModel["week_schedule"] = []map[string]interface{}{weekScheduleModel}
		fullScheduleModel["month_schedule"] = []map[string]interface{}{monthScheduleModel}
		fullScheduleModel["year_schedule"] = []map[string]interface{}{yearScheduleModel}

		fullBackupPolicyModel := make(map[string]interface{})
		fullBackupPolicyModel["schedule"] = []map[string]interface{}{fullScheduleModel}

		dataLockConfigModel := make(map[string]interface{})
		dataLockConfigModel["mode"] = "Compliance"
		dataLockConfigModel["unit"] = "Days"
		dataLockConfigModel["duration"] = int(1)
		dataLockConfigModel["enable_worm_on_external_target"] = true

		retentionModel := make(map[string]interface{})
		retentionModel["unit"] = "Days"
		retentionModel["duration"] = int(1)
		retentionModel["data_lock_config"] = []map[string]interface{}{dataLockConfigModel}

		fullScheduleAndRetentionModel := make(map[string]interface{})
		fullScheduleAndRetentionModel["schedule"] = []map[string]interface{}{fullScheduleModel}
		fullScheduleAndRetentionModel["retention"] = []map[string]interface{}{retentionModel}

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

		primaryArchivalTargetModel := make(map[string]interface{})
		primaryArchivalTargetModel["target_id"] = int(26)
		primaryArchivalTargetModel["tier_settings"] = []map[string]interface{}{tierLevelSettingsModel}

		primaryBackupTargetModel := make(map[string]interface{})
		primaryBackupTargetModel["target_type"] = "Local"
		primaryBackupTargetModel["archival_target_settings"] = []map[string]interface{}{primaryArchivalTargetModel}
		primaryBackupTargetModel["use_default_backup_target"] = true

		regularBackupPolicyModel := make(map[string]interface{})
		regularBackupPolicyModel["incremental"] = []map[string]interface{}{incrementalBackupPolicyModel}
		regularBackupPolicyModel["full"] = []map[string]interface{}{fullBackupPolicyModel}
		regularBackupPolicyModel["full_backups"] = []map[string]interface{}{fullScheduleAndRetentionModel}
		regularBackupPolicyModel["retention"] = []map[string]interface{}{retentionModel}
		regularBackupPolicyModel["primary_backup_target"] = []map[string]interface{}{primaryBackupTargetModel}

		logScheduleModel := make(map[string]interface{})
		logScheduleModel["unit"] = "Minutes"
		logScheduleModel["minute_schedule"] = []map[string]interface{}{minuteScheduleModel}
		logScheduleModel["hour_schedule"] = []map[string]interface{}{hourScheduleModel}

		logBackupPolicyModel := make(map[string]interface{})
		logBackupPolicyModel["schedule"] = []map[string]interface{}{logScheduleModel}
		logBackupPolicyModel["retention"] = []map[string]interface{}{retentionModel}

		bmrScheduleModel := make(map[string]interface{})
		bmrScheduleModel["unit"] = "Days"
		bmrScheduleModel["day_schedule"] = []map[string]interface{}{dayScheduleModel}
		bmrScheduleModel["week_schedule"] = []map[string]interface{}{weekScheduleModel}
		bmrScheduleModel["month_schedule"] = []map[string]interface{}{monthScheduleModel}
		bmrScheduleModel["year_schedule"] = []map[string]interface{}{yearScheduleModel}

		bmrBackupPolicyModel := make(map[string]interface{})
		bmrBackupPolicyModel["schedule"] = []map[string]interface{}{bmrScheduleModel}
		bmrBackupPolicyModel["retention"] = []map[string]interface{}{retentionModel}

		cdpRetentionModel := make(map[string]interface{})
		cdpRetentionModel["unit"] = "Minutes"
		cdpRetentionModel["duration"] = int(1)
		cdpRetentionModel["data_lock_config"] = []map[string]interface{}{dataLockConfigModel}

		cdpBackupPolicyModel := make(map[string]interface{})
		cdpBackupPolicyModel["retention"] = []map[string]interface{}{cdpRetentionModel}

		storageArraySnapshotScheduleModel := make(map[string]interface{})
		storageArraySnapshotScheduleModel["unit"] = "Minutes"
		storageArraySnapshotScheduleModel["minute_schedule"] = []map[string]interface{}{minuteScheduleModel}
		storageArraySnapshotScheduleModel["hour_schedule"] = []map[string]interface{}{hourScheduleModel}
		storageArraySnapshotScheduleModel["day_schedule"] = []map[string]interface{}{dayScheduleModel}
		storageArraySnapshotScheduleModel["week_schedule"] = []map[string]interface{}{weekScheduleModel}
		storageArraySnapshotScheduleModel["month_schedule"] = []map[string]interface{}{monthScheduleModel}
		storageArraySnapshotScheduleModel["year_schedule"] = []map[string]interface{}{yearScheduleModel}

		storageArraySnapshotBackupPolicyModel := make(map[string]interface{})
		storageArraySnapshotBackupPolicyModel["schedule"] = []map[string]interface{}{storageArraySnapshotScheduleModel}
		storageArraySnapshotBackupPolicyModel["retention"] = []map[string]interface{}{retentionModel}

		cancellationTimeoutParamsModel := make(map[string]interface{})
		cancellationTimeoutParamsModel["timeout_mins"] = int(26)
		cancellationTimeoutParamsModel["backup_type"] = "kRegular"

		model := make(map[string]interface{})
		model["regular"] = []map[string]interface{}{regularBackupPolicyModel}
		model["log"] = []map[string]interface{}{logBackupPolicyModel}
		model["bmr"] = []map[string]interface{}{bmrBackupPolicyModel}
		model["cdp"] = []map[string]interface{}{cdpBackupPolicyModel}
		model["storage_array_snapshot"] = []map[string]interface{}{storageArraySnapshotBackupPolicyModel}
		model["run_timeouts"] = []map[string]interface{}{cancellationTimeoutParamsModel}

		assert.Equal(t, result, model)
	}

	minuteScheduleModel := new(backuprecoveryv1.MinuteSchedule)
	minuteScheduleModel.Frequency = core.Int64Ptr(int64(1))

	hourScheduleModel := new(backuprecoveryv1.HourSchedule)
	hourScheduleModel.Frequency = core.Int64Ptr(int64(1))

	dayScheduleModel := new(backuprecoveryv1.DaySchedule)
	dayScheduleModel.Frequency = core.Int64Ptr(int64(1))

	weekScheduleModel := new(backuprecoveryv1.WeekSchedule)
	weekScheduleModel.DayOfWeek = []string{"Sunday"}

	monthScheduleModel := new(backuprecoveryv1.MonthSchedule)
	monthScheduleModel.DayOfWeek = []string{"Sunday"}
	monthScheduleModel.WeekOfMonth = core.StringPtr("First")
	monthScheduleModel.DayOfMonth = core.Int64Ptr(int64(38))

	yearScheduleModel := new(backuprecoveryv1.YearSchedule)
	yearScheduleModel.DayOfYear = core.StringPtr("First")

	incrementalScheduleModel := new(backuprecoveryv1.IncrementalSchedule)
	incrementalScheduleModel.Unit = core.StringPtr("Minutes")
	incrementalScheduleModel.MinuteSchedule = minuteScheduleModel
	incrementalScheduleModel.HourSchedule = hourScheduleModel
	incrementalScheduleModel.DaySchedule = dayScheduleModel
	incrementalScheduleModel.WeekSchedule = weekScheduleModel
	incrementalScheduleModel.MonthSchedule = monthScheduleModel
	incrementalScheduleModel.YearSchedule = yearScheduleModel

	incrementalBackupPolicyModel := new(backuprecoveryv1.IncrementalBackupPolicy)
	incrementalBackupPolicyModel.Schedule = incrementalScheduleModel

	fullScheduleModel := new(backuprecoveryv1.FullSchedule)
	fullScheduleModel.Unit = core.StringPtr("Days")
	fullScheduleModel.DaySchedule = dayScheduleModel
	fullScheduleModel.WeekSchedule = weekScheduleModel
	fullScheduleModel.MonthSchedule = monthScheduleModel
	fullScheduleModel.YearSchedule = yearScheduleModel

	fullBackupPolicyModel := new(backuprecoveryv1.FullBackupPolicy)
	fullBackupPolicyModel.Schedule = fullScheduleModel

	dataLockConfigModel := new(backuprecoveryv1.DataLockConfig)
	dataLockConfigModel.Mode = core.StringPtr("Compliance")
	dataLockConfigModel.Unit = core.StringPtr("Days")
	dataLockConfigModel.Duration = core.Int64Ptr(int64(1))
	dataLockConfigModel.EnableWormOnExternalTarget = core.BoolPtr(true)

	retentionModel := new(backuprecoveryv1.Retention)
	retentionModel.Unit = core.StringPtr("Days")
	retentionModel.Duration = core.Int64Ptr(int64(1))
	retentionModel.DataLockConfig = dataLockConfigModel

	fullScheduleAndRetentionModel := new(backuprecoveryv1.FullScheduleAndRetention)
	fullScheduleAndRetentionModel.Schedule = fullScheduleModel
	fullScheduleAndRetentionModel.Retention = retentionModel

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

	primaryArchivalTargetModel := new(backuprecoveryv1.PrimaryArchivalTarget)
	primaryArchivalTargetModel.TargetID = core.Int64Ptr(int64(26))
	primaryArchivalTargetModel.TierSettings = tierLevelSettingsModel

	primaryBackupTargetModel := new(backuprecoveryv1.PrimaryBackupTarget)
	primaryBackupTargetModel.TargetType = core.StringPtr("Local")
	primaryBackupTargetModel.ArchivalTargetSettings = primaryArchivalTargetModel
	primaryBackupTargetModel.UseDefaultBackupTarget = core.BoolPtr(true)

	regularBackupPolicyModel := new(backuprecoveryv1.RegularBackupPolicy)
	regularBackupPolicyModel.Incremental = incrementalBackupPolicyModel
	regularBackupPolicyModel.Full = fullBackupPolicyModel
	regularBackupPolicyModel.FullBackups = []backuprecoveryv1.FullScheduleAndRetention{*fullScheduleAndRetentionModel}
	regularBackupPolicyModel.Retention = retentionModel
	regularBackupPolicyModel.PrimaryBackupTarget = primaryBackupTargetModel

	logScheduleModel := new(backuprecoveryv1.LogSchedule)
	logScheduleModel.Unit = core.StringPtr("Minutes")
	logScheduleModel.MinuteSchedule = minuteScheduleModel
	logScheduleModel.HourSchedule = hourScheduleModel

	logBackupPolicyModel := new(backuprecoveryv1.LogBackupPolicy)
	logBackupPolicyModel.Schedule = logScheduleModel
	logBackupPolicyModel.Retention = retentionModel

	bmrScheduleModel := new(backuprecoveryv1.BmrSchedule)
	bmrScheduleModel.Unit = core.StringPtr("Days")
	bmrScheduleModel.DaySchedule = dayScheduleModel
	bmrScheduleModel.WeekSchedule = weekScheduleModel
	bmrScheduleModel.MonthSchedule = monthScheduleModel
	bmrScheduleModel.YearSchedule = yearScheduleModel

	bmrBackupPolicyModel := new(backuprecoveryv1.BmrBackupPolicy)
	bmrBackupPolicyModel.Schedule = bmrScheduleModel
	bmrBackupPolicyModel.Retention = retentionModel

	cdpRetentionModel := new(backuprecoveryv1.CdpRetention)
	cdpRetentionModel.Unit = core.StringPtr("Minutes")
	cdpRetentionModel.Duration = core.Int64Ptr(int64(1))
	cdpRetentionModel.DataLockConfig = dataLockConfigModel

	cdpBackupPolicyModel := new(backuprecoveryv1.CdpBackupPolicy)
	cdpBackupPolicyModel.Retention = cdpRetentionModel

	storageArraySnapshotScheduleModel := new(backuprecoveryv1.StorageArraySnapshotSchedule)
	storageArraySnapshotScheduleModel.Unit = core.StringPtr("Minutes")
	storageArraySnapshotScheduleModel.MinuteSchedule = minuteScheduleModel
	storageArraySnapshotScheduleModel.HourSchedule = hourScheduleModel
	storageArraySnapshotScheduleModel.DaySchedule = dayScheduleModel
	storageArraySnapshotScheduleModel.WeekSchedule = weekScheduleModel
	storageArraySnapshotScheduleModel.MonthSchedule = monthScheduleModel
	storageArraySnapshotScheduleModel.YearSchedule = yearScheduleModel

	storageArraySnapshotBackupPolicyModel := new(backuprecoveryv1.StorageArraySnapshotBackupPolicy)
	storageArraySnapshotBackupPolicyModel.Schedule = storageArraySnapshotScheduleModel
	storageArraySnapshotBackupPolicyModel.Retention = retentionModel

	cancellationTimeoutParamsModel := new(backuprecoveryv1.CancellationTimeoutParams)
	cancellationTimeoutParamsModel.TimeoutMins = core.Int64Ptr(int64(26))
	cancellationTimeoutParamsModel.BackupType = core.StringPtr("kRegular")

	model := new(backuprecoveryv1.BackupPolicy)
	model.Regular = regularBackupPolicyModel
	model.Log = logBackupPolicyModel
	model.Bmr = bmrBackupPolicyModel
	model.Cdp = cdpBackupPolicyModel
	model.StorageArraySnapshot = storageArraySnapshotBackupPolicyModel
	model.RunTimeouts = []backuprecoveryv1.CancellationTimeoutParams{*cancellationTimeoutParamsModel}

	result, err := backuprecovery.DataSourceIbmBaasProtectionPolicyBackupPolicyToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasProtectionPolicyRegularBackupPolicyToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		minuteScheduleModel := make(map[string]interface{})
		minuteScheduleModel["frequency"] = int(1)

		hourScheduleModel := make(map[string]interface{})
		hourScheduleModel["frequency"] = int(1)

		dayScheduleModel := make(map[string]interface{})
		dayScheduleModel["frequency"] = int(1)

		weekScheduleModel := make(map[string]interface{})
		weekScheduleModel["day_of_week"] = []string{"Sunday"}

		monthScheduleModel := make(map[string]interface{})
		monthScheduleModel["day_of_week"] = []string{"Sunday"}
		monthScheduleModel["week_of_month"] = "First"
		monthScheduleModel["day_of_month"] = int(38)

		yearScheduleModel := make(map[string]interface{})
		yearScheduleModel["day_of_year"] = "First"

		incrementalScheduleModel := make(map[string]interface{})
		incrementalScheduleModel["unit"] = "Minutes"
		incrementalScheduleModel["minute_schedule"] = []map[string]interface{}{minuteScheduleModel}
		incrementalScheduleModel["hour_schedule"] = []map[string]interface{}{hourScheduleModel}
		incrementalScheduleModel["day_schedule"] = []map[string]interface{}{dayScheduleModel}
		incrementalScheduleModel["week_schedule"] = []map[string]interface{}{weekScheduleModel}
		incrementalScheduleModel["month_schedule"] = []map[string]interface{}{monthScheduleModel}
		incrementalScheduleModel["year_schedule"] = []map[string]interface{}{yearScheduleModel}

		incrementalBackupPolicyModel := make(map[string]interface{})
		incrementalBackupPolicyModel["schedule"] = []map[string]interface{}{incrementalScheduleModel}

		fullScheduleModel := make(map[string]interface{})
		fullScheduleModel["unit"] = "Days"
		fullScheduleModel["day_schedule"] = []map[string]interface{}{dayScheduleModel}
		fullScheduleModel["week_schedule"] = []map[string]interface{}{weekScheduleModel}
		fullScheduleModel["month_schedule"] = []map[string]interface{}{monthScheduleModel}
		fullScheduleModel["year_schedule"] = []map[string]interface{}{yearScheduleModel}

		fullBackupPolicyModel := make(map[string]interface{})
		fullBackupPolicyModel["schedule"] = []map[string]interface{}{fullScheduleModel}

		dataLockConfigModel := make(map[string]interface{})
		dataLockConfigModel["mode"] = "Compliance"
		dataLockConfigModel["unit"] = "Days"
		dataLockConfigModel["duration"] = int(1)
		dataLockConfigModel["enable_worm_on_external_target"] = true

		retentionModel := make(map[string]interface{})
		retentionModel["unit"] = "Days"
		retentionModel["duration"] = int(1)
		retentionModel["data_lock_config"] = []map[string]interface{}{dataLockConfigModel}

		fullScheduleAndRetentionModel := make(map[string]interface{})
		fullScheduleAndRetentionModel["schedule"] = []map[string]interface{}{fullScheduleModel}
		fullScheduleAndRetentionModel["retention"] = []map[string]interface{}{retentionModel}

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

		primaryArchivalTargetModel := make(map[string]interface{})
		primaryArchivalTargetModel["target_id"] = int(26)
		primaryArchivalTargetModel["tier_settings"] = []map[string]interface{}{tierLevelSettingsModel}

		primaryBackupTargetModel := make(map[string]interface{})
		primaryBackupTargetModel["target_type"] = "Local"
		primaryBackupTargetModel["archival_target_settings"] = []map[string]interface{}{primaryArchivalTargetModel}
		primaryBackupTargetModel["use_default_backup_target"] = true

		model := make(map[string]interface{})
		model["incremental"] = []map[string]interface{}{incrementalBackupPolicyModel}
		model["full"] = []map[string]interface{}{fullBackupPolicyModel}
		model["full_backups"] = []map[string]interface{}{fullScheduleAndRetentionModel}
		model["retention"] = []map[string]interface{}{retentionModel}
		model["primary_backup_target"] = []map[string]interface{}{primaryBackupTargetModel}

		assert.Equal(t, result, model)
	}

	minuteScheduleModel := new(backuprecoveryv1.MinuteSchedule)
	minuteScheduleModel.Frequency = core.Int64Ptr(int64(1))

	hourScheduleModel := new(backuprecoveryv1.HourSchedule)
	hourScheduleModel.Frequency = core.Int64Ptr(int64(1))

	dayScheduleModel := new(backuprecoveryv1.DaySchedule)
	dayScheduleModel.Frequency = core.Int64Ptr(int64(1))

	weekScheduleModel := new(backuprecoveryv1.WeekSchedule)
	weekScheduleModel.DayOfWeek = []string{"Sunday"}

	monthScheduleModel := new(backuprecoveryv1.MonthSchedule)
	monthScheduleModel.DayOfWeek = []string{"Sunday"}
	monthScheduleModel.WeekOfMonth = core.StringPtr("First")
	monthScheduleModel.DayOfMonth = core.Int64Ptr(int64(38))

	yearScheduleModel := new(backuprecoveryv1.YearSchedule)
	yearScheduleModel.DayOfYear = core.StringPtr("First")

	incrementalScheduleModel := new(backuprecoveryv1.IncrementalSchedule)
	incrementalScheduleModel.Unit = core.StringPtr("Minutes")
	incrementalScheduleModel.MinuteSchedule = minuteScheduleModel
	incrementalScheduleModel.HourSchedule = hourScheduleModel
	incrementalScheduleModel.DaySchedule = dayScheduleModel
	incrementalScheduleModel.WeekSchedule = weekScheduleModel
	incrementalScheduleModel.MonthSchedule = monthScheduleModel
	incrementalScheduleModel.YearSchedule = yearScheduleModel

	incrementalBackupPolicyModel := new(backuprecoveryv1.IncrementalBackupPolicy)
	incrementalBackupPolicyModel.Schedule = incrementalScheduleModel

	fullScheduleModel := new(backuprecoveryv1.FullSchedule)
	fullScheduleModel.Unit = core.StringPtr("Days")
	fullScheduleModel.DaySchedule = dayScheduleModel
	fullScheduleModel.WeekSchedule = weekScheduleModel
	fullScheduleModel.MonthSchedule = monthScheduleModel
	fullScheduleModel.YearSchedule = yearScheduleModel

	fullBackupPolicyModel := new(backuprecoveryv1.FullBackupPolicy)
	fullBackupPolicyModel.Schedule = fullScheduleModel

	dataLockConfigModel := new(backuprecoveryv1.DataLockConfig)
	dataLockConfigModel.Mode = core.StringPtr("Compliance")
	dataLockConfigModel.Unit = core.StringPtr("Days")
	dataLockConfigModel.Duration = core.Int64Ptr(int64(1))
	dataLockConfigModel.EnableWormOnExternalTarget = core.BoolPtr(true)

	retentionModel := new(backuprecoveryv1.Retention)
	retentionModel.Unit = core.StringPtr("Days")
	retentionModel.Duration = core.Int64Ptr(int64(1))
	retentionModel.DataLockConfig = dataLockConfigModel

	fullScheduleAndRetentionModel := new(backuprecoveryv1.FullScheduleAndRetention)
	fullScheduleAndRetentionModel.Schedule = fullScheduleModel
	fullScheduleAndRetentionModel.Retention = retentionModel

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

	primaryArchivalTargetModel := new(backuprecoveryv1.PrimaryArchivalTarget)
	primaryArchivalTargetModel.TargetID = core.Int64Ptr(int64(26))
	primaryArchivalTargetModel.TierSettings = tierLevelSettingsModel

	primaryBackupTargetModel := new(backuprecoveryv1.PrimaryBackupTarget)
	primaryBackupTargetModel.TargetType = core.StringPtr("Local")
	primaryBackupTargetModel.ArchivalTargetSettings = primaryArchivalTargetModel
	primaryBackupTargetModel.UseDefaultBackupTarget = core.BoolPtr(true)

	model := new(backuprecoveryv1.RegularBackupPolicy)
	model.Incremental = incrementalBackupPolicyModel
	model.Full = fullBackupPolicyModel
	model.FullBackups = []backuprecoveryv1.FullScheduleAndRetention{*fullScheduleAndRetentionModel}
	model.Retention = retentionModel
	model.PrimaryBackupTarget = primaryBackupTargetModel

	result, err := backuprecovery.DataSourceIbmBaasProtectionPolicyRegularBackupPolicyToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasProtectionPolicyIncrementalBackupPolicyToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		minuteScheduleModel := make(map[string]interface{})
		minuteScheduleModel["frequency"] = int(1)

		hourScheduleModel := make(map[string]interface{})
		hourScheduleModel["frequency"] = int(1)

		dayScheduleModel := make(map[string]interface{})
		dayScheduleModel["frequency"] = int(1)

		weekScheduleModel := make(map[string]interface{})
		weekScheduleModel["day_of_week"] = []string{"Sunday"}

		monthScheduleModel := make(map[string]interface{})
		monthScheduleModel["day_of_week"] = []string{"Sunday"}
		monthScheduleModel["week_of_month"] = "First"
		monthScheduleModel["day_of_month"] = int(38)

		yearScheduleModel := make(map[string]interface{})
		yearScheduleModel["day_of_year"] = "First"

		incrementalScheduleModel := make(map[string]interface{})
		incrementalScheduleModel["unit"] = "Minutes"
		incrementalScheduleModel["minute_schedule"] = []map[string]interface{}{minuteScheduleModel}
		incrementalScheduleModel["hour_schedule"] = []map[string]interface{}{hourScheduleModel}
		incrementalScheduleModel["day_schedule"] = []map[string]interface{}{dayScheduleModel}
		incrementalScheduleModel["week_schedule"] = []map[string]interface{}{weekScheduleModel}
		incrementalScheduleModel["month_schedule"] = []map[string]interface{}{monthScheduleModel}
		incrementalScheduleModel["year_schedule"] = []map[string]interface{}{yearScheduleModel}

		model := make(map[string]interface{})
		model["schedule"] = []map[string]interface{}{incrementalScheduleModel}

		assert.Equal(t, result, model)
	}

	minuteScheduleModel := new(backuprecoveryv1.MinuteSchedule)
	minuteScheduleModel.Frequency = core.Int64Ptr(int64(1))

	hourScheduleModel := new(backuprecoveryv1.HourSchedule)
	hourScheduleModel.Frequency = core.Int64Ptr(int64(1))

	dayScheduleModel := new(backuprecoveryv1.DaySchedule)
	dayScheduleModel.Frequency = core.Int64Ptr(int64(1))

	weekScheduleModel := new(backuprecoveryv1.WeekSchedule)
	weekScheduleModel.DayOfWeek = []string{"Sunday"}

	monthScheduleModel := new(backuprecoveryv1.MonthSchedule)
	monthScheduleModel.DayOfWeek = []string{"Sunday"}
	monthScheduleModel.WeekOfMonth = core.StringPtr("First")
	monthScheduleModel.DayOfMonth = core.Int64Ptr(int64(38))

	yearScheduleModel := new(backuprecoveryv1.YearSchedule)
	yearScheduleModel.DayOfYear = core.StringPtr("First")

	incrementalScheduleModel := new(backuprecoveryv1.IncrementalSchedule)
	incrementalScheduleModel.Unit = core.StringPtr("Minutes")
	incrementalScheduleModel.MinuteSchedule = minuteScheduleModel
	incrementalScheduleModel.HourSchedule = hourScheduleModel
	incrementalScheduleModel.DaySchedule = dayScheduleModel
	incrementalScheduleModel.WeekSchedule = weekScheduleModel
	incrementalScheduleModel.MonthSchedule = monthScheduleModel
	incrementalScheduleModel.YearSchedule = yearScheduleModel

	model := new(backuprecoveryv1.IncrementalBackupPolicy)
	model.Schedule = incrementalScheduleModel

	result, err := backuprecovery.DataSourceIbmBaasProtectionPolicyIncrementalBackupPolicyToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasProtectionPolicyIncrementalScheduleToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		minuteScheduleModel := make(map[string]interface{})
		minuteScheduleModel["frequency"] = int(1)

		hourScheduleModel := make(map[string]interface{})
		hourScheduleModel["frequency"] = int(1)

		dayScheduleModel := make(map[string]interface{})
		dayScheduleModel["frequency"] = int(1)

		weekScheduleModel := make(map[string]interface{})
		weekScheduleModel["day_of_week"] = []string{"Sunday"}

		monthScheduleModel := make(map[string]interface{})
		monthScheduleModel["day_of_week"] = []string{"Sunday"}
		monthScheduleModel["week_of_month"] = "First"
		monthScheduleModel["day_of_month"] = int(38)

		yearScheduleModel := make(map[string]interface{})
		yearScheduleModel["day_of_year"] = "First"

		model := make(map[string]interface{})
		model["unit"] = "Minutes"
		model["minute_schedule"] = []map[string]interface{}{minuteScheduleModel}
		model["hour_schedule"] = []map[string]interface{}{hourScheduleModel}
		model["day_schedule"] = []map[string]interface{}{dayScheduleModel}
		model["week_schedule"] = []map[string]interface{}{weekScheduleModel}
		model["month_schedule"] = []map[string]interface{}{monthScheduleModel}
		model["year_schedule"] = []map[string]interface{}{yearScheduleModel}

		assert.Equal(t, result, model)
	}

	minuteScheduleModel := new(backuprecoveryv1.MinuteSchedule)
	minuteScheduleModel.Frequency = core.Int64Ptr(int64(1))

	hourScheduleModel := new(backuprecoveryv1.HourSchedule)
	hourScheduleModel.Frequency = core.Int64Ptr(int64(1))

	dayScheduleModel := new(backuprecoveryv1.DaySchedule)
	dayScheduleModel.Frequency = core.Int64Ptr(int64(1))

	weekScheduleModel := new(backuprecoveryv1.WeekSchedule)
	weekScheduleModel.DayOfWeek = []string{"Sunday"}

	monthScheduleModel := new(backuprecoveryv1.MonthSchedule)
	monthScheduleModel.DayOfWeek = []string{"Sunday"}
	monthScheduleModel.WeekOfMonth = core.StringPtr("First")
	monthScheduleModel.DayOfMonth = core.Int64Ptr(int64(38))

	yearScheduleModel := new(backuprecoveryv1.YearSchedule)
	yearScheduleModel.DayOfYear = core.StringPtr("First")

	model := new(backuprecoveryv1.IncrementalSchedule)
	model.Unit = core.StringPtr("Minutes")
	model.MinuteSchedule = minuteScheduleModel
	model.HourSchedule = hourScheduleModel
	model.DaySchedule = dayScheduleModel
	model.WeekSchedule = weekScheduleModel
	model.MonthSchedule = monthScheduleModel
	model.YearSchedule = yearScheduleModel

	result, err := backuprecovery.DataSourceIbmBaasProtectionPolicyIncrementalScheduleToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasProtectionPolicyMinuteScheduleToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["frequency"] = int(1)

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.MinuteSchedule)
	model.Frequency = core.Int64Ptr(int64(1))

	result, err := backuprecovery.DataSourceIbmBaasProtectionPolicyMinuteScheduleToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasProtectionPolicyHourScheduleToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["frequency"] = int(1)

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.HourSchedule)
	model.Frequency = core.Int64Ptr(int64(1))

	result, err := backuprecovery.DataSourceIbmBaasProtectionPolicyHourScheduleToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasProtectionPolicyDayScheduleToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["frequency"] = int(1)

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.DaySchedule)
	model.Frequency = core.Int64Ptr(int64(1))

	result, err := backuprecovery.DataSourceIbmBaasProtectionPolicyDayScheduleToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasProtectionPolicyWeekScheduleToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["day_of_week"] = []string{"Sunday"}

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.WeekSchedule)
	model.DayOfWeek = []string{"Sunday"}

	result, err := backuprecovery.DataSourceIbmBaasProtectionPolicyWeekScheduleToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasProtectionPolicyMonthScheduleToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["day_of_week"] = []string{"Sunday"}
		model["week_of_month"] = "First"
		model["day_of_month"] = int(38)

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.MonthSchedule)
	model.DayOfWeek = []string{"Sunday"}
	model.WeekOfMonth = core.StringPtr("First")
	model.DayOfMonth = core.Int64Ptr(int64(38))

	result, err := backuprecovery.DataSourceIbmBaasProtectionPolicyMonthScheduleToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasProtectionPolicyYearScheduleToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["day_of_year"] = "First"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.YearSchedule)
	model.DayOfYear = core.StringPtr("First")

	result, err := backuprecovery.DataSourceIbmBaasProtectionPolicyYearScheduleToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasProtectionPolicyFullBackupPolicyToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		dayScheduleModel := make(map[string]interface{})
		dayScheduleModel["frequency"] = int(1)

		weekScheduleModel := make(map[string]interface{})
		weekScheduleModel["day_of_week"] = []string{"Sunday"}

		monthScheduleModel := make(map[string]interface{})
		monthScheduleModel["day_of_week"] = []string{"Sunday"}
		monthScheduleModel["week_of_month"] = "First"
		monthScheduleModel["day_of_month"] = int(38)

		yearScheduleModel := make(map[string]interface{})
		yearScheduleModel["day_of_year"] = "First"

		fullScheduleModel := make(map[string]interface{})
		fullScheduleModel["unit"] = "Days"
		fullScheduleModel["day_schedule"] = []map[string]interface{}{dayScheduleModel}
		fullScheduleModel["week_schedule"] = []map[string]interface{}{weekScheduleModel}
		fullScheduleModel["month_schedule"] = []map[string]interface{}{monthScheduleModel}
		fullScheduleModel["year_schedule"] = []map[string]interface{}{yearScheduleModel}

		model := make(map[string]interface{})
		model["schedule"] = []map[string]interface{}{fullScheduleModel}

		assert.Equal(t, result, model)
	}

	dayScheduleModel := new(backuprecoveryv1.DaySchedule)
	dayScheduleModel.Frequency = core.Int64Ptr(int64(1))

	weekScheduleModel := new(backuprecoveryv1.WeekSchedule)
	weekScheduleModel.DayOfWeek = []string{"Sunday"}

	monthScheduleModel := new(backuprecoveryv1.MonthSchedule)
	monthScheduleModel.DayOfWeek = []string{"Sunday"}
	monthScheduleModel.WeekOfMonth = core.StringPtr("First")
	monthScheduleModel.DayOfMonth = core.Int64Ptr(int64(38))

	yearScheduleModel := new(backuprecoveryv1.YearSchedule)
	yearScheduleModel.DayOfYear = core.StringPtr("First")

	fullScheduleModel := new(backuprecoveryv1.FullSchedule)
	fullScheduleModel.Unit = core.StringPtr("Days")
	fullScheduleModel.DaySchedule = dayScheduleModel
	fullScheduleModel.WeekSchedule = weekScheduleModel
	fullScheduleModel.MonthSchedule = monthScheduleModel
	fullScheduleModel.YearSchedule = yearScheduleModel

	model := new(backuprecoveryv1.FullBackupPolicy)
	model.Schedule = fullScheduleModel

	result, err := backuprecovery.DataSourceIbmBaasProtectionPolicyFullBackupPolicyToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasProtectionPolicyFullScheduleToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		dayScheduleModel := make(map[string]interface{})
		dayScheduleModel["frequency"] = int(1)

		weekScheduleModel := make(map[string]interface{})
		weekScheduleModel["day_of_week"] = []string{"Sunday"}

		monthScheduleModel := make(map[string]interface{})
		monthScheduleModel["day_of_week"] = []string{"Sunday"}
		monthScheduleModel["week_of_month"] = "First"
		monthScheduleModel["day_of_month"] = int(38)

		yearScheduleModel := make(map[string]interface{})
		yearScheduleModel["day_of_year"] = "First"

		model := make(map[string]interface{})
		model["unit"] = "Days"
		model["day_schedule"] = []map[string]interface{}{dayScheduleModel}
		model["week_schedule"] = []map[string]interface{}{weekScheduleModel}
		model["month_schedule"] = []map[string]interface{}{monthScheduleModel}
		model["year_schedule"] = []map[string]interface{}{yearScheduleModel}

		assert.Equal(t, result, model)
	}

	dayScheduleModel := new(backuprecoveryv1.DaySchedule)
	dayScheduleModel.Frequency = core.Int64Ptr(int64(1))

	weekScheduleModel := new(backuprecoveryv1.WeekSchedule)
	weekScheduleModel.DayOfWeek = []string{"Sunday"}

	monthScheduleModel := new(backuprecoveryv1.MonthSchedule)
	monthScheduleModel.DayOfWeek = []string{"Sunday"}
	monthScheduleModel.WeekOfMonth = core.StringPtr("First")
	monthScheduleModel.DayOfMonth = core.Int64Ptr(int64(38))

	yearScheduleModel := new(backuprecoveryv1.YearSchedule)
	yearScheduleModel.DayOfYear = core.StringPtr("First")

	model := new(backuprecoveryv1.FullSchedule)
	model.Unit = core.StringPtr("Days")
	model.DaySchedule = dayScheduleModel
	model.WeekSchedule = weekScheduleModel
	model.MonthSchedule = monthScheduleModel
	model.YearSchedule = yearScheduleModel

	result, err := backuprecovery.DataSourceIbmBaasProtectionPolicyFullScheduleToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasProtectionPolicyFullScheduleAndRetentionToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		dayScheduleModel := make(map[string]interface{})
		dayScheduleModel["frequency"] = int(1)

		weekScheduleModel := make(map[string]interface{})
		weekScheduleModel["day_of_week"] = []string{"Sunday"}

		monthScheduleModel := make(map[string]interface{})
		monthScheduleModel["day_of_week"] = []string{"Sunday"}
		monthScheduleModel["week_of_month"] = "First"
		monthScheduleModel["day_of_month"] = int(38)

		yearScheduleModel := make(map[string]interface{})
		yearScheduleModel["day_of_year"] = "First"

		fullScheduleModel := make(map[string]interface{})
		fullScheduleModel["unit"] = "Days"
		fullScheduleModel["day_schedule"] = []map[string]interface{}{dayScheduleModel}
		fullScheduleModel["week_schedule"] = []map[string]interface{}{weekScheduleModel}
		fullScheduleModel["month_schedule"] = []map[string]interface{}{monthScheduleModel}
		fullScheduleModel["year_schedule"] = []map[string]interface{}{yearScheduleModel}

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
		model["schedule"] = []map[string]interface{}{fullScheduleModel}
		model["retention"] = []map[string]interface{}{retentionModel}

		assert.Equal(t, result, model)
	}

	dayScheduleModel := new(backuprecoveryv1.DaySchedule)
	dayScheduleModel.Frequency = core.Int64Ptr(int64(1))

	weekScheduleModel := new(backuprecoveryv1.WeekSchedule)
	weekScheduleModel.DayOfWeek = []string{"Sunday"}

	monthScheduleModel := new(backuprecoveryv1.MonthSchedule)
	monthScheduleModel.DayOfWeek = []string{"Sunday"}
	monthScheduleModel.WeekOfMonth = core.StringPtr("First")
	monthScheduleModel.DayOfMonth = core.Int64Ptr(int64(38))

	yearScheduleModel := new(backuprecoveryv1.YearSchedule)
	yearScheduleModel.DayOfYear = core.StringPtr("First")

	fullScheduleModel := new(backuprecoveryv1.FullSchedule)
	fullScheduleModel.Unit = core.StringPtr("Days")
	fullScheduleModel.DaySchedule = dayScheduleModel
	fullScheduleModel.WeekSchedule = weekScheduleModel
	fullScheduleModel.MonthSchedule = monthScheduleModel
	fullScheduleModel.YearSchedule = yearScheduleModel

	dataLockConfigModel := new(backuprecoveryv1.DataLockConfig)
	dataLockConfigModel.Mode = core.StringPtr("Compliance")
	dataLockConfigModel.Unit = core.StringPtr("Days")
	dataLockConfigModel.Duration = core.Int64Ptr(int64(1))
	dataLockConfigModel.EnableWormOnExternalTarget = core.BoolPtr(true)

	retentionModel := new(backuprecoveryv1.Retention)
	retentionModel.Unit = core.StringPtr("Days")
	retentionModel.Duration = core.Int64Ptr(int64(1))
	retentionModel.DataLockConfig = dataLockConfigModel

	model := new(backuprecoveryv1.FullScheduleAndRetention)
	model.Schedule = fullScheduleModel
	model.Retention = retentionModel

	result, err := backuprecovery.DataSourceIbmBaasProtectionPolicyFullScheduleAndRetentionToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasProtectionPolicyRetentionToMap(t *testing.T) {
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

	result, err := backuprecovery.DataSourceIbmBaasProtectionPolicyRetentionToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasProtectionPolicyDataLockConfigToMap(t *testing.T) {
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

	result, err := backuprecovery.DataSourceIbmBaasProtectionPolicyDataLockConfigToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasProtectionPolicyPrimaryBackupTargetToMap(t *testing.T) {
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

		tierLevelSettingsModel := make(map[string]interface{})
		tierLevelSettingsModel["aws_tiering"] = []map[string]interface{}{awsTiersModel}
		tierLevelSettingsModel["azure_tiering"] = []map[string]interface{}{azureTiersModel}
		tierLevelSettingsModel["cloud_platform"] = "AWS"
		tierLevelSettingsModel["google_tiering"] = []map[string]interface{}{googleTiersModel}
		tierLevelSettingsModel["oracle_tiering"] = []map[string]interface{}{oracleTiersModel}

		primaryArchivalTargetModel := make(map[string]interface{})
		primaryArchivalTargetModel["target_id"] = int(26)
		primaryArchivalTargetModel["tier_settings"] = []map[string]interface{}{tierLevelSettingsModel}

		model := make(map[string]interface{})
		model["target_type"] = "Local"
		model["archival_target_settings"] = []map[string]interface{}{primaryArchivalTargetModel}
		model["use_default_backup_target"] = true

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

	tierLevelSettingsModel := new(backuprecoveryv1.TierLevelSettings)
	tierLevelSettingsModel.AwsTiering = awsTiersModel
	tierLevelSettingsModel.AzureTiering = azureTiersModel
	tierLevelSettingsModel.CloudPlatform = core.StringPtr("AWS")
	tierLevelSettingsModel.GoogleTiering = googleTiersModel
	tierLevelSettingsModel.OracleTiering = oracleTiersModel

	primaryArchivalTargetModel := new(backuprecoveryv1.PrimaryArchivalTarget)
	primaryArchivalTargetModel.TargetID = core.Int64Ptr(int64(26))
	primaryArchivalTargetModel.TierSettings = tierLevelSettingsModel

	model := new(backuprecoveryv1.PrimaryBackupTarget)
	model.TargetType = core.StringPtr("Local")
	model.ArchivalTargetSettings = primaryArchivalTargetModel
	model.UseDefaultBackupTarget = core.BoolPtr(true)

	result, err := backuprecovery.DataSourceIbmBaasProtectionPolicyPrimaryBackupTargetToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasProtectionPolicyPrimaryArchivalTargetToMap(t *testing.T) {
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

		tierLevelSettingsModel := make(map[string]interface{})
		tierLevelSettingsModel["aws_tiering"] = []map[string]interface{}{awsTiersModel}
		tierLevelSettingsModel["azure_tiering"] = []map[string]interface{}{azureTiersModel}
		tierLevelSettingsModel["cloud_platform"] = "AWS"
		tierLevelSettingsModel["google_tiering"] = []map[string]interface{}{googleTiersModel}
		tierLevelSettingsModel["oracle_tiering"] = []map[string]interface{}{oracleTiersModel}

		model := make(map[string]interface{})
		model["target_id"] = int(26)
		model["target_name"] = "testString"
		model["tier_settings"] = []map[string]interface{}{tierLevelSettingsModel}

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

	tierLevelSettingsModel := new(backuprecoveryv1.TierLevelSettings)
	tierLevelSettingsModel.AwsTiering = awsTiersModel
	tierLevelSettingsModel.AzureTiering = azureTiersModel
	tierLevelSettingsModel.CloudPlatform = core.StringPtr("AWS")
	tierLevelSettingsModel.GoogleTiering = googleTiersModel
	tierLevelSettingsModel.OracleTiering = oracleTiersModel

	model := new(backuprecoveryv1.PrimaryArchivalTarget)
	model.TargetID = core.Int64Ptr(int64(26))
	model.TargetName = core.StringPtr("testString")
	model.TierSettings = tierLevelSettingsModel

	result, err := backuprecovery.DataSourceIbmBaasProtectionPolicyPrimaryArchivalTargetToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasProtectionPolicyTierLevelSettingsToMap(t *testing.T) {
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

	result, err := backuprecovery.DataSourceIbmBaasProtectionPolicyTierLevelSettingsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasProtectionPolicyAWSTiersToMap(t *testing.T) {
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

	result, err := backuprecovery.DataSourceIbmBaasProtectionPolicyAWSTiersToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasProtectionPolicyAWSTierToMap(t *testing.T) {
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

	result, err := backuprecovery.DataSourceIbmBaasProtectionPolicyAWSTierToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasProtectionPolicyAzureTiersToMap(t *testing.T) {
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

	result, err := backuprecovery.DataSourceIbmBaasProtectionPolicyAzureTiersToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasProtectionPolicyAzureTierToMap(t *testing.T) {
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

	result, err := backuprecovery.DataSourceIbmBaasProtectionPolicyAzureTierToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasProtectionPolicyGoogleTiersToMap(t *testing.T) {
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

	result, err := backuprecovery.DataSourceIbmBaasProtectionPolicyGoogleTiersToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasProtectionPolicyGoogleTierToMap(t *testing.T) {
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

	result, err := backuprecovery.DataSourceIbmBaasProtectionPolicyGoogleTierToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasProtectionPolicyOracleTiersToMap(t *testing.T) {
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

	result, err := backuprecovery.DataSourceIbmBaasProtectionPolicyOracleTiersToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasProtectionPolicyOracleTierToMap(t *testing.T) {
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

	result, err := backuprecovery.DataSourceIbmBaasProtectionPolicyOracleTierToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasProtectionPolicyLogBackupPolicyToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		minuteScheduleModel := make(map[string]interface{})
		minuteScheduleModel["frequency"] = int(1)

		hourScheduleModel := make(map[string]interface{})
		hourScheduleModel["frequency"] = int(1)

		logScheduleModel := make(map[string]interface{})
		logScheduleModel["unit"] = "Minutes"
		logScheduleModel["minute_schedule"] = []map[string]interface{}{minuteScheduleModel}
		logScheduleModel["hour_schedule"] = []map[string]interface{}{hourScheduleModel}

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
		model["schedule"] = []map[string]interface{}{logScheduleModel}
		model["retention"] = []map[string]interface{}{retentionModel}

		assert.Equal(t, result, model)
	}

	minuteScheduleModel := new(backuprecoveryv1.MinuteSchedule)
	minuteScheduleModel.Frequency = core.Int64Ptr(int64(1))

	hourScheduleModel := new(backuprecoveryv1.HourSchedule)
	hourScheduleModel.Frequency = core.Int64Ptr(int64(1))

	logScheduleModel := new(backuprecoveryv1.LogSchedule)
	logScheduleModel.Unit = core.StringPtr("Minutes")
	logScheduleModel.MinuteSchedule = minuteScheduleModel
	logScheduleModel.HourSchedule = hourScheduleModel

	dataLockConfigModel := new(backuprecoveryv1.DataLockConfig)
	dataLockConfigModel.Mode = core.StringPtr("Compliance")
	dataLockConfigModel.Unit = core.StringPtr("Days")
	dataLockConfigModel.Duration = core.Int64Ptr(int64(1))
	dataLockConfigModel.EnableWormOnExternalTarget = core.BoolPtr(true)

	retentionModel := new(backuprecoveryv1.Retention)
	retentionModel.Unit = core.StringPtr("Days")
	retentionModel.Duration = core.Int64Ptr(int64(1))
	retentionModel.DataLockConfig = dataLockConfigModel

	model := new(backuprecoveryv1.LogBackupPolicy)
	model.Schedule = logScheduleModel
	model.Retention = retentionModel

	result, err := backuprecovery.DataSourceIbmBaasProtectionPolicyLogBackupPolicyToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasProtectionPolicyLogScheduleToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		minuteScheduleModel := make(map[string]interface{})
		minuteScheduleModel["frequency"] = int(1)

		hourScheduleModel := make(map[string]interface{})
		hourScheduleModel["frequency"] = int(1)

		model := make(map[string]interface{})
		model["unit"] = "Minutes"
		model["minute_schedule"] = []map[string]interface{}{minuteScheduleModel}
		model["hour_schedule"] = []map[string]interface{}{hourScheduleModel}

		assert.Equal(t, result, model)
	}

	minuteScheduleModel := new(backuprecoveryv1.MinuteSchedule)
	minuteScheduleModel.Frequency = core.Int64Ptr(int64(1))

	hourScheduleModel := new(backuprecoveryv1.HourSchedule)
	hourScheduleModel.Frequency = core.Int64Ptr(int64(1))

	model := new(backuprecoveryv1.LogSchedule)
	model.Unit = core.StringPtr("Minutes")
	model.MinuteSchedule = minuteScheduleModel
	model.HourSchedule = hourScheduleModel

	result, err := backuprecovery.DataSourceIbmBaasProtectionPolicyLogScheduleToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasProtectionPolicyBmrBackupPolicyToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		dayScheduleModel := make(map[string]interface{})
		dayScheduleModel["frequency"] = int(1)

		weekScheduleModel := make(map[string]interface{})
		weekScheduleModel["day_of_week"] = []string{"Sunday"}

		monthScheduleModel := make(map[string]interface{})
		monthScheduleModel["day_of_week"] = []string{"Sunday"}
		monthScheduleModel["week_of_month"] = "First"
		monthScheduleModel["day_of_month"] = int(38)

		yearScheduleModel := make(map[string]interface{})
		yearScheduleModel["day_of_year"] = "First"

		bmrScheduleModel := make(map[string]interface{})
		bmrScheduleModel["unit"] = "Days"
		bmrScheduleModel["day_schedule"] = []map[string]interface{}{dayScheduleModel}
		bmrScheduleModel["week_schedule"] = []map[string]interface{}{weekScheduleModel}
		bmrScheduleModel["month_schedule"] = []map[string]interface{}{monthScheduleModel}
		bmrScheduleModel["year_schedule"] = []map[string]interface{}{yearScheduleModel}

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
		model["schedule"] = []map[string]interface{}{bmrScheduleModel}
		model["retention"] = []map[string]interface{}{retentionModel}

		assert.Equal(t, result, model)
	}

	dayScheduleModel := new(backuprecoveryv1.DaySchedule)
	dayScheduleModel.Frequency = core.Int64Ptr(int64(1))

	weekScheduleModel := new(backuprecoveryv1.WeekSchedule)
	weekScheduleModel.DayOfWeek = []string{"Sunday"}

	monthScheduleModel := new(backuprecoveryv1.MonthSchedule)
	monthScheduleModel.DayOfWeek = []string{"Sunday"}
	monthScheduleModel.WeekOfMonth = core.StringPtr("First")
	monthScheduleModel.DayOfMonth = core.Int64Ptr(int64(38))

	yearScheduleModel := new(backuprecoveryv1.YearSchedule)
	yearScheduleModel.DayOfYear = core.StringPtr("First")

	bmrScheduleModel := new(backuprecoveryv1.BmrSchedule)
	bmrScheduleModel.Unit = core.StringPtr("Days")
	bmrScheduleModel.DaySchedule = dayScheduleModel
	bmrScheduleModel.WeekSchedule = weekScheduleModel
	bmrScheduleModel.MonthSchedule = monthScheduleModel
	bmrScheduleModel.YearSchedule = yearScheduleModel

	dataLockConfigModel := new(backuprecoveryv1.DataLockConfig)
	dataLockConfigModel.Mode = core.StringPtr("Compliance")
	dataLockConfigModel.Unit = core.StringPtr("Days")
	dataLockConfigModel.Duration = core.Int64Ptr(int64(1))
	dataLockConfigModel.EnableWormOnExternalTarget = core.BoolPtr(true)

	retentionModel := new(backuprecoveryv1.Retention)
	retentionModel.Unit = core.StringPtr("Days")
	retentionModel.Duration = core.Int64Ptr(int64(1))
	retentionModel.DataLockConfig = dataLockConfigModel

	model := new(backuprecoveryv1.BmrBackupPolicy)
	model.Schedule = bmrScheduleModel
	model.Retention = retentionModel

	result, err := backuprecovery.DataSourceIbmBaasProtectionPolicyBmrBackupPolicyToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasProtectionPolicyBmrScheduleToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		dayScheduleModel := make(map[string]interface{})
		dayScheduleModel["frequency"] = int(1)

		weekScheduleModel := make(map[string]interface{})
		weekScheduleModel["day_of_week"] = []string{"Sunday"}

		monthScheduleModel := make(map[string]interface{})
		monthScheduleModel["day_of_week"] = []string{"Sunday"}
		monthScheduleModel["week_of_month"] = "First"
		monthScheduleModel["day_of_month"] = int(38)

		yearScheduleModel := make(map[string]interface{})
		yearScheduleModel["day_of_year"] = "First"

		model := make(map[string]interface{})
		model["unit"] = "Days"
		model["day_schedule"] = []map[string]interface{}{dayScheduleModel}
		model["week_schedule"] = []map[string]interface{}{weekScheduleModel}
		model["month_schedule"] = []map[string]interface{}{monthScheduleModel}
		model["year_schedule"] = []map[string]interface{}{yearScheduleModel}

		assert.Equal(t, result, model)
	}

	dayScheduleModel := new(backuprecoveryv1.DaySchedule)
	dayScheduleModel.Frequency = core.Int64Ptr(int64(1))

	weekScheduleModel := new(backuprecoveryv1.WeekSchedule)
	weekScheduleModel.DayOfWeek = []string{"Sunday"}

	monthScheduleModel := new(backuprecoveryv1.MonthSchedule)
	monthScheduleModel.DayOfWeek = []string{"Sunday"}
	monthScheduleModel.WeekOfMonth = core.StringPtr("First")
	monthScheduleModel.DayOfMonth = core.Int64Ptr(int64(38))

	yearScheduleModel := new(backuprecoveryv1.YearSchedule)
	yearScheduleModel.DayOfYear = core.StringPtr("First")

	model := new(backuprecoveryv1.BmrSchedule)
	model.Unit = core.StringPtr("Days")
	model.DaySchedule = dayScheduleModel
	model.WeekSchedule = weekScheduleModel
	model.MonthSchedule = monthScheduleModel
	model.YearSchedule = yearScheduleModel

	result, err := backuprecovery.DataSourceIbmBaasProtectionPolicyBmrScheduleToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasProtectionPolicyCdpBackupPolicyToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		dataLockConfigModel := make(map[string]interface{})
		dataLockConfigModel["mode"] = "Compliance"
		dataLockConfigModel["unit"] = "Days"
		dataLockConfigModel["duration"] = int(1)
		dataLockConfigModel["enable_worm_on_external_target"] = true

		cdpRetentionModel := make(map[string]interface{})
		cdpRetentionModel["unit"] = "Minutes"
		cdpRetentionModel["duration"] = int(1)
		cdpRetentionModel["data_lock_config"] = []map[string]interface{}{dataLockConfigModel}

		model := make(map[string]interface{})
		model["retention"] = []map[string]interface{}{cdpRetentionModel}

		assert.Equal(t, result, model)
	}

	dataLockConfigModel := new(backuprecoveryv1.DataLockConfig)
	dataLockConfigModel.Mode = core.StringPtr("Compliance")
	dataLockConfigModel.Unit = core.StringPtr("Days")
	dataLockConfigModel.Duration = core.Int64Ptr(int64(1))
	dataLockConfigModel.EnableWormOnExternalTarget = core.BoolPtr(true)

	cdpRetentionModel := new(backuprecoveryv1.CdpRetention)
	cdpRetentionModel.Unit = core.StringPtr("Minutes")
	cdpRetentionModel.Duration = core.Int64Ptr(int64(1))
	cdpRetentionModel.DataLockConfig = dataLockConfigModel

	model := new(backuprecoveryv1.CdpBackupPolicy)
	model.Retention = cdpRetentionModel

	result, err := backuprecovery.DataSourceIbmBaasProtectionPolicyCdpBackupPolicyToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasProtectionPolicyCdpRetentionToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		dataLockConfigModel := make(map[string]interface{})
		dataLockConfigModel["mode"] = "Compliance"
		dataLockConfigModel["unit"] = "Days"
		dataLockConfigModel["duration"] = int(1)
		dataLockConfigModel["enable_worm_on_external_target"] = true

		model := make(map[string]interface{})
		model["unit"] = "Minutes"
		model["duration"] = int(1)
		model["data_lock_config"] = []map[string]interface{}{dataLockConfigModel}

		assert.Equal(t, result, model)
	}

	dataLockConfigModel := new(backuprecoveryv1.DataLockConfig)
	dataLockConfigModel.Mode = core.StringPtr("Compliance")
	dataLockConfigModel.Unit = core.StringPtr("Days")
	dataLockConfigModel.Duration = core.Int64Ptr(int64(1))
	dataLockConfigModel.EnableWormOnExternalTarget = core.BoolPtr(true)

	model := new(backuprecoveryv1.CdpRetention)
	model.Unit = core.StringPtr("Minutes")
	model.Duration = core.Int64Ptr(int64(1))
	model.DataLockConfig = dataLockConfigModel

	result, err := backuprecovery.DataSourceIbmBaasProtectionPolicyCdpRetentionToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasProtectionPolicyStorageArraySnapshotBackupPolicyToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		minuteScheduleModel := make(map[string]interface{})
		minuteScheduleModel["frequency"] = int(1)

		hourScheduleModel := make(map[string]interface{})
		hourScheduleModel["frequency"] = int(1)

		dayScheduleModel := make(map[string]interface{})
		dayScheduleModel["frequency"] = int(1)

		weekScheduleModel := make(map[string]interface{})
		weekScheduleModel["day_of_week"] = []string{"Sunday"}

		monthScheduleModel := make(map[string]interface{})
		monthScheduleModel["day_of_week"] = []string{"Sunday"}
		monthScheduleModel["week_of_month"] = "First"
		monthScheduleModel["day_of_month"] = int(38)

		yearScheduleModel := make(map[string]interface{})
		yearScheduleModel["day_of_year"] = "First"

		storageArraySnapshotScheduleModel := make(map[string]interface{})
		storageArraySnapshotScheduleModel["unit"] = "Minutes"
		storageArraySnapshotScheduleModel["minute_schedule"] = []map[string]interface{}{minuteScheduleModel}
		storageArraySnapshotScheduleModel["hour_schedule"] = []map[string]interface{}{hourScheduleModel}
		storageArraySnapshotScheduleModel["day_schedule"] = []map[string]interface{}{dayScheduleModel}
		storageArraySnapshotScheduleModel["week_schedule"] = []map[string]interface{}{weekScheduleModel}
		storageArraySnapshotScheduleModel["month_schedule"] = []map[string]interface{}{monthScheduleModel}
		storageArraySnapshotScheduleModel["year_schedule"] = []map[string]interface{}{yearScheduleModel}

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
		model["schedule"] = []map[string]interface{}{storageArraySnapshotScheduleModel}
		model["retention"] = []map[string]interface{}{retentionModel}

		assert.Equal(t, result, model)
	}

	minuteScheduleModel := new(backuprecoveryv1.MinuteSchedule)
	minuteScheduleModel.Frequency = core.Int64Ptr(int64(1))

	hourScheduleModel := new(backuprecoveryv1.HourSchedule)
	hourScheduleModel.Frequency = core.Int64Ptr(int64(1))

	dayScheduleModel := new(backuprecoveryv1.DaySchedule)
	dayScheduleModel.Frequency = core.Int64Ptr(int64(1))

	weekScheduleModel := new(backuprecoveryv1.WeekSchedule)
	weekScheduleModel.DayOfWeek = []string{"Sunday"}

	monthScheduleModel := new(backuprecoveryv1.MonthSchedule)
	monthScheduleModel.DayOfWeek = []string{"Sunday"}
	monthScheduleModel.WeekOfMonth = core.StringPtr("First")
	monthScheduleModel.DayOfMonth = core.Int64Ptr(int64(38))

	yearScheduleModel := new(backuprecoveryv1.YearSchedule)
	yearScheduleModel.DayOfYear = core.StringPtr("First")

	storageArraySnapshotScheduleModel := new(backuprecoveryv1.StorageArraySnapshotSchedule)
	storageArraySnapshotScheduleModel.Unit = core.StringPtr("Minutes")
	storageArraySnapshotScheduleModel.MinuteSchedule = minuteScheduleModel
	storageArraySnapshotScheduleModel.HourSchedule = hourScheduleModel
	storageArraySnapshotScheduleModel.DaySchedule = dayScheduleModel
	storageArraySnapshotScheduleModel.WeekSchedule = weekScheduleModel
	storageArraySnapshotScheduleModel.MonthSchedule = monthScheduleModel
	storageArraySnapshotScheduleModel.YearSchedule = yearScheduleModel

	dataLockConfigModel := new(backuprecoveryv1.DataLockConfig)
	dataLockConfigModel.Mode = core.StringPtr("Compliance")
	dataLockConfigModel.Unit = core.StringPtr("Days")
	dataLockConfigModel.Duration = core.Int64Ptr(int64(1))
	dataLockConfigModel.EnableWormOnExternalTarget = core.BoolPtr(true)

	retentionModel := new(backuprecoveryv1.Retention)
	retentionModel.Unit = core.StringPtr("Days")
	retentionModel.Duration = core.Int64Ptr(int64(1))
	retentionModel.DataLockConfig = dataLockConfigModel

	model := new(backuprecoveryv1.StorageArraySnapshotBackupPolicy)
	model.Schedule = storageArraySnapshotScheduleModel
	model.Retention = retentionModel

	result, err := backuprecovery.DataSourceIbmBaasProtectionPolicyStorageArraySnapshotBackupPolicyToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasProtectionPolicyStorageArraySnapshotScheduleToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		minuteScheduleModel := make(map[string]interface{})
		minuteScheduleModel["frequency"] = int(1)

		hourScheduleModel := make(map[string]interface{})
		hourScheduleModel["frequency"] = int(1)

		dayScheduleModel := make(map[string]interface{})
		dayScheduleModel["frequency"] = int(1)

		weekScheduleModel := make(map[string]interface{})
		weekScheduleModel["day_of_week"] = []string{"Sunday"}

		monthScheduleModel := make(map[string]interface{})
		monthScheduleModel["day_of_week"] = []string{"Sunday"}
		monthScheduleModel["week_of_month"] = "First"
		monthScheduleModel["day_of_month"] = int(38)

		yearScheduleModel := make(map[string]interface{})
		yearScheduleModel["day_of_year"] = "First"

		model := make(map[string]interface{})
		model["unit"] = "Minutes"
		model["minute_schedule"] = []map[string]interface{}{minuteScheduleModel}
		model["hour_schedule"] = []map[string]interface{}{hourScheduleModel}
		model["day_schedule"] = []map[string]interface{}{dayScheduleModel}
		model["week_schedule"] = []map[string]interface{}{weekScheduleModel}
		model["month_schedule"] = []map[string]interface{}{monthScheduleModel}
		model["year_schedule"] = []map[string]interface{}{yearScheduleModel}

		assert.Equal(t, result, model)
	}

	minuteScheduleModel := new(backuprecoveryv1.MinuteSchedule)
	minuteScheduleModel.Frequency = core.Int64Ptr(int64(1))

	hourScheduleModel := new(backuprecoveryv1.HourSchedule)
	hourScheduleModel.Frequency = core.Int64Ptr(int64(1))

	dayScheduleModel := new(backuprecoveryv1.DaySchedule)
	dayScheduleModel.Frequency = core.Int64Ptr(int64(1))

	weekScheduleModel := new(backuprecoveryv1.WeekSchedule)
	weekScheduleModel.DayOfWeek = []string{"Sunday"}

	monthScheduleModel := new(backuprecoveryv1.MonthSchedule)
	monthScheduleModel.DayOfWeek = []string{"Sunday"}
	monthScheduleModel.WeekOfMonth = core.StringPtr("First")
	monthScheduleModel.DayOfMonth = core.Int64Ptr(int64(38))

	yearScheduleModel := new(backuprecoveryv1.YearSchedule)
	yearScheduleModel.DayOfYear = core.StringPtr("First")

	model := new(backuprecoveryv1.StorageArraySnapshotSchedule)
	model.Unit = core.StringPtr("Minutes")
	model.MinuteSchedule = minuteScheduleModel
	model.HourSchedule = hourScheduleModel
	model.DaySchedule = dayScheduleModel
	model.WeekSchedule = weekScheduleModel
	model.MonthSchedule = monthScheduleModel
	model.YearSchedule = yearScheduleModel

	result, err := backuprecovery.DataSourceIbmBaasProtectionPolicyStorageArraySnapshotScheduleToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasProtectionPolicyCancellationTimeoutParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["timeout_mins"] = int(26)
		model["backup_type"] = "kRegular"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.CancellationTimeoutParams)
	model.TimeoutMins = core.Int64Ptr(int64(26))
	model.BackupType = core.StringPtr("kRegular")

	result, err := backuprecovery.DataSourceIbmBaasProtectionPolicyCancellationTimeoutParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasProtectionPolicyBlackoutWindowToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		timeOfDayModel := make(map[string]interface{})
		timeOfDayModel["hour"] = int(0)
		timeOfDayModel["minute"] = int(0)
		timeOfDayModel["time_zone"] = "America/Los_Angeles"

		model := make(map[string]interface{})
		model["day"] = "Sunday"
		model["start_time"] = []map[string]interface{}{timeOfDayModel}
		model["end_time"] = []map[string]interface{}{timeOfDayModel}
		model["config_id"] = "testString"

		assert.Equal(t, result, model)
	}

	timeOfDayModel := new(backuprecoveryv1.TimeOfDay)
	timeOfDayModel.Hour = core.Int64Ptr(int64(0))
	timeOfDayModel.Minute = core.Int64Ptr(int64(0))
	timeOfDayModel.TimeZone = core.StringPtr("America/Los_Angeles")

	model := new(backuprecoveryv1.BlackoutWindow)
	model.Day = core.StringPtr("Sunday")
	model.StartTime = timeOfDayModel
	model.EndTime = timeOfDayModel
	model.ConfigID = core.StringPtr("testString")

	result, err := backuprecovery.DataSourceIbmBaasProtectionPolicyBlackoutWindowToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasProtectionPolicyTimeOfDayToMap(t *testing.T) {
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

	result, err := backuprecovery.DataSourceIbmBaasProtectionPolicyTimeOfDayToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasProtectionPolicyExtendedRetentionPolicyToMap(t *testing.T) {
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

	result, err := backuprecovery.DataSourceIbmBaasProtectionPolicyExtendedRetentionPolicyToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasProtectionPolicyExtendedRetentionScheduleToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["unit"] = "Runs"
		model["frequency"] = int(1)

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.ExtendedRetentionSchedule)
	model.Unit = core.StringPtr("Runs")
	model.Frequency = core.Int64Ptr(int64(1))

	result, err := backuprecovery.DataSourceIbmBaasProtectionPolicyExtendedRetentionScheduleToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasProtectionPolicyTargetsConfigurationToMap(t *testing.T) {
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

	result, err := backuprecovery.DataSourceIbmBaasProtectionPolicyTargetsConfigurationToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasProtectionPolicyReplicationTargetConfigurationToMap(t *testing.T) {
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

	result, err := backuprecovery.DataSourceIbmBaasProtectionPolicyReplicationTargetConfigurationToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasProtectionPolicyTargetScheduleToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["unit"] = "Runs"
		model["frequency"] = int(1)

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.TargetSchedule)
	model.Unit = core.StringPtr("Runs")
	model.Frequency = core.Int64Ptr(int64(1))

	result, err := backuprecovery.DataSourceIbmBaasProtectionPolicyTargetScheduleToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasProtectionPolicyLogRetentionToMap(t *testing.T) {
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

	result, err := backuprecovery.DataSourceIbmBaasProtectionPolicyLogRetentionToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasProtectionPolicyAWSTargetConfigToMap(t *testing.T) {
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

	result, err := backuprecovery.DataSourceIbmBaasProtectionPolicyAWSTargetConfigToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasProtectionPolicyAzureTargetConfigToMap(t *testing.T) {
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

	result, err := backuprecovery.DataSourceIbmBaasProtectionPolicyAzureTargetConfigToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasProtectionPolicyRemoteTargetConfigToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["cluster_id"] = int(26)
		model["cluster_name"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.RemoteTargetConfig)
	model.ClusterID = core.Int64Ptr(int64(26))
	model.ClusterName = core.StringPtr("testString")

	result, err := backuprecovery.DataSourceIbmBaasProtectionPolicyRemoteTargetConfigToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasProtectionPolicyArchivalTargetConfigurationToMap(t *testing.T) {
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

	result, err := backuprecovery.DataSourceIbmBaasProtectionPolicyArchivalTargetConfigurationToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasProtectionPolicyCloudSpinTargetConfigurationToMap(t *testing.T) {
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

	result, err := backuprecovery.DataSourceIbmBaasProtectionPolicyCloudSpinTargetConfigurationToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasProtectionPolicyCloudSpinTargetToMap(t *testing.T) {
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

	result, err := backuprecovery.DataSourceIbmBaasProtectionPolicyCloudSpinTargetToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasProtectionPolicyAwsCloudSpinParamsToMap(t *testing.T) {
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

	result, err := backuprecovery.DataSourceIbmBaasProtectionPolicyAwsCloudSpinParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasProtectionPolicyCustomTagParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["key"] = "testString"
		model["value"] = "testString"

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.CustomTagParams)
	model.Key = core.StringPtr("testString")
	model.Value = core.StringPtr("testString")

	result, err := backuprecovery.DataSourceIbmBaasProtectionPolicyCustomTagParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasProtectionPolicyAzureCloudSpinParamsToMap(t *testing.T) {
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

	result, err := backuprecovery.DataSourceIbmBaasProtectionPolicyAzureCloudSpinParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasProtectionPolicyOnpremDeployTargetConfigurationToMap(t *testing.T) {
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

	result, err := backuprecovery.DataSourceIbmBaasProtectionPolicyOnpremDeployTargetConfigurationToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasProtectionPolicyOnpremDeployParamsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["id"] = int(26)

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.OnpremDeployParams)
	model.ID = core.Int64Ptr(int64(26))

	result, err := backuprecovery.DataSourceIbmBaasProtectionPolicyOnpremDeployParamsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasProtectionPolicyRpaasTargetConfigurationToMap(t *testing.T) {
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

	result, err := backuprecovery.DataSourceIbmBaasProtectionPolicyRpaasTargetConfigurationToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasProtectionPolicyCascadedTargetConfigurationToMap(t *testing.T) {
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

		targetsConfigurationModel := make(map[string]interface{})
		targetsConfigurationModel["replication_targets"] = []map[string]interface{}{replicationTargetConfigurationModel}
		targetsConfigurationModel["archival_targets"] = []map[string]interface{}{archivalTargetConfigurationModel}
		targetsConfigurationModel["cloud_spin_targets"] = []map[string]interface{}{cloudSpinTargetConfigurationModel}
		targetsConfigurationModel["onprem_deploy_targets"] = []map[string]interface{}{onpremDeployTargetConfigurationModel}
		targetsConfigurationModel["rpaas_targets"] = []map[string]interface{}{rpaasTargetConfigurationModel}

		model := make(map[string]interface{})
		model["source_cluster_id"] = int(26)
		model["remote_targets"] = []map[string]interface{}{targetsConfigurationModel}

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

	targetsConfigurationModel := new(backuprecoveryv1.TargetsConfiguration)
	targetsConfigurationModel.ReplicationTargets = []backuprecoveryv1.ReplicationTargetConfiguration{*replicationTargetConfigurationModel}
	targetsConfigurationModel.ArchivalTargets = []backuprecoveryv1.ArchivalTargetConfiguration{*archivalTargetConfigurationModel}
	targetsConfigurationModel.CloudSpinTargets = []backuprecoveryv1.CloudSpinTargetConfiguration{*cloudSpinTargetConfigurationModel}
	targetsConfigurationModel.OnpremDeployTargets = []backuprecoveryv1.OnpremDeployTargetConfiguration{*onpremDeployTargetConfigurationModel}
	targetsConfigurationModel.RpaasTargets = []backuprecoveryv1.RpaasTargetConfiguration{*rpaasTargetConfigurationModel}

	model := new(backuprecoveryv1.CascadedTargetConfiguration)
	model.SourceClusterID = core.Int64Ptr(int64(26))
	model.RemoteTargets = targetsConfigurationModel

	result, err := backuprecovery.DataSourceIbmBaasProtectionPolicyCascadedTargetConfigurationToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}

func TestDataSourceIbmBaasProtectionPolicyRetryOptionsToMap(t *testing.T) {
	checkResult := func(result map[string]interface{}) {
		model := make(map[string]interface{})
		model["retries"] = int(0)
		model["retry_interval_mins"] = int(1)

		assert.Equal(t, result, model)
	}

	model := new(backuprecoveryv1.RetryOptions)
	model.Retries = core.Int64Ptr(int64(0))
	model.RetryIntervalMins = core.Int64Ptr(int64(1))

	result, err := backuprecovery.DataSourceIbmBaasProtectionPolicyRetryOptionsToMap(model)
	assert.Nil(t, err)
	checkResult(result)
}
