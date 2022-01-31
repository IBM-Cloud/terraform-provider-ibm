---
subcategory: "Satellite"
layout: "ibm"
page_title: "IBM : satellite_cluster_worker_pool_zone_attachment"
description: |-
  Manages IBM Cloud satellite cluster worker pool zone attachment.
---

# ibm_satellite_cluster_worker_pool_zone_attachment

Provides a resource for satellite_cluster_worker_pool_zone_attachment. This allows satellite_cluster_worker_pool_zone_attachment to be created, updated and deleted.

## Example Usage

```hcl
resource "ibm_satellite_cluster_worker_pool_zone_attachment" "satellite_cluster_worker_pool_zone_attachment" {
  cluster     = var.cluster
  worker_pool = var.worker_pool
  zone        = var.zone_name
}
```

## Argument Reference

Review the argument reference that you can specify for your resource.

* `cluster` - (Optional, Forces new resource, String) The name of the cluster.
* `resource_group_id` - (Optional, Forces new resource, String) The ID of the resource group that the Satellite location is in. To list the resource group ID of the location, use the `GET /v2/satellite/getController` API method.
* `worker_pool` - (Optional, Forces new resource, String) The name of the worker pool.
* `zone` - (Optional, Forces new resource, String) (String) The name of the zone to attach.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

* `id` - (String) The unique identifier of the satellite_cluster_worker_pool_zone_attachment.
* `autobalance_enabled` - (Optional, Boolean) Auto enabled status.
* `messages` - (Optional, List) Messages.
* `worker_count` - (Optional, Integer) Number of workers in worker pool.

## Import

You can import the `ibm_satellite_cluster_worker_pool_zone_attachment` resource can be imported by using the `cluster` and `worker pool`, `zone name`.

```
<cluster>/<worker_pool>/<zone_name>
```
* `cluster`: A string. The cluster ID.
* `worker_pool`: A string. The worker pool name.
* `zone_name`: A string. The zone name.

# Syntax
```
$ terraform import ibm_satellite_cluster_worker_pool_zone_attachment.satellite_cluster_worker_pool_zone_attachment <cluster>/<worker_pool>/<zone_name>
```