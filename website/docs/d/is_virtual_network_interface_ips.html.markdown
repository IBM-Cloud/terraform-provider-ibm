---
layout: "ibm"
page_title: "IBM : ibm_is_virtual_network_interface_ips"
description: |-
  Get information about ReservedIP Collection of a Virtual Network Interface
subcategory: "VPC infrastructure"
---

# ibm_is_virtual_network_interface_ips

Provides a read-only data source to retrieve information about a ReservedIP Collection bound to a virtual network interface. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```terraform
data "ibm_is_virtual_network_interface_ips" "is_reserved_ips" {
	virtual_network_interface = ibm_is_virtual_network_interface.testacc_vni.id
}
```

## Argument Reference

You can specify the following arguments for this data source.

- `virtual_network_interface` - (Required, Forces new resource, String) The virtual network interface identifier.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

- `id` - The unique identifier of the virtual network interface reserved IP Collection.

- `reserved_ips` - (List) Collection of reserved IPs in this subnet.
	Nested schema for **reserved_ips**:
	- `address` - (String) The IP address.If the address has not yet been selected, the value will be `0.0.0.0`.This property may add support for IPv6 addresses in the future. When processing a value in this property, verify that the address is in an expected format. If it is not, log an error. Optionally halt processing and surface the error, or bypass the resource on which the unexpected IP address format was encountered.
	- `href` - (String) The URL for this reserved IP.
	- `deleted` - (List) 	If present, this property indicates the referenced resource has been deleted and provides some supplementary information.
		Nested scheme for **deleted**:
		- `more_info` - (String) Link to documentation about deleted resources.
	- `id` - (String) The unique identifier for this reserved IP.
	- `name` - (String) The name for this reserved IP. The name is unique across all reserved IPs in a subnet.
	- `resource_type` - (String) The resource type.

