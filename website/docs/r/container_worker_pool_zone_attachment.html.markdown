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
resource "ibm_container_worker_pool_zone_attachment" "test_zone" {
  cluster         = "my_cluster"
  worker_pool     = "my_cluster_pool"
  zone            = "dal12"
  private_vlan_id = "2320267"
  public_vlan_id  = "2320265"
  region          = "eu-de"

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

* `zone` - (Required, string) The name of the zone. To list available zones, run 'bx cs zones'
* `cluster` - (Required, string) The name or id of the cluster.
* `worker_pool` - (Required, string) The name or id of the worker pool.
* `private_vlan_id` - (Required, string) The private VLAN of the worker node. You can retrieve the value by running the `bx cs vlans <data-center>` command in the IBM Cloud CLI.
* `public_vlan_id` - (Optional, string) The public VLAN of the worker node. You can retrieve the value by running the `bx cs vlans <data-center>` command in the IBM Cloud CLI..
* `region` - (Optional, string) The region where the cluster is provisioned. If the region is not specified it will be defaulted to provider region(BM_REGION/BLUEMIX_REGION). To get the list of supported regions please access this [link](https://containers.bluemix.net/v1/regions) and use the alias.


## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the worker pool zone attachment resource. The id is composed of \<cluster_name_id\>/\< worker_pool_name_id\>/\<zone/>
* `worker_count` - Number of workers attached to this zone.

## Import

ibm_container_worker_pool_zone_attachment can be imported using cluster_name_id, worker_pool_name_id and zone, eg

```
$ terraform import ibm_container_worker_pool_zone_attachment.example mycluster/5c4f4d06e0dc402084922dea70850e3b-7cafe35/dal10