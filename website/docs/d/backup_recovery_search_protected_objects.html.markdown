---
layout: "ibm"
page_title: "IBM : ibm_backup_recovery_search_protected_objects"
description: |-
  Get information about Protected Objects Search Result
subcategory: "IBM Backup Recovery"
---

# ibm_backup_recovery_search_protected_objects

Provides a read-only data source to retrieve information about a Protected Objects Search Result. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_backup_recovery_search_protected_objects" "backup_recovery_search_protected_objects" {
	x_ibm_tenant_id = "x_ibm_tenant_id"
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `cdp_protected_only` - (Optional, Boolean) Specifies whether to only return the CDP protected objects.
* `endpoint_type` - (Optional, String) Backup Recovery Endpoint type. By default set to "public".
* `instance_id` - (Optional, String) Backup Recovery instance ID. If provided here along with region, the provider constructs the endpoint URL using them, which overrides any value set through environment variables or the `endpoints.json` file.
* `region` - (Optional, String) Backup Recovery region. If provided here along with instance_id, the provider constructs the endpoint URL using them, which overrides any value set through environment variables or the `endpoints.json` file.  
* `environments` - (Optional, List) Specifies the environment type to filter objects.
  * Constraints: Allowable list items are: `kPhysical`, `kSQL`.
* `filter_snapshot_from_usecs` - (Optional, Integer) Specifies the timestamp in Unix time epoch in microseconds to filter the objects if the Object has a successful snapshot after this value.
* `filter_snapshot_to_usecs` - (Optional, Integer) Specifies the timestamp in Unix time epoch in microseconds to filter the objects if the Object has a successful snapshot before this value.
* `object_action_key` - (Optional, String) Filter by ObjectActionKey, which uniquely represents protection of an object. An object can be protected in multiple ways but atmost once for a given combination of ObjectActionKey. When specified, latest snapshot info matching the objectActionKey is for corresponding object.
  * Constraints: Allowable values are: `kPhysical`, `kSQL`.
* `object_ids` - (Optional, List) Specifies a list of Object ids to filter.
* `os_types` - (Optional, List) Specifies the operating system types to filter objects on.
  * Constraints: Allowable list items are: `kLinux`, `kWindows`.
* `protection_group_ids` - (Optional, List) Specifies a list of Protection Group ids to filter the objects. If specified, the objects protected by specified Protection Group ids will be returned.
* `request_initiator_type` - (Optional, String) Specifies the type of request from UI, which is used for services like magneto to determine the priority of requests.
  * Constraints: Allowable values are: `UIUser`, `UIAuto`, `Helios`.
* `run_instance_ids` - (Optional, List) Specifies a list of run instance ids. If specified only objects belonging to the provided run id will be retunrned.
* `search_string` - (Optional, String) Specifies the search string to filter the objects. This search string will be applicable for objectnames and Protection Group names. User can specify a wildcard character '*' as a suffix to a string where all object and their Protection Group names are matched with the prefix string. For example, if vm1 and vm2 are the names of objects, user can specify vm* to list the objects. If not specified, then all the objects with Protection Groups will be returned which will match other filtering criteria.
* `snapshot_actions` - (Optional, List) Specifies a list of recovery actions. Only snapshots that applies to these actions will be returned.
  * Constraints: Allowable list items are: `RecoverVMs`, `RecoverFiles`, `InstantVolumeMount`, `RecoverVmDisks`, `MountVolumes`, `RecoverVApps`, `RecoverRDS`, `RecoverAurora`, `RecoverS3Buckets`, `RecoverApps`, `RecoverNasVolume`, `RecoverPhysicalVolumes`, `RecoverSystem`, `RecoverSanVolumes`, `RecoverNamespaces`, `RecoverObjects`, `DownloadFilesAndFolders`, `RecoverPublicFolders`, `RecoverVAppTemplates`, `RecoverMailbox`, `RecoverOneDrive`, `RecoverMsTeam`, `RecoverMsGroup`, `RecoverSharePoint`, `ConvertToPst`, `RecoverSfdcRecords`, `RecoverAzureSQL`, `DownloadChats`, `RecoverRDSPostgres`, `RecoverMailboxCSM`, `RecoverOneDriveCSM`, `RecoverSharePointCSM`.
* `source_ids` - (Optional, List) Specifies a list of Protection Source object ids to filter the objects. If specified, the object which are present in those Sources will be returned.
* `sub_result_size` - (Optional, Integer) Specifies the size of objects to be fetched for a single subresult.
* `use_cached_data` - (Optional, Boolean) Specifies whether we can serve the GET request to the read replica cache cache. There is a lag of 15 seconds between the read replica and primary data source.
* `x_ibm_tenant_id` - (Required, String) Specifies the key to be used to encrypt the source credential. If includeSourceCredentials is set to true this key must be specified.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the Protected Objects Search Result.
* `metadata` - (List) Specifies the metadata information about the Protection Groups, Protection Policy etc., for search result.
Nested schema for **metadata**:
	* `unique_protection_group_identifiers` - (List) Specifies the list of unique Protection Group identifiers for all the Objects returned in the response.
	Nested schema for **unique_protection_group_identifiers**:
		* `protection_group_id` - (String) Specifies Protection Group id.
		* `protection_group_name` - (String) Specifies Protection Group name.
* `num_results` - (Integer) Specifies the total number of search results which matches the search criteria.
* `objects` - (List) Specifies the list of Protected Objects.
Nested schema for **objects**:
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
	* `latest_snapshots_info` - (List) Specifies the latest snapshot information for every Protection Group for a given object.
	Nested schema for **latest_snapshots_info**:
		* `archival_snapshots_info` - (List) Specifies the archival snapshots information.
		Nested schema for **archival_snapshots_info**:
			* `archival_task_id` - (String) Specifies the archival task id. This is a protection group UID which only applies when archival type is 'Tape'.
			* `logical_size_bytes` - (Integer) Specifies the logical size of this snapshot in bytes.
			* `ownership_context` - (String) Specifies the ownership context for the target.
			  * Constraints: Allowable values are: `Local`, `FortKnox`.
			* `snapshot_id` - (String) Specifies the id of the archival snapshot for the object.
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
		* `indexing_status` - (String) Specifies the indexing status of objects in this snapshot.<br> 'InProgress' indicates the indexing is in progress.<br> 'Done' indicates indexing is done.<br> 'NoIndex' indicates indexing is not applicable.<br> 'Error' indicates indexing failed with error.
		  * Constraints: Allowable values are: `InProgress`, `Done`, `NoIndex`, `Error`.
		* `local_snapshot_info` - (List) Specifies the local snapshot information.
		Nested schema for **local_snapshot_info**:
			* `logical_size_bytes` - (Integer) Specifies the logical size of this snapshot in bytes.
			* `snapshot_id` - (String) Specifies the id of the local snapshot for the object.
		* `protection_group_id` - (String) Specifies id of the Protection Group.
		* `protection_group_name` - (String) Specifies name of the Protection Group.
		* `protection_run_end_time_usecs` - (Integer) Specifies the end time of Protection Group Run in Unix timestamp epoch in microseconds.
		* `protection_run_id` - (String) Specifies the id of Protection Group Run.
		* `protection_run_start_time_usecs` - (Integer) Specifies the start time of Protection Group Run in Unix timestamp epoch in microseconds.
		* `run_instance_id` - (Integer) Specifies the instance id of the protection run which create the snapshot.
		* `run_type` - (String) Specifies the type of protection run created this snapshot.
		  * Constraints: Allowable values are: `kRegular`, `kFull`, `kLog`, `kSystem`, `kHydrateCDP`, `kStorageArraySnapshot`.
		* `source_group_id` - (String) Specifies the source protection group id in case of replication.
	* `logical_size_bytes` - (Integer) Specifies the logical size of object in bytes.
	* `mssql_params` - (List) Specifies the parameters for Msssql object.
	Nested schema for **mssql_params**:
		* `aag_info` - (List) Object details for Mssql.
		Nested schema for **aag_info**:
			* `name` - (String) Specifies the AAG name.
			* `object_id` - (Integer) Specifies the AAG object Id.
		* `host_info` - (List) Specifies the host information for a objects. This is mainly populated in case of App objects where app object is hosted by another object such as VM or physical server.
		Nested schema for **host_info**:
			* `environment` - (String) Specifies the environment of the object.
			  * Constraints: Allowable values are: `kPhysical`, `kSQL`.
			* `id` - (String) Specifies the id of the host object.
			* `name` - (String) Specifies the name of the host object.
		* `is_encrypted` - (Boolean) Specifies whether the database is TDE enabled.
	* `name` - (String) Specifies the name of the object.
	* `object_hash` - (String) Specifies the hash identifier of the object.
	* `object_type` - (String) Specifies the type of the object.
	  * Constraints: Allowable values are: `kCluster`, `kVserver`, `kVolume`, `kVCenter`, `kStandaloneHost`, `kvCloudDirector`, `kFolder`, `kDatacenter`, `kComputeResource`, `kClusterComputeResource`, `kResourcePool`, `kDatastore`, `kHostSystem`, `kVirtualMachine`, `kVirtualApp`, `kStoragePod`, `kNetwork`, `kDistributedVirtualPortgroup`, `kTagCategory`, `kTag`, `kOpaqueNetwork`, `kOrganization`, `kVirtualDatacenter`, `kCatalog`, `kOrgMetadata`, `kStoragePolicy`, `kVirtualAppTemplate`, `kDomain`, `kOutlook`, `kMailbox`, `kUsers`, `kGroups`, `kSites`, `kUser`, `kGroup`, `kSite`, `kApplication`, `kGraphUser`, `kPublicFolders`, `kPublicFolder`, `kTeams`, `kTeam`, `kRootPublicFolder`, `kO365Exchange`, `kO365OneDrive`, `kO365Sharepoint`, `kKeyspace`, `kTable`, `kDatabase`, `kCollection`, `kBucket`, `kNamespace`, `kSCVMMServer`, `kStandaloneCluster`, `kHostGroup`, `kHypervHost`, `kHostCluster`, `kCustomProperty`, `kTenant`, `kSubscription`, `kResourceGroup`, `kStorageAccount`, `kStorageKey`, `kStorageContainer`, `kStorageBlob`, `kNetworkSecurityGroup`, `kVirtualNetwork`, `kSubnet`, `kComputeOptions`, `kSnapshotManagerPermit`, `kAvailabilitySet`, `kOVirtManager`, `kHost`, `kStorageDomain`, `kVNicProfile`, `kIAMUser`, `kRegion`, `kAvailabilityZone`, `kEC2Instance`, `kVPC`, `kInstanceType`, `kKeyPair`, `kRDSOptionGroup`, `kRDSParameterGroup`, `kRDSInstance`, `kRDSSubnet`, `kRDSTag`, `kAuroraTag`, `kAuroraCluster`, `kAccount`, `kSubTaskPermit`, `kS3Bucket`, `kS3Tag`, `kKmsKey`, `kProject`, `kLabel`, `kMetadata`, `kVPCConnector`, `kPrismCentral`, `kOtherHypervisorCluster`, `kZone`, `kMountPoint`, `kStorageArray`, `kFileSystem`, `kContainer`, `kFilesystem`, `kFileset`, `kPureProtectionGroup`, `kVolumeGroup`, `kStoragePool`, `kViewBox`, `kView`, `kWindowsCluster`, `kOracleRACCluster`, `kOracleAPCluster`, `kService`, `kPVC`, `kPersistentVolumeClaim`, `kPersistentVolume`, `kRootContainer`, `kDAGRootContainer`, `kExchangeNode`, `kExchangeDAGDatabaseCopy`, `kExchangeStandaloneDatabase`, `kExchangeDAG`, `kExchangeDAGDatabase`, `kDomainController`, `kInstance`, `kAAG`, `kAAGRootContainer`, `kAAGDatabase`, `kRACRootContainer`, `kTableSpace`, `kPDB`, `kObject`, `kOrg`, `kAppInstance`.
	* `os_type` - (String) Specifies the operating system type of the object.
	  * Constraints: Allowable values are: `kLinux`, `kWindows`.
	* `permissions` - (List) Specifies the list of users, groups and users that have permissions for a given object.
	Nested schema for **permissions**:
		* `groups` - (List) Specifies the list of user groups which has permissions to the object.
		Nested schema for **groups**:
			* `domain` - (String) Specifies the domain of the user group.
			* `name` - (String) Specifies the name of the user group.
			* `sid` - (String) Specifies the sid of the user group.
		* `object_id` - (Integer) Specifies the id of the object.
		* `tenant` - (List) Specifies a tenant object.
		Nested schema for **tenant**:
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
		* `users` - (List) Specifies the list of users which has the permissions to the object.
		Nested schema for **users**:
			* `domain` - (String) Specifies the domain of the user.
			* `name` - (String) Specifies the name of the user.
			* `sid` - (String) Specifies the sid of the user.
	* `physical_params` - (List) Specifies the parameters for Physical object.
	Nested schema for **physical_params**:
		* `enable_system_backup` - (Boolean) Specifies if system backup was enabled for the source in a particular run.
	* `protection_stats` - (List) Specifies the count and size of protected and unprotected objects for the size.
	Nested schema for **protection_stats**:
		* `deleted_protected_count` - (Integer) Specifies the count of protected leaf objects which were deleted from the source after being protected.
		* `environment` - (String) Specifies the environment of the object.
		  * Constraints: Allowable values are: `kPhysical`, `kSQL`, `kOracle`.
		* `protected_count` - (Integer) Specifies the count of the protected leaf objects.
		* `protected_size_bytes` - (Integer) Specifies the protected logical size in bytes.
		* `unprotected_count` - (Integer) Specifies the count of the unprotected leaf objects.
		* `unprotected_size_bytes` - (Integer) Specifies the unprotected logical size in bytes.
	* `protection_type` - (String) Specifies the protection type of the object if any.
	  * Constraints: Allowable values are: `kAgent`, `kNative`, `kSnapshotManager`, `kRDSSnapshotManager`, `kAuroraSnapshotManager`, `kAwsS3`, `kAwsRDSPostgresBackup`, `kAwsAuroraPostgres`, `kAwsRDSPostgres`, `kAzureSQL`, `kFile`, `kVolume`.
	* `sharepoint_site_summary` - (List) Specifies the common parameters for Sharepoint site objects.
	Nested schema for **sharepoint_site_summary**:
		* `site_web_url` - (String) Specifies the web url for the Sharepoint site.
	* `source_id` - (Integer) Specifies registered source id to which object belongs.
	* `source_info` - (List) Specifies the Source Object information.
	Nested schema for **source_info**:
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
	* `source_name` - (String) Specifies registered source name to which object belongs.
	* `uuid` - (String) Specifies the uuid which is a unique identifier of the object.
	* `v_center_summary` - (List)
	Nested schema for **v_center_summary**:
		* `is_cloud_env` - (Boolean) Specifies that registered vCenter source is a VMC (VMware Cloud) environment or not.
	* `windows_cluster_summary` - (List)
	Nested schema for **windows_cluster_summary**:
		* `cluster_source_type` - (String) Specifies the type of cluster resource this source represents.

