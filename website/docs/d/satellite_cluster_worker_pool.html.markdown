---
subcategory: "Satellite"
layout: "ibm"
page_title: "IBM : satellite_cluster_worker_pool"
description: |-
  Get information about an IBM Cloud satellite cluster worker pool.
---

# ibm_satellite_cluster

Retrieve information about an existing Satellite cluster worker pool. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax. For more information, about Satellite cluster worker pool, see [Setting up Satellite hosts](https://cloud.ibm.com/docs/satellite?topic=satellite-hosts).

## Example usage

```terraform
data "ibm_satellite_cluster_worker_pool" "worker_pool" {
  name     = var.worker_pool_name
  cluster  = var.cluster
}
```

## Argument reference

Review the argument reference that you can specify for your data source.

- `name` - (Required, String) The name or ID of the worker pool.
- `cluster` - (Required, String) The name or ID of the satellite.
cluster.
- `region` - (Optional, String) The name of the region.
- `resource_group_id` - (Optional, String) The ID of the resource group.

## Attributes reference

In addition to all argument references list, you can access the following attribute references after your resource is created.

- `id`  - (String) The unique identifier of the worker pool.
- `cluster`  - (String) The name or ID of the satellite cluster.
- `flavor`  - (String) The flavor of the satellite worker node.
- `operating_system` (String) The operating system of the hosts in the worker pool.
- `provider`  - (String) Provider of this offering.
- `state`  - (String) The state of the worker pool.
- `zones`- (List) A nested block describing the zones of this worker_pool. 

  Nested scheme for `zones`:
    - `zone`- (String) The name of the zone.
    - ` workercount`- (String) The number of worker nodes in the current worker pool.
- `worker_pool_labels` -  (String) Labels on all the workers in the worker pool.
- `host_labels`  - (String) Host labels on the workers.
- `isolation`  - (String) Isolation of the worker node.
- `auto_scale_enabled`  - (String) Enable auto scalling for worker pool.
- `worker_count` - (String) The number of workers that are attached.
- `openshift_license_source` - (String) The license source for OpenShift.
