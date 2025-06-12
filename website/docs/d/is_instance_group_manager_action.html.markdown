---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : instance_group_manager_action"
description: |-
  Get information about IBM VPC instance group manager action.
---

# ibm_is_instance_group_manager_action
Retrive information about an instance group manager. For more information, about VPC instance group manager action, see [scheduled scaling](https://cloud.ibm.com/docs/vpc?topic=vpc-scheduled-scaling-vpc).

**Note:** 
VPC infrastructure services are a regional specific based endpoint, by default targets to `us-south`. Please make sure to target right region in the provider block as shown in the `provider.tf` file, if VPC service is created in region other than `us-south`.

**provider.tf**

```terraform
provider "ibm" {
  region = "eu-gb"
}
```

## Example usage

```terraform
data "ibm_is_instance_group_manager_action" "example" {
  instance_group         = ibm_is_instance_group.example.id
  instance_group_manager = ibm_is_instance_group_manager.example.manager_id
  name                   = "example-instance-group-manager-action"
}
```

## Argument reference
Review the argument references that you can specify for your data source. 

- `instance_group` - (Required, String) The instance group identifier.
- `instance_group_manager` - (Required, String) The instance group manager identifier of type scheduled.
- `name` - (Required, String) The instance group manager action name.

## Attribute reference
In addition to the argument reference list, you can access the following attribute references after your data source is created. 

- `action_id` - (String) The unique identifier of the ibm_is_instance_group_manager_action.
- `auto_delete` - (String) If set to `true`, this scheduled action will be automatically deleted after it has finished and the `auto_delete_timeout` time has passed.
- `auto_delete_timeout` - (String) Amount of time in hours that are required to pass before the scheduled action will be automatically deleted once it has finished. If this value is 0, the action will be deleted on completion.
- `action_type` - (String) The type of action for the instance group.
- `cron_spec` - (String) The cron specification for a recurring scheduled action. Actions can be applied a maximum of one time within a 5 minimum period.
- `created_at` - (Timestamp) The date and time that the instance group manager action was created.
- `id` -  (String) The ID is the combination of instance group ID, instance group manager ID and instance group manager action ID.
- `last_applied_at` - (Timestamp) The date and time the scheduled action was last applied. If empty the action has never been applied.
- `membership_count` - (Integer) The number of members the instance group should have at the scheduled time.
- `max_membership_count` - (Integer) The maximum number of members the instance group should have at the scheduled time.
- `min_membership_count` - (Integer) The minimum number of members the instance group should have at the scheduled time. Default value is set to 1.
- `name` - (String) The user-defined name for this instance group manager action. Names must be unique within the instance group manager.
- `next_run_at` - (Timestamp) The date and time the scheduled action will next run. If empty the system is currently calculating the next run time.
- `resource_type` - (String) The resource type.
- `status` - (String) The status of the instance group action.

    ->**Supported Status**
      </br>&#x2022; **active**: Action is ready to be run.
      </br>&#x2022; **completed**: Action was completed successfully.
      </br>&#x2022; **failed**: Action could not be completed successfully.
      </br>&#x2022; **incompatible**: Action parameters are not compatible with the group or manager.
      </br>&#x2022; **omitted**: Action was not applied because this action's manager was disabled.

- `target_manager` - (String) The unique identifier for this instance group manager of type autoscale.
- `target_manager_name` - (String) Name of instance group manager of type autoscale.
- `updated_at` - (String) The date and time that the instance group manager action was modified.
