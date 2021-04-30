---
subcategory: "Kubernetes Service"
layout: "ibm"
page_title: "IBM: container_worker_pool"
description: |-
  Manages IBM container worker pool.
---
​
# ibm\_container_worker_pool
​
Import the details of a Kubernetes cluster on IBM Cloud as a read-only data source.
​
## Example Usage
```hcl
data "ibm_container_worker_pool" "testacc_ds_worker_pool"{
  worker_pool_name = ibm_container_worker_pool.test_pool.worker_pool_name
  cluster          = ibm_container_cluster.testacc_cluster.id
}
```
​
## Argument Reference
​
The following arguments are supported:

* `cluster` - (Required, string) Name/ID  of the ClusterID.
* `worker_pool_name` - (Required, string) The Name of the worker Pool whose details are to be retrieved.
​
## Attribute Reference
​
In addition to all arguments above, the following attributes are exported:
​
* `id` - The unique identifier of the worker pool data source.
* `state` - Worker pool state.
* `zones` - List of zones attached to the worker_pool.
   * `zone` - Zone name.
   * `private_vlan` - The ID of the private VLAN.
   * `public_vlan` - The ID of the public VLAN.
   * `worker_count` - Number of workers attached to this zone.
* `machine_type` - The machine type of the worker node.
* `size_per_zone` - Number of workers per zone in this pool.
* `hardware` - The level of hardware isolation for your worker node. 
* `disk_encryption` - Disk encryption on a worker
* `labels` - Labels on all the workers in the worker pool.
* `resource_group_id` - The ID of the worker pool resource group.