---

subcategory: "Classic infrastructure"
layout: "ibm"
page_title: "IBM : firewall"
description: |-
  Manages rules for IBM rirewall.
---

# ibm_firewall
Provides a firewall in IBM. One firewall protects one public VLAN and provides in-bound network packet filtering. You can order or find firewalls in the IBM Cloud infrastructure customer portal by navigating to **Network > IP Management > VLANs** and clicking the **Gateway/Firewall** column. For an overview of supported firewalls, see [exploring firewalls](https://cloud.ibm.com/docs/hardware-firewall-dedicated?topic=fortigate-10g-exploring-firewalls).

## Example usage

```terraform
resource "ibm_firewall" "testfw" {
  firewall_type  = "HARDWARE_FIREWALL_DEDICATED"
  ha_enabled     = false
  public_vlan_id = 12345678
  tags = [
    "collectd",
    "mesos-master",
  ]
}

```

## Argument reference
Review the argument references that you can specify for your resource. 

- `firewall_type`- (Optional, Forces new resource, String) Specifies the type of firewall to create. Valid options are HARDWARE_FIREWALL_DEDICATED or FORTIGATE_SECURITY_APPLIANCE. Defaults to HARDWARE_FIREWALL_DEDICATED-
- `ha_enabled`- (Required, Forces new resource, Bool) Specifies whether the local load balancer needs to be HA-enabled.
- `public_vlan_id` - (Required, Forces new resource, Integer)  - The target public VLAN ID that you want the firewall to protect. You can find accepted values [here](https://cloud.ibm.com/classic/network/vlans). Click the VLAN that you want to use and note the ID number in the resulting URL. You can also refer to a VLAN name by using a data source.
- `tags`- (Optiona, Array of Strings) Set tags on the firewall. Permitted characters include: A-Z, 0-9, whitespace, `_` (underscore), `- ` (hyphen), `.` (period), and `:` (colon). All other characters are removed.


## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id` - (String) The unique identifier of the VLAN.
- `location` - (String) The location/datacenter of the created firewall device.
- `primary_ip:` - (String) The public gateway IP address.
- `username:` - (String) The username that is used to log in into the device, in case of Forti Gate Appliances.
- `password:` - (String) The password that is used to log in into the device, in case of Forti Gate Appliances.

