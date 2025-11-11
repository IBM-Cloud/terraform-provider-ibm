---
layout: "ibm"
page_title: "IBM : ibm_is_instance_network_attachment"
description: |-
  Manages Instance Network Attachment.
subcategory: "VPC infrastructure"
---

# ibm_is_instance_network_attachment

Create, update, and delete Instance Network Attachment with this resource.
Instance network attachments allow you to create and manage additional virtual network interfaces to your compute instances using virtual service instance network attachment.

## Example Usage

### Basic usage with existing virtual network interface

```terraform
resource "ibm_is_instance_network_attachment" "example" {
  instance = ibm_is_instance.example.id
  name     = "example-networkatt"
  virtual_network_interface {
    id = ibm_is_virtual_network_interface.example.id
  }
}
```

### Inline virtual network interface with auto-assigned IP

```terraform
resource "ibm_is_instance_network_attachment" "example2" {
  instance = ibm_is_instance.example.id
  name     = "example-networkatt2"
  virtual_network_interface {
    name                     = "example-vni-2"
    subnet                   = ibm_is_subnet.example.id
    auto_delete              = true
    allow_ip_spoofing        = false
    enable_infrastructure_nat = true
    security_groups = [
      ibm_is_security_group.example.id
    ]
  }
}
```

### Inline virtual network interface with specific primary IP address

```terraform
resource "ibm_is_instance_network_attachment" "example3" {
  instance = ibm_is_instance.example.id
  name     = "example-networkatt3"
  virtual_network_interface {
    name                     	= "example-vni-3"
    subnet                   	= ibm_is_subnet.example.id
    auto_delete              	= true
    allow_ip_spoofing        	= false
    enable_infrastructure_nat 	= true
    primary_ip {
      address     = "10.240.64.100"
      auto_delete = true
      name        = "example-primary-ip-1"
    }
  }
}
```

### Advanced configuration with reserved IPs

```terraform
resource "ibm_is_subnet_reserved_ip" "example_primary" {
  subnet  = ibm_is_subnet.example.id
  name    = "example-primary-ip"
  address = "10.240.65.50"
}

resource "ibm_is_subnet_reserved_ip" "example_additional" {
  subnet  = ibm_is_subnet.example.id
  name    = "example-additional-ip"
  address = "10.240.65.51"
}

resource "ibm_is_instance_network_attachment" "example4" {
  instance = ibm_is_instance.example.id
  name     = "example-networkatt4"
  virtual_network_interface {
    name                     = "example-vni-5"
    subnet                   = ibm_is_subnet.example.id
    auto_delete              = true
    allow_ip_spoofing        = false
    enable_infrastructure_nat = true
    primary_ip {
      reserved_ip = ibm_is_subnet_reserved_ip.example_primary.reserved_ip
    }
    ips {
      reserved_ip = ibm_is_subnet_reserved_ip.example_additional.reserved_ip
    }
  }
}
```

## Argument Reference

You can specify the following arguments for this resource.

- `instance` - (Required, Forces new resource, String) The virtual server instance identifier.
- `name` - (Optional, String) The name for this instance network attachment. The name is unique across all network attachments for the instance.
- `virtual_network_interface` - (Required, List) The virtual network interface for this instance network attachment. This can be specified using an existing virtual network interface ID, or a prototype object for a new virtual network interface.

	Nested schema for **virtual_network_interface**:
  - `crn` - (String) The CRN for this virtual network interface.
	- `id` - (Optional, String) The unique identifier for an existing virtual network interface. When specified, all other nested arguments are ignored and will conflict if provided.
	- `allow_ip_spoofing` - (Optional, Boolean) Indicates whether source IP spoofing is allowed on this interface. If `false`, source IP spoofing is prevented on this interface. If `true`, source IP spoofing is allowed on this interface. Conflicts with `id`.
	- `auto_delete` - (Optional, Boolean) Indicates whether this virtual network interface will be automatically deleted when `target` is deleted. Conflicts with `id`.
	- `enable_infrastructure_nat` - (Optional, Boolean) If `true`: The VPC infrastructure performs any needed NAT operations. `floating_ips` must not have more than one floating IP. If `false`: Packets are passed unchanged to/from the network interface, allowing the workload to perform any needed NAT operations. `allow_ip_spoofing` must be `false`. If the virtual network interface is attached: The target `resource_type` must be `bare_metal_server_network_attachment`. The target `interface_type` must not be `hipersocket`. Conflicts with `id`.
	- `ips` - (Optional, Set) The reserved IPs bound to this virtual network interface. May be empty when `lifecycle_state` is `pending`. Conflicts with `id`.
		~> **NOTE** to add `ips` only existing `reserved_ip` is supported, new reserved_ip creation is not supported as it leads to unmanaged(dangling) reserved ips. Use `ibm_is_subnet_reserved_ip` to create a reserved_ip

		Nested schema for **ips**:
		- `reserved_ip` - (Required, String) The unique identifier for this reserved IP.
		- `address` - (Computed, String) The IP address. If the address has not yet been selected, the value will be `0.0.0.0`. This property may add support for IPv6 addresses in the future. When processing a value in this property, verify that the address is in an expected format. If it is not, log an error. Optionally halt processing and surface the error, or bypass the resource on which the unexpected IP address format was encountered.
		- `auto_delete` - (Computed, Boolean) Indicates whether this reserved IP member will be automatically deleted when either target is deleted, or the reserved IP is unbound.
		- `deleted` - (Computed, List) If present, this property indicates the referenced resource has been deleted, and provides some supplementary information.
			Nested schema for **deleted**:
			- `more_info` - (String) Link to documentation about deleted resources.
		- `href` - (Computed, String) The URL for this reserved IP.
		- `name` - (Computed, String) The name for this reserved IP. The name is unique across all reserved IPs in a subnet.
		- `resource_type` - (Computed, String) The resource type.
	- `name` - (Optional, String) The name for this virtual network interface. The name is unique across all virtual network interfaces in the VPC. Conflicts with `id`.
	- `primary_ip` - (Optional, List) The primary IP address of the virtual network interface for the instance network attachment. Conflicts with `id`.
		Nested schema for **primary_ip**:
		- `address` - (Optional, String) The IP address. If the address has not yet been selected, the value will be `0.0.0.0`. This property may add support for IPv6 addresses in the future. When processing a value in this property, verify that the address is in an expected format. If it is not, log an error. Optionally halt processing and surface the error, or bypass the resource on which the unexpected IP address format was encountered.
		- `auto_delete` - (Optional, Boolean) Indicates whether this primary_ip will be automatically deleted when `vni` is deleted. Default value: `true`.
		- `deleted` - (Computed, List) If present, this property indicates the referenced resource has been deleted, and provides some supplementary information.
			Nested schema for **deleted**:
			- `more_info` - (String) Link to documentation about deleted resources.
		- `href` - (Computed, String) The URL for this reserved IP.
		- `name` - (Optional, String) The name for this reserved IP. The name is unique across all reserved IPs in a subnet.
		- `reserved_ip` - (Optional, String) The unique identifier for this reserved IP.
		- `resource_type` - (Computed, String) The resource type.
	- `resource_group` - (Optional, String) The resource group id for this virtual network interface. Conflicts with `id`.
	- `resource_type` - (Computed, String) The resource type.
	- `security_groups` - (Optional, Forces new resource, Set of Strings) The security groups for this virtual network interface. Conflicts with `id`.
	- `subnet` - (Optional, Forces new resource, String) The associated subnet id. Conflicts with `id`.
	- `crn` - (Computed, String) The CRN for this virtual network interface.
	- `href` - (Computed, String) The URL for this virtual network interface.

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

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import the `ibm_is_instance_network_attachment` resource by using `id`.
The `id` property can be formed from `instance`, and `id`. For example:

```terraform
import {
  to = ibm_is_instance_network_attachment.is_instance_network_attachment
  id = "<instance>/<id>"
}
```

Using `terraform import`. For example:

```console
% terraform import ibm_is_instance_network_attachment.is_instance_network_attachment <instance>/<id>
```