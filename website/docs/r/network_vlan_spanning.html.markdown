---
layout: "ibm"
page_title: "IBM: network_vlan_spanning"
sidebar_current: "docs-ibm-resource-network-vlan-spanning"
description: |-
  Configures VLAN Spanning for the IaaS account.
---

# ibm\_network_vlan_spanning

This resource configures the VLAN spanning attribute for an IaaS account. By default VLAN spanning on the private networt is disabled (off) and servers provisioned on separate private VLANs will not be able to communicate with each other over the private network. When enabled, the private network VLAN spanning service allows all private network VLANs to communicate with one another and hence all servers in the account to communicate with each other. Future servers will be added as they are provisioned. VLAN spanning enables servers to communicate across VLANs in the same data center and across data centers. 

VLAN Spanning must be enabled to use Security Groups containing servers provisioned over multiple VLANs or across multiple data centers and regions. Note VLAN Spanning does not implement network security or firewalls and must be used with Security Groups or Virtual Router Appliances (Vyatta or Juniper) to provide network isolation. 

VRF at an IaaS account level can be used as an alternative to VLAN Spanning and is required if DirectLink is used.  



## Example Usage

```hcl
resource "ibm_network_vlan_spanning" "spanning" {
   "vlan_spanning" = "on"
}`
```


## Argument Reference

The following arguments are supported:

* `vlan_spanning` - (Required, string) The desired state of VLAN spanning for the account. Accepted values are `on`, `off`.


## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the VLAN spanning resource.
