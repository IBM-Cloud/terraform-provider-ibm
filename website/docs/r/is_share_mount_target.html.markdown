---
layout: "ibm"
page_title: "IBM : is_share_mount_target"
description: |-
  Manages ShareTarget.
subcategory: "VPC infrastructure"
---


# ibm\_is_share_mount_target

Provides a resource for ShareMountTarget. This allows ShareTarget to be created, updated and deleted.


~> **NOTE**
IBM Cloud® File Storage for VPC is available for customers with special approval. Contact your IBM Sales representative if you are interested in getting access.

~> **NOTE**
This is a Beta feature and it is subject to change in the GA release 

## Example Usage

```hcl
resource "ibm_is_vpc" "vpc" {
  name = "my-vpc"
}

resource "ibm_is_share" "is_share" {
  name = "my-share"
  size = 200
  profile = "tier-3iops"
  zone = "us-south-2"
}

resource "ibm_is_share_mount_target" "is_share_target" {
  share = ibm_is_share.is_share.id
  vpc = ibm_is_vpc.vpc.id
  name = "my-share-target"
}`
```

## Argument Reference

The following arguments are supported:

- `share` - (Required, String) The file share identifier.
- `vpc` - (Required, String) The VPC in which instances can mount the file share using this share target.This property will be removed in a future release.The `subnet` property should be used instead.
- `name` - (Required, String) The user-defined name for this share target. Names must be unique within the share the share target resides in. If unspecified, the name will be a hyphenated list of randomly-selected words.
- `transit_encryption` - (Optional, String) The transit encryption mode for this share target. Supported values are **none**, **user_managed**. Default is **none**
## Attribute Reference

The following attributes are exported:


- `mount_target` - The unique identifier of the share target
- `created_at` - The date and time that the share target was created.
- `href` - The URL for this share target.
- `id` - The unique identifier of the ShareTarget. The id is composed of \<ibm_is_share_id\>/\<ibm_is_share_target_id\>
- `lifecycle_state` - The lifecycle state of the mount target.
- `mount_path` - The mount path for the share.The IP addresses used in the mount path are currently within the IBM services IP range, but are expected to change to be within one of the VPC's subnets in the future.
- `resource_type` - The type of resource referenced.
- `transit_encryption` - (String) The transit encryption mode for this share target.

## Import

The `ibm_is_share_target` can be imported using ID.

**Syntax**

```
$ terraform import ibm_is_share_target.example `\<ibm_is_share_id\>/\<ibm_is_share_target_id\>`
```

**Example**

```
$ terraform import ibm_is_share_target.example d7bec597-4726-451f-8a63-e62e6f19c32c/d7bec597-4726-451f-8a63-e62e6f19c32c
```