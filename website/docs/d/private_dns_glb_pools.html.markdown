---
layout: "ibm"
page_title: "IBM : "
sidebar_current: "docs-ibm-datasources-dns-glb-pools"
description: |-
  Manages IBM Cloud Infrastructure Private Domain Name Service GLB Pools.
---

# ibm_dns_glb_pools

Import the details of an existing IBM Cloud Infrastructure private domain name service GLB Pools as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

```hcl

data "ibm_dns_glb_pools" "ds_pdns_glb_pools" {
  instance_id = "resource_instance_guid"
}

```

## Argument Reference

The following arguments are supported:

- `instance_id` - (Required, string) The resource instance id of the private DNS on which zones were created.

## Attribute Reference

The following attributes are exported:

- `dns_glb_pools` - List of all private domain name service GLB in the IBM PoolsCloud Infrastructure.
  - `name` - Name of the load balancer pool.
  - `description` - Descriptive text of the load balancer pool.
  - `enabled` - Whether the load balancer pool is enabled.
  - `health` - The status of GLB Pool's health.
  - `healthy_origins_threshold` - The minimum number of origins that must be healthy for this pool to serve traffic. If the number of healthy origins falls below this number, the pool will be marked unhealthy and we will failover to the next available pool.
  - `origins` - The list of origins within this pool. Traffic directed at this pool is balanced across all currently healthy origins, provided the pool itself is healthy.
    - `name` - The name of the origin server.
    - `description` - Description of the origin server.
    - `address` - The address of the origin server. It can be a hostname or an IP address.
    - `enabled` - Whether the origin server is enabled.
    - `health` - Whether the health is `true` or `false`.
    - `health_failure_reason` - The Reason for health check failure.
  - `monitor` - The ID of the load balancer monitor to be associated to this pool.
  - `notification_channel` - The notification channel.
  - `healthcheck_region` - Health check region of VSIs.
  - `healthcheck_subnets` - Health check subnet IDs of VSIs.
  - `created_on` - The time (Created On) of the DNS glb pool.
  - `modified_on` - The time (Modified On) of the DNS glb pool.
