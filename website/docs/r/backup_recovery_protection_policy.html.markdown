---
layout: "ibm"
page_title: "IBM : ibm_backup_recovery_protection_policy"
description: |-
  Manages backup_recovery_protection_policy.
subcategory: "IBM Backup Recovery"
---

# ibm_backup_recovery_protection_policy

Create, update, and delete backup_recovery_protection_policys with this resource.

## Example Usage

```hcl
resource "ibm_backup_recovery_protection_policy" "backup_recovery_protection_policy_instance" {
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
  name = "name"
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
  retry_options {
		retries = 0
		retry_interval_mins = 1
  }
  x_ibm_tenant_id = "x_ibm_tenant_id"
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `backup_policy` - (Required, List) Specifies the backup schedule and retentions of a Protection Policy.
Nested schema for **backup_policy**:
	* `bmr` - (Optional, List) Specifies the BMR schedule in case of physical source protection.
	Nested schema for **bmr**:
		* `retention` - (Required, List) Specifies the retention of a backup.
		Nested schema for **retention**:
			* `data_lock_config` - (Optional, List) Specifies WORM retention type for the snapshots. When a WORM retention type is specified, the snapshots of the Protection Groups using this policy will be kept for the last N days as specified in the duration of the datalock. During that time, the snapshots cannot be deleted.
			Nested schema for **data_lock_config**:
				* `duration` - (Required, Integer) Specifies the duration for a dataLock. <br> Example. If duration is 7 and unit is Months, the dataLock is enabled for last 7 * 30 = 210 days of the backup.
				  * Constraints: The minimum value is `1`.
				* `enable_worm_on_external_target` - (Optional, Boolean) Specifies whether objects in the external target associated with this policy need to be made immutable.
				* `mode` - (Required, String) Specifies the type of WORM retention type. 'Compliance' implies WORM retention is set for compliance reason. 'Administrative' implies WORM retention is set for administrative purposes.
				  * Constraints: Allowable values are: `Compliance`, `Administrative`.
				* `unit` - (Required, String) Specificies the Retention Unit of a dataLock measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Months' then number of retention days will be 365 * 2 = 730 days.
				  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
			* `duration` - (Required, Integer) Specifies the duration for a backup retention. <br> Example. If duration is 7 and unit is Months, the retention of a backup is 7 * 30 = 210 days.
			  * Constraints: The minimum value is `1`.
			* `unit` - (Required, String) Specificies the Retention Unit of a backup measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Years' then number of retention days will be 365 * 2 = 730 days.
			  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
		* `schedule` - (Required, List) Specifies settings that defines how frequent bmr backup will be performed for a Protection Group.
		Nested schema for **schedule**:
			* `day_schedule` - (Optional, List) Specifies settings that define a schedule for a Protection Group runs to start after certain number of days.
			Nested schema for **day_schedule**:
				* `frequency` - (Required, Integer) Specifies a factor to multiply the unit by, to determine the backup schedule. <br> Example: If 'frequency' set to 2 and the unit is 'Hours', then Snapshots are backed up every 2 hours. <br> This field is only applicable if unit is 'Minutes', 'Hours' or 'Days'.
				  * Constraints: The minimum value is `1`.
			* `month_schedule` - (Optional, List) Specifies settings that define a schedule for a Protection Group runs to on specific week and specific days of that week.
			Nested schema for **month_schedule**:
				* `day_of_month` - (Optional, Integer) Specifies the exact date of the month (such as 18) in a Monthly Schedule specified by unit field as 'Years'. <br> Example: if 'dayOfMonth' is set to '18', a backup is performed on the 18th of every month.
				* `day_of_week` - (Optional, List) Specifies a list of days of the week when to start Protection Group Runs. <br> Example: To run a Protection Group on every Monday and Tuesday, set the schedule with following values: <br>  unit: 'Weeks' <br>  dayOfWeek: ['Monday','Tuesday'].
				  * Constraints: Allowable list items are: `Sunday`, `Monday`, `Tuesday`, `Wednesday`, `Thursday`, `Friday`, `Saturday`.
				* `week_of_month` - (Optional, String) Specifies the week of the month (such as 'Third') or nth day of month (such as 'First' or 'Last') in a Monthly Schedule specified by unit field as 'Months'. <br>This field can be used in combination with 'dayOfWeek' to define the day in the month to start the Protection Group Run. <br> Example: if 'weekOfMonth' is set to 'Third' and day is set to 'Monday', a backup is performed on the third Monday of every month. <br> Example: if 'weekOfMonth' is set to 'Last' and dayOfWeek is not set, a backup is performed on the last day of every month.
				  * Constraints: Allowable values are: `First`, `Second`, `Third`, `Fourth`, `Last`.
			* `unit` - (Required, String) Specifies how often to start new runs of a Protection Group. <br>'Weeks' specifies that new Protection Group runs start weekly on certain days specified using 'dayOfWeek' field. <br>'Months' specifies that new Protection Group runs start monthly on certain day of specific week.
			  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
			* `week_schedule` - (Optional, List) Specifies settings that define a schedule for a Protection Group runs to start on certain days of week.
			Nested schema for **week_schedule**:
				* `day_of_week` - (Required, List) Specifies a list of days of the week when to start Protection Group Runs. <br> Example: To run a Protection Group on every Monday and Tuesday, set the schedule with following values: <br>  unit: 'Weeks' <br>  dayOfWeek: ['Monday','Tuesday'].
				  * Constraints: Allowable list items are: `Sunday`, `Monday`, `Tuesday`, `Wednesday`, `Thursday`, `Friday`, `Saturday`.
			* `year_schedule` - (Optional, List) Specifies settings that define a schedule for a Protection Group to run on specific year and specific day of that year.
			Nested schema for **year_schedule**:
				* `day_of_year` - (Required, String) Specifies the day of the Year (such as 'First' or 'Last') in a Yearly Schedule. <br>This field is used to define the day in the year to start the Protection Group Run. <br> Example: if 'dayOfYear' is set to 'First', a backup is performed on the first day of every year. <br> Example: if 'dayOfYear' is set to 'Last', a backup is performed on the last day of every year.
				  * Constraints: Allowable values are: `First`, `Last`.
	* `cdp` - (Optional, List) Specifies CDP (Continious Data Protection) backup settings for a Protection Group.
	Nested schema for **cdp**:
		* `retention` - (Required, List) Specifies the retention of a CDP backup.
		Nested schema for **retention**:
			* `data_lock_config` - (Optional, List) Specifies WORM retention type for the snapshots. When a WORM retention type is specified, the snapshots of the Protection Groups using this policy will be kept for the last N days as specified in the duration of the datalock. During that time, the snapshots cannot be deleted.
			Nested schema for **data_lock_config**:
				* `duration` - (Required, Integer) Specifies the duration for a dataLock. <br> Example. If duration is 7 and unit is Months, the dataLock is enabled for last 7 * 30 = 210 days of the backup.
				  * Constraints: The minimum value is `1`.
				* `enable_worm_on_external_target` - (Optional, Boolean) Specifies whether objects in the external target associated with this policy need to be made immutable.
				* `mode` - (Required, String) Specifies the type of WORM retention type. 'Compliance' implies WORM retention is set for compliance reason. 'Administrative' implies WORM retention is set for administrative purposes.
				  * Constraints: Allowable values are: `Compliance`, `Administrative`.
				* `unit` - (Required, String) Specificies the Retention Unit of a dataLock measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Months' then number of retention days will be 365 * 2 = 730 days.
				  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
			* `duration` - (Required, Integer) Specifies the duration for a cdp backup retention.
			  * Constraints: The minimum value is `1`.
			* `unit` - (Required, String) Specificies the Retention Unit of a CDP backup measured in minutes or hours.
			  * Constraints: Allowable values are: `Minutes`, `Hours`.
	* `log` - (Optional, List) Specifies log backup settings for a Protection Group.
	Nested schema for **log**:
		* `retention` - (Required, List) Specifies the retention of a backup.
		Nested schema for **retention**:
			* `data_lock_config` - (Optional, List) Specifies WORM retention type for the snapshots. When a WORM retention type is specified, the snapshots of the Protection Groups using this policy will be kept for the last N days as specified in the duration of the datalock. During that time, the snapshots cannot be deleted.
			Nested schema for **data_lock_config**:
				* `duration` - (Required, Integer) Specifies the duration for a dataLock. <br> Example. If duration is 7 and unit is Months, the dataLock is enabled for last 7 * 30 = 210 days of the backup.
				  * Constraints: The minimum value is `1`.
				* `enable_worm_on_external_target` - (Optional, Boolean) Specifies whether objects in the external target associated with this policy need to be made immutable.
				* `mode` - (Required, String) Specifies the type of WORM retention type. 'Compliance' implies WORM retention is set for compliance reason. 'Administrative' implies WORM retention is set for administrative purposes.
				  * Constraints: Allowable values are: `Compliance`, `Administrative`.
				* `unit` - (Required, String) Specificies the Retention Unit of a dataLock measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Months' then number of retention days will be 365 * 2 = 730 days.
				  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
			* `duration` - (Required, Integer) Specifies the duration for a backup retention. <br> Example. If duration is 7 and unit is Months, the retention of a backup is 7 * 30 = 210 days.
			  * Constraints: The minimum value is `1`.
			* `unit` - (Required, String) Specificies the Retention Unit of a backup measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Years' then number of retention days will be 365 * 2 = 730 days.
			  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
		* `schedule` - (Required, List) Specifies settings that defines how frequent log backup will be performed for a Protection Group.
		Nested schema for **schedule**:
			* `hour_schedule` - (Optional, List) Specifies settings that define a schedule for a Protection Group runs to start after certain number of hours.
			Nested schema for **hour_schedule**:
				* `frequency` - (Required, Integer) Specifies a factor to multiply the unit by, to determine the backup schedule. <br> Example: If 'frequency' set to 2 and the unit is 'Hours', then Snapshots are backed up every 2 hours. <br> This field is only applicable if unit is 'Minutes', 'Hours' or 'Days'.
				  * Constraints: The minimum value is `1`.
			* `minute_schedule` - (Optional, List) Specifies settings that define a schedule for a Protection Group runs to start after certain number of minutes.
			Nested schema for **minute_schedule**:
				* `frequency` - (Required, Integer) Specifies a factor to multiply the unit by, to determine the backup schedule. <br> Example: If 'frequency' set to 2 and the unit is 'Hours', then Snapshots are backed up every 2 hours. <br> This field is only applicable if unit is 'Minutes', 'Hours' or 'Days'.
				  * Constraints: The minimum value is `1`.
			* `unit` - (Required, String) Specifies how often to start new Protection Group Runs of a Protection Group. <br>'Minutes' specifies that Protection Group run starts periodically after certain number of minutes specified in 'frequency' field. <br>'Hours' specifies that Protection Group run starts periodically after certain number of hours specified in 'frequency' field.
			  * Constraints: Allowable values are: `Minutes`, `Hours`.
	* `regular` - (Required, List) Specifies the Incremental and Full policy settings and also the common Retention policy settings.".
	Nested schema for **regular**:
		* `full` - (Optional, List) Specifies full backup settings for a Protection Group. Currently, full backup settings can be specified by using either of 'schedule' or 'schdulesAndRetentions' field. Using 'schdulesAndRetentions' is recommended when multiple full backups need to be configured. If full and incremental backup has common retention then only setting 'schedule' is recommended.
		Nested schema for **full**:
			* `schedule` - (Optional, List) Specifies settings that defines how frequent full backup will be performed for a Protection Group.
			Nested schema for **schedule**:
				* `day_schedule` - (Optional, List) Specifies settings that define a schedule for a Protection Group runs to start after certain number of days.
				Nested schema for **day_schedule**:
					* `frequency` - (Required, Integer) Specifies a factor to multiply the unit by, to determine the backup schedule. <br> Example: If 'frequency' set to 2 and the unit is 'Hours', then Snapshots are backed up every 2 hours. <br> This field is only applicable if unit is 'Minutes', 'Hours' or 'Days'.
					  * Constraints: The minimum value is `1`.
				* `month_schedule` - (Optional, List) Specifies settings that define a schedule for a Protection Group runs to on specific week and specific days of that week.
				Nested schema for **month_schedule**:
					* `day_of_month` - (Optional, Integer) Specifies the exact date of the month (such as 18) in a Monthly Schedule specified by unit field as 'Years'. <br> Example: if 'dayOfMonth' is set to '18', a backup is performed on the 18th of every month.
					* `day_of_week` - (Optional, List) Specifies a list of days of the week when to start Protection Group Runs. <br> Example: To run a Protection Group on every Monday and Tuesday, set the schedule with following values: <br>  unit: 'Weeks' <br>  dayOfWeek: ['Monday','Tuesday'].
					  * Constraints: Allowable list items are: `Sunday`, `Monday`, `Tuesday`, `Wednesday`, `Thursday`, `Friday`, `Saturday`.
					* `week_of_month` - (Optional, String) Specifies the week of the month (such as 'Third') or nth day of month (such as 'First' or 'Last') in a Monthly Schedule specified by unit field as 'Months'. <br>This field can be used in combination with 'dayOfWeek' to define the day in the month to start the Protection Group Run. <br> Example: if 'weekOfMonth' is set to 'Third' and day is set to 'Monday', a backup is performed on the third Monday of every month. <br> Example: if 'weekOfMonth' is set to 'Last' and dayOfWeek is not set, a backup is performed on the last day of every month.
					  * Constraints: Allowable values are: `First`, `Second`, `Third`, `Fourth`, `Last`.
				* `unit` - (Required, String) Specifies how often to start new runs of a Protection Group. <br>'Days' specifies that Protection Group run starts periodically on every day. For full backup schedule, currently we only support frequecny of 1 which indicates that full backup will be performed daily. <br>'Weeks' specifies that new Protection Group runs start weekly on certain days specified using 'dayOfWeek' field. <br>'Months' specifies that new Protection Group runs start monthly on certain day of specific week. This schedule needs 'weekOfMonth' and 'dayOfWeek' fields to be set. <br>'ProtectOnce' specifies that groups using this policy option will run only once and after that group will permanently be disabled. <br> Example: To run the Protection Group on Second Sunday of Every Month, following schedule need to be set: <br> unit: 'Month' <br> dayOfWeek: 'Sunday' <br> weekOfMonth: 'Second'.
				  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`, `ProtectOnce`.
				* `week_schedule` - (Optional, List) Specifies settings that define a schedule for a Protection Group runs to start on certain days of week.
				Nested schema for **week_schedule**:
					* `day_of_week` - (Required, List) Specifies a list of days of the week when to start Protection Group Runs. <br> Example: To run a Protection Group on every Monday and Tuesday, set the schedule with following values: <br>  unit: 'Weeks' <br>  dayOfWeek: ['Monday','Tuesday'].
					  * Constraints: Allowable list items are: `Sunday`, `Monday`, `Tuesday`, `Wednesday`, `Thursday`, `Friday`, `Saturday`.
				* `year_schedule` - (Optional, List) Specifies settings that define a schedule for a Protection Group to run on specific year and specific day of that year.
				Nested schema for **year_schedule**:
					* `day_of_year` - (Required, String) Specifies the day of the Year (such as 'First' or 'Last') in a Yearly Schedule. <br>This field is used to define the day in the year to start the Protection Group Run. <br> Example: if 'dayOfYear' is set to 'First', a backup is performed on the first day of every year. <br> Example: if 'dayOfYear' is set to 'Last', a backup is performed on the last day of every year.
					  * Constraints: Allowable values are: `First`, `Last`.
		* `full_backups` - (Optional, List) Specifies multiple schedules and retentions for full backup. Specify either of the 'full' or 'fullBackups' values. Its recommended to use 'fullBaackups' value since 'full' will be deprecated after few releases.
		Nested schema for **full_backups**:
			* `retention` - (Required, List) Specifies the retention of a backup.
			Nested schema for **retention**:
				* `data_lock_config` - (Optional, List) Specifies WORM retention type for the snapshots. When a WORM retention type is specified, the snapshots of the Protection Groups using this policy will be kept for the last N days as specified in the duration of the datalock. During that time, the snapshots cannot be deleted.
				Nested schema for **data_lock_config**:
					* `duration` - (Required, Integer) Specifies the duration for a dataLock. <br> Example. If duration is 7 and unit is Months, the dataLock is enabled for last 7 * 30 = 210 days of the backup.
					  * Constraints: The minimum value is `1`.
					* `enable_worm_on_external_target` - (Optional, Boolean) Specifies whether objects in the external target associated with this policy need to be made immutable.
					* `mode` - (Required, String) Specifies the type of WORM retention type. 'Compliance' implies WORM retention is set for compliance reason. 'Administrative' implies WORM retention is set for administrative purposes.
					  * Constraints: Allowable values are: `Compliance`, `Administrative`.
					* `unit` - (Required, String) Specificies the Retention Unit of a dataLock measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Months' then number of retention days will be 365 * 2 = 730 days.
					  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
				* `duration` - (Required, Integer) Specifies the duration for a backup retention. <br> Example. If duration is 7 and unit is Months, the retention of a backup is 7 * 30 = 210 days.
				  * Constraints: The minimum value is `1`.
				* `unit` - (Required, String) Specificies the Retention Unit of a backup measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Years' then number of retention days will be 365 * 2 = 730 days.
				  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
			* `schedule` - (Required, List) Specifies settings that defines how frequent full backup will be performed for a Protection Group.
			Nested schema for **schedule**:
				* `day_schedule` - (Optional, List) Specifies settings that define a schedule for a Protection Group runs to start after certain number of days.
				Nested schema for **day_schedule**:
					* `frequency` - (Required, Integer) Specifies a factor to multiply the unit by, to determine the backup schedule. <br> Example: If 'frequency' set to 2 and the unit is 'Hours', then Snapshots are backed up every 2 hours. <br> This field is only applicable if unit is 'Minutes', 'Hours' or 'Days'.
					  * Constraints: The minimum value is `1`.
				* `month_schedule` - (Optional, List) Specifies settings that define a schedule for a Protection Group runs to on specific week and specific days of that week.
				Nested schema for **month_schedule**:
					* `day_of_month` - (Optional, Integer) Specifies the exact date of the month (such as 18) in a Monthly Schedule specified by unit field as 'Years'. <br> Example: if 'dayOfMonth' is set to '18', a backup is performed on the 18th of every month.
					* `day_of_week` - (Optional, List) Specifies a list of days of the week when to start Protection Group Runs. <br> Example: To run a Protection Group on every Monday and Tuesday, set the schedule with following values: <br>  unit: 'Weeks' <br>  dayOfWeek: ['Monday','Tuesday'].
					  * Constraints: Allowable list items are: `Sunday`, `Monday`, `Tuesday`, `Wednesday`, `Thursday`, `Friday`, `Saturday`.
					* `week_of_month` - (Optional, String) Specifies the week of the month (such as 'Third') or nth day of month (such as 'First' or 'Last') in a Monthly Schedule specified by unit field as 'Months'. <br>This field can be used in combination with 'dayOfWeek' to define the day in the month to start the Protection Group Run. <br> Example: if 'weekOfMonth' is set to 'Third' and day is set to 'Monday', a backup is performed on the third Monday of every month. <br> Example: if 'weekOfMonth' is set to 'Last' and dayOfWeek is not set, a backup is performed on the last day of every month.
					  * Constraints: Allowable values are: `First`, `Second`, `Third`, `Fourth`, `Last`.
				* `unit` - (Required, String) Specifies how often to start new runs of a Protection Group. <br>'Days' specifies that Protection Group run starts periodically on every day. For full backup schedule, currently we only support frequecny of 1 which indicates that full backup will be performed daily. <br>'Weeks' specifies that new Protection Group runs start weekly on certain days specified using 'dayOfWeek' field. <br>'Months' specifies that new Protection Group runs start monthly on certain day of specific week. This schedule needs 'weekOfMonth' and 'dayOfWeek' fields to be set. <br>'ProtectOnce' specifies that groups using this policy option will run only once and after that group will permanently be disabled. <br> Example: To run the Protection Group on Second Sunday of Every Month, following schedule need to be set: <br> unit: 'Month' <br> dayOfWeek: 'Sunday' <br> weekOfMonth: 'Second'.
				  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`, `ProtectOnce`.
				* `week_schedule` - (Optional, List) Specifies settings that define a schedule for a Protection Group runs to start on certain days of week.
				Nested schema for **week_schedule**:
					* `day_of_week` - (Required, List) Specifies a list of days of the week when to start Protection Group Runs. <br> Example: To run a Protection Group on every Monday and Tuesday, set the schedule with following values: <br>  unit: 'Weeks' <br>  dayOfWeek: ['Monday','Tuesday'].
					  * Constraints: Allowable list items are: `Sunday`, `Monday`, `Tuesday`, `Wednesday`, `Thursday`, `Friday`, `Saturday`.
				* `year_schedule` - (Optional, List) Specifies settings that define a schedule for a Protection Group to run on specific year and specific day of that year.
				Nested schema for **year_schedule**:
					* `day_of_year` - (Required, String) Specifies the day of the Year (such as 'First' or 'Last') in a Yearly Schedule. <br>This field is used to define the day in the year to start the Protection Group Run. <br> Example: if 'dayOfYear' is set to 'First', a backup is performed on the first day of every year. <br> Example: if 'dayOfYear' is set to 'Last', a backup is performed on the last day of every year.
					  * Constraints: Allowable values are: `First`, `Last`.
		* `incremental` - (Optional, List) Specifies incremental backup settings for a Protection Group.
		Nested schema for **incremental**:
			* `schedule` - (Required, List) Specifies settings that defines how frequent backup will be performed for a Protection Group.
			Nested schema for **schedule**:
				* `day_schedule` - (Optional, List) Specifies settings that define a schedule for a Protection Group runs to start after certain number of days.
				Nested schema for **day_schedule**:
					* `frequency` - (Required, Integer) Specifies a factor to multiply the unit by, to determine the backup schedule. <br> Example: If 'frequency' set to 2 and the unit is 'Hours', then Snapshots are backed up every 2 hours. <br> This field is only applicable if unit is 'Minutes', 'Hours' or 'Days'.
					  * Constraints: The minimum value is `1`.
				* `hour_schedule` - (Optional, List) Specifies settings that define a schedule for a Protection Group runs to start after certain number of hours.
				Nested schema for **hour_schedule**:
					* `frequency` - (Required, Integer) Specifies a factor to multiply the unit by, to determine the backup schedule. <br> Example: If 'frequency' set to 2 and the unit is 'Hours', then Snapshots are backed up every 2 hours. <br> This field is only applicable if unit is 'Minutes', 'Hours' or 'Days'.
					  * Constraints: The minimum value is `1`.
				* `minute_schedule` - (Optional, List) Specifies settings that define a schedule for a Protection Group runs to start after certain number of minutes.
				Nested schema for **minute_schedule**:
					* `frequency` - (Required, Integer) Specifies a factor to multiply the unit by, to determine the backup schedule. <br> Example: If 'frequency' set to 2 and the unit is 'Hours', then Snapshots are backed up every 2 hours. <br> This field is only applicable if unit is 'Minutes', 'Hours' or 'Days'.
					  * Constraints: The minimum value is `1`.
				* `month_schedule` - (Optional, List) Specifies settings that define a schedule for a Protection Group runs to on specific week and specific days of that week.
				Nested schema for **month_schedule**:
					* `day_of_month` - (Optional, Integer) Specifies the exact date of the month (such as 18) in a Monthly Schedule specified by unit field as 'Years'. <br> Example: if 'dayOfMonth' is set to '18', a backup is performed on the 18th of every month.
					* `day_of_week` - (Optional, List) Specifies a list of days of the week when to start Protection Group Runs. <br> Example: To run a Protection Group on every Monday and Tuesday, set the schedule with following values: <br>  unit: 'Weeks' <br>  dayOfWeek: ['Monday','Tuesday'].
					  * Constraints: Allowable list items are: `Sunday`, `Monday`, `Tuesday`, `Wednesday`, `Thursday`, `Friday`, `Saturday`.
					* `week_of_month` - (Optional, String) Specifies the week of the month (such as 'Third') or nth day of month (such as 'First' or 'Last') in a Monthly Schedule specified by unit field as 'Months'. <br>This field can be used in combination with 'dayOfWeek' to define the day in the month to start the Protection Group Run. <br> Example: if 'weekOfMonth' is set to 'Third' and day is set to 'Monday', a backup is performed on the third Monday of every month. <br> Example: if 'weekOfMonth' is set to 'Last' and dayOfWeek is not set, a backup is performed on the last day of every month.
					  * Constraints: Allowable values are: `First`, `Second`, `Third`, `Fourth`, `Last`.
				* `unit` - (Required, String) Specifies how often to start new runs of a Protection Group. <br>'Minutes' specifies that Protection Group run starts periodically after certain number of minutes specified in 'frequency' field. <br>'Hours' specifies that Protection Group run starts periodically after certain number of hours specified in 'frequency' field. <br>'Days' specifies that Protection Group run starts periodically after certain number of days specified in 'frequency' field. <br>'Week' specifies that new Protection Group runs start weekly on certain days specified using 'dayOfWeek' field. <br>'Month' specifies that new Protection Group runs start monthly on certain day of specific week. This schedule needs 'weekOfMonth' and 'dayOfWeek' fields to be set. <br> Example: To run the Protection Group on Second Sunday of Every Month, following schedule need to be set: <br> unit: 'Month' <br> dayOfWeek: 'Sunday' <br> weekOfMonth: 'Second'.
				  * Constraints: Allowable values are: `Minutes`, `Hours`, `Days`, `Weeks`, `Months`, `Years`.
				* `week_schedule` - (Optional, List) Specifies settings that define a schedule for a Protection Group runs to start on certain days of week.
				Nested schema for **week_schedule**:
					* `day_of_week` - (Required, List) Specifies a list of days of the week when to start Protection Group Runs. <br> Example: To run a Protection Group on every Monday and Tuesday, set the schedule with following values: <br>  unit: 'Weeks' <br>  dayOfWeek: ['Monday','Tuesday'].
					  * Constraints: Allowable list items are: `Sunday`, `Monday`, `Tuesday`, `Wednesday`, `Thursday`, `Friday`, `Saturday`.
				* `year_schedule` - (Optional, List) Specifies settings that define a schedule for a Protection Group to run on specific year and specific day of that year.
				Nested schema for **year_schedule**:
					* `day_of_year` - (Required, String) Specifies the day of the Year (such as 'First' or 'Last') in a Yearly Schedule. <br>This field is used to define the day in the year to start the Protection Group Run. <br> Example: if 'dayOfYear' is set to 'First', a backup is performed on the first day of every year. <br> Example: if 'dayOfYear' is set to 'Last', a backup is performed on the last day of every year.
					  * Constraints: Allowable values are: `First`, `Last`.
		* `primary_backup_target` - (Optional, List) Specifies the primary backup target settings for regular backups. If the backup target field is not specified then backup will be taken locally on the Cohesity cluster.
		Nested schema for **primary_backup_target**:
			* `archival_target_settings` - (Optional, List) Specifies the primary archival settings. Mainly used for cloud direct archive (CAD) policy where primary backup is stored on archival target.
			Nested schema for **archival_target_settings**:
				* `target_id` - (Required, Integer) Specifies the Archival target id to take primary backup.
				* `target_name` - (Computed, String) Specifies the Archival target name where Snapshots are copied.
				* `tier_settings` - (Optional, List) Specifies the settings tier levels configured with each archival target. The tier settings need to be applied in specific order and default tier should always be passed as first entry in tiers array. The following example illustrates how to configure tiering input for AWS tiering. Same type of input structure applied to other cloud platforms also. <br>If user wants to achieve following tiering for backup, <br>User Desired Tiering- <br><t>1.Archive Full back up for 12 Months <br><t>2.Tier Levels <br><t><t>[1,12] [ <br><t><t><t>s3 (1 to 2 months), (default tier) <br><t><t><t>s3 Intelligent tiering (3 to 6 months), <br><t><t><t>s3 One Zone (7 to 9 months) <br><t><t><t>Glacier (10 to 12 months)] <br><t>API Input <br><t><t>1.tiers-[ <br><t><t><t>{'tierType': 'S3','moveAfterUnit':'months', <br><t><t><t>'moveAfter':2 - move from s3 to s3Inte after 2 months}, <br><t><t><t>{'tierType': 'S3Inte','moveAfterUnit':'months', <br><t><t><t>'moveAfter':4 - move from S3Inte to Glacier after 4 months}, <br><t><t><t>{'tierType': 'Glacier', 'moveAfterUnit':'months', <br><t><t><t>'moveAfter': 3 - move from Glacier to S3 One Zone after 3 months }, <br><t><t><t>{'tierType': 'S3 One Zone', 'moveAfterUnit': nil, <br><t><t><t>'moveAfter': nil - For the last record, 'moveAfter' and 'moveAfterUnit' <br><t><t><t>will be ignored since there are no further tier for data movement } <br><t><t><t>}].
				Nested schema for **tier_settings**:
					* `aws_tiering` - (Optional, List) Specifies aws tiers.
					Nested schema for **aws_tiering**:
						* `tiers` - (Required, List) Specifies the tiers that are used to move the archived backup from current tier to next tier. The order of the tiers determines which tier will be used next for moving the archived backup. The first tier input should always be default tier where backup will be acrhived. Each tier specifies how much time after the backup will be moved to next tier from the current tier.
						Nested schema for **tiers**:
							* `move_after` - (Computed, Integer) Specifies the time period after which the backup will be moved from current tier to next tier.
							* `move_after_unit` - (Computed, String) Specifies the unit for moving the data from current tier to next tier. This unit will be a base unit for the 'moveAfter' field specified below.
							  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
							* `tier_type` - (Computed, String) Specifies the AWS tier types.
							  * Constraints: Allowable values are: `kAmazonS3Standard`, `kAmazonS3StandardIA`, `kAmazonS3OneZoneIA`, `kAmazonS3IntelligentTiering`, `kAmazonS3Glacier`, `kAmazonS3GlacierDeepArchive`.
					* `azure_tiering` - (Optional, List) Specifies Azure tiers.
					Nested schema for **azure_tiering**:
						* `tiers` - (Optional, List) Specifies the tiers that are used to move the archived backup from current tier to next tier. The order of the tiers determines which tier will be used next for moving the archived backup. The first tier input should always be default tier where backup will be acrhived. Each tier specifies how much time after the backup will be moved to next tier from the current tier.
						Nested schema for **tiers**:
							* `move_after` - (Computed, Integer) Specifies the time period after which the backup will be moved from current tier to next tier.
							* `move_after_unit` - (Computed, String) Specifies the unit for moving the data from current tier to next tier. This unit will be a base unit for the 'moveAfter' field specified below.
							  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
							* `tier_type` - (Computed, String) Specifies the Azure tier types.
							  * Constraints: Allowable values are: `kAzureTierHot`, `kAzureTierCool`, `kAzureTierArchive`.
					* `cloud_platform` - (Optional, String) Specifies the cloud platform to enable tiering.
					  * Constraints: Allowable values are: `AWS`, `Azure`, `Oracle`, `Google`.
					* `google_tiering` - (Optional, List) Specifies Google tiers.
					Nested schema for **google_tiering**:
						* `tiers` - (Required, List) Specifies the tiers that are used to move the archived backup from current tier to next tier. The order of the tiers determines which tier will be used next for moving the archived backup. The first tier input should always be default tier where backup will be acrhived. Each tier specifies how much time after the backup will be moved to next tier from the current tier.
						Nested schema for **tiers**:
							* `move_after` - (Computed, Integer) Specifies the time period after which the backup will be moved from current tier to next tier.
							* `move_after_unit` - (Computed, String) Specifies the unit for moving the data from current tier to next tier. This unit will be a base unit for the 'moveAfter' field specified below.
							  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
							* `tier_type` - (Computed, String) Specifies the Google tier types.
							  * Constraints: Allowable values are: `kGoogleStandard`, `kGoogleRegional`, `kGoogleMultiRegional`, `kGoogleNearline`, `kGoogleColdline`.
					* `oracle_tiering` - (Optional, List) Specifies Oracle tiers.
					Nested schema for **oracle_tiering**:
						* `tiers` - (Required, List) Specifies the tiers that are used to move the archived backup from current tier to next tier. The order of the tiers determines which tier will be used next for moving the archived backup. The first tier input should always be default tier where backup will be acrhived. Each tier specifies how much time after the backup will be moved to next tier from the current tier.
						Nested schema for **tiers**:
							* `move_after` - (Computed, Integer) Specifies the time period after which the backup will be moved from current tier to next tier.
							* `move_after_unit` - (Computed, String) Specifies the unit for moving the data from current tier to next tier. This unit will be a base unit for the 'moveAfter' field specified below.
							  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
							* `tier_type` - (Computed, String) Specifies the Oracle tier types.
							  * Constraints: Allowable values are: `kOracleTierStandard`, `kOracleTierArchive`.
			* `target_type` - (Optional, String) Specifies the primary backup location where backups will be stored. If not specified, then default is assumed as local backup on Cohesity cluster.
			  * Constraints: Allowable values are: `Local`, `Archival`.
			* `use_default_backup_target` - (Optional, Boolean) Specifies if the default primary backup target must be used for backups. If this is not specified or set to false, then targets specified in 'archivalTargetSettings' will be used for backups. If the value is specified as true, then default backup target is used internally. This field should only be set in the environment where tenant policy management is enabled and external targets are assigned to tenant when provisioning tenants.
		* `retention` - (Optional, List) Specifies the retention of a backup.
		Nested schema for **retention**:
			* `data_lock_config` - (Optional, List) Specifies WORM retention type for the snapshots. When a WORM retention type is specified, the snapshots of the Protection Groups using this policy will be kept for the last N days as specified in the duration of the datalock. During that time, the snapshots cannot be deleted.
			Nested schema for **data_lock_config**:
				* `duration` - (Required, Integer) Specifies the duration for a dataLock. <br> Example. If duration is 7 and unit is Months, the dataLock is enabled for last 7 * 30 = 210 days of the backup.
				  * Constraints: The minimum value is `1`.
				* `enable_worm_on_external_target` - (Optional, Boolean) Specifies whether objects in the external target associated with this policy need to be made immutable.
				* `mode` - (Required, String) Specifies the type of WORM retention type. 'Compliance' implies WORM retention is set for compliance reason. 'Administrative' implies WORM retention is set for administrative purposes.
				  * Constraints: Allowable values are: `Compliance`, `Administrative`.
				* `unit` - (Required, String) Specificies the Retention Unit of a dataLock measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Months' then number of retention days will be 365 * 2 = 730 days.
				  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
			* `duration` - (Required, Integer) Specifies the duration for a backup retention. <br> Example. If duration is 7 and unit is Months, the retention of a backup is 7 * 30 = 210 days.
			  * Constraints: The minimum value is `1`.
			* `unit` - (Required, String) Specificies the Retention Unit of a backup measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Years' then number of retention days will be 365 * 2 = 730 days.
			  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
	* `run_timeouts` - (Optional, List) Specifies the backup timeouts for different type of runs(kFull, kRegular etc.).
	Nested schema for **run_timeouts**:
		* `backup_type` - (Optional, String) The scheduled backup type(kFull, kRegular etc.).
		  * Constraints: Allowable values are: `kRegular`, `kFull`, `kLog`, `kSystem`, `kHydrateCDP`, `kStorageArraySnapshot`.
		* `timeout_mins` - (Optional, Integer) Specifies the timeout in mins.
	* `storage_array_snapshot` - (Optional, List) Specifies storage snapshot managment backup settings for a Protection Group.
	Nested schema for **storage_array_snapshot**:
		* `retention` - (Required, List) Specifies the retention of a backup.
		Nested schema for **retention**:
			* `data_lock_config` - (Optional, List) Specifies WORM retention type for the snapshots. When a WORM retention type is specified, the snapshots of the Protection Groups using this policy will be kept for the last N days as specified in the duration of the datalock. During that time, the snapshots cannot be deleted.
			Nested schema for **data_lock_config**:
				* `duration` - (Required, Integer) Specifies the duration for a dataLock. <br> Example. If duration is 7 and unit is Months, the dataLock is enabled for last 7 * 30 = 210 days of the backup.
				  * Constraints: The minimum value is `1`.
				* `enable_worm_on_external_target` - (Optional, Boolean) Specifies whether objects in the external target associated with this policy need to be made immutable.
				* `mode` - (Required, String) Specifies the type of WORM retention type. 'Compliance' implies WORM retention is set for compliance reason. 'Administrative' implies WORM retention is set for administrative purposes.
				  * Constraints: Allowable values are: `Compliance`, `Administrative`.
				* `unit` - (Required, String) Specificies the Retention Unit of a dataLock measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Months' then number of retention days will be 365 * 2 = 730 days.
				  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
			* `duration` - (Required, Integer) Specifies the duration for a backup retention. <br> Example. If duration is 7 and unit is Months, the retention of a backup is 7 * 30 = 210 days.
			  * Constraints: The minimum value is `1`.
			* `unit` - (Required, String) Specificies the Retention Unit of a backup measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Years' then number of retention days will be 365 * 2 = 730 days.
			  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
		* `schedule` - (Required, List) Specifies settings that defines how frequent Storage Snapshot Management backup will be performed for a Protection Group.
		Nested schema for **schedule**:
			* `day_schedule` - (Optional, List) Specifies settings that define a schedule for a Protection Group runs to start after certain number of days.
			Nested schema for **day_schedule**:
				* `frequency` - (Required, Integer) Specifies a factor to multiply the unit by, to determine the backup schedule. <br> Example: If 'frequency' set to 2 and the unit is 'Hours', then Snapshots are backed up every 2 hours. <br> This field is only applicable if unit is 'Minutes', 'Hours' or 'Days'.
				  * Constraints: The minimum value is `1`.
			* `hour_schedule` - (Optional, List) Specifies settings that define a schedule for a Protection Group runs to start after certain number of hours.
			Nested schema for **hour_schedule**:
				* `frequency` - (Required, Integer) Specifies a factor to multiply the unit by, to determine the backup schedule. <br> Example: If 'frequency' set to 2 and the unit is 'Hours', then Snapshots are backed up every 2 hours. <br> This field is only applicable if unit is 'Minutes', 'Hours' or 'Days'.
				  * Constraints: The minimum value is `1`.
			* `minute_schedule` - (Optional, List) Specifies settings that define a schedule for a Protection Group runs to start after certain number of minutes.
			Nested schema for **minute_schedule**:
				* `frequency` - (Required, Integer) Specifies a factor to multiply the unit by, to determine the backup schedule. <br> Example: If 'frequency' set to 2 and the unit is 'Hours', then Snapshots are backed up every 2 hours. <br> This field is only applicable if unit is 'Minutes', 'Hours' or 'Days'.
				  * Constraints: The minimum value is `1`.
			* `month_schedule` - (Optional, List) Specifies settings that define a schedule for a Protection Group runs to on specific week and specific days of that week.
			Nested schema for **month_schedule**:
				* `day_of_month` - (Optional, Integer) Specifies the exact date of the month (such as 18) in a Monthly Schedule specified by unit field as 'Years'. <br> Example: if 'dayOfMonth' is set to '18', a backup is performed on the 18th of every month.
				* `day_of_week` - (Optional, List) Specifies a list of days of the week when to start Protection Group Runs. <br> Example: To run a Protection Group on every Monday and Tuesday, set the schedule with following values: <br>  unit: 'Weeks' <br>  dayOfWeek: ['Monday','Tuesday'].
				  * Constraints: Allowable list items are: `Sunday`, `Monday`, `Tuesday`, `Wednesday`, `Thursday`, `Friday`, `Saturday`.
				* `week_of_month` - (Optional, String) Specifies the week of the month (such as 'Third') or nth day of month (such as 'First' or 'Last') in a Monthly Schedule specified by unit field as 'Months'. <br>This field can be used in combination with 'dayOfWeek' to define the day in the month to start the Protection Group Run. <br> Example: if 'weekOfMonth' is set to 'Third' and day is set to 'Monday', a backup is performed on the third Monday of every month. <br> Example: if 'weekOfMonth' is set to 'Last' and dayOfWeek is not set, a backup is performed on the last day of every month.
				  * Constraints: Allowable values are: `First`, `Second`, `Third`, `Fourth`, `Last`.
			* `unit` - (Required, String) Specifies how often to start new Protection Group Runs of a Protection Group. <br>'Minutes' specifies that Protection Group run starts periodically after certain number of minutes specified in 'frequency' field. <br>'Hours' specifies that Protection Group run starts periodically after certain number of hours specified in 'frequency' field.
			  * Constraints: Allowable values are: `Minutes`, `Hours`, `Days`, `Weeks`, `Months`, `Years`.
			* `week_schedule` - (Optional, List) Specifies settings that define a schedule for a Protection Group runs to start on certain days of week.
			Nested schema for **week_schedule**:
				* `day_of_week` - (Required, List) Specifies a list of days of the week when to start Protection Group Runs. <br> Example: To run a Protection Group on every Monday and Tuesday, set the schedule with following values: <br>  unit: 'Weeks' <br>  dayOfWeek: ['Monday','Tuesday'].
				  * Constraints: Allowable list items are: `Sunday`, `Monday`, `Tuesday`, `Wednesday`, `Thursday`, `Friday`, `Saturday`.
			* `year_schedule` - (Optional, List) Specifies settings that define a schedule for a Protection Group to run on specific year and specific day of that year.
			Nested schema for **year_schedule**:
				* `day_of_year` - (Required, String) Specifies the day of the Year (such as 'First' or 'Last') in a Yearly Schedule. <br>This field is used to define the day in the year to start the Protection Group Run. <br> Example: if 'dayOfYear' is set to 'First', a backup is performed on the first day of every year. <br> Example: if 'dayOfYear' is set to 'Last', a backup is performed on the last day of every year.
				  * Constraints: Allowable values are: `First`, `Last`.
* `blackout_window` - (Optional, List) List of Blackout Windows. If specified, this field defines blackout periods when new Group Runs are not started. If a Group Run has been scheduled but not yet executed and the blackout period starts, the behavior depends on the policy field AbortInBlackoutPeriod.
Nested schema for **blackout_window**:
	* `config_id` - (Optional, String) Specifies the unique identifier for the target getting added. This field need to be passed olny when policies are updated.
	* `day` - (Required, String) Specifies a day in the week when no new Protection Group Runs should be started such as 'Sunday'. Specifies a day in a week such as 'Sunday', 'Monday', etc.
	  * Constraints: Allowable values are: `Sunday`, `Monday`, `Tuesday`, `Wednesday`, `Thursday`, `Friday`, `Saturday`.
	* `end_time` - (Required, List) Specifies the time of day. Used for scheduling purposes.
	Nested schema for **end_time**:
		* `hour` - (Required, Integer) Specifies the hour of the day (0-23).
		  * Constraints: The maximum value is `23`. The minimum value is `0`.
		* `minute` - (Required, Integer) Specifies the minute of the hour (0-59).
		  * Constraints: The maximum value is `59`. The minimum value is `0`.
		* `time_zone` - (Optional, String) Specifies the time zone of the user. If not specified, default value is assumed as America/Los_Angeles.
		  * Constraints: The default value is `America/Los_Angeles`.
	* `start_time` - (Required, List) Specifies the time of day. Used for scheduling purposes.
	Nested schema for **start_time**:
		* `hour` - (Required, Integer) Specifies the hour of the day (0-23).
		  * Constraints: The maximum value is `23`. The minimum value is `0`.
		* `minute` - (Required, Integer) Specifies the minute of the hour (0-59).
		  * Constraints: The maximum value is `59`. The minimum value is `0`.
		* `time_zone` - (Optional, String) Specifies the time zone of the user. If not specified, default value is assumed as America/Los_Angeles.
		  * Constraints: The default value is `America/Los_Angeles`.
* `cascaded_targets_config` - (Optional, List) Specifies the configuration for cascaded replications. Using cascaded replication, replication cluster(Rx) can further replicate and archive the snapshot copies to further targets. Its recommended to create cascaded configuration where protection group will be created.
Nested schema for **cascaded_targets_config**:
	* `remote_targets` - (Required, List) Specifies the replication, archival and cloud spin targets of Protection Policy.
	Nested schema for **remote_targets**:
		* `archival_targets` - (Optional, List)
		Nested schema for **archival_targets**:
			* `backup_run_type` - (Optional, String) Specifies which type of run should be copied, if not set, all types of runs will be eligible for copying. If set, this will ensure that the first run of given type in the scheduled period will get copied. Currently, this can only be set to Full.
			  * Constraints: Allowable values are: `Regular`, `Full`, `Log`, `System`, `StorageArraySnapshot`.
			* `config_id` - (Optional, String) Specifies the unique identifier for the target getting added. This field need to be passed only when policies are being updated.
			* `copy_on_run_success` - (Optional, Boolean) Specifies if Snapshots are copied from the first completely successful Protection Group Run or the first partially successful Protection Group Run occurring at the start of the replication schedule. <br> If true, Snapshots are copied from the first Protection Group Run occurring at the start of the replication schedule that was completely successful i.e. Snapshots for all the Objects in the Protection Group were successfully captured. <br> If false, Snapshots are copied from the first Protection Group Run occurring at the start of the replication schedule, even if first Protection Group Run was not completely successful i.e. Snapshots were not captured for all Objects in the Protection Group.
			* `extended_retention` - (Optional, List) Specifies additional retention policies that should be applied to the archived backup. Archived backup snapshot will be retained up to a time that is the maximum of all retention policies that are applicable to it.
			Nested schema for **extended_retention**:
				* `config_id` - (Optional, String) Specifies the unique identifier for the target getting added. This field need to be passed olny when policies are updated.
				* `retention` - (Required, List) Specifies the retention of a backup.
				Nested schema for **retention**:
					* `data_lock_config` - (Optional, List) Specifies WORM retention type for the snapshots. When a WORM retention type is specified, the snapshots of the Protection Groups using this policy will be kept for the last N days as specified in the duration of the datalock. During that time, the snapshots cannot be deleted.
					Nested schema for **data_lock_config**:
						* `duration` - (Required, Integer) Specifies the duration for a dataLock. <br> Example. If duration is 7 and unit is Months, the dataLock is enabled for last 7 * 30 = 210 days of the backup.
						  * Constraints: The minimum value is `1`.
						* `enable_worm_on_external_target` - (Optional, Boolean) Specifies whether objects in the external target associated with this policy need to be made immutable.
						* `mode` - (Required, String) Specifies the type of WORM retention type. 'Compliance' implies WORM retention is set for compliance reason. 'Administrative' implies WORM retention is set for administrative purposes.
						  * Constraints: Allowable values are: `Compliance`, `Administrative`.
						* `unit` - (Required, String) Specificies the Retention Unit of a dataLock measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Months' then number of retention days will be 365 * 2 = 730 days.
						  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
					* `duration` - (Required, Integer) Specifies the duration for a backup retention. <br> Example. If duration is 7 and unit is Months, the retention of a backup is 7 * 30 = 210 days.
					  * Constraints: The minimum value is `1`.
					* `unit` - (Required, String) Specificies the Retention Unit of a backup measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Years' then number of retention days will be 365 * 2 = 730 days.
					  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
				* `run_type` - (Optional, String) The backup run type to which this extended retention applies to. If this is not set, the extended retention will be applicable to all non-log backup types. Currently, the only value that can be set here is Full.'Regular' indicates a incremental (CBT) backup. Incremental backups utilizing CBT (if supported) are captured of the target protection objects. The first run of a Regular schedule captures all the blocks.'Full' indicates a full (no CBT) backup. A complete backup (all blocks) of the target protection objects are always captured and Change Block Tracking (CBT) is not utilized.'Log' indicates a Database Log backup. Capture the database transaction logs to allow rolling back to a specific point in time.'System' indicates a system backup. System backups are used to do bare metal recovery of the system to a specific point in time.
				  * Constraints: Allowable values are: `Regular`, `Full`, `Log`, `System`, `StorageArraySnapshot`.
				* `schedule` - (Required, List) Specifies a schedule frequency and schedule unit for Extended Retentions.
				Nested schema for **schedule**:
					* `frequency` - (Optional, Integer) Specifies a factor to multiply the unit by, to determine the retention schedule. For example if set to 2 and the unit is hourly, then Snapshots from the first eligible Job Run for every 2 hour period is retained.
					  * Constraints: The minimum value is `1`.
					* `unit` - (Required, String) Specifies the unit interval for retention of Snapshots. <br>'Runs' means that the Snapshot copy retained after the number of Protection Group Runs equals the number specified in the frequency. <br>'Hours' means that the Snapshot copy retained hourly at the frequency set in the frequency, for example if scheduleFrequency is 2, the copy occurs every 2 hours. <br>'Days' means that the Snapshot copy gets retained daily at the frequency set in the frequency. <br>'Weeks' means that the Snapshot copy is retained weekly at the frequency set in the frequency. <br>'Months' means that the Snapshot copy is retained monthly at the frequency set in the Frequency. <br>'Years' means that the Snapshot copy is retained yearly at the frequency set in the Frequency.
					  * Constraints: Allowable values are: `Runs`, `Hours`, `Days`, `Weeks`, `Months`, `Years`.
			* `log_retention` - (Optional, List) Specifies the retention of a backup.
			Nested schema for **log_retention**:
				* `data_lock_config` - (Optional, List) Specifies WORM retention type for the snapshots. When a WORM retention type is specified, the snapshots of the Protection Groups using this policy will be kept for the last N days as specified in the duration of the datalock. During that time, the snapshots cannot be deleted.
				Nested schema for **data_lock_config**:
					* `duration` - (Required, Integer) Specifies the duration for a dataLock. <br> Example. If duration is 7 and unit is Months, the dataLock is enabled for last 7 * 30 = 210 days of the backup.
					  * Constraints: The minimum value is `1`.
					* `enable_worm_on_external_target` - (Optional, Boolean) Specifies whether objects in the external target associated with this policy need to be made immutable.
					* `mode` - (Required, String) Specifies the type of WORM retention type. 'Compliance' implies WORM retention is set for compliance reason. 'Administrative' implies WORM retention is set for administrative purposes.
					  * Constraints: Allowable values are: `Compliance`, `Administrative`.
					* `unit` - (Required, String) Specificies the Retention Unit of a dataLock measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Months' then number of retention days will be 365 * 2 = 730 days.
					  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
				* `duration` - (Required, Integer) Specifies the duration for a backup retention. <br> Example. If duration is 7 and unit is Months, the retention of a backup is 7 * 30 = 210 days.
				  * Constraints: The minimum value is `0`.
				* `unit` - (Required, String) Specificies the Retention Unit of a backup measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Years' then number of retention days will be 365 * 2 = 730 days.
				  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
			* `retention` - (Required, List) Specifies the retention of a backup.
			Nested schema for **retention**:
				* `data_lock_config` - (Optional, List) Specifies WORM retention type for the snapshots. When a WORM retention type is specified, the snapshots of the Protection Groups using this policy will be kept for the last N days as specified in the duration of the datalock. During that time, the snapshots cannot be deleted.
				Nested schema for **data_lock_config**:
					* `duration` - (Required, Integer) Specifies the duration for a dataLock. <br> Example. If duration is 7 and unit is Months, the dataLock is enabled for last 7 * 30 = 210 days of the backup.
					  * Constraints: The minimum value is `1`.
					* `enable_worm_on_external_target` - (Optional, Boolean) Specifies whether objects in the external target associated with this policy need to be made immutable.
					* `mode` - (Required, String) Specifies the type of WORM retention type. 'Compliance' implies WORM retention is set for compliance reason. 'Administrative' implies WORM retention is set for administrative purposes.
					  * Constraints: Allowable values are: `Compliance`, `Administrative`.
					* `unit` - (Required, String) Specificies the Retention Unit of a dataLock measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Months' then number of retention days will be 365 * 2 = 730 days.
					  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
				* `duration` - (Required, Integer) Specifies the duration for a backup retention. <br> Example. If duration is 7 and unit is Months, the retention of a backup is 7 * 30 = 210 days.
				  * Constraints: The minimum value is `1`.
				* `unit` - (Required, String) Specificies the Retention Unit of a backup measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Years' then number of retention days will be 365 * 2 = 730 days.
				  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
			* `run_timeouts` - (Optional, List) Specifies the replication/archival timeouts for different type of runs(kFull, kRegular etc.).
			Nested schema for **run_timeouts**:
				* `backup_type` - (Optional, String) The scheduled backup type(kFull, kRegular etc.).
				  * Constraints: Allowable values are: `kRegular`, `kFull`, `kLog`, `kSystem`, `kHydrateCDP`, `kStorageArraySnapshot`.
				* `timeout_mins` - (Optional, Integer) Specifies the timeout in mins.
			* `schedule` - (Required, List) Specifies a schedule fregquency and schedule unit for copying Snapshots to backup targets.
			Nested schema for **schedule**:
				* `frequency` - (Optional, Integer) Specifies a factor to multiply the unit by, to determine the copy schedule. For example if set to 2 and the unit is hourly, then Snapshots from the first eligible Job Run for every 2 hour period is copied.
				  * Constraints: The minimum value is `1`.
				* `unit` - (Required, String) Specifies the frequency that Snapshots should be copied to the specified target. Used in combination with multiplier. <br>'Runs' means that the Snapshot copy occurs after the number of Protection Group Runs equals the number specified in the frequency. <br>'Hours' means that the Snapshot copy occurs hourly at the frequency set in the frequency, for example if scheduleFrequency is 2, the copy occurs every 2 hours. <br>'Days' means that the Snapshot copy occurs daily at the frequency set in the frequency. <br>'Weeks' means that the Snapshot copy occurs weekly at the frequency set in the frequency. <br>'Months' means that the Snapshot copy occurs monthly at the frequency set in the Frequency. <br>'Years' means that the Snapshot copy occurs yearly at the frequency set in the scheduleFrequency.
				  * Constraints: Allowable values are: `Runs`, `Hours`, `Days`, `Weeks`, `Months`, `Years`.
			* `target_id` - (Required, Integer) Specifies the Archival target to copy the Snapshots to.
			* `target_name` - (Computed, String) Specifies the Archival target name where Snapshots are copied.
			* `target_type` - (Computed, String) Specifies the Archival target type where Snapshots are copied.
			  * Constraints: Allowable values are: `Tape`, `Cloud`, `Nas`.
			* `tier_settings` - (Optional, List) Specifies the settings tier levels configured with each archival target. The tier settings need to be applied in specific order and default tier should always be passed as first entry in tiers array. The following example illustrates how to configure tiering input for AWS tiering. Same type of input structure applied to other cloud platforms also. <br>If user wants to achieve following tiering for backup, <br>User Desired Tiering- <br><t>1.Archive Full back up for 12 Months <br><t>2.Tier Levels <br><t><t>[1,12] [ <br><t><t><t>s3 (1 to 2 months), (default tier) <br><t><t><t>s3 Intelligent tiering (3 to 6 months), <br><t><t><t>s3 One Zone (7 to 9 months) <br><t><t><t>Glacier (10 to 12 months)] <br><t>API Input <br><t><t>1.tiers-[ <br><t><t><t>{'tierType': 'S3','moveAfterUnit':'months', <br><t><t><t>'moveAfter':2 - move from s3 to s3Inte after 2 months}, <br><t><t><t>{'tierType': 'S3Inte','moveAfterUnit':'months', <br><t><t><t>'moveAfter':4 - move from S3Inte to Glacier after 4 months}, <br><t><t><t>{'tierType': 'Glacier', 'moveAfterUnit':'months', <br><t><t><t>'moveAfter': 3 - move from Glacier to S3 One Zone after 3 months }, <br><t><t><t>{'tierType': 'S3 One Zone', 'moveAfterUnit': nil, <br><t><t><t>'moveAfter': nil - For the last record, 'moveAfter' and 'moveAfterUnit' <br><t><t><t>will be ignored since there are no further tier for data movement } <br><t><t><t>}].
			Nested schema for **tier_settings**:
				* `aws_tiering` - (Optional, List) Specifies aws tiers.
				Nested schema for **aws_tiering**:
					* `tiers` - (Required, List) Specifies the tiers that are used to move the archived backup from current tier to next tier. The order of the tiers determines which tier will be used next for moving the archived backup. The first tier input should always be default tier where backup will be acrhived. Each tier specifies how much time after the backup will be moved to next tier from the current tier.
					Nested schema for **tiers**:
						* `move_after` - (Computed, Integer) Specifies the time period after which the backup will be moved from current tier to next tier.
						* `move_after_unit` - (Computed, String) Specifies the unit for moving the data from current tier to next tier. This unit will be a base unit for the 'moveAfter' field specified below.
						  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
						* `tier_type` - (Computed, String) Specifies the AWS tier types.
						  * Constraints: Allowable values are: `kAmazonS3Standard`, `kAmazonS3StandardIA`, `kAmazonS3OneZoneIA`, `kAmazonS3IntelligentTiering`, `kAmazonS3Glacier`, `kAmazonS3GlacierDeepArchive`.
				* `azure_tiering` - (Optional, List) Specifies Azure tiers.
				Nested schema for **azure_tiering**:
					* `tiers` - (Optional, List) Specifies the tiers that are used to move the archived backup from current tier to next tier. The order of the tiers determines which tier will be used next for moving the archived backup. The first tier input should always be default tier where backup will be acrhived. Each tier specifies how much time after the backup will be moved to next tier from the current tier.
					Nested schema for **tiers**:
						* `move_after` - (Computed, Integer) Specifies the time period after which the backup will be moved from current tier to next tier.
						* `move_after_unit` - (Computed, String) Specifies the unit for moving the data from current tier to next tier. This unit will be a base unit for the 'moveAfter' field specified below.
						  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
						* `tier_type` - (Computed, String) Specifies the Azure tier types.
						  * Constraints: Allowable values are: `kAzureTierHot`, `kAzureTierCool`, `kAzureTierArchive`.
				* `cloud_platform` - (Optional, String) Specifies the cloud platform to enable tiering.
				  * Constraints: Allowable values are: `AWS`, `Azure`, `Oracle`, `Google`.
				* `google_tiering` - (Optional, List) Specifies Google tiers.
				Nested schema for **google_tiering**:
					* `tiers` - (Required, List) Specifies the tiers that are used to move the archived backup from current tier to next tier. The order of the tiers determines which tier will be used next for moving the archived backup. The first tier input should always be default tier where backup will be acrhived. Each tier specifies how much time after the backup will be moved to next tier from the current tier.
					Nested schema for **tiers**:
						* `move_after` - (Computed, Integer) Specifies the time period after which the backup will be moved from current tier to next tier.
						* `move_after_unit` - (Computed, String) Specifies the unit for moving the data from current tier to next tier. This unit will be a base unit for the 'moveAfter' field specified below.
						  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
						* `tier_type` - (Computed, String) Specifies the Google tier types.
						  * Constraints: Allowable values are: `kGoogleStandard`, `kGoogleRegional`, `kGoogleMultiRegional`, `kGoogleNearline`, `kGoogleColdline`.
				* `oracle_tiering` - (Optional, List) Specifies Oracle tiers.
				Nested schema for **oracle_tiering**:
					* `tiers` - (Required, List) Specifies the tiers that are used to move the archived backup from current tier to next tier. The order of the tiers determines which tier will be used next for moving the archived backup. The first tier input should always be default tier where backup will be acrhived. Each tier specifies how much time after the backup will be moved to next tier from the current tier.
					Nested schema for **tiers**:
						* `move_after` - (Computed, Integer) Specifies the time period after which the backup will be moved from current tier to next tier.
						* `move_after_unit` - (Computed, String) Specifies the unit for moving the data from current tier to next tier. This unit will be a base unit for the 'moveAfter' field specified below.
						  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
						* `tier_type` - (Computed, String) Specifies the Oracle tier types.
						  * Constraints: Allowable values are: `kOracleTierStandard`, `kOracleTierArchive`.
		* `cloud_spin_targets` - (Optional, List)
		Nested schema for **cloud_spin_targets**:
			* `backup_run_type` - (Optional, String) Specifies which type of run should be copied, if not set, all types of runs will be eligible for copying. If set, this will ensure that the first run of given type in the scheduled period will get copied. Currently, this can only be set to Full.
			  * Constraints: Allowable values are: `Regular`, `Full`, `Log`, `System`, `StorageArraySnapshot`.
			* `config_id` - (Optional, String) Specifies the unique identifier for the target getting added. This field need to be passed only when policies are being updated.
			* `copy_on_run_success` - (Optional, Boolean) Specifies if Snapshots are copied from the first completely successful Protection Group Run or the first partially successful Protection Group Run occurring at the start of the replication schedule. <br> If true, Snapshots are copied from the first Protection Group Run occurring at the start of the replication schedule that was completely successful i.e. Snapshots for all the Objects in the Protection Group were successfully captured. <br> If false, Snapshots are copied from the first Protection Group Run occurring at the start of the replication schedule, even if first Protection Group Run was not completely successful i.e. Snapshots were not captured for all Objects in the Protection Group.
			* `log_retention` - (Optional, List) Specifies the retention of a backup.
			Nested schema for **log_retention**:
				* `data_lock_config` - (Optional, List) Specifies WORM retention type for the snapshots. When a WORM retention type is specified, the snapshots of the Protection Groups using this policy will be kept for the last N days as specified in the duration of the datalock. During that time, the snapshots cannot be deleted.
				Nested schema for **data_lock_config**:
					* `duration` - (Required, Integer) Specifies the duration for a dataLock. <br> Example. If duration is 7 and unit is Months, the dataLock is enabled for last 7 * 30 = 210 days of the backup.
					  * Constraints: The minimum value is `1`.
					* `enable_worm_on_external_target` - (Optional, Boolean) Specifies whether objects in the external target associated with this policy need to be made immutable.
					* `mode` - (Required, String) Specifies the type of WORM retention type. 'Compliance' implies WORM retention is set for compliance reason. 'Administrative' implies WORM retention is set for administrative purposes.
					  * Constraints: Allowable values are: `Compliance`, `Administrative`.
					* `unit` - (Required, String) Specificies the Retention Unit of a dataLock measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Months' then number of retention days will be 365 * 2 = 730 days.
					  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
				* `duration` - (Required, Integer) Specifies the duration for a backup retention. <br> Example. If duration is 7 and unit is Months, the retention of a backup is 7 * 30 = 210 days.
				  * Constraints: The minimum value is `0`.
				* `unit` - (Required, String) Specificies the Retention Unit of a backup measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Years' then number of retention days will be 365 * 2 = 730 days.
				  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
			* `retention` - (Required, List) Specifies the retention of a backup.
			Nested schema for **retention**:
				* `data_lock_config` - (Optional, List) Specifies WORM retention type for the snapshots. When a WORM retention type is specified, the snapshots of the Protection Groups using this policy will be kept for the last N days as specified in the duration of the datalock. During that time, the snapshots cannot be deleted.
				Nested schema for **data_lock_config**:
					* `duration` - (Required, Integer) Specifies the duration for a dataLock. <br> Example. If duration is 7 and unit is Months, the dataLock is enabled for last 7 * 30 = 210 days of the backup.
					  * Constraints: The minimum value is `1`.
					* `enable_worm_on_external_target` - (Optional, Boolean) Specifies whether objects in the external target associated with this policy need to be made immutable.
					* `mode` - (Required, String) Specifies the type of WORM retention type. 'Compliance' implies WORM retention is set for compliance reason. 'Administrative' implies WORM retention is set for administrative purposes.
					  * Constraints: Allowable values are: `Compliance`, `Administrative`.
					* `unit` - (Required, String) Specificies the Retention Unit of a dataLock measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Months' then number of retention days will be 365 * 2 = 730 days.
					  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
				* `duration` - (Required, Integer) Specifies the duration for a backup retention. <br> Example. If duration is 7 and unit is Months, the retention of a backup is 7 * 30 = 210 days.
				  * Constraints: The minimum value is `1`.
				* `unit` - (Required, String) Specificies the Retention Unit of a backup measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Years' then number of retention days will be 365 * 2 = 730 days.
				  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
			* `run_timeouts` - (Optional, List) Specifies the replication/archival timeouts for different type of runs(kFull, kRegular etc.).
			Nested schema for **run_timeouts**:
				* `backup_type` - (Optional, String) The scheduled backup type(kFull, kRegular etc.).
				  * Constraints: Allowable values are: `kRegular`, `kFull`, `kLog`, `kSystem`, `kHydrateCDP`, `kStorageArraySnapshot`.
				* `timeout_mins` - (Optional, Integer) Specifies the timeout in mins.
			* `schedule` - (Required, List) Specifies a schedule fregquency and schedule unit for copying Snapshots to backup targets.
			Nested schema for **schedule**:
				* `frequency` - (Optional, Integer) Specifies a factor to multiply the unit by, to determine the copy schedule. For example if set to 2 and the unit is hourly, then Snapshots from the first eligible Job Run for every 2 hour period is copied.
				  * Constraints: The minimum value is `1`.
				* `unit` - (Required, String) Specifies the frequency that Snapshots should be copied to the specified target. Used in combination with multiplier. <br>'Runs' means that the Snapshot copy occurs after the number of Protection Group Runs equals the number specified in the frequency. <br>'Hours' means that the Snapshot copy occurs hourly at the frequency set in the frequency, for example if scheduleFrequency is 2, the copy occurs every 2 hours. <br>'Days' means that the Snapshot copy occurs daily at the frequency set in the frequency. <br>'Weeks' means that the Snapshot copy occurs weekly at the frequency set in the frequency. <br>'Months' means that the Snapshot copy occurs monthly at the frequency set in the Frequency. <br>'Years' means that the Snapshot copy occurs yearly at the frequency set in the scheduleFrequency.
				  * Constraints: Allowable values are: `Runs`, `Hours`, `Days`, `Weeks`, `Months`, `Years`.
			* `target` - (Required, List) Specifies the details about Cloud Spin target where backup snapshots may be converted and stored.
			Nested schema for **target**:
				* `aws_params` - (Optional, List) Specifies various resources when converting and deploying a VM to AWS.
				Nested schema for **aws_params**:
					* `custom_tag_list` - (Optional, List) Specifies tags of various resources when converting and deploying a VM to AWS.
					Nested schema for **custom_tag_list**:
						* `key` - (Computed, String) Specifies key of the custom tag.
						* `value` - (Computed, String) Specifies value of the custom tag.
					* `region` - (Required, Integer) Specifies id of the AWS region in which to deploy the VM.
					* `subnet_id` - (Optional, Integer) Specifies id of the subnet within above VPC.
					* `vpc_id` - (Optional, Integer) Specifies id of the Virtual Private Cloud to chose for the instance type.
				* `azure_params` - (Optional, List) Specifies various resources when converting and deploying a VM to Azure.
				Nested schema for **azure_params**:
					* `availability_set_id` - (Optional, Integer) Specifies the availability set.
					* `network_resource_group_id` - (Optional, Integer) Specifies id of the resource group for the selected virtual network.
					* `resource_group_id` - (Optional, Integer) Specifies id of the Azure resource group. Its value is globally unique within Azure.
					* `storage_account_id` - (Optional, Integer) Specifies id of the storage account that will contain the storage container within which we will create the blob that will become the VHD disk for the cloned VM.
					* `storage_container_id` - (Optional, Integer) Specifies id of the storage container within the above storage account.
					* `storage_resource_group_id` - (Optional, Integer) Specifies id of the resource group for the selected storage account.
					* `temp_vm_resource_group_id` - (Optional, Integer) Specifies id of the temporary Azure resource group.
					* `temp_vm_storage_account_id` - (Optional, Integer) Specifies id of the temporary VM storage account that will contain the storage container within which we will create the blob that will become the VHD disk for the cloned VM.
					* `temp_vm_storage_container_id` - (Optional, Integer) Specifies id of the temporary VM storage container within the above storage account.
					* `temp_vm_subnet_id` - (Optional, Integer) Specifies Id of the temporary VM subnet within the above virtual network.
					* `temp_vm_virtual_network_id` - (Optional, Integer) Specifies Id of the temporary VM Virtual Network.
				* `id` - (Optional, Integer) Specifies the unique id of the cloud spin entity.
				* `name` - (Computed, String) Specifies the name of the already added cloud spin target.
		* `onprem_deploy_targets` - (Optional, List)
		Nested schema for **onprem_deploy_targets**:
			* `backup_run_type` - (Optional, String) Specifies which type of run should be copied, if not set, all types of runs will be eligible for copying. If set, this will ensure that the first run of given type in the scheduled period will get copied. Currently, this can only be set to Full.
			  * Constraints: Allowable values are: `Regular`, `Full`, `Log`, `System`, `StorageArraySnapshot`.
			* `config_id` - (Optional, String) Specifies the unique identifier for the target getting added. This field need to be passed only when policies are being updated.
			* `copy_on_run_success` - (Optional, Boolean) Specifies if Snapshots are copied from the first completely successful Protection Group Run or the first partially successful Protection Group Run occurring at the start of the replication schedule. <br> If true, Snapshots are copied from the first Protection Group Run occurring at the start of the replication schedule that was completely successful i.e. Snapshots for all the Objects in the Protection Group were successfully captured. <br> If false, Snapshots are copied from the first Protection Group Run occurring at the start of the replication schedule, even if first Protection Group Run was not completely successful i.e. Snapshots were not captured for all Objects in the Protection Group.
			* `log_retention` - (Optional, List) Specifies the retention of a backup.
			Nested schema for **log_retention**:
				* `data_lock_config` - (Optional, List) Specifies WORM retention type for the snapshots. When a WORM retention type is specified, the snapshots of the Protection Groups using this policy will be kept for the last N days as specified in the duration of the datalock. During that time, the snapshots cannot be deleted.
				Nested schema for **data_lock_config**:
					* `duration` - (Required, Integer) Specifies the duration for a dataLock. <br> Example. If duration is 7 and unit is Months, the dataLock is enabled for last 7 * 30 = 210 days of the backup.
					  * Constraints: The minimum value is `1`.
					* `enable_worm_on_external_target` - (Optional, Boolean) Specifies whether objects in the external target associated with this policy need to be made immutable.
					* `mode` - (Required, String) Specifies the type of WORM retention type. 'Compliance' implies WORM retention is set for compliance reason. 'Administrative' implies WORM retention is set for administrative purposes.
					  * Constraints: Allowable values are: `Compliance`, `Administrative`.
					* `unit` - (Required, String) Specificies the Retention Unit of a dataLock measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Months' then number of retention days will be 365 * 2 = 730 days.
					  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
				* `duration` - (Required, Integer) Specifies the duration for a backup retention. <br> Example. If duration is 7 and unit is Months, the retention of a backup is 7 * 30 = 210 days.
				  * Constraints: The minimum value is `0`.
				* `unit` - (Required, String) Specificies the Retention Unit of a backup measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Years' then number of retention days will be 365 * 2 = 730 days.
				  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
			* `params` - (Optional, List) Specifies the details about OnpremDeploy target where backup snapshots may be converted and deployed.
			Nested schema for **params**:
				* `id` - (Optional, Integer) Specifies the unique id of the onprem entity.
			* `retention` - (Required, List) Specifies the retention of a backup.
			Nested schema for **retention**:
				* `data_lock_config` - (Optional, List) Specifies WORM retention type for the snapshots. When a WORM retention type is specified, the snapshots of the Protection Groups using this policy will be kept for the last N days as specified in the duration of the datalock. During that time, the snapshots cannot be deleted.
				Nested schema for **data_lock_config**:
					* `duration` - (Required, Integer) Specifies the duration for a dataLock. <br> Example. If duration is 7 and unit is Months, the dataLock is enabled for last 7 * 30 = 210 days of the backup.
					  * Constraints: The minimum value is `1`.
					* `enable_worm_on_external_target` - (Optional, Boolean) Specifies whether objects in the external target associated with this policy need to be made immutable.
					* `mode` - (Required, String) Specifies the type of WORM retention type. 'Compliance' implies WORM retention is set for compliance reason. 'Administrative' implies WORM retention is set for administrative purposes.
					  * Constraints: Allowable values are: `Compliance`, `Administrative`.
					* `unit` - (Required, String) Specificies the Retention Unit of a dataLock measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Months' then number of retention days will be 365 * 2 = 730 days.
					  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
				* `duration` - (Required, Integer) Specifies the duration for a backup retention. <br> Example. If duration is 7 and unit is Months, the retention of a backup is 7 * 30 = 210 days.
				  * Constraints: The minimum value is `1`.
				* `unit` - (Required, String) Specificies the Retention Unit of a backup measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Years' then number of retention days will be 365 * 2 = 730 days.
				  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
			* `run_timeouts` - (Optional, List) Specifies the replication/archival timeouts for different type of runs(kFull, kRegular etc.).
			Nested schema for **run_timeouts**:
				* `backup_type` - (Optional, String) The scheduled backup type(kFull, kRegular etc.).
				  * Constraints: Allowable values are: `kRegular`, `kFull`, `kLog`, `kSystem`, `kHydrateCDP`, `kStorageArraySnapshot`.
				* `timeout_mins` - (Optional, Integer) Specifies the timeout in mins.
			* `schedule` - (Required, List) Specifies a schedule fregquency and schedule unit for copying Snapshots to backup targets.
			Nested schema for **schedule**:
				* `frequency` - (Optional, Integer) Specifies a factor to multiply the unit by, to determine the copy schedule. For example if set to 2 and the unit is hourly, then Snapshots from the first eligible Job Run for every 2 hour period is copied.
				  * Constraints: The minimum value is `1`.
				* `unit` - (Required, String) Specifies the frequency that Snapshots should be copied to the specified target. Used in combination with multiplier. <br>'Runs' means that the Snapshot copy occurs after the number of Protection Group Runs equals the number specified in the frequency. <br>'Hours' means that the Snapshot copy occurs hourly at the frequency set in the frequency, for example if scheduleFrequency is 2, the copy occurs every 2 hours. <br>'Days' means that the Snapshot copy occurs daily at the frequency set in the frequency. <br>'Weeks' means that the Snapshot copy occurs weekly at the frequency set in the frequency. <br>'Months' means that the Snapshot copy occurs monthly at the frequency set in the Frequency. <br>'Years' means that the Snapshot copy occurs yearly at the frequency set in the scheduleFrequency.
				  * Constraints: Allowable values are: `Runs`, `Hours`, `Days`, `Weeks`, `Months`, `Years`.
		* `replication_targets` - (Optional, List)
		Nested schema for **replication_targets**:
			* `aws_target_config` - (Optional, List) Specifies the configuration for adding AWS as repilcation target.
			Nested schema for **aws_target_config**:
				* `name` - (Computed, String) Specifies the name of the AWS Replication target.
				* `region` - (Required, Integer) Specifies id of the AWS region in which to replicate the Snapshot to. Applicable if replication target is AWS target.
				* `region_name` - (Computed, String) Specifies name of the AWS region in which to replicate the Snapshot to. Applicable if replication target is AWS target.
				* `source_id` - (Required, Integer) Specifies the source id of the AWS protection source registered on IBM cluster.
			* `azure_target_config` - (Optional, List) Specifies the configuration for adding Azure as replication target.
			Nested schema for **azure_target_config**:
				* `name` - (Computed, String) Specifies the name of the Azure Replication target.
				* `resource_group` - (Optional, Integer) Specifies id of the Azure resource group used to filter regions in UI.
				* `resource_group_name` - (Computed, String) Specifies name of the Azure resource group used to filter regions in UI.
				* `source_id` - (Required, Integer) Specifies the source id of the Azure protection source registered on IBM cluster.
				* `storage_account` - (Computed, Integer) Specifies id of the storage account of Azure replication target which will contain storage container.
				* `storage_account_name` - (Computed, String) Specifies name of the storage account of Azure replication target which will contain storage container.
				* `storage_container` - (Computed, Integer) Specifies id of the storage container of Azure Replication target.
				* `storage_container_name` - (Computed, String) Specifies name of the storage container of Azure Replication target.
				* `storage_resource_group` - (Computed, Integer) Specifies id of the storage resource group of Azure Replication target.
				* `storage_resource_group_name` - (Computed, String) Specifies name of the storage resource group of Azure Replication target.
			* `backup_run_type` - (Optional, String) Specifies which type of run should be copied, if not set, all types of runs will be eligible for copying. If set, this will ensure that the first run of given type in the scheduled period will get copied. Currently, this can only be set to Full.
			  * Constraints: Allowable values are: `Regular`, `Full`, `Log`, `System`, `StorageArraySnapshot`.
			* `config_id` - (Optional, String) Specifies the unique identifier for the target getting added. This field need to be passed only when policies are being updated.
			* `copy_on_run_success` - (Optional, Boolean) Specifies if Snapshots are copied from the first completely successful Protection Group Run or the first partially successful Protection Group Run occurring at the start of the replication schedule. <br> If true, Snapshots are copied from the first Protection Group Run occurring at the start of the replication schedule that was completely successful i.e. Snapshots for all the Objects in the Protection Group were successfully captured. <br> If false, Snapshots are copied from the first Protection Group Run occurring at the start of the replication schedule, even if first Protection Group Run was not completely successful i.e. Snapshots were not captured for all Objects in the Protection Group.
			* `log_retention` - (Optional, List) Specifies the retention of a backup.
			Nested schema for **log_retention**:
				* `data_lock_config` - (Optional, List) Specifies WORM retention type for the snapshots. When a WORM retention type is specified, the snapshots of the Protection Groups using this policy will be kept for the last N days as specified in the duration of the datalock. During that time, the snapshots cannot be deleted.
				Nested schema for **data_lock_config**:
					* `duration` - (Required, Integer) Specifies the duration for a dataLock. <br> Example. If duration is 7 and unit is Months, the dataLock is enabled for last 7 * 30 = 210 days of the backup.
					  * Constraints: The minimum value is `1`.
					* `enable_worm_on_external_target` - (Optional, Boolean) Specifies whether objects in the external target associated with this policy need to be made immutable.
					* `mode` - (Required, String) Specifies the type of WORM retention type. 'Compliance' implies WORM retention is set for compliance reason. 'Administrative' implies WORM retention is set for administrative purposes.
					  * Constraints: Allowable values are: `Compliance`, `Administrative`.
					* `unit` - (Required, String) Specificies the Retention Unit of a dataLock measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Months' then number of retention days will be 365 * 2 = 730 days.
					  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
				* `duration` - (Required, Integer) Specifies the duration for a backup retention. <br> Example. If duration is 7 and unit is Months, the retention of a backup is 7 * 30 = 210 days.
				  * Constraints: The minimum value is `0`.
				* `unit` - (Required, String) Specificies the Retention Unit of a backup measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Years' then number of retention days will be 365 * 2 = 730 days.
				  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
			* `remote_target_config` - (Optional, List) Specifies the configuration for adding remote cluster as repilcation target.
			Nested schema for **remote_target_config**:
				* `cluster_id` - (Required, Integer) Specifies the cluster id of the target replication cluster.
				* `cluster_name` - (Computed, String) Specifies the cluster name of the target replication cluster.
			* `retention` - (Required, List) Specifies the retention of a backup.
			Nested schema for **retention**:
				* `data_lock_config` - (Optional, List) Specifies WORM retention type for the snapshots. When a WORM retention type is specified, the snapshots of the Protection Groups using this policy will be kept for the last N days as specified in the duration of the datalock. During that time, the snapshots cannot be deleted.
				Nested schema for **data_lock_config**:
					* `duration` - (Required, Integer) Specifies the duration for a dataLock. <br> Example. If duration is 7 and unit is Months, the dataLock is enabled for last 7 * 30 = 210 days of the backup.
					  * Constraints: The minimum value is `1`.
					* `enable_worm_on_external_target` - (Optional, Boolean) Specifies whether objects in the external target associated with this policy need to be made immutable.
					* `mode` - (Required, String) Specifies the type of WORM retention type. 'Compliance' implies WORM retention is set for compliance reason. 'Administrative' implies WORM retention is set for administrative purposes.
					  * Constraints: Allowable values are: `Compliance`, `Administrative`.
					* `unit` - (Required, String) Specificies the Retention Unit of a dataLock measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Months' then number of retention days will be 365 * 2 = 730 days.
					  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
				* `duration` - (Required, Integer) Specifies the duration for a backup retention. <br> Example. If duration is 7 and unit is Months, the retention of a backup is 7 * 30 = 210 days.
				  * Constraints: The minimum value is `1`.
				* `unit` - (Required, String) Specificies the Retention Unit of a backup measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Years' then number of retention days will be 365 * 2 = 730 days.
				  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
			* `run_timeouts` - (Optional, List) Specifies the replication/archival timeouts for different type of runs(kFull, kRegular etc.).
			Nested schema for **run_timeouts**:
				* `backup_type` - (Optional, String) The scheduled backup type(kFull, kRegular etc.).
				  * Constraints: Allowable values are: `kRegular`, `kFull`, `kLog`, `kSystem`, `kHydrateCDP`, `kStorageArraySnapshot`.
				* `timeout_mins` - (Optional, Integer) Specifies the timeout in mins.
			* `schedule` - (Required, List) Specifies a schedule fregquency and schedule unit for copying Snapshots to backup targets.
			Nested schema for **schedule**:
				* `frequency` - (Optional, Integer) Specifies a factor to multiply the unit by, to determine the copy schedule. For example if set to 2 and the unit is hourly, then Snapshots from the first eligible Job Run for every 2 hour period is copied.
				  * Constraints: The minimum value is `1`.
				* `unit` - (Required, String) Specifies the frequency that Snapshots should be copied to the specified target. Used in combination with multiplier. <br>'Runs' means that the Snapshot copy occurs after the number of Protection Group Runs equals the number specified in the frequency. <br>'Hours' means that the Snapshot copy occurs hourly at the frequency set in the frequency, for example if scheduleFrequency is 2, the copy occurs every 2 hours. <br>'Days' means that the Snapshot copy occurs daily at the frequency set in the frequency. <br>'Weeks' means that the Snapshot copy occurs weekly at the frequency set in the frequency. <br>'Months' means that the Snapshot copy occurs monthly at the frequency set in the Frequency. <br>'Years' means that the Snapshot copy occurs yearly at the frequency set in the scheduleFrequency.
				  * Constraints: Allowable values are: `Runs`, `Hours`, `Days`, `Weeks`, `Months`, `Years`.
			* `target_type` - (Required, String) Specifies the type of target to which replication need to be performed.
			  * Constraints: Allowable values are: `RemoteCluster`, `AWS`, `Azure`.
		* `rpaas_targets` - (Optional, List)
		Nested schema for **rpaas_targets**:
			* `backup_run_type` - (Optional, String) Specifies which type of run should be copied, if not set, all types of runs will be eligible for copying. If set, this will ensure that the first run of given type in the scheduled period will get copied. Currently, this can only be set to Full.
			  * Constraints: Allowable values are: `Regular`, `Full`, `Log`, `System`, `StorageArraySnapshot`.
			* `config_id` - (Optional, String) Specifies the unique identifier for the target getting added. This field need to be passed only when policies are being updated.
			* `copy_on_run_success` - (Optional, Boolean) Specifies if Snapshots are copied from the first completely successful Protection Group Run or the first partially successful Protection Group Run occurring at the start of the replication schedule. <br> If true, Snapshots are copied from the first Protection Group Run occurring at the start of the replication schedule that was completely successful i.e. Snapshots for all the Objects in the Protection Group were successfully captured. <br> If false, Snapshots are copied from the first Protection Group Run occurring at the start of the replication schedule, even if first Protection Group Run was not completely successful i.e. Snapshots were not captured for all Objects in the Protection Group.
			* `log_retention` - (Optional, List) Specifies the retention of a backup.
			Nested schema for **log_retention**:
				* `data_lock_config` - (Optional, List) Specifies WORM retention type for the snapshots. When a WORM retention type is specified, the snapshots of the Protection Groups using this policy will be kept for the last N days as specified in the duration of the datalock. During that time, the snapshots cannot be deleted.
				Nested schema for **data_lock_config**:
					* `duration` - (Required, Integer) Specifies the duration for a dataLock. <br> Example. If duration is 7 and unit is Months, the dataLock is enabled for last 7 * 30 = 210 days of the backup.
					  * Constraints: The minimum value is `1`.
					* `enable_worm_on_external_target` - (Optional, Boolean) Specifies whether objects in the external target associated with this policy need to be made immutable.
					* `mode` - (Required, String) Specifies the type of WORM retention type. 'Compliance' implies WORM retention is set for compliance reason. 'Administrative' implies WORM retention is set for administrative purposes.
					  * Constraints: Allowable values are: `Compliance`, `Administrative`.
					* `unit` - (Required, String) Specificies the Retention Unit of a dataLock measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Months' then number of retention days will be 365 * 2 = 730 days.
					  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
				* `duration` - (Required, Integer) Specifies the duration for a backup retention. <br> Example. If duration is 7 and unit is Months, the retention of a backup is 7 * 30 = 210 days.
				  * Constraints: The minimum value is `0`.
				* `unit` - (Required, String) Specificies the Retention Unit of a backup measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Years' then number of retention days will be 365 * 2 = 730 days.
				  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
			* `retention` - (Required, List) Specifies the retention of a backup.
			Nested schema for **retention**:
				* `data_lock_config` - (Optional, List) Specifies WORM retention type for the snapshots. When a WORM retention type is specified, the snapshots of the Protection Groups using this policy will be kept for the last N days as specified in the duration of the datalock. During that time, the snapshots cannot be deleted.
				Nested schema for **data_lock_config**:
					* `duration` - (Required, Integer) Specifies the duration for a dataLock. <br> Example. If duration is 7 and unit is Months, the dataLock is enabled for last 7 * 30 = 210 days of the backup.
					  * Constraints: The minimum value is `1`.
					* `enable_worm_on_external_target` - (Optional, Boolean) Specifies whether objects in the external target associated with this policy need to be made immutable.
					* `mode` - (Required, String) Specifies the type of WORM retention type. 'Compliance' implies WORM retention is set for compliance reason. 'Administrative' implies WORM retention is set for administrative purposes.
					  * Constraints: Allowable values are: `Compliance`, `Administrative`.
					* `unit` - (Required, String) Specificies the Retention Unit of a dataLock measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Months' then number of retention days will be 365 * 2 = 730 days.
					  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
				* `duration` - (Required, Integer) Specifies the duration for a backup retention. <br> Example. If duration is 7 and unit is Months, the retention of a backup is 7 * 30 = 210 days.
				  * Constraints: The minimum value is `1`.
				* `unit` - (Required, String) Specificies the Retention Unit of a backup measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Years' then number of retention days will be 365 * 2 = 730 days.
				  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
			* `run_timeouts` - (Optional, List) Specifies the replication/archival timeouts for different type of runs(kFull, kRegular etc.).
			Nested schema for **run_timeouts**:
				* `backup_type` - (Optional, String) The scheduled backup type(kFull, kRegular etc.).
				  * Constraints: Allowable values are: `kRegular`, `kFull`, `kLog`, `kSystem`, `kHydrateCDP`, `kStorageArraySnapshot`.
				* `timeout_mins` - (Optional, Integer) Specifies the timeout in mins.
			* `schedule` - (Required, List) Specifies a schedule fregquency and schedule unit for copying Snapshots to backup targets.
			Nested schema for **schedule**:
				* `frequency` - (Optional, Integer) Specifies a factor to multiply the unit by, to determine the copy schedule. For example if set to 2 and the unit is hourly, then Snapshots from the first eligible Job Run for every 2 hour period is copied.
				  * Constraints: The minimum value is `1`.
				* `unit` - (Required, String) Specifies the frequency that Snapshots should be copied to the specified target. Used in combination with multiplier. <br>'Runs' means that the Snapshot copy occurs after the number of Protection Group Runs equals the number specified in the frequency. <br>'Hours' means that the Snapshot copy occurs hourly at the frequency set in the frequency, for example if scheduleFrequency is 2, the copy occurs every 2 hours. <br>'Days' means that the Snapshot copy occurs daily at the frequency set in the frequency. <br>'Weeks' means that the Snapshot copy occurs weekly at the frequency set in the frequency. <br>'Months' means that the Snapshot copy occurs monthly at the frequency set in the Frequency. <br>'Years' means that the Snapshot copy occurs yearly at the frequency set in the scheduleFrequency.
				  * Constraints: Allowable values are: `Runs`, `Hours`, `Days`, `Weeks`, `Months`, `Years`.
			* `target_id` - (Required, Integer) Specifies the RPaaS target to copy the Snapshots.
			* `target_name` - (Computed, String) Specifies the RPaaS target name where Snapshots are copied.
			* `target_type` - (Optional, String) Specifies the RPaaS target type where Snapshots are copied.
			  * Constraints: Allowable values are: `Tape`, `Cloud`, `Nas`.
	* `source_cluster_id` - (Required, Integer) Specifies the source cluster id from where the remote operations will be performed to the next set of remote targets.
* `data_lock` - (Optional, String) This field is now deprecated. Please use the DataLockConfig in the backup retention.
  * Constraints: Allowable values are: `Compliance`, `Administrative`.
* `description` - (Optional, String) Specifies the description of the Protection Policy.
* `endpoint_type` - (Optional, String) Backup Recovery Endpoint type. By default set to "public".
* `instance_id` - (Optional, String) Backup Recovery instance ID. If provided here along with region, the provider constructs the endpoint URL using them, which overrides any value set through environment variables or the `endpoints.json` file.
* `region` - (Optional, String) Backup Recovery region. If provided here along with instance_id, the provider constructs the endpoint URL using them, which overrides any value set through environment variables or the `endpoints.json` file.  
* `extended_retention` - (Optional, List) Specifies additional retention policies that should be applied to the backup snapshots. A backup snapshot will be retained up to a time that is the maximum of all retention policies that are applicable to it.
Nested schema for **extended_retention**:
	* `config_id` - (Optional, String) Specifies the unique identifier for the target getting added. This field need to be passed olny when policies are updated.
	* `retention` - (Required, List) Specifies the retention of a backup.
	Nested schema for **retention**:
		* `data_lock_config` - (Optional, List) Specifies WORM retention type for the snapshots. When a WORM retention type is specified, the snapshots of the Protection Groups using this policy will be kept for the last N days as specified in the duration of the datalock. During that time, the snapshots cannot be deleted.
		Nested schema for **data_lock_config**:
			* `duration` - (Required, Integer) Specifies the duration for a dataLock. <br> Example. If duration is 7 and unit is Months, the dataLock is enabled for last 7 * 30 = 210 days of the backup.
			  * Constraints: The minimum value is `1`.
			* `enable_worm_on_external_target` - (Optional, Boolean) Specifies whether objects in the external target associated with this policy need to be made immutable.
			* `mode` - (Required, String) Specifies the type of WORM retention type. 'Compliance' implies WORM retention is set for compliance reason. 'Administrative' implies WORM retention is set for administrative purposes.
			  * Constraints: Allowable values are: `Compliance`, `Administrative`.
			* `unit` - (Required, String) Specificies the Retention Unit of a dataLock measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Months' then number of retention days will be 365 * 2 = 730 days.
			  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
		* `duration` - (Required, Integer) Specifies the duration for a backup retention. <br> Example. If duration is 7 and unit is Months, the retention of a backup is 7 * 30 = 210 days.
		  * Constraints: The minimum value is `1`.
		* `unit` - (Required, String) Specificies the Retention Unit of a backup measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Years' then number of retention days will be 365 * 2 = 730 days.
		  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
	* `run_type` - (Optional, String) The backup run type to which this extended retention applies to. If this is not set, the extended retention will be applicable to all non-log backup types. Currently, the only value that can be set here is Full.'Regular' indicates a incremental (CBT) backup. Incremental backups utilizing CBT (if supported) are captured of the target protection objects. The first run of a Regular schedule captures all the blocks.'Full' indicates a full (no CBT) backup. A complete backup (all blocks) of the target protection objects are always captured and Change Block Tracking (CBT) is not utilized.'Log' indicates a Database Log backup. Capture the database transaction logs to allow rolling back to a specific point in time.'System' indicates a system backup. System backups are used to do bare metal recovery of the system to a specific point in time.
	  * Constraints: Allowable values are: `Regular`, `Full`, `Log`, `System`, `StorageArraySnapshot`.
	* `schedule` - (Required, List) Specifies a schedule frequency and schedule unit for Extended Retentions.
	Nested schema for **schedule**:
		* `frequency` - (Optional, Integer) Specifies a factor to multiply the unit by, to determine the retention schedule. For example if set to 2 and the unit is hourly, then Snapshots from the first eligible Job Run for every 2 hour period is retained.
		  * Constraints: The minimum value is `1`.
		* `unit` - (Required, String) Specifies the unit interval for retention of Snapshots. <br>'Runs' means that the Snapshot copy retained after the number of Protection Group Runs equals the number specified in the frequency. <br>'Hours' means that the Snapshot copy retained hourly at the frequency set in the frequency, for example if scheduleFrequency is 2, the copy occurs every 2 hours. <br>'Days' means that the Snapshot copy gets retained daily at the frequency set in the frequency. <br>'Weeks' means that the Snapshot copy is retained weekly at the frequency set in the frequency. <br>'Months' means that the Snapshot copy is retained monthly at the frequency set in the Frequency. <br>'Years' means that the Snapshot copy is retained yearly at the frequency set in the Frequency.
		  * Constraints: Allowable values are: `Runs`, `Hours`, `Days`, `Weeks`, `Months`, `Years`.
* `is_cbs_enabled` - (Optional, Boolean) Specifies true if Calender Based Schedule is supported by client. Default value is assumed as false for this feature.
* `last_modification_time_usecs` - (Optional, Integer) Specifies the last time this Policy was updated. If this is passed into a PUT request, then the backend will validate that the timestamp passed in matches the time that the policy was actually last modified. If the two timestamps do not match, then the request will be rejected with a stale error.
* `name` - (Required, String) Specifies the name of the Protection Policy.
* `remote_target_policy` - (Optional, List) Specifies the replication, archival and cloud spin targets of Protection Policy.
Nested schema for **remote_target_policy**:
	* `archival_targets` - (Optional, List)
	Nested schema for **archival_targets**:
		* `backup_run_type` - (Optional, String) Specifies which type of run should be copied, if not set, all types of runs will be eligible for copying. If set, this will ensure that the first run of given type in the scheduled period will get copied. Currently, this can only be set to Full.
		  * Constraints: Allowable values are: `Regular`, `Full`, `Log`, `System`, `StorageArraySnapshot`.
		* `config_id` - (Optional, String) Specifies the unique identifier for the target getting added. This field need to be passed only when policies are being updated.
		* `copy_on_run_success` - (Optional, Boolean) Specifies if Snapshots are copied from the first completely successful Protection Group Run or the first partially successful Protection Group Run occurring at the start of the replication schedule. <br> If true, Snapshots are copied from the first Protection Group Run occurring at the start of the replication schedule that was completely successful i.e. Snapshots for all the Objects in the Protection Group were successfully captured. <br> If false, Snapshots are copied from the first Protection Group Run occurring at the start of the replication schedule, even if first Protection Group Run was not completely successful i.e. Snapshots were not captured for all Objects in the Protection Group.
		* `extended_retention` - (Optional, List) Specifies additional retention policies that should be applied to the archived backup. Archived backup snapshot will be retained up to a time that is the maximum of all retention policies that are applicable to it.
		Nested schema for **extended_retention**:
			* `config_id` - (Optional, String) Specifies the unique identifier for the target getting added. This field need to be passed olny when policies are updated.
			* `retention` - (Required, List) Specifies the retention of a backup.
			Nested schema for **retention**:
				* `data_lock_config` - (Optional, List) Specifies WORM retention type for the snapshots. When a WORM retention type is specified, the snapshots of the Protection Groups using this policy will be kept for the last N days as specified in the duration of the datalock. During that time, the snapshots cannot be deleted.
				Nested schema for **data_lock_config**:
					* `duration` - (Required, Integer) Specifies the duration for a dataLock. <br> Example. If duration is 7 and unit is Months, the dataLock is enabled for last 7 * 30 = 210 days of the backup.
					  * Constraints: The minimum value is `1`.
					* `enable_worm_on_external_target` - (Optional, Boolean) Specifies whether objects in the external target associated with this policy need to be made immutable.
					* `mode` - (Required, String) Specifies the type of WORM retention type. 'Compliance' implies WORM retention is set for compliance reason. 'Administrative' implies WORM retention is set for administrative purposes.
					  * Constraints: Allowable values are: `Compliance`, `Administrative`.
					* `unit` - (Required, String) Specificies the Retention Unit of a dataLock measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Months' then number of retention days will be 365 * 2 = 730 days.
					  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
				* `duration` - (Required, Integer) Specifies the duration for a backup retention. <br> Example. If duration is 7 and unit is Months, the retention of a backup is 7 * 30 = 210 days.
				  * Constraints: The minimum value is `1`.
				* `unit` - (Required, String) Specificies the Retention Unit of a backup measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Years' then number of retention days will be 365 * 2 = 730 days.
				  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
			* `run_type` - (Optional, String) The backup run type to which this extended retention applies to. If this is not set, the extended retention will be applicable to all non-log backup types. Currently, the only value that can be set here is Full.'Regular' indicates a incremental (CBT) backup. Incremental backups utilizing CBT (if supported) are captured of the target protection objects. The first run of a Regular schedule captures all the blocks.'Full' indicates a full (no CBT) backup. A complete backup (all blocks) of the target protection objects are always captured and Change Block Tracking (CBT) is not utilized.'Log' indicates a Database Log backup. Capture the database transaction logs to allow rolling back to a specific point in time.'System' indicates a system backup. System backups are used to do bare metal recovery of the system to a specific point in time.
			  * Constraints: Allowable values are: `Regular`, `Full`, `Log`, `System`, `StorageArraySnapshot`.
			* `schedule` - (Required, List) Specifies a schedule frequency and schedule unit for Extended Retentions.
			Nested schema for **schedule**:
				* `frequency` - (Optional, Integer) Specifies a factor to multiply the unit by, to determine the retention schedule. For example if set to 2 and the unit is hourly, then Snapshots from the first eligible Job Run for every 2 hour period is retained.
				  * Constraints: The minimum value is `1`.
				* `unit` - (Required, String) Specifies the unit interval for retention of Snapshots. <br>'Runs' means that the Snapshot copy retained after the number of Protection Group Runs equals the number specified in the frequency. <br>'Hours' means that the Snapshot copy retained hourly at the frequency set in the frequency, for example if scheduleFrequency is 2, the copy occurs every 2 hours. <br>'Days' means that the Snapshot copy gets retained daily at the frequency set in the frequency. <br>'Weeks' means that the Snapshot copy is retained weekly at the frequency set in the frequency. <br>'Months' means that the Snapshot copy is retained monthly at the frequency set in the Frequency. <br>'Years' means that the Snapshot copy is retained yearly at the frequency set in the Frequency.
				  * Constraints: Allowable values are: `Runs`, `Hours`, `Days`, `Weeks`, `Months`, `Years`.
		* `log_retention` - (Optional, List) Specifies the retention of a backup.
		Nested schema for **log_retention**:
			* `data_lock_config` - (Optional, List) Specifies WORM retention type for the snapshots. When a WORM retention type is specified, the snapshots of the Protection Groups using this policy will be kept for the last N days as specified in the duration of the datalock. During that time, the snapshots cannot be deleted.
			Nested schema for **data_lock_config**:
				* `duration` - (Required, Integer) Specifies the duration for a dataLock. <br> Example. If duration is 7 and unit is Months, the dataLock is enabled for last 7 * 30 = 210 days of the backup.
				  * Constraints: The minimum value is `1`.
				* `enable_worm_on_external_target` - (Optional, Boolean) Specifies whether objects in the external target associated with this policy need to be made immutable.
				* `mode` - (Required, String) Specifies the type of WORM retention type. 'Compliance' implies WORM retention is set for compliance reason. 'Administrative' implies WORM retention is set for administrative purposes.
				  * Constraints: Allowable values are: `Compliance`, `Administrative`.
				* `unit` - (Required, String) Specificies the Retention Unit of a dataLock measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Months' then number of retention days will be 365 * 2 = 730 days.
				  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
			* `duration` - (Required, Integer) Specifies the duration for a backup retention. <br> Example. If duration is 7 and unit is Months, the retention of a backup is 7 * 30 = 210 days.
			  * Constraints: The minimum value is `0`.
			* `unit` - (Required, String) Specificies the Retention Unit of a backup measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Years' then number of retention days will be 365 * 2 = 730 days.
			  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
		* `retention` - (Required, List) Specifies the retention of a backup.
		Nested schema for **retention**:
			* `data_lock_config` - (Optional, List) Specifies WORM retention type for the snapshots. When a WORM retention type is specified, the snapshots of the Protection Groups using this policy will be kept for the last N days as specified in the duration of the datalock. During that time, the snapshots cannot be deleted.
			Nested schema for **data_lock_config**:
				* `duration` - (Required, Integer) Specifies the duration for a dataLock. <br> Example. If duration is 7 and unit is Months, the dataLock is enabled for last 7 * 30 = 210 days of the backup.
				  * Constraints: The minimum value is `1`.
				* `enable_worm_on_external_target` - (Optional, Boolean) Specifies whether objects in the external target associated with this policy need to be made immutable.
				* `mode` - (Required, String) Specifies the type of WORM retention type. 'Compliance' implies WORM retention is set for compliance reason. 'Administrative' implies WORM retention is set for administrative purposes.
				  * Constraints: Allowable values are: `Compliance`, `Administrative`.
				* `unit` - (Required, String) Specificies the Retention Unit of a dataLock measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Months' then number of retention days will be 365 * 2 = 730 days.
				  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
			* `duration` - (Required, Integer) Specifies the duration for a backup retention. <br> Example. If duration is 7 and unit is Months, the retention of a backup is 7 * 30 = 210 days.
			  * Constraints: The minimum value is `1`.
			* `unit` - (Required, String) Specificies the Retention Unit of a backup measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Years' then number of retention days will be 365 * 2 = 730 days.
			  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
		* `run_timeouts` - (Optional, List) Specifies the replication/archival timeouts for different type of runs(kFull, kRegular etc.).
		Nested schema for **run_timeouts**:
			* `backup_type` - (Optional, String) The scheduled backup type(kFull, kRegular etc.).
			  * Constraints: Allowable values are: `kRegular`, `kFull`, `kLog`, `kSystem`, `kHydrateCDP`, `kStorageArraySnapshot`.
			* `timeout_mins` - (Optional, Integer) Specifies the timeout in mins.
		* `schedule` - (Required, List) Specifies a schedule fregquency and schedule unit for copying Snapshots to backup targets.
		Nested schema for **schedule**:
			* `frequency` - (Optional, Integer) Specifies a factor to multiply the unit by, to determine the copy schedule. For example if set to 2 and the unit is hourly, then Snapshots from the first eligible Job Run for every 2 hour period is copied.
			  * Constraints: The minimum value is `1`.
			* `unit` - (Required, String) Specifies the frequency that Snapshots should be copied to the specified target. Used in combination with multiplier. <br>'Runs' means that the Snapshot copy occurs after the number of Protection Group Runs equals the number specified in the frequency. <br>'Hours' means that the Snapshot copy occurs hourly at the frequency set in the frequency, for example if scheduleFrequency is 2, the copy occurs every 2 hours. <br>'Days' means that the Snapshot copy occurs daily at the frequency set in the frequency. <br>'Weeks' means that the Snapshot copy occurs weekly at the frequency set in the frequency. <br>'Months' means that the Snapshot copy occurs monthly at the frequency set in the Frequency. <br>'Years' means that the Snapshot copy occurs yearly at the frequency set in the scheduleFrequency.
			  * Constraints: Allowable values are: `Runs`, `Hours`, `Days`, `Weeks`, `Months`, `Years`.
		* `target_id` - (Required, Integer) Specifies the Archival target to copy the Snapshots to.
		* `target_name` - (Computed, String) Specifies the Archival target name where Snapshots are copied.
		* `target_type` - (Computed, String) Specifies the Archival target type where Snapshots are copied.
		  * Constraints: Allowable values are: `Tape`, `Cloud`, `Nas`.
		* `tier_settings` - (Optional, List) Specifies the settings tier levels configured with each archival target. The tier settings need to be applied in specific order and default tier should always be passed as first entry in tiers array. The following example illustrates how to configure tiering input for AWS tiering. Same type of input structure applied to other cloud platforms also. <br>If user wants to achieve following tiering for backup, <br>User Desired Tiering- <br><t>1.Archive Full back up for 12 Months <br><t>2.Tier Levels <br><t><t>[1,12] [ <br><t><t><t>s3 (1 to 2 months), (default tier) <br><t><t><t>s3 Intelligent tiering (3 to 6 months), <br><t><t><t>s3 One Zone (7 to 9 months) <br><t><t><t>Glacier (10 to 12 months)] <br><t>API Input <br><t><t>1.tiers-[ <br><t><t><t>{'tierType': 'S3','moveAfterUnit':'months', <br><t><t><t>'moveAfter':2 - move from s3 to s3Inte after 2 months}, <br><t><t><t>{'tierType': 'S3Inte','moveAfterUnit':'months', <br><t><t><t>'moveAfter':4 - move from S3Inte to Glacier after 4 months}, <br><t><t><t>{'tierType': 'Glacier', 'moveAfterUnit':'months', <br><t><t><t>'moveAfter': 3 - move from Glacier to S3 One Zone after 3 months }, <br><t><t><t>{'tierType': 'S3 One Zone', 'moveAfterUnit': nil, <br><t><t><t>'moveAfter': nil - For the last record, 'moveAfter' and 'moveAfterUnit' <br><t><t><t>will be ignored since there are no further tier for data movement } <br><t><t><t>}].
		Nested schema for **tier_settings**:
			* `aws_tiering` - (Optional, List) Specifies aws tiers.
			Nested schema for **aws_tiering**:
				* `tiers` - (Required, List) Specifies the tiers that are used to move the archived backup from current tier to next tier. The order of the tiers determines which tier will be used next for moving the archived backup. The first tier input should always be default tier where backup will be acrhived. Each tier specifies how much time after the backup will be moved to next tier from the current tier.
				Nested schema for **tiers**:
					* `move_after` - (Computed, Integer) Specifies the time period after which the backup will be moved from current tier to next tier.
					* `move_after_unit` - (Computed, String) Specifies the unit for moving the data from current tier to next tier. This unit will be a base unit for the 'moveAfter' field specified below.
					  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
					* `tier_type` - (Computed, String) Specifies the AWS tier types.
					  * Constraints: Allowable values are: `kAmazonS3Standard`, `kAmazonS3StandardIA`, `kAmazonS3OneZoneIA`, `kAmazonS3IntelligentTiering`, `kAmazonS3Glacier`, `kAmazonS3GlacierDeepArchive`.
			* `azure_tiering` - (Optional, List) Specifies Azure tiers.
			Nested schema for **azure_tiering**:
				* `tiers` - (Optional, List) Specifies the tiers that are used to move the archived backup from current tier to next tier. The order of the tiers determines which tier will be used next for moving the archived backup. The first tier input should always be default tier where backup will be acrhived. Each tier specifies how much time after the backup will be moved to next tier from the current tier.
				Nested schema for **tiers**:
					* `move_after` - (Computed, Integer) Specifies the time period after which the backup will be moved from current tier to next tier.
					* `move_after_unit` - (Computed, String) Specifies the unit for moving the data from current tier to next tier. This unit will be a base unit for the 'moveAfter' field specified below.
					  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
					* `tier_type` - (Computed, String) Specifies the Azure tier types.
					  * Constraints: Allowable values are: `kAzureTierHot`, `kAzureTierCool`, `kAzureTierArchive`.
			* `cloud_platform` - (Optional, String) Specifies the cloud platform to enable tiering.
			  * Constraints: Allowable values are: `AWS`, `Azure`, `Oracle`, `Google`.
			* `google_tiering` - (Optional, List) Specifies Google tiers.
			Nested schema for **google_tiering**:
				* `tiers` - (Required, List) Specifies the tiers that are used to move the archived backup from current tier to next tier. The order of the tiers determines which tier will be used next for moving the archived backup. The first tier input should always be default tier where backup will be acrhived. Each tier specifies how much time after the backup will be moved to next tier from the current tier.
				Nested schema for **tiers**:
					* `move_after` - (Computed, Integer) Specifies the time period after which the backup will be moved from current tier to next tier.
					* `move_after_unit` - (Computed, String) Specifies the unit for moving the data from current tier to next tier. This unit will be a base unit for the 'moveAfter' field specified below.
					  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
					* `tier_type` - (Computed, String) Specifies the Google tier types.
					  * Constraints: Allowable values are: `kGoogleStandard`, `kGoogleRegional`, `kGoogleMultiRegional`, `kGoogleNearline`, `kGoogleColdline`.
			* `oracle_tiering` - (Optional, List) Specifies Oracle tiers.
			Nested schema for **oracle_tiering**:
				* `tiers` - (Required, List) Specifies the tiers that are used to move the archived backup from current tier to next tier. The order of the tiers determines which tier will be used next for moving the archived backup. The first tier input should always be default tier where backup will be acrhived. Each tier specifies how much time after the backup will be moved to next tier from the current tier.
				Nested schema for **tiers**:
					* `move_after` - (Computed, Integer) Specifies the time period after which the backup will be moved from current tier to next tier.
					* `move_after_unit` - (Computed, String) Specifies the unit for moving the data from current tier to next tier. This unit will be a base unit for the 'moveAfter' field specified below.
					  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
					* `tier_type` - (Computed, String) Specifies the Oracle tier types.
					  * Constraints: Allowable values are: `kOracleTierStandard`, `kOracleTierArchive`.
	* `cloud_spin_targets` - (Optional, List)
	Nested schema for **cloud_spin_targets**:
		* `backup_run_type` - (Optional, String) Specifies which type of run should be copied, if not set, all types of runs will be eligible for copying. If set, this will ensure that the first run of given type in the scheduled period will get copied. Currently, this can only be set to Full.
		  * Constraints: Allowable values are: `Regular`, `Full`, `Log`, `System`, `StorageArraySnapshot`.
		* `config_id` - (Optional, String) Specifies the unique identifier for the target getting added. This field need to be passed only when policies are being updated.
		* `copy_on_run_success` - (Optional, Boolean) Specifies if Snapshots are copied from the first completely successful Protection Group Run or the first partially successful Protection Group Run occurring at the start of the replication schedule. <br> If true, Snapshots are copied from the first Protection Group Run occurring at the start of the replication schedule that was completely successful i.e. Snapshots for all the Objects in the Protection Group were successfully captured. <br> If false, Snapshots are copied from the first Protection Group Run occurring at the start of the replication schedule, even if first Protection Group Run was not completely successful i.e. Snapshots were not captured for all Objects in the Protection Group.
		* `log_retention` - (Optional, List) Specifies the retention of a backup.
		Nested schema for **log_retention**:
			* `data_lock_config` - (Optional, List) Specifies WORM retention type for the snapshots. When a WORM retention type is specified, the snapshots of the Protection Groups using this policy will be kept for the last N days as specified in the duration of the datalock. During that time, the snapshots cannot be deleted.
			Nested schema for **data_lock_config**:
				* `duration` - (Required, Integer) Specifies the duration for a dataLock. <br> Example. If duration is 7 and unit is Months, the dataLock is enabled for last 7 * 30 = 210 days of the backup.
				  * Constraints: The minimum value is `1`.
				* `enable_worm_on_external_target` - (Optional, Boolean) Specifies whether objects in the external target associated with this policy need to be made immutable.
				* `mode` - (Required, String) Specifies the type of WORM retention type. 'Compliance' implies WORM retention is set for compliance reason. 'Administrative' implies WORM retention is set for administrative purposes.
				  * Constraints: Allowable values are: `Compliance`, `Administrative`.
				* `unit` - (Required, String) Specificies the Retention Unit of a dataLock measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Months' then number of retention days will be 365 * 2 = 730 days.
				  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
			* `duration` - (Required, Integer) Specifies the duration for a backup retention. <br> Example. If duration is 7 and unit is Months, the retention of a backup is 7 * 30 = 210 days.
			  * Constraints: The minimum value is `0`.
			* `unit` - (Required, String) Specificies the Retention Unit of a backup measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Years' then number of retention days will be 365 * 2 = 730 days.
			  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
		* `retention` - (Required, List) Specifies the retention of a backup.
		Nested schema for **retention**:
			* `data_lock_config` - (Optional, List) Specifies WORM retention type for the snapshots. When a WORM retention type is specified, the snapshots of the Protection Groups using this policy will be kept for the last N days as specified in the duration of the datalock. During that time, the snapshots cannot be deleted.
			Nested schema for **data_lock_config**:
				* `duration` - (Required, Integer) Specifies the duration for a dataLock. <br> Example. If duration is 7 and unit is Months, the dataLock is enabled for last 7 * 30 = 210 days of the backup.
				  * Constraints: The minimum value is `1`.
				* `enable_worm_on_external_target` - (Optional, Boolean) Specifies whether objects in the external target associated with this policy need to be made immutable.
				* `mode` - (Required, String) Specifies the type of WORM retention type. 'Compliance' implies WORM retention is set for compliance reason. 'Administrative' implies WORM retention is set for administrative purposes.
				  * Constraints: Allowable values are: `Compliance`, `Administrative`.
				* `unit` - (Required, String) Specificies the Retention Unit of a dataLock measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Months' then number of retention days will be 365 * 2 = 730 days.
				  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
			* `duration` - (Required, Integer) Specifies the duration for a backup retention. <br> Example. If duration is 7 and unit is Months, the retention of a backup is 7 * 30 = 210 days.
			  * Constraints: The minimum value is `1`.
			* `unit` - (Required, String) Specificies the Retention Unit of a backup measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Years' then number of retention days will be 365 * 2 = 730 days.
			  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
		* `run_timeouts` - (Optional, List) Specifies the replication/archival timeouts for different type of runs(kFull, kRegular etc.).
		Nested schema for **run_timeouts**:
			* `backup_type` - (Optional, String) The scheduled backup type(kFull, kRegular etc.).
			  * Constraints: Allowable values are: `kRegular`, `kFull`, `kLog`, `kSystem`, `kHydrateCDP`, `kStorageArraySnapshot`.
			* `timeout_mins` - (Optional, Integer) Specifies the timeout in mins.
		* `schedule` - (Required, List) Specifies a schedule fregquency and schedule unit for copying Snapshots to backup targets.
		Nested schema for **schedule**:
			* `frequency` - (Optional, Integer) Specifies a factor to multiply the unit by, to determine the copy schedule. For example if set to 2 and the unit is hourly, then Snapshots from the first eligible Job Run for every 2 hour period is copied.
			  * Constraints: The minimum value is `1`.
			* `unit` - (Required, String) Specifies the frequency that Snapshots should be copied to the specified target. Used in combination with multiplier. <br>'Runs' means that the Snapshot copy occurs after the number of Protection Group Runs equals the number specified in the frequency. <br>'Hours' means that the Snapshot copy occurs hourly at the frequency set in the frequency, for example if scheduleFrequency is 2, the copy occurs every 2 hours. <br>'Days' means that the Snapshot copy occurs daily at the frequency set in the frequency. <br>'Weeks' means that the Snapshot copy occurs weekly at the frequency set in the frequency. <br>'Months' means that the Snapshot copy occurs monthly at the frequency set in the Frequency. <br>'Years' means that the Snapshot copy occurs yearly at the frequency set in the scheduleFrequency.
			  * Constraints: Allowable values are: `Runs`, `Hours`, `Days`, `Weeks`, `Months`, `Years`.
		* `target` - (Required, List) Specifies the details about Cloud Spin target where backup snapshots may be converted and stored.
		Nested schema for **target**:
			* `aws_params` - (Optional, List) Specifies various resources when converting and deploying a VM to AWS.
			Nested schema for **aws_params**:
				* `custom_tag_list` - (Optional, List) Specifies tags of various resources when converting and deploying a VM to AWS.
				Nested schema for **custom_tag_list**:
					* `key` - (Computed, String) Specifies key of the custom tag.
					* `value` - (Computed, String) Specifies value of the custom tag.
				* `region` - (Required, Integer) Specifies id of the AWS region in which to deploy the VM.
				* `subnet_id` - (Optional, Integer) Specifies id of the subnet within above VPC.
				* `vpc_id` - (Optional, Integer) Specifies id of the Virtual Private Cloud to chose for the instance type.
			* `azure_params` - (Optional, List) Specifies various resources when converting and deploying a VM to Azure.
			Nested schema for **azure_params**:
				* `availability_set_id` - (Optional, Integer) Specifies the availability set.
				* `network_resource_group_id` - (Optional, Integer) Specifies id of the resource group for the selected virtual network.
				* `resource_group_id` - (Optional, Integer) Specifies id of the Azure resource group. Its value is globally unique within Azure.
				* `storage_account_id` - (Optional, Integer) Specifies id of the storage account that will contain the storage container within which we will create the blob that will become the VHD disk for the cloned VM.
				* `storage_container_id` - (Optional, Integer) Specifies id of the storage container within the above storage account.
				* `storage_resource_group_id` - (Optional, Integer) Specifies id of the resource group for the selected storage account.
				* `temp_vm_resource_group_id` - (Optional, Integer) Specifies id of the temporary Azure resource group.
				* `temp_vm_storage_account_id` - (Optional, Integer) Specifies id of the temporary VM storage account that will contain the storage container within which we will create the blob that will become the VHD disk for the cloned VM.
				* `temp_vm_storage_container_id` - (Optional, Integer) Specifies id of the temporary VM storage container within the above storage account.
				* `temp_vm_subnet_id` - (Optional, Integer) Specifies Id of the temporary VM subnet within the above virtual network.
				* `temp_vm_virtual_network_id` - (Optional, Integer) Specifies Id of the temporary VM Virtual Network.
			* `id` - (Optional, Integer) Specifies the unique id of the cloud spin entity.
			* `name` - (Computed, String) Specifies the name of the already added cloud spin target.
	* `onprem_deploy_targets` - (Optional, List)
	Nested schema for **onprem_deploy_targets**:
		* `backup_run_type` - (Optional, String) Specifies which type of run should be copied, if not set, all types of runs will be eligible for copying. If set, this will ensure that the first run of given type in the scheduled period will get copied. Currently, this can only be set to Full.
		  * Constraints: Allowable values are: `Regular`, `Full`, `Log`, `System`, `StorageArraySnapshot`.
		* `config_id` - (Optional, String) Specifies the unique identifier for the target getting added. This field need to be passed only when policies are being updated.
		* `copy_on_run_success` - (Optional, Boolean) Specifies if Snapshots are copied from the first completely successful Protection Group Run or the first partially successful Protection Group Run occurring at the start of the replication schedule. <br> If true, Snapshots are copied from the first Protection Group Run occurring at the start of the replication schedule that was completely successful i.e. Snapshots for all the Objects in the Protection Group were successfully captured. <br> If false, Snapshots are copied from the first Protection Group Run occurring at the start of the replication schedule, even if first Protection Group Run was not completely successful i.e. Snapshots were not captured for all Objects in the Protection Group.
		* `log_retention` - (Optional, List) Specifies the retention of a backup.
		Nested schema for **log_retention**:
			* `data_lock_config` - (Optional, List) Specifies WORM retention type for the snapshots. When a WORM retention type is specified, the snapshots of the Protection Groups using this policy will be kept for the last N days as specified in the duration of the datalock. During that time, the snapshots cannot be deleted.
			Nested schema for **data_lock_config**:
				* `duration` - (Required, Integer) Specifies the duration for a dataLock. <br> Example. If duration is 7 and unit is Months, the dataLock is enabled for last 7 * 30 = 210 days of the backup.
				  * Constraints: The minimum value is `1`.
				* `enable_worm_on_external_target` - (Optional, Boolean) Specifies whether objects in the external target associated with this policy need to be made immutable.
				* `mode` - (Required, String) Specifies the type of WORM retention type. 'Compliance' implies WORM retention is set for compliance reason. 'Administrative' implies WORM retention is set for administrative purposes.
				  * Constraints: Allowable values are: `Compliance`, `Administrative`.
				* `unit` - (Required, String) Specificies the Retention Unit of a dataLock measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Months' then number of retention days will be 365 * 2 = 730 days.
				  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
			* `duration` - (Required, Integer) Specifies the duration for a backup retention. <br> Example. If duration is 7 and unit is Months, the retention of a backup is 7 * 30 = 210 days.
			  * Constraints: The minimum value is `0`.
			* `unit` - (Required, String) Specificies the Retention Unit of a backup measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Years' then number of retention days will be 365 * 2 = 730 days.
			  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
		* `params` - (Optional, List) Specifies the details about OnpremDeploy target where backup snapshots may be converted and deployed.
		Nested schema for **params**:
			* `id` - (Optional, Integer) Specifies the unique id of the onprem entity.
		* `retention` - (Required, List) Specifies the retention of a backup.
		Nested schema for **retention**:
			* `data_lock_config` - (Optional, List) Specifies WORM retention type for the snapshots. When a WORM retention type is specified, the snapshots of the Protection Groups using this policy will be kept for the last N days as specified in the duration of the datalock. During that time, the snapshots cannot be deleted.
			Nested schema for **data_lock_config**:
				* `duration` - (Required, Integer) Specifies the duration for a dataLock. <br> Example. If duration is 7 and unit is Months, the dataLock is enabled for last 7 * 30 = 210 days of the backup.
				  * Constraints: The minimum value is `1`.
				* `enable_worm_on_external_target` - (Optional, Boolean) Specifies whether objects in the external target associated with this policy need to be made immutable.
				* `mode` - (Required, String) Specifies the type of WORM retention type. 'Compliance' implies WORM retention is set for compliance reason. 'Administrative' implies WORM retention is set for administrative purposes.
				  * Constraints: Allowable values are: `Compliance`, `Administrative`.
				* `unit` - (Required, String) Specificies the Retention Unit of a dataLock measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Months' then number of retention days will be 365 * 2 = 730 days.
				  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
			* `duration` - (Required, Integer) Specifies the duration for a backup retention. <br> Example. If duration is 7 and unit is Months, the retention of a backup is 7 * 30 = 210 days.
			  * Constraints: The minimum value is `1`.
			* `unit` - (Required, String) Specificies the Retention Unit of a backup measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Years' then number of retention days will be 365 * 2 = 730 days.
			  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
		* `run_timeouts` - (Optional, List) Specifies the replication/archival timeouts for different type of runs(kFull, kRegular etc.).
		Nested schema for **run_timeouts**:
			* `backup_type` - (Optional, String) The scheduled backup type(kFull, kRegular etc.).
			  * Constraints: Allowable values are: `kRegular`, `kFull`, `kLog`, `kSystem`, `kHydrateCDP`, `kStorageArraySnapshot`.
			* `timeout_mins` - (Optional, Integer) Specifies the timeout in mins.
		* `schedule` - (Required, List) Specifies a schedule fregquency and schedule unit for copying Snapshots to backup targets.
		Nested schema for **schedule**:
			* `frequency` - (Optional, Integer) Specifies a factor to multiply the unit by, to determine the copy schedule. For example if set to 2 and the unit is hourly, then Snapshots from the first eligible Job Run for every 2 hour period is copied.
			  * Constraints: The minimum value is `1`.
			* `unit` - (Required, String) Specifies the frequency that Snapshots should be copied to the specified target. Used in combination with multiplier. <br>'Runs' means that the Snapshot copy occurs after the number of Protection Group Runs equals the number specified in the frequency. <br>'Hours' means that the Snapshot copy occurs hourly at the frequency set in the frequency, for example if scheduleFrequency is 2, the copy occurs every 2 hours. <br>'Days' means that the Snapshot copy occurs daily at the frequency set in the frequency. <br>'Weeks' means that the Snapshot copy occurs weekly at the frequency set in the frequency. <br>'Months' means that the Snapshot copy occurs monthly at the frequency set in the Frequency. <br>'Years' means that the Snapshot copy occurs yearly at the frequency set in the scheduleFrequency.
			  * Constraints: Allowable values are: `Runs`, `Hours`, `Days`, `Weeks`, `Months`, `Years`.
	* `replication_targets` - (Optional, List)
	Nested schema for **replication_targets**:
		* `aws_target_config` - (Optional, List) Specifies the configuration for adding AWS as repilcation target.
		Nested schema for **aws_target_config**:
			* `name` - (Computed, String) Specifies the name of the AWS Replication target.
			* `region` - (Required, Integer) Specifies id of the AWS region in which to replicate the Snapshot to. Applicable if replication target is AWS target.
			* `region_name` - (Computed, String) Specifies name of the AWS region in which to replicate the Snapshot to. Applicable if replication target is AWS target.
			* `source_id` - (Required, Integer) Specifies the source id of the AWS protection source registered on IBM cluster.
		* `azure_target_config` - (Optional, List) Specifies the configuration for adding Azure as replication target.
		Nested schema for **azure_target_config**:
			* `name` - (Computed, String) Specifies the name of the Azure Replication target.
			* `resource_group` - (Optional, Integer) Specifies id of the Azure resource group used to filter regions in UI.
			* `resource_group_name` - (Computed, String) Specifies name of the Azure resource group used to filter regions in UI.
			* `source_id` - (Required, Integer) Specifies the source id of the Azure protection source registered on IBM cluster.
			* `storage_account` - (Computed, Integer) Specifies id of the storage account of Azure replication target which will contain storage container.
			* `storage_account_name` - (Computed, String) Specifies name of the storage account of Azure replication target which will contain storage container.
			* `storage_container` - (Computed, Integer) Specifies id of the storage container of Azure Replication target.
			* `storage_container_name` - (Computed, String) Specifies name of the storage container of Azure Replication target.
			* `storage_resource_group` - (Computed, Integer) Specifies id of the storage resource group of Azure Replication target.
			* `storage_resource_group_name` - (Computed, String) Specifies name of the storage resource group of Azure Replication target.
		* `backup_run_type` - (Optional, String) Specifies which type of run should be copied, if not set, all types of runs will be eligible for copying. If set, this will ensure that the first run of given type in the scheduled period will get copied. Currently, this can only be set to Full.
		  * Constraints: Allowable values are: `Regular`, `Full`, `Log`, `System`, `StorageArraySnapshot`.
		* `config_id` - (Optional, String) Specifies the unique identifier for the target getting added. This field need to be passed only when policies are being updated.
		* `copy_on_run_success` - (Optional, Boolean) Specifies if Snapshots are copied from the first completely successful Protection Group Run or the first partially successful Protection Group Run occurring at the start of the replication schedule. <br> If true, Snapshots are copied from the first Protection Group Run occurring at the start of the replication schedule that was completely successful i.e. Snapshots for all the Objects in the Protection Group were successfully captured. <br> If false, Snapshots are copied from the first Protection Group Run occurring at the start of the replication schedule, even if first Protection Group Run was not completely successful i.e. Snapshots were not captured for all Objects in the Protection Group.
		* `log_retention` - (Optional, List) Specifies the retention of a backup.
		Nested schema for **log_retention**:
			* `data_lock_config` - (Optional, List) Specifies WORM retention type for the snapshots. When a WORM retention type is specified, the snapshots of the Protection Groups using this policy will be kept for the last N days as specified in the duration of the datalock. During that time, the snapshots cannot be deleted.
			Nested schema for **data_lock_config**:
				* `duration` - (Required, Integer) Specifies the duration for a dataLock. <br> Example. If duration is 7 and unit is Months, the dataLock is enabled for last 7 * 30 = 210 days of the backup.
				  * Constraints: The minimum value is `1`.
				* `enable_worm_on_external_target` - (Optional, Boolean) Specifies whether objects in the external target associated with this policy need to be made immutable.
				* `mode` - (Required, String) Specifies the type of WORM retention type. 'Compliance' implies WORM retention is set for compliance reason. 'Administrative' implies WORM retention is set for administrative purposes.
				  * Constraints: Allowable values are: `Compliance`, `Administrative`.
				* `unit` - (Required, String) Specificies the Retention Unit of a dataLock measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Months' then number of retention days will be 365 * 2 = 730 days.
				  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
			* `duration` - (Required, Integer) Specifies the duration for a backup retention. <br> Example. If duration is 7 and unit is Months, the retention of a backup is 7 * 30 = 210 days.
			  * Constraints: The minimum value is `0`.
			* `unit` - (Required, String) Specificies the Retention Unit of a backup measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Years' then number of retention days will be 365 * 2 = 730 days.
			  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
		* `remote_target_config` - (Optional, List) Specifies the configuration for adding remote cluster as repilcation target.
		Nested schema for **remote_target_config**:
			* `cluster_id` - (Required, Integer) Specifies the cluster id of the target replication cluster.
			* `cluster_name` - (Computed, String) Specifies the cluster name of the target replication cluster.
		* `retention` - (Required, List) Specifies the retention of a backup.
		Nested schema for **retention**:
			* `data_lock_config` - (Optional, List) Specifies WORM retention type for the snapshots. When a WORM retention type is specified, the snapshots of the Protection Groups using this policy will be kept for the last N days as specified in the duration of the datalock. During that time, the snapshots cannot be deleted.
			Nested schema for **data_lock_config**:
				* `duration` - (Required, Integer) Specifies the duration for a dataLock. <br> Example. If duration is 7 and unit is Months, the dataLock is enabled for last 7 * 30 = 210 days of the backup.
				  * Constraints: The minimum value is `1`.
				* `enable_worm_on_external_target` - (Optional, Boolean) Specifies whether objects in the external target associated with this policy need to be made immutable.
				* `mode` - (Required, String) Specifies the type of WORM retention type. 'Compliance' implies WORM retention is set for compliance reason. 'Administrative' implies WORM retention is set for administrative purposes.
				  * Constraints: Allowable values are: `Compliance`, `Administrative`.
				* `unit` - (Required, String) Specificies the Retention Unit of a dataLock measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Months' then number of retention days will be 365 * 2 = 730 days.
				  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
			* `duration` - (Required, Integer) Specifies the duration for a backup retention. <br> Example. If duration is 7 and unit is Months, the retention of a backup is 7 * 30 = 210 days.
			  * Constraints: The minimum value is `1`.
			* `unit` - (Required, String) Specificies the Retention Unit of a backup measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Years' then number of retention days will be 365 * 2 = 730 days.
			  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
		* `run_timeouts` - (Optional, List) Specifies the replication/archival timeouts for different type of runs(kFull, kRegular etc.).
		Nested schema for **run_timeouts**:
			* `backup_type` - (Optional, String) The scheduled backup type(kFull, kRegular etc.).
			  * Constraints: Allowable values are: `kRegular`, `kFull`, `kLog`, `kSystem`, `kHydrateCDP`, `kStorageArraySnapshot`.
			* `timeout_mins` - (Optional, Integer) Specifies the timeout in mins.
		* `schedule` - (Required, List) Specifies a schedule fregquency and schedule unit for copying Snapshots to backup targets.
		Nested schema for **schedule**:
			* `frequency` - (Optional, Integer) Specifies a factor to multiply the unit by, to determine the copy schedule. For example if set to 2 and the unit is hourly, then Snapshots from the first eligible Job Run for every 2 hour period is copied.
			  * Constraints: The minimum value is `1`.
			* `unit` - (Required, String) Specifies the frequency that Snapshots should be copied to the specified target. Used in combination with multiplier. <br>'Runs' means that the Snapshot copy occurs after the number of Protection Group Runs equals the number specified in the frequency. <br>'Hours' means that the Snapshot copy occurs hourly at the frequency set in the frequency, for example if scheduleFrequency is 2, the copy occurs every 2 hours. <br>'Days' means that the Snapshot copy occurs daily at the frequency set in the frequency. <br>'Weeks' means that the Snapshot copy occurs weekly at the frequency set in the frequency. <br>'Months' means that the Snapshot copy occurs monthly at the frequency set in the Frequency. <br>'Years' means that the Snapshot copy occurs yearly at the frequency set in the scheduleFrequency.
			  * Constraints: Allowable values are: `Runs`, `Hours`, `Days`, `Weeks`, `Months`, `Years`.
		* `target_type` - (Required, String) Specifies the type of target to which replication need to be performed.
		  * Constraints: Allowable values are: `RemoteCluster`, `AWS`, `Azure`.
	* `rpaas_targets` - (Optional, List)
	Nested schema for **rpaas_targets**:
		* `backup_run_type` - (Optional, String) Specifies which type of run should be copied, if not set, all types of runs will be eligible for copying. If set, this will ensure that the first run of given type in the scheduled period will get copied. Currently, this can only be set to Full.
		  * Constraints: Allowable values are: `Regular`, `Full`, `Log`, `System`, `StorageArraySnapshot`.
		* `config_id` - (Optional, String) Specifies the unique identifier for the target getting added. This field need to be passed only when policies are being updated.
		* `copy_on_run_success` - (Optional, Boolean) Specifies if Snapshots are copied from the first completely successful Protection Group Run or the first partially successful Protection Group Run occurring at the start of the replication schedule. <br> If true, Snapshots are copied from the first Protection Group Run occurring at the start of the replication schedule that was completely successful i.e. Snapshots for all the Objects in the Protection Group were successfully captured. <br> If false, Snapshots are copied from the first Protection Group Run occurring at the start of the replication schedule, even if first Protection Group Run was not completely successful i.e. Snapshots were not captured for all Objects in the Protection Group.
		* `log_retention` - (Optional, List) Specifies the retention of a backup.
		Nested schema for **log_retention**:
			* `data_lock_config` - (Optional, List) Specifies WORM retention type for the snapshots. When a WORM retention type is specified, the snapshots of the Protection Groups using this policy will be kept for the last N days as specified in the duration of the datalock. During that time, the snapshots cannot be deleted.
			Nested schema for **data_lock_config**:
				* `duration` - (Required, Integer) Specifies the duration for a dataLock. <br> Example. If duration is 7 and unit is Months, the dataLock is enabled for last 7 * 30 = 210 days of the backup.
				  * Constraints: The minimum value is `1`.
				* `enable_worm_on_external_target` - (Optional, Boolean) Specifies whether objects in the external target associated with this policy need to be made immutable.
				* `mode` - (Required, String) Specifies the type of WORM retention type. 'Compliance' implies WORM retention is set for compliance reason. 'Administrative' implies WORM retention is set for administrative purposes.
				  * Constraints: Allowable values are: `Compliance`, `Administrative`.
				* `unit` - (Required, String) Specificies the Retention Unit of a dataLock measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Months' then number of retention days will be 365 * 2 = 730 days.
				  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
			* `duration` - (Required, Integer) Specifies the duration for a backup retention. <br> Example. If duration is 7 and unit is Months, the retention of a backup is 7 * 30 = 210 days.
			  * Constraints: The minimum value is `0`.
			* `unit` - (Required, String) Specificies the Retention Unit of a backup measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Years' then number of retention days will be 365 * 2 = 730 days.
			  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
		* `retention` - (Required, List) Specifies the retention of a backup.
		Nested schema for **retention**:
			* `data_lock_config` - (Optional, List) Specifies WORM retention type for the snapshots. When a WORM retention type is specified, the snapshots of the Protection Groups using this policy will be kept for the last N days as specified in the duration of the datalock. During that time, the snapshots cannot be deleted.
			Nested schema for **data_lock_config**:
				* `duration` - (Required, Integer) Specifies the duration for a dataLock. <br> Example. If duration is 7 and unit is Months, the dataLock is enabled for last 7 * 30 = 210 days of the backup.
				  * Constraints: The minimum value is `1`.
				* `enable_worm_on_external_target` - (Optional, Boolean) Specifies whether objects in the external target associated with this policy need to be made immutable.
				* `mode` - (Required, String) Specifies the type of WORM retention type. 'Compliance' implies WORM retention is set for compliance reason. 'Administrative' implies WORM retention is set for administrative purposes.
				  * Constraints: Allowable values are: `Compliance`, `Administrative`.
				* `unit` - (Required, String) Specificies the Retention Unit of a dataLock measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Months' then number of retention days will be 365 * 2 = 730 days.
				  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
			* `duration` - (Required, Integer) Specifies the duration for a backup retention. <br> Example. If duration is 7 and unit is Months, the retention of a backup is 7 * 30 = 210 days.
			  * Constraints: The minimum value is `1`.
			* `unit` - (Required, String) Specificies the Retention Unit of a backup measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Years' then number of retention days will be 365 * 2 = 730 days.
			  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
		* `run_timeouts` - (Optional, List) Specifies the replication/archival timeouts for different type of runs(kFull, kRegular etc.).
		Nested schema for **run_timeouts**:
			* `backup_type` - (Optional, String) The scheduled backup type(kFull, kRegular etc.).
			  * Constraints: Allowable values are: `kRegular`, `kFull`, `kLog`, `kSystem`, `kHydrateCDP`, `kStorageArraySnapshot`.
			* `timeout_mins` - (Optional, Integer) Specifies the timeout in mins.
		* `schedule` - (Required, List) Specifies a schedule fregquency and schedule unit for copying Snapshots to backup targets.
		Nested schema for **schedule**:
			* `frequency` - (Optional, Integer) Specifies a factor to multiply the unit by, to determine the copy schedule. For example if set to 2 and the unit is hourly, then Snapshots from the first eligible Job Run for every 2 hour period is copied.
			  * Constraints: The minimum value is `1`.
			* `unit` - (Required, String) Specifies the frequency that Snapshots should be copied to the specified target. Used in combination with multiplier. <br>'Runs' means that the Snapshot copy occurs after the number of Protection Group Runs equals the number specified in the frequency. <br>'Hours' means that the Snapshot copy occurs hourly at the frequency set in the frequency, for example if scheduleFrequency is 2, the copy occurs every 2 hours. <br>'Days' means that the Snapshot copy occurs daily at the frequency set in the frequency. <br>'Weeks' means that the Snapshot copy occurs weekly at the frequency set in the frequency. <br>'Months' means that the Snapshot copy occurs monthly at the frequency set in the Frequency. <br>'Years' means that the Snapshot copy occurs yearly at the frequency set in the scheduleFrequency.
			  * Constraints: Allowable values are: `Runs`, `Hours`, `Days`, `Weeks`, `Months`, `Years`.
		* `target_id` - (Required, Integer) Specifies the RPaaS target to copy the Snapshots.
		* `target_name` - (Computed, String) Specifies the RPaaS target name where Snapshots are copied.
		* `target_type` - (Optional, String) Specifies the RPaaS target type where Snapshots are copied.
		  * Constraints: Allowable values are: `Tape`, `Cloud`, `Nas`.
* `retry_options` - (Optional, List) Retry Options of a Protection Policy when a Protection Group run fails.
Nested schema for **retry_options**:
	* `retries` - (Optional, Integer) Specifies the number of times to retry capturing Snapshots before the Protection Group Run fails.
	  * Constraints: The minimum value is `0`.
	* `retry_interval_mins` - (Optional, Integer) Specifies the number of minutes before retrying a failed Protection Group.
	  * Constraints: The minimum value is `1`.
* `template_id` - (Optional, String) Specifies the parent policy template id to which the policy is linked to. This field is set only when policy is created from template.
* `version` - (Optional, Integer) Specifies the current policy verison. Policy version is incremented for optionally supporting new features and differentialting across releases.
* `x_ibm_tenant_id` - (Required, String) Specifies the key to be used to encrypt the source credential. If includeSourceCredentials is set to true this key must be specified.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the backup_recovery_protection_policy.
* `is_replicated` - (Boolean) This field is set to true when policy is the replicated policy.
* `policy_id` - The unique identifier of the policy ID.
* `is_usable` - (Boolean) This field is set to true if the linked policy which is internally created from a policy templates qualifies as usable to create more policies on the cluster. If the linked policy is partially filled and can not create a working policy then this field will be set to false. In case of normal policy created on the cluster, this field wont be populated.
* `num_protected_objects` - (Integer) Specifies the number of protected objects using the protection policy.
* `num_protection_groups` - (Integer) Specifies the number of protection groups using the protection policy.


## Import

You can import the `ibm_backup_recovery_protection_policy` resource by using `id`. The ID is formed using tenantID and resourceId.
`id = <tenantId>::<policy_id>`. 


#### Syntax
```
import {
	to = <ibm_backup_recovery_resource>
	id = "<tenantId>::<policy_id>"
}
```

#### Example
```
resource "ibm_backup_recovery_protection_policy" "terra_policy_1" {
    x_ibm_tenant_id = "jhxqx715r9/"
    name = "test-terra-policy-2"
    backup_policy {
      regular {
        incremental{
          schedule{
            day_schedule {
              frequency = 2
            }
            unit = "Days"
          }
        }
        retention {
          duration = 3
          unit = "Weeks"
        }
        primary_backup_target {
          use_default_backup_target = true
        }
      }
    }
    retry_options {
      retries = 3
      retry_interval_mins = 10
    }
}

import {
	to = ibm_backup_recovery_protection_policy.terra_policy_1
	id = "jhxqx715r9/::5170815044477768:1732541085048:238"
}
```
