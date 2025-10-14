---
layout: "ibm"
page_title: "IBM : ibm_backup_recovery_download_files_folders"
description: |-
  Manages Download Files And Folders Recovery Params..
subcategory: "IBM Backup Recovery"
---

# ibm_backup_recovery_download_files_folders

Create Download Files And Folders Recovery Paramss with this resource.

**Note**
ibm_backup_recovery_download_files_folders resource does not support update or delete operations due to the absence of corresponding API endpoints. As a result, Terraform cannot manage these operations for those resources. Users should be aware that removing these resources from the configuration (main.tf) will only remove them from the Terraform state and will not affect the actual resources in the backend. Similarly updating these resources will throw an error in the plan phase stating that the resource cannot be updated.

**Important:** When managing resources that lack complete CRUD operations, users should exercise caution and consider the limitations described above. Manual intervention may be required to manage these resources in the backend if updates or deletions are necessary.**

## Example Usage

```hcl
resource "ibm_backup_recovery_download_files_folders" "backup_recovery_download_files_folders_instance" {
  documents {
		is_directory = true
		item_id = "item_id"
  }
  files_and_folders {
		absolute_path = "absolute_path"
		is_directory = true
  }
  name = "name"
  object {
		snapshot_id = "snapshot_id"
		point_in_time_usecs = 1
		protection_group_id = "protection_group_id"
		protection_group_name = "protection_group_name"
		snapshot_creation_time_usecs = 1
		object_info {
			id = 1
			name = "name"
			source_id = 1
			source_name = "source_name"
			environment = "kPhysical"
			object_hash = "object_hash"
			object_type = "kCluster"
			logical_size_bytes = 1
			uuid = "uuid"
			global_id = "global_id"
			protection_type = "kAgent"
			sharepoint_site_summary {
				site_web_url = "site_web_url"
			}
			os_type = "kLinux"
			child_objects {
				id = 1
				name = "name"
				source_id = 1
				source_name = "source_name"
				environment = "kPhysical"
				object_hash = "object_hash"
				object_type = "kCluster"
				logical_size_bytes = 1
				uuid = "uuid"
				global_id = "global_id"
				protection_type = "kAgent"
				sharepoint_site_summary {
					site_web_url = "site_web_url"
				}
				os_type = "kLinux"
				v_center_summary {
					is_cloud_env = true
				}
				windows_cluster_summary {
					cluster_source_type = "cluster_source_type"
				}
			}
			v_center_summary {
				is_cloud_env = true
			}
			windows_cluster_summary {
				cluster_source_type = "cluster_source_type"
			}
		}
		snapshot_target_type = "Local"
		archival_target_info {
			target_id = 1
			archival_task_id = "archival_task_id"
			target_name = "target_name"
			target_type = "Tape"
			usage_type = "Archival"
			ownership_context = "Local"
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
				current_tier_type = "kAmazonS3Standard"
			}
		}
		progress_task_id = "progress_task_id"
		recover_from_standby = true
		status = "Accepted"
		start_time_usecs = 1
		end_time_usecs = 1
		messages = [ "messages" ]
		bytes_restored = 1
  }
  x_ibm_tenant_id = "x_ibm_tenant_id"
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `documents` - (Optional, Forces new resource, List) Specifies the list of documents to download using item ids. Only one of filesAndFolders or documents should be used. Currently only files are supported by documents.
  * Constraints: The minimum length is `1` item.
Nested schema for **documents**:
	* `is_directory` - (Optional, Boolean) Specifies whether the document is a directory. Since currently only files are supported this should always be false.
	* `item_id` - (Required, String) Specifies the item id of the document.
* `files_and_folders` - (Required, Forces new resource, List) Specifies the list of files and folders to download.
  * Constraints: The minimum length is `1` item.
Nested schema for **files_and_folders**:
	* `absolute_path` - (Required, String) Specifies the absolute path of the file or folder.
	* `is_directory` - (Optional, Boolean) Specifies whether the file or folder object is a directory.
* `glacier_retrieval_type` - (Optional, Forces new resource, String) Specifies the glacier retrieval type when restoring or downloding files or folders from a Glacier-based cloud snapshot.
  * Constraints: Allowable values are: `kStandard`, `kExpeditedNoPCU`, `kExpeditedWithPCU`.
* `name` - (Required, Forces new resource, String) Specifies the name of the recovery task. This field must be set and must be a unique name.
* `backup_recovery_endpoint` - (Optional, String) Backup Recovery Endpoint URL. If provided here, it overrides values configured via environment variable (IBMCLOUD_BACKUP_RECOVERY_ENDPOINT) or endpoints.json.   
* `object` - (Required, Forces new resource, List) Specifies the common snapshot parameters for a protected object.
Nested schema for **object**:
	* `archival_target_info` - (Optional, List) Specifies the archival target information if the snapshot is an archival snapshot.
	Nested schema for **archival_target_info**:
		* `archival_task_id` - (Computed, String) Specifies the archival task id. This is a protection group UID which only applies when archival type is 'Tape'.
		* `ownership_context` - (Computed, String) Specifies the ownership context for the target.
		  * Constraints: Allowable values are: `Local`, `FortKnox`.
		* `target_id` - (Computed, Integer) Specifies the archival target ID.
		* `target_name` - (Computed, String) Specifies the archival target name.
		* `target_type` - (Computed, String) Specifies the archival target type.
		  * Constraints: Allowable values are: `Tape`, `Cloud`, `Nas`.
		* `tier_settings` - (Optional, List) Specifies the tier info for archival.
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
			* `cloud_platform` - (Computed, String) Specifies the cloud platform to enable tiering.
			  * Constraints: Allowable values are: `AWS`, `Azure`, `Oracle`, `Google`.
			* `current_tier_type` - (Computed, String) Specifies the type of the current tier where the snapshot resides. This will be specified if the run is a CAD run.
			  * Constraints: Allowable values are: `kAmazonS3Standard`, `kAmazonS3StandardIA`, `kAmazonS3OneZoneIA`, `kAmazonS3IntelligentTiering`, `kAmazonS3Glacier`, `kAmazonS3GlacierDeepArchive`, `kAzureTierHot`, `kAzureTierCool`, `kAzureTierArchive`, `kGoogleStandard`, `kGoogleRegional`, `kGoogleMultiRegional`, `kGoogleNearline`, `kGoogleColdline`, `kOracleTierStandard`, `kOracleTierArchive`.
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
		* `usage_type` - (Computed, String) Specifies the usage type for the target.
		  * Constraints: Allowable values are: `Archival`, `Tiering`, `Rpaas`.
	* `bytes_restored` - (Computed, Integer) Specify the total bytes restored.
	* `end_time_usecs` - (Computed, Integer) Specifies the end time of the Recovery in Unix timestamp epoch in microseconds. This field will be populated only after Recovery is finished.
	* `messages` - (Computed, List) Specify error messages about the object.
	* `object_info` - (Optional, List) Specifies the information about the object for which the snapshot is taken.
	Nested schema for **object_info**:
		* `child_objects` - (Optional, List) Specifies child object details.
		Nested schema for **child_objects**:
			* `child_objects` - (Optional, List) Specifies child object details.
			Nested schema for **child_objects**:
			* `environment` - (Computed, String) Specifies the environment of the object.
			  * Constraints: Allowable values are: `kPhysical`, `kSQL`.
			* `global_id` - (Computed, String) Specifies the global id which is a unique identifier of the object.
			* `id` - (Computed, Integer) Specifies object id.
			* `logical_size_bytes` - (Computed, Integer) Specifies the logical size of object in bytes.
			* `name` - (Computed, String) Specifies the name of the object.
			* `object_hash` - (Computed, String) Specifies the hash identifier of the object.
			* `object_type` - (Computed, String) Specifies the type of the object.
			  * Constraints: Allowable values are: `kCluster`, `kVserver`, `kVolume`, `kVCenter`, `kStandaloneHost`, `kvCloudDirector`, `kFolder`, `kDatacenter`, `kComputeResource`, `kClusterComputeResource`, `kResourcePool`, `kDatastore`, `kHostSystem`, `kVirtualMachine`, `kVirtualApp`, `kStoragePod`, `kNetwork`, `kDistributedVirtualPortgroup`, `kTagCategory`, `kTag`, `kOpaqueNetwork`, `kOrganization`, `kVirtualDatacenter`, `kCatalog`, `kOrgMetadata`, `kStoragePolicy`, `kVirtualAppTemplate`, `kDomain`, `kOutlook`, `kMailbox`, `kUsers`, `kGroups`, `kSites`, `kUser`, `kGroup`, `kSite`, `kApplication`, `kGraphUser`, `kPublicFolders`, `kPublicFolder`, `kTeams`, `kTeam`, `kRootPublicFolder`, `kO365Exchange`, `kO365OneDrive`, `kO365Sharepoint`, `kKeyspace`, `kTable`, `kDatabase`, `kCollection`, `kBucket`, `kNamespace`, `kSCVMMServer`, `kStandaloneCluster`, `kHostGroup`, `kHypervHost`, `kHostCluster`, `kCustomProperty`, `kTenant`, `kSubscription`, `kResourceGroup`, `kStorageAccount`, `kStorageKey`, `kStorageContainer`, `kStorageBlob`, `kNetworkSecurityGroup`, `kVirtualNetwork`, `kSubnet`, `kComputeOptions`, `kSnapshotManagerPermit`, `kAvailabilitySet`, `kOVirtManager`, `kHost`, `kStorageDomain`, `kVNicProfile`, `kIAMUser`, `kRegion`, `kAvailabilityZone`, `kEC2Instance`, `kVPC`, `kInstanceType`, `kKeyPair`, `kRDSOptionGroup`, `kRDSParameterGroup`, `kRDSInstance`, `kRDSSubnet`, `kRDSTag`, `kAuroraTag`, `kAuroraCluster`, `kAccount`, `kSubTaskPermit`, `kS3Bucket`, `kS3Tag`, `kKmsKey`, `kProject`, `kLabel`, `kMetadata`, `kVPCConnector`, `kPrismCentral`, `kOtherHypervisorCluster`, `kZone`, `kMountPoint`, `kStorageArray`, `kFileSystem`, `kContainer`, `kFilesystem`, `kFileset`, `kPureProtectionGroup`, `kVolumeGroup`, `kStoragePool`, `kViewBox`, `kView`, `kWindowsCluster`, `kOracleRACCluster`, `kOracleAPCluster`, `kService`, `kPVC`, `kPersistentVolumeClaim`, `kPersistentVolume`, `kRootContainer`, `kDAGRootContainer`, `kExchangeNode`, `kExchangeDAGDatabaseCopy`, `kExchangeStandaloneDatabase`, `kExchangeDAG`, `kExchangeDAGDatabase`, `kDomainController`, `kInstance`, `kAAG`, `kAAGRootContainer`, `kAAGDatabase`, `kRACRootContainer`, `kTableSpace`, `kPDB`, `kObject`, `kOrg`, `kAppInstance`.
			* `os_type` - (Computed, String) Specifies the operating system type of the object.
			  * Constraints: Allowable values are: `kLinux`, `kWindows`.
			* `protection_type` - (Computed, String) Specifies the protection type of the object if any.
			  * Constraints: Allowable values are: `kAgent`, `kNative`, `kSnapshotManager`, `kRDSSnapshotManager`, `kAuroraSnapshotManager`, `kAwsS3`, `kAwsRDSPostgresBackup`, `kAwsAuroraPostgres`, `kAwsRDSPostgres`, `kAzureSQL`, `kFile`, `kVolume`.
			* `sharepoint_site_summary` - (Optional, List) Specifies the common parameters for Sharepoint site objects.
			Nested schema for **sharepoint_site_summary**:
				* `site_web_url` - (Computed, String) Specifies the web url for the Sharepoint site.
			* `source_id` - (Computed, Integer) Specifies registered source id to which object belongs.
			* `source_name` - (Computed, String) Specifies registered source name to which object belongs.
			* `uuid` - (Computed, String) Specifies the uuid which is a unique identifier of the object.
			* `v_center_summary` - (Optional, List)
			Nested schema for **v_center_summary**:
				* `is_cloud_env` - (Computed, Boolean) Specifies that registered vCenter source is a VMC (VMware Cloud) environment or not.
			* `windows_cluster_summary` - (Optional, List)
			Nested schema for **windows_cluster_summary**:
				* `cluster_source_type` - (Computed, String) Specifies the type of cluster resource this source represents.
		* `environment` - (Computed, String) Specifies the environment of the object.
		  * Constraints: Allowable values are: `kPhysical`, `kSQL`.
		* `global_id` - (Computed, String) Specifies the global id which is a unique identifier of the object.
		* `id` - (Computed, Integer) Specifies object id.
		* `logical_size_bytes` - (Computed, Integer) Specifies the logical size of object in bytes.
		* `name` - (Computed, String) Specifies the name of the object.
		* `object_hash` - (Computed, String) Specifies the hash identifier of the object.
		* `object_type` - (Computed, String) Specifies the type of the object.
		  * Constraints: Allowable values are: `kCluster`, `kVserver`, `kVolume`, `kVCenter`, `kStandaloneHost`, `kvCloudDirector`, `kFolder`, `kDatacenter`, `kComputeResource`, `kClusterComputeResource`, `kResourcePool`, `kDatastore`, `kHostSystem`, `kVirtualMachine`, `kVirtualApp`, `kStoragePod`, `kNetwork`, `kDistributedVirtualPortgroup`, `kTagCategory`, `kTag`, `kOpaqueNetwork`, `kOrganization`, `kVirtualDatacenter`, `kCatalog`, `kOrgMetadata`, `kStoragePolicy`, `kVirtualAppTemplate`, `kDomain`, `kOutlook`, `kMailbox`, `kUsers`, `kGroups`, `kSites`, `kUser`, `kGroup`, `kSite`, `kApplication`, `kGraphUser`, `kPublicFolders`, `kPublicFolder`, `kTeams`, `kTeam`, `kRootPublicFolder`, `kO365Exchange`, `kO365OneDrive`, `kO365Sharepoint`, `kKeyspace`, `kTable`, `kDatabase`, `kCollection`, `kBucket`, `kNamespace`, `kSCVMMServer`, `kStandaloneCluster`, `kHostGroup`, `kHypervHost`, `kHostCluster`, `kCustomProperty`, `kTenant`, `kSubscription`, `kResourceGroup`, `kStorageAccount`, `kStorageKey`, `kStorageContainer`, `kStorageBlob`, `kNetworkSecurityGroup`, `kVirtualNetwork`, `kSubnet`, `kComputeOptions`, `kSnapshotManagerPermit`, `kAvailabilitySet`, `kOVirtManager`, `kHost`, `kStorageDomain`, `kVNicProfile`, `kIAMUser`, `kRegion`, `kAvailabilityZone`, `kEC2Instance`, `kVPC`, `kInstanceType`, `kKeyPair`, `kRDSOptionGroup`, `kRDSParameterGroup`, `kRDSInstance`, `kRDSSubnet`, `kRDSTag`, `kAuroraTag`, `kAuroraCluster`, `kAccount`, `kSubTaskPermit`, `kS3Bucket`, `kS3Tag`, `kKmsKey`, `kProject`, `kLabel`, `kMetadata`, `kVPCConnector`, `kPrismCentral`, `kOtherHypervisorCluster`, `kZone`, `kMountPoint`, `kStorageArray`, `kFileSystem`, `kContainer`, `kFilesystem`, `kFileset`, `kPureProtectionGroup`, `kVolumeGroup`, `kStoragePool`, `kViewBox`, `kView`, `kWindowsCluster`, `kOracleRACCluster`, `kOracleAPCluster`, `kService`, `kPVC`, `kPersistentVolumeClaim`, `kPersistentVolume`, `kRootContainer`, `kDAGRootContainer`, `kExchangeNode`, `kExchangeDAGDatabaseCopy`, `kExchangeStandaloneDatabase`, `kExchangeDAG`, `kExchangeDAGDatabase`, `kDomainController`, `kInstance`, `kAAG`, `kAAGRootContainer`, `kAAGDatabase`, `kRACRootContainer`, `kTableSpace`, `kPDB`, `kObject`, `kOrg`, `kAppInstance`.
		* `os_type` - (Computed, String) Specifies the operating system type of the object.
		  * Constraints: Allowable values are: `kLinux`, `kWindows`.
		* `protection_type` - (Computed, String) Specifies the protection type of the object if any.
		  * Constraints: Allowable values are: `kAgent`, `kNative`, `kSnapshotManager`, `kRDSSnapshotManager`, `kAuroraSnapshotManager`, `kAwsS3`, `kAwsRDSPostgresBackup`, `kAwsAuroraPostgres`, `kAwsRDSPostgres`, `kAzureSQL`, `kFile`, `kVolume`.
		* `sharepoint_site_summary` - (Optional, List) Specifies the common parameters for Sharepoint site objects.
		Nested schema for **sharepoint_site_summary**:
			* `site_web_url` - (Computed, String) Specifies the web url for the Sharepoint site.
		* `source_id` - (Computed, Integer) Specifies registered source id to which object belongs.
		* `source_name` - (Computed, String) Specifies registered source name to which object belongs.
		* `uuid` - (Computed, String) Specifies the uuid which is a unique identifier of the object.
		* `v_center_summary` - (Optional, List)
		Nested schema for **v_center_summary**:
			* `is_cloud_env` - (Computed, Boolean) Specifies that registered vCenter source is a VMC (VMware Cloud) environment or not.
		* `windows_cluster_summary` - (Optional, List)
		Nested schema for **windows_cluster_summary**:
			* `cluster_source_type` - (Computed, String) Specifies the type of cluster resource this source represents.
	* `point_in_time_usecs` - (Optional, Integer) Specifies the timestamp (in microseconds. from epoch) for recovering to a point-in-time in the past.
	* `progress_task_id` - (Computed, String) Progress monitor task id for Recovery of VM.
	* `protection_group_id` - (Optional, String) Specifies the protection group id of the object snapshot.
	* `protection_group_name` - (Optional, String) Specifies the protection group name of the object snapshot.
	* `recover_from_standby` - (Optional, Boolean) Specifies that user wants to perform standby restore if it is enabled for this object.
	* `snapshot_creation_time_usecs` - (Computed, Integer) Specifies the time when the snapshot is created in Unix timestamp epoch in microseconds.
	* `snapshot_id` - (Required, String) Specifies the snapshot id.
	* `snapshot_target_type` - (Computed, String) Specifies the snapshot target type.
	  * Constraints: Allowable values are: `Local`, `Archival`, `RpaasArchival`, `StorageArraySnapshot`, `Remote`.
	* `start_time_usecs` - (Computed, Integer) Specifies the start time of the Recovery in Unix timestamp epoch in microseconds.
	* `status` - (Computed, String) Status of the Recovery. 'Running' indicates that the Recovery is still running. 'Canceled' indicates that the Recovery has been cancelled. 'Canceling' indicates that the Recovery is in the process of being cancelled. 'Failed' indicates that the Recovery has failed. 'Succeeded' indicates that the Recovery has finished successfully. 'SucceededWithWarning' indicates that the Recovery finished successfully, but there were some warning messages. 'Skipped' indicates that the Recovery task was skipped.
	  * Constraints: Allowable values are: `Accepted`, `Running`, `Canceled`, `Canceling`, `Failed`, `Missed`, `Succeeded`, `SucceededWithWarning`, `OnHold`, `Finalizing`, `Skipped`, `LegalHold`.
* `parent_recovery_id` - (Optional, Forces new resource, String) If current recovery is child task triggered through another parent recovery operation, then this field will specify the id of the parent recovery.
  * Constraints: The value must match regular expression `/^\\d+:\\d+:\\d+$/`.
* `x_ibm_tenant_id` - (Required, Forces new resource, String) Specifies the key to be used to encrypt the source credential. If includeSourceCredentials is set to true this key must be specified.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the Download Files And Folders Recovery Params..


### Import
Not Supported