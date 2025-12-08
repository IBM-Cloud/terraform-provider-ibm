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


### Sample to create a application load balancer as member target for private path network load balancer.

```terraform
resource "ibm_is_lb_pool_member" "example" {
  lb        = ibm_is_lb.example.id
  pool      = element(split("/", ibm_is_lb_pool.example.id), 1)
  port      = 8080
  weight    = 60
  target_id = ibm_is_lb.example.id
}
```


### Sample to create a application load balancer as member target for private path network load balancer.

```terraform
resource "ibm_is_lb_pool_member" "example" {
  lb        = ibm_is_lb.example.id
  pool      = element(split("/", ibm_is_lb_pool.example.id), 1)
  port      = 8080
  weight    = 60
  target_id = ibm_is_lb.example.id
}
```

### Sample to create a reserved ip as a member target for network load balancer.

```terraform
  resource "ibm_is_lb_pool_member" "example" {
    lb        = ibm_is_lb.example.id
    pool      = element(split("/", ibm_is_lb_pool.example.id), 1)
    port      = 8080
    weight    = 20
    target_id = ibm_is_subnet_reserved_ip.example.id
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
- `target_address` - (Required, String) The IP address of the pool member.(Mutually exclusive with `target_id`)
- `target_id` - (Required, String) The unique identifier for the virtual server instance or application load balancer pool member or subnet reserved ip. Required for network load balancer. (Mutually exclusive with `target_address`)

- `weight` - (Optional, Integer) Weight of the server member. This option takes effect only when the load-balancing algorithm of its belonging pool is `weighted_round_robin`, Minimum allowed weight is `0` and Maximum allowed weight is `100`. Default: 50, Weight of the server member. Applicable only if the pool algorithm is weighted_round_robin.


## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id` - (String) The unique identifier of the load balancer pool member.
- `href` - (String) The member’s canonical URL.
- `health` - (String) The health of the server member in the pool.

## Import

In Terraform v1.5.0 and later, use an [`import` block](https://developer.hashicorp.com/terraform/language/import) to import the `ibm_is_lb_pool_member` resource by using `id`.
The `id` property can be formed from `load balancer ID`, `pool ID`, and `pool member ID`. For example:

```terraform
import {
  to = ibm_is_lb_pool_member.example
  id = "<loadbalancer_ID>/<pool_ID>/<pool_member_ID>"
}
```

Using `terraform import`. For example:

```console
% terraform import ibm_is_lb_pool_member.example <loadbalancer_ID>/<pool_ID>/<pool_member_ID>
```