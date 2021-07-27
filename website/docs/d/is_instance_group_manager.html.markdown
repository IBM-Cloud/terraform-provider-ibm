---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM: instance_group_manager"
description: |-
  Get IBM VPC instance group manager.
---

# ibm_is_instance_group_manager
Retrieve information about an instance group manager. For more information, about instance group manager, see [managing an instance group](https://cloud.ibm.com/docs/vpc?topic=vpc-managing-instance-group).

## Example usage
The following example can retrieve instance group manager info.

```terraform
data "ibm_is_instance_group_manager" "instance_group_manager" {
  instance_group = "r006-76740f94-fcc4-11e9-96e7-a77723715315"
  name = "testmanager"
}
```


## Argument reference
Review the argument references that you can specify for your data source.

- `instance_group` - (Required, String) The instance group ID where instance group manager is created.
- `name` - (Required, String) The name of an instance group manager.

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created. 

- `actions` - (String) The list of actions of an instance group manager.
- `aggregation_window` - (String) The time window in seconds to aggregate metrics prior to evaluation.
- `cooldown` - (String) The duration of time in seconds to pause further scale actions after scaling has taken place.
- `id` - (String) ID is the the combination of instance group ID and instance group manager ID.
- `manager_type` - (String) The type of instance group manager.
- `manager_id` - (String) The instance group manager ID.
- `max_membership_count` - (String) The maximum number of members in a managed instance group.
- `min_membership_count` - (String) The minimum number of members in a managed instance group. 
- `policies` - (String) The list of policies associated with the instance group manager.
