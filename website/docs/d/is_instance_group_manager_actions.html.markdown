---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : ibm_is_instance_group_manager_actions"
description: |-
  Get information about ibm_is_instance_group_manager_actions
---

# ibm\_ibm_is_instance_group_manager_action

Retrive all the action info of an instance group manager

## Example Usage

```hcl
data "ibm_is_instance_group_manager_action" "ibm_is_instance_group_manager_action" {
	instance_group = "r006-76770f94-f7654-11e9-96e7-a77724435315"
	instance_group_manager = "r006-76770f94-f8764-11e9-96e7-a77726534315"
}
```

## Argument Reference

The following arguments are supported:

* `instance_group` - (Required, string) The instance group identifier.
* `instance_group_manager_scheduled` - (Required, string) The instance group manager identifier.

## Attribute Reference

The following attributes are exported:

* `instance_group_manager_actions` - Nested block containing list of instance manager actions
    * `instance_group_manager_action` - The unique identifier of the ibm_is_instance_group_manager_action.
    * `auto_delete` - If set to `true`, this scheduled action will be automatically deleted after it has finished and the `auto_delete_timeout` time has passed.
    * `auto_delete_timeout` - Amount of time in hours that are required to pass before the scheduled action will be automatically deleted once it has finished. If this value is 0, the action will be deleted on completion.
    * `created_at` - The date and time that the instance group manager action was created.
    * `name` - The user-defined name for this instance group manager action. Names must be unique within the instance group manager.
    * `resource_type` - The resource type.
    * `status` - The status of the instance group action
        `active`: Action is ready to be run
        `completed`: Action was completed successfully
        `failed`: Action could not be completed successfully
        `incompatible`: Action parameters are not compatible with the group or manager
        `omitted`: Action was not applied because this action's manager was disabled.
    * `updated_at` - The date and time that the instance group manager action was modified.
    * `action_type` - The type of action for the instance group.
    * `cron_spec` - The cron specification for a recurring scheduled action. Actions can be applied a maximum of one time within a 5 min period.
    * `last_applied_at` - The date and time the scheduled action was last applied. If empty the action has never been applied.
    * `next_run_at` - The date and time the scheduled action will next run. If empty the system is currently calculating the next run time.
    * `membership_count` - (Optional, int) "The number of members the instance group should have at the scheduled time."
    * `instance_group_manager_autoscale` - (Optional, string) The unique identifier for this instance group manager of type autoscale.
    * `instance_group_manager_autoscale_name` - (Optional, string) Name of instance group manager of type autoscale.
    * `max_membership_count` - (Optional, int) The maximum number of members the instance group should have at the scheduled time.
    * `min_membership_count` - (Optional, int) The minimum number of members the instance group should have at the scheduled time. Default valeue is set to 1.

