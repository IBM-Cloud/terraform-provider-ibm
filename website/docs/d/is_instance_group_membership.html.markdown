---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : is_instance_group_membership"
description: |-
  Get information about instance group membership.
---

# ibm_is_instance_group_membership
Retrieve information of an instance group memership. For more information, about instance group membership, see [bulk provisioning instances with instance groups](https://cloud.ibm.com/docs/vpc?topic=vpc-bulk-provisioning).

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
data "ibm_is_instance_group_membership" "example" {
  instance_group = ibm_is_instance_group.example.id
  name           = "example-ig-membership"
}
```

## Argument reference
Review the argument references that you can specify for your data source. 

- `instance_group` - (Required, String) The instance group identifier.
- `name` - (Required, String) The name of the instance group membership.

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created.

- `id` - (String) The ID is the combination of instance group ID and instance group membership ID.
- `delete_instance_on_membership_delete` - (Bool) If set to **true**, when deleting the membership the instance will also be deleted.
- `instance_group_membership` - (String) The unique identifier for this instance group membership.
- `instance` - (List) Nested `instance` blocks have the following structure:

  Nested scheme for `instance`:
  - `crn` - (String) The CRN for this virtual server instance.
  - `virtual_server_instance` - (String) The unique identifier for this virtual server instance.
  - `name` - (String) The user-defined name for this virtual server instance (and default system hostname).
- `instance_template` - (List) Nested `instance_template` blocks have the following structure:
 
   Nested scheme for `instance_template`:
   - `crn` - (String) The CRN for this instance template.
   - `instance_template` - (String) The unique identifier for this instance template.
   - `name` - (String) The unique user-defined name for this instance template.
- `name` - (String) The user-defined name for this instance group membership. Names must be unique within the instance group.
- `load_balancer_pool_member` - (String) The unique identifier for this load balancer pool member.
- `status` - (String) The status of the instance group membership

  ->**Supported Status**
    &#x2022; **deleting** Membership is deleting dependent resources. </br>
    &#x2022; **failed** Membership was unable to maintain dependent resources.</br>
    &#x2022; **healthy** Membership is active and serving in the group. </br>
    &#x2022; **pending** Membership is waiting for dependent resources. </br>
    &#x2022; **unhealthy** Membership has unhealthy dependent resources. 
