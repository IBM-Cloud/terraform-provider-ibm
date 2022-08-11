---

subcategory: "Classic infrastructure"
layout: "ibm"
page_title: "IBM : hardware firewall shared"
description: |-
  Manages rules for IBM firewall shared.
---

# ibm_hardware_firewall_shared
Create, delete, and update a shared hardware firewall. One firewall protects one public VLAN and provides in-bound network packet filtering. You can order or find firewalls in the IBM Cloud infrastructure customer portal by navigating to **Network > IP Management > VLANs** and clicking the **Gateway/Firewall** column.

For more information, about how to configure a firewall, see [about hardware firewalls](https://cloud.ibm.com/docs/hardware-firewall-shared?topic=hardware-firewall-shared-about-hardware-firewall-shared-).

## Example usage

```terraform
resource "ibm_hardware_firewall_shared" "test_firewall" {
    firewall_type="100MBPS_HARDWARE_FIREWALL"
    hardware_instance_id="12345678"
    
}
```

## Timeouts

The `ibm_hardware_firewall_shared` resource provides the following [Timeouts](https://www.terraform.io/docs/language/resources/syntax.html) configuration options:

- **create** - (Default 10 minutes) Used for creating shared firewall.

## Argument reference
Review the argument references that you can specify for your resource. 

- `firewall_type`- (Required, String) Specifies whether it needs to be of particular speed. Firewall type is in between <br> [10MBPS_HARDWARE_FIREWALL, <br>20MBPS_HARDWARE_FIREWALL, <br>100MBPS_HARDWARE_FIREWALL, <br>1000MBPS_HARDWARE_FIREWALL].
- `virtual_instance_id`- (Optional, String) Specifies the ID of particular guest on which firewall shared is to be deployed. **Note** This reference conflicts with `hardware_instance_id`.
- `hardware_instance_id`- (Optional, String) Specifies the ID of particular guest on which firewall shared is to be deployed. **Note** This reference conflicts with `virtual_instance_id`.


## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

 - `id`- (String) The unique identifier of the hardware firewall.
