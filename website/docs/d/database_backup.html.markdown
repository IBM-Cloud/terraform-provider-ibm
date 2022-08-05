---
layout: "ibm"
page_title: "IBM : ibm_database_backup"
description: |-
  Get information about Backup
subcategory: "Cloud Databases"
---

# ibm_database_backup

Provides a read-only data source for Backup. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_database_backup" "database_backup" {
	backup_id = "<backup_crn>"
}
```

## Argument Reference

Review the argument reference that you can specify for your data source.

* `backup_id` - (Required, Forces new resource, String) Backup ID.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

* `backup_id` - The unique identifier of the Backup.
* `created_at` - (Optional, String) Date and time when this backup was created.

* `deployment_id` - (Optional, String) ID of the deployment this backup relates to.

* `download_link` - (Optional, String) URI which is currently available for file downloading.

* `is_downloadable` - (Optional, Boolean) Is this backup available to download?.

* `is_restorable` - (Optional, Boolean) Can this backup be used to restore an instance?.

* `status` - (Optional, String) The status of this backup.
  * Constraints: Allowable values are: `running`, `completed`, `failed`.

* `type` - (Optional, String) The type of backup.
  * Constraints: Allowable values are: `scheduled`, `on_demand`.

