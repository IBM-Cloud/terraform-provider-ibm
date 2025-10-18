---
layout: "ibm"
page_title: "IBM : ibm_backup_recoveries"
description: |-
  Get information about List of Recoveries.
subcategory: "IBM Backup Recovery API"
---

# ibm_backup_recoveries

Provides a read-only data source to retrieve information about a List of Recoveries.. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_backup_recoveries" "backup_recoveries" {
	ids = ["11:111:11"]
	x_ibm_tenant_id = ibm_backup_recovery.backup_recovery_instance.x_ibm_tenant_id
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `archival_target_type` - (Optional, List) Specifies the snapshot's archival target type from which recovery has been performed. This parameter applies only if 'snapshotTargetType' is 'Archival'.
  * Constraints: Allowable list items are: `Tape`, `Cloud`, `Nas`.
* `end_time_usecs` - (Optional, Integer) Returns the recoveries which are started before the specific time. This value should be in Unix timestamp epoch in microseconds.
* `ids` - (Optional, List) Filter Recoveries for given ids.
  * Constraints: The list items must match regular expression `/^\\d+:\\d+:\\d+$/`.
* `recovery_actions` - (Optional, List) Specifies the list of recovery actions to filter Recoveries. If empty, Recoveries related to all actions will be returned.
  * Constraints: Allowable list items are: `RecoverVMs`, `RecoverFiles`, `InstantVolumeMount`, `RecoverVmDisks`, `RecoverVApps`, `RecoverVAppTemplates`, `UptierSnapshot`, `RecoverRDS`, `RecoverAurora`, `RecoverS3Buckets`, `RecoverRDSPostgres`, `RecoverAzureSQL`, `RecoverApps`, `CloneApps`, `RecoverNasVolume`, `RecoverPhysicalVolumes`, `RecoverSystem`, `RecoverExchangeDbs`, `CloneAppView`, `RecoverSanVolumes`, `RecoverSanGroup`, `RecoverMailbox`, `RecoverOneDrive`, `RecoverSharePoint`, `RecoverPublicFolders`, `RecoverMsGroup`, `RecoverMsTeam`, `ConvertToPst`, `DownloadChats`, `RecoverMailboxCSM`, `RecoverOneDriveCSM`, `RecoverSharePointCSM`, `RecoverNamespaces`, `RecoverObjects`, `RecoverSfdcObjects`, `RecoverSfdcOrg`, `RecoverSfdcRecords`, `DownloadFilesAndFolders`, `CloneVMs`, `CloneView`, `CloneRefreshApp`, `CloneVMsToView`, `ConvertAndDeployVMs`, `DeployVMs`.
* `return_only_child_recoveries` - (Optional, Boolean) Returns only child recoveries if passed as true. This filter should always be used along with 'ids' filter.
* `snapshot_environments` - (Optional, List) Specifies the list of snapshot environment types to filter Recoveries. If empty, Recoveries related to all environments will be returned.
  * Constraints: Allowable list items are: `kPhysical`, `kSQL`, `kKubernetes`.
* `snapshot_target_type` - (Optional, List) Specifies the snapshot's target type from which recovery has been performed.
  * Constraints: Allowable list items are: `Local`, `Archival`, `RpaasArchival`, `StorageArraySnapshot`, `Remote`.
* `start_time_usecs` - (Optional, Integer) Returns the recoveries which are started after the specific time. This value should be in Unix timestamp epoch in microseconds.
* `status` - (Optional, List) Specifies the list of run status to filter Recoveries. If empty, Recoveries with all run status will be returned.
  * Constraints: Allowable list items are: `Accepted`, `Running`, `Canceled`, `Canceling`, `Failed`, `Missed`, `Succeeded`, `SucceededWithWarning`, `OnHold`, `Finalizing`, `Skipped`, `LegalHold`.
* `x_ibm_tenant_id` - (Required, String) Specifies the key to be used to encrypt the source credential. If includeSourceCredentials is set to true this key must be specified.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the List of Recoveries..
* `recoveries` - (List) Specifies list of Recoveries.
Nested schema for **recoveries**:
	* `can_tear_down` - (Boolean) Specifies whether it's possible to tear down the objects created by the recovery.
	* `creation_info` - (List) Specifies the information about the creation of the protection group or recovery.
	Nested schema for **creation_info**:
		* `user_name` - (String) Specifies the name of the user who created the protection group or recovery.
	* `end_time_usecs` - (Integer) Specifies the end time of the Recovery in Unix timestamp epoch in microseconds. This field will be populated only after Recovery is finished.
	* `id` - (String) Specifies the id of the Recovery.
	  * Constraints: The value must match regular expression `/^\\d+:\\d+:\\d+$/`.
	* `is_multi_stage_restore` - (Boolean) Specifies whether the current recovery operation is a multi-stage restore operation. This is currently used by VMware recoveres for the migration/hot-standby use case.
	* `is_parent_recovery` - (Boolean) Specifies whether the current recovery operation has created child recoveries. This is currently used in SQL recovery where multiple child recoveries can be tracked under a common/parent recovery.
	* `kubernetes_params` - (List) Specifies the recovery options specific to Kubernetes environment.
	Nested schema for **kubernetes_params**:
		* `download_file_and_folder_params` - (List) Specifies the parameters to download files and folders.
		Nested schema for **download_file_and_folder_params**:
			* `download_file_path` - (String) Specifies the path location to download the files and folders.
			* `expiry_time_usecs` - (Integer) Specifies the time upto which the download link is available.
			* `files_and_folders` - (List) Specifies the info about the files and folders to be recovered.
			Nested schema for **files_and_folders**:
				* `absolute_path` - (String) Specifies the absolute path to the file or folder.
				* `destination_dir` - (String) Specifies the destination directory where the file/directory was copied.
				* `is_directory` - (Boolean) Specifies whether this is a directory or not.
				* `is_view_file_recovery` - (Boolean) Specify if the recovery is of type view file/folder.
				* `messages` - (List) Specify error messages about the file during recovery.
				* `status` - (String) Specifies the recovery status for this file or folder.
				  * Constraints: Allowable values are: `NotStarted`, `EstimationInProgress`, `EstimationDone`, `CopyInProgress`, `Finished`.
		* `objects` - (List) Specifies the list of objects which need to be recovered.
		Nested schema for **objects**:
			* `archival_target_info` - (List) Specifies the archival target information if the snapshot is an archival snapshot.
			Nested schema for **archival_target_info**:
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
			* `bytes_restored` - (Integer) Specify the total bytes restored.
			* `end_time_usecs` - (Integer) Specifies the end time of the Recovery in Unix timestamp epoch in microseconds. This field will be populated only after Recovery is finished.
			* `messages` - (List) Specify error messages about the object.
			* `object_info` - (List) Specifies the information about the object for which the snapshot is taken.
			Nested schema for **object_info**:
				* `child_objects` - (List) Specifies child object details.
				Nested schema for **child_objects**:
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
			* `point_in_time_usecs` - (Integer) Specifies the timestamp (in microseconds. from epoch) for recovering to a point-in-time in the past.
			* `progress_task_id` - (String) Progress monitor task id for Recovery of VM.
			* `protection_group_id` - (String) Specifies the protection group id of the object snapshot.
			* `protection_group_name` - (String) Specifies the protection group name of the object snapshot.
			* `recover_from_standby` - (Boolean) Specifies that user wants to perform standby restore if it is enabled for this object.
			* `snapshot_creation_time_usecs` - (Integer) Specifies the time when the snapshot is created in Unix timestamp epoch in microseconds.
			* `snapshot_id` - (String) Specifies the snapshot id.
			* `snapshot_target_type` - (String) Specifies the snapshot target type.
			  * Constraints: Allowable values are: `Local`, `Archival`, `RpaasArchival`, `StorageArraySnapshot`, `Remote`.
			* `start_time_usecs` - (Integer) Specifies the start time of the Recovery in Unix timestamp epoch in microseconds.
			* `status` - (String) Status of the Recovery. 'Running' indicates that the Recovery is still running. 'Canceled' indicates that the Recovery has been cancelled. 'Canceling' indicates that the Recovery is in the process of being cancelled. 'Failed' indicates that the Recovery has failed. 'Succeeded' indicates that the Recovery has finished successfully. 'SucceededWithWarning' indicates that the Recovery finished successfully, but there were some warning messages. 'Skipped' indicates that the Recovery task was skipped.
			  * Constraints: Allowable values are: `Accepted`, `Running`, `Canceled`, `Canceling`, `Failed`, `Missed`, `Succeeded`, `SucceededWithWarning`, `OnHold`, `Finalizing`, `Skipped`, `LegalHold`.
		* `recover_file_and_folder_params` - (List) Specifies the parameters to perform a file and folder recovery.
		Nested schema for **recover_file_and_folder_params**:
			* `files_and_folders` - (List) Specifies the information about the files and folders to be recovered.
			Nested schema for **files_and_folders**:
				* `absolute_path` - (String) Specifies the absolute path to the file or folder.
				* `destination_dir` - (String) Specifies the destination directory where the file/directory was copied.
				* `is_directory` - (Boolean) Specifies whether this is a directory or not.
				* `is_view_file_recovery` - (Boolean) Specify if the recovery is of type view file/folder.
				* `messages` - (List) Specify error messages about the file during recovery.
				* `status` - (String) Specifies the recovery status for this file or folder.
				  * Constraints: Allowable values are: `NotStarted`, `EstimationInProgress`, `EstimationDone`, `CopyInProgress`, `Finished`.
			* `kubernetes_target_params` - (List) Specifies the parameters to recover to a Kubernetes target.
			Nested schema for **kubernetes_target_params**:
				* `continue_on_error` - (Boolean) Specifies whether to continue recovering other files if one of files or folders failed to recover. Default value is false.
				* `new_target_config` - (List) Specifies the configuration for recovering to a new target.
				Nested schema for **new_target_config**:
					* `absolute_path` - (String) Specifies the absolute path of the file.
					* `target_namespace` - (List) Specifies the target namespace to recover files and folders to.
					Nested schema for **target_namespace**:
						* `id` - (Integer) Specifies the id of the object.
						* `name` - (String) Specifies the name of the object.
						* `parent_source_id` - (Integer) Specifies the id of the parent source of the target.
						* `parent_source_name` - (String) Specifies the name of the parent source of the target.
					* `target_pvc` - (List) Specifies the target PVC(Persistent Volume Claim) to recover files and folders to.
					Nested schema for **target_pvc**:
						* `id` - (Integer) Specifies the id of the object.
						* `name` - (String) Specifies the name of the object.
						* `parent_source_id` - (Integer) Specifies the id of the parent source of the target.
						* `parent_source_name` - (String) Specifies the name of the parent source of the target.
					* `target_source` - (List) Specifies the target kubernetes to recover files and folders to.
					Nested schema for **target_source**:
						* `id` - (Integer) Specifies the id of the object.
						* `name` - (String) Specifies the name of the object.
						* `parent_source_id` - (Integer) Specifies the id of the parent source of the target.
						* `parent_source_name` - (String) Specifies the name of the parent source of the target.
				* `original_target_config` - (List) Specifies the configuration for recovering to the original target.
				Nested schema for **original_target_config**:
					* `alternate_path` - (String) Specifies the alternate path location to recover files to.
					* `recover_to_original_path` - (Boolean) Specifies whether to recover files and folders to the original path location. If false, alternatePath must be specified.
				* `overwrite_existing` - (Boolean) Specifies whether to overwrite the existing files. Default is true.
				* `preserve_attributes` - (Boolean) Specifies whether to preserve original attributes. Default is true.
				* `recover_to_original_target` - (Boolean) Specifies whether to recover to the original target. If true, originalTargetConfig must be specified. If false, newTargetConfig must be specified.
				* `vlan_config` - (List) Specifies VLAN Params associated with the recovered files and folders. If this is not specified, then the VLAN settings will be automatically selected from one of the below options: a. If VLANs are configured on Cohesity, then the VLAN host/VIP will be automatically based on the client's (e.g. ESXI host) IP address. b. If VLANs are not configured on Cohesity, then the partition hostname or VIPs will be used for Recovery.
				Nested schema for **vlan_config**:
					* `disable_vlan` - (Boolean) If this is set to true, then even if VLANs are configured on the system, the partition VIPs will be used for the Recovery.
					* `id` - (Integer) If this is set, then the Cohesity host name or the IP address associated with this vlan is used for mounting Cohesity's view on the remote host.
					* `interface_name` - (String) Interface group to use for Recovery.
			* `target_environment` - (String) Specifies the environment of the recovery target. The corresponding params below must be filled out.
			  * Constraints: Allowable values are: `kKubernetes`.
		* `recover_namespace_params` - (List) Specifies the parameters to recover Kubernetes Namespaces.
		Nested schema for **recover_namespace_params**:
			* `kubernetes_target_params` - (List) Specifies the params for recovering to a Kubernetes host.
			Nested schema for **kubernetes_target_params**:
				* `exclude_params` - (List) Specifies the parameters to in/exclude objects (e.g.: volumes). An object satisfying any of these criteria will be included by this filter.
				Nested schema for **exclude_params**:
					* `label_combination_method` - (String) Whether to include all the labels or any of them while performing inclusion/exclusion of objects.
					  * Constraints: Allowable values are: `AND`, `OR`.
					* `label_vector` - (List) Array of Object to represent Label that Specify Objects (e.g.: Persistent Volumes and Persistent Volume Claims) to Include or Exclude.It will be a two-dimensional array, where each inner array will consist of a key and value representing labels. Using this two dimensional array of Labels, the Cluster generates a list of items to include in the filter, which are derived from intersections or the union of these labels, as decided by operation parameter.
					Nested schema for **label_vector**:
						* `key` - (String) The key of the label, used to identify the label.
						* `value` - (String) The value associated with the label key.
					* `objects` - (List) Array of objects that are to be included.
				* `excluded_pvcs` - (List) Specifies the list of pvc to be excluded from recovery.
				Nested schema for **excluded_pvcs**:
					* `id` - (Integer) Specifies the id of the pvc.
					* `name` - (String) Name of the pvc.
				* `include_params` - (List) Specifies the parameters to in/exclude objects (e.g.: volumes). An object satisfying any of these criteria will be included by this filter.
				Nested schema for **include_params**:
					* `label_combination_method` - (String) Whether to include all the labels or any of them while performing inclusion/exclusion of objects.
					  * Constraints: Allowable values are: `AND`, `OR`.
					* `label_vector` - (List) Array of Object to represent Label that Specify Objects (e.g.: Persistent Volumes and Persistent Volume Claims) to Include or Exclude.It will be a two-dimensional array, where each inner array will consist of a key and value representing labels. Using this two dimensional array of Labels, the Cluster generates a list of items to include in the filter, which are derived from intersections or the union of these labels, as decided by operation parameter.
					Nested schema for **label_vector**:
						* `key` - (String) The key of the label, used to identify the label.
						* `value` - (String) The value associated with the label key.
					* `objects` - (List) Array of objects that are to be included.
				* `objects` - (List) Specifies the objects to be recovered.
				Nested schema for **objects**:
					* `archival_target_info` - (List) Specifies the archival target information if the snapshot is an archival snapshot.
					Nested schema for **archival_target_info**:
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
					* `bytes_restored` - (Integer) Specify the total bytes restored.
					* `end_time_usecs` - (Integer) Specifies the end time of the Recovery in Unix timestamp epoch in microseconds. This field will be populated only after Recovery is finished.
					* `messages` - (List) Specify error messages about the object.
					* `object_info` - (List) Specifies the information about the object for which the snapshot is taken.
					Nested schema for **object_info**:
						* `child_objects` - (List) Specifies child object details.
						Nested schema for **child_objects**:
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
					* `point_in_time_usecs` - (Integer) Specifies the timestamp (in microseconds. from epoch) for recovering to a point-in-time in the past.
					* `progress_task_id` - (String) Progress monitor task id for Recovery of VM.
					* `protection_group_id` - (String) Specifies the protection group id of the object snapshot.
					* `protection_group_name` - (String) Specifies the protection group name of the object snapshot.
					* `recover_from_standby` - (Boolean) Specifies that user wants to perform standby restore if it is enabled for this object.
					* `snapshot_creation_time_usecs` - (Integer) Specifies the time when the snapshot is created in Unix timestamp epoch in microseconds.
					* `snapshot_id` - (String) Specifies the snapshot id.
					* `snapshot_target_type` - (String) Specifies the snapshot target type.
					  * Constraints: Allowable values are: `Local`, `Archival`, `RpaasArchival`, `StorageArraySnapshot`, `Remote`.
					* `start_time_usecs` - (Integer) Specifies the start time of the Recovery in Unix timestamp epoch in microseconds.
					* `status` - (String) Status of the Recovery. 'Running' indicates that the Recovery is still running. 'Canceled' indicates that the Recovery has been cancelled. 'Canceling' indicates that the Recovery is in the process of being cancelled. 'Failed' indicates that the Recovery has failed. 'Succeeded' indicates that the Recovery has finished successfully. 'SucceededWithWarning' indicates that the Recovery finished successfully, but there were some warning messages. 'Skipped' indicates that the Recovery task was skipped.
					  * Constraints: Allowable values are: `Accepted`, `Running`, `Canceled`, `Canceling`, `Failed`, `Missed`, `Succeeded`, `SucceededWithWarning`, `OnHold`, `Finalizing`, `Skipped`, `LegalHold`.
				* `recover_protection_group_runs_params` - (List) Specifies the Protection Group Runs params to recover. All the VM's that are successfully backed up by specified Runs will be recovered. This can be specified along with individual snapshots of VMs. User has to make sure that specified Object snapshots and Protection Group Runs should not have any intersection. For example, user cannot specify multiple Runs which has same Object or an Object snapshot and a Run which has same Object's snapshot.
				Nested schema for **recover_protection_group_runs_params**:
					* `archival_target_id` - (Integer) Specifies the archival target id. If specified and Protection Group run has an archival snapshot then VMs are recovered from the specified archival snapshot. If not specified (default), VMs are recovered from local snapshot.
					* `protection_group_id` - (String) Specifies the local Protection Group id. In case of recovering a replication Run, this field should be provided with local Protection Group id.
					* `protection_group_instance_id` - (Integer) Specifies the Protection Group Instance id.
					* `protection_group_run_id` - (String) Specifies the Protection Group Run id from which to recover VMs. All the VM's that are successfully protected by this Run will be recovered.
					  * Constraints: The value must match regular expression `/^\\d+:\\d+$/`.
				* `recover_pvcs_only` - (Boolean) Specifies whether to recover PVCs only during recovery.
				* `recovery_target_config` - (List) Specifies the recovery target configuration of the Namespace recovery.
				Nested schema for **recovery_target_config**:
					* `new_source_config` - (List) Specifies the new source configuration if a Kubernetes Namespace is being restored to a different source than the one from which it was protected.
					Nested schema for **new_source_config**:
						* `source` - (List) Specifies the id of the parent source to recover the Namespaces.
						Nested schema for **source**:
							* `id` - (Integer) Specifies the id of the object.
							* `name` - (String) Specifies the name of the object.
					* `recover_to_new_source` - (Boolean) Specifies whether or not to recover the Namespaces to a different source than they were backed up from.
				* `rename_recovered_namespaces_params` - (List) Specifies params to rename the Namespaces that are recovered. If not specified, the original names of the Namespaces are preserved. If a name collision occurs then the Namespace being recovered will overwrite the Namespace already present on the source.
				Nested schema for **rename_recovered_namespaces_params**:
					* `prefix` - (String) Specifies the prefix string to be added to recovered or cloned object names.
					* `suffix` - (String) Specifies the suffix string to be added to recovered or cloned object names.
				* `skip_cluster_compatibility_check` - (Boolean) Specifies whether to skip checking if the target cluster, to restore to, is compatible or not. By default restore allowed to compatible cluster only.
				* `storage_class` - (List) Specifies the storage class parameters for recovery of namespace.
				Nested schema for **storage_class**:
					* `storage_class_mapping` - (List) Specifies mapping of storage classes.
					Nested schema for **storage_class_mapping**:
						* `key` - (String) The key of the label, used to identify the label.
						* `value` - (String) The value associated with the label key.
					* `use_storage_class_mapping` - (Boolean) Specifies whether or not to use storage class mapping.
			* `target_environment` - (String) Specifies the environment of the recovery target. The corresponding params below must be filled out. As of now only kubernetes target environment is supported.
			  * Constraints: Allowable values are: `kKubernetes`.
			* `vlan_config` - (List) Specifies VLAN Params associated with the recovered. If this is not specified, then the VLAN settings will be automatically selected from one of the below options: a. If VLANs are configured on Cohesity, then the VLAN host/VIP will be automatically based on the client's (e.g. ESXI host) IP address. b. If VLANs are not configured on Cohesity, then the partition hostname or VIPs will be used for Recovery.
			Nested schema for **vlan_config**:
				* `disable_vlan` - (Boolean) If this is set to true, then even if VLANs are configured on the system, the partition VIPs will be used for the Recovery.
				* `id` - (Integer) If this is set, then the Cohesity host name or the IP address associated with this vlan is used for mounting Cohesity's view on the remote host.
				* `interface_name` - (String) Interface group to use for Recovery.
		* `recovery_action` - (String) Specifies the type of recover action to be performed.
		  * Constraints: Allowable values are: `RecoverNamespaces`, `RecoverFiles`, `DownloadFilesAndFolders`.
	* `messages` - (List) Specifies messages about the recovery.
	* `mssql_params` - (List) Specifies the recovery options specific to Sql environment.
	Nested schema for **mssql_params**:
		* `recover_app_params` - (List) Specifies the parameters to recover Sql databases.
		  * Constraints: The minimum length is `1` item.
		Nested schema for **recover_app_params**:
			* `aag_info` - (List) Object details for Mssql.
			Nested schema for **aag_info**:
				* `name` - (String) Specifies the AAG name.
				* `object_id` - (Integer) Specifies the AAG object Id.
			* `archival_target_info` - (List) Specifies the archival target information if the snapshot is an archival snapshot.
			Nested schema for **archival_target_info**:
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
			* `bytes_restored` - (Integer) Specify the total bytes restored.
			* `end_time_usecs` - (Integer) Specifies the end time of the Recovery in Unix timestamp epoch in microseconds. This field will be populated only after Recovery is finished.
			* `host_info` - (List) Specifies the host information for a objects. This is mainly populated in case of App objects where app object is hosted by another object such as VM or physical server.
			Nested schema for **host_info**:
				* `environment` - (String) Specifies the environment of the object.
				  * Constraints: Allowable values are: `kPhysical`, `kSQL`.
				* `id` - (String) Specifies the id of the host object.
				* `name` - (String) Specifies the name of the host object.
			* `is_encrypted` - (Boolean) Specifies whether the database is TDE enabled.
			* `messages` - (List) Specify error messages about the object.
			* `object_info` - (List) Specifies the information about the object for which the snapshot is taken.
			Nested schema for **object_info**:
				* `child_objects` - (List) Specifies child object details.
				Nested schema for **child_objects**:
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
			* `point_in_time_usecs` - (Integer) Specifies the timestamp (in microseconds. from epoch) for recovering to a point-in-time in the past.
			* `progress_task_id` - (String) Progress monitor task id for Recovery of VM.
			* `protection_group_id` - (String) Specifies the protection group id of the object snapshot.
			* `protection_group_name` - (String) Specifies the protection group name of the object snapshot.
			* `recover_from_standby` - (Boolean) Specifies that user wants to perform standby restore if it is enabled for this object.
			* `snapshot_creation_time_usecs` - (Integer) Specifies the time when the snapshot is created in Unix timestamp epoch in microseconds.
			* `snapshot_id` - (String) Specifies the snapshot id.
			* `snapshot_target_type` - (String) Specifies the snapshot target type.
			  * Constraints: Allowable values are: `Local`, `Archival`, `RpaasArchival`, `StorageArraySnapshot`, `Remote`.
			* `sql_target_params` - (List) Specifies the params for recovering to a sql host. Specifiy seperate settings for each db object that need to be recovered. Provided sql backup should be recovered to same type of target host. For Example: If you have sql backup taken from a physical host then that should be recovered to physical host only.
			Nested schema for **sql_target_params**:
				* `new_source_config` - (List) Specifies the destination Source configuration parameters where the databases will be recovered. This is mandatory if recoverToNewSource is set to true.
				Nested schema for **new_source_config**:
					* `data_file_directory_location` - (String) Specifies the directory where to put the database data files. Missing directory will be automatically created.
					* `database_name` - (String) Specifies a new name for the restored database. If this field is not specified, then the original database will be overwritten after recovery.
					* `host` - (List) Specifies the source id of target host where databases will be recovered. This source id can be a physical host or virtual machine.
					Nested schema for **host**:
						* `id` - (Integer) Specifies the id of the object.
						* `name` - (String) Specifies the name of the object.
					* `instance_name` - (String) Specifies an instance name of the Sql Server that should be used for restoring databases to.
					* `keep_cdc` - (Boolean) Specifies whether to keep CDC (Change Data Capture) on recovered databases or not. If not passed, this is assumed to be true. If withNoRecovery is passed as true, then this field must not be set to true. Passing this field as true in this scenario will be a invalid request.
					* `log_file_directory_location` - (String) Specifies the directory where to put the database log files. Missing directory will be automatically created.
					* `multi_stage_restore_options` - (List) Specifies the parameters related to multi stage Sql restore.
					Nested schema for **multi_stage_restore_options**:
						* `enable_auto_sync` - (Boolean) Set this to true if you want to enable auto sync for multi stage restore.
						* `enable_multi_stage_restore` - (Boolean) Set this to true if you are creating a multi-stage Sql restore task needed for features such as Hot-Standby.
					* `native_log_recovery_with_clause` - (String) Specifies the WITH clause to be used in native sql log restore command. This is only applicable for native log restore.
					* `native_recovery_with_clause` - (String) 'with_clause' contains 'with clause' to be used in native sql restore command. This is only applicable for database restore of native sql backup. Here user can specify multiple restore options. Example: 'WITH BUFFERCOUNT = 575, MAXTRANSFERSIZE = 2097152'.
					* `overwriting_policy` - (String) Specifies a policy to be used while recovering existing databases.
					  * Constraints: Allowable values are: `FailIfExists`, `Overwrite`.
					* `replay_entire_last_log` - (Boolean) Specifies the option to set replay last log bit while creating the sql restore task and doing restore to latest point-in-time. If this is set to true, we will replay the entire last log without STOPAT.
					* `restore_time_usecs` - (Integer) Specifies the time in the past to which the Sql database needs to be restored. This allows for granular recovery of Sql databases. If this is not set, the Sql database will be restored from the full/incremental snapshot.
					* `secondary_data_files_dir_list` - (List) Specifies the secondary data filename pattern and corresponding direcories of the DB. Secondary data files are optional and are user defined. The recommended file extention for secondary files is ".ndf". If this option is specified and the destination folders do not exist they will be automatically created.
					Nested schema for **secondary_data_files_dir_list**:
						* `directory` - (String) Specifies the directory where to keep the files matching the pattern.
						* `filename_pattern` - (String) Specifies a pattern to be matched with filenames. This can be a regex expression.
					* `with_no_recovery` - (Boolean) Specifies the flag to bring DBs online or not after successful recovery. If this is passed as true, then it means DBs won't be brought online.
				* `original_source_config` - (List) Specifies the Source configuration if databases are being recovered to Original Source. If not specified, all the configuration parameters will be retained.
				Nested schema for **original_source_config**:
					* `capture_tail_logs` - (Boolean) Set this to true if tail logs are to be captured before the recovery operation. This is only applicable if database is not being renamed.
					* `data_file_directory_location` - (String) Specifies the directory where to put the database data files. Missing directory will be automatically created. If you are overwriting the existing database then this field will be ignored.
					* `keep_cdc` - (Boolean) Specifies whether to keep CDC (Change Data Capture) on recovered databases or not. If not passed, this is assumed to be true. If withNoRecovery is passed as true, then this field must not be set to true. Passing this field as true in this scenario will be a invalid request.
					* `log_file_directory_location` - (String) Specifies the directory where to put the database log files. Missing directory will be automatically created. If you are overwriting the existing database then this field will be ignored.
					* `multi_stage_restore_options` - (List) Specifies the parameters related to multi stage Sql restore.
					Nested schema for **multi_stage_restore_options**:
						* `enable_auto_sync` - (Boolean) Set this to true if you want to enable auto sync for multi stage restore.
						* `enable_multi_stage_restore` - (Boolean) Set this to true if you are creating a multi-stage Sql restore task needed for features such as Hot-Standby.
					* `native_log_recovery_with_clause` - (String) Specifies the WITH clause to be used in native sql log restore command. This is only applicable for native log restore.
					* `native_recovery_with_clause` - (String) 'with_clause' contains 'with clause' to be used in native sql restore command. This is only applicable for database restore of native sql backup. Here user can specify multiple restore options. Example: 'WITH BUFFERCOUNT = 575, MAXTRANSFERSIZE = 2097152'.
					* `new_database_name` - (String) Specifies a new name for the restored database. If this field is not specified, then the original database will be overwritten after recovery.
					* `overwriting_policy` - (String) Specifies a policy to be used while recovering existing databases.
					  * Constraints: Allowable values are: `FailIfExists`, `Overwrite`.
					* `replay_entire_last_log` - (Boolean) Specifies the option to set replay last log bit while creating the sql restore task and doing restore to latest point-in-time. If this is set to true, we will replay the entire last log without STOPAT.
					* `restore_time_usecs` - (Integer) Specifies the time in the past to which the Sql database needs to be restored. This allows for granular recovery of Sql databases. If this is not set, the Sql database will be restored from the full/incremental snapshot.
					* `secondary_data_files_dir_list` - (List) Specifies the secondary data filename pattern and corresponding direcories of the DB. Secondary data files are optional and are user defined. The recommended file extention for secondary files is ".ndf". If this option is specified and the destination folders do not exist they will be automatically created.
					Nested schema for **secondary_data_files_dir_list**:
						* `directory` - (String) Specifies the directory where to keep the files matching the pattern.
						* `filename_pattern` - (String) Specifies a pattern to be matched with filenames. This can be a regex expression.
					* `with_no_recovery` - (Boolean) Specifies the flag to bring DBs online or not after successful recovery. If this is passed as true, then it means DBs won't be brought online.
				* `recover_to_new_source` - (Boolean) Specifies the parameter whether the recovery should be performed to a new sources or an original Source Target.
			* `start_time_usecs` - (Integer) Specifies the start time of the Recovery in Unix timestamp epoch in microseconds.
			* `status` - (String) Status of the Recovery. 'Running' indicates that the Recovery is still running. 'Canceled' indicates that the Recovery has been cancelled. 'Canceling' indicates that the Recovery is in the process of being cancelled. 'Failed' indicates that the Recovery has failed. 'Succeeded' indicates that the Recovery has finished successfully. 'SucceededWithWarning' indicates that the Recovery finished successfully, but there were some warning messages. 'Skipped' indicates that the Recovery task was skipped.
			  * Constraints: Allowable values are: `Accepted`, `Running`, `Canceled`, `Canceling`, `Failed`, `Missed`, `Succeeded`, `SucceededWithWarning`, `OnHold`, `Finalizing`, `Skipped`, `LegalHold`.
			* `target_environment` - (String) Specifies the environment of the recovery target. The corresponding params below must be filled out.
			  * Constraints: Allowable values are: `kSQL`.
		* `recovery_action` - (String) Specifies the type of recover action to be performed.
		  * Constraints: Allowable values are: `RecoverApps`, `CloneApps`.
		* `vlan_config` - (List) Specifies VLAN Params associated with the recovered. If this is not specified, then the VLAN settings will be automatically selected from one of the below options: a. If VLANs are configured on IBM, then the VLAN host/VIP will be automatically based on the client's (e.g. ESXI host) IP address. b. If VLANs are not configured on IBM, then the partition hostname or VIPs will be used for Recovery.
		Nested schema for **vlan_config**:
			* `disable_vlan` - (Boolean) If this is set to true, then even if VLANs are configured on the system, the partition VIPs will be used for the Recovery.
			* `id` - (Integer) If this is set, then the Cohesity host name or the IP address associated with this vlan is used for mounting Cohesity's view on the remote host.
			* `interface_name` - (String) Interface group to use for Recovery.
	* `name` - (String) Specifies the name of the Recovery.
	* `parent_recovery_id` - (String) If current recovery is child recovery triggered by another parent recovery operation, then this field willt specify the id of the parent recovery.
	  * Constraints: The value must match regular expression `/^\\d+:\\d+:\\d+$/`.
	* `permissions` - (List) Specifies the list of tenants that have permissions for this recovery.
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
	* `physical_params` - (List) Specifies the recovery options specific to Physical environment.
	Nested schema for **physical_params**:
		* `download_file_and_folder_params` - (List) Specifies the parameters to download files and folders.
		Nested schema for **download_file_and_folder_params**:
			* `download_file_path` - (String) Specifies the path location to download the files and folders.
			* `expiry_time_usecs` - (Integer) Specifies the time upto which the download link is available.
			* `files_and_folders` - (List) Specifies the info about the files and folders to be recovered.
			Nested schema for **files_and_folders**:
				* `absolute_path` - (String) Specifies the absolute path to the file or folder.
				* `destination_dir` - (String) Specifies the destination directory where the file/directory was copied.
				* `is_directory` - (Boolean) Specifies whether this is a directory or not.
				* `is_view_file_recovery` - (Boolean) Specify if the recovery is of type view file/folder.
				* `messages` - (List) Specify error messages about the file during recovery.
				* `status` - (String) Specifies the recovery status for this file or folder.
				  * Constraints: Allowable values are: `NotStarted`, `EstimationInProgress`, `EstimationDone`, `CopyInProgress`, `Finished`.
		* `mount_volume_params` - (List) Specifies the parameters to mount Physical Volumes.
		Nested schema for **mount_volume_params**:
			* `physical_target_params` - (List) Specifies the params for recovering to a physical target.
			Nested schema for **physical_target_params**:
				* `mount_to_original_target` - (Boolean) Specifies whether to mount to the original target. If true, originalTargetConfig must be specified. If false, newTargetConfig must be specified.
				* `mounted_volume_mapping` - (List) Specifies the mapping of original volumes and mounted volumes.
				Nested schema for **mounted_volume_mapping**:
					* `file_system_type` - (String) Specifies the type of the file system of the volume.
					* `mounted_volume` - (String) Specifies the name of the point where the volume is mounted.
					* `original_volume` - (String) Specifies the name of the original volume.
				* `new_target_config` - (List) Specifies the configuration for mounting to a new target.
				Nested schema for **new_target_config**:
					* `mount_target` - (List) Specifies the target entity to recover to.
					Nested schema for **mount_target**:
						* `id` - (Integer) Specifies the id of the object.
						* `name` - (String) Specifies the name of the object.
						* `parent_source_id` - (Integer) Specifies the id of the parent source of the target.
						* `parent_source_name` - (String) Specifies the name of the parent source of the target.
					* `server_credentials` - (List) Specifies credentials to access the target server. This is required if the server is of Linux OS.
					Nested schema for **server_credentials**:
						* `password` - (String) Specifies the password to access target entity.
						* `username` - (String) Specifies the username to access target entity.
				* `original_target_config` - (List) Specifies the configuration for mounting to the original target.
				Nested schema for **original_target_config**:
					* `server_credentials` - (List) Specifies credentials to access the target server. This is required if the server is of Linux OS.
					Nested schema for **server_credentials**:
						* `password` - (String) Specifies the password to access target entity.
						* `username` - (String) Specifies the username to access target entity.
				* `read_only_mount` - (Boolean) Specifies whether to perform a read-only mount. Default is false.
				* `vlan_config` - (List) Specifies VLAN Params associated with the recovered. If this is not specified, then the VLAN settings will be automatically selected from one of the below options: a. If VLANs are configured on Cohesity, then the VLAN host/VIP will be automatically based on the client's (e.g. ESXI host) IP address. b. If VLANs are not configured on Cohesity, then the partition hostname or VIPs will be used for Recovery.
				Nested schema for **vlan_config**:
					* `disable_vlan` - (Boolean) If this is set to true, then even if VLANs are configured on the system, the partition VIPs will be used for the Recovery.
					* `id` - (Integer) If this is set, then the Cohesity host name or the IP address associated with this vlan is used for mounting Cohesity's view on the remote host.
					* `interface_name` - (String) Interface group to use for Recovery.
				* `volume_names` - (List) Specifies the names of volumes that need to be mounted. If this is not specified then all volumes that are part of the source VM will be mounted on the target VM.
			* `target_environment` - (String) Specifies the environment of the recovery target. The corresponding params below must be filled out.
			  * Constraints: Allowable values are: `kPhysical`.
		* `objects` - (List) Specifies the list of Recover Object parameters. For recovering files, specifies the object contains the file to recover.
		Nested schema for **objects**:
			* `archival_target_info` - (List) Specifies the archival target information if the snapshot is an archival snapshot.
			Nested schema for **archival_target_info**:
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
			* `bytes_restored` - (Integer) Specify the total bytes restored.
			* `end_time_usecs` - (Integer) Specifies the end time of the Recovery in Unix timestamp epoch in microseconds. This field will be populated only after Recovery is finished.
			* `messages` - (List) Specify error messages about the object.
			* `object_info` - (List) Specifies the information about the object for which the snapshot is taken.
			Nested schema for **object_info**:
				* `child_objects` - (List) Specifies child object details.
				Nested schema for **child_objects**:
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
			* `point_in_time_usecs` - (Integer) Specifies the timestamp (in microseconds. from epoch) for recovering to a point-in-time in the past.
			* `progress_task_id` - (String) Progress monitor task id for Recovery of VM.
			* `protection_group_id` - (String) Specifies the protection group id of the object snapshot.
			* `protection_group_name` - (String) Specifies the protection group name of the object snapshot.
			* `recover_from_standby` - (Boolean) Specifies that user wants to perform standby restore if it is enabled for this object.
			* `snapshot_creation_time_usecs` - (Integer) Specifies the time when the snapshot is created in Unix timestamp epoch in microseconds.
			* `snapshot_id` - (String) Specifies the snapshot id.
			* `snapshot_target_type` - (String) Specifies the snapshot target type.
			  * Constraints: Allowable values are: `Local`, `Archival`, `RpaasArchival`, `StorageArraySnapshot`, `Remote`.
			* `start_time_usecs` - (Integer) Specifies the start time of the Recovery in Unix timestamp epoch in microseconds.
			* `status` - (String) Status of the Recovery. 'Running' indicates that the Recovery is still running. 'Canceled' indicates that the Recovery has been cancelled. 'Canceling' indicates that the Recovery is in the process of being cancelled. 'Failed' indicates that the Recovery has failed. 'Succeeded' indicates that the Recovery has finished successfully. 'SucceededWithWarning' indicates that the Recovery finished successfully, but there were some warning messages. 'Skipped' indicates that the Recovery task was skipped.
			  * Constraints: Allowable values are: `Accepted`, `Running`, `Canceled`, `Canceling`, `Failed`, `Missed`, `Succeeded`, `SucceededWithWarning`, `OnHold`, `Finalizing`, `Skipped`, `LegalHold`.
		* `recover_file_and_folder_params` - (List) Specifies the parameters to perform a file and folder recovery.
		Nested schema for **recover_file_and_folder_params**:
			* `files_and_folders` - (List) Specifies the information about the files and folders to be recovered.
			Nested schema for **files_and_folders**:
				* `absolute_path` - (String) Specifies the absolute path to the file or folder.
				* `destination_dir` - (String) Specifies the destination directory where the file/directory was copied.
				* `is_directory` - (Boolean) Specifies whether this is a directory or not.
				* `is_view_file_recovery` - (Boolean) Specify if the recovery is of type view file/folder.
				* `messages` - (List) Specify error messages about the file during recovery.
				* `status` - (String) Specifies the recovery status for this file or folder.
				  * Constraints: Allowable values are: `NotStarted`, `EstimationInProgress`, `EstimationDone`, `CopyInProgress`, `Finished`.
			* `physical_target_params` - (List) Specifies the parameters to recover to a Physical target.
			Nested schema for **physical_target_params**:
				* `alternate_restore_directory` - (String) Specifies the directory path where restore should happen if restore_to_original_paths is set to false.
				* `continue_on_error` - (Boolean) Specifies whether to continue recovering other volumes if one of the volumes fails to recover. Default value is false.
				* `overwrite_existing` - (Boolean) Specifies whether to overwrite existing file/folder during recovery.
				* `preserve_acls` - (Boolean) Whether to preserve the ACLs of the original file.
				* `preserve_attributes` - (Boolean) Specifies whether to preserve file/folder attributes during recovery.
				* `preserve_timestamps` - (Boolean) Whether to preserve the original time stamps.
				* `recover_target` - (List) Specifies the target entity where the volumes are being mounted.
				Nested schema for **recover_target**:
					* `id` - (Integer) Specifies the id of the object.
					* `name` - (String) Specifies the name of the object.
					* `parent_source_id` - (Integer) Specifies the id of the parent source of the target.
					* `parent_source_name` - (String) Specifies the name of the parent source of the target.
				* `restore_entity_type` - (String) Specifies the restore type (restore everything or ACLs only) when restoring or downloading files or folders from a Physical file based or block based backup snapshot.
				  * Constraints: Allowable values are: `kRegular`, `kACLOnly`.
				* `restore_to_original_paths` - (Boolean) If this is true, then files will be restored to original paths.
				* `save_success_files` - (Boolean) Specifies whether to save success files or not. Default value is false.
				* `vlan_config` - (List) Specifies VLAN Params associated with the recovered. If this is not specified, then the VLAN settings will be automatically selected from one of the below options: a. If VLANs are configured on Cohesity, then the VLAN host/VIP will be automatically based on the client's (e.g. ESXI host) IP address. b. If VLANs are not configured on Cohesity, then the partition hostname or VIPs will be used for Recovery.
				Nested schema for **vlan_config**:
					* `disable_vlan` - (Boolean) If this is set to true, then even if VLANs are configured on the system, the partition VIPs will be used for the Recovery.
					* `id` - (Integer) If this is set, then the Cohesity host name or the IP address associated with this vlan is used for mounting Cohesity's view on the remote host.
					* `interface_name` - (String) Interface group to use for Recovery.
			* `target_environment` - (String) Specifies the environment of the recovery target. The corresponding params below must be filled out.
			  * Constraints: Allowable values are: `kPhysical`.
		* `recover_volume_params` - (List) Specifies the parameters to recover Physical Volumes.
		Nested schema for **recover_volume_params**:
			* `physical_target_params` - (List) Specifies the params for recovering to a physical target.
			Nested schema for **physical_target_params**:
				* `force_unmount_volume` - (Boolean) Specifies whether volume would be dismounted first during LockVolume failure. If not specified, default is false.
				* `mount_target` - (List) Specifies the target entity where the volumes are being mounted.
				Nested schema for **mount_target**:
					* `id` - (Integer) Specifies the id of the object.
					* `name` - (String) Specifies the name of the object.
				* `vlan_config` - (List) Specifies VLAN Params associated with the recovered. If this is not specified, then the VLAN settings will be automatically selected from one of the below options: a. If VLANs are configured on Cohesity, then the VLAN host/VIP will be automatically based on the client's (e.g. ESXI host) IP address. b. If VLANs are not configured on Cohesity, then the partition hostname or VIPs will be used for Recovery.
				Nested schema for **vlan_config**:
					* `disable_vlan` - (Boolean) If this is set to true, then even if VLANs are configured on the system, the partition VIPs will be used for the Recovery.
					* `id` - (Integer) If this is set, then the Cohesity host name or the IP address associated with this vlan is used for mounting Cohesity's view on the remote host.
					* `interface_name` - (String) Interface group to use for Recovery.
				* `volume_mapping` - (List) Specifies the mapping from source volumes to destination volumes.
				Nested schema for **volume_mapping**:
					* `destination_volume_guid` - (String) Specifies the guid of the destination volume.
					* `source_volume_guid` - (String) Specifies the guid of the source volume.
			* `target_environment` - (String) Specifies the environment of the recovery target. The corresponding params below must be filled out.
			  * Constraints: Allowable values are: `kPhysical`.
		* `recovery_action` - (String) Specifies the type of recover action to be performed.
		  * Constraints: Allowable values are: `RecoverPhysicalVolumes`, `InstantVolumeMount`, `RecoverFiles`, `RecoverSystem`.
		* `system_recovery_params` - (List) Specifies the parameters to perform a system recovery.
		Nested schema for **system_recovery_params**:
			* `full_nas_path` - (String) Specifies the path to the recovery view.
	* `progress_task_id` - (String) Progress monitor task id for Recovery.
	* `recovery_action` - (String) Specifies the type of recover action.
	  * Constraints: Allowable values are: `RecoverVMs`, `RecoverFiles`, `InstantVolumeMount`, `RecoverVmDisks`, `RecoverVApps`, `RecoverVAppTemplates`, `UptierSnapshot`, `RecoverRDS`, `RecoverAurora`, `RecoverS3Buckets`, `RecoverRDSPostgres`, `RecoverAzureSQL`, `RecoverApps`, `CloneApps`, `RecoverNasVolume`, `RecoverPhysicalVolumes`, `RecoverSystem`, `RecoverExchangeDbs`, `CloneAppView`, `RecoverSanVolumes`, `RecoverSanGroup`, `RecoverMailbox`, `RecoverOneDrive`, `RecoverSharePoint`, `RecoverPublicFolders`, `RecoverMsGroup`, `RecoverMsTeam`, `ConvertToPst`, `DownloadChats`, `RecoverMailboxCSM`, `RecoverOneDriveCSM`, `RecoverSharePointCSM`, `RecoverNamespaces`, `RecoverObjects`, `RecoverSfdcObjects`, `RecoverSfdcOrg`, `RecoverSfdcRecords`, `DownloadFilesAndFolders`, `CloneVMs`, `CloneView`, `CloneRefreshApp`, `CloneVMsToView`, `ConvertAndDeployVMs`, `DeployVMs`.
	* `retrieve_archive_tasks` - (List) Specifies the list of persistent state of a retrieve of an archive task.
	Nested schema for **retrieve_archive_tasks**:
		* `task_uid` - (String) Specifies the globally unique id for this retrieval of an archive task.
		  * Constraints: The value must match regular expression `/^\\d+:\\d+:\\d+$/`.
		* `uptier_expiry_times` - (List) Specifies how much time the retrieved entity is present in the hot-tiers.
	* `snapshot_environment` - (String) Specifies the type of snapshot environment for which the Recovery was performed.
	  * Constraints: Allowable values are: `kPhysical`, `kSQL`, `kKubernetes`.
	* `start_time_usecs` - (Integer) Specifies the start time of the Recovery in Unix timestamp epoch in microseconds.
	* `status` - (String) Status of the Recovery. 'Running' indicates that the Recovery is still running. 'Canceled' indicates that the Recovery has been cancelled. 'Canceling' indicates that the Recovery is in the process of being cancelled. 'Failed' indicates that the Recovery has failed. 'Succeeded' indicates that the Recovery has finished successfully. 'SucceededWithWarning' indicates that the Recovery finished successfully, but there were some warning messages. 'Skipped' indicates that the Recovery task was skipped.
	  * Constraints: Allowable values are: `Accepted`, `Running`, `Canceled`, `Canceling`, `Failed`, `Missed`, `Succeeded`, `SucceededWithWarning`, `OnHold`, `Finalizing`, `Skipped`, `LegalHold`.
	* `tear_down_message` - (String) Specifies the error message about the tear down operation if it fails.
	* `tear_down_status` - (String) Specifies the status of the tear down operation. This is only set when the canTearDown is set to true. 'DestroyScheduled' indicates that the tear down is ready to schedule. 'Destroying' indicates that the tear down is still running. 'Destroyed' indicates that the tear down succeeded. 'DestroyError' indicates that the tear down failed.
	  * Constraints: Allowable values are: `DestroyScheduled`, `Destroying`, `Destroyed`, `DestroyError`.

