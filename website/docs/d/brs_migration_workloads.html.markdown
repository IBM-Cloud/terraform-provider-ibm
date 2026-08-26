---
layout: "ibm"
page_title: "IBM : ibm_brs_migration_workloads"
description: |-
  Get information about brs_migration_workloads
subcategory: "IBM Cloud Backup and Recovery Migration API"
---

# ibm_brs_migration_workloads

Provides a read-only data source to retrieve information about brs_migration_workloads. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_brs_migration_workloads" "brs_migration_workloads" {
	migration_id = ibm_brs_migration_workload.brs_migration_workload_instance.migration_id
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `migration_id` - (Required, Forces new resource, String) The migration project ID (mgr-{uuid4} format).
  * Constraints: Length must be `40` characters. The value must match regular expression `/^mgr-[0-9a-f-]{36}$/`.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the brs_migration_workloads.
* `workloads` - (List) The list of workloads on this page.
  * Constraints: The maximum length is `100` items. The minimum length is `0` items.
Nested schema for **workloads**:
	* `completed_at` - (String) Timestamp when `POST /migrations/{migration_id}/workloads/{workload_id}/complete` finished.
	* `created_at` - (String) Timestamp when this workload was created.
	* `id` - (String) Migration service workload ID (wl-{uuid4} format).
	  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[0-9a-zA-Z-]+$/`.
	* `name` - (String) Human-readable name for this workload.
	  * Constraints: The maximum length is `1024` characters. The minimum length is `1` character. The value must match regular expression `/^.+$/`.
	* `payloads` - (List) List of source-to-destination payload mappings.
	  * Constraints: The maximum length is `1000` items. The minimum length is `0` items.
	Nested schema for **payloads**:
		* `data_specs` - (List) One or more data specifications for this payload.
		  * Constraints: The maximum length is `1000` items. The minimum length is `0` items.
		Nested schema for **data_specs**:
			* `data_format` - (String) How this data unit is structured. When `source.volume_id` is a registered `vol-*` ID, the server derives this from the volume registry — pass `raw` as a default if unsure.
			  * Constraints: Allowable values are: `raw`, `file_system`, `database`.
			* `destination` - (List) Destination data location and format for this mapping.
			Nested schema for **destination**:
				* `path` - (String) OS mount path, database data directory path, or block device path.
				  * Constraints: The maximum length is `1024` characters. The minimum length is `1` character. The value must match regular expression `/^.+$/`.
				* `type` - (String) Filesystem type, database engine, or raw block device type marker.
				  * Constraints: The maximum length is `1024` characters. The minimum length is `1` character. The value must match regular expression `/^.+$/`.
				* `volume_id` - (String) Migration service vol-* ID from POST /migrations/{migration_id}/volumes.
				  * Constraints: The maximum length is `1024` characters. The minimum length is `1` character. The value must match regular expression `/^.+$/`.
			* `source` - (List) Source data location and format for this mapping.
			Nested schema for **source**:
				* `path` - (String) OS mount path, database data directory path, or block device path.
				  * Constraints: The maximum length is `1024` characters. The minimum length is `1` character. The value must match regular expression `/^.+$/`.
				* `type` - (String) Filesystem type, database engine, or raw block device type marker.
				  * Constraints: The maximum length is `1024` characters. The minimum length is `1` character. The value must match regular expression `/^.+$/`.
				* `volume_id` - (String) Migration service vol-* ID from POST /migrations/{migration_id}/volumes.
				  * Constraints: The maximum length is `1024` characters. The minimum length is `1` character. The value must match regular expression `/^.+$/`.
		* `destination_host_id` - (String) Migration service ID of the destination host.
		  * Constraints: The maximum length is `1024` characters. The minimum length is `1` character. The value must match regular expression `/^.+$/`.
		* `id` - (String) Migration service payload ID (pl-{uuid4} format).
		  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[0-9a-zA-Z-]+$/`.
		* `source_host_id` - (String) Migration service ID of the source host.
		  * Constraints: The maximum length is `1024` characters. The minimum length is `1` character. The value must match regular expression `/^.+$/`.
	* `schedule` - (List) Populated once setup completes. Null while settingUp or in created/failed state.
	Nested schema for **schedule**:
		* `backup_policy` - (List) Specifies the backup schedule and retentions of a protection policy.
		Nested schema for **backup_policy**:
			* `log` - (List) Specifies log backup schedule settings for database workloads.
			Nested schema for **log**:
				* `schedule` - (List) Defines how frequently log backup runs are started.
				Nested schema for **schedule**:
					* `hour_schedule` - (List) A simple multiplier schedule — repeat every N units.
					Nested schema for **hour_schedule**:
						* `frequency` - (Integer) How many units between each run.
						  * Constraints: The maximum value is `999999`. The minimum value is `1`.
					* `minute_schedule` - (List) A simple multiplier schedule — repeat every N units.
					Nested schema for **minute_schedule**:
						* `frequency` - (Integer) How many units between each run.
						  * Constraints: The maximum value is `999999`. The minimum value is `1`.
					* `unit` - (String) Frequency unit for the log backup schedule.
					  * Constraints: Allowable values are: `minutes`, `hours`.
			* `regular` - (List) Specifies the incremental and full backup schedule settings with retention. When `incremental` or `full` is provided, `retention` must also be provided — the BRS API returns a 400 if a schedule is present without retention. When neither schedule is supplied, a default daily incremental with 30-day retention is applied automatically by the service.
			Nested schema for **regular**:
				* `full` - (List) Specifies full backup schedule settings.
				Nested schema for **full**:
					* `schedule` - (List) Defines when full backup runs are started. The sub-schedule field matching `unit` must be provided: `days` → `day_schedule`, `weeks` → `week_schedule` (with at least one day), `months` → `month_schedule`, `years` → `year_schedule`. The BRS API returns a 400 if the matching sub-schedule is absent.
					Nested schema for **schedule**:
						* `day_schedule` - (List) A simple multiplier schedule — repeat every N units.
						Nested schema for **day_schedule**:
							* `frequency` - (Integer) How many units between each run.
							  * Constraints: The maximum value is `999999`. The minimum value is `1`.
						* `month_schedule` - (List) Runs on a specific week and day within a month, or on a fixed day-of-month. Use `week_of_month` + `day_of_week` for a week-relative run (e.g. "first Sunday"), or `day_of_month` alone for a fixed calendar date. These two forms are mutually exclusive; `day_of_month` takes precedence when both are supplied.
						Nested schema for **month_schedule**:
							* `day_of_month` - (Integer) Exact day of the month (1-31) for the run. When set, `week_of_month` and `day_of_week` are ignored.
							  * Constraints: The maximum value is `31`. The minimum value is `1`.
							* `day_of_week` - (List) Days of the week when runs are started (used with `week_of_month`). Must contain at least one entry.
							  * Constraints: Allowable list items are: `sunday`, `monday`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`. The maximum length is `7` items. The minimum length is `1` item.
							* `week_of_month` - (String) Which week of the month the run falls on. Use with `day_of_week`.
							  * Constraints: Allowable values are: `first`, `second`, `third`, `fourth`, `last`.
						* `unit` - (String) Frequency unit for the full backup schedule.
						  * Constraints: Allowable values are: `days`, `weeks`, `months`, `years`.
						* `week_schedule` - (List) Runs on specific days of the week. At least one day must be specified.
						Nested schema for **week_schedule**:
							* `day_of_week` - (List) Days of the week on which runs are started. At least one day is required.
							  * Constraints: Allowable list items are: `sunday`, `monday`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`. The maximum length is `7` items. The minimum length is `1` item.
						* `year_schedule` - (List) Runs on the first or last day of the year.
						Nested schema for **year_schedule**:
							* `day_of_year` - (String) Day of the year when the run starts.
							  * Constraints: Allowable values are: `first`, `last`.
				* `incremental` - (List) Specifies incremental backup schedule settings.
				Nested schema for **incremental**:
					* `schedule` - (List) Defines how frequently incremental backup runs are started. The sub-schedule field matching `unit` must be provided: `minutes` → `minute_schedule`, `hours` → `hour_schedule`, `days` → `day_schedule`, `weeks` → `week_schedule` (with at least one day), `months` → `month_schedule`, `years` → `year_schedule`. The BRS API returns a 400 if the matching sub-schedule is absent.
					Nested schema for **schedule**:
						* `day_schedule` - (List) A simple multiplier schedule — repeat every N units.
						Nested schema for **day_schedule**:
							* `frequency` - (Integer) How many units between each run.
							  * Constraints: The maximum value is `999999`. The minimum value is `1`.
						* `hour_schedule` - (List) A simple multiplier schedule — repeat every N units.
						Nested schema for **hour_schedule**:
							* `frequency` - (Integer) How many units between each run.
							  * Constraints: The maximum value is `999999`. The minimum value is `1`.
						* `minute_schedule` - (List) A simple multiplier schedule — repeat every N units.
						Nested schema for **minute_schedule**:
							* `frequency` - (Integer) How many units between each run.
							  * Constraints: The maximum value is `999999`. The minimum value is `1`.
						* `month_schedule` - (List) Runs on a specific week and day within a month, or on a fixed day-of-month. Use `week_of_month` + `day_of_week` for a week-relative run (e.g. "first Sunday"), or `day_of_month` alone for a fixed calendar date. These two forms are mutually exclusive; `day_of_month` takes precedence when both are supplied.
						Nested schema for **month_schedule**:
							* `day_of_month` - (Integer) Exact day of the month (1-31) for the run. When set, `week_of_month` and `day_of_week` are ignored.
							  * Constraints: The maximum value is `31`. The minimum value is `1`.
							* `day_of_week` - (List) Days of the week when runs are started (used with `week_of_month`). Must contain at least one entry.
							  * Constraints: Allowable list items are: `sunday`, `monday`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`. The maximum length is `7` items. The minimum length is `1` item.
							* `week_of_month` - (String) Which week of the month the run falls on. Use with `day_of_week`.
							  * Constraints: Allowable values are: `first`, `second`, `third`, `fourth`, `last`.
						* `unit` - (String) Frequency unit for the incremental schedule.
						  * Constraints: Allowable values are: `minutes`, `hours`, `days`, `weeks`, `months`, `years`.
						* `week_schedule` - (List) Runs on specific days of the week. At least one day must be specified.
						Nested schema for **week_schedule**:
							* `day_of_week` - (List) Days of the week on which runs are started. At least one day is required.
							  * Constraints: Allowable list items are: `sunday`, `monday`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`. The maximum length is `7` items. The minimum length is `1` item.
						* `year_schedule` - (List) Runs on the first or last day of the year.
						Nested schema for **year_schedule**:
							* `day_of_year` - (String) Day of the year when the run starts.
							  * Constraints: Allowable values are: `first`, `last`.
				* `retention` - (List) Specifies how long backup snapshots are retained.
				Nested schema for **retention**:
					* `duration` - (Integer) Number of `unit`s to retain the snapshot.
					  * Constraints: The maximum value is `99999`. The minimum value is `1`.
					* `unit` - (String) Retention time unit.
					  * Constraints: Allowable values are: `days`, `weeks`, `months`, `years`.
		* `blackout_window` - (List) List of blackout windows during which new runs will not be started.
		  * Constraints: The maximum length is `1000` items. The minimum length is `0` items.
		Nested schema for **blackout_window**:
			* `day` - (String) Day of the week the blackout applies to.
			  * Constraints: Allowable values are: `sunday`, `monday`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`.
			* `end_time` - (List) Specifies a time of day used in scheduling and blackout windows.
			Nested schema for **end_time**:
				* `hour` - (Integer) Hour of the day (0-23).
				  * Constraints: The maximum value is `23`. The minimum value is `0`.
				* `minute` - (Integer) Minute of the hour (0-59).
				  * Constraints: The maximum value is `59`. The minimum value is `0`.
				* `timezone` - (String) IANA time zone name. Defaults to `America/Los_Angeles` if omitted.
				  * Constraints: The default value is `America/Los_Angeles`. The maximum length is `1024` characters. The minimum length is `1` character. The value must match regular expression `/^.+$/`.
			* `start_time` - (List) Specifies a time of day used in scheduling and blackout windows.
			Nested schema for **start_time**:
				* `hour` - (Integer) Hour of the day (0-23).
				  * Constraints: The maximum value is `23`. The minimum value is `0`.
				* `minute` - (Integer) Minute of the hour (0-59).
				  * Constraints: The maximum value is `59`. The minimum value is `0`.
				* `timezone` - (String) IANA time zone name. Defaults to `America/Los_Angeles` if omitted.
				  * Constraints: The default value is `America/Los_Angeles`. The maximum length is `1024` characters. The minimum length is `1` character. The value must match regular expression `/^.+$/`.
		* `description` - (String) Optional description of the schedule.
		  * Constraints: The maximum length is `1024` characters. The minimum length is `1` character. The value must match regular expression `/^.+$/`.
		* `extended_retention` - (List) Additional retention policies applied on top of regular retention.
		  * Constraints: The maximum length is `1000` items. The minimum length is `0` items.
		Nested schema for **extended_retention**:
			* `retention` - (List) Specifies how long backup snapshots are retained.
			Nested schema for **retention**:
				* `duration` - (Integer) Number of `unit`s to retain the snapshot.
				  * Constraints: The maximum value is `99999`. The minimum value is `1`.
				* `unit` - (String) Retention time unit.
				  * Constraints: Allowable values are: `days`, `weeks`, `months`, `years`.
			* `run_type` - (String) The backup run type this extended retention applies to.
			  * Constraints: Allowable values are: `regular`, `full`, `log`, `system`, `storage_array_snapshot`.
			* `schedule` - (List) Frequency at which the extended retention rule applies.
			Nested schema for **schedule**:
				* `frequency` - (Integer) Multiplier applied to the unit.
				  * Constraints: The maximum value is `999999`. The minimum value is `1`.
				* `unit` - (String) Schedule unit for extended retention.
				  * Constraints: Allowable values are: `runs`, `hours`, `days`, `weeks`, `months`, `years`.
		* `name` - (String) Human-readable name for the schedule / protection policy.
		  * Constraints: The maximum length is `1024` characters. The minimum length is `1` character. The value must match regular expression `/^.+$/`.
		* `retry_options` - (List) Controls how failed runs are retried.
		Nested schema for **retry_options**:
			* `retries` - (Integer) Number of retry attempts before the run is marked failed.
			  * Constraints: The maximum value is `10`. The minimum value is `0`.
			* `retry_interval_mins` - (Integer) Minutes to wait between retry attempts.
			  * Constraints: The maximum value is `1440`. The minimum value is `1`.
	* `scheduling_error` - (String) Non-empty when `state` is `failed` due to async setup failure.
	  * Constraints: The maximum length is `1024` characters. The minimum length is `1` character. The value must match regular expression `/^.+$/`.
	* `state` - (String) Current lifecycle state of the workload.
	  * Constraints: Allowable values are: `created`, `setting_up`, `scheduled`, `running`, `canceling`, `canceled`, `succeeded`, `failed`, `completing`, `completed`.
	* `updated_at` - (String) Timestamp of the last update to this workload.
	* `volume_ownership_map` - (Map) Server-computed map of `vol-*` volume_id → owning payload_id. Only populated when the same `source.volume_id` appears in two or more payloads. Null in all other cases.

