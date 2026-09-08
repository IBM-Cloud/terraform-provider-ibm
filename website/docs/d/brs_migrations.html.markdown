---
layout: "ibm"
page_title: "IBM : ibm_brs_migrations"
description: |-
  Get information about brs_migrations
subcategory: "IBM Cloud Backup and Recovery Migration API"
---

# ibm_brs_migrations

Provides a read-only data source to retrieve information about brs_migrations. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_brs_migrations" "brs_migrations" {
}
```

## Argument Reference

You can specify the following arguments for this data source.

* `state` - (Optional, String) Filter by migration state.
  * Constraints: Allowable values are: `active`, `completed`, `deleted`.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

* `id` - The unique identifier of the brs_migrations.
* `migrations` - (List) The list of migration projects on this page.
  * Constraints: The maximum length is `100` items. The minimum length is `0` items.
Nested schema for **migrations**:
	* `brs_crn` - (String) CRN of the IBM Cloud Backup and Recovery instance backing this migration.
	  * Constraints: The maximum length is `1024` characters. The minimum length is `1` character. The value must match regular expression `/^crn:.+$/`.
	* `created_at` - (String) Timestamp when this migration was created.
	* `crn` - (String) Server-assigned CRN for this migration resource.
	  * Constraints: The maximum length is `1024` characters. The minimum length is `1` character. The value must match regular expression `/^crn:.+$/`.
	* `description` - (String) Optional human-readable description.
	  * Constraints: The maximum length is `1024` characters. The minimum length is `1` character. The value must match regular expression `/^.+$/`.
	* `id` - (String) Migration project ID (mgr-{uuid4} format).
	  * Constraints: Length must be `40` characters. The value must match regular expression `/^mgr-[0-9a-f-]{36}$/`.
	* `name` - (String) Human-readable name for this migration project.
	  * Constraints: The maximum length is `1024` characters. The minimum length is `1` character. The value must match regular expression `/^.+$/`.
	* `state` - (String) Current lifecycle state of the migration project.
	  * Constraints: Allowable values are: `active`, `completed`, `deleted`.
	* `updated_at` - (String) Timestamp of the last update to this migration.

