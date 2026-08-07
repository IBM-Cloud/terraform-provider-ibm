---
layout: "ibm"
page_title: "IBM : ibm_backup_recovery_protection_run_progress"
description: |-
  Get information about backup_recovery_protection_run_progress
subcategory: "IBM Backup Recovery API"
---

# ibm_backup_recovery_protection_run_progress

Provides a read-only data source to retrieve information about backup_recovery_protection_run_progress. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_backup_recovery_protection_run_progress" "backup_recovery_protection_run_progress" {
	run_id = "run_id"
	x_ibm_tenant_id = "tenantId"
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `end_time_usecs` - (Optional, Integer) Specifies the time before which the progress task ends in Unix epoch Timestamp(in microseconds).
* `exclude_object_details` - (Optional, Boolean) Specifies whether to return objects. By default all the task tree are returned.
* `include_event_logs` - (Optional, Boolean) Specifies whether to include event logs.
* `include_finished_tasks` - (Optional, Boolean) Specifies whether to return finished tasks. By default only active tasks are returned.
* `include_tenants` - (Optional, Boolean) If true, the response will include Protection Group Runs which were created by all tenants which the current user has permission to see. If false, then only Protection Groups created by the current user will be returned. If it's not specified, it is true by default.
* `max_log_level` - (Optional, Integer) Specifies the number of levels till which to fetch the event logs. This is applicable only when includeEventLogs is true.
* `max_tasks_num` - (Optional, Integer) Specifies the maximum number of tasks to return.
* `object_task_paths` - (Optional, List) Specifies the object level task path. This relates to the objectID. If provided this will take precedence over the objects, and will be used to fetch progress details directly without looking actuall task path of the object.
* `objects` - (Optional, List) Specifies the objects whose progress will be returned. This only applies to protection group runs and will be ignored for object runs. If the objects are specified, the run progress will not be returned and only the progress of the specified objects will be returned.
* `run_id` - (Required, Forces new resource, String) Specifies a unique run id of the Protection Run.
* `run_task_path` - (Optional, String) Specifies the task path of the run or object run. This is applicable only if progress of a protection group with one or more object is required.If provided this will be used to fetch progress details directly without looking actual task path of the object. Objects field is stil expected else it changes the response format.
* `start_time_usecs` - (Optional, Integer) Specifies the time after which the progress task starts in Unix epoch Timestamp(in microseconds).
* `tenant_ids` - (Optional, List) TenantIds contains ids of the tenants for which the run is to be returned.
* `x_ibm_tenant_id` - (Required, String) Id of the tenant accessing the cluster.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the backup_recovery_protection_run_progress.
* `archival_run` - (List) Progress for the archival run.
Nested schema for **archival_run**:
	* `archival_task_id` - (String) Specifies the archival task id. This is a protection group UID which only applies when archival type is 'Tape'.
	* `end_time_usecs` - (Integer) Specifies the end time of the progress task in Unix epoch Timestamp(in microseconds).
	* `events` - (List) Specifies the event log created for progress Task.
	Nested schema for **events**:
		* `message` - (String) Specifies the log message describing the current event.
		* `occured_at_usecs` - (Integer) Specifies the time of the event occurance in Unix epoch Timestamp(in microseconds).
	* `expected_remaining_time_usecs` - (Integer) Specifies the expected remaining time of the progress task in Unix epoch Timestamp(in microseconds).
	* `objects` - (List) Specifies progress for objects.
	Nested schema for **objects**:
		* `end_time_usecs` - (Integer) Specifies the end time of the progress task in Unix epoch Timestamp(in microseconds).
		* `environment` - (String) Specifies the environment of the object.
		  * Constraints: Allowable values are: `kPhysical`, `kSQL`.
		* `events` - (List) Specifies the event log created for progress Task.
		Nested schema for **events**:
			* `message` - (String) Specifies the log message describing the current event.
			* `occured_at_usecs` - (Integer) Specifies the time of the event occurance in Unix epoch Timestamp(in microseconds).
		* `expected_remaining_time_usecs` - (Integer) Specifies the expected remaining time of the progress task in Unix epoch Timestamp(in microseconds).
		* `failed_attempts` - (List) Specifies progress for failed attempts of this object.
		Nested schema for **failed_attempts**:
			* `end_time_usecs` - (Integer) Specifies the end time of the progress task in Unix epoch Timestamp(in microseconds).
			* `events` - (List) Specifies the event log created for progress Task.
			Nested schema for **events**:
				* `message` - (String) Specifies the log message describing the current event.
				* `occured_at_usecs` - (Integer) Specifies the time of the event occurance in Unix epoch Timestamp(in microseconds).
			* `expected_remaining_time_usecs` - (Integer) Specifies the expected remaining time of the progress task in Unix epoch Timestamp(in microseconds).
			* `percentage_completed` - (Float) Specifies the current completed percentage of the progress task.
			* `start_time_usecs` - (Integer) Specifies the start time of the progress task in Unix epoch Timestamp(in microseconds).
			* `stats` - (List) Specifies the stats within progress.
			Nested schema for **stats**:
				* `backup_file_count` - (Integer) Specifies the total number of file and directory entities that are backed up in this run. Only applicable to file based backups.
				* `file_walk_done` - (Boolean) Specifies whether the file system walk is done. Only applicable to file based backups.
				* `total_file_count` - (Integer) Specifies the total number of file and directory entities visited in this backup. Only applicable to file based backups.
			* `status` - (String) Specifies the current status of the progress task.
			  * Constraints: Allowable values are: `Active`, `Finished`, `FinishedWithError`, `Canceled`, `FinishedGarbageCollected`.
		* `id` - (Integer) Specifies object id.
		* `name` - (String) Specifies the name of the object.
		* `percentage_completed` - (Float) Specifies the current completed percentage of the progress task.
		* `source_id` - (Integer) Specifies registered source id to which object belongs.
		* `source_name` - (String) Specifies registered source name to which object belongs.
		* `start_time_usecs` - (Integer) Specifies the start time of the progress task in Unix epoch Timestamp(in microseconds).
		* `stats` - (List) Specifies the stats within progress.
		Nested schema for **stats**:
			* `backup_file_count` - (Integer) Specifies the total number of file and directory entities that are backed up in this run. Only applicable to file based backups.
			* `file_walk_done` - (Boolean) Specifies whether the file system walk is done. Only applicable to file based backups.
			* `total_file_count` - (Integer) Specifies the total number of file and directory entities visited in this backup. Only applicable to file based backups.
		* `status` - (String) Specifies the current status of the progress task.
		  * Constraints: Allowable values are: `Active`, `Finished`, `FinishedWithError`, `Canceled`, `FinishedGarbageCollected`.
	* `ownership_context` - (String) Specifies the ownership context for the target.
	  * Constraints: Allowable values are: `Local`, `FortKnox`.
	* `percentage_completed` - (Float) Specifies the current completed percentage of the progress task.
	* `start_time_usecs` - (Integer) Specifies the start time of the progress task in Unix epoch Timestamp(in microseconds).
	* `stats` - (List) Specifies the stats within progress.
	Nested schema for **stats**:
		* `backup_file_count` - (Integer) Specifies the total number of file and directory entities that are backed up in this run. Only applicable to file based backups.
		* `file_walk_done` - (Boolean) Specifies whether the file system walk is done. Only applicable to file based backups.
		* `total_file_count` - (Integer) Specifies the total number of file and directory entities visited in this backup. Only applicable to file based backups.
	* `status` - (String) Specifies the current status of the progress task.
	  * Constraints: Allowable values are: `Active`, `Finished`, `FinishedWithError`, `Canceled`, `FinishedGarbageCollected`.
	* `target_id` - (Integer) Specifies the archival target ID.
	* `target_name` - (String) Specifies the archival target name.
	* `target_type` - (String) Specifies the archival target type.
	  * Constraints: Allowable values are: `Tape`, `Cloud`, `Nas`.
	* `tier_settings` - (List) Specifies the tier info for archival.
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
		* `current_tier_type` - (String) Specifies the type of the current tier where the snapshot resides. This will be specified if the run is a CAD run.
		  * Constraints: Allowable values are: `kAmazonS3Standard`, `kAmazonS3StandardIA`, `kAmazonS3OneZoneIA`, `kAmazonS3IntelligentTiering`, `kAmazonS3Glacier`, `kAmazonS3GlacierDeepArchive`, `kAzureTierHot`, `kAzureTierCool`, `kAzureTierArchive`, `kGoogleStandard`, `kGoogleRegional`, `kGoogleMultiRegional`, `kGoogleNearline`, `kGoogleColdline`, `kOracleTierStandard`, `kOracleTierArchive`.
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
	* `usage_type` - (String) Specifies the usage type for the target.
	  * Constraints: Allowable values are: `Archival`, `Tiering`, `Rpaas`.
* `local_run` - (List) Specifies the progress of a local backup run.
Nested schema for **local_run**:
	* `end_time_usecs` - (Integer) Specifies the end time of the progress task in Unix epoch Timestamp(in microseconds).
	* `events` - (List) Specifies the event log created for progress Task.
	Nested schema for **events**:
		* `message` - (String) Specifies the log message describing the current event.
		* `occured_at_usecs` - (Integer) Specifies the time of the event occurance in Unix epoch Timestamp(in microseconds).
	* `expected_remaining_time_usecs` - (Integer) Specifies the expected remaining time of the progress task in Unix epoch Timestamp(in microseconds).
	* `objects` - (List) Specifies progress for objects.
	Nested schema for **objects**:
		* `end_time_usecs` - (Integer) Specifies the end time of the progress task in Unix epoch Timestamp(in microseconds).
		* `environment` - (String) Specifies the environment of the object.
		  * Constraints: Allowable values are: `kPhysical`, `kSQL`.
		* `events` - (List) Specifies the event log created for progress Task.
		Nested schema for **events**:
			* `message` - (String) Specifies the log message describing the current event.
			* `occured_at_usecs` - (Integer) Specifies the time of the event occurance in Unix epoch Timestamp(in microseconds).
		* `expected_remaining_time_usecs` - (Integer) Specifies the expected remaining time of the progress task in Unix epoch Timestamp(in microseconds).
		* `failed_attempts` - (List) Specifies progress for failed attempts of this object.
		Nested schema for **failed_attempts**:
			* `end_time_usecs` - (Integer) Specifies the end time of the progress task in Unix epoch Timestamp(in microseconds).
			* `events` - (List) Specifies the event log created for progress Task.
			Nested schema for **events**:
				* `message` - (String) Specifies the log message describing the current event.
				* `occured_at_usecs` - (Integer) Specifies the time of the event occurance in Unix epoch Timestamp(in microseconds).
			* `expected_remaining_time_usecs` - (Integer) Specifies the expected remaining time of the progress task in Unix epoch Timestamp(in microseconds).
			* `percentage_completed` - (Float) Specifies the current completed percentage of the progress task.
			* `start_time_usecs` - (Integer) Specifies the start time of the progress task in Unix epoch Timestamp(in microseconds).
			* `stats` - (List) Specifies the stats within progress.
			Nested schema for **stats**:
				* `backup_file_count` - (Integer) Specifies the total number of file and directory entities that are backed up in this run. Only applicable to file based backups.
				* `file_walk_done` - (Boolean) Specifies whether the file system walk is done. Only applicable to file based backups.
				* `total_file_count` - (Integer) Specifies the total number of file and directory entities visited in this backup. Only applicable to file based backups.
			* `status` - (String) Specifies the current status of the progress task.
			  * Constraints: Allowable values are: `Active`, `Finished`, `FinishedWithError`, `Canceled`, `FinishedGarbageCollected`.
		* `id` - (Integer) Specifies object id.
		* `name` - (String) Specifies the name of the object.
		* `percentage_completed` - (Float) Specifies the current completed percentage of the progress task.
		* `source_id` - (Integer) Specifies registered source id to which object belongs.
		* `source_name` - (String) Specifies registered source name to which object belongs.
		* `start_time_usecs` - (Integer) Specifies the start time of the progress task in Unix epoch Timestamp(in microseconds).
		* `stats` - (List) Specifies the stats within progress.
		Nested schema for **stats**:
			* `backup_file_count` - (Integer) Specifies the total number of file and directory entities that are backed up in this run. Only applicable to file based backups.
			* `file_walk_done` - (Boolean) Specifies whether the file system walk is done. Only applicable to file based backups.
			* `total_file_count` - (Integer) Specifies the total number of file and directory entities visited in this backup. Only applicable to file based backups.
		* `status` - (String) Specifies the current status of the progress task.
		  * Constraints: Allowable values are: `Active`, `Finished`, `FinishedWithError`, `Canceled`, `FinishedGarbageCollected`.
	* `percentage_completed` - (Float) Specifies the current completed percentage of the progress task.
	* `start_time_usecs` - (Integer) Specifies the start time of the progress task in Unix epoch Timestamp(in microseconds).
	* `stats` - (List) Specifies the stats within progress.
	Nested schema for **stats**:
		* `backup_file_count` - (Integer) Specifies the total number of file and directory entities that are backed up in this run. Only applicable to file based backups.
		* `file_walk_done` - (Boolean) Specifies whether the file system walk is done. Only applicable to file based backups.
		* `total_file_count` - (Integer) Specifies the total number of file and directory entities visited in this backup. Only applicable to file based backups.
	* `status` - (String) Specifies the current status of the progress task.
	  * Constraints: Allowable values are: `Active`, `Finished`, `FinishedWithError`, `Canceled`, `FinishedGarbageCollected`.
* `replication_run` - (List) Progress for the replication run.
Nested schema for **replication_run**:
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
	* `cluster_id` - (Integer) Specifies the id of the cluster.
	* `cluster_incarnation_id` - (Integer) Specifies the incarnation id of the cluster.
	* `cluster_name` - (String) Specifies the name of the cluster.
	* `end_time_usecs` - (Integer) Specifies the end time of the progress task in Unix epoch Timestamp(in microseconds).
	* `events` - (List) Specifies the event log created for progress Task.
	Nested schema for **events**:
		* `message` - (String) Specifies the log message describing the current event.
		* `occured_at_usecs` - (Integer) Specifies the time of the event occurance in Unix epoch Timestamp(in microseconds).
	* `expected_remaining_time_usecs` - (Integer) Specifies the expected remaining time of the progress task in Unix epoch Timestamp(in microseconds).
	* `objects` - (List) Specifies progress for objects.
	Nested schema for **objects**:
		* `end_time_usecs` - (Integer) Specifies the end time of the progress task in Unix epoch Timestamp(in microseconds).
		* `environment` - (String) Specifies the environment of the object.
		  * Constraints: Allowable values are: `kPhysical`, `kSQL`.
		* `events` - (List) Specifies the event log created for progress Task.
		Nested schema for **events**:
			* `message` - (String) Specifies the log message describing the current event.
			* `occured_at_usecs` - (Integer) Specifies the time of the event occurance in Unix epoch Timestamp(in microseconds).
		* `expected_remaining_time_usecs` - (Integer) Specifies the expected remaining time of the progress task in Unix epoch Timestamp(in microseconds).
		* `failed_attempts` - (List) Specifies progress for failed attempts of this object.
		Nested schema for **failed_attempts**:
			* `end_time_usecs` - (Integer) Specifies the end time of the progress task in Unix epoch Timestamp(in microseconds).
			* `events` - (List) Specifies the event log created for progress Task.
			Nested schema for **events**:
				* `message` - (String) Specifies the log message describing the current event.
				* `occured_at_usecs` - (Integer) Specifies the time of the event occurance in Unix epoch Timestamp(in microseconds).
			* `expected_remaining_time_usecs` - (Integer) Specifies the expected remaining time of the progress task in Unix epoch Timestamp(in microseconds).
			* `percentage_completed` - (Float) Specifies the current completed percentage of the progress task.
			* `start_time_usecs` - (Integer) Specifies the start time of the progress task in Unix epoch Timestamp(in microseconds).
			* `stats` - (List) Specifies the stats within progress.
			Nested schema for **stats**:
				* `backup_file_count` - (Integer) Specifies the total number of file and directory entities that are backed up in this run. Only applicable to file based backups.
				* `file_walk_done` - (Boolean) Specifies whether the file system walk is done. Only applicable to file based backups.
				* `total_file_count` - (Integer) Specifies the total number of file and directory entities visited in this backup. Only applicable to file based backups.
			* `status` - (String) Specifies the current status of the progress task.
			  * Constraints: Allowable values are: `Active`, `Finished`, `FinishedWithError`, `Canceled`, `FinishedGarbageCollected`.
		* `id` - (Integer) Specifies object id.
		* `name` - (String) Specifies the name of the object.
		* `percentage_completed` - (Float) Specifies the current completed percentage of the progress task.
		* `source_id` - (Integer) Specifies registered source id to which object belongs.
		* `source_name` - (String) Specifies registered source name to which object belongs.
		* `start_time_usecs` - (Integer) Specifies the start time of the progress task in Unix epoch Timestamp(in microseconds).
		* `stats` - (List) Specifies the stats within progress.
		Nested schema for **stats**:
			* `backup_file_count` - (Integer) Specifies the total number of file and directory entities that are backed up in this run. Only applicable to file based backups.
			* `file_walk_done` - (Boolean) Specifies whether the file system walk is done. Only applicable to file based backups.
			* `total_file_count` - (Integer) Specifies the total number of file and directory entities visited in this backup. Only applicable to file based backups.
		* `status` - (String) Specifies the current status of the progress task.
		  * Constraints: Allowable values are: `Active`, `Finished`, `FinishedWithError`, `Canceled`, `FinishedGarbageCollected`.
	* `percentage_completed` - (Float) Specifies the current completed percentage of the progress task.
	* `start_time_usecs` - (Integer) Specifies the start time of the progress task in Unix epoch Timestamp(in microseconds).
	* `stats` - (List) Specifies the stats within progress.
	Nested schema for **stats**:
		* `backup_file_count` - (Integer) Specifies the total number of file and directory entities that are backed up in this run. Only applicable to file based backups.
		* `file_walk_done` - (Boolean) Specifies whether the file system walk is done. Only applicable to file based backups.
		* `total_file_count` - (Integer) Specifies the total number of file and directory entities visited in this backup. Only applicable to file based backups.
	* `status` - (String) Specifies the current status of the progress task.
	  * Constraints: Allowable values are: `Active`, `Finished`, `FinishedWithError`, `Canceled`, `FinishedGarbageCollected`.

