---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM: instance_group_manager_policies"
description: |-
  Get all the IBM VPC instance group manager policies information.
---

# ibm_is_instance_group_manager_policies
Retrieve all the policies information of an instance group manager. For more information, about instance group manager policies information, see [required permissions](https://cloud.ibm.com/docs/vpc?topic=vpc-resource-authorizations-required-for-api-and-cli-calls).

**Note:** 
VPC infrastructure services are a regional specific based endpoint, by default targets to `us-south`. Please make sure to target right region in the provider block as shown in the `provider.tf` file, if VPC service is created in region other than `us-south`.

**provider.tf**

```terraform
provider "ibm" {
  region = "eu-gb"
}
```

## Example usage
In the following example, you can retrieve a policy info of an instance group manager.

```terraform
data "ibm_is_instance_group_manager_policies" "example" {
  instance_group         = ibm_is_instance_group.example.id
  instance_group_manager = ibm_is_instance_group_manager.example.manager_id
}
```

## Argument reference
Review the argument references that you can specify for your data source.

- `instance_group` - (Required, String) The instance group ID.
- `instance_group_manager` - (Required, String) The instance group manager ID.

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created.

- `instance_group_manager_policies` - (List) The list of instance group manager policies.

  Nested scheme for `instance_group_manager_policies`:
  - `id`- (Object) This ID is the combination of instance group ID, instance group manager ID and instance group manager policy ID.
  - `metric_type` - (String) The type of metric to evaluate. The possible values are `cpu`, `memory`, `network_in` and `network_out`.
  - `metric_value` -  (String) The metric value to evaluate.
  - `name` - (String) The policy name.
  - `policy_type` - (String) The type of metric to evaluate.
  - `policy_id` - (String) The policy ID.
