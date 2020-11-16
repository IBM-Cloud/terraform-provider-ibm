---
layout: "ibm"
page_title: "IBM : dns_glb_pool"
sidebar_current: "docs-ibm-resource-dns-glb-pool"
description: |-
  Manages IBM Private DNS glb pool.
---

# ibm_dns_glb_pool

Provides a private dns glb pool resource. This allows dns glb pool to be created,updated and deleted.

## Example Usage

```hcl

resource "ibm_dns_glb_pool" "test-pdns-pool-nw" {
  depends_on                = [ibm_dns_zone.test-pdns-glb-pool-zone]
  name                      = "testpool"
  instance_id               = ibm_resource_instance.test-pdns-glb-pool-instance.guid
  description               = "New test pool"
  enabled                   = true
  healthy_origins_threshold = 1
  origins {
    name        = "example-1"
    address     = "www.google.com"
    enabled     = true
    description = "origin pool"
  }
  monitor              = ibm_dns_glb_monitor.test-pdns-glb-monitor.monitor_id
  notification_channel = "https://mywebsite.com/dns/webhook"
  healthcheck_region   = "us-south"
  healthcheck_subnets  = [ibm_is_subnet.test-pdns-glb-subnet.resource_crn]
}

```

## Argument Reference

The following arguments are supported:

- `instance_id` - (Required, string,ForceNew) The guid of the private DNS on which zone has to be created.
- `name` - (Required, string) Name of the load balancer pool.
- `description` - (Optional,string) Descriptive text of the load balancer pool.
- `enabled` - (Optional,bool) Whether the load balancer pool is enabled.
- `healthy_origins_threshold` - (Optional,int) The minimum number of origins that must be healthy for this pool to serve traffic. If the number of healthy origins falls below this number, the pool will be marked unhealthy and we will failover to the next available pool.
- `origins` - (Required, set) The list of origins within this pool. Traffic directed at this pool is balanced across all currently healthy origins, provided the pool itself is healthy.
  - `name` - (Required,string) The name of the origin server.
  - `description` - (Optional,string) Description of the origin server.
  - `address` - (Required,string) The address of the origin server. It can be a hostname or an IP address.
  - `enabled` - (Required,bool) Whether the origin server is enabled.
- `monitor` - (Optional,string) The ID of the load balancer monitor to be associated to this pool.
- `notification_channel` - (Optional,string) The notification channel,It is a webhook url.
- `healthcheck_region` - (Optional,string) Health check region of VSIs.Allowable values: [,us-south,us-east,eu-gb,eu-du,au-syd,jp-tok]
- `healthcheck_subnets` - (Optional,List) Health check subnet crn of VSIs.

## Attribute Reference

The following attributes are exported:

- `id` - The unique identifier of the private DNS zone. The id is composed of <instance_id>/<glb_pool_id>.
- `pool_id` - Pool Id.
- `created_on` - The time (Created On) of the DNS glb pool.
- `modified_on` - The time (Modified On) of the DNS glb pool.
- `health` - The status of GLB Pool's health.Possible values: [DOWN,UP,DEGRADED]
- `origins`
  - `health` - Whether the health is `true` or `false`.
  - `health_failure_reason` - The Reason for health check failure.

## Import

ibm_dns_glb_pool can be imported using private DNS instance ID,glb pool ID, eg

```
$ terraform import ibm_dns_glb_pool.example 6ffda12064634723b079acdb018ef308/435da12064634723b079acdb018ef308
```
