---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : ibm_is_backup_policy"
description: |-
  Manages BackupPolicy.
---

# ibm_is_backup_policy

Provides a resource for BackupPolicy. This allows BackupPolicy to be created, updated and deleted. For more information, about backup policy in your IBM Cloud VPC, see [Backup policy](https://cloud.ibm.com/docs/vpc?topic=vpc-creating-backup-policy).

## Example Usage

```hcl
resource "ibm_is_backup_policy" "example" {
  name = "example-backup-policy"
  match_user_tags = ["tags"]
}
```

## Argument Reference

Review the argument reference that you can specify for your resource.

- `match_resource_types` - (Optional, List) A resource type this backup policy applies to. Resources that have both a matching type and a matching user tag will be subject to the backup policy. The default value is `["volume"]`. Allowable list items are: `volume`.
- `match_user_tags` - (Required, List) The user tags this backup policy applies to. Resources that have both a matching user tag and a matching type will be subject to the backup policy.
- `name` - (Required, String) The user-defined name for this backup policy. Names must be unique within the region this backup policy resides in. 
- `plans` - (Optional, List) The prototype objects for backup plans to be created for this backup policy.
  Nested scheme for `plans`:
	- `active` - (Optional, Boolean) Indicates whether the plan is active.
	- `attach_user_tags` - (Optional, List) User tags to attach to each backup (snapshot) created by this plan. If unspecified, no user tags will be attached.
	- `clone_policy` - (Optional, List)
	Nested scheme for `clone_policy`:
		- `max_snapshots` - (Optional, Integer) The maximum number of recent snapshots (per source) that will keep clones.
		- `zones` - (Required, List) The zone this backup policy plan will create snapshot clones in.
		Nested scheme for `zones`:
			- `href` - (Optional, String) The URL for this zone.
			- `name` - (Required, String) The globally unique name for this zone.
	- `copy_user_tags` - (Optional, Boolean) Indicates whether to copy the source's user tags to the created backups (snapshots).
	- `cron_spec` - (Required, String) The cron specification for the backup schedule. The value must match regular expression `/^((((\\d+,)+\\d+|([\\d\\*]+(\/|-)\\d+)|\\d+|\\*) ?){5,7})$/`.
	- `deletion_trigger` - (Optional, List)
    Nested scheme for `deletion_trigger`:
      - `delete_after` - (Optional, Integer) The maximum number of days to keep each backup after creation.
      - `delete_over_count` - (Optional, Integer) The maximum number of recent backups to keep. If unspecified, there will be no maximum.
	- `name` - (Optional, String) The user-defined name for this backup policy plan. Names must be unique within the backup policy this plan resides in. If unspecified, the name will be a hyphenated list of randomly-selected words.
- `resource_group` - (Optional, List) The resource group to use. If unspecified, the account's [default resource group](https://cloud.ibm.com/apidocs/resource-manager#introduction) is used.
  Nested scheme for `resource_group`:
    - `id` - (Optional, String) The unique identifier for this resource group.
  
## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

- `id` - The unique identifier of the BackupPolicy.
- `created_at` - (String) The date and time that the backup policy was created.
- `crn` - (String) The CRN for this backup policy.
- `href` - (String) The URL for this backup policy.
- `last_job_completed_at` - (String) The date and time that the most recent job for this backup policy completed.
- `lifecycle_state` - (String) The lifecycle state of the backup policy.
- `resource_type` - (String) The resource type.
- `version` - Version of the BackupPolicy.

## Import

You can import the `ibm_is_backup_policy` resource by using `id`. The unique identifier for this backup policy.

# Syntax
```
$ terraform import ibm_is_backup_policy.is_backup_policy <id>
```

# Example
```
$ terraform import ibm_is_backup_policy.is_backup_policy 0fe9e5d8-0a4d-4818-96ec-e99708644a58
```
