---
subcategory: "Classic infrastructure"
layout: "ibm"
page_title: "IBM : ibm_network_vlan"
description: |-
  Get information on a IBM Cloud network VLAN.
---

# ibm_network_vlan
Retrieve information of an existing network VLAN as a read-only data source. For more information, about network VLAN, see [network services](https://cloud.ibm.com/docs/cloud-infrastructure?topic=cloud-infrastructure-ha-introduction#network-services).


## Example usage
The following example shows how you can use this data source to reference a VLAN ID in the `ibm_compute_bare_metal` resource because the numeric IDs are often unknown.

```terraform
data "ibm_network_vlan" "vlan_foo" {
    name = "FOO"
}

resource "ibm_compute_bare_metal" "bm1" {
    public_vlan_id = data.ibm_network_vlan.vlan_foo.id
}
```

## Argument reference
Review the argument references that you can specify for your data source.

- `name` - (Conditional, String) The name of the VLAN. This value is required if neither the VLAN number nor the router host name is specified. To retrieve the name, go to the [IBM Cloud infrastructure portal](https://cloud.ibm.com/classic/network/vlans) and navigate to **Network > IP Management > VLANs**. **Note** you need right permission to access the classic infrastructure.
- `number`- (Conditional, Integer) The VLAN number. This value is required if no VLAN name is provided. To find the number, go to the [IBM Cloud infrastructure portal](https://cloud.ibm.com/classic/network/vlans).
- `router_hostname` - (Conditional, String) The host name of the primary VLAN router. This value is required if no VLAN name is provided. To find the host name, see [IBM Cloud infrastructure portal](https://cloud.ibm.com/classic/network/vlans).


## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created.

- `id` - (String) The unique identifier of the VLAN.
- `subnets` - (List of Objects) The collection of subnets associated with the VLAN.

  Nested scheme for `subnets`:
  - `cidr`- (Integer) A subnet Classless Inter-Domain Routing prefix. The number in the range of 0 - 32 signifying the number of bits in a subnet mask.
  - `gateway` - (String) A subnet gateway address.
  - `id` - (String) The ID of the subnet.
  - `subnet` - (String) The subnet for the VLAN.
  - `subnet-type` - (String) A subnet can be one of several types. `PRIMARY, ADDITIONAL_PRIMARY, SECONDARY, ROUTED_TO_VLAN, SECONDARY_ON_VLAN, STORAGE_NETWORK, and STATIC_IP_ROUTED`. A `PRIMARY` subnet is the primary network that is bound to a VLAN within the IBM Cloud network. An `ADDITIONAL_PRIMARY` subnet is bound to a network VLAN to augment the pool of available primary IP addresses that might be assigned to a server. A `SECONDARY` subnet is any of the secondary subnets bound to a VLAN interface. A `ROUTED_TO_VLAN` subnet is a portable subnet that can be routed to any server on a VLAN. A `SECONDARY_ON_VLAN` subnet also doesn't exist as a VLAN interface, but is routed directly to a VLAN instead of a single IP address.
  - `subnet-size` - (String) The size of the subnet for the VLAN.
- `virtual_guests` - (List of Objects) A nested block describes the VSIs attached to the VLAN.

  Nested scheme for `virtual_guests`:
	- `domain` - (String) The domain of the virtual guest.
	- `id` - (String) The ID of the virtual guest.
	- `hostname` - (String) The hostname of the virtual guest.
