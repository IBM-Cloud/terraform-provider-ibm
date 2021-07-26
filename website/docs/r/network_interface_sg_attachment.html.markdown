---

subcategory: "Classic infrastructure"
layout: "ibm"
page_title: "IBM: ibm_network_interface_sg_attachment"
description: |-
  Manages binding between a security group and a network interface.
---

# ibm_network_interface_sg_attachment
Create, delete, and update to attach security group to a network interface. For more information, about security group and an network interface, see [integrating an IBM Cloud Application Load Balancer for VPC with security groups](https://cloud.ibm.com/docs/vpc?topic=vpc-alb-integration-with-security-groups).

**Note**

For more information, see the [IBM Cloud Classic Infrastructure  (SoftLayer) API docs](http://sldn.softlayer.com/reference/datatypes/SoftLayer_Virtual_Network_SecurityGroup_NetworkComponentBinding).

## Example usage

```terraform
data "ibm_security_group" "allowssh" {
  name = "allow_ssh"
}

resource "ibm_compute_vm_instance" "vsi" {
}

resource "ibm_network_interface_sg_attachment" "sg1" {
  security_group_id    = data.ibm_security_group.allowssh.id
  network_interface_id = ibm_compute_vm_instance.vsi.public_interface_id

  //User can increase timeouts
  timeouts {
    create = "15m"
  }
}
```

## Timeouts

The `ibm_network_interface_sg_attachment` resource provides the following [Timeouts](https://www.terraform.io/docs/language/resources/syntax.html) configuration options:

- **create** - (Default 15 minutes) Used for creating instance.

## Argument reference 
Review the argument references that you can specify for your resource.

- `security_group_id` - (Required, Forces new resource, Integer) The ID of the security group.
- `network_interface_id` - (Required, Forces new resource, Integer) The ID of the network interface to which the security group must be applied.
- `soft_reboot`-  (Optional, Forces new resource, Bool) Default **true**. If set to **true** and a reboot is required to apply the security group attachment for the virtual server instance, then a soft reboot is performed. If set to **false**, no reboot is performed. 

**Note** 

A reboot is always required the first time a security group is applied to a network interface of a virtual server instance that was never rebooted before.
