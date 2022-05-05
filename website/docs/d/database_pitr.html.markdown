---
layout: "ibm"
page_title: "IBM : ibm_database_point_in_time_recovery"
description: |-
  Get information about database_pitr
subcategory: "Cloud Databases"
---

# ibm_database_point_in_time_recovery

Provides a read-only data source for database_pitr. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_database_point_in_time_recovery" "database_pitr" {
	deployment_id = data.ibm_database.database.id
}
```

## Argument Reference

Review the argument reference that you can specify for your data source.

* `deployment_id` - (Required, Forces new resource, String) Deployment ID.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

* `deployment_id` - The unique identifier of the database_pitr.
* `earliest_point_in_time_recovery_time` - (String) - The earliest point in time recovery

