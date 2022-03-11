---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : ibm_is_lb_pool_member"
description: |-
  Get information about LoadBalancerPoolMember
---

# ibm_is_lb_pool_member

Provides a read-only data source for LoadBalancerPoolMember.

## Example Usage

```terraform
data "ibm_is_lb_pool_member" "example" {
	member = element(split("/",ibm_is_lb_pool_member.example.id),2)
	lb = ibm_is_lb.example.id
	pool = ibm_is_lb_pool.example.pool_id
}
```

## Argument Reference

Review the argument reference that you can specify for your data source.

- `member` - (Required, String) The member identifier.
- `lb` - (Required, String) The load balancer identifier.
- `pool` - (Required, String) The pool identifier.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

- `id` - The unique identifier of the LoadBalancerPoolMember.
- `created_at` - (Required, String) The date and time that this member was created.
- `health` - (Required, String) Health of the server member in the pool.
- `href` - (Required, String) The member's canonical URL.
- `port` - (Required, Integer) The port number of the application running in the server member.
- `provisioning_status` - (Required, String) The provisioning status of this member.
- `target` - (Required, List) The pool member target. Load balancers in the `network` family support virtual serverinstances. Load balancers in the `application` family support IP addresses.
	Nested scheme for `target`:
    	- `address` - (Optional, String) The IP address.This property may add support for IPv6 addresses in the future. When processing a value in this property, verify that the address is in an expected format. If it is not, log an error. Optionally halt processing and surface the error, or bypass the resource on which the unexpected IP address format was encountered.
    	- `crn` - (Optional, String) The CRN for this virtual server instance.
    	- `deleted` - (Optional, List) If present, this property indicates the referenced resource has been deleted and providessome supplementary information.
			Nested scheme for `deleted`:
        		- `more_info` - (Required, String) Link to documentation about deleted resources.
    	- `href` - (Optional, String) The URL for this virtual server instance.
    	- `id` - (Optional, String) The unique identifier for this virtual server instance.
    	- `name` - (Optional, String) The user-defined name for this virtual server instance (and default system hostname).
- `weight` - (Optional, Integer) Weight of the server member. Applicable only if the pool algorithm is`weighted_round_robin`.