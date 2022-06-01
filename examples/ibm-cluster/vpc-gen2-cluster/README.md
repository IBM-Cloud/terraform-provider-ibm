# IBM Cloud VPC Gen 2 cluster example

This example shows how to create a Kubernetes VPC Gen-2 Cluster under a specified resource group ID, in a default worker node with given zone and subnets. To have a multizone cluster, update the zones with new zone-name and subnet-id. 
 
Following types of resources are supported:

* [VPC Gen 2 cluster resource](https://cloud.ibm.com/docs/ibm-cloud-provider-for-terraform?topic=ibm-cloud-provider-for-terraform-index-of-terraform-on-ibm-cloud-resources-and-data-sources#vpc-infrastructure_rd)

## Usage

To run this example you need to execute:

```sh
$ terraform init
$ terraform plan
$ terraform apply
```

Run `terraform destroy` when you don't need these resources.

## Example usage

Create a container cluster:

```terraform
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

```terraform
data "ibm_container_vpc_cluster" "cluster" {
  cluster_name_id   = "vpccluster"
  resource_group_id = data.ibm_resource_group.group.id
}
```

## Examples

* [VPC Gen 2 cluster](https://github.com/IBM-Cloud/terraform-provider-ibm/tree/master/examples/ibm-cluster/vpc-gen2-cluster)

<!-- BEGINNING OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
## Requirements

| Name | Version |
|------|---------|
| terraform | >=1.0.0, <2.0 |

## Providers

| Name | Version |
|------|---------|
| ibm  | latest |

## Inputs

| Name | Description | Type | Required |
|------|-------------|------|---------|
| name | Name of the cluster. | `string` | yes |
| flavor | The flavor of the VPC worker node that you want to use. | `string` | yes |
| worker_count | The number of worker nodes per zone in the default worker pool. Default value `1`.| `integer` | no |
| zone | Name of the zone.| `string` | yes |
| resource_group | Name of the resource group.| `string` | yes |
{: caption="inputs"}

## Outputs

| Name | Description |
|------|-------------|
| cluster_config_file_path | Path where cluster config file is written to. |
{: caption="outputs"}
