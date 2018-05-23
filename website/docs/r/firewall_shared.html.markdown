---
layout: "ibm"
page_title: "IBM : firewall shared"
sidebar_current: "docs-ibm-resource-firewall-shared"
description: |-
  Manages rules for IBM Firewall shared.
---

# ibm\_firewall\_shared

Provides a firewall in IBM. One firewall protects one public VLAN and provides in-bound network packet filtering. 

<!-- You can order or find firewalls in the IBM Cloud infrastructure customer portal by navigating to **Network > IP Management > VLANs** and clicking the **Gateway/Firewall** column. -->

For more information about how to configure a firewall, see the [docs](https://knowledgelayer.softlayer.com/procedure/configure-hardware-firewall).

## Example Usage

```hcl
resource "ibm_firewall_shared" "test_firewall" {
    firewall_type="100MBPS_HARDWARE_FIREWALL"
    guest_type="baremetal"
    guest_id="12345678"
}
```

## Argument Reference

The following arguments are supported:

* `firewall_type` - (Required, string) Specifies whether it needs to be of particular speed. Firewall type is in between [10MBPS_HARDWARE_FIREWALL, 20MBPS_HARDWARE_FIREWALL,100MBPS_HARDWARE_FIREWALL, 1024MBPS_HARDWARE_FIREWALL]
* `guest_type` - (Required, string) Specifies whether the guest is baremetal server or virtual guest server.
* `guest_id` - (Required, string) Specifies the id of particular guest on which firewall shared is to be deployed.