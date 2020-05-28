---
layout: "ibm"
page_title: "IBM: container_vpc_worker_pool"
sidebar_current: "docs-ibm-resource-container-vpc-worker-pool"
description: |-
  Manages IBM container vpc worker pool.
---

# ibm\_container_vpc_worker_pool

Create or delete a worker pool. The worker pool will be attached to the specified cluster.


## Example Usage

In the following example, you can create a worker pool for a vpc cluster:

```hcl
resource "ibm_container_vpc_worker_pool" "test_pool" {
  cluster          = "my_vpc_cluster"
  worker_pool_name = "my_vpc_pool"
  flavor           = "c2.2x4"
  vpc_id           = "6015365a-9d93-4bb4-8248-79ae0db2dc21"
  worker_count     = "1"

  zones = [
    {
      name      = "us-south-1"
      subnet_id = "015ffb8b-efb1-4c03-8757-29335a07493b"
    },
  ]
}
```

In the follwoing example, you can create a worker pool for openshift cluster type with entitlement.
 ```hcl
 resource "ibm_container_vpc_worker_pool" "test_pool" {
   cluster          = "my_openshift_cluster"
   worker_pool_name = "my_openshift_vpc_pool"
   flavor           = "b3c.4x16"
   vpc_id           = "6015365a-9d93-4bb4-8248-79ae0db2dc21"
   worker_count     = "1"
   entitlement      = "cloud_pak"

   zones {
     name      = "us-south-1"
     subnet_id = "015ffb8b-efb1-4c03-8757-29335a07493b"
   }
 }
 ```

## Timeouts

ibm_container_vpc_worker_pool provides the following [Timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) configuration options:

* `create` - (Default 90 minutes) Used for creating Instance.
* `delete` - (Default 90 minutes) Used for deleting Instance.


## Argument Reference

The following arguments are supported:

* `worker_pool_name` - (Required, Forces new resource, string) The name of the worker pool.
* `cluster` - (Required, Forces new resource, string) The name or id of the cluster.
* `vpc_id` - (Required, Forces new resource, string) The Id of VPC 
* `worker_count` - (Required, Int) The number of worker nodes per zone in the worker pool.
* `flavor` - (Required, Forces new resource, string) The flavour of the worker node.
* `zones` - (Required, Forces new resource, list) A nested block describing the zones of this worker_pool. Nested zones blocks have the following structure:
  * `subnet-id` - (Required, string) The worker pool subnet to assign the cluster. 
  * `name` - (Required, string) Name of the zone.
* `labels` - (Optional, map) Labels on all the workers in the worker pool.
* `resource_group_id` - (Optional, Forces new resource, string) The ID of the resource group.  You can retrieve the value from data source `ibm_resource_group`. If not provided defaults to default resource group.
* `entitlement` - (Optional, string) The openshift cluster entitlement avoids the OCP licence charges incurred. Use cloud paks with OCP Licence entitlement to add the Openshift cluster worker pool.
    **NOTE**:
    1. It is set only for the first time creation of the worker pool, modification in the further runs will not have any impacts.
    2. Set this argument to 'cloud_pak' only if you use this cluster with a Cloud Pak that has an OpenShift entitlement
 

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the worker pool resource. The id is composed of \<cluster_name_id\>/\<worker_pool_id\>.<br/>

## Import

ibm_container_vpc_worker_pool can be imported using cluster_name_id, worker_pool_id eg

```
$ terraform import ibm_container_vpc_worker_pool.example mycluster/5c4f4d06e0dc402084922dea70850e3b-7cafe35
