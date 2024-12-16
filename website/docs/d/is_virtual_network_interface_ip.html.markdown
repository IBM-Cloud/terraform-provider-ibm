---
layout: "ibm"
page_title: "IBM : ibm_is_virtual_network_interface_ip"
description: |-
  Get information about Virtual Network Interface Reserved Ip
subcategory: "VPC infrastructure"
---

# ibm_is_virtual_network_interface_ip

Provides a read-only data source to retrieve information about a Virtual Network Interface Reserved Ip. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_is_virtual_network_interface_ip" "is_reserved_ip" {
	reserved_ip = "id"
	virtual_network_interface = ibm_is_virtual_network_interface_ip.is_reserved_ip.subnet_id
}
```

## Argument Reference

You can specify the following arguments for this data source.

- `reserved_ip` - (Required, Forces new resource, String) The reserved IP identifier.
- `virtual_network_interface` - (Required, Forces new resource, String)  The virtual network interface identifier.

## Attribute Reference

After your data source is created, you can read values from the following attributes.

- `address` - (String) The IP address.If the address has not yet been selected, the value will be `0.0.0.0`.This property may add support for IPv6 addresses in the future. When processing a value in this property, verify that the address is in an expected format. If it is not, log an error. Optionally halt processing and surface the error, or bypass the resource on which the unexpected IP address format was encountered.
- `href` - (String) The URL for this reserved IP.
- `id` - (String) The unique identifier for this reserved IP.
- `deleted` - (List) 	If present, this property indicates the referenced resource has been deleted and provides some supplementary information.
	Nested scheme for **deleted**:
	- `more_info` - (String) Link to documentation about deleted resources.
- `name` - (String) The name for this reserved IP. The name is unique across all reserved IPs in a subnet.
- `resource_type` - (String) The resource type.
