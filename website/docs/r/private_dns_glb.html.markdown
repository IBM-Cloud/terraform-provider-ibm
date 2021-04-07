---

subcategory: "DNS Services"
layout: "ibm"
page_title: "IBM : ibm_dns_glb"
description: |-
  Manages IBM Private DNS global load balancer.
---

# ibm_dns_glb

Provides a private dns Global load balancer resource. This allows dns Global load balancer to be created,updated and deleted.

## Example Usage

```hcl

resource "ibm_dns_glb" "test_pdns_glb" {
  depends_on    = [ibm_dns_glb_pool.test_pdns_glb_pool]
  name          = "testglb"
  instance_id   = ibm_resource_instance.test_pdns_instance.guid
  zone_id       = ibm_dns_zone.test_pdns_glb_zone.zone_id
  description   = "new glb"
  ttl           = 120
  fallback_pool = ibm_dns_glb_pool.test_pdns_glb_pool.pool_id
  default_pools = [ibm_dns_glb_pool.test_pdns_glb_pool.pool_id]
  az_pools {
    availability_zone = "us-south-1"
    pools             = [ibm_dns_glb_pool.test_pdns_glb_pool.pool_id]
  }
}

```

## Argument Reference

The following arguments are supported:

- `instance_id` - (Required, string,ForceNew) The GUID of the private DNS.
- `name` - (Required, string) The name of the load balancer.
- `description` - (Optional,string) Descriptive text of the load balancer.
- `zone_id` - (Required, string,ForceNew) The ID of the private DNS Zone.
- `ttl` - (Optional, int) Time to live in second.
- `fallback_pool` - (Required, string) The pool ID to use when all other pools are detected as unhealthy.
- `default_pools` - (Required, list of strings) TA list of pool IDs ordered by their failover priority.
- `az_pools` - (Optional, set) Map availability zones to pool ID's.
  - `availability_zone` - (Required, string) Availability zone..
  - `pools` - (Required, list of string) List of load balancer pools.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

- `id` - The unique identifier of the private DNS Global load balancer. The id is composed of <instance_id>/<zone_id>/<glb_id>.
- `created_on` - Load Balancer creation date.
- `modified_on` - Load Balancer Modification date.
- `glb_id` - Load balancer Id.
- `health` - Healthy state of the load balancer.Possible values: [DOWN,UP,DEGRADED]

## Import

ibm_dns_glb can be imported using private DNS instance ID,zone ID and global load balancer ID, eg

```
$ terraform import ibm_dns_glb.example 6ffda12064634723b079acdb018ef308/5ffda12064634723b079acdb018ef308/435da12064634723b079acdb018ef308
```
