---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : ibm_is_instance_group_manager_action"
description: |-
  Manages ibm_is_instance_group_manager_action.
---

# ibm\_ibm_is_instance_group_manager_action

This allows ibm_is_instance_group_manager_action to be created, updated and deleted.

## Example Usage

```hcl
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
	
	   vpc       = ibm_is_vpc.vpc2.id
	   zone      = "us-south-2"
	   keys      = [ibm_is_ssh_key.sshkey.id]
	 }
		
	resource "ibm_is_instance_group" "instance_group" {
		name =  "testinstancegroup"
		instance_template = ibm_is_instance_template.instancetemplate1.id
		instance_count =  2
		subnets = [ibm_is_subnet.subnet2.id]
	}

	resource "ibm_is_instance_group_manager" "instance_group_manager" {
		name = "testinstancegroupmanager"
		instance_group = ibm_is_instance_group.instance_group.id
		manager_type = "scheduled"
		enable_manager = true
	}

	resource "ibm_is_instance_group_manager_action" "instance_group_manager_action" {
		name = "testinstancegroupmanageraction"
		instance_group = ibm_is_instance_group.instance_group.id
		instance_group_manager_scheduled = ibm_is_instance_group_manager.instance_group_manager.manager_id
		cron_spec = "*/5 1,2,3 * * *"
		membership_count = 1
	}
    
```

## Argument Reference

The following arguments are supported:

* `instance_group` - (Required, string) The instance group identifier.
* `instance_group_manager_scheduled` - (Required, string) The instance group manager identifier.
* `name` - (Optional, string) The user-defined name for this instance group manager action. Names must be unique within the instance group manager.
* `run_at` - (Optional, string) The date and time the scheduled action will run.
* `cron_spec` - (Optional, string) The cron specification for a recurring scheduled action. Actions can be applied a maximum of one time within a 5 min period.
* `membership_count` - (Optional, int) "The number of members the instance group should have at the scheduled time."
* `instance_group_manager_autoscale` - (Optional, string) The unique identifier for this instance group manager of type autoscale.
* `max_membership_count` - (Optional, int) The maximum number of members the instance group should have at the scheduled time.
* `min_membership_count` - (Optional, int) The minimum number of members the instance group should have at the scheduled time. Default valeue is set to 1.
 

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `instance_group_manager_action` - The unique identifier of the ibm_is_instance_group_manager_action.
* `auto_delete` - If set to `true`, this scheduled action will be automatically deleted after it has finished and the `auto_delete_timeout` time has passed.
* `auto_delete_timeout` - Amount of time in hours that are required to pass before the scheduled action will be automatically deleted once it has finished. If this value is 0, the action will be deleted on completion.
* `created_at` - The date and time that the instance group manager action was created.
* `name` - The user-defined name for this instance group manager action. Names must be unique within the instance group manager.
* `resource_type` - The resource type.
* `status` - The status of the instance group action- `active`: Action is ready to be run- `completed`: Action was completed successfully- `failed`: Action could not be completed successfully- `incompatible`: Action parameters are not compatible with the group or manager- `omitted`: Action was not applied because this action's manager was disabled.
* `updated_at` - The date and time that the instance group manager action was modified.
* `action_type` - The type of action for the instance group.
* `cron_spec` - The cron specification for a recurring scheduled action. Actions can be applied a maximum of one time within a 5 min period.
* `last_applied_at` - The date and time the scheduled action was last applied. If empty the action has never been applied.
* `next_run_at` - The date and time the scheduled action will next run. If empty the system is currently calculating the next run time.
* `instance_group_manager_autoscale_name` - (Optional, string) Name of instance group manager of type autoscale.