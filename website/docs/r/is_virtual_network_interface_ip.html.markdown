---
layout: "ibm"
page_title: "IBM : ibm_is_virtual_network_interface_ip"
description: |-
  Manages virtual_network_interface reserved ip attachment.
subcategory: "VPC infrastructure"
---

# ibm_is_virtual_network_interface_ip

Create, update, and delete ReservedIP virtual network instance attachment with this resource.

## Example Usage

```hcl
resource "ibm_is_virtual_network_interface_ip" "is_reserved_ip_instance" {
  reserved_ip               = ibm_is_subnet_reserved_ip.example.reserved_ip
  virtual_network_interface = ibm_is_virtual_network_interface.example.id
}
```

## Argument Reference

You can specify the following arguments for this resource.

- `reserved_ip` - (Required, Forces new resource, String) The reserved IP identifier.
- `virtual_network_interface` - (Required, Forces new resource, String)  The virtual network interface identifier.

## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

- `address` - (String) The IP address.If the address has not yet been selected, the value will be `0.0.0.0`.This property may add support for IPv6 addresses in the future. When processing a value in this property, verify that the address is in an expected format. If it is not, log an error. Optionally halt processing and surface the error, or bypass the resource on which the unexpected IP address format was encountered.
- `deleted` - (List) 	If present, this property indicates the referenced resource has been deleted and provides some supplementary information.
	Nested scheme for **deleted**:
	- `more_info` - (String) Link to documentation about deleted resources.
- `href` - (String) The URL for this reserved IP.
- `id` - (String) The unique identifier for this reserved IP.
- `name` - (String) The name for this reserved IP. The name is unique across all reserved IPs in a subnet.
- `resource_type` - (String) The resource type.


## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import the `ibm_is_virtual_network_interface_ip` resource by using `id`.
The `id` property can be formed from `virtual_network_interface`, and `reserved_ip`. For example:

```terraform
import {
  to = ibm_is_virtual_network_interface_ip.is_reserved_ip
  id = "<virtual_network_interface>/<reserved_ip>"
}
```

Using `terraform import`. For example:

```console
% terraform import ibm_is_virtual_network_interface_ip.is_reserved_ip <virtual_network_interface>/<reserved_ip>
```