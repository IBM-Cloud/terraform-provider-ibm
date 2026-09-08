---
layout: "ibm"
page_title: "IBM : ibm_brs_migration_workload_run"
description: |-
  Manages brs_migration_workload_run.
subcategory: "IBM Cloud Backup and Recovery Migration API"
---

# ibm_brs_migration_workload_run

Create, update, and delete brs_migration_workload_runs with this resource.

## Example Usage

```hcl
resource "ibm_brs_migration_workload_run" "brs_migration_workload_run_instance" {
  migration_id = "mgr-a1b2c3d4-e5f6-7890-abcd-ef12345678ab"
  workload_id = "wl-a1b2c3d4-e5f6-7890-abcd-ef1234567890"
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `migration_id` - (Required, Forces new resource, String) The migration project ID (mgr-{uuid4} format).
  * Constraints: Length must be `40` characters. The value must match regular expression `/^mgr-[0-9a-f-]{36}$/`.
* `workload_id` - (Required, Forces new resource, String) The migration service workload ID (wl-{uuid4} format).
  * Constraints: Length must be `39` characters. The value must match regular expression `/^wl-[0-9a-f-]{36}$/`.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the brs_migration_workload_run.
* `completed_at` - (String) Time the run completed. Null if still in progress.
* `duration_seconds` - (Integer) Wall-clock duration of the run in seconds.
  * Constraints: The maximum value is `2147483647`. The minimum value is `0`.
* `message` - (String) Human-readable status message or error detail.
  * Constraints: The maximum length is `1024` characters. The minimum length is `1` character. The value must match regular expression `/^.+$/`.
* `operation_type` - (String) Whether this run is a backup or a restore operation.
  * Constraints: Allowable values are: `backup`, `restore`.
* `payload_results` - (List) Per-payload breakdown of the run.
  * Constraints: The maximum length is `1000` items. The minimum length is `0` items.
Nested schema for **payload_results**:
	* `message` - (String) Error or warning detail specific to this payload.
	  * Constraints: The maximum length is `1024` characters. The minimum length is `1` character. The value must match regular expression `/^.+$/`.
	* `payload_id` - (String) ID of the workload payload this result belongs to.
	  * Constraints: The maximum length is `1024` characters. The minimum length is `1` character. The value must match regular expression `/^.+$/`.
	* `stats` - (List) Data-transfer statistics for a workload run or payload.
	Nested schema for **stats**:
		* `bytes_read` - (Integer) Number of bytes read from the source.
		  * Constraints: The maximum value is `9007199254740991`. The minimum value is `0`.
		* `bytes_transferred` - (Integer) Number of bytes successfully transferred.
		  * Constraints: The maximum value is `9007199254740991`. The minimum value is `0`.
		* `logical_size_bytes` - (Integer) Total logical size of all data processed, in bytes.
		  * Constraints: The maximum value is `9007199254740991`. The minimum value is `0`.
		* `total_file_count` - (Integer) Total number of files or objects processed.
		  * Constraints: The maximum value is `9007199254740991`. The minimum value is `0`.
		* `transferred_file_count` - (Integer) Number of files or objects successfully transferred.
		  * Constraints: The maximum value is `9007199254740991`. The minimum value is `0`.
	* `status` - (String) Status of this individual payload transfer.
	  * Constraints: Allowable values are: `accepted`, `running`, `canceling`, `canceled`, `succeeded`, `failed`.
* `run_id` - (String) Unique run ID (run-{uuid4} format).
  * Constraints: The maximum length is `100` characters. The minimum length is `1` character. The value must match regular expression `/^[0-9a-zA-Z-]+$/`.
* `run_type` - (String) Whether this run was triggered on-demand or by the schedule.
  * Constraints: Allowable values are: `scheduled`, `on_demand`.
* `started_at` - (String) Time the run started (ISO 8601 UTC).
* `stats` - (List) Data-transfer statistics for a workload run or payload.
Nested schema for **stats**:
	* `bytes_read` - (Integer) Number of bytes read from the source.
	  * Constraints: The maximum value is `9007199254740991`. The minimum value is `0`.
	* `bytes_transferred` - (Integer) Number of bytes successfully transferred.
	  * Constraints: The maximum value is `9007199254740991`. The minimum value is `0`.
	* `logical_size_bytes` - (Integer) Total logical size of all data processed, in bytes.
	  * Constraints: The maximum value is `9007199254740991`. The minimum value is `0`.
	* `total_file_count` - (Integer) Total number of files or objects processed.
	  * Constraints: The maximum value is `9007199254740991`. The minimum value is `0`.
	* `transferred_file_count` - (Integer) Number of files or objects successfully transferred.
	  * Constraints: The maximum value is `9007199254740991`. The minimum value is `0`.
* `status` - (String) Current execution status of the run.
  * Constraints: Allowable values are: `accepted`, `running`, `canceling`, `canceled`, `succeeded`, `failed`.


## Import

You can import the `ibm_brs_migration_workload_run` resource by using `id`.
The `id` property can be formed from `migration_id`, `workload_id`, and `run_id` in the following format:

<pre>
&lt;migration_id&gt;/&lt;workload_id&gt;/&lt;run_id&gt;
</pre>
* `migration_id`: A string in the format `mgr-a1b2c3d4-e5f6-7890-abcd-ef12345678ab`. The migration project ID (mgr-{uuid4} format).
* `workload_id`: A string in the format `wl-a1b2c3d4-e5f6-7890-abcd-ef1234567890`. The migration service workload ID (wl-{uuid4} format).
* `run_id`: A string in the format `run-e5f6a7b8-c9d0-1234-ef01-234567890123`. Unique run ID (run-{uuid4} format).

# Syntax
<pre>
$ terraform import ibm_brs_migration_workload_run.brs_migration_workload_run &lt;migration_id&gt;/&lt;workload_id&gt;/&lt;run_id&gt;
</pre>
