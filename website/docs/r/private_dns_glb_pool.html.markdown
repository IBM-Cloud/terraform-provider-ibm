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

resource "ibm_dns_glb_pool" "test-pdns-glb-pool-nw" {
    instance_id = ibm_resource_instance.test-pdns-instance.guid
    vpc_crn = ibm_is_vpc.test_pdns_vpc.crn
    type = "vpc"
    description = "new nimi pool"
    enabled=true
    healthy_origins_threshold=1
    origins {
	    name    = "example-1"
     	address = "www.google.com"
	    enabled = true
	    description="origin pool"
    }
    monitor="7dd6841c-264e-11ea-88df-062967242a6a"
    notification_channel="https://mywebsite.com/dns/webhook"
    healthcheck_region="us-south"
    healthcheck_subnets=["0716-a4c0c123-594c-4ef4-ace3-a08858540b5e"]
}

```

## Argument Reference

The following arguments are supported:

- `instance_id` - (Required, string) The id of the private DNS on which zone has to be created.
- `name` - (Required, string) Name of the load balancer pool.
- `description` - (Optional,string) Descriptive text of the load balancer pool.
- `enabled` - (Optional,bool) Whether the load balancer pool is enabled.
- `healthy_origins_threshold` - (Optional,string) The minimum number of origins that must be healthy for this pool to serve traffic. If the number of healthy origins falls below this number, the pool will be marked unhealthy and we will failover to the next available pool.
- `origins` - (Required, list) The list of origins within this pool. Traffic directed at this pool is balanced across all currently healthy origins, provided the pool itself is healthy.
  - `name` - (Required,string) The name of the origin server.
  - `description` - (Optional,string) Description of the origin server.
  - `address` - (Required,string) The address of the origin server. It can be a hostname or an IP address.
  - `enabled` - (Optional,bool) Whether the origin server is enabled.
- `monitor` - (Optional,string) The ID of the load balancer monitor to be associated to this pool.
- `notification_channel` - (Requiredstring) The notification channel.
- `healthcheck_region` - (Requiredstring) Health check region of VSIs.
- `healthcheck_subnets` - (Required,string) Health check subnet IDs of VSIs.

## Attribute Reference

The following attributes are exported:

- `id` - The unique identifier of the private DNS zone. The id is composed of <instance_id>/<glb_pool_id>.
- `created_on` - The time (Created On) of the DNS glb pool.
- `modified_on` - The time (Modified On) of the DNS glb pool.
- `health` - The status of GLB Pool's health.
- `origins`
  - `health` - Whether the health is `true` or `false`.
  - `health_failure_reason` - The Reason for health check failure.

## Import

ibm_dns_glb_pool can be imported using private DNS instance ID,glb pool ID, eg

```
$ terraform import ibm_dns_glb_pool.example 6ffda12064634723b079acdb018ef308/435da12064634723b079acdb018ef308
```
