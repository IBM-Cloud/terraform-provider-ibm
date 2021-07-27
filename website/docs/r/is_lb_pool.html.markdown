---

subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : lb_pool"
description: |-
  Manages IBM load balancer pool.
---

# ibm_is_lb_pool
Create, update, or delete a VPC load balancer pool.  For more information, about load balancer pool, see [working with pool](https://cloud.ibm.com/docs/vpc?topic=vpc-nlb-pools).

## Example usage

### Sample to create a load balancer pool.

```terraform
resource "ibm_is_lb_pool" "testacc_pool" {
  name           = "test_pool"
  lb             = "addfd-gg4r4-12345"
  algorithm      = "round_robin"
  protocol       = "http"
  health_delay   = 60
  health_retries = 5
  health_timeout = 30
  health_type    = "http"
  proxy_protocol = "v1"
}

```

### Sample to create a load balancer pool with `https` protocol.

```terraform
resource "ibm_is_lb_pool" "testacc_pool" {
  name           = "test_pool"
  lb             = "addfd-gg4r4-12345"
  algorithm      = "round_robin"
  protocol       = "https"
  health_delay   = 60
  health_retries = 5
  health_timeout = 30
  health_type    = "https"
  proxy_protocol = "v1"
}

```

## Timeouts
The `ibm_is_lb_pool` resource provides the following [Timeouts](https://www.terraform.io/docs/language/resources/syntax.html) configuration options:

- **create** - (Default 10 minutes) Used for creating Instance.
- **update** - (Default 10 minutes) Used for updating Instance.
- **delete** - (Default 10 minutes) Used for deleting Instance.


## Argument reference
Review the argument references that you can specify for your resource. 

- `algorithm` - (Required, String) The load-balancing algorithm. Supported values are `round_robin`, `weighted_round_robin`, or `least_connections`.
- `health_delay`- (Required, Integer) The health check interval in seconds. Interval must be greater than `timeout` value.
- `health_retries`- (Required, Integer) The health check max retries.
- `health_timeout`- (Required, Integer) The health check timeout in seconds.
- `health_type` - (Required, String) The pool protocol. Enumeration type: `http`, `https`, `tcp` are supported.
- `health_monitor_url` - (Optional, String) The health check URL. This option is applicable only to the HTTP `health-type`.
- `health_monitor_port` - (Optional, Integer) The health check port number.
- `lb`  - (Required, Forces new resource, String) The load balancer unique identifier.
- `name` - (Required, String) The name of the pool.
- `protocol` - (Required, String) The pool protocol. Enumeration type: `http`, `https`, `tcp` are supported.
- `proxy_protocol` - (Optional, String) The proxy protocol setting for the pool that is supported by the load balancers in the application family. Valid values are `disabled`, `v1`, and `v2`. Default value is `disabled`.
- `session_persistence_type` - (Optional, String) The persistence session type. Supported enumeration type is `source_ip`. <hidden>`http_cookie`, and `app_cookie` are yet to be supported.</hidden>

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id` - (String) The unique identifier of the load balancer pool. The ID is composed of `<lb_id>/<pool_id>`.
- `provisioning_status` - (String) The status of load balancer pool.
- `pool_id` - (String) ID of the load balancer pool.
- `related_crn` - (String) The CRN of the load balancer resource.

## Import
The `ibm_is_lb_pool` resource can be imported by using the load balancer ID and pool ID. 

**Syntax**

```
$ terraform import ibm_is_lb_pool.example <loadbalancer_ID>/<pool_ID>
```

**Example**

```
$ terraform import ibm_is_lb_pool.example d7bec597-4726-451f-8a63-e62e6f19c32c/cea6651a-bc0a-4438-9f8a-a0770bbf3ebb
```
