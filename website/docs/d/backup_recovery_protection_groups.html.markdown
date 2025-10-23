---
layout: "ibm"
page_title: "IBM : ibm_backup_recovery_protection_groups"
description: |-
  Get information about backup_recovery_protection_groups
subcategory: "IBM Backup Recovery"
---

# ibm_backup_recovery_protection_groups

Provides a read-only data source to retrieve information about backup_recovery_protection_groups. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_backup_recovery_protection_groups" "backup_recovery_protection_groups" {
	x_ibm_tenant_id = ibm_backup_recovery_protection_group.backup_recovery_protection_group_instance.x_ibm_tenant_id
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `environments` - (Optional, List) Filter by environment types such as 'kVMware', 'kView', etc. Only Protection Groups protecting the specified environment types are returned.
  * Constraints: Allowable list items are: `kPhysical`, `kSQL`, `kKubernetes`.
* `ids` - (Optional, List) Filter by a list of Protection Group ids.
* `endpoint_type` - (Optional, String) Backup Recovery Endpoint type. By default set to "public".
* `instance_id` - (Optional, String) Backup Recovery instance ID. If provided here along with region, the provider constructs the endpoint URL using them, which overrides any value set through environment variables or the `endpoints.json` file.
* `region` - (Optional, String) Backup Recovery region. If provided here along with instance_id, the provider constructs the endpoint URL using them, which overrides any value set through environment variables or the `endpoints.json` file.  
* `include_groups_with_datalock_only` - (Optional, Boolean) Whether to only return Protection Groups with a datalock.
* `include_last_run_info` - (Optional, Boolean) If true, the response will include last run info. If it is false or not specified, the last run info won't be returned.
* `is_active` - (Optional, Boolean) Filter by Inactive or Active Protection Groups. If not set, all Inactive and Active Protection Groups are returned. If true, only Active Protection Groups are returned. If false, only Inactive Protection Groups are returned. When you create a Protection Group on a Primary Cluster with a replication schedule, the Cluster creates an Inactive copy of the Protection Group on the Remote Cluster. In addition, when an Active and running Protection Group is deactivated, the Protection Group becomes Inactive.
* `is_deleted` - (Optional, Boolean) If true, return only Protection Groups that have been deleted but still have Snapshots associated with them. If false, return all Protection Groups except those Protection Groups that have been deleted and still have Snapshots associated with them. A Protection Group that is deleted with all its Snapshots is not returned for either of these cases.
* `is_last_run_sla_violated` - (Optional, Boolean) If true, return Protection Groups for which last run SLA was violated.
* `is_paused` - (Optional, Boolean) Filter by paused or non paused Protection Groups, If not set, all paused and non paused Protection Groups are returned. If true, only paused Protection Groups are returned. If false, only non paused Protection Groups are returned.
* `last_run_any_status` - (Optional, List) Filter by last any run status.<br> 'Running' indicates that the run is still running.<br> 'Canceled' indicates that the run has been canceled.<br> 'Canceling' indicates that the run is in the process of being canceled.<br> 'Failed' indicates that the run has failed.<br> 'Missed' indicates that the run was unable to take place at the scheduled time because the previous run was still happening.<br> 'Succeeded' indicates that the run has finished successfully.<br> 'SucceededWithWarning' indicates that the run finished successfully, but there were some warning messages.<br> 'Paused' indicates that the ongoing run has been paused.<br> 'Skipped' indicates that the run was skipped.
  * Constraints: Allowable list items are: `Accepted`, `Running`, `Canceled`, `Canceling`, `Failed`, `Missed`, `Succeeded`, `SucceededWithWarning`, `OnHold`, `Finalizing`, `Skipped`, `Paused`.
* `last_run_archival_status` - (Optional, List) Filter by last cloud archival run status.<br> 'Running' indicates that the run is still running.<br> 'Canceled' indicates that the run has been canceled.<br> 'Canceling' indicates that the run is in the process of being canceled.<br> 'Failed' indicates that the run has failed.<br> 'Missed' indicates that the run was unable to take place at the scheduled time because the previous run was still happening.<br> 'Succeeded' indicates that the run has finished successfully.<br> 'SucceededWithWarning' indicates that the run finished successfully, but there were some warning messages.<br> 'Paused' indicates that the ongoing run has been paused.<br> 'Skipped' indicates that the run was skipped.
  * Constraints: Allowable list items are: `Accepted`, `Running`, `Canceled`, `Canceling`, `Failed`, `Missed`, `Succeeded`, `SucceededWithWarning`, `OnHold`, `Finalizing`, `Skipped`, `Paused`.
* `last_run_cloud_spin_status` - (Optional, List) Filter by last cloud spin run status.<br> 'Running' indicates that the run is still running.<br> 'Canceled' indicates that the run has been canceled.<br> 'Canceling' indicates that the run is in the process of being canceled.<br> 'Failed' indicates that the run has failed.<br> 'Missed' indicates that the run was unable to take place at the scheduled time because the previous run was still happening.<br> 'Succeeded' indicates that the run has finished successfully.<br> 'SucceededWithWarning' indicates that the run finished successfully, but there were some warning messages.<br> 'Paused' indicates that the ongoing run has been paused.<br> 'Skipped' indicates that the run was skipped.
  * Constraints: Allowable list items are: `Accepted`, `Running`, `Canceled`, `Canceling`, `Failed`, `Missed`, `Succeeded`, `SucceededWithWarning`, `OnHold`, `Finalizing`, `Skipped`, `Paused`.
* `last_run_local_backup_status` - (Optional, List) Filter by last local backup run status.<br> 'Running' indicates that the run is still running.<br> 'Canceled' indicates that the run has been canceled.<br> 'Canceling' indicates that the run is in the process of being canceled.<br> 'Failed' indicates that the run has failed.<br> 'Missed' indicates that the run was unable to take place at the scheduled time because the previous run was still happening.<br> 'Succeeded' indicates that the run has finished successfully.<br> 'SucceededWithWarning' indicates that the run finished successfully, but there were some warning messages.<br> 'Paused' indicates that the ongoing run has been paused.<br> 'Skipped' indicates that the run was skipped.
  * Constraints: Allowable list items are: `Accepted`, `Running`, `Canceled`, `Canceling`, `Failed`, `Missed`, `Succeeded`, `SucceededWithWarning`, `OnHold`, `Finalizing`, `Skipped`, `Paused`.
* `last_run_replication_status` - (Optional, List) Filter by last remote replication run status.<br> 'Running' indicates that the run is still running.<br> 'Canceled' indicates that the run has been canceled.<br> 'Canceling' indicates that the run is in the process of being canceled.<br> 'Failed' indicates that the run has failed.<br> 'Missed' indicates that the run was unable to take place at the scheduled time because the previous run was still happening.<br> 'Succeeded' indicates that the run has finished successfully.<br> 'SucceededWithWarning' indicates that the run finished successfully, but there were some warning messages.<br> 'Paused' indicates that the ongoing run has been paused.<br> 'Skipped' indicates that the run was skipped.
  * Constraints: Allowable list items are: `Accepted`, `Running`, `Canceled`, `Canceling`, `Failed`, `Missed`, `Succeeded`, `SucceededWithWarning`, `OnHold`, `Finalizing`, `Skipped`, `Paused`.
* `names` - (Optional, List) Filter by a list of Protection Group names.
* `policy_ids` - (Optional, List) Filter by Policy ids that are associated with Protection Groups. Only Protection Groups associated with the specified Policy ids, are returned.
* `prune_excluded_source_ids` - (Optional, Boolean) If true, the response will not include the list of excluded source IDs in groups that contain this field. This can be set to true in order to improve performance if excluded source IDs are not needed by the user.
* `prune_source_ids` - (Optional, Boolean) If true, the response will exclude the list of source IDs within the group specified.
* `request_initiator_type` - (Optional, String) Specifies the type of request from UI, which is used for services like magneto to determine the priority of requests.
  * Constraints: Allowable values are: `UIUser`, `UIAuto`, `Helios`.
* `source_ids` - (Optional, List) Filter by Source ids that are associated with Protection Groups. Only Protection Groups associated with the specified Source ids, are returned.
* `use_cached_data` - (Optional, Boolean) Specifies whether we can serve the GET request from the read replica cache. There is a lag of 15 seconds between the read replica and primary data source.
* `x_ibm_tenant_id` - (Required, String) Specifies the key to be used to encrypt the source credential. If includeSourceCredentials is set to true this key must be specified.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the backup_recovery_protection_groups.
* `protection_groups` - (List) Specifies the list of Protection Groups which were returned by the request.
Nested schema for **protection_groups**:
	* `abort_in_blackouts` - (Boolean) Specifies whether currently executing jobs should abort if a blackout period specified by a policy starts. Available only if the selected policy has at least one blackout period. Default value is false.
	* `advanced_configs` - (List) Specifies the advanced configuration for a protection job.
	Nested schema for **advanced_configs**:
		* `key` - (String) key.
		* `value` - (String) value.
	* `alert_policy` - (List) Specifies a policy for alerting users of the status of a Protection Group.
	Nested schema for **alert_policy**:
		* `alert_targets` - (List) Specifies a list of targets to receive the alerts.
		Nested schema for **alert_targets**:
			* `email_address` - (String) Specifies an email address to receive an alert.
			* `language` - (String) Specifies the language of the delivery target. Default value is 'en-us'.
			  * Constraints: Allowable values are: `en-us`, `ja-jp`, `zh-cn`.
			* `recipient_type` - (String) Specifies the recipient type of email recipient. Default value is 'kTo'.
			  * Constraints: Allowable values are: `kTo`, `kCc`.
		* `backup_run_status` - (List) Specifies the run status for which the user would like to receive alerts.
		  * Constraints: Allowable list items are: `kSuccess`, `kFailure`, `kSlaViolation`, `kWarning`. The minimum length is `1` item.
		* `raise_object_level_failure_alert` - (Boolean) Specifies whether object level alerts are raised for backup failures after the backup run.
		* `raise_object_level_failure_alert_after_each_attempt` - (Boolean) Specifies whether object level alerts are raised for backup failures after each backup attempt.
		* `raise_object_level_failure_alert_after_last_attempt` - (Boolean) Specifies whether object level alerts are raised for backup failures after last backup attempt.
	* `cluster_id` - (String) Specifies the cluster ID.
	* `description` - (String) Specifies a description of the Protection Group.
	* `end_time_usecs` - (Integer) Specifies the end time in micro seconds for this Protection Group. If this is not specified, the Protection Group won't be ended.
	* `environment` - (String) Specifies the environment of the Protection Group.
	  * Constraints: Allowable values are: `kPhysical`, `kSQL`.
	* `id` - (String) Specifies the ID of the Protection Group.
	* `kubernetes_params` - (List) Specifies the parameters which are related to Kubernetes Protection Groups.
	Nested schema for **kubernetes_params**:
		* `enable_indexing` - (Boolean) Specifies if indexing of files and folders is allowed or not while backing up namespace. If allowed files and folder can be recovered.
		* `exclude_label_ids` - (List) Array of arrays of label IDs that specify labels to exclude. Optionally specify a list of labels to exclude from protecting by listing protection source ids of labels in this two dimensional array. Using this two dimensional array of label IDs, the Cluster generates a list of namespaces to exclude from protecting, which are derived from intersections of the inner arrays and union of the outer array.
		* `exclude_object_ids` - (List) Specifies the objects to be excluded in the Protection Group.
		* `exclude_params` - (List) Specifies the parameters to in/exclude objects (e.g.: volumes). An object satisfying any of these criteria will be included by this filter.
		Nested schema for **exclude_params**:
			* `label_combination_method` - (String) Whether to include all the labels or any of them while performing inclusion/exclusion of objects.
			  * Constraints: Allowable values are: `AND`, `OR`.
			* `label_vector` - (List) Array of Object to represent Label that Specify Objects (e.g.: Persistent Volumes and Persistent Volume Claims) to Include or Exclude.It will be a two-dimensional array, where each inner array will consist of a key and value representing labels. Using this two dimensional array of Labels, the Cluster generates a list of items to include in the filter, which are derived from intersections or the union of these labels, as decided by operation parameter.
			Nested schema for **label_vector**:
				* `key` - (String) The key of the label, used to identify the label.
				* `value` - (String) The value associated with the label key.
			* `objects` - (List) Array of objects that are to be included.
		* `include_params` - (List) Specifies the parameters to in/exclude objects (e.g.: volumes). An object satisfying any of these criteria will be included by this filter.
		Nested schema for **include_params**:
			* `label_combination_method` - (String) Whether to include all the labels or any of them while performing inclusion/exclusion of objects.
			  * Constraints: Allowable values are: `AND`, `OR`.
			* `label_vector` - (List) Array of Object to represent Label that Specify Objects (e.g.: Persistent Volumes and Persistent Volume Claims) to Include or Exclude.It will be a two-dimensional array, where each inner array will consist of a key and value representing labels. Using this two dimensional array of Labels, the Cluster generates a list of items to include in the filter, which are derived from intersections or the union of these labels, as decided by operation parameter.
			Nested schema for **label_vector**:
				* `key` - (String) The key of the label, used to identify the label.
				* `value` - (String) The value associated with the label key.
			* `objects` - (List) Array of objects that are to be included.
		* `label_ids` - (List) Array of array of label IDs that specify labels to protect. Optionally specify a list of labels to protect by listing protection source ids of labels in this two dimensional array. Using this two dimensional array of label IDs, the cluster generates a list of namespaces to protect, which are derived from intersections of the inner arrays and union of the outer array.
		* `leverage_csi_snapshot` - (Boolean) Specifies if CSI snapshots should be used for backup of namespaces.
		* `non_snapshot_backup` - (Boolean) Specifies if snapshot backup fails, non-snapshot backup will be proceeded.
		* `objects` - (List) Specifies the objects included in the Protection Group.
		Nested schema for **objects**:
			* `backup_only_pvc` - (Boolean) Specifies whether to backup pvc and related resources only.
			* `exclude_pvcs` - (List) Specifies a list of pvcs to exclude from being protected. This is only applicable to kubernetes.
			Nested schema for **exclude_pvcs**:
				* `id` - (Integer) Specifies the id of the pvc.
				* `name` - (String) Name of the pvc.
			* `excluded_resources` - (List) Specifies the resources to exclude during backup.
			* `id` - (Integer) Specifies the id of the object.
			* `include_pvcs` - (List) Specifies a list of Pvcs to include in the protection. This is only applicable to kubernetes.
			Nested schema for **include_pvcs**:
				* `id` - (Integer) Specifies the id of the pvc.
				* `name` - (String) Name of the pvc.
			* `included_resources` - (List) Specifies the resources to include during backup.
			* `name` - (String) Specifies the name of the object.
			* `quiesce_groups` - (List) Specifies the quiescing rules are which specified by the user for doing backup.
			Nested schema for **quiesce_groups**:
				* `quiesce_mode` - (String) Specifies quiesce mode for applying quiesce rules.
				  * Constraints: Allowable values are: `kQuiesceTogether`, `kQuiesceIndependently`.
				* `quiesce_rules` - (List) Specifies a list of quiesce rules.
				Nested schema for **quiesce_rules**:
					* `pod_selector_labels` - (List) Specifies the labels to select a pod.
					Nested schema for **pod_selector_labels**:
						* `key` - (String) The key of the label, used to identify the label.
						* `value` - (String) The value associated with the label key.
					* `post_snapshot_hooks` - (List) Specifies the hooks to be applied after taking snapshot.
					Nested schema for **post_snapshot_hooks**:
						* `commands` - (List) Specifies the commands.
						* `container` - (String) Specifies the name of the container.
						* `fail_on_error` - (Boolean) Specifies whether to fail on error or not.
						* `timeout` - (Integer) Specifies timeout for the operation.
					* `pre_snapshot_hooks` - (List) Specifies the hooks to be applied before taking snapshot.
					Nested schema for **pre_snapshot_hooks**:
						* `commands` - (List) Specifies the commands.
						* `container` - (String) Specifies the name of the container.
						* `fail_on_error` - (Boolean) Specifies whether to fail on error or not.
						* `timeout` - (Integer) Specifies timeout for the operation.
		* `source_id` - (Integer) Specifies the id of the parent of the objects.
		* `source_name` - (String) Specifies the name of the parent of the objects.
		* `vlan_params` - (List) Specifies VLAN params associated with the backup/restore operation.
		Nested schema for **vlan_params**:
			* `disable_vlan` - (Boolean) If this is set to true, then even if VLANs are configured on the system, the partition VIPs will be used for the restore.
			* `interface_name` - (String) Interface group to use for backup/restore. If this is not specified, primary interface group for the cluster will be used.
			* `vlan_id` - (Integer) If this is set, then the Cohesity host name or the IP address associated with this VLAN is used for mounting Cohesity's view on the remote host.
		* `volume_backup_failure` - (Boolean) Specifies whether to process with backup if volumes backup fails.
	* `invalid_entities` - (List) Specifies the Information about invalid entities. An entity will be considered invalid if it is part of an active protection group but has lost compatibility for the given backup type.
	Nested schema for **invalid_entities**:
		* `id` - (Integer) Specifies the ID of the object.
		* `name` - (String) Specifies the name of the object.
		* `parent_source_id` - (Integer) Specifies the id of the parent source of the object.
		* `parent_source_name` - (String) Specifies the name of the parent source of the object.
	* `is_active` - (Boolean) Specifies if the Protection Group is active or not.
	* `is_deleted` - (Boolean) Specifies if the Protection Group has been deleted.
	* `is_paused` - (Boolean) Specifies if the the Protection Group is paused. New runs are not scheduled for the paused Protection Groups. Active run if any is not impacted.
	* `is_protect_once` - (Boolean) Specifies if the the Protection Group is using a protect once type of policy. This field is helpful to identify run happen for this group.
	* `last_modified_timestamp_usecs` - (Integer) Specifies the last time this protection group was updated. If this is passed into a PUT request, then the backend will validate that the timestamp passed in matches the time that the protection group was actually last modified. If the two timestamps do not match, then the request will be rejected with a stale error.
	* `last_run` - (List) Specifies the parameters which are common between Protection Group runs of all Protection Groups.
	Nested schema for **last_run**:
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
	* `missing_entities` - (List) Specifies the Information about missing entities.
	Nested schema for **missing_entities**:
		* `id` - (Integer) Specifies the ID of the object.
		* `name` - (String) Specifies the name of the object.
		* `parent_source_id` - (Integer) Specifies the id of the parent source of the object.
		* `parent_source_name` - (String) Specifies the name of the parent source of the object.
	* `mssql_params` - (List) Specifies the parameters specific to MSSQL Protection Group.
	Nested schema for **mssql_params**:
		* `file_protection_type_params` - (List) Specifies the params to create a File based MSSQL Protection Group.
		Nested schema for **file_protection_type_params**:
			* `aag_backup_preference_type` - (String) Specifies the preference type for backing up databases that are part of an AAG. If not specified, then default preferences of the AAG server are applied. This field wont be applicable if user DB preference is set to skip AAG databases.
			  * Constraints: Allowable values are: `kPrimaryReplicaOnly`, `kSecondaryReplicaOnly`, `kPreferSecondaryReplica`, `kAnyReplica`.
			* `additional_host_params` - (List) Specifies settings which are to be applied to specific host containers in this protection group.
			Nested schema for **additional_host_params**:
				* `disable_source_side_deduplication` - (Boolean) Specifies whether or not to disable source side deduplication on this source. The default behavior is false unless the user has set 'performSourceSideDeduplication' to true.
				* `host_id` - (Integer) Specifies the id of the host container on which databases are hosted.
				* `host_name` - (String) Specifies the name of the host container on which databases are hosted.
			* `advanced_settings` - (List) This is used to regulate certain gflag values from the UI. The values passed by the user from the UI will be used for the respective gflags.
			Nested schema for **advanced_settings**:
				* `cloned_db_backup_status` - (String) Whether to report error if SQL database is cloned.
				  * Constraints: Allowable values are: `kError`, `kWarn`, `kIgnore`.
				* `db_backup_if_not_online_status` - (String) Whether to report error if SQL database is not online.
				  * Constraints: Allowable values are: `kError`, `kWarn`, `kIgnore`.
				* `missing_db_backup_status` - (String) Fail the backup job when the database is missing. The database may be missing if it is deleted or corrupted.
				  * Constraints: Allowable values are: `kError`, `kWarn`, `kIgnore`.
				* `offline_restoring_db_backup_status` - (String) Fail the backup job when database is offline or restoring.
				  * Constraints: Allowable values are: `kError`, `kWarn`, `kIgnore`.
				* `read_only_db_backup_status` - (String) Whether to skip backup for read-only SQL databases.
				  * Constraints: Allowable values are: `kError`, `kWarn`, `kIgnore`.
				* `report_all_non_autoprotect_db_errors` - (String) Whether to report error for all dbs in non-autoprotect jobs.
				  * Constraints: Allowable values are: `kError`, `kWarn`, `kIgnore`.
			* `backup_system_dbs` - (Boolean) Specifies whether to backup system databases. If not specified then parameter is set to true.
			* `exclude_filters` - (List) Specifies the list of exclusion filters applied during the group creation or edit. These exclusion filters can be wildcard supported strings or regular expressions. Objects satisfying the will filters will be excluded during backup and also auto protected objects will be ignored if filtered by any of the filters.
			Nested schema for **exclude_filters**:
				* `filter_string` - (String) Specifies the filter string using wildcard supported strings or regular expressions.
				* `is_regular_expression` - (Boolean) Specifies whether the provided filter string is a regular expression or not. This needs to be explicitly set to true if user is trying to filter by regular expressions. Not providing this value in case of regular expression can result in unintended results. The default value is assumed to be false.
				  * Constraints: The default value is `false`.
			* `full_backups_copy_only` - (Boolean) Specifies whether full backups should be copy-only.
			* `log_backup_num_streams` - (Integer) Specifies the number of streams to be used for log backups.
			* `log_backup_with_clause` - (String) Specifies the WithClause to be used for log backups.
			* `objects` - (List) Specifies the list of object params to be protected.
			  * Constraints: The minimum length is `1` item.
			Nested schema for **objects**:
				* `id` - (Integer) Specifies the ID of the object being protected. If this is a non leaf level object, then the object will be auto-protected unless leaf objects are specified for exclusion.
				* `name` - (String) Specifies the name of the object being protected.
				* `source_type` - (String) Specifies the type of source being protected.
			* `perform_source_side_deduplication` - (Boolean) Specifies whether or not to perform source side deduplication on this Protection Group.
			* `pre_post_script` - (List) Specifies the params for pre and post scripts.
			Nested schema for **pre_post_script**:
				* `post_script` - (List) Specifies the common params for PostBackup scripts.
				Nested schema for **post_script**:
					* `is_active` - (Boolean) Specifies whether the script should be enabled, default value set to true.
					* `params` - (String) Specifies the arguments or parameters and values to pass into the remote script. For example if the script expects values for the 'database' and 'user' parameters, specify the parameters and values using the following string: "database=myDatabase user=me".
					* `path` - (String) Specifies the absolute path to the script on the remote host.
					* `timeout_secs` - (Integer) Specifies the timeout of the script in seconds. The script will be killed if it exceeds this value. By default, no timeout will occur if left empty.
					  * Constraints: The minimum value is `1`.
				* `pre_script` - (List) Specifies the common params for PreBackup scripts.
				Nested schema for **pre_script**:
					* `continue_on_error` - (Boolean) Specifies if the script needs to continue even if there is an occurence of an error. If this flag is set to true, then Backup Run will start even if the pre backup script fails. If not specified or false, then backup run will not start when script fails.
					* `is_active` - (Boolean) Specifies whether the script should be enabled, default value set to true.
					* `params` - (String) Specifies the arguments or parameters and values to pass into the remote script. For example if the script expects values for the 'database' and 'user' parameters, specify the parameters and values using the following string: "database=myDatabase user=me".
					* `path` - (String) Specifies the absolute path to the script on the remote host.
					* `timeout_secs` - (Integer) Specifies the timeout of the script in seconds. The script will be killed if it exceeds this value. By default, no timeout will occur if left empty.
					  * Constraints: The minimum value is `1`.
			* `use_aag_preferences_from_server` - (Boolean) Specifies whether or not the AAG backup preferences specified on the SQL Server host should be used.
			* `user_db_backup_preference_type` - (String) Specifies the preference type for backing up user databases on the host.
			  * Constraints: Allowable values are: `kBackupAllDatabases`, `kBackupAllExceptAAGDatabases`, `kBackupOnlyAAGDatabases`.
		* `native_protection_type_params` - (List) Specifies the params to create a Native based MSSQL Protection Group.
		Nested schema for **native_protection_type_params**:
			* `aag_backup_preference_type` - (String) Specifies the preference type for backing up databases that are part of an AAG. If not specified, then default preferences of the AAG server are applied. This field wont be applicable if user DB preference is set to skip AAG databases.
			  * Constraints: Allowable values are: `kPrimaryReplicaOnly`, `kSecondaryReplicaOnly`, `kPreferSecondaryReplica`, `kAnyReplica`.
			* `advanced_settings` - (List) This is used to regulate certain gflag values from the UI. The values passed by the user from the UI will be used for the respective gflags.
			Nested schema for **advanced_settings**:
				* `cloned_db_backup_status` - (String) Whether to report error if SQL database is cloned.
				  * Constraints: Allowable values are: `kError`, `kWarn`, `kIgnore`.
				* `db_backup_if_not_online_status` - (String) Whether to report error if SQL database is not online.
				  * Constraints: Allowable values are: `kError`, `kWarn`, `kIgnore`.
				* `missing_db_backup_status` - (String) Fail the backup job when the database is missing. The database may be missing if it is deleted or corrupted.
				  * Constraints: Allowable values are: `kError`, `kWarn`, `kIgnore`.
				* `offline_restoring_db_backup_status` - (String) Fail the backup job when database is offline or restoring.
				  * Constraints: Allowable values are: `kError`, `kWarn`, `kIgnore`.
				* `read_only_db_backup_status` - (String) Whether to skip backup for read-only SQL databases.
				  * Constraints: Allowable values are: `kError`, `kWarn`, `kIgnore`.
				* `report_all_non_autoprotect_db_errors` - (String) Whether to report error for all dbs in non-autoprotect jobs.
				  * Constraints: Allowable values are: `kError`, `kWarn`, `kIgnore`.
			* `backup_system_dbs` - (Boolean) Specifies whether to backup system databases. If not specified then parameter is set to true.
			* `exclude_filters` - (List) Specifies the list of exclusion filters applied during the group creation or edit. These exclusion filters can be wildcard supported strings or regular expressions. Objects satisfying the will filters will be excluded during backup and also auto protected objects will be ignored if filtered by any of the filters.
			Nested schema for **exclude_filters**:
				* `filter_string` - (String) Specifies the filter string using wildcard supported strings or regular expressions.
				* `is_regular_expression` - (Boolean) Specifies whether the provided filter string is a regular expression or not. This needs to be explicitly set to true if user is trying to filter by regular expressions. Not providing this value in case of regular expression can result in unintended results. The default value is assumed to be false.
				  * Constraints: The default value is `false`.
			* `full_backups_copy_only` - (Boolean) Specifies whether full backups should be copy-only.
			* `log_backup_num_streams` - (Integer) Specifies the number of streams to be used for log backups.
			* `log_backup_with_clause` - (String) Specifies the WithClause to be used for log backups.
			* `num_streams` - (Integer) Specifies the number of streams to be used.
			* `objects` - (List) Specifies the list of object params to be protected.
			  * Constraints: The minimum length is `1` item.
			Nested schema for **objects**:
				* `id` - (Integer) Specifies the ID of the object being protected. If this is a non leaf level object, then the object will be auto-protected unless leaf objects are specified for exclusion.
				* `name` - (String) Specifies the name of the object being protected.
				* `source_type` - (String) Specifies the type of source being protected.
			* `pre_post_script` - (List) Specifies the params for pre and post scripts.
			Nested schema for **pre_post_script**:
				* `post_script` - (List) Specifies the common params for PostBackup scripts.
				Nested schema for **post_script**:
					* `is_active` - (Boolean) Specifies whether the script should be enabled, default value set to true.
					* `params` - (String) Specifies the arguments or parameters and values to pass into the remote script. For example if the script expects values for the 'database' and 'user' parameters, specify the parameters and values using the following string: "database=myDatabase user=me".
					* `path` - (String) Specifies the absolute path to the script on the remote host.
					* `timeout_secs` - (Integer) Specifies the timeout of the script in seconds. The script will be killed if it exceeds this value. By default, no timeout will occur if left empty.
					  * Constraints: The minimum value is `1`.
				* `pre_script` - (List) Specifies the common params for PreBackup scripts.
				Nested schema for **pre_script**:
					* `continue_on_error` - (Boolean) Specifies if the script needs to continue even if there is an occurence of an error. If this flag is set to true, then Backup Run will start even if the pre backup script fails. If not specified or false, then backup run will not start when script fails.
					* `is_active` - (Boolean) Specifies whether the script should be enabled, default value set to true.
					* `params` - (String) Specifies the arguments or parameters and values to pass into the remote script. For example if the script expects values for the 'database' and 'user' parameters, specify the parameters and values using the following string: "database=myDatabase user=me".
					* `path` - (String) Specifies the absolute path to the script on the remote host.
					* `timeout_secs` - (Integer) Specifies the timeout of the script in seconds. The script will be killed if it exceeds this value. By default, no timeout will occur if left empty.
					  * Constraints: The minimum value is `1`.
			* `use_aag_preferences_from_server` - (Boolean) Specifies whether or not the AAG backup preferences specified on the SQL Server host should be used.
			* `user_db_backup_preference_type` - (String) Specifies the preference type for backing up user databases on the host.
			  * Constraints: Allowable values are: `kBackupAllDatabases`, `kBackupAllExceptAAGDatabases`, `kBackupOnlyAAGDatabases`.
			* `with_clause` - (String) Specifies the WithClause to be used.
		* `protection_type` - (String) Specifies the MSSQL Protection Group type.
		  * Constraints: Allowable values are: `kFile`, `kVolume`, `kNative`.
		* `volume_protection_type_params` - (List) Specifies the params to create a Volume based MSSQL Protection Group.
		Nested schema for **volume_protection_type_params**:
			* `aag_backup_preference_type` - (String) Specifies the preference type for backing up databases that are part of an AAG. If not specified, then default preferences of the AAG server are applied. This field wont be applicable if user DB preference is set to skip AAG databases.
			  * Constraints: Allowable values are: `kPrimaryReplicaOnly`, `kSecondaryReplicaOnly`, `kPreferSecondaryReplica`, `kAnyReplica`.
			* `additional_host_params` - (List) Specifies settings which are to be applied to specific host containers in this protection group.
			Nested schema for **additional_host_params**:
				* `enable_system_backup` - (Boolean) Specifies whether to enable system/bmr backup using 3rd party tools installed on agent host.
				* `host_id` - (Integer) Specifies the id of the host container on which databases are hosted.
				* `host_name` - (String) Specifies the name of the host container on which databases are hosted.
				* `volume_guids` - (List) Specifies the list of volume GUIDs to be protected. If not specified, all the volumes of the host will be protected. Note that volumes of host on which databases are hosted are protected even if its not mentioned in this list.
			* `advanced_settings` - (List) This is used to regulate certain gflag values from the UI. The values passed by the user from the UI will be used for the respective gflags.
			Nested schema for **advanced_settings**:
				* `cloned_db_backup_status` - (String) Whether to report error if SQL database is cloned.
				  * Constraints: Allowable values are: `kError`, `kWarn`, `kIgnore`.
				* `db_backup_if_not_online_status` - (String) Whether to report error if SQL database is not online.
				  * Constraints: Allowable values are: `kError`, `kWarn`, `kIgnore`.
				* `missing_db_backup_status` - (String) Fail the backup job when the database is missing. The database may be missing if it is deleted or corrupted.
				  * Constraints: Allowable values are: `kError`, `kWarn`, `kIgnore`.
				* `offline_restoring_db_backup_status` - (String) Fail the backup job when database is offline or restoring.
				  * Constraints: Allowable values are: `kError`, `kWarn`, `kIgnore`.
				* `read_only_db_backup_status` - (String) Whether to skip backup for read-only SQL databases.
				  * Constraints: Allowable values are: `kError`, `kWarn`, `kIgnore`.
				* `report_all_non_autoprotect_db_errors` - (String) Whether to report error for all dbs in non-autoprotect jobs.
				  * Constraints: Allowable values are: `kError`, `kWarn`, `kIgnore`.
			* `backup_db_volumes_only` - (Boolean) Specifies whether to only backup volumes on which the specified databases reside. If not specified (default), all the volumes of the host will be protected.
			* `backup_system_dbs` - (Boolean) Specifies whether to backup system databases. If not specified then parameter is set to true.
			* `exclude_filters` - (List) Specifies the list of exclusion filters applied during the group creation or edit. These exclusion filters can be wildcard supported strings or regular expressions. Objects satisfying the will filters will be excluded during backup and also auto protected objects will be ignored if filtered by any of the filters.
			Nested schema for **exclude_filters**:
				* `filter_string` - (String) Specifies the filter string using wildcard supported strings or regular expressions.
				* `is_regular_expression` - (Boolean) Specifies whether the provided filter string is a regular expression or not. This needs to be explicitly set to true if user is trying to filter by regular expressions. Not providing this value in case of regular expression can result in unintended results. The default value is assumed to be false.
				  * Constraints: The default value is `false`.
			* `full_backups_copy_only` - (Boolean) Specifies whether full backups should be copy-only.
			* `incremental_backup_after_restart` - (Boolean) Specifies whether or to perform incremental backups the first time after a server restarts. By default, a full backup will be performed.
			* `indexing_policy` - (List) Specifies settings for indexing files found in an Object (such as a VM) so these files can be searched and recovered. This also specifies inclusion and exclusion rules that determine the directories to index.
			Nested schema for **indexing_policy**:
				* `enable_indexing` - (Boolean) Specifies if the files found in an Object (such as a VM) should be indexed. If true (the default), files are indexed.
				* `exclude_paths` - (List) Array of Excluded Directories. Specifies a list of directories to exclude from indexing.Regular expression can also be specified to provide the directory paths. Example: /Users/<wildcard>/AppData.
				* `include_paths` - (List) Array of Indexed Directories. Specifies a list of directories to index. Regular expression can also be specified to provide the directory paths. Example: /Users/<wildcard>/AppData.
			* `log_backup_num_streams` - (Integer) Specifies the number of streams to be used for log backups.
			* `log_backup_with_clause` - (String) Specifies the WithClause to be used for log backups.
			* `objects` - (List) Specifies the list of object ids to be protected.
			Nested schema for **objects**:
				* `id` - (Integer) Specifies the ID of the object being protected. If this is a non leaf level object, then the object will be auto-protected unless leaf objects are specified for exclusion.
				* `name` - (String) Specifies the name of the object being protected.
				* `source_type` - (String) Specifies the type of source being protected.
			* `pre_post_script` - (List) Specifies the params for pre and post scripts.
			Nested schema for **pre_post_script**:
				* `post_script` - (List) Specifies the common params for PostBackup scripts.
				Nested schema for **post_script**:
					* `is_active` - (Boolean) Specifies whether the script should be enabled, default value set to true.
					* `params` - (String) Specifies the arguments or parameters and values to pass into the remote script. For example if the script expects values for the 'database' and 'user' parameters, specify the parameters and values using the following string: "database=myDatabase user=me".
					* `path` - (String) Specifies the absolute path to the script on the remote host.
					* `timeout_secs` - (Integer) Specifies the timeout of the script in seconds. The script will be killed if it exceeds this value. By default, no timeout will occur if left empty.
					  * Constraints: The minimum value is `1`.
				* `pre_script` - (List) Specifies the common params for PreBackup scripts.
				Nested schema for **pre_script**:
					* `continue_on_error` - (Boolean) Specifies if the script needs to continue even if there is an occurence of an error. If this flag is set to true, then Backup Run will start even if the pre backup script fails. If not specified or false, then backup run will not start when script fails.
					* `is_active` - (Boolean) Specifies whether the script should be enabled, default value set to true.
					* `params` - (String) Specifies the arguments or parameters and values to pass into the remote script. For example if the script expects values for the 'database' and 'user' parameters, specify the parameters and values using the following string: "database=myDatabase user=me".
					* `path` - (String) Specifies the absolute path to the script on the remote host.
					* `timeout_secs` - (Integer) Specifies the timeout of the script in seconds. The script will be killed if it exceeds this value. By default, no timeout will occur if left empty.
					  * Constraints: The minimum value is `1`.
			* `use_aag_preferences_from_server` - (Boolean) Specifies whether or not the AAG backup preferences specified on the SQL Server host should be used.
			* `user_db_backup_preference_type` - (String) Specifies the preference type for backing up user databases on the host.
			  * Constraints: Allowable values are: `kBackupAllDatabases`, `kBackupAllExceptAAGDatabases`, `kBackupOnlyAAGDatabases`.
	* `name` - (String) Specifies the name of the Protection Group.
	* `num_protected_objects` - (Integer) Specifies the number of protected objects of the Protection Group.
	* `pause_in_blackouts` - (Boolean) Specifies whether currently executing jobs should be paused if a blackout period specified by a policy starts. Available only if the selected policy has at least one blackout period. Default value is false. This field should not be set to true if 'abortInBlackouts' is sent as true.
	* `permissions` - (List) Specifies the list of tenants that have permissions for this protection group.
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
	* `physical_params` - (List)
	Nested schema for **physical_params**:
		* `file_protection_type_params` - (List) Specifies the parameters which are specific to Physical related Protection Groups.
		Nested schema for **file_protection_type_params**:
			* `allow_parallel_runs` - (Boolean) Specifies whether or not this job can have parallel runs.
			* `cobmr_backup` - (Boolean) Specifies whether to take CoBMR backup.
			* `continue_on_quiesce_failure` - (Boolean) Specifies whether to continue backing up on quiesce failure.
			* `dedup_exclusion_source_ids` - (List) Specifies ids of sources for which deduplication has to be disabled.
			* `excluded_vss_writers` - (List) Specifies writer names which should be excluded from physical file based backups.
			* `global_exclude_fs` - (List) Specifies global exclude filesystems which are applied to all sources in a job.
			* `global_exclude_paths` - (List) Specifies global exclude filters which are applied to all sources in a job.
			* `ignorable_errors` - (List) Specifies the Errors to be ignored in error db.
			  * Constraints: Allowable list items are: `kEOF`, `kNonExistent`.
			* `indexing_policy` - (List) Specifies settings for indexing files found in an Object (such as a VM) so these files can be searched and recovered. This also specifies inclusion and exclusion rules that determine the directories to index.
			Nested schema for **indexing_policy**:
				* `enable_indexing` - (Boolean) Specifies if the files found in an Object (such as a VM) should be indexed. If true (the default), files are indexed.
				* `exclude_paths` - (List) Array of Excluded Directories. Specifies a list of directories to exclude from indexing.Regular expression can also be specified to provide the directory paths. Example: /Users/<wildcard>/AppData.
				* `include_paths` - (List) Array of Indexed Directories. Specifies a list of directories to index. Regular expression can also be specified to provide the directory paths. Example: /Users/<wildcard>/AppData.
			* `objects` - (List) Specifies the list of objects protected by this Protection Group.
			  * Constraints: The minimum length is `1` item.
			Nested schema for **objects**:
				* `excluded_vss_writers` - (List) Specifies writer names which should be excluded from physical file based backups.
				* `file_paths` - (List) Specifies a list of file paths to be protected by this Protection Group.
				Nested schema for **file_paths**:
					* `excluded_paths` - (List) Specifies a set of paths nested under the include path which should be excluded from the Protection Group.
					* `included_path` - (String) Specifies a path to be included on the source. All paths under this path will be included unless they are specifically mentioned in excluded paths.
					* `skip_nested_volumes` - (Boolean) Specifies whether to skip any nested volumes (both local and network) that are mounted under include path. Applicable only for windows sources.
				* `follow_nas_symlink_target` - (Boolean) Specifies whether to follow NAS target pointed by symlink for windows sources.
				* `id` - (Integer) Specifies the ID of the object protected.
				* `metadata_file_path` - (String) Specifies the path of metadatafile on source. This file contains absolute paths of files that needs to be backed up on the same source.
				* `name` - (String) Specifies the name of the object protected.
				* `nested_volume_types_to_skip` - (List) Specifies mount types of nested volumes to be skipped.
				* `uses_path_level_skip_nested_volume_setting` - (Boolean) Specifies whether path level or object level skip nested volume setting will be used.
			* `perform_brick_based_deduplication` - (Boolean) Specifies whether or not to perform brick based deduplication on this Protection Group.
			* `perform_source_side_deduplication` - (Boolean) Specifies whether or not to perform source side deduplication on this Protection Group.
			* `pre_post_script` - (List) Specifies the params for pre and post scripts.
			Nested schema for **pre_post_script**:
				* `post_script` - (List) Specifies the common params for PostBackup scripts.
				Nested schema for **post_script**:
					* `is_active` - (Boolean) Specifies whether the script should be enabled, default value set to true.
					* `params` - (String) Specifies the arguments or parameters and values to pass into the remote script. For example if the script expects values for the 'database' and 'user' parameters, specify the parameters and values using the following string: "database=myDatabase user=me".
					* `path` - (String) Specifies the absolute path to the script on the remote host.
					* `timeout_secs` - (Integer) Specifies the timeout of the script in seconds. The script will be killed if it exceeds this value. By default, no timeout will occur if left empty.
					  * Constraints: The minimum value is `1`.
				* `pre_script` - (List) Specifies the common params for PreBackup scripts.
				Nested schema for **pre_script**:
					* `continue_on_error` - (Boolean) Specifies if the script needs to continue even if there is an occurence of an error. If this flag is set to true, then Backup Run will start even if the pre backup script fails. If not specified or false, then backup run will not start when script fails.
					* `is_active` - (Boolean) Specifies whether the script should be enabled, default value set to true.
					* `params` - (String) Specifies the arguments or parameters and values to pass into the remote script. For example if the script expects values for the 'database' and 'user' parameters, specify the parameters and values using the following string: "database=myDatabase user=me".
					* `path` - (String) Specifies the absolute path to the script on the remote host.
					* `timeout_secs` - (Integer) Specifies the timeout of the script in seconds. The script will be killed if it exceeds this value. By default, no timeout will occur if left empty.
					  * Constraints: The minimum value is `1`.
			* `quiesce` - (Boolean) Specifies Whether to take app-consistent snapshots by quiescing apps and the filesystem before taking a backup.
			* `task_timeouts` - (List) Specifies the timeouts for all the objects inside this Protection Group, for both full and incremental backups.
			Nested schema for **task_timeouts**:
				* `backup_type` - (String) The scheduled backup type(kFull, kRegular etc.).
				  * Constraints: Allowable values are: `kRegular`, `kFull`, `kLog`, `kSystem`, `kHydrateCDP`, `kStorageArraySnapshot`.
				* `timeout_mins` - (Integer) Specifies the timeout in mins.
		* `protection_type` - (String) Specifies the Physical Protection Group type.
		  * Constraints: Allowable values are: `kFile`, `kVolume`.
		* `volume_protection_type_params` - (List) Specifies the parameters which are specific to Volume based physical Protection Groups.
		Nested schema for **volume_protection_type_params**:
			* `cobmr_backup` - (Boolean) Specifies whether to take a CoBMR backup.
			* `continue_on_quiesce_failure` - (Boolean) Specifies whether to continue backing up on quiesce failure.
			* `dedup_exclusion_source_ids` - (List) Specifies ids of sources for which deduplication has to be disabled.
			* `excluded_vss_writers` - (List) Specifies writer names which should be excluded from physical volume based backups.
			* `incremental_backup_after_restart` - (Boolean) Specifies whether or not to perform an incremental backup after the server restarts. This is applicable to windows environments.
			* `indexing_policy` - (List) Specifies settings for indexing files found in an Object (such as a VM) so these files can be searched and recovered. This also specifies inclusion and exclusion rules that determine the directories to index.
			Nested schema for **indexing_policy**:
				* `enable_indexing` - (Boolean) Specifies if the files found in an Object (such as a VM) should be indexed. If true (the default), files are indexed.
				* `exclude_paths` - (List) Array of Excluded Directories. Specifies a list of directories to exclude from indexing.Regular expression can also be specified to provide the directory paths. Example: /Users/<wildcard>/AppData.
				* `include_paths` - (List) Array of Indexed Directories. Specifies a list of directories to index. Regular expression can also be specified to provide the directory paths. Example: /Users/<wildcard>/AppData.
			* `objects` - (List)
			  * Constraints: The minimum length is `1` item.
			Nested schema for **objects**:
				* `enable_system_backup` - (Boolean) Specifies whether or not to take a system backup. Applicable only for windows sources.
				* `excluded_vss_writers` - (List) Specifies writer names which should be excluded from physical volume based backups for a given source.
				* `id` - (Integer) Specifies the ID of the object protected.
				* `name` - (String) Specifies the name of the object protected.
				* `volume_guids` - (List) Specifies the list of GUIDs of volumes protected. If empty, then all volumes will be protected by default.
			* `perform_source_side_deduplication` - (Boolean) Specifies whether or not to perform source side deduplication on this Protection Group.
			* `pre_post_script` - (List) Specifies the params for pre and post scripts.
			Nested schema for **pre_post_script**:
				* `post_script` - (List) Specifies the common params for PostBackup scripts.
				Nested schema for **post_script**:
					* `is_active` - (Boolean) Specifies whether the script should be enabled, default value set to true.
					* `params` - (String) Specifies the arguments or parameters and values to pass into the remote script. For example if the script expects values for the 'database' and 'user' parameters, specify the parameters and values using the following string: "database=myDatabase user=me".
					* `path` - (String) Specifies the absolute path to the script on the remote host.
					* `timeout_secs` - (Integer) Specifies the timeout of the script in seconds. The script will be killed if it exceeds this value. By default, no timeout will occur if left empty.
					  * Constraints: The minimum value is `1`.
				* `pre_script` - (List) Specifies the common params for PreBackup scripts.
				Nested schema for **pre_script**:
					* `continue_on_error` - (Boolean) Specifies if the script needs to continue even if there is an occurence of an error. If this flag is set to true, then Backup Run will start even if the pre backup script fails. If not specified or false, then backup run will not start when script fails.
					* `is_active` - (Boolean) Specifies whether the script should be enabled, default value set to true.
					* `params` - (String) Specifies the arguments or parameters and values to pass into the remote script. For example if the script expects values for the 'database' and 'user' parameters, specify the parameters and values using the following string: "database=myDatabase user=me".
					* `path` - (String) Specifies the absolute path to the script on the remote host.
					* `timeout_secs` - (Integer) Specifies the timeout of the script in seconds. The script will be killed if it exceeds this value. By default, no timeout will occur if left empty.
					  * Constraints: The minimum value is `1`.
			* `quiesce` - (Boolean) Specifies Whether to take app-consistent snapshots by quiescing apps and the filesystem before taking a backup.
	* `policy_id` - (String) Specifies the unique id of the Protection Policy associated with the Protection Group. The Policy provides retry settings Protection Schedules, Priority, SLA, etc.
	* `priority` - (String) Specifies the priority of the Protection Group.
	  * Constraints: Allowable values are: `kLow`, `kMedium`, `kHigh`.
	* `qos_policy` - (String) Specifies whether the Protection Group will be written to HDD or SSD.
	  * Constraints: Allowable values are: `kBackupHDD`, `kBackupSSD`, `kTestAndDevHigh`, `kBackupAll`.
	* `region_id` - (String) Specifies the region ID.
	* `sla` - (List) Specifies the SLA parameters for this Protection Group.
	Nested schema for **sla**:
		* `backup_run_type` - (String) Specifies the type of run this rule should apply to.
		  * Constraints: Allowable values are: `kIncremental`, `kFull`, `kLog`.
		* `sla_minutes` - (Integer) Specifies the number of minutes allotted to a run of the specified type before SLA is considered violated.
		  * Constraints: The minimum value is `1`.
	* `start_time` - (List) Specifies the time of day. Used for scheduling purposes.
	Nested schema for **start_time**:
		* `hour` - (Integer) Specifies the hour of the day (0-23).
		  * Constraints: The maximum value is `23`. The minimum value is `0`.
		* `minute` - (Integer) Specifies the minute of the hour (0-59).
		  * Constraints: The maximum value is `59`. The minimum value is `0`.
		* `time_zone` - (String) Specifies the time zone of the user. If not specified, default value is assumed as America/Los_Angeles.
		  * Constraints: The default value is `America/Los_Angeles`.

