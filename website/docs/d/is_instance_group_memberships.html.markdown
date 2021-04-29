---
layout: "ibm"
page_title: "IBM : is_instance_group_memberships"
description: |-
  Get information about InstanceGroupMembershipCollection
subcategory: "Virtual Private Cloud API"
---

# ibm\_is_instance_group_memberships

Retrive all the instance group memership info of an instance group

## Example Usage

```hcl
data "is_instance_group_memberships" "is_instance_group_memberships" {
	instance_group = "r006-76740f94-fcc4-11e9-96e7-a77723715315"
}
```

## Argument Reference

The following arguments are supported:

* `instance_group` - (Required, string) The instance group identifier.

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the InstanceGroupMembershipCollection.

* `memberships` - Collection of instance group memberships. Nested `memberships` blocks have the following structure:
	* `delete_instance_on_membership_delete` - If set to true, when deleting the membership the instance will also be deleted.
	* `instance_group_membership` - The unique identifier for this instance group membership.
	* `instance`  Nested `instance` blocks have the following structure:
		* `crn` - The CRN for this virtual server instance.
		* `virtual_server_instance` - The unique identifier for this virtual server instance.
		* `name` - The user-defined name for this virtual server instance (and default system hostname).
	* `instance_template`  Nested `instance_template` blocks have the following structure:
		* `crn` - The CRN for this instance template.
		* `instance_template` - The unique identifier for this instance template.
		* `name` - The unique user-defined name for this instance template.
	* `name` - The user-defined name for this instance group membership. Names must be unique within the instance group.
	* `load_balancer_pool_member` - The unique identifier for this load balancer pool member.
	* `status` - The status of the instance group membership
		`deleting`: Membership is deleting dependent resources
		`failed`: Membership was unable to maintain dependent resources
		`healthy`: Membership is active and serving in the group
		`pending`: Membership is waiting for dependent resources
		`unhealthy`: Membership has unhealthy dependent resources.

