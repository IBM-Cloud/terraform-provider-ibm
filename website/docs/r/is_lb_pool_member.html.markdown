---

subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : lb_pool_member"
description: |-
  Manages IBM load balancer pool member.
---

# ibm_is_lb_pool_member
Create, update, or delete a pool member for a VPC load balancer. For more information, about load balancer listener pool member, see [Creating managed pools and instance groups](https://cloud.ibm.com/docs/vpc?topic=vpc-lbaas-integration-with-instance-groups).

**Note:** 
VPC infrastructure services are a regional specific based endpoint, by default targets to `us-south`. Please make sure to target right region in the provider block as shown in the `provider.tf` file, if VPC service is created in region other than `us-south`.

**provider.tf**

```terraform
provider "ibm" {
  region = "eu-gb"
}
```

## Example usage

### Sample to create a load balancer pool member for application load balancer.

```terraform
resource "ibm_is_lb_pool_member" "example" {
  lb             = ibm_is_lb.example.id
  pool           = element(split("/", ibm_is_lb_pool.example.id), 1)
  port           = 8080
  target_address = "127.0.0.1"
  weight         = 60
}
```

### Sample to create a load balancer pool member for network load balancer.

```terraform
resource "ibm_is_lb_pool_member" "example" {
  lb        = ibm_is_lb.example.id
  pool      = element(split("/", ibm_is_lb_pool.example.id), 1)
  port      = 8080
  target_id = ibm_is_instance.example.id
  weight    = 60
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
- `weight` - (Optional, Integer) Weight of the server member. This option takes effect only when the load-balancing algorithm of its belonging pool is `weighted_round_robin`, Minimum allowed weight is `0` and Maximum allowed weight is `100`. When weight is not provided a default of 40 is returned.

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
