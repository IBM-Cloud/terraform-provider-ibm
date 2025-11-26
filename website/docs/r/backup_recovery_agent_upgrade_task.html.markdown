---
layout: "ibm"
page_title: "IBM : ibm_backup_recovery_agent_upgrade_task"
description: |-
  Manages Agent upgrade task state.
subcategory: "IBM Backup Recovery"
---

# ibm_backup_recovery_agent_upgrade_task

Create Agent upgrade task states with this resource.

**Note**
ibm_backup_recovery_agent_upgrade_task resource does not support update or delete operations due to the absence of corresponding API endpoints. As a result, Terraform cannot manage these operations for those resources. Users should be aware that removing these resources from the configuration (main.tf) will only remove them from the Terraform state and will not affect the actual resources in the backend. Similarly updating these resources will throw an error in the plan phase stating that the resource cannot be updated.

**Important:** When managing resources that lack complete CRUD operations, users should exercise caution and consider the limitations described above. Manual intervention may be required to manage these resources in the backend if updates or deletions are necessary.**


## Example Usage

```hcl
resource "ibm_backup_recovery_agent_upgrade_task" "backup_recovery_agent_upgrade_task_instance" {
  x_ibm_tenant_id = "x_ibm_tenant_id"
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `agent_ids` - (Optional, Forces new resource, List) Specifies the agents upgraded in the task.
* `description` - (Optional, Forces new resource, String) Specifies the description of the task.
* `name` - (Optional, Forces new resource, String) Specifies the name of the task.
* `endpoint_type` - (Optional, String) Backup Recovery Endpoint type. By default set to "public".
* `instance_id` - (Optional, String) Backup Recovery instance ID. If provided here along with region, the provider constructs the endpoint URL using them, which overrides any value set through environment variables or the `endpoints.json` file.
* `region` - (Optional, String) Backup Recovery region. If provided here along with instance_id, the provider constructs the endpoint URL using them, which overrides any value set through environment variables or the `endpoints.json` file.  
* `schedule_end_time_usecs` - (Optional, Forces new resource, Integer) Specifies the time before which the upgrade task should start execution as a Unix epoch Timestamp (in microseconds). If this is not specified the task will start anytime after scheduleTimeUsecs.
* `schedule_time_usecs` - (Optional, Forces new resource, Integer) Specifies the time when the task should start execution as a Unix epoch Timestamp (in microseconds). If no schedule is specified, the task will start immediately.
* `x_ibm_tenant_id` - (Required, Forces new resource, String) Specifies the key to be used to encrypt the source credential. If includeSourceCredentials is set to true this key must be specified.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the Agent upgrade task state.
* `agents` - (List) Specifies the upgrade information for each agent.
Nested schema for **agents**:
	* `id` - (Integer) Specifies the ID of the agent.
	* `info` - (List) Specifies the upgrade state of the agent.
	Nested schema for **info**:
		* `end_time_usecs` - (Integer) Specifies the time when the upgrade for an agent completed as a Unix epoch Timestamp (in microseconds).
		* `error` - (List) Object that holds the error object.
		Nested schema for **error**:
			* `error_code` - (String) Specifies the error code.
			* `message` - (String) Specifies the error message.
			* `task_log_id` - (String) Specifies the TaskLogId of the failed task.
		* `name` - (String) Specifies the name of the source where the agent is installed.
		* `previous_software_version` - (String) Specifies the software version of the agent before upgrade.
		* `start_time_usecs` - (Integer) Specifies the time when the upgrade for an agent started as a Unix epoch Timestamp (in microseconds).
		* `status` - (String) Specifies the upgrade status of the agent.<br> 'Scheduled' indicates that upgrade for the agent is yet to start.<br> 'Started' indicates that upgrade for the agent is started.<br> 'Succeeded' indicates that agent was upgraded successfully.<br> 'Failed' indicates that upgrade for the agent has failed.<br> 'Skipped' indicates that upgrade for the agent was skipped.
		  * Constraints: Allowable values are: `Scheduled`, `Started`, `Succeeded`, `Failed`, `Skipped`.
* `cluster_version` - (String) Specifies the version to which agents are upgraded.
* `end_time_usecs` - (Integer) Specifies the time when the upgrade task completed execution as a Unix epoch Timestamp (in microseconds).
* `error` - (List) Object that holds the error object.
Nested schema for **error**:
	* `error_code` - (String) Specifies the error code.
	* `message` - (String) Specifies the error message.
	* `task_log_id` - (String) Specifies the TaskLogId of the failed task.
* `is_retryable` - (Boolean) Specifies if a task can be retried.
* `retried_task_id` - (Integer) Specifies ID of a task which was retried if type is 'Retry'.
* `start_time_usecs` - (Integer) Specifies the time, as a Unix epoch timestamp in microseconds, when the task started execution.
* `status` - (String) Specifies the status of the task.<br> 'Scheduled' indicates that the upgrade task is yet to start.<br> 'Running' indicates that the upgrade task has started execution.<br> 'Succeeded' indicates that the upgrade task completed without an error.<br> 'Failed' indicates that upgrade has failed for all agents. 'PartiallyFailed' indicates that upgrade has failed for some agents.
  * Constraints: Allowable values are: `Scheduled`, `Running`, `Succeeded`, `Failed`, `PartiallyFailed`, `Canceled`.
* `type` - (String) Specifes the type of task.<br> 'Auto' indicates an auto agent upgrade task which is started after a cluster upgrade.<br> 'Manual' indicates a schedule based agent upgrade task.<br> 'Retry' indicates an agent upgrade task which was retried.
  * Constraints: Allowable values are: `Auto`, `Manual`, `Retry`.


### Import
Not Supported