---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM: instance_group"
description: |-
  Get IBM VPC instance group information.
---

# ibm_is_instance_group
Retrieve information of an exisitng VPC instance group. For more information, about VPC instance group information, see [creating an instance group for auto scaling](https://cloud.ibm.com/docs/vpc?topic=vpc-creating-auto-scale-instance-group).

## Example usage
The following example gets an instance group information.

```terraform
data "ibm_is_instance_group" "instance_group_data" {
	name =  ibm_is_instance_group.instance_group.name
}

```

## Argument reference
Review the argument references that you can specify for your data source. 

- `name` - (Required, String) The name of an instance group.

## Attribute reference
In addition to the argument reference list, you can access the following attribute references after your data source is created. 

- `application_port` - (String) Scales an instances to supply the port for the Load Balancer pool member.
- `id`- (Object) The ID of an instance group.
- `instance_template` -  (String) The ID of an instance template to create an instance group.
- `instance_count` - (String) The number of instances created in an instance group.
- `load_balancer_pool` - (String) The Load Balancer pool ID.
- `managers` - (String) List of managers associated with the instance group.
- `status` - (String) The status of an instance group.
- `resource_group`-  (String) The resource group ID.
- `subnets`-  (String) The list of subnet IDs used by an instances.
- `vpc` - (String) The VPC ID.

