---

subcategory: "Classic infrastructure"
layout: "ibm"
page_title: "IBM: network_gateway_vlan_association"
description: |-
  Manages association and dis-association of VLAN to gateway.
---

# ibm_network_gateway_vlan_association
Create, update, and delete a VLAN with a network gateway. The VLANs can be disassociated or updated later to be bypassed or routed. For more information, about association of VLAN in gateway, see [viewing gateway appliance details](https://cloud.ibm.com/docs/gateway-appliance?topic=gateway-appliance-viewing-gateway-appliance-details).

**Note**

For more information, see the [IBM Cloud Classic Infrastructure (SoftLayer) API docs](http://sldn.softlayer.com/reference/services/SoftLayer_Network_Gateway_Vlan).

For more information, about getting started, see the [IBM Virtual Router Appliance docs](https://cloud.ibm.com/docs/infrastructure/virtual-router-appliance/getting-started.html#getting-started).

## Example usage

```terraform
resource "ibm_network_gateway" "gateway" {
  name = "gateway"

  members {
    hostname             = "my-virtual-router"
    domain               = "terraformuat1.ibm.com"
    datacenter           = "ams01"
    network_speed        = 100
    private_network_only = false
    tcp_monitoring       = true
    process_key_name     = "INTEL_SINGLE_XEON_1270_3_40_2"
    os_key_name          = "OS_VYATTA_5600_5_X_UP_TO_1GBPS_SUBSCRIPTION_EDITION_64_BIT"
    redundant_network    = false
    disk_key_names       = ["HARD_DRIVE_2_00TB_SATA_II"]
    public_bandwidth     = 20000
    memory               = 4
    ipv6_enabled         = true
  }
}

resource "ibm_network_gateway_vlan_association" "gateway_vlan_association" {
  gateway_id      = ibm_network_gateway.gateway.id
  network_vlan_id = 645086
}


```
## Argument reference 
Review the argument references that you can specify for your resource.

- `bypass`-  (Optional, Bool) Indicates if the VLAN should be in bypass or routed mode. Default value is **true**.
- `gateway_id` - (Required, Forces new resource, Integer) The ID of the network gateway.
- `network_vlan_id` - (Required, Forces new resource, Integer) The ID of the network VLAN to associate with the network gateway.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id`- (String) The unique identifier of the gateway/VLAN association.
