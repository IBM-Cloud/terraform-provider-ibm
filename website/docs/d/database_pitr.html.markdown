---
layout: "ibm"
page_title: "IBM : ibm_database_pitr"
description: |-
  Get information about database_pitr
subcategory: "The IBM Cloud Databases API"
---

# ibm_database_pitr

Provides a read-only data source for database_pitr. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_database_pitr" "database_pitr" {
	deployment_id = "id"
}
```

## Argument Reference

Review the argument reference that you can specify for your data source.

* `id` - (Required, Forces new resource, String) Deployment ID.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

* `id` - The unique identifier of the database_pitr.
* `earliest_point_in_time_recovery_time` - (Optional, String) 

