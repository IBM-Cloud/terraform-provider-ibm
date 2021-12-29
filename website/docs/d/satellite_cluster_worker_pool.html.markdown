---
subcategory: "Satellite"
layout: "ibm"
page_title: "IBM : satellite_cluster_worker_pool"
description: |-
  Get information about an IBM Cloud satellite cluster worker pool.
---

# ibm_satellite_cluster

Import the details of an existing satellite cluster worker pool as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration by using interpolation syntax.


## Example usage

```terraform
data "ibm_satellite_cluster_worker_pool" "worker_pool" {
  name     = var.worker_pool_name
  cluster  = var.cluster
}
```

## Argument reference

The following arguments are supported:

* `name` - (Required, string) The name or ID of the worker pool.
* `cluster` - (Required, string) The name or ID of the satellite.
cluster.
* `region` - The name of the region.
* `resource_group_id` - The ID of the resource group.

## Attributes Reference

In addition to all arguments above, the following attributes are exported:

* `id`  - The unique identifier of the worker pool.
* `cluster`  - The name or ID of the satellite cluster.
* `flavor`  - The flavor of the satellite worker node.
* `provider`  - Provider of this offering.
* `state`  - The state of the worker pool.
* `zones`- A nested block describing the zones of this worker_pool. Nested zones blocks have the following structure:
    * `zone`- The name of the zone
    * ` workercount`- The number of worker nodes in the current worker pool
* `worker_pool_labels` -  Labels on all the workers in the worker pool.
* `host_labels`  - Host labels on the workers.
* `isolation`  - Isolation of the worker node.
* `auto_scale_enabled`  - Enable auto scalling for worker pool.
* `worker_count` - The number of workers that are attached.
