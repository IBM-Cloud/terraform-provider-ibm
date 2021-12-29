---

subcategory: "Kubernetes Service"
layout: "ibm"
page_title: "IBM: container_worker_pool_zone_attachment"
description: |-
  Manages IBM container worker pool zone attachment.
---

# ibm_container_worker_pool_zone_attachment
Create, update, or delete a zone from a worker pool. For more information, about IBM container worker pool zone, see [adding worker nodes and zones to clusters](https://cloud.ibm.com/docs/containers?topic=containers-add_workers).

## Example usage
In the following example, you can create a zone:

```terraform
resource "ibm_container_worker_pool" "test_pool" {
  worker_pool_name = "my_pool"
  machine_type     = "u2c.2x4"
  cluster          = "my_cluster"
  size_per_zone    = 2
  hardware         = "shared"
  disk_encryption  = "true"
  labels = {
    "test"  = "test-pool"
    "test1" = "test-pool1"
  }
}

resource "ibm_container_worker_pool_zone_attachment" "test_zone" {
  cluster         = "my_cluster"
  worker_pool     = element(split("/", ibm_container_worker_pool.test_pool.id), 1)
  zone            = "dal12"
  private_vlan_id = "2320267"
  public_vlan_id  = "2320265"

  //User can increase timeouts
  timeouts {
    create = "90m"
    update = "3h"
    delete = "30m"
  }
}

```

## Timeouts

The `ibm_container_worker_pool_zone_attachment` provides the following [Timeouts](https://www.terraform.io/docs/language/resources/syntax.html) configuration options:

- **create**: The attachment of the zone is considered `failed` if no response is received for 90 minutes. 
- **update**: The update of the zone is considered `failed` if no response is received for 90 minutes. 
- **delete**: The detachment of the zone is considered `failed` if no response is received for 90 minutes. 


## Argument reference
Review the argument references that you can specify for your resource. 

- `cluster` - (Required, Forces new resource, String)The name or ID of the cluster that the worker pool belongs to.
- `private_vlan_id` - (Optional, String) The ID of the private VLAN that you want to use for the zone. To find available zones, run `ibmcloud ks vlans --zone <zone>`. If you do not have a private VLAN for that zone, do not specify this option. A private VLAN is automatically created for you.
- `public_vlan_id` - (Optional, String) The ID of the public VLAN that you want to use for the zone. To find available zones, run `ibmcloud ks vlans --zone <zone>`.  If you do not have a public VLAN for that zone, do not specify this option. A public VLAN is automatically created for you.
- `resource_group_id` - (Optional, Forces new resource, String) The ID of the resource group where your cluster is provisioned into. To list resource groups, run `ibmcloud resource groups` or use the `ibm_resource_group` data source.
- `wait_till_albs` - (Optional, Bool) When you add a zone to a worker pool, worker nodes are provisioned in that zone with the configuration that you defined in your worker pool. This process and enabling the ALBs on those worker nodes can take a few minutes to complete. To avoid long wait times when you run your  Terraform code, you can specify the stage when you want  Terraform to mark the zone attachment complete. Set to **true** to wait until all worker nodes are successfully provisioned in the zone that you added to your worker pool and all ALBs are available and healthy. If you want the worker node creation and ALB enablement to continue in the background, set this option to **false**. **Note**: `wait_till_albs` is set only for the first time creation of the resource, modification in the further executes will not any impacts.
- `worker_pool` - (Required, Forces new resource, String) The name or ID of the worker pool to which you want to add a zone.
- `zone` - (Required, Forces new resource, String) The name of the zone that you want to attach to the worker pool. To list available zones, run `ibmcloud ks zones`.

**Deprecated reference**

- `region` - (Deprecated, Forces new resource, string) The region where the cluster is provisioned. If the region is not specified it defaults to provider region (`IC_REGION/IBMCLOUD_REGION`). To get the list of supported regions, see [link](https://containers.bluemix.net/v1/regions) and use the alias.

## Attribute reference
In addition to all argument reference list, you can access the following attribute reference after your resource is created.

- `id` - (String) The unique identifier of the worker pool zone attachment in the format `<cluster_name_id>/< worker_pool_name_id>/<zone>`- 
- `worker_count` - (Integer) The number of worker nodes that are attached to this zone.

## Import
The `ibm_container_worker_pool_zone_attachment` can be imported by using `cluster_name_id`, `worker_pool_name_id` and `zone`.

**Example**

```
$ terraform import ibm_container_worker_pool_zone_attachment.example mycluster/5c4f4d06e0dc402084922dea70850e3b-7cafe35/dal10
```
