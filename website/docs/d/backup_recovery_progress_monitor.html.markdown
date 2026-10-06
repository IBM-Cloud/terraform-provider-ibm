---
layout: "ibm"
page_title: "IBM : ibm_backup_recovery_progress_monitor"
description: |-
  Get information about Result returned by Pulse's GetTasks API.
subcategory: "IBM Backup Recovery API"
---

# ibm_backup_recovery_progress_monitor

Provides a read-only data source to retrieve information about a Result returned by Pulse's GetTasks API.. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_backup_recovery_progress_monitor" "backup_recovery_progress_monitor" {
	x_ibm_tenant_id = "tenantId"
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `attribute_vec` - (Optional, List) If specified, tasks matching the current query are futher filtered by these KeyValuePairs. This gives client an ability to search by custom attributes that they specified during the task creation. Only the Tasks having 'all' of the specified key=value pairs will be returned.
Nested schema for **attribute_vec**:
	* `key` - (Optional, String) Specifies the name of the key.
	* `value` - (Optional, List) Specifies a value for the key.
	Nested schema for **value**:
		* `data` - (Optional, List) Specifies the fields to store data of a given type.Specify data in the appropriate field for the current data type.
		Nested schema for **data**:
			* `oneof_data` - (Optional, List) Types that are valid to be assigned to OneofData:Value_Data_Int64ValueValue_Data_DoubleValueValue_Data_StringValueValue_Data_BytesValue.
			Nested schema for **oneof_data**:
		* `type` - (Optional, Integer) Specifies the type of value. 0 specifies a data point of type Int64. 1 specifies a data point of type Double. 2 specifies a data point of type String. 3 specifies a data point of type Bytes.
* `end_time_secs` - (Optional, Integer) Tasks that ended before this time.
* `end_time_secs` - (Optional, Integer) Tasks that ended before this time.
* `exclude_sub_tasks` - (Optional, Boolean) Skip information about the sub tasks of the matching root and sub tasks. By default, the entire task tree will be returned for matching tasks.
* `exclude_sub_tasks` - (Optional, Boolean) Skip information about the sub tasks of the matching root and sub tasks. By default, the entire task tree will be returned for matching tasks.
* `fetch_logs_max_level` - (Optional, Integer) Number of levels till which we need to fetch the event logs for a pulse tree. Note that it is applicable only when include_event_logs is true.
* `include_event_logs` - (Optional, Boolean) If set, the event logs will be included in the response message. Otherwise they will be cleared out.
* `include_finished_tasks` - (Optional, Boolean) Returns finished tasks as well. By default, Pulse only returns active tasks.
* `include_finished_tasks` - (Optional, Boolean) Returns finished tasks as well. By default, Pulse only returns active tasks.
* `max_tasks` - (Optional, Integer) Only return at most these many matching tasks. This constraint is applied with each query's result group.
* `max_tasks` - (Optional, Integer) Only return at most these many matching tasks. This constraint is applied with each query's result group.
* `start_time_secs` - (Optional, Integer) Tasks that started after this time.
* `start_time_secs` - (Optional, Integer) Tasks that started after this time.
* `task_path_vec` - (Optional, List) The hierarchical paths to the names of the tasks being queried. The task path-name specified here can be a prefix. Clients can specify multiple paths/prefixes. Pulse will return one ResultGroup for each path query.Each path is treated separately by Pulse, so if there are duplicate paths, Pulse will return duplicate results.Both root tasks and sub tasks can be specified in @task_path_vec.
* `task_path_vec` - (Optional, List) The hierarchical paths to the names of the tasks being queried. The task path-name specified here can be a prefix. Clients can specify multiple paths/prefixes. Pulse will return one ResultGroup for each path query.Each path is treated separately by Pulse, so if there are duplicate paths, Pulse will return duplicate results.Both root tasks and sub tasks can be specified in @task_path_vec.
* `x_ibm_tenant_id` - (Required, String) Specifies the unique id of the tenant.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the Result returned by Pulse's GetTasks API..
* `error` - (List) Proto to describe the error returned by pulse.
Nested schema for **error**:
	* `error_msg` - (String) A string describing the errors encountered.
	* `type` - (Integer) The type of error encountered.
* `result_group_vec` - (List) 
Nested schema for **result_group_vec**:
	* `task_vec` - (List) All tasks that match the corresponding query.
	Nested schema for **task_vec**:
		* `progress` - (List) The progress on this task.
		Nested schema for **progress**:
			* `approx_percent_unknown_work` - (Float) If set this indicate the percentage of work which is not know at this time. This will be useful if client does not know total amount of work that has to done. But client know how much work it has completed and approximate how much more work need to be done. This is usually reported by the clients for leaf tasks. For non-leaf tasks, the progress may be dynamically inferred.(see ReportTaskProgressArg).
			* `attribute_vec` - (List) The latest attributes (if any) reported for this task.
			Nested schema for **attribute_vec**:
				* `key` - (String) key.
				* `value` - (String) value.
			* `end_time_secs` - (Integer) The time when the task finished.
			* `event_vec` - (List) The events (if any) reported for this task.
			Nested schema for **event_vec**:
				* `event_msg` - (String) Message associated with the event.
				* `owner_percent_finished` - (Float) How much the owning task completed when this event occurred.
				* `owner_remaining_work_count` - (Integer) How much work was remaining for the owning task when this event occurred.
				* `timestamp_secs` - (Integer) The timestamp at which the event occurred.
			* `expected_end_time_secs` - (Integer) The expected end time of this task (if it hasn't ended). This is extrapolated using the current progress, and any historic data about this task if it occurs periodically. TODO(gaurav): Deprecate this field once Iris has stopped using it.
			* `expected_time_remaining_secs` - (Integer) Expected time remaining for this task (if it hasn't ended).
			* `expected_total_work_count` - (Integer) The expected raw count of the total work remaining. This is the highest work count value reported by the client. This field can be set to let pulse compute percent_finished by looking at the currently reported remaining_work_count and the expected_total_work_count.
			* `last_update_time_secs` - (Integer) The timestamp at which task progress was last reported.
			* `percent_finished` - (Float) The reported progress on this task. This is usually reported by clients for leaf tasks. For non-leaf tasks, the progress may be dynamically inferred.(see ReportTaskProgressArg).
			* `start_time_secs` - (Integer) The time when the task was started.
			* `status` - (List) The status of the task.
			Nested schema for **status**:
				* `error_msg` - (String) The error message (if any).
				* `type` - (Integer) The return type.
		* `sub_task_vec` - (List) Information about all the sub tasks for this task.
		* `task_path` - (String) The hierarchical name of the task.
		* `weight` - (Integer) The weight of this task.

