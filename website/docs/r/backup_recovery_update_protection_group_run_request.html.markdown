---
layout: "ibm"
page_title: "IBM : ibm_backup_recovery_update_protection_group_run_request"
description: |-
  Manages Update Protection Group Run Request Body..
subcategory: "IBM Backup Recovery"
---

# ibm_backup_recovery_update_protection_group_run_request

Create Update Protection Group Run Request Body.s with this resource.

**Note**
ibm_backup_recovery_update_protection_group_run_request resource does not support update or delete operations due to the absence of corresponding API endpoints. As a result, Terraform cannot manage these operations for those resources. Users should be aware that removing these resources from the configuration (main.tf) will only remove them from the Terraform state and will not affect the actual resources in the backend. Similarly updating these resources will throw an error in the plan phase stating that the resource cannot be updated.

**Important:** When managing resources that lack complete CRUD operations, users should exercise caution and consider the limitations described above. Manual intervention may be required to manage these resources in the backend if updates or deletions are necessary.**

## Example Usage

```hcl
resource "ibm_backup_recovery_update_protection_group_run_request" "backup_recovery_update_protection_group_run_request_instance" {
  group_id = "123:11:1"
  update_protection_group_run_params {
		run_id = "run_id"
		local_snapshot_config {
			enable_legal_hold = true
			delete_snapshot = true
			data_lock = "Compliance"
			days_to_keep = 1
		}
		replication_snapshot_config {
			new_snapshot_config {
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
			update_existing_snapshot_config {
				id = 1
				name = "name"
				enable_legal_hold = true
				delete_snapshot = true
				resync = true
				data_lock = "Compliance"
				days_to_keep = 1
			}
		}
		archival_snapshot_config {
			new_snapshot_config {
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
			update_existing_snapshot_config {
				id = 1
				name = "name"
				archival_target_type = "Tape"
				enable_legal_hold = true
				delete_snapshot = true
				resync = true
				data_lock = "Compliance"
				days_to_keep = 1
			}
		}
  }
  x_ibm_tenant_id = "x_ibm_tenant_id"
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `update_protection_group_run_params` - (Required, Forces new resource, List) 
  * Constraints: The minimum length is `1` item.
Nested schema for **update_protection_group_run_params**:
	* `archival_snapshot_config` - (Optional, List) Specifies the params to perform actions on archival snapshots taken by a Protection Group Run.
	Nested schema for **archival_snapshot_config**:
		* `new_snapshot_config` - (Optional, List) Specifies the new configuration about adding Archival Snapshot to existing Protection Group Run.
		Nested schema for **new_snapshot_config**:
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
		* `update_existing_snapshot_config` - (Optional, List) Specifies the configuration about updating an existing Archival Snapshot Run.
		Nested schema for **update_existing_snapshot_config**:
			* `archival_target_type` - (Required, String) Specifies the snapshot's archival target type from which recovery has been performed.
			  * Constraints: Allowable values are: `Tape`, `Cloud`, `Nas`.
			* `data_lock` - (Optional, String) Specifies WORM retention type for the snapshots. When a WORM retention type is specified, the snapshots of the Protection Groups using this policy will be kept until the maximum of the snapshot retention time. During that time, the snapshots cannot be deleted. <br>'Compliance' implies WORM retention is set for compliance reason. <br>'Administrative' implies WORM retention is set for administrative purposes.
			  * Constraints: Allowable values are: `Compliance`, `Administrative`.
			* `days_to_keep` - (Optional, Integer) Specifies number of days to retain the snapshots. If positive, then this value is added to exisiting expiry time thereby increasing  the retention period of the snapshot. Conversly, if this value is negative, then value is subtracted to existing expiry time thereby decreasing the retention period of the snaphot. Here, by this operation if expiry time goes below current time then snapshot is immediately deleted.
			* `delete_snapshot` - (Optional, Boolean) Specifies whether to delete the snapshot. When this is set to true, all other params will be ignored.
			* `enable_legal_hold` - (Optional, Boolean) Specifies whether to retain the snapshot for legal purpose. If set to true, the snapshots cannot be deleted until the retention period. Note that using this option may cause the Cluster to run out of space. If set to false explicitly, the hold is removed, and the snapshots will expire as specified in the policy of the Protection Group. If this field is not specified, there is no change to the hold of the run. This field can be set only by a User having Data Security Role.
			* `id` - (Required, Integer) Specifies the id of the archival target.
			* `name` - (Optional, String) Specifies the name of the archival target.
			* `resync` - (Optional, Boolean) Specifies whether to retry the archival operation in case if earlier attempt failed. If not specified or set to false, archival is not retried.
	* `local_snapshot_config` - (Optional, List) Specifies the params to perform actions on local snapshot taken by a Protection Group Run.
	Nested schema for **local_snapshot_config**:
		* `data_lock` - (Optional, String) Specifies WORM retention type for the snapshots. When a WORM retention type is specified, the snapshots of the Protection Groups using this policy will be kept until the maximum of the snapshot retention time. During that time, the snapshots cannot be deleted. <br>'Compliance' implies WORM retention is set for compliance reason. <br>'Administrative' implies WORM retention is set for administrative purposes.
		  * Constraints: Allowable values are: `Compliance`, `Administrative`.
		* `days_to_keep` - (Optional, Integer) Specifies number of days to retain the snapshots. If positive, then this value is added to exisiting expiry time thereby increasing  the retention period of the snapshot. Conversly, if this value is negative, then value is subtracted to existing expiry time thereby decreasing the retention period of the snaphot. Here, by this operation if expiry time goes below current time then snapshot is immediately deleted.
		* `delete_snapshot` - (Optional, Boolean) Specifies whether to delete the snapshot. When this is set to true, all other params will be ignored.
		* `enable_legal_hold` - (Optional, Boolean) Specifies whether to retain the snapshot for legal purpose. If set to true, the snapshots cannot be deleted until the retention period. Note that using this option may cause the Cluster to run out of space. If set to false explicitly, the hold is removed, and the snapshots will expire as specified in the policy of the Protection Group. If this field is not specified, there is no change to the hold of the run. This field can be set only by a User having Data Security Role.
	* `replication_snapshot_config` - (Optional, List) Specifies the params to perform actions on replication snapshots taken by a Protection Group Run.
	Nested schema for **replication_snapshot_config**:
		* `new_snapshot_config` - (Optional, List) Specifies the new configuration about adding Replication Snapshot to existing Protection Group Run.
		Nested schema for **new_snapshot_config**:
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
		* `update_existing_snapshot_config` - (Optional, List) Specifies the configuration about updating an existing Replication Snapshot Run.
		Nested schema for **update_existing_snapshot_config**:
			* `data_lock` - (Optional, String) Specifies WORM retention type for the snapshots. When a WORM retention type is specified, the snapshots of the Protection Groups using this policy will be kept until the maximum of the snapshot retention time. During that time, the snapshots cannot be deleted. <br>'Compliance' implies WORM retention is set for compliance reason. <br>'Administrative' implies WORM retention is set for administrative purposes.
			  * Constraints: Allowable values are: `Compliance`, `Administrative`.
			* `days_to_keep` - (Optional, Integer) Specifies number of days to retain the snapshots. If positive, then this value is added to exisiting expiry time thereby increasing  the retention period of the snapshot. Conversly, if this value is negative, then value is subtracted to existing expiry time thereby decreasing the retention period of the snaphot. Here, by this operation if expiry time goes below current time then snapshot is immediately deleted.
			* `delete_snapshot` - (Optional, Boolean) Specifies whether to delete the snapshot. When this is set to true, all other params will be ignored.
			* `enable_legal_hold` - (Optional, Boolean) Specifies whether to retain the snapshot for legal purpose. If set to true, the snapshots cannot be deleted until the retention period. Note that using this option may cause the Cluster to run out of space. If set to false explicitly, the hold is removed, and the snapshots will expire as specified in the policy of the Protection Group. If this field is not specified, there is no change to the hold of the run. This field can be set only by a User having Data Security Role.
			* `id` - (Required, Integer) Specifies the cluster id of the replication cluster.
			* `name` - (Optional, String) Specifies the cluster name of the replication cluster.
			* `resync` - (Optional, Boolean) Specifies whether to retry the replication operation in case if earlier attempt failed. If not specified or set to false, replication is not retried.
	* `run_id` - (Required, String) Specifies a unique Protection Group Run id.
	  * Constraints: The value must match regular expression `/^\\d+:\\d+$/`.
* `x_ibm_tenant_id` - (Required, Forces new resource, String) Specifies the key to be used to encrypt the source credential. If includeSourceCredentials is set to true this key must be specified.
* `group_id` - (Required, String) Specifies the protection group ID
* `backup_recovery_endpoint` - (Optional, String) Backup Recovery Endpoint URL. If provided here, it overrides values configured via environment variable (IBMCLOUD_BACKUP_RECOVERY_ENDPOINT) or endpoints.json.   

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the Update Protection Group Run Request Body..


## Import
Not Supported
