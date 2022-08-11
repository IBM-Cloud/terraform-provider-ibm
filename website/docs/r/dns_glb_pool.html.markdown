---
subcategory: "DNS Services"
layout: "ibm"
page_title: "IBM : dns_glb_pool"
description: |-
  Manages IBM Private DNS GLB pool.
---

# ibm_dns_glb_pool

Provides a private DNS Global Load Balancer (GLB) pool resource. This allows DNS GLB pool to  create, update, and delete. For more information, see [Viewing GLB events](https://cloud.ibm.com/docs/dns-svcs?topic=dns-svcs-health-check-events#health-check-event-properties)


## Example usage

```terraform
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

## Argument reference
Review the argument reference that you can specify for your resource. 

- `description` - (Optional, String) Descriptive text of the origin server.
- `enabled`- (Required, Bool) Whether the origin server is enabled.
- `healthy_origins_threshold`- (Required, Integer) The minimum number of origins that must be healthy for this pool to serve traffic. If the number of healthy origins falls below this number, the pool will be marked unhealthy and will failover to the next available pool.
- `healthcheck_region` - (Optional, String) Health check region of VSIs. Examples: `us-south`,`us-east`, `eu-gb`, `eu-de`, `au-syd`, `jp-tok`, `jp-osa`, `ca-tor`, `br-sao`.
- `healthcheck_subnets` - (List, Optional) The health check subnet CRN of VSIs.
- `instance_id` - (Required, Forces new resource, String) The GUID of the private DNS on which zone has to be created.
- `monitor` - (Optional, String) The ID of the Load Balancer monitor to be associated to this pool.
- `name` - (Required, String) The name of the origin server.
- `notification_channel` - (Optional, String) The webhook URL as a notification channel.
- `origins`- (Required, Set) The list of origins within this pool. Traffic directed at this pool is balanced across all currently healthy origins, provided the pool itself is healthy.
  
  Nested scheme for `origins`:
  - `address` - (Required, String) The address of the origin server. It can be a hostname or an IP address.
  - `description` - (Optional, String)  Descriptive text of the origin server.
  - `enabled`- (Required, Bool) Whether the origin server is enabled.
  - `name` - (Required, String) The name of the origin server.

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your resource is created. 

- `created_on` - (Timestamp) The time (created On) of the DNS GLB pool. 
- `id` - (String) The unique ID of the private DNS zone. The ID is composed of `<instance_id>/<glb_pool_id>`. 
- `pool_id`- (String) The pool ID.
- `modified_on` - (Timestamp) The time (modified On) of the DNS GLB pool.
- `health`- (String) The status of DNS GLB pool's health. Possible values are `DOWN`, `UP`, `DEGRADED`.
- `origins`
  
  Nested scheme for `origins`:
  - `health`- (String) Whether the health is **true** or **false**.
  - `health_failure_reason`- (String) The reason for health check failure.

## Import
The `ibm_dns_glb_pool` can be imported by using private DNS instance ID, GLB pool ID.

**Example**

```
$ terraform import ibm_dns_glb_pool.example 6ffda12064634723b079acdb018ef308/435da12064634723b079acdb018ef308
```
