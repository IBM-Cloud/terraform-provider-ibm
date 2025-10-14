---
layout: "ibm"
page_title: "IBM : ibm_backup_recovery_perform_action_on_protection_group_run_request"
description: |-
  Manages backup_recovery_perform_action_on_protection_group_run_request.
subcategory: "IBM Backup Recovery"
---

# ibm_backup_recovery_perform_action_on_protection_group_run_request

Create backup_recovery_perform_action_on_protection_group_run_requests with this resource.

**Note**
ibm_backup_recovery_perform_action_on_protection_group_run_request resource does not support update or delete operations due to the absence of corresponding API endpoints. As a result, Terraform cannot manage these operations for those resources. Users should be aware that removing these resources from the configuration (main.tf) will only remove them from the Terraform state and will not affect the actual resources in the backend. Similarly updating these resources will throw an error in the plan phase stating that the resource cannot be updated.

**Important:** When managing resources that lack complete CRUD operations, users should exercise caution and consider the limitations described above. Manual intervention may be required to manage these resources in the backend if updates or deletions are necessary.**

## Example Usage

```hcl
resource "ibm_backup_recovery_perform_action_on_protection_group_run_request" "backup_recovery_perform_action_on_protection_group_run_request_instance" {
  group_id = "123:11:1"
  action = "Pause"
  cancel_params {
		run_id = "run_id"
		local_task_id = "local_task_id"
		object_ids = [ 1 ]
		replication_task_id = [ "replicationTaskId" ]
		archival_task_id = [ "archivalTaskId" ]
		cloud_spin_task_id = [ "cloudSpinTaskId" ]
  }
  pause_params {
		run_id = "run_id"
  }
  resume_params {
		run_id = "run_id"
  }
  x_ibm_tenant_id = "x_ibm_tenant_id"
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `action` - (Required, Forces new resource, String) Specifies the type of the action which will be performed on protection runs.
  * Constraints: Allowable values are: `Pause`, `Resume`, `Cancel`.
* `cancel_params` - (Optional, Forces new resource, List) Specifies the cancel action params for a protection run.
Nested schema for **cancel_params**:
	* `archival_task_id` - (Optional, List) Specifies the task id of the archival run.
	  * Constraints: The list items must match regular expression `/^\\d+:\\d+:\\d+$/`.
	* `cloud_spin_task_id` - (Optional, List) Specifies the task id of the cloudSpin run.
	  * Constraints: The list items must match regular expression `/^\\d+:\\d+:\\d+$/`.
	* `local_task_id` - (Optional, String) Specifies the task id of the local run.
	  * Constraints: The value must match regular expression `/^\\d+:\\d+:\\d+$/`.
	* `object_ids` - (Optional, List) List of entity ids for which we need to cancel the backup tasks. If this is provided it will not cancel the complete run but will cancel only subset of backup tasks (if backup tasks are cancelled correspoding copy task will also get cancelled). If the backup tasks are completed successfully it will not cancel those backup tasks.
	* `replication_task_id` - (Optional, List) Specifies the task id of the replication run.
	  * Constraints: The list items must match regular expression `/^\\d+:\\d+:\\d+$/`.
	* `run_id` - (Required, String) Specifies a unique run id of the Protection Group run.
	  * Constraints: The value must match regular expression `/^\\d+:\\d+$/`.
* `pause_params` - (Optional, Forces new resource, List) Specifies the pause action params for a protection run.
Nested schema for **pause_params**:
	* `run_id` - (Required, String) Specifies a unique run id of the Protection Group run.
	  * Constraints: The value must match regular expression `/^\\d+:\\d+$/`.
* `resume_params` - (Optional, Forces new resource, List) Specifies the resume action params for a protection run.
Nested schema for **resume_params**:
	* `run_id` - (Required, String) Specifies a unique run id of the Protection Group run.
	  * Constraints: The value must match regular expression `/^\\d+:\\d+$/`.
* `x_ibm_tenant_id` - (Required, Forces new resource, String) Specifies the key to be used to encrypt the source credential. If includeSourceCredentials is set to true this key must be specified.
* `group_id` - (Required, String) Specifies the protection group ID
* `backup_recovery_endpoint` - (Optional, String) Backup Recovery Endpoint URL. If provided here, it overrides values configured via environment variable (IBMCLOUD_BACKUP_RECOVERY_ENDPOINT) or endpoints.json.   

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the backup_recovery_perform_action_on_protection_group_run_request.

### Import
Not Supported