---
layout: "ibm"
page_title: "IBM: instance_group"
sidebar_current: "docs-ibm-resource-is-instance-group"
description: |-
  Manages IBM VPC instance group.
---

# ibm\_is_instance_group

Create, update or delete a instance group on VPC

## Example Usage

In the following example, you can create a instance group on VPC gen-2 infrastructure.
```hcl
provider "ibm" {
  generation = 2
}

resource "ibm_is_vpc" "vpc2" {
  name = "vpc2test"
}

resource "ibm_is_subnet" "subnet2" {
  name            = "subnet2"
  vpc             = ibm_is_vpc.vpc2.id
  zone            = "us-south-2"
  ipv4_cidr_block = "10.240.64.0/28"
}

resource "ibm_is_ssh_key" "sshkey" {
  name       = "ssh1"
  public_key = "SSH KEY"
}

resource "ibm_is_instance_template" "instancetemplate1" {
  name    = "testtemplate"
  image   = "r006-14140f94-fcc4-11e9-96e7-a72723715315"
  profile = "bx2-8x32"

  primary_network_interface {
    subnet = ibm_is_subnet.subnet2.id
  }

  vpc  = ibm_is_vpc.vpc2.id
  zone = "us-south-2"
  keys = [ibm_is_ssh_key.sshkey.id]
}

resource "ibm_is_instance_group" "instance_group" {
  name              = "testgroup"
  instance_template = ibm_is_instance_template.instancetemplate1.id
  instance_count    = 2
  subnets           = [ibm_is_subnet.subnet2.id]

  //User can configure timeouts
  timeouts {
    create = "15m"
    delete = "15m"
    update = "10m"
  }
}
```

## Timeouts

ibm_is_instance_group provides the following [Timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) configuration options:

* `create` - (Default 10 minutes) Used for creating Instance.
* `update` - (Default 10 minutes) Used for updating Instance.
* `delete` - (Default 5 minutes) Used for deleting Instance.

## Argument Reference

The following arguments are supported:

* `name` - (Required, string) The name of the instance group.
* `instance_template` - (Required, Forces new resource, string) The ID of the instance template to create the instance group.
* `instance_count` - (Optional, int) The number of instances to be created under the instance group. Default is set to 0.
  **NOTE**: instance group manager should be in disabled state to update the `instance_count`.
* `resource_group` - (Optional, string) Resource group ID.
* `subnets` - (Required, list) The list of subnet IDs used by the instances.
* `application_port` - (Optional, int) Used by the instance group when scaling up instances to supply the port for the load balancer pool member
* `load_balancer` - (Optional, string) Load blamcer ID.
* `load_balancer_pool` - (Optional, string) Load blamcer pool ID.

## Attribute Reference

The following attributes are exported:

* `id` - Id of the instance group
* `instances` - The number of instances in the intances group
* `managers` - list of managers associated with the instance group.
* `vpc` - The VPC ID
* `status` - Status of instance group.

## Import

`ibm_is_instance_group` can be imported using instance group ID, eg ibm_is_instance_group.instance_group

```
$ terraform import ibm_is_instance_group.instance_group r006-14140f94-fcc4-11e9-96e7-a72723715315
```
