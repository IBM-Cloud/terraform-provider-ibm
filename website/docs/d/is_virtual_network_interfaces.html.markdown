---
layout: "ibm"
page_title: "IBM : ibm_is_virtual_network_interfaces"
description: |-
  Get information about Virtual Network Interfaces Collection
subcategory: "VPC infrastructure"
---

# ibm_is_virtual_network_interfaces

Provides a read-only data source for VirtualNetworkInterfaceCollection. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl
data "ibm_is_virtual_network_interfaces" "example" {
}
```


## Attribute Reference

- `resource_group` - (Optional, String) The ID of the Resource group these virtual network interfaces belong to.

In addition to all argument references listed, you can access the following attribute references after your data source is created.

- `virtual_network_interfaces` - (List) Collection of virtual network interfaces.
	Nested scheme for **virtual_network_interfaces**:
	- `access_tags`  - (Array of Strings) Access management tags associated for the virtual network interface.
	- `allow_ip_spoofing` - (Boolean) Indicates whether source IP spoofing is allowed on this interface. If `false`, source IP spoofing is prevented on this interface. If `true`, source IP spoofing is allowed on this interface.
	- `auto_delete` - (Boolean) Indicates whether this virtual network interface will be automatically deleted when`target` is deleted.
	- `created_at` - (String) The date and time that the virtual network interface was created.
	- `crn` - (String) The CRN for this virtual network interface.
	- `enable_infrastructure_nat` - (Boolean) If `true`:- The VPC infrastructure performs any needed NAT operations.- `floating_ips` must not have more than one floating IP.If `false`:- Packets are passed unchanged to/from the virtual network interface,  allowing the workload to perform any needed NAT operations.- `allow_ip_spoofing` must be `false`.- If the virtual network interface is attached:  - The target `resource_type` must be `bare_metal_server_network_attachment`.  - The target `interface_type` must not be `hipersocket`.
	- `href` - (String) The URL for this virtual network interface.
	- `id` - (String) The unique identifier for this virtual network interface.
	- `ips` - (List) The reserved IPs bound to this virtual network interface.May be empty when `lifecycle_state` is `pending`.
		Nested schema for **ips**:
		- `address` - (String) The IP address.If the address has not yet been selected, the value will be `0.0.0.0`.This property may add support for IPv6 addresses in the future. When processing a value in this property, verify that the address is in an expected format. If it is not, log an error. Optionally halt processing and surface the error, or bypass the resource on which the unexpected IP address format was encountered.
		- `deleted` - (List) If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.
			Nested schema for **deleted**:
			- `more_info` - (String) Link to documentation about deleted resources.
		- `href` - (String) The URL for this reserved IP.
		- `id` - (String) The unique identifier for this reserved IP.
		- `name` - (String) The name for this reserved IP. The name is unique across all reserved IPs in a subnet.
		- `resource_type` - (String) The resource type.
	- `lifecycle_state` - (String) The lifecycle state of the virtual network interface.
	- `mac_address` - (String) The MAC address of the virtual network interface. May be absent if `lifecycle_state` is `pending`.
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
	- `protocol_state_filtering_mode` - (String) The protocol state filtering mode to use for this virtual network interface.
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
			Nested scheme for --deleted--:
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
	- `tags` - (Array of Strings) The tags associated with the virtual netork interface.
	- `target` - (List) The target of this virtual network interface.If absent, this virtual network interface is not attached to a target.
		Nested scheme for **target**:
		- `deleted` - (List) If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.
			Nested scheme for --deleted--:
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

