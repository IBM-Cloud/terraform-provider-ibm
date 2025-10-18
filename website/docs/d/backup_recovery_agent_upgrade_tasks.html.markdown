---
layout: "ibm"
page_title: "IBM : ibm_backup_recovery_agent_upgrade_tasks"
description: |-
  Get information about Agent upgrade task list
subcategory: "IBM Backup Recovery"
---

# ibm_backup_recovery_agent_upgrade_tasks

Provides a read-only data source to retrieve information about an Agent upgrade task list. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_backup_recovery_agent_upgrade_tasks" "backup_recovery_agent_upgrade_tasks" {
	x_ibm_tenant_id = ibm_backup_recovery_agent_upgrade_task.backup_recovery_agent_upgrade_task_instance.x_ibm_tenant_id
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `ids` - (Optional, List) Specifies IDs of tasks to be fetched.
* `endpoint_type` - (Optional, String) Backup Recovery Endpoint type. By default set to "public".
* `instance_id` - (Optional, String) Backup Recovery instance ID. If provided here along with region, the provider constructs the endpoint URL using them, which overrides any value set through environment variables or the `endpoints.json` file.
* `region` - (Optional, String) Backup Recovery region. If provided here along with instance_id, the provider constructs the endpoint URL using them, which overrides any value set through environment variables or the `endpoints.json` file.  
* `x_ibm_tenant_id` - (Required, String) Specifies the key to be used to encrypt the source credential. If includeSourceCredentials is set to true this key must be specified.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the Agent upgrade task list.
* `tasks` - (List) Specifies the list of agent upgrade tasks.
Nested schema for **tasks**:
	* `agent_ids` - (List) Specifies the agents upgraded in the task.
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
	* `description` - (String) Specifies the description of the task.
	* `end_time_usecs` - (Integer) Specifies the time when the upgrade task completed execution as a Unix epoch Timestamp (in microseconds).
	* `error` - (List) Object that holds the error object.
	Nested schema for **error**:
		* `error_code` - (String) Specifies the error code.
		* `message` - (String) Specifies the error message.
		* `task_log_id` - (String) Specifies the TaskLogId of the failed task.
	* `id` - (Integer) Specifies the ID of the task.
	* `is_retryable` - (Boolean) Specifies if a task can be retried.
	* `name` - (String) Specifies the name of the task.
	* `retried_task_id` - (Integer) Specifies ID of a task which was retried if type is 'Retry'.
	* `schedule_end_time_usecs` - (Integer) Specifies the time before which the upgrade task should start execution as a Unix epoch Timestamp (in microseconds). If this is not specified the task will start anytime after scheduleTimeUsecs.
	* `schedule_time_usecs` - (Integer) Specifies the time when the task should start execution as a Unix epoch Timestamp (in microseconds). If no schedule is specified, the task will start immediately.
	* `start_time_usecs` - (Integer) Specifies the time, as a Unix epoch timestamp in microseconds, when the task started execution.
	* `status` - (String) Specifies the status of the task.<br> 'Scheduled' indicates that the upgrade task is yet to start.<br> 'Running' indicates that the upgrade task has started execution.<br> 'Succeeded' indicates that the upgrade task completed without an error.<br> 'Failed' indicates that upgrade has failed for all agents. 'PartiallyFailed' indicates that upgrade has failed for some agents.
	  * Constraints: Allowable values are: `Scheduled`, `Running`, `Succeeded`, `Failed`, `PartiallyFailed`, `Canceled`.
	* `type` - (String) Specifes the type of task.<br> 'Auto' indicates an auto agent upgrade task which is started after a cluster upgrade.<br> 'Manual' indicates a schedule based agent upgrade task.<br> 'Retry' indicates an agent upgrade task which was retried.
	  * Constraints: Allowable values are: `Auto`, `Manual`, `Retry`.

