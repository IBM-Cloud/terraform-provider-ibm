---
subcategory: "Db2 SaaS"
layout: "ibm"
page_title: "IBM : ibm_db2_backup"
description: |-
  Get Information about Backups of IBM Db2 instance.
---

# ibm_db2_backup

Provides a read-only data source to retrieve information about a Backups of an existing [IBM Db2 Instance](https://cloud.ibm.com/docs/Db2onCloud).

## Example Usage

```hcl
data "ibm_db2_backup" "db2_backup" {
	deployment_id = "crn%3Av1%3Astaging%3Apublic%3Adashdb-for-transactions%3Aus-south%3Aa%2Fe7e3e87b512f474381c0684a5ecbba03%3A39269573-e43f-43e8-8b93-09f44c2ff875%3A%3A"
}
```

## Argument Reference

Review the argument reference that you can specify for your data source.

* `deployment_id` - (Required, String) Encoded CRN of the instance this backup relates to.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

* `id` - The unique identifier of the db2_backup.
* `backups` - (List) 
Nested schema for **backups**:
	* `created_at` - (String) Timestamp of the backup created.
	* `duration` - (Integer) The duration of the backup operation in seconds.
	* `id` - (String) CRN of the db2 instance.
	* `size` - (Integer) Size of the backup or data set.
	* `status` - (String) Status of the backup.
	* `type` - (String) Defines the type of execution of backup.

