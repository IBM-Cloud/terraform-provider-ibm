# IBM Container Worker Pool Zone example

This example shows how to create a Kubernetes Cluster under a specified resource group id, with a default worker pool with 2 workers, edit the default worker pool to add a new zone to it, add a worker pool with different zone with 2 workers and binds a service instance to a cluster.

Following types of resources are supported:

* [ Container Worker Pool Zone Attachement](https://cloud.ibm.com/docs/terraform?topic=terraform-container-resources#container-pool-zone)


## Terraform versions

Terraform 0.12. Pin module version to `~> v1.7.1`. Branch - `master`.

Terraform 0.11. Pin module version to `~> v0.29.1`. Branch - `terraform_v0.11.x`.

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
resource "ibm_is_vpc" "vpc1" {
  name = "vpc"
}

resource "ibm_is_subnet" "subnet1" {
  name                     = "subnet-1"
  vpc                      = ibm_is_vpc.vpc1.id
  zone                     = "us-south-1"
  total_ipv4_address_count = 256
}

data "ibm_resource_group" "resource_group" {
  name = var.name
}

resource "ibm_container_cluster" "cluster" {
  name              = "mycluster"
  datacenter        = "dal12"
  no_subnet         = true
  subnet_id         = [ibm_is_subnet.subnet1.id]
  default_pool_size = 2
  hardware          = "shared"
  resource_group_id = data.ibm_resource_group.resource_group.id
  machine_type      = "u2c.2x4"
}

resource "ibm_container_worker_pool" "test_pool" {
  worker_pool_name = "mypool"
  machine_type     = "u2c.2x4"
  cluster          = ibm_container_cluster.cluster.id
  size_per_zone    = 2
  hardware         = "shared"
  disk_encryption  = "true"
  labels {
    "test" = "test-pool"

    "test1" = "test-pool1"
  }
}
resource "ibm_container_worker_pool_zone_attachment" "test_zone" {
  cluster         = "mycluster"
  worker_pool     = element(split("/",ibm_container_worker_pool.test_pool.id),1)
  zone            = "dal12"
}
```

## Examples

* [ VPC Classic Cluster  ](https://github.com/umarali-nagoor/terraform-provider-ibm/tree/v12_iks_openshift_example_update/examples/ibm-cluster/cluster-worker-pool-zone)


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
| datacenter| The datacenter where you want to provision the worker nodes.  | `string` | yes |
| machine_type | The machine type for your worker node.   | `string` | yes |
| name | Name of the resource group. | `string` | yes |
| service_instance_name | The name of the service that you want to bind to the cluster.  | `string` | no |
| private_vlan_id | The ID of the private VLAN that you want to use for the zone.| `string` | no |
| public_vlan_id | The ID of the public VLAN that you want to use for the zone.  | `string` | no |
| subnet_id | The ID of an existing subnet that you want to use for your worker nodes.  | `string` | yes |
| worker_pool_name | The name of the worker pool. | `string` | no |
| zone | The name of the zone that you want to attach to the worker pool.  | `string` | yes |

## Outputs

| Name | Description |
|------|-------------|
| cluster_config_file_path | Path where cluster config file is written to. |