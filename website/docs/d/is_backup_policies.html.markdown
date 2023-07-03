---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : is_backup_policies"
description: |-
  Get information about backup policies.
---

# ibm_is_backup_policies

Provides a read-only data source for BackupPolicyCollection. For more information, about backup policy in your IBM Cloud VPC, see [Backup policy](https://cloud.ibm.com/docs/vpc?topic=vpc-backup-view-policies).

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
data "ibm_is_backup_policies" "example" {
}
```
 
## Argument Reference

Review the argument reference that you can specify for your data source.

- `name` - (Optional, String) Filters the collection to resources with the exact specified name.
- `resource_group` - (Optional, String) Filters the collection to resources in the resource group with the specified identifier.
- `tag` - (Optional, String) Filters the collection to resources with the exact tag value.

## Attribute Reference
In addition to all argument reference list, you can access the following attribute references after your data source is created.

- `id` - The unique identifier of the BackupPolicyCollection.
- `backup_policies` - (List) Collection of backup policies. 
  
  Nested `backup_policies` blocks have the following structure:
  - `created_at` -  (String) The date and time that the backup policy was created.
  - `crn` - (String) The CRN for this backup policy.
  - `href` - (String) The URL for this backup policy.
  - `id` - (String) The unique identifier for this backup policy.
  - `last_job_completed_at` - (String) he date and time that the most recent job for this backup policy completed.
  - `lifecycle_state` - (String) The lifecycle state of the backup policy.
  - `match_resource_types` - (List) A resource type this backup policy applies to. Resources that have both a matching type and a matching user tag will be subject to the backup policy.
  - `match_user_tags` - (List) The user tags this backup policy applies to. Resources that have both a matching user tag and a matching type will be subject to the backup policy.
  - `name` - (String) The unique user-defined name for this backup policy.
  - `plans` - (List) The plans for the backup policy. 
    
      Nested `plans` blocks have the following structure:
      - `deleted` - (List) If present, this property indicates the referenced resource has been deleted and provides some supplementary information.
        
          Nested `deleted` blocks have the following structure: 
          - `more_info` - (String) Link to documentation about deleted resources.
      - `href` - (String) The URL for this backup policy plan.
      - `id` - (String) The unique identifier for this backup policy plan.
      - `name` - (String) The unique user-defined name for this backup policy plan.
      - `resource_type` - (String) The type of resource referenced.
    - `resource_group` - (List) The resource group object, for this backup policy. 
    
      Nested `resource_group` blocks have the following structure:
        - `href` - (String) The URL for this resource group.
        - `id` - (String) The unique identifier for this resource group.
        - `name` - (String) The user-defined name for this resource group.


