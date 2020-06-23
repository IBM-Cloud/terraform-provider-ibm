# IBM VPC Gen2 Cluster example

This example shows how to create a Kubernetes VPC Gen-2 Cluster under a specified resource group id, with default worker node with given zone and subnets. To have a multizone cluster, update the zones with new zone-name and subnet-id. 
 
Note : To create a Gen-2 VPC cluster user need to set the generation parameter inside provider blcok to 2 or export the environment varaibale IC_GENERATION as value 2. By default, value of IC_GENERATION is 2.

Following types of resources are supported:

* [ VPC Gen-2 Cluster Resource ](https://cloud.ibm.com/docs/terraform?topic=terraform-container-resources#vpc-gen2)


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

provider "ibm" {
  generation = 2
}

resource "ibm_is_vpc" "vpc1" {
  name = "vpc"
}

data "ibm_resource_group" "resource_group" {
  name = var.resource_group
}

resource "ibm_is_subnet" "subnet1" {
  name                     = "subnet-1"
  vpc                      = ibm_is_vpc.vpc1.id
  zone                     = var.zone
  total_ipv4_address_count = 256
}

resource "ibm_container_vpc_cluster" "cluster" {
  name              = var.name
  vpc_id            = ibm_is_vpc.vpc1.id
  flavor            = var.flavor
  worker_count      = var.worker_count
  resource_group_id = data.ibm_resource_group.resource_group.id

  zones {
    subnet_id = ibm_is_subnet.subnet1.id
    name      = "us-south-1"
  }
}
```

```hcl
data "ibm_container_vpc_cluster" "cluster" {
  cluster_name_id   = "vpccluster"
  resource_group_id = data.ibm_resource_group.group.id
}
```

## Examples

* [ VPC Gen-2 Cluster  ](https://github.com/IBM-Cloud/terraform-provider-ibm/tree/master/examples/ibm-cluster/vpc-gen2-cluster)

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

## Outputs

| Name | Description |
|------|-------------|
| cluster_config_file_path | Path where cluster config file is written to. |