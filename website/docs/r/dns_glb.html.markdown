---
subcategory: "DNS Services"
layout: "ibm"
page_title: "IBM : ibm_dns_glb"
description: |-
  Manages IBM Private DNS GLB.
---

# ibm_dns_glb

Provides a private DNS Global Load Balancer (GLB) resource. This allows DNS GLB to create, update, and delete. For more information, see [Working with GLBs](https://cloud.ibm.com/docs/dns-svcs?topic=dns-svcs-global-load-balancers). 


## Example usage

```terraform
resource "ibm_dns_glb" "test_pdns_glb" {
  depends_on    = [ibm_dns_glb_pool.test_pdns_glb_pool]
  name          = "testglb"
  instance_id   = ibm_resource_instance.test_pdns_instance.guid
  zone_id       = ibm_dns_zone.test_pdns_glb_zone.zone_id
  description   = "new glb"
  ttl           = 120
  enabled       = true
  fallback_pool = ibm_dns_glb_pool.test_pdns_glb_pool.pool_id
  default_pools = [ibm_dns_glb_pool.test_pdns_glb_pool.pool_id]
  az_pools {
    availability_zone = "us-south-1"
    pools             = [ibm_dns_glb_pool.test_pdns_glb_pool.pool_id]
  }
}
```

## Argument reference
Review the argument reference that you can specify for your resource. 

- `az_pools` - (Optional, Set)  Map availability zones to pool ID's.

  Nested scheme for `az_pools`:
  - `availability_zone` - (Required, String) Availability of the zone.
  - `pools`- (Required, List of string) List of Load Balancer pools.
- `default_pools`- (Required, List of string) TA list of pool IDs ordered by their failover priority.
- `description` - (Optional, String)  Descriptive text of the Load Balancer.
- `fallback_pool`- (Required, Integer) The pool ID to use when all other pools are detected as unhealthy.
- `instance_id` - (Required, Forces new resource, String) The GUID of the private DNS.
- `name` - (Required, String) The name of the Load Balancer.
- `ttl` - (Optional, Integer) The time to live (TTL) in seconds.
- `zone_id` - (Required, Forces new resource, String) The ID of the private DNS Zone.
- `enabled` - (Optional, Bool) Whether the load balancer is enabled.

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your resource is created. 

- `created_on` - (Timestamp) The time when the Load Balancer was created. 
- `glb_id` - (String) The Load Balancer ID. 
- `health` - (String) Healthy state of the Load Balancer. Possible values are `DOWN`, `UP`, or `DEGRADED`. 
- `id` - (String) The unique identifier of the DNS record. The ID is composed of `<instance_id>/<zone_id>/<glb_id>`.
- `modified_on` - (Timestamp) The time when the Load Balancer was modified.

## Import
The `ibm_dns_glb` can be imported by using private DNS instance ID, zone ID, and GLB ID.

**Example**

```
$ terraform import ibm_dns_glb.example 6ffda12064634723b079acdb018ef308/5ffda12064634723b079acdb018ef308/435da12064634723b079acdb018ef308
```
