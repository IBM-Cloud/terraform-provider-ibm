---

subcategory: "Kubernetes Service"
layout: "ibm"
page_title: "IBM: container_dedicated_host_pool"
description: |-
  Manages dedicated host pool.
---

# ibm_container_dedicated_host_pool

Create or delete a dedicated host pool. For more information, about dedicated host pool, see [Creating and managing dedicated hosts on VPC Gen 2 infrastructure](https://cloud.ibm.com/docs/containers?topic=containers-dedicated-hosts).


## Example usage
In the following example, you can create a dedicated host pool:

```terraform
resource "ibm_container_dedicated_host_pool" "test_dhostpool" {
  name         = "test_dhostpool"
  flavor_class = "bx2d"
  metro        = "dal"
}
```

## Argument reference
Review the argument references that you can specify for your resource. 

- `name` - (Required, Forces new resource, String) The name of the dedicated host pool.
- `metro`- (Required, Forces new resource, String) The metro to create the dedicated host pool in.
- `flavor_class` - (Required, Forces new resource, String) The flavor class of the dedicated host pool.
 
## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id` - (String) The unique identifier of the dedicated host pool.
- `host_count` (Int) The count of the hosts under the dedicated host pool.
- `state` - (String) The state of the dedicated host pool.
- `zones` - (List) A nested block describes the zones of this dedicated host pool.

  Nested scheme for `zones`:
  - `capacity` - (Map) A nested block describes the capacity of the zone.
    Nested scheme for `capacity`:
    - `memory_bytes` - (Int) Memory capacity of the zone.
    - `vcpu` - (Int) VCPU capacity of the zone.
  - `host_count` - (Int) The count of the hosts under the zone.
  - `zone` - (String) The name of the zone.
- `worker_pools` - (List) A nested block describes the worker pools of this dedicated host pool.

  Nested scheme for `worker_pools`:
  - `cluster_id` - (String) The ID of the cluster.
  - `worker_pool_id` -  (String) The unique identifier of the worker pool.

## Import

The `ibm_container_dedicated_host_pool` can be imported by using `id`.

**Example**

```
$ terraform import ibm_container_dedicated_host_pool.test_dhostpool <dedicated host id>
