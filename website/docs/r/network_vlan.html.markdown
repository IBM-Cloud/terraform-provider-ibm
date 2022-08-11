---

subcategory: "Classic infrastructure"
layout: "ibm"
page_title: "IBM : vlan"
description: |-
  Manages IBM Cloud network VLAN.
---

# ibm_network_vlan
Create, delete, and update an IBM Cloud network VLAN. For more information, about network VLAN, see [getting started with VLANs](https://cloud.ibm.com/docs/vlans?topic=vlans-getting-started).

If you have a default SoftLayer account, you do not have permission to create a VLAN using the SoftLayer API. If you want to create a VLAN with Terraform, you must get the required permissions in advance. Contact a SoftLayer sales person or open a ticket.

You can manage existing VLANs with Terraform by using the `terraform import` command. The command requires the VLAN IDs, which you can find in the [IBM Cloud infrastructure customer portal](https://cloud.ibm.com/classic/network/vlans). After the VLAN IDs are imported inot Softlayer, the IDs provide useful information such as subnets and child resource counts. When you run the `terraform destroy` command, the billing item for the VLAN is deleted. The VLAN remains in SoftLayer until you delete remaining resources on the VLAN, such as virtual guests, secondary subnets, and firewalls.

For more information, see [SoftLayer API docs](http://sldn.softlayer.com/reference/datatypes/SoftLayer_Network_Vlan).

## Example usage
In the following example, you can create a VLAN:

```terraform
resource "ibm_network_vlan" "test_vlan" {
  name            = "test_vlan"
  datacenter      = "dal06"
  type            = "PUBLIC"
  router_hostname = "fcr01a.dal06"
  tags = [
    "collectd",
    "mesos-master",
  ]
}

```

## Timeouts
The `ibm_network_vlan` block allows you to specify [Timeouts](https://www.terraform.io/docs/language/resources/syntax.html) for certain actions:

- **create**: (Defaults to 10 mins) Used when creating the VLAN.
- **delete**: (Defaults to 10 mins) Used when deleting the VLAN. There might be some resources (like Virtual Guests) on the VLAN. The VLAN delete request is issued once there are no Virtual Guests on the VLAN.

## Argument reference 
Review the argument references that you can specify for your resource.

- `datacenter` - (Required, Forces new resource, String) The data center in which the VLAN resides.
- `type` - (Required, Forces new resource, String)The type of VLAN. Accepted values are `PRIVATE` and `PUBLIC`.
- `name` - (Optional, String) The name of the VLAN. Maximum length of 20 characters.
- `router_hostname` - (Optional, Forces new resource, String) The hostname of the primary router associated with the VLAN.
- `tags` - (Optional, Array of string) Tags associated with the VLAN. Permitted characters include: A-Z, 0-9, whitespace, `_` (underscore), `- ` (hyphen), `.` (period), and `:` (colon). All other characters are removed.


## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `child_resource_count` - (String) A count of the resources, such as virtual servers and other network components, that are connected to the VLAN.
- `id` - (String) The unique identifier of the VLAN.
- `softlayer_managed` - (String) SoftLayer manages the VLAN. If SoftLayer creates the VLAN automatically when SoftLayer creates other resources, this attribute is set to **true**. If a user creates the VLAN by using the SoftLayer API, portal, or ticket, this attribute is set to **false**.
- `subnets` - (List) The collection of subnets associated with the VLAN.

  Nested scheme for `subnets`:
  - `cidr` - (String) A subnet Classless Inter-Domain Routing prefix. The number in the range of 0 - 32 signifying the number of bits in a subnet mask.
  - `gateway` - (String) A subnet gateway address.
  - `subnet` - (String) The subnet for the VLAN. .
  - `subnet-type` - (String) A subnet can be one of several types. `PRIMARY, ADDITIONAL_PRIMARY, SECONDARY, ROUTED_TO_VLAN, SECONDARY_ON_VLAN, STORAGE_NETWORK, and STATIC_IP_ROUTED`. A `PRIMARY` subnet is the primary network that is bound to a VLAN within the SoftLayer network. An `ADDITIONAL_PRIMARY` subnet is bound to a network VLAN to augment the pool of available primary IP addresses that might be assigned to a server. A `SECONDARY` subnet is any of the secondary subnets bound to a VLAN interface. A `ROUTED_TO_VLAN` subnet is a portable subnet that can be routed to any server on a VLAN. A `SECONDARY_ON_VLAN` subnet also doesn't exist as a VLAN interface, but is routed directly to a VLAN instead of a single IP address by SoftLayer's.
  - `subnet-size` - (String) The size of the subnet for the VLAN.
- `vlan_number` - (String) The VLAN number as recorded within the SoftLayer network. This attribute is configured directly on SoftLayer's networking equipment.
