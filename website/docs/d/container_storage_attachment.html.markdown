---
subcategory: "Kubernetes Service"
layout: "ibm"
page_title: "IBM: container_vpc_worker_storage"
description: |-
  Fetch the information about IBM container vpc worker storage attachment.
---

# ibm_container_storage_attachment

Import the details of a VPC storage volume attachment of a VPC worker node as a read-only data source. You can then reference the fields of the data source in other resources within the same configuration using interpolation syntax. For more information, about VPC storage volume attachment, see [Attaching a block storage volume](https://cloud.ibm.com/docs/vpc?topic=vpc-attaching-block-storage&interface=ui).


## Example usage

In the following example, you can import a storage attachment of a VPC cluster worker node:

```terraform

data "ibm_container_storage_attachment" "volume_attach"{
	volume_attachment_id = "3567365d-7b9a-cc44-97ac-ef201653ea21"
	cluster = "tf-cluster"
	worker = "kube-c08evsgd0anad0v8c76g-gen2newvpc-default-00000116"
}
```

## Argument reference

Review the argument references that you can specify for your data source.

* `cluster` - (Required, String) The name or ID of the cluster.
* `volume_attachment_id` - (Required, String) The VPC volume attachment ID.
* `worker` - (Required, String) The VPC cluster worker node ID.

## Attribute reference

In addition to all argument reference list, you can access the following attribute references after your data source is created.

* `id` - (String) The unique identifier of the worker storage resource. The id is composed of <cluster_name_id>/<worker_id><volume_attachment_id>.
* `status` - (String) The volume attachment status.
* `volume` - (String) The VPC volume ID.
* `volume_attachment_name` - (String) The volume attachment name.
* `volume_type` - (String) The volume attachment type.
