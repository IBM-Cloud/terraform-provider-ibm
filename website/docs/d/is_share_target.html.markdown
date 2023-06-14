---
layout: "ibm"
page_title: "IBM : is_share_mount_target"
description: |-
  Get information about ShareTarget
subcategory: "VPC infrastructure"
---

# ibm\is_share_mount_target

Provides a read-only data source for ShareTarget. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.


~> **NOTE**
IBM Cloud® File Storage for VPC is available for customers with special approval. Contact your IBM Sales representative if you are interested in getting access.

~> **NOTE**
This is a Beta feature and it is subject to change in the GA release 

~> **NOTE**
This data source is being deprecated. Please use `ibm_is_share_mount_target` instead


## Example Usage

```hcl
resource "ibm_is_vpc" "example" {
  name = "example-vpc"
}
resource "ibm_is_share" "example" {
  name    = "example-share"
  size    = 200
  profile = "tier-3iops"
  zone    = "us-south-2"
}

resource "ibm_is_share_mount_target" "example" {
  share = ibm_is_share.is_share.id
  vpc   = ibm_is_vpc.example.id
  name  = "example-share-target"
}

data "ibm_is_share_mount_target" "example" {
  share        = ibm_is_share.example.id
  mount_target = ibm_is_share_mount_target.example.mount_target
}
```

## Argument Reference

The following arguments are supported:

- `share` - (Required, string) The file share identifier.
- `mount_target` - (Required, string) The share target identifier.

## Attribute Reference

The following attributes are exported:

- `created_at` - The date and time that the share target was created.
- `href` - The URL for this share target.
- `lifecycle_state` - The lifecycle state of the mount target.
- `mount_path` - The mount path for the share.The IP addresses used in the mount path are currently within the IBM services IP range, but are expected to change to be within one of the VPC's subnets in the future.
- `name` - The user-defined name for this share target.
- `resource_type` - The type of resource referenced.
- `subnet` - The subnet associated with this file share target. Nested `subnet` blocks have the following structure:
	- `crn` - The CRN for this subnet.
	- `deleted` - If present, this property indicates the referenced resource has been deleted and providessome supplementary information. Nested `deleted` blocks have the following structure:
		- `more_info` - Link to documentation about deleted resources.
	- `href` - The URL for this subnet.
	- `id` - The unique identifier for this subnet.
	- `name` - The user-defined name for this subnet.
	- `resource_type` - The resource type.
- `vpc` - The VPC to which this share target is allowing to mount the file share. Nested `vpc` blocks have the following structure:
	- `crn` - The CRN for this VPC.
	- `deleted` - If present, this property indicates the referenced resource has been deleted and providessome supplementary information. Nested `deleted` blocks have the following structure:
		- `more_info` - Link to documentation about deleted resources.
	- `href` - The URL for this VPC.
	- `id` - The unique identifier for this VPC.
	- `name` - The unique user-defined name for this VPC.
	- `resource_type` - The resource type.

