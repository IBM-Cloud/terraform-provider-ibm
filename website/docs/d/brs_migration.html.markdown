---
layout: "ibm"
page_title: "IBM : ibm_brs_migration"
description: |-
  Get information about brs_migration
subcategory: "IBM Cloud Backup and Recovery Migration API"
---

# ibm_brs_migration

Provides a read-only data source to retrieve information about a brs_migration. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_brs_migration" "brs_migration" {
	migration_id = "mgr-a1b2c3d4-e5f6-7890-abcd-ef12345678ab"
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `migration_id` - (Required, Forces new resource, String) The migration project ID (mgr-{uuid4} format).
  * Constraints: Length must be `40` characters. The value must match regular expression `/^mgr-[0-9a-f-]{36}$/`.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the brs_migration.
* `brs_crn` - (String) CRN of the IBM Cloud Backup and Recovery instance backing this migration.
  * Constraints: The maximum length is `1024` characters. The minimum length is `1` character. The value must match regular expression `/^crn:.+$/`.
* `created_at` - (String) Timestamp when this migration was created.
* `crn` - (String) Server-assigned CRN for this migration resource.
  * Constraints: The maximum length is `1024` characters. The minimum length is `1` character. The value must match regular expression `/^crn:.+$/`.
* `description` - (String) Optional human-readable description.
  * Constraints: The maximum length is `1024` characters. The minimum length is `1` character. The value must match regular expression `/^.+$/`.
* `name` - (String) Human-readable name for this migration project.
  * Constraints: The maximum length is `1024` characters. The minimum length is `1` character. The value must match regular expression `/^.+$/`.
* `state` - (String) Current lifecycle state of the migration project.
  * Constraints: Allowable values are: `active`, `completed`, `deleted`.
* `updated_at` - (String) Timestamp of the last update to this migration.

