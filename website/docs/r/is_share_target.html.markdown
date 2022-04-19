---
layout: "ibm"
page_title: "IBM : is_share_target"
description: |-
  Manages ShareTarget.
subcategory: "Virtual Private Cloud API"
---

# ibm\_is_share_target

Provides a resource for ShareTarget. This allows ShareTarget to be created, updated and deleted.

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

resource "is_share_target" "is_share_target" {
  share = is_share.is_share.id
  vpc = ibm_is_vpc.vpc.id
  name = "my-share-target"
}`
```

## Argument Reference

The following arguments are supported:

* `share` - (Required, string) The file share identifier.
* `vpc` - (Required, string) The VPC in which instances can mount the file share using this share target.This property will be removed in a future release.The `subnet` property should be used instead.
* `name` - (Required, string) The user-defined name for this share target. Names must be unique within the share the share target resides in. If unspecified, the name will be a hyphenated list of randomly-selected words.
* `subnet` - (Optional, string) The unique identifier of the subnet associated with this file share target.Only virtual server instances in the same VPC as this subnetwill be allowed to mount the file share.In the future, this property may be required and used to assignan IP address for the file share target.
## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the ShareTarget. The id is composed of \<ibm_is_share_id\>/\<ibm_is_share_target_id\>
* `share_target` - The unique identifier of the share target
* `created_at` - The date and time that the share target was created.
* `href` - The URL for this share target.
* `lifecycle_state` - The lifecycle state of the mount target.
* `mount_path` - The mount path for the share.The IP addresses used in the mount path are currently within the IBM services IP range, but are expected to change to be within one of the VPC's subnets in the future.
* `resource_type` - The type of resource referenced.
