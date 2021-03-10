---
layout: "ibm"
page_title: "IBM : load balancer"
sidebar_current: "docs-ibm-data-source-is-lb"
description: |-
  Manages IBM load balancer.
---

# ibm\_is_lb

Import the details of an existing IBM VPC Load Balancer as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.


## Example Usage

```hcl
resource "ibm_is_vpc" "testacc_vpc" {
  name = "testvpc"
}

resource "ibm_is_subnet" "testacc_subnet" {
  name            = "testsubnet"
  vpc             = ibm_is_vpc.testacc_vpc.id
  zone            = "us-south-1"
  ipv4_cidr_block = "10.240.0.0/24"
}

resource "ibm_is_lb" "testacc_lb" {
  name    = "testlb"
  subnets = [ibm_is_subnet.testacc_subnet.id]
}

data "ibm_is_lb" "ds_lb" {
  name = ibm_is_lb.testacc_lb.name
}
```

## Argument Reference

The following arguments are supported:

* `name` -  Name of the loadbalancer.

## Attribute Reference

The following attributes are exported:


* `subnets` - ID of the subnets to provision this load balancer.
* `listeners` - ID of the listeners attached to this load balancer.
* `pools` - List all pools of this load balancer.
  * `algorithm` - The load balancing algorithm..
  * `created_at` - The date and time pool was created.
  * `href` - The pool's canonical URL.
  * `id` - The unique identifier for this load balancer pool.
  * `name` - The user-defined name for this load balancer pool.
  * `protocol` - The protocol used for this load balancer pool.
  * `provisioning_status` - The provisioning status of this pool.
  * `health_monitor` - The health monitor of this pool.
    * `delay` - The health check interval in seconds. Interval must be greater than timeout value.
    * `max_retries` - The health check max retries.
    * `timeout` - The health check timeout in seconds.
    * `type` - The protocol type of this load balancer pool health monitor.
    * `url_path` - The health check URL. This is applicable only to http type of health monitor.
  * `instance_group` - The instance group that is managing this pool.
    * `crn` - The CRN for this instance group.
    * `href` - The URL for this instance group.
    * `id` - The unique identifier for this instance group.
    * `name` - he user-defined name for this instance group.
  * `members` - The backend server members of the pool.
    * `href` - he member's canonical URL.
    * `id` - The unique identifier for this load balancer pool member. 
  * `session_persistence` - The session persistence of this pool.
    * `type` - The session persistence type.
* `type` - The type of the load balancer.
* `resource_group` - The resource group where the load balancer is created.
* `tags` - Tags associated with the load balancer.
* `id` - The unique identifier of the load balancer.
* `public_ips` - The public IP addresses assigned to this load balancer.
* `private_ips` - The private IP addresses assigned to this load balancer.
* `status` - The status of load balancer.
* `operating_status` - The operating status of this load balancer.
* `hostname` - Fully qualified domain name assigned to this load balancer.
* `logging` - (Optional, bool) Enable or disable datapath logging for this load balancer. If unspecified, datapath logging is disabled. This is applicable only for application load balancer. One of: false, true.

