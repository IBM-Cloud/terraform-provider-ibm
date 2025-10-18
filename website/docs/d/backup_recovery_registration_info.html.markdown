---
layout: "ibm"
page_title: "IBM : ibm_backup_recovery_registration_info"
description: |-
  Get information about backup_recovery_registration_info
subcategory: "IBM Backup Recovery API"
---

# ibm_backup_recovery_registration_info

Provides a read-only data source to retrieve information about a backup_recovery_registration_info. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_backup_recovery_registration_info" "backup_recovery_registration_info" {
	x_ibm_tenant_id = "tenantId"
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `all_under_hierarchy` - (Optional, Boolean) AllUnderHierarchy specifies if objects of all the tenants under the hierarchy of the logged in user's organization should be returned.
* `encryption_key` - (Optional, String) Key to be used to encrypt the source credential. If include_source_credentials is set to true this key must be specified.
* `environments` - (Optional, List) Return only Protection Sources that match the passed in environment type such as 'kVMware', 'kSQL', 'kView' 'kPhysical', 'kPuppeteer', 'kPure', 'kNetapp', 'kGenericNas', 'kHyperV', 'kAcropolis', or 'kAzure'. For example, set this parameter to 'kVMware' to only return the Sources (and their Object subtrees) found in the 'kVMware' (VMware vCenter Server) environment. NOTE: 'kPuppeteer' refers to Cohesity's Remote Adapter.
  * Constraints: Allowable list items are: `kVMware`, `kSQL`, `kView`, `kPuppeteer`, `kPhysical`, `kPure`, `kNetapp`, `kGenericNas`, `kHyperV`, `kAcropolis`, `kAzure`, `kPhysicalFiles`, `kIsilon`, `kGPFS`, `kKVM`, `kAWS`, `kExchange`, `kHyperVVSS`, `kOracle`, `kGCP`, `kFlashBlade`, `kAWSNative`, `kVCD`, `kO365`, `kO365Outlook`, `kHyperFlex`, `kGCPNative`, `kKubernetes`, `kCassandra`, `kMongoDB`, `kCouchbase`, `kHdfs`, `kHive`, `kHBase`, `kUDA`, `kAwsS3`.
* `ids` - (Optional, List) Return only the registered root nodes whose Ids are given in the list.
* `include_applications_tree_info` - (Optional, Boolean) Specifies whether to return applications tree info or not.
* `include_entity_permission_info` - (Optional, Boolean) If specified, then a list of entities with permissions assigned to them are returned.
* `include_external_metadata` - (Optional, Boolean) Specifies if entity external metadata should be included within the response to get entity hierarchy call.
* `include_source_credentials` - (Optional, Boolean) If specified, then crednetial for the registered sources will be included. Credential is first encrypted with internal key and then reencrypted with user supplied 'encryption_key'.
* `maintenance_status` - (Optional, String) Specifies the maintenance status of a source 'UnderMaintenance' indicates the source is currently under maintenance. 'ScheduledMaintenance' indicates the source is scheduled for maintenance. 'NotConfigured' indicates maintenance is not configured on the source.
  * Constraints: Allowable values are: `UnderMaintenance`, `ScheduledMaintenance`, `NotConfigured`.
* `prune_non_critical_info` - (Optional, Boolean) Specifies whether to prune non critical info within entities. Incase of VMs, virtual disk information will be pruned. Incase of Office365, metadata about user entities will be pruned. This can be used to limit the size of the response by caller.
* `request_initiator_type` - (Optional, String) Specifies the type of the request. Possible values are UIUser and UIAuto, which means the request is triggered by user or is an auto refresh request. Services like magneto will use this to determine the priority of the requests, so that it can more intelligently handle overload situations by prioritizing higher priority requests.
* `sids` - (Optional, List) Filter the registered root nodes for the sids given in the list.
* `tenant_ids` - (Optional, List) TenantIds contains ids of the tenants for which objects are to be returned.
* `use_cached_data` - (Optional, Boolean) Specifies whether we can serve the GET request to the read replica cache. setting this to true ensures that the API request is served to the read replica. setting this to false will serve the request to the master.
* `x_ibm_tenant_id` - (Required, String) Specifies the key to be used to encrypt the source credential. If includeSourceCredentials is set to true this key must be specified.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the backup_recovery_registration_info.
* `root_nodes` - (List) Specifies the registration, protection and permission information of either all or a subset of registered Protection Sources matching the filter parameters. overrideDescription: true.
Nested schema for **root_nodes**:
	* `applications` - (List) Array of applications hierarchy registered on this node. Specifies the application type and the list of instances of the application objects. For example for SQL Server, this list provides the SQL Server instances running on a VM or a Physical Server.
	Nested schema for **applications**:
		* `application_tree_info` - (List) Application Server and the subtrees below them. Specifies the application subtree used to store additional application level Objects. Different environments use the subtree to store application level information. For example for SQL Server, this subtree stores the SQL Server instances running on a VM or a Physical Server. overrideDescription: true.
		Nested schema for **application_tree_info**:
			* `connection_id` - (Integer) Specifies the connection id of the tenant.
			* `connector_group_id` - (Integer) Specifies the connector group id of the connector groups.
			* `custom_name` - (String) Specifies the user provided custom name of the Protection Source.
			* `environment` - (String) Specifies the environment (such as 'kVMware' or 'kSQL') where the Protection Source exists. Depending on the environment, one of the following Protection Sources are initialized.
			  * Constraints: Allowable values are: `kPhysical`, `kPhysicalFiles`, `kSQL`, `kAgent`.
			* `id` - (Integer) Specifies an id of the Protection Source.
			* `kubernetes_protection_source` - (List) Specifies a Protection Source in Kubernetes environment.
			Nested schema for **kubernetes_protection_source**:
				* `datamover_image_location` - (String) Specifies the location of Datamover image in private registry.
				* `datamover_service_type` - (Integer) Specifies Type of service to be deployed for communication with DataMover pods. Currently, LoadBalancer and NodePort are supported. [default = kNodePort].
				* `datamover_upgradability` - (Integer) Specifies if the deployed Datamover image needs to be upgraded for this kubernetes entity.
				* `default_vlan_params` - (List) Specifies VLAN parameters for the restore operation.
				Nested schema for **default_vlan_params**:
					* `disable_vlan` - (Boolean) Specifies whether to use the VIPs even when VLANs are configured on the Cluster. If configured, VLAN IP addresses are used by default. If VLANs are not configured, this flag is ignored. Set this flag to true to force using the partition VIPs when VLANs are configured on the Cluster.
					* `interface_name` - (String) Specifies the physical interface group name to use for mounting Cohesity's view on the remote host. If specified, Cohesity hostname or the IP address on this VLAN is used.
					* `vlan` - (Integer) Specifies the VLAN to use for mounting Cohesity's view on the remote host. If specified, Cohesity hostname or the IP address on this VLAN is used.
				* `description` - (String) Specifies an optional description of the object.
				* `distribution` - (String) Specifies the type of the entity in a Kubernetes environment. Determines the K8s distribution. kIKS, kROKS.
				  * Constraints: Allowable values are: `kMainline`, `kOpenshift`, `kRancher`, `kEKS`, `kGKE`, `kAKS`, `kVMwareTanzu`.
				* `init_container_image_location` - (String) Specifies the location of the image for init containers.
				* `label_attributes` - (List) Specifies the list of label attributes of this source.
				Nested schema for **label_attributes**:
					* `id` - (Integer) Specifies the Cohesity id of the K8s label.
					* `name` - (String) Specifies the appended key and value of the K8s label.
					* `uuid` - (String) Specifies Kubernetes Unique Identifier (UUID) of the K8s label.
				* `name` - (String) Specifies a unique name of the Protection Source.
				* `priority_class_name` - (String) Specifies the pritority class name during registration.
				* `resource_annotation_list` - (List) Specifies resource Annotations information provided during registration.
				Nested schema for **resource_annotation_list**:
					* `key` - (String) Key for label.
					* `value` - (String) Value for label.
				* `resource_label_list` - (List) Specifies resource labels information provided during registration.
				Nested schema for **resource_label_list**:
					* `key` - (String) Key for label.
					* `value` - (String) Value for label.
				* `san_field` - (List) Specifies the SAN field for agent certificate.
				* `service_annotations` - (List) Specifies annotations to be put on services for IP allocation. Applicable only when service is of type LoadBalancer.
				Nested schema for **service_annotations**:
					* `key` - (String)
					* `value` - (String)
				* `storage_class` - (List) Specifies storage class information of source.
				Nested schema for **storage_class**:
					* `name` - (String) Specifies name of storage class.
					* `provisioner` - (String) specifies provisioner of storage class.
				* `type` - (String) Specifies the type of the entity in a Kubernetes environment. Specifies the type of a Kubernetes Protection Source. 'kCluster' indicates a Kubernetes Cluster. 'kNamespace' indicates a namespace in a Kubernetes Cluster. 'kService' indicates a service running on a Kubernetes Cluster.
				  * Constraints: Allowable values are: `kCluster`, `kNamespace`, `kService`.
				* `uuid` - (String) Specifies the UUID of the object.
				* `velero_aws_plugin_image_location` - (String) Specifies the location of Velero AWS plugin image in private registry.
				* `velero_image_location` - (String) Specifies the location of Velero image in private registry.
				* `velero_openshift_plugin_image_location` - (String) Specifies the location of the image for openshift plugin container.
				* `velero_upgradability` - (String) Specifies if the deployed Velero image needs to be upgraded for this kubernetes entity.
				* `vlan_info_vec` - (List) Specifies VLAN information provided during registration.
				Nested schema for **vlan_info_vec**:
					* `service_annotations` - (List) Specifies annotations to be put on services for IP allocation. Applicable only when service is of type LoadBalancer.
					Nested schema for **service_annotations**:
						* `key` - (String) Specifies the service annotation key value.
						* `value` - (String) Specifies the service annotation value.
					* `vlan_params` - (List) Specifies VLAN params associated with the backup/restore operation.
					Nested schema for **vlan_params**:
						* `disable_vlan` - (Boolean) If this is set to true, then even if VLANs are configured on the system, the partition VIPs will be used for the restore.
						* `interface_name` - (String) Interface group to use for backup/restore. If this is not specified, primary interface group for the cluster will be used.
						* `vlan_id` - (Integer) If this is set, then the Cohesity host name or the IP address associated with this VLAN is used for mounting Cohesity's view on the remote host.
			* `name` - (String) Specifies a name of the Protection Source.
			* `parent_id` - (Integer) Specifies an id of the parent of the Protection Source.
			* `physical_protection_source` - (List) Specifies a Protection Source in a Physical environment.
			Nested schema for **physical_protection_source**:
				* `agents` - (List) Specifiles the agents running on the Physical Protection Source and the status information.
				Nested schema for **agents**:
					* `cbmr_version` - (String) Specifies the version if Cristie BMR product is installed on the host.
					* `file_cbt_info` - (List) CBT version and service state info.
					Nested schema for **file_cbt_info**:
						* `file_version` - (List) Subcomponent version. The interpretation of the version is based on operating system.
						Nested schema for **file_version**:
							* `build_ver` - (Float)
							* `major_ver` - (Float)
							* `minor_ver` - (Float)
							* `revision_num` - (Float)
						* `is_installed` - (Boolean) Indicates whether the cbt driver is installed.
						* `reboot_status` - (String) Indicates whether host is rebooted post VolCBT installation.
						  * Constraints: Allowable values are: `kRebooted`, `kNeedsReboot`, `kInternalError`.
						* `service_state` - (List) Structure to Hold Service Status.
						Nested schema for **service_state**:
							* `name` - (String)
							* `state` - (String)
					* `host_type` - (String) Specifies the host type where the agent is running. This is only set for persistent agents.
					  * Constraints: Allowable values are: `kLinux`, `kWindows`, `kAix`, `kSolaris`, `kSapHana`, `kSapOracle`, `kCockroachDB`, `kMySQL`, `kOther`, `kSapSybase`, `kSapMaxDB`, `kSapSybaseIQ`, `kDB2`, `kSapASE`, `kMariaDB`, `kPostgreSQL`, `kVOS`, `kHPUX`.
					* `id` - (Integer) Specifies the agent's id.
					* `name` - (String) Specifies the agent's name.
					* `oracle_multi_node_channel_supported` - (Boolean) Specifies whether oracle multi node multi channel is supported or not.
					* `registration_info` - (List) Specifies information about a registered Source.
					Nested schema for **registration_info**:
						* `access_info` - (List) Specifies the parameters required to establish a connection with a particular environment.
						Nested schema for **access_info**:
							* `connection_id` - (Integer) ID of the Bifrost (HyX or Rigel) network realm (i.e. a connection) associated with the source.
							* `connector_group_id` - (Integer) Specifies the Id of the connector group. Each connector group is collection of Rigel/hyx. Each entity will be tagged with connector group id.
							* `endpoint` - (String) Specify an IP address or URL of the environment. (such as the IP address of the vCenter Server for a VMware environment).
							* `environment` - (String) Specifies the environment like VMware, SQL, where the Protection Source exists. Supported environment types such as 'kView', 'kSQL', 'kVMware', etc.
							  * Constraints: Allowable values are: `kPhysical`, `kPhysicalFiles`, `kSQL`, `kAgent`.
							* `id` - (Integer) Specifies a Unique id that is generated when the Source is registered. This is a convenience field that is used to maintain an index to different connection params.
							* `version` - (Integer) Version is updated each time the connector parameters are updated. This is used to discard older connector parameters.
						* `allowed_ip_addresses` - (List) Specifies the list of IP Addresses on the registered source to be exclusively allowed for doing any type of IO operations.
						* `authentication_error_message` - (String) Specifies an authentication error message. This indicates the given credentials are rejected and the registration of the source is not successful.
						* `authentication_status` - (String) Specifies the status of the authenticating to the Protection Source when registering it with Cohesity Cluster.
						  * Constraints: Allowable values are: `kPending`, `kScheduled`, `kFinished`, `kRefreshInProgress`.
						* `blacklisted_ip_addresses` - (List) This field is deprecated. Use DeniedIpAddresses instead.
						* `denied_ip_addresses` - (List) Specifies the list of IP Addresses on the registered source to be denied for doing any type of IO operations.
						* `environments` - (List) Specifies a list of applications environment that are registered with this Protection Source such as 'kSQL'. Supported environment types such as 'kView', 'kSQL', 'kVMware', etc.
						  * Constraints: Allowable list items are: `kPhysical`, `kPhysicalFiles`, `kSQL`, `kAgent`.
						* `is_db_authenticated` - (Boolean) Specifies if application entity dbAuthenticated or not.
						* `is_storage_array_snapshot_enabled` - (Boolean) Specifies if this source entity has enabled storage array snapshot or not.
						* `link_vms_across_vcenter` - (Boolean) Specifies if the VM linking feature is enabled for this VCenter This means that VMs present in this VCenter which earlier belonged to some other VCenter(also registerd on same cluster) and were migrated, will be linked during EH refresh. This will enable preserving snapshot chains for migrated VMs.
						* `minimum_free_space_gb` - (Integer) Specifies the minimum free space in GiB of the space expected to be available on the datastore where the virtual disks of the VM being backed up. If the amount of free space(in GiB) is lower than the value given by this field, backup will be aborted. Note that this field is applicable only to 'kVMware' type of environments.
						* `minimum_free_space_percent` - (Integer) Specifies the minimum free space in percentage of the space expected to be available on the datastore where the virtual disks of the VM being backed up. If the amount of free space(in percentage) is lower than the value given by this field, backup will be aborted. Note that this field is applicable only to 'kVMware' type of environments.
						* `password` - (String) Specifies password of the username to access the target source.
						* `physical_params` - (List) Specifies the parameters required to register Application Servers running in a Protection Source specific to a physical adapter.
						Nested schema for **physical_params**:
							* `applications` - (List) Specifies the types of applications such as 'kSQL', 'kExchange', 'kAD' running on the Protection Source. Supported environment types such as 'kView', 'kSQL', 'kVMware', etc.
							  * Constraints: Allowable list items are: `kSQL`, `kOracle`.
							* `password` - (String) Specifies password of the username to access the target source.
							* `throttling_config` - (List) Specifies the source side throttling configuration.
							Nested schema for **throttling_config**:
								* `cpu_throttling_config` - (List) Specifies the Throttling Configuration Parameters.
								Nested schema for **cpu_throttling_config**:
									* `fixed_threshold` - (Integer) Fixed baseline threshold for throttling. This is mandatory for any other throttling type than kNoThrottling.
									* `pattern_type` - (String) Type of the throttling pattern. 'kNoThrottling' indicates that throttling is not in force. 'kBaseThrottling' indicates indicates a constant base level throttling. 'kFixed' indicates a constant base level throttling.
									  * Constraints: Allowable values are: `kNoThrottling`, `kBaseThrottling`, `kFixed`.
									* `throttling_windows` - (List)
									Nested schema for **throttling_windows**:
										* `day_time_window` - (List) Specifies the Day Time Window Parameters.
										Nested schema for **day_time_window**:
											* `end_time` - (List) Specifies the Day Time Parameters.
											Nested schema for **end_time**:
												* `day` - (String) Specifies the day of the week (such as 'kMonday') for scheduling throttling. Specifies a day in a week such as 'kSunday', 'kMonday', etc.
												  * Constraints: Allowable values are: `kSunday`, `kMonday`, `kTuesday`, `kWednesday`, `kThursday`, `kFriday`, `kSaturday`.
												* `time` - (List) Specifies the time in hours and minutes.
												Nested schema for **time**:
													* `hour` - (Integer) Specifies the hour of this time.
													* `minute` - (Integer) Specifies the minute of this time.
											* `start_time` - (List) Specifies the Day Time Parameters.
											Nested schema for **start_time**:
												* `day` - (String) Specifies the day of the week (such as 'kMonday') for scheduling throttling. Specifies a day in a week such as 'kSunday', 'kMonday', etc.
												  * Constraints: Allowable values are: `kSunday`, `kMonday`, `kTuesday`, `kWednesday`, `kThursday`, `kFriday`, `kSaturday`.
												* `time` - (List) Specifies the time in hours and minutes.
												Nested schema for **time**:
													* `hour` - (Integer) Specifies the hour of this time.
													* `minute` - (Integer) Specifies the minute of this time.
										* `threshold` - (Integer) Throttling threshold applicable in the window.
								* `network_throttling_config` - (List) Specifies the Throttling Configuration Parameters.
								Nested schema for **network_throttling_config**:
									* `fixed_threshold` - (Integer) Fixed baseline threshold for throttling. This is mandatory for any other throttling type than kNoThrottling.
									* `pattern_type` - (String) Type of the throttling pattern. 'kNoThrottling' indicates that throttling is not in force. 'kBaseThrottling' indicates indicates a constant base level throttling. 'kFixed' indicates a constant base level throttling.
									  * Constraints: Allowable values are: `kNoThrottling`, `kBaseThrottling`, `kFixed`.
									* `throttling_windows` - (List)
									Nested schema for **throttling_windows**:
										* `day_time_window` - (List) Specifies the Day Time Window Parameters.
										Nested schema for **day_time_window**:
											* `end_time` - (List) Specifies the Day Time Parameters.
											Nested schema for **end_time**:
												* `day` - (String) Specifies the day of the week (such as 'kMonday') for scheduling throttling. Specifies a day in a week such as 'kSunday', 'kMonday', etc.
												  * Constraints: Allowable values are: `kSunday`, `kMonday`, `kTuesday`, `kWednesday`, `kThursday`, `kFriday`, `kSaturday`.
												* `time` - (List) Specifies the time in hours and minutes.
												Nested schema for **time**:
													* `hour` - (Integer) Specifies the hour of this time.
													* `minute` - (Integer) Specifies the minute of this time.
											* `start_time` - (List) Specifies the Day Time Parameters.
											Nested schema for **start_time**:
												* `day` - (String) Specifies the day of the week (such as 'kMonday') for scheduling throttling. Specifies a day in a week such as 'kSunday', 'kMonday', etc.
												  * Constraints: Allowable values are: `kSunday`, `kMonday`, `kTuesday`, `kWednesday`, `kThursday`, `kFriday`, `kSaturday`.
												* `time` - (List) Specifies the time in hours and minutes.
												Nested schema for **time**:
													* `hour` - (Integer) Specifies the hour of this time.
													* `minute` - (Integer) Specifies the minute of this time.
										* `threshold` - (Integer) Throttling threshold applicable in the window.
							* `username` - (String) Specifies username to access the target source.
						* `progress_monitor_path` - (String) Captures the current progress and pulse details w.r.t to either the registration or refresh.
						* `refresh_error_message` - (String) Specifies a message if there was any error encountered during the last rebuild of the Protection Source tree. If there was no error during the last rebuild, this field is reset.
						* `refresh_time_usecs` - (Integer) Specifies the Unix epoch time (in microseconds) when the Protection Source tree was most recently fetched and built.
						* `registered_apps_info` - (List) Specifies information of the applications registered on this protection source.
						Nested schema for **registered_apps_info**:
							* `authentication_error_message` - (String) pecifies an authentication error message. This indicates the given credentials are rejected and the registration of the application is not successful.
							* `authentication_status` - (String) Specifies the status of authenticating to the Protection Source when registering this application with Cohesity Cluster. If the status is 'kFinished' and there is no error, registration is successful. Specifies the status of the authentication during the registration of a Protection Source. 'kPending' indicates the authentication is in progress. 'kScheduled' indicates the authentication is scheduled. 'kFinished' indicates the authentication is completed. 'kRefreshInProgress' indicates the refresh is in progress.
							  * Constraints: Allowable values are: `kPending`, `kScheduled`, `kFinished`, `kRefreshInProgress`.
							* `environment` - (String) Specifies the application environment. Supported environment types such as 'kView', 'kSQL', 'kVMware', etc.
							  * Constraints: Allowable values are: `kPhysical`, `kPhysicalFiles`, `kSQL`, `kAgent`, `kVMware`, `kHyperV`, `kPure`, `kNimble`, `kView`, `kPuppeteer`.
							* `host_settings_check_results` - (List)
							Nested schema for **host_settings_check_results**:
								* `check_type` - (String) Specifies the type of the check internally performed. Specifies the type of the host check performed internally. 'kIsAgentPortAccessible' indicates the check for agent port access. 'kIsAgentRunning' indicates the status for the Cohesity agent service. 'kIsSQLWriterRunning' indicates the status for SQLWriter service. 'kAreSQLInstancesRunning' indicates the run status for all the SQL instances in the host. 'kCheckServiceLoginsConfig' checks the privileges and sysadmin status of the logins used by the SQL instance services, Cohesity agent service and the SQLWriter service. 'kCheckSQLFCIVIP' checks whether the SQL FCI is registered with a valid VIP or FQDN. 'kCheckSQLDiskSpace' checks whether volumes containing SQL DBs have at least 10% free space.
								  * Constraints: Allowable values are: `kIsAgentPortAccessible`, `kIsAgentRunning`, `kIsSQLWriterRunning`, `kAreSQLInstancesRunning`, `kCheckServiceLoginsConfig`, `kCheckSQLFCIVIP`, `kCheckSQLDiskSpace`.
								* `result_type` - (String) Specifies the type of the result returned after performing the internal host check. Specifies the type of the host check result performed internally. 'kPass' indicates that the respective check was successful. 'kFail' indicates that the respective check failed as some mandatory setting is not met 'kWarning' indicates that the respective check has warning as certain non-mandatory setting is not met.
								  * Constraints: Allowable values are: `kPass`, `kFail`, `kWarning`.
								* `user_message` - (String) Specifies a descriptive message for failed/warning types.
							* `refresh_error_message` - (String) Specifies a message if there was any error encountered during the last rebuild of the application tree. If there was no error during the last rebuild, this field is reset.
						* `registration_time_usecs` - (Integer) Specifies the Unix epoch time (in microseconds) when the Protection Source was registered.
						* `subnets` - (List) Specifies the list of subnets added during creation or updation of vmare source. Currently, this field will only be populated in case of VMware registration.
						Nested schema for **subnets**:
							* `component` - (String) Component that has reserved the subnet.
							* `description` - (String) Description of the subnet.
							* `id` - (Float) ID of the subnet.
							* `ip` - (String) Specifies either an IPv6 address or an IPv4 address.
							* `netmask_bits` - (Float) netmaskBits.
							* `netmask_ip4` - (String) Specifies the netmask using an IP4 address. The netmask can only be set using netmaskIp4 if the IP address is an IPv4 address.
							* `nfs_access` - (String) Component that has reserved the subnet.
							  * Constraints: Allowable values are: `kDisabled`, `kReadOnly`, `kReadWrite`.
							* `nfs_all_squash` - (Boolean) Specifies whether all clients from this subnet can map view with view_all_squash_uid/view_all_squash_gid configured in the view.
							* `nfs_root_squash` - (Boolean) Specifies whether clients from this subnet can mount as root on NFS.
							* `s3_access` - (String) Specifies whether clients from this subnet can access using S3 protocol. Protocol access level. 'kDisabled' indicates Protocol access level 'Disabled' 'kReadOnly' indicates Protocol access level 'ReadOnly' 'kReadWrite' indicates Protocol access level 'ReadWrite'.
							  * Constraints: Allowable values are: `kDisabled`, `kReadOnly`, `kReadWrite`.
							* `smb_access` - (String) Specifies whether clients from this subnet can mount using SMB protocol. Protocol access level. 'kDisabled' indicates Protocol access level 'Disabled' 'kReadOnly' indicates Protocol access level 'ReadOnly' 'kReadWrite' indicates Protocol access level 'ReadWrite'.
							  * Constraints: Allowable values are: `kDisabled`, `kReadOnly`, `kReadWrite`.
							* `tenant_id` - (String) Specifies the unique id of the tenant.
						* `throttling_policy` - (List) Specifies the throttling policy for a registered Protection Source.
						Nested schema for **throttling_policy**:
							* `enforce_max_streams` - (Boolean) Specifies whether datastore streams are configured for all datastores that are part of the registered entity. If set to true, number of streams from Cohesity cluster to the registered entity will be limited to the value set for maxConcurrentStreams. If not set or set to false, there is no max limit for the number of concurrent streams.
							* `enforce_registered_source_max_backups` - (Boolean) Specifies whether no. of backups are configured for the registered entity. If set to true, number of backups made by Cohesity cluster in the registered entity will be limited to the value set for RegisteredSourceMaxConcurrentBackups. If not set or set to false, there is no max limit for the number of concurrent backups.
							* `is_enabled` - (Boolean) Indicates whether read operations to the datastores, which are part of the registered Protection Source, are throttled.
							* `latency_thresholds` - (List) Specifies latency thresholds that trigger throttling for all datastores found in the registered Protection Source or specific to one datastore.
							Nested schema for **latency_thresholds**:
								* `active_task_msecs` - (Integer) If the latency of a datastore is above this value, existing backup tasks using the datastore are throttled.
								* `new_task_msecs` - (Integer) If the latency of a datastore is above this value, then new backup tasks using the datastore will not be started.
							* `max_concurrent_streams` - (Float) Specifies the limit on the number of streams Cohesity cluster will make concurrently to the datastores of the registered entity. This limit is enforced only when the flag enforceMaxStreams is set to true.
							* `nas_source_params` - (List) Specifies the NAS specific source throttling parameters during source registration or during backup of the source.
							Nested schema for **nas_source_params**:
								* `max_parallel_metadata_fetch_full_percentage` - (Float) Specifies the percentage value of maximum concurrent metadata to be fetched during full backup of the source.
								* `max_parallel_metadata_fetch_incremental_percentage` - (Float) Specifies the percentage value of maximum concurrent metadata to be fetched during incremental backup of the source.
								* `max_parallel_read_write_full_percentage` - (Float) Specifies the percentage value of maximum concurrent IO during full backup of the source.
								* `max_parallel_read_write_incremental_percentage` - (Float) Specifies the percentage value of maximum concurrent IO during incremental backup of the source.
							* `registered_source_max_concurrent_backups` - (Float) Specifies the limit on the number of backups Cohesity cluster will make concurrently to the registered entity. This limit is enforced only when the flag enforceRegisteredSourceMaxBackups is set to true.
							* `storage_array_snapshot_config` - (List) Specifies Storage Array Snapshot Configuration.
							Nested schema for **storage_array_snapshot_config**:
								* `is_max_snapshots_config_enabled` - (Boolean) Specifies if the storage array snapshot max snapshots config is enabled or not.
								* `is_max_space_config_enabled` - (Boolean) Specifies if the storage array snapshot max space config is enabled or not.
								* `storage_array_snapshot_max_space_config` - (List) Specifies Storage Array Snapshot Max Space Config.
								Nested schema for **storage_array_snapshot_max_space_config**:
									* `max_snapshot_space_percentage` - (Float) Max number of storage snapshots allowed per volume/lun.
								* `storage_array_snapshot_throttling_policies` - (List) Specifies throttling policies configured for individual volume/lun.
								Nested schema for **storage_array_snapshot_throttling_policies**:
									* `id` - (Integer) Specifies the volume id of the storage array snapshot config.
									* `is_max_snapshots_config_enabled` - (Boolean) Specifies if the storage array snapshot max snapshots config is enabled or not.
									* `is_max_space_config_enabled` - (Boolean) Specifies if the storage array snapshot max space config is enabled or not.
									* `max_snapshot_config` - (List) Specifies Storage Array Snapshot Max Snapshots Config.
									Nested schema for **max_snapshot_config**:
										* `max_snapshots` - (Float) Max number of storage snapshots allowed per volume/lun.
									* `max_space_config` - (List) Specifies Storage Array Snapshot Max Space Config.
									Nested schema for **max_space_config**:
										* `max_snapshot_space_percentage` - (Float) Max number of storage snapshots allowed per volume/lun.
						* `throttling_policy_overrides` - (List)
						Nested schema for **throttling_policy_overrides**:
							* `datastore_id` - (Integer) Specifies the Protection Source id of the Datastore.
							* `datastore_name` - (String) Specifies the display name of the Datastore.
							* `throttling_policy` - (List) Specifies the throttling policy for a registered Protection Source.
							Nested schema for **throttling_policy**:
								* `enforce_max_streams` - (Boolean) Specifies whether datastore streams are configured for all datastores that are part of the registered entity. If set to true, number of streams from Cohesity cluster to the registered entity will be limited to the value set for maxConcurrentStreams. If not set or set to false, there is no max limit for the number of concurrent streams.
								* `enforce_registered_source_max_backups` - (Boolean) Specifies whether no. of backups are configured for the registered entity. If set to true, number of backups made by Cohesity cluster in the registered entity will be limited to the value set for RegisteredSourceMaxConcurrentBackups. If not set or set to false, there is no max limit for the number of concurrent backups.
								* `is_enabled` - (Boolean) Indicates whether read operations to the datastores, which are part of the registered Protection Source, are throttled.
								* `latency_thresholds` - (List) Specifies latency thresholds that trigger throttling for all datastores found in the registered Protection Source or specific to one datastore.
								Nested schema for **latency_thresholds**:
									* `active_task_msecs` - (Integer) If the latency of a datastore is above this value, existing backup tasks using the datastore are throttled.
									* `new_task_msecs` - (Integer) If the latency of a datastore is above this value, then new backup tasks using the datastore will not be started.
								* `max_concurrent_streams` - (Float) Specifies the limit on the number of streams Cohesity cluster will make concurrently to the datastores of the registered entity. This limit is enforced only when the flag enforceMaxStreams is set to true.
								* `nas_source_params` - (List) Specifies the NAS specific source throttling parameters during source registration or during backup of the source.
								Nested schema for **nas_source_params**:
									* `max_parallel_metadata_fetch_full_percentage` - (Float) Specifies the percentage value of maximum concurrent metadata to be fetched during full backup of the source.
									* `max_parallel_metadata_fetch_incremental_percentage` - (Float) Specifies the percentage value of maximum concurrent metadata to be fetched during incremental backup of the source.
									* `max_parallel_read_write_full_percentage` - (Float) Specifies the percentage value of maximum concurrent IO during full backup of the source.
									* `max_parallel_read_write_incremental_percentage` - (Float) Specifies the percentage value of maximum concurrent IO during incremental backup of the source.
								* `registered_source_max_concurrent_backups` - (Float) Specifies the limit on the number of backups Cohesity cluster will make concurrently to the registered entity. This limit is enforced only when the flag enforceRegisteredSourceMaxBackups is set to true.
								* `storage_array_snapshot_config` - (List) Specifies Storage Array Snapshot Configuration.
								Nested schema for **storage_array_snapshot_config**:
									* `is_max_snapshots_config_enabled` - (Boolean) Specifies if the storage array snapshot max snapshots config is enabled or not.
									* `is_max_space_config_enabled` - (Boolean) Specifies if the storage array snapshot max space config is enabled or not.
									* `storage_array_snapshot_max_space_config` - (List) Specifies Storage Array Snapshot Max Space Config.
									Nested schema for **storage_array_snapshot_max_space_config**:
										* `max_snapshot_space_percentage` - (Float) Max number of storage snapshots allowed per volume/lun.
									* `storage_array_snapshot_throttling_policies` - (List) Specifies throttling policies configured for individual volume/lun.
									Nested schema for **storage_array_snapshot_throttling_policies**:
										* `id` - (Integer) Specifies the volume id of the storage array snapshot config.
										* `is_max_snapshots_config_enabled` - (Boolean) Specifies if the storage array snapshot max snapshots config is enabled or not.
										* `is_max_space_config_enabled` - (Boolean) Specifies if the storage array snapshot max space config is enabled or not.
										* `max_snapshot_config` - (List) Specifies Storage Array Snapshot Max Snapshots Config.
										Nested schema for **max_snapshot_config**:
											* `max_snapshots` - (Float) Max number of storage snapshots allowed per volume/lun.
										* `max_space_config` - (List) Specifies Storage Array Snapshot Max Space Config.
										Nested schema for **max_space_config**:
											* `max_snapshot_space_percentage` - (Float) Max number of storage snapshots allowed per volume/lun.
						* `use_o_auth_for_exchange_online` - (Boolean) Specifies whether OAuth should be used for authentication in case of Exchange Online.
						* `use_vm_bios_uuid` - (Boolean) Specifies if registered vCenter is using BIOS UUID to track virtual machines.
						* `user_messages` - (List) Specifies the additional details encountered during registration. Though the registration may succeed, user messages imply the host environment requires some cleanup or fixing.
						* `username` - (String) Specifies username to access the target source.
						* `vlan_params` - (List) Specifies the VLAN configuration for Recovery.
						Nested schema for **vlan_params**:
							* `disable_vlan` - (Boolean) Specifies whether to use the VIPs even when VLANs are configured on the Cluster. If configured, VLAN IP addresses are used by default. If VLANs are not configured, this flag is ignored. Set this flag to true to force using the partition VIPs when VLANs are configured on the Cluster.
							* `interface_name` - (String) Specifies the physical interface group name to use for mounting Cohesity's view on the remote host. If specified, Cohesity hostname or the IP address on this VLAN is used.
							* `vlan` - (Float) Specifies the VLAN to use for mounting Cohesity's view on the remote host. If specified, Cohesity hostname or the IP address on this VLAN is used.
						* `warning_messages` - (List) Specifies a list of warnings encountered during registration. Though the registration may succeed, warning messages imply the host environment requires some cleanup or fixing.
					* `source_side_dedup_enabled` - (Boolean) Specifies whether source side dedup is enabled or not.
					* `status` - (String) Specifies the agent status. Specifies the status of the agent running on a physical source.
					  * Constraints: Allowable values are: `kUnknown`, `kUnreachable`, `kHealthy`, `kDegraded`.
					* `status_message` - (String) Specifies additional details about the agent status.
					* `upgradability` - (String) Specifies the upgradability of the agent running on the physical server. Specifies the upgradability of the agent running on the physical server.
					  * Constraints: Allowable values are: `kUpgradable`, `kCurrent`, `kUnknown`, `kNonUpgradableInvalidVersion`, `kNonUpgradableAgentIsNewer`, `kNonUpgradableAgentIsOld`.
					* `upgrade_status` - (String) Specifies the status of the upgrade of the agent on a physical server. Specifies the status of the upgrade of the agent on a physical server.
					  * Constraints: Allowable values are: `kIdle`, `kAccepted`, `kStarted`, `kFinished`, `kScheduled`.
					* `upgrade_status_message` - (String) Specifies detailed message about the agent upgrade failure. This field is not set for successful upgrade.
					* `version` - (String) Specifies the version of the Agent software.
					* `vol_cbt_info` - (List) CBT version and service state info.
					Nested schema for **vol_cbt_info**:
						* `file_version` - (List) Subcomponent version. The interpretation of the version is based on operating system.
						Nested schema for **file_version**:
							* `build_ver` - (Float)
							* `major_ver` - (Float)
							* `minor_ver` - (Float)
							* `revision_num` - (Float)
						* `is_installed` - (Boolean) Indicates whether the cbt driver is installed.
						* `reboot_status` - (String) Indicates whether host is rebooted post VolCBT installation.
						  * Constraints: Allowable values are: `kRebooted`, `kNeedsReboot`, `kInternalError`.
						* `service_state` - (List) Structure to Hold Service Status.
						Nested schema for **service_state**:
							* `name` - (String)
							* `state` - (String)
				* `cluster_source_type` - (String) Specifies the type of cluster resource this source represents.
				* `host_name` - (String) Specifies the hostname.
				* `host_type` - (String) Specifies the environment type for the host.
				  * Constraints: Allowable values are: `kLinux`, `kWindows`, `kAix`, `kSolaris`, `kSapHana`, `kSapOracle`, `kCockroachDB`, `kMySQL`, `kOther`, `kSapSybase`, `kSapMaxDB`, `kSapSybaseIQ`, `kDB2`, `kSapASE`, `kMariaDB`, `kPostgreSQL`, `kVOS`, `kHPUX`.
				* `id` - (List) Specifies an id for an object that is unique across Cohesity Clusters. The id is composite of all the ids listed below.
				Nested schema for **id**:
					* `cluster_id` - (Integer) Specifies the Cohesity Cluster id where the object was created.
					* `cluster_incarnation_id` - (Integer) Specifies an id for the Cohesity Cluster that is generated when a Cohesity Cluster is initially created.
					* `id` - (Integer) Specifies a unique id assigned to an object (such as a Job) by the Cohesity Cluster.
				* `is_proxy_host` - (Boolean) Specifies if the physical host is a proxy host.
				* `memory_size_bytes` - (Integer) Specifies the total memory on the host in bytes.
				* `name` - (String) Specifies a human readable name of the Protection Source.
				* `networking_info` - (List) Specifies the struct containing information about network addresses configured on the given box. This is needed for dealing with Windows/Oracle Cluster resources that we discover and protect automatically.
				Nested schema for **networking_info**:
					* `resource_vec` - (List) The list of resources on the system that are accessible by an IP address.
					Nested schema for **resource_vec**:
						* `endpoints` - (List) The endpoints by which the resource is accessible.
						Nested schema for **endpoints**:
							* `fqdn` - (String) The Fully Qualified Domain Name.
							* `ipv4_addr` - (String) The IPv4 address.
							* `ipv6_addr` - (String) The IPv6 address.
						* `type` - (String) The type of the resource.
				* `num_processors` - (Integer) Specifies the number of processors on the host.
				* `os_name` - (String) Specifies a human readable name of the OS of the Protection Source.
				* `type` - (String) Specifies the type of managed Object in a Physical Protection Source. 'kGroup' indicates the EH container.
				  * Constraints: Allowable values are: `kGroup`, `kHost`, `kWindowsCluster`, `kOracleRACCluster`, `kOracleAPCluster`.
				* `vcs_version` - (String) Specifies cluster version for VCS host.
				* `volumes` - (List) Array of Physical Volumes. Specifies the volumes available on the physical host. These fields are populated only for the kPhysicalHost type.
				Nested schema for **volumes**:
					* `device_path` - (String) Specifies the path to the device that hosts the volume locally.
					* `guid` - (String) Specifies an id for the Physical Volume.
					* `is_boot_volume` - (Boolean) Specifies whether the volume is boot volume.
					* `is_extended_attributes_supported` - (Boolean) Specifies whether this volume supports extended attributes (like ACLs) when performing file backups.
					* `is_protected` - (Boolean) Specifies if a volume is protected by a Job.
					* `is_shared_volume` - (Boolean) Specifies whether the volume is shared volume.
					* `label` - (String) Specifies a volume label that can be used for displaying additional identifying information about a volume.
					* `logical_size_bytes` - (Float) Specifies the logical size of the volume in bytes that is not reduced by change-block tracking, compression and deduplication.
					* `mount_points` - (List) Specifies the mount points where the volume is mounted, for example- 'C:', '/mnt/foo' etc.
					* `mount_type` - (String) Specifies mount type of volume e.g. nfs, autofs, ext4 etc.
					* `network_path` - (String) Specifies the full path to connect to the network attached volume. For example, (IP or hostname):/path/to/share for NFS volumes).
					* `used_size_bytes` - (Float) Specifies the size used by the volume in bytes.
				* `vsswriters` - (List)
				Nested schema for **vsswriters**:
					* `is_writer_excluded` - (Boolean) If true, the writer will be excluded by default.
					* `writer_name` - (Boolean) Specifies the name of the writer.
			* `sql_protection_source` - (List) Specifies an Object representing one SQL Server instance or database.
			Nested schema for **sql_protection_source**:
				* `created_timestamp` - (String) Specifies the time when the database was created. It is displayed in the timezone of the SQL server on which this database is running.
				* `database_name` - (String) Specifies the database name of the SQL Protection Source, if the type is database.
				* `db_aag_entity_id` - (Integer) Specifies the AAG entity id if the database is part of an AAG. This field is set only for type 'kDatabase'.
				* `db_aag_name` - (String) Specifies the name of the AAG if the database is part of an AAG. This field is set only for type 'kDatabase'.
				* `db_compatibility_level` - (Integer) Specifies the versions of SQL server that the database is compatible with.
				* `db_file_groups` - (List) Specifies the information about the set of file groups for this db on the host. This is only set if the type is kDatabase.
				* `db_files` - (List) Specifies the last known information about the set of database files on the host. This field is set only for type 'kDatabase'.
				Nested schema for **db_files**:
					* `file_type` - (String) Specifies the format type of the file that SQL database stores the data. Specifies the format type of the file that SQL database stores the data. 'kRows' refers to a data file 'kLog' refers to a log file 'kFileStream' refers to a directory containing FILESTREAM data 'kNotSupportedType' is for information purposes only. Not supported. 'kFullText' refers to a full-text catalog.
					  * Constraints: Allowable values are: `kRows`, `kLog`, `kFileStream`, `kNotSupportedType`, `kFullText`.
					* `full_path` - (String) Specifies the full path of the database file on the SQL host machine.
					* `size_bytes` - (Integer) Specifies the last known size of the database file.
				* `db_owner_username` - (String) Specifies the name of the database owner.
				* `default_database_location` - (String) Specifies the default path for data files for DBs in an instance.
				* `default_log_location` - (String) Specifies the default path for log files for DBs in an instance.
				* `id` - (List) Specifies a unique id for a SQL Protection Source.
				Nested schema for **id**:
					* `created_date_msecs` - (Integer) Specifies a unique identifier generated from the date the database is created or renamed. Cohesity uses this identifier in combination with the databaseId to uniquely identify a database.
					* `database_id` - (Integer) Specifies a unique id of the database but only for the life of the database. SQL Server may reuse database ids. Cohesity uses the createDateMsecs in combination with this databaseId to uniquely identify a database.
					* `instance_id` - (String) Specifies unique id for the SQL Server instance. This id does not change during the life of the instance.
				* `is_available_for_vss_backup` - (Boolean) Specifies whether the database is marked as available for backup according to the SQL Server VSS writer. This may be false if either the state of the databases is not online, or if the VSS writer is not online. This field is set only for type 'kDatabase'.
				* `is_encrypted` - (Boolean) Specifies whether the database is TDE enabled.
				* `name` - (String) Specifies the instance name of the SQL Protection Source.
				* `owner_id` - (Integer) Specifies the id of the container VM for the SQL Protection Source.
				* `recovery_model` - (String) Specifies the Recovery Model for the database in SQL environment. Only meaningful for the 'kDatabase' SQL Protection Source. Specifies the Recovery Model set for the Microsoft SQL Server. 'kSimpleRecoveryModel' indicates the Simple SQL Recovery Model which does not utilize log backups. 'kFullRecoveryModel' indicates the Full SQL Recovery Model which requires log backups and allows recovery to a single point in time. 'kBulkLoggedRecoveryModel' indicates the Bulk Logged SQL Recovery Model which requires log backups and allows high-performance bulk copy operations.
				  * Constraints: Allowable values are: `kSimpleRecoveryModel`, `kFullRecoveryModel`, `kBulkLoggedRecoveryModel`.
				* `sql_server_db_state` - (String) The state of the database as returned by SQL Server. Indicates the state of the database. The values correspond to the 'state' field in the system table sys.databases. See https://goo.gl/P66XqM. 'kOnline' indicates that database is in online state. 'kRestoring' indicates that database is in restore state. 'kRecovering' indicates that database is in recovery state. 'kRecoveryPending' indicates that database recovery is in pending state. 'kSuspect' indicates that primary filegroup is suspect and may be damaged. 'kEmergency' indicates that manually forced emergency state. 'kOffline' indicates that database is in offline state. 'kCopying' indicates that database is in copying state. 'kOfflineSecondary' indicates that secondary database is in offline state.
				  * Constraints: Allowable values are: `kOnline`, `kRestoring`, `kRecovering`, `kRecoveryPending`, `kSuspect`, `kEmergency`, `kOffline`, `kCopying`, `kOfflineSecondary`.
				* `sql_server_instance_version` - (List) Specifies the Server Instance Version.
				Nested schema for **sql_server_instance_version**:
					* `build` - (Float) Specifies the build.
					* `major_version` - (Float) Specifies the major version.
					* `minor_version` - (Float) Specifies the minor version.
					* `revision` - (Float) Specifies the revision.
					* `version_string` - (Float) Specifies the version string.
				* `type` - (String) Specifies the type of the managed Object in a SQL Protection Source. Examples of SQL Objects include 'kInstance' and 'kDatabase'. 'kInstance' indicates that SQL server instance is being protected. 'kDatabase' indicates that SQL server database is being protected. 'kAAG' indicates that SQL AAG (AlwaysOn Availability Group) is being protected. 'kAAGRootContainer' indicates that SQL AAG's root container is being protected. 'kRootContainer' indicates root container for SQL sources.
				  * Constraints: Allowable values are: `kInstance`, `kDatabase`, `kAAG`, `kAAGRootContainer`, `kRootContainer`.
		* `environment` - (String) Specifies the environment type of the application such as 'kSQL', 'kExchange' registered on the Protection Source. overrideDescription: true Supported environment types such as 'kView', 'kSQL', 'kVMware', etc. NOTE: 'kPuppeteer' refers to Cohesity's Remote Adapter.'kVMware' indicates the VMware Protection Source environment. 'kHyperV' indicates the HyperV Protection Source environment.'kSQL' indicates the SQL Protection Source environment.'kView' indicates the View Protection Source environment. 'kPuppeteer' indicates the Cohesity's Remote Adapter. 'kPhysical' indicates the physical Protection Source environment. 'kPure' indicates the Pure Storage Protection Source environment.'kNimble' indicates the Nimble Storage Protection Source environment. 'kHpe3Par' indicates the Hpe 3Par Storage Protection Source environment. 'kAzure' indicates the Microsoft's Azure Protection Source environment. 'kNetapp' indicates the Netapp Protection Source environment.'kAgent' indicates the Agent Protection Source environment. 'kGenericNas' indicates the Generic Network Attached Storage Protection Source environment. 'kAcropolis' indicates the Acropolis Protection Source environment. 'kPhysicalFiles' indicates the Physical Files Protection Source environment. 'kIbmFlashSystem' indicates the IBM Flash System Protection Source environment. 'kIsilon' indicates the Dell EMC's Isilon Protection Source environment. 'kGPFS' indicates IBM's GPFS Protection Source environment. 'kKVM' indicates the KVM Protection Source environment. 'kAWS' indicates the AWS Protection Source environment. 'kExchange' indicates the Exchange Protection Source environment. 'kHyperVVSS' indicates the HyperV VSS Protection Source environment. 'kOracle' indicates the Oracle Protection Source environment. 'kGCP' indicates the Google Cloud Platform Protection Source environment. 'kFlashBlade' indicates the Flash Blade Protection Source environment. 'kAWSNative' indicates the AWS Native Protection Source environment. 'kO365' indicates the Office 365 Protection Source environment. 'kO365Outlook' indicates Office 365 outlook Protection Source environment. 'kHyperFlex' indicates the Hyper Flex Protection Source environment. 'kGCPNative' indicates the GCP Native Protection Source environment. 'kAzureNative' indicates the Azure Native Protection Source environment. 'kKubernetes' indicates a Kubernetes Protection Source environment. 'kElastifile' indicates Elastifile Protection Source environment. 'kAD' indicates Active Directory Protection Source environment. 'kRDSSnapshotManager' indicates AWS RDS Protection Source environment. 'kCassandra' indicates Cassandra Protection Source environment. 'kMongoDB' indicates MongoDB Protection Source environment. 'kCouchbase' indicates Couchbase Protection Source environment. 'kHdfs' indicates Hdfs Protection Source environment. 'kHive' indicates Hive Protection Source environment. 'kHBase' indicates HBase Protection Source environment. 'kUDA' indicates Universal Data Adapter Protection Source environment. 'kSAPHANA' indicates SAP HANA protection source environment. 'kO365Teams' indicates the Office365 Teams Protection Source environment. 'kO365Group' indicates the Office365 Groups Protection Source environment. 'kO365Exchange' indicates the Office365 Mailbox Protection Source environment. 'kO365OneDrive' indicates the Office365 OneDrive Protection Source environment. 'kO365Sharepoint' indicates the Office365 SharePoint Protection Source environment. 'kO365PublicFolders' indicates the Office365 PublicFolders Protection Source environment. kHpe3Par, kIbmFlashSystem, kAzure, kNetapp, kAgent, kGenericNas, kAcropolis, kPhysicalFiles, kIsilon, kGPFS, kKVM, kAWS, kExchange, kHyperVVSS, kOracle, kGCP, kFlashBlade, kAWSNative, kO365, kO365Outlook, kHyperFlex, kGCPNative, kAzureNative, kKubernetes, kElastifile, kAD, kRDSSnapshotManager, kCassandra, kMongoDB, kCouchbase, kHdfs, kHive, kHBase, kUDA, kSAPHANA, kO365Teams, kO365Group, kO365Exchange, kO365OneDrive, kO365Sharepoint, kO365PublicFolders.
		  * Constraints: Allowable values are: `kVMware`, `kHyperV`, `kSQL`, `kView`, `kPuppeteer`, `kPhysical`, `kPure`, `kNimble`.
	* `entity_permission_info` - (List) Specifies the permission information of entities.
	Nested schema for **entity_permission_info**:
		* `entity_id` - (Integer) Specifies the entity id.
		* `groups` - (List) Specifies groups that have access to entity in case of restricted user.
		Nested schema for **groups**:
			* `domain` - (String) Specifies domain name of the user.
			* `group_name` - (String) Specifies group name of the group.
			* `sid` - (String) Specifies unique Security ID (SID) of the user.
			* `tenant_ids` - (List) Specifies the tenants to which the group belongs to.
		* `is_inferred` - (Boolean) Specifies whether the Entity Permission Information is inferred or not. For example, SQL application hosted over vCenter will have inferred entity permission information.
		* `is_registered_by_sp` - (Boolean) Specifies whether this entity is registered by the SP or not. This will be populated only if the entity is a root entity. Refer to magneto/base/permissions.proto for details.
		* `registering_tenant_id` - (String) Specifies the tenant id that registered this entity. This will be populated only if the entity is a root entity.
		* `tenant` - (List) Specifies struct with basic tenant details.
		Nested schema for **tenant**:
			* `bifrost_enabled` - (Boolean) Specifies if this tenant is bifrost enabled or not.
			* `is_managed_on_helios` - (Boolean) Specifies whether this tenant is manged on helios.
			* `name` - (String) Specifies name of the tenant.
			* `tenant_id` - (String) Specifies the unique id of the tenant.
		* `users` - (List) Specifies users that have access to entity in case of restricted user.
		Nested schema for **users**:
			* `domain` - (String) Specifies domain name of the user.
			* `sid` - (String) Specifies unique Security ID (SID) of the user.
			* `tenant_id` - (String) Specifies the tenant to which the user belongs to.
			* `user_name` - (String) Specifies user name of the user.
	* `logical_size_bytes` - (Integer) Specifies the logical size of the Protection Source in bytes.
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
	* `registration_info` - (List) Specifies registration information for a root node in a Protection Sources tree. A root node represents a registered Source on the Cohesity Cluster, such as a vCenter Server.
	Nested schema for **registration_info**:
		* `access_info` - (List) Specifies the parameters required to establish a connection with a particular environment.
		Nested schema for **access_info**:
			* `connection_id` - (Integer) ID of the Bifrost (HyX or Rigel) network realm (i.e. a connection) associated with the source.
			* `connector_group_id` - (Integer) Specifies the Id of the connector group. Each connector group is collection of Rigel/hyx. Each entity will be tagged with connector group id.
			* `endpoint` - (String) Specify an IP address or URL of the environment. (such as the IP address of the vCenter Server for a VMware environment).
			* `environment` - (String) Specifies the environment like VMware, SQL, where the Protection Source exists. Supported environment types such as 'kView', 'kSQL', 'kVMware', etc. NOTE: 'kPuppeteer' refers to Cohesity's Remote Adapter. 'kVMware' indicates the VMware Protection Source environment. 'kHyperV' indicates the HyperV Protection Source environment. 'kSQL' indicates the SQL Protection Source environment. 'kView' indicates the View Protection Source environment. 'kPuppeteer' indicates the Cohesity's Remote Adapter. 'kPhysical' indicates the physical Protection Source environment. 'kPure' indicates the Pure Storage Protection Source environment. 'kNimble' indicates the Nimble Storage Protection Source environment. 'kHpe3Par' indicates the Hpe 3Par Storage Protection Source environment.'kAzure' indicates the Microsoft's Azure Protection Source environment. 'kNetapp' indicates the Netapp Protection Source environment. 'kAgent' indicates the Agent Protection Source environment. 'kGenericNas' indicates the Generic Network Attached Storage Protection Source environment. 'kAcropolis' indicates the Acropolis Protection Source environment. 'kPhysicalFiles' indicates the Physical Files Protection Source environment. 'kIbmFlashSystem' indicates the IBM Flash System Protection Source environment. 'kIsilon' indicates the Dell EMC's Isilon Protection Source environment. 'kGPFS' indicates IBM's GPFS Protection Source environment. 'kKVM' indicates the KVM Protection Source environment. 'kAWS' indicates the AWS Protection Source environment. 'kExchange' indicates the Exchange Protection Source environment. 'kHyperVVSS' indicates the HyperV VSS Protection Source environment. 'kOracle' indicates the Oracle Protection Source environment. 'kGCP' indicates the Google Cloud Platform Protection Source environment. 'kFlashBlade' indicates the Flash Blade Protection Source environment. 'kAWSNative' indicates the AWS Native Protection Source environment. 'kO365' indicates the Office 365 Protection Source environment. 'kO365Outlook' indicates Office 365 outlook Protection Source environment. 'kHyperFlex' indicates the Hyper Flex Protection Source environment. 'kGCPNative' indicates the GCP Native Protection Source environment. 'kAzureNative' indicates the Azure Native Protection Source environment. 'kKubernetes' indicates a Kubernetes Protection Source environment. 'kElastifile' indicates Elastifile Protection Source environment. 'kAD' indicates Active Directory Protection Source environment. 'kRDSSnapshotManager' indicates AWS RDS Protection Source environment. 'kCassandra' indicates Cassandra Protection Source environment. 'kMongoDB' indicates MongoDB Protection Source environment. 'kCouchbase' indicates Couchbase Protection Source environment. 'kHdfs' indicates Hdfs Protection Source environment. 'kHive' indicates Hive Protection Source environment. 'kHBase' indicates HBase Protection Source environment. 'kUDA' indicates Universal Data Adapter Protection Source environment. 'kSAPHANA' indicates SAP HANA protection source environment. 'kO365Teams' indicates the Office365 Teams Protection Source environment. 'kO365Group' indicates the Office365 Groups Protection Source environment. 'kO365Exchange' indicates the Office365 Mailbox Protection Source environment. 'kO365OneDrive' indicates the Office365 OneDrive Protection Source environment. 'kO365Sharepoint' indicates the Office365 SharePoint Protection Source environment. 'kO365PublicFolders' indicates the Office365 PublicFolders Protection Source environment. kHpe3Par, kIbmFlashSystem, kAzure, kNetapp, kAgent, kGenericNas, kAcropolis, kPhysicalFiles, kIsilon, kGPFS, kKVM, kAWS, kExchange, kHyperVVSS, kOracle, kGCP, kFlashBlade, kAWSNative, kO365, kO365Outlook, kHyperFlex, kGCPNative, kAzureNative, kKubernetes, kElastifile, kAD, kRDSSnapshotManager, kCassandra, kMongoDB, kCouchbase, kHdfs, kHive, kHBase, kUDA, kSAPHANA, kO365Teams, kO365Group, kO365Exchange, kO365OneDrive, kO365Sharepoint, kO365PublicFolders.
			  * Constraints: Allowable values are: `kVMware`, `kHyperV`, `kSQL`, `kView`, `kPuppeteer`, `kPhysical`, `kPure`, `kNimble`.
			* `id` - (Integer) Specifies a Unique id that is generated when the Source is registered. This is a convenience field that is used to maintain an index to different connection params.
			* `version` - (Integer) Version is updated each time the connector parameters are updated. This is used to discard older connector parameters.
		* `allowed_ip_addresses` - (List) Specifies the list of IP Addresses on the registered source to be exclusively allowed for doing any type of IO operations.
		* `authentication_error_message` - (String) Specifies an authentication error message. This indicates the given credentials are rejected and the registration of the source is not successful.
		* `authentication_status` - (String) Specifies the status of the authenticating to the Protection Source when registering it with Cohesity Cluster. If the status is 'kFinished' and there is no error, registration is successful. Specifies the status of the authentication during the registration of a Protection Source. 'kPending' indicates the authentication is in progress. 'kScheduled' indicates the authentication is scheduled. 'kFinished' indicates the authentication is completed. 'kRefreshInProgress' indicates the refresh is in progress.
		  * Constraints: Allowable values are: `kPending`, `kScheduled`, `kFinished`, `kRefreshInProgress`.
		* `blacklisted_ip_addresses` - (List) This field is deprecated. Use DeniedIpAddresses instead. deprecated: true.
		* `cassandra_params` - (List) Specifies an Object containing information about a registered cassandra source.
		Nested schema for **cassandra_params**:
			* `cassandra_ports_info` - (List) Specifies an Object containing information on various Cassandra ports.
			Nested schema for **cassandra_ports_info**:
				* `jmx_port` - (Integer) Specifies the Cassandra JMX port.
				* `native_transport_port` - (Integer) Specifies the port for the CQL native transport.
				* `rpc_port` - (Integer) Specifies the Remote Procedure Call (RPC) port for general mechanism for client-server applications.
				* `ssl_storage_port` - (Integer) Specifies the SSL port for encrypted communication.
				* `storage_port` - (Integer) Specifies the TCP port for data. Internally used by Cassandra bulk loader.
			* `cassandra_security_info` - (List) Specifies an Object containing information on Cassandra security.
			Nested schema for **cassandra_security_info**:
				* `cassandra_auth_required` - (Boolean) Is Cassandra authentication required ?.
				* `cassandra_auth_type` - (String) Cassandra Authentication type. Enum: [PASSWORD KERBEROS LDAP] Specifies the Cassandra auth type.'PASSWORD' 'KERBEROS' 'LDAP'.
				  * Constraints: Allowable values are: `PASSWORD`, `KERBEROS`, `LDAP`.
				* `cassandra_authorizer` - (String) Cassandra Authenticator/Authorizer.
				* `client_encryption` - (Boolean) Is Client Encryption enabled for this cluster ?.
				* `dse_authorization` - (Boolean) Is DSE Authorization enabled for this cluster ?.
				* `server_encryption_req_client_auth` - (Boolean) Is 'Server encryption request client authentication' enabled for this cluster ?.
				* `server_internode_encryption_type` - (String) 'Server internal node Encryption' type for this cluster.
			* `cassandra_version` - (String) Cassandra version.
			* `commit_log_backup_location` - (String) Specifies the commit log archival location for cassandra node.
			* `config_directory` - (String) Specifies the Directory path containing Config YAML for discovery.
			* `data_centers` - (List) Specifies the List of all physical data center or virtual data center. In most cases, the data centers will be listed after discovery operation however, if they are not listed, you must manually type the data center names. Leaving the field blank will disallow data center-specific backup or restore. Entering a subset of all data centers may cause problems in data movement.
			* `dse_config_directory` - (String) Specifies the Directory from where DSE specific configuration can be read.
			* `dse_version` - (String) DSE version.
			* `is_dse_authenticator` - (Boolean) Specifies whether this cluster has DSE Authenticator.
			* `is_dse_tiered_storage` - (Boolean) Specifies whether this cluster has DSE tiered storage.
			* `is_jmx_auth_enable` - (Boolean) Specifies if JMX Authentication enabled in this cluster.
			* `kerberos_principal` - (String) Specifies the Kerberos Principal for Kerberos connection.
			* `primary_host` - (String) Specifies the Primary Host for the Cassandra cluster.
			* `seeds` - (List) Specifies the Seed nodes of this Cassandra cluster.
			* `solr_nodes` - (List) Specifies the Solr node IP Addresses.
			* `solr_port` - (Integer) Specifies the Solr node Port.
		* `couchbase_params` - (List) Specifies an Object containing information about a registered couchbase source.
		Nested schema for **couchbase_params**:
			* `carrier_direct_port` - (Integer) Specifies the Carrier direct/sll port.
			* `http_direct_port` - (Integer) Specifies the HTTP direct/sll port.
			* `requires_ssl` - (Boolean) Specifies whether this cluster allows connection through SSL only.
			* `seeds` - (List) Specifies the Seeds of this Couchbase Cluster.
		* `denied_ip_addresses` - (List) Specifies the list of IP Addresses on the registered source to be denied for doing any type of IO operations.
		* `environments` - (List) Specifies a list of applications environment that are registered with this Protection Source such as 'kSQL'. Supported environment types such as 'kView', 'kSQL', 'kVMware', etc. NOTE: 'kPuppeteer' refers to Cohesity's Remote Adapter. 'kVMware' indicates the VMware Protection Source environment. 'kHyperV' indicates the HyperV Protection Source environment. 'kSQL' indicates the SQL Protection Source environment. 'kView' indicates the View Protection Source environment. 'kPuppeteer' indicates the Cohesity's Remote Adapter. 'kPhysical' indicates the physical Protection Source environment. 'kPure' indicates the Pure Storage Protection Source environment. 'kNimble' indicates the Nimble Storage Protection Source environment. 'kHpe3Par' indicates the Hpe 3Par Storage Protection Source environment. 'kAzure' indicates the Microsoft's Azure Protection Source environment. 'kNetapp' indicates the Netapp Protection Source environment. 'kAgent' indicates the Agent Protection Source environment. 'kGenericNas' indicates the Generic Network Attached Storage Protection Source environment. 'kAcropolis' indicates the Acropolis Protection Source environment. 'kPhysicalFiles' indicates the Physical Files Protection Source environment. 'kIbmFlashSystem' indicates the IBM Flash System Protection Source environment. 'kIsilon' indicates the Dell EMC's Isilon Protection Source environment. 'kGPFS' indicates IBM's GPFS Protection Source environment. 'kKVM' indicates the KVM Protection Source environment. 'kAWS' indicates the AWS Protection Source environment. 'kExchange' indicates the Exchange Protection Source environment. 'kHyperVVSS' indicates the HyperV VSS Protection Source environment. 'kOracle' indicates the Oracle Protection Source environment. 'kGCP' indicates the Google Cloud Platform Protection Source environment. 'kFlashBlade' indicates the Flash Blade Protection Source environment. 'kAWSNative' indicates the AWS Native Protection Source environment. 'kO365' indicates the Office 365 Protection Source environment. 'kO365Outlook' indicates Office 365 outlook Protection Source environment. 'kHyperFlex' indicates the Hyper Flex Protection Source environment. 'kGCPNative' indicates the GCP Native Protection Source environment. 'kAzureNative' indicates the Azure Native Protection Source environment. 'kKubernetes' indicates a Kubernetes Protection Source environment. 'kElastifile' indicates Elastifile Protection Source environment. 'kAD' indicates Active Directory Protection Source environment. 'kRDSSnapshotManager' indicates AWS RDS Protection Source environment. 'kCassandra' indicates Cassandra Protection Source environment. 'kMongoDB' indicates MongoDB Protection Source environment. 'kCouchbase' indicates Couchbase Protection Source environment. 'kHdfs' indicates Hdfs Protection Source environment. 'kHive' indicates Hive Protection Source environment. 'kHBase' indicates HBase Protection Source environment. 'kUDA' indicates Universal Data Adapter Protection Source environment. 'kSAPHANA' indicates SAP HANA protection source environment. 'kO365Teams' indicates the Office365 Teams Protection Source environment. 'kO365Group' indicates the Office365 Groups Protection Source environment. 'kO365Exchange' indicates the Office365 Mailbox Protection Source environment. 'kO365OneDrive' indicates the Office365 OneDrive Protection Source environment. 'kO365Sharepoint' indicates the Office365 SharePoint Protection Source environment. 'kO365PublicFolders' indicates the Office365 PublicFolders Protection Source environment. kHpe3Par, kIbmFlashSystem, kAzure, kNetapp, kAgent, kGenericNas, kAcropolis, kPhysicalFiles, kIsilon, kGPFS, kKVM, kAWS, kExchange, kHyperVVSS, kOracle, kGCP, kFlashBlade, kAWSNative, kO365, kO365Outlook, kHyperFlex, kGCPNative, kAzureNative, kKubernetes, kElastifile, kAD, kRDSSnapshotManager, kCassandra, kMongoDB, kCouchbase, kHdfs, kHive, kHBase, kUDA, kSAPHANA, kO365Teams, kO365Group, kO365Exchange, kO365OneDrive, kO365Sharepoint, kO365PublicFolders.
		  * Constraints: Allowable list items are: `kVMware`, `kHyperV`, `kSQL`, `kView`, `kPuppeteer`, `kPhysical`, `kPure`, `kNimble`.
		* `hbase_params` - (List) Specifies an Object containing information about a registered HBase source.
		Nested schema for **hbase_params**:
			* `hbase_discovery_params` - (List) Specifies an Object containing information about discovering a Hadoop source.
			Nested schema for **hbase_discovery_params**:
				* `config_directory` - (String) Specifies the configuration directory.
				* `host` - (String) Specifies the host IP.
			* `hdfs_entity_id` - (Integer) The entity id of the HDFS source for this HBase.
			* `kerberos_principal` - (String) Specifies the kerberos principal.
			* `root_data_directory` - (String) Specifies the HBase data root directory.
			* `zookeeper_quorum` - (List) Specifies the HBase zookeeper quorum.
		* `hdfs_params` - (List) Specifies an Object containing information about a registered Hdfs source.
		Nested schema for **hdfs_params**:
			* `hadoop_distribution` - (String) Specifies the Hadoop Distribution. Hadoop distribution. 'CDH' indicates Hadoop distribution type Cloudera. 'HDP' indicates Hadoop distribution type Hortonworks.
			  * Constraints: Allowable values are: `CDH`, `HDP`.
			* `hadoop_version` - (String) Specifies the Hadoop version.
			* `hdfs_connection_type` - (String) Specifies the Hdfs connection type. Hdfs connection type. 'DFS' indicates Hdfs connection type DFS. 'WEBHDFS' indicates Hdfs connection type WEBHDFS. 'HTTPFSLB' indicates Hdfs connection type HTTPFS_LB. 'HTTPFS' indicates Hdfs connection type HTTPFS.
			  * Constraints: Allowable values are: `DFS`, `WEBHDFS`, `HTTPFSLB`, `HTTPFS`.
			* `hdfs_discovery_params` - (List) Specifies an Object containing information about discovering a Hadoop source.
			Nested schema for **hdfs_discovery_params**:
				* `config_directory` - (String) Specifies the configuration directory.
				* `host` - (String) Specifies the host IP.
			* `kerberos_principal` - (String) Specifies the kerberos principal.
			* `namenode` - (String) Specifies the Namenode host or Nameservice.
			* `port` - (Integer) Specifies the Webhdfs Port.
		* `hive_params` - (List) Specifies an Object containing information about a registered Hive source.
		Nested schema for **hive_params**:
			* `entity_threshold_exceeded` - (Boolean) Specifies if max entity count exceeded for protection source view.
			* `hdfs_entity_id` - (Integer) Specifies the entity id of the HDFS source for this Hive.
			* `hive_discovery_params` - (List) Specifies an Object containing information about discovering a Hadoop source.
			Nested schema for **hive_discovery_params**:
				* `config_directory` - (String) Specifies the configuration directory.
				* `host` - (String) Specifies the host IP.
			* `kerberos_principal` - (String) Specifies the kerberos principal.
			* `metastore` - (String) Specifies the Hive metastore host.
			* `thrift_port` - (Integer) Specifies the Hive metastore thrift Port.
		* `is_db_authenticated` - (Boolean) Specifies if application entity dbAuthenticated or not. ex: oracle database.
		* `is_storage_array_snapshot_enabled` - (Boolean) Specifies if this source entity has enabled storage array snapshot or not.
		* `isilon_params` - (List) Specifies the Isilon specific Registered Protection Source params. This definition is used to send isilion source params in update protection source params to magneto.
		Nested schema for **isilon_params**:
			* `zone_config_list` - (List) List of access zone info in an Isilion Cluster.
			Nested schema for **zone_config_list**:
				* `dynamic_network_pool_config` - (List) While caonfiguring the isilon protection source, this is the selected network pool config for the isilon access zone.
				Nested schema for **dynamic_network_pool_config**:
					* `pool_name` - (String) Specifies the name of the Network pool.
					* `subnet` - (String) Specifies the name of the subnet the network pool belongs to.
					* `use_smart_connect` - (Boolean) Specifies whether to use SmartConnect if available. If true, DNS name for the SmartConnect zone will be used to balance the IPs. Otherwise, pool IPs will be balanced manually.
		* `link_vms_across_vcenter` - (Boolean) Specifies if the VM linking feature is enabled for this VCenter This means that VMs present in this VCenter which earlier belonged to some other VCenter(also registerd on same cluster) and were migrated, will be linked during EH refresh. This will enable preserving snapshot chains for migrated VMs.
		* `minimum_free_space_gb` - (Integer) Specifies the minimum free space in GiB of the space expected to be available on the datastore where the virtual disks of the VM being backed up. If the amount of free space(in GiB) is lower than the value given by this field, backup will be aborted. Note that this field is applicable only to 'kVMware' type of environments.
		* `minimum_free_space_percent` - (Integer) Specifies the minimum free space in percentage of the space expected to be available on the datastore where the virtual disks of the VM being backed up. If the amount of free space(in percentage) is lower than the value given by this field, backup will be aborted. Note that this field is applicable only to 'kVMware' type of environments.
		* `mongodb_params` - (List) Specifies an Object containing information about a registered mongodb source.
		Nested schema for **mongodb_params**:
			* `auth_type` - (String) Specifies whether authentication is configured on this MongoDB cluster. Specifies the type of an MongoDB source entity. 'SCRAM' 'LDAP' 'NONE' 'KERBEROS'.
			  * Constraints: Allowable values are: `SCRAM`, `LDAP`, `NONE`, `KERBEROS`.
			* `authenticating_database_name` - (String) Specifies the Authenticating Database for this MongoDB cluster.
			* `requires_ssl` - (Boolean) Specifies whether connection is allowed through SSL only in this cluster.
			* `secondary_node_tag` - (String) MongoDB Secondary node tag. Required only if 'useSecondaryForBackup' is true. The system will use this to identify the secondary nodes for reading backup data.
			* `seeds` - (List) Specifies the seeds of this MongoDB Cluster.
			* `use_fixed_node_for_backup` - (Boolean) Set this to true if you want the system to peform backups from fixed nodes.
			* `use_secondary_for_backup` - (Boolean) Set this to true if you want the system to peform backups from secondary nodes.
		* `nas_mount_credentials` - (List) Specifies the credentials required to mount directories on the NetApp server if given.
		Nested schema for **nas_mount_credentials**:
			* `domain` - (String) Specifies the domain in which this credential is valid.
			* `nas_protocol` - (String) Specifies the protocol used by the NAS server. Specifies the protocol used by a NAS server. 'kNoProtocol' indicates no protocol set. 'kNfs3' indicates NFS v3 protocol. 'kNfs4_1' indicates NFS v4.1 protocol. 'kCifs1' indicates CIFS v1.0 protocol. 'kCifs2' indicates CIFS v2.0 protocol. 'kCifs3' indicates CIFS v3.0 protocol.
			  * Constraints: Allowable values are: `kNoProtocol`, `kNfs3`, `kNfs4_1`, `kCifs1`, `kCifs2`, `kCifs3`.
		* `o365_params` - (List) Specifies an Object containing information about a registered Office 365 source.
		Nested schema for **o365_params**:
			* `csm_params` - (List)
			Nested schema for **csm_params**:
				* `backup_allowed` - (Boolean) Specifies whether the current source allows data backup through M365 Backup Storage APIs. Enabling this, data can be optionally backed up within either Cohesity or MSFT or both depending on the backup configuration.
			* `objects_discovery_params` - (List) Specifies the parameters used for discovering the office 365 objects selectively during source registration or refresh.
			Nested schema for **objects_discovery_params**:
				* `discoverable_object_type_list` - (List) Specifies the list of object types that will be discovered as part of source registration or refresh.
				* `sites_discovery_params` - (List) Specifies discovery params for kSite entities. It should only be populated when the 'DiscoveryParams.discoverableObjectTypeList' includes 'kSites'.
				Nested schema for **sites_discovery_params**:
					* `enable_site_tagging` - (Boolean) Specifies whether the SharePoint Sites will be tagged whether they belong to a group site or teams site.
				* `teams_additional_params` - (List) Specifies additional params for Teams entities. It should only be populated if the 'DiscoveryParams.discoverableObjectTypeList' includes 'kTeams' otherwise this will be ignored.
				Nested schema for **teams_additional_params**:
					* `allow_posts_backup` - (Boolean) Specifies whether the Teams posts/conversations will be backed up or not. If this is false or not specified teams' posts backup will not be done.
				* `users_discovery_params` - (List) Specifies discovery params for kUser entities. It should only be populated when the 'DiscoveryParams.discoverableObjectTypeList' includes 'kUsers'.
				Nested schema for **users_discovery_params**:
					* `allow_chats_backup` - (Boolean) Specifies whether users' chats should be backed up or not. If this is false or not specified users' chats backup will not be done.
					* `discover_users_with_mailbox` - (Boolean) Specifies if office 365 users with valid mailboxes should be discovered or not.
					* `discover_users_with_onedrive` - (Boolean) Specifies if office 365 users with valid Onedrives should be discovered or not.
					* `fetch_mailbox_info` - (Boolean) Specifies whether users' mailbox info including the provisioning status, mailbox type & in-place archival support will be fetched and processed.
					* `fetch_one_drive_info` - (Boolean) Specifies whether users' onedrive info including the provisioning status & storage quota will be fetched and processed.
					* `skip_users_without_my_site` - (Boolean) Specifies whether to skip processing user who have uninitialized OneDrive or are without MySite.
		* `office365_credentials_list` - (List) Office365 Source Credentials. Specifies credentials needed to authenticate & authorize user for Office365.
		Nested schema for **office365_credentials_list**:
			* `client_id` - (String) Specifies the application ID that the registration portal (apps.dev.microsoft.com) assigned.
			* `client_secret` - (String) Specifies the application secret that was created in app registration portal.
			* `grant_type` - (String) Specifies the application grant type. eg: For client credentials flow, set this to "client_credentials"; For refreshing access-token, set this to "refresh_token".
			* `scope` - (String) Specifies a space separated list of scopes/permissions for the user. eg: Incase of MS Graph APIs for Office365, scope is set to default: https://graph.microsoft.com/.default.
			* `use_o_auth_for_exchange_online` - (Boolean) This field is deprecated from here and placed in RegisteredSourceInfo  and ProtectionSourceParameters. deprecated: true.
		* `office365_region` - (String) Specifies the region for Office365. Inorder to truly categorize M365 region, clients should not depend upon the endpoint, instead look at this attribute for the same.
		* `office365_service_account_credentials_list` - (List) Office365 Service Account Credentials. Specifies credentials for improving mailbox backup performance for O365.
		Nested schema for **office365_service_account_credentials_list**:
			* `password` - (String) Specifies the password to access target entity.
			* `username` - (String) Specifies the username to access target entity.
		* `password` - (String) Specifies password of the username to access the target source.
		* `physical_params` - (List) Specifies the parameters required to register Application Servers running in a Protection Source specific to a physical adapter.
		Nested schema for **physical_params**:
			* `applications` - (List) Specifies the types of applications such as 'kSQL', 'kExchange', 'kAD' running on the Protection Source. overrideDescription: true Supported environment types such as 'kView', 'kSQL', 'kVMware', etc. NOTE: 'kPuppeteer' refers to Cohesity's Remote Adapter. 'kVMware' indicates the VMware Protection Source environment. 'kHyperV' indicates the HyperV Protection Source environment. 'kSQL' indicates the SQL Protection Source environment. 'kView' indicates the View Protection Source environment. 'kPuppeteer' indicates the Cohesity's Remote Adapter. 'kPhysical' indicates the physical Protection Source environment. 'kPure' indicates the Pure Storage Protection Source environment. 'kNimble' indicates the Nimble Storage Protection Source environment. 'kHpe3Par' indicates the Hpe 3Par Storage Protection Source environment. 'kAzure' indicates the Microsoft's Azure Protection Source environment. 'kNetapp' indicates the Netapp Protection Source environment. 'kAgent' indicates the Agent Protection Source environment. 'kGenericNas' indicates the Generic Network Attached Storage Protection Source environment. 'kAcropolis' indicates the Acropolis Protection Source environment. 'kPhysicalFiles' indicates the Physical Files Protection Source environment. 'kIbmFlashSystem' indicates the IBM Flash System Protection Source environment. 'kIsilon' indicates the Dell EMC's Isilon Protection Source environment. 'kGPFS' indicates IBM's GPFS Protection Source environment. 'kKVM' indicates the KVM Protection Source environment. 'kAWS' indicates the AWS Protection Source environment. 'kExchange' indicates the Exchange Protection Source environment. 'kHyperVVSS' indicates the HyperV VSS Protection Source environment. 'kOracle' indicates the Oracle Protection Source environment. 'kGCP' indicates the Google Cloud Platform Protection Source environment. 'kFlashBlade' indicates the Flash Blade Protection Source environment. 'kAWSNative' indicates the AWS Native Protection Source environment. 'kO365' indicates the Office 365 Protection Source environment. 'kO365Outlook' indicates Office 365 outlook Protection Source environment. 'kHyperFlex' indicates the Hyper Flex Protection Source environment. 'kGCPNative' indicates the GCP Native Protection Source environment. 'kAzureNative' indicates the Azure Native Protection Source environment. 'kKubernetes' indicates a Kubernetes Protection Source environment. 'kElastifile' indicates Elastifile Protection Source environment. 'kAD' indicates Active Directory Protection Source environment. 'kRDSSnapshotManager' indicates AWS RDS Protection Source environment. 'kCassandra' indicates Cassandra Protection Source environment. 'kMongoDB' indicates MongoDB Protection Source environment. 'kCouchbase' indicates Couchbase Protection Source environment. 'kHdfs' indicates Hdfs Protection Source environment. 'kHive' indicates Hive Protection Source environment. 'kHBase' indicates HBase Protection Source environment. 'kUDA' indicates Universal Data Adapter Protection Source environment. 'kSAPHANA' indicates SAP HANA protection source environment. 'kO365Teams' indicates the Office365 Teams Protection Source environment. 'kO365Group' indicates the Office365 Groups Protection Source environment. 'kO365Exchange' indicates the Office365 Mailbox Protection Source environment. 'kO365OneDrive' indicates the Office365 OneDrive Protection Source environment. 'kO365Sharepoint' indicates the Office365 SharePoint Protection Source environment. 'kO365PublicFolders' indicates the Office365 PublicFolders Protection Source environment. kHpe3Par, kIbmFlashSystem, kAzure, kNetapp, kAgent, kGenericNas, kAcropolis, kPhysicalFiles, kIsilon, kGPFS, kKVM, kAWS, kExchange, kHyperVVSS, kOracle, kGCP, kFlashBlade, kAWSNative, kO365, kO365Outlook, kHyperFlex, kGCPNative, kAzureNative, kKubernetes, kElastifile, kAD, kRDSSnapshotManager, kCassandra, kMongoDB, kCouchbase, kHdfs, kHive, kHBase, kUDA, kSAPHANA, kO365Teams, kO365Group, kO365Exchange, kO365OneDrive, kO365Sharepoint, kO365PublicFolders.
			  * Constraints: Allowable list items are: `kVMware`, `kHyperV`, `kSQL`, `kView`, `kPuppeteer`, `kPhysical`, `kPure`, `kNimble`.
			* `password` - (String) Specifies password of the username to access the target source.
			* `throttling_config` - (List) Specifies the source side throttling configuration.
			Nested schema for **throttling_config**:
				* `cpu_throttling_config` - (List)
				Nested schema for **cpu_throttling_config**:
					* `fixed_threshold` - (Integer) Fixed baseline threshold for throttling. This is mandatory for any other throttling type than kNoThrottling.
					* `pattern_type` - (String) Type of the throttling pattern. 'kNoThrottling' indicates that throttling is not in force. 'kBaseThrottling' indicates indicates a constant base level throttling. 'kFixed' indicates a constant base level throttling.
					  * Constraints: Allowable values are: `kNoThrottling`, `kBaseThrottling`, `kFixed`.
					* `throttling_windows` - (List) Throttling windows which will be applicable in case of pattern_typec = kScheduleBased.
					Nested schema for **throttling_windows**:
						* `day_time_window` - (List) Specifies the Day Time Window Parameters.
						Nested schema for **day_time_window**:
							* `end_time` - (List) Specifies the Day Time Parameters.
							Nested schema for **end_time**:
								* `day` - (String) Specifies the day of the week (such as 'kMonday') for scheduling throttling. Specifies a day in a week such as 'kSunday', 'kMonday', etc.
								  * Constraints: Allowable values are: `kSunday`, `kMonday`, `kTuesday`, `kWednesday`, `kThursday`, `kFriday`, `kSaturday`.
								* `time` - (List) Specifies the time in hours and minutes.
								Nested schema for **time**:
									* `hour` - (Integer) Specifies the hour of this time.
									* `minute` - (Integer) Specifies the minute of this time.
							* `start_time` - (List) Specifies the Day Time Parameters.
							Nested schema for **start_time**:
								* `day` - (String) Specifies the day of the week (such as 'kMonday') for scheduling throttling. Specifies a day in a week such as 'kSunday', 'kMonday', etc.
								  * Constraints: Allowable values are: `kSunday`, `kMonday`, `kTuesday`, `kWednesday`, `kThursday`, `kFriday`, `kSaturday`.
								* `time` - (List) Specifies the time in hours and minutes.
								Nested schema for **time**:
									* `hour` - (Integer) Specifies the hour of this time.
									* `minute` - (Integer) Specifies the minute of this time.
						* `threshold` - (Integer) Throttling threshold applicable in the window.
				* `network_throttling_config` - (List)
				Nested schema for **network_throttling_config**:
					* `fixed_threshold` - (Integer) Fixed baseline threshold for throttling. This is mandatory for any other throttling type than kNoThrottling.
					* `pattern_type` - (String) Type of the throttling pattern. 'kNoThrottling' indicates that throttling is not in force. 'kBaseThrottling' indicates indicates a constant base level throttling. 'kFixed' indicates a constant base level throttling.
					  * Constraints: Allowable values are: `kNoThrottling`, `kBaseThrottling`, `kFixed`.
					* `throttling_windows` - (List) Throttling windows which will be applicable in case of pattern_typec = kScheduleBased.
					Nested schema for **throttling_windows**:
						* `day_time_window` - (List) Specifies the Day Time Window Parameters.
						Nested schema for **day_time_window**:
							* `end_time` - (List) Specifies the Day Time Parameters.
							Nested schema for **end_time**:
								* `day` - (String) Specifies the day of the week (such as 'kMonday') for scheduling throttling. Specifies a day in a week such as 'kSunday', 'kMonday', etc.
								  * Constraints: Allowable values are: `kSunday`, `kMonday`, `kTuesday`, `kWednesday`, `kThursday`, `kFriday`, `kSaturday`.
								* `time` - (List) Specifies the time in hours and minutes.
								Nested schema for **time**:
									* `hour` - (Integer) Specifies the hour of this time.
									* `minute` - (Integer) Specifies the minute of this time.
							* `start_time` - (List) Specifies the Day Time Parameters.
							Nested schema for **start_time**:
								* `day` - (String) Specifies the day of the week (such as 'kMonday') for scheduling throttling. Specifies a day in a week such as 'kSunday', 'kMonday', etc.
								  * Constraints: Allowable values are: `kSunday`, `kMonday`, `kTuesday`, `kWednesday`, `kThursday`, `kFriday`, `kSaturday`.
								* `time` - (List) Specifies the time in hours and minutes.
								Nested schema for **time**:
									* `hour` - (Integer) Specifies the hour of this time.
									* `minute` - (Integer) Specifies the minute of this time.
						* `threshold` - (Integer) Throttling threshold applicable in the window.
			* `username` - (String) Specifies username to access the target source.
		* `progress_monitor_path` - (String) Captures the current progress and pulse details w.r.t to either the registration or refresh.
		* `refresh_error_message` - (String) Specifies a message if there was any error encountered during the last rebuild of the Protection Source tree. If there was no error during the last rebuild, this field is reset.
		* `refresh_time_usecs` - (Integer) Specifies the Unix epoch time (in microseconds) when the Protection Source tree was most recently fetched and built.
		* `registered_apps_info` - (List) Specifies information of the applications registered on this protection source.
		Nested schema for **registered_apps_info**:
			* `authentication_error_message` - (String) pecifies an authentication error message. This indicates the given credentials are rejected and the registration of the application is not successful.
			* `authentication_status` - (String) Specifies the status of authenticating to the Protection Source when registering this application with Cohesity Cluster. If the status is 'kFinished' and there is no error, registration is successful. Specifies the status of the authentication during the registration of a Protection Source. 'kPending' indicates the authentication is in progress. 'kScheduled' indicates the authentication is scheduled. 'kFinished' indicates the authentication is completed. 'kRefreshInProgress' indicates the refresh is in progress.
			  * Constraints: Allowable values are: `kPending`, `kScheduled`, `kFinished`, `kRefreshInProgress`.
			* `environment` - (String) Specifies the application environment. Supported environment types such as 'kView', 'kSQL', 'kVMware', etc.
			  * Constraints: Allowable values are: `kPhysical`, `kPhysicalFiles`, `kSQL`, `kAgent`, `kVMware`, `kHyperV`, `kPure`, `kNimble`, `kView`, `kPuppeteer`.
			* `host_settings_check_results` - (List)
			Nested schema for **host_settings_check_results**:
				* `check_type` - (String) Specifies the type of the check internally performed. Specifies the type of the host check performed internally. 'kIsAgentPortAccessible' indicates the check for agent port access. 'kIsAgentRunning' indicates the status for the Cohesity agent service. 'kIsSQLWriterRunning' indicates the status for SQLWriter service. 'kAreSQLInstancesRunning' indicates the run status for all the SQL instances in the host. 'kCheckServiceLoginsConfig' checks the privileges and sysadmin status of the logins used by the SQL instance services, Cohesity agent service and the SQLWriter service. 'kCheckSQLFCIVIP' checks whether the SQL FCI is registered with a valid VIP or FQDN. 'kCheckSQLDiskSpace' checks whether volumes containing SQL DBs have at least 10% free space.
				  * Constraints: Allowable values are: `kIsAgentPortAccessible`, `kIsAgentRunning`, `kIsSQLWriterRunning`, `kAreSQLInstancesRunning`, `kCheckServiceLoginsConfig`, `kCheckSQLFCIVIP`, `kCheckSQLDiskSpace`.
				* `result_type` - (String) Specifies the type of the result returned after performing the internal host check. Specifies the type of the host check result performed internally. 'kPass' indicates that the respective check was successful. 'kFail' indicates that the respective check failed as some mandatory setting is not met 'kWarning' indicates that the respective check has warning as certain non-mandatory setting is not met.
				  * Constraints: Allowable values are: `kPass`, `kFail`, `kWarning`.
				* `user_message` - (String) Specifies a descriptive message for failed/warning types.
			* `refresh_error_message` - (String) Specifies a message if there was any error encountered during the last rebuild of the application tree. If there was no error during the last rebuild, this field is reset.
		* `registration_time_usecs` - (Integer) Specifies the Unix epoch time (in microseconds) when the Protection Source was registered.
		* `sfdc_params` - (List) Specifies an Object containing information about a registered Salesforce source.
		Nested schema for **sfdc_params**:
			* `access_token` - (String) Token that will be used in subsequent api requests.
			* `concurrent_api_requests_limit` - (Integer) Specifies the maximum number of concurrent API requests allowed for salesforce.
			* `consumer_key` - (String) Consumer key from the connected app in Sfdc.
			* `consumer_secret` - (String) Consumer secret from the connected app in Sfdc.
			* `daily_api_limit` - (Integer) Maximum daily api limit.
			* `endpoint` - (String) Sfdc Endpoint URL.
			* `endpoint_type` - (String) Specifies the Environment type for salesforce. 'PROD' 'SANDBOX' 'OTHER'.
			  * Constraints: Allowable values are: `PROD`, `SANDBOX`, `OTHER`.
			* `metadata_endpoint_url` - (String) Metadata endpoint url. All metadata requests must be made to this url.
			* `refresh_token` - (String) Token that will be used to refresh the access token.
			* `soap_endpoint_url` - (String) Soap endpoint url. All soap requests must be made to this url.
			* `use_bulk_api` - (Boolean) use bulk api if set to true.
		* `subnets` - (List) Specifies the list of subnets added during creation or updation of vmare source. Currently, this field will only be populated in case of VMware registration.
		Nested schema for **subnets**:
			* `component` - (String) Component that has reserved the subnet.
			* `description` - (String) Description of the subnet.
			* `id` - (Float) ID of the subnet.
			* `ip` - (String) Specifies either an IPv6 address or an IPv4 address.
			* `netmask_bits` - (Float) netmaskBits.
			* `netmask_ip4` - (String) Specifies the netmask using an IP4 address. The netmask can only be set using netmaskIp4 if the IP address is an IPv4 address.
			* `nfs_access` - (String) Component that has reserved the subnet.
			  * Constraints: Allowable values are: `kDisabled`, `kReadOnly`, `kReadWrite`.
			* `nfs_all_squash` - (Boolean) Specifies whether all clients from this subnet can map view with view_all_squash_uid/view_all_squash_gid configured in the view.
			* `nfs_root_squash` - (Boolean) Specifies whether clients from this subnet can mount as root on NFS.
			* `s3_access` - (String) Specifies whether clients from this subnet can access using S3 protocol. Protocol access level. 'kDisabled' indicates Protocol access level 'Disabled' 'kReadOnly' indicates Protocol access level 'ReadOnly' 'kReadWrite' indicates Protocol access level 'ReadWrite'.
			  * Constraints: Allowable values are: `kDisabled`, `kReadOnly`, `kReadWrite`.
			* `smb_access` - (String) Specifies whether clients from this subnet can mount using SMB protocol. Protocol access level. 'kDisabled' indicates Protocol access level 'Disabled' 'kReadOnly' indicates Protocol access level 'ReadOnly' 'kReadWrite' indicates Protocol access level 'ReadWrite'.
			  * Constraints: Allowable values are: `kDisabled`, `kReadOnly`, `kReadWrite`.
			* `tenant_id` - (String) Specifies the unique id of the tenant.
		* `throttling_policy` - (List) Specifies the throttling policy for a registered Protection Source.
		Nested schema for **throttling_policy**:
			* `enforce_max_streams` - (Boolean) Specifies whether datastore streams are configured for all datastores that are part of the registered entity. If set to true, number of streams from Cohesity cluster to the registered entity will be limited to the value set for maxConcurrentStreams. If not set or set to false, there is no max limit for the number of concurrent streams.
			* `enforce_registered_source_max_backups` - (Boolean) Specifies whether no. of backups are configured for the registered entity. If set to true, number of backups made by Cohesity cluster in the registered entity will be limited to the value set for RegisteredSourceMaxConcurrentBackups. If not set or set to false, there is no max limit for the number of concurrent backups.
			* `is_enabled` - (Boolean) Indicates whether read operations to the datastores, which are part of the registered Protection Source, are throttled.
			* `latency_thresholds` - (List) Specifies latency thresholds that trigger throttling for all datastores found in the registered Protection Source or specific to one datastore.
			Nested schema for **latency_thresholds**:
				* `active_task_msecs` - (Integer) If the latency of a datastore is above this value, existing backup tasks using the datastore are throttled.
				* `new_task_msecs` - (Integer) If the latency of a datastore is above this value, then new backup tasks using the datastore will not be started.
			* `max_concurrent_streams` - (Integer) Specifies the limit on the number of streams Cohesity cluster will make concurrently to the datastores of the registered entity. This limit is enforced only when the flag enforceMaxStreams is set to true.
			* `nas_source_params` - (List) Specifies the NAS specific source throttling parameters during source registration or during backup of the source.
			Nested schema for **nas_source_params**:
				* `max_parallel_metadata_fetch_full_percentage` - (Integer) Specifies the percentage value of maximum concurrent metadata to be fetched during full backup of the source.
				* `max_parallel_metadata_fetch_incremental_percentage` - (Integer) Specifies the percentage value of maximum concurrent metadata to be fetched during incremental backup of the source.
				* `max_parallel_read_write_full_percentage` - (Integer) Specifies the percentage value of maximum concurrent IO during full backup of the source.
				* `max_parallel_read_write_incremental_percentage` - (Integer) Specifies the percentage value of maximum concurrent IO during incremental backup of the source.
			* `registered_source_max_concurrent_backups` - (Integer) Specifies the limit on the number of backups Cohesity cluster will make concurrently to the registered entity. This limit is enforced only when the flag enforceRegisteredSourceMaxBackups is set to true.
			* `storage_array_snapshot_config` - (List)
			Nested schema for **storage_array_snapshot_config**:
				* `is_max_snapshots_config_enabled` - (Boolean) Specifies if the storage array snapshot max snapshots config is enabled or not.
				* `is_max_space_config_enabled` - (Boolean) Specifies if the storage array snapshot max space config is enabled or not.
				* `storage_array_snapshot_max_space_config` - (List)
				Nested schema for **storage_array_snapshot_max_space_config**:
					* `max_snapshot_space_percentage` - (Integer) Max number of storage snapshots allowed per volume/lun.
				* `storage_array_snapshot_throttling_policies` - (List) Specifies throttling policies configured for individual volume/lun.
				Nested schema for **storage_array_snapshot_throttling_policies**:
					* `id` - (Integer) Specifies the volume id of the storage array snapshot config.
					* `is_max_snapshots_config_enabled` - (Boolean) Specifies if the storage array snapshot max snapshots config is enabled or not.
					* `is_max_space_config_enabled` - (Boolean) Specifies if the storage array snapshot max space config is enabled or not.
					* `max_snapshot_config` - (List)
					Nested schema for **max_snapshot_config**:
						* `max_snapshots` - (Integer) Max number of storage snapshots allowed per volume/lun.
					* `max_space_config` - (List)
					Nested schema for **max_space_config**:
						* `max_snapshot_space_percentage` - (Integer) Max number of storage snapshots allowed per volume/lun.
		* `throttling_policy_overrides` - (List) Array of Throttling Policy Overrides for Datastores. Specifies a list of Throttling Policy for datastores that override the common throttling policy specified for the registered Protection Source. For datastores not in this list, common policy will still apply.
		Nested schema for **throttling_policy_overrides**:
			* `datastore_id` - (Integer) Specifies the Protection Source id of the Datastore.
			* `datastore_name` - (String) Specifies the display name of the Datastore.
			* `throttling_policy` - (List) Specifies the throttling policy for a registered Protection Source.
			Nested schema for **throttling_policy**:
				* `enforce_max_streams` - (Boolean) Specifies whether datastore streams are configured for all datastores that are part of the registered entity. If set to true, number of streams from Cohesity cluster to the registered entity will be limited to the value set for maxConcurrentStreams. If not set or set to false, there is no max limit for the number of concurrent streams.
				* `enforce_registered_source_max_backups` - (Boolean) Specifies whether no. of backups are configured for the registered entity. If set to true, number of backups made by Cohesity cluster in the registered entity will be limited to the value set for RegisteredSourceMaxConcurrentBackups. If not set or set to false, there is no max limit for the number of concurrent backups.
				* `is_enabled` - (Boolean) Indicates whether read operations to the datastores, which are part of the registered Protection Source, are throttled.
				* `latency_thresholds` - (List) Specifies latency thresholds that trigger throttling for all datastores found in the registered Protection Source or specific to one datastore.
				Nested schema for **latency_thresholds**:
					* `active_task_msecs` - (Integer) If the latency of a datastore is above this value, existing backup tasks using the datastore are throttled.
					* `new_task_msecs` - (Integer) If the latency of a datastore is above this value, then new backup tasks using the datastore will not be started.
				* `max_concurrent_streams` - (Integer) Specifies the limit on the number of streams Cohesity cluster will make concurrently to the datastores of the registered entity. This limit is enforced only when the flag enforceMaxStreams is set to true.
				* `nas_source_params` - (List) Specifies the NAS specific source throttling parameters during source registration or during backup of the source.
				Nested schema for **nas_source_params**:
					* `max_parallel_metadata_fetch_full_percentage` - (Integer) Specifies the percentage value of maximum concurrent metadata to be fetched during full backup of the source.
					* `max_parallel_metadata_fetch_incremental_percentage` - (Integer) Specifies the percentage value of maximum concurrent metadata to be fetched during incremental backup of the source.
					* `max_parallel_read_write_full_percentage` - (Integer) Specifies the percentage value of maximum concurrent IO during full backup of the source.
					* `max_parallel_read_write_incremental_percentage` - (Integer) Specifies the percentage value of maximum concurrent IO during incremental backup of the source.
				* `registered_source_max_concurrent_backups` - (Integer) Specifies the limit on the number of backups Cohesity cluster will make concurrently to the registered entity. This limit is enforced only when the flag enforceRegisteredSourceMaxBackups is set to true.
				* `storage_array_snapshot_config` - (List)
				Nested schema for **storage_array_snapshot_config**:
					* `is_max_snapshots_config_enabled` - (Boolean) Specifies if the storage array snapshot max snapshots config is enabled or not.
					* `is_max_space_config_enabled` - (Boolean) Specifies if the storage array snapshot max space config is enabled or not.
					* `storage_array_snapshot_max_space_config` - (List)
					Nested schema for **storage_array_snapshot_max_space_config**:
						* `max_snapshot_space_percentage` - (Integer) Max number of storage snapshots allowed per volume/lun.
					* `storage_array_snapshot_throttling_policies` - (List) Specifies throttling policies configured for individual volume/lun.
					Nested schema for **storage_array_snapshot_throttling_policies**:
						* `id` - (Integer) Specifies the volume id of the storage array snapshot config.
						* `is_max_snapshots_config_enabled` - (Boolean) Specifies if the storage array snapshot max snapshots config is enabled or not.
						* `is_max_space_config_enabled` - (Boolean) Specifies if the storage array snapshot max space config is enabled or not.
						* `max_snapshot_config` - (List)
						Nested schema for **max_snapshot_config**:
							* `max_snapshots` - (Integer) Max number of storage snapshots allowed per volume/lun.
						* `max_space_config` - (List)
						Nested schema for **max_space_config**:
							* `max_snapshot_space_percentage` - (Integer) Max number of storage snapshots allowed per volume/lun.
		* `uda_params` - (List) Specifies an Object containing information about a registered Universal Data Adapter source.
		Nested schema for **uda_params**:
			* `capabilities` - (List)
			Nested schema for **capabilities**:
				* `auto_log_backup` - (Boolean)
				* `dynamic_config` - (Boolean) Specifies whether the source supports the 'Dynamic Configuration' capability.
				* `entity_support` - (Boolean) Indicates if source has entity capability.
				* `et_log_backup` - (Boolean) Specifies whether the source supports externally triggered log backups.
				* `external_disks` - (Boolean) Only for sources in the cloud. A temporary external disk is provisoned in the cloud and mounted on the control node selected during backup / recovery for dump-sweep workflows that need a local disk to dump data. Prereq - non-mount, AGENT_ON_RIGEL.
				* `full_backup` - (Boolean)
				* `incr_backup` - (Boolean)
				* `log_backup` - (Boolean)
				* `multi_object_restore` - (Boolean) Whether the source supports restore of multiple objects.
				* `pause_resume_backup` - (Boolean)
				* `post_backup_job_script` - (Boolean) Triggers a post backup script on all nodes.
				* `post_restore_job_script` - (Boolean) Triggers a post restore script on all nodes.
				* `pre_backup_job_script` - (Boolean) Make a source call before actual start backup call.
				* `pre_restore_job_script` - (Boolean) Triggers a pre restore script on all nodes.
				* `resource_throttling` - (Boolean)
				* `snapfs_cert` - (Boolean)
			* `credentials` - (List) Specifies the object to hold username and password.
			Nested schema for **credentials**:
				* `password` - (String) Specifies the password to access target entity.
				* `username` - (String) Specifies the username to access target entity.
			* `et_enable_log_backup_policy` - (Boolean) Specifies whether to enable cohesity policy triggered log backups along with externally triggered backups. Only applicable if etLogBackup capability is true.
			* `et_enable_run_now` - (Boolean) Specifies if the user triggered runs are allowed along with externally triggered backups. Only applicable if etLogBackup is true.
			* `host_type` - (String) Specifies the environment type for the host.
			  * Constraints: Allowable values are: `kLinux`, `kWindows`, `kAix`, `kSolaris`, `kSapHana`, `kOther`, `kHPUX`, `kVOS`.
			* `hosts` - (List) List of hosts forming the Universal Data Adapter cluster.
			* `live_data_view` - (Boolean) Whether to use a live view for data backups.
			* `live_log_view` - (Boolean) Whether to use a live view for log backups.
			* `mount_dir` - (String) This field is deprecated and its value will be ignored. It was used to specify the absolute path on the host where the view would be mounted. This is now controlled by the agent configuration and is common for all the Universal Data Adapter sources. deprecated: true.
			* `mount_view` - (Boolean) Whether to mount a view during the source backup.
			* `script_dir` - (String) Path where various source scripts will be located.
			* `source_args` - (String) Custom arguments which will be provided to the source registration scripts. This is deprecated. Use 'sourceRegistrationArguments' instead.
			* `source_registration_arguments` - (List) Specifies a map of custom arguments to be supplied to the source registration scripts.
			Nested schema for **source_registration_arguments**:
				* `key` - (String)
				* `value` - (String)
			* `source_type` - (String) Global app source type.
		* `update_last_backup_details` - (Boolean) Specifies if the last backup time and status should be updated for the VMs protected from the vCenter.
		* `use_o_auth_for_exchange_online` - (Boolean) Specifies whether OAuth should be used for authentication in case of  Exchange Online.
		* `use_vm_bios_uuid` - (Boolean) Specifies if registered vCenter is using BIOS UUID to track virtual  machines.
		* `user_messages` - (List) Specifies the additional details encountered during registration. Though the registration may succeed, user messages imply the host environment requires some cleanup or fixing.
		* `username` - (String) Specifies username to access the target source.
		* `vlan_params` - (List) Specifies VLAN parameters for the restore operation.
		Nested schema for **vlan_params**:
			* `disable_vlan` - (Boolean) Specifies whether to use the VIPs even when VLANs are configured on the Cluster. If configured, VLAN IP addresses are used by default. If VLANs are not configured, this flag is ignored. Set this flag to true to force using the partition VIPs when VLANs are configured on the Cluster.
			* `interface_name` - (String) Specifies the physical interface group name to use for mounting Cohesity's view on the remote host. If specified, Cohesity hostname or the IP address on this VLAN is used.
			* `vlan` - (Integer) Specifies the VLAN to use for mounting Cohesity's view on the remote host. If specified, Cohesity hostname or the IP address on this VLAN is used.
		* `warning_messages` - (List) Specifies a list of warnings encountered during registration. Though the registration may succeed, warning messages imply the host environment requires some cleanup or fixing.
	* `root_node` - (List) Specifies the Protection Source for the root node of the Protection Source tree.
	Nested schema for **root_node**:
		* `connection_id` - (Integer) Specifies the connection id of the tenant.
		* `connector_group_id` - (Integer) Specifies the connector group id of the connector groups.
		* `custom_name` - (String) Specifies the user provided custom name of the Protection Source.
		* `environment` - (String) Specifies the environment (such as 'kVMware' or 'kSQL') where the Protection Source exists. Depending on the environment, one of the following Protection Sources are initialized.
		  * Constraints: Allowable values are: `kPhysical`, `kPhysicalFiles`, `kSQL`, `kAgent`.
		* `id` - (Integer) Specifies an id of the Protection Source.
		* `kubernetes_protection_source` - (List) Specifies a Protection Source in Kubernetes environment.
		Nested schema for **kubernetes_protection_source**:
			* `datamover_image_location` - (String) Specifies the location of Datamover image in private registry.
			* `datamover_service_type` - (Integer) Specifies Type of service to be deployed for communication with DataMover pods. Currently, LoadBalancer and NodePort are supported. [default = kNodePort].
			* `datamover_upgradability` - (Integer) Specifies if the deployed Datamover image needs to be upgraded for this kubernetes entity.
			* `default_vlan_params` - (List) Specifies VLAN parameters for the restore operation.
			Nested schema for **default_vlan_params**:
				* `disable_vlan` - (Boolean) Specifies whether to use the VIPs even when VLANs are configured on the Cluster. If configured, VLAN IP addresses are used by default. If VLANs are not configured, this flag is ignored. Set this flag to true to force using the partition VIPs when VLANs are configured on the Cluster.
				* `interface_name` - (String) Specifies the physical interface group name to use for mounting Cohesity's view on the remote host. If specified, Cohesity hostname or the IP address on this VLAN is used.
				* `vlan` - (Integer) Specifies the VLAN to use for mounting Cohesity's view on the remote host. If specified, Cohesity hostname or the IP address on this VLAN is used.
			* `description` - (String) Specifies an optional description of the object.
			* `distribution` - (String) Specifies the type of the entity in a Kubernetes environment. Determines the K8s distribution. kIKS, kROKS.
			  * Constraints: Allowable values are: `kMainline`, `kOpenshift`, `kRancher`, `kEKS`, `kGKE`, `kAKS`, `kVMwareTanzu`.
			* `init_container_image_location` - (String) Specifies the location of the image for init containers.
			* `label_attributes` - (List) Specifies the list of label attributes of this source.
			Nested schema for **label_attributes**:
				* `id` - (Integer) Specifies the Cohesity id of the K8s label.
				* `name` - (String) Specifies the appended key and value of the K8s label.
				* `uuid` - (String) Specifies Kubernetes Unique Identifier (UUID) of the K8s label.
			* `name` - (String) Specifies a unique name of the Protection Source.
			* `priority_class_name` - (String) Specifies the pritority class name during registration.
			* `resource_annotation_list` - (List) Specifies resource Annotations information provided during registration.
			Nested schema for **resource_annotation_list**:
				* `key` - (String) Key for label.
				* `value` - (String) Value for label.
			* `resource_label_list` - (List) Specifies resource labels information provided during registration.
			Nested schema for **resource_label_list**:
				* `key` - (String) Key for label.
				* `value` - (String) Value for label.
			* `san_field` - (List) Specifies the SAN field for agent certificate.
			* `service_annotations` - (List) Specifies annotations to be put on services for IP allocation. Applicable only when service is of type LoadBalancer.
			Nested schema for **service_annotations**:
				* `key` - (String)
				* `value` - (String)
			* `storage_class` - (List) Specifies storage class information of source.
			Nested schema for **storage_class**:
				* `name` - (String) Specifies name of storage class.
				* `provisioner` - (String) specifies provisioner of storage class.
			* `type` - (String) Specifies the type of the entity in a Kubernetes environment. Specifies the type of a Kubernetes Protection Source. 'kCluster' indicates a Kubernetes Cluster. 'kNamespace' indicates a namespace in a Kubernetes Cluster. 'kService' indicates a service running on a Kubernetes Cluster.
			  * Constraints: Allowable values are: `kCluster`, `kNamespace`, `kService`.
			* `uuid` - (String) Specifies the UUID of the object.
			* `velero_aws_plugin_image_location` - (String) Specifies the location of Velero AWS plugin image in private registry.
			* `velero_image_location` - (String) Specifies the location of Velero image in private registry.
			* `velero_openshift_plugin_image_location` - (String) Specifies the location of the image for openshift plugin container.
			* `velero_upgradability` - (String) Specifies if the deployed Velero image needs to be upgraded for this kubernetes entity.
			* `vlan_info_vec` - (List) Specifies VLAN information provided during registration.
			Nested schema for **vlan_info_vec**:
				* `service_annotations` - (List) Specifies annotations to be put on services for IP allocation. Applicable only when service is of type LoadBalancer.
				Nested schema for **service_annotations**:
					* `key` - (String) Specifies the service annotation key value.
					* `value` - (String) Specifies the service annotation value.
				* `vlan_params` - (List) Specifies VLAN params associated with the backup/restore operation.
				Nested schema for **vlan_params**:
					* `disable_vlan` - (Boolean) If this is set to true, then even if VLANs are configured on the system, the partition VIPs will be used for the restore.
					* `interface_name` - (String) Interface group to use for backup/restore. If this is not specified, primary interface group for the cluster will be used.
					* `vlan_id` - (Integer) If this is set, then the Cohesity host name or the IP address associated with this VLAN is used for mounting Cohesity's view on the remote host.
		* `name` - (String) Specifies a name of the Protection Source.
		* `parent_id` - (Integer) Specifies an id of the parent of the Protection Source.
		* `physical_protection_source` - (List) Specifies a Protection Source in a Physical environment.
		Nested schema for **physical_protection_source**:
			* `agents` - (List) Specifiles the agents running on the Physical Protection Source and the status information.
			Nested schema for **agents**:
				* `cbmr_version` - (String) Specifies the version if Cristie BMR product is installed on the host.
				* `file_cbt_info` - (List) CBT version and service state info.
				Nested schema for **file_cbt_info**:
					* `file_version` - (List) Subcomponent version. The interpretation of the version is based on operating system.
					Nested schema for **file_version**:
						* `build_ver` - (Float)
						* `major_ver` - (Float)
						* `minor_ver` - (Float)
						* `revision_num` - (Float)
					* `is_installed` - (Boolean) Indicates whether the cbt driver is installed.
					* `reboot_status` - (String) Indicates whether host is rebooted post VolCBT installation.
					  * Constraints: Allowable values are: `kRebooted`, `kNeedsReboot`, `kInternalError`.
					* `service_state` - (List) Structure to Hold Service Status.
					Nested schema for **service_state**:
						* `name` - (String)
						* `state` - (String)
				* `host_type` - (String) Specifies the host type where the agent is running. This is only set for persistent agents.
				  * Constraints: Allowable values are: `kLinux`, `kWindows`, `kAix`, `kSolaris`, `kSapHana`, `kSapOracle`, `kCockroachDB`, `kMySQL`, `kOther`, `kSapSybase`, `kSapMaxDB`, `kSapSybaseIQ`, `kDB2`, `kSapASE`, `kMariaDB`, `kPostgreSQL`, `kVOS`, `kHPUX`.
				* `id` - (Integer) Specifies the agent's id.
				* `name` - (String) Specifies the agent's name.
				* `oracle_multi_node_channel_supported` - (Boolean) Specifies whether oracle multi node multi channel is supported or not.
				* `registration_info` - (List) Specifies information about a registered Source.
				Nested schema for **registration_info**:
					* `access_info` - (List) Specifies the parameters required to establish a connection with a particular environment.
					Nested schema for **access_info**:
						* `connection_id` - (Integer) ID of the Bifrost (HyX or Rigel) network realm (i.e. a connection) associated with the source.
						* `connector_group_id` - (Integer) Specifies the Id of the connector group. Each connector group is collection of Rigel/hyx. Each entity will be tagged with connector group id.
						* `endpoint` - (String) Specify an IP address or URL of the environment. (such as the IP address of the vCenter Server for a VMware environment).
						* `environment` - (String) Specifies the environment like VMware, SQL, where the Protection Source exists. Supported environment types such as 'kView', 'kSQL', 'kVMware', etc.
						  * Constraints: Allowable values are: `kPhysical`, `kPhysicalFiles`, `kSQL`, `kAgent`.
						* `id` - (Integer) Specifies a Unique id that is generated when the Source is registered. This is a convenience field that is used to maintain an index to different connection params.
						* `version` - (Integer) Version is updated each time the connector parameters are updated. This is used to discard older connector parameters.
					* `allowed_ip_addresses` - (List) Specifies the list of IP Addresses on the registered source to be exclusively allowed for doing any type of IO operations.
					* `authentication_error_message` - (String) Specifies an authentication error message. This indicates the given credentials are rejected and the registration of the source is not successful.
					* `authentication_status` - (String) Specifies the status of the authenticating to the Protection Source when registering it with Cohesity Cluster.
					  * Constraints: Allowable values are: `kPending`, `kScheduled`, `kFinished`, `kRefreshInProgress`.
					* `blacklisted_ip_addresses` - (List) This field is deprecated. Use DeniedIpAddresses instead.
					* `denied_ip_addresses` - (List) Specifies the list of IP Addresses on the registered source to be denied for doing any type of IO operations.
					* `environments` - (List) Specifies a list of applications environment that are registered with this Protection Source such as 'kSQL'. Supported environment types such as 'kView', 'kSQL', 'kVMware', etc.
					  * Constraints: Allowable list items are: `kPhysical`, `kPhysicalFiles`, `kSQL`, `kAgent`.
					* `is_db_authenticated` - (Boolean) Specifies if application entity dbAuthenticated or not.
					* `is_storage_array_snapshot_enabled` - (Boolean) Specifies if this source entity has enabled storage array snapshot or not.
					* `link_vms_across_vcenter` - (Boolean) Specifies if the VM linking feature is enabled for this VCenter This means that VMs present in this VCenter which earlier belonged to some other VCenter(also registerd on same cluster) and were migrated, will be linked during EH refresh. This will enable preserving snapshot chains for migrated VMs.
					* `minimum_free_space_gb` - (Integer) Specifies the minimum free space in GiB of the space expected to be available on the datastore where the virtual disks of the VM being backed up. If the amount of free space(in GiB) is lower than the value given by this field, backup will be aborted. Note that this field is applicable only to 'kVMware' type of environments.
					* `minimum_free_space_percent` - (Integer) Specifies the minimum free space in percentage of the space expected to be available on the datastore where the virtual disks of the VM being backed up. If the amount of free space(in percentage) is lower than the value given by this field, backup will be aborted. Note that this field is applicable only to 'kVMware' type of environments.
					* `password` - (String) Specifies password of the username to access the target source.
					* `physical_params` - (List) Specifies the parameters required to register Application Servers running in a Protection Source specific to a physical adapter.
					Nested schema for **physical_params**:
						* `applications` - (List) Specifies the types of applications such as 'kSQL', 'kExchange', 'kAD' running on the Protection Source. Supported environment types such as 'kView', 'kSQL', 'kVMware', etc.
						  * Constraints: Allowable list items are: `kSQL`, `kOracle`.
						* `password` - (String) Specifies password of the username to access the target source.
						* `throttling_config` - (List) Specifies the source side throttling configuration.
						Nested schema for **throttling_config**:
							* `cpu_throttling_config` - (List) Specifies the Throttling Configuration Parameters.
							Nested schema for **cpu_throttling_config**:
								* `fixed_threshold` - (Integer) Fixed baseline threshold for throttling. This is mandatory for any other throttling type than kNoThrottling.
								* `pattern_type` - (String) Type of the throttling pattern. 'kNoThrottling' indicates that throttling is not in force. 'kBaseThrottling' indicates indicates a constant base level throttling. 'kFixed' indicates a constant base level throttling.
								  * Constraints: Allowable values are: `kNoThrottling`, `kBaseThrottling`, `kFixed`.
								* `throttling_windows` - (List)
								Nested schema for **throttling_windows**:
									* `day_time_window` - (List) Specifies the Day Time Window Parameters.
									Nested schema for **day_time_window**:
										* `end_time` - (List) Specifies the Day Time Parameters.
										Nested schema for **end_time**:
											* `day` - (String) Specifies the day of the week (such as 'kMonday') for scheduling throttling. Specifies a day in a week such as 'kSunday', 'kMonday', etc.
											  * Constraints: Allowable values are: `kSunday`, `kMonday`, `kTuesday`, `kWednesday`, `kThursday`, `kFriday`, `kSaturday`.
											* `time` - (List) Specifies the time in hours and minutes.
											Nested schema for **time**:
												* `hour` - (Integer) Specifies the hour of this time.
												* `minute` - (Integer) Specifies the minute of this time.
										* `start_time` - (List) Specifies the Day Time Parameters.
										Nested schema for **start_time**:
											* `day` - (String) Specifies the day of the week (such as 'kMonday') for scheduling throttling. Specifies a day in a week such as 'kSunday', 'kMonday', etc.
											  * Constraints: Allowable values are: `kSunday`, `kMonday`, `kTuesday`, `kWednesday`, `kThursday`, `kFriday`, `kSaturday`.
											* `time` - (List) Specifies the time in hours and minutes.
											Nested schema for **time**:
												* `hour` - (Integer) Specifies the hour of this time.
												* `minute` - (Integer) Specifies the minute of this time.
									* `threshold` - (Integer) Throttling threshold applicable in the window.
							* `network_throttling_config` - (List) Specifies the Throttling Configuration Parameters.
							Nested schema for **network_throttling_config**:
								* `fixed_threshold` - (Integer) Fixed baseline threshold for throttling. This is mandatory for any other throttling type than kNoThrottling.
								* `pattern_type` - (String) Type of the throttling pattern. 'kNoThrottling' indicates that throttling is not in force. 'kBaseThrottling' indicates indicates a constant base level throttling. 'kFixed' indicates a constant base level throttling.
								  * Constraints: Allowable values are: `kNoThrottling`, `kBaseThrottling`, `kFixed`.
								* `throttling_windows` - (List)
								Nested schema for **throttling_windows**:
									* `day_time_window` - (List) Specifies the Day Time Window Parameters.
									Nested schema for **day_time_window**:
										* `end_time` - (List) Specifies the Day Time Parameters.
										Nested schema for **end_time**:
											* `day` - (String) Specifies the day of the week (such as 'kMonday') for scheduling throttling. Specifies a day in a week such as 'kSunday', 'kMonday', etc.
											  * Constraints: Allowable values are: `kSunday`, `kMonday`, `kTuesday`, `kWednesday`, `kThursday`, `kFriday`, `kSaturday`.
											* `time` - (List) Specifies the time in hours and minutes.
											Nested schema for **time**:
												* `hour` - (Integer) Specifies the hour of this time.
												* `minute` - (Integer) Specifies the minute of this time.
										* `start_time` - (List) Specifies the Day Time Parameters.
										Nested schema for **start_time**:
											* `day` - (String) Specifies the day of the week (such as 'kMonday') for scheduling throttling. Specifies a day in a week such as 'kSunday', 'kMonday', etc.
											  * Constraints: Allowable values are: `kSunday`, `kMonday`, `kTuesday`, `kWednesday`, `kThursday`, `kFriday`, `kSaturday`.
											* `time` - (List) Specifies the time in hours and minutes.
											Nested schema for **time**:
												* `hour` - (Integer) Specifies the hour of this time.
												* `minute` - (Integer) Specifies the minute of this time.
									* `threshold` - (Integer) Throttling threshold applicable in the window.
						* `username` - (String) Specifies username to access the target source.
					* `progress_monitor_path` - (String) Captures the current progress and pulse details w.r.t to either the registration or refresh.
					* `refresh_error_message` - (String) Specifies a message if there was any error encountered during the last rebuild of the Protection Source tree. If there was no error during the last rebuild, this field is reset.
					* `refresh_time_usecs` - (Integer) Specifies the Unix epoch time (in microseconds) when the Protection Source tree was most recently fetched and built.
					* `registered_apps_info` - (List) Specifies information of the applications registered on this protection source.
					Nested schema for **registered_apps_info**:
						* `authentication_error_message` - (String) pecifies an authentication error message. This indicates the given credentials are rejected and the registration of the application is not successful.
						* `authentication_status` - (String) Specifies the status of authenticating to the Protection Source when registering this application with Cohesity Cluster. If the status is 'kFinished' and there is no error, registration is successful. Specifies the status of the authentication during the registration of a Protection Source. 'kPending' indicates the authentication is in progress. 'kScheduled' indicates the authentication is scheduled. 'kFinished' indicates the authentication is completed. 'kRefreshInProgress' indicates the refresh is in progress.
						  * Constraints: Allowable values are: `kPending`, `kScheduled`, `kFinished`, `kRefreshInProgress`.
						* `environment` - (String) Specifies the application environment. Supported environment types such as 'kView', 'kSQL', 'kVMware', etc.
						  * Constraints: Allowable values are: `kPhysical`, `kPhysicalFiles`, `kSQL`, `kAgent`, `kVMware`, `kHyperV`, `kPure`, `kNimble`, `kView`, `kPuppeteer`.
						* `host_settings_check_results` - (List)
						Nested schema for **host_settings_check_results**:
							* `check_type` - (String) Specifies the type of the check internally performed. Specifies the type of the host check performed internally. 'kIsAgentPortAccessible' indicates the check for agent port access. 'kIsAgentRunning' indicates the status for the Cohesity agent service. 'kIsSQLWriterRunning' indicates the status for SQLWriter service. 'kAreSQLInstancesRunning' indicates the run status for all the SQL instances in the host. 'kCheckServiceLoginsConfig' checks the privileges and sysadmin status of the logins used by the SQL instance services, Cohesity agent service and the SQLWriter service. 'kCheckSQLFCIVIP' checks whether the SQL FCI is registered with a valid VIP or FQDN. 'kCheckSQLDiskSpace' checks whether volumes containing SQL DBs have at least 10% free space.
							  * Constraints: Allowable values are: `kIsAgentPortAccessible`, `kIsAgentRunning`, `kIsSQLWriterRunning`, `kAreSQLInstancesRunning`, `kCheckServiceLoginsConfig`, `kCheckSQLFCIVIP`, `kCheckSQLDiskSpace`.
							* `result_type` - (String) Specifies the type of the result returned after performing the internal host check. Specifies the type of the host check result performed internally. 'kPass' indicates that the respective check was successful. 'kFail' indicates that the respective check failed as some mandatory setting is not met 'kWarning' indicates that the respective check has warning as certain non-mandatory setting is not met.
							  * Constraints: Allowable values are: `kPass`, `kFail`, `kWarning`.
							* `user_message` - (String) Specifies a descriptive message for failed/warning types.
						* `refresh_error_message` - (String) Specifies a message if there was any error encountered during the last rebuild of the application tree. If there was no error during the last rebuild, this field is reset.
					* `registration_time_usecs` - (Integer) Specifies the Unix epoch time (in microseconds) when the Protection Source was registered.
					* `subnets` - (List) Specifies the list of subnets added during creation or updation of vmare source. Currently, this field will only be populated in case of VMware registration.
					Nested schema for **subnets**:
						* `component` - (String) Component that has reserved the subnet.
						* `description` - (String) Description of the subnet.
						* `id` - (Float) ID of the subnet.
						* `ip` - (String) Specifies either an IPv6 address or an IPv4 address.
						* `netmask_bits` - (Float) netmaskBits.
						* `netmask_ip4` - (String) Specifies the netmask using an IP4 address. The netmask can only be set using netmaskIp4 if the IP address is an IPv4 address.
						* `nfs_access` - (String) Component that has reserved the subnet.
						  * Constraints: Allowable values are: `kDisabled`, `kReadOnly`, `kReadWrite`.
						* `nfs_all_squash` - (Boolean) Specifies whether all clients from this subnet can map view with view_all_squash_uid/view_all_squash_gid configured in the view.
						* `nfs_root_squash` - (Boolean) Specifies whether clients from this subnet can mount as root on NFS.
						* `s3_access` - (String) Specifies whether clients from this subnet can access using S3 protocol. Protocol access level. 'kDisabled' indicates Protocol access level 'Disabled' 'kReadOnly' indicates Protocol access level 'ReadOnly' 'kReadWrite' indicates Protocol access level 'ReadWrite'.
						  * Constraints: Allowable values are: `kDisabled`, `kReadOnly`, `kReadWrite`.
						* `smb_access` - (String) Specifies whether clients from this subnet can mount using SMB protocol. Protocol access level. 'kDisabled' indicates Protocol access level 'Disabled' 'kReadOnly' indicates Protocol access level 'ReadOnly' 'kReadWrite' indicates Protocol access level 'ReadWrite'.
						  * Constraints: Allowable values are: `kDisabled`, `kReadOnly`, `kReadWrite`.
						* `tenant_id` - (String) Specifies the unique id of the tenant.
					* `throttling_policy` - (List) Specifies the throttling policy for a registered Protection Source.
					Nested schema for **throttling_policy**:
						* `enforce_max_streams` - (Boolean) Specifies whether datastore streams are configured for all datastores that are part of the registered entity. If set to true, number of streams from Cohesity cluster to the registered entity will be limited to the value set for maxConcurrentStreams. If not set or set to false, there is no max limit for the number of concurrent streams.
						* `enforce_registered_source_max_backups` - (Boolean) Specifies whether no. of backups are configured for the registered entity. If set to true, number of backups made by Cohesity cluster in the registered entity will be limited to the value set for RegisteredSourceMaxConcurrentBackups. If not set or set to false, there is no max limit for the number of concurrent backups.
						* `is_enabled` - (Boolean) Indicates whether read operations to the datastores, which are part of the registered Protection Source, are throttled.
						* `latency_thresholds` - (List) Specifies latency thresholds that trigger throttling for all datastores found in the registered Protection Source or specific to one datastore.
						Nested schema for **latency_thresholds**:
							* `active_task_msecs` - (Integer) If the latency of a datastore is above this value, existing backup tasks using the datastore are throttled.
							* `new_task_msecs` - (Integer) If the latency of a datastore is above this value, then new backup tasks using the datastore will not be started.
						* `max_concurrent_streams` - (Float) Specifies the limit on the number of streams Cohesity cluster will make concurrently to the datastores of the registered entity. This limit is enforced only when the flag enforceMaxStreams is set to true.
						* `nas_source_params` - (List) Specifies the NAS specific source throttling parameters during source registration or during backup of the source.
						Nested schema for **nas_source_params**:
							* `max_parallel_metadata_fetch_full_percentage` - (Float) Specifies the percentage value of maximum concurrent metadata to be fetched during full backup of the source.
							* `max_parallel_metadata_fetch_incremental_percentage` - (Float) Specifies the percentage value of maximum concurrent metadata to be fetched during incremental backup of the source.
							* `max_parallel_read_write_full_percentage` - (Float) Specifies the percentage value of maximum concurrent IO during full backup of the source.
							* `max_parallel_read_write_incremental_percentage` - (Float) Specifies the percentage value of maximum concurrent IO during incremental backup of the source.
						* `registered_source_max_concurrent_backups` - (Float) Specifies the limit on the number of backups Cohesity cluster will make concurrently to the registered entity. This limit is enforced only when the flag enforceRegisteredSourceMaxBackups is set to true.
						* `storage_array_snapshot_config` - (List) Specifies Storage Array Snapshot Configuration.
						Nested schema for **storage_array_snapshot_config**:
							* `is_max_snapshots_config_enabled` - (Boolean) Specifies if the storage array snapshot max snapshots config is enabled or not.
							* `is_max_space_config_enabled` - (Boolean) Specifies if the storage array snapshot max space config is enabled or not.
							* `storage_array_snapshot_max_space_config` - (List) Specifies Storage Array Snapshot Max Space Config.
							Nested schema for **storage_array_snapshot_max_space_config**:
								* `max_snapshot_space_percentage` - (Float) Max number of storage snapshots allowed per volume/lun.
							* `storage_array_snapshot_throttling_policies` - (List) Specifies throttling policies configured for individual volume/lun.
							Nested schema for **storage_array_snapshot_throttling_policies**:
								* `id` - (Integer) Specifies the volume id of the storage array snapshot config.
								* `is_max_snapshots_config_enabled` - (Boolean) Specifies if the storage array snapshot max snapshots config is enabled or not.
								* `is_max_space_config_enabled` - (Boolean) Specifies if the storage array snapshot max space config is enabled or not.
								* `max_snapshot_config` - (List) Specifies Storage Array Snapshot Max Snapshots Config.
								Nested schema for **max_snapshot_config**:
									* `max_snapshots` - (Float) Max number of storage snapshots allowed per volume/lun.
								* `max_space_config` - (List) Specifies Storage Array Snapshot Max Space Config.
								Nested schema for **max_space_config**:
									* `max_snapshot_space_percentage` - (Float) Max number of storage snapshots allowed per volume/lun.
					* `throttling_policy_overrides` - (List)
					Nested schema for **throttling_policy_overrides**:
						* `datastore_id` - (Integer) Specifies the Protection Source id of the Datastore.
						* `datastore_name` - (String) Specifies the display name of the Datastore.
						* `throttling_policy` - (List) Specifies the throttling policy for a registered Protection Source.
						Nested schema for **throttling_policy**:
							* `enforce_max_streams` - (Boolean) Specifies whether datastore streams are configured for all datastores that are part of the registered entity. If set to true, number of streams from Cohesity cluster to the registered entity will be limited to the value set for maxConcurrentStreams. If not set or set to false, there is no max limit for the number of concurrent streams.
							* `enforce_registered_source_max_backups` - (Boolean) Specifies whether no. of backups are configured for the registered entity. If set to true, number of backups made by Cohesity cluster in the registered entity will be limited to the value set for RegisteredSourceMaxConcurrentBackups. If not set or set to false, there is no max limit for the number of concurrent backups.
							* `is_enabled` - (Boolean) Indicates whether read operations to the datastores, which are part of the registered Protection Source, are throttled.
							* `latency_thresholds` - (List) Specifies latency thresholds that trigger throttling for all datastores found in the registered Protection Source or specific to one datastore.
							Nested schema for **latency_thresholds**:
								* `active_task_msecs` - (Integer) If the latency of a datastore is above this value, existing backup tasks using the datastore are throttled.
								* `new_task_msecs` - (Integer) If the latency of a datastore is above this value, then new backup tasks using the datastore will not be started.
							* `max_concurrent_streams` - (Float) Specifies the limit on the number of streams Cohesity cluster will make concurrently to the datastores of the registered entity. This limit is enforced only when the flag enforceMaxStreams is set to true.
							* `nas_source_params` - (List) Specifies the NAS specific source throttling parameters during source registration or during backup of the source.
							Nested schema for **nas_source_params**:
								* `max_parallel_metadata_fetch_full_percentage` - (Float) Specifies the percentage value of maximum concurrent metadata to be fetched during full backup of the source.
								* `max_parallel_metadata_fetch_incremental_percentage` - (Float) Specifies the percentage value of maximum concurrent metadata to be fetched during incremental backup of the source.
								* `max_parallel_read_write_full_percentage` - (Float) Specifies the percentage value of maximum concurrent IO during full backup of the source.
								* `max_parallel_read_write_incremental_percentage` - (Float) Specifies the percentage value of maximum concurrent IO during incremental backup of the source.
							* `registered_source_max_concurrent_backups` - (Float) Specifies the limit on the number of backups Cohesity cluster will make concurrently to the registered entity. This limit is enforced only when the flag enforceRegisteredSourceMaxBackups is set to true.
							* `storage_array_snapshot_config` - (List) Specifies Storage Array Snapshot Configuration.
							Nested schema for **storage_array_snapshot_config**:
								* `is_max_snapshots_config_enabled` - (Boolean) Specifies if the storage array snapshot max snapshots config is enabled or not.
								* `is_max_space_config_enabled` - (Boolean) Specifies if the storage array snapshot max space config is enabled or not.
								* `storage_array_snapshot_max_space_config` - (List) Specifies Storage Array Snapshot Max Space Config.
								Nested schema for **storage_array_snapshot_max_space_config**:
									* `max_snapshot_space_percentage` - (Float) Max number of storage snapshots allowed per volume/lun.
								* `storage_array_snapshot_throttling_policies` - (List) Specifies throttling policies configured for individual volume/lun.
								Nested schema for **storage_array_snapshot_throttling_policies**:
									* `id` - (Integer) Specifies the volume id of the storage array snapshot config.
									* `is_max_snapshots_config_enabled` - (Boolean) Specifies if the storage array snapshot max snapshots config is enabled or not.
									* `is_max_space_config_enabled` - (Boolean) Specifies if the storage array snapshot max space config is enabled or not.
									* `max_snapshot_config` - (List) Specifies Storage Array Snapshot Max Snapshots Config.
									Nested schema for **max_snapshot_config**:
										* `max_snapshots` - (Float) Max number of storage snapshots allowed per volume/lun.
									* `max_space_config` - (List) Specifies Storage Array Snapshot Max Space Config.
									Nested schema for **max_space_config**:
										* `max_snapshot_space_percentage` - (Float) Max number of storage snapshots allowed per volume/lun.
					* `use_o_auth_for_exchange_online` - (Boolean) Specifies whether OAuth should be used for authentication in case of Exchange Online.
					* `use_vm_bios_uuid` - (Boolean) Specifies if registered vCenter is using BIOS UUID to track virtual machines.
					* `user_messages` - (List) Specifies the additional details encountered during registration. Though the registration may succeed, user messages imply the host environment requires some cleanup or fixing.
					* `username` - (String) Specifies username to access the target source.
					* `vlan_params` - (List) Specifies the VLAN configuration for Recovery.
					Nested schema for **vlan_params**:
						* `disable_vlan` - (Boolean) Specifies whether to use the VIPs even when VLANs are configured on the Cluster. If configured, VLAN IP addresses are used by default. If VLANs are not configured, this flag is ignored. Set this flag to true to force using the partition VIPs when VLANs are configured on the Cluster.
						* `interface_name` - (String) Specifies the physical interface group name to use for mounting Cohesity's view on the remote host. If specified, Cohesity hostname or the IP address on this VLAN is used.
						* `vlan` - (Float) Specifies the VLAN to use for mounting Cohesity's view on the remote host. If specified, Cohesity hostname or the IP address on this VLAN is used.
					* `warning_messages` - (List) Specifies a list of warnings encountered during registration. Though the registration may succeed, warning messages imply the host environment requires some cleanup or fixing.
				* `source_side_dedup_enabled` - (Boolean) Specifies whether source side dedup is enabled or not.
				* `status` - (String) Specifies the agent status. Specifies the status of the agent running on a physical source.
				  * Constraints: Allowable values are: `kUnknown`, `kUnreachable`, `kHealthy`, `kDegraded`.
				* `status_message` - (String) Specifies additional details about the agent status.
				* `upgradability` - (String) Specifies the upgradability of the agent running on the physical server. Specifies the upgradability of the agent running on the physical server.
				  * Constraints: Allowable values are: `kUpgradable`, `kCurrent`, `kUnknown`, `kNonUpgradableInvalidVersion`, `kNonUpgradableAgentIsNewer`, `kNonUpgradableAgentIsOld`.
				* `upgrade_status` - (String) Specifies the status of the upgrade of the agent on a physical server. Specifies the status of the upgrade of the agent on a physical server.
				  * Constraints: Allowable values are: `kIdle`, `kAccepted`, `kStarted`, `kFinished`, `kScheduled`.
				* `upgrade_status_message` - (String) Specifies detailed message about the agent upgrade failure. This field is not set for successful upgrade.
				* `version` - (String) Specifies the version of the Agent software.
				* `vol_cbt_info` - (List) CBT version and service state info.
				Nested schema for **vol_cbt_info**:
					* `file_version` - (List) Subcomponent version. The interpretation of the version is based on operating system.
					Nested schema for **file_version**:
						* `build_ver` - (Float)
						* `major_ver` - (Float)
						* `minor_ver` - (Float)
						* `revision_num` - (Float)
					* `is_installed` - (Boolean) Indicates whether the cbt driver is installed.
					* `reboot_status` - (String) Indicates whether host is rebooted post VolCBT installation.
					  * Constraints: Allowable values are: `kRebooted`, `kNeedsReboot`, `kInternalError`.
					* `service_state` - (List) Structure to Hold Service Status.
					Nested schema for **service_state**:
						* `name` - (String)
						* `state` - (String)
			* `cluster_source_type` - (String) Specifies the type of cluster resource this source represents.
			* `host_name` - (String) Specifies the hostname.
			* `host_type` - (String) Specifies the environment type for the host.
			  * Constraints: Allowable values are: `kLinux`, `kWindows`, `kAix`, `kSolaris`, `kSapHana`, `kSapOracle`, `kCockroachDB`, `kMySQL`, `kOther`, `kSapSybase`, `kSapMaxDB`, `kSapSybaseIQ`, `kDB2`, `kSapASE`, `kMariaDB`, `kPostgreSQL`, `kVOS`, `kHPUX`.
			* `id` - (List) Specifies an id for an object that is unique across Cohesity Clusters. The id is composite of all the ids listed below.
			Nested schema for **id**:
				* `cluster_id` - (Integer) Specifies the Cohesity Cluster id where the object was created.
				* `cluster_incarnation_id` - (Integer) Specifies an id for the Cohesity Cluster that is generated when a Cohesity Cluster is initially created.
				* `id` - (Integer) Specifies a unique id assigned to an object (such as a Job) by the Cohesity Cluster.
			* `is_proxy_host` - (Boolean) Specifies if the physical host is a proxy host.
			* `memory_size_bytes` - (Integer) Specifies the total memory on the host in bytes.
			* `name` - (String) Specifies a human readable name of the Protection Source.
			* `networking_info` - (List) Specifies the struct containing information about network addresses configured on the given box. This is needed for dealing with Windows/Oracle Cluster resources that we discover and protect automatically.
			Nested schema for **networking_info**:
				* `resource_vec` - (List) The list of resources on the system that are accessible by an IP address.
				Nested schema for **resource_vec**:
					* `endpoints` - (List) The endpoints by which the resource is accessible.
					Nested schema for **endpoints**:
						* `fqdn` - (String) The Fully Qualified Domain Name.
						* `ipv4_addr` - (String) The IPv4 address.
						* `ipv6_addr` - (String) The IPv6 address.
					* `type` - (String) The type of the resource.
			* `num_processors` - (Integer) Specifies the number of processors on the host.
			* `os_name` - (String) Specifies a human readable name of the OS of the Protection Source.
			* `type` - (String) Specifies the type of managed Object in a Physical Protection Source. 'kGroup' indicates the EH container.
			  * Constraints: Allowable values are: `kGroup`, `kHost`, `kWindowsCluster`, `kOracleRACCluster`, `kOracleAPCluster`.
			* `vcs_version` - (String) Specifies cluster version for VCS host.
			* `volumes` - (List) Array of Physical Volumes. Specifies the volumes available on the physical host. These fields are populated only for the kPhysicalHost type.
			Nested schema for **volumes**:
				* `device_path` - (String) Specifies the path to the device that hosts the volume locally.
				* `guid` - (String) Specifies an id for the Physical Volume.
				* `is_boot_volume` - (Boolean) Specifies whether the volume is boot volume.
				* `is_extended_attributes_supported` - (Boolean) Specifies whether this volume supports extended attributes (like ACLs) when performing file backups.
				* `is_protected` - (Boolean) Specifies if a volume is protected by a Job.
				* `is_shared_volume` - (Boolean) Specifies whether the volume is shared volume.
				* `label` - (String) Specifies a volume label that can be used for displaying additional identifying information about a volume.
				* `logical_size_bytes` - (Float) Specifies the logical size of the volume in bytes that is not reduced by change-block tracking, compression and deduplication.
				* `mount_points` - (List) Specifies the mount points where the volume is mounted, for example- 'C:', '/mnt/foo' etc.
				* `mount_type` - (String) Specifies mount type of volume e.g. nfs, autofs, ext4 etc.
				* `network_path` - (String) Specifies the full path to connect to the network attached volume. For example, (IP or hostname):/path/to/share for NFS volumes).
				* `used_size_bytes` - (Float) Specifies the size used by the volume in bytes.
			* `vsswriters` - (List)
			Nested schema for **vsswriters**:
				* `is_writer_excluded` - (Boolean) If true, the writer will be excluded by default.
				* `writer_name` - (Boolean) Specifies the name of the writer.
		* `sql_protection_source` - (List) Specifies an Object representing one SQL Server instance or database.
		Nested schema for **sql_protection_source**:
			* `created_timestamp` - (String) Specifies the time when the database was created. It is displayed in the timezone of the SQL server on which this database is running.
			* `database_name` - (String) Specifies the database name of the SQL Protection Source, if the type is database.
			* `db_aag_entity_id` - (Integer) Specifies the AAG entity id if the database is part of an AAG. This field is set only for type 'kDatabase'.
			* `db_aag_name` - (String) Specifies the name of the AAG if the database is part of an AAG. This field is set only for type 'kDatabase'.
			* `db_compatibility_level` - (Integer) Specifies the versions of SQL server that the database is compatible with.
			* `db_file_groups` - (List) Specifies the information about the set of file groups for this db on the host. This is only set if the type is kDatabase.
			* `db_files` - (List) Specifies the last known information about the set of database files on the host. This field is set only for type 'kDatabase'.
			Nested schema for **db_files**:
				* `file_type` - (String) Specifies the format type of the file that SQL database stores the data. Specifies the format type of the file that SQL database stores the data. 'kRows' refers to a data file 'kLog' refers to a log file 'kFileStream' refers to a directory containing FILESTREAM data 'kNotSupportedType' is for information purposes only. Not supported. 'kFullText' refers to a full-text catalog.
				  * Constraints: Allowable values are: `kRows`, `kLog`, `kFileStream`, `kNotSupportedType`, `kFullText`.
				* `full_path` - (String) Specifies the full path of the database file on the SQL host machine.
				* `size_bytes` - (Integer) Specifies the last known size of the database file.
			* `db_owner_username` - (String) Specifies the name of the database owner.
			* `default_database_location` - (String) Specifies the default path for data files for DBs in an instance.
			* `default_log_location` - (String) Specifies the default path for log files for DBs in an instance.
			* `id` - (List) Specifies a unique id for a SQL Protection Source.
			Nested schema for **id**:
				* `created_date_msecs` - (Integer) Specifies a unique identifier generated from the date the database is created or renamed. Cohesity uses this identifier in combination with the databaseId to uniquely identify a database.
				* `database_id` - (Integer) Specifies a unique id of the database but only for the life of the database. SQL Server may reuse database ids. Cohesity uses the createDateMsecs in combination with this databaseId to uniquely identify a database.
				* `instance_id` - (String) Specifies unique id for the SQL Server instance. This id does not change during the life of the instance.
			* `is_available_for_vss_backup` - (Boolean) Specifies whether the database is marked as available for backup according to the SQL Server VSS writer. This may be false if either the state of the databases is not online, or if the VSS writer is not online. This field is set only for type 'kDatabase'.
			* `is_encrypted` - (Boolean) Specifies whether the database is TDE enabled.
			* `name` - (String) Specifies the instance name of the SQL Protection Source.
			* `owner_id` - (Integer) Specifies the id of the container VM for the SQL Protection Source.
			* `recovery_model` - (String) Specifies the Recovery Model for the database in SQL environment. Only meaningful for the 'kDatabase' SQL Protection Source. Specifies the Recovery Model set for the Microsoft SQL Server. 'kSimpleRecoveryModel' indicates the Simple SQL Recovery Model which does not utilize log backups. 'kFullRecoveryModel' indicates the Full SQL Recovery Model which requires log backups and allows recovery to a single point in time. 'kBulkLoggedRecoveryModel' indicates the Bulk Logged SQL Recovery Model which requires log backups and allows high-performance bulk copy operations.
			  * Constraints: Allowable values are: `kSimpleRecoveryModel`, `kFullRecoveryModel`, `kBulkLoggedRecoveryModel`.
			* `sql_server_db_state` - (String) The state of the database as returned by SQL Server. Indicates the state of the database. The values correspond to the 'state' field in the system table sys.databases. See https://goo.gl/P66XqM. 'kOnline' indicates that database is in online state. 'kRestoring' indicates that database is in restore state. 'kRecovering' indicates that database is in recovery state. 'kRecoveryPending' indicates that database recovery is in pending state. 'kSuspect' indicates that primary filegroup is suspect and may be damaged. 'kEmergency' indicates that manually forced emergency state. 'kOffline' indicates that database is in offline state. 'kCopying' indicates that database is in copying state. 'kOfflineSecondary' indicates that secondary database is in offline state.
			  * Constraints: Allowable values are: `kOnline`, `kRestoring`, `kRecovering`, `kRecoveryPending`, `kSuspect`, `kEmergency`, `kOffline`, `kCopying`, `kOfflineSecondary`.
			* `sql_server_instance_version` - (List) Specifies the Server Instance Version.
			Nested schema for **sql_server_instance_version**:
				* `build` - (Float) Specifies the build.
				* `major_version` - (Float) Specifies the major version.
				* `minor_version` - (Float) Specifies the minor version.
				* `revision` - (Float) Specifies the revision.
				* `version_string` - (Float) Specifies the version string.
			* `type` - (String) Specifies the type of the managed Object in a SQL Protection Source. Examples of SQL Objects include 'kInstance' and 'kDatabase'. 'kInstance' indicates that SQL server instance is being protected. 'kDatabase' indicates that SQL server database is being protected. 'kAAG' indicates that SQL AAG (AlwaysOn Availability Group) is being protected. 'kAAGRootContainer' indicates that SQL AAG's root container is being protected. 'kRootContainer' indicates root container for SQL sources.
			  * Constraints: Allowable values are: `kInstance`, `kDatabase`, `kAAG`, `kAAGRootContainer`, `kRootContainer`.
	* `stats` - (List) Specifies the stats of protection for a Protection Source Tree.
	Nested schema for **stats**:
		* `protected_count` - (Integer) Specifies the number of objects that are protected under the given entity.
		* `protected_size` - (Integer) Specifies the total size of the protected objects under the given entity.
		* `unprotected_count` - (Integer) Specifies the number of objects that are not protected under the given entity.
		* `unprotected_size` - (Integer) Specifies the total size of the unprotected objects under the given entity.
	* `stats_by_env` - (List) Specifies the breakdown of the stats of protection by environment. overrideDescription: true.
	Nested schema for **stats_by_env**:
		* `environment` - (String) Specifies the type of environment of the source object like kSQL etc.  Supported environment types such as 'kView', 'kSQL', 'kVMware', etc. NOTE: 'kPuppeteer' refers to Cohesity's Remote Adapter. 'kVMware' indicates the VMware Protection Source environment. 'kHyperV' indicates the HyperV Protection Source environment. 'kSQL' indicates the SQL Protection Source environment. 'kView' indicates the View Protection Source environment. 'kPuppeteer' indicates the Cohesity's Remote Adapter. 'kPhysical' indicates the physical Protection Source environment. 'kPure' indicates the Pure Storage Protection Source environment. 'kNimble' indicates the Nimble Storage Protection Source environment. 'kHpe3Par' indicates the Hpe 3Par Storage Protection Source environment. 'kAzure' indicates the Microsoft's Azure Protection Source environment. 'kNetapp' indicates the Netapp Protection Source environment. 'kAgent' indicates the Agent Protection Source environment. 'kGenericNas' indicates the Generic Network Attached Storage Protection Source environment. 'kAcropolis' indicates the Acropolis Protection Source environment. 'kPhysicalFiles' indicates the Physical Files Protection Source environment. 'kIbmFlashSystem' indicates the IBM Flash System Protection Source environment. 'kIsilon' indicates the Dell EMC's Isilon Protection Source environment. 'kGPFS' indicates IBM's GPFS Protection Source environment. 'kKVM' indicates the KVM Protection Source environment. 'kAWS' indicates the AWS Protection Source environment. 'kExchange' indicates the Exchange Protection Source environment. 'kHyperVVSS' indicates the HyperV VSS Protection Source environment. 'kOracle' indicates the Oracle Protection Source environment. 'kGCP' indicates the Google Cloud Platform Protection Source environment. 'kFlashBlade' indicates the Flash Blade Protection Source environment. 'kAWSNative' indicates the AWS Native Protection Source environment. 'kO365' indicates the Office 365 Protection Source environment. 'kO365Outlook' indicates Office 365 outlook Protection Source environment. 'kHyperFlex' indicates the Hyper Flex Protection Source environment. 'kGCPNative' indicates the GCP Native Protection Source environment. 'kAzureNative' indicates the Azure Native Protection Source environment. 'kKubernetes' indicates a Kubernetes Protection Source environment. 'kElastifile' indicates Elastifile Protection Source environment. 'kAD' indicates Active Directory Protection Source environment. 'kRDSSnapshotManager' indicates AWS RDS Protection Source environment. 'kCassandra' indicates Cassandra Protection Source environment. 'kMongoDB' indicates MongoDB Protection Source environment. 'kCouchbase' indicates Couchbase Protection Source environment. 'kHdfs' indicates Hdfs Protection Source environment. 'kHive' indicates Hive Protection Source environment. 'kHBase' indicates HBase Protection Source environment. 'kUDA' indicates Universal Data Adapter Protection Source environment. 'kSAPHANA' indicates SAP HANA protection source environment. 'kO365Teams' indicates the Office365 Teams Protection Source environment. 'kO365Group' indicates the Office365 Groups Protection Source environment. 'kO365Exchange' indicates the Office365 Mailbox Protection Source environment. 'kO365OneDrive' indicates the Office365 OneDrive Protection Source environment. 'kO365Sharepoint' indicates the Office365 SharePoint Protection Source environment. 'kO365PublicFolders' indicates the Office365 PublicFolders ProtectionSource environment. kHpe3Par, kIbmFlashSystem, kAzure, kNetapp, kAgent, kGenericNas, kAcropolis, kPhysicalFiles, kIsilon, kGPFS, kKVM, kAWS, kExchange, kHyperVVSS, kOracle, kGCP, kFlashBlade, kAWSNative, kO365, kO365Outlook, kHyperFlex, kGCPNative, kAzureNative, kKubernetes, kElastifile, kAD, kRDSSnapshotManager, kCassandra, kMongoDB, kCouchbase, kHdfs, kHive, kHBase, kUDA, kSAPHANA, kO365Teams, kO365Group, kO365Exchange, kO365OneDrive, kO365Sharepoint, kO365PublicFolders.
		  * Constraints: Allowable values are: `kVMware`, `kHyperV`, `kSQL`, `kView`, `kPuppeteer`, `kPhysical`, `kPure`, `kNimble`.
		* `kubernetes_distribution_stats` - (List) Specifies the breakdown of the kubernetes clusters by distribution type.
		Nested schema for **kubernetes_distribution_stats**:
			* `distribution` - (String) Specifies the type of Kuberentes distribution Determines the K8s distribution. kIKS, kROKS.
			  * Constraints: Allowable values are: `kMainline`, `kOpenshift`, `kRancher`, `kEKS`, `kGKE`, `kAKS`, `kVMwareTanzu`.
			* `protected_count` - (Integer) Specifies the number of objects that are protected for that distribution.
			* `protected_size` - (Integer) Specifies the total size of objects that are protected for that distribution.
			* `total_registered_clusters` - (Integer) Specifies the number of registered clusters for that distribution.
			* `unprotected_count` - (Integer) Specifies the number of objects that are not protected for that distribution.
			* `unprotected_size` - (Integer) Specifies the total size of objects that are not protected for that distribution.
		* `protected_count` - (Integer) Specifies the number of objects that are protected under the given entity.
		* `protected_size` - (Integer) Specifies the total size of the protected objects under the given entity.
		* `unprotected_count` - (Integer) Specifies the number of objects that are not protected under the given entity.
		* `unprotected_size` - (Integer) Specifies the total size of the unprotected objects under the given entity.
	* `total_downtiered_size_in_bytes` - (Integer) Specifies the total bytes downtiered from the source so far.
	* `total_uptiered_size_in_bytes` - (Integer) Specifies the total bytes uptiered to the source so far.
* `stats` - (List) Specifies the sum of all the stats of protection of Protection Sources and views selected by the query parameters.
Nested schema for **stats**:
	* `protected_count` - (Integer) Specifies the number of objects that are protected under the given entity.
	* `protected_size` - (Integer) Specifies the total size of the protected objects under the given entity.
	* `unprotected_count` - (Integer) Specifies the number of objects that are not protected under the given entity.
	* `unprotected_size` - (Integer) Specifies the total size of the unprotected objects under the given entity.
* `stats_by_env` - (List) Specifies the breakdown of the stats by environment overrideDescription: true.
Nested schema for **stats_by_env**:
	* `environment` - (String) Specifies the type of environment of the source object like kSQL etc.  Supported environment types such as 'kView', 'kSQL', 'kVMware', etc. NOTE: 'kPuppeteer' refers to Cohesity's Remote Adapter. 'kVMware' indicates the VMware Protection Source environment. 'kHyperV' indicates the HyperV Protection Source environment. 'kSQL' indicates the SQL Protection Source environment. 'kView' indicates the View Protection Source environment. 'kPuppeteer' indicates the Cohesity's Remote Adapter. 'kPhysical' indicates the physical Protection Source environment. 'kPure' indicates the Pure Storage Protection Source environment. 'kNimble' indicates the Nimble Storage Protection Source environment. 'kHpe3Par' indicates the Hpe 3Par Storage Protection Source environment. 'kAzure' indicates the Microsoft's Azure Protection Source environment. 'kNetapp' indicates the Netapp Protection Source environment. 'kAgent' indicates the Agent Protection Source environment. 'kGenericNas' indicates the Generic Network Attached Storage Protection Source environment. 'kAcropolis' indicates the Acropolis Protection Source environment. 'kPhysicalFiles' indicates the Physical Files Protection Source environment. 'kIbmFlashSystem' indicates the IBM Flash System Protection Source environment. 'kIsilon' indicates the Dell EMC's Isilon Protection Source environment. 'kGPFS' indicates IBM's GPFS Protection Source environment. 'kKVM' indicates the KVM Protection Source environment. 'kAWS' indicates the AWS Protection Source environment. 'kExchange' indicates the Exchange Protection Source environment. 'kHyperVVSS' indicates the HyperV VSS Protection Source environment. 'kOracle' indicates the Oracle Protection Source environment. 'kGCP' indicates the Google Cloud Platform Protection Source environment. 'kFlashBlade' indicates the Flash Blade Protection Source environment. 'kAWSNative' indicates the AWS Native Protection Source environment. 'kO365' indicates the Office 365 Protection Source environment. 'kO365Outlook' indicates Office 365 outlook Protection Source environment. 'kHyperFlex' indicates the Hyper Flex Protection Source environment. 'kGCPNative' indicates the GCP Native Protection Source environment. 'kAzureNative' indicates the Azure Native Protection Source environment. 'kKubernetes' indicates a Kubernetes Protection Source environment. 'kElastifile' indicates Elastifile Protection Source environment. 'kAD' indicates Active Directory Protection Source environment. 'kRDSSnapshotManager' indicates AWS RDS Protection Source environment. 'kCassandra' indicates Cassandra Protection Source environment. 'kMongoDB' indicates MongoDB Protection Source environment. 'kCouchbase' indicates Couchbase Protection Source environment. 'kHdfs' indicates Hdfs Protection Source environment. 'kHive' indicates Hive Protection Source environment. 'kHBase' indicates HBase Protection Source environment. 'kUDA' indicates Universal Data Adapter Protection Source environment. 'kSAPHANA' indicates SAP HANA protection source environment. 'kO365Teams' indicates the Office365 Teams Protection Source environment. 'kO365Group' indicates the Office365 Groups Protection Source environment. 'kO365Exchange' indicates the Office365 Mailbox Protection Source environment. 'kO365OneDrive' indicates the Office365 OneDrive Protection Source environment. 'kO365Sharepoint' indicates the Office365 SharePoint Protection Source environment. 'kO365PublicFolders' indicates the Office365 PublicFolders ProtectionSource environment. kHpe3Par, kIbmFlashSystem, kAzure, kNetapp, kAgent, kGenericNas, kAcropolis, kPhysicalFiles, kIsilon, kGPFS, kKVM, kAWS, kExchange, kHyperVVSS, kOracle, kGCP, kFlashBlade, kAWSNative, kO365, kO365Outlook, kHyperFlex, kGCPNative, kAzureNative, kKubernetes, kElastifile, kAD, kRDSSnapshotManager, kCassandra, kMongoDB, kCouchbase, kHdfs, kHive, kHBase, kUDA, kSAPHANA, kO365Teams, kO365Group, kO365Exchange, kO365OneDrive, kO365Sharepoint, kO365PublicFolders.
	  * Constraints: Allowable values are: `kVMware`, `kHyperV`, `kSQL`, `kView`, `kPuppeteer`, `kPhysical`, `kPure`, `kNimble`.
	* `kubernetes_distribution_stats` - (List) Specifies the breakdown of the kubernetes clusters by distribution type.
	Nested schema for **kubernetes_distribution_stats**:
		* `distribution` - (String) Specifies the type of Kuberentes distribution Determines the K8s distribution. kIKS, kROKS.
		  * Constraints: Allowable values are: `kMainline`, `kOpenshift`, `kRancher`, `kEKS`, `kGKE`, `kAKS`, `kVMwareTanzu`.
		* `protected_count` - (Integer) Specifies the number of objects that are protected for that distribution.
		* `protected_size` - (Integer) Specifies the total size of objects that are protected for that distribution.
		* `total_registered_clusters` - (Integer) Specifies the number of registered clusters for that distribution.
		* `unprotected_count` - (Integer) Specifies the number of objects that are not protected for that distribution.
		* `unprotected_size` - (Integer) Specifies the total size of objects that are not protected for that distribution.
	* `protected_count` - (Integer) Specifies the number of objects that are protected under the given entity.
	* `protected_size` - (Integer) Specifies the total size of the protected objects under the given entity.
	* `unprotected_count` - (Integer) Specifies the number of objects that are not protected under the given entity.
	* `unprotected_size` - (Integer) Specifies the total size of the unprotected objects under the given entity.

