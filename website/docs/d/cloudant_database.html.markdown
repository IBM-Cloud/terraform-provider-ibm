---
layout: "ibm"
page_title: "IBM : cloudant_database"
description: |-
  Get information about cloudant_database
subcategory: "Cloudant Databases"
---

# ibm\_cloudant_database

Provides a read-only data source for cloudant_database. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_cloudant_database" "cloudant_database" {
	db            = var.db_name
	instance_crn  = ibm_cloudant.cloudant.crn
}
```

## Argument Reference

The following arguments are supported:

* `db` - (Required, string) Path parameter to specify the database name.
* `instance_crn` - (Required, string) Path parameter to specify the cloudant instance CRN.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `cluster` - Schema for database cluster information. Nested `cluster` blocks have the following structure:
	* `read_quorum` - Read quorum. The number of consistent copies of a document that need to be read before a successful reply.
	* `replicas` - Schema for the number of replicas of a database in a cluster.
	* `shards` - Schema for the number of shards in a database. Each shard is a partition of the hash value range.
	* `write_quorum` - Write quorum. The number of copies of a document that need to be written before a successful reply.

* `committed_update_seq` - An opaque string that describes the committed state of the database.

* `compact_running` - True if the database compaction routine is operating on this database.

* `compacted_seq` - An opaque string that describes the compaction state of the database.

* `db_name` - The name of the database.

* `disk_format_version` - The version of the physical format used for the data when it is stored on disk.

* `doc_count` - A count of the documents in the specified database.

* `doc_del_count` - Number of deleted documents.

* `engine` - The engine used for the database.

* `id` - The unique identifier of the cloudant_database.

* `props` - Schema for database properties. Nested `props` blocks have the following structure:
	* `partitioned` - The value is `true` for a partitioned database.

* `sizes` - Schema for size information of content. Nested `sizes` blocks have the following structure:
	* `active` - The active size of the content, in bytes.
	* `external` - The total uncompressed size of the content, in bytes.
	* `file` - The total size of the content as stored on disk, in bytes.

* `update_seq` - An opaque string that describes the state of the database. Do not rely on this string for counting the number of updates.

* `uuid` - The UUID of the database.

