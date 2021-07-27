---

subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : lb_pool_member"
description: |-
  Manages IBM load balancer pool member.
---

# ibm_is_lb_pool_member
Create, update, or delete a pool member for a VPC load balancer.


## Example usage

### Sample to create a load balancer pool member for application load balancer.

```terraform
resource "ibm_is_lb_pool_member" "testacc_lb_mem" {
  lb             = "daac2b08-fe8a-443b-9b06-1cef79922dce"
  pool           = "f087d3bd-3da8-452d-9ce4-c1010c9fec04"
  port           = 8080
  target_address = "127.0.0.1"
  weight         = 60
}

```

### Sample to create a load balancer pool member for network load balancer.

```terraform
resource "ibm_is_lb_pool_member" "testacc_lb_mem" {
  lb             = "daac2b08-fe8a-443b-9b06-1cef79922dce"
  pool           = "f087d3bd-3da8-452d-9ce4-c1010c9fec04"
  port           = 8080
  target_id      = "54ad563a-0261-11e9-8317-bec54e704988"
  weight         = 60
}

```

## Timeouts
The `ibm_is_lb_pool_member` resource provides the following [Timeouts](https://www.terraform.io/docs/language/resources/syntax.html) configuration options:

- **create** - (Default 10 minutes) Used for creating Instance.
- **update** - (Default 10 minutes) Used for updating Instance.
- **delete** - (Default 10 minutes) Used for deleting Instance.


## Argument reference
Review the argument references that you can specify for your resource. 

 - `lb` - (Required, Forces new resource, String) The load balancer unique identifier.
- `pool` - (Required, Forces new resource, String) The load balancer pool unique identifier.
- `port`- (Required, Integer) The port number of the application running in the server member.
- `target_address` - (Required, String) The IP address of the pool member.
- `target_id` - (Required, String) The unique identifier for the virtual server instance pool member. Required for network load balancer.
- `weight` - (Optional, Integer) Weight of the server member. This option takes effect only when the load-balancing algorithm of its belonging pool is `weighted_round_robin`.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id` - (String) The unique identifier of the load balancer pool member.
- `href` - (String) The member’s canonical URL.
- `health` - (String) The health of the server member in the pool.

## Import
The `ibm_is_lb_pool_member` resource can be imported by using the load balancer ID, pool ID, pool member ID.

**Syntax**

```
$ terraform import ibm_is_lb_pool_member.example <loadbalancer_ID>/<pool_ID>/<pool_member_ID>
```

**Example**

```
$ terraform import ibm_is_lb_pool_member.example d7bec597-4726-451f-8a63-e62e6f19c32c/cea6651a-bc0a-4438-9f8a-a0770bbf3ebb/gfe6651a-bc0a-5538-8h8a-b0770bbf32cc
```
