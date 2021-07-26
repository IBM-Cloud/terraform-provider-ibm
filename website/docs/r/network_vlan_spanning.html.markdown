---

subcategory: "Classic infrastructure"
layout: "ibm"
page_title: "IBM: network_vlan_spanning"
description: |-
  Configures VLAN Spanning for the IaaS account.
---

# ibm_network_vlan_spanning
This resource configures the VLAN spanning attribute for an IaaS account. By default VLAN spanning on the private network is disabled (off) and servers provisioned on separate private VLANs will not be able to communicate with each other over the private network. When enabled, the private network VLAN spanning service allows all private network VLANs to communicate with one another and hence all servers in the account to communicate with each other. Future servers will be added as they are provisioned. VLAN spanning enables servers to communicate across VLANs in the same data center and across data centers. 

VLAN Spanning must be enabled to use Security Groups containing servers provisioned over multiple VLANs or across multiple data centers and regions. Note VLAN Spanning does not implement network security or firewalls and must be used with Security Groups or Virtual Router Appliances (Vyatta or Juniper) to provide network isolation. 

VRF at an IaaS account level can be used as an alternative to VLAN Spanning and is required if DirectLink is used. For more information, see [VLAN spanning](https://cloud.ibm.com/docs/vlans?topic=vlans-vlan-spanning).

## Example usage

```terraform
resource "ibm_network_vlan_spanning" "spanning" {
   "vlan_spanning" = "on"
   
}`
```


## Argument reference 
Review the argument references that you can specify for your resource.

- `vlan_spanning` - (Required, String) The desired state of VLAN spanning for the account. Accepted values are **on**, **off**.


## Attribute reference
In addition to the argument reference list, you can access the following attribute reference after your resource is created.

- `id` - (String) The unique identifier of the VLAN spanning resource.
