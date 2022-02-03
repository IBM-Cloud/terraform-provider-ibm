`-
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : ibm_is_backup_policy_plan"
description: |-
  Manages BackupPolicyPlan.
`-

# ibm_is_backup_policy_plan

Provides a resource for BackupPolicyPlan. This allows BackupPolicyPlan to be created, updated and deleted.For more information, about backup policy plan in your IBM Cloud VPC, see [Backup policy plan](https://cloud.ibm.com/docs/vpc?topic=vpc-creating-backup-policy-plan).

## Example Usage

```hcl
resource "ibm_is_backup_policy_plan" "example" {
  backup_policy_id = "backup_policy_id"
  cron_spec = "-/5 1,2,3 - - -"
  name = "example-policy-plan"
}
```

## Argument Reference

Review the argument reference that you can specify for your resource.

- `active` - (Optional, Boolean) Indicates whether the plan is active.
- `attach_user_tags` - (Optional, List) User tags to attach to each backup (snapshot) created by this plan. If unspecified, no user tags will be attached.
- `backup_policy_id` - (Required, Forces new resource, String) The backup policy identifier.
- `clone_policy` - (Optional, List) 
Nested scheme for `clone_policy`:
	- `max_snapshots` - (Optional, Integer) The maximum number of recent snapshots (per source) that will keep clones.
	- `zones` - (Required, List) The zone this backup policy plan will create snapshot clones in.
	Nested scheme for `zones`:
		- `href` - (Optional, String) The URL for this zone.
		- `name` - (Optional, String) The globally unique name for this zone.
- `copy_user_tags` - (Optional, Boolean) Indicates whether to copy the source's user tags to the created backups (snapshots). The default value is `true`.
- `cron_spec` - (Required, String) The cron specification for the backup schedule. The value must match regular expression `/^((((\\d+,)+\\d+|([\\d\\-]+(\/|-)\\d+)|\\d+|\\-) ?){5,7})$/`.
- `deletion_trigger` - (Optional, List) 
Nested scheme for `deletion_trigger`:
	- `delete_after` - (Optional, Integer) The maximum number of days to keep each backup after creation.
	- `delete_over_count` - (Optional, Integer) The maximum number of recent backups to keep. If unspecified, there will be no maximum.
- `name` - (Optional, String) The user-defined name for this backup policy plan. Names must be unique within the backup policy this plan resides in. If unspecified, the name will be a hyphenated list of randomly-selected words.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

- `id` - The unique identifier of the BackupPolicyPlan.
- `created_at` - (String) The date and time that the backup policy plan was created.
- `href` - (String) The URL for this backup policy plan.
- `lifecycle_state` - (String) The lifecycle state of this backup policy plan.
  - Constraints: Allowable values are: `deleting`, `failed`, `pending`, `stable`, `suspended`, `updating`, `waiting`.
- `resource_type` - (String) The resource type.
- `version` - Version of the BackupPolicyPlan.

## Import

You can import the `ibm_is_backup_policy_plan` resource by using `id`.
The `id` property can be formed from `backup_policy_id`, and `id` in the following format:

```
<backup_policy_id>/<id>
```
- `backup_policy_id`: A string. The backup policy identifier.
- `id`: A string. The backup policy plan identifier.

# Syntax
```
$ terraform import ibm_is_backup_policy_plan.is_backup_policy_plan <backup_policy_id>/<id>
```
