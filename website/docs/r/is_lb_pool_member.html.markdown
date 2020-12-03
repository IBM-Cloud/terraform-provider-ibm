---
layout: "ibm"
page_title: "IBM : lb_pool_member"
sidebar_current: "docs-ibm-resource-is-lb-pool-member"
description: |-
  Manages IBM load balancer pool member.
---

# ibm\_is_lb_pool_member

Provides a load balancer pool member resource. This allows load balancer pool member to be created, updated, and cancelled.


## Example Usage

In the following example, you can create a load balancer pool member for application load balancer:

```hcl
resource "ibm_is_lb_pool_member" "testacc_lb_mem" {
  lb             = "daac2b08-fe8a-443b-9b06-1cef79922dce"
  pool           = "f087d3bd-3da8-452d-9ce4-c1010c9fec04"
  port           = 8080
  target_address = "127.0.0.1"
  weight         = 60
}

```

In the following example, you can create a load balancer pool member for network load balancer:

```hcl
resource "ibm_is_lb_pool_member" "testacc_lb_mem" {
  lb             = "daac2b08-fe8a-443b-9b06-1cef79922dce"
  pool           = "f087d3bd-3da8-452d-9ce4-c1010c9fec04"
  port           = 8080
  target_id      = "54ad563a-0261-11e9-8317-bec54e704988"
  weight         = 60
}

```

## Timeouts

ibm_is_lb_pool_member provides the following [Timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) configuration options:

* `create` - (Default 10 minutes) Used for creating Instance.
* `update` - (Default 10 minutes) Used for updating Instance.
* `delete` - (Default 10 minutes) Used for deleting Instance.

## Argument Reference

The following arguments are supported:

* `pool` - (Required, Forces new resource, string) The load balancer pool unique identifier.
* `lb` - (Required, Forces new resource, string)  The load balancer unique identifier.
* `port` - (Required, int) The port number of the application running in the server member.
* `target_address` - (Required for application load balancer, string) The IP address of the pool member.
* `target_id` - (Required for network load balancer, string) The unique identifier for the virtual server instance pool member.
* `weight` - (Optional, int) Weight of the server member. This option takes effect only when the load balancing algorithm of its belonging pool is weighted_round_robin

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the load balancer pool member.
* `href` - The member's canonical URL.
* `health` - Health of the server member in the pool.

## Import

ibm_is_lb_pool_member can be imported using lbID, poolID and poolmemebrID, eg

```
$ terraform import ibm_is_lb_pool_member.example d7bec597-4726-451f-8a63-e62e6f19c32c/cea6651a-bc0a-4438-9f8a-a0770bbf3ebb/gfe6651a-bc0a-5538-8h8a-b0770bbf32cc
```
