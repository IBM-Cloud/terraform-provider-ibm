---
layout: "ibm"
page_title: "IBM : ibm_backup_recovery_protection_sources"
description: |-
  Get information about Protection Sources Response
subcategory: "IBM Backup Recovery"
---

# ibm_backup_recovery_protection_sources

Provides a read-only data source to retrieve information about a Protection Sources Response. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_backup_recovery_protection_sources" "backup_recovery_protection_sources" {
	x_ibm_tenant_id = "x_ibm_tenant_id"
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `after_cursor_entity_id` - (Optional, Integer) Specifies the entity id starting from which the items are to be returned.
* `endpoint_type` - (Optional, String) Backup Recovery Endpoint type. By default set to "public".
* `instance_id` - (Optional, String) Backup Recovery instance ID. If provided here along with region, the provider constructs the endpoint URL using them, which overrides any value set through environment variables or the `endpoints.json` file.
* `region` - (Optional, String) Backup Recovery region. If provided here along with instance_id, the provider constructs the endpoint URL using them, which overrides any value set through environment variables or the `endpoints.json` file.  
* `all_under_hierarchy` - (Optional, Boolean) AllUnderHierarchy specifies if objects of all the tenants under the hierarchy of the logged in user's organization should be returned.
* `backup_recovery_protection_sources_id` - (Optional, Integer) Return the Object subtree for the passed in Protection Source id.
* `before_cursor_entity_id` - (Optional, Integer) Specifies the entity id upto which the items are to be returned.
* `encryption_key` - (Optional, String) Key to be used to encrypt the source credential. If include_source_credentials is set to true this key must be specified.
* `environment` - (Deprecated, Optional, String) This field is deprecated. Use environments instead.
* `environments` - (Optional, List) Return only Protection Sources that match the passed in environment type such as 'kVMware', 'kSQL', 'kView' 'kPhysical', 'kPuppeteer', 'kPure', 'kNetapp', 'kGenericNas', 'kHyperV', 'kAcropolis', or 'kAzure'. For example, set this parameter to 'kVMware' to only return the Sources (and their Object subtrees) found in the 'kVMware' (VMware vCenter Server) environment.
  * Constraints: Allowable list items are: `kVMware`, `kHyperV`, `kSQL`, `kView`, `kPuppeteer`, `kPhysical`, `kPure`, `kNimble`, `kAzure`, `kNetapp`, `kAgent`, `kGenericNas`,`kKubernetes`, `kAcropolis`, `kPhysicalFiles`, `kIsilon`, `kGPFS`, `kKVM`, `kAWS`, `kExchange`, `kHyperVVSS`, `kOracle`, `kGCP`, `kFlashBlade`, `kAWSNative`, `kO365`, `kO365Outlook`, `kHyperFlex`, `kGCPNative`, `kAzureNative`, `kKubernetes`, `kElastifile`, `kAD`, `kRDSSnapshotManager`, `kCassandra`, `kMongoDB`, `kCouchbase`, `kHdfs`, `kHBase`, `kUDA`, `KSfdc`, `kAwsS3`.
* `exclude_aws_types` - (Optional, List) Specifies the Object types to be filtered out for AWS that match the passed in types such as 'kEC2Instance', 'kRDSInstance', 'kAuroraCluster', 'kTag', 'kAuroraTag', 'kRDSTag', kS3Bucket, kS3Tag. For example, set this parameter to 'kEC2Instance' to exclude ec2 instance from being returned.
  * Constraints: Allowable list items are: `kEC2Instance`, `kRDSInstance`, `kAuroraCluster`, `kS3Bucket`, `kTag`, `kRDSTag`, `kAuroraTag`, `kS3Tag`.
* `exclude_kubernetes_types` - (Optional, List) Specifies the Object types to be filtered out for Kubernetes that match the passed in types such as 'kService'. For example, set this parameter to 'kService' to exclude services from being returned.
  * Constraints: Allowable list items are: `kService`.
* `exclude_office365_types` - (Optional, List) Specifies the Object types to be filtered out for Office 365 that match the passed in types such as 'kDomain', 'kOutlook', 'kMailbox', etc. For example, set this parameter to 'kMailbox' to exclude Mailbox Objects from being returned.
  * Constraints: Allowable list items are: `kDomain`, `kOutlook`, `kMailbox`, `kUsers`, `kUser`, `kGroups`, `kGroup`, `kSites`, `kSite`.
* `exclude_types` - (Optional, List) Filter out the Object types (and their subtrees) that match the passed in types such as 'kVCenter', 'kFolder', 'kDatacenter', 'kComputeResource', 'kResourcePool', 'kDatastore', 'kHostSystem', 'kVirtualMachine', etc. For example, set this parameter to 'kResourcePool' to exclude Resource Pool Objects from being returned.
  * Constraints: Allowable list items are: `kVCenter`, `kFolder`, `kDatacenter`, `kComputeResource`, `kClusterComputeResource`, `kResourcePool`, `kDatastore`, `kHostSystem`, `kVirtualMachine`, `kVirtualApp`, `kStandaloneHost`, `kStoragePod`, `kNetwork`, `kDistributedVirtualPortgroup`, `kTagCategory`, `kTag`.
* `get_teams_channels` - (Optional, Boolean) Filter policies by a list of policy ids.
* `has_valid_mailbox` - (Optional, Boolean) If set to true, users with valid mailbox will be returned.
* `has_valid_onedrive` - (Optional, Boolean) If set to true, users with valid onedrive will be returned.
* `include_datastores` - (Optional, Boolean) Set this parameter to true to also return kDatastore object types found in the Source in addition to their Object subtrees. By default, datastores are not returned.
* `include_entity_permission_info` - (Optional, Boolean) If specified, then a list of entites with permissions assigned to them are returned.
* `include_networks` - (Optional, Boolean) Set this parameter to true to also return kNetwork object types found in the Source in addition to their Object subtrees. By default, network objects are not returned.
* `include_object_protection_info` - (Optional, Boolean) If specified, the object protection of entities(if any) will be returned.
* `include_sfdc_fields` - (Optional, Boolean) Set this parameter to true to also return fields of the object found in the Source in addition to their Object subtrees. By default, Sfdc object fields are not returned.
* `include_source_credentials` - (Optional, Boolean) If specified, then crednetial for the registered sources will be included. Credential is first encrypted with internal key and then reencrypted with user supplied 'encryption_key'.
* `include_system_v_apps` - (Optional, Boolean) Set this parameter to true to also return system VApp object types found in the Source in addition to their Object subtrees. By default, VM folder objects are not returned.
* `include_vm_folders` - (Optional, Boolean) Set this parameter to true to also return kVMFolder object types found in the Source in addition to their Object subtrees. By default, VM folder objects are not returned.
* `is_security_group` - (Optional, Boolean) If set to true, Groups which are security enabled will be returned.
* `node_id` - (Optional, Integer) Specifies the entity id for the Node at any level within the Source entity hierarchy whose children are to be paginated.
* `num_levels` - (Optional, Float) Specifies the expected number of levels from the root node to be returned in the entity hierarchy response.
* `page_size` - (Optional, Integer) Specifies the maximum number of entities to be returned within the page.
* `prune_aggregation_info` - (Optional, Boolean) Specifies whether to prune the aggregation information about the number of entities protected/unprotected.
* `prune_non_critical_info` - (Optional, Boolean) Specifies whether to prune non critical info within entities. Incase of VMs, virtual disk information will be pruned. Incase of Office365, metadata about user entities will be pruned. This can be used to limit the size of the response by caller.
* `request_initiator_type` - (Optional, String) Specifies the type of the request. Possible values are UIUser and UIAuto, which means the request is triggered by user or is an auto refresh request. Services like magneto will use this to determine the priority of the requests, so that it can more intelligently handle overload situations by prioritizing higher priority requests.
* `sids` - (Optional, List) Filter the object subtree for the sids given in the list.
* `use_cached_data` - (Optional, Boolean) Specifies whether we can serve the GET request to the read replica cache. setting this to true ensures that the API request is served to the read replica. setting this to false will serve the request to the master.
* `x_ibm_tenant_id` - (Required, String) Specifies the key to be used to encrypt the source credential. If includeSourceCredentials is set to true this key must be specified.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the Protection Sources Response.
* `protection_sources` - (List) Specifies list of protection sources.
Nested schema for **protection_sources**:
	* `application_nodes` - (List) Specifies the child subtree used to store additional application-level Objects. Different environments use the subtree to store application-level information. For example for SQL Server, this subtree stores the SQL Server instances running on a VM.
	Nested schema for **application_nodes**:
		* `nodes` - (List) Specifies children of the current node in the Protection Sources hierarchy.
		Nested schema for **nodes**:
	* `entity_pagination_parameters` - (List) Specifies the cursor based pagination parameters for Protection Source and its children. Pagination is supported at a given level within the Protection Source Hierarchy with the help of before or after cursors. A Cursor will always refer to a specific source within the source dataset but will be invalidated if the item is removed.
	Nested schema for **entity_pagination_parameters**:
		* `after_cursor_entity_id` - (Integer) Specifies the entity id starting from which the items are to be returned.
		* `before_cursor_entity_id` - (Integer) Specifies the entity id upto which the items are to be returned.
		* `node_id` - (Integer) Specifies the entity id for the Node at any level within the Source entity hierarchy whose children are to be paginated.
		* `page_size` - (Integer) Specifies the maximum number of entities to be returned within the page.
	* `entity_permission_info` - (List) Specifies the permission information of entities.
	Nested schema for **entity_permission_info**:
		* `entity_id` - (Integer) Specifies the entity id.
		* `groups` - (List)
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
			* `name` - (Boolean) Specifies name of the tenant.
			* `tenant_id` - (Boolean) Specifies the unique id of the tenant.
		* `users` - (List)
		Nested schema for **users**:
			* `domain` - (String) Specifies domain name of the user.
			* `sid` - (String) Specifies unique Security ID (SID) of the user.
			* `tenant_id` - (String) Specifies the tenant to which the user belongs to.
			* `user_name` - (String) Specifies user name of the user.
	* `logical_size` - (Integer) Specifies the logical size of the data in bytes for the Object on this node. Presence of this field indicates this node is a leaf node.
	* `nodes` - (List) Specifies children of the current node in the Protection Sources hierarchy. When representing Objects in memory, the entire Object subtree hierarchy is represented. You can use this subtree to navigate down the Object hierarchy.
	Nested schema for **nodes**:
	* `object_protection_info` - (List) Specifies the Object Protection Info of the Protection Source.
	Nested schema for **object_protection_info**:
		* `auto_protect_parent_id` - (Integer) Specifies the auto protect parent id if this entity is protected based on auto protection. This is only specified for leaf entities.
		* `entity_id` - (Integer) Specifies the entity id.
		* `has_active_object_protection_spec` - (Integer) Specifies if the entity is under object protection.
	* `protected_sources_summary` - (List) Array of Protected Objects. Specifies aggregated information about all the child Objects of this node that are currently protected by a Protection Job. There is one entry for each environment that is being backed up. The aggregated information for the Object hierarchy's environment will be available at the 0th index of the vector.
	Nested schema for **protected_sources_summary**:
		* `environment` - (String) Specifies the environment such as 'kSQL' or 'kVMware', where the Protection Source exists. Supported environment types such as 'kView', 'kSQL', 'kVMware', etc. NOTE 'kPuppeteer' refers to Cohesity's Remote Adapter. 'kVMware' indicates the VMware Protection Source environment. 'kHyperV' indicates the HyperV Protection Source environment. 'kSQL' indicates the SQL Protection Source environment. 'kView' indicates the View Protection Source environment. 'kPuppeteer' indicates the Cohesity's Remote Adapter. 'kPhysical' indicates the physical Protection Source environment. 'kPure' indicates the Pure Storage Protection Source environment. 'kNimble' indicates the Nimble Storage Protection Source environment. 'kAzure' indicates the Microsoft's Azure Protection Source environment. 'kNetapp' indicates the Netapp Protection Source environment. 'kAgent' indicates the Agent Protection Source environment. 'kGenericNas' indicates the Generic Network Attached Storage Protection Source environment. 'kAcropolis' indicates the Acropolis Protection Source environment. 'kPhysicalFiles' indicates the Physical Files Protection Source environment. 'kIbmFlashSystem' indicates the IBM Flash System Protection Source environment. 'kIsilon' indicates the Dell EMC's Isilon Protection Source environment. 'kGPFS' indicates IBM's GPFS Protection Source environment. 'kKVM' indicates the KVM Protection Source environment. 'kAWS' indicates the AWS Protection Source environment. 'kExchange' indicates the Exchange Protection Source environment. 'kHyperVVSS' indicates the HyperV VSS Protection Source environment. 'kOracle' indicates the Oracle Protection Source environment. 'kGCP' indicates the Google Cloud Platform Protection Source environment. 'kFlashBlade' indicates the Flash Blade Protection Source environment. 'kAWSNative' indicates the AWS Native Protection Source environment. 'kO365' indicates the Office 365 Protection Source environment. 'kO365Outlook' indicates Office 365 outlook Protection Source environment. 'kHyperFlex' indicates the Hyper Flex Protection Source environment. 'kGCPNative' indicates the GCP Native Protection Source environment. 'kAzureNative' indicates the Azure Native Protection Source environment. 'kKubernetes' indicates a Kubernetes Protection Source environment. 'kElastifile' indicates Elastifile Protection Source environment. 'kAD' indicates Active Directory Protection Source environment. 'kRDSSnapshotManager' indicates AWS RDS Protection Source environment. 'kCassandra' indicates Cassandra Protection Source environment. 'kMongoDB' indicates MongoDB Protection Source environment. 'kCouchbase' indicates Couchbase Protection Source environment. 'kHdfs' indicates Hdfs Protection Source environment. 'kHive' indicates Hive Protection Source environment. 'kHBase' indicates HBase Protection Source environment. 'kUDA' indicates Universal Data Adapter Protection Source environment. 'kO365Teams' indicates the Office365 Teams Protection Source environment. 'kO365Group' indicates the Office365 Groups Protection Source environment. 'kO365Exchange' indicates the Office365 Mailbox Protection Source environment. 'kO365OneDrive' indicates the Office365 OneDrive Protection Source environment. 'kO365Sharepoint' indicates the Office365 SharePoint Protection Source environment. 'kO365PublicFolders' indicates the Office365 PublicFolders Protection Source environment.
		  * Constraints: Allowable values are: `kVMware`, `kHyperV`, `kSQL`, `kView`, `kPhysical`, `kPure`, `kNimble`, `kIbmFlashSystem`, `kAzure`, `kNetapp`, `kAgent`, `kGenericNas`, `kAcropolis`, `kPhysicalFiles`, `kIsilon`, `kGPFS`, `kKVM`, `kAWS`, `kExchange`, `kHyperVVSS`, `kOracle`, `kGCP`, `kFlashBlade`, `kAWSNative`, `kO365`, `kO365Outlook`, `kHyperFlex`, `kGCPNative`, `kAzureNative`, `kKubernetes`, `kElastifile`, `kAD`, `kRDSSnapshotManager`, `kCassandra`, `kMongoDB`, `kCouchbase`, `kHdfs`, `kHive`, `kHBase`, `kUDA`, `kO365Teams`, `kO365Group`, `kO365Exchange`, `kO365OneDrive`, `kO365Sharepoint`, `kO365PublicFolders`.
		* `leaves_count` - (Integer) Specifies the number of leaf nodes under the subtree of this node.
		* `total_logical_size` - (Integer) Specifies the total logical size of the data under the subtree of this node.
	* `protection_source` - (List) Specifies details about an Acropolis Protection Source when the environment is set to 'kAcropolis'.
	Nested schema for **protection_source**:
		* `connection_id` - (Integer) Specifies the connection id of the tenant.
		* `connector_group_id` - (Integer) Specifies the connector group id of the connector groups.
		* `custom_name` - (String) Specifies the user provided custom name of the Protection Source.
		* `environment` - (String) Specifies the environment (such as 'kVMware' or 'kSQL') where the Protection Source exists. Depending on the environment, one of the following Protection Sources are initialized.
		  * Constraints: Allowable values are: `kPhysical`, `kPhysicalFiles`, `kSQL`, `kAgent`.
		* `id` - (Integer) Specifies an id of the Protection Source.
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
						  * Constraints: Allowable values are: `kPhysical`, `kPhysicalFiles`, `kSQL`, `kAgent`.
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
			  * Constraints: Allowable values are: `kPhysical`, `kPhysicalFiles`, `kSQL`, `kAgent`.
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
	* `total_downtiered_size_in_bytes` - (Integer) Specifies the total bytes downtiered from the source so far.
	* `total_uptiered_size_in_bytes` - (Integer) Specifies the total bytes uptiered to the source so far.
	* `unprotected_sources_summary` - (List)
	Nested schema for **unprotected_sources_summary**:
		* `environment` - (String) Specifies the environment such as 'kSQL' or 'kVMware', where the Protection Source exists. Supported environment types such as 'kView', 'kSQL', 'kVMware', etc. NOTE 'kPuppeteer' refers to Cohesity's Remote Adapter. 'kVMware' indicates the VMware Protection Source environment. 'kHyperV' indicates the HyperV Protection Source environment. 'kSQL' indicates the SQL Protection Source environment. 'kView' indicates the View Protection Source environment. 'kPuppeteer' indicates the Cohesity's Remote Adapter. 'kPhysical' indicates the physical Protection Source environment. 'kPure' indicates the Pure Storage Protection Source environment. 'kNimble' indicates the Nimble Storage Protection Source environment. 'kAzure' indicates the Microsoft's Azure Protection Source environment. 'kNetapp' indicates the Netapp Protection Source environment. 'kAgent' indicates the Agent Protection Source environment. 'kGenericNas' indicates the Generic Network Attached Storage Protection Source environment. 'kAcropolis' indicates the Acropolis Protection Source environment. 'kPhysicalFiles' indicates the Physical Files Protection Source environment. 'kIbmFlashSystem' indicates the IBM Flash System Protection Source environment. 'kIsilon' indicates the Dell EMC's Isilon Protection Source environment. 'kGPFS' indicates IBM's GPFS Protection Source environment. 'kKVM' indicates the KVM Protection Source environment. 'kAWS' indicates the AWS Protection Source environment. 'kExchange' indicates the Exchange Protection Source environment. 'kHyperVVSS' indicates the HyperV VSS Protection Source environment. 'kOracle' indicates the Oracle Protection Source environment. 'kGCP' indicates the Google Cloud Platform Protection Source environment. 'kFlashBlade' indicates the Flash Blade Protection Source environment. 'kAWSNative' indicates the AWS Native Protection Source environment. 'kO365' indicates the Office 365 Protection Source environment. 'kO365Outlook' indicates Office 365 outlook Protection Source environment. 'kHyperFlex' indicates the Hyper Flex Protection Source environment. 'kGCPNative' indicates the GCP Native Protection Source environment. 'kAzureNative' indicates the Azure Native Protection Source environment. 'kKubernetes' indicates a Kubernetes Protection Source environment. 'kElastifile' indicates Elastifile Protection Source environment. 'kAD' indicates Active Directory Protection Source environment. 'kRDSSnapshotManager' indicates AWS RDS Protection Source environment. 'kCassandra' indicates Cassandra Protection Source environment. 'kMongoDB' indicates MongoDB Protection Source environment. 'kCouchbase' indicates Couchbase Protection Source environment. 'kHdfs' indicates Hdfs Protection Source environment. 'kHive' indicates Hive Protection Source environment. 'kHBase' indicates HBase Protection Source environment. 'kUDA' indicates Universal Data Adapter Protection Source environment. 'kO365Teams' indicates the Office365 Teams Protection Source environment. 'kO365Group' indicates the Office365 Groups Protection Source environment. 'kO365Exchange' indicates the Office365 Mailbox Protection Source environment. 'kO365OneDrive' indicates the Office365 OneDrive Protection Source environment. 'kO365Sharepoint' indicates the Office365 SharePoint Protection Source environment. 'kO365PublicFolders' indicates the Office365 PublicFolders Protection Source environment.
		  * Constraints: Allowable values are: `kVMware`, `kHyperV`, `kSQL`, `kView`, `kPhysical`, `kPure`, `kNimble`, `kIbmFlashSystem`, `kAzure`, `kNetapp`, `kAgent`, `kGenericNas`, `kAcropolis`, `kPhysicalFiles`, `kIsilon`, `kGPFS`, `kKVM`, `kAWS`, `kExchange`, `kHyperVVSS`, `kOracle`, `kGCP`, `kFlashBlade`, `kAWSNative`, `kO365`, `kO365Outlook`, `kHyperFlex`, `kGCPNative`, `kAzureNative`, `kKubernetes`, `kElastifile`, `kAD`, `kRDSSnapshotManager`, `kCassandra`, `kMongoDB`, `kCouchbase`, `kHdfs`, `kHive`, `kHBase`, `kUDA`, `kO365Teams`, `kO365Group`, `kO365Exchange`, `kO365OneDrive`, `kO365Sharepoint`, `kO365PublicFolders`.
		* `leaves_count` - (Integer) Specifies the number of leaf nodes under the subtree of this node.
		* `total_logical_size` - (Integer) Specifies the total logical size of the data under the subtree of this node.

