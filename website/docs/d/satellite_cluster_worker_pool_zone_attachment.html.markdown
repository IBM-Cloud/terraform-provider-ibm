---
subcategory: "Satellite"
layout: "ibm"
page_title: "IBM : satellite_cluster_worker_pool_zone_attachment"
description: |-
  Get information about IBM Cloud satellite cluster worker pool zone attachment.
---

# ibm_satellite_cluster_worker_pool_zone_attachment

Import the details of an existing satellite cluster worker pool zone attached as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.

## Example Usage

```hcl
data "ibm_satellite_cluster_worker_pool_zone_attachment" "read_worker_pool_zone_attachment" {
  cluster     = "satellite-cluster"
  worker_pool = "default"
  zone        = "zone-4"
}
```

## Argument Reference

Review the argument reference that you can specify for your resource.

* `cluster` - (Required, String) The name or id of the cluster.
* `resource_group_id` - The ID of the resource group that the Satellite location is in. To list the resource group ID of the location, use the `GET /v2/satellite/getController` API method.
* `worker_pool` - (Required, String) The name of the worker pool.
* `zone` - (Required, String) The name of the zone to attach.

## Attribute Reference

In addition to all argument references listed, you can access the following attribute references after your resource is created.

* `id` - The unique identifier of the satellite_cluster_worker_pool_zone_attachment.
* `autobalance_enabled` - Auto enabled status.
* `messages` - Messages.
* `worker_count` - Number of workers in worker pool.