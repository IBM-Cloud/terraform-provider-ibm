---
layout: "ibm"
page_title: "IBM : ibm_is_virtual_network_interface_floating_ip"
description: |-
  Manages Virtual Network Interface Floating IP.
subcategory: "VPC infrastructure"
---

# ibm_is_virtual_network_interface_floating_ip

Create, read, and delete Virtual Network Interface Floating IP with this resource.

## Example Usage

```terrform
resource "ibm_is_virtual_network_interface_floating_ip" "vni_fip" {
  virtual_network_interface   =   ibm_is_virtual_network_interface.example.id
  floating_ip                 =   ibm_is_floating_ip.example.id
}
```

## Argument Reference

You can specify the following arguments for this resource.

- `virtual_network_interface` - (Required, String) The virtual network interface identifier
- `floating_ip` - (Required, String) The floating IP identifier
## Attribute Reference

After your resource is created, you can read values from the listed arguments and the following attributes.

- `id` - The unique identifier of the VirtualNetworkInterfaceFloatingIP. The ID is composed of `<vni_id>/<floating_ip_id>`.
- `address` - (String) The globally unique IP address.
- `crn` - (String) The CRN for this floating IP.
- `href` - (String) The URL for this floating IP.
- `deleted` - (List) 	If present, this property indicates the referenced resource has been deleted and provides some supplementary information.
	Nested scheme for **deleted**:
	- `more_info` - (String) Link to documentation about deleted resources.
- `name` - (String) The name for this floating IP. The name is unique across all floating IPs in the region.

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import the `ibm_is_virtual_network_interface_floating_ip` resource by using `id`.
The `id` property can be formed using the virtual_network_interface and floating_ip identifier(s). For example:

```terraform
import {
  to = ibm_is_virtual_network_interface_floating_ip.vni_fip
  id = "< vni_id >/< floating_ip_id >"
}
```

Using `terraform import`. For example:

```console
% terraform import ibm_is_virtual_network_interface_floating_ip.vni_fip < vni_id >/< floating_ip_id >
```