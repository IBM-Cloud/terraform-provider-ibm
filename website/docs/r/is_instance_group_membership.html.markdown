---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : is_instance_group_membership"
description: |-
  Manages Instance Group Membership
---

# ibm\_is_instance_group_membership

Update or delete a instance group memership on of an instance group

## Example Usage

```hcl
resource "is_instance_group_membership" "is_instance_group_membership" {
  instance_group            = "r006-76740f94-fcc4-11e9-96e7-a77723715315"
  instance_group_membership = "eff45fa0-de38-4368-8a0d-28dc88c2ef9b"
  name                      = "membershipname"
}
```

## Argument Reference

The following arguments are supported:

* `instance_group` - (Required, Forces new resource, string) The instance group identifier.
* `instance_group_membership` - (Required, string) The instance group membership identifier.
* `name` - (Optional, string) The name of the instance group membership.
* `action_delete` - (Optional, string) "The delete flag for this instance group membership. Must be set to true to delete instance group membership.".

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - Id is the the combination of instance group ID and instance group membership ID.
* `delete_instance_on_membership_delete` - If set to true, when deleting the membership the instance will also be deleted.
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

