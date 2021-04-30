---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM: instance_group"
description: |-
  Get IBM VPC instance group info.
---

# ibm\_is_instance_group

Retrieves the instance group info.

## Example Usage

In the following example, you can get the instance group info.
```hcl
data "ibm_is_instance_group" "instance_group_data" {
	name =  ibm_is_instance_group.instance_group.name
}

```

## Argument Reference

The following arguments are supported:

* `name` - (Required, string) The name of the instance group.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Id of the instance group
* `managers` - list of managers associated with the instance group.
* `vpc` - The VPC ID
* `status` - Status of instance group.
* `instance_template` - The ID of the instance template to create the instance group.
* `instance_count` - The number of instances to be created under the instance group.
* `resource_group` - Resource group ID.
* `subnets` - The list of subnet IDs used by the instances.
* `application_port` - Used by the instance group when scaling up instances to supply the port for the load balancer pool member
* `load_balancer_pool` - Load blamcer pool ID.

