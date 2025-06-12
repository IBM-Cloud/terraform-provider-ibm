---
layout: "ibm"
page_title: "IBM : is_share_mount_targets"
description: |-
  Get information about ShareMountTargetCollection
subcategory: "VPC infrastructure"
---

# ibm\_is_share_mount_targets

Provides a read-only data source for ShareMountTargetCollection. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

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
	- `primary_ip` - The primary IP address of the virtual network interface for the share mount target. Nested `primary_ip` blocks have the following structure:
		- `address` - The IP address.
		- `deleted` - If present, this property indicates the referenced resource has been deleted and providessome supplementary information. Nested `deleted` blocks have the following structure:
			- `more_info` - Link to documentation about deleted resources.
		- `href` - The URL for this reserved IP.
		- `id` - The unique identifier for this reserved IP.
		- `name` - The name for this reserved IP. The name is unique across all reserved IPs in a subnet.
		- `resource_type` - The resource type.
	- `subnet` - The subnet of the virtual network interface for the share mount target. Nested `vpc` blocks have the following structure:
		- `crn` - The CRN for this subnet.
		- `deleted` - If present, this property indicates the referenced resource has been deleted and providessome supplementary information. Nested `deleted` blocks have the following structure:
			- `more_info` - Link to documentation about deleted resources.
		- `href` - The URL for this subnet.
		- `id` - The unique identifier for this subnet.
		- `name` - The unique user-defined name for this subnet.
		- `resource_type` - The resource type.
	- `virtual_network_interface` - The virtual network interface for this file share mount target.. Nested `subnet` blocks have the following structure:
		- `crn` - The CRN for this virtual network interface.
		- `href` - The URL for this virtual network interface.
		- `id` - The unique identifier for this virtual network interface.
		- `name` - The unique user-defined name for this virtual network interface.
		- `resource_type` - The resource type.

