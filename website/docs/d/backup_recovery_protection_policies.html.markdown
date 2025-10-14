---
layout: "ibm"
page_title: "IBM : ibm_backup_recovery_protection_policies"
description: |-
  Get information about Specifies the details about the Protection Policy.
subcategory: "IBM Backup Recovery"
---

# ibm_backup_recovery_protection_policies

Provides a read-only data source to retrieve information about a Specifies the details about the Protection Policy.. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_backup_recovery_protection_policies" "backup_recovery_protection_policies" {
	x_ibm_tenant_id = ibm_backup_recovery_protection_policy.backup_recovery_protection_policy_instance.x_ibm_tenant_id
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `exclude_linked_policies` - (Optional, Boolean) If excludeLinkedPolicies is set to true then only local policies created on cluster will be returned. The result will exclude all linked policies created from policy templates.
* `backup_recovery_endpoint` - (Optional, String) Backup Recovery Endpoint URL. If provided here, it overrides values configured via environment variable (IBMCLOUD_BACKUP_RECOVERY_ENDPOINT) or endpoints.json.   
* `ids` - (Optional, List) Filter policies by a list of policy ids.
* `include_replicated_policies` - (Optional, Boolean) If includeReplicatedPolicies is set to true, then response will also contain replicated policies. By default, replication policies are not included in the response.
* `include_stats` - (Optional, Boolean) If includeStats is set to true, then response will return number of protection groups and objects. By default, the protection stats are not included in the response.
* `policy_names` - (Optional, List) Filter policies by a list of policy names.
* `request_initiator_type` - (Optional, String) Specifies the type of request from UI, which is used for services like magneto to determine the priority of requests.
  * Constraints: Allowable values are: `UIUser`, `UIAuto`, `Helios`.
* `types` - (Optional, List) Types specifies the policy type of policies to be returned.
  * Constraints: Allowable list items are: `Regular`, `Internal`.
* `x_ibm_tenant_id` - (Required, String) Specifies the key to be used to encrypt the source credential. If includeSourceCredentials is set to true this key must be specified.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the Specifies the details about the Protection Policy..
* `policies` - (List) Specifies a list of protection policies.
Nested schema for **policies**:
	* `backup_policy` - (List) Specifies the backup schedule and retentions of a Protection Policy.
	Nested schema for **backup_policy**:
		* `bmr` - (List) Specifies the BMR schedule in case of physical source protection.
		Nested schema for **bmr**:
			* `retention` - (List) Specifies the retention of a backup.
			Nested schema for **retention**:
				* `data_lock_config` - (List) Specifies WORM retention type for the snapshots. When a WORM retention type is specified, the snapshots of the Protection Groups using this policy will be kept for the last N days as specified in the duration of the datalock. During that time, the snapshots cannot be deleted.
				Nested schema for **data_lock_config**:
					* `duration` - (Integer) Specifies the duration for a dataLock. <br> Example. If duration is 7 and unit is Months, the dataLock is enabled for last 7 * 30 = 210 days of the backup.
					  * Constraints: The minimum value is `1`.
					* `enable_worm_on_external_target` - (Boolean) Specifies whether objects in the external target associated with this policy need to be made immutable.
					* `mode` - (String) Specifies the type of WORM retention type. 'Compliance' implies WORM retention is set for compliance reason. 'Administrative' implies WORM retention is set for administrative purposes.
					  * Constraints: Allowable values are: `Compliance`, `Administrative`.
					* `unit` - (String) Specificies the Retention Unit of a dataLock measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Months' then number of retention days will be 365 * 2 = 730 days.
					  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
				* `duration` - (Integer) Specifies the duration for a backup retention. <br> Example. If duration is 7 and unit is Months, the retention of a backup is 7 * 30 = 210 days.
				  * Constraints: The minimum value is `1`.
				* `unit` - (String) Specificies the Retention Unit of a backup measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Years' then number of retention days will be 365 * 2 = 730 days.
				  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
			* `schedule` - (List) Specifies settings that defines how frequent bmr backup will be performed for a Protection Group.
			Nested schema for **schedule**:
				* `day_schedule` - (List) Specifies settings that define a schedule for a Protection Group runs to start after certain number of days.
				Nested schema for **day_schedule**:
					* `frequency` - (Integer) Specifies a factor to multiply the unit by, to determine the backup schedule. <br> Example: If 'frequency' set to 2 and the unit is 'Hours', then Snapshots are backed up every 2 hours. <br> This field is only applicable if unit is 'Minutes', 'Hours' or 'Days'.
					  * Constraints: The minimum value is `1`.
				* `month_schedule` - (List) Specifies settings that define a schedule for a Protection Group runs to on specific week and specific days of that week.
				Nested schema for **month_schedule**:
					* `day_of_month` - (Integer) Specifies the exact date of the month (such as 18) in a Monthly Schedule specified by unit field as 'Years'. <br> Example: if 'dayOfMonth' is set to '18', a backup is performed on the 18th of every month.
					* `day_of_week` - (List) Specifies a list of days of the week when to start Protection Group Runs. <br> Example: To run a Protection Group on every Monday and Tuesday, set the schedule with following values: <br>  unit: 'Weeks' <br>  dayOfWeek: ['Monday','Tuesday'].
					  * Constraints: Allowable list items are: `Sunday`, `Monday`, `Tuesday`, `Wednesday`, `Thursday`, `Friday`, `Saturday`.
					* `week_of_month` - (String) Specifies the week of the month (such as 'Third') or nth day of month (such as 'First' or 'Last') in a Monthly Schedule specified by unit field as 'Months'. <br>This field can be used in combination with 'dayOfWeek' to define the day in the month to start the Protection Group Run. <br> Example: if 'weekOfMonth' is set to 'Third' and day is set to 'Monday', a backup is performed on the third Monday of every month. <br> Example: if 'weekOfMonth' is set to 'Last' and dayOfWeek is not set, a backup is performed on the last day of every month.
					  * Constraints: Allowable values are: `First`, `Second`, `Third`, `Fourth`, `Last`.
				* `unit` - (String) Specifies how often to start new runs of a Protection Group. <br>'Weeks' specifies that new Protection Group runs start weekly on certain days specified using 'dayOfWeek' field. <br>'Months' specifies that new Protection Group runs start monthly on certain day of specific week.
				  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
				* `week_schedule` - (List) Specifies settings that define a schedule for a Protection Group runs to start on certain days of week.
				Nested schema for **week_schedule**:
					* `day_of_week` - (List) Specifies a list of days of the week when to start Protection Group Runs. <br> Example: To run a Protection Group on every Monday and Tuesday, set the schedule with following values: <br>  unit: 'Weeks' <br>  dayOfWeek: ['Monday','Tuesday'].
					  * Constraints: Allowable list items are: `Sunday`, `Monday`, `Tuesday`, `Wednesday`, `Thursday`, `Friday`, `Saturday`.
				* `year_schedule` - (List) Specifies settings that define a schedule for a Protection Group to run on specific year and specific day of that year.
				Nested schema for **year_schedule**:
					* `day_of_year` - (String) Specifies the day of the Year (such as 'First' or 'Last') in a Yearly Schedule. <br>This field is used to define the day in the year to start the Protection Group Run. <br> Example: if 'dayOfYear' is set to 'First', a backup is performed on the first day of every year. <br> Example: if 'dayOfYear' is set to 'Last', a backup is performed on the last day of every year.
					  * Constraints: Allowable values are: `First`, `Last`.
		* `cdp` - (List) Specifies CDP (Continious Data Protection) backup settings for a Protection Group.
		Nested schema for **cdp**:
			* `retention` - (List) Specifies the retention of a CDP backup.
			Nested schema for **retention**:
				* `data_lock_config` - (List) Specifies WORM retention type for the snapshots. When a WORM retention type is specified, the snapshots of the Protection Groups using this policy will be kept for the last N days as specified in the duration of the datalock. During that time, the snapshots cannot be deleted.
				Nested schema for **data_lock_config**:
					* `duration` - (Integer) Specifies the duration for a dataLock. <br> Example. If duration is 7 and unit is Months, the dataLock is enabled for last 7 * 30 = 210 days of the backup.
					  * Constraints: The minimum value is `1`.
					* `enable_worm_on_external_target` - (Boolean) Specifies whether objects in the external target associated with this policy need to be made immutable.
					* `mode` - (String) Specifies the type of WORM retention type. 'Compliance' implies WORM retention is set for compliance reason. 'Administrative' implies WORM retention is set for administrative purposes.
					  * Constraints: Allowable values are: `Compliance`, `Administrative`.
					* `unit` - (String) Specificies the Retention Unit of a dataLock measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Months' then number of retention days will be 365 * 2 = 730 days.
					  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
				* `duration` - (Integer) Specifies the duration for a cdp backup retention.
				  * Constraints: The minimum value is `1`.
				* `unit` - (String) Specificies the Retention Unit of a CDP backup measured in minutes or hours.
				  * Constraints: Allowable values are: `Minutes`, `Hours`.
		* `log` - (List) Specifies log backup settings for a Protection Group.
		Nested schema for **log**:
			* `retention` - (List) Specifies the retention of a backup.
			Nested schema for **retention**:
				* `data_lock_config` - (List) Specifies WORM retention type for the snapshots. When a WORM retention type is specified, the snapshots of the Protection Groups using this policy will be kept for the last N days as specified in the duration of the datalock. During that time, the snapshots cannot be deleted.
				Nested schema for **data_lock_config**:
					* `duration` - (Integer) Specifies the duration for a dataLock. <br> Example. If duration is 7 and unit is Months, the dataLock is enabled for last 7 * 30 = 210 days of the backup.
					  * Constraints: The minimum value is `1`.
					* `enable_worm_on_external_target` - (Boolean) Specifies whether objects in the external target associated with this policy need to be made immutable.
					* `mode` - (String) Specifies the type of WORM retention type. 'Compliance' implies WORM retention is set for compliance reason. 'Administrative' implies WORM retention is set for administrative purposes.
					  * Constraints: Allowable values are: `Compliance`, `Administrative`.
					* `unit` - (String) Specificies the Retention Unit of a dataLock measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Months' then number of retention days will be 365 * 2 = 730 days.
					  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
				* `duration` - (Integer) Specifies the duration for a backup retention. <br> Example. If duration is 7 and unit is Months, the retention of a backup is 7 * 30 = 210 days.
				  * Constraints: The minimum value is `1`.
				* `unit` - (String) Specificies the Retention Unit of a backup measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Years' then number of retention days will be 365 * 2 = 730 days.
				  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
			* `schedule` - (List) Specifies settings that defines how frequent log backup will be performed for a Protection Group.
			Nested schema for **schedule**:
				* `hour_schedule` - (List) Specifies settings that define a schedule for a Protection Group runs to start after certain number of hours.
				Nested schema for **hour_schedule**:
					* `frequency` - (Integer) Specifies a factor to multiply the unit by, to determine the backup schedule. <br> Example: If 'frequency' set to 2 and the unit is 'Hours', then Snapshots are backed up every 2 hours. <br> This field is only applicable if unit is 'Minutes', 'Hours' or 'Days'.
					  * Constraints: The minimum value is `1`.
				* `minute_schedule` - (List) Specifies settings that define a schedule for a Protection Group runs to start after certain number of minutes.
				Nested schema for **minute_schedule**:
					* `frequency` - (Integer) Specifies a factor to multiply the unit by, to determine the backup schedule. <br> Example: If 'frequency' set to 2 and the unit is 'Hours', then Snapshots are backed up every 2 hours. <br> This field is only applicable if unit is 'Minutes', 'Hours' or 'Days'.
					  * Constraints: The minimum value is `1`.
				* `unit` - (String) Specifies how often to start new Protection Group Runs of a Protection Group. <br>'Minutes' specifies that Protection Group run starts periodically after certain number of minutes specified in 'frequency' field. <br>'Hours' specifies that Protection Group run starts periodically after certain number of hours specified in 'frequency' field.
				  * Constraints: Allowable values are: `Minutes`, `Hours`.
		* `regular` - (List) Specifies the Incremental and Full policy settings and also the common Retention policy settings.".
		Nested schema for **regular**:
			* `full` - (List) Specifies full backup settings for a Protection Group. Currently, full backup settings can be specified by using either of 'schedule' or 'schdulesAndRetentions' field. Using 'schdulesAndRetentions' is recommended when multiple full backups need to be configured. If full and incremental backup has common retention then only setting 'schedule' is recommended.
			Nested schema for **full**:
				* `schedule` - (List) Specifies settings that defines how frequent full backup will be performed for a Protection Group.
				Nested schema for **schedule**:
					* `day_schedule` - (List) Specifies settings that define a schedule for a Protection Group runs to start after certain number of days.
					Nested schema for **day_schedule**:
						* `frequency` - (Integer) Specifies a factor to multiply the unit by, to determine the backup schedule. <br> Example: If 'frequency' set to 2 and the unit is 'Hours', then Snapshots are backed up every 2 hours. <br> This field is only applicable if unit is 'Minutes', 'Hours' or 'Days'.
						  * Constraints: The minimum value is `1`.
					* `month_schedule` - (List) Specifies settings that define a schedule for a Protection Group runs to on specific week and specific days of that week.
					Nested schema for **month_schedule**:
						* `day_of_month` - (Integer) Specifies the exact date of the month (such as 18) in a Monthly Schedule specified by unit field as 'Years'. <br> Example: if 'dayOfMonth' is set to '18', a backup is performed on the 18th of every month.
						* `day_of_week` - (List) Specifies a list of days of the week when to start Protection Group Runs. <br> Example: To run a Protection Group on every Monday and Tuesday, set the schedule with following values: <br>  unit: 'Weeks' <br>  dayOfWeek: ['Monday','Tuesday'].
						  * Constraints: Allowable list items are: `Sunday`, `Monday`, `Tuesday`, `Wednesday`, `Thursday`, `Friday`, `Saturday`.
						* `week_of_month` - (String) Specifies the week of the month (such as 'Third') or nth day of month (such as 'First' or 'Last') in a Monthly Schedule specified by unit field as 'Months'. <br>This field can be used in combination with 'dayOfWeek' to define the day in the month to start the Protection Group Run. <br> Example: if 'weekOfMonth' is set to 'Third' and day is set to 'Monday', a backup is performed on the third Monday of every month. <br> Example: if 'weekOfMonth' is set to 'Last' and dayOfWeek is not set, a backup is performed on the last day of every month.
						  * Constraints: Allowable values are: `First`, `Second`, `Third`, `Fourth`, `Last`.
					* `unit` - (String) Specifies how often to start new runs of a Protection Group. <br>'Days' specifies that Protection Group run starts periodically on every day. For full backup schedule, currently we only support frequecny of 1 which indicates that full backup will be performed daily. <br>'Weeks' specifies that new Protection Group runs start weekly on certain days specified using 'dayOfWeek' field. <br>'Months' specifies that new Protection Group runs start monthly on certain day of specific week. This schedule needs 'weekOfMonth' and 'dayOfWeek' fields to be set. <br>'ProtectOnce' specifies that groups using this policy option will run only once and after that group will permanently be disabled. <br> Example: To run the Protection Group on Second Sunday of Every Month, following schedule need to be set: <br> unit: 'Month' <br> dayOfWeek: 'Sunday' <br> weekOfMonth: 'Second'.
					  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`, `ProtectOnce`.
					* `week_schedule` - (List) Specifies settings that define a schedule for a Protection Group runs to start on certain days of week.
					Nested schema for **week_schedule**:
						* `day_of_week` - (List) Specifies a list of days of the week when to start Protection Group Runs. <br> Example: To run a Protection Group on every Monday and Tuesday, set the schedule with following values: <br>  unit: 'Weeks' <br>  dayOfWeek: ['Monday','Tuesday'].
						  * Constraints: Allowable list items are: `Sunday`, `Monday`, `Tuesday`, `Wednesday`, `Thursday`, `Friday`, `Saturday`.
					* `year_schedule` - (List) Specifies settings that define a schedule for a Protection Group to run on specific year and specific day of that year.
					Nested schema for **year_schedule**:
						* `day_of_year` - (String) Specifies the day of the Year (such as 'First' or 'Last') in a Yearly Schedule. <br>This field is used to define the day in the year to start the Protection Group Run. <br> Example: if 'dayOfYear' is set to 'First', a backup is performed on the first day of every year. <br> Example: if 'dayOfYear' is set to 'Last', a backup is performed on the last day of every year.
						  * Constraints: Allowable values are: `First`, `Last`.
			* `full_backups` - (List) Specifies multiple schedules and retentions for full backup. Specify either of the 'full' or 'fullBackups' values. Its recommended to use 'fullBaackups' value since 'full' will be deprecated after few releases.
			Nested schema for **full_backups**:
				* `retention` - (List) Specifies the retention of a backup.
				Nested schema for **retention**:
					* `data_lock_config` - (List) Specifies WORM retention type for the snapshots. When a WORM retention type is specified, the snapshots of the Protection Groups using this policy will be kept for the last N days as specified in the duration of the datalock. During that time, the snapshots cannot be deleted.
					Nested schema for **data_lock_config**:
						* `duration` - (Integer) Specifies the duration for a dataLock. <br> Example. If duration is 7 and unit is Months, the dataLock is enabled for last 7 * 30 = 210 days of the backup.
						  * Constraints: The minimum value is `1`.
						* `enable_worm_on_external_target` - (Boolean) Specifies whether objects in the external target associated with this policy need to be made immutable.
						* `mode` - (String) Specifies the type of WORM retention type. 'Compliance' implies WORM retention is set for compliance reason. 'Administrative' implies WORM retention is set for administrative purposes.
						  * Constraints: Allowable values are: `Compliance`, `Administrative`.
						* `unit` - (String) Specificies the Retention Unit of a dataLock measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Months' then number of retention days will be 365 * 2 = 730 days.
						  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
					* `duration` - (Integer) Specifies the duration for a backup retention. <br> Example. If duration is 7 and unit is Months, the retention of a backup is 7 * 30 = 210 days.
					  * Constraints: The minimum value is `1`.
					* `unit` - (String) Specificies the Retention Unit of a backup measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Years' then number of retention days will be 365 * 2 = 730 days.
					  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
				* `schedule` - (List) Specifies settings that defines how frequent full backup will be performed for a Protection Group.
				Nested schema for **schedule**:
					* `day_schedule` - (List) Specifies settings that define a schedule for a Protection Group runs to start after certain number of days.
					Nested schema for **day_schedule**:
						* `frequency` - (Integer) Specifies a factor to multiply the unit by, to determine the backup schedule. <br> Example: If 'frequency' set to 2 and the unit is 'Hours', then Snapshots are backed up every 2 hours. <br> This field is only applicable if unit is 'Minutes', 'Hours' or 'Days'.
						  * Constraints: The minimum value is `1`.
					* `month_schedule` - (List) Specifies settings that define a schedule for a Protection Group runs to on specific week and specific days of that week.
					Nested schema for **month_schedule**:
						* `day_of_month` - (Integer) Specifies the exact date of the month (such as 18) in a Monthly Schedule specified by unit field as 'Years'. <br> Example: if 'dayOfMonth' is set to '18', a backup is performed on the 18th of every month.
						* `day_of_week` - (List) Specifies a list of days of the week when to start Protection Group Runs. <br> Example: To run a Protection Group on every Monday and Tuesday, set the schedule with following values: <br>  unit: 'Weeks' <br>  dayOfWeek: ['Monday','Tuesday'].
						  * Constraints: Allowable list items are: `Sunday`, `Monday`, `Tuesday`, `Wednesday`, `Thursday`, `Friday`, `Saturday`.
						* `week_of_month` - (String) Specifies the week of the month (such as 'Third') or nth day of month (such as 'First' or 'Last') in a Monthly Schedule specified by unit field as 'Months'. <br>This field can be used in combination with 'dayOfWeek' to define the day in the month to start the Protection Group Run. <br> Example: if 'weekOfMonth' is set to 'Third' and day is set to 'Monday', a backup is performed on the third Monday of every month. <br> Example: if 'weekOfMonth' is set to 'Last' and dayOfWeek is not set, a backup is performed on the last day of every month.
						  * Constraints: Allowable values are: `First`, `Second`, `Third`, `Fourth`, `Last`.
					* `unit` - (String) Specifies how often to start new runs of a Protection Group. <br>'Days' specifies that Protection Group run starts periodically on every day. For full backup schedule, currently we only support frequecny of 1 which indicates that full backup will be performed daily. <br>'Weeks' specifies that new Protection Group runs start weekly on certain days specified using 'dayOfWeek' field. <br>'Months' specifies that new Protection Group runs start monthly on certain day of specific week. This schedule needs 'weekOfMonth' and 'dayOfWeek' fields to be set. <br>'ProtectOnce' specifies that groups using this policy option will run only once and after that group will permanently be disabled. <br> Example: To run the Protection Group on Second Sunday of Every Month, following schedule need to be set: <br> unit: 'Month' <br> dayOfWeek: 'Sunday' <br> weekOfMonth: 'Second'.
					  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`, `ProtectOnce`.
					* `week_schedule` - (List) Specifies settings that define a schedule for a Protection Group runs to start on certain days of week.
					Nested schema for **week_schedule**:
						* `day_of_week` - (List) Specifies a list of days of the week when to start Protection Group Runs. <br> Example: To run a Protection Group on every Monday and Tuesday, set the schedule with following values: <br>  unit: 'Weeks' <br>  dayOfWeek: ['Monday','Tuesday'].
						  * Constraints: Allowable list items are: `Sunday`, `Monday`, `Tuesday`, `Wednesday`, `Thursday`, `Friday`, `Saturday`.
					* `year_schedule` - (List) Specifies settings that define a schedule for a Protection Group to run on specific year and specific day of that year.
					Nested schema for **year_schedule**:
						* `day_of_year` - (String) Specifies the day of the Year (such as 'First' or 'Last') in a Yearly Schedule. <br>This field is used to define the day in the year to start the Protection Group Run. <br> Example: if 'dayOfYear' is set to 'First', a backup is performed on the first day of every year. <br> Example: if 'dayOfYear' is set to 'Last', a backup is performed on the last day of every year.
						  * Constraints: Allowable values are: `First`, `Last`.
			* `incremental` - (List) Specifies incremental backup settings for a Protection Group.
			Nested schema for **incremental**:
				* `schedule` - (List) Specifies settings that defines how frequent backup will be performed for a Protection Group.
				Nested schema for **schedule**:
					* `day_schedule` - (List) Specifies settings that define a schedule for a Protection Group runs to start after certain number of days.
					Nested schema for **day_schedule**:
						* `frequency` - (Integer) Specifies a factor to multiply the unit by, to determine the backup schedule. <br> Example: If 'frequency' set to 2 and the unit is 'Hours', then Snapshots are backed up every 2 hours. <br> This field is only applicable if unit is 'Minutes', 'Hours' or 'Days'.
						  * Constraints: The minimum value is `1`.
					* `hour_schedule` - (List) Specifies settings that define a schedule for a Protection Group runs to start after certain number of hours.
					Nested schema for **hour_schedule**:
						* `frequency` - (Integer) Specifies a factor to multiply the unit by, to determine the backup schedule. <br> Example: If 'frequency' set to 2 and the unit is 'Hours', then Snapshots are backed up every 2 hours. <br> This field is only applicable if unit is 'Minutes', 'Hours' or 'Days'.
						  * Constraints: The minimum value is `1`.
					* `minute_schedule` - (List) Specifies settings that define a schedule for a Protection Group runs to start after certain number of minutes.
					Nested schema for **minute_schedule**:
						* `frequency` - (Integer) Specifies a factor to multiply the unit by, to determine the backup schedule. <br> Example: If 'frequency' set to 2 and the unit is 'Hours', then Snapshots are backed up every 2 hours. <br> This field is only applicable if unit is 'Minutes', 'Hours' or 'Days'.
						  * Constraints: The minimum value is `1`.
					* `month_schedule` - (List) Specifies settings that define a schedule for a Protection Group runs to on specific week and specific days of that week.
					Nested schema for **month_schedule**:
						* `day_of_month` - (Integer) Specifies the exact date of the month (such as 18) in a Monthly Schedule specified by unit field as 'Years'. <br> Example: if 'dayOfMonth' is set to '18', a backup is performed on the 18th of every month.
						* `day_of_week` - (List) Specifies a list of days of the week when to start Protection Group Runs. <br> Example: To run a Protection Group on every Monday and Tuesday, set the schedule with following values: <br>  unit: 'Weeks' <br>  dayOfWeek: ['Monday','Tuesday'].
						  * Constraints: Allowable list items are: `Sunday`, `Monday`, `Tuesday`, `Wednesday`, `Thursday`, `Friday`, `Saturday`.
						* `week_of_month` - (String) Specifies the week of the month (such as 'Third') or nth day of month (such as 'First' or 'Last') in a Monthly Schedule specified by unit field as 'Months'. <br>This field can be used in combination with 'dayOfWeek' to define the day in the month to start the Protection Group Run. <br> Example: if 'weekOfMonth' is set to 'Third' and day is set to 'Monday', a backup is performed on the third Monday of every month. <br> Example: if 'weekOfMonth' is set to 'Last' and dayOfWeek is not set, a backup is performed on the last day of every month.
						  * Constraints: Allowable values are: `First`, `Second`, `Third`, `Fourth`, `Last`.
					* `unit` - (String) Specifies how often to start new runs of a Protection Group. <br>'Minutes' specifies that Protection Group run starts periodically after certain number of minutes specified in 'frequency' field. <br>'Hours' specifies that Protection Group run starts periodically after certain number of hours specified in 'frequency' field. <br>'Days' specifies that Protection Group run starts periodically after certain number of days specified in 'frequency' field. <br>'Week' specifies that new Protection Group runs start weekly on certain days specified using 'dayOfWeek' field. <br>'Month' specifies that new Protection Group runs start monthly on certain day of specific week. This schedule needs 'weekOfMonth' and 'dayOfWeek' fields to be set. <br> Example: To run the Protection Group on Second Sunday of Every Month, following schedule need to be set: <br> unit: 'Month' <br> dayOfWeek: 'Sunday' <br> weekOfMonth: 'Second'.
					  * Constraints: Allowable values are: `Minutes`, `Hours`, `Days`, `Weeks`, `Months`, `Years`.
					* `week_schedule` - (List) Specifies settings that define a schedule for a Protection Group runs to start on certain days of week.
					Nested schema for **week_schedule**:
						* `day_of_week` - (List) Specifies a list of days of the week when to start Protection Group Runs. <br> Example: To run a Protection Group on every Monday and Tuesday, set the schedule with following values: <br>  unit: 'Weeks' <br>  dayOfWeek: ['Monday','Tuesday'].
						  * Constraints: Allowable list items are: `Sunday`, `Monday`, `Tuesday`, `Wednesday`, `Thursday`, `Friday`, `Saturday`.
					* `year_schedule` - (List) Specifies settings that define a schedule for a Protection Group to run on specific year and specific day of that year.
					Nested schema for **year_schedule**:
						* `day_of_year` - (String) Specifies the day of the Year (such as 'First' or 'Last') in a Yearly Schedule. <br>This field is used to define the day in the year to start the Protection Group Run. <br> Example: if 'dayOfYear' is set to 'First', a backup is performed on the first day of every year. <br> Example: if 'dayOfYear' is set to 'Last', a backup is performed on the last day of every year.
						  * Constraints: Allowable values are: `First`, `Last`.
			* `primary_backup_target` - (List) Specifies the primary backup target settings for regular backups. If the backup target field is not specified then backup will be taken locally on the Cohesity cluster.
			Nested schema for **primary_backup_target**:
				* `archival_target_settings` - (List) Specifies the primary archival settings. Mainly used for cloud direct archive (CAD) policy where primary backup is stored on archival target.
				Nested schema for **archival_target_settings**:
					* `target_id` - (Integer) Specifies the Archival target id to take primary backup.
					* `target_name` - (String) Specifies the Archival target name where Snapshots are copied.
					* `tier_settings` - (List) Specifies the settings tier levels configured with each archival target. The tier settings need to be applied in specific order and default tier should always be passed as first entry in tiers array. The following example illustrates how to configure tiering input for AWS tiering. Same type of input structure applied to other cloud platforms also. <br>If user wants to achieve following tiering for backup, <br>User Desired Tiering- <br><t>1.Archive Full back up for 12 Months <br><t>2.Tier Levels <br><t><t>[1,12] [ <br><t><t><t>s3 (1 to 2 months), (default tier) <br><t><t><t>s3 Intelligent tiering (3 to 6 months), <br><t><t><t>s3 One Zone (7 to 9 months) <br><t><t><t>Glacier (10 to 12 months)] <br><t>API Input <br><t><t>1.tiers-[ <br><t><t><t>{'tierType': 'S3','moveAfterUnit':'months', <br><t><t><t>'moveAfter':2 - move from s3 to s3Inte after 2 months}, <br><t><t><t>{'tierType': 'S3Inte','moveAfterUnit':'months', <br><t><t><t>'moveAfter':4 - move from S3Inte to Glacier after 4 months}, <br><t><t><t>{'tierType': 'Glacier', 'moveAfterUnit':'months', <br><t><t><t>'moveAfter': 3 - move from Glacier to S3 One Zone after 3 months }, <br><t><t><t>{'tierType': 'S3 One Zone', 'moveAfterUnit': nil, <br><t><t><t>'moveAfter': nil - For the last record, 'moveAfter' and 'moveAfterUnit' <br><t><t><t>will be ignored since there are no further tier for data movement } <br><t><t><t>}].
					Nested schema for **tier_settings**:
						* `aws_tiering` - (List) Specifies aws tiers.
						Nested schema for **aws_tiering**:
							* `tiers` - (List) Specifies the tiers that are used to move the archived backup from current tier to next tier. The order of the tiers determines which tier will be used next for moving the archived backup. The first tier input should always be default tier where backup will be acrhived. Each tier specifies how much time after the backup will be moved to next tier from the current tier.
							Nested schema for **tiers**:
								* `move_after` - (Integer) Specifies the time period after which the backup will be moved from current tier to next tier.
								* `move_after_unit` - (String) Specifies the unit for moving the data from current tier to next tier. This unit will be a base unit for the 'moveAfter' field specified below.
								  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
								* `tier_type` - (String) Specifies the AWS tier types.
								  * Constraints: Allowable values are: `kAmazonS3Standard`, `kAmazonS3StandardIA`, `kAmazonS3OneZoneIA`, `kAmazonS3IntelligentTiering`, `kAmazonS3Glacier`, `kAmazonS3GlacierDeepArchive`.
						* `azure_tiering` - (List) Specifies Azure tiers.
						Nested schema for **azure_tiering**:
							* `tiers` - (List) Specifies the tiers that are used to move the archived backup from current tier to next tier. The order of the tiers determines which tier will be used next for moving the archived backup. The first tier input should always be default tier where backup will be acrhived. Each tier specifies how much time after the backup will be moved to next tier from the current tier.
							Nested schema for **tiers**:
								* `move_after` - (Integer) Specifies the time period after which the backup will be moved from current tier to next tier.
								* `move_after_unit` - (String) Specifies the unit for moving the data from current tier to next tier. This unit will be a base unit for the 'moveAfter' field specified below.
								  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
								* `tier_type` - (String) Specifies the Azure tier types.
								  * Constraints: Allowable values are: `kAzureTierHot`, `kAzureTierCool`, `kAzureTierArchive`.
						* `cloud_platform` - (String) Specifies the cloud platform to enable tiering.
						  * Constraints: Allowable values are: `AWS`, `Azure`, `Oracle`, `Google`.
						* `google_tiering` - (List) Specifies Google tiers.
						Nested schema for **google_tiering**:
							* `tiers` - (List) Specifies the tiers that are used to move the archived backup from current tier to next tier. The order of the tiers determines which tier will be used next for moving the archived backup. The first tier input should always be default tier where backup will be acrhived. Each tier specifies how much time after the backup will be moved to next tier from the current tier.
							Nested schema for **tiers**:
								* `move_after` - (Integer) Specifies the time period after which the backup will be moved from current tier to next tier.
								* `move_after_unit` - (String) Specifies the unit for moving the data from current tier to next tier. This unit will be a base unit for the 'moveAfter' field specified below.
								  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
								* `tier_type` - (String) Specifies the Google tier types.
								  * Constraints: Allowable values are: `kGoogleStandard`, `kGoogleRegional`, `kGoogleMultiRegional`, `kGoogleNearline`, `kGoogleColdline`.
						* `oracle_tiering` - (List) Specifies Oracle tiers.
						Nested schema for **oracle_tiering**:
							* `tiers` - (List) Specifies the tiers that are used to move the archived backup from current tier to next tier. The order of the tiers determines which tier will be used next for moving the archived backup. The first tier input should always be default tier where backup will be acrhived. Each tier specifies how much time after the backup will be moved to next tier from the current tier.
							Nested schema for **tiers**:
								* `move_after` - (Integer) Specifies the time period after which the backup will be moved from current tier to next tier.
								* `move_after_unit` - (String) Specifies the unit for moving the data from current tier to next tier. This unit will be a base unit for the 'moveAfter' field specified below.
								  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
								* `tier_type` - (String) Specifies the Oracle tier types.
								  * Constraints: Allowable values are: `kOracleTierStandard`, `kOracleTierArchive`.
				* `target_type` - (String) Specifies the primary backup location where backups will be stored. If not specified, then default is assumed as local backup on Cohesity cluster.
				  * Constraints: Allowable values are: `Local`, `Archival`.
				* `use_default_backup_target` - (Boolean) Specifies if the default primary backup target must be used for backups. If this is not specified or set to false, then targets specified in 'archivalTargetSettings' will be used for backups. If the value is specified as true, then default backup target is used internally. This field should only be set in the environment where tenant policy management is enabled and external targets are assigned to tenant when provisioning tenants.
			* `retention` - (List) Specifies the retention of a backup.
			Nested schema for **retention**:
				* `data_lock_config` - (List) Specifies WORM retention type for the snapshots. When a WORM retention type is specified, the snapshots of the Protection Groups using this policy will be kept for the last N days as specified in the duration of the datalock. During that time, the snapshots cannot be deleted.
				Nested schema for **data_lock_config**:
					* `duration` - (Integer) Specifies the duration for a dataLock. <br> Example. If duration is 7 and unit is Months, the dataLock is enabled for last 7 * 30 = 210 days of the backup.
					  * Constraints: The minimum value is `1`.
					* `enable_worm_on_external_target` - (Boolean) Specifies whether objects in the external target associated with this policy need to be made immutable.
					* `mode` - (String) Specifies the type of WORM retention type. 'Compliance' implies WORM retention is set for compliance reason. 'Administrative' implies WORM retention is set for administrative purposes.
					  * Constraints: Allowable values are: `Compliance`, `Administrative`.
					* `unit` - (String) Specificies the Retention Unit of a dataLock measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Months' then number of retention days will be 365 * 2 = 730 days.
					  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
				* `duration` - (Integer) Specifies the duration for a backup retention. <br> Example. If duration is 7 and unit is Months, the retention of a backup is 7 * 30 = 210 days.
				  * Constraints: The minimum value is `1`.
				* `unit` - (String) Specificies the Retention Unit of a backup measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Years' then number of retention days will be 365 * 2 = 730 days.
				  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
		* `run_timeouts` - (List) Specifies the backup timeouts for different type of runs(kFull, kRegular etc.).
		Nested schema for **run_timeouts**:
			* `backup_type` - (String) The scheduled backup type(kFull, kRegular etc.).
			  * Constraints: Allowable values are: `kRegular`, `kFull`, `kLog`, `kSystem`, `kHydrateCDP`, `kStorageArraySnapshot`.
			* `timeout_mins` - (Integer) Specifies the timeout in mins.
		* `storage_array_snapshot` - (List) Specifies storage snapshot managment backup settings for a Protection Group.
		Nested schema for **storage_array_snapshot**:
			* `retention` - (List) Specifies the retention of a backup.
			Nested schema for **retention**:
				* `data_lock_config` - (List) Specifies WORM retention type for the snapshots. When a WORM retention type is specified, the snapshots of the Protection Groups using this policy will be kept for the last N days as specified in the duration of the datalock. During that time, the snapshots cannot be deleted.
				Nested schema for **data_lock_config**:
					* `duration` - (Integer) Specifies the duration for a dataLock. <br> Example. If duration is 7 and unit is Months, the dataLock is enabled for last 7 * 30 = 210 days of the backup.
					  * Constraints: The minimum value is `1`.
					* `enable_worm_on_external_target` - (Boolean) Specifies whether objects in the external target associated with this policy need to be made immutable.
					* `mode` - (String) Specifies the type of WORM retention type. 'Compliance' implies WORM retention is set for compliance reason. 'Administrative' implies WORM retention is set for administrative purposes.
					  * Constraints: Allowable values are: `Compliance`, `Administrative`.
					* `unit` - (String) Specificies the Retention Unit of a dataLock measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Months' then number of retention days will be 365 * 2 = 730 days.
					  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
				* `duration` - (Integer) Specifies the duration for a backup retention. <br> Example. If duration is 7 and unit is Months, the retention of a backup is 7 * 30 = 210 days.
				  * Constraints: The minimum value is `1`.
				* `unit` - (String) Specificies the Retention Unit of a backup measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Years' then number of retention days will be 365 * 2 = 730 days.
				  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
			* `schedule` - (List) Specifies settings that defines how frequent Storage Snapshot Management backup will be performed for a Protection Group.
			Nested schema for **schedule**:
				* `day_schedule` - (List) Specifies settings that define a schedule for a Protection Group runs to start after certain number of days.
				Nested schema for **day_schedule**:
					* `frequency` - (Integer) Specifies a factor to multiply the unit by, to determine the backup schedule. <br> Example: If 'frequency' set to 2 and the unit is 'Hours', then Snapshots are backed up every 2 hours. <br> This field is only applicable if unit is 'Minutes', 'Hours' or 'Days'.
					  * Constraints: The minimum value is `1`.
				* `hour_schedule` - (List) Specifies settings that define a schedule for a Protection Group runs to start after certain number of hours.
				Nested schema for **hour_schedule**:
					* `frequency` - (Integer) Specifies a factor to multiply the unit by, to determine the backup schedule. <br> Example: If 'frequency' set to 2 and the unit is 'Hours', then Snapshots are backed up every 2 hours. <br> This field is only applicable if unit is 'Minutes', 'Hours' or 'Days'.
					  * Constraints: The minimum value is `1`.
				* `minute_schedule` - (List) Specifies settings that define a schedule for a Protection Group runs to start after certain number of minutes.
				Nested schema for **minute_schedule**:
					* `frequency` - (Integer) Specifies a factor to multiply the unit by, to determine the backup schedule. <br> Example: If 'frequency' set to 2 and the unit is 'Hours', then Snapshots are backed up every 2 hours. <br> This field is only applicable if unit is 'Minutes', 'Hours' or 'Days'.
					  * Constraints: The minimum value is `1`.
				* `month_schedule` - (List) Specifies settings that define a schedule for a Protection Group runs to on specific week and specific days of that week.
				Nested schema for **month_schedule**:
					* `day_of_month` - (Integer) Specifies the exact date of the month (such as 18) in a Monthly Schedule specified by unit field as 'Years'. <br> Example: if 'dayOfMonth' is set to '18', a backup is performed on the 18th of every month.
					* `day_of_week` - (List) Specifies a list of days of the week when to start Protection Group Runs. <br> Example: To run a Protection Group on every Monday and Tuesday, set the schedule with following values: <br>  unit: 'Weeks' <br>  dayOfWeek: ['Monday','Tuesday'].
					  * Constraints: Allowable list items are: `Sunday`, `Monday`, `Tuesday`, `Wednesday`, `Thursday`, `Friday`, `Saturday`.
					* `week_of_month` - (String) Specifies the week of the month (such as 'Third') or nth day of month (such as 'First' or 'Last') in a Monthly Schedule specified by unit field as 'Months'. <br>This field can be used in combination with 'dayOfWeek' to define the day in the month to start the Protection Group Run. <br> Example: if 'weekOfMonth' is set to 'Third' and day is set to 'Monday', a backup is performed on the third Monday of every month. <br> Example: if 'weekOfMonth' is set to 'Last' and dayOfWeek is not set, a backup is performed on the last day of every month.
					  * Constraints: Allowable values are: `First`, `Second`, `Third`, `Fourth`, `Last`.
				* `unit` - (String) Specifies how often to start new Protection Group Runs of a Protection Group. <br>'Minutes' specifies that Protection Group run starts periodically after certain number of minutes specified in 'frequency' field. <br>'Hours' specifies that Protection Group run starts periodically after certain number of hours specified in 'frequency' field.
				  * Constraints: Allowable values are: `Minutes`, `Hours`, `Days`, `Weeks`, `Months`, `Years`.
				* `week_schedule` - (List) Specifies settings that define a schedule for a Protection Group runs to start on certain days of week.
				Nested schema for **week_schedule**:
					* `day_of_week` - (List) Specifies a list of days of the week when to start Protection Group Runs. <br> Example: To run a Protection Group on every Monday and Tuesday, set the schedule with following values: <br>  unit: 'Weeks' <br>  dayOfWeek: ['Monday','Tuesday'].
					  * Constraints: Allowable list items are: `Sunday`, `Monday`, `Tuesday`, `Wednesday`, `Thursday`, `Friday`, `Saturday`.
				* `year_schedule` - (List) Specifies settings that define a schedule for a Protection Group to run on specific year and specific day of that year.
				Nested schema for **year_schedule**:
					* `day_of_year` - (String) Specifies the day of the Year (such as 'First' or 'Last') in a Yearly Schedule. <br>This field is used to define the day in the year to start the Protection Group Run. <br> Example: if 'dayOfYear' is set to 'First', a backup is performed on the first day of every year. <br> Example: if 'dayOfYear' is set to 'Last', a backup is performed on the last day of every year.
					  * Constraints: Allowable values are: `First`, `Last`.
	* `blackout_window` - (List) List of Blackout Windows. If specified, this field defines blackout periods when new Group Runs are not started. If a Group Run has been scheduled but not yet executed and the blackout period starts, the behavior depends on the policy field AbortInBlackoutPeriod.
	Nested schema for **blackout_window**:
		* `config_id` - (String) Specifies the unique identifier for the target getting added. This field need to be passed olny when policies are updated.
		* `day` - (String) Specifies a day in the week when no new Protection Group Runs should be started such as 'Sunday'. Specifies a day in a week such as 'Sunday', 'Monday', etc.
		  * Constraints: Allowable values are: `Sunday`, `Monday`, `Tuesday`, `Wednesday`, `Thursday`, `Friday`, `Saturday`.
		* `end_time` - (List) Specifies the time of day. Used for scheduling purposes.
		Nested schema for **end_time**:
			* `hour` - (Integer) Specifies the hour of the day (0-23).
			  * Constraints: The maximum value is `23`. The minimum value is `0`.
			* `minute` - (Integer) Specifies the minute of the hour (0-59).
			  * Constraints: The maximum value is `59`. The minimum value is `0`.
			* `time_zone` - (String) Specifies the time zone of the user. If not specified, default value is assumed as America/Los_Angeles.
			  * Constraints: The default value is `America/Los_Angeles`.
		* `start_time` - (List) Specifies the time of day. Used for scheduling purposes.
		Nested schema for **start_time**:
			* `hour` - (Integer) Specifies the hour of the day (0-23).
			  * Constraints: The maximum value is `23`. The minimum value is `0`.
			* `minute` - (Integer) Specifies the minute of the hour (0-59).
			  * Constraints: The maximum value is `59`. The minimum value is `0`.
			* `time_zone` - (String) Specifies the time zone of the user. If not specified, default value is assumed as America/Los_Angeles.
			  * Constraints: The default value is `America/Los_Angeles`.
	* `cascaded_targets_config` - (List) Specifies the configuration for cascaded replications. Using cascaded replication, replication cluster(Rx) can further replicate and archive the snapshot copies to further targets. Its recommended to create cascaded configuration where protection group will be created.
	Nested schema for **cascaded_targets_config**:
		* `remote_targets` - (List) Specifies the replication, archival and cloud spin targets of Protection Policy.
		Nested schema for **remote_targets**:
			* `archival_targets` - (List)
			Nested schema for **archival_targets**:
				* `backup_run_type` - (String) Specifies which type of run should be copied, if not set, all types of runs will be eligible for copying. If set, this will ensure that the first run of given type in the scheduled period will get copied. Currently, this can only be set to Full.
				  * Constraints: Allowable values are: `Regular`, `Full`, `Log`, `System`, `StorageArraySnapshot`.
				* `config_id` - (String) Specifies the unique identifier for the target getting added. This field need to be passed only when policies are being updated.
				* `copy_on_run_success` - (Boolean) Specifies if Snapshots are copied from the first completely successful Protection Group Run or the first partially successful Protection Group Run occurring at the start of the replication schedule. <br> If true, Snapshots are copied from the first Protection Group Run occurring at the start of the replication schedule that was completely successful i.e. Snapshots for all the Objects in the Protection Group were successfully captured. <br> If false, Snapshots are copied from the first Protection Group Run occurring at the start of the replication schedule, even if first Protection Group Run was not completely successful i.e. Snapshots were not captured for all Objects in the Protection Group.
				* `extended_retention` - (List) Specifies additional retention policies that should be applied to the archived backup. Archived backup snapshot will be retained up to a time that is the maximum of all retention policies that are applicable to it.
				Nested schema for **extended_retention**:
					* `config_id` - (String) Specifies the unique identifier for the target getting added. This field need to be passed olny when policies are updated.
					* `retention` - (List) Specifies the retention of a backup.
					Nested schema for **retention**:
						* `data_lock_config` - (List) Specifies WORM retention type for the snapshots. When a WORM retention type is specified, the snapshots of the Protection Groups using this policy will be kept for the last N days as specified in the duration of the datalock. During that time, the snapshots cannot be deleted.
						Nested schema for **data_lock_config**:
							* `duration` - (Integer) Specifies the duration for a dataLock. <br> Example. If duration is 7 and unit is Months, the dataLock is enabled for last 7 * 30 = 210 days of the backup.
							  * Constraints: The minimum value is `1`.
							* `enable_worm_on_external_target` - (Boolean) Specifies whether objects in the external target associated with this policy need to be made immutable.
							* `mode` - (String) Specifies the type of WORM retention type. 'Compliance' implies WORM retention is set for compliance reason. 'Administrative' implies WORM retention is set for administrative purposes.
							  * Constraints: Allowable values are: `Compliance`, `Administrative`.
							* `unit` - (String) Specificies the Retention Unit of a dataLock measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Months' then number of retention days will be 365 * 2 = 730 days.
							  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
						* `duration` - (Integer) Specifies the duration for a backup retention. <br> Example. If duration is 7 and unit is Months, the retention of a backup is 7 * 30 = 210 days.
						  * Constraints: The minimum value is `1`.
						* `unit` - (String) Specificies the Retention Unit of a backup measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Years' then number of retention days will be 365 * 2 = 730 days.
						  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
					* `run_type` - (String) The backup run type to which this extended retention applies to. If this is not set, the extended retention will be applicable to all non-log backup types. Currently, the only value that can be set here is Full.'Regular' indicates a incremental (CBT) backup. Incremental backups utilizing CBT (if supported) are captured of the target protection objects. The first run of a Regular schedule captures all the blocks.'Full' indicates a full (no CBT) backup. A complete backup (all blocks) of the target protection objects are always captured and Change Block Tracking (CBT) is not utilized.'Log' indicates a Database Log backup. Capture the database transaction logs to allow rolling back to a specific point in time.'System' indicates a system backup. System backups are used to do bare metal recovery of the system to a specific point in time.
					  * Constraints: Allowable values are: `Regular`, `Full`, `Log`, `System`, `StorageArraySnapshot`.
					* `schedule` - (List) Specifies a schedule frequency and schedule unit for Extended Retentions.
					Nested schema for **schedule**:
						* `frequency` - (Integer) Specifies a factor to multiply the unit by, to determine the retention schedule. For example if set to 2 and the unit is hourly, then Snapshots from the first eligible Job Run for every 2 hour period is retained.
						  * Constraints: The minimum value is `1`.
						* `unit` - (String) Specifies the unit interval for retention of Snapshots. <br>'Runs' means that the Snapshot copy retained after the number of Protection Group Runs equals the number specified in the frequency. <br>'Hours' means that the Snapshot copy retained hourly at the frequency set in the frequency, for example if scheduleFrequency is 2, the copy occurs every 2 hours. <br>'Days' means that the Snapshot copy gets retained daily at the frequency set in the frequency. <br>'Weeks' means that the Snapshot copy is retained weekly at the frequency set in the frequency. <br>'Months' means that the Snapshot copy is retained monthly at the frequency set in the Frequency. <br>'Years' means that the Snapshot copy is retained yearly at the frequency set in the Frequency.
						  * Constraints: Allowable values are: `Runs`, `Hours`, `Days`, `Weeks`, `Months`, `Years`.
				* `log_retention` - (List) Specifies the retention of a backup.
				Nested schema for **log_retention**:
					* `data_lock_config` - (List) Specifies WORM retention type for the snapshots. When a WORM retention type is specified, the snapshots of the Protection Groups using this policy will be kept for the last N days as specified in the duration of the datalock. During that time, the snapshots cannot be deleted.
					Nested schema for **data_lock_config**:
						* `duration` - (Integer) Specifies the duration for a dataLock. <br> Example. If duration is 7 and unit is Months, the dataLock is enabled for last 7 * 30 = 210 days of the backup.
						  * Constraints: The minimum value is `1`.
						* `enable_worm_on_external_target` - (Boolean) Specifies whether objects in the external target associated with this policy need to be made immutable.
						* `mode` - (String) Specifies the type of WORM retention type. 'Compliance' implies WORM retention is set for compliance reason. 'Administrative' implies WORM retention is set for administrative purposes.
						  * Constraints: Allowable values are: `Compliance`, `Administrative`.
						* `unit` - (String) Specificies the Retention Unit of a dataLock measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Months' then number of retention days will be 365 * 2 = 730 days.
						  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
					* `duration` - (Integer) Specifies the duration for a backup retention. <br> Example. If duration is 7 and unit is Months, the retention of a backup is 7 * 30 = 210 days.
					  * Constraints: The minimum value is `0`.
					* `unit` - (String) Specificies the Retention Unit of a backup measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Years' then number of retention days will be 365 * 2 = 730 days.
					  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
				* `retention` - (List) Specifies the retention of a backup.
				Nested schema for **retention**:
					* `data_lock_config` - (List) Specifies WORM retention type for the snapshots. When a WORM retention type is specified, the snapshots of the Protection Groups using this policy will be kept for the last N days as specified in the duration of the datalock. During that time, the snapshots cannot be deleted.
					Nested schema for **data_lock_config**:
						* `duration` - (Integer) Specifies the duration for a dataLock. <br> Example. If duration is 7 and unit is Months, the dataLock is enabled for last 7 * 30 = 210 days of the backup.
						  * Constraints: The minimum value is `1`.
						* `enable_worm_on_external_target` - (Boolean) Specifies whether objects in the external target associated with this policy need to be made immutable.
						* `mode` - (String) Specifies the type of WORM retention type. 'Compliance' implies WORM retention is set for compliance reason. 'Administrative' implies WORM retention is set for administrative purposes.
						  * Constraints: Allowable values are: `Compliance`, `Administrative`.
						* `unit` - (String) Specificies the Retention Unit of a dataLock measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Months' then number of retention days will be 365 * 2 = 730 days.
						  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
					* `duration` - (Integer) Specifies the duration for a backup retention. <br> Example. If duration is 7 and unit is Months, the retention of a backup is 7 * 30 = 210 days.
					  * Constraints: The minimum value is `1`.
					* `unit` - (String) Specificies the Retention Unit of a backup measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Years' then number of retention days will be 365 * 2 = 730 days.
					  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
				* `run_timeouts` - (List) Specifies the replication/archival timeouts for different type of runs(kFull, kRegular etc.).
				Nested schema for **run_timeouts**:
					* `backup_type` - (String) The scheduled backup type(kFull, kRegular etc.).
					  * Constraints: Allowable values are: `kRegular`, `kFull`, `kLog`, `kSystem`, `kHydrateCDP`, `kStorageArraySnapshot`.
					* `timeout_mins` - (Integer) Specifies the timeout in mins.
				* `schedule` - (List) Specifies a schedule fregquency and schedule unit for copying Snapshots to backup targets.
				Nested schema for **schedule**:
					* `frequency` - (Integer) Specifies a factor to multiply the unit by, to determine the copy schedule. For example if set to 2 and the unit is hourly, then Snapshots from the first eligible Job Run for every 2 hour period is copied.
					  * Constraints: The minimum value is `1`.
					* `unit` - (String) Specifies the frequency that Snapshots should be copied to the specified target. Used in combination with multiplier. <br>'Runs' means that the Snapshot copy occurs after the number of Protection Group Runs equals the number specified in the frequency. <br>'Hours' means that the Snapshot copy occurs hourly at the frequency set in the frequency, for example if scheduleFrequency is 2, the copy occurs every 2 hours. <br>'Days' means that the Snapshot copy occurs daily at the frequency set in the frequency. <br>'Weeks' means that the Snapshot copy occurs weekly at the frequency set in the frequency. <br>'Months' means that the Snapshot copy occurs monthly at the frequency set in the Frequency. <br>'Years' means that the Snapshot copy occurs yearly at the frequency set in the scheduleFrequency.
					  * Constraints: Allowable values are: `Runs`, `Hours`, `Days`, `Weeks`, `Months`, `Years`.
				* `target_id` - (Integer) Specifies the Archival target to copy the Snapshots to.
				* `target_name` - (String) Specifies the Archival target name where Snapshots are copied.
				* `target_type` - (String) Specifies the Archival target type where Snapshots are copied.
				  * Constraints: Allowable values are: `Tape`, `Cloud`, `Nas`.
				* `tier_settings` - (List) Specifies the settings tier levels configured with each archival target. The tier settings need to be applied in specific order and default tier should always be passed as first entry in tiers array. The following example illustrates how to configure tiering input for AWS tiering. Same type of input structure applied to other cloud platforms also. <br>If user wants to achieve following tiering for backup, <br>User Desired Tiering- <br><t>1.Archive Full back up for 12 Months <br><t>2.Tier Levels <br><t><t>[1,12] [ <br><t><t><t>s3 (1 to 2 months), (default tier) <br><t><t><t>s3 Intelligent tiering (3 to 6 months), <br><t><t><t>s3 One Zone (7 to 9 months) <br><t><t><t>Glacier (10 to 12 months)] <br><t>API Input <br><t><t>1.tiers-[ <br><t><t><t>{'tierType': 'S3','moveAfterUnit':'months', <br><t><t><t>'moveAfter':2 - move from s3 to s3Inte after 2 months}, <br><t><t><t>{'tierType': 'S3Inte','moveAfterUnit':'months', <br><t><t><t>'moveAfter':4 - move from S3Inte to Glacier after 4 months}, <br><t><t><t>{'tierType': 'Glacier', 'moveAfterUnit':'months', <br><t><t><t>'moveAfter': 3 - move from Glacier to S3 One Zone after 3 months }, <br><t><t><t>{'tierType': 'S3 One Zone', 'moveAfterUnit': nil, <br><t><t><t>'moveAfter': nil - For the last record, 'moveAfter' and 'moveAfterUnit' <br><t><t><t>will be ignored since there are no further tier for data movement } <br><t><t><t>}].
				Nested schema for **tier_settings**:
					* `aws_tiering` - (List) Specifies aws tiers.
					Nested schema for **aws_tiering**:
						* `tiers` - (List) Specifies the tiers that are used to move the archived backup from current tier to next tier. The order of the tiers determines which tier will be used next for moving the archived backup. The first tier input should always be default tier where backup will be acrhived. Each tier specifies how much time after the backup will be moved to next tier from the current tier.
						Nested schema for **tiers**:
							* `move_after` - (Integer) Specifies the time period after which the backup will be moved from current tier to next tier.
							* `move_after_unit` - (String) Specifies the unit for moving the data from current tier to next tier. This unit will be a base unit for the 'moveAfter' field specified below.
							  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
							* `tier_type` - (String) Specifies the AWS tier types.
							  * Constraints: Allowable values are: `kAmazonS3Standard`, `kAmazonS3StandardIA`, `kAmazonS3OneZoneIA`, `kAmazonS3IntelligentTiering`, `kAmazonS3Glacier`, `kAmazonS3GlacierDeepArchive`.
					* `azure_tiering` - (List) Specifies Azure tiers.
					Nested schema for **azure_tiering**:
						* `tiers` - (List) Specifies the tiers that are used to move the archived backup from current tier to next tier. The order of the tiers determines which tier will be used next for moving the archived backup. The first tier input should always be default tier where backup will be acrhived. Each tier specifies how much time after the backup will be moved to next tier from the current tier.
						Nested schema for **tiers**:
							* `move_after` - (Integer) Specifies the time period after which the backup will be moved from current tier to next tier.
							* `move_after_unit` - (String) Specifies the unit for moving the data from current tier to next tier. This unit will be a base unit for the 'moveAfter' field specified below.
							  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
							* `tier_type` - (String) Specifies the Azure tier types.
							  * Constraints: Allowable values are: `kAzureTierHot`, `kAzureTierCool`, `kAzureTierArchive`.
					* `cloud_platform` - (String) Specifies the cloud platform to enable tiering.
					  * Constraints: Allowable values are: `AWS`, `Azure`, `Oracle`, `Google`.
					* `google_tiering` - (List) Specifies Google tiers.
					Nested schema for **google_tiering**:
						* `tiers` - (List) Specifies the tiers that are used to move the archived backup from current tier to next tier. The order of the tiers determines which tier will be used next for moving the archived backup. The first tier input should always be default tier where backup will be acrhived. Each tier specifies how much time after the backup will be moved to next tier from the current tier.
						Nested schema for **tiers**:
							* `move_after` - (Integer) Specifies the time period after which the backup will be moved from current tier to next tier.
							* `move_after_unit` - (String) Specifies the unit for moving the data from current tier to next tier. This unit will be a base unit for the 'moveAfter' field specified below.
							  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
							* `tier_type` - (String) Specifies the Google tier types.
							  * Constraints: Allowable values are: `kGoogleStandard`, `kGoogleRegional`, `kGoogleMultiRegional`, `kGoogleNearline`, `kGoogleColdline`.
					* `oracle_tiering` - (List) Specifies Oracle tiers.
					Nested schema for **oracle_tiering**:
						* `tiers` - (List) Specifies the tiers that are used to move the archived backup from current tier to next tier. The order of the tiers determines which tier will be used next for moving the archived backup. The first tier input should always be default tier where backup will be acrhived. Each tier specifies how much time after the backup will be moved to next tier from the current tier.
						Nested schema for **tiers**:
							* `move_after` - (Integer) Specifies the time period after which the backup will be moved from current tier to next tier.
							* `move_after_unit` - (String) Specifies the unit for moving the data from current tier to next tier. This unit will be a base unit for the 'moveAfter' field specified below.
							  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
							* `tier_type` - (String) Specifies the Oracle tier types.
							  * Constraints: Allowable values are: `kOracleTierStandard`, `kOracleTierArchive`.
			* `cloud_spin_targets` - (List)
			Nested schema for **cloud_spin_targets**:
				* `backup_run_type` - (String) Specifies which type of run should be copied, if not set, all types of runs will be eligible for copying. If set, this will ensure that the first run of given type in the scheduled period will get copied. Currently, this can only be set to Full.
				  * Constraints: Allowable values are: `Regular`, `Full`, `Log`, `System`, `StorageArraySnapshot`.
				* `config_id` - (String) Specifies the unique identifier for the target getting added. This field need to be passed only when policies are being updated.
				* `copy_on_run_success` - (Boolean) Specifies if Snapshots are copied from the first completely successful Protection Group Run or the first partially successful Protection Group Run occurring at the start of the replication schedule. <br> If true, Snapshots are copied from the first Protection Group Run occurring at the start of the replication schedule that was completely successful i.e. Snapshots for all the Objects in the Protection Group were successfully captured. <br> If false, Snapshots are copied from the first Protection Group Run occurring at the start of the replication schedule, even if first Protection Group Run was not completely successful i.e. Snapshots were not captured for all Objects in the Protection Group.
				* `log_retention` - (List) Specifies the retention of a backup.
				Nested schema for **log_retention**:
					* `data_lock_config` - (List) Specifies WORM retention type for the snapshots. When a WORM retention type is specified, the snapshots of the Protection Groups using this policy will be kept for the last N days as specified in the duration of the datalock. During that time, the snapshots cannot be deleted.
					Nested schema for **data_lock_config**:
						* `duration` - (Integer) Specifies the duration for a dataLock. <br> Example. If duration is 7 and unit is Months, the dataLock is enabled for last 7 * 30 = 210 days of the backup.
						  * Constraints: The minimum value is `1`.
						* `enable_worm_on_external_target` - (Boolean) Specifies whether objects in the external target associated with this policy need to be made immutable.
						* `mode` - (String) Specifies the type of WORM retention type. 'Compliance' implies WORM retention is set for compliance reason. 'Administrative' implies WORM retention is set for administrative purposes.
						  * Constraints: Allowable values are: `Compliance`, `Administrative`.
						* `unit` - (String) Specificies the Retention Unit of a dataLock measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Months' then number of retention days will be 365 * 2 = 730 days.
						  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
					* `duration` - (Integer) Specifies the duration for a backup retention. <br> Example. If duration is 7 and unit is Months, the retention of a backup is 7 * 30 = 210 days.
					  * Constraints: The minimum value is `0`.
					* `unit` - (String) Specificies the Retention Unit of a backup measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Years' then number of retention days will be 365 * 2 = 730 days.
					  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
				* `retention` - (List) Specifies the retention of a backup.
				Nested schema for **retention**:
					* `data_lock_config` - (List) Specifies WORM retention type for the snapshots. When a WORM retention type is specified, the snapshots of the Protection Groups using this policy will be kept for the last N days as specified in the duration of the datalock. During that time, the snapshots cannot be deleted.
					Nested schema for **data_lock_config**:
						* `duration` - (Integer) Specifies the duration for a dataLock. <br> Example. If duration is 7 and unit is Months, the dataLock is enabled for last 7 * 30 = 210 days of the backup.
						  * Constraints: The minimum value is `1`.
						* `enable_worm_on_external_target` - (Boolean) Specifies whether objects in the external target associated with this policy need to be made immutable.
						* `mode` - (String) Specifies the type of WORM retention type. 'Compliance' implies WORM retention is set for compliance reason. 'Administrative' implies WORM retention is set for administrative purposes.
						  * Constraints: Allowable values are: `Compliance`, `Administrative`.
						* `unit` - (String) Specificies the Retention Unit of a dataLock measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Months' then number of retention days will be 365 * 2 = 730 days.
						  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
					* `duration` - (Integer) Specifies the duration for a backup retention. <br> Example. If duration is 7 and unit is Months, the retention of a backup is 7 * 30 = 210 days.
					  * Constraints: The minimum value is `1`.
					* `unit` - (String) Specificies the Retention Unit of a backup measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Years' then number of retention days will be 365 * 2 = 730 days.
					  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
				* `run_timeouts` - (List) Specifies the replication/archival timeouts for different type of runs(kFull, kRegular etc.).
				Nested schema for **run_timeouts**:
					* `backup_type` - (String) The scheduled backup type(kFull, kRegular etc.).
					  * Constraints: Allowable values are: `kRegular`, `kFull`, `kLog`, `kSystem`, `kHydrateCDP`, `kStorageArraySnapshot`.
					* `timeout_mins` - (Integer) Specifies the timeout in mins.
				* `schedule` - (List) Specifies a schedule fregquency and schedule unit for copying Snapshots to backup targets.
				Nested schema for **schedule**:
					* `frequency` - (Integer) Specifies a factor to multiply the unit by, to determine the copy schedule. For example if set to 2 and the unit is hourly, then Snapshots from the first eligible Job Run for every 2 hour period is copied.
					  * Constraints: The minimum value is `1`.
					* `unit` - (String) Specifies the frequency that Snapshots should be copied to the specified target. Used in combination with multiplier. <br>'Runs' means that the Snapshot copy occurs after the number of Protection Group Runs equals the number specified in the frequency. <br>'Hours' means that the Snapshot copy occurs hourly at the frequency set in the frequency, for example if scheduleFrequency is 2, the copy occurs every 2 hours. <br>'Days' means that the Snapshot copy occurs daily at the frequency set in the frequency. <br>'Weeks' means that the Snapshot copy occurs weekly at the frequency set in the frequency. <br>'Months' means that the Snapshot copy occurs monthly at the frequency set in the Frequency. <br>'Years' means that the Snapshot copy occurs yearly at the frequency set in the scheduleFrequency.
					  * Constraints: Allowable values are: `Runs`, `Hours`, `Days`, `Weeks`, `Months`, `Years`.
				* `target` - (List) Specifies the details about Cloud Spin target where backup snapshots may be converted and stored.
				Nested schema for **target**:
					* `aws_params` - (List) Specifies various resources when converting and deploying a VM to AWS.
					Nested schema for **aws_params**:
						* `custom_tag_list` - (List) Specifies tags of various resources when converting and deploying a VM to AWS.
						Nested schema for **custom_tag_list**:
							* `key` - (String) Specifies key of the custom tag.
							* `value` - (String) Specifies value of the custom tag.
						* `region` - (Integer) Specifies id of the AWS region in which to deploy the VM.
						* `subnet_id` - (Integer) Specifies id of the subnet within above VPC.
						* `vpc_id` - (Integer) Specifies id of the Virtual Private Cloud to chose for the instance type.
					* `azure_params` - (List) Specifies various resources when converting and deploying a VM to Azure.
					Nested schema for **azure_params**:
						* `availability_set_id` - (Integer) Specifies the availability set.
						* `network_resource_group_id` - (Integer) Specifies id of the resource group for the selected virtual network.
						* `resource_group_id` - (Integer) Specifies id of the Azure resource group. Its value is globally unique within Azure.
						* `storage_account_id` - (Integer) Specifies id of the storage account that will contain the storage container within which we will create the blob that will become the VHD disk for the cloned VM.
						* `storage_container_id` - (Integer) Specifies id of the storage container within the above storage account.
						* `storage_resource_group_id` - (Integer) Specifies id of the resource group for the selected storage account.
						* `temp_vm_resource_group_id` - (Integer) Specifies id of the temporary Azure resource group.
						* `temp_vm_storage_account_id` - (Integer) Specifies id of the temporary VM storage account that will contain the storage container within which we will create the blob that will become the VHD disk for the cloned VM.
						* `temp_vm_storage_container_id` - (Integer) Specifies id of the temporary VM storage container within the above storage account.
						* `temp_vm_subnet_id` - (Integer) Specifies Id of the temporary VM subnet within the above virtual network.
						* `temp_vm_virtual_network_id` - (Integer) Specifies Id of the temporary VM Virtual Network.
					* `id` - (Integer) Specifies the unique id of the cloud spin entity.
					* `name` - (String) Specifies the name of the already added cloud spin target.
			* `onprem_deploy_targets` - (List)
			Nested schema for **onprem_deploy_targets**:
				* `backup_run_type` - (String) Specifies which type of run should be copied, if not set, all types of runs will be eligible for copying. If set, this will ensure that the first run of given type in the scheduled period will get copied. Currently, this can only be set to Full.
				  * Constraints: Allowable values are: `Regular`, `Full`, `Log`, `System`, `StorageArraySnapshot`.
				* `config_id` - (String) Specifies the unique identifier for the target getting added. This field need to be passed only when policies are being updated.
				* `copy_on_run_success` - (Boolean) Specifies if Snapshots are copied from the first completely successful Protection Group Run or the first partially successful Protection Group Run occurring at the start of the replication schedule. <br> If true, Snapshots are copied from the first Protection Group Run occurring at the start of the replication schedule that was completely successful i.e. Snapshots for all the Objects in the Protection Group were successfully captured. <br> If false, Snapshots are copied from the first Protection Group Run occurring at the start of the replication schedule, even if first Protection Group Run was not completely successful i.e. Snapshots were not captured for all Objects in the Protection Group.
				* `log_retention` - (List) Specifies the retention of a backup.
				Nested schema for **log_retention**:
					* `data_lock_config` - (List) Specifies WORM retention type for the snapshots. When a WORM retention type is specified, the snapshots of the Protection Groups using this policy will be kept for the last N days as specified in the duration of the datalock. During that time, the snapshots cannot be deleted.
					Nested schema for **data_lock_config**:
						* `duration` - (Integer) Specifies the duration for a dataLock. <br> Example. If duration is 7 and unit is Months, the dataLock is enabled for last 7 * 30 = 210 days of the backup.
						  * Constraints: The minimum value is `1`.
						* `enable_worm_on_external_target` - (Boolean) Specifies whether objects in the external target associated with this policy need to be made immutable.
						* `mode` - (String) Specifies the type of WORM retention type. 'Compliance' implies WORM retention is set for compliance reason. 'Administrative' implies WORM retention is set for administrative purposes.
						  * Constraints: Allowable values are: `Compliance`, `Administrative`.
						* `unit` - (String) Specificies the Retention Unit of a dataLock measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Months' then number of retention days will be 365 * 2 = 730 days.
						  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
					* `duration` - (Integer) Specifies the duration for a backup retention. <br> Example. If duration is 7 and unit is Months, the retention of a backup is 7 * 30 = 210 days.
					  * Constraints: The minimum value is `0`.
					* `unit` - (String) Specificies the Retention Unit of a backup measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Years' then number of retention days will be 365 * 2 = 730 days.
					  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
				* `params` - (List) Specifies the details about OnpremDeploy target where backup snapshots may be converted and deployed.
				Nested schema for **params**:
					* `id` - (Integer) Specifies the unique id of the onprem entity.
				* `retention` - (List) Specifies the retention of a backup.
				Nested schema for **retention**:
					* `data_lock_config` - (List) Specifies WORM retention type for the snapshots. When a WORM retention type is specified, the snapshots of the Protection Groups using this policy will be kept for the last N days as specified in the duration of the datalock. During that time, the snapshots cannot be deleted.
					Nested schema for **data_lock_config**:
						* `duration` - (Integer) Specifies the duration for a dataLock. <br> Example. If duration is 7 and unit is Months, the dataLock is enabled for last 7 * 30 = 210 days of the backup.
						  * Constraints: The minimum value is `1`.
						* `enable_worm_on_external_target` - (Boolean) Specifies whether objects in the external target associated with this policy need to be made immutable.
						* `mode` - (String) Specifies the type of WORM retention type. 'Compliance' implies WORM retention is set for compliance reason. 'Administrative' implies WORM retention is set for administrative purposes.
						  * Constraints: Allowable values are: `Compliance`, `Administrative`.
						* `unit` - (String) Specificies the Retention Unit of a dataLock measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Months' then number of retention days will be 365 * 2 = 730 days.
						  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
					* `duration` - (Integer) Specifies the duration for a backup retention. <br> Example. If duration is 7 and unit is Months, the retention of a backup is 7 * 30 = 210 days.
					  * Constraints: The minimum value is `1`.
					* `unit` - (String) Specificies the Retention Unit of a backup measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Years' then number of retention days will be 365 * 2 = 730 days.
					  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
				* `run_timeouts` - (List) Specifies the replication/archival timeouts for different type of runs(kFull, kRegular etc.).
				Nested schema for **run_timeouts**:
					* `backup_type` - (String) The scheduled backup type(kFull, kRegular etc.).
					  * Constraints: Allowable values are: `kRegular`, `kFull`, `kLog`, `kSystem`, `kHydrateCDP`, `kStorageArraySnapshot`.
					* `timeout_mins` - (Integer) Specifies the timeout in mins.
				* `schedule` - (List) Specifies a schedule fregquency and schedule unit for copying Snapshots to backup targets.
				Nested schema for **schedule**:
					* `frequency` - (Integer) Specifies a factor to multiply the unit by, to determine the copy schedule. For example if set to 2 and the unit is hourly, then Snapshots from the first eligible Job Run for every 2 hour period is copied.
					  * Constraints: The minimum value is `1`.
					* `unit` - (String) Specifies the frequency that Snapshots should be copied to the specified target. Used in combination with multiplier. <br>'Runs' means that the Snapshot copy occurs after the number of Protection Group Runs equals the number specified in the frequency. <br>'Hours' means that the Snapshot copy occurs hourly at the frequency set in the frequency, for example if scheduleFrequency is 2, the copy occurs every 2 hours. <br>'Days' means that the Snapshot copy occurs daily at the frequency set in the frequency. <br>'Weeks' means that the Snapshot copy occurs weekly at the frequency set in the frequency. <br>'Months' means that the Snapshot copy occurs monthly at the frequency set in the Frequency. <br>'Years' means that the Snapshot copy occurs yearly at the frequency set in the scheduleFrequency.
					  * Constraints: Allowable values are: `Runs`, `Hours`, `Days`, `Weeks`, `Months`, `Years`.
			* `replication_targets` - (List)
			Nested schema for **replication_targets**:
				* `aws_target_config` - (List) Specifies the configuration for adding AWS as repilcation target.
				Nested schema for **aws_target_config**:
					* `name` - (String) Specifies the name of the AWS Replication target.
					* `region` - (Integer) Specifies id of the AWS region in which to replicate the Snapshot to. Applicable if replication target is AWS target.
					* `region_name` - (String) Specifies name of the AWS region in which to replicate the Snapshot to. Applicable if replication target is AWS target.
					* `source_id` - (Integer) Specifies the source id of the AWS protection source registered on IBM cluster.
				* `azure_target_config` - (List) Specifies the configuration for adding Azure as replication target.
				Nested schema for **azure_target_config**:
					* `name` - (String) Specifies the name of the Azure Replication target.
					* `resource_group` - (Integer) Specifies id of the Azure resource group used to filter regions in UI.
					* `resource_group_name` - (String) Specifies name of the Azure resource group used to filter regions in UI.
					* `source_id` - (Integer) Specifies the source id of the Azure protection source registered on IBM cluster.
					* `storage_account` - (Integer) Specifies id of the storage account of Azure replication target which will contain storage container.
					* `storage_account_name` - (String) Specifies name of the storage account of Azure replication target which will contain storage container.
					* `storage_container` - (Integer) Specifies id of the storage container of Azure Replication target.
					* `storage_container_name` - (String) Specifies name of the storage container of Azure Replication target.
					* `storage_resource_group` - (Integer) Specifies id of the storage resource group of Azure Replication target.
					* `storage_resource_group_name` - (String) Specifies name of the storage resource group of Azure Replication target.
				* `backup_run_type` - (String) Specifies which type of run should be copied, if not set, all types of runs will be eligible for copying. If set, this will ensure that the first run of given type in the scheduled period will get copied. Currently, this can only be set to Full.
				  * Constraints: Allowable values are: `Regular`, `Full`, `Log`, `System`, `StorageArraySnapshot`.
				* `config_id` - (String) Specifies the unique identifier for the target getting added. This field need to be passed only when policies are being updated.
				* `copy_on_run_success` - (Boolean) Specifies if Snapshots are copied from the first completely successful Protection Group Run or the first partially successful Protection Group Run occurring at the start of the replication schedule. <br> If true, Snapshots are copied from the first Protection Group Run occurring at the start of the replication schedule that was completely successful i.e. Snapshots for all the Objects in the Protection Group were successfully captured. <br> If false, Snapshots are copied from the first Protection Group Run occurring at the start of the replication schedule, even if first Protection Group Run was not completely successful i.e. Snapshots were not captured for all Objects in the Protection Group.
				* `log_retention` - (List) Specifies the retention of a backup.
				Nested schema for **log_retention**:
					* `data_lock_config` - (List) Specifies WORM retention type for the snapshots. When a WORM retention type is specified, the snapshots of the Protection Groups using this policy will be kept for the last N days as specified in the duration of the datalock. During that time, the snapshots cannot be deleted.
					Nested schema for **data_lock_config**:
						* `duration` - (Integer) Specifies the duration for a dataLock. <br> Example. If duration is 7 and unit is Months, the dataLock is enabled for last 7 * 30 = 210 days of the backup.
						  * Constraints: The minimum value is `1`.
						* `enable_worm_on_external_target` - (Boolean) Specifies whether objects in the external target associated with this policy need to be made immutable.
						* `mode` - (String) Specifies the type of WORM retention type. 'Compliance' implies WORM retention is set for compliance reason. 'Administrative' implies WORM retention is set for administrative purposes.
						  * Constraints: Allowable values are: `Compliance`, `Administrative`.
						* `unit` - (String) Specificies the Retention Unit of a dataLock measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Months' then number of retention days will be 365 * 2 = 730 days.
						  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
					* `duration` - (Integer) Specifies the duration for a backup retention. <br> Example. If duration is 7 and unit is Months, the retention of a backup is 7 * 30 = 210 days.
					  * Constraints: The minimum value is `0`.
					* `unit` - (String) Specificies the Retention Unit of a backup measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Years' then number of retention days will be 365 * 2 = 730 days.
					  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
				* `remote_target_config` - (List) Specifies the configuration for adding remote cluster as repilcation target.
				Nested schema for **remote_target_config**:
					* `cluster_id` - (Integer) Specifies the cluster id of the target replication cluster.
					* `cluster_name` - (String) Specifies the cluster name of the target replication cluster.
				* `retention` - (List) Specifies the retention of a backup.
				Nested schema for **retention**:
					* `data_lock_config` - (List) Specifies WORM retention type for the snapshots. When a WORM retention type is specified, the snapshots of the Protection Groups using this policy will be kept for the last N days as specified in the duration of the datalock. During that time, the snapshots cannot be deleted.
					Nested schema for **data_lock_config**:
						* `duration` - (Integer) Specifies the duration for a dataLock. <br> Example. If duration is 7 and unit is Months, the dataLock is enabled for last 7 * 30 = 210 days of the backup.
						  * Constraints: The minimum value is `1`.
						* `enable_worm_on_external_target` - (Boolean) Specifies whether objects in the external target associated with this policy need to be made immutable.
						* `mode` - (String) Specifies the type of WORM retention type. 'Compliance' implies WORM retention is set for compliance reason. 'Administrative' implies WORM retention is set for administrative purposes.
						  * Constraints: Allowable values are: `Compliance`, `Administrative`.
						* `unit` - (String) Specificies the Retention Unit of a dataLock measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Months' then number of retention days will be 365 * 2 = 730 days.
						  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
					* `duration` - (Integer) Specifies the duration for a backup retention. <br> Example. If duration is 7 and unit is Months, the retention of a backup is 7 * 30 = 210 days.
					  * Constraints: The minimum value is `1`.
					* `unit` - (String) Specificies the Retention Unit of a backup measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Years' then number of retention days will be 365 * 2 = 730 days.
					  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
				* `run_timeouts` - (List) Specifies the replication/archival timeouts for different type of runs(kFull, kRegular etc.).
				Nested schema for **run_timeouts**:
					* `backup_type` - (String) The scheduled backup type(kFull, kRegular etc.).
					  * Constraints: Allowable values are: `kRegular`, `kFull`, `kLog`, `kSystem`, `kHydrateCDP`, `kStorageArraySnapshot`.
					* `timeout_mins` - (Integer) Specifies the timeout in mins.
				* `schedule` - (List) Specifies a schedule fregquency and schedule unit for copying Snapshots to backup targets.
				Nested schema for **schedule**:
					* `frequency` - (Integer) Specifies a factor to multiply the unit by, to determine the copy schedule. For example if set to 2 and the unit is hourly, then Snapshots from the first eligible Job Run for every 2 hour period is copied.
					  * Constraints: The minimum value is `1`.
					* `unit` - (String) Specifies the frequency that Snapshots should be copied to the specified target. Used in combination with multiplier. <br>'Runs' means that the Snapshot copy occurs after the number of Protection Group Runs equals the number specified in the frequency. <br>'Hours' means that the Snapshot copy occurs hourly at the frequency set in the frequency, for example if scheduleFrequency is 2, the copy occurs every 2 hours. <br>'Days' means that the Snapshot copy occurs daily at the frequency set in the frequency. <br>'Weeks' means that the Snapshot copy occurs weekly at the frequency set in the frequency. <br>'Months' means that the Snapshot copy occurs monthly at the frequency set in the Frequency. <br>'Years' means that the Snapshot copy occurs yearly at the frequency set in the scheduleFrequency.
					  * Constraints: Allowable values are: `Runs`, `Hours`, `Days`, `Weeks`, `Months`, `Years`.
				* `target_type` - (String) Specifies the type of target to which replication need to be performed.
				  * Constraints: Allowable values are: `RemoteCluster`, `AWS`, `Azure`.
			* `rpaas_targets` - (List)
			Nested schema for **rpaas_targets**:
				* `backup_run_type` - (String) Specifies which type of run should be copied, if not set, all types of runs will be eligible for copying. If set, this will ensure that the first run of given type in the scheduled period will get copied. Currently, this can only be set to Full.
				  * Constraints: Allowable values are: `Regular`, `Full`, `Log`, `System`, `StorageArraySnapshot`.
				* `config_id` - (String) Specifies the unique identifier for the target getting added. This field need to be passed only when policies are being updated.
				* `copy_on_run_success` - (Boolean) Specifies if Snapshots are copied from the first completely successful Protection Group Run or the first partially successful Protection Group Run occurring at the start of the replication schedule. <br> If true, Snapshots are copied from the first Protection Group Run occurring at the start of the replication schedule that was completely successful i.e. Snapshots for all the Objects in the Protection Group were successfully captured. <br> If false, Snapshots are copied from the first Protection Group Run occurring at the start of the replication schedule, even if first Protection Group Run was not completely successful i.e. Snapshots were not captured for all Objects in the Protection Group.
				* `log_retention` - (List) Specifies the retention of a backup.
				Nested schema for **log_retention**:
					* `data_lock_config` - (List) Specifies WORM retention type for the snapshots. When a WORM retention type is specified, the snapshots of the Protection Groups using this policy will be kept for the last N days as specified in the duration of the datalock. During that time, the snapshots cannot be deleted.
					Nested schema for **data_lock_config**:
						* `duration` - (Integer) Specifies the duration for a dataLock. <br> Example. If duration is 7 and unit is Months, the dataLock is enabled for last 7 * 30 = 210 days of the backup.
						  * Constraints: The minimum value is `1`.
						* `enable_worm_on_external_target` - (Boolean) Specifies whether objects in the external target associated with this policy need to be made immutable.
						* `mode` - (String) Specifies the type of WORM retention type. 'Compliance' implies WORM retention is set for compliance reason. 'Administrative' implies WORM retention is set for administrative purposes.
						  * Constraints: Allowable values are: `Compliance`, `Administrative`.
						* `unit` - (String) Specificies the Retention Unit of a dataLock measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Months' then number of retention days will be 365 * 2 = 730 days.
						  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
					* `duration` - (Integer) Specifies the duration for a backup retention. <br> Example. If duration is 7 and unit is Months, the retention of a backup is 7 * 30 = 210 days.
					  * Constraints: The minimum value is `0`.
					* `unit` - (String) Specificies the Retention Unit of a backup measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Years' then number of retention days will be 365 * 2 = 730 days.
					  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
				* `retention` - (List) Specifies the retention of a backup.
				Nested schema for **retention**:
					* `data_lock_config` - (List) Specifies WORM retention type for the snapshots. When a WORM retention type is specified, the snapshots of the Protection Groups using this policy will be kept for the last N days as specified in the duration of the datalock. During that time, the snapshots cannot be deleted.
					Nested schema for **data_lock_config**:
						* `duration` - (Integer) Specifies the duration for a dataLock. <br> Example. If duration is 7 and unit is Months, the dataLock is enabled for last 7 * 30 = 210 days of the backup.
						  * Constraints: The minimum value is `1`.
						* `enable_worm_on_external_target` - (Boolean) Specifies whether objects in the external target associated with this policy need to be made immutable.
						* `mode` - (String) Specifies the type of WORM retention type. 'Compliance' implies WORM retention is set for compliance reason. 'Administrative' implies WORM retention is set for administrative purposes.
						  * Constraints: Allowable values are: `Compliance`, `Administrative`.
						* `unit` - (String) Specificies the Retention Unit of a dataLock measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Months' then number of retention days will be 365 * 2 = 730 days.
						  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
					* `duration` - (Integer) Specifies the duration for a backup retention. <br> Example. If duration is 7 and unit is Months, the retention of a backup is 7 * 30 = 210 days.
					  * Constraints: The minimum value is `1`.
					* `unit` - (String) Specificies the Retention Unit of a backup measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Years' then number of retention days will be 365 * 2 = 730 days.
					  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
				* `run_timeouts` - (List) Specifies the replication/archival timeouts for different type of runs(kFull, kRegular etc.).
				Nested schema for **run_timeouts**:
					* `backup_type` - (String) The scheduled backup type(kFull, kRegular etc.).
					  * Constraints: Allowable values are: `kRegular`, `kFull`, `kLog`, `kSystem`, `kHydrateCDP`, `kStorageArraySnapshot`.
					* `timeout_mins` - (Integer) Specifies the timeout in mins.
				* `schedule` - (List) Specifies a schedule fregquency and schedule unit for copying Snapshots to backup targets.
				Nested schema for **schedule**:
					* `frequency` - (Integer) Specifies a factor to multiply the unit by, to determine the copy schedule. For example if set to 2 and the unit is hourly, then Snapshots from the first eligible Job Run for every 2 hour period is copied.
					  * Constraints: The minimum value is `1`.
					* `unit` - (String) Specifies the frequency that Snapshots should be copied to the specified target. Used in combination with multiplier. <br>'Runs' means that the Snapshot copy occurs after the number of Protection Group Runs equals the number specified in the frequency. <br>'Hours' means that the Snapshot copy occurs hourly at the frequency set in the frequency, for example if scheduleFrequency is 2, the copy occurs every 2 hours. <br>'Days' means that the Snapshot copy occurs daily at the frequency set in the frequency. <br>'Weeks' means that the Snapshot copy occurs weekly at the frequency set in the frequency. <br>'Months' means that the Snapshot copy occurs monthly at the frequency set in the Frequency. <br>'Years' means that the Snapshot copy occurs yearly at the frequency set in the scheduleFrequency.
					  * Constraints: Allowable values are: `Runs`, `Hours`, `Days`, `Weeks`, `Months`, `Years`.
				* `target_id` - (Integer) Specifies the RPaaS target to copy the Snapshots.
				* `target_name` - (String) Specifies the RPaaS target name where Snapshots are copied.
				* `target_type` - (String) Specifies the RPaaS target type where Snapshots are copied.
				  * Constraints: Allowable values are: `Tape`, `Cloud`, `Nas`.
		* `source_cluster_id` - (Integer) Specifies the source cluster id from where the remote operations will be performed to the next set of remote targets.
	* `data_lock` - (String) This field is now deprecated. Please use the DataLockConfig in the backup retention.
	  * Constraints: Allowable values are: `Compliance`, `Administrative`.
	* `description` - (String) Specifies the description of the Protection Policy.
	* `extended_retention` - (List) Specifies additional retention policies that should be applied to the backup snapshots. A backup snapshot will be retained up to a time that is the maximum of all retention policies that are applicable to it.
	Nested schema for **extended_retention**:
		* `config_id` - (String) Specifies the unique identifier for the target getting added. This field need to be passed olny when policies are updated.
		* `retention` - (List) Specifies the retention of a backup.
		Nested schema for **retention**:
			* `data_lock_config` - (List) Specifies WORM retention type for the snapshots. When a WORM retention type is specified, the snapshots of the Protection Groups using this policy will be kept for the last N days as specified in the duration of the datalock. During that time, the snapshots cannot be deleted.
			Nested schema for **data_lock_config**:
				* `duration` - (Integer) Specifies the duration for a dataLock. <br> Example. If duration is 7 and unit is Months, the dataLock is enabled for last 7 * 30 = 210 days of the backup.
				  * Constraints: The minimum value is `1`.
				* `enable_worm_on_external_target` - (Boolean) Specifies whether objects in the external target associated with this policy need to be made immutable.
				* `mode` - (String) Specifies the type of WORM retention type. 'Compliance' implies WORM retention is set for compliance reason. 'Administrative' implies WORM retention is set for administrative purposes.
				  * Constraints: Allowable values are: `Compliance`, `Administrative`.
				* `unit` - (String) Specificies the Retention Unit of a dataLock measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Months' then number of retention days will be 365 * 2 = 730 days.
				  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
			* `duration` - (Integer) Specifies the duration for a backup retention. <br> Example. If duration is 7 and unit is Months, the retention of a backup is 7 * 30 = 210 days.
			  * Constraints: The minimum value is `1`.
			* `unit` - (String) Specificies the Retention Unit of a backup measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Years' then number of retention days will be 365 * 2 = 730 days.
			  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
		* `run_type` - (String) The backup run type to which this extended retention applies to. If this is not set, the extended retention will be applicable to all non-log backup types. Currently, the only value that can be set here is Full.'Regular' indicates a incremental (CBT) backup. Incremental backups utilizing CBT (if supported) are captured of the target protection objects. The first run of a Regular schedule captures all the blocks.'Full' indicates a full (no CBT) backup. A complete backup (all blocks) of the target protection objects are always captured and Change Block Tracking (CBT) is not utilized.'Log' indicates a Database Log backup. Capture the database transaction logs to allow rolling back to a specific point in time.'System' indicates a system backup. System backups are used to do bare metal recovery of the system to a specific point in time.
		  * Constraints: Allowable values are: `Regular`, `Full`, `Log`, `System`, `StorageArraySnapshot`.
		* `schedule` - (List) Specifies a schedule frequency and schedule unit for Extended Retentions.
		Nested schema for **schedule**:
			* `frequency` - (Integer) Specifies a factor to multiply the unit by, to determine the retention schedule. For example if set to 2 and the unit is hourly, then Snapshots from the first eligible Job Run for every 2 hour period is retained.
			  * Constraints: The minimum value is `1`.
			* `unit` - (String) Specifies the unit interval for retention of Snapshots. <br>'Runs' means that the Snapshot copy retained after the number of Protection Group Runs equals the number specified in the frequency. <br>'Hours' means that the Snapshot copy retained hourly at the frequency set in the frequency, for example if scheduleFrequency is 2, the copy occurs every 2 hours. <br>'Days' means that the Snapshot copy gets retained daily at the frequency set in the frequency. <br>'Weeks' means that the Snapshot copy is retained weekly at the frequency set in the frequency. <br>'Months' means that the Snapshot copy is retained monthly at the frequency set in the Frequency. <br>'Years' means that the Snapshot copy is retained yearly at the frequency set in the Frequency.
			  * Constraints: Allowable values are: `Runs`, `Hours`, `Days`, `Weeks`, `Months`, `Years`.
	* `id` - (String) Specifies a unique Policy id assigned by the Cohesity Cluster.
	* `is_cbs_enabled` - (Boolean) Specifies true if Calender Based Schedule is supported by client. Default value is assumed as false for this feature.
	* `is_replicated` - (Boolean) This field is set to true when policy is the replicated policy.
	* `is_usable` - (Boolean) This field is set to true if the linked policy which is internally created from a policy templates qualifies as usable to create more policies on the cluster. If the linked policy is partially filled and can not create a working policy then this field will be set to false. In case of normal policy created on the cluster, this field wont be populated.
	* `last_modification_time_usecs` - (Integer) Specifies the last time this Policy was updated. If this is passed into a PUT request, then the backend will validate that the timestamp passed in matches the time that the policy was actually last modified. If the two timestamps do not match, then the request will be rejected with a stale error.
	* `name` - (String) Specifies the name of the Protection Policy.
	* `num_protected_objects` - (Integer) Specifies the number of protected objects using the protection policy.
	* `num_protection_groups` - (Integer) Specifies the number of protection groups using the protection policy.
	* `remote_target_policy` - (List) Specifies the replication, archival and cloud spin targets of Protection Policy.
	Nested schema for **remote_target_policy**:
		* `archival_targets` - (List)
		Nested schema for **archival_targets**:
			* `backup_run_type` - (String) Specifies which type of run should be copied, if not set, all types of runs will be eligible for copying. If set, this will ensure that the first run of given type in the scheduled period will get copied. Currently, this can only be set to Full.
			  * Constraints: Allowable values are: `Regular`, `Full`, `Log`, `System`, `StorageArraySnapshot`.
			* `config_id` - (String) Specifies the unique identifier for the target getting added. This field need to be passed only when policies are being updated.
			* `copy_on_run_success` - (Boolean) Specifies if Snapshots are copied from the first completely successful Protection Group Run or the first partially successful Protection Group Run occurring at the start of the replication schedule. <br> If true, Snapshots are copied from the first Protection Group Run occurring at the start of the replication schedule that was completely successful i.e. Snapshots for all the Objects in the Protection Group were successfully captured. <br> If false, Snapshots are copied from the first Protection Group Run occurring at the start of the replication schedule, even if first Protection Group Run was not completely successful i.e. Snapshots were not captured for all Objects in the Protection Group.
			* `extended_retention` - (List) Specifies additional retention policies that should be applied to the archived backup. Archived backup snapshot will be retained up to a time that is the maximum of all retention policies that are applicable to it.
			Nested schema for **extended_retention**:
				* `config_id` - (String) Specifies the unique identifier for the target getting added. This field need to be passed olny when policies are updated.
				* `retention` - (List) Specifies the retention of a backup.
				Nested schema for **retention**:
					* `data_lock_config` - (List) Specifies WORM retention type for the snapshots. When a WORM retention type is specified, the snapshots of the Protection Groups using this policy will be kept for the last N days as specified in the duration of the datalock. During that time, the snapshots cannot be deleted.
					Nested schema for **data_lock_config**:
						* `duration` - (Integer) Specifies the duration for a dataLock. <br> Example. If duration is 7 and unit is Months, the dataLock is enabled for last 7 * 30 = 210 days of the backup.
						  * Constraints: The minimum value is `1`.
						* `enable_worm_on_external_target` - (Boolean) Specifies whether objects in the external target associated with this policy need to be made immutable.
						* `mode` - (String) Specifies the type of WORM retention type. 'Compliance' implies WORM retention is set for compliance reason. 'Administrative' implies WORM retention is set for administrative purposes.
						  * Constraints: Allowable values are: `Compliance`, `Administrative`.
						* `unit` - (String) Specificies the Retention Unit of a dataLock measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Months' then number of retention days will be 365 * 2 = 730 days.
						  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
					* `duration` - (Integer) Specifies the duration for a backup retention. <br> Example. If duration is 7 and unit is Months, the retention of a backup is 7 * 30 = 210 days.
					  * Constraints: The minimum value is `1`.
					* `unit` - (String) Specificies the Retention Unit of a backup measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Years' then number of retention days will be 365 * 2 = 730 days.
					  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
				* `run_type` - (String) The backup run type to which this extended retention applies to. If this is not set, the extended retention will be applicable to all non-log backup types. Currently, the only value that can be set here is Full.'Regular' indicates a incremental (CBT) backup. Incremental backups utilizing CBT (if supported) are captured of the target protection objects. The first run of a Regular schedule captures all the blocks.'Full' indicates a full (no CBT) backup. A complete backup (all blocks) of the target protection objects are always captured and Change Block Tracking (CBT) is not utilized.'Log' indicates a Database Log backup. Capture the database transaction logs to allow rolling back to a specific point in time.'System' indicates a system backup. System backups are used to do bare metal recovery of the system to a specific point in time.
				  * Constraints: Allowable values are: `Regular`, `Full`, `Log`, `System`, `StorageArraySnapshot`.
				* `schedule` - (List) Specifies a schedule frequency and schedule unit for Extended Retentions.
				Nested schema for **schedule**:
					* `frequency` - (Integer) Specifies a factor to multiply the unit by, to determine the retention schedule. For example if set to 2 and the unit is hourly, then Snapshots from the first eligible Job Run for every 2 hour period is retained.
					  * Constraints: The minimum value is `1`.
					* `unit` - (String) Specifies the unit interval for retention of Snapshots. <br>'Runs' means that the Snapshot copy retained after the number of Protection Group Runs equals the number specified in the frequency. <br>'Hours' means that the Snapshot copy retained hourly at the frequency set in the frequency, for example if scheduleFrequency is 2, the copy occurs every 2 hours. <br>'Days' means that the Snapshot copy gets retained daily at the frequency set in the frequency. <br>'Weeks' means that the Snapshot copy is retained weekly at the frequency set in the frequency. <br>'Months' means that the Snapshot copy is retained monthly at the frequency set in the Frequency. <br>'Years' means that the Snapshot copy is retained yearly at the frequency set in the Frequency.
					  * Constraints: Allowable values are: `Runs`, `Hours`, `Days`, `Weeks`, `Months`, `Years`.
			* `log_retention` - (List) Specifies the retention of a backup.
			Nested schema for **log_retention**:
				* `data_lock_config` - (List) Specifies WORM retention type for the snapshots. When a WORM retention type is specified, the snapshots of the Protection Groups using this policy will be kept for the last N days as specified in the duration of the datalock. During that time, the snapshots cannot be deleted.
				Nested schema for **data_lock_config**:
					* `duration` - (Integer) Specifies the duration for a dataLock. <br> Example. If duration is 7 and unit is Months, the dataLock is enabled for last 7 * 30 = 210 days of the backup.
					  * Constraints: The minimum value is `1`.
					* `enable_worm_on_external_target` - (Boolean) Specifies whether objects in the external target associated with this policy need to be made immutable.
					* `mode` - (String) Specifies the type of WORM retention type. 'Compliance' implies WORM retention is set for compliance reason. 'Administrative' implies WORM retention is set for administrative purposes.
					  * Constraints: Allowable values are: `Compliance`, `Administrative`.
					* `unit` - (String) Specificies the Retention Unit of a dataLock measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Months' then number of retention days will be 365 * 2 = 730 days.
					  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
				* `duration` - (Integer) Specifies the duration for a backup retention. <br> Example. If duration is 7 and unit is Months, the retention of a backup is 7 * 30 = 210 days.
				  * Constraints: The minimum value is `0`.
				* `unit` - (String) Specificies the Retention Unit of a backup measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Years' then number of retention days will be 365 * 2 = 730 days.
				  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
			* `retention` - (List) Specifies the retention of a backup.
			Nested schema for **retention**:
				* `data_lock_config` - (List) Specifies WORM retention type for the snapshots. When a WORM retention type is specified, the snapshots of the Protection Groups using this policy will be kept for the last N days as specified in the duration of the datalock. During that time, the snapshots cannot be deleted.
				Nested schema for **data_lock_config**:
					* `duration` - (Integer) Specifies the duration for a dataLock. <br> Example. If duration is 7 and unit is Months, the dataLock is enabled for last 7 * 30 = 210 days of the backup.
					  * Constraints: The minimum value is `1`.
					* `enable_worm_on_external_target` - (Boolean) Specifies whether objects in the external target associated with this policy need to be made immutable.
					* `mode` - (String) Specifies the type of WORM retention type. 'Compliance' implies WORM retention is set for compliance reason. 'Administrative' implies WORM retention is set for administrative purposes.
					  * Constraints: Allowable values are: `Compliance`, `Administrative`.
					* `unit` - (String) Specificies the Retention Unit of a dataLock measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Months' then number of retention days will be 365 * 2 = 730 days.
					  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
				* `duration` - (Integer) Specifies the duration for a backup retention. <br> Example. If duration is 7 and unit is Months, the retention of a backup is 7 * 30 = 210 days.
				  * Constraints: The minimum value is `1`.
				* `unit` - (String) Specificies the Retention Unit of a backup measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Years' then number of retention days will be 365 * 2 = 730 days.
				  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
			* `run_timeouts` - (List) Specifies the replication/archival timeouts for different type of runs(kFull, kRegular etc.).
			Nested schema for **run_timeouts**:
				* `backup_type` - (String) The scheduled backup type(kFull, kRegular etc.).
				  * Constraints: Allowable values are: `kRegular`, `kFull`, `kLog`, `kSystem`, `kHydrateCDP`, `kStorageArraySnapshot`.
				* `timeout_mins` - (Integer) Specifies the timeout in mins.
			* `schedule` - (List) Specifies a schedule fregquency and schedule unit for copying Snapshots to backup targets.
			Nested schema for **schedule**:
				* `frequency` - (Integer) Specifies a factor to multiply the unit by, to determine the copy schedule. For example if set to 2 and the unit is hourly, then Snapshots from the first eligible Job Run for every 2 hour period is copied.
				  * Constraints: The minimum value is `1`.
				* `unit` - (String) Specifies the frequency that Snapshots should be copied to the specified target. Used in combination with multiplier. <br>'Runs' means that the Snapshot copy occurs after the number of Protection Group Runs equals the number specified in the frequency. <br>'Hours' means that the Snapshot copy occurs hourly at the frequency set in the frequency, for example if scheduleFrequency is 2, the copy occurs every 2 hours. <br>'Days' means that the Snapshot copy occurs daily at the frequency set in the frequency. <br>'Weeks' means that the Snapshot copy occurs weekly at the frequency set in the frequency. <br>'Months' means that the Snapshot copy occurs monthly at the frequency set in the Frequency. <br>'Years' means that the Snapshot copy occurs yearly at the frequency set in the scheduleFrequency.
				  * Constraints: Allowable values are: `Runs`, `Hours`, `Days`, `Weeks`, `Months`, `Years`.
			* `target_id` - (Integer) Specifies the Archival target to copy the Snapshots to.
			* `target_name` - (String) Specifies the Archival target name where Snapshots are copied.
			* `target_type` - (String) Specifies the Archival target type where Snapshots are copied.
			  * Constraints: Allowable values are: `Tape`, `Cloud`, `Nas`.
			* `tier_settings` - (List) Specifies the settings tier levels configured with each archival target. The tier settings need to be applied in specific order and default tier should always be passed as first entry in tiers array. The following example illustrates how to configure tiering input for AWS tiering. Same type of input structure applied to other cloud platforms also. <br>If user wants to achieve following tiering for backup, <br>User Desired Tiering- <br><t>1.Archive Full back up for 12 Months <br><t>2.Tier Levels <br><t><t>[1,12] [ <br><t><t><t>s3 (1 to 2 months), (default tier) <br><t><t><t>s3 Intelligent tiering (3 to 6 months), <br><t><t><t>s3 One Zone (7 to 9 months) <br><t><t><t>Glacier (10 to 12 months)] <br><t>API Input <br><t><t>1.tiers-[ <br><t><t><t>{'tierType': 'S3','moveAfterUnit':'months', <br><t><t><t>'moveAfter':2 - move from s3 to s3Inte after 2 months}, <br><t><t><t>{'tierType': 'S3Inte','moveAfterUnit':'months', <br><t><t><t>'moveAfter':4 - move from S3Inte to Glacier after 4 months}, <br><t><t><t>{'tierType': 'Glacier', 'moveAfterUnit':'months', <br><t><t><t>'moveAfter': 3 - move from Glacier to S3 One Zone after 3 months }, <br><t><t><t>{'tierType': 'S3 One Zone', 'moveAfterUnit': nil, <br><t><t><t>'moveAfter': nil - For the last record, 'moveAfter' and 'moveAfterUnit' <br><t><t><t>will be ignored since there are no further tier for data movement } <br><t><t><t>}].
			Nested schema for **tier_settings**:
				* `aws_tiering` - (List) Specifies aws tiers.
				Nested schema for **aws_tiering**:
					* `tiers` - (List) Specifies the tiers that are used to move the archived backup from current tier to next tier. The order of the tiers determines which tier will be used next for moving the archived backup. The first tier input should always be default tier where backup will be acrhived. Each tier specifies how much time after the backup will be moved to next tier from the current tier.
					Nested schema for **tiers**:
						* `move_after` - (Integer) Specifies the time period after which the backup will be moved from current tier to next tier.
						* `move_after_unit` - (String) Specifies the unit for moving the data from current tier to next tier. This unit will be a base unit for the 'moveAfter' field specified below.
						  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
						* `tier_type` - (String) Specifies the AWS tier types.
						  * Constraints: Allowable values are: `kAmazonS3Standard`, `kAmazonS3StandardIA`, `kAmazonS3OneZoneIA`, `kAmazonS3IntelligentTiering`, `kAmazonS3Glacier`, `kAmazonS3GlacierDeepArchive`.
				* `azure_tiering` - (List) Specifies Azure tiers.
				Nested schema for **azure_tiering**:
					* `tiers` - (List) Specifies the tiers that are used to move the archived backup from current tier to next tier. The order of the tiers determines which tier will be used next for moving the archived backup. The first tier input should always be default tier where backup will be acrhived. Each tier specifies how much time after the backup will be moved to next tier from the current tier.
					Nested schema for **tiers**:
						* `move_after` - (Integer) Specifies the time period after which the backup will be moved from current tier to next tier.
						* `move_after_unit` - (String) Specifies the unit for moving the data from current tier to next tier. This unit will be a base unit for the 'moveAfter' field specified below.
						  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
						* `tier_type` - (String) Specifies the Azure tier types.
						  * Constraints: Allowable values are: `kAzureTierHot`, `kAzureTierCool`, `kAzureTierArchive`.
				* `cloud_platform` - (String) Specifies the cloud platform to enable tiering.
				  * Constraints: Allowable values are: `AWS`, `Azure`, `Oracle`, `Google`.
				* `google_tiering` - (List) Specifies Google tiers.
				Nested schema for **google_tiering**:
					* `tiers` - (List) Specifies the tiers that are used to move the archived backup from current tier to next tier. The order of the tiers determines which tier will be used next for moving the archived backup. The first tier input should always be default tier where backup will be acrhived. Each tier specifies how much time after the backup will be moved to next tier from the current tier.
					Nested schema for **tiers**:
						* `move_after` - (Integer) Specifies the time period after which the backup will be moved from current tier to next tier.
						* `move_after_unit` - (String) Specifies the unit for moving the data from current tier to next tier. This unit will be a base unit for the 'moveAfter' field specified below.
						  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
						* `tier_type` - (String) Specifies the Google tier types.
						  * Constraints: Allowable values are: `kGoogleStandard`, `kGoogleRegional`, `kGoogleMultiRegional`, `kGoogleNearline`, `kGoogleColdline`.
				* `oracle_tiering` - (List) Specifies Oracle tiers.
				Nested schema for **oracle_tiering**:
					* `tiers` - (List) Specifies the tiers that are used to move the archived backup from current tier to next tier. The order of the tiers determines which tier will be used next for moving the archived backup. The first tier input should always be default tier where backup will be acrhived. Each tier specifies how much time after the backup will be moved to next tier from the current tier.
					Nested schema for **tiers**:
						* `move_after` - (Integer) Specifies the time period after which the backup will be moved from current tier to next tier.
						* `move_after_unit` - (String) Specifies the unit for moving the data from current tier to next tier. This unit will be a base unit for the 'moveAfter' field specified below.
						  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
						* `tier_type` - (String) Specifies the Oracle tier types.
						  * Constraints: Allowable values are: `kOracleTierStandard`, `kOracleTierArchive`.
		* `cloud_spin_targets` - (List)
		Nested schema for **cloud_spin_targets**:
			* `backup_run_type` - (String) Specifies which type of run should be copied, if not set, all types of runs will be eligible for copying. If set, this will ensure that the first run of given type in the scheduled period will get copied. Currently, this can only be set to Full.
			  * Constraints: Allowable values are: `Regular`, `Full`, `Log`, `System`, `StorageArraySnapshot`.
			* `config_id` - (String) Specifies the unique identifier for the target getting added. This field need to be passed only when policies are being updated.
			* `copy_on_run_success` - (Boolean) Specifies if Snapshots are copied from the first completely successful Protection Group Run or the first partially successful Protection Group Run occurring at the start of the replication schedule. <br> If true, Snapshots are copied from the first Protection Group Run occurring at the start of the replication schedule that was completely successful i.e. Snapshots for all the Objects in the Protection Group were successfully captured. <br> If false, Snapshots are copied from the first Protection Group Run occurring at the start of the replication schedule, even if first Protection Group Run was not completely successful i.e. Snapshots were not captured for all Objects in the Protection Group.
			* `log_retention` - (List) Specifies the retention of a backup.
			Nested schema for **log_retention**:
				* `data_lock_config` - (List) Specifies WORM retention type for the snapshots. When a WORM retention type is specified, the snapshots of the Protection Groups using this policy will be kept for the last N days as specified in the duration of the datalock. During that time, the snapshots cannot be deleted.
				Nested schema for **data_lock_config**:
					* `duration` - (Integer) Specifies the duration for a dataLock. <br> Example. If duration is 7 and unit is Months, the dataLock is enabled for last 7 * 30 = 210 days of the backup.
					  * Constraints: The minimum value is `1`.
					* `enable_worm_on_external_target` - (Boolean) Specifies whether objects in the external target associated with this policy need to be made immutable.
					* `mode` - (String) Specifies the type of WORM retention type. 'Compliance' implies WORM retention is set for compliance reason. 'Administrative' implies WORM retention is set for administrative purposes.
					  * Constraints: Allowable values are: `Compliance`, `Administrative`.
					* `unit` - (String) Specificies the Retention Unit of a dataLock measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Months' then number of retention days will be 365 * 2 = 730 days.
					  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
				* `duration` - (Integer) Specifies the duration for a backup retention. <br> Example. If duration is 7 and unit is Months, the retention of a backup is 7 * 30 = 210 days.
				  * Constraints: The minimum value is `0`.
				* `unit` - (String) Specificies the Retention Unit of a backup measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Years' then number of retention days will be 365 * 2 = 730 days.
				  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
			* `retention` - (List) Specifies the retention of a backup.
			Nested schema for **retention**:
				* `data_lock_config` - (List) Specifies WORM retention type for the snapshots. When a WORM retention type is specified, the snapshots of the Protection Groups using this policy will be kept for the last N days as specified in the duration of the datalock. During that time, the snapshots cannot be deleted.
				Nested schema for **data_lock_config**:
					* `duration` - (Integer) Specifies the duration for a dataLock. <br> Example. If duration is 7 and unit is Months, the dataLock is enabled for last 7 * 30 = 210 days of the backup.
					  * Constraints: The minimum value is `1`.
					* `enable_worm_on_external_target` - (Boolean) Specifies whether objects in the external target associated with this policy need to be made immutable.
					* `mode` - (String) Specifies the type of WORM retention type. 'Compliance' implies WORM retention is set for compliance reason. 'Administrative' implies WORM retention is set for administrative purposes.
					  * Constraints: Allowable values are: `Compliance`, `Administrative`.
					* `unit` - (String) Specificies the Retention Unit of a dataLock measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Months' then number of retention days will be 365 * 2 = 730 days.
					  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
				* `duration` - (Integer) Specifies the duration for a backup retention. <br> Example. If duration is 7 and unit is Months, the retention of a backup is 7 * 30 = 210 days.
				  * Constraints: The minimum value is `1`.
				* `unit` - (String) Specificies the Retention Unit of a backup measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Years' then number of retention days will be 365 * 2 = 730 days.
				  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
			* `run_timeouts` - (List) Specifies the replication/archival timeouts for different type of runs(kFull, kRegular etc.).
			Nested schema for **run_timeouts**:
				* `backup_type` - (String) The scheduled backup type(kFull, kRegular etc.).
				  * Constraints: Allowable values are: `kRegular`, `kFull`, `kLog`, `kSystem`, `kHydrateCDP`, `kStorageArraySnapshot`.
				* `timeout_mins` - (Integer) Specifies the timeout in mins.
			* `schedule` - (List) Specifies a schedule fregquency and schedule unit for copying Snapshots to backup targets.
			Nested schema for **schedule**:
				* `frequency` - (Integer) Specifies a factor to multiply the unit by, to determine the copy schedule. For example if set to 2 and the unit is hourly, then Snapshots from the first eligible Job Run for every 2 hour period is copied.
				  * Constraints: The minimum value is `1`.
				* `unit` - (String) Specifies the frequency that Snapshots should be copied to the specified target. Used in combination with multiplier. <br>'Runs' means that the Snapshot copy occurs after the number of Protection Group Runs equals the number specified in the frequency. <br>'Hours' means that the Snapshot copy occurs hourly at the frequency set in the frequency, for example if scheduleFrequency is 2, the copy occurs every 2 hours. <br>'Days' means that the Snapshot copy occurs daily at the frequency set in the frequency. <br>'Weeks' means that the Snapshot copy occurs weekly at the frequency set in the frequency. <br>'Months' means that the Snapshot copy occurs monthly at the frequency set in the Frequency. <br>'Years' means that the Snapshot copy occurs yearly at the frequency set in the scheduleFrequency.
				  * Constraints: Allowable values are: `Runs`, `Hours`, `Days`, `Weeks`, `Months`, `Years`.
			* `target` - (List) Specifies the details about Cloud Spin target where backup snapshots may be converted and stored.
			Nested schema for **target**:
				* `aws_params` - (List) Specifies various resources when converting and deploying a VM to AWS.
				Nested schema for **aws_params**:
					* `custom_tag_list` - (List) Specifies tags of various resources when converting and deploying a VM to AWS.
					Nested schema for **custom_tag_list**:
						* `key` - (String) Specifies key of the custom tag.
						* `value` - (String) Specifies value of the custom tag.
					* `region` - (Integer) Specifies id of the AWS region in which to deploy the VM.
					* `subnet_id` - (Integer) Specifies id of the subnet within above VPC.
					* `vpc_id` - (Integer) Specifies id of the Virtual Private Cloud to chose for the instance type.
				* `azure_params` - (List) Specifies various resources when converting and deploying a VM to Azure.
				Nested schema for **azure_params**:
					* `availability_set_id` - (Integer) Specifies the availability set.
					* `network_resource_group_id` - (Integer) Specifies id of the resource group for the selected virtual network.
					* `resource_group_id` - (Integer) Specifies id of the Azure resource group. Its value is globally unique within Azure.
					* `storage_account_id` - (Integer) Specifies id of the storage account that will contain the storage container within which we will create the blob that will become the VHD disk for the cloned VM.
					* `storage_container_id` - (Integer) Specifies id of the storage container within the above storage account.
					* `storage_resource_group_id` - (Integer) Specifies id of the resource group for the selected storage account.
					* `temp_vm_resource_group_id` - (Integer) Specifies id of the temporary Azure resource group.
					* `temp_vm_storage_account_id` - (Integer) Specifies id of the temporary VM storage account that will contain the storage container within which we will create the blob that will become the VHD disk for the cloned VM.
					* `temp_vm_storage_container_id` - (Integer) Specifies id of the temporary VM storage container within the above storage account.
					* `temp_vm_subnet_id` - (Integer) Specifies Id of the temporary VM subnet within the above virtual network.
					* `temp_vm_virtual_network_id` - (Integer) Specifies Id of the temporary VM Virtual Network.
				* `id` - (Integer) Specifies the unique id of the cloud spin entity.
				* `name` - (String) Specifies the name of the already added cloud spin target.
		* `onprem_deploy_targets` - (List)
		Nested schema for **onprem_deploy_targets**:
			* `backup_run_type` - (String) Specifies which type of run should be copied, if not set, all types of runs will be eligible for copying. If set, this will ensure that the first run of given type in the scheduled period will get copied. Currently, this can only be set to Full.
			  * Constraints: Allowable values are: `Regular`, `Full`, `Log`, `System`, `StorageArraySnapshot`.
			* `config_id` - (String) Specifies the unique identifier for the target getting added. This field need to be passed only when policies are being updated.
			* `copy_on_run_success` - (Boolean) Specifies if Snapshots are copied from the first completely successful Protection Group Run or the first partially successful Protection Group Run occurring at the start of the replication schedule. <br> If true, Snapshots are copied from the first Protection Group Run occurring at the start of the replication schedule that was completely successful i.e. Snapshots for all the Objects in the Protection Group were successfully captured. <br> If false, Snapshots are copied from the first Protection Group Run occurring at the start of the replication schedule, even if first Protection Group Run was not completely successful i.e. Snapshots were not captured for all Objects in the Protection Group.
			* `log_retention` - (List) Specifies the retention of a backup.
			Nested schema for **log_retention**:
				* `data_lock_config` - (List) Specifies WORM retention type for the snapshots. When a WORM retention type is specified, the snapshots of the Protection Groups using this policy will be kept for the last N days as specified in the duration of the datalock. During that time, the snapshots cannot be deleted.
				Nested schema for **data_lock_config**:
					* `duration` - (Integer) Specifies the duration for a dataLock. <br> Example. If duration is 7 and unit is Months, the dataLock is enabled for last 7 * 30 = 210 days of the backup.
					  * Constraints: The minimum value is `1`.
					* `enable_worm_on_external_target` - (Boolean) Specifies whether objects in the external target associated with this policy need to be made immutable.
					* `mode` - (String) Specifies the type of WORM retention type. 'Compliance' implies WORM retention is set for compliance reason. 'Administrative' implies WORM retention is set for administrative purposes.
					  * Constraints: Allowable values are: `Compliance`, `Administrative`.
					* `unit` - (String) Specificies the Retention Unit of a dataLock measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Months' then number of retention days will be 365 * 2 = 730 days.
					  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
				* `duration` - (Integer) Specifies the duration for a backup retention. <br> Example. If duration is 7 and unit is Months, the retention of a backup is 7 * 30 = 210 days.
				  * Constraints: The minimum value is `0`.
				* `unit` - (String) Specificies the Retention Unit of a backup measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Years' then number of retention days will be 365 * 2 = 730 days.
				  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
			* `params` - (List) Specifies the details about OnpremDeploy target where backup snapshots may be converted and deployed.
			Nested schema for **params**:
				* `id` - (Integer) Specifies the unique id of the onprem entity.
			* `retention` - (List) Specifies the retention of a backup.
			Nested schema for **retention**:
				* `data_lock_config` - (List) Specifies WORM retention type for the snapshots. When a WORM retention type is specified, the snapshots of the Protection Groups using this policy will be kept for the last N days as specified in the duration of the datalock. During that time, the snapshots cannot be deleted.
				Nested schema for **data_lock_config**:
					* `duration` - (Integer) Specifies the duration for a dataLock. <br> Example. If duration is 7 and unit is Months, the dataLock is enabled for last 7 * 30 = 210 days of the backup.
					  * Constraints: The minimum value is `1`.
					* `enable_worm_on_external_target` - (Boolean) Specifies whether objects in the external target associated with this policy need to be made immutable.
					* `mode` - (String) Specifies the type of WORM retention type. 'Compliance' implies WORM retention is set for compliance reason. 'Administrative' implies WORM retention is set for administrative purposes.
					  * Constraints: Allowable values are: `Compliance`, `Administrative`.
					* `unit` - (String) Specificies the Retention Unit of a dataLock measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Months' then number of retention days will be 365 * 2 = 730 days.
					  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
				* `duration` - (Integer) Specifies the duration for a backup retention. <br> Example. If duration is 7 and unit is Months, the retention of a backup is 7 * 30 = 210 days.
				  * Constraints: The minimum value is `1`.
				* `unit` - (String) Specificies the Retention Unit of a backup measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Years' then number of retention days will be 365 * 2 = 730 days.
				  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
			* `run_timeouts` - (List) Specifies the replication/archival timeouts for different type of runs(kFull, kRegular etc.).
			Nested schema for **run_timeouts**:
				* `backup_type` - (String) The scheduled backup type(kFull, kRegular etc.).
				  * Constraints: Allowable values are: `kRegular`, `kFull`, `kLog`, `kSystem`, `kHydrateCDP`, `kStorageArraySnapshot`.
				* `timeout_mins` - (Integer) Specifies the timeout in mins.
			* `schedule` - (List) Specifies a schedule fregquency and schedule unit for copying Snapshots to backup targets.
			Nested schema for **schedule**:
				* `frequency` - (Integer) Specifies a factor to multiply the unit by, to determine the copy schedule. For example if set to 2 and the unit is hourly, then Snapshots from the first eligible Job Run for every 2 hour period is copied.
				  * Constraints: The minimum value is `1`.
				* `unit` - (String) Specifies the frequency that Snapshots should be copied to the specified target. Used in combination with multiplier. <br>'Runs' means that the Snapshot copy occurs after the number of Protection Group Runs equals the number specified in the frequency. <br>'Hours' means that the Snapshot copy occurs hourly at the frequency set in the frequency, for example if scheduleFrequency is 2, the copy occurs every 2 hours. <br>'Days' means that the Snapshot copy occurs daily at the frequency set in the frequency. <br>'Weeks' means that the Snapshot copy occurs weekly at the frequency set in the frequency. <br>'Months' means that the Snapshot copy occurs monthly at the frequency set in the Frequency. <br>'Years' means that the Snapshot copy occurs yearly at the frequency set in the scheduleFrequency.
				  * Constraints: Allowable values are: `Runs`, `Hours`, `Days`, `Weeks`, `Months`, `Years`.
		* `replication_targets` - (List)
		Nested schema for **replication_targets**:
			* `aws_target_config` - (List) Specifies the configuration for adding AWS as repilcation target.
			Nested schema for **aws_target_config**:
				* `name` - (String) Specifies the name of the AWS Replication target.
				* `region` - (Integer) Specifies id of the AWS region in which to replicate the Snapshot to. Applicable if replication target is AWS target.
				* `region_name` - (String) Specifies name of the AWS region in which to replicate the Snapshot to. Applicable if replication target is AWS target.
				* `source_id` - (Integer) Specifies the source id of the AWS protection source registered on IBM cluster.
			* `azure_target_config` - (List) Specifies the configuration for adding Azure as replication target.
			Nested schema for **azure_target_config**:
				* `name` - (String) Specifies the name of the Azure Replication target.
				* `resource_group` - (Integer) Specifies id of the Azure resource group used to filter regions in UI.
				* `resource_group_name` - (String) Specifies name of the Azure resource group used to filter regions in UI.
				* `source_id` - (Integer) Specifies the source id of the Azure protection source registered on IBM cluster.
				* `storage_account` - (Integer) Specifies id of the storage account of Azure replication target which will contain storage container.
				* `storage_account_name` - (String) Specifies name of the storage account of Azure replication target which will contain storage container.
				* `storage_container` - (Integer) Specifies id of the storage container of Azure Replication target.
				* `storage_container_name` - (String) Specifies name of the storage container of Azure Replication target.
				* `storage_resource_group` - (Integer) Specifies id of the storage resource group of Azure Replication target.
				* `storage_resource_group_name` - (String) Specifies name of the storage resource group of Azure Replication target.
			* `backup_run_type` - (String) Specifies which type of run should be copied, if not set, all types of runs will be eligible for copying. If set, this will ensure that the first run of given type in the scheduled period will get copied. Currently, this can only be set to Full.
			  * Constraints: Allowable values are: `Regular`, `Full`, `Log`, `System`, `StorageArraySnapshot`.
			* `config_id` - (String) Specifies the unique identifier for the target getting added. This field need to be passed only when policies are being updated.
			* `copy_on_run_success` - (Boolean) Specifies if Snapshots are copied from the first completely successful Protection Group Run or the first partially successful Protection Group Run occurring at the start of the replication schedule. <br> If true, Snapshots are copied from the first Protection Group Run occurring at the start of the replication schedule that was completely successful i.e. Snapshots for all the Objects in the Protection Group were successfully captured. <br> If false, Snapshots are copied from the first Protection Group Run occurring at the start of the replication schedule, even if first Protection Group Run was not completely successful i.e. Snapshots were not captured for all Objects in the Protection Group.
			* `log_retention` - (List) Specifies the retention of a backup.
			Nested schema for **log_retention**:
				* `data_lock_config` - (List) Specifies WORM retention type for the snapshots. When a WORM retention type is specified, the snapshots of the Protection Groups using this policy will be kept for the last N days as specified in the duration of the datalock. During that time, the snapshots cannot be deleted.
				Nested schema for **data_lock_config**:
					* `duration` - (Integer) Specifies the duration for a dataLock. <br> Example. If duration is 7 and unit is Months, the dataLock is enabled for last 7 * 30 = 210 days of the backup.
					  * Constraints: The minimum value is `1`.
					* `enable_worm_on_external_target` - (Boolean) Specifies whether objects in the external target associated with this policy need to be made immutable.
					* `mode` - (String) Specifies the type of WORM retention type. 'Compliance' implies WORM retention is set for compliance reason. 'Administrative' implies WORM retention is set for administrative purposes.
					  * Constraints: Allowable values are: `Compliance`, `Administrative`.
					* `unit` - (String) Specificies the Retention Unit of a dataLock measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Months' then number of retention days will be 365 * 2 = 730 days.
					  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
				* `duration` - (Integer) Specifies the duration for a backup retention. <br> Example. If duration is 7 and unit is Months, the retention of a backup is 7 * 30 = 210 days.
				  * Constraints: The minimum value is `0`.
				* `unit` - (String) Specificies the Retention Unit of a backup measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Years' then number of retention days will be 365 * 2 = 730 days.
				  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
			* `remote_target_config` - (List) Specifies the configuration for adding remote cluster as repilcation target.
			Nested schema for **remote_target_config**:
				* `cluster_id` - (Integer) Specifies the cluster id of the target replication cluster.
				* `cluster_name` - (String) Specifies the cluster name of the target replication cluster.
			* `retention` - (List) Specifies the retention of a backup.
			Nested schema for **retention**:
				* `data_lock_config` - (List) Specifies WORM retention type for the snapshots. When a WORM retention type is specified, the snapshots of the Protection Groups using this policy will be kept for the last N days as specified in the duration of the datalock. During that time, the snapshots cannot be deleted.
				Nested schema for **data_lock_config**:
					* `duration` - (Integer) Specifies the duration for a dataLock. <br> Example. If duration is 7 and unit is Months, the dataLock is enabled for last 7 * 30 = 210 days of the backup.
					  * Constraints: The minimum value is `1`.
					* `enable_worm_on_external_target` - (Boolean) Specifies whether objects in the external target associated with this policy need to be made immutable.
					* `mode` - (String) Specifies the type of WORM retention type. 'Compliance' implies WORM retention is set for compliance reason. 'Administrative' implies WORM retention is set for administrative purposes.
					  * Constraints: Allowable values are: `Compliance`, `Administrative`.
					* `unit` - (String) Specificies the Retention Unit of a dataLock measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Months' then number of retention days will be 365 * 2 = 730 days.
					  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
				* `duration` - (Integer) Specifies the duration for a backup retention. <br> Example. If duration is 7 and unit is Months, the retention of a backup is 7 * 30 = 210 days.
				  * Constraints: The minimum value is `1`.
				* `unit` - (String) Specificies the Retention Unit of a backup measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Years' then number of retention days will be 365 * 2 = 730 days.
				  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
			* `run_timeouts` - (List) Specifies the replication/archival timeouts for different type of runs(kFull, kRegular etc.).
			Nested schema for **run_timeouts**:
				* `backup_type` - (String) The scheduled backup type(kFull, kRegular etc.).
				  * Constraints: Allowable values are: `kRegular`, `kFull`, `kLog`, `kSystem`, `kHydrateCDP`, `kStorageArraySnapshot`.
				* `timeout_mins` - (Integer) Specifies the timeout in mins.
			* `schedule` - (List) Specifies a schedule fregquency and schedule unit for copying Snapshots to backup targets.
			Nested schema for **schedule**:
				* `frequency` - (Integer) Specifies a factor to multiply the unit by, to determine the copy schedule. For example if set to 2 and the unit is hourly, then Snapshots from the first eligible Job Run for every 2 hour period is copied.
				  * Constraints: The minimum value is `1`.
				* `unit` - (String) Specifies the frequency that Snapshots should be copied to the specified target. Used in combination with multiplier. <br>'Runs' means that the Snapshot copy occurs after the number of Protection Group Runs equals the number specified in the frequency. <br>'Hours' means that the Snapshot copy occurs hourly at the frequency set in the frequency, for example if scheduleFrequency is 2, the copy occurs every 2 hours. <br>'Days' means that the Snapshot copy occurs daily at the frequency set in the frequency. <br>'Weeks' means that the Snapshot copy occurs weekly at the frequency set in the frequency. <br>'Months' means that the Snapshot copy occurs monthly at the frequency set in the Frequency. <br>'Years' means that the Snapshot copy occurs yearly at the frequency set in the scheduleFrequency.
				  * Constraints: Allowable values are: `Runs`, `Hours`, `Days`, `Weeks`, `Months`, `Years`.
			* `target_type` - (String) Specifies the type of target to which replication need to be performed.
			  * Constraints: Allowable values are: `RemoteCluster`, `AWS`, `Azure`.
		* `rpaas_targets` - (List)
		Nested schema for **rpaas_targets**:
			* `backup_run_type` - (String) Specifies which type of run should be copied, if not set, all types of runs will be eligible for copying. If set, this will ensure that the first run of given type in the scheduled period will get copied. Currently, this can only be set to Full.
			  * Constraints: Allowable values are: `Regular`, `Full`, `Log`, `System`, `StorageArraySnapshot`.
			* `config_id` - (String) Specifies the unique identifier for the target getting added. This field need to be passed only when policies are being updated.
			* `copy_on_run_success` - (Boolean) Specifies if Snapshots are copied from the first completely successful Protection Group Run or the first partially successful Protection Group Run occurring at the start of the replication schedule. <br> If true, Snapshots are copied from the first Protection Group Run occurring at the start of the replication schedule that was completely successful i.e. Snapshots for all the Objects in the Protection Group were successfully captured. <br> If false, Snapshots are copied from the first Protection Group Run occurring at the start of the replication schedule, even if first Protection Group Run was not completely successful i.e. Snapshots were not captured for all Objects in the Protection Group.
			* `log_retention` - (List) Specifies the retention of a backup.
			Nested schema for **log_retention**:
				* `data_lock_config` - (List) Specifies WORM retention type for the snapshots. When a WORM retention type is specified, the snapshots of the Protection Groups using this policy will be kept for the last N days as specified in the duration of the datalock. During that time, the snapshots cannot be deleted.
				Nested schema for **data_lock_config**:
					* `duration` - (Integer) Specifies the duration for a dataLock. <br> Example. If duration is 7 and unit is Months, the dataLock is enabled for last 7 * 30 = 210 days of the backup.
					  * Constraints: The minimum value is `1`.
					* `enable_worm_on_external_target` - (Boolean) Specifies whether objects in the external target associated with this policy need to be made immutable.
					* `mode` - (String) Specifies the type of WORM retention type. 'Compliance' implies WORM retention is set for compliance reason. 'Administrative' implies WORM retention is set for administrative purposes.
					  * Constraints: Allowable values are: `Compliance`, `Administrative`.
					* `unit` - (String) Specificies the Retention Unit of a dataLock measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Months' then number of retention days will be 365 * 2 = 730 days.
					  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
				* `duration` - (Integer) Specifies the duration for a backup retention. <br> Example. If duration is 7 and unit is Months, the retention of a backup is 7 * 30 = 210 days.
				  * Constraints: The minimum value is `0`.
				* `unit` - (String) Specificies the Retention Unit of a backup measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Years' then number of retention days will be 365 * 2 = 730 days.
				  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
			* `retention` - (List) Specifies the retention of a backup.
			Nested schema for **retention**:
				* `data_lock_config` - (List) Specifies WORM retention type for the snapshots. When a WORM retention type is specified, the snapshots of the Protection Groups using this policy will be kept for the last N days as specified in the duration of the datalock. During that time, the snapshots cannot be deleted.
				Nested schema for **data_lock_config**:
					* `duration` - (Integer) Specifies the duration for a dataLock. <br> Example. If duration is 7 and unit is Months, the dataLock is enabled for last 7 * 30 = 210 days of the backup.
					  * Constraints: The minimum value is `1`.
					* `enable_worm_on_external_target` - (Boolean) Specifies whether objects in the external target associated with this policy need to be made immutable.
					* `mode` - (String) Specifies the type of WORM retention type. 'Compliance' implies WORM retention is set for compliance reason. 'Administrative' implies WORM retention is set for administrative purposes.
					  * Constraints: Allowable values are: `Compliance`, `Administrative`.
					* `unit` - (String) Specificies the Retention Unit of a dataLock measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Months' then number of retention days will be 365 * 2 = 730 days.
					  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
				* `duration` - (Integer) Specifies the duration for a backup retention. <br> Example. If duration is 7 and unit is Months, the retention of a backup is 7 * 30 = 210 days.
				  * Constraints: The minimum value is `1`.
				* `unit` - (String) Specificies the Retention Unit of a backup measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Years' then number of retention days will be 365 * 2 = 730 days.
				  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
			* `run_timeouts` - (List) Specifies the replication/archival timeouts for different type of runs(kFull, kRegular etc.).
			Nested schema for **run_timeouts**:
				* `backup_type` - (String) The scheduled backup type(kFull, kRegular etc.).
				  * Constraints: Allowable values are: `kRegular`, `kFull`, `kLog`, `kSystem`, `kHydrateCDP`, `kStorageArraySnapshot`.
				* `timeout_mins` - (Integer) Specifies the timeout in mins.
			* `schedule` - (List) Specifies a schedule fregquency and schedule unit for copying Snapshots to backup targets.
			Nested schema for **schedule**:
				* `frequency` - (Integer) Specifies a factor to multiply the unit by, to determine the copy schedule. For example if set to 2 and the unit is hourly, then Snapshots from the first eligible Job Run for every 2 hour period is copied.
				  * Constraints: The minimum value is `1`.
				* `unit` - (String) Specifies the frequency that Snapshots should be copied to the specified target. Used in combination with multiplier. <br>'Runs' means that the Snapshot copy occurs after the number of Protection Group Runs equals the number specified in the frequency. <br>'Hours' means that the Snapshot copy occurs hourly at the frequency set in the frequency, for example if scheduleFrequency is 2, the copy occurs every 2 hours. <br>'Days' means that the Snapshot copy occurs daily at the frequency set in the frequency. <br>'Weeks' means that the Snapshot copy occurs weekly at the frequency set in the frequency. <br>'Months' means that the Snapshot copy occurs monthly at the frequency set in the Frequency. <br>'Years' means that the Snapshot copy occurs yearly at the frequency set in the scheduleFrequency.
				  * Constraints: Allowable values are: `Runs`, `Hours`, `Days`, `Weeks`, `Months`, `Years`.
			* `target_id` - (Integer) Specifies the RPaaS target to copy the Snapshots.
			* `target_name` - (String) Specifies the RPaaS target name where Snapshots are copied.
			* `target_type` - (String) Specifies the RPaaS target type where Snapshots are copied.
			  * Constraints: Allowable values are: `Tape`, `Cloud`, `Nas`.
	* `retry_options` - (List) Retry Options of a Protection Policy when a Protection Group run fails.
	Nested schema for **retry_options**:
		* `retries` - (Integer) Specifies the number of times to retry capturing Snapshots before the Protection Group Run fails.
		  * Constraints: The minimum value is `0`.
		* `retry_interval_mins` - (Integer) Specifies the number of minutes before retrying a failed Protection Group.
		  * Constraints: The minimum value is `1`.
	* `template_id` - (String) Specifies the parent policy template id to which the policy is linked to. This field is set only when policy is created from template.
	* `version` - (Integer) Specifies the current policy verison. Policy version is incremented for optionally supporting new features and differentialting across releases.

