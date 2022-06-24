---
layout: "ibm"
page_title: "IBM : ibm_database_tasks"
description: |-
  Get information about database_tasks
subcategory: "Cloud Databases"
---

# ibm_database_tasks

Provides a read-only data source for database_tasks. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_database_tasks" "database_tasks" {
	deployment_id = data.ibm_database.database.id
}
```

## Argument Reference

Review the argument reference that you can specify for your data source.

* `deployment_id` - (Required, Forces new resource, String) Deployment ID.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

* `deployment_id` - The unique identifier of the database_tasks.
* `tasks` - (Optional, List) 
Nested scheme for **tasks**:
	* `created_at` - (Optional, String) Date and time when the task was created.
	* `deployment_id` - (Optional, String) ID of the deployment the task is being performed on.
	* `description` - (Optional, String) Human-readable description of the task.
	* `task_id` - (Optional, String) ID of the task.
	* `progress_percent` - (Optional, Integer) Indicator as percentage of progress of the task.
	* `status` - (Optional, String) The status of the task.
	  * Constraints: Allowable values are: `running`, `completed`, `failed`.

