---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM: instance_group_manager_policy"
description: |-
  Get IBM VPC instance group manager policy info.
---

# ibm\_is_instance_group_manager_policy

Retrive policy info of an instance group manager

## Example Usage

In the following example, you can retrieve a policy info of an instance group manager.
```hcl
data "ibm_is_instance_group_manager_policy" "instance_group_manager_policy" {
  instance_group = "r006-76770f94-f7654-11e9-96e7-a77724435315"
  instance_group_manager = "r006-76770f94-f8764-11e9-96e7-a77726534315"
	name = "testpolicy
}
```

## Argument Reference

The following arguments are supported:
* `name` - (Required, string) The name of the policy.
* `instance_group` - (Required, string) The instance group ID.
* `instance_group_manager` - (Required, string) The instance group manager ID.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Id is the combination of instance group ID, instance group manager ID and instance group manager policy ID
* `policy_type` - The type of metric to be evaluated.
* `metric_type` - The type of metric to be evaluated. The possible values for metric types are cpu, memory, network_in and network_out
* `metric_value` - The metric value to be evaluated.
* `policy_id` - The ID of the Policy.
