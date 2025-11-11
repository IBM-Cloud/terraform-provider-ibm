---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : ibm_is_snapshot_consistency_group"
description: |-
  Manages SnapshotConsistencyGroup.
---

# ibm_is_snapshot_consistency_group

Create, update, and delete SnapshotConsistencyGroups with this resource. For more information, about snapshot consistency group, see [creating snapshot consistency group](https://cloud.ibm.com/docs/vpc?topic=vpc-snapshots-vpc-create).

**Note**
VPC infrastructure services are a regional specific based endpoint, by default targets to `us-south`. Please make sure to target right region in the provider block as shown in the `provider.tf` file, if VPC service is created in region other than `us-south`.

**provider.tf**

```terraform
provider "ibm" {
  region = "eu-gb"
}
```

## Example Usage

```terraform
resource "ibm_is_snapshot_consistency_group" "is_snapshot_consistency_group" {
  delete_snapshots_on_delete = true
  snapshots {
    name          = "example-snapshot"
    source_volume = ibm_is_instance.testacc_instance.volume_attachments[0].volume_id
  }
  name = "example-snapshot-consistency-group"
}
```

## Argument Reference

You can specify the following arguments for this resource.

- `access_tags`  - (Optional, List of Strings) A list of access management tags to attach to the bare metal server.

  ~> **Note:** 
  **&#x2022;** You can attach only those access tags that already exists.</br>
  **&#x2022;** For more information, about creating access tags, see [working with tags](https://cloud.ibm.com/docs/account?topic=account-tag&interface=ui#create-access-console).</br>
  **&#x2022;** You must have the access listed in the [Granting users access to tag resources](https://cloud.ibm.com/docs/account?topic=account-access) for `access_tags`</br>
  **&#x2022;** `access_tags` must be in the format `key:value`.
- `delete_snapshots_on_delete` - (Optional, Boolean) Indicates whether deleting the snapshot consistency group will also delete the snapshots in the group.
- `name` - (Optional, String) The name for this snapshot consistency group. The name is unique across all snapshot consistency groups in the region.
- `resource_group` - (Optional, Forces new resource, String) The resource group ID where the snapshot consistency group is to be created.
- `snapshots` - (Required, List) The member snapshots that are data-consistent with respect to captured time. (may be[deleted](https://cloud.ibm.com/apidocs/vpc#deleted-resources)).
    
    Nested schema for `snapshots`:
    - `name` - (Optional, String) The name for this snapshot.
    - `source_volume` - (Required, Forces new resource, String)  The unique identifier for the volume for which snapshot is to be created.
    - `tags`- (Optional, Array of Strings) A list of user tags that you want to add to your snapshot. (https://cloud.ibm.com/apidocs/tagging#types-of-tags)
- `tags`- (Optional, Array of Strings) A list of user tags that you want to add to your snapshot. (https://cloud.ibm.com/apidocs/tagging#types-of-tags)


## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

- `id` - The unique identifier of the SnapshotConsistencyGroup.
- `backup_policy_plan` - (List) If present, the backup policy plan which created this snapshot consistency group.
  
   Nested scheme for `backup_policy_plan`:
    - `deleted` - (List) If present, this property indicates the referenced resource has been deleted and provides some supplementary information.
   
      Nested scheme for `deleted`:
      - `more_info` - (String) Link to documentation about deleted resources.
    - `href` - (String) The URL for this backup policy plan.
    - `id` - (String) The unique identifier for this backup policy plan.
    - `name` - (String) The unique user defined name for this backup policy plan. If unspecified, the name will be a hyphenated list of randomly selected words.
    - `resource_type` - (String) The type of resource referenced.
    - `remote` - (List) If present, this property indicates that the resource associated with this referenceis remote and therefore may not be directly retrievable.
      
      Nested schema for `remote`:
      - `href` - (String) The URL for this region.
      - `name` - (String) The globally unique name for this region.

- `created_at` - (String) The date and time that this snapshot consistency group was created.
- `crn` - (String) The CRN of this snapshot consistency group.
- `href` - (String) The URL for this snapshot consistency group.
- `lifecycle_state` - (String) The lifecycle state of this snapshot consistency group.
- `resource_type` - (String) The resource type.
- `snapshots_reference` - (List) The member snapshots that are data-consistent with respect to captured time. (may be[deleted](https://cloud.ibm.com/apidocs/vpc#deleted-resources)).

  Nested scheme for `snapshots_reference`:
  - `crn` - (String) The CRN of the source snapshot.
  - `deleted` - (List) If present, this property indicates the referenced resource has been deleted and provides some supplementary information.
   
      Nested scheme for `deleted`:
      - `more_info` - (String) Link to documentation about deleted resources.
  - `href` - (String) The URL for the source snapshot.
  - `id` - (String) The unique identifier for the source snapshot.
  - `name` - (String) The name for the source snapshot. The name is unique across all snapshots in the source snapshot's native region.
  - `remote` - (List) If present, this property indicates the referenced resource is remote to this region,and identifies the native region
    
      Nested scheme for `remote`:
      - `href` - (String) The URL for this region.
      - `name` - (String) The globally unique name for this region.
  - `resource_type` - (String) The resource type.
- `service_tags` - (List) The [service tags](https://cloud.ibm.com/apidocs/tagging#types-of-tags)[`is.instance:` prefix](https://cloud.ibm.com/docs/vpc?topic=vpc-snapshots-vpc-faqs) associated with this snapshot consistency group.


## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import the `ibm_is_snapshot_consistency_group` resource by using `id`.
The `id` property can be formed using the snapshot_consistency_group identifier. For example:

```terraform
import {
  to = ibm_is_snapshot_consistency_group.is_snapshot_consistency_group
  id = "<id>"
}
```

Using `terraform import`. For example:

```console
% terraform import ibm_is_snapshot_consistency_group.is_snapshot_consistency_group <id>
```