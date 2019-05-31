---
layout: "ibm"
page_title: "IBM : lb_pool"
sidebar_current: "docs-ibm-resource-is-lb-pool"
description: |-
  Manages IBM load balancer pool.
---

# ibm\_is_lb_pool

Provides a load balancer pool resource. This allows load balancer pool to be created, updated, and cancelled.


## Example Usage

In the following example, you can create a load balancer pool:

```hcl
resource "ibm_is_lb_pool" "testacc_pool" {
  name           = "test_pool"
  lb             = "addfd-gg4r4-12345"
  algorithm      = "round_robin"
  protocol       = "http"
  health_delay   = 60
  health_retries = 5
  health_timeout = 30
  health_type    = "http"
}

```

## Timeouts

ibm_is_lb_pool provides the following [Timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) configuration options:

* `create` - (Default 60 minutes) Used for creating Instance.
* `update` - (Default 60 minutes) Used for updating Instance.
* `delete` - (Default 60 minutes) Used for deleting Instance.

## Argument Reference

The following arguments are supported:

* `name` - (Required, string) The name of the pool
* `lb` - (Required, string)  The load balancer unique identifier.
* `algorithm` - (Required, string) The load balancing algorithm. Enumeration type: round_robin, weighted_round_robin, least_connections
* `protocol` - (Required, string) The pool protocol. Enumeration type: http, tcp
* `health_delay` - (Required, int) The health check interval in seconds. Interval must be greater than timeout value
* `health_retries` - (Required, int) The health check max retries
* `health_timeout` - (Required, int) The health check timeout in seconds
* `health_type` - (Required, string) The pool protocol. Enumeration type: http, tcp
* `health_monitor_url` - (Required, string) The health check url. This option is applicable only to http type of --health-type
* `session_persistence_type` - (Optional, string) The session persistence type, Enumeration type: source_ip, http_cookie, app_cookie
* `session_persistence_cookie_name` - (Optional, string) Session persistence cookie name. This option is applicable only to --session-persistence-type

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the load balancer pool.The id is composed of \<lb_id\>/\<pool_id\>.
* `provisioning_status` - The status of load balancer pool.

## Import

ibm_is_lb_pool can be imported using lbID and poolID, eg

```
$ terraform import ibm_is_lb_pool.example d7bec597-4726-451f-8a63-e62e6f19c32c/cea6651a-bc0a-4438-9f8a-a0770bbf3ebb
```
