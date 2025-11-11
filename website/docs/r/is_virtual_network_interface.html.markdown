---
layout: "ibm"
page_title: "IBM : ibm_is_virtual_network_interface"
description: |-
  Manages Virtual Network Interface.
subcategory: "VPC infrastructure"
---

# ibm_is_virtual_network_interface

Create, update, and delete VirtualNetworkInterfaces with this resource.

## Example Usage

```terraform
resource "ibm_is_virtual_network_interface" "is_virtual_network_interface_instance" {
  allow_ip_spoofing = true
  auto_delete = false
  enable_infrastructure_nat = true
  name = "my-virtual-network-interface"
  subnet = "7ec86020-1c6e-4889-b3f0-a15f2e50f87e"
}
```
## Example Usage with protocol_state_filtering_mode enabled

```terraform
resource "ibm_is_virtual_network_interface" "is_virtual_network_interface_instance" {
  allow_ip_spoofing = true
  auto_delete = false
  enable_infrastructure_nat = true
  name = "my-virtual-network-interface"
  subnet = "7ec86020-1c6e-4889-b3f0-a15f2e50f87e"
  protocol_state_filtering_mode = "enabled"
}
```
## Argument Reference

You can specify the following arguments for this resource.


- `access_tags`  - (Optional, List of Strings) A list of access management tags to attach to the virtual network interface.

  ~> **Note:** 
  **&#x2022;** You can attach only those access tags that already exists.</br>
  **&#x2022;** For more information, about creating access tags, see [working with tags](https://cloud.ibm.com/docs/account?topic=account-tag&interface=ui#create-access-console).</br>
  **&#x2022;** You must have the access listed in the [Granting users access to tag resources](https://cloud.ibm.com/docs/account?topic=account-access) for `access_tags`</br>
  **&#x2022;** `access_tags` must be in the format `key:value`.

- `allow_ip_spoofing` - (Optional, Boolean) Indicates whether source IP spoofing is allowed on this interface. If `false`, source IP spoofing is prevented on this interface. If `true`, source IP spoofing is allowed on this interface.
- `auto_delete` - (Optional, Boolean) Indicates whether this virtual network interface will be automatically deleted when`target` is deleted. Must be false if the virtual network interface is unbound.
- `enable_infrastructure_nat` - (Optional, Boolean) If `true`:- The VPC infrastructure performs any needed NAT operations.- `floating_ips` must not have more than one floating IP.If `false`:- Packets are passed unchanged to/from the network interface,  allowing the workload to perform any needed NAT operations.- `allow_ip_spoofing` must be `false`.- If the virtual network interface is attached:  - The target `resource_type` must be `bare_metal_server_network_attachment`.  - The target `interface_type` must not be `hipersocket`.

~> **NOTE** to add `ips` only existing `reserved_ip` is supported, new reserved_ip creation is not supported as it leads to unmanaged(dangling) reserved ips. Use `ibm_is_subnet_reserved_ip` to create a reserved_ip
- `ips` - (Optional, List) The reserved IPs bound to this virtual network interface.May be empty when `lifecycle_state` is `pending`.
	Nested schema for **ips**:
	- `reserved_ip` - (Required, String) The unique identifier for this reserved IP.
- `name` - (Optional, String) The name for this virtual network interface. The name is unique across all virtual network interfaces in the VPC.
- `protocol_state_filtering_mode` - (Optional, String) The protocol state filtering mode to use for this virtual network interface. 

  ~> **If auto, protocol state packet filtering is enabled or disabled based on the virtual network interface's target resource type:** 
  **&#x2022;** bare_metal_server_network_attachment: disabled </br>
  **&#x2022;** instance_network_attachment: enabled </br>
  **&#x2022;** share_mount_target: enabled </br>
- `primary_ip` - (Optional, List) The reserved IP for this virtual network interface.
	Nested schema for **primary_ip**:
	- `address` - (Required, String) The IP address.If the address has not yet been selected, the value will be `0.0.0.0`.This property may add support for IPv6 addresses in the future. When processing a value in this property, verify that the address is in an expected format. If it is not, log an error. Optionally halt processing and surface the error, or bypass the resource on which the unexpected IP address format was encountered.
	- `deleted` - (Optional, List) If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.
	Nested schema for **deleted**:
		- `more_info` - (Required, String) Link to documentation about deleted resources.
	- `href` - (Required, String) The URL for this reserved IP.
	- `reserved_ip` - (Required, String) The unique identifier for this reserved IP.
	- `name` - (Required, String) The name for this reserved IP. The name is unique across all reserved IPs in a subnet.
	- `resource_type` - (Computed, String) The resource type.
- `resource_group` - (Optional, String) The resource group id for this virtual network interface.
- `security_groups` - (Optional, Array of string) The security group ids list for this virtual network interface.
- `subnet` - (Optional, List) The associated subnet id.
- `tags` (Optional, Array of Strings) Enter any tags that you want to associate with your VPC. Tags might help you find your VPC more easily after it is created. Separate multiple tags with a comma (`,`).

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

- `id` - The unique identifier of the VirtualNetworkInterface.
- `created_at` - (String) The date and time that the virtual network interface was created.
- `crn` - (String) The CRN for this virtual network interface.
- `href` - (String) The URL for this virtual network interface.
- `lifecycle_state` - (String) The lifecycle state of the virtual network interface. Allowable values are: `deleting`, `failed`, `pending`, `stable`, `suspended`, `updating`, `waiting`.
- `mac_address` - (String) The MAC address of the interface. Absent when the interface is not attached to a target.
- `resource_type` - (String) The resource type.
- `target` - (List) The target of this virtual network interface.If absent, this virtual network interface is not attached to a target.
	Nested schema for **target**:
	- `deleted` - (List) If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.
	Nested schema for **deleted**:
		- `more_info` - (String) Link to documentation about deleted resources.
	- `href` - (String) The URL for this share mount target.
	- `id` - (String) The unique identifier for this share mount target.
	- `name` - (String) The name for this share mount target. The name is unique across all mount targets for the file share.
	- `resource_type` - (String) The resource type.
- `vpc` - (List) The VPC this virtual network interface resides in.
	Nested schema for **vpc**:
	- `crn` - (String) The CRN for this VPC.
	- `deleted` - (List) If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.
	Nested schema for **deleted**:
		- `more_info` - (String) Link to documentation about deleted resources.
	- `href` - (String) The URL for this VPC.
	- `id` - (String) The unique identifier for this VPC.
	- `name` - (String) The name for this VPC. The name is unique across all VPCs in the region.
	- `resource_type` - (String) The resource type.
- `zone` - (String) The zone name of the zone this virtual network interface resides in.


## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import the `ibm_is_virtual_network_interface` resource by using `id`.
The `id` property can be formed using the VNI identifier. For example:

```terraform
import {
  to = ibm_is_virtual_network_interface.is_virtual_network_interface
  id = "<id>"
}
```

Using `terraform import`. For example:

```console
% terraform import ibm_is_virtual_network_interface.is_virtual_network_interface <id>
```