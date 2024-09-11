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
