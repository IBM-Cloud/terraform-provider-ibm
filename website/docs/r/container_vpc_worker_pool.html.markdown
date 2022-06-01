---

subcategory: "Kubernetes Service"
layout: "ibm"
page_title: "IBM: container_vpc_worker_pool"
description: |-
  Manages IBM container VPC worker pool.
---

# ibm_container_vpc_worker_pool

Create or delete a worker pool. The worker pool will be attached to the specified cluster. For more information, about VPC worker pool, see [creating clusters](https://cloud.ibm.com/docs/containers?topic=containers-clusters).


## Example usage
In the following example, you can create a worker pool for a vpc cluster:

```terraform
resource "ibm_container_vpc_worker_pool" "test_pool" {
  cluster          = "my_vpc_cluster"
  worker_pool_name = "my_vpc_pool"
  flavor           = "c2.2x4"
  vpc_id           = "6015365a-9d93-4bb4-8248-79ae0db2dc21"
  worker_count     = "1"

  zones {
    name      = "us-south-1"
    subnet_id = "015ffb8b-efb1-4c03-8757-29335a07493b"
  }
}
```

In the following example, you can create a worker pool for a vpc cluster with boot volume encryption enabled:

```terraform
resource "ibm_container_vpc_worker_pool" "test_pool" {
  cluster          = "my_vpc_cluster"
  worker_pool_name = "my_vpc_pool"
  flavor           = "c2.2x4"
  vpc_id           = "6015365a-9d93-4bb4-8248-79ae0db2dc21"
  worker_count     = "1"

  zones {
    name      = "us-south-1"
    subnet_id = "015ffb8b-efb1-4c03-8757-29335a07493b"
  }

  kms_instance_id = "8e9056e6-1936-4dd9-a0a1-51d824765e11"
  crk = "804cb251-fa0a-46f5-a442-fe42cfb0ed5f"
}
```

In the follwoing example, you can create a worker pool for openshift cluster type with entitlement.
```terraform
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

The `ibm_container_vpc_worker_pool` provides the following [Timeouts](https://www.terraform.io/docs/language/resources/syntax.html) configuration options:

- **Create** The creation of the worker pool is considered failed when no response is received for 90 minutes. 
- **Delete** The deletion of the worker pool is considered failed when no response is received for 90 minutes. 

## Argument reference
Review the argument references that you can specify for your resource. 

- `cluster` - (Required, Forces new resource, String) The name or ID of the cluster.
- `entitlement`- (Optional, String) The OpenShift cluster entitlement avoids incurred OCP license charges and use cloud pak with OCP license entitlement to add the OpenShift cluster worker pool. **Note** <ul><li> It is set as one time creation of the worker pool. There is no impacts on any modification.</li><li> Set the argument to `entitlement` only when you use cluster with a cloud pak that has an OpenShift entitlement. </li></ul>
- `flavor` - (Required, Forces new resource, String) The flavor of the worker node.
- `host_pool_id` - (Optional, String) The ID of the dedicated host pool the worker pool is associated with.
- `labels` (Optional, Map) A list of labels that you want to add to all the worker nodes in the worker pool.
- `resource_group_id` - (Optional, Forces new resource, String) The ID of the resource group. To retrieve the ID, run `ibmcloud resource groups` or use the `ibm_resource_group` data source. If no value is provided, the `default` resource group is used.
- `taints` - (Optional, Set) A nested block that sets or removes Kubernetes taints for all worker nodes in a worker pool

  Nested scheme for `taints`:
  - `key` - (Required, String) Key for taint.
  - `value` - (Required, String) Value for taint.
  - `effect` - (Required, String) Effect for taint. Accepted values are `NoSchedule`, `PreferNoSchedule`, and `NoExecute`.
 
- `vpc_id` - (Required, Forces new resource, String) The ID of the VPC.
- `worker_count`- (Required, Integer) The number of worker nodes per zone in the worker pool.
- `worker_pool_name` - (Required, Forces new resource, String) The name of the worker pool.
- `zones` - (Required, List) A nested block describes the zones of this worker pool.

  Nested scheme for `zones`:
  - `name` - (Required, String) The name of the zone.
  - `subnet_id` - (Required, String) The subnet that you want to use for your worker pool.

- `crk` - Root Key ID for boot volume encryption.
- `kms_instance_id` - Instance ID for boot volume encryption. 

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id` - (String) The unique identifier of the worker pool. The ID is composed of `<cluster_name_id>/<worker_pool_id>`.
- `worker_pool_id` -  (String) The unique identifier of the worker pool.

## Import

The `ibm_container_vpc_worker_pool` can be imported by using `cluster_name_id`, `worker_pool_id`.

**Example**

```
$ terraform import ibm_container_vpc_worker_pool.example mycluster/5c4f4d06e0dc402084922dea70850e3b-7cafe35
