---
layout: "ibm"
page_title: "IBM : firewall"
sidebar_current: "docs-ibm-resource-firewall"
description: |-
  Manages rules for IBM Firewall.
---

# ibm\_firewall

Provides a firewall in IBM. One firewall protects one public VLAN and provides in-bound network packet filtering.

You can order or find firewalls in the IBM Cloud infrastructure customer portal by navigating to **Network > IP Management > VLANs** and clicking the **Gateway/Firewall** column.

For more information about how to configure a firewall, see the [docs](https://knowledgelayer.softlayer.com/procedure/configure-hardware-firewall-dedicated).

## Example Usage

```hcl
resource "ibm_firewall" "testfw" {
  ha_enabled = false
  public_vlan_id = 12345678
  tags = [
     "collectd",
     "mesos-master"
   ]
}
```

## Argument Reference

The following arguments are supported:

* `ha_enabled` - (Required, boolean) Specifies whether the local load balancer needs to be HA-enabled.
* `public_vlan_id` - (Required, integer) The target public VLAN ID that you want the firewall to protect. You can find accepted values [here](https://control.softlayer.com/network/vlans). Click the desired VLAN and note the ID number in the resulting URL. You can also [refer to a VLAN by name using a data source](../d/network_vlan.html).
* `tags` - (Optional, array of strings) Set tages on the firewall. Permitted characters include: A-Z, 0-9, whitespace, _ (underscore), - (hyphen), . (period), and : (colon). All other characters are removed.
