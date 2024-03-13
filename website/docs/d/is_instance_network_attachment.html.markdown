---
layout: "ibm"
page_title: "IBM : ibm_is_instance_network_attachment"
description: |-
  Get information about InstanceNetworkAttachment
subcategory: "VPC infrastructure"
---

# ibm_is_instance_network_attachment

Provides a read-only data source to retrieve information about an Instance NetworkAttachment. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```terraform
data "ibm_is_instance_network_attachment" "example" {
	instance				= ibm_is_instance.example.id
	network_attachment		= ibm_is_instance.example.primary_network_attachment.0.id
}
```

## Argument Reference

You can specify the following arguments for this data source.

- `instance` - (Required, Forces new resource, String) The virtual server instance identifier.
- `network_attachment` - (Required, Forces new resource, String) The instance network attachment identifier.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

- `id` - The unique identifier of the Instance NetworkAttachment.`<instance>/<network_attachment>`
- `created_at` - (String) The date and time that the instance network attachment was created.
- `href` - (String) The URL for this instance network attachment.
- `lifecycle_state` - (String) The lifecycle state of the instance network attachment. Allowable values are: `deleting`, `failed`, `pending`, `stable`, `suspended`, `updating`, `waiting`.
- `name` - (String) The name for this instance network attachment. The name is unique across all network attachments for the instance.
- `port_speed` - (Integer) The port speed for this instance network attachment in Mbps.
- `primary_ip` - (List) The primary IP address of the virtual network interface for the instance networkattachment.
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
- `subnet` - (List) The subnet of the virtual network interface for the instance network attachment.
	Nested schema for **subnet**:
	- `crn` - (String) The CRN for this subnet.
	- `deleted` - (List) If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.
	Nested schema for **deleted**:
		- `more_info` - (String) Link to documentation about deleted resources.
	- `href` - (String) The URL for this subnet.
	- `id` - (String) The unique identifier for this subnet.
	- `name` - (String) The name for this subnet. The name is unique across all subnets in the VPC.
	- `resource_type` - (String) The resource type.

- `type` - (String) The instance network attachment type.
- `virtual_network_interface` - (List) The virtual network interface for this instance network attachment.
	Nested schema for **virtual_network_interface**:
	- `crn` - (String) The CRN for this virtual network interface.
	- `href` - (String) The URL for this virtual network interface.
	- `id` - (String) The unique identifier for this virtual network interface.
	- `name` - (String) The name for this virtual network interface. The name is unique across all virtual network interfaces in the VPC.
	- `resource_type` - (String) The resource type.

