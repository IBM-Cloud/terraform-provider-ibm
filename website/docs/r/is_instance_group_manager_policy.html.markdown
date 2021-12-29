---

subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM: instance_group_manager_policy"
description: |-
  Manages IBM VPC instance group manager policy.
---

# ibm_is_instance_group_manager_policy

Create, update or delete a policy of an instance group manager. For more information, about instance group manager policy, see [creating an instance group for auto scaling](https://cloud.ibm.com/docs/vpc?topic=vpc-creating-auto-scale-instance-group).

**Note:** 
VPC infrastructure services are a regional specific based endpoint, by default targets to `us-south`. Please make sure to target right region in the provider block as shown in the `provider.tf` file, if VPC service is created in region other than `us-south`.

**provider.tf**

```terraform
provider "ibm" {
  region = "eu-gb"
}
```

## Example usage
In the following example, you can create a policy for instance group manager.

```terraform
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

## Argument reference
Review the argument references that you can specify for your resource. 

- `instance_group` - (Required, String) The instance group ID.
- `instance_group_manager` - (Required, String) The instance group manager ID for policy creation.
- `policy_type` - (Required, String) The type of metric to evaluate.
- `metric_type` - (Required, String) The type of metric to evaluate. The possible values for metric types are `cpu`, `memory`, `network_in`, and `network_out`.
- `metric_value`- (Required, Integer) The metric value to evaluate.
- `name` - (Optional, String) The name of the policy.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id` - (String) The ID in the combination of instance group ID, instance group manager ID, and instance group manager policy ID.
- `policy_id` - (String) The policy ID.

## Import
The `ibm_is_instance_group_manager_policy` resource can be imported by using instance group ID, instance group manager ID and instance group manager policy ID.

**Example**

```
$ terraform import ibm_is_instance_group_manager_policy.policy r006-eea6b0b7-babd-47a8-82c5-ad73d1e10bef/r006-160b9a68-58c8-4ec3-84b0-ad553ccb1e5a/r006-94d99d1d-be65-4939-9006-1a1a767245b5
```
