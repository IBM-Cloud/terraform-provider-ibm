---
subcategory: "Kubernetes Service"
layout: "ibm"
page_title: "IBM: container_worker_pool"
description: |-
  Manages IBM container worker pool.
---
​
# ibm_container_worker_pool
For more information, about container worker pool, see [adding worker nodes and zones to clusters](https://cloud.ibm.com/docs/containers?topic=containers-add_workers).
​
## Example usage
The following example shows how to import information about Kubernetes clusters.

```terraform
data "ibm_container_worker_pool" "testacc_ds_worker_pool"{
  worker_pool_name = ibm_container_worker_pool.test_pool.worker_pool_name
  cluster          = ibm_container_cluster.testacc_cluster.id
}
```
## Argument reference
Review the argument references that you can specify for your data source. 

- `cluster` - (Required, String) The name or ID of the cluster.
- `worker_pool_name` - (Required, String) The name of the worker pool that need to be retrieved.

## Attribute reference
Review the attribute references that are exported.

- `disk_encryption` - (String) Disk encryption on a worker.
- `id` - (String) The unique identifier of the worker pool. 
- `hardware` - (String) The level of hardware isolation for your worker node.
- `labels` - (String) Labels on all the workers in the worker pool.
- `machine_type` - (String) The machine type of the worker node.
- `resource_group_id` - (String) The ID of the worker pool resource group.
- `size_per_zone` - (String) Number of workers per zone in this pool.
- `state` - (String) Worker pool state. 
- `zones` - (String) List of zones attached to the worker_pool.
	- `private_vlan` - (String) The ID of the private VLAN.
	- `public_vlan` - (String) The ID of the public VLAN.
	- `worker_count` - (String) Number of workers attached to this zone.
  - `zone` - (String) Zone name.
- `crk` - Root Key ID for boot volume encryption.
- `kms_instance_id` - Instance ID for boot volume encryption.
