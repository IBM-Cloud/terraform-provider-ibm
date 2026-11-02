---
layout: "ibm"
page_title: "IBM : ibm_brs_migration"
description: |-
  Manages brs_migration.
subcategory: "IBM Cloud Backup and Recovery Migration API"
---

# ibm_brs_migration

Create, update, and delete brs_migrations with this resource.

## Example Usage

```hcl
resource "ibm_brs_migration" "brs_migration_instance" {
  brs_crn = "crn:v1:bluemix:public:backup-recovery:us-south:a/123456:abcdef01-2345-6789-abcd-ef0123456789::"
  description = "Migrate production Classic workloads to VPC"
  name = "prod-classic-to-vpc"
}
```

## Argument Reference

You can specify the following arguments for this resource.

* `brs_crn` - (Required, String) CRN of the IBM Cloud Backup and Recovery instance backing this migration.
  * Constraints: The maximum length is `1024` characters. The minimum length is `1` character. The value must match regular expression `/^crn:.+$/`.
* `description` - (Optional, String) Optional human-readable description.
  * Constraints: The maximum length is `1024` characters. The minimum length is `1` character. The value must match regular expression `/^.+$/`.
* `name` - (Required, String) Human-readable name for this migration project.
  * Constraints: The maximum length is `1024` characters. The minimum length is `1` character. The value must match regular expression `/^.+$/`.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

* `id` - The unique identifier of the brs_migration.
* `created_at` - (String) Timestamp when this migration was created.
* `crn` - (String) Server-assigned CRN for this migration resource.
  * Constraints: The maximum length is `1024` characters. The minimum length is `1` character. The value must match regular expression `/^crn:.+$/`.
* `state` - (String) Current lifecycle state of the migration project.
  * Constraints: Allowable values are: `active`, `completed`, `deleted`.
* `updated_at` - (String) Timestamp of the last update to this migration.


## Import

You can import the `ibm_brs_migration` resource by using `id`. Migration project ID (mgr-{uuid4} format).

# Syntax
<pre>
$ terraform import ibm_brs_migration.brs_migration &lt;id&gt;
</pre>

# Example
```
$ terraform import ibm_brs_migration.brs_migration mgr-a1b2c3d4-e5f6-7890-abcd-ef12345678ab
```
