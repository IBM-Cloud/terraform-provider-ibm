---
layout: "ibm"
page_title: "IBM : hardware firewall shared"
sidebar_current: "docs-ibm-resource-hardware-firewall-shared"
description: |-
  Manages rules for IBM Firewall shared.
---

# ibm\_hardware\_firewall\_shared

Provides a firewall in IBM. One firewall protects one public VLAN and provides in-bound network packet filtering. 

<!-- You can order or find firewalls in the IBM Cloud infrastructure customer portal by navigating to **Network > IP Management > VLANs** and clicking the **Gateway/Firewall** column. -->

For more information about how to configure a firewall, see the [docs](https://knowledgelayer.softlayer.com/procedure/configure-hardware-firewall).

## Example Usage

```hcl
resource "ibm_hardware-firewall_shared" "test_firewall" {
    firewall_type="100MBPS_HARDWARE_FIREWALL"
    hardware_instance_id="12345678"
}
```

## Argument Reference

The following arguments are supported:

* `firewall_type` - (Required, string) Specifies whether it needs to be of particular speed. Firewall type is in between [10MBPS_HARDWARE_FIREWALL, 20MBPS_HARDWARE_FIREWALL,100MBPS_HARDWARE_FIREWALL, 1000MBPS_HARDWARE_FIREWALL]
* `virtual_instance_id` - (Optional, string) Specifies the id of particular guest on which firewall shared is to be deployed.**NOTE**: This is conflicting parameter with hardware_instance_id.
* `hardware_instance_id` - (Optional, string) Specifies the id of particular guest on which firewall shared is to be deployed.**NOTE**: This is conflicting parameter with virtual_instance_id.

## Attribute Reference

The following attributes are exported:
 * `id` - The unique identifier of the hardware firewall.
