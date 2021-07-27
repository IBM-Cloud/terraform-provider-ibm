---

subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM: instance_group_manager"
description: |-
  Manages IBM VPC instance group manager.
---

# ibm_is_instance_group_manager
Create, update, or delete an instance group manager on VPC of an instance group. For more information, about instance group manager, see [creating an instance group for auto scaling](https://cloud.ibm.com/docs/vpc?topic=vpc-creating-auto-scale-instance-group).

## Example usage
The following example creates an instance group manager.

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

resource "ibm_is_instance_group_manager" "instance_group_manager_scheduled" {
  name           = "testinstancegroupmanager"
  instance_group = ibm_is_instance_group.instance_group.id
  manager_type   = "scheduled"
  enable_manager = true
}
```

## Argument reference
Review the argument references that you can specify for your resource. 

- `aggregation_window` - (Optional, Integer) The time window in seconds to aggregate metrics prior to evaluation.
- `cooldown` - (Optional, Integer) The duration of time in seconds to pause further scale actions after scaling has taken place.
- `enable_manager` - (Optional, Bool)  Enable or disable the instance group manager. Default value is **true**.
- `instance_group` - (Required, String) The instance group ID where instance group manager is created.
- `manager_type` - (Optional, String) The type of instance group manager. Default value is `autoscale`.
- `max_membership_count`- (Required, Integer) The maximum number of members in a managed instance group.
- `min_membership_count` - (Optional, Integer) The minimum number of members in a managed instance group. Default value is `1`.
- `name` - (Optional, String) The name of the instance group manager.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `actions` - (String) List of actions of the instance group manager.
- `id` - (String) The ID in the combination of instance group ID and instance group manager ID.
- `policies` - (String) List of policies associated with the instance group manager.
- `manager_id` - (String) The ID of the instance group manager.

## Import
The `ibm_is_instance_group_manager` resource can be imported by using the instance group ID and instance group manager ID.

**Example**

```
$ terraform import ibm_is_instance_group_manager.manager r006-eea6b0b7-babd-47a8-82c5-ad73d1e10bef/r006-160b9a68-58c8-4ec3-84b0-ad553c111115a
```
