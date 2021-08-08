---
layout: "ibm"
page_title: "IBM : cloudant_database"
description: |-
  Manages cloudant_database.
subcategory: "Cloudant"
---

# ibm\_cloudant_database

Provides a resource for cloudant_database. This allows cloudant_database to be created, updated and deleted.

## Example Usage

```hcl
resource "ibm_cloudant_database" "cloudant_database" {
  cloudant_guid = var.cloudant_guid
  db            = var.db_name
}  
```

## Argument Reference

The following arguments are supported:

* `cloudant_guid` - (Required, string) Path parameter to specify the cloudant instance GUID.
* `db` - (Required, Forces new resource, string) Path parameter to specify the database name.
* `partitioned` - (Optional, Forces new resource, bool) Query parameter to specify whether to enable database partitions when creating a database.
  * Constraints: The default value is `false`.
* `q` - (Optional, Forces new resource, int) The number of shards in the database. Each shard is a partition of the hash value range. Default is 8, unless overridden in the `cluster config`.
  * Constraints: The maximum value is `5120`. The minimum value is `1`.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The unique identifier of the cloudant_database.

## Import

You can import the `cloudant_database` resource by using `ID`.
The `ID` property can be formed from `cloudant_guid`, and `db` in the following format:

```
<cloudant_guid>/<db>
```
* `cloudant_guid`: A string. Path parameter to specify the cloudant instance GUID.
* `db`: A string. Path parameter to specify the database name.

```
$ terraform import ibm_cloudant_database.cloudant_database <cloudant_guid>/<db>
```
