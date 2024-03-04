---
layout: "ibm"
page_title: "IBM : ibm_is_instance_network_attachment"
description: |-
  Manages Instance NetworkAttachment.
subcategory: "Virtual Private Cloud API"
---

# ibm_is_instance_network_attachment

Create, update, and delete Instance NetworkAttachment with this resource.

## Example Usage

```terraform
resource "ibm_is_instance_network_attachment" "example" {
  instance = "<instance_id>"
  virtual_network_interface {
		id = "<virtual_network_interface_id>"
  }
}
```

## Argument Reference

You can specify the following arguments for this resource.

- `instance` - (Required, Forces new resource, String) The virtual server instance identifier.
- `name` - (Optional, String) The name for this instance network attachment. The name is unique across all network attachments for the instance.
- `virtual_network_interface` - (Required, List) The virtual network interface for this instance network attachment.
	Nested schema for **virtual_network_interface**:
	- `crn` - (Required, String) The CRN for this virtual network interface.
	- `href` - (Required, String) The URL for this virtual network interface.
	- `id` - (Required, String) The unique identifier for this virtual network interface.
	~> **NOTE** to add `ips` only existing `reserved_ip` is supported, new reserved_ip creation is not supported as it leads to unmanaged(dangling) reserved ips. Use `ibm_is_subnet_reserved_ip` to create a reserved_ip
	- `name` - (Required, String) The name for this virtual network interface. The name is unique across all virtual network interfaces in the VPC.
	- `resource_type` - (Computed, String) The resource type.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

- `id` - The unique identifier of the Instance NetworkAttachment.
- `created_at` - (String) The date and time that the instance network attachment was created.
- `href` - (String) The URL for this instance network attachment.
- `lifecycle_state` - (String) The lifecycle state of the instance network attachment.
  * Allowable values are: `deleting`, `failed`, `pending`, `stable`, `suspended`, `updating`, `waiting`.
- `network_attachment` - (String) The id of the network attachment.
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


## Import

You can import the `ibm_is_instance_network_attachment` resource by using `id`.
The `id` property can be formed from `instance`, and `id` in the following format:

```
<instance>/<id>
```
- `instance`: A string. The virtual server instance identifier.
- `id`: A string. The instance network attachment identifier.

# Syntax
```
$ terraform import ibm_is_instance_network_attachment.is_instance_network_attachment <instance>/<id>
```
