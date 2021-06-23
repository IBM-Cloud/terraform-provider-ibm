---

subcategory: "Classic infrastructure"
layout: "ibm"
page_title: "IBM : multivlan_firewall"
description: |-
  Manages IBM Cloud multi-Vlan firewall.
---

# ibm_multivlan_firewall
Create, delete, and update a multi-Vlan firewall resource. For more information, about IBM Cloud multi-Vlan firewall, see [getting started with VLANs](https://cloud.ibm.com/docs/vlans?topic=vlans-getting-started).

For more information, see the [IBM Cloud (SoftLayer) multi VLAN firewall Request Docs](https://sldn.softlayer.com/reference/datatypes/SoftLayer_Container_Product_Order_Network_Protection_Firewall_Dedicated/).

## Example usage

In the following example, you can create a multi-vlan firewall:

```terraform
resource "ibm_multi_vlan_firewall" "firewall_first" {
  datacenter          = "dal13"
  pod                 = "pod01"
  name                = "Checkdelete1"
  firewall_type       = "FortiGate Security Appliance"
  addon_configuration = ["FortiGate Security Appliance - Web Filtering Add-on (High Availability)", "FortiGate Security Appliance - NGFW Add-on (High Availability)", "FortiGate Security Appliance - AV Add-on (High Availability)"]
}

```

## Argument reference 
Review the argument references that you can specify for your resource.

- `addon_configuration`- (Required, List) The list of add-ons that are allowed. Allowed values are **FortiGate Security Appliance) Web Filtering Add-on (High Availability)**,**FortiGate Security Appliance - NGFW Add-on (High Availability)**,**FortiGate Security Appliance - AV Add-on (High Availability)** or **FortiGate Security Appliance - Web Filtering Add-on**, **FortiGate Security Appliance - NGFW Add-on**,**FortiGate Security Appliance - AV Add-on**.
- `datacenter` - (Required, Forces new resource, String)The data center in which the firewall appliance resides.
- `firewall_type` - (Required, Forces new resource, String)The type of the firewall device. Supported values include **FortiGate Security Appliance** and **FortiGate Firewall Appliance HA Option**.
- `name` - (Required, Forces new resource, String)The name of the firewall device.
- `pod` - (Required, Forces new resource, String)The pod in which the firewall resides.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id` - (String) The unique identifier of the multi VLAN firewall.
- `public_vlan_id` - (String) The ID of the Public VLAN for accessing this gateway.
- `private_vlan_id` - (String) The ID of the Private VLAN for accessing this gateway.
- `public_ip` - (String) The public gateway IP address.
- `public_ipv6` - (String) The public gateway IPv6 address.
- `private_ip` - (String) The private gateway IP address.
- `password` - (String) The password that is used to log in into the device.
- `username` - (String) The username that is used to log in into the device.
