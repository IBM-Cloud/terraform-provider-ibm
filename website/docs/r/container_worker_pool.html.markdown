---

subcategory: "Kubernetes Service"
layout: "ibm"
page_title: "IBM: container_worker_pool"
description: |-
  Manages IBM container worker pool.
---

# ibm_container_worker_pool

Create, update, or delete a worker pool. For more information, about container worker pool, see [adding worker nodes and zones to clusters](https://cloud.ibm.com/docs/containers?topic=containers-add_workers).

## Example usage
The following example creates the worker pool `mypool` for the cluster that is named `mycluster`. 

```terraform
resource "ibm_container_worker_pool" "testacc_workerpool" {
  worker_pool_name = "terraform_test_pool"
  machine_type     = "u2c.2x4"
  cluster          = "my_cluster"
  size_per_zone    = 1
  hardware         = "shared"
  disk_encryption  = "true"

  labels = {
    "test" = "test-pool"
  }

  //User can increase timeouts 
  timeouts {
    update = "180m"
  }
}
```

### Create the Openshift cluster worker Pool with entitlement:

```terraform
resource "ibm_container_worker_pool" "test_pool" {
  worker_pool_name = "test_openshift_wpool"
  machine_type     = "b3c.4x16"
  cluster          = "openshift_cluster_example"
  size_per_zone    = 3
  hardware         = "shared"
  disk_encryption  = "true"
  entitlement = "cloud_pak"

  labels = {
    "test" = "oc-pool"
  }
}
```

## Timeouts

ibm_container_worker_pool provides the following [Timeouts](https://www.terraform.io/docs/language/resources/syntax.html) configuration options:

- **Update**: The update of the worker pool is considered `failed` if no response is received for 90 minutes.

## Argument reference
Review the argument references that you can specify for your resource. 

- `cluster` - (Required, Forces new resource, String) The name or ID of the cluster where you want to enable or disable the feature.
- `disk_encryption` -  (Bool) Optional-If set to **true**, the worker node disks are set up with an AES 256-bit encryption. If set to **false**, the disk encryption for the worker node is disabled. For more information, see [Encrypted disks](https://cloud.ibm.com/docs/containers?topic=containers-security).Yes.
- `entitlement` - (Optional, String) If you purchased an IBM Cloud Cloud Pak that includes an entitlement to run worker nodes that are installed with OpenShift Container Platform, enter `entitlement` to create your worker pool with that entitlement so that you are not charged twice for the OpenShift license. **Note** that this option can be set only when you create the worker pool. After the worker pool is created, the cost for the OpenShift license automates when you add worker nodes to your worker pool. **Note** <ul><li> It is set only for the first time creation of the worker pool, modification in the further executes will not have any impacts.</li><li> Set this argument to `cloud_pak` only if you use this cluster with a cloud pak that has an OpenShift entitlement.</li></ul>
- `hardware` - (Optional, Forces new resource, String) The level of hardware isolation for your worker node. Use `dedicated` to have available physical resources dedicated to you only, or `shared` to allow physical resources to be shared with other IBM customers. This option is available for virtual machine worker node flavors only.
- `labels` - (Optional, Map) A list of labels that you want to add to your worker pool. The labels can help you find the worker pool more easily later.
- `machine_type` - (Required, Forces new resource, String) The machine type for your worker node. The machine type determines the amount of memory, CPU, and disk space that is available to the worker node. For an overview of supported machine types, see [Planning your worker node setup](https://cloud.ibm.com/docs/containers?topic=containers-planning_worker_nodes).
- `name` - (Required, Forces new resource, String) The name of the worker pool.
- `operating_system` - (Optional, String) The operating system of the workers in the worker pool. For supported options, see [Red Hat OpenShift on IBM Cloud version information](https://cloud.ibm.com/docs/openshift?topic=openshift-openshift_versions) or [IBM Cloud Kubernetes Service version information](https://cloud.ibm.com/docs/containers?topic=containers-cs_versions). **Note:** You will need to update or replace your workers for the change to take effect. Using terraform you can set the `ibm_container_cluster.update_all_workers` parameter to `true`.
- `resource_group_id` - (Optional, Forces new resource, String) The ID of the resource group where your cluster is provisioned into. To list resource groups, run `ibmcloud resource groups` or use the `ibm_resource_group` data source.
- `size_per_zone`  - (Required, Integer) The number of worker nodes per zone that you want to add to the worker pool.
- `taints` - (Optional, Set) A nested block that sets or removes Kubernetes taints for all worker nodes in a worker pool

  Nested scheme for `taints`:
  - `key` - (Required, String) Key for taint.
  - `value` - (Required, String) Value for taint.
  - `effect` - (Required, String) Effect for taint. Accepted values are `NoSchedule`, `PreferNoSchedule`, and `NoExecute`.
- `import_on_create` - (Optional, Bool) Import an existing WorkerPool from the cluster, instead of creating a new.
- `orphan_on_delete` - (Optional, Bool) Orphan the Worker Pool resource, instead of deleting it. The argument allows the user to remove the worker pool from the state, without deleting the actual cloud resource. The worker pool can be re-imported into the state using the `import_on_create` argument.
 

**Deprecated reference**

- `region` - (Deprecated, Forces new resource, string) The region where the cluster is provisioned. If the region is not specified it will be defaulted to provider region(IC_REGION/IBMCLOUD_REGION). To get the list of supported regions please access this [link](https://containers.bluemix.net/v1/regions) and use the alias.

 
## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id` - (String) The unique identifier of the worker pool in the format `<cluster_name_id>/<worker_pool_id>`. **Note** To reference the worker pool ID in other resources use below interpolation syntax. For example, 
`: ${element(split("/",ibm_container_worker_pool.testacc_workerpool.id),1)}`
- `state` - (String) The state of the worker pool.
- `worker_pool_id` - (String) The unique identifier of the worker pool.
- `zones` - List - A list of zones that are attached to the worker pool. 

  Nested scheme for `zones`:
  - `private_vlan` - (String) The ID of the private VLAN that is used in the zone. 
  - `public_vlan` - (String) The ID of the public VLAN that is used in the zone. 
  - `worker_count` - (Integer) The number of worker nodes that are attached to the zone.
  - `zone` - (String) The name of the zone. 
- `autoscale_enabled` - (Bool) Autoscaling is enabled on the workerpool

## Import
The `ibm_container_worker_pool` can be imported by using `cluster_name_id`, `worker_pool_id`.

**Example**

```
$ terraform import ibm_container_worker_pool.example mycluster/5c4f4d06e0dc402084922dea70850e3b-7cafe35
```
