---
subcategory: "Kubernetes Service"
layout: "ibm"
page_title: "IBM: container_vpc_worker_pool"
description: |-
  Get information about a Kubernetes container vpc worker pool.
---

# ibm_container_vpc_worker_pool
Retrieve information about a Kubernetes cluster worker pool on IBM Cloud as a read-only data source. For more information, about VPC cluster, see [creating clusters](https://cloud.ibm.com/docs/containers?topic=containers-clusters).

## Example usage
In the following example, you can create a worker pool for a VPC cluster.

```terraform
data "ibm_container_vpc_worker_pool" "testacc_ds_worker_pool" {
    cluster = "cluster_name"
    worker_pool_name = i"worker_pool_name
}
```


## Argument reference
Review the argument references that you can specify for your data source. 

- `cluster` - (Required, String) The name or ID of the cluster.
- `worker_pool_name` - (Required, String) The name of the worker pool.

## Attribute reference
In addition to all argument reference list, you can access the following attribute references after your data source is created. 

- `flavor` - (String) The flavour of the worker node.
- `host_pool_id` -(String) The ID of the dedicated host pool the worker pool is associated with.
- `id` - (String) The unique identifier of the worker pool resource, as <cluster_name_id>/<worker_pool_id>.
- `isolation` - (String) Isolation for the worker node.
- `labels` - (String) Labels on all the workers in the worker pool. 
- `provider` - (String) Provider Details of the worker Pool.
- `resource_group_id` - (String) The ID of the resource group.
- `vpc_id` - (String) The ID of the VPC.
- `worker_count` - (String) The number of worker nodes per zone in the worker pool.
- `zones` - (String) A nested block describes the zones of the worker_pool. Nested zones blocks has `subnet-id` and `name`.

  Nested scheme for `zones`:
	- `subnet-id` - (String) The worker pool subnet to assign the cluster.
	- `subnet-name` - (String) Name of the zone.
