---
subcategory: "Kubernetes Service"
layout: "ibm"
page_title: "IBM: container_vpc_worker_pool"
description: |-
  Get information about a Kubernetes container vpc worker pool.
---

# ibm\_container_vpc_worker_pool

Import the details of a Kubernetes cluster worker pool on IBM Cloud as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.

## Example Usage

In the following example, you can create a worker pool for a vpc cluster:

```hcl
data "ibm_container_vpc_worker_pool" "testacc_ds_worker_pool" {
    cluster = "cluster_name"
    worker_pool_name = i"worker_pool_name
}
```


## Argument Reference

The following arguments are supported:

* `worker_pool_name` - (Required, string) The name of the worker pool.
* `cluster` - (Required, string) The name or id of the cluster.

## Attribute Reference

In addition to all arguments above, the following attributes are exported:

* `id` - The unique identifier of the worker pool resource. The id is composed of \<cluster_name_id\>/\<worker_pool_id\>.<br/>
* `vpc_id` -  The Id of VPC 
* `worker_count` - The number of worker nodes per zone in the worker pool.
* `flavor` - The flavour of the worker node.
* `zones` - A nested block describing the zones of this worker_pool. Nested zones blocks have the following structure:
  * `subnet-id` -  The worker pool subnet to assign the cluster. 
  * `name` -  Name of the zone.
* `labels` -  Labels on all the workers in the worker pool.
* `resource_group_id` -  The ID of the resource group.
* `provider` -  Provider Details of the worker Pool.
* `isolation` -  Isolation for the worker node