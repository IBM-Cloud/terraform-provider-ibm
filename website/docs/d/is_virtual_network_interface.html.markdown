---
layout: "ibm"
page_title: "IBM : ibm_is_virtual_network_interface"
description: |-
  Get information about Virtual Network Interface
subcategory: "VPC infrastructure"
---

# ibm_is_virtual_network_interface

Provides a read-only data source for VirtualNetworkInterface. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
resource "ibm_is_vpc" "example" {
  name = "my-vpc"
}
resource "ibm_is_subnet" "example" {
  name                     = "example-subnet"
  vpc                      = ibm_is_vpc.example.id
  zone                     = "us-south-1"
  total_ipv4_address_count = 256
}
resource "ibm_is_share" "example" {
  name = "my-share"
  access_control_mode = "security_group"
  size = 200
  profile = "dp2"
  zone = "us-south-2"
}
resource "ibm_is_share_mount_target" "example" {
  share = ibm_is_share.example.id
  virtual_network_interface {
    name = "my-virtual_network_interface"
    primary_ip {
      address = "10.240.64.5"
      auto_delete = true
      name = "my-reserved-ip"
    }
    name = "my-share-target"
  }
}

data "ibm_is_virtual_network_interface" "example" {
	virtual_network_interface = ibm_is_share_mount_target.example.virtual_network_interface.0.id
}
```

## Argument Reference

Review the argument reference that you can specify for your data source.

- `virtual_network_interface` - (Required, String) The virtual network interface identifier.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

- `auto_delete` - (Boolean) Indicates whether this virtual network interface will be automatically deleted when`target` is deleted.
- `created_at` - (String) The date and time that the virtual network interface was created.
- `crn` - (String) The CRN for this virtual network interface.
- `href` - (String) The URL for this virtual network interface.
- `id` - The unique identifier of the VirtualNetworkInterface.
- `lifecycle_state` - (String) The lifecycle state of the virtual network interface.
- `name` - (String) The name for this virtual network interface. The name is unique across all virtual network interfaces in the VPC.
- `primary_ip` - (List) The reserved IP for this virtual network interface.May be absent when `lifecycle_state` is `pending`.
	Nested scheme for **primary_ip**:
	- `address` - (String) The IP address.If the address has not yet been selected, the value will be `0.0.0.0`.This property may add support for IPv6 addresses in the future. When processing a value in this property, verify that the address is in an expected format. If it is not, log an error. Optionally halt processing and surface the error, or bypass the resource on which the unexpected IP address format was encountered.
	  - `deleted` - (List) If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.
		Nested scheme for **deleted**:
		- `more_info` - (String) Link to documentation about deleted resources.
	- `href` - (String) The URL for this reserved IP.
	- `id` - (String) The unique identifier for this reserved IP.
	- `name` - (String) The name for this reserved IP. The name is unique across all reserved IPs in a subnet.
	- `resource_type` - (String) The resource type.
- `resource_group` - (List) The resource group for this virtual network interface.
	Nested scheme for **resource_group**:
	- `href` - (String) The URL for this resource group.
	- `id` - (String) The unique identifier for this resource group.
	- `name` - (String) The name for this resource group.
- `resource_type` - (String) The resource type.
- `security_groups` - (List) The security groups for this virtual network interface.
	Nested scheme for **security_groups**:
	- `crn` - (String) The security group's CRN.
	- `deleted` - (List) If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.
		Nested scheme for **deleted**:
		- `more_info` - (String) Link to documentation about deleted resources.
	- `href` - (String) The security group's canonical URL.
	- `id` - (String) The unique identifier for this security group.
	- `name` - (String) The name for this security group. The name is unique across all security groups for the VPC.
- `subnet` - (List) The associated subnet.
	Nested scheme for **subnet**:
	- `crn` - (String) The CRN for this subnet.
	- `deleted` - (List) If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.
		Nested scheme for **deleted**:
		- `more_info` - (String) Link to documentation about deleted resources.
	- `href` - (String) The URL for this subnet.
	- `id` - (String) The unique identifier for this subnet.
	- `name` - (String) The name for this subnet. The name is unique across all subnets in the VPC.
	- `resource_type` - (String) The resource type.
- `target` - (List) The target of this virtual network interface.If absent, this virtual network interface is not attached to a target.
	Nested scheme for **target**:
	- `deleted` - (List) If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.
		Nested scheme for **deleted**:
		- `more_info` - (String) Link to documentation about deleted resources.
	- `href` - (String) The URL for this share mount target.  
	- `id` - (String) The unique identifier for this share mount target.  
	- `name` - (String) The name for this share mount target. The name is unique across all targets for the file share.
	- `resource_type` - (String) The resource type.
- `vpc` - (List) The VPC this virtual network interface resides in.
	Nested scheme for **vpc**:
	- `crn` - (String) The CRN for this VPC.
	- `deleted` - (List) If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.
		Nested scheme for **deleted**:
		- `more_info` - (String) Link to documentation about deleted resources.
	- `href` - (String) The URL for this VPC.
	- `id` - (String) The unique identifier for this VPC.
	- `name` - (String) The name for this VPC. The name is unique across all VPCs in the region.
	- `resource_type` - (String) The resource type.
- `zone` - (List) The zone this virtual network interface resides in.
	Nested scheme for **zone**:
	- `href` - (String) The URL for this zone.
	- `name` - (String) The globally unique name for this zone.
	  

