---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : ibm_is_lb_pools"
description: |-
  Get information about LoadBalancerPoolCollection
---

# ibm_is_lb_pools

Provides a read-only data source for LoadBalancerPoolCollection. For more information, about load balancer pool, see [working with pool](https://cloud.ibm.com/docs/vpc?topic=vpc-nlb-pools).

## Example Usage

```terraform
data "ibm_is_lb_pools" "example" {
	lb = ibm_is_lb.example.id
}
```

## Argument Reference

Review the argument reference that you can specify for your data source.

- `lb` - (Required, Forces new resource, String) The load balancer identifier.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your data source is created.

- `id` - The unique identifier of the LoadBalancerPoolCollection.
- `pools` - (List) Collection of pools.

	Nested scheme for `pools`:

	- `algorithm` - (String) The load balancing algorithm.
	- `created_at` - (String) The date and time that this pool was created.
	- `failsafe_policy` - (List) A load balancer failsafe policy.
		Nested schema for **failsafe_policy**:
		- `action` - (String) A load balancer failsafe policy action. The enumerated values for this property may [expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) in the future, currently allowed values are:
			- `bypass`: Bypasses the members and sends requests directly to their destination IPs.
			- `drop`: Drops requests.
			- `fail`: Fails requests with an HTTP 503 status code.
			- `forward`: Forwards requests to the target pool.
		- `healthy_member_threshold_count` - (Integer) The healthy member count at which the failsafe policy action will be triggered. At present, this is always `0`, but may be modifiable in the future.
		- `target` - (List) If `action` is `forward`, the target pool to forward to.If `action` is `fail`, this property will be absent.The targets supported by this property may[expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) in the future.
			Nested schema for **target**:
			- `deleted` - (List) If present, this property indicates the referenced resource has been deleted, and providessome supplementary information.
				Nested schema for **deleted**:
				- `more_info` - (String) Link to documentation about deleted resources.
			- `href` - (String) The URL for this load balancer pool.
			- `id` - (String) The unique identifier for this load balancer pool.
			- `name` - (String) The name for this load balancer pool. The name is unique across all pools for the load balancer.
	- `health_monitor` - (List) The health monitor of this pool.
		Nested scheme for `health_monitor`:
		- `delay` - (Integer) The health check interval in seconds. Interval must be greater than timeout value.
		- `max_retries` - (Integer) The health check max retries.
		- `port` - (Integer) The health check port number. If specified, this overrides the ports specified in the server member resources.
		- `timeout` - (Integer) The health check timeout in seconds.
		- `type` - (String) The protocol type of this load balancer pool health monitor.The enumerated values for this property are expected to expand in the future. When processing this property, check for and log unknown values. Optionally halt processing and surface the error, or bypass the health monitor on which the unexpected property value was encountered.
		- `url_path` - (String) The health check URL path. Applicable only if the health monitor `type` is `http` or`https`. This value must be in the format of an [origin-form request target](https://tools.ietf.org/html/rfc7230#section-5.3.1).
	- `href` - (String) The pool's canonical URL.
	- `id` - (String) The unique identifier for this load balancer pool.
	- `instance_group` - (List) The instance group that is managing this pool.
		Nested scheme for `instance_group`:
			- `crn` - (String) The CRN for this instance group.
			- `deleted` - (List) If present, this property indicates the referenced resource has been deleted and providessome supplementary information.
				Nested scheme for `deleted`:
				- `more_info` - (String) Link to documentation about deleted resources.
				- `href` - (String) The URL for this instance group.
				- `id` - (String) The unique identifier for this instance group.
				- `name` - (String) The user-defined name for this instance group.
	- `members` - (List) The backend server members of the pool.
		Nested scheme for `members`:
			- `deleted` - (List) If present, this property indicates the referenced resource has been deleted and providessome supplementary information.
			Nested scheme for `deleted`:
				- `more_info` - (String) Link to documentation about deleted resources.
				- `href` - (String) The member's canonical URL.
				- `id` - (String) The unique identifier for this load balancer pool member.
	- `name` - (String) The user-defined name for this load balancer pool.
	- `protocol` - (String) The protocol used for this load balancer pool.The enumerated values for this property are expected to expand in the future. When processing this property, check for and log unknown values. Optionally halt processing and surface the error, or bypass the pool on which the unexpected property value was encountered.
	- `provisioning_status` - (String) The provisioning status of this pool.
	- `proxy_protocol` - (String) The PROXY protocol setting for this pool:- `v1`: Enabled with version 1 (human-readable header format)- `v2`: Enabled with version 2 (binary header format)- `disabled`: DisabledSupported by load balancers in the `application` family (otherwise always `disabled`).
	- `session_persistence` - (List) The session persistence of this pool.The enumerated values for this property are expected to expand in the future. Whenprocessing this property, check for and log unknown values. Optionally haltprocessing and surface the error, or bypass the pool on which the unexpectedproperty value was encountered.
		Nested scheme for `session_persistence`:
			- `cookie_name` - (String) The session persistence cookie name. Applicable only for type `app_cookie`. Names starting with `IBM` are not allowed.
			- `type` - (String) The session persistence type. The `http_cookie` and `app_cookie` types are applicable only to the `http` and `https` protocols.