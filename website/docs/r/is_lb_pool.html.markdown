---

subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : lb_pool"
description: |-
  Manages IBM load balancer pool.
---

# ibm_is_lb_pool
Create, update, or delete a VPC load balancer pool. For more information about load balancer pools, see [working with pools](https://cloud.ibm.com/docs/vpc?topic=vpc-nlb-pools).

**Note:** 
VPC infrastructure services are region-specific endpoints. By default, the provider targets `us-south`. Ensure you target the correct region in the provider block as shown in the `provider.tf` file if your VPC service is created in a different region.

**provider.tf**

```terraform
provider "ibm" {
  region = "eu-gb"
}
```

## Example usage

### Basic load balancer pool with HTTP protocol

```terraform
resource "ibm_is_lb_pool" "example" {
  name           = "example-pool"
  lb             = ibm_is_lb.example.id
  algorithm      = "round_robin"
  protocol       = "http"
  health_delay   = 60
  health_retries = 5
  health_timeout = 30
  health_type    = "http"
  proxy_protocol = "v1"
}
```

### Load balancer pool with HTTPS protocol and enhanced security

```terraform
resource "ibm_is_lb_pool" "example" {
  name           = "example-pool"
  lb             = ibm_is_lb.example.id
  algorithm      = "round_robin"
  protocol       = "https"
  health_delay   = 60
  health_retries = 5
  health_timeout = 30
  health_type    = "https"
  health_monitor_url = "/health"
  health_monitor_port = 8080
  proxy_protocol = "v1"
}

```

### Load balancer pool with app_cookie session persistence

This example demonstrates session persistence using application cookies, ideal for applications that manage their own session tokens:

```terraform
resource "ibm_is_lb_pool" "example" {
  name           = "example-pool"
  lb             = ibm_is_lb.example.id
  algorithm      = "round_robin"
  protocol       = "https"
  health_delay   = 60
  health_retries = 5
  health_timeout = 30
  health_type    = "https"
  proxy_protocol = "v1"
  session_persistence_type = "app_cookie"
  session_persistence_app_cookie_name = "JSESSIONID"
}

```

### Load balancer pool with http_cookie session persistence

This configuration uses HTTP cookies managed by the load balancer for session stickiness:

```terraform
resource "ibm_is_lb_pool" "example" {
  name           = "example-pool"
  lb             = ibm_is_lb.example.id
  algorithm      = "round_robin"
  protocol       = "https"
  health_delay   = 60
  health_retries = 5
  health_timeout = 30
  health_type    = "https"
  proxy_protocol = "v1"
  session_persistence_type = "http_cookie"
}

```

### Load balancer pool with source_ip session persistence

Source IP-based session persistence ensures requests from the same client IP are routed to the same backend:

```terraform
resource "ibm_is_lb_pool" "example" {
  name           = "example-pool"
  lb             = ibm_is_lb.example.id
  algorithm      = "round_robin"
  protocol       = "https"
  health_delay   = 60
  health_retries = 5
  health_timeout = 30
  health_type    = "https"
  proxy_protocol = "v1"
  session_persistence_type = "source_ip"
}
```

### Load balancer pool without session persistence (Route Mode Compatible)

For route mode load balancers or when session persistence isn't required, omit the session persistence parameters entirely:

```terraform
resource "ibm_is_lb_pool" "route_mode_example" {
  name           = "route-mode-pool"
  lb             = ibm_is_lb.route_mode.id
  algorithm      = "round_robin"
  protocol       = "tcp"
  health_delay   = 60
  health_retries = 5
  health_timeout = 30
  health_type    = "tcp"
  # No session_persistence_type specified - required for route mode
}
```

### Load balancer pool with failsafe policy

Configure failsafe behavior when all pool members become unhealthy:

```terraform
resource "ibm_is_lb_pool" "with_failsafe" {
  name           = "failsafe-pool"
  lb             = ibm_is_lb.example.id
  algorithm      = "least_connections"
  protocol       = "https"
  health_delay   = 30
  health_retries = 3
  health_timeout = 15
  health_type    = "https"
  
  failsafe_policy {
    action = "forward"
    target {
      id = ibm_is_lb_pool.backup_pool.pool_id
    }
  }
}
```

## Timeouts
The `ibm_is_lb_pool` resource provides the following [Timeouts](https://www.terraform.io/docs/language/resources/syntax.html) configuration options:

- **create** - (Default 10 minutes) Used for creating the load balancer pool.
- **update** - (Default 10 minutes) Used for updating the load balancer pool.
- **delete** - (Default 10 minutes) Used for deleting the load balancer pool.


## Argument reference
Review the argument references that you can specify for your resource. 

- `algorithm` - (Required, String) The load-balancing algorithm. Supported values are `round_robin`, `weighted_round_robin`, or `least_connections`. Choose `least_connections` for workloads with varying response times.
- `failsafe_policy` - (Optional, List) The failsafe policy defines behavior when all pool members are unhealthy. If unspecified, the default failsafe policy from the load balancer profile applies.
	Nested schema for **failsafe_policy**:
	- `action` - (Optional, String) Failsafe policy action: `forward` (routes requests to target pool) or `fail` (rejects requests with HTTP 503). The enumerated values for this property may [expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) in the future.
	- `target` - (Optional, List) Target pool for `forward` action. Not applicable when action is `fail`. The targets supported by this property may [expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) in the future. 
		Nested schema for **target**:
		- `href` - (Optional, String) The URL for the target load balancer pool. Mutually exclusive with `id`. Specify "null" during update to remove an existing failsafe target pool.
		- `id` - (Optional, String) The unique identifier for the target load balancer pool. Mutually exclusive with `href`. Specify "null" during update to remove an existing failsafe target pool.
- `health_delay`- (Required, Integer) Health check interval in seconds. Must be greater than the `health_timeout` value. Recommended range: 30-300 seconds.
- `health_retries`- (Required, Integer) Maximum number of health check retries before marking a member unhealthy. Recommended range: 2-10.
- `health_timeout`- (Required, Integer) Health check timeout in seconds. Must be less than `health_delay`. Recommended range: 5-60 seconds.
- `health_type` - (Required, String) The health check protocol. Supported values: `http`, `https`, `tcp`. Should typically match the pool protocol for optimal compatibility.
- `health_monitor_url` - (Optional, String) The health check URL path (e.g., `/health`, `/status`). Only applicable for `http` and `https` health check types. Defaults to `/` if not specified.
- `health_monitor_port` - (Optional, Integer) Custom health check port number. Specify `0` to remove an existing custom health check port and use the member's port. If not specified, uses the same port as the pool member.
- `lb`  - (Required, Forces new resource, String) The unique identifier of the load balancer. Changing this forces recreation of the resource.
- `name` - (Required, String) The name of the pool. Must be unique within the load balancer and follow standard naming conventions.
- `protocol` - (Required, String) The pool protocol for traffic forwarding. Supported values: `http`, `https`, `tcp`, `udp`. Choose based on your application requirements.
- `proxy_protocol` - (Optional, String) Proxy protocol setting for preserving client connection information. Supported values: `disabled` (default), `v1`, `v2`. Only supported by application load balancers, not network load balancers.
- `session_persistence_type` - (Optional, String) Session persistence method to ensure client requests are routed to the same backend server. Supported values: `source_ip`, `app_cookie`, `http_cookie`. **Important notes:**
  - Omit this parameter entirely when no session persistence is needed
  - Must be omitted for route mode load balancers
  - To remove session persistence from an existing pool, remove this parameter from your configuration and apply
  - Cannot be used with UDP protocol
- `session_persistence_app_cookie_name` - (Optional, String) Name of the application cookie used for session persistence. Required and only applicable when `session_persistence_type = "app_cookie"`. Common examples include `JSESSIONID`, `PHPSESSID`, or custom application cookies.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id` - (String) The unique identifier of the load balancer pool. The ID is composed of `<lb_id>/<pool_id>`.
- `failsafe_policy` - (List) The configured failsafe policy for this pool. If unspecified, the default failsafe policy action from the profile is used.
	Nested schema for **failsafe_policy**:
	- `healthy_member_threshold_count` - (Integer) The healthy member count threshold that triggers the failsafe policy action. Currently always `0`, but may be configurable in future versions. The minimum value is `0`.
	- `target` - (List) Target pool configuration when `action` is `forward`. Not present when `action` is `fail`. The targets supported by this property may [expand](https://cloud.ibm.com/apidocs/vpc#property-value-expansion) in the future.
		Nested schema for **target**:
		- `deleted` - (List) Indicates if the referenced target resource has been deleted, with supplementary information.
			Nested schema for **deleted**:
			- `more_info` - (String) Link to documentation about deleted resources.
		- `href` - (String) The URL for this load balancer pool.
		- `id` - (String) The unique identifier for this load balancer pool.
		- `name` - (String) The name for this load balancer pool. The name is unique across all pools for the load balancer.
- `provisioning_status` - (String) The current provisioning status of the load balancer pool. Possible values: `create_pending`, `active`, `delete_pending`, `failed`, `maintenance_pending`, `update_pending`.
- `pool_id` - (String) The unique identifier of the load balancer pool (without the load balancer prefix).
- `related_crn` - (String) The Cloud Resource Name (CRN) of the associated load balancer resource.
- `session_persistence_http_cookie_name` - (String) The HTTP cookie name used for session persistence. Only present when `session_persistence_type = "http_cookie"`. This is system-generated and read-only.

## Import
The `ibm_is_lb_pool` resource can be imported using the load balancer ID and pool ID separated by a forward slash. 

**Syntax**

```
$ terraform import ibm_is_lb_pool.example <loadbalancer_ID>/<pool_ID>
```

**Example**

```
$ terraform import ibm_is_lb_pool.example d7bec597-4726-451f-8a63-e62e6f19c32c/cea6651a-bc0a-4438-9f8a-a0770bbf3ebb
```
