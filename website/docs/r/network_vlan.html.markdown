---
layout: "ibm"
page_title: "IBM : vlan"
sidebar_current: "docs-ibm-resource-network-vlan"
description: |-
  Manages IBM Network VLAN.
---

# ibm\_network_vlan

Provides a VLAN resource. This allows public and private VLANs to be created, updated, and cancelled.

If you have a default SoftLayer account, you do not have permission to create a VLAN using the SoftLayer API. If you want to create a VLAN with Terraform, you must get the required permissions in advance. Contact a SoftLayer sales person or open a ticket.

You can manage existing VLANs with Terraform by using the `terraform import` command. The command requires the VLAN IDs, which you can find in the [IBM Cloud infrastructure customer portal](https://control.softlayer.com/network/vlans). After the VLAN IDs are imported inot Softlayer, the IDs provide useful information such as subnets and child resource counts. When you run the `terraform destroy` command, the billing item for the VLAN is deleted. The VLAN remains in SoftLayer until you delete remaining resources on the VLAN, such as virtual guests, secondary subnets, and firewalls.

For additional details please refer to the [SoftLayer API docs](http://sldn.softlayer.com/reference/datatypes/SoftLayer_Network_Vlan).

## Example Usage

In the following example, you can create a VLAN:

```hcl
resource "ibm_network_vlan" "test_vlan" {
   name = "test_vlan"
   datacenter = "dal06"
   type = "PUBLIC"
   router_hostname = "fcr01a.dal06"
   tags = [
     "collectd",
     "mesos-master"
   ]
}

```

## Argument Reference

The following arguments are supported:

* `datacenter` - (Required, string) The data center in which the VLAN resides.
* `type` - (Required, string) The type of VLAN. Accepted values are `PRIVATE` and `PUBLIC`.
* `subnet_size` - (Removed, integer) The size of the primary subnet for the VLAN. Accepted values are `8`, `16`, `32`, and `64`. This field has been removed.
* `name` - (Optional, string) The name of the VLAN.
* `router_hostname` - (Optional, string) The hostname of the primary router associated with the VLAN.
* `tags` - (Optional, array of strings) Tags associated with the VLAN. Permitted characters include: A-Z, 0-9, whitespace, _ (underscore), - (hyphen), . (period), and : (colon). All other characters are removed.

### Timeouts

The `timeouts` block allows you to specify [timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) for certain actions:
* `delete` - (Defaults to 10 mins) Used when deleting the VLAN. There might be some resources(like Virtual Guests) on the VLAN. The VLAN delete request is issued once there are no Virtual Guests on the VLAN.

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the VLAN.
* `vlan_number` - The VLAN number as recorded within the SoftLayer network. This attribute is configured directly on SoftLayer's networking equipment.
* `softlayer_managed` - Whether or not Softlayer manages the VLAN. If Softlayer creates the VLAN automatically when Softlayer creates other resources, this attribute is set to `true`. If a user creates the VLAN using the SoftLayer API, portal, or ticket, this attribute is set to `false`.
* `child_resource_count` - A count of the resources, such as virtual servers and other network components, that are connected to the VLAN.
* `subnets` - The collection of subnets associated with the VLAN.
    * `subnet` - The subnet for the vlan.
    * `subnet-type` - A subnet can be one of several types. `PRIMARY, ADDITIONAL_PRIMARY, SECONDARY, ROUTED_TO_VLAN, SECONDARY_ON_VLAN, STORAGE_NETWORK, and STATIC_IP_ROUTED`. A `PRIMARY` subnet is the primary network bound to a VLAN within the softlayer network. An `ADDITIONAL_PRIMARY` subnet is bound to a network VLAN to augment the pool of available primary IP addresses that may be assigned to a server. A `SECONDARY` subnet is any of the secondary subnet's bound to a VLAN interface. A `ROUTED_TO_VLAN` subnet is a portable subnet that can be routed to any server on a vlan. A `SECONDARY_ON_VLAN` subnet also doesn't exist as a VLAN interface, but is routed directly to a VLAN instead of a single IP address by SoftLayer's.
    * `subnet-size` - The size of the subnet for the VLAN.
    * `gateway` - A subnet's gateway address.
    * `cidr` - A subnet's Classless Inter-Domain Routing prefix. This is a number between 0 and 32 signifying the number of bits in a subnet's netmask. 
