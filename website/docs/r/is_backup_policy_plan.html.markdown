---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : ibm_is_backup_policy_plan"
description: |-
  Manages Backup Policy Plan.
---

# ibm_is_backup_policy_plan

Provides a resource for BackupPolicyPlan. This allows BackupPolicyPlan to be created, updated and deleted.For more information, about backup policy plan in your IBM Cloud VPC, see [Backup policy plan](https://cloud.ibm.com/docs/vpc?topic=vpc-backup-policy-create).

**Note:** 
VPC infrastructure services are a regional specific based endpoint, by default targets to `us-south`. Please make sure to target right region in the provider block as shown in the `provider.tf` file, if VPC service is created in region other than `us-south`.


**provider.tf**

```terraform
provider "ibm" {
  region = "eu-gb"
}
```

## Example Usage

```terraform
resource "ibm_is_backup_policy_plan" "example" {
  backup_policy_id = ibm_is_backup_policy.example.id
  cron_spec        = "0 12 * * *"
  name             = "example-backup-policy-plan"
}
```
## Example Usage (clones)

```terraform
resource "ibm_is_backup_policy_plan" "example" {
  backup_policy_id = ibm_is_backup_policy.example.id
  cron_spec        = "0 12 * * *"
  name             = "example-backup-policy-plan"
  clone_policy {
    zones 			    = ["us-south-1", "us-south-2"]
    max_snapshots 	= 3
  }
}
```

## Example Usage for Cross Region Copy
```terraform
resource "ibm_is_backup_policy_plan" "example" {
  backup_policy_id = ibm_is_backup_policy.example.id
  cron_spec        = "30 */2 * * 1-5"
  name             = "my-policy-plan"
  deletion_trigger {
    delete_after      = 20
    delete_over_count = 20
  }
  remote_region_policy {
    delete_over_count = 1
    encryption_key    = "crn:v1:bluemix:public:kms:us-south:a/dffefjgfeg88992eb3b752286b87:e398349-2ef0-42a6-8fd2-0348r34:key:i4ouo34u-c4b1-447f-9646-3498349kr"
    region            = "us-south"
  }
}
```

->**Note:**  Backup Policy Jobs are getting enhanced, will be available soon.

## Argument Reference

Review the argument reference that you can specify for your resource.

- `active` - (Optional, Boolean) Indicates whether the plan is active.
- `attach_user_tags` - (Optional, List) User tags to attach to each backup (snapshot) created by this plan. If unspecified, no user tags will be attached.
- `backup_policy_id` - (Required, Forces new resource, String) The backup policy identifier.
backup_policy_plan_id
- `copy_user_tags` - (Optional, Boolean) Indicates whether to copy the source's user tags to the created backups (snapshots). The default value is `true`.
- `cron_spec` - (Required, String) The cron specification for the backup schedule. The value must match regular expression `^((((\d+,)+\d+|([\d\*]+(\/|-)\d+)|\d+|\*) ?){5,7})$`.

	->**Note** The backup policy jobs (which create and delete backups for this plan) will not start until this time, and may start for up to 90 minutes after this time.All backup schedules for plans in the same policy must be at least an hour apart.
- `clone_policy` - (Optional, List)
  
  Nested scheme for `clone_policy`:
  - `max_snapshots` - (Optional, Integer) The maximum number of recent snapshots (per source) that will keep clones.
  - `zones` - (Optional, List) The zone list this backup policy plan will create snapshot clones in.

- `deletion_trigger` - (Optional, List)
  
  Nested scheme for `deletion_trigger`:
  - `delete_after` - (Optional, Integer) The maximum number of days to keep each backup after creation. Default value is 30.
  - `delete_over_count` - (Optional, String) The maximum number of recent backups to keep. If unspecified, there will be no maximum.
    
      ->**Note** Assign back to "null" to reset to no maximum.

- `name` - (Optional, String) The user-defined name for this backup policy plan. Names must be unique within the backup policy this plan resides in. If unspecified, the name will be a hyphenated list of randomly-selected words.

- `remote_region_policy` - (Optional, List) Backup policy plan cross region rule.

  Nested scheme for `remote_region_policy`:
	- `delete_over_count` - (Optional, Integer) The maximum number of recent remote copies to keep in this region. If no value is passed, then by default `delete_over_count` is 5. Range for `delete_over_count` is [1-100].
	- `encryption_key` - (Optional, String) The root key to use to rewrap the data encryption key for the snapshot.If unspecified, the source's `encryption_key` will be used.The specified key may be in a different account, subject to IAM policies. The CRN of the [Key Protect Root Key](https://cloud.ibm.com/docs/key-protect?topic=key-protect-getting-started-tutorial) or [Hyper Protect Crypto Services Root Key](https://cloud.ibm.com/docs/hs-crypto?topic=hs-crypto-get-started) for this resource.
	- `region` - (Required, String) Identifies a region by a unique property. The globally unique name for this region.


## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

- `id` - The unique identifier of the BackupPolicyPlan.
- `backup_policy_id` - (String) The backup policy identifier.
- `created_at` - (String) The date and time that the backup policy plan was created.
- `href` - (String) The URL for this backup policy plan.
- `lifecycle_state` - (String) The lifecycle state of this backup policy plan.
- `resource_type` - (String) The resource type.
- `version` - Version of the BackupPolicyPlan.

## Import

You can import the `ibm_is_backup_policy_plan` resource by using `id`.
The `id` property can be formed from `backup_policy_id`, and `id` in the following format:

```
<0fe9e5d8-0a4d-4818-96ec-e99708644a58>/<0fg9e5d8-0a4d-4818-96ec-e99708634a58>
```
- `backup_policy_id`: A string. The backup policy identifier.
- `id`: A string. The backup policy plan identifier.

# Syntax
```
$ terraform import ibm_is_backup_policy_plan.is_backup_policy_plan <0fe9e5d8-0a4d-4818-96ec-e99708644a58>/<0fg9e5d8-0a4d-4818-96ec-e99708634a58>

```
