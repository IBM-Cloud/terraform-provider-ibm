---

subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM: instance_group_manager_policy"
description: |-
  Manages IBM VPC instance group manager policy.
---

# ibm\_is_instance_group_manager_policy

Create update or delete a policy of an instance group manager

## Example Usage

In the following example, you can create a policy for instance group manager.
```hcl
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

resource "ibm_is_instance_group_manager_policy" "cpuPolicy" {
  instance_group         = ibm_is_instance_group.instance_group.id
  instance_group_manager = ibm_is_instance_group_manager.instance_group_manager.manager_id
  metric_type            = "cpu"
  metric_value           = 70
  policy_type            = "target"
  name                   = "testpolicy"
}

```

## Argument Reference

The following arguments are supported:

* `name` - (Optional, string) The name of the policy.
* `policy_type` - (Required, string) The type of metric to be evaluated.
* `instance_group` - (Required, string) The instance group ID.
* `instance_group_manager` - (Required, string) The instance group manager ID for policy creation.
* `metric_type` - (Required, string) The type of metric to be evaluated. The possible values for metric types are cpu, memory, network_in and network_out
* `metric_value` - (Required, int) The metric value to be evaluated.

## Attribute Reference

The following attributes are exported:

* `id` - Id is the comination of the instance group ID, insatnce group manager ID and instance group manager policy ID.
* `policy_id` - ID of the policy

## Import

`ibm_is_instance_group_manager_policy` can be imported using instance group ID,  insatnce group manager ID and instance group manager policy ID.
eg; ibm_is_instance_group_manager_policy.policy

```
$ terraform import ibm_is_instance_group_manager_policy.policy r006-eea6b0b7-babd-47a8-82c5-ad73d1e10bef/r006-160b9a68-58c8-4ec3-84b0-ad553ccb1e5a/r006-94d99d1d-be65-4939-9006-1a1a767245b5
```
