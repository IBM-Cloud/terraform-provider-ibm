---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : load balancer"
description: |-
  Manages IBM Cloud VPC load balancer.
---

# ibm_is_lb
Retrieve information of an existing IBM VPC Load Balancer as a read-only data source. For more information, about VPC load balancer, see [load balancers for VPC overview](https://cloud.ibm.com/docs/vpc?topic=vpc-nlb-vs-elb).


## Example usage

```terraform
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

## Argument reference
Review the argument references that you can specify for your data source. 
 
- `name` - (Required, String) The name of the load balancer. 

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created. 

- `hostname` - (String) Fully qualified domain name assigned to this load balancer.
- `id` - (String) The ID of the load balancer.
- `listeners` - (String) The ID of the listeners attached to this load balancer.
- `logging`-  (Bool) Enable (**true**) or disable (**false**) datapath logging for this load balancer. If unspecified, datapath logging is disabled. This option is supported only for application load balancers.
- `operating_status` - (String) The operating status of this load balancer.
- `pools` - (List) List all the Pools attached to this load balancer.

  Nested scheme for `pools`:
	- `algorithm` - (String) The load balancing algorithm.
	- `created_at` -  (String) The date and time pool was created.
	- `href` - (String) The pool's canonical URL.
	- `id` - (String) The unique identifier for this load balancer pool.
	- `name` - (String) The user-defined name for this load balancer pool.
	- `protocol` - (String) The protocol used for this load balancer pool.
	- `provisioning_status` - (String) The provisioning status of this pool.
	- `health_monitor` - (List) The health monitor of this pool.

	  Nested scheme for `health_monitor`:
	  - `delay` - (String) The health check interval in seconds. Interval must be greater than timeout value.
	  - `max_retries` - (String) The health check max retries.
	  - `timeout` - (String) The health check timeout in seconds.
	  - `type` - (String) The protocol type of this load balancer pool health monitor.
	  - `url_path` - (String) The health monitor of this pool.
  - `instance_group` - (List) The instance group that is managing this pool.

    Nested scheme for `instance_group`:
	- `crn` - (String) The CRN for this instance group.
	- `href` - (String) The URL for this instance group.
	- `id` - (String) The unique identifier for this instance group.
	- `name` - (String) The user-defined name for this instance group.
  - `members` - (List) The backend server members of the pool.

    Nested scheme for `members`:
	- `href` - (String) The canonical URL of the member.
	- `id` - (String) The unique identifier for this load balancer pool member.
  - `session_persistence` - (List) The session persistence of this pool.

    Nested scheme for `session_persistence`:
	- `type` - (String) The session persistence type.
- `public_ips` - (String) The public IP addresses assigned to this load balancer.
- `private_ips` - (String) The private IP addresses assigned to this load balancer.
- `resource_group` - (String) The resource group where the load balancer is created.
- `security_groups`- (String) A list of security groups that are used with this load balancer. This option is supported only for application load balancers.
- `security_groups_supported`- (Bool) Indicates if this load balancer supports security groups.
- `subnets` - (String) The ID of the subnets to provision this load balancer.
- `status` - (String) The status of load balancer.
- `tags` - (String) The tags associated with the load balancer.
- `type` - (String) The type of the load balancer.
