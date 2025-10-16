---
layout: "ibm"
page_title: "IBM : ibm_backup_recovery_protection_group_run_request"
description: |-
  Manages backup_recovery_protection_group_run_request.
subcategory: "IBM Backup Recovery"
---

# ibm_backup_recovery_protection_group_run_request

Create backup_recovery_protection_group_run_requests with this resource.

**Note**
ibm_backup_recovery_protection_group_run_request resource does not support delete operation due to the absence of corresponding API endpoints. protection group run updates are managed through a separate resource `ibm_backup_recovery_update_protection_group_run_request`. 
Note that the `ibm_backup_recovery_protection_group_run_request` and `ibm_backup_recovery_update_protection_group_run_request` resources must be used in tandem to manage Protection Group Runs. There is no delete operation available for the protection group run resource. If  ibm_backup_recovery_update_protection_group_run_request or ibm_backup_recovery_protection_group_run_request resource is removed from the `main.tf` file, Terraform will remove it from the state file but not from the backend. The resource will continue to exist in the backend system.

**Important:** When managing resources that lack complete CRUD operations, users should exercise caution and consider the limitations described above. Manual intervention may be required to manage these resources in the backend if updates or deletions are necessary.**

## Example Usage

```hcl
resource "ibm_backup_recovery_protection_group_run_request" "backup_recovery_protection_group_run_request_instance" {
  group_id = "123:11:1"
  objects {
		id = 1
		app_ids = [ 1 ]
		physical_params {
			metadata_file_path = "metadata_file_path"
		}
  }
  run_type = "kRegular"
  targets_config {
		use_policy_defaults = true
		replications {
			id = 1
			retention {
				unit = "Days"
				duration = 1
				data_lock_config {
					mode = "Compliance"
					unit = "Days"
					duration = 1
					enable_worm_on_external_target = true
				}
			}
		}
		archivals {
			id = 1
			archival_target_type = "Tape"
			retention {
				unit = "Days"
				duration = 1
				data_lock_config {
					mode = "Compliance"
					unit = "Days"
					duration = 1
					enable_worm_on_external_target = true
				}
			}
			copy_only_fully_successful = true
		}
		cloud_replications {
			aws_target {
				name = "name"
				region = 1
				region_name = "region_name"
				source_id = 1
			}
			azure_target {
				name = "name"
				resource_group = 1
				resource_group_name = "resource_group_name"
				source_id = 1
				storage_account = 1
				storage_account_name = "storage_account_name"
				storage_container = 1
				storage_container_name = "storage_container_name"
				storage_resource_group = 1
				storage_resource_group_name = "storage_resource_group_name"
			}
			target_type = "AWS"
			retention {
				unit = "Days"
				duration = 1
				data_lock_config {
					mode = "Compliance"
					unit = "Days"
					duration = 1
					enable_worm_on_external_target = true
				}
			}
		}
  }
  x_ibm_tenant_id = "x_ibm_tenant_id"
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `objects` - (Optional, Forces new resource, List) Specifies the list of objects to be protected by this Protection Group run. These can be leaf objects or non-leaf objects in the protection hierarchy. This must be specified only if a subset of objects from the Protection Groups needs to be protected.
Nested schema for **objects**:
	* `app_ids` - (Optional, List) Specifies a list of ids of applications.
	* `id` - (Required, Integer) Specifies the id of object.
	* `physical_params` - (Optional, List) Specifies physical parameters for this run.
	Nested schema for **physical_params**:
		* `metadata_file_path` - (Optional, String) Specifies metadata file path during run-now requests for physical file based backups for some specific source. If specified, it will override any default metadata/directive file path set at the object level for the source. Also note that if the job default does not specify a metadata/directive file path for the source, then specifying this field for that source during run-now request will be rejected.
* `run_type` - (Required, Forces new resource, String) Type of protection run. 'kRegular' indicates an incremental (CBT) backup. Incremental backups utilizing CBT (if supported) are captured of the target protection objects. The first run of a kRegular schedule captures all the blocks. 'kFull' indicates a full (no CBT) backup. A complete backup (all blocks) of the target protection objects are always captured and Change Block Tracking (CBT) is not utilized. 'kLog' indicates a Database Log backup. Capture the database transaction logs to allow rolling back to a specific point in time. 'kSystem' indicates system volume backup. It produces an image for bare metal recovery.
  * Constraints: Allowable values are: `kRegular`, `kFull`, `kLog`, `kSystem`, `kHydrateCDP`, `kStorageArraySnapshot`.
* `targets_config` - (Optional, Forces new resource, List) Specifies the replication and archival targets.
Nested schema for **targets_config**:
	* `archivals` - (Optional, List) Specifies a list of archival targets configurations.
	Nested schema for **archivals**:
		* `archival_target_type` - (Required, String) Specifies the snapshot's archival target type from which recovery has been performed.
		  * Constraints: Allowable values are: `Tape`, `Cloud`, `Nas`.
		* `copy_only_fully_successful` - (Optional, Boolean) Specifies if Snapshots are copied from a fully successful Protection Group Run or a partially successful Protection Group Run. If false, Snapshots are copied the Protection Group Run, even if the Run was not fully successful i.e. Snapshots were not captured for all Objects in the Protection Group. If true, Snapshots are copied only when the run is fully successful.
		* `id` - (Required, Integer) Specifies the Archival target to copy the Snapshots to.
		* `retention` - (Optional, List) Specifies the retention of a backup.
		Nested schema for **retention**:
			* `data_lock_config` - (Optional, List) Specifies WORM retention type for the snapshots. When a WORM retention type is specified, the snapshots of the Protection Groups using this policy will be kept for the last N days as specified in the duration of the datalock. During that time, the snapshots cannot be deleted.
			Nested schema for **data_lock_config**:
				* `duration` - (Required, Integer) Specifies the duration for a dataLock. <br> Example. If duration is 7 and unit is Months, the dataLock is enabled for last 7 * 30 = 210 days of the backup.
				  * Constraints: The minimum value is `1`.
				* `enable_worm_on_external_target` - (Optional, Boolean) Specifies whether objects in the external target associated with this policy need to be made immutable.
				* `mode` - (Required, String) Specifies the type of WORM retention type. 'Compliance' implies WORM retention is set for compliance reason. 'Administrative' implies WORM retention is set for administrative purposes.
				  * Constraints: Allowable values are: `Compliance`, `Administrative`.
				* `unit` - (Required, String) Specificies the Retention Unit of a dataLock measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Months' then number of retention days will be 365 * 2 = 730 days.
				  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
			* `duration` - (Required, Integer) Specifies the duration for a backup retention. <br> Example. If duration is 7 and unit is Months, the retention of a backup is 7 * 30 = 210 days.
			  * Constraints: The minimum value is `1`.
			* `unit` - (Required, String) Specificies the Retention Unit of a backup measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Years' then number of retention days will be 365 * 2 = 730 days.
			  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
	* `cloud_replications` - (Optional, List) Specifies a list of cloud replication targets configurations.
	Nested schema for **cloud_replications**:
		* `aws_target` - (Optional, List) Specifies the configuration for adding AWS as repilcation target.
		Nested schema for **aws_target**:
			* `name` - (Computed, String) Specifies the name of the AWS Replication target.
			* `region` - (Required, Integer) Specifies id of the AWS region in which to replicate the Snapshot to. Applicable if replication target is AWS target.
			* `region_name` - (Computed, String) Specifies name of the AWS region in which to replicate the Snapshot to. Applicable if replication target is AWS target.
			* `source_id` - (Required, Integer) Specifies the source id of the AWS protection source registered on IBM cluster.
		* `azure_target` - (Optional, List) Specifies the configuration for adding Azure as replication target.
		Nested schema for **azure_target**:
			* `name` - (Computed, String) Specifies the name of the Azure Replication target.
			* `resource_group` - (Optional, Integer) Specifies id of the Azure resource group used to filter regions in UI.
			* `resource_group_name` - (Computed, String) Specifies name of the Azure resource group used to filter regions in UI.
			* `source_id` - (Required, Integer) Specifies the source id of the Azure protection source registered on IBM cluster.
			* `storage_account` - (Computed, Integer) Specifies id of the storage account of Azure replication target which will contain storage container.
			* `storage_account_name` - (Computed, String) Specifies name of the storage account of Azure replication target which will contain storage container.
			* `storage_container` - (Computed, Integer) Specifies id of the storage container of Azure Replication target.
			* `storage_container_name` - (Computed, String) Specifies name of the storage container of Azure Replication target.
			* `storage_resource_group` - (Computed, Integer) Specifies id of the storage resource group of Azure Replication target.
			* `storage_resource_group_name` - (Computed, String) Specifies name of the storage resource group of Azure Replication target.
		* `retention` - (Optional, List) Specifies the retention of a backup.
		Nested schema for **retention**:
			* `data_lock_config` - (Optional, List) Specifies WORM retention type for the snapshots. When a WORM retention type is specified, the snapshots of the Protection Groups using this policy will be kept for the last N days as specified in the duration of the datalock. During that time, the snapshots cannot be deleted.
			Nested schema for **data_lock_config**:
				* `duration` - (Required, Integer) Specifies the duration for a dataLock. <br> Example. If duration is 7 and unit is Months, the dataLock is enabled for last 7 * 30 = 210 days of the backup.
				  * Constraints: The minimum value is `1`.
				* `enable_worm_on_external_target` - (Optional, Boolean) Specifies whether objects in the external target associated with this policy need to be made immutable.
				* `mode` - (Required, String) Specifies the type of WORM retention type. 'Compliance' implies WORM retention is set for compliance reason. 'Administrative' implies WORM retention is set for administrative purposes.
				  * Constraints: Allowable values are: `Compliance`, `Administrative`.
				* `unit` - (Required, String) Specificies the Retention Unit of a dataLock measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Months' then number of retention days will be 365 * 2 = 730 days.
				  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
			* `duration` - (Required, Integer) Specifies the duration for a backup retention. <br> Example. If duration is 7 and unit is Months, the retention of a backup is 7 * 30 = 210 days.
			  * Constraints: The minimum value is `1`.
			* `unit` - (Required, String) Specificies the Retention Unit of a backup measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Years' then number of retention days will be 365 * 2 = 730 days.
			  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
		* `target_type` - (Required, String) Specifies the type of target to which replication need to be performed.
		  * Constraints: Allowable values are: `AWS`, `Azure`.
	* `replications` - (Optional, List) Specifies a list of replication targets configurations.
	Nested schema for **replications**:
		* `id` - (Required, Integer) Specifies id of Remote Cluster to copy the Snapshots to.
		* `retention` - (Optional, List) Specifies the retention of a backup.
		Nested schema for **retention**:
			* `data_lock_config` - (Optional, List) Specifies WORM retention type for the snapshots. When a WORM retention type is specified, the snapshots of the Protection Groups using this policy will be kept for the last N days as specified in the duration of the datalock. During that time, the snapshots cannot be deleted.
			Nested schema for **data_lock_config**:
				* `duration` - (Required, Integer) Specifies the duration for a dataLock. <br> Example. If duration is 7 and unit is Months, the dataLock is enabled for last 7 * 30 = 210 days of the backup.
				  * Constraints: The minimum value is `1`.
				* `enable_worm_on_external_target` - (Optional, Boolean) Specifies whether objects in the external target associated with this policy need to be made immutable.
				* `mode` - (Required, String) Specifies the type of WORM retention type. 'Compliance' implies WORM retention is set for compliance reason. 'Administrative' implies WORM retention is set for administrative purposes.
				  * Constraints: Allowable values are: `Compliance`, `Administrative`.
				* `unit` - (Required, String) Specificies the Retention Unit of a dataLock measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Months' then number of retention days will be 365 * 2 = 730 days.
				  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
			* `duration` - (Required, Integer) Specifies the duration for a backup retention. <br> Example. If duration is 7 and unit is Months, the retention of a backup is 7 * 30 = 210 days.
			  * Constraints: The minimum value is `1`.
			* `unit` - (Required, String) Specificies the Retention Unit of a backup measured in days, months or years. <br> If unit is 'Months', then number specified in duration is multiplied to 30. <br> Example: If duration is 4 and unit is 'Months' then number of retention days will be 30 * 4 = 120 days. <br> If unit is 'Years', then number specified in duration is multiplied to 365. <br> If duration is 2 and unit is 'Years' then number of retention days will be 365 * 2 = 730 days.
			  * Constraints: Allowable values are: `Days`, `Weeks`, `Months`, `Years`.
	* `use_policy_defaults` - (Optional, Boolean) Specifies whether to use default policy settings or not. If specified as true then 'replications' and 'arcihvals' should not be specified. In case of true value, replicatioan targets congfigured in the policy will be added internally.
	  * Constraints: The default value is `false`.
* `x_ibm_tenant_id` - (Required, Forces new resource, String) Specifies the key to be used to encrypt the source credential. If includeSourceCredentials is set to true this key must be specified.
* `group_id` - (Required, String) Specifies the protection group ID
* `endpoint_type` - (Optional, String) Backup Recovery Endpoint type. By default set to "public".
* `instance_id` - (Optional, String) Backup Recovery instance ID. If provided here along with region, the provider constructs the endpoint URL using them, which overrides any value set through environment variables or the `endpoints.json` file.
* `region` - (Optional, String) Backup Recovery region. If provided here along with instance_id, the provider constructs the endpoint URL using them, which overrides any value set through environment variables or the `endpoints.json` file.  

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the backup_recovery_protection_group_run_request.


## Import
Not Supported