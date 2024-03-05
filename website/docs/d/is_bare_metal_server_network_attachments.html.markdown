---
layout: "ibm"
page_title: "IBM : ibm_is_bare_metal_server_network_attachments"
description: |-
  Get information about BareMetalServerNetworkAttachmentCollection
subcategory: "VPC infrastructure"
---

# ibm_is_bare_metal_server_network_attachments

Provides a read-only data source to retrieve information about a BareMetalServerNetworkAttachmentCollection. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```terraform
data "ibm_is_bare_metal_server_network_attachments" "example" {
	bare_metal_server = ibm_is_bare_metal_server_network_attachment.example.bare_metal_server
}
```

## Argument Reference

You can specify the following arguments for this data source.

- `bare_metal_server` - (Required, Forces new resource, String) The bare metal server identifier.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

- `id` - The unique identifier of the BareMetalServerNetworkAttachmentCollection.
- `network_attachments` - (List) Collection of bare metal server network attachments.
	Nested schema for **network_attachments**:
	- `allow_to_float` - (Boolean) Indicates if the bare metal server network attachment can automatically float to any other server within the same `resource_group`. The bare metal server network attachment will float automatically if the network detects a GARP or RARP on another bare metal server in the resource group. Applies only to bare metal server network attachments with `vlan` interface type.
	- `allowed_vlans` - (List)
	- `created_at` - (String) The date and time that the bare metal server network attachment was created.
	- `href` - (String) The URL for this bare metal server network attachment.
	- `id` - (String) The unique identifier for this bare metal server network attachment.
	- `interface_type` - (String) The network attachment's interface type:- `pci`: a physical PCI device which can only be created or deleted when the bare metal  server is stopped  - Has an `allowed_vlans` property which controls the VLANs that will be permitted    to use the PCI attachment  - Cannot directly use an IEEE 802.1q VLAN tag.- `vlan`: a virtual device, used through a `pci` device that has the `vlan` in its  array of `allowed_vlans`.  - Must use an IEEE 802.1q tag.The enumerated values for this property are expected to expand in the future. When processing this property, check for and log unknown values. Optionally halt processing and surface the error, or bypass the resource on which the unexpected property value was encountered.
	- `lifecycle_state` - (String) The lifecycle state of the bare metal server network attachment.
	- `name` - (String) The name for this bare metal server network attachment. The name is unique across all network attachments for the bare metal server.
	- `port_speed` - (Integer) The port speed for this bare metal server network attachment in Mbps.
	- `primary_ip` - (List) The primary IP address of the virtual network interface for the bare metal servernetwork attachment.
		Nested schema for **primary_ip**:
		- `address` - (String) The IP address.If the address has not yet been selected, the value will be `0.0.0.0`.This property may add support for IPv6 addresses in the future. When processing a value in this property, verify that the address is in an expected format. If it is not, log an error. Optionally halt processing and surface the error, or bypass the resource on which the unexpected IP address format was encountered.
		- `deleted` - (List) If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.
			Nested schema for **deleted**:
			- `more_info` - (String) Link to documentation about deleted resources.
		- `href` - (String) The URL for this reserved IP.
		- `id` - (String) The unique identifier for this reserved IP.
		- `name` - (String) The name for this reserved IP. The name is unique across all reserved IPs in a subnet.
		- `resource_type` - (String) The resource type.
	- `resource_type` - (String) The resource type.
	- `subnet` - (List) The subnet of the virtual network interface for the bare metal server networkattachment.
		Nested schema for **subnet**:
		- `crn` - (String) The CRN for this subnet.
		- `deleted` - (List) If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.
			Nested schema for **deleted**:
			- `more_info` - (String) Link to documentation about deleted resources.
		- `href` - (String) The URL for this subnet.
		- `id` - (String) The unique identifier for this subnet.
		- `name` - (String) The name for this subnet. The name is unique across all subnets in the VPC.
		- `resource_type` - (String) The resource type.
	- `type` - (String) The bare metal server network attachment type.
	- `virtual_network_interface` - (List) The virtual network interface for this bare metal server network attachment.
		Nested schema for **virtual_network_interface**:
		- `crn` - (String) The CRN for this virtual network interface.
		- `href` - (String) The URL for this virtual network interface.
		- `id` - (String) The unique identifier for this virtual network interface.
		- `name` - (String) The name for this virtual network interface. The name is unique across all virtual network interfaces in the VPC.
		- `resource_type` - (String) The resource type.
	- `vlan` - (Integer) Indicates the 802.1Q VLAN ID tag that must be used for all traffic on this attachment.