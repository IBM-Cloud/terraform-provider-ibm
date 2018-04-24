---
layout: "ibm"
page_title: "IBM : ibm_network_vlan"
sidebar_current: "docs-ibm-datasource-network-vlan"
description: |-
  Get information on a IBM Network VLAN.
---

# ibm\_network_vlan


Import the details of an existing VLAN as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.


## Example Usage

```hcl
data "ibm_network_vlan" "vlan_foo" {
    name = "FOO"
}
```


The following example shows how you can use this data source to reference a VLAN ID in the _ibm_compute_bare_metal_ resource because the numeric IDs are often unknown.

```hcl
resource "ibm_compute_bare_metal" "bm1" {
    ...
    public_vlan_id = "${data.ibm_network_vlan.vlan_foo.id}"
    ...
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required if neither the number nor the router hostname are provided, string) The name of the VLAN, as it was defined in IBM Cloud Infrastructure (SoftLayer). You can find names in the [IBM Cloud infrastructure customer portal](https://control.softlayer.com/network/vlans) by navigating to **Network > IP Management > VLANs**.
* `number` - (Required if the name is not provided, integer) The VLAN number. You can find the numbers in the [IBM Cloud infrastructure customer portal](https://control.softlayer.com/network/vlans).
* `router_hostname` - (Required if the name is not provided, string) The primary VLAN router hostname. You can find the  hostnames in the [IBM Cloud infrastructure customer portal](https://control.softlayer.com/network/vlans).

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the VLAN.
* `subnets` - The collection of subnets associated with the VLAN.
    * `id` - The ID of the subnet.  
    * `subnet` - The subnet for the vlan.
    * `subnet-type` - A subnet can be one of several types. `PRIMARY, ADDITIONAL_PRIMARY, SECONDARY, ROUTED_TO_VLAN, SECONDARY_ON_VLAN, STORAGE_NETWORK, and STATIC_IP_ROUTED`. A `PRIMARY` subnet is the primary network bound to a VLAN within the softlayer network. An `ADDITIONAL_PRIMARY` subnet is bound to a network VLAN to augment the pool of available primary IP addresses that may be assigned to a server. A `SECONDARY` subnet is any of the secondary subnet's bound to a VLAN interface. A `ROUTED_TO_VLAN` subnet is a portable subnet that can be routed to any server on a vlan. A `SECONDARY_ON_VLAN` subnet also doesn't exist as a VLAN interface, but is routed directly to a VLAN instead of a single IP address by SoftLayer's.
    * `subnet-size` - The size of the subnet for the VLAN.
    * `gateway` - A subnet's gateway address.
    * `cidr` - A subnet's Classless Inter-Domain Routing prefix. This is a number between 0 and 32 signifying the number of bits in a subnet's netmask. 
* `virtual_guests` - A nested block describing the VSIs attached to the VLAN. Nested `virtual_guests` blocks have the following structure:
  * `id` - The ID of the virtual guest.
  * `domain` - The domain of the virtual guest.
  * `hostname` - The hostname of the virtual guest.

