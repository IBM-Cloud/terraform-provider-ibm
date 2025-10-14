---
layout: "ibm"
page_title: "IBM : ibm_backup_recovery_protection_group_runs"
description: |-
  Get information about backup_recovery_protection_group_runs
subcategory: "IBM Backup Recovery"
---

# ibm_backup_recovery_protection_group_runs

Provides a read-only data source to retrieve information about backup_recovery_protection_group_runs. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_backup_recovery_protection_group_runs" "backup_recovery_protection_group_runs" {
	protection_group_id = ibm_backup_recovery_protection_group.backup_recovery_protection_group_instance.backup_recovery_protection_group_id
	x_ibm_tenant_id = "x_ibm_tenant_id"
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `archival_run_status` - (Optional, List) Specifies a list of archival status, runs matching the status will be returned.<br> 'Running' indicates that the run is still running.<br> 'Canceled' indicates that the run has been canceled.<br> 'Canceling' indicates that the run is in the process of being canceled.<br> 'Failed' indicates that the run has failed.<br> 'Missed' indicates that the run was unable to take place at the scheduled time because the previous run was still happening.<br> 'Succeeded' indicates that the run has finished successfully.<br> 'SucceededWithWarning' indicates that the run finished successfully, but there were some warning messages.<br> 'Paused' indicates that the ongoing run has been paused.<br> 'Skipped' indicates that the run was skipped.
* `backup_recovery_endpoint` - (Optional, String) Backup Recovery Endpoint URL. If provided here, it overrides values configured via environment variable (IBMCLOUD_BACKUP_RECOVERY_ENDPOINT) or endpoints.json.   
  * Constraints: Allowable list items are: `Accepted`, `Running`, `Canceled`, `Canceling`, `Failed`, `Missed`, `Succeeded`, `SucceededWithWarning`, `OnHold`, `Finalizing`, `Skipped`, `Paused`.
* `protection_group_id` - (Required, Forces new resource, String) Specifies a unique id of the Protection Group.
  * Constraints: The value must match regular expression `/^\\d+:\\d+:\\d+$/`.
* `cloud_spin_run_status` - (Optional, List) Specifies a list of cloud spin status, runs matching the status will be returned.<br> 'Running' indicates that the run is still running.<br> 'Canceled' indicates that the run has been canceled.<br> 'Canceling' indicates that the run is in the process of being canceled.<br> 'Failed' indicates that the run has failed.<br> 'Missed' indicates that the run was unable to take place at the scheduled time because the previous run was still happening.<br> 'Succeeded' indicates that the run has finished successfully.<br> 'SucceededWithWarning' indicates that the run finished successfully, but there were some warning messages.<br> 'Paused' indicates that the ongoing run has been paused.<br> 'Skipped' indicates that the run was skipped.
  * Constraints: Allowable list items are: `Accepted`, `Running`, `Canceled`, `Canceling`, `Failed`, `Missed`, `Succeeded`, `SucceededWithWarning`, `OnHold`, `Finalizing`, `Skipped`, `Paused`.
* `end_time_usecs` - (Optional, Integer) End time for time range filter. Specify the end time as a Unix epoch Timestamp (in microseconds), only runs executing before this time will be returned. By default it is current time.
* `exclude_non_restorable_runs` - (Optional, Boolean) Specifies whether to exclude non restorable runs. Run is treated restorable only if there is atleast one object snapshot (which may be either a local or an archival snapshot) which is not deleted or expired. Default value is false.
  * Constraints: The default value is `false`.
* `filter_by_copy_task_end_time` - (Optional, Boolean) If true, then the details of the runs for which any copyTask completed in the given timerange will be returned. Only one of filterByEndTime and filterByCopyTaskEndTime can be set.
* `filter_by_end_time` - (Optional, Boolean) If true, the runs with backup end time within the specified time range will be returned. Otherwise, the runs with start time in the time range are returned.
* `include_object_details` - (Optional, Boolean) Specifies if the result includes the object details for each protection run. If set to true, details of the protected object will be returned. If set to false or not specified, details will not be returned.
* `local_backup_run_status` - (Optional, List) Specifies a list of local backup status, runs matching the status will be returned.<br> 'Running' indicates that the run is still running.<br> 'Canceled' indicates that the run has been canceled.<br> 'Canceling' indicates that the run is in the process of being canceled.<br> 'Failed' indicates that the run has failed.<br> 'Missed' indicates that the run was unable to take place at the scheduled time because the previous run was still happening.<br> 'Succeeded' indicates that the run has finished successfully.<br> 'SucceededWithWarning' indicates that the run finished successfully, but there were some warning messages.<br> 'Paused' indicates that the ongoing run has been paused.<br> 'Skipped' indicates that the run was skipped.
  * Constraints: Allowable list items are: `Accepted`, `Running`, `Canceled`, `Canceling`, `Failed`, `Missed`, `Succeeded`, `SucceededWithWarning`, `OnHold`, `Finalizing`, `Skipped`, `Paused`.
* `num_runs` - (Optional, Integer) Specifies the max number of runs. If not specified, at most 100 runs will be returned.
* `only_return_successful_copy_run` - (Optional, Boolean) only successful copyruns are returned.
* `replication_run_status` - (Optional, List) Specifies a list of replication status, runs matching the status will be returned.<br> 'Running' indicates that the run is still running.<br> 'Canceled' indicates that the run has been canceled.<br> 'Canceling' indicates that the run is in the process of being canceled.<br> 'Failed' indicates that the run has failed.<br> 'Missed' indicates that the run was unable to take place at the scheduled time because the previous run was still happening.<br> 'Succeeded' indicates that the run has finished successfully.<br> 'SucceededWithWarning' indicates that the run finished successfully, but there were some warning messages.<br> 'Paused' indicates that the ongoing run has been paused.<br> 'Skipped' indicates that the run was skipped.
  * Constraints: Allowable list items are: `Accepted`, `Running`, `Canceled`, `Canceling`, `Failed`, `Missed`, `Succeeded`, `SucceededWithWarning`, `OnHold`, `Finalizing`, `Skipped`, `Paused`.
* `request_initiator_type` - (Optional, String) Specifies the type of request from UI, which is used for services like magneto to determine the priority of requests.
  * Constraints: Allowable values are: `UIUser`, `UIAuto`, `Helios`.
* `run_id` - (Optional, String) Specifies the protection run id.
  * Constraints: The value must match regular expression `/^\\d+:\\d+$/`.
* `run_tags` - (Optional, List) Specifies a list of tags for protection runs. If this is specified, only the runs which match these tags will be returned.
* `run_types` - (Optional, List) Filter by run type. Only protection run matching the specified types will be returned.
  * Constraints: Allowable list items are: `kAll`, `kHydrateCDP`, `kSystem`, `kStorageArraySnapshot`, `kIncremental`, `kFull`, `kLog`.
* `snapshot_target_types` - (Optional, List) Specifies the snapshot's target type which should be filtered.
  * Constraints: Allowable list items are: `Local`, `Archival`, `RpaasArchival`, `StorageArraySnapshot`, `Remote`.
* `start_time_usecs` - (Optional, Integer) Start time for time range filter. Specify the start time as a Unix epoch Timestamp (in microseconds), only runs executing after this time will be returned. By default it is endTimeUsecs minus an hour.
* `use_cached_data` - (Optional, Boolean) Specifies whether we can serve the GET request from the read replica cache. There is a lag of 15 seconds between the read replica and primary data source.
* `x_ibm_tenant_id` - (Required, String) Specifies the key to be used to encrypt the source credential. If includeSourceCredentials is set to true this key must be specified.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the backup_recovery_protection_group_runs.
* `runs` - (List) Specifies the list of Protection Group runs.
Nested schema for **runs**:
	* `archival_info` - (List) Specifies summary information about archival run.
	Nested schema for **archival_info**:
		* `archival_target_results` - (List) Archival results for each archival target.
		Nested schema for **archival_target_results**:
			* `archival_task_id` - (String) Specifies the archival task id. This is a protection group UID which only applies when archival type is 'Tape'.
			* `cancelled_app_objects_count` - (Integer) Specifies the count of app objects for which backup was cancelled.
			* `cancelled_objects_count` - (Integer) Specifies the count of objects for which backup was cancelled.
			* `data_lock_constraints` - (List) Specifies the dataLock constraints for local or target snapshot.
			Nested schema for **data_lock_constraints**:
				* `expiry_time_usecs` - (Integer) Specifies the expiry time of attempt in Unix epoch Timestamp (in microseconds).
				* `mode` - (String) Specifies the type of WORM retention type. <br>'Compliance' implies WORM retention is set for compliance reason. <br>'Administrative' implies WORM retention is set for administrative purposes.
				  * Constraints: Allowable values are: `Compliance`, `Administrative`.
			* `end_time_usecs` - (Integer) Specifies the end time of replication run in Unix epoch Timestamp(in microseconds) for an archival target.
			* `expiry_time_usecs` - (Integer) Specifies the expiry time of attempt in Unix epoch Timestamp (in microseconds).
			* `failed_app_objects_count` - (Integer) Specifies the count of app objects for which backup failed.
			* `failed_objects_count` - (Integer) Specifies the count of objects for which backup failed.
			* `indexing_task_id` - (String) Progress monitor task for indexing.
			* `is_cad_archive` - (Boolean) Whether this is CAD archive or not.
			* `is_forever_incremental` - (Boolean) Whether this is forever incremental or not.
			* `is_incremental` - (Boolean) Whether this is an incremental archive. If set to true, this is an incremental archive, otherwise this is a full archive.
			* `is_manually_deleted` - (Boolean) Specifies whether the snapshot is deleted manually.
			* `is_sla_violated` - (Boolean) Indicated if SLA has been violated for this run.
			* `message` - (String) Message about the archival run.
			* `on_legal_hold` - (Boolean) Specifies the legal hold status for a archival target.
			* `ownership_context` - (String) Specifies the ownership context for the target.
			  * Constraints: Allowable values are: `Local`, `FortKnox`.
			* `progress_task_id` - (String) Progress monitor task id for archival.
			* `queued_time_usecs` - (Integer) Specifies the time when the archival is queued for schedule in Unix epoch Timestamp(in microseconds) for a target.
			* `run_type` - (String) Type of Protection Group run. 'kRegular' indicates an incremental (CBT) backup. Incremental backups utilizing CBT (if supported) are captured of the target protection objects. The first run of a kRegular schedule captures all the blocks. 'kFull' indicates a full (no CBT) backup. A complete backup (all blocks) of the target protection objects are always captured and Change Block Tracking (CBT) is not utilized. 'kLog' indicates a Database Log backup. Capture the database transaction logs to allow rolling back to a specific point in time. 'kSystem' indicates system volume backup. It produces an image for bare metal recovery.
			  * Constraints: Allowable values are: `kRegular`, `kFull`, `kLog`, `kSystem`, `kHydrateCDP`, `kStorageArraySnapshot`.
			* `snapshot_id` - (String) Snapshot id for a successful snapshot. This field will not be set if the archival Run fails to take the snapshot.
			* `start_time_usecs` - (Integer) Specifies the start time of replication run in Unix epoch Timestamp(in microseconds) for an archival target.
			* `stats` - (List) Specifies statistics about archival data.
			Nested schema for **stats**:
				* `avg_logical_transfer_rate_bps` - (Integer) Specifies the average rate of transfer in bytes per second.
				* `backup_file_count` - (Integer) Specifies the total number of file and directory entities that are backed up in this run. Only applicable to file based backups.
				* `bytes_read` - (Integer) Specifies total logical bytes read for creating the snapshot.
				* `file_walk_done` - (Boolean) Specifies whether the file system walk is done. Only applicable to file based backups.
				* `logical_bytes_transferred` - (Integer) Specifies the logical bytes transferred.
				* `logical_size_bytes` - (Integer) Specifies the logicalSizeBytes.
				* `physical_bytes_transferred` - (Integer) Specifies the physical bytes transferred.
				* `total_file_count` - (Integer) Specifies the total number of file and directory entities visited in this backup. Only applicable to file based backups.
			* `stats_task_id` - (String) Run Stats task id for archival.
			* `status` - (String) Status of the replication run for an archival target. 'Running' indicates that the run is still running. 'Canceled' indicates that the run has been canceled. 'Canceling' indicates that the run is in the process of being canceled. 'Paused' indicates that the ongoing run has been paused. 'Failed' indicates that the run has failed. 'Missed' indicates that the run was unable to take place at the scheduled time because the previous run was still happening. 'Succeeded' indicates that the run has finished successfully. 'SucceededWithWarning' indicates that the run finished successfully, but there were some warning messages. 'Skipped' indicates that the run was skipped.
			  * Constraints: Allowable values are: `Accepted`, `Running`, `Canceled`, `Canceling`, `Failed`, `Missed`, `Succeeded`, `SucceededWithWarning`, `OnHold`, `Finalizing`, `Skipped`, `Paused`.
			* `successful_app_objects_count` - (Integer) Specifies the count of app objects for which backup was successful.
			* `successful_objects_count` - (Integer) Specifies the count of objects for which backup was successful.
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
			* `worm_properties` - (List) Specifies the WORM related properties for this archive.
			Nested schema for **worm_properties**:
				* `is_archive_worm_compliant` - (Boolean) Specifies whether this archive run is WORM compliant.
				* `worm_expiry_time_usecs` - (Integer) Specifies the time at which the WORM protection expires.
				* `worm_non_compliance_reason` - (String) Specifies reason of archive not being worm compliant.
	* `cloud_spin_info` - (List) Specifies summary information about cloud spin run.
	Nested schema for **cloud_spin_info**:
		* `cloud_spin_target_results` - (List) Cloud Spin results for each Cloud Spin target.
		Nested schema for **cloud_spin_target_results**:
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
			* `cloudspin_task_id` - (String) Task ID for a CloudSpin protection run.
			* `data_lock_constraints` - (List) Specifies the dataLock constraints for local or target snapshot.
			Nested schema for **data_lock_constraints**:
				* `expiry_time_usecs` - (Integer) Specifies the expiry time of attempt in Unix epoch Timestamp (in microseconds).
				* `mode` - (String) Specifies the type of WORM retention type. <br>'Compliance' implies WORM retention is set for compliance reason. <br>'Administrative' implies WORM retention is set for administrative purposes.
				  * Constraints: Allowable values are: `Compliance`, `Administrative`.
			* `end_time_usecs` - (Integer) Specifies the end time of Cloud Spin in Unix epoch Timestamp(in microseconds) for a target.
			* `expiry_time_usecs` - (Integer) Specifies the expiry time of attempt in Unix epoch Timestamp (in microseconds) for an object.
			* `id` - (Integer) Specifies the unique id of the cloud spin entity.
			* `is_manually_deleted` - (Boolean) Specifies whether the snapshot is deleted manually.
			* `message` - (String) Message about the Cloud Spin run.
			* `name` - (String) Specifies the name of the already added cloud spin target.
			* `on_legal_hold` - (Boolean) Specifies the legal hold status for a cloud spin target.
			* `progress_task_id` - (String) Progress monitor task id for Cloud Spin run.
			* `start_time_usecs` - (Integer) Specifies the start time of Cloud Spin in Unix epoch Timestamp(in microseconds) for a target.
			* `stats` - (List) Specifies statistics about Cloud Spin data.
			Nested schema for **stats**:
				* `physical_bytes_transferred` - (Integer) Specifies the physical bytes transferred.
			* `status` - (String) Status of the Cloud Spin for a target. 'Running' indicates that the run is still running. 'Canceled' indicates that the run has been canceled. 'Canceling' indicates that the run is in the process of being canceled. 'Paused' indicates that the ongoing run has been paused. 'Failed' indicates that the run has failed. 'Missed' indicates that the run was unable to take place at the scheduled time because the previous run was still happening. 'Succeeded' indicates that the run has finished successfully. 'SucceededWithWarning' indicates that the run finished successfully, but there were some warning messages. 'Skipped' indicates that the run was skipped.
			  * Constraints: Allowable values are: `Accepted`, `Running`, `Canceled`, `Canceling`, `Failed`, `Missed`, `Succeeded`, `SucceededWithWarning`, `OnHold`, `Finalizing`, `Skipped`, `Paused`.
	* `environment` - (String) Specifies the environment of the Protection Group.
	* `externally_triggered_backup_tag` - (String) The tag of externally triggered backup job.
	* `has_local_snapshot` - (Boolean) Specifies whether the run has a local snapshot. For cloud retrieved runs there may not be local snapshots.
	* `id` - (String) Specifies the ID of the Protection Group run.
	* `is_cloud_archival_direct` - (Boolean) Specifies whether the run is a CAD run if cloud archive direct feature is enabled. If this field is true, the primary backup copy will only be available at the given archived location.
	* `is_local_snapshots_deleted` - (Boolean) Specifies if snapshots for this run has been deleted.
	* `is_replication_run` - (Boolean) Specifies if this protection run is a replication run.
	* `local_backup_info` - (List) Specifies summary information about local snapshot run across all objects.
	Nested schema for **local_backup_info**:
		* `cancelled_app_objects_count` - (Integer) Specifies the count of app objects for which backup was cancelled.
		* `cancelled_objects_count` - (Integer) Specifies the count of objects for which backup was cancelled.
		* `data_lock` - (String) This field is deprecated. Use DataLockConstraints field instead.
		  * Constraints: Allowable values are: `Compliance`, `Administrative`.
		* `data_lock_constraints` - (List) Specifies the dataLock constraints for local or target snapshot.
		Nested schema for **data_lock_constraints**:
			* `expiry_time_usecs` - (Integer) Specifies the expiry time of attempt in Unix epoch Timestamp (in microseconds).
			* `mode` - (String) Specifies the type of WORM retention type. <br>'Compliance' implies WORM retention is set for compliance reason. <br>'Administrative' implies WORM retention is set for administrative purposes.
			  * Constraints: Allowable values are: `Compliance`, `Administrative`.
		* `end_time_usecs` - (Integer) Specifies the end time of backup run in Unix epoch Timestamp(in microseconds).
		* `failed_app_objects_count` - (Integer) Specifies the count of app objects for which backup failed.
		* `failed_objects_count` - (Integer) Specifies the count of objects for which backup failed.
		* `indexing_task_id` - (String) Progress monitor task for indexing.
		* `is_sla_violated` - (Boolean) Indicated if SLA has been violated for this run.
		* `local_snapshot_stats` - (List) Specifies statistics about local snapshot.
		Nested schema for **local_snapshot_stats**:
			* `bytes_read` - (Integer) Specifies total logical bytes read for creating the snapshot.
			* `bytes_written` - (Integer) Specifies total size of data in bytes written after taking backup.
			* `logical_size_bytes` - (Integer) Specifies total logical size of object(s) in bytes.
		* `local_task_id` - (String) Task ID for a local protection run.
		* `messages` - (List) Message about the backup run.
		* `progress_task_id` - (String) Progress monitor task id for local backup run.
		* `run_type` - (String) Type of Protection Group run. 'kRegular' indicates an incremental (CBT) backup. Incremental backups utilizing CBT (if supported) are captured of the target protection objects. The first run of a kRegular schedule captures all the blocks. 'kFull' indicates a full (no CBT) backup. A complete backup (all blocks) of the target protection objects are always captured and Change Block Tracking (CBT) is not utilized. 'kLog' indicates a Database Log backup. Capture the database transaction logs to allow rolling back to a specific point in time. 'kSystem' indicates system volume backup. It produces an image for bare metal recovery. 'kStorageArraySnapshot' indicates storage array snapshot backup.
		  * Constraints: Allowable values are: `kRegular`, `kFull`, `kLog`, `kSystem`, `kHydrateCDP`, `kStorageArraySnapshot`.
		* `skipped_objects_count` - (Integer) Specifies the count of objects for which backup was skipped.
		* `start_time_usecs` - (Integer) Specifies the start time of backup run in Unix epoch Timestamp(in microseconds).
		* `stats_task_id` - (String) Stats task id for local backup run.
		* `status` - (String) Status of the backup run. 'Running' indicates that the run is still running. 'Canceled' indicates that the run has been canceled. 'Canceling' indicates that the run is in the process of being canceled. 'Paused' indicates that the ongoing run has been paused. 'Failed' indicates that the run has failed. 'Missed' indicates that the run was unable to take place at the scheduled time because the previous run was still happening. 'Succeeded' indicates that the run has finished successfully. 'SucceededWithWarning' indicates that the run finished successfully, but there were some warning messages. 'Skipped' indicates that the run was skipped.
		  * Constraints: Allowable values are: `Accepted`, `Running`, `Canceled`, `Canceling`, `Failed`, `Missed`, `Succeeded`, `SucceededWithWarning`, `OnHold`, `Finalizing`, `Skipped`, `Paused`.
		* `successful_app_objects_count` - (Integer) Specifies the count of app objects for which backup was successful.
		* `successful_objects_count` - (Integer) Specifies the count of objects for which backup was successful.
	* `objects` - (List) Snapahot, replication, archival results for each object.
	Nested schema for **objects**:
		* `archival_info` - (List) Specifies information about archival run for an object.
		Nested schema for **archival_info**:
			* `archival_target_results` - (List) Archival result for an archival target.
			Nested schema for **archival_target_results**:
				* `archival_task_id` - (String) Specifies the archival task id. This is a protection group UID which only applies when archival type is 'Tape'.
				* `cancelled_app_objects_count` - (Integer) Specifies the count of app objects for which backup was cancelled.
				* `cancelled_objects_count` - (Integer) Specifies the count of objects for which backup was cancelled.
				* `data_lock_constraints` - (List) Specifies the dataLock constraints for local or target snapshot.
				Nested schema for **data_lock_constraints**:
					* `expiry_time_usecs` - (Integer) Specifies the expiry time of attempt in Unix epoch Timestamp (in microseconds).
					* `mode` - (String) Specifies the type of WORM retention type. <br>'Compliance' implies WORM retention is set for compliance reason. <br>'Administrative' implies WORM retention is set for administrative purposes.
					  * Constraints: Allowable values are: `Compliance`, `Administrative`.
				* `end_time_usecs` - (Integer) Specifies the end time of replication run in Unix epoch Timestamp(in microseconds) for an archival target.
				* `expiry_time_usecs` - (Integer) Specifies the expiry time of attempt in Unix epoch Timestamp (in microseconds).
				* `failed_app_objects_count` - (Integer) Specifies the count of app objects for which backup failed.
				* `failed_objects_count` - (Integer) Specifies the count of objects for which backup failed.
				* `indexing_task_id` - (String) Progress monitor task for indexing.
				* `is_cad_archive` - (Boolean) Whether this is CAD archive or not.
				* `is_forever_incremental` - (Boolean) Whether this is forever incremental or not.
				* `is_incremental` - (Boolean) Whether this is an incremental archive. If set to true, this is an incremental archive, otherwise this is a full archive.
				* `is_manually_deleted` - (Boolean) Specifies whether the snapshot is deleted manually.
				* `is_sla_violated` - (Boolean) Indicated if SLA has been violated for this run.
				* `message` - (String) Message about the archival run.
				* `on_legal_hold` - (Boolean) Specifies the legal hold status for a archival target.
				* `ownership_context` - (String) Specifies the ownership context for the target.
				  * Constraints: Allowable values are: `Local`, `FortKnox`.
				* `progress_task_id` - (String) Progress monitor task id for archival.
				* `queued_time_usecs` - (Integer) Specifies the time when the archival is queued for schedule in Unix epoch Timestamp(in microseconds) for a target.
				* `run_type` - (String) Type of Protection Group run. 'kRegular' indicates an incremental (CBT) backup. Incremental backups utilizing CBT (if supported) are captured of the target protection objects. The first run of a kRegular schedule captures all the blocks. 'kFull' indicates a full (no CBT) backup. A complete backup (all blocks) of the target protection objects are always captured and Change Block Tracking (CBT) is not utilized. 'kLog' indicates a Database Log backup. Capture the database transaction logs to allow rolling back to a specific point in time. 'kSystem' indicates system volume backup. It produces an image for bare metal recovery.
				  * Constraints: Allowable values are: `kRegular`, `kFull`, `kLog`, `kSystem`, `kHydrateCDP`, `kStorageArraySnapshot`.
				* `snapshot_id` - (String) Snapshot id for a successful snapshot. This field will not be set if the archival Run fails to take the snapshot.
				* `start_time_usecs` - (Integer) Specifies the start time of replication run in Unix epoch Timestamp(in microseconds) for an archival target.
				* `stats` - (List) Specifies statistics about archival data.
				Nested schema for **stats**:
					* `avg_logical_transfer_rate_bps` - (Integer) Specifies the average rate of transfer in bytes per second.
					* `backup_file_count` - (Integer) Specifies the total number of file and directory entities that are backed up in this run. Only applicable to file based backups.
					* `bytes_read` - (Integer) Specifies total logical bytes read for creating the snapshot.
					* `file_walk_done` - (Boolean) Specifies whether the file system walk is done. Only applicable to file based backups.
					* `logical_bytes_transferred` - (Integer) Specifies the logical bytes transferred.
					* `logical_size_bytes` - (Integer) Specifies the logicalSizeBytes.
					* `physical_bytes_transferred` - (Integer) Specifies the physical bytes transferred.
					* `total_file_count` - (Integer) Specifies the total number of file and directory entities visited in this backup. Only applicable to file based backups.
				* `stats_task_id` - (String) Run Stats task id for archival.
				* `status` - (String) Status of the replication run for an archival target. 'Running' indicates that the run is still running. 'Canceled' indicates that the run has been canceled. 'Canceling' indicates that the run is in the process of being canceled. 'Paused' indicates that the ongoing run has been paused. 'Failed' indicates that the run has failed. 'Missed' indicates that the run was unable to take place at the scheduled time because the previous run was still happening. 'Succeeded' indicates that the run has finished successfully. 'SucceededWithWarning' indicates that the run finished successfully, but there were some warning messages. 'Skipped' indicates that the run was skipped.
				  * Constraints: Allowable values are: `Accepted`, `Running`, `Canceled`, `Canceling`, `Failed`, `Missed`, `Succeeded`, `SucceededWithWarning`, `OnHold`, `Finalizing`, `Skipped`, `Paused`.
				* `successful_app_objects_count` - (Integer) Specifies the count of app objects for which backup was successful.
				* `successful_objects_count` - (Integer) Specifies the count of objects for which backup was successful.
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
				* `worm_properties` - (List) Specifies the WORM related properties for this archive.
				Nested schema for **worm_properties**:
					* `is_archive_worm_compliant` - (Boolean) Specifies whether this archive run is WORM compliant.
					* `worm_expiry_time_usecs` - (Integer) Specifies the time at which the WORM protection expires.
					* `worm_non_compliance_reason` - (String) Specifies reason of archive not being worm compliant.
		* `cloud_spin_info` - (List) Specifies information about Cloud Spin run for an object.
		Nested schema for **cloud_spin_info**:
			* `cloud_spin_target_results` - (List) Cloud Spin result for a target.
			Nested schema for **cloud_spin_target_results**:
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
				* `cloudspin_task_id` - (String) Task ID for a CloudSpin protection run.
				* `data_lock_constraints` - (List) Specifies the dataLock constraints for local or target snapshot.
				Nested schema for **data_lock_constraints**:
					* `expiry_time_usecs` - (Integer) Specifies the expiry time of attempt in Unix epoch Timestamp (in microseconds).
					* `mode` - (String) Specifies the type of WORM retention type. <br>'Compliance' implies WORM retention is set for compliance reason. <br>'Administrative' implies WORM retention is set for administrative purposes.
					  * Constraints: Allowable values are: `Compliance`, `Administrative`.
				* `end_time_usecs` - (Integer) Specifies the end time of Cloud Spin in Unix epoch Timestamp(in microseconds) for a target.
				* `expiry_time_usecs` - (Integer) Specifies the expiry time of attempt in Unix epoch Timestamp (in microseconds) for an object.
				* `id` - (Integer) Specifies the unique id of the cloud spin entity.
				* `is_manually_deleted` - (Boolean) Specifies whether the snapshot is deleted manually.
				* `message` - (String) Message about the Cloud Spin run.
				* `name` - (String) Specifies the name of the already added cloud spin target.
				* `on_legal_hold` - (Boolean) Specifies the legal hold status for a cloud spin target.
				* `progress_task_id` - (String) Progress monitor task id for Cloud Spin run.
				* `start_time_usecs` - (Integer) Specifies the start time of Cloud Spin in Unix epoch Timestamp(in microseconds) for a target.
				* `stats` - (List) Specifies statistics about Cloud Spin data.
				Nested schema for **stats**:
					* `physical_bytes_transferred` - (Integer) Specifies the physical bytes transferred.
				* `status` - (String) Status of the Cloud Spin for a target. 'Running' indicates that the run is still running. 'Canceled' indicates that the run has been canceled. 'Canceling' indicates that the run is in the process of being canceled. 'Paused' indicates that the ongoing run has been paused. 'Failed' indicates that the run has failed. 'Missed' indicates that the run was unable to take place at the scheduled time because the previous run was still happening. 'Succeeded' indicates that the run has finished successfully. 'SucceededWithWarning' indicates that the run finished successfully, but there were some warning messages. 'Skipped' indicates that the run was skipped.
				  * Constraints: Allowable values are: `Accepted`, `Running`, `Canceled`, `Canceling`, `Failed`, `Missed`, `Succeeded`, `SucceededWithWarning`, `OnHold`, `Finalizing`, `Skipped`, `Paused`.
		* `local_snapshot_info` - (List) Specifies information about backup run for an object.
		Nested schema for **local_snapshot_info**:
			* `failed_attempts` - (List) Failed backup attempts for an object.
			Nested schema for **failed_attempts**:
				* `admitted_time_usecs` - (Integer) Specifies the time at which the backup task was admitted to run in Unix epoch Timestamp(in microseconds) for an object.
				* `end_time_usecs` - (Integer) Specifies the end time of attempt in Unix epoch Timestamp(in microseconds) for an object.
				* `message` - (String) A message about the error if encountered while performing backup.
				* `permit_grant_time_usecs` - (Integer) Specifies the time when gatekeeper permit is granted to the backup task. If the backup task is rescheduled due to errors, the field is updated to the time when permit is granted again.
				* `progress_task_id` - (String) Progress monitor task for an object.
				* `queue_duration_usecs` - (Integer) Specifies the duration between the startTime and when gatekeeper permit is granted to the backup task. If the backup task is rescheduled due to errors, the field is updated considering the time when permit is granted again. Queue duration = PermitGrantTimeUsecs - StartTimeUsecs.
				* `snapshot_creation_time_usecs` - (Integer) Specifies the time at which the source snapshot was taken in Unix epoch Timestamp(in microseconds) for an object.
				* `start_time_usecs` - (Integer) Specifies the start time of attempt in Unix epoch Timestamp(in microseconds) for an object.
				* `stats` - (List) Specifies statistics about local snapshot.
				Nested schema for **stats**:
					* `bytes_read` - (Integer) Specifies total logical bytes read for creating the snapshot.
					* `bytes_written` - (Integer) Specifies total size of data in bytes written after taking backup.
					* `logical_size_bytes` - (Integer) Specifies total logical size of object(s) in bytes.
				* `status` - (String) Status of the attempt for an object. 'Running' indicates that the run is still running. 'Canceled' indicates that the run has been canceled. 'Canceling' indicates that the run is in the process of being canceled. 'Paused' indicates that the ongoing run has been paused. 'Pausing' indicates that the ongoing run is in the process of being paused. 'Resuming' indicates that the already paused run is in the process of being running again. 'Failed' indicates that the run has failed. 'Missed' indicates that the run was unable to take place at the scheduled time because the previous run was still happening. 'Succeeded' indicates that the run has finished successfully. 'SucceededWithWarning' indicates that the run finished successfully, but there were some warning messages. 'Skipped' indicates that the run was skipped.
				  * Constraints: Allowable values are: `Accepted`, `Running`, `Canceled`, `Canceling`, `Failed`, `Missed`, `Succeeded`, `SucceededWithWarning`, `OnHold`, `Finalizing`, `Skipped`, `Paused`.
			* `snapshot_info` - (List) Snapshot info for an object.
			Nested schema for **snapshot_info**:
				* `admitted_time_usecs` - (Integer) Specifies the time at which the backup task was admitted to run in Unix epoch Timestamp(in microseconds) for an object.
				* `backup_file_count` - (Integer) The total number of file and directory entities that are backed up in this run. Only applicable to file based backups.
				* `data_lock_constraints` - (List) Specifies the dataLock constraints for local or target snapshot.
				Nested schema for **data_lock_constraints**:
					* `expiry_time_usecs` - (Integer) Specifies the expiry time of attempt in Unix epoch Timestamp (in microseconds).
					* `mode` - (String) Specifies the type of WORM retention type. <br>'Compliance' implies WORM retention is set for compliance reason. <br>'Administrative' implies WORM retention is set for administrative purposes.
					  * Constraints: Allowable values are: `Compliance`, `Administrative`.
				* `end_time_usecs` - (Integer) Specifies the end time of attempt in Unix epoch Timestamp(in microseconds) for an object.
				* `expiry_time_usecs` - (Integer) Specifies the expiry time of attempt in Unix epoch Timestamp (in microseconds) for an object.
				* `indexing_task_id` - (String) Progress monitor task for the indexing of documents in an object.
				* `is_manually_deleted` - (Boolean) Specifies whether the snapshot is deleted manually.
				* `permit_grant_time_usecs` - (Integer) Specifies the time when gatekeeper permit is granted to the backup task. If the backup task is rescheduled due to errors, the field is updated to the time when permit is granted again.
				* `progress_task_id` - (String) Progress monitor task for backup of the object.
				* `queue_duration_usecs` - (Integer) Specifies the duration between the startTime and when gatekeeper permit is granted to the backup task. If the backup task is rescheduled due to errors, the field is updated considering the time when permit is granted again. Queue duration = PermitGrantTimeUsecs - StartTimeUsecs.
				* `snapshot_creation_time_usecs` - (Integer) Specifies the time at which the source snapshot was taken in Unix epoch Timestamp(in microseconds) for an object.
				* `snapshot_id` - (String) Snapshot id for a successful snapshot. This field will not be set if the Protection Group Run has no successful attempt.
				* `start_time_usecs` - (Integer) Specifies the start time of attempt in Unix epoch Timestamp(in microseconds) for an object.
				* `stats` - (List) Specifies statistics about local snapshot.
				Nested schema for **stats**:
					* `bytes_read` - (Integer) Specifies total logical bytes read for creating the snapshot.
					* `bytes_written` - (Integer) Specifies total size of data in bytes written after taking backup.
					* `logical_size_bytes` - (Integer) Specifies total logical size of object(s) in bytes.
				* `stats_task_id` - (String) Stats task for an object.
				* `status` - (String) Status of snapshot.
				  * Constraints: Allowable values are: `kInProgress`, `kSuccessful`, `kFailed`, `kWaitingForNextAttempt`, `kWarning`, `kCurrentAttemptPaused`, `kCurrentAttemptResuming`, `kCurrentAttemptPausing`, `kWaitingForOlderBackupRun`.
				* `status_message` - (String) A message decribing the status. This will be populated currently only for kWaitingForOlderBackupRun status.
				* `total_file_count` - (Integer) The total number of file and directory entities visited in this backup. Only applicable to file based backups.
				* `warnings` - (List) Specifies a list of warning messages.
		* `object` - (List) Specifies the Object Summary.
		Nested schema for **object**:
			* `child_objects` - (List) Specifies child object details.
			Nested schema for **child_objects**:
			* `environment` - (String) Specifies the environment of the object.
			  * Constraints: Allowable values are: `kPhysical`, `kSQL`.
			* `global_id` - (String) Specifies the global id which is a unique identifier of the object.
			* `id` - (Integer) Specifies object id.
			* `logical_size_bytes` - (Integer) Specifies the logical size of object in bytes.
			* `name` - (String) Specifies the name of the object.
			* `object_hash` - (String) Specifies the hash identifier of the object.
			* `object_type` - (String) Specifies the type of the object.
			  * Constraints: Allowable values are: `kCluster`, `kVserver`, `kVolume`, `kVCenter`, `kStandaloneHost`, `kvCloudDirector`, `kFolder`, `kDatacenter`, `kComputeResource`, `kClusterComputeResource`, `kResourcePool`, `kDatastore`, `kHostSystem`, `kVirtualMachine`, `kVirtualApp`, `kStoragePod`, `kNetwork`, `kDistributedVirtualPortgroup`, `kTagCategory`, `kTag`, `kOpaqueNetwork`, `kOrganization`, `kVirtualDatacenter`, `kCatalog`, `kOrgMetadata`, `kStoragePolicy`, `kVirtualAppTemplate`, `kDomain`, `kOutlook`, `kMailbox`, `kUsers`, `kGroups`, `kSites`, `kUser`, `kGroup`, `kSite`, `kApplication`, `kGraphUser`, `kPublicFolders`, `kPublicFolder`, `kTeams`, `kTeam`, `kRootPublicFolder`, `kO365Exchange`, `kO365OneDrive`, `kO365Sharepoint`, `kKeyspace`, `kTable`, `kDatabase`, `kCollection`, `kBucket`, `kNamespace`, `kSCVMMServer`, `kStandaloneCluster`, `kHostGroup`, `kHypervHost`, `kHostCluster`, `kCustomProperty`, `kTenant`, `kSubscription`, `kResourceGroup`, `kStorageAccount`, `kStorageKey`, `kStorageContainer`, `kStorageBlob`, `kNetworkSecurityGroup`, `kVirtualNetwork`, `kSubnet`, `kComputeOptions`, `kSnapshotManagerPermit`, `kAvailabilitySet`, `kOVirtManager`, `kHost`, `kStorageDomain`, `kVNicProfile`, `kIAMUser`, `kRegion`, `kAvailabilityZone`, `kEC2Instance`, `kVPC`, `kInstanceType`, `kKeyPair`, `kRDSOptionGroup`, `kRDSParameterGroup`, `kRDSInstance`, `kRDSSubnet`, `kRDSTag`, `kAuroraTag`, `kAuroraCluster`, `kAccount`, `kSubTaskPermit`, `kS3Bucket`, `kS3Tag`, `kKmsKey`, `kProject`, `kLabel`, `kMetadata`, `kVPCConnector`, `kPrismCentral`, `kOtherHypervisorCluster`, `kZone`, `kMountPoint`, `kStorageArray`, `kFileSystem`, `kContainer`, `kFilesystem`, `kFileset`, `kPureProtectionGroup`, `kVolumeGroup`, `kStoragePool`, `kViewBox`, `kView`, `kWindowsCluster`, `kOracleRACCluster`, `kOracleAPCluster`, `kService`, `kPVC`, `kPersistentVolumeClaim`, `kPersistentVolume`, `kRootContainer`, `kDAGRootContainer`, `kExchangeNode`, `kExchangeDAGDatabaseCopy`, `kExchangeStandaloneDatabase`, `kExchangeDAG`, `kExchangeDAGDatabase`, `kDomainController`, `kInstance`, `kAAG`, `kAAGRootContainer`, `kAAGDatabase`, `kRACRootContainer`, `kTableSpace`, `kPDB`, `kObject`, `kOrg`, `kAppInstance`.
			* `os_type` - (String) Specifies the operating system type of the object.
			  * Constraints: Allowable values are: `kLinux`, `kWindows`.
			* `protection_type` - (String) Specifies the protection type of the object if any.
			  * Constraints: Allowable values are: `kAgent`, `kNative`, `kSnapshotManager`, `kRDSSnapshotManager`, `kAuroraSnapshotManager`, `kAwsS3`, `kAwsRDSPostgresBackup`, `kAwsAuroraPostgres`, `kAwsRDSPostgres`, `kAzureSQL`, `kFile`, `kVolume`.
			* `sharepoint_site_summary` - (List) Specifies the common parameters for Sharepoint site objects.
			Nested schema for **sharepoint_site_summary**:
				* `site_web_url` - (String) Specifies the web url for the Sharepoint site.
			* `source_id` - (Integer) Specifies registered source id to which object belongs.
			* `source_name` - (String) Specifies registered source name to which object belongs.
			* `uuid` - (String) Specifies the uuid which is a unique identifier of the object.
			* `v_center_summary` - (List)
			Nested schema for **v_center_summary**:
				* `is_cloud_env` - (Boolean) Specifies that registered vCenter source is a VMC (VMware Cloud) environment or not.
			* `windows_cluster_summary` - (List)
			Nested schema for **windows_cluster_summary**:
				* `cluster_source_type` - (String) Specifies the type of cluster resource this source represents.
		* `on_legal_hold` - (Boolean) Specifies if object's snapshot is on legal hold.
		* `original_backup_info` - (List) Specifies information about backup run for an object.
		Nested schema for **original_backup_info**:
			* `failed_attempts` - (List) Failed backup attempts for an object.
			Nested schema for **failed_attempts**:
				* `admitted_time_usecs` - (Integer) Specifies the time at which the backup task was admitted to run in Unix epoch Timestamp(in microseconds) for an object.
				* `end_time_usecs` - (Integer) Specifies the end time of attempt in Unix epoch Timestamp(in microseconds) for an object.
				* `message` - (String) A message about the error if encountered while performing backup.
				* `permit_grant_time_usecs` - (Integer) Specifies the time when gatekeeper permit is granted to the backup task. If the backup task is rescheduled due to errors, the field is updated to the time when permit is granted again.
				* `progress_task_id` - (String) Progress monitor task for an object.
				* `queue_duration_usecs` - (Integer) Specifies the duration between the startTime and when gatekeeper permit is granted to the backup task. If the backup task is rescheduled due to errors, the field is updated considering the time when permit is granted again. Queue duration = PermitGrantTimeUsecs - StartTimeUsecs.
				* `snapshot_creation_time_usecs` - (Integer) Specifies the time at which the source snapshot was taken in Unix epoch Timestamp(in microseconds) for an object.
				* `start_time_usecs` - (Integer) Specifies the start time of attempt in Unix epoch Timestamp(in microseconds) for an object.
				* `stats` - (List) Specifies statistics about local snapshot.
				Nested schema for **stats**:
					* `bytes_read` - (Integer) Specifies total logical bytes read for creating the snapshot.
					* `bytes_written` - (Integer) Specifies total size of data in bytes written after taking backup.
					* `logical_size_bytes` - (Integer) Specifies total logical size of object(s) in bytes.
				* `status` - (String) Status of the attempt for an object. 'Running' indicates that the run is still running. 'Canceled' indicates that the run has been canceled. 'Canceling' indicates that the run is in the process of being canceled. 'Paused' indicates that the ongoing run has been paused. 'Pausing' indicates that the ongoing run is in the process of being paused. 'Resuming' indicates that the already paused run is in the process of being running again. 'Failed' indicates that the run has failed. 'Missed' indicates that the run was unable to take place at the scheduled time because the previous run was still happening. 'Succeeded' indicates that the run has finished successfully. 'SucceededWithWarning' indicates that the run finished successfully, but there were some warning messages. 'Skipped' indicates that the run was skipped.
				  * Constraints: Allowable values are: `Accepted`, `Running`, `Canceled`, `Canceling`, `Failed`, `Missed`, `Succeeded`, `SucceededWithWarning`, `OnHold`, `Finalizing`, `Skipped`, `Paused`.
			* `snapshot_info` - (List) Snapshot info for an object.
			Nested schema for **snapshot_info**:
				* `admitted_time_usecs` - (Integer) Specifies the time at which the backup task was admitted to run in Unix epoch Timestamp(in microseconds) for an object.
				* `backup_file_count` - (Integer) The total number of file and directory entities that are backed up in this run. Only applicable to file based backups.
				* `data_lock_constraints` - (List) Specifies the dataLock constraints for local or target snapshot.
				Nested schema for **data_lock_constraints**:
					* `expiry_time_usecs` - (Integer) Specifies the expiry time of attempt in Unix epoch Timestamp (in microseconds).
					* `mode` - (String) Specifies the type of WORM retention type. <br>'Compliance' implies WORM retention is set for compliance reason. <br>'Administrative' implies WORM retention is set for administrative purposes.
					  * Constraints: Allowable values are: `Compliance`, `Administrative`.
				* `end_time_usecs` - (Integer) Specifies the end time of attempt in Unix epoch Timestamp(in microseconds) for an object.
				* `expiry_time_usecs` - (Integer) Specifies the expiry time of attempt in Unix epoch Timestamp (in microseconds) for an object.
				* `indexing_task_id` - (String) Progress monitor task for the indexing of documents in an object.
				* `is_manually_deleted` - (Boolean) Specifies whether the snapshot is deleted manually.
				* `permit_grant_time_usecs` - (Integer) Specifies the time when gatekeeper permit is granted to the backup task. If the backup task is rescheduled due to errors, the field is updated to the time when permit is granted again.
				* `progress_task_id` - (String) Progress monitor task for backup of the object.
				* `queue_duration_usecs` - (Integer) Specifies the duration between the startTime and when gatekeeper permit is granted to the backup task. If the backup task is rescheduled due to errors, the field is updated considering the time when permit is granted again. Queue duration = PermitGrantTimeUsecs - StartTimeUsecs.
				* `snapshot_creation_time_usecs` - (Integer) Specifies the time at which the source snapshot was taken in Unix epoch Timestamp(in microseconds) for an object.
				* `snapshot_id` - (String) Snapshot id for a successful snapshot. This field will not be set if the Protection Group Run has no successful attempt.
				* `start_time_usecs` - (Integer) Specifies the start time of attempt in Unix epoch Timestamp(in microseconds) for an object.
				* `stats` - (List) Specifies statistics about local snapshot.
				Nested schema for **stats**:
					* `bytes_read` - (Integer) Specifies total logical bytes read for creating the snapshot.
					* `bytes_written` - (Integer) Specifies total size of data in bytes written after taking backup.
					* `logical_size_bytes` - (Integer) Specifies total logical size of object(s) in bytes.
				* `stats_task_id` - (String) Stats task for an object.
				* `status` - (String) Status of snapshot.
				  * Constraints: Allowable values are: `kInProgress`, `kSuccessful`, `kFailed`, `kWaitingForNextAttempt`, `kWarning`, `kCurrentAttemptPaused`, `kCurrentAttemptResuming`, `kCurrentAttemptPausing`, `kWaitingForOlderBackupRun`.
				* `status_message` - (String) A message decribing the status. This will be populated currently only for kWaitingForOlderBackupRun status.
				* `total_file_count` - (Integer) The total number of file and directory entities visited in this backup. Only applicable to file based backups.
				* `warnings` - (List) Specifies a list of warning messages.
		* `replication_info` - (List) Specifies information about replication run for an object.
		Nested schema for **replication_info**:
			* `replication_target_results` - (List) Replication result for a target.
			Nested schema for **replication_target_results**:
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
				* `data_lock_constraints` - (List) Specifies the dataLock constraints for local or target snapshot.
				Nested schema for **data_lock_constraints**:
					* `expiry_time_usecs` - (Integer) Specifies the expiry time of attempt in Unix epoch Timestamp (in microseconds).
					* `mode` - (String) Specifies the type of WORM retention type. <br>'Compliance' implies WORM retention is set for compliance reason. <br>'Administrative' implies WORM retention is set for administrative purposes.
					  * Constraints: Allowable values are: `Compliance`, `Administrative`.
				* `end_time_usecs` - (Integer) Specifies the end time of replication in Unix epoch Timestamp(in microseconds) for a target.
				* `entries_changed` - (Integer) Specifies the number of metadata actions completed during the protection run.
				* `expiry_time_usecs` - (Integer) Specifies the expiry time of attempt in Unix epoch Timestamp (in microseconds) for an object.
				* `is_in_bound` - (Boolean) Specifies the direction of the replication. If the snapshot is replicated to this cluster, then isInBound is true. If the snapshot is replicated from this cluster to another cluster, then isInBound is false.
				* `is_manually_deleted` - (Boolean) Specifies whether the snapshot is deleted manually.
				* `message` - (String) Message about the replication run.
				* `multi_object_replication` - (Boolean) Specifies whether view based replication was used. In this case, the view containing all objects is replicated as a whole instead of replicating on a per object basis.
				* `on_legal_hold` - (Boolean) Specifies the legal hold status for a replication target.
				* `percentage_completed` - (Integer) Specifies the progress in percentage.
				* `queued_time_usecs` - (Integer) Specifies the time when the replication is queued for schedule in Unix epoch Timestamp(in microseconds) for a target.
				* `replication_task_id` - (String) Task UID for a replication protection run. This is for tasks that are replicated from another cluster.
				* `start_time_usecs` - (Integer) Specifies the start time of replication in Unix epoch Timestamp(in microseconds) for a target.
				* `stats` - (List) Specifies statistics about replication data.
				Nested schema for **stats**:
					* `logical_bytes_transferred` - (Integer) Specifies the total logical bytes transferred.
					* `logical_size_bytes` - (Integer) Specifies the total logical size in bytes.
					* `physical_bytes_transferred` - (Integer) Specifies the total physical bytes transferred.
				* `status` - (String) Status of the replication for a target. 'Running' indicates that the run is still running. 'Canceled' indicates that the run has been canceled. 'Canceling' indicates that the run is in the process of being canceled. 'Paused' indicates that the ongoing run has been paused. 'Failed' indicates that the run has failed. 'Missed' indicates that the run was unable to take place at the scheduled time because the previous run was still happening. 'Succeeded' indicates that the run has finished successfully. 'SucceededWithWarning' indicates that the run finished successfully, but there were some warning messages. 'Skipped' indicates that the run was skipped.
				  * Constraints: Allowable values are: `Accepted`, `Running`, `Canceled`, `Canceling`, `Failed`, `Missed`, `Succeeded`, `SucceededWithWarning`, `OnHold`, `Finalizing`, `Skipped`, `Paused`.
	* `on_legal_hold` - (Boolean) Specifies if the Protection Run is on legal hold.
	* `origin_cluster_identifier` - (List) Specifies the information about a cluster.
	Nested schema for **origin_cluster_identifier**:
		* `cluster_id` - (Integer) Specifies the id of the cluster.
		* `cluster_incarnation_id` - (Integer) Specifies the incarnation id of the cluster.
		* `cluster_name` - (String) Specifies the name of the cluster.
	* `origin_protection_group_id` - (String) ProtectionGroupId to which this run belongs on the primary cluster if this run is a replication run.
	* `original_backup_info` - (List) Specifies summary information about local snapshot run across all objects.
	Nested schema for **original_backup_info**:
		* `cancelled_app_objects_count` - (Integer) Specifies the count of app objects for which backup was cancelled.
		* `cancelled_objects_count` - (Integer) Specifies the count of objects for which backup was cancelled.
		* `data_lock` - (String) This field is deprecated. Use DataLockConstraints field instead.
		  * Constraints: Allowable values are: `Compliance`, `Administrative`.
		* `data_lock_constraints` - (List) Specifies the dataLock constraints for local or target snapshot.
		Nested schema for **data_lock_constraints**:
			* `expiry_time_usecs` - (Integer) Specifies the expiry time of attempt in Unix epoch Timestamp (in microseconds).
			* `mode` - (String) Specifies the type of WORM retention type. <br>'Compliance' implies WORM retention is set for compliance reason. <br>'Administrative' implies WORM retention is set for administrative purposes.
			  * Constraints: Allowable values are: `Compliance`, `Administrative`.
		* `end_time_usecs` - (Integer) Specifies the end time of backup run in Unix epoch Timestamp(in microseconds).
		* `failed_app_objects_count` - (Integer) Specifies the count of app objects for which backup failed.
		* `failed_objects_count` - (Integer) Specifies the count of objects for which backup failed.
		* `indexing_task_id` - (String) Progress monitor task for indexing.
		* `is_sla_violated` - (Boolean) Indicated if SLA has been violated for this run.
		* `local_snapshot_stats` - (List) Specifies statistics about local snapshot.
		Nested schema for **local_snapshot_stats**:
			* `bytes_read` - (Integer) Specifies total logical bytes read for creating the snapshot.
			* `bytes_written` - (Integer) Specifies total size of data in bytes written after taking backup.
			* `logical_size_bytes` - (Integer) Specifies total logical size of object(s) in bytes.
		* `local_task_id` - (String) Task ID for a local protection run.
		* `messages` - (List) Message about the backup run.
		* `progress_task_id` - (String) Progress monitor task id for local backup run.
		* `run_type` - (String) Type of Protection Group run. 'kRegular' indicates an incremental (CBT) backup. Incremental backups utilizing CBT (if supported) are captured of the target protection objects. The first run of a kRegular schedule captures all the blocks. 'kFull' indicates a full (no CBT) backup. A complete backup (all blocks) of the target protection objects are always captured and Change Block Tracking (CBT) is not utilized. 'kLog' indicates a Database Log backup. Capture the database transaction logs to allow rolling back to a specific point in time. 'kSystem' indicates system volume backup. It produces an image for bare metal recovery. 'kStorageArraySnapshot' indicates storage array snapshot backup.
		  * Constraints: Allowable values are: `kRegular`, `kFull`, `kLog`, `kSystem`, `kHydrateCDP`, `kStorageArraySnapshot`.
		* `skipped_objects_count` - (Integer) Specifies the count of objects for which backup was skipped.
		* `start_time_usecs` - (Integer) Specifies the start time of backup run in Unix epoch Timestamp(in microseconds).
		* `stats_task_id` - (String) Stats task id for local backup run.
		* `status` - (String) Status of the backup run. 'Running' indicates that the run is still running. 'Canceled' indicates that the run has been canceled. 'Canceling' indicates that the run is in the process of being canceled. 'Paused' indicates that the ongoing run has been paused. 'Failed' indicates that the run has failed. 'Missed' indicates that the run was unable to take place at the scheduled time because the previous run was still happening. 'Succeeded' indicates that the run has finished successfully. 'SucceededWithWarning' indicates that the run finished successfully, but there were some warning messages. 'Skipped' indicates that the run was skipped.
		  * Constraints: Allowable values are: `Accepted`, `Running`, `Canceled`, `Canceling`, `Failed`, `Missed`, `Succeeded`, `SucceededWithWarning`, `OnHold`, `Finalizing`, `Skipped`, `Paused`.
		* `successful_app_objects_count` - (Integer) Specifies the count of app objects for which backup was successful.
		* `successful_objects_count` - (Integer) Specifies the count of objects for which backup was successful.
	* `permissions` - (List) Specifies the list of tenants that have permissions for this protection group run.
	Nested schema for **permissions**:
		* `created_at_time_msecs` - (Integer) Epoch time when tenant was created.
		* `deleted_at_time_msecs` - (Integer) Epoch time when tenant was last updated.
		* `description` - (String) Description about the tenant.
		* `external_vendor_metadata` - (List) Specifies the additional metadata for the tenant that is specifically set by the external vendors who are responsible for managing tenants. This field will only applicable if tenant creation is happening for a specially provisioned clusters for external vendors.
		Nested schema for **external_vendor_metadata**:
			* `ibm_tenant_metadata_params` - (List) Specifies the additional metadata for the tenant that is specifically set by the external vendor of type 'IBM'.
			Nested schema for **ibm_tenant_metadata_params**:
				* `account_id` - (String) Specifies the unique identifier of the IBM's account ID.
				* `crn` - (String) Specifies the unique CRN associated with the tenant.
				* `custom_properties` - (List) Specifies the list of custom properties associated with the tenant. External vendors can choose to set any properties inside following list. Note that the fields set inside the following will not be available for direct filtering. API callers should make sure that no sensitive information such as passwords is sent in these fields.
				Nested schema for **custom_properties**:
					* `key` - (String) Specifies the unique key for custom property.
					* `value` - (String) Specifies the value for the above custom key.
				* `liveness_mode` - (String) Specifies the current liveness mode of the tenant. This mode may change based on AZ failures when vendor chooses to failover or failback the tenants to other AZs.
				  * Constraints: Allowable values are: `Active`, `Standby`.
				* `metrics_config` - (List) Specifies the metadata for metrics configuration. The metadata defined here will be used by cluster to send the usgae metrics to IBM cloud metering service for calculating the tenant billing.
				Nested schema for **metrics_config**:
					* `cos_resource_config` - (List) Specifies the details of COS resource configuration required for posting metrics and trackinb billing information for IBM tenants.
					Nested schema for **cos_resource_config**:
						* `resource_url` - (String) Specifies the resource COS resource configuration endpoint that will be used for fetching bucket usage for a given tenant.
					* `iam_metrics_config` - (List) Specifies the IAM configuration that will be used for accessing the billing service in IBM cloud.
					Nested schema for **iam_metrics_config**:
						* `billing_api_key_secret_id` - (String) Specifies Id of the secret that contains the API key.
						* `iam_url` - (String) Specifies the IAM URL needed to fetch the operator token from IBM. The operator token is needed to make service API calls to IBM billing service.
					* `metering_config` - (List) Specifies the metering configuration that will be used for IBM cluster to send the billing details to IBM billing service.
					Nested schema for **metering_config**:
						* `part_ids` - (List) Specifies the list of part identifiers used for metrics identification.
						  * Constraints: Allowable list items are: `USAGETERABYTE`. The minimum length is `1` item.
						* `submission_interval_in_secs` - (Integer) Specifies the frequency in seconds at which the metrics will be pushed to IBM billing service from cluster.
						* `url` - (String) Specifies the base metering URL that will be used by cluster to send the billing information.
				* `ownership_mode` - (String) Specifies the current ownership mode for the tenant. The ownership of the tenant represents the active role for functioning of the tenant.
				  * Constraints: Allowable values are: `Primary`, `Secondary`.
				* `plan_id` - (String) Specifies the Plan Id associated with the tenant. This field is introduced for tracking purposes inside IBM enviournment.
				* `resource_group_id` - (String) Specifies the Resource Group ID associated with the tenant.
				* `resource_instance_id` - (String) Specifies the Resource Instance ID associated with the tenant. This field is introduced for tracking purposes inside IBM enviournment.
			* `type` - (String) Specifies the type of the external vendor. The type specific parameters must be specified the provided type.
			  * Constraints: Allowable values are: `IBM`.
		* `id` - (String) The tenant id.
		* `is_managed_on_helios` - (Boolean) Flag to indicate if tenant is managed on helios.
		* `last_updated_at_time_msecs` - (Integer) Epoch time when tenant was last updated.
		* `name` - (String) Name of the Tenant.
		* `network` - (List) Networking information about a Tenant on a Cluster.
		Nested schema for **network**:
			* `cluster_hostname` - (String) The hostname for Cohesity cluster as seen by tenants and as is routable from the tenant's network. Tenant's VLAN's hostname, if available can be used instead but it is mandatory to provide this value if there's no VLAN hostname to use. Also, when set, this field would take precedence over VLAN hostname.
			* `cluster_ips` - (List) Set of IPs as seen from the tenant's network for the Cohesity cluster. Only one from 'clusterHostname' and 'clusterIps' is needed.
			* `connector_enabled` - (Boolean) Whether connector (hybrid extender) is enabled.
		* `status` - (String) Current Status of the Tenant.
		  * Constraints: Allowable values are: `Active`, `Inactive`, `MarkedForDeletion`, `Deleted`.
	* `protection_group_id` - (String) ProtectionGroupId to which this run belongs.
	* `protection_group_instance_id` - (Integer) Protection Group instance Id. This field will be removed later.
	* `protection_group_name` - (String) Name of the Protection Group to which this run belongs.
	* `replication_info` - (List) Specifies summary information about replication run.
	Nested schema for **replication_info**:
		* `replication_target_results` - (List) Replication results for each replication target.
		Nested schema for **replication_target_results**:
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
			* `data_lock_constraints` - (List) Specifies the dataLock constraints for local or target snapshot.
			Nested schema for **data_lock_constraints**:
				* `expiry_time_usecs` - (Integer) Specifies the expiry time of attempt in Unix epoch Timestamp (in microseconds).
				* `mode` - (String) Specifies the type of WORM retention type. <br>'Compliance' implies WORM retention is set for compliance reason. <br>'Administrative' implies WORM retention is set for administrative purposes.
				  * Constraints: Allowable values are: `Compliance`, `Administrative`.
			* `end_time_usecs` - (Integer) Specifies the end time of replication in Unix epoch Timestamp(in microseconds) for a target.
			* `entries_changed` - (Integer) Specifies the number of metadata actions completed during the protection run.
			* `expiry_time_usecs` - (Integer) Specifies the expiry time of attempt in Unix epoch Timestamp (in microseconds) for an object.
			* `is_in_bound` - (Boolean) Specifies the direction of the replication. If the snapshot is replicated to this cluster, then isInBound is true. If the snapshot is replicated from this cluster to another cluster, then isInBound is false.
			* `is_manually_deleted` - (Boolean) Specifies whether the snapshot is deleted manually.
			* `message` - (String) Message about the replication run.
			* `multi_object_replication` - (Boolean) Specifies whether view based replication was used. In this case, the view containing all objects is replicated as a whole instead of replicating on a per object basis.
			* `on_legal_hold` - (Boolean) Specifies the legal hold status for a replication target.
			* `percentage_completed` - (Integer) Specifies the progress in percentage.
			* `queued_time_usecs` - (Integer) Specifies the time when the replication is queued for schedule in Unix epoch Timestamp(in microseconds) for a target.
			* `replication_task_id` - (String) Task UID for a replication protection run. This is for tasks that are replicated from another cluster.
			* `start_time_usecs` - (Integer) Specifies the start time of replication in Unix epoch Timestamp(in microseconds) for a target.
			* `stats` - (List) Specifies statistics about replication data.
			Nested schema for **stats**:
				* `logical_bytes_transferred` - (Integer) Specifies the total logical bytes transferred.
				* `logical_size_bytes` - (Integer) Specifies the total logical size in bytes.
				* `physical_bytes_transferred` - (Integer) Specifies the total physical bytes transferred.
			* `status` - (String) Status of the replication for a target. 'Running' indicates that the run is still running. 'Canceled' indicates that the run has been canceled. 'Canceling' indicates that the run is in the process of being canceled. 'Paused' indicates that the ongoing run has been paused. 'Failed' indicates that the run has failed. 'Missed' indicates that the run was unable to take place at the scheduled time because the previous run was still happening. 'Succeeded' indicates that the run has finished successfully. 'SucceededWithWarning' indicates that the run finished successfully, but there were some warning messages. 'Skipped' indicates that the run was skipped.
			  * Constraints: Allowable values are: `Accepted`, `Running`, `Canceled`, `Canceling`, `Failed`, `Missed`, `Succeeded`, `SucceededWithWarning`, `OnHold`, `Finalizing`, `Skipped`, `Paused`.
* `total_runs` - (Integer) Specifies the count of total runs exist for the given set of filters. The number of runs in single API call are limited and this count can be used to estimate query filter values to get next set of remaining runs. Please note that this field will only be populated if startTimeUsecs or endTimeUsecs or both are specified in query parameters.

