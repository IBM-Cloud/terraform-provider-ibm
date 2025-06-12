---
layout: "ibm"
page_title: "IBM : ibm_is_virtual_network_interface_floating_ips"
description: |-
  List information about Virtual Network Interface Floating IPs.
subcategory: "VPC infrastructure"
---

# ibm_is_virtual_network_interface_floating_ips

Provides a read-only data source to retrieve information about a list of Virtual Network Interface Floating IPs. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```terraform
data "ibm_is_virtual_network_interface_floating_ips" "vni_fips" {
  virtual_network_interface = <vni_id>
}
```

## Argument Reference

You can specify the following arguments for this data source.

- `virtual_network_interface` - (Required, String) The virtual network interface identifier

## Attribute Reference

After your data source is created, you can read values from the following attributes.
- `floating_ips` - (List)  Lists all floatings ips in the virtual network interface.

  Nested scheme for `floating_ips`:
	- `id` - The unique identifier of the FloatingIP
	- `address` - (String) The globally unique IP address.
	- `crn` - (String) The CRN for this floating IP.
	- `href` - (String) The URL for this floating IP.
	- `deleted` - (List) 	If present, this property indicates the referenced resource has been deleted and provides some supplementary information.
		Nested scheme for **deleted**:
		- `more_info` - (String) Link to documentation about deleted resources.
	- `name` - (String) The name for this floating IP. The name is unique across all floating IPs in the region.

	
