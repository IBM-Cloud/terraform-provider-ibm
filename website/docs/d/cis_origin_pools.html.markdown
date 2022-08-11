---
subcategory: "Internet services"
layout: "ibm"
page_title: "IBM: ibm_cis_origin_pools"
description: |-
  Provides a IBM Cloud Internet Services Global Load Balancer origin pool resource.
---

# ibm_cis_origin_pools
Retrieve information of an IBM Cloud Internet Services origin pool resource. This provides a pool of origins that is used by an IBM Cloud Internet Services Global Load Balancer. For more information, about CIS origin pool, see [setting up origin pools](https://cloud.ibm.com/docs/cis?topic=cis-glb-features-pools).

## Example usage

```terraform
data "ibm_cis_origin_pools" "test" {
  cis_id = var.cis_crn
}
```

## Argument reference
Review the argument references that you can specify for your data source. 

- `cis_id` - (Required, String) The ID of the CIS service instance.  

## Attribute reference
In addition to the argument reference list, you can access the following attribute references after your data source is created. 

- `cis_origin_pools` - Collection of GLB pools detail.

  Nested scheme for `cis_origin_pools`:
  - `created_on` - (String) Created RFC3339 timestamp of the Load Balancer.
  - `description` - (String) The description of the origin pool.
  - `enabled` - (String) The default value is `enabled`. Disabled pools do not receive traffic, and are excluded from health checks. Disabling a pool cause any Load Balancers using it to failover to the next pool (if any).
  - `healthy` - (String) The status of the origin pool.
  - `id` - (String) The ID of the Load Balancer pool.
  - `modified_on` - (String) Last modified RFC3339 timestamp of the Load Balancer.
  - `monitor` - (String) The ID of the monitor to use for health checking origins within this pool.
  - `notification_email` - (String) The Email address to send health status notifications. This can be an individual mailbox or a mailing list.
  - `name` - (String) A short name `tag` for the pool. Only alphanumeric characters, hyphens, and underscores are allowed.
  - `origins` - (String) The list of origins within this pool. Traffic directed at this pool is balanced across all currently healthy origins, provided the pool itself is healthy. Description of it's complex value is stated.

    Nested scheme for `origins`:
	- `address` - (String) The IP address `IPv4` or `IPv6` of the origin, or the publicly addressable hostname. Hostnames entered is resolved directly to the origin, and not be a hostname proxied by CIS.
	- `enabled` - (String) The default value is `enable`. Disabled origins do not receive traffic, and are excluded from health checks. The origin is disabled only for the current pool.
	- `disabled_at` - (String) The disabled date and time.
	- `failure_reason` - (String) The failure reason.
	- `healthy` - (String) The status of origins health.
	- `name` - (String) A human-identifiable name of the origin.
	- `weight` - (String) The weight of the origin pool.
- `id` - (String) ID of the data source.