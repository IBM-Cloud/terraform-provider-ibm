---
layout: "ibm"
page_title: "IBM : ibm_database_tasks"
description: |-
  Get information about database_tasks
subcategory: "Cloud Databases"
---

# ibm_database_tasks

Provides a read-only data source for database_tasks. Supports both Classic and Gen2 database instances. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

### Classic

```hcl
data "ibm_database_tasks" "database_tasks" {
  deployment_id = data.ibm_database.database.id
}
```

### Gen2

For Gen2 instances, the `deployment_id` is a Gen2 database instance CRN. The data source returns a single task representing the current instance state.

```hcl
data "ibm_database_tasks" "database_tasks" {
  deployment_id = "crn:v1:bluemix:public:databases-for-postgresql:<region>:a/<account_id>:<instance_id>::"
}
```

## Argument Reference

Review the argument reference that you can specify for your data source.

* `deployment_id` - (Required, Forces new resource, String) Deployment ID.

  **Classic:** The database instance CRN, for example:
  ```
  crn:v1:bluemix:public:databases-for-postgresql:us-east:a/<account_id>:<instance_id>::
  ```

  **Gen2:** The Gen2 database instance CRN, for example:
  ```
  crn:v1:bluemix:public:databases-for-postgresql:<region>:a/<account_id>:<instance_id>::
  ```
  The Gen2 deployment CRN can be retrieved from:
  - The `id` attribute of an `ibm_database` resource configured with a Gen2 plan.
  - The IBM Cloud UI under **Databases → your instance → Overview**.
  - The IBM Cloud CLI: `ibmcloud resource service-instance <instance_name> --output json | jq -r '.[0].crn'`.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

* `deployment_id` - The unique identifier of the database_tasks.
* `tasks` - (Optional, List) An array of tasks. Always contains exactly one entry for Gen2 instances.
Nested scheme for **tasks**:
	* `task_id` - (Optional, String) ID of the task. Empty string `""` for Gen2 instances.
	* `deployment_id` - (Optional, String) ID of the deployment the task is being performed on.
	* `description` - (Optional, String) Human-readable description of the task.
	* `status` - (Optional, String) The status of the task.
	  * Constraints: Allowable values are: `running`, `completed`, `failed`, `queued`.
	* `progress_percent` - (Optional, Integer) Indicator as percentage of progress of the task.
	* `created_at` - (Optional, String) Date and time when the task was created.

## Gen2 Behaviour

For Gen2 database instances, the data source derives task information from the current instance state via the Resource Controller API rather than the ICD tasks API. `tasks` always contains exactly one entry. The instance state is mapped to task attributes as follows:

| Instance State | `status` | `progress_percent` |
|----------------|----------|--------------------|
| `active` | `completed` | `100` |
| `provisioning`, `preparing` | `running` | `50` |
| `inactive` | `queued` | `0` |
| `failed`, `removed` | `failed` | `0` |

