---
layout: "ibm"
page_title: "IBM : firewall"
sidebar_current: "docs-ibm-resource-firewall"
description: |-
  Manages rules for IBM Firewall.
---

# ibm\_firewall

Provides a firewall in IBM. One firewall protects one public VLAN and provides in-bound network packet filtering. Firewall type can be choosen in between the Dedicated Firewall, and Fortigate Security Appliance based on their specific performance and feature requirements.

You can order or find firewalls in the IBM Cloud infrastructure customer portal by navigating to **Network > IP Management > VLANs** and clicking the **Gateway/Firewall** column.

For more information about how to configure a firewall, see the [docs](https://knowledgelayer.softlayer.com/procedure/configure-hardware-firewall-dedicated).

## Example Usage

```hcl
resource "ibm_firewall" "testfw" {
  firewall_type = "HARDWARE_FIREWALL_DEDICATED"
  public_vlan_id = 12345678
  tags = [
     "collectd",
     "mesos-master"
   ]
}
```

## Argument Reference

The following arguments are supported:

* `firewall_type` - (Required, string) Specifies whether it needs to be HA-enabled. Firewall type is in between [HARDWARE_FIREWALL_DEDICATED, HARDWARE_FIREWALL_HIGH_AVAILABILITY,
FORTIGATE_SECURITY_APPLIANCE, FORTIGATE_SECURITY_APPLIANCE_HIGH_AVAILABILITY]
* `public_vlan_id` - (Required, integer) The target public VLAN ID that you want the firewall to protect. You can find accepted values [here](https://control.softlayer.com/network/vlans). Click the desired VLAN and note the ID number in the resulting URL. You can also [refer to a VLAN by name using a data source](../d/network_vlan.html).
* `tags` - (Optional, array of strings) Set tages on the firewall. Permitted characters include: A-Z, 0-9, whitespace, _ (underscore), - (hyphen), . (period), and : (colon). All other characters are removed.
