---
layout: "ibm"
page_title: "IBM : ibm_backup_recovery_source_registration"
description: |-
  Get information about backup_recovery_source_registration
subcategory: "IBM Backup Recovery"
---

# ibm_backup_recovery_source_registration

Provides a read-only data source to retrieve information about a backup_recovery_source_registration. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_backup_recovery_source_registration" "backup_recovery_source_registration" {
	source_registration_id = ibm_backup_recovery_source_registration.backup_recovery_source_registration_instance.id
	x_ibm_tenant_id = ibm_backup_recovery_source_registration.backup_recovery_source_registration_instance.x_ibm_tenant_id
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `source_registration_id` - (Required, Forces new resource, Integer) Specifies the id of the Protection Source registration.
* `request_initiator_type` - (Optional, String) Specifies the type of request from UI, which is used for services like magneto to determine the priority of requests.
  * Constraints: Allowable values are: `UIUser`, `UIAuto`, `Helios`.
* `x_ibm_tenant_id` - (Required, String) Specifies the key to be used to encrypt the source credential. If includeSourceCredentials is set to true this key must be specified.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the backup_recovery_source_registration.
* `advanced_configs` - (List) Specifies the advanced configuration for a protection source.
Nested schema for **advanced_configs**:
	* `key` - (String) key.
	* `value` - (String) value.
* `authentication_status` - (String) Specifies the status of the authentication during the registration of a Protection Source. 'Pending' indicates the authentication is in progress. 'Scheduled' indicates the authentication is scheduled. 'Finished' indicates the authentication is completed. 'RefreshInProgress' indicates the refresh is in progress.
  * Constraints: Allowable values are: `Pending`, `Scheduled`, `Finished`, `RefreshInProgress`.
* `connection_id` - (Integer) Specifies the id of the connection from where this source is reachable. This should only be set for a source being registered by a tenant user. This field will be depricated in future. Use connections field.
* `connections` - (List) Specfies the list of connections for the source.
Nested schema for **connections**:
	* `connection_id` - (Integer) Specifies the id of the connection.
	* `connector_group_id` - (Integer) Specifies the connector group id of connector groups.
	* `data_source_connection_id` - (String) Specifies the id of the connection in string format.
	* `entity_id` - (Integer) Specifies the entity id of the source. The source can a non-root entity.
* `connector_group_id` - (Integer) Specifies the connector group id of connector groups.
* `data_source_connection_id` - (String) Specifies the id of the connection from where this source is reachable. This should only be set for a source being registered by a tenant user. Also, this is the 'string' of connectionId. This property was added to accommodate for ID values that exceed 2^53 - 1, which is the max value for which JS maintains precision.
* `environment` - (String) Specifies the environment type of the Protection Source.
  * Constraints: Allowable values are: `kPhysical`, `kSQL`.
* `external_metadata` - (List) Specifies the External metadata of an entity.
Nested schema for **external_metadata**:
	* `maintenance_mode_config` - (List) Specifies the entity metadata for maintenance mode.
	Nested schema for **maintenance_mode_config**:
		* `activation_time_intervals` - (List) Specifies the absolute intervals where the maintenance schedule is valid, i.e. maintenance_shedule is considered only for these time ranges. (For example, if there is one time range with [now_usecs, now_usecs + 10 days], the action will be done during the maintenance_schedule for the next 10 days.)The start time must be specified. The end time can be -1 which would denote an indefinite maintenance mode.
		Nested schema for **activation_time_intervals**:
			* `end_time_usecs` - (Integer) Specifies the end time of this time range.
			* `start_time_usecs` - (Integer) Specifies the start time of this time range.
		* `maintenance_schedule` - (List) Specifies a schedule for actions to be taken.
		Nested schema for **maintenance_schedule**:
			* `periodic_time_windows` - (List) Specifies the time range within the days of the week.
			Nested schema for **periodic_time_windows**:
				* `day_of_the_week` - (String) Specifies the week day.
				  * Constraints: Allowable values are: `Sunday`, `Monday`, `Tuesday`, `Wednesday`, `Thursday`, `Friday`, `Saturday`.
				* `end_time` - (List) Specifies the time in hours and minutes.
				Nested schema for **end_time**:
					* `hour` - (Integer) Specifies the hour of this time.
					* `minute` - (Integer) Specifies the minute of this time.
				* `start_time` - (List) Specifies the time in hours and minutes.
				Nested schema for **start_time**:
					* `hour` - (Integer) Specifies the hour of this time.
					* `minute` - (Integer) Specifies the minute of this time.
			* `schedule_type` - (String) Specifies the type of schedule for this ScheduleProto.
			  * Constraints: Allowable values are: `PeriodicTimeWindows`, `CustomIntervals`.
			* `time_ranges` - (List) Specifies the time ranges in usecs.
			Nested schema for **time_ranges**:
				* `end_time_usecs` - (Integer) Specifies the end time of this time range.
				* `start_time_usecs` - (Integer) Specifies the start time of this time range.
			* `timezone` - (String) Specifies the timezone of the user of this ScheduleProto. The timezones have unique names of the form 'Area/Location'.
		* `user_message` - (String) User provided message associated with this maintenance mode.
		* `workflow_intervention_spec_list` - (List) Specifies the type of intervention for different workflows when the source goes into maintenance mode.
		Nested schema for **workflow_intervention_spec_list**:
			* `intervention` - (String) Specifies the intervention type for ongoing tasks.
			  * Constraints: Allowable values are: `NoIntervention`, `Cancel`.
			* `workflow_type` - (String) Specifies the workflow type for which an intervention would be needed when maintenance mode begins.
			  * Constraints: Allowable values are: `BackupRun`.
* `last_refreshed_time_msecs` - (Integer) Specifies the time when the source was last refreshed in milliseconds.
* `name` - (String) The user specified name for this source.
* `physical_params` - (List) Specifies parameters to register physical server.
Nested schema for **physical_params**:
	* `applications` - (List) Specifies the list of applications to be registered with Physical Source.
	  * Constraints: Allowable list items are: `kSQL`, `kOracle`.
	* `endpoint` - (String) Specifies the endpoint IPaddress, URL or hostname of the physical host.
	* `force_register` - (Boolean) The agent running on a physical host will fail the registration if it is already registered as part of another cluster. By setting this option to true, agent can be forced to register with the current cluster.
	* `host_type` - (String) Specifies the type of host.
	  * Constraints: Allowable values are: `kLinux`, `kWindows`.
	* `physical_type` - (String) Specifies the type of physical server.
	  * Constraints: Allowable values are: `kGroup`, `kHost`, `kWindowsCluster`, `kOracleRACCluster`, `kOracleAPCluster`, `kUnixCluster`.
* `registration_time_msecs` - (Integer) Specifies the time when the source was registered in milliseconds.
* `source_id` - (Integer) ID of top level source object discovered after the registration.
* `source_info` - (List) Specifies information about an object.
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

