---
layout: "ibm"
page_title: "IBM: ibm_cis_origin_pools"
sidebar_current: "docs-ibm-cis-origin-pools"
description: |-
  Provides a IBM Cloud Internet Services Global Load Balancer Origin Pool resource.
---

# ibm_cis_origin_pool

Provides a IBM Cloud Internet Services origin pool resource. This provides a pool of origins that can be used by a IBM CIS Global Load Balancer. This resource is associated with an IBM Cloud Internet Services instance and optionally a CIS Healthcheck monitor resource.

## Example Usage

```hcl
data "ibm_cis_origin_pools" "test" {
  cis_id = var.cis_crn
}
```

## Argument Reference

The following arguments are supported:

- `cis_id` - (Required,string) The ID of the CIS service instance

## Attributes Reference

The following attributes are exported:

- `id` - ID for this load balancer pool.
- `created_on` - The RFC3339 timestamp of when the load balancer was created.
- `modified_on` - The RFC3339 timestamp of when the load balancer was last modified.
- `healthy` - The status of the origin pool.
- `name` - A short name (tag) for the pool. Only alphanumeric characters, hyphens, and underscores are allowed.
- `origins` - The list of origins within this pool. Traffic directed at this pool is balanced across all currently healthy origins, provided the pool itself is healthy. It's a complex value. See description below.
  - `name` - A human-identifiable name for the origin.
  - `address` - The IP address (IPv4 or IPv6) of the origin, or the publicly addressable hostname. Hostnames entered here should resolve directly to the origin, and not be a hostname proxied by CIS.
  - `enabled` - Whether to enable (the default) this origin within the Pool. Disabled origins will not receive traffic and are excluded from health checks. The origin will only be disabled for the current pool.
  - `weight` - The origin pool weight.
  - `healthy` - The status of origins health.
  - `disabled_at` - The disabled date and time.
  - `failure_reason` - The reason of failure.
- `description` - Free text description.
- `enabled` - Whether to enable (the default) this pool. Disabled pools will not receive traffic and are excluded from health checks. Disabling a pool will cause any load balancers using it to failover to the next pool (if any).
- `monitor` - The ID of the Monitor to use for health checking origins within this pool.
- `notification_email` - The email address to send health status notifications to. This can be an individual mailbox or a mailing list.
