# IBM VPC Cluster Volume Attchment 

This example shows how to attach block storage volume to VPC Cluster's worker node. To attach storage volume blocks to different worker nodes, change the worker node ID.
 
Note : 
1. When a VPC storage volume is attached to the worker node, It can not be shared/attached to the other worker nodes.
2. The worker node and the storage volume should be in the same zone.

Following types of resources are supported:

* [ VPC Worker Storage ](https://cloud.ibm.com/docs/terraform?topic=terraform-container-resources#vpc-gen2)


## Terraform versions

Terraform 0.12. Pin module version to `~> v1.21.0`. Branch - `master`.

## Usage

To run this example you need to execute:

```bash
$ terraform init
$ terraform plan
$ terraform apply
```

Run `terraform destroy` when you don't need these resources.

## Example Usage

Create a container cluster:

```hcl

provider "ibm" {
  generation = 2
}

resource "ibm_is_vpc" "vpc1" {
  name = "testvpc"
}

resource "ibm_is_subnet" "subnet1" {
  name                     = "test-subnet1"
  vpc                      = ibm_is_vpc.vpc1.id
  zone                     = "us-south-1"
  total_ipv4_address_count = 256
}

resource "ibm_is_subnet" "subnet2" {
  name                     = "test-subnet2"
  vpc                      = ibm_is_vpc.vpc1.id
  zone                     = "us-south-2"
  total_ipv4_address_count = 256
}

data "ibm_resource_group" "resource_group" {
  name = "Default"
}

resource "ibm_container_vpc_cluster" "cluster" {
  name              = "testcluster"
  vpc_id            = ibm_is_vpc.vpc1.id
  kube_version      = "1.18"
  flavor            = "bx2.2x8"
  worker_count      = 3
  resource_group_id = data.ibm_resource_group.resource_group.id

  zones {
    subnet_id = ibm_is_subnet.subnet1.id
    name      = "us-south-1"
  }
}

resource "ibm_is_volume" "storage_block"{
    name = var.volume_name
    profile = var.volume_profile
    zone = "us-south-1"
}

data "ibm_container_vpc_cluster" "cluster_info"{
    name = ibm_container_vpc_cluster.cluster.name
}

resource "ibm_container_storage_attachment" "volume_attach"{
    volume = ibm_is_volume.storage_block.id
    cluster = ibm_container_vpc_cluster.cluster.id
    worker = data.ibm_container_vpc_cluster.cluster_info.workers[0]
}

```

```hcl
data "ibm_container_storage_attachment" "cluster" {
	volume_attachment_id = "3567365d-7b9a-cc44-97ac-ef201653ea21"
	cluster = "tf-cluster"
	worker = "kube-c08evsgd0anad0v8c76g-gen2newvpc-default-00000116"
}
```

## Examples

* [ VPC Cluster Volume attachment  ](https://github.com/Mavrickk3/terraform-provider-ibm/tree/master/examples/ibm-cluster/vpc-cluster-vol-attachment)

<!-- BEGINNING OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
## Requirements

| Name | Version |
|------|---------|
| terraform | ~> 0.12 |

## Providers

| Name | Version |
|------|---------|
| ibm | n/a |

## Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| name | Name of the cluster. | `string` | yes |
| flavor | The flavor of the VPC worker node that you want to use. | `string` | yes |
| worker\_count | The number of worker nodes per zone in the default worker pool. Default value `1`.| `integer` | no |
| zone | Name of the zone.| `string` | yes |
| resource\_group | Name of the resource group.| `string` | yes |
| volume\_name | Name of the storage volume. | `string` | yes |
| volume\_profile | Type of the volume profile. | `string` | yes |

## Outputs

| Name | Description |
|------|-------------|
| volume_attachment_id | ID of the volume attachment |
| volume_attachment_status | Status of the volume attachment |
