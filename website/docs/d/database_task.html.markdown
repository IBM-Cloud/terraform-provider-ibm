---
layout: "ibm"
page_title: "IBM : ibm_database_task"
description: |-
  Get information about database_task
subcategory: "Cloud Databases"
---

# ibm_database_task

Provides a read-only data source for database_task. Supports both Classic and Gen2 database instances. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

### Classic

```hcl
data "ibm_database_task" "database_task" {
  task_id = "<task_crn>"
}
```

### Gen2

For Gen2 instances, the `task_id` is the database instance CRN (not a task-specific CRN). The data source returns the current instance state as a single task.

```hcl
data "ibm_database_task" "database_task" {
  task_id = "crn:v1:bluemix:public:databases-for-postgresql:<region>:a/<account_id>:<instance_id>::"
}
```

## Argument Reference

Review the argument reference that you can specify for your data source.

* `task_id` - (Required, Forces new resource, String) Task ID.

  **Classic:** The task CRN returned by the ICD API, identifiable by the `:task:` segment, for example:
  ```
  crn:v1:bluemix:public:databases-for-postgresql:us-east:a/<account_id>:<instance_id>:task:<task_uuid>
  ```

  **Gen2:** The database instance CRN (the same CRN used as `deployment_id`), for example:
  ```
  crn:v1:bluemix:public:databases-for-postgresql:<region>:a/<account_id>:<instance_id>::
  ```
  The Gen2 instance CRN can be retrieved from:
  - The `id` attribute of an `ibm_database` resource configured with a Gen2 plan.
  - The IBM Cloud UI under **Databases → your instance → Overview**.

  **Note:** For Gen2 instances, `task_id` is set to an empty string `""` in the returned attributes because Gen2 does not expose individual task IDs.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

* `task_id` - The unique identifier of the task. Empty string `""` for Gen2 instances.
* `created_at` - (String) Date and time when the task was created.

* `deployment_id` - (String) ID of the deployment the task is being performed on.

* `description` - (String) Human-readable description of the task.

* `progress_percent` - (Integer) Indicator as percentage of progress of the task.

* `status` - (String) The status of the task.
  * Constraints: Allowable values are: `running`, `completed`, `failed`, `queued`.

## Gen2 Behaviour

For Gen2 database instances, the data source derives task information from the current instance state via the Resource Controller API rather than the ICD tasks API. The instance state is mapped to task attributes as follows:

| Instance State | `status` | `progress_percent` |
|----------------|----------|--------------------|
| `active` | `completed` | `100` |
| `provisioning`, `preparing` | `running` | `50` |
| `inactive` | `queued` | `0` |
| `failed`, `removed` | `failed` | `0` |
