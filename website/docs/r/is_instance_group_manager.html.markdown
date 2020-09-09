---
layout: "ibm"
page_title: "IBM: instance_group_manager"
sidebar_current: "docs-ibm-resource-is-instance-group-manager"
description: |-
  Manages IBM VPC instance group manager.
---

# ibm\_is_instance_group_manager

Create, Update or delete a instance group manager on of an instance group

## Example Usage

In the following example, you can create a instance group manager.

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
  public_key = "SSH_KEY"
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

resource "ibm_is_instance_group_manager" "instance_group_manager" {
  name                 = "testmanager"
  aggregation_window   = 120
  instance_group       = ibm_is_instance_group.instance_group.id
  cooldown             = 300
  manager_type         = "autoscale"
  enable_manager       = true
  max_membership_count = 2
  min_membership_count = 1
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Optional, string) The name of the instance group manager.
* `enable_manager` - (Optional, bool) Enable or disbale the instance group manager. Default is set to True.
* `instance_group` - (Required, string) The instance group ID where instance group manager is created.
* `manager_type` - (Optional, string) The type of instance group manager. Default is set to 'autoscale'
* `aggregation_window` - (Optional, int) The time window in seconds to aggregate metrics prior to evaluation
* `cooldown` - (Optional, int) The duration of time in seconds to pause further scale actions after scaling has taken place
* `max_membership_count` - (Required, int) The maximum number of members in a managed instance group
* `main_membership_count` - (Optional, int) The minimum number of members in a managed instance group. Default valeue is set to 1.

## Attribute Reference

The following attributes are exported:

* `id` - Id is the comination of instance group ID and instance group manager ID
* `policies` - list of policies associated with the instance group manager.
* `manager_id` - Id of the instance group manager

## Import

`ibm_is_instance_group_manager` can be imported using instance group ID and instance group manager ID, eg ibm_is_instance_group_manager.manager

```
$ terraform import ibm_is_instance_group_manager.manager r006-eea6b0b7-babd-47a8-82c5-ad73d1e10bef/r006-160b9a68-58c8-4ec3-84b0-ad553ccb1e5a
```
