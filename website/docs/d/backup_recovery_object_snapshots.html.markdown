---
layout: "ibm"
page_title: "IBM : ibm_backup_recovery_object_snapshots"
description: |-
  Get information about Object Snapshots.
subcategory: "IBM Backup Recovery"
---

# ibm_backup_recovery_object_snapshots

Provides a read-only data source to retrieve information about an Object Snapshots.. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_backup_recovery_object_snapshots" "backup_recovery_object_snapshots" {
	object_id = 2
	x_ibm_tenant_id = "x_ibm_tenant_id"
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `object_id` - (Required, Forces new resource, Integer) Specifies the id of the Object.
* `from_time_usecs` - (Optional, Integer) Specifies the timestamp in Unix time epoch in microseconds to filter Object's snapshots which were taken after this value.
* `endpoint_type` - (Optional, String) Backup Recovery Endpoint type. By default set to "public".
* `instance_id` - (Optional, String) Backup Recovery instance ID. If provided here along with region, the provider constructs the endpoint URL using them, which overrides any value set through environment variables or the `endpoints.json` file.
* `region` - (Optional, String) Backup Recovery region. If provided here along with instance_id, the provider constructs the endpoint URL using them, which overrides any value set through environment variables or the `endpoints.json` file.  
* `object_action_keys` - (Optional, List) Filter by ObjectActionKey, which uniquely represents the protection of an object. An object can be protected in multiple ways but at most once for a given combination of ObjectActionKey. When specified, only snapshots matching the given action keys are returned for the corresponding object.
  * Constraints: Allowable list items are: `kVMware`, `kHyperV`, `kVCD`, `kAzure`, `kGCP`, `kKVM`, `kAcropolis`, `kAWS`, `kAWSNative`, `kAwsS3`, `kAWSSnapshotManager`, `kRDSSnapshotManager`, `kAuroraSnapshotManager`, `kAwsRDSPostgresBackup`, `kAwsRDSPostgres`, `kAwsAuroraPostgres`, `kAzureNative`, `kAzureSQL`, `kAzureSnapshotManager`, `kPhysical`, `kPhysicalFiles`, `kGPFS`, `kElastifile`, `kNetapp`, `kGenericNas`, `kIsilon`, `kFlashBlade`, `kPure`, `kIbmFlashSystem`, `kSQL`, `kExchange`, `kAD`, `kOracle`, `kView`, `kRemoteAdapter`, `kO365`, `kO365PublicFolders`, `kO365Teams`, `kO365Group`, `kO365Exchange`, `kO365OneDrive`, `kO365Sharepoint`, `kKubernetes`, `kCassandra`, `kMongoDB`, `kCouchbase`, `kHdfs`, `kHive`, `kHBase`, `kSAPHANA`, `kUDA`, `kSfdc`, `kO365ExchangeCSM`, `kO365OneDriveCSM`, `kO365SharepointCSM`.
* `protection_group_ids` - (Optional, List) If specified, this returns only the snapshots of the specified object ID, which belong to the provided protection group IDs.
* `region_ids` - (Optional, List) Filter by a list of region IDs.
* `run_instance_ids` - (Optional, List) Filter by a list of run instance IDs. If specified, only snapshots created by these protection runs will be returned.
* `run_start_from_time_usecs` - (Optional, Integer) Specifies the timestamp in Unix time epoch in microseconds to filter Object's snapshots which were run after this value.
* `run_start_to_time_usecs` - (Optional, Integer) Specifies the timestamp in Unix time epoch in microseconds to filter Object's snapshots which were run before this value.
* `run_types` - (Optional, List) Filter by run type. Only protection runs matching the specified types will be returned. By default, CDP hydration snapshots are not included unless explicitly queried using this field.
  * Constraints: Allowable list items are: `kRegular`, `kFull`, `kLog`, `kSystem`, `kHydrateCDP`, `kStorageArraySnapshot`.
* `snapshot_actions` - (Optional, List) Specifies a list of recovery actions. Only snapshots that apply to these actions will be returned.
  * Constraints: Allowable list items are: `RecoverVMs`, `RecoverFiles`, `InstantVolumeMount`, `RecoverVmDisks`, `MountVolumes`, `RecoverVApps`, `RecoverRDS`, `RecoverAurora`, `RecoverS3Buckets`, `RecoverApps`, `RecoverNasVolume`, `RecoverPhysicalVolumes`, `RecoverSystem`, `RecoverSanVolumes`, `RecoverNamespaces`, `RecoverObjects`, `DownloadFilesAndFolders`, `RecoverPublicFolders`, `RecoverVAppTemplates`, `RecoverMailbox`, `RecoverOneDrive`, `RecoverMsTeam`, `RecoverMsGroup`, `RecoverSharePoint`, `ConvertToPst`, `RecoverSfdcRecords`, `RecoverAzureSQL`, `DownloadChats`, `RecoverRDSPostgres`, `RecoverMailboxCSM`, `RecoverOneDriveCSM`, `RecoverSharePointCSM`.
* `to_time_usecs` - (Optional, Integer) Specifies the timestamp in Unix time epoch in microseconds to filter Object's snapshots which were taken before this value.
* `x_ibm_tenant_id` - (Required, String) Specifies the key to be used to encrypt the source credential. If includeSourceCredentials is set to true this key must be specified.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the Object Snapshots..
* `snapshots` - (List) Specifies the list of snapshots.
Nested schema for **snapshots**:
	* `aws_params` - (List) Specifies parameters of AWS type snapshots.
	Nested schema for **aws_params**:
		* `protection_type` - (String) Specifies the protection type of AWS snapshots.
		  * Constraints: Allowable values are: `kAgent`, `kNative`, `kSnapshotManager`, `kRDSSnapshotManager`, `kAuroraSnapshotManager`, `kAwsS3`, `kAwsRDSPostgresBackup`, `kAwsAuroraPostgres`, `kAwsRDSPostgres`.
	* `azure_params` - (List) Specifies parameters of Azure type snapshots.
	Nested schema for **azure_params**:
		* `protection_type` - (String) Specifies the protection type of Azure snapshots.
		  * Constraints: Allowable values are: `kAgent`, `kNative`, `kSnapshotManager`, `kAzureSQL`.
	* `cluster_id` - (Integer) Specifies the cluster id where this snapshot belongs to.
	* `cluster_incarnation_id` - (Integer) Specifies the cluster incarnation id where this snapshot belongs to.
	* `elastifile_params` - (List) Specifies the common parameters for NAS objects.
	Nested schema for **elastifile_params**:
		* `supported_nas_mount_protocols` - (List) Specifies a list of NAS mount protocols supported by this object.
		  * Constraints: Allowable list items are: `kNoProtocol`, `kNfs3`, `kNfs4_1`, `kCifs1`, `kCifs2`, `kCifs3`.
	* `environment` - (String) Specifies the snapshot environment.
	  * Constraints: Allowable values are: `kVMware`, `kHyperV`, `kAzure`, `kKVM`, `kAWS`, `kAcropolis`, `kGCP`, `kPhysical`, `kPhysicalFiles`, `kIsilon`, `kNetapp`, `kGenericNas`, `kFlashBlade`, `kElastifile`, `kGPFS`, `kPure`, `kIbmFlashSystem`, `kNimble`, `kSQL`, `kOracle`, `kExchange`, `kAD`, `kView`, `kO365`, `kHyperFlex`, `kKubernetes`, `kCassandra`, `kMongoDB`, `kCouchbase`, `kHdfs`, `kHive`, `kHBase`, `kSAPHANA`, `kUDA`, `kSfdc`.
	* `expiry_time_usecs` - (Integer) Specifies the expiry time of the snapshot in Unix timestamp epoch in microseconds. If the snapshot has no expiry, this property will not be set.
	* `external_target_info` - (List) Specifies archival target summary information.
	Nested schema for **external_target_info**:
		* `archival_task_id` - (String) Specifies the archival task id. This is a protection group UID which only applies when archival type is 'Tape'.
		* `ownership_context` - (String) Specifies the ownership context for the target.
		  * Constraints: Allowable values are: `Local`, `FortKnox`.
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
	* `flashblade_params` - (List) Specifies the common parameters for Flashblade objects.
	Nested schema for **flashblade_params**:
		* `supported_nas_mount_protocols` - (List) Specifies a list of NAS mount protocols supported by this object.
		  * Constraints: Allowable list items are: `kNfs`, `kCifs2`, `kHttp`.
	* `generic_nas_params` - (List) Specifies the common parameters for NAS objects.
	Nested schema for **generic_nas_params**:
		* `supported_nas_mount_protocols` - (List) Specifies a list of NAS mount protocols supported by this object.
		  * Constraints: Allowable list items are: `kNoProtocol`, `kNfs3`, `kNfs4_1`, `kCifs1`, `kCifs2`, `kCifs3`.
	* `gpfs_params` - (List) Specifies the common parameters for NAS objects.
	Nested schema for **gpfs_params**:
		* `supported_nas_mount_protocols` - (List) Specifies a list of NAS mount protocols supported by this object.
		  * Constraints: Allowable list items are: `kNoProtocol`, `kNfs3`, `kNfs4_1`, `kCifs1`, `kCifs2`, `kCifs3`.
	* `has_data_lock` - (Boolean) Specifies if this snapshot has datalock.
	* `hyperv_params` - (List) Specifies parameters of HyperV type snapshots.
	Nested schema for **hyperv_params**:
		* `protection_type` - (String) Specifies the protection type of HyperV snapshots.
		  * Constraints: Allowable values are: `kAuto`, `kRCT`, `kVSS`.
	* `id` - (String) Specifies the id of the snapshot.
	* `indexing_status` - (String) Specifies the indexing status of objects in this snapshot.<br> 'InProgress' indicates the indexing is in progress.<br> 'Done' indicates indexing is done.<br> 'NoIndex' indicates indexing is not applicable.<br> 'Error' indicates indexing failed with error.
	  * Constraints: Allowable values are: `InProgress`, `Done`, `NoIndex`, `Error`.
	* `isilon_params` - (List) Specifies the common parameters for Isilon objects.
	Nested schema for **isilon_params**:
		* `supported_nas_mount_protocols` - (List) Specifies a list of NAS mount protocols supported by this object.
		  * Constraints: Allowable list items are: `kNfs`, `kSmb`.
	* `netapp_params` - (List) Specifies the common parameters for Netapp objects.
	Nested schema for **netapp_params**:
		* `supported_nas_mount_protocols` - (List) Specifies a list of NAS mount protocols supported by this object.
		  * Constraints: Allowable list items are: `kNfs`, `kCifs`, `kIscsi`, `kFc`, `kFcache`, `kHttp`, `kNdmp`, `kManagement`, `kNvme`.
		* `volume_extended_style` - (String) Specifies the extended style of a NetApp volume.
		  * Constraints: Allowable values are: `kFlexVol`, `kFlexGroup`.
		* `volume_type` - (String) Specifies the Netapp volume type.
		  * Constraints: Allowable values are: `ReadWrite`, `LoadSharing`, `DataProtection`, `DataCache`, `Temp`, `UnkownType`.
	* `object_id` - (Integer) Specifies the object id which the snapshot is taken from.
	* `object_name` - (String) Specifies the object name which the snapshot is taken from.
	* `on_legal_hold` - (Boolean) Specifies if this snapshot is on legalhold.
	* `ownership_context` - (String) Specifies the ownership context for the target.
	  * Constraints: Allowable values are: `Local`, `FortKnox`.
	* `physical_params` - (List) Specifies parameters of Physical type snapshots.
	Nested schema for **physical_params**:
		* `enable_system_backup` - (Boolean) Specifies if system backup was enabled for the source in that particular run.
		* `protection_type` - (String) Specifies the protection type of Physical snapshots.
		  * Constraints: Allowable values are: `kFile`, `kVolume`.
	* `protection_group_id` - (String) Specifies id of the Protection Group.
	  * Constraints: The value must match regular expression `/^\\d+:\\d+:\\d+$/`.
	* `protection_group_name` - (String) Specifies name of the Protection Group.
	* `protection_group_run_id` - (String) Specifies id of the Protection Group Run.
	  * Constraints: The value must match regular expression `/^\\d+:\\d+$/`.
	* `region_id` - (String) Specifies the region id where this snapshot belongs to.
	* `run_instance_id` - (Integer) Specifies the instance id of the protection run which create the snapshot.
	* `run_start_time_usecs` - (Integer) Specifies the start time of the run in micro seconds.
	* `run_type` - (String) Specifies the type of protection run created this snapshot.
	  * Constraints: Allowable values are: `kRegular`, `kFull`, `kLog`, `kSystem`, `kHydrateCDP`, `kStorageArraySnapshot`.
	* `sfdc_params` - (List) Specifies the Salesforce objects mutation parameters.
	Nested schema for **sfdc_params**:
		* `records_added` - (Integer) Specifies the number of records added for the Object.
		* `records_modified` - (Integer) Specifies the number of records updated for the Object.
		* `records_removed` - (Integer) Specifies the number of records removed from the Object.
	* `snapshot_target_type` - (String) Specifies the target type where the Object's snapshot resides.
	  * Constraints: Allowable values are: `Local`, `Archival`, `RpaasArchival`, `StorageArraySnapshot`, `Remote`.
	* `snapshot_timestamp_usecs` - (Integer) Specifies the timestamp in Unix time epoch in microseconds when the snapshot is taken for the specified Object.
	* `source_group_id` - (String) Specifies the source protection group id in case of replication.
	* `source_id` - (Integer) Specifies the object source id which the snapshot is taken from.
	* `storage_domain_id` - (Integer) Specifies the Storage Domain id where the snapshot of object is present.

