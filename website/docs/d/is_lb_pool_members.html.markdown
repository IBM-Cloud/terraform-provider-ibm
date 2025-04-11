---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : ibm_is_lb_pool_members"
description: |-
  Get information about LoadBalancerPoolMemberCollection
---

# ibm_is_lb_pool_members

Provides a read-only data source for LoadBalancerPoolMemberCollection. 

## Example Usage

```terraform
data "ibm_is_lb_pool_members" "example" {
	lb = ibm_is_lb.example.id
	pool = ibm_is_lb_pool.example.pool_id
}
```

## Argument Reference

Review the argument reference that you can specify for your data source.

- `lb` - (Required, String) The load balancer identifier.
- `pool` - (Required, String) The pool identifier.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

- `id` - The unique identifier of the LoadBalancerPoolMemberCollection.
- `members` - (List) Collection of members.
	Nested scheme for `members`:
	- `created_at` - (String) The date and time that this member was created.
	- `health` - (String) Health of the server member in the pool.
	- `href` - (String) The member's canonical URL.
	- `id` - (String) The unique identifier for this load balancer pool member.
	- `port` - (Integer) The port number of the application running in the server member.
	- `provisioning_status` - (String) The provisioning status of this member.
	- `target` - (List) The pool member target. Load balancers in the `network` family support virtual serverinstances. Load balancers in the `application` family support IP addresses.
		Nested scheme for `target`:
		- `address` - (String) The IP address.This property may add support for IPv6 addresses in the future. When processing a value in this property, verify that the address is in an expected format. If it is not, log an error. Optionally halt processing and surface the error, or bypass the resource on which the unexpected IP address format was encountered.
		- `crn` - (String) The CRN for this virtual server instance.
		- `deleted` - (List) If present, this property indicates the referenced resource has been deleted and providessome supplementary information.
			Nested scheme for `deleted`:
			- `more_info` - (String) Link to documentation about deleted resources.
		- `href` - (String) The URL for this virtual server instance.
		- `id` - (String) The unique identifier for this virtual server instance.
		- `name` - (String) The user-defined name for this virtual server instance (and default system hostname).
	- `weight` - (Integer) Weight of the server member. Applicable only if the pool algorithm is`weighted_round_robin`.
