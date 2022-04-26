---
layout: "ibm"
page_title: "IBM : ibm_database_remotes"
description: |-
  Get information about database_remotes
subcategory: "Cloud Databases"
---

# ibm_database_remotes

Provides a read-only data source for database_remotes. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_database" "database" {
  name = "mydatabase"
  location = "us-east"
}

data "ibm_database_remotes" "database_remotes" {
	deployment_id = data.ibm_database.database.id
}
```

## Argument Reference

Review the argument reference that you can specify for your data source.

* `deployment_id` - (Required, Forces new resource, String) Deployment ID.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

* `id` - The unique identifier of the database_remotes.
* `leader` - (String) Leader ID, if applicable.

* `replicas` - (List) Replica IDs, if applicable.

