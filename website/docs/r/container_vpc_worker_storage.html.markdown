---
layout: "ibm"
page_title: "IBM: container_vpc_worker_storage"
sidebar_current: "docs-ibm-resource-container-vpc-worker-storage"
description: |-
  Manages IBM container vpc worker storage attachment.
---

# ibm\_container_vpc_worker_storage

Create or delete a vpc storage attachment of a vpc woker node.


## Example Usage

In the following example, you can create a storage attachment for a vpc cluster worker node:

```hcl

	resource "ibm_container_vpc_worker_volume" "volume_attach"{
		volume = "3567365d-7b9a-cc44-97ac-ef201653ea21"
		cluster = "c08evsgd0anad0v8c76g"
		worker = "kube-c08evsgd0anad0v8c76g-gen2newvpc-default-00000116"
	}
```

## Timeouts

ibm_container_vpc_worker_storage provides the following [Timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) configuration options:

* `create` - (Default 15 minutes) Used for creating storage attachment Instance.
* `delete` - (Default 10 minutes) Used for deleting storage attachment Instance.


## Argument Reference

The following arguments are supported:

* `volume` - (Required, Forces new resource, string) ID of the vpc block volume
* `cluster` - (Required, Forces new resource, string) The name or id of the cluster.
* `worker` - (Required, Forces new resource, string) The ID of the vpc cluster worker node.

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the worker storage resource. The id is composed of \<cluster_name_id\>/\<worker_id\>\<volume_attachment_id\>.<br/>
* `volume_attachment_id` - The volume attachment ID
* `volume_attachment_name` - The volume attachment name
* `status` - The volume attachment status
* `volume_type` - The volume attachment type

## Import

ibm_container_vpc_worker_storage can be imported using cluster_name_id, worker_id and volume_attachment_id eg;

```
$ terraform import ibm_container_vpc_worker_storage.example mycluster/kube-c08evsgd0anad0v8c76g-gen2newvpc-default-00000116/5c4f4d06e0dc402084922dea70850e3b-7cafe35
