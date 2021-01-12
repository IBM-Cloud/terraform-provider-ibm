---
layout: "ibm"
page_title: "IBM: container_vpc_worker_storage"
sidebar_current: "docs-ibm-resource-container-vpc-worker-storage"
description: |-
  Get information about IBM container vpc worker storage attachment.
---

# ibm\_container_vpc_worker_storage

Import the details of a vpc storage volume attachment of a vpc woker node as a read-only data source.You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax.


## Example Usage

In the following example, you can import a storage attachment of a vpc cluster worker node:

```hcl

	data "ibm_container_vpc_worker_volume" "volume_attach"{
		volume_attachment_id = "3567365d-7b9a-cc44-97ac-ef201653ea21"
		cluster = "tf-cluster"
		worker = "kube-c08evsgd0anad0v8c76g-gen2newvpc-default-00000116"
	}
```

## Argument Reference

The following arguments are supported:

* `volume_attachment_id` - (Required, string) ID of the vpc volume attachment.
* `cluster` - (Required, string) The name or id of the cluster.
* `worker` - (Required, string) The ID of the vpc cluster worker node.

## Attribute Reference

The following attributes are exported:

* `id` - The unique identifier of the worker storage resource. The id is composed of \<cluster_name_id\>/\<worker_id\>\<volume_attachment_id\>.<br/>
* `volume` - The VPC volume ID
* `volume_attachment_name` - The volume attachment name
* `status` - The volume attachment status
* `volume_type` - The volume attachment type
