---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : is_instance_group_membership"
description: |-
  Manages instance group membership.
---

# ibm_is_instance_group_membership
Create, update, or delete a instance group memership of an instance group. For more information, about instance group membership, see [bulk provisioning instances with instance groups](https://cloud.ibm.com/docs/vpc?topic=vpc-bulk-provisioning).

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
resource "ibm_is_instance_group_membership" "example" {
  instance_group            = ibm_is_instance_group.example.id
  instance_group_membership = data.ibm_is_instance_group_memberships.example.memberships.0.instance_group_membership
  name                      = "example-ig-membership"
}
```

## Argument reference
Review the argument references that you can specify for your resource. 

- `instance_group` - (Required, Forces new resource, String) The ID of the instance group.
- `instance_group_membership` - (Required, String) The ID of the instance group membership.
- `name` - (Optional, String) The name of the instance group membership.
- `action_delete` - (Optional, Bool) The delete flag for the instance group membership. You must set to **true** to delete the instance group membership.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `delete_instance_on_membership_delete` - (String) If set to **true**, when deleting the membership, the instance gets deleted.
- `id` - (String) The combination ID of the instance group and the instance group membership ID's.
- `instance` - (List) Nested `instance` blocks have the following structure:

  Nested scheme for `instance`:
  - `crn` - (String) The CRN for this virtual server instance.
  - `name` - (String) The user defined name for this virtual server instance (and default system hostname).
  - `virtual_server_instance` - (String) The unique identifier for this virtual server instance.
- `instance_template` - (List)  Nested `instance_template` blocks have the following structure:

  Nested scheme for `instance_template`:
  - `crn` - (String) The CRN for this instance template.
  - `instance_template` - (String) The unique identifier for this instance template.
  - `name` - (String) The unique user defined name for this instance template.
- `load_balancer_pool_member` - (String) The unique identifier for this load balancer pool member.
- `name` - (String) The user-defined name for this instance group membership. Names must be unique within the instance group.
- `status` - (String) The status of the instance group membership are:
  </br>&#x2022; **deleting** Membership is deleting dependent resources.
  </br>&#x2022; **failed** Membership is unable to maintain dependent resources.
	</br>&#x2022; **healthy** Membership is active and serving in the group.
	</br>&#x2022; **pending** Membership is waiting for dependent resources.
	</br>&#x2022; **unhealthy** Membership contains unhealthy dependent resources.
