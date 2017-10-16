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

* `name` - (Required if neither the number nor router hostname are provided, string) The name of the VLAN, as it was defined in Bluemix Infrastructure (SoftLayer). You can find names in the [SoftLayer Customer Portal](https://control.softlayer.com/network/vlans) by navigating to **Network > IP Management > VLANs**.
* `number` - (Required if the name is not provided, integer) The VLAN number. You can find  numbers in the [SoftLayer Customer Portal](https://control.softlayer.com/network/vlans).
* `router_hostname` - (Required if the name is not provided, string) The primary VLAN router hostname. You can find hostnames in the [SoftLayer Customer Portal](https://control.softlayer.com/network/vlans).

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the VLAN.
* `subnets` - A list of subnets associated with this VLAN.
* `virtual_guests` - A nested block describing the VSIs attached to the VLAN. Nested `virtual_guests` blocks have the following structure:
  * `id` - The ID of the virtual guest.
  * `domain` - The domain of the virtual guest.
  * `hostname` - The hostname of the virtual guest.

