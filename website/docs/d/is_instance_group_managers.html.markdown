---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM: instance_group_managers"
description: |-
  Get IBM VPC instance group managers of an instance group.
---

# ibm\_is_instance_group_managers

Retrive all the instance group managers info of an instance group

## Example Usage

In the following example, you can retrive list of instance group managers info.
```
data "ibm_is_instance_group_managers" "instance_group_managers" {
    instance_group = "r006-76740f94-fcc4-11e9-96e7-a77723715315"
}
```

## Argument Reference

The following arguments are supported:

* `instance_group` - (Required, string) The instance group ID where instance group manager is created.

## Attribute Reference

The following attributes are exported:

* `instance_group_managers` - Nested block containing list of instance manager properties
  * `id` - d is the the combination of instance group ID and instance group manager ID.
  * `policies` - list of policies associated with the instance group manager.
  * `manager_type` - The type of instance group manager.
  * `aggregation_window` - The time window in seconds to aggregate metrics prior to evaluation.
  * `manager_id` - instance group manager ID
  * `cooldown` - The duration of time in seconds to pause further scale actions after scaling has taken place.
  * `max_membership_count` - The maximum number of members in a managed instance group.
  * `min_membership_count` - The minimum number of members in a managed instance group. 
  * `actions` - list of actions of the instance group manager.

