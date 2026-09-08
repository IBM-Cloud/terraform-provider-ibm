---
layout: "ibm"
page_title: "IBM : ibm_brs_migration_workload_history"
description: |-
  Get information about brs_migration_workload_history
subcategory: "IBM Cloud Backup and Recovery Migration API"
---

# ibm_brs_migration_workload_history

Provides a read-only data source to retrieve information about a brs_migration_workload_history. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_brs_migration_workload_history" "brs_migration_workload_history" {
	migration_id = "mgr-a1b2c3d4-e5f6-7890-abcd-ef12345678ab"
	workload_id = "wl-a1b2c3d4-e5f6-7890-abcd-ef1234567890"
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `migration_id` - (Required, Forces new resource, String) The migration project ID (mgr-{uuid4} format).
  * Constraints: Length must be `40` characters. The value must match regular expression `/^mgr-[0-9a-f-]{36}$/`.
* `workload_id` - (Required, Forces new resource, String) The migration service workload ID (wl-{uuid4} format).
  * Constraints: Length must be `39` characters. The value must match regular expression `/^wl-[0-9a-f-]{36}$/`.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the brs_migration_workload_history.
* `history` - (List) Workload execution history entries on this page.
  * Constraints: The maximum length is `1000` items. The minimum length is `0` items.
Nested schema for **history**:
	* `completed_at` - (String) Timestamp when this run completed.
	* `id` - (String) Unique identifier for this history entry.
	  * Constraints: The maximum length is `1024` characters. The minimum length is `1` character. The value must match regular expression `/^.+$/`.
	* `message` - (String) Human-readable status or error message.
	  * Constraints: The maximum length is `1024` characters. The minimum length is `1` character. The value must match regular expression `/^.+$/`.
	* `started_at` - (String) Timestamp when this run started.
	* `state` - (String) Final execution state of this history entry.
	  * Constraints: Allowable values are: `scheduled`, `running`, `canceling`, `canceled`, `succeeded`, `failed`.

