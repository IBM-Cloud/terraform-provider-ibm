---
layout: "ibm"
page_title: "IBM : is_share_mount_targets"
description: |-
  Get information about ShareMountTargetCollection
subcategory: "VPC infrastructure"
---

# ibm\_is_share_mount_targets

Provides a read-only data source for ShareMountTargetCollection. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.


~> **NOTE**
IBM Cloud® File Storage for VPC is available for customers with special approval. Contact your IBM Sales representative if you are interested in getting access.

~> **NOTE**
This is a Beta feature and it is subject to change in the GA release 



## Example Usage

```hcl
resource "ibm_is_vpc" "example" {
  name = "example-vpc"
}
resource "ibm_is_share" "example" {
  name    = "example-share"
  size    = 200
  profile = "dp2"
  zone    = "us-south-2"
}

data "ibm_is_share_mount_targets" "example" {
  share = ibm_is_share.example.id
}
```

## Argument Reference

The following arguments are supported:

- `share` - (Required, string) The file share identifier.

## Attribute Reference

The following attributes are exported:

- `id` - The unique identifier of the ShareTargetCollection.
- `mount_targets` - Collection of share targets. Nested `targets` blocks have the following structure:
	- `created_at` - The date and time that the share target was created.
	- `href` - The URL for this share target.
	- `id` - The unique identifier for this share target.
	- `lifecycle_state` - The lifecycle state of the mount target.
	- `mount_path` - The mount path for the share.The IP addresses used in the mount path are currently within the IBM services IP range, but are expected to change to be within one of the VPC's subnets in the future.
	- `name` - The user-defined name for this share target.
	- `resource_type` - The type of resource referenced.
	- `transit_encryption` - (String) The transit encryption mode for this share target.
	- `vpc` - The VPC to which this share target is allowing to mount the file share. Nested `vpc` blocks have the following structure:
		- `crn` - The CRN for this VPC.
		- `deleted` - If present, this property indicates the referenced resource has been deleted and providessome supplementary information. Nested `deleted` blocks have the following structure:
			- `more_info` - Link to documentation about deleted resources.
		- `href` - The URL for this VPC.
		- `id` - The unique identifier for this VPC.
		- `name` - The unique user-defined name for this VPC.
		- `resource_type` - The resource type.

