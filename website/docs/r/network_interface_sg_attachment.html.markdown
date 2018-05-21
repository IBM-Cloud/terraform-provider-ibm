---
layout: "ibm"
page_title: "IBM: ibm_network_interface_sg_attachment"
sidebar_current:      "docs-ibm-resource-network-interface-sg-attachment"
description: |-
  Manages binding between a security group and a network interface
---

# ibm\_network_interface_sg_attachment

Provide a resource to attach security group to a network interface. This allows attachments to be created and deleted.

For additional details, see the [IBM Cloud Infrastructure (SoftLayer) API docs](http://sldn.softlayer.com/reference/datatypes/SoftLayer_Virtual_Network_SecurityGroup_NetworkComponentBinding).

## Example Usage

```
data "ibm_security_group" "allowssh" {
  name = "allow_ssh"
}

resource "ibm_compute_vm_instance" "vsi"{
   ....
}
resource "ibm_network_interface_sg_attachment" "sg1" {
    security_group_id = "${data.ibm_security_group.allowssh.id}"
    network_interface_id = "${ibm_compute_vm_instance.vsi.public_interface_id}"
    //User can increase timeouts 
    timeouts {
      create = "15m"
    }
}
```

## Timeouts

ibm_network_interface_sg_attachment provides the following [Timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) configuration options:

* `create` - (Default 15 minutes) Used for Creating Instance.

## Argument Reference

The following arguments are supported:

* `security_group_id` - (Required, int) The ID of the security group.
* `network_interface_id` - (Required, int) The ID of the network interface to which the security group must be applied.
* `soft_reboot` - (Optional, boolean) Default `true`. If true and if a reboot is required to apply the attachment then VSI on which the network interface lies would be soft rebooted. If false then no reboot is perfomed.

**Note**: A reboot is required if this is first time any security group is applied to this network interface and it has never been rebooted since then.
