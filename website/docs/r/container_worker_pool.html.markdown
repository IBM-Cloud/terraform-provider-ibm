---
layout: "ibm"
page_title: "IBM: container_worker_pool"
sidebar_current: "docs-ibm-resource-container-worker-pool"
description: |-
  Manages IBM container worker pool.
---

# ibm\_container_worker_pool

Create, update, or delete a worker pool. The worker pool will be attached to the specified cluster.


## Example Usage

In the following example, you can create a worker pool:

```hcl
resource "ibm_container_worker_pool" "testacc_workerpool" {
  worker_pool_name = "terraform_test_pool"
  machine_type     = "u2c.2x4"
  cluster          = "my_cluster"
  size_per_zone    = 1
  hardware         = "shared"
  disk_encryption  = "true"
  region           = "eu-de"

  labels = {
    "test" = "test-pool"
  }

  //User can increase timeouts 
    timeouts {
      update = "180m"
    }
}
```

## Timeouts

ibm_container_worker_pool provides the following [Timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) configuration options:

* `update` - (Default 90 minutes) Used for updating Instance.

## Argument Reference

The following arguments are supported:

* `name` - (Required, string) The name of the worker pool.
* `cluster` - (Required, string) The name or id of the cluster.
* `machine_type` - (Required, string) The machine type of the worker node.
* `size_per_zone` - (Required, int) Number of workers per zone in this pool.
* `hardware` - (Optional, string) The level of hardware isolation for your worker node. Use `dedicated` to have available physical resources dedicated to you only, or `shared` to allow physical resources to be shared with other IBM customers. For IBM Cloud Public accounts, the default value is shared. For IBM Cloud Dedicated accounts, dedicated is the only available option.
* `disk_encryption` - (Optional, boolean) Set to `false` to disable encryption on a worker. Default is true.
* `labels` - (Optional, map) Labels on all the workers in the worker pool.
* `region` - (Optional, string) The region where the cluster is provisioned. If the region is not specified it will be defaulted to provider region(BM_REGION/BLUEMIX_REGION). To get the list of supported regions please access this [link](https://containers.bluemix.net/v1/regions) and use the alias.
* `resource_group_id` - (Optional, string) The ID of the resource group.  You can retrieve the value from data source `ibm_resource_group`. If not provided defaults to default resource group.
 
## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the worker pool resource. The id is composed of \<cluster_name_id\>/\<worker_pool_id\>
* `state` - Worker pool state.
* `zones` - List of zones attached to the worker_pool.
   * `zone` - Zone name.
   * `private_vlan` - The ID of the private VLAN.
   * `public_vlan` - The ID of the public VLAN.
   * `worker_count` - Number of workers attached to this zone.

## Import

ibm_container_worker_pool can be imported using cluster_name_id, worker_pool_id eg

```
$ terraform import ibm_container_worker_pool.example mycluster/5c4f4d06e0dc402084922dea70850e3b-7cafe35
