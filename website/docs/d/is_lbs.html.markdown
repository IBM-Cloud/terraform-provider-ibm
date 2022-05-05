---
subcategory: "VPC infrastructure"
layout: "ibm"
page_title: "IBM : load balancer"
description: |-
  Manages IBM load balancers.
---

# ibm_is_lbs
Retrieve information of an existing IBM VPC load balancers as a read-only data source. For more information, about VPC load balancer, see [load balancers for VPC overview](https://cloud.ibm.com/docs/vpc?topic=vpc-nlb-vs-elb).

**Note:** 
VPC infrastructure services are a regional specific based endpoint, by default targets to `us-south`. Please make sure to target right region in the provider block as shown in the `provider.tf` file, if VPC service is created in region other than `us-south`.

**provider.tf**

```terraform
provider "ibm" {
  region = "eu-gb"
}
```

## Example usage

```terraform
data "ibm_is_lbs" "example" {
}
```


## Attribute reference
Review the attribute references that you can access after you retrieve your data source. 

- `load_balancers` - (List) The Collection of load balancers.

  Nested scheme for `load_balancers`:
	- `id` - (String) The unique identifier of the load balancer.
	- `created_at` - (String) The date and time this load balancer was created.
	- `crn` - (String) The load balancer's CRN.
	- `name` - (String) Name of the load balancer.
	- `subnets` - (List) The subnets this load balancer is part of.

      Nested scheme for `subnets`:
	  - `crn` - (String) The CRN for the subnet.
	  - `id` - (String) The unique identifier for this subnet.
	  - `href` - (String) The URL for this subnet.
	  - `name` - (String) The user-defined name for this subnet.
	- `hostname` - (String) The Fully qualified domain name assigned to this load balancer.
	- `listeners` - (List) The listeners of this load balancer.

	  Nested scheme for `listeners`:
	  - `id` - (String) The unique identifier for this load balancer listener.
	  - `href` - (String) The listener's canonical URL.
	- `operating_status` - (String) The operating status of this load balancer.
	- `pools` - (List) The pools of this load balancer.

	  Nested scheme for `pools`:
	  - `href` - (String) The pool's canonical URL.
	  - `id` - (String) The unique identifier for this load balancer pool.
	  - `name` - (String) The user-defined name for this load balancer pool.
	- `profile` - (List) The profile to use for this load balancer.

	  Nested scheme for `profile`:
	  - `family` - (String) The product family this load balancer profile belongs to.
	  - `href` - (String) The URL for this load balancer profile.
	  - `name` - (String) The name for this load balancer profile.
	- `private_ip` - (List) The primary IP address to bind to the network interface. This can be specified using an existing reserved IP, or a prototype object for a new reserved IP.

		Nested scheme for `private_ip`:
		- `address` - (String) The IP address. If the address has not yet been selected, the value will be 0.0.0.0. This property may add support for IPv6 addresses in the future. When processing a value in this property, verify that the address is in an expected format. If it is not, log an error. Optionally halt processing and surface the error, or bypass the resource on which the unexpected IP address format was encountered.
		- `href`- (String) The URL for this reserved IP
		- `name`- (String) The user-defined or system-provided name for this reserved IP
		- `reserved_ip`- (String) The unique identifier for this reserved IP
		- `resource_type`- (String) The resource type.	  
	- `private_ips` - (String) The private IP addresses assigned to this load balancer. Same as `private_ip.[].address`
	- `provisioning_status` - (String) The provisioning status of this load balancer. Possible values are: **active**, **create_pending**, **delete_pending**, **failed**, **maintenance_pending**, **update_pending**-
	- `public_ips` - (String) The public IP addresses assigned to this load balancer.
	- `resource_group` - (String) The resource group where the load balancer is created.
	- `route_mode` - (Bool) Indicates whether route mode is enabled for this load balancer.
	- `status` - (String) The status of the load balancers.
	- `type` - (String) The type of the load balancer.
	- `tags` - (String) Tags associated with the load balancer.
	- `udp_supported`- (Bool) Indicates whether this load balancer supports UDP.
