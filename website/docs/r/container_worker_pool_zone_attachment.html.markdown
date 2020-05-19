---
layout: "ibm"
page_title: "IBM: container_worker_pool_zone_attachment"
sidebar_current: "docs-ibm-resource-container-worker-pool-zone-attachment"
description: |-
  Manages IBM container worker pool zone attachment.
---

# ibm\_container_worker_pool_zone_attachment

Create, update, or delete a zone. This resource creates the zone and attaches it to the specified worker pool.

## Example Usage

In the following example, you can create a zone:

```hcl
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

ibm_container_worker_pool_zone_attachment provides the following [Timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) configuration options:

* `create` - (Default 90 minutes) Used for creating Instance.
* `update` - (Default 90 minutes) Used for updating Instance.
* `delete` - (Default 90 minutes) Used for deleting Instance.

## Argument Reference

The following arguments are supported:

* `zone` - (Required, Forces new resource, string) The name of the zone. To list available zones, run `ibmcloud ks zones`
* `cluster` - (Required, Forces new resource, string) The name or id of the cluster.
* `worker_pool` - (Required, Forces new resource, string) The name or id of the worker pool.
* `private_vlan_id` - (Optional, string) The private VLAN of the worker node. You can retrieve the value by running the `ibmcloud ks vlans <zone>` command in the IBM Cloud CLI.
* `public_vlan_id` - (Optional, string) The public VLAN of the worker node. The public vlan id cannot be specified alone, it should be specified along with the private vlan id. You can retrieve the value by running the `ibmcloud ks vlans <zone>` command in the IBM Cloud CLI.
**Note**: If you do not have a private or public VLAN in that zone, do not specify `private_vlan_id` and `public_vlan_id`. A private and a public VLAN are automatically created for you when you initially add a new zone to your worker pool.
* `region` - (Deprecated, Forces new resource, string) The region where the cluster is provisioned. If the region is not specified it will be defaulted to provider region(IC_REGION/IBMCLOUD_REGION). To get the list of supported regions please access this [link](https://containers.bluemix.net/v1/regions) and use the alias.
* `resource_group_id` - (Optional, Forces new resource, string) The ID of the resource group.  You can retrieve the value from data source `ibm_resource_group`. If not provided defaults to default resource group.
* `wait_till_albs` - (Optional, Bool) The woker-pool zone attachment adds the zone specified to the worker-pool. Post addition of zone, resource waits for added nodes to be normal and albs to be avialbe in the added zone. To avoid the longer wait times for resource execution, this field is introduced.
Resource will wait for ALBs to availbale in the zone added if attribute is set to true.

**NOTE**: 
1. `wait_till_albs` is set only for the first time creation of the resource, modification in the further runs will not any impacts.

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the worker pool zone attachment resource. The id is composed of \<cluster_name_id\>/\< worker_pool_name_id\>/\<zone/>
* `worker_count` - Number of workers attached to this zone.

## Import

ibm_container_worker_pool_zone_attachment can be imported using cluster_name_id, worker_pool_name_id and zone, eg

```
$ terraform import ibm_container_worker_pool_zone_attachment.example mycluster/5c4f4d06e0dc402084922dea70850e3b-7cafe35/dal10
