---
layout: "ibm"
page_title: "IBM : ibm_backup_recovery_search_objects"
description: |-
  Get information about Objects Search Result
subcategory: "IBM Backup Recovery"
---

# ibm_backup_recovery_search_objects

Provides a read-only data source to retrieve information about an Objects Search Result. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_backup_recovery_search_objects" "backup_recovery_search_objects" {
	x_ibm_tenant_id = "x_ibm_tenant_id"
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `cluster_identifiers` - (Optional, List) Specifies the list of cluster identifiers. Format is clusterId:clusterIncarnationId. Only records from clusters having these identifiers will be returned.
* `count` - (Optional, Integer) Specifies the number of objects to be fetched for the specified pagination cookie.
* `environments` - (Optional, List) Specifies the environment type to filter objects.
  * Constraints: Allowable list items are: `kPhysical`, `kSQL`.
* `external_filters` - (Optional, List) Specifies the key-value pairs to filtering the results for the search. Each filter is of the form 'key:value'. The filter 'externalFilters:k1:v1&externalFilters:k2:v2&externalFilters:k2:v3' returns the documents where each document will match the query (k1=v1) AND (k2=v2 OR k2 = v3). Allowed keys: - vmBiosUuid - graphUuid - arn - instanceId - bucketName - azureId.
* `include_deleted_objects` - (Optional, Boolean) Specifies whether to include deleted objects in response. These objects can't be protected but can be recovered. This field is deprecated.
* `include_helios_tag_info_for_objects` - (Optional, Boolean) pecifies whether to include helios tags information for objects in response. Default value is false.
* `is_deleted` - (Optional, Boolean) If set to true, then objects which are deleted on atleast one cluster will be returned. If not set or set to false then objects which are registered on atleast one cluster are returned.
* `is_protected` - (Optional, Boolean) Specifies the protection status of objects. If set to true, only protected objects will be returned. If set to false, only unprotected objects will be returned. If not specified, all objects will be returned.
* `last_run_status_list` - (Optional, List) Specifies a list of status of the object's last protection run. Only objects with last run status of these will be returned.
  * Constraints: Allowable list items are: `Accepted`, `Running`, `Canceled`, `Canceling`, `Failed`, `Missed`, `Succeeded`, `SucceededWithWarning`, `OnHold`, `Finalizing`, `Skipped`, `LegalHold`.
* `might_have_snapshot_tag_ids` - (Optional, List) Specifies list of snapshot tags, one or more of which might be present in the document. These are OR'ed together and the resulting criteria AND'ed with the rest of the query.
  * Constraints: The list items must match regular expression `/^\\d+:\\d+:[A-Z0-9-]+$/`.
* `might_have_tag_ids` - (Optional, List) Specifies list of tags, one or more of which might be present in the document. These are OR'ed together and the resulting criteria AND'ed with the rest of the query.
  * Constraints: The list items must match regular expression `/^\\d+:\\d+:[A-Z0-9-]+$/`.
* `must_have_snapshot_tag_ids` - (Optional, List) Specifies snapshot tags which must be all present in the document.
  * Constraints: The list items must match regular expression `/^\\d+:\\d+:[A-Z0-9-]+$/`.
* `must_have_tag_ids` - (Optional, List) Specifies tags which must be all present in the document.
  * Constraints: The list items must match regular expression `/^\\d+:\\d+:[A-Z0-9-]+$/`.
* `object_ids` - (Optional, List) Specifies a list of Object ids to filter.
* `os_types` - (Optional, List) Specifies the operating system types to filter objects on.
  * Constraints: Allowable list items are: `kLinux`, `kWindows`.
* `pagination_cookie` - (Optional, String) Specifies the pagination cookie with which subsequent parts of the response can be fetched.
* `protection_group_ids` - (Optional, List) Specifies a list of Protection Group ids to filter the objects. If specified, the objects protected by specified Protection Group ids will be returned.
* `protection_types` - (Optional, List) Specifies the protection type to filter objects.
  * Constraints: Allowable list items are: `kAgent`, `kNative`, `kSnapshotManager`, `kRDSSnapshotManager`, `kAuroraSnapshotManager`, `kAwsS3`, `kAwsRDSPostgresBackup`, `kAwsAuroraPostgres`, `kAwsRDSPostgres`, `kAzureSQL`, `kFile`, `kVolume`.
* `request_initiator_type` - (Optional, String) Specifies the type of request from UI, which is used for services like magneto to determine the priority of requests.
  * Constraints: Allowable values are: `UIUser`, `UIAuto`, `Helios`.
* `search_string` - (Optional, String) Specifies the search string to filter the objects. This search string will be applicable for objectnames. User can specify a wildcard character '*' as a suffix to a string where all object names are matched with the prefix string. For example, if vm1 and vm2 are the names of objects, user can specify vm* to list the objects. If not specified, then all the objects will be returned which will match other filtering criteria.
* `source_ids` - (Optional, List) Specifies a list of Protection Source object ids to filter the objects. If specified, the object which are present in those Sources will be returned.
* `source_uuids` - (Optional, List) Specifies a list of Protection Source object uuids to filter the objects. If specified, the object which are present in those Sources will be returned.
* `tag_categories` - (Optional, List) Specifies the tag category to filter the objects and snapshots.
  * Constraints: Allowable list items are: `Security`.
* `tag_names` - (Optional, List) Specifies the tag names to filter the tagged objects and snapshots.
* `tag_search_name` - (Optional, String) Specifies the tag name to filter the tagged objects and snapshots. User can specify a wildcard character '*' as a suffix to a string where all object's tag names are matched with the prefix string.
* `tag_sub_categories` - (Optional, List) Specifies the tag subcategory to filter the objects and snapshots.
  * Constraints: Allowable list items are: `Classification`, `Threats`, `Anomalies`, `Dspm`.
* `tag_types` - (Optional, List) Specifies the tag names to filter the tagged objects and snapshots.
  * Constraints: Allowable list items are: `System`, `Custom`, `ThirdParty`.
* `x_ibm_tenant_id` - (Required, String) Specifies the key to be used to encrypt the source credential. If includeSourceCredentials is set to true this key must be specified.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the Objects Search Result.
* `objects` - (List) Specifies the list of Objects.
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
	* `helios_tags` - (List) Specifies the helios tag information for the object.
	Nested schema for **helios_tags**:
		* `category` - (String) Specifies category of tag applied to the object.
		  * Constraints: Allowable values are: `Security`.
		* `name` - (String) Specifies name of tag applied to the object.
		* `sub_category` - (String) Specifies subCategory of tag applied to the object.
		  * Constraints: Allowable values are: `Classification`, `Threats`, `Anomalies`, `Dspm`.
		* `third_party_name` - (String) Specifies thirdPartyName of tag applied to the object.
		* `type` - (String) Specifies the type (ex custom, thirdparty, system) of tag applied to the object.
		  * Constraints: Allowable values are: `System`, `Custom`, `ThirdParty`.
		* `ui_color` - (String) Specifies the color of tag applied to the object.
		* `updated_time_usecs` - (Integer) Specifies update time of tag applied to the object.
		* `uuid` - (String) Specifies Uuid of tag applied to the object.
	* `id` - (Integer) Specifies object id.
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
	* `object_protection_infos` - (List) Specifies the object info on each cluster.
	Nested schema for **object_protection_infos**:
		* `cluster_id` - (Integer) Specifies the cluster id where this object belongs to.
		* `cluster_incarnation_id` - (Integer) Specifies the cluster incarnation id where this object belongs to.
		* `is_deleted` - (Boolean) Specifies whether the object is deleted. Deleted objects can't be protected but can be recovered or unprotected.
		* `last_run_status` - (String) Specifies the status of the object's last protection run.
		  * Constraints: Allowable values are: `Accepted`, `Running`, `Canceled`, `Canceling`, `Failed`, `Missed`, `Succeeded`, `SucceededWithWarning`, `OnHold`, `Finalizing`, `Skipped`, `LegalHold`, `Paused`.
		* `object_backup_configuration` - (List) Specifies a list of object protections.
		Nested schema for **object_backup_configuration**:
			* `last_archival_run_status` - (String) Specifies the status of last archival run.
			  * Constraints: Allowable values are: `Accepted`, `Running`, `Canceled`, `Canceling`, `Failed`, `Missed`, `Succeeded`, `SucceededWithWarning`, `OnHold`, `Finalizing`, `Skipped`, `LegalHold`, `Paused`.
			* `last_backup_run_status` - (String) Specifies the status of last local back up run.
			  * Constraints: Allowable values are: `Accepted`, `Running`, `Canceled`, `Canceling`, `Failed`, `Missed`, `Succeeded`, `SucceededWithWarning`, `OnHold`, `Finalizing`, `Skipped`, `LegalHold`, `Paused`.
			* `last_replication_run_status` - (String) Specifies the status of last replication run.
			  * Constraints: Allowable values are: `Accepted`, `Running`, `Canceled`, `Canceling`, `Failed`, `Missed`, `Succeeded`, `SucceededWithWarning`, `OnHold`, `Finalizing`, `Skipped`, `LegalHold`, `Paused`.
			* `last_run_sla_violated` - (Boolean) Specifies if the sla is violated in last run.
			* `policy_id` - (String) Specifies the policy id for this protection.
			* `policy_name` - (String) Specifies the policy name for this group.
		* `object_id` - (Integer) Specifies the object id.
		* `protection_groups` - (List) Specifies a list of protection groups protecting this object.
		Nested schema for **protection_groups**:
			* `id` - (String) Specifies the protection group id.
			* `last_archival_run_status` - (String) Specifies the status of last archival run.
			  * Constraints: Allowable values are: `Accepted`, `Running`, `Canceled`, `Canceling`, `Failed`, `Missed`, `Succeeded`, `SucceededWithWarning`, `OnHold`, `Finalizing`, `Skipped`, `LegalHold`.
			* `last_backup_run_status` - (String) Specifies the status of last local back up run.
			  * Constraints: Allowable values are: `Accepted`, `Running`, `Canceled`, `Canceling`, `Failed`, `Missed`, `Succeeded`, `SucceededWithWarning`, `OnHold`, `Finalizing`, `Skipped`, `LegalHold`.
			* `last_replication_run_status` - (String) Specifies the status of last replication run.
			  * Constraints: Allowable values are: `Accepted`, `Running`, `Canceled`, `Canceling`, `Failed`, `Missed`, `Succeeded`, `SucceededWithWarning`, `OnHold`, `Finalizing`, `Skipped`, `LegalHold`.
			* `last_run_sla_violated` - (Boolean) Specifies if the sla is violated in last run.
			* `name` - (String) Specifies the protection group name.
			* `policy_id` - (String) Specifies the policy id for this group.
			* `policy_name` - (String) Specifies the policy name for this group.
			* `protection_env_type` - (String) Specifies the protection type of the job if any.
			  * Constraints: Allowable values are: `kAgent`, `kNative`, `kSnapshotManager`, `kRDSSnapshotManager`, `kAuroraSnapshotManager`, `kAwsS3`, `kAwsRDSPostgresBackup`, `kAwsAuroraPostgres`, `kAwsRDSPostgres`, `kAzureSQL`, `kFile`, `kVolume`.
		* `region_id` - (String) Specifies the region id where this object belongs to.
		* `source_id` - (Integer) Specifies the source id.
		* `tenant_ids` - (List) List of Tenants the object belongs to.
		* `view_id` - (Integer) Specifies the view id for the object.
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
	* `secondary_ids` - (List) Specifies secondary IDs associated to the object.
	Nested schema for **secondary_ids**:
		* `name` - (String) Specifies name of the secondary ID for an object.
		* `value` - (String) Specifies value of the secondary ID for an object.
	* `sharepoint_site_summary` - (List) Specifies the common parameters for Sharepoint site objects.
	Nested schema for **sharepoint_site_summary**:
		* `site_web_url` - (String) Specifies the web url for the Sharepoint site.
	* `snapshot_tags` - (List) Specifies snapshot tags applied to the object.
	Nested schema for **snapshot_tags**:
		* `run_ids` - (List) Specifies runs the tags are applied to.
		* `tag_id` - (String) Specifies Id of tag applied to the object.
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
		* `source_name` - (String) Specifies registered source name to which object belongs.
		* `uuid` - (String) Specifies the uuid which is a unique identifier of the object.
		* `v_center_summary` - (List)
		Nested schema for **v_center_summary**:
			* `is_cloud_env` - (Boolean) Specifies that registered vCenter source is a VMC (VMware Cloud) environment or not.
		* `windows_cluster_summary` - (List)
		Nested schema for **windows_cluster_summary**:
			* `cluster_source_type` - (String) Specifies the type of cluster resource this source represents.
	* `source_name` - (String) Specifies registered source name to which object belongs.
	* `tagged_snapshots` - (List) Specifies the helios tagged snapshots (snapshots which are tagged by user or thirdparty in control plane) for the object.
	Nested schema for **tagged_snapshots**:
		* `cluster_id` - (Integer) Specifies the cluster Id of the tagged snapshot.
		* `cluster_incarnation_id` - (Integer) Specifies the clusterIncarnationId of the tagged snapshot.
		* `job_id` - (Integer) Specifies the jobId of the tagged snapshot.
		* `object_uuid` - (String) Specifies the object uuid of the tagged snapshot.
		* `run_start_time_usecs` - (Integer) Specifies the runStartTimeUsecs of the tagged snapshot.
		* `tags` - (List) Specifies tag applied to the object.
		Nested schema for **tags**:
			* `category` - (String) Specifies category of tag applied to the object.
			  * Constraints: Allowable values are: `Security`.
			* `name` - (String) Specifies name of tag applied to the object.
			* `sub_category` - (String) Specifies subCategory of tag applied to the object.
			  * Constraints: Allowable values are: `Classification`, `Threats`, `Anomalies`, `Dspm`.
			* `third_party_name` - (String) Specifies thirdPartyName of tag applied to the object.
			* `type` - (String) Specifies the type (ex custom, thirdparty, system) of tag applied to the object.
			  * Constraints: Allowable values are: `System`, `Custom`, `ThirdParty`.
			* `ui_color` - (String) Specifies the color of tag applied to the object.
			* `updated_time_usecs` - (Integer) Specifies update time of tag applied to the object.
			* `uuid` - (String) Specifies Uuid of tag applied to the object.
	* `tags` - (List) Specifies tag applied to the object.
	Nested schema for **tags**:
		* `tag_id` - (String) Specifies Id of tag applied to the object.
	* `uuid` - (String) Specifies the uuid which is a unique identifier of the object.
	* `v_center_summary` - (List)
	Nested schema for **v_center_summary**:
		* `is_cloud_env` - (Boolean) Specifies that registered vCenter source is a VMC (VMware Cloud) environment or not.
	* `windows_cluster_summary` - (List)
	Nested schema for **windows_cluster_summary**:
		* `cluster_source_type` - (String) Specifies the type of cluster resource this source represents.

