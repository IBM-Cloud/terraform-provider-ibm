---
layout: "ibm"
page_title: "IBM: instance_group_manager_policies"
sidebar_current: "docs-ibm-resource-is-instance-group-manager-policies"
description: |-
  Get all the IBM VPC instance group manager policies info.
---

# ibm\_is_instance_group_manager_policies

Retrive all the policies info of an instance group manager

## Example Usage

In the following example, you can retrieve a policy info of an instance group manager.
```hcl
data "ibm_is_instance_group_manager_policy" "instance_group_manager_policy" {
  instance_group = "r006-76770f94-f7654-11e9-96e7-a77724435315"
  instance_group_manager = "r006-76770f94-f8764-11e9-96e7-a77726534315"
}
```

## Argument Reference

The following arguments are supported:
* `instance_group` - (Required, string) The instance group ID.
* `instance_group_manager` - (Required, string) The instance group manager ID.

## Attribute Reference

The following attributes are exported:

* `instance_group_manager_policies` - instance group manager policies list
  * `id` - Id is the combination of instance group ID, instance group manager ID and instance group manager policy ID
  * `name` - Name of the Policy
  * `policy_type` - The type of metric to be evaluated.
  * `metric_type` - The type of metric to be evaluated. The possible values for metric types are cpu, memory, network_in and network_out
  * `metric_value` - The metric value to be evaluated.
  * `policy_id` - The ID of the Policy.
