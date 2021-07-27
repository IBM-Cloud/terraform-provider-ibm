---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : instance_group_manager_action"
description: |-
  Manages IBM VPC instance group manager action.
---

# ibm_is_instance_group_manager_action
Create, update, or delete an instance group manager action on VPC. For more information, about instance group manager action, see [creating an instance group for auto scaling](https://cloud.ibm.com/docs/vpc?topic=vpc-creating-auto-scale-instance-group).

## Example usage

```terraform
provider "ibm" {
  generation = 2
}

resource "ibm_is_vpc" "vpc2" {
  name = "testvpc"
}

resource "ibm_is_subnet" "subnet2" {
  name            = "testsubnet"
  vpc             = ibm_is_vpc.vpc2.id
  zone            = "us-south-2"
  ipv4_cidr_block = "10.240.64.0/28"
}

resource "ibm_is_ssh_key" "sshkey" {
  name       = "testssh"
  public_key = "SSH_KEY"
}

resource "ibm_is_instance_template" "instancetemplate1" {
  name    = "testinstancetemplate"
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
  name              = "testinstancegroup"
  instance_template = ibm_is_instance_template.instancetemplate1.id
  instance_count    = 2
  subnets           = [ibm_is_subnet.subnet2.id]
}

resource "ibm_is_instance_group_manager" "instance_group_manager" {
  name           = "testinstancegroupmanager"
  instance_group = ibm_is_instance_group.instance_group.id
  manager_type   = "scheduled"
  enable_manager = true
}

resource "ibm_is_instance_group_manager_action" "instance_group_manager_action" {
  name                   = "testinstancegroupmanageraction"
  instance_group         = ibm_is_instance_group.instance_group.id
  instance_group_manager = ibm_is_instance_group_manager.instance_group_manager.manager_id
  cron_spec              = "*/5 1,2,3 * * *"
  membership_count       = 1
}
    
```
## Argument reference
Review the argument references that you can specify for your resource. 

- `cron_spec` - (Optional, String) The cron specification for a recurring scheduled action. Actions can be applied a maximum of one time within a 5 min period.
- `instance_group` - (Required, String) The instance group identifier.
- `instance_group_manager` - (Required, String) The instance group manager identifier of type scheduled.
- `membership_count` - (Optional, Integer) The number of members the instance group should have at the scheduled time.
- `max_membership_count` - (Optional, Integer) The maximum number of members the instance group should have at the scheduled time.
- `min_membership_count` - (Optional, Integer) The minimum number of members the instance group should have at the scheduled time. Default value is set to 1.
- `name` - (Optional, String) The user-defined name for this instance group manager action. Names must be unique within the instance group manager.
- `run_at` - (Optional, String) The date and time that is specified for the scheduled action. The format is in ISO 8601 format. Example: 2024-03-05T15:31:50.701Z or 2024-03-05T15:31:50.701+8:00.
- `target_manager` - (Optional, String) The unique identifier for this instance group manager of type autoscale.
 

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `action_id` - (String) The unique identifier of the `ibm_is_instance_group_manager_action`.
- `action_type` - (String) The type of action for the instance group.
- `auto_delete` - (Bool) If set to **true**, this scheduled action automatically deletes after finishing and the `auto_delete_timeout` time has passed.
- `auto_delete_timeout` - (String) An amount of time in hours that are required to pass before the scheduled action is automatically deleted once it is finish. If this value is `0`, the action will be deleted on completion.
- `created_at` - (Timestamp) The date and time that the instance group manager action was created.
- `cron_spec` - (String) The cron specification for a recurring scheduled action. Actions can be applied a maximum of one time within a 5 minutes period.
- `id` - (String) The combination ID of the instance group ID, instance group manager ID and instance group manager action ID.
- `last_applied_at` - (Timestamp) The date and time the scheduled action last applied. If empty the action has never been applied.
- `name` - (String) The user defined name for the instance group manager action. Names must be unique within the instance group manager.
- `next_run_at` - (Timestamp) The date and time the scheduled action will next run. If empty the system is currently calculating the next run time.
- `resource_type` - (String) The resource type.
- `status` - (String) The status of the instance group action. **active** Action is ready to be run. **completed** Action completed successfully. **failed** Action could not be completed successfully. **incompatible** Action parameters are not compatible with the group or manager. **omitted** Action not applied when the action's manager is disabled.
- `target_manager_name` - (String) The name of the instance group manager of type autoscale.
- `updated_at` - (Timestamp) The date and time that the instance group manager action was modified.


## Import

The `ibm_is_instance_group_manager_action` resource can be imported by using instance group ID,  instance group manager ID and instance group manager action ID.

**Example**

```
$ terraform import ibm_is_instance_group_manager_action.action r006-eea6b0b7-babd-47a8-82c5-ad73d1e10bef/r006-160b9a68-58c8-4ec3-84b0-ad553ccb1e5a/r006-94d99d1d-be65-4939-9006-1a1a767245b5
```
