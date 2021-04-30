---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : load balancer"
description: |-
  Manages IBM load balancer.
---

# ibm\_is_lbs

Import the details of existing IBM VPC Load Balancers as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.


## Example Usage

```hcl
data "ibm_is_lbs" "ds_lbs" {
 }
```

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `load_balancers` - Collection of load balancers
  * `name` -  Name of the loadbalancer.
  * `subnets` - The subnets this load balancer is part of.
    * `crn` - The CRN for this subnet.
    * `id` - The unique identifier for this subnet.
    * `href` - The URL for this subnet.
    * `name` - The user-defined name for this subnet.
  * `listeners` - The listeners of this load balancer.
    * `id` - The unique identifier for this load balancer listener.
    * `href` - The listener's canonical URL.
  * `pools` - The pools of this load balancer.
    * `href` - The pool's canonical URL.
    * `id` - The unique identifier for this load balancer pool.
    * `name` - The user-defined name for this load balancer pool.
  * `profile` - The profile to use for this load balancer.
    * `family` - The product family this load balancer profile belongs to.
    * `href` - The URL for this load balancer profile
    * `name` - The name for this load balancer profile
  * `type` - The type of the load balancer.
  * `resource_group` - The resource group where the load balancer is created.
  * `tags` - Tags associated with the load balancer.
  * `id` - The unique identifier of the load balancer.
  * `public_ips` - The public IP addresses assigned to this load balancer.
  * `private_ips` - The private IP addresses assigned to this load balancer.
  * `status` - The status of load balancer.
  * `operating_status` - The operating status of this load balancer.
  * `hostname` - Fully qualified domain name assigned to this load balancer.
  * `crn` - The load balancer's CRN.
  * `created_at` - The date and time that this load balancer was created.
  * `provisioning_status` - The provisioning status of this load balancer.Possible values: [active,create_pending,delete_pending,failed,maintenance_pending,update_pending]

