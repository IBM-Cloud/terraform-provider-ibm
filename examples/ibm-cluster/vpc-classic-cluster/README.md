# IBM VPC Classic Cluster example

This example shows how to create a Kubernetes VPC Cluster under a specified resource group id, with default worker node with given zone and subnets. To have a multizone cluster, update the zones with new zone-name and subnet-id. 

Following types of resources are supported:

* [ VPC Classic Cluster Resource](https://cloud.ibm.com/docs/terraform?topic=terraform-container-resources#vpc-cluster)


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

data "ibm_resource_group" "resource_group" {
  name = var.resource_group
}

resource "ibm_is_subnet" "subnet1" {
  name                     = "subnet-1"
  vpc                      = ibm_is_vpc.vpc1.id
  zone                     = "us-south-1"
  total_ipv4_address_count = 256
}

resource "ibm_resource_instance" "kms_instance1" {
    name              = "test_kms"
    service           = "kms"
    plan              = "tiered-pricing"
    location          = "us-south"
}
  
resource "ibm_kms_key" "test" {
    instance_id = "${ibm_resource_instance.kms_instance1.guid}"
    key_name = "test_root_key"
    standard_key =  false
    force_delete = true
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

  kms_config {
    instance_id = ibm_resource_instance.kms_instance1.guid
    crk_id = ibm_kms_key.test.id
    private_endpoint = false
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

* [ VPC Classic Cluster  ](https://github.com/umarali-nagoor/terraform-provider-ibm/tree/v12_iks_openshift_example_update/examples/ibm-cluster/vpc-classic-cluster)

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
| resource_group | The ID of the resource group. | `string` | yes |

## Outputs

| Name | Description |
|------|-------------|
| cluster_config_file_path | Path where cluster config file is written to. |