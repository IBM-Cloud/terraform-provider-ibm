---

subcategory: "Kubernetes Service"
layout: "ibm"
page_title: "IBM: container_vpc_worker_storage"
description: |-
  Manages IBM Cloud Container VPC worker storage attachment
---

# ibm_container_storage_attachment

Create, update, or delete a VPC storage attachment of a VPC worker node. For more information, about VPC storage attachment, see [Attaching a block storage volume](https://cloud.ibm.com/docs/vpc?topic=vpc-attaching-block-storage&interface=ui).


## Example usage

In the following example, you can create a storage attachment for a VPC cluster worker node:

```terraform
provider "ibm" {
	region ="us-south"
}	

data "ibm_resource_group" "resource_group" {
	is_default = "true"
}
		
resource "ibm_is_vpc" "vpc" {
	name = "vpc"
}
		
resource "ibm_is_subnet" "subnet" {
	name                     = "subnet"
	vpc                      = ibm_is_vpc.vpc.id
	zone                     = "us-south-1"
	total_ipv4_address_count = 256
}
		
resource "ibm_container_vpc_cluster" "cluster" {
	name              = "cluster"
	vpc_id            = ibm_is_vpc.vpc.id
	flavor            = "cx2.2x4"
	worker_count      = 1
	wait_till         = "OneWorkerNodeReady"
	resource_group_id = data.ibm_resource_group.resource_group.id
	zones {
		subnet_id = ibm_is_subnet.subnet.id
		name      = "us-south-1"
	}
			
    worker_labels = {
	    "test"  = "test-default-pool"
		"test1" = "test-default-pool1"
		"test2" = "test-default-pool2"
	}			
}

resource "ibm_is_volume" "storage"{
	name = "volume"
	profile = "10iops-tier"
	zone = "us-south-1"
}

data "ibm_container_vpc_cluster" "cluster" {
	name = ibm_container_vpc_cluster.cluster.id
}

resource "ibm_container_storage_attachment" "volume_attach"{
	volume = ibm_is_volume.storage.id
	cluster = ibm_container_vpc_cluster.cluster.id
	worker = data.ibm_container_vpc_cluster.cluster.workers[0]
}
```

## Timeouts

The ibm_container_storage_attachment provides the following [Timeouts](https://www.terraform.io/docs/configuration/resources.html#timeouts) configuration options:

* `Create` - (Default 15 minutes) Used for creating storage attachment Instance.
* `Delete` - (Default 10 minutes) Used for deleting storage attachment Instance.


## Argument reference

Review the argument references that you can specify for your resource.

* `cluster` - (Required, Forces new resource, String) The name or ID of the cluster.
* `volume` - (Required, Forces new resource, String) The ID of the VPC block volume.
* `worker` - (Required, Forces new resource, String) The ID of the VPC cluster worker node.


## Attribute reference

In addition to all argument reference list, you can access the following attribute reference after your resource is created.

* `id` - The unique identifier of the worker storage resource. The id is composed of <cluster_name_id>/<worker_id><volume_attachment_id>.
* `status` - (String) The volume attachment status.
* `volume_attachment_id` - (String) The volume attachment ID.
* `volume_attachment_name` - (String) The volume attachment name.
* `volume_type` - (String) The volume attachment type.

## Import

The ibm_container_storage_attachment can be imported using `cluster_name_id`, `worker_id` and `volume_attachment_id`

Example

```
$ terraform import ibm_container_storage_attachment.example mycluster/kube-c08evsgd0anad0v8c76g-gen2newvpc-default-00000116/5c4f4d06e0dc402084922dea70850e3b-7cafe35
```