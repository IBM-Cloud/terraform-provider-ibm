---
layout: "ibm"
page_title: "IBM : ibm_database_task"
description: |-
  Get information about database_task
subcategory: "The IBM Cloud Databases API"
---

# ibm_database_task

Provides a read-only data source for database_task. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_database_task" "database_task" {
	id = "id"
}
```

## Argument Reference

Review the argument reference that you can specify for your data source.

* `id` - (Required, Forces new resource, String) Task ID.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

* `id` - The unique identifier of the database_task.
* `created_at` - (Optional, String) Date and time when the task was created.

* `deployment_id` - (Optional, String) ID of the deployment the task is being performed on.

* `description` - (Optional, String) Human-readable description of the task.

* `progress_percent` - (Optional, Integer) Indicator as percentage of progress of the task.

* `status` - (Optional, String) The status of the task.
  * Constraints: Allowable values are: `running`, `completed`, `failed`.

