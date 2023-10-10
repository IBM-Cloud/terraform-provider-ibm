---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : ibm_is_backup_policy"
description: |-
  Manages BackupPolicy.
---

# ibm_is_backup_policy

Provides a resource for BackupPolicy. This allows BackupPolicy to be created, updated and deleted. For more information, about backup policy in your IBM Cloud VPC, see [Backup policy](https://cloud.ibm.com/docs/vpc?topic=vpc-backup-policy-create).

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
resource "ibm_is_backup_policy" "example" {
  match_user_tags = ["tag1"]
  name            = "example-backup-policy"
}
```

## Argument Reference

Review the argument reference that you can specify for your resource.

- `match_resource_types` - (Optional, List) A resource type this backup policy applies to. Resources that have both a matching type and a matching user tag will be subject to the backup policy. The default value is `["volume"]`.
  
~> **Note**
  `match_resource_types` is deprecated. Please use `match_resource_type` instead.
- `match_resource_type` - (Optional, String) The resource type this backup policy will apply to. Resources that have both a matching type and a matching user tag will be subject to the backup policy. The default value is `["volume"]`.
- `match_user_tags` - (Required, List) The user tags this backup policy applies to. Resources that have both a matching user tag and a matching type will be subject to the backup policy.
- `name` - (Required, String) The user-defined name for this backup policy. Names must be unique within the region this backup policy resides in. 
- `resource_group` - (Optional, List) The resource group id, to use. If unspecified, the account's [default resource group](https://cloud.ibm.com/apidocs/resource-manager#introduction) is used.

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
