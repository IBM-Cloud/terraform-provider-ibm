---
layout: "ibm"
page_title: "IBM : ibm_backup_recovery_source_registration"
description: |-
  Manages backup_recovery_source_registration.
subcategory: "IBM Backup Recovery"
---

# ibm_backup_recovery_source_registration

Create, update, and delete backup_recovery_source_registrations with this resource.

## Example Usage

```hcl
resource "ibm_backup_recovery_source_registration" "backup_recovery_source_registration_instance" {
  advanced_configs {
		key = "key"
		value = "value"
  }

  connection_id = 1

  environment = "kPhysical"
    kubernetes_params {
		auto_protect_config {
			error_message = "error_message"
			is_default_auto_protected = true
			policy_id = "policy_id"
			protection_group_id = "protection_group_id"
			storage_domain_id = 1
		}
		client_private_key = "client_private_key"
		data_mover_image_location = "data_mover_image_location"
		datamover_service_type = "kNodePort"
		default_vlan_params {
			disable_vlan = true
			interface_name = "interface_name"
			vlan_id = 1
		}
		endpoint = "endpoint"
		init_container_image_location = "init_container_image_location"
		kubernetes_distribution = "kOpenshift"
		kubernetes_type = "kCluster"
		priority_class_name = "priority_class_name"
		resource_annotations {
			key = "key"
			value = "value"
		}
		resource_labels {
			key = "key"
			value = "value"
		}
		san_fields = [ "sanFields" ]
		service_annotations {
			key = "key"
			value = "value"
		}
		velero_aws_plugin_image_location = "velero_aws_plugin_image_location"
		velero_image_location = "velero_image_location"
		velero_openshift_plugin_image_location = "velero_openshift_plugin_image_location"
		vlan_info_vec {
			service_annotations {
				key = "key"
				value = "value"
			}
			vlan_params {
				disable_vlan = true
				interface_name = "interface_name"
				vlan_id = 1
			}
		}
  }
  physical_params {
		endpoint = "endpoint"
		force_register = true
		host_type = "kLinux"
		physical_type = "kGroup"
		applications = [ "kSQL" ]
  }
  x_ibm_tenant_id = "x_ibm_tenant_id"
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `advanced_configs` - (Optional, List) Specifies the advanced configuration for a protection source.
Nested schema for **advanced_configs**:
	* `key` - (Required, String) key.
	* `value` - (Required, String) value.
* `connection_id` - (Optional, Integer) Specifies the id of the connection from where this source is reachable. This should only be set for a source being registered by a tenant user. This field will be depricated in future. Use connections field.
* `connections` - (Optional, List) Specfies the list of connections for the source.
Nested schema for **connections**:
	* `connection_id` - (Optional, Integer) Specifies the id of the connection.
	* `connector_group_id` - (Optional, Integer) Specifies the connector group id of connector groups.
	* `data_source_connection_id` - (Optional, String) Specifies the id of the connection in string format.
	* `entity_id` - (Optional, Integer) Specifies the entity id of the source. The source can a non-root entity.
* `connector_group_id` - (Optional, Integer) Specifies the connector group id of connector groups.
* `data_source_connection_id` - (Optional, String) Specifies the id of the connection from where this source is reachable. This should only be set for a source being registered by a tenant user. Also, this is the 'string' of connectionId. This property was added to accommodate for ID values that exceed 2^53 - 1, which is the max value for which JS maintains precision.
* `environment` - (Required, String) Specifies the environment type of the Protection Source.
  * Constraints: Allowable values are: `kPhysical`, `kSQL`, `kKubernetes`.
* `kubernetes_params` - (Optional, List) Specifies the parameters to register a Kubernetes source.
Nested schema for **kubernetes_params**:
	* `auto_protect_config` - (Optional, List) Specifies the parameters to auto protect the source after registration.
	Nested schema for **auto_protect_config**:
		* `error_message` - (Optional, String) Specifies the error message in case source registration is successful but protection job creation fails.
		* `is_default_auto_protected` - (Required, Boolean) Specifies if entire source should be auto protected after registration. Default: False.
		* `policy_id` - (Required, String) Specifies the protection policy to auto protect the source with.
		* `protection_group_id` - (Optional, String) Specifies the protection group Id after it is successfully created.
		* `storage_domain_id` - (Optional, Integer) Specifies the storage domain id for the protection job.
	* `client_private_key` - (Required, String) Specifies the bearer token or private key of Kubernetes source.
	* `cohesity_dataprotect_plugin_image_location` - (Optional, String) Specifies the custom Cohesity Dataprotect plugin image location of the Kubernetes source.
	* `data_mover_image_location` - (Required, String) Specifies the datamover image location of Kubernetes source.
	* `datamover_service_type` - (Optional, String) Specifies the data mover service type of Kubernetes source.
	  * Constraints: Allowable values are: `kNodePort`, `kLoadBalancer`, `kClusterIp`.
	* `default_vlan_params` - (Optional, List) Specifies VLAN params associated with the backup/restore operation.
	Nested schema for **default_vlan_params**:
		* `disable_vlan` - (Optional, Boolean) If this is set to true, then even if VLANs are configured on the system, the partition VIPs will be used for the restore.
		* `interface_name` - (Optional, String) Interface group to use for backup/restore. If this is not specified, primary interface group for the cluster will be used.
		* `vlan_id` - (Optional, Integer) If this is set, then the Cohesity host name or the IP address associated with this VLAN is used for mounting Cohesity's view on the remote host.
	* `endpoint` - (Required, String) Specifies the endpoint of Kubernetes source.
	* `init_container_image_location` - (Optional, String) Specifies the initial container image location of Kubernetes source.
	* `kubernetes_distribution` - (Required, String) Specifies the distribution type of Kubernetes source.
	  * Constraints: Allowable values are: `kOpenshift`, `kMainline`, `kVMwareTanzu`, `kRancher`, `kEKS`, `kGKE`, `kAKS`, `kIKS`, `kROKS`.
	* `kubernetes_type` - (Optional, String) Specifies the type of kubernetes source.
	  * Constraints: Allowable values are: `kCluster`, `kNamespace`, `kService`, `kPVC`, `kPersistentVolumeClaim`, `kPersistentVolume`, `kLabel`.
	* `priority_class_name` - (Optional, String) Specifies the priority class name for cohesity resources.
	* `resource_annotations` - (Optional, List) Specifies resource annotations to be applied on cohesity resources.
	Nested schema for **resource_annotations**:
		* `key` - (Required, String) Specifies the label key.
		* `value` - (Optional, String) Specifies the label value.
	* `resource_labels` - (Optional, List) Specifies resource label to be applied on cohesity resources.
	Nested schema for **resource_labels**:
		* `key` - (Required, String) Specifies the label key.
		* `value` - (Optional, String) Specifies the label value.
	* `san_fields` - (Optional, List) Specifies the SAN field for agent certificate.
	* `service_annotations` - (Optional, List) Specifies the service annotation object of Kubernetes source.
	Nested schema for **service_annotations**:
		* `key` - (Optional, String) Specifies the service annotation key value.
		* `value` - (Optional, String) Specifies the service annotation value.
	* `velero_aws_plugin_image_location` - (Optional, String) Specifies the velero AWS plugin image location of the Kubernetes source.
	* `velero_image_location` - (Optional, String) Specifies the velero image location of the Kubernetes source.
	* `velero_openshift_plugin_image_location` - (Optional, String) Specifies the velero open shift plugin image for the Kubernetes source.
	* `vlan_info_vec` - (Optional, List) Specifies VLAN information provided during registration.
	Nested schema for **vlan_info_vec**:
		* `service_annotations` - (Optional, List) Specifies annotations to be put on services for IP allocation. Applicable only when service is of type LoadBalancer.
		Nested schema for **service_annotations**:
			* `key` - (Optional, String) Specifies the service annotation key value.
			* `value` - (Optional, String) Specifies the service annotation value.
		* `vlan_params` - (Optional, List) Specifies VLAN params associated with the backup/restore operation.
		Nested schema for **vlan_params**:
			* `disable_vlan` - (Optional, Boolean) If this is set to true, then even if VLANs are configured on the system, the partition VIPs will be used for the restore.
			* `interface_name` - (Optional, String) Interface group to use for backup/restore. If this is not specified, primary interface group for the cluster will be used.
			* `vlan_id` - (Optional, Integer) If this is set, then the Cohesity host name or the IP address associated with this VLAN is used for mounting Cohesity's view on the remote host.
* `name` - (Optional, String) The user specified name for this source.
* `endpoint_type` - (Optional, String) Backup Recovery Endpoint type. By default set to "public".
* `instance_id` - (Optional, String) Backup Recovery instance ID. If provided here along with region, the provider constructs the endpoint URL using them, which overrides any value set through environment variables or the `endpoints.json` file.
* `region` - (Optional, String) Backup Recovery region. If provided here along with instance_id, the provider constructs the endpoint URL using them, which overrides any value set through environment variables or the `endpoints.json` file.  
* `physical_params` - (Optional, List) Specifies parameters to register physical server.
Nested schema for **physical_params**:
	* `applications` - (Optional, List) Specifies the list of applications to be registered with Physical Source.
	  * Constraints: Allowable list items are: `kSQL`, `kOracle`.
	* `endpoint` - (Required, String) Specifies the endpoint IPaddress, URL or hostname of the physical host.
	* `force_register` - (Optional, Boolean) The agent running on a physical host will fail the registration if it is already registered as part of another cluster. By setting this option to true, agent can be forced to register with the current cluster.
	* `host_type` - (Optional, String) Specifies the type of host.
	  * Constraints: Allowable values are: `kLinux`, `kWindows`.
	* `physical_type` - (Optional, String) Specifies the type of physical server.
	  * Constraints: Allowable values are: `kGroup`, `kHost`, `kWindowsCluster`, `kOracleRACCluster`, `kOracleAPCluster`, `kUnixCluster`.
* `x_ibm_tenant_id` - (Required, String) Specifies the key to be used to encrypt the source credential. If includeSourceCredentials is set to true this key must be specified.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the backup_recovery_source_registration.
* `authentication_status` - (String) Specifies the status of the authentication during the registration of a Protection Source. 'Pending' indicates the authentication is in progress. 'Scheduled' indicates the authentication is scheduled. 'Finished' indicates the authentication is completed. 'RefreshInProgress' indicates the refresh is in progress.
  * Constraints: Allowable values are: `Pending`, `Scheduled`, `Finished`, `RefreshInProgress`.
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
* `registration_time_msecs` - (Integer) Specifies the time when the source was registered in milliseconds.
* `source_id` - (Integer) ID of top level source object discovered after the registration.
* `auto_proetction_group_id` - (String) Id of the protection group created using auto protect config
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


## Import

You can import the `ibm_backup_recovery_source_registration` resource by using `id`. Source Registration ID.  The ID is formed using tenantID and resourceId.
`id = <tenantId>::<source_id>`. 


#### Syntax
```
import {
	to = <ibm_backup_recovery_resource>
	id = "<tenantId>::<source_id>"
}
```

#### Example
```
resource "ibm_backup_recovery_source_registration" "terra_source_registration_2" {
  x_ibm_tenant_id = "jhxqx715r9/"
  environment = "kPhysical"
  connection_id = "6456"
  physical_params {
    endpoint = "172.26.1.1"
    host_type = "kLinux"
    physical_type = "kHost"
  }
}

import {
	to = ibm_backup_recovery_source_registration.terra_source_registration_1
	id = "jhxqx715r9/::3"
}
```

