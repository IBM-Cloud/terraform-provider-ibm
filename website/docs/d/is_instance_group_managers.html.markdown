---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM: instance_group_managers"
description: |-
  Get IBM VPC instance group managers of an instance group.
---

# ibm_is_instance_group_managers
Retrieve information of an instance group managers information of an instance group. For more information, about instance group manager, see [creating an instance group for auto scaling](https://cloud.ibm.com/docs/vpc?topic=vpc-creating-auto-scale-instance-group).

**Note:** 
VPC infrastructure services are a regional specific based endpoint, by default targets to `us-south`. Please make sure to target right region in the provider block as shown in the `provider.tf` file, if VPC service is created in region other than `us-south`.

**provider.tf**

```terraform
provider "ibm" {
  region = "eu-gb"
}
```

## Example usage
In the following example, you can retrieve list of instance group managers information.

```terraform
data "ibm_is_instance_group_managers" "example" {
  instance_group = ibm_is_instance_group.example.id
}
```

## Argument reference
Review the argument references that you can specify for your data source. 

- `instance_group` - (Required, String) The instance group ID where the instance group manager is created.

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created.

- `instance_group_managers` - (List) Nested block with list of instance manager properties.

  Nested scheme for `instance_group_manager`:
  - `actions` - (String) The list of actions of an instance group manager.
  - `aggregation_window` - (String) The time window in seconds to aggregate metrics prior to evaluation.
  - `cooldown`- (Timestamp) The duration of time in seconds to pause further scale actions after scaling has taken place.
  - `id`- (Object) This ID is the combination of instance group ID, and instance group manager ID.
  - `manager_type` - (String) The type of an instance group manager.
  - `manager_id` - (String) The instance group manager ID.
  - `max_membership_count` - (String) The maximum number of members in a managed instance group.
  - `min_membership_count` - (String) The minimum number of members in a managed instance group.
  - `policies` - (String) List of policies associated with the instance group manager.
