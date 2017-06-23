---
layout: "ibm"
page_title: "IBM : firewall"
sidebar_current: "docs-ibm-resource-firewall"
description: |-
  Manages rules for IBM Firewall.
---

# ibm\_firewall

Represents firewall resources in IBM. One firewall protects one public VLAN and provides in-bound network packet filtering. 

You can order or find firewalls in the SoftLayer Customer Portal. Navigate to **Network > IP Management > VLANs** and view the  **Gateway/Firewall** column. 

For more information about how to configure a firewall, see the [docs](https://knowledgelayer.softlayer.com/procedure/configure-hardware-firewall-dedicated).

```hcl
resource "ibm_firewall" "testfw" {
  ha_enabled = false
  public_vlan_id = 12345678
}
```

## Argument Reference

The following arguments are supported:

* `ha_enabled` - (Required, boolean) Set whether the local load balancer needs to be HA enabled or not.
* `public_vlan_id` - (Required, integer) Target public VLAN ID to be protected by the firewall. Accepted values can be found [here](https://control.softlayer.com/network/vlans). Click the desired VLAN and note the ID on the resulting URL. Or, you can [refer to a VLAN by name using a data source](../d/network_vlan.html).
