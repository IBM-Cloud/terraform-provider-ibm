---
layout: "ibm"
page_title: "IBM : is_share_mount_target"
description: |-
  Get information about ShareMountTarget
subcategory: "VPC infrastructure"
---

# ibm\_is_share_mount_target

Provides a read-only data source for ShareMountTarget. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.


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

- `access_control_mode` - (String) The access control mode for the share.
- `created_at` - (String) The date and time that the share target was created.
- `href` - (String) The URL for this share target.
- `lifecycle_state` - (String) The lifecycle state of the mount target.
- `mount_path` - (String) The mount path for the share. The server component of the mount path may be either an IP address or a fully qualified domain name.

    This property will be absent if the lifecycle_state of the mount target is 'pending', failed, or deleting.

    -> **If the share's access_control_mode is:**
    &#x2022; security_group: The IP address used in the mount path is the primary_ip address of the virtual network interface for this share mount target. </br>
    &#x2022; vpc: The fully-qualified domain name used in the mount path is an address that resolves to the share mount target. </br>
- `name` - The user-defined name for this share target.
- `resource_type` - (String) The type of resource referenced.
- `transit_encryption` - (String) The transit encryption mode for this share target.
- `vpc` - (List) The VPC to which this share target is allowing to mount the file share. Nested `vpc` blocks have the following structure:
	- `crn` - (String) The CRN for this VPC.
	- `deleted` - (String) If present, this property indicates the referenced resource has been deleted and providessome supplementary information. Nested `deleted` blocks have the following structure:
		- `more_info` - (String) Link to documentation about deleted resources.
	- `href` - (String) The URL for this VPC.
	- `id` - (String) The unique identifier for this VPC.
	- `name` - (String) The unique user-defined name for this VPC.
	- `resource_type` - (String) The resource type.
- `subnet` - (List) The subnet of the virtual network interface for the share mount target. Nested `subnet` blocks have the following structure:
	- `crn` - (String) The CRN for this subnet.
	- `deleted` - (String) If present, this property indicates the referenced resource has been deleted and providessome supplementary information. Nested `deleted` blocks have the following structure:
		- `more_info` - (String) Link to documentation about deleted resources.
	- `href` - (String) The URL for this subnet.
	- `id` - (String) The unique identifier for this subnet.
	- `name` - (String) The unique user-defined name for this subnet.
	- `resource_type` - (String) The resource type.
- `virtual_network_interface` - (List) The virtual network interface for this file share mount target.. Nested `virtual_network_interface` blocks have the following structure:
	- `crn` - (String) The CRN for this virtual network interface.
	- `href` - (String) The URL for this virtual network interface.
	- `id` - (String) The unique identifier for this virtual network interface.
	- `name` - (String) The unique user-defined name for this virtual network interface.
	- `resource_type` - (String) The resource type.
